// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) Home(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Home", args).URL
}


type tAuthentication struct {}
var Authentication tAuthentication


func (_ tAuthentication) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Authentication.Login", args).URL
}

func (_ tAuthentication) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Authentication.Logout", args).URL
}


type tMaterialApi struct {}
var MaterialApi tMaterialApi


func (_ tMaterialApi) Home(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MaterialApi.Home", args).URL
}

func (_ tMaterialApi) IndexMaterial(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MaterialApi.IndexMaterial", args).URL
}

func (_ tMaterialApi) GradeMaterials(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MaterialApi.GradeMaterials", args).URL
}

func (_ tMaterialApi) MyMaterials(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MaterialApi.MyMaterials", args).URL
}

func (_ tMaterialApi) SelectGrade(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MaterialApi.SelectGrade", args).URL
}

func (_ tMaterialApi) PostMaterial(
		file interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "file", file)
	return revel.MainRouter.Reverse("MaterialApi.PostMaterial", args).URL
}

func (_ tMaterialApi) File(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MaterialApi.File", args).URL
}

func (_ tMaterialApi) ViewMaterial(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MaterialApi.ViewMaterial", args).URL
}

func (_ tMaterialApi) DeleteMaterial(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MaterialApi.DeleteMaterial", args).URL
}


type tPassword struct {}
var Password tPassword


func (_ tPassword) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Password.Index", args).URL
}

func (_ tPassword) Input(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Password.Input", args).URL
}

func (_ tPassword) Password(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Password.Password", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


