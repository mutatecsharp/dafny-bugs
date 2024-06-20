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
datatype D0 = DC1 | DC2(cf1: int) | DC0(cf0: int)
datatype D1 = DC4(cf3: int, cf4: int) | DC3(cf2: bool)
datatype D2 = DC6 | DC7(cf6: int) | DC5(cf5: array<bool>)
datatype D3 = DC9(cf8: bool, cf9: int, cf10: bool) | DC10(cf11: int) | DC8(cf7: seq<bool>) | DC11(cf12: D3)
datatype D4 = DC13(cf14: int, cf15: bool, cf16: int, cf17: seq<bool>) | DC12(cf13: char)
datatype D5 = DC14(cf18: multiset<int>)
datatype D6 = DC16(cf20: bool, cf21: bool, cf22: seq<int>) | DC15(cf19: multiset<seq<bool>>)
datatype D7 = DC17(cf23: set<int>)
datatype D8 = DC18(cf24: map<array<int>, bool>)
datatype D9 = DC20(cf26: bool, cf27: char, cf28: int) | DC19(cf25: string)
datatype D10 = DC22(cf30: string, cf31: bool, cf32: bool, cf33: array<bool>, cf34: int) | DC21(cf29: map<bool, int>)
datatype D11 = DC24(cf36: bool, cf37: bool, cf38: int) | DC25(cf39: bool, cf40: int) | DC26(cf41: bool) | DC23(cf35: array<map<bool, bool>>)
datatype D12 = DC28(cf43: int, cf44: int, cf45: bool, cf46: bool, cf47: int) | DC27(cf42: seq<seq<bool>>)
datatype D13 = DC29(cf48: set<array<int>>)
datatype D14 = DC30(cf49: map<int, int>)
datatype D15 = DC31(cf50: array<array<bool>>)
datatype D16 = DC33(cf52: bool, cf53: int, cf54: multiset<bool>, cf55: int, cf56: set<bool>) | DC32(cf51: array<int>)
datatype D17 = DC34(cf57: seq<D9>)
datatype D18 = DC36(cf59: int, cf60: int, cf61: int) | DC35(cf58: seq<map<D11, bool>>)
datatype D19 = DC38(cf63: bool, cf64: bool) | DC37(cf62: map<char, array<bool>>) | DC39(cf65: D19)
datatype D20 = DC41(cf67: int) | DC40(cf66: map<bool, D12>)
datatype D21 = DC43(cf69: bool, cf70: bool, cf71: seq<bool>) | DC44(cf72: bool) | DC42(cf68: array<D0>)
datatype D22 = DC46(cf74: int, cf75: int, cf76: bool) | DC47(cf77: int, cf78: bool, cf79: bool, cf80: int, cf81: int) | DC45(cf73: array<map<bool, D11>>)
datatype D23 = DC49(cf83: int, cf84: int, cf85: int) | DC48(cf82: C0) | DC50(cf86: D23)
datatype D24 = DC51(cf87: T0)
datatype D25 = DC52(cf88: multiset<seq<int>>)
datatype D26 = DC54(cf90: int, cf91: int, cf92: bool, cf93: int) | DC53(cf89: seq<C5>) | DC55(cf94: D26)
datatype D27 = DC57(cf96: seq<array<int>>) | DC56(cf95: C2)
datatype D28 = DC58(cf97: map<map<bool, bool>, bool>)
trait T0 {
	const f4 : bool
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> 
	function fm7(globalState: GlobalState): int 
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) 
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) 
}

trait T1 extends T0 {
	var f5 : array<array<bool>>
	const f6 : int
	function fm8(p0: int, p1: bool, p2: seq<char>, p3: seq<int>, globalState: GlobalState): bool 
	function fm9(p0: bool, p1: bool, p2: int, globalState: GlobalState): int 
}

class GlobalState {
	const f0 : multiset<bool>
	const f1 : int
	const f2 : bool
	const f3 : int
	constructor (f0 : multiset<bool>, f1 : int, f2 : bool, f3 : int) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
	}
	
}

class C0 {
	const f12 : int
	constructor (f12 : int) {
		this.f12 := f12;
	}
	
	function fm17(p0: bool, globalState: GlobalState): bool {
		true
	}
}

class C1 extends T1 {
	const f15 : int
	var f16 : int
	constructor (f15 : int, f16 : int, f5 : array<array<bool>>, f6 : int, f4 : bool) {
		this.f15 := f15;
		this.f16 := f16;
		this.f5 := f5;
		this.f6 := f6;
		this.f4 := f4;
	}
	
	function fm8(p0: int, p1: bool, p2: seq<char>, p3: seq<int>, globalState: GlobalState): bool {
		f4
	}
	function fm9(p0: bool, p1: bool, p2: int, globalState: GlobalState): int {
		safeDivisionInt(f15, -f6)
	}
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> {
		if (f4 || true) then multiset(seq(abs(0x2b), i0  => (f6))) * multiset{f6, 0x1d5} else multiset{f16} + multiset{f6}
	}
	function fm7(globalState: GlobalState): int {
		match DC3(!f4) {
			case DC4(cf3, cf4) => |"iyylxfgh"|
			case DC3(cf2) => |(if (true) then [cf2, !cf2] else [cf2, cf2])|
		}
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) {
		var v0: set<int> := {f15};
		var v1: multiset<int> := multiset{f16, |v0|, f16};
		var v2 := DC9(f4, f6, !f4);
		match (if (v1 < multiset{f6}) then v2 else v2.(cf8 := false)) {
			case DC9(cf8, cf9, cf10) =>
				var v3 := 'k';
				if (-0x33c <= (|fm37(cf8, !cf10, p0, fm0(v3, globalState), globalState)| - cf9)) {
					var v4: array<array<int>> := new array<int>[4];
					var v5: array<int> := new int[22](i0 => i0 * p0);
					v4[safeIndex(500, v4.Length)] := v5;
					v4[safeIndex(500, v4.Length)] := v5;
					var v6 := "tqnumrpc";
					v6 := (seq(abs(435), i1  => (v6[safeIndex(|DC16(f4, f4, seq(abs(0x236), i2  => (|{0x2dc, |map[cf9 := p0]|}|))).cf22|, |v6|)])))[safeIndex(0x342, |(seq(abs(435), i1  => (v6[safeIndex(|DC16(f4, f4, seq(abs(0x236), i2  => (|{0x2dc, |map[cf9 := p0]|}|))).cf22|, |v6|)])))|) := v3];
					var v7: seq<int> := [cf9];
					var v8: map<bool, int> := map[fm8(fm7(globalState), f4, v6, v7, globalState) := |v7|];
					r0 := v8;
					cf8 := cf10;
					var v9: seq<bool> := [f4];
					v9 := [cf8] + v9;
				} else {
					var v11: seq<bool> := [f4, cf8];
					var v12: map<map<int, bool>, int> := map[map v10 : int | v10 in multiset{f15} :: (safeModuloInt(v10, p0)) := (f4) := fm3(cf9, v11[safeIndex(cf9, |v11|)], globalState)];
					v12 := v12 + v12;
					var v13 := new C0(cf9);
					f16 := safeDivisionInt(f15, v13.f12);
					var v14: seq<int> := [-p0, 886];
					f16 := v14[safeIndex(cf9, |v14|)] - 0x2f8;
					var v15: multiset<bool> := multiset{!!cf8};
					var v16 := "tplcf";
					f16 := cf9 - |(fm37(cf10, f4, if (f4 in v15) then v15[f4] else 150, f4, globalState) + v16[safeIndex(f16, |v16|) := v3])|;
				}
				
				cf8 := cf8;
				f5 := f5;
				if (cf10) {
					cf8 := !cf8;
					f16 := safeModuloInt(p0, f6);
					cf9 := cf9 - 905;
					var v17: C0 := new C0(cf9 - cf9);
					v17 := new C0(0x399);
					var v18: seq<bool> := [f4, true, !cf10];
					var v19: seq<int> := [f6 * |DC27([v18]).cf42|];
					v19 := v19;
				} else {
					f16 := -cf9;
					var v20: array<seq<int>> := new seq<int>[8](i3 => [cf9, |"lyl"|]);
					var v21 := DC16(cf10, cf10, [p0]);
					v20[safeIndex(686, v20.Length)] := v21.cf22;
					var v22: seq<int> := [f6];
					v20[safeIndex(686, v20.Length)] := v22;
					var v23: map<bool, int> := map[f4 := f16];
					var v24 := "tlcjkdars";
					cf10 := if (fm8(if (cf8 in v23) then v23[cf8] else |v24|, true, v24, v22, globalState)) then cf10 else !f4;
					var v25: array<int> := new int[18];
					v25 := v25;
					v3 := if (f4) then v3 else v3;
				}
				
			case DC10(cf11) =>
				if (f4) {
					var v26 := new C0(fm7(globalState));
					var v27 := "nfgsjgori";
					var v28 := DC19(v27);
					var v29: multiset<D9> := multiset{v28.(cf25 := v27), v28};
					var v30: set<bool> := {true, f4};
					var v31: seq<int> := [|v30|];
					var v32: map<multiset<D9>, seq<int>> := map[v29 := v31];
					var v34: set<multiset<D9>> := {multiset([v28]), multiset{v28}, v29, v29};
					var v35: seq<bool> := [f4];
					var v36: array<bool> := new bool[11];
					var v37: map<array<bool>, int> := map[v36 := cf11];
					var v38: array<int> := new int[14] [f16, |(v32 + (map v33 : multiset<D9> | v33 in v34 :: (v33) := ([v26.f12, -0x267])))|, p0, cf11, f6, safeDivisionInt(|v35|, p0), f6, if (fm8(f6, v35[safeIndex(fm3(0x82, false, globalState), |v35|)], v27, v31, globalState)) then f6 else |v31|, f16, p0, f16, -(if (v36 in v37) then v37[v36] else |v35|), v26.f12, safeDivisionInt(cf11, p0)];
					v38[safeIndex(159, v38.Length)] := safeDivisionInt(-0x156, fm3(f6, f4, globalState));
					v38[safeIndex(159, v38.Length)] := fm3(f6 - v26.f12, f4, globalState);
					var v39: seq<set<bool>> := [v30];
					var v40: map<seq<bool>, seq<set<bool>>> := map[v35 + v35 := v39];
					var v41: map<int, set<bool>> := map[-f15 := v30];
					v40 := v40[v35 := [if (0x174 in v41) then v41[0x174] else v30]];
					var v42 := true;
					var v43: multiset<bool> := multiset{v42, v42, {f4, f4, v42, fm8(v38[safeIndex(159, v38.Length)], f4, v27, v31, globalState), v42} > v30};
					f16, v42, v0, v43 := v26.f12, false, set v44 : int | (0x2a9 <= v44) && (v44 < 893) :: (v44 * v26.f12), (multiset{f4, false})[false <== f4 := abs(|v27| + v26.f12)];
					var v45 := DC17(v0);
					var v46: set<string> := {seq(abs(0x2c6), i4  => ('w'))};
					var v47: map<D7, bool> := map[v45 := v46 !! v46];
					v47 := v47[v45 := f4];
				} else {
					var v48 := "qrngovhln";
					v48 := v48 + v48;
					var v49: multiset<string> := multiset{v48, v48, v48};
					cf11 := safeModuloInt(fm3(f6, f4, globalState), if (v48 in v49) then v49[v48] else f15);
					var v50: array<int> := new int[9];
					v50[safeIndex(442, v50.Length)] := -|v0|;
					var v51: set<bool> := {f4};
					v50[safeIndex(442, v50.Length)] := safeModuloInt(fm7(globalState), -|map[p0 := v51]|);
					cf11 := v50[safeIndex(442, v50.Length)];
					var v52: array<multiset<bool>> := new multiset<bool>[9](i5 => multiset{!f4} * multiset{f4, f4});
					var v53: multiset<bool> := multiset{f4};
					var v54: map<bool, multiset<bool>> := map[f4 := v53];
					v52[safeIndex(258, v52.Length)] := if (!f4) then multiset{f4} else if (f4 in v54) then v54[f4] else v53;
					v52[safeIndex(258, v52.Length)] := v53 - (multiset{f4} - v53);
				}
				
				var v55 := "dohmu";
				var v56: map<string, int> := map[v55 := f6];
				var v57: seq<map<string, int>> := [v56];
				var v58: array<bool> := new bool[26];
				var v59 := m11(f4, !f4, (v57[safeIndex(f6, |v57|)])[v55 := cf11], v58, globalState);
				var v60: array<string> := new string[25](i6 => v55 + v55);
				v60[safeIndex(910, v60.Length)] := v55;
				var v61: map<int, int> := map[f15 := cf11];
				var v62: map<int, map<int, int>> := map[-0x57 := v61];
				v60[safeIndex(910, v60.Length)] := fm37(f4, f4, p0, f15 !in (if (p0 in v62) then v62[p0] else v61[fm9(f4, true, cf11, globalState) := f15]), globalState);
				var v63: array<multiset<D4>> := new multiset<D4>[21](i7 => multiset{DC12('f')} * multiset{DC12('l')});
				v63 := v63;
			case DC8(cf7) =>
				var v64: array<bool> := new bool[28];
				v64[safeIndex(707, v64.Length)] := f16 > 0x224;
				v64[safeIndex(707, v64.Length)] := f4;
				var v65 := "dv";
				v65 := "qssuchryd";
				var v66 := new C0(p0);
				var v67: array<int> := new int[14];
				var v68: set<array<int>> := {v67, v67};
				var v69 := DC29(v68);
				var v70: map<bool, set<array<int>>> := map[f4 := v69.cf48];
				var v71: seq<set<array<int>>> := [v68];
				v70 := v70[!f4 := if (f4) then v68 else v71[safeIndex(f6, |v71|)]];
			case DC11(cf12) =>
				match (DC24(!f4, f4, safeModuloInt(f16, -0x271))) {
					case DC24(cf36, cf37, cf38) =>
						f16 := cf38;
						cf36 := cf36;
						var v72: array<bool> := new bool[11];
						var v73 := DC5(v72);
						v73 := v73;
						var v74: array<int> := new int[2] [0x1bd, f15];
						var v75: seq<array<int>> := [v74];
						v75 := ([v74] + v75[safeIndex(f6, |v75|) := v74]) + (v75 + v75);
					case DC25(cf39, cf40) =>
						var v76 := 'b';
						var v77: map<string, int> := map["odwhwaejb" := |multiset{cf39}|];
						var v79 := "dqtrskur";
						var v80: set<string> := {v79};
						var v81: array<bool> := new bool[20](i9 => !f4);
						var v82 := m11(fm37(cf39, f4, f6, false, globalState) < (seq(abs(0x38a), i8  => ('c'))), fm0(v76, globalState), v77 + (map v78 : string | v78 in v80 :: (v78) := (|{f4, cf39, cf39}|)), v81, globalState);
						v79 := seq(abs(0x32c), i10  => (v76));
						var v83: seq<int> := [f16];
						var v84: map<multiset<int>, string> := map[multiset(v83) := v79];
						var v85, v86 := m0(v84, "lcxrsang" <= v79, v1 * v1, v76 !in v79[safeIndex(p0, |v79|) := v76], globalState);
						var v87: seq<bool> := [cf39];
						var v88: map<int, char> := map[v86 := v76];
						var v89: map<int, int> := map[v86 := cf40];
						var v90: map<int, int> := map[|v89| := cf40];
						var v91: seq<seq<bool>> := [fm40(f4, v85, v86, globalState), v87, v87];
						v85 := !((v1 - fm6(fm38(f15, p0, v85, |v87|, globalState), DC0(f15), v88, |v90|, globalState)) > (v1 * multiset(fm39(v79, DC27(v91), cf39, globalState))));
					case DC26(cf41) =>
						f16 := safeDivisionInt(f6, f15);
						var v92: array<bool> := new bool[18];
						v92[safeIndex(725, v92.Length)] := f16 <= p0;
						var v93: seq<int> := [106, f15];
						v92[safeIndex(725, v92.Length)] := v93 == v93;
						v92[safeIndex(725, v92.Length)] := v92[safeIndex(725, v92.Length)] <==> v92[safeIndex(725, v92.Length)];
						var v94 := DC0(f15);
						var v95 := "r";
						var v96: map<int, string> := map[p0 := v95];
						var v97: seq<string> := [v95, if (-0x2c2 in v96) then v96[-0x2c2] else v95];
						var v98 := 'y';
						v94, v97, v92[safeIndex(725, v92.Length)], v98 := v94.(cf0 := p0 + p0), v97 + v97, f4, v98;
					case DC23(cf35) =>
						var v99: array<map<string, array<int>>> := new map<string, array<int>>[12];
						var v100 := "yfnegvsi";
						var v101: array<int> := new int[19](i11 => i11 - f16);
						var v102: map<string, array<int>> := map[v100 := v101];
						v99[safeIndex(305, v99.Length)] := v102[seq(abs(252), i12  => ('h')) := v101];
						v99[safeIndex(305, v99.Length)] := v102;
						var v103: seq<int> := [f16, 357];
						v100 := fm37(f4, p0 != |v103|, -f6, f4, globalState);
						var v104: array<bool> := new bool[28];
						v104[safeIndex(205, v104.Length)] := f4;
						v104[safeIndex(205, v104.Length)] := f4 && f4;
						var v105: map<int, bool> := map[fm3(f15, f4, globalState) := v104[safeIndex(205, v104.Length)]];
						var v106: map<array<bool>, int> := map[v104 := f15];
						var v107 := 'f';
						var v108: map<char, bool> := map[v107 := f4];
						var v109 := DC22(v100, f4, if (v107 in v108) then v108[v107] else !f4, v104, p0);
						v105 := v105[if (v104 in v106) then v106[v104] else if (f16 in v1) then v1[f16] else f15 := v109.cf32];
				}
				
				var v110: array<int> := new int[9];
				v110[safeIndex(1, v110.Length)] := safeModuloInt(0x2e3, f15);
				v110[safeIndex(1, v110.Length)] := f6;
				f16 := f15 + -p0;
				var v111 := new C0(f15);
		}
		
		var v112 := "jmhajedk";
		var v113: set<string> := {v112};
		var i13 := 0;
		while (!((v113 * v113) > v113))
			decreases 100 - i13
		{
			if (i13 >= 100) {
				break;
			}
			
			i13 := i13 + 1;
			f16 := |(v112 + v112)|;
			var v114 := new C0(f16);
			var v115 := 'n';
			var v116: map<bool, char> := map[f4 := v115];
			var v117: seq<int> := [f15 + |v116|, v114.f12, f15];
			var v118: array<int> := new int[22];
			v118[safeIndex(23, v118.Length)] := f15;
			var v119: set<bool> := {f4, f4};
			var v120: map<int, set<bool>> := map[f16 := v119];
			var v121: map<int, bool> := map[f15 := f4];
			v117, f16, v118[safeIndex(23, v118.Length)] := v117 + [|v120|, 798], v114.f12, f6 * |v121|;
			var v122: array<bool> := new bool[10];
			v122[safeIndex(6, v122.Length)] := if (f4) then true else f4;
			v118[safeIndex(23, v118.Length)], v122[safeIndex(6, v122.Length)] := -(f6 * v118[safeIndex(23, v118.Length)]), f4;
		}
		var v123 := 'a';
		var v124: map<int, bool> := map[f15 := f4];
		var v125: multiset<bool> := multiset{f4};
		var v126 := DC28(|fm40(f4, if ((if (f4 in v125) then v125[f4] else f6) in v124) then v124[if (f4 in v125) then v125[f4] else f6] else f4, f16, globalState)|, -f6, f4, f4, f15);
		v123 := match v126 {
			case DC28(cf43, cf44, cf45, cf46, cf47) => v123
			case DC27(cf42) => v123
		};
		var v127: seq<int> := [f15];
		var v128: set<D6> := {DC16(f4, f4, v127), fm43(globalState)};
		var v129: array<bool> := new bool[20] [f4, !(if (f6 in v124) then v124[f6] else f4), !false, p0 < -f15, f4, f4, f4, !fm8(f15, f4, v112, v127, globalState), f4, p0 <= f15, false, f4, fm41(|v112|, v123, f4, globalState) <= fm41(0x1b8, v123, f4, globalState), fm0(v123, globalState), f4, false, f4 ==> f4, true, p0 in fm42(v123, v128, globalState), f4];
		forall i14 | 0 <= i14 < v129.Length {
			v129[i14] := f4;
		}
		var i15 := 0;
		while (f4)
			decreases 100 - i15
		{
			if (i15 >= 100) {
				break;
			}
			
			i15 := i15 + 1;
			f16 := -(f16 * p0);
			var v130: seq<bool> := [false, f4, f4];
			var v131 := DC13(f16, false, p0, DC13(0x16a, true, f16, v130).cf17);
			var v132: map<D4, bool> := map[v131 := f4];
			v132 := v132[v131 := f4];
			var v133 := new C0(0x2f1);
			v133 := v133;
		}
		var v134 := false;
		v129[safeIndex(395, v129.Length)] := true;
		var v135: array<int> := new int[3];
		var v136: set<array<int>> := {v135};
		var v137: array<D0> := new D0[20](i16 => DC1());
		var v138: map<D13, array<D0>> := map[DC29(v136) := v137];
		var v139: seq<map<D13, array<D0>>> := [v138];
		v135[safeIndex(693, v135.Length)] := |v139[safeIndex(f15, |v139|)]|;
		var v140 := DC25(v134, p0);
		f16, v134, v129[safeIndex(395, v129.Length)], v135[safeIndex(693, v135.Length)], v134 := v140.cf40, v134 <==> v134, v134, p0, f4;
		var v141: map<bool, int> := map[fm0(v123, globalState) := p0];
		r0 := v141;
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) {
		var v0: array<bool> := new bool[28];
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := f4 <== f4;
		}
		var i1 := 0;
		while (true)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v1: map<int, bool> := map[f6 := true];
			v1 := v1 + v1;
			var v2: multiset<bool> := multiset{p0};
			var v3: map<bool, int> := map[multiset{f4} !! v2 := f6];
			v3 := fm44(globalState) + (v3 + v3);
			var v4: seq<bool> := [!f4];
			f16 := |map[f15 := |v4|]|;
			var v5 := new C0(f6);
		}
		var v6: seq<bool> := [p0];
		var v7 := DC2(|v6|);
		match (v7) {
			case DC1() =>
				var v8 := false;
				var v9 := 'n';
				v8 := fm0(v9, globalState);
				var v10: set<bool> := {f4, v8, v8, false};
				if ((v10 * v10) == v10) {
					var v11: seq<array<bool>> := [v0];
					var v12: array<int> := new int[1];
					var v13: map<bool, bool> := map[p0 := !f4];
					var v14: multiset<char> := multiset{v9, v9, v9};
					var v15: seq<int> := [f15];
					var v16: array<int> := new int[26] [f6, f16, fm3(f6, p0, globalState), f16, f16, |map[v12 := fm7(globalState)]|, f6, f15, f6, |v13|, f6, f15, |v6|, -f15, f16, f15, f6, f16, f15, f15, |v14|, |v15|, f16, f16, -0x3a7, 0x282];
					var v17: multiset<array<int>> := multiset{v16, v16};
					var v18: map<bool, int> := map[!([v0] < v11) := (if (v12 in v17) then v17[v12] else v15[safeIndex(f16, |v15|)]) + 0x38b];
					var v19: map<int, int> := map[155 := f16];
					v18 := v18[f4 := if (f16 in v19) then v19[f16] else f16];
					var v20 := DC8(fm40(p0, v8, |fm37(v8, f4, f6, f4, globalState)|, globalState));
					var v21 := DC11(v20);
					v21 := v21;
					v0[safeIndex(299, v0.Length)] := p0;
					var v22 := "shvveca";
					f16, f16, v8, v0[safeIndex(299, v0.Length)], v8 := --|(v22 + v22)|, fm3(f15, v8, globalState), v8, v8, v8;
					var v23: array<multiset<bool>> := new multiset<bool>[5];
					var v24: map<char, bool> := map[v9 := f4];
					var v25: seq<array<multiset<bool>>> := [v23, v23];
					var v26: array<array<multiset<bool>>> := new array<multiset<bool>>[20] [v23, v23, if (if (v9 in v24) then v24[v9] else v0[safeIndex(299, v0.Length)]) then v23 else v23, v23, v23, v23, v23, v23, v23, v25[safeIndex(f16, |v25|)], v23, v23, v23, v23, v23, v23, v23, v23, v23, v23];
					v26[safeIndex(525, v26.Length)] := v23;
					v26[safeIndex(525, v26.Length)] := v23;
					v22 := (v22 + v22) + v22;
				} else {
					var v27: map<int, bool> := map[f6 := f4];
					var v28: map<bool, map<int, bool>> := map[p0 := v27];
					v28 := v28[p0 := map[0x367 := v8]];
					var v29: map<int, int> := map[f16 := f15];
					f16 := if (f15 in v29) then v29[f15] else f16;
					f16 := f6;
					var v30: multiset<int> := multiset{f15, -639};
					var v31: array<int> := new int[28];
					var v32: seq<array<int>> := [v31, v31, v31];
					f16, v30, v8, v32 := f16 + f6, v30, f4, v32;
					f16 := f16;
				}
				
				f16 := f16 - |(seq(abs(0x24f), i2  => (v9)))|;
				v0 := v0;
			case DC2(cf1) =>
				var v33 := false;
				var v34: map<bool, array<bool>> := map[f4 := v0];
				var v35: seq<array<bool>> := [v0, if (f4 in v34) then v34[f4] else v0];
				var v36 := "gx";
				var v37: map<array<bool>, int> := map[v35[safeIndex(|"htbf"|, |v35|)] := fm3(|v36|, p0, globalState)];
				var v38: map<int, int> := map[cf1 := f15];
				var v39: set<map<int, int>> := {v38, v38};
				var v40: map<bool, bool> := map[v33 := p0];
				var v41: map<bool, int> := map[p0 := cf1];
				var v42 := 'r';
				var v43: set<char> := {v42, v42};
				var v44: array<int> := new int[24] [f16, |v40|, f16, 0x179, cf1, |[f6, cf1, f6, |v34|, |v36|]|, f15, 0x1ef, f16, |multiset{f16, f16}|, if (f4 in v41) then v41[f4] else f16, f6, cf1, f6, f16, |v43|, f16, f16, cf1, cf1, |v36|, |v41|, |v36|, 180];
				var v45: seq<array<int>> := [v44, v44, v44, v44, v44];
				var v46: map<string, int> := map[v36 := |(seq(abs(0x3e), i3  => (v38)))|];
				var v47 := DC3(p0);
				var v48: map<D1, int> := map[v47 := cf1];
				var v49: array<int> := new int[26] [safeModuloInt(f16, |v36|), f6, -f15, safeModuloInt(|v6|, fm3(|v39|, v33, globalState)), safeModuloInt(-cf1, f16), 813, |v45|, f6, f16, 0xce, f15 * f6, -0x363, f16, safeDivisionInt(f15, --cf1), safeDivisionInt(cf1, 0x393), cf1, fm9(p0, p0, if ((seq(abs(0x376), i4  => (v42))) in v46) then v46[seq(abs(0x376), i4  => (v42))] else f16, globalState), safeDivisionInt(|v6|, f16), cf1, f15, |v48|, |"eadjey"|, cf1, safeModuloInt(|v6|, |v36|), f15, if (v33 in v41) then v41[v33] else -0xde];
				v44[safeIndex(987, v44.Length)] := cf1;
				var v50: seq<int> := [676, cf1];
				v33, v33, v37, v44[safeIndex(987, v44.Length)], v6 := !(f4 <==> p0), f4, v37, (if (v33 in v41) then v41[v33] else 94) + safeDivisionInt(cf1, |v50|), v6;
				v33, v36 := v33, v36;
				var v51: map<seq<array<bool>>, bool> := map[v35 := f4];
				var v52 := m11(if (v35 in v51) then v51[v35] else v33, v50 < (seq(abs(-0x34c), i5  => (f6))), v46 + v46, v0, globalState);
				if (p0) {
					v33 := v6[safeIndex(|v36|, |v6|)];
					v0[safeIndex(513, v0.Length)] := v33;
					v0[safeIndex(513, v0.Length)] := v33;
					v33 := f4;
					v40 := v40[p0 := fm0(v42, globalState)];
					v52 := 0x274 * -(-|v38| * v44[safeIndex(987, v44.Length)]);
				} else {
					v44[safeIndex(987, v44.Length)] := safeModuloInt(safeModuloInt(cf1, f15), v50[safeIndex(f16, |v50|)]);
					v0[safeIndex(413, v0.Length)] := v6[safeIndex(f6, |v6|)];
					v0[safeIndex(413, v0.Length)] := if (fm0(v42, globalState)) then !v33 else !true;
					var v53: multiset<bool> := multiset{v33, v33};
					var v54 := DC25(v0[safeIndex(413, v0.Length)], f16);
					v0[safeIndex(413, v0.Length)], v44[safeIndex(987, v44.Length)], r0 := (v53 > multiset(v6)) in ([f4, p0] + [true, f4]), f15 * v54.cf40, v42;
					v52, v52, v49, v0[safeIndex(413, v0.Length)] := -|(v38 + map[v44[safeIndex(987, v44.Length)] := cf1])|, v44[safeIndex(987, v44.Length)], v49, !!v33;
					v33 := v0[safeIndex(413, v0.Length)];
				}
				
			case DC0(cf0) =>
				var v55 := new C0(f16 + cf0);
				var v56 := true;
				var v57: array<int> := new int[4];
				var v58 := 'c';
				var v59: map<C0, bool> := map[v55 := p0];
				var v60: seq<char> := [v58];
				var v61: multiset<array<bool>> := multiset{v0, v0};
				var v62: seq<int> := [|v61|];
				var v63: seq<int> := [|map[v57 := v58]|, f6, fm3(f15, fm8(-cf0, !(if (v55 in v59) then v59[v55] else p0), v60, v62, globalState), globalState)];
				var v64: array<int> := new int[4] [-v55.f12, f15, v63[safeIndex(cf0, |v63|)], cf0];
				v57[safeIndex(78, v57.Length)] := v55.f12;
				var v65: multiset<bool> := multiset{p0, v56};
				v56, cf0, v57[safeIndex(78, v57.Length)] := if (v56) then v65 <= v65 else f4, if (p0) then f16 else f16, cf0;
				var v66: map<int, int> := map[cf0 := v57[safeIndex(78, v57.Length)]];
				var v67: multiset<int> := multiset{cf0, -806, |v6|, if (f16 in v66) then v66[f16] else v55.f12};
				var v68 := DC14(v67);
				var v69: map<D5, map<int, int>> := map[v68 := DC30(map[v57[safeIndex(78, v57.Length)] := f15]).cf49 + v66];
				v69 := v69;
				v57[safeIndex(78, v57.Length)] := v57[safeIndex(78, v57.Length)] + |v62|;
		}
		
		v0[safeIndex(243, v0.Length)] := f4;
		var v70 := "uehfebb";
		var v71 := 'v';
		v0[safeIndex(243, v0.Length)] := if (v70 < (seq(abs(0x110), i6  => ('i')))) then v71 !in v70 else f4;
		var v73: multiset<int> := multiset{f6};
		for i7 := -|(map v72 : int | v72 in v73 :: (v72 + -|v73|) := (if (f6 in v73) then v73[f6] else f6))| to f16 - f16 {
			var v75: array<int> := new int[6](i8 => i8 * |[{0x388}, set v74 : int | (0x2e3 <= v74) && (v74 < 0x3a2) :: (safeDivisionInt(v74, f16)), {f6, f15}]|);
			v75[safeIndex(117, v75.Length)] := fm9(p0, f4, i7, globalState);
			v75[safeIndex(117, v75.Length)] := f15;
			var v76 := new C0(v75[safeIndex(117, v75.Length)]);
			var v77: multiset<bool> := multiset{p0, v0[safeIndex(243, v0.Length)]};
			var v78: seq<multiset<bool>> := [v77, multiset{true, f4, f4, v0[safeIndex(243, v0.Length)], f4} * v77];
			v78 := (v78 + fm45(p0, f15, v0[safeIndex(243, v0.Length)], v76.f12, globalState)) + [v77];
			var v79 := DC3(fm8(779, v0[safeIndex(243, v0.Length)], v70, seq(abs(520), i9  => (v75[safeIndex(117, v75.Length)])), globalState));
			var v80: array<array<int>> := new array<int>[4];
			var v81: map<bool, array<array<int>>> := map[(v79.(cf2 := p0)).cf2 := v80];
			v81 := v81[!p0 := v80];
		}
		v0[safeIndex(243, v0.Length)] := f4;
		r0 := v71;
		var v82: array<D1> := new D1[3](i10 => DC3(f4));
		r1 := v82;
	}
	method m11(p0: bool, p1: bool, p2: map<string, int>, p3: array<bool>, globalState: GlobalState) returns (r0: int) {
		var v0 := "h";
		v0 := v0;
		var v1 := false;
		f5[safeIndex(595, f5.Length)] := p3;
		var v2: seq<bool> := [v1];
		var v3 := DC4(f15, f16);
		v1, f5[safeIndex(595, f5.Length)], v1, v1, v1 := v2[safeIndex(f15 - f6, |v2|)], p3, p0, match v3.(cf3 := -0x3b4) {
			case DC4(cf3, cf4) => p1
			case DC3(cf2) => p0
		}, false;
		var v4 := 'd';
		var v5: map<char, bool> := map[v4 := p0];
		var i0 := 0;
		while (if (v0 != v0) then |v5| > f6 else p1)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v6: array<int> := new int[14];
			v6 := v6;
			var v7: map<int, int> := map[f6 := -f6];
			var v8 := new C0(if (f4) then |v7| else f6);
			v1 := fm40(f4, p1, v8.f12, globalState) == [v1];
			var v9: set<bool> := {v1 || f4, p0};
			v9 := v9;
		}
		if (false) {
			var v10: set<array<bool>> := {p3};
			var v11: array<array<bool>> := new array<bool>[23];
			var v12: array<array<array<bool>>> := new array<array<bool>>[12] [v11, v11, v11, v11, v11, v11, v11, v11, v11, v11, v11, v11];
			v12[safeIndex(401, v12.Length)] := v11;
			var v13: map<int, int> := map[safeDivisionInt(|[f16]|, fm9(f4, p1, f15, globalState)) := f15];
			var v14: seq<array<array<bool>>> := [v11, v11, v11, v11];
			var v15: multiset<bool> := multiset{p0};
			v10, v12[safeIndex(401, v12.Length)], v13 := v10, v14[safeIndex(if (!false in v15) then v15[!false] else f15, |v14|)], v13 + v13;
			f16 := fm7(globalState);
			if (f4) {
				var v16: array<int> := new int[6];
				var v17: map<int, array<int>> := map[|v0| := v16];
				v17 := v17;
				var v18: multiset<int> := multiset{f15};
				v18 := (v18 * multiset{f15})[safeDivisionInt(-|v0|, f15) := abs(f15)];
				f5[safeIndex(595, f5.Length)][safeIndex(253, f5[safeIndex(595, f5.Length)].Length)] := f4;
				var v19: map<bool, string> := map[p0 := v0];
				var v20: seq<int> := [-f16];
				var v21: map<bool, int> := map[true := -|v20|];
				var v22 := DC28(f6, f16, v1, !v1, if (f4 in v21) then v21[f4] else 0x92);
				f16, r0, f5[safeIndex(595, f5.Length)][safeIndex(253, f5[safeIndex(595, f5.Length)].Length)] := |v19|, v22.cf44, p0;
				var v23: seq<string> := [v0, v0];
				v1 := !((v23 + v23) <= v23);
				f5[safeIndex(595, f5.Length)][safeIndex(253, f5[safeIndex(595, f5.Length)].Length)] := !f4;
			} else {
				r0 := f15 * f16;
				var v24 := DC2(-f6);
				v24 := v24;
				f16 := f16;
				var v25: map<seq<bool>, bool> := map[v2 := p0];
				var v26: set<int> := {f15};
				v25 := v25[v2 := v26 >= v26];
				r0 := |v0|;
			}
			
			var v27 := DC0(f6);
			var v28: C0 := new C0(|v2|);
			var v29: map<C0, bool> := map[v28 := v1];
			var v30: map<int, map<C0, bool>> := map[f15 := v29];
			var v31: seq<int> := [-0x9f, -f16];
			v27, v29, r0, v1, v1 := v27, (map[v28 := f4])[v28 := !f4] + (if (f15 in v30) then v30[f15] else v29), f6, fm8(f15 * v28.f12, v31[safeIndex(f6, |v31|)] <= v28.f12, v0, v31, globalState), p1 && p0;
			var v32: map<multiset<int>, int> := map[multiset{v28.f12} := if (!v1 in v15) then v15[!v1] else fm7(globalState)];
			var v33: map<int, C0> := map[f16 := v28];
			var v34: map<bool, map<multiset<int>, int>> := map[v33 != v33 := v32];
			v32 := if (p0 in v34) then v34[p0] else v32;
		} else {
			f16 := f16;
			var v35 := new C0(520);
			var v36: multiset<bool> := multiset{p1};
			var v37 := DC10(|v36|);
			var v38: map<int, bool> := map[f16 := v37.cf11 == v35.f12];
			var v39: seq<int> := [f16];
			var v40: seq<seq<int>> := [[v35.f12], v39, v39, seq(abs(-0x3d6), i1  => (-v35.f12)), v39];
			v38 := v38[f6 := v40 != v40];
			var v41: array<int> := new int[17](i3 => i3 - f6);
			var v42: map<string, map<bool, array<int>>> := map["wirva" + (seq(abs(-0x3e1), i2  => ('u'))) := map[p1 := v41]];
			var v43: map<bool, array<int>> := map[p0 := v41];
			v42 := v42[v0 := v43 + v43];
			v1 := f4;
		}
		
		var v44: map<bool, int> := map[p0 := f16];
		r0 := f16 + safeDivisionInt(|v44|, -0x2b2);
		var v45: map<char, int> := map[v4 := f15];
		if (v45 == (v45 + v45[v4 := 220])) {
			f16 := safeModuloInt(f16, f15);
			var v46: seq<int> := [|v0| * 668, f16 * f15];
			var v47 := DC27([v2]);
			v46 := fm39(v0, v47, p1, globalState);
			var v48: multiset<bool> := multiset{!false, p0};
			var v49: multiset<int> := multiset{f16, f16, |v48|};
			var v50 := DC19("imf");
			var v51: map<multiset<int>, string> := map[v49 := v50.cf25];
			var v52: set<int> := {f16};
			var v53, v54 := m0(v51, {f6} >= v52, v49, v1, globalState);
			var v55: map<int, int> := map[f6 := v46[safeIndex(|v0|, |v46|)]];
			var v56: array<int> := new int[6] [|v0| + -f6, 0x18e, 0x3e, safeModuloInt(f16, -|v48|), if (f6 in v55) then v55[f6] else f16, v54 * f16];
			v56, v53, f16, f16 := v56, p1, f16 + (|(seq(abs(0x137), i4  => (|"csd"|)))| + |fm46(v4, globalState)|), if (v1 in v48) then v48[v1] else f6;
			if (f4) {
				var v57, v58 := m0(v51[v49 := v0], p0, v49, p1, globalState);
				var v60: set<multiset<int>> := {v49};
				var v61, v62 := m0(map v59 : multiset<int> | v59 in v60 :: (v59) := ("lgs"), |v44| >= 0x244, v49 + v49, v1, globalState);
				v1 := p1;
				v61 := p1;
				r0 := v54;
			} else {
				var v63: map<int, bool> := map[f15 := p1];
				v44 := v44[f4 := f15 * fm9(false, p1, |v63|, globalState)];
				var v64: map<bool, bool> := map[v53 := p0];
				var v65: seq<multiset<int>> := [v49];
				var v66: map<map<bool, bool>, int> := map[v64 + v64 := |v65[safeIndex(-f16, |v65|)]|];
				v66 := v66;
				var v67 := new C0(f16 + 0x34b);
				v48 := (v48 + v48) + (v48 + fm47(f6, globalState));
				v1 := v53;
			}
			
		} else {
			v1 := p0 || p1;
			v4 := v4;
			var v68: array<seq<int>> := new seq<int>[15](i5 => if (f4) then [|multiset{f6, -0x252, f16}|, -f15] else [f16]);
			var v69: map<bool, bool> := map[p1 := v1];
			var v70: seq<int> := [|v69|];
			var v71: seq<int> := [f15, |v70|, f6];
			v68[safeIndex(489, v68.Length)] := v71;
			var v72: map<bool, char> := map[fm0(v4, globalState) := v4];
			v68[safeIndex(489, v68.Length)] := ([f15] + v71)[safeIndex(876, |([f15] + v71)|) := |v72|] + [f16, f15];
			var v73: map<string, int> := map[v0 + v0 := f6];
			v73 := v73[v0 := f16 + 424];
			v1 := p0;
		}
		
		r0 := f16;
	}
}

