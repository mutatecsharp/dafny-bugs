datatype D0 = DC0(cf0: bool)
datatype D1 = DC2(cf2: bool, cf3: bool, cf4: bool, cf5: char, cf6: char) | DC1(cf1: set<bool>) | DC3(cf7: D1)
datatype D2 = DC5(cf9: int, cf10: int, cf11: int, cf12: int, cf13: seq<bool>) | DC4(cf8: int)
datatype D3 = DC7(cf15: int, cf16: int, cf17: C0) | DC8(cf18: int, cf19: int) | DC9(cf20: int, cf21: int, cf22: C0, cf23: int) | DC6(cf14: array<bool>)
datatype D4 = DC11(cf25: bool, cf26: int, cf27: int, cf28: int) | DC10(cf24: array<int>)
datatype D5 = DC13(cf30: int, cf31: bool, cf32: int, cf33: array<string>, cf34: bool) | DC14 | DC12(cf29: seq<int>)
datatype D6 = DC16(cf36: int, cf37: bool, cf38: bool) | DC15(cf35: multiset<int>)
datatype D7 = DC18(cf40: int, cf41: string, cf42: bool, cf43: bool, cf44: set<int>) | DC17(cf39: map<D5, bool>)
datatype D8 = DC20(cf46: bool, cf47: bool, cf48: bool) | DC21(cf49: bool, cf50: array<array<bool>>, cf51: int, cf52: array<bool>, cf53: int) | DC19(cf45: multiset<bool>) | DC22(cf54: D8)
datatype D9 = DC24(cf56: string, cf57: bool, cf58: int) | DC25(cf59: string, cf60: bool, cf61: int) | DC26(cf62: int) | DC23(cf55: map<bool, array<C0>>)
datatype D10 = DC28(cf64: C2, cf65: bool, cf66: bool, cf67: bool) | DC29(cf68: bool) | DC27(cf63: map<bool, int>)
datatype D11 = DC30(cf69: map<bool, bool>)
datatype D12 = DC32(cf71: int, cf72: C2, cf73: C0, cf74: int) | DC31(cf70: seq<set<D5>>) | DC33(cf75: D12)
datatype D13 = DC34(cf76: map<seq<int>, int>)
datatype D14 = DC35(cf77: multiset<char>)
datatype D15 = DC37(cf79: seq<bool>, cf80: bool, cf81: bool, cf82: char) | DC36(cf78: map<char, int>)
class GlobalState {
	const f0 : map<map<int, bool>, char>
	var f1 : array<int>
	const f2 : bool
	var f3 : bool
	var f4 : array<int>
	const f5 : int
	var f6 : seq<bool>
	var f7 : bool
	var f8 : bool
	var f9 : int
	const f10 : bool
	constructor (f0 : map<map<int, bool>, char>, f1 : array<int>, f2 : bool, f3 : bool, f4 : array<int>, f5 : int, f6 : seq<bool>, f7 : bool, f8 : bool, f9 : int, f10 : bool) {
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

function fm0(p0: char, globalState: GlobalState): bool {
	(0x2fb % 0x126) > |((seq(934, i0  => ('a'))) + "goxbqw")|
}
function fm1(p0: bool, p1: int, p2: int, globalState: GlobalState): int {
	if (true <==> true) then |((seq(882, i0  => ('y'))) + "ov")| else |(['f', 'v', 'd', 'y'] + ['k', 'p', 'w', 'p'])|
}
function fm5(p0: bool, p1: bool, p2: bool, p3: int, globalState: GlobalState): seq<bool> {
	[|{|map[|map[710 := |map[0x330 := 'l']|]| := true]|}| < |"irm"|, DC25("ap", false, 0xd6).cf60, false]
}
function fm7(globalState: GlobalState): seq<bool> {
	if (!(|map[true := "sfmw"]| >= 0x132)) then [true, false] else [true, !!true]
}
function fm9(p0: int, p1: bool, globalState: GlobalState): string {
	"x"
}
function fm10(p0: bool, p1: bool, p2: bool, globalState: GlobalState): seq<bool> {
	if (true || false) then [false] else [true]
}
function fm13(p0: char, p1: bool, globalState: GlobalState): multiset<int> {
	(multiset{|(map v0 : int | (-834 <= v0) && (v0 < -0x308) :: (v0 % |multiset{true}|) := (54))|, |map[false := -|map[false := |map[true := |[false]|]|]|]|, |map[[false] := false]|, |{!DC20(true, false, false).cf48}|} * multiset{|{false, true}|, 0x0, -|[|map[DC18(|map[|[true]| := 0x228]|, "dvkfi", true, !false, set v1 : int | v1 in {-0x2ab} :: (v1 / |"mqnecwgoe"|)).cf43 := 'r']|]|, |[0x3b7, |map[-0x35f := true]|, 159, |map[!true := true]|, |[false]|]|}) + (multiset([-0x2a5]) * multiset{|[|[true]|, 0x35d]|, |"od"|, |[0x1]|})
}
function fm14(p0: D1, p1: int, globalState: GlobalState): char {
	'n'
}
function fm15(p0: int, p1: int, p2: bool, globalState: GlobalState): D1 {
	DC2(DC25("klphfov", !true, -|map[|"arkhat"| := false]|).cf60, true, false, if (true) then 'k' else 'g', 'k')
}
function fm16(p0: int, globalState: GlobalState): map<bool, bool> {
	map[true := true] + map[!!false := true]
}
function fm17(p0: bool, p1: bool, p2: int, p3: int, globalState: GlobalState): multiset<set<D1>> {
	multiset(seq(172, i0  => ({DC3(DC1({true}))})))
}
function fm18(p0: int, p1: bool, globalState: GlobalState): set<bool> {
	({false, false, !true, true, false} + {true, !false}) + {false, true}
}
function fm19(p0: int, globalState: GlobalState): multiset<map<bool, int>> {
	multiset{DC27(map[true := 0x8f]).cf63, map[false := 0xb0] + map[true := 0x2b6]}
}
function fm20(p0: char, p1: int, p2: int, p3: map<int, char>, globalState: GlobalState): set<int> {
	(set v0 : int | (0x202 <= v0) && (v0 < -0xec) :: (v0 + |(seq(-0xb7, i0  => (|map[!true := |(map v1 : int | (740 <= v1) && (v1 < 0x12f) :: (v1 + |multiset{0x116}|) := (|"wag"|))|]|)))|)) - ((set v2 : int | (-0x2e <= v2) && (v2 < 0x35) :: (v2 * 370)) + {-604, -|{DC2(!true, true, false, 'g', 'i').cf2}|, 0x326})
}
function fm21(p0: bool, p1: bool, p2: bool, globalState: GlobalState): map<char, int> {
	DC36(map['o' := 123]).cf78
}
function fm22(p0: int, p1: bool, globalState: GlobalState): multiset<bool> {
	multiset{492 <= 754}
}
function fm23(p0: int, globalState: GlobalState): map<bool, int> {
	map[true := |map[|"oot"| := false]|]
}
function fm24(globalState: GlobalState): seq<int> {
	[120, -(|{|map[false := false]|}| % 0x2c8), -0x32d, if (!true) then 0x3ad else |"ldpqylslw"|]
}
function fm25(p0: char, p1: char, p2: bool, globalState: GlobalState): map<string, string> {
	((map v0 : string | v0 in {"aua", seq(0x42, i0  => ('c'))} :: (v0) := ("jhdbecap")) + map[seq(-0x41, i1  => ('t')) := seq(0x3e1, i2  => ('o'))]) + map["iqvwwh" := "oudm"]
}
function fm26(p0: int, p1: bool, globalState: GlobalState): D8 {
	DC20(if (false) then !true else true, true, "qoxnse" <= "rjxfs")
}
function fm27(p0: D14, p1: string, globalState: GlobalState): set<seq<int>> {
	{[|[0x18c]|] + [0x19a, 0xa6]}
}
function fm28(globalState: GlobalState): D0 {
	if (false) then DC0(true) else DC0(true)
}
function fm29(p0: bool, p1: int, globalState: GlobalState): map<int, map<bool, int>> {
	map[-0x225 := map[false := -0xdc]] + map[-0x1c3 := map[true := 0x21c]]
}
method m0(globalState: GlobalState) {
	var v0 := false;
	globalState.f3 := v0;
	var v1 := 'f';
	var v2: map<bool, char> := map[v0 := v1];
	var v3 := 692;
	v2 := v2[v3 >= 0x186 := v1];
	var v4: seq<int> := [v3, 164];
	var i0 := 0;
	while (v3 <= (|[|v4|]| % v3))
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		globalState.f9 := v3;
		if (!!!!v0) {
			var v5: array<bool> := new bool[18](i1 => v3 != v3);
			var v6 := "luog";
			v5[830] := v6 != v6;
			v5[830] := fm0(v1, globalState);
			var v7: array<int> := new int[27];
			globalState.f3, v5[830] := {v7} > {v7}, if (v5[830]) then true else v5[830];
			globalState.f7 := false;
			globalState.f9 := v3;
			globalState.f9 := v3;
		} else {
			var v8: multiset<bool> := multiset{v0, false};
			var v9: seq<bool> := [v0, v0];
			var v10: set<bool> := {v0, v0};
			var v11: map<set<int>, int> := map[{130, v3} := fm1(v0, |v10|, v3, globalState)];
			var v12: set<int> := {v3};
			var v13 := "jlgqgkipu";
			v3 := (if (v9[v3] in v8) then v8[v9[v3]] else v3) + (if (v12 in v11) then v11[v12] else |v13|);
			var v14 := new C2(v0, v8, fm0(v13[v3], globalState));
			globalState.f9, v3, v3, v0 := v3 - (if (v14.f15) then |v13| else v3), -264, -v3, v0;
			var v15: multiset<int> := multiset{v3, v3, v3, v3, |v13|};
			globalState.f9 := (--(if (v3 in v15) then v15[v3] else |v4|) % |{v14.f15}|) - v3;
			globalState.f7 := v14.f15;
		}
		
		var v16: map<bool, int> := map[v0 := v3];
		var v17 := DC27(v16);
		match (v17.(cf63 := map[!v0 := v3] + v16)) {
			case DC28(cf64, cf65, cf66, cf67) =>
				var v18: seq<bool> := [cf66];
				var v19: map<char, seq<bool>> := map[v1 := v18];
				v19 := v19[v1 := [cf67, v0]];
				var v20: set<bool> := {v0};
				cf67 := (v20 - v20) > v20;
				var v21: set<seq<int>> := {fm24(globalState)};
				var v22: multiset<char> := multiset{v1};
				var v23 := DC35(v22);
				var v24 := "kjr";
				var v26: map<int, set<seq<int>>> := map[v3 := v21];
				var v28: map<seq<int>, char> := map[fm24(globalState) := v1];
				var v30: array<set<seq<int>>> := new set<seq<int>>[15] [v21, fm27(v23, v24, globalState), v21, {v4, v4}, fm27(DC35(v22).(cf77 := v22), v24, globalState), v21, v21, set v25 : seq<int> | v25 in (seq(-30, i2  => (v4))) :: (v25), v21, (if (v3 in v26) then v26[v3] else {v4, v4}) - {v4[v3 := v3]}, v21, set v27 : seq<int> | v27 in v21 :: (v27), v21, set v29 : seq<int> | v29 in v28 :: (v29), v21];
				v30[210] := fm27(DC35(v22), v24, globalState);
				v30[210] := v21;
				globalState.f9 := 0xa2;
			case DC29(cf68) =>
				v1 := v1;
				globalState.f9 := v4[v3 + v3];
				var v31: array<map<D5, int>> := new map<D5, int>[11](i3 => map[DC14() := v3] + map[DC14() := |multiset{v0, cf68}|]);
				var v32 := DC14();
				var v33: map<D5, int> := map[v32 := v3];
				v31[36] := v33;
				v31[36] := v33 + v33;
				var v34: seq<bool> := [true];
				var v35 := DC2(cf68, v0, cf68, v1, v1);
				var v36 := DC2(|v34| == v3, v0, cf68, fm14(v35, |v34|, globalState), v1);
				v35 := v35.(cf6 := v1, cf3 := v0).(cf3 := cf68, cf4 := false);
			case DC27(cf63) =>
				var v37: array<int> := new int[2];
				globalState.f1 := if (v0) then v37 else v37;
				var v38: multiset<bool> := multiset{v0};
				var v39: C2 := new C2(true, v38, v0);
				var v40: map<C2, int> := map[v39 := v3];
				var v41: set<int> := {if (v0) then |v40[v39 := v3]| else v3, v3 + v3};
				v41 := v41 + v41;
				globalState.f9 := v3 + v3;
				var v42 := "bdgl";
				v42 := v42;
		}
		
		var v43: multiset<bool> := multiset{v0, false, v0, false};
		var v44: T0 := new C2(v0, v43, v0);
		var v45: multiset<T0> := multiset{v44, v44};
		var v46: map<int, bool> := map[v3 := v44.f12];
		v3 := fm1(fm0(v1, globalState), -|v45|, |v46|, globalState);
	}
	var i4 := 0;
	while (v0)
		decreases 100 - i4
	{
		if (i4 >= 100) {
			break;
		}
		
		i4 := i4 + 1;
		var v47: seq<bool> := [v0 <==> v0, v0, v0];
		if (v47[v3]) {
			var v48: multiset<bool> := multiset{v0, v0};
			var v49 := new C2(v0, v48, false);
			v48 := if (!(v3 < v3)) then v48 else v48;
			globalState.f9 := v3;
			var v50 := "s";
			var v51: map<int, bool> := map[v3 := v49.f15];
			var v52: seq<map<int, bool>> := [v51];
			var v53: map<int, int> := map[-|(v50 + v50)[v3 := v1]| := v3 - |v52[v3]|];
			v53 := v53[v3 := v3];
			var v54: T0 := new C1(v3, v48, v49.f15);
			var v55: map<T0, seq<int>> := map[v54 := v4];
			var v56: set<int> := {v3, v3, |multiset(if (v54 in v55) then v55[v54] else v4)|, |{v1, v1}|};
			var v57: map<char, multiset<bool>> := map['a' := v48];
			var v58 := new C3(|multiset{v0, v0}|, v56 < v56, if ('y' in v57) then v57['y'] else v48, {v3} > {v3});
		} else {
			var v59 := DC0(true <== v0);
			v59 := fm28(globalState).(cf0 := v0 <== v0);
			globalState.f7 := v0;
			var v60: C0 := new C0(if (v0) then v0 else v0, v3, multiset{true, v0}, v3 !in v4);
			v60 := v60;
			var v61: multiset<bool> := multiset{v0};
			var v62 := new C2(v0, v61 + multiset{v60.f17}, v0);
			var v63 := "r";
			var v64: multiset<int> := multiset{fm1(v60.f17, v60.f18, |v2|, globalState)};
			var v65: map<bool, int> := map[v0 := v3];
			v63, globalState.f9, v3 := fm9(v3, v60.f17 ==> v60.fm2(v64, -v3, v65, v0, globalState), globalState), v3, (0x15e + v3) * v3;
		}
		
		var v66: array<bool> := new bool[7](i5 => if (false) then v0 else v0);
		v66[460] := !v0;
		var v67: map<bool, int> := map[v0 := v3];
		v66[460] := true in (v67 + v67);
		globalState.f6 := v47;
		var v68: multiset<char> := multiset{v1};
		var v69: seq<char> := [v1];
		v66[460] := v68 > multiset(v69);
	}
	var v70 := "dslsfkrf";
	var v71: map<string, int> := map[v70 := |v70|];
	var v72: set<bool> := {fm0(v1, globalState)};
	var v73: multiset<bool> := multiset{v0};
	var v74: array<bool> := new bool[26] [v0, v0, true <==> v0, v0, v0, v0, v71 != v71[v70 := v3], v0, v0 !in v72, false, v0, !v0, v72 !! v72, true, fm0(v1, globalState), true, if (v0) then v0 else false, v0, v0, v0, |v73| == |v70|, (DC0(v0).(cf0 := v0)).cf0, v0, v0, v0, v0];
	var v75: seq<array<bool>> := [v74, v74, if (v0) then v74 else v74];
	var v76: multiset<int> := multiset{v3, v3, v3, fm1(v0, v3, 496, globalState)};
	v74, v74, v0 := v75[|(seq(0x318, i6  => (v70[0x36d])))| % v3], v74, if (v70 < v70) then !v0 else v3 in v76;
	var v77: array<string> := new string[2];
	var v78 := DC13(v3, v0, v3, v77, v0);
	for i7 := v3 to v78.cf30 {
		var v79 := DC35(multiset{v1} * multiset{'g'});
		match (v79) {
			case DC35(cf77) =>
				var v80: seq<bool> := [true];
				var v81: seq<bool> := [v0, v0, false, v80[-|"jcvbfjwrl"|]];
				globalState.f9 := if (i7 in v76) then v3 else i7 + |v81|;
				var v82: array<int> := new int[1](i8 => i8 + v3);
				v82[618] := v3 * v3;
				var v83: multiset<array<int>> := multiset{v82, v82};
				var v84: set<int> := {i7, i7, v3};
				var v85 := DC5(-i7, |v83|, |v84|, |(multiset(v80))[v0 := fm1(v0, i7, v3, globalState)]|, v81);
				globalState.f4, v82[618], v85 := v82, fm1(v0, v3 * v3, v3, globalState), v85;
				v1 := if (true) then v1 else v1;
				var v86: map<char, int> := map[v1 := v3];
				v86 := v86[v1 := fm1(v0, v82[618], -v82[618], globalState)];
		}
		
		v0 := false;
		var v87: seq<seq<bool>> := [[v0]];
		var v88: map<seq<bool>, multiset<string>> := map[v87[v3] := multiset([seq(529, i9  => (v1)), v70, v70, "bultv", "yvdibk"] + [v70, v70])];
		var v89: seq<bool> := [v0];
		var v90: multiset<string> := multiset{v70, v70, v70};
		v88 := v88[v89 := v90];
		if (!!v0) {
			var v91 := new C4(multiset(v89) + v73, !!v0);
			globalState.f9 := i7 / i7;
			v70 := v70;
			var v92: set<int> := {i7};
			var v93 := DC8(--v3, i7);
			var v94: map<bool, D3> := map[|v92| in v92 := v93];
			v94 := v94[v0 := v93];
			var v95: map<string, string> := map["o" := "uli"];
			var v96: map<map<int, map<bool, int>>, int> := map[fm29(v89[|v92|], v3, globalState) := |{0x395, |v4|, v91.fm3(v0, v95, v0, i7, globalState)}|];
			var v97: map<bool, int> := map[v0 := i7];
			var v98: map<int, map<bool, int>> := map[i7 := v97];
			var v99: map<bool, bool> := map[v0 := v0];
			v96 := v96[v98 + map[v3 := v97] := |v99| * i7];
		} else {
			var v100: array<C0> := new C0[2];
			var v101 := DC23(map[v0 := v100]);
			var v102: map<int, map<bool, array<C0>>> := map[v3 := map[v0 := v100]];
			var v103: map<bool, array<C0>> := map[v0 := v100];
			var v104: set<D9> := {v101.(cf55 := if (v3 in v102) then v102[v3] else v103), v101, if (true) then v101 else DC23(v103), v101, v101};
			v104 := v104 + v104;
			v74[705] := v0;
			var v105: map<bool, bool> := map[!(if (v0) then v0 else !v0) := !(v0 && v0)];
			v74[705] := if ((if (!v0) then v0 else v0) in v105) then v105[if (!v0) then v0 else v0] else v0;
			var v106 := DC24(v70, v0, v3);
			var v107: array<bool> := new bool[20] [v74[705], v0, v74[705], false, v74[705], v106.cf57, v74[705], !v0, v89[v3], v74[705], v74[705], v74[705], v74[705], v74[705], v0, v0, v0, v0, v0, v0];
			var v108: set<array<bool>> := {v107, v107};
			var v109: map<int, int> := map[i7 := |v105|];
			var v110: map<bool, map<int, int>> := map[v0 := v109];
			var v111: map<bool, map<bool, map<int, int>>> := map[|v105| != |v108| := v110[v74[705] := map[v3 := |[0x131]|]] + v110];
			v111 := v111[if (v89[i7]) then v74[705] else v0 := v110 + map[v0 := v109]];
			var v112: C1 := new C1(0x12c, multiset{v74[705]}, v0);
			var v113: map<bool, int> := map[true := v3];
			var v114: array<int> := new int[10] [|v113|, v3, v3, |v89|, i7, 93, v4[v112.f16], v112.f16, -v3, |v4|];
			var v115: seq<array<int>> := [v114];
			var v116: array<array<bool>> := new array<bool>[17] [v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107];
			var v117 := DC21(!v0, v116, v3, v107, v3);
			v112, v107, globalState.f8 := v112, if (v115[-v112.f16 := v114] <= v115) then v117.cf52 else v107, v74[705];
			v114[553] := v112.f16;
			v114[553] := i7;
		}
		
	}
}
trait T0 {
	var f11 : multiset<bool>
	var f12 : bool
	function fm2(p0: multiset<int>, p1: int, p2: map<bool, int>, p3: bool, globalState: GlobalState): bool 
}

class C0 extends T0 {
	const f17 : bool
	const f18 : int
	constructor (f17 : bool, f18 : int, f11 : multiset<bool>, f12 : bool) {
		this.f17 := f17;
		this.f18 := f18;
		this.f11 := f11;
		this.f12 := f12;
	}
	
