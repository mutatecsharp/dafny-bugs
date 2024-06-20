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
datatype D0 = DC1(cf1: multiset<seq<bool>>, cf2: int) | DC0(cf0: bool) | DC2(cf3: D0)
datatype D1 = DC4 | DC3(cf4: seq<int>) | DC5(cf5: D1)
datatype D2 = DC7(cf7: int, cf8: C0) | DC6(cf6: C0) | DC8(cf9: D2)
datatype D3 = DC10(cf11: char, cf12: bool) | DC11(cf13: int, cf14: int, cf15: bool, cf16: bool) | DC9(cf10: seq<bool>)
datatype D4 = DC13(cf18: int, cf19: int, cf20: bool) | DC14(cf21: C0, cf22: bool) | DC12(cf17: array<int>)
datatype D5 = DC15(cf23: T0)
datatype D6 = DC17(cf25: bool, cf26: bool, cf27: bool) | DC18(cf28: C0) | DC16(cf24: set<int>)
datatype D7 = DC20(cf30: bool, cf31: bool, cf32: C0) | DC19(cf29: multiset<int>) | DC21(cf33: D7)
datatype D8 = DC23(cf35: bool, cf36: bool, cf37: C0) | DC24(cf38: bool, cf39: array<int>, cf40: bool) | DC25(cf41: D6) | DC22(cf34: set<seq<int>>) | DC26(cf42: D8)
datatype D9 = DC28(cf44: bool, cf45: int, cf46: array<C0>) | DC29(cf47: int, cf48: bool, cf49: bool) | DC30(cf50: int, cf51: int, cf52: int, cf53: array<T0>) | DC27(cf43: multiset<set<bool>>)
datatype D10 = DC31(cf54: string)
datatype D11 = DC33(cf56: bool) | DC32(cf55: multiset<string>) | DC34(cf57: D11)
datatype D12 = DC36(cf59: int, cf60: int, cf61: bool, cf62: int, cf63: bool) | DC37(cf64: int, cf65: bool, cf66: bool) | DC35(cf58: map<bool, bool>)
trait T0 {
	var f7 : bool
	function fm4(globalState: GlobalState): int 
}

class GlobalState {
	const f0 : bool
	const f1 : string
	const f2 : bool
	var f3 : int
	var f4 : map<char, bool>
	const f5 : multiset<multiset<int>>
	var f6 : int
	constructor (f0 : bool, f1 : string, f2 : bool, f3 : int, f4 : map<char, bool>, f5 : multiset<multiset<int>>, f6 : int) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
	}
	
}

class C0 extends T0 {
	const f8 : int
	const f9 : int
	constructor (f8 : int, f9 : int, f7 : bool) {
		this.f8 := f8;
		this.f9 := f9;
		this.f7 := f7;
	}
	
	function fm4(globalState: GlobalState): int {
		-0x27b - f9
	}
	function fm5(p0: map<bool, bool>, p1: char, p2: int, p3: bool, globalState: GlobalState): int {
		|map[if (f7) then f7 else f7 := f7]|
	}
}

