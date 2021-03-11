package step

import (
	"fmt"
	"reflect"
	"strings"
)

/* this is the only function that is not part of the STEP ISO template */
func (p *Parser) skipUtility(key string) {
	fmt.Println("doesn't loop like i've implemented `", p.lexemes[key].Type, "` skipping to next States..")
}

func (p *Parser) productDefinitionShape(key string) {

	productDefinitionState := parseState{fn: p.productDefinition, key: p.lexemes[key].Arguments[2]}
	p.nextStates = append(p.nextStates, productDefinitionState)

}
func (p *Parser) productDefinition(key string) {

	productDefinitionFormationState := parseState{fn: p.productDefinitionFormation, key: p.lexemes[key].Arguments[2]}
	productDefinitionContextState := parseState{fn: p.productDefinitionContext, key: p.lexemes[key].Arguments[2]}

	p.nextStates = append(p.nextStates, productDefinitionFormationState, productDefinitionContextState)

}

func (p *Parser) productDefinitionFormation(key string) {

}

func (p *Parser) productDefinitionContext(key string) {

}

func (p *Parser) axis2Placement3D(key string) {

	cartesianVectorState := parseState{fn: p.cartesianPoint, key: p.lexemes[key].Arguments[1]}
	xDirectionState := parseState{fn: p.direction, key: p.lexemes[key].Arguments[1]}
	yDirectionState := parseState{fn: p.direction, key: p.lexemes[key].Arguments[1]}
	p.nextStates = append(p.nextStates, cartesianVectorState, xDirectionState, yDirectionState)

}

func (p *Parser) cartesianPoint(key string) {
	//fmt.Println("key:", key)
	//fmt.Println("value", p.lexemes[key])

	arguments := p.lexemes[key].Arguments
	coordinates := getArray(arguments[1], reflect.Float64)
	_ = coordinates
	//fmt.Println("coordinates:", coordinates)

}

func (p *Parser) direction(key string) {

}

func (p *Parser) manifoldSolidBRep(key string) {

	parseStates := make([]parseState, 0)

	closedShellState := parseState{fn: p.closedShell, key: p.lexemes[key].Arguments[1]}
	parseStates = append(parseStates, closedShellState)
	// vectorLoopState := parseState{fn: parseVectorLoop, key: l[key].Arguments[1]}
	// if l[key].Type == "MANIFOLD_SOLID_BREP" {
	// 	parseStates = append(parseStates, vectorLoopState)
	//
	// }
	//

}

/* TODO : add function for adding to lexemes */
func (p *Parser) closedShell(key string) {

}

func (p *Parser) vectorLoop(key string) {
	parseFunctions := make([]parseState, 0)

	parseFunctions = append(parseFunctions)

}

func (p *Parser) advancedFace(key string) {

	fmt.Println("parsing advanced face")

	advancedFace := p.lexemes[key]
	neighbourID := advancedFace.Arguments[1]
	fmt.Println(p.lexemes[neighbourID].Type)

	switch faceType := strings.TrimSpace(p.lexemes[neighbourID].Type); {
	case faceType == "SPHERICAL_SURFACE":
		fmt.Println("it's a sphere")
		sphericalSurfaceState := parseState{fn: p.sphericalSurface, key: neighbourID}
		p.nextStates = append(p.nextStates, sphericalSurfaceState)
		fmt.Println("STATES", p.nextStates)

	case faceType == "TOROIDAL_SURFACE":
		ToroidalSurfaceState := parseState{fn: p.toroidalSurface, key: neighbourID}
		p.nextStates = append(p.nextStates, ToroidalSurfaceState)

	default:

	}

}

func (p *Parser) sphericalSurface(key string) {
	// add rendering data to the channel
	// and propagate to new p.nextStates of the parser.
	//p.nextStates := make([]parseState, 0)
	fmt.Println("Bazinga!!!!!!")
	radius := p.lexemes[key].Arguments[2]
	fmt.Println("radius:", radius)

}

func (p *Parser) toroidalSurface(key string) {
	// add rendering data to the channel
	// and propagate to new p.nextStates of the parser.
	fmt.Println("Bazinga!!!!!!")

}

func (p *Parser) vertexPoint(key string) {

}

func (p *Parser) vertexLoop(key string) {
}

func (p *Parser) face(key string) {
	parseFunctions := make([]parseState, 0)
	parseFunctions = append(parseFunctions)

}

func (p *Parser) bSplineSurfaceWithKnots(key string) {
	/* parameters:
	----Name				string
	----UDegree				int
	----VDegree				int
	----SurfaceForm			Enum of BSpline Surface: https://www.steptools.com/stds/stp_aim/html/t_b_spline_surface_form.html
	----controlPointsList	matrix of CARTESIAN_POINT
	----UClosed				bool
	----VClosed				bool
	----selfIntersect		bool
	----UMultiplicities		List of integers
	----VMultiplicities		List of integers
	----UKnots				List of reals
	----VKnots				List of reals
	----KnotSpec			Enum of KnotType : https://www.steptools.com/stds/stp_aim/html/t_knot_type.html
	*/
	fmt.Println("Hooray! we found an awesome bSplineSurface with KNOTS")
	fmt.Println("len(arguemnts):", len(p.lexemes[key].Arguments))
	name := p.lexemes[key].Arguments[0]
	uDegree := p.lexemes[key].Arguments[1]
	vDegree := p.lexemes[key].Arguments[2]
	controlPointsList := getMatrix(p.lexemes[key].Arguments[3])
	surfaceForm := p.lexemes[key].Arguments[4] /* this is buggy */
	uClosed := getLogical(p.lexemes[key].Arguments[5])
	vClosed := getLogical(p.lexemes[key].Arguments[6])
	selfIntersect := getLogical(p.lexemes[key].Arguments[7])
	uMultiplicities := getArray(p.lexemes[key].Arguments[8], reflect.Int)
	vMultiplicities := getArray(p.lexemes[key].Arguments[9], reflect.Int)
	uKnots := getArray(p.lexemes[key].Arguments[10], reflect.Float64)
	vKnots := getArray(p.lexemes[key].Arguments[11], reflect.Float64)

	surfaceJob := bSplineSurfaceWithKnots{
		name:              name,
		uDegree:           uDegree,
		vDegree:           vDegree,
		controlPointsList: controlPointsList,
		surfaceForm:       surfaceForm,
		uClosed:           uClosed,
		vClosed:           vClosed,
		selfIntersect:     selfIntersect,
		uMultiplicities:   uMultiplicities,
		vMultiplicities:   vMultiplicities,
		uKnots:            uKnots,
		vKnots:            vKnots,
	}
	surfaceJob.print()
}
