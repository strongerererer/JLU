package env

type EnvService struct {
	config string
}

func NewService(envpath string) *EnvService {
	return &EnvService{config: envpath}
}

func (this *EnvService) Init() bool {

	_evnData = make(map[string]string)
	if this.config != "" {
		return Load(this.config)
	}

	return true
}

func (this *EnvService) Name() string { return "env" }

func (this *EnvService) Reload() bool { return true }

func (this *EnvService) Final() {}
