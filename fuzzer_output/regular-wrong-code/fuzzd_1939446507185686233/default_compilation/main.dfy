function abs(x: int): int {
	if (x < 0) then -1 * x else x
}
function safeIndex(x: int, length: int): int 
	requires length > 0
{
	if (x < 0) then 0 else if (x >= length) then x % length else x
}
function safeDivisionInt(x1: int, x2: int): int {
	if (x2 == 0) then x1 else x1 / x2
}
function safeModuloInt(x1: int, x2: int): int {
	if (x2 == 0) then x1 else x1 % x2
}
datatype D0 = DC0(cf0: bool)
datatype D1 = DC1(cf1: map<int, bool>)
datatype D2 = DC3(cf3: bool, cf4: bool) | DC4(cf5: int, cf6: char) | DC2(cf2: char)
datatype D3 = DC5(cf7: map<multiset<bool>, C0>)
datatype D4 = DC7(cf9: int, cf10: bool) | DC6(cf8: array<bool>)
datatype D5 = DC9(cf12: string, cf13: bool, cf14: C1, cf15: char) | DC8(cf11: map<int, array<bool>>)
datatype D6 = DC11(cf17: string, cf18: C0, cf19: int, cf20: seq<int>, cf21: int) | DC12(cf22: bool) | DC13 | DC10(cf16: map<C1, bool>) | DC14(cf23: D6)
datatype D7 = DC16(cf25: string, cf26: multiset<array<bool>>, cf27: bool) | DC15(cf24: seq<bool>)
datatype D8 = DC18(cf29: int, cf30: bool) | DC17(cf28: multiset<map<bool, int>>) | DC19(cf31: D8)
class GlobalState {
	const f0 : array<bool>
	const f1 : int
	const f2 : bool
	var f3 : array<bool>
	var f4 : int
	const f5 : int
	const f6 : bool
	const f7 : int
	const f8 : bool
	var f9 : bool
	var f10 : bool
	constructor (f0 : array<bool>, f1 : int, f2 : bool, f3 : array<bool>, f4 : int, f5 : int, f6 : bool, f7 : int, f8 : bool, f9 : bool, f10 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
		this.f7 := f7;
		this.f8 := f8;
		this.f9 := f9;
		this.f10 := f10;
	}
	
}

class C0 {
	const f12 : char
	constructor (f12 : char) {
		this.f12 := f12;
	}
	
	function fm3(p0: set<int>, globalState: GlobalState): D0 {
		if (if (DC0(!true).cf0) then false else true) then DC0(false) else DC0(false)
	}
}

class C1 {
	const f11 : D0
	constructor (f11 : D0) {
		this.f11 := f11;
	}
	
