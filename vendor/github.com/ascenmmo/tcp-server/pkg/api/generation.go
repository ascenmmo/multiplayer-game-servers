// @tg version=1.0.3
// @tg backend="Asenmmo"
// @tg title=`Ascenmmo Rest API`
// @tg servers=`http://stage.ascenmmo.com;stage cluster`
//
//go:generate go tool github.com/seniorGolang/tg/v2/cmd/tg transport --services . --out ../transport --outSwagger ../swagger.yaml
//go:generate go tool github.com/seniorGolang/tg/v2/cmd/tg client -go --services . --outPath ../../clients/suppliers
//go:generate goimports -l -w ../transport ../../clients/suppliers

package api