	function fm2(p0: multiset<int>, p1: int, p2: map<bool, int>, p3: bool, globalState: GlobalState): bool {
		f12
	}
	function fm11(p0: bool, p1: int, p2: int, globalState: GlobalState): map<int, D1> {
		map[f18 / f18 := DC2(f12, false, f17, 'f', 'x')]
	}
	function fm12(globalState: GlobalState): int {
		f18 / 0x123
	}
}

class C1 extends T0 {
	var f16 : int
	constructor (f16 : int, f11 : multiset<bool>, f12 : bool) {
		this.f16 := f16;
		this.f11 := f11;
		this.f12 := f12;
	}
	
	function fm2(p0: multiset<int>, p1: int, p2: map<bool, int>, p3: bool, globalState: GlobalState): bool {
		match DC5(f16, f16, f16, -0x1ca, [f12]) {
			case DC5(cf9, cf10, cf11, cf12, cf13) => f12
			case DC4(cf8) => !(if (f12) then f12 else false)
		}
	}
	method m6(globalState: GlobalState) returns (r0: int) {
		var i0 := 0;
		while (f12)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			if ((if (f12) then f16 else |map[f12 := f12]|) == (f16 / f16)) {
				var v0 := DC4(0x2b);
				globalState.f9 := v0.cf8 % f16;
				var v1: array<int> := new int[27](i1 => i1 + f16);
				var v2: seq<array<int>> := [v1];
				var v3: map<bool, int> := map[f12 := |(v2 + v2)[f16 := v2[f16]]|];
				v3 := v3 + v3;
				v1[524] := -f16;
				v1[524] := 0x114;
				var v4: array<bool> := new bool[3];
				v4 := v4;
				var v5: array<array<bool>> := new array<bool>[27];
				v5[596] := v4;
				v4[240] := false;
				var v6: array<string> := new string[22];
				var v7 := 't';
				var v8 := DC2(f12, f12, true, v7, v7);
				v6[789] := fm9(v1[524], v8.cf3, globalState);
				var v9 := "xcdmomb";
				var v10: multiset<string> := multiset{seq(0xf8, i2  => (v7)), "cqudam", v9};
				var v11: map<bool, bool> := map[f12 := f12];
				var v12: seq<int> := [|v11|, v1[524], |v9|, v1[524]];
				var v13: seq<string> := [v9];
				var v14: multiset<int> := multiset{f16};
				var v15: map<int, int> := map[v1[524] := f16];
				v5[596], f12, v4[240], v6[789] := v4, (v10 - v10) > v10, [|v11|] == v12, v9 + v13[fm1(fm2(v14, f16, map[f12 := -0x2fa], f12, globalState), v1[524], |v15|, globalState)];
			} else {
				var v16 := 't';
				var v17: seq<char> := [v16];
				var v18: seq<bool> := [f12, 'p' !in v17, f12, f12];
				globalState.f6 := v18[f16 * f16 := true];
				var v19: multiset<int> := multiset{|v17|};
				var v20: array<int> := new int[27];
				var v21: map<seq<bool>, array<int>> := map[v18 := v20];
				globalState.f9 := if (!v18[f16]) then if (|v21| in v19) then v19[|v21|] else 0x162 else -844;
				globalState.f3 := !f12;
				globalState.f9 := f16;
				var v22 := DC2(f12, f12, f12, v16, v16);
				globalState.f9 := -|map[if (f12) then f12 else f12 := v22.cf5]|;
			}
			
			var v23 := 'q';
			var v24: seq<bool> := [fm0(v23, globalState)];
			var v25 := "yfreelwss";
			var v26: array<bool> := new bool[23] [f12, if (f12) then f12 else false, f12, f12, f12, true, f12 <==> !f12, v24 == v24, f12, f12 <== f12, f12 || !f12, f12, f12, |map[f12 := f12]| > f16, f12, f12, |v25| > f16, !f12, f12, if (f12) then f12 else f12, f12, f12, f12];
			v26[19] := f12;
			var v27: map<bool, bool> := map[!f12 := v25 == v25];
			var v28: seq<seq<bool>> := [v24, v24, [f12, f12], v24, fm10(f12, f12, true, globalState)];
			var v29: multiset<int> := multiset{|v28[f16]|};
			var v30: map<bool, int> := map[f12 := f16];
			globalState.f9, v26[19] := f16 - |"ddcprsn"|, if (f12 in v27) then v27[f12] else fm2(v29, f16, v30, f12, globalState) <== f12;
			var v31: array<array<bool>> := new array<bool>[2];
			v31[811] := v26;
			v29, v25, v31[811], v26 := v29, ("nvme" + v25)[f16 - f16 := v23], v26, v26;
			var v32: multiset<char> := multiset{v23, v23, v23};
			var v33: seq<int> := [f16, f16];
			var v34: map<multiset<char>, seq<int>> := map[v32 := v33];
			v34 := v34[v32 := v33[f16 := f16]];
		}
		var v35: array<int> := new int[26];
		v35[829] := f16;
		var v36 := "bso";
		v35[829] := |v36| / |(seq(0x34, i3  => (DC3(DC2(f12, f12, f12, 'k', 'd')))))|;
		var v37: map<bool, bool> := map[true := f12];
		v37 := v37;
		var v38 := DC4(-f16);
		var v39: array<bool> := new bool[28](i4 => false);
		var v40: map<D2, array<bool>> := map[if (f12) then v38 else v38 := v39];
		v35[829], v40 := f16 + (f16 * v35[829]), v40;
		var v41: seq<bool> := [f12];
		var v42: map<int, seq<bool>> := map[f16 := v41];
		var v43: array<seq<bool>> := new seq<bool>[10] [v41, v41, v41, v41[v35[829] := f12], v41, v41, fm10(!f12, f12, v41[f16], globalState), v41, (if (fm1(f12, f16, f16, globalState) in v42) then v42[fm1(f12, f16, f16, globalState)] else v41) + v41, [f12]];
		var v44 := DC0(f12);
		var v45: set<bool> := {f12};
		var v46: seq<seq<bool>> := [v41, v41, ([f12, v44.cf0, f12])[|v45| := f12], v41, v41];
		var v47: map<int, bool> := map[v35[829] := f12];
		v43[21] := v46[|v47|];
		v43[21] := v41;
		v44 := v44;
		r0 := |map[!f12 := v35[829]]|;
	}
	method m7(p0: multiset<int>, p1: string, globalState: GlobalState) returns (r0: array<string>) {
		var v0 := DC0(f12);
		var v1 := 's';
		var v2 := DC2(v0.cf0, f12, f12, v1, 'w');
		match (v2) {
			case DC2(cf2, cf3, cf4, cf5, cf6) =>
				globalState.f9 := f16;
				var v3: seq<int> := [f16, f16, f16];
				f16 := v3[f16];
				globalState.f7, v3, globalState.f3 := cf3, v3 + (v3 + v3), false;
				if ((p0 - multiset{f16}) < p0) {
					globalState.f7 := if (cf3) then cf4 else f12 <== f12;
					var v4: map<int, bool> := map[|fm9(f16, cf3, globalState)| := cf3];
					var v5: array<bool> := new bool[6](i0 => cf4);
					var v6: map<array<bool>, bool> := map[v5 := cf4];
					v4 := v4[-f16 := if (v5 in v6) then v6[v5] else cf4];
					var v7: seq<bool> := [cf4];
					globalState.f9, globalState.f9, cf3, f11, globalState.f3 := |fm9(-0x194, cf3, globalState)|, f16 / 0x12e, cf2, multiset(v7) + (if (cf4) then f11 else multiset{cf2, cf3, cf3}), cf3;
					globalState.f9 := f16;
					var v8: seq<string> := [p1];
					var v9: set<int> := {f16};
					var v10: array<string> := new string[23] [p1, p1, v8[f16], "exlcjc", p1, p1, p1, fm9(|v9|, cf3, globalState), p1, seq(527, i1  => (v1)), p1, seq(0x6e, i2  => (cf6)), p1, p1, p1, ("mw")[|v7| := cf6], p1, p1, p1, p1, seq(0x319, i3  => (v1)), seq(0xe9, i4  => (cf6)), p1];
					var v11: map<char, array<string>> := map[cf5 := v10];
					var v12: map<bool, array<string>> := map[cf2 := v10];
					var v13: seq<array<string>> := [v10];
					var v14: array<array<string>> := new array<string>[19] [v10, v10, v10, if (v1 in v11) then v11[v1] else v10, if (f12) then v10 else v10, v10, v10, v10, v10, if (cf4 in v12) then v12[cf4] else v10, v10, v10, v10, v10, if (cf4) then v13[f16] else v10, v10, v10, if (cf2 in v12) then v12[cf2] else v10, v10];
					v14[160] := v10;
					v14[160] := v10;
				} else {
					var v15: map<bool, int> := map[f12 := f16];
					globalState.f7 := fm2(p0[f16 := f16] + p0, 0x300 / -f16, v15, f16 == f16, globalState);
					v1 := v1;
					globalState.f9 := f16 - f16;
					var v16: seq<bool> := [cf4, cf2];
					var v17: map<string, seq<bool>> := map[seq(0x176, i5  => (v1)) := v16 + v16];
					var v18: T0 := new C0(cf3, f16, f11, cf2);
					var v19: map<T0, seq<bool>> := map[v18 := v16];
					v17 := v17[if (cf3) then p1 else p1 := if (v18 in v19) then v19[v18] else v16];
					var v20: array<C0> := new C0[26];
					var v21: map<int, int> := map[f16 := f16];
					var v22: map<char, string> := map[cf6 := p1];
					globalState.f9, globalState.f7, f16, v20, globalState.f9 := |(v16 + v16)| / f16, v16[f16 * f16], (-|map[false := cf4]| / |[cf2]|) % (f16 + f16), v20, if (cf2) then f16 else -(f16 - (if (f16 in v21) then v21[f16] else |v22|));
				}
				
			case DC1(cf1) =>
				var v23 := DC1(cf1);
				var v24: map<D1, set<bool>> := map[v23 := cf1];
				var v25: array<int> := new int[8] [|p0|, |(p1 + "q")|, 0x87, f16, f16 * fm1(f12, 0x1d3, f16, globalState), f16, f16, |(if (v23 in v24) then v24[v23] else cf1)|];
				globalState.f1 := v25;
				var v27: map<int, map<map<int, bool>, int>> := map[f16 := map[(map v26 : int | (893 <= v26) && (v26 < 0x1f5) :: (v26 + |map[f16 := v1]|) := (f12))[73 := f12] := f16]];
				var v29: set<int> := {f16};
				if (if (f12) then f12 else f16 < |(if (-670 in v27) then v27[-670] else map[map v28 : int | v28 in v29 :: (v28 % 0x229) := (f12) := f16])|) {
					globalState.f7 := f12;
					f12 := f12;
					f16, globalState.f3 := fm1(f12, f16, f16, globalState), false;
					globalState.f9 := fm1(f12, 0x342, if (f12 in f11) then f11[f12] else 0x351, globalState) % (f16 + 0x3b1);
					var v30: array<array<int>> := new array<int>[21];
					var v31: map<multiset<int>, array<array<int>>> := map[multiset{f16} := v30];
					var v32 := DC4(f16);
					var v33: map<bool, D2> := map[f12 := v32];
					v31 := v31[p0 * p0[|v33| := f16] := v30];
				} else {
					var v34 := "pekwurnd";
					var v35: map<char, string> := map[v1 := "tfxhr"];
					v34 := if ('w' in v35) then v35['w'] else p1;
					var v36: array<bool> := new bool[14](i6 => !(if (true) then f12 else true));
					v36[726] := |"gdkv"| > f16;
					v36[726] := f12;
					var v37 := DC6(v36);
					v36 := v37.cf14;
					var v38: array<multiset<int>> := new multiset<int>[9];
					v38[284] := fm13(v1, f12, globalState);
					v38[284] := p0 * p0;
					v36[726] := f12;
				}
				
				var v39: map<bool, int> := map[!false := f16];
				var v40: seq<map<bool, int>> := [v39, v39];
				v40 := if (true) then ([v39])[|p1| := v39] + [v39] else v40;
				var v41: array<bool> := new bool[26];
				v41 := v41;
			case DC3(cf7) =>
				if (!fm0(v1, globalState)) {
					var v42: array<char> := new char[24] [v1, v1, v1, v1, 'f', 'w', v1, v1, v1, v1, v1, 'c', v1, v1, v1, 'r', v1, v1, v1, v1, v1, v1, v1, 'w'];
					var v43: seq<array<char>> := [v42];
					v43 := v43;
					var v44: map<int, bool> := map[588 := f12];
					v44 := v44[f16 := f12];
					var v45: set<bool> := {f12, f12, f12, f12};
					globalState.f8 := f16 == |v45|;
					var v46: array<bool> := new bool[11];
					v46 := v46;
					var v47 := "njgjri";
					var v48: array<int> := new int[27](i7 => i7 * -f16);
					var v49 := DC10(v48);
					var v50: seq<array<int>> := [v48];
					var v51: array<array<int>> := new array<int>[11] [v48, v48, if (f12) then v48 else v48, (v49.(cf24 := v48)).cf24, v50[f16], v48, v48, v48, v48, v48, v48];
					var v52: map<bool, char> := map[false := v1];
					var v53: map<bool, int> := map[f12 := f16];
					v46[387] := fm2(multiset{if (!f12 in f11) then f11[!f12] else |v52|}, f16, v53, f12, globalState) && f12;
					v47, v51, v46[387] := v47[f16 := fm14(v2, f16, globalState)], v51, v47[f16 := 'c'] < (v47[f16 := v1] + p1);
				} else {
					var v54: map<int, bool> := map[f16 := f12];
					v1 := if (v54 == v54) then v1 else v1;
					var v55: array<string> := new string[1](i8 => "nol");
					v55[481] := p1;
					v55[481] := p1;
					globalState.f9 := f16 * f16;
					var v56: seq<int> := [f16, 0x116];
					var v57 := DC12(v56);
					v56 := v57.cf29[if (f12) then |multiset{f12, f12}| else -|v55[481]| := f16];
					var v58: seq<bool> := [f12, f12];
					globalState.f9 := f16 / (f16 - |v58|);
				}
				
				var v59: array<string> := new string[29];
				r0 := v59;
				match (v0) {
					case DC0(cf0) =>
						var v60: array<int> := new int[24];
						v60[693] := f16;
						var v61 := "jljlfkl";
						var v62: array<D1> := new D1[21];
						var v63 := DC3(fm15(|v61|, f16, cf0, globalState));
						var v64 := DC3(v63);
						v62[188] := v64;
						var v65 := DC8(f16, fm1(true, f16, f16, globalState));
						var v66: seq<bool> := [true, f12];
						var v67: map<D3, int> := map[DC8(f16, |[f16, |v66|]|) := f16];
						var v68: seq<int> := [f16];
						var v69: set<string> := {fm9(f16, f12, globalState)};
						var v70: multiset<string> := multiset{p1, "sarchm", p1};
						var v71: array<char> := new char[25] [v1, v1, v1, v1, v1, v1, fm14(v2, f16, globalState), v1, v1, v1, v1, v1, v1, v1, v1, fm14(fm15(f16, f16, f12, globalState), 651, globalState), 'o', v1, v1, 'p', v1, v1, v1, v61[f16], v1];
						var v72: map<array<char>, bool> := map[v71 := false];
						var v73: set<bool> := {false, cf0, !cf0};
						var v74: set<int> := {f16};
						var v75 := DC11(cf0, f16, fm1(f12, -|[f16]|, f16, globalState), |v74|);
						var v76: map<string, int> := map[v61 := |v68|];
						var v77: C0 := new C0(cf0, |v76|, f11, v66[if (|f11| in p0) then p0[|f11|] else f16]);
						var v78: set<C0> := {v77, v77, v77, v77};
						var v79: array<bool> := new bool[11] [v65 in v67, f16 <= |v68|, -0x5c == |v69|, f12, v61 in v70, if (v71 in v72) then v72[v71] else f12, cf0, v73 <= {v75.cf25}, cf0, cf0, |v78| > v77.f18];
						var v80: map<bool, int> := map[v77.f17 := 0x35];
						v79[765] := v77.fm2(p0, f16, v80, v77.f17, globalState);
						var v81: map<string, bool> := map[p1 := v77.f17];
						v60[693], v61, v62[188], v79, v79[765] := f16, v61, v64, v79, !(if (p1 in v81) then v81[p1] else cf0);
						v60[693] := -0x12;
						v59[736] := p1;
						v59[736] := fm9(v60[693], -v60[693] >= (if (v77.f17 in v80) then v80[v77.f17] else f16), globalState);
						v2 := v2;
				}
				
				var v82: array<char> := new char[11];
				v82[113] := v1;
				v82[113] := if (f12) then v1 else v1;
		}
		
		for i9 := f16 to f16 {
			var v83: map<int, bool> := map[|(seq(434, i10  => (f16)))| := f12];
			var v84 := DC11(!false, i9, i9, |v83|);
			var v85 := new C0(v84.cf25 <==> f12, -i9, f11, false);
			var v86: array<map<bool, seq<int>>> := new map<bool, seq<int>>[22];
			var v87: seq<int> := [f16, v85.f18];
			var v88: seq<int> := [|v87|];
			var v89: map<bool, seq<int>> := map[f12 := v88];
			v86[225] := v89 + v89;
			var v90: array<bool> := new bool[2];
			var v91: seq<array<bool>> := [v90, v90];
			v86[225] := if (v91 == v91) then if (f12) then v89 else v89 else v89 + v89;
			f12 := v85.f17 && (v85.f17 && f12);
			var v92: map<int, seq<int>> := map[v85.f18 := seq(0x42, i11  => (v85.f18))];
			var v93: set<seq<int>> := {v87, v87, if (i9 in v92) then v92[i9] else v88};
			var v94: map<set<seq<int>>, bool> := map[v93 := f12];
			v94 := v94[v93 := v85.f17];
		}
		var v95: array<bool> := new bool[20];
		v95[617] := f12;
		v95[617] := (f11 + multiset{f12, true}) > f11;
		var v96: array<char> := new char[16] ['b', v1, v1, v1, v1, v1, v1, fm14(v2, -f16, globalState), v1, v1, v1, v1, v1, v1, v1, v1];
		forall i12 | 0 <= i12 < v96.Length {
			v96[i12] := v1;
		}
		var v97: array<string> := new string[3](i13 => p1);
		var v98 := DC13(f16, v95[617], f16, v97, v95[617]);
		var v99: seq<bool> := [v95[617], v95[617], f12, false];
		var v100: array<multiset<int>> := new multiset<int>[3] [p0, p0, (p0[v98.cf32 := -f16])[|v99[0x8e := v95[617]]| := f16]];
		v100[505] := p0 - p0;
		v100[505] := p0 * p0;
		if (v95[617]) {
			f12 := (|multiset{v95[617]}| / f16) <= f16;
			globalState.f9 := f16;
			globalState.f7 := fm0(v1, globalState);
			var v101 := DC6(v95);
			match (v101) {
				case DC7(cf15, cf16, cf17) =>
					globalState.f7 := true;
					cf15, globalState.f9, cf17, globalState.f9 := -(|"eokqyem"| / cf16), (cf16 + cf15) % f16, cf17, -cf17.f18;
					var v102: map<int, int> := map[cf15 := f16];
					var v103: set<int> := {cf15, cf17.f18, cf16, 0xe4 * |v102|, cf17.fm12(globalState)};
					v103 := v103;
					var v104: set<bool> := {cf17.f17};
					var v105: map<bool, bool> := map[true := f12];
					var v106: set<bool> := {f12, v95[617], !f12, fm16(|v104|, globalState) != v105};
					var v107: seq<int> := [cf15, |v100[505]|];
					v95[617], cf16, v106, globalState.f8 := (set v108 : int | v108 in v107[f16 := cf15] :: (v108 - 0x3e0)) > (v103 - v103), |v102|, v106 * {cf17.f17}, {f12} >= v104;
				case DC8(cf18, cf19) =>
					v97[474] := p1;
					v97[474], globalState.f6, globalState.f7, globalState.f3, v95 := p1, (if (!!false) then [v95[617], f12] else v99) + v99, !fm0(v1, globalState), f12, (v101.(cf14 := v95)).cf14;
					var v109: T0 := new C0(false, f16, f11, v99[cf19]);
					var v110: map<int, T0> := map[cf19 := v109];
					var v111: map<bool, int> := map[v95[617] := f16];
					globalState.f7 := fm2(fm13(v1, v95[617], globalState), |v110| % cf18, v111 + v111, v95[617], globalState);
					var v112: seq<int> := [cf18, cf18];
					var v113: array<int> := new int[25];
					var v114: map<multiset<int>, array<int>> := map[multiset((if (f12) then v112 else v112)[f16 := f16]) := v113];
					v114 := v114[p0 - v100[505] := v113];
					var v115: C0 := new C0(v95[617], f16, v109.f11 + f11, v95[617] || v95[617]);
					v115 := v115;
				case DC9(cf20, cf21, cf22, cf23) =>
					var v116: array<int> := new int[16];
					v116[12] := f16;
					var v117: map<string, bool> := map[p1 := f12];
					v116[12] := |v117| % cf22.f18;
					var v118 := DC2(f12, v95[617], f12, v1, v1);
					var v119 := DC3(v118);
					var v120 := DC3(v119);
					var v121 := DC3(v118);
					var v122: set<D1> := {DC3(v120).(cf7 := v118), v121};
					var v123: multiset<set<D1>> := multiset{v122 + v122, v122};
					v123 := v123 - (v123 - fm17(cf22.f17, !false, cf22.f18, 0xbe, globalState));
					cf20 := -cf20;
					var v124: map<bool, array<bool>> := map[f12 := v95];
					v124 := v124[v95[617] := v95];
				case DC6(cf14) =>
					var v125: set<int> := {f16};
					var v126: map<bool, int> := map[f12 := f16];
					var v127: set<bool> := {f12, v95[617]};
					var v128 := DC4(f16);
					var v129: array<int> := new int[6] [f16, f16, -|map[DC5(|v99|, f16, f16, f16, v99).cf9 := f12]|, |(v125 - {-f16})|, -623 - |v126[!f12 := fm1(!v95[617], |v127|, f16, globalState)]|, fm1(f12, v128.cf8, 0x31d, globalState)];
					v129[273] := f16;
					var v130: map<char, bool> := map['o' := v95[617]];
					var v131: multiset<char> := multiset{v1, v2.cf5, v1, p1[f16]};
					var v132: seq<int> := [f16, |(v130 + v130)|, |v99[0x29f := v95[617]]| % f16, f16, |(v131 - v131)|];
					v129[273] := v132[f16];
					var v133 := new C0(f12, f16, multiset{v95[617], f12} + f11, f16 == f16);
					v129[273] := f16;
					v129[273] := v129[273];
			}
			
			var v135: set<int> := {f16};
			var v136 := new C0(false, f16, multiset(([v95[617]])[f16 := f12]), (set v134 : int | (447 <= v134) && (v134 < 0x26f) :: (v134 % f16)) > v135);
		} else {
			globalState.f9 := f16;
			var v137: map<bool, int> := map[!v95[617] := f16];
			globalState.f3 := (-201 % (if (v95[617] in v137) then v137[v95[617]] else f16)) != |(seq(0x25e, i14  => (-f16)))|;
			if ((f16 == fm1(f12, f16, f16, globalState)) !in map[f12 := f16]) {
				var v138: set<bool> := {f12, f12};
				var v139: C0 := new C0(f12, |map[f16 := |v138|]|, f11[f12 := |p1|], f12);
				var v140 := DC9(f16, |fm18(f16, f12, globalState)|, v139, v139.f18);
				var v141: map<int, int> := map[f16 := f16 / |map[v140 := v139.f17]|];
				v141 := (v141 + v141) + v141;
				v95[617] := true;
				var v142: array<int> := new int[10](i15 => i15 * 263);
				v142[122] := f16;
				var v143 := DC6(v95);
				v142[712] := f16;
				v1, v142[122], v143, v142[712] := v1, -(f16 / (f16 + |f11|)), v143, f16;
				globalState.f3 := v95[617];
				v95[617], v100[505], globalState.f7, f12 := f12, p0, v139.fm2(fm13(v1, false, globalState), -f16, v137, v95[617] <== f12, globalState), v139.f17;
			} else {
				var v144: array<D3> := new D3[24];
				var v145 := DC6(v95);
				v144[190] := v145.(cf14 := v95).(cf14 := v95);
				v144[190] := v145;
				var v146 := new C0(f12, |v100[505]|, f11 * f11, !(v95[617] ==> v95[617]));
				var v147: T0 := new C0(if (v95[617]) then v95[617] else f12, v146.f18 * f16, f11[f12 := v146.f18], f12);
				v147 := v147;
				var v148 := new C0(v95[617], v146.f18, v147.f11 + v147.f11, f12);
				var v149 := "ggsc";
				v149 := p1;
			}
			
			var v150: seq<array<bool>> := [v95];
			var v151: multiset<char> := multiset{v1};
			var v152: map<bool, multiset<char>> := map[false := v151];
			var v153: array<array<bool>> := new array<bool>[10] [v95, v95, v95, v95, v150[|v152|], v95, v95, v150[f16], v95, v95];
			v153[807] := v95;
			v153[807] := v95;
			globalState.f3 := v95[617];
		}
		
		r0 := v97;
	}
}

