import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_
import _dafny
import System_

# Module: module_

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def abs(x):
        if (x) < (0):
            return (-1) * (x)
        elif True:
            return x

    @staticmethod
    def safeIndex(x, length):
        if (x) < (0):
            return 0
        elif (x) >= (length):
            return _dafny.euclidian_modulus(x, length)
        elif True:
            return x

    @staticmethod
    def safeDivisionInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_division(x1, x2)

    @staticmethod
    def safeModuloInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_modulus(x1, x2)

    @staticmethod
    def fm0(p0, p1, globalState):
        return (856) - ((964) + (len(_dafny.SeqWithoutIsStrInference([False]))))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        return not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxdfowin"))) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imhcdljya"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gng")))))

    @staticmethod
    def fm3(globalState):
        return _dafny.SeqWithoutIsStrInference([-386])

    @staticmethod
    def fm4(globalState):
        return _dafny.CodePoint('w')

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: str
            for compr_0_ in (_dafny.Map({_dafny.CodePoint('p'): -604})).keys.Elements:
                d_0_v0_: str = compr_0_
                if (d_0_v0_) in (_dafny.Map({_dafny.CodePoint('p'): -604})):
                    coll0_[d_0_v0_] = not(False)
            return _dafny.Map(coll0_)
        return ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oydcocpi"))) if True else 904)) != ((len(iife0_()
        )) + (len(_dafny.SeqWithoutIsStrInference([820]))))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1_i0_ in range(default__.abs(176))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')]))

    @staticmethod
    def fm9(p0, globalState):
        return True

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return ((_dafny.Map({False: True}))) | (_dafny.Map({True: True}))

    @staticmethod
    def fm11(p0, globalState):
        return (D0_DC2(not(True), len(_dafny.Set({185})), _dafny.SeqWithoutIsStrInference([True, False, True]), 273, len(_dafny.SeqWithoutIsStrInference([not(True), False])))).cf6

    @staticmethod
    def fm12(p0, p1, globalState):
        if True:
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(492, 328):
                    d_2_v0_: int = compr_1_
                    if ((492) <= (d_2_v0_)) and ((d_2_v0_) < (328)):
                        coll1_[(d_2_v0_) + (956)] = len(_dafny.SeqWithoutIsStrInference([True]))
                return _dafny.Map(coll1_)
            return _dafny.MultiSet([676, len(iife1_()
            ), len(_dafny.Map({False: True})), 49])
        elif True:
            return _dafny.MultiSet([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([-771]): len(_dafny.SeqWithoutIsStrInference([True, True]))}))])

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return _dafny.CodePoint('y')

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        return ((_dafny.Map({False: len(_dafny.Map({True: _dafny.CodePoint('i')}))})) | (_dafny.Map({False: 491}))) | ((_dafny.Map({False: len(_dafny.Set({False}))})) | (_dafny.Map({not(not(not(True))): 644})))

    @staticmethod
    def fm17(p0, globalState):
        return (_dafny.Set({True, not(False), True, True})) - ((_dafny.Set({not(True), not(True), False})).intersection(_dafny.Set({True})))

    @staticmethod
    def fm18(p0, globalState):
        return _dafny.MultiSet([D3_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsnhlv")))])

    @staticmethod
    def fm19(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_3_i0_ in range(default__.abs(818))])).Elements:
                d_4_v0_: int = compr_2_
                if (d_4_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_3_i0_ in range(default__.abs(818))])):
                    coll2_[default__.safeModuloInt(d_4_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkhneayht"))))] = 583
            return _dafny.Map(coll2_)
        return (iife2_()
        ) | ((_dafny.Map({-379: 887})) | (_dafny.Map({len(_dafny.Map({False: D5_DC17((0) - ((_dafny.MultiSet([True])).cardinality), (_dafny.MultiSet([-257, 187, 911, 21])).cardinality, D0_DC1(True, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, False, False])): False})): _dafny.CodePoint('w')}))])), 729])), 798), False, True)})): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality})))

    @staticmethod
    def fm20(p0, p1, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False, True]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, True, True, True])))).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([not(False), False, False, False])) + (_dafny.SeqWithoutIsStrInference([False, True]))))

    @staticmethod
    def fm21(p0, globalState):
        if not((192) == (358)):
            return D0_DC0(_dafny.MultiSet([False, False]))
        elif True:
            return D0_DC0(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False])))

    @staticmethod
    def fm22(p0, p1, globalState):
        return _dafny.Map({(_dafny.MultiSet([False, True]) if True else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D5_DC17((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([not(False)])))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "naexdwq"))), D0_DC1(True, len(_dafny.SeqWithoutIsStrInference([376])), len(_dafny.Map({True: len(_dafny.Map({752: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "famhky")))}))}))), True, True)).cf33]))): (889) - (len(_dafny.Map({_dafny.CodePoint('l'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwit")))})))})

    @staticmethod
    def fm23(globalState):
        source0_ = D1_DC4(_dafny.SeqWithoutIsStrInference([D0_DC2(True, len(_dafny.Set({True, False})), _dafny.SeqWithoutIsStrInference([True]), 467, -669), D0_DC2(True, 195, _dafny.SeqWithoutIsStrInference([False]), -382, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqavgwus"))))]))
        if source0_.is_DC5:
            d_5___mcc_h0_ = source0_.cf11
            d_6_cf11_ = d_5___mcc_h0_
            return D2_DC8(231, (0) - (-575))
        elif source0_.is_DC4:
            d_7___mcc_h1_ = source0_.cf10
            d_8_cf10_ = d_7___mcc_h1_
            return D2_DC8(527, 49)
        elif True:
            d_9___mcc_h2_ = source0_.cf12
            d_10_cf12_ = d_9___mcc_h2_
            return D2_DC8(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etcx"))), len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djn")): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdomg")))})))

    @staticmethod
    def fm24(globalState):
        return (_dafny.SeqWithoutIsStrInference([D0_DC2(True, (_dafny.MultiSet([False, True, True, True, not(False)])).cardinality, _dafny.SeqWithoutIsStrInference([False]), 74, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrcsueoj"))))])) + (_dafny.SeqWithoutIsStrInference([D0_DC2(False, (0) - (len(_dafny.SeqWithoutIsStrInference([-536, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jva")))]))), _dafny.SeqWithoutIsStrInference([False, False]), 546, -596) for d_11_i0_ in range(default__.abs(762))]))

    @staticmethod
    def fm25(globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(660, 885):
                d_12_v0_: int = compr_3_
                if ((660) <= (d_12_v0_)) and ((d_12_v0_) < (885)):
                    coll3_[(d_12_v0_) - (641)] = False
            return _dafny.Map(coll3_)
        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, (D0_DC2(False, len(_dafny.SeqWithoutIsStrInference([541])), _dafny.SeqWithoutIsStrInference([True]), len(_dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.Set({True}))})), 810])))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ats"))))).cf4, not(not(False)), True])): False})) | (iife3_()
        )

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        return D1_DC5(not (False) or (False))

    @staticmethod
    def fm27(p0, p1, globalState):
        return D0_DC1((False) == (False), 192, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnqykwbvy"))) if False else 443))

    @staticmethod
    def fm28(p0, p1, globalState):
        return _dafny.Set({(_dafny.Set({(D0_DC1(True, 2, -271)).cf1, False, False, True, True})).intersection(_dafny.Set({not(True)})), (_dafny.Set({True, False})) | (_dafny.Set({False, True})), (_dafny.Set({False})) - (_dafny.Set({True, True})), (_dafny.Set({not(True)})) | (_dafny.Set({True}))})

    @staticmethod
    def Main(noArgsParameter__):
        d_13_v0_: int
        d_13_v0_ = 888
        d_14_v1_: _dafny.Seq
        d_14_v1_ = _dafny.SeqWithoutIsStrInference([d_13_v0_, d_13_v0_])
        d_15_v2_: _dafny.Seq
        d_15_v2_ = _dafny.SeqWithoutIsStrInference([d_13_v0_, len(d_14_v1_)])
        d_16_v3_: _dafny.Array
        nw0_ = _dafny.Array(int(0), 11)
        d_16_v3_ = nw0_
        d_17_v4_: _dafny.Set
        d_17_v4_ = _dafny.Set({d_13_v0_})
        d_18_v5_: _dafny.Map
        d_18_v5_ = _dafny.Map({d_17_v4_: d_13_v0_})
        d_19_v6_: _dafny.Array
        def lambda0_(d_20_i0_):
            return True

        init0_ = lambda0_
        nw1_ = _dafny.Array(None, 14)
        for i0_0_ in range(nw1_.length(0)):
            nw1_[i0_0_] = init0_(i0_0_)
        d_19_v6_ = nw1_
        d_21_v7_: _dafny.Seq
        d_21_v7_ = _dafny.SeqWithoutIsStrInference([d_16_v3_])
        d_22_v8_: _dafny.Array
        nw2_ = _dafny.Array(None, 3)
        nw2_[int(0)] = d_16_v3_
        nw2_[int(1)] = d_16_v3_
        nw2_[int(2)] = (d_21_v7_)[default__.safeIndex(491, len(d_21_v7_))]
        d_22_v8_ = nw2_
        d_23_globalState_: GlobalState
        nw3_ = GlobalState()
        nw3_.ctor__(True, d_15_v2_, 957, d_16_v3_, d_18_v5_, 116, -467, _dafny.Map({d_19_v6_: d_13_v0_}), True, d_22_v8_, d_16_v3_, True, False, -352, d_16_v3_, 833, False, True, 294)
        d_23_globalState_ = nw3_
        d_24_v9_: bool
        d_24_v9_ = True
        hi0_ = 379
        for d_25_i1_ in range(default__.fm0(d_24_v9_, d_24_v9_, d_23_globalState_), hi0_):
            (d_23_globalState_).f18 = (0) - ((0) - (default__.fm0(d_24_v9_, d_24_v9_, d_23_globalState_)))
            d_26_v10_: C0
            nw4_ = C0()
            nw4_.ctor__(d_13_v0_)
            d_26_v10_ = nw4_
            d_27_v11_: C2
            nw5_ = C2()
            nw5_.ctor__(d_26_v10_, d_24_v9_, d_19_v6_, d_24_v9_)
            d_27_v11_ = nw5_
            d_28_v12_: _dafny.Map
            d_28_v12_ = _dafny.Map({not(d_27_v11_.f26): d_24_v9_})
            d_24_v9_ = not((_dafny.Map({d_27_v11_.f26: default__.fm9(d_24_v9_, d_23_globalState_)})) == ((d_28_v12_) | (d_28_v12_)))
            (d_23_globalState_).f13 = (0) - (d_25_i1_)
        (d_23_globalState_).f0 = default__.fm7(d_13_v0_, d_13_v0_, d_13_v0_, d_23_globalState_)
        d_29_v13_: str
        d_29_v13_ = _dafny.CodePoint('j')
        d_30_v14_: _dafny.Seq
        d_30_v14_ = _dafny.SeqWithoutIsStrInference([d_29_v13_])
        d_31_v15_: C3
        nw6_ = C3()
        nw6_.ctor__(default__.fm8(d_30_v14_, _dafny.Map({d_13_v0_: d_24_v9_}), d_13_v0_, d_23_globalState_), d_19_v6_)
        d_31_v15_ = nw6_
        d_32_v16_: _dafny.Seq
        d_32_v16_ = _dafny.SeqWithoutIsStrInference([d_31_v15_, d_31_v15_, d_31_v15_, d_31_v15_])
        hi1_ = len(d_17_v4_)
        for d_33_i2_ in range((d_13_v0_) * (len((d_32_v16_).set(default__.safeIndex(d_13_v0_, len(d_32_v16_)), d_31_v15_))), hi1_):
            (d_23_globalState_).f0 = not (False) or (d_24_v9_)
            (d_31_v15_).f22 = d_31_v15_.f22
            d_34_v17_: _dafny.Map
            d_34_v17_ = _dafny.Map({d_24_v9_: d_24_v9_})
            d_34_v17_ = d_34_v17_
            (d_23_globalState_).f0 = not((d_13_v0_) != (d_13_v0_))
        d_35_i3_: int
        d_35_i3_ = 0
        with _dafny.label("0"):
            while default__.fm2(d_24_v9_, d_13_v0_, d_13_v0_, d_23_globalState_):
                with _dafny.c_label("0"):
                    if (d_35_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_35_i3_ = (d_35_i3_) + (1)
                    d_36_v18_: _dafny.Array
                    def lambda1_(d_37_v1_):
                        def lambda2_(d_38_i4_):
                            return d_37_v1_

                        return lambda2_

                    init1_ = lambda1_(d_14_v1_)
                    nw7_ = _dafny.Array(None, 2)
                    for i0_1_ in range(nw7_.length(0)):
                        nw7_[i0_1_] = init1_(i0_1_)
                    d_36_v18_ = nw7_
                    d_39_v19_: D4
                    d_39_v19_ = D4_DC12(d_29_v13_)
                    d_40_v20_: _dafny.Map
                    d_40_v20_ = _dafny.Map({d_36_v18_: d_39_v19_})
                    d_40_v20_ = (d_40_v20_).set(d_36_v18_, d_39_v19_)
                    d_41_v21_: _dafny.Seq
                    d_41_v21_ = _dafny.SeqWithoutIsStrInference([d_24_v9_])
                    d_42_v22_: T0
                    nw8_ = C1()
                    nw8_.ctor__(d_31_v15_.f22, (d_31_v15_).f23, d_24_v9_)
                    d_42_v22_ = nw8_
                    d_43_v23_: _dafny.Map
                    d_43_v23_ = _dafny.Map({d_42_v22_: d_13_v0_})
                    d_44_v24_: C4
                    nw9_ = C4()
                    nw9_.ctor__((d_31_v15_).fm5(default__.fm7(len(d_41_v21_), ((d_43_v23_)[d_42_v22_] if (d_42_v22_) in (d_43_v23_) else d_13_v0_), len(d_30_v14_), d_23_globalState_), d_23_globalState_), d_19_v6_, (d_42_v22_).f20)
                    d_44_v24_ = nw9_
                    d_44_v24_ = d_44_v24_
                    index0_ = default__.safeIndex(254, ((d_31_v15_).f23).length(0))
                    ((d_31_v15_).f23)[index0_] = (d_24_v9_) == ((d_42_v22_).f20)
                    d_45_v25_: _dafny.Map
                    d_45_v25_ = _dafny.Map({d_13_v0_: d_24_v9_})
                    index1_ = default__.safeIndex(254, ((d_31_v15_).f23).length(0))
                    ((d_31_v15_).f23)[index1_] = (d_30_v14_) != (default__.fm8(d_30_v14_, d_45_v25_, (d_44_v24_).f21, d_23_globalState_))
                    d_46_v26_: _dafny.Array
                    nw10_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
                    d_46_v26_ = nw10_
                    index2_ = default__.safeIndex(233, (d_46_v26_).length(0))
                    (d_46_v26_)[index2_] = d_31_v15_.f22
                    index3_ = default__.safeIndex(233, (d_46_v26_).length(0))
                    (d_46_v26_)[index3_] = d_31_v15_.f22
                    pass
            pass
        d_47_v27_: _dafny.Seq
        d_48_v28_: bool
        out0_: _dafny.Seq
        out1_: bool
        out0_, out1_ = (d_31_v15_).m2(d_23_globalState_)
        d_47_v27_ = out0_
        d_48_v28_ = out1_
        d_49_i5_: int
        d_49_i5_ = 0
        with _dafny.label("1"):
            while d_24_v9_:
                with _dafny.c_label("1"):
                    if (d_49_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_49_i5_ = (d_49_i5_) + (1)
                    (d_23_globalState_).f0 = (default__.safeModuloInt(d_13_v0_, 155)) <= (845)
                    d_50_v29_: _dafny.Map
                    d_50_v29_ = _dafny.Map({d_48_v28_: d_48_v28_})
                    (d_23_globalState_).f18 = default__.safeDivisionInt(d_13_v0_, (d_31_v15_).fm5(((d_50_v29_)[d_48_v28_] if (d_48_v28_) in (d_50_v29_) else d_24_v9_), d_23_globalState_))
                    if (d_13_v0_) not in (d_14_v1_):
                        (d_23_globalState_).f13 = d_13_v0_
                        d_51_v30_: T0
                        nw11_ = C1()
                        nw11_.ctor__(_dafny.SeqWithoutIsStrInference([d_29_v13_ for d_52_i6_ in range(default__.abs(963))]), (d_31_v15_).f23, d_24_v9_)
                        d_51_v30_ = nw11_
                        d_53_v31_: _dafny.Seq
                        d_53_v31_ = _dafny.SeqWithoutIsStrInference([d_51_v30_, d_51_v30_])
                        rhs0_ = (_dafny.SeqWithoutIsStrInference([d_51_v30_, d_51_v30_, d_51_v30_])) + (d_53_v31_)
                        rhs1_ = (0) - (d_13_v0_)
                        lhs0_ = d_23_globalState_
                        d_53_v31_ = rhs0_
                        lhs0_.f18 = rhs1_
                        d_54_v32_: _dafny.Seq
                        d_55_v33_: bool
                        out2_: _dafny.Seq
                        out3_: bool
                        out2_, out3_ = (d_31_v15_).m2(d_23_globalState_)
                        d_54_v32_ = out2_
                        d_55_v33_ = out3_
                        (d_23_globalState_).f18 = (default__.fm0(d_24_v9_, d_24_v9_, d_23_globalState_)) * (default__.safeModuloInt(d_13_v0_, d_13_v0_))
                        d_30_v14_ = (d_30_v14_ if (d_13_v0_) <= (d_13_v0_) else d_31_v15_.f22)
                    elif True:
                        d_56_v34_: C1
                        nw12_ = C1()
                        nw12_.ctor__((d_30_v14_) + (default__.fm8(_dafny.SeqWithoutIsStrInference([default__.fm4(d_23_globalState_)]), _dafny.Map({d_13_v0_: d_24_v9_}), len(d_30_v14_), d_23_globalState_)), (d_31_v15_).f23, d_24_v9_)
                        d_56_v34_ = nw12_
                        d_48_v28_ = d_48_v28_
                        d_57_v35_: C1
                        nw13_ = C1()
                        nw13_.ctor__(d_30_v14_, (d_31_v15_).f23, d_24_v9_)
                        d_57_v35_ = nw13_
                        d_58_v36_: C0
                        nw14_ = C0()
                        nw14_.ctor__(len(_dafny.SeqWithoutIsStrInference([len(d_47_v27_), d_13_v0_])))
                        d_58_v36_ = nw14_
                        d_59_v37_: C2
                        nw15_ = C2()
                        nw15_.ctor__(d_58_v36_, d_24_v9_, (d_31_v15_).f23, d_48_v28_)
                        d_59_v37_ = nw15_
                        d_60_v38_: D8
                        d_60_v38_ = D8_DC22(d_59_v37_)
                        d_61_v39_: _dafny.Set
                        d_61_v39_ = _dafny.Set({(d_60_v38_).cf42, d_59_v37_, d_59_v37_, d_59_v37_, d_59_v37_})
                        d_61_v39_ = (d_61_v39_).intersection(d_61_v39_)
                        d_62_v40_: bool
                        d_63_v41_: int
                        d_64_v42_: int
                        d_65_v43_: int
                        out4_: bool
                        out5_: int
                        out6_: int
                        out7_: int
                        out4_, out5_, out6_, out7_ = (d_59_v37_).m4(((d_58_v36_).f24) < ((d_58_v36_).f24), d_23_globalState_)
                        d_62_v40_ = out4_
                        d_63_v41_ = out5_
                        d_64_v42_ = out6_
                        d_65_v43_ = out7_
                    (d_31_v15_).m3(d_23_globalState_)
                    pass
            pass
        d_66_v44_: _dafny.Array
        nw16_ = _dafny.Array(_dafny.MultiSet({}), 1)
        d_66_v44_ = nw16_
        d_67_v45_: _dafny.MultiSet
        d_67_v45_ = _dafny.MultiSet([d_13_v0_, -409])
        index4_ = default__.safeIndex(957, (d_66_v44_).length(0))
        (d_66_v44_)[index4_] = d_67_v45_
        index5_ = default__.safeIndex(957, (d_66_v44_).length(0))
        (d_66_v44_)[index5_] = (d_67_v45_).intersection((d_67_v45_) - (d_67_v45_))
        d_68_v46_: _dafny.Map
        d_68_v46_ = _dafny.Map({d_24_v9_: d_13_v0_})
        d_69_v47_: D1
        d_69_v47_ = D1_DC6(default__.fm26(d_13_v0_, d_24_v9_, d_24_v9_, d_68_v46_, d_23_globalState_))
        d_70_v48_: D1
        d_70_v48_ = D1_DC6(d_69_v47_)
        pat_let_tv0_ = d_17_v4_
        pat_let_tv1_ = d_17_v4_
        pat_let_tv2_ = d_17_v4_
        pat_let_tv3_ = d_66_v44_
        pat_let_tv4_ = d_66_v44_
        pat_let_tv5_ = d_66_v44_
        pat_let_tv6_ = d_66_v44_
        def lambda3_(source1_):
            if source1_.is_DC5:
                d_71___mcc_h0_ = source1_.cf11
                d_72_cf11_ = d_71___mcc_h0_
                return pat_let_tv0_
            elif source1_.is_DC4:
                d_73___mcc_h1_ = source1_.cf10
                d_74_cf10_ = d_73___mcc_h1_
                return pat_let_tv1_
            elif True:
                d_75___mcc_h2_ = source1_.cf12
                d_76_cf12_ = d_75___mcc_h2_
                def iife4_():
                    coll4_ = _dafny.Set()
                    compr_4_: int
                    for compr_4_ in ((pat_let_tv4_)[default__.safeIndex(957, (pat_let_tv3_).length(0))]).Elements:
                        d_77_v49_: int = compr_4_
                        if (d_77_v49_) in ((pat_let_tv6_)[default__.safeIndex(957, (pat_let_tv5_).length(0))]):
                            coll4_ = coll4_.union(_dafny.Set([(d_77_v49_) * (len(_dafny.SeqWithoutIsStrInference([258])))]))
                    return _dafny.Set(coll4_)
                return (pat_let_tv2_) | (iife4_()
                )

        d_17_v4_ = lambda3_(d_70_v48_)
        index6_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
        ((d_31_v15_).f23)[index6_] = False
        d_78_v50_: C0
        nw17_ = C0()
        nw17_.ctor__(d_13_v0_)
        d_78_v50_ = nw17_
        d_79_v51_: T0
        nw18_ = C2()
        nw18_.ctor__(d_78_v50_, False, (d_31_v15_).f23, False)
        d_79_v51_ = nw18_
        d_80_v52_: _dafny.Map
        d_80_v52_ = _dafny.Map({d_79_v51_: default__.fm7((0) - (len(d_30_v14_)), (0) - ((d_78_v50_).f24), (d_78_v50_).f24, d_23_globalState_)})
        d_81_v53_: _dafny.MultiSet
        d_81_v53_ = _dafny.MultiSet([d_29_v13_, _dafny.CodePoint('o'), default__.fm15((d_78_v50_).f24, (d_79_v51_).f20, (d_78_v50_).f24, d_23_globalState_)])
        index7_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
        ((d_31_v15_).f23)[index7_] = (len(d_80_v52_)) != ((d_81_v53_).cardinality)
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(712, 713):
                d_82_v54_: int = compr_5_
                if ((712) <= (d_82_v54_)) and ((d_82_v54_) < (713)):
                    coll5_[default__.safeDivisionInt(d_82_v54_, (0) - ((0) - (d_13_v0_)))] = D3_DC11(d_13_v0_, d_13_v0_, d_48_v28_)
            return _dafny.Map(coll5_)
        source2_ = default__.fm27(not(d_48_v28_), iife5_()
        , d_23_globalState_)
        if source2_.is_DC1:
            d_83___mcc_h3_ = source2_.cf1
            d_84___mcc_h4_ = source2_.cf2
            d_85___mcc_h5_ = source2_.cf3
            d_86_cf3_ = d_85___mcc_h5_
            d_87_cf2_ = d_84___mcc_h4_
            d_88_cf1_ = d_83___mcc_h3_
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(828, 775):
                    d_89_v55_: int = compr_6_
                    if ((828) <= (d_89_v55_)) and ((d_89_v55_) < (775)):
                        coll6_[(d_89_v55_) - ((d_78_v50_).f24)] = _dafny.Set({d_29_v13_})
                return _dafny.Map(coll6_)
            d_14_v1_ = (d_15_v2_) + ((d_15_v2_) + (_dafny.SeqWithoutIsStrInference([648, len(iife6_()
            )])))
            d_86_cf3_ = d_87_cf2_
            d_90_v56_: C2
            nw19_ = C2()
            nw19_.ctor__(d_78_v50_, (d_79_v51_).f20, (d_31_v15_).f23, ((d_31_v15_).f23)[default__.safeIndex(23, ((d_31_v15_).f23).length(0))])
            d_90_v56_ = nw19_
            if (d_90_v56_.f26) == ((d_79_v51_).f20):
                d_91_v57_: _dafny.MultiSet
                d_91_v57_ = _dafny.MultiSet([(d_79_v51_).f20])
                d_87_cf2_ = (0) - (default__.safeDivisionInt(((d_91_v57_).set(d_90_v56_.f26, default__.abs((d_14_v1_)[default__.safeIndex(len(_dafny.Map({d_88_cf1_: True})), len(d_14_v1_))]))).cardinality, (d_78_v50_).f24))
                d_92_v58_: _dafny.Map
                d_92_v58_ = _dafny.Map({d_13_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvadfbvx"))})
                d_93_v59_: _dafny.Map
                d_93_v59_ = _dafny.Map({d_92_v58_: (len(d_31_v15_.f22)) - ((d_78_v50_).f24)})
                d_93_v59_ = (d_93_v59_).set(_dafny.Map({(d_78_v50_).f24: d_31_v15_.f22}), (d_87_cf2_) + ((d_78_v50_).f24))
                d_94_v60_: _dafny.Array
                def lambda4_(d_95_v4_):
                    def lambda5_(d_96_i7_):
                        return d_95_v4_

                    return lambda5_

                init2_ = lambda4_(d_17_v4_)
                nw20_ = _dafny.Array(None, 26)
                for i0_2_ in range(nw20_.length(0)):
                    nw20_[i0_2_] = init2_(i0_2_)
                d_94_v60_ = nw20_
                index8_ = default__.safeIndex(725, (d_94_v60_).length(0))
                (d_94_v60_)[index8_] = d_17_v4_
                d_97_v61_: _dafny.Set
                d_97_v61_ = _dafny.Set({d_13_v0_})
                index9_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
                index10_ = default__.safeIndex(725, (d_94_v60_).length(0))
                rhs2_ = default__.safeModuloInt((d_78_v50_).f24, d_13_v0_)
                rhs3_ = (d_78_v50_).f24
                rhs4_ = (d_79_v51_).f20
                rhs5_ = (d_97_v61_)
                lhs1_ = d_23_globalState_
                lhs2_ = (d_31_v15_).f23
                lhs3_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
                lhs4_ = d_94_v60_
                lhs5_ = default__.safeIndex(725, (d_94_v60_).length(0))
                d_86_cf3_ = rhs2_
                lhs1_.f13 = rhs3_
                lhs2_[lhs3_] = rhs4_
                lhs4_[lhs5_] = rhs5_
                d_98_v62_: _dafny.Seq
                d_98_v62_ = _dafny.SeqWithoutIsStrInference([(d_31_v15_).f23])
                d_99_v63_: C3
                nw21_ = C3()
                nw21_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfcjm")), (d_98_v62_)[default__.safeIndex(d_13_v0_, len(d_98_v62_))])
                d_99_v63_ = nw21_
                d_100_v64_: C1
                nw22_ = C1()
                nw22_.ctor__(d_31_v15_.f22, (d_31_v15_).f23, True)
                d_100_v64_ = nw22_
                d_101_v65_: _dafny.Seq
                d_101_v65_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_100_v64_, d_100_v64_])])
                d_17_v4_ = ((d_94_v60_)[default__.safeIndex(725, (d_94_v60_).length(0))]) - (_dafny.Set({len((d_101_v65_)[default__.safeIndex((d_78_v50_).f24, len(d_101_v65_))])}))
            elif True:
                d_102_v66_: _dafny.Array
                def lambda6_(d_103_v15_):
                    def lambda7_(d_104_i8_):
                        return d_103_v15_.f22

                    return lambda7_

                init3_ = lambda6_(d_31_v15_)
                nw23_ = _dafny.Array(None, 21)
                for i0_3_ in range(nw23_.length(0)):
                    nw23_[i0_3_] = init3_(i0_3_)
                d_102_v66_ = nw23_
                index11_ = default__.safeIndex(746, (d_102_v66_).length(0))
                (d_102_v66_)[index11_] = d_30_v14_
                index12_ = default__.safeIndex(746, (d_102_v66_).length(0))
                (d_102_v66_)[index12_] = d_31_v15_.f22
                d_105_v67_: _dafny.Set
                d_105_v67_ = _dafny.Set({d_19_v6_})
                d_105_v67_ = (d_105_v67_).intersection(d_105_v67_)
                d_24_v9_ = d_24_v9_
                (d_90_v56_).f26 = not((not((_dafny.SeqWithoutIsStrInference([d_87_cf2_, (d_78_v50_).f24])) <= (d_14_v1_))) not in (_dafny.Set({(d_79_v51_).f20, d_88_cf1_, d_90_v56_.f26})))
                d_106_v68_: _dafny.Array
                def lambda8_(d_107_v27_, d_108_v50_):
                    def lambda9_(d_109_i9_):
                        return (d_107_v27_).set(default__.safeIndex((d_108_v50_).f24, len(d_107_v27_)), True)

                    return lambda9_

                init4_ = lambda8_(d_47_v27_, d_78_v50_)
                nw24_ = _dafny.Array(None, 1)
                for i0_4_ in range(nw24_.length(0)):
                    nw24_[i0_4_] = init4_(i0_4_)
                d_106_v68_ = nw24_
                d_106_v68_ = d_106_v68_
        elif source2_.is_DC2:
            d_110___mcc_h6_ = source2_.cf4
            d_111___mcc_h7_ = source2_.cf5
            d_112___mcc_h8_ = source2_.cf6
            d_113___mcc_h9_ = source2_.cf7
            d_114___mcc_h10_ = source2_.cf8
            d_115_cf8_ = d_114___mcc_h10_
            d_116_cf7_ = d_113___mcc_h9_
            d_117_cf6_ = d_112___mcc_h8_
            d_118_cf5_ = d_111___mcc_h7_
            d_119_cf4_ = d_110___mcc_h6_
            (d_23_globalState_).f0 = d_119_cf4_
            index13_ = default__.safeIndex(447, (d_16_v3_).length(0))
            (d_16_v3_)[index13_] = (d_78_v50_).f24
            index14_ = default__.safeIndex(447, (d_16_v3_).length(0))
            (d_16_v3_)[index14_] = (d_78_v50_).f24
            (d_23_globalState_).f18 = (d_78_v50_).f24
            (d_31_v15_).f22 = d_31_v15_.f22
        elif source2_.is_DC3:
            d_120___mcc_h11_ = source2_.cf9
            d_121_cf9_ = d_120___mcc_h11_
            d_19_v6_ = (d_31_v15_).f23
            d_122_v69_: _dafny.Set
            d_122_v69_ = _dafny.Set({(d_79_v51_).f20})
            d_123_v70_: _dafny.Seq
            d_123_v70_ = _dafny.SeqWithoutIsStrInference([d_122_v69_, (d_122_v69_) - (d_122_v69_), (d_122_v69_).intersection(d_122_v69_), (default__.fm17(len(d_31_v15_.f22), d_23_globalState_)).intersection(d_122_v69_)])
            index15_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
            index16_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
            rhs6_ = d_123_v70_
            rhs7_ = (d_79_v51_).f20
            rhs8_ = not(d_24_v9_)
            rhs9_ = True
            rhs10_ = ((d_31_v15_).f23)[default__.safeIndex(23, ((d_31_v15_).f23).length(0))]
            lhs6_ = d_23_globalState_
            lhs7_ = (d_31_v15_).f23
            lhs8_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
            lhs9_ = (d_31_v15_).f23
            lhs10_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
            d_123_v70_ = rhs6_
            lhs6_.f0 = rhs7_
            lhs7_[lhs8_] = rhs8_
            lhs9_[lhs10_] = rhs9_
            d_24_v9_ = rhs10_
            d_124_v71_: C0
            nw25_ = C0()
            nw25_.ctor__((d_13_v0_) - (d_13_v0_))
            d_124_v71_ = nw25_
            (d_31_v15_).m3(d_23_globalState_)
        elif True:
            d_125___mcc_h12_ = source2_.cf0
            d_126_cf0_ = d_125___mcc_h12_
            d_127_v72_: _dafny.Map
            d_127_v72_ = _dafny.Map({default__.safeModuloInt(len(default__.fm19(d_24_v9_, d_23_globalState_)), (d_78_v50_).f24): False})
            d_127_v72_ = (d_127_v72_).set(len((d_30_v14_).set(default__.safeIndex((d_31_v15_).fm5((d_79_v51_).f20, d_23_globalState_), len(d_30_v14_)), d_29_v13_)), not(d_48_v28_))
            d_128_v73_: _dafny.Set
            d_128_v73_ = _dafny.Set({not((d_79_v51_).f20), d_24_v9_, d_48_v28_, d_24_v9_})
            d_129_v74_: D4
            d_129_v74_ = D4_DC14(d_24_v9_, d_128_v73_)
            d_129_v74_ = d_129_v74_
            d_130_v75_: D1
            d_130_v75_ = D1_DC5((d_79_v51_).f20)
            (d_23_globalState_).f0 = default__.fm2((d_130_v75_).cf11, (d_13_v0_ if (d_79_v51_).f20 else (d_78_v50_).f24), ((d_78_v50_).f24) * ((d_78_v50_).f24), d_23_globalState_)
            index17_ = default__.safeIndex(897, (d_16_v3_).length(0))
            (d_16_v3_)[index17_] = (d_78_v50_).f24
            d_131_v76_: _dafny.Set
            d_131_v76_ = _dafny.Set({d_31_v15_})
            index18_ = default__.safeIndex(897, (d_16_v3_).length(0))
            (d_16_v3_)[index18_] = ((((d_126_cf0_)[d_24_v9_] if (d_24_v9_) in (d_126_cf0_) else len(d_17_v4_))) + (len(d_131_v76_))) + ((0) - (d_13_v0_))
        d_15_v2_ = d_14_v1_
        (d_31_v15_).m3(d_23_globalState_)
        d_132_v77_: _dafny.Map
        d_132_v77_ = _dafny.Map({_dafny.Map({d_31_v15_.f22: d_79_v51_}): ((d_31_v15_).f23)[default__.safeIndex(23, ((d_31_v15_).f23).length(0))]})
        d_133_v78_: _dafny.Map
        d_133_v78_ = _dafny.Map({(d_79_v51_).f20: d_79_v51_})
        d_134_v79_: _dafny.Map
        d_134_v79_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqqlffck")): ((d_133_v78_)[(d_79_v51_).f20] if ((d_79_v51_).f20) in (d_133_v78_) else d_79_v51_)})
        d_135_v80_: C4
        nw26_ = C4()
        nw26_.ctor__(default__.safeModuloInt(d_13_v0_, d_13_v0_), d_19_v6_, ((d_132_v77_)[d_134_v79_] if (d_134_v79_) in (d_132_v77_) else d_48_v28_))
        d_135_v80_ = nw26_
        index19_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
        rhs11_ = d_29_v13_
        rhs12_ = d_13_v0_
        rhs13_ = d_135_v80_
        rhs14_ = d_48_v28_
        lhs11_ = d_23_globalState_
        lhs12_ = (d_31_v15_).f23
        lhs13_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
        d_29_v13_ = rhs11_
        lhs11_.f13 = rhs12_
        d_135_v80_ = rhs13_
        lhs12_[lhs13_] = rhs14_
        d_16_v3_ = d_16_v3_
        hi2_ = d_13_v0_
        for d_136_i10_ in range(d_13_v0_, hi2_):
            d_137_v81_: C2
            nw27_ = C2()
            nw27_.ctor__(d_78_v50_, (d_79_v51_).f20, (d_79_v51_).f19, True)
            d_137_v81_ = nw27_
            d_138_v82_: _dafny.Seq
            d_138_v82_ = _dafny.SeqWithoutIsStrInference([d_137_v81_, d_137_v81_])
            d_139_v83_: _dafny.Map
            d_139_v83_ = _dafny.Map({d_13_v0_: d_138_v82_})
            d_140_v84_: _dafny.Seq
            d_140_v84_ = _dafny.SeqWithoutIsStrInference([(d_138_v82_) + (d_138_v82_), d_138_v82_, ((d_139_v83_)[d_13_v0_] if (d_13_v0_) in (d_139_v83_) else d_138_v82_)])
            d_140_v84_ = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_137_v81_]), d_138_v82_, d_138_v82_, d_138_v82_, (d_140_v84_)[default__.safeIndex(d_13_v0_, len(d_140_v84_))]])) + ((d_140_v84_) + (d_140_v84_))
            d_141_v85_: D0
            d_141_v85_ = D0_DC2(d_137_v81_.f26, len(_dafny.Set({d_31_v15_})), _dafny.SeqWithoutIsStrInference([((d_31_v15_).f23)[default__.safeIndex(23, ((d_31_v15_).f23).length(0))]]), (d_135_v80_).f21, 945)
            d_142_v86_: D0
            d_142_v86_ = D0_DC2((d_141_v85_).cf4, d_136_i10_, (d_47_v27_).set(default__.safeIndex(d_136_i10_, len(d_47_v27_)), d_24_v9_), (d_135_v80_).f21, (d_78_v50_).f24)
            d_143_v87_: _dafny.Map
            d_143_v87_ = _dafny.Map({d_31_v15_: d_137_v81_})
            d_144_v88_: _dafny.Seq
            d_144_v88_ = _dafny.SeqWithoutIsStrInference([d_141_v85_, d_141_v85_, d_141_v85_, D0_DC2(((d_31_v15_).f23)[default__.safeIndex(23, ((d_31_v15_).f23).length(0))], d_136_i10_, d_47_v27_, (0) - (len(d_30_v14_)), d_136_i10_), D0_DC2(d_137_v81_.f26, d_136_i10_, _dafny.SeqWithoutIsStrInference([False]), len(d_143_v87_), len(d_15_v2_))])
            source3_ = D1_DC4(d_144_v88_)
            if source3_.is_DC5:
                d_145___mcc_h13_ = source3_.cf11
                d_146_cf11_ = d_145___mcc_h13_
                nw28_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_22_v8_ = nw28_
                (d_23_globalState_).f13 = ((d_31_v15_).fm6((d_79_v51_).f20, (_dafny.Map({len(d_47_v27_): 40})).set((d_135_v80_).f21, (d_135_v80_).f21), _dafny.SeqWithoutIsStrInference([d_29_v13_ for d_147_i11_ in range(default__.abs(433))]), False, d_23_globalState_) if (d_79_v51_).f20 else (d_78_v50_).f24)
                d_148_v89_: C1
                nw29_ = C1()
                nw29_.ctor__(d_31_v15_.f22, d_19_v6_, ((d_31_v15_).f23)[default__.safeIndex(23, ((d_31_v15_).f23).length(0))])
                d_148_v89_ = nw29_
                d_149_v90_: _dafny.Map
                d_149_v90_ = _dafny.Map({(D8_DC23(d_48_v28_, len(d_15_v2_), (d_135_v80_).f21, d_24_v9_)).cf45: d_148_v89_})
                (d_23_globalState_).f18 = (((d_78_v50_).f24) * (len(d_149_v90_))) * (d_13_v0_)
                d_13_v0_ = 75
            elif source3_.is_DC4:
                d_150___mcc_h14_ = source3_.cf10
                d_151_cf10_ = d_150___mcc_h14_
                d_152_v91_: _dafny.Seq
                d_152_v91_ = _dafny.SeqWithoutIsStrInference([d_135_v80_])
                d_153_v92_: _dafny.Array
                nw30_ = _dafny.Array(None, 8)
                nw30_[int(0)] = d_135_v80_
                nw30_[int(1)] = d_135_v80_
                nw30_[int(2)] = d_135_v80_
                nw30_[int(3)] = d_135_v80_
                nw30_[int(4)] = d_135_v80_
                nw30_[int(5)] = d_135_v80_
                nw30_[int(6)] = (d_152_v91_)[default__.safeIndex((d_78_v50_).f24, len(d_152_v91_))]
                nw30_[int(7)] = d_135_v80_
                d_153_v92_ = nw30_
                index20_ = default__.safeIndex(949, (d_153_v92_).length(0))
                (d_153_v92_)[index20_] = d_135_v80_
                index21_ = default__.safeIndex(949, (d_153_v92_).length(0))
                (d_153_v92_)[index21_] = d_135_v80_
                (d_23_globalState_).f0 = ((d_31_v15_).f23)[default__.safeIndex(23, ((d_31_v15_).f23).length(0))]
                d_154_v93_: bool
                d_155_v94_: int
                d_156_v95_: int
                d_157_v96_: int
                out8_: bool
                out9_: int
                out10_: int
                out11_: int
                out8_, out9_, out10_, out11_ = (d_137_v81_).m4((d_79_v51_).f20, d_23_globalState_)
                d_154_v93_ = out8_
                d_155_v94_ = out9_
                d_156_v95_ = out10_
                d_157_v96_ = out11_
                d_24_v9_ = (d_79_v51_).f20
            elif True:
                d_158___mcc_h15_ = source3_.cf12
                d_159_cf12_ = d_158___mcc_h15_
                index22_ = default__.safeIndex(23, ((d_31_v15_).f23).length(0))
                ((d_31_v15_).f23)[index22_] = d_48_v28_
                d_160_v97_: C2
                nw31_ = C2()
                nw31_.ctor__(d_137_v81_.f25, (d_79_v51_).f20, (d_79_v51_).f19, d_137_v81_.f26)
                d_160_v97_ = nw31_
                d_161_v98_: _dafny.Set
                d_161_v98_ = _dafny.Set({(d_79_v51_).f20})
                d_162_v99_: _dafny.Set
                d_162_v99_ = _dafny.Set({(_dafny.Set({(d_79_v51_).f20})) - (d_161_v98_)})
                d_162_v99_ = default__.fm28(d_136_i10_, default__.safeModuloInt(default__.fm0((d_79_v51_).f20, d_160_v97_.f26, d_23_globalState_), (d_135_v80_).f21), d_23_globalState_)
                (d_23_globalState_).f13 = d_13_v0_
            d_163_v100_: bool
            d_164_v101_: int
            d_165_v102_: int
            d_166_v103_: int
            out12_: bool
            out13_: int
            out14_: int
            out15_: int
            out12_, out13_, out14_, out15_ = (d_137_v81_).m4(not (d_48_v28_) or (d_137_v81_.f26), d_23_globalState_)
            d_163_v100_ = out12_
            d_164_v101_ = out13_
            d_165_v102_ = out14_
            d_166_v103_ = out15_
            index23_ = default__.safeIndex(659, (d_16_v3_).length(0))
            (d_16_v3_)[index23_] = d_166_v103_
            index24_ = default__.safeIndex(659, (d_16_v3_).length(0))
            rhs15_ = (d_31_v15_)
            rhs16_ = d_166_v103_
            lhs14_ = d_16_v3_
            lhs15_ = default__.safeIndex(659, (d_16_v3_).length(0))
            d_31_v15_ = rhs15_
            lhs14_[lhs15_] = rhs16_
        d_167_v104_: _dafny.Map
        d_167_v104_ = _dafny.Map({d_13_v0_: default__.fm7(d_13_v0_, d_13_v0_, len(d_30_v14_), d_23_globalState_)})
        d_168_v107_: D0
        d_168_v107_ = D0_DC1(True, (d_135_v80_).f21, (d_78_v50_).f24)
        d_169_v108_: _dafny.MultiSet
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(929, 294):
                d_170_v105_: int = compr_7_
                if ((929) <= (d_170_v105_)) and ((d_170_v105_) < (294)):
                    coll7_[(d_170_v105_) + ((d_135_v80_).f21)] = ((d_31_v15_).f23)[default__.safeIndex(23, ((d_31_v15_).f23).length(0))]
            return _dafny.Map(coll7_)
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in (d_14_v1_).Elements:
                d_171_v106_: int = compr_8_
                if (d_171_v106_) in (d_14_v1_):
                    coll8_[(d_171_v106_) * ((d_135_v80_).f21)] = d_48_v28_
            return _dafny.Map(coll8_)
        d_169_v108_ = _dafny.MultiSet([(d_167_v104_) | (_dafny.Map({(d_78_v50_).f24: d_48_v28_})), iife7_()
        , (iife8_()
        ) | ((d_167_v104_).set((d_78_v50_).f24, (d_168_v107_).cf1))])
        (d_23_globalState_).f13 = ((d_169_v108_)[(_dafny.Map({(0) - ((d_78_v50_).f24): not(d_48_v28_)})) | (d_167_v104_)] if ((_dafny.Map({(0) - ((d_78_v50_).f24): not(d_48_v28_)})) | (d_167_v104_)) in (d_169_v108_) else (d_135_v80_).f21)
        _dafny.print(_dafny.string_of(d_13_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_14_v1_) == (_dafny.SeqWithoutIsStrInference([888, 2, 888, 2, 648, 0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_15_v2_) == (_dafny.SeqWithoutIsStrInference([888, 2, 888, 2, 648, 0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_17_v4_) == (_dafny.Set({888}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_v5_) == (_dafny.Map({_dafny.Set({888}): 888}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_19_v6_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_21_v7_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_23_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_23_globalState_).f1) == (_dafny.SeqWithoutIsStrInference([888, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_.f4) == (_dafny.Map({_dafny.Set({888}): 888}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_23_globalState_).f7)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_23_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_23_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_23_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_24_v9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_29_v13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_30_v14_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_31_v15_.f22) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('d')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_31_v15_).f23)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_32_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_35_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_v27_) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_48_v28_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_49_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v44_)[0]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_67_v45_) == (_dafny.MultiSet([888, -409]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_68_v46_) == (_dafny.Map({True: 888}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_69_v47_).cf12).cf11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_70_v48_).cf12).cf12).cf11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v50_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_v51_).f19)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_v51_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_80_v52_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_81_v53_) == (_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('o'), _dafny.CodePoint('y')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_132_v77_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_133_v78_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_134_v79_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v80_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_167_v104_) == (_dafny.Map({888: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v107_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v107_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v107_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_v108_) == (_dafny.MultiSet([_dafny.Map({888: True}), _dafny.Map({}), _dafny.Map({0: True, 888: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC5(D1, NamedTuple('DC5', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(_dafny.Array(None, 0), _dafny.Map({}), _dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({self.cf16.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC13(D4, NamedTuple('DC13', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC17(int(0), int(0), D0.default()(), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC17(D5, NamedTuple('DC17', [('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

class D6_DC20(D6, NamedTuple('DC20', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC21(D7, NamedTuple('DC21', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(False, int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC23(D8, NamedTuple('DC23', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC26(D9, NamedTuple('DC26', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)

class D10_DC27(D10, NamedTuple('DC27', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)

class D11_DC28(D11, NamedTuple('DC28', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f3: _dafny.Array = _dafny.Array(None, 0)
        self.f4: _dafny.Map = _dafny.Map({})
        self.f10: _dafny.Array = _dafny.Array(None, 0)
        self.f13: int = int(0)
        self.f18: int = int(0)
        self._f1: _dafny.Seq = _dafny.Seq({})
        self._f2: int = int(0)
        self._f5: int = int(0)
        self._f6: int = int(0)
        self._f7: _dafny.Map = _dafny.Map({})
        self._f8: bool = False
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f11: bool = False
        self._f12: bool = False
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f15: int = int(0)
        self._f16: bool = False
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self).f18 = f18

    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17

class C0:
    def  __init__(self):
        self._f24: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f24):
        (self)._f24 = f24

    @property
    def f24(self):
        return self._f24

class C1(T0):
    def  __init__(self):
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f20: bool = False
        self._f27: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f27, f19, f20):
        (self)._f27 = f27
        (self)._f19 = f19
        (self)._f20 = f20

    def fm13(self, globalState):
        return (self).f20

    def fm14(self, p0, globalState):
        return (_dafny.Map({D1_DC4(_dafny.SeqWithoutIsStrInference([D0_DC2((self).f20, len(_dafny.Map({(self).f20: len((self).f27)})), _dafny.SeqWithoutIsStrInference([not((self).f20), (self).f20]), 572, len(_dafny.Map({293: D0_DC3(_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f20, (self).f20])): (_dafny.MultiSet([(self).f20, (self).f20, (self).f20, (self).f20])).cardinality}))}))) for d_172_i0_ in range(default__.abs(780))])): (self).f20})) | ((_dafny.Map({D1_DC4(_dafny.SeqWithoutIsStrInference([D0_DC2((self).f20, 756, _dafny.SeqWithoutIsStrInference([(self).f20]), len(_dafny.SeqWithoutIsStrInference([319])), 449) for d_173_i1_ in range(default__.abs(398))])): (self).f20})) | (_dafny.Map({D1_DC4(_dafny.SeqWithoutIsStrInference([D0_DC2((self).f20, len(_dafny.Map({(self).f20: (D0_DC1((self).f20, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len((self).f27): (_dafny.MultiSet([(self).f20])).cardinality})), -88])), len((self).f27))).cf1})), _dafny.SeqWithoutIsStrInference([False]), len(_dafny.Map({951: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhnf")))})), len(_dafny.Map({(self).f20: (self).f20}))) for d_174_i2_ in range(default__.abs(802))])): not((self).f20)})))

    @property
    def f27(self):
        return self._f27

class C2(T0):
    def  __init__(self):
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f20: bool = False
        self.f25: C0 = None
        self.f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f25, f26, f19, f20):
        (self).f25 = f25
        (self).f26 = f26
        (self)._f19 = f19
        (self)._f20 = f20

    def m4(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        r3: int = int(0)
        d_175_v0_: _dafny.Seq
        d_175_v0_ = _dafny.SeqWithoutIsStrInference([(self.f25).f24])
        hi3_ = ((self.f25).f24) + (len(d_175_v0_))
        for d_176_i0_ in range((self.f25).f24, hi3_):
            d_177_v1_: _dafny.Seq
            d_177_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')])
            d_178_v2_: _dafny.Map
            d_178_v2_ = _dafny.Map({424: (self).f20})
            d_177_v1_ = default__.fm8(d_177_v1_, d_178_v2_, d_176_i0_, globalState)
            r1 = (0) - (((self.f25).f24) + ((self.f25).f24))
            d_179_v3_: _dafny.Seq
            d_179_v3_ = _dafny.SeqWithoutIsStrInference([self.f26, p0])
            d_180_v4_: C0
            nw32_ = C0()
            nw32_.ctor__(len(d_179_v3_))
            d_180_v4_ = nw32_
            (globalState).f0 = self.f26
        d_181_v5_: _dafny.Seq
        d_181_v5_ = _dafny.SeqWithoutIsStrInference([self.f26])
        d_182_v6_: _dafny.Map
        d_182_v6_ = _dafny.Map({p0: (self.f25).f24})
        d_183_v7_: _dafny.MultiSet
        d_183_v7_ = _dafny.MultiSet([d_182_v6_, d_182_v6_])
        d_184_v8_: D0
        d_184_v8_ = D0_DC2(True, (self.f25).f24, d_181_v5_, ((d_183_v7_)[d_182_v6_] if (d_182_v6_) in (d_183_v7_) else len(_dafny.Set({137}))), (self.f25).f24)
        d_185_v9_: _dafny.Seq
        d_185_v9_ = _dafny.SeqWithoutIsStrInference([d_184_v8_, d_184_v8_])
        d_186_v10_: _dafny.Map
        d_186_v10_ = _dafny.Map({len(_dafny.Set({self.f26, (self).f20})): (self.f25).f24})
        d_187_v11_: _dafny.Map
        d_187_v11_ = _dafny.Map({len(d_186_v10_): p0})
        d_188_v12_: D1
        d_188_v12_ = D1_DC4(_dafny.SeqWithoutIsStrInference([D0_DC2((self).f20, (self.f25).f24, d_181_v5_, (0) - ((self.f25).f24), len(d_187_v11_))]))
        d_189_v13_: _dafny.Array
        nw33_ = _dafny.Array(None, 2)
        nw33_[int(0)] = (D1_DC4((d_185_v9_).set(default__.safeIndex((self.f25).f24, len(d_185_v9_)), d_184_v8_)) if (self).f20 else d_188_v12_)
        nw33_[int(1)] = d_188_v12_
        d_189_v13_ = nw33_
        pat_let_tv7_ = d_185_v9_
        index25_ = default__.safeIndex(549, (d_189_v13_).length(0))
        def iife9_(_pat_let0_0):
            def iife10_(d_190_dt__update__tmp_h0_):
                def iife11_(_pat_let1_0):
                    def iife12_(d_191_dt__update_hcf10_h0_):
                        return D1_DC4(d_191_dt__update_hcf10_h0_)
                    return iife12_(_pat_let1_0)
                return iife11_(pat_let_tv7_)
            return iife10_(_pat_let0_0)
        (d_189_v13_)[index25_] = iife9_(d_188_v12_)
        index26_ = default__.safeIndex(549, (d_189_v13_).length(0))
        (d_189_v13_)[index26_] = d_188_v12_
        if (self).f20:
            r0 = default__.fm9((d_181_v5_)[default__.safeIndex(-479, len(d_181_v5_))], globalState)
            d_192_v14_: _dafny.Set
            d_192_v14_ = _dafny.Set({(self).f20, self.f26})
            d_193_v15_: D1
            d_193_v15_ = D1_DC5(False)
            d_194_v16_: _dafny.MultiSet
            d_194_v16_ = _dafny.MultiSet([d_193_v15_])
            d_195_v17_: _dafny.Map
            d_195_v17_ = _dafny.Map({d_192_v14_: d_194_v16_})
            d_196_v18_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_196_v18_ = nw34_
            d_197_v19_: _dafny.Seq
            d_197_v19_ = _dafny.SeqWithoutIsStrInference([d_196_v18_, d_196_v18_])
            d_198_v20_: _dafny.Map
            d_198_v20_ = _dafny.Map({d_195_v17_: (d_197_v19_)[default__.safeIndex(686, len(d_197_v19_))]})
            d_199_v21_: _dafny.Map
            d_199_v21_ = _dafny.Map({(self.f25).f24: d_192_v14_})
            d_198_v20_ = (d_198_v20_).set(_dafny.Map({((d_199_v21_)[(self.f25).f24] if ((self.f25).f24) in (d_199_v21_) else d_192_v14_): _dafny.MultiSet([d_193_v15_, D1_DC5((self).f20), d_193_v15_])}), d_196_v18_)
            (globalState).f13 = ((self.f25).f24) * ((self.f25).f24)
            d_200_v22_: C0
            nw35_ = C0()
            nw35_.ctor__((self.f25).f24)
            d_200_v22_ = nw35_
            d_201_v23_: _dafny.Array
            nw36_ = _dafny.Array(int(0), 7)
            d_201_v23_ = nw36_
            index27_ = default__.safeIndex(629, (d_201_v23_).length(0))
            (d_201_v23_)[index27_] = (d_200_v22_).f24
            index28_ = default__.safeIndex(629, (d_201_v23_).length(0))
            (d_201_v23_)[index28_] = (self.f25).f24
        elif True:
            d_202_v24_: _dafny.Array
            nw37_ = _dafny.Array(_dafny.Map({}), 28)
            d_202_v24_ = nw37_
            d_203_v25_: _dafny.Map
            d_203_v25_ = _dafny.Map({(self).f20: (self).f20})
            index29_ = default__.safeIndex(576, (d_202_v24_).length(0))
            (d_202_v24_)[index29_] = (_dafny.Map({p0: self.f26})) | (d_203_v25_)
            d_204_v26_: _dafny.Seq
            d_204_v26_ = _dafny.SeqWithoutIsStrInference([(default__.fm10((self.f25).f24, len(d_175_v0_), (self.f25).f24, globalState)) | (d_203_v25_)])
            d_205_v27_: _dafny.MultiSet
            d_205_v27_ = _dafny.MultiSet([p0, not(self.f26)])
            d_206_v28_: _dafny.Array
            nw38_ = _dafny.Array(int(0), 27)
            d_206_v28_ = nw38_
            d_207_v29_: _dafny.MultiSet
            d_207_v29_ = _dafny.MultiSet([d_206_v28_])
            d_208_v30_: _dafny.Set
            d_208_v30_ = _dafny.Set({(self.f25).f24, ((d_207_v29_)[d_206_v28_] if (d_206_v28_) in (d_207_v29_) else (self.f25).f24)})
            d_209_v31_: _dafny.Map
            d_209_v31_ = _dafny.Map({(0) - ((d_205_v27_).cardinality): d_208_v30_})
            d_210_v32_: _dafny.Seq
            d_210_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
            index30_ = default__.safeIndex(576, (d_202_v24_).length(0))
            (d_202_v24_)[index30_] = (d_204_v26_)[default__.safeIndex((len(d_209_v31_)) * (len(d_210_v32_)), len(d_204_v26_))]
            d_186_v10_ = (d_186_v10_) | ((d_186_v10_ if p0 else (d_186_v10_).set(len((_dafny.SeqWithoutIsStrInference([(self.f25).f24])).set(default__.safeIndex((self.f25).f24, len(_dafny.SeqWithoutIsStrInference([(self.f25).f24]))), (self.f25).f24)), (self.f25).f24)))
            d_211_v33_: D2
            d_211_v33_ = D2_DC7(_dafny.SeqWithoutIsStrInference([(self.f25).f24]))
            d_212_v34_: _dafny.Set
            d_212_v34_ = _dafny.Set({d_211_v33_})
            d_213_v35_: _dafny.Seq
            d_213_v35_ = _dafny.SeqWithoutIsStrInference([d_212_v34_, d_212_v34_])
            (self).f26 = (_dafny.Set({d_211_v33_, d_211_v33_})).isdisjoint((d_213_v35_)[default__.safeIndex(len(d_210_v32_), len(d_213_v35_))])
            d_214_v36_: _dafny.Array
            def lambda10_(d_215_i1_):
                return _dafny.CodePoint('i')

            init5_ = lambda10_
            nw39_ = _dafny.Array(None, 16)
            for i0_5_ in range(nw39_.length(0)):
                nw39_[i0_5_] = init5_(i0_5_)
            d_214_v36_ = nw39_
            d_214_v36_ = d_214_v36_
            (globalState).f0 = (self).f20
        d_216_v37_: _dafny.Array
        def lambda11_(d_217_i2_):
            return (d_217_i2_) + ((self.f25).f24)

        init6_ = lambda11_
        nw40_ = _dafny.Array(None, 8)
        for i0_6_ in range(nw40_.length(0)):
            nw40_[i0_6_] = init6_(i0_6_)
        d_216_v37_ = nw40_
        index31_ = default__.safeIndex(296, (d_216_v37_).length(0))
        (d_216_v37_)[index31_] = len(default__.fm11((self.f25).f24, globalState))
        index32_ = default__.safeIndex(296, (d_216_v37_).length(0))
        (d_216_v37_)[index32_] = (self.f25).f24
        d_218_v38_: _dafny.Seq
        d_218_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmahe"))
        d_219_v39_: str
        d_219_v39_ = _dafny.CodePoint('m')
        d_220_v40_: D3
        d_220_v40_ = D3_DC9(d_218_v38_)
        d_221_v42_: _dafny.Set
        d_221_v42_ = _dafny.Set({(d_216_v37_)[default__.safeIndex(296, (d_216_v37_).length(0))]})
        d_222_v43_: _dafny.Array
        nw41_ = _dafny.Array(None, 17)
        nw41_[int(0)] = d_218_v38_
        nw41_[int(1)] = d_218_v38_
        nw41_[int(2)] = (d_218_v38_) + (d_218_v38_)
        nw41_[int(3)] = d_218_v38_
        nw41_[int(4)] = d_218_v38_
        nw41_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnp"))
        nw41_[int(6)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqfjf"))).set(default__.safeIndex((self.f25).f24, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqfjf")))), d_219_v39_)) + (d_218_v38_)
        nw41_[int(7)] = d_218_v38_
        nw41_[int(8)] = (d_220_v40_).cf16
        nw41_[int(9)] = (d_218_v38_) + (d_218_v38_)
        nw41_[int(10)] = (d_218_v38_).set(default__.safeIndex((self.f25).f24, len(d_218_v38_)), d_219_v39_)
        nw41_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbiytnt"))
        nw41_[int(12)] = d_218_v38_
        nw41_[int(13)] = d_218_v38_
        nw41_[int(14)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
        nw41_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uusrltbmy"))
        def iife13_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in (d_175_v0_).Elements:
                d_224_v41_: int = compr_9_
                if (d_224_v41_) in (d_175_v0_):
                    coll9_[(d_224_v41_) * ((self.f25).f24)] = (self).f20
            return _dafny.Map(coll9_)
        nw41_[int(16)] = default__.fm8(_dafny.SeqWithoutIsStrInference([d_219_v39_ for d_223_i4_ in range(default__.abs(88))]), (iife13_()
        ).set(len(d_221_v42_), self.f26), (d_216_v37_)[default__.safeIndex(296, (d_216_v37_).length(0))], globalState)
        d_222_v43_ = nw41_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_222_v43_).length(0)):
            d_225_i3_: int = guard_loop_0_
            if (True) and (((0) <= (d_225_i3_)) and ((d_225_i3_) < ((d_222_v43_).length(0)))):
                (d_222_v43_)[(d_225_i3_)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (d_218_v38_)) + (_dafny.SeqWithoutIsStrInference([d_219_v39_ for d_226_i5_ in range(default__.abs(40))]))
        hi4_ = (d_216_v37_)[default__.safeIndex(296, (d_216_v37_).length(0))]
        for d_227_i6_ in range(((self.f25).f24) - ((self.f25).f24), hi4_):
            r0 = self.f26
            (globalState).f13 = ((self.f25).f24 if (self.f26) == (not((self).f20)) else 128)
            d_228_v44_: _dafny.Map
            d_228_v44_ = _dafny.Map({d_227_i6_: (self).f19})
            d_229_v45_: _dafny.Array
            nw42_ = _dafny.Array(None, 13)
            nw42_[int(0)] = (self).f19
            nw42_[int(1)] = (self).f19
            nw42_[int(2)] = (self).f19
            nw42_[int(3)] = (self).f19
            nw42_[int(4)] = ((d_228_v44_)[default__.fm0(p0, self.f26, globalState)] if (default__.fm0(p0, self.f26, globalState)) in (d_228_v44_) else (self).f19)
            nw42_[int(5)] = (self).f19
            nw42_[int(6)] = (self).f19
            nw42_[int(7)] = (self).f19
            nw42_[int(8)] = (self).f19
            nw42_[int(9)] = (self).f19
            nw42_[int(10)] = ((self).f19 if (self).f20 else (self).f19)
            nw42_[int(11)] = (self).f19
            nw42_[int(12)] = (self).f19
            d_229_v45_ = nw42_
            index33_ = default__.safeIndex(716, (d_229_v45_).length(0))
            (d_229_v45_)[index33_] = (self).f19
            d_230_v46_: _dafny.Set
            d_230_v46_ = _dafny.Set({p0})
            d_231_v47_: _dafny.Map
            d_231_v47_ = _dafny.Map({p0: p0})
            d_232_v48_: _dafny.Map
            d_232_v48_ = _dafny.Map({len(d_175_v0_): d_231_v47_})
            index34_ = default__.safeIndex(716, (d_229_v45_).length(0))
            nw43_ = _dafny.Array(None, 7)
            nw43_[int(0)] = self.f26
            nw43_[int(1)] = (d_230_v46_).ispropersubset(d_230_v46_)
            nw43_[int(2)] = (d_219_v39_) in (d_218_v38_)
            nw43_[int(3)] = (self).f20
            nw43_[int(4)] = (self).f20
            nw43_[int(5)] = (d_221_v42_).isdisjoint(_dafny.Set({len(((d_232_v48_)[120] if (120) in (d_232_v48_) else d_231_v47_))}))
            nw43_[int(6)] = p0
            (d_229_v45_)[index34_] = nw43_
            d_233_v49_: _dafny.MultiSet
            d_233_v49_ = _dafny.MultiSet([default__.fm0(self.f26, self.f26, globalState), (self.f25).f24, (self.f25).f24])
            d_234_v50_: _dafny.MultiSet
            d_234_v50_ = _dafny.MultiSet([self.f26, p0])
            (globalState).f0 = (len(d_230_v46_)) in (default__.fm12(d_233_v49_, ((d_234_v50_)[not((self).f20)] if (not((self).f20)) in (d_234_v50_) else (self.f25).f24), globalState))
        r0 = default__.fm9(p0, globalState)
        r1 = 893
        r2 = ((d_182_v6_)[not(((self.f25).f24) >= ((self.f25).f24))] if (not(((self.f25).f24) >= ((self.f25).f24))) in (d_182_v6_) else ((d_186_v10_)[(self.f25).f24] if ((self.f25).f24) in (d_186_v10_) else (self.f25).f24))
        r3 = (d_216_v37_)[default__.safeIndex(296, (d_216_v37_).length(0))]
        return r0, r1, r2, r3

    def m5(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_235_v0_: _dafny.Seq
        d_235_v0_ = _dafny.SeqWithoutIsStrInference([self.f26])
        d_236_v1_: _dafny.Seq
        d_236_v1_ = _dafny.SeqWithoutIsStrInference([D0_DC2((self).f20, (self.f25).f24, d_235_v0_, (self.f25).f24, (self.f25).f24)])
        d_237_v2_: D1
        d_237_v2_ = D1_DC4(d_236_v1_)
        pat_let_tv8_ = d_236_v1_
        def iife14_(_pat_let2_0):
            def iife15_(d_238_dt__update__tmp_h0_):
                def iife16_(_pat_let3_0):
                    def iife17_(d_239_dt__update_hcf10_h0_):
                        return D1_DC4(d_239_dt__update_hcf10_h0_)
                    return iife17_(_pat_let3_0)
                return iife16_(pat_let_tv8_)
            return iife15_(_pat_let2_0)
        source4_ = iife14_(d_237_v2_)
        if source4_.is_DC5:
            d_240___mcc_h0_ = source4_.cf11
            d_241_cf11_ = d_240___mcc_h0_
            d_242_v3_: C0
            nw44_ = C0()
            nw44_.ctor__(default__.fm0((self).f20, d_241_cf11_, globalState))
            d_242_v3_ = nw44_
            d_243_v4_: D1
            d_243_v4_ = D1_DC5(default__.fm9(self.f26, globalState))
            d_244_v5_: D1
            d_244_v5_ = D1_DC6(d_243_v4_)
            d_244_v5_ = d_244_v5_
            d_241_cf11_ = d_241_cf11_
            d_245_v6_: _dafny.Seq
            d_245_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tclfl"))
            d_246_v7_: T0
            nw45_ = C1()
            nw45_.ctor__(d_245_v6_, (self).f19, d_241_cf11_)
            d_246_v7_ = nw45_
            d_247_v8_: _dafny.Map
            d_247_v8_ = _dafny.Map({(self).f20: d_246_v7_})
            d_247_v8_ = (d_247_v8_).set(True, (d_246_v7_ if d_241_cf11_ else d_246_v7_))
        elif source4_.is_DC4:
            d_248___mcc_h1_ = source4_.cf10
            d_249_cf10_ = d_248___mcc_h1_
            d_250_v9_: _dafny.Array
            nw46_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_250_v9_ = nw46_
            d_251_v10_: str
            d_251_v10_ = _dafny.CodePoint('i')
            d_252_v11_: _dafny.MultiSet
            d_252_v11_ = _dafny.MultiSet([(self).f20, self.f26])
            d_253_v12_: _dafny.Array
            nw47_ = _dafny.Array(None, 26)
            nw47_[int(0)] = d_251_v10_
            nw47_[int(1)] = d_251_v10_
            nw47_[int(2)] = d_251_v10_
            nw47_[int(3)] = d_251_v10_
            nw47_[int(4)] = d_251_v10_
            nw47_[int(5)] = d_251_v10_
            nw47_[int(6)] = d_251_v10_
            nw47_[int(7)] = d_251_v10_
            nw47_[int(8)] = d_251_v10_
            nw47_[int(9)] = d_251_v10_
            nw47_[int(10)] = d_251_v10_
            nw47_[int(11)] = d_251_v10_
            nw47_[int(12)] = d_251_v10_
            nw47_[int(13)] = d_251_v10_
            nw47_[int(14)] = d_251_v10_
            nw47_[int(15)] = _dafny.CodePoint('x')
            nw47_[int(16)] = d_251_v10_
            nw47_[int(17)] = default__.fm15((d_252_v11_).cardinality, self.f26, (self.f25).f24, globalState)
            nw47_[int(18)] = d_251_v10_
            nw47_[int(19)] = d_251_v10_
            nw47_[int(20)] = _dafny.CodePoint('i')
            nw47_[int(21)] = _dafny.CodePoint('u')
            nw47_[int(22)] = d_251_v10_
            nw47_[int(23)] = d_251_v10_
            nw47_[int(24)] = d_251_v10_
            nw47_[int(25)] = d_251_v10_
            d_253_v12_ = nw47_
            index35_ = default__.safeIndex(508, (d_250_v9_).length(0))
            (d_250_v9_)[index35_] = d_253_v12_
            index36_ = default__.safeIndex(508, (d_250_v9_).length(0))
            (d_250_v9_)[index36_] = d_253_v12_
            d_254_v13_: _dafny.Array
            nw48_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
            d_254_v13_ = nw48_
            d_254_v13_ = d_254_v13_
            d_255_v14_: _dafny.Seq
            d_255_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            d_256_v15_: D3
            d_256_v15_ = D3_DC9(d_255_v14_)
            d_257_v16_: D3
            d_257_v16_ = D3_DC9((_dafny.SeqWithoutIsStrInference([d_251_v10_ for d_258_i0_ in range(default__.abs(273))])) + ((d_256_v15_).cf16))
            source5_ = d_256_v15_
            if source5_.is_DC10:
                d_259___mcc_h3_ = source5_.cf17
                d_260___mcc_h4_ = source5_.cf18
                d_261___mcc_h5_ = source5_.cf19
                d_262___mcc_h6_ = source5_.cf20
                d_263_cf20_ = d_262___mcc_h6_
                d_264_cf19_ = d_261___mcc_h5_
                d_265_cf18_ = d_260___mcc_h4_
                d_266_cf17_ = d_259___mcc_h3_
                d_267_v17_: _dafny.Seq
                d_267_v17_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_251_v10_ for d_268_i1_ in range(default__.abs(132))])), (self.f25).f24])
                d_269_v18_: _dafny.Map
                d_269_v18_ = _dafny.Map({(d_267_v17_) + (d_267_v17_): ((self.f25).f24) * (len(d_255_v14_))})
                d_269_v18_ = (d_269_v18_).set(d_267_v17_, (self.f25).f24)
                d_270_v19_: _dafny.Map
                d_270_v19_ = _dafny.Map({d_251_v10_: (self.f25).f24})
                d_271_v20_: _dafny.MultiSet
                d_271_v20_ = _dafny.MultiSet([(self.f25).f24, len((d_265_cf18_).set(self.f26, 850)), ((d_252_v11_)[(self).f20] if ((self).f20) in (d_252_v11_) else d_263_cf20_), (self.f25).f24, ((d_270_v19_)[d_251_v10_] if (d_251_v10_) in (d_270_v19_) else (self.f25).f24)])
                r1 = (d_271_v20_).isdisjoint((d_271_v20_).intersection(_dafny.MultiSet([len(d_265_cf18_), (self.f25).f24])))
                index37_ = default__.safeIndex(103, ((self).f19).length(0))
                ((self).f19)[index37_] = (self).f20
                index38_ = default__.safeIndex(103, ((self).f19).length(0))
                ((self).f19)[index38_] = not((self).f20)
                d_272_v21_: _dafny.Set
                d_272_v21_ = _dafny.Set({(d_251_v10_) in (d_255_v14_), ((self).f20) and (((self).f19)[default__.safeIndex(103, ((self).f19).length(0))])})
                d_272_v21_ = _dafny.Set({((self).f19)[default__.safeIndex(103, ((self).f19).length(0))], (d_235_v0_)[default__.safeIndex((self.f25).f24, len(d_235_v0_))], True})
            elif source5_.is_DC11:
                d_273___mcc_h7_ = source5_.cf21
                d_274___mcc_h8_ = source5_.cf22
                d_275___mcc_h9_ = source5_.cf23
                d_276_cf23_ = d_275___mcc_h9_
                d_277_cf22_ = d_274___mcc_h8_
                d_278_cf21_ = d_273___mcc_h7_
                d_277_cf22_ = len((d_235_v0_).set(default__.safeIndex(default__.safeModuloInt(d_278_cf21_, (self.f25).f24), len(d_235_v0_)), (d_277_cf22_) <= ((self.f25).f24)))
                index39_ = default__.safeIndex(317, ((self).f19).length(0))
                ((self).f19)[index39_] = ((self).f20) and (self.f26)
                d_279_v22_: C1
                nw49_ = C1()
                nw49_.ctor__(d_255_v14_, (self).f19, d_276_cf23_)
                d_279_v22_ = nw49_
                d_280_v23_: _dafny.Seq
                d_280_v23_ = _dafny.SeqWithoutIsStrInference([d_279_v22_, d_279_v22_])
                d_281_v24_: _dafny.Array
                nw50_ = _dafny.Array(None, 26)
                nw50_[int(0)] = (self.f25).f24
                nw50_[int(1)] = (self.f25).f24
                nw50_[int(2)] = 858
                nw50_[int(3)] = len(d_255_v14_)
                nw50_[int(4)] = 46
                nw50_[int(5)] = (self.f25).f24
                nw50_[int(6)] = d_278_cf21_
                nw50_[int(7)] = (self.f25).f24
                nw50_[int(8)] = len(d_255_v14_)
                nw50_[int(9)] = (self.f25).f24
                nw50_[int(10)] = (_dafny.MultiSet([d_251_v10_, d_251_v10_])).cardinality
                nw50_[int(11)] = (self.f25).f24
                nw50_[int(12)] = len(d_280_v23_)
                nw50_[int(13)] = 479
                nw50_[int(14)] = (self.f25).f24
                nw50_[int(15)] = 708
                nw50_[int(16)] = (self.f25).f24
                nw50_[int(17)] = d_277_cf22_
                nw50_[int(18)] = 777
                nw50_[int(19)] = d_278_cf21_
                nw50_[int(20)] = (self.f25).f24
                nw50_[int(21)] = (self.f25).f24
                nw50_[int(22)] = len(d_235_v0_)
                nw50_[int(23)] = d_277_cf22_
                nw50_[int(24)] = 29
                nw50_[int(25)] = (self.f25).f24
                d_281_v24_ = nw50_
                d_282_v25_: D0
                d_282_v25_ = D0_DC0(d_252_v11_)
                d_283_v26_: _dafny.Map
                d_283_v26_ = _dafny.Map({(self.f25).f24: d_282_v25_})
                d_284_v27_: _dafny.Map
                d_284_v27_ = _dafny.Map({d_281_v24_: d_283_v26_})
                index40_ = default__.safeIndex(317, ((self).f19).length(0))
                ((self).f19)[index40_] = (d_284_v27_) != (d_284_v27_)
                d_285_v28_: _dafny.Seq
                d_285_v28_ = _dafny.SeqWithoutIsStrInference([default__.fm8(d_255_v14_, _dafny.Map({944: self.f26}), (self.f25).f24, globalState)])
                d_286_v29_: C0
                nw51_ = C0()
                nw51_.ctor__(len(d_285_v28_))
                d_286_v29_ = nw51_
                d_287_v30_: D4
                d_287_v30_ = D4_DC12(d_251_v10_)
                d_251_v10_ = (d_287_v30_).cf24
            elif True:
                d_288___mcc_h10_ = source5_.cf16
                d_289_cf16_ = d_288___mcc_h10_
                d_290_v31_: C1
                nw52_ = C1()
                nw52_.ctor__(d_255_v14_, (self).f19, ((self.f25).f24) > ((self.f25).f24))
                d_290_v31_ = nw52_
                d_290_v31_ = d_290_v31_
                d_291_v32_: _dafny.Map
                d_291_v32_ = _dafny.Map({d_290_v31_: self.f26})
                d_292_v33_: _dafny.Map
                d_292_v33_ = _dafny.Map({len(d_291_v32_): self.f26})
                d_293_v34_: _dafny.Map
                d_293_v34_ = _dafny.Map({len(d_292_v33_): (self.f25).f24})
                d_293_v34_ = (d_293_v34_).set((self.f25).f24, ((self.f25).f24) + (((d_252_v11_)[self.f26] if (self.f26) in (d_252_v11_) else len((d_290_v31_).f27))))
                (globalState).f0 = self.f26
                d_294_v35_: _dafny.Set
                d_294_v35_ = _dafny.Set({(self.f25).f24})
                index41_ = default__.safeIndex(666, ((self).f19).length(0))
                ((self).f19)[index41_] = not((670) != (((d_252_v11_)[(self).f20] if ((self).f20) in (d_252_v11_) else len(d_294_v35_))))
                d_295_v36_: _dafny.Array
                nw53_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_295_v36_ = nw53_
                d_296_v37_: _dafny.Seq
                d_296_v37_ = _dafny.SeqWithoutIsStrInference([(self.f25).f24])
                d_297_v38_: _dafny.Map
                d_297_v38_ = _dafny.Map({(len(d_296_v37_) if (self).f20 else (self.f25).f24): d_295_v36_})
                d_298_v39_: _dafny.Seq
                d_298_v39_ = _dafny.SeqWithoutIsStrInference([d_255_v14_])
                d_299_v40_: _dafny.Map
                d_299_v40_ = _dafny.Map({d_298_v39_: d_295_v36_})
                index42_ = default__.safeIndex(666, ((self).f19).length(0))
                rhs17_ = default__.fm0((d_252_v11_).issubset(d_252_v11_), self.f26, globalState)
                rhs18_ = (self).f20
                rhs19_ = (self.f25).f24
                rhs20_ = ((d_297_v38_)[(self.f25).f24] if ((self.f25).f24) in (d_297_v38_) else ((d_299_v40_)[_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mti"))])] if (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mti"))])) in (d_299_v40_) else d_295_v36_))
                lhs16_ = globalState
                lhs17_ = (self).f19
                lhs18_ = default__.safeIndex(666, ((self).f19).length(0))
                lhs19_ = globalState
                lhs16_.f18 = rhs17_
                lhs17_[lhs18_] = rhs18_
                lhs19_.f18 = rhs19_
                d_295_v36_ = rhs20_
            (globalState).f0 = self.f26
        elif True:
            d_300___mcc_h2_ = source4_.cf12
            d_301_cf12_ = d_300___mcc_h2_
            index43_ = default__.safeIndex(597, ((self).f19).length(0))
            ((self).f19)[index43_] = (self).f20
            d_302_v41_: _dafny.Seq
            d_302_v41_ = _dafny.SeqWithoutIsStrInference([(self.f25).f24])
            index44_ = default__.safeIndex(597, ((self).f19).length(0))
            ((self).f19)[index44_] = (d_302_v41_) < (d_302_v41_)
            d_303_v42_: _dafny.Map
            d_303_v42_ = _dafny.Map({((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]: self.f25})
            d_304_v43_: _dafny.Map
            d_304_v43_ = _dafny.Map({self.f26: (_dafny.Map({((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]: self.f25})).set((self).f20, self.f25)})
            d_305_v44_: _dafny.Array
            nw54_ = _dafny.Array(None, 28)
            nw54_[int(0)] = d_303_v42_
            nw54_[int(1)] = d_303_v42_
            nw54_[int(2)] = d_303_v42_
            nw54_[int(3)] = d_303_v42_
            nw54_[int(4)] = _dafny.Map({True: self.f25})
            nw54_[int(5)] = d_303_v42_
            nw54_[int(6)] = d_303_v42_
            nw54_[int(7)] = d_303_v42_
            nw54_[int(8)] = d_303_v42_
            nw54_[int(9)] = d_303_v42_
            nw54_[int(10)] = d_303_v42_
            nw54_[int(11)] = d_303_v42_
            nw54_[int(12)] = d_303_v42_
            nw54_[int(13)] = (d_303_v42_).set(((self).f19)[default__.safeIndex(597, ((self).f19).length(0))], self.f25)
            nw54_[int(14)] = d_303_v42_
            nw54_[int(15)] = d_303_v42_
            nw54_[int(16)] = _dafny.Map({(self).f20: self.f25})
            nw54_[int(17)] = d_303_v42_
            nw54_[int(18)] = ((d_304_v43_)[(self).f20] if ((self).f20) in (d_304_v43_) else d_303_v42_)
            nw54_[int(19)] = _dafny.Map({((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]: self.f25})
            nw54_[int(20)] = d_303_v42_
            nw54_[int(21)] = d_303_v42_
            nw54_[int(22)] = d_303_v42_
            nw54_[int(23)] = d_303_v42_
            nw54_[int(24)] = d_303_v42_
            nw54_[int(25)] = d_303_v42_
            nw54_[int(26)] = _dafny.Map({(d_235_v0_)[default__.safeIndex((self.f25).f24, len(d_235_v0_))]: self.f25})
            nw54_[int(27)] = _dafny.Map({(self).f20: self.f25})
            d_305_v44_ = nw54_
            d_306_v45_: _dafny.Map
            d_306_v45_ = _dafny.Map({((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]: default__.fm0(self.f26, ((self).f19)[default__.safeIndex(597, ((self).f19).length(0))], globalState)})
            d_307_v46_: _dafny.Seq
            d_307_v46_ = _dafny.SeqWithoutIsStrInference([d_306_v45_])
            d_308_v47_: D3
            d_308_v47_ = D3_DC10(d_305_v44_, (d_307_v46_)[default__.safeIndex(-519, len(d_307_v46_))], _dafny.Set({d_237_v2_}), (932) - ((self.f25).f24))
            d_309_v48_: str
            d_309_v48_ = _dafny.CodePoint('l')
            d_310_v49_: _dafny.Seq
            d_310_v49_ = _dafny.SeqWithoutIsStrInference([(d_309_v48_ if self.f26 else d_309_v48_), d_309_v48_])
            d_311_v50_: _dafny.MultiSet
            d_311_v50_ = _dafny.MultiSet([((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]])
            d_312_v51_: _dafny.Set
            d_312_v51_ = _dafny.Set({(d_311_v50_).intersection(d_311_v50_)})
            d_313_v52_: _dafny.MultiSet
            d_313_v52_ = _dafny.MultiSet([d_237_v2_, d_237_v2_])
            d_314_v53_: _dafny.MultiSet
            d_314_v53_ = _dafny.MultiSet([(0) - ((self.f25).f24), (self.f25).f24])
            d_315_v54_: _dafny.Array
            nw55_ = _dafny.Array(None, 24)
            nw55_[int(0)] = ((d_313_v52_)[d_237_v2_] if (d_237_v2_) in (d_313_v52_) else (self.f25).f24)
            nw55_[int(1)] = (0) - ((self.f25).f24)
            nw55_[int(2)] = default__.safeModuloInt((self.f25).f24, (self.f25).f24)
            nw55_[int(3)] = (0) - (default__.safeModuloInt((0) - ((self.f25).f24), (0) - ((self.f25).f24)))
            nw55_[int(4)] = (self.f25).f24
            nw55_[int(5)] = (self.f25).f24
            nw55_[int(6)] = 40
            nw55_[int(7)] = 670
            nw55_[int(8)] = default__.safeDivisionInt((0) - ((self.f25).f24), (self.f25).f24)
            nw55_[int(9)] = len(d_310_v49_)
            nw55_[int(10)] = 337
            nw55_[int(11)] = (self.f25).f24
            nw55_[int(12)] = ((self.f25).f24) + ((self.f25).f24)
            nw55_[int(13)] = (self.f25).f24
            nw55_[int(14)] = len(d_310_v49_)
            nw55_[int(15)] = (self.f25).f24
            nw55_[int(16)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxgapglw"))) if ((self).f19)[default__.safeIndex(597, ((self).f19).length(0))] else (self.f25).f24)
            nw55_[int(17)] = (self.f25).f24
            nw55_[int(18)] = (611 if self.f26 else (0) - (default__.fm0(self.f26, self.f26, globalState)))
            nw55_[int(19)] = ((self.f25).f24) * ((self.f25).f24)
            nw55_[int(20)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))
            nw55_[int(21)] = (283) - ((self.f25).f24)
            nw55_[int(22)] = (self.f25).f24
            nw55_[int(23)] = ((d_311_v50_).cardinality) + ((d_314_v53_).cardinality)
            d_315_v54_ = nw55_
            index45_ = default__.safeIndex(929, (d_315_v54_).length(0))
            (d_315_v54_)[index45_] = (self.f25).f24
            d_316_v55_: _dafny.Seq
            d_316_v55_ = _dafny.SeqWithoutIsStrInference([d_305_v44_])
            d_317_v56_: D0
            d_317_v56_ = D0_DC2((self).f20, 682, d_235_v0_, len(d_310_v49_), (self.f25).f24)
            d_318_v57_: _dafny.Map
            d_318_v57_ = _dafny.Map({(d_317_v56_).cf7: ((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]})
            d_319_v58_: _dafny.Map
            d_319_v58_ = _dafny.Map({len(d_318_v57_): _dafny.SeqWithoutIsStrInference([d_309_v48_, d_309_v48_, d_309_v48_])})
            pat_let_tv9_ = d_306_v45_
            pat_let_tv10_ = d_316_v55_
            pat_let_tv11_ = d_316_v55_
            index46_ = default__.safeIndex(929, (d_315_v54_).length(0))
            def iife18_(_pat_let4_0):
                def iife19_(d_320_dt__update__tmp_h1_):
                    def iife20_(_pat_let5_0):
                        def iife21_(d_321_dt__update_hcf18_h0_):
                            def iife22_(_pat_let6_0):
                                def iife23_(d_322_dt__update_hcf17_h0_):
                                    return D3_DC10(d_322_dt__update_hcf17_h0_, d_321_dt__update_hcf18_h0_, (d_320_dt__update__tmp_h1_).cf19, (d_320_dt__update__tmp_h1_).cf20)
                                return iife23_(_pat_let6_0)
                            return iife22_((pat_let_tv10_)[default__.safeIndex((self.f25).f24, len(pat_let_tv11_))])
                        return iife21_(_pat_let5_0)
                    return iife20_(pat_let_tv9_)
                return iife19_(_pat_let4_0)
            rhs21_ = iife18_(d_308_v47_)
            rhs22_ = ((d_319_v58_)[len(d_235_v0_)] if (len(d_235_v0_)) in (d_319_v58_) else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e')])) + (_dafny.SeqWithoutIsStrInference([d_309_v48_, d_309_v48_])))
            rhs23_ = d_312_v51_
            rhs24_ = (0) - ((self.f25).f24)
            lhs20_ = d_315_v54_
            lhs21_ = default__.safeIndex(929, (d_315_v54_).length(0))
            d_308_v47_ = rhs21_
            d_310_v49_ = rhs22_
            d_312_v51_ = rhs23_
            lhs20_[lhs21_] = rhs24_
            d_323_v59_: D3
            d_323_v59_ = D3_DC11((self.f25).f24, (self.f25).f24, not (((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]) or (((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]))
            source6_ = d_323_v59_
            if source6_.is_DC10:
                d_324___mcc_h11_ = source6_.cf17
                d_325___mcc_h12_ = source6_.cf18
                d_326___mcc_h13_ = source6_.cf19
                d_327___mcc_h14_ = source6_.cf20
                d_328_cf20_ = d_327___mcc_h14_
                d_329_cf19_ = d_326___mcc_h13_
                d_330_cf18_ = d_325___mcc_h12_
                d_331_cf17_ = d_324___mcc_h11_
                d_332_v60_: D4
                d_332_v60_ = D4_DC12(d_309_v48_)
                pat_let_tv12_ = d_309_v48_
                d_333_v61_: _dafny.Map
                def iife24_(_pat_let7_0):
                    def iife25_(d_334_dt__update__tmp_h2_):
                        def iife26_(_pat_let8_0):
                            def iife27_(d_335_dt__update_hcf24_h0_):
                                return D4_DC12(d_335_dt__update_hcf24_h0_)
                            return iife27_(_pat_let8_0)
                        return iife26_(pat_let_tv12_)
                    return iife25_(_pat_let7_0)
                d_333_v61_ = _dafny.Map({iife24_(d_332_v60_): (self.f25).f24})
                d_333_v61_ = (d_333_v61_).set(d_332_v60_, default__.safeModuloInt((self.f25).f24, (self.f25).f24))
                (globalState).f0 = not(((d_318_v57_)[(0) - ((len(d_235_v0_)) + ((_dafny.MultiSet([(self.f25).f24])).cardinality))] if ((0) - ((len(d_235_v0_)) + ((_dafny.MultiSet([(self.f25).f24])).cardinality))) in (d_318_v57_) else (self).f20))
                (self).f26 = ((self).f19)[default__.safeIndex(597, ((self).f19).length(0))]
                d_336_v62_: _dafny.Set
                d_336_v62_ = _dafny.Set({self.f26, ((d_311_v50_).set(self.f26, default__.abs((self.f25).f24))) != (d_311_v50_), (self).f20, ((self.f25).f24) >= ((self.f25).f24)})
                d_337_v63_: D4
                d_337_v63_ = D4_DC14(((self).f19)[default__.safeIndex(597, ((self).f19).length(0))], d_336_v62_)
                pat_let_tv13_ = d_336_v62_
                def iife28_(_pat_let9_0):
                    def iife29_(d_338_dt__update__tmp_h3_):
                        def iife30_(_pat_let10_0):
                            def iife31_(d_339_dt__update_hcf26_h0_):
                                return D4_DC14((d_338_dt__update__tmp_h3_).cf25, d_339_dt__update_hcf26_h0_)
                            return iife31_(_pat_let10_0)
                        return iife30_(pat_let_tv13_)
                    return iife29_(_pat_let9_0)
                d_336_v62_ = (iife28_(d_337_v63_)).cf26
            elif source6_.is_DC11:
                d_340___mcc_h15_ = source6_.cf21
                d_341___mcc_h16_ = source6_.cf22
                d_342___mcc_h17_ = source6_.cf23
                d_343_cf23_ = d_342___mcc_h17_
                d_344_cf22_ = d_341___mcc_h16_
                d_345_cf21_ = d_340___mcc_h15_
                d_346_v64_: _dafny.Array
                def lambda12_(d_347_cf22_):
                    def lambda13_(d_348_i2_):
                        return ((self.f25).f24) > (d_347_cf22_)

                    return lambda13_

                init7_ = lambda12_(d_344_cf22_)
                nw56_ = _dafny.Array(None, 14)
                for i0_7_ in range(nw56_.length(0)):
                    nw56_[i0_7_] = init7_(i0_7_)
                d_346_v64_ = nw56_
                d_346_v64_ = d_346_v64_
                (globalState).f18 = (d_315_v54_)[default__.safeIndex(929, (d_315_v54_).length(0))]
                d_349_v65_: C1
                nw57_ = C1()
                nw57_.ctor__(d_310_v49_, d_346_v64_, self.f26)
                d_349_v65_ = nw57_
                d_350_v66_: _dafny.Map
                d_350_v66_ = _dafny.Map({(d_349_v65_).f27: self.f26})
                d_350_v66_ = (d_350_v66_).set((d_349_v65_).f27, not((self).f20))
            elif True:
                d_351___mcc_h18_ = source6_.cf16
                d_352_cf16_ = d_351___mcc_h18_
                d_353_v68_: _dafny.Array
                def lambda14_(d_354_v48_):
                    def lambda15_(d_355_i3_):
                        def iife32_():
                            coll10_ = _dafny.Set()
                            compr_10_: str
                            for compr_10_ in (_dafny.Set({d_354_v48_, d_354_v48_})).Elements:
                                d_356_v67_: str = compr_10_
                                if (d_356_v67_) in (_dafny.Set({d_354_v48_, d_354_v48_})):
                                    coll10_ = coll10_.union(_dafny.Set([d_356_v67_]))
                            return _dafny.Set(coll10_)
                        return (_dafny.Set({d_354_v48_})).isdisjoint(iife32_()
                        )

                    return lambda15_

                init8_ = lambda14_(d_309_v48_)
                nw58_ = _dafny.Array(None, 8)
                for i0_8_ in range(nw58_.length(0)):
                    nw58_[i0_8_] = init8_(i0_8_)
                d_353_v68_ = nw58_
                d_353_v68_ = d_353_v68_
                (globalState).f10 = d_315_v54_
                d_357_v69_: C0
                nw59_ = C0()
                nw59_.ctor__(((self.f25).f24) - ((self.f25).f24))
                d_357_v69_ = nw59_
                d_358_v70_: _dafny.Map
                d_358_v70_ = _dafny.Map({(self).f20: self.f26})
                d_358_v70_ = d_358_v70_
            d_359_v71_: bool
            d_360_v72_: int
            d_361_v73_: int
            d_362_v74_: int
            out16_: bool
            out17_: int
            out18_: int
            out19_: int
            out16_, out17_, out18_, out19_ = (self).m4(not((self).f20), globalState)
            d_359_v71_ = out16_
            d_360_v72_ = out17_
            d_361_v73_ = out18_
            d_362_v74_ = out19_
        d_363_v75_: _dafny.Seq
        d_363_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eneiyaye"))
        d_363_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbiwky"))
        if self.f26:
            d_364_v76_: _dafny.Map
            d_364_v76_ = _dafny.Map({955: (self.f25).f24})
            d_365_v77_: _dafny.MultiSet
            d_365_v77_ = _dafny.MultiSet([(self.f25).f24, ((d_364_v76_)[(self.f25).f24] if ((self.f25).f24) in (d_364_v76_) else len(d_363_v75_)), 403, (self.f25).f24])
            d_366_v78_: _dafny.Map
            d_366_v78_ = _dafny.Map({(self).f20: (d_365_v77_).set((_dafny.MultiSet(d_235_v0_)).cardinality, default__.abs(len(d_235_v0_)))})
            d_365_v77_ = (((d_366_v78_)[self.f26] if (self.f26) in (d_366_v78_) else d_365_v77_)) - (d_365_v77_)
            d_367_v79_: C1
            nw60_ = C1()
            nw60_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_368_i4_ in range(default__.abs(276))]), (self).f19, self.f26)
            d_367_v79_ = nw60_
            d_369_v80_: _dafny.Seq
            d_369_v80_ = _dafny.SeqWithoutIsStrInference([d_367_v79_])
            d_370_v81_: _dafny.Map
            d_370_v81_ = _dafny.Map({not(self.f26): (self).f20})
            d_371_v82_: D0
            d_371_v82_ = D0_DC2(default__.fm9((self).f20, globalState), len(d_370_v81_), _dafny.SeqWithoutIsStrInference([(self).f20, self.f26]), (self.f25).f24, (self.f25).f24)
            d_372_v83_: _dafny.Array
            nw61_ = _dafny.Array(int(0), 7)
            d_372_v83_ = nw61_
            pat_let_tv14_ = d_235_v0_
            d_373_v84_: _dafny.Map
            def iife33_(_pat_let11_0):
                def iife34_(d_374_dt__update__tmp_h4_):
                    def iife35_(_pat_let12_0):
                        def iife36_(d_375_dt__update_hcf5_h0_):
                            def iife37_(_pat_let13_0):
                                def iife38_(d_376_dt__update_hcf4_h0_):
                                    def iife39_(_pat_let14_0):
                                        def iife40_(d_377_dt__update_hcf6_h0_):
                                            return D0_DC2(d_376_dt__update_hcf4_h0_, d_375_dt__update_hcf5_h0_, d_377_dt__update_hcf6_h0_, (d_374_dt__update__tmp_h4_).cf7, (d_374_dt__update__tmp_h4_).cf8)
                                        return iife40_(_pat_let14_0)
                                    return iife39_(pat_let_tv14_)
                                return iife38_(_pat_let13_0)
                            return iife37_(True)
                        return iife36_(_pat_let12_0)
                    return iife35_((self.f25).f24)
                return iife34_(_pat_let11_0)
            d_373_v84_ = _dafny.Map({(len(d_369_v80_)) - ((iife33_(d_371_v82_)).cf7): d_372_v83_})
            d_378_v85_: D3
            d_378_v85_ = D3_DC9(d_363_v75_)
            d_379_v86_: _dafny.MultiSet
            d_379_v86_ = _dafny.MultiSet([D3_DC9((d_367_v79_).f27), d_378_v85_, D3_DC9(d_363_v75_), d_378_v85_, d_378_v85_])
            d_373_v84_ = (d_373_v84_).set((d_379_v86_).cardinality, d_372_v83_)
            d_380_v87_: _dafny.MultiSet
            d_380_v87_ = _dafny.MultiSet([(self).f20])
            d_380_v87_ = d_380_v87_
            if not((self.f26) or ((self.f26) or (not((self).f20)))):
                (globalState).f0 = (self).f20
                d_381_v88_: _dafny.Set
                d_381_v88_ = _dafny.Set({(0) - ((self.f25).f24)})
                rhs25_ = (d_381_v88_).ispropersubset((d_381_v88_) - (d_381_v88_))
                rhs26_ = (self.f25).f24
                lhs22_ = self
                lhs23_ = globalState
                lhs22_.f26 = rhs25_
                lhs23_.f13 = rhs26_
                d_382_v89_: D2
                d_382_v89_ = D2_DC8((self.f25).f24, default__.fm0(self.f26, self.f26, globalState))
                rhs27_ = (d_382_v89_).cf15
                rhs28_ = (d_372_v83_ if (self).f20 else d_372_v83_)
                r3 = rhs27_
                d_372_v83_ = rhs28_
                (globalState).f18 = (self.f25).f24
                d_383_v90_: _dafny.Map
                d_383_v90_ = _dafny.Map({(self.f25).f24: (self).f20})
                (globalState).f18 = default__.safeDivisionInt((self.f25).f24, (default__.fm12(d_365_v77_, len((d_383_v90_).set(len((d_367_v79_).f27), (self).f20)), globalState)).cardinality)
            elif True:
                index47_ = default__.safeIndex(291, ((self).f19).length(0))
                ((self).f19)[index47_] = self.f26
                d_384_v91_: D1
                d_384_v91_ = D1_DC5(True)
                d_385_v92_: str
                d_385_v92_ = _dafny.CodePoint('o')
                d_386_v93_: T0
                nw62_ = C1()
                nw62_.ctor__(((d_367_v79_).f27).set(default__.safeIndex((0) - ((self.f25).f24), len((d_367_v79_).f27)), d_385_v92_), (self).f19, self.f26)
                d_386_v93_ = nw62_
                d_387_v94_: _dafny.Seq
                d_387_v94_ = _dafny.SeqWithoutIsStrInference([d_386_v93_, d_386_v93_])
                d_388_v95_: _dafny.MultiSet
                d_388_v95_ = _dafny.MultiSet([d_386_v93_, (d_387_v94_)[default__.safeIndex((self.f25).f24, len(d_387_v94_))], d_386_v93_])
                index48_ = default__.safeIndex(291, ((self).f19).length(0))
                rhs29_ = ((d_388_v95_).cardinality) != (((self.f25).f24) * ((self.f25).f24))
                rhs30_ = (self.f25).f24
                rhs31_ = (D1_DC5((d_386_v93_).f20) if self.f26 else d_384_v91_)
                lhs24_ = (self).f19
                lhs25_ = default__.safeIndex(291, ((self).f19).length(0))
                lhs24_[lhs25_] = rhs29_
                r3 = rhs30_
                d_384_v91_ = rhs31_
                r3 = (default__.safeModuloInt(len(d_363_v75_), (0) - ((self.f25).f24))) - ((self.f25).f24)
                d_389_v96_: _dafny.Set
                d_389_v96_ = _dafny.Set({(self).f20})
                d_390_v97_: C0
                nw63_ = C0()
                nw63_.ctor__(default__.safeDivisionInt(len(d_389_v96_), (self.f25).f24))
                d_390_v97_ = nw63_
                d_391_v98_: _dafny.Array
                nw64_ = _dafny.Array(D1.default()(), 26)
                d_391_v98_ = nw64_
                index49_ = default__.safeIndex(378, (d_391_v98_).length(0))
                (d_391_v98_)[index49_] = D1_DC4(d_236_v1_)
                index50_ = default__.safeIndex(378, (d_391_v98_).length(0))
                (d_391_v98_)[index50_] = d_237_v2_
                (globalState).f10 = d_372_v83_
            index51_ = default__.safeIndex(534, ((self).f19).length(0))
            ((self).f19)[index51_] = (self).f20
            index52_ = default__.safeIndex(534, ((self).f19).length(0))
            ((self).f19)[index52_] = (self).f20
        elif True:
            index53_ = default__.safeIndex(329, ((self).f19).length(0))
            ((self).f19)[index53_] = True
            index54_ = default__.safeIndex(329, ((self).f19).length(0))
            ((self).f19)[index54_] = self.f26
            d_392_v99_: _dafny.Array
            nw65_ = _dafny.Array(None, 10)
            nw65_[int(0)] = (0) - ((self.f25).f24)
            nw65_[int(1)] = (self.f25).f24
            nw65_[int(2)] = (self.f25).f24
            nw65_[int(3)] = 226
            nw65_[int(4)] = (_dafny.MultiSet([(self.f25).f24])).cardinality
            nw65_[int(5)] = (self.f25).f24
            nw65_[int(6)] = 67
            nw65_[int(7)] = 418
            nw65_[int(8)] = (self.f25).f24
            nw65_[int(9)] = (self.f25).f24
            d_392_v99_ = nw65_
            (globalState).f10 = (D5_DC16(d_392_v99_)).cf28
            d_393_v100_: _dafny.Array
            nw66_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_393_v100_ = nw66_
            d_394_v101_: _dafny.Array
            def lambda16_(d_395_i5_):
                return D0_DC3(_dafny.Map({_dafny.MultiSet([((self).f19)[default__.safeIndex(329, ((self).f19).length(0))], (self).f20, (self).f20, ((self).f19)[default__.safeIndex(329, ((self).f19).length(0))]]): 118}))

            init9_ = lambda16_
            nw67_ = _dafny.Array(None, 16)
            for i0_9_ in range(nw67_.length(0)):
                nw67_[i0_9_] = init9_(i0_9_)
            d_394_v101_ = nw67_
            index55_ = default__.safeIndex(174, (d_393_v100_).length(0))
            (d_393_v100_)[index55_] = d_394_v101_
            index56_ = default__.safeIndex(174, (d_393_v100_).length(0))
            (d_393_v100_)[index56_] = d_394_v101_
            d_396_v102_: _dafny.Array
            def lambda17_(d_397_i6_):
                return (_dafny.Map({((self).f19)[default__.safeIndex(329, ((self).f19).length(0))]: (self).f20})) | (_dafny.Map({(self).f20: ((self).f19)[default__.safeIndex(329, ((self).f19).length(0))]}))

            init10_ = lambda17_
            nw68_ = _dafny.Array(None, 11)
            for i0_10_ in range(nw68_.length(0)):
                nw68_[i0_10_] = init10_(i0_10_)
            d_396_v102_ = nw68_
            d_396_v102_ = d_396_v102_
            d_398_v103_: _dafny.Array
            nw69_ = _dafny.Array(_dafny.Map({}), 25)
            d_398_v103_ = nw69_
            d_399_v104_: _dafny.Set
            d_399_v104_ = _dafny.Set({d_237_v2_})
            d_400_v105_: D3
            d_400_v105_ = D3_DC10(d_398_v103_, default__.fm16(self.f26, self.f26, len(d_363_v75_), globalState), d_399_v104_, (0) - ((self.f25).f24))
            r2 = (0) - ((d_400_v105_).cf20)
        d_401_v106_: _dafny.Array
        nw70_ = _dafny.Array(int(0), 8)
        d_401_v106_ = nw70_
        (globalState).f3 = d_401_v106_
        d_402_v107_: _dafny.Array
        nw71_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_402_v107_ = nw71_
        d_403_v108_: _dafny.Array
        nw72_ = _dafny.Array(_dafny.CodePoint('D'), 24)
        d_403_v108_ = nw72_
        index57_ = default__.safeIndex(856, (d_402_v107_).length(0))
        (d_402_v107_)[index57_] = d_403_v108_
        index58_ = default__.safeIndex(856, (d_402_v107_).length(0))
        (d_402_v107_)[index58_] = d_403_v108_
        r3 = ((self.f25).f24) - (((self.f25).f24) + ((self.f25).f24))
        r0 = self.f26
        r1 = ((self).f20 if False else self.f26)
        r2 = len((d_363_v75_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_404_i7_ in range(default__.abs(910))])))
        r3 = len((d_363_v75_) + ((d_363_v75_) + (d_363_v75_)))
        return r0, r1, r2, r3


class C3:
    def  __init__(self):
        self.f22: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f23: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f22, f23):
        (self).f22 = f22
        (self)._f23 = f23

    def fm5(self, p0, globalState):
        source7_ = D0_DC3(_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True])): len(_dafny.SeqWithoutIsStrInference([True, False, True]))}))
        if source7_.is_DC1:
            d_405___mcc_h0_ = source7_.cf1
            d_406___mcc_h1_ = source7_.cf2
            d_407___mcc_h2_ = source7_.cf3
            d_408_cf3_ = d_407___mcc_h2_
            d_409_cf2_ = d_406___mcc_h1_
            d_410_cf1_ = d_405___mcc_h0_
            return 543
        elif source7_.is_DC2:
            d_411___mcc_h3_ = source7_.cf4
            d_412___mcc_h4_ = source7_.cf5
            d_413___mcc_h5_ = source7_.cf6
            d_414___mcc_h6_ = source7_.cf7
            d_415___mcc_h7_ = source7_.cf8
            d_416_cf8_ = d_415___mcc_h7_
            d_417_cf7_ = d_414___mcc_h6_
            d_418_cf6_ = d_413___mcc_h5_
            d_419_cf5_ = d_412___mcc_h4_
            d_420_cf4_ = d_411___mcc_h3_
            return (len(_dafny.SeqWithoutIsStrInference([d_420_cf4_]))) - (d_419_cf5_)
        elif source7_.is_DC3:
            d_421___mcc_h8_ = source7_.cf9
            d_422_cf9_ = d_421___mcc_h8_
            return 4
        elif True:
            d_423___mcc_h9_ = source7_.cf0
            d_424_cf0_ = d_423___mcc_h9_
            return default__.safeModuloInt(206, -935)

    def fm6(self, p0, p1, p2, p3, globalState):
        return len((_dafny.SeqWithoutIsStrInference([D0_DC2(not(True), len(_dafny.Map({283: not(True)})), _dafny.SeqWithoutIsStrInference([True]), 78, 231) for d_425_i0_ in range(default__.abs(13))])) + ((_dafny.SeqWithoutIsStrInference([D0_DC2(not(True), 872, _dafny.SeqWithoutIsStrInference([False, True]), 945, len(_dafny.Set({True})))])) + ((D1_DC4(_dafny.SeqWithoutIsStrInference([D0_DC2(False, (_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([not(False), False]))}))])).cardinality, _dafny.SeqWithoutIsStrInference([False]), -264, len(_dafny.SeqWithoutIsStrInference([428]))) for d_426_i1_ in range(default__.abs(462))]))).cf10)))

    def m2(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_427_v0_: int
        d_427_v0_ = 419
        d_428_v1_: bool
        d_428_v1_ = True
        d_429_v2_: _dafny.Set
        d_429_v2_ = _dafny.Set({False, d_428_v1_, d_428_v1_, d_428_v1_})
        d_430_v3_: _dafny.Map
        d_430_v3_ = _dafny.Map({-146: _dafny.Set({d_428_v1_})})
        d_431_v4_: _dafny.Seq
        d_431_v4_ = _dafny.SeqWithoutIsStrInference([len(d_429_v2_), len(((d_430_v3_)[len(_dafny.Map({len(d_429_v2_): self.f22}))] if (len(_dafny.Map({len(d_429_v2_): self.f22}))) in (d_430_v3_) else d_429_v2_)), d_427_v0_, d_427_v0_])
        d_432_v5_: _dafny.Array
        nw73_ = _dafny.Array(None, 4)
        nw73_[int(0)] = d_427_v0_
        nw73_[int(1)] = (581) + (d_427_v0_)
        nw73_[int(2)] = (d_431_v4_)[default__.safeIndex((0) - (d_427_v0_), len(d_431_v4_))]
        nw73_[int(3)] = d_427_v0_
        d_432_v5_ = nw73_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_432_v5_).length(0)):
            d_433_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_433_i0_)) and ((d_433_i0_) < ((d_432_v5_).length(0)))):
                (d_432_v5_)[(d_433_i0_)] = default__.safeDivisionInt(d_433_i0_, d_427_v0_)
        d_434_v6_: C0
        nw74_ = C0()
        nw74_.ctor__(default__.safeModuloInt(d_427_v0_, d_427_v0_))
        d_434_v6_ = nw74_
        d_435_v7_: _dafny.MultiSet
        d_435_v7_ = _dafny.MultiSet([d_428_v1_])
        d_436_v8_: D0
        d_436_v8_ = D0_DC0(d_435_v7_)
        d_436_v8_ = D0_DC0(d_435_v7_)
        d_437_v9_: D2
        d_437_v9_ = D2_DC7(_dafny.SeqWithoutIsStrInference([len(self.f22) for d_438_i1_ in range(default__.abs(22))]))
        d_431_v4_ = (d_437_v9_).cf13
        (globalState).f18 = (d_427_v0_) - (d_427_v0_)
        r1 = (len(d_429_v2_)) < (len(_dafny.SeqWithoutIsStrInference([False, d_428_v1_, True, d_428_v1_, d_428_v1_])))
        d_439_v10_: _dafny.Seq
        d_439_v10_ = _dafny.SeqWithoutIsStrInference([default__.fm7(d_427_v0_, (0) - (d_427_v0_), d_427_v0_, globalState), (d_431_v4_) < (_dafny.SeqWithoutIsStrInference([d_427_v0_ for d_440_i2_ in range(default__.abs(595))]))])
        r0 = d_439_v10_
        r1 = d_428_v1_
        return r0, r1

    def m3(self, globalState):
        d_441_v0_: bool
        d_441_v0_ = True
        d_442_v1_: _dafny.Seq
        d_442_v1_ = _dafny.SeqWithoutIsStrInference([d_441_v0_, False, d_441_v0_, d_441_v0_])
        (globalState).f13 = len((d_442_v1_ if d_441_v0_ else d_442_v1_))
        d_443_v2_: _dafny.Array
        nw75_ = _dafny.Array(_dafny.CodePoint('D'), 4)
        d_443_v2_ = nw75_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_443_v2_).length(0)):
            d_444_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_444_i0_)) and ((d_444_i0_) < ((d_443_v2_).length(0)))):
                (d_443_v2_)[(d_444_i0_)] = _dafny.CodePoint('d')
        d_445_v3_: T0
        nw76_ = C1()
        nw76_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxxd")), (self).f23, d_441_v0_)
        d_445_v3_ = nw76_
        d_446_v4_: _dafny.Map
        d_446_v4_ = _dafny.Map({d_445_v3_: (self).f23})
        d_446_v4_ = d_446_v4_
        d_447_v5_: _dafny.Array
        d_447_v5_ = (self).f23
        d_448_v6_: _dafny.Set
        d_448_v6_ = _dafny.Set({(self).f23, (self).f23})
        d_449_v7_: int
        d_449_v7_ = -719
        d_450_v8_: _dafny.Map
        d_450_v8_ = _dafny.Map({(d_448_v6_).ispropersubset(_dafny.Set({(self).f23, (d_447_v5_), (self).f23, (self).f23, (d_445_v3_).f19})): d_449_v7_})
        d_451_v9_: _dafny.Seq
        d_451_v9_ = _dafny.SeqWithoutIsStrInference([d_450_v8_])
        d_452_v10_: _dafny.Map
        d_452_v10_ = _dafny.Map({d_450_v8_: (d_451_v9_)[default__.safeIndex(d_449_v7_, len(d_451_v9_))]})
        d_453_v11_: _dafny.Map
        d_453_v11_ = _dafny.Map({False: d_450_v8_})
        d_454_v12_: _dafny.Array
        nw77_ = _dafny.Array(_dafny.Map({}), 8)
        d_454_v12_ = nw77_
        d_455_v13_: _dafny.Set
        d_455_v13_ = _dafny.Set({D1_DC4(_dafny.SeqWithoutIsStrInference([D0_DC2(not((d_445_v3_).f20), d_449_v7_, _dafny.SeqWithoutIsStrInference([(d_442_v1_)[default__.safeIndex(d_449_v7_, len(d_442_v1_))]]), d_449_v7_, d_449_v7_) for d_456_i1_ in range(default__.abs(724))]))})
        d_457_v14_: D3
        d_457_v14_ = D3_DC10(d_454_v12_, _dafny.Map({default__.fm7(d_449_v7_, 844, d_449_v7_, globalState): d_449_v7_}), d_455_v13_, d_449_v7_)
        d_450_v8_ = ((((d_452_v10_)[_dafny.Map({(d_445_v3_).f20: d_449_v7_})] if (_dafny.Map({(d_445_v3_).f20: d_449_v7_})) in (d_452_v10_) else default__.fm16((d_445_v3_).f20, not(d_441_v0_), d_449_v7_, globalState)) if (d_445_v3_).f20 else ((d_453_v11_)[(d_445_v3_).f20] if ((d_445_v3_).f20) in (d_453_v11_) else d_450_v8_))) | ((d_457_v14_).cf18)
        if (d_445_v3_).f20:
            (globalState).f18 = d_449_v7_
            (globalState).f0 = d_441_v0_
            rhs32_ = (-257) - (d_449_v7_)
            rhs33_ = d_449_v7_
            lhs26_ = globalState
            d_449_v7_ = rhs32_
            lhs26_.f18 = rhs33_
            d_458_v15_: D1
            d_458_v15_ = D1_DC5((d_445_v3_).f20)
            rhs34_ = d_441_v0_
            rhs35_ = d_458_v15_
            rhs36_ = not (False) or ((d_445_v3_).f20)
            lhs27_ = globalState
            lhs28_ = globalState
            lhs27_.f0 = rhs34_
            d_458_v15_ = rhs35_
            lhs28_.f0 = rhs36_
            d_459_v16_: _dafny.Seq
            d_459_v16_ = _dafny.SeqWithoutIsStrInference([(d_445_v3_).f19, (d_445_v3_).f19, (d_445_v3_).f19])
            d_460_v17_: str
            d_460_v17_ = _dafny.CodePoint('o')
            d_461_v18_: _dafny.Map
            d_461_v18_ = _dafny.Map({d_460_v17_: d_449_v7_})
            d_462_v19_: _dafny.Seq
            d_462_v19_ = _dafny.SeqWithoutIsStrInference([d_461_v18_])
            d_463_v20_: D0
            d_463_v20_ = D0_DC1((d_445_v3_).f20, 514, d_449_v7_)
            d_464_v21_: D5
            d_464_v21_ = D5_DC17(len(d_442_v1_), d_449_v7_, d_463_v20_, (d_445_v3_).f20, (d_445_v3_).f20)
            d_465_v22_: _dafny.Array
            def lambda18_(d_466_i2_):
                return (d_466_i2_) * (len(self.f22))

            init11_ = lambda18_
            nw78_ = _dafny.Array(None, 13)
            for i0_11_ in range(nw78_.length(0)):
                nw78_[i0_11_] = init11_(i0_11_)
            d_465_v22_ = nw78_
            d_467_v23_: _dafny.MultiSet
            d_467_v23_ = _dafny.MultiSet([d_465_v22_, d_465_v22_])
            d_468_v24_: _dafny.MultiSet
            d_468_v24_ = _dafny.MultiSet([(d_467_v23_).cardinality, d_449_v7_])
            d_469_v25_: _dafny.Seq
            d_469_v25_ = _dafny.SeqWithoutIsStrInference([d_449_v7_])
            d_470_v26_: _dafny.Seq
            d_470_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([((d_468_v24_)[len(d_450_v8_)] if (len(d_450_v8_)) in (d_468_v24_) else d_449_v7_), d_449_v7_, len(d_469_v25_)]), _dafny.SeqWithoutIsStrInference([len(d_450_v8_) for d_471_i3_ in range(default__.abs(453))])])
            pat_let_tv15_ = d_449_v7_
            pat_let_tv16_ = d_442_v1_
            d_472_v27_: _dafny.Seq
            def iife41_(_pat_let15_0):
                def iife42_(d_473_dt__update__tmp_h0_):
                    def iife43_(_pat_let16_0):
                        def iife44_(d_474_dt__update_hcf8_h0_):
                            def iife45_(_pat_let17_0):
                                def iife46_(d_475_dt__update_hcf6_h0_):
                                    return D0_DC2((d_473_dt__update__tmp_h0_).cf4, (d_473_dt__update__tmp_h0_).cf5, d_475_dt__update_hcf6_h0_, (d_473_dt__update__tmp_h0_).cf7, d_474_dt__update_hcf8_h0_)
                                return iife46_(_pat_let17_0)
                            return iife45_(pat_let_tv16_)
                        return iife44_(_pat_let16_0)
                    return iife43_(pat_let_tv15_)
                return iife42_(_pat_let15_0)
            d_472_v27_ = _dafny.SeqWithoutIsStrInference([iife41_(D0_DC2((d_445_v3_).f20, d_449_v7_, d_442_v1_, d_449_v7_, len(self.f22)))])
            d_476_v28_: D1
            d_476_v28_ = D1_DC4(d_472_v27_)
            d_477_v29_: _dafny.Seq
            d_477_v29_ = _dafny.SeqWithoutIsStrInference([self.f22])
            pat_let_tv17_ = d_445_v3_
            pat_let_tv18_ = d_449_v7_
            pat_let_tv19_ = d_442_v1_
            pat_let_tv20_ = d_449_v7_
            pat_let_tv21_ = d_477_v29_
            pat_let_tv22_ = d_449_v7_
            pat_let_tv23_ = d_477_v29_
            d_478_v30_: _dafny.MultiSet
            def iife47_(_pat_let18_0):
                def iife48_(d_479_dt__update__tmp_h1_):
                    def iife49_(_pat_let19_0):
                        def iife50_(d_480_dt__update_hcf10_h0_):
                            return D1_DC4(d_480_dt__update_hcf10_h0_)
                        return iife50_(_pat_let19_0)
                    return iife49_(_dafny.SeqWithoutIsStrInference([D0_DC2(not((pat_let_tv17_).f20), pat_let_tv18_, pat_let_tv19_, pat_let_tv20_, len((pat_let_tv21_)[default__.safeIndex(pat_let_tv22_, len(pat_let_tv23_))]))]))
                return iife48_(_pat_let18_0)
            d_478_v30_ = _dafny.MultiSet([iife47_(d_476_v28_), d_476_v28_])
            d_481_v31_: _dafny.Array
            nw79_ = _dafny.Array(None, 27)
            nw79_[int(0)] = len(d_459_v16_)
            nw79_[int(1)] = len(self.f22)
            nw79_[int(2)] = (d_449_v7_ if d_441_v0_ else 851)
            nw79_[int(3)] = (default__.fm0(False, (d_445_v3_).f20, globalState)) * (d_449_v7_)
            nw79_[int(4)] = (0) - (d_449_v7_)
            nw79_[int(5)] = (d_449_v7_) + (d_449_v7_)
            nw79_[int(6)] = (0) - (d_449_v7_)
            nw79_[int(7)] = 398
            nw79_[int(8)] = len((_dafny.Set({(d_445_v3_).f20})) - (default__.fm17(len(self.f22), globalState)))
            nw79_[int(9)] = d_449_v7_
            nw79_[int(10)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnxjphmtm")))
            nw79_[int(11)] = len(d_462_v19_)
            nw79_[int(12)] = d_449_v7_
            nw79_[int(13)] = default__.fm0((d_445_v3_).f20, (d_464_v21_).cf33, globalState)
            nw79_[int(14)] = default__.safeModuloInt(d_449_v7_, d_449_v7_)
            nw79_[int(15)] = (0) - (((0) - (d_449_v7_)) * (d_449_v7_))
            nw79_[int(16)] = d_449_v7_
            nw79_[int(17)] = d_449_v7_
            nw79_[int(18)] = len((d_470_v26_)[default__.safeIndex(d_449_v7_, len(d_470_v26_))])
            nw79_[int(19)] = d_449_v7_
            nw79_[int(20)] = d_449_v7_
            nw79_[int(21)] = d_449_v7_
            nw79_[int(22)] = d_449_v7_
            nw79_[int(23)] = d_449_v7_
            nw79_[int(24)] = (d_449_v7_) * ((d_478_v30_).cardinality)
            nw79_[int(25)] = len((self.f22) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))))
            nw79_[int(26)] = d_449_v7_
            d_481_v31_ = nw79_
            index59_ = default__.safeIndex(209, (d_465_v22_).length(0))
            (d_465_v22_)[index59_] = d_449_v7_
            d_482_v32_: _dafny.Array
            nw80_ = _dafny.Array(_dafny.Seq({}), 13)
            d_482_v32_ = nw80_
            d_483_v33_: D5
            d_483_v33_ = D5_DC19((d_468_v24_).cardinality, (d_445_v3_).f20, d_482_v32_, (d_445_v3_).f20, d_460_v17_)
            index60_ = default__.safeIndex(209, (d_465_v22_).length(0))
            rhs37_ = (d_449_v7_) - (d_449_v7_)
            rhs38_ = (d_442_v1_)[default__.safeIndex(d_449_v7_, len(d_442_v1_))]
            rhs39_ = (d_483_v33_).cf36
            rhs40_ = d_449_v7_
            lhs29_ = d_465_v22_
            lhs30_ = default__.safeIndex(209, (d_465_v22_).length(0))
            lhs31_ = globalState
            lhs32_ = globalState
            lhs33_ = globalState
            lhs29_[lhs30_] = rhs37_
            lhs31_.f0 = rhs38_
            lhs32_.f0 = rhs39_
            lhs33_.f13 = rhs40_
        elif True:
            (globalState).f0 = ((_dafny.MultiSet([(d_445_v3_).f20])).set((d_445_v3_).f20, default__.abs(d_449_v7_))).isdisjoint(_dafny.MultiSet([d_441_v0_, not((d_445_v3_).f20), d_441_v0_]))
            if d_441_v0_:
                d_484_v34_: C0
                nw81_ = C0()
                nw81_.ctor__(d_449_v7_)
                d_484_v34_ = nw81_
                d_485_v35_: str
                d_485_v35_ = _dafny.CodePoint('p')
                d_486_v36_: _dafny.Map
                d_486_v36_ = _dafny.Map({d_485_v35_: (d_445_v3_).f19})
                d_487_v37_: C2
                nw82_ = C2()
                nw82_.ctor__(d_484_v34_, (d_445_v3_).f20, ((d_445_v3_).f19 if not((d_445_v3_).f20) else ((d_486_v36_)[d_485_v35_] if (d_485_v35_) in (d_486_v36_) else (d_445_v3_).f19)), default__.fm9((d_445_v3_).f20, globalState))
                d_487_v37_ = nw82_
                d_488_v38_: _dafny.Array
                nw83_ = _dafny.Array(int(0), 23)
                d_488_v38_ = nw83_
                d_489_v39_: _dafny.MultiSet
                d_489_v39_ = _dafny.MultiSet([False, (d_445_v3_).f20])
                index61_ = default__.safeIndex(621, (d_488_v38_).length(0))
                (d_488_v38_)[index61_] = len(_dafny.SeqWithoutIsStrInference([d_489_v39_]))
                index62_ = default__.safeIndex(621, (d_488_v38_).length(0))
                (d_488_v38_)[index62_] = d_449_v7_
                d_449_v7_ = default__.safeModuloInt((d_449_v7_) + ((d_484_v34_).f24), (d_488_v38_)[default__.safeIndex(621, (d_488_v38_).length(0))])
                index63_ = default__.safeIndex(300, ((d_445_v3_).f19).length(0))
                ((d_445_v3_).f19)[index63_] = not(((d_445_v3_).f20) or ((d_445_v3_).f20))
                index64_ = default__.safeIndex(300, ((d_445_v3_).f19).length(0))
                ((d_445_v3_).f19)[index64_] = False
                d_490_v40_: _dafny.Seq
                d_490_v40_ = _dafny.SeqWithoutIsStrInference([((d_488_v38_)[default__.safeIndex(621, (d_488_v38_).length(0))]) * ((d_488_v38_)[default__.safeIndex(621, (d_488_v38_).length(0))]), d_449_v7_])
                (globalState).f13 = (d_490_v40_)[default__.safeIndex((d_484_v34_).f24, len(d_490_v40_))]
            elif True:
                d_491_v41_: D3
                d_491_v41_ = D3_DC9(self.f22)
                d_492_v42_: _dafny.Map
                d_492_v42_ = _dafny.Map({d_441_v0_: _dafny.MultiSet([d_491_v41_, d_491_v41_])})
                d_493_v43_: _dafny.MultiSet
                d_493_v43_ = _dafny.MultiSet([d_491_v41_])
                d_494_v44_: _dafny.Seq
                d_494_v44_ = _dafny.SeqWithoutIsStrInference([d_449_v7_, 448, d_449_v7_, d_449_v7_])
                (globalState).f0 = (default__.fm18(d_494_v44_, globalState)).issubset((((d_492_v42_)[(d_445_v3_).f20] if ((d_445_v3_).f20) in (d_492_v42_) else d_493_v43_)) | (d_493_v43_))
                d_495_v45_: _dafny.Map
                d_495_v45_ = _dafny.Map({d_441_v0_: (d_445_v3_).f20})
                d_495_v45_ = (d_495_v45_).set((d_445_v3_).f20, (d_445_v3_).f20)
                d_496_v46_: _dafny.Map
                d_496_v46_ = _dafny.Map({d_449_v7_: not((d_445_v3_).f20)})
                d_497_v47_: C0
                nw84_ = C0()
                nw84_.ctor__(len(d_496_v46_))
                d_497_v47_ = nw84_
                d_498_v48_: C2
                nw85_ = C2()
                nw85_.ctor__(d_497_v47_, (d_445_v3_).f20, (d_445_v3_).f19, (self.f22) < (self.f22))
                d_498_v48_ = nw85_
                d_499_v49_: _dafny.Array
                nw86_ = _dafny.Array(D0.default()(), 1)
                d_499_v49_ = nw86_
                d_500_v50_: _dafny.Array
                d_500_v50_ = d_499_v49_
                d_501_v51_: _dafny.Array
                nw87_ = _dafny.Array(None, 6)
                nw87_[int(0)] = d_499_v49_
                nw87_[int(1)] = d_499_v49_
                nw87_[int(2)] = d_499_v49_
                nw87_[int(3)] = d_499_v49_
                nw87_[int(4)] = d_499_v49_
                nw87_[int(5)] = (d_500_v50_)
                d_501_v51_ = nw87_
                index65_ = default__.safeIndex(56, (d_501_v51_).length(0))
                (d_501_v51_)[index65_] = d_499_v49_
                index66_ = default__.safeIndex(56, (d_501_v51_).length(0))
                (d_501_v51_)[index66_] = d_499_v49_
                d_502_v52_: _dafny.MultiSet
                d_502_v52_ = _dafny.MultiSet([not(d_441_v0_), (d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20])
                (globalState).f0 = (((d_502_v52_)[not(d_441_v0_)] if (not(d_441_v0_)) in (d_502_v52_) else len(default__.fm19(True, globalState)))) < ((d_497_v47_).f24)
            d_503_v53_: _dafny.Seq
            d_504_v54_: bool
            out20_: _dafny.Seq
            out21_: bool
            out20_, out21_ = (self).m2(globalState)
            d_503_v53_ = out20_
            d_504_v54_ = out21_
            d_505_v55_: str
            d_505_v55_ = _dafny.CodePoint('o')
            d_505_v55_ = d_505_v55_
            (globalState).f13 = ((0) - (d_449_v7_)) - ((d_449_v7_) + (d_449_v7_))
        if not(d_441_v0_):
            d_449_v7_ = -455
            d_506_v56_: _dafny.Set
            d_506_v56_ = _dafny.Set({d_449_v7_})
            rhs41_ = d_449_v7_
            rhs42_ = (d_445_v3_).f20
            rhs43_ = (d_445_v3_).f20
            rhs44_ = ((d_506_v56_ if (d_445_v3_).f20 else d_506_v56_)).issubset(d_506_v56_)
            lhs34_ = globalState
            lhs35_ = globalState
            lhs36_ = globalState
            lhs37_ = globalState
            lhs34_.f13 = rhs41_
            lhs35_.f0 = rhs42_
            lhs36_.f0 = rhs43_
            lhs37_.f0 = rhs44_
            (globalState).f13 = (len(d_450_v8_) if (d_445_v3_).f20 else d_449_v7_)
            d_507_v57_: str
            d_507_v57_ = _dafny.CodePoint('p')
            (globalState).f13 = len(((self.f22) + ((self.f22).set(default__.safeIndex(d_449_v7_, len(self.f22)), d_507_v57_))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usturyl"))))
            if (d_445_v3_).f20:
                d_508_v58_: D4
                d_508_v58_ = D4_DC13()
                d_508_v58_ = d_508_v58_
                d_509_v59_: _dafny.Map
                d_509_v59_ = _dafny.Map({(d_445_v3_).f20: d_507_v57_})
                d_510_v60_: _dafny.Map
                d_510_v60_ = _dafny.Map({(d_445_v3_).f20: (d_441_v0_) in (d_509_v59_)})
                d_510_v60_ = ((d_510_v60_) | (d_510_v60_)) | (d_510_v60_)
                d_511_v61_: _dafny.Seq
                d_511_v61_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))])
                d_512_v62_: _dafny.Array
                nw88_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_512_v62_ = nw88_
                d_513_v63_: _dafny.Map
                d_513_v63_ = _dafny.Map({len(d_511_v61_): d_512_v62_})
                d_514_v64_: _dafny.Map
                d_514_v64_ = _dafny.Map({d_449_v7_: ((d_513_v63_)[680] if (680) in (d_513_v63_) else d_512_v62_)})
                d_514_v64_ = (d_514_v64_).set((d_449_v7_ if (d_445_v3_).f20 else d_449_v7_), d_512_v62_)
                (globalState).f0 = (not((d_445_v3_).f20) if default__.fm7(len(default__.fm19(True, globalState)), d_449_v7_, d_449_v7_, globalState) else (d_441_v0_) in (d_442_v1_))
                d_515_v65_: D0
                d_515_v65_ = D0_DC1((d_445_v3_).f20, d_449_v7_, d_449_v7_)
                d_516_v66_: _dafny.Array
                nw89_ = _dafny.Array(None, 9)
                nw89_[int(0)] = d_449_v7_
                nw89_[int(1)] = default__.fm0((d_515_v65_).cf1, (d_445_v3_).f20, globalState)
                nw89_[int(2)] = d_449_v7_
                nw89_[int(3)] = d_449_v7_
                nw89_[int(4)] = len(d_442_v1_)
                nw89_[int(5)] = 139
                nw89_[int(6)] = d_449_v7_
                nw89_[int(7)] = d_449_v7_
                nw89_[int(8)] = d_449_v7_
                d_516_v66_ = nw89_
                d_517_v67_: D5
                d_517_v67_ = D5_DC16(d_516_v66_)
                d_518_v68_: _dafny.Map
                d_518_v68_ = _dafny.Map({d_517_v67_: (d_445_v3_).f19})
                d_519_v69_: C1
                nw90_ = C1()
                nw90_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjlkirkt")), ((d_518_v68_)[d_517_v67_] if (d_517_v67_) in (d_518_v68_) else (d_445_v3_).f19), (d_442_v1_)[default__.safeIndex(d_449_v7_, len(d_442_v1_))])
                d_519_v69_ = nw90_
            elif True:
                (globalState).f18 = ((d_450_v8_)[(d_445_v3_).f20] if ((d_445_v3_).f20) in (d_450_v8_) else len(_dafny.SeqWithoutIsStrInference([(d_445_v3_).f20])))
                d_449_v7_ = len(self.f22)
                d_520_v70_: D0
                d_520_v70_ = D0_DC1((d_445_v3_).f20, 310, (0) - (d_449_v7_))
                d_521_v71_: _dafny.MultiSet
                d_521_v71_ = _dafny.MultiSet([D0_DC1((d_445_v3_).f20, d_449_v7_, d_449_v7_), d_520_v70_, d_520_v70_, d_520_v70_])
                d_522_v72_: C1
                nw91_ = C1()
                nw91_.ctor__(_dafny.SeqWithoutIsStrInference([d_507_v57_ for d_523_i4_ in range(default__.abs(38))]), (self).f23, (d_445_v3_).f20)
                d_522_v72_ = nw91_
                d_524_v73_: _dafny.Map
                d_524_v73_ = _dafny.Map({d_449_v7_: d_522_v72_})
                d_525_v74_: _dafny.Map
                d_525_v74_ = _dafny.Map({d_449_v7_: len(d_524_v73_)})
                d_526_v75_: _dafny.Seq
                d_526_v75_ = _dafny.SeqWithoutIsStrInference([((d_521_v71_)[D0_DC1((d_445_v3_).f20, ((d_525_v74_)[d_449_v7_] if (d_449_v7_) in (d_525_v74_) else d_449_v7_), d_449_v7_)] if (D0_DC1((d_445_v3_).f20, ((d_525_v74_)[d_449_v7_] if (d_449_v7_) in (d_525_v74_) else d_449_v7_), d_449_v7_)) in (d_521_v71_) else d_449_v7_), d_449_v7_, (_dafny.MultiSet([d_449_v7_])).cardinality])
                d_527_v76_: C0
                nw92_ = C0()
                nw92_.ctor__(d_449_v7_)
                d_527_v76_ = nw92_
                d_528_v77_: _dafny.Seq
                d_528_v77_ = _dafny.SeqWithoutIsStrInference([d_527_v76_])
                d_529_v78_: D5
                d_529_v78_ = D5_DC18(d_528_v77_)
                d_530_v79_: _dafny.Map
                d_530_v79_ = _dafny.Map({d_529_v78_: (d_445_v3_).f20})
                d_531_v80_: _dafny.Array
                nw93_ = _dafny.Array(None, 16)
                nw93_[int(0)] = d_449_v7_
                nw93_[int(1)] = d_449_v7_
                nw93_[int(2)] = d_449_v7_
                nw93_[int(3)] = d_449_v7_
                nw93_[int(4)] = d_449_v7_
                nw93_[int(5)] = (d_526_v75_)[default__.safeIndex(d_449_v7_, len(d_526_v75_))]
                nw93_[int(6)] = d_449_v7_
                nw93_[int(7)] = (0) - (default__.safeDivisionInt(d_449_v7_, d_449_v7_))
                nw93_[int(8)] = (d_449_v7_) - (d_449_v7_)
                nw93_[int(9)] = len(d_530_v79_)
                nw93_[int(10)] = (self).fm5((d_445_v3_).f20, globalState)
                nw93_[int(11)] = (len(d_442_v1_)) + ((d_527_v76_).f24)
                nw93_[int(12)] = d_449_v7_
                nw93_[int(13)] = len((d_442_v1_ if (d_445_v3_).f20 else d_442_v1_))
                nw93_[int(14)] = len(((self.f22) + (self.f22)).set(default__.safeIndex((d_527_v76_).f24, len((self.f22) + (self.f22))), d_507_v57_))
                nw93_[int(15)] = (d_449_v7_) * ((0) - ((d_527_v76_).f24))
                d_531_v80_ = nw93_
                index67_ = default__.safeIndex(546, (d_531_v80_).length(0))
                (d_531_v80_)[index67_] = (d_527_v76_).f24
                index68_ = default__.safeIndex(825, (d_531_v80_).length(0))
                (d_531_v80_)[index68_] = default__.safeDivisionInt((d_527_v76_).f24, len(_dafny.SeqWithoutIsStrInference([(d_445_v3_).f20, d_441_v0_])))
                index69_ = default__.safeIndex(546, (d_531_v80_).length(0))
                index70_ = default__.safeIndex(825, (d_531_v80_).length(0))
                rhs45_ = (d_449_v7_) - ((d_527_v76_).f24)
                rhs46_ = d_449_v7_
                rhs47_ = d_449_v7_
                lhs38_ = globalState
                lhs39_ = d_531_v80_
                lhs40_ = default__.safeIndex(546, (d_531_v80_).length(0))
                lhs41_ = d_531_v80_
                lhs42_ = default__.safeIndex(825, (d_531_v80_).length(0))
                lhs38_.f13 = rhs45_
                lhs39_[lhs40_] = rhs46_
                lhs41_[lhs42_] = rhs47_
                d_532_v81_: D4
                d_532_v81_ = D4_DC14((d_445_v3_).f20, _dafny.Set({(d_445_v3_).f20}))
                (globalState).f0 = ((d_532_v81_).cf25 if (d_445_v3_).f20 else (d_445_v3_).f20)
                (globalState).f13 = d_449_v7_
        elif True:
            d_533_v82_: _dafny.Set
            d_533_v82_ = _dafny.Set({(d_445_v3_).f20})
            d_534_v83_: _dafny.MultiSet
            d_534_v83_ = _dafny.MultiSet([d_449_v7_, d_449_v7_])
            d_535_v84_: _dafny.Seq
            d_535_v84_ = _dafny.SeqWithoutIsStrInference([(d_534_v83_).cardinality])
            d_536_v85_: _dafny.Map
            d_536_v85_ = _dafny.Map({d_533_v82_: (d_535_v84_)[default__.safeIndex(d_449_v7_, len(d_535_v84_))]})
            d_537_v86_: _dafny.Map
            d_537_v86_ = _dafny.Map({d_449_v7_: d_449_v7_})
            d_538_v87_: _dafny.Map
            d_538_v87_ = _dafny.Map({d_537_v86_: True})
            if (((d_536_v85_)[d_533_v82_] if (d_533_v82_) in (d_536_v85_) else d_449_v7_)) < ((0) - (len(d_538_v87_))):
                d_539_v88_: _dafny.Set
                d_539_v88_ = _dafny.Set({default__.safeModuloInt(d_449_v7_, d_449_v7_), d_449_v7_, default__.fm0((d_445_v3_).f20, (d_445_v3_).f20, globalState), d_449_v7_})
                def iife51_():
                    coll11_ = _dafny.Set()
                    compr_11_: int
                    for compr_11_ in (_dafny.MultiSet([len(self.f22)])).Elements:
                        d_540_v89_: int = compr_11_
                        if (d_540_v89_) in (_dafny.MultiSet([len(self.f22)])):
                            coll11_ = coll11_.union(_dafny.Set([default__.safeModuloInt(d_540_v89_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False])]))).cardinality)]))
                    return _dafny.Set(coll11_)
                d_539_v88_ = (iife51_()
                ).intersection(d_539_v88_)
                d_541_v90_: str
                d_541_v90_ = _dafny.CodePoint('p')
                d_542_v91_: _dafny.Map
                d_542_v91_ = _dafny.Map({(d_445_v3_).f20: d_541_v90_})
                d_542_v91_ = (d_542_v91_).set(d_441_v0_, d_541_v90_)
                d_543_v92_: _dafny.Map
                d_543_v92_ = _dafny.Map({d_449_v7_: (self).f23})
                rhs48_ = d_543_v92_
                rhs49_ = not(not((d_445_v3_).f20))
                d_543_v92_ = rhs48_
                d_441_v0_ = rhs49_
                d_544_v93_: _dafny.Set
                d_544_v93_ = _dafny.Set({d_533_v82_})
                d_545_v94_: _dafny.MultiSet
                d_545_v94_ = _dafny.MultiSet([d_541_v90_])
                d_546_v95_: _dafny.Map
                d_546_v95_ = _dafny.Map({d_449_v7_: d_545_v94_})
                d_442_v1_ = _dafny.SeqWithoutIsStrInference([(d_544_v93_).ispropersubset(d_544_v93_), (d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20, (((d_546_v95_)[len(d_535_v84_)] if (len(d_535_v84_)) in (d_546_v95_) else d_545_v94_)) != (d_545_v94_)])
                d_547_v96_: _dafny.Array
                def lambda19_(d_548_v1_, d_549_v3_):
                    def lambda20_(d_550_i5_):
                        return _dafny.SeqWithoutIsStrInference([D0_DC0(_dafny.MultiSet((d_548_v1_).set(default__.safeIndex(len(d_548_v1_), len(d_548_v1_)), (d_549_v3_).f20))), D0_DC0(_dafny.MultiSet([(d_549_v3_).f20, (d_549_v3_).f20, (d_549_v3_).f20, not((d_549_v3_).f20), (d_549_v3_).f20])), D0_DC0(_dafny.MultiSet(d_548_v1_))])

                    return lambda20_

                init12_ = lambda19_(d_442_v1_, d_445_v3_)
                nw94_ = _dafny.Array(None, 23)
                for i0_12_ in range(nw94_.length(0)):
                    nw94_[i0_12_] = init12_(i0_12_)
                d_547_v96_ = nw94_
                d_551_v97_: _dafny.MultiSet
                d_551_v97_ = _dafny.MultiSet([d_441_v0_])
                d_552_v98_: D0
                d_552_v98_ = D0_DC0(d_551_v97_)
                d_553_v99_: _dafny.Seq
                d_553_v99_ = _dafny.SeqWithoutIsStrInference([D0_DC0(d_551_v97_), d_552_v98_])
                index71_ = default__.safeIndex(420, (d_547_v96_).length(0))
                (d_547_v96_)[index71_] = (d_553_v99_ if (d_445_v3_).f20 else d_553_v99_)
                index72_ = default__.safeIndex(420, (d_547_v96_).length(0))
                (d_547_v96_)[index72_] = ((d_553_v99_).set(default__.safeIndex((0) - (d_449_v7_), len(d_553_v99_)), D0_DC0(default__.fm20((d_445_v3_).f20, d_449_v7_, globalState)))).set(default__.safeIndex(d_449_v7_, len((d_553_v99_).set(default__.safeIndex((0) - (d_449_v7_), len(d_553_v99_)), D0_DC0(default__.fm20((d_445_v3_).f20, d_449_v7_, globalState))))), (d_552_v98_ if (d_445_v3_).f20 else default__.fm21(d_449_v7_, globalState)))
            elif True:
                d_554_v100_: _dafny.MultiSet
                d_554_v100_ = _dafny.MultiSet([self.f22, self.f22])
                d_555_v102_: C1
                nw95_ = C1()
                nw95_.ctor__(self.f22, (d_445_v3_).f19, d_441_v0_)
                d_555_v102_ = nw95_
                d_556_v103_: _dafny.MultiSet
                d_556_v103_ = _dafny.MultiSet([d_555_v102_])
                d_557_v104_: _dafny.Map
                d_557_v104_ = _dafny.Map({default__.fm22(((d_556_v103_)[d_555_v102_] if (d_555_v102_) in (d_556_v103_) else d_449_v7_), d_449_v7_, globalState): len((d_555_v102_).f27)})
                d_558_v105_: _dafny.MultiSet
                d_558_v105_ = _dafny.MultiSet([d_441_v0_, (d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20])
                d_559_v106_: str
                d_559_v106_ = _dafny.CodePoint('d')
                d_560_v107_: _dafny.Map
                d_560_v107_ = _dafny.Map({d_559_v106_: d_441_v0_})
                d_561_v108_: _dafny.Set
                d_561_v108_ = _dafny.Set({len(d_560_v107_)})
                d_562_v109_: _dafny.Map
                d_562_v109_ = _dafny.Map({d_558_v105_: len(d_561_v108_)})
                d_563_v110_: _dafny.Array
                nw96_ = _dafny.Array(None, 28)
                nw96_[int(0)] = len(self.f22)
                nw96_[int(1)] = d_449_v7_
                nw96_[int(2)] = (d_449_v7_ if (d_445_v3_).f20 else (0) - (d_449_v7_))
                nw96_[int(3)] = default__.safeModuloInt((0) - (d_449_v7_), 645)
                nw96_[int(4)] = (0) - (d_449_v7_)
                nw96_[int(5)] = d_449_v7_
                nw96_[int(6)] = -369
                nw96_[int(7)] = len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brlgftm")): False}))
                nw96_[int(8)] = d_449_v7_
                nw96_[int(9)] = 122
                nw96_[int(10)] = d_449_v7_
                nw96_[int(11)] = (d_449_v7_) * ((0) - (d_449_v7_))
                nw96_[int(12)] = default__.safeDivisionInt((0) - (d_449_v7_), d_449_v7_)
                nw96_[int(13)] = d_449_v7_
                def iife52_():
                    coll12_ = _dafny.Set()
                    compr_12_: _dafny.Seq
                    for compr_12_ in (d_554_v100_).Elements:
                        d_564_v101_: _dafny.Seq = compr_12_
                        if (d_564_v101_) in (d_554_v100_):
                            coll12_ = coll12_.union(_dafny.Set([d_564_v101_]))
                    return _dafny.Set(coll12_)
                nw96_[int(14)] = (d_449_v7_) * (len(iife52_()
                ))
                nw96_[int(15)] = (0) - ((0) - ((d_449_v7_ if (d_445_v3_).f20 else d_449_v7_)))
                nw96_[int(16)] = default__.safeModuloInt(len(self.f22), len(d_453_v11_))
                nw96_[int(17)] = d_449_v7_
                nw96_[int(18)] = (d_554_v100_).cardinality
                nw96_[int(19)] = d_449_v7_
                nw96_[int(20)] = (942) - ((0) - (d_449_v7_))
                nw96_[int(21)] = d_449_v7_
                nw96_[int(22)] = default__.safeModuloInt(406, d_449_v7_)
                nw96_[int(23)] = d_449_v7_
                nw96_[int(24)] = ((d_557_v104_)[d_562_v109_] if (d_562_v109_) in (d_557_v104_) else d_449_v7_)
                nw96_[int(25)] = (d_449_v7_) - (d_449_v7_)
                nw96_[int(26)] = d_449_v7_
                nw96_[int(27)] = (d_449_v7_) * (d_449_v7_)
                d_563_v110_ = nw96_
                index73_ = default__.safeIndex(266, (d_563_v110_).length(0))
                (d_563_v110_)[index73_] = d_449_v7_
                index74_ = default__.safeIndex(266, (d_563_v110_).length(0))
                (d_563_v110_)[index74_] = d_449_v7_
                d_565_v111_: C1
                nw97_ = C1()
                nw97_.ctor__((d_555_v102_).f27, (self).f23, (d_445_v3_).f20)
                d_565_v111_ = nw97_
                d_566_v112_: _dafny.Map
                d_566_v112_ = _dafny.Map({((d_563_v110_)[default__.safeIndex(266, (d_563_v110_).length(0))]) + ((d_563_v110_)[default__.safeIndex(266, (d_563_v110_).length(0))]): False})
                d_566_v112_ = (d_566_v112_).set((d_449_v7_) * ((d_563_v110_)[default__.safeIndex(266, (d_563_v110_).length(0))]), (d_445_v3_).f20)
                d_533_v82_ = ((d_533_v82_).intersection(d_533_v82_)).intersection(d_533_v82_)
                d_450_v8_ = (d_450_v8_).set((d_445_v3_).f20, (d_563_v110_)[default__.safeIndex(266, (d_563_v110_).length(0))])
            d_567_v113_: D4
            d_567_v113_ = D4_DC14((d_445_v3_).f20, d_533_v82_)
            d_568_v114_: _dafny.Seq
            d_568_v114_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_441_v0_, (d_567_v113_).cf25})])
            d_569_v115_: str
            d_569_v115_ = _dafny.CodePoint('y')
            d_570_v116_: _dafny.Map
            d_570_v116_ = _dafny.Map({d_569_v115_: (d_445_v3_).f20})
            d_571_v117_: _dafny.Map
            d_571_v117_ = _dafny.Map({(d_445_v3_).f20: (d_445_v3_).f20})
            d_572_v118_: _dafny.Map
            d_572_v118_ = _dafny.Map({917: (d_445_v3_).f20})
            d_573_v119_: _dafny.Map
            d_573_v119_ = _dafny.Map({d_449_v7_: _dafny.Set({not(((d_572_v118_)[-341] if (-341) in (d_572_v118_) else (d_442_v1_)[default__.safeIndex(d_449_v7_, len(d_442_v1_))]))})})
            d_574_v120_: _dafny.Array
            nw98_ = _dafny.Array(None, 19)
            nw98_[int(0)] = (d_568_v114_)[default__.safeIndex((0) - (d_449_v7_), len(d_568_v114_))]
            nw98_[int(1)] = _dafny.Set({(d_445_v3_).f20, ((d_570_v116_)[_dafny.CodePoint('w')] if (_dafny.CodePoint('w')) in (d_570_v116_) else (d_445_v3_).f20)})
            nw98_[int(2)] = d_533_v82_
            nw98_[int(3)] = d_533_v82_
            nw98_[int(4)] = d_533_v82_
            nw98_[int(5)] = d_533_v82_
            nw98_[int(6)] = d_533_v82_
            nw98_[int(7)] = default__.fm17(len(d_571_v117_), globalState)
            nw98_[int(8)] = (d_533_v82_) - (d_533_v82_)
            nw98_[int(9)] = _dafny.Set({(d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20})
            nw98_[int(10)] = (d_533_v82_) | (((d_573_v119_)[len(self.f22)] if (len(self.f22)) in (d_573_v119_) else d_533_v82_))
            nw98_[int(11)] = d_533_v82_
            nw98_[int(12)] = d_533_v82_
            nw98_[int(13)] = (default__.fm17(d_449_v7_, globalState)) - (_dafny.Set({(d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20, (d_445_v3_).f20}))
            nw98_[int(14)] = d_533_v82_
            nw98_[int(15)] = (d_568_v114_)[default__.safeIndex(d_449_v7_, len(d_568_v114_))]
            nw98_[int(16)] = d_533_v82_
            nw98_[int(17)] = (d_533_v82_) - (d_533_v82_)
            nw98_[int(18)] = (d_533_v82_) - (d_533_v82_)
            d_574_v120_ = nw98_
            index75_ = default__.safeIndex(274, (d_574_v120_).length(0))
            (d_574_v120_)[index75_] = d_533_v82_
            index76_ = default__.safeIndex(274, (d_574_v120_).length(0))
            (d_574_v120_)[index76_] = ((d_533_v82_) | (d_533_v82_)).intersection(_dafny.Set({(d_445_v3_).f20, d_441_v0_}))
            d_575_v121_: _dafny.Array
            def lambda21_(d_576_v7_, d_577_v118_, d_578_v3_, d_579_v84_, d_580_v117_):
                def lambda22_(d_581_i6_):
                    return _dafny.SeqWithoutIsStrInference([D3_DC11(d_576_v7_, len(d_577_v118_), (d_578_v3_).f20), D3_DC11(len(d_579_v84_), len(d_580_v117_), (d_578_v3_).f20), D3_DC11(len(_dafny.SeqWithoutIsStrInference([len(d_580_v117_) for d_582_i7_ in range(default__.abs(-641))])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_583_i8_ in range(default__.abs(830))])), (d_578_v3_).f20), D3_DC11(d_576_v7_, d_576_v7_, True), D3_DC11(d_576_v7_, d_576_v7_, (d_578_v3_).f20)])

                return lambda22_

            init13_ = lambda21_(d_449_v7_, d_572_v118_, d_445_v3_, d_535_v84_, d_571_v117_)
            nw99_ = _dafny.Array(None, 24)
            for i0_13_ in range(nw99_.length(0)):
                nw99_[i0_13_] = init13_(i0_13_)
            d_575_v121_ = nw99_
            d_584_v122_: D5
            d_584_v122_ = D5_DC19(d_449_v7_, (d_445_v3_).f20, d_575_v121_, d_441_v0_, d_569_v115_)
            d_571_v117_ = (d_571_v117_).set((d_584_v122_).cf38, not(True))
            rhs50_ = d_449_v7_
            rhs51_ = (d_442_v1_)[default__.safeIndex(len(self.f22), len(d_442_v1_))]
            lhs43_ = globalState
            lhs44_ = globalState
            lhs43_.f18 = rhs50_
            lhs44_.f0 = rhs51_
            d_449_v7_ = d_449_v7_

    @property
    def f23(self):
        return self._f23