class C2 extends T0 {
	const f17 : string
	constructor (f17 : string, f4 : bool) {
		this.f17 := f17;
		this.f4 := f4;
	}
	
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> {
		DC14(multiset{0x1e6}).cf18
	}
	function fm7(globalState: GlobalState): int {
		0x33f
	}
	function fm48(p0: bool, p1: int, p2: int, globalState: GlobalState): bool {
		DC3(f4).cf2
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) {
		for i0 := fm7(globalState) to p0 {
			var v0: array<bool> := new bool[18];
			v0 := new bool[2](i1 => f4);
			var v1 := 0x325;
			v1 := -78;
			var v2: set<bool> := {f4};
			var v3, v4 := m12(v1, v2 - v2, map[v1 := |v2|], f4, globalState);
			if (f4) {
				var v5 := true;
				v5 := v2 < {f4, f4, f4};
				var v6: set<int> := {i0, v3};
				var v7: multiset<bool> := multiset{false, v5, v5, v5};
				var v8 := 'p';
				var v9: set<set<int>> := {v6, v6 * {i0, |v7|, |f17[safeIndex(v4, |f17|) := v8]|, |v7[v5 := abs(-p0)]|}};
				var v10 := DC26(v5);
				var v11: map<D11, string> := map[v10 := "hjgsgbnw" + f17];
				var v12: array<D2> := new D2[16];
				v12[safeIndex(838, v12.Length)] := DC6();
				var v14: multiset<D11> := multiset{v10};
				var v15 := DC6();
				v9, v11, v12[safeIndex(838, v12.Length)], v5, v5 := if (false) then v9 else v9 - v9, (v11 + v11) + ((map v13 : D11 | v13 in v14 :: (v13) := (f17)) + v11), v15, !f4 && (v1 >= v1), f4;
				var v16 := new C0(-0x3d4);
				v5 := !(if (v5) then f4 else v5);
				var v17 := DC5(v0);
				var v18: map<bool, array<bool>> := map[v5 || f4 := v17.cf5];
				v18 := v18[v5 <== v5 := v0];
			} else {
				var v19: multiset<int> := multiset{i0};
				v19 := v19 * v19;
				var v20: array<int> := new int[21](i2 => i2 - v4);
				v20[safeIndex(217, v20.Length)] := i0;
				var v21: seq<int> := [v3];
				var v22: seq<seq<int>> := [v21];
				var v23 := false;
				var v24: seq<bool> := [f4, v23];
				var v25: seq<string> := [f17, f17];
				v20[safeIndex(217, v20.Length)], v22, v23, v4, v23 := fm7(globalState), (fm49(f17, v24, v3, globalState))[safeIndex(|v25[safeIndex(v3, |v25|)]|, |fm49(f17, v24, v3, globalState)|) := seq(abs(345), i3  => (i0))] + [v21], true, -(v3 + -v1), f4;
				var v26: multiset<bool> := multiset{v23, v23, false};
				v20[safeIndex(217, v20.Length)] := -(if (v23 in v26) then v26[v23] else v4);
				v2 := {!f4};
				var v27: array<char> := new char[1];
				v27[safeIndex(331, v27.Length)] := 'l';
				var v28 := 'c';
				v27[safeIndex(331, v27.Length)] := v28;
			}
			
		}
		var v29: array<bool> := new bool[25];
		var v30 := DC22("jdyfjcy", -p0 == 0x7a, !f4, v29, -p0);
		match (v30) {
			case DC22(cf30, cf31, cf32, cf33, cf34) =>
				cf34 := if (cf31) then safeDivisionInt(cf34, p0) else 0x239;
				var v31: map<bool, int> := map[if (f4) then cf32 else f4 := cf34];
				var v32: seq<int> := [p0];
				var v33: set<bool> := {cf32, true};
				var v34: array<int> := new int[9] [p0, cf34, |v32|, 0x9c, p0, p0, |cf30|, fm7(globalState), |v33|];
				var v35: map<array<int>, int> := map[v34 := p0];
				v31 := v31[map[v34 := fm7(globalState)] == v35[v34 := -0x168] := p0];
				var v36 := DC7(cf34 + p0);
				cf34, cf32, v36 := p0, cf32, v36;
				var v37 := new C0(p0);
			case DC21(cf29) =>
				var v38 := DC6();
				match (v38) {
					case DC6() =>
						var v39 := 'w';
						var v40 := DC20(f4, v39, |f17|);
						var v41: seq<D9> := [v40];
						var v42 := DC34(v41);
						v41 := (v42.(cf57 := v41)).cf57;
						v29[safeIndex(751, v29.Length)] := if (f4) then f4 else f4;
						var v43: array<int> := new int[16];
						v43[safeIndex(437, v43.Length)] := safeDivisionInt(p0, p0);
						var v44: array<array<int>> := new array<int>[27];
						var v45 := true;
						v29[safeIndex(751, v29.Length)], v43[safeIndex(437, v43.Length)], v44, v45 := f4, p0, v44, !v45;
						var v46: array<bool> := new bool[14];
						v46 := v46;
						var v47: seq<bool> := [v29[safeIndex(751, v29.Length)]];
						var v48: seq<int> := [|v47|, p0];
						var v49: multiset<bool> := multiset{false, v29[safeIndex(751, v29.Length)]};
						var v50: map<int, int> := map[v48[safeIndex(|f17|, |v48|)] := if (v29[safeIndex(751, v29.Length)]) then -p0 else |v49|];
						var v51: set<int> := {v43[safeIndex(437, v43.Length)]};
						var v52: multiset<int> := multiset{|v51|, p0};
						v50 := v50[v43[safeIndex(437, v43.Length)] := if (v45) then 0x1cc else |v52|];
					case DC7(cf6) =>
						var v53: seq<set<bool>> := [{f4, f4, !true}];
						v53 := v53;
						var v54: seq<bool> := [f4];
						var v55: seq<bool> := [if (f4) then v54[safeIndex(-cf6, |v54|)] else false];
						v55 := v54;
						var v56: multiset<int> := multiset{p0, p0};
						var v57: map<string, int> := map[if (fm48(f4, cf6, p0, globalState)) then f17 else fm50(0x37c, p0, |v56|, f4, globalState) := p0];
						var v58: seq<int> := [p0, p0];
						var v59: map<int, string> := map[|v58| := f17];
						var v60: seq<int> := [|v59|];
						var v61: seq<int> := [325, cf6, p0, v60[safeIndex(cf6, |v60|)]];
						v57 := map[seq(abs(228), i4  => ('d')) := v61[safeIndex(p0, |v61|)] * |(seq(abs(761), i5  => ('x')))|];
						var v62 := false;
						var v63: array<int> := new int[22](i6 => safeDivisionInt(i6, cf6));
						var v64: map<int, int> := map[p0 := p0];
						v63[safeIndex(0, v63.Length)] := |[cf6]| * (if (p0 in v64) then v64[p0] else cf6);
						v62, v63[safeIndex(0, v63.Length)], v62 := v62, if (cf6 in v64) then v64[cf6] else p0, "hasb" < f17;
					case DC5(cf5) =>
						var v65: array<set<map<bool, bool>>> := new set<map<bool, bool>>[10];
						var v66: map<bool, bool> := map[f4 := f4];
						var v67: set<map<bool, bool>> := {v66};
						v65[safeIndex(210, v65.Length)] := if (f4) then v67 else v67;
						var v68: multiset<bool> := multiset{f4};
						v65[safeIndex(210, v65.Length)] := set v69 : map<bool, bool> | v69 in (map[v66 := multiset{false}])[v66 := v68] :: (v69);
						var v70: map<int, int> := map[p0 := p0];
						var v71: set<bool> := {f4, f4};
						var v72: array<int> := new int[16] [|v70[p0 := p0]|, p0, 0xf5, p0, |v71|, p0, p0, p0, p0, p0, |cf29|, p0, p0, p0, |f17|, |(v70 + v70)|];
						v72[safeIndex(918, v72.Length)] := p0;
						v72[safeIndex(918, v72.Length)] := 0x34c;
						var v73 := new C0(v72[safeIndex(918, v72.Length)] * v72[safeIndex(918, v72.Length)]);
						var v74: map<int, set<bool>> := map[74 := v71];
						var v75: seq<int> := [v73.f12, |v74|];
						var v76: seq<int> := [-|v75| - p0, v73.f12];
						v76 := v76;
				}
				
				var v77: seq<bool> := [f4];
				var v78: seq<seq<bool>> := [v77, v77, [f4, f4]];
				v77 := DC13(p0, true, -|f17|, v78[safeIndex(p0, |v78|)]).cf17;
				var v79: array<int> := new int[3];
				var v80: map<array<int>, bool> := map[v79 := f4];
				v80 := v80[v79 := p0 >= p0];
				if ((p0 <= p0) || (p0 >= p0)) {
					var v81: array<array<bool>> := new array<bool>[4] [v29, v29, v29, v29];
					var v82 := new C1(p0, p0, v81, |"bw"| * 0xce, f4);
					var v83: seq<int> := [|"mxhimf"|];
					var v84: map<seq<int>, bool> := map[v83 := f4 ==> f4];
					var v86: seq<seq<int>> := [v83, v83];
					v84 := ((map v85 : seq<int> | v85 in v86 :: (v85) := (f4)) + v84) + (v84 + v84);
					var v87 := new C0(p0);
					var v88 := false;
					var v89: array<string> := new string[4] [f17 + f17, f17, f17 + f17, f17];
					v89[safeIndex(575, v89.Length)] := f17;
					var v90 := 'e';
					var v91: map<char, array<bool>> := map[v90 := v29];
					var v92: multiset<array<bool>> := multiset{v29, if ('o' in v91) then v91['o'] else v29, v29, v29, if (v88) then v29 else v29};
					var v93: map<int, int> := map[v82.f15 := |f17|];
					v82.f16, v88, v88, v88, v89[safeIndex(575, v89.Length)] := if (v29 in v92) then v92[v29] else (if (-v82.f16 in v93) then v93[-v82.f16] else v83[safeIndex(p0, |v83|)]) - v82.f15, v87.f12 != v87.f12, v88, v88, seq(abs(-0x324), i7  => (v90));
					v90 := 'x';
				} else {
					var v94 := 297;
					v94 := p0;
					var v95 := false;
					var v96 := DC4(p0, p0);
					var v97: array<array<bool>> := new array<bool>[10];
					var v98: C1 := new C1(p0, p0, v97, v94, f4);
					var v99: seq<C1> := [v98];
					var v100: set<string> := {f17, f17};
					var v101: map<set<string>, seq<C1>> := map[v100 := v99];
					var v102: map<bool, bool> := map[f4 := v95];
					var v103: set<bool> := {f4, if (v95 in v102) then v102[v95] else true};
					var v104 := 'e';
					var v105: map<char, set<bool>> := map[v104 := v103];
					v94, v95, v96, v95, v94 := p0, v99 != (if (v100 in v101) then v101[v100] else v99), v96, {v95, f4} >= (v103 + (if (v104 in v105) then v105[v104] else {!true})), p0;
					v95 := false;
					var v106: seq<array<int>> := [v79];
					var v107 := DC5(v29);
					v95 := -safeDivisionInt(|v106|, v94) <= safeDivisionInt(-p0, -|[v107]|);
					var v108: map<int, int> := map[-0xe := v98.f15];
					var v109, v110 := m12(p0, v103, v108 + v108, f4, globalState);
				}
				
		}
		
		var v111 := new C0(p0);
		var i8 := 0;
		while (f4 && f4)
			decreases 100 - i8
		{
			if (i8 >= 100) {
				break;
			}
			
			i8 := i8 + 1;
			v29[safeIndex(978, v29.Length)] := f4;
			v29[safeIndex(978, v29.Length)] := f4;
			var v112: map<int, bool> := map[p0 := !f4];
			var v113: map<bool, bool> := map[f4 := v29[safeIndex(978, v29.Length)]];
			var v114: set<bool> := {f4};
			v112 := fm51(map[if (fm0('c', globalState) in v113) then v113[fm0('c', globalState)] else false := v114], true, globalState);
			var v115: seq<bool> := [f17 <= f17, if (f4 in v113) then v113[f4] else v29[safeIndex(978, v29.Length)], v29[safeIndex(978, v29.Length)], f4];
			var v116 := DC26(v29[safeIndex(978, v29.Length)]);
			var v117: multiset<bool> := multiset{f4};
			var v118 := DC33(v116.cf41, v111.f12, v117, p0, v114);
			var v119 := DC13(|([v118])[safeIndex(p0, |[v118]|) := v118]|, v29[safeIndex(978, v29.Length)], p0, v115);
			v115 := (v119.(cf17 := v115, cf14 := p0)).cf17 + v115;
			v29[safeIndex(978, v29.Length)], v29[safeIndex(978, v29.Length)] := v114 >= (if (f4) then v114 else v114), !f4;
		}
		var v120 := 906;
		v120 := safeModuloInt(v120, v111.f12);
		v120 := |f17| + v120;
		var v121: map<bool, int> := map[if (f4) then v111.fm17(f4, globalState) else f4 := v111.f12 - 518];
		r0 := v121;
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) {
		var v0: seq<bool> := [f4, f4];
		var v1 := DC8(v0);
		var v2: array<D3> := new D3[26] [v1, DC8(v0), v1, v1, DC8(v0), DC8(v0), v1, v1, v1, v1, DC8(v0), v1, fm52(-402, globalState), v1, v1, v1, v1, v1, v1, v1, v1, v1, DC8(v0), DC8(v0), v1, v1];
		v2[safeIndex(878, v2.Length)] := v1;
		v2[safeIndex(878, v2.Length)] := v1.(cf7 := v0);
		var v3 := 0xbb;
		var v4: array<int> := new int[9] [safeModuloInt(v3, |([v3, v3])[safeIndex(0x198, |[v3, v3]|) := v3]|), v3, v3, v3, v3, v3 + 0x241, fm3(v3, f4, globalState), v3, -|f17|];
		forall i0 | 0 <= i0 < v4.Length {
			v4[i0] := safeModuloInt(i0, -(v3 - v3));
		}
		var v5: array<bool> := new bool[1] [f4];
		var v6: map<bool, array<bool>> := map[p0 := v5];
		var v7: map<int, int> := map[-(|v6| * v3) := -v3];
		v7 := v7[v3 := v3];
		var v8: multiset<bool> := multiset{!f4, p0};
		var v11 := 'n';
		var v12: array<array<bool>> := new array<bool>[13] [v5, v5, v5, v5, v5, v5, v5, v5, v5, v5, v5, v5, v5];
		var v14 := new C1(|fm50(if (f4 in v8) then v8[f4] else v3, |(map v9 : int | (0x141 <= v9) && (v9 < 0x179) :: (safeDivisionInt(v9, |(map v10 : int | (72 <= v10) && (v10 < 693) :: (v10 - |"cii"|) := (f4))|)) := (|f17|))|, v3, fm0(v11, globalState), globalState)|, v3, v12, fm3(|(map v13 : int | (0x19c <= v13) && (v13 < -0x2a4) :: (v13 * v3) := (p0))|, p0, globalState), p0);
		var v15: set<bool> := {p0, p0, p0};
		v15 := fm53(safeModuloInt(v3, v3), fm54(v14.f16, globalState) + (seq(abs(644), i1  => (map[false := p0]))), globalState);
		var i2 := 0;
		while (!((if (p0) then v3 else v14.f15) < (v14.f16 * v14.f15)))
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			v5[safeIndex(549, v5.Length)] := v14.f15 < 21;
			v5[safeIndex(549, v5.Length)] := fm0(v11, globalState);
			v14.f16 := fm3(fm7(globalState), p0, globalState);
			v5[safeIndex(549, v5.Length)] := f4;
			v5[safeIndex(549, v5.Length)] := |f17| <= v3;
		}
		r0 := v11;
		var v16 := DC3(p0);
		r1 := new D1[12] [v16, v16, v16, v16, v16, v16, v16, v16, v16, v16.(cf2 := f4), fm55(!p0, f4, 0x223, globalState), DC3(f4)];
	}
	method m12(p0: int, p1: set<bool>, p2: map<int, int>, p3: bool, globalState: GlobalState) returns (r0: int, r1: int) {
		var v0: array<bool> := new bool[16];
		for i0 := |(p1 * {f4, f4})| to |[v0, v0, v0, v0]| {
			var v1 := new C0(0x10f);
			var v2: map<bool, int> := map[f4 := p0];
			var v3: seq<int> := [i0];
			var v4 := DC19(seq(abs(-0x285), i1  => ('i')));
			var v5: array<int> := new int[21] [p0, i0, i0, 0x380, v1.f12, |multiset(v3)|, v1.f12, i0, p0, i0, i0, p0, |f17|, |f17|, v1.f12, |v4.cf25|, -0x288, i0, p0, i0, p0];
			var v6: seq<array<int>> := [v5];
			var v7: map<map<bool, int>, seq<array<int>>> := map[v2 := v6];
			var v8: seq<seq<array<int>>> := [v6];
			v7 := v7[v2 := v8[safeIndex(p0, |v8|)]];
			var v9: array<array<bool>> := new array<bool>[7];
			v9[safeIndex(170, v9.Length)] := v0;
			v0[safeIndex(985, v0.Length)] := false;
			var v10 := 'c';
			var v11: T1 := new C1(i0, p0, v9, fm7(globalState), p3);
			var v12: map<T1, bool> := map[v11 := v11.f4];
			v0[safeIndex(964, v0.Length)] := if (v11 in v12) then v12[v11] else v11.f4;
			var v13: seq<bool> := [v11.f4, v11.f4];
			v9[safeIndex(170, v9.Length)], r0, v0[safeIndex(985, v0.Length)], v10, v0[safeIndex(964, v0.Length)] := v0, fm7(globalState), !(if (v11.f4) then !false else v11.f4), 'k', if (v13[safeIndex(p0, |v13|)]) then f4 <== p3 else p3;
			var v14: map<int, bool> := map[p0 := v11.f4];
			var v15: seq<map<int, bool>> := [v14, map[|v14| := p3] + map[v1.f12 := v11.f4], v14];
			v15 := v15;
		}
		var v16 := 'l';
		var i2 := 0;
		while (fm0(v16, globalState))
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v17 := DC21(map[f4 := p0]);
			var v18: seq<D10> := [v17, v17];
			var v19: multiset<int> := multiset{390 - p0, p0, -|v18| - |p2|};
			r0 := if (-(p0 * -p0) in v19) then v19[-(p0 * -p0)] else p0;
			var v20 := "ulsxdir";
			v20 := v20 + (if (!f4) then v20 else f17);
			var v21: multiset<bool> := multiset{p3};
			var v22: array<array<bool>> := new array<bool>[8];
			var v23 := new C1(p0, -(|[v16, v16, v16]| + (if (true in v21) then v21[true] else p0)), if (p3) then v22 else v22, p0, p0 != 0x1ed);
			r0 := v23.fm9(f4, f4, v23.f15, globalState) * v23.f15;
		}
		r0 := -((p0 - p0) - safeModuloInt(p0, p0));
		var v24: array<int> := new int[14];
		var v25: map<int, array<int>> := map[-fm7(globalState) + p0 := v24];
		v24 := if (0x2ce in v25) then v25[0x2ce] else v24;
		var v26 := DC7(if (f4) then p0 else p0);
		match (v26) {
			case DC6() =>
				var v27 := new C0(safeDivisionInt(--298, p0));
				var v28 := true;
				v28 := v28;
				r0 := p0;
				r1 := v27.f12;
			case DC7(cf6) =>
				r1 := p0;
				var v29 := false;
				v29 := p0 != |f17|;
				var v30: multiset<int> := multiset{-p0};
				var v31: map<multiset<int>, string> := map[v30 := seq(abs(-304), i3  => (v16))];
				var v32: seq<bool> := [p3];
				var v33: seq<seq<bool>> := [fm56(f4, cf6, p3, globalState), v32, v32];
				var v34: multiset<seq<bool>> := multiset{v32};
				var v35, v36 := m0(v31, multiset(v33) <= v34, v30, |p2| == p0, globalState);
				var v37: array<array<bool>> := new array<bool>[17];
				var v38: C1 := new C1(p0, 0x8f, v37, v36, !false);
				var v39: multiset<bool> := multiset{f4, p3};
				var v40 := DC16(false, p3, ([cf6])[safeIndex(|map[v35 := v38]|, |[cf6]|) := |v39|]);
				var v41: map<int, D6> := map[fm3(cf6, f4, globalState) := v40];
				var v42: set<array<bool>> := {v0};
				v41 := v41[|(v42 + v42)| := v40];
			case DC5(cf5) =>
				var v43: multiset<bool> := multiset{true};
				var v44: seq<int> := [p0, if (p3 in v43) then v43[p3] else -p0, p0];
				cf5, r0, r1 := v0, p0, v44[safeIndex(safeModuloInt(p0, p0), |v44|)];
				var v45 := true;
				v45 := !f4;
				var v46: map<array<int>, bool> := map[v24 := !f4];
				var v47: map<bool, map<array<int>, bool>> := map[v45 := v46];
				v47 := v47[p3 := if (v45) then v46 else v46];
				var v48: map<string, bool> := map[(f17 + "x")[safeIndex(298, |(f17 + "x")|) := v16] := multiset{p3} >= v43];
				v48 := v48[if (false) then "jpo" else ("qlx")[safeIndex(p0, |"qlx"|) := v16] := f4];
		}
		
		if (f4 && p3) {
			v0[safeIndex(936, v0.Length)] := p3;
			v0[safeIndex(936, v0.Length)] := f4;
			var v49: seq<int> := [p0, p0];
			var v50: seq<int> := [if (v49[safeIndex(-|[v0[safeIndex(936, v0.Length)], p3]|, |v49|)] in p2) then p2[v49[safeIndex(-|[v0[safeIndex(936, v0.Length)], p3]|, |v49|)]] else p0, p0];
			v50 := v49[safeIndex(fm7(globalState), |v49|) := p0];
			var v51: array<string> := new string[19];
			v51[safeIndex(893, v51.Length)] := f17;
			var v52: seq<bool> := [!v0[safeIndex(936, v0.Length)]];
			v51[safeIndex(893, v51.Length)] := (if (v52[safeIndex(p0, |v52|)]) then f17 else "dehlggigm") + f17;
			var v53: array<char> := new char[15] [v16, v16, v16, v16, 'e', 's', v16, v16, v51[safeIndex(893, v51.Length)][safeIndex(p0, |v51[safeIndex(893, v51.Length)]|)], 'g', fm1(p0, p3, globalState), v16, v16, fm1(p0, v0[safeIndex(936, v0.Length)], globalState), v16];
			v53 := v53;
			var v54: array<bool> := new bool[2](i4 => p3);
			var v55: map<int, array<bool>> := map[safeDivisionInt(p0, 0x87) := v54];
			v55 := map[p0 := v54];
		} else {
			var v56 := false;
			var v57: seq<bool> := [f4];
			var v58: map<int, bool> := map[p0 := false];
			var v59: seq<int> := [|v58|];
			var v60: map<bool, int> := map[v56 := |v59|];
			var v61: multiset<bool> := multiset{v57[safeIndex(if (p3 in v60) then v60[p3] else p0, |v57|)], p3};
			var v62: array<multiset<bool>> := new multiset<bool>[2] [v61, v61];
			var v63: array<array<multiset<bool>>> := new array<multiset<bool>>[7] [v62, v62, v62, v62, v62, if (f4) then v62 else v62, v62];
			v63[safeIndex(906, v63.Length)] := v62;
			r1, v0, v56, v63[safeIndex(906, v63.Length)] := -p0, v0, !f4, v62;
			var v64: array<array<bool>> := new array<bool>[20];
			var v65 := new C1(v59[safeIndex(p0, |v59|)], safeModuloInt(p0, p0), v64, safeDivisionInt(|multiset{|v57|, 0x9}|, p0), p3);
			v56 := v56;
			v56 := v59[safeIndex(v65.f15, |v59|) := p0] != (v59 + fm57(v65.f15, -v65.f16, globalState))[safeIndex(v65.f16, |(v59 + fm57(v65.f15, -v65.f16, globalState))|) := v65.f16];
			v56 := v65.f16 <= v65.f16;
		}
		
		r0 := -0x103;
		r1 := -p0;
	}
}

class C3 extends T1 {
	constructor (f5 : array<array<bool>>, f6 : int, f4 : bool) {
		this.f5 := f5;
		this.f6 := f6;
		this.f4 := f4;
	}
	