	method m0(p0: int, p1: string, p2: bool, p3: int, globalState: GlobalState) returns (r0: seq<string>) {
		var v0: seq<int> := [|p1|, p3];
		var i0 := 0;
		while (v0 < [--0x21])
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			globalState.f4 := -p3;
			var v1: map<int, string> := map[-0x300 := p1];
			v1 := v1[p0 := p1];
			globalState.f10 := fm0(p2, if (p2) then p2 else true, p3, safeDivisionInt(p3, p3), globalState);
			var v2: array<string> := new string[26](i1 => p1);
			v2[safeIndex(105, v2.Length)] := p1;
			v2[safeIndex(105, v2.Length)] := "tqgbnrs";
		}
		globalState.f10 := p2;
		for i2 := p3 to p3 {
			var v3 := 'p';
			var v4: map<int, bool> := map[|map[v3 := p2]| := p2];
			var v5: seq<string> := [p1];
			var v6: array<int> := new int[6] [p3, i2, i2, fm2(p2, |multiset{p2, p2, p2}|, |v4|, v5[safeIndex(|p1|, |v5|)], globalState), p3, p0 - p3];
			v6[safeIndex(323, v6.Length)] := i2;
			var v7: seq<bool> := [p2];
			v6[safeIndex(323, v6.Length)] := -(|v7| - i2);
			if (p2) {
				var v8: multiset<int> := multiset{866, i2};
				globalState.f4 := v6[safeIndex(323, v6.Length)] * -(-(if (|p1[safeIndex(p0, |p1|) := v3]| in v8) then v8[|p1[safeIndex(p0, |p1|) := v3]|] else i2) * -p3);
				v6[safeIndex(323, v6.Length)] := p3;
				globalState.f9 := p2 ==> fm0(p2, false, v0[safeIndex(v6[safeIndex(323, v6.Length)], |v0|)], -|v7|, globalState);
				v6[safeIndex(323, v6.Length)] := (if (|p1| in v8) then v8[|p1|] else -p3) * v0[safeIndex(p3, |v0|)];
				var v9: map<bool, seq<int>> := map[p2 := v0];
				globalState.f10 := (if (p2 in v9) then v9[p2] else v0) == v0;
			} else {
				var v10 := new C0(v3);
				var v11: array<map<int, bool>> := new map<int, bool>[6];
				v11[safeIndex(209, v11.Length)] := map[i2 := p2];
				var v12 := DC1(map[i2 := p2]);
				var v13: seq<map<int, bool>> := [v4, v12.cf1 + map[p3 := p2], v4 + v4, v4, v4];
				v11[safeIndex(209, v11.Length)] := v13[safeIndex(v6[safeIndex(323, v6.Length)] * p3, |v13|)];
				var v14 := new C0(v10.f12);
				var v15: array<string> := new string[13](i3 => p1);
				v15[safeIndex(933, v15.Length)] := seq(abs(-0x2c2), i4  => (DC4(p3, v3).cf6));
				var v16: set<int> := {p3};
				var v17: seq<array<int>> := [v6];
				var v18: map<array<int>, int> := map[v17[safeIndex(p0, |v17|)] := |v0|];
				var v19: multiset<D0> := multiset{DC0(p2)};
				var v20: map<int, int> := map[615 := i2];
				v15[safeIndex(933, v15.Length)], globalState.f9, globalState.f4, globalState.f9 := p1, v16 >= v16, if (v18 != v18) then if (DC0(v7[safeIndex(|v0|, |v7|)]) in v19) then v19[DC0(v7[safeIndex(|v0|, |v7|)])] else i2 else if (p0 in v20) then v20[p0] else p3, !p2;
				var v21: multiset<C0> := multiset{v10, v10};
				v15[safeIndex(933, v15.Length)] := fm4(-|v21|, p2, !p2, v6[safeIndex(323, v6.Length)] * v6[safeIndex(323, v6.Length)], globalState);
			}
			
			var v22 := DC1(v4);
			var v23: seq<D1> := [v22, v22];
			globalState.f10 := (if (p2) then v23 else v23) <= v23;
			globalState.f4 := safeModuloInt(safeDivisionInt(fm2(p2, v6[safeIndex(323, v6.Length)], |(seq(abs(0xc7), i5  => (v3)))|, p1, globalState), i2), v6[safeIndex(323, v6.Length)]);
		}
		if (p2) {
			var v24 := new C0(fm5(-|fm4(p3, p2, p2, |"aroliv"|, globalState)|, globalState));
			var v25: seq<seq<int>> := [v0];
			v25 := v25;
			globalState.f10 := !(false <==> (0x27c > p0));
			var v26: array<int> := new int[26];
			v26[safeIndex(933, v26.Length)] := p0;
			var v27: set<bool> := {p2, p2, p2, !p2};
			var v28: array<bool> := new bool[15];
			var v29: seq<bool> := [p2];
			var v30 := DC3(p2, p2);
			v26[safeIndex(933, v26.Length)], globalState.f9, globalState.f4, globalState.f4, globalState.f9 := p3, v27 >= v27, -safeDivisionInt(p0, p0), safeDivisionInt(p0, p0), |fm4(|multiset{v28}|, p2, v29[safeIndex(p0, |v29|)], p0, globalState)| > (if (v30.cf4) then p3 else p0);
			var v31: multiset<bool> := multiset{p2, p2, p2, p2, !p2};
			var v32: map<multiset<bool>, C0> := map[v31 := v24];
			var v33 := DC5(v32);
			var v34: map<bool, map<multiset<bool>, C0>> := map[p2 := v33.cf7];
			var v35: array<map<multiset<bool>, C0>> := new map<multiset<bool>, C0>[8] [if (p2 in v34) then v34[p2] else v32, map[fm6(true, v26[safeIndex(933, v26.Length)], globalState) := v24], map[v31 := v24], v32, map[v31 := v24], v32, v32, v32];
			v35[safeIndex(352, v35.Length)] := v32;
			v35[safeIndex(352, v35.Length)] := v33.cf7;
		} else {
			globalState.f10 := p2;
			var v36 := DC3(p2, p2);
			match (v36) {
				case DC3(cf3, cf4) =>
					var v37 := 'b';
					var v38: array<bool> := new bool[22](i6 => p2);
					var v39: map<char, array<bool>> := map[v37 := v38];
					v39 := map[if (p2) then fm5(|v0|, globalState) else v37 := v38];
					var v40 := "ovmhyj";
					v38[safeIndex(415, v38.Length)] := cf3;
					var v41: multiset<bool> := multiset{p2};
					globalState.f10, v40, v38[safeIndex(415, v38.Length)] := cf4, p1 + (v40 + p1), (false in v41) <==> p2;
					var v42: array<int> := new int[14] [fm2(!p2, -|{p0, p0}|, p0, seq(abs(-0x14b), i7  => (v37)), globalState), -31, p3, p3, p3, p3, p3, p0, p3, p0, p0, -p0, p3, p3];
					var v43: map<bool, array<int>> := map[p2 := v42];
					var v44: seq<bool> := [cf3];
					var v45: map<char, int> := map[v37 := |"w"|];
					var v46: multiset<int> := multiset{43};
					var v47: seq<bool> := [false, fm0(fm0(cf3, v44[safeIndex(p0, |v44|)], 0x328, p0, globalState), p2, if ('r' in v45) then v45['r'] else 0x226, |v46|, globalState), cf4, cf3, p2];
					v43 := v43[v44[safeIndex(-0xd9, |v44|)] := v42];
					globalState.f9 := cf3;
				case DC4(cf5, cf6) =>
					var v48 := DC4(cf5, cf6);
					var v49: map<D2, bool> := map[v48.(cf6 := 'd') := true];
					v49 := v49[v48 := p1 < p1];
					var v50: C0 := new C0(cf6);
					var v51: multiset<C0> := multiset{v50, v50, v50, v50};
					v51 := v51[v50 := abs(cf5)];
					cf6 := v50.f12;
					var v52: set<bool> := {!p2, p2, p2, fm0(true, p2, 0x249, cf5, globalState), true};
					var v53: multiset<int> := multiset{p3};
					var v54: map<C0, int> := map[v50 := p3];
					var v55: map<multiset<bool>, C0> := map[multiset{p2} := v50];
					var v56 := DC5(v55);
					var v57: map<bool, D3> := map[v36.cf3 := v56];
					var v58: seq<bool> := [p2];
					var v59: array<int> := new int[17] [|v0|, p0, p3, p0, |v52|, 0x260, |v53|, |p1|, if (cf5 in v53) then v53[cf5] else -fm2(true, p0, p0, p1, globalState), cf5, |v54|, |v57[p2 := v56]|, |v58|, p0, p3, cf5, fm2(p2, p0, cf5, p1, globalState)];
					var v60: array<int> := new int[7] [|p1|, cf5, |map[{v50} := v59]|, p3, 0x29f, cf5, p3];
					v60[safeIndex(655, v60.Length)] := -0x26 + p3;
					v60[safeIndex(655, v60.Length)] := p0;
				case DC2(cf2) =>
					var v61: set<int> := {p3};
					var v62: map<D0, int> := map[f11 := p3];
					var v63: map<bool, int> := map[p2 := p0];
					var v64: C0 := new C0(cf2);
					var v65: map<C0, bool> := map[v64 := p2];
					var v66: multiset<int> := multiset{p0, p3};
					var v67: multiset<bool> := multiset{p2, p2};
					var v68: array<int> := new int[21] [fm2(p2, p3, |p1|, p1, globalState), --fm2(p2, p3, p0, p1, globalState), if (p2) then -p0 else p3, p0, if (!p2) then |v61| else p0, safeDivisionInt(807, p0), safeModuloInt(|p1|, p3), p3, p3, |multiset(seq(abs(0x2ee), i8  => (p0)))|, p3, p3, if (p2) then p0 else if (f11 in v62) then v62[f11] else p3, p0, |p1|, -0x275, -safeModuloInt(-0x6f, |v63|), |v65|, -|v66|, p3, if (p2) then p3 else |[|multiset{if (p2 in v67) then v67[p2] else p0}|]|];
					var v69: array<char> := new char[21](i9 => v64.f12);
					v69[safeIndex(650, v69.Length)] := cf2;
					v68, v69[safeIndex(650, v69.Length)], globalState.f10, globalState.f9 := v68, cf2, p2, p2;
					v64 := new C0(v64.f12);
					var v70: map<int, bool> := map[p3 := fm0(p2, p2, p3, p3, globalState)];
					v70 := map v71 : int | (531 <= v71) && (v71 < 0x111) :: (v71 - |v0|) := (!([p2] == [p2, p2]));
					var v72 := new C0(v64.f12);
			}
			
			var v73: map<bool, string> := map[p2 := "nqmai"];
			var v74: map<bool, bool> := map[p2 := p2];
			var v75: set<int> := {|v74|};
			var v76 := 'o';
			var v77: multiset<bool> := multiset{p2};
			var v78: seq<string> := ["twnjp", p1];
			var v79: array<int> := new int[28] [|{p3}|, p3, -safeModuloInt(p3, p3), p0, if (p2) then p0 else -p3, |fm7(p2, p0, p2, |v73|, globalState)| - -p3, p0, p3, p3, p3, |v75|, p0, safeDivisionInt(p0, |v74|), -p0, |("m")[safeIndex(p0, |"m"|) := v76]|, v0[safeIndex(250, |v0|)], p3, |fm8(p2, p1, globalState)|, -p3, p3, -652 - p0, |"xlpfc"|, 739, |((multiset{p2})[p2 := abs(p3)] + v77)|, p3, p3, 655, |(p1 + v78[safeIndex(p0, |v78|)])|];
			var v80: multiset<int> := multiset{|p1|, p3};
			var v81: C0 := new C0('j');
			var v82: map<multiset<bool>, C0> := map[v77 := v81];
			var v83 := DC5(v82);
			var v84: set<D3> := {v83};
			v79[safeIndex(566, v79.Length)] := if (p0 in v80) then v80[p0] else -|v84|;
			var v85: map<int, int> := map[p0 := p0];
			var v86: map<int, int> := map[p3 := if (p0 in v85) then v85[p0] else p3];
			var v87: map<int, set<D2>> := map[p3 := {DC3(p2, p2), v36, v36, v36}];
			v79[safeIndex(566, v79.Length)], v76, globalState.f10 := |fm1(v85, !p2, |p1|, -p0, globalState)|, p1[safeIndex(414, |p1|)], fm9(p2, globalState) != v87;
			globalState.f9 := p2;
			var v88: map<bool, int> := map[p2 := p0];
			v88 := v88[p2 := -0x386 + |v88|];
		}
		
