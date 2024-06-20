datatype D0 = DC0(cf0: int, cf1: bool) | DC1(cf2: set<set<bool>>, cf3: bool) | DC2(cf4: set<int>, cf5: int, cf6: bool) | DC3(cf7: char)
datatype D1 = DC5(cf9: bool) | DC4(cf8: string)
datatype D2 = DC7(cf11: bool, cf12: bool, cf13: bool) | DC6(cf10: seq<bool>)
datatype D3 = DC9(cf15: char) | DC8(cf14: C1) | DC10(cf16: D3)
datatype D4 = DC12(cf18: int, cf19: char, cf20: multiset<int>) | DC11(cf17: set<seq<int>>) | DC13(cf21: D4)
datatype D5 = DC15(cf23: set<int>, cf24: seq<C1>) | DC16 | DC14(cf22: array<array<int>>)
datatype D6 = DC17(cf25: set<T0>)
datatype D7 = DC19(cf27: bool) | DC18(cf26: multiset<char>)
datatype D8 = DC21(cf29: seq<bool>, cf30: char, cf31: set<bool>) | DC22(cf32: bool, cf33: map<map<bool, int>, bool>, cf34: int) | DC20(cf28: seq<seq<int>>) | DC23(cf35: D8)
datatype D9 = DC25(cf37: bool, cf38: bool, cf39: bool) | DC26(cf40: bool, cf41: int) | DC24(cf36: seq<int>) | DC27(cf42: D9)
class GlobalState {
	const f0 : int
	const f1 : int
	const f2 : multiset<multiset<int>>
	const f3 : bool
	var f4 : char
	var f5 : set<array<char>>
	var f6 : set<seq<int>>
	const f7 : char
	const f8 : multiset<multiset<bool>>
	const f9 : bool
	const f10 : int
	const f11 : seq<bool>
	var f12 : array<int>
	const f13 : bool
	constructor (f0 : int, f1 : int, f2 : multiset<multiset<int>>, f3 : bool, f4 : char, f5 : set<array<char>>, f6 : set<seq<int>>, f7 : char, f8 : multiset<multiset<bool>>, f9 : bool, f10 : int, f11 : seq<bool>, f12 : array<int>, f13 : bool) {
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
		this.f11 := f11;
		this.f12 := f12;
		this.f13 := f13;
	}
	
}