	function fm8(p0: int, p1: bool, p2: seq<char>, p3: seq<int>, globalState: GlobalState): bool {
		f4 <== (f6 != -f6)
	}
	function fm9(p0: bool, p1: bool, p2: int, globalState: GlobalState): int {
		0x1dc
	}
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> {
		multiset{f6, f6} - multiset{f6, f6}
	}
	function fm7(globalState: GlobalState): int {
		f6
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) {
		var v0: array<int> := new int[22](i1 => safeDivisionInt(i1, p0));
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := i0 - p0;
		}
		var v1: seq<bool> := [false];
		var v2: multiset<seq<bool>> := multiset{v1};
		var v3 := DC15(v2);
		v3 := fm33(f4, -|"dqdyv"|, if (!f4) then false else !f4, globalState);
		var v4: seq<int> := [p0];
		var v5: set<int> := {p0, |v4|};
		for i2 := p0 to p0 - |v5| {
			var v6 := "sj";
			var v7: map<multiset<int>, string> := map[multiset(seq(abs(871), i3  => (p0))) := v6];
			var v8: multiset<int> := multiset{|v6|};
			var v9, v10 := m0(fm34(i2, p0, !f4, f4, globalState) + v7, f4, v8, f4, globalState);
			v10 := f6 * (if (!f4) then p0 else if (i2 in v8) then v8[i2] else p0);
			var v11 := new C0(0x377);
			var v12 := DC17(fm35(i2, globalState));
			match (v12) {
				case DC17(cf23) =>
					var v13: map<int, seq<int>> := map[v11.f12 := seq(abs(832), i5  => (p0))];
					var v14: array<seq<int>> := new seq<int>[28] [seq(abs(0x366), i4  => (v11.f12)), v4 + v4, v4 + v4, v4 + v4[safeIndex(0x20b, |v4|) := v10], (if (v11.f12 in v13) then v13[v11.f12] else v4) + v4, v4, v4 + [i2, fm7(globalState)], v4, v4, if (f4) then v4 else [i2], v4, v4, v4, v4, v4, seq(abs(0x182), i6  => (v10)), [-v11.f12], v4, fm36(true, globalState), [|v8|, -v10], v4, [i2, v11.f12, |v1|, v10], v4, v4[safeIndex(-i2, |v4|) := |v6|], seq(abs(0x279), i7  => (|[f4, false]|)), (seq(abs(0x210), i8  => (p0))) + [-466], fm36(fm8(|multiset{v9, f4, v9}|, false, v6, v4, globalState), globalState), v4];
					v14[safeIndex(161, v14.Length)] := v4;
					v14[safeIndex(161, v14.Length)] := v4;
					v10 := v11.f12;
					var v15: seq<string> := [v6];
					var v16: seq<string> := [v6, v15[safeIndex(f6, |v15|)], v6];
					var v17 := DC10(|v16[safeIndex(|[v10]|, |v16|)]|);
					v17 := v17;
					v9 := v9;
			}
			
		}
		var v18: array<seq<bool>> := new seq<bool>[23];
		forall i9 | 0 <= i9 < v18.Length {
			v18[i9] := v1;
		}
		var v19 := true;
		v19 := f4;
		match (DC2(-988)) {
			case DC1() =>
				var v20: map<int, int> := map[fm3(p0, v19, globalState) := p0];
				var v21 := "vtke";
				v20 := v20[|v21| := safeDivisionInt(|v5|, -p0)];
				v19 := f4;
				var v22 := -0x3af;
				var v23: set<array<int>> := {v0, v0};
				v22 := |v23|;
				v19 := f4;
			case DC2(cf1) =>
				var v24: map<set<int>, int> := map[v5 := p0];
				cf1 := safeDivisionInt(if (v5 in v24) then v24[v5] else cf1, f6);
				var v25: map<bool, int> := map[f4 := f6];
				var v26 := DC21(if (!f4) then v25 else v25);
				match (v26) {
					case DC22(cf30, cf31, cf32, cf33, cf34) =>
						var v27 := 't';
						var v28 := m9(fm0(v27, globalState), globalState);
						cf1 := (cf1 * cf1) + -0x230;
						var v29: array<map<bool, bool>> := new map<bool, bool>[23];
						var v30: map<array<map<bool, bool>>, bool> := map[v29 := v19];
						var v31 := DC23(v29);
						v30 := v30[v31.cf35 := f4];
						var v32: multiset<int> := multiset{p0};
						v19 := if (v19) then [f6] <= v4 else v32 <= v32;
					case DC21(cf29) =>
						v2 := v2;
						v0[safeIndex(483, v0.Length)] := f6;
						var v33: set<bool> := {v19};
						var v34: map<int, set<bool>> := map[p0 := v33];
						var v35 := 'k';
						var v36: C0 := new C0(f6);
						var v37: multiset<int> := multiset{p0, -(p0 * f6), cf1 * 0x182, p0, fm9(f4, f4, v36.f12, globalState)};
						var v38 := "wcnfyfu";
						var v39: multiset<multiset<int>> := multiset{v37};
						cf1, cf1, cf1, v0[safeIndex(483, v0.Length)] := |map[if (f6 in v34) then v34[f6] else v33 := map[v35 := v36]]|, if (|v38| in v37) then v37[|v38|] else cf1, if ((v37 * v37) in v39) then v39[v37 * v37] else p0, cf1;
						var v40: map<int, string> := map[safeDivisionInt(-194, 0x155) := "mou"];
						v40 := v40[v0[safeIndex(483, v0.Length)] := seq(abs(0x46), i10  => (v35))];
						var v41: map<bool, char> := map[true := v35];
						var v42: map<bool, bool> := map[v19 := f4];
						var v43: map<int, int> := map[cf1 := |v42|];
						var v44: map<int, bool> := map[f6 := false];
						v41 := v41[v0[safeIndex(483, v0.Length)] <= (if (cf1 in v43) then v43[cf1] else |v44|) := v35];
				}
				
				v19 := v19;
				match (v3) {
					case DC16(cf20, cf21, cf22) =>
						var v45: T1 := new C1(-cf1, f6, f5, 0x1ea, false);
						var v46: map<T1, bool> := map[v45 := v45.f4];
						var v47: map<char, map<T1, bool>> := map[fm1(p0, cf21, globalState) := v46];
						v47 := v47;
						var v48 := 'r';
						var v49: map<char, char> := map[v48 := v48];
						var v50: multiset<map<char, char>> := multiset{v49};
						var v51: set<bool> := {v19, false};
						var v52: multiset<int> := multiset{f6};
						cf20 := multiset{|v50|, |v51|, p0, cf1} >= (v52 + v52);
						var v53: multiset<bool> := multiset{v19};
						v1 := DC13(v45.f6, v19, |v53|, v1).cf17;
						var v54 := "wfi";
						var v55: map<multiset<int>, string> := map[v52 := v54];
						var v56, v57 := m0(v55, cf22 <= v4, v52, v45.f4, globalState);
					case DC15(cf19) =>
						cf1 := -safeModuloInt(fm9(v19, true, f6, globalState), if (f4 in v25) then v25[f4] else p0);
						var v58: array<bool> := new bool[2] [v19, f4];
						var v59: map<array<bool>, int> := map[v58 := cf1];
						var v60 := DC31(f5);
						var v61: C1 := new C1(f6, f6, v60.cf50, f6, v19);
						var v62: map<array<bool>, C1> := map[v58 := v61];
						var v63: map<int, int> := map[cf1 := |v62|];
						var v64: T1 := new C1(|v63|, f6, f5, -|v4|, f4);
						var v65: map<T1, array<bool>> := map[v64 := v58];
						v59 := v59[if (v64 in v65) then v65[v64] else v58 := v61.f16];
						var v66: multiset<bool> := multiset{false, !v19};
						var v67 := DC28(p0, |v66|, false, v19 || v19, fm9(v19, v19, v61.f16, globalState));
						v67 := v67;
						var v68: map<int, char> := map[v61.f16 := 'm'];
						v61.f16 := safeModuloInt(p0, |v68|);
				}
				
			case DC0(cf0) =>
				var v69 := "qlh";
				var v70: map<multiset<int>, string> := map[multiset(fm39(v69, DC27([[f4]]), fm0('r', globalState), globalState)) := v69];
				var v71, v72 := m0(v70 + v70, f4, multiset(v4 + (seq(abs(68), i11  => (|map[v69 := 'i']|)))), !v19, globalState);
				if (false) {
					v71 := v19;
					v0 := v0;
					var v74 := new C1(-v72, v4[safeIndex(f6, |v4|)], f5, |(v5 + (set v73 : int | (0x30b <= v73) && (v73 < -0x1e9) :: (v73 - p0)))|, v72 > v72);
					v19 := f4;
					var v75: array<bool> := new bool[7];
					v75 := v75;
				} else {
					cf0 := v72;
					var v76: map<int, int> := map[|v1| := f6];
					var v77: multiset<int> := multiset{0xe2};
					var v78: seq<multiset<int>> := [v77];
					v71 := v1[safeIndex(safeDivisionInt(|v76|, |v78[safeIndex(0x9, |v78|)]|), |v1|)];
					var v79: map<bool, bool> := map[cf0 < cf0 := v71];
					v19 := if (f4 in v79) then v79[f4] else v19;
					var v80 := 'v';
					v80 := v80;
					var v81: array<bool> := new bool[24];
					v81 := DC5(v81).cf5;
				}
				
				var v82 := 'l';
				v82 := v82;
				var v85: set<bool> := {f4, v71, fm0(v82, globalState)};
				var v86: array<set<int>> := new set<int>[6] [v5, set v83 : int | v83 in v4 :: (safeDivisionInt(v83, |(seq(abs(-0x3a4), i12  => ('v')))|)), v5, if (v71) then v5 else set v84 : int | v84 in v4 :: (safeDivisionInt(v84, --0xdf)), v5, fm35(|v85|, globalState)];
				v86[safeIndex(869, v86.Length)] := v5;
				var v87: multiset<int> := multiset{p0};
				v86[safeIndex(869, v86.Length)] := v5 * {if (f6 in v87) then v87[f6] else f6, p0, f6, cf0, p0};
		}
		
		var v88: map<bool, int> := map[v19 := f6];
		r0 := v88;
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) {
		var v0: multiset<int> := multiset{f6};
		var v1: array<int> := new int[3] [-(if (f6 in v0) then v0[f6] else f6), f6, -(f6 * fm7(globalState))];
		v1[safeIndex(699, v1.Length)] := f6;
		v1[safeIndex(699, v1.Length)] := f6;
		var v2: map<int, array<int>> := map[v1[safeIndex(699, v1.Length)] := v1];
		var v3: map<int, array<int>> := map[442 := if (f6 in v2) then v2[f6] else v1];
		var v4 := "ivk";
		var v5: seq<array<int>> := [v1, v1, v1, v1];
		var v6: array<array<int>> := new array<int>[24] [v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, if (!f4) then v1 else v1, v1, v1, if (f4) then v1 else v1, if (|v4| in v3) then v3[|v4|] else v1, v1, v1, v1, v1, v1, v5[safeIndex(v1[safeIndex(699, v1.Length)], |v5|)], v1];
		v6[safeIndex(463, v6.Length)] := v1;
		var v7 := DC7(|fm37(p0, f4, f6, p0, globalState)|);
		var v8: multiset<D2> := multiset{v7, v7};
		var v9: set<bool> := {p0, p0};
		var v10: map<bool, bool> := map[|v8| == v1[safeIndex(699, v1.Length)] := |v9| != v1[safeIndex(699, v1.Length)]];
		var v11 := true;
		var v12 := DC32(v1);
		var v13: map<array<int>, array<int>> := map[v1 := v12.cf51];
		v6[safeIndex(463, v6.Length)], v1[safeIndex(699, v1.Length)], v10, v11 := if (v1 in v13) then v13[v1] else v1, f6, v10 + (v10 + v10), f4 ==> f4;
		var v14: array<array<char>> := new array<char>[18];
		var v15: seq<bool> := [f4];
		var v16: map<bool, int> := map[f4 := v1[safeIndex(699, v1.Length)]];
		var v17: seq<map<bool, int>> := [v16, map[p0 := -v1[safeIndex(699, v1.Length)]], v16];
		v1[safeIndex(699, v1.Length)], v1[safeIndex(699, v1.Length)], v11, v11, v14 := v1[safeIndex(699, v1.Length)], |(if (f4) then v15 + v15 else v15 + v15)|, v16 !in (v17 + v17), v11, v14;
		var v18: T0 := new C2(seq(abs(587), i0  => ('f')), p0);
		var v19: map<bool, T0> := map[v18.f4 := v18];
		var v20: map<int, T0> := map[f6 := v18];
		var v21: seq<T0> := [v18, if (p0 in v19) then v19[p0] else if (-0xe4 in v20) then v20[-0xe4] else v18, v18, v18, v18];
		v1[safeIndex(699, v1.Length)] := |v21| - |(seq(abs(-0x2bd), i1  => ('u')))|;
		var v22: array<string> := new string[14](i2 => "q" + "fag");
		var v23: seq<string> := [v4, seq(abs(0x333), i3  => ('h'))];
		var v24 := 'f';
		v22[safeIndex(290, v22.Length)] := (v23[safeIndex(f6, |v23|)])[safeIndex(f6, |v23[safeIndex(f6, |v23|)]|) := v24] + "nusipv";
		v22[safeIndex(290, v22.Length)] := seq(abs(0), i4  => ('k'));
		if (fm0(v24, globalState)) {
			var v25: array<char> := new char[21];
			v14[safeIndex(647, v14.Length)] := v25;
			v14[safeIndex(647, v14.Length)] := v25;
			v11 := !p0;
			var v26 := DC26(v18.f4);
			var v27: map<D11, bool> := map[v26 := v18.f4];
			var v28: seq<map<D11, bool>> := [v27];
			var v29 := DC35(v28);
			var v30: seq<int> := [|(seq(abs(0x1c3), i5  => (v24)))|, v1[safeIndex(699, v1.Length)]];
			var v31: seq<int> := [v30[safeIndex(v1[safeIndex(699, v1.Length)], |v30|)]];
			var v32: map<int, map<bool, bool>> := map[|v30| := v10];
			var v33: seq<map<int, map<bool, bool>>> := [v32, v32];
			var v34 := new C1(-|v29.cf58|, |v33[safeIndex(|{v11}|, |v33|)]|, f5, v31[safeIndex(v1[safeIndex(699, v1.Length)], |v31|)] - v18.fm7(globalState), true);
			var v35: set<multiset<bool>> := {multiset(v15)};
			var v36: array<bool> := new bool[27] [fm0(v24, globalState), !(f6 <= v34.f15), v35 >= v35, f4, false, v11, f4, f4, true, v18.f4, f4, p0, v11, v9 >= v9, v18.f4, f4, p0, v18.f4 || p0, p0, v18.f4, f4, p0, p0, v11, true, v18.f4, p0];
			var v37: set<string> := {v22[safeIndex(290, v22.Length)]};
			v36[safeIndex(815, v36.Length)] := v37 > v37;
			v1[safeIndex(699, v1.Length)], v11, v36[safeIndex(815, v36.Length)] := safeModuloInt(f6, v1[safeIndex(699, v1.Length)] * v34.f16), p0, v18.f4;
			v11 := !DC28(fm3(|(v15[safeIndex(v1[safeIndex(699, v1.Length)], |v15|) := !p0])[safeIndex(v34.f16, |v15[safeIndex(v1[safeIndex(699, v1.Length)], |v15|) := !p0]|) := v36[safeIndex(815, v36.Length)]]|, f4, globalState), v34.f16, v18.f4, v36[safeIndex(815, v36.Length)], f6).cf46;
		} else {
			v1[safeIndex(699, v1.Length)] := f6;
			v1[safeIndex(699, v1.Length)] := f6;
			v11 := true;
			var v38: C2 := new C2(v22[safeIndex(290, v22.Length)], v11);
			var v39: multiset<C2> := multiset{v38, v38, v38};
			var v40: array<bool> := new bool[10] [v38 !in v39[v38 := abs(f6)], true, fm0(v24, globalState), if (v18.f4) then v11 else true, v9 != v9, v0 == v0, f6 > f6, v11, f4, v18.f4];
			v40[safeIndex(580, v40.Length)] := v11;
			v40[safeIndex(580, v40.Length)] := v18.f4;
			var v41 := new C1(v1[safeIndex(699, v1.Length)], 0x326, f5, f6 * f6, v11);
		}
		
		r0 := v24;
		var v42: array<D1> := new D1[23];
		r1 := v42;
	}
	method m9(p0: bool, globalState: GlobalState) returns (r0: bool) {
		r0 := !p0;
		var v0: array<string> := new seq<char>[22](i0 => (seq(abs(773), i1  => ('p'))) + "uqvrar");
		v0[safeIndex(721, v0.Length)] := "t";
		var v1 := "w";
		v0[safeIndex(721, v0.Length)] := fm37(!f4, f4, |v1|, f4, globalState) + v1;
		var v2 := 'v';
		v0[safeIndex(721, v0.Length)] := v1[safeIndex(-f6, |v1|) := v2];
		var v3: seq<int> := [f6, f6, f6];
		var v4 := new C1(-f6, |(v3 + v3)|, f5, f6, !f4);
		if (!f4) {
			var v5: array<bool> := new bool[5](i2 => f4);
			var v6: map<bool, array<bool>> := map[true := v5];
			var v7: seq<array<bool>> := [v5, v5, v5];
			v5 := if (f4 in v6) then v6[f4] else v7[safeIndex(v4.f15, |v7|)];
			var v8 := new C2(v1, p0);
			v4.f16 := v4.f15;
			var v9: map<int, bool> := map[v4.f15 := f4];
			v9 := map[v4.f15 := f4] + map[v4.f15 := p0];
			var v10 := DC22(seq(abs(0x317), i3  => (v2)), false, f4, v5, 0xb8);
			var v11 := DC4(v10.cf34, |"lroyicek"| - 0x1c7);
			v11 := v11;
		} else {
			r0 := "vtossvsa" < (v0[safeIndex(721, v0.Length)] + v0[safeIndex(721, v0.Length)]);
			r0 := (if (f4) then v4.f16 else f6) >= |v0[safeIndex(721, v0.Length)]|;
			var v12: array<char> := new char[2] [v2, v2];
			var v13: map<bool, array<char>> := map[p0 := v12];
			var v14: seq<array<char>> := [v12];
			var v15: array<array<char>> := new array<char>[8] [v12, v12, v12, v12, v12, if (p0 in v13) then v13[p0] else v14[safeIndex(-v4.f15, |v14|)], v12, v12];
			v15[safeIndex(472, v15.Length)] := v12;
			var v16: multiset<int> := multiset{v4.f16};
			v0[safeIndex(721, v0.Length)], v4.f16, v15[safeIndex(472, v15.Length)] := ("bsbsmmac" + "ps") + v0[safeIndex(721, v0.Length)], if (v4.f16 in v16) then v16[v4.f16] else v4.f15, v12;
			var v18: array<set<int>> := new set<int>[1](i4 => set v17 : int | v17 in v3 :: (v17 + -0x8e));
			var v19: array<array<set<int>>> := new array<set<int>>[4] [v18, if (p0) then v18 else v18, if (p0) then v18 else v18, v18];
			v19[safeIndex(373, v19.Length)] := v18;
			v19[safeIndex(373, v19.Length)] := if (f4) then v18 else v18;
			if (true) {
				var v20: array<map<bool, int>> := new map<bool, int>[25](i5 => map[f4 := |v16|]);
				var v21: map<bool, int> := map[f4 := v4.f16];
				v20[safeIndex(655, v20.Length)] := v21;
				var v22: map<bool, bool> := map[f4 := f4];
				var v23: map<bool, bool> := map[f4 := if (f4 in v22) then v22[f4] else p0];
				var v24: seq<map<bool, bool>> := [v22, v23, v22];
				r0, v15[safeIndex(472, v15.Length)], v20[safeIndex(655, v20.Length)], v16 := v24 == v24, v12, v21, v16;
				var v25: array<map<int, int>> := new map<int, int>[22];
				var v26: multiset<bool> := multiset{p0};
				var v27: map<int, int> := map[|v26| := 0x224];
				v25[safeIndex(11, v25.Length)] := v27;
				var v28: set<bool> := {p0};
				v25[safeIndex(11, v25.Length)] := map[|v28| := 580];
				r0 := f4;
				v4.f16 := v4.f16;
				v4.f16 := v4.f16 + (v4.f15 * v4.f16);
			} else {
				var v29: map<string, int> := map[seq(abs(4), i6  => (v2)) := 0x133];
				var v30: map<bool, map<string, int>> := map[f4 := v29];
				var v31: array<bool> := new bool[16](i7 => true);
				var v32 := v4.m11(p0, p0, if (p0 in v30) then v30[p0] else map[v0[safeIndex(721, v0.Length)] := v4.f16], v31, globalState);
				var v33: map<int, bool> := map[v4.f15 := p0];
				var v34: seq<bool> := [if (v4.f15 in v33) then v33[v4.f15] else false];
				var v35: seq<seq<bool>> := [v34, v34];
				var v36 := DC27(v35);
				v36 := v36.(cf42 := v35[safeIndex(|v1|, |v35|) := v35[safeIndex(v4.f16, |v35|)]]).(cf42 := v35);
				f5[safeIndex(533, f5.Length)] := v31;
				f5[safeIndex(533, f5.Length)] := v31;
				var v37: map<bool, array<bool>> := map[f4 := v31];
				var v38: map<char, array<bool>> := map[v2 := if (f4 in v37) then v37[f4] else f5[safeIndex(533, f5.Length)]];
				var v39 := DC37(v38);
				r0, v38 := f4, v39.cf62;
				var v40, v41 := v4.m2(p0, globalState);
			}
			
		}
		
		var v42: map<bool, bool> := map[p0 := p0];
		var v43: map<map<bool, bool>, char> := map[v42 := v2];
		var v44: map<int, string> := map[f6 := "mcrpadd"];
		var v45, v46, v47, v48 := m10(multiset{|v43|, f6, v4.f16}, v1 < (if (v4.f16 in v44) then v44[v4.f16] else "h"), globalState);
		r0 := p0;
	}
	method m10(p0: multiset<int>, p1: bool, globalState: GlobalState) returns (r0: bool, r1: int, r2: int, r3: bool) {
		var v0: array<int> := new int[28];
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := safeModuloInt(i0, f6);
		}
		for i1 := f6 to -0x16d {
			var v1 := DC32(v0);
			match (v1) {
				case DC33(cf52, cf53, cf54, cf55, cf56) =>
					cf52 := !true;
					var v2: seq<bool> := [cf52, false];
					var v3 := DC13(f6, cf52, i1, v2);
					r0 := (if (p1) then v3 else v3).cf15;
					var v4: array<bool> := new bool[14];
					v4 := v4;
					r2 := -i1;
				case DC32(cf51) =>
					var v5: array<bool> := new bool[16];
					v5[safeIndex(287, v5.Length)] := f4;
					var v6 := DC0(|map[p1 := i1]|);
					var v7: set<D0> := {v6};
					var v8: multiset<bool> := multiset{f4};
					var v10: set<set<D0>> := {v7, v7, set v9 : D0 | v9 in [fm38(|(v8[f4 := abs(i1)])[f4 := abs(i1)]|, i1, p1, f6, globalState)] :: (v9)};
					var v11: multiset<int> := multiset{|({v7} - v10)|};
					v5[safeIndex(287, v5.Length)], v11, r1 := f4 ==> f4, v11, |[|"skjggxj"|]|;
					v5[safeIndex(287, v5.Length)] := v5[safeIndex(287, v5.Length)] && (v5[safeIndex(287, v5.Length)] <==> false);
					var v12: seq<int> := [-f6, f6];
					r2 := |v12|;
					v5[safeIndex(287, v5.Length)] := p1;
			}
			
			var v13: array<string> := new string[1];
			var v14 := DC26(f4);
			v13, v14 := v13, v14;
			var v15: map<int, bool> := map[f6 := p1];
			var v16 := new C1(|v15| * i1, i1, f5, i1, !false <== f4);
			var v17 := 'b';
			r3 := fm0(v17, globalState);
		}
		if (f4) {
			var v18 := DC10(f6 * f6);
			match (v18) {
				case DC9(cf8, cf9, cf10) =>
					r2 := f6;
					var v19 := "uig";
					var v20: map<multiset<int>, string> := map[p0 := v19];
					var v21 := DC0(0x2ff);
					var v22 := 't';
					var v23: set<string> := {v19};
					var v24, v25 := m0(v20, cf8, fm6(DC0(|p0|), v21, map[cf9 := v22], |v23|, globalState), cf10, globalState);
					var v26: array<map<bool, D12>> := new map<bool, D12>[8];
					var v27: map<D3, map<bool, D12>> := map[v18 := fm58(globalState)];
					var v28: seq<int> := [f6];
					var v29: seq<bool> := [fm8(v25, false, v19, v28, globalState)];
					var v30: seq<seq<bool>> := [v29];
					var v31 := DC27(v30[safeIndex(|v29|, |v30|) := v29]);
					var v32: map<bool, D12> := map[f4 := v31];
					v26[safeIndex(776, v26.Length)] := if (v18 in v27) then v27[v18] else v32;
					var v33 := DC40(v32);
					v26[safeIndex(776, v26.Length)] := v32 + (v32 + DC40(v33.cf66).cf66);
					var v34: map<int, bool> := map[v25 := cf10];
					var v35 := new C1(|v34| + cf9, safeModuloInt(|v19|, v25), f5, f6, v24);
				case DC10(cf11) =>
					r1 := 0x1eb;
					var v36 := "fvlmlhsge";
					r2 := |(v36 + v36)|;
					var v37 := new C1(f6, safeModuloInt(0x2f1, f6), f5, safeModuloInt(f6, f6), f4);
					v0 := v0;
				case DC8(cf7) =>
					var v38 := 'e';
					var v39: map<char, bool> := map[v38 := false];
					r0 := (if (v38 in v39) then v39[v38] else !f4) <== p1;
					var v40 := "honftv";
					r2 := f6 - fm9(f4, f4, |v40|, globalState);
					var v41: array<bool> := new bool[26](i2 => true);
					var v42: map<char, array<bool>> := map['p' := v41];
					var v43 := DC37(v42);
					var v44: map<bool, D19> := map[f4 || true := v43];
					v44 := v44[f4 := v43];
					var v45: C2 := new C2(v40, f4);
					var v46: seq<C2> := [v45, v45, v45];
					var v47: seq<seq<C2>> := [v46, [v45]];
					r3 := v47[safeIndex(f6, |v47|) := v46] != v47;
				case DC11(cf12) =>
					var v48: seq<int> := [f6, f6];
					v48 := seq(abs(0x388), i3  => (|([p1] + [f4, f4])|));
					var v49 := "bwcahcko";
					var v50 := 'h';
					var v51 := DC19(v49);
					v49 := (if (f4) then "ylnncpc" else v49)[safeIndex(f6, |(if (f4) then "ylnncpc" else v49)|) := v50] + v51.cf25;
					var v52 := DC28(f6, -f6, f4, p1, |(if (f4) then "iunuhhphg" else seq(abs(651), i4  => (v50)))|);
					v52, r0, v49 := v52, true || f4, v49;
					r3 := p1;
			}
			
			var v53 := "ukokag";
			r1 := |fm37(p1, !p1, |v53|, fm8(f6, p1, v53, [|v53|, f6], globalState), globalState)|;
			r2 := -f6;
			var v54: set<bool> := {f4, p1};
			var v55: seq<bool> := [v54 >= v54];
			r0 := v55[safeIndex(f6 - 0x30b, |v55|)];
			var v56 := 'i';
			v56 := v56;
		} else {
			r1 := f6;
			var v57 := "kcdy";
			var v58: map<bool, string> := map[f4 := v57];
			var v59: map<string, map<bool, string>> := map[v57 := v58];
			var v60: map<int, map<bool, string>> := map[f6 := v58];
			v59 := v59[v57 := (if (|v57| in v60) then v60[|v57|] else v58) + v58];
			var v61: seq<int> := [f6];
			r3 := fm8(f6, p1 && f4, v57, v61, globalState);
			v0[safeIndex(459, v0.Length)] := f6;
			v0[safeIndex(459, v0.Length)] := f6;
			var v62: array<set<int>> := new set<int>[14];
			var v63: set<int> := {f6, |v57|};
			v62[safeIndex(216, v62.Length)] := v63;
			v62[safeIndex(216, v62.Length)] := v63;
		}
		
		var v64 := "iktm";
		for i5 := f6 - |v64| to f6 {
			var v65: seq<int> := [f6];
			r2 := v65[safeIndex(f6, |v65|)];
			var v66: multiset<bool> := multiset{f4, !f4};
			match (DC28(f6, |p0[i5 := abs(|v66|)]|, p1, true, i5).(cf43 := f6).(cf47 := |v64|)) {
				case DC28(cf43, cf44, cf45, cf46, cf47) =>
					var v67 := DC0(0x36f - cf47);
					v67 := v67;
					v64 := v64;
					var v68: set<int> := {f6, cf47};
					cf46 := i5 == -(|v68| * v65[safeIndex(-|v68|, |v65|)]);
					var v69: map<int, bool> := map[cf47 := p1];
					var v71: multiset<map<int, bool>> := multiset{if (cf45) then v69 else v69, map v70 : int | (0x1a0 <= v70) && (v70 < -0x1c0) :: (safeModuloInt(v70, f6)) := (cf46), v69};
					v71 := v71;
				case DC27(cf42) =>
					r3 := (|v64| <= f6) ==> f4;
					r1 := i5;
					var v72: array<bool> := new bool[3] [if (!!p1) then !p1 else f4, f4, f4];
					v72[safeIndex(154, v72.Length)] := p1;
					r0, v72[safeIndex(154, v72.Length)] := p1, p1;
					v72[safeIndex(154, v72.Length)] := v72[safeIndex(154, v72.Length)];
			}
			
			var v73: map<bool, int> := map[f4 := |v64|];
			var v74 := DC21(v73);
			var v75: map<D10, int> := map[v74 := |v64|];
			v75 := fm59(globalState);
			r1 := -f6;
		}
		var v76: seq<int> := [f6];
		v76 := if (p1) then [-0x377, f6, f6, f6, f6] else seq(abs(0x383), i6  => (f6));
		var v77: map<bool, int> := map[f4 := f6 - f6];
		v77 := fm44(globalState) + (v77 + v77);
		r0 := true;
		r1 := f6;
		r2 := f6;
		r3 := !f4;
	}
}

class C4 extends T1, T0 {
	constructor (f5 : array<array<bool>>, f6 : int, f4 : bool) {
		this.f5 := f5;
		this.f6 := f6;
		this.f4 := f4;
	}
	
