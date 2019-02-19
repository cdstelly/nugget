// Generated from /home/cyclops/code/nugget/Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NuggetListener is a complete listener for a parse tree produced by NuggetParser.
type NuggetListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDefine_assign is called when entering the define_assign production.
	EnterDefine_assign(c *Define_assignContext)

	// EnterDefine is called when entering the define production.
	EnterDefine(c *DefineContext)

	// EnterDefine_tuple is called when entering the define_tuple production.
	EnterDefine_tuple(c *Define_tupleContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterOperation_on_singleton is called when entering the operation_on_singleton production.
	EnterOperation_on_singleton(c *Operation_on_singletonContext)

	// EnterOutput_as is called when entering the output_as production.
	EnterOutput_as(c *Output_asContext)

	// EnterOutput_type is called when entering the output_type production.
	EnterOutput_type(c *Output_typeContext)

	// EnterSingleton_op is called when entering the singleton_op production.
	EnterSingleton_op(c *Singleton_opContext)

	// EnterSingleton_var is called when entering the singleton_var production.
	EnterSingleton_var(c *Singleton_varContext)

	// EnterNugget_type is called when entering the nugget_type production.
	EnterNugget_type(c *Nugget_typeContext)

	// EnterNugget_action is called when entering the nugget_action production.
	EnterNugget_action(c *Nugget_actionContext)

	// EnterAction_word is called when entering the action_word production.
	EnterAction_word(c *Action_wordContext)

	// EnterAsType is called when entering the asType production.
	EnterAsType(c *AsTypeContext)

	// EnterByField is called when entering the byField production.
	EnterByField(c *ByFieldContext)

	// EnterByteOffsetSize is called when entering the byteOffsetSize production.
	EnterByteOffsetSize(c *ByteOffsetSizeContext)

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

	// ExitDefine_tuple is called when exiting the define_tuple production.
	ExitDefine_tuple(c *Define_tupleContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitOperation_on_singleton is called when exiting the operation_on_singleton production.
	ExitOperation_on_singleton(c *Operation_on_singletonContext)

	// ExitOutput_as is called when exiting the output_as production.
	ExitOutput_as(c *Output_asContext)

	// ExitOutput_type is called when exiting the output_type production.
	ExitOutput_type(c *Output_typeContext)

	// ExitSingleton_op is called when exiting the singleton_op production.
	ExitSingleton_op(c *Singleton_opContext)

	// ExitSingleton_var is called when exiting the singleton_var production.
	ExitSingleton_var(c *Singleton_varContext)

	// ExitNugget_type is called when exiting the nugget_type production.
	ExitNugget_type(c *Nugget_typeContext)

	// ExitNugget_action is called when exiting the nugget_action production.
	ExitNugget_action(c *Nugget_actionContext)

	// ExitAction_word is called when exiting the action_word production.
	ExitAction_word(c *Action_wordContext)

	// ExitAsType is called when exiting the asType production.
	ExitAsType(c *AsTypeContext)

	// ExitByField is called when exiting the byField production.
	ExitByField(c *ByFieldContext)

	// ExitByteOffsetSize is called when exiting the byteOffsetSize production.
	ExitByteOffsetSize(c *ByteOffsetSizeContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitFilter_term is called when exiting the filter_term production.
	ExitFilter_term(c *Filter_termContext)
}