function fm0(p0: map<bool, bool>, p1: bool, p2: bool, globalState: GlobalState): int {
	-184
}
function fm3(p0: int, globalState: GlobalState): D1 {
	match DC7(false, false, true) {
		case DC7(cf11, cf12, cf13) => DC4("ljgfacm")
		case DC6(cf10) => if (!false) then DC4("g") else DC4("dewv")
	}
}
function fm4(p0: bool, p1: set<bool>, p2: int, globalState: GlobalState): bool {
	true
}
function fm6(p0: bool, p1: int, globalState: GlobalState): bool {
	(|map[true := |map[true := map v0 : int | (-0x224 <= v0) && (v0 < 0xbb) :: (v0 * 953) := (-0x174)]|]| - |map[false := true]|) in [|map[true := -0x1ae]|]
}
function fm7(globalState: GlobalState): D0 {
	DC2({|"sfekhaf"|, 774, -567, 0x192}, |(multiset{|[false]|, --165} + multiset{|['d']|, |multiset{false}|, |(seq(-0x3, i0  => ('b')))|, 0x2e4})|, false)
}
function fm8(p0: bool, p1: bool, p2: bool, p3: bool, globalState: GlobalState): seq<int> {
	([0x355] + [|"ks"|]) + ([0x79, 0x379, |"xikkpo"|, |[0x102]|, |multiset{!!true, true, true, !true}|] + [0x17b])
}
function fm9(p0: char, p1: int, p2: int, p3: map<bool, int>, globalState: GlobalState): set<bool> {
	{[0x340, |(seq(0xf2, i0  => ('p')))|, --61] < (seq(0x1fe, i1  => (0x102)))}
}
function fm10(p0: bool, p1: char, p2: int, globalState: GlobalState): seq<bool> {
	([false, !true] + [DC1({{true, !true}}, true).cf3]) + [DC2({-|"ajxt"|, -490}, 836, true).cf6]
}
function fm11(p0: string, globalState: GlobalState): multiset<char> {
	(multiset{'o', 'm', 'f'} - multiset{'f'}) * (DC18(multiset{'c'}).cf26 * multiset{'j', 'y', 'j'})
}
function fm12(p0: bool, p1: int, globalState: GlobalState): D4 {
	if (!false) then DC13(DC12(0x3c, 'e', multiset{|(seq(115, i0  => ('t')))|})) else DC13(DC12(|map[-0x2e3 := |[|[380]|, |"bntttie"|]|]|, 'g', multiset{|multiset{|"gix"|}|}))
}
function fm13(p0: D0, p1: bool, globalState: GlobalState): seq<string> {
	(["kem"] + ["cqlomrr", "hxb"]) + (["dugneonhq"] + ["awqmcmgf"])
}
function fm14(p0: int, p1: char, p2: int, p3: int, globalState: GlobalState): set<char> {
	{'d', 'y', 'y', if (false) then 'd' else 'g', 'j'}
}
function fm15(p0: char, p1: map<int, int>, globalState: GlobalState): char {
	'b'
}
function fm16(p0: bool, p1: bool, globalState: GlobalState): multiset<int> {
	multiset{|{!false, false}|} - (multiset{|['g']|} - multiset([|{951}|, 849]))
}
function fm17(p0: bool, p1: multiset<bool>, p2: int, p3: string, globalState: GlobalState): seq<seq<int>> {
	DC20([[|[true, false]|, -|"aapfmbeu"|], [|"lytyvmod"|]]).cf28 + ([[|"rdnvcfxx"|]] + [[|(map v0 : int | v0 in map[0x381 := false] :: (v0 + |(seq(0xd8, i0  => (0x12f)))|) := ('t'))|], [|[|"gwpfdjks"|]|], [0x2fc], DC24([369, -0x2dc]).cf36, [0x3c1, |{false}|, 812, 50]])
}
function fm18(p0: char, globalState: GlobalState): string {
	"dh" + "nrki"
}
method m0(p0: int, p1: seq<bool>, globalState: GlobalState) returns (r0: bool, r1: array<D0>, r2: int, r3: int) {
	r3 := p0;
	r2 := p0;
	var v0 := true;
	var i0 := 0;
	while (v0)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		if (v0) {
			r3 := p0 - (p0 * p0);
			var v1 := "ym";
			var v2: set<int> := {|v1|};
			var v3: map<bool, set<int>> := map[v0 <==> v0 := v2];
			v3 := v3 + map[!v0 := set v4 : int | (-0x346 <= v4) && (v4 < 345) :: (v4 + p0)];
			var v5: array<D4> := new D4[5];
			var v6 := 't';
			var v7: map<int, int> := map[0xa := -0x34c];
			var v8: multiset<int> := multiset{p0};
			var v9 := DC12(p0, fm15(v6, v7, globalState), v8);
			v5[163] := v9;
			v5[163] := v9;
			var v10: map<int, string> := map[768 := v1];
			var v11 := DC3(v6);
			v10 := v10[p0 := fm18(v11.cf7, globalState)];
			var v12: C1 := new C1(if (v0) then v0 else v0, -333);
			v12 := new C1(p0 <= 667, p0);
		} else {
			var v13 := 'u';
			globalState.f4 := v13;
			var v14: array<array<bool>> := new array<bool>[7];
			var v15: map<bool, array<array<bool>>> := map[v0 := v14];
			v15 := v15[v0 || v0 := v14];
			r2 := p0;
			var v16: array<bool> := new bool[29];
			v16 := v16;
			var v17: seq<int> := [p0];
			r0 := p0 in v17;
		}
		
		var v18: array<int> := new int[8];
		var v19 := DC0(p0, v0);
		var v20: multiset<int> := multiset{v19.cf0, p0};
		var v21: map<int, array<int>> := map[|v20| := v18];
		var v22: seq<array<int>> := [v18];
		var v23: map<bool, array<int>> := map[v0 := v18];
		var v24: array<array<int>> := new array<int>[24] [v18, v18, if (p0 in v21) then v21[p0] else v18, v18, v18, v18, v22[p0], v18, v18, v18, v18, v18, v18, if (v0 in v23) then v23[v0] else v18, v18, if (p0 in v21) then v21[p0] else v18, v18, v18, v18, v18, v18, v18, v18, v18];
		var v25 := DC14(v24);
		v25 := v25;
		if (false) {
			var v26 := "esfxlsqkn";
			r3 := |v26| % p0;
			var v27: array<bool> := new bool[12](i1 => v0);
			v27[92] := v0 && fm6(p1[p0], p0, globalState);
			v27[92] := false || false;
			r3 := p0;
			var v28 := 'b';
			var v29: map<bool, char> := map[true := v28];
			var v30: C0 := new C0(v29 == v29);
			var v31 := DC3(v28);
			var v32: array<D0> := new D0[16] [DC3(v28), DC3(v28), v31, DC3(v28), v31, DC3(v28), DC3('c'), v31, v31, v31, v31, v31, v31, DC3(v28), v31, v31];
			var v33: array<array<D0>> := new array<D0>[9] [v32, v32, v32, v32, v32, v32, v32, v32, v32];
			v33[929] := v32;
			v30, v33[929] := v30, v32;
			v27[92] := !v27[92];
		} else {
			r0 := v0;
			v18[438] := p0;
			v18[438] := p0;
			var v34: seq<int> := [479, fm0(map[true := v0], v0, v0, globalState)];
			var v35: set<bool> := {v0, v0};
			var v36: map<bool, bool> := map[v0 := fm4(v0, v35, p0, globalState)];
			var v37: seq<int> := [v18[438], |v34| - fm0(v36, v0, v0, globalState), v18[438] % p0];
			v37 := v37 + v34;
			var v38 := 'g';
			var v39 := "sir";
			var v40: set<multiset<int>> := {v20, v20};
			var v41: array<bool> := new bool[17] [v0, v0, v0, !!(v38 !in v39), v0, v0, fm0(v36, v0, v0, globalState) > p0, v0, p1[p0], v18[438] >= v18[438], v0, v0, p0 <= 0x4d, v20 in v40, v0, v0 || !p1[p0], !v0];
			v18[438], r3, v41, v0 := p0, fm0(v36, !!v0, v0, globalState), v41, if (if (v0) then v0 else p1[696]) then fm6(v0, v18[438], globalState) else v18[438] > v18[438];
			v18[438] := v18[438];
		}
		
		var v42 := 'x';
		var v43: map<int, int> := map[p0 := p0];
		var v44: seq<char> := [fm15(v42, v43, globalState)];
		var v45: map<bool, seq<char>> := map[v0 := v44];
		var v46: set<bool> := {false !in v45, p0 < p0, v0};
		v46 := (v46 + v46) + v46;
	}
	var v47: array<C0> := new C0[3];
	v47 := v47;
	var v48 := "mns";
	var v49 := DC4(v48);
	match (v49) {
		case DC5(cf9) =>
			var v50: set<bool> := {cf9};
			var v51: map<int, bool> := map[-p0 := cf9];
			var v52: set<bool> := {fm4(v0, v50, |multiset{|v51|, p0, p0, p0, p0}|, globalState)};
			if (if (v0) then v50 >= v52 else v0) {
				var v53: multiset<seq<bool>> := multiset{[v0, cf9], p1, [v0, cf9], [cf9, cf9, v0, cf9, !cf9]};
				var v54: multiset<multiset<seq<bool>>> := multiset{v53, multiset{p1}};
				var v55: multiset<bool> := multiset{true};
				var v56: multiset<int> := multiset{-p0, if (v53 in v54) then v54[v53] else if (v0 in v55) then v55[v0] else p0, p0};
				v56 := v56;
				var v57 := 'c';
				var v58: map<string, int> := map[v48[p0 := v57] := p0];
				var v59: map<string, int> := map[fm18(v57, globalState) + v48 := |v58|];
				var v60: set<int> := {p0};
				var v61: set<int> := {p0, |multiset{v60, {p0}}|};
				v59 := v59[if (cf9) then v48 else "bmdwy" := |v60|];
				var v62: array<int> := new int[24];
				v62[948] := p0;
				v62[948] := p0 - (|v48| / -|v55|);
				r2 := p0;
				var v63: C0 := new C0(v0);
				var v64: map<int, C0> := map[-p0 := v63];
				v64 := v64[v62[948] := v63];
			} else {
				var v65: array<int> := new int[5](i2 => i2 + 0x136);
				v65[910] := p0;
				var v66: map<bool, bool> := map[v0 := cf9];
				var v67: map<string, map<bool, bool>> := map[v48 := v66];
				r0, v65[910], r0, r2 := fm6(v0, -fm0(if (v48 in v67) then v67[v48] else v66, p1[|v48|], v0, globalState), globalState), -p0, !cf9, (p0 + fm0(map[!v0 := !true], if (v0 in v66) then v66[v0] else cf9, v0, globalState)) - p0;
				var v68: T0 := new C0(v0);
				r2, v0, v68 := p0 % -220, v0, v68;
				var v69: array<bool> := new bool[7](i3 => v0);
				v69[515] := cf9;
				cf9, v69[515] := cf9, v0;
				r3 := p0 - (p0 / p0);
				v69[515] := !cf9;
			}
			
			var v70: array<bool> := new bool[26];
			v70[116] := cf9;
			v70[116] := p1[0x261 % |v48|];
			var v71: C1 := new C1(v70[116], p0);
			var v72: map<C1, int> := map[v71 := p0];
			var v73: map<bool, int> := map[cf9 := |v72|];
			var v74: seq<int> := [v71.f15];
			var v75 := new C1(fm6(v0, |v73|, globalState), |v74|);
			var v76: set<array<bool>> := {v70};
			var v77: map<set<array<bool>>, map<bool, int>> := map[v76 := v73];
			v77 := v77[v76 := v73];
		case DC4(cf8) =>
			var v78: seq<int> := [p0];
			var v79: map<int, seq<int>> := map[p0 := v78];
			var v80: map<bool, bool> := map[v0 := v0];
			v78 := (if (p0 in v79) then v79[p0] else v78)[fm0(v80, v0, v0, globalState) := p0] + (v78 + v78);
			var v81 := new C1(p0 > |"sm"|, p0);
			var v82: map<int, bool> := map[v81.f15 := v0];
			var v83: map<int, map<int, bool>> := map[v81.f15 := v82];
			var v84: map<int, bool> := map[v81.f15 - |(if (p0 in v83) then v83[p0] else v82)| := v81.f14];
			v82 := v82[p0 := v0];
			var v85 := 'y';
			var v86: T0 := new C0(v81.f14);
			var v87: map<bool, char> := map[v0 := v85];
			var v88: map<T0, map<bool, char>> := map[v86 := v87];
			var v89: multiset<int> := multiset{-|(if (v86 in v88) then v88[v86] else v87)|};
			var v90 := DC12(v81.f15, v85, v89);
			r3 := |v90.cf20| + --v78[v81.f15];
	}
	
	r0 := false;
	var v91 := 'l';
	var v92: multiset<int> := multiset{|p1|};
	var v93 := DC12(p0, v91, v92);
	r0 := !match v93 {
		case DC12(cf18, cf19, cf20) => v0 ==> !p1[cf18]
		case DC11(cf17) => v0
		case DC13(cf21) => v0
	};
	var v94: seq<int> := [p0, p0, p0, p0, p0];
	var v95 := DC0(p0, v0);
	var v96: array<D0> := new D0[14] [v95, DC0(p0, v0), DC0(p0, fm6(v0, p0, globalState)), v95, v95, v95, DC0(p0, fm6(v0, p0, globalState)), v95, v95, v95.(cf0 := p0), v95, v95, DC0(p0, v0), v95];
	r1 := if (if (!v0) then fm4(v0, {v0, v0}, |v94|, globalState) else v0) then if (v0) then v96 else v96 else v96;
	r2 := if (v0) then p0 - |v92| else p0;
	r3 := p0;
}
trait T0 {
	function fm1(p0: string, p1: int, p2: int, globalState: GlobalState): int 
	function fm2(p0: set<bool>, p1: D0, p2: int, globalState: GlobalState): map<char, string> 
	method m1(globalState: GlobalState) 
}

