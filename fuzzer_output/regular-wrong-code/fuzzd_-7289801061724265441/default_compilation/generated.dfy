datatype D0 = DC1(cf1: array<int>, cf2: bool, cf3: array<string>) | DC2(cf4: bool) | DC3(cf5: array<int>, cf6: int) | DC0(cf0: seq<int>) | DC4(cf7: D0)
datatype D1 = DC6(cf9: int, cf10: int, cf11: int) | DC5(cf8: string) | DC7(cf12: D1)
datatype D2 = DC9(cf14: int, cf15: int) | DC8(cf13: seq<bool>)
datatype D3 = DC11(cf17: C0, cf18: C0, cf19: bool) | DC10(cf16: set<int>)
class GlobalState {
	var f0 : string
	const f1 : string
	const f2 : bool
	const f3 : array<array<bool>>
	const f4 : seq<int>
	const f5 : bool
	const f6 : int
	const f7 : bool
	const f8 : array<int>
	const f9 : map<bool, int>
	const f10 : int
	var f11 : string
	const f12 : seq<bool>
	const f13 : int
	const f14 : seq<int>
	const f15 : array<int>
	const f16 : bool
	const f17 : int
	const f18 : map<int, map<char, int>>
	var f19 : seq<string>
	var f20 : int
	const f21 : int
	constructor (f0 : string, f1 : string, f2 : bool, f3 : array<array<bool>>, f4 : seq<int>, f5 : bool, f6 : int, f7 : bool, f8 : array<int>, f9 : map<bool, int>, f10 : int, f11 : string, f12 : seq<bool>, f13 : int, f14 : seq<int>, f15 : array<int>, f16 : bool, f17 : int, f18 : map<int, map<char, int>>, f19 : seq<string>, f20 : int, f21 : int) {
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
	}
	
}

