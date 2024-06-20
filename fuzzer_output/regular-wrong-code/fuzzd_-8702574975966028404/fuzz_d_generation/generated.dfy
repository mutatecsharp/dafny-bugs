datatype D0 = DC0(cf0: array<int>, cf1: int) | DC1(cf2: bool, cf3: int, cf4: int, cf5: int) | DC2(cf6: D0)
datatype D1 = DC4(cf8: bool, cf9: int) | DC3(cf7: map<int, set<C0>>)
datatype D2 = DC6(cf11: int, cf12: int, cf13: int) | DC5(cf10: multiset<int>)
datatype D3 = DC8 | DC7(cf14: multiset<array<C0>>)
datatype D4 = DC10(cf16: D3, cf17: int, cf18: char) | DC11(cf19: int, cf20: bool) | DC9(cf15: set<seq<bool>>)
datatype D5 = DC12(cf21: array<bool>)
datatype D6 = DC14(cf23: char, cf24: int, cf25: map<int, bool>, cf26: int) | DC13(cf22: seq<multiset<int>>) | DC15(cf27: D6)
class GlobalState {
	const f0 : bool
	const f1 : int
	const f2 : int
	const f3 : array<bool>
	var f4 : int
	const f5 : char
	const f6 : bool
	constructor (f0 : bool, f1 : int, f2 : int, f3 : array<bool>, f4 : int, f5 : char, f6 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
	}
	
}

