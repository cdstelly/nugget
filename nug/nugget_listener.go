// Generated from Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NuggetListener is a complete listener for a parse tree produced by NuggetParser.
type NuggetListener interface {
	antlr.ParseTreeListener

	// EnterNugget is called when entering the nugget production.
	EnterNugget(c *NuggetContext)

	// EnterInitextract is called when entering the initextract production.
	EnterInitextract(c *InitextractContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// EnterStreamTask is called when entering the streamTask production.
	EnterStreamTask(c *StreamTaskContext)

	// EnterSave is called when entering the save production.
	EnterSave(c *SaveContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterTimefilter is called when entering the timefilter production.
	EnterTimefilter(c *TimefilterContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterSubtype is called when entering the subtype production.
	EnterSubtype(c *SubtypeContext)

	// EnterTask is called when entering the task production.
	EnterTask(c *TaskContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterSourceidentifier is called when entering the sourceidentifier production.
	EnterSourceidentifier(c *SourceidentifierContext)

	// EnterPrintId is called when entering the printId production.
	EnterPrintId(c *PrintIdContext)

	// EnterSin is called when entering the sin production.
	EnterSin(c *SinContext)

	// ExitNugget is called when exiting the nugget production.
	ExitNugget(c *NuggetContext)

	// ExitInitextract is called when exiting the initextract production.
	ExitInitextract(c *InitextractContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitStreamTask is called when exiting the streamTask production.
	ExitStreamTask(c *StreamTaskContext)

	// ExitSave is called when exiting the save production.
	ExitSave(c *SaveContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitTimefilter is called when exiting the timefilter production.
	ExitTimefilter(c *TimefilterContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitSubtype is called when exiting the subtype production.
	ExitSubtype(c *SubtypeContext)

	// ExitTask is called when exiting the task production.
	ExitTask(c *TaskContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitSourceidentifier is called when exiting the sourceidentifier production.
	ExitSourceidentifier(c *SourceidentifierContext)

	// ExitPrintId is called when exiting the printId production.
	ExitPrintId(c *PrintIdContext)

	// ExitSin is called when exiting the sin production.
	ExitSin(c *SinContext)
}