class C4(T0):
    def  __init__(self):
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f20: bool = False
        self._f21: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f21, f19, f20):
        (self)._f21 = f21
        (self)._f19 = f19
        (self)._f20 = f20

    def fm1(self, p0, p1, p2, globalState):
        if (self).f20:
            return _dafny.MultiSet([(self).f20])
        elif True:
            return ((D0_DC0(_dafny.MultiSet([(self).f20]))).cf0).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f20])))

    def m0(self, p0, p1, globalState):
        d_585_v0_: _dafny.Seq
        d_585_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrrji"))
        (globalState).f13 = len((d_585_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqvxm"))))
        d_585_v0_ = d_585_v0_
        (globalState).f0 = not((self).f20)
        d_586_v1_: _dafny.MultiSet
        d_586_v1_ = _dafny.MultiSet([(self).f20, p0])
        d_587_v2_: D0
        d_587_v2_ = D0_DC0(d_586_v1_)
        pat_let_tv24_ = p0
        pat_let_tv25_ = p0
        pat_let_tv26_ = p0
        def lambda23_(source8_):
            if source8_.is_DC1:
                d_588___mcc_h0_ = source8_.cf1
                d_589___mcc_h1_ = source8_.cf2
                d_590___mcc_h2_ = source8_.cf3
                d_591_cf3_ = d_590___mcc_h2_
                d_592_cf2_ = d_589___mcc_h1_
                d_593_cf1_ = d_588___mcc_h0_
                return pat_let_tv24_
            elif source8_.is_DC2:
                d_594___mcc_h3_ = source8_.cf4
                d_595___mcc_h4_ = source8_.cf5
                d_596___mcc_h5_ = source8_.cf6
                d_597___mcc_h6_ = source8_.cf7
                d_598___mcc_h7_ = source8_.cf8
                d_599_cf8_ = d_598___mcc_h7_
                d_600_cf7_ = d_597___mcc_h6_
                d_601_cf6_ = d_596___mcc_h5_
                d_602_cf5_ = d_595___mcc_h4_
                d_603_cf4_ = d_594___mcc_h3_
                return d_603_cf4_
            elif source8_.is_DC3:
                d_604___mcc_h8_ = source8_.cf9
                d_605_cf9_ = d_604___mcc_h8_
                return pat_let_tv25_
            elif True:
                d_606___mcc_h9_ = source8_.cf0
                d_607_cf0_ = d_606___mcc_h9_
                return pat_let_tv26_

        if lambda23_(d_587_v2_):
            (globalState).f0 = p0
            (globalState).f0 = (d_585_v0_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqfpsses")))
            d_608_v3_: _dafny.Map
            d_608_v3_ = _dafny.Map({(self).f21: (self).f21})
            d_609_v4_: _dafny.Map
            d_609_v4_ = _dafny.Map({(d_585_v0_) == (d_585_v0_): d_608_v3_})
            d_609_v4_ = d_609_v4_
            if p0:
                (globalState).f13 = (self).f21
                d_610_v5_: _dafny.Seq
                d_610_v5_ = _dafny.SeqWithoutIsStrInference([p0])
                d_611_v6_: D0
                d_611_v6_ = D0_DC2((self).f20, default__.fm0((self).f20, True, globalState), d_610_v5_, (self).f21, (self).f21)
                (globalState).f13 = (d_611_v6_).cf5
                (globalState).f0 = not((True) and (((self).f21) > ((self).f21)))
                (globalState).f0 = p0
                d_612_v7_: _dafny.Map
                d_612_v7_ = _dafny.Map({(0) - (len(d_585_v0_)): _dafny.MultiSet([(self).f20, p0, p0, p0])})
                d_613_v8_: _dafny.Seq
                d_613_v8_ = _dafny.SeqWithoutIsStrInference([d_586_v1_, d_586_v1_, d_586_v1_, d_586_v1_, d_586_v1_])
                d_614_v9_: _dafny.Seq
                d_614_v9_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                d_615_v10_: _dafny.Array
                nw100_ = _dafny.Array(None, 26)
                nw100_[int(0)] = _dafny.MultiSet([p0, (self).f20, p0])
                nw100_[int(1)] = (((d_612_v7_)[(self).f21] if ((self).f21) in (d_612_v7_) else d_586_v1_)).set(not((self).f20), default__.abs(138))
                nw100_[int(2)] = (d_586_v1_) - (d_586_v1_)
                nw100_[int(3)] = (_dafny.MultiSet(d_610_v5_)).set(p0, default__.abs((self).f21))
                nw100_[int(4)] = d_586_v1_
                nw100_[int(5)] = d_586_v1_
                nw100_[int(6)] = _dafny.MultiSet(d_610_v5_)
                nw100_[int(7)] = d_586_v1_
                nw100_[int(8)] = d_586_v1_
                nw100_[int(9)] = d_586_v1_
                nw100_[int(10)] = ((_dafny.MultiSet([p0])).set(p0, default__.abs((self).f21))) - (d_586_v1_)
                nw100_[int(11)] = d_586_v1_
                nw100_[int(12)] = (d_586_v1_ if (self).f20 else d_586_v1_)
                nw100_[int(13)] = (d_613_v8_)[default__.safeIndex(len(d_614_v9_), len(d_613_v8_))]
                nw100_[int(14)] = (_dafny.MultiSet([default__.fm2((self).f20, (self).f21, (self).f21, globalState)])) | (d_586_v1_)
                nw100_[int(15)] = (self).fm1((self).f21, (self).f21, (self).f20, globalState)
                nw100_[int(16)] = ((self).fm1((self).f21, len(p1), True, globalState)).intersection(_dafny.MultiSet([(self).f20, False]))
                nw100_[int(17)] = (d_587_v2_).cf0
                nw100_[int(18)] = _dafny.MultiSet(d_610_v5_)
                nw100_[int(19)] = (self).fm1((self).f21, (self).f21, p0, globalState)
                nw100_[int(20)] = (self).fm1((self).f21, (self).f21, (self).f20, globalState)
                nw100_[int(21)] = _dafny.MultiSet([(self).f20, (self).f20, not(True), False, p0])
                nw100_[int(22)] = (_dafny.MultiSet([False, (self).f20])) | (d_586_v1_)
                nw100_[int(23)] = (d_613_v8_)[default__.safeIndex((d_586_v1_).cardinality, len(d_613_v8_))]
                nw100_[int(24)] = d_586_v1_
                nw100_[int(25)] = (d_587_v2_).cf0
                d_615_v10_ = nw100_
                index77_ = default__.safeIndex(580, (d_615_v10_).length(0))
                (d_615_v10_)[index77_] = (self).fm1((0) - ((self).f21), (self).f21, default__.fm2(p0, (self).f21, (self).f21, globalState), globalState)
                index78_ = default__.safeIndex(580, (d_615_v10_).length(0))
                (d_615_v10_)[index78_] = d_586_v1_
            elif True:
                d_616_v11_: _dafny.Map
                d_616_v11_ = _dafny.Map({p0: p0})
                index79_ = default__.safeIndex(883, ((self).f19).length(0))
                ((self).f19)[index79_] = p0
                index80_ = default__.safeIndex(883, ((self).f19).length(0))
                rhs52_ = d_616_v11_
                rhs53_ = (d_585_v0_) == ((d_585_v0_) + (d_585_v0_))
                rhs54_ = True
                rhs55_ = not ((self).f20) or (not((self).f20))
                lhs45_ = globalState
                lhs46_ = globalState
                lhs47_ = (self).f19
                lhs48_ = default__.safeIndex(883, ((self).f19).length(0))
                d_616_v11_ = rhs52_
                lhs45_.f0 = rhs53_
                lhs46_.f0 = rhs54_
                lhs47_[lhs48_] = rhs55_
                d_617_v12_: _dafny.Seq
                d_617_v12_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                d_618_v13_: _dafny.Map
                d_618_v13_ = _dafny.Map({((self).f19)[default__.safeIndex(883, ((self).f19).length(0))]: d_617_v12_})
                d_619_v14_: _dafny.Set
                d_619_v14_ = _dafny.Set({(((d_618_v13_)[p0] if (p0) in (d_618_v13_) else _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21]))) + (d_617_v12_), ((_dafny.SeqWithoutIsStrInference([430 for d_620_i0_ in range(default__.abs(292))])).set(default__.safeIndex((self).f21, len(_dafny.SeqWithoutIsStrInference([430 for d_621_i0_ in range(default__.abs(292))]))), (self).f21)) + (d_617_v12_)})
                rhs56_ = (len((d_616_v11_) | (d_616_v11_))) > (((0) - ((self).f21)) - ((self).f21))
                rhs57_ = _dafny.Set({default__.fm3(globalState)})
                rhs58_ = default__.safeDivisionInt((self).f21, len(d_608_v3_))
                lhs49_ = globalState
                lhs50_ = globalState
                lhs49_.f0 = rhs56_
                d_619_v14_ = rhs57_
                lhs50_.f13 = rhs58_
                d_622_v15_: _dafny.Array
                nw101_ = _dafny.Array(int(0), 5)
                d_622_v15_ = nw101_
                index81_ = default__.safeIndex(320, (d_622_v15_).length(0))
                (d_622_v15_)[index81_] = 789
                index82_ = default__.safeIndex(320, (d_622_v15_).length(0))
                rhs59_ = p0
                rhs60_ = (self).f21
                lhs51_ = globalState
                lhs52_ = d_622_v15_
                lhs53_ = default__.safeIndex(320, (d_622_v15_).length(0))
                lhs51_.f0 = rhs59_
                lhs52_[lhs53_] = rhs60_
                (globalState).f0 = (self).f20
                d_585_v0_ = d_585_v0_
            d_623_v16_: _dafny.Map
            d_623_v16_ = _dafny.Map({_dafny.CodePoint('p'): False})
            d_624_v17_: str
            d_624_v17_ = _dafny.CodePoint('j')
            (globalState).f0 = ((d_623_v16_)[d_624_v17_] if (d_624_v17_) in (d_623_v16_) else (self).f20)
        elif True:
            d_625_v18_: str
            d_625_v18_ = _dafny.CodePoint('x')
            d_626_v19_: _dafny.Map
            d_626_v19_ = _dafny.Map({d_625_v18_: p0})
            d_627_v20_: D0
            d_627_v20_ = D0_DC1(p0, (0) - ((d_586_v1_).cardinality), 360)
            d_628_v21_: _dafny.MultiSet
            d_628_v21_ = _dafny.MultiSet([d_627_v20_])
            d_629_v22_: _dafny.Seq
            d_629_v22_ = _dafny.SeqWithoutIsStrInference([p0, (self).f20])
            d_630_v23_: D0
            d_630_v23_ = D0_DC2(((d_626_v19_)[default__.fm4(globalState)] if (default__.fm4(globalState)) in (d_626_v19_) else (self).f20), (d_628_v21_).cardinality, d_629_v22_, len(d_585_v0_), (self).f21)
            d_630_v23_ = d_630_v23_
            (globalState).f0 = not (not((self).f20)) or (((self).f20 if (self).f20 else (self).f20))
            d_631_v24_: _dafny.Array
            nw102_ = _dafny.Array(int(0), 14)
            d_631_v24_ = nw102_
            (globalState).f3 = d_631_v24_
            if (self).f20:
                index83_ = default__.safeIndex(410, ((self).f19).length(0))
                ((self).f19)[index83_] = (True) == (not(p0))
                d_632_v25_: _dafny.Map
                d_632_v25_ = _dafny.Map({(self).f20: (self).f21})
                index84_ = default__.safeIndex(410, ((self).f19).length(0))
                rhs61_ = p0
                rhs62_ = (d_632_v25_) != (d_632_v25_)
                rhs63_ = (_dafny.SeqWithoutIsStrInference([(self).f20])) + (d_629_v22_)
                lhs54_ = globalState
                lhs55_ = (self).f19
                lhs56_ = default__.safeIndex(410, ((self).f19).length(0))
                lhs54_.f0 = rhs61_
                lhs55_[lhs56_] = rhs62_
                d_629_v22_ = rhs63_
                d_633_v26_: _dafny.Array
                nw103_ = _dafny.Array(None, 23)
                nw103_[int(0)] = ((self).f19)[default__.safeIndex(410, ((self).f19).length(0))]
                nw103_[int(1)] = (self).f20
                nw103_[int(2)] = False
                nw103_[int(3)] = True
                nw103_[int(4)] = False
                nw103_[int(5)] = p0
                nw103_[int(6)] = False
                nw103_[int(7)] = (self).f20
                nw103_[int(8)] = (self).f20
                nw103_[int(9)] = ((self).f19)[default__.safeIndex(410, ((self).f19).length(0))]
                nw103_[int(10)] = p0
                nw103_[int(11)] = p0
                nw103_[int(12)] = ((self).f19)[default__.safeIndex(410, ((self).f19).length(0))]
                nw103_[int(13)] = ((self).f19)[default__.safeIndex(410, ((self).f19).length(0))]
                nw103_[int(14)] = (self).f20
                nw103_[int(15)] = True
                nw103_[int(16)] = (self).f20
                nw103_[int(17)] = (self).f20
                nw103_[int(18)] = ((self).f19)[default__.safeIndex(410, ((self).f19).length(0))]
                nw103_[int(19)] = p0
                nw103_[int(20)] = ((self).f19)[default__.safeIndex(410, ((self).f19).length(0))]
                nw103_[int(21)] = p0
                nw103_[int(22)] = (self).f20
                d_633_v26_ = nw103_
                d_634_v27_: C1
                nw104_ = C1()
                nw104_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucn"))) + (d_585_v0_), d_633_v26_, False)
                d_634_v27_ = nw104_
                d_635_v28_: _dafny.Map
                d_635_v28_ = _dafny.Map({d_631_v24_: _dafny.CodePoint('e')})
                d_635_v28_ = (d_635_v28_).set(d_631_v24_, _dafny.CodePoint('g'))
                d_636_v29_: D4
                d_636_v29_ = D4_DC13()
                d_637_v30_: D4
                d_637_v30_ = D4_DC15(d_636_v29_)
                d_637_v30_ = d_637_v30_
                d_638_v31_: C0
                nw105_ = C0()
                nw105_.ctor__((self).f21)
                d_638_v31_ = nw105_
                d_639_v32_: C2
                nw106_ = C2()
                nw106_.ctor__(d_638_v31_, ((self).f19)[default__.safeIndex(410, ((self).f19).length(0))], d_633_v26_, (self).f20)
                d_639_v32_ = nw106_
            elif True:
                d_640_v33_: _dafny.MultiSet
                d_640_v33_ = _dafny.MultiSet([(self).f21])
                d_641_v34_: _dafny.Map
                d_641_v34_ = _dafny.Map({(0) - (((d_640_v33_).set((self).f21, default__.abs((self).f21))).cardinality): (self).f20})
                d_642_v35_: _dafny.Map
                d_642_v35_ = _dafny.Map({(self).f21: len(default__.fm8(d_585_v0_, d_641_v34_, (self).f21, globalState))})
                d_643_v36_: C0
                nw107_ = C0()
                nw107_.ctor__((self).f21)
                d_643_v36_ = nw107_
                d_644_v37_: T0
                nw108_ = C2()
                nw108_.ctor__(d_643_v36_, (self).f20, (self).f19, p0)
                d_644_v37_ = nw108_
                d_645_v38_: _dafny.Seq
                d_645_v38_ = _dafny.SeqWithoutIsStrInference([d_644_v37_, d_644_v37_, d_644_v37_, d_644_v37_, d_644_v37_])
                d_646_v39_: _dafny.MultiSet
                d_646_v39_ = _dafny.MultiSet([len(d_629_v22_), (self).f21, len(d_642_v35_), len(d_645_v38_)])
                (globalState).f13 = ((d_646_v39_).cardinality) - ((d_643_v36_).f24)
                d_647_v40_: _dafny.Map
                d_647_v40_ = _dafny.Map({p0: 126})
                d_647_v40_ = (d_647_v40_).set((self).f20, (d_643_v36_).f24)
                index85_ = default__.safeIndex(978, (d_631_v24_).length(0))
                (d_631_v24_)[index85_] = default__.safeModuloInt((self).f21, (self).f21)
                index86_ = default__.safeIndex(627, ((d_644_v37_).f19).length(0))
                ((d_644_v37_).f19)[index86_] = p0
                d_648_v41_: _dafny.Map
                d_648_v41_ = _dafny.Map({(d_644_v37_).f20: d_585_v0_})
                index87_ = default__.safeIndex(978, (d_631_v24_).length(0))
                index88_ = default__.safeIndex(627, ((d_644_v37_).f19).length(0))
                rhs64_ = 948
                rhs65_ = 542
                rhs66_ = (self).f21
                rhs67_ = default__.fm7(default__.safeDivisionInt((0) - (len(((d_648_v41_)[default__.fm7((self).f21, (self).f21, (self).f21, globalState)] if (default__.fm7((self).f21, (self).f21, (self).f21, globalState)) in (d_648_v41_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxledcrb"))))), ((d_642_v35_)[(self).f21] if ((self).f21) in (d_642_v35_) else (self).f21)), (d_643_v36_).f24, (self).f21, globalState)
                rhs68_ = (0) - ((self).f21)
                lhs57_ = globalState
                lhs58_ = globalState
                lhs59_ = d_631_v24_
                lhs60_ = default__.safeIndex(978, (d_631_v24_).length(0))
                lhs61_ = (d_644_v37_).f19
                lhs62_ = default__.safeIndex(627, ((d_644_v37_).f19).length(0))
                lhs63_ = globalState
                lhs57_.f18 = rhs64_
                lhs58_.f13 = rhs65_
                lhs59_[lhs60_] = rhs66_
                lhs61_[lhs62_] = rhs67_
                lhs63_.f18 = rhs68_
                d_649_v42_: D3
                d_649_v42_ = D3_DC9(d_585_v0_)
                d_650_v43_: _dafny.Array
                nw109_ = _dafny.Array(None, 17)
                nw109_[int(0)] = default__.fm0(((d_644_v37_).f19)[default__.safeIndex(627, ((d_644_v37_).f19).length(0))], (self).f20, globalState)
                nw109_[int(1)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xw")))
                nw109_[int(2)] = 836
                nw109_[int(3)] = 350
                nw109_[int(4)] = (self).f21
                nw109_[int(5)] = len(d_585_v0_)
                nw109_[int(6)] = (d_631_v24_)[default__.safeIndex(978, (d_631_v24_).length(0))]
                nw109_[int(7)] = (self).f21
                nw109_[int(8)] = (self).f21
                nw109_[int(9)] = (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([d_625_v18_ for d_651_i1_ in range(default__.abs(233))]))))
                nw109_[int(10)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqgtfnndn")))
                nw109_[int(11)] = default__.safeModuloInt((d_643_v36_).f24, (d_643_v36_).f24)
                nw109_[int(12)] = (len((d_649_v42_).cf16)) + (len(p1))
                nw109_[int(13)] = (d_631_v24_)[default__.safeIndex(978, (d_631_v24_).length(0))]
                nw109_[int(14)] = (0) - ((self).f21)
                nw109_[int(15)] = (d_643_v36_).f24
                nw109_[int(16)] = ((self).f21) - ((d_631_v24_)[default__.safeIndex(978, (d_631_v24_).length(0))])
                d_650_v43_ = nw109_
                (globalState).f10 = d_650_v43_
                index89_ = default__.safeIndex(627, ((d_644_v37_).f19).length(0))
                ((d_644_v37_).f19)[index89_] = not((p0) == (default__.fm7((self).f21, len(_dafny.Set({(d_644_v37_).f20, (d_644_v37_).f20, False})), (d_643_v36_).f24, globalState)))
            d_652_v44_: D3
            d_652_v44_ = D3_DC9(_dafny.SeqWithoutIsStrInference([d_625_v18_ for d_653_i2_ in range(default__.abs(990))]))
            pat_let_tv27_ = d_652_v44_
            pat_let_tv28_ = d_585_v0_
            pat_let_tv29_ = d_585_v0_
            pat_let_tv30_ = d_625_v18_
            d_654_v45_: _dafny.Array
            nw110_ = _dafny.Array(None, 22)
            nw110_[int(0)] = d_652_v44_
            nw110_[int(1)] = d_652_v44_
            nw110_[int(2)] = d_652_v44_
            nw110_[int(3)] = d_652_v44_
            nw110_[int(4)] = d_652_v44_
            nw110_[int(5)] = d_652_v44_
            nw110_[int(6)] = D3_DC9(d_585_v0_)
            nw110_[int(7)] = d_652_v44_
            nw110_[int(8)] = d_652_v44_
            def iife53_(_pat_let20_0):
                def iife54_(d_655_dt__update__tmp_h0_):
                    def iife55_(_pat_let21_0):
                        def iife56_(d_656_dt__update_hcf16_h0_):
                            return D3_DC9(d_656_dt__update_hcf16_h0_)
                        return iife56_(_pat_let21_0)
                    return iife55_((pat_let_tv27_).cf16)
                return iife54_(_pat_let20_0)
            nw110_[int(9)] = iife53_(d_652_v44_)
            nw110_[int(10)] = d_652_v44_
            nw110_[int(11)] = d_652_v44_
            nw110_[int(12)] = d_652_v44_
            nw110_[int(13)] = d_652_v44_
            nw110_[int(14)] = d_652_v44_
            nw110_[int(15)] = d_652_v44_
            nw110_[int(16)] = d_652_v44_
            nw110_[int(17)] = (d_652_v44_ if p0 else d_652_v44_)
            nw110_[int(18)] = d_652_v44_
            nw110_[int(19)] = d_652_v44_
            def iife57_(_pat_let22_0):
                def iife58_(d_657_dt__update__tmp_h1_):
                    def iife59_(_pat_let23_0):
                        def iife60_(d_658_dt__update_hcf16_h1_):
                            return D3_DC9(d_658_dt__update_hcf16_h1_)
                        return iife60_(_pat_let23_0)
                    return iife59_((pat_let_tv28_).set(default__.safeIndex((self).f21, len(pat_let_tv29_)), pat_let_tv30_))
                return iife58_(_pat_let22_0)
            nw110_[int(20)] = iife57_(d_652_v44_)
            nw110_[int(21)] = d_652_v44_
            d_654_v45_ = nw110_
            d_654_v45_ = d_654_v45_
        (globalState).f18 = (self).f21
        d_659_v46_: C0
        nw111_ = C0()
        nw111_.ctor__(107)
        d_659_v46_ = nw111_
        d_660_v47_: T0
        nw112_ = C2()
        nw112_.ctor__(d_659_v46_, (self).f20, (self).f19, (self).f20)
        d_660_v47_ = nw112_
        d_661_v48_: _dafny.Map
        d_661_v48_ = _dafny.Map({d_660_v47_: (self).f20})
        d_662_v49_: _dafny.Map
        d_662_v49_ = _dafny.Map({((d_661_v48_)[d_660_v47_] if (d_660_v47_) in (d_661_v48_) else not(p0)): ((d_659_v46_).f24) == ((d_659_v46_).f24)})
        d_663_v50_: D4
        d_663_v50_ = D4_DC14(p0, p1)
        pat_let_tv31_ = p1
        def iife61_(_pat_let24_0):
            def iife62_(d_664_dt__update__tmp_h2_):
                def iife63_(_pat_let25_0):
                    def iife64_(d_665_dt__update_hcf26_h0_):
                        return D4_DC14((d_664_dt__update__tmp_h2_).cf25, d_665_dt__update_hcf26_h0_)
                    return iife64_(_pat_let25_0)
                return iife63_(pat_let_tv31_)
            return iife62_(_pat_let24_0)
        d_662_v49_ = (d_662_v49_).set((iife61_(d_663_v50_)).cf25, p0)

    def m1(self, globalState):
        r0: int = int(0)
        (globalState).f13 = ((self).f21) * ((self).f21)
        hi5_ = (self).f21
        for d_666_i0_ in range((self).f21, hi5_):
            d_667_v0_: _dafny.Seq
            d_667_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
            d_668_v1_: C1
            nw113_ = C1()
            nw113_.ctor__(d_667_v0_, (self).f19, (self).f20)
            d_668_v1_ = nw113_
            d_669_v2_: C3
            nw114_ = C3()
            nw114_.ctor__(((d_668_v1_).f27) + (d_667_v0_), (self).f19)
            d_669_v2_ = nw114_
            d_670_v3_: D2
            d_670_v3_ = D2_DC8(((self).f21 if (self).f20 else d_666_i0_), (self).f21)
            d_671_v4_: _dafny.Array
            nw115_ = _dafny.Array(int(0), 28)
            d_671_v4_ = nw115_
            index90_ = default__.safeIndex(86, (d_671_v4_).length(0))
            (d_671_v4_)[index90_] = d_666_i0_
            index91_ = default__.safeIndex(722, (d_671_v4_).length(0))
            (d_671_v4_)[index91_] = d_666_i0_
            d_672_v5_: _dafny.Seq
            d_672_v5_ = _dafny.SeqWithoutIsStrInference([(self).f20, (self).f20])
            d_673_v6_: _dafny.MultiSet
            d_673_v6_ = _dafny.MultiSet([(self).f21])
            pat_let_tv32_ = d_673_v6_
            index92_ = default__.safeIndex(86, (d_671_v4_).length(0))
            index93_ = default__.safeIndex(722, (d_671_v4_).length(0))
            def iife65_(_pat_let26_0):
                def iife66_(d_674_dt__update__tmp_h0_):
                    def iife67_(_pat_let27_0):
                        def iife68_(d_675_dt__update_hcf15_h0_):
                            return D2_DC8((d_674_dt__update__tmp_h0_).cf14, d_675_dt__update_hcf15_h0_)
                        return iife68_(_pat_let27_0)
                    return iife67_(((pat_let_tv32_) | (_dafny.MultiSet([28]))).cardinality)
                return iife66_(_pat_let26_0)
            rhs69_ = (0) - (len(_dafny.SeqWithoutIsStrInference([d_672_v5_, d_672_v5_])))
            rhs70_ = iife65_(default__.fm23(globalState))
            rhs71_ = len(d_669_v2_.f22)
            rhs72_ = default__.safeModuloInt(d_666_i0_, (self).f21)
            rhs73_ = (self).f20
            lhs64_ = globalState
            lhs65_ = d_671_v4_
            lhs66_ = default__.safeIndex(86, (d_671_v4_).length(0))
            lhs67_ = d_671_v4_
            lhs68_ = default__.safeIndex(722, (d_671_v4_).length(0))
            lhs69_ = globalState
            lhs64_.f13 = rhs69_
            d_670_v3_ = rhs70_
            lhs65_[lhs66_] = rhs71_
            lhs67_[lhs68_] = rhs72_
            lhs69_.f0 = rhs73_
            if (self).f20:
                d_676_v7_: _dafny.Map
                d_676_v7_ = _dafny.Map({(self).f20: (self).f20})
                d_677_v8_: _dafny.Seq
                d_677_v8_ = _dafny.SeqWithoutIsStrInference([(d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))], (self).f21, (self).f21, (self).f21, (d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))]])
                d_678_v9_: _dafny.Array
                nw116_ = _dafny.Array(None, 28)
                nw116_[int(0)] = (self).f20
                nw116_[int(1)] = not(((d_676_v7_)[(self).f20] if ((self).f20) in (d_676_v7_) else (self).f20))
                nw116_[int(2)] = (d_669_v2_.f22) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_679_i1_ in range(default__.abs(563))]))
                nw116_[int(3)] = (self).f20
                nw116_[int(4)] = (self).f20
                nw116_[int(5)] = not(True)
                nw116_[int(6)] = (self).f20
                nw116_[int(7)] = not((_dafny.MultiSet(d_677_v8_)).issubset(_dafny.MultiSet([(d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))]])))
                nw116_[int(8)] = (self).f20
                nw116_[int(9)] = (self).f20
                nw116_[int(10)] = (self).f20
                nw116_[int(11)] = False
                nw116_[int(12)] = (d_668_v1_).fm13(globalState)
                nw116_[int(13)] = (self).f20
                nw116_[int(14)] = (self).f20
                nw116_[int(15)] = (self).f20
                nw116_[int(16)] = (d_673_v6_).isdisjoint(d_673_v6_)
                nw116_[int(17)] = ((self).f21) > (d_666_i0_)
                nw116_[int(18)] = (self).f20
                nw116_[int(19)] = (self).f20
                nw116_[int(20)] = (self).f20
                nw116_[int(21)] = (self).f20
                nw116_[int(22)] = (self).f20
                nw116_[int(23)] = (self).f20
                nw116_[int(24)] = (self).f20
                nw116_[int(25)] = not(True)
                nw116_[int(26)] = (self).f20
                nw116_[int(27)] = (self).f20
                d_678_v9_ = nw116_
                d_678_v9_ = (self).f19
                d_680_v10_: str
                d_680_v10_ = _dafny.CodePoint('q')
                d_681_v12_: _dafny.Seq
                def iife69_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in (d_677_v8_).Elements:
                        d_682_v11_: int = compr_13_
                        if (d_682_v11_) in (d_677_v8_):
                            coll13_[default__.safeModuloInt(d_682_v11_, (d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))])] = (self).f20
                    return _dafny.Map(coll13_)
                d_681_v12_ = _dafny.SeqWithoutIsStrInference([iife69_()
                ])
                d_683_v13_: _dafny.Map
                d_683_v13_ = _dafny.Map({(self).f21: (self).f20})
                d_684_v14_: C3
                nw117_ = C3()
                nw117_.ctor__(default__.fm8(default__.fm8(_dafny.SeqWithoutIsStrInference([d_680_v10_]), (d_681_v12_)[default__.safeIndex(d_666_i0_, len(d_681_v12_))], len(d_669_v2_.f22), globalState), d_683_v13_, (self).f21, globalState), d_678_v9_)
                d_684_v14_ = nw117_
                d_685_v15_: _dafny.Set
                d_685_v15_ = _dafny.Set({(d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))]})
                d_686_v16_: _dafny.Map
                d_686_v16_ = _dafny.Map({d_680_v10_: (d_685_v15_).isdisjoint(_dafny.Set({(d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))], len(d_669_v2_.f22)}))})
                d_686_v16_ = d_686_v16_
                index94_ = default__.safeIndex(86, (d_671_v4_).length(0))
                (d_671_v4_)[index94_] = (0) - (-178)
                d_687_v17_: _dafny.Array
                nw118_ = _dafny.Array(D1.default()(), 2)
                d_687_v17_ = nw118_
                d_688_v18_: D1
                d_688_v18_ = D1_DC5((self).f20)
                index95_ = default__.safeIndex(29, (d_687_v17_).length(0))
                (d_687_v17_)[index95_] = d_688_v18_
                d_689_v19_: D0
                d_689_v19_ = D0_DC2((self).f20, (self).f21, d_672_v5_, default__.safeDivisionInt((d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))], 853), (0) - (len(d_669_v2_.f22)))
                d_690_v20_: _dafny.Map
                d_690_v20_ = _dafny.Map({(d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))]: d_666_i0_})
                d_691_v21_: _dafny.MultiSet
                d_691_v21_ = _dafny.MultiSet([((d_668_v1_).f27) != (d_669_v2_.f22), (d_672_v5_)[default__.safeIndex(d_666_i0_, len(d_672_v5_))], not((self).f20), (d_666_i0_) == ((d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))]), (self).f20])
                d_692_v22_: _dafny.MultiSet
                d_692_v22_ = _dafny.MultiSet([d_680_v10_, d_680_v10_])
                index96_ = default__.safeIndex(29, (d_687_v17_).length(0))
                rhs74_ = D1_DC5((self).f20)
                rhs75_ = d_689_v19_
                rhs76_ = (d_690_v20_) | (_dafny.Map({((d_692_v22_)[d_680_v10_] if (d_680_v10_) in (d_692_v22_) else 817): d_666_i0_}))
                rhs77_ = (self).f20
                rhs78_ = (d_691_v21_).intersection((d_691_v21_).intersection(d_691_v21_))
                lhs70_ = d_687_v17_
                lhs71_ = default__.safeIndex(29, (d_687_v17_).length(0))
                lhs72_ = globalState
                lhs70_[lhs71_] = rhs74_
                d_689_v19_ = rhs75_
                d_690_v20_ = rhs76_
                lhs72_.f0 = rhs77_
                d_691_v21_ = rhs78_
            elif True:
                r0 = d_666_i0_
                rhs79_ = d_672_v5_
                rhs80_ = (self).f20
                lhs73_ = globalState
                d_672_v5_ = rhs79_
                lhs73_.f0 = rhs80_
                (d_669_v2_).m3(globalState)
                d_693_v23_: _dafny.Map
                d_693_v23_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqw")): (d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))]})
                d_694_v24_: _dafny.Map
                d_694_v24_ = _dafny.Map({(self).f21: (d_668_v1_).f27})
                d_695_v25_: _dafny.Map
                d_695_v25_ = _dafny.Map({False: (d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))]})
                d_696_v26_: _dafny.MultiSet
                d_696_v26_ = _dafny.MultiSet([_dafny.Map({(self).f20: (d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))]}), (d_695_v25_).set((self).f20, d_666_i0_)])
                d_697_v27_: _dafny.Map
                d_697_v27_ = _dafny.Map({-831: ((d_696_v26_)[d_695_v25_] if (d_695_v25_) in (d_696_v26_) else -915)})
                d_693_v23_ = (d_693_v23_).set((d_667_v0_) + (((d_694_v24_)[d_666_i0_] if (d_666_i0_) in (d_694_v24_) else d_667_v0_)), len((d_697_v27_).set((d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))], (d_671_v4_)[default__.safeIndex(86, (d_671_v4_).length(0))])))
                d_698_v28_: str
                d_698_v28_ = _dafny.CodePoint('y')
                d_698_v28_ = d_698_v28_
        d_699_v29_: _dafny.Array
        nw119_ = _dafny.Array(None, 4)
        nw119_[int(0)] = (self).f21
        nw119_[int(1)] = (self).f21
        nw119_[int(2)] = (self).f21
        nw119_[int(3)] = default__.fm0((self).f20, (self).f20, globalState)
        d_699_v29_ = nw119_
        index97_ = default__.safeIndex(183, (d_699_v29_).length(0))
        (d_699_v29_)[index97_] = (self).f21
        index98_ = default__.safeIndex(183, (d_699_v29_).length(0))
        (d_699_v29_)[index98_] = (self).f21
        d_700_v30_: str
        d_700_v30_ = _dafny.CodePoint('a')
        d_701_v31_: D4
        d_701_v31_ = D4_DC12(d_700_v30_)
        d_702_v32_: _dafny.Seq
        d_702_v32_ = _dafny.SeqWithoutIsStrInference([d_700_v30_, (d_701_v31_).cf24])
        d_703_v33_: _dafny.Map
        d_703_v33_ = _dafny.Map({968: False})
        d_704_v34_: T0
        nw120_ = C1()
        nw120_.ctor__((default__.fm8(d_702_v32_, d_703_v33_, (self).f21, globalState)) + (d_702_v32_), ((self).f19 if False else (self).f19), (self).f20)
        d_704_v34_ = nw120_
        d_705_v35_: D0
        d_705_v35_ = D0_DC1(True, (self).f21, (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))])
        nw121_ = C1()
        nw121_.ctor__(d_702_v32_, ((d_704_v34_).f19 if (D5_DC17((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qedpt"))), d_705_v35_, (self).f20, True)).cf32 else (self).f19), (d_704_v34_).f20)
        d_704_v34_ = nw121_
        hi6_ = (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]
        for d_706_i2_ in range((len(_dafny.SeqWithoutIsStrInference([d_700_v30_ for d_713_i3_ in range(default__.abs(-527))]))) - ((self).f21), hi6_):
            d_707_v36_: _dafny.MultiSet
            d_707_v36_ = _dafny.MultiSet([(d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))], default__.fm0((d_704_v34_).f20, (self).f20, globalState)])
            d_708_v37_: _dafny.Map
            d_708_v37_ = _dafny.Map({default__.fm7((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))], (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))], d_706_i2_, globalState): d_707_v36_})
            d_708_v37_ = _dafny.Map({(self).f20: (d_707_v36_) | (d_707_v36_)})
            (globalState).f0 = (self).f20
            d_709_v38_: _dafny.MultiSet
            d_709_v38_ = _dafny.MultiSet([(d_702_v32_) + (_dafny.SeqWithoutIsStrInference([d_700_v30_ for d_710_i4_ in range(default__.abs(42))])), (d_702_v32_) + (d_702_v32_), d_702_v32_, d_702_v32_, ((d_702_v32_ if (d_704_v34_).f20 else d_702_v32_)).set(default__.safeIndex((self).f21, len((d_702_v32_ if (d_704_v34_).f20 else d_702_v32_))), d_700_v30_)])
            d_709_v38_ = d_709_v38_
            d_711_v39_: _dafny.Array
            nw122_ = _dafny.Array(False, 12)
            d_711_v39_ = nw122_
            index99_ = default__.safeIndex(306, ((d_704_v34_).f19).length(0))
            ((d_704_v34_).f19)[index99_] = not((self).f20)
            d_712_v40_: _dafny.MultiSet
            d_712_v40_ = _dafny.MultiSet([(d_704_v34_).f20, (d_704_v34_).f20, False])
            index100_ = default__.safeIndex(183, (d_699_v29_).length(0))
            index101_ = default__.safeIndex(306, ((d_704_v34_).f19).length(0))
            rhs81_ = (d_704_v34_).f20
            rhs82_ = (d_704_v34_).f19
            rhs83_ = default__.safeDivisionInt((self).f21, (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))])
            rhs84_ = (d_712_v40_).isdisjoint((d_712_v40_).set((d_704_v34_).f20, default__.abs(d_706_i2_)))
            lhs74_ = globalState
            lhs75_ = d_699_v29_
            lhs76_ = default__.safeIndex(183, (d_699_v29_).length(0))
            lhs77_ = (d_704_v34_).f19
            lhs78_ = default__.safeIndex(306, ((d_704_v34_).f19).length(0))
            lhs74_.f0 = rhs81_
            d_711_v39_ = rhs82_
            lhs75_[lhs76_] = rhs83_
            lhs77_[lhs78_] = rhs84_
        hi7_ = (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]
        for d_714_i5_ in range((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))], hi7_):
            if (((self).f21) <= ((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))])) == ((self).f20):
                d_715_v41_: _dafny.Map
                d_715_v41_ = _dafny.Map({(d_704_v34_).f20: (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]})
                d_716_v42_: _dafny.Map
                d_716_v42_ = _dafny.Map({(self).f21: (self).f21})
                (globalState).f13 = ((d_715_v41_)[(True if False else (d_704_v34_).f20)] if ((True if False else (d_704_v34_).f20)) in (d_715_v41_) else len((d_716_v42_) | (d_716_v42_)))
                index102_ = default__.safeIndex(183, (d_699_v29_).length(0))
                (d_699_v29_)[index102_] = default__.safeDivisionInt((0) - ((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]), (0) - ((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]))
                (globalState).f0 = (self).f20
                d_717_v43_: _dafny.Array
                def lambda24_(d_718_v34_):
                    def lambda25_(d_719_i6_):
                        return (d_718_v34_).f20

                    return lambda25_

                init14_ = lambda24_(d_704_v34_)
                nw123_ = _dafny.Array(None, 5)
                for i0_14_ in range(nw123_.length(0)):
                    nw123_[i0_14_] = init14_(i0_14_)
                d_717_v43_ = nw123_
                d_720_v44_: C0
                nw124_ = C0()
                nw124_.ctor__((self).f21)
                d_720_v44_ = nw124_
                d_721_v45_: C2
                nw125_ = C2()
                nw125_.ctor__(d_720_v44_, (d_704_v34_).f20, d_717_v43_, (d_704_v34_).f20)
                d_721_v45_ = nw125_
                d_722_v46_: _dafny.Array
                nw126_ = _dafny.Array(D0.default()(), 2)
                d_722_v46_ = nw126_
                d_723_v47_: _dafny.Array
                d_723_v47_ = d_722_v46_
                d_724_v48_: _dafny.Array
                nw127_ = _dafny.Array(None, 18)
                nw127_[int(0)] = d_723_v47_
                nw127_[int(1)] = d_722_v46_
                nw127_[int(2)] = d_722_v46_
                nw127_[int(3)] = d_723_v47_
                nw127_[int(4)] = d_723_v47_
                nw127_[int(5)] = d_723_v47_
                nw127_[int(6)] = d_723_v47_
                nw127_[int(7)] = d_723_v47_
                nw127_[int(8)] = d_723_v47_
                nw127_[int(9)] = d_723_v47_
                nw127_[int(10)] = d_723_v47_
                nw127_[int(11)] = d_723_v47_
                nw127_[int(12)] = d_723_v47_
                nw127_[int(13)] = d_722_v46_
                nw127_[int(14)] = (d_723_v47_ if (d_704_v34_).f20 else d_722_v46_)
                nw127_[int(15)] = d_722_v46_
                nw127_[int(16)] = d_723_v47_
                nw127_[int(17)] = (d_722_v46_ if (d_704_v34_).f20 else d_722_v46_)
                d_724_v48_ = nw127_
                index103_ = default__.safeIndex(250, (d_724_v48_).length(0))
                (d_724_v48_)[index103_] = d_723_v47_
                index104_ = default__.safeIndex(183, (d_699_v29_).length(0))
                index105_ = default__.safeIndex(250, (d_724_v48_).length(0))
                rhs85_ = d_717_v43_
                rhs86_ = d_721_v45_
                rhs87_ = (self).f21
                rhs88_ = d_723_v47_
                lhs79_ = d_699_v29_
                lhs80_ = default__.safeIndex(183, (d_699_v29_).length(0))
                lhs81_ = d_724_v48_
                lhs82_ = default__.safeIndex(250, (d_724_v48_).length(0))
                d_717_v43_ = rhs85_
                d_721_v45_ = rhs86_
                lhs79_[lhs80_] = rhs87_
                lhs81_[lhs82_] = rhs88_
                d_715_v41_ = (d_715_v41_).set((d_702_v32_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnij"))), (d_720_v44_).f24)
            elif True:
                d_725_v49_: _dafny.MultiSet
                d_725_v49_ = _dafny.MultiSet([(self).f20, (self).f20])
                d_726_v50_: D0
                d_726_v50_ = D0_DC0(d_725_v49_)
                d_726_v50_ = d_726_v50_
                (globalState).f0 = (self).f20
                d_727_v51_: _dafny.Array
                nw128_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_727_v51_ = nw128_
                d_727_v51_ = d_727_v51_
                d_728_v52_: _dafny.Map
                d_728_v52_ = _dafny.Map({(d_704_v34_).f20: (0) - ((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))])})
                d_725_v49_ = ((d_725_v49_).set((d_704_v34_).f20, default__.abs(len(d_728_v52_)))) | (d_725_v49_)
                d_729_v53_: C0
                nw129_ = C0()
                nw129_.ctor__(len(default__.fm19(False, globalState)))
                d_729_v53_ = nw129_
            d_730_v54_: _dafny.Seq
            d_730_v54_ = _dafny.SeqWithoutIsStrInference([(self).f20])
            d_731_v55_: _dafny.Map
            d_731_v55_ = _dafny.Map({(d_702_v32_) + (d_702_v32_): (default__.fm0((self).f20, (d_704_v34_).f20, globalState)) + (len(default__.fm10(len(d_730_v54_), 462, (0) - ((0) - (-433)), globalState)))})
            rhs89_ = d_731_v55_
            rhs90_ = (d_704_v34_).f20
            lhs83_ = globalState
            d_731_v55_ = rhs89_
            lhs83_.f0 = rhs90_
            d_732_v56_: _dafny.Seq
            d_732_v56_ = _dafny.SeqWithoutIsStrInference([d_702_v32_, d_702_v32_])
            d_733_v57_: _dafny.Map
            d_733_v57_ = _dafny.Map({d_732_v56_: _dafny.SeqWithoutIsStrInference([d_700_v30_ for d_734_i7_ in range(default__.abs(182))])})
            d_733_v57_ = (d_733_v57_).set((d_732_v56_ if default__.fm9((self).f20, globalState) else d_732_v56_), _dafny.SeqWithoutIsStrInference([d_700_v30_ for d_735_i8_ in range(default__.abs(-291))]))
            if (self).f20:
                d_736_v58_: _dafny.MultiSet
                d_736_v58_ = _dafny.MultiSet([(self).f20])
                d_737_v59_: D0
                d_737_v59_ = D0_DC0(d_736_v58_)
                pat_let_tv33_ = d_736_v58_
                d_738_v60_: _dafny.Map
                def iife70_(_pat_let28_0):
                    def iife71_(d_739_dt__update__tmp_h4_):
                        def iife72_(_pat_let29_0):
                            def iife73_(d_740_dt__update_hcf0_h0_):
                                return D0_DC0(d_740_dt__update_hcf0_h0_)
                            return iife73_(_pat_let29_0)
                        return iife72_(pat_let_tv33_)
                    return iife71_(_pat_let28_0)
                d_738_v60_ = _dafny.Map({False: iife70_(d_737_v59_)})
                d_741_v61_: _dafny.Seq
                d_741_v61_ = _dafny.SeqWithoutIsStrInference([d_738_v60_, _dafny.Map({False: d_737_v59_})])
                (globalState).f0 = ((d_741_v61_) + (d_741_v61_)) != (d_741_v61_)
                (globalState).f18 = (0) - (((0) - (default__.fm0((d_704_v34_).f20, True, globalState))) - ((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]))
                index106_ = default__.safeIndex(183, (d_699_v29_).length(0))
                (d_699_v29_)[index106_] = 260
                d_742_v62_: _dafny.Array
                nw130_ = _dafny.Array(_dafny.Map({}), 24)
                d_742_v62_ = nw130_
                d_743_v63_: _dafny.Map
                d_743_v63_ = _dafny.Map({default__.fm9((d_704_v34_).f20, globalState): d_714_i5_})
                d_744_v64_: D0
                d_744_v64_ = D0_DC2((self).f20, d_714_i5_, d_730_v54_, (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))], len(_dafny.SeqWithoutIsStrInference([(self).f20])))
                d_745_v65_: _dafny.Seq
                d_745_v65_ = _dafny.SeqWithoutIsStrInference([d_744_v64_])
                d_746_v66_: _dafny.Set
                d_746_v66_ = _dafny.Set({D1_DC4(d_745_v65_)})
                d_747_v67_: _dafny.Set
                d_747_v67_ = _dafny.Set({(self).f20, (self).f20, (d_704_v34_).f20})
                d_748_v68_: D4
                d_748_v68_ = D4_DC14(not(False), d_747_v67_)
                d_703_v33_ = (d_703_v33_).set((D3_DC10(d_742_v62_, d_743_v63_, d_746_v66_, (self).f21)).cf20, (d_748_v68_).cf25)
                d_749_v69_: _dafny.Array
                nw131_ = _dafny.Array(None, 5)
                nw131_[int(0)] = d_700_v30_
                nw131_[int(1)] = d_700_v30_
                nw131_[int(2)] = _dafny.CodePoint('o')
                nw131_[int(3)] = d_700_v30_
                nw131_[int(4)] = (d_702_v32_)[default__.safeIndex((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))], len(d_702_v32_))]
                d_749_v69_ = nw131_
                index107_ = default__.safeIndex(855, (d_749_v69_).length(0))
                (d_749_v69_)[index107_] = d_700_v30_
                index108_ = default__.safeIndex(855, (d_749_v69_).length(0))
                (d_749_v69_)[index108_] = _dafny.CodePoint('u')
            elif True:
                d_750_v70_: _dafny.MultiSet
                d_750_v70_ = _dafny.MultiSet([(d_714_i5_) < (548), (d_704_v34_).f20, (d_704_v34_).f20, ((d_704_v34_).f20) == ((self).f20)])
                d_750_v70_ = (d_750_v70_) | ((d_750_v70_ if (d_704_v34_).f20 else d_750_v70_))
                d_751_v71_: C0
                nw132_ = C0()
                nw132_.ctor__((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))])
                d_751_v71_ = nw132_
                d_752_v72_: _dafny.Map
                d_752_v72_ = _dafny.Map({(d_704_v34_).f20: d_751_v71_})
                d_753_v73_: _dafny.Seq
                d_753_v73_ = _dafny.SeqWithoutIsStrInference([d_751_v71_])
                d_754_v74_: _dafny.Map
                d_754_v74_ = _dafny.Map({(self).f20: (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]})
                d_755_v75_: _dafny.Array
                nw133_ = _dafny.Array(None, 17)
                nw133_[int(0)] = d_752_v72_
                nw133_[int(1)] = d_752_v72_
                nw133_[int(2)] = d_752_v72_
                nw133_[int(3)] = d_752_v72_
                nw133_[int(4)] = d_752_v72_
                nw133_[int(5)] = d_752_v72_
                nw133_[int(6)] = d_752_v72_
                nw133_[int(7)] = d_752_v72_
                nw133_[int(8)] = _dafny.Map({(d_704_v34_).f20: (d_753_v73_)[default__.safeIndex(len(d_754_v74_), len(d_753_v73_))]})
                nw133_[int(9)] = d_752_v72_
                nw133_[int(10)] = d_752_v72_
                nw133_[int(11)] = d_752_v72_
                nw133_[int(12)] = d_752_v72_
                nw133_[int(13)] = d_752_v72_
                nw133_[int(14)] = d_752_v72_
                nw133_[int(15)] = d_752_v72_
                nw133_[int(16)] = d_752_v72_
                d_755_v75_ = nw133_
                d_756_v76_: D1
                d_756_v76_ = D1_DC4(default__.fm24(globalState))
                d_757_v77_: D0
                d_757_v77_ = D0_DC2((d_704_v34_).f20, (d_751_v71_).f24, d_730_v54_, 604, (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))])
                d_758_v78_: _dafny.Seq
                d_758_v78_ = _dafny.SeqWithoutIsStrInference([d_757_v77_])
                pat_let_tv34_ = d_758_v78_
                d_759_v79_: _dafny.Set
                def iife74_(_pat_let30_0):
                    def iife75_(d_760_dt__update__tmp_h5_):
                        def iife76_(_pat_let31_0):
                            def iife77_(d_761_dt__update_hcf10_h0_):
                                return D1_DC4(d_761_dt__update_hcf10_h0_)
                            return iife77_(_pat_let31_0)
                        return iife76_(pat_let_tv34_)
                    return iife75_(_pat_let30_0)
                d_759_v79_ = _dafny.Set({iife74_(d_756_v76_), D1_DC4(d_758_v78_)})
                d_762_v80_: D3
                d_762_v80_ = D3_DC10(d_755_v75_, _dafny.Map({True: (self).f21}), d_759_v79_, -727)
                index109_ = default__.safeIndex(183, (d_699_v29_).length(0))
                (d_699_v29_)[index109_] = (d_762_v80_).cf20
                def iife78_():
                    coll14_ = _dafny.Map()
                    compr_14_: int
                    for compr_14_ in _dafny.IntegerRange(527, 117):
                        d_763_v81_: int = compr_14_
                        if ((527) <= (d_763_v81_)) and ((d_763_v81_) < (117)):
                            coll14_[(d_763_v81_) * ((d_751_v71_).f24)] = (d_704_v34_).f20
                    return _dafny.Map(coll14_)
                (globalState).f0 = ((d_704_v34_).f20 if (iife78_()
                ) not in (_dafny.SeqWithoutIsStrInference([default__.fm25(globalState)])) else (d_704_v34_).f20)
                index110_ = default__.safeIndex(183, (d_699_v29_).length(0))
                (d_699_v29_)[index110_] = (d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]
                d_750_v70_ = _dafny.MultiSet([(self).f20])
        r0 = default__.safeModuloInt(((self).f21) + ((d_699_v29_)[default__.safeIndex(183, (d_699_v29_).length(0))]), (self).f21)
        return r0

    @property
    def f21(self):
        return self._f21
