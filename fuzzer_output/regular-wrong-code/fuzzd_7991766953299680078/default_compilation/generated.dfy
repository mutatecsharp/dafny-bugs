datatype D0 = DC0(cf0: set<int>, cf1: string) | DC1(cf2: bool, cf3: map<bool, bool>, cf4: C0, cf5: int) | DC2(cf6: int, cf7: multiset<int>, cf8: bool) | DC3(cf9: D0)
datatype D1 = DC5(cf11: int) | DC6(cf12: set<bool>) | DC7(cf13: int, cf14: int) | DC4(cf10: array<int>) | DC8(cf15: D1)
datatype D2 = DC10 | DC11(cf17: char, cf18: int) | DC9(cf16: map<int, int>)
datatype D3 = DC13(cf20: bool, cf21: bool, cf22: D1) | DC12(cf19: array<bool>) | DC14(cf23: D3)
datatype D4 = DC15(cf24: multiset<bool>)
datatype D5 = DC17(cf26: D2) | DC18(cf27: D1, cf28: int, cf29: char, cf30: multiset<int>, cf31: int) | DC16(cf25: map<seq<int>, bool>) | DC19(cf32: D5)
class GlobalState {
	var f0 : int
	const f1 : int
	const f2 : int
	const f3 : bool
	var f4 : bool
	const f5 : int
	var f6 : seq<bool>
	const f7 : int
	const f8 : array<int>
	const f9 : int
	var f10 : int
	var f11 : map<bool, int>
	const f12 : bool
	var f13 : array<int>
	const f14 : array<bool>
	const f15 : multiset<array<bool>>
	var f16 : int
	const f17 : char
	constructor (f0 : int, f1 : int, f2 : int, f3 : bool, f4 : bool, f5 : int, f6 : seq<bool>, f7 : int, f8 : array<int>, f9 : int, f10 : int, f11 : map<bool, int>, f12 : bool, f13 : array<int>, f14 : array<bool>, f15 : multiset<array<bool>>, f16 : int, f17 : char) {
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
	}
	
}

