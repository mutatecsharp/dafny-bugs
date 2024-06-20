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
datatype D0 = DC1(cf1: bool, cf2: int) | DC0(cf0: multiset<int>) | DC2(cf3: D0)
datatype D1 = DC4(cf5: bool) | DC3(cf4: array<int>) | DC5(cf6: D1)
datatype D2 = DC7(cf8: int, cf9: int) | DC6(cf7: multiset<bool>) | DC8(cf10: D2)
datatype D3 = DC10(cf12: bool, cf13: int, cf14: bool, cf15: int, cf16: array<int>) | DC9(cf11: map<int, bool>)
datatype D4 = DC11(cf17: array<bool>)
datatype D5 = DC12(cf18: string)
datatype D6 = DC14(cf20: seq<bool>) | DC15(cf21: bool, cf22: bool, cf23: int) | DC16(cf24: seq<D1>, cf25: string) | DC13(cf19: char)
datatype D7 = DC18(cf27: bool) | DC17(cf26: C0) | DC19(cf28: D7)
datatype D8 = DC21(cf30: int, cf31: bool) | DC20(cf29: set<bool>)
class GlobalState {
	var f0 : bool
	const f1 : set<bool>
	var f2 : string
	const f3 : seq<bool>
	var f4 : seq<int>
	const f5 : int
	const f6 : int
	const f7 : bool
	const f8 : multiset<map<int, char>>
	const f9 : int
	const f10 : int
	const f11 : int
	const f12 : int
	const f13 : array<int>
	const f14 : int
	const f15 : int
	const f16 : int
	var f17 : map<string, multiset<int>>
	const f18 : bool
	const f19 : seq<seq<int>>
	const f20 : string
	const f21 : int
	const f22 : int
	const f23 : seq<int>
	const f24 : bool
	const f25 : int
	const f26 : bool
	const f27 : map<int, bool>
	constructor (f0 : bool, f1 : set<bool>, f2 : string, f3 : seq<bool>, f4 : seq<int>, f5 : int, f6 : int, f7 : bool, f8 : multiset<map<int, char>>, f9 : int, f10 : int, f11 : int, f12 : int, f13 : array<int>, f14 : int, f15 : int, f16 : int, f17 : map<string, multiset<int>>, f18 : bool, f19 : seq<seq<int>>, f20 : string, f21 : int, f22 : int, f23 : seq<int>, f24 : bool, f25 : int, f26 : bool, f27 : map<int, bool>) {
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
	const f28 : bool
	const f29 : array<int>
	constructor (f28 : bool, f29 : array<int>) {
		this.f28 := f28;
		this.f29 := f29;
	}
	
	function fm1(p0: D0, p1: int, p2: int, p3: int, globalState: GlobalState): char {
		if (!true) then 'u' else 'o'
	}
	function fm2(p0: bool, globalState: GlobalState): bool {
		f28
	}
}

function fm0(globalState: GlobalState): bool {
	!(0x237 >= (|['q']| * |DC14([true, false]).cf20|))
}
function fm3(p0: int, p1: bool, p2: int, p3: map<int, bool>, globalState: GlobalState): int {
	safeModuloInt(|[map v0 : int | v0 in (map v1 : int | v1 in [151] :: (v1 - 829) := (|[|{|(seq(abs(0x50), i0  => ('y')))|}|, 523]|)) :: (v0 - 0x21e) := (921), map[0x273 := 0x1d0], map[|[map[true := 0x27e]]| := 126], map[-0xdc := |(seq(abs(487), i1  => (|map[true := !false]|)))|]]|, -511) + 0x217
}
function fm4(globalState: GlobalState): multiset<int> {
	(multiset{0x1ce} + multiset{|[|"rnvmueh"|, |map[614 := multiset(seq(abs(0x336), i0  => ('k')))]|]|, -|map[true := 0xef]|, -650, |multiset([false])|}) - multiset([-0x249] + [913, |{map['e' := |"f"|], map['j' := |map[0x12d := false]|]}|])
}
function fm5(p0: char, globalState: GlobalState): map<int, bool> {
	(map[-|[true]| := false] + map[|[|map[-0x26c := |multiset{false}|]|]| := true]) + (map v0 : int | (0x3c3 <= v0) && (v0 < 0x297) :: (safeModuloInt(v0, 441)) := (false))
}
function fm6(p0: bool, p1: int, globalState: GlobalState): multiset<bool> {
	(if (true) then multiset{false, false, true} else multiset([true, false, false, !false, false])) + multiset{!true}
}
function fm7(p0: bool, p1: bool, p2: bool, p3: int, globalState: GlobalState): map<bool, bool> {
	map[true := false]
}
function fm8(p0: int, globalState: GlobalState): set<bool> {
	{0x5a >= -|map[|"qjl"| := -53]|}
}
function fm9(globalState: GlobalState): string {
	"usw"
}
function fm10(p0: int, p1: int, p2: int, globalState: GlobalState): seq<bool> {
	[(map v0 : int | (0x14f <= v0) && (v0 < 0x21b) :: (v0 * -|{false}|) := (-|(seq(abs(0x81), i0  => ('p')))|)) != map[0xba := |map[true := true]|], true !in [true, true, !true, true, true]]
}
function fm11(p0: int, p1: char, globalState: GlobalState): multiset<string> {
	(multiset{"xqmuuiom", "sjc", "f"} * multiset{"ofdbvnl"}) + (multiset{"bc"} - multiset(seq(abs(0x17b), i0  => (seq(abs(0x2a), i1  => ('w'))))))
}
function fm12(p0: int, p1: int, p2: int, p3: D7, globalState: GlobalState): multiset<map<bool, char>> {
	multiset{map[false := 'v']} - multiset{map[true := 'c'], map[false := 's'], map[true := 's']}
}
method m0(globalState: GlobalState) {
	var v0 := false;
	var v1 := -0x31e;
	var v2: map<bool, bool> := map[v0 := v0];
	var v3: seq<bool> := [!v0];
	var v4: set<int> := {v1, -v1};
	var v5: set<bool> := {v0};
	var v6: array<int> := new int[22] [v1, |v2|, -v1, |v3|, |v4|, v1, v1, v1, v1, v1, v1, v1, v1, v1, v1, |v5|, |v3|, -0x124, v1, v1, v1, v1];
	var v7 := new C0(v0, if (v0) then v6 else v6);
	var i0 := 0;
	while (v7.f28)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		if (fm0(globalState)) {
			var v8 := 'd';
			v8 := if (v0) then v8 else v8;
			var v9: array<bool> := new bool[19] [v0, v7.f28, v0, v7.f28, v7.f28, v7.fm2(v7.f28, globalState), v7.f28, v7.f28, !v7.f28, v0, v7.f28, v0, v7.f28, v7.f28, v7.f28, v7.f28, !v7.f28, v0, v7.f28];
			var v10 := DC11(v9);
			var v11: seq<array<bool>> := [v9, v9, v9, v9, v9];
			var v12 := "btlenkt";
			var v13: set<array<bool>> := {v9, v10.cf17, v9, v11[safeIndex(-|v12|, |v11|)]};
			var v14 := DC1(v7.f28, v1);
			var v15 := DC2(v14);
			var v16: multiset<bool> := multiset{v0, !v7.f28};
			var v17: map<int, array<bool>> := map[v1 := v9];
			var v18: seq<set<array<bool>>> := [{if (|{v1, v1, |v5|, v1}| in v17) then v17[|{v1, v1, |v5|, v1}|] else v9}, v13, v13 - v13, {v9}, v13 + v13];
			var v19 := DC4(v7.f28);
			v13, v15, v9, v16 := v18[safeIndex(v1, |v18|)], v15, v9, (v16 + v16) * fm6(v19.cf5, v1, globalState);
			var v20 := new C0(v7.f28, v7.f29);
			var v21: map<bool, seq<bool>> := map[v7.f28 := v3];
			var v22: seq<seq<bool>> := [[v7.f28], v3, (if (v0 in v21) then v21[v0] else v3) + v3, [v20.f28]];
			var v23: map<int, int> := map[v1 := v1];
			var v24 := DC10(v7.fm2(v7.f28, globalState), |v23|, !v7.f28, 0x11a, v7.f29);
			v9[safeIndex(970, v9.Length)] := v24.cf14;
			globalState.f0, globalState.f0, v22, v9[safeIndex(970, v9.Length)] := false, v1 > (v1 * |map[v1 := -0x39c]|), v22, !v7.f28;
			var v25 := DC6(v16);
			var v26: map<D2, int> := map[if (v20.fm2(v7.f28, globalState)) then v25 else v25 := 0x349];
			v26 := v26[DC6(v16) := v1];
		} else {
			var v27: multiset<bool> := multiset{v1 <= v1, v1 == v1};
			v1, v1, v6 := v1, if (v7.f28 in v27) then v27[v7.f28] else 591, v7.f29;
			var v28: array<bool> := new bool[25];
			v28[safeIndex(941, v28.Length)] := v7.f28;
			var v29 := 'o';
			var v30: seq<char> := [v29, v29, v29];
			var v31 := DC1(v0, -832);
			v1, v28[safeIndex(941, v28.Length)], v0 := |v30| + v31.cf2, v7.f28, if (v7.f28) then v1 in (set v32 : int | v32 in (seq(abs(0x3b0), i1  => (v1))) :: (safeModuloInt(v32, |"sbrnrtwv"|))) else v1 < v1;
			v29 := v29;
			var v33: multiset<C0> := multiset{v7, v7, v7};
			v1 := if (v7 in v33) then v33[v7] else v1 + v1;
			var v34: map<string, bool> := map[v30 := v0];
			var v35: array<map<bool, bool>> := new map<bool, bool>[7] [v2[v0 := false], v2, v2, v2 + fm7(v7.f28, v7.f28, v28[safeIndex(941, v28.Length)], v1, globalState), v2, (fm7(v0, fm0(globalState), v7.f28, |v34|, globalState))[v0 := v0], v2];
			v35[safeIndex(382, v35.Length)] := map[v7.f28 := v7.fm2(v7.f28, globalState)] + map[v28[safeIndex(941, v28.Length)] := v7.f28];
			v35[safeIndex(382, v35.Length)] := v2;
		}
		
		if ((!v7.f28 <== v7.f28) <== v0) {
			var v36 := "vithblii";
			globalState.f0 := 'i' in (v36 + v36);
			var v37: array<char> := new char[7](i2 => 'y');
			var v38 := DC12(v36);
			globalState.f0, v37 := (seq(abs(0x2a8), i3  => ('h'))) <= ((v38.(cf18 := seq(abs(0x366), i4  => ('x')))).cf18 + v36), if (v7.f28) then v37 else if (v7.f28) then v37 else v37;
			v0 := v0;
			v1 := safeDivisionInt(v1, v1) - fm3(v1, v7.f28, v1, map[v1 := v7.f28], globalState);
			var v39: array<set<bool>> := new set<bool>[23];
			var v40: map<int, C0> := map[v1 := v7];
			var v41: map<int, set<bool>> := map[|v40| := fm8(v1, globalState)];
			v39[safeIndex(908, v39.Length)] := {v7.f28, v0, v0, v7.f28, v7.f28} - (if (330 in v41) then v41[330] else fm8(v1, globalState));
			v39[safeIndex(908, v39.Length)] := v5 - (v5 + v5);
		} else {
			var v42 := 'l';
			var v43 := DC10(v0, v1, v7.f28, v1, v7.f29);
			var v44: map<bool, C0> := map[v43.cf14 := v7];
			var v45: seq<C0> := [if (!v7.fm2(v7.f28, globalState) in v44) then v44[!v7.fm2(v7.f28, globalState)] else v7, v7, v7];
			var v46 := "squrxyb";
			var v47 := DC12(v46);
			var v48: map<D5, bool> := map[v47 := false];
			var v49 := DC13(v42);
			var v50 := DC13((v49.(cf19 := v42)).cf19);
			v1, v7, v42 := v1, v45[safeIndex(|(v48 + v48)|, |v45|)], v49.cf19;
			var v51: array<bool> := new bool[5] [v7.f28, v7.f28, v0, v7.f28, v0];
			var v52: array<array<int>> := new array<int>[18] [v7.f29, v6, v7.f29, v7.f29, v7.f29, v7.f29, v7.f29, v7.f29, v7.f29, v7.f29, v7.f29, v7.f29, v7.f29, v7.f29, v6, v7.f29, v7.f29, v6];
			var v53: map<array<bool>, array<array<int>>> := map[v51 := v52];
			v53 := v53[v51 := v52];
			var v54 := new C0(v0, v7.f29);
			var v55: multiset<int> := multiset{v1};
			var v56: map<string, int> := map[seq(abs(-0x2df), i5  => (v42)) := |v55|];
			v56 := v56[fm9(globalState) := v1];
			var v57: map<int, int> := map[0x113 := |("v" + v46)|];
			var v58: map<int, bool> := map[v1 := v7.f28];
			var v59: seq<map<int, bool>> := [v58];
			v57 := v57[if (v3[safeIndex(765, |v3|)]) then v1 else fm3(v1, v7.f28, v1, v59[safeIndex(|v46|, |v59|)], globalState) := v1];
		}
		
		v1 := if (v7.fm2(v0, globalState)) then v1 else v1;
		globalState.f0 := v7.f28;
	}
	var v60: multiset<bool> := multiset{v0, v7.f28, v0, v7.f28};
	var v61: array<bool> := new bool[8];
	var v62: map<int, array<bool>> := map[0x26a := v61];
	var v63: seq<map<int, array<bool>>> := [v62, v62, v62];
	globalState.f0, v0, globalState.f0, v1 := v60 !! v60, true, v7.f28, |(v63[safeIndex(|[v1, v1]|, |v63|)] + v62)| - v1;
	if (!v7.f28) {
		v1 := v1;
		globalState.f0 := v7.f28;
		var v64: seq<seq<bool>> := [[v0]];
		var v65 := DC1(v0, |v64|);
		if (v65.cf1) {
			v61[safeIndex(769, v61.Length)] := v7.f28;
			v61[safeIndex(769, v61.Length)] := v7.f28 <==> v7.f28;
			var v66: array<char> := new char[26];
			var v67: set<array<char>> := {v66};
			globalState.f0 := {v66} <= v67;
			var v68 := DC14(fm10(v1, 0x2a1, v1, globalState));
			var v69 := DC3(v7.f29);
			var v70: map<int, D1> := map[|v68.cf20| := v69];
			var v71 := DC5(if (v1 in v70) then v70[v1] else v69);
			var v72 := DC16([v71], "am");
			var v73 := "jrtiwitvl";
			var v74: seq<D1> := [v71.(cf6 := v69), v71];
			var v75 := 'i';
			var v76: array<D6> := new D6[10] [v72, v72.(cf25 := v73).(cf24 := v74), v72, v72, v72, v72, v72.(cf25 := v73[safeIndex(|v2|, |v73|) := v75]), v72, v72, v72];
			v76[safeIndex(929, v76.Length)] := DC16(v74, v73);
			v76[safeIndex(929, v76.Length)] := if (v7.f28) then v72 else v72;
			globalState.f0 := v7.f28;
			v72 := v72;
		} else {
			var v77: set<string> := {"xsxaj"};
			globalState.f0 := v77 <= v77;
			var v78: multiset<int> := multiset{777};
			var v79: map<int, multiset<int>> := map[v1 * |v4| := v78];
			v78 := if (|("ysstjicth" + (seq(abs(-0x2f6), i6  => ('h'))))| in v79) then v79[|("ysstjicth" + (seq(abs(-0x2f6), i6  => ('h'))))|] else v78;
			var v80 := "ipw";
			var v81: map<int, string> := map[v1 := v80];
			v3, globalState.f2, v1, globalState.f0 := fm10(-0x2a8 + v1, v1, v1, globalState), (v80 + (if (v1 in v81) then v81[v1] else v80)) + v80, safeModuloInt(491, v1), v7.f28;
			v7.f29[safeIndex(738, v7.f29.Length)] := v1;
			v7.f29[safeIndex(738, v7.f29.Length)] := v1;
			var v82: map<C0, array<int>> := map[v7 := v7.f29];
			v82 := v82[v7 := v7.f29];
		}
		
		var v83 := 'p';
		v83 := v83;
		var v84: array<set<int>> := new set<int>[14];
		v84[safeIndex(57, v84.Length)] := set v85 : int | (-913 <= v85) && (v85 < 0x148) :: (v85 - |multiset{v1, v1}|);
		v84[safeIndex(57, v84.Length)] := v4;
	} else {
		var v86: map<bool, int> := map[v0 := v1];
		v1 := (if (!!v7.f28 in v86) then v86[!!v7.f28] else v1) + v1;
		globalState.f0 := true;
		v61[safeIndex(488, v61.Length)] := v7.f28 <== v7.f28;
		var v87 := "dicp";
		var v88: seq<int> := [v1];
		v3, v1, v61[safeIndex(488, v61.Length)] := v3, if (v0) then |(v87 + "dqnhuuat")| else -v88[safeIndex(v1, |v88|)], v7.f28;
		var v89: array<bool> := new bool[13] [v61[safeIndex(488, v61.Length)], v60 !! multiset{v0, v0, v7.f28, v0}, v0, false, v7.f28, v61[safeIndex(488, v61.Length)], true, v7.f28, v61[safeIndex(488, v61.Length)], multiset{0x2a0} == multiset(seq(abs(0x2f0), i7  => (|v88|))), v61[safeIndex(488, v61.Length)], v7.f28, v7.f28];
		v89 := v89;
		var v90: seq<C0> := [v7, v7, v7];
		var v91: map<int, bool> := map[v1 := v61[safeIndex(488, v61.Length)]];
		v7 := v90[safeIndex(fm3(0x1f6, false, v1, v91, globalState), |v90|)];
	}
	