class C0 extends T0 {
	const f16 : bool
	constructor (f16 : bool) {
		this.f16 := f16;
	}
	
	function fm1(p0: string, p1: int, p2: int, globalState: GlobalState): int {
		|multiset{f16, !f16}|
	}
	function fm2(p0: set<bool>, p1: D0, p2: int, globalState: GlobalState): map<char, string> {
		(map v0 : char | v0 in ['p', 'm', 'q', 'g'] :: (v0) := ("tmqlcma")) + (map['r' := "qky"] + map['l' := seq(0x1d3, i0  => ('r'))])
	}
	function fm5(p0: bool, p1: int, p2: int, p3: bool, globalState: GlobalState): string {
		"tv" + (if (f16) then "usgh" else "xjrsoedx")
	}
	method m1(globalState: GlobalState) {
		var v0: array<bool> := new bool[4](i0 => f16);
		v0[848] := f16;
		var v1 := DC5(f16);
		v0[848] := match v1 {
			case DC5(cf9) => |{f16}| == 0x1bc
			case DC4(cf8) => f16
		};
		var v2 := 0x60;
		var v3: set<int> := {v2, v2};
		var v4: seq<set<int>> := [v3];
		var v5 := "oobutohe";
		var v6: seq<int> := [--v2];
		var i1 := 0;
		while (|(v4[|v5|] + v3)| <= |v6|)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v7: array<seq<bool>> := new seq<bool>[14](i2 => [f16, f16] + [v0[848]]);
			var v8: seq<bool> := [f16, false, v0[848], !v0[848], fm6(v0[848], v2, globalState)];
			v7[22] := v8 + v8;
			v7[22] := v8[v2 := v0[848]];
			var v9, v10, v11, v12 := m0(-(v2 * v2), v7[22], globalState);
			var v13: array<int> := new int[4];
			var v14 := 'f';
			var v15: map<string, char> := map["wluaxmjb" := v14];
			v13[462] := v2 * |v15|;
			v13[462] := v2 % |(v6 + v6)|;
			var v16: multiset<int> := multiset{|v5|, v13[462], 876};
			match (fm7(globalState).(cf5 := |v16|)) {
				case DC0(cf0, cf1) =>
					v0[848] := v9;
					var v17: map<array<int>, bool> := map[v13 := !v9];
					v17 := v17[v13 := fm6(v9, |v3|, globalState)];
					var v18 := DC0(|{v2, 682}|, v9);
					cf1 := !!!!v18.cf1;
					var v19, v20, v21, v22 := m0(v13[462] - cf0, [v0[848]], globalState);
				case DC1(cf2, cf3) =>
					v2 := 868;
					var v23: array<map<char, bool>> := new map<char, bool>[7](i3 => map[v14 := v0[848]]);
					v23 := v23;
					var v24 := DC3(v14);
					v24 := v24;
					globalState.f12 := v13;
				case DC2(cf4, cf5, cf6) =>
					v5 := "qf";
					v13[462] := v11;
					var v25: array<char> := new char[2] [v14, v14];
					v25[885] := v14;
					v25[885] := v14;
					var v26: map<bool, set<int>> := map[f16 := {|v5|, v2, cf5}];
					var v27 := DC2(cf4, |v8|, v9);
					var v28 := DC2(if (v9 in v26) then v26[v9] else v3, v27.cf5, true);
					v28 := fm7(globalState);
				case DC3(cf7) =>
					var v29, v30, v31, v32 := m0(v11, [v0[848]], globalState);
					cf7 := v5[104];
					var v33: map<int, bool> := map[-0x386 := if (!v9) then f16 else f16];
					v33 := v33 + v33;
					v0 := v0;
			}
			
		}
		var v34: array<int> := new int[26];
		v34[332] := v2;
		v34[332] := fm1(v5, v2, if (f16) then |v5| else v2, globalState);
		var v35 := 'f';
		var v36: map<bool, char> := map[f16 := v35];
		globalState.f4 := if (f16 in v36) then v36[f16] else 'v';
		var v37: map<bool, int> := map[v34[332] > v34[332] := v2];
		var v38: map<set<int>, int> := map[v3 := -717];
		var v39 := DC2(v3, v34[332], v0[848]);
		var v40: map<int, bool> := map[v2 := v0[848]];
		v37 := v37[f16 := if (f16) then v34[332] else if (v39.cf4 in v38) then v38[v39.cf4] else |v40|];
		v34[332] := v34[332];
	}
}

class C1 extends T0 {
	const f14 : bool
	const f15 : int
	constructor (f14 : bool, f15 : int) {
		this.f14 := f14;
		this.f15 := f15;
	}
	
