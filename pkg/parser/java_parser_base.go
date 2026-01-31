// it's manually converted from https://github.com/antlr/grammars-v4/blob/master/java/java/Java/JavaParserBase.java

package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

// JavaParserBase corresponds to JavaParserBase.java
// (external base class expected by the grammar)
type JavaParserBase struct {
	*antlr.BaseParser
}

func NewJavaParserBase(input antlr.TokenStream) *JavaParserBase {
	return &JavaParserBase{
		BaseParser: antlr.NewBaseParser(input),
	}
}

// DoLastRecordComponent replicates JavaParserBase.DoLastRecordComponent()
func (p *JavaParserBase) DoLastRecordComponent() bool {
	ctx := p.GetParserRuleContext()
	tctx, ok := ctx.(*RecordComponentListContext)
	if !ok {
		// Same behavior as Java: unexpected state -> true
		return true
	}

	rcs := tctx.AllRecordComponent()
	if len(rcs) == 0 {
		return true
	}

	for i, rcAny := range rcs {
		rc, ok := rcAny.(*RecordComponentContext)
		if !ok {
			continue
		}
		// In generated Go parser, token accessor usually exists as rc.ELLIPSIS()
		// and returns antlr.TerminalNode (or nil).
		if rc.ELLIPSIS() != nil && i+1 < len(rcs) {
			return false
		}
	}
	return true
}

// IsNotIdentifierAssign replicates JavaParserBase.IsNotIdentifierAssign()
func (p *JavaParserBase) IsNotIdentifierAssign() bool {
	la := p.GetTokenStream().LA(1)

	// If not identifier-ish, return true because it can't be "identifier = ..."
	switch la {
	case JavaParserIDENTIFIER,
		JavaParserMODULE,
		JavaParserOPEN,
		JavaParserREQUIRES,
		JavaParserEXPORTS,
		JavaParserOPENS,
		JavaParserTO,
		JavaParserUSES,
		JavaParserPROVIDES,
		JavaParserWHEN,
		JavaParserWITH,
		JavaParserTRANSITIVE,
		JavaParserYIELD,
		JavaParserSEALED,
		JavaParserPERMITS,
		JavaParserRECORD,
		JavaParserVAR:
		// continue
	default:
		return true
	}

	la2 := p.GetTokenStream().LA(2)
	if la2 != JavaParserASSIGN {
		return true
	}
	return false
}