function fm2(globalState: GlobalState): bool {
	((set v0 : char | v0 in ['l'] :: (v0)) - {'q', 'i'}) >= ({'h'} * {'n', 'i'})
}
function fm3(globalState: GlobalState): multiset<bool> {
	multiset{{DC4(true, |map[128 := 6]|), DC4(false, 0xe), DC4(false, -0x2b2), DC4(true, |multiset([!true])|), DC4(true, |(seq(321, i0  => ('b')))|)} <= {DC4(true, 997), DC4(false, |[false]|)}}
}
function fm4(p0: int, globalState: GlobalState): multiset<int> {
	multiset{0xc3}
}
function fm5(p0: int, globalState: GlobalState): map<set<bool>, bool> {
	((map v0 : set<bool> | v0 in multiset{{false}} :: (v0) := (true)) + map[{false, true, false} := false]) + map[{false, true, true} := false]
}
function fm6(p0: bool, p1: char, p2: D1, p3: bool, globalState: GlobalState): string {
	"wmqadti"
}
function fm7(p0: bool, p1: bool, p2: int, globalState: GlobalState): char {
	if (true ==> false) then 'p' else if (!!!false) then 'l' else 'b'
}
function fm8(p0: bool, p1: bool, globalState: GlobalState): map<bool, int> {
	match DC9({[false]}) {
		case DC10(cf16, cf17, cf18) => map[!!!true := |map[cf17 := !!true]|] + map[true := cf17]
		case DC11(cf19, cf20) => map[cf20 := cf19]
		case DC9(cf15) => map[false := |[|multiset{true}|]|]
	}
}
function fm9(p0: int, p1: bool, globalState: GlobalState): D2 {
	DC6(|(map[true := 0x31] + map[true := -0x3b6])|, 132 * |multiset{true, true}|, 811 - |multiset{true}|)
}
function fm10(p0: int, p1: int, globalState: GlobalState): seq<seq<bool>> {
	((seq(0x2de, i0  => ([true]))) + (seq(0x399, i1  => ([true])))) + ((seq(0x200, i2  => ([!!true]))) + [[false], [true, !true]])
}
function fm11(p0: bool, globalState: GlobalState): map<bool, bool> {
	map[false := true] + (map[!false := !true] + map[true := true])
}
function fm12(p0: int, p1: int, globalState: GlobalState): seq<multiset<int>> {
	[multiset{761}] + ([multiset([|[|"ttdvaiasm"|, |[false, false, false]|, -0x33e, |[map[|{true, true}| := 817]]|]|, 0x4b])] + DC13([multiset([0x1ca, -|multiset{true}|]), multiset{0x382}, multiset{-0x2b3, |[|map[false := true]|]|}, multiset{0x230, |{false}|}]).cf22)
}
function fm13(p0: int, p1: bool, p2: bool, p3: bool, globalState: GlobalState): set<int> {
	{|multiset{false}| * 0x1b, 0xab / 0x185}
}
method m0(p0: int, globalState: GlobalState) returns (r0: int) {
	r0 := p0;
	var v0 := true;
	v0 := v0;
	var v1: C0 := new C0();
	v1 := v1;
	var v2: array<int> := new int[27];
	var v3 := "eo";
	globalState.f4, v2, globalState.f4, r0 := p0, v2, p0, if (v0) then p0 else |v3|;
	if (v3 < v3) {
		v1 := v1;
		globalState.f4 := p0;
		var v4: array<bool> := new bool[16];
		v4[938] := v0;
		v4[938] := if (v0) then v0 else fm2(globalState);
		var v5: multiset<int> := multiset{p0, p0};
		v4[938], v0, globalState.f4, v0 := (p0 / p0) != p0, !!(v5 >= v5), p0, |multiset{v4[938], v4[938]}| >= p0;
		var v6: array<char> := new char[2](i0 => 'r');
		var v7 := 'h';
		v6[694] := v7;
		v6[694] := v7;
	} else {
		var v8: set<bool> := {v0};
		var v9: map<int, set<bool>> := map[-p0 := v8];
		globalState.f4 := -(-(|v9| - p0) - (p0 - p0));
		var v10: map<int, int> := map[p0 := v1.fm1(false, 0x25d, p0, globalState)];
		var v11: seq<int> := [|v10|, p0, p0];
		var v12: seq<int> := [|v11|, p0, p0, p0, p0];
		var v13: multiset<int> := multiset{|fm13(v11[p0], v0, v0, v0, globalState)|};
		var v14: multiset<bool> := multiset{v0, v0, true};
		var v15: map<bool, bool> := map[true := v0];
		r0 := if ((|v14| % p0) in v13) then v13[|v14| % p0] else if (p0 in v10) then v10[p0] else |v15|;
		var v16 := new C0();
		var v17: array<bool> := new bool[27];
		v17[866] := false;
		v17[866] := v0;
		globalState.f4 := p0;
	}
	
	for i1 := p0 to p0 * p0 {
		var v18 := DC4(v0, i1);
		if (if (v18.cf8) then v0 else v0) {
			v3 := v3;
			var v19: array<bool> := new bool[6];
			v19[892] := v0;
			var v20: map<array<int>, string> := map[v2 := v3];
			v19[892] := v2 !in v20;
			var v21: array<array<int>> := new array<int>[10] [v2, v2, v2, v2, v2, v2, v2, v2, v2, v2];
			v21[388] := v2;
			v21[388], v19[892] := v2, v0;
			var v22: map<int, int> := map[p0 := i1];
			var v23: map<C0, map<int, int>> := map[v1 := v22 + v22];
			var v24: map<bool, map<C0, map<int, int>>> := map[!fm2(globalState) := map[v1 := v22]];
			v23 := v23 + (if (v0 in v24) then v24[v0] else map[v1 := v22]);
			var v25: map<bool, C0> := map[v0 := v1];
			var v26 := 'p';
			v19[892] := (0x32f / |v25|) <= (|v3[453 := v26]| * p0);
		} else {
			var v27: array<bool> := new bool[14](i2 => |multiset([v0])| < p0);
			var v28 := DC12(v27);
			v27 := v28.cf21;
			globalState.f4 := 617;
			globalState.f4 := -p0;
			var v29 := 'q';
			v29 := v29;
			var v30: map<seq<char>, D0> := map[v3 := DC0(v2, p0)];
			var v31: multiset<bool> := multiset{v0, v0};
			var v32: multiset<int> := multiset{v1.fm1(true, p0, |map[false := |v30|]|, globalState), p0, p0, |v31|, p0};
			var v33: array<multiset<int>> := new multiset<int>[3] [v32, v32, v32];
			v33 := new multiset<int>[28](i3 => multiset{418});
		}
		
		var v34: map<int, int> := map[p0 := p0];
		var v35: seq<map<int, int>> := [v34 + v34, v34, v34, v34, v34];
		v1, v34, globalState.f4 := v1, v35[p0], 0x26a + i1;
		var v36 := new C0();
		var v37 := new C0();
	}
	r0 := if (false && v0) then |v3| else p0 % p0;
}
class C0 {
	constructor () {
	}
	
