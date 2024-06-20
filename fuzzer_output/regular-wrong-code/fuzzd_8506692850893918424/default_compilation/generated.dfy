datatype D0 = DC1(cf1: bool, cf2: int, cf3: bool, cf4: int) | DC0(cf0: char)
datatype D1 = DC3(cf6: int, cf7: bool) | DC2(cf5: map<bool, int>)
datatype D2 = DC5(cf9: int, cf10: int, cf11: D0) | DC6 | DC4(cf8: array<int>)
datatype D3 = DC8(cf13: int) | DC7(cf12: map<bool, array<int>>) | DC9(cf14: D3)
class GlobalState {
	const f0 : multiset<int>
	const f1 : array<int>
	const f2 : bool
	var f3 : int
	const f4 : int
	const f5 : int
	var f6 : int
	var f7 : int
	var f8 : bool
	const f9 : bool
	const f10 : set<string>
	const f11 : bool
	const f12 : int
	const f13 : int
	constructor (f0 : multiset<int>, f1 : array<int>, f2 : bool, f3 : int, f4 : int, f5 : int, f6 : int, f7 : int, f8 : bool, f9 : bool, f10 : set<string>, f11 : bool, f12 : int, f13 : int) {
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

function fm1(p0: int, p1: bool, globalState: GlobalState): seq<char> {
	['v'] + ['a', 'w']
}
function fm2(p0: string, p1: bool, p2: bool, globalState: GlobalState): char {
	's'
}
function fm3(p0: bool, p1: bool, p2: D0, p3: bool, globalState: GlobalState): set<bool> {
	({!false} * {!false}) + {true}
}
function fm4(p0: string, p1: seq<int>, p2: bool, p3: bool, globalState: GlobalState): bool {
	(0x224 + |[0x19c]|) >= -0x13a
}
function fm5(p0: int, p1: bool, p2: bool, globalState: GlobalState): multiset<multiset<int>> {
	multiset([multiset{0x1a5, -0x11}] + [multiset{|[false, !false, false]|, |[false]|}])
}
function fm6(p0: int, p1: bool, globalState: GlobalState): int {
	0x64 + DC3(0x153, !!!!true).cf6
}
function fm7(p0: bool, p1: bool, globalState: GlobalState): D1 {
	match DC8(|multiset{-463, -|"u"|, 0x353, |["cdurb", "amh"]|, 0x14b}|) {
		case DC8(cf13) => DC3(cf13, false)
		case DC7(cf12) => DC3(|multiset{true, true}|, false)
		case DC9(cf14) => DC3(|multiset{|{!false}|}|, false)
	}
}
function fm8(p0: bool, p1: map<int, int>, p2: int, globalState: GlobalState): seq<int> {
	[0x2ff % |"nux"|, |{-0x367, |"rqqxre"|}| * |map[!true := true]|]
}
function fm9(p0: int, p1: int, p2: bool, p3: bool, globalState: GlobalState): seq<D0> {
	match DC6() {
		case DC5(cf9, cf10, cf11) => [DC0('y'), DC0('j'), DC0('w'), DC0('h')]
		case DC6() => [DC0('b'), DC0('p')]
		case DC4(cf8) => [DC0('y'), DC0('r')]
	}
}
method m0(globalState: GlobalState) returns (r0: bool, r1: set<int>, r2: string) {
	var v0 := 0x1fa;
	var v1 := "skwjs";
	var v2: map<int, string> := map[v0 % |v1| := "k"];
	var v3: seq<string> := [v1];
	v2 := v2[v0 * 0x3d := v3[v0] + v1];
	var v4: array<set<D1>> := new set<D1>[17];
	var v5 := true;
	var v6 := DC3(-v0, v5);
	var v7: map<bool, set<D1>> := map[v5 := {v6}];
	v4[874] := if (v5 in v7) then v7[v5] else {v6, v6.(cf6 := v0), v6};
	var v8: set<D1> := {v6};
	v4[874] := v8;
	var i0 := 0;
	while (v5 <==> v5)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		v0 := -v0;
		var v9: array<int> := new int[29](i1 => i1 / v0);
		v9[54] := v0;
		v9[54] := v0;
		var v10 := 'k';
		var v11 := DC0(v10);
		var v12: seq<D0> := [v11];
		var v13: seq<int> := [v9[54]];
		var v14: seq<bool> := [true, v5, v5, !v5];
		var v15: seq<int> := [v9[54], v13[|map[0x28a := v5]|], |v14|, v0, |v1|];
		v12, v5 := (v12 + fm9(v0, v9[54], v5, v5, globalState)) + v12, fm4(v1, (seq(0x32b, i2  => (|v1|))) + v13, v5, v5, globalState);
		r0 := v5;
	}
	var v16: seq<int> := [v0];
	globalState.f6 := (0x2c2 / v0) / (v0 - -|map[!fm4(seq(0x1dc, i3  => ('r')), v16, v5, v5, globalState) := v0]|);
	var v17: C0 := new C0();
	var v18: map<int, C0> := map[v0 := v17];
	var v19: multiset<map<int, C0>> := multiset{v18, v18};
	var v20: set<bool> := {v5};
	var v21: map<int, bool> := map[if (v18 in v19) then v19[v18] else fm6(v0, v5, globalState) := {!v5} > v20];
	if (if ((v0 * v0) in v21) then v21[v0 * v0] else true) {
		var v22 := new C0();
		var v23 := DC1(v0 == v0, v0, v5, v0 * v0);
		v23 := v23;
		var v24: seq<bool> := [false, v5];
		var v25: seq<seq<bool>> := [[v5, v5], [false]];
		var v26: seq<C0> := [v17];
		v5, r0, globalState.f6 := v24 !in v25, v5, |(seq(0x373, i4  => ('c')))| / |v26|;
		var v27: multiset<int> := multiset{v0, v0, v0};
		var v28: multiset<bool> := multiset{v5};
		var v29 := 'g';
		var v30: map<char, bool> := map[v29 := !true];
		var v31: array<bool> := new bool[23] [v5, v6.cf7, v1 == v1, if (!v5) then true else v5, v5, v5, v5, v5, v0 != v0, |v16| < v0, v5, v5, v5, !(v5 || false), -v0 > |v20|, v27 !! v27, v5, v28 > v28, v5, true, if (v29 in v30) then v30[v29] else v5, !v5, v5];
		v31[602] := v5;
		v31[602] := v5 <== v5;
		var v32: map<bool, char> := map[v5 := v29];
		r0, v31[602], v31[602], r2, v32 := v31[602], (-v0 - v0) < v0, |v1| >= v0, v1[v0 := v29], v32;
	} else {
		globalState.f7 := v6.cf6;
		globalState.f6 := -v16[v0];
		var v33: array<int> := new int[21](i5 => i5 - v0);
		var v34: multiset<array<int>> := multiset{v33, v33};
		if (multiset{v33} == v34) {
			var v35: seq<bool> := [v5, v5];
			v35 := v35 + v35;
			v5 := (v0 - v0) < v0;
			globalState.f8 := (|v35| * v0) >= 0xf5;
			v33[699] := v0 - v0;
			var v36: map<array<int>, int> := map[v33 := v0 / v0];
			var v37: seq<array<int>> := [v33];
			var v38: map<int, int> := map[|v37| := v0];
			var v39: set<map<int, int>> := {v38, v38};
			var v40: map<bool, bool> := map[v5 := v5];
			var v41: map<bool, C0> := map[v5 := v17];
			r0, v33[699], globalState.f6, r0, globalState.f8 := v5, if (v33 in v36) then v36[v33] else v0, v0, v5, DC1(v5, |v39|, if (fm4(v1, [v0], v5, v5, globalState) in v40) then v40[fm4(v1, [v0], v5, v5, globalState)] else v5, |v41|).cf1;
			v35 := v35 + v35;
		} else {
			globalState.f6 := v0;
			var v42: seq<C0> := [v17, v17];
			v17 := v42[|(seq(0x1d9, i6  => ('u')))|];
			var v43: seq<bool> := [v5, v5];
			var v44: multiset<int> := multiset{v0};
			var v45: map<seq<bool>, bool> := map[v43 := v44 !! v44];
			globalState.f3, globalState.f3, v45 := v0 / (-380 / v0), |(v1 + v1)| * v0, v45;
			var v46 := new C0();
			globalState.f3 := v0;
		}
		
		var v47: array<array<bool>> := new array<bool>[28];
		var v48: array<bool> := new bool[17](i7 => v5);
		v47[702] := v48;
		v47[702] := v48;
		var v49 := DC1(v5, 0x22, !v5, -|(v20 + v20)|);
		match (v49) {
			case DC1(cf1, cf2, cf3, cf4) =>
				r2 := v1;
				v48 := v47[702];
				globalState.f3 := |"vspevqskp"|;
				var v50: map<bool, int> := map[cf3 := cf4];
				v2 := v2[|v50| * v0 := seq(655, i8  => ('s'))];
			case DC0(cf0) =>
				globalState.f3 := v16[|v20|];
				var v51: multiset<bool> := multiset{v5};
				var v52: set<multiset<bool>> := {v51, v51, v51};
				var v53: map<bool, set<multiset<bool>>> := map[true := v52];
				v52 := if (fm4("x", v16, !!v5, !v5, globalState) in v53) then v53[fm4("x", v16, !!v5, !v5, globalState)] else v52;
				globalState.f8 := v5;
				var v54: array<array<seq<bool>>> := new array<seq<bool>>[4];
				var v55: array<seq<bool>> := new seq<bool>[19](i9 => [false]);
				v54[478] := v55;
				var v56 := DC2(map[fm4(v1, v16, v5, v5, globalState) := v0]);
				var v57: seq<array<seq<bool>>> := [v55, v55];
				var v58: seq<array<seq<bool>>> := [v55, v57[-v0], v55];
				v54[478], globalState.f6, v56 := v57[v0], fm6(v0, v5, globalState) * fm6(v0, v5, globalState), v56;
		}
		
	}
	
	var v59: seq<set<bool>> := [v20, v20];
	var i10 := 0;
	while (|v16| <= |v59[v0]|)
		decreases 100 - i10
	{
		if (i10 >= 100) {
			break;
		}
		
		i10 := i10 + 1;
		var v60: array<int> := new int[7](i11 => i11 * v0);
		var v61: multiset<array<int>> := multiset{v60, v60, v60, v60};
		v60, v0, globalState.f8, globalState.f8, globalState.f6 := DC4(v60).cf8, v0, v5, v5, |v61[v60 := v0 - v0]|;
		if (v5) {
			var v62: map<int, int> := map[292 := v0];
			var v63 := 'p';
			var v64: multiset<char> := multiset{v63};
			var v65: array<bool> := new bool[21] [v5, v5, fm4("owms", v16, v5, v5, globalState), v5, v5, v5, v5, |v62| > v0, false, v5, v64 > v64, v5, v5, v5, v5, v5, v5, v5, v5, v5, v5];
			v65[836] := if (v5) then v5 else true;
			v65[836] := v5;
			v17 := v17;
			var v66: array<C0> := new C0[23];
			v66[624] := v17;
			var v67: seq<bool> := [v5, v65[836]];
			var v68: multiset<bool> := multiset{v5, false};
			var v69: set<int> := {|v67|, |v68|, v0};
			v17, v66[624], globalState.f3, globalState.f3 := v17, v17, (|v69| * v0) + v0, 0x2c4 * |v21|;
			globalState.f6 := v0;
			var v70 := new C0();
		} else {
			globalState.f8 := v5;
			globalState.f3 := v0;
			var v71: set<int> := {v0, -v0};
			v5 := |v71| in (v16 + v16);
			var v72 := DC1(false, -v0, !v5, v0);
			var v73 := DC5(v0, v0, v72);
			v60[531] := v0 % v73.cf9;
			v60[531] := |"poflqgduh"|;
			var v74: array<int> := new int[19];
			var v75: map<bool, array<int>> := map[false := v74];
			var v76: seq<map<bool, array<int>>> := [v75, v75];
			var v77 := DC7(v75);
			var v78: seq<bool> := [v5];
			var v79: array<map<bool, array<int>>> := new map<bool, array<int>>[20] [v75, v75, v75, v75, v75, v75, map[fm4(seq(-688, i12  => ('k')), [|v21|, v0], v5, v5, globalState) := v74], v75, map[v5 := v74], v75 + v76[v0], v75, v75, v77.cf12, v75, v75, v75, v75 + map[true := v74], v75 + v75, map[v78[v60[531]] := v74], v75];
			v79[438] := v75 + map[v5 := v74];
			v79[438] := v75;
		}
		
		var v80 := DC1(true, v0, v5, 0x350);
		globalState.f7 := v80.cf2 % 34;
		var v81 := new C0();
	}
	r0 := v5;
	r1 := set v82 : int | v82 in ([|map[v0 := v0]|] + v16) :: (v82 / 669);
	var v83: map<bool, int> := map[v5 := v0];
	r2 := (fm1(|v83|, v5, globalState) + v1) + v1;
}
class C0 {
	constructor () {
	}
	