	var v92 := DC1(if (!v7.f28) then v7.f28 else v0, v1);
	match (v92) {
		case DC1(cf1, cf2) =>
			v1 := v1;
			var v93: map<int, bool> := map[cf2 := !cf1];
			cf2 := -fm3(cf2 - cf2, v7.f28, v1, v93, globalState);
			var v94: map<D0, map<int, bool>> := map[v92.(cf2 := v1) := v93];
			var v95: seq<set<int>> := [v4, v4];
			var v96: seq<int> := [v1, cf2, v1, v1];
			v94, globalState.f0, cf2, v1 := v94 + (v94 + v94), !v7.f28 <== (v95[safeIndex(cf2, |v95|)] == v4), safeDivisionInt(cf2, -|v96|), cf2;
			if (v1 != v1) {
				var v97 := new C0(v7.f28, v7.f29);
				v7.f29[safeIndex(804, v7.f29.Length)] := 0x1fb * cf2;
				v7.f29[safeIndex(804, v7.f29.Length)] := |(v2 + v2)|;
				var v98: seq<array<bool>> := [v61];
				var v99: multiset<int> := multiset{v1};
				var v100 := "ckgrfoe";
				var v101: map<int, string> := map[v1 := v100];
				var v102 := DC0(v99);
				var v103: map<int, multiset<int>> := map[v7.f29[safeIndex(804, v7.f29.Length)] := v99];
				var v104: array<multiset<int>> := new multiset<int>[18] [v99 + v99, v99[0xbb := abs(v7.f29[safeIndex(804, v7.f29.Length)])] - v99, v99, multiset{|v100|, cf2, v1, |v100|, v1} * v99, multiset{---|v101|}, v99, (multiset(v96))[v1 := abs(v1)], v99, v99, multiset(v96), v99 * v99, v102.cf0, multiset([v1]), v99, v99, v99, (if (705 in v103) then v103[705] else v99)[cf2 := abs(v1)], multiset{cf2, -0x297}];
				v104[safeIndex(414, v104.Length)] := multiset(v96);
				var v105 := 'b';
				var v106: map<bool, char> := map[true := v105];
				var v107: multiset<map<bool, char>> := multiset{map[cf1 := v105], map[true := 'k'], v106[v7.f28 := v105], v106, if (v97.f28) then v106 else v106};
				var v108: map<int, C0> := map[|v100| := v97];
				var v109 := DC17(if (|v3| in v108) then v108[|v3|] else v7);
				var v110: seq<C0> := [v109.cf26];
				var v111: seq<string> := [v100, v100];
				var v112: multiset<string> := multiset{"f", v100, v100, "a", seq(abs(55), i8  => (v105))};
				var v113 := DC18(v7.f28);
				v98, v7.f29[safeIndex(804, v7.f29.Length)], v104[safeIndex(414, v104.Length)], globalState.f0, v107 := v98, |(v110 + [v97, v97, v97])[safeIndex(0x17d, |(v110 + [v97, v97, v97])|) := v97]|, v99, (multiset(v111) - v112) > (if (v97.f28) then fm11(-0x282, 'a', globalState) else v112), fm12(v1, v7.f29[safeIndex(804, v7.f29.Length)], if (v97.f28 in v60) then v60[v97.f28] else v7.f29[safeIndex(804, v7.f29.Length)], v113, globalState);
				v61[safeIndex(991, v61.Length)] := v7.f28;
				v61[safeIndex(991, v61.Length)] := v7.f28;
				v7.f29[safeIndex(804, v7.f29.Length)] := cf2 + (v1 * v7.f29[safeIndex(804, v7.f29.Length)]);
			} else {
				v93 := v93[v1 := v7.f28];
				v7 := v7;
				var v114 := DC11(v61);
				var v115: array<char> := new char[3];
				v115[safeIndex(346, v115.Length)] := 'k';
				var v116 := 'p';
				v114, v115[safeIndex(346, v115.Length)] := DC11(v61).(cf17 := v61), v116;
				v7.f29[safeIndex(767, v7.f29.Length)] := 0x15;
				var v117: multiset<int> := multiset{-0x144};
				v7.f29[safeIndex(767, v7.f29.Length)] := -(if (v1 in v117) then v117[v1] else v1);
				var v118: seq<array<bool>> := [v61, if (v7.f28) then v61 else v61, v61];
				v7.f29[safeIndex(767, v7.f29.Length)], v61, cf2 := v1, v118[safeIndex(v7.f29[safeIndex(767, v7.f29.Length)], |v118|)], v1;
			}
			
		case DC0(cf0) =>
			v1 := fm3(v1 + v1, v7.f28, v1, map[v1 := v7.f28], globalState);
			var v119 := DC15(v7.f28, v7.f28, v1);
			v1 := v119.cf23;
			v1 := safeDivisionInt(v1, v1);
			var v120 := "sudyiqf";
			globalState.f2 := v120 + (v120 + v120);
		case DC2(cf3) =>
			var v121 := DC5(DC3(v7.f29));
			var v122 := DC5(v121);
			var v123 := DC5(v122);
			var v124: seq<D1> := [v123];
			var v125 := "yxedcgbw";
			var v126 := DC16(v124, v125);
			match (v126) {
				case DC14(cf20) =>
					var v127: map<bool, int> := map[v7.f28 := |"kmsp"|];
					v1 := |((map[v7.f28 := v1] + v127) + v127)|;
					v7.f29[safeIndex(878, v7.f29.Length)] := v1;
					v7.f29[safeIndex(878, v7.f29.Length)] := v1;
					v1 := v7.f29[safeIndex(878, v7.f29.Length)] - v7.f29[safeIndex(878, v7.f29.Length)];
					var v128: seq<string> := [v125];
					globalState.f2 := v128[safeIndex(v1, |v128|)];
				case DC15(cf21, cf22, cf23) =>
					v7.f29[safeIndex(646, v7.f29.Length)] := -0x3d5;
					v7.f29[safeIndex(646, v7.f29.Length)] := v1;
					globalState.f2 := v125;
					globalState.f0 := true;
					globalState.f0 := !v7.f28;
				case DC16(cf24, cf25) =>
					var v129: set<C0> := {v7};
					var v130: map<set<C0>, bool> := map[v129 := v7.f28];
					var v131: array<map<set<C0>, bool>> := new map<set<C0>, bool>[13] [map[v129 := v7.f28], map[{v7} := !v7.f28] + v130, v130 + v130, v130, v130, v130, v130 + v130, v130, (map[v129 := v7.f28])[{v7} := v7.f28], v130[v129 := v7.f28], v130, v130, v130];
					v131[safeIndex(24, v131.Length)] := v130 + v130;
					var v132 := DC18(v7.f28);
					var v133 := DC19(v132);
					var v134 := DC19(v132);
					v131[safeIndex(24, v131.Length)], v7, globalState.f2, v134 := v130, if (v0) then v7 else v7, fm9(globalState), v134;
					v4 := v4;
					v1 := safeDivisionInt(|cf25|, v1);
					var v135: map<int, bool> := map[-safeDivisionInt(v1, v1) := v7.f28];
					var v136: seq<int> := [v1];
					v135 := v135[|{[294], v136, v136}| := v7.f28];
				case DC13(cf19) =>
					globalState.f2 := v125 + v126.cf25;
					var v137 := new C0(v7.f28, v6);
					var v138 := DC4(v1 <= v1);
					v1, v138, globalState.f0, globalState.f0, v6 := safeModuloInt(v1, v1), v138, v7.fm2(v7.f28, globalState), !v7.f28 || v137.f28, v137.f29;
					v1 := v1 - -v1;
			}
			
			var v139 := 'g';
			var v140: array<char> := new char[2] [v139, v139];
			var v141: map<char, array<char>> := map[v139 := v140];
			var v142: seq<map<char, array<char>>> := [map[v139 := v140]];
			var v143: seq<C0> := [v7];
			var v144: multiset<int> := multiset{|(seq(abs(597), i9  => ('f')))|, |v143|};
			var v145 := DC20(v5);
			var v146: array<map<char, array<char>>> := new map<char, array<char>>[24] [v141, v141, v142[safeIndex(if (v1 in v144) then v144[v1] else |v145.cf29|, |v142|)], v141, v141, map[v139 := v140], v141, v141, v141 + v141, v142[safeIndex(-0x1a9, |v142|)], v141, v141, v142[safeIndex(267, |v142|)], map[v139 := v140] + v141, map[v139 := v140], v141, v141, v141, v141[v139 := v140], v141, v141, v141, map[v139 := v140] + v141[v139 := v140], v141[v139 := v140] + v141];
			v146[safeIndex(957, v146.Length)] := v141;
			v146[safeIndex(957, v146.Length)] := v141 + map[v139 := v140];
			var v147: map<int, bool> := map[v1 := v7.f28];
			v1 := safeModuloInt(v1, --(|v147| - v1));
			v7.f29[safeIndex(649, v7.f29.Length)] := 0x15a;
			v7.f29[safeIndex(649, v7.f29.Length)] := v1 - v92.cf2;
	}
	