	function fm1(p0: string, p1: int, p2: int, globalState: GlobalState): int {
		-|((if (f14) then [[f15, f15, |map[f15 := f14]|, f15]] else [[f15, f15]]) + [[0x200], [|{f14}|, |map[f14 := f15]|]])|
	}
	function fm2(p0: set<bool>, p1: D0, p2: int, globalState: GlobalState): map<char, string> {
		(map v0 : char | v0 in map['x' := seq(861, i0  => ('j'))] :: (v0) := ("evyrhbeq")) + (map['v' := "ycuwefvef"] + map['h' := seq(0x245, i1  => ('x'))])
	}
	method m1(globalState: GlobalState) {
		var v0 := "qqptj";
		var v1: multiset<string> := multiset{v0, v0 + v0, v0};
		var v2: seq<string> := [v0, v0, (fm3(f15, globalState)).cf8];
		v1 := (multiset(v2) * v1) + v1;
		v2 := v2;
		for i0 := f15 to f15 {
			var v3: set<bool> := {f14, f14, f14};
			var v4: multiset<int> := multiset{|v0|};
			var v5 := 'r';
			var v6: map<bool, char> := map[f14 := v5];
			var v7: map<bool, bool> := map[f14 := f14];
			var v8: seq<bool> := [true, f14];
			var v9: array<bool> := new bool[28] [i0 < i0, f14, f14, true, DC2({f15}, 0xb4, true).cf6, f14, f14, f14, f14, v3 >= v3, f14, f14, f14, f14, i0 > (if (488 in v4) then v4[488] else |v6|), !f14 ==> f14, f14, f14, i0 < |v0|, !(|v7| != 0x2f7), f14, v4[i0 := f15] > v4, !(if (f14) then f14 else f14), v8 != v8, !f14, v5 in v0, v8 <= v8, f14];
			v9[466] := f15 >= i0;
			v9[466] := if ((i0 == f15) in v7) then v7[i0 == f15] else f14 || f14;
			var v11: array<map<int, int>> := new map<int, int>[22](i1 => map v10 : int | v10 in {f15, i0, -|v3|, i0, f15} :: (v10 % |multiset{f14, !f14, f14, f14, f14}|) := (-0x1a4));
			v11 := new map<int, int>[19](i2 => map[i0 := DC0(f15, f14).cf0]);
			var v12: seq<int> := [f15];
			var v13 := DC0(|v12|, !f14);
			match (v13.(cf1 := f14)) {
				case DC0(cf0, cf1) =>
					cf1 := if (v9[466]) then !fm4(f14, v3, f15, globalState) else f14;
					globalState.f4 := v5;
					var v14 := new C0(v3 < v3);
					v9[466] := i0 <= i0;
				case DC1(cf2, cf3) =>
					var v15 := 0x3a3;
					var v16: array<map<bool, bool>> := new map<bool, bool>[19];
					v16[865] := v7;
					v15, v16[865], v15, v9[466] := v12[f15], v7 + (v7 + v7), (v15 % i0) * i0, if (f14) then v9[466] else !f14;
					var v17: map<int, set<bool>> := map[i0 := {false}];
					var v18: map<map<int, set<bool>>, int> := map[v17 := f15];
					v18 := v18[v17 := -424];
					var v19: map<multiset<int>, int> := map[multiset{i0} := f15];
					v15 := (f15 * f15) / |v19|;
					cf3 := !cf3;
				case DC2(cf4, cf5, cf6) =>
					cf6 := !(cf4 == cf4);
					globalState.f4 := v5;
					v9[466] := f15 != fm1(v0, f15, 0x25a, globalState);
					var v20: C0 := new C0(v9[466]);
					var v21: seq<C0> := [v20, v20, v20];
					cf6 := if (f15 != |v21|) then v20.f16 else false;
				case DC3(cf7) =>
					v0 := v0;
					v9[466] := f14;
					var v22: multiset<bool> := multiset{if (v9[466]) then v9[466] else true};
					v22 := v22;
					v9 := v9;
			}
			
			var v23: set<set<bool>> := {v3};
			var v24 := DC1(v23 - v23, f14);
			v24 := v24;
		}
		var v25 := 0xfd;
		v25 := v25;
		var v26: set<int> := {v25, v25};
		var v27: map<int, set<bool>> := map[v25 := {f14, f14}];
		var v28: set<bool> := {f14, f14, f14};
		var v29 := DC2(v26, |({f14} + (if (f15 in v27) then v27[f15] else v28))|, f14);
		match (v29) {
			case DC0(cf0, cf1) =>
				var v30: multiset<int> := multiset{f15, v25};
				var v31: map<multiset<int>, int> := map[v30 - v30 := v25];
				v31 := v31 + (v31 + v31);
				if (cf1 || (if (cf1) then cf1 else cf1)) {
					v25 := f15;
					var v32: multiset<bool> := multiset{cf1, false};
					cf0 := |(v32 * v32)|;
					var v33: seq<bool> := [cf1];
					var v34: map<int, bool> := map[|v0| := f14];
					var v35: array<bool> := new bool[25] [!cf1, cf1, cf1, f14, if (cf1) then v33[cf0] else false, f14, cf1, false, !(f15 != -f15), v0 == "wmjh", false, !cf1, f14, cf1 ==> f14, v30 > v30, false, if (cf0 in v34) then v34[cf0] else true, v32 <= v32, cf1, f14, f14, true, if (f14) then cf1 else true, !cf1, cf1];
					var v36 := 's';
					var v37: set<char> := {v36};
					v35 := new bool[15] [cf1, !!f14, f14, cf1, false, cf1 && f14, !f14, true <==> f14, cf1, !cf1, true, f14, cf1, v25 < |v37|, false];
					var v38 := new C0(cf1);
					v35[394] := cf1;
					v35[394] := !cf1;
				} else {
					var v39: multiset<D0> := multiset{v29};
					var v40: seq<D0> := [DC2(v26, cf0, cf1), v29];
					var v41: set<multiset<D0>> := {v39, v39, v39 + multiset(v40), multiset{v29}, multiset{v29}};
					var v42: map<bool, bool> := map[cf1 := cf1];
					v41, v0, cf0, v28, cf1 := {v39 - v39}, v0, fm0(v42 + v42, cf1 || f14, (seq(0x9e, i3  => (DC3('k')))) < (seq(-891, i4  => (DC3('m')))), globalState), v28, f14;
					cf1 := !(fm0(v42, !!f14, cf1, globalState) == fm0(v42, f14, f14, globalState));
					cf0 := f15 % |fm8(f14, f14, f14, cf1, globalState)|;
					cf1 := f14;
					cf0 := f15;
				}
				
				var v43: array<bool> := new bool[26];
				v43 := new bool[7](i5 => f14);
				var v44: array<int> := new int[4];
				v44[24] := |(seq(-0x98, i6  => (v0[v25])))|;
				v44[24] := -(if (f14) then cf0 else if (cf1) then |[f14]| else cf0);
			case DC1(cf2, cf3) =>
				var v45 := 'u';
				var v46 := new C0((seq(0x99, i7  => ('x'))) < v0[f15 := v45]);
				var v47: array<bool> := new bool[10];
				v47 := v47;
				v47[591] := v46.f16;
				v47[591] := v46.f16;
				var v48: T0 := new C0(v46.f16);
				var v49: seq<T0> := [v48];
				v49 := v49;
			case DC2(cf4, cf5, cf6) =>
				var v50: seq<int> := [v25, |v26|, cf5, cf5];
				var v51: seq<bool> := [false, fm4(!f14, v28, f15, globalState)];
				var v52: map<bool, seq<bool>> := map[f14 := v51];
				var v53: array<int> := new int[15] [f15, cf5, cf5, if (f14) then v25 else f15, v25, 0x396, f15, 0x2ce, v25 - cf5, |v50|, f15, -v25, 0x43, |v52| - v25, v25];
				v53[241] := f15;
				var v54: array<string> := new string[15];
				v54[678] := v0;
				var v55 := 'd';
				var v56: multiset<char> := multiset{v55};
				var v57: map<bool, int> := map[f14 := v25];
				cf6, v28, v53[241], v54[678] := v56 !! v56[v55 := cf5], fm9(v55, f15, |[cf6]|, v57, globalState) - v28, -|([v0, v0] + (v2 + v2))|, v0;
				var v58: map<array<int>, bool> := map[v53 := f14];
				var v59: array<bool> := new bool[23] [cf6, f14, f14, if (v53 in v58) then v58[v53] else !cf6, !true, cf6, fm4(true, v28, |v54[678]|, globalState), !!(cf5 !in v50), cf6, f14, false, cf6, f14, cf5 == v53[241], v54[678] != v54[678], f14, f14, if (cf6) then cf6 else f14, f14 <== cf6, f14, false, cf6, cf6];
				v59[274] := f14;
				v59[445] := fm4(f14, v28, f15, globalState);
				var v60 := DC6([cf6]);
				var v61: multiset<bool> := multiset{!f14};
				v29, v53[241], cf6, v59[274], v59[445] := v29.(cf5 := 0x168, cf6 := f14 <== f14), f15, false, true, if (multiset(v60.cf10) < v61) then v61 > v61 else !false;
				var v62: array<C0> := new C0[16];
				var v63: C0 := new C0(!true);
				v62[851] := v63;
				v62[851] := v63;
				var v64: map<bool, string> := map[v59[274] := v54[678]];
				cf5 := |(v64[cf6 := v54[678]] + (v64 + v64))|;
			case DC3(cf7) =>
				var v65 := DC7(fm6(f14, v25, globalState), f14, |"lawxj"| != 0x333);
				v65 := v65;
				if ((v25 / v25) != f15) {
					var v66 := new C0(f14);
					var v67: seq<bool> := [f14, !(v66.f16 <== f14), v66.f16, -f15 != f15];
					var v68: array<bool> := new bool[20];
					v68[777] := f14;
					v67, v68[777], v66, v68 := v67, f15 <= v25, v66, v68;
					var v69 := new C0(v68[777]);
					var v70 := new C0(v69.fm5(false, v25, 0x2b, f14, globalState) <= v0);
					var v71 := DC0(f15, v69.f16);
					var v72 := new C0(v71.cf1 || fm6(v70.f16, v25, globalState));
				} else {
					v25 := if (f14) then f15 else v25 % v25;
					var v73: map<int, int> := map[v25 := v25];
					v73 := map[f15 := f15];
					var v74 := new C0(false);
					var v75: seq<bool> := [f14, fm4(v74.f16, v28, 0xe7, globalState)];
					var v76, v77, v78, v79 := m0(v25, v75, globalState);
					var v80 := DC3(cf7);
					v80 := v80;
				}
				
				var v81: map<bool, int> := map[f14 := f15];
				var v82: map<bool, bool> := map[f14 := f14];
				var v83: C0 := new C0(f14);
				var v84: map<int, C0> := map[v25 := v83];
				var v85: array<bool> := new bool[16] [f14 ==> f14, f14, f14, fm4(f14, v28, |v81|, globalState) <== f14, f14, !f14, if (f14 in v82) then v82[f14] else f14, f14, f14, f14, f14, f14, |v84| <= -0xca, f14, -f15 > |(seq(0x1b6, i8  => (cf7)))|, !v83.f16];
				v85[51] := f14;
				v85[51] := f14;
				var v86: map<bool, string> := map[fm4(true, v28, -0x123, globalState) := if (v85[51]) then "kyo" else v0];
				v86 := v86[f14 := v0 + v0];
		}
		
		var v87 := false;
		var v88: array<bool> := new bool[23](i9 => v87);
		var v89: multiset<array<bool>> := multiset{v88, v88};
		var v90: multiset<int> := multiset{fm1(v0, f15, v25, globalState)};
		v87 := v89[v88 := -|v90|] <= v89;
	}
}