	function fm0(p0: bool, p1: string, globalState: GlobalState): seq<bool> {
		(if (false) then [true] else [false]) + ([true] + [true])
	}
}

method Main() {
	var v0 := -0xb7;
	var v1: multiset<int> := multiset{v0};
	var v2: array<int> := new int[19];
	var v3 := "cqq";
	var v4: set<string> := {"rwijw", v3};
	var globalState := new GlobalState(v1, v2, false, 0x1b2, 0x24c, 0x366, 0x4d, -390, false, true, v4, true, 980, 0x2ce);
	var v5 := false;
	var v6: map<map<int, bool>, int> := map[map[v0 := !v5] := v0];
	var v7: map<int, bool> := map[-0x1a5 := false];
	v6 := v6[v7 := -v0];
	var v8 := new C0();
	var v9: array<bool> := new bool[28];
	forall i0 | 0 <= i0 < v9.Length {
		v9[i0] := v5;
	}
	var v10, v11, v12 := m0(globalState);
	var i1 := 0;
	while (v10)
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		v10 := !(v0 < v0) <== !v5;
		var v13, v14, v15 := m0(globalState);
		globalState.f3 := |"tiuwpcrys"|;
		var v16: map<int, string> := map[v0 := v12];
		v12 := v15 + (if (v0 in v16) then v16[v0] else seq(0x2eb, i2  => ('k')));
	}
	var i3 := 0;
	while (if (v7 == map[v0 := true]) then v10 else v10)
		decreases 100 - i3
	{
		if (i3 >= 100) {
			break;
		}
		
		i3 := i3 + 1;
		globalState.f3 := v0;
		var v17 := new C0();
		v3 := v3 + ((seq(-0x9b, i4  => (DC0('x').cf0))) + (seq(-0x254, i5  => ('s'))));
		var v18: map<bool, bool> := map[v5 := v5 <==> v10];
		globalState.f6 := |v18|;
	}
	var i6 := 0;
	while (v5)
		decreases 100 - i6
	{
		if (i6 >= 100) {
			break;
		}
		
		i6 := i6 + 1;
		var v19 := 'v';
		var v20 := DC0(v19);
		v20, globalState.f8 := v20, false;
		v3 := v3;
		var v21: array<array<int>> := new array<int>[5];
		v21[128] := v2;
		v21[128] := v2;
		v5 := v5;
	}
	var v22 := 'q';
	var v23 := DC0(v22);
	v23 := v23;
	v10 := v5;
	var v24 := DC1(v10, v0, true, v0);
	var i7 := 0;
	while (match v24 {
		case DC1(cf1, cf2, cf3, cf4) => v10
		case DC0(cf0) => v5
	})
		decreases 100 - i7
	{
		if (i7 >= 100) {
			break;
		}
		
		i7 := i7 + 1;
		globalState.f6 := v0;
		var v25 := new C0();
		var v26: map<int, C0> := map[v0 := v25];
		globalState.f8 := v0 in (v26 + v26);
		if (v5) {
			globalState.f8 := v10;
			var v27: array<seq<char>> := new seq<char>[5];
			v27[578] := v3;
			v27[578], globalState.f8, globalState.f3 := (fm1(v0, v10, globalState))[v0 := fm2(v12, v5, v10, globalState)], v10, (v0 - v0) - |v3|;
			v5 := v10;
			var v28: array<D0> := new D0[15] [v23, DC0(v22), v23, v23, v23, v23, v23, v23, DC0(v22), v23, DC0(v22), v23, DC0('c'), v23, v23];
			v28[170] := DC0(v22);
			v27[578], v28[170] := v3 + v3, v23;
			globalState.f7 := v0;
		} else {
			v24 := v24;
			var v29 := new C0();
			v22 := v22;
			var v30: seq<bool> := [v5];
			globalState.f8 := v30 != v8.fm0(v10, v3, globalState);
			var v31 := new C0();
		}
		
	}
	var v32: set<bool> := {v5};
	var v33: seq<set<bool>> := [{v5, v10, v10, v10, v10}, v32, v32];
	if (v33 != (v33 + [v32])) {
		var v34: map<string, bool> := map[v3 := v10];
		v34 := v34;
		if (!(fm3(v10, v10, v24, v5, globalState) >= v32)) {
			v32 := v32;
			var v35: seq<int> := [v0, -v0, v0];
			globalState.f8 := fm4(v12, v35, false, v5, globalState);
			var v37: array<map<bool, int>> := new map<bool, int>[25](i8 => DC2(map[v5 := v0]).cf5 + map[v10 := |(map v36 : seq<char> | v36 in [v12[-v0 := v22], [v22]] :: (v36) := (DC1(!v5, v0, v5, |map[v22 := |map[false := false]|]|)))|]);
			v37 := v37;
			v9[98] := v10;
			v9[98] := v5;
			v2[833] := |("ngxstq" + v12)[v0 := v22]|;
			v2[833] := -v0;
		} else {
			var v38: array<char> := new char[16](i9 => v22);
			v38[867] := v22;
			v38[867] := v22;
			globalState.f6 := if (true) then -|[fm2(seq(0x37f, i10  => (v38[867])), v10, if (v12 in v34) then v34[v12] else fm4(v12, [v0], v10, true, globalState), globalState)]| else v0;
			var v39: seq<string> := [v3, (seq(966, i11  => (v38[867]))) + (seq(-0xdb, i12  => (v38[867])))];
			v39 := [v3, v3, (seq(0x117, i13  => (v22))) + "ngtw", seq(0x30e, i14  => (v38[867])), v12];
			var v40: seq<bool> := [!v5];
			globalState.f8 := (v32 != {!v10}) in (v40 + v40);
			var v41: array<multiset<multiset<int>>> := new multiset<multiset<int>>[2];
			var v42: multiset<multiset<int>> := multiset{v1};
			v41[706] := v42 - multiset{v1, v1};
			v41[706] := fm5(330, v10 !in v32, v5 && true, globalState);
		}
		
		var v43, v44, v45 := m0(globalState);
		v43 := v5;
		var v46: array<D0> := new D0[12];
		v46[27] := v24;
		var v47: seq<int> := [v0];
		var v48: seq<bool> := [fm4("okbhypfhd", v47, v5, v5, globalState)];
		var v49: seq<array<bool>> := [v9, v9];
		globalState.f7, v22, v5, v46[27], v9 := (|v45| * |multiset(v48)|) - (if (v0 in v1) then v1[v0] else fm6(0x383, v5, globalState)), 't', v43 <== v43, DC1(v43, v0, v5, v0), v49[v0];
	} else {
		globalState.f6 := fm6(fm6(v0, v10, globalState), v10, globalState) + (if (v5) then v0 else |"nowxxf"|);
		v9[856] := v5 <== v10;
		v9[856] := false;
		var v50: array<seq<bool>> := new seq<bool>[22](i15 => [v9[856], v9[856], v5]);
		v50[762] := [v10];
		var v51: seq<bool> := [v10];
		v50[762] := v51 + v51;
		var v52: seq<C0> := [v8, v8];
		v52 := v52;
		globalState.f3 := -v0;
	}
	