function fm0(p0: int, p1: bool, globalState: GlobalState): int {
	|map[--836 := 'b']| / |("vmxs" + "nrusd")|
}
function fm1(p0: char, globalState: GlobalState): map<int, multiset<int>> {
	(map[911 := multiset{-0x263, 0x2}] + map[DC2(0x30c, multiset{0xa0}, true).cf6 := multiset{253}]) + map[0xeb := multiset([|map[|map[true := -|[true]|]| := 'b']|])]
}
function fm2(p0: bool, p1: bool, globalState: GlobalState): string {
	"bgrolje"
}
function fm3(p0: bool, p1: int, p2: int, p3: bool, globalState: GlobalState): bool {
	match DC10() {
		case DC10() => !false
		case DC11(cf17, cf18) => [0x1e] < [-0x1bf]
		case DC9(cf16) => !DC2(|cf16|, multiset{-769}, true).cf8 <==> false
	}
}
function fm5(globalState: GlobalState): seq<bool> {
	(if (true) then [true] else [true]) + ([true] + [false])
}
function fm6(p0: int, p1: string, p2: bool, globalState: GlobalState): set<bool> {
	(if (false) then {false, true} else {!false, false, false}) * (if (false) then {false, true} else {false})
}
function fm7(globalState: GlobalState): char {
	'h'
}
function fm8(p0: int, p1: D2, p2: bool, p3: string, globalState: GlobalState): multiset<bool> {
	multiset{true, 0x7 < 211, if (true) then false else true}
}
function fm9(p0: bool, globalState: GlobalState): map<int, int> {
	match DC6({false}) {
		case DC5(cf11) => map[cf11 := cf11] + map[|"dtisl"| := cf11]
		case DC6(cf12) => map[|"wf"| := 589]
		case DC7(cf13, cf14) => map[cf14 := cf13] + map[cf13 := |multiset([false, false])|]
		case DC4(cf10) => map[|DC16(map[[0x106] := false]).cf25| := |[|"bv"|, 0x209]|] + map[|[|map[multiset{-0x3cc} := |map[-630 := 0x22e]|]|, 0x4e]| := 0x2aa]
		case DC8(cf15) => map[|map[!!true := false]| := 0x250]
	}
}
function fm10(p0: bool, p1: int, globalState: GlobalState): D1 {
	if ("ttrqiyhpy" != "twandbl") then DC6({false}) else DC6({false, !true, true})
}
function fm11(p0: map<seq<int>, int>, p1: bool, p2: set<bool>, globalState: GlobalState): multiset<int> {
	multiset{|"mwuiwydjh"|} - (multiset{|"nwoc"|, |{!false}|} + multiset{-|multiset{0x42, 0x25d}|})
}
method m0(p0: string, p1: bool, globalState: GlobalState) returns (r0: int) {
	var v0 := 0x146;
	var v1: seq<int> := [v0, v0];
	r0 := (222 / v0) / (|v1| / v0);
	var v2 := DC10();
	match (v2) {
		case DC10() =>
			var v3: array<D0> := new D0[10];
			var v4: seq<array<D0>> := [v3, v3, v3];
			var v5: array<array<D0>> := new array<D0>[27] [v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v4[v0], v3, v3, v3];
			v5[430] := v3;
			v5[430] := v3;
			var v6: array<set<int>> := new set<int>[7];
			var v7: seq<bool> := [p1, p1, p1];
			var v8: set<int> := {|v7|};
			v6[530] := {v0} + v8;
			var v9 := 'v';
			var v10: set<char> := {v9, v9};
			var v11: map<bool, int> := map[p1 := |v7|];
			var v12: map<bool, map<bool, int>> := map[false := v11];
			var v13: map<char, int> := map[v9 := |v1|];
			var v14: multiset<int> := multiset{|p0|, v0, |{v9}|, |v13|, v0};
			var v15: array<int> := new int[24] [0x38, v0, |("a" + p0[v0 := v9])|, v0, |v7|, -v0, v0 / v0, |v8|, v0 % v0, |p0|, v0, v0, v0, |v7| - v0, v0, |(if (p1) then v10 else {v9})|, v0 * v0, |(if (p1 in v12) then v12[p1] else v11)| + v0, |v14|, -(v0 + v0), v0, v0, -v0, v0];
			v15[838] := v0;
			var v16 := DC0(v8, p0);
			v6[530], globalState.f16, v9, v15[838] := (v16.(cf1 := "dqexxyx")).cf0, -v0, v9, -|v14|;
			globalState.f16 := v0;
			var v17 := DC2(|[v15[838], 906, v0, v15[838], v15[838]]|, v14, p1);
			globalState.f10 := -(|{p1, p1}| - (v0 * v17.cf6));
		case DC11(cf17, cf18) =>
			var v18 := DC5(v0);
			match (v18) {
				case DC5(cf11) =>
					var v19 := new C0(v0, p1);
					var v20: map<bool, bool> := map[p1 := p1];
					var v21 := DC7(v0, |v20|);
					var v22: map<map<int, bool>, bool> := map[(map[cf11 := p1])[|p0| := p1] := fm3(p1, cf18, |p0|, false, globalState)];
					var v23: multiset<map<map<int, bool>, bool>> := multiset{v22 + v22, v22};
					var v24: map<int, bool> := map[cf18 := p1];
					var v25: set<int> := {cf18, |v24|};
					var v26: map<int, bool> := map[|v25| := p1];
					v21, globalState.f10 := v21, if (map[v24 := p1] in v23) then v23[map[v24 := p1]] else v0 * 0x378;
					v0 := cf11 * v0;
					var v27: seq<bool> := [p1, p1, p1];
					var v28: T0 := new C0(cf18, v27 == v27[v0 := p1]);
					v28 := v28;
				case DC6(cf12) =>
					var v29 := DC7(v0 / v0, cf18);
					v29 := v29.(cf14 := -|p0|);
					var v30: multiset<int> := multiset{v0, |cf12|, v0};
					var v31: seq<multiset<int>> := [v30];
					var v32: map<int, bool> := map[cf18 := p1];
					var v33: seq<bool> := [p1];
					var v34: array<bool> := new bool[20] [p1, cf18 != cf18, fm3(p1, v0, cf18, p1, globalState), !p1 && p1, p1, p1, multiset{-v0} > v31[v0], !(if (v0 in v32) then v32[v0] else p1), p1, p1, p1, v33[DC7(v0, -298).cf13], p1, fm3(p1, v0, |v1|, p1, globalState), (if (cf18 in v32) then v32[cf18] else p1) && p1, p1, p1, p0 == p0, !false, p1];
					var v35 := DC12(v34);
					v34 := DC12(v35.cf19).cf19;
					var v36: array<array<bool>> := new array<bool>[10];
					v34[900] := p1;
					var v37: map<bool, bool> := map[p1 := p1];
					var v38: multiset<bool> := multiset{p1, p1};
					var v39: C0 := new C0(|v38|, p1);
					var v40 := DC1(p1, v37, v39, |v33|);
					v36, globalState.f16, v34[900], globalState.f4 := v36, (v40.(cf3 := v37)).cf5, !(v32 == map[cf18 := false]), p1;
					var v41 := new C0(v0, p1 || p1);
				case DC7(cf13, cf14) =>
					globalState.f16 := cf13;
					var v42: C0 := new C0(v0, p1);
					var v43: map<bool, C0> := map[true := v42];
					var v44: array<C0> := new C0[17] [if (p1) then v42 else v42, v42, v42, v42, v42, v42, v42, v42, v42, v42, v42, v42, v42, v42, if (p1 in v43) then v43[p1] else v42, v42, v42];
					v44[716] := v42;
					v44[716] := v42;
					v42 := v42;
					var v45 := DC11(cf17, -0xf3);
					var v46: map<C0, int> := map[v42 := v45.cf18];
					var v47: seq<string> := [p0];
					v46 := v46[v44[716] := |(p0 + v47[cf18])|];
				case DC4(cf10) =>
					globalState.f4 := p1;
					var v48: map<bool, bool> := map[p1 := p1];
					var v49: C0 := new C0(cf18, p1);
					var v50: set<int> := {cf18};
					var v51 := DC1(p1, v48, v49, |v50|);
					var v52: seq<bool> := [true, p1, p1];
					var v53: array<bool> := new bool[27] [p1, p1, if (true) then p1 else p1, p1, true, p1, p1, p1, p1, !!v51.cf2, p1, p0 < p0, v51.cf2, !p1, !p1, p1, v52[v1[cf18]], p1, p1, !(757 == cf18), p1, true, p1, p1, p1 ==> true, p1, p1 && p1];
					v53[369] := p1;
					cf10[455] := cf18;
					var v54: set<bool> := {p1};
					var v55 := DC6(v54);
					var v56 := DC8(v55);
					var v57 := DC14(DC14(DC13(p1, p1, v56)));
					var v58: map<char, D3> := map[cf17 := v57];
					v53[137] := map['g' := v57] != v58;
					var v59 := DC0({cf18}, "xsto");
					var v60: map<D0, bool> := map[v59 := if (p1) then !p1 else !p1];
					var v61 := DC6({p1});
					var v62: multiset<D1> := multiset{v61, v61, v61, v61, v61.(cf12 := {v52[v0], p1})};
					var v63: set<string> := {p0};
					var v64: seq<map<D0, bool>> := [map[v59 := p1]];
					var v66: map<D0, char> := map[v59 := cf17];
					v53[369], cf10[455], v53[137], v2, v60 := v62 != v62, v0 / |v63|, p1, DC10(), v64[-v0] + (map v65 : D0 | v65 in v66 :: (v65) := (p1));
					globalState.f4 := v51.cf2;
					r0 := cf18;
				case DC8(cf15) =>
					var v67: seq<string> := [fm2(p1, p1, globalState), p0 + p0, seq(0x19d, i0  => ('e')), p0];
					v67 := v67;
					var v68: map<bool, char> := map[p1 := fm7(globalState)];
					globalState.f4 := (|v68| - |(seq(242, i1  => (cf17)))|) < (if (p1) then v0 else cf18);
					var v69: T0 := new C0(0x153, p1);
					var v70: array<int> := new int[18];
					var v71 := DC4(v70);
					var v72 := DC8(DC8(v71));
					var v73: map<T0, D3> := map[v69 := DC14(DC13(v69.f19, p1, v72))];
					var v74: array<bool> := new bool[10] [p1, v69.f19, p1, p1, v69.f19, p1, p1, v69.f19, p1, false];
					var v75 := DC12(v74);
					var v76 := DC14(v75);
					var v77: seq<map<T0, D3>> := [v73, map[v69 := v76]];
					var v78: multiset<bool> := multiset{p1, v69.f19};
					var v79: seq<map<T0, D3>> := [v77[|v78|]];
					var v80: set<bool> := {v69.f19, v69.f19, p1};
					var v81: map<int, set<bool>> := map[v69.f18 := v80];
					var v82: C0 := new C0(|p0|, p1);
					var v83: set<C0> := {v82};
					var v84: map<bool, bool> := map[v69.f19 := v69.f19];
					var v85 := DC11(cf17, cf18);
					var v86: map<int, bool> := map[-0x376 := p1];
					var v87: array<D2> := new D2[5] [v85, v85, v85, DC11('k', |v86|), v85];
					var v88: multiset<int> := multiset{v82.fm4(v69.f18, v69.f19, p1, globalState), cf18};
					var v89: array<int> := new int[28] [cf18, v0 * v0, 0x1f7 % cf18, |v79[v0]|, cf18, 0x2ea * fm0(|(if (v0 in v81) then v81[v0] else v80)|, true, globalState), cf18, |v83|, cf18, v69.f18, |p0|, v69.f18, -|p0|, |(v84 + v84)|, v69.f18, 0x252, -(v69.f18 + v0), v0, -v69.f18 % -v69.f18, if (fm3(false, cf18, v69.f18, v69.f19, globalState)) then |[p1, !p1, p1, v69.f19, v69.f19]| else |p0|, |map[--v69.f18 := v87]|, if (false) then v0 else v69.f18, v0 % (if (-v0 in v88) then v88[-v0] else v69.f18), v69.f18 * cf18, v69.f18 - v0, v69.f18 % v0, v69.f18 % v69.f18, 0x2c9];
					v70[727] := v69.f18;
					v70[727] := -(v69.f18 % v0) - (v69.f18 / v0);
					var v90: map<T0, bool> := map[v69 := v69.f19];
					var v91: map<bool, int> := map[p1 := |v90|];
					var v92: set<int> := {|v91[true := |p0|]|, -v69.f18};
					globalState.f4 := |v92| != v69.f18;
			}
			
			var v93: T0 := new C0(v0, !p1);
			var v94: map<seq<bool>, T0> := map[[!p1] := v93];
			var v95: seq<bool> := [p1];
			v94 := v94[v95 := v93];
			var v96: set<bool> := {v93.f19};
			var v97: map<map<int, int>, string> := map[(map[v0 := v1[v0]])[|v96| := v93.f18] := "iyjjr"];
			var v98: map<int, int> := map[v0 := |multiset{v93.f18}|];
			v97 := v97[v98 + v98 := p0 + p0];
			var v99: map<bool, int> := map[p1 := v93.f18];
			if ((if (p1) then if (v93.f19 in v99) then v99[v93.f19] else v93.f18 else v93.f18) < cf18) {
				var v100: multiset<bool> := multiset{v93.f19};
				var v101: array<multiset<bool>> := new multiset<bool>[14] [v100, v100 * multiset{false, p1}, v100[v93.f19 := 0x240], multiset(v95), multiset(v95), v100, v100, v100 * v100, v100, multiset{v93.f19, !v93.f19}, v100 - v100, v100[p1 := |p0|], v100, v100];
				v101[578] := v100;
				var v102 := DC15(v100);
				var v103: C0 := new C0(0x107, v102.cf24 > multiset(v95));
				var v104: array<bool> := new bool[16] [v93.f19, p1, p1, v93.f19, !p1, v93.f19, v93.f19, v93.f19, v93.f19, !p1, p1, p1, v93.f19, v93.f19, v93.f19, v93.f19];
				var v105 := DC12(v104);
				var v106: set<int> := {|v96|};
				var v107 := DC0(v106, p0);
				v101[578], v103, v105, v107 := if (p1) then v100 else v100, v103, v105, v107;
				v102 := v102;
				v98 := v98[v93.f18 % |p0| := v0];
				globalState.f4 := cf18 > cf18;
				var v108: set<C0> := {v103, v103};
				globalState.f16 := (0xaf - v0) % |(v108 * v108)|;
			} else {
				globalState.f4 := fm3(p1, if (v93.f19) then cf18 else v93.f18, cf18, p1, globalState);
				globalState.f0 := (-v93.f18 + fm0(v93.f18, v93.f19, globalState)) * -271;
				var v109: array<map<D0, bool>> := new map<D0, bool>[9];
				var v110: map<T0, bool> := map[v93 := v93.f19];
				var v111: set<map<T0, bool>> := {v110, v110, v110, v110, v110};
				var v112 := DC2(|v111|, multiset{v93.f18, v93.f18, v0}, v93.f19);
				var v113: map<D0, bool> := map[v112 := false];
				v109[165] := v113;
				v109[165] := v113;
				var v114: set<int> := {|v98|, cf18, v0};
				var v115: C0 := new C0(-v93.f18, v93.f19);
				var v116: seq<C0> := [v115, v115];
				var v117: map<int, bool> := map[|v114| := p1];
				var v118: multiset<bool> := multiset{p1, v93.f19, p1, v93.f19, v93.f19};
				var v119: multiset<T0> := multiset{v93};
				var v120 := DC11(cf17, cf18);
				var v121: array<int> := new int[29] [|p0|, v1[0x21b], v93.f18, v93.f18, if (p1) then fm0(v93.f18, v93.f19, globalState) else v0, v0 % v93.f18, v0, v1[|v114|], cf18 * |v116|, v93.f18, v93.f18, cf18 + v93.f18, 195, 0x3a2, v93.f18, |v95|, if (v93.f19) then v93.f18 else |v117|, cf18, v93.f18, if (fm3(v93.f19, v0, 0x68, p1, globalState) in v118) then v118[fm3(v93.f19, v0, 0x68, p1, globalState)] else v1[cf18], v93.f18, cf18, |"lcwsjh"|, |v119|, |p0|, v120.cf18, v115.fm4(v0, v93.f19, v93.f19, globalState), v0, v0];
				v121[481] := v0 % cf18;
				v121[481] := v93.f18;
				var v122: array<set<bool>> := new set<bool>[20] [v96 - v96, v96, v96, {p1}, v96, v96, v96, v96 - v96, v96 + v96, fm6(v121[481], p0[v93.f18 := cf17], !v93.f19, globalState) + v96, v96, v96 * v96, v96, v96 + v96, v96 * v96, v96, v96, {v93.f19, v93.f19}, v96 - v96, v96];
				v122[150] := v96;
				v122[150] := {v93.f19, !true} - v96;
			}
			
		case DC9(cf16) =>
			var v123: multiset<int> := multiset{|(multiset{p1, p1})[p1 := v0]|};
			var v124: T0 := new C0(v0, p1);
			var v125: map<T0, bool> := map[v124 := p1];
			globalState.f4 := fm3(p1, v1[if (fm0(|v125|, p1, globalState) in v123) then v123[fm0(|v125|, p1, globalState)] else v124.f18], v0, v124.f19, globalState);
			var v126: map<bool, bool> := map[p1 := !p1];
			var v127: multiset<bool> := multiset{p1, v124.f19, v124.f19};
			var v128: map<multiset<bool>, map<bool, bool>> := map[v127 := v126 + (map[p1 := p1])[p1 := true]];
			var v129: seq<bool> := [p1];
			var v130 := DC9(cf16);
			v126, globalState.f4, globalState.f4 := if ((multiset(v129) - fm8(fm0(-v124.f18, v124.f19, globalState), v130.(cf16 := cf16), !p1, p0, globalState)) in v128) then v128[multiset(v129) - fm8(fm0(-v124.f18, v124.f19, globalState), v130.(cf16 := cf16), !p1, p0, globalState)] else v126, true, v124.f19;
			var v131 := 'p';
			v131 := v131;
			globalState.f4 := v131 !in p0;
	}
	
	globalState.f16 := fm0(v0, p1, globalState);
	if (!!p1) {
		globalState.f16 := v1[v0];
		var v132: map<bool, int> := map[p1 := |map[v1 := p1]|];
		var v133: C0 := new C0(|v132|, p1);
		v133, globalState.f0 := v133, if (p1) then v0 else v0;
		var v134: map<int, bool> := map[v1[v0] / --fm0(v0, fm3(p1, v0, 0x3a4, p1, globalState), globalState) := p1];
		v134 := v134[v0 := p1];
		globalState.f0 := 0xa5;
		var v135: array<bool> := new bool[27];
		v135 := v135;
	} else {
		var v136: array<int> := new int[24] [v0, |p0|, v0 % v0, v0, v0 * v0, v0, v0, v0, 521, v0, v0, v0, 0x3e7, v0, v0, if (p1) then -v0 else v0, --(v0 - v0), v0, v0, 681, v0, v0, v0, fm0(v0, false, globalState)];
		v136[773] := if (p1) then v0 else 0x250;
		v136[773] := v0;
		globalState.f4 := p0 < p0;
		var v137: multiset<int> := multiset{v0};
		v136[773] := if (p1) then |v137| % v0 else v136[773];
		var v138: set<bool> := {!p1};
		match (fm10(p1, -v136[773], globalState).(cf12 := v138)) {
			case DC5(cf11) =>
				var v139 := new C0(|v137|, p1);
				var v140: T0 := new C0(v136[773], if (p1) then !!p1 else p1);
				v140 := new C0(|v1| / v136[773], p1);
				var v141: map<bool, int> := map[v140.f19 := v0];
				globalState.f16 := 0x279 * (if (p1 in v141) then v141[p1] else v136[773]);
				var v142: map<seq<int>, int> := map[[v140.f18] := v140.f18];
				var v143: multiset<multiset<int>> := multiset{v137, fm11(v142, p1, v138, globalState), v137[v0 := v0], v137};
				var v144: seq<multiset<int>> := [v137];
				v143 := v143 + multiset(v144);
			case DC6(cf12) =>
				var v146: array<set<map<int, int>>> := new set<map<int, int>>[16](i2 => (set v145 : map<int, int> | v145 in map[map[|p0| := v136[773]] := p1] :: (v145)) + {map[0xa5 := v0], map[v0 := |[p1]|]});
				var v147: map<int, int> := map[v136[773] := v0];
				var v148: set<map<int, int>> := {v147, v147};
				v146[673] := v148;
				var v149: seq<map<int, int>> := [v147, v147];
				v146[673] := v148 + (set v150 : map<int, int> | v150 in v149 :: (v150));
				var v151: map<int, bool> := map[-v0 := p1];
				v151 := v151 + v151;
				var v152: array<C0> := new C0[10];
				var v153: map<int, array<C0>> := map[v0 := v152];
				v153 := v153[|(map v154 : int | v154 in (set v155 : int | (-0x9a <= v155) && (v155 < -0x2d2) :: (v155 * v0)) :: (v154 - v0) := ('g'))| + v136[773] := v152];
				globalState.f4 := !p1;
			case DC7(cf13, cf14) =>
				var v156: array<bool> := new bool[3](i3 => false);
				v156[752] := p1;
				var v157: map<bool, bool> := map[p1 := v1[-cf14] !in v1];
				var v158: C0 := new C0(cf14, p1);
				var v159 := DC1(p1, v157, v158, v136[773]);
				v156[752] := !!(if (v159.cf2 in v157) then v157[v159.cf2] else p1);
				var v160 := DC12(v156);
				var v161: seq<D3> := [v160];
				var v162: map<bool, seq<D3>> := map[if (fm3(v156[752], cf13, v136[773], false, globalState)) then v156[752] else v156[752] := v161[|p0| := v160]];
				v162 := v162[v156[752] := v161];
				var v163 := DC4(v136);
				var v164: seq<bool> := [v156[752], v156[752]];
				v163, v158, globalState.f6 := v163, v158, if (v156[752]) then v164 else v164;
				cf14 := v1[cf14];
			case DC4(cf10) =>
				globalState.f4 := p1;
				globalState.f16 := v136[773];
				globalState.f4 := (v0 * v0) < v136[773];
				var v165: C0 := new C0(v136[773], p1);
				v165 := v165;
			case DC8(cf15) =>
				var v166: array<set<C0>> := new set<C0>[1];
				var v167: C0 := new C0(2, true);
				var v168: set<C0> := {v167, v167};
				v166[247] := v168;
				v166[247] := if (p1) then v168 else v168;
				v0 := 0x63;
				var v169: array<array<D2>> := new array<D2>[29];
				var v170: map<array<array<D2>>, bool> := map[v169 := p1];
				v170 := v170[v169 := p1];
				var v171: set<int> := {|p0|};
				globalState.f4 := fm3(p1, v0 % |p0|, |v171|, p1, globalState);
		}
		
		var v172: array<bool> := new bool[1] [p1];
		var v173: seq<string> := ["rlhubva"];
		v172[925] := p0 in v173;
		var v174: set<int> := {v136[773]};
		v172[925] := (v174 + v174) !! v174;
	}
	
	globalState.f4 := p1;
	var v175: seq<bool> := [p1, !p1, p1];
	for i4 := fm0(|v175|, p1, globalState) to v0 {
		var v176: C0 := new C0(i4, p1);
		v176 := v176;
		r0 := i4;
		var v177: array<bool> := new bool[4](i5 => false || p1);
		v177[718] := p1;
		var v178 := DC1(p1, map[p1 := p1], v176, i4);
		var v179: multiset<bool> := multiset{p1};
		var v180: map<D0, int> := map[v178 := |v179|];
		v177[718] := v1[|v180[v178 := i4]|] > (i4 - 258);
		if (v177[718]) {
			var v181 := new C0(v0, p1);
			globalState.f16 := v0;
			var v182 := 'f';
			var v183: array<string> := new string[25] [p0 + p0, p0 + (seq(0x93, i6  => ('s'))), seq(0x122, i7  => ('y')), p0, p0 + p0[v0 := v182], fm2(true, p1, globalState), fm2(p1, v177[718], globalState), seq(0x6a, i8  => (v182)), p0, p0, "lo", p0, p0, p0, "pyu" + p0, p0, (p0 + p0)[i4 := v182], p0, "ujrexjxx", seq(-0x268, i9  => (v182)), ("ekt")[v0 := v182] + p0, p0, fm2(p1, v177[718], globalState), fm2(p1, v177[718], globalState), p0];
			v183[606] := p0;
			v183[606] := p0;
			globalState.f4 := p1;
			globalState.f4, v183[606], v182, globalState.f16 := v182 !in (p0 + v183[606]), (seq(0x11c, i10  => (v182))) + (v183[606] + (seq(0x110, i11  => (v182)))), fm7(globalState), i4;
		} else {
			globalState.f4 := true;
			var v184: seq<seq<int>> := [v1];
			var v185: array<seq<int>> := new seq<int>[12] [v1[v0 := |"r"|], v1, [331, i4] + v1, v1 + v1, v1, [v0], (v184[v1[v0]])[i4 := v0], v1 + (seq(0x2e1, i12  => (|v179|))), v1, v1, [|v179|, v0], v1 + v1];
			v185[838] := v1;
			var v186: multiset<C0> := multiset{v176};
			var v187: multiset<multiset<C0>> := multiset{v186, v186, v186};
			var v188 := "lhknlfumr";
			var v189: set<bool> := {p1};
			var v190: map<int, int> := map[v0 := i4];
			var v191: map<map<int, int>, set<bool>> := map[v190 := v189];
			v185[838], v177[718], v177[718], v187, v188 := v1, p1, (v189 * (if (v190 in v191) then v191[v190] else v189)) > v189, v187, v188 + "jduntgx";
			var v192 := new C0(|v189|, false);
			globalState.f16 := |(v189 * (v189 + v189))|;
			globalState.f16 := (v0 % i4) / (|v190| / v0);
		}
		
	}
	r0 := 0x35b / 457;
}
trait T0 {
	var f18 : int
	const f19 : bool
}

