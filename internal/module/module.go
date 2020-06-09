package module

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"log"
	"path"
	"regexp"

	"github.com/commitdev/zero/internal/config"
	"github.com/commitdev/zero/internal/config/moduleconfig"
	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/internal/util"
	"github.com/hashicorp/go-getter"
)

// TemplateModule merges a module instance params with the static configs
type TemplateModule struct {
	config.ModuleInstance // @TODO Move this
	Config                moduleconfig.ModuleConfig
}

// FetchModule downloads the remote module source (or loads the local files) and parses the module config yaml
func FetchModule(source string) (moduleconfig.ModuleConfig, error) {
	config := moduleconfig.ModuleConfig{}
	localPath := GetSourceDir(source)
	if !isLocal(source) {
		err := getter.Get(localPath, source)
		if err != nil {
			return config, err
		}
	}

	configPath := path.Join(localPath, constants.ZeroModuleYml)
	config, err := moduleconfig.LoadModuleConfig(configPath)
	return config, err
}

// GetSourcePath gets a unique local source directory name. For local modules, it use the local directory
func GetSourceDir(source string) string {
	if !isLocal(source) {
		h := md5.New()
		io.WriteString(h, source)
		source = base64.StdEncoding.EncodeToString(h.Sum(nil))
		return path.Join(constants.TemplatesDir, source)
	} else {
		return source
	}
}

// IsLocal uses the go-getter FileDetector to check if source is a file
func isLocal(source string) bool {
	pwd := util.GetCwd()

	// ref: https://github.com/hashicorp/go-getter/blob/master/detect_test.go
	out, err := getter.Detect(source, pwd, getter.Detectors)

	match, err := regexp.MatchString("^file://.*", out)
	if err != nil {
		log.Panicf("invalid source format %s", err)
	}

	return match
}

func withPWD(pwd string) func(*getter.Client) error {
	return func(c *getter.Client) error {
		c.Pwd = pwd
		return nil
	}
}