	function fm8(p0: int, p1: bool, p2: seq<char>, p3: seq<int>, globalState: GlobalState): bool {
		!f4
	}
	function fm9(p0: bool, p1: bool, p2: int, globalState: GlobalState): int {
		f6
	}
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> {
		multiset{safeModuloInt(117, 0x34a), if (f4) then f6 else f6, f6, safeDivisionInt(f6, f6), f6}
	}
	function fm7(globalState: GlobalState): int {
		|(match DC19("sxfde") {
			case DC20(cf26, cf27, cf28) => "uaimtxsem" + (seq(abs(0x2e8), i0  => (cf27)))
			case DC19(cf25) => cf25
		})|
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) {
		var v0 := "tv";
		var v1: multiset<int> := multiset{|map[f4 := f4]|, |[v0]|};
		var v2, v3 := m0(map[v1 := v0], !(if (f4) then f4 else f4), v1, f4, globalState);
		var v4: array<char> := new char[9](i0 => 'x');
		var v5 := 'i';
		v4[safeIndex(290, v4.Length)] := v5;
		var v6: seq<int> := [safeDivisionInt(v3, f6)];
		var v7: map<bool, int> := map[v2 := -|fm30(v6, true, globalState)|];
		var v8: map<bool, bool> := map[v2 := f4];
		var v9: map<map<bool, bool>, int> := map[v8 := |{f4, v2}|];
		var v11: array<int> := new int[22] [-0x32b, p0, |(seq(abs(-0x1dc), i1  => (v5)))|, f6, f6, f6, |v7|, f6, f6, v3, |[p0, p0, -v3, f6]|, f6, p0, f6, |"ejh"|, fm3(v3, f4, globalState), v3, v3, if (map[v2 := f4] in v9) then v9[map[v2 := f4]] else |(map v10 : int | (374 <= v10) && (v10 < 0x132) :: (v10 - f6) := (|map[|multiset(v6)| := f4]|))|, p0, v3, f6];
		var v12: seq<array<int>> := [v11];
		var v13: multiset<array<int>> := multiset{v11, v12[safeIndex(v3, |v12|)]};
		var v14: multiset<multiset<array<int>>> := multiset{v13, v13 - v13[v11 := abs(p0)]};
		v4[safeIndex(290, v4.Length)], v3, v3, v6 := v5, f6, if ((v13 + v13) in v14) then v14[v13 + v13] else if (v2) then v3 else p0, v6;
		var v15 := new C0(safeDivisionInt(v3, |v6|));
		var v16: map<char, bool> := map[v5 := v2];
		v16 := v16[v4[safeIndex(290, v4.Length)] := f4];
		for i2 := -659 to -(v15.f12 - 281) {
			var v17 := new C0(v3);
			var v18: map<int, set<int>> := map[|"jl"| := fm31(i2, v3, |v0|, f4, globalState)];
			var v19: map<int, bool> := map[v3 := f4];
			var v20: set<int> := {|v1|, |v7|, DC0(f6).cf0, |v19|, |v6|};
			var v23: array<set<int>> := new set<int>[8] [if (v15.f12 in v18) then v18[v15.f12] else v20, set v21 : int | (380 <= v21) && (v21 < 0x2ac) :: (v21 - v15.f12), {f6, v15.f12, v17.f12, |v1|}, (fm32(558, v2, globalState)).cf23, v20, {i2}, v20, set v22 : int | (0x314 <= v22) && (v22 < 0x91) :: (v22 - v17.f12)];
			var v24: set<bool> := {v17.fm17(f4, globalState), !!f4, v2, f4, !f4};
			var v25: map<set<bool>, bool> := map[v24 := v2];
			var v26: map<array<set<int>>, int> := map[if (v2) then v23 else v23 := |v25|];
			v26 := v26[v23 := safeModuloInt(p0, v15.f12)];
			v11[safeIndex(911, v11.Length)] := |v7|;
			var v27: seq<bool> := [f4, v2, v2, v2];
			v11[safeIndex(911, v11.Length)] := |(v27[safeIndex(|v27|, |v27|) := f4] + v27)|;
			var v28: multiset<bool> := multiset{v15.fm17(f4, globalState), v15.fm17(true, globalState), v2};
			v28 := v28;
		}
		v4[safeIndex(290, v4.Length)] := v4[safeIndex(290, v4.Length)];
		r0 := v7;
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) {
		var v0: map<int, int> := map[-f6 * f6 := fm7(globalState)];
		var v1: array<int> := new int[4];
		var v2: map<array<int>, map<int, int>> := map[v1 := map[f6 := -f6]];
		v0 := if (v1 in v2) then v2[v1] else map[f6 := f6];
		var v3 := "tjatmipct";
		var v4: seq<int> := [|v3|, f6];
		v4 := seq(abs(355), i0  => (f6));
		for i1 := f6 to |v3| {
			var v5: seq<bool> := [p0];
			var v6 := DC13(safeModuloInt(|v0|, i1), f4, -262, v5);
			match (v6) {
				case DC13(cf14, cf15, cf16, cf17) =>
					var v7 := 'y';
					cf16 := if (v7 !in v3) then cf14 else cf14;
					var v8: C0 := new C0(f6);
					var v9: map<C0, int> := map[v8 := i1];
					var v10: map<bool, seq<char>> := map[cf15 := v3];
					var v11: array<bool> := new bool[22] [f4, f4, p0, p0, f4, f4, cf15, p0, p0, f4, !!cf15, f4, f4, cf15, !p0, p0, f4, true, cf15, fm8(cf16, cf15, if (p0 in v10) then v10[p0] else v3, v4, globalState), cf15, f4];
					v0 := v0[|v9| - 0x1c2 := DC22(v3, p0, false, v11, v8.f12).cf34];
					v1[safeIndex(991, v1.Length)] := 0x310;
					v1[safeIndex(991, v1.Length)] := i1 * 0x39d;
					var v12 := DC2(cf16);
					var v13: multiset<char> := multiset{v7};
					var v14: T1 := new C3(f5, |v13|, fm0('f', globalState));
					var v15: seq<T1> := [v14, v14, v14, v14, v14];
					var v16: array<D0> := new D0[21] [v12, v12, DC2(cf16).(cf1 := |v5|), v12, v12, v12, v12, DC2(0x123).(cf1 := cf16), DC2(0x1af), v12.(cf1 := cf16), v12, v12, v12.(cf1 := v1[safeIndex(991, v1.Length)]), v12, v12, DC2(|v15|), v12, DC2(i1), v12, v12, DC2(v8.f12)];
					v16 := new D0[22](i2 => v12);
				case DC12(cf13) =>
					var v17: map<int, array<int>> := map[f6 := v1];
					v17 := v17[i1 := v1];
					var v18 := 277;
					var v20: map<int, char> := map[-i1 := cf13];
					var v21: set<map<int, char>> := {v20};
					v18 := safeDivisionInt(|(map v19 : int | (0x1bf <= v19) && (v19 < -0x11a) :: (safeDivisionInt(v19, 0x250)) := (i1))|, safeDivisionInt(|v21|, f6));
					var v22: array<seq<int>> := new seq<int>[29];
					v22[safeIndex(528, v22.Length)] := v4;
					v22[safeIndex(528, v22.Length)] := [v18];
					var v23: array<char> := new char[21];
					v23[safeIndex(885, v23.Length)] := cf13;
					v23[safeIndex(885, v23.Length)] := cf13;
			}
			
			var v24 := true;
			var v25: multiset<int> := multiset{i1, f6};
			var v26: set<bool> := {p0, f4};
			var v27: seq<seq<int>> := [v4, seq(abs(0x3d4), i3  => (i1))];
			v24 := fm8(if (|v26| in v25) then v25[|v26|] else 622, v27 == [v4], v3, v4 + v4, globalState);
			var v28 := DC24(f4, f4, |v3|);
			var v29 := DC41(if (v24) then i1 else v28.cf38);
			v29 := DC41(0x82);
			var v30 := 174;
			var v31: set<int> := {i1};
			var v32 := DC17(v31);
			var v33: set<D7> := {v32};
			v30 := -(-v30 - f6) + (if (f4) then v30 else |v33|);
		}
		var v34 := false;
		v1[safeIndex(485, v1.Length)] := f6 * f6;
		var v35 := 0x3d2;
		var v36: map<int, bool> := map[v4[safeIndex(v35, |v4|)] := v34];
		var v37: map<seq<bool>, map<int, bool>> := map[[f4] := v36];
		var v38: seq<multiset<int>> := [multiset{f6}];
		var v39: multiset<bool> := multiset{p0};
		v34, v1[safeIndex(485, v1.Length)], v35, v34, v34 := v37 != fm60(|v38|, f6, v35, globalState), v35, safeDivisionInt(|v39|, v35), !!v34, f4;
		var v40: array<D12> := new D12[28](i4 => DC27([[false], [p0]]));
		var v41: set<bool> := {f4, f4};
		var v42 := DC27(fm61(v41, globalState));
		v40[safeIndex(785, v40.Length)] := v42;
		v35, v40[safeIndex(785, v40.Length)] := v35, v42;
		var v43: C1 := new C1(v35, safeModuloInt(v1[safeIndex(485, v1.Length)], v35), f5, f6, fm8(f6, f4, seq(abs(0x236), i5  => ('g')), seq(abs(-0x214), i6  => (v35)), globalState));
		var v44 := 'f';
		var v45 := DC20(f4, v44, v43.f16);
		var v46: seq<D9> := [v45];
		v43, v34, v34, v1[safeIndex(485, v1.Length)] := if ((if (v1[safeIndex(485, v1.Length)] in v0) then v0[v1[safeIndex(485, v1.Length)]] else v1[safeIndex(485, v1.Length)]) <= v1[safeIndex(485, v1.Length)]) then v43 else v43, p0, p0, -match DC34(v46) {
			case DC34(cf57) => 0x38f
		};
		r0 := v44;
		var v47: array<D1> := new D1[2];
		r1 := v47;
	}
}

class C5 extends T1 {
	const f13 : bool
	var f14 : int
	constructor (f13 : bool, f14 : int, f5 : array<array<bool>>, f6 : int, f4 : bool) {
		this.f13 := f13;
		this.f14 := f14;
		this.f5 := f5;
		this.f6 := f6;
		this.f4 := f4;
	}
	
	function fm8(p0: int, p1: bool, p2: seq<char>, p3: seq<int>, globalState: GlobalState): bool {
		if (f13) then map[f4 := f4] == map[f13 := f13] else -f14 !in multiset{f14}
	}
	function fm9(p0: bool, p1: bool, p2: int, globalState: GlobalState): int {
		0x1e
	}
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> {
		multiset{f6, 309, f6} - multiset{-f14}
	}
	function fm7(globalState: GlobalState): int {
		match DC9(f13, f6, f4) {
			case DC9(cf8, cf9, cf10) => f6
			case DC10(cf11) => f14 + f14
			case DC8(cf7) => 64 - |DC19("fafdlayl").cf25|
			case DC11(cf12) => f6
		}
	}
	function fm26(p0: int, p1: bool, p2: int, globalState: GlobalState): int {
		-f14
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) {
		var v0: multiset<bool> := multiset{f13, f4, f4, f13, f13};
		var v1: map<D6, int> := map[DC15(fm27(f13, 35, v0, f13, globalState)) := |(v0 * v0)|];
		v1 := v1;
		f14 := f6;
		var v2: set<bool> := {f4};
		var v3: seq<set<bool>> := [v2];
		var v4: map<bool, int> := map[f13 := |v3|];
		var v5 := 'a';
		var v6 := DC20(f13, v5, f6);
		var v7 := new C0(if (v6.cf26 in v4) then v4[v6.cf26] else p0);
		var v8: seq<bool> := [f13];
		var v9: seq<seq<bool>> := [[f4, f4, f13, f4, f4]];
		var v11 := "trc";
		var v12: array<int> := new int[1](i0 => safeModuloInt(i0, f6));
		var v13: map<bool, array<int>> := map[f4 := v12];
		var v14: array<int> := new int[20] [f14, -f14, |[0x3b, v7.f12, v7.f12]|, f14 - f14, v7.f12, f6, p0, f6, |map[|multiset(v8)| := v7.f12]|, f6, f14, v7.f12, |v9|, |(set v10 : int | (0xfd <= v10) && (v10 < -0x123) :: (v10 * v7.f12))|, f6 - |v11|, |v13|, f6, safeModuloInt(438, f14), f6, v7.f12];
		v12[safeIndex(396, v12.Length)] := f14;
		var v15 := DC21(fm28(globalState));
		var v16: seq<seq<set<bool>>> := [v3, v3];
		v12[safeIndex(396, v12.Length)], v4, v3 := if (f13) then v7.f12 else v7.f12, v15.cf29, if (f14 != 0x209) then [v2] else v16[safeIndex(v7.f12, |v16|)];
		var v17: map<char, char> := map[v5 := v5];
		v17 := v17['w' := v5];
		var v18: multiset<int> := multiset{|v11|, |v2|};
		var v19: map<multiset<int>, string> := map[v18 := "ydutpac"];
		var v20, v21 := m0(v19 + v19, false, v18, f4, globalState);
		r0 := v4;
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) {
		var v1 := "rptj";
		var v2: multiset<int> := multiset{|v1|};
		var v3: seq<int> := [f14, f6, f14, f14, f6];
		var v4: seq<multiset<int>> := [v2, multiset(v3)];
		var v5, v6 := m0(map v0 : multiset<int> | v0 in v4 :: (v0) := (v1), f4, v2, f6 != f6, globalState);
		v1 := seq(abs(0x250), i0  => ('s'));
		var v7 := new C0(|[f4, !v5, !f13]|);
		var v8 := 'p';
		var v9: map<int, bool> := map[f14 := fm0(v8, globalState)];
		var v10: multiset<char> := multiset{v8};
		v9 := v9[-|v10| := v6 == -f14];
		for i1 := f6 - f6 to 0x1f0 {
			if (false && !p0) {
				var v11: array<bool> := new bool[10](i2 => f4);
				var v12: map<array<bool>, bool> := map[v11 := p0];
				var v13 := new C0(fm26(v7.f12, fm8(f6, true, v1, v3, globalState), |v12|, globalState));
				var v14: map<bool, int> := map[f13 := f6];
				var v15: map<int, int> := map[|v14| := |(seq(abs(218), i3  => (v13.f12)))|];
				v15 := v15[|(map v16 : int | (0x1bc <= v16) && (v16 < 0x203) :: (v16 * v7.f12) := (|multiset{0x262, f14}|))| := |[|(seq(abs(0x52), i4  => (v3)))[safeIndex(|map[f6 := 0x21d]|, |(seq(abs(0x52), i4  => (v3)))|) := v3]|]|];
				var v17: array<int> := new int[6](i5 => safeDivisionInt(i5, v7.f12));
				v17[safeIndex(99, v17.Length)] := 0x28 * f14;
				var v18: multiset<bool> := multiset{f4, p0, f4};
				v17[safeIndex(99, v17.Length)] := safeDivisionInt(f14 - |v18|, f6);
				v5 := f4;
				v6 := v7.f12;
			} else {
				var v19: map<int, int> := map[f14 - v7.f12 := -0x256 - v7.f12];
				v19 := fm29(globalState);
				v6 := v7.f12;
				r0 := v8;
				v8 := v8;
				v6, f14, f14, r0, f14 := f14, v7.f12 * |v1|, f6, v8, v7.f12;
			}
			
			if (v5) {
				f14 := i1;
				f14 := i1 * v6;
				v6 := v7.f12;
				v5 := !f13;
				var v20: array<string> := new string[23];
				v20[safeIndex(239, v20.Length)] := v1;
				var v21: seq<string> := [v1];
				v20[safeIndex(239, v20.Length)] := v21[safeIndex(v7.f12, |v21|)] + (v21[safeIndex(|v3|, |v21|)] + v1);
			} else {
				var v22: T0 := new C2(v1, p0);
				var v23: map<bool, T0> := map[v5 := v22];
				v23 := v23[f13 ==> f13 := v22];
				v5 := !p0;
				var v24 := DC12(v8);
				var v25: array<char> := new char[9](i6 => 'd');
				v25[safeIndex(656, v25.Length)] := v8;
				var v26: seq<bool> := [!v5];
				var v27: array<bool> := new bool[11] [!f4, f4, f4, f13, false, f13, false, p0 ==> true, !f4, v26 == v26, fm0(v8, globalState)];
				v27[safeIndex(225, v27.Length)] := fm8(f14, p0, v1, [f14], globalState);
				var v28 := DC22(v1, p0, f13, v27, i1);
				v24, v25[safeIndex(656, v25.Length)], v27[safeIndex(225, v27.Length)], v27 := fm62(globalState), if (v28.cf31) then v8 else 'd', p0 <==> (i1 == v7.f12), v27;
				v27[safeIndex(225, v27.Length)], v5 := p0, v6 >= (if (v5) then 909 else |v1|);
				v5 := p0;
			}
			
			var v29: array<bool> := new bool[4] [f13, v5, 0x31 < v6, f4];
			v29 := v29;
			var v30: set<bool> := {p0, false};
			f14 := |DC33(v5, |v1|, multiset{p0, true}, v6, v30).cf54|;
		}
		var v31: seq<bool> := [p0, v5, v5];
		var v32: seq<seq<bool>> := [v31];
		var v33 := DC27(v32);
		match (v33) {
			case DC28(cf43, cf44, cf45, cf46, cf47) =>
				var v34: array<map<bool, bool>> := new map<bool, bool>[29](i7 => map[f13 := false]);
				v34 := v34;
				if (fm0(v8, globalState)) {
					var v35: set<bool> := {f13, cf46};
					var v36: map<set<bool>, bool> := map[v35 * v35 := f4];
					v36 := v36[v35 + {f4} := v5 <== f4];
					v5 := f13;
					var v37: C3 := new C3(f5, -f14, p0);
					var v38: map<C3, bool> := map[v37 := f13];
					v38 := v38[v37 := cf46];
					v8 := v8;
					v1 := v1 + v1;
				} else {
					cf44 := 483;
					var v39: array<int> := new int[7](i8 => i8 * 0x92);
					var v40 := DC32(v39);
					v39 := v40.cf51;
					v5 := cf45;
					v39[safeIndex(769, v39.Length)] := v7.f12;
					v39[safeIndex(769, v39.Length)] := 342;
					cf47 := v7.f12;
				}
				
				var v41: set<int> := {fm7(globalState)};
				var v42: map<bool, bool> := map[cf45 := cf45];
				var v43: map<int, int> := map[|v3| := |v42|];
				v5, cf43 := v41 >= ({if (f6 in v43) then v43[f6] else v7.f12, v7.f12} - (set v44 : int | v44 in v2 :: (v44 + |{true, true, false, false}|))), (cf44 * f6) * |"reowmtl"|;
				var v45 := new C2(v1, cf45);
			case DC27(cf42) =>
				var v46 := DC19(v1);
				var v47: array<D9> := new D9[3] [DC19(fm37(f13, false, |v1|, p0, globalState)), v46, v46];
				var v48: set<bool> := {v5};
				v47 := new D9[20] [DC19(v1), v46, DC19(v1), v46.(cf25 := fm50(f14, v6, |v48|, !!!false, globalState)), v46, DC19("ruvt"), v46, DC19(v1), v46, if (p0) then v46.(cf25 := v1).(cf25 := v1) else v46, v46, v46, v46, if (f13) then DC19("khiwbj") else v46, DC19("qmqtcd"), v46, v46, DC19(v1), v46.(cf25 := v1), v46];
				var v49: array<bool> := new bool[9] [!f13, v5, fm53(v6, (seq(abs(0x16), i9  => (map[v5 := f13])))[safeIndex(v7.f12, |(seq(abs(0x16), i9  => (map[v5 := f13])))|) := (map[f13 := v5])[p0 := f13]], globalState) <= v48, if (p0) then f4 else fm0(v8, globalState), v31 <= v31, f13, f4, f13, false];
				v49[safeIndex(379, v49.Length)] := true;
				var v50: array<D18> := new D18[1];
				var v51 := DC26(v5);
				var v52: map<D11, bool> := map[v51 := true];
				var v53: seq<map<D11, bool>> := [v52];
				var v54 := DC35(v53);
				v50[safeIndex(848, v50.Length)] := v54;
				v49[safeIndex(379, v49.Length)], v50[safeIndex(848, v50.Length)] := (fm37(p0, f13, v6, f13, globalState) + v1) <= v1, v54;
				var v55 := DC28(f6, v7.f12, p0, true, v6);
				f14, v5 := safeModuloInt(f6, if (!v5) then v6 else v6), {f13, true, v55.cf45, f4, false} >= (v48 * v48);
				var v56 := new C1(f14, 0x1a9, if (f13) then f5 else f5, f14, !p0);
		}
		
		r0 := match DC24(p0, !f4, v7.f12).(cf38 := |v1|) {
			case DC24(cf36, cf37, cf38) => 'w'
			case DC25(cf39, cf40) => v8
			case DC26(cf41) => v8
			case DC23(cf35) => v8
		};
		var v57: array<D1> := new D1[19];
		r1 := v57;
	}
}

class C6 extends T0, T1 {
	const f11 : bool
	constructor (f11 : bool, f4 : bool, f5 : array<array<bool>>, f6 : int) {
		this.f11 := f11;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
	}
	
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> {
		multiset([f6]) * multiset{f6}
	}
	function fm7(globalState: GlobalState): int {
		(f6 * |multiset{|map[f11 := f6]|, f6}|) * safeDivisionInt(f6, |map[!false := "hkuc"]|)
	}
	function fm8(p0: int, p1: bool, p2: seq<char>, p3: seq<int>, globalState: GlobalState): bool {
		!(f11 && (f6 < |multiset{f4, f4}|))
	}
	function fm9(p0: bool, p1: bool, p2: int, globalState: GlobalState): int {
		f6
	}
	function fm15(globalState: GlobalState): int {
		match DC1() {
			case DC1() => -f6
			case DC2(cf1) => safeDivisionInt(cf1, |[|[f11]|]|)
			case DC0(cf0) => |([f4, f11] + DC8([false]).cf7)|
		}
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) {
		var v0 := "gteox";
		var v1: map<bool, string> := map[!f11 := v0];
		var v2 := DC6();
		var v3 := DC3(f11);
		var v4: array<map<bool, string>> := new map<bool, string>[17] [map[!f11 := v0], map[f11 := v0], v1, map[f11 := (fm16(f4, f4, f11, v2, globalState))[safeIndex(f6, |fm16(f4, f4, f11, v2, globalState)|) := 'l']], v1, v1, v1, v1, if (v3.cf2) then v1 else map[f11 := "jwkwygjor"], v1 + v1, v1 + v1, v1, v1, v1, v1, v1, v1];
		v4[safeIndex(480, v4.Length)] := v1 + v1;
		v4[safeIndex(480, v4.Length)] := v1;
		for i0 := p0 to fm7(globalState) {
			var v5: array<int> := new int[23](i1 => i1 - i0);
			var v6: map<int, bool> := map[f6 := f4];
			var v7: map<array<int>, map<int, bool>> := map[v5 := v6];
			var v8: map<bool, int> := map[f4 := i0];
			v7 := v7[v5 := map[if (f11 in v8) then v8[f11] else f6 := f4]];
			var v9: multiset<int> := multiset{i0};
			if (v9 !! v9) {
				var v10: array<seq<array<int>>> := new seq<array<int>>[25];
				var v11: seq<array<int>> := [v5];
				v10[safeIndex(952, v10.Length)] := v11;
				v10[safeIndex(952, v10.Length)] := v11;
				var v12 := 85;
				var v13 := false;
				v12, v13 := p0 - -f6, !f4;
				var v14: array<D0> := new D0[16](i2 => DC0(v12));
				v14 := v14;
				var v15: set<int> := {-0x372};
				v12 := |v15| + p0;
				var v16: map<bool, array<int>> := map[f11 := v5];
				var v17: map<map<bool, array<int>>, int> := map[v16 := if (f4) then if (false in v8) then v8[false] else fm9(!v13, f11, p0, globalState) else i0];
				v17 := v17[v16[f4 := v5] := i0];
			} else {
				var v18 := new C0(i0);
				var v19 := 0x2f7;
				var v20: map<int, int> := map[p0 := 0x19a];
				var v21: map<int, map<int, int>> := map[p0 := v20];
				v19 := |v21|;
				v19 := v18.f12 * -v19;
				var v22 := true;
				v5[safeIndex(426, v5.Length)] := fm3(v19, v22, globalState) + v19;
				v5[safeIndex(914, v5.Length)] := if (f4 in v8) then v8[f4] else i0;
				var v23 := DC7(v18.f12);
				v22, v5[safeIndex(426, v5.Length)], v5[safeIndex(914, v5.Length)], v19, v22 := v22, -(v18.f12 - (|fm16(f4, v22, f11, v2, globalState)| + p0)), p0 * v23.cf6, f6, v3.cf2;
				v19 := |v9|;
			}
			
			var v24: set<bool> := {f4};
			var v25: map<set<bool>, bool> := map[v24 := f11];
			var v26: set<int> := {|v25|, p0, p0, p0};
			v26 := v26;
			var v27 := -0x10f;
			v27 := v27;
		}
		var i3 := 0;
		while (f11)
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v28: array<int> := new int[12];
			v28 := v28;
			var v29 := false;
			var v30: seq<int> := [0x247];
			v29 := v30 == v30;
			v29 := f11;
			var v31 := -0x3d5;
			v29, v31 := f11, 759;
		}
		var v32: multiset<bool> := multiset{f11};
		var v33: map<int, bool> := map[-f6 := f4];
		if ((f6 + f6) == (p0 * (if (f4 in v32) then v32[f4] else |v33|))) {
			var v34: seq<int> := [0x147, p0];
			var v35 := new C0(|v34|);
			var v36: seq<string> := ["wfhdckno"];
			var v37 := new C0(|v36|);
			var v38: array<string> := new string[28](i4 => (v0 + v0)[safeIndex(v34[safeIndex(0xef, |v34|)], |(v0 + v0)|) := 'j']);
			v38[safeIndex(734, v38.Length)] := v0;
			var v39 := 's';
			v38[safeIndex(734, v38.Length)] := v0 + (v0 + v0[safeIndex(f6, |v0|) := v39]);
			var v40: set<int> := {p0};
			var v41: array<int> := new int[13];
			var v42: map<int, array<int>> := map[|(v40 * v40)| := v41];
			v42 := v42[-824 := v41];
			var v43: map<string, int> := map[v0 := v37.f12];
			var v44: array<bool> := new bool[6] [f4, f11, f11, f4, f4, f11];
			var v45: map<array<bool>, bool> := map[v44 := f11];
			v43 := v43[seq(abs(0x35f), i5  => (v39)) := |fm18(v0, |v45|, globalState)|];
		} else {
			var v46 := 0x25a;
			var v47 := 'j';
			v46, v46, v47 := p0, p0, v47;
			if (f4) {
				v46 := p0;
				v46 := if (v46 > p0) then |fm16(f11, f4, DC9(f4, v46, f11).cf10, v2, globalState)| else v46;
				var v48: set<int> := {0x112, p0, |v0|};
				v48 := v48 + {p0};
				var v49: array<int> := new int[11](i6 => i6 * p0);
				v49[safeIndex(996, v49.Length)] := if (f4) then f6 else p0;
				v49[safeIndex(996, v49.Length)] := 0x15c;
				var v50 := new C0(if (f11) then -0x10c else p0);
			} else {
				var v51 := DC12(v47);
				v47 := v51.cf13;
				var v52: seq<int> := [|fm16(f11, fm0(v47, globalState), f11, v2, globalState)| * fm15(globalState)];
				var v53: seq<bool> := [f4];
				var v54: array<int> := new int[29];
				var v55: array<seq<bool>> := new seq<bool>[19] [v53, [f4], fm19(f11, fm3(-|map[v47 := v54]|, f11, globalState), globalState), v53, v53, v53, v53, v53, if (f4) then v53 else v53, v53, v53, (v53[safeIndex(v46, |v53|) := true])[safeIndex(f6, |v53[safeIndex(v46, |v53|) := true]|) := false], v53 + v53, v53, v53, v53, v53, v53, v53];
				v55[safeIndex(737, v55.Length)] := DC13(f6, f11, v46, v53).cf17 + v53;
				v52, v46, v46, v55[safeIndex(737, v55.Length)] := v52, p0, -0x31a, v53;
				var v56 := DC10(p0);
				var v57 := DC11(v56);
				var v58: array<bool> := new bool[6] [f11, f4, f4, f4, f4, f4];
				var v59: multiset<array<bool>> := multiset{v58};
				var v60: map<D3, multiset<array<bool>>> := map[v57 := v59 - v59];
				v60 := v60[v57 := v59];
				var v61: multiset<int> := multiset{if (f4) then -0x85 else p0, v46, |(seq(abs(-0x150), i7  => ('j')))|, p0};
				v61 := DC14(multiset(v52)).cf18;
				var v62: map<int, int> := map[p0 := p0];
				var v63: map<int, string> := map[|v61| := v0];
				var v64: seq<map<int, string>> := [v63, v63];
				v54[safeIndex(300, v54.Length)] := if (f4) then fm3(if (v46 in v62) then v62[v46] else v46, f4, globalState) else |v64[safeIndex(|v0|, |v64|)]|;
				v54[safeIndex(300, v54.Length)] := if (f11) then p0 else |DC8((fm20(f6, f4, globalState)).cf7).cf7|;
			}
			
			var v65: array<bool> := new bool[8](i8 => multiset{v46, |v0|} < multiset{p0, v46, v46});
			v65 := v65;
			v33 := v33[0x1c3 := f4];
			var v66 := true;
			var v67: map<bool, int> := map[f4 := f6];
			v66 := |v67| in multiset{f6};
		}
		
