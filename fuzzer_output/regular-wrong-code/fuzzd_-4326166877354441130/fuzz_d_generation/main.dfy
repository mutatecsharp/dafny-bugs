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
datatype D0 = DC0(cf0: array<bool>)
datatype D1 = DC2(cf2: bool, cf3: bool, cf4: int, cf5: int) | DC1(cf1: int)
datatype D2 = DC4(cf7: map<multiset<int>, int>, cf8: int, cf9: bool, cf10: int) | DC3(cf6: char) | DC5(cf11: D2)
class GlobalState {
	var f0 : int
	const f1 : set<seq<bool>>
	var f2 : int
	const f3 : int
	var f4 : bool
	const f5 : bool
	const f6 : int
	const f7 : bool
	var f8 : int
	var f9 : seq<map<char, int>>
	const f10 : int
	var f11 : bool
	var f12 : map<int, bool>
	const f13 : array<char>
	var f14 : int
	var f15 : multiset<bool>
	var f16 : string
	const f17 : bool
	const f18 : set<int>
	const f19 : bool
	var f20 : array<char>
	var f21 : multiset<set<int>>
	const f22 : bool
	const f23 : array<int>
	const f24 : bool
	var f25 : int
	const f26 : int
	var f27 : map<int, string>
	constructor (f0 : int, f1 : set<seq<bool>>, f2 : int, f3 : int, f4 : bool, f5 : bool, f6 : int, f7 : bool, f8 : int, f9 : seq<map<char, int>>, f10 : int, f11 : bool, f12 : map<int, bool>, f13 : array<char>, f14 : int, f15 : multiset<bool>, f16 : string, f17 : bool, f18 : set<int>, f19 : bool, f20 : array<char>, f21 : multiset<set<int>>, f22 : bool, f23 : array<int>, f24 : bool, f25 : int, f26 : int, f27 : map<int, string>) {
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
		this.f14 := f14;
		this.f15 := f15;
		this.f16 := f16;
		this.f17 := f17;
		this.f18 := f18;
		this.f19 := f19;
		this.f20 := f20;
		this.f21 := f21;
		this.f22 := f22;
		this.f23 := f23;
		this.f24 := f24;
		this.f25 := f25;
		this.f26 := f26;
		this.f27 := f27;
	}
	
}

class C0 {
	const f28 : map<D0, seq<int>>
	const f29 : int
	constructor (f28 : map<D0, seq<int>>, f29 : int) {
		this.f28 := f28;
		this.f29 := f29;
	}
	
	function fm3(globalState: GlobalState): int {
		--f29
	}
	function fm4(globalState: GlobalState): bool {
		!({f29} > (set v0 : int | (869 <= v0) && (v0 < 0xc4) :: (safeDivisionInt(v0, f29)))) !in [!!!false, false]
	}
}

