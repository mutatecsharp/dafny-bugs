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
    def fm0(p0, p1, p2, globalState):
        if (len(_dafny.SeqWithoutIsStrInference([not(not(False))]))) < (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_0_i0_ in range(default__.abs(-917))]))):
            return (_dafny.CodePoint('h'))
        elif True:
            return _dafny.CodePoint('c')

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return True

    @staticmethod
    def fm2(p0, globalState):
        return ((len(_dafny.Map({_dafny.CodePoint('x'): True}))) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([490]))).cardinality]))])))) * (len((_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtycevgkj"))})) | (_dafny.Map({not(not(True)): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqmpa")))}))))

    @staticmethod
    def fm3(globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({229: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1_i0_ in range(default__.abs(838))]))})), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2_i1_ in range(default__.abs(294))])))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auo"))) for d_3_i2_ in range(default__.abs(720))]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([318 for d_4_i3_ in range(default__.abs(-51))])), len(_dafny.SeqWithoutIsStrInference([True, not(not(not(True))), True]))])) + (_dafny.SeqWithoutIsStrInference([245 for d_5_i4_ in range(default__.abs(497))])))

    @staticmethod
    def fm7(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjwgaoud")))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        def iife0_():
            def iife3_():
                def iife4_():
                    coll4_ = _dafny.Map()
                    compr_4_: int
                    for compr_4_ in (_dafny.MultiSet([-312, (_dafny.MultiSet([220])).cardinality])).Elements:
                        d_7_v1_: int = compr_4_
                        if (d_7_v1_) in (_dafny.MultiSet([-312, (_dafny.MultiSet([220])).cardinality])):
                            coll4_[default__.safeDivisionInt(d_7_v1_, -651)] = 555
                    return _dafny.Map(coll4_)
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(46, 483):
                    d_6_v0_: int = compr_3_
                    if ((46) <= (d_6_v0_)) and ((d_6_v0_) < (483)):
                        coll3_[default__.safeModuloInt(d_6_v0_, len(iife4_()
                        ))] = 899
                return _dafny.Map(coll3_)
            coll0_ = _dafny.Set()
            def iife1_():
                def iife2_():
                    coll2_ = _dafny.Map()
                    compr_2_: int
                    for compr_2_ in (_dafny.MultiSet([-312, (_dafny.MultiSet([220])).cardinality])).Elements:
                        d_7_v1_: int = compr_2_
                        if (d_7_v1_) in (_dafny.MultiSet([-312, (_dafny.MultiSet([220])).cardinality])):
                            coll2_[default__.safeDivisionInt(d_7_v1_, -651)] = 555
                    return _dafny.Map(coll2_)
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(46, 483):
                    d_6_v0_: int = compr_1_
                    if ((46) <= (d_6_v0_)) and ((d_6_v0_) < (483)):
                        coll1_[default__.safeModuloInt(d_6_v0_, len(iife2_()
                        ))] = 899
                return _dafny.Map(coll1_)
            compr_0_: int
            for compr_0_ in (iife1_()
            ).keys.Elements:
                d_8_v2_: int = compr_0_
                if (d_8_v2_) in (iife3_()
                ):
                    coll0_ = coll0_.union(_dafny.Set([(d_8_v2_) * (len(_dafny.Map({len(_dafny.Map({not(False): len(_dafny.Set({899, 703, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llm")))}))})): -706})))]))
            return _dafny.Set(coll0_)
        source0_ = D7_DC14(D7_DC14(D7_DC14(D7_DC14(D7_DC12(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([989, len(iife0_()
)])]))))))
        if source0_.is_DC13:
            d_9___mcc_h0_ = source0_.cf17
            d_10_cf17_ = d_9___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_11_i0_ in range(default__.abs(37))])
        elif source0_.is_DC12:
            d_12___mcc_h1_ = source0_.cf16
            d_13_cf16_ = d_12___mcc_h1_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
        elif True:
            d_14___mcc_h2_ = source0_.cf18
            d_15_cf18_ = d_14___mcc_h2_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvccdvgrp"))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return _dafny.Set({(_dafny.Map({916: True})) | (_dafny.Map({len(_dafny.Map({False: -611})): False})), _dafny.Map({991: not(True)})})

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        source1_ = (D6_DC11(-673) if not(True) else D6_DC11(len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqh")))}))))
        if source1_.is_DC11:
            d_16___mcc_h0_ = source1_.cf15
            d_17_cf15_ = d_16___mcc_h0_
            return _dafny.CodePoint('g')
        elif True:
            d_18___mcc_h1_ = source1_.cf14
            d_19_cf14_ = d_18___mcc_h1_
            return _dafny.CodePoint('r')

    @staticmethod
    def fm11(globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: str
            for compr_5_ in (_dafny.Map({_dafny.CodePoint('b'): -514})).keys.Elements:
                d_20_v0_: str = compr_5_
                if (d_20_v0_) in (_dafny.Map({_dafny.CodePoint('b'): -514})):
                    coll5_ = coll5_.union(_dafny.Set([d_20_v0_]))
            return _dafny.Set(coll5_)
        return ((_dafny.Map({False: _dafny.Set({_dafny.CodePoint('i')})})) | (_dafny.Map({False: _dafny.Set({_dafny.CodePoint('g')})}))) | (_dafny.Map({False: iife5_()
        }))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        def iife6_():
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in _dafny.IntegerRange(481, 593):
                    d_24_v1_: int = compr_8_
                    if ((481) <= (d_24_v1_)) and ((d_24_v1_) < (593)):
                        coll8_[default__.safeModuloInt(d_24_v1_, 843)] = len(_dafny.SeqWithoutIsStrInference([False]))
                return _dafny.Map(coll8_)
            coll6_ = _dafny.Map()
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(481, 593):
                    d_24_v1_: int = compr_7_
                    if ((481) <= (d_24_v1_)) and ((d_24_v1_) < (593)):
                        coll7_[default__.safeModuloInt(d_24_v1_, 843)] = len(_dafny.SeqWithoutIsStrInference([False]))
                return _dafny.Map(coll7_)
            compr_6_: int
            for compr_6_ in (iife7_()
            ).keys.Elements:
                d_25_v0_: int = compr_6_
                if (d_25_v0_) in (iife8_()
                ):
                    coll6_[default__.safeDivisionInt(d_25_v0_, -620)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyeoofmh"))
            return _dafny.Map(coll6_)
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([943 for d_21_i0_ in range(default__.abs(134))])]) if False else _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([7, (0) - (len(_dafny.Map({520: not(True)}))), 496, (0) - ((0) - (-47))])]))) | (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([975])]) if not(True) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - ((0) - (-786)) for d_22_i1_ in range(default__.abs(907))]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([811, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_23_i2_ in range(default__.abs(973))])), 948])).cardinality, len(iife6_()
        ), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmjx")))])]))))

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        return (_dafny.Set({True})).intersection(_dafny.Set({False, False}))

    @staticmethod
    def fm14(p0, globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: _dafny.MultiSet
            for compr_9_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, True]), _dafny.MultiSet([True])])).Elements:
                d_27_v0_: _dafny.MultiSet = compr_9_
                if (d_27_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, True]), _dafny.MultiSet([True])])):
                    coll9_[d_27_v0_] = len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({True: False})))]))
            return _dafny.Map(coll9_)
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({not(False), True}))) for d_26_i0_ in range(default__.abs(-173))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vetmckar")))}))]))])), _dafny.SeqWithoutIsStrInference([len(iife9_()
        ), len(_dafny.Set({not(not(not(True)))})), 670]), _dafny.SeqWithoutIsStrInference([818, len((D10_DC21(_dafny.SeqWithoutIsStrInference([not(True)]))).cf33), 699])})

    @staticmethod
    def fm15(globalState):
        return ((_dafny.Map({-782: False})) | (_dafny.Map({(_dafny.MultiSet([(_dafny.MultiSet([-698])).cardinality])).cardinality: not(True)}))) | (_dafny.Map({-381: not(False)}))

    @staticmethod
    def fm16(globalState):
        return _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, True])]))})

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        if (_dafny.Set({_dafny.SeqWithoutIsStrInference([False])})).isdisjoint(_dafny.Set({_dafny.SeqWithoutIsStrInference([False, True])})):
            if True:
                return D1_DC3(_dafny.MultiSet([not(False), False, False]), False, True)
            elif True:
                return D1_DC3(_dafny.MultiSet([True]), False, False)
        elif True:
            return D1_DC3(_dafny.MultiSet([not(not(True)), True]), not((D7_DC13(True)).cf17), True)

    @staticmethod
    def m0(p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        r3: _dafny.Set = _dafny.Set({})
        d_28_v0_: int
        d_28_v0_ = 232
        d_29_v1_: _dafny.Seq
        d_29_v1_ = _dafny.SeqWithoutIsStrInference([d_28_v0_])
        d_30_v2_: D8
        d_30_v2_ = D8_DC15(d_29_v1_)
        d_29_v1_ = (d_30_v2_).cf19
        d_31_v3_: D7
        d_31_v3_ = D7_DC13(p1)
        d_32_v4_: D7
        d_32_v4_ = D7_DC14(d_31_v3_)
        d_33_v5_: _dafny.Map
        d_33_v5_ = _dafny.Map({d_32_v4_: p1})
        d_33_v5_ = (d_33_v5_).set(d_32_v4_, False)
        d_34_v6_: _dafny.Map
        d_34_v6_ = _dafny.Map({p1: d_28_v0_})
        d_35_v7_: _dafny.Set
        d_35_v7_ = _dafny.Set({d_34_v6_, d_34_v6_})
        d_36_i0_: int
        d_36_i0_ = 0
        with _dafny.label("0"):
            def iife10_():
                coll10_ = _dafny.Set()
                compr_10_: _dafny.Map
                for compr_10_ in (d_35_v7_).Elements:
                    d_53_v8_: _dafny.Map = compr_10_
                    if (d_53_v8_) in (d_35_v7_):
                        coll10_ = coll10_.union(_dafny.Set([d_53_v8_]))
                return _dafny.Set(coll10_)
            while (d_35_v7_).ispropersubset(iife10_()
            ):
                with _dafny.c_label("0"):
                    if (d_36_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_36_i0_ = (d_36_i0_) + (1)
                    d_37_v9_: _dafny.Array
                    def lambda0_(d_38_v0_):
                        def lambda1_(d_39_i1_):
                            return (d_39_i1_) * (d_38_v0_)

                        return lambda1_

                    init0_ = lambda0_(d_28_v0_)
                    nw0_ = _dafny.Array(None, 8)
                    for i0_0_ in range(nw0_.length(0)):
                        nw0_[i0_0_] = init0_(i0_0_)
                    d_37_v9_ = nw0_
                    d_37_v9_ = (D1_DC1(d_37_v9_)).cf1
                    d_40_v10_: _dafny.Array
                    def lambda2_(d_41_v1_, d_42_v0_, d_43_p1_):
                        def lambda3_(d_44_i2_):
                            return (_dafny.MultiSet([d_41_v1_])) - (_dafny.MultiSet([d_41_v1_, (_dafny.SeqWithoutIsStrInference([d_42_v0_, len(d_41_v1_)])).set(default__.safeIndex(d_42_v0_, len(_dafny.SeqWithoutIsStrInference([d_42_v0_, len(d_41_v1_)]))), (0) - (len(_dafny.SeqWithoutIsStrInference([d_43_p1_]))))]))

                        return lambda3_

                    init1_ = lambda2_(d_29_v1_, d_28_v0_, p1)
                    nw1_ = _dafny.Array(None, 18)
                    for i0_1_ in range(nw1_.length(0)):
                        nw1_[i0_1_] = init1_(i0_1_)
                    d_40_v10_ = nw1_
                    d_45_v11_: _dafny.Seq
                    d_45_v11_ = _dafny.SeqWithoutIsStrInference([d_29_v1_])
                    index0_ = default__.safeIndex(232, (d_40_v10_).length(0))
                    (d_40_v10_)[index0_] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_28_v0_])])) + (d_45_v11_))
                    d_46_v12_: C0
                    nw2_ = C0()
                    nw2_.ctor__(p1)
                    d_46_v12_ = nw2_
                    d_47_v13_: _dafny.Map
                    d_47_v13_ = _dafny.Map({d_46_v12_: d_29_v1_})
                    d_48_v14_: _dafny.MultiSet
                    d_48_v14_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_28_v0_, d_28_v0_]), ((d_47_v13_)[d_46_v12_] if (d_46_v12_) in (d_47_v13_) else d_29_v1_)])
                    index1_ = default__.safeIndex(232, (d_40_v10_).length(0))
                    (d_40_v10_)[index1_] = (D7_DC12(d_48_v14_)).cf16
                    d_49_v15_: bool
                    d_49_v15_ = False
                    d_50_v16_: str
                    d_50_v16_ = _dafny.CodePoint('r')
                    d_51_v17_: str
                    d_51_v17_ = d_50_v16_
                    d_52_v18_: _dafny.Seq
                    d_52_v18_ = _dafny.SeqWithoutIsStrInference([(d_51_v17_)])
                    rhs0_ = (default__.fm7(len(d_52_v18_), d_28_v0_, globalState)) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrtno")) if p1 else d_52_v18_))
                    rhs1_ = d_46_v12_
                    d_49_v15_ = rhs0_
                    d_46_v12_ = rhs1_
                    d_29_v1_ = (_dafny.SeqWithoutIsStrInference([d_28_v0_])) + (d_29_v1_)
                    pass
            pass
        d_54_v19_: bool
        d_54_v19_ = True
        d_54_v19_ = p1
        d_55_i3_: int
        d_55_i3_ = 0
        with _dafny.label("1"):
            while p1:
                with _dafny.c_label("1"):
                    if (d_55_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_55_i3_ = (d_55_i3_) + (1)
                    d_54_v19_ = p1
                    r0 = d_28_v0_
                    d_56_v20_: _dafny.Seq
                    d_56_v20_ = _dafny.SeqWithoutIsStrInference([False])
                    d_57_v21_: _dafny.Seq
                    d_57_v21_ = _dafny.SeqWithoutIsStrInference([d_32_v4_, D7_DC14(d_31_v3_)])
                    d_58_v22_: _dafny.Map
                    d_58_v22_ = _dafny.Map({d_28_v0_: d_28_v0_})
                    d_59_v23_: D8
                    d_59_v23_ = D8_DC16(d_54_v19_, _dafny.MultiSet(d_56_v20_), d_57_v21_, d_58_v22_)
                    d_60_v24_: D8
                    d_60_v24_ = D8_DC18(d_59_v23_)
                    d_61_v25_: D8
                    d_61_v25_ = D8_DC18(d_59_v23_)
                    d_62_v26_: D8
                    d_62_v26_ = D8_DC18(d_59_v23_)
                    d_63_v27_: D8
                    d_63_v27_ = D8_DC18(d_62_v26_)
                    d_64_v28_: D8
                    d_64_v28_ = D8_DC18(d_62_v26_)
                    source2_ = d_64_v28_
                    if source2_.is_DC16:
                        d_65___mcc_h0_ = source2_.cf20
                        d_66___mcc_h1_ = source2_.cf21
                        d_67___mcc_h2_ = source2_.cf22
                        d_68___mcc_h3_ = source2_.cf23
                        d_69_cf23_ = d_68___mcc_h3_
                        d_70_cf22_ = d_67___mcc_h2_
                        d_71_cf21_ = d_66___mcc_h1_
                        d_72_cf20_ = d_65___mcc_h0_
                        d_73_v29_: _dafny.Array
                        nw3_ = _dafny.Array(_dafny.Map({}), 15)
                        d_73_v29_ = nw3_
                        index2_ = default__.safeIndex(447, (d_73_v29_).length(0))
                        (d_73_v29_)[index2_] = default__.fm15(globalState)
                        d_74_v30_: _dafny.Map
                        d_74_v30_ = _dafny.Map({d_28_v0_: d_72_cf20_})
                        d_75_v31_: _dafny.Seq
                        d_75_v31_ = _dafny.SeqWithoutIsStrInference([d_74_v30_, d_74_v30_])
                        d_76_v32_: _dafny.Seq
                        d_76_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nixsur"))
                        d_77_v33_: _dafny.Map
                        d_77_v33_ = _dafny.Map({d_76_v32_: d_28_v0_})
                        d_78_v34_: _dafny.Seq
                        d_78_v34_ = _dafny.SeqWithoutIsStrInference([d_77_v33_, d_77_v33_])
                        d_79_v35_: _dafny.Set
                        d_79_v35_ = _dafny.Set({d_54_v19_, d_54_v19_, p1})
                        index3_ = default__.safeIndex(447, (d_73_v29_).length(0))
                        (d_73_v29_)[index3_] = ((d_75_v31_)[default__.safeIndex(len((d_78_v34_)[default__.safeIndex(len(_dafny.Set({d_28_v0_, d_28_v0_})), len(d_78_v34_))]), len(d_75_v31_))]) | (((d_74_v30_).set((0) - (d_28_v0_), p1)).set(d_28_v0_, default__.fm1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmgxqq")), d_28_v0_, d_79_v35_, globalState)))
                        d_80_v36_: T0
                        nw4_ = C0()
                        nw4_.ctor__(p1)
                        d_80_v36_ = nw4_
                        d_80_v36_ = d_80_v36_
                        d_81_v37_: _dafny.Set
                        d_81_v37_ = _dafny.Set({(0) - (len(_dafny.Map({(d_74_v30_).set(d_28_v0_, d_72_cf20_): d_54_v19_})))})
                        d_82_v38_: _dafny.Array
                        nw5_ = _dafny.Array(None, 4)
                        nw5_[int(0)] = d_81_v37_
                        nw5_[int(1)] = d_81_v37_
                        nw5_[int(2)] = _dafny.Set({d_28_v0_, d_28_v0_})
                        nw5_[int(3)] = d_81_v37_
                        d_82_v38_ = nw5_
                        index4_ = default__.safeIndex(317, (d_82_v38_).length(0))
                        (d_82_v38_)[index4_] = d_81_v37_
                        d_83_v39_: _dafny.Array
                        def lambda4_(d_84_v0_):
                            def lambda5_(d_85_i4_):
                                return default__.safeDivisionInt(d_85_i4_, d_84_v0_)

                            return lambda5_

                        init2_ = lambda4_(d_28_v0_)
                        nw6_ = _dafny.Array(None, 22)
                        for i0_2_ in range(nw6_.length(0)):
                            nw6_[i0_2_] = init2_(i0_2_)
                        d_83_v39_ = nw6_
                        index5_ = default__.safeIndex(833, (d_83_v39_).length(0))
                        (d_83_v39_)[index5_] = default__.safeModuloInt(d_28_v0_, len(d_58_v22_))
                        index6_ = default__.safeIndex(982, (d_83_v39_).length(0))
                        (d_83_v39_)[index6_] = len(d_29_v1_)
                        d_86_v40_: C0
                        nw7_ = C0()
                        nw7_.ctor__(p1)
                        d_86_v40_ = nw7_
                        d_87_v41_: _dafny.Seq
                        d_87_v41_ = _dafny.SeqWithoutIsStrInference([d_86_v40_])
                        d_88_v42_: _dafny.Seq
                        d_88_v42_ = ((_dafny.SeqWithoutIsStrInference([(d_87_v41_)[default__.safeIndex(d_28_v0_, len(d_87_v41_))]])).set(default__.safeIndex(len(d_87_v41_), len(_dafny.SeqWithoutIsStrInference([(d_87_v41_)[default__.safeIndex(d_28_v0_, len(d_87_v41_))]]))), d_86_v40_)) + (d_87_v41_)
                        index7_ = default__.safeIndex(317, (d_82_v38_).length(0))
                        index8_ = default__.safeIndex(833, (d_83_v39_).length(0))
                        index9_ = default__.safeIndex(982, (d_83_v39_).length(0))
                        rhs2_ = ((d_81_v37_) - (default__.fm16(globalState))).intersection(d_81_v37_)
                        rhs3_ = d_76_v32_
                        rhs4_ = (d_29_v1_)[default__.safeIndex(d_28_v0_, len(d_29_v1_))]
                        rhs5_ = (d_28_v0_) * (len(d_76_v32_))
                        rhs6_ = d_88_v42_
                        lhs0_ = d_82_v38_
                        lhs1_ = default__.safeIndex(317, (d_82_v38_).length(0))
                        lhs2_ = d_83_v39_
                        lhs3_ = default__.safeIndex(833, (d_83_v39_).length(0))
                        lhs4_ = d_83_v39_
                        lhs5_ = default__.safeIndex(982, (d_83_v39_).length(0))
                        lhs0_[lhs1_] = rhs2_
                        d_76_v32_ = rhs3_
                        lhs2_[lhs3_] = rhs4_
                        lhs4_[lhs5_] = rhs5_
                        d_88_v42_ = rhs6_
                        d_89_v43_: _dafny.MultiSet
                        d_89_v43_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_28_v0_, (d_83_v39_)[default__.safeIndex(833, (d_83_v39_).length(0))]]), d_29_v1_, d_29_v1_, d_29_v1_])
                        d_90_v44_: D7
                        d_90_v44_ = D7_DC12(d_89_v43_)
                        pat_let_tv0_ = d_89_v43_
                        def iife11_(_pat_let0_0):
                            def iife12_(d_91_dt__update__tmp_h0_):
                                def iife13_(_pat_let1_0):
                                    def iife14_(d_92_dt__update_hcf16_h0_):
                                        return D7_DC12(d_92_dt__update_hcf16_h0_)
                                    return iife14_(_pat_let1_0)
                                return iife13_(pat_let_tv0_)
                            return iife12_(_pat_let0_0)
                        rhs7_ = (d_72_cf20_ if d_80_v36_.f0 else d_80_v36_.f0)
                        rhs8_ = (d_83_v39_)[default__.safeIndex(833, (d_83_v39_).length(0))]
                        rhs9_ = iife11_(d_90_v44_)
                        d_72_cf20_ = rhs7_
                        r2 = rhs8_
                        d_90_v44_ = rhs9_
                    elif source2_.is_DC17:
                        d_93___mcc_h4_ = source2_.cf24
                        d_94___mcc_h5_ = source2_.cf25
                        d_95___mcc_h6_ = source2_.cf26
                        d_96___mcc_h7_ = source2_.cf27
                        d_97_cf27_ = d_96___mcc_h7_
                        d_98_cf26_ = d_95___mcc_h6_
                        d_99_cf25_ = d_94___mcc_h5_
                        d_100_cf24_ = d_93___mcc_h4_
                        d_101_v45_: str
                        d_101_v45_ = _dafny.CodePoint('f')
                        d_102_v46_: D3
                        d_102_v46_ = D3_DC6(d_28_v0_, d_98_cf26_, p1)
                        d_103_v47_: _dafny.Set
                        d_103_v47_ = _dafny.Set({(d_102_v46_).cf10, p1, p1})
                        d_104_v48_: _dafny.Map
                        d_104_v48_ = _dafny.Map({(d_29_v1_)[default__.safeIndex(d_99_cf25_, len(d_29_v1_))]: default__.fm1((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efgeywy"))).set(default__.safeIndex(d_99_cf25_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efgeywy")))), d_101_v45_), d_28_v0_, d_103_v47_, globalState)})
                        d_104_v48_ = (d_104_v48_).set(d_97_cf27_, d_54_v19_)
                        d_54_v19_ = (d_99_cf25_) <= ((d_29_v1_)[default__.safeIndex((0) - ((D6_DC11(d_97_cf27_)).cf15), len(d_29_v1_))])
                        d_105_v49_: C0
                        nw8_ = C0()
                        nw8_.ctor__(d_54_v19_)
                        d_105_v49_ = nw8_
                        d_105_v49_ = d_105_v49_
                        d_106_v50_: _dafny.Array
                        nw9_ = _dafny.Array(int(0), 4)
                        d_106_v50_ = nw9_
                        d_106_v50_ = d_106_v50_
                    elif source2_.is_DC15:
                        d_107___mcc_h8_ = source2_.cf19
                        d_108_cf19_ = d_107___mcc_h8_
                        d_109_v51_: int
                        d_109_v51_ = d_28_v0_
                        d_110_v52_: _dafny.Map
                        d_110_v52_ = _dafny.Map({d_54_v19_: d_109_v51_})
                        d_111_v53_: _dafny.Map
                        d_111_v53_ = _dafny.Map({p1: d_54_v19_})
                        d_112_v54_: _dafny.Map
                        d_112_v54_ = _dafny.Map({d_110_v52_: ((d_111_v53_)[d_54_v19_] if (d_54_v19_) in (d_111_v53_) else p1)})
                        d_112_v54_ = (d_112_v54_).set(d_110_v52_, p1)
                        d_54_v19_ = ((d_111_v53_)[d_54_v19_] if (d_54_v19_) in (d_111_v53_) else not (not(p1)) or (p1))
                        d_54_v19_ = p1
                        d_113_v55_: _dafny.Set
                        d_113_v55_ = _dafny.Set({d_54_v19_, d_54_v19_})
                        d_54_v19_ = (d_113_v55_).ispropersubset((d_113_v55_).intersection(d_113_v55_))
                    elif True:
                        d_114___mcc_h9_ = source2_.cf28
                        d_115_cf28_ = d_114___mcc_h9_
                        d_116_v56_: _dafny.MultiSet
                        d_116_v56_ = _dafny.MultiSet([(d_56_v20_)[default__.safeIndex(d_28_v0_, len(d_56_v20_))], p1])
                        d_117_v57_: _dafny.Set
                        d_117_v57_ = _dafny.Set({p1, p1})
                        d_54_v19_ = not(default__.fm1(default__.fm7(((d_116_v56_)[d_54_v19_] if (d_54_v19_) in (d_116_v56_) else d_28_v0_), d_28_v0_, globalState), d_28_v0_, d_117_v57_, globalState))
                        d_29_v1_ = ((d_29_v1_ if p1 else _dafny.SeqWithoutIsStrInference([d_28_v0_ for d_118_i5_ in range(default__.abs(907))]))).set(default__.safeIndex(d_28_v0_, len((d_29_v1_ if p1 else _dafny.SeqWithoutIsStrInference([d_28_v0_ for d_119_i5_ in range(default__.abs(907))])))), d_28_v0_)
                        d_120_v58_: C0
                        nw10_ = C0()
                        nw10_.ctor__(p1)
                        d_120_v58_ = nw10_
                        d_121_v59_: _dafny.Map
                        d_121_v59_ = _dafny.Map({d_120_v58_: p1})
                        d_122_v60_: _dafny.Map
                        d_122_v60_ = _dafny.Map({d_121_v59_: p1})
                        d_123_v61_: C0
                        nw11_ = C0()
                        nw11_.ctor__(((d_122_v60_)[d_121_v59_] if (d_121_v59_) in (d_122_v60_) else p1))
                        d_123_v61_ = nw11_
                        d_124_v62_: _dafny.Seq
                        d_124_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owkgj"))
                        d_125_v63_: _dafny.Map
                        d_125_v63_ = _dafny.Map({d_28_v0_: d_124_v62_})
                        d_126_v64_: str
                        d_126_v64_ = _dafny.CodePoint('d')
                        d_127_v65_: _dafny.Map
                        d_127_v65_ = _dafny.Map({d_54_v19_: d_126_v64_})
                        d_128_v66_: _dafny.Map
                        d_128_v66_ = _dafny.Map({len((((d_125_v63_)[d_28_v0_] if (d_28_v0_) in (d_125_v63_) else d_124_v62_) if d_54_v19_ else (d_124_v62_).set(default__.safeIndex(d_28_v0_, len(d_124_v62_)), ((d_127_v65_)[p1] if (p1) in (d_127_v65_) else d_126_v64_)))): d_123_v61_})
                        d_129_v67_: D6
                        d_129_v67_ = D6_DC10(d_123_v61_)
                        d_123_v61_ = ((d_128_v66_)[len(((d_56_v20_).set(default__.safeIndex(302, len(d_56_v20_)), d_54_v19_)) + (_dafny.SeqWithoutIsStrInference([d_54_v19_, p1, True, d_54_v19_])))] if (len(((d_56_v20_).set(default__.safeIndex(302, len(d_56_v20_)), d_54_v19_)) + (_dafny.SeqWithoutIsStrInference([d_54_v19_, p1, True, d_54_v19_])))) in (d_128_v66_) else (d_129_v67_).cf14)
                        d_54_v19_ = p1
                    d_130_v68_: _dafny.Seq
                    d_130_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emnf"))
                    if default__.fm1(d_130_v68_, d_28_v0_, default__.fm13(d_54_v19_, d_54_v19_, p1, globalState), globalState):
                        d_54_v19_ = d_54_v19_
                        d_131_v69_: _dafny.Set
                        d_131_v69_ = _dafny.Set({d_54_v19_, d_54_v19_})
                        d_132_v70_: C0
                        nw12_ = C0()
                        nw12_.ctor__((p1) or (default__.fm1(d_130_v68_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ertqblss"))), d_131_v69_, globalState)))
                        d_132_v70_ = nw12_
                        d_133_v71_: _dafny.Array
                        nw13_ = _dafny.Array(int(0), 16)
                        d_133_v71_ = nw13_
                        index10_ = default__.safeIndex(624, (d_133_v71_).length(0))
                        (d_133_v71_)[index10_] = d_28_v0_
                        index11_ = default__.safeIndex(624, (d_133_v71_).length(0))
                        rhs10_ = (d_29_v1_) <= (d_29_v1_)
                        rhs11_ = d_130_v68_
                        rhs12_ = d_28_v0_
                        lhs6_ = d_133_v71_
                        lhs7_ = default__.safeIndex(624, (d_133_v71_).length(0))
                        d_54_v19_ = rhs10_
                        d_130_v68_ = rhs11_
                        lhs6_[lhs7_] = rhs12_
                        d_134_v72_: _dafny.Array
                        def lambda6_(d_135_i6_):
                            return False

                        init3_ = lambda6_
                        nw14_ = _dafny.Array(None, 22)
                        for i0_3_ in range(nw14_.length(0)):
                            nw14_[i0_3_] = init3_(i0_3_)
                        d_134_v72_ = nw14_
                        index12_ = default__.safeIndex(852, (d_134_v72_).length(0))
                        (d_134_v72_)[index12_] = p1
                        index13_ = default__.safeIndex(852, (d_134_v72_).length(0))
                        (d_134_v72_)[index13_] = p1
                        d_136_v73_: C0
                        nw15_ = C0()
                        nw15_.ctor__((default__.fm17(d_28_v0_, 755, p1, globalState)).cf5)
                        d_136_v73_ = nw15_
                    elif True:
                        d_137_v74_: _dafny.Array
                        nw16_ = _dafny.Array(int(0), 23)
                        d_137_v74_ = nw16_
                        index14_ = default__.safeIndex(860, (d_137_v74_).length(0))
                        (d_137_v74_)[index14_] = d_28_v0_
                        d_138_v75_: str
                        d_138_v75_ = _dafny.CodePoint('d')
                        d_139_v76_: _dafny.Seq
                        d_139_v76_ = _dafny.SeqWithoutIsStrInference([d_130_v68_])
                        d_140_v77_: int
                        d_140_v77_ = d_28_v0_
                        d_141_v78_: _dafny.Map
                        d_141_v78_ = _dafny.Map({d_28_v0_: d_139_v76_})
                        d_142_v79_: _dafny.Seq
                        d_142_v79_ = _dafny.SeqWithoutIsStrInference([((d_141_v78_)[-272] if (-272) in (d_141_v78_) else (d_139_v76_).set(default__.safeIndex(d_28_v0_, len(d_139_v76_)), d_130_v68_))])
                        index15_ = default__.safeIndex(860, (d_137_v74_).length(0))
                        rhs13_ = (d_140_v77_)
                        rhs14_ = d_138_v75_
                        rhs15_ = (d_130_v68_)[default__.safeIndex(d_28_v0_, len(d_130_v68_))]
                        rhs16_ = ((d_142_v79_)[default__.safeIndex(default__.fm2(p1, globalState), len(d_142_v79_))]) + (d_139_v76_)
                        rhs17_ = default__.safeDivisionInt(d_28_v0_, (d_29_v1_)[default__.safeIndex(d_28_v0_, len(d_29_v1_))])
                        lhs8_ = d_137_v74_
                        lhs9_ = default__.safeIndex(860, (d_137_v74_).length(0))
                        lhs8_[lhs9_] = rhs13_
                        d_138_v75_ = rhs14_
                        d_138_v75_ = rhs15_
                        d_139_v76_ = rhs16_
                        d_28_v0_ = rhs17_
                        d_143_v80_: _dafny.Array
                        nw17_ = _dafny.Array(None, 5)
                        d_143_v80_ = nw17_
                        nw18_ = _dafny.Array(None, 7)
                        d_143_v80_ = nw18_
                        d_144_v82_: D9
                        def iife15_():
                            coll11_ = _dafny.Map()
                            compr_11_: int
                            for compr_11_ in _dafny.IntegerRange(-359, 203):
                                d_145_v81_: int = compr_11_
                                if ((-359) <= (d_145_v81_)) and ((d_145_v81_) < (203)):
                                    coll11_[default__.safeModuloInt(d_145_v81_, d_28_v0_)] = d_56_v20_
                            return _dafny.Map(coll11_)
                        d_144_v82_ = D9_DC19(iife15_()
)
                        index16_ = default__.safeIndex(860, (d_137_v74_).length(0))
                        (d_137_v74_)[index16_] = len((d_144_v82_).cf29)
                        d_54_v19_ = d_54_v19_
                        d_146_v83_: C0
                        nw19_ = C0()
                        nw19_.ctor__(d_54_v19_)
                        d_146_v83_ = nw19_
                    pass
            pass
        d_147_i7_: int
        d_147_i7_ = 0
        with _dafny.label("2"):
            while not(not(d_54_v19_)):
                with _dafny.c_label("2"):
                    if (d_147_i7_) >= (100):
                        raise _dafny.Break("2")
                    d_147_i7_ = (d_147_i7_) + (1)
                    d_148_v84_: _dafny.Seq
                    d_148_v84_ = _dafny.SeqWithoutIsStrInference([not(p1), d_54_v19_])
                    d_149_v85_: _dafny.MultiSet
                    d_149_v85_ = _dafny.MultiSet([(d_148_v84_)[default__.safeIndex(d_28_v0_, len(d_148_v84_))]])
                    d_150_v86_: D1
                    d_150_v86_ = D1_DC3(d_149_v85_, d_54_v19_, d_54_v19_)
                    d_151_v87_: _dafny.Set
                    d_151_v87_ = _dafny.Set({p1, d_54_v19_, p1, (d_150_v86_).cf4})
                    d_54_v19_ = default__.fm1(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_152_i8_ in range(default__.abs(253))]), d_28_v0_, d_151_v87_, globalState)
                    if (448) == (d_28_v0_):
                        d_153_v88_: _dafny.Seq
                        d_153_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlwntcgy"))
                        d_154_v89_: C0
                        nw20_ = C0()
                        nw20_.ctor__(p1)
                        d_154_v89_ = nw20_
                        d_155_v90_: str
                        d_155_v90_ = _dafny.CodePoint('m')
                        d_153_v88_ = (d_153_v88_).set(default__.safeIndex((d_28_v0_) * (len(_dafny.Set({d_154_v89_}))), len(d_153_v88_)), (d_155_v90_ if p1 else _dafny.CodePoint('j')))
                        d_156_v91_: _dafny.Array
                        def lambda7_(d_157_v2_):
                            def lambda8_(d_158_i9_):
                                return d_157_v2_

                            return lambda8_

                        init4_ = lambda7_(d_30_v2_)
                        nw21_ = _dafny.Array(None, 19)
                        for i0_4_ in range(nw21_.length(0)):
                            nw21_[i0_4_] = init4_(i0_4_)
                        d_156_v91_ = nw21_
                        d_159_v92_: _dafny.Set
                        d_159_v92_ = _dafny.Set({d_156_v91_})
                        d_54_v19_ = ((521 if d_54_v19_ else d_28_v0_)) != (len((d_159_v92_).intersection(d_159_v92_)))
                        d_29_v1_ = d_29_v1_
                        d_160_v93_: _dafny.Array
                        nw22_ = _dafny.Array(False, 26)
                        d_160_v93_ = nw22_
                        index17_ = default__.safeIndex(75, (d_160_v93_).length(0))
                        (d_160_v93_)[index17_] = d_54_v19_
                        index18_ = default__.safeIndex(75, (d_160_v93_).length(0))
                        (d_160_v93_)[index18_] = False
                        d_153_v88_ = default__.fm7(d_28_v0_, d_28_v0_, globalState)
                    elif True:
                        d_161_v94_: _dafny.Map
                        d_161_v94_ = _dafny.Map({False: d_54_v19_})
                        d_161_v94_ = (d_161_v94_).set(d_54_v19_, not (p1) or (p1))
                        r1 = d_28_v0_
                        d_162_v95_: _dafny.Array
                        nw23_ = _dafny.Array(None, 20)
                        d_162_v95_ = nw23_
                        d_163_v96_: C0
                        nw24_ = C0()
                        nw24_.ctor__(not(p1))
                        d_163_v96_ = nw24_
                        index19_ = default__.safeIndex(849, (d_162_v95_).length(0))
                        (d_162_v95_)[index19_] = d_163_v96_
                        index20_ = default__.safeIndex(849, (d_162_v95_).length(0))
                        (d_162_v95_)[index20_] = d_163_v96_
                        d_164_v97_: C0
                        nw25_ = C0()
                        nw25_.ctor__(p1)
                        d_164_v97_ = nw25_
                        d_165_v98_: _dafny.Map
                        d_165_v98_ = _dafny.Map({d_28_v0_: not(p1)})
                        d_166_v99_: _dafny.Set
                        d_166_v99_ = _dafny.Set({d_29_v1_, d_29_v1_})
                        d_167_v100_: D1
                        d_167_v100_ = D1_DC2(d_166_v99_)
                        d_168_v101_: str
                        d_168_v101_ = _dafny.CodePoint('w')
                        d_169_v102_: _dafny.Seq
                        d_169_v102_ = _dafny.SeqWithoutIsStrInference([d_161_v94_])
                        d_170_v103_: _dafny.MultiSet
                        d_170_v103_ = _dafny.MultiSet([d_168_v101_, (d_168_v101_ if (d_164_v97_).fm6(d_28_v0_, _dafny.MultiSet(d_169_v102_), d_28_v0_, globalState) else d_168_v101_), d_168_v101_, d_168_v101_, d_168_v101_])
                        def iife16_():
                            coll12_ = _dafny.Map()
                            compr_12_: int
                            for compr_12_ in (default__.fm3(globalState)).Elements:
                                d_171_v104_: int = compr_12_
                                if (d_171_v104_) in (default__.fm3(globalState)):
                                    coll12_[(d_171_v104_) - ((0) - (d_28_v0_))] = False
                            return _dafny.Map(coll12_)
                        rhs18_ = (d_28_v0_) in (d_165_v98_)
                        rhs19_ = (iife16_()
                        ) | ((d_165_v98_).set(556, p1))
                        rhs20_ = d_167_v100_
                        rhs21_ = d_162_v95_
                        rhs22_ = (d_170_v103_).intersection(d_170_v103_)
                        d_54_v19_ = rhs18_
                        d_165_v98_ = rhs19_
                        d_167_v100_ = rhs20_
                        d_162_v95_ = rhs21_
                        d_170_v103_ = rhs22_
                    d_172_v105_: _dafny.Seq
                    d_172_v105_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hr"))
                    d_28_v0_ = (0) - (default__.safeModuloInt((20) + (655), default__.safeModuloInt(d_28_v0_, len(d_172_v105_))))
                    d_54_v19_ = (p1) and ((d_54_v19_ if d_54_v19_ else False))
                    pass
            pass
        r0 = d_28_v0_
        r1 = d_28_v0_
        r2 = d_28_v0_
        d_173_v106_: _dafny.Set
        d_173_v106_ = _dafny.Set({True})
        d_174_v107_: _dafny.Set
        d_174_v107_ = _dafny.Set({d_173_v106_, d_173_v106_, d_173_v106_})
        r3 = d_174_v107_
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_175_globalState_: GlobalState
        nw26_ = GlobalState()
        nw26_.ctor__()
        d_175_globalState_ = nw26_
        d_176_v0_: _dafny.Array
        nw27_ = _dafny.Array(False, 24)
        d_176_v0_ = nw27_
        d_177_v1_: int
        d_177_v1_ = 365
        d_178_v2_: bool
        d_178_v2_ = True
        d_179_v3_: int
        d_180_v4_: int
        d_181_v5_: int
        d_182_v6_: _dafny.Set
        out0_: int
        out1_: int
        out2_: int
        out3_: _dafny.Set
        out0_, out1_, out2_, out3_ = default__.m0(_dafny.Map({d_176_v0_: d_177_v1_}), d_178_v2_, d_175_globalState_)
        d_179_v3_ = out0_
        d_180_v4_ = out1_
        d_181_v5_ = out2_
        d_182_v6_ = out3_
        d_183_v7_: _dafny.Seq
        d_183_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbehkj"))
        d_180_v4_ = default__.safeDivisionInt(default__.safeModuloInt((0) - (d_180_v4_), len(d_183_v7_)), 813)
        d_184_v8_: str
        d_184_v8_ = _dafny.CodePoint('k')
        d_185_v9_: _dafny.Map
        d_185_v9_ = _dafny.Map({d_178_v2_: d_177_v1_})
        d_186_v10_: str
        d_186_v10_ = _dafny.CodePoint('u')
        d_187_v11_: _dafny.Array
        nw28_ = _dafny.Array(None, 20)
        nw28_[int(0)] = d_184_v8_
        nw28_[int(1)] = d_184_v8_
        nw28_[int(2)] = d_184_v8_
        nw28_[int(3)] = d_184_v8_
        nw28_[int(4)] = default__.fm0(d_183_v7_, (0) - (d_181_v5_), d_181_v5_, d_175_globalState_)
        nw28_[int(5)] = default__.fm0(d_183_v7_, len(_dafny.SeqWithoutIsStrInference([len(d_185_v9_), d_181_v5_, d_177_v1_])), d_177_v1_, d_175_globalState_)
        nw28_[int(6)] = d_184_v8_
        nw28_[int(7)] = d_184_v8_
        nw28_[int(8)] = d_184_v8_
        nw28_[int(9)] = d_184_v8_
        nw28_[int(10)] = _dafny.CodePoint('c')
        nw28_[int(11)] = d_184_v8_
        nw28_[int(12)] = (d_186_v10_)
        nw28_[int(13)] = default__.fm0(_dafny.SeqWithoutIsStrInference([d_184_v8_ for d_188_i0_ in range(default__.abs(441))]), d_181_v5_, d_179_v3_, d_175_globalState_)
        nw28_[int(14)] = d_184_v8_
        nw28_[int(15)] = d_184_v8_
        nw28_[int(16)] = (d_183_v7_)[default__.safeIndex(d_179_v3_, len(d_183_v7_))]
        nw28_[int(17)] = d_184_v8_
        nw28_[int(18)] = default__.fm0(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_189_i1_ in range(default__.abs(916))]), d_179_v3_, d_179_v3_, d_175_globalState_)
        nw28_[int(19)] = _dafny.CodePoint('x')
        d_187_v11_ = nw28_
        d_187_v11_ = d_187_v11_
        d_190_v12_: _dafny.Map
        d_190_v12_ = _dafny.Map({d_180_v4_: d_180_v4_})
        if (((d_190_v12_)[d_180_v4_] if (d_180_v4_) in (d_190_v12_) else d_180_v4_)) != (d_181_v5_):
            if True:
                d_191_v13_: _dafny.Array
                nw29_ = _dafny.Array(int(0), 28)
                d_191_v13_ = nw29_
                d_191_v13_ = d_191_v13_
                d_192_v14_: _dafny.Seq
                d_192_v14_ = _dafny.SeqWithoutIsStrInference([d_180_v4_, d_180_v4_])
                d_185_v9_ = (d_185_v9_).set(d_178_v2_, len(d_192_v14_))
                d_193_v15_: D1
                d_193_v15_ = D1_DC1(d_191_v13_)
                d_191_v13_ = (d_191_v13_ if not(d_178_v2_) else (d_193_v15_).cf1)
                index21_ = default__.safeIndex(925, (d_176_v0_).length(0))
                (d_176_v0_)[index21_] = d_178_v2_
                index22_ = default__.safeIndex(925, (d_176_v0_).length(0))
                (d_176_v0_)[index22_] = d_178_v2_
                index23_ = default__.safeIndex(925, (d_176_v0_).length(0))
                (d_176_v0_)[index23_] = (d_176_v0_)[default__.safeIndex(925, (d_176_v0_).length(0))]
            elif True:
                d_194_v16_: _dafny.Map
                d_194_v16_ = _dafny.Map({d_176_v0_: d_180_v4_})
                d_195_v17_: int
                d_196_v18_: int
                d_197_v19_: int
                d_198_v20_: _dafny.Set
                out4_: int
                out5_: int
                out6_: int
                out7_: _dafny.Set
                out4_, out5_, out6_, out7_ = default__.m0((d_194_v16_) | (d_194_v16_), True, d_175_globalState_)
                d_195_v17_ = out4_
                d_196_v18_ = out5_
                d_197_v19_ = out6_
                d_198_v20_ = out7_
                d_178_v2_ = d_178_v2_
                d_199_v21_: _dafny.Seq
                d_199_v21_ = _dafny.SeqWithoutIsStrInference([d_178_v2_, d_178_v2_])
                d_200_v22_: _dafny.Set
                d_200_v22_ = _dafny.Set({len(d_199_v21_)})
                d_201_v23_: int
                d_202_v24_: int
                d_203_v25_: int
                d_204_v26_: _dafny.Set
                out8_: int
                out9_: int
                out10_: int
                out11_: _dafny.Set
                out8_, out9_, out10_, out11_ = default__.m0(d_194_v16_, (d_200_v22_) == (d_200_v22_), d_175_globalState_)
                d_201_v23_ = out8_
                d_202_v24_ = out9_
                d_203_v25_ = out10_
                d_204_v26_ = out11_
                d_201_v23_ = default__.safeDivisionInt(d_203_v25_, d_203_v25_)
                d_205_v27_: D1
                d_205_v27_ = D1_DC3(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_178_v2_, d_178_v2_])), d_178_v2_, d_178_v2_)
                pat_let_tv1_ = d_178_v2_
                def iife17_(_pat_let2_0):
                    def iife18_(d_206_dt__update__tmp_h0_):
                        def iife19_(_pat_let3_0):
                            def iife20_(d_207_dt__update_hcf5_h0_):
                                return D1_DC3((d_206_dt__update__tmp_h0_).cf3, (d_206_dt__update__tmp_h0_).cf4, d_207_dt__update_hcf5_h0_)
                            return iife20_(_pat_let3_0)
                        return iife19_(pat_let_tv1_)
                    return iife18_(_pat_let2_0)
                d_178_v2_ = ((iife17_(d_205_v27_)).cf4) and (d_178_v2_)
            d_208_v28_: _dafny.MultiSet
            d_208_v28_ = _dafny.MultiSet([d_181_v5_, len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([31]) for d_209_i2_ in range(default__.abs(448))])), d_180_v4_, -248])
            d_178_v2_ = (d_208_v28_).isdisjoint((d_208_v28_) | (d_208_v28_))
            d_210_v29_: _dafny.MultiSet
            d_210_v29_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_184_v8_ for d_211_i3_ in range(default__.abs(259))])])
            d_212_v30_: _dafny.Seq
            d_212_v30_ = _dafny.SeqWithoutIsStrInference([d_178_v2_])
            d_213_v31_: int
            d_213_v31_ = len(d_212_v30_)
            d_214_v32_: _dafny.Set
            d_214_v32_ = _dafny.Set({d_178_v2_, d_178_v2_})
            d_215_v33_: _dafny.Set
            d_215_v33_ = _dafny.Set({not(d_178_v2_), d_178_v2_, default__.fm1(d_183_v7_, (d_213_v31_), d_214_v32_, d_175_globalState_), not(not(d_178_v2_))})
            rhs23_ = d_178_v2_
            rhs24_ = not((_dafny.MultiSet([d_183_v7_])).issubset((d_210_v29_).intersection(d_210_v29_)))
            rhs25_ = _dafny.SeqWithoutIsStrInference([d_184_v8_ for d_216_i4_ in range(default__.abs(232))])
            rhs26_ = d_184_v8_
            rhs27_ = len(d_214_v32_)
            d_178_v2_ = rhs23_
            d_178_v2_ = rhs24_
            d_183_v7_ = rhs25_
            d_184_v8_ = rhs26_
            d_179_v3_ = rhs27_
            d_217_v34_: _dafny.Map
            d_217_v34_ = _dafny.Map({d_176_v0_: d_180_v4_})
            d_218_v35_: D3
            d_218_v35_ = D3_DC5(d_217_v34_)
            d_219_v36_: int
            d_220_v37_: int
            d_221_v38_: int
            d_222_v39_: _dafny.Set
            out12_: int
            out13_: int
            out14_: int
            out15_: _dafny.Set
            out12_, out13_, out14_, out15_ = default__.m0((d_218_v35_).cf7, not(True), d_175_globalState_)
            d_219_v36_ = out12_
            d_220_v37_ = out13_
            d_221_v38_ = out14_
            d_222_v39_ = out15_
            d_223_v40_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.Seq({}), 28)
            d_223_v40_ = nw30_
            d_224_v41_: _dafny.Seq
            d_224_v41_ = _dafny.SeqWithoutIsStrInference([d_183_v7_, d_183_v7_, d_183_v7_])
            index24_ = default__.safeIndex(468, (d_223_v40_).length(0))
            (d_223_v40_)[index24_] = (d_224_v41_) + (d_224_v41_)
            index25_ = default__.safeIndex(468, (d_223_v40_).length(0))
            (d_223_v40_)[index25_] = (_dafny.SeqWithoutIsStrInference([d_183_v7_, d_183_v7_])) + (_dafny.SeqWithoutIsStrInference([d_183_v7_]))
        elif True:
            d_177_v1_ = d_177_v1_
            d_178_v2_ = False
            d_190_v12_ = (d_190_v12_).set(d_179_v3_, (265) + (default__.fm2(d_178_v2_, d_175_globalState_)))
            rhs28_ = d_178_v2_
            rhs29_ = d_184_v8_
            rhs30_ = d_179_v3_
            d_178_v2_ = rhs28_
            d_184_v8_ = rhs29_
            d_180_v4_ = rhs30_
            index26_ = default__.safeIndex(603, (d_176_v0_).length(0))
            (d_176_v0_)[index26_] = d_178_v2_
            index27_ = default__.safeIndex(603, (d_176_v0_).length(0))
            (d_176_v0_)[index27_] = ((d_180_v4_) < (d_181_v5_)) == (d_178_v2_)
        d_225_v42_: _dafny.Set
        d_225_v42_ = _dafny.Set({d_178_v2_})
        if (default__.fm1((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkkborlrx"))).set(default__.safeIndex(d_181_v5_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkkborlrx")))), d_184_v8_), default__.fm2(d_178_v2_, d_175_globalState_), d_225_v42_, d_175_globalState_) if True else (d_225_v42_) == (d_225_v42_)):
            d_226_v43_: _dafny.MultiSet
            d_226_v43_ = _dafny.MultiSet([d_178_v2_])
            d_227_v44_: _dafny.Map
            d_227_v44_ = _dafny.Map({d_176_v0_: (d_226_v43_).cardinality})
            d_228_v45_: D3
            d_228_v45_ = D3_DC5(d_227_v44_)
            d_229_v46_: int
            d_230_v47_: int
            d_231_v48_: int
            d_232_v49_: _dafny.Set
            out16_: int
            out17_: int
            out18_: int
            out19_: _dafny.Set
            out16_, out17_, out18_, out19_ = default__.m0((d_227_v44_ if d_178_v2_ else (d_228_v45_).cf7), d_178_v2_, d_175_globalState_)
            d_229_v46_ = out16_
            d_230_v47_ = out17_
            d_231_v48_ = out18_
            d_232_v49_ = out19_
            d_233_v50_: int
            d_234_v51_: int
            d_235_v52_: int
            d_236_v53_: _dafny.Set
            out20_: int
            out21_: int
            out22_: int
            out23_: _dafny.Set
            out20_, out21_, out22_, out23_ = default__.m0(d_227_v44_, d_178_v2_, d_175_globalState_)
            d_233_v50_ = out20_
            d_234_v51_ = out21_
            d_235_v52_ = out22_
            d_236_v53_ = out23_
            d_237_v54_: int
            d_238_v55_: int
            d_239_v56_: int
            d_240_v57_: _dafny.Set
            out24_: int
            out25_: int
            out26_: int
            out27_: _dafny.Set
            out24_, out25_, out26_, out27_ = default__.m0((d_227_v44_) | (d_227_v44_), d_178_v2_, d_175_globalState_)
            d_237_v54_ = out24_
            d_238_v55_ = out25_
            d_239_v56_ = out26_
            d_240_v57_ = out27_
            d_241_v58_: int
            d_242_v59_: int
            d_243_v60_: int
            d_244_v61_: _dafny.Set
            out28_: int
            out29_: int
            out30_: int
            out31_: _dafny.Set
            out28_, out29_, out30_, out31_ = default__.m0((d_227_v44_) | (d_227_v44_), d_178_v2_, d_175_globalState_)
            d_241_v58_ = out28_
            d_242_v59_ = out29_
            d_243_v60_ = out30_
            d_244_v61_ = out31_
            rhs31_ = not((d_178_v2_) or (d_178_v2_))
            rhs32_ = d_178_v2_
            rhs33_ = d_178_v2_
            rhs34_ = (d_187_v11_ if d_178_v2_ else d_187_v11_)
            rhs35_ = d_178_v2_
            d_178_v2_ = rhs31_
            d_178_v2_ = rhs32_
            d_178_v2_ = rhs33_
            d_187_v11_ = rhs34_
            d_178_v2_ = rhs35_
        elif True:
            d_245_v62_: _dafny.Map
            d_245_v62_ = _dafny.Map({d_176_v0_: d_177_v1_})
            d_246_v63_: int
            d_247_v64_: int
            d_248_v65_: int
            d_249_v66_: _dafny.Set
            out32_: int
            out33_: int
            out34_: int
            out35_: _dafny.Set
            out32_, out33_, out34_, out35_ = default__.m0(d_245_v62_, d_178_v2_, d_175_globalState_)
            d_246_v63_ = out32_
            d_247_v64_ = out33_
            d_248_v65_ = out34_
            d_249_v66_ = out35_
            d_248_v65_ = (d_179_v3_) * (len(d_185_v9_))
            d_178_v2_ = not(((d_181_v5_ if d_178_v2_ else ((d_185_v9_)[d_178_v2_] if (d_178_v2_) in (d_185_v9_) else len(_dafny.SeqWithoutIsStrInference([d_178_v2_, d_178_v2_]))))) not in (default__.fm3(d_175_globalState_)))
            d_250_v67_: _dafny.Seq
            d_250_v67_ = _dafny.SeqWithoutIsStrInference([d_183_v7_, d_183_v7_, d_183_v7_])
            d_251_v68_: C0
            nw31_ = C0()
            nw31_.ctor__(default__.fm1((d_250_v67_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_179_v3_])), len(d_250_v67_))], d_177_v1_, d_225_v42_, d_175_globalState_))
            d_251_v68_ = nw31_
            d_252_v69_: _dafny.Array
            nw32_ = _dafny.Array(_dafny.Seq({}), 9)
            d_252_v69_ = nw32_
            d_253_v70_: _dafny.Seq
            d_253_v70_ = _dafny.SeqWithoutIsStrInference([default__.fm2(d_178_v2_, d_175_globalState_)])
            index28_ = default__.safeIndex(525, (d_252_v69_).length(0))
            (d_252_v69_)[index28_] = d_253_v70_
            index29_ = default__.safeIndex(525, (d_252_v69_).length(0))
            (d_252_v69_)[index29_] = d_253_v70_
        d_254_v71_: C0
        nw33_ = C0()
        nw33_.ctor__(d_178_v2_)
        d_254_v71_ = nw33_
        d_255_v72_: _dafny.Seq
        d_255_v72_ = _dafny.SeqWithoutIsStrInference([True])
        hi0_ = (0) - ((len(_dafny.SeqWithoutIsStrInference([d_254_v71_])) if False else (0) - (len(d_255_v72_))))
        for d_256_i5_ in range(796, hi0_):
            d_257_v73_: _dafny.Map
            d_257_v73_ = _dafny.Map({d_176_v0_: d_180_v4_})
            d_258_v74_: int
            d_259_v75_: int
            d_260_v76_: int
            d_261_v77_: _dafny.Set
            out36_: int
            out37_: int
            out38_: int
            out39_: _dafny.Set
            out36_, out37_, out38_, out39_ = default__.m0((d_257_v73_) | (d_257_v73_), default__.fm1(_dafny.SeqWithoutIsStrInference([d_184_v8_ for d_262_i6_ in range(default__.abs(-795))]), d_179_v3_, d_225_v42_, d_175_globalState_), d_175_globalState_)
            d_258_v74_ = out36_
            d_259_v75_ = out37_
            d_260_v76_ = out38_
            d_261_v77_ = out39_
            d_263_v78_: int
            d_264_v79_: int
            d_265_v80_: int
            d_266_v81_: _dafny.Set
            out40_: int
            out41_: int
            out42_: int
            out43_: _dafny.Set
            out40_, out41_, out42_, out43_ = default__.m0((d_257_v73_ if d_178_v2_ else d_257_v73_), (d_225_v42_).isdisjoint(_dafny.Set({d_178_v2_, d_178_v2_, False, d_178_v2_, d_178_v2_})), d_175_globalState_)
            d_263_v78_ = out40_
            d_264_v79_ = out41_
            d_265_v80_ = out42_
            d_266_v81_ = out43_
            d_178_v2_ = not(False)
            d_267_v82_: _dafny.Seq
            d_267_v82_ = _dafny.SeqWithoutIsStrInference([d_254_v71_, d_254_v71_])
            d_268_v83_: _dafny.Array
            nw34_ = _dafny.Array(None, 3)
            nw34_[int(0)] = (d_267_v82_).set(default__.safeIndex(d_264_v79_, len(d_267_v82_)), d_254_v71_)
            nw34_[int(1)] = d_267_v82_
            nw34_[int(2)] = d_267_v82_
            d_268_v83_ = nw34_
            index30_ = default__.safeIndex(180, (d_268_v83_).length(0))
            (d_268_v83_)[index30_] = d_267_v82_
            d_269_v84_: _dafny.Map
            d_269_v84_ = _dafny.Map({d_178_v2_: d_178_v2_})
            d_270_v85_: D1
            d_270_v85_ = D1_DC3(_dafny.MultiSet([d_178_v2_, ((d_269_v84_)[d_178_v2_] if (d_178_v2_) in (d_269_v84_) else d_178_v2_)]), True, d_178_v2_)
            d_271_v86_: _dafny.Set
            d_271_v86_ = _dafny.Set({38})
            d_272_v87_: _dafny.Array
            nw35_ = _dafny.Array(None, 7)
            nw35_[int(0)] = 826
            nw35_[int(1)] = d_179_v3_
            nw35_[int(2)] = ((d_185_v9_)[(d_270_v85_).cf4] if ((d_270_v85_).cf4) in (d_185_v9_) else d_260_v76_)
            nw35_[int(3)] = len(d_271_v86_)
            nw35_[int(4)] = d_179_v3_
            nw35_[int(5)] = d_260_v76_
            nw35_[int(6)] = d_265_v80_
            d_272_v87_ = nw35_
            d_273_v88_: _dafny.Map
            d_273_v88_ = _dafny.Map({d_272_v87_: d_267_v82_})
            d_274_v89_: _dafny.Seq
            d_274_v89_ = _dafny.SeqWithoutIsStrInference([d_272_v87_])
            index31_ = default__.safeIndex(180, (d_268_v83_).length(0))
            (d_268_v83_)[index31_] = ((d_273_v88_)[(d_274_v89_)[default__.safeIndex(len(d_183_v7_), len(d_274_v89_))]] if ((d_274_v89_)[default__.safeIndex(len(d_183_v7_), len(d_274_v89_))]) in (d_273_v88_) else (_dafny.SeqWithoutIsStrInference([d_254_v71_])) + (d_267_v82_))
        d_275_v90_: C0
        nw36_ = C0()
        nw36_.ctor__((len(d_255_v72_)) < (len((_dafny.SeqWithoutIsStrInference([(d_255_v72_)[default__.safeIndex(-989, len(d_255_v72_))]])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_180_v4_, d_177_v1_])), len(_dafny.SeqWithoutIsStrInference([(d_255_v72_)[default__.safeIndex(-989, len(d_255_v72_))]]))), d_178_v2_))))
        d_275_v90_ = nw36_
        hi1_ = d_181_v5_
        for d_276_i7_ in range(default__.fm2(d_178_v2_, d_175_globalState_), hi1_):
            d_277_v91_: _dafny.Seq
            d_277_v91_ = _dafny.SeqWithoutIsStrInference([len(d_183_v7_)])
            d_278_v92_: _dafny.Map
            d_278_v92_ = _dafny.Map({d_176_v0_: len(d_277_v91_)})
            d_279_v93_: int
            d_280_v94_: int
            d_281_v95_: int
            d_282_v96_: _dafny.Set
            out44_: int
            out45_: int
            out46_: int
            out47_: _dafny.Set
            out44_, out45_, out46_, out47_ = default__.m0((d_278_v92_) | (d_278_v92_), d_178_v2_, d_175_globalState_)
            d_279_v93_ = out44_
            d_280_v94_ = out45_
            d_281_v95_ = out46_
            d_282_v96_ = out47_
            d_178_v2_ = d_178_v2_
            d_283_v97_: _dafny.Array
            nw37_ = _dafny.Array(_dafny.Seq({}), 15)
            d_283_v97_ = nw37_
            d_284_v98_: _dafny.Seq
            d_284_v98_ = _dafny.SeqWithoutIsStrInference([d_254_v71_, d_275_v90_])
            index32_ = default__.safeIndex(935, (d_283_v97_).length(0))
            (d_283_v97_)[index32_] = (d_284_v98_) + (d_284_v98_)
            d_285_v99_: _dafny.Seq
            d_285_v99_ = _dafny.SeqWithoutIsStrInference([d_254_v71_])
            index33_ = default__.safeIndex(935, (d_283_v97_).length(0))
            (d_283_v97_)[index33_] = (d_284_v98_) + ((d_285_v99_))
            d_254_v71_ = d_254_v71_
        pat_let_tv2_ = d_181_v5_
        def lambda9_(source3_):
            d_286___mcc_h0_ = source3_
            d_287_cf0_ = d_286___mcc_h0_
            return (pat_let_tv2_) >= (len(_dafny.Map({45: d_287_cf0_})))

        if lambda9_((d_184_v8_ if d_178_v2_ else d_184_v8_)):
            d_288_v100_: T0
            nw38_ = C0()
            nw38_.ctor__(d_178_v2_)
            d_288_v100_ = nw38_
            d_289_v101_: _dafny.Map
            d_289_v101_ = _dafny.Map({False: d_288_v100_})
            d_290_v102_: _dafny.Array
            nw39_ = _dafny.Array(None, 18)
            nw39_[int(0)] = d_288_v100_
            nw39_[int(1)] = d_288_v100_
            nw39_[int(2)] = d_288_v100_
            nw39_[int(3)] = d_288_v100_
            nw39_[int(4)] = d_288_v100_
            nw39_[int(5)] = d_288_v100_
            nw39_[int(6)] = d_288_v100_
            nw39_[int(7)] = d_288_v100_
            nw39_[int(8)] = d_288_v100_
            nw39_[int(9)] = d_288_v100_
            nw39_[int(10)] = d_288_v100_
            nw39_[int(11)] = d_288_v100_
            nw39_[int(12)] = d_288_v100_
            nw39_[int(13)] = (d_288_v100_ if True else d_288_v100_)
            nw39_[int(14)] = d_288_v100_
            nw39_[int(15)] = ((d_289_v101_)[not(d_288_v100_.f0)] if (not(d_288_v100_.f0)) in (d_289_v101_) else d_288_v100_)
            nw39_[int(16)] = d_288_v100_
            nw39_[int(17)] = d_288_v100_
            d_290_v102_ = nw39_
            rhs36_ = d_290_v102_
            rhs37_ = (d_288_v100_.f0 if (not(d_288_v100_.f0) if d_288_v100_.f0 else d_288_v100_.f0) else d_288_v100_.f0)
            d_290_v102_ = rhs36_
            d_178_v2_ = rhs37_
            d_180_v4_ = d_180_v4_
            d_291_v103_: _dafny.Seq
            d_291_v103_ = _dafny.SeqWithoutIsStrInference([d_183_v7_, d_183_v7_, d_183_v7_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hucmr"))])
            d_292_v104_: _dafny.Seq
            d_292_v104_ = _dafny.SeqWithoutIsStrInference([len(d_255_v72_), 479])
            d_293_v105_: _dafny.Map
            d_293_v105_ = _dafny.Map({(d_291_v103_) + ((d_291_v103_).set(default__.safeIndex((d_292_v104_)[default__.safeIndex(d_180_v4_, len(d_292_v104_))], len(d_291_v103_)), _dafny.SeqWithoutIsStrInference([d_184_v8_ for d_294_i8_ in range(default__.abs(5))]))): d_184_v8_})
            d_293_v105_ = (d_293_v105_).set((_dafny.SeqWithoutIsStrInference([default__.fm7(d_179_v3_, d_180_v4_, d_175_globalState_), d_183_v7_, d_183_v7_])) + (_dafny.SeqWithoutIsStrInference([d_183_v7_])), _dafny.CodePoint('e'))
            d_295_v106_: _dafny.Array
            def lambda10_(d_296_i9_):
                return (d_296_i9_) * (439)

            init5_ = lambda10_
            nw40_ = _dafny.Array(None, 7)
            for i0_5_ in range(nw40_.length(0)):
                nw40_[i0_5_] = init5_(i0_5_)
            d_295_v106_ = nw40_
            index34_ = default__.safeIndex(858, (d_295_v106_).length(0))
            (d_295_v106_)[index34_] = d_180_v4_
            d_297_v107_: int
            d_297_v107_ = d_179_v3_
            d_298_v108_: _dafny.Set
            d_298_v108_ = _dafny.Set({(d_177_v1_ if d_288_v100_.f0 else d_297_v107_), d_297_v107_})
            index35_ = default__.safeIndex(858, (d_295_v106_).length(0))
            (d_295_v106_)[index35_] = len(d_298_v108_)
            d_299_v109_: _dafny.Seq
            d_299_v109_ = d_183_v7_
            d_183_v7_ = ((d_299_v109_ if False else default__.fm8(d_178_v2_, not(d_178_v2_), d_177_v1_, d_175_globalState_)))
        elif True:
            d_300_v110_: C0
            nw41_ = C0()
            nw41_.ctor__((d_180_v4_) > (d_177_v1_))
            d_300_v110_ = nw41_
            d_301_v111_: D1
            d_301_v111_ = D1_DC3(_dafny.MultiSet([d_178_v2_, d_178_v2_]), d_178_v2_, d_178_v2_)
            d_302_v112_: D6
            d_302_v112_ = D6_DC10(d_275_v90_)
            d_303_v113_: _dafny.Seq
            d_303_v113_ = _dafny.SeqWithoutIsStrInference([d_254_v71_, d_275_v90_, (d_254_v71_ if d_178_v2_ else d_254_v71_), d_275_v90_, (d_300_v110_ if (d_301_v111_).cf4 else (d_302_v112_).cf14)])
            d_304_v114_: _dafny.Array
            nw42_ = _dafny.Array(_dafny.Set({}), 17)
            d_304_v114_ = nw42_
            d_305_v115_: _dafny.Map
            d_305_v115_ = _dafny.Map({d_177_v1_: d_178_v2_})
            d_306_v116_: _dafny.Set
            d_306_v116_ = _dafny.Set({d_305_v115_})
            index36_ = default__.safeIndex(537, (d_304_v114_).length(0))
            (d_304_v114_)[index36_] = d_306_v116_
            d_307_v117_: T0
            nw43_ = C0()
            nw43_.ctor__(d_178_v2_)
            d_307_v117_ = nw43_
            d_308_v118_: _dafny.Map
            d_308_v118_ = _dafny.Map({d_183_v7_: d_307_v117_})
            index37_ = default__.safeIndex(537, (d_304_v114_).length(0))
            rhs38_ = 721
            rhs39_ = ((d_303_v113_) + ((d_303_v113_) + (d_303_v113_))).set(default__.safeIndex(len(d_308_v118_), len((d_303_v113_) + ((d_303_v113_) + (d_303_v113_)))), d_300_v110_)
            rhs40_ = default__.fm9(len((d_183_v7_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icsjbibk")))), len((d_255_v72_).set(default__.safeIndex(d_180_v4_, len(d_255_v72_)), d_307_v117_.f0)), (len(_dafny.Map({d_307_v117_.f0: -841}))) + ((d_307_v117_).fm5(d_179_v3_, True, d_175_globalState_)), d_184_v8_, d_175_globalState_)
            lhs10_ = d_304_v114_
            lhs11_ = default__.safeIndex(537, (d_304_v114_).length(0))
            d_180_v4_ = rhs38_
            d_303_v113_ = rhs39_
            lhs10_[lhs11_] = rhs40_
            d_309_v119_: _dafny.MultiSet
            d_309_v119_ = _dafny.MultiSet([d_178_v2_])
            d_310_v120_: _dafny.MultiSet
            d_310_v120_ = _dafny.MultiSet([d_309_v119_])
            d_180_v4_ = (d_310_v120_).cardinality
            d_177_v1_ = default__.fm2((d_307_v117_.f0) and (not(False)), d_175_globalState_)
            if d_178_v2_:
                d_311_v121_: _dafny.Seq
                d_311_v121_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({-218, d_180_v4_, -520, (D6_DC11(d_177_v1_)).cf15, d_177_v1_})])
                (d_307_v117_).f0 = (d_311_v121_) == (d_311_v121_)
                d_312_v122_: _dafny.Seq
                d_312_v122_ = _dafny.SeqWithoutIsStrInference([d_179_v3_, (0) - (d_179_v3_)])
                d_312_v122_ = d_312_v122_
                d_184_v8_ = _dafny.CodePoint('o')
                d_313_v123_: C0
                nw44_ = C0()
                nw44_.ctor__(d_178_v2_)
                d_313_v123_ = nw44_
                d_178_v2_ = d_178_v2_
            elif True:
                d_314_v124_: C0
                nw45_ = C0()
                nw45_.ctor__(d_307_v117_.f0)
                d_314_v124_ = nw45_
                d_315_v125_: C0
                nw46_ = C0()
                nw46_.ctor__(d_178_v2_)
                d_315_v125_ = nw46_
                d_316_v126_: _dafny.Array
                nw47_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_316_v126_ = nw47_
                index38_ = default__.safeIndex(692, (d_316_v126_).length(0))
                (d_316_v126_)[index38_] = _dafny.SeqWithoutIsStrInference([d_184_v8_ for d_317_i10_ in range(default__.abs(347))])
                index39_ = default__.safeIndex(692, (d_316_v126_).length(0))
                (d_316_v126_)[index39_] = d_183_v7_
                (d_307_v117_).f0 = d_307_v117_.f0
                d_318_v127_: C0
                nw48_ = C0()
                nw48_.ctor__(d_178_v2_)
                d_318_v127_ = nw48_
        d_319_v128_: C0
        nw49_ = C0()
        nw49_.ctor__(d_178_v2_)
        d_319_v128_ = nw49_
        if (d_177_v1_) < ((d_180_v4_) - (d_181_v5_)):
            d_320_v129_: _dafny.Map
            d_320_v129_ = _dafny.Map({d_178_v2_: d_178_v2_})
            d_321_v130_: _dafny.MultiSet
            d_321_v130_ = _dafny.MultiSet([d_320_v129_])
            d_178_v2_ = (d_254_v71_).fm6((0) - (d_181_v5_), (d_321_v130_) | (d_321_v130_), (0) - ((0) - (d_180_v4_)), d_175_globalState_)
            rhs41_ = (d_186_v10_)
            rhs42_ = _dafny.SeqWithoutIsStrInference([d_184_v8_ for d_322_i11_ in range(default__.abs(-426))])
            rhs43_ = 406
            d_184_v8_ = rhs41_
            d_183_v7_ = rhs42_
            d_180_v4_ = rhs43_
            d_323_v131_: _dafny.Map
            d_323_v131_ = _dafny.Map({default__.fm0(d_183_v7_, d_179_v3_, len(d_185_v9_), d_175_globalState_): len(d_183_v7_)})
            d_324_v132_: _dafny.Set
            d_324_v132_ = _dafny.Set({d_323_v131_})
            d_324_v132_ = d_324_v132_
            d_178_v2_ = d_178_v2_
            d_325_v133_: _dafny.Map
            d_325_v133_ = _dafny.Map({d_176_v0_: 940})
            d_326_v134_: _dafny.MultiSet
            d_326_v134_ = _dafny.MultiSet([d_178_v2_, d_178_v2_])
            pat_let_tv3_ = d_326_v134_
            d_327_v135_: int
            d_328_v136_: int
            d_329_v137_: int
            d_330_v138_: _dafny.Set
            out48_: int
            out49_: int
            out50_: int
            out51_: _dafny.Set
            def iife21_(_pat_let4_0):
                def iife22_(d_331_dt__update__tmp_h1_):
                    def iife23_(_pat_let5_0):
                        def iife24_(d_332_dt__update_hcf3_h0_):
                            return D1_DC3(d_332_dt__update_hcf3_h0_, (d_331_dt__update__tmp_h1_).cf4, (d_331_dt__update__tmp_h1_).cf5)
                        return iife24_(_pat_let5_0)
                    return iife23_(pat_let_tv3_)
                return iife22_(_pat_let4_0)
            out48_, out49_, out50_, out51_ = default__.m0(d_325_v133_, (iife21_(D1_DC3(d_326_v134_, not(d_178_v2_), d_178_v2_))).cf5, d_175_globalState_)
            d_327_v135_ = out48_
            d_328_v136_ = out49_
            d_329_v137_ = out50_
            d_330_v138_ = out51_
        elif True:
            d_333_v139_: _dafny.Array
            nw50_ = _dafny.Array(_dafny.Map({}), 16)
            d_333_v139_ = nw50_
            def lambda11_(d_334_v2_):
                def lambda12_(d_335_i12_):
                    return (_dafny.Map({d_334_v2_: d_334_v2_})) | (_dafny.Map({d_334_v2_: d_334_v2_}))

                return lambda12_

            init6_ = lambda11_(d_178_v2_)
            nw51_ = _dafny.Array(None, 2)
            for i0_6_ in range(nw51_.length(0)):
                nw51_[i0_6_] = init6_(i0_6_)
            d_333_v139_ = nw51_
            d_336_v140_: _dafny.Seq
            d_336_v140_ = _dafny.SeqWithoutIsStrInference([d_225_v42_])
            d_337_v141_: _dafny.Map
            d_337_v141_ = _dafny.Map({d_181_v5_: d_336_v140_})
            d_337_v141_ = (d_337_v141_).set(default__.safeModuloInt(d_180_v4_, d_180_v4_), (d_336_v140_) + (d_336_v140_))
            d_338_v142_: _dafny.Map
            d_338_v142_ = _dafny.Map({d_179_v3_: d_183_v7_})
            d_339_v143_: _dafny.Seq
            d_339_v143_ = _dafny.SeqWithoutIsStrInference([d_183_v7_, d_183_v7_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffchstvvi"))) + (d_183_v7_)])
            rhs44_ = d_255_v72_
            rhs45_ = d_338_v142_
            rhs46_ = (d_319_v128_).fm5(d_181_v5_, False, d_175_globalState_)
            rhs47_ = d_183_v7_
            rhs48_ = (_dafny.SeqWithoutIsStrInference([d_183_v7_ for d_340_i13_ in range(default__.abs(-882))])) + (d_339_v143_)
            d_255_v72_ = rhs44_
            d_338_v142_ = rhs45_
            d_180_v4_ = rhs46_
            d_183_v7_ = rhs47_
            d_339_v143_ = rhs48_
            d_275_v90_ = d_275_v90_
            d_341_v144_: _dafny.Array
            def lambda13_(d_342_v12_):
                def lambda14_(d_343_i14_):
                    return default__.safeDivisionInt(d_343_i14_, len(d_342_v12_))

                return lambda14_

            init7_ = lambda13_(d_190_v12_)
            nw52_ = _dafny.Array(None, 4)
            for i0_7_ in range(nw52_.length(0)):
                nw52_[i0_7_] = init7_(i0_7_)
            d_341_v144_ = nw52_
            index40_ = default__.safeIndex(620, (d_341_v144_).length(0))
            (d_341_v144_)[index40_] = ((0) - (d_180_v4_) if False else d_181_v5_)
            d_344_v145_: _dafny.MultiSet
            d_344_v145_ = _dafny.MultiSet([d_183_v7_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stnwqxv"))), _dafny.SeqWithoutIsStrInference([d_184_v8_ for d_345_i15_ in range(default__.abs(61))])])
            index41_ = default__.safeIndex(620, (d_341_v144_).length(0))
            rhs49_ = (d_344_v145_).issubset((d_344_v145_) - (d_344_v145_))
            rhs50_ = (d_181_v5_) * (d_179_v3_)
            rhs51_ = (d_275_v90_).fm5(d_180_v4_, not(d_178_v2_), d_175_globalState_)
            lhs12_ = d_341_v144_
            lhs13_ = default__.safeIndex(620, (d_341_v144_).length(0))
            d_178_v2_ = rhs49_
            lhs12_[lhs13_] = rhs50_
            d_181_v5_ = rhs51_
        d_346_i16_: int
        d_346_i16_ = 0
        with _dafny.label("3"):
            while (d_255_v72_)[default__.safeIndex(d_177_v1_, len(d_255_v72_))]:
                with _dafny.c_label("3"):
                    if (d_346_i16_) >= (100):
                        raise _dafny.Break("3")
                    d_346_i16_ = (d_346_i16_) + (1)
                    d_319_v128_ = d_275_v90_
                    d_347_v146_: _dafny.Map
                    d_347_v146_ = _dafny.Map({(d_180_v4_) * ((d_319_v128_).fm5(len(d_225_v42_), d_178_v2_, d_175_globalState_)): d_178_v2_})
                    d_347_v146_ = d_347_v146_
                    d_348_v147_: _dafny.MultiSet
                    d_348_v147_ = _dafny.MultiSet([d_180_v4_, 122])
                    d_349_v148_: _dafny.Seq
                    d_349_v148_ = _dafny.SeqWithoutIsStrInference([176, d_181_v5_, -170, default__.fm2(default__.fm1(d_183_v7_, 943, _dafny.Set({d_178_v2_}), d_175_globalState_), d_175_globalState_)])
                    rhs52_ = (_dafny.MultiSet(d_349_v148_)).issubset(d_348_v147_)
                    rhs53_ = d_184_v8_
                    rhs54_ = (default__.safeDivisionInt(d_181_v5_, d_181_v5_) if (not(False)) or (d_178_v2_) else len(d_183_v7_))
                    d_178_v2_ = rhs52_
                    d_184_v8_ = rhs53_
                    d_180_v4_ = rhs54_
                    d_181_v5_ = len((d_183_v7_))
                    pass
            pass
        d_181_v5_ = (0) - (d_177_v1_)
        d_350_v149_: _dafny.MultiSet
        d_350_v149_ = _dafny.MultiSet([_dafny.Map({d_178_v2_: d_178_v2_})])
        d_351_v150_: _dafny.Map
        d_351_v150_ = _dafny.Map({d_178_v2_: d_178_v2_})
        if (d_319_v128_).fm6(len(d_183_v7_), (d_350_v149_) | ((d_350_v149_).set(d_351_v150_, default__.abs(d_180_v4_))), 342, d_175_globalState_):
            rhs55_ = d_176_v0_
            rhs56_ = (0) - ((d_181_v5_) + (d_181_v5_))
            rhs57_ = d_178_v2_
            d_176_v0_ = rhs55_
            d_179_v3_ = rhs56_
            d_178_v2_ = rhs57_
            d_183_v7_ = _dafny.SeqWithoutIsStrInference([d_184_v8_ for d_352_i17_ in range(default__.abs(522))])
            if d_178_v2_:
                d_176_v0_ = d_176_v0_
                d_179_v3_ = default__.safeDivisionInt(-845, d_180_v4_)
                d_183_v7_ = d_183_v7_
                d_353_v151_: C0
                nw53_ = C0()
                nw53_.ctor__(not(d_178_v2_))
                d_353_v151_ = nw53_
                d_354_v152_: _dafny.Array
                def lambda15_(d_355_v10_):
                    def lambda16_(d_356_i18_):
                        return d_355_v10_

                    return lambda16_

                init8_ = lambda15_(d_186_v10_)
                nw54_ = _dafny.Array(None, 16)
                for i0_8_ in range(nw54_.length(0)):
                    nw54_[i0_8_] = init8_(i0_8_)
                d_354_v152_ = nw54_
                index42_ = default__.safeIndex(273, (d_354_v152_).length(0))
                (d_354_v152_)[index42_] = d_186_v10_
                d_357_v153_: _dafny.Seq
                d_357_v153_ = _dafny.SeqWithoutIsStrInference([d_181_v5_])
                d_358_v154_: _dafny.Seq
                d_358_v154_ = _dafny.SeqWithoutIsStrInference([d_357_v153_])
                index43_ = default__.safeIndex(273, (d_354_v152_).length(0))
                (d_354_v152_)[index43_] = d_184_v8_
            elif True:
                index44_ = default__.safeIndex(457, (d_176_v0_).length(0))
                (d_176_v0_)[index44_] = d_178_v2_
                index45_ = default__.safeIndex(457, (d_176_v0_).length(0))
                (d_176_v0_)[index45_] = d_178_v2_
                d_359_v155_: C0
                nw55_ = C0()
                nw55_.ctor__(not ((d_176_v0_)[default__.safeIndex(457, (d_176_v0_).length(0))]) or (d_178_v2_))
                d_359_v155_ = nw55_
                d_360_v156_: _dafny.Set
                d_360_v156_ = _dafny.Set({d_184_v8_})
                d_361_v157_: _dafny.Map
                d_361_v157_ = _dafny.Map({((d_255_v72_)[default__.safeIndex(d_177_v1_, len(d_255_v72_))] if d_178_v2_ else d_178_v2_): (_dafny.Set({d_184_v8_, d_184_v8_, d_184_v8_})) | (d_360_v156_)})
                d_362_v158_: _dafny.Set
                d_362_v158_ = _dafny.Set({d_181_v5_})
                d_363_v159_: _dafny.MultiSet
                d_363_v159_ = _dafny.MultiSet([(d_176_v0_)[default__.safeIndex(457, (d_176_v0_).length(0))]])
                d_364_v160_: D1
                d_364_v160_ = D1_DC3(d_363_v159_, d_178_v2_, d_178_v2_)
                index46_ = default__.safeIndex(457, (d_176_v0_).length(0))
                rhs58_ = default__.safeDivisionInt(d_177_v1_, (-967) - (d_177_v1_))
                rhs59_ = (d_178_v2_ if d_178_v2_ else (d_176_v0_)[default__.safeIndex(457, (d_176_v0_).length(0))])
                rhs60_ = default__.safeModuloInt(d_180_v4_, (len(d_362_v158_)) * (len(_dafny.Map({d_183_v7_: (d_364_v160_).cf4}))))
                rhs61_ = (d_361_v157_) | ((default__.fm11(d_175_globalState_)) | (d_361_v157_))
                lhs14_ = d_176_v0_
                lhs15_ = default__.safeIndex(457, (d_176_v0_).length(0))
                d_180_v4_ = rhs58_
                lhs14_[lhs15_] = rhs59_
                d_181_v5_ = rhs60_
                d_361_v157_ = rhs61_
                d_365_v161_: _dafny.Array
                nw56_ = _dafny.Array(False, 15)
                d_365_v161_ = nw56_
                d_366_v162_: _dafny.Map
                d_366_v162_ = _dafny.Map({d_365_v161_: (d_359_v155_).fm5(d_181_v5_, (d_176_v0_)[default__.safeIndex(457, (d_176_v0_).length(0))], d_175_globalState_)})
                d_367_v163_: _dafny.Map
                d_367_v163_ = _dafny.Map({d_319_v128_: d_366_v162_})
                d_368_v164_: int
                d_369_v165_: int
                d_370_v166_: int
                d_371_v167_: _dafny.Set
                out52_: int
                out53_: int
                out54_: int
                out55_: _dafny.Set
                out52_, out53_, out54_, out55_ = default__.m0(((d_367_v163_)[d_275_v90_] if (d_275_v90_) in (d_367_v163_) else d_366_v162_), d_178_v2_, d_175_globalState_)
                d_368_v164_ = out52_
                d_369_v165_ = out53_
                d_370_v166_ = out54_
                d_371_v167_ = out55_
                d_372_v168_: _dafny.Array
                def lambda17_(d_373_v72_):
                    def lambda18_(d_374_i19_):
                        return d_373_v72_

                    return lambda18_

                init9_ = lambda17_(d_255_v72_)
                nw57_ = _dafny.Array(None, 3)
                for i0_9_ in range(nw57_.length(0)):
                    nw57_[i0_9_] = init9_(i0_9_)
                d_372_v168_ = nw57_
                d_375_v169_: _dafny.MultiSet
                d_375_v169_ = _dafny.MultiSet([d_319_v128_])
                index47_ = default__.safeIndex(457, (d_176_v0_).length(0))
                rhs62_ = ((_dafny.MultiSet([d_319_v128_, d_319_v128_])) - (d_375_v169_)).issubset(d_375_v169_)
                rhs63_ = (d_372_v168_ if not((d_176_v0_)[default__.safeIndex(457, (d_176_v0_).length(0))]) else d_372_v168_)
                rhs64_ = d_180_v4_
                lhs16_ = d_176_v0_
                lhs17_ = default__.safeIndex(457, (d_176_v0_).length(0))
                lhs16_[lhs17_] = rhs62_
                d_372_v168_ = rhs63_
                d_179_v3_ = rhs64_
            if (d_319_v128_).fm6(len(d_225_v42_), d_350_v149_, d_180_v4_, d_175_globalState_):
                d_181_v5_ = 826
                d_376_v170_: _dafny.MultiSet
                d_376_v170_ = _dafny.MultiSet([d_180_v4_])
                d_377_v171_: _dafny.Map
                d_377_v171_ = _dafny.Map({d_178_v2_: d_176_v0_})
                d_378_v172_: _dafny.Array
                nw58_ = _dafny.Array(None, 18)
                nw58_[int(0)] = (d_181_v5_) - (d_177_v1_)
                nw58_[int(1)] = len(d_190_v12_)
                nw58_[int(2)] = ((_dafny.MultiSet([d_180_v4_])).set(d_179_v3_, default__.abs(d_180_v4_))).cardinality
                nw58_[int(3)] = d_181_v5_
                nw58_[int(4)] = default__.safeDivisionInt(len(d_183_v7_), len(d_225_v42_))
                nw58_[int(5)] = d_180_v4_
                nw58_[int(6)] = (0) - ((d_179_v3_) * (d_181_v5_))
                nw58_[int(7)] = ((0) - ((d_376_v170_).cardinality) if d_178_v2_ else d_181_v5_)
                nw58_[int(8)] = len((_dafny.Map({d_178_v2_: d_176_v0_})) | (d_377_v171_))
                nw58_[int(9)] = d_179_v3_
                nw58_[int(10)] = d_180_v4_
                nw58_[int(11)] = d_180_v4_
                nw58_[int(12)] = d_180_v4_
                nw58_[int(13)] = -876
                nw58_[int(14)] = d_181_v5_
                nw58_[int(15)] = len(_dafny.Set({d_178_v2_, d_178_v2_, d_178_v2_}))
                nw58_[int(16)] = (0) - ((d_275_v90_).fm5(947, d_178_v2_, d_175_globalState_))
                nw58_[int(17)] = d_180_v4_
                d_378_v172_ = nw58_
                index48_ = default__.safeIndex(493, (d_378_v172_).length(0))
                (d_378_v172_)[index48_] = d_181_v5_
                index49_ = default__.safeIndex(493, (d_378_v172_).length(0))
                (d_378_v172_)[index49_] = -73
                d_379_v173_: int
                d_380_v174_: int
                d_381_v175_: int
                d_382_v176_: _dafny.Set
                out56_: int
                out57_: int
                out58_: int
                out59_: _dafny.Set
                out56_, out57_, out58_, out59_ = default__.m0(_dafny.Map({d_176_v0_: d_181_v5_}), d_178_v2_, d_175_globalState_)
                d_379_v173_ = out56_
                d_380_v174_ = out57_
                d_381_v175_ = out58_
                d_382_v176_ = out59_
                d_383_v177_: _dafny.MultiSet
                d_383_v177_ = _dafny.MultiSet([d_178_v2_])
                d_179_v3_ = (d_380_v174_) * ((d_383_v177_).cardinality)
                d_177_v1_ = (d_319_v128_).fm5(((d_378_v172_)[default__.safeIndex(493, (d_378_v172_).length(0))]) * (d_381_v175_), False, d_175_globalState_)
            elif True:
                d_181_v5_ = (d_181_v5_) + (d_180_v4_)
                d_384_v178_: D3
                d_384_v178_ = D3_DC6((0) - (len(d_255_v72_)), d_180_v4_, d_178_v2_)
                d_178_v2_ = (d_384_v178_).cf10
                d_385_v179_: _dafny.Map
                d_385_v179_ = _dafny.Map({d_176_v0_: d_179_v3_})
                d_386_v180_: int
                d_387_v181_: int
                d_388_v182_: int
                d_389_v183_: _dafny.Set
                out60_: int
                out61_: int
                out62_: int
                out63_: _dafny.Set
                out60_, out61_, out62_, out63_ = default__.m0((d_385_v179_) | (d_385_v179_), ((d_351_v150_)[d_178_v2_] if (d_178_v2_) in (d_351_v150_) else d_178_v2_), d_175_globalState_)
                d_386_v180_ = out60_
                d_387_v181_ = out61_
                d_388_v182_ = out62_
                d_389_v183_ = out63_
                d_390_v185_: _dafny.Seq
                d_390_v185_ = d_183_v7_
                d_391_v186_: _dafny.Map
                d_391_v186_ = _dafny.Map({417: d_390_v185_})
                d_392_v187_: _dafny.Map
                d_392_v187_ = _dafny.Map({len(d_255_v72_): d_184_v8_})
                d_393_v188_: _dafny.Seq
                def iife25_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in (d_391_v186_).keys.Elements:
                        d_394_v184_: int = compr_13_
                        if (d_394_v184_) in (d_391_v186_):
                            coll13_[(d_394_v184_) * (d_387_v181_)] = d_184_v8_
                    return _dafny.Map(coll13_)
                d_393_v188_ = _dafny.SeqWithoutIsStrInference([iife25_()
                , (d_392_v187_) | (d_392_v187_), (d_392_v187_) | (d_392_v187_), d_392_v187_])
                d_395_v189_: _dafny.Array
                nw59_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_395_v189_ = nw59_
                d_396_v190_: _dafny.Array
                def lambda19_(d_397_v1_):
                    def lambda20_(d_398_i20_):
                        return default__.safeModuloInt(d_398_i20_, d_397_v1_)

                    return lambda20_

                init10_ = lambda19_(d_177_v1_)
                nw60_ = _dafny.Array(None, 17)
                for i0_10_ in range(nw60_.length(0)):
                    nw60_[i0_10_] = init10_(i0_10_)
                d_396_v190_ = nw60_
                index50_ = default__.safeIndex(693, (d_395_v189_).length(0))
                (d_395_v189_)[index50_] = d_396_v190_
                d_399_v191_: _dafny.Map
                d_399_v191_ = _dafny.Map({d_225_v42_: d_178_v2_})
                index51_ = default__.safeIndex(693, (d_395_v189_).length(0))
                rhs65_ = (d_393_v188_ if False else d_393_v188_)
                rhs66_ = d_275_v90_
                rhs67_ = d_396_v190_
                rhs68_ = (d_387_v181_) + ((0) - (len(d_183_v7_)))
                rhs69_ = (d_399_v191_) == (d_399_v191_)
                lhs18_ = d_395_v189_
                lhs19_ = default__.safeIndex(693, (d_395_v189_).length(0))
                d_393_v188_ = rhs65_
                d_254_v71_ = rhs66_
                lhs18_[lhs19_] = rhs67_
                d_388_v182_ = rhs68_
                d_178_v2_ = rhs69_
                rhs70_ = ((0) - (d_179_v3_)) + (229)
                rhs71_ = (d_384_v178_).cf8
                rhs72_ = d_184_v8_
                rhs73_ = d_176_v0_
                d_387_v181_ = rhs70_
                d_386_v180_ = rhs71_
                d_184_v8_ = rhs72_
                d_176_v0_ = rhs73_
            d_400_v192_: _dafny.Map
            d_400_v192_ = _dafny.Map({d_176_v0_: len(d_255_v72_)})
            d_401_v193_: int
            d_402_v194_: int
            d_403_v195_: int
            d_404_v196_: _dafny.Set
            out64_: int
            out65_: int
            out66_: int
            out67_: _dafny.Set
            out64_, out65_, out66_, out67_ = default__.m0(d_400_v192_, d_178_v2_, d_175_globalState_)
            d_401_v193_ = out64_
            d_402_v194_ = out65_
            d_403_v195_ = out66_
            d_404_v196_ = out67_
        elif True:
            d_178_v2_ = d_178_v2_
            d_405_v197_: _dafny.Array
            nw61_ = _dafny.Array(int(0), 4)
            d_405_v197_ = nw61_
            index52_ = default__.safeIndex(942, (d_405_v197_).length(0))
            (d_405_v197_)[index52_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odeh"))) + (d_183_v7_))
            index53_ = default__.safeIndex(942, (d_405_v197_).length(0))
            (d_405_v197_)[index53_] = d_181_v5_
            d_406_v198_: D7
            d_406_v198_ = D7_DC12(default__.fm12(True, d_184_v8_, d_178_v2_, d_177_v1_, d_175_globalState_))
            d_407_v199_: _dafny.Map
            d_407_v199_ = _dafny.Map({d_176_v0_: ((d_406_v198_).cf16).cardinality})
            d_408_v200_: int
            d_409_v201_: int
            d_410_v202_: int
            d_411_v203_: _dafny.Set
            out68_: int
            out69_: int
            out70_: int
            out71_: _dafny.Set
            out68_, out69_, out70_, out71_ = default__.m0(d_407_v199_, d_178_v2_, d_175_globalState_)
            d_408_v200_ = out68_
            d_409_v201_ = out69_
            d_410_v202_ = out70_
            d_411_v203_ = out71_
            d_178_v2_ = not(d_178_v2_)
            d_412_v204_: _dafny.Array
            nw62_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_412_v204_ = nw62_
            rhs74_ = d_177_v1_
            rhs75_ = d_412_v204_
            rhs76_ = not (d_178_v2_) or (d_178_v2_)
            rhs77_ = d_178_v2_
            d_177_v1_ = rhs74_
            d_412_v204_ = rhs75_
            d_178_v2_ = rhs76_
            d_178_v2_ = rhs77_
        index54_ = default__.safeIndex(901, (d_176_v0_).length(0))
        (d_176_v0_)[index54_] = d_178_v2_
        index55_ = default__.safeIndex(901, (d_176_v0_).length(0))
        (d_176_v0_)[index55_] = default__.fm1(d_183_v7_, len(d_183_v7_), d_225_v42_, d_175_globalState_)
        d_413_v205_: _dafny.Seq
        d_413_v205_ = _dafny.SeqWithoutIsStrInference([len(d_183_v7_), len(d_255_v72_)])
        d_414_v206_: _dafny.MultiSet
        d_414_v206_ = _dafny.MultiSet([d_413_v205_])
        d_415_v207_: D7
        d_415_v207_ = D7_DC12(d_414_v206_)
        d_416_v208_: D7
        d_416_v208_ = D7_DC12((d_415_v207_).cf16)
        d_417_v209_: D7
        d_417_v209_ = D7_DC14(d_416_v208_)
        pat_let_tv4_ = d_416_v208_
        def iife26_(_pat_let6_0):
            def iife27_(d_418_dt__update__tmp_h4_):
                def iife28_(_pat_let7_0):
                    def iife29_(d_419_dt__update_hcf18_h0_):
                        return D7_DC14(d_419_dt__update_hcf18_h0_)
                    return iife29_(_pat_let7_0)
                return iife28_(pat_let_tv4_)
            return iife27_(_pat_let6_0)
        source4_ = iife26_(d_417_v209_)
        if source4_.is_DC13:
            d_420___mcc_h1_ = source4_.cf17
            d_421_cf17_ = d_420___mcc_h1_
            d_180_v4_ = (0) - (d_179_v3_)
            d_421_cf17_ = d_421_cf17_
            d_422_v210_: _dafny.Map
            d_422_v210_ = _dafny.Map({d_421_cf17_: d_254_v71_})
            d_423_v211_: _dafny.Array
            nw63_ = _dafny.Array(None, 18)
            nw63_[int(0)] = d_254_v71_
            nw63_[int(1)] = d_275_v90_
            nw63_[int(2)] = d_275_v90_
            nw63_[int(3)] = d_275_v90_
            nw63_[int(4)] = d_319_v128_
            nw63_[int(5)] = d_254_v71_
            nw63_[int(6)] = d_254_v71_
            nw63_[int(7)] = d_319_v128_
            nw63_[int(8)] = d_254_v71_
            nw63_[int(9)] = d_254_v71_
            nw63_[int(10)] = d_319_v128_
            nw63_[int(11)] = ((d_422_v210_)[d_178_v2_] if (d_178_v2_) in (d_422_v210_) else d_319_v128_)
            nw63_[int(12)] = d_319_v128_
            nw63_[int(13)] = d_319_v128_
            nw63_[int(14)] = d_275_v90_
            nw63_[int(15)] = d_275_v90_
            nw63_[int(16)] = d_275_v90_
            nw63_[int(17)] = d_254_v71_
            d_423_v211_ = nw63_
            index56_ = default__.safeIndex(917, (d_423_v211_).length(0))
            (d_423_v211_)[index56_] = d_275_v90_
            index57_ = default__.safeIndex(917, (d_423_v211_).length(0))
            (d_423_v211_)[index57_] = d_319_v128_
            d_177_v1_ = (default__.safeDivisionInt(d_179_v3_, d_181_v5_)) * ((len(d_255_v72_)) - (283))
        elif source4_.is_DC12:
            d_424___mcc_h2_ = source4_.cf16
            d_425_cf16_ = d_424___mcc_h2_
            index58_ = default__.safeIndex(901, (d_176_v0_).length(0))
            rhs78_ = d_181_v5_
            rhs79_ = d_180_v4_
            rhs80_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xya"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbshdbhf")))) != (_dafny.SeqWithoutIsStrInference([d_184_v8_ for d_426_i21_ in range(default__.abs(-774))]))
            rhs81_ = d_184_v8_
            rhs82_ = d_178_v2_
            lhs20_ = d_176_v0_
            lhs21_ = default__.safeIndex(901, (d_176_v0_).length(0))
            d_180_v4_ = rhs78_
            d_179_v3_ = rhs79_
            d_178_v2_ = rhs80_
            d_184_v8_ = rhs81_
            lhs20_[lhs21_] = rhs82_
            d_427_v212_: _dafny.Set
            d_427_v212_ = _dafny.Set({d_179_v3_})
            index59_ = default__.safeIndex(901, (d_176_v0_).length(0))
            (d_176_v0_)[index59_] = (d_254_v71_).fm6(default__.safeDivisionInt(d_181_v5_, len(d_183_v7_)), d_350_v149_, (0) - ((d_181_v5_) - (len(d_427_v212_))), d_175_globalState_)
            d_428_v213_: _dafny.MultiSet
            d_428_v213_ = _dafny.MultiSet([d_178_v2_])
            d_429_v214_: D1
            d_429_v214_ = D1_DC3(d_428_v213_, d_178_v2_, (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))])
            source5_ = (d_429_v214_ if True else d_429_v214_)
            if source5_.is_DC2:
                d_430___mcc_h4_ = source5_.cf2
                d_431_cf2_ = d_430___mcc_h4_
                d_181_v5_ = (0) - ((0) - (d_180_v4_))
                d_432_v215_: int
                d_433_v216_: int
                d_434_v217_: int
                d_435_v218_: _dafny.Set
                out72_: int
                out73_: int
                out74_: int
                out75_: _dafny.Set
                out72_, out73_, out74_, out75_ = default__.m0(_dafny.Map({d_176_v0_: d_179_v3_}), d_178_v2_, d_175_globalState_)
                d_432_v215_ = out72_
                d_433_v216_ = out73_
                d_434_v217_ = out74_
                d_435_v218_ = out75_
                d_436_v219_: _dafny.Map
                d_436_v219_ = _dafny.Map({d_176_v0_: len(d_413_v205_)})
                d_437_v220_: int
                d_438_v221_: int
                d_439_v222_: int
                d_440_v223_: _dafny.Set
                out76_: int
                out77_: int
                out78_: int
                out79_: _dafny.Set
                out76_, out77_, out78_, out79_ = default__.m0((d_436_v219_) | (d_436_v219_), d_178_v2_, d_175_globalState_)
                d_437_v220_ = out76_
                d_438_v221_ = out77_
                d_439_v222_ = out78_
                d_440_v223_ = out79_
                d_441_v224_: D3
                d_441_v224_ = D3_DC6(default__.safeDivisionInt((d_428_v213_).cardinality, d_181_v5_), (225) - ((0) - (d_180_v4_)), (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))])
                pat_let_tv5_ = d_183_v7_
                pat_let_tv6_ = d_183_v7_
                def iife30_(_pat_let8_0):
                    def iife31_(d_442_dt__update__tmp_h5_):
                        def iife32_(_pat_let9_0):
                            def iife33_(d_443_dt__update_hcf9_h0_):
                                return D3_DC6((d_442_dt__update__tmp_h5_).cf8, d_443_dt__update_hcf9_h0_, (d_442_dt__update__tmp_h5_).cf10)
                            return iife33_(_pat_let9_0)
                        return iife32_(len((pat_let_tv5_) + (pat_let_tv6_)))
                    return iife31_(_pat_let8_0)
                d_441_v224_ = iife30_(d_441_v224_)
            elif source5_.is_DC3:
                d_444___mcc_h5_ = source5_.cf3
                d_445___mcc_h6_ = source5_.cf4
                d_446___mcc_h7_ = source5_.cf5
                d_447_cf5_ = d_446___mcc_h7_
                d_448_cf4_ = d_445___mcc_h6_
                d_449_cf3_ = d_444___mcc_h5_
                d_255_v72_ = (d_255_v72_) + (d_255_v72_)
                d_450_v225_: _dafny.Map
                d_450_v225_ = _dafny.Map({d_176_v0_: 691})
                d_451_v226_: _dafny.Seq
                d_451_v226_ = _dafny.SeqWithoutIsStrInference([d_254_v71_])
                d_452_v227_: _dafny.Map
                d_452_v227_ = _dafny.Map({d_451_v226_: (d_254_v71_).fm6(d_181_v5_, d_350_v149_, d_177_v1_, d_175_globalState_)})
                d_453_v228_: _dafny.Map
                d_453_v228_ = _dafny.Map({d_452_v227_: False})
                d_454_v229_: int
                d_455_v230_: int
                d_456_v231_: int
                d_457_v232_: _dafny.Set
                out80_: int
                out81_: int
                out82_: int
                out83_: _dafny.Set
                out80_, out81_, out82_, out83_ = default__.m0(d_450_v225_, ((d_453_v228_)[d_452_v227_] if (d_452_v227_) in (d_453_v228_) else (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]), d_175_globalState_)
                d_454_v229_ = out80_
                d_455_v230_ = out81_
                d_456_v231_ = out82_
                d_457_v232_ = out83_
                d_183_v7_ = (default__.fm7(d_177_v1_, d_454_v229_, d_175_globalState_)) + (d_183_v7_)
                d_180_v4_ = (0) - (d_180_v4_)
            elif True:
                d_458___mcc_h8_ = source5_.cf1
                d_459_cf1_ = d_458___mcc_h8_
                d_460_v233_: _dafny.Map
                d_460_v233_ = _dafny.Map({d_176_v0_: len(d_183_v7_)})
                d_461_v234_: _dafny.Map
                d_461_v234_ = _dafny.Map({d_275_v90_: d_180_v4_})
                d_462_v235_: int
                d_463_v236_: int
                d_464_v237_: int
                d_465_v238_: _dafny.Set
                out84_: int
                out85_: int
                out86_: int
                out87_: _dafny.Set
                out84_, out85_, out86_, out87_ = default__.m0((d_460_v233_) | ((d_460_v233_).set(d_176_v0_, d_180_v4_)), (d_177_v1_) < (len(d_461_v234_)), d_175_globalState_)
                d_462_v235_ = out84_
                d_463_v236_ = out85_
                d_464_v237_ = out86_
                d_465_v238_ = out87_
                d_466_v239_: int
                d_467_v240_: int
                d_468_v241_: int
                d_469_v242_: _dafny.Set
                out88_: int
                out89_: int
                out90_: int
                out91_: _dafny.Set
                out88_, out89_, out90_, out91_ = default__.m0(d_460_v233_, (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))], d_175_globalState_)
                d_466_v239_ = out88_
                d_467_v240_ = out89_
                d_468_v241_ = out90_
                d_469_v242_ = out91_
                d_225_v42_ = default__.fm13((d_183_v7_) <= (d_183_v7_), not((d_178_v2_) in (d_428_v213_)), False, d_175_globalState_)
                d_470_v243_: D6
                d_470_v243_ = D6_DC11(d_468_v241_)
                d_471_v244_: _dafny.Set
                d_471_v244_ = _dafny.Set({d_470_v243_, D6_DC11(d_467_v240_), d_470_v243_})
                d_472_v245_: T0
                nw64_ = C0()
                nw64_.ctor__(d_178_v2_)
                d_472_v245_ = nw64_
                d_473_v246_: _dafny.Seq
                d_473_v246_ = _dafny.SeqWithoutIsStrInference([d_472_v245_, d_472_v245_])
                d_474_v247_: _dafny.Array
                nw65_ = _dafny.Array(_dafny.Set({}), 2)
                d_474_v247_ = nw65_
                rhs83_ = d_471_v244_
                rhs84_ = (_dafny.SeqWithoutIsStrInference([d_472_v245_, d_472_v245_])) + ((_dafny.SeqWithoutIsStrInference([d_472_v245_])) + (d_473_v246_))
                rhs85_ = d_474_v247_
                d_471_v244_ = rhs83_
                d_473_v246_ = rhs84_
                d_474_v247_ = rhs85_
            d_475_v248_: _dafny.Map
            d_475_v248_ = _dafny.Map({d_178_v2_: d_176_v0_})
            d_475_v248_ = (d_475_v248_).set((_dafny.CodePoint('m')) in (d_183_v7_), d_176_v0_)
        elif True:
            d_476___mcc_h3_ = source4_.cf18
            d_477_cf18_ = d_476___mcc_h3_
            if (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]:
                d_180_v4_ = d_181_v5_
                d_478_v249_: _dafny.Set
                d_478_v249_ = _dafny.Set({d_413_v205_, d_413_v205_, d_413_v205_, d_413_v205_})
                index60_ = default__.safeIndex(901, (d_176_v0_).length(0))
                (d_176_v0_)[index60_] = not((d_478_v249_).issubset(default__.fm14(len(d_185_v9_), d_175_globalState_)))
                d_184_v8_ = d_184_v8_
                d_479_v250_: _dafny.Map
                d_479_v250_ = _dafny.Map({d_176_v0_: d_181_v5_})
                d_480_v251_: int
                d_481_v252_: int
                d_482_v253_: int
                d_483_v254_: _dafny.Set
                out92_: int
                out93_: int
                out94_: int
                out95_: _dafny.Set
                out92_, out93_, out94_, out95_ = default__.m0(d_479_v250_, (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))], d_175_globalState_)
                d_480_v251_ = out92_
                d_481_v252_ = out93_
                d_482_v253_ = out94_
                d_483_v254_ = out95_
                d_484_v255_: _dafny.Set
                d_484_v255_ = _dafny.Set({((d_185_v9_)[not(not(d_178_v2_))] if (not(not(d_178_v2_))) in (d_185_v9_) else (0) - (len(d_183_v7_)))})
                d_485_v256_: _dafny.Map
                d_485_v256_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([not((d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]), (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))], d_178_v2_]): len(d_484_v255_)})
                d_486_v257_: _dafny.Seq
                d_486_v257_ = _dafny.SeqWithoutIsStrInference([d_255_v72_])
                d_485_v256_ = (d_485_v256_).set(((d_486_v257_)[default__.safeIndex(d_480_v251_, len(d_486_v257_))]) + (d_255_v72_), d_181_v5_)
            elif True:
                index61_ = default__.safeIndex(901, (d_176_v0_).length(0))
                (d_176_v0_)[index61_] = d_178_v2_
                d_487_v258_: _dafny.MultiSet
                d_487_v258_ = _dafny.MultiSet([d_183_v7_, d_183_v7_])
                d_488_v259_: _dafny.Map
                d_488_v259_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_177_v1_ for d_489_i22_ in range(default__.abs(669))]): _dafny.SeqWithoutIsStrInference([d_184_v8_ for d_490_i23_ in range(default__.abs(23))])})
                nw66_ = _dafny.Array(None, 19)
                nw66_[int(0)] = (d_178_v2_ if (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))] else (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))])
                nw66_[int(1)] = d_178_v2_
                nw66_[int(2)] = (d_255_v72_) <= (_dafny.SeqWithoutIsStrInference([d_178_v2_]))
                nw66_[int(3)] = (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]
                nw66_[int(4)] = (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]
                nw66_[int(5)] = d_178_v2_
                nw66_[int(6)] = (d_255_v72_)[default__.safeIndex((d_487_v258_).cardinality, len(d_255_v72_))]
                nw66_[int(7)] = (True) == ((d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))])
                nw66_[int(8)] = d_178_v2_
                nw66_[int(9)] = d_178_v2_
                nw66_[int(10)] = (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]
                nw66_[int(11)] = (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]
                nw66_[int(12)] = d_178_v2_
                nw66_[int(13)] = d_178_v2_
                nw66_[int(14)] = (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]
                nw66_[int(15)] = default__.fm1(((d_488_v259_)[d_413_v205_] if (d_413_v205_) in (d_488_v259_) else d_183_v7_), d_179_v3_, d_225_v42_, d_175_globalState_)
                nw66_[int(16)] = (D7_DC13((d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))])).cf17
                nw66_[int(17)] = (d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))]
                nw66_[int(18)] = not((d_176_v0_)[default__.safeIndex(901, (d_176_v0_).length(0))])
                d_176_v0_ = nw66_
                d_178_v2_ = not(not((d_183_v7_) < (d_183_v7_)))
                d_184_v8_ = d_184_v8_
                d_491_v260_: _dafny.Set
                d_491_v260_ = _dafny.Set({324, d_177_v1_})
                d_492_v261_: C0
                nw67_ = C0()
                nw67_.ctor__((d_491_v260_).isdisjoint(d_491_v260_))
                d_492_v261_ = nw67_
            def lambda21_(d_493_i24_):
                return False

            init11_ = lambda21_
            nw68_ = _dafny.Array(None, 15)
            for i0_11_ in range(nw68_.length(0)):
                nw68_[i0_11_] = init11_(i0_11_)
            d_176_v0_ = nw68_
            d_494_v262_: _dafny.Array
            nw69_ = _dafny.Array(_dafny.Map({}), 20)
            d_494_v262_ = nw69_
            index62_ = default__.safeIndex(131, (d_494_v262_).length(0))
            (d_494_v262_)[index62_] = (d_351_v150_) | (d_351_v150_)
            index63_ = default__.safeIndex(131, (d_494_v262_).length(0))
            (d_494_v262_)[index63_] = (d_351_v150_) | (d_351_v150_)
            d_495_v263_: _dafny.Map
            d_495_v263_ = _dafny.Map({(d_413_v205_) + (d_413_v205_): d_181_v5_})
            d_495_v263_ = d_495_v263_
        _dafny.print(_dafny.string_of((d_176_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_177_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_178_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_179_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_180_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_181_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v6_) == (_dafny.Set({_dafny.Set({True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v7_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_184_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_v9_) == (_dafny.Map({True: 365}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v10_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v11_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v12_) == (_dafny.Map({0: 265}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v42_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_v72_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_346_i16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_350_v149_) == (_dafny.MultiSet([_dafny.Map({False: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_351_v150_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_413_v205_) == (_dafny.SeqWithoutIsStrInference([426, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_414_v206_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([426, 1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_415_v207_).cf16) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([426, 1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_416_v208_).cf16) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([426, 1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_417_v209_).cf18).cf16) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([426, 1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

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
        return lambda: D1_DC2(_dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: int(0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC4(D2, NamedTuple('DC4', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC6(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC6(D3, NamedTuple('DC6', [('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC5(D3, NamedTuple('DC5', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)

class D4_DC8(D4, NamedTuple('DC8', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D5_DC9)

class D5_DC9(D5, NamedTuple('DC9', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC9({self.cf13.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC9) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC11(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D6_DC10)

class D6_DC11(D6, NamedTuple('DC11', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC10(D6, NamedTuple('DC10', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC10({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC10) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC13(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D7_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D7_DC12)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)

class D7_DC13(D7, NamedTuple('DC13', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC13({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC13) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC12(D7, NamedTuple('DC12', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC12({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC12) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC16(False, _dafny.MultiSet({}), _dafny.Seq({}), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D8_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D8_DC15)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC16(D8, NamedTuple('DC16', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC16({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC16) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC17(D8, NamedTuple('DC17', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({self.cf24.VerbatimString(True)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC15(D8, NamedTuple('DC15', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC15({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC15) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC20(_dafny.Array(None, 0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)

class D9_DC20(D9, NamedTuple('DC20', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC19(D9, NamedTuple('DC19', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC22(False, int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)

class D10_DC22(D10, NamedTuple('DC22', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC23(D10, NamedTuple('DC23', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f0(self):
        return self._f0
    @f0.setter
    def f0(self, value):
        self._f0 = value

class GlobalState:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self):
        pass
        pass


class C0(T0):
    def  __init__(self):
        self._f0: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f0(self):
        return self._f0
    @f0.setter
    def f0(self, value):
        self._f0 = value
    def ctor__(self, f0):
        (self).f0 = f0

    def fm4(self, globalState):
        source6_ = D1_DC2(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({840: self.f0}))])}))
        if source6_.is_DC2:
            d_496___mcc_h0_ = source6_.cf2
            d_497_cf2_ = d_496___mcc_h0_
            return -299
        elif source6_.is_DC3:
            d_498___mcc_h1_ = source6_.cf3
            d_499___mcc_h2_ = source6_.cf4
            d_500___mcc_h3_ = source6_.cf5
            d_501_cf5_ = d_500___mcc_h3_
            d_502_cf4_ = d_499___mcc_h2_
            d_503_cf3_ = d_498___mcc_h1_
            return len(_dafny.Set({d_502_cf4_}))
        elif True:
            d_504___mcc_h4_ = source6_.cf1
            d_505_cf1_ = d_504___mcc_h4_
            def iife34_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(617, 586):
                    d_506_v0_: int = compr_14_
                    if ((617) <= (d_506_v0_)) and ((d_506_v0_) < (586)):
                        coll14_[(d_506_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vo"))))] = 129
                return _dafny.Map(coll14_)
            return (0) - (len(iife34_()
            ))

    def fm5(self, p0, p1, globalState):
        return (0) - ((39) * ((452)))

    def fm6(self, p0, p1, p2, globalState):
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([-394 for d_507_i0_ in range(default__.abs(222))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpmca"))) for d_508_i1_ in range(default__.abs(-784))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hs"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjhdka"))), -765]), _dafny.SeqWithoutIsStrInference([240 for d_509_i2_ in range(default__.abs(511))]), _dafny.SeqWithoutIsStrInference([815, 633, (len(_dafny.Set({-347}))), len(_dafny.Map({False: True})), 390])})).issubset(_dafny.Set({_dafny.SeqWithoutIsStrInference([583])}))