	function fm0(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState): seq<bool> {
		[true] + [true, true, true]
	}
	function fm1(p0: bool, p1: int, p2: int, globalState: GlobalState): int {
		-214
	}
}

method Main() {
	var v0: array<bool> := new bool[22];
	var globalState := new GlobalState(true, -0x209, 537, v0, 0x2a9, 'e', false);
	var v1 := 0xb0;
	var v2 := true;
	for i0 := v1 to if (v2) then v1 else -845 {
		globalState.f4 := i0;
		var v3 := "mjypph";
		v3 := (seq(163, i1  => ('k'))) + v3;
		v1 := 0x3e7;
		var v4 := new C0();
	}
	v2 := v2;
	var v5: C0 := new C0();
	var v6: map<C0, int> := map[v5 := v1];
	var v7: map<bool, bool> := map[v1 <= |v6| := v2];
	var v8: map<int, int> := map[v1 := 0xbc];
	v7 := v7[v2 := |v8| in v8];
	var v9: seq<array<bool>> := [v0];
	globalState.f4 := |(v9 + [v0, v0, v0])|;
	var i2 := 0;
	while (v2 || fm2(globalState))
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		var v10: array<int> := new int[5];
		var v11: seq<int> := [0x41];
		var v12 := 'w';
		var v13: map<char, int> := map[v12 := v1];
		var v14: seq<int> := [v11[if (v12 in v13) then v13[v12] else v1]];
		v10[128] := -v14[v1];
		var v15: set<int> := {-v1, v1, v1};
		v10[128] := |(v15 - (v15 * v15))|;
		var v16 := m0(v10[128] / -|v14|, globalState);
		var v18: array<map<map<bool, bool>, bool>> := new map<map<bool, bool>, bool>[22](i3 => (map v17 : map<bool, bool> | v17 in map[v7 := |v15|] :: (v17) := (v2)) + map[v7[v2 := v2] := false]);
		v18[802] := map[v7 := v2] + map[v7 := v2];
		var v19: array<char> := new char[3](i4 => v12);
		var v20: array<array<char>> := new array<char>[3] [v19, v19, v19];
		var v21: seq<array<array<char>>> := [v20];
		var v22: array<array<array<char>>> := new array<array<char>>[6] [v21[v16], v20, v20, v20, v20, v20];
		v22[420] := v20;
		v5, v18[802], v2, v22[420] := v5, map[v7 := v2], v2, v20;
		v2 := fm2(globalState);
	}
	v2 := v2 && v2;
	var v23 := "grpdnh";
	globalState.f4 := v1 + -|v23|;
	var v24 := DC1(v2, v1, v1, v1);
	var v25: map<bool, D0> := map[!!v2 := v24];
	var v26: multiset<map<bool, D0>> := multiset{(map[v2 := v24])[v2 := v24], v25};
	var v27 := DC1(v2, v1, v1, if (v25 in v26) then v26[v25] else v5.fm1(!v2, v1, v1, globalState));
	match (v27) {
		case DC0(cf0, cf1) =>
			var v28 := new C0();
			var v29 := 'm';
			cf0[418] := v1;
			var v30: map<char, bool> := map[v29 := v2];
			var v31: seq<multiset<int>> := [multiset{|v30|}];
			var v32: multiset<int> := multiset{v1};
			var v33: seq<bool> := [v2];
			var v34: map<int, seq<bool>> := map[v1 := v33];
			var v35: array<seq<multiset<int>>> := new seq<multiset<int>>[16] [v31, v31 + ([v32, v32])[v1 := v32], v31, v31, v31, [v32], v31, v31, v31[|(if (v1 in v34) then v34[v1] else v33)| := v32], v31, v31 + v31, v31, v31, [v32, v31[v1], v32, v32], v31, v31 + v31];
			v35[549] := v31 + v31;
			var v36: map<array<int>, int> := map[cf0 := -v1];
			var v37: map<bool, string> := map[false := seq(-0xe5, i5  => (v29))];
			v29, cf0[418], v35[549], v28, globalState.f4 := v29, if (cf0 in v36) then v36[cf0] else cf1 * cf1, v31[cf1 := v32] + (v31 + [v32, v32]), v28, |v37[v2 := seq(0x2a2, i6  => (v29))]| / (v1 / v28.fm1(true, cf1, cf1, globalState));
			v29 := v29;
			var v38: set<bool> := {v2};
			var v39: map<multiset<int>, int> := map[v32 * v32 := |v38|];
			cf1 := if (v32 in v39) then v39[v32] else |v23|;
		case DC1(cf2, cf3, cf4, cf5) =>
			if (v2) {
				v7 := v7;
				var v40: seq<int> := [0xe3];
				v1 := if (v2) then 0x26f else v40[v1] * -cf4;
				var v41: seq<C0> := [v5, v5];
				var v42: seq<seq<C0>> := [v41 + [v5]];
				v42 := v42;
				v8 := (v8 + v8) + map[cf3 := cf3];
				var v43: map<int, set<C0>> := map[v1 := {v5}];
				var v44 := DC3(v43);
				v43 := v44.cf7;
			} else {
				v5 := v5;
				var v45: array<seq<bool>> := new seq<bool>[8];
				var v46: seq<bool> := [v2];
				var v47: seq<seq<bool>> := [v46, [cf2], v5.fm0(cf5, |v46|, cf4, v2, globalState)];
				v45[387] := v47[|v23|];
				v45[387] := v46;
				var v48: array<array<bool>> := new array<bool>[10] [v0, v0, v0, v0, v0, v0, v0, v0, v0, v0];
				v48[386] := v0;
				var v49: array<int> := new int[3];
				var v50: map<array<int>, bool> := map[v49 := cf2];
				v48[386] := if (v50 != v50) then v0 else v0;
				var v51: array<char> := new char[14];
				var v52 := 'u';
				v51[16] := v52;
				v51[16] := v52;
				v49[3] := cf5;
				var v53: multiset<bool> := multiset{cf2, cf2};
				v1, v5, v49, v49[3], v0 := if ((cf3 % cf4) in v8) then v8[cf3 % cf4] else v1, v5, if (v2) then v49 else v49, |((fm3(globalState) * v53) + v53[v2 := cf4])|, if (if (v2) then cf2 else v2) then v48[386] else v48[386];
			}
			
			cf2 := v2;
			if ((cf3 + v1) > 290) {
				cf2 := cf2;
				var v54 := m0(-cf5, globalState);
				var v55: array<C0> := new C0[13] [v5, v5, v5, v5, v5, v5, v5, v5, v5, v5, v5, if (v2) then v5 else v5, v5];
				v55[84] := v5;
				v55[84] := v5;
				var v56: array<D1> := new D1[15];
				var v57 := DC4(v2, -cf4);
				v56[838] := v57;
				v56[838] := v57;
				var v58 := m0(cf3, globalState);
			} else {
				var v59: set<C0> := {v5, v5};
				var v60: map<int, set<C0>> := map[cf3 := v59];
				var v61 := DC3(v60);
				v61 := v61;
				v2 := cf2;
				v5 := v5;
				var v62: array<int> := new int[6];
				var v63: map<bool, array<int>> := map[v2 := v62];
				var v64: map<bool, map<bool, array<int>>> := map[v2 <==> v2 := v63];
				var v65: seq<bool> := [cf2];
				var v66: map<int, seq<bool>> := map[cf3 := v65];
				var v67: seq<int> := [cf4];
				v64 := v64[(if (|v67| in v66) then v66[|v67|] else v65) <= [!v2] := v63];
				v65 := v65 + [!true, cf2, v2, !true, cf2];
			}
			
			var v68 := new C0();
		case DC2(cf6) =>
			var v69: map<bool, C0> := map[v2 := v5];
			var v70: map<map<bool, C0>, array<bool>> := map[map[v2 := v5] + v69 := v0];
			v0 := if (v69[v2 := v5] in v70) then v70[v69[v2 := v5]] else if (v2) then v0 else v0;
			if (fm2(globalState)) {
				v1 := v1 + v1;
				globalState.f4 := |v23|;
				var v71: array<int> := new int[25];
				v71[370] := v1 % v1;
				v71[370] := v1;
				v5 := new C0();
				var v72: multiset<int> := multiset{0x27, |v23| + v1};
				v72 := v72 * fm4(0x36d, globalState);
			} else {
				var v73: map<int, bool> := map[v1 := !v2];
				var v74: seq<map<int, bool>> := [v73, v73[|map[v1 := |v7|]| := v2], v73, v73, map[v1 := true]];
				var v75: map<seq<map<int, bool>>, int> := map[v74 := v1];
				v75 := v75[v74 := -41];
				var v76 := 'n';
				v76 := v76;
				var v77: array<string> := new string[26](i7 => v23);
				var v78: array<int> := new int[7];
				v77, v78 := v77, v78;
				globalState.f4 := 0x17e;
				var v79: map<int, set<C0>> := map[0x12f := {v5}];
				var v80 := DC3(v79);
				var v81 := m0(0x2f2 + |multiset{v80}|, globalState);
			}
			
			var v82: multiset<int> := multiset{v1};
			v82 := v82;
			var v83: array<int> := new int[4] [v1, v1, v1, v1];
			var v84 := DC0(v83, v1);
			v84 := v84;
	}
	
	v0[999] := v2;
	var v85: seq<bool> := [fm2(globalState), fm2(globalState)];
	v0[999] := v85[v1];
	if (v2) {
		var v86 := new C0();
		v85 := if (v1 >= v1) then v85 else v85;
		globalState.f4 := 0x91;
		v5 := v86;
		var v87: seq<int> := [v1, if (v2) then v1 else v1, v1 + 0x3e3];
		v87 := v87;
	} else {
		var v88: map<string, int> := map[v23 := v1];
		v88 := v88[v23 := v1];
		var v89: seq<int> := [v1];
		v0[999] := (v89 != v89) <==> v2;
		var v90: array<int> := new int[5](i8 => i8 / -0x22a);
		var v91: seq<array<int>> := [v90, v90];
		var v92: multiset<int> := multiset{v1};
		var v93: map<int, C0> := map[v1 := v5];
		var v94 := DC0(v91[0x372], if (|v93| in v92) then v92[|v93|] else v1);
		var v95 := DC2(v94);
		match (v95) {
			case DC0(cf0, cf1) =>
				var v96 := 'u';
				v96 := v96;
				v8 := v8[207 := 274];
				cf0[62] := |v23[|fm5(cf1, globalState)| := v96]|;
				cf0[62] := v1;
				v0 := v0;
			case DC1(cf2, cf3, cf4, cf5) =>
				v1 := v1;
				globalState.f4 := 0x140;
				cf2 := cf2;
				var v97: seq<C0> := [v5];
				var v98: multiset<map<bool, bool>> := multiset{map[!false := v0[999]]};
				v5 := if (true) then v5 else v97[|v98|];
			case DC2(cf6) =>
				var v99: array<C0> := new C0[5] [v5, v5, v5, v5, v5];
				v99[987] := v5;
				var v100: map<int, string> := map[-0x1cd := v23];
				v0[999], v99[987] := ((seq(0x3cf, i9  => ('r'))) + "rrodo") <= (if (v1 in v100) then v100[v1] else v23), v5;
				globalState.f4, globalState.f4, globalState.f4, v2 := v1 + -v1, v1, v1, v2;
				var v101 := new C0();
				globalState.f4 := v1;
		}
		
		var v102 := DC0(v90, v1);
		match (v102) {
			case DC0(cf0, cf1) =>
				var v103: multiset<bool> := multiset{v2};
				v0[999] := v103 <= v103;
				var v104 := DC4(v0[999], |v85|);
				globalState.f4 := v104.cf9;
				var v105: set<C0> := {v5};
				v2 := v105 < (v105 + v105);
				var v106 := new C0();
			case DC1(cf2, cf3, cf4, cf5) =>
				var v107 := new C0();
				var v108 := 't';
				var v109 := DC4(true, 0x163);
				var v110 := m0(-|fm6(false, v108, v109, cf2, globalState)| * -cf5, globalState);
				var v111: array<map<C0, int>> := new map<C0, int>[20];
				v111, v0[999] := v111, v24.cf2;
				globalState.f4, cf2, v5 := -cf5, v2, v107;
			case DC2(cf6) =>
				var v112 := new C0();
				var v113 := DC4(v2, v1);
				var v114 := 'y';
				var v115: multiset<bool> := multiset{v2};
				var v116: map<char, multiset<bool>> := map[v114 := v115];
				var v117: map<C0, string> := map[v112 := v23];
				var v118: map<char, array<int>> := map[fm7(v2, false, |(if (v5 in v117) then v117[v5] else v23)[v1 := v114]|, globalState) := v90];
				var v119: array<array<int>> := new array<int>[14] [v90, if (v113.cf8) then v90 else v90, v90, v91[|(if (v114 in v116) then v116[v114] else v115)|], v90, v90, v90, v90, v90, if (v114 in v118) then v118[v114] else v90, v90, v90, v90, v90];
				v119[319] := v90;
				v119[319] := v90;
				var v120 := m0(v1, globalState);
				var v121 := new C0();
		}
		
		v89 := (v89 + v89)[v1 - v1 := v5.fm1(v0[999], v1, v1, globalState)];
	}
	
	var v122: seq<seq<bool>> := [v85];
	var v123: seq<int> := [v1];
	var v124: array<int> := new int[14] [v1, v1, v1 % v1, v1, v1 + v1, v1 / v1, v5.fm1(fm2(globalState), v1, |"lyoe"|, globalState), v1, v1, |fm8(v2, v2, globalState)|, |(v85 + v122[|v123|])|, v1 + v1, -0x1e0 - v1, v1];
	var v125: set<int> := {0x269};
	v124 := if (v125 >= v125) then v124 else if (v2) then v124 else v124;
	var v128: array<map<int, int>> := new map<int, int>[8] [v8, map v126 : int | v126 in (map v127 : int | v127 in v123 :: (v127 / |{v1, v1}|) := (DC5(multiset{v1}).cf10)) :: (v126 / |v8|) := (|v85|), v8, v8, v8, v8, map[67 := v1], v8];
	v128 := v128;
	v125 := v125;
	if (v2) {
		var v129 := DC6(v1, v1, v1);
		var v130: set<C0> := {v5, v5};
		var v131: multiset<int> := multiset{v1};
		v2, v129, v2 := v2 && v0[999], fm9(v1, v0[999], globalState).(cf12 := v1, cf11 := if (false) then |v130| else 0x108), v131 > multiset{v1};
		var v132 := m0(v1, globalState);
		if (if (fm2(globalState) || v0[999]) then v0[999] else v0[999]) {
			var v133: array<C0> := new C0[10];
			var v134: multiset<array<C0>> := multiset{v133};
			var v135 := DC7(v134);
			v134 := (v134 - v134) * v135.cf14;
			globalState.f4 := DC0(v124, |v85|).cf1;
			v2 := v2;
			v124[257] := if (v5.fm1(v0[999], 0x2d0, v1, globalState) in v131) then v131[v5.fm1(v0[999], 0x2d0, v1, globalState)] else v132;
			v124[257] := v1 + -|v85|;
			v23 := v23;
		} else {
			v0[999] := v2;
			var v137 := 'u';
			var v138: multiset<char> := multiset{v137};
			v2 := DC1(v0[999], v1, |v23|, |(map v136 : char | v136 in v138 :: (v136) := (v0[999]))|).cf2;
			var v139: set<bool> := {v0[999], v0[999]};
			var v140: map<set<bool>, C0> := map[v139 := v5];
			var v141: multiset<C0> := multiset{v5, if (v139 in v140) then v140[v139] else v5, v5};
			var v142: map<bool, int> := map[fm2(globalState) := |v141|];
			v2 := |(if (v2) then v142 else v142)| == v132;
			var v143: array<string> := new string[5](i10 => v23 + v23);
			v143 := if (v2) then v143 else v143;
			v2 := v0[999] ==> (false <==> v85[v1]);
		}
		
		v1 := v1;
		v5 := v5;
	} else {
		globalState.f4 := v1;
		var v144: array<string> := new string[21];
		v144[487] := v23;
		v144[487] := "aqmgw";
		var v145: multiset<C0> := multiset{v5, v5};
		var v146 := DC1(v0[999], v1, 188, |v145|);
		var v147 := DC2(if ((v24.(cf2 := true, cf5 := v1)).cf2) then v146 else v146);
		match (v147) {
			case DC0(cf0, cf1) =>
				globalState.f4 := -cf1;
				var v148 := m0(cf1, globalState);
				var v149 := m0(v148, globalState);
				var v150: array<C0> := new C0[20];
				v150 := v150;
			case DC1(cf2, cf3, cf4, cf5) =>
				var v151 := new C0();
				globalState.f4 := -cf5;
				v9, v144[487], v5, v0, v144[487] := (v9 + v9) + v9, "bsuwye", v5, v0, v23;
				var v152: map<int, bool> := map[cf5 := cf2];
				v152 := map v153 : int | (0x3 <= v153) && (v153 < 0x2b6) :: (v153 * cf5) := (v0[999]);
			case DC2(cf6) =>
				var v155: seq<map<int, bool>> := [map[v1 := false], map[v1 := v2] + (map v154 : int | (-0x2ea <= v154) && (v154 < 0x31c) :: (v154 % v1) := (v2)), map[v1 := !v2]];
				v155 := v155;
				var v156 := new C0();
				v0[999] := v2;
				globalState.f4 := v156.fm1(v2, 0xb6, v1, globalState);
		}
		
		v0[999], v2 := v24.cf2, v2;
		v23, globalState.f4 := v23, v1 + v1;
	}
	
	var v157: set<seq<bool>> := {v85};
	var v158: seq<seq<seq<bool>>> := [fm10(v1, v1, globalState), v122];
	var v160: array<set<seq<bool>>> := new set<seq<bool>>[11] [v157, v157 - {v85[v1 := v0[999]], v85, v85}, v157, v157 + v157, v157, set v159 : seq<bool> | v159 in v158[|v23|] :: (v159), {v85}, v157, {v85} * v157, v157, v157];
	v160[592] := v157;
	var v161: array<map<map<bool, bool>, bool>> := new map<map<bool, bool>, bool>[22];
	var v162: map<map<bool, bool>, bool> := map[map[false := v2] := false];
	v161[867] := map[map[v2 := v2] := v0[999]] + v162;
	var v163 := DC9({v85});
	var v165 := DC6(v1, v1, v1);
	var v166: map<map<bool, bool>, D2> := map[fm11(v2, globalState) := v165];
	v123, v160[592], v2, v161[867] := if (v2) then v123 else v123, (v157 - v163.cf15) + (v157 * v157), v0[999], map[v7 := v2] + ((map v164 : map<bool, bool> | v164 in v166 :: (v164) := (v0[999])) + v162);
	if (v2) {
		v124[836] := v1;
		v124[836] := v1;
		var v167 := new C0();
		v124[836] := v124[836] / (if (v0[999]) then v124[836] else v124[836]);
		globalState.f4 := v124[836];
		var v168: multiset<int> := multiset{v124[836]};
		var v169: seq<multiset<int>> := [multiset{v5.fm1(v0[999], v1, v1, globalState)}, v168];
		v0[999] := !(fm12(DC11(v124[836], v0[999]).cf19, |[true]|, globalState) <= v169);
	} else {
		var v170: array<string> := new string[25](i11 => v23[v1 := 'a']);
		var v171: set<bool> := {fm2(globalState), v0[999], v0[999], v2};
		v0[999], v170 := v171 !! (v171 - v171), if (v0[999]) then v170 else v170;
		var v172 := m0(v1, globalState);
		v5 := new C0();
		v0, v172 := v0, v1;
		var v173: multiset<bool> := multiset{v0[999], true};
		if (v173 != multiset(v85)) {
			v23 := v23 + (seq(0x143, i12  => ('m')));
			var v174 := new C0();
			v1, globalState.f4, globalState.f4 := -685, v172, v172;
			var v175: map<int, bool> := map[|v85| := !v2];
			var v176: map<map<int, bool>, int> := map[v175 := |([v0, v0])[v1 := v0]|];
			v2 := true <==> (v175 !in v176);
			var v177: seq<seq<int>> := [v123];
			v177 := v177 + v177;
		} else {
			var v178: multiset<array<bool>> := multiset{v0, v0, v0, v0, v0};
			v0[999] := v178 !! v178;
			var v179: multiset<seq<bool>> := multiset{v85, v85, v85};
			v0[999] := (v179 + v179) > (v179 - v179);
			v124[846] := v172;
			v124[846] := v5.fm1(v1 >= v172, v1, v123[v172], globalState);
			var v180 := m0(v1, globalState);
			var v181 := DC1(false, v1, v172, v1);
			var v182 := DC2(v181);
			var v183: seq<D0> := [v181];
			var v184 := DC2(v183[v180]);
			var v185 := DC2(v184);
			v185 := v185;
		}
		
	}
	
}