		var v68 := true;
		v68 := f4;
		var v69 := 0x2a6;
		v69 := -(if (!(v68 <==> v68)) then f6 else p0);
		var v70: map<bool, int> := map[v68 := p0];
		r0 := v70;
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) {
		var v0: seq<bool> := [f11, p0, f4];
		var v1: seq<bool> := [v0[safeIndex(f6, |v0|)], p0, f4];
		var v2 := DC8(v0[safeIndex(f6, |v0|) := f11]);
		match (v2) {
			case DC9(cf8, cf9, cf10) =>
				var v3: array<array<int>> := new array<int>[9];
				var v4: array<bool> := new bool[18](i0 => !f11);
				v4[safeIndex(874, v4.Length)] := f11;
				var v5: seq<array<array<int>>> := [v3, v3, v3, v3, v3];
				var v6: multiset<array<bool>> := multiset{v4, v4};
				var v7 := "rmjgkq";
				cf9, v3, cf9, cf9, v4[safeIndex(874, v4.Length)] := fm9(cf10, f11, cf9, globalState), v5[safeIndex(if (v4 in v6) then v6[v4] else |v7|, |v5|)], -(|v7| + safeModuloInt(fm9(true, f11, cf9, globalState), f6)), if (p0) then |v7| else cf9, cf10;
				if (f4) {
					cf10 := false;
					f5[safeIndex(333, f5.Length)] := v4;
					f5[safeIndex(333, f5.Length)] := DC5(v4).cf5;
					var v8 := DC10(|v7|);
					var v9: array<D3> := new D3[25] [if (cf10) then v8 else v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, v8, DC10(cf9), v8, v8, DC10(f6), v8, DC10(f6), v8, if (cf10) then v8 else v8, v8, v8, v8, v8, v8, v8];
					v9[safeIndex(515, v9.Length)] := v8;
					v9[safeIndex(515, v9.Length)] := v8;
					cf9 := safeModuloInt(cf9, f6);
					var v10: map<bool, bool> := map[true := v4[safeIndex(874, v4.Length)]];
					var v11: seq<map<bool, bool>> := [v10[cf10 := cf8]];
					v10 := v11[safeIndex(f6, |v11|)] + v10;
				} else {
					var v12: map<bool, int> := map[!v1[safeIndex(cf9, |v1|)] := f6];
					var v13 := new C0(if (cf10 in v12) then v12[cf10] else f6);
					var v14 := DC7(safeModuloInt(f6, 0x1b1));
					var v15: set<int> := {v13.f12};
					var v16: multiset<set<int>> := multiset{v15, v15, v15};
					var v17: map<C0, array<bool>> := map[v13 := v4];
					var v18: map<int, int> := map[0x178 := |v17|];
					v14 := if ({f6} !in v16[v15 := abs(v13.f12)]) then DC7(|v18|) else v14;
					var v19 := 's';
					r0 := v19;
					cf10 := f4;
					cf9 := cf9;
				}
				
				var v20 := new C0(safeDivisionInt(|v0|, f6));
				v4[safeIndex(874, v4.Length)] := v20.fm17(p0, globalState);
			case DC10(cf11) =>
				var v21: map<int, bool> := map[safeModuloInt(0x89, cf11) := if (f11) then f11 else p0];
				v21 := v21[cf11 := f11];
				var v22 := false;
				var v23: array<int> := new int[26];
				var v24: seq<int> := [cf11];
				v22, cf11, v22, v23 := p0, cf11, v24 != v24, v23;
				var v25 := "gan";
				var v26 := 'j';
				v25 := if (if (f4) then fm0(v26, globalState) else !f4) then v25 else v25 + v25;
				var v27: array<array<D0>> := new array<D0>[24];
				v27 := v27;
			case DC8(cf7) =>
				var v28: array<int> := new int[7];
				var v29 := DC4(-f6, f6);
				var v30: seq<int> := [v29.cf3];
				v28[safeIndex(538, v28.Length)] := v30[safeIndex(f6, |v30|)];
				var v31: map<bool, int> := map[f4 := 476];
				var v32: multiset<map<bool, int>> := multiset{v31, v31};
				var v33 := 'v';
				var v34: seq<char> := [v33];
				var v35: multiset<bool> := multiset{p0, fm8(0x2d, f11, v34, v30, globalState), p0, f11, f4};
				v28[safeIndex(538, v28.Length)] := |(fm21(--f6, v32, 0xc1, f6, globalState) - v35)|;
				var v36 := new C0(-(v28[safeIndex(538, v28.Length)] - v28[safeIndex(538, v28.Length)]));
				var v37: array<bool> := new bool[27];
				v37[safeIndex(189, v37.Length)] := f4 || f11;
				v37[safeIndex(189, v37.Length)] := !v0[safeIndex(0xaf + v28[safeIndex(538, v28.Length)], |v0|)];
				var v38 := DC13(v28[safeIndex(538, v28.Length)], f11, |v34|, v1);
				var v39: set<bool> := {v38.cf15, p0, f4};
				v37[safeIndex(189, v37.Length)], v28[safeIndex(538, v28.Length)], v28[safeIndex(538, v28.Length)] := v1[safeIndex(|(v39 + v39)|, |v1|)], if (if (p0) then f4 else p0) then f6 else v36.f12, v30[safeIndex(f6, |v30|)];
			case DC11(cf12) =>
				var v40: array<int> := new int[19](i1 => safeModuloInt(i1, f6));
				v40[safeIndex(30, v40.Length)] := 0x101;
				v40[safeIndex(30, v40.Length)] := f6;
				v40[safeIndex(30, v40.Length)] := f6;
				var v41 := false;
				var v42: array<bool> := new bool[26];
				var v43: multiset<bool> := multiset{p0, !v41, f4};
				var v44: map<bool, multiset<bool>> := map[f11 := v43];
				var v45: multiset<multiset<bool>> := multiset{v43, multiset(v1), if (v41 in v44) then v44[v41] else v43, v43};
				v42[safeIndex(420, v42.Length)] := v45 >= v45;
				v41, v40[safeIndex(30, v40.Length)], v42[safeIndex(420, v42.Length)], v40[safeIndex(30, v40.Length)] := true, v40[safeIndex(30, v40.Length)], f11, safeModuloInt(v40[safeIndex(30, v40.Length)], if (f11) then f6 else v40[safeIndex(30, v40.Length)]);
				if (true) {
					var v46 := DC2(|[{f6, f6}]|);
					var v47: multiset<int> := multiset{v40[safeIndex(30, v40.Length)], v46.cf1};
					var v48 := DC9(f11, |v47|, f11);
					var v49: seq<int> := [v48.cf9, v40[safeIndex(30, v40.Length)]];
					var v50, v51 := m0(fm22(globalState), p0 <==> f4, multiset(v49) * v47, v41, globalState);
					var v52: map<int, int> := map[v40[safeIndex(30, v40.Length)] + v51 := 0xc8];
					v52 := v52[fm3(v51, v41, globalState) := v49[safeIndex(-562, |v49|)]];
					var v53 := "owcvyfr";
					var v54: map<int, seq<bool>> := map[|v53| * v51 := v0];
					v54 := v54[v51 := [v42[safeIndex(420, v42.Length)], f11, f11, p0]];
					var v55 := 'd';
					v42[safeIndex(420, v42.Length)] := fm0(v55, globalState);
					var v56: set<bool> := {v42[safeIndex(420, v42.Length)]};
					var v57: seq<array<int>> := [v40, v40];
					v56, v57, v50 := fm23(v51, f6, globalState), v57, v41;
				} else {
					v41 := !f11;
					var v58: map<int, bool> := map[|v43| := true];
					var v59: multiset<map<int, bool>> := multiset{v58};
					var v60: seq<map<int, bool>> := [v58, v58];
					var v61 := "wgpjpb";
					var v62: map<int, multiset<map<int, bool>>> := map[f6 := v59 + multiset{v58, v58, v58, v60[safeIndex(|v61|, |v60|)], v58}];
					v62 := v62[v40[safeIndex(30, v40.Length)] := v59];
					v40[safeIndex(30, v40.Length)] := f6 - f6;
					v42[safeIndex(420, v42.Length)] := 0x323 >= (fm9(p0, p0, fm15(globalState), globalState) * (if (p0 in v43) then v43[p0] else f6));
					var v63 := 'e';
					r0 := v63;
				}
				
		}
		
		var i2 := 0;
		while (f4)
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v64: array<bool> := new bool[21](i3 => p0);
			v64 := v64;
			var v65 := true;
			v65 := f4;
			var v66: multiset<int> := multiset{126, f6};
			var v67 := DC2(|v66|);
			var v68: map<bool, int> := map[f11 := v67.cf1];
			v68 := v68[f4 := f6 + f6];
			v64[safeIndex(910, v64.Length)] := !p0;
			v64[safeIndex(910, v64.Length)] := p0;
		}
		var v69: map<int, int> := map[f6 := f6];
		v69 := v69 + (v69 + v69);
		v0 := v0 + v1;
		for i4 := f6 to safeModuloInt(f6, f6) {
			var v70 := 0x241;
			v70 := f6;
			var v71 := DC7(f6);
			match (v71) {
				case DC6() =>
					var v72, v73, v74 := m8(safeModuloInt(f6, v70), v70, globalState);
					var v75 := 'j';
					var v76 := "ftxvovkv";
					var v77 := DC0(|v76|);
					var v78: map<D0, bool> := map[v77 := v72];
					var v79: map<bool, map<D0, bool>> := map[fm0(v75, globalState) := v78];
					var v80: array<int> := new int[5] [v70, -|(if (!p0 in v79) then v79[!p0] else map[DC0(v70) := f4])|, f6, -352, f6];
					var v81: set<array<int>> := {v80};
					v72 := v81 !! v81;
					v69 := v69[f6 := f6];
					var v82 := DC2(v70);
					v82 := v82;
				case DC7(cf6) =>
					var v83: array<bool> := new bool[27](i5 => p0);
					var v84 := 'w';
					var v85: map<char, int> := map[v84 := i4];
					var v87: seq<char> := [v84, v84];
					var v88: seq<map<char, int>> := [map v86 : char | v86 in v87 :: (v86) := (v70)];
					var v89: map<seq<array<bool>>, seq<map<char, int>>> := map[[v83] := [v85, v85] + v88[safeIndex(cf6, |v88|) := map[fm1(fm3(cf6, f11, globalState), f11, globalState) := -f6]]];
					var v90: seq<array<bool>> := [v83, v83, v83];
					var v91: seq<array<bool>> := [v83, v83, v90[safeIndex(-553, |v90|)], DC5(v83).cf5, v83];
					v89 := v89[v90 := seq(abs(707), i6  => (v85))];
					var v92: map<string, int> := map[v87 := f6];
					v92 := v92["qpjli" := cf6];
					var v93: map<bool, bool> := map[p0 := p0];
					r0 := fm1(|(v93 + map[false := false])|, f4, globalState);
					var v94: array<array<int>> := new array<int>[4];
					var v95: array<int> := new int[1](i7 => i7 - 964);
					v94[safeIndex(906, v94.Length)] := v95;
					v94[safeIndex(906, v94.Length)] := v95;
				case DC5(cf5) =>
					var v96 := new C0(f6);
					var v97: multiset<int> := multiset{v96.f12, v96.f12, f6, v70};
					var v98: map<multiset<int>, string> := map[v97 := "synxuae"];
					var v99, v100 := m0(v98, f11, multiset{v70}, f11, globalState);
					var v101 := "adljl";
					var v102: set<string> := {if (v0[safeIndex(v70, |v0|)]) then v101 else v101};
					v102 := v102;
					var v103: map<bool, int> := map[!!p0 := -v70];
					var v104: seq<map<bool, int>> := [map[v99 := v96.f12], v103, v103, v103];
					var v105 := new C0(|v104|);
			}
			
			var v106 := "tasx";
			var v107: array<int> := new int[17];
			v107[safeIndex(289, v107.Length)] := 415;
			v106, v107[safeIndex(289, v107.Length)] := "xohxe", |v106|;
			v107[safeIndex(289, v107.Length)] := v70;
		}
		var v108 := 'h';
		if (fm0(v108, globalState)) {
			var v109 := 0x27f;
			var v110: multiset<bool> := multiset{f4};
			v109 := f6 * |[v109, |v110|, f6]|;
			var v111 := false;
			var v112 := "s";
			var v113: map<int, bool> := map[|v112| := v111];
			v111 := if (v109 in v113) then v113[v109] else f11;
			var v114: map<seq<bool>, bool> := map[v0 := v111];
			v114 := v114[[p0] := p0];
			var v115: multiset<seq<bool>> := multiset{fm19(p0, v109, globalState)};
			var v116 := DC3(v111);
			var v117: seq<seq<bool>> := [v0, v0];
			var v118: array<multiset<seq<bool>>> := new multiset<seq<bool>>[22] [v115, multiset{v1, v0, [f11], v0}, v115, v115, multiset{v1}, v115, v115, multiset{v1, [f11, true], v0, v0}, if (v116.cf2) then multiset(seq(abs(0x92), i8  => (v0))) else fm24(v109, v108, v109, globalState), v115, v115 * multiset(v117), v115, v115 * v115, v115, v115 * v115, multiset{[f4]} * v115, v115, v115, v115, v115, v115, v115];
			var v119: map<bool, seq<seq<bool>>> := map[v111 := [v1]];
			var v120 := DC15(v115);
			v118[safeIndex(270, v118.Length)] := multiset(if (p0 in v119) then v119[p0] else v117) + v120.cf19;
			var v121: seq<int> := [v109, -v109];
			var v122: multiset<seq<int>> := multiset{v121};
			var v123: array<seq<bool>> := new seq<bool>[1];
			v118[safeIndex(270, v118.Length)], v122, v123, v111 := v115[[false] := abs(f6)], v122, v123, f11;
			v109 := safeDivisionInt(313, v109);
		} else {
			var v124 := DC10(f6);
			v124 := DC10(f6);
			var v125 := 0xd8;
			v125 := if (p0) then if (-103 in v69) then v69[-103] else v125 else |(v0[safeIndex(v125, |v0|) := fm0(v108, globalState)])[safeIndex(0x3b4, |v0[safeIndex(v125, |v0|) := fm0(v108, globalState)]|) := true]|;
			match (DC8(v0)) {
				case DC9(cf8, cf9, cf10) =>
					var v126: multiset<int> := multiset{|(seq(abs(0x2a4), i9  => (v108)))|, |{f11, !f4, cf8}|, |fm25(cf9, f6, f4, f11, globalState)|, v125};
					var v127: seq<multiset<int>> := [v126];
					var v128: seq<int> := [cf9, f6, v125];
					v126 := v127[safeIndex(|(map[v125 := v128])[f6 := seq(abs(0x327), i10  => (cf9))]|, |v127|)];
					var v129: set<int> := {v125};
					cf9 := fm9(f11, cf8, |DC17(v129).cf23|, globalState);
					var v130: array<int> := new int[8](i11 => i11 * f6);
					var v131: array<D0> := new D0[10](i12 => DC0(f6));
					var v132 := DC0(cf9);
					v131[safeIndex(258, v131.Length)] := v132;
					var v133: array<bool> := new bool[27];
					v124, v130, cf9, v131[safeIndex(258, v131.Length)], v133 := v124, v130, --(-safeDivisionInt(0x215, cf9) + (if (cf10) then cf9 else -614)), v132, v133;
					v125 := --cf9;
				case DC10(cf11) =>
					var v134: array<int> := new int[19](i13 => i13 + f6);
					var v135: map<array<int>, bool> := map[v134 := p0];
					v135 := DC18((v135[v134 := f4])[v134 := f4]).cf24;
					var v136 := true;
					cf11, cf11, v136 := f6, f6 + cf11, DC8(v0).cf7 != [p0, p0, f11, v136];
					var v137 := "fjydewlp";
					var v138: T1 := new C3(f5, |v137|, false);
					var v139: map<bool, T1> := map[v136 := v138];
					var v140: map<int, T1> := map[f6 := v138];
					v139 := v139[if (p0) then f11 else v138.f4 := if (v138.f6 in v140) then v140[v138.f6] else v138];
					v125 := v125;
				case DC8(cf7) =>
					var v141: array<bool> := new bool[11];
					var v142: seq<array<bool>> := [v141];
					var v143 := "sxv";
					var v144 := DC13(-v125, v141 !in v142, safeModuloInt(-|v143|, f6), v0);
					var v145 := false;
					var v146: seq<int> := [f6, v125];
					v125, v144, v125, v145 := f6 + f6, fm63(globalState), v146[safeIndex(v125, |v146|)], f4;
					var v147: multiset<bool> := multiset{f11};
					v145 := !f11 <==> (v147 > v147);
					var v148: map<char, string> := map[v108 := v143 + v143];
					v148 := v148;
					var v150: array<multiset<set<int>>> := new multiset<set<int>>[11](i14 => multiset{set v149 : int | (0x210 <= v149) && (v149 < 26) :: (safeDivisionInt(v149, v125)), {f6}});
					var v151: map<int, bool> := map[-0x231 := DC9(v145, fm15(globalState), f11).cf8];
					var v152: set<int> := {|v151|};
					var v153: multiset<set<int>> := multiset{v152, v152};
					v150[safeIndex(763, v150.Length)] := v153;
					var v154: seq<set<int>> := [v152, v152, v152];
					v150[safeIndex(763, v150.Length)] := v153[v154[safeIndex(|v143|, |v154|)] - v152 := abs(-safeModuloInt(|v146|, v125))];
				case DC11(cf12) =>
					v125 := fm15(globalState);
					var v155: map<bool, int> := map[f11 := 0x2d5];
					var v156: multiset<map<bool, int>> := multiset{v155, fm44(globalState), v155, v155, v155};
					v156 := v156 + (v156 + (multiset{v155})[v155 := abs(fm3(v125, p0, globalState))]);
					v108 := v108;
					v155 := v155;
			}
			
			v125 := if (p0) then -f6 else f6 + f6;
			var v157 := true;
			v157 := p0;
		}
		
		var v158 := "ehhyianjb";
		var v159: C2 := new C2(v158, p0);
		var v160: map<int, C2> := map[0x2b4 := v159];
		r0 := fm1(fm3(f6, p0, globalState), fm7(globalState) <= |v160|, globalState);
		var v161 := DC3(f4);
		var v162: map<bool, bool> := map[f11 := p0];
		var v163: array<D1> := new D1[29] [v161, v161, v161, fm55(if (p0 in v162) then v162[p0] else f11, f4, f6, globalState), v161, v161, v161, v161, v161, v161, v161, v161, v161, DC3(fm0(v108, globalState)), v161, v161, DC3(f11), v161, v161, v161, v161, v161, v161, v161, v161, v161, v161, v161, v161];
		var v164: map<bool, array<D1>> := map[p0 := v163];
		var v165: map<bool, int> := map[p0 := f6];
		var v166: map<map<bool, int>, bool> := map[v165 := f4];
		r1 := if ((if (map[p0 := f6] in v166) then v166[map[p0 := f6]] else f4) in v164) then v164[if (map[p0 := f6] in v166) then v166[map[p0 := f6]] else f4] else v163;
	}
	method m7(globalState: GlobalState) {
		var v0 := true;
		v0 := f11;
		var v1 := 0x34b;
		var v2: seq<int> := [f6];
		var v3 := "gnnf";
		var v4: seq<bool> := [true];
		var v5: seq<seq<bool>> := [v4];
		var v6 := DC27(v5);
		var v7: seq<seq<int>> := [v2, fm39(v3, v6, f11, globalState) + [v1, f6]];
		v1 := |v7[safeIndex(if (f4) then f6 else f6, |v7|)]|;
		var v8: array<bool> := new bool[14];
		v8 := v8;
		var v9 := new C4(f5, v1, f4);
		v1, v0 := safeDivisionInt(v1, f6), |(if (v0) then v4 else v4[safeIndex(f6, |v4|) := f11])| != v1;
		var v10: map<int, int> := map[f6 := v1];
		var v11 := new C1(v1, (if (v1 in v10) then v10[v1] else v1) * |v3|, f5, v1, !v0);
	}
	method m8(p0: int, p1: int, globalState: GlobalState) returns (r0: bool, r1: bool, r2: D1) {
		var v0 := 'a';
		var v1: array<bool> := new bool[15];
		var v2: map<char, array<bool>> := map[v0 := v1];
		var v3 := DC39(DC37(v2));
		var v4 := DC37(v2[v0 := v1]);
		v3 := DC39(v4);
		if (f11) {
			var v5 := 594;
			v5 := p1 + fm3(f6, f11, globalState);
			var v6: set<bool> := {!f4};
			v6 := v6;
			r1 := f11 <==> (if (f4) then f11 else f11);
			v5 := v5 + f6;
			r0 := f4;
		} else {
			f5[safeIndex(951, f5.Length)] := v1;
			f5[safeIndex(951, f5.Length)] := v1;
			var v7: seq<bool> := [f11, f4, true, f4];
			var v8: seq<seq<bool>> := [v7 + v7, v7 + v7, [f11] + v7];
			v8 := seq(abs(0x6c), i0  => (v7));
			var v9: multiset<string> := multiset{"brp"};
			var v10: multiset<int> := multiset{255, if ("cufybjlks" in v9) then v9["cufybjlks"] else p0};
			var v11 := DC0(f6);
			var v12: map<int, char> := map[p0 := v0];
			var v13: map<bool, bool> := map[f4 := f4];
			var v14, v15 := m0(map[v10 := seq(abs(0x241), i1  => (v0))], f4, fm6(v11, v11, v12, p1, globalState) + v10, false ==> (if (f11 in v13) then v13[f11] else f4), globalState);
			var v16 := "wwdadof";
			var v17: map<int, string> := map[v15 := v16];
			var v18 := DC6();
			v17 := v17[safeDivisionInt(p1, 0x1ba) := (fm16(v14, f11, false, v18, globalState))[safeIndex(p1, |fm16(v14, f11, false, v18, globalState)|) := v0] + v16];
			var v19: array<array<bool>> := new array<bool>[2] [f5[safeIndex(951, f5.Length)], v1];
			var v20 := DC31(v19);
			var v21 := new C5(f4, -p0, v20.cf50, |v16| * p0, v14);
		}
		
		var v22 := 0x35d;
		v22 := v22 * (p1 - p1);
		if (f11) {
			var v23 := "senka";
			var v24: map<array<array<bool>>, bool> := map[f5 := v23 == v23];
			v24 := v24[f5 := f4];
			var v25: set<bool> := {f4, !f11, f11};
			v25 := v25 * v25;
			r0 := f4;
			m7(globalState);
			v0 := v0;
		} else {
			var v26 := DC3(!true);
			r2 := v26;
			var v27: map<bool, bool> := map[!f4 := f4];
			var v28: map<int, bool> := map[p1 := if (f4 in v27) then v27[f4] else !f11];
			v28 := v28[v22 := f11];
			r0 := f11;
			v1[safeIndex(250, v1.Length)] := f4;
			v1[safeIndex(250, v1.Length)] := f11;
			var v29 := "pgj";
			var v30 := new C2(v29, v1[safeIndex(250, v1.Length)]);
		}
		
		var v31 := "a";
		var v32: T0 := new C2(v31, f11);
		v32 := v32;
		var v33: map<bool, bool> := map[if (f11) then f11 else f4 := f11];
		v33 := v33;
		r0 := f4;
		r1 := f6 >= (f6 + v22);
		var v34 := DC3(f11);
		r2 := v34;
	}
}

class C7 extends T1 {
	const f9 : map<bool, string>
	const f10 : array<bool>
	constructor (f9 : map<bool, string>, f10 : array<bool>, f5 : array<array<bool>>, f6 : int, f4 : bool) {
		this.f9 := f9;
		this.f10 := f10;
		this.f5 := f5;
		this.f6 := f6;
		this.f4 := f4;
	}
	
	function fm8(p0: int, p1: bool, p2: seq<char>, p3: seq<int>, globalState: GlobalState): bool {
		if (if (f4) then f4 else f4) then (seq(abs(0x31f), i0  => ('i'))) <= ['h'] else f6 == -f6
	}
	function fm9(p0: bool, p1: bool, p2: int, globalState: GlobalState): int {
		f6
	}
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> {
		(if (false) then multiset{f6} else multiset{f6}) - (multiset{f6} - multiset{0x1d8})
	}
	function fm7(globalState: GlobalState): int {
		-safeDivisionInt(f6, f6) - 0x2a5
	}
	function fm13(p0: int, p1: bool, p2: bool, globalState: GlobalState): seq<bool> {
		[f4] + [!f4, f4, f4]
	}
	function fm14(p0: map<int, bool>, p1: bool, p2: int, p3: bool, globalState: GlobalState): bool {
		f6 == 0x298
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) {
		var v0 := new C1(f6, p0, f5, p0, f4);
		var v1: seq<bool> := [f4, f4, true, f4, !f4];
		match (DC13(0xc2, !f4, f6 - v0.f16, v1)) {
			case DC13(cf14, cf15, cf16, cf17) =>
				var v2: set<int> := {-|fm64(p0, f6, cf15, globalState)|};
				var v3 := DC20(cf15, 'n', v0.f16);
				var v4 := new C6(|v2| < 272, f4, f5, v3.cf28);
				cf15 := true;
				cf15 := !!v4.f11;
				v0.f16 := v0.f15 - -0x391;
			case DC12(cf13) =>
				var v5: array<int> := new int[2];
				v5[safeIndex(382, v5.Length)] := |v1|;
				v5[safeIndex(382, v5.Length)] := v0.f15;
				var v6: array<multiset<int>> := new multiset<int>[23](i0 => multiset{v0.f16} + multiset{p0});
				var v7: multiset<int> := multiset{p0};
				v6[safeIndex(126, v6.Length)] := if (f4) then v7 else v7;
				var v8 := true;
				var v9: map<int, bool> := map[v0.f15 := true];
				var v10: C3 := new C3(f5, v0.f15, if (fm14(v9, true, v0.f16, false, globalState)) then f4 else v8);
				var v11: map<bool, bool> := map[v8 := v8];
				var v12: multiset<bool> := multiset{v8, if (f4 in v11) then v11[f4] else v8};
				var v13 := "sdatbjvc";
				var v14: C2 := new C2("xyw", f4);
				var v15: map<int, C2> := map[v0.f16 := v14];
				var v16 := DC33(false, f6, v12, |v13|, fm23(|v15|, v0.f16, globalState));
				var v17 := DC20(v16.cf52, cf13, v0.f16);
				var v18: array<char> := new char[6] [cf13, cf13, (v17.(cf28 := |v11|, cf26 := f4)).cf27, if (f4) then cf13 else cf13, cf13, cf13];
				v6[safeIndex(126, v6.Length)], v8, v10, v18 := v7, (f4 || true) && (v0.f15 != v0.f16), v10, v18;
				v0.f16 := v0.f16;
				v0.f16, v0.f16, v5[safeIndex(382, v5.Length)] := safeDivisionInt(0x3dd, v0.f15) - |"cfu"|, 0xc4, -v0.f16;
		}
		
		var v19: map<int, int> := map[p0 := v0.f16];
		var v20: map<bool, map<int, int>> := map[true := v19];
		var v21: seq<map<int, int>> := [v19, if (fm0('a', globalState) in v20) then v20[fm0('a', globalState)] else v19];
		var v22: multiset<int> := multiset{|v21[safeIndex(-p0, |v21|)]|, v0.f16, p0};
		var v23: map<multiset<int>, string> := map[v22 := "rcxfakrum"];
		var v24, v25 := m0(v23, f4, v22, f4, globalState);
		f10[safeIndex(270, f10.Length)] := v24;
		var v26 := DC26(!f4);
		var v27: set<D11> := {v26, v26, v26, v26};
		f10[safeIndex(270, f10.Length)] := v25 != |(v27 * v27)|;
		var v28 := "lvmi";
		var v29: map<bool, bool> := map[false := f10[safeIndex(270, f10.Length)]];
		var v30: seq<int> := [f6, |v1|];
		var v31: seq<int> := [v0.f16, |v30|];
		var v32: array<int> := new int[14] [fm7(globalState), v0.f15, safeModuloInt(|v28|, v0.f15), |v29|, f6, safeModuloInt(v25, |v1|), f6 * p0, safeDivisionInt(0x22, 0x247), safeDivisionInt(|v31|, 967), v0.f16, v0.f15, v0.f16, v25, v0.f15];
		var v33: map<int, bool> := map[0x1e := f10[safeIndex(270, f10.Length)]];
		var v34: array<bool> := new bool[25];
		var v35 := DC22("tsfhsrurn", v24, f4, v34, f6);
		var v36: multiset<bool> := multiset{f4, v24};
		var v37: C4 := new C4(f5, v25, v24);
		var v38: map<C4, string> := map[v37 := v28];
		var v39: set<bool> := {f4, v24};
		var v40 := DC33(if (v0.f16 in v33) then v33[v0.f16] else f10[safeIndex(270, f10.Length)], v37.fm7(globalState), multiset{f10[safeIndex(270, f10.Length)]}, |v22|, v39);
		var v41 := DC33(fm14(v33, f10[safeIndex(270, f10.Length)], v0.f15, v35.cf32, globalState), -((if (v0.f16 in v19) then v19[v0.f16] else |v22|) + 0x1af), v36[v24 := abs(|v38|)] * v36, -(v0.f16 * v0.f16), {v24, v24, true, !v24, !v40.cf52});
		v0.f16, v25, v32, v40, f10[safeIndex(270, f10.Length)] := 0x9a, 0x34a, v32, v40, f10[safeIndex(270, f10.Length)];
		v0.f16, f10[safeIndex(270, f10.Length)] := p0, if (v25 == v0.f16) then false else false;
		var v42: map<bool, int> := map[f4 := p0];
		r0 := v42;
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) {
		var v0 := -0x33e;
		var v1 := "q";
		var v2: C3 := new C3(f5, |v1|, p0);
		var v3: map<C3, bool> := map[v2 := false];
		v0 := safeDivisionInt(|map[v3 := f6]|, v0) * f6;
		var i0 := 0;
		while (p0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v4: array<string> := new string[23];
			v4[safeIndex(251, v4.Length)] := seq(abs(0x242), i1  => ('o'));
			var v5 := false;
			var v6: array<char> := new char[13](i2 => 'd');
			var v7 := 'h';
			v6[safeIndex(294, v6.Length)] := v7;
			var v8: set<int> := {474, f6, f6, -fm7(globalState)};
			var v9: seq<bool> := [!f4];
			var v10: seq<string> := [v1];
			var v11: map<array<bool>, seq<string>> := map[f10 := v10];
			v4[safeIndex(251, v4.Length)], v5, v0, v6[safeIndex(294, v6.Length)], v8 := v1, true in v9, |(if (f10 in v11) then v11[f10] else v10 + v10)|, v7, v8;
			v0 := safeDivisionInt(0x194, |v1|);
			var v12: map<bool, char> := map[p0 := v7];
			var v13: seq<map<bool, char>> := [(fm65(v0, 0x271, 0xd1, globalState))[true := v6[safeIndex(294, v6.Length)]], v12];
			v0 := -|(v12[v5 := v6[safeIndex(294, v6.Length)]] + (v12 + v13[safeIndex(|v1|, |v13|)]))|;
			v5 := p0;
		}
		var v14: T0 := new C6(f4, p0, f5, v0);
		var v15: map<int, T0> := map[f6 := v14];
		var v16: map<int, int> := map[f6 := f6];
		var v17 := DC9(false, |v15| + |v16|, v14.f4);
		match (v17) {
			case DC9(cf8, cf9, cf10) =>
				var v18: set<bool> := {v14.f4, p0, f4};
				var v19: array<int> := new int[7](i3 => safeModuloInt(i3, cf9));
				var v20: map<array<int>, set<bool>> := map[v19 := v18];
				var v21 := 'x';
				var v22: map<bool, bool> := map[true := v14.f4];
				var v23: array<set<bool>> := new set<bool>[25] [v18, {true}, v18, v18, {cf10, p0, v14.f4} - v18, if (cf8) then v18 else v18, v18, v18, {cf8}, v18, v18, v18, {p0, v14.f4, v14.f4, cf10, v14.f4}, v18, v18, v18, v18, v18 - v18, v18, v18, (if (v19 in v20) then v20[v19] else v18) + fm46(v21, globalState), v18, v18 * v18, v18, v18 + fm53(f6, [v22, v22], globalState)];
				var v24: multiset<bool> := multiset{v14.f4};
				v23[safeIndex(856, v23.Length)] := fm53(cf9, fm54(|v24|, globalState), globalState);
				v23[safeIndex(856, v23.Length)] := v18;
				var v25 := DC31(f5);
				var v26 := new C5(cf10, f6, v25.cf50, cf9, v14.f4);
				if (!v26.f13 || (v0 < v0)) {
					var v27: array<set<int>> := new set<int>[25];
					var v28: array<bool> := new bool[6];
					v19[safeIndex(176, v19.Length)] := fm3(v0, cf8, globalState);
					v27, v28, v19[safeIndex(176, v19.Length)] := v27, v28, -safeDivisionInt(-v0, 0x3d);
					var v29: multiset<array<bool>> := multiset{f10};
					v29 := multiset{f10};
					cf10 := (v0 + f6) < (--v19[safeIndex(176, v19.Length)] * v26.f14);
					var v30: seq<int> := [-v26.f14];
					v0 := v30[safeIndex(v26.f14, |v30|)];
					var v31 := new C4(f5, v26.f14, cf8);
				} else {
					var v32: map<bool, int> := map[p0 := v0];
					v16 := v16[f6 + |v32| := cf9];
					v0 := safeModuloInt(-v26.f14, v26.f14);
					v0 := |v1|;
					var v33 := DC5(f10);
					v33 := v33;
					var v34: map<int, bool> := map[v26.f14 := cf10];
					v34 := v34[f6 - f6 := !v14.f4];
				}
				
				cf10 := !v14.f4;
			case DC10(cf11) =>
				var v35: map<int, array<bool>> := map[-|[f4]| := f10];
				v35 := v35[f6 := f10];
				var v36 := DC28(v0, cf11, false, f4, cf11);
				if (v36.cf46) {
					var v37: map<bool, string> := map[v14.f4 := v1];
					var v39: array<int> := new int[7] [154, safeModuloInt(|v1|, |(map[f4 := v0])[p0 := v0]|), fm7(globalState), safeModuloInt(|(set v38 : int | (0x24b <= v38) && (v38 < 174) :: (v38 - f6))|, cf11), f6, |multiset{v14.f4, f4, p0}|, f6];
					var v40: seq<bool> := [p0];
					v39[safeIndex(891, v39.Length)] := |(v40 + v40)|;
					var v41: array<array<int>> := new array<int>[27];
					v41[safeIndex(166, v41.Length)] := v39;
					var v42 := false;
					var v43 := 'h';
					var v44: map<bool, int> := map[v14.f4 := v0];
					v37, v39[safeIndex(891, v39.Length)], v41[safeIndex(166, v41.Length)], v42, v42 := map[f4 := v1] + (f9 + v37), v0, v39, fm0(v43, globalState), --|(v44 + v44)| == -0x201;
					var v45: set<bool> := {!v14.f4, !v14.f4};
					var v46 := DC33(false, cf11, multiset{false}, f6, v45);
					v42 := multiset{v14.f4} !! v46.cf54;
					var v47: array<map<int, array<char>>> := new map<int, array<char>>[23];
					var v48: multiset<int> := multiset{v39[safeIndex(891, v39.Length)], |v44|, cf11};
					var v49: array<char> := new char[12] [v43, fm1(v39[safeIndex(891, v39.Length)], v14.f4, globalState), 'c', v43, v43, v43, v43, 'y', v43, v43, 'l', v43];
					var v50: map<int, array<char>> := map[if (cf11 in v48) then v48[cf11] else 879 := v49];
					v47[safeIndex(836, v47.Length)] := v50;
					var v51 := DC25(true, cf11);
					var v52: set<int> := {v0};
					var v53: seq<int> := [v51.cf40, |[v52]|];
					var v54: map<int, bool> := map[|v53| := v14.f4];
					var v55: set<map<int, bool>> := {v54};
					var v56: multiset<bool> := multiset{v14.f4, v55 != v55, v42};
					v42, v42, v0, v47[safeIndex(836, v47.Length)] := p0 || v14.f4, v42, |v56|, if (p0) then v50 + v50 else v50 + v50;
					var v57: C1 := new C1(v39[safeIndex(891, v39.Length)], v39[safeIndex(891, v39.Length)], f5, if (p0 in v44) then v44[p0] else f6, v14.f4);
					var v58: seq<C1> := [if (f4) then v57 else v57, v57, v57, v57, v57];
					var v59: seq<seq<C1>> := [v58 + v58[safeIndex(|"y"|, |v58|) := v57]];
					v58 := v59[safeIndex(0x274, |v59|)];
					var v60 := DC20(v14.f4, v43, safeDivisionInt(f6, f6));
					var v61: map<char, array<bool>> := map[v43 := f10];
					var v62 := DC37(v61);
					v39[safeIndex(891, v39.Length)], v60, v42, v62 := |(set v63 : int | (0x177 <= v63) && (v63 < 806) :: (v63 * v57.f15))| * v0, v60, v14.f4, v62;
				} else {
					cf11 := v0;
					var v64 := false;
					f10[safeIndex(635, f10.Length)] := !v14.f4;
					var v65: array<int> := new int[2];
					var v66: map<int, array<int>> := map[v0 := v65];
					var v67: set<bool> := {false};
					var v68: seq<set<bool>> := [v67, v67, v67, v67];
					var v69 := DC30(v16);
					v64, v64, f10[safeIndex(635, f10.Length)], v64 := v0 >= |v66|, !(multiset{cf11} != (multiset{f6, 0x2fa})[f6 := abs(|v68|)]), true, !((v16 != v69.cf49) ==> v14.f4);
					v0 := v0;
					f10[safeIndex(635, f10.Length)] := p0 ==> (v14.f4 || p0);
					var v70: map<bool, bool> := map[p0 := f10[safeIndex(635, f10.Length)]];
					var v71: seq<int> := [cf11];
					var v72: map<D6, bool> := map[DC16(true, if (v14.f4 in v70) then v70[v14.f4] else v14.f4, v71) := !v64];
					v72 := v72;
				}
				
				var v73: seq<bool> := [true];
				var v74: set<int> := {v0, |v73|, v0};
				if (if (!v14.f4) then v74 > v74 else p0) {
					var v75 := false;
					var v76: multiset<int> := multiset{|DC22(fm50(v0, v0, cf11, v14.f4, globalState), p0, f4, f10, f6).cf30|, -0xa3, cf11};
					v75 := |v76[cf11 := abs(f6)]| == v0;
					cf11 := f6;
					cf11 := (f6 + f6) - |v1|;
					cf11 := if ((fm66(v0, globalState)).cf52) then v2.fm7(globalState) else |v73|;
					f10[safeIndex(31, f10.Length)] := p0;
					f10[safeIndex(31, f10.Length)] := v0 == v0;
				} else {
					var v77 := true;
					var v78 := DC25(v77, f6);
					v77 := v78.cf39;
					var v79: map<seq<int>, bool> := map[[cf11, f6] := v14.f4];
					var v80: seq<int> := [|multiset{f6}|, cf11];
					var v81 := new C1(|v79[v80 := v14.f4]|, |(seq(abs(0x190), i4  => (cf11)))|, f5, cf11, v14.f4);
					var v82: map<bool, int> := map[v77 := -f6];
					v81.f16 := v81.f16 + safeModuloInt(|v82|, v14.fm7(globalState));
					var v83: multiset<int> := multiset{v81.f16, |map[v81.f16 := v77]|, v81.f15};
					var v84: seq<seq<bool>> := [v73];
					var v85: multiset<int> := multiset{--v81.f16 * cf11, safeDivisionInt(f6, v0), if (0xdc in v83) then v83[0xdc] else if (f4 in v82) then v82[f4] else DC13(|v16|, v77, v0, v84[safeIndex(v81.f16, |v84|)]).cf16, f6};
					var v86: map<int, bool> := map[v81.f16 := f4];
					v85 := ((multiset(v80))[|"tfiwvdm"| := abs(|v86|)])[v0 := abs(--f6)];
					var v87: array<seq<map<D10, int>>> := new seq<map<D10, int>>[24];
					var v88 := DC21(v82);
					var v89: map<D10, int> := map[v88 := v0];
					var v90: seq<map<D10, int>> := [v89];
					v87[safeIndex(101, v87.Length)] := v90;
					f10[safeIndex(172, f10.Length)] := v14.f4;
					var v91: set<string> := {v1, "ugmeom", v1};
					var v92: map<string, int> := map[v1 + "q" := |multiset{f6, v81.f16, cf11, cf11, -|v91|}| * |v80|];
					var v93 := 'w';
					v77, v87[safeIndex(101, v87.Length)], f10[safeIndex(172, f10.Length)], r0, v92 := p0, v90, v14.f4, v93, v92[v1 := v81.f16];
				}
				
				f5[safeIndex(896, f5.Length)] := f10;
				var v94: seq<array<bool>> := [f10, f10, if (v14.f4) then f10 else f10, f10, if (v0 in v35) then v35[v0] else f10];
				f5[safeIndex(896, f5.Length)] := v94[safeIndex(safeModuloInt(v0, 0x7a), |v94|)];
			case DC8(cf7) =>
				var v95 := DC30(fm29(globalState));
				var v96: seq<seq<bool>> := [cf7];
				var v97 := DC40(map[f4 := DC27(v96)]);
				var v98: map<int, D20> := map[|cf7| := v97];
				f10[safeIndex(211, f10.Length)] := f4;
				var v99: multiset<seq<bool>> := multiset{cf7, [p0, f4, p0]};
				var v100 := DC15(v99[cf7[safeIndex(-524, |cf7|) := v14.f4] := abs(v0)]);
				v0, v95, v98, f10[safeIndex(211, f10.Length)], v100 := f6, v95, map[f6 := v97], p0, v100;
				var v101: array<string> := new string[4](i5 => (v1 + v1)[safeIndex(f6, |(v1 + v1)|) := 's']);
				v101[safeIndex(642, v101.Length)] := v1;
				var v102: seq<int> := [f6];
				var v103 := 'o';
				v101[safeIndex(642, v101.Length)] := (fm37(f4, if (v14.f4) then v14.f4 else v2.fm8(v0, true, v1, v102, globalState), f6, v14.f4, globalState))[safeIndex(v0, |fm37(f4, if (v14.f4) then v14.f4 else v2.fm8(v0, true, v1, v102, globalState), f6, v14.f4, globalState)|) := v103];
				if (p0) {
					f10[safeIndex(211, f10.Length)] := (cf7[safeIndex(4, |cf7|)] ==> v14.f4) && f4;
					var v104: array<bool> := new bool[6](i6 => DC3(f10[safeIndex(211, f10.Length)]).cf2);
					v104 := if (v102 <= v102) then v104 else v104;
					v0 := safeDivisionInt(safeDivisionInt(f6, |cf7|), 846);
					f10[safeIndex(211, f10.Length)] := !f10[safeIndex(211, f10.Length)];
					v0 := v0;
				} else {
					f10[safeIndex(211, f10.Length)], v101[safeIndex(642, v101.Length)], f10[safeIndex(211, f10.Length)], v0 := !(v14.f4 ==> true), v101[safeIndex(642, v101.Length)] + v1, if (v14.f4) then v0 <= v0 else f6 >= v0, v0;
					f10[safeIndex(211, f10.Length)] := DC9(fm0(v103, globalState), f6, f10[safeIndex(211, f10.Length)]).cf8 in map[true := v14.f4];
					var v105 := DC27(v96[safeIndex(f6, |v96|) := cf7]);
					v102, v0, f10[safeIndex(211, f10.Length)], f10[safeIndex(211, f10.Length)] := fm39(v101[safeIndex(642, v101.Length)], v105, v14.f4, globalState) + ([v0] + v102), v102[safeIndex(f6, |v102|)], v14.f4, !false;
					var v106: T1 := new C6(f10[safeIndex(211, f10.Length)], f4, f5, f6);
					var v107: map<T1, string> := map[v106 := v1];
					var v108: map<bool, int> := map[v14.f4 := f6];
					var v109: multiset<int> := multiset{v106.f6, if (!v2.fm8(|cf7|, v106.f4, v1, v102, globalState) in v108) then v108[!v2.fm8(|cf7|, v106.f4, v1, v102, globalState)] else f6};
					var v110: map<int, string> := map[v106.f6 := v1];
					var v111: map<bool, bool> := map[f10[safeIndex(211, f10.Length)] := v14.f4];
					var v112: multiset<bool> := multiset{f10[safeIndex(211, f10.Length)]};
					var v113: array<int> := new int[26] [v2.fm7(globalState), v0, fm9(true, false, f6, globalState), |(v16 + v16)|, v0, v0, |v107| * |cf7|, v0, f6 + v0, v106.f6, |(v102 + v102)|, |(v109 - v109)|, v106.f6, f6, if (v106.f4) then |v110| else f6, v0, |[false]|, f6, v106.f6, f6, |{f10[safeIndex(211, f10.Length)]}| * f6, v0, |v1|, |(v101[safeIndex(642, v101.Length)] + v101[safeIndex(642, v101.Length)])|, |v111|, |v112|];
					var v114: map<map<int, int>, bool> := map[v16 := v14.f4];
					v113[safeIndex(983, v113.Length)] := |(if (v14.f4) then [|v114|] else v102)|;
					v113[safeIndex(983, v113.Length)] := safeDivisionInt(v0, f6);
					var v115: array<bool> := new bool[1](i7 => v106.f4);
					var v116: map<char, array<bool>> := map[v103 := v115];
					var v117 := DC37(v116);
					var v118 := DC39(v117);
					var v119: map<D19, int> := map[v118 := -f6];
					var v120: map<char, int> := map[v103 := |v119|];
					v120 := v120[v103 := 0xeb];
				}
				
				var v121: T1 := new C1(669, f6, f5, |"kv"|, v14.f4);
				var v122: seq<seq<int>> := [fm57(v0, v0, globalState), v102, seq(abs(0x36b), i8  => (0xf6)), v102];
				v121 := new C5(v121.f4, fm3(fm3(f6, true, globalState), f10[safeIndex(211, f10.Length)], globalState), v121.f5, |v122|, p0);
			case DC11(cf12) =>
				var v123: map<string, char> := map[v1 := 'a'];
				var v124: multiset<bool> := multiset{v14.f4, true, v14.f4};
				var v125: map<D2, int> := map[DC6() := if (|v123| in v16) then v16[|v123|] else if (v14.f4 in v124) then v124[v14.f4] else v2.fm9(true, false, 0x278, globalState)];
				var v126 := DC6();
				v125 := v125[v126 := safeModuloInt(f6, f6)];
				var v127: array<int> := new int[25](i9 => safeDivisionInt(i9, v0));
				var v128 := DC32(v127);
				match (if (p0) then v128 else v128) {
					case DC33(cf52, cf53, cf54, cf55, cf56) =>
						var v129: map<bool, int> := map[v14.f4 := cf55];
						v0 := safeModuloInt(f6, |v129| + |v1|);
						var v130 := 'f';
						r0 := v130;
						var v131 := DC12('p');
						v131 := v131;
						var v132: array<set<seq<int>>> := new set<seq<int>>[6](i10 => {[f6, v0, f6], [cf55, |"dv"|, |{v130, v130}|], [v0]} - {[741]});
						var v133: seq<int> := [-cf55];
						var v134: multiset<T0> := multiset{v14};
						var v135: set<seq<int>> := {[|"tm"|], seq(abs(0x11b), i11  => (v0)), v133, v133, [0x3d6, |v134|]};
						v132[safeIndex(265, v132.Length)] := v135;
						f10[safeIndex(963, f10.Length)] := cf52;
						var v136 := DC16(v14.f4, f4, v133);
						var v137: set<D6> := {v136};
						var v138: map<map<int, int>, char> := map[fm42(v130, v137, globalState) := v130];
						v127[safeIndex(895, v127.Length)] := v0;
						var v139: seq<bool> := [cf52, cf52];
						var v140: seq<seq<bool>> := [v139, v139, v139];
						var v141 := DC27(v140);
						var v142: map<bool, bool> := map[v14.f4 := f4];
						var v143: map<int, bool> := map[f6 := if (p0 in v142) then v142[p0] else v14.f4];
						var v144 := DC19("mskr");
						v132[safeIndex(265, v132.Length)], f10[safeIndex(963, f10.Length)], v138, v127[safeIndex(895, v127.Length)] := {v133 + fm39(v1, v141, if (|v133| in v143) then v143[|v133|] else p0, globalState)}, f4, v138 + v138, |v144.cf25|;
					case DC32(cf51) =>
						var v145 := DC10(f6);
						v0 := safeModuloInt(v145.cf11, 0x2fd);
						var v146: seq<bool> := [false, v2.fm8(v0, f4, v1, seq(abs(133), i12  => (f6)), globalState)];
						f10[safeIndex(257, f10.Length)] := v146[safeIndex(f6, |v146|)];
						var v147: array<D0> := new D0[12](i13 => DC1());
						var v148 := true;
						var v149 := 'p';
						var v150: array<string> := new string[10] [v1, "xkjwa", v1, ("tc" + v1)[safeIndex(-|map[v14.f4 := v149]|, |("tc" + v1)|) := 'c'], v1, "vtrm", v1, if (v14.f4) then v1 else v1, seq(abs(0x1d7), i14  => ('h')), v1];
						v150[safeIndex(38, v150.Length)] := if (f4) then v1 else v1;
						f10[safeIndex(257, f10.Length)], v147, v148, v0, v150[safeIndex(38, v150.Length)] := v14.f4, DC42(v147).cf68, v14.f4, v0, v1 + v1;
						var v151: map<bool, bool> := map[false := f10[safeIndex(257, f10.Length)]];
						var v152: set<int> := {f6};
						var v153: map<map<bool, bool>, int> := map[v151 := |v152|];
						v148, v148 := fm0('o', globalState), !(v153 == v153);
						v0 := v0;
				}
				
				var v154: map<string, int> := map[v1 := |v1|];
				var v155: map<map<string, int>, int> := map[v154 := v0];
				v155 := v155[map[v1 := -f6] := f6];
				var v156: seq<int> := [v0];
				var v157: set<bool> := {p0};
				v0 := v156[safeIndex(|v157|, |v156|)];
		}
		
		var v158 := true;
		v158 := f4;
		var v159 := 'k';
		r0 := v159;
		v0 := v0;
		r0 := v159;
		var v160: array<D1> := new D1[19];
		r1 := v160;
	}
	method m6(globalState: GlobalState) returns (r0: D1, r1: string) {
		var i0 := 0;
		while (f4)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v0: array<map<array<bool>, bool>> := new map<array<bool>, bool>[13];
			var v1: map<array<bool>, bool> := map[f10 := f4];
			v0[safeIndex(868, v0.Length)] := v1 + v1;
			v0[safeIndex(868, v0.Length)] := v1 + ((map[f10 := f4])[f10 := f4] + v1);
			var v2: array<int> := new int[10];
			v2[safeIndex(304, v2.Length)] := f6;
			var v3 := 'w';
			f10[safeIndex(597, f10.Length)] := fm0(v3, globalState);
			var v4 := true;
			v2[safeIndex(304, v2.Length)], f10[safeIndex(597, f10.Length)], v4 := f6, f4, v4;
			var v5: array<map<int, int>> := new map<int, int>[13](i1 => map[f6 := v2[safeIndex(304, v2.Length)]] + DC30(map[f6 := f6]).cf49);
			var v6: seq<int> := [DC4(v2[safeIndex(304, v2.Length)], v2[safeIndex(304, v2.Length)]).cf3];
			var v7: map<int, int> := map[|v6| := f6];
			v5[safeIndex(958, v5.Length)] := v7;
			v5[safeIndex(958, v5.Length)] := v7;
			var v8 := new C6(fm0(v3, globalState), false, f5, safeModuloInt(f6, f6));
		}
		for i2 := f6 to f6 {
			f10[safeIndex(324, f10.Length)] := false || true;
			f10[safeIndex(324, f10.Length)] := f6 >= i2;
			f10[safeIndex(324, f10.Length)] := f10[safeIndex(324, f10.Length)];
			var v9: array<map<bool, bool>> := new map<bool, bool>[24](i3 => map[true := f10[safeIndex(324, f10.Length)]]);
			var v10 := DC23(v9);
			var v11: map<int, D11> := map[f6 := v10];
			var v12 := 472;
			v11, v12 := v11, i2;
			f10[safeIndex(324, f10.Length)], v12, v12, f10[safeIndex(324, f10.Length)] := !f10[safeIndex(324, f10.Length)], --(i2 * 0x2b9), -v12 - fm7(globalState), false;
		}
		var v13 := false;
		v13 := v13;
		f10[safeIndex(304, f10.Length)] := v13 || f4;
		f10[safeIndex(304, f10.Length)] := v13;
		var v14 := "ptmrel";
		var v15: seq<int> := [|(if (f4 in f9) then f9[f4] else v14)|, f6, f6, f6];
		var v16: seq<bool> := [f10[safeIndex(304, f10.Length)]];
		var v17: multiset<seq<bool>> := multiset{v16};
		var v18: map<bool, D6> := map[true := DC15(v17)];
		if (v15[safeIndex(f6, |v15|)] > |v18|) {
			var v19: set<int> := {f6};
			var v20: map<bool, int> := map[f10[safeIndex(304, f10.Length)] := |v19|];
			var v21 := DC21(v20);
			match (v21) {
				case DC22(cf30, cf31, cf32, cf33, cf34) =>
					cf30 := cf30 + ("mlwyh" + cf30);
					var v22: array<int> := new int[28](i4 => i4 - 0xc7);
					var v23: map<int, int> := map[|v15| := f6];
					v22[safeIndex(968, v22.Length)] := if (|v14| in v23) then v23[|v14|] else cf34;
					v22[safeIndex(968, v22.Length)] := f6;
					var v24 := DC25(cf31, f6);
					var v25 := new C5((v24.(cf39 := v13)).cf39, -f6, f5, |"romg"|, f10[safeIndex(304, f10.Length)]);
					var v26: array<multiset<bool>> := new multiset<bool>[18](i5 => multiset{cf32});
					v26 := v26;
				case DC21(cf29) =>
					var v27 := 0x173;
					v27 := v27;
					var v28: multiset<bool> := multiset{f10[safeIndex(304, f10.Length)], f4, f4, v13};
					var v29: multiset<int> := multiset{v27, |v28|};
					var v30: map<multiset<int>, string> := map[v29 := v14];
					var v31, v32 := m0(v30, f10[safeIndex(304, f10.Length)], v29, 0xfd > |v15[safeIndex(v27, |v15|) := v27]|, globalState);
					v27 := f6;
					f10[safeIndex(304, f10.Length)] := f4;
			}
			
			if (v13) {
				var v33: multiset<int> := multiset{f6};
				var v34: map<multiset<int>, string> := map[v33 := v14];
				var v35, v36 := m0(v34, true, multiset(v15), v13, globalState);
				var v37: array<set<string>> := new set<string>[6](i6 => {seq(abs(0x3a1), i7  => ('a'))} + {v14, "mj"});
				var v38: set<string> := {v14};
				v37[safeIndex(140, v37.Length)] := v38;
				v37[safeIndex(140, v37.Length)] := v38;
				v36 := v36;
				var v39 := new C4(f5, v15[safeIndex(f6, |v15|)], f4);
				var v40 := DC13(v36, v35, |"wdwcp"|, v16);
				var v41 := 'i';
				v36, v35, v36, v36, v36 := f6 - v40.cf14, v13, |(v14[safeIndex(v36, |v14|) := v41] + (v14[safeIndex(f6, |v14|) := v41] + v14))|, v36, v36 + 0x1b7;
			} else {
				var v42: array<int> := new int[26](i8 => i8 - f6);
				v42[safeIndex(291, v42.Length)] := f6 - f6;
				v42[safeIndex(291, v42.Length)] := f6;
				v42[safeIndex(291, v42.Length)] := v42[safeIndex(291, v42.Length)] * v42[safeIndex(291, v42.Length)];
				var v44 := DC14(multiset(v15));
				var v45: map<D5, bool> := map[v44 := f10[safeIndex(304, f10.Length)]];
				var v46: map<int, bool> := map[|(map v43 : D5 | v43 in v45 :: (v43) := (|v14|))| := f4];
				v46 := v46;
				v13 := v13;
				v42[safeIndex(291, v42.Length)] := f6;
			}
			
			v15 := [safeModuloInt(f6, |v20|), f6];
			var v47 := 0x163;
			v47 := if (v13) then f6 else v47;
			f10[safeIndex(304, f10.Length)] := f10[safeIndex(304, f10.Length)];
		} else {
			var v48: array<D21> := new D21[1];
			var v49: array<D0> := new D0[1](i9 => DC1());
			var v50 := DC42(v49);
			v48[safeIndex(777, v48.Length)] := v50.(cf68 := v49);
			v48[safeIndex(777, v48.Length)] := v50;
			var v51 := 'a';
			v51 := v51;
			var v52: set<bool> := {true};
			f10[safeIndex(304, f10.Length)] := v52 >= v52;
			var v53: seq<seq<int>> := [v15, v15, v15];
			var v54: multiset<int> := multiset{f6, f6, |v53|};
			var v55: map<int, bool> := map[|v54| := f10[safeIndex(304, f10.Length)]];
			f10[safeIndex(304, f10.Length)] := v13 <==> fm14(v55, !f4, f6, f10[safeIndex(304, f10.Length)], globalState);
			var v56: array<int> := new int[13](i10 => safeDivisionInt(i10, f6));
			v56[safeIndex(193, v56.Length)] := f6 + fm3(f6, true, globalState);
			v56[safeIndex(193, v56.Length)] := f6;
		}
		
