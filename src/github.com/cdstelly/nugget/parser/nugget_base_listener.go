// Generated from /home/cyclops/code/nugget/Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNuggetListener is a complete listener for a parse tree produced by NuggetParser.
type BaseNuggetListener struct{}

var _ NuggetListener = &BaseNuggetListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNuggetListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNuggetListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNuggetListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNuggetListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseNuggetListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseNuggetListener) ExitProg(ctx *ProgContext) {}

// EnterDefine_assign is called when production define_assign is entered.
func (s *BaseNuggetListener) EnterDefine_assign(ctx *Define_assignContext) {}

// ExitDefine_assign is called when production define_assign is exited.
func (s *BaseNuggetListener) ExitDefine_assign(ctx *Define_assignContext) {}

// EnterDefine is called when production define is entered.
func (s *BaseNuggetListener) EnterDefine(ctx *DefineContext) {}

// ExitDefine is called when production define is exited.
func (s *BaseNuggetListener) ExitDefine(ctx *DefineContext) {}

// EnterDefine_tuple is called when production define_tuple is entered.
func (s *BaseNuggetListener) EnterDefine_tuple(ctx *Define_tupleContext) {}

// ExitDefine_tuple is called when production define_tuple is exited.
func (s *BaseNuggetListener) ExitDefine_tuple(ctx *Define_tupleContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseNuggetListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseNuggetListener) ExitAssign(ctx *AssignContext) {}

// EnterOperation_on_singleton is called when production operation_on_singleton is entered.
func (s *BaseNuggetListener) EnterOperation_on_singleton(ctx *Operation_on_singletonContext) {}

// ExitOperation_on_singleton is called when production operation_on_singleton is exited.
func (s *BaseNuggetListener) ExitOperation_on_singleton(ctx *Operation_on_singletonContext) {}

// EnterOutput_as is called when production output_as is entered.
func (s *BaseNuggetListener) EnterOutput_as(ctx *Output_asContext) {}

// ExitOutput_as is called when production output_as is exited.
func (s *BaseNuggetListener) ExitOutput_as(ctx *Output_asContext) {}

// EnterOutput_type is called when production output_type is entered.
func (s *BaseNuggetListener) EnterOutput_type(ctx *Output_typeContext) {}

// ExitOutput_type is called when production output_type is exited.
func (s *BaseNuggetListener) ExitOutput_type(ctx *Output_typeContext) {}

// EnterSingleton_op is called when production singleton_op is entered.
func (s *BaseNuggetListener) EnterSingleton_op(ctx *Singleton_opContext) {}

// ExitSingleton_op is called when production singleton_op is exited.
func (s *BaseNuggetListener) ExitSingleton_op(ctx *Singleton_opContext) {}

// EnterSingleton_var is called when production singleton_var is entered.
func (s *BaseNuggetListener) EnterSingleton_var(ctx *Singleton_varContext) {}

// ExitSingleton_var is called when production singleton_var is exited.
func (s *BaseNuggetListener) ExitSingleton_var(ctx *Singleton_varContext) {}

// EnterNugget_type is called when production nugget_type is entered.
func (s *BaseNuggetListener) EnterNugget_type(ctx *Nugget_typeContext) {}

// ExitNugget_type is called when production nugget_type is exited.
func (s *BaseNuggetListener) ExitNugget_type(ctx *Nugget_typeContext) {}

// EnterNugget_action is called when production nugget_action is entered.
func (s *BaseNuggetListener) EnterNugget_action(ctx *Nugget_actionContext) {}

// ExitNugget_action is called when production nugget_action is exited.
func (s *BaseNuggetListener) ExitNugget_action(ctx *Nugget_actionContext) {}

// EnterAction_word is called when production action_word is entered.
func (s *BaseNuggetListener) EnterAction_word(ctx *Action_wordContext) {}

// ExitAction_word is called when production action_word is exited.
func (s *BaseNuggetListener) ExitAction_word(ctx *Action_wordContext) {}

// EnterAsType is called when production asType is entered.
func (s *BaseNuggetListener) EnterAsType(ctx *AsTypeContext) {}

// ExitAsType is called when production asType is exited.
func (s *BaseNuggetListener) ExitAsType(ctx *AsTypeContext) {}

// EnterByField is called when production byField is entered.
func (s *BaseNuggetListener) EnterByField(ctx *ByFieldContext) {}

// ExitByField is called when production byField is exited.
func (s *BaseNuggetListener) ExitByField(ctx *ByFieldContext) {}

// EnterByteOffsetSize is called when production byteOffsetSize is entered.
func (s *BaseNuggetListener) EnterByteOffsetSize(ctx *ByteOffsetSizeContext) {}

// ExitByteOffsetSize is called when production byteOffsetSize is exited.
func (s *BaseNuggetListener) ExitByteOffsetSize(ctx *ByteOffsetSizeContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseNuggetListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseNuggetListener) ExitFilter(ctx *FilterContext) {}

// EnterFilter_term is called when production filter_term is entered.
func (s *BaseNuggetListener) EnterFilter_term(ctx *Filter_termContext) {}

// ExitFilter_term is called when production filter_term is exited.
func (s *BaseNuggetListener) ExitFilter_term(ctx *Filter_termContext) {}