class C2 extends T0 {
	var f15 : bool
	constructor (f15 : bool, f11 : multiset<bool>, f12 : bool) {
		this.f15 := f15;
		this.f11 := f11;
		this.f12 := f12;
	}
	
	function fm2(p0: multiset<int>, p1: int, p2: map<bool, int>, p3: bool, globalState: GlobalState): bool {
		f12
	}
	function fm8(p0: bool, p1: bool, globalState: GlobalState): seq<int> {
		[622] + (seq(392, i0  => (0x3bf)))
	}
	method m4(p0: int, p1: bool, p2: D1, globalState: GlobalState) returns (r0: bool, r1: string, r2: int, r3: bool) {
		if (|[p0]| == p0) {
			if (p1) {
				var v0: array<bool> := new bool[16];
				v0 := if (f12 && p1) then v0 else v0;
				r2 := p0 / p0;
				var v1 := "nbffrd";
				globalState.f7 := ((seq(0x5, i0  => ('j'))) + v1) <= "xycvusgt";
				var v2: array<array<int>> := new array<int>[17];
				var v3: map<array<array<int>>, bool> := map[v2 := f15];
				v3 := v3[v2 := f15];
				var v4: map<char, bool> := map['j' := f12];
				var v5: set<array<bool>> := {v0};
				var v6: map<int, bool> := map[|v4| := v5 > v5];
				v1, v6, v0, globalState.f7, globalState.f3 := v1 + v1, v6, v0, f12, f12 <== f15;
			} else {
				var v7: seq<int> := [p0, p0];
				v7 := v7;
				var v8: array<bool> := new bool[28];
				v8[166] := (seq(0x17, i1  => ('u'))) != (seq(924, i2  => ('h')));
				v8[166] := f12;
				var v9: array<int> := new int[4];
				v9[459] := fm1(f15, p0, p0, globalState);
				var v10: map<int, bool> := map[p0 := f12];
				v9[459] := (p0 % p0) - |v10|;
				m0(globalState);
				var v11: map<array<bool>, array<int>> := map[v8 := v9];
				v11 := v11[v8 := v9];
			}
			
			globalState.f9 := -|"xswkq"|;
			globalState.f9 := p0;
			r2 := if ((f12 ==> f15) in f11) then f11[f12 ==> f15] else p0;
			var v12: set<int> := {-0x1df, p0, p0};
			var v13: map<bool, set<int>> := map[p1 ==> f12 := {|v12|}];
			v12 := if (false in v13) then v13[false] else v12;
		} else {
			var v14: seq<int> := [p0];
			var v16 := new C0(f12, p0, f11, {|v14|} != (set v15 : int | (0x117 <= v15) && (v15 < 0x12b) :: (v15 % p0)));
			var v17: map<bool, int> := map[v16.f17 := p0];
			v17 := v17[p1 := p0];
			var v18: seq<bool> := [f15];
			var v19: array<bool> := new bool[2] [v16.f17, !(|v18| > v16.f18)];
			var v20: map<int, bool> := map[v16.f18 := f12];
			var v21: map<bool, bool> := map[f15 := if (0x2b6 in v20) then v20[0x2b6] else p1];
			v19[21] := if (!!true in v21) then v21[!!true] else v16.f17;
			v19[21] := f15;
			globalState.f3 := p0 == (p0 / p0);
			v17 := v17[v19[21] := v16.f18];
		}
		
		var i3 := 0;
		while (p1)
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v22: seq<int> := [p0];
			var v23: map<int, bool> := map[-p0 := true];
			v22 := fm8(if (p0 in v23) then v23[p0] else f12, p1, globalState);
			var v24: set<bool> := {f12};
			var v25 := 'q';
			var v26: map<int, char> := map[|v24| - p0 := v25];
			v26 := v26;
			var v27: multiset<int> := multiset{0x64, p0, p0, 494, p0};
			globalState.f9 := (if (!f15) then |f11| else if (p0 in v27) then v27[p0] else 0x13c) * p0;
			var v28: map<seq<int>, int> := map[v22[-0x2e4 := p0] := 0x3ba];
			var v30: multiset<seq<int>> := multiset{v22};
			var v31: seq<map<seq<int>, int>> := [v28];
			v28 := (v28 + (map v29 : seq<int> | v29 in v30 :: (v29) := (p0))) + v31[p0];
		}
		f15 := f15;
		var v32: seq<bool> := [f12];
		var v33: seq<int> := [fm1(p1, p0, p0, globalState), p0, p0];
		var v34: map<int, int> := map[|f11| * (if (v32[|v33|] in f11) then f11[v32[|v33|]] else p0) := p0];
		var v35 := "hptlhvxc";
		var v36 := 'c';
		v34 := v34[if (f12) then |v35[-p0 := v36]| else p0 := |v35|];
		var v37: array<int> := new int[29](i5 => i5 + 0x2f6);
		forall i4 | 0 <= i4 < v37.Length {
			v37[i4] := i4 * DC4(-115).cf8;
		}
		r2 := -p0 - (p0 + -p0);
		var v38: multiset<int> := multiset{p0};
		var v39 := DC15(v38);
		var v40: set<bool> := {p1, true};
		var v41: map<bool, int> := map[f12 := p0];
		r0 := fm2(v39.cf35 * multiset{p0, |v35|, |v40|}, p0, v41, f15, globalState);
		r1 := seq(0x370, i6  => (v36));
		r2 := p0 % fm1(p1, if (f15 in f11) then f11[f15] else -0x37, |v40|, globalState);
		r3 := !f12;
	}
	method m5(p0: int, p1: string, globalState: GlobalState) returns (r0: D0, r1: int) {
		r1 := p0;
		var v0 := 'a';
		var v1 := DC2(f15, f12, true, v0, v0);
		var v2: map<bool, char> := map[f15 := fm14(v1, 0x158, globalState)];
		v2 := v2[f15 := v0];
		var v3: array<bool> := new bool[3] [f15, f12, f15];
		forall i0 | 0 <= i0 < v3.Length {
			v3[i0] := !f15;
		}
		var v4: set<bool> := {f12, f15};
		var v5 := DC1(v4);
		var v6: set<D1> := {v5};
		var v7: map<bool, bool> := map[f15 := v6 < {DC1(v4)}];
		if (if (f15 in v7) then v7[f15] else f15) {
			globalState.f9 := p0;
			f12 := f12;
			var v8: map<bool, int> := map[false := p0];
			var v9: seq<string> := [p1, ((seq(591, i1  => ('h')))[if (f15 in v8) then v8[f15] else p0 := v0])[p0 := v0]];
			var v10: seq<int> := [-|p1|, |v9| - p0, if (f15) then p0 else p0, p0, p0];
			var v11: C0 := new C0(f12, p0, f11, f15);
			globalState.f9 := v10[|[v11, v11]|];
			var v12: seq<seq<int>> := [v10 + v10, v10];
			v10 := v12[v11.f18];
			var v13: set<char> := {v0};
			var v14: map<D5, bool> := map[DC14() := !false];
			var v15 := DC17(v14);
			var v16: array<int> := new int[26] [p0, p0, v11.f18, p0, 0x12c * v11.f18, 0x277 - v11.f18, p0 + p0, p0, p0, v11.f18 + p0, |v4|, 0x0, |v13| % p0, -|v15.cf39|, p0, p0, fm1(f12, v11.f18, v11.f18, globalState), v11.f18, if (f12 in f11) then f11[f12] else v11.f18, v11.f18, v11.f18, |p1|, -v11.f18 * v11.f18, v11.f18 + |f11|, v11.f18, v11.f18];
			v16[371] := v10[p0];
			v16[371] := -(p0 + 0x27b);
		} else {
			var v17: array<seq<int>> := new seq<int>[17];
			v17[387] := fm8(f15, f12, globalState);
			var v18: seq<int> := [p0];
			var v19 := DC12(v18);
			v17[387] := DC12(v18).cf29 + v19.cf29;
			var v20: array<int> := new int[25](i2 => i2 % p0);
			var v21: map<bool, array<int>> := map[f15 := v20];
			var v22: seq<array<int>> := [if (f15 in v21) then v21[f15] else v20];
			v3[260] := f12;
			var v23 := "muarlkhi";
			v22, v3[260], v23 := (v22 + v22) + v22, v7 == v7, p1;
			var v24: map<int, string> := map[p0 := v23];
			var v25: multiset<int> := multiset{p0};
			globalState.f3, globalState.f9, v24 := p1 <= (seq(370, i3  => ('a'))), ((if (p0 in v25) then v25[p0] else p0) % p0) / (---p0 / |p1|), v24[|(seq(108, i4  => ('f')))| := p1] + v24;
			var v26: map<int, int> := map[p0 := v18[p0]];
			var v27: set<map<int, int>> := {v26, v26};
			var v28: map<bool, int> := map[f12 := fm1(!f12, p0, |v27|, globalState)];
			v28 := v28;
			var v29: map<D6, multiset<int>> := map[DC15(v25) := v25];
			var v30 := DC15(v25);
			var v31: set<int> := {|v28|};
			var v32 := DC11(fm2(if (v30 in v29) then v29[v30] else v25, p0, v28, v3[260], globalState), p0, p0, fm1(v3[260], -|v31|, |v7|, globalState));
			var v33: map<array<int>, bool> := map[v20 := v32.cf25];
			v33 := v33[v20 := f12];
		}
		
		var v34 := DC19(f11);
		var v35: seq<bool> := [f15];
		var v36 := new C1(|p1|, v34.cf45[v35[-0x359] := -462], f12);
		globalState.f3 := fm0(v0, globalState);
		var v37: set<int> := {p0};
		var v38 := DC0(false);
		r0 := if ({p0, v36.f16} <= v37) then v38 else v38;
		r1 := v36.f16 + p0;
	}
}