function fm0(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): char {
	match DC2(!false, !!false, -|"rhjaet"|, 0x2ec) {
		case DC2(cf2, cf3, cf4, cf5) => 'c'
		case DC1(cf1) => DC3('g').cf6
	}
}
function fm1(p0: int, globalState: GlobalState): bool {
	!(([-0x31b] + [0x3b6]) != (if (false) then [|map[|multiset{-194, -0x240}| := !true]|] else [0x2b, -|"koxnyli"|]))
}
function fm2(globalState: GlobalState): set<int> {
	{0x326, -0xb3 + |{0x1f, -|map[--0x163 := 655]|}|}
}
function fm5(p0: bool, p1: char, globalState: GlobalState): seq<int> {
	seq(abs(153), i0  => (|"kfihj"|))
}
function fm6(p0: bool, p1: bool, p2: int, globalState: GlobalState): string {
	match DC1(|"ywoabtk"|) {
		case DC2(cf2, cf3, cf4, cf5) => seq(abs(0x2fa), i0  => ('f'))
		case DC1(cf1) => "cncea"
	}
}
function fm7(p0: int, p1: int, p2: string, globalState: GlobalState): D1 {
	match DC3('u') {
		case DC4(cf7, cf8, cf9, cf10) => DC2(cf9, DC4(cf7, |map[|multiset{cf8}| := cf8]|, cf9, -|{cf9}|).cf9, cf8, 0x162)
		case DC3(cf6) => DC2(!true, true, 0x373, |map["xep" := false]|)
		case DC5(cf11) => DC2(!false, !true, |"cijfftd"|, |map[false := true]|)
	}
}
function fm8(p0: bool, p1: int, globalState: GlobalState): map<int, D1> {
	(map[-|multiset{DC4(map[multiset{-0x1a3} := 0x398], |[false, true]|, false, -0x395).cf9, !false, !false, true}| := DC2(true, true, -480, -|map[|"ttkwy"| := 0x38f]|)] + map[|(seq(abs(66), i0  => ('j')))| := DC2(false, true, 0x29, |map[|{|map[|map[|map[true := !false]| := |map[false := true]|]| := true]|, |multiset{|(map v0 : int | v0 in [|[true]|, 0x27d] :: (v0 + |{|map[true := 334]|}|) := (0x8c))|}|, |(seq(abs(0x341), i1  => ('w')))|}| := |(seq(abs(563), i2  => (map[true := 0x25])))|]|)]) + (map v1 : int | (0x29b <= v1) && (v1 < 0x384) :: (v1 * |(seq(abs(9), i3  => ('r')))|) := (DC2(false, !false, 0x344, |"msq"|)))
}
function fm9(p0: int, p1: seq<map<bool, bool>>, p2: int, globalState: GlobalState): int {
	-match DC4(map[multiset{|(map v0 : int | v0 in [|"bururqtaf"|, |[true]|] :: (v0 - |"g"|) := (0x1d8))|} := |"kabs"|], |[false]|, false, --|[|multiset{915, 0x68}|]|) {
		case DC4(cf7, cf8, cf9, cf10) => cf10
		case DC3(cf6) => -(if (false) then |"sjnmc"| else -0x3)
		case DC5(cf11) => 0x12d
	}
}
method m0(p0: set<string>, p1: int, globalState: GlobalState) {
	globalState.f25 := 0x7a;
	var v0 := 'u';
	var v1: seq<char> := [v0];
	var v2: seq<seq<char>> := [v1];
	var v3: array<bool> := new bool[3];
	var v4 := DC0(v3);
	var v5: seq<int> := [0xf1, |{|v1|}|];
	var v6: map<D0, seq<int>> := map[v4 := v5];
	var v7: C0 := new C0(v6, p1);
	var v8: map<int, int> := map[p1 := |{v7}|];
	var v9: array<int> := new int[5] [|v2[safeIndex(p1, |v2|)]|, 0x33b, p1, safeModuloInt(0x3c9, if (v7.f29 in v8) then v8[v7.f29] else v7.f29), safeModuloInt(p1, v7.f29)];
	v9 := v9;
	var i0 := 0;
	while (!true)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v10: multiset<int> := multiset{v7.f29};
		globalState.f2 := if (-0x55 in v10) then v10[-0x55] else p1;
		var v11: map<C0, string> := map[v7 := v1];
		var v12 := true;
		v0 := fm0(|(v11[v7 := v1] + v11)|, 0x1ea, v12, -v7.f29, globalState);
		var v13: map<string, bool> := map[v1 := v12];
		if ((|v13| - p1) <= v7.f29) {
			globalState.f16 := (seq(abs(904), i1  => (v0)))[safeIndex(v7.f29, |(seq(abs(904), i1  => (v0)))|) := v0];
			var v14: array<map<int, bool>> := new map<int, bool>[19];
			v14 := v14;
			globalState.f25 := v7.fm3(globalState) * v7.f29;
			globalState.f11 := v12;
			globalState.f2 := safeModuloInt(v7.f29, v7.f29);
		} else {
			globalState.f2 := v7.f29 * p1;
			globalState.f11 := v10 >= (v10 + v10);
			v9[safeIndex(294, v9.Length)] := p1;
			v9[safeIndex(294, v9.Length)] := v7.f29;
			globalState.f4 := v12;
			v3[safeIndex(808, v3.Length)] := v12;
			var v15: set<int> := {v7.f29};
			var v16: set<bool> := {v12, v12};
			var v17: multiset<set<bool>> := multiset{v16};
			v3[safeIndex(808, v3.Length)], v15 := v17 !! (v17[v16 := abs(v7.f29)] * v17), v15;
		}
		
		var v18: map<multiset<bool>, int> := map[multiset{v12} := v7.f29];
		var v19: multiset<bool> := multiset{v12};
		globalState.f2 := if (v19 in v18) then v18[v19] else v7.f29 * p1;
	}
	var v20 := false;
	globalState.f4 := v20;
	globalState.f25 := v7.f29;
	forall i2 | 0 <= i2 < v9.Length {
		v9[i2] := i2 * -safeModuloInt(v7.f29, v7.f29);
	}
}
method Main() {
	var v0 := true;
	var v1: seq<bool> := [!v0, v0];
	var v2: set<seq<bool>> := {v1, v1};
	var v4 := 530;
	var v5: map<bool, bool> := map[!v0 := v0];
	var v6: seq<int> := [v4, v4, |v5|];
	var v7: array<char> := new char[7](i1 => 'i');
	var v9: map<map<int, bool>, bool> := map[map v8 : int | v8 in [v4] :: (safeDivisionInt(v8, |v6|)) := (v0) := v0];
	var v10: map<int, bool> := map[v6[safeIndex(v4, |v6|)] := v0];
	var v11: multiset<bool> := multiset{if (v10 in v9) then v9[v10] else v0, v0, v1[safeIndex(v4, |v1|)], v0, v0};
	var v12 := "ic";
	var v13: set<int> := {v4, v4, v4, v4};
	var v14: array<int> := new int[10] [-78, v4, |v13|, v4, v4, v4, 687, v4, v4, v4];
	var v15: map<int, string> := map[v4 := v12];
	var globalState := new GlobalState(-129, v2, 0x200, 0x5a, false, false, 0x394, true, 0x391, seq(abs(0x1d9), i0  => (map['k' := 624])), -0x36c, false, map v3 : int | v3 in v6 :: (safeModuloInt(v3, -v4)) := (false), v7, 0x167, multiset(v1) + v11, v12 + v12, false, v13, false, v7, multiset{v13}, true, v14, true, 0x3de, 255, v15);
	var v16 := 'f';
	var v17: multiset<char> := multiset{if (v0) then v16 else v16, fm0(v4, |v1|, !v0, v4, globalState)};
	v17 := v17 * multiset{v16};
	globalState.f4 := v0;
	globalState.f11 := fm1(v4 * 0x11a, globalState);
	var v18: array<bool> := new bool[5](i2 => !(v13 >= v13));
	v18[safeIndex(17, v18.Length)] := v1[safeIndex(v4, |v1|)];
	v18[safeIndex(17, v18.Length)] := v0;
	var v19: map<int, int> := map[v4 := v4];
	globalState.f25 := |((v19 + v19) + v19)|;
	v18 := v18;
	globalState.f2 := v4 + (v4 + v4);
	globalState.f4, v0, v14, v5 := false, v0, v14, v5[v0 := v0] + v5;
	v18[safeIndex(17, v18.Length)] := v1[safeIndex(|v1|, |v1|)] || (|map[v0 := -v4]| < |v10|);
	v14[safeIndex(687, v14.Length)] := v4;
	v14[safeIndex(687, v14.Length)] := if (v4 in v19) then v19[v4] else v4;
	if (v18[safeIndex(17, v18.Length)]) {
		var v20: map<int, array<bool>> := map[v14[safeIndex(687, v14.Length)] := v18];
		var v21 := DC0(v18);
		var v22: array<array<bool>> := new array<bool>[26] [v18, v18, v18, v18, v18, if (v14[safeIndex(687, v14.Length)] in v20) then v20[v14[safeIndex(687, v14.Length)]] else v18, v18, v18, v18, v18, v18, v18, v18, DC0(v18).cf0, v18, v18, v18, v18, v21.cf0, v18, v18, v18, v18, v18, v18, v18];
		v22[safeIndex(787, v22.Length)] := v18;
		v22[safeIndex(787, v22.Length)] := v18;
		var v24: map<int, set<int>> := map[--v4 := set v23 : int | (0x334 <= v23) && (v23 < 0x21e) :: (v23 * v4)];
		if (if (-|(if (v4 in v24) then v24[v4] else fm2(globalState))| <= -v14[safeIndex(687, v14.Length)]) then v0 else fm1(v4, globalState)) {
			globalState.f11 := !v18[safeIndex(17, v18.Length)];
			v16 := v16;
			var v25: set<string> := {"tvycemat"};
			m0(v25, v4, globalState);
			var v26: multiset<int> := multiset{v14[safeIndex(687, v14.Length)]};
			v19 := v19[|v26| := safeModuloInt(v4, v4)];
			v14[safeIndex(687, v14.Length)] := v4;
		} else {
			var v27: map<bool, int> := map[v18[safeIndex(17, v18.Length)] := v14[safeIndex(687, v14.Length)] + |v1|];
			v27 := v27[v0 := v14[safeIndex(687, v14.Length)]];
			var v28: map<D0, seq<int>> := map[v21 := v6];
			var v29 := DC1(v4);
			var v30 := new C0(v28, v29.cf1);
			v16 := v16;
			var v31: array<set<bool>> := new set<bool>[1];
			v31 := v31;
			var v32: map<char, bool> := map[v16 := v0];
			v32 := v32[v16 := v0];
		}
		
		if (v18[safeIndex(17, v18.Length)]) {
			globalState.f11 := v0 <==> (if (v0) then v18[safeIndex(17, v18.Length)] else v0);
			var v33: array<C0> := new C0[29];
			var v34: map<D0, seq<int>> := map[v21 := ([681, v4])[safeIndex(v14[safeIndex(687, v14.Length)], |[681, v4]|) := |"am"|]];
			var v35: C0 := new C0(v34, v14[safeIndex(687, v14.Length)]);
			v33[safeIndex(859, v33.Length)] := v35;
			v33[safeIndex(859, v33.Length)], globalState.f4 := v35, v0;
			v16 := fm0(v14[safeIndex(687, v14.Length)], v35.fm3(globalState), v18[safeIndex(17, v18.Length)], v35.f29, globalState);
			var v36: seq<string> := [v12, "caxv"];
			var v37 := new C0(v35.f28, if (v18[safeIndex(17, v18.Length)]) then |v1| else |v36[safeIndex(v4, |v36|)]|);
			globalState.f25 := v14[safeIndex(687, v14.Length)];
		} else {
			v14[safeIndex(687, v14.Length)] := v4 * -|(set v38 : int | (0x308 <= v38) && (v38 < 0x33b) :: (safeDivisionInt(v38, |v6|)))|;
			v22 := v22;
			var v39: seq<array<int>> := [v14, v14];
			var v40: map<bool, string> := map[true := v12];
			var v41: multiset<string> := multiset{v12, (if (v0 in v40) then v40[v0] else v12)[safeIndex(v14[safeIndex(687, v14.Length)], |(if (v0 in v40) then v40[v0] else v12)|) := v16]};
			v14, globalState.f11, v14, globalState.f8, v1 := v39[safeIndex(v6[safeIndex(v14[safeIndex(687, v14.Length)], |v6|)], |v39|)], !(v12 in v41), v14, v6[safeIndex(v14[safeIndex(687, v14.Length)], |v6|)], v1;
			var v42: map<D0, seq<int>> := map[v21 := v6];
			var v43 := new C0(v42[v21 := v6] + v42, v4);
			var v44 := DC1(v6[safeIndex(v14[safeIndex(687, v14.Length)], |v6|)]);
			v44 := v44;
		}
		
		v16 := v16;
		var v45 := DC2(v0, !(if (v18[safeIndex(17, v18.Length)] in v5) then v5[v18[safeIndex(17, v18.Length)]] else v0), v4, v4);
		var v46: array<D1> := new D1[13] [v45, v45, v45, v45, DC2(v0, v18[safeIndex(17, v18.Length)], v4, v4), v45, v45.(cf2 := v0), DC2(v18[safeIndex(17, v18.Length)], v18[safeIndex(17, v18.Length)], v14[safeIndex(687, v14.Length)], v14[safeIndex(687, v14.Length)]), v45, v45.(cf3 := v18[safeIndex(17, v18.Length)], cf4 := v6[safeIndex(v4, |v6|)]), v45, v45, v45];
		v46[safeIndex(575, v46.Length)] := v45;
		var v47: multiset<D1> := multiset{DC2(v18[safeIndex(17, v18.Length)], v0, v14[safeIndex(687, v14.Length)], |v6|)};
		v18[safeIndex(17, v18.Length)], v46[safeIndex(575, v46.Length)], globalState.f4, globalState.f4, v18[safeIndex(17, v18.Length)] := v47 >= v47, DC2(v4 != v14[safeIndex(687, v14.Length)], v0, v4, v4 * v4), true && ('p' !in "ruxv"), v18[safeIndex(17, v18.Length)], if (v0) then false else v0 && !v18[safeIndex(17, v18.Length)];
	} else {
		v14[safeIndex(687, v14.Length)], globalState.f25, v1, globalState.f0 := -v14[safeIndex(687, v14.Length)], -(0x10 + safeDivisionInt(|v5|, v14[safeIndex(687, v14.Length)])), v1, safeModuloInt(|(v10 + v10)|, v14[safeIndex(687, v14.Length)]);
		if (v12 < v12) {
			globalState.f11 := v11 >= (multiset(v1))[v18[safeIndex(17, v18.Length)] := abs(-0xc1)];
			var v48 := DC0(v18);
			var v49: map<D0, seq<int>> := map[v48 := v6];
			var v50: multiset<int> := multiset{|v12|, v4, v4};
			var v51: C0 := new C0(v49, if (|v50| in v19) then v19[|v50|] else v4);
			v51 := new C0(map[v48 := v6], v14[safeIndex(687, v14.Length)]);
			v14[safeIndex(687, v14.Length)] := safeDivisionInt(v51.f29, |v13|);
			v19 := v19[|(if (v0) then v1 else v1)| := v4];
			var v52: array<string> := new string[28](i3 => "ycg");
			var v53: map<bool, array<string>> := map[v0 := v52];
			v53 := v53[v18[safeIndex(17, v18.Length)] := v52];
		} else {
			v18[safeIndex(17, v18.Length)] := v12 < v12;
			v14 := v14;
			var v54: map<D0, seq<int>> := map[DC0(v18) := [v4]];
			var v55 := new C0(v54 + v54, v4 + v14[safeIndex(687, v14.Length)]);
			globalState.f4 := fm1(v4, globalState);
			globalState.f11 := v1[safeIndex(|v12| * v14[safeIndex(687, v14.Length)], |v1|)];
		}
		
		var v56: array<C0> := new C0[19];
		var v57 := DC0(v18);
		var v58: map<D0, seq<int>> := map[v57 := v6];
		var v59: C0 := new C0(v58[v57 := fm5(v0, v16, globalState)], v14[safeIndex(687, v14.Length)]);
		v56[safeIndex(672, v56.Length)] := v59;
		v56[safeIndex(672, v56.Length)] := new C0(v59.f28, |v12|);
		var v60: map<char, bool> := map[v16 := v0];
		v60 := v60[v16 := false];
		v14 := v14;
	}
	
	for i4 := if (false) then v14[safeIndex(687, v14.Length)] else v4 to v14[safeIndex(687, v14.Length)] {
		globalState.f16 := fm6(v0, v0, |v5[fm1(|v1|, globalState) := v0]|, globalState);
		v18 := v18;
		var v61: set<array<bool>> := {v18};
		v61 := v61;
		var v62: set<string> := {seq(abs(0x145), i5  => (v16)), v12};
		if (v62 >= v62) {
			m0(v62, safeModuloInt(i4, i4), globalState);
			var v63 := DC0(v18);
			var v64: map<D0, seq<int>> := map[v63 := v6];
			var v65: seq<map<D0, seq<int>>> := [v64];
			var v66: seq<map<D0, seq<int>>> := [v65[safeIndex(v4, |v65|)]];
			var v67: C0 := new C0(v66[safeIndex(i4, |v66|)], v4);
			v67 := v67;
			m0(v62, v67.f29, globalState);
			var v68: set<bool> := {v18[safeIndex(17, v18.Length)], v18[safeIndex(17, v18.Length)], v0};
			v68 := v68 + v68;
			var v69: map<int, D1> := map[i4 := fm7(v67.f29, v67.f29, v12, globalState)];
			var v70: map<bool, map<int, D1>> := map[v0 := v69];
			var v71: seq<map<int, D1>> := [if (!v18[safeIndex(17, v18.Length)] in v70) then v70[!v18[safeIndex(17, v18.Length)]] else v69, fm8(v18[safeIndex(17, v18.Length)], v4, globalState)];
			globalState.f4, globalState.f4, v69, globalState.f16 := (v12 + v12) != (v12 + v12), v18[safeIndex(17, v18.Length)], v71[safeIndex(-v4, |v71|)], v12;
		} else {
			var v72: seq<map<bool, bool>> := [map[v18[safeIndex(17, v18.Length)] := v0]];
			var v73: multiset<int> := multiset{|v1|};
			v19 := v19[if (v18[safeIndex(17, v18.Length)]) then |v1| else fm9(v14[safeIndex(687, v14.Length)], v72, i4, globalState) := |(multiset{v4, v4, v4} + v73)|];
			var v74 := DC0(v18);
			var v75: map<D0, seq<int>> := map[v74 := [i4]];
			var v76: C0 := new C0(v75[v74 := v6], 222 + v4);
			v76 := v76;
			v0 := true;
			v4 := safeModuloInt(i4, v14[safeIndex(687, v14.Length)]);
			globalState.f11, globalState.f25 := v18[safeIndex(17, v18.Length)], i4;
		}
		
	}
	var v77 := DC0(v18);
	var v78: map<D0, seq<int>> := map[v77 := v6];
	var v79: C0 := new C0(v78, v4);
	var v80: seq<C0> := [v79];
	var v81 := DC2(v0, v18[safeIndex(17, v18.Length)], v4, |v80|);
	var v82: seq<D1> := [fm7(v14[safeIndex(687, v14.Length)], v14[safeIndex(687, v14.Length)], seq(abs(0x2f8), i6  => (v16)), globalState), v81, v81];
	var v83: array<seq<int>> := new seq<int>[7];
	v13, v82, v18[safeIndex(17, v18.Length)], globalState.f11, v83 := {|v12[safeIndex(638, |v12|) := v16]|}, [DC2(v0, !v0, v79.f29, v4).(cf5 := v4)], match v81 {
		case DC2(cf2, cf3, cf4, cf5) => cf3
		case DC1(cf1) => v0 && false
	}, !v18[safeIndex(17, v18.Length)], v83;
	globalState.f4, v12, globalState.f8, v16 := v0, if (if (fm1(|v12|, globalState)) then v0 else v0) then v12 + v12 else v12, -v79.f29, v16;
	var v84: map<char, string> := map[v16 := v12];
	globalState.f16 := if (v16 in v84) then v84[v16] else v12;
	globalState.f4 := false;
	print v0, "\n";
	print v1 == [false, true], "\n";
	print v2 == {[false, true]}, "\n";
	print v4, "\n";
	print v5 == map[false := true, true := true], "\n";
	print v6 == [530, 530, 1], "\n";
	print v7[0], "\n";
	print v7[1], "\n";
	print v7[2], "\n";
	print v7[3], "\n";
	print v7[4], "\n";
	print v7[5], "\n";
	print v7[6], "\n";
	print v9 == map[map[176 := true] := true], "\n";
	print v10 == map[1 := true], "\n";
	print v11 == multiset{true, true, true, true, false}, "\n";
	print v12, "\n";
	print v13 == {2}, "\n";
	print v14[0], "\n";
	print v14[1], "\n";
	print v14[2], "\n";
	print v14[3], "\n";
	print v14[4], "\n";
	print v14[5], "\n";
	print v14[6], "\n";
	print v14[7], "\n";
	print v14[8], "\n";
	print v14[9], "\n";
	print v15 == map[530 := "ic"], "\n";
	print globalState.f0, "\n";
	print globalState.f1 == {[false, true]}, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9 == [map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624], map['k' := 624]], "\n";
	print globalState.f10, "\n";
	print globalState.f11, "\n";
	print globalState.f12 == map[0 := false, 1 := false], "\n";
	print globalState.f13[0], "\n";
	print globalState.f13[1], "\n";
	print globalState.f13[2], "\n";
	print globalState.f13[3], "\n";
	print globalState.f13[4], "\n";
	print globalState.f13[5], "\n";
	print globalState.f13[6], "\n";
	print globalState.f14, "\n";
	print globalState.f15 == multiset{false, false, true, true, true, true, true}, "\n";
	print globalState.f16, "\n";
	print globalState.f17, "\n";
	print globalState.f18 == {530}, "\n";
	print globalState.f19, "\n";
	print globalState.f20[0], "\n";
	print globalState.f20[1], "\n";
	print globalState.f20[2], "\n";
	print globalState.f20[3], "\n";
	print globalState.f20[4], "\n";
	print globalState.f20[5], "\n";
	print globalState.f20[6], "\n";
	print globalState.f21 == multiset{{530}}, "\n";
	print globalState.f22, "\n";
	print globalState.f23[0], "\n";
	print globalState.f23[1], "\n";
	print globalState.f23[2], "\n";
	print globalState.f23[3], "\n";
	print globalState.f23[4], "\n";
	print globalState.f23[5], "\n";
	print globalState.f23[6], "\n";
	print globalState.f23[7], "\n";
	print globalState.f23[8], "\n";
	print globalState.f23[9], "\n";
	print globalState.f24, "\n";
	print globalState.f25, "\n";
	print globalState.f26, "\n";
	print globalState.f27 == map[530 := "ic"], "\n";
	print v16, "\n";
	print v17 == multiset{'f'}, "\n";
	print v18[0], "\n";
	print v18[1], "\n";
	print v18[2], "\n";
	print v18[3], "\n";
	print v18[4], "\n";
	print v19 == map[530 := 530, 2 := 530], "\n";
	print v77.cf0[0], "\n";
	print v77.cf0[1], "\n";
	print v77.cf0[2], "\n";
	print v77.cf0[3], "\n";
	print v77.cf0[4], "\n";
	print |v78|, "\n";
	print |v79.f28|, "\n";
	print v79.f29, "\n";
	print |v80|, "\n";
	print v81.cf2, "\n";
	print v81.cf3, "\n";
	print v81.cf4, "\n";
	print v81.cf5, "\n";
	print v82 == [DC2(true, false, 530, 530)], "\n";
	print v84 == map['f' := "icic"], "\n";
}