	var v148: map<int, int> := map[v1 := v1];
	var v149: map<seq<map<int, int>>, bool> := map[[v148, v148] := v7.f28];
	var v150 := 'q';
	var v151: map<bool, char> := map[v7.f28 := v150];
	v149 := v149[[v148[v1 := |v151|]] := v7.f28];
}
method Main() {
	var v0 := false;
	var v1: map<bool, bool> := map[v0 := v0];
	var v2: set<bool> := {v0};
	var v3 := "sk";
	var v4: seq<bool> := [v0, v0, v0, v0];
	var v5 := 0xe4;
	var v6: seq<int> := [v5];
	var v7: set<int> := {|v6|};
	var v8: seq<int> := [-v5, |v7|];
	var v9 := 'v';
	var v10: map<int, char> := map[-v5 := v9];
	var v11: multiset<map<int, char>> := multiset{v10};
	var v12: array<int> := new int[13](i0 => i0 + v5);
	var v13: map<string, multiset<int>> := map["ybllip" := multiset{v5, v5}];
	var v15: seq<seq<int>> := [v6];
	var v16 := DC1(v0, |v1|);
	var v17 := DC2(v16);
	var v18: multiset<D0> := multiset{v17, v17};
	var v19: multiset<multiset<D0>> := multiset{v18};
	var v20: map<int, bool> := map[if (v18 in v19) then v19[v18] else v5 := v0];
	var globalState := new GlobalState(true, {if (v0 in v1) then v1[v0] else v0} - v2, v3, v4, v8, -0x3a1, 851, true, v11, 0x5a, -909, 917, 0x150, v12, 0xce, 664, -0x1b0, v13 + (map v14 : string | v14 in (seq(abs(-0xbf), i1  => ("incqt")))[safeIndex(v5, |(seq(abs(-0xbf), i1  => ("incqt")))|) := v3] :: (v14) := (DC0(multiset{v5}).cf0)), true, v15, v3[safeIndex(v5, |v3|) := v9], -661, -0x59, [v5], true, 0x351, false, v20);
	m0(globalState);
	if (|v4| == |v20|) {
		v5 := |v3|;
		globalState.f0 := !!fm0(globalState);
		var v21 := new C0(v0, v12);
		var v22 := DC1(v0, fm3(v5, v0, fm3(v5, v21.f28, 0x220, v20, globalState), v20, globalState));
		var v23 := DC3(v12);
		var v24: C0 := new C0(v22.cf1, v23.cf4);
		v24 := v24;
		v9 := v24.fm1(DC1(v21.f28, v5), v5, v5, |v3|, globalState);
	} else {
		v12[safeIndex(59, v12.Length)] := |v7|;
		v12[safeIndex(59, v12.Length)] := -|(v1 + v1[v0 := v0])|;
		var v25: array<int> := new int[21](i2 => safeModuloInt(i2, |v8|));
		var v26 := new C0(v0, v25);
		globalState.f0 := v26.f28;
		v5 := v12[safeIndex(59, v12.Length)];
		var v27: array<bool> := new bool[11];
		var v29: multiset<map<int, bool>> := multiset{v20, map[v12[safeIndex(59, v12.Length)] := v26.f28], map v28 : int | v28 in (seq(abs(-0x188), i3  => (|DC6(multiset{v26.f28}).cf7|))) :: (safeDivisionInt(v28, v5)) := (v26.f28), v20, v20};
		v27[safeIndex(414, v27.Length)] := |v29| < v12[safeIndex(59, v12.Length)];
		v27[safeIndex(414, v27.Length)] := v26.f28;
	}
	
	v0 := v0;
	v5 := v5;
	var v30: seq<string> := [v3, v3];
	v30 := v30 + (v30 + v30);
	var v31: multiset<int> := multiset{safeDivisionInt(-v5, -v5)};
	v31 := fm4(globalState) - v31;
	var v32 := DC9(v20);
	globalState.f0 := if (v5 <= fm3(-v5, v0, v5, v32.cf11, globalState)) then !(|map[v0 := true]| in multiset([v5])) else !v0 <== v0;
	var v33: array<string> := new string[20](i4 => "skjjyean");
	var v34: map<array<string>, int> := map[v33 := v5];
	v34 := v34[v33 := v5];
	v5 := v5 + |v3|;
	globalState.f0 := v0;
	globalState.f0 := !(if (v0 in v1) then v1[v0] else !v0 <==> v0);
	var v35: map<bool, array<int>> := map[v0 := v12];
	var v36: C0 := new C0(v0, v12);
	var v37: set<C0> := {v36, v36, v36};
	var v38: map<map<bool, array<int>>, bool> := map[v35 := v37 <= v37];
	var v39: seq<map<bool, array<int>>> := [map[v36.f28 := v12]];
	var v40: seq<map<bool, array<int>>> := [v39[safeIndex(v5, |v39|)]];
	v38 := v38[v40[safeIndex(v5, |v40|)] + v35 := true];
	var v41: map<map<bool, bool>, C0> := map[(map[v0 := v36.f28])[true := v0] := v36];
	for i5 := |(v30 + v30)| to |v41| {
		globalState.f0 := !!(v7 != v7);
		globalState.f0 := v36.f28;
		var v42: multiset<C0> := multiset{v36};
		var v43 := new C0(multiset{v36} <= v42, v12);
		var v44: map<int, seq<bool>> := map[i5 := [v36.f28, v0, !v36.f28]];
		v4 := (if (|(map[v43 := v5] + (map[v36 := -0xc7])[v43 := i5])| in v44) then v44[|(map[v43 := v5] + (map[v36 := -0xc7])[v43 := i5])|] else [!v36.f28])[safeIndex(i5 + v5, |(if (|(map[v43 := v5] + (map[v36 := -0xc7])[v43 := i5])| in v44) then v44[|(map[v43 := v5] + (map[v36 := -0xc7])[v43 := i5])|] else [!v36.f28])|) := v36.f28];
	}
	v32 := v32;
	var v45 := DC10(v36.f28, 0x119, v0, v5, v36.f29);
	v5 := safeModuloInt(|"iy"|, fm3(v5, v36.f28, |[!v45.cf14]|, fm5(v9, globalState), globalState));
	forall i6 | 0 <= i6 < v33.Length {
		v33[i6] := "lnqf";
	}
	print v0, "\n";
	print v1 == map[false := false], "\n";
	print v2 == {false}, "\n";
	print v3, "\n";
	print v4 == [false, false, false, false], "\n";
	print v5, "\n";
	print v6 == [228], "\n";
	print v7 == {1}, "\n";
	print v8 == [-228, 1], "\n";
	print v9, "\n";
	print v10 == map[-228 := 'v'], "\n";
	print v11 == multiset{map[-228 := 'v']}, "\n";
	print v12[0], "\n";
	print v12[1], "\n";
	print v12[2], "\n";
	print v12[3], "\n";
	print v12[4], "\n";
	print v12[5], "\n";
	print v12[6], "\n";
	print v12[7], "\n";
	print v12[8], "\n";
	print v12[9], "\n";
	print v12[10], "\n";
	print v12[11], "\n";
	print v12[12], "\n";
	print v13 == map["ybllip" := multiset{228, 228}], "\n";
	print v15 == [[228]], "\n";
	print v16.cf1, "\n";
	print v16.cf2, "\n";
	print v17.cf3.cf1, "\n";
	print v17.cf3.cf2, "\n";
	print v18 == multiset{DC2(DC1(false, 1)), DC2(DC1(false, 1))}, "\n";
	print v19 == multiset{multiset{DC2(DC1(false, 1)), DC2(DC1(false, 1))}}, "\n";
	print v20 == map[1 := false], "\n";
	print globalState.f0, "\n";
	print globalState.f1 == {}, "\n";
	print globalState.f2, "\n";
	print globalState.f3 == [false, false, false, false], "\n";
	print globalState.f4 == [-228, 1], "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8 == multiset{map[-228 := 'v']}, "\n";
	print globalState.f9, "\n";
	print globalState.f10, "\n";
	print globalState.f11, "\n";
	print globalState.f12, "\n";
	print globalState.f13[0], "\n";
	print globalState.f13[1], "\n";
	print globalState.f13[2], "\n";
	print globalState.f13[3], "\n";
	print globalState.f13[4], "\n";
	print globalState.f13[5], "\n";
	print globalState.f13[6], "\n";
	print globalState.f13[7], "\n";
	print globalState.f13[8], "\n";
	print globalState.f13[9], "\n";
	print globalState.f13[10], "\n";
	print globalState.f13[11], "\n";
	print globalState.f13[12], "\n";
	print globalState.f14, "\n";
	print globalState.f15, "\n";
	print globalState.f16, "\n";
	print globalState.f17 == map["ybllip" := multiset{228, 228}, "incqt" := multiset{228}, "sk" := multiset{228}], "\n";
	print globalState.f18, "\n";
	print globalState.f19 == [[228]], "\n";
	print globalState.f20, "\n";
	print globalState.f21, "\n";
	print globalState.f22, "\n";
	print globalState.f23 == [228], "\n";
	print globalState.f24, "\n";
	print globalState.f25, "\n";
	print globalState.f26, "\n";
	print globalState.f27 == map[1 := false], "\n";
	print v30 == ["sk", "sk", "sk", "sk", "sk", "sk"], "\n";
	print v31 == multiset{462, -1, -650}, "\n";
	print v32.cf11 == map[1 := false], "\n";
	print v33[0], "\n";
	print v33[1], "\n";
	print v33[2], "\n";
	print v33[3], "\n";
	print v33[4], "\n";
	print v33[5], "\n";
	print v33[6], "\n";
	print v33[7], "\n";
	print v33[8], "\n";
	print v33[9], "\n";
	print v33[10], "\n";
	print v33[11], "\n";
	print v33[12], "\n";
	print v33[13], "\n";
	print v33[14], "\n";
	print v33[15], "\n";
	print v33[16], "\n";
	print v33[17], "\n";
	print v33[18], "\n";
	print v33[19], "\n";
	print |v34|, "\n";
	print |v35|, "\n";
	print v36.f28, "\n";
	print v36.f29[0], "\n";
	print v36.f29[1], "\n";
	print v36.f29[2], "\n";
	print v36.f29[3], "\n";
	print v36.f29[4], "\n";
	print v36.f29[5], "\n";
	print v36.f29[6], "\n";
	print v36.f29[7], "\n";
	print v36.f29[8], "\n";
	print v36.f29[9], "\n";
	print v36.f29[10], "\n";
	print v36.f29[11], "\n";
	print v36.f29[12], "\n";
	print |v37|, "\n";
	print |v38|, "\n";
	print |v39|, "\n";
	print |v40|, "\n";
	print |v41|, "\n";
	print v45.cf12, "\n";
	print v45.cf13, "\n";
	print v45.cf14, "\n";
	print v45.cf15, "\n";
	print v45.cf16[0], "\n";
	print v45.cf16[1], "\n";
	print v45.cf16[2], "\n";
	print v45.cf16[3], "\n";
	print v45.cf16[4], "\n";
	print v45.cf16[5], "\n";
	print v45.cf16[6], "\n";
	print v45.cf16[7], "\n";
	print v45.cf16[8], "\n";
	print v45.cf16[9], "\n";
	print v45.cf16[10], "\n";
	print v45.cf16[11], "\n";
	print v45.cf16[12], "\n";
}