class C3 extends T0 {
	const f13 : int
	const f14 : bool
	constructor (f13 : int, f14 : bool, f11 : multiset<bool>, f12 : bool) {
		this.f13 := f13;
		this.f14 := f14;
		this.f11 := f11;
		this.f12 := f12;
	}
	
	function fm2(p0: multiset<int>, p1: int, p2: map<bool, int>, p3: bool, globalState: GlobalState): bool {
		f12 <== f14
	}
	function fm6(p0: bool, p1: int, globalState: GlobalState): bool {
		if (if (true) then f12 else f12) then if (f12) then f14 else !f12 else f13 < f13
	}
	method m2(p0: string, p1: bool, p2: array<string>, globalState: GlobalState) returns (r0: int, r1: bool) {
		if (f14) {
			var v0: array<multiset<char>> := new multiset<char>[29];
			var v1 := 'd';
			var v2: multiset<char> := multiset{'a', v1};
			var v3: seq<multiset<char>> := [v2];
			v0[206] := v3[f13];
			v0[206] := multiset(p0 + (p0 + (([v1, v1])[f13 := v1])[f13 := 'i']));
			if (p1) {
				var v4: set<int> := {f13, f13};
				var v5: array<bool> := new bool[5] [f12, f13 != f13, v4 !! v4, f14, p1];
				v5[1] := p1 <==> f14;
				var v6: seq<array<bool>> := [v5, v5, v5];
				var v7: map<seq<array<bool>>, bool> := map[v6 := f12];
				v5[1] := if (v6 in v7) then v7[v6] else if (p1) then false else p1;
				var v8: array<int> := new int[15];
				v8[557] := f13;
				v8[557] := f13;
				var v9: map<int, seq<char>> := map[|[v8, v8, v8]| := p0];
				v9 := v9[v8[557] * f13 := p0];
				var v10 := DC2(true, f14, v5[1], v1, 't');
				var v11 := DC3(v10);
				var v12: map<int, D1> := map[-76 := v11];
				v12 := v12[f13 := DC3(v10)];
				var v13 := DC0(p1);
				v13, globalState.f7 := v13, p1;
			} else {
				var v14: seq<bool> := [f12];
				var v15 := DC5(f13, f13, f13, f13, v14);
				r0 := |p0| * (v15.(cf10 := f13, cf13 := fm7(globalState))).cf10;
				globalState.f9 := f13;
				var v16 := new C1(f13, f11, p1);
				var v17 := v16.m6(globalState);
				var v18 := new C1(f13, f11, f14);
			}
			
			if (!f14) {
				var v19: seq<char> := ['k'];
				var v20: array<int> := new int[11];
				v20[265] := f13;
				var v21: set<int> := {f13, f13, 0x0};
				var v23 := DC18(f13, seq(0x313, i0  => (v1)), f12, f14, set v22 : int | v22 in v21 :: (v22 + |multiset{false}|));
				r0, v19, v20[265], globalState.f7 := if (p1) then 43 else f13, v23.cf41, f13 * f13, !f12;
				var v24: map<bool, bool> := map[p1 := f14];
				var v25: array<C2> := new C2[10];
				var v26: map<bool, array<C2>> := map[f11 > multiset{p1} := v25];
				var v27 := DC11(false, v20[265], f13, f13);
				var v28: map<int, array<int>> := map[v20[265] := v20];
				v24, v26, v27 := v24, v26, DC11(DC20(p1, f12, p1).cf48, v20[265], v20[265], |(if (f14) then v28 else v28)|);
				var v29: array<seq<array<D6>>> := new seq<array<D6>>[23];
				var v30: array<D6> := new D6[2];
				v29[684] := [v30, v30, v30, v30, v30];
				v29[684] := [v30, v30];
				var v31 := DC0(f12);
				var v32: map<bool, char> := map[v31.cf0 := v1];
				m3(f13 - |v32|, globalState);
				r0, globalState.f9, f11, v20[265] := f13, 386, f11, v20[265] + (0x396 % v27.cf28);
			} else {
				r1 := false;
				globalState.f9 := f13;
				var v33: array<set<set<int>>> := new set<set<int>>[14](i1 => {{|p0|}} + {{|p0|, f13}});
				var v34: map<int, bool> := map[f13 := f12];
				var v35: set<int> := {f13, |v34|};
				var v36: seq<bool> := [f14];
				var v37: T0 := new C1(f13, multiset{f12, true}, true);
				var v38: seq<T0> := [v37];
				var v39: map<bool, T0> := map[p1 := v38[f13]];
				var v40 := DC5(|v36|, |v39|, f13, f13, [v37.f12, v37.f12]);
				var v41: set<set<int>> := {v35, v35, {v40.cf12, |(seq(0x2c2, i2  => (v1)))|}};
				v33[974] := v41;
				var v42: array<bool> := new bool[10];
				v33[974], v42, v37.f12 := v41, v42, if (p1) then p1 else f12;
				v42[888] := v37.f12;
				v42[888] := p1;
				var v43: multiset<D2> := multiset{v40, v40};
				var v44: seq<multiset<D2>> := [multiset{v40, v40}, v43[v40 := f13], v43];
				v44 := v44 + (seq(-0x390, i3  => (v43)));
			}
			
			var v45: array<C0> := new C0[27];
			var v46: map<bool, array<C0>> := map[p1 := v45];
			var v47: seq<bool> := [f14];
			var v48: map<int, map<bool, array<C0>>> := map[fm1(f14, f13, |v47|, globalState) := map[f14 := v45]];
			var v49 := DC23(v46);
			var v50: array<map<bool, array<C0>>> := new map<bool, array<C0>>[25] [v46, if (f12) then v46 else v46, v46[true := v45], v46, v46 + v46, v46, v46, v46, v46, if ((if (true in f11) then f11[true] else f13) in v48) then v48[if (true in f11) then f11[true] else f13] else v46, v49.cf55, if (f12) then v46 else v46, map[f12 := v45], v46, v46, (map[f14 := v45])[f12 := v45], v46, v46, v46, v46 + v46, v46, v46, v46, map[true := v45], v46[f12 := v45]];
			v50[514] := if (p1) then v46 else v46;
			var v51: multiset<int> := multiset{-f13};
			var v52: seq<multiset<int>> := [v51, v51];
			var v53: map<bool, int> := map[f14 := f13];
			r0, v50[514], globalState.f3, r0 := f13, v46 + v49.cf55, !fm2(v52[-|v47|], f13, v53[false := f13], f14, globalState), (f13 / f13) % (|f11| + f13);
			var v54: array<int> := new int[7];
			v54[68] := f13;
			v54[68] := f13;
		} else {
			var v55 := 'y';
			var v56: multiset<char> := multiset{v55, v55};
			var v57: array<seq<int>> := new seq<int>[8];
			var v58: seq<int> := [f13];
			v57[204] := v58;
			v56, r1, globalState.f9, v57[204] := multiset{v55}, f14, -f13 / f13, (v58 + v58) + v58;
			var v59: multiset<int> := multiset{f13};
			var v60 := DC2(f12, f12, f12, v55, v55);
			var v61 := DC2(p1, false, !p1, v55, fm14(v60, f13, globalState));
			var v62: array<string> := new string[23] [p0, seq(0x341, i4  => (v55)), p0, p0, p0, p0, p0, p0, p0, p0, p0, p0, p0, "ppym", p0 + p0, p0, p0, p0 + p0, if (f14) then p0 else p0[-|v59[f13 := -(if (f13 in v59) then v59[f13] else f13)]| := v61.cf6], p0, seq(-0x21e, i5  => (v55)), "drfnfyqmg", p0];
			var v63: map<bool, int> := map[f14 := |"whjni"|];
			var v64: seq<bool> := [true];
			var v65 := DC5(f13, |v56|, f13, if (f14 in v63) then v63[f14] else f13, v64);
			var v66: map<D2, array<string>> := map[if (f12) then v65 else v65.(cf9 := f13, cf13 := v64) := v62];
			var v67: array<int> := new int[26];
			var v68: seq<array<int>> := [v67, v67];
			var v70: set<int> := {f13, f13};
			var v71: seq<set<int>> := [set v69 : int | v69 in v59 :: (v69 * -0x184), v70, {|v64|, f13}];
			var v72: seq<array<int>> := [v67, v67, v68[v57[204][fm1(p1, f13, |v71|, globalState)]]];
			var v73 := DC24("c", f12, f13);
			globalState.f9, v62, v57[204], globalState.f4, globalState.f7 := -0x2a4, if (v65.(cf9 := 697) in v66) then v66[v65.(cf9 := 697)] else p2, v58 + (seq(-730, i6  => (f13))), v68[f13 % f13], (v73.(cf58 := f13)).cf57;
			var v74: map<int, bool> := map[f13 := f12];
			v74 := v74[f13 / f13 := p1 ==> f12];
			r0 := -f13;
			var v75: C0 := new C0(f14, f13, f11[f12 := f13], f14);
			var v76 := DC9(f13, (if (fm1(f12, f13, f13, globalState) in v59) then v59[fm1(f12, f13, f13, globalState)] else f13) % -|v57[204][|v57[204]| := -0x3c8]|, v75, v75.f18);
			v76 := v76;
		}
		
		var v77: seq<bool> := [!f14, f12, p1];
		r1, globalState.f3 := f14, v77[|map[f13 := f13]|];
		var v78: map<bool, int> := map[!f12 := f13];
		var v79: multiset<map<bool, int>> := multiset{v78};
		var v80 := DC27(v78);
		var v81: seq<int> := [|v78[p1 := -f13]|, f13, f13];
		var v82: array<multiset<map<bool, int>>> := new multiset<map<bool, int>>[13] [multiset{map[f14 := |p0|], v78, map[f12 := f13]}, multiset{v78, v78}, v79 * v79, v79, v79, multiset{v80.cf63}, v79, v79, v79, v79, v79, v79[v78 := |v81|], v79 - fm19(f13, globalState)];
		v82[599] := v79;
		v82[599] := v79 - (v79 + v79);
		var v83: map<seq<int>, int> := map[[f13] := f13];
		if (v81 in v83) {
			var v84: map<seq<bool>, bool> := map[v77 := p1];
			v84 := v84[v77 + v77[|f11| := !p1] := f12];
			var v85: C2 := new C2(f14, f11, f14);
			v85 := v85;
			var v86 := "kfcc";
			v86 := v86 + p0;
			var v87 := DC14();
			var v88 := DC17(map[v87 := p1]);
			var v89: map<D5, bool> := map[v87 := f12];
			v88 := v88.(cf39 := v89 + v89);
			globalState.f7 := if (v85.f15) then v85.f15 else f12;
		} else {
			globalState.f8 := f14;
			var v90 := 's';
			var v91 := DC2(f12, false, f14, v90, v90);
			var v92: map<int, char> := map[|v81| := fm14(v91, f13, globalState)];
			var v93: set<bool> := {f14};
			var v94: multiset<int> := multiset{f13};
			var v95: map<int, int> := map[-f13 := f13];
			var v96: set<int> := {f13, |v95|, f13};
			var v97: map<bool, set<int>> := map[fm2(v94, f13, v78, p1, globalState) := v96];
			var v98: array<int> := new int[24] [-|v81|, fm1(f14, f13, 375, globalState), |v92|, f13, f13, f13, f13, |v93|, f13, |"mwn"|, |p0|, f13, f13, f13, f13, |v97|, f13, f13, fm1(p1, |v78|, f13, globalState), f13, -f13, f13, f13, f13];
			var v99: set<array<int>> := {v98};
			v99 := if (f14) then v99 - v99 else v99 + v99;
			var v100: map<bool, array<int>> := map[f14 := v98];
			v100 := v100[f12 := v98];
			f12 := ((seq(0x22e, i7  => (f13))) + v81) != v81;
			var v101 := "m";
			v101 := p0;
		}
		
		var v102 := new C0(!f14, f13, f11, p1);
		f12 := f12;
		r0 := -fm1(f12, f13, f13 % 0x78, globalState);
		var v103 := DC12(seq(179, i8  => (382)));
		var v104: multiset<D5> := multiset{v103, v103, v103};
		r1 := v104 <= v104;
	}
	method m3(p0: int, globalState: GlobalState) {
		var v0: array<int> := new int[20];
		v0[3] := p0;
		v0[3] := p0;
		var v1: array<D4> := new D4[13];
		var v2 := DC10(v0);
		v1[525] := v2.(cf24 := v0);
		v1[525] := DC10(v0);
		if (f14) {
			var v3 := 'm';
			var v4 := DC2(f14, f14, f12, v3, v3);
			var v5: map<int, char> := map[p0 := fm14(v4, p0, globalState)];
			var v6: set<int> := {p0, v0[3]};
			f12, globalState.f9 := !(fm20(v3, 678, 388, v5, globalState) >= v6), -0x133;
			v0[3] := f13;
			var v7: map<array<int>, int> := map[v0 := |v6|];
			var v8: seq<array<int>> := [v0, v0];
			var v9: set<bool> := {f12};
			var v10: seq<int> := [|v9|, p0, f13, v0[3]];
			v7 := v7[v8[|v10|] := p0];
			var v11 := "hoftygcb";
			v11 := v11;
			globalState.f9 := p0 * f13;
		} else {
			globalState.f9 := 0x2bd;
			var v12: array<string> := new string[28](i0 => "ipdww");
			var v13 := DC13(f13, f12, p0 - p0, v12, f14);
			match (v13) {
				case DC13(cf30, cf31, cf32, cf33, cf34) =>
					var v14: C2 := new C2(f12, f11 - multiset{f14}, cf34);
					v14 := v14;
					cf32 := v0[3] % (|"klkgrol"| * cf32);
					var v15: array<bool> := new bool[8](i1 => f12);
					v15[803] := f12;
					v15[803] := f14;
					var v16: map<bool, array<bool>> := map[v14.f15 := v15];
					globalState.f3 := (v16 + v16) == v16;
				case DC14() =>
					globalState.f9 := fm1(f14, 0xc, |[0x56, 617, |f11|]| - v0[3], globalState);
					var v17: seq<bool> := [f12];
					var v18: map<map<bool, bool>, int> := map[map[false := false] := p0];
					var v19 := "x";
					var v20: set<int> := {|v19|};
					var v21: C0 := new C0(f12, p0, f11 - (multiset(v17))[f14 := |v17|], {-(if (map[DC0(false).cf0 := f14] in v18) then v18[map[DC0(false).cf0 := f14]] else -0x274), v0[3]} <= v20);
					v21 := new C0(f14, f13, multiset{v21.f17}, f14);
					var v22 := DC25(v19, v21.f17, p0);
					var v23 := 'y';
					v19 := v22.cf59[v0[3] := v23];
					var v24 := DC5(p0, v0[3], 223, p0, v17);
					globalState.f8, v23, globalState.f7, globalState.f9 := f12, v23, (v21.f18 % v21.fm12(globalState)) < (|v24.cf13| + f13), |v20| % |map[v23 := f12]|;
				case DC12(cf29) =>
					var v25: T0 := new C2(f14, f11, f12);
					v25 := v25;
					var v26 := 'c';
					var v27: seq<bool> := [f14];
					var v28: map<int, int> := map[|v27| := v0[3]];
					var v29 := new C0(fm0(v26, globalState), fm1(v25.f12, f13, f13, globalState), multiset{!f14, f14, f12, f12, !f14}, !((if (f13 in v28) then v28[f13] else v0[3]) == -0x1a8));
					var v30: set<bool> := {f12, f14};
					var v31 := new C0(v30 > {f12, f14}, 0x1f6, multiset{f14, v29.f17}, !v29.f17);
					var v32: array<bool> := new bool[6](i2 => map[v31.f17 := v31.f17] !in [map[v25.f12 := v31.f17]]);
					v32[186] := v31.f17;
					v32[186] := fm0(v26, globalState) ==> !v29.f17;
			}
			
			v0[3] := (-f13 % v0[3]) + f13;
			var v33: seq<array<int>> := [v0];
			var v34 := 'l';
			var v35 := DC2(f14, !f14, f12, 'c', v34);
			match (if (f12) then DC13(|f11|, f12, |v33|, v12, v35.cf2) else v13) {
				case DC13(cf30, cf31, cf32, cf33, cf34) =>
					globalState.f3 := cf34;
					var v36: array<bool> := new bool[19](i3 => cf34);
					v36[763] := !cf34;
					var v38: set<char> := {'c', v34, v34};
					v36[763] := (set v37 : char | v37 in fm21(cf34, true, cf31, globalState) :: (v37)) <= v38;
					var v39: seq<array<bool>> := [v36, v36, v36, v36, v36];
					v39 := v39;
					var v40: map<array<int>, int> := map[v0 := f13];
					v34 := if (v40 == v40) then v34 else v34;
				case DC14() =>
					var v41: seq<bool> := [f12];
					v0[3], globalState.f9, v1[525] := f13, |v41| / fm1(f12, p0, |"scarno"|, globalState), DC10(v0);
					var v42: seq<seq<bool>> := [v41, v41, v41[0x32f := f12]];
					var v43: map<bool, seq<seq<bool>>> := map[false := v42];
					v43 := v43[f14 := v42 + v42];
					var v44: array<bool> := new bool[28];
					v44[468] := f14;
					globalState.f9, v44[468], globalState.f8 := p0, f12, f12;
					var v45: array<array<bool>> := new array<bool>[15];
					v45 := v45;
				case DC12(cf29) =>
					var v46: seq<bool> := [f12];
					var v47: T0 := new C0(true, f13, multiset(v46), false);
					var v48: seq<T0> := [v47];
					var v49: map<int, seq<T0>> := map[p0 := v48 + v48];
					v49 := v49[p0 := v48 + v48];
					var v50 := new C2(v47.f12, f11, f14);
					var v51: array<set<int>> := new set<int>[5];
					var v52 := "tdx";
					var v53: set<int> := {|v52|, |(seq(583, i4  => (v34)))|};
					v51[668] := v53;
					v51[668] := set v54 : int | (0xbc <= v54) && (v54 < 25) :: (v54 - v0[3]);
					globalState.f8 := if (fm0(v34, globalState)) then v50.f15 else false;
			}
			
			globalState.f7 := fm6(!f12, f13, globalState);
		}
		
		var i5 := 0;
		while (f12)
			decreases 100 - i5
		{
			if (i5 >= 100) {
				break;
			}
			
			i5 := i5 + 1;
			m0(globalState);
			v0[3] := -0x136 * (-f13 % p0);
			var v55 := 'p';
			var v56: map<int, char> := map[v0[3] := v55];
			if (fm0(if (p0 in v56) then v56[p0] else v55, globalState)) {
				var v57: array<bool> := new bool[25](i6 => f13 < --|"ad"|);
				var v58 := "xmboh";
				var v59: set<int> := {|(seq(134, i7  => (v55)))|};
				var v60 := DC18(p0, v58, f14, f14, v59);
				var v61: seq<bool> := [f12];
				v57 := new bool[11] [!f14, f14, f14, false, f12, v55 !in v58, f14, !(v60.(cf44 := v59)).cf43, f14, f12, v61 != v61];
				var v62 := new C2(f12, (multiset{f14})[fm0('c', globalState) := f13], f14);
				var v63: map<int, array<bool>> := map[p0 := v57];
				globalState.f8 := (v63 == v63) ==> v62.f15;
				var v64 := new C2(f12, f11, f14);
				globalState.f9 := v0[3];
			} else {
				var v65: map<bool, int> := map[f12 := v0[3]];
				var v66: seq<int> := [|"tpnnpacy"|];
				var v67: seq<seq<int>> := [v66, v66];
				var v68: map<map<bool, int>, seq<int>> := map[v65 := v67[v66[p0]] + v66];
				v68 := v68;
				var v69: array<array<int>> := new array<int>[4];
				v69[251] := v0;
				v69[251] := v0;
				var v70: C2 := new C2(fm0(v55, globalState), (fm22(p0, f12, globalState))[f14 := v0[3]], f14);
				var v71: map<multiset<bool>, int> := map[f11[v70.f15 := p0] := v0[3]];
				var v72: map<char, int> := map[v55 := v0[3]];
				var v74: array<map<char, int>> := new map<char, int>[8] [v72[v55 := f13], v72, v72, v72, map v73 : char | v73 in {v55} :: (v73) := (v0[3]), v72, v72, v72];
				v74[354] := map[v55 := v0[3]];
				var v76: multiset<multiset<bool>> := multiset{f11};
				var v78: seq<multiset<bool>> := [f11, f11];
				var v79: set<multiset<bool>> := {f11, v78[p0]};
				v70, v0[3], v71, v0[3], v74[354] := v70, (fm1(f14, v0[3], f13, globalState) / -v0[3]) * p0, if (f13 != 0x19b) then (map v75 : multiset<bool> | v75 in v76 :: (v75) := (|v65|)) + (map v77 : multiset<bool> | v77 in v79 :: (v77) := (f13)) else map[f11 := f13] + v71, 815 / |v65|, v72;
				var v80: T0 := new C0(f12, v0[3], f11, f12);
				var v81 := "xcisubu";
				var v82: map<T0, string> := map[v80 := v81];
				v82 := v82[v80 := v81];
				globalState.f3 := fm0(v55, globalState);
			}
			
			var v83: set<int> := {p0};
			var v84: array<bool> := new bool[2] [f12 ==> f14, v83 >= v83];
			v84[705] := if (f12) then fm0(v55, globalState) else f14;
			var v85 := "hbi";
			v84[705] := (if (fm6(f14, p0, globalState)) then fm9(p0, false, globalState) else v85) == v85;
		}
		var v86: set<bool> := {f12, false, f14};
		var v87 := "kjje";
		var v88: set<int> := {f13};
		for i8 := v0[3] to DC18(|v86|, v87, false, f14, v88).cf40 {
			if (f14) {
				var v89: array<bool> := new bool[21](i9 => f14);
				var v90: map<int, array<bool>> := map[i8 := v89];
				v90 := (v90 + v90) + v90;
				var v91: seq<int> := [i8];
				globalState.f9 := if (f12) then v0[3] else v91[p0] - f13;
				v90 := v90;
				v0[3] := f13 % f13;
				var v92: map<bool, int> := map[f14 := p0];
				var v93: seq<map<bool, int>> := [v92];
				var v94: map<bool, string> := map[f14 := v87];
				var v95: map<int, map<bool, int>> := map[p0 := v92];
				var v96: map<bool, bool> := map[f12 := f12];
				var v97: array<map<bool, int>> := new map<bool, int>[19] [v92, v92, v93[|v94|] + v92, v92, map[f12 := v0[3]] + v92, v92, v92 + v92, v92[false := f13], fm23(p0, globalState) + v92, v92, v92, (fm23(f13, globalState))[f14 := f13], v92 + v92, v92, v92 + v92, v92 + map[f14 := p0], if (|v96| in v95) then v95[|v96|] else map[false := i8], v92 + v92, v92];
				v97, v87 := v97, "l";
			} else {
				var v98: array<bool> := new bool[2](i10 => true);
				v98[627] := f12;
				v98[627] := f14;
				globalState.f1 := v0;
				var v99: seq<set<bool>> := [v86, v86 - v86];
				var v100 := 'q';
				var v101: seq<string> := [v87];
				globalState.f9, f12, v87, globalState.f8, v98[627] := |v99[-fm1(true, i8, -i8, globalState)]|, if (f14) then f12 else v98[627], v87[|(f11 + f11)| := v100], 'd' in v101[v0[3]], !f12;
				m0(globalState);
				var v102: array<C0> := new C0[10];
				var v103: map<bool, array<C0>> := map[v98[627] := v102];
				var v104 := DC23(v103);
				v104, v0[3], globalState.f9, v0[3] := v104.(cf55 := v103), f13, p0, -637;
			}
			
			v0[3] := f13;
			var v105: seq<bool> := [f12, f12];
			var v106: map<bool, bool> := map[false := f14];
			var v107: array<bool> := new bool[8] [f14, !(if (f14) then true else f12), !(-|v105[p0 := f12]| == v0[3]), if (f14 in v106) then v106[f14] else f12, f12, f14, f12, f12];
			v107[174] := v87 == v87;
			var v108: array<D8> := new D8[9];
			var v109 := DC19(f11);
			var v110 := DC22(v109);
			v108[659] := v110;
			v107[174], v0[3], v108[659], globalState.f9, globalState.f9 := !(v88 >= v88), if (f12) then f13 + 466 else i8, v110, f13, p0;
			var v111: map<array<int>, bool> := map[v0 := f12];
			v111 := v111[v0 := v87 == v87];
		}
		var v112 := 'x';
		var v113: seq<bool> := [!f12, fm0(v112, globalState)];
		var v114: map<bool, int> := map[!f12 := |v113|];
		var v115: seq<int> := [p0, v0[3], |"tyfnjxj"|, f13, if (f12 in v114) then v114[f12] else v0[3]];
		v0[3] := (v0[3] + |v115|) + v0[3];
	}
}

