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
        return ((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality for d_0_i0_ in range(default__.abs(573))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(not(True)): True}))]))) + ((_dafny.SeqWithoutIsStrInference([-708])) + (_dafny.SeqWithoutIsStrInference([224, (_dafny.MultiSet([not(True)])).cardinality])))

    @staticmethod
    def fm1(p0, globalState):
        return _dafny.MultiSet([(_dafny.CodePoint('c') if True else _dafny.CodePoint('h'))])

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        source0_ = D5_DC14(D5_DC13(_dafny.CodePoint('y'), False, 349, _dafny.CodePoint('u')))
        if source0_.is_DC12:
            d_1___mcc_h0_ = source0_.cf17
            d_2_cf17_ = d_1___mcc_h0_
            return D0_DC0(_dafny.CodePoint('d'))
        elif source0_.is_DC13:
            d_3___mcc_h1_ = source0_.cf18
            d_4___mcc_h2_ = source0_.cf19
            d_5___mcc_h3_ = source0_.cf20
            d_6___mcc_h4_ = source0_.cf21
            d_7_cf21_ = d_6___mcc_h4_
            d_8_cf20_ = d_5___mcc_h3_
            d_9_cf19_ = d_4___mcc_h2_
            d_10_cf18_ = d_3___mcc_h1_
            return D0_DC0(_dafny.CodePoint('v'))
        elif source0_.is_DC11:
            d_11___mcc_h5_ = source0_.cf16
            d_12_cf16_ = d_11___mcc_h5_
            return D0_DC0(_dafny.CodePoint('o'))
        elif True:
            d_13___mcc_h6_ = source0_.cf22
            d_14_cf22_ = d_13___mcc_h6_
            return D0_DC0(_dafny.CodePoint('y'))

    @staticmethod
    def fm4(p0, globalState):
        source1_ = D5_DC11(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_15_i0_ in range(default__.abs(588))])]))
        if source1_.is_DC12:
            d_16___mcc_h0_ = source1_.cf17
            d_17_cf17_ = d_16___mcc_h0_
            return d_17_cf17_
        elif source1_.is_DC13:
            d_18___mcc_h1_ = source1_.cf18
            d_19___mcc_h2_ = source1_.cf19
            d_20___mcc_h3_ = source1_.cf20
            d_21___mcc_h4_ = source1_.cf21
            d_22_cf21_ = d_21___mcc_h4_
            d_23_cf20_ = d_20___mcc_h3_
            d_24_cf19_ = d_19___mcc_h2_
            d_25_cf18_ = d_18___mcc_h1_
            return len(_dafny.SeqWithoutIsStrInference([d_23_cf20_]))
        elif source1_.is_DC11:
            d_26___mcc_h5_ = source1_.cf16
            d_27_cf16_ = d_26___mcc_h5_
            return 890
        elif True:
            d_28___mcc_h6_ = source1_.cf22
            d_29_cf22_ = d_28___mcc_h6_
            return -862

    @staticmethod
    def fm5(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hab"))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({245: len(_dafny.SeqWithoutIsStrInference([True]))})) | (_dafny.Map({478: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_30_i0_ in range(default__.abs(30))]))}))) | ((_dafny.Map({(_dafny.MultiSet([True])).cardinality: 247})))

    @staticmethod
    def fm7(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): 994})).keys.Elements:
                d_31_v0_: _dafny.Seq = compr_0_
                if (d_31_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): 994})):
                    coll0_[d_31_v0_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "af")))
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_32_i0_ in range(default__.abs(625))]))).Elements:
                d_33_v1_: _dafny.Seq = compr_1_
                if (d_33_v1_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_32_i0_ in range(default__.abs(625))]))):
                    coll1_[d_33_v1_] = 467
            return _dafny.Map(coll1_)
        return (iife0_()
        ) | ((iife1_()
        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(False), False, False, False, True]): 461})))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('y')

    @staticmethod
    def fm9(globalState):
        return ((_dafny.SeqWithoutIsStrInference([not(False), True, False, False, True])) + (_dafny.SeqWithoutIsStrInference([False, False]))) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([not(False)])))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return _dafny.Set({not((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_34_i0_ in range(default__.abs(274))])) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuonoct"))))})

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: str
            for compr_2_ in (_dafny.Set({_dafny.CodePoint('u')})).Elements:
                d_35_v0_: str = compr_2_
                if (d_35_v0_) in (_dafny.Set({_dafny.CodePoint('u')})):
                    coll2_[d_35_v0_] = len(_dafny.SeqWithoutIsStrInference([149, len(_dafny.SeqWithoutIsStrInference([False]))]))
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: str
            for compr_3_ in (_dafny.Map({_dafny.CodePoint('k'): False})).keys.Elements:
                d_36_v1_: str = compr_3_
                if (d_36_v1_) in (_dafny.Map({_dafny.CodePoint('k'): False})):
                    coll3_[d_36_v1_] = (_dafny.MultiSet([72])).cardinality
            return _dafny.Map(coll3_)
        return ((iife2_()
        ) | (_dafny.Map({_dafny.CodePoint('v'): 46}))) | ((iife3_()
        ) | (_dafny.Map({_dafny.CodePoint('n'): len(_dafny.SeqWithoutIsStrInference([518]))})))

    @staticmethod
    def fm12(globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(841, 902):
                d_37_v0_: int = compr_4_
                if ((841) <= (d_37_v0_)) and ((d_37_v0_) < (902)):
                    coll4_[(d_37_v0_) - ((_dafny.MultiSet([True])).cardinality)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_38_i0_ in range(default__.abs(-948))]))
            return _dafny.Map(coll4_)
        return not (True) or ((len(_dafny.Map({False: iife4_()
        }))) >= (-711))

    @staticmethod
    def fm13(globalState):
        return (_dafny.MultiSet([True, True])) | ((_dafny.MultiSet([False]) if True else _dafny.MultiSet([False, True])))

    @staticmethod
    def m1(p0, p1, p2, p3, globalState):
        r0: C0 = None
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: int = int(0)
        d_39_v0_: _dafny.Set
        d_39_v0_ = _dafny.Set({True, default__.fm12(globalState)})
        d_40_v1_: int
        d_40_v1_ = 952
        d_41_v4_: _dafny.Map
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(327, 247):
                d_42_v2_: int = compr_5_
                if ((327) <= (d_42_v2_)) and ((d_42_v2_) < (247)):
                    def iife6_():
                        coll6_ = _dafny.Set()
                        compr_6_: int
                        for compr_6_ in (_dafny.MultiSet(p3)).Elements:
                            d_43_v3_: int = compr_6_
                            if (d_43_v3_) in (_dafny.MultiSet(p3)):
                                coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_43_v3_, 665)]))
                        return _dafny.Set(coll6_)
                    coll5_ = coll5_.union(_dafny.Set([(d_42_v2_) * (len(iife6_()
))]))
            return _dafny.Set(coll5_)
        d_41_v4_ = _dafny.Map({iife5_()
        : p2})
        d_44_v5_: _dafny.Map
        d_44_v5_ = _dafny.Map({d_40_v1_: len(d_41_v4_)})
        d_45_i0_: int
        d_45_i0_ = 0
        with _dafny.label("0"):
            while (len(d_39_v0_)) > ((0) - (((d_44_v5_)[d_40_v1_] if (d_40_v1_) in (d_44_v5_) else d_40_v1_))):
                with _dafny.c_label("0"):
                    if (d_45_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_45_i0_ = (d_45_i0_) + (1)
                    (globalState).f2 = d_40_v1_
                    d_44_v5_ = (d_44_v5_).set(d_40_v1_, default__.safeModuloInt(d_40_v1_, 196))
                    d_46_v6_: D4
                    d_46_v6_ = D4_DC9(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: d_40_v1_})) for d_47_i1_ in range(default__.abs(234))])), p2)
                    source2_ = (d_46_v6_ if p2 else d_46_v6_)
                    if source2_.is_DC9:
                        d_48___mcc_h0_ = source2_.cf14
                        d_49___mcc_h1_ = source2_.cf15
                        d_50_cf15_ = d_49___mcc_h1_
                        d_51_cf14_ = d_48___mcc_h0_
                        d_52_v7_: _dafny.Array
                        nw0_ = _dafny.Array(None, 1)
                        nw0_[int(0)] = d_51_cf14_
                        d_52_v7_ = nw0_
                        d_53_v8_: _dafny.Seq
                        d_53_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        d_54_v9_: _dafny.MultiSet
                        d_54_v9_ = _dafny.MultiSet([d_53_v8_])
                        d_55_v10_: D5
                        d_55_v10_ = D5_DC11(d_54_v9_)
                        index0_ = default__.safeIndex(961, (d_52_v7_).length(0))
                        (d_52_v7_)[index0_] = ((d_54_v9_) | ((d_55_v10_).cf16)).cardinality
                        index1_ = default__.safeIndex(961, (d_52_v7_).length(0))
                        (d_52_v7_)[index1_] = ((d_44_v5_)[d_51_cf14_] if (d_51_cf14_) in (d_44_v5_) else d_40_v1_)
                        (globalState).f10 = p2
                        d_56_v11_: C0
                        nw1_ = C0()
                        nw1_.ctor__()
                        d_56_v11_ = nw1_
                        r0 = d_56_v11_
                        d_57_v12_: _dafny.Array
                        nw2_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
                        d_57_v12_ = nw2_
                        index2_ = default__.safeIndex(253, (d_57_v12_).length(0))
                        (d_57_v12_)[index2_] = d_53_v8_
                        index3_ = default__.safeIndex(253, (d_57_v12_).length(0))
                        (d_57_v12_)[index3_] = (d_53_v8_) + (d_53_v8_)
                    elif source2_.is_DC10:
                        d_58_v13_: _dafny.Set
                        d_58_v13_ = _dafny.Set({d_40_v1_, d_40_v1_, d_40_v1_, d_40_v1_, d_40_v1_})
                        d_59_v14_: _dafny.Map
                        d_59_v14_ = _dafny.Map({len(d_58_v13_): p2})
                        d_60_v15_: D0
                        d_60_v15_ = D0_DC1(not(((d_59_v14_)[d_40_v1_] if (d_40_v1_) in (d_59_v14_) else p2)), False, p2)
                        d_61_v16_: _dafny.Seq
                        d_61_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgxywina"))
                        d_62_v17_: _dafny.Map
                        d_62_v17_ = _dafny.Map({p2: p2})
                        r3 = ((len(d_61_v16_) if p2 else d_40_v1_) if (d_60_v15_).cf3 else len((_dafny.Map({p2: (d_46_v6_).cf15})) | (d_62_v17_)))
                        (globalState).f9 = -423
                        d_63_v18_: _dafny.Seq
                        d_63_v18_ = _dafny.SeqWithoutIsStrInference([p2])
                        d_64_v19_: _dafny.Array
                        nw3_ = _dafny.Array(None, 4)
                        nw3_[int(0)] = d_63_v18_
                        nw3_[int(1)] = _dafny.SeqWithoutIsStrInference([(d_63_v18_)[default__.safeIndex((0) - (d_40_v1_), len(d_63_v18_))]])
                        nw3_[int(2)] = _dafny.SeqWithoutIsStrInference([p2, not(p2)])
                        nw3_[int(3)] = (d_63_v18_) + (d_63_v18_)
                        d_64_v19_ = nw3_
                        r1 = d_64_v19_
                        d_65_v20_: C0
                        nw4_ = C0()
                        nw4_.ctor__()
                        d_65_v20_ = nw4_
                    elif True:
                        d_66___mcc_h2_ = source2_.cf13
                        d_67_cf13_ = d_66___mcc_h2_
                        d_68_v21_: C0
                        nw5_ = C0()
                        nw5_.ctor__()
                        d_68_v21_ = nw5_
                        r0 = (d_68_v21_ if p2 else d_68_v21_)
                        (globalState).f2 = ((d_40_v1_) * (d_40_v1_)) * ((d_40_v1_) - (-153))
                        (globalState).f9 = d_40_v1_
                        (globalState).f2 = (p3)[default__.safeIndex(d_40_v1_, len(p3))]
                    d_69_v22_: _dafny.Seq
                    d_69_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnyebc"))
                    d_40_v1_ = default__.safeModuloInt(len(d_69_v22_), (-808) + (133))
                    pass
            pass
        d_70_v23_: C0
        nw6_ = C0()
        nw6_.ctor__()
        d_70_v23_ = nw6_
        d_71_v24_: _dafny.Array
        nw7_ = _dafny.Array(None, 8)
        nw7_[int(0)] = d_70_v23_
        nw7_[int(1)] = d_70_v23_
        nw7_[int(2)] = d_70_v23_
        nw7_[int(3)] = d_70_v23_
        nw7_[int(4)] = d_70_v23_
        nw7_[int(5)] = d_70_v23_
        nw7_[int(6)] = d_70_v23_
        nw7_[int(7)] = d_70_v23_
        d_71_v24_ = nw7_
        rhs0_ = d_40_v1_
        rhs1_ = d_71_v24_
        rhs2_ = p2
        rhs3_ = d_40_v1_
        rhs4_ = d_70_v23_
        lhs0_ = globalState
        lhs1_ = globalState
        r3 = rhs0_
        d_71_v24_ = rhs1_
        lhs0_.f13 = rhs2_
        lhs1_.f2 = rhs3_
        d_70_v23_ = rhs4_
        d_72_v25_: _dafny.Seq
        d_72_v25_ = _dafny.SeqWithoutIsStrInference([p2])
        d_73_v26_: _dafny.Map
        d_73_v26_ = _dafny.Map({d_72_v25_: d_40_v1_})
        hi0_ = d_40_v1_
        for d_74_i2_ in range(default__.fm4(d_73_v26_, globalState), hi0_):
            d_75_v27_: C0
            nw8_ = C0()
            nw8_.ctor__()
            d_75_v27_ = nw8_
            (globalState).f10 = p2
            d_76_v28_: _dafny.Seq
            d_76_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dw"))
            d_77_v29_: _dafny.Map
            d_77_v29_ = _dafny.Map({d_40_v1_: d_76_v28_})
            d_78_v30_: _dafny.Array
            nw9_ = _dafny.Array(None, 6)
            nw9_[int(0)] = (d_76_v28_ if p2 else d_76_v28_)
            nw9_[int(1)] = d_76_v28_
            nw9_[int(2)] = ((d_77_v29_)[619] if (619) in (d_77_v29_) else d_76_v28_)
            nw9_[int(3)] = (d_76_v28_) + (d_76_v28_)
            nw9_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hicn"))
            nw9_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etaeup"))) + (d_76_v28_)
            d_78_v30_ = nw9_
            d_78_v30_ = d_78_v30_
            d_79_v31_: _dafny.Map
            d_79_v31_ = _dafny.Map({p2: p2})
            d_80_v32_: _dafny.Map
            d_80_v32_ = _dafny.Map({((d_79_v31_)[p2] if (p2) in (d_79_v31_) else p2): p2})
            d_81_v33_: _dafny.Map
            d_81_v33_ = _dafny.Map({p2: default__.fm10(p2, p2, d_40_v1_, p0, globalState)})
            (globalState).f10 = ((d_79_v31_)[True] if (True) in (d_79_v31_) else (_dafny.Set({p2, p2, p2, p2, p2})).issubset(((d_81_v33_)[p2] if (p2) in (d_81_v33_) else _dafny.Set({p2, p2}))))
        d_82_v34_: str
        d_82_v34_ = _dafny.CodePoint('s')
        d_82_v34_ = p0
        d_83_v35_: _dafny.Array
        def lambda0_(d_84_v1_):
            def lambda1_(d_85_i4_):
                return (d_85_i4_) - (d_84_v1_)

            return lambda1_

        init0_ = lambda0_(d_40_v1_)
        nw10_ = _dafny.Array(None, 23)
        for i0_0_ in range(nw10_.length(0)):
            nw10_[i0_0_] = init0_(i0_0_)
        d_83_v35_ = nw10_
        d_86_v36_: D1
        d_86_v36_ = D1_DC3(p3, p2)
        d_87_v37_: _dafny.Map
        d_87_v37_ = _dafny.Map({d_83_v35_: (d_86_v36_).cf5})
        d_88_v38_: _dafny.Map
        d_88_v38_ = _dafny.Map({p2: -464})
        d_89_v39_: _dafny.Seq
        d_89_v39_ = _dafny.SeqWithoutIsStrInference([d_88_v38_])
        d_90_v40_: _dafny.Array
        nw11_ = _dafny.Array(None, 28)
        nw11_[int(0)] = p2
        nw11_[int(1)] = p2
        nw11_[int(2)] = not(p2)
        nw11_[int(3)] = p2
        nw11_[int(4)] = default__.fm12(globalState)
        nw11_[int(5)] = False
        nw11_[int(6)] = False
        nw11_[int(7)] = p2
        nw11_[int(8)] = not(p2)
        nw11_[int(9)] = True
        nw11_[int(10)] = default__.fm12(globalState)
        nw11_[int(11)] = (d_87_v37_) == (_dafny.Map({d_83_v35_: p3}))
        nw11_[int(12)] = p2
        nw11_[int(13)] = p2
        nw11_[int(14)] = (d_70_v23_).fm3(not(p2), p2, default__.fm13(globalState), d_72_v25_, globalState)
        nw11_[int(15)] = p2
        nw11_[int(16)] = p2
        nw11_[int(17)] = p2
        nw11_[int(18)] = p2
        nw11_[int(19)] = True
        nw11_[int(20)] = p2
        nw11_[int(21)] = p2
        nw11_[int(22)] = (d_72_v25_) <= (_dafny.SeqWithoutIsStrInference([p2, p2]))
        nw11_[int(23)] = p2
        nw11_[int(24)] = p2
        nw11_[int(25)] = (d_89_v39_) <= (d_89_v39_)
        nw11_[int(26)] = p2
        nw11_[int(27)] = (d_40_v1_) < (default__.fm4(d_73_v26_, globalState))
        d_90_v40_ = nw11_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_90_v40_).length(0)):
            d_91_i3_: int = guard_loop_0_
            if (True) and (((0) <= (d_91_i3_)) and ((d_91_i3_) < ((d_90_v40_).length(0)))):
                (d_90_v40_)[(d_91_i3_)] = (_dafny.MultiSet([d_40_v1_])).issubset((_dafny.MultiSet([d_40_v1_]) if True else _dafny.MultiSet([d_40_v1_, len(_dafny.SeqWithoutIsStrInference([d_40_v1_])), d_40_v1_, d_40_v1_, d_40_v1_])))
        d_92_v41_: _dafny.Array
        nw12_ = _dafny.Array(_dafny.Map({}), 29)
        d_92_v41_ = nw12_
        index4_ = default__.safeIndex(789, (d_92_v41_).length(0))
        (d_92_v41_)[index4_] = _dafny.Map({p2: 842})
        index5_ = default__.safeIndex(789, (d_92_v41_).length(0))
        (d_92_v41_)[index5_] = (d_88_v38_) | ((d_88_v38_) | (d_88_v38_))
        nw13_ = C0()
        nw13_.ctor__()
        r0 = nw13_
        d_93_v42_: _dafny.Array
        nw14_ = _dafny.Array(None, 14)
        nw14_[int(0)] = d_72_v25_
        nw14_[int(1)] = _dafny.SeqWithoutIsStrInference([p2])
        nw14_[int(2)] = d_72_v25_
        nw14_[int(3)] = d_72_v25_
        nw14_[int(4)] = d_72_v25_
        nw14_[int(5)] = default__.fm9(globalState)
        nw14_[int(6)] = d_72_v25_
        nw14_[int(7)] = d_72_v25_
        nw14_[int(8)] = d_72_v25_
        nw14_[int(9)] = default__.fm9(globalState)
        nw14_[int(10)] = d_72_v25_
        nw14_[int(11)] = d_72_v25_
        nw14_[int(12)] = ((d_72_v25_).set(default__.safeIndex((0) - (d_40_v1_), len(d_72_v25_)), default__.fm12(globalState))).set(default__.safeIndex(d_40_v1_, len((d_72_v25_).set(default__.safeIndex((0) - (d_40_v1_), len(d_72_v25_)), default__.fm12(globalState)))), not(p2))
        nw14_[int(13)] = d_72_v25_
        d_93_v42_ = nw14_
        d_94_v43_: _dafny.Map
        d_94_v43_ = _dafny.Map({d_70_v23_: d_93_v42_})
        r1 = ((d_94_v43_)[d_70_v23_] if (d_70_v23_) in (d_94_v43_) else d_93_v42_)
        r2 = default__.safeModuloInt(d_40_v1_, d_40_v1_)
        r3 = d_40_v1_
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_95_v0_: _dafny.Array
        nw15_ = _dafny.Array(False, 20)
        d_95_v0_ = nw15_
        d_96_v1_: _dafny.Seq
        d_96_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwjfy"))
        d_97_globalState_: GlobalState
        nw16_ = GlobalState()
        nw16_.ctor__(511, _dafny.CodePoint('o'), -937, 29, d_95_v0_, False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_98_i0_ in range(default__.abs(220))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_99_i1_ in range(default__.abs(402))]), _dafny.CodePoint('y'), 949, False, _dafny.CodePoint('j'), 820, True, (d_96_v1_) + (d_96_v1_), False)
        d_97_globalState_ = nw16_
        d_100_v2_: bool
        d_100_v2_ = False
        d_101_v3_: int
        d_101_v3_ = 397
        if ((d_100_v2_) == (d_100_v2_)) == (not((d_101_v3_) >= (d_101_v3_))):
            d_102_v4_: _dafny.Seq
            d_102_v4_ = _dafny.SeqWithoutIsStrInference([d_101_v3_, 92, d_101_v3_, d_101_v3_, (_dafny.MultiSet([d_100_v2_])).cardinality])
            (d_97_globalState_).f0 = (0) - (len(default__.fm0(d_102_v4_, default__.safeModuloInt(d_101_v3_, d_101_v3_), d_97_globalState_)))
            d_103_v5_: _dafny.MultiSet
            d_103_v5_ = _dafny.MultiSet([-498, d_101_v3_])
            (d_97_globalState_).f10 = not(((d_103_v5_) | (d_103_v5_)).issubset(d_103_v5_))
            if d_100_v2_:
                (d_97_globalState_).f7 = d_96_v1_
                d_104_v6_: _dafny.Map
                d_104_v6_ = _dafny.Map({(241 if d_100_v2_ else (d_103_v5_).cardinality): default__.fm1(d_101_v3_, d_97_globalState_)})
                d_105_v7_: str
                d_105_v7_ = _dafny.CodePoint('a')
                d_106_v8_: _dafny.MultiSet
                d_106_v8_ = _dafny.MultiSet([(default__.fm2(d_100_v2_, d_100_v2_, d_101_v3_, d_97_globalState_)).cf0, d_105_v7_, _dafny.CodePoint('t'), d_105_v7_])
                d_104_v6_ = (d_104_v6_).set(len(_dafny.SeqWithoutIsStrInference([936 for d_107_i2_ in range(default__.abs(96))])), d_106_v8_)
                d_108_v9_: C0
                nw17_ = C0()
                nw17_.ctor__()
                d_108_v9_ = nw17_
                d_109_v10_: _dafny.Array
                nw18_ = _dafny.Array(int(0), 24)
                d_109_v10_ = nw18_
                d_110_v11_: _dafny.Map
                d_110_v11_ = _dafny.Map({d_100_v2_: d_109_v10_})
                d_111_v12_: _dafny.Seq
                d_111_v12_ = _dafny.SeqWithoutIsStrInference([d_109_v10_, ((d_110_v11_)[d_100_v2_] if (d_100_v2_) in (d_110_v11_) else d_109_v10_), d_109_v10_])
                (d_97_globalState_).f9 = len(d_111_v12_)
                d_112_v13_: C0
                nw19_ = C0()
                nw19_.ctor__()
                d_112_v13_ = nw19_
            elif True:
                d_113_v14_: str
                d_113_v14_ = _dafny.CodePoint('l')
                d_114_v15_: D0
                d_114_v15_ = D0_DC0(d_113_v14_)
                d_115_v16_: _dafny.Seq
                d_115_v16_ = _dafny.SeqWithoutIsStrInference([d_114_v15_])
                d_116_v17_: C0
                d_117_v18_: _dafny.Array
                d_118_v19_: int
                d_119_v20_: int
                out0_: C0
                out1_: _dafny.Array
                out2_: int
                out3_: int
                out0_, out1_, out2_, out3_ = default__.m1(_dafny.CodePoint('j'), d_115_v16_, d_100_v2_, (d_102_v4_) + (_dafny.SeqWithoutIsStrInference([d_101_v3_])), d_97_globalState_)
                d_116_v17_ = out0_
                d_117_v18_ = out1_
                d_118_v19_ = out2_
                d_119_v20_ = out3_
                (d_97_globalState_).f10 = d_100_v2_
                d_120_v21_: _dafny.Map
                d_120_v21_ = _dafny.Map({_dafny.CodePoint('t'): d_100_v2_})
                d_121_v22_: _dafny.Set
                d_121_v22_ = _dafny.Set({d_120_v21_})
                d_122_v23_: _dafny.Seq
                d_122_v23_ = _dafny.SeqWithoutIsStrInference([d_96_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckdgwau")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vu"))])
                d_123_v24_: _dafny.Map
                d_123_v24_ = _dafny.Map({d_100_v2_: d_122_v23_})
                d_124_v25_: _dafny.Array
                nw20_ = _dafny.Array(None, 29)
                nw20_[int(0)] = len(d_121_v22_)
                nw20_[int(1)] = d_118_v19_
                nw20_[int(2)] = d_118_v19_
                nw20_[int(3)] = d_118_v19_
                nw20_[int(4)] = (0) - (d_119_v20_)
                nw20_[int(5)] = d_101_v3_
                nw20_[int(6)] = len(d_96_v1_)
                nw20_[int(7)] = d_118_v19_
                nw20_[int(8)] = d_118_v19_
                nw20_[int(9)] = d_101_v3_
                nw20_[int(10)] = d_119_v20_
                nw20_[int(11)] = d_118_v19_
                nw20_[int(12)] = d_101_v3_
                nw20_[int(13)] = d_119_v20_
                nw20_[int(14)] = 597
                nw20_[int(15)] = len(((d_123_v24_)[d_100_v2_] if (d_100_v2_) in (d_123_v24_) else d_122_v23_))
                nw20_[int(16)] = d_101_v3_
                nw20_[int(17)] = d_101_v3_
                nw20_[int(18)] = d_118_v19_
                nw20_[int(19)] = (0) - (d_119_v20_)
                nw20_[int(20)] = d_101_v3_
                nw20_[int(21)] = d_118_v19_
                nw20_[int(22)] = d_119_v20_
                nw20_[int(23)] = d_101_v3_
                nw20_[int(24)] = d_118_v19_
                nw20_[int(25)] = (_dafny.MultiSet([d_100_v2_])).cardinality
                nw20_[int(26)] = d_118_v19_
                nw20_[int(27)] = d_119_v20_
                nw20_[int(28)] = 291
                d_124_v25_ = nw20_
                d_125_v26_: _dafny.Array
                nw21_ = _dafny.Array(None, 26)
                nw21_[int(0)] = d_124_v25_
                nw21_[int(1)] = d_124_v25_
                nw21_[int(2)] = d_124_v25_
                nw21_[int(3)] = d_124_v25_
                nw21_[int(4)] = d_124_v25_
                nw21_[int(5)] = d_124_v25_
                nw21_[int(6)] = d_124_v25_
                nw21_[int(7)] = d_124_v25_
                nw21_[int(8)] = d_124_v25_
                nw21_[int(9)] = d_124_v25_
                nw21_[int(10)] = d_124_v25_
                nw21_[int(11)] = d_124_v25_
                nw21_[int(12)] = d_124_v25_
                nw21_[int(13)] = d_124_v25_
                nw21_[int(14)] = d_124_v25_
                nw21_[int(15)] = d_124_v25_
                nw21_[int(16)] = d_124_v25_
                nw21_[int(17)] = d_124_v25_
                nw21_[int(18)] = d_124_v25_
                nw21_[int(19)] = d_124_v25_
                nw21_[int(20)] = d_124_v25_
                nw21_[int(21)] = d_124_v25_
                nw21_[int(22)] = d_124_v25_
                nw21_[int(23)] = d_124_v25_
                nw21_[int(24)] = d_124_v25_
                nw21_[int(25)] = d_124_v25_
                d_125_v26_ = nw21_
                index6_ = default__.safeIndex(168, (d_125_v26_).length(0))
                (d_125_v26_)[index6_] = d_124_v25_
                index7_ = default__.safeIndex(168, (d_125_v26_).length(0))
                (d_125_v26_)[index7_] = d_124_v25_
                d_126_v27_: _dafny.Array
                nw22_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                d_126_v27_ = nw22_
                nw23_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
                d_126_v27_ = nw23_
                d_127_v28_: _dafny.Seq
                d_127_v28_ = _dafny.SeqWithoutIsStrInference([d_100_v2_])
                (d_97_globalState_).f0 = (0) - ((0) - (default__.safeModuloInt(len(d_96_v1_), len(d_127_v28_))))
            d_128_v29_: str
            d_128_v29_ = _dafny.CodePoint('d')
            d_129_v30_: D0
            d_129_v30_ = D0_DC0(d_128_v29_)
            d_130_v31_: _dafny.Seq
            d_130_v31_ = _dafny.SeqWithoutIsStrInference([d_129_v30_, d_129_v30_, D0_DC0(d_128_v29_)])
            d_131_v32_: _dafny.Map
            d_131_v32_ = _dafny.Map({d_101_v3_: d_101_v3_})
            d_132_v33_: C0
            d_133_v34_: _dafny.Array
            d_134_v35_: int
            d_135_v36_: int
            out4_: C0
            out5_: _dafny.Array
            out6_: int
            out7_: int
            out4_, out5_, out6_, out7_ = default__.m1(d_128_v29_, d_130_v31_, d_100_v2_, _dafny.SeqWithoutIsStrInference([len(d_131_v32_), d_101_v3_]), d_97_globalState_)
            d_132_v33_ = out4_
            d_133_v34_ = out5_
            d_134_v35_ = out6_
            d_135_v36_ = out7_
            d_136_v37_: _dafny.Array
            def lambda2_(d_137_v29_):
                def lambda3_(d_138_i3_):
                    return _dafny.SeqWithoutIsStrInference([d_137_v29_ for d_139_i4_ in range(default__.abs(168))])

                return lambda3_

            init1_ = lambda2_(d_128_v29_)
            nw24_ = _dafny.Array(None, 1)
            for i0_1_ in range(nw24_.length(0)):
                nw24_[i0_1_] = init1_(i0_1_)
            d_136_v37_ = nw24_
            index8_ = default__.safeIndex(191, (d_136_v37_).length(0))
            (d_136_v37_)[index8_] = d_96_v1_
            d_140_v38_: _dafny.Seq
            d_140_v38_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alnbjixe")), d_96_v1_])
            d_141_v39_: _dafny.Map
            d_141_v39_ = _dafny.Map({d_100_v2_: d_100_v2_})
            d_142_v40_: _dafny.Map
            d_142_v40_ = _dafny.Map({len(d_141_v39_): d_96_v1_})
            d_143_v41_: _dafny.Seq
            d_143_v41_ = _dafny.SeqWithoutIsStrInference([d_100_v2_, d_100_v2_])
            index9_ = default__.safeIndex(191, (d_136_v37_).length(0))
            rhs5_ = ((d_140_v38_)[default__.safeIndex(4, len(d_140_v38_))]) <= ((default__.fm5(d_100_v2_, d_100_v2_, d_97_globalState_)) + ((d_96_v1_).set(default__.safeIndex((d_103_v5_).cardinality, len(d_96_v1_)), d_128_v29_)))
            rhs6_ = (_dafny.SeqWithoutIsStrInference([d_128_v29_ for d_144_i5_ in range(default__.abs(-282))])) + (((d_142_v40_)[(0) - (len(d_143_v41_))] if ((0) - (len(d_143_v41_))) in (d_142_v40_) else d_96_v1_))
            rhs7_ = d_96_v1_
            lhs2_ = d_97_globalState_
            lhs3_ = d_97_globalState_
            lhs4_ = d_136_v37_
            lhs5_ = default__.safeIndex(191, (d_136_v37_).length(0))
            lhs2_.f15 = rhs5_
            lhs3_.f7 = rhs6_
            lhs4_[lhs5_] = rhs7_
        elif True:
            d_145_v42_: _dafny.Map
            d_145_v42_ = _dafny.Map({d_100_v2_: 726})
            d_146_v43_: _dafny.Seq
            d_146_v43_ = _dafny.SeqWithoutIsStrInference([((d_145_v42_)[d_100_v2_] if (d_100_v2_) in (d_145_v42_) else -708), len(d_96_v1_), -797, d_101_v3_])
            d_147_v44_: _dafny.Seq
            d_147_v44_ = _dafny.SeqWithoutIsStrInference([d_146_v43_, d_146_v43_, _dafny.SeqWithoutIsStrInference([d_101_v3_ for d_148_i6_ in range(default__.abs(671))])])
            d_147_v44_ = d_147_v44_
            index10_ = default__.safeIndex(47, (d_95_v0_).length(0))
            (d_95_v0_)[index10_] = d_100_v2_
            index11_ = default__.safeIndex(47, (d_95_v0_).length(0))
            (d_95_v0_)[index11_] = True
            d_149_v45_: D1
            d_149_v45_ = D1_DC4((d_95_v0_)[default__.safeIndex(47, (d_95_v0_).length(0))], d_101_v3_, d_96_v1_)
            d_150_v46_: _dafny.Set
            d_150_v46_ = _dafny.Set({d_149_v45_, d_149_v45_, d_149_v45_})
            d_150_v46_ = d_150_v46_
            d_151_v47_: _dafny.MultiSet
            d_151_v47_ = _dafny.MultiSet([d_100_v2_])
            (d_97_globalState_).f0 = (((d_151_v47_) | (d_151_v47_)).cardinality) * (d_101_v3_)
            d_152_v48_: _dafny.Map
            d_152_v48_ = _dafny.Map({len(d_145_v42_): d_101_v3_})
            d_152_v48_ = (d_152_v48_).set(d_101_v3_, d_101_v3_)
        if d_100_v2_:
            (d_97_globalState_).f13 = ((d_100_v2_ if False else d_100_v2_)) == (False)
            d_153_v49_: C0
            nw25_ = C0()
            nw25_.ctor__()
            d_153_v49_ = nw25_
            d_154_v50_: str
            d_154_v50_ = _dafny.CodePoint('f')
            d_154_v50_ = d_154_v50_
            d_155_v51_: _dafny.Set
            d_155_v51_ = _dafny.Set({d_100_v2_})
            d_156_v52_: _dafny.Seq
            d_156_v52_ = _dafny.SeqWithoutIsStrInference([d_101_v3_, len(_dafny.Map({d_100_v2_: d_101_v3_})), len(d_155_v51_), 298])
            d_157_v53_: D1
            d_157_v53_ = D1_DC3(d_156_v52_, d_100_v2_)
            if not((d_157_v53_).cf6):
                d_158_v54_: _dafny.Seq
                d_158_v54_ = _dafny.SeqWithoutIsStrInference([True])
                (d_97_globalState_).f10 = ((d_101_v3_) * (d_101_v3_)) == (default__.fm4(_dafny.Map({d_158_v54_: 904}), d_97_globalState_))
                (d_97_globalState_).f10 = d_100_v2_
                d_159_v55_: _dafny.Array
                nw26_ = _dafny.Array(_dafny.Map({}), 18)
                d_159_v55_ = nw26_
                d_160_v56_: _dafny.Seq
                d_160_v56_ = _dafny.SeqWithoutIsStrInference([d_153_v49_])
                d_161_v57_: _dafny.Map
                d_161_v57_ = _dafny.Map({d_101_v3_: len(d_160_v56_)})
                d_162_v58_: D1
                d_162_v58_ = D1_DC4(d_100_v2_, d_101_v3_, d_96_v1_)
                d_163_v59_: _dafny.Seq
                d_163_v59_ = _dafny.SeqWithoutIsStrInference([d_161_v57_])
                d_164_v60_: _dafny.Map
                d_164_v60_ = _dafny.Map({d_100_v2_: d_153_v49_})
                d_165_v61_: _dafny.Map
                d_165_v61_ = d_161_v57_
                d_166_v62_: _dafny.Map
                d_166_v62_ = (d_165_v61_)
                nw27_ = _dafny.Array(None, 17)
                nw27_[int(0)] = d_161_v57_
                nw27_[int(1)] = d_161_v57_
                nw27_[int(2)] = d_161_v57_
                nw27_[int(3)] = _dafny.Map({d_101_v3_: (d_162_v58_).cf8})
                nw27_[int(4)] = ((d_161_v57_).set(d_101_v3_, d_101_v3_)) | (d_161_v57_)
                nw27_[int(5)] = default__.fm6(False, d_100_v2_, d_101_v3_, 702, d_97_globalState_)
                nw27_[int(6)] = (d_161_v57_) | ((d_163_v59_)[default__.safeIndex(-488, len(d_163_v59_))])
                nw27_[int(7)] = d_161_v57_
                nw27_[int(8)] = (d_161_v57_).set(d_101_v3_, len(d_164_v60_))
                nw27_[int(9)] = (d_165_v61_)
                nw27_[int(10)] = d_161_v57_
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(-838, -263):
                        d_167_v63_: int = compr_7_
                        if ((-838) <= (d_167_v63_)) and ((d_167_v63_) < (-263)):
                            coll7_[(d_167_v63_) - (d_101_v3_)] = d_101_v3_
                    return _dafny.Map(coll7_)
                nw27_[int(11)] = iife7_()
                
                nw27_[int(12)] = d_161_v57_
                nw27_[int(13)] = (d_161_v57_).set(d_101_v3_, d_101_v3_)
                nw27_[int(14)] = _dafny.Map({446: d_101_v3_})
                nw27_[int(15)] = d_161_v57_
                def iife8_():
                    coll8_ = _dafny.Map()
                    compr_8_: int
                    for compr_8_ in (d_156_v52_).Elements:
                        d_168_v64_: int = compr_8_
                        if (d_168_v64_) in (d_156_v52_):
                            coll8_[default__.safeModuloInt(d_168_v64_, d_101_v3_)] = 991
                    return _dafny.Map(coll8_)
                nw27_[int(16)] = iife8_()
                
                d_159_v55_ = nw27_
                d_169_v65_: _dafny.Map
                d_169_v65_ = _dafny.Map({d_100_v2_: len((d_156_v52_).set(default__.safeIndex(d_101_v3_, len(d_156_v52_)), d_101_v3_))})
                d_170_v66_: _dafny.MultiSet
                d_170_v66_ = _dafny.MultiSet([d_100_v2_])
                d_169_v65_ = (d_169_v65_).set((d_96_v1_) <= (d_96_v1_), (d_101_v3_ if not((d_153_v49_).fm3(d_100_v2_, False, d_170_v66_, _dafny.SeqWithoutIsStrInference([d_100_v2_]), d_97_globalState_)) else d_101_v3_))
                index12_ = default__.safeIndex(355, (d_95_v0_).length(0))
                (d_95_v0_)[index12_] = d_100_v2_
                index13_ = default__.safeIndex(355, (d_95_v0_).length(0))
                (d_95_v0_)[index13_] = False
            elif True:
                d_95_v0_ = d_95_v0_
                d_171_v67_: C0
                nw28_ = C0()
                nw28_.ctor__()
                d_171_v67_ = nw28_
                (d_97_globalState_).f0 = d_101_v3_
                (d_97_globalState_).f10 = d_100_v2_
                d_101_v3_ = (0) - ((0) - ((d_101_v3_ if d_100_v2_ else 421)))
            (d_97_globalState_).f15 = d_100_v2_
        elif True:
            (d_97_globalState_).f2 = d_101_v3_
            d_172_v68_: _dafny.Seq
            d_172_v68_ = _dafny.SeqWithoutIsStrInference([True])
            d_173_v69_: _dafny.Map
            d_173_v69_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_100_v2_, d_100_v2_])) + (d_172_v68_): d_100_v2_})
            d_173_v69_ = d_173_v69_
            d_174_v70_: _dafny.MultiSet
            d_174_v70_ = _dafny.MultiSet([d_100_v2_, d_100_v2_])
            d_175_v71_: _dafny.Array
            nw29_ = _dafny.Array(None, 4)
            nw29_[int(0)] = d_174_v70_
            nw29_[int(1)] = (d_174_v70_) - (d_174_v70_)
            nw29_[int(2)] = d_174_v70_
            nw29_[int(3)] = (d_174_v70_).set(d_100_v2_, default__.abs(d_101_v3_))
            d_175_v71_ = nw29_
            index14_ = default__.safeIndex(947, (d_175_v71_).length(0))
            (d_175_v71_)[index14_] = (d_174_v70_).intersection(d_174_v70_)
            index15_ = default__.safeIndex(947, (d_175_v71_).length(0))
            (d_175_v71_)[index15_] = (_dafny.MultiSet([d_100_v2_])) | (d_174_v70_)
            d_176_v72_: C0
            nw30_ = C0()
            nw30_.ctor__()
            d_176_v72_ = nw30_
            d_177_v73_: C0
            d_177_v73_ = d_176_v72_
            d_176_v72_ = (d_176_v72_ if d_100_v2_ else ((d_177_v73_) if d_100_v2_ else d_176_v72_))
            if d_100_v2_:
                (d_176_v72_).m0(d_172_v68_, d_101_v3_, d_97_globalState_)
                d_178_v74_: _dafny.Array
                def lambda4_(d_179_i7_):
                    return default__.safeDivisionInt(d_179_i7_, len(_dafny.Set({_dafny.CodePoint('y')})))

                init2_ = lambda4_
                nw31_ = _dafny.Array(None, 18)
                for i0_2_ in range(nw31_.length(0)):
                    nw31_[i0_2_] = init2_(i0_2_)
                d_178_v74_ = nw31_
                d_180_v75_: _dafny.Map
                d_180_v75_ = _dafny.Map({default__.fm5(d_100_v2_, True, d_97_globalState_): d_178_v74_})
                d_181_v76_: _dafny.Seq
                d_181_v76_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_96_v1_: d_178_v74_}), (d_180_v75_).set(d_96_v1_, d_178_v74_), d_180_v75_])
                d_182_v77_: _dafny.Seq
                d_182_v77_ = _dafny.SeqWithoutIsStrInference([d_101_v3_, 510, -366])
                d_183_v78_: _dafny.Map
                d_183_v78_ = _dafny.Map({(d_181_v76_)[default__.safeIndex(len(d_96_v1_), len(d_181_v76_))]: default__.safeDivisionInt(d_101_v3_, (d_182_v77_)[default__.safeIndex((0) - (len(_dafny.Set({d_100_v2_, d_100_v2_}))), len(d_182_v77_))])})
                d_183_v78_ = (d_183_v78_).set((d_180_v75_) | (d_180_v75_), -798)
                (d_97_globalState_).f13 = d_100_v2_
                index16_ = default__.safeIndex(549, (d_178_v74_).length(0))
                (d_178_v74_)[index16_] = d_101_v3_
                d_184_v79_: _dafny.MultiSet
                d_184_v79_ = _dafny.MultiSet([d_101_v3_, d_101_v3_])
                d_185_v80_: _dafny.Map
                d_185_v80_ = _dafny.Map({d_176_v72_: d_184_v79_})
                index17_ = default__.safeIndex(549, (d_178_v74_).length(0))
                (d_178_v74_)[index17_] = (618) + ((len(d_185_v80_)) * (d_101_v3_))
                d_186_v81_: _dafny.Set
                d_186_v81_ = _dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_100_v2_, d_100_v2_]))).cardinality})
                d_187_v82_: _dafny.Map
                d_187_v82_ = _dafny.Map({d_100_v2_: (d_178_v74_)[default__.safeIndex(549, (d_178_v74_).length(0))]})
                d_188_v83_: _dafny.Map
                d_188_v83_ = _dafny.Map({(d_186_v81_).issubset(d_186_v81_): (d_96_v1_)[default__.safeIndex(default__.fm4(default__.fm7(len(default__.fm5(True, False, d_97_globalState_)), len(d_187_v82_), d_97_globalState_), d_97_globalState_), len(d_96_v1_))]})
                d_189_v84_: _dafny.Array
                def lambda5_(d_190_i8_):
                    return _dafny.CodePoint('q')

                init3_ = lambda5_
                nw32_ = _dafny.Array(None, 16)
                for i0_3_ in range(nw32_.length(0)):
                    nw32_[i0_3_] = init3_(i0_3_)
                d_189_v84_ = nw32_
                d_191_v85_: _dafny.MultiSet
                d_191_v85_ = _dafny.MultiSet([d_189_v84_])
                d_192_v86_: _dafny.Seq
                d_192_v86_ = _dafny.SeqWithoutIsStrInference([d_191_v85_, d_191_v85_])
                d_193_v87_: _dafny.Seq
                d_193_v87_ = _dafny.SeqWithoutIsStrInference([d_189_v84_])
                d_194_v88_: str
                d_194_v88_ = _dafny.CodePoint('s')
                d_188_v83_ = (d_188_v83_).set(((d_192_v86_)[default__.safeIndex(976, len(d_192_v86_))]).isdisjoint(_dafny.MultiSet([(d_193_v87_)[default__.safeIndex(d_101_v3_, len(d_193_v87_))], d_189_v84_, d_189_v84_])), d_194_v88_)
            elif True:
                d_195_v89_: _dafny.Seq
                d_195_v89_ = _dafny.SeqWithoutIsStrInference([d_101_v3_])
                d_196_v90_: _dafny.MultiSet
                d_196_v90_ = _dafny.MultiSet([d_101_v3_])
                d_197_v91_: _dafny.Map
                d_197_v91_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_100_v2_]): d_101_v3_})
                d_198_v92_: str
                d_198_v92_ = _dafny.CodePoint('r')
                d_199_v93_: _dafny.MultiSet
                d_199_v93_ = _dafny.MultiSet([default__.fm8(d_195_v89_, (d_196_v90_).set(-538, default__.abs((_dafny.MultiSet([d_101_v3_, d_101_v3_, d_101_v3_, d_101_v3_])).cardinality)), d_100_v2_, default__.fm4(d_197_v91_, d_97_globalState_), d_97_globalState_), d_198_v92_, d_198_v92_, d_198_v92_])
                d_200_v94_: _dafny.Seq
                d_200_v94_ = _dafny.SeqWithoutIsStrInference([d_101_v3_, (d_199_v93_).cardinality, default__.fm4(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): d_101_v3_}), d_97_globalState_), d_101_v3_])
                d_101_v3_ = (0) - ((d_200_v94_)[default__.safeIndex((0) - (d_101_v3_), len(d_200_v94_))])
                d_201_v95_: _dafny.Set
                d_201_v95_ = _dafny.Set({d_100_v2_})
                d_202_v96_: _dafny.Seq
                d_202_v96_ = _dafny.SeqWithoutIsStrInference([d_201_v95_, d_201_v95_, d_201_v95_, d_201_v95_])
                d_203_v97_: _dafny.Map
                d_203_v97_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len((d_202_v96_)[default__.safeIndex(d_101_v3_, len(d_202_v96_))])]): d_176_v72_})
                d_204_v98_: C0
                d_205_v99_: _dafny.Array
                d_206_v100_: int
                d_207_v101_: int
                out8_: C0
                out9_: _dafny.Array
                out10_: int
                out11_: int
                out8_, out9_, out10_, out11_ = default__.m1(d_198_v92_, _dafny.SeqWithoutIsStrInference([D0_DC0(d_198_v92_) for d_208_i9_ in range(default__.abs(981))]), d_100_v2_, (d_195_v89_).set(default__.safeIndex(d_101_v3_, len(d_195_v89_)), len(d_203_v97_)), d_97_globalState_)
                d_204_v98_ = out8_
                d_205_v99_ = out9_
                d_206_v100_ = out10_
                d_207_v101_ = out11_
                d_101_v3_ = default__.fm4(default__.fm7(d_206_v100_, d_207_v101_, d_97_globalState_), d_97_globalState_)
                d_209_v102_: _dafny.Seq
                d_209_v102_ = _dafny.SeqWithoutIsStrInference([d_172_v68_, default__.fm9(d_97_globalState_)])
                d_210_v103_: _dafny.Map
                d_210_v103_ = _dafny.Map({d_100_v2_: d_100_v2_})
                d_211_v104_: _dafny.Seq
                d_211_v104_ = _dafny.SeqWithoutIsStrInference([d_210_v103_])
                (d_176_v72_).m0((d_209_v102_)[default__.safeIndex(len(d_211_v104_), len(d_209_v102_))], default__.fm4(d_197_v91_, d_97_globalState_), d_97_globalState_)
                d_212_v105_: D0
                d_212_v105_ = D0_DC0(d_198_v92_)
                d_213_v106_: _dafny.Seq
                d_213_v106_ = _dafny.SeqWithoutIsStrInference([d_212_v105_])
                d_214_v108_: C0
                d_215_v109_: _dafny.Array
                d_216_v110_: int
                d_217_v111_: int
                out12_: C0
                out13_: _dafny.Array
                out14_: int
                out15_: int
                def iife9_():
                    coll9_ = _dafny.Map()
                    compr_9_: _dafny.Seq
                    for compr_9_ in (_dafny.SeqWithoutIsStrInference([d_172_v68_])).Elements:
                        d_218_v107_: _dafny.Seq = compr_9_
                        if (d_218_v107_) in (_dafny.SeqWithoutIsStrInference([d_172_v68_])):
                            coll9_[d_218_v107_] = len(d_201_v95_)
                    return _dafny.Map(coll9_)
                out12_, out13_, out14_, out15_ = default__.m1(d_198_v92_, d_213_v106_, d_100_v2_, _dafny.SeqWithoutIsStrInference([(0) - (d_207_v101_), default__.fm4(iife9_()
                , d_97_globalState_)]), d_97_globalState_)
                d_214_v108_ = out12_
                d_215_v109_ = out13_
                d_216_v110_ = out14_
                d_217_v111_ = out15_
        d_219_v112_: _dafny.Array
        nw33_ = _dafny.Array(None, 17)
        d_219_v112_ = nw33_
        d_220_v113_: _dafny.Seq
        d_220_v113_ = _dafny.SeqWithoutIsStrInference([d_219_v112_])
        (d_97_globalState_).f2 = len(d_220_v113_)
        d_221_v114_: C0
        nw34_ = C0()
        nw34_.ctor__()
        d_221_v114_ = nw34_
        if d_100_v2_:
            d_222_v115_: _dafny.Array
            nw35_ = _dafny.Array(int(0), 5)
            d_222_v115_ = nw35_
            index18_ = default__.safeIndex(789, (d_222_v115_).length(0))
            (d_222_v115_)[index18_] = d_101_v3_
            index19_ = default__.safeIndex(789, (d_222_v115_).length(0))
            (d_222_v115_)[index19_] = default__.safeDivisionInt(d_101_v3_, d_101_v3_)
            if not (d_100_v2_) or ((not(d_100_v2_)) and (d_100_v2_)):
                d_223_v116_: _dafny.Array
                nw36_ = _dafny.Array(None, 9)
                nw36_[int(0)] = d_95_v0_
                nw36_[int(1)] = d_95_v0_
                nw36_[int(2)] = d_95_v0_
                nw36_[int(3)] = d_95_v0_
                nw36_[int(4)] = d_95_v0_
                nw36_[int(5)] = d_95_v0_
                nw36_[int(6)] = d_95_v0_
                nw36_[int(7)] = d_95_v0_
                nw36_[int(8)] = d_95_v0_
                d_223_v116_ = nw36_
                d_224_v117_: _dafny.Map
                d_224_v117_ = _dafny.Map({d_223_v116_: False})
                d_225_v118_: _dafny.MultiSet
                d_225_v118_ = _dafny.MultiSet([d_100_v2_])
                d_226_v119_: _dafny.Seq
                d_226_v119_ = _dafny.SeqWithoutIsStrInference([True, d_100_v2_, d_100_v2_])
                d_224_v117_ = (d_224_v117_).set(d_223_v116_, not((d_221_v114_).fm3(d_100_v2_, d_100_v2_, d_225_v118_, d_226_v119_, d_97_globalState_)))
                (d_97_globalState_).f10 = not((d_100_v2_) and (d_100_v2_))
                (d_221_v114_).m0(d_226_v119_, d_101_v3_, d_97_globalState_)
                index20_ = default__.safeIndex(789, (d_222_v115_).length(0))
                (d_222_v115_)[index20_] = (0) - (default__.safeModuloInt(default__.safeDivisionInt(142, 898), (d_222_v115_)[default__.safeIndex(789, (d_222_v115_).length(0))]))
                (d_97_globalState_).f2 = (0) - (((0) - ((d_222_v115_)[default__.safeIndex(789, (d_222_v115_).length(0))]) if d_100_v2_ else d_101_v3_))
            elif True:
                d_96_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkgtl"))
                d_221_v114_ = d_221_v114_
                index21_ = default__.safeIndex(789, (d_222_v115_).length(0))
                (d_222_v115_)[index21_] = (d_222_v115_)[default__.safeIndex(789, (d_222_v115_).length(0))]
                d_227_v120_: _dafny.Map
                d_227_v120_ = _dafny.Map({d_100_v2_: d_100_v2_})
                d_228_v121_: _dafny.Map
                d_228_v121_ = _dafny.Map({d_100_v2_: d_227_v120_})
                d_229_v122_: _dafny.Map
                d_229_v122_ = _dafny.Map({_dafny.MultiSet([d_227_v120_, ((d_228_v121_)[True] if (True) in (d_228_v121_) else d_227_v120_)]): d_100_v2_})
                d_230_v123_: _dafny.MultiSet
                d_230_v123_ = _dafny.MultiSet([_dafny.Map({d_100_v2_: d_100_v2_}), d_227_v120_, d_227_v120_])
                (d_97_globalState_).f13 = ((d_229_v122_)[d_230_v123_] if (d_230_v123_) in (d_229_v122_) else d_100_v2_)
                d_231_v124_: str
                d_231_v124_ = _dafny.CodePoint('w')
                d_232_v125_: _dafny.Map
                d_232_v125_ = _dafny.Map({d_231_v124_: d_95_v0_})
                (d_97_globalState_).f13 = (d_231_v124_) in (d_232_v125_)
            nw37_ = _dafny.Array(False, 11)
            d_95_v0_ = nw37_
            if d_100_v2_:
                d_101_v3_ = 285
                d_233_v126_: _dafny.Seq
                d_233_v126_ = _dafny.SeqWithoutIsStrInference([d_100_v2_, d_100_v2_, d_100_v2_])
                (d_221_v114_).m0(d_233_v126_, (d_222_v115_)[default__.safeIndex(789, (d_222_v115_).length(0))], d_97_globalState_)
                d_234_v127_: _dafny.Set
                d_234_v127_ = _dafny.Set({d_100_v2_, d_100_v2_})
                d_235_v128_: _dafny.Map
                d_235_v128_ = _dafny.Map({(d_222_v115_)[default__.safeIndex(789, (d_222_v115_).length(0))]: len(d_234_v127_)})
                d_236_v129_: _dafny.MultiSet
                d_236_v129_ = _dafny.MultiSet([d_100_v2_])
                d_237_v130_: _dafny.Map
                d_237_v130_ = _dafny.Map({d_100_v2_: (d_221_v114_).fm3(False, d_100_v2_, d_236_v129_, d_233_v126_, d_97_globalState_)})
                d_238_v131_: _dafny.Map
                d_238_v131_ = _dafny.Map({d_100_v2_: ((d_235_v128_)[577] if (577) in (d_235_v128_) else len(d_237_v130_))})
                d_239_v132_: _dafny.Map
                d_239_v132_ = _dafny.Map({d_101_v3_: d_238_v131_})
                d_240_v133_: D1
                d_240_v133_ = D1_DC4(False, (d_222_v115_)[default__.safeIndex(789, (d_222_v115_).length(0))], d_96_v1_)
                d_241_v134_: _dafny.Seq
                d_241_v134_ = _dafny.SeqWithoutIsStrInference([d_240_v133_])
                d_242_v135_: D1
                d_242_v135_ = D1_DC4(d_100_v2_, len(d_241_v134_), d_96_v1_)
                d_239_v132_ = (d_239_v132_).set(((d_242_v135_).cf8) - (d_101_v3_), d_238_v131_)
                d_243_v136_: _dafny.Set
                d_243_v136_ = _dafny.Set({d_221_v114_})
                d_244_v137_: _dafny.Map
                d_244_v137_ = _dafny.Map({d_101_v3_: d_243_v136_})
                d_244_v137_ = (_dafny.Map({(d_222_v115_)[default__.safeIndex(789, (d_222_v115_).length(0))]: d_243_v136_})) | (d_244_v137_)
                (d_97_globalState_).f10 = d_100_v2_
            elif True:
                d_245_v139_: _dafny.Seq
                def iife10_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in _dafny.IntegerRange(-420, 960):
                        d_246_v138_: int = compr_10_
                        if ((-420) <= (d_246_v138_)) and ((d_246_v138_) < (960)):
                            coll10_[(d_246_v138_) + (len(_dafny.Map({(d_222_v115_)[default__.safeIndex(789, (d_222_v115_).length(0))]: d_100_v2_})))] = _dafny.MultiSet([d_100_v2_, d_100_v2_, d_100_v2_, d_100_v2_, not(d_100_v2_)])
                    return _dafny.Map(coll10_)
                d_245_v139_ = _dafny.SeqWithoutIsStrInference([len(iife10_()
                ), len(d_96_v1_)])
                d_247_v140_: D1
                d_247_v140_ = D1_DC3(d_245_v139_, d_100_v2_)
                (d_97_globalState_).f10 = (d_247_v140_).cf6
                d_248_v141_: C0
                nw38_ = C0()
                nw38_.ctor__()
                d_248_v141_ = nw38_
                (d_97_globalState_).f15 = False
                d_249_v142_: _dafny.Array
                def lambda6_(d_250_v1_, d_251_v115_):
                    def lambda7_(d_252_i10_):
                        return (d_250_v1_).set(default__.safeIndex((d_251_v115_)[default__.safeIndex(789, (d_251_v115_).length(0))], len(d_250_v1_)), _dafny.CodePoint('m'))

                    return lambda7_

                init4_ = lambda6_(d_96_v1_, d_222_v115_)
                nw39_ = _dafny.Array(None, 23)
                for i0_4_ in range(nw39_.length(0)):
                    nw39_[i0_4_] = init4_(i0_4_)
                d_249_v142_ = nw39_
                d_253_v143_: _dafny.Map
                d_253_v143_ = _dafny.Map({not(d_100_v2_): d_96_v1_})
                index22_ = default__.safeIndex(242, (d_249_v142_).length(0))
                (d_249_v142_)[index22_] = ((d_253_v143_)[d_100_v2_] if (d_100_v2_) in (d_253_v143_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pojsssw")))
                index23_ = default__.safeIndex(242, (d_249_v142_).length(0))
                (d_249_v142_)[index23_] = d_96_v1_
                d_254_v144_: D1
                d_254_v144_ = D1_DC2(d_101_v3_)
                index24_ = default__.safeIndex(789, (d_222_v115_).length(0))
                (d_222_v115_)[index24_] = (d_254_v144_).cf4
            d_255_v145_: _dafny.MultiSet
            d_255_v145_ = _dafny.MultiSet([d_100_v2_])
            d_256_v146_: _dafny.Seq
            d_256_v146_ = _dafny.SeqWithoutIsStrInference([d_100_v2_])
            d_257_v147_: _dafny.Map
            d_257_v147_ = _dafny.Map({(d_221_v114_).fm3(d_100_v2_, d_100_v2_, d_255_v145_, d_256_v146_, d_97_globalState_): d_101_v3_})
            d_257_v147_ = d_257_v147_
        elif True:
            d_258_v148_: _dafny.Seq
            d_258_v148_ = _dafny.SeqWithoutIsStrInference([not(d_100_v2_)])
            (d_221_v114_).m0(d_258_v148_, default__.safeDivisionInt(d_101_v3_, d_101_v3_), d_97_globalState_)
            index25_ = default__.safeIndex(911, (d_219_v112_).length(0))
            (d_219_v112_)[index25_] = d_221_v114_
            d_259_v149_: _dafny.Map
            d_259_v149_ = _dafny.Map({d_96_v1_: d_221_v114_})
            index26_ = default__.safeIndex(911, (d_219_v112_).length(0))
            rhs8_ = True
            rhs9_ = ((d_259_v149_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucwu"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucwu"))) in (d_259_v149_) else d_221_v114_)
            lhs6_ = d_97_globalState_
            lhs7_ = d_219_v112_
            lhs8_ = default__.safeIndex(911, (d_219_v112_).length(0))
            lhs6_.f13 = rhs8_
            lhs7_[lhs8_] = rhs9_
            d_260_v150_: _dafny.Map
            d_260_v150_ = _dafny.Map({d_101_v3_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dioapv")))})
            d_260_v150_ = (d_260_v150_).set(d_101_v3_, d_101_v3_)
            d_261_v151_: str
            d_261_v151_ = _dafny.CodePoint('v')
            rhs10_ = d_261_v151_
            rhs11_ = d_95_v0_
            d_261_v151_ = rhs10_
            d_95_v0_ = rhs11_
            (d_97_globalState_).f10 = (not (d_100_v2_) or ((d_258_v148_)[default__.safeIndex((0) - (len(d_96_v1_)), len(d_258_v148_))]) if (d_101_v3_) < (d_101_v3_) else d_100_v2_)
        d_262_i11_: int
        d_262_i11_ = 0
        with _dafny.label("1"):
            while d_100_v2_:
                with _dafny.c_label("1"):
                    if (d_262_i11_) >= (100):
                        raise _dafny.Break("1")
                    d_262_i11_ = (d_262_i11_) + (1)
                    d_263_v152_: _dafny.Array
                    def lambda8_(d_264_v2_):
                        def lambda9_(d_265_i12_):
                            return (_dafny.SeqWithoutIsStrInference([d_264_v2_, d_264_v2_])) + (_dafny.SeqWithoutIsStrInference([False]))

                        return lambda9_

                    init5_ = lambda8_(d_100_v2_)
                    nw40_ = _dafny.Array(None, 17)
                    for i0_5_ in range(nw40_.length(0)):
                        nw40_[i0_5_] = init5_(i0_5_)
                    d_263_v152_ = nw40_
                    index27_ = default__.safeIndex(815, (d_263_v152_).length(0))
                    (d_263_v152_)[index27_] = _dafny.SeqWithoutIsStrInference([d_100_v2_])
                    d_266_v153_: _dafny.Seq
                    d_266_v153_ = _dafny.SeqWithoutIsStrInference([d_100_v2_, not(d_100_v2_)])
                    index28_ = default__.safeIndex(881, (d_263_v152_).length(0))
                    (d_263_v152_)[index28_] = (d_266_v153_) + (d_266_v153_)
                    d_267_v154_: _dafny.Map
                    d_267_v154_ = _dafny.Map({d_266_v153_: len(d_96_v1_)})
                    d_268_v155_: _dafny.Set
                    d_268_v155_ = _dafny.Set({False})
                    index29_ = default__.safeIndex(815, (d_263_v152_).length(0))
                    index30_ = default__.safeIndex(881, (d_263_v152_).length(0))
                    rhs12_ = d_101_v3_
                    rhs13_ = d_101_v3_
                    rhs14_ = d_266_v153_
                    rhs15_ = default__.fm4(d_267_v154_, d_97_globalState_)
                    rhs16_ = (d_266_v153_ if (_dafny.Set({True, d_100_v2_, d_100_v2_, d_100_v2_})).ispropersubset(d_268_v155_) else d_266_v153_)
                    lhs9_ = d_97_globalState_
                    lhs10_ = d_263_v152_
                    lhs11_ = default__.safeIndex(815, (d_263_v152_).length(0))
                    lhs12_ = d_97_globalState_
                    lhs13_ = d_263_v152_
                    lhs14_ = default__.safeIndex(881, (d_263_v152_).length(0))
                    d_101_v3_ = rhs12_
                    lhs9_.f0 = rhs13_
                    lhs10_[lhs11_] = rhs14_
                    lhs12_.f9 = rhs15_
                    lhs13_[lhs14_] = rhs16_
                    (d_97_globalState_).f9 = default__.safeDivisionInt((d_101_v3_ if d_100_v2_ else d_101_v3_), (0) - (default__.safeDivisionInt(d_101_v3_, d_101_v3_)))
                    (d_97_globalState_).f2 = d_101_v3_
                    d_269_v156_: D1
                    d_269_v156_ = D1_DC2(d_101_v3_)
                    d_270_v157_: D1
                    d_270_v157_ = D1_DC2((0) - ((0) - (((d_269_v156_).cf4 if d_100_v2_ else len(d_268_v155_)))))
                    source3_ = d_269_v156_
                    if source3_.is_DC3:
                        d_271___mcc_h0_ = source3_.cf5
                        d_272___mcc_h1_ = source3_.cf6
                        d_273_cf6_ = d_272___mcc_h1_
                        d_274_cf5_ = d_271___mcc_h0_
                        (d_97_globalState_).f10 = not (d_273_cf6_) or ((d_96_v1_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_275_i13_ in range(default__.abs(679))])))
                        d_276_v158_: _dafny.Seq
                        d_276_v158_ = _dafny.SeqWithoutIsStrInference([(d_96_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijprsydy"))), (d_96_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xceaadkiw"))])
                        d_277_v159_: _dafny.MultiSet
                        d_277_v159_ = _dafny.MultiSet([d_100_v2_, d_100_v2_])
                        d_278_v160_: _dafny.Map
                        d_278_v160_ = _dafny.Map({d_101_v3_: _dafny.CodePoint('b')})
                        d_279_v161_: str
                        d_279_v161_ = _dafny.CodePoint('n')
                        rhs17_ = d_276_v158_
                        rhs18_ = d_221_v114_
                        rhs19_ = (d_101_v3_) - (((d_277_v159_)[d_273_cf6_] if (d_273_cf6_) in (d_277_v159_) else d_101_v3_))
                        rhs20_ = _dafny.SeqWithoutIsStrInference([((d_278_v160_)[-218] if (-218) in (d_278_v160_) else d_279_v161_), d_279_v161_, d_279_v161_, d_279_v161_, d_279_v161_])
                        rhs21_ = (d_101_v3_) <= (d_101_v3_)
                        lhs15_ = d_97_globalState_
                        lhs16_ = d_97_globalState_
                        d_276_v158_ = rhs17_
                        d_221_v114_ = rhs18_
                        lhs15_.f0 = rhs19_
                        d_96_v1_ = rhs20_
                        lhs16_.f15 = rhs21_
                        index31_ = default__.safeIndex(333, (d_95_v0_).length(0))
                        (d_95_v0_)[index31_] = d_273_cf6_
                        d_280_v162_: _dafny.Array
                        def lambda10_(d_281_v155_):
                            def lambda11_(d_282_i14_):
                                return d_281_v155_

                            return lambda11_

                        init6_ = lambda10_(d_268_v155_)
                        nw41_ = _dafny.Array(None, 2)
                        for i0_6_ in range(nw41_.length(0)):
                            nw41_[i0_6_] = init6_(i0_6_)
                        d_280_v162_ = nw41_
                        index32_ = default__.safeIndex(821, (d_280_v162_).length(0))
                        (d_280_v162_)[index32_] = (d_268_v155_).intersection(_dafny.Set({d_273_cf6_, d_100_v2_}))
                        index33_ = default__.safeIndex(333, (d_95_v0_).length(0))
                        index34_ = default__.safeIndex(821, (d_280_v162_).length(0))
                        rhs22_ = d_100_v2_
                        rhs23_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rugsxcnc"))), len(d_274_cf5_))
                        rhs24_ = (d_268_v155_) - (default__.fm10(d_100_v2_, d_273_cf6_, d_101_v3_, d_279_v161_, d_97_globalState_))
                        rhs25_ = d_274_cf5_
                        lhs17_ = d_95_v0_
                        lhs18_ = default__.safeIndex(333, (d_95_v0_).length(0))
                        lhs19_ = d_97_globalState_
                        lhs20_ = d_280_v162_
                        lhs21_ = default__.safeIndex(821, (d_280_v162_).length(0))
                        lhs17_[lhs18_] = rhs22_
                        lhs19_.f2 = rhs23_
                        lhs20_[lhs21_] = rhs24_
                        d_274_cf5_ = rhs25_
                        index35_ = default__.safeIndex(333, (d_95_v0_).length(0))
                        rhs26_ = d_221_v114_
                        rhs27_ = d_273_cf6_
                        rhs28_ = (d_95_v0_)[default__.safeIndex(333, (d_95_v0_).length(0))]
                        lhs22_ = d_97_globalState_
                        lhs23_ = d_95_v0_
                        lhs24_ = default__.safeIndex(333, (d_95_v0_).length(0))
                        d_221_v114_ = rhs26_
                        lhs22_.f15 = rhs27_
                        lhs23_[lhs24_] = rhs28_
                    elif source3_.is_DC4:
                        d_283___mcc_h2_ = source3_.cf7
                        d_284___mcc_h3_ = source3_.cf8
                        d_285___mcc_h4_ = source3_.cf9
                        d_286_cf9_ = d_285___mcc_h4_
                        d_287_cf8_ = d_284___mcc_h3_
                        d_288_cf7_ = d_283___mcc_h2_
                        d_289_v163_: _dafny.MultiSet
                        d_289_v163_ = _dafny.MultiSet([default__.fm4(d_267_v154_, d_97_globalState_)])
                        d_290_v164_: str
                        d_290_v164_ = _dafny.CodePoint('n')
                        d_291_v165_: _dafny.Map
                        d_291_v165_ = _dafny.Map({not(d_100_v2_): d_290_v164_})
                        d_292_v166_: _dafny.MultiSet
                        d_292_v166_ = _dafny.MultiSet([d_100_v2_])
                        d_293_v167_: _dafny.Seq
                        d_293_v167_ = _dafny.SeqWithoutIsStrInference([((d_289_v163_)[d_101_v3_] if (d_101_v3_) in (d_289_v163_) else (d_292_v166_).cardinality), d_287_cf8_])
                        d_294_v168_: _dafny.Map
                        d_294_v168_ = _dafny.Map({d_100_v2_: _dafny.MultiSet(d_293_v167_)})
                        d_295_v169_: _dafny.Seq
                        d_295_v169_ = _dafny.SeqWithoutIsStrInference([len(d_291_v165_), default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfd")))), d_101_v3_), (d_287_cf8_) - (len(d_294_v168_))])
                        rhs29_ = ((_dafny.SeqWithoutIsStrInference([d_290_v164_, d_290_v164_])) + ((d_286_cf9_).set(default__.safeIndex(d_287_cf8_, len(d_286_cf9_)), d_290_v164_))) + ((d_96_v1_ if not(d_100_v2_) else d_96_v1_))
                        rhs30_ = (d_295_v169_)[default__.safeIndex(d_101_v3_, len(d_295_v169_))]
                        rhs31_ = d_287_cf8_
                        rhs32_ = (_dafny.MultiSet([d_287_cf8_, d_287_cf8_, d_101_v3_, (0) - (d_287_cf8_)])).intersection((d_289_v163_).intersection(d_289_v163_))
                        lhs25_ = d_97_globalState_
                        lhs26_ = d_97_globalState_
                        lhs27_ = d_97_globalState_
                        lhs25_.f7 = rhs29_
                        lhs26_.f9 = rhs30_
                        lhs27_.f2 = rhs31_
                        d_289_v163_ = rhs32_
                        d_296_v170_: _dafny.Array
                        nw42_ = _dafny.Array(None, 8)
                        nw42_[int(0)] = (d_96_v1_) + (d_96_v1_)
                        nw42_[int(1)] = d_96_v1_
                        nw42_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ws"))
                        nw42_[int(3)] = (d_96_v1_) + (d_286_cf9_)
                        nw42_[int(4)] = ((_dafny.SeqWithoutIsStrInference([d_290_v164_ for d_297_i15_ in range(default__.abs(741))])) + (d_286_cf9_)).set(default__.safeIndex(d_101_v3_, len((_dafny.SeqWithoutIsStrInference([d_290_v164_ for d_298_i15_ in range(default__.abs(741))])) + (d_286_cf9_))), d_290_v164_)
                        nw42_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_299_i16_ in range(default__.abs(-212))])
                        nw42_[int(6)] = d_286_cf9_
                        nw42_[int(7)] = d_96_v1_
                        d_296_v170_ = nw42_
                        index36_ = default__.safeIndex(58, (d_296_v170_).length(0))
                        (d_296_v170_)[index36_] = d_286_cf9_
                        d_300_v171_: _dafny.Seq
                        d_300_v171_ = _dafny.SeqWithoutIsStrInference([d_292_v166_])
                        pat_let_tv0_ = d_100_v2_
                        index37_ = default__.safeIndex(58, (d_296_v170_).length(0))
                        def iife11_(_pat_let0_0):
                            def iife12_(d_301_dt__update__tmp_h0_):
                                def iife13_(_pat_let1_0):
                                    def iife14_(d_302_dt__update_hcf2_h0_):
                                        return D0_DC1((d_301_dt__update__tmp_h0_).cf1, d_302_dt__update_hcf2_h0_, (d_301_dt__update__tmp_h0_).cf3)
                                    return iife14_(_pat_let1_0)
                                return iife13_(pat_let_tv0_)
                            return iife12_(_pat_let0_0)
                        (d_296_v170_)[index37_] = ((d_286_cf9_) + (d_286_cf9_) if (d_221_v114_).fm3(d_288_cf7_, d_100_v2_, (d_300_v171_)[default__.safeIndex(557, len(d_300_v171_))], d_266_v153_, d_97_globalState_) else default__.fm5((iife11_(D0_DC1((d_221_v114_).fm3(True, d_100_v2_, d_292_v166_, d_266_v153_, d_97_globalState_), d_100_v2_, d_100_v2_))).cf2, (d_221_v114_).fm3(not(d_288_cf7_), d_100_v2_, d_292_v166_, (d_263_v152_)[default__.safeIndex(815, (d_263_v152_).length(0))], d_97_globalState_), d_97_globalState_))
                        d_303_v172_: _dafny.Seq
                        d_303_v172_ = _dafny.SeqWithoutIsStrInference([d_266_v153_])
                        d_303_v172_ = (d_303_v172_) + (d_303_v172_)
                        d_290_v164_ = d_290_v164_
                    elif source3_.is_DC2:
                        d_304___mcc_h5_ = source3_.cf4
                        d_305_cf4_ = d_304___mcc_h5_
                        d_306_v173_: C0
                        nw43_ = C0()
                        nw43_.ctor__()
                        d_306_v173_ = nw43_
                        d_307_v174_: _dafny.Array
                        nw44_ = _dafny.Array(_dafny.Set({}), 28)
                        d_307_v174_ = nw44_
                        d_308_v175_: _dafny.Map
                        d_308_v175_ = _dafny.Map({671: d_305_cf4_})
                        d_309_v176_: _dafny.Set
                        d_309_v176_ = _dafny.Set({d_308_v175_, d_308_v175_, _dafny.Map({d_305_cf4_: default__.fm4(d_267_v154_, d_97_globalState_)})})
                        index38_ = default__.safeIndex(947, (d_307_v174_).length(0))
                        (d_307_v174_)[index38_] = d_309_v176_
                        index39_ = default__.safeIndex(947, (d_307_v174_).length(0))
                        (d_307_v174_)[index39_] = d_309_v176_
                        d_310_v177_: D1
                        d_310_v177_ = D1_DC4((d_268_v155_).isdisjoint(d_268_v155_), d_101_v3_, d_96_v1_)
                        d_310_v177_ = d_310_v177_
                        (d_97_globalState_).f9 = default__.fm4(d_267_v154_, d_97_globalState_)
                    elif True:
                        d_311___mcc_h6_ = source3_.cf10
                        d_312_cf10_ = d_311___mcc_h6_
                        (d_221_v114_).m0(_dafny.SeqWithoutIsStrInference([d_100_v2_]), default__.safeModuloInt(d_101_v3_, d_101_v3_), d_97_globalState_)
                        d_313_v178_: _dafny.Seq
                        d_313_v178_ = _dafny.SeqWithoutIsStrInference([d_221_v114_, d_221_v114_])
                        d_314_v179_: _dafny.Map
                        d_314_v179_ = _dafny.Map({d_221_v114_: d_313_v178_})
                        d_314_v179_ = (d_314_v179_).set(d_221_v114_, (d_313_v178_ if d_100_v2_ else d_313_v178_))
                        d_315_v180_: str
                        d_315_v180_ = _dafny.CodePoint('b')
                        d_316_v181_: D0
                        d_316_v181_ = D0_DC0(d_315_v180_)
                        d_317_v182_: _dafny.Seq
                        d_317_v182_ = _dafny.SeqWithoutIsStrInference([D0_DC0(d_315_v180_), d_316_v181_, d_316_v181_])
                        d_318_v183_: _dafny.Map
                        d_318_v183_ = _dafny.Map({d_100_v2_: _dafny.SeqWithoutIsStrInference([d_101_v3_])})
                        d_319_v184_: _dafny.Seq
                        d_319_v184_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgeiw")))])
                        d_320_v185_: C0
                        d_321_v186_: _dafny.Array
                        d_322_v187_: int
                        d_323_v188_: int
                        out16_: C0
                        out17_: _dafny.Array
                        out18_: int
                        out19_: int
                        def iife15_(_pat_let2_0):
                            def iife16_(d_324_dt__update__tmp_h1_):
                                def iife17_(_pat_let3_0):
                                    def iife18_(d_325_dt__update_hcf0_h0_):
                                        return D0_DC0(d_325_dt__update_hcf0_h0_)
                                    return iife18_(_pat_let3_0)
                                return iife17_(_dafny.CodePoint('m'))
                            return iife16_(_pat_let2_0)
                        out16_, out17_, out18_, out19_ = default__.m1(d_315_v180_, (_dafny.SeqWithoutIsStrInference([D0_DC0(d_315_v180_), iife15_(D0_DC0(d_315_v180_)), D0_DC0(_dafny.CodePoint('o')), D0_DC0(d_315_v180_)]) if d_100_v2_ else d_317_v182_), d_100_v2_, ((d_318_v183_)[d_100_v2_] if (d_100_v2_) in (d_318_v183_) else d_319_v184_), d_97_globalState_)
                        d_320_v185_ = out16_
                        d_321_v186_ = out17_
                        d_322_v187_ = out18_
                        d_323_v188_ = out19_
                        (d_97_globalState_).f2 = (389) * (default__.fm4((d_267_v154_).set(d_266_v153_, d_101_v3_), d_97_globalState_))
                    pass
            pass
        d_326_v189_: str
        d_326_v189_ = _dafny.CodePoint('p')
        d_327_v190_: _dafny.Map
        d_327_v190_ = _dafny.Map({d_326_v189_: d_101_v3_})
        d_328_v191_: _dafny.Map
        d_328_v191_ = _dafny.Map({False: d_100_v2_})
        d_329_v192_: _dafny.MultiSet
        d_329_v192_ = _dafny.MultiSet([d_100_v2_, d_100_v2_])
        d_330_v193_: _dafny.Seq
        d_330_v193_ = _dafny.SeqWithoutIsStrInference([d_100_v2_])
        d_331_v194_: _dafny.Seq
        d_331_v194_ = _dafny.SeqWithoutIsStrInference([(d_221_v114_).fm3(((d_328_v191_)[d_100_v2_] if (d_100_v2_) in (d_328_v191_) else d_100_v2_), d_100_v2_, d_329_v192_, d_330_v193_, d_97_globalState_), False])
        d_332_v195_: _dafny.Seq
        d_332_v195_ = _dafny.SeqWithoutIsStrInference([d_327_v190_, default__.fm11(_dafny.MultiSet(d_330_v193_), d_101_v3_, d_100_v2_, d_97_globalState_)])
        d_333_i17_: int
        d_333_i17_ = 0
        with _dafny.label("2"):
            while (d_221_v114_).fm3((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipj"))) < (d_96_v1_), (((d_332_v195_).set(default__.safeIndex(d_101_v3_, len(d_332_v195_)), d_327_v190_)).set(default__.safeIndex(d_101_v3_, len((d_332_v195_).set(default__.safeIndex(d_101_v3_, len(d_332_v195_)), d_327_v190_))), d_327_v190_)) == ((_dafny.SeqWithoutIsStrInference([d_327_v190_])).set(default__.safeIndex(d_101_v3_, len(_dafny.SeqWithoutIsStrInference([d_327_v190_]))), default__.fm11(_dafny.MultiSet([d_100_v2_]), d_101_v3_, d_100_v2_, d_97_globalState_))), (d_329_v192_).intersection(_dafny.MultiSet([d_100_v2_, d_100_v2_])), d_331_v194_, d_97_globalState_):
                with _dafny.c_label("2"):
                    if (d_333_i17_) >= (100):
                        raise _dafny.Break("2")
                    d_333_i17_ = (d_333_i17_) + (1)
                    d_334_v196_: _dafny.Array
                    def lambda12_(d_335_i18_):
                        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cixtikg"))

                    init7_ = lambda12_
                    nw45_ = _dafny.Array(None, 27)
                    for i0_7_ in range(nw45_.length(0)):
                        nw45_[i0_7_] = init7_(i0_7_)
                    d_334_v196_ = nw45_
                    index40_ = default__.safeIndex(676, (d_334_v196_).length(0))
                    (d_334_v196_)[index40_] = d_96_v1_
                    index41_ = default__.safeIndex(676, (d_334_v196_).length(0))
                    (d_334_v196_)[index41_] = d_96_v1_
                    d_336_v197_: _dafny.Array
                    nw46_ = _dafny.Array(_dafny.Set({}), 14)
                    d_336_v197_ = nw46_
                    d_337_v198_: _dafny.Array
                    nw47_ = _dafny.Array(None, 22)
                    nw47_[int(0)] = d_336_v197_
                    nw47_[int(1)] = d_336_v197_
                    nw47_[int(2)] = d_336_v197_
                    nw47_[int(3)] = d_336_v197_
                    nw47_[int(4)] = d_336_v197_
                    nw47_[int(5)] = d_336_v197_
                    nw47_[int(6)] = d_336_v197_
                    nw47_[int(7)] = d_336_v197_
                    nw47_[int(8)] = d_336_v197_
                    nw47_[int(9)] = d_336_v197_
                    nw47_[int(10)] = d_336_v197_
                    nw47_[int(11)] = d_336_v197_
                    nw47_[int(12)] = d_336_v197_
                    nw47_[int(13)] = d_336_v197_
                    nw47_[int(14)] = d_336_v197_
                    nw47_[int(15)] = d_336_v197_
                    nw47_[int(16)] = d_336_v197_
                    nw47_[int(17)] = d_336_v197_
                    nw47_[int(18)] = (d_336_v197_ if d_100_v2_ else d_336_v197_)
                    nw47_[int(19)] = d_336_v197_
                    nw47_[int(20)] = d_336_v197_
                    nw47_[int(21)] = d_336_v197_
                    d_337_v198_ = nw47_
                    index42_ = default__.safeIndex(331, (d_337_v198_).length(0))
                    (d_337_v198_)[index42_] = d_336_v197_
                    index43_ = default__.safeIndex(331, (d_337_v198_).length(0))
                    (d_337_v198_)[index43_] = d_336_v197_
                    d_95_v0_ = d_95_v0_
                    d_338_v199_: C0
                    nw48_ = C0()
                    nw48_.ctor__()
                    d_338_v199_ = nw48_
                    pass
            pass
        d_339_v200_: _dafny.Seq
        d_339_v200_ = _dafny.SeqWithoutIsStrInference([d_101_v3_])
        (d_221_v114_).m0(d_331_v194_, (d_339_v200_)[default__.safeIndex(len(d_96_v1_), len(d_339_v200_))], d_97_globalState_)
        d_331_v194_ = (_dafny.SeqWithoutIsStrInference([d_100_v2_, d_100_v2_, d_100_v2_, d_100_v2_])) + ((_dafny.SeqWithoutIsStrInference([d_100_v2_, d_100_v2_, d_100_v2_]) if d_100_v2_ else d_331_v194_))
        d_340_i19_: int
        d_340_i19_ = 0
        with _dafny.label("3"):
            while (not((d_221_v114_).fm3(d_100_v2_, d_100_v2_, _dafny.MultiSet([d_100_v2_, d_100_v2_, d_100_v2_]), d_330_v193_, d_97_globalState_))) and (d_100_v2_):
                with _dafny.c_label("3"):
                    if (d_340_i19_) >= (100):
                        raise _dafny.Break("3")
                    d_340_i19_ = (d_340_i19_) + (1)
                    (d_97_globalState_).f10 = d_100_v2_
                    d_100_v2_ = d_100_v2_
                    (d_97_globalState_).f15 = not(d_100_v2_)
                    (d_97_globalState_).f15 = ((d_329_v192_) | (d_329_v192_)).ispropersubset((d_329_v192_) - (d_329_v192_))
                    pass
            pass
        d_341_i20_: int
        d_341_i20_ = 0
        with _dafny.label("4"):
            while (d_100_v2_) == (d_100_v2_):
                with _dafny.c_label("4"):
                    if (d_341_i20_) >= (100):
                        raise _dafny.Break("4")
                    d_341_i20_ = (d_341_i20_) + (1)
                    d_342_v201_: _dafny.Map
                    d_342_v201_ = _dafny.Map({d_100_v2_: d_329_v192_})
                    (d_97_globalState_).f15 = (((d_342_v201_)[d_100_v2_] if (d_100_v2_) in (d_342_v201_) else _dafny.MultiSet([d_100_v2_, d_100_v2_]))).issubset(d_329_v192_)
                    (d_97_globalState_).f15 = (d_101_v3_) <= (d_101_v3_)
                    d_343_v202_: _dafny.Array
                    nw49_ = _dafny.Array(int(0), 12)
                    d_343_v202_ = nw49_
                    d_344_v203_: _dafny.Map
                    d_344_v203_ = _dafny.Map({d_100_v2_: d_101_v3_})
                    d_345_v204_: _dafny.Set
                    d_345_v204_ = _dafny.Set({(d_344_v203_) | (_dafny.Map({d_100_v2_: d_101_v3_})), (d_344_v203_).set(d_100_v2_, d_101_v3_), d_344_v203_})
                    d_346_v205_: _dafny.Map
                    d_346_v205_ = _dafny.Map({d_96_v1_: d_96_v1_})
                    rhs33_ = d_221_v114_
                    rhs34_ = d_343_v202_
                    rhs35_ = (_dafny.SeqWithoutIsStrInference([d_326_v189_ for d_347_i21_ in range(default__.abs(-480))])) not in ((d_346_v205_).set(d_96_v1_, default__.fm5(d_100_v2_, d_100_v2_, d_97_globalState_)))
                    rhs36_ = d_345_v204_
                    lhs28_ = d_97_globalState_
                    d_221_v114_ = rhs33_
                    d_343_v202_ = rhs34_
                    lhs28_.f10 = rhs35_
                    d_345_v204_ = rhs36_
                    if (d_101_v3_) <= (d_101_v3_):
                        (d_97_globalState_).f9 = len((d_96_v1_).set(default__.safeIndex(793, len(d_96_v1_)), d_326_v189_))
                        d_348_v206_: _dafny.Array
                        nw50_ = _dafny.Array(_dafny.Array(None, 0), 5)
                        d_348_v206_ = nw50_
                        index44_ = default__.safeIndex(989, (d_348_v206_).length(0))
                        (d_348_v206_)[index44_] = d_343_v202_
                        index45_ = default__.safeIndex(989, (d_348_v206_).length(0))
                        (d_348_v206_)[index45_] = d_343_v202_
                        d_349_v207_: _dafny.Map
                        d_349_v207_ = _dafny.Map({72: d_101_v3_})
                        d_350_v208_: _dafny.Seq
                        d_350_v208_ = _dafny.SeqWithoutIsStrInference([d_349_v207_])
                        d_351_v209_: _dafny.MultiSet
                        d_351_v209_ = _dafny.MultiSet([d_101_v3_, (len((d_350_v208_)[default__.safeIndex(d_101_v3_, len(d_350_v208_))])) + (d_101_v3_), (d_101_v3_) + (d_101_v3_), default__.safeModuloInt(d_101_v3_, d_101_v3_), d_101_v3_])
                        d_352_v211_: _dafny.Map
                        d_352_v211_ = _dafny.Map({(d_351_v209_).cardinality: (d_221_v114_).fm3(d_100_v2_, d_100_v2_, _dafny.MultiSet([d_100_v2_, d_100_v2_]), d_330_v193_, d_97_globalState_)})
                        d_353_v212_: _dafny.Map
                        d_353_v212_ = _dafny.Map({d_352_v211_: (D1_DC4(d_100_v2_, d_101_v3_, d_96_v1_)).cf7})
                        d_354_v215_: _dafny.Set
                        d_354_v215_ = _dafny.Set({d_352_v211_})
                        d_355_v216_: _dafny.Seq
                        d_355_v216_ = _dafny.SeqWithoutIsStrInference([d_354_v215_])
                        d_356_v217_: _dafny.Seq
                        def iife19_():
                            coll11_ = _dafny.Map()
                            compr_11_: _dafny.Map
                            for compr_11_ in ((d_355_v216_)[default__.safeIndex(d_101_v3_, len(d_355_v216_))]).Elements:
                                d_357_v214_: _dafny.Map = compr_11_
                                if (d_357_v214_) in ((d_355_v216_)[default__.safeIndex(d_101_v3_, len(d_355_v216_))]):
                                    coll11_[d_357_v214_] = d_100_v2_
                            return _dafny.Map(coll11_)
                        d_356_v217_ = _dafny.SeqWithoutIsStrInference([iife19_()
                        ])
                        d_358_v218_: _dafny.Set
                        d_358_v218_ = _dafny.Set({d_100_v2_})
                        d_359_v219_: D4
                        d_359_v219_ = D4_DC8(d_353_v212_)
                        d_360_v220_: _dafny.Array
                        nw51_ = _dafny.Array(None, 20)
                        def iife20_():
                            coll12_ = _dafny.Map()
                            compr_12_: _dafny.Map
                            for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_101_v3_: d_100_v2_}) for d_361_i22_ in range(default__.abs(-660))])).Elements:
                                d_362_v210_: _dafny.Map = compr_12_
                                if (d_362_v210_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_101_v3_: d_100_v2_}) for d_361_i22_ in range(default__.abs(-660))])):
                                    coll12_[d_362_v210_] = d_100_v2_
                            return _dafny.Map(coll12_)
                        nw51_[int(0)] = (iife20_()
                        ) | (d_353_v212_)
                        nw51_[int(1)] = d_353_v212_
                        nw51_[int(2)] = d_353_v212_
                        nw51_[int(3)] = d_353_v212_
                        nw51_[int(4)] = d_353_v212_
                        nw51_[int(5)] = d_353_v212_
                        def iife21_():
                            coll13_ = _dafny.Map()
                            compr_13_: int
                            for compr_13_ in (d_352_v211_).keys.Elements:
                                d_363_v213_: int = compr_13_
                                if (d_363_v213_) in (d_352_v211_):
                                    coll13_[(d_363_v213_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))] = True
                            return _dafny.Map(coll13_)
                        nw51_[int(6)] = (d_353_v212_).set(iife21_()
                        , d_100_v2_)
                        nw51_[int(7)] = d_353_v212_
                        nw51_[int(8)] = d_353_v212_
                        nw51_[int(9)] = (d_353_v212_).set(d_352_v211_, d_100_v2_)
                        nw51_[int(10)] = (d_353_v212_).set(d_352_v211_, not(d_100_v2_))
                        nw51_[int(11)] = ((d_356_v217_)[default__.safeIndex(d_101_v3_, len(d_356_v217_))]) | (_dafny.Map({(_dafny.Map({d_101_v3_: d_100_v2_})).set(d_101_v3_, d_100_v2_): (d_221_v114_).fm3(d_100_v2_, d_100_v2_, _dafny.MultiSet([d_100_v2_]), d_331_v194_, d_97_globalState_)}))
                        nw51_[int(12)] = d_353_v212_
                        nw51_[int(13)] = d_353_v212_
                        nw51_[int(14)] = (d_353_v212_) | (d_353_v212_)
                        nw51_[int(15)] = d_353_v212_
                        nw51_[int(16)] = _dafny.Map({_dafny.Map({(0) - (len(d_358_v218_)): False}): ((d_352_v211_)[len(d_328_v191_)] if (len(d_328_v191_)) in (d_352_v211_) else True)})
                        nw51_[int(17)] = d_353_v212_
                        nw51_[int(18)] = (d_359_v219_).cf13
                        nw51_[int(19)] = d_353_v212_
                        d_360_v220_ = nw51_
                        index46_ = default__.safeIndex(351, (d_360_v220_).length(0))
                        def iife22_():
                            coll14_ = _dafny.Map()
                            compr_14_: int
                            for compr_14_ in _dafny.IntegerRange(-640, 23):
                                d_364_v221_: int = compr_14_
                                if ((-640) <= (d_364_v221_)) and ((d_364_v221_) < (23)):
                                    coll14_[(d_364_v221_) - (len(d_344_v203_))] = d_100_v2_
                            return _dafny.Map(coll14_)
                        (d_360_v220_)[index46_] = (_dafny.Map({d_352_v211_: d_100_v2_})) | ((d_353_v212_).set(iife22_()
                        , d_100_v2_))
                        index47_ = default__.safeIndex(351, (d_360_v220_).length(0))
                        def iife23_():
                            coll15_ = _dafny.Map()
                            compr_15_: _dafny.Map
                            for compr_15_ in (_dafny.Set({d_352_v211_})).Elements:
                                d_365_v222_: _dafny.Map = compr_15_
                                if (d_365_v222_) in (_dafny.Set({d_352_v211_})):
                                    coll15_[d_365_v222_] = False
                            return _dafny.Map(coll15_)
                        rhs37_ = _dafny.MultiSet([(d_101_v3_) - ((0) - (((d_344_v203_)[d_100_v2_] if (d_100_v2_) in (d_344_v203_) else d_101_v3_)))])
                        rhs38_ = d_221_v114_
                        rhs39_ = (d_353_v212_ if not(d_100_v2_) else ((d_356_v217_)[default__.safeIndex(d_101_v3_, len(d_356_v217_))]) | (iife23_()
                        ))
                        rhs40_ = d_221_v114_
                        lhs29_ = d_360_v220_
                        lhs30_ = default__.safeIndex(351, (d_360_v220_).length(0))
                        d_351_v209_ = rhs37_
                        d_221_v114_ = rhs38_
                        lhs29_[lhs30_] = rhs39_
                        d_221_v114_ = rhs40_
                        d_366_v223_: _dafny.Array
                        nw52_ = _dafny.Array(D1.default()(), 3)
                        d_366_v223_ = nw52_
                        index48_ = default__.safeIndex(972, (d_95_v0_).length(0))
                        (d_95_v0_)[index48_] = d_100_v2_
                        index49_ = default__.safeIndex(972, (d_95_v0_).length(0))
                        rhs41_ = (0) - ((d_101_v3_) * (d_101_v3_))
                        rhs42_ = _dafny.CodePoint('v')
                        rhs43_ = d_366_v223_
                        rhs44_ = (d_100_v2_) and ((d_330_v193_)[default__.safeIndex(d_101_v3_, len(d_330_v193_))])
                        lhs31_ = d_97_globalState_
                        lhs32_ = d_95_v0_
                        lhs33_ = default__.safeIndex(972, (d_95_v0_).length(0))
                        lhs31_.f2 = rhs41_
                        d_326_v189_ = rhs42_
                        d_366_v223_ = rhs43_
                        lhs32_[lhs33_] = rhs44_
                        d_367_v224_: C0
                        d_368_v225_: _dafny.Array
                        d_369_v226_: int
                        d_370_v227_: int
                        out20_: C0
                        out21_: _dafny.Array
                        out22_: int
                        out23_: int
                        out20_, out21_, out22_, out23_ = default__.m1(d_326_v189_, _dafny.SeqWithoutIsStrInference([D0_DC0(d_326_v189_) for d_371_i23_ in range(default__.abs(-33))]), d_100_v2_, d_339_v200_, d_97_globalState_)
                        d_367_v224_ = out20_
                        d_368_v225_ = out21_
                        d_369_v226_ = out22_
                        d_370_v227_ = out23_
                    elif True:
                        (d_221_v114_).m0((_dafny.SeqWithoutIsStrInference([d_100_v2_])) + (d_331_v194_), len(d_339_v200_), d_97_globalState_)
                        d_221_v114_ = d_221_v114_
                        d_372_v228_: C0
                        nw53_ = C0()
                        nw53_.ctor__()
                        d_372_v228_ = nw53_
                        d_373_v229_: _dafny.Map
                        d_373_v229_ = _dafny.Map({d_101_v3_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emrwvbq"))})
                        d_374_v230_: _dafny.MultiSet
                        d_374_v230_ = _dafny.MultiSet([len(d_339_v200_)])
                        d_375_v231_: _dafny.MultiSet
                        d_375_v231_ = _dafny.MultiSet([d_374_v230_, d_374_v230_])
                        (d_97_globalState_).f7 = ((((d_373_v229_)[(d_375_v231_).cardinality] if ((d_375_v231_).cardinality) in (d_373_v229_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skwwff")))) + (default__.fm5(d_100_v2_, d_100_v2_, d_97_globalState_))) + (d_96_v1_)
                        (d_97_globalState_).f2 = default__.safeModuloInt(d_101_v3_, (d_101_v3_) * (d_101_v3_))
                    pass
            pass
        d_101_v3_ = d_101_v3_
        d_376_v232_: _dafny.Array
        def lambda13_(d_377_v3_):
            def lambda14_(d_378_i24_):
                return default__.safeModuloInt(d_378_i24_, (0) - (d_377_v3_))

            return lambda14_

        init8_ = lambda13_(d_101_v3_)
        nw54_ = _dafny.Array(None, 20)
        for i0_8_ in range(nw54_.length(0)):
            nw54_[i0_8_] = init8_(i0_8_)
        d_376_v232_ = nw54_
        d_376_v232_ = d_376_v232_
        rhs45_ = d_100_v2_
        rhs46_ = (d_96_v1_) + ((d_96_v1_).set(default__.safeIndex(d_101_v3_, len(d_96_v1_)), d_326_v189_))
        lhs34_ = d_97_globalState_
        d_100_v2_ = rhs45_
        lhs34_.f7 = rhs46_
        d_100_v2_ = (d_329_v192_).issubset(_dafny.MultiSet([d_100_v2_]))
        d_376_v232_ = d_376_v232_
        _dafny.print(_dafny.string_of((d_95_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_95_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_96_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_97_globalState_).f4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_97_globalState_).f4)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_97_globalState_).f6) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_97_globalState_.f7).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_97_globalState_).f14).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_97_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_100_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_101_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_220_v113_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_262_i11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_326_v189_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v190_) == (_dafny.Map({_dafny.CodePoint('p'): 890}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_328_v191_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_329_v192_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_330_v193_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_331_v194_) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_332_v195_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('p'): 890}), _dafny.Map({_dafny.CodePoint('u'): 2, _dafny.CodePoint('v'): 46, _dafny.CodePoint('k'): 1, _dafny.CodePoint('n'): 1})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_333_i17_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v200_) == (_dafny.SeqWithoutIsStrInference([890]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_340_i19_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_341_i20_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_376_v232_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
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
        return lambda: D1_DC3(_dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {self.cf9.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC7(D3, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC9(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)

class D4_DC9(D4, NamedTuple('DC9', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC12(D5, NamedTuple('DC12', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f2: int = int(0)
        self.f7: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f9: int = int(0)
        self.f10: bool = False
        self.f13: bool = False
        self.f15: bool = False
        self._f1: str = _dafny.CodePoint('D')
        self._f3: int = int(0)
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: bool = False
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f8: str = _dafny.CodePoint('D')
        self._f11: str = _dafny.CodePoint('D')
        self._f12: int = int(0)
        self._f14: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self).f15 = f15

    @property
    def f1(self):
        return self._f1
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    @property
    def f8(self):
        return self._f8
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    @property
    def f14(self):
        return self._f14

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm3(self, p0, p1, p2, p3, globalState):
        return True

    def m0(self, p0, p1, globalState):
        d_379_v0_: D1
        d_379_v0_ = D1_DC2(p1)
        d_380_v1_: bool
        d_380_v1_ = True
        d_381_v2_: _dafny.Map
        d_381_v2_ = _dafny.Map({d_380_v1_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snasecdf")))})
        d_382_v3_: _dafny.Map
        d_382_v3_ = _dafny.Map({385: _dafny.Map({False: p1})})
        d_383_v4_: _dafny.Seq
        d_383_v4_ = _dafny.SeqWithoutIsStrInference([d_381_v2_, ((d_382_v3_)[len(p0)] if (len(p0)) in (d_382_v3_) else d_381_v2_)])
        d_384_v5_: _dafny.Array
        nw55_ = _dafny.Array(None, 14)
        nw55_[int(0)] = p1
        nw55_[int(1)] = p1
        nw55_[int(2)] = p1
        nw55_[int(3)] = (0) - ((d_379_v0_).cf4)
        nw55_[int(4)] = len(d_383_v4_)
        nw55_[int(5)] = p1
        nw55_[int(6)] = (d_379_v0_).cf4
        nw55_[int(7)] = p1
        nw55_[int(8)] = p1
        nw55_[int(9)] = p1
        nw55_[int(10)] = p1
        nw55_[int(11)] = (len(p0)) * (p1)
        nw55_[int(12)] = p1
        nw55_[int(13)] = default__.safeModuloInt((0) - (p1), p1)
        d_384_v5_ = nw55_
        index50_ = default__.safeIndex(192, (d_384_v5_).length(0))
        (d_384_v5_)[index50_] = (p1) * (p1)
        d_385_v6_: _dafny.Map
        d_385_v6_ = _dafny.Map({p0: p1})
        index51_ = default__.safeIndex(192, (d_384_v5_).length(0))
        (d_384_v5_)[index51_] = default__.fm4(d_385_v6_, globalState)
        d_386_v7_: _dafny.Map
        d_386_v7_ = _dafny.Map({(0) - ((d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))]): d_380_v1_})
        d_387_v8_: _dafny.Set
        d_387_v8_ = _dafny.Set({d_386_v7_})
        d_388_v9_: _dafny.Map
        d_388_v9_ = _dafny.Map({len(d_387_v8_): p1})
        d_389_v10_: _dafny.Set
        d_389_v10_ = _dafny.Set({d_380_v1_})
        d_390_v11_: _dafny.Seq
        d_390_v11_ = _dafny.SeqWithoutIsStrInference([d_389_v10_, d_389_v10_])
        d_388_v9_ = (d_388_v9_).set(default__.safeModuloInt((d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))), (0) - (len(d_390_v11_)))
        d_391_v12_: _dafny.Map
        d_391_v12_ = _dafny.Map({d_380_v1_: d_384_v5_})
        index52_ = default__.safeIndex(192, (d_384_v5_).length(0))
        (d_384_v5_)[index52_] = (len((d_391_v12_) | (d_391_v12_))) - (800)
        d_392_v13_: _dafny.Array
        nw56_ = _dafny.Array(False, 15)
        d_392_v13_ = nw56_
        index53_ = default__.safeIndex(904, (d_392_v13_).length(0))
        (d_392_v13_)[index53_] = True
        index54_ = default__.safeIndex(904, (d_392_v13_).length(0))
        (d_392_v13_)[index54_] = d_380_v1_
        hi1_ = (d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))]
        for d_393_i0_ in range(p1, hi1_):
            (globalState).f0 = default__.safeModuloInt((p1 if d_380_v1_ else (d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))]), (0) - ((d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))]))
            d_394_v14_: _dafny.Seq
            d_394_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cshphu"))
            index55_ = default__.safeIndex(904, (d_392_v13_).length(0))
            index56_ = default__.safeIndex(192, (d_384_v5_).length(0))
            rhs47_ = ((d_388_v9_)[(len(d_389_v10_)) * ((d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))])] if ((len(d_389_v10_)) * ((d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))])) in (d_388_v9_) else default__.safeModuloInt(d_393_i0_, (d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))]))
            rhs48_ = (d_394_v14_) == (d_394_v14_)
            rhs49_ = 686
            lhs35_ = globalState
            lhs36_ = d_392_v13_
            lhs37_ = default__.safeIndex(904, (d_392_v13_).length(0))
            lhs38_ = d_384_v5_
            lhs39_ = default__.safeIndex(192, (d_384_v5_).length(0))
            lhs35_.f9 = rhs47_
            lhs36_[lhs37_] = rhs48_
            lhs38_[lhs39_] = rhs49_
            (globalState).f13 = False
            (globalState).f0 = -981
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_384_v5_).length(0)):
            d_395_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_395_i1_)) and ((d_395_i1_) < ((d_384_v5_).length(0)))):
                _ingredientsd_0.append((d_384_v5_, int(d_395_i1_), (d_395_i1_) * (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_396_i2_ in range(default__.abs(381))])), (d_384_v5_)[default__.safeIndex(192, (d_384_v5_).length(0))]))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]

