package module

import (
	"crypto/md5"
	"io"
	"log"
	"os"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/hashicorp/go-getter"
	"github.com/manifoldco/promptui"
)

type TemplateModule struct {
	config.Module
	Config config.ModuleConfig
}

type ProgressTracking struct {
	sync.Mutex
	downloaded map[string]int
}

// Init downloads the remote template files and parses the module config yaml
func NewTemplateModule(moduleCfg config.Module) (*TemplateModule, error) {
	var templateModule TemplateModule
	templateModule.Source = moduleCfg.Source
	templateModule.Params = moduleCfg.Params

	p := &ProgressTracking{}
	sourcePath := templateModule.GetSourceDir()

	if !templateModule.IsLocal() {
		err := getter.Get(sourcePath, templateModule.Source, getter.WithProgress(p))
		if err != nil {
			return nil, err
		}
	}

	config.LoadModuleConfig(sourcePath, templateModule.Config)

	return &templateModule, nil
}

// PromptParams renders series of prompt UI based on the config
func (m *TemplateModule) PromptParams() error {
	for _, promptConfig := range m.Config.Prompts {

		label := promptConfig.Label
		if promptConfig.Label == "" {
			label = promptConfig.Field
		}
		var err error
		var result string

		if len(promptConfig.Options) > 0 {
			prompt := promptui.Select{
				Label: label,
				Items: promptConfig.Options,
			}
			_, result, err = prompt.Run()

		} else {
			prompt := promptui.Prompt{
				Label: label,
			}
			result, err = prompt.Run()
		}
		if err != nil {
			return err
		}

		m.Params[promptConfig.Field] = result
	}

	return nil
}

// GetSourcePath gets a unique local source directory name. For local modules, it use the local directory
func (m *TemplateModule) GetSourceDir() string {
	identifier := m.Source

	// paths := strings.SplitN(m.Source, "/", -1)
	// re, _ := regexp.Compile("[^a-zA-Z0-9]+")
	// return re.ReplaceAllString(strings.Join(idPaths, "_"), "")
	h := md5.New()
	io.WriteString(h, m.Source)
	identifier = string(h.Sum(nil))

	return "templates/" + identifier
}

func (m *TemplateModule) IsLocal() bool {
	f := new(getter.FileDetector)

	pwd, err := os.Getwd()
	if err != nil {
		log.Panicf("failed to get current working directory: %v", err)
	}

	_, ok, err := f.Detect(m.Source, pwd)
	if ok && err != nil {
		return true
	} else {
		return false
	}
}

// func withPWD(pwd string) func(*getter.Client) error {
// 	return func(c *getter.Client) error {
// 		c.Pwd = pwd
// 		return nil
// 	}
// }

func (p *ProgressTracking) TrackProgress(src string, currentSize, totalSize int64, stream io.ReadCloser) (body io.ReadCloser) {
	p.Lock()
	defer p.Unlock()

	if p.downloaded == nil {
		p.downloaded = map[string]int{}
	}

	v, _ := p.downloaded[src]
	p.downloaded[src] = v + 1
	return stream
}