class C4 extends T0 {
	constructor (f11 : multiset<bool>, f12 : bool) {
		this.f11 := f11;
		this.f12 := f12;
	}
	
	function fm2(p0: multiset<int>, p1: int, p2: map<bool, int>, p3: bool, globalState: GlobalState): bool {
		0x380 != (--394 * 46)
	}
	function fm3(p0: bool, p1: map<string, string>, p2: bool, p3: int, globalState: GlobalState): int {
		0x9
	}
	function fm4(p0: int, p1: int, globalState: GlobalState): bool {
		(if (f12) then -538 else -0x272) == |(multiset{0x38c} + multiset{-0x13d, 0xa3})|
	}
	method m1(p0: set<bool>, globalState: GlobalState) returns (r0: array<map<int, char>>, r1: array<seq<bool>>, r2: int) {
		var v0 := -0x343;
		var v1 := "fahubdda";
		globalState.f9, globalState.f7, f11, globalState.f6 := -(fm1(f12, v0, v0, globalState) * v0), f12, f11[f12 := -(0x24 + v0)], fm5(f12, f12, f12, v0, globalState) + fm5(f12, f12, f12, |v1|, globalState);
		var v2: array<bool> := new bool[3];
		var v3: set<array<bool>> := {v2, v2, v2, v2};
		var v4: seq<int> := [v0, fm1(f12, v0, v0, globalState), v0, |"kd"| + v0];
		var v5: seq<set<bool>> := [p0, p0];
		var v6 := DC1(p0);
		var v7: array<set<bool>> := new set<bool>[29] [p0, p0, p0 * p0, p0 + p0, {f12, DC0(f12).cf0, f12, !f12}, p0 + p0, p0, p0 + p0, p0, p0 - v5[v0], p0, v6.cf1, p0, p0, p0, p0, p0, v5[-0xf9], p0, p0 * p0, p0, p0 - {f12, f12}, if (f12) then p0 else p0, {!false} + p0, p0 * {f12}, p0, {f12, f12}, p0 - p0, p0];
		v7[185] := p0 - p0;
		var v8: seq<set<array<bool>>> := [v3];
		v3, v4, v7[185], globalState.f3, v0 := v8[0x296 % -v0], v4, p0, f12, 0x39a;
		v2 := v2;
		var i0 := 0;
		while (f12)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			globalState.f9 := v0 - v0;
			globalState.f3 := !f12;
			if (f12) {
				var v9 := DC1(v7[185]);
				var v10: multiset<D1> := multiset{DC3(v9)};
				var v11 := DC3(v9);
				globalState.f3 := v10[v11 := v0] > (v10 - v10);
				var v12 := new C3(-(if (f12) then v0 else v0), f12, multiset{f12, f12} - multiset{!f12, f12}, f11 !! f11);
				var v13: map<int, int> := map[v12.f13 := --v12.f13];
				var v14: set<int> := {v12.f13, v12.f13, if (-708 in v13) then v13[-708] else v12.f13};
				v14, globalState.f9, globalState.f9, v4 := v14, v0, v0, if (v12.f14) then v4 else v4 + v4;
				var v15: array<array<bool>> := new array<bool>[19];
				v15[748] := if (f12) then v2 else v2;
				v0, v15[748] := v0 - (if (v12.f14) then v12.f13 else v0), v2;
				var v16: map<int, string> := map[v4[v12.f13] := v1];
				globalState.f9, globalState.f8, globalState.f8, v1 := v12.f13, v12.f14, v12.f14, if (v0 in v16) then v16[v0] else "tghqo" + v1;
			} else {
				var v17 := DC25(v1, f12, v0);
				var v18: seq<D9> := [v17, v17];
				var v19: C3 := new C3(v0, f12, f11, f12);
				var v20: seq<C3> := [v19, v19];
				var v21: map<bool, bool> := map[!v19.f14 := f12];
				var v22: seq<bool> := [true, f12, v19.f14];
				var v23 := 'v';
				var v24 := DC2(v19.f14, f12, false, v23, v23);
				var v25: T0 := new C1(|v22|, multiset{v24.cf4}, !true);
				var v26: set<int> := {v0};
				var v27: array<string> := new string[7](i1 => v1);
				var v28 := DC13(v19.f13, true, v19.f13, v27, v25.f12);
				var v29: map<bool, int> := map[v28.cf34 := v19.f13];
				var v30: C0 := new C0(true, v0, f11, v25.f12);
				var v31: seq<C0> := [v30];
				var v32: array<int> := new int[19] [v0, v0, (if (f12 in f11) then f11[f12] else v0) + v0, |v18|, v0, v0, v0, v0, |v20|, |v21|, |map[v19.f13 := v25]|, v19.f13 / v19.f13, |(v26 - v26)|, -v19.f13, if (v25.f12 in v29) then v29[v25.f12] else |v31|, v30.f18, 278, v19.f13, v19.f13];
				v32[355] := v0 / v0;
				v32[355] := v0;
				var v33 := new C0(v19.f14, |p0| % v19.f13, f11, |multiset{v4}| != v0);
				v2[162] := v30.f17;
				v2[162] := v30.f17;
				globalState.f8 := if (v19.f13 == 0x24f) then v25.f12 else f12;
				var v34: map<int, seq<int>> := map[-(v32[355] + |(seq(0x318, i2  => (v23)))|) := seq(0x150, i3  => (v19.f13))];
				v34 := if (if (!v2[162]) then v33.f17 else v33.f17) then v34 else map[v33.f18 := v4] + (map[v32[355] := v4])[0x38b := v4];
			}
			
			v1 := fm9(|multiset{v0}|, f12, globalState);
		}
		var v35 := 'm';
		var v36: array<char> := new char[5] [v35, v35, v35, v35, v35];
		var v37 := DC2(f12, f12, f12, v35, v35);
		var v38: array<array<bool>> := new array<bool>[19];
		var v39: T0 := new C3(v0, f12, f11, f12);
		var v40: set<T0> := {v39, v39};
		var v41 := DC21(true, v38, |v1|, v2, |v40|);
		v36[164] := fm14(v37, -v41.cf53, globalState);
		globalState.f9, v7[185], v36[164], v4 := |v4|, p0 * {true}, v35, fm24(globalState);
		for i4 := v0 / v0 to v0 {
			v7[185] := p0;
			var v42: array<string> := new string[11];
			v42[383] := v1;
			v42[383] := v1;
			var v43: seq<bool> := [f12];
			var v44: map<multiset<bool>, seq<bool>> := map[multiset([v39.f12, f12]) := v43];
			var v45: map<char, bool> := map[v36[164] := v39.f12];
			globalState.f6, v0, globalState.f7 := if (v39.f11 in v44) then v44[v39.f11] else v43, v0, v36[164] in (v45 + v45);
			var v46: map<T0, T0> := map[v39 := v39];
			v46 := map[v39 := v39];
		}
		var v47: array<map<int, char>> := new map<int, char>[5];
		r0 := v47;
		var v48: array<seq<bool>> := new seq<bool>[10];
		r1 := v48;
		r2 := v0;
	}
}

