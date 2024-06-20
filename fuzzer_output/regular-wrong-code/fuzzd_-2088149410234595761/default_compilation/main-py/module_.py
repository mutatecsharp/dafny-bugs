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
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(565, 642):
                d_0_v0_: int = compr_0_
                if ((565) <= (d_0_v0_)) and ((d_0_v0_) < (642)):
                    def iife1_():
                        coll1_ = _dafny.Map()
                        compr_1_: _dafny.Seq
                        for compr_1_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])])).Elements:
                            d_1_v1_: _dafny.Seq = compr_1_
                            if (d_1_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])])):
                                coll1_[d_1_v1_] = 24
                        return _dafny.Map(coll1_)
                    coll0_[default__.safeDivisionInt(d_0_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doerfd"))))] = len(iife1_()
                    )
            return _dafny.Map(coll0_)
        return default__.safeModuloInt(833, len(iife0_()
        ))

    @staticmethod
    def fm1(globalState):
        return _dafny.CodePoint('w')

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([-329, (0) - ((0) - ((_dafny.MultiSet([True])).cardinality)), 696, len(_dafny.Set({True, True, not(True)}))])) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omngle")) for d_2_i0_ in range(default__.abs(982))]))])).intersection(_dafny.MultiSet([len(_dafny.Map({(0) - (-597): not(False)}))])))

    @staticmethod
    def fm4(p0, globalState):
        return ((-92 if True else len(_dafny.Map({not(True): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_3_i0_ in range(default__.abs(605))])))]))).cardinality})))) <= (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esnepq"))), 414))

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([False, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aahphab")))) == ((0) - ((0) - (-978))), not(True), (-747) >= (-992), (True) and (not(not(True)))])

    @staticmethod
    def fm6(p0, p1, globalState):
        source0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkgbfxm"))
        d_4___mcc_h0_ = source0_
        d_5_cf5_ = d_4___mcc_h0_
        return d_5_cf5_

    @staticmethod
    def fm7(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, True, not(False), True])).cardinality, len(_dafny.SeqWithoutIsStrInference([True, True]))])) + (_dafny.SeqWithoutIsStrInference([-334]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cudaub")))])) + (_dafny.SeqWithoutIsStrInference([650])))

    @staticmethod
    def fm8(globalState):
        return ((_dafny.MultiSet([False])) | (_dafny.MultiSet([True, False, True]))) - (_dafny.MultiSet([True, True]))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        (globalState).f2 = p0
        d_6_v0_: bool
        d_6_v0_ = False
        d_7_v1_: _dafny.Seq
        d_7_v1_ = _dafny.SeqWithoutIsStrInference([d_6_v0_, d_6_v0_, d_6_v0_])
        d_7_v1_ = d_7_v1_
        d_8_v2_: _dafny.Set
        d_8_v2_ = _dafny.Set({p1})
        d_9_i0_: int
        d_9_i0_ = 0
        with _dafny.label("0"):
            while not(not(((d_8_v2_).issubset(d_8_v2_)) or (True))):
                with _dafny.c_label("0"):
                    if (d_9_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_9_i0_ = (d_9_i0_) + (1)
                    d_10_v3_: str
                    d_10_v3_ = _dafny.CodePoint('r')
                    d_11_v4_: _dafny.Map
                    d_11_v4_ = _dafny.Map({d_6_v0_: d_10_v3_})
                    d_12_v5_: _dafny.Array
                    nw0_ = _dafny.Array(None, 15)
                    nw0_[int(0)] = d_10_v3_
                    nw0_[int(1)] = default__.fm1(globalState)
                    nw0_[int(2)] = d_10_v3_
                    nw0_[int(3)] = d_10_v3_
                    nw0_[int(4)] = d_10_v3_
                    nw0_[int(5)] = d_10_v3_
                    nw0_[int(6)] = d_10_v3_
                    nw0_[int(7)] = d_10_v3_
                    nw0_[int(8)] = (d_10_v3_ if d_6_v0_ else d_10_v3_)
                    nw0_[int(9)] = ((d_11_v4_)[d_6_v0_] if (d_6_v0_) in (d_11_v4_) else _dafny.CodePoint('b'))
                    nw0_[int(10)] = d_10_v3_
                    nw0_[int(11)] = d_10_v3_
                    nw0_[int(12)] = d_10_v3_
                    nw0_[int(13)] = d_10_v3_
                    nw0_[int(14)] = d_10_v3_
                    d_12_v5_ = nw0_
                    index0_ = default__.safeIndex(34, (d_12_v5_).length(0))
                    (d_12_v5_)[index0_] = d_10_v3_
                    index1_ = default__.safeIndex(34, (d_12_v5_).length(0))
                    (d_12_v5_)[index1_] = d_10_v3_
                    r1 = default__.fm0(d_6_v0_, p0, globalState)
                    d_13_v6_: C0
                    nw1_ = C0()
                    nw1_.ctor__()
                    d_13_v6_ = nw1_
                    d_14_v7_: _dafny.MultiSet
                    d_14_v7_ = _dafny.MultiSet([True, d_6_v0_, d_6_v0_])
                    d_6_v0_ = (((d_14_v7_) | (d_14_v7_)).cardinality) != (p0)
                    pass
            pass
        d_15_v8_: _dafny.Seq
        d_15_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhpgj"))
        hi0_ = default__.safeDivisionInt(380, 60)
        for d_16_i1_ in range(len(d_15_v8_), hi0_):
            d_17_v9_: _dafny.Array
            nw2_ = _dafny.Array(None, 1)
            nw2_[int(0)] = d_16_i1_
            d_17_v9_ = nw2_
            d_18_v10_: _dafny.Seq
            d_18_v10_ = _dafny.SeqWithoutIsStrInference([d_17_v9_, d_17_v9_, d_17_v9_])
            (globalState).f2 = len(d_18_v10_)
            d_15_v8_ = (d_15_v8_) + (d_15_v8_)
            d_19_v11_: str
            d_19_v11_ = _dafny.CodePoint('n')
            d_15_v8_ = (d_15_v8_) + ((default__.fm6(not(False), d_6_v0_, globalState)).set(default__.safeIndex(p1, len(default__.fm6(not(False), d_6_v0_, globalState))), d_19_v11_))
            if d_6_v0_:
                d_20_v12_: _dafny.Map
                d_20_v12_ = _dafny.Map({(830) * (d_16_i1_): d_15_v8_})
                d_20_v12_ = (d_20_v12_).set(p1, d_15_v8_)
                (globalState).f2 = p1
                d_21_v13_: C0
                nw3_ = C0()
                nw3_.ctor__()
                d_21_v13_ = nw3_
                d_22_v14_: _dafny.Array
                nw4_ = _dafny.Array(None, 24)
                nw4_[int(0)] = d_21_v13_
                nw4_[int(1)] = d_21_v13_
                nw4_[int(2)] = d_21_v13_
                nw4_[int(3)] = d_21_v13_
                nw4_[int(4)] = d_21_v13_
                nw4_[int(5)] = d_21_v13_
                nw4_[int(6)] = d_21_v13_
                nw4_[int(7)] = d_21_v13_
                nw4_[int(8)] = d_21_v13_
                nw4_[int(9)] = d_21_v13_
                nw4_[int(10)] = (d_21_v13_ if d_6_v0_ else d_21_v13_)
                nw4_[int(11)] = d_21_v13_
                nw4_[int(12)] = d_21_v13_
                nw4_[int(13)] = d_21_v13_
                nw4_[int(14)] = d_21_v13_
                nw4_[int(15)] = d_21_v13_
                nw4_[int(16)] = (d_21_v13_ if d_6_v0_ else d_21_v13_)
                nw4_[int(17)] = d_21_v13_
                nw4_[int(18)] = d_21_v13_
                nw4_[int(19)] = d_21_v13_
                nw4_[int(20)] = d_21_v13_
                nw4_[int(21)] = d_21_v13_
                nw4_[int(22)] = (d_21_v13_ if d_6_v0_ else d_21_v13_)
                nw4_[int(23)] = d_21_v13_
                d_22_v14_ = nw4_
                nw5_ = _dafny.Array(None, 27)
                d_22_v14_ = nw5_
                d_23_v15_: _dafny.MultiSet
                d_23_v15_ = _dafny.MultiSet([d_6_v0_, d_6_v0_])
                (globalState).f2 = ((d_23_v15_)[not((d_19_v11_) not in (d_15_v8_))] if (not((d_19_v11_) not in (d_15_v8_))) in (d_23_v15_) else d_16_i1_)
                d_24_v16_: _dafny.Seq
                d_24_v16_ = _dafny.SeqWithoutIsStrInference([p1, default__.safeDivisionInt(p0, p0)])
                d_24_v16_ = default__.fm7(d_6_v0_, (569 if d_6_v0_ else d_16_i1_), globalState)
            elif True:
                d_25_v17_: _dafny.Array
                def lambda0_(d_26_v8_):
                    def lambda1_(d_27_i2_):
                        return d_26_v8_

                    return lambda1_

                init0_ = lambda0_(d_15_v8_)
                nw6_ = _dafny.Array(None, 16)
                for i0_0_ in range(nw6_.length(0)):
                    nw6_[i0_0_] = init0_(i0_0_)
                d_25_v17_ = nw6_
                d_28_v18_: D0
                d_28_v18_ = D0_DC2(d_6_v0_, d_25_v17_, p0)
                d_29_v19_: _dafny.Map
                d_29_v19_ = _dafny.Map({p0: (d_28_v18_).cf2})
                d_30_v20_: _dafny.MultiSet
                d_30_v20_ = _dafny.MultiSet([d_16_i1_, d_16_i1_, p1])
                d_6_v0_ = ((d_29_v19_)[default__.safeModuloInt(d_16_i1_, len((d_15_v8_).set(default__.safeIndex(((d_30_v20_)[len(d_15_v8_)] if (len(d_15_v8_)) in (d_30_v20_) else p0), len(d_15_v8_)), d_19_v11_)))] if (default__.safeModuloInt(d_16_i1_, len((d_15_v8_).set(default__.safeIndex(((d_30_v20_)[len(d_15_v8_)] if (len(d_15_v8_)) in (d_30_v20_) else p0), len(d_15_v8_)), d_19_v11_)))) in (d_29_v19_) else d_6_v0_)
                d_31_v21_: _dafny.Seq
                d_31_v21_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_6_v0_, p0, globalState), p1])
                d_32_v22_: _dafny.Set
                d_32_v22_ = _dafny.Set({d_19_v11_, d_19_v11_})
                d_33_v23_: _dafny.Seq
                d_33_v23_ = _dafny.SeqWithoutIsStrInference([d_31_v21_])
                d_34_v24_: D0
                d_34_v24_ = D0_DC1(p0, (d_31_v21_).set(default__.safeIndex(len(d_32_v22_), len(d_31_v21_)), len(d_33_v23_)))
                (globalState).f2 = (d_34_v24_).cf0
                d_6_v0_ = d_6_v0_
                d_35_v25_: D0
                d_35_v25_ = D0_DC0()
                rhs0_ = d_16_i1_
                rhs1_ = d_35_v25_
                lhs0_ = globalState
                lhs0_.f2 = rhs0_
                d_35_v25_ = rhs1_
                d_36_v26_: _dafny.MultiSet
                d_36_v26_ = _dafny.MultiSet([d_6_v0_])
                d_6_v0_ = ((default__.fm8(globalState)).intersection(d_36_v26_)).isdisjoint((d_36_v26_) | (_dafny.MultiSet([d_6_v0_])))
        d_37_v27_: _dafny.Array
        nw7_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
        d_37_v27_ = nw7_
        d_38_v28_: _dafny.Seq
        d_38_v28_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_39_v29_: D0
        d_39_v29_ = D0_DC2(d_6_v0_, d_37_v27_, (D0_DC1(p1, d_38_v28_)).cf0)
        d_40_i3_: int
        d_40_i3_ = 0
        with _dafny.label("1"):
            while not((d_39_v29_).cf2):
                with _dafny.c_label("1"):
                    if (d_40_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_40_i3_ = (d_40_i3_) + (1)
                    d_41_v30_: _dafny.Array
                    nw8_ = _dafny.Array(None, 1)
                    d_41_v30_ = nw8_
                    d_42_v31_: C0
                    nw9_ = C0()
                    nw9_.ctor__()
                    d_42_v31_ = nw9_
                    index2_ = default__.safeIndex(252, (d_41_v30_).length(0))
                    (d_41_v30_)[index2_] = d_42_v31_
                    index3_ = default__.safeIndex(252, (d_41_v30_).length(0))
                    (d_41_v30_)[index3_] = d_42_v31_
                    d_43_v32_: D0
                    d_43_v32_ = D0_DC1(p1, d_38_v28_)
                    (globalState).f2 = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_6_v0_, d_6_v0_])), (p1) - ((d_43_v32_).cf0))
                    d_38_v28_ = d_38_v28_
                    d_44_v33_: _dafny.Map
                    d_44_v33_ = _dafny.Map({d_38_v28_: (d_15_v8_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odihkxdux")))})
                    r1 = len(d_44_v33_)
                    pass
            pass
        d_45_v35_: _dafny.Array
        def lambda2_(d_46_v8_):
            def lambda3_(d_47_i5_):
                def iife2_():
                    coll2_ = _dafny.Set()
                    compr_2_: str
                    for compr_2_ in (_dafny.MultiSet(d_46_v8_)).Elements:
                        d_48_v34_: str = compr_2_
                        if (d_48_v34_) in (_dafny.MultiSet(d_46_v8_)):
                            coll2_ = coll2_.union(_dafny.Set([d_48_v34_]))
                    return _dafny.Set(coll2_)
                return (_dafny.Set({_dafny.CodePoint('b')})) - (iife2_()
                )

            return lambda3_

        init1_ = lambda2_(d_15_v8_)
        nw10_ = _dafny.Array(None, 29)
        for i0_1_ in range(nw10_.length(0)):
            nw10_[i0_1_] = init1_(i0_1_)
        d_45_v35_ = nw10_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_45_v35_).length(0)):
            d_49_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_49_i4_)) and ((d_49_i4_) < ((d_45_v35_).length(0)))):
                (d_45_v35_)[(d_49_i4_)] = (_dafny.Set({_dafny.CodePoint('b')})) - (_dafny.Set({_dafny.CodePoint('f')}))
        d_50_v36_: _dafny.Array
        nw11_ = _dafny.Array(False, 19)
        d_50_v36_ = nw11_
        r0 = d_50_v36_
        r1 = p1
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_51_v0_: _dafny.Seq
        d_51_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhvwcv"))
        d_52_globalState_: GlobalState
        nw12_ = GlobalState()
        nw12_.ctor__((d_51_v0_) + (d_51_v0_), True, 411, d_51_v0_)
        d_52_globalState_ = nw12_
        d_53_v1_: int
        d_53_v1_ = 387
        (d_52_globalState_).f2 = d_53_v1_
        d_54_v2_: bool
        d_54_v2_ = False
        d_54_v2_ = d_54_v2_
        if d_54_v2_:
            d_55_v3_: _dafny.Array
            d_56_v4_: int
            out0_: _dafny.Array
            out1_: int
            out0_, out1_ = default__.m0(d_53_v1_, (d_53_v1_) - (d_53_v1_), d_52_globalState_)
            d_55_v3_ = out0_
            d_56_v4_ = out1_
            d_57_v5_: _dafny.MultiSet
            d_57_v5_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-8, d_56_v4_])])
            d_58_v6_: _dafny.Seq
            d_58_v6_ = _dafny.SeqWithoutIsStrInference([d_57_v5_, d_57_v5_])
            d_54_v2_ = ((d_58_v6_)[default__.safeIndex((0) - (d_56_v4_), len(d_58_v6_))]).ispropersubset(d_57_v5_)
            d_59_v7_: _dafny.Array
            def lambda4_(d_60_v2_, d_61_v0_):
                def lambda5_(d_62_i0_):
                    return _dafny.Map({d_60_v2_: len(d_61_v0_)})

                return lambda5_

            init2_ = lambda4_(d_54_v2_, d_51_v0_)
            nw13_ = _dafny.Array(None, 16)
            for i0_2_ in range(nw13_.length(0)):
                nw13_[i0_2_] = init2_(i0_2_)
            d_59_v7_ = nw13_
            index4_ = default__.safeIndex(98, (d_55_v3_).length(0))
            (d_55_v3_)[index4_] = d_54_v2_
            d_63_v8_: _dafny.Map
            d_63_v8_ = _dafny.Map({d_53_v1_: default__.safeModuloInt(len(d_51_v0_), d_56_v4_)})
            d_64_v9_: _dafny.Set
            d_64_v9_ = _dafny.Set({d_53_v1_, d_56_v4_, d_53_v1_})
            index5_ = default__.safeIndex(98, (d_55_v3_).length(0))
            rhs2_ = len(d_63_v8_)
            rhs3_ = d_59_v7_
            rhs4_ = (d_64_v9_).isdisjoint((d_64_v9_).intersection(d_64_v9_))
            lhs1_ = d_52_globalState_
            lhs2_ = d_55_v3_
            lhs3_ = default__.safeIndex(98, (d_55_v3_).length(0))
            lhs1_.f2 = rhs2_
            d_59_v7_ = rhs3_
            lhs2_[lhs3_] = rhs4_
            d_65_v10_: _dafny.Array
            d_66_v11_: int
            out2_: _dafny.Array
            out3_: int
            out2_, out3_ = default__.m0(default__.safeDivisionInt(d_53_v1_, d_56_v4_), d_53_v1_, d_52_globalState_)
            d_65_v10_ = out2_
            d_66_v11_ = out3_
            d_51_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwstnawx"))
        elif True:
            d_67_v12_: _dafny.Array
            nw14_ = _dafny.Array(int(0), 4)
            d_67_v12_ = nw14_
            d_68_v13_: _dafny.Map
            d_68_v13_ = _dafny.Map({d_53_v1_: d_67_v12_})
            d_69_v14_: _dafny.Array
            nw15_ = _dafny.Array(None, 1)
            nw15_[int(0)] = ((d_68_v13_)[d_53_v1_] if (d_53_v1_) in (d_68_v13_) else d_67_v12_)
            d_69_v14_ = nw15_
            index6_ = default__.safeIndex(281, (d_69_v14_).length(0))
            (d_69_v14_)[index6_] = d_67_v12_
            index7_ = default__.safeIndex(281, (d_69_v14_).length(0))
            (d_69_v14_)[index7_] = d_67_v12_
            (d_52_globalState_).f2 = d_53_v1_
            d_70_v15_: _dafny.Array
            d_71_v16_: int
            out4_: _dafny.Array
            out5_: int
            out4_, out5_ = default__.m0(d_53_v1_, (0) - (-856), d_52_globalState_)
            d_70_v15_ = out4_
            d_71_v16_ = out5_
            index8_ = default__.safeIndex(281, (d_69_v14_).length(0))
            (d_69_v14_)[index8_] = d_67_v12_
            d_72_v17_: _dafny.Map
            d_72_v17_ = _dafny.Map({d_70_v15_: (d_71_v16_) >= (d_71_v16_)})
            d_72_v17_ = (d_72_v17_) | (_dafny.Map({d_70_v15_: False}))
        hi1_ = d_53_v1_
        for d_73_i1_ in range(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suibphoto"))) + (d_51_v0_)), hi1_):
            d_74_v18_: _dafny.Array
            def lambda6_(d_75_v1_, d_76_i1_):
                def lambda7_(d_77_i2_):
                    return (d_75_v1_) != (d_76_i1_)

                return lambda7_

            init3_ = lambda6_(d_53_v1_, d_73_i1_)
            nw16_ = _dafny.Array(None, 12)
            for i0_3_ in range(nw16_.length(0)):
                nw16_[i0_3_] = init3_(i0_3_)
            d_74_v18_ = nw16_
            d_74_v18_ = d_74_v18_
            d_53_v1_ = d_73_i1_
            d_78_v19_: _dafny.Seq
            d_78_v19_ = _dafny.SeqWithoutIsStrInference([False, d_54_v2_])
            d_79_v20_: _dafny.Array
            d_80_v21_: int
            out6_: _dafny.Array
            out7_: int
            out6_, out7_ = default__.m0((663) * (default__.fm0(d_54_v2_, len(d_78_v19_), d_52_globalState_)), d_73_i1_, d_52_globalState_)
            d_79_v20_ = out6_
            d_80_v21_ = out7_
            rhs5_ = d_54_v2_
            rhs6_ = -337
            lhs4_ = d_52_globalState_
            d_54_v2_ = rhs5_
            lhs4_.f2 = rhs6_
        d_81_v22_: _dafny.Set
        d_81_v22_ = _dafny.Set({d_54_v2_, d_54_v2_, True})
        d_82_v23_: _dafny.Array
        nw17_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
        d_82_v23_ = nw17_
        d_83_v24_: D0
        d_83_v24_ = D0_DC2((d_81_v22_) != (d_81_v22_), d_82_v23_, d_53_v1_)
        source1_ = d_83_v24_
        if source1_.is_DC0:
            d_84_v25_: _dafny.Map
            d_84_v25_ = _dafny.Map({d_53_v1_: d_54_v2_})
            d_85_v26_: _dafny.Seq
            d_85_v26_ = _dafny.SeqWithoutIsStrInference([d_54_v2_])
            d_86_v27_: _dafny.Map
            d_86_v27_ = _dafny.Map({d_85_v26_: True})
            d_87_v28_: _dafny.Array
            nw18_ = _dafny.Array(None, 12)
            nw18_[int(0)] = d_53_v1_
            nw18_[int(1)] = default__.fm0(d_54_v2_, d_53_v1_, d_52_globalState_)
            nw18_[int(2)] = -526
            nw18_[int(3)] = -912
            nw18_[int(4)] = d_53_v1_
            nw18_[int(5)] = d_53_v1_
            nw18_[int(6)] = (0) - (-894)
            nw18_[int(7)] = d_53_v1_
            nw18_[int(8)] = (d_53_v1_) - (len(d_84_v25_))
            nw18_[int(9)] = len(d_86_v27_)
            nw18_[int(10)] = (d_53_v1_ if True else d_53_v1_)
            nw18_[int(11)] = d_53_v1_
            d_87_v28_ = nw18_
            d_88_v29_: _dafny.Array
            def lambda8_(d_89_i3_):
                return _dafny.CodePoint('d')

            init4_ = lambda8_
            nw19_ = _dafny.Array(None, 5)
            for i0_4_ in range(nw19_.length(0)):
                nw19_[i0_4_] = init4_(i0_4_)
            d_88_v29_ = nw19_
            d_90_v30_: _dafny.Map
            d_90_v30_ = _dafny.Map({d_88_v29_: d_53_v1_})
            d_91_v31_: _dafny.Map
            d_91_v31_ = _dafny.Map({d_53_v1_: len(d_90_v30_)})
            d_92_v32_: _dafny.Map
            d_92_v32_ = _dafny.Map({True: d_91_v31_})
            index9_ = default__.safeIndex(377, (d_87_v28_).length(0))
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(955, 328):
                    d_93_v33_: int = compr_3_
                    if ((955) <= (d_93_v33_)) and ((d_93_v33_) < (328)):
                        coll3_[(d_93_v33_) * (d_53_v1_)] = d_53_v1_
                return _dafny.Map(coll3_)
            (d_87_v28_)[index9_] = len(((d_92_v32_)[d_54_v2_] if (d_54_v2_) in (d_92_v32_) else iife3_()
            ))
            index10_ = default__.safeIndex(377, (d_87_v28_).length(0))
            (d_87_v28_)[index10_] = default__.fm0(d_54_v2_, 348, d_52_globalState_)
            d_94_v34_: str
            d_94_v34_ = _dafny.CodePoint('o')
            d_95_v35_: _dafny.Array
            nw20_ = _dafny.Array(False, 4)
            d_95_v35_ = nw20_
            index11_ = default__.safeIndex(709, (d_95_v35_).length(0))
            (d_95_v35_)[index11_] = d_54_v2_
            d_96_v36_: _dafny.Map
            d_96_v36_ = _dafny.Map({(d_83_v24_ if d_54_v2_ else D0_DC2(False, d_82_v23_, d_53_v1_)): d_94_v34_})
            index12_ = default__.safeIndex(377, (d_87_v28_).length(0))
            index13_ = default__.safeIndex(709, (d_95_v35_).length(0))
            rhs7_ = d_54_v2_
            rhs8_ = default__.safeModuloInt((d_87_v28_)[default__.safeIndex(377, (d_87_v28_).length(0))], (d_87_v28_)[default__.safeIndex(377, (d_87_v28_).length(0))])
            rhs9_ = ((d_96_v36_)[(D0_DC2(d_54_v2_, d_82_v23_, (d_87_v28_)[default__.safeIndex(377, (d_87_v28_).length(0))]) if False else d_83_v24_)] if ((D0_DC2(d_54_v2_, d_82_v23_, (d_87_v28_)[default__.safeIndex(377, (d_87_v28_).length(0))]) if False else d_83_v24_)) in (d_96_v36_) else default__.fm1(d_52_globalState_))
            rhs10_ = not ((d_54_v2_ if not(d_54_v2_) else d_54_v2_)) or (d_54_v2_)
            lhs5_ = d_87_v28_
            lhs6_ = default__.safeIndex(377, (d_87_v28_).length(0))
            lhs7_ = d_95_v35_
            lhs8_ = default__.safeIndex(709, (d_95_v35_).length(0))
            d_54_v2_ = rhs7_
            lhs5_[lhs6_] = rhs8_
            d_94_v34_ = rhs9_
            lhs7_[lhs8_] = rhs10_
            d_97_v37_: C0
            nw21_ = C0()
            nw21_.ctor__()
            d_97_v37_ = nw21_
            d_98_v38_: _dafny.Seq
            d_98_v38_ = _dafny.SeqWithoutIsStrInference([-296, d_53_v1_])
            rhs11_ = d_51_v0_
            rhs12_ = (d_98_v38_)[default__.safeIndex(default__.safeDivisionInt((d_87_v28_)[default__.safeIndex(377, (d_87_v28_).length(0))], d_53_v1_), len(d_98_v38_))]
            lhs9_ = d_52_globalState_
            d_51_v0_ = rhs11_
            lhs9_.f2 = rhs12_
        elif source1_.is_DC1:
            d_99___mcc_h0_ = source1_.cf0
            d_100___mcc_h1_ = source1_.cf1
            d_101_cf1_ = d_100___mcc_h1_
            d_102_cf0_ = d_99___mcc_h0_
            d_103_v39_: _dafny.Array
            nw22_ = _dafny.Array(int(0), 18)
            d_103_v39_ = nw22_
            index14_ = default__.safeIndex(335, (d_103_v39_).length(0))
            (d_103_v39_)[index14_] = d_53_v1_
            index15_ = default__.safeIndex(335, (d_103_v39_).length(0))
            (d_103_v39_)[index15_] = (d_102_cf0_) - (d_102_cf0_)
            d_54_v2_ = d_54_v2_
            d_103_v39_ = d_103_v39_
            d_54_v2_ = d_54_v2_
        elif True:
            d_104___mcc_h2_ = source1_.cf2
            d_105___mcc_h3_ = source1_.cf3
            d_106___mcc_h4_ = source1_.cf4
            d_107_cf4_ = d_106___mcc_h4_
            d_108_cf3_ = d_105___mcc_h3_
            d_109_cf2_ = d_104___mcc_h2_
            d_51_v0_ = d_51_v0_
            d_109_cf2_ = (d_83_v24_).cf2
            (d_52_globalState_).f2 = (d_53_v1_) + ((d_53_v1_) - (d_53_v1_))
            d_110_v41_: C0
            nw23_ = C0()
            nw23_.ctor__()
            d_110_v41_ = nw23_
            d_111_v42_: _dafny.Map
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(927, -180):
                    d_112_v40_: int = compr_4_
                    if ((927) <= (d_112_v40_)) and ((d_112_v40_) < (-180)):
                        coll4_[default__.safeDivisionInt(d_112_v40_, d_53_v1_)] = d_54_v2_
                return _dafny.Map(coll4_)
            d_111_v42_ = _dafny.Map({iife4_()
            : d_110_v41_})
            d_113_v43_: _dafny.Array
            d_114_v44_: int
            out8_: _dafny.Array
            out9_: int
            out8_, out9_ = default__.m0(len(d_111_v42_), (0) - ((d_53_v1_ if d_109_cf2_ else 506)), d_52_globalState_)
            d_113_v43_ = out8_
            d_114_v44_ = out9_
        d_115_v45_: C0
        nw24_ = C0()
        nw24_.ctor__()
        d_115_v45_ = nw24_
        d_116_v46_: _dafny.MultiSet
        d_116_v46_ = _dafny.MultiSet([d_115_v45_, d_115_v45_])
        d_117_v47_: _dafny.Seq
        d_117_v47_ = _dafny.SeqWithoutIsStrInference([d_53_v1_, d_53_v1_, d_53_v1_, (d_116_v46_).cardinality])
        d_118_v48_: _dafny.MultiSet
        d_118_v48_ = _dafny.MultiSet([-254, (d_117_v47_)[default__.safeIndex(len(d_51_v0_), len(d_117_v47_))], d_53_v1_])
        hi2_ = ((d_118_v48_)[-779] if (-779) in (d_118_v48_) else d_53_v1_)
        for d_119_i4_ in range(d_53_v1_, hi2_):
            d_120_v49_: _dafny.Seq
            d_120_v49_ = _dafny.SeqWithoutIsStrInference([d_54_v2_])
            d_121_v50_: _dafny.Map
            d_121_v50_ = _dafny.Map({d_120_v49_: d_54_v2_})
            d_122_v51_: str
            d_122_v51_ = _dafny.CodePoint('s')
            d_51_v0_ = (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_123_i5_ in range(default__.abs(218))])) + (d_51_v0_)).set(default__.safeIndex((len(d_121_v50_)) + (d_119_i4_), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_124_i5_ in range(default__.abs(218))])) + (d_51_v0_))), d_122_v51_)).set(default__.safeIndex((0) - (d_53_v1_), len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_125_i5_ in range(default__.abs(218))])) + (d_51_v0_)).set(default__.safeIndex((len(d_121_v50_)) + (d_119_i4_), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_126_i5_ in range(default__.abs(218))])) + (d_51_v0_))), d_122_v51_))), _dafny.CodePoint('l'))
            d_127_v52_: _dafny.Map
            d_127_v52_ = _dafny.Map({(d_53_v1_) + (d_119_i4_): 807})
            d_127_v52_ = (d_127_v52_).set((len(_dafny.Set({d_119_i4_, d_119_i4_, d_53_v1_}))) + (d_53_v1_), default__.safeDivisionInt(d_119_i4_, d_53_v1_))
            d_128_v53_: _dafny.Array
            nw25_ = _dafny.Array(int(0), 20)
            d_128_v53_ = nw25_
            d_129_v54_: _dafny.Set
            d_129_v54_ = _dafny.Set({d_128_v53_, d_128_v53_})
            d_54_v2_ = (d_129_v54_).issubset(d_129_v54_)
            d_54_v2_ = (d_120_v49_)[default__.safeIndex(d_53_v1_, len(d_120_v49_))]
        d_130_v55_: _dafny.Array
        nw26_ = _dafny.Array(_dafny.Array(None, 0), 2)
        d_130_v55_ = nw26_
        d_130_v55_ = d_130_v55_
        d_131_v56_: _dafny.Seq
        d_131_v56_ = _dafny.SeqWithoutIsStrInference([d_54_v2_, d_54_v2_])
        if ((d_53_v1_) * (len(d_131_v56_))) in (d_117_v47_):
            d_53_v1_ = (d_53_v1_ if False else (d_53_v1_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_132_i6_ in range(default__.abs(561))]))))
            d_133_v57_: _dafny.Map
            d_133_v57_ = _dafny.Map({default__.fm1(d_52_globalState_): d_54_v2_})
            d_134_v58_: str
            d_134_v58_ = _dafny.CodePoint('a')
            d_54_v2_ = not(((d_133_v57_)[d_134_v58_] if (d_134_v58_) in (d_133_v57_) else not((d_51_v0_) == (d_51_v0_))))
            (d_52_globalState_).f2 = d_53_v1_
            d_83_v24_ = d_83_v24_
            d_135_v59_: _dafny.Map
            d_135_v59_ = _dafny.Map({993: d_53_v1_})
            d_54_v2_ = (774) <= (default__.safeDivisionInt((0) - ((0) - (len(d_51_v0_))), len(d_135_v59_)))
        elif True:
            d_136_v60_: C0
            nw27_ = C0()
            nw27_.ctor__()
            d_136_v60_ = nw27_
            d_54_v2_ = (len(d_117_v47_)) != (d_53_v1_)
            d_137_v61_: _dafny.Array
            d_138_v62_: int
            out10_: _dafny.Array
            out11_: int
            out10_, out11_ = default__.m0(d_53_v1_, d_53_v1_, d_52_globalState_)
            d_137_v61_ = out10_
            d_138_v62_ = out11_
            if d_54_v2_:
                d_139_v63_: _dafny.Array
                nw28_ = _dafny.Array(None, 25)
                nw28_[int(0)] = default__.fm0(not(d_54_v2_), d_53_v1_, d_52_globalState_)
                nw28_[int(1)] = d_53_v1_
                nw28_[int(2)] = d_53_v1_
                nw28_[int(3)] = d_53_v1_
                nw28_[int(4)] = (d_138_v62_) + (d_53_v1_)
                nw28_[int(5)] = (d_138_v62_) - (d_53_v1_)
                nw28_[int(6)] = d_53_v1_
                nw28_[int(7)] = d_53_v1_
                nw28_[int(8)] = (10) + (d_53_v1_)
                nw28_[int(9)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fatyy")))
                nw28_[int(10)] = d_138_v62_
                nw28_[int(11)] = d_53_v1_
                nw28_[int(12)] = d_138_v62_
                nw28_[int(13)] = d_53_v1_
                nw28_[int(14)] = d_53_v1_
                nw28_[int(15)] = d_53_v1_
                nw28_[int(16)] = (d_138_v62_) * ((0) - ((0) - (default__.fm0(d_54_v2_, len(d_81_v22_), d_52_globalState_))))
                nw28_[int(17)] = (len(d_131_v56_)) * (7)
                nw28_[int(18)] = (0) - (d_138_v62_)
                nw28_[int(19)] = (0) - (d_138_v62_)
                nw28_[int(20)] = default__.safeModuloInt(d_138_v62_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iasqrm"))))
                nw28_[int(21)] = (0) - (d_138_v62_)
                nw28_[int(22)] = d_53_v1_
                nw28_[int(23)] = (d_117_v47_)[default__.safeIndex(d_53_v1_, len(d_117_v47_))]
                nw28_[int(24)] = d_138_v62_
                d_139_v63_ = nw28_
                d_139_v63_ = d_139_v63_
                d_54_v2_ = not((d_118_v48_).issubset((default__.fm3(_dafny.SeqWithoutIsStrInference([d_53_v1_]), d_54_v2_, d_54_v2_, d_54_v2_, d_52_globalState_)) - (d_118_v48_)))
                index16_ = default__.safeIndex(57, (d_139_v63_).length(0))
                (d_139_v63_)[index16_] = d_53_v1_
                index17_ = default__.safeIndex(57, (d_139_v63_).length(0))
                (d_139_v63_)[index17_] = len(_dafny.Set({(d_54_v2_ if d_54_v2_ else True), ((d_118_v48_).set(d_138_v62_, default__.abs(d_53_v1_))).issubset(d_118_v48_), (d_51_v0_) <= (d_51_v0_)}))
                d_140_v64_: _dafny.Map
                d_140_v64_ = _dafny.Map({d_54_v2_: default__.fm4(d_54_v2_, d_52_globalState_)})
                d_141_v65_: _dafny.Map
                d_141_v65_ = _dafny.Map({default__.fm0(d_54_v2_, d_138_v62_, d_52_globalState_): d_54_v2_})
                d_54_v2_ = ((d_140_v64_)[(len(d_141_v65_)) >= ((d_139_v63_)[default__.safeIndex(57, (d_139_v63_).length(0))])] if ((len(d_141_v65_)) >= ((d_139_v63_)[default__.safeIndex(57, (d_139_v63_).length(0))])) in (d_140_v64_) else not(d_54_v2_))
                d_142_v66_: C0
                nw29_ = C0()
                nw29_.ctor__()
                d_142_v66_ = nw29_
            elif True:
                d_143_v67_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.Seq({}), 20)
                d_143_v67_ = nw30_
                d_144_v68_: _dafny.Array
                def lambda9_(d_145_v1_):
                    def lambda10_(d_146_i7_):
                        return (d_146_i7_) - (d_145_v1_)

                    return lambda10_

                init5_ = lambda9_(d_53_v1_)
                nw31_ = _dafny.Array(None, 8)
                for i0_5_ in range(nw31_.length(0)):
                    nw31_[i0_5_] = init5_(i0_5_)
                d_144_v68_ = nw31_
                index18_ = default__.safeIndex(580, (d_144_v68_).length(0))
                (d_144_v68_)[index18_] = d_53_v1_
                d_147_v69_: _dafny.Map
                d_147_v69_ = _dafny.Map({len(d_51_v0_): len(d_81_v22_)})
                d_148_v70_: _dafny.Map
                d_148_v70_ = _dafny.Map({d_137_v61_: d_147_v69_})
                d_149_v71_: _dafny.Map
                d_149_v71_ = _dafny.Map({d_144_v68_: d_143_v67_})
                index19_ = default__.safeIndex(580, (d_144_v68_).length(0))
                rhs13_ = (d_148_v70_) != (d_148_v70_)
                rhs14_ = d_54_v2_
                rhs15_ = ((d_149_v71_)[d_144_v68_] if (d_144_v68_) in (d_149_v71_) else d_143_v67_)
                rhs16_ = (0) - (d_138_v62_)
                lhs10_ = d_144_v68_
                lhs11_ = default__.safeIndex(580, (d_144_v68_).length(0))
                d_54_v2_ = rhs13_
                d_54_v2_ = rhs14_
                d_143_v67_ = rhs15_
                lhs10_[lhs11_] = rhs16_
                d_150_v72_: _dafny.Set
                d_150_v72_ = _dafny.Set({((d_131_v56_).set(default__.safeIndex(d_53_v1_, len(d_131_v56_)), d_54_v2_)) + (d_131_v56_), (d_131_v56_) + (d_131_v56_)})
                d_150_v72_ = _dafny.Set({d_131_v56_, (default__.fm5((d_144_v68_)[default__.safeIndex(580, (d_144_v68_).length(0))], d_54_v2_, d_54_v2_, d_52_globalState_)) + (_dafny.SeqWithoutIsStrInference([d_54_v2_])), d_131_v56_})
                d_151_v73_: str
                d_151_v73_ = _dafny.CodePoint('a')
                d_152_v74_: _dafny.Map
                d_152_v74_ = _dafny.Map({d_118_v48_: d_151_v73_})
                d_152_v74_ = (d_152_v74_).set(d_118_v48_, (d_151_v73_ if False else d_151_v73_))
                d_153_v75_: _dafny.Array
                nw32_ = _dafny.Array(D0.default()(), 24)
                d_153_v75_ = nw32_
                index20_ = default__.safeIndex(483, (d_153_v75_).length(0))
                (d_153_v75_)[index20_] = d_83_v24_
                pat_let_tv0_ = d_82_v23_
                index21_ = default__.safeIndex(483, (d_153_v75_).length(0))
                def iife5_(_pat_let0_0):
                    def iife6_(d_154_dt__update__tmp_h0_):
                        def iife7_(_pat_let1_0):
                            def iife8_(d_155_dt__update_hcf3_h0_):
                                return D0_DC2((d_154_dt__update__tmp_h0_).cf2, d_155_dt__update_hcf3_h0_, (d_154_dt__update__tmp_h0_).cf4)
                            return iife8_(_pat_let1_0)
                        return iife7_(pat_let_tv0_)
                    return iife6_(_pat_let0_0)
                (d_153_v75_)[index21_] = iife5_(d_83_v24_)
                d_156_v76_: _dafny.Array
                def lambda11_(d_157_i8_):
                    return _dafny.CodePoint('d')

                init6_ = lambda11_
                nw33_ = _dafny.Array(None, 24)
                for i0_6_ in range(nw33_.length(0)):
                    nw33_[i0_6_] = init6_(i0_6_)
                d_156_v76_ = nw33_
                index22_ = default__.safeIndex(942, (d_156_v76_).length(0))
                (d_156_v76_)[index22_] = d_151_v73_
                index23_ = default__.safeIndex(942, (d_156_v76_).length(0))
                rhs17_ = d_151_v73_
                rhs18_ = (d_118_v48_).set((d_144_v68_)[default__.safeIndex(580, (d_144_v68_).length(0))], default__.abs((d_144_v68_)[default__.safeIndex(580, (d_144_v68_).length(0))]))
                rhs19_ = (d_144_v68_)[default__.safeIndex(580, (d_144_v68_).length(0))]
                lhs12_ = d_156_v76_
                lhs13_ = default__.safeIndex(942, (d_156_v76_).length(0))
                lhs12_[lhs13_] = rhs17_
                d_118_v48_ = rhs18_
                d_138_v62_ = rhs19_
            d_158_v77_: _dafny.MultiSet
            d_158_v77_ = _dafny.MultiSet([not(d_54_v2_)])
            d_158_v77_ = d_158_v77_
        d_115_v45_ = d_115_v45_
        d_159_v78_: str
        d_159_v78_ = _dafny.CodePoint('i')
        rhs20_ = d_54_v2_
        rhs21_ = d_54_v2_
        rhs22_ = d_159_v78_
        rhs23_ = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_53_v1_, d_53_v1_]))).intersection((d_118_v48_ if d_54_v2_ else d_118_v48_))).cardinality
        d_54_v2_ = rhs20_
        d_54_v2_ = rhs21_
        d_159_v78_ = rhs22_
        d_53_v1_ = rhs23_
        d_160_v79_: D0
        d_160_v79_ = D0_DC0()
        d_161_v80_: _dafny.Array
        nw34_ = _dafny.Array(None, 1)
        nw34_[int(0)] = len(_dafny.Map({d_53_v1_: d_160_v79_}))
        d_161_v80_ = nw34_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_161_v80_).length(0)):
            d_162_i9_: int = guard_loop_1_
            if (True) and (((0) <= (d_162_i9_)) and ((d_162_i9_) < ((d_161_v80_).length(0)))):
                (d_161_v80_)[(d_162_i9_)] = default__.safeModuloInt(d_162_i9_, d_53_v1_)
        d_163_v81_: C0
        nw35_ = C0()
        nw35_.ctor__()
        d_163_v81_ = nw35_
        d_54_v2_ = d_54_v2_
        d_164_v82_: _dafny.Map
        d_164_v82_ = _dafny.Map({d_54_v2_: d_53_v1_})
        if ((d_164_v82_) != (d_164_v82_) if d_54_v2_ else not(d_54_v2_)):
            d_165_v83_: _dafny.Array
            d_166_v84_: int
            out12_: _dafny.Array
            out13_: int
            out12_, out13_ = default__.m0(d_53_v1_, d_53_v1_, d_52_globalState_)
            d_165_v83_ = out12_
            d_166_v84_ = out13_
            d_54_v2_ = not((_dafny.SeqWithoutIsStrInference([d_159_v78_ for d_167_i10_ in range(default__.abs(625))])) != (d_51_v0_))
            d_168_v85_: _dafny.Array
            d_169_v86_: int
            out14_: _dafny.Array
            out15_: int
            out14_, out15_ = default__.m0(default__.safeDivisionInt((d_118_v48_).cardinality, d_53_v1_), d_53_v1_, d_52_globalState_)
            d_168_v85_ = out14_
            d_169_v86_ = out15_
            d_54_v2_ = (d_159_v78_) not in ((d_51_v0_) + (d_51_v0_))
            d_170_v87_: C0
            nw36_ = C0()
            nw36_.ctor__()
            d_170_v87_ = nw36_
        elif True:
            d_53_v1_ = d_53_v1_
            rhs24_ = (0) - (d_53_v1_)
            rhs25_ = d_53_v1_
            rhs26_ = d_164_v82_
            lhs14_ = d_52_globalState_
            lhs14_.f2 = rhs24_
            d_53_v1_ = rhs25_
            d_164_v82_ = rhs26_
            d_171_v88_: _dafny.Array
            nw37_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_171_v88_ = nw37_
            d_172_v89_: _dafny.Array
            def lambda12_(d_173_v78_):
                def lambda13_(d_174_i11_):
                    return d_173_v78_

                return lambda13_

            init7_ = lambda12_(d_159_v78_)
            nw38_ = _dafny.Array(None, 7)
            for i0_7_ in range(nw38_.length(0)):
                nw38_[i0_7_] = init7_(i0_7_)
            d_172_v89_ = nw38_
            index24_ = default__.safeIndex(65, (d_171_v88_).length(0))
            (d_171_v88_)[index24_] = d_172_v89_
            index25_ = default__.safeIndex(65, (d_171_v88_).length(0))
            (d_171_v88_)[index25_] = d_172_v89_
            d_175_v90_: _dafny.Array
            d_176_v91_: int
            out16_: _dafny.Array
            out17_: int
            out16_, out17_ = default__.m0(len(d_51_v0_), (d_53_v1_) + (d_53_v1_), d_52_globalState_)
            d_175_v90_ = out16_
            d_176_v91_ = out17_
            d_51_v0_ = d_51_v0_
        d_177_v92_: _dafny.Array
        nw39_ = _dafny.Array(False, 9)
        d_177_v92_ = nw39_
        d_177_v92_ = d_177_v92_
        d_178_v93_: _dafny.Seq
        d_178_v93_ = _dafny.SeqWithoutIsStrInference([d_161_v80_])
        d_179_v94_: _dafny.Map
        d_179_v94_ = _dafny.Map({d_117_v47_: (d_178_v93_) + (d_178_v93_)})
        d_179_v94_ = (d_179_v94_).set(d_117_v47_, d_178_v93_)
        _dafny.print((d_51_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_52_globalState_).f0).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_52_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_52_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_52_globalState_).f3).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_53_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_54_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_81_v22_) == (_dafny.Set({False, True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v24_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v24_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_v46_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_117_v47_) == (_dafny.SeqWithoutIsStrInference([386, 386, 386, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v48_) == (_dafny.MultiSet([-254, 386, 386]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v56_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_159_v78_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v80_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v82_) == (_dafny.Map({False: 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_178_v93_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_179_v94_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC0(D0, NamedTuple('DC0', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf0', Any), ('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({self.cf5.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self._f0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f1: bool = False
        self._f3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f3(self):
        return self._f3

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm2(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ieh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewuihk")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))