		var v57 := 'n';
		if (fm0(v57, globalState)) {
			f10[safeIndex(304, f10.Length)] := v13;
			var v58: C6 := new C6(f4, v13, f5, if (false) then f6 else f6);
			var v59 := -0x318;
			v58, f10[safeIndex(304, f10.Length)], v59 := v58, fm9(v58.f11, f10[safeIndex(304, f10.Length)], |v14|, globalState) == v59, v59;
			v59 := f6;
			v59 := v59;
			var v60: array<int> := new int[15](i11 => i11 * v59);
			v60[safeIndex(231, v60.Length)] := safeDivisionInt(v59, f6);
			var v61: multiset<int> := multiset{-v59};
			var v62: map<multiset<int>, bool> := map[v61 := v13];
			var v63 := DC14(v61);
			var v64: map<int, multiset<int>> := map[v59 := v63.cf18];
			var v65: map<bool, bool> := map[f10[safeIndex(304, f10.Length)] := if ((if (v59 in v64) then v64[v59] else v61) in v62) then v62[if (v59 in v64) then v64[v59] else v61] else v58.f11];
			v60[safeIndex(231, v60.Length)], v13, v59, v58 := -(v59 - (|v15| * v59)), if (v58.f11 in v65) then v65[v58.f11] else v58.f11, fm7(globalState), v58;
		} else {
			var v66 := 0x139;
			v66 := safeDivisionInt(v66, f6 + |"gpm"|);
			var v67: array<bool> := new bool[27];
			v67 := v67;
			var v68 := DC19("p");
			var v69: map<D9, bool> := map[v68 := f4];
			var v70: map<map<D9, bool>, char> := map[v69 + v69 := v57];
			v70 := v70[v69 := v57];
			v66 := -(v66 + v66);
			f10[safeIndex(304, f10.Length)] := (f6 == f6) && !(v15 != v15);
		}
		
		var v71: set<int> := {f6};
		var v72: map<int, bool> := map[-f6 := !!v13];
		var v73: array<int> := new int[28] [f6, f6, f6, |v71|, f6, f6, f6, 617, f6, 0x34c, f6, f6, -0x3bd, f6, 0x35c, |v14|, -f6, f6, f6, f6, |v14|, f6, f6, f6, f6, f6, |v72|, 0x2aa];
		var v74: map<array<int>, int> := map[v73 := f6];
		r0 := DC3(fm8(f6, v13, seq(abs(-0x1ed), i12  => (v57)), v15, globalState)).(cf2 := f6 >= |v74|);
		r1 := v14;
	}
}

class C8 {
	constructor () {
	}
	
	function fm11(globalState: GlobalState): bool {
		"prxjdttbg" <= "or"
	}
	function fm12(p0: bool, p1: D0, p2: int, p3: bool, globalState: GlobalState): bool {
		|map[true := map[0x356 := |[true, false]|]]| !in ((seq(abs(0x9d), i0  => (-0x278))) + [-0x151, |map[false := false]|])
	}
	method m4(p0: int, p1: int, p2: bool, globalState: GlobalState) {
		var v0 := DC0(0x3da);
		var i0 := 0;
		while (!fm12(p2, v0, p1, p2, globalState))
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v1 := 'h';
			v1 := v1;
			var v2 := "mhlv";
			v2 := v2 + v2;
			var v3 := true;
			var v4: array<int> := new int[25];
			v3, v4 := p2, v4;
			var v5 := 0x151;
			v3, v5 := -p0 >= |v2|, |v2|;
		}
		var v6 := 0x277;
		v6 := p1 * p0;
		var v7: array<int> := new int[4];
		v7[safeIndex(803, v7.Length)] := if (true) then 0xc2 else p1;
		var v8 := DC26(p2);
		var v9: map<bool, bool> := map[p2 := p2];
		var v10 := "drtvj";
		var v11 := DC28(p0, |v10|, p2, !p2, p0);
		var v12 := 'a';
		var v13: seq<bool> := [p2];
		var v14 := DC43(p2, false, v13);
		var v15: seq<bool> := [p2, v14.cf70, p2];
		var v16: array<bool> := new bool[21] [p2, p2, true, p2, p2, v8.cf41, p2, p2, p2, p2, p2, p2, p2, p2, p2, p2, if (p2 in v9) then v9[p2] else true, p2, DC20(v11.cf46, v12, v6).cf26, !v13[safeIndex(p0, |v13|)], p2];
		var v17: array<array<bool>> := new array<bool>[28] [v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16, v16];
		var v18: multiset<int> := multiset{-0x1f6};
		var v19: T1 := new C3(v17, v6 * p0, v18 < v18);
		v19.f5[safeIndex(632, v19.f5.Length)] := v16;
		var v20: map<int, int> := map[p0 := v6];
		var v21: multiset<seq<bool>> := multiset{v15[safeIndex(|v20|, |v15|) := v19.f4]};
		var v22 := DC15(v21);
		v7[safeIndex(803, v7.Length)], v19, v12, v19.f5[safeIndex(632, v19.f5.Length)] := p0, v19, match v22 {
			case DC16(cf20, cf21, cf22) => v12
			case DC15(cf19) => v10[safeIndex(v6, |v10|)]
		}, v16;
		var i1 := 0;
		while (v19.f4)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			v7 := v7;
			var v23: array<string> := new seq<char>[23](i2 => seq(abs(0xcd), i3  => (v12)));
			var v24 := DC4(v7[safeIndex(803, v7.Length)], v6);
			var v25: seq<int> := [|[fm47(v19.f6, globalState)]|];
			var v26: C1 := new C1(v7[safeIndex(803, v7.Length)], -616, v19.f5, |v13|, p2);
			v23[safeIndex(626, v23.Length)] := fm50(v24.cf3, |v25[safeIndex(p1, |v25|) := |[p1, v19.f6]|]|, |map[fm0(v12, globalState) := v26]|, v19.f4, globalState);
			v23[safeIndex(626, v23.Length)] := (if (v19.f4) then v10 + v10 else v10)[safeIndex(|v10| + |v10|, |(if (v19.f4) then v10 + v10 else v10)|) := v12];
			var v27: map<char, int> := map['a' := p0];
			v27 := v27[v12 := p1];
			v17 := v19.f5;
		}
		var v28 := DC14(v18);
		var v29: seq<D5> := [v28.(cf18 := v18), v28, v28];
		var v30: set<seq<D5>> := {v29};
		v30 := v30;
		var v31 := false;
		v31 := v19.f4;
	}
	method m5(p0: int, globalState: GlobalState) returns (r0: bool, r1: bool) {
		var v0 := false;
		if (v0) {
			r0 := !v0;
			var v1: set<bool> := {v0, v0};
			r1 := v1 >= v1;
			var v2: map<int, int> := map[p0 := p0 - |[259]|];
			v2 := v2[p0 := p0];
			r1 := p0 <= p0;
			var v3: array<array<bool>> := new array<bool>[4];
			var v4 := new C3(v3, p0, v0);
		} else {
			var v5 := "h";
			v0 := (p0 * -|v5|) == 0x319;
			if (!v0) {
				var v6 := DC6();
				v6 := v6;
				var v7: seq<int> := [0x10e, -p0];
				var v8: seq<bool> := [v0, v0];
				var v9: multiset<int> := multiset{|multiset(v8)|};
				var v10: array<int> := new int[23] [p0, |v5|, p0, |v7|, -0x5b, p0, p0, p0, |v9|, p0, p0, p0, -870, p0, p0, p0, p0, p0, 0x3b7, 0x1a8, p0, p0, p0];
				var v11: set<array<int>> := {v10, v10, v10};
				var v12 := DC29(v11);
				var v13: set<D13> := {v12.(cf48 := v11), v12};
				v0 := v13 > v13;
				var v14: map<int, bool> := map[p0 * p0 := v0];
				v14 := v14[p0 - v7[safeIndex(-179, |v7|)] := v0];
				var v15 := 'd';
				var v16: map<bool, char> := map[v0 := v15];
				var v17: map<bool, int> := map[v0 := |v16|];
				var v18 := DC21(v17);
				var v19: set<int> := {p0, 85};
				var v22: array<set<int>> := new set<int>[19] [v19 * v19, v19 - v19, fm31(|v14|, p0, p0, false, globalState), v19, {|v5|}, v19, {738, p0} - v19, v19, v19, v19, v19, set v20 : int | (0x8d <= v20) && (v20 < 0x3d6) :: (v20 - p0), {0xe6}, v19, v19, {p0} * v19, set v21 : int | (-0x81 <= v21) && (v21 < 0x2dc) :: (safeDivisionInt(v21, p0)), v19, v19];
				v22[safeIndex(947, v22.Length)] := v19 - v19;
				var v23: seq<string> := [v5];
				var v24: multiset<bool> := multiset{v0, false, !v0};
				var v25: set<bool> := {v0};
				var v26 := DC33(v0, |v5|, v24, p0, v25);
				var v27 := DC17({p0, p0, p0, p0, p0});
				v0, v18, r1, v22[safeIndex(947, v22.Length)] := false ==> (v15 !in v23[safeIndex(p0, |v23|)]), DC21(v17), v26.cf52, v27.cf23;
				var v28 := 0x71;
				v28 := -p0;
			} else {
				var v29 := 0x364;
				v29 := -p0;
				r0 := v0;
				var v30: multiset<int> := multiset{v29};
				var v31: seq<multiset<int>> := [v30];
				v29 := |v31|;
				var v32: array<array<bool>> := new array<bool>[27];
				var v33: map<int, int> := map[p0 := p0];
				var v34: map<map<int, int>, bool> := map[v33 := v0];
				var v35 := new C4(v32, p0, fm67(v0, globalState) == v34);
				var v36: array<bool> := new bool[28](i0 => {v0, v0} >= {v0, v0});
				var v37 := DC44(v0);
				v36[safeIndex(521, v36.Length)] := (v37.(cf72 := v0)).cf72;
				v36[safeIndex(521, v36.Length)] := true;
			}
			
			v5 := v5;
			var v38: array<int> := new int[4](i1 => i1 + p0);
			v38[safeIndex(377, v38.Length)] := 647;
			var v39: multiset<bool> := multiset{!v0};
			v38[safeIndex(377, v38.Length)] := |v39| * p0;
			var v40: seq<bool> := [v0, v0, v0, v0];
			var v41 := DC13(0x89, v0, |v5|, v40);
			var v42: multiset<int> := multiset{v38[safeIndex(377, v38.Length)]};
			v38[safeIndex(377, v38.Length)] := |((v41.cf17[safeIndex(-(if (p0 in v42) then v42[p0] else p0), |v41.cf17|) := v0] + v40) + (v40 + v40))|;
		}
		
		var v43: set<bool> := {false};
		v0 := (v43 - v43) !! (v43 - v43);
		if (v0) {
			var v44 := 'p';
			var v45 := DC12(v44);
			match (v45.(cf13 := if (v0) then v44 else v44)) {
				case DC13(cf14, cf15, cf16, cf17) =>
					cf16 := fm3(cf14, cf15, globalState);
					var v46 := new C2("bgj", v0);
					var v47: seq<int> := [cf16 + p0, -328];
					v47 := v47 + v47;
					r0 := cf15;
				case DC12(cf13) =>
					var v48: array<array<bool>> := new array<bool>[10];
					var v49: C1 := new C1(p0, p0, v48, p0, v0);
					v49 := v49;
					v49.f16 := v49.fm7(globalState);
					var v50: map<bool, bool> := map[v0 := v0];
					var v51: multiset<bool> := multiset{v0};
					v50 := v50[true := multiset([v0]) == v51];
					var v52: multiset<char> := multiset{v44, v44, cf13, v44};
					v52 := if (v0) then v52 else multiset{cf13};
			}
			
			var v53: array<char> := new char[5];
			v53[safeIndex(201, v53.Length)] := v44;
			v53[safeIndex(201, v53.Length)] := v44;
			var v54 := -0x245;
			v54 := v54;
			var v55 := "pbuouf";
			var v56: multiset<string> := multiset{v55, v55};
			var v57 := DC25(-p0 == p0, safeModuloInt(--|v56|, -p0));
			v57 := v57;
			var v58: array<int> := new int[4] [-v54, p0, v54, v54];
			var v59 := DC32(v58);
			match (v59) {
				case DC33(cf52, cf53, cf54, cf55, cf56) =>
					r0 := !(if (cf52) then if (v0) then v0 else true else v0);
					var v60 := DC26(v0);
					var v61: map<bool, bool> := map[v0 := v0];
					var v62: map<D11, bool> := map[v60 := if (v0 in v61) then v61[v0] else cf52];
					var v63: seq<map<D11, bool>> := [v62];
					var v64 := DC35(v63);
					var v65: array<array<int>> := new array<int>[24];
					var v66: map<char, bool> := map[v44 := v0];
					var v67: array<array<bool>> := new array<bool>[10];
					var v68: T1 := new C6(v0, if ('g' in v66) then v66['g'] else v0, v67, cf55);
					var v69: array<T1> := new T1[27] [v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68, v68];
					var v70: multiset<array<T1>> := multiset{v69, v69};
					var v71: map<multiset<array<T1>>, char> := map[if (cf52) then v70 else v70 := v44];
					v53[safeIndex(201, v53.Length)], cf52, v64, v65, v44 := fm1(|v55|, v0, globalState), cf52, v64, v65, if (multiset{v69} in v71) then v71[multiset{v69}] else v44;
					var v72: seq<set<bool>> := [cf56];
					var v73: map<int, int> := map[-cf55 := -0x8b];
					var v74: seq<bool> := [cf52, v68.f4];
					var v75 := DC19(v55);
					var v76: array<bool> := new bool[29] [cf52, v0, cf52, true, !cf52, v68.f4, cf52 && false, |(seq(abs(614), i2  => (v53[safeIndex(201, v53.Length)])))| < v54, cf52, cf52, v0, v54 > -p0, !v0 in v72[safeIndex(if (v68.f6 in v73) then v73[v68.f6] else v54, |v72|)], v0, fm0(v44, globalState), v74[safeIndex(p0, |v74|)], v0, false, cf52, v68.f4, v0, cf52, cf52, v0, v68.f4, (seq(abs(0x37f), i3  => (v53[safeIndex(201, v53.Length)]))) <= (v75.(cf25 := v55)).cf25, false, v68.f4, v68.f4 ==> cf52];
					v76[safeIndex(172, v76.Length)] := cf52;
					var v77: map<map<int, int>, bool> := map[v73 := v68.f4];
					var v78: map<int, array<int>> := map[v68.f6 := v58];
					var v79: map<int, array<int>> := map[fm3(|v77|, v0, globalState) := if (cf53 in v78) then v78[cf53] else v58];
					var v80 := DC41(v54);
					var v81: multiset<int> := multiset{v68.f6};
					var v82: seq<int> := [v68.f6, cf55];
					var v83: multiset<int> := multiset{|v81|, |v82|, v54, cf53, 0x3af};
					var v84: C1 := new C1(0x39d, p0, v67, cf53, !v0);
					var v85: map<int, seq<int>> := map[625 := [|v81[cf55 := abs(|map[v84 := |cf54|]|)]|]];
					v76[safeIndex(172, v76.Length)], cf52, r0, v79 := v68.fm8(v80.cf67, cf52, seq(abs(69), i4  => (v53[safeIndex(201, v53.Length)])), if (0x354 in v85) then v85[0x354] else v82, globalState), v68.f4, !v68.f4, v79 + v79;
					var v86: array<map<bool, D11>> := new map<bool, D11>[29];
					var v87: set<array<bool>> := {v76, v76};
					var v88 := DC45(v86);
					var v89: map<set<array<bool>>, array<map<bool, D11>>> := map[v87 := v88.cf73];
					v86 := if (v87 in v89) then v89[v87] else v86;
				case DC32(cf51) =>
					var v90: array<bool> := new bool[22];
					var v91: seq<int> := [v54, p0, p0, 0xd];
					var v92: map<seq<array<bool>>, int> := map[[v90, v90, v90, v90, v90] := |multiset(v91)|];
					var v93: map<int, seq<bool>> := map[if ([v90, v90, v90, v90] in v92) then v92[[v90, v90, v90, v90]] else v54 := fm40(false, true, p0, globalState)];
					v93 := v93[0xc0 := fm19(v0, p0, globalState)];
					v90[safeIndex(700, v90.Length)] := fm11(globalState);
					v90[safeIndex(700, v90.Length)] := v0;
					var v94: array<string> := new seq<char>[5](i5 => seq(abs(370), i6  => (v44)));
					var v95: array<map<int, bool>> := new map<int, bool>[24](i7 => map[|map[p0 := -v54]| := v90[safeIndex(700, v90.Length)]]);
					v94, v95 := v94, v95;
					var v96: seq<string> := [v55, seq(abs(0x20d), i8  => (v44)), v55, v55];
					var v97 := DC27(seq(abs(195), i9  => ([v0])));
					var v98: set<int> := {v54, p0};
					var v99: map<bool, set<int>> := map[!v90[safeIndex(700, v90.Length)] := v98];
					r1, v91, v90[safeIndex(700, v90.Length)], v54, v54 := p0 >= p0, fm39(v55[safeIndex(v54, |v55|) := v53[safeIndex(201, v53.Length)]] + v96[safeIndex(p0, |v96|)], v97, v0, globalState), fm11(globalState), |(if (false in v99) then v99[false] else v98)| + v54, v54;
			}
			
		} else {
			var v100: multiset<bool> := multiset{v0};
			r0 := (v100 - v100) !! multiset{v0, v0, v0};
			var v101: array<bool> := new bool[24](i10 => v0 <==> v0);
			v101[safeIndex(935, v101.Length)] := v0;
			var v102: map<bool, int> := map[v0 := p0];
			v101[safeIndex(935, v101.Length)] := (if (v0 in v102) then v102[v0] else p0) == fm3(|v43|, v0, globalState);
			var v103: array<int> := new int[11](i11 => i11 - p0);
			v103[safeIndex(247, v103.Length)] := -p0;
			v103[safeIndex(247, v103.Length)] := p0;
			var v104: array<map<bool, D11>> := new map<bool, D11>[20];
			var v105 := DC45(v104);
			var v106: array<D22> := new D22[4] [v105, v105, v105.(cf73 := v105.cf73), v105];
			v106[safeIndex(442, v106.Length)] := v105;
			v106[safeIndex(442, v106.Length)] := v105;
			var v107: C0 := new C0(v103[safeIndex(247, v103.Length)]);
			var v108: seq<C0> := [v107, v107];
			var v109: map<int, C0> := map[|map[v0 := p0]| := v107];
			var v110 := DC48(v107);
			var v111: array<C0> := new C0[24] [v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v107, v108[safeIndex(|[v107.f12, p0]|, |v108|)], v107, v107, if (p0 in v109) then v109[p0] else v107, v107, v110.cf82, v107, v107, v107, v107, v107, v110.cf82];
			v111 := v111;
		}
		
		var v112: array<int> := new int[29](i13 => i13 * p0);
		forall i12 | 0 <= i12 < v112.Length {
			v112[i12] := safeModuloInt(i12, p0 + p0);
		}
		var v113: array<C3> := new C3[19];
		var v114: array<bool> := new bool[26](i14 => v0);
		var v115: array<array<bool>> := new array<bool>[23] [v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114, v114];
		var v116 := "clpbitvi";
		var v117: seq<string> := ["vvtigi", v116];
		var v118: C3 := new C3(v115, |v117|, v0);
		v113[safeIndex(124, v113.Length)] := v118;
		var v119: map<bool, int> := map[v0 := p0];
		var v120: C0 := new C0(-|v119|);
		r1, v113[safeIndex(124, v113.Length)], v120 := v0, v118, v120;
		var v121 := -0xbf;
		v121 := -v120.f12 - v120.f12;
		var v122 := DC47(|v116|, !!v0, v0, v120.f12, -p0);
		r0 := v122.cf79;
		r1 := v0;
	}
}