	var v53: array<array<D1>> := new array<D1>[8];
	var v54: array<seq<C0>> := new seq<C0>[18];
	var v55: seq<C0> := [v8];
	v54[689] := if (v10) then v55 else v55;
	var v56: map<int, int> := map[|[v22, v22]| := v0];
	var v57: seq<int> := [v0];
	v5, v53, v10, v54[689], v5 := match fm7(v5, v5, globalState) {
		case DC3(cf6, cf7) => v10
		case DC2(cf5) => v11 !! v11
	}, v53, v5, ((v55 + v55) + v55)[v0 := v8], if (fm4(fm1(v0, v10, globalState), fm8(v10, v56[0x13c := v0], v0, globalState), v5, fm4(v12, v57, v10, false, globalState), globalState)) then [v0] <= v57 else !v5;
	var v58, v59, v60 := m0(globalState);
	v0 := -v0;
	var v61 := new C0();
	var v62: multiset<bool> := multiset{v58};
	var v63 := DC3(|v62|, v58);
	var v64: array<set<bool>> := new set<bool>[6] [v32, v32, v32 * v33[v0], v32, v32, v32 - fm3(v10, !v5, v24, v63.cf7, globalState)];
	forall i16 | 0 <= i16 < v64.Length {
		v64[i16] := v32;
	}
}