package main

import (
	"testing"
	"reflect"
	"go/ast"
	"go/token"
	"go/parser"
)

func getSampleAst() *ast.File {
	const sourceContent = `package main
import "fmt"
func main() {
    // This is a comment
    msg := "hello, world\n"
    fmt.Printf( msg )
	if (len(msg)) > 0 {
		fmt.Println(msg)
    }
}
`
	fileSet := token.NewFileSet()
	sourceFileName := "main.go"
	astFile, err := parser.ParseFile(fileSet, sourceFileName, sourceContent, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return astFile
}

func getSampleUast() *Node {
	return MapFile(getSampleAst())
}

func Test_mapFile(t *testing.T) {
	uast := getSampleUast()
	if expected := []Kind{COMPILATION_UNIT}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 1; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if expected := []Kind{DECL_LIST}; !reflect.DeepEqual(expected, uast.Children[0].Kinds) {
		t.Fatalf("got %v as kinds of first child; expected %v", uast.Children[0].Kinds, expected)
	}

	if expected := 1; expected != uast.Token.offset {
		t.Fatalf("got %v as Token.offset; expected %v", uast.Token.offset, expected)
	}

	if expected := "main"; expected != uast.Token.Value {
		t.Fatalf("got %v as Value; expected %v", uast.Token.Value, expected)
	}

	if expected := "*ast.File"; expected != uast.NativeNode {
		t.Fatalf("got %v as NativeValue; expected %v", uast.NativeNode, expected)
	}
}

func Test_mapFuncDecl(t *testing.T) {
	funcDecl := getSampleAst().Decls[1].(*ast.FuncDecl)
	uast := mapFuncDecl(funcDecl)

	if expected := []Kind{FUNCTION}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 2; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if uast.Token != nil {
		t.Fatalf("got %v as Token; expected nil", uast.Token)
	}

	if expected := "*ast.FuncDecl"; expected != uast.NativeNode {
		t.Fatalf("got %v as NativeValue; expected %v", uast.NativeNode, expected)
	}
}

func Test_mapFuncDecl_Name(t *testing.T) {
	funcDecl := getSampleAst().Decls[1].(*ast.FuncDecl)
	uast := mapFuncDecl(funcDecl).Children[0]

	if expected := []Kind{IDENTIFIER}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 0; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if expected := 32; expected != uast.Token.offset {
		t.Fatalf("got %v as Token.offset; expected %v", uast.Token.offset, expected)
	}

	if expected := "main"; expected != uast.Token.Value {
		t.Fatalf("got %v as Value; expected %v", uast.Token.Value, expected)
	}

	if expected := "*ast.Ident"; expected != uast.NativeNode {
		t.Fatalf("got %v as NativeValue; expected %v", uast.NativeNode, expected)
	}
}

func Test_mapAssignStmt(t *testing.T) {
	blockStmt := getSampleAst().Decls[1].(*ast.FuncDecl).Body
	uast := mapAssignStmt(blockStmt.List[0].(*ast.AssignStmt))

	if expected := []Kind{ASSIGNMENT}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 3; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if uast.Token != nil {
		t.Fatalf("got %v as Token; expected nil", uast.Token)
	}

	if expected := "*ast.AssignStmt"; expected != uast.NativeNode {
		t.Fatalf("got %v as NativeValue; expected %v", uast.NativeNode, expected)
	}
}

func Test_mapExprList(t *testing.T) {
	blockStmt := getSampleAst().Decls[1].(*ast.FuncDecl).Body
	uast := mapExprList(EXPR_LIST, blockStmt.List[0].(*ast.AssignStmt).Lhs)

	if expected := []Kind{EXPR_LIST}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 1; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if uast.Token != nil {
		t.Fatalf("got %v as Token; expected nil", uast.Token)
	}

	if expected := "[]ast.Expr"; expected != uast.NativeNode {
		t.Fatalf("got %v as NativeValue; expected %v", uast.NativeNode, expected)
	}
}

func Test_mapExpr_Ident(t *testing.T) {
	blockStmt := getSampleAst().Decls[1].(*ast.FuncDecl).Body
	uast := mapExpr(blockStmt.List[0].(*ast.AssignStmt).Lhs[0])

	if uast == nil {
		t.Fatalf("got nil; expected an identifier")
	}

	if expected := []Kind{IDENTIFIER}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 0; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if expected := 70; expected != uast.Token.offset {
		t.Fatalf("got %v as Token.offset; expected %v", uast.Token.offset, expected)
	}

	if expected := "msg"; expected != uast.Token.Value {
		t.Fatalf("got %v as Value; expected %v", uast.Token.Value, expected)
	}

	if expected := "*ast.Ident"; expected != uast.NativeNode {
		t.Fatalf("got %v as NativeValue; expected %v", uast.NativeNode, expected)
	}
}

func Test_mapExpr_BasicLit(t *testing.T) {
	blockStmt := getSampleAst().Decls[1].(*ast.FuncDecl).Body
	uast := mapExpr(blockStmt.List[0].(*ast.AssignStmt).Rhs[0])

	if uast == nil {
		t.Fatalf("got nil; expected a literal")
	}

	if expected := []Kind{LITERAL}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 0; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if expected := 77; expected != uast.Token.offset {
		t.Fatalf("got %v as Token.offset; expected %v", uast.Token.offset, expected)
	}

	if expected := "\"hello, world\\n\""; expected != uast.Token.Value {
		t.Fatalf("got %v as Value; expected %v", uast.Token.Value, expected)
	}

	if expected := "*ast.BasicLit"; expected != uast.NativeNode {
		t.Fatalf("got %v as NativeValue; expected %v", uast.NativeNode, expected)
	}
}

func Test_mapExprStmt(t *testing.T) {
	blockStmt := getSampleAst().Decls[1].(*ast.FuncDecl).Body
	uast := mapExprStmt(blockStmt.List[1].(*ast.ExprStmt))

	if expected := []Kind{EXPR_STMT}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 1; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if uast.Token != nil {
		t.Fatalf("got %v as Token; expected nil", uast.Token)
	}
}

func Test_mapCallExpr(t *testing.T) {
	blockStmt := getSampleAst().Decls[1].(*ast.FuncDecl).Body
	uast := mapCallExpr(blockStmt.List[1].(*ast.ExprStmt).X.(*ast.CallExpr))

	if expected := []Kind{CALL}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 4; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if uast.Token != nil {
		t.Fatalf("got %v as Token; expected nil", uast.Token)
	}
}

func Test_mapIfStmt(t *testing.T) {
	blockStmt := getSampleAst().Decls[1].(*ast.FuncDecl).Body
	uast := mapIfStmt(blockStmt.List[2].(*ast.IfStmt))

	if expected := []Kind{IF_STMT}; !reflect.DeepEqual(expected, uast.Kinds) {
		t.Fatalf("got %v as Kinds; expected %v", uast.Kinds, expected)
	}

	if expected := 4; expected != len(uast.Children) {
		t.Fatalf("got %v as number of Children; expected %v", len(uast.Children), expected)
	}

	if uast.Token != nil {
		t.Fatalf("got %v as Token; expected nil", uast.Token)
	}
}