class C0 extends T0 {
	constructor (f18 : int, f19 : bool) {
		this.f18 := f18;
		this.f19 := f19;
	}
	
	function fm4(p0: int, p1: bool, p2: bool, globalState: GlobalState): int {
		-f18
	}
}

method Main() {
	var v0 := true;
	var v1: seq<bool> := [v0, v0, v0];
	var v2: array<int> := new int[15];
	var v3 := -235;
	var v4: map<bool, int> := map[v0 := v3];
	var v5: seq<array<int>> := [v2];
	var v6: array<bool> := new bool[8] [!v0, !v0, v0, v0, true, v0, v0, v0];
	var v7: multiset<array<bool>> := multiset{v6};
	var globalState := new GlobalState(0x1e5, -62, 0x20d, false, false, 646, v1, 0x1ce, v2, 0x315, 0x262, v4, false, v5[v3], v6, v7 - v7, 0x315, 'd');
	var v8: map<bool, bool> := map[v0 := v0];
	var v9: set<bool> := {v0, if (v0 in v8) then v8[v0] else false};
	var i0 := 0;
	while (v1[-|(v9 - v9)|])
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v10: array<char> := new char[6];
		var v11 := 'n';
		v10[53] := v11;
		v10[53] := v11;
		var v12 := "jx";
		var v13 := m0(v12, !!!v0, globalState);
		var v14 := m0(v12, v0, globalState);
		globalState.f10 := v14;
	}
	v6[460] := !!v0;
	v6[460] := !(if (v0) then v0 else v0);
	if (if (false in v8) then v8[false] else v6[460]) {
		var v15 := "p";
		globalState.f16 := |v15|;
		var v16 := m0(v15, !(if (false) then v6[460] else v0), globalState);
		var v17: map<bool, seq<bool>> := map[v0 := v1];
		var v18: seq<map<bool, seq<bool>>> := [v17, v17];
		var v19: map<int, map<bool, seq<bool>>> := map[v3 := map[v6[460] := [true, v6[460], v6[460]]]];
		v6[460] := (v17 + v17) == ((v18[-v16])[v0 := v1] + (if (fm0(v16, v0, globalState) in v19) then v19[fm0(v16, v0, globalState)] else v17));
		var v20 := 'b';
		var v21: map<map<int, multiset<int>>, bool> := map[fm1(v20, globalState) := -v3 < v16];
		var v22: multiset<int> := multiset{|map[v0 := v20]|};
		var v23: map<int, multiset<int>> := map[-0x2b7 := v22];
		v0 := if (v23 in v21) then v21[v23] else v6[460];
		var v24: array<seq<bool>> := new seq<bool>[28](i1 => [v0]);
		var v27: array<map<int, set<bool>>> := new map<int, set<bool>>[1](i2 => (map v25 : int | (175 <= v25) && (v25 < -0xf8) :: (v25 + |(map v26 : int | (149 <= v26) && (v26 < 0x333) :: (v26 / v16) := (v20))|) := (v9)) + map[v3 := v9]);
		var v29: set<int> := {v16};
		v27[3] := map v28 : int | v28 in v29 :: (v28 % -v16) := ({v0});
		var v31: map<int, set<bool>> := map[v16 := v9];
		v20, v6[460], v24, v27[3], v2 := v20, (if (v6[460]) then v29 else v29) <= (set v30 : int | (0xb6 <= v30) && (v30 < 0x21b) :: (v30 / v16)), v24, v31 + v31, v2;
	} else {
		var v32 := "tkkvwrsmk";
		var v33: set<map<bool, int>> := {v4};
		var v34 := m0(v32, v33 > v33, globalState);
		v0 := true;
		v32 := ((seq(176, i3  => ('d'))) + fm2(v6[460], v0, globalState)) + (seq(-0x277, i4  => ('x')));
		v3 := |fm2(v0, v6[460], globalState)| + |(seq(62, i5  => ('o')))|;
		globalState.f4 := fm3(v0, v34, v34, v0, globalState);
	}
	
	var v35 := "rgyxvjy";
	var v36 := m0(v35, v1[v3], globalState);
	var v37: seq<int> := [v36, fm0(v36, !false, globalState), |v35|, 135];
	v35, v6, globalState.f0, v6[460], v6[460] := v35 + "ryj", v6, v36, fm3(v0, v3, --v3, v6[460], globalState) || (fm0(v36, v6[460], globalState) in v37), v1[v36];
	if (v35 <= v35) {
		var v38 := new C0(0x25e, v0);
		var v39: map<string, bool> := map[v35 := !false];
		var v40: map<int, int> := map[v36 := v38.fm4(v36, if (v35 in v39) then v39[v35] else v6[460], v6[460], globalState)];
		v40 := v40[v36 := v3];
		var v41: seq<seq<bool>> := [v1];
		var v42: map<int, seq<seq<bool>>> := map[v3 := v41];
		v41, v6, globalState.f4 := (((if (v36 in v42) then v42[v36] else v41) + v41)[-929 := fm5(globalState)])[v3 := [v0]], v6, fm0(|(seq(512, i6  => ('q')))|, v6[460], globalState) >= |(v35 + v35)|;
		var v43: map<seq<int>, int> := map[v37 := -v3];
		var v44: map<int, map<seq<int>, int>> := map[-|v37| / v36 := v43];
		var v45: multiset<int> := multiset{v37[v3], v36, v37[0x5f], --v36};
		v44 := v44[v36 := v43[v37 := if (v3 in v45) then v45[v3] else |v9|]];
		match (DC1(v0, map[v0 := v6[460]], v38, v36)) {
			case DC0(cf0, cf1) =>
				v38 := v38;
				globalState.f4 := [v6[460]] < v1;
				var v46: set<array<bool>> := {v6, v6, v6, v6};
				v46 := {v6, v6, v6, v6, v6};
				var v47: map<bool, seq<array<int>>> := map[v0 := v5];
				var v48 := DC4(v2);
				v47 := v47[v6[460] := [v48.cf10]];
			case DC1(cf2, cf3, cf4, cf5) =>
				cf4 := cf4;
				var v49: seq<C0> := [v38, cf4];
				v49 := v49 + [v38];
				v8 := v8[cf2 := v6[460]];
				var v50: map<int, bool> := map[|v37| := cf2];
				v6[460], globalState.f4, cf2 := v37[0xf8] <= (v36 % cf5), if (v3 in v50) then v50[v3] else v3 <= v37[0x1ce], true;
			case DC2(cf6, cf7, cf8) =>
				var v51 := m0(v35, v0, globalState);
				globalState.f4 := v6[460];
				var v52 := DC1(v0, v8, v38, |v9|);
				globalState.f4 := v52.cf2;
				globalState.f10 := v36;
			case DC3(cf9) =>
				var v53 := m0("tiftisjra", fm3(true, -0x232, v3, v6[460], globalState), globalState);
				var v54: map<C0, int> := map[v38 := v36];
				globalState.f10 := v36 * |v54[v38 := v53]|;
				v53 := 0x125;
				globalState.f16 := v53;
		}
		
	} else {
		var v56: map<int, int> := map[|(set v55 : int | (-0x4b <= v55) && (v55 < 0x1ff) :: (v55 / |v37|))| := v36];
		var v57 := DC9(v56);
		var v58: T0 := new C0(|v57.cf16|, v6[460]);
		globalState.f4, v58, v3, globalState.f4, v6[460] := !v6[460] <==> (v56 != v56), v58, (v3 + v58.f18) % v58.f18, v0, v0;
		var v59: C0 := new C0(|v1|, v58.f19);
		var v60: map<char, C0> := map['m' := v59];
		v60 := v60;
		var v61: set<int> := {v36};
		var v62 := DC2(|v61|, multiset{-v3, v36, v58.f18, 0x1d9, v58.f18}, v6[460]);
		globalState.f4, v37 := v62.cf8, v37;
		v36 := fm0(v3, fm3(v6[460], |v35|, |v9|, v58.f19, globalState), globalState);
		var v63 := DC5(0x320);
		v6[460] := fm3(!v0, fm0(|v4|, v0, globalState), v63.cf11, !false, globalState);
	}
	
	v1 := (v1 + v1) + ([v6[460]])[v36 := false];
	var v64: array<D2> := new D2[7];
	var v65: map<bool, array<D2>> := map[v0 := v64];
	if (v65[v6[460] := v64] != (v65 + v65)) {
		var v66: set<set<bool>> := {fm6(-990, v35, v6[460], globalState), v9};
		v36 := |(v66 + (v66 * v66))|;
		v37 := v37 + ([0xfa, v36] + v37);
		v0 := v0;
		if (!((v37 + [0x198, v36]) <= v37)) {
			globalState.f0 := if (v0 in v4) then v4[v0] else v3 + v3;
			var v67 := DC4(v2);
			var v68: map<bool, D1> := map[v6[460] := v67];
			var v69: C0 := new C0(v36, v6[460]);
			var v70: seq<C0> := [v69];
			v68 := v68[v1[|v70|] := DC4(v2)];
			var v71 := 'o';
			var v72: map<int, char> := map[v3 := v71];
			var v73: multiset<int> := multiset{v3, --|v72|, v36, v36, v36};
			globalState.f0 := |v73[v3 := |v35|]|;
			var v74 := DC5(v36);
			var v75 := new C0(-(v74.cf11 - v3), v0);
			var v76: map<C0, bool> := map[v75 := v0];
			v76 := v76[if (v0) then v75 else v75 := !(if (v6[460]) then v6[460] else v6[460])];
		} else {
			var v77 := 'o';
			globalState.f4 := v36 <= fm0(|v35[-0x10c := v77]|, v0, globalState);
			var v78: multiset<bool> := multiset{v6[460]};
			var v79: set<int> := {v3, |v78|};
			var v80 := new C0(v3 - DC7(|v79|, v36).cf13, v6[460]);
			var v81: array<char> := new char[1];
			var v82: multiset<int> := multiset{v3, v36};
			var v83: map<int, multiset<int>> := map[if (v3 in v82) then v82[v3] else 125 := v82];
			var v84: map<map<int, multiset<int>>, array<char>> := map[v83 + v83 := v81];
			v81 := if (v83 in v84) then v84[v83] else v81;
			var v85: array<string> := new string[9];
			v85 := v85;
			globalState.f10 := v3;
		}
		
		var v86: array<string> := new string[17];
		v86[265] := v35;
		var v87: C0 := new C0(v3, v6[460]);
		var v88: seq<string> := [v35, "hl" + v35];
		var v89: seq<C0> := [v87, v87, v87, v87];
		v35, v86[265], v87, v3 := seq(-560, i7  => ('j')), v88[-v36], v89[v36 + |v37|], -v36;
	} else {
		var v90 := 'n';
		globalState.f4 := v90 !in v35[0x312 := v90];
		if (!v6[460]) {
			globalState.f16 := v36 + v3;
			var v91: array<seq<bool>> := new seq<bool>[22];
			v91[585] := fm5(globalState) + v1;
			v91[585] := [false, v35 == v35, v6[460]];
			var v92 := new C0(-0xee, v6[460] <==> fm3(false, v36, v36, v0, globalState));
			globalState.f10 := v92.fm4(v36, v0, v6[460], globalState);
			var v93 := DC4(v2);
			var v94: map<D1, array<bool>> := map[v93 := v6];
			v94 := v94[v93 := v6];
		} else {
			var v95: multiset<int> := multiset{v3};
			var v96: multiset<bool> := multiset{!!(v95 !! v95)};
			v96 := v96 * v96;
			v90 := 'k';
			var v97: array<char> := new char[22] [v90, v90, 't', 'k', v90, v90, fm7(globalState), v90, v90, v90, v90, v90, v90, v90, v90, v90, v90, v90, 'n', v90, v90, v90];
			var v98: map<bool, array<char>> := map[v0 := v97];
			v98 := v98;
			v6[460] := v90 in ((seq(-0x329, i8  => (v90))) + "qhsmdmwu");
			v0, globalState.f4, v6 := true, v95 >= (v95 * v95), v6;
		}
		
		var v99: C0 := new C0(0x28b, v1 <= [v6[460]]);
		v99 := v99;
		globalState.f10 := v36;
		var v100: map<int, string> := map[v36 := v35];
		var v101: seq<string> := [v35, v35, seq(0x54, i11  => (v90))];
		var v102: array<string> := new string[17] [v35, v35 + (seq(0x3bd, i9  => (v90))), v35, v35, if (v6[460]) then if (v36 in v100) then v100[v36] else v35 else v35[v36 := v90], v35, v35, "pgyyvl" + v35, "llc", "cqphv", v35 + "npufrcjrg", v35 + v35, "sa", v35 + v35[v3 := v90], if (v6[460]) then v35 else seq(0x39b, i10  => (v90)), fm2(!!v6[460], v6[460], globalState), v101[v3]];
		v102[353] := v35;
		v102[353] := if (v6[460]) then v35 else v35[v3 := v90];
	}
	
	var v103: multiset<bool> := multiset{v6[460]};
	var v104: map<multiset<bool>, string> := map[v103 := v35];
	var v105: map<bool, string> := map[v0 := v35];
	v104 := (v104[v103 := v35])[fm8(|v8|, DC9(fm9(v0, globalState)), v6[460], v35, globalState) := ("xifslf")[v3 := v35[v36]] + (if (v0 in v105) then v105[v0] else "m")];
	var i12 := 0;
	while (v0)
		decreases 100 - i12
	{
		if (i12 >= 100) {
			break;
		}
		
		i12 := i12 + 1;
		var v106: array<map<int, bool>> := new map<int, bool>[18];
		var v107: array<string> := new string[6](i13 => "tof");
		var v108: map<array<map<int, bool>>, array<string>> := map[v106 := v107];
		v108 := v108[v106 := v107];
		var v109 := 'r';
		var v110 := DC11(v109, v36);
		var v111: map<D2, string> := map[v110 := v35];
		v111 := v111[v110 := v35];
		v36 := v3;
		var v112: set<multiset<int>> := {multiset(v37)};
		var v113: multiset<int> := multiset{v36};
		if (v112 > {v113, v113}) {
			v109 := v109;
			v0 := v6[460];
			var v114: map<int, int> := map[v36 := 0x322];
			var v115 := DC2(-(if (v3 in v114) then v114[v3] else v36), v113, v6[460]);
			var v116: set<D0> := {DC2(v36, v113, v0)};
			var v117 := m0(v35, !!({v115} !! v116), globalState);
			var v118: map<int, bool> := map[v117 := v0];
			var v119: C0 := new C0(-v36, v0);
			var v120: multiset<C0> := multiset{if (if (v117 in v118) then v118[v117] else v0) then v119 else v119, v119};
			var v121: seq<C0> := [v119];
			v120 := v120 * multiset(v121);
			globalState.f16 := v3;
		} else {
			var v122: C0 := new C0(v36, v6[460]);
			v122 := v122;
			globalState.f4 := v0;
			globalState.f16 := v3;
			var v123: array<C0> := new C0[16];
			var v124 := DC10();
			var v125: map<array<C0>, D2> := map[v123 := v124];
			globalState.f16 := |v125[v123 := v124]|;
			v4 := v4[v6[460] := v36];
		}
		
	}
	globalState.f4 := v0;
	forall i14 | 0 <= i14 < v6.Length {
		v6[i14] := |(v35 + v35)| in (map[v36 := false] + map[v3 := v0]);
	}
	v6[460] := (v3 == |(if (false in v105) then v105[false] else v35)|) && v0;
	if (v6[460]) {
		v7 := (if (v0) then v7 else multiset{v6, v6, v6, v6, v6}) - v7;
		if (v3 != -0xbb) {
			var v126 := m0(v35 + v35, v6[460], globalState);
			globalState.f6 := (v1 + v1) + [v0];
			globalState.f4 := v0;
			globalState.f4 := (if (v0) then false else v6[460]) && (v6[460] ==> v0);
			var v127: C0 := new C0(v36, v0);
			var v128: multiset<int> := multiset{563, DC1(v0, v8, v127, v3).cf5};
			globalState.f4 := v128 != (v128 + multiset{v36});
		} else {
			var v129: map<int, bool> := map[394 := v6[460]];
			v129 := v129;
			v6[460] := v0;
			v6[460] := v6[460];
			var v130: array<char> := new char[2];
			var v131 := 'b';
			var v132: seq<array<char>> := [v130, v130, v130];
			var v133: C0 := new C0(v3, v0);
			var v134: multiset<C0> := multiset{v133};
			v130, v35, v6[460], globalState.f0, v131 := v132[|v134|], v35 + v35, true, |(if (v0) then {!v6[460], v0} else v9)|, v131;
			globalState.f0 := |v1| * (|v35| % v3);
		}
		
		var v135 := new C0(v36 / v36, v3 > v3);
		var v136: array<multiset<bool>> := new multiset<bool>[29];
		v136[74] := v103;
		v136[74] := (v103 * v103) - v103;
		var v137: array<set<D0>> := new set<D0>[4];
		v137 := v137;
	} else {
		var v138 := new C0(0x10c, v6[460]);
		v6[460] := v0;
		var v139: multiset<int> := multiset{826};
		var v140: T0 := new C0(|v139|, true);
		var v141 := 'v';
		v6[460], v140, v141, globalState.f4, globalState.f16 := !!v0, v140, v141, !!v0, fm0(fm0(v140.f18, v140.f19, globalState), v140.f19, globalState) / |[v141, v141, v141, v141]|;
		var v142 := new C0(v36, v0);
		v6[460] := v35 == v35;
	}
	
	v35 := v35;
	var v143: seq<string> := [v35];
	var v144 := m0(v143[v36], v1 != v1, globalState);
}