function fm0(p0: int, p1: bool, p2: seq<int>, p3: char, globalState: GlobalState): int {
	safeDivisionInt(|multiset{DC13(|[0x2f]|, -0x10b, true).cf20, false, false, !true, false}|, |(seq(abs(0x27), i0  => ("hrdfpvvp")))| - |[true, DC0(true).cf0]|)
}
function fm1(globalState: GlobalState): D0 {
	match DC2(DC1(multiset{[true, true]}, 0x1e4)) {
		case DC1(cf1, cf2) => DC1(cf1, |map['m' := cf2]|)
		case DC0(cf0) => DC1(multiset([[cf0]]), 0xaa)
		case DC2(cf3) => DC1(multiset{[!false]}, 0x105)
	}
}
function fm2(p0: map<multiset<int>, bool>, p1: char, globalState: GlobalState): set<int> {
	{-|[-262]|} + (set v0 : int | (0x1e1 <= v0) && (v0 < 0x15) :: (v0 - -863))
}
function fm3(p0: int, globalState: GlobalState): bool {
	map[map[false := 'n'] := false] == map[map[!true := 'i'] := true]
}
function fm6(p0: int, p1: char, p2: bool, p3: bool, globalState: GlobalState): seq<bool> {
	DC9([false]).cf10
}
function fm7(globalState: GlobalState): seq<int> {
	[|[true]|, |DC32(multiset(seq(abs(0x2be), i0  => ("re")))).cf55|, |[0x205, 0x2c]|] + ([|[false]|] + [0x341, -|map[!true := !false]|])
}
function fm8(p0: int, p1: int, p2: int, p3: int, globalState: GlobalState): map<bool, bool> {
	(map[false := false] + DC35(map[false := true]).cf58) + map[false := !false]
}
function fm9(globalState: GlobalState): D1 {
	DC3([|(map v0 : int | (-0xa <= v0) && (v0 < 0x2b3) :: (v0 + 155) := (|['f']|))|])
}
function fm10(p0: int, p1: bool, p2: int, p3: multiset<seq<int>>, globalState: GlobalState): string {
	"hrux" + (seq(abs(0x15a), i0  => ('k')))
}
function fm11(p0: multiset<map<bool, int>>, p1: int, p2: string, p3: string, globalState: GlobalState): char {
	'l'
}
function fm12(p0: int, p1: int, p2: map<bool, int>, globalState: GlobalState): multiset<bool> {
	multiset{false} + multiset([DC11(|[false]|, |"ihqlrtng"|, true, !true).cf15])
}
function fm13(p0: int, globalState: GlobalState): map<bool, int> {
	(map[false := -488] + map[true := --168]) + (map[false := -929] + map[false := -0x2cb])
}
function fm14(globalState: GlobalState): set<char> {
	{'u'} - ((set v0 : char | v0 in multiset{'b', 'd'} :: (v0)) + {'q', 'n'})
}
function fm15(p0: bool, p1: bool, p2: int, globalState: GlobalState): map<int, bool> {
	map[0x2e2 := true <==> !!false]
}
function fm16(p0: bool, p1: int, p2: bool, p3: bool, globalState: GlobalState): map<int, int> {
	map[-0x15e := |{'q', 'l'}|] + ((map v0 : int | v0 in {0x104} :: (safeModuloInt(v0, |"uvbomwi"|)) := (|(seq(abs(0x312), i0  => ('y')))|)) + (map v1 : int | (0x3ca <= v1) && (v1 < -978) :: (v1 * 0x89) := (0x1e4)))
}
function fm17(p0: bool, p1: bool, p2: bool, globalState: GlobalState): seq<seq<bool>> {
	[[true]]
}
method m0(p0: int, p1: multiset<bool>, p2: seq<bool>, globalState: GlobalState) returns (r0: string) {
	for i0 := p0 to p0 {
		var v0: array<seq<seq<int>>> := new seq<seq<int>>[25](i1 => if (false) then seq(abs(899), i2  => ([p0, |{p0, |"wq"|}|, |multiset{p0}|])) else [[i0, i0], seq(abs(0x1d2), i3  => (0x8f)), [p0]]);
		var v1: set<int> := {-p0};
		var v2: seq<int> := [p0, |v1|, p0];
		var v3: seq<seq<int>> := [v2];
		v0[safeIndex(570, v0.Length)] := v3;
		var v4 := false;
		var v5: array<set<int>> := new set<int>[26];
		v5[safeIndex(698, v5.Length)] := v1;
		var v6 := 'n';
		var v7: map<int, char> := map[p0 := v6];
		var v8: C0 := new C0(i0, -i0, true);
		var v9 := DC7(|v7|, v8);
		v0[safeIndex(570, v0.Length)], v4, globalState.f3, globalState.f6, v5[safeIndex(698, v5.Length)] := (seq(abs(0x271), i4  => (v2))) + [seq(abs(-0x392), i5  => (p0))], v4, v9.cf7, v8.f9, v1;
		var v10: array<C0> := new C0[21];
		v10[safeIndex(694, v10.Length)] := v8;
		v10[safeIndex(694, v10.Length)] := v8;
		v4 := v4;
		if (v4) {
			var v11: array<int> := new int[3](i6 => safeModuloInt(i6, |p1|));
			var v12 := "kn";
			var v13: set<C0> := {v8};
			var v14: multiset<set<C0>> := multiset{v13};
			v11[safeIndex(243, v11.Length)] := safeDivisionInt(|v12|, if ({v10[safeIndex(694, v10.Length)], v8} in v14) then v14[{v10[safeIndex(694, v10.Length)], v8}] else v2[safeIndex(|v12|, |v2|)]);
			v11[safeIndex(243, v11.Length)] := 0x227;
			var v15: map<int, int> := map[v8.f9 := v8.f9];
			globalState.f6 := 42 * safeDivisionInt(v8.f9, |v15|);
			var v16: seq<bool> := [v4, v4, !!v4 !in p1];
			var v17 := DC14(v10[safeIndex(694, v10.Length)], v4);
			var v18: map<C0, char> := map[v10[safeIndex(694, v10.Length)] := v6];
			var v19: map<D4, char> := map[v17 := if (v10[safeIndex(694, v10.Length)] in v18) then v18[v10[safeIndex(694, v10.Length)]] else v6];
			var v20: multiset<int> := multiset{|v19|};
			var v21: T0 := new C0(i0, |v16|, v4);
			var v22: seq<T0> := [v21, v21, v21];
			v16, v6, globalState.f6, v11[safeIndex(243, v11.Length)] := v16, v6, -|v20|, (-|v22| * v11[safeIndex(243, v11.Length)]) + v8.f8;
			var v23 := DC17(v21.f7, false, v4);
			var v24: seq<D6> := [v23, v23, v23];
			v4 := (v24 + [v23]) < ((seq(abs(576), i7  => (v23))) + (v24[safeIndex(|"ljh"|, |v24|) := v23])[safeIndex(v8.f8, |v24[safeIndex(|"ljh"|, |v24|) := v23]|) := DC17(v21.f7, v21.f7, v21.f7)]);
			var v25: seq<seq<bool>> := [v16];
			var v26: multiset<seq<bool>> := multiset{v16, v16, p2, p2};
			var v27: seq<multiset<seq<bool>>> := [multiset(v25), v26];
			var v28 := DC1(v27[safeIndex(|v12|, |v27|)], safeModuloInt(0xe9, p0));
			v16, v8, v28 := v16, if (fm3(v11[safeIndex(243, v11.Length)], globalState)) then v8 else v10[safeIndex(694, v10.Length)], v28;
		} else {
			v1 := v5[safeIndex(698, v5.Length)];
			var v29 := "dled";
			var v30: seq<string> := [v29, "vsxctdar", v29, "gmyfssb"];
			v30 := v30;
			var v31 := new C0(v8.f8, v8.f9, p0 > |(seq(abs(0x281), i8  => (v6)))|);
			var v32 := new C0(|(seq(abs(38), i9  => (v6)))[safeIndex(v8.f8, |(seq(abs(38), i9  => (v6)))|) := v6]|, 0x193, v4);
			var v33: map<int, bool> := map[p0 := false];
			var v34: map<bool, int> := map[if (-i0 in v33) then v33[-i0] else v4 := v2[safeIndex(-p0, |v2|)]];
			var v35 := DC10(v6, v4);
			v34 := v34[v35.cf12 := safeModuloInt(v8.f9, -0x187)];
		}
		
	}
	var v36 := true;
	var v37 := "deqbibs";
	for i10 := if (!v36) then |v37| else |v37| to p0 {
		var v38: array<multiset<multiset<int>>> := new multiset<multiset<int>>[8];
		var v39: seq<int> := [p0];
		var v40 := 'x';
		v38[safeIndex(102, v38.Length)] := multiset([multiset{fm0(0x154, v36, v39, v40, globalState), p0, i10}]);
		var v41: array<bool> := new bool[5](i11 => v40 in "uyntxwc");
		v41[safeIndex(756, v41.Length)] := p0 == p0;
		var v42: array<multiset<bool>> := new multiset<bool>[10];
		v42[safeIndex(780, v42.Length)] := p1;
		var v43 := DC17(v36, v36, v36);
		var v44: map<D6, int> := map[v43 := p0];
		var v45: multiset<map<D6, int>> := multiset{v44, v44, v44, v44};
		var v46: C0 := new C0(0x29a, i10, false);
		var v47 := DC7(p0, v46);
		var v48: multiset<int> := multiset{-p0, v47.cf7, p0, p0};
		var v49 := DC3(v39);
		var v50 := DC19(multiset(v39));
		var v51: multiset<multiset<int>> := multiset{v48, multiset(v49.cf4), v50.cf29, v48};
		v36, v38[safeIndex(102, v38.Length)], v41, v41[safeIndex(756, v41.Length)], v42[safeIndex(780, v42.Length)] := !false ==> (v45 !! v45), v51, v41, fm7(globalState) != v39, multiset(p2);
		var v52: T0 := new C0(v46.f9, p0, !v41[safeIndex(756, v41.Length)]);
		var v53: seq<T0> := [v52];
		var v54: array<int> := new int[4] [p0, safeDivisionInt(v46.f8, |map["xpqepkn" := v53[safeIndex(921, |v53|)]]|), v46.f9, safeModuloInt(i10, v46.f9)];
		v54[safeIndex(916, v54.Length)] := p0;
		v54[safeIndex(916, v54.Length)] := v46.f8;
		var v55: seq<C0> := [v46];
		var v56: set<int> := {i10, |v55|};
		var v57: set<set<int>> := {v56 + {|p2|}};
		v57 := {v56};
		globalState.f3 := -0x125;
	}
	var v58: map<int, string> := map[p0 := v37];
	v37 := if (-(if (false in p1) then p1[false] else p0) in v58) then v58[-(if (false in p1) then p1[false] else p0)] else v37;
	if (p0 != 784) {
		var v59: set<string> := {"lmjjoibd" + v37, v37, v37, v37};
		v59 := (v59 * v59) - v59;
		if (|v37| >= 106) {
			var v60: seq<int> := [p0];
			var v61 := 'd';
			var v62: T0 := new C0(fm0(p0, v36, v60, v61, globalState), |v37|, v36);
			var v63: multiset<T0> := multiset{v62, v62};
			v36 := (if (v36) then v63 else v63) < v63;
			var v64: C0 := new C0(520, p0, v36);
			var v65: seq<C0> := [v64];
			var v66: array<int> := new int[16] [fm0(p0, v36, [|p2|], v61, globalState), 0x1e3, |v65|, p0, v64.f9, v64.f8, v64.f9, v64.f8, -166, v64.f9, v64.f9, v64.f8, v64.f9, v64.f8, -v60[safeIndex(p0, |v60|)], -v64.f8];
			var v67 := DC20(v36, v62.f7, v64);
			var v68: map<array<int>, D7> := map[v66 := if (v36) then v67 else v67];
			var v69: map<bool, array<int>> := map[v36 := v66];
			var v70: seq<map<array<int>, D7>> := [v68];
			v68 := map[if (v36 in v69) then v69[v36] else v66 := v67] + v70[safeIndex(v64.f8, |v70|)];
			var v71: map<char, T0> := map[v61 := v62];
			var v72: map<bool, int> := map[v62.f7 := v64.f9];
			var v73: multiset<map<bool, int>> := multiset{v72, v72};
			var v74: array<map<char, T0>> := new map<char, T0>[15] [v71[fm11(v73, p0, v37, "aicyyk", globalState) := v62], v71, v71, v71, v71, map[v61 := v62], v71, v71, v71, v71, v71, v71, v71, v71, v71[v61 := v62]];
			var v75: set<seq<int>> := {v60, v60};
			var v76: map<array<map<char, T0>>, set<seq<int>>> := map[v74 := v75];
			var v77 := DC22(v75);
			v76 := v76[v74 := v77.cf34 * v75];
			var v78: array<bool> := new bool[28](i12 => v62.f7);
			var v79: map<string, array<bool>> := map[v37 := v78];
			v79 := (v79 + v79) + v79;
			var v80: seq<array<bool>> := [v78, v78];
			v80 := v80;
		} else {
			var v81: map<int, bool> := map[p0 := v36];
			var v82: seq<map<int, bool>> := [v81, fm15(true, v36, p0, globalState)];
			v36 := [v81] != ([v81, v81, fm15(v36, fm3(p0, globalState), 137, globalState), map[-p0 := v36]] + v82);
			var v83: array<char> := new char[7];
			var v84 := 'f';
			v83[safeIndex(974, v83.Length)] := v84;
			v83[safeIndex(974, v83.Length)] := v84;
			var v85: map<bool, bool> := map[v36 := false];
			v85 := v85[true := v36 || v36];
			var v86: set<int> := {p0, -0x8a};
			var v88: multiset<set<int>> := multiset{v86 * (set v87 : int | v87 in v86 :: (safeDivisionInt(v87, 948)))};
			var v90: set<bool> := {v36};
			var v91: multiset<set<bool>> := multiset{v90};
			var v92 := DC27(v91);
			var v93: seq<multiset<set<bool>>> := [v91, v91, v92.cf43, v91, v91];
			var v94: multiset<int> := multiset{p0};
			var v95: map<int, int> := map[if (123 in v94) then v94[123] else p0 := p0];
			v36, r0, globalState.f6 := v36, "uownf", if ((v86 + (set v89 : int | (0x71 <= v89) && (v89 < 0x39e) :: (v89 - p0))) in v88) then v88[v86 + (set v89 : int | (0x71 <= v89) && (v89 < 0x39e) :: (v89 - p0))] else |v93[safeIndex(|v95|, |v93|)]|;
			var v96: C0 := new C0(|fm16(v36, |v37|, v36, v36, globalState)|, p0, v36);
			var v97 := DC18(v96);
			var v98: array<D6> := new D6[6] [v97, v97, DC18(v96), v97, v97, v97];
			v98[safeIndex(70, v98.Length)] := DC18(v96);
			v98[safeIndex(70, v98.Length)] := v97.(cf28 := v96);
		}
		
		var v99: T0 := new C0(-p0, p0, v36);
		var v100 := new C0(safeDivisionInt(|[v99, v99]|, if (v99.f7 in p1) then p1[v99.f7] else p0), p0, v99.f7);
		globalState.f6 := v100.f8;
		v99.f7 := v36;
	} else {
		globalState.f3 := p0;
		v36 := v36;
		var v101: T0 := new C0(|p1|, p0, fm3(p0, globalState));
		var v102: set<T0> := {v101};
		if (v101 !in v102) {
			globalState.f3 := p0 - p0;
			var v103: array<bool> := new bool[28](i13 => v101.f7);
			var v104: multiset<array<bool>> := multiset{v103};
			v104 := multiset{v103} + v104;
			var v105: C0 := new C0(p0, -p0, v36);
			var v106 := DC20(true, false, v105);
			v105 := v106.cf32;
			var v107 := new C0(safeDivisionInt(p0, |map[p0 := v101]|), safeDivisionInt(v105.f9, v105.f8), fm3(v105.f8, globalState));
			v105 := v105;
		} else {
			v101.f7 := p2[safeIndex(p0, |p2|)];
			var v108: array<bool> := new bool[27];
			v108[safeIndex(243, v108.Length)] := false;
			var v109: map<int, multiset<bool>> := map[p0 := p1];
			v108[safeIndex(243, v108.Length)] := p1 != (if (!v101.f7) then if (p0 in v109) then v109[p0] else p1 else p1);
			var v110: map<bool, bool> := map[true := v108[safeIndex(243, v108.Length)]];
			var v111: seq<int> := [|v110|];
			var v112: seq<int> := [|v111|];
			var v113 := DC22({v111});
			var v114: set<seq<int>> := {v112, v111};
			v113 := DC22(v114);
			var v115: C0 := new C0(v101.fm4(globalState), |map[false := p2]|, v108[safeIndex(243, v108.Length)]);
			var v116: map<C0, int> := map[v115 := p0];
			v116 := v116 + v116;
			v101 := v101;
		}
		
		var v117: C0 := new C0(|v37|, p0, v101.f7);
		var v118 := DC14(v117, v101.f7);
		var v119 := new C0(p0, p0 - p0, v118.cf22);
		var v120 := DC0(v36);
		v101.f7 := (v120.(cf0 := v101.f7)).cf0;
	}
	
	var v121: array<int> := new int[7](i14 => i14 * -0x303);
	var v122: map<bool, array<int>> := map[v36 := v121];
	v122 := v122[fm3(p0, globalState) := v121];
	var v123: multiset<seq<bool>> := multiset{p2};
	var v124 := DC1(multiset{p2, p2}, -p0);
	var v125: multiset<D0> := multiset{v124};
	for i15 := p0 to |(multiset{DC1(v123, p0)} + v125)| {
		var v126: seq<int> := [p0];
		match (DC3(v126 + v126)) {
			case DC4() =>
				r0 := v37;
				var v127: array<C0> := new C0[27];
				var v128: C0 := new C0(i15, v126[safeIndex(i15, |v126|)], v36);
				var v129: map<C0, C0> := map[v128 := v128];
				var v130: map<bool, map<C0, C0>> := map[v36 := v129];
				var v131 := 'q';
				var v132: multiset<int> := multiset{480, fm0(|(if (v36 in v130) then v130[v36] else v129)|, v36, seq(abs(0x39f), i16  => (v128.f9)), v131, globalState), -|v126|, |(seq(abs(0x1da), i17  => (v131)))|};
				var v133: C0 := new C0(p0, |v132|, v36);
				v127[safeIndex(808, v127.Length)] := v128;
				v127[safeIndex(808, v127.Length)] := v133;
				globalState.f6 := v128.f9 * safeDivisionInt(-448, i15);
				globalState.f3 := v133.f9;
			case DC3(cf4) =>
				var v134 := DC31(v37);
				var v135: multiset<seq<int>> := multiset{cf4, cf4};
				var v136 := 't';
				var v137: array<string> := new string[29] [v37, v37, v37, seq(abs(-0xf6), i18  => ('r')), "iqglrvh", v37, if (v36) then v37 else v37, v134.cf54, v37, v37, v37 + v37, v37 + v37, "ugemjeiny", v37, v37, v37, v37, v37, v37, fm10(p0, v36, i15, v135, globalState) + v37, v37, "foi" + v37, ("ptuxtb")[safeIndex(p0, |"ptuxtb"|) := v136], v37, v37, v37, v37, v37, v37];
				v137[safeIndex(936, v137.Length)] := v37;
				v137[safeIndex(936, v137.Length)] := "jquu";
				v121[safeIndex(933, v121.Length)] := p0;
				v121[safeIndex(933, v121.Length)] := i15;
				var v138: map<seq<bool>, int> := map[p2 := v121[safeIndex(933, v121.Length)]];
				var v139: map<map<seq<bool>, int>, char> := map[v138 := v136];
				v136 := if ((map v140 : seq<bool> | v140 in fm17(v36, v36, v36, globalState) :: (v140) := (i15)) in v139) then v139[map v140 : seq<bool> | v140 in fm17(v36, v36, v36, globalState) :: (v140) := (i15)] else 'g';
				var v141 := new C0(|p2|, i15, false);
			case DC5(cf5) =>
				v36 := v36;
				v36 := v37 !in multiset{v37, v37, v37, v37, v37};
				var v142: multiset<bool> := multiset{v36};
				v142 := p1;
				var v143: map<bool, int> := map[v36 := i15];
				v142 := (fm12(i15, i15, v143, globalState))[v36 := abs(i15)] + p1;
		}
		
		var v144: set<int> := {i15};
		var v145: multiset<set<int>> := multiset{v144};
		var v146: T0 := new C0((if (v144 in v145) then v145[v144] else p0) + i15, -i15, v36);
		var v147: map<int, T0> := map[i15 := v146];
		v146 := if (i15 in v147) then v147[i15] else v146;
		var v148: array<array<int>> := new array<int>[26];
		v148[safeIndex(300, v148.Length)] := v121;
		v148[safeIndex(300, v148.Length)] := new int[13];
		v146.f7 := v146.f7;
	}
	r0 := (v37 + v37) + ("oc" + v37);
}
method Main() {
	var v0 := "qfxbbb";
	var v1 := 0x30b;
	var v2 := 'p';
	var v3: map<int, char> := map[v1 := v2];
	var v4 := false;
	var v5: map<char, bool> := map[if (|v0| in v3) then v3[|v0|] else v2 := v4];
	var v6: multiset<int> := multiset{-0xa6};
	var v7: multiset<multiset<int>> := multiset{multiset{|v6|}};
	var globalState := new GlobalState(true, v0, true, 0xe1, v5 + v5, v7, 852);
	globalState.f6 := 0x9b;
	for i0 := |([v1])[safeIndex(|(map v8 : char | v8 in v0 :: (v8) := (v1))|, |[v1]|) := v1]| + 0x53 to v1 {
		var v9: seq<int> := [i0, 0x340, v1, v1];
		if (v9 < v9) {
			var v10: array<multiset<int>> := new multiset<int>[8](i1 => v6);
			v10 := v10;
			var v11: multiset<bool> := multiset{v4, !v4};
			var v12 := m0(i0, v11 * v11, ([v4])[safeIndex(|v9|, |[v4]|) := v4], globalState);
			globalState.f6, globalState.f3 := if (v4) then i0 else |v0|, i0;
			v4 := v4;
			var v13 := DC0(v4);
			v4 := v13.cf0;
		} else {
			var v14: set<int> := {i0, v1};
			globalState.f3, globalState.f3 := fm0((fm1(globalState)).cf2, true, v9, v2, globalState), safeDivisionInt(v1, |v14|);
			var v15: map<bool, multiset<int>> := map[v4 := multiset{0x3a0, v1}];
			var v16: map<multiset<int>, bool> := map[if (v4 in v15) then v15[v4] else v6 := false];
			v14 := fm2(v16, 'd', globalState);
			var v17: map<int, int> := map[i0 := v1];
			var v18: seq<bool> := [v4, v4];
			var v19: multiset<seq<bool>> := multiset{v18};
			var v20: map<bool, bool> := map[v4 := false];
			var v21 := DC1(v19, |v20|);
			var v22: map<string, bool> := map[v0 := v4];
			var v23: map<map<string, bool>, bool> := map[v22 := v4];
			var v24: set<bool> := {v4};
			globalState.f3, v17, v4, v4 := if (fm3(v1, globalState)) then if (v4) then v1 else i0 else v21.cf2, v17, if (v22 in v23) then v23[v22] else v24 > v24, v4;
			v4 := v4;
			var v25: array<int> := new int[14];
			v25[safeIndex(794, v25.Length)] := i0;
			v25[safeIndex(794, v25.Length)] := v1;
		}
		
		var v26: seq<bool> := [v4];
		var v27 := DC0(false);
		var v28: array<bool> := new bool[10] [v4, v26[safeIndex(0x1d0, |v26|)], v4, v4, v4, v4, true, v4, v4, v27.cf0];
		v28[safeIndex(685, v28.Length)] := v4;
		v28[safeIndex(685, v28.Length)] := v4;
		var v29: set<char> := {v2};
		var v30: multiset<set<char>> := multiset{v29 * v29, v29};
		var v31: seq<set<char>> := [v29];
		v30 := (multiset(v31) + v30)[v29 := abs(i0)];
		var v32: map<bool, int> := map[false := i0];
		var v33: seq<array<bool>> := [v28, v28];
		var v34: seq<multiset<int>> := [multiset{-v1}, v6];
		var v35 := DC3(v9);
		var v36: array<int> := new int[21] [|(seq(abs(-0x14c), i2  => (-|multiset{false}|)))[safeIndex(|v0|, |(seq(abs(-0x14c), i2  => (-|multiset{false}|)))|) := |v32|]|, 0x2e6, |v32|, v1, i0, v1, i0, 0x3c7, i0, 0x238, |v33|, v1, 0x365, safeModuloInt(v1, i0), v1, i0 - v1, v1, safeDivisionInt(fm0(|v34|, v28[safeIndex(685, v28.Length)], v9, v2, globalState), fm0(|v0|, v28[safeIndex(685, v28.Length)], v35.cf4, v2, globalState)), -i0, 0x3cf, -i0 + i0];
		v36[safeIndex(938, v36.Length)] := v1;
		v36[safeIndex(938, v36.Length)], v1, globalState.f6 := i0, i0, i0;
	}
	var v37: set<char> := {v2};
	var v38: set<set<char>> := {v37, v37, v37, {v2, v2}, v37};
	var v39: seq<set<set<char>>> := [v38];
	var v40: map<int, int> := map[v1 := v1];
	for i3 := |v39[safeIndex(v1, |v39|)]| to v1 - |v40| {
		if (v4) {
			var v41: array<seq<bool>> := new seq<bool>[1];
			var v42: array<int> := new int[18](i4 => safeModuloInt(i4, v1));
			var v43: map<array<seq<bool>>, array<int>> := map[v41 := v42];
			v43 := v43[v41 := v42];
			var v44 := new C0(v1, safeModuloInt(i3, 163), v4);
			v42[safeIndex(326, v42.Length)] := safeDivisionInt(v44.f8, i3);
			v42[safeIndex(326, v42.Length)] := -v1;
			var v45: set<bool> := {v4};
			var v46: multiset<bool> := multiset{fm3(v1, globalState)};
			var v47 := DC4();
			var v48: map<D1, bool> := map[v47 := v4];
			var v49 := m0(|v45| + i3, v46, fm6(v44.f9, 'p', if (v47 in v48) then v48[v47] else v4, v4, globalState) + fm6(v1, v2, v4, v4, globalState), globalState);
			var v50: seq<bool> := [v4, v4];
			var v51: seq<bool> := [v50[safeIndex(v44.f8, |v50|)]];
			var v52 := m0(0x1f0, v46, v51, globalState);
		} else {
			v4 := v6 !! v6;
			var v53: array<bool> := new bool[26](i5 => v4);
			var v54: map<int, array<bool>> := map[i3 := v53];
			v1 := |((v54 + map[v1 := v53]) + v54)|;
			var v55: map<bool, int> := map[v4 := i3];
			var v56: seq<bool> := [!true];
			var v57: seq<int> := [v1, v1];
			var v58: map<int, bool> := map[fm0(|v55|, v56[safeIndex(-|v0|, |v56|)], v57, v2, globalState) := v4];
			var v59 := new C0(0x3c9 - i3, safeDivisionInt(v1, |v55|), if (v1 in v58) then v58[v1] else v4);
			var v60: array<D1> := new D1[12];
			var v61 := DC3([0x92]);
			v60[safeIndex(998, v60.Length)] := v61;
			v60[safeIndex(998, v60.Length)] := v61;
			var v62: multiset<bool> := multiset{v4};
			var v63 := m0(if (v4) then v59.f9 else |v56|, multiset(v56) * v62, v56[safeIndex(i3, |v56|) := v4], globalState);
		}
		
		globalState.f6 := safeDivisionInt(safeDivisionInt(-i3, |v0|), v1);
		var v64: seq<int> := [v1];
		var v65: array<bool> := new bool[11] [true, v4, v4, false, !v4, v4, !v4, v4, fm3(-i3, globalState), !false, v4];
		var v66: map<int, array<bool>> := map[v64[safeIndex(i3, |v64|)] := v65];
		var v67: set<bool> := {v4};
		v66 := v66[|v67| := v65];
		if (v4) {
			var v68: multiset<bool> := multiset{v4};
			globalState.f3, v4 := v1 * |{-i3, v1, v1}|, v68 != multiset{v4};
			var v69 := new C0(i3, i3, v4 <==> v4);
			v4 := !DC0(v4).cf0;
			var v70: seq<multiset<bool>> := [v68];
			var v71: seq<bool> := [v4];
			var v72 := m0(safeDivisionInt(i3, fm0(v1, v4, v64, v2, globalState)), v68[v4 := abs(-i3)] - v70[safeIndex(v69.f9, |v70|)], v71, globalState);
			v4 := fm3(v69.f9, globalState);
		} else {
			v4 := v4;
			var v73: seq<bool> := [v4];
			var v74: seq<bool> := [v4, v73[safeIndex(|v6|, |v73|)]];
			v65[safeIndex(891, v65.Length)] := v73[safeIndex(v1, |v73|)];
			v65[safeIndex(891, v65.Length)] := v4;
			var v75: multiset<seq<bool>> := multiset{v73};
			var v76 := DC1(v75, v1);
			var v77 := DC2(v76);
			v77 := (if (v4) then v77 else v77).(cf3 := fm1(globalState));
			v65[safeIndex(891, v65.Length)] := fm3(if (v4) then i3 else i3, globalState);
			globalState.f3 := v1;
		}
		
	}
	for i6 := v1 to v1 {
		var v78 := new C0(--v1, i6, v4);
		v3 := v3[v78.f9 := v2];
		var v79: seq<bool> := [v4, v4, v4];
		var v80 := m0(-v78.f8, multiset{v4, false}, v79, globalState);
		globalState.f3 := -i6;
	}
	var v81: seq<int> := [v1];
	var v82: seq<int> := [v1, v81[safeIndex(v1, |v81|)]];
	var v83 := DC3(v81);
	var i7 := 0;
	while (match v83 {
		case DC4() => v4
		case DC3(cf4) => v4
		case DC5(cf5) => |"nol"| > v1
	})
		decreases 100 - i7
	{
		if (i7 >= 100) {
			break;
		}
		
		i7 := i7 + 1;
		var v84: array<bool> := new bool[14];
		var v85: seq<bool> := [v4];
		v84[safeIndex(798, v84.Length)] := v85 != v85;
		var v86 := DC0(v4);
		var v87: C0 := new C0(v1, 0x2cb, v86.cf0);
		var v88: map<C0, bool> := map[v87 := v4];
		var v89: multiset<seq<bool>> := multiset{[!v4, v4], v85};
		v84[safeIndex(798, v84.Length)] := if (v87 in v88) then v88[v87] else v89 >= v89;
		var v90: array<int> := new int[16];
		v90[safeIndex(384, v90.Length)] := v87.f9;
		var v91: map<bool, bool> := map[v4 := v84[safeIndex(798, v84.Length)]];
		var v92: set<bool> := {v4, v86.cf0, v91 != v91, fm3(v87.f9, globalState), v84[safeIndex(798, v84.Length)] in v85};
		v84[safeIndex(798, v84.Length)], globalState.f6, v90[safeIndex(384, v90.Length)] := v4, v87.f9, |v92|;
		var v93: array<C0> := new C0[11];
		v93 := new C0[27];
		var v94: multiset<bool> := multiset{!false, !!!v4, v4, !true, v84[safeIndex(798, v84.Length)]};
		var v95 := m0(-v1, v94, v85, globalState);
	}
	v4 := v4;
	var v96 := DC0(fm3(v1, globalState));
	if (v96.cf0) {
		var v97: array<bool> := new bool[18];
		v97[safeIndex(29, v97.Length)] := v4;
		var v98: map<bool, int> := map[v4 := v1];
		var v99: array<int> := new int[4] [-v1, v1, if (v4 in v98) then v98[v4] else |map[v97 := v1]|, v1];
		v99[safeIndex(938, v99.Length)] := 0x2e0;
		var v100: seq<bool> := [v4, v4];
		var v101: multiset<seq<bool>> := multiset{v100};
		var v102 := DC1(v101, v1);
		var v103: map<seq<int>, int> := map[v82 := v1];
		globalState.f6, v97[safeIndex(29, v97.Length)], v99[safeIndex(938, v99.Length)] := v102.cf2, fm7(globalState) !in v103, v1;
		var v104: array<map<int, bool>> := new map<int, bool>[24];
		var v105: array<array<map<int, bool>>> := new array<map<int, bool>>[4] [v104, v104, v104, v104];
		var v106: map<int, array<map<int, bool>>> := map[v99[safeIndex(938, v99.Length)] := v104];
		var v107: multiset<bool> := multiset{false, v97[safeIndex(29, v97.Length)]};
		v105[safeIndex(718, v105.Length)] := if (|v107| in v106) then v106[|v107|] else v104;
		v4, v105[safeIndex(718, v105.Length)] := v100[safeIndex(0x10c, |v100|)], v104;
		var v108 := new C0(v99[safeIndex(938, v99.Length)], v1, v4);
		var v109 := DC4();
		var v110: set<D1> := {v109};
		v99[safeIndex(938, v99.Length)], v97[safeIndex(29, v97.Length)] := safeModuloInt(safeModuloInt(v108.f9, |v110|), v108.f9), multiset(v81 + v82) !! v6;
		if (v97[safeIndex(29, v97.Length)]) {
			var v111: map<bool, bool> := map[false := v97[safeIndex(29, v97.Length)]];
			var v112: map<map<bool, bool>, char> := map[v111 := v2];
			v112 := v112[map[v97[safeIndex(29, v97.Length)] := v4] := v2];
			var v113: T0 := new C0(v108.f9, v99[safeIndex(938, v99.Length)], v4);
			var v114: map<T0, bool> := map[v113 := v97[safeIndex(29, v97.Length)]];
			v97[safeIndex(29, v97.Length)] := if (v113 in v114) then v114[v113] else true;
			var v115 := m0(0x38b, multiset{v4, v97[safeIndex(29, v97.Length)]}, v100, globalState);
			globalState.f6, v99, v111, globalState.f3, v108 := v108.f8, v99, v111, v108.f9, v108;
			v99[safeIndex(938, v99.Length)] := v108.f8;
		} else {
			v97[safeIndex(29, v97.Length)], v99[safeIndex(938, v99.Length)] := v4, v108.f9;
			var v116: map<bool, bool> := map[v97[safeIndex(29, v97.Length)] := v4];
			globalState.f3 := safeModuloInt(|v116|, v82[safeIndex(|multiset{v4, v4}|, |v82|)]) * fm0(v99[safeIndex(938, v99.Length)], v97[safeIndex(29, v97.Length)], [v108.f9, v108.f8, v99[safeIndex(938, v99.Length)]], v2, globalState);
			var v117: seq<string> := [v0];
			v97[safeIndex(29, v97.Length)], v83, v97[safeIndex(29, v97.Length)] := v108.f8 < v108.fm5(fm8(v108.f9, v99[safeIndex(938, v99.Length)], 0x164, v99[safeIndex(938, v99.Length)], globalState), v2, v1, v4, globalState), if (v117[safeIndex(v108.f9, |v117|)] != v0) then if (v97[safeIndex(29, v97.Length)]) then v83 else v83 else fm9(globalState), v4 <== v97[safeIndex(29, v97.Length)];
			v116 := v116[v97[safeIndex(29, v97.Length)] := v97[safeIndex(29, v97.Length)]];
			var v118 := m0(v108.f8, multiset{v97[safeIndex(29, v97.Length)]}, v100, globalState);
		}
		
	} else {
		var v119: seq<bool> := [v4];
		var v120: multiset<seq<bool>> := multiset{v119, v119};
		var v121 := DC1(v120, v1);
		var v122: multiset<D0> := multiset{DC1(v120, |v119|), v121};
		v4 := !(v122 > multiset([v121]));
		var v123: multiset<bool> := multiset{v4, v4};
		var v124 := m0(|v0|, v123, v119[safeIndex(|[v4]|, |v119|) := v4], globalState);
		var v125: map<int, bool> := map[v1 := false];
		v4 := !(fm3(fm0(v1, v4, v81, v2, globalState), globalState) && (if (v1 in v125) then v125[v1] else v4));
		var v126 := m0(v1, v123 - v123, fm6(v1, v2, v4, v4, globalState), globalState);
		var v127: map<bool, int> := map[v4 := -v1];
		var v128: set<int> := {if (!v4 in v127) then v127[!v4] else v1};
		v128 := v128;
	}
	
	var v129: multiset<bool> := multiset{v4, v4, v4, v4, fm3(v1, globalState)};
	var v130: map<char, multiset<bool>> := map[v2 := v129];
	v130 := v130[v2 := v129];
	v4 := v4;
	v4 := v96.cf0;
	var v132: C0 := new C0(v1, |(set v131 : int | (0x1a0 <= v131) && (v131 < 0x2f2) :: (v131 * v1))|, v4);
	var v133 := DC6(v132);
	var v134: array<C0> := new C0[29] [v132, v132, v133.cf6, v132, v132, v132, if (false) then v132 else v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, v132, if (v4) then v132 else v132, if (v4) then v132 else v132, v132, v132];
	v134[safeIndex(5, v134.Length)] := v132;
	var v135: array<D0> := new D0[4](i8 => DC1(multiset{[v4, v4], DC9([v4, v4]).cf10, [DC10(v2, v4).cf12]}, v132.f9));
	var v136: seq<bool> := [v4];
	var v137: multiset<seq<bool>> := multiset{v136};
	var v138 := DC1(v137, v132.f9);
	v135[safeIndex(568, v135.Length)] := v138;
	var v139: array<int> := new int[16];
	var v140: seq<array<int>> := [v139];
	var v141: multiset<seq<int>> := multiset{v81};
	globalState.f6, v134[safeIndex(5, v134.Length)], v135[safeIndex(568, v135.Length)], v0 := |(v140 + [v139])|, v132, v138, fm10(|v136|, !(v132.f8 <= v132.f8), v1, v141 - v141, globalState);
	var v142 := DC10(v2, v4);
	match (v142) {
		case DC10(cf11, cf12) =>
			var v143: multiset<map<bool, int>> := multiset{map[v4 := v132.f8]};
			cf11 := fm11(v143, |v6| + v1, v0, v0, globalState);
			v4 := v4;
			var v144: map<bool, int> := map[v4 := v1];
			var v145: map<array<int>, map<bool, int>> := map[v139 := v144];
			var v146: array<char> := new char[8] [v2, fm11(multiset{v144, if (v139 in v145) then v145[v139] else v144}, v1, "unvm", v0, globalState), 'l', cf11, 'a', cf11, cf11, cf11];
			v146[safeIndex(842, v146.Length)] := v2;
			v146[safeIndex(842, v146.Length)] := cf11;
			match (v96) {
				case DC1(cf1, cf2) =>
					globalState.f3 := v132.f9;
					cf2 := v132.f9 + -v1;
					var v147: map<map<int, int>, seq<int>> := map[v40 := v81];
					var v148: set<int> := {|v82|};
					v147 := v147[v40[v132.f9 := |v148|] := v82];
					globalState.f6 := v132.f8;
				case DC0(cf0) =>
					var v149 := m0(safeModuloInt(v132.f9, v1), v129, v136, globalState);
					globalState.f6 := safeModuloInt(v1, v132.f9);
					var v150 := m0(v132.f9, v129, v136, globalState);
					globalState.f6 := v132.f8;
				case DC2(cf3) =>
					cf11 := cf11;
					var v151: map<int, map<array<int>, int>> := map[v132.f8 := map[v139 := v1]];
					var v152: map<array<int>, int> := map[v139 := v1];
					v151 := v151[v132.f9 := v152];
					var v153 := m0(v132.f9, fm12(v132.f8, |(seq(abs(546), i9  => (v146[safeIndex(842, v146.Length)])))|, fm13(v132.f8, globalState), globalState), v136, globalState);
					globalState.f3 := |v153|;
			}
			
		case DC11(cf13, cf14, cf15, cf16) =>
			var v154: array<bool> := new bool[1](i10 => cf16 && cf16);
			v154[safeIndex(526, v154.Length)] := v4;
			v154[safeIndex(526, v154.Length)] := !(0x186 >= (v132.f8 * v1));
			var v155 := DC12(v139);
			var v156: array<array<int>> := new array<int>[12] [v139, v139, v139, v139, v139, v155.cf17, v139, if (v4) then v139 else v139, v139, v139, v139, v139];
			v156[safeIndex(603, v156.Length)] := v139;
			v156[safeIndex(603, v156.Length)], cf13 := v139, v132.f8;
			var v157: map<bool, bool> := map[!true := fm3(v1, globalState)];
			var v158: set<map<int, int>> := {v40, map[|v5| := |v157|], v40};
			if (!(v158 !! v158)) {
				globalState.f3 := if (cf16) then v132.f9 else v132.f8;
				var v159: map<array<bool>, bool> := map[v154 := v154[safeIndex(526, v154.Length)]];
				v134[safeIndex(5, v134.Length)] := new C0(v1, |v159|, cf15);
				v0 := (v0 + (seq(abs(-203), i11  => ('b')))) + ("uw" + v0);
				var v160: array<T0> := new T0[14];
				var v161: T0 := new C0(v132.f8, -634, v154[safeIndex(526, v154.Length)]);
				v160[safeIndex(585, v160.Length)] := v161;
				v160[safeIndex(585, v160.Length)] := v161;
				v0 := v0;
			} else {
				var v162: array<D1> := new D1[1];
				v162[safeIndex(240, v162.Length)] := v83;
				v162[safeIndex(240, v162.Length)] := DC3(v82);
				globalState.f6 := (v132.f8 + v132.f9) - safeModuloInt(v132.f9, cf13);
				var v163: map<int, bool> := map[v132.f9 := v154[safeIndex(526, v154.Length)]];
				v163 := v163[v132.f9 := cf15];
				cf15 := cf15;
				v154[safeIndex(526, v154.Length)] := v154[safeIndex(526, v154.Length)];
			}
			
			var v164: set<int> := {v132.f9, -v132.f8};
			if ((set v165 : int | v165 in v164 :: (v165 + |[true, true]|)) == (v164 - v164)) {
				var v166 := new C0(v1, v132.f8, v154[safeIndex(526, v154.Length)]);
				v1 := (0x26c * v132.f8) + v138.cf2;
				var v167 := DC11(|v157|, 0x282, fm3(v166.fm4(globalState), globalState), v154[safeIndex(526, v154.Length)]);
				var v168: seq<D3> := [v167];
				v168 := v168;
				cf15 := !(if (cf16) then v154[safeIndex(526, v154.Length)] else cf15) && true;
				var v169: map<bool, int> := map[false := v166.f9];
				v154[safeIndex(526, v154.Length)] := if (v136[safeIndex(cf14, |v136|)]) then |v169| == v166.f9 else cf15;
			} else {
				var v170: map<int, string> := map[cf13 := v0[safeIndex(cf14, |v0|) := v2]];
				var v171 := m0(v132.f9, v129, fm6(|(if (v132.f8 in v170) then v170[v132.f8] else v0)|, v2, !false, fm3(cf13, globalState), globalState), globalState);
				v129, globalState.f3 := v129, v1 + |[true, v4, v154[safeIndex(526, v154.Length)]]|;
				v154[safeIndex(526, v154.Length)] := cf15;
				var v172: array<array<bool>> := new array<bool>[12];
				v172[safeIndex(306, v172.Length)] := v154;
				globalState.f3, v172[safeIndex(306, v172.Length)], v129 := cf14, v154, v129;
				var v173 := m0(v132.fm5(v157, v2, 0xd5, v154[safeIndex(526, v154.Length)], globalState), v129, [v4], globalState);
			}
			
		case DC9(cf10) =>
			if (!!!(v4 <==> v4)) {
				var v174: array<bool> := new bool[20];
				v174[safeIndex(369, v174.Length)] := !v4 <==> v4;
				v174[safeIndex(369, v174.Length)] := |v40| != v132.f8;
				globalState.f6 := |(v0 + v0)|;
				var v175: array<T0> := new T0[1];
				var v176: T0 := new C0(-v132.f8, v132.f8, v4);
				v175[safeIndex(6, v175.Length)] := v176;
				v175[safeIndex(6, v175.Length)] := v176;
				var v177 := m0(v132.f8, v129, cf10, globalState);
				v174[safeIndex(369, v174.Length)] := fm3(v132.f8 * v132.f9, globalState);
			} else {
				var v178: seq<C0> := [v132, v132];
				var v179: T0 := new C0(|v0|, |v178|, v4);
				var v180: seq<T0> := [v179];
				var v181: multiset<T0> := multiset{v179};
				var v182: array<T0> := new T0[23] [v179, v179, v179, v179, v180[safeIndex(|v181[v179 := abs(v132.f8)]|, |v180|)], v179, v179, v179, v179, v179, v179, v179, v179, v179, v179, v179, v179, v179, v179, if (true) then v179 else v179, v179, v179, v179];
				var v183 := DC15(v179);
				var v184: map<int, T0> := map[0x319 := v179];
				v182 := new T0[20] [v179, v179, v179, v179, v179, v179, v179, v179, v179, v179, v179, v179, v183.cf23, v179, v179, v179, v179, if (v132.f8 in v184) then v184[v132.f8] else v179, v179, v179];
				var v185 := m0(v1, v129, v136, globalState);
				var v186: array<bool> := new bool[26](i12 => true);
				var v187: map<C0, int> := map[v132 := |map[0x90 := v179]|];
				v186, v4, globalState.f6, v139, v1 := v186, v179.f7, v1 * (v132.f9 - -|[v187]|), v139, v132.f8 * -v132.f8;
				var v188: multiset<array<bool>> := multiset{v186};
				var v189 := DC11(v132.f9, v179.fm4(globalState), v188 >= v188, true);
				v189 := v189;
				cf10 := cf10 + v136;
			}
			
			v4 := !v4 ==> (v132.f9 == v132.f9);
			var v190 := m0(v132.f9, v129, cf10, globalState);
			v4, globalState.f6 := v4, v132.f8;
	}
	
	v4 := v4;
	v4, v4, globalState.f3, globalState.f3 := v136[safeIndex(if (v4) then v132.f8 else v132.f9, |v136|)], true, -(v1 + v1), v132.f9;
	if (v4 ==> true) {
		var v191 := m0(v1, multiset(v136), v136 + [v4, v4, v4], globalState);
		v4 := true;
		v4 := v0 != v191;
		var v192: array<bool> := new bool[18];
		var v193 := DC3(v82[safeIndex(|map[v192 := v1]|, |v82|) := v132.f8]);
		var v194 := DC5(v193);
		match (v194) {
			case DC4() =>
				var v195: array<seq<int>> := new seq<int>[11];
				v195[safeIndex(719, v195.Length)] := v81;
				v195[safeIndex(719, v195.Length)] := v81;
				var v196: set<int> := {-0x3ab, v132.f8};
				var v197 := DC16(v196);
				globalState.f3 := |v197.cf24| + v1;
				var v198: map<int, bool> := map[v132.f8 := false];
				v198 := v198[v1 := false];
				var v199: set<bool> := {v4};
				var v200: map<int, set<bool>> := map[safeDivisionInt(|v136[safeIndex(v132.f8, |v136|) := v4]|, |v0|) := v199];
				v200 := v200[safeModuloInt(-v82[safeIndex(v1, |v82|)], -0x318) := v199 - v199];
			case DC3(cf4) =>
				v4 := !v4;
				var v201: T0 := new C0(0x3c5, v132.f8, fm3(v1, globalState));
				var v202: set<bool> := {v201.f7, v201.f7, v201.f7, false, v201.f7};
				v201 := if (v132.f9 != fm0(|v202|, true, [v132.f9, v1, v132.f9], v2, globalState)) then v201 else v201;
				var v203: map<bool, int> := map[v201.f7 := -v1];
				var v204 := new C0(|(v203 + v203)|, v132.f8, false);
				globalState.f3 := (v204.fm4(globalState) - v132.f8) + v132.f9;
			case DC5(cf5) =>
				v134[safeIndex(5, v134.Length)] := if (v132.f9 == v132.f8) then v132 else v132;
				var v205: map<int, array<bool>> := map[v132.f9 := v192];
				v205 := v205[if (v4) then v132.f9 else v132.f9 := v192];
				v4 := v37 >= fm14(globalState);
				v139[safeIndex(748, v139.Length)] := v132.f8;
				v139[safeIndex(748, v139.Length)] := v132.f9;
		}
		
		v4 := v4;
	} else {
		globalState.f3 := v1;
		var v206: map<set<bool>, array<int>> := map[{!v4} := v139];
		var v207: set<bool> := {v4};
		v206 := v206[v207 := v139];
		globalState.f3 := v132.f8;
		if (v4) {
			var v208 := new C0(v132.f8 * v132.f9, v1, v4);
			var v209: multiset<C0> := multiset{v132, v132};
			var v210: map<multiset<C0>, int> := map[v209 := |{true, v4, v4}|];
			v4 := !((v209 * v209) !in v210);
			globalState.f6 := 0x41;
			var v212: set<int> := {|v0|, safeModuloInt(|(set v211 : int | v211 in v81 :: (v211 * |map[[false] := 0x3a3]|))|, v132.f8)};
			v212 := (v212 * v212) * v212;
			v2 := v2;
		} else {
			v0 := fm10(v132.f8, v4, |(multiset{v4, v4})[fm3(v132.f8, globalState) := abs(v132.f8)]|, v141, globalState) + v0;
			v132 := v134[safeIndex(5, v134.Length)];
			v139[safeIndex(54, v139.Length)] := -v132.f8;
			v139[safeIndex(54, v139.Length)] := v1 - v132.f9;
			var v213: map<bool, set<bool>> := map[!v4 := {v4}];
			var v214 := m0(|(if (v4 in v213) then v213[v4] else v207)|, v129, v136, globalState);
			var v215 := new C0(v132.f8, safeModuloInt(v132.f9, v132.f8), v139[safeIndex(54, v139.Length)] >= v132.f8);
		}
		
		var v216: map<bool, bool> := map[v4 := v4];
		var v217: map<bool, bool> := map[v4 := !!v136[safeIndex(|v216|, |v136|)]];
		var v218 := DC9((v136[safeIndex(v132.f9, |v136|) := false])[safeIndex(v132.fm5(v216, v2, 0x1d8, v4, globalState), |v136[safeIndex(v132.f9, |v136|) := false]|) := v4]);
		var v219 := m0(safeDivisionInt(v132.f8, v1), multiset{v4, v4}, v218.cf10, globalState);
	}
	
	v1 := v1 - |v141|;
	print v0, "\n";
	print v1, "\n";
	print v2, "\n";
	print v3 == map[779 := 'p'], "\n";
	print v4, "\n";
	print v5 == map['p' := false], "\n";
	print v6 == multiset{-166}, "\n";
	print v7 == multiset{multiset{1}}, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4 == map['p' := false], "\n";
	print globalState.f5 == multiset{multiset{1}}, "\n";
	print globalState.f6, "\n";
	print v37 == {'p'}, "\n";
	print v38 == {{'p'}}, "\n";
	print v39 == [{{'p'}}], "\n";
	print v40 == map[778 := 778], "\n";
	print v81 == [2], "\n";
	print v82 == [2, 2], "\n";
	print v83.cf4 == [2], "\n";
	print i7, "\n";
	print v96.cf0, "\n";
	print v129 == multiset{true, true, true, true, false}, "\n";
	print v130 == map['p' := multiset{true, true, true, true, false}], "\n";
	print v132.f8, "\n";
	print v132.f9, "\n";
	print v133.cf6.f8, "\n";
	print v133.cf6.f9, "\n";
	print v133.cf6.f7, "\n";
	print v134[0].f8, "\n";
	print v134[0].f9, "\n";
	print v134[0].f7, "\n";
	print v134[1].f8, "\n";
	print v134[1].f9, "\n";
	print v134[1].f7, "\n";
	print v134[2].f8, "\n";
	print v134[2].f9, "\n";
	print v134[2].f7, "\n";
	print v134[3].f8, "\n";
	print v134[3].f9, "\n";
	print v134[3].f7, "\n";
	print v134[4].f8, "\n";
	print v134[4].f9, "\n";
	print v134[4].f7, "\n";
	print v134[5].f8, "\n";
	print v134[5].f9, "\n";
	print v134[5].f7, "\n";
	print v134[6].f8, "\n";
	print v134[6].f9, "\n";
	print v134[6].f7, "\n";
	print v134[7].f8, "\n";
	print v134[7].f9, "\n";
	print v134[7].f7, "\n";
	print v134[8].f8, "\n";
	print v134[8].f9, "\n";
	print v134[8].f7, "\n";
	print v134[9].f8, "\n";
	print v134[9].f9, "\n";
	print v134[9].f7, "\n";
	print v134[10].f8, "\n";
	print v134[10].f9, "\n";
	print v134[10].f7, "\n";
	print v134[11].f8, "\n";
	print v134[11].f9, "\n";
	print v134[11].f7, "\n";
	print v134[12].f8, "\n";
	print v134[12].f9, "\n";
	print v134[12].f7, "\n";
	print v134[13].f8, "\n";
	print v134[13].f9, "\n";
	print v134[13].f7, "\n";
	print v134[14].f8, "\n";
	print v134[14].f9, "\n";
	print v134[14].f7, "\n";
	print v134[15].f8, "\n";
	print v134[15].f9, "\n";
	print v134[15].f7, "\n";
	print v134[16].f8, "\n";
	print v134[16].f9, "\n";
	print v134[16].f7, "\n";
	print v134[17].f8, "\n";
	print v134[17].f9, "\n";
	print v134[17].f7, "\n";
	print v134[18].f8, "\n";
	print v134[18].f9, "\n";
	print v134[18].f7, "\n";
	print v134[19].f8, "\n";
	print v134[19].f9, "\n";
	print v134[19].f7, "\n";
	print v134[20].f8, "\n";
	print v134[20].f9, "\n";
	print v134[20].f7, "\n";
	print v134[21].f8, "\n";
	print v134[21].f9, "\n";
	print v134[21].f7, "\n";
	print v134[22].f8, "\n";
	print v134[22].f9, "\n";
	print v134[22].f7, "\n";
	print v134[23].f8, "\n";
	print v134[23].f9, "\n";
	print v134[23].f7, "\n";
	print v134[24].f8, "\n";
	print v134[24].f9, "\n";
	print v134[24].f7, "\n";
	print v134[25].f8, "\n";
	print v134[25].f9, "\n";
	print v134[25].f7, "\n";
	print v134[26].f8, "\n";
	print v134[26].f9, "\n";
	print v134[26].f7, "\n";
	print v134[27].f8, "\n";
	print v134[27].f9, "\n";
	print v134[27].f7, "\n";
	print v134[28].f8, "\n";
	print v134[28].f9, "\n";
	print v134[28].f7, "\n";
	print v135[0].cf1 == multiset{[false]}, "\n";
	print v135[0].cf2, "\n";
	print v135[1].cf1 == multiset{[false, false], [false, false], [false]}, "\n";
	print v135[1].cf2, "\n";
	print v135[2].cf1 == multiset{[false, false], [false, false], [false]}, "\n";
	print v135[2].cf2, "\n";
	print v135[3].cf1 == multiset{[false, false], [false, false], [false]}, "\n";
	print v135[3].cf2, "\n";
	print v136 == [false], "\n";
	print v137 == multiset{[false]}, "\n";
	print v138.cf1 == multiset{[false]}, "\n";
	print v138.cf2, "\n";
	print v139[12], "\n";
	print |v140|, "\n";
	print v141 == multiset{[2]}, "\n";
	print v142.cf11, "\n";
	print v142.cf12, "\n";
}