class C9 extends T1 {
	var f7 : int
	const f8 : set<seq<bool>>
	constructor (f7 : int, f8 : set<seq<bool>>, f5 : array<array<bool>>, f6 : int, f4 : bool) {
		this.f7 := f7;
		this.f8 := f8;
		this.f5 := f5;
		this.f6 := f6;
		this.f4 := f4;
	}
	
	function fm8(p0: int, p1: bool, p2: seq<char>, p3: seq<int>, globalState: GlobalState): bool {
		f4
	}
	function fm9(p0: bool, p1: bool, p2: int, globalState: GlobalState): int {
		match if (f4) then DC4(|[f7]|, 0x390) else DC4(f6, 0x188) {
			case DC4(cf3, cf4) => -cf3
			case DC3(cf2) => f6
		}
	}
	function fm6(p0: D0, p1: D0, p2: map<int, char>, p3: int, globalState: GlobalState): multiset<int> {
		multiset{|multiset([f4])|, f6, f7, -280} + (multiset{|"fku"|} * multiset{|map[f7 := f4]|})
	}
	function fm7(globalState: GlobalState): int {
		f7
	}
	method m1(p0: int, globalState: GlobalState) returns (r0: map<bool, int>) {
		var v0: array<bool> := new bool[18](i1 => f4);
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := f6 >= p0;
		}
		for i2 := safeModuloInt(f7, f7) to -safeDivisionInt(f7, f7) {
			var v1 := DC4(f6, p0);
			var v2: map<bool, D1> := map[!f4 := if (false) then v1 else v1];
			var v3: seq<char> := ['f'];
			var v4: map<int, seq<char>> := map[f6 := v3];
			var v5 := 'h';
			var v6: map<bool, int> := map[f4 := |v3[safeIndex(i2, |v3|) := v5]|];
			v2 := v2[f4 := fm10(f4, false, v4, |v6|, globalState)];
			var v7 := true;
			v7 := f4;
			var v8 := new C0(i2 + -f6);
			if (v7) {
				var v9: multiset<int> := multiset{v8.f12, |multiset{|v6|, 0x376}|};
				var v10: map<multiset<int>, int> := map[v9 * v9 := 451];
				var v11 := DC0(i2);
				var v12: map<int, char> := map[f6 := v5];
				v10 := v10[multiset{p0, -v8.f12, f7} + fm6(v11, v11, v12, 0xfb, globalState) := f7];
				v7, f7 := f4, |[[i2]]|;
				v7 := v7;
				var v13: array<string> := new seq<char>[14](i3 => v3);
				v13[safeIndex(656, v13.Length)] := fm16(v7, f4, !v7, DC6(), globalState);
				var v14: multiset<D0> := multiset{v11, v11};
				var v15: array<multiset<D0>> := new multiset<D0>[2] [v14, v14];
				v15[safeIndex(913, v15.Length)] := multiset{v11};
				var v16: seq<D0> := [v11.(cf0 := -i2)];
				v13[safeIndex(656, v13.Length)], v15[safeIndex(913, v15.Length)] := v3, v14 - multiset(v16 + v16);
				var v17 := DC28(p0, f7, f4, true, p0);
				var v18: seq<bool> := [v17.cf46, f4, v7 <==> v7];
				var v19: seq<int> := [f6];
				v18 := (v18[safeIndex(i2, |v18|) := v18[safeIndex(|v19|, |v18|)]])[safeIndex(safeModuloInt(f6, f7), |v18[safeIndex(i2, |v18|) := v18[safeIndex(|v19|, |v18|)]]|) := !true];
			} else {
				var v20: map<string, char> := map[v3 := fm1(f7, v7, globalState)];
				v20 := v20["yrkp" := v5];
				var v21 := new C6(f4, f4, if (v7) then f5 else f5, v8.f12);
				var v22: set<bool> := {!v7};
				v7 := v22 != fm46('t', globalState);
				var v23 := DC9(true, p0, f4);
				var v24: map<int, int> := map[-f6 := v23.cf9];
				var v25: seq<bool> := [v7, f4];
				var v26: seq<int> := [|v25|, f6];
				var v27 := DC16(false, v21.f11, v26);
				var v28: set<D6> := {v27, v27, v27};
				v24, v7, v7, f7, f7 := fm42(v5, v28, globalState) + v24, i2 < v8.f12, v8.f12 != fm7(globalState), p0 - 0x159, p0;
				v7 := v8.f12 != (f6 - f7);
			}
			
		}
		var i4 := 0;
		while (f4)
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			f7 := -f6;
			var v29 := "cnnryc";
			var v30: T0 := new C2(v29, f4);
			var v31: seq<T0> := [v30, v30, v30];
			var v32 := DC51(v30);
			var v33: array<T0> := new T0[18] [v30, v31[safeIndex(p0, |v31|)], v31[safeIndex(|v29|, |v31|)], if (f4) then v30 else v30, v32.cf87, v30, if (f4) then v30 else v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30, v30];
			v33[safeIndex(704, v33.Length)] := v30;
			v33[safeIndex(704, v33.Length)] := v30;
			match (fm68(globalState)) {
				case DC28(cf43, cf44, cf45, cf46, cf47) =>
					cf43 := |(v29 + v29)|;
					var v34: map<bool, bool> := map[v30.f4 := !cf46];
					v34 := v34;
					var v35: array<set<int>> := new set<int>[28](i5 => {|{cf46, v30.f4, !v30.f4, f4, cf45}|});
					v35[safeIndex(441, v35.Length)] := {cf44};
					v35[safeIndex(441, v35.Length)] := {p0, p0, |[cf46]|, cf43, fm7(globalState)} * {f6};
					cf45, cf46 := cf45, (v35[safeIndex(441, v35.Length)] * v35[safeIndex(441, v35.Length)]) > v35[safeIndex(441, v35.Length)];
				case DC27(cf42) =>
					var v36: seq<bool> := [f4, v30.f4, v30.f4];
					v36 := v36[safeIndex(f7, |v36|) := f4];
					var v37 := false;
					v37 := v37;
					v0[safeIndex(68, v0.Length)] := v37;
					var v38: multiset<bool> := multiset{f4, v30.f4, true};
					v0[safeIndex(68, v0.Length)] := v38 >= multiset{v37};
					f7 := p0;
			}
			
			var v39: map<T0, bool> := map[v30 := v30.f4];
			var v40 := DC28(f7, |v29|, v30.f4, true, |v29|);
			var v41: set<bool> := {if (v33[safeIndex(704, v33.Length)] in v39) then v39[v33[safeIndex(704, v33.Length)]] else v30.f4, false, v40.cf45, f4, f4};
			var v42: map<bool, bool> := map[v30.f4 := v30.f4];
			var v43: seq<map<bool, bool>> := [v42, map[f4 := v30.f4]];
			var v44: seq<set<bool>> := [v41, fm53(f6, v43, globalState)];
			f7 := safeDivisionInt(|v44[safeIndex(p0, |v44|)]|, 0x134 + f6);
		}
		var v45 := DC31(f5);
		var v46 := false;
		var v47: set<int> := {f7, f7};
		var v48: seq<bool> := [false, v47 < v47, v46, v46 <==> v46];
		var v49: multiset<int> := multiset{p0, f6};
		v45, f7, v46 := v45, |v48|, v49 == (v49 - v49);
		var v50 := new C0(p0);
		var v51 := 'd';
		var v52 := "nom";
		var v53 := DC9(f4, p0, true);
		v0[safeIndex(544, v0.Length)] := v53.cf10;
		var v54: multiset<multiset<int>> := multiset{v49};
		var v55: map<char, multiset<multiset<int>>> := map[v51 := v54];
		f7, v51, v52, v46, v0[safeIndex(544, v0.Length)] := if (v46) then f7 else -0x2e1, v51, "dxypm", v54 > (if (v51 in v55) then v55[v51] else v54), v46;
		var v56: map<bool, int> := map[v0[safeIndex(544, v0.Length)] := v50.f12];
		r0 := v56 + map[!v0[safeIndex(544, v0.Length)] := v50.f12];
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: char, r1: array<D1>) {
		var v0 := new C4(f5, safeDivisionInt(fm7(globalState), f7), !p0);
		var i0 := 0;
		while (p0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v1 := new C8();
			var v2 := "bj";
			var v3: map<int, string> := map[f7 := v2];
			v3 := fm69(globalState);
			var v4: array<int> := new int[22](i1 => safeModuloInt(i1, |[f4, false]|));
			v4 := v4;
			f7 := f6;
		}
		var v5: array<int> := new int[3];
		v5 := v5;
		if (p0) {
			v5[safeIndex(176, v5.Length)] := f6;
			v5[safeIndex(176, v5.Length)] := safeModuloInt(f6, f6);
			var v6 := true;
			var v7: set<bool> := {p0};
			var v8 := 'w';
			v6 := (v7 * v7) > (fm46(v8, globalState) * v7);
			var v9 := new C3(f5, v5[safeIndex(176, v5.Length)], v6);
			var v10: array<seq<int>> := new seq<int>[24];
			v10[safeIndex(799, v10.Length)] := seq(abs(0x2ff), i2  => (v5[safeIndex(176, v5.Length)]));
			var v11: seq<int> := [f7];
			var v12 := "xondmryd";
			var v13: seq<bool> := [p0, f4, f4];
			var v14 := DC27([v13[safeIndex(v5[safeIndex(176, v5.Length)], |v13|) := v6]]);
			var v15: seq<seq<int>> := [v11 + [f6], v11 + fm39(v12, v14, f4, globalState)];
			var v16: multiset<int> := multiset{f6, |"cprbau"|};
			v10[safeIndex(799, v10.Length)] := v15[safeIndex(|v16|, |v15|)];
			var v17: multiset<string> := multiset{v12};
			var v18: map<bool, int> := map[f4 := v5[safeIndex(176, v5.Length)]];
			v17 := v17 - v17[v12 := abs(|v18|)];
		} else {
			var v19: seq<int> := [f6, safeModuloInt(f7, f7), f7];
			v19 := v19 + (v19 + (seq(abs(484), i3  => (f6))));
			var v20 := 'f';
			var v21 := new C5(!!p0, f7, f5, f6, fm0(v20, globalState));
			var v22 := true;
			var v23: set<bool> := {p0};
			v5[safeIndex(617, v5.Length)] := safeDivisionInt(v21.f14, -|v23|);
			var v24: set<int> := {f6, 0x12};
			var v25: multiset<bool> := multiset{v21.f13, v21.f13};
			var v26: map<bool, multiset<bool>> := map[v21.f13 := multiset{p0}];
			v22, v21.f14, v22, v5[safeIndex(617, v5.Length)], v22 := (v24 - v24) >= v24, safeModuloInt(f6, f6 * f7), v25 >= multiset{false, p0}, (f6 * f7) * (f6 - v21.fm7(globalState)), (p0 <== v21.f13) !in v26;
			v22 := p0;
			v5[safeIndex(617, v5.Length)] := 0x2d;
		}
		
		var v27 := "packlnnxk";
		v27 := v27;
		var v28: seq<bool> := [f4];
		var v29 := 'y';
		var v30: multiset<char> := multiset{v29};
		var v31: seq<multiset<char>> := [v30, v30];
		if (fm70(f4, 0x2e0, v28, p0, globalState) <= ([v30] + v31)) {
			f7 := f7;
			var v32: map<int, bool> := map[f7 := f4];
			f7, f7, f7, v28 := f7, if (p0 <== (if (|v27| in v32) then v32[|v27|] else f4)) then f7 else f6, f7, v28;
			var v33 := true;
			v33 := !v33;
			var v34: seq<int> := [f7];
			v5[safeIndex(227, v5.Length)] := |v34| * f6;
			v5[safeIndex(13, v5.Length)] := if (fm8(f7, v33, v27, v34[safeIndex(f7, |v34|) := f7], globalState)) then |v28| else fm3(f6, false, globalState);
			var v35: map<bool, bool> := map[p0 := false];
			var v36: multiset<bool> := multiset{if (v33 in v35) then v35[v33] else v28[safeIndex(f7, |v28|)], v33, !v33};
			v5[safeIndex(227, v5.Length)], v5[safeIndex(13, v5.Length)], f7, v33, v33 := --safeDivisionInt(f6, if (v0.fm8(f7, false, v27, v34, globalState) in v36) then v36[v0.fm8(f7, false, v27, v34, globalState)] else |[p0, p0]|), f7, fm3(f6, v33, globalState), v28[safeIndex(f6, |v28|)], p0;
			var v37: seq<seq<char>> := [v27];
			v33 := v0.fm8(f6, f4, v37[safeIndex(f6, |v37|)], v34, globalState);
		} else {
			var v38: array<bool> := new bool[1];
			var v39 := DC5(v38);
			v39 := v39;
			v38 := v38;
			var v40: array<array<C8>> := new array<C8>[22];
			var v41: seq<array<int>> := [v5, v5];
			var v42: map<array<int>, array<array<C8>>> := map[v41[safeIndex(f7, |v41|)] := v40];
			v40 := if (v41[safeIndex(f6, |v41|)] in v42) then v42[v41[safeIndex(f6, |v41|)]] else v40;
			var v43: map<int, int> := map[-f6 := f7];
			var v44: seq<int> := [f6];
			f7 := safeDivisionInt(0x1e0 + (if (v44[safeIndex(0x385, |v44|)] in v43) then v43[v44[safeIndex(0x385, |v44|)]] else |v28|), 0x1bc);
			var v45: map<array<int>, array<bool>> := map[v5 := v38];
			f7 := |v45|;
		}
		
		r0 := v29;
		r1 := new D1[9];
	}
	method m3(p0: int, p1: int, p2: char, globalState: GlobalState) returns (r0: map<int, int>) {
		var v0: seq<bool> := [true];
		var v1: seq<seq<bool>> := [v0];
		var v2 := DC27(v1);
		var v3: seq<int> := [|fm71(globalState)|, 0x3b9];
		var v4 := DC16(f4, f4 <== f4, v3 + v3);
		var v5 := "wweol";
		var v6: map<bool, bool> := map[false := f4];
		var v7: set<bool> := {false};
		var v8: C5 := new C5(f4, |v7|, f5, f6, !fm0('s', globalState));
		var v9: seq<C5> := [v8];
		var v10: set<C5> := {v8, v9[safeIndex(f7, |v9|)]};
		var v11: multiset<bool> := multiset{!v0[safeIndex(v8.f14, |v0|)], f4, f4};
		var v12: array<int> := new int[15] [f6, |(v6 + v6)|, |{f4, f4, f4, f4}|, p0, safeModuloInt(p0, |v10|), f6, p0, v8.fm9(v8.f13, v8.f13, p1, globalState), p1, v8.fm7(globalState), -(if (false in v11) then v11[false] else p0), p0 - f7, |v3| + p0, |v7|, p0];
		v12[safeIndex(662, v12.Length)] := DC13(|v5|, f4, f7, v0).cf14;
		var v13 := 'r';
		var v14 := DC19("dks");
		v2, v4, v5, v12[safeIndex(662, v12.Length)], v13 := v2, DC16(v8.f13, fm0(fm1(|v5|, f4, globalState), globalState), v3[safeIndex(p0, |v3|) := p1]), (if (v0[safeIndex(-0x8, |v0|)]) then v14 else v14).cf25, |"lvpib"|, v13;
		if (false) {
			var v15 := false;
			var v16: multiset<int> := multiset{|v0|};
			v15 := (v16 - multiset([f7])) >= v16;
			v12[safeIndex(662, v12.Length)] := -p0;
			v12[safeIndex(662, v12.Length)] := f7;
			v12[safeIndex(662, v12.Length)] := v8.f14;
			var v17: array<set<map<bool, int>>> := new set<map<bool, int>>[17];
			var v18: map<bool, int> := map[true := v8.f14];
			var v19: set<map<bool, int>> := {v18};
			v17[safeIndex(990, v17.Length)] := v19;
			v17[safeIndex(990, v17.Length)], v8.f14, v15 := {map[v15 := v8.f14], v18} + {map[f4 := p0], v18, v18, v18}, (f6 - -f7) - p0, v8.f13;
		} else {
			var v20 := true;
			var v21: array<bool> := new bool[21];
			var v22: map<array<bool>, bool> := map[v21 := true];
			v20 := if (v21 in v22) then v22[v21] else v8.f13;
			v20 := v20;
			var v23 := new C5(v8.f13 || v20, f7, f5, f6, v8.f13 && v8.f13);
			v20 := !v23.f13;
			var v24: multiset<seq<int>> := multiset{[0x2ad], v3, v3, v3, fm39(v5, v2, v8.f13, globalState)};
			var v25 := DC52(multiset{v3});
			v24 := v25.cf88;
		}
		
		v12[safeIndex(662, v12.Length)], v12 := v12[safeIndex(662, v12.Length)], v12;
		var v26: map<int, bool> := map[p0 := f4];
		var v27: multiset<array<int>> := multiset{v12};
		var v28: array<bool> := new bool[12](i0 => v8.f13);
		var v29: map<int, int> := map[|v5[safeIndex(|v27|, |v5|) := p2]| := |[v28]|];
		var v30: map<bool, int> := map[false := |v29|];
		var v31: array<bool> := new bool[16] [v8.f13, v8.f13, v8.fm8(p1, v8.f13, v5, v3, globalState), !f4, f4, v8.f13 <==> f4, !fm0(p2, globalState), f4, v8.f13, f4 || DC13(v12[safeIndex(662, v12.Length)], v8.f13, p0, v0[safeIndex(|v0|, |v0|) := v8.f13]).cf15, v8.f13, v11 != multiset{if (|v30| in v26) then v26[|v30|] else v8.f13, false}, !((seq(abs(-0x1a4), i1  => (v13))) < v5), v8.f13, f4, f4];
		v28 := v28;
		var v32 := true;
		v32 := false;
		if (f4) {
			v8.f14 := safeDivisionInt(v8.f14, -v12[safeIndex(662, v12.Length)]);
			v32 := v8.f13 || (multiset{|v0|, f7} !! multiset{p0});
			v28[safeIndex(467, v28.Length)] := v8.f13;
			v28[safeIndex(467, v28.Length)] := v8.f13;
			var v33 := new C0(v8.f14);
			v12[safeIndex(662, v12.Length)] := v33.f12;
		} else {
			v8.f14 := v12[safeIndex(662, v12.Length)];
			v6 := v6[f4 := v8.f13];
			v8.f14 := v8.f14 * 0xf8;
			v12[safeIndex(662, v12.Length)] := p0;
			v8.f14 := f6;
		}
		
		r0 := (map[f7 := f6])[|map[v8.f14 := |v5|]| := p0] + v29;
	}
}

