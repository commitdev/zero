package module

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"log"
	"path"
	"regexp"
	"sync"

	"github.com/commitdev/zero/internal/config/moduleconfig"
	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/hashicorp/go-getter"
)

// FetchModule downloads the remote module source if necessary. Meant to be run in a goroutine.
func FetchModule(source string, wg *sync.WaitGroup) {
	defer wg.Done()

	localPath := GetSourceDir(source)
	if !IsLocal(source) {
		flog.Debugf("Downloading module: %s to %s", source, localPath)
		err := getter.Get(localPath, source)
		if err != nil {
			exit.Fatal("Failed to fetch remote module from %s: %v\n", source, err)
		}
	}
	return
}

// ParseModuleConfig loads the local config file for a module and parses the yaml
func ParseModuleConfig(source string) (moduleconfig.ModuleConfig, error) {
	localPath := GetSourceDir(source)
	config := moduleconfig.ModuleConfig{}
	configPath := path.Join(localPath, constants.ZeroModuleYml)
	config, err := moduleconfig.LoadModuleConfig(configPath)
	return config, err
}

// GetSourcePath gets a unique local source directory name. For local modules, it use the local directory
func GetSourceDir(source string) string {
	if !IsLocal(source) {
		h := md5.New()
		io.WriteString(h, source)
		source = base64.StdEncoding.EncodeToString(h.Sum(nil))
		return path.Join(constants.TemplatesDir, source)
	} else {
		return source
	}
}

// IsLocal uses the go-getter FileDetector to check if source is a file
func IsLocal(source string) bool {
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