		var v89: multiset<bool> := multiset{p2};
		var v90 := 'y';
		var v91: C0 := new C0(v90);
		var v92: map<multiset<bool>, C0> := map[v89 := v91];
		var v93 := DC5(v92);
		var v94: seq<D3> := [v93];
		for i10 := safeDivisionInt(|multiset(v94)|, p0) to fm2(p2, --285, p0, p1, globalState) {
			var v95: map<int, bool> := map[p3 := p2];
			var v96: map<int, int> := map[p0 := |v95|];
			var v97: map<int, bool> := map[if (p0 in v96) then v96[p0] else p3 := p2];
			var v98: seq<bool> := [p2 <==> !(if (p3 in v97) then v97[p3] else p2)];
			globalState.f10 := v98[safeIndex(i10, |v98|)];
			var v99: array<D3> := new D3[21] [v93, if (p2) then v93 else v93, v93, v93, DC5(map[v89 := v91]), v93, v93, v93, v93, v93, v93, v93, v93, v93.(cf7 := v92), v93.(cf7 := v92), v93, v93, DC5(v92), DC5(v92), DC5(v92), DC5((v93.(cf7 := v92)).cf7)];
			v99[safeIndex(950, v99.Length)] := v93;
			v99[safeIndex(950, v99.Length)] := v93.(cf7 := v92).(cf7 := (map[v89 := v91])[v89 := v91]).(cf7 := v92);
			var v100: map<bool, int> := map[p2 := i10];
			var v101: array<bool> := new bool[11];
			var v102: map<map<bool, int>, array<bool>> := map[v100 := v101];
			v102 := v102[v100 := v101];
			var v103 := new C0(v90);
		}
		var v104: seq<bool> := [p2];
		var v105: array<D0> := new D0[26] [DC0(p2), DC0(p2), f11, f11, f11, f11, f11, f11, f11, f11.(cf0 := true), f11, f11, f11, DC0(p2), f11, f11, f11, f11, f11, f11, f11.(cf0 := v104[safeIndex(p3, |v104|)]), f11, f11, f11, DC0(p2), f11];
		v105[safeIndex(463, v105.Length)] := f11;
		v105[safeIndex(463, v105.Length)] := f11;
		r0 := fm10(globalState);
	}
}

