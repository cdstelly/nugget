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

	// EnterOperation_on_singleton is called when entering the operation_on_singleton production.
	EnterOperation_on_singleton(c *Operation_on_singletonContext)

	// EnterSingleton_op is called when entering the singleton_op production.
	EnterSingleton_op(c *Singleton_opContext)

	// EnterAsType is called when entering the asType production.
	EnterAsType(c *AsTypeContext)

	// EnterNugget_type is called when entering the nugget_type production.
	EnterNugget_type(c *Nugget_typeContext)

	// EnterNugget_action is called when entering the nugget_action production.
	EnterNugget_action(c *Nugget_actionContext)

	// EnterAction_word is called when entering the action_word production.
	EnterAction_word(c *Action_wordContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterFilter_term is called when entering the filter_term production.
	EnterFilter_term(c *Filter_termContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDefine_assign is called when exiting the define_assign production.
	ExitDefine_assign(c *Define_assignContext)

	// ExitDefine is called when exiting the define production.
	ExitDefine(c *DefineContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitOperation_on_singleton is called when exiting the operation_on_singleton production.
	ExitOperation_on_singleton(c *Operation_on_singletonContext)

	// ExitSingleton_op is called when exiting the singleton_op production.
	ExitSingleton_op(c *Singleton_opContext)

	// ExitAsType is called when exiting the asType production.
	ExitAsType(c *AsTypeContext)

	// ExitNugget_type is called when exiting the nugget_type production.
	ExitNugget_type(c *Nugget_typeContext)

	// ExitNugget_action is called when exiting the nugget_action production.
	ExitNugget_action(c *Nugget_actionContext)

	// ExitAction_word is called when exiting the action_word production.
	ExitAction_word(c *Action_wordContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitFilter_term is called when exiting the filter_term production.
	ExitFilter_term(c *Filter_termContext)
}