method Main() {
	var v1 := -351;
	var v2: seq<int> := [v1, 0xe7, v1];
	var v3 := 'p';
	var v4: map<map<int, bool>, char> := map[map v0 : int | v0 in v2 :: (v0 - v1) := (false) := v3];
	var v5: array<int> := new int[29](i0 => i0 % v1);
	var v6 := false;
	var v7: seq<bool> := [v6, true];
	var globalState := new GlobalState(v4, v5, true, false, v5, 0x2d4, v7, true, true, -0x297, false);
	m0(globalState);
	var i1 := 0;
	while (if (false) then v6 else !!!v6)
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		globalState.f8 := fm0(v3, globalState);
		var v8: map<int, bool> := map[v1 % v1 := fm0(v3, globalState)];
		var v10: map<bool, int> := map[v6 := v1];
		v8 := map v9 : int | v9 in [fm1(v6, v1, |v10|, globalState)] :: (v9 / |multiset{v1, |"kengs"|, |(set v11 : int | v11 in v8 :: (v11 - 206))|}|) := (multiset{v6, v6} > multiset{v6, v6});
		var v12: multiset<bool> := multiset{v6};
		var v13 := new C4(v12, v6);
		if (v6) {
			v5[968] := |(seq(284, i2  => (v3)))| - v1;
			v5[968] := v1;
			v5[968] := v1;
			globalState.f9 := v1;
			var v14 := new C4(v12, v6);
			globalState.f7 := v6;
		} else {
			var v15 := "fuaink";
			var v16: multiset<string> := multiset{v15, if (v6) then v15 else seq(0xd9, i3  => (v3)), v15 + "tdvl", v15, if (if (-v1 in v8) then v8[-v1] else v6) then v15 else DC25(v15, v6, |"gsmf"|).cf59};
			globalState.f9 := |v16|;
			v10 := v10[v6 := v13.fm3(v6, fm25(v3, v3, !v6, globalState), !v6, v1, globalState)];
			var v17 := new C2(v6, v12, v6);
			var v18: T0 := new C4(multiset{v17.f15}, v17.f15);
			var v19: map<int, T0> := map[v1 := v18];
			var v20 := DC28(v17, v17.f15, v18.f12, v18.f12);
			var v21: array<bool> := new bool[8] [true, v17.f15, v17.f15, v1 != |v19|, v20.cf67, true, v17.f15, v18.f12];
			v21[591] := v1 >= 489;
			v21[591] := v18.f12;
			globalState.f7 := v17.f15;
		}
		
	}
	var i4 := 0;
	while (v6)
		decreases 100 - i4
	{
		if (i4 >= 100) {
			break;
		}
		
		i4 := i4 + 1;
		var v22 := "lrpi";
		var v23: multiset<bool> := multiset{v6};
		var v24: T0 := new C3(|v22|, !v6, v23, v6);
		v24 := if (v24.f12) then v24 else v24;
		var v25: array<multiset<bool>> := new multiset<bool>[1](i5 => v24.f11);
		v25[509] := multiset{!true, v24.f12};
		var v26: C2 := new C2(v24.f12, fm22(v1, !v24.f12, globalState), v6);
		var v27: seq<C2> := [v26, v26, v26];
		var v28: C0 := new C0(v6, v1, multiset{v26.f15, v24.f12}, v24.f12);
		var v29: set<int> := {v28.f18};
		var v30: map<set<int>, C0> := map[v29 := v28];
		var v31: array<C0> := new C0[22] [v28, v28, v28, v28, v28, v28, v28, v28, if ({v1, v28.f18} in v30) then v30[{v1, v28.f18}] else v28, v28, v28, v28, v28, v28, v28, v28, v28, v28, v28, v28, v28, v28];
		var v32: map<bool, array<C0>> := map[v6 := v31];
		var v33: map<D9, bool> := map[DC23(v32) := v24.f12];
		var v34: map<bool, map<bool, array<C0>>> := map[v26.f15 := v32];
		var v35: map<bool, int> := map[v26.f15 := v1];
		var v36 := DC24(v22, v28.f17, v1);
		var v37: map<bool, bool> := map[v28.fm2(multiset{|v24.f11|}, v1, v35, false, globalState) := v36.cf57];
		var v38: multiset<map<bool, bool>> := multiset{fm16(v28.f18, globalState), v37};
		var v39: multiset<char> := multiset{v3, v3, v3, v3, v3};
		var v40: multiset<int> := multiset{v28.f18, 0x50, v28.f18, v1, if (v3 in v39) then v39[v3] else v28.f18};
		globalState.f3, globalState.f7, v25[509], v27, globalState.f7 := if (DC23(if (v28.f17 in v34) then v34[v28.f17] else v32) in v33) then v33[DC23(if (v28.f17 in v34) then v34[v28.f17] else v32)] else v6 <==> v28.f17, (v22 + "svi") <= v22, fm22(if (map[v28.f17 := v28.f17] in v38) then v38[map[v28.f17 := v28.f17]] else v28.f18, v24.f12, globalState), ([v26] + v27) + v27, !!v24.fm2(v40, v28.f18, fm23(v28.f18, globalState), v26.f15, globalState);
		var v41 := DC16(v1, v39[v3 := if (v26.f15 in v24.f11) then v24.f11[v26.f15] else v1] > v39, !v6);
		match (v41) {
			case DC16(cf36, cf37, cf38) =>
				v1 := v28.f18;
				var v42: map<bool, map<bool, bool>> := map[cf37 := v37];
				var v43 := DC30(v37);
				v24, v37 := v24, (if (fm0(v3, globalState)) then ((if (v28.f17 in v42) then v42[v28.f17] else v37)[v26.f15 := v24.f12])[cf37 := v24.f12] else v37) + v43.cf69;
				var v44: array<bool> := new bool[19] [v28.f17, v28.f17, cf38, v24.f12, v26.f15, v26.f15, v28.f17, cf37, v26.f15, v26.f15, v28.f17, v28.f17, v24.f12, v28.f17, !v26.f15, v24.f12, v28.f17, cf38, v26.fm2(multiset{v28.f18, v1}, v28.f18, v35, v26.f15, globalState)];
				var v45: seq<array<bool>> := [v44, v44, v44, v44, v44];
				var v46: set<array<bool>> := {v44, v44, v45[0x351]};
				v46 := v46 * v46;
				var v47 := DC2(cf38, v26.f15, v28.f17, v3, v3);
				var v48: map<int, char> := map[v28.f18 := fm14(v47, |v24.f11|, globalState)];
				var v49 := DC2(cf38, if (v24.f12 in v37) then v37[v24.f12] else true, false, v3, if (v28.f18 in v48) then v48[v28.f18] else v3);
				v3 := fm14(v49, v1, globalState);
			case DC15(cf35) =>
				var v50, v51 := v26.m5(v28.f18 % v28.f18, v22, globalState);
				v22 := v22 + ("alxsdoim" + v22);
				v51 := |((v7 + [v6, v26.f15]) + ([v24.f12, v24.f12, v26.f15] + v7))|;
				globalState.f3 := v28.f18 != |v22|;
		}
		
		var v52, v53 := v26.m5(fm1(v26.f15, v1, v28.f18, globalState), v22, globalState);
	}
	v2 := v2 + v2;
	globalState.f3 := v6;
	var v54 := "eyxck";
	var v55: set<int> := {v1};
	var v56 := DC18(|v7|, v54, v6, v6, v55);
	var v57: seq<string> := [v56.cf41];
	if (v1 <= |v57[v1]|) {
		var v58: multiset<bool> := multiset{v6};
		var v59: T0 := new C2(v6, v58, v6);
		var v60: multiset<T0> := multiset{v59, v59, v59, v59, v59};
		var v61: map<int, bool> := map[v1 := v6];
		var v62: map<int, int> := map[-v1 := v1];
		var v63: multiset<int> := multiset{v1, |v62|, v1};
		var v64 := new C1(|{v1, v1, 0x1d2, |v60|, |v61|}|, multiset(fm10(v59.f12, v59.f12, v6, globalState)), |v63| < fm1(v6, |v55|, v1, globalState));
		var v66: array<string> := new string[14] [v54, v54, seq(60, i6  => (v54[|(map v65 : int | v65 in [v64.f16] :: (v65 - v1) := (-0xf4))|])), v54, v54, "r", v54[v64.f16 := v3], v54, v54, v54, v54, "foxyq", v54, v54];
		var v67 := DC13(v64.f16, v6, -v64.f16, v66, v59.f12);
		v1 := v64.f16 * v67.cf30;
		v64.f16, v1 := |"qtkfwgn"| - (if (|v61| in v63) then v63[|v61|] else v1), v1;
		var v68 := DC25(v54 + "ak", v59.f12, v64.f16);
		match (v68) {
			case DC24(cf56, cf57, cf58) =>
				var v69: map<bool, bool> := map[v59.f12 := false];
				v69 := v69[|{cf57, v59.f12}| > --0x1f0 := if (v59.f12) then v59.f12 else v59.f12];
				var v70: map<bool, int> := map[v59.f12 := v64.f16];
				var v71: map<int, map<bool, int>> := map[v64.f16 := v70];
				v71 := v71[0xc := v70];
				v5[222] := |(cf56 + cf56)|;
				v5[222] := v64.f16;
				v3 := v3;
			case DC25(cf59, cf60, cf61) =>
				var v72: array<seq<bool>> := new seq<bool>[12];
				var v73: map<bool, array<seq<bool>>> := map[if (v64.f16 in v61) then v61[v64.f16] else v59.f12 := v72];
				v73 := v73[v59.f12 := v72];
				var v74: map<bool, int> := map[v59.f12 := v1];
				var v75: map<int, seq<bool>> := map[if (v64.f16 in v62) then v62[v64.f16] else |v74| := [cf60]];
				var v76: C3 := new C3(|(if (0x2a7 in v75) then v75[0x2a7] else v7)|, v6, fm22(v64.f16, v59.f12, globalState), v59.f12);
				var v77: map<C3, int> := map[v76 := |(seq(0x11b, i7  => (v3)))| - v64.f16];
				globalState.f9 := |v77|;
				globalState.f8 := v59.f12;
				var v78 := DC16(v64.f16, cf60, v6);
				v59.f12 := v78.cf38;
			case DC26(cf62) =>
				globalState.f8 := v59.f12;
				var v79 := new C0(!v59.f12, v64.f16 % v64.f16, v58, !v6);
				cf62 := -0x2b9;
				var v80: array<C3> := new C3[4];
				var v81: seq<multiset<bool>> := [v59.f11];
				var v82: C3 := new C3(v64.f16, v6, v81[v1], v59.f12);
				v80[648] := v82;
				v80[648] := v82;
			case DC23(cf55) =>
				var v83 := new C4(v59.f11, v59.f12);
				var v84: array<array<bool>> := new array<bool>[21];
				var v85: map<bool, array<array<bool>>> := map[v6 := v84];
				var v86: seq<multiset<bool>> := [v58];
				var v87: map<array<array<bool>>, seq<multiset<bool>>> := map[if (v6 in v85) then v85[v6] else v84 := v86];
				v87 := v87[v84 := v86];
				var v88 := DC20(v59.f12, v59.f12, false);
				v88 := DC20(v59.f12, v6, v83.fm4(v64.f16, v64.f16, globalState));
				var v89: set<bool> := {v59.f12, true, v59.f12, v6, !v59.f12};
				var v90, v91, v92 := v83.m1(v89, globalState);
		}
		
		v64.f16, globalState.f9, globalState.f9, globalState.f3, globalState.f8 := v64.f16, if (v1 >= v64.f16) then v64.f16 else v1, v1, v6, v6;
	} else {
		var v93: multiset<bool> := multiset{!v6, false, v6};
		var v94: map<bool, seq<int>> := map[multiset([true]) != v93 := v2];
		v94 := v94[v7[v1] := seq(226, i8  => (v1))];
		globalState.f8 := v1 > v1;
		v5[537] := v1 * v1;
		globalState.f9, v5[537] := if (v6) then v1 else v1, v2[0x39f];
		m0(globalState);
		var v95: map<bool, int> := map[true := v5[537]];
		var v96: map<int, int> := map[v5[537] := v5[537]];
		globalState.f9 := (if (!true in v95) then v95[!true] else fm1(v6, v5[537], -|v96|, globalState)) / v5[537];
	}
	
	globalState.f3 := |(if (v6) then v54 else ("gul")[v1 := 'd'])| == v1;
	if (v6) {
		var v97: set<bool> := {v6, v6};
		if (v97 < v97) {
			var v98: map<bool, bool> := map[v6 := fm0(v3, globalState)];
			var v99: map<seq<int>, map<bool, bool>> := map[v2 := v98];
			var v100 := DC29(v6);
			var v101: set<set<bool>> := {v97, fm18(|v99[[|v57|] := v98]|, (v100.(cf68 := v6)).cf68, globalState), {v6, v6}};
			globalState.f3, v1, v101, globalState.f7, globalState.f7 := !fm0(v3, globalState), v1 + |"vomrh"|, v101, v6, v1 > (-0x138 * v1);
			var v102: array<set<int>> := new set<int>[20](i9 => v55);
			v102[838] := v55;
			v102[838] := v55;
			var v103: map<int, bool> := map[v1 := v6];
			var v104: map<map<int, bool>, int> := map[v103 := v1];
			globalState.f9 := |(v104 + map[map[v1 := v6] := -|v102[838]|])|;
			v1 := v1;
			globalState.f7 := v6;
		} else {
			var v105: map<set<bool>, int> := map[v97 := v1];
			v5[579] := if (v97 in v105) then v105[v97] else v1;
			v5[76] := v1;
			var v106: C4 := new C4(multiset{v6, v6}, v6);
			var v107: multiset<C4> := multiset{v106, v106};
			globalState.f3, v5[579], v5[76], globalState.f9, globalState.f9 := v6, v1, 574 / v1, v1 / -498, |(if (map[v107 := v6] == map[v107 := v6]) then seq(543, i10  => (v1)) else seq(0x331, i11  => (v1)))|;
			globalState.f9 := 0x2f8;
			var v108: multiset<bool> := multiset{v6, v6, false, v6};
			var v109: map<int, multiset<bool>> := map[v1 := v108];
			var v110 := DC2(true, v6, v6, v3, v3);
			var v111: C0 := new C0(v6, v1, if (v5[579] in v109) then v109[v5[579]] else multiset{v6, v6, (v110.(cf2 := false, cf6 := v3)).cf4}, v6);
			var v112: map<C0, char> := map[v111 := v3];
			v112 := v112[v111 := 'd'];
			v5[579], v5[579], globalState.f9 := v111.f18 * (v111.f18 * 0x37), v5[579] / (if (v6 in v108) then v108[v6] else 0x21e), -v111.f18;
			var v113, v114, v115 := v106.m1(fm18(0x10d, v111.f17, globalState), globalState);
		}
		
		globalState.f3 := v1 != (v1 / v1);
		var v116: array<C0> := new C0[9];
		var v117: seq<array<C0>> := [v116];
		var v118: array<array<C0>> := new array<C0>[14] [v116, v116, v116, v116, v116, v117[-v1], v116, v116, v116, v116, v116, v116, v116, v116];
		v118[802] := if (v6) then v116 else v116;
		v118[802] := v116;
		var v119: map<char, int> := map[v3 := v1];
		var v120: set<map<char, int>> := {v119};
		var v121: map<set<map<char, int>>, set<int>> := map[v120 * v120 := v55 - v55];
		v55 := if ({v119, v119, map[v3 := v1]} in v121) then v121[{v119, v119, map[v3 := v1]}] else v56.cf44 - v55;
		if (fm0(v3, globalState)) {
			globalState.f9 := 678;
			var v122: array<bool> := new bool[28];
			v122 := v122;
			var v123: array<D8> := new D8[29];
			var v124: array<array<D8>> := new array<D8>[8] [v123, v123, v123, v123, v123, v123, v123, v123];
			v124[301] := v123;
			v124[301] := v123;
			var v125: map<char, bool> := map[v3 := v6];
			v125 := v125[v3 := v6];
			var v126: array<array<bool>> := new array<bool>[8];
			var v127 := DC21(v6, v126, v1, v122, 0x145);
			var v128: multiset<bool> := multiset{fm0(v3, globalState) ==> !v127.cf49, v6, v6, v6};
			v128 := (fm22(v1, v6, globalState) + v128) * (v128 * v128);
		} else {
			globalState.f3 := v6;
			globalState.f9 := -0x3a;
			m0(globalState);
			var v129: array<C4> := new C4[16];
			var v130: map<array<C4>, bool> := map[v129 := v6];
			v130 := v130[v129 := v6];
			v2 := [v1 * v1, 0x3cf, v1];
		}
		
	} else {
		globalState.f7 := v6;
		var v131: multiset<bool> := multiset{v6};
		var v132 := new C4(v131, v6);
		var v133: T0 := new C0(v6, v1, v131[v6 := v1], v6);
		v133 := v133;
		var v134 := DC20(v6, v6, v133.f12);
		var v135 := new C0(!v133.f12, v1, v133.f11, v134.cf48);
		var v136: map<bool, int> := map[v6 := v1];
		var v137: set<map<bool, int>> := {v136};
		var v138: set<bool> := {true, v135.f17};
		var v139 := new C0(v6, |v137|, v133.f11, v135.f18 >= |v138|);
	}
	
	var v140: map<string, bool> := map[("muwivve")[v1 := v3] := (fm26(v1, !false, globalState)).cf46];
	if (if ((seq(0x334, i12  => (v3))) in v140) then v140[seq(0x334, i12  => (v3))] else v55 < (set v141 : int | (0x237 <= v141) && (v141 < 0x3ae) :: (v141 + v1))) {
		var v142: array<bool> := new bool[7](i13 => v6);
		var v143: seq<array<bool>> := [v142];
		var v144: array<array<bool>> := new array<bool>[23] [v142, v142, v142, v143[fm1(v6, fm1(v6, v1, v1, globalState), |(seq(0x32a, i14  => (v3)))|, globalState)], v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142];
		var v145 := DC21(v6, v144, v1, v142, fm1(v6, 0x2f7, v1, globalState));
		var v146: multiset<bool> := multiset{v6, v145.cf49};
		var v147: C2 := new C2(true, v146, !v6);
		var v148: map<C2, array<bool>> := map[v147 := v142];
		var v149: array<array<bool>> := new array<bool>[23] [v142, v142, v142, v142, if (v147 in v148) then v148[v147] else v142, v142, v142, v142, (v145.(cf50 := v144)).cf52, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142, v142];
		v144[586] := v142;
		var v150: map<int, bool> := map[v1 := v147.f15];
		var v151: map<bool, array<bool>> := map[if (v1 in v150) then v150[v1] else v147.f15 := v142];
		var v152: multiset<int> := multiset{v1, v1, |v7|};
		var v153: map<bool, int> := map[v6 := v1];
		v144[586] := if (v147.fm2(v152, v1, v153, !v6, globalState) in v151) then v151[v147.fm2(v152, v1, v153, !v6, globalState)] else v142;
		v144[586][94] := v147.f15;
		v144[586][94] := v6;
		var v154 := new C0(v144[586][94], v1, v146, v6);
		if (v147.f15) {
			v6 := (|v153| * v154.f18) >= v154.f18;
			globalState.f8 := v144[586][94];
			var v155: set<char> := {v3, v3};
			var v156: array<string> := new string[27] [v54, seq(-0x12a, i15  => (v3)), "ecbkysbx", "mekb", v54, v54, v54, v54, v54, v54, v54, v54, v54, v54, "cnflkujmg", v54, v54, v54, fm9(|v155|, true, globalState), seq(0x1b0, i16  => ('f')), v54, v54, v54, "c", seq(325, i17  => (v3)), v54, v54];
			var v157 := DC13(|v153|, true, v1, v156, v154.f17);
			var v158: set<D5> := {v157, v157};
			var v159: seq<set<D5>> := [v158];
			var v160 := DC31(v159);
			var v161: seq<seq<set<D5>>> := [v159, [v158]];
			v147.f15 := v160.cf70 == v161[v154.f18];
			var v162 := new C1(-v154.f18, multiset{v6}, v154.f18 == v154.f18);
			var v163: multiset<string> := multiset{"ygwxn"};
			var v164: map<bool, multiset<string>> := map[v154.f17 := v163];
			v163 := (if (v6 in v164) then v164[v6] else v163) - v163;
		} else {
			m0(globalState);
			v5[527] := v154.f18;
			v5[527] := if (v147.fm2(v152, v154.f18, map[true := v154.f18], v6, globalState) in v146) then v146[v147.fm2(v152, v154.f18, map[true := v154.f18], v6, globalState)] else v154.f18;
			var v165: multiset<char> := multiset{v3};
			var v166: map<bool, multiset<char>> := map[v144[586][94] := v165];
			var v167: set<bool> := {!v6};
			var v168: array<int> := new int[14] [-v154.f18, if (v154.f17) then -0x1a4 else |(if (false in v166) then v166[false] else v165)|, -(|multiset{v154.f17, v6}| - v154.f18), v1, |v167|, v1, v1 - 844, v154.fm12(globalState) * v154.f18, 620, |map[0x262 := v154.f18]|, v154.f18, v1 % v154.f18, v5[527], v1 / |v54|];
			globalState.f1 := v168;
			var v169: set<char> := {v3, v3};
			var v170: map<int, multiset<int>> := map[|v169| := v152];
			var v171: map<int, multiset<bool>> := map[v1 := v146];
			var v172: array<string> := new string[20];
			var v173 := DC13(v1, false, |v171|, v172, v154.f17);
			v54, v5[527], globalState.f7, v152 := v54 + v54, v154.f18, v147.fm2(multiset{v154.f18, |(if (|v55| in v170) then v170[|v55|] else v152)|}, v173.cf32, (fm23(-v5[527], globalState))[false := v1], v144[586][94], globalState), (multiset(v2) * fm13(v54[v154.f18], false, globalState)) + v152;
			globalState.f9 := v154.fm12(globalState);
		}
		
		var v174 := DC14();
		match (v174) {
			case DC13(cf30, cf31, cf32, cf33, cf34) =>
				var v175: C4 := new C4(v146, v154.f18 in v55);
				v5[539] := 0x87;
				var v176: C3 := new C3(cf32, v154.f17, v146, false);
				var v177: map<int, int> := map[v176.f13 := cf30];
				var v178: map<C3, int> := map[v176 := |v177|];
				var v179: multiset<multiset<int>> := multiset{multiset{if (v176 in v178) then v178[v176] else v154.f18, fm1(!v147.f15, v1, v1, globalState), v1}, v152 - v152, v152, v152};
				var v180 := DC24(v54, v147.f15, v1);
				v3, v175, v55, v5[539], v179 := v3, v175, v55 * v55, |v180.cf56|, v179;
				globalState.f9 := v154.f18;
				var v181: seq<multiset<bool>> := [v146];
				var v182: T0 := new C4(v181[v5[539]], !cf34);
				var v183: map<T0, string> := map[v182 := v54[-v176.f13 := v3]];
				var v184: seq<map<bool, int>> := [fm23(|v2|, globalState), fm23(v5[539], globalState)];
				var v185: multiset<T0> := multiset{v182, v182};
				var v186: map<int, array<array<bool>>> := map[0x1d3 := v149];
				var v187: set<bool> := {v182.f12};
				var v188: set<set<bool>> := {v187};
				v183, cf30, v144[586][94], v145, v184 := v183, if (v182 in v185) then v185[v182] else v154.f18 / 222, (v54 + v54) < (seq(0x325, i18  => (v3))), DC21(v147.f15, if (v5[539] in v186) then v186[v5[539]] else v144, v5[539] % v154.f18, v144[586], |v188| * fm1(cf31, cf30, cf30, globalState)), v184 + v184;
				var v189, v190 := v176.m2("jb", v176.f14, cf33, globalState);
			case DC14() =>
				v147.f15 := true;
				globalState.f3 := v147.f15;
				v153 := v153[v154.fm2(v152, |v143|, v153, v6, globalState) := v154.f18];
				v6, v152, globalState.f3, v144[586][94] := !!false, v152 - (v152 + v152), v154.f17, !v144[586][94];
			case DC12(cf29) =>
				var v191 := new C2(v147.f15, v146, v144[586][94]);
				globalState.f9 := |v54|;
				var v192, v193 := v191.m5(0x2fd, v54 + v54, globalState);
				v193 := -v1;
		}
		
	} else {
		var v194: array<D6> := new D6[6];
		var v195: map<bool, bool> := map[v6 := v6];
		var v196: C1 := new C1(v1, (multiset(v7))[v6 := |v195|], v6);
		var v197: multiset<int> := multiset{v1, v196.f16};
		var v198: map<C1, multiset<int>> := map[v196 := v197];
		var v199 := DC15(if (v196 in v198) then v198[v196] else v197);
		v194[601] := v199;
		v194[601] := v199;
		var v200 := DC2(false, v6, v6, v3, v3);
		var v201 := DC3(v200);
		var v202 := DC3(v201);
		var v203 := DC3(v202);
		var v204 := DC3(v203);
		v204 := v204;
		var v206: map<int, map<bool, bool>> := map[|v195| := v195];
		v1 := |((map v205 : int | (0x167 <= v205) && (v205 < -0x36f) :: (v205 / |multiset(v7)|) := (v195[!v6 := v6])) + v206[|v54| := v195])| / |v54|;
		var v207: array<bool> := new bool[15](i19 => v7[v1]);
		v207[717] := v6;
		var v208: set<bool> := {true};
		v207[717] := ({v6, v6, v6} != v208) ==> (if (v6) then v6 else v6);
		var v209: set<string> := {v57[v196.f16]};
		if ((v209 * (set v210 : string | v210 in map[v54 := v6] :: (v210))) < v209) {
			v7 := (v7 + v7)[v1 := v6];
			v54 := seq(32, i20  => (v3));
			var v211: multiset<bool> := multiset{v207[717], v6, v207[717]};
			var v212 := new C4(v211, false);
			var v213 := new C0(v207[717], 449, v211, v207[717]);
			v6 := v213.f17;
		} else {
			var v214: multiset<bool> := multiset{v207[717]};
			var v215: array<array<bool>> := new array<bool>[20] [v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207, v207];
			var v216 := DC21(false, v215, v1, v207, |{v207[717], v207[717]}|);
			var v217: C4 := new C4(v214, v216.cf49);
			var v218: map<bool, C4> := map[true := v217];
			var v219: map<int, bool> := map[v196.f16 := v207[717]];
			v218 := v218[!v6 && (if (v196.f16 in v219) then v219[v196.f16] else v6) := v217];
			v214 := v214;
			globalState.f1 := if (v217.fm4(v196.f16, v196.f16, globalState)) then v5 else v5;
			globalState.f9 := v196.f16;
			v196.f16 := v1;
		}
		
	}
	
	var v220: set<bool> := {v6};
	var v221: C0 := new C0(v6, |v220|, multiset{v6, v6}, v6);
	var v222: multiset<bool> := multiset{v221.f17, v6, v6};
	var v223: C3 := new C3(v1, v221.f17, v222, v6);
	var v224: map<int, C3> := map[v1 := v223];
	var v225: seq<multiset<bool>> := [v222, v222, v222, fm22(|v224|, v6, globalState)];
	var v226: map<C0, bool> := map[v221 := v225[v221.f18] !! fm22(v223.f13, v223.f14, globalState)];
	v226 := v226[v221 := !v223.f14];
	var v227: array<string> := new string[27](i22 => v54);
	forall i21 | 0 <= i21 < v227.Length {
		v227[i21] := v54;
	}
	var v228: C1 := new C1(v223.f13, v222, v221.f17);
	var v229: seq<C1> := [v228, v228];
	var i23 := 0;
	while (!!((v229 + [v228])[v1 := v228] == v229))
		decreases 100 - i23
	{
		if (i23 >= 100) {
			break;
		}
		
		i23 := i23 + 1;
		var v230: T0 := new C1(v223.f13, multiset{v6}, v221.f17);
		var v231: set<T0> := {v230};
		v57 := [seq(0x1e5, i24  => (v3)), v54, ("ot")[-|v231| := 'r'], v54];
		var v233: array<map<seq<int>, int>> := new map<seq<int>, int>[8](i25 => map[[v221.f18, v1] := v223.f13] + (map v232 : seq<int> | v232 in multiset{v2[v228.f16 := v1], v2} :: (v232) := (v223.f13)));
		var v234: map<seq<int>, int> := map[[|v54|] := v1];
		var v235 := DC34(v234);
		v233[824] := v235.cf76;
		v233[824] := (v234 + map[v2 := v1]) + (v234 + (map v236 : seq<int> | v236 in (seq(0x52, i26  => (v2))) :: (v236) := (v221.f18)));
		var v237: multiset<int> := multiset{v228.f16, 0x32d, v221.f18, |v7|};
		var v238: map<bool, int> := map[v230.f12 := v223.f13];
		var v239: seq<seq<int>> := [v2, if (v228.fm2(v237, 0x2a1, v238, v223.fm6(v223.f14, v1, globalState), globalState)) then v2 else v2, v2 + v2];
		var v240: seq<seq<seq<int>>> := [v239, v239 + v239];
		var v241 := DC24(v54, true, 0x2c3);
		v239 := (v240[v223.f13])[v221.f18 := [v241.cf58] + v2];
		v230.f12 := !v6;
	}
	m0(globalState);
	if (v6) {
		v54, v2 := v54, v2;
		var v242: map<int, bool> := map[|v54| := v221.f17];
		var v243: T0 := new C0(v223.f14, v221.f18, v222, v6);
		var v244: map<char, T0> := map[v3 := v243];
		var v245: seq<map<char, T0>> := [v244[v3 := v243]];
		var v246 := new C1(0x34f % |v54|, multiset{if (|v245| in v242) then v242[|v245|] else v243.f12, v223.f14}, !v221.f17);
		if (v243.f12 <== v223.fm6(!!v6, v221.fm12(globalState), globalState)) {
			var v247: array<bool> := new bool[10];
			var v248: seq<array<bool>> := [v247];
			var v249: multiset<int> := multiset{v228.f16, v1, |v248|, -v228.f16, v228.f16};
			var v250: map<int, multiset<int>> := map[v1 := v249];
			var v251: array<multiset<int>> := new multiset<int>[12] [fm13(v3, !v221.f17, globalState), v249, v249, v249, if (v221.f18 in v250) then v250[v221.f18] else multiset{v228.f16}, v249, v249[v223.f13 := v223.f13], v249, multiset{|v2|} + v249, multiset{v228.f16, v246.f16, |"co"|}, if (v243.f12) then v249 else v249, (multiset{0xa2, v228.f16})[-v1 := v228.f16]];
			v251[845] := fm13('h', v223.f14, globalState);
			v251[845] := v249;
			var v252: map<seq<int>, int> := map[v2[v246.f16 := v221.f18] := 0x162];
			var v253 := DC34(v252);
			var v254: map<D13, bool> := map[v253 := v223.f14];
			var v255: array<D8> := new D8[10];
			var v256: map<map<D13, bool>, array<D8>> := map[v254 := v255];
			v5[120] := v221.f18 % v228.f16;
			v256, v243.f12, v223, v5[120] := v256, v223.f13 != v221.f18, if (v7[v221.f18]) then if (v221.f17) then v223 else v223 else v223, v228.f16;
			var v257: C2 := new C2(v243.f12, v243.f11, v243.f12);
			var v258 := DC32(v221.f18 / |map[-0x35d := v246]|, v257, v221, v223.f13);
			var v259: map<int, C0> := map[v246.f16 := v221];
			var v260: seq<C0> := [v221, if (v221.f18 in v259) then v259[v221.f18] else v221, v221];
			v258 := DC32(v228.f16, v257, v260[|v220|], -|v54|);
			var v261 := new C3(-v223.f13, v243.f12, multiset{false, v221.f17} + v222, v221.f17);
			var v262: array<int> := new int[13] [v223.f13 + v228.f16, v223.f13, v228.f16 + -v246.f16, fm1(v257.f15, v223.f13, |v55|, globalState), -(v223.f13 * v221.f18), v228.f16, v223.f13, -v228.f16, v223.f13, v228.f16, if (!!v257.f15) then |v54| else v1, |v7|, |v54| - v221.f18];
			globalState.f3, globalState.f1, v1 := !(true ==> v243.f12), v262, if (v261.f13 <= -0x36c) then v1 + v221.f18 else 0x3d3;
		} else {
			v6 := v243.f12;
			var v263 := DC34(map[v2 := v228.f16]);
			var v264: map<seq<int>, int> := map[v2 := v246.f16];
			var v265: array<D13> := new D13[4] [v263, DC34(v264), v263, v263];
			var v266: map<bool, array<D13>> := map[v223.f14 := v265];
			var v267: array<bool> := new bool[10] [v55 > {|v220|, v1, v228.f16}, v223.f13 > 0x215, v221.f17 && v223.f14, v6, v223.f14, v223.f14, v243.f12, !v243.f12 !in v266, v7[fm1(v6, v1, v228.f16, globalState)], v221.f17];
			v267[622] := v223.f14;
			globalState.f7, globalState.f9, v267[622], v3 := (v2 <= v2) || v221.f17, v221.f18, v221.f17, v3;
			var v268 := v228.m6(globalState);
			var v269: map<bool, int> := map[false := v223.f13];
			var v270: map<bool, bool> := map[v223.f14 := v223.f14];
			var v271: multiset<int> := multiset{|v2|, v221.fm12(globalState), |v270|, v228.f16, v221.f18};
			var v272: map<array<int>, bool> := map[v5 := true];
			var v273: map<int, array<bool>> := map[if (v243.fm2(v271, -|v270|, v269, v267[622], globalState) in v269) then v269[v243.fm2(v271, -|v270|, v269, v267[622], globalState)] else |v272| := v267];
			var v274: set<char> := {v3};
			var v275: map<bool, array<bool>> := map[v274 == v274 := if (v223.f14) then v267 else v267];
			v273, globalState.f3, v228.f16, v228.f16, v275 := v273, if (v223.f14) then v267[622] else v243.f12, |v54|, v221.f18 + -v228.f16, v275 + v275;
			v270 := v270[v221.f17 := v223.f14];
		}
		
		v223 := v223;
		var v276 := DC2(v221.f17, v223.f14, v221.f17, v3, v3);
		var v277: array<char> := new char[12] [fm14(v276, v246.f16, globalState), v3, v3, v3, fm14(v276.(cf3 := v223.f14), v221.f18, globalState), v3, v3, v3, v3, v3, if (!!v223.f14) then v54[|map[v223.fm6(v6, v223.f13, globalState) := v1]|] else v3, v3];
		v277[314] := fm14(v276, v221.f18, globalState);
		var v278: C2 := new C2(v242 == v242, v243.f11, v243.f12);
		var v279: multiset<int> := multiset{|v7|, |v2|, v228.f16, v246.f16};
		var v280: map<int, multiset<int>> := map[-0x1d3 := v279];
		var v281: array<multiset<int>> := new multiset<int>[12] [v279[v228.f16 := v1], v279 * v279, v279, v279, v279, multiset(v2), if (v221.f18 in v280) then v280[v221.f18] else v279, v279, v279, (fm13(v3, v221.f17, globalState))[v221.f18 := -|v54|], multiset{v246.f16} + v279, v279];
		v281[81] := v279;
		globalState.f9, v277[314], v278, v281[81] := v228.f16, v3, v278, v279;
	} else {
		var v282: T0 := new C1(-|v222|, v222, true);
		var v283: map<T0, int> := map[v282 := if (v221.f17) then |(v57[|v55|])[v221.f18 := v3]| else |v7|];
		v283 := v283[v282 := fm1(v7[v221.f18], -v221.f18, v223.f13, globalState)];
		var v284: map<bool, bool> := map[!v223.f14 := false];
		var v285: seq<map<bool, bool>> := [v284, v284[v282.f12 := !v223.fm6(v223.f14, if (v223.f14 in v222) then v222[v223.f14] else v1, globalState)], v284];
		v285 := [v284 + v284[v221.f17 := v6], v284, v284, v284 + v284];
		var v286: array<bool> := new bool[25](i27 => true);
		var v287: map<bool, array<bool>> := map[false := v286];
		var v288: array<array<bool>> := new array<bool>[22] [v286, v286, v286, v286, v286, v286, v286, v286, v286, v286, v286, v286, v286, v286, v286, if (v6 in v287) then v287[v6] else v286, v286, v286, v286, v286, v286, v286];
		var v289 := DC21(v221.f17, v288, v221.f18, v286, v223.f13 % v221.f18);
		match (v289) {
			case DC20(cf46, cf47, cf48) =>
				var v290: seq<set<bool>> := [v220];
				v1 := -|(v290[v228.f16] - (v220 * v220))|;
				var v291: map<int, char> := map[803 := v3];
				v3 := if (v228.f16 in v291) then v291[v228.f16] else v3;
				var v292: multiset<char> := multiset{v3, v3, v3, v54[v221.f18]};
				var v293 := DC35(multiset{v3});
				var v294: seq<multiset<char>> := [v292];
				cf47 := (if (true) then v292 else v293.cf77) >= v294[|v54|];
				globalState.f3 := v223.f14;
			case DC21(cf49, cf50, cf51, cf52, cf53) =>
				cf51 := |v57[if (v282.f12) then v1 else v221.f18]|;
				v282.f12 := cf49;
				cf52 := v286;
				var v295: multiset<seq<int>> := multiset{seq(0xeb, i28  => (v228.f16))};
				cf51 := if ((seq(0x383, i29  => (--0xc6))) in v295) then v295[seq(0x383, i29  => (--0xc6))] else v221.f18 / v221.f18;
			case DC19(cf45) =>
				m0(globalState);
				globalState.f9 := if (v282.f12) then v223.f13 else v223.f13;
				globalState.f9 := v221.f18 * v228.f16;
				v282 := v282;
			case DC22(cf54) =>
				v282.f12 := v282.f12;
				var v296 := DC2(v282.f12, v221.f17, v223.f14, v3, v3);
				v3 := fm14(v296, v1, globalState);
				var v297: map<bool, int> := map[v223.f14 := v221.f18];
				var v298: array<seq<int>> := new seq<int>[1];
				var v299: map<array<seq<int>>, map<bool, int>> := map[v298 := v297];
				var v300: seq<map<bool, int>> := [v297, v297];
				var v301: map<int, int> := map[|v54| := v221.f18];
				var v302: multiset<int> := multiset{v228.f16, v223.f13};
				var v303: map<set<int>, int> := map[{if (|v302| in v301) then v301[|v302|] else v228.f16, v1} := v221.f18];
				var v304: array<map<bool, int>> := new map<bool, int>[19] [v297, v297, if (v298 in v299) then v299[v298] else v297, v297, fm23(v228.f16, globalState), v297, map[v221.f17 := v223.f13], v297, v297 + v297, v297 + v297, map[v221.f17 := v221.f18], v297, v297, v297, v297, v297 + v297, v297, v297, v297 + v300[|v303|]];
				var v305: array<D9> := new D9[19](i30 => DC26(v1));
				v228.f16, v228, v228.f16, globalState.f8, v304 := |multiset{v305}|, v228, v221.f18, v223.f14, v304;
				var v306: map<bool, C1> := map[true := v228];
				var v307: array<C1> := new C1[4] [if (v6 in v306) then v306[v6] else v228, v228, v228, v228];
				v307 := v307;
		}
		
		m0(globalState);
		v286 := v286;
	}
	
	globalState.f7 := v6;
	var i31 := 0;
	while (v221.f17)
		decreases 100 - i31
	{
		if (i31 >= 100) {
			break;
		}
		
		i31 := i31 + 1;
		v228.f16 := v1;
		var v308: C2 := new C2(v221.f17, v222, v223.f14);
		var v309 := DC32(|v54|, v308, v221, v223.f13);
		var v310: map<seq<int>, D12> := map[v2 := v309];
		var v311: map<int, int> := map[v228.f16 := |{v308.f15}|];
		var v312: map<int, int> := map[v223.f13 := |v311|];
		v310 := v310[v2 := DC32(v228.f16, v308, v221, if (-0xfc in v311) then v311[-0xfc] else v221.f18)];
		var v313: seq<C0> := [v221];
		var v314: multiset<int> := multiset{|v313|};
		var v315 := DC25(v54, v308.f15, v221.f18);
		var v316: map<bool, int> := map[v228.fm2(v314, v228.f16, map[v315.cf60 := v228.f16], v223.f14, globalState) := 0x175];
		var v317: C4 := new C4(v222, v221.f17);
		var v318: map<int, C4> := map[if (v308.f15 in v316) then v316[v308.f15] else 0x138 := v317];
		v318 := v318[615 := if (v223.f14) then v317 else v317];
		var v319, v320 := v223.m2(v54, v223.f14, v227, globalState);
	}
}