method Main() {
	var v0 := -0xfa;
	var v1: multiset<multiset<int>> := multiset{multiset{v0}};
	var v2: array<char> := new char[29];
	var v3: set<array<char>> := {v2};
	var v4: seq<int> := [v0];
	var v5: set<seq<int>> := {v4};
	var v6 := false;
	var v7: multiset<bool> := multiset{v6, v6};
	var v8: multiset<multiset<bool>> := multiset{v7};
	var v9: seq<bool> := [!v6];
	var v10 := "jxxycbwg";
	var v11: map<bool, bool> := map[v6 := true];
	var v12: map<int, int> := map[v0 := |"ue"|];
	var v13: array<int> := new int[8] [v0, |v10|, |map[v6 := v11]|, v0, |v4|, |v10|, |v12|, v0];
	var globalState := new GlobalState(870, -430, v1, false, 'x', v3, v5 + v5, 'u', v8, false, 270, v9, v13, true);
	v13[749] := v0;
	v13[749] := (-|v10| * 0x1ab) % -v0;
	var v14: map<int, bool> := map[v13[749] := v6];
	var v15: set<int> := {|v14|};
	var v16 := DC2(v15, v0, v6);
	match (if (v16.cf6) then v16 else v16) {
		case DC0(cf0, cf1) =>
			var v17: array<bool> := new bool[5];
			v17[151] := false;
			var v18: multiset<string> := multiset{v10 + v10};
			v17[50] := v6;
			v17[151], v13[749], cf0, v18, v17[50] := cf1, cf0 % v0, |v10| * -(v0 % cf0), v18, v0 == (v0 % v13[749]);
			if (!cf1) {
				var v19, v20, v21, v22 := m0(|v10| % -6, v9, globalState);
				var v23, v24, v25, v26 := m0(fm0(v11, cf1, false, globalState) * -0x118, v9, globalState);
				var v27: multiset<map<int, int>> := multiset{v12, v12, v12};
				v27 := v27;
				v22 := v25;
				v13[749] := v22;
			} else {
				var v28 := 'l';
				var v29: multiset<char> := multiset{v28};
				var v30: map<multiset<char>, string> := map[v29 := v10];
				v13[749] := |v14[v13[749] := v17[151]]| + |(if (cf1) then v30 else v30)|;
				var v31, v32, v33, v34 := m0(-|v5|, v9, globalState);
				var v35 := new C0(cf1);
				var v36: map<bool, int> := map[!v6 := cf0];
				var v37: map<map<bool, int>, string> := map[v36 := v10];
				cf1, cf0 := v31, -(v34 % |(if (v36 in v37) then v37[v36] else v10)|);
				var v38: set<map<bool, bool>> := {v11};
				var v39 := new C0(v38 <= v38);
			}
			
			v13[749] := 0x216;
			var v40 := DC6([v17[151]]);
			var v41: map<int, D2> := map[cf0 * |v4| := v40];
			var v42 := 'r';
			cf0, globalState.f4, cf0 := |v41|, v42, cf0 / v13[749];
		case DC1(cf2, cf3) =>
			var v43: set<bool> := {!v6};
			if (!((v43 - v43) !! {!v6, cf3})) {
				var v44: array<D0> := new D0[12] [if (cf3) then v16 else v16, v16, v16, v16, v16, v16, v16, v16, DC2(v15, v13[749], v6), DC2(v15, fm0(v11, cf3, v6, globalState), false), v16, DC2(v15, v0, cf3).(cf4 := v15)];
				v44 := v44;
				v13[749] := v13[749];
				v0 := (v0 - v0) % v13[749];
				v10 := v10;
				var v45 := new C0(fm4(v6, v43, |multiset{v13[749]}|, globalState));
			} else {
				var v46 := DC5(v13[749] <= v13[749]);
				v46 := v46;
				v0 := -(|v14[v0 := v6]| % v0);
				var v47: C1 := new C1(cf3, v13[749]);
				var v48: map<C1, char> := map[v47 := v10[v0]];
				var v49 := DC8(v47);
				var v50 := 'w';
				globalState.f4, v0, v6, v13[749] := if (v49.cf14 in v48) then v48[v49.cf14] else v50, v47.f15, fm6(if (fm6(fm4(!cf3, v43, |v4|, globalState), v13[749], globalState)) then v47.f14 else v47.f14, v0, globalState), v13[749] / v47.f15;
				v47.m1(globalState);
				v0 := v13[749];
			}
			
			var v51 := DC4(v10);
			var v52: array<D1> := new D1[7] [DC4(v10), v51, v51, v51, v51.(cf8 := v10), v51, v51.(cf8 := v10)];
			var v53: array<bool> := new bool[29];
			v52, v13, v53 := v52, v13, v53;
			var v54, v55, v56, v57 := m0(v13[749], v9, globalState);
			var v58: seq<seq<bool>> := [v9, v9];
			var v59, v60, v61, v62 := m0(-0x283, v58[v57], globalState);
		case DC2(cf4, cf5, cf6) =>
			var v63: multiset<seq<bool>> := multiset{v9, v9, v9};
			var v64 := DC0(|v9|, v6);
			var v65: seq<seq<bool>> := [v9];
			var v66: map<bool, seq<seq<bool>>> := map[v64.cf1 := v65];
			var v67: array<multiset<seq<bool>>> := new multiset<seq<bool>>[8] [v63, v63 + v63[v9 := 596], if (cf6) then v63 else v63, v63, multiset{v9, v9}, v63, if (cf6) then multiset(if (cf6 in v66) then v66[cf6] else v65) else v63, v63];
			v67[276] := multiset{v9, v9, v9};
			var v68: seq<multiset<seq<bool>>> := [v63, v63 + v63, v63 + v63, v63, multiset{v9, v9}];
			v67[276] := v68[v4[0x1f8] / |[false]|];
			if (v6) {
				var v69 := new C0(fm6(v6, -0x385, globalState));
				v10 := v69.fm5(cf6, cf5, |v11| / cf5, v69.f16, globalState);
				var v70 := DC11(v5);
				var v71 := new C1(fm6(!true, 999, globalState), |v70.cf17|);
				v69.m1(globalState);
				var v72 := new C1(v10 < (seq(0x1b4, i0  => ('w'))), |v11|);
			} else {
				var v73: C0 := new C0(cf6);
				var v74: seq<C0> := [v73, v73];
				cf5, v73, v4, v13[749], v74 := v16.cf5 + v13[749], v73, v4 + (v4 + v4), cf5, [v73, v73, v73];
				v13[749] := |([cf5] + v4)|;
				cf6 := cf6;
				cf6 := v73.f16 <== !v6;
				v0 := v0 % v0;
			}
			
			v13[749] := v0;
			if (cf6) {
				v0 := v64.cf0;
				var v75: array<bool> := new bool[23](i1 => !v6);
				v75[380] := cf6;
				v75[380] := !v6;
				v9 := v9[cf5 := false];
				v6 := if (0x9d in v14) then v14[0x9d] else v75[380];
				var v76 := 'j';
				var v77: seq<string> := [v10];
				var v78: array<D0> := new D0[10];
				var v79: map<int, multiset<bool>> := map[v0 := v7];
				var v80: map<array<D0>, map<int, multiset<bool>>> := map[v78 := v79];
				globalState.f4, v75, v10, v6, v6 := v76, v75, v10 + v77[cf5], |[v75[380]]| in (if (v78 in v80) then v80[v78] else v79), v75[380];
			} else {
				var v81 := 'h';
				v2[172] := v81;
				v2[172] := v81;
				cf5 := -v13[749];
				var v82: multiset<int> := multiset{|v14|, -cf5};
				globalState.f12, v6, v13[749], v82 := v13, !false, cf5, v82;
				var v83, v84, v85, v86 := m0(-v13[749], fm10(!cf6, v2[172], v0, globalState), globalState);
				v6 := v6 <==> true;
			}
			
		case DC3(cf7) =>
			var v87: array<T0> := new T0[8];
			v87 := v87;
			v6 := v0 > v0;
			var v88: map<bool, map<int, bool>> := map[v9[v13[749]] := if (v6) then v14 else v14];
			v88 := v88 + v88;
			v6 := !(v13[749] > (v0 / v0));
	}
	
	v6 := v6 || (v13[749] != v13[749]);
	var i2 := 0;
	while (v6)
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		v12 := v12[v0 * v0 := v13[749] - v13[749]];
		v6 := v0 in (if (v6) then v14 else v14);
		var v89: array<bool> := new bool[11](i3 => v0 < v13[749]);
		v89 := new bool[16](i4 => v6);
		var v90: C1 := new C1(v6, |(seq(0x231, i5  => ('d')))|);
		v90, v13[749] := v90, 0x1fd;
	}
	var v91: array<bool> := new bool[8](i6 => v6);
	var v92: C0 := new C0(false);
	var v93: array<C0> := new C0[2] [if (v6) then v92 else v92, v92];
	v93[538] := v92;
	var v94: C1 := new C1(v6, v0);
	var v95: seq<C1> := [v94, v94];
	var v96 := 'c';
	var v97: map<int, char> := map[|v95| := v96];
	v91, v93[538], globalState.f4 := v91, v92, if (v0 in v97) then v97[v0] else v10[v4[-0x102]];
	v92.m1(globalState);
	match (v16) {
		case DC0(cf0, cf1) =>
			v0 := cf0;
			var v98, v99, v100, v101 := m0(v94.f15, v9 + [cf1, !cf1, v6, v94.f14, v94.f14], globalState);
			v13[749] := 0x187 / cf0;
			v10 := "fxrlkl";
		case DC1(cf2, cf3) =>
			var v102: array<set<bool>> := new set<bool>[2];
			var v103: set<bool> := {false, v94.f14};
			v102[723] := {cf3, v92.f16} - v103;
			v102[723] := if (v94.f14) then v103 else v103;
			v13[749] := v94.f15;
			var v104 := DC5(true);
			var v105: map<D0, int> := map[DC3(v96) := v94.f15];
			var v106 := DC3(v96);
			var v107 := new C1(v104.cf9, if (v106 in v105) then v105[v106] else v13[749]);
			var v108: array<map<set<int>, T0>> := new map<set<int>, T0>[11];
			var v109: T0 := new C0(v92.f16);
			var v110: map<set<int>, T0> := map[{v94.f15} := v109];
			v108[337] := v110;
			v0, v0, v108[337], v13[749] := v94.f15, |(map v111 : string | v111 in [v10, v10] :: (v111) := (DC1(cf2, v94.f14)))|, (v110 + v110) + v110, |v7|;
		case DC2(cf4, cf5, cf6) =>
			var v112 := new C0(v6);
			v6 := !(v112.f16 && v92.f16);
			v6 := |fm9(v96, v0, 904, map[!v94.f14 := cf5], globalState)| >= v94.f15;
			var v113: array<map<string, array<int>>> := new map<string, array<int>>[13];
			var v114: map<string, array<int>> := map[v10 := v13];
			v113[703] := v114;
			v113[703] := map[v10 := v13] + (v114 + v114);
		case DC3(cf7) =>
			v13[749] := v94.f15;
			if (v94.f14) {
				var v115 := DC10(DC9(cf7));
				var v116: map<int, D3> := map[|[v92.f16, v9[|v95|]]| := v115.(cf16 := DC8(v94))];
				v116 := v116[v94.f15 := DC10(DC9(v96))];
				var v117: T0 := new C1(v92.f16, v13[749]);
				v117 := v117;
				var v118 := DC8(v94);
				v118 := v118;
				v10 := v10;
				var v119: map<bool, seq<bool>> := map[fm6(v92.f16, v94.f15, globalState) := v9];
				v119 := v119;
			} else {
				var v120: multiset<char> := multiset{cf7};
				var v121: map<bool, multiset<char>> := map[v6 := v120];
				var v122: set<bool> := {false, !v94.f14};
				var v123: seq<multiset<char>> := [if (v92.f16 in v121) then v121[v92.f16] else v120['j' := |v122|]];
				v13[749] := |(v123 + v123)| % v94.f15;
				v0 := v94.f15;
				var v124 := DC5(v6);
				var v125: map<int, array<C0>> := map[v94.f15 := v93];
				globalState.f12, v124, v6, v125, v93[538] := v13, v124, if ((if (v92.f16 in v11) then v11[v92.f16] else v92.f16) in v11) then v11[if (v92.f16 in v11) then v11[v92.f16] else v92.f16] else v92.f16, v125 + v125[v0 := v93], v92;
				v0 := v0;
				v13[749], v11, cf7, v10 := (v94.f15 / v0) + (v94.f15 / v94.f15), v11, cf7, v10 + v10;
			}
			
			v13 := v13;
			v6 := v4 < v4;
	}
	
	var v126: multiset<char> := multiset{v96};
	v126 := (v126 * fm11("avyxie", globalState))[v96 := v13[749] * v94.f15];
	v6 := v94.f15 > |v10|;
	match (match fm12(v6, v94.f15, globalState) {
		case DC12(cf18, cf19, cf20) => v16
		case DC11(cf17) => DC2(v15, v94.f15, false)
		case DC13(cf21) => v16
	}) {
		case DC0(cf0, cf1) =>
			var v127 := new C1(v92.f16, v94.f15);
			v6 := v6;
			if (v127.f14) {
				var v128: T0 := new C1(true, v94.f15);
				var v129: multiset<int> := multiset{v94.f15, v13[749]};
				var v130 := DC12(v94.f15, v96, v129);
				var v131: map<seq<int>, D4> := map[v4 := v130];
				var v132: map<T0, int> := map[v128 := |v131|];
				cf1 := v132 == v132;
				var v133: map<bool, seq<bool>> := map[v127.f14 := v9];
				var v134: map<map<bool, seq<bool>>, array<int>> := map[v133 := v13];
				v13 := if (v133[v94.f14 := v9] in v134) then v134[v133[v94.f14 := v9]] else if (!fm6(v127.f14, v94.f15, globalState)) then v13 else v13;
				v13[749] := -v94.f15 * v94.f15;
				var v135: array<D3> := new D3[28];
				var v136 := DC8(v94);
				v135[844] := v136;
				v135[844] := DC8(v127);
				globalState.f4 := 'w';
			} else {
				v13[749] := v94.fm1(v10, v127.f15, v94.f15, globalState);
				v94.m1(globalState);
				v6 := cf1;
				v91[321] := !true;
				var v137: multiset<C0> := multiset{v93[538], v93[538], v92, v92, v93[538]};
				v91[321] := v137 !! v137;
				var v138: set<bool> := {v91[321], v92.f16};
				v6 := fm4(v94.f14, v138, v4[v13[749]], globalState);
			}
			
			var v139: map<multiset<bool>, C0> := map[v7 := if (v94.f14) then v93[538] else v92];
			var v140: seq<string> := [v10 + v10, seq(854, i7  => (v96)), seq(240, i8  => (v96))];
			var v141: set<bool> := {v6, cf1, cf1, v92.f16};
			var v142: map<bool, int> := map[v92.f16 := v94.f15];
			var v143: set<set<bool>> := {v141, v141, v141, fm9(v96, cf0, v94.f15, v142[v6 := cf0], globalState), v141};
			cf1, v6, v139, v92, v140 := v15 <= {v94.f15}, fm6(false && v127.f14, cf0, globalState), map[v7 := v93[538]], v93[538], fm13(DC1(v143, true), v9[-0xc4], globalState);
		case DC1(cf2, cf3) =>
			var v144: array<seq<char>> := new seq<char>[5];
			v144[140] := v10;
			v144[140] := v10;
			v13[749] := v94.f15;
			v94.m1(globalState);
			var v145: map<array<bool>, int> := map[v91 := |v10|];
			v145 := v145[v91 := v94.f15];
		case DC2(cf4, cf5, cf6) =>
			v6 := v94.f14;
			v6 := false;
			var v146 := new C1(!(v13[749] < cf5), v0 / v94.f15);
			var v147, v148, v149, v150 := m0(if (v94.f14) then v0 else 0x20c, v9, globalState);
		case DC3(cf7) =>
			v94 := v94;
			v91[615] := true;
			var v151: set<char> := {cf7};
			var v153: seq<set<char>> := [v151, v151];
			var v155: array<set<char>> := new set<char>[20] [v151, v151, v151, v151, {cf7}, fm14(0x1c8, cf7, v94.f15, v94.f15, globalState), v151, set v154 : char | v154 in (map v152 : char | v152 in v153[v94.f15] :: (v152) := (cf7)) :: (v154), v151, {fm15('k', map[v94.f15 := v0], globalState)}, if (v92.f16) then v151 else v151, if (!false) then v151 else v151, v151, v151, v151, v151 + v151, v151, {v96, cf7}, fm14(v0, cf7, |v7|, 0x2a, globalState), v151 + v151];
			var v156: map<seq<bool>, bool> := map[v9 := true];
			var v157: seq<C0> := [v93[538], v92];
			var v158: T0 := new C1(fm6(v94.f14, v94.f15, globalState), |[819, |v156|, |v157|, fm0(v11, !false, v94.f14, globalState)]|);
			var v159: map<C1, T0> := map[if (v94.f14) then v94 else v94 := v158];
			var v160: seq<array<set<char>>> := [v155];
			v9, v91[615], v6, v0, v155 := [v94.f14] + v9, v9[v0], v92.f16, |v159|, v160[v0];
			var v161: set<bool> := {v92.f16, v6, v91[615], v92.f16, v6};
			v6 := fm4(v6 || v92.f16, v161, v0 % -752, globalState);
			var v162: multiset<int> := multiset{|{v94.f14, !v91[615], v94.f14, v92.f16, v91[615]}|, v0};
			var v163: map<char, bool> := map[cf7 := v94.f14];
			var v164: seq<multiset<int>> := [v162];
			var v165: array<multiset<int>> := new multiset<int>[13] [v162, fm16(v6, if (|v163| in v14) then v14[|v163|] else if (v92.f16 in v11) then v11[v92.f16] else v94.f14, globalState) + v162, v162 - multiset{-0x104}, fm16(v94.f14, v94.f14, globalState), v162 - multiset{-0x367}, v164[v94.f15], multiset{fm0(v11, v92.f16, v92.f16, globalState)}, v162 - v162, multiset{v94.f15} * v162, v164[v0], multiset{v0}, v162 + v162, multiset{|v9|, v0}];
			v165[787] := v162;
			v165[787] := v162 * v162;
	}
	
	var i9 := 0;
	while (!v6)
		decreases 100 - i9
	{
		if (i9 >= 100) {
			break;
		}
		
		i9 := i9 + 1;
		var v166 := DC9(v96);
		var v167 := DC10(v166);
		var v168: set<D3> := {DC10(v166), v167};
		var v169: seq<set<D3>> := [v168, v168, v168];
		v13[749] := |v169[if (true) then v13[749] else v13[749]]|;
		var v170: array<array<int>> := new array<int>[19];
		var v171 := DC14(v170);
		v170 := (if (v6) then v171 else v171).cf22;
		var v172 := new C1(v6, v94.f15);
		var v173: map<seq<seq<int>>, array<bool>> := map[fm17(v6, v7, |v10[v172.f15 := v96]|, v10, globalState) := v91];
		var v174: seq<seq<int>> := [v4, [v94.f15, -|v10|]];
		v173 := v173[v174 := v91];
	}
	var i10 := 0;
	while (v6 && false)
		decreases 100 - i10
	{
		if (i10 >= 100) {
			break;
		}
		
		i10 := i10 + 1;
		var v175: map<bool, int> := map[v94.f14 := v0];
		v175 := (v175 + v175[true := v13[749]]) + v175;
		v92.m1(globalState);
		var v176 := new C0(v92.f16);
		var v178 := DC1(set v177 : set<bool> | v177 in (seq(-0x394, i11  => ({v92.f16, v92.f16, !v92.f16}))) :: (v177), v94.f14);
		v91[360] := v178.cf3;
		v91[360] := v94.f14 !in v175;
	}
	var v179: array<map<int, bool>> := new map<int, bool>[28];
	v179[534] := map v180 : int | v180 in v4 :: (v180 + v94.f15) := (v94.f14);
	v179[534] := v14;
	v92.m1(globalState);
	var i12 := 0;
	while (v13[749] != |(v10 + v10[--v13[749] := v96])|)
		decreases 100 - i12
	{
		if (i12 >= 100) {
			break;
		}
		
		i12 := i12 + 1;
		var v181: T0 := new C1(v6, v94.f15);
		var v182: map<int, T0> := map[0x310 := v181];
		var v183 := DC8(v94);
		var v184: array<D3> := new D3[3] [v183, v183, v183];
		var v185: map<T0, array<D3>> := map[if (v0 in v182) then v182[v0] else v181 := v184];
		var v186: map<array<D3>, array<int>> := map[v184 := v13];
		v6 := (if (v181 in v185) then v185[v181] else v184) in (v186 + map[v184 := v13]);
		if (v94.f14) {
			v6 := v94.f14;
			v6 := v92.f16;
			var v187 := new C1(v92.f16, v94.f15);
			v6 := fm6(v94.f14, v94.f15, globalState);
			var v188: set<bool> := {v94.f14};
			var v189: multiset<set<bool>> := multiset{v188};
			v91[171] := v189 > v189;
			v91[171] := v6;
		} else {
			globalState.f4 := v96;
			v13 := v13;
			var v190: seq<seq<bool>> := [([v94.f14])[v4[v94.f15] := v6], v9, v9 + v9, v9];
			v190 := v190 + v190;
			v91[864] := v94.f14;
			var v191: set<bool> := {v6, v94.f14};
			v91[864] := if (v6) then v13[749] > v94.f15 else fm4(v94.f14, v191, v0, globalState);
			v182 := v182[v0 + v13[749] := v181];
		}
		
		if (v92.f16) {
			var v192: seq<T0> := [v181, v181, v181, v181, v181];
			var v193: array<seq<T0>> := new seq<T0>[1] [v192];
			v193[941] := v192;
			v193[941] := v192;
			v10 := v10 + v10;
			v13[749] := v94.f15 * |(map[|"nadw"| := v94.f14] + v179[534])|;
			var v194: array<T0> := new T0[9] [v181, v181, v181, v181, v181, v181, v181, v181, v181];
			v194 := v194;
			var v195 := new C1(v92.f16, if (false) then v94.f15 else -v181.fm1(v10, -0x149, v0, globalState));
		} else {
			var v196 := DC17({v181, v181});
			var v197: set<set<T0>> := {v196.cf25};
			var v198: set<T0> := {v181};
			var v199: seq<set<set<T0>>> := [v197];
			var v200: array<set<set<T0>>> := new set<set<T0>>[25] [v197, v197, v197, v197, v197, v197, v197 * v197, v197, v197, v197 * v197, v197 - {v198}, v197, v197 - v197, v197, v197, v197, v197 - v197, v197, v197, v197, v197, v197, v197, v199[v94.f15], v197];
			v200[982] := v197;
			v200[982] := v197;
			v6 := v92.f16;
			v92.m1(globalState);
			var v201 := DC0(v94.f15, v94.f14);
			var v202: map<bool, map<int, bool>> := map[v94.f14 := map[v94.f15 := v92.f16]];
			v201 := DC0(|(v202[false := v14] + v202)|, v6 <== v92.f16);
			v13[749] := v13[749];
		}
		
		var v203: set<array<bool>> := {v91, v91};
		v203 := v203 * v203;
	}
	var v204 := DC5(v6);
	v6 := match v204 {
		case DC5(cf9) => cf9
		case DC4(cf8) => if (v94.f14 in v11) then v11[v94.f14] else v94.f14
	};
}