// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget2.g4 by ANTLR 4.7.

package parser // Nugget2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNugget2Listener is a complete listener for a parse tree produced by Nugget2Parser.
type BaseNugget2Listener struct{}

var _ Nugget2Listener = &BaseNugget2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNugget2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNugget2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNugget2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNugget2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseNugget2Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseNugget2Listener) ExitProg(ctx *ProgContext) {}

// EnterDefine_assign is called when production define_assign is entered.
func (s *BaseNugget2Listener) EnterDefine_assign(ctx *Define_assignContext) {}

// ExitDefine_assign is called when production define_assign is exited.
func (s *BaseNugget2Listener) ExitDefine_assign(ctx *Define_assignContext) {}

// EnterDefine is called when production define is entered.
func (s *BaseNugget2Listener) EnterDefine(ctx *DefineContext) {}

// ExitDefine is called when production define is exited.
func (s *BaseNugget2Listener) ExitDefine(ctx *DefineContext) {}

// EnterDefine_tuple is called when production define_tuple is entered.
func (s *BaseNugget2Listener) EnterDefine_tuple(ctx *Define_tupleContext) {}

// ExitDefine_tuple is called when production define_tuple is exited.
func (s *BaseNugget2Listener) ExitDefine_tuple(ctx *Define_tupleContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseNugget2Listener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseNugget2Listener) ExitAssign(ctx *AssignContext) {}

// EnterOperation_on_singleton is called when production operation_on_singleton is entered.
func (s *BaseNugget2Listener) EnterOperation_on_singleton(ctx *Operation_on_singletonContext) {}

// ExitOperation_on_singleton is called when production operation_on_singleton is exited.
func (s *BaseNugget2Listener) ExitOperation_on_singleton(ctx *Operation_on_singletonContext) {}

// EnterSingleton_op is called when production singleton_op is entered.
func (s *BaseNugget2Listener) EnterSingleton_op(ctx *Singleton_opContext) {}

// ExitSingleton_op is called when production singleton_op is exited.
func (s *BaseNugget2Listener) ExitSingleton_op(ctx *Singleton_opContext) {}

// EnterSingleton_var is called when production singleton_var is entered.
func (s *BaseNugget2Listener) EnterSingleton_var(ctx *Singleton_varContext) {}

// ExitSingleton_var is called when production singleton_var is exited.
func (s *BaseNugget2Listener) ExitSingleton_var(ctx *Singleton_varContext) {}

// EnterNugget_type is called when production nugget_type is entered.
func (s *BaseNugget2Listener) EnterNugget_type(ctx *Nugget_typeContext) {}

// ExitNugget_type is called when production nugget_type is exited.
func (s *BaseNugget2Listener) ExitNugget_type(ctx *Nugget_typeContext) {}

// EnterNugget_action is called when production nugget_action is entered.
func (s *BaseNugget2Listener) EnterNugget_action(ctx *Nugget_actionContext) {}

// ExitNugget_action is called when production nugget_action is exited.
func (s *BaseNugget2Listener) ExitNugget_action(ctx *Nugget_actionContext) {}

// EnterAction_word is called when production action_word is entered.
func (s *BaseNugget2Listener) EnterAction_word(ctx *Action_wordContext) {}

// ExitAction_word is called when production action_word is exited.
func (s *BaseNugget2Listener) ExitAction_word(ctx *Action_wordContext) {}

// EnterAsType is called when production asType is entered.
func (s *BaseNugget2Listener) EnterAsType(ctx *AsTypeContext) {}

// ExitAsType is called when production asType is exited.
func (s *BaseNugget2Listener) ExitAsType(ctx *AsTypeContext) {}

// EnterByField is called when production byField is entered.
func (s *BaseNugget2Listener) EnterByField(ctx *ByFieldContext) {}

// ExitByField is called when production byField is exited.
func (s *BaseNugget2Listener) ExitByField(ctx *ByFieldContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseNugget2Listener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseNugget2Listener) ExitFilter(ctx *FilterContext) {}

// EnterFilter_term is called when production filter_term is entered.
func (s *BaseNugget2Listener) EnterFilter_term(ctx *Filter_termContext) {}

// ExitFilter_term is called when production filter_term is exited.
func (s *BaseNugget2Listener) ExitFilter_term(ctx *Filter_termContext) {}