function fm0(p0: bool, p1: bool, p2: int, p3: int, globalState: GlobalState): bool {
	0x360 != (|{572}| * |(seq(abs(-0x1e3), i0  => ('s')))|)
}
function fm1(p0: map<int, int>, p1: bool, p2: int, p3: int, globalState: GlobalState): set<int> {
	((set v0 : int | v0 in map[0x24b := false] :: (v0 - |multiset{!true}|)) + (set v1 : int | (0xb6 <= v1) && (v1 < 0x36d) :: (v1 + |{false}|))) * {-0xd8}
}
function fm2(p0: bool, p1: int, p2: int, p3: string, globalState: GlobalState): int {
	safeDivisionInt(|({false} * {false})|, 0x51)
}
function fm4(p0: int, p1: bool, p2: bool, p3: int, globalState: GlobalState): string {
	"nargnu"
}
function fm5(p0: int, globalState: GlobalState): char {
	match DC2('a') {
		case DC3(cf3, cf4) => 'x'
		case DC4(cf5, cf6) => DC2('l').cf2
		case DC2(cf2) => cf2
	}
}
function fm6(p0: bool, p1: int, globalState: GlobalState): multiset<bool> {
	(multiset{!false, true, !true} - multiset{false, false}) - (multiset([true, true, true]) - multiset{false, false, !!!true})
}
function fm7(p0: bool, p1: int, p2: bool, p3: int, globalState: GlobalState): map<int, int> {
	(map[-0x284 := 0x140] + map[|multiset([true, false])| := -|[-0x17c]|]) + (map[-569 := -889] + map[0x1b0 := 0xd2])
}
function fm8(p0: bool, p1: string, globalState: GlobalState): seq<int> {
	((seq(abs(0x259), i0  => (0x213))) + [238]) + (seq(abs(0x30b), i1  => (0xd1)))
}
function fm9(p0: bool, globalState: GlobalState): map<int, set<D2>> {
	(map[|multiset{|{827}|, |map[true := true]|}| := {DC3(false, true)}] + map[-0x2f := set v0 : D2 | v0 in map[DC3(!!false, true) := {false}] :: (v0)]) + (map[-0x1e7 := {DC3(false, true)}] + map[|map[-0x28e := !true]| := {DC3(true, true)}])
}
function fm10(globalState: GlobalState): seq<string> {
	seq(abs(0xdc), i0  => (seq(abs(0x1ed), i1  => ('n'))))
}
function fm11(p0: int, globalState: GlobalState): seq<bool> {
	["vwsaa" < "dmyeoj", map[|map[false := false]| := false] != map[|[-|(set v0 : int | (808 <= v0) && (v0 < -796) :: (v0 + 0xa))|, 754, 462, |"tkpurq"|]| := !false], [0x382, |{|"agfup"|}|] != [|map[true := -0x2f0]|, |[false]|, |multiset{|[multiset{|{!true}|, DC7(0x66, true).cf9}]|}|, 0x56, |map[false := 741]|]]
}
function fm12(p0: int, globalState: GlobalState): D2 {
	if (!(DC17(multiset{map[false := 977], map[false := 0x65]}).cf28 !! multiset{map[true := |(map v0 : int | (0x359 <= v0) && (v0 < 0x163) :: (safeDivisionInt(v0, 0x376)) := ([true]))|]})) then DC3(!true, false) else DC3(!true, false)
}
function fm13(p0: int, p1: int, p2: string, p3: bool, globalState: GlobalState): set<string> {
	if (true) then {"b", seq(abs(-0x227), i0  => ('q')), "w"} + {"l"} else {"ll"}
}
function fm14(p0: bool, p1: int, p2: set<int>, globalState: GlobalState): map<multiset<bool>, bool> {
	(map[multiset{true} := !!false] + (map v0 : multiset<bool> | v0 in [multiset{!true}] :: (v0) := (false))) + ((map v1 : multiset<bool> | v1 in multiset{multiset([false]), multiset{true}, multiset{true, false}, multiset{false}} :: (v1) := (false)) + map[multiset{DC18(-|{-0x3e6}|, true).cf30} := !true])
}
function fm15(globalState: GlobalState): map<bool, bool> {
	(map[true := true] + map[true := false]) + (map[true := false] + map[false := true])
}
function fm16(globalState: GlobalState): map<set<bool>, int> {
	match DC13() {
		case DC11(cf17, cf18, cf19, cf20, cf21) => map[{true, true} := cf21]
		case DC12(cf22) => map[{cf22} := -0x1ef]
		case DC13() => map[{!false, true, true, true, false} := -|"hw"|]
		case DC10(cf16) => map[{true} := 414]
		case DC14(cf23) => map[{true} := -560]
	}
}
function fm17(p0: map<int, map<int, bool>>, p1: string, p2: int, globalState: GlobalState): set<bool> {
	{true}
}
method m1(p0: int, p1: array<bool>, globalState: GlobalState) returns (r0: array<bool>) {
	var v0 := false;
	p1[safeIndex(666, p1.Length)] := v0;
	var v1 := "xdfamt";
	var v2: multiset<int> := multiset{p0, -p0};
	var v3: seq<int> := [p0, p0, p0, |v1|, if (p0 in v2) then v2[p0] else p0];
	p1[safeIndex(666, p1.Length)] := v0 <== ([p0, p0, |"benlgnvi"|, p0] < v3);
	var v4 := DC3(p1[safeIndex(666, p1.Length)], p1[safeIndex(666, p1.Length)]);
	var v5: array<D2> := new D2[3] [DC3(p1[safeIndex(666, p1.Length)], p1[safeIndex(666, p1.Length)]), v4, v4];
	var v6: seq<array<D2>> := [v5];
	v6 := v6;
	var v7: seq<bool> := [p1[safeIndex(666, p1.Length)]];
	if ((p1[safeIndex(666, p1.Length)] !in v7) && !(fm0(v0, true, 0x1bc, p0, globalState) ==> !p1[safeIndex(666, p1.Length)])) {
		var v8 := DC0(p1[safeIndex(666, p1.Length)]);
		var v9: C1 := new C1(v8);
		var v10 := 'i';
		var v11 := DC9(v1, v0, v9, v10);
		match ((if (!v0) then DC9("j", v0, v9, v10) else v11).(cf14 := v9)) {
			case DC9(cf12, cf13, cf14, cf15) =>
				var v12: set<bool> := {cf13};
				var v13: map<set<bool>, int> := map[v12 := |multiset{p1[safeIndex(666, p1.Length)]}|];
				var v14: C0 := new C0(cf12[safeIndex(p0, |cf12|)]);
				var v15: map<bool, C0> := map[fm16(globalState) != v13 := v14];
				v0, globalState.f9, v15 := |(if (p1[safeIndex(666, p1.Length)]) then v7 else v7)| !in (set v16 : int | (0x126 <= v16) && (v16 < 0x1e5) :: (v16 - |[v1, cf12]|)), v0 <==> fm0(v0, v0, p0, p0, globalState), v15;
				var v17 := new C0(v10);
				var v18: map<int, array<bool>> := map[|v2| := p1];
				var v19 := DC8(v18 + v18);
				v19 := if (v0) then v19 else v19;
				var v20: map<int, char> := map[-fm2(v0, p0, p0, cf12, globalState) := v10];
				var v21: map<bool, int> := map[v0 := p0];
				var v22: map<int, map<bool, int>> := map[p0 := v21];
				var v23: array<char> := new char[28] ['k', cf15, cf15, cf15, if (0x11d in v20) then v20[0x11d] else v10, cf15, 'e', v14.f12, v17.f12, v10, v17.f12, v10, v17.f12, cf15, v1[safeIndex(-0x239, |v1|)], v10, v14.f12, v14.f12, fm5(|v22|, globalState), v10, v10, v17.f12, v10, cf15, v10, v14.f12, 'f', v14.f12];
				var v24: map<bool, array<char>> := map[cf13 := v23];
				var v25: seq<array<char>> := [if (p1[safeIndex(666, p1.Length)] in v24) then v24[p1[safeIndex(666, p1.Length)]] else v23];
				v0 := true ==> (v23 !in v25[safeIndex(p0, |v25|) := v23]);
			case DC8(cf11) =>
				p1[safeIndex(666, p1.Length)] := p1[safeIndex(666, p1.Length)];
				globalState.f4 := p0 * p0;
				var v26: array<char> := new char[20];
				v26[safeIndex(104, v26.Length)] := v10;
				globalState.f4, v1, v26[safeIndex(104, v26.Length)] := -951, v1, v10;
				var v27: map<int, char> := map[p0 := v26[safeIndex(104, v26.Length)]];
				var v28: map<int, map<int, char>> := map[p0 := v27];
				globalState.f4 := safeModuloInt(safeDivisionInt(p0, p0), if (p1[safeIndex(666, p1.Length)]) then p0 else --|(if (p0 in v28) then v28[p0] else map[p0 := v10])|);
		}
		
		var v29: multiset<string> := multiset{"hqp", seq(abs(0xc), i0  => (v10))};
		var v30: multiset<bool> := multiset{v0};
		var v31: map<int, multiset<bool>> := map[p0 := v30];
		var v32: map<bool, bool> := map[p1[safeIndex(666, p1.Length)] := v0];
		var v33: map<int, bool> := map[p0 := p1[safeIndex(666, p1.Length)]];
		var v34: map<int, int> := map[|v33| := p0];
		var v35: array<int> := new int[21] [p0, p0 * p0, p0, if (v1 in v29) then v29[v1] else |v1|, |(seq(abs(0xff), i1  => (v1)))|, p0, p0, -p0, p0 - p0, p0, |v31|, 0x3a8, p0, -|v32[v0 := true]|, p0, p0, p0, p0, p0 * (if (v0 in v30) then v30[v0] else |v34|), 405, if (v0) then p0 else p0];
		v35[safeIndex(878, v35.Length)] := p0;
		v35[safeIndex(878, v35.Length)] := if (DC7(if (v0 in v30) then v30[v0] else p0, p1[safeIndex(666, p1.Length)]).cf9 in v34) then v34[DC7(if (v0 in v30) then v30[v0] else p0, p1[safeIndex(666, p1.Length)]).cf9] else p0;
		var v36 := new C0(if (fm0(p1[safeIndex(666, p1.Length)], true, -p0, v35[safeIndex(878, v35.Length)], globalState)) then fm5(|v1[safeIndex(768, |v1|) := v10]|, globalState) else v10);
		globalState.f10 := !p1[safeIndex(666, p1.Length)];
		var v37: set<string> := {v1, v1};
		v35 := new int[12] [v35[safeIndex(878, v35.Length)], v35[safeIndex(878, v35.Length)], p0, p0, p0, v35[safeIndex(878, v35.Length)], v35[safeIndex(878, v35.Length)], |v37|, v35[safeIndex(878, v35.Length)], safeModuloInt(p0, p0), |(v7 + v7)|, v35[safeIndex(878, v35.Length)]];
	} else {
		var v38 := DC0(!v0);
		var v39: C1 := new C1(v38);
		var v40: map<C1, bool> := map[v39 := v0];
		var v41 := DC10(v40);
		var v42: map<multiset<int>, map<C1, bool>> := map[v2 := v41.cf16];
		var v43: seq<map<C1, bool>> := [v40];
		var v44: map<int, map<C1, bool>> := map[-0x3dd := v40];
		var v45: map<bool, C1> := map[p1[safeIndex(666, p1.Length)] := v39];
		var v46: seq<C1> := [v39];
		var v47: array<map<C1, bool>> := new map<C1, bool>[25] [if (v2 in v42) then v42[v2] else v40, v40, v40, v40, v40, v40 + v40, map[v39 := p1[safeIndex(666, p1.Length)]], v40, v43[safeIndex(p0, |v43|)], v40, v40 + v40, v40, v40, v40[v39 := true] + map[v39 := v0], v40, map[v39 := fm0(v0, fm0(true, v0, p0, p0, globalState), p0, p0, globalState)], v40, (if (|v3| in v44) then v44[|v3|] else v40) + v40, v40, v40, v40, map[if (v0 in v45) then v45[v0] else v39 := v0], v40, map[v46[safeIndex(p0, |v46|)] := p1[safeIndex(666, p1.Length)]], v40 + map[v39 := v0]];
		var v48: map<int, int> := map[p0 := -p0];
		var v49: array<int> := new int[3] [p0, if (p0 in v48) then v48[p0] else p0, p0];
		v49[safeIndex(242, v49.Length)] := fm2(v0, p0, p0, v1, globalState);
		var v50: array<set<bool>> := new set<bool>[7];
		var v51: set<bool> := {v0};
		v50[safeIndex(267, v50.Length)] := v51;
		var v52: map<multiset<int>, array<map<C1, bool>>> := map[v2 := v47];
		var v53 := 'h';
		var v54: C0 := new C0(v53);
		var v55: map<int, map<int, bool>> := map[p0 := map[p0 := p1[safeIndex(666, p1.Length)]]];
		v47, v49[safeIndex(242, v49.Length)], globalState.f4, globalState.f9, v50[safeIndex(267, v50.Length)] := if (multiset{|DC11(v1, v54, |[v54]|, [p0, p0], p0).cf20|, p0, p0, p0, -582} in v52) then v52[multiset{|DC11(v1, v54, |[v54]|, [p0, p0], p0).cf20|, p0, p0, p0, -582}] else v47, p0, -p0, v0, fm17(v55, v1, p0, globalState);
		globalState.f9 := v0;
		globalState.f4 := p0 + p0;
		var v56: array<seq<bool>> := new seq<bool>[16](i2 => v7);
		v56[safeIndex(633, v56.Length)] := v7 + v7;
		var v57 := DC15(v7);
		v56[safeIndex(633, v56.Length)] := v57.cf24;
		var v58: array<array<bool>> := new array<bool>[17] [p1, if (v0) then p1 else p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1, p1];
		v58[safeIndex(85, v58.Length)] := p1;
		v58[safeIndex(85, v58.Length)] := p1;
	}
	
	var v59: array<seq<int>> := new seq<int>[10](i3 => v3);
	v59[safeIndex(21, v59.Length)] := v3;
	var v60: array<C1> := new C1[18];
	var v61 := DC0(false);
	var v62: C1 := new C1(v61);
	v60[safeIndex(430, v60.Length)] := v62;
	var v63: set<int> := {p0, |v2|};
	var v64: multiset<set<int>> := multiset{v63};
	v59[safeIndex(21, v59.Length)], v60[safeIndex(430, v60.Length)], v62, v64 := v3 + (v3 + v3), v62, v62, v64;
	var v65: set<bool> := {p1[safeIndex(666, p1.Length)]};
	if (v65 >= v65) {
		var v66: map<D0, bool> := map[v61 := !p1[safeIndex(666, p1.Length)]];
		v66 := v66[v62.f11 := v0];
		v1 := if (v0) then if (true) then v1 else v1 else v1 + v1;
		var v67: map<seq<bool>, bool> := map[v7 + v7 := v0];
		v67 := v67[[!p1[safeIndex(666, p1.Length)]] + v7 := p1[safeIndex(666, p1.Length)]];
		globalState.f3 := p1;
		globalState.f4 := 0x1f;
	} else {
		var v68 := new C1(v62.f11);
		var v69 := DC1(map[p0 := v0]);
		var v70: set<D1> := {v69, v69};
		v70 := v70;
		var v71 := DC13();
		var v72: seq<D6> := [v71, v71, v71];
		var v73 := DC14(v72[safeIndex(|v1|, |v72|)]);
		v73 := v73;
		v1 := v1 + (seq(abs(0x303), i4  => ('y')));
		globalState.f9 := p1[safeIndex(666, p1.Length)];
	}
	
	var v74 := v62.m0(|v1|, v1, v0, p0, globalState);
	r0 := p1;
}
method Main() {
	var v0: array<bool> := new bool[14];
	var globalState := new GlobalState(v0, 0x13f, false, v0, 0xfa, 888, false, 746, false, false, true);
	var v1 := true;
	var v2: multiset<bool> := multiset{v1};
	var v3: map<multiset<bool>, bool> := map[v2 + v2 := v1];
	v3 := v3 + v3;
	var v4 := 0x180;
	for i0 := v4 to -v4 - v4 {
		var v5: seq<bool> := [v1];
		v1 := if (fm0(v1, v1, |"bfjvu"|, |v5|, globalState)) then v1 else fm0(v1, !true, v4, i0, globalState);
		var v6: seq<array<bool>> := [v0];
		v6 := [v0, v0];
		var v7: array<map<int, char>> := new map<int, char>[23](i1 => map[i0 := 'g']);
		var v8 := 'o';
		var v9: map<int, char> := map[v4 := v8];
		v7[safeIndex(840, v7.Length)] := v9;
		var v10 := DC0(false);
		v7[safeIndex(840, v7.Length)] := map[i0 - |fm1(map[i0 := v4], v10.cf0, i0, v4, globalState)| := v8];
		var v11 := "ewk";
		v11 := v11;
	}
	var v12: map<int, bool> := map[v4 - 0x105 := !fm0(v1, v1, v4, 789, globalState)];
	v12 := v12;
	var v13: map<bool, string> := map[v1 := "du"];
	var v14: seq<int> := [v4, safeModuloInt(-|v13|, |(seq(abs(0x29f), i2  => ('u')))|)];
	var v15 := DC0(v1);
	var v16: map<D0, int> := map[v15.(cf0 := v1) := v4];
	var v17 := "ajftdewh";
	globalState.f4 := v14[safeIndex(fm2(v1, if (v15 in v16) then v16[v15] else v4, v4, v17, globalState), |v14|)];
	var v18: map<bool, bool> := map[v1 := if (v1) then false else v1];
	v18 := v18[v1 := v1];
	match (v15) {
		case DC0(cf0) =>
			var v19 := 't';
			var v20 := new C0(v19);
			var v21: seq<string> := [v17];
			var v22: array<string> := new string[18] [v17 + v17, "kwmdrfli", v17, v21[safeIndex(v4, |v21|)], seq(abs(0x65), i3  => (v19)), v17, if (fm0(v1, v1, v4, v4, globalState) in v13) then v13[fm0(v1, v1, v4, v4, globalState)] else v17, (seq(abs(195), i4  => (v19))) + v17, "krisnrwt", v17 + (seq(abs(0x1f0), i5  => (v20.f12))), v17, seq(abs(0x289), i6  => (v20.f12)), v17 + v17, v17, v17, "dkm", fm4(-v4, cf0, cf0, -v4, globalState), (v17 + "qfkthhxah")[safeIndex(v4, |(v17 + "qfkthhxah")|) := v19]];
			v22[safeIndex(280, v22.Length)] := v17;
			v22[safeIndex(280, v22.Length)] := v17;
			var v23: array<char> := new char[10];
			v23[safeIndex(16, v23.Length)] := v20.f12;
			v23[safeIndex(16, v23.Length)] := v19;
			globalState.f10 := (!v1 ==> v1) && v1;
	}
	
	globalState.f10 := v1;
	var v24: array<int> := new int[29](i8 => i8 - |v17|);
	forall i7 | 0 <= i7 < v24.Length {
		v24[i7] := i7 * v4;
	}
	if (v1) {
		v24[safeIndex(891, v24.Length)] := -v4;
		var v25: seq<bool> := [v1, v1];
		v24[safeIndex(891, v24.Length)], v1, globalState.f4, globalState.f9 := v4 - v4, v1 !in v25[safeIndex(v4, |v25|) := v1], safeDivisionInt(v4, |fm11(-v4, globalState)|), v1;
		match (fm12(v4, globalState)) {
			case DC3(cf3, cf4) =>
				var v26: set<bool> := {if (v4 in v12) then v12[v4] else cf3};
				globalState.f9 := v26 != v26;
				v14 := v14;
				v24[safeIndex(891, v24.Length)] := |fm11(v24[safeIndex(891, v24.Length)], globalState)|;
				var v27 := DC3(v25[safeIndex(v24[safeIndex(891, v24.Length)], |v25|)], false);
				var v28 := new C0(if (v27.cf4) then 'l' else 'k');
			case DC4(cf5, cf6) =>
				var v29: C1 := new C1(v15);
				v29 := v29;
				var v30: array<int> := new int[18];
				v30, globalState.f9, v25, globalState.f9 := v30, v1, v25, v1;
				v1 := v1;
				var v31: set<multiset<bool>> := {v2, v2};
				var v32: multiset<multiset<bool>> := multiset{v2, v2};
				v31 := set v33 : multiset<bool> | v33 in v32 :: (v33);
			case DC2(cf2) =>
				v17 := v17 + v17;
				var v34: C1 := new C1(v15);
				v34 := v34;
				var v35: C0 := new C0(cf2);
				v35 := v35;
				var v36: map<multiset<bool>, int> := map[v2 := v4];
				var v37 := v34.m0(|v36|, v17 + (seq(abs(0x396), i9  => (cf2))), v1 <== v1, v24[safeIndex(891, v24.Length)], globalState);
		}
		
		v24[safeIndex(891, v24.Length)] := v24[safeIndex(891, v24.Length)];
		var v38 := 'f';
		var v39 := DC3(v17 <= v17[safeIndex(v4, |v17|) := v38], v17 < v17[safeIndex(v4, |v17|) := v38]);
		v39 := v39;
		var v40: array<int> := new int[25];
		v40 := v40;
	} else {
		var v41: set<string> := {v17, v17, v17};
		v41 := fm13(839, v4, v17, v1, globalState);
		var v42 := 'w';
		var v43: C0 := new C0(v42);
		var v44: seq<C0> := [v43];
		var v45: array<C0> := new C0[21] [v43, v43, v43, v43, v43, v43, v44[safeIndex(v4, |v44|)], v43, v43, v43, v43, v43, v43, v43, v43, v43, v43, v43, v43, v43, v43];
		v45 := v45;
		globalState.f4 := fm2(v1, 942, v4, "p", globalState) + v4;
		var v46 := m1(v4, v0, globalState);
		var v47 := DC6(v46);
		var v48 := m1(safeDivisionInt(-v4, v4), v47.cf8, globalState);
	}
	
	var v49: C1 := new C1(v15);
	v49 := v49;
	var v50: seq<bool> := [v1, v1];
	v1 := v50 != v50;
	if (v1) {
		globalState.f10 := v1;
		var v52: set<int> := {|(map v51 : int | (-0x34 <= v51) && (v51 < -65) :: (v51 + v4) := (|map[false := |v17|]|))| * v4};
		v52 := v52 * v52;
		var v53 := 'y';
		var v54 := DC4(v4, v53);
		var v55: map<seq<int>, int> := map[[0xe] := if (v1) then v4 else v54.cf5];
		v55 := v55[((([v4])[safeIndex(v4, |[v4]|) := v4])[safeIndex(v4, |([v4])[safeIndex(v4, |[v4]|) := v4]|) := v4])[safeIndex(|[v1]|, |(([v4])[safeIndex(v4, |[v4]|) := v4])[safeIndex(v4, |([v4])[safeIndex(v4, |[v4]|) := v4]|) := v4]|) := |map[v1 := v1]|] := -0x72];
		var v56: C0 := new C0(v53);
		var v57: array<seq<int>> := new seq<int>[11](i10 => v14[safeIndex(v4, |v14|) := v4]);
		var v58: seq<array<seq<int>>> := [v57];
		var v59: array<array<seq<int>>> := new array<seq<int>>[18] [v57, v57, v57, v58[safeIndex(0x13c, |v58|)], v57, v57, v57, v57, v57, v57, v57, v57, v57, if (!v1) then v57 else v57, v57, v57, v57, v57];
		v59[safeIndex(757, v59.Length)] := v57;
		var v60: array<array<int>> := new array<int>[15];
		var v61: map<int, C0> := map[-v4 := v56];
		v56, globalState.f10, v59[safeIndex(757, v59.Length)], v60, globalState.f4 := if (|fm8(v1, v17, globalState)| in v61) then v61[|fm8(v1, v17, globalState)|] else if (v1) then v56 else v56, if ((if (v1) then v1 else v1) in v18) then v18[if (v1) then v1 else v1] else v1, v57, v60, v4;
		var v62 := new C1(v49.f11);
	} else {
		var v63: array<array<bool>> := new array<bool>[13] [v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0, v0];
		v63 := v63;
		var v64: seq<string> := ["jfb"];
		var v65 := DC1(v12);
		if (fm0(v1, v17 < v64[safeIndex(v4, |v64|)], -0x2f4, fm2(fm0(v1, v1, v4, v4, globalState), |[v65]|, |v50|, v17, globalState), globalState)) {
			var v66 := 'w';
			var v67 := new C0(v66);
			globalState.f4 := v4;
			var v69: array<map<int, int>> := new map<int, int>[20](i11 => map[v4 := v4] + (map v68 : int | (-874 <= v68) && (v68 < 754) :: (safeModuloInt(v68, v4)) := (v4)));
			var v70: map<int, int> := map[v4 := v4];
			v69[safeIndex(565, v69.Length)] := v70;
			var v71: seq<C1> := [v49];
			var v72: map<int, array<bool>> := map[v4 := v0];
			var v73 := DC8(v72);
			var v74: map<multiset<C1>, int> := map[multiset(v71) := |(v73.cf11 + v72)|];
			var v77: map<char, bool> := map[v67.f12 := v1];
			var v78: map<multiset<bool>, map<char, int>> := map[v2 := map v76 : char | v76 in v77 :: (v76) := (v4)];
			var v79: set<int> := {|v17|, v4, v4};
			var v80: array<C0> := new C0[5];
			var v81: map<bool, array<C0>> := map[v1 := v80];
			var v82: map<map<bool, array<C0>>, map<multiset<C1>, int>> := map[v81 := v74];
			var v83: multiset<C1> := multiset{v49};
			v69[safeIndex(565, v69.Length)], v3, v74 := v70, map[v2 := v1] + ((map v75 : multiset<bool> | v75 in v78 :: (v75) := (v1)) + fm14(true, v4, v79, globalState)), if (v81 in v82) then v82[v81] else v74[v83 := v4];
			var v84 := m1(v4, v0, globalState);
			var v85 := DC2(v67.f12);
			var v86: multiset<D2> := multiset{v85};
			globalState.f10 := !((v86 * v86[v85 := abs(v4)]) > v86);
		} else {
			var v87: map<bool, int> := map[if (v1) then v1 else v1 := v4];
			var v88 := 'n';
			globalState.f10, globalState.f4, globalState.f10, v24, v87 := !v1, v4, v17 == v17[safeIndex(v4, |v17|) := v88], v24, v87 + map[v1 := v4];
			var v89: array<C1> := new C1[4];
			v89[safeIndex(264, v89.Length)] := if (v1) then v49 else v49;
			v24[safeIndex(183, v24.Length)] := |fm15(globalState)|;
			var v90: set<array<bool>> := {v0};
			v89[safeIndex(264, v89.Length)], v1, v24[safeIndex(183, v24.Length)], v90, globalState.f10 := v49, fm0(v1, v1, v4, v4, globalState), safeModuloInt(fm2(v1, v4, v4, v17, globalState), |(seq(abs(-0x144), i12  => (v4)))| * v4), v90, true;
			var v91 := m1(v24[safeIndex(183, v24.Length)], v0, globalState);
			var v92: seq<array<bool>> := [v0, v91, v91, v0];
			var v93 := m1(|v17|, v92[safeIndex(v4, |v92|)], globalState);
			v24[safeIndex(183, v24.Length)] := 0x19c;
		}
		
		v4 := v14[safeIndex(v4, |v14|)];
		var v94: multiset<int> := multiset{v4, v4};
		var v95: seq<multiset<int>> := [v94];
		var v96: map<seq<multiset<int>>, bool> := map[v95 := !!(if (v1) then v1 else v1)];
		v96 := v96[v95 := v1];
		v4 := v4;
	}
	
	v4 := v4;
	globalState.f4 := v4;
	if (v1) {
		var v97 := m1(v4, v0, globalState);
		v24[safeIndex(144, v24.Length)] := v4;
		v24[safeIndex(144, v24.Length)] := v4;
		globalState.f10 := v1;
		var v98 := new C1(v49.f11);
		globalState.f4 := if (v1) then v4 else v4;
	} else {
		var v99 := 'r';
		var v100: map<char, int> := map[fm5(v4, globalState) := 0x2b4];
		var v101: map<array<bool>, map<bool, bool>> := map[v0 := v18];
		var v102: map<map<char, int>, int> := map[if (v1) then map[v99 := -0x1fd] else v100 := |v101|];
		v102 := v102[v100 := v4];
		globalState.f4 := v4;
		globalState.f3 := v0;
		var v103 := v49.m0(v4, v17 + v17, !v1, v4, globalState);
		globalState.f10 := v49.f11.cf0;
	}
	
	globalState.f4 := fm2(v1, v4, v4, v17, globalState);
	print v0[8], "\n";
	print globalState.f0[8], "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3[8], "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9, "\n";
	print globalState.f10, "\n";
	print v1, "\n";
	print v2 == multiset{true}, "\n";
	print v3 == map[multiset{true} := false, multiset{false} := false, multiset{true, false} := false], "\n";
	print v4, "\n";
	print v12 == map[123 := false], "\n";
	print v13 == map[true := "du"], "\n";
	print v14 == [384, 670], "\n";
	print v15.cf0, "\n";
	print v16 == map[DC0(true) := 384], "\n";
	print v17, "\n";
	print v18 == map[true := true], "\n";
	print v24[0], "\n";
	print v24[1], "\n";
	print v24[2], "\n";
	print v24[3], "\n";
	print v24[4], "\n";
	print v24[5], "\n";
	print v24[6], "\n";
	print v24[7], "\n";
	print v24[8], "\n";
	print v24[9], "\n";
	print v24[10], "\n";
	print v24[11], "\n";
	print v24[12], "\n";
	print v24[13], "\n";
	print v24[14], "\n";
	print v24[15], "\n";
	print v24[16], "\n";
	print v24[17], "\n";
	print v24[18], "\n";
	print v24[19], "\n";
	print v24[20], "\n";
	print v24[21], "\n";
	print v24[22], "\n";
	print v24[23], "\n";
	print v24[24], "\n";
	print v24[25], "\n";
	print v24[26], "\n";
	print v24[27], "\n";
	print v24[28], "\n";
	print v49.f11.cf0, "\n";
	print v50 == [false, false], "\n";
}