function fm0(p0: char, globalState: GlobalState): bool {
	(multiset{false} + multiset([false, false, false, false, false])) !in (seq(abs(0x394), i0  => (multiset([true]))))
}
function fm1(p0: int, p1: bool, globalState: GlobalState): char {
	'm'
}
function fm2(globalState: GlobalState): D0 {
	match if (!!!false) then DC52(multiset{seq(abs(0x2c5), i0  => (-|{|{-0x186}|}|)), [|(map v0 : D3 | v0 in (set v1 : D3 | v1 in map[DC8([false]) := true] :: (v1)) :: (v0) := (true))|, 0x78], [0x31c], [0x7e], [0x11d, |"ngfivc"|]}) else DC52(multiset([[-0x217, -0x35e], [0x239]])) {
		case DC52(cf88) => DC1()
	}
}
function fm3(p0: int, p1: bool, globalState: GlobalState): int {
	-|multiset{-560 * |[|[-0x191]|]|}|
}
function fm4(p0: bool, p1: bool, globalState: GlobalState): seq<multiset<int>> {
	(seq(abs(0x2c6), i0  => (multiset{|['m']|, -0x1d, -|[map[true := false]]|, |map[-0x2b3 := |multiset{true, !true}|]|, -0x9}))) + (if (true) then seq(abs(0x13a), i1  => (multiset{|[[true, false, false, true]]|})) else [multiset{|map[|map[967 := 0x1c0]| := |(map v0 : int | (366 <= v0) && (v0 < 0xe4) :: (v0 - |DC33(false, |map[false := 734]|, multiset{false}, -|"i"|, {true}).cf54|) := ('y'))|]|}])
}
function fm5(p0: int, p1: bool, globalState: GlobalState): map<D0, int> {
	(map[DC1() := 603] + map[DC1() := 0x9d]) + ((map v0 : D0 | v0 in {DC1()} :: (v0) := (0x22e)) + map[DC1() := |multiset{'c', 'e', 'u'}|])
}
function fm10(p0: bool, p1: bool, p2: map<int, seq<char>>, p3: int, globalState: GlobalState): D1 {
	DC4(|map[[-0x25e, |[false]|] := |[|{0xa7}|, 0x349, 0x38d]|]|, -(|map[464 := DC43(true, false, [!false]).cf70]| * 159))
}
function fm16(p0: bool, p1: bool, p2: bool, p3: D2, globalState: GlobalState): string {
	("eh" + "rgorcc") + "oxshrry"
}
function fm18(p0: string, p1: int, globalState: GlobalState): map<int, set<bool>> {
	map[safeModuloInt(|['e', 'n']|, 0x1c) := {true} * {false}]
}
function fm19(p0: bool, p1: int, globalState: GlobalState): seq<bool> {
	[false] + [false, false, false, false]
}
function fm20(p0: int, p1: bool, globalState: GlobalState): D3 {
	match DC15(multiset{[false]}) {
		case DC16(cf20, cf21, cf22) => DC8([false, cf20])
		case DC15(cf19) => if (false) then DC8([true]) else DC8([!false, true])
	}
}
function fm21(p0: int, p1: multiset<map<bool, int>>, p2: int, p3: int, globalState: GlobalState): multiset<bool> {
	match DC38(false, !true) {
		case DC38(cf63, cf64) => multiset{cf64, cf64, cf63, cf64}
		case DC37(cf62) => DC33(true, |[|map[true := 0xfa]|]|, multiset{true}, 0x20f, {!true}).cf54
		case DC39(cf65) => multiset{false, true} + multiset{true, true}
	}
}
function fm22(globalState: GlobalState): map<multiset<int>, string> {
	map[multiset{|multiset{true}|} := "n"] + (map[multiset{0x357, |"xt"|} := seq(abs(0x122), i0  => ('w'))] + map[multiset{-|map[false := false]|} := "pmrrdi"])
}
function fm23(p0: int, p1: int, globalState: GlobalState): set<bool> {
	{false}
}
function fm24(p0: int, p1: char, p2: int, globalState: GlobalState): multiset<seq<bool>> {
	(multiset{[!!true], [true]} + multiset{[false]}) - multiset{[true]}
}
function fm25(p0: int, p1: int, p2: bool, p3: bool, globalState: GlobalState): seq<D0> {
	[DC1(), DC1()] + ([DC1(), DC1(), DC1(), DC1()] + [DC1()])
}
function fm27(p0: bool, p1: int, p2: multiset<bool>, p3: bool, globalState: GlobalState): multiset<seq<bool>> {
	multiset{[false]} + (multiset{[false, !true, !true, false], [false, true, true]} - multiset{[true]})
}
function fm28(globalState: GlobalState): map<bool, int> {
	if (0x3be == 0x18d) then map[false := -268] else map[!false := 0x184]
}
function fm29(globalState: GlobalState): map<int, int> {
	map[0x1c4 := 0x35f] + map[|(map v0 : seq<bool> | v0 in {[true, false], [true, true, false, !true, !true], [true, true, false]} :: (v0) := (!true))| := -0x162]
}
function fm30(p0: seq<int>, p1: bool, globalState: GlobalState): seq<seq<int>> {
	[[0x32c], [-0x3c8, |[false, true]|]] + (seq(abs(573), i0  => ([|map[true := false]|, |[true, true]|, |(map v0 : int | (122 <= v0) && (v0 < -0xf) :: (v0 * -0x3d1) := (true))|])))
}
function fm31(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState): set<int> {
	DC17({-0x1f9, |{-506}|}).cf23
}
function fm32(p0: int, p1: bool, globalState: GlobalState): D7 {
	DC17({|{|multiset{|[true]|, |[false]|}|, |map[!false := "vyojj"]|}|} * {0x360})
}
function fm33(p0: bool, p1: int, p2: bool, globalState: GlobalState): D6 {
	DC15(multiset{[true, true]} + multiset{[false, true]})
}
function fm34(p0: int, p1: int, p2: bool, p3: bool, globalState: GlobalState): map<multiset<int>, string> {
	((map v0 : multiset<int> | v0 in [multiset{|multiset{true}|, 0xec}] :: (v0) := ("e")) + map[multiset([|{299, 0x3a7, -557, |"chaced"|, 0x3a6}|, 0x177]) := "jnbkik"]) + (if (false) then map[multiset(seq(abs(0x157), i0  => (|multiset{true}|))) := "wdpr"] else map[multiset{|map[seq(abs(0x19), i1  => (|[true]|)) := -972]|} := seq(abs(0x14b), i2  => ('k'))])
}
function fm35(p0: int, globalState: GlobalState): set<int> {
	({|"tghobrq"|} * {-0x32f}) * ({--|{|{false, true}|, |"d"|}|, 0xba} + {0x35a, -0x20e, DC24(!true, true, -|"tfoksstqw"|).cf38, |(seq(abs(653), i0  => ('d')))|})
}
function fm36(p0: bool, globalState: GlobalState): seq<int> {
	([0x23] + [|multiset{!!false, true}|]) + (seq(abs(0x396), i0  => (0x387)))
}
function fm37(p0: bool, p1: bool, p2: int, p3: bool, globalState: GlobalState): string {
	"gcdnnimh"
}
function fm38(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): D0 {
	if (0x1d8 < |{|"ky"|}|) then DC0(-|[|"ys"|]|) else if (!true) then DC0(0x26d) else DC0(0x334)
}
function fm39(p0: string, p1: D12, p2: bool, globalState: GlobalState): seq<int> {
	[safeModuloInt(|["tjnq", seq(abs(462), i0  => ('n'))]|, |map[multiset{false} := 'd']|)]
}
function fm40(p0: bool, p1: bool, p2: int, globalState: GlobalState): seq<bool> {
	([true] + [true]) + ([false] + [false])
}
function fm41(p0: int, p1: char, p2: bool, globalState: GlobalState): set<int> {
	({|map[true := |[multiset{|['y', 'a']|}]|]|, |[false, true]|, -0xc1} - (set v0 : int | (0x2fb <= v0) && (v0 < 0x45) :: (v0 - 0x100))) * ((set v2 : int | v2 in (map v1 : int | (0x120 <= v1) && (v1 < 859) :: (v1 + 0x254) := (|{false, false, true}|)) :: (v2 + -500)) + {|[true, true, false]|})
}
function fm42(p0: char, p1: set<D6>, globalState: GlobalState): map<int, int> {
	map[459 := --980] + (map v0 : int | (0x245 <= v0) && (v0 < 197) :: (safeDivisionInt(v0, 0x1f5)) := (|multiset([!!!!!false, false])|))
}
function fm43(globalState: GlobalState): D6 {
	DC16("a" <= "exmbhmdo", true ==> true, [-0x15a] + [|"onku"|])
}
function fm44(globalState: GlobalState): map<bool, int> {
	map[false := 0x1bd]
}
function fm45(p0: bool, p1: int, p2: bool, p3: int, globalState: GlobalState): seq<multiset<bool>> {
	[multiset{true, false}, multiset{true, false, false, true, true}, multiset{false}] + (if (!true) then [multiset{false}, multiset{true}] else seq(abs(824), i0  => (multiset{!false})))
}
function fm46(p0: char, globalState: GlobalState): set<bool> {
	if (true) then {false} else DC33(false, |[|{false}|]|, multiset{false, true}, 0x369, {!true, true}).cf56
}
function fm47(p0: int, globalState: GlobalState): multiset<bool> {
	multiset([true] + (if (true) then [false, true] else [!true]))
}
function fm49(p0: string, p1: seq<bool>, p2: int, globalState: GlobalState): seq<seq<int>> {
	([[191], [|{0x2e1, -741}|, 0x158]] + (seq(abs(-224), i0  => ([|[|map[0x25e := DC16(true, false, [|map[0x240 := |map[true := false]|]|]).cf20]|]|])))) + [[|[false, true]|]]
}
function fm50(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState): string {
	match DC3(!!true) {
		case DC4(cf3, cf4) => "dyg"
		case DC3(cf2) => "dstoi"
	}
}
function fm51(p0: map<bool, set<bool>>, p1: bool, globalState: GlobalState): map<int, bool> {
	map[0x37b := false] + map[|[0x174]| := false]
}
function fm52(p0: int, globalState: GlobalState): D3 {
	DC8(if (false) then [DC33(false, -0x261, multiset([true, true, false, false]), -0x1c6, {false}).cf52] else [false])
}
function fm53(p0: int, p1: seq<map<bool, bool>>, globalState: GlobalState): set<bool> {
	{true, false} - ({false, false} * {true})
}
function fm54(p0: int, globalState: GlobalState): seq<map<bool, bool>> {
	[map[true := false]] + ([map[false := true]] + (seq(abs(0x9b), i0  => (map[true := !false]))))
}
function fm55(p0: bool, p1: bool, p2: int, globalState: GlobalState): D1 {
	match DC43(false, true, [!!true]) {
		case DC43(cf69, cf70, cf71) => DC3(false)
		case DC44(cf72) => DC3(false)
		case DC42(cf68) => DC3(true)
	}
}
function fm56(p0: bool, p1: int, p2: bool, globalState: GlobalState): seq<bool> {
	if (false && !true) then [false] else [false] + [!false]
}
function fm57(p0: int, p1: int, globalState: GlobalState): seq<int> {
	((seq(abs(0xbd), i0  => (|[map[false := 681]]|))) + [-469]) + (seq(abs(-0x295), i1  => (|{map['b' := false], map['x' := false]}|)))
}
function fm58(globalState: GlobalState): map<bool, D12> {
	(map[false := DC27([[true]])] + map[true := DC27([[true, false]])]) + map[true := DC27([[false], [true]])]
}
function fm59(globalState: GlobalState): map<D10, int> {
	match DC43(true, false, [false, false]) {
		case DC43(cf69, cf70, cf71) => map v0 : D10 | v0 in [DC21(map[cf69 := 804])] :: (v0) := (|"sg"|)
		case DC44(cf72) => map[DC21(map[cf72 := 0x12]) := --|"c"|]
		case DC42(cf68) => map[DC21(map[true := |map[false := 2]|]) := -|[0x308]|] + map[DC21(map[false := 0x14f]) := |{0x179, --|multiset{0x2fa, -0x399}|}|]
	}
}
function fm60(p0: int, p1: int, p2: int, globalState: GlobalState): map<seq<bool>, map<int, bool>> {
	map[[false, false] + [false] := map[|map[909 := 0x170]| := !!!false] + (map v0 : int | v0 in multiset{896} :: (v0 - 75) := (true))]
}
function fm61(p0: set<bool>, globalState: GlobalState): seq<seq<bool>> {
	[[true, !true, false, true, false]] + [[!!false, true, true, false, !false]]
}
function fm62(globalState: GlobalState): D4 {
	DC12(if (!true) then 'p' else 'a')
}
function fm63(globalState: GlobalState): D4 {
	DC13(|(map[!false := |"hxnhw"|] + map[false := |"aog"|])|, false <==> false, 0x352, [!false])
}
function fm64(p0: int, p1: int, p2: bool, globalState: GlobalState): set<char> {
	({'e'} + {'m', 'm', 'x'}) + ((set v0 : char | v0 in multiset{'f'} :: (v0)) - {'a'})
}
function fm65(p0: int, p1: int, p2: int, globalState: GlobalState): map<bool, char> {
	(map[true := 'f'] + map[false := 'r']) + (map[true := 'q'] + map[!false := 'j'])
}
function fm66(p0: int, globalState: GlobalState): D16 {
	DC33(true ==> false, safeModuloInt(|"gk"|, 925), DC33(false, |[false]|, multiset{!false}, 0xf9, {false}).cf54 * multiset{true}, -0x98, {false} - {false})
}
function fm67(p0: bool, globalState: GlobalState): map<map<int, int>, bool> {
	if (true) then map[map[876 := 0x35d] := !!true] else map[map[|{0x2d0, 0x224}| := |multiset{"rss", "s"}|] := true] + map[map[|"udygyb"| := -0x375] := !true]
}
function fm68(globalState: GlobalState): D12 {
	DC27(if (false) then seq(abs(0xdf), i0  => ([true])) else [[false]])
}
function fm69(globalState: GlobalState): map<int, string> {
	map[|(multiset{--3} + multiset{-0x3d6})| := "igvytirrh" + (seq(abs(0x266), i0  => ('m')))]
}
function fm70(p0: bool, p1: int, p2: seq<bool>, p3: bool, globalState: GlobalState): seq<multiset<char>> {
	((seq(abs(0x7), i0  => (multiset{'d'}))) + (seq(abs(866), i1  => (multiset{'j', 'q'})))) + [multiset{'k'}, multiset{'r', 'c', 'f', 'b'}, multiset(['q']), multiset{'b', 'c'}, multiset{'h'}]
}
function fm71(globalState: GlobalState): map<bool, bool> {
	map[true := true] + map[!false := false]
}
function fm72(p0: map<string, D12>, p1: multiset<int>, globalState: GlobalState): D25 {
	match DC17({0x15e, |[true, !true, true]|, |(seq(abs(0x34e), i0  => ('m')))|}) {
		case DC17(cf23) => DC52(multiset(seq(abs(510), i1  => ([0xb4]))))
	}
}
function fm73(globalState: GlobalState): multiset<seq<int>> {
	(multiset{[|(map v0 : int | v0 in map[0x220 := |"qkivckp"|] :: (safeModuloInt(v0, 0x2b0)) := (false))|], [|DC58(map[map[false := true] := !false]).cf97|]} * multiset([[0x2dc]])) + multiset{seq(abs(0x207), i0  => (|map[0x25a := -700]|))}
}
method m0(p0: map<multiset<int>, string>, p1: bool, p2: multiset<int>, p3: bool, globalState: GlobalState) returns (r0: bool, r1: int) {
	if (true) {
		var v0: array<seq<int>> := new seq<int>[5];
		var v1: map<bool, array<seq<int>>> := map[p1 := v0];
		v0 := if (p3 in v1) then v1[p3] else v0;
		var v3 := 0x32;
		var v4: seq<int> := [v3, v3, v3];
		var v5: set<int> := {v3};
		var v6: multiset<bool> := multiset{p3};
		var v7: set<bool> := {p3, p1, p3, p1, p3};
		var v8: map<bool, bool> := map[p3 := p3];
		var v9: seq<map<bool, bool>> := [v8, map[p1 := p3], map[p1 := p1], v8];
		var v10: array<bool> := new bool[19] [p3, p3, p1, |(map v2 : int | v2 in v4 :: (safeDivisionInt(v2, 0x58)) := (-v3))| != 0x323, p1, p1, v5 > v5, p1, p1, v3 <= v3, v6 !! v6, v7 < fm53(v3, v9, globalState), p1, p1, p1, p1, p1 ==> p3, p3, p3];
		var v11: seq<bool> := [true];
		v10[safeIndex(167, v10.Length)] := v11[safeIndex(v3, |v11|)];
		var v12 := "fku";
		v10[safeIndex(167, v10.Length)] := v5 < {|v12|, v3};
		var v13: array<array<bool>> := new array<bool>[10];
		var v14: C4 := new C4(v13, --v3, v10[safeIndex(167, v10.Length)]);
		v14 := v14;
		var v15: array<int> := new int[5];
		v15[safeIndex(173, v15.Length)] := |v6| - v3;
		v15[safeIndex(173, v15.Length)] := ((if (|v7| in p2) then p2[|v7|] else |v12|) + v3) + v3;
		var v16: C2 := new C2(v12, v11[safeIndex(-v15[safeIndex(173, v15.Length)], |v11|)]);
		var v17 := DC56(v16);
		v16 := v17.cf95;
	} else {
		var v18 := 'r';
		var v19: array<bool> := new bool[8] [p1, fm0(v18, globalState), p3, p3, p1, false, p3, p1];
		var v20: map<bool, array<bool>> := map[p3 := v19];
		var v21: array<array<bool>> := new array<bool>[22] [v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, v19, if (p3 in v20) then v20[p3] else v19, v19];
		var v22 := 0x354;
		var v23: T1 := new C4(v21, v22, p1);
		var v24: set<T1> := {v23};
		var v25 := DC28(|v24|, v22, v23.f4, v23.f4, safeDivisionInt(|"tbi"|, v23.f6));
		v25 := v25;
		var v26 := "filnlnc";
		v26 := v26;
		var v27: seq<int> := [v23.f6, v22];
		var v28: multiset<bool> := multiset{p3};
		var v29: array<seq<int>> := new seq<int>[22] [v27, v27, v27, [|v26|], v27 + v27, v27[safeIndex(|v26|, |v27|) := v23.f6], v27 + v27, v27 + v27, v27, v27, [if (p1 in v28) then v28[p1] else v22], v27 + [v27[safeIndex(v23.f6, |v27|)], |v26|, v22, if (false in v28) then v28[false] else |p2|, v22], [v22, |v26|], [v23.f6] + v27, v27, v27, v27, v27, v27, v27, v27, fm57(v23.f6, v23.f6, globalState)];
		var v30: map<bool, int> := map[p1 := v22];
		var v31 := DC21(v30[p3 := v23.f6]);
		var v32: seq<D10> := [v31];
		var v33: seq<bool> := [fm0(v18, globalState), p3];
		r0, v29, v32, r1 := !v33[safeIndex(v22, |v33|)], v29, if (false) then v32 else v32[safeIndex(v23.f6, |v32|) := v31], v23.fm7(globalState);
		var v34 := DC33(v23.f4, v23.fm9(p3, p3, v23.f6, globalState), v28, |v26|, {v23.f4, p1});
		r0 := !(v34.cf52 <== (v28 >= v34.cf54));
		var v35: map<int, bool> := map[v22 := p3];
		var v36: map<char, seq<bool>> := map[v18 := v33];
		var v37: C1 := new C1(-safeDivisionInt(185, if ((if (0x374 in v35) then v35[0x374] else p1) in v28) then v28[if (0x374 in v35) then v35[0x374] else p1] else v22), v22, if (p1) then v21 else v23.f5, |v36|, p2 < p2);
		v37, v27, v19 := v37, v27 + (seq(abs(0x1dc), i0  => (-v23.f6))), v19;
	}
	
	var v38 := 197;
	for i1 := v38 to fm3(v38, true, globalState) {
		v38 := safeModuloInt(v38, v38);
		var v39: seq<int> := [v38];
		var v40 := "etegojdr";
		r0 := if (false) then p3 else v39 != v39[safeIndex(|v40|, |v39|) := i1];
		var v41: C2 := new C2(v40 + v40, p3);
		v41 := if (p1) then v41 else v41;
		var v42: seq<bool> := [!p1];
		var v43: set<seq<bool>> := {v42, v42, v42, v42};
		var v44: seq<seq<bool>> := [v42, v42];
		var v46: array<array<bool>> := new array<bool>[2];
		var v47 := new C9(i1, v43 - (set v45 : seq<bool> | v45 in v44 :: (v45)), v46, i1, p3);
	}
	var v48: array<int> := new int[4](i2 => i2 - |map['j' := 33]|);
	var v49: seq<array<int>> := [v48, v48, v48, v48];
	v49 := v49 + v49;
	v38 := safeDivisionInt(fm3(v38, p1, globalState), v38);
	if (p1) {
		v48[safeIndex(171, v48.Length)] := 946;
		v48[safeIndex(171, v48.Length)] := -0x1cf;
		var v50: array<char> := new char[4](i3 => 'l');
		var v51 := 'j';
		v50[safeIndex(838, v50.Length)] := v51;
		v50[safeIndex(838, v50.Length)], r0, r1, r0 := v51, fm0(v51, globalState) <== p1, v48[safeIndex(171, v48.Length)], !!false <== p3;
		v50[safeIndex(838, v50.Length)] := 'b';
		var v52: seq<bool> := [p1];
		var v53 := "bjpsw";
		var v54: set<seq<bool>> := {v52, ([false, p3])[safeIndex(|v53|, |[false, p3]|) := p3]};
		var v55 := DC24(p1, p3, v48[safeIndex(171, v48.Length)]);
		var v56 := DC3(p1);
		var v57: array<bool> := new bool[27] [p3, p3, p3, p1, true, v55.cf36, p1, p3, p3, v56.cf2, !p3, p3, p1, p3, p3, p3, p3, p1, p1, !p1, p3, false, p1, p3, !p1, true, p1];
		var v58: array<array<bool>> := new array<bool>[8] [v57, v57, v57, v57, v57, v57, v57, v57];
		var v59: seq<array<array<bool>>> := [v58];
		var v60: set<int> := {v38};
		var v61 := DC28(|v60|, v38, p3, p3, 0x204);
		var v62: C9 := new C9(|(seq(abs(544), i4  => (v50[safeIndex(838, v50.Length)])))| + v48[safeIndex(171, v48.Length)], v54, v59[safeIndex(v38, |v59|)], v48[safeIndex(171, v48.Length)], (v61.(cf45 := true, cf44 := v38)).cf46);
		var v63: seq<int> := [v62.f7];
		var v64: seq<int> := [v48[safeIndex(171, v48.Length)], v62.f7, -v62.f7, v62.f7, v63[safeIndex(0x153, |v63|)]];
		var v65 := DC20(p3, v50[safeIndex(838, v50.Length)], v62.f7);
		var v66: array<int> := new int[13] [v38, safeModuloInt(|map[975 := false]|, v62.f7), v62.f7, v38, v62.f7, -v38, |(v63 + (seq(abs(0x2d9), i5  => (v48[safeIndex(171, v48.Length)]))))|, v48[safeIndex(171, v48.Length)], fm3(v65.cf28, p3, globalState), v62.f7, v48[safeIndex(171, v48.Length)], if (p1) then v38 else 0x2cb, v62.f7];
		v62, v48[safeIndex(171, v48.Length)], v66, r0 := v62, fm3(v38, p3, globalState), v66, p1;
		var v67: C6 := new C6(p3, p3, v58, -356);
		var v68: set<C6> := {v67};
		var v69: map<set<C6>, bool> := map[v68 := p1];
		v48[safeIndex(171, v48.Length)], v38, v48[safeIndex(171, v48.Length)], r1 := v38 * fm3(v38, false, globalState), v38 - v48[safeIndex(171, v48.Length)], v48[safeIndex(171, v48.Length)] + safeModuloInt(v62.f7, v62.f7), (if (!p1) then |v69| else -v48[safeIndex(171, v48.Length)]) + (v38 + v62.f7);
	} else {
		var v70 := 'g';
		var v71: map<char, int> := map[v70 := v38];
		var v72: seq<int> := [v38, v38, if (v70 in v71) then v71[v70] else v38, |(seq(abs(0x21c), i6  => (v38)))|];
		var v73 := DC4(v38, v38);
		var v74: map<bool, int> := map[p3 := v38];
		var v75: set<int> := {safeDivisionInt(|v72|, v38), v38, v73.cf4, if (p1 in v74) then v74[p1] else v38};
		v75 := (if (!p1) then v75 else v75) - (v75 * v75);
		var v76: map<bool, bool> := map[p1 := p1];
		var v77: multiset<bool> := multiset{p3, p3};
		var v78 := DC24(fm0(v70, globalState), if (false in v76) then v76[false] else p3, |v77|);
		var v79: multiset<D11> := multiset{v78};
		var v80: array<bool> := new bool[26] [p3, p3, p1, p1, p1, p1, p3, fm0(v70, globalState), !false, p3, p3, p3, p3, !p1, p1, !p3, true, p3, false, p1, p3, p3, p1, p1, p3, p1];
		var v81: array<array<bool>> := new array<bool>[26] [v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80, v80];
		var v82: C5 := new C5(p3 <==> p3, |v79|, v81, v38, p3);
		v82 := v82;
		var v83: array<seq<T0>> := new seq<T0>[21];
		var v84: map<array<seq<T0>>, bool> := map[v83 := v82.f13];
		v84 := v84[v83 := p3];
		var v85: seq<bool> := [p1];
		var v86: set<seq<bool>> := {v85};
		var v87: C9 := new C9(-|(seq(abs(-613), i7  => (v70)))|, v86, v81, v82.f14, p1);
		var v88: map<C9, bool> := map[v87 := p3];
		v88 := v88[if (v82.f13) then v87 else v87 := p1];
		var v89 := "xu";
		var v90: T0 := new C2(v89, p1);
		var v91 := DC51(v90);
		v91 := v91;
	}
	
	if (p3) {
		var v92 := 'k';
		var v93 := "hvok";
		var v94 := DC19(v93);
		var v95: T0 := new C2(seq(abs(-0x153), i8  => ('o')), v92 in v94.cf25);
		v95, v38 := v95, -v38;
		var v96: multiset<bool> := multiset{v95.f4, p3};
		var v97: map<bool, bool> := map[p3 := p1];
		var v98: seq<map<bool, bool>> := [v97, v97];
		match (DC33(v95.f4, v38, v96 - v96, v38, fm53(v38, v98, globalState))) {
			case DC33(cf52, cf53, cf54, cf55, cf56) =>
				var v99: array<C4> := new C4[20];
				var v100: array<array<bool>> := new array<bool>[14];
				var v101: seq<int> := [cf53];
				var v102: C4 := new C4(v100, v101[safeIndex(|v101|, |v101|)], p3);
				v99[safeIndex(60, v99.Length)] := v102;
				var v103: seq<bool> := [!p3];
				var v106: multiset<map<int, bool>> := multiset{map v104 : int | (438 <= v104) && (v104 < 0x187) :: (safeDivisionInt(v104, |v97|)) := (p1), map v105 : int | (0x163 <= v105) && (v105 < 267) :: (v105 * v38) := (p1)};
				r1, r0, v92, v99[safeIndex(60, v99.Length)] := cf53, v103[safeIndex(|v106|, |v103|)], if (v102.fm8(cf53, cf52, seq(abs(0x107), i9  => (v92)), v101, globalState)) then v92 else 'p', v102;
				var v107: array<C9> := new C9[3];
				var v108: set<seq<bool>> := {v103, v103, [v95.f4], v103, v103};
				var v109: C9 := new C9(cf53, v108, v100, cf55, cf52);
				v107[safeIndex(913, v107.Length)] := v109;
				v107[safeIndex(913, v107.Length)] := v109;
				cf52 := !false;
				cf52 := -v102.fm7(globalState) < cf53;
			case DC32(cf51) =>
				v92 := if (if (p1) then v95.f4 else v95.f4) then v92 else fm1(v38, v95.f4, globalState);
				var v110: array<bool> := new bool[20];
				var v111: seq<int> := [|{v110, v110, v110, v110, v110}|];
				r0 := v111[safeIndex(v38, |v111|)] > -(if (v95.f4) then fm3(v38, v95.f4, globalState) else v38);
				var v112: array<set<set<D15>>> := new set<set<D15>>[19];
				var v113: array<array<bool>> := new array<bool>[9];
				var v114 := DC31(v113);
				var v115: set<D15> := {v114, v114};
				var v116: set<set<D15>> := {{v114}, v115};
				v112[safeIndex(506, v112.Length)] := v116;
				v112[safeIndex(506, v112.Length)] := {v115};
				r1 := -v38;
		}
		
		var v117: map<int, array<int>> := map[v38 := v48];
		v117 := v117[v38 := v48];
		v38 := v38;
		r0 := p1;
	} else {
		r0 := |(seq(abs(863), i10  => ('l')))| >= v38;
		var v118: array<map<bool, D11>> := new map<bool, D11>[1];
		var v119 := DC45(v118);
		v119 := DC45(v118);
		var v120: map<int, bool> := map[v38 := p3];
		var v122: array<array<bool>> := new array<bool>[22];
		var v123: seq<int> := [v38, v38];
		var v124 := new C1(|(set v121 : int | v121 in v120 :: (v121 + |{true, true}|))|, v38, v122, |v123|, p1);
		v48[safeIndex(240, v48.Length)] := v38;
		v48[safeIndex(240, v48.Length)] := v38;
		v38 := |(seq(abs(0x26f), i11  => ('b')))|;
	}
	
	var v125 := "rjhatt";
	var v126: array<bool> := new bool[22];
	var v127: seq<array<bool>> := [v126, v126];
	var v128 := DC41(v38);
	var v129 := DC22(v125, p1, p3, v127[safeIndex(v38, |v127|)], v128.cf67);
	var v130: map<int, bool> := map[v38 := p3];
	var v131: map<bool, bool> := map[p1 := p1];
	var v132: array<array<bool>> := new array<bool>[24];
	var v133: C6 := new C6(if (v38 in v130) then v130[v38] else p1, !(if (p1 in v131) then v131[p1] else p1), v132, fm3(v38, false, globalState));
	var v134: map<C6, int> := map[v133 := v38];
	r0 := |v129.cf30| >= |v134|;
	var v135 := 't';
	r1 := |(("ewmfdfes")[safeIndex(|(seq(abs(0x2ca), i12  => ('r')))|, |"ewmfdfes"|) := v135] + "nv")|;
}
method Main() {
	var v0 := false;
	var v1: multiset<bool> := multiset{v0, v0, v0};
	var globalState := new GlobalState(v1, -0x3e1, false, 0x35e);
	var v2 := 598;
	var v3: array<int> := new int[14] [v2, v2, v2, if (v0) then v2 else v2, safeDivisionInt(|(seq(abs(532), i0  => ('m')))|, 0x14e), -0x3a2, v2, v2, v2, DC2(v2).cf1 - -v2, v2, v2 - v2, --v2 * 743, v2];
	v3 := v3;
	var v4 := 'j';
	v4 := v4;
	var v5: map<int, bool> := map[v2 * v2 := v0];
	v5 := v5[v2 - v2 := v0];
	var v6: set<int> := {v2, v2};
	var v7: set<set<int>> := {v6};
	var v8 := "smw";
	var v9: map<set<set<int>>, bool> := map[v7 := |v8| != v2];
	v0 := if (v7 in v9) then v9[v7] else v0;
	var v10: array<set<bool>> := new set<bool>[2];
	forall i1 | 0 <= i1 < v10.Length {
		v10[i1] := ({v0} - {v0, false}) - ({v0, true, v0, false} * {v0, v0, DC3(v0).cf2, v0, v0});
	}
	v0 := v0 <== v0;
	var v11: array<bool> := new bool[6](i2 => v0);
	var v12: seq<array<bool>> := [v11, v11];
	var v13: map<bool, char> := map[!v0 := v4];
	var v14: seq<int> := [v2];
	var v15: array<bool> := new bool[21] [!false, fm0(fm1(-v2, v0, globalState), globalState), v0, v0, !fm0('e', globalState), v0, v0, !v0, v0, !v0, v0, v11 in v12, v0, v0, v0, v2 != (if (v0 in v1) then v1[v0] else |v8|), if (v0) then v0 else !v0, v0, v0, !fm0(if (v0 in v13) then v13[v0] else v4, globalState), |v8| < |v14|];
	v15[safeIndex(99, v15.Length)] := false;
	var v16 := DC1();
	var v17: seq<D0> := [v16, fm2(globalState)];
	v0, v15[safeIndex(99, v15.Length)] := !([DC1()] == (v17 + v17)), fm0(v4, globalState);
	var v18: multiset<int> := multiset{v2};
	var v19: map<multiset<int>, string> := map[v18 := v8];
	var v20, v21 := m0(v19, if (v15[safeIndex(99, v15.Length)]) then !v0 else v15[safeIndex(99, v15.Length)], v18[v2 := abs(v2)], v15[safeIndex(99, v15.Length)], globalState);
	var v22: array<seq<int>> := new seq<int>[24](i3 => v14 + v14);
	v22[safeIndex(766, v22.Length)] := v14;
	var v23 := DC2(v2);
	v16, v15[safeIndex(99, v15.Length)], v20, v22[safeIndex(766, v22.Length)], v2 := match v23 {
		case DC1() => v16
		case DC2(cf1) => v16
		case DC0(cf0) => v16
	}, true, fm0(v4, globalState), v14, v2;
	v23 := v23;
	if (v15[safeIndex(99, v15.Length)]) {
		var v24: map<int, array<int>> := map[v2 := v3];
		v3 := if (v2 in v24) then v24[v2] else v3;
		v3[safeIndex(124, v3.Length)] := v21 * fm3(v2, v15[safeIndex(99, v15.Length)], globalState);
		var v25: seq<multiset<int>> := [v18, v18 + multiset{v2, 0x2a1}];
		v20, v3[safeIndex(124, v3.Length)], v25, v23 := !v0, safeDivisionInt(v21, v2), v25 + fm4(false, true, globalState), v23;
		v3[safeIndex(124, v3.Length)] := 0x386;
		v15[safeIndex(99, v15.Length)] := v20;
		var v26 := DC5(v15);
		v11 := v26.cf5;
	} else {
		v3[safeIndex(547, v3.Length)] := |(if (v20) then v8 else v8)|;
		var v27: array<string> := new string[16];
		v27[safeIndex(397, v27.Length)] := "nwenxedn";
		var v28: seq<string> := [v8, v8, "qystgqh", seq(abs(0x27b), i4  => (v4)), v8 + v8];
		v3[safeIndex(547, v3.Length)], v27[safeIndex(397, v27.Length)] := |v28|, seq(abs(912), i5  => ('w'));
		var v30: seq<multiset<int>> := [v18];
		var v31: map<bool, map<multiset<int>, string>> := map[true := v19];
		var v32, v33 := m0((map v29 : multiset<int> | v29 in v30 :: (v29) := (v8)) + (if (v20 in v31) then v31[v20] else v19), v0, v18, !v0, globalState);
		v0 := v15[safeIndex(99, v15.Length)];
		var v34: set<string> := {v8};
		var v35: map<bool, bool> := map[v0 := v32];
		v21, v18, v3[safeIndex(547, v3.Length)], v3[safeIndex(547, v3.Length)], v3[safeIndex(547, v3.Length)] := fm3(|v34| + v33, !v15[safeIndex(99, v15.Length)], globalState), v18, v21, -safeModuloInt(v2, -(if (v20) then v33 else |v35|)), v33;
		v0 := v32 <== v32;
	}
	
	var v36: map<D0, int> := map[v16 := v21];
	var v40: set<bool> := {fm0(v4, globalState)};
	var v41: array<map<D0, int>> := new map<D0, int>[12] [v36, v36 + v36, v36, v36, map v37 : D0 | v37 in v36 :: (v37) := (v14[safeIndex(v2, |v14|)]), v36, v36, v36, (map v38 : D0 | v38 in [v16, v16] :: (v38) := (|(map v39 : D2 | v39 in multiset([DC7(v2)]) :: (v39) := (v21))|))[DC1() := fm3(|v40|, v0, globalState)], v36, fm5(v21, v15[safeIndex(99, v15.Length)], globalState), v36];
	v41 := if (!v0) then v41 else v41;
	var v42: multiset<char> := multiset{v4};
	var v43, v44 := m0(v19, false, v18, 'n' in v42, globalState);
	v3[safeIndex(964, v3.Length)] := v44;
	var v45: map<array<int>, seq<int>> := map[v3 := seq(abs(-0x2ee), i6  => (v21))];
	v3[safeIndex(964, v3.Length)] := -(|(if (v20) then v45 else v45)| * v2);
	for i7 := v44 to v21 {
		v3[safeIndex(964, v3.Length)] := v2 + v3[safeIndex(964, v3.Length)];
		v3[safeIndex(964, v3.Length)] := v2;
		var v46 := DC38(v20, v43);
		var v47: seq<bool> := [false, false];
		var v48: set<seq<bool>> := {v47, v47, v47};
		var v49: map<bool, array<bool>> := map[v0 := v15];
		var v50 := DC22(v8, v0, v20, v15, 0x363);
		var v51: array<array<bool>> := new array<bool>[18] [v11, v11, v15, v15, v15, v11, v11, v15, v11, v11, v15, v11, if (v0 in v49) then v49[v0] else v11, v15, v11, v15, v11, v50.cf33];
		var v52 := new C9(fm3(i7, v46.cf64, globalState), v48, v51, v2, fm0(v4, globalState));
		v3 := v3;
	}
	if (!!v20) {
		var v53: array<array<bool>> := new array<bool>[10] [v11, v11, v15, v15, v11, v11, v11, v11, v15, v15];
		var v54: C5 := new C5(v20, v2, v53, |v8|, v43);
		v3[safeIndex(964, v3.Length)] := |DC53([v54, v54, v54]).cf89|;
		var v55: set<char> := {v4, v4};
		var v56 := DC16(v54.f13, if (|v55| in v5) then v5[|v55|] else v43, v22[safeIndex(766, v22.Length)]);
		var v57: seq<bool> := [true, v54.fm8(v21, v0, v8, v56.cf22, globalState)];
		v4 := v8[safeIndex(safeModuloInt(|v57|, v2), |v8|)];
		var v59: seq<array<array<bool>>> := [v53];
		var v60: T1 := new C9(v2, set v58 : seq<bool> | v58 in (seq(abs(-0x58), i8  => (v57))) :: (v58), v59[safeIndex(v2, |v59|)], v44, v54.f13);
		var v61: set<T1> := {v60};
		v61 := v61;
		v3[safeIndex(964, v3.Length)] := v60.fm9(v54.f13, v54.fm8(v3[safeIndex(964, v3.Length)], v43, v8, [v54.f14, v3[safeIndex(964, v3.Length)]], globalState), v21, globalState);
		v43 := v54.f13;
	} else {
		var v63: multiset<multiset<int>> := multiset{v18};
		var v64, v65 := m0(map v62 : multiset<int> | v62 in v63 :: (v62) := (v8), v20, multiset(v14), v20 || v0, globalState);
		v3[safeIndex(964, v3.Length)] := |v18| * |v6|;
		var v66 := DC33(v43, v2, v1, v44, v40);
		if ((0x3d4 != v3[safeIndex(964, v3.Length)]) && v66.cf52) {
			var v67, v68 := m0(v19, v43, v18, v20, globalState);
			v67 := v64;
			var v69: map<int, multiset<int>> := map[0x140 := v18];
			var v70: seq<bool> := [v0, v69 != v69];
			v15[safeIndex(99, v15.Length)] := v70[safeIndex(v2, |v70|)];
			v15[safeIndex(99, v15.Length)] := v20;
			v0 := if (if (v64) then v0 else v64) then v0 else v15[safeIndex(99, v15.Length)];
		} else {
			var v71: seq<bool> := [v64, v43];
			v0 := v71[safeIndex(v3[safeIndex(964, v3.Length)], |v71|)];
			v8 := "becsmvx";
			var v72, v73 := m0(v19, v2 <= v2, v18, v43, globalState);
			var v74 := DC44(false);
			v74 := v74.(cf72 := v43);
			v72 := fm3(v3[safeIndex(964, v3.Length)], v72, globalState) !in (v18 + multiset{-977, 547, if (v20 in v1) then v1[v20] else v44});
		}
		
		var v75: multiset<seq<int>> := multiset{[v2, v65], v14};
		var v76 := DC52(v75);
		var v77: seq<bool> := [v20];
		var v78: seq<seq<bool>> := [v77];
		var v79 := DC27(v78);
		var v80: map<string, D12> := map[v8 := v79];
		var v81: seq<map<string, D12>> := [v80, v80];
		var v82: array<D25> := new D25[25] [v76, v76, DC52(v75), v76, v76, DC52(v75), fm72(v81[safeIndex(v44, |v81|)], v18, globalState), v76, v76, DC52(fm73(globalState)), v76.(cf88 := v75), v76, v76, v76, v76, v76, v76, v76, v76, v76, DC52(multiset(seq(abs(790), i9  => (v14)))), v76, v76, v76, v76.(cf88 := v75[v22[safeIndex(766, v22.Length)] := abs(v3[safeIndex(964, v3.Length)])])];
		v82[safeIndex(428, v82.Length)] := v76;
		v82[safeIndex(428, v82.Length)] := v76;
		v4, v20, v2 := v4, !v15[safeIndex(99, v15.Length)], safeDivisionInt(|(v8 + v8)|, v21);
	}
	
	print v0, "\n";
	print v1 == multiset{false, false, false}, "\n";
	print globalState.f0 == multiset{false, false, false}, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print v2, "\n";
	print v3[0], "\n";
	print v3[1], "\n";
	print v3[2], "\n";
	print v3[3], "\n";
	print v3[4], "\n";
	print v3[5], "\n";
	print v3[6], "\n";
	print v3[7], "\n";
	print v3[8], "\n";
	print v3[9], "\n";
	print v3[10], "\n";
	print v3[11], "\n";
	print v3[12], "\n";
	print v3[13], "\n";
	print v4, "\n";
	print v5 == map[357604 := false, 0 := false], "\n";
	print v6 == {598}, "\n";
	print v7 == {{598}}, "\n";
	print v8, "\n";
	print v9 == map[{{598}} := true], "\n";
	print v10[0] == {}, "\n";
	print v10[1] == {}, "\n";
	print v11[0], "\n";
	print v11[1], "\n";
	print v11[2], "\n";
	print v11[3], "\n";
	print v11[4], "\n";
	print v11[5], "\n";
	print v11[6], "\n";
	print v11[7], "\n";
	print v11[8], "\n";
	print v11[9], "\n";
	print v11[10], "\n";
	print v11[11], "\n";
	print v11[12], "\n";
	print v11[13], "\n";
	print v11[14], "\n";
	print v11[15], "\n";
	print v11[16], "\n";
	print v11[17], "\n";
	print v11[18], "\n";
	print v11[19], "\n";
	print v11[20], "\n";
	print |v12|, "\n";
	print v13 == map[false := 'j'], "\n";
	print v14 == [598], "\n";
	print v15[0], "\n";
	print v15[1], "\n";
	print v15[2], "\n";
	print v15[3], "\n";
	print v15[4], "\n";
	print v15[5], "\n";
	print v15[6], "\n";
	print v15[7], "\n";
	print v15[8], "\n";
	print v15[9], "\n";
	print v15[10], "\n";
	print v15[11], "\n";
	print v15[12], "\n";
	print v15[13], "\n";
	print v15[14], "\n";
	print v15[15], "\n";
	print v15[16], "\n";
	print v15[17], "\n";
	print v15[18], "\n";
	print v15[19], "\n";
	print v15[20], "\n";
	print v17 == [DC1(), DC1()], "\n";
	print v18 == multiset{598}, "\n";
	print v19 == map[multiset{598} := "smw"], "\n";
	print v20, "\n";
	print v21, "\n";
	print v22[0] == [598, 598], "\n";
	print v22[1] == [598, 598], "\n";
	print v22[2] == [598, 598], "\n";
	print v22[3] == [598, 598], "\n";
	print v22[4] == [598, 598], "\n";
	print v22[5] == [598, 598], "\n";
	print v22[6] == [598, 598], "\n";
	print v22[7] == [598, 598], "\n";
	print v22[8] == [598, 598], "\n";
	print v22[9] == [598, 598], "\n";
	print v22[10] == [598, 598], "\n";
	print v22[11] == [598, 598], "\n";
	print v22[12] == [598, 598], "\n";
	print v22[13] == [598, 598], "\n";
	print v22[14] == [598, 598], "\n";
	print v22[15] == [598, 598], "\n";
	print v22[16] == [598, 598], "\n";
	print v22[17] == [598, 598], "\n";
	print v22[18] == [598, 598], "\n";
	print v22[19] == [598, 598], "\n";
	print v22[20] == [598, 598], "\n";
	print v22[21] == [598, 598], "\n";
	print v22[22] == [598], "\n";
	print v22[23] == [598, 598], "\n";
	print v23.cf1, "\n";
	print v36 == map[DC1() := 10], "\n";
	print v40 == {true}, "\n";
	print v41[0] == map[DC1() := 10], "\n";
	print v41[1] == map[DC1() := 10], "\n";
	print v41[2] == map[DC1() := 10], "\n";
	print v41[3] == map[DC1() := 10], "\n";
	print v41[4] == map[DC1() := 598], "\n";
	print v41[5] == map[DC1() := 10], "\n";
	print v41[6] == map[DC1() := 10], "\n";
	print v41[7] == map[DC1() := 10], "\n";
	print v41[8] == map[DC1() := -1], "\n";
	print v41[9] == map[DC1() := 10], "\n";
	print v41[10] == map[DC1() := 3], "\n";
	print v41[11] == map[DC1() := 10], "\n";
	print v42 == multiset{'j'}, "\n";
	print v43, "\n";
	print v44, "\n";
	print |v45|, "\n";
}