function fm0(p0: int, p1: bool, p2: bool, p3: string, globalState: GlobalState): bool {
	false
}
function fm1(p0: int, p1: bool, p2: set<int>, globalState: GlobalState): int {
	--0x12c
}
function fm2(p0: int, globalState: GlobalState): D0 {
	DC0(seq(905, i0  => (-0x1be)))
}
function fm3(globalState: GlobalState): seq<map<int, int>> {
	([map[|(seq(-787, i0  => ('x')))| := |"fivufydg"|]] + (seq(0x186, i1  => (map[|{--0x1b5, 0x3c7}| := |"pevdpvb"|])))) + ((seq(523, i2  => (map[|[!true]| := 0x1]))) + [map[-719 := |multiset{true, true, false, false, false}|]])
}
function fm5(globalState: GlobalState): seq<bool> {
	[!!true]
}
function fm6(globalState: GlobalState): multiset<string> {
	match DC8([false, false, false]) {
		case DC9(cf14, cf15) => multiset{"gm", "byjjmnp", "l", "dkpsd"}
		case DC8(cf13) => multiset{"ihjrffkga", "xiypy", "kxg"}
	}
}
function fm7(p0: bool, p1: int, globalState: GlobalState): string {
	"ruoqh" + ((seq(0x18b, i0  => ('c'))) + "irmgafus")
}
function fm8(p0: int, p1: bool, p2: map<bool, char>, p3: int, globalState: GlobalState): multiset<int> {
	(multiset{|map[false := 'n']|} + multiset{|['c']|}) - (multiset{|(seq(0x36f, i0  => (|(seq(508, i1  => ('x')))|)))|} * multiset{0x300, 353})
}
method m0(globalState: GlobalState) returns (r0: map<int, array<int>>) {
	var v0 := 0xea;
	var v1 := "kmmfs";
	var v2: multiset<int> := multiset{|v1|};
	var v3 := false;
	var v4: map<bool, int> := map[v3 := v0];
	var v5: multiset<bool> := multiset{v3};
	var v6: seq<bool> := [v3, v3];
	var v7: set<int> := {v0, -v0, |v6|, 0x1a1};
	var v8: array<int> := new int[16] [v0, v0, v0, 0x2b, v0, v0, v0, if (v0 in v2) then v2[v0] else if (fm0(v0, v3, v3, v1, globalState) in v4) then v4[fm0(v0, v3, v3, v1, globalState)] else v0, 554, |v5| + |v1|, if (v3 in v5) then v5[v3] else v0, v0, v0, v0, 0x17c * fm1(-|v6|, v3, v7, globalState), v0];
	var v9: map<bool, array<int>> := map[v3 := v8];
	v8[381] := |v9|;
	v8[381] := v0;
	var v10: array<bool> := new bool[19];
	v10[692] := v3;
	var v11 := DC10(v7);
	var v12: seq<int> := [fm1(v0, v3, v11.cf16, globalState), v0];
	var v13: map<int, map<bool, int>> := map[985 := v4];
	var v14 := 'p';
	v6, globalState.f20, v3, v2, v10[692] := v6, v0, v3, (multiset(v12) + fm8(v12[|(if (v8[381] in v13) then v13[v8[381]] else map[v3 := v0])|], v3, map[v3 := v14], |v6|, globalState)) * multiset{v0}, v3;
	for i0 := |[v0, v0, v8[381]]| to v0 + v8[381] {
		v10[692] := (|multiset{v0}| < i0) && (v10[692] <==> v3);
		var v15: map<bool, array<bool>> := map[v10[692] := v10];
		v15 := v15[v10[692] || v3 := v10];
		var v16: map<array<int>, int> := map[v8 := v8[381]];
		var v17: map<bool, bool> := map[true := v3];
		var v18: map<map<bool, bool>, bool> := map[v17 := v3];
		var v19: seq<map<map<bool, bool>, bool>> := [v18, v18];
		var v20: array<map<array<int>, int>> := new map<array<int>, int>[13] [v16 + v16, v16, v16, v16, v16, v16, map[v8 := 0x60], (v16[v8 := |v2|])[if (v3 in v9) then v9[v3] else v8 := v8[381]], v16, map[v8 := v8[381]], v16, v16[v8 := |v19[v12[|v12|]]|] + v16, v16];
		var v21: C0 := new C0(v10[692]);
		var v22: multiset<C0> := multiset{v21};
		v20[195] := map[v8 := if (v21 in v22) then v22[v21] else v0] + map[v8 := |v17|];
		v3, v3, globalState.f20, v20[195] := v10[692], (v1 + v1) <= ("g")[|v12| := 'e'], fm1(-v8[381] + i0, v6 != [v3], v7, globalState), v16;
		var v23 := DC0(v12);
		var v24: array<D0> := new D0[3] [DC0(v12), v23, fm2(v8[381], globalState)];
		var v25: map<C0, array<D0>> := map[v21 := v24];
		v25 := v25[v21 := v24];
	}
	v10[692] := v10[692];
	var v26: map<int, int> := map[v8[381] := v8[381]];
	var v27: map<int, bool> := map[if (!v3) then v0 else v0 := if (fm0(|v26|, v3, v3, "whejvaw", globalState)) then v3 else true];
	v3 := !(if (v0 in v27) then v27[v0] else v10[692]);
	var v28: map<array<bool>, bool> := map[v10 := fm0(v8[381], v3, false, v1, globalState)];
	v3 := if (v10 in v28) then v28[v10] else v3 <==> v3;
	var v29: C0 := new C0(v3);
	var v30: multiset<C0> := multiset{v29};
	var v31: map<int, array<int>> := map[|v30| := v8];
	r0 := v31;
}
class C0 {
	const f22 : bool
	constructor (f22 : bool) {
		this.f22 := f22;
	}
	
	function fm4(p0: string, p1: bool, p2: map<string, seq<char>>, p3: string, globalState: GlobalState): bool {
		f22
	}
}

method Main() {
	var v0 := "ljlofb";
	var v1: array<array<bool>> := new array<bool>[5];
	var v2 := 0x87;
	var v3: seq<int> := [v2];
	var v4: array<int> := new int[27];
	var v5 := false;
	var v6: map<bool, int> := map[v5 := 0x248];
	var v7 := 't';
	var v8: seq<bool> := [v5];
	var v9: map<char, int> := map[v0[v2] := |"apnhjorfd"|];
	var v10: seq<map<char, int>> := [v9, map[v7 := v2]];
	var v11: map<int, map<char, int>> := map[v2 := v10[v2]];
	var v12: seq<string> := [v0];
	var globalState := new GlobalState(v0, v0, false, v1, v3, true, 0x25c, false, v4, v6 + v6, 503, v0[v2 := v7], v8 + v8[0x9 := v5], 0x3e4, v3, v4, false, 0x3aa, v11 + v11, v12, 0x1d1, 48);
	var v13: array<bool> := new bool[8](i0 => false);
	var v14: seq<array<bool>> := [v13];
	v13 := v14[v2];
	for i1 := |([v2, v2, v2, v2, |map[|v0| := v2]|])[v2 := v2]| % v2 to v2 {
		var v15: map<int, int> := map[v2 := 174];
		v15, globalState.f20 := v15, v2 * i1;
		var v16: seq<array<int>> := [v4];
		var v17: map<int, array<int>> := map[v2 := v16[v2]];
		var v18: seq<array<int>> := [v4, v4, if (v2 in v17) then v17[v2] else v4];
		v18 := v16[|"vpxkaki"| := v4];
		v8 := v8;
		v5 := v5;
	}
	v6 := v6[v5 := v2];
	var v19 := m0(globalState);
	if (false || false) {
		var v20: seq<seq<int>> := [v3 + v3, v3, v3];
		v20 := v20;
		v13[324] := v5;
		v13[324] := v5;
		var v21: multiset<bool> := multiset{v13[324]};
		v13[324] := !!(v5 !in v21);
		v0 := ("dyxqgdxc" + v0)[|(map v22 : int | (0x3aa <= v22) && (v22 < 0x192) :: (v22 * v2) := (v13[324]))| := v7] + v0;
		var v23: array<bool> := new bool[17] [fm0(v2, v13[324], v5, v0, globalState), v5, v5, !v5, v5, v5, fm0(-0xfb, false, v5, v0, globalState), v5, v13[324], v5, v13[324], fm0(|v0|, v13[324], v5, v0, globalState), v5, v13[324], v13[324], v13[324], v5];
		var v24: map<array<bool>, int> := map[v23 := v2];
		v24 := (v24 + v24) + v24;
	} else {
		v4 := new int[24](i2 => i2 / v2);
		var v25: array<multiset<char>> := new multiset<char>[26];
		var v26: map<seq<int>, array<multiset<char>>> := map[v3[-|v0| := v2] := v25];
		var v27: map<seq<seq<bool>>, seq<int>> := map[[v8] := v3];
		var v28: seq<seq<bool>> := [v8];
		v26 := v26[if (v28 in v27) then v27[v28] else v3 := v25];
		var v29: set<array<bool>> := {v13};
		v2 := |v29|;
		v13[602] := v5;
		v13[602] := v5;
		var v30: array<array<array<int>>> := new array<array<int>>[16];
		var v31: array<array<int>> := new array<int>[7] [v4, v4, v4, v4, v4, v4, v4];
		var v32: seq<array<array<int>>> := [v31];
		v30[588] := v32[v2];
		var v33: map<int, int> := map[v2 := |v8|];
		var v34: multiset<bool> := multiset{v13[602]};
		var v35: map<int, bool> := map[if (fm0(v2, v13[602], v13[602], v0, globalState) in v34) then v34[fm0(v2, v13[602], v13[602], v0, globalState)] else v2 := v13[602]];
		var v36: set<int> := {|v6|};
		v13[602], v30[588], v33, globalState.f20 := true, v31, map[v2 % |v35| := v2], 633 / (v2 - fm1(v2, false, v36, globalState));
	}
	
	for i3 := 0x3ad + v2 to v2 {
		var v37 := DC0(v3);
		v3 := (v37.(cf0 := v3)).cf0;
		var v38 := DC3(v4, i3);
		match (v38) {
			case DC1(cf1, cf2, cf3) =>
				v4 := v4;
				v13[928] := cf2 ==> cf2;
				var v39: multiset<bool> := multiset{!v5, -v2 in v3};
				v13[928], v39 := cf2 <== v5, v39 + v39;
				v13[928] := v5 <==> (i3 >= i3);
				v37 := if (cf2) then v37 else fm2(|v8|, globalState);
			case DC2(cf4) =>
				var v40: seq<array<int>> := [v4, v4];
				v40, v13, v5, globalState.f20 := [v4, v4, v4, v4], v13, false, v2 + (i3 - v2);
				v5 := v5 ==> (if (v5) then !v5 else v5);
				globalState.f20 := v2;
				var v41: set<D0> := {v37, v37};
				v41 := v41 + v41;
			case DC3(cf5, cf6) =>
				v5 := !true || (!v5 ==> v5);
				globalState.f20 := i3 * -(i3 - 0x21a);
				cf6, v5 := |v3| + cf6, v5;
				var v42: set<int> := {v2};
				globalState.f20 := fm1(i3, v5, v42, globalState) - i3;
			case DC0(cf0) =>
				var v43: set<int> := {v2, v2, i3};
				v43 := v43;
				var v44: set<bool> := {v5, v8[0x1ee], !v8[v2], v5};
				v5 := v44 >= v44;
				var v45: array<seq<map<int, int>>> := new seq<map<int, int>>[19];
				var v46: map<int, int> := map[v2 := 0x12f];
				var v47: seq<map<int, int>> := [v46, v46];
				v45[335] := v47 + v47;
				v45[335] := fm3(globalState);
				v13[842] := !(if (v5) then v5 else v5);
				var v48 := DC2(v5);
				v13[842] := v48.cf4;
			case DC4(cf7) =>
				var v49: map<int, int> := map[i3 := i3];
				var v50 := new C0(v49 == map[i3 := v2]);
				var v51: multiset<bool> := multiset{v50.f22};
				v5 := (multiset(fm5(globalState)))[v5 := v2] > v51;
				v19 := v19;
				globalState.f20 := i3;
		}
		
		globalState.f20 := if (v5) then |v6| / v2 else v2;
		var v52: set<map<bool, int>> := {v6};
		var v53 := DC5(v0);
		var v54: map<bool, D0> := map[v52 > v52 := DC3(v4, |v53.cf8|)];
		var v55: multiset<bool> := multiset{v5 <==> v5};
		var v56: array<string> := new string[26](i4 => v0);
		var v57 := DC1(v4, v5, v56);
		v54, v55 := map[v5 := if (false) then v38 else v38], multiset(v8) + multiset{v57.cf2, v5};
	}
	v0 := v0;
	var v58 := m0(globalState);
	globalState.f20 := v2;
	var v59 := DC2(false);
	var i5 := 0;
	while (v59.cf4 ==> true)
		decreases 100 - i5
	{
		if (i5 >= 100) {
			break;
		}
		
		i5 := i5 + 1;
		var v60: array<multiset<string>> := new multiset<string>[24];
		v60[801] := fm6(globalState);
		var v61: map<bool, string> := map[v5 := v0];
		var v62: multiset<string> := multiset{if (v5 in v61) then v61[v5] else "lrw"};
		v60[801] := (v62 + v62)[v0 := v2];
		globalState.f11 := v0;
		v5 := v5;
		globalState.f20 := -v3[v2];
	}
	var v63 := m0(globalState);
	if (!!(v2 != 0x1b)) {
		var v64 := m0(globalState);
		if (v5) {
			v4 := v4;
			v5 := v5;
			var v65 := DC0(v3);
			var v66: seq<D0> := [DC4(v65), v65, DC4(v65), v65, v65];
			var v67: map<char, D0> := map[v7 := v66[257]];
			var v68 := DC4(if (v7 in v67) then v67[v7] else v65);
			var v69: array<string> := new string[12];
			var v70 := DC4(if (v5) then v65 else DC1(v4, v5, v69));
			v70 := DC4(v65);
			var v71: set<bool> := {v5};
			var v72: seq<set<bool>> := [v71, v71];
			var v73: array<seq<int>> := new seq<int>[9] [[v2], [v2], v3, v3 + [|v71|], v3 + ([v2])[v2 := |[v8, v8, v8, v8]|], [v2, v2], v3, (seq(0x243, i6  => (v2)))[-v2 := v2], ([v2])[-v2 := |v72|]];
			v73[705] := v3;
			v73[705] := v3[v2 := |[v2]|] + v3;
			var v74: map<int, int> := map[v2 := v2];
			var v75: set<int> := {v2, v2};
			globalState.f20 := -fm1(if (v2 in v74) then v74[v2] else fm1(v2, true, v75, globalState), {v5} != v71, v75, globalState);
		} else {
			var v76: C0 := new C0(v5);
			var v77: map<int, C0> := map[|v3| := v76];
			globalState.f20 := |(if (v5) then v77 else v77)| - 0x3d;
			v76 := new C0(v76.f22);
			var v78 := m0(globalState);
			var v79 := m0(globalState);
			var v80: array<set<int>> := new set<int>[5];
			var v81: set<int> := {v2};
			v80[399] := {v2, |v3|} - v81;
			v80[399] := v81 - v81;
		}
		
		v5 := v5;
		if (v5) {
			var v82 := m0(globalState);
			var v83: multiset<array<int>> := multiset{v4};
			var v84: map<int, bool> := map[-v2 % (if (v4 in v83) then v83[v4] else -v2) := v5];
			v84 := v84[v2 := v5];
			v5 := v5;
			var v85: C0 := new C0(v8[v2]);
			var v86: map<C0, bool> := map[v85 := v5];
			var v87: set<map<C0, bool>> := {v86[v85 := v5], v86, v86};
			var v88: set<int> := {v2};
			var v89: map<int, int> := map[fm1(|v87|, v5, v88, globalState) % |(seq(0x148, i7  => (v7)))| := v2 - |"gkst"|];
			v89 := v89;
			v5 := v0 <= [v7];
		} else {
			globalState.f20 := -v2;
			v5 := v5;
			var v90 := m0(globalState);
			var v91: map<int, char> := map[|v6| % v2 := v7];
			v13[596] := v5;
			var v92: C0 := new C0(v5);
			v13[262] := -v2 != v2;
			var v94: multiset<int> := multiset{v2};
			var v95: map<array<int>, C0> := map[v4 := v92];
			v91, globalState.f0, v13[596], v92, v13[262] := map v93 : int | v93 in v94 :: (v93 % v2) := (v7), fm7(v92.f22, v2, globalState) + ("oqkka")[|v6| := v0[v2]], v5, v92, map[v4 := v92] == v95;
			v4[698] := v2;
			var v96: set<int> := {v2};
			v4[698] := fm1(-810, (seq(-0x120, i8  => (v2))) <= v3, v96, globalState);
		}
		
		v2 := v2;
	} else {
		v7, v5, globalState.f20 := v7, v5, v2 % v2;
		v5 := v5;
		v2 := v2;
		v2 := v2;
		if ((if (v5) then --v2 else v2) < |v0|) {
			v7 := v7;
			var v97: map<int, int> := map[v2 % v2 := -0xb2 / v2];
			v97 := v97;
			var v98 := DC3(v4, -(v2 * v2));
			var v99: array<array<D0>> := new array<D0>[21];
			var v100: set<int> := {v2, v2};
			v4[699] := fm1(v2, v5, v100, globalState) - v2;
			var v101: map<seq<int>, int> := map[v3 := v2];
			v98, v99, v4[699], globalState.f20, v2 := v98, if (v5 <== !v5) then v99 else v99, fm1(|v101| + v2, v0 != v0, v100, globalState), v2 * v2, v2 / v2;
			var v102 := new C0(true);
			var v103: map<bool, seq<bool>> := map[true := v8];
			var v104 := DC8([v5, v5]);
			var v105: array<seq<bool>> := new seq<bool>[27] [v8, [v5], v8 + v8, v8, v8, v8, fm5(globalState) + v8, v8, v8, [v5], v8, [true], v8, v8[|"fllbype"| := v5], [v5, v102.f22, !!!v5, v5, v5], v8, [v5], if (v102.f22 in v103) then v103[v102.f22] else [v5], [!v5], v104.cf13, v8, v8, v8, v8 + v8, v8, v8 + [v5, false, !v5, v5], v8];
			v105[828] := fm5(globalState) + v8;
			v105[828] := fm5(globalState) + (v8 + v8);
		} else {
			v4[557] := |v8|;
			var v106: array<string> := new string[17](i9 => v0);
			v106[696] := fm7(v5, v2, globalState);
			var v107: set<map<bool, int>> := {map[v5 := v2]};
			v0, v4[557], v106[696], v4, globalState.f20 := v0, (|v107| % v2) / |v0|, seq(0x2d3, i10  => (v7)), v4, v2;
			var v108 := m0(globalState);
			v13[127] := v5;
			v13[127] := v5;
			var v109: multiset<int> := multiset{v2};
			v109 := v109;
			var v110: set<D0> := {v59};
			var v111: map<bool, bool> := map[v59.(cf4 := v13[127]) !in v110 := !!v13[127] ==> !false];
			v111 := v111[false := fm0(-v2, !v13[127], fm0(v2, v5, v5, "t", globalState), v106[696], globalState)];
		}
		
	}
	
	var v112 := m0(globalState);
	var v113: map<int, seq<int>> := map[v2 := v3];
	var v114: map<int, int> := map[v2 := v2];
	v113, globalState.f20, v5 := v113[v2 := [if (v2 in v114) then v114[v2] else v2]], v2, (v0 + v0) !in v12;
	forall i11 | 0 <= i11 < v4.Length {
		v4[i11] := i11 * v2;
	}
	var i12 := 0;
	while (v5)
		decreases 100 - i12
	{
		if (i12 >= 100) {
			break;
		}
		
		i12 := i12 + 1;
		var v115 := new C0(v5);
		globalState.f20 := 0x264;
		globalState.f20 := v2;
		if (v115.f22) {
			var v116: seq<C0> := [v115, v115];
			var v117: set<int> := {|([v115] + v116)|, 0x49, v2};
			v117 := v117;
			globalState.f20 := v2 - (v2 / v2);
			var v118: multiset<bool> := multiset{true};
			var v119: map<int, multiset<bool>> := map[v2 := v118];
			v119 := if (v2 >= 0x32) then v119 else v119;
			var v120 := new C0(true);
			var v121 := m0(globalState);
		} else {
			var v122: array<string> := new string[12](i13 => v0);
			var v123 := DC1(v4, !false, v122);
			v5, globalState.f20, v13, globalState.f20 := (v115.f22 && v115.f22) || v123.cf2, v2, v13, v2;
			globalState.f20 := v2;
			var v125: seq<map<int, int>> := [(map v124 : int | (0x10a <= v124) && (v124 < 0x11) :: (v124 - |v0|) := (v2)) + v114, v114 + map[v2 := |{v2}|], v114];
			var v126: map<D0, seq<map<int, int>>> := map[DC2(v115.f22) := v125];
			v125 := if (DC2(v115.f22) in v126) then v126[DC2(v115.f22)] else v125;
			var v127: map<int, array<bool>> := map[v2 := v13];
			v127 := v127;
			var v128 := DC0(v3);
			var v129 := DC4(v128);
			var v130: array<D0> := new D0[17] [v129, v129, v129, DC4(v128), v129, v129, v129, v129, v129, v129, v129.(cf7 := DC2(v115.f22)), v129.(cf7 := v128), v129, v129, v129, DC4(v128), v129];
			v130[393] := v129;
			v130[393] := if (v5) then if (v115.f22) then v129 else v129 else v129;
		}
		
	}
}