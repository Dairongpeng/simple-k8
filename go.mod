module simple-k8

go 1.12

require (
	github.com/Joker/jade v0.0.0-20161230135920-35b3f5bdbcc9 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/aymerick/raymond v0.0.0-20161209220724-72acac220747 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/eknkc/amber v0.0.0-20170716093748-b8bd8b03e4f7 // indirect
	github.com/elastic/go-ucfg v0.4.6
	github.com/esemplastic/unis v0.0.0-20170509161724-6e30ed034e8c // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20170704123420-58f1f3387f7c // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/go-sql-driver/mysql v0.0.0-20170616131631-8fefef06da77
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/httpexpect v0.0.0-20180314041918-ebe99fcebbce // indirect
	github.com/jmoiron/sqlx v0.0.0-20180124204410-05cef0741ade
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/juju/errors v0.0.0-20170703010042-c7d06af17c68 // indirect
	github.com/juju/loggo v0.0.0-20190526231331-6e530bcce5d8 // indirect
	github.com/juju/testing v0.0.0-20191001232224-ce9dec17d28b // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris v0.0.0-20170616235417-20f30022705b
	github.com/klauspost/compress v0.0.0-20170528132359-f3dce52e0576 // indirect
	github.com/klauspost/cpuid v0.0.0-20160302075316-09cded8978dc // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-sqlite3 v2.0.2+incompatible // indirect
	github.com/microcosm-cc/bluemonday v0.0.0-20161202143824-e79763773ab6 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/monoculum/formam v0.0.0-20170713212123-6292a2420ab5 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/natefinch/lumberjack v0.0.0-20170531180850-df99d62fd42d
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/russross/blackfriday v0.0.0-20170610170232-067529f716f4 // indirect
	github.com/satori/go.uuid v0.0.0-20170321230731-5bf94b69c6b6
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/urfave/cli v0.0.0-20170612000038-cf33a9befefd
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/genproto v0.0.0-20190418145605-e7d98fc518a7 // indirect
	google.golang.org/grpc v1.19.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	k8s.io/api v0.17.0 // indirect
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20191114184206-e782cd3c129f // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20191004102349-159aefb8556b
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191004105649-b14e3c49469a
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004074956-c5d2f014d689
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.3.0
)
