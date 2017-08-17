// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget2.g4 by ANTLR 4.7.

package parser // Nugget2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Nugget2Listener is a complete listener for a parse tree produced by Nugget2Parser.
type Nugget2Listener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDefine_assign is called when entering the define_assign production.
	EnterDefine_assign(c *Define_assignContext)

	// EnterDefine is called when entering the define production.
	EnterDefine(c *DefineContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterAsType is called when entering the asType production.
	EnterAsType(c *AsTypeContext)

	// EnterNugget_type is called when entering the nugget_type production.
	EnterNugget_type(c *Nugget_typeContext)

	// EnterNuggetaction is called when entering the nuggetaction production.
	EnterNuggetaction(c *NuggetactionContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDefine_assign is called when exiting the define_assign production.
	ExitDefine_assign(c *Define_assignContext)

	// ExitDefine is called when exiting the define production.
	ExitDefine(c *DefineContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitAsType is called when exiting the asType production.
	ExitAsType(c *AsTypeContext)

	// ExitNugget_type is called when exiting the nugget_type production.
	ExitNugget_type(c *Nugget_typeContext)

	// ExitNuggetaction is called when exiting the nuggetaction production.
	ExitNuggetaction(c *NuggetactionContext)
}
