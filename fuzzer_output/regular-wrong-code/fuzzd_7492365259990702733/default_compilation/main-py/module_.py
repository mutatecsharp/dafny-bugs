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
    def fm0(globalState):
        return (529) + ((64) - (len(_dafny.Set({(_dafny.MultiSet([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))})), 695])).cardinality, 768}))))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return ((True if False else True)) or (True)

    @staticmethod
    def fm2(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Map
            for compr_0_ in (_dafny.Map({_dafny.Map({len(_dafny.Set({_dafny.MultiSet([True, True]), _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True]))), _dafny.MultiSet([True])})): D0_DC2(True)}): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_0_i0_ in range(default__.abs(380))]))})).keys.Elements:
                d_1_v0_: _dafny.Map = compr_0_
                if (d_1_v0_) in (_dafny.Map({_dafny.Map({len(_dafny.Set({_dafny.MultiSet([True, True]), _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True]))), _dafny.MultiSet([True])})): D0_DC2(True)}): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_0_i0_ in range(default__.abs(380))]))})):
                    coll0_[d_1_v0_] = 329
            return _dafny.Map(coll0_)
        return ((iife0_()
        ) | (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): D0_DC2(False)}): len(_dafny.SeqWithoutIsStrInference([False, True, False, False, False]))}))) | ((_dafny.Map({_dafny.Map({-222: D0_DC2(True)}): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swxgyamao")))})) | (_dafny.Map({_dafny.Map({(_dafny.MultiSet([679])).cardinality: D0_DC2(True)}): len(_dafny.SeqWithoutIsStrInference([273, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([511]))])), len(_dafny.Map({127: not(False)}))]))})))

    @staticmethod
    def fm3(p0, globalState):
        return D0_DC2(not (False) or (not(True)))

    @staticmethod
    def fm4(globalState):
        return ((_dafny.SeqWithoutIsStrInference([not(not(False))]))) + (_dafny.SeqWithoutIsStrInference([True, True]))

    @staticmethod
    def fm5(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-864: 937}))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False}))]): 311}))]))

    @staticmethod
    def fm8(globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxltqisx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbiom")))

    @staticmethod
    def fm9(p0, p1, globalState):
        if False:
            return _dafny.CodePoint('h')
        elif True:
            return _dafny.CodePoint('v')

    @staticmethod
    def fm10(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([D0_DC0(566), D0_DC0(607)])) + (_dafny.SeqWithoutIsStrInference([D0_DC0(len(_dafny.Map({False: _dafny.CodePoint('g')}))), D0_DC0(len(_dafny.Set({True, False, False, False, not(False)})))]))

    @staticmethod
    def fm11(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwmmnrp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obby")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2_i0_ in range(default__.abs(247))])])).Elements:
                d_3_v0_: _dafny.Seq = compr_1_
                if (d_3_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwmmnrp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obby")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2_i0_ in range(default__.abs(247))])])):
                    coll1_[d_3_v0_] = (405) + (921)
            return _dafny.Map(coll1_)
        return iife1_()
        

    @staticmethod
    def fm12(p0, p1, globalState):
        return ((_dafny.Set({False})).intersection(_dafny.Set({True}))) | (_dafny.Set({not(True), False, False}))

    @staticmethod
    def fm13(p0, p1, globalState):
        return _dafny.Map({(-131) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqfasi")))): not (True) or (not(True))})

    @staticmethod
    def fm14(p0, p1, globalState):
        return _dafny.MultiSet([(True) == (False)])

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        if (531) not in (_dafny.Map({771: False})):
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in (_dafny.SeqWithoutIsStrInference([-874 for d_4_i0_ in range(default__.abs(942))])).Elements:
                    d_5_v0_: int = compr_2_
                    if (d_5_v0_) in (_dafny.SeqWithoutIsStrInference([-874 for d_4_i0_ in range(default__.abs(942))])):
                        def iife3_():
                            def iife5_():
                                coll5_ = _dafny.Map()
                                compr_5_: int
                                for compr_5_ in _dafny.IntegerRange(976, 745):
                                    d_6_v2_: int = compr_5_
                                    if ((976) <= (d_6_v2_)) and ((d_6_v2_) < (745)):
                                        coll5_[default__.safeModuloInt(d_6_v2_, 710)] = len(_dafny.SeqWithoutIsStrInference([not(False), False]))
                                return _dafny.Map(coll5_)
                            def iife6_():
                                coll6_ = _dafny.Set()
                                compr_6_: str
                                for compr_6_ in (_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('x')])).Elements:
                                    d_8_v3_: str = compr_6_
                                    if (d_8_v3_) in (_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('x')])):
                                        coll6_ = coll6_.union(_dafny.Set([d_8_v3_]))
                                return _dafny.Set(coll6_)
                            coll3_ = _dafny.Map()
                            def iife4_():
                                coll4_ = _dafny.Map()
                                compr_4_: int
                                for compr_4_ in _dafny.IntegerRange(976, 745):
                                    d_6_v2_: int = compr_4_
                                    if ((976) <= (d_6_v2_)) and ((d_6_v2_) < (745)):
                                        coll4_[default__.safeModuloInt(d_6_v2_, 710)] = len(_dafny.SeqWithoutIsStrInference([not(False), False]))
                                return _dafny.Map(coll4_)
                            compr_3_: int
                            for compr_3_ in ((D8_DC16(_dafny.Map({760: iife4_()
}))).cf32).keys.Elements:
                                d_7_v1_: int = compr_3_
                                if (d_7_v1_) in ((D8_DC16(_dafny.Map({760: iife5_()
}))).cf32):
                                    coll3_[(d_7_v1_) - (len(_dafny.Map({len(iife6_()
                                    ): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mamhl"))) for d_9_i1_ in range(default__.abs(228))]))})))] = _dafny.Map({(_dafny.MultiSet([False])).cardinality: True})
                            return _dafny.Map(coll3_)
                        coll2_[(d_5_v0_) * (871)] = len(iife3_()
                        )
                return _dafny.Map(coll2_)
            return iife2_()
            
        elif True:
            return (_dafny.Map({95: -232})) | (_dafny.Map({len(_dafny.Set({967})): -697}))

    @staticmethod
    def fm16(p0, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Seq
            for compr_7_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([not(False)])})).Elements:
                d_10_v0_: _dafny.Seq = compr_7_
                if (d_10_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([not(False)])})):
                    coll7_[d_10_v0_] = (len(_dafny.Set({True, True}))) == (-608)
            return _dafny.Map(coll7_)
        return iife7_()
        

    @staticmethod
    def fm17(p0, globalState):
        return _dafny.Map({(_dafny.MultiSet([True])).isdisjoint(_dafny.MultiSet([True, False])): _dafny.CodePoint('d')})

    @staticmethod
    def m0(p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        hi0_ = p0
        for d_11_i0_ in range(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))), hi0_):
            d_12_v0_: _dafny.Set
            d_12_v0_ = _dafny.Set({len(_dafny.Set({d_11_i0_})), p0, p0})
            d_13_v1_: bool
            d_13_v1_ = False
            d_14_v2_: C0
            nw0_ = C0()
            nw0_.ctor__()
            d_14_v2_ = nw0_
            d_15_v3_: _dafny.Map
            d_15_v3_ = _dafny.Map({d_14_v2_: 803})
            d_16_v4_: D3
            d_16_v4_ = D3_DC7(d_15_v3_)
            d_17_v5_: _dafny.Map
            d_17_v5_ = _dafny.Map({d_13_v1_: d_16_v4_})
            def iife8_():
                coll8_ = _dafny.Set()
                compr_8_: int
                for compr_8_ in ((d_12_v0_) | (d_12_v0_)).Elements:
                    d_18_v6_: int = compr_8_
                    if (d_18_v6_) in ((d_12_v0_) | (d_12_v0_)):
                        coll8_ = coll8_.union(_dafny.Set([(d_18_v6_) + ((201) + ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-128 for d_19_i1_ in range(default__.abs(-511))]))).cardinality))]))
                return _dafny.Set(coll8_)
            rhs0_ = iife8_()

            rhs1_ = d_17_v5_
            d_12_v0_ = rhs0_
            d_17_v5_ = rhs1_
            (globalState).f11 = (0) - (d_11_i0_)
            d_20_v7_: _dafny.Array
            def lambda0_(d_21_v1_):
                def lambda1_(d_22_i2_):
                    return d_21_v1_

                return lambda1_

            init0_ = lambda0_(d_13_v1_)
            nw1_ = _dafny.Array(None, 21)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_20_v7_ = nw1_
            d_23_v8_: str
            d_23_v8_ = _dafny.CodePoint('t')
            d_24_v9_: _dafny.Map
            d_24_v9_ = _dafny.Map({d_23_v8_: d_20_v7_})
            d_25_v10_: D0
            d_25_v10_ = D0_DC1(443, 315, p0, p0, ((d_24_v9_)[d_23_v8_] if (d_23_v8_) in (d_24_v9_) else d_20_v7_))
            d_26_v11_: _dafny.Map
            d_26_v11_ = _dafny.Map({d_25_v10_: d_20_v7_})
            d_27_v12_: _dafny.Array
            nw2_ = _dafny.Array(None, 17)
            nw2_[int(0)] = d_20_v7_
            nw2_[int(1)] = d_20_v7_
            nw2_[int(2)] = d_20_v7_
            nw2_[int(3)] = d_20_v7_
            nw2_[int(4)] = d_20_v7_
            nw2_[int(5)] = ((d_26_v11_)[d_25_v10_] if (d_25_v10_) in (d_26_v11_) else d_20_v7_)
            nw2_[int(6)] = d_20_v7_
            nw2_[int(7)] = d_20_v7_
            nw2_[int(8)] = d_20_v7_
            nw2_[int(9)] = d_20_v7_
            nw2_[int(10)] = d_20_v7_
            nw2_[int(11)] = d_20_v7_
            nw2_[int(12)] = d_20_v7_
            nw2_[int(13)] = d_20_v7_
            nw2_[int(14)] = d_20_v7_
            nw2_[int(15)] = (d_25_v10_).cf5
            nw2_[int(16)] = d_20_v7_
            d_27_v12_ = nw2_
            index0_ = default__.safeIndex(241, (d_27_v12_).length(0))
            (d_27_v12_)[index0_] = (d_25_v10_).cf5
            index1_ = default__.safeIndex(241, (d_27_v12_).length(0))
            (d_27_v12_)[index1_] = d_20_v7_
            d_28_v13_: _dafny.Array
            nw3_ = _dafny.Array(int(0), 29)
            d_28_v13_ = nw3_
            r0 = d_28_v13_
        d_29_v14_: _dafny.Map
        d_29_v14_ = _dafny.Map({p0: True})
        d_30_v15_: bool
        d_30_v15_ = False
        d_31_v16_: _dafny.Seq
        d_31_v16_ = _dafny.SeqWithoutIsStrInference([False, not(d_30_v15_)])
        d_32_v17_: _dafny.Set
        d_32_v17_ = _dafny.Set({p0})
        d_33_v18_: _dafny.Array
        nw4_ = _dafny.Array(None, 16)
        nw4_[int(0)] = d_30_v15_
        nw4_[int(1)] = False
        nw4_[int(2)] = d_30_v15_
        nw4_[int(3)] = False
        nw4_[int(4)] = d_30_v15_
        nw4_[int(5)] = d_30_v15_
        nw4_[int(6)] = d_30_v15_
        nw4_[int(7)] = ((d_29_v14_)[p0] if (p0) in (d_29_v14_) else d_30_v15_)
        nw4_[int(8)] = not(d_30_v15_)
        nw4_[int(9)] = d_30_v15_
        nw4_[int(10)] = d_30_v15_
        nw4_[int(11)] = not(d_30_v15_)
        nw4_[int(12)] = d_30_v15_
        nw4_[int(13)] = not(d_30_v15_)
        nw4_[int(14)] = not(d_30_v15_)
        nw4_[int(15)] = (d_31_v16_)[default__.safeIndex((0) - (len(d_32_v17_)), len(d_31_v16_))]
        d_33_v18_ = nw4_
        d_34_v19_: D0
        d_34_v19_ = D0_DC1(default__.fm0(globalState), len((d_29_v14_ if False else d_29_v14_)), p0, p0, d_33_v18_)
        source0_ = d_34_v19_
        if source0_.is_DC1:
            d_35___mcc_h0_ = source0_.cf1
            d_36___mcc_h1_ = source0_.cf2
            d_37___mcc_h2_ = source0_.cf3
            d_38___mcc_h3_ = source0_.cf4
            d_39___mcc_h4_ = source0_.cf5
            d_40_cf5_ = d_39___mcc_h4_
            d_41_cf4_ = d_38___mcc_h3_
            d_42_cf3_ = d_37___mcc_h2_
            d_43_cf2_ = d_36___mcc_h1_
            d_44_cf1_ = d_35___mcc_h0_
            d_45_v20_: _dafny.Set
            d_45_v20_ = _dafny.Set({True})
            d_46_v21_: C1
            nw5_ = C1()
            nw5_.ctor__((len(d_45_v20_)) - (d_43_cf2_))
            d_46_v21_ = nw5_
            d_47_v22_: _dafny.Array
            def lambda2_(d_48_v15_, d_49_v16_):
                def lambda3_(d_50_i3_):
                    return (d_50_i3_) * (len(_dafny.Map({d_48_v15_: len(d_49_v16_)})))

                return lambda3_

            init1_ = lambda2_(d_30_v15_, d_31_v16_)
            nw6_ = _dafny.Array(None, 26)
            for i0_1_ in range(nw6_.length(0)):
                nw6_[i0_1_] = init1_(i0_1_)
            d_47_v22_ = nw6_
            r0 = d_47_v22_
            d_41_cf4_ = d_46_v21_.f18
            (globalState).f12 = d_30_v15_
        elif source0_.is_DC2:
            d_51___mcc_h5_ = source0_.cf6
            d_52_cf6_ = d_51___mcc_h5_
            d_53_v23_: C1
            nw7_ = C1()
            nw7_.ctor__(-998)
            d_53_v23_ = nw7_
            d_53_v23_ = d_53_v23_
            d_33_v18_ = d_33_v18_
            d_52_cf6_ = (d_52_cf6_ if not (d_30_v15_) or (d_30_v15_) else (p0) <= (p0))
            d_54_v24_: C0
            nw8_ = C0()
            nw8_.ctor__()
            d_54_v24_ = nw8_
        elif source0_.is_DC3:
            d_55___mcc_h6_ = source0_.cf7
            d_56___mcc_h7_ = source0_.cf8
            d_57___mcc_h8_ = source0_.cf9
            d_58___mcc_h9_ = source0_.cf10
            d_59_cf10_ = d_58___mcc_h9_
            d_60_cf9_ = d_57___mcc_h8_
            d_61_cf8_ = d_56___mcc_h7_
            d_62_cf7_ = d_55___mcc_h6_
            d_63_v25_: _dafny.Seq
            d_63_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbjalwtvy"))
            d_64_v26_: D0
            d_64_v26_ = D0_DC0(len(d_63_v25_))
            d_64_v26_ = d_64_v26_
            (globalState).f12 = (d_63_v25_) <= (((default__.fm8(globalState)).set(default__.safeIndex(d_62_cf7_, len(default__.fm8(globalState))), _dafny.CodePoint('l'))) + (d_63_v25_))
            d_65_v27_: _dafny.Set
            d_65_v27_ = _dafny.Set({d_30_v15_, (d_31_v16_)[default__.safeIndex(d_62_cf7_, len(d_31_v16_))], d_61_cf8_, d_30_v15_, d_30_v15_})
            d_66_v28_: str
            d_66_v28_ = _dafny.CodePoint('s')
            d_65_v27_ = default__.fm12(d_66_v28_, p0, globalState)
            (globalState).f12 = d_61_cf8_
        elif True:
            d_67___mcc_h10_ = source0_.cf0
            d_68_cf0_ = d_67___mcc_h10_
            d_30_v15_ = d_30_v15_
            d_69_v29_: _dafny.Seq
            d_69_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dw"))
            d_70_v30_: _dafny.Seq
            d_70_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igmkken")), d_69_v29_])
            if ((d_31_v16_)[default__.safeIndex(105, len(d_31_v16_))]) == ((d_69_v29_) in (d_70_v30_)):
                d_71_v31_: _dafny.Seq
                d_71_v31_ = _dafny.SeqWithoutIsStrInference([p0, d_68_cf0_])
                d_72_v32_: _dafny.Map
                d_72_v32_ = _dafny.Map({d_30_v15_: d_71_v31_})
                d_73_v33_: D3
                d_73_v33_ = D3_DC8(d_71_v31_, True, d_30_v15_, d_69_v29_, d_30_v15_)
                d_72_v32_ = (d_72_v32_).set(default__.fm1(d_30_v15_, d_30_v15_, p0, globalState), (d_73_v33_).cf16)
                d_74_v34_: _dafny.MultiSet
                d_74_v34_ = _dafny.MultiSet([len((d_69_v29_) + (d_69_v29_)), (d_68_cf0_ if d_30_v15_ else p0), (p0) * (p0), p0, p0])
                d_75_v35_: _dafny.Map
                d_75_v35_ = _dafny.Map({default__.fm1(d_30_v15_, d_30_v15_, p0, globalState): (d_68_cf0_) < (default__.fm0(globalState))})
                d_76_v36_: _dafny.Set
                d_76_v36_ = _dafny.Set({d_30_v15_})
                d_77_v37_: _dafny.MultiSet
                d_77_v37_ = _dafny.MultiSet([d_31_v16_])
                d_78_v38_: _dafny.Map
                d_78_v38_ = _dafny.Map({d_30_v15_: d_77_v37_})
                d_79_v39_: _dafny.Map
                d_79_v39_ = _dafny.Map({(d_76_v36_) - (d_76_v36_): (((d_78_v38_)[d_30_v15_] if (d_30_v15_) in (d_78_v38_) else d_77_v37_)).isdisjoint(d_77_v37_)})
                d_80_v40_: str
                d_80_v40_ = _dafny.CodePoint('b')
                d_81_v41_: _dafny.Set
                d_81_v41_ = _dafny.Set({d_80_v40_})
                d_82_v42_: D0
                d_82_v42_ = D0_DC2(d_30_v15_)
                rhs2_ = ((d_79_v39_)[d_76_v36_] if (d_76_v36_) in (d_79_v39_) else d_30_v15_)
                rhs3_ = d_74_v34_
                rhs4_ = _dafny.Map({d_30_v15_: not(d_30_v15_)})
                rhs5_ = (d_81_v41_) == ((_dafny.Set({d_80_v40_, d_80_v40_, d_80_v40_, d_80_v40_, d_80_v40_})) - (d_81_v41_))
                rhs6_ = (d_68_cf0_ if ((d_82_v42_).cf6) and (d_30_v15_) else d_68_cf0_)
                lhs0_ = globalState
                lhs1_ = globalState
                lhs2_ = globalState
                lhs0_.f12 = rhs2_
                d_74_v34_ = rhs3_
                d_75_v35_ = rhs4_
                lhs1_.f12 = rhs5_
                lhs2_.f11 = rhs6_
                d_83_v43_: C0
                nw9_ = C0()
                nw9_.ctor__()
                d_83_v43_ = nw9_
                d_84_v44_: D4
                d_84_v44_ = D4_DC9(d_83_v43_)
                d_85_v45_: D4
                d_85_v45_ = D4_DC11(d_84_v44_)
                d_86_v46_: D4
                d_86_v46_ = D4_DC11(d_84_v44_)
                d_87_v47_: D4
                d_87_v47_ = D4_DC11(d_86_v46_)
                rhs7_ = len(d_69_v29_)
                rhs8_ = d_87_v47_
                rhs9_ = d_80_v40_
                lhs3_ = globalState
                lhs3_.f11 = rhs7_
                d_87_v47_ = rhs8_
                d_80_v40_ = rhs9_
                d_88_v48_: _dafny.Array
                nw10_ = _dafny.Array(_dafny.Seq({}), 29)
                d_88_v48_ = nw10_
                d_89_v49_: _dafny.Seq
                d_89_v49_ = _dafny.SeqWithoutIsStrInference([d_81_v41_])
                index2_ = default__.safeIndex(291, (d_88_v48_).length(0))
                (d_88_v48_)[index2_] = d_89_v49_
                d_90_v50_: _dafny.MultiSet
                d_90_v50_ = _dafny.MultiSet([d_30_v15_])
                index3_ = default__.safeIndex(291, (d_88_v48_).length(0))
                rhs10_ = d_33_v18_
                rhs11_ = d_89_v49_
                rhs12_ = (d_90_v50_).intersection(d_90_v50_)
                lhs4_ = d_88_v48_
                lhs5_ = default__.safeIndex(291, (d_88_v48_).length(0))
                d_33_v18_ = rhs10_
                lhs4_[lhs5_] = rhs11_
                d_90_v50_ = rhs12_
                d_91_v51_: _dafny.Array
                def lambda4_(d_92_p0_):
                    def lambda5_(d_93_i4_):
                        return (d_93_i4_) + (d_92_p0_)

                    return lambda5_

                init2_ = lambda4_(p0)
                nw11_ = _dafny.Array(None, 21)
                for i0_2_ in range(nw11_.length(0)):
                    nw11_[i0_2_] = init2_(i0_2_)
                d_91_v51_ = nw11_
                index4_ = default__.safeIndex(4, (d_91_v51_).length(0))
                (d_91_v51_)[index4_] = default__.safeModuloInt(d_68_cf0_, d_68_cf0_)
                index5_ = default__.safeIndex(4, (d_91_v51_).length(0))
                rhs13_ = p0
                rhs14_ = len(d_31_v16_)
                rhs15_ = d_68_cf0_
                rhs16_ = ((-722 if d_30_v15_ else d_68_cf0_)) + (default__.fm0(globalState))
                rhs17_ = default__.safeDivisionInt((d_71_v31_)[default__.safeIndex(p0, len(d_71_v31_))], (d_71_v31_)[default__.safeIndex(len(d_69_v29_), len(d_71_v31_))])
                lhs6_ = globalState
                lhs7_ = globalState
                lhs8_ = d_91_v51_
                lhs9_ = default__.safeIndex(4, (d_91_v51_).length(0))
                lhs10_ = globalState
                lhs6_.f11 = rhs13_
                d_68_cf0_ = rhs14_
                lhs7_.f11 = rhs15_
                lhs8_[lhs9_] = rhs16_
                lhs10_.f10 = rhs17_
            elif True:
                (globalState).f12 = (D3_DC8(_dafny.SeqWithoutIsStrInference([p0]), default__.fm1(d_30_v15_, not(d_30_v15_), d_68_cf0_, globalState), d_30_v15_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_94_i5_ in range(default__.abs(272))])).set(default__.safeIndex(d_68_cf0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_95_i5_ in range(default__.abs(272))]))), (d_69_v29_)[default__.safeIndex(p0, len(d_69_v29_))]), True)).cf20
                (globalState).f12 = d_30_v15_
                d_96_v52_: _dafny.Map
                d_96_v52_ = _dafny.Map({d_30_v15_: (d_68_cf0_) == (491)})
                d_96_v52_ = (d_96_v52_).set((d_69_v29_) != (d_69_v29_), d_30_v15_)
                d_33_v18_ = d_33_v18_
                d_97_v53_: _dafny.Array
                def lambda6_(d_98_cf0_, d_99_p0_):
                    def lambda7_(d_100_i6_):
                        return (d_100_i6_) + (len(_dafny.Map({d_98_cf0_: d_99_p0_})))

                    return lambda7_

                init3_ = lambda6_(d_68_cf0_, p0)
                nw12_ = _dafny.Array(None, 11)
                for i0_3_ in range(nw12_.length(0)):
                    nw12_[i0_3_] = init3_(i0_3_)
                d_97_v53_ = nw12_
                index6_ = default__.safeIndex(43, (d_97_v53_).length(0))
                (d_97_v53_)[index6_] = p0
                index7_ = default__.safeIndex(43, (d_97_v53_).length(0))
                (d_97_v53_)[index7_] = p0
            d_68_cf0_ = (p0) - (d_68_cf0_)
            d_101_v54_: _dafny.MultiSet
            d_101_v54_ = _dafny.MultiSet([d_30_v15_, False])
            (globalState).f12 = ((d_101_v54_).cardinality) == (d_68_cf0_)
        if (True) and (d_30_v15_):
            (globalState).f11 = p0
            (globalState).f12 = not(d_30_v15_)
            d_102_v55_: str
            d_102_v55_ = _dafny.CodePoint('v')
            d_103_v56_: _dafny.Map
            d_103_v56_ = _dafny.Map({d_30_v15_: d_102_v55_})
            d_104_v57_: _dafny.Map
            d_104_v57_ = _dafny.Map({d_103_v56_: d_30_v15_})
            index8_ = default__.safeIndex(708, (d_33_v18_).length(0))
            (d_33_v18_)[index8_] = ((d_104_v57_)[default__.fm17(d_30_v15_, globalState)] if (default__.fm17(d_30_v15_, globalState)) in (d_104_v57_) else d_30_v15_)
            d_105_v58_: _dafny.MultiSet
            d_105_v58_ = _dafny.MultiSet([p0])
            d_106_v59_: _dafny.Seq
            d_106_v59_ = _dafny.SeqWithoutIsStrInference([p0])
            d_107_v60_: _dafny.Map
            d_107_v60_ = _dafny.Map({d_30_v15_: d_30_v15_})
            index9_ = default__.safeIndex(708, (d_33_v18_).length(0))
            (d_33_v18_)[index9_] = (((d_105_v58_)[p0] if (p0) in (d_105_v58_) else p0)) <= ((d_106_v59_)[default__.safeIndex(len(d_107_v60_), len(d_106_v59_))])
            d_108_v61_: _dafny.MultiSet
            d_108_v61_ = _dafny.MultiSet([(d_33_v18_)[default__.safeIndex(708, (d_33_v18_).length(0))]])
            d_109_v62_: _dafny.Seq
            d_109_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjuprfrne"))
            d_110_v63_: _dafny.Set
            d_110_v63_ = _dafny.Set({(d_108_v61_).set(True, default__.abs(len(d_109_v62_))), d_108_v61_, d_108_v61_})
            d_111_v64_: _dafny.Map
            d_111_v64_ = _dafny.Map({p0: len(d_110_v63_)})
            d_112_v65_: C1
            nw13_ = C1()
            nw13_.ctor__(len(d_111_v64_))
            d_112_v65_ = nw13_
            d_113_v66_: _dafny.Array
            def lambda8_(d_114_v55_):
                def lambda9_(d_115_i7_):
                    return d_114_v55_

                return lambda9_

            init4_ = lambda8_(d_102_v55_)
            nw14_ = _dafny.Array(None, 3)
            for i0_4_ in range(nw14_.length(0)):
                nw14_[i0_4_] = init4_(i0_4_)
            d_113_v66_ = nw14_
            d_116_v67_: D6
            d_116_v67_ = D6_DC14(d_30_v15_, d_102_v55_, d_113_v66_)
            pat_let_tv0_ = d_102_v55_
            pat_let_tv1_ = d_30_v15_
            pat_let_tv2_ = d_30_v15_
            def iife9_(_pat_let0_0):
                def iife10_(d_117_dt__update__tmp_h0_):
                    def iife11_(_pat_let1_0):
                        def iife12_(d_118_dt__update_hcf29_h0_):
                            def iife13_(_pat_let2_0):
                                def iife14_(d_119_dt__update_hcf28_h0_):
                                    return D6_DC14(d_119_dt__update_hcf28_h0_, d_118_dt__update_hcf29_h0_, (d_117_dt__update__tmp_h0_).cf30)
                                return iife14_(_pat_let2_0)
                            return iife13_((pat_let_tv1_) and (pat_let_tv2_))
                        return iife12_(_pat_let1_0)
                    return iife11_(pat_let_tv0_)
                return iife10_(_pat_let0_0)
            source1_ = iife9_(d_116_v67_)
            if source1_.is_DC14:
                d_120___mcc_h11_ = source1_.cf28
                d_121___mcc_h12_ = source1_.cf29
                d_122___mcc_h13_ = source1_.cf30
                d_123_cf30_ = d_122___mcc_h13_
                d_124_cf29_ = d_121___mcc_h12_
                d_125_cf28_ = d_120___mcc_h11_
                (globalState).f10 = d_112_v65_.f18
                d_107_v60_ = (d_107_v60_).set(not(d_30_v15_), (d_33_v18_)[default__.safeIndex(708, (d_33_v18_).length(0))])
                d_126_v68_: _dafny.Array
                nw15_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                d_126_v68_ = nw15_
                index10_ = default__.safeIndex(394, (d_126_v68_).length(0))
                (d_126_v68_)[index10_] = d_109_v62_
                index11_ = default__.safeIndex(394, (d_126_v68_).length(0))
                rhs18_ = d_102_v55_
                rhs19_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgxwsdkj"))).set(default__.safeIndex(d_112_v65_.f18, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgxwsdkj")))), d_124_cf29_)
                lhs11_ = d_126_v68_
                lhs12_ = default__.safeIndex(394, (d_126_v68_).length(0))
                d_124_cf29_ = rhs18_
                lhs11_[lhs12_] = rhs19_
                d_127_v69_: _dafny.Array
                nw16_ = _dafny.Array(int(0), 28)
                d_127_v69_ = nw16_
                index12_ = default__.safeIndex(195, (d_127_v69_).length(0))
                (d_127_v69_)[index12_] = d_112_v65_.f18
                index13_ = default__.safeIndex(195, (d_127_v69_).length(0))
                (d_127_v69_)[index13_] = (0) - (p0)
            elif True:
                d_128___mcc_h14_ = source1_.cf27
                d_129_cf27_ = d_128___mcc_h14_
                (globalState).f15 = d_102_v55_
                d_130_v70_: _dafny.Array
                nw17_ = _dafny.Array(int(0), 29)
                d_130_v70_ = nw17_
                d_131_v71_: _dafny.Array
                nw18_ = _dafny.Array(None, 28)
                nw18_[int(0)] = d_130_v70_
                nw18_[int(1)] = d_130_v70_
                nw18_[int(2)] = d_130_v70_
                nw18_[int(3)] = d_130_v70_
                nw18_[int(4)] = d_130_v70_
                nw18_[int(5)] = d_130_v70_
                nw18_[int(6)] = d_130_v70_
                nw18_[int(7)] = d_130_v70_
                nw18_[int(8)] = (d_130_v70_ if d_30_v15_ else d_130_v70_)
                nw18_[int(9)] = d_130_v70_
                nw18_[int(10)] = d_130_v70_
                nw18_[int(11)] = d_130_v70_
                nw18_[int(12)] = d_130_v70_
                nw18_[int(13)] = d_130_v70_
                nw18_[int(14)] = d_130_v70_
                nw18_[int(15)] = d_130_v70_
                nw18_[int(16)] = d_130_v70_
                nw18_[int(17)] = d_130_v70_
                nw18_[int(18)] = d_130_v70_
                nw18_[int(19)] = d_130_v70_
                nw18_[int(20)] = d_130_v70_
                nw18_[int(21)] = d_130_v70_
                nw18_[int(22)] = d_130_v70_
                nw18_[int(23)] = d_130_v70_
                nw18_[int(24)] = d_130_v70_
                nw18_[int(25)] = d_130_v70_
                nw18_[int(26)] = d_130_v70_
                nw18_[int(27)] = d_130_v70_
                d_131_v71_ = nw18_
                index14_ = default__.safeIndex(425, (d_131_v71_).length(0))
                (d_131_v71_)[index14_] = d_130_v70_
                d_132_v72_: C0
                nw19_ = C0()
                nw19_.ctor__()
                d_132_v72_ = nw19_
                d_133_v73_: _dafny.Array
                nw20_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_133_v73_ = nw20_
                d_134_v74_: _dafny.Array
                def lambda10_(d_135_v14_):
                    def lambda11_(d_136_i8_):
                        return d_135_v14_

                    return lambda11_

                init5_ = lambda10_(d_29_v14_)
                nw21_ = _dafny.Array(None, 5)
                for i0_5_ in range(nw21_.length(0)):
                    nw21_[i0_5_] = init5_(i0_5_)
                d_134_v74_ = nw21_
                index15_ = default__.safeIndex(371, (d_133_v73_).length(0))
                (d_133_v73_)[index15_] = d_134_v74_
                d_137_v75_: _dafny.Set
                d_137_v75_ = _dafny.Set({d_130_v70_, d_130_v70_})
                index16_ = default__.safeIndex(425, (d_131_v71_).length(0))
                index17_ = default__.safeIndex(371, (d_133_v73_).length(0))
                rhs20_ = d_130_v70_
                rhs21_ = d_132_v72_
                rhs22_ = (d_134_v74_ if (d_137_v75_).ispropersubset(d_137_v75_) else d_134_v74_)
                lhs13_ = d_131_v71_
                lhs14_ = default__.safeIndex(425, (d_131_v71_).length(0))
                lhs15_ = d_133_v73_
                lhs16_ = default__.safeIndex(371, (d_133_v73_).length(0))
                lhs13_[lhs14_] = rhs20_
                d_132_v72_ = rhs21_
                lhs15_[lhs16_] = rhs22_
                d_138_v76_: _dafny.Array
                nw22_ = _dafny.Array(_dafny.Map({}), 5)
                d_138_v76_ = nw22_
                d_139_v77_: _dafny.Map
                d_139_v77_ = _dafny.Map({d_109_v62_: d_112_v65_.f18})
                index18_ = default__.safeIndex(313, (d_138_v76_).length(0))
                (d_138_v76_)[index18_] = d_139_v77_
                index19_ = default__.safeIndex(313, (d_138_v76_).length(0))
                (d_138_v76_)[index19_] = d_139_v77_
                d_140_v78_: _dafny.MultiSet
                d_140_v78_ = _dafny.MultiSet([d_109_v62_])
                d_141_v79_: _dafny.Map
                d_141_v79_ = _dafny.Map({(d_33_v18_)[default__.safeIndex(708, (d_33_v18_).length(0))]: p0})
                index20_ = default__.safeIndex(205, (d_130_v70_).length(0))
                (d_130_v70_)[index20_] = ((d_140_v78_)[(d_132_v72_).fm7(d_30_v15_, len(_dafny.SeqWithoutIsStrInference([d_102_v55_, d_102_v55_])), d_141_v79_, d_109_v62_, globalState)] if ((d_132_v72_).fm7(d_30_v15_, len(_dafny.SeqWithoutIsStrInference([d_102_v55_, d_102_v55_])), d_141_v79_, d_109_v62_, globalState)) in (d_140_v78_) else (0) - ((d_108_v61_).cardinality))
                d_142_v81_: _dafny.Map
                def iife15_():
                    coll9_ = _dafny.Set()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(-954, 990):
                        d_143_v80_: int = compr_9_
                        if ((-954) <= (d_143_v80_)) and ((d_143_v80_) < (990)):
                            coll9_ = coll9_.union(_dafny.Set([(d_143_v80_) * (d_112_v65_.f18)]))
                    return _dafny.Set(coll9_)
                d_142_v81_ = _dafny.Map({iife15_()
                : p0})
                d_144_v83_: D6
                d_144_v83_ = D6_DC13(d_129_cf27_)
                d_145_v84_: _dafny.MultiSet
                d_145_v84_ = _dafny.MultiSet([d_144_v83_])
                d_146_v85_: _dafny.Map
                d_146_v85_ = _dafny.Map({d_145_v84_: (d_33_v18_)[default__.safeIndex(708, (d_33_v18_).length(0))]})
                index21_ = default__.safeIndex(708, (d_33_v18_).length(0))
                index22_ = default__.safeIndex(205, (d_130_v70_).length(0))
                def iife16_():
                    coll10_ = _dafny.Set()
                    compr_10_: int
                    for compr_10_ in (d_29_v14_).keys.Elements:
                        d_147_v82_: int = compr_10_
                        if (d_147_v82_) in (d_29_v14_):
                            coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_147_v82_, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): 420})))]))
                    return _dafny.Set(coll10_)
                def iife17_():
                    coll11_ = _dafny.Set()
                    compr_11_: int
                    for compr_11_ in (d_29_v14_).keys.Elements:
                        d_148_v82_: int = compr_11_
                        if (d_148_v82_) in (d_29_v14_):
                            coll11_ = coll11_.union(_dafny.Set([default__.safeModuloInt(d_148_v82_, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): 420})))]))
                    return _dafny.Set(coll11_)
                rhs23_ = d_106_v59_
                rhs24_ = not(d_30_v15_)
                rhs25_ = ((d_142_v81_)[iife16_()
                ] if (iife17_()
                ) in (d_142_v81_) else p0)
                rhs26_ = (d_112_v65_.f18) - (default__.safeModuloInt(len(d_146_v85_), len(_dafny.Set({d_30_v15_}))))
                rhs27_ = (d_109_v62_) + (d_109_v62_)
                lhs17_ = d_33_v18_
                lhs18_ = default__.safeIndex(708, (d_33_v18_).length(0))
                lhs19_ = globalState
                lhs20_ = d_130_v70_
                lhs21_ = default__.safeIndex(205, (d_130_v70_).length(0))
                d_106_v59_ = rhs23_
                lhs17_[lhs18_] = rhs24_
                lhs19_.f10 = rhs25_
                lhs20_[lhs21_] = rhs26_
                d_109_v62_ = rhs27_
        elif True:
            d_149_v86_: _dafny.MultiSet
            d_149_v86_ = _dafny.MultiSet([not(d_30_v15_)])
            d_150_v87_: _dafny.Map
            d_150_v87_ = _dafny.Map({d_33_v18_: d_33_v18_})
            rhs28_ = d_30_v15_
            rhs29_ = d_149_v86_
            rhs30_ = len((d_32_v17_) | (d_32_v17_))
            rhs31_ = ((d_150_v87_) | (d_150_v87_) if (p0) > (p0) else d_150_v87_)
            lhs22_ = globalState
            lhs23_ = globalState
            lhs22_.f12 = rhs28_
            d_149_v86_ = rhs29_
            lhs23_.f11 = rhs30_
            d_150_v87_ = rhs31_
            d_151_v88_: _dafny.Set
            d_151_v88_ = _dafny.Set({True})
            (globalState).f10 = (p0) + (len(d_151_v88_))
            d_152_v89_: _dafny.Map
            d_152_v89_ = _dafny.Map({d_30_v15_: d_30_v15_})
            d_153_v90_: _dafny.MultiSet
            d_153_v90_ = _dafny.MultiSet([d_152_v89_, _dafny.Map({default__.fm1(d_30_v15_, d_30_v15_, p0, globalState): d_30_v15_})])
            d_153_v90_ = ((d_153_v90_).set(d_152_v89_, default__.abs(p0)) if (d_31_v16_)[default__.safeIndex(p0, len(d_31_v16_))] else (d_153_v90_) | (d_153_v90_))
            d_154_v91_: _dafny.Array
            nw23_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_154_v91_ = nw23_
            d_155_v92_: _dafny.Array
            nw24_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_155_v92_ = nw24_
            index23_ = default__.safeIndex(130, (d_154_v91_).length(0))
            (d_154_v91_)[index23_] = d_155_v92_
            d_156_v93_: _dafny.Seq
            d_156_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpyho"))
            d_157_v94_: _dafny.Seq
            d_157_v94_ = d_156_v93_
            d_158_v95_: str
            d_158_v95_ = _dafny.CodePoint('c')
            index24_ = default__.safeIndex(130, (d_154_v91_).length(0))
            nw25_ = _dafny.Array(None, 5)
            nw25_[int(0)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lksj"))) + (d_156_v93_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lksj"))) + (d_156_v93_))), (d_156_v93_)[default__.safeIndex(p0, len(d_156_v93_))])
            nw25_[int(1)] = (d_157_v94_)
            nw25_[int(2)] = (d_156_v93_) + (d_156_v93_)
            nw25_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_159_i9_ in range(default__.abs(116))])
            nw25_[int(4)] = ((d_156_v93_) + (d_156_v93_)).set(default__.safeIndex(default__.fm0(globalState), len((d_156_v93_) + (d_156_v93_))), d_158_v95_)
            (d_154_v91_)[index24_] = nw25_
            rhs32_ = d_30_v15_
            rhs33_ = d_158_v95_
            rhs34_ = p0
            rhs35_ = (default__.fm8(globalState)) + ((d_156_v93_).set(default__.safeIndex(p0, len(d_156_v93_)), d_158_v95_))
            lhs24_ = globalState
            lhs25_ = globalState
            lhs26_ = globalState
            lhs24_.f12 = rhs32_
            d_158_v95_ = rhs33_
            lhs25_.f10 = rhs34_
            lhs26_.f3 = rhs35_
        d_160_i10_: int
        d_160_i10_ = 0
        with _dafny.label("0"):
            while not(d_30_v15_):
                with _dafny.c_label("0"):
                    if (d_160_i10_) >= (100):
                        raise _dafny.Break("0")
                    d_160_i10_ = (d_160_i10_) + (1)
                    d_161_v96_: C0
                    nw26_ = C0()
                    nw26_.ctor__()
                    d_161_v96_ = nw26_
                    d_162_v97_: _dafny.Map
                    d_162_v97_ = _dafny.Map({d_161_v96_: d_30_v15_})
                    d_163_v98_: _dafny.Map
                    d_163_v98_ = _dafny.Map({d_30_v15_: d_162_v97_})
                    d_163_v98_ = (d_163_v98_).set(d_30_v15_, (d_162_v97_) | (d_162_v97_))
                    (globalState).f11 = (len(default__.fm15(p0, d_30_v15_, 794, globalState))) - (p0)
                    d_164_v99_: _dafny.Seq
                    d_164_v99_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                    d_165_v100_: C1
                    nw27_ = C1()
                    nw27_.ctor__(451)
                    d_165_v100_ = nw27_
                    d_166_v102_: str
                    d_166_v102_ = _dafny.CodePoint('u')
                    d_167_v103_: _dafny.Seq
                    d_167_v103_ = _dafny.SeqWithoutIsStrInference([d_166_v102_, d_166_v102_])
                    d_168_v104_: _dafny.Map
                    def iife18_():
                        coll12_ = _dafny.Map()
                        compr_12_: str
                        for compr_12_ in (d_167_v103_).Elements:
                            d_169_v101_: str = compr_12_
                            if (d_169_v101_) in (d_167_v103_):
                                coll12_[d_169_v101_] = d_165_v100_.f18
                        return _dafny.Map(coll12_)
                    d_168_v104_ = _dafny.Map({d_165_v100_: iife18_()
                    })
                    d_170_v105_: _dafny.Array
                    nw28_ = _dafny.Array(None, 23)
                    nw28_[int(0)] = (p0 if d_30_v15_ else 899)
                    nw28_[int(1)] = (p0) * (p0)
                    nw28_[int(2)] = 586
                    nw28_[int(3)] = 120
                    nw28_[int(4)] = p0
                    nw28_[int(5)] = p0
                    nw28_[int(6)] = p0
                    nw28_[int(7)] = p0
                    nw28_[int(8)] = p0
                    nw28_[int(9)] = p0
                    nw28_[int(10)] = default__.fm0(globalState)
                    nw28_[int(11)] = len((d_164_v99_).set(default__.safeIndex(p0, len(d_164_v99_)), (0) - (p0)))
                    nw28_[int(12)] = p0
                    nw28_[int(13)] = p0
                    nw28_[int(14)] = (len((d_164_v99_).set(default__.safeIndex(p0, len(d_164_v99_)), 472))) * (len(d_168_v104_))
                    nw28_[int(15)] = d_165_v100_.f18
                    nw28_[int(16)] = d_165_v100_.f18
                    nw28_[int(17)] = p0
                    nw28_[int(18)] = len((d_167_v103_).set(default__.safeIndex(p0, len(d_167_v103_)), (d_167_v103_)[default__.safeIndex(p0, len(d_167_v103_))]))
                    nw28_[int(19)] = d_165_v100_.f18
                    nw28_[int(20)] = (default__.fm0(globalState) if default__.fm1(d_30_v15_, d_30_v15_, len(d_167_v103_), globalState) else (d_164_v99_)[default__.safeIndex(d_165_v100_.f18, len(d_164_v99_))])
                    nw28_[int(21)] = d_165_v100_.f18
                    nw28_[int(22)] = p0
                    d_170_v105_ = nw28_
                    r0 = d_170_v105_
                    (globalState).f12 = not(d_30_v15_)
                    pass
            pass
        d_33_v18_ = d_33_v18_
        d_171_v106_: _dafny.Array
        nw29_ = _dafny.Array(_dafny.CodePoint('D'), 20)
        d_171_v106_ = nw29_
        d_171_v106_ = d_171_v106_
        d_172_v107_: _dafny.Array
        nw30_ = _dafny.Array(None, 1)
        nw30_[int(0)] = p0
        d_172_v107_ = nw30_
        r0 = d_172_v107_
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_173_v0_: _dafny.Seq
        d_173_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yo"))
        d_174_v2_: str
        d_174_v2_ = _dafny.CodePoint('k')
        d_175_v3_: _dafny.Set
        d_175_v3_ = _dafny.Set({d_174_v2_})
        d_176_v4_: _dafny.MultiSet
        def iife19_():
            coll13_ = _dafny.Set()
            compr_13_: str
            for compr_13_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n')])).Elements:
                d_177_v1_: str = compr_13_
                if (d_177_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n')])):
                    coll13_ = coll13_.union(_dafny.Set([d_177_v1_]))
            return _dafny.Set(coll13_)
        d_176_v4_ = _dafny.MultiSet([iife19_()
        , d_175_v3_])
        d_178_globalState_: GlobalState
        nw31_ = GlobalState()
        nw31_.ctor__(134, True, 107, d_173_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ci")), 380, -147, False, False, 727, 540, 576, False, False, d_176_v4_, _dafny.CodePoint('c'), 403, -68)
        d_178_globalState_ = nw31_
        d_179_v5_: int
        d_179_v5_ = 980
        hi1_ = d_179_v5_
        for d_180_i0_ in range(-735, hi1_):
            d_181_v6_: _dafny.Array
            nw32_ = _dafny.Array(False, 7)
            d_181_v6_ = nw32_
            d_182_v7_: _dafny.Map
            d_182_v7_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_181_v6_, d_181_v6_, d_181_v6_]): False})
            d_183_v8_: _dafny.Seq
            d_183_v8_ = _dafny.SeqWithoutIsStrInference([d_181_v6_, d_181_v6_, d_181_v6_, d_181_v6_, d_181_v6_])
            d_182_v7_ = (d_182_v7_).set((d_183_v8_) + (d_183_v8_), False)
            d_184_v9_: _dafny.Array
            nw33_ = _dafny.Array(None, 2)
            nw33_[int(0)] = -13
            nw33_[int(1)] = (D0_DC0(d_180_i0_)).cf0
            d_184_v9_ = nw33_
            d_185_v10_: bool
            d_185_v10_ = True
            d_186_v11_: _dafny.Seq
            d_186_v11_ = _dafny.SeqWithoutIsStrInference([not(d_185_v10_), d_185_v10_])
            d_187_v12_: _dafny.Map
            d_187_v12_ = _dafny.Map({(d_179_v5_) * (d_180_i0_): d_185_v10_})
            rhs36_ = (d_184_v9_ if (-438) > (880) else d_184_v9_)
            rhs37_ = not((d_173_v0_) != (d_173_v0_))
            rhs38_ = (d_185_v10_ if d_185_v10_ else (d_186_v11_)[default__.safeIndex(d_179_v5_, len(d_186_v11_))])
            rhs39_ = ((d_187_v12_)[716] if (716) in (d_187_v12_) else d_185_v10_)
            rhs40_ = ((0) - (d_180_i0_)) - ((len(d_173_v0_)) * ((0) - ((0) - (default__.fm0(d_178_globalState_)))))
            lhs27_ = d_178_globalState_
            lhs28_ = d_178_globalState_
            lhs29_ = d_178_globalState_
            d_184_v9_ = rhs36_
            lhs27_.f12 = rhs37_
            lhs28_.f12 = rhs38_
            lhs29_.f12 = rhs39_
            d_179_v5_ = rhs40_
            d_188_v13_: _dafny.MultiSet
            d_188_v13_ = _dafny.MultiSet([default__.fm0(d_178_globalState_)])
            d_188_v13_ = (d_188_v13_) | (d_188_v13_)
            d_189_v14_: _dafny.MultiSet
            d_189_v14_ = _dafny.MultiSet([d_185_v10_, True, True])
            d_190_v15_: _dafny.Map
            d_190_v15_ = _dafny.Map({d_185_v10_: (d_189_v14_).cardinality})
            d_191_v16_: _dafny.Map
            d_191_v16_ = _dafny.Map({d_179_v5_: d_190_v15_})
            d_191_v16_ = (d_191_v16_).set(d_180_i0_, d_190_v15_)
        d_192_v17_: bool
        d_192_v17_ = False
        d_193_v18_: _dafny.Map
        d_193_v18_ = _dafny.Map({d_192_v17_: d_179_v5_})
        d_194_v19_: _dafny.Map
        d_194_v19_ = _dafny.Map({len(_dafny.Set({d_179_v5_, d_179_v5_, len(d_193_v18_)})): not(d_192_v17_)})
        if ((d_194_v19_)[d_179_v5_] if (d_179_v5_) in (d_194_v19_) else (d_179_v5_) == (d_179_v5_)):
            d_195_v20_: _dafny.Array
            nw34_ = _dafny.Array(False, 22)
            d_195_v20_ = nw34_
            d_196_v21_: D0
            d_196_v21_ = D0_DC3(d_179_v5_, d_192_v17_, d_179_v5_, len(d_173_v0_))
            index25_ = default__.safeIndex(972, (d_195_v20_).length(0))
            (d_195_v20_)[index25_] = (default__.fm1(d_192_v17_, d_192_v17_, d_179_v5_, d_178_globalState_) if ((d_194_v19_)[d_179_v5_] if (d_179_v5_) in (d_194_v19_) else (d_196_v21_).cf8) else d_192_v17_)
            d_197_v22_: _dafny.Set
            d_197_v22_ = _dafny.Set({d_192_v17_, not(d_192_v17_)})
            index26_ = default__.safeIndex(972, (d_195_v20_).length(0))
            (d_195_v20_)[index26_] = default__.fm1(False, (d_197_v22_).isdisjoint(d_197_v22_), default__.safeModuloInt(d_179_v5_, d_179_v5_), d_178_globalState_)
            d_198_v23_: _dafny.Array
            nw35_ = _dafny.Array(_dafny.Map({}), 21)
            d_198_v23_ = nw35_
            d_199_v25_: _dafny.Map
            d_199_v25_ = _dafny.Map({d_179_v5_: d_179_v5_})
            d_200_v26_: _dafny.Map
            def iife20_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in (d_199_v25_).keys.Elements:
                    d_201_v24_: int = compr_14_
                    if (d_201_v24_) in (d_199_v25_):
                        coll14_[(d_201_v24_) * (d_179_v5_)] = D0_DC2((d_195_v20_)[default__.safeIndex(972, (d_195_v20_).length(0))])
                return _dafny.Map(coll14_)
            d_200_v26_ = _dafny.Map({iife20_()
            : d_179_v5_})
            index27_ = default__.safeIndex(397, (d_198_v23_).length(0))
            (d_198_v23_)[index27_] = d_200_v26_
            index28_ = default__.safeIndex(397, (d_198_v23_).length(0))
            (d_198_v23_)[index28_] = default__.fm2(d_178_globalState_)
            source2_ = default__.fm3(len(d_173_v0_), d_178_globalState_)
            if source2_.is_DC1:
                d_202___mcc_h0_ = source2_.cf1
                d_203___mcc_h1_ = source2_.cf2
                d_204___mcc_h2_ = source2_.cf3
                d_205___mcc_h3_ = source2_.cf4
                d_206___mcc_h4_ = source2_.cf5
                d_207_cf5_ = d_206___mcc_h4_
                d_208_cf4_ = d_205___mcc_h3_
                d_209_cf3_ = d_204___mcc_h2_
                d_210_cf2_ = d_203___mcc_h1_
                d_211_cf1_ = d_202___mcc_h0_
                d_212_v27_: _dafny.Array
                out0_: _dafny.Array
                out0_ = default__.m0((d_210_cf2_ if (d_195_v20_)[default__.safeIndex(972, (d_195_v20_).length(0))] else d_208_cf4_), d_178_globalState_)
                d_212_v27_ = out0_
                (d_178_globalState_).f10 = (0) - ((0) - (d_208_cf4_))
                d_213_v28_: _dafny.Array
                nw36_ = _dafny.Array(_dafny.Set({}), 10)
                d_213_v28_ = nw36_
                d_214_v29_: _dafny.Seq
                d_214_v29_ = _dafny.SeqWithoutIsStrInference([d_210_cf2_])
                d_215_v30_: _dafny.Set
                d_215_v30_ = _dafny.Set({d_214_v29_})
                index29_ = default__.safeIndex(2, (d_213_v28_).length(0))
                (d_213_v28_)[index29_] = d_215_v30_
                index30_ = default__.safeIndex(2, (d_213_v28_).length(0))
                (d_213_v28_)[index30_] = d_215_v30_
                index31_ = default__.safeIndex(972, (d_195_v20_).length(0))
                (d_195_v20_)[index31_] = not((d_174_v2_) in (d_173_v0_))
            elif source2_.is_DC2:
                d_216___mcc_h5_ = source2_.cf6
                d_217_cf6_ = d_216___mcc_h5_
                d_218_v31_: _dafny.Seq
                d_218_v31_ = _dafny.SeqWithoutIsStrInference([(d_195_v20_)[default__.safeIndex(972, (d_195_v20_).length(0))]])
                d_218_v31_ = (d_218_v31_) + (d_218_v31_)
                rhs41_ = default__.safeModuloInt(d_179_v5_, d_179_v5_)
                rhs42_ = ((0) - (d_179_v5_)) - (d_179_v5_)
                lhs30_ = d_178_globalState_
                lhs31_ = d_178_globalState_
                lhs30_.f10 = rhs41_
                lhs31_.f10 = rhs42_
                (d_178_globalState_).f12 = False
                d_219_v32_: _dafny.Array
                def lambda12_(d_220_v5_):
                    def lambda13_(d_221_i1_):
                        return (d_221_i1_) + (d_220_v5_)

                    return lambda13_

                init6_ = lambda12_(d_179_v5_)
                nw37_ = _dafny.Array(None, 4)
                for i0_6_ in range(nw37_.length(0)):
                    nw37_[i0_6_] = init6_(i0_6_)
                d_219_v32_ = nw37_
                index32_ = default__.safeIndex(302, (d_219_v32_).length(0))
                (d_219_v32_)[index32_] = d_179_v5_
                d_222_v33_: _dafny.Map
                d_222_v33_ = _dafny.Map({_dafny.MultiSet([d_217_cf6_]): (d_195_v20_)[default__.safeIndex(972, (d_195_v20_).length(0))]})
                index33_ = default__.safeIndex(302, (d_219_v32_).length(0))
                rhs43_ = (d_179_v5_ if d_217_cf6_ else default__.safeModuloInt(d_179_v5_, len(d_222_v33_)))
                rhs44_ = (d_173_v0_) + (_dafny.SeqWithoutIsStrInference([d_174_v2_ for d_223_i2_ in range(default__.abs(401))]))
                lhs32_ = d_219_v32_
                lhs33_ = default__.safeIndex(302, (d_219_v32_).length(0))
                lhs32_[lhs33_] = rhs43_
                d_173_v0_ = rhs44_
            elif source2_.is_DC3:
                d_224___mcc_h6_ = source2_.cf7
                d_225___mcc_h7_ = source2_.cf8
                d_226___mcc_h8_ = source2_.cf9
                d_227___mcc_h9_ = source2_.cf10
                d_228_cf10_ = d_227___mcc_h9_
                d_229_cf9_ = d_226___mcc_h8_
                d_230_cf8_ = d_225___mcc_h7_
                d_231_cf7_ = d_224___mcc_h6_
                d_232_v34_: _dafny.Seq
                d_232_v34_ = _dafny.SeqWithoutIsStrInference([d_194_v19_, (d_194_v19_).set(959, d_230_cf8_), d_194_v19_])
                d_233_v35_: _dafny.MultiSet
                d_233_v35_ = _dafny.MultiSet([d_174_v2_])
                d_234_v36_: _dafny.Seq
                d_234_v36_ = _dafny.SeqWithoutIsStrInference([d_230_cf8_])
                d_235_v37_: _dafny.MultiSet
                d_235_v37_ = _dafny.MultiSet([d_179_v5_, d_228_cf10_])
                d_236_v38_: _dafny.Array
                nw38_ = _dafny.Array(None, 29)
                nw38_[int(0)] = default__.safeModuloInt(d_179_v5_, d_231_cf7_)
                nw38_[int(1)] = (0) - ((d_179_v5_ if d_230_cf8_ else d_228_cf10_))
                nw38_[int(2)] = 11
                nw38_[int(3)] = (d_231_cf7_) + (d_229_cf9_)
                nw38_[int(4)] = d_229_cf9_
                nw38_[int(5)] = (default__.fm0(d_178_globalState_)) * (len(d_232_v34_))
                nw38_[int(6)] = d_179_v5_
                nw38_[int(7)] = d_179_v5_
                nw38_[int(8)] = len(_dafny.SeqWithoutIsStrInference([d_229_cf9_ for d_237_i3_ in range(default__.abs(693))]))
                nw38_[int(9)] = (d_229_cf9_) * (419)
                nw38_[int(10)] = 437
                nw38_[int(11)] = (((d_233_v35_).set(d_174_v2_, default__.abs(d_179_v5_))).cardinality) - (d_231_cf7_)
                nw38_[int(12)] = d_231_cf7_
                nw38_[int(13)] = len((default__.fm4(d_178_globalState_)) + (d_234_v36_))
                nw38_[int(14)] = d_231_cf7_
                nw38_[int(15)] = -540
                nw38_[int(16)] = d_228_cf10_
                nw38_[int(17)] = ((d_235_v37_)[d_228_cf10_] if (d_228_cf10_) in (d_235_v37_) else d_228_cf10_)
                nw38_[int(18)] = d_179_v5_
                nw38_[int(19)] = d_228_cf10_
                nw38_[int(20)] = d_179_v5_
                nw38_[int(21)] = d_179_v5_
                nw38_[int(22)] = (0) - (((0) - (len(d_173_v0_))) - (len(default__.fm5(d_228_cf10_, d_178_globalState_))))
                nw38_[int(23)] = d_228_cf10_
                nw38_[int(24)] = (d_179_v5_) * (-530)
                nw38_[int(25)] = len(_dafny.SeqWithoutIsStrInference([d_179_v5_, (0) - (len(d_173_v0_))]))
                nw38_[int(26)] = d_228_cf10_
                nw38_[int(27)] = d_179_v5_
                nw38_[int(28)] = d_228_cf10_
                d_236_v38_ = nw38_
                index34_ = default__.safeIndex(882, (d_236_v38_).length(0))
                (d_236_v38_)[index34_] = (d_229_cf9_) + (d_231_cf7_)
                index35_ = default__.safeIndex(882, (d_236_v38_).length(0))
                (d_236_v38_)[index35_] = len(d_173_v0_)
                d_238_v39_: D1
                d_238_v39_ = D1_DC4(d_174_v2_)
                pat_let_tv3_ = d_174_v2_
                def iife21_(_pat_let3_0):
                    def iife22_(d_239_dt__update__tmp_h0_):
                        def iife23_(_pat_let4_0):
                            def iife24_(d_240_dt__update_hcf11_h0_):
                                return D1_DC4(d_240_dt__update_hcf11_h0_)
                            return iife24_(_pat_let4_0)
                        return iife23_(pat_let_tv3_)
                    return iife22_(_pat_let3_0)
                (d_178_globalState_).f15 = (iife21_(d_238_v39_)).cf11
                d_241_v40_: _dafny.Array
                nw39_ = _dafny.Array(D0.default()(), 25)
                d_241_v40_ = nw39_
                d_242_v41_: _dafny.Set
                d_242_v41_ = _dafny.Set({d_241_v40_})
                index36_ = default__.safeIndex(882, (d_236_v38_).length(0))
                (d_236_v38_)[index36_] = len((d_242_v41_) | (d_242_v41_))
                d_230_cf8_ = ((d_195_v20_)[default__.safeIndex(972, (d_195_v20_).length(0))]) and (d_192_v17_)
            elif True:
                d_243___mcc_h10_ = source2_.cf0
                d_244_cf0_ = d_243___mcc_h10_
                (d_178_globalState_).f12 = not (d_192_v17_) or ((d_195_v20_)[default__.safeIndex(972, (d_195_v20_).length(0))])
                (d_178_globalState_).f3 = d_173_v0_
                (d_178_globalState_).f15 = d_174_v2_
                d_245_v42_: _dafny.Array
                out1_: _dafny.Array
                out1_ = default__.m0(default__.fm0(d_178_globalState_), d_178_globalState_)
                d_245_v42_ = out1_
            d_246_v43_: _dafny.MultiSet
            d_246_v43_ = _dafny.MultiSet([len(d_173_v0_)])
            index37_ = default__.safeIndex(972, (d_195_v20_).length(0))
            (d_195_v20_)[index37_] = not((d_246_v43_).issubset((d_246_v43_) | (d_246_v43_)))
            d_247_v44_: _dafny.Map
            d_247_v44_ = _dafny.Map({d_192_v17_: not(d_192_v17_)})
            d_248_v45_: _dafny.Array
            out2_: _dafny.Array
            out2_ = default__.m0(((d_246_v43_)[len(d_247_v44_)] if (len(d_247_v44_)) in (d_246_v43_) else d_179_v5_), d_178_globalState_)
            d_248_v45_ = out2_
        elif True:
            d_249_v46_: _dafny.Array
            out3_: _dafny.Array
            out3_ = default__.m0(d_179_v5_, d_178_globalState_)
            d_249_v46_ = out3_
            d_250_v47_: D0
            d_250_v47_ = D0_DC3(d_179_v5_, d_192_v17_, -428, d_179_v5_)
            if (d_250_v47_).cf8:
                d_251_v48_: _dafny.Seq
                d_251_v48_ = d_173_v0_
                (d_178_globalState_).f3 = (d_251_v48_)
                index38_ = default__.safeIndex(421, (d_249_v46_).length(0))
                (d_249_v46_)[index38_] = d_179_v5_
                index39_ = default__.safeIndex(421, (d_249_v46_).length(0))
                rhs45_ = d_179_v5_
                rhs46_ = d_192_v17_
                rhs47_ = (len(d_173_v0_)) - (d_179_v5_)
                rhs48_ = ((d_173_v0_) + (d_173_v0_)) < (d_173_v0_)
                lhs34_ = d_249_v46_
                lhs35_ = default__.safeIndex(421, (d_249_v46_).length(0))
                lhs36_ = d_178_globalState_
                d_179_v5_ = rhs45_
                d_192_v17_ = rhs46_
                lhs34_[lhs35_] = rhs47_
                lhs36_.f12 = rhs48_
                d_252_v49_: _dafny.Array
                out4_: _dafny.Array
                out4_ = default__.m0((d_249_v46_)[default__.safeIndex(421, (d_249_v46_).length(0))], d_178_globalState_)
                d_252_v49_ = out4_
                d_253_v50_: D0
                d_253_v50_ = D0_DC2(d_192_v17_)
                pat_let_tv4_ = d_192_v17_
                d_254_v51_: _dafny.Seq
                def iife25_(_pat_let5_0):
                    def iife26_(d_255_dt__update__tmp_h1_):
                        def iife27_(_pat_let6_0):
                            def iife28_(d_256_dt__update_hcf6_h0_):
                                return D0_DC2(d_256_dt__update_hcf6_h0_)
                            return iife28_(_pat_let6_0)
                        return iife27_(not(pat_let_tv4_))
                    return iife26_(_pat_let5_0)
                d_254_v51_ = _dafny.SeqWithoutIsStrInference([iife25_(d_253_v50_), d_253_v50_])
                d_254_v51_ = _dafny.SeqWithoutIsStrInference([d_253_v50_ for d_257_i4_ in range(default__.abs(344))])
                d_258_v52_: _dafny.Map
                d_258_v52_ = _dafny.Map({d_192_v17_: d_192_v17_})
                d_259_v53_: _dafny.Set
                d_259_v53_ = _dafny.Set({((d_258_v52_)[d_192_v17_] if (d_192_v17_) in (d_258_v52_) else d_192_v17_)})
                d_260_v54_: _dafny.Array
                out5_: _dafny.Array
                out5_ = default__.m0((97 if not(not(d_192_v17_)) else len(d_259_v53_)), d_178_globalState_)
                d_260_v54_ = out5_
            elif True:
                rhs49_ = (d_179_v5_) + (d_179_v5_)
                rhs50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfnyb"))
                lhs37_ = d_178_globalState_
                lhs38_ = d_178_globalState_
                lhs37_.f11 = rhs49_
                lhs38_.f3 = rhs50_
                d_261_v55_: _dafny.Set
                d_261_v55_ = _dafny.Set({d_192_v17_, d_192_v17_})
                d_262_v56_: _dafny.Map
                d_262_v56_ = _dafny.Map({d_192_v17_: d_261_v55_})
                (d_178_globalState_).f10 = (d_179_v5_) - (len(d_262_v56_))
                d_263_v57_: _dafny.Array
                out6_: _dafny.Array
                out6_ = default__.m0(d_179_v5_, d_178_globalState_)
                d_263_v57_ = out6_
                d_264_v58_: _dafny.Array
                out7_: _dafny.Array
                out7_ = default__.m0(d_179_v5_, d_178_globalState_)
                d_264_v58_ = out7_
                (d_178_globalState_).f12 = d_192_v17_
            d_265_v59_: _dafny.Seq
            d_265_v59_ = _dafny.SeqWithoutIsStrInference([d_192_v17_, True, d_192_v17_])
            d_266_v60_: _dafny.Array
            nw40_ = _dafny.Array(None, 8)
            nw40_[int(0)] = d_192_v17_
            nw40_[int(1)] = (d_179_v5_) > (d_179_v5_)
            nw40_[int(2)] = d_192_v17_
            nw40_[int(3)] = d_192_v17_
            nw40_[int(4)] = d_192_v17_
            nw40_[int(5)] = not(((d_194_v19_)[d_179_v5_] if (d_179_v5_) in (d_194_v19_) else d_192_v17_))
            nw40_[int(6)] = not((d_265_v59_)[default__.safeIndex(d_179_v5_, len(d_265_v59_))])
            nw40_[int(7)] = d_192_v17_
            d_266_v60_ = nw40_
            index40_ = default__.safeIndex(694, (d_266_v60_).length(0))
            (d_266_v60_)[index40_] = d_192_v17_
            d_267_v61_: _dafny.MultiSet
            d_267_v61_ = _dafny.MultiSet([d_192_v17_])
            index41_ = default__.safeIndex(694, (d_266_v60_).length(0))
            (d_266_v60_)[index41_] = default__.fm1(d_192_v17_, (_dafny.MultiSet([(d_267_v61_).cardinality, d_179_v5_, d_179_v5_])) != ((_dafny.MultiSet([d_179_v5_, d_179_v5_])).set(d_179_v5_, default__.abs(d_179_v5_))), (0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbxk"))), d_179_v5_)), d_178_globalState_)
            d_268_v62_: _dafny.Array
            out8_: _dafny.Array
            out8_ = default__.m0(5, d_178_globalState_)
            d_268_v62_ = out8_
            index42_ = default__.safeIndex(694, (d_266_v60_).length(0))
            (d_266_v60_)[index42_] = not (not((d_266_v60_)[default__.safeIndex(694, (d_266_v60_).length(0))])) or (not ((d_266_v60_)[default__.safeIndex(694, (d_266_v60_).length(0))]) or ((d_266_v60_)[default__.safeIndex(694, (d_266_v60_).length(0))]))
        d_269_v63_: _dafny.Array
        nw41_ = _dafny.Array(False, 7)
        d_269_v63_ = nw41_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_269_v63_).length(0)):
            d_270_i5_: int = guard_loop_0_
            if (True) and (((0) <= (d_270_i5_)) and ((d_270_i5_) < ((d_269_v63_).length(0)))):
                (d_269_v63_)[(d_270_i5_)] = (((_dafny.MultiSet([d_173_v0_, d_173_v0_, _dafny.SeqWithoutIsStrInference([d_174_v2_ for d_271_i6_ in range(default__.abs(873))])])).intersection(_dafny.MultiSet([d_173_v0_]))).cardinality) not in (_dafny.MultiSet([d_179_v5_]))
        d_272_v64_: _dafny.Set
        d_272_v64_ = _dafny.Set({d_192_v17_, default__.fm1(d_192_v17_, d_192_v17_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhaahdov"))), d_178_globalState_), d_192_v17_})
        d_192_v17_ = (d_272_v64_).ispropersubset(d_272_v64_)
        d_273_v65_: _dafny.Map
        d_273_v65_ = _dafny.Map({d_269_v63_: d_179_v5_})
        d_274_v66_: _dafny.Seq
        d_274_v66_ = _dafny.SeqWithoutIsStrInference([not(default__.fm1(not(not(d_192_v17_)), d_192_v17_, len(d_273_v65_), d_178_globalState_)), d_192_v17_, d_192_v17_, d_192_v17_, d_192_v17_])
        d_275_v67_: _dafny.Set
        d_275_v67_ = _dafny.Set({d_179_v5_, 706, d_179_v5_})
        d_276_v68_: D0
        d_276_v68_ = D0_DC3(d_179_v5_, (len(d_274_v66_)) < (d_179_v5_), d_179_v5_, (len(d_275_v67_)) * (d_179_v5_))
        d_276_v68_ = d_276_v68_
        d_277_v69_: _dafny.Map
        d_277_v69_ = _dafny.Map({d_193_v18_: d_192_v17_})
        hi2_ = 66
        for d_278_i7_ in range((0) - (len(d_277_v69_)), hi2_):
            d_279_v70_: C1
            nw42_ = C1()
            nw42_.ctor__(d_278_i7_)
            d_279_v70_ = nw42_
            (d_178_globalState_).f12 = (not (d_192_v17_) or (d_192_v17_)) == (d_192_v17_)
            (d_279_v70_).m2(d_269_v63_, (d_173_v0_) + (d_173_v0_), len(_dafny.SeqWithoutIsStrInference([d_279_v70_.f18])), d_178_globalState_)
            if not (d_192_v17_) or (d_192_v17_):
                nw43_ = C1()
                nw43_.ctor__(d_279_v70_.f18)
                d_279_v70_ = nw43_
                d_280_v71_: _dafny.Map
                d_280_v71_ = _dafny.Map({806: d_279_v70_})
                d_281_v72_: _dafny.Seq
                d_281_v72_ = _dafny.SeqWithoutIsStrInference([d_280_v71_])
                d_282_v73_: _dafny.Seq
                d_282_v73_ = _dafny.SeqWithoutIsStrInference([(d_281_v72_)[default__.safeIndex(d_279_v70_.f18, len(d_281_v72_))]])
                (d_178_globalState_).f12 = ((d_282_v73_)[default__.safeIndex(d_278_i7_, len(d_282_v73_))]) == (d_280_v71_)
                d_283_v74_: _dafny.Array
                nw44_ = _dafny.Array(_dafny.MultiSet({}), 29)
                d_283_v74_ = nw44_
                d_284_v75_: _dafny.Map
                d_284_v75_ = _dafny.Map({False: d_192_v17_})
                d_285_v76_: _dafny.Map
                d_285_v76_ = _dafny.Map({d_278_i7_: 556})
                d_286_v77_: _dafny.Array
                nw45_ = _dafny.Array(None, 22)
                nw45_[int(0)] = d_179_v5_
                nw45_[int(1)] = len(d_284_v75_)
                nw45_[int(2)] = d_278_i7_
                nw45_[int(3)] = d_279_v70_.f18
                nw45_[int(4)] = d_278_i7_
                nw45_[int(5)] = d_179_v5_
                nw45_[int(6)] = d_278_i7_
                nw45_[int(7)] = (0) - (-670)
                nw45_[int(8)] = d_179_v5_
                nw45_[int(9)] = 297
                nw45_[int(10)] = d_279_v70_.f18
                nw45_[int(11)] = d_279_v70_.f18
                nw45_[int(12)] = d_278_i7_
                nw45_[int(13)] = d_279_v70_.f18
                nw45_[int(14)] = d_279_v70_.f18
                nw45_[int(15)] = d_279_v70_.f18
                nw45_[int(16)] = d_279_v70_.f18
                nw45_[int(17)] = d_179_v5_
                nw45_[int(18)] = default__.fm0(d_178_globalState_)
                nw45_[int(19)] = len(d_285_v76_)
                nw45_[int(20)] = d_278_i7_
                nw45_[int(21)] = d_179_v5_
                d_286_v77_ = nw45_
                d_287_v78_: _dafny.Set
                d_287_v78_ = _dafny.Set({d_286_v77_})
                d_288_v79_: _dafny.Map
                d_288_v79_ = _dafny.Map({(d_287_v78_) | (d_287_v78_): (len(d_193_v18_)) * (d_179_v5_)})
                rhs51_ = ((d_288_v79_)[d_287_v78_] if (d_287_v78_) in (d_288_v79_) else 110)
                rhs52_ = d_283_v74_
                rhs53_ = default__.safeDivisionInt(d_279_v70_.f18, d_179_v5_)
                rhs54_ = d_192_v17_
                lhs39_ = d_178_globalState_
                lhs40_ = d_178_globalState_
                lhs39_.f10 = rhs51_
                d_283_v74_ = rhs52_
                d_179_v5_ = rhs53_
                lhs40_.f12 = rhs54_
                d_289_v80_: _dafny.Map
                d_289_v80_ = _dafny.Map({d_192_v17_: d_279_v70_})
                d_290_v81_: C1
                nw46_ = C1()
                nw46_.ctor__(len((d_289_v80_) | (d_289_v80_)))
                d_290_v81_ = nw46_
                (d_178_globalState_).f11 = d_279_v70_.f18
            elif True:
                d_291_v82_: C0
                nw47_ = C0()
                nw47_.ctor__()
                d_291_v82_ = nw47_
                d_292_v83_: _dafny.Map
                d_292_v83_ = _dafny.Map({d_291_v82_: d_279_v70_.f18})
                d_293_v84_: D3
                d_293_v84_ = D3_DC7(d_292_v83_)
                d_294_v85_: _dafny.Map
                d_294_v85_ = _dafny.Map({d_293_v84_: d_278_i7_})
                (d_279_v70_).f18 = default__.safeDivisionInt(len((_dafny.Map({d_293_v84_: d_279_v70_.f18})) | (d_294_v85_)), len(d_173_v0_))
                d_295_v86_: _dafny.MultiSet
                d_295_v86_ = _dafny.MultiSet([d_291_v82_, d_291_v82_])
                d_296_v87_: C1
                nw48_ = C1()
                nw48_.ctor__((d_295_v86_).cardinality)
                d_296_v87_ = nw48_
                index43_ = default__.safeIndex(948, (d_269_v63_).length(0))
                (d_269_v63_)[index43_] = d_192_v17_
                index44_ = default__.safeIndex(948, (d_269_v63_).length(0))
                (d_269_v63_)[index44_] = (d_279_v70_.f18) >= (d_296_v87_.f18)
                d_297_v88_: C0
                nw49_ = C0()
                nw49_.ctor__()
                d_297_v88_ = nw49_
                (d_178_globalState_).f10 = (0) - (d_279_v70_.f18)
        d_298_v89_: _dafny.Seq
        d_298_v89_ = _dafny.SeqWithoutIsStrInference([d_179_v5_, -221, d_179_v5_, d_179_v5_, len(d_173_v0_)])
        hi3_ = d_179_v5_
        for d_299_i8_ in range(len(d_298_v89_), hi3_):
            d_300_v90_: C0
            nw50_ = C0()
            nw50_.ctor__()
            d_300_v90_ = nw50_
            d_192_v17_ = d_192_v17_
            d_301_v91_: _dafny.Map
            d_301_v91_ = _dafny.Map({True: d_173_v0_})
            d_302_v92_: C1
            nw51_ = C1()
            nw51_.ctor__(default__.safeDivisionInt(741, len(d_301_v91_)))
            d_302_v92_ = nw51_
            d_303_v93_: _dafny.Map
            d_303_v93_ = _dafny.Map({d_192_v17_: d_269_v63_})
            d_303_v93_ = d_303_v93_
        d_304_v94_: _dafny.Array
        def lambda14_(d_305_v0_, d_306_v2_):
            def lambda15_(d_307_i9_):
                return (d_305_v0_) + ((d_305_v0_).set(default__.safeIndex(len(d_305_v0_), len(d_305_v0_)), d_306_v2_))

            return lambda15_

        init7_ = lambda14_(d_173_v0_, d_174_v2_)
        nw52_ = _dafny.Array(None, 7)
        for i0_7_ in range(nw52_.length(0)):
            nw52_[i0_7_] = init7_(i0_7_)
        d_304_v94_ = nw52_
        d_304_v94_ = d_304_v94_
        (d_178_globalState_).f12 = False
        d_308_v95_: _dafny.Array
        nw53_ = _dafny.Array(_dafny.Array(None, 0), 23)
        d_308_v95_ = nw53_
        d_308_v95_ = ((d_308_v95_ if d_192_v17_ else d_308_v95_) if d_192_v17_ else d_308_v95_)
        (d_178_globalState_).f10 = d_179_v5_
        d_309_v96_: _dafny.Map
        d_309_v96_ = _dafny.Map({d_179_v5_: d_179_v5_})
        (d_178_globalState_).f12 = (d_179_v5_) != ((-135) + (len((default__.fm4(d_178_globalState_)).set(default__.safeIndex(len(d_309_v96_), len(default__.fm4(d_178_globalState_))), d_192_v17_))))
        index45_ = default__.safeIndex(551, (d_304_v94_).length(0))
        (d_304_v94_)[index45_] = _dafny.SeqWithoutIsStrInference([d_174_v2_ for d_310_i10_ in range(default__.abs(238))])
        d_311_v97_: _dafny.Seq
        d_311_v97_ = _dafny.SeqWithoutIsStrInference([d_173_v0_])
        index46_ = default__.safeIndex(551, (d_304_v94_).length(0))
        (d_304_v94_)[index46_] = (d_311_v97_)[default__.safeIndex(d_179_v5_, len(d_311_v97_))]
        if (d_192_v17_ if d_192_v17_ else d_192_v17_):
            d_312_v98_: _dafny.Array
            out9_: _dafny.Array
            out9_ = default__.m0(d_179_v5_, d_178_globalState_)
            d_312_v98_ = out9_
            (d_178_globalState_).f12 = (default__.safeDivisionInt(d_179_v5_, d_179_v5_)) >= (d_179_v5_)
            (d_178_globalState_).f10 = ((d_309_v96_)[465] if (465) in (d_309_v96_) else len(d_298_v89_))
            (d_178_globalState_).f10 = d_179_v5_
            d_313_v99_: _dafny.Array
            nw54_ = _dafny.Array(_dafny.Seq({}), 17)
            d_313_v99_ = nw54_
            index47_ = default__.safeIndex(514, (d_313_v99_).length(0))
            (d_313_v99_)[index47_] = _dafny.SeqWithoutIsStrInference([d_179_v5_, default__.fm0(d_178_globalState_)])
            index48_ = default__.safeIndex(514, (d_313_v99_).length(0))
            (d_313_v99_)[index48_] = (_dafny.SeqWithoutIsStrInference([-759])) + (_dafny.SeqWithoutIsStrInference([d_179_v5_, d_179_v5_, 877]))
        elif True:
            d_179_v5_ = (d_179_v5_) - (len(d_274_v66_))
            d_314_v100_: _dafny.Seq
            d_314_v100_ = _dafny.SeqWithoutIsStrInference([d_269_v63_, d_269_v63_, d_269_v63_])
            d_315_v101_: C1
            nw55_ = C1()
            nw55_.ctor__(len((d_314_v100_ if d_192_v17_ else d_314_v100_)))
            d_315_v101_ = nw55_
            d_315_v101_ = d_315_v101_
            d_315_v101_ = d_315_v101_
            if not((d_179_v5_) >= (d_179_v5_)):
                d_316_v102_: _dafny.Array
                nw56_ = _dafny.Array(_dafny.Set({}), 8)
                d_316_v102_ = nw56_
                d_317_v103_: _dafny.Array
                nw57_ = _dafny.Array(_dafny.Set({}), 20)
                d_317_v103_ = nw57_
                d_318_v104_: _dafny.Array
                d_318_v104_ = d_317_v103_
                d_319_v105_: _dafny.Set
                d_319_v105_ = _dafny.Set({d_318_v104_, d_318_v104_, d_317_v103_, d_318_v104_})
                index49_ = default__.safeIndex(230, (d_316_v102_).length(0))
                (d_316_v102_)[index49_] = (d_319_v105_ if d_192_v17_ else d_319_v105_)
                index50_ = default__.safeIndex(230, (d_316_v102_).length(0))
                (d_316_v102_)[index50_] = d_319_v105_
                d_320_v106_: _dafny.Array
                out10_: _dafny.Array
                out10_ = default__.m0(d_179_v5_, d_178_globalState_)
                d_320_v106_ = out10_
                d_192_v17_ = d_192_v17_
                index51_ = default__.safeIndex(553, (d_320_v106_).length(0))
                (d_320_v106_)[index51_] = d_315_v101_.f18
                index52_ = default__.safeIndex(553, (d_320_v106_).length(0))
                rhs55_ = d_315_v101_.f18
                rhs56_ = d_269_v63_
                lhs41_ = d_320_v106_
                lhs42_ = default__.safeIndex(553, (d_320_v106_).length(0))
                lhs41_[lhs42_] = rhs55_
                d_269_v63_ = rhs56_
                d_321_v107_: C0
                nw58_ = C0()
                nw58_.ctor__()
                d_321_v107_ = nw58_
                d_322_v108_: _dafny.Map
                d_322_v108_ = _dafny.Map({not(d_192_v17_): d_321_v107_})
                (d_178_globalState_).f12 = default__.fm1(not((len((d_322_v108_).set(d_192_v17_, d_321_v107_))) == (len(d_275_v67_))), d_192_v17_, d_179_v5_, d_178_globalState_)
            elif True:
                (d_315_v101_).m2(d_269_v63_, d_173_v0_, default__.safeModuloInt((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_174_v2_ for d_323_i11_ in range(default__.abs(237))])])).cardinality, 243), d_178_globalState_)
                d_274_v66_ = (_dafny.SeqWithoutIsStrInference([default__.fm1(d_192_v17_, d_192_v17_, d_179_v5_, d_178_globalState_)])) + ((d_274_v66_) + (d_274_v66_))
                (d_178_globalState_).f11 = (0) - (len((d_304_v94_)[default__.safeIndex(551, (d_304_v94_).length(0))]))
                (d_178_globalState_).f12 = ((0) - (d_179_v5_)) == ((len(d_193_v18_)) - (d_179_v5_))
                d_324_v109_: C1
                nw59_ = C1()
                nw59_.ctor__(d_315_v101_.f18)
                d_324_v109_ = nw59_
            d_325_v110_: _dafny.Seq
            d_325_v110_ = _dafny.SeqWithoutIsStrInference([d_315_v101_, d_315_v101_, d_315_v101_, d_315_v101_])
            d_315_v101_ = (d_325_v110_)[default__.safeIndex(d_179_v5_, len(d_325_v110_))]
        d_326_v111_: _dafny.MultiSet
        d_326_v111_ = _dafny.MultiSet([d_192_v17_, d_192_v17_, default__.fm1(d_192_v17_, d_192_v17_, d_179_v5_, d_178_globalState_), d_192_v17_])
        d_327_v112_: _dafny.Map
        d_327_v112_ = _dafny.Map({d_192_v17_: d_192_v17_})
        d_328_v113_: _dafny.Map
        d_328_v113_ = _dafny.Map({d_274_v66_: not(((d_327_v112_)[d_192_v17_] if (d_192_v17_) in (d_327_v112_) else d_192_v17_))})
        (d_178_globalState_).f12 = (default__.fm16((d_326_v111_).cardinality, d_178_globalState_)) != (d_328_v113_)
        d_329_v114_: _dafny.Array
        nw60_ = _dafny.Array(None, 3)
        d_329_v114_ = nw60_
        d_330_v115_: C0
        nw61_ = C0()
        nw61_.ctor__()
        d_330_v115_ = nw61_
        index53_ = default__.safeIndex(204, (d_329_v114_).length(0))
        (d_329_v114_)[index53_] = d_330_v115_
        index54_ = default__.safeIndex(204, (d_329_v114_).length(0))
        (d_329_v114_)[index54_] = d_330_v115_
        _dafny.print((d_173_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_174_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v3_) == (_dafny.Set({_dafny.CodePoint('k')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v4_) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('n')}), _dafny.Set({_dafny.CodePoint('k')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_178_globalState_.f3).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_178_globalState_).f4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_178_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_178_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_178_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_178_globalState_).f14) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('n')}), _dafny.Set({_dafny.CodePoint('k')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_178_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_179_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_192_v17_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v18_) == (_dafny.Map({False: -2161}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v19_) == (_dafny.Map({2: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v63_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v63_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v63_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v63_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v63_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v63_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v63_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v64_) == (_dafny.Set({False, True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_273_v65_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v66_) == (_dafny.SeqWithoutIsStrInference([True, False, False, False, False, False, False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v67_) == (_dafny.Set({-2161, 706}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_276_v68_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_276_v68_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_276_v68_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_276_v68_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_v69_) == (_dafny.Map({_dafny.Map({False: -2161}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_298_v89_) == (_dafny.SeqWithoutIsStrInference([1, -221, 1, 1, 403]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_304_v94_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_304_v94_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_304_v94_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_304_v94_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_304_v94_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_304_v94_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_304_v94_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_v96_) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v97_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yokkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_326_v111_) == (_dafny.MultiSet([False, False, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v112_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_328_v113_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, False, False, False, False, False, False, False, False, False, False]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), int(0), int(0), _dafny.Array(None, 0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
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
        return lambda: D1_DC5(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({self.cf14.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(_dafny.Seq({}), False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC8(D3, NamedTuple('DC8', [('cf16', Any), ('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {self.cf19.VerbatimString(True)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(_dafny.CodePoint('D'), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC10(D4, NamedTuple('DC10', [('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC12(D5, NamedTuple('DC12', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC14(False, _dafny.CodePoint('D'), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC14(D6, NamedTuple('DC14', [('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)

class D7_DC15(D7, NamedTuple('DC15', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC17(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D8_DC16)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC17(D8, NamedTuple('DC17', [('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({self.cf35.VerbatimString(True)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC16(D8, NamedTuple('DC16', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC16({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC16) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f10: int = int(0)
        self.f11: int = int(0)
        self.f12: bool = False
        self.f15: str = _dafny.CodePoint('D')
        self._f0: int = int(0)
        self._f1: bool = False
        self._f2: int = int(0)
        self._f4: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f5: int = int(0)
        self._f6: int = int(0)
        self._f7: bool = False
        self._f8: bool = False
        self._f9: int = int(0)
        self._f13: bool = False
        self._f14: _dafny.MultiSet = _dafny.MultiSet({})
        self._f16: int = int(0)
        self._f17: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
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
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, p1, globalState):
        return _dafny.MultiSet([(0) - ((-495) * (len(_dafny.Set({not(False), False})))), (-643 if False else 993)])

    def fm7(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdmixbs"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "blliiliu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vthendu"))))


class C1:
    def  __init__(self):
        self.f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f18):
        (self).f18 = f18

    def m1(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_331_v0_: _dafny.Array
        nw62_ = _dafny.Array(_dafny.CodePoint('D'), 22)
        d_331_v0_ = nw62_
        d_332_v1_: str
        d_332_v1_ = _dafny.CodePoint('a')
        index55_ = default__.safeIndex(96, (d_331_v0_).length(0))
        (d_331_v0_)[index55_] = d_332_v1_
        index56_ = default__.safeIndex(96, (d_331_v0_).length(0))
        (d_331_v0_)[index56_] = _dafny.CodePoint('n')
        d_333_v2_: _dafny.Array
        nw63_ = _dafny.Array(False, 27)
        d_333_v2_ = nw63_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_333_v2_).length(0)):
            d_334_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_334_i0_)) and ((d_334_i0_) < ((d_333_v2_).length(0)))):
                (d_333_v2_)[(d_334_i0_)] = p1
        (self).f18 = p0
        d_335_v3_: _dafny.Array
        out11_: _dafny.Array
        out11_ = default__.m0(p0, globalState)
        d_335_v3_ = out11_
        hi4_ = default__.fm0(globalState)
        for d_336_i1_ in range(self.f18, hi4_):
            index57_ = default__.safeIndex(862, (d_333_v2_).length(0))
            (d_333_v2_)[index57_] = True
            d_337_v4_: _dafny.MultiSet
            d_337_v4_ = _dafny.MultiSet([d_336_i1_])
            index58_ = default__.safeIndex(862, (d_333_v2_).length(0))
            (d_333_v2_)[index58_] = (d_337_v4_).ispropersubset(d_337_v4_)
            d_338_v5_: C0
            nw64_ = C0()
            nw64_.ctor__()
            d_338_v5_ = nw64_
            d_339_v6_: _dafny.Seq
            d_339_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfkwvkpem"))
            d_340_v7_: _dafny.Map
            d_340_v7_ = _dafny.Map({d_339_v6_: 938})
            d_340_v7_ = d_340_v7_
            index59_ = default__.safeIndex(941, (d_335_v3_).length(0))
            (d_335_v3_)[index59_] = (0) - (p3)
            index60_ = default__.safeIndex(941, (d_335_v3_).length(0))
            (d_335_v3_)[index60_] = (0) - (((0) - ((self.f18) + (d_336_i1_))) + (self.f18))
        if default__.fm1(p1, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrme"))) < (default__.fm8(globalState)), p3, globalState):
            d_341_v8_: _dafny.Map
            d_341_v8_ = _dafny.Map({d_335_v3_: D0_DC0(p3)})
            d_341_v8_ = d_341_v8_
            d_342_v9_: _dafny.MultiSet
            d_342_v9_ = _dafny.MultiSet([True, p1])
            d_343_v10_: _dafny.Map
            d_343_v10_ = _dafny.Map({d_331_v0_: (d_342_v9_).ispropersubset(_dafny.MultiSet([p1, p1, p1, p1]))})
            d_344_v11_: _dafny.Map
            d_344_v11_ = _dafny.Map({default__.fm0(globalState): d_331_v0_})
            d_343_v10_ = (d_343_v10_).set(((d_344_v11_)[p0] if (p0) in (d_344_v11_) else d_331_v0_), p1)
            (globalState).f11 = p0
            (globalState).f12 = p1
            d_345_v12_: D0
            d_345_v12_ = D0_DC0(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_346_i2_ in range(default__.abs(-927))])))
            d_347_v13_: _dafny.Set
            d_347_v13_ = _dafny.Set({D0_DC0(p3), d_345_v12_, D0_DC0(537)})
            index61_ = default__.safeIndex(916, (d_333_v2_).length(0))
            (d_333_v2_)[index61_] = (d_347_v13_).ispropersubset(d_347_v13_)
            index62_ = default__.safeIndex(916, (d_333_v2_).length(0))
            rhs57_ = False
            rhs58_ = p1
            lhs43_ = globalState
            lhs44_ = d_333_v2_
            lhs45_ = default__.safeIndex(916, (d_333_v2_).length(0))
            lhs43_.f12 = rhs57_
            lhs44_[lhs45_] = rhs58_
        elif True:
            d_348_v14_: _dafny.Seq
            d_348_v14_ = _dafny.SeqWithoutIsStrInference([(d_331_v0_)[default__.safeIndex(96, (d_331_v0_).length(0))], (d_331_v0_)[default__.safeIndex(96, (d_331_v0_).length(0))]])
            d_349_v15_: _dafny.Seq
            d_349_v15_ = d_348_v14_
            d_350_v16_: _dafny.Map
            d_350_v16_ = _dafny.Map({d_349_v15_: p1})
            d_350_v16_ = _dafny.Map({d_348_v14_: True})
            d_351_v17_: _dafny.Seq
            d_351_v17_ = _dafny.SeqWithoutIsStrInference([p1])
            d_352_v18_: _dafny.Map
            d_352_v18_ = _dafny.Map({d_348_v14_: d_351_v17_})
            d_352_v18_ = _dafny.Map({(d_348_v14_) + (d_348_v14_): d_351_v17_})
            d_353_v19_: _dafny.Map
            d_353_v19_ = _dafny.Map({d_333_v2_: p1})
            d_354_v20_: D1
            d_354_v20_ = D1_DC4(d_332_v1_)
            rhs59_ = (d_353_v19_) | (d_353_v19_)
            rhs60_ = (0) - (p0)
            rhs61_ = p3
            rhs62_ = d_354_v20_
            lhs46_ = globalState
            lhs47_ = globalState
            d_353_v19_ = rhs59_
            lhs46_.f11 = rhs60_
            lhs47_.f10 = rhs61_
            d_354_v20_ = rhs62_
            r0 = d_333_v2_
            d_355_v21_: _dafny.Map
            d_355_v21_ = _dafny.Map({p1: p1})
            d_356_v22_: _dafny.Set
            d_356_v22_ = _dafny.Set({d_355_v21_, _dafny.Map({p1: p1}), d_355_v21_})
            (globalState).f12 = (d_356_v22_).isdisjoint(d_356_v22_)
        r0 = d_333_v2_
        r1 = p3
        return r0, r1

    def m2(self, p0, p1, p2, globalState):
        d_357_v0_: _dafny.Array
        nw65_ = _dafny.Array(None, 2)
        nw65_[int(0)] = (0) - (p2)
        nw65_[int(1)] = (self.f18) - (-962)
        d_357_v0_ = nw65_
        d_358_v1_: bool
        d_358_v1_ = False
        d_357_v0_ = (d_357_v0_ if d_358_v1_ else d_357_v0_)
        d_359_i0_: int
        d_359_i0_ = 0
        with _dafny.label("1"):
            while ((self.f18) >= (181) if True else (p2) >= (p2)):
                with _dafny.c_label("1"):
                    if (d_359_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_359_i0_ = (d_359_i0_) + (1)
                    d_360_v2_: _dafny.Array
                    def lambda16_(d_361_v1_, d_362_p2_):
                        def lambda17_(d_363_i1_):
                            return (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_361_v1_, d_361_v1_, d_361_v1_]))).cardinality for d_364_i2_ in range(default__.abs(145))]) if d_361_v1_ else _dafny.SeqWithoutIsStrInference([d_362_p2_ for d_365_i3_ in range(default__.abs(154))]))

                        return lambda17_

                    init8_ = lambda16_(d_358_v1_, p2)
                    nw66_ = _dafny.Array(None, 25)
                    for i0_8_ in range(nw66_.length(0)):
                        nw66_[i0_8_] = init8_(i0_8_)
                    d_360_v2_ = nw66_
                    d_366_v3_: _dafny.Seq
                    d_366_v3_ = _dafny.SeqWithoutIsStrInference([p2])
                    index63_ = default__.safeIndex(869, (d_360_v2_).length(0))
                    (d_360_v2_)[index63_] = d_366_v3_
                    index64_ = default__.safeIndex(869, (d_360_v2_).length(0))
                    (d_360_v2_)[index64_] = d_366_v3_
                    (globalState).f12 = (not(d_358_v1_) if (len(p1)) >= (len(p1)) else d_358_v1_)
                    d_367_v4_: C0
                    nw67_ = C0()
                    nw67_.ctor__()
                    d_367_v4_ = nw67_
                    d_368_v5_: _dafny.Map
                    d_368_v5_ = _dafny.Map({d_367_v4_: (self.f18) - (p2)})
                    index65_ = default__.safeIndex(864, (p0).length(0))
                    (p0)[index65_] = not(d_358_v1_)
                    d_369_v6_: D3
                    d_369_v6_ = D3_DC7(d_368_v5_)
                    d_370_v7_: D0
                    d_370_v7_ = D0_DC1(len(_dafny.SeqWithoutIsStrInference([self.f18 for d_371_i4_ in range(default__.abs(908))])), self.f18, p2, self.f18, p0)
                    d_372_v8_: _dafny.Seq
                    d_372_v8_ = _dafny.SeqWithoutIsStrInference([D0_DC1(250, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mocct"))), (0) - (self.f18), self.f18, p0)])
                    index66_ = default__.safeIndex(864, (p0).length(0))
                    rhs63_ = (d_369_v6_).cf15
                    rhs64_ = d_358_v1_
                    rhs65_ = (d_370_v7_) not in (d_372_v8_)
                    rhs66_ = len((p1) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_373_i5_ in range(default__.abs(25))])))
                    lhs48_ = p0
                    lhs49_ = default__.safeIndex(864, (p0).length(0))
                    lhs50_ = globalState
                    d_368_v5_ = rhs63_
                    lhs48_[lhs49_] = rhs64_
                    d_358_v1_ = rhs65_
                    lhs50_.f10 = rhs66_
                    d_374_v9_: D0
                    d_374_v9_ = D0_DC2(d_358_v1_)
                    d_375_v10_: _dafny.Map
                    d_375_v10_ = _dafny.Map({(p0)[default__.safeIndex(864, (p0).length(0))]: d_374_v9_})
                    d_376_v11_: _dafny.Map
                    d_376_v11_ = _dafny.Map({d_375_v10_: self.f18})
                    d_377_v12_: str
                    d_377_v12_ = _dafny.CodePoint('c')
                    d_378_v13_: _dafny.MultiSet
                    d_378_v13_ = _dafny.MultiSet([len((p1).set(default__.safeIndex(p2, len(p1)), d_377_v12_))])
                    d_379_v15_: _dafny.Set
                    d_379_v15_ = _dafny.Set({d_375_v10_, d_375_v10_})
                    def iife29_():
                        coll15_ = _dafny.Map()
                        compr_15_: _dafny.Map
                        for compr_15_ in (d_379_v15_).Elements:
                            d_380_v14_: _dafny.Map = compr_15_
                            if (d_380_v14_) in (d_379_v15_):
                                coll15_[d_380_v14_] = 2
                        return _dafny.Map(coll15_)
                    (globalState).f12 = (d_376_v11_) == ((_dafny.Map({d_375_v10_: (d_378_v13_).cardinality})) | (iife29_()
                    ))
                    pass
            pass
        d_381_v16_: _dafny.Array
        def lambda18_(d_382_p1_, d_383_p2_, d_384_v1_):
            def lambda19_(d_385_i7_):
                return D3_DC8(_dafny.SeqWithoutIsStrInference([896, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len((d_382_p1_).set(default__.safeIndex(505, len(d_382_p1_)), _dafny.CodePoint('v'))), d_383_p2_]))).cardinality]), d_384_v1_, d_384_v1_, d_382_p1_, d_384_v1_)

            return lambda19_

        init9_ = lambda18_(p1, p2, d_358_v1_)
        nw68_ = _dafny.Array(None, 17)
        for i0_9_ in range(nw68_.length(0)):
            nw68_[i0_9_] = init9_(i0_9_)
        d_381_v16_ = nw68_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_381_v16_).length(0)):
            d_386_i6_: int = guard_loop_2_
            if (True) and (((0) <= (d_386_i6_)) and ((d_386_i6_) < ((d_381_v16_).length(0)))):
                (d_381_v16_)[(d_386_i6_)] = (D3_DC8(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_358_v1_: self.f18})) for d_387_i8_ in range(default__.abs(998))]), d_358_v1_, d_358_v1_, p1, d_358_v1_) if False else D3_DC8(_dafny.SeqWithoutIsStrInference([p2 for d_388_i9_ in range(default__.abs(49))]), not(d_358_v1_), d_358_v1_, p1, d_358_v1_))
        if False:
            if (p2) >= (p2):
                d_389_v17_: C0
                nw69_ = C0()
                nw69_.ctor__()
                d_389_v17_ = nw69_
                rhs67_ = d_358_v1_
                rhs68_ = True
                rhs69_ = p2
                rhs70_ = p2
                lhs51_ = globalState
                lhs52_ = globalState
                lhs53_ = globalState
                lhs54_ = globalState
                lhs51_.f12 = rhs67_
                lhs52_.f12 = rhs68_
                lhs53_.f10 = rhs69_
                lhs54_.f11 = rhs70_
                d_390_v18_: _dafny.Map
                d_390_v18_ = _dafny.Map({not(d_358_v1_): p1})
                d_391_v19_: _dafny.Seq
                d_391_v19_ = p1
                d_392_v20_: _dafny.Seq
                d_392_v20_ = _dafny.SeqWithoutIsStrInference([self.f18])
                d_393_v21_: D3
                d_393_v21_ = D3_DC8(d_392_v20_, d_358_v1_, d_358_v1_, p1, d_358_v1_)
                d_394_v22_: _dafny.Array
                nw70_ = _dafny.Array(None, 29)
                nw70_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lglnyxquc"))
                nw70_[int(1)] = p1
                nw70_[int(2)] = ((d_390_v18_)[d_358_v1_] if (d_358_v1_) in (d_390_v18_) else p1)
                nw70_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_395_i10_ in range(default__.abs(-156))])
                nw70_[int(4)] = p1
                nw70_[int(5)] = p1
                nw70_[int(6)] = (p1)
                nw70_[int(7)] = p1
                nw70_[int(8)] = _dafny.SeqWithoutIsStrInference([(D1_DC4(_dafny.CodePoint('l'))).cf11 for d_396_i11_ in range(default__.abs(738))])
                nw70_[int(9)] = p1
                nw70_[int(10)] = p1
                nw70_[int(11)] = (p1) + ((d_393_v21_).cf19)
                nw70_[int(12)] = p1
                nw70_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfqwixbp"))
                nw70_[int(14)] = p1
                nw70_[int(15)] = p1
                nw70_[int(16)] = p1
                nw70_[int(17)] = (p1) + (p1)
                nw70_[int(18)] = p1
                nw70_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyqwj"))
                nw70_[int(20)] = (d_393_v21_).cf19
                nw70_[int(21)] = p1
                nw70_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqevck"))
                nw70_[int(23)] = p1
                nw70_[int(24)] = (p1) + (p1)
                nw70_[int(25)] = p1
                nw70_[int(26)] = p1
                nw70_[int(27)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_397_i12_ in range(default__.abs(307))])
                nw70_[int(28)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgyrwllh"))
                d_394_v22_ = nw70_
                index67_ = default__.safeIndex(767, (d_394_v22_).length(0))
                (d_394_v22_)[index67_] = p1
                index68_ = default__.safeIndex(767, (d_394_v22_).length(0))
                (d_394_v22_)[index68_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_398_i13_ in range(default__.abs(360))])
                d_399_v23_: _dafny.Seq
                d_399_v23_ = _dafny.SeqWithoutIsStrInference([d_389_v17_])
                d_389_v17_ = (d_399_v23_)[default__.safeIndex((self.f18) - ((0) - (self.f18)), len(d_399_v23_))]
                d_400_v24_: _dafny.Array
                nw71_ = _dafny.Array(_dafny.Seq({}), 7)
                d_400_v24_ = nw71_
                d_401_v25_: _dafny.Seq
                d_401_v25_ = _dafny.SeqWithoutIsStrInference([d_357_v0_, d_357_v0_, d_357_v0_, d_357_v0_])
                index69_ = default__.safeIndex(798, (d_400_v24_).length(0))
                (d_400_v24_)[index69_] = ((_dafny.SeqWithoutIsStrInference([d_357_v0_, d_357_v0_, d_357_v0_])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_357_v0_, d_357_v0_, d_357_v0_]))), d_357_v0_)) + (d_401_v25_)
                index70_ = default__.safeIndex(798, (d_400_v24_).length(0))
                (d_400_v24_)[index70_] = (d_401_v25_) + (_dafny.SeqWithoutIsStrInference([d_357_v0_, d_357_v0_]))
            elif True:
                d_357_v0_ = d_357_v0_
                (globalState).f15 = default__.fm9(d_358_v1_, d_358_v1_, globalState)
                d_402_v26_: _dafny.Array
                def lambda20_(d_403_i14_):
                    return _dafny.CodePoint('l')

                init10_ = lambda20_
                nw72_ = _dafny.Array(None, 20)
                for i0_10_ in range(nw72_.length(0)):
                    nw72_[i0_10_] = init10_(i0_10_)
                d_402_v26_ = nw72_
                d_402_v26_ = d_402_v26_
                d_358_v1_ = (d_358_v1_ if d_358_v1_ else d_358_v1_)
                d_404_v27_: C0
                nw73_ = C0()
                nw73_.ctor__()
                d_404_v27_ = nw73_
            d_405_v28_: str
            d_405_v28_ = _dafny.CodePoint('h')
            d_406_v29_: _dafny.Map
            d_406_v29_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_407_i15_ in range(default__.abs(350))]): d_405_v28_})
            d_406_v29_ = (d_406_v29_).set((p1) + (p1), d_405_v28_)
            d_408_v30_: C0
            nw74_ = C0()
            nw74_.ctor__()
            d_408_v30_ = nw74_
            d_409_v31_: _dafny.Map
            d_409_v31_ = _dafny.Map({d_408_v30_: (0) - (p2)})
            d_409_v31_ = (d_409_v31_).set(d_408_v30_, self.f18)
            d_410_v32_: D0
            d_410_v32_ = D0_DC1(len(p1), p2, p2, (0) - (default__.safeDivisionInt((0) - (self.f18), p2)), p0)
            source3_ = d_410_v32_
            if source3_.is_DC1:
                d_411___mcc_h0_ = source3_.cf1
                d_412___mcc_h1_ = source3_.cf2
                d_413___mcc_h2_ = source3_.cf3
                d_414___mcc_h3_ = source3_.cf4
                d_415___mcc_h4_ = source3_.cf5
                d_416_cf5_ = d_415___mcc_h4_
                d_417_cf4_ = d_414___mcc_h3_
                d_418_cf3_ = d_413___mcc_h2_
                d_419_cf2_ = d_412___mcc_h1_
                d_420_cf1_ = d_411___mcc_h0_
                (globalState).f12 = (default__.safeModuloInt(d_417_cf4_, self.f18)) <= (d_418_cf3_)
                index71_ = default__.safeIndex(843, (d_416_cf5_).length(0))
                (d_416_cf5_)[index71_] = d_358_v1_
                d_421_v33_: _dafny.MultiSet
                d_421_v33_ = _dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_420_cf1_ for d_422_i16_ in range(default__.abs(-308))]))])
                d_423_v34_: _dafny.MultiSet
                d_423_v34_ = _dafny.MultiSet([853, self.f18])
                d_424_v35_: _dafny.Seq
                d_424_v35_ = _dafny.SeqWithoutIsStrInference([d_358_v1_, d_358_v1_])
                d_425_v36_: _dafny.Map
                d_425_v36_ = _dafny.Map({d_405_v28_: False})
                d_426_v37_: _dafny.MultiSet
                d_426_v37_ = _dafny.MultiSet([((d_425_v36_)[d_405_v28_] if (d_405_v28_) in (d_425_v36_) else d_358_v1_)])
                d_427_v38_: _dafny.Map
                d_427_v38_ = _dafny.Map({d_358_v1_: self.f18})
                index72_ = default__.safeIndex(843, (d_416_cf5_).length(0))
                rhs71_ = (d_421_v33_) == (((d_421_v33_).set((d_423_v34_).set(p2, default__.abs(len(_dafny.Map({not(d_358_v1_): d_358_v1_})))), default__.abs(p2))) | (d_421_v33_))
                rhs72_ = (d_426_v37_).issubset(_dafny.MultiSet((d_424_v35_) + (d_424_v35_)))
                rhs73_ = ((d_408_v30_).fm7(d_358_v1_, self.f18, d_427_v38_, p1, globalState)) + (p1)
                rhs74_ = (0) - (default__.safeDivisionInt(d_417_cf4_, d_420_cf1_))
                lhs55_ = d_416_cf5_
                lhs56_ = default__.safeIndex(843, (d_416_cf5_).length(0))
                lhs57_ = globalState
                lhs58_ = globalState
                lhs55_[lhs56_] = rhs71_
                d_358_v1_ = rhs72_
                lhs57_.f3 = rhs73_
                lhs58_.f10 = rhs74_
                (globalState).f3 = p1
                d_428_v39_: D0
                d_428_v39_ = D0_DC0(len(d_427_v38_))
                d_429_v40_: _dafny.Seq
                d_429_v40_ = _dafny.SeqWithoutIsStrInference([d_428_v39_])
                d_429_v40_ = (default__.fm10(d_358_v1_, d_405_v28_, globalState)) + (d_429_v40_)
            elif source3_.is_DC2:
                d_430___mcc_h5_ = source3_.cf6
                d_431_cf6_ = d_430___mcc_h5_
                rhs75_ = False
                rhs76_ = (default__.fm0(globalState)) - (self.f18)
                rhs77_ = default__.safeModuloInt(default__.safeDivisionInt(p2, len(p1)), self.f18)
                lhs59_ = globalState
                lhs60_ = globalState
                d_358_v1_ = rhs75_
                lhs59_.f11 = rhs76_
                lhs60_.f10 = rhs77_
                d_432_v41_: _dafny.Map
                d_432_v41_ = _dafny.Map({d_431_cf6_: d_431_cf6_})
                d_432_v41_ = ((d_432_v41_).set(d_358_v1_, d_431_cf6_)) | (d_432_v41_)
                index73_ = default__.safeIndex(647, (p0).length(0))
                (p0)[index73_] = d_431_cf6_
                index74_ = default__.safeIndex(647, (p0).length(0))
                (p0)[index74_] = False
                d_433_v42_: _dafny.Seq
                d_433_v42_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(647, (p0).length(0))], (p0)[default__.safeIndex(647, (p0).length(0))], d_431_cf6_])
                d_433_v42_ = d_433_v42_
            elif source3_.is_DC3:
                d_434___mcc_h6_ = source3_.cf7
                d_435___mcc_h7_ = source3_.cf8
                d_436___mcc_h8_ = source3_.cf9
                d_437___mcc_h9_ = source3_.cf10
                d_438_cf10_ = d_437___mcc_h9_
                d_439_cf9_ = d_436___mcc_h8_
                d_440_cf8_ = d_435___mcc_h7_
                d_441_cf7_ = d_434___mcc_h6_
                index75_ = default__.safeIndex(301, (p0).length(0))
                (p0)[index75_] = d_358_v1_
                index76_ = default__.safeIndex(301, (p0).length(0))
                (p0)[index76_] = not (d_358_v1_) or (d_358_v1_)
                d_440_cf8_ = (_dafny.MultiSet(default__.fm4(globalState))).isdisjoint(_dafny.MultiSet([d_440_cf8_]))
                d_442_v43_: D4
                d_442_v43_ = D4_DC9(d_408_v30_)
                d_443_v44_: _dafny.Map
                d_443_v44_ = _dafny.Map({default__.fm1(d_440_cf8_, d_440_cf8_, self.f18, globalState): d_408_v30_})
                d_444_v45_: _dafny.Array
                nw75_ = _dafny.Array(None, 17)
                nw75_[int(0)] = d_408_v30_
                nw75_[int(1)] = d_408_v30_
                nw75_[int(2)] = (d_442_v43_).cf21
                nw75_[int(3)] = d_408_v30_
                nw75_[int(4)] = d_408_v30_
                nw75_[int(5)] = d_408_v30_
                nw75_[int(6)] = d_408_v30_
                nw75_[int(7)] = d_408_v30_
                nw75_[int(8)] = d_408_v30_
                nw75_[int(9)] = d_408_v30_
                nw75_[int(10)] = d_408_v30_
                nw75_[int(11)] = d_408_v30_
                nw75_[int(12)] = (d_408_v30_ if d_358_v1_ else d_408_v30_)
                nw75_[int(13)] = d_408_v30_
                nw75_[int(14)] = d_408_v30_
                nw75_[int(15)] = d_408_v30_
                nw75_[int(16)] = ((d_443_v44_)[(p0)[default__.safeIndex(301, (p0).length(0))]] if ((p0)[default__.safeIndex(301, (p0).length(0))]) in (d_443_v44_) else d_408_v30_)
                d_444_v45_ = nw75_
                index77_ = default__.safeIndex(135, (d_444_v45_).length(0))
                (d_444_v45_)[index77_] = d_408_v30_
                index78_ = default__.safeIndex(301, (p0).length(0))
                index79_ = default__.safeIndex(135, (d_444_v45_).length(0))
                rhs78_ = (d_440_cf8_) or (d_440_cf8_)
                rhs79_ = d_408_v30_
                lhs61_ = p0
                lhs62_ = default__.safeIndex(301, (p0).length(0))
                lhs63_ = d_444_v45_
                lhs64_ = default__.safeIndex(135, (d_444_v45_).length(0))
                lhs61_[lhs62_] = rhs78_
                lhs63_[lhs64_] = rhs79_
                (globalState).f10 = d_441_cf7_
            elif True:
                d_445___mcc_h10_ = source3_.cf0
                d_446_cf0_ = d_445___mcc_h10_
                d_447_v46_: _dafny.Map
                d_447_v46_ = _dafny.Map({self.f18: d_358_v1_})
                d_448_v48_: _dafny.Map
                def iife30_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(306, 513):
                        d_449_v47_: int = compr_16_
                        if ((306) <= (d_449_v47_)) and ((d_449_v47_) < (513)):
                            coll16_[(d_449_v47_) - (d_446_cf0_)] = d_358_v1_
                    return _dafny.Map(coll16_)
                d_448_v48_ = _dafny.Map({305: iife30_()
                })
                d_450_v49_: D0
                d_450_v49_ = D0_DC3(p2, d_358_v1_, self.f18, self.f18)
                d_451_v50_: _dafny.Map
                d_451_v50_ = _dafny.Map({(d_447_v46_) | (((d_448_v48_)[d_446_cf0_] if (d_446_cf0_) in (d_448_v48_) else d_447_v46_)): d_450_v49_})
                d_451_v50_ = (d_451_v50_).set(d_447_v46_, d_450_v49_)
                d_452_v51_: _dafny.Map
                d_452_v51_ = _dafny.Map({d_358_v1_: self.f18})
                d_453_v52_: _dafny.Seq
                d_453_v52_ = p1
                (globalState).f3 = (d_408_v30_).fm7(d_358_v1_, self.f18, (d_452_v51_ if d_358_v1_ else _dafny.Map({default__.fm1(d_358_v1_, True, d_446_cf0_, globalState): 175})), (d_453_v52_), globalState)
                (globalState).f12 = d_358_v1_
                (self).f18 = (252) - (p2)
            (globalState).f12 = False
        elif True:
            d_454_v53_: D1
            d_454_v53_ = D1_DC5(d_358_v1_, self.f18)
            source4_ = d_454_v53_
            if source4_.is_DC5:
                d_455___mcc_h11_ = source4_.cf12
                d_456___mcc_h12_ = source4_.cf13
                d_457_cf13_ = d_456___mcc_h12_
                d_458_cf12_ = d_455___mcc_h11_
                d_459_v54_: _dafny.Map
                d_459_v54_ = _dafny.Map({528: d_458_cf12_})
                d_460_v55_: _dafny.Seq
                d_460_v55_ = _dafny.SeqWithoutIsStrInference([d_458_cf12_, d_458_cf12_])
                d_461_v56_: str
                d_461_v56_ = _dafny.CodePoint('a')
                d_462_v57_: _dafny.Map
                d_462_v57_ = _dafny.Map({(((d_459_v54_)[d_457_cf13_] if (d_457_cf13_) in (d_459_v54_) else d_458_cf12_)) not in (d_460_v55_): _dafny.Map({p2: d_461_v56_})})
                d_462_v57_ = (d_462_v57_) | (d_462_v57_)
                d_463_v58_: _dafny.Array
                nw76_ = _dafny.Array(_dafny.Set({}), 21)
                d_463_v58_ = nw76_
                d_464_v59_: _dafny.Array
                d_464_v59_ = d_463_v58_
                rhs80_ = (d_464_v59_)
                rhs81_ = p2
                lhs65_ = self
                d_463_v58_ = rhs80_
                lhs65_.f18 = rhs81_
                d_465_v60_: _dafny.Map
                d_465_v60_ = _dafny.Map({self.f18: ((0) - (len(default__.fm11(d_358_v1_, globalState)))) + (p2)})
                d_465_v60_ = (d_465_v60_).set(-983, p2)
                (globalState).f10 = -458
            elif True:
                d_466___mcc_h13_ = source4_.cf11
                d_467_cf11_ = d_466___mcc_h13_
                d_468_v61_: C0
                nw77_ = C0()
                nw77_.ctor__()
                d_468_v61_ = nw77_
                d_468_v61_ = d_468_v61_
                d_358_v1_ = d_358_v1_
                d_358_v1_ = d_358_v1_
                index80_ = default__.safeIndex(317, (p0).length(0))
                (p0)[index80_] = (self.f18) >= (self.f18)
                index81_ = default__.safeIndex(317, (p0).length(0))
                (p0)[index81_] = d_358_v1_
            d_469_v63_: _dafny.Set
            d_469_v63_ = _dafny.Set({d_358_v1_})
            d_470_v64_: C0
            nw78_ = C0()
            nw78_.ctor__()
            d_470_v64_ = nw78_
            d_471_v65_: _dafny.Seq
            d_471_v65_ = _dafny.SeqWithoutIsStrInference([self.f18, 55])
            d_472_v66_: _dafny.Seq
            def iife31_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(905, 243):
                    d_473_v62_: int = compr_17_
                    if ((905) <= (d_473_v62_)) and ((d_473_v62_) < (243)):
                        coll17_[(d_473_v62_) * (p2)] = d_358_v1_
                return _dafny.Map(coll17_)
            d_472_v66_ = _dafny.SeqWithoutIsStrInference([len(iife31_()
            ), len(d_469_v63_), len(_dafny.SeqWithoutIsStrInference([d_470_v64_])), len(d_471_v65_)])
            (globalState).f12 = default__.fm1(d_358_v1_, False, len(d_472_v66_), globalState)
            (globalState).f12 = False
            d_474_v67_: _dafny.Array
            nw79_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_474_v67_ = nw79_
            d_475_v68_: _dafny.Array
            def lambda21_(d_476_i17_):
                return _dafny.MultiSet([self.f18])

            init11_ = lambda21_
            nw80_ = _dafny.Array(None, 2)
            for i0_11_ in range(nw80_.length(0)):
                nw80_[i0_11_] = init11_(i0_11_)
            d_475_v68_ = nw80_
            index82_ = default__.safeIndex(589, (d_474_v67_).length(0))
            (d_474_v67_)[index82_] = d_475_v68_
            d_477_v69_: _dafny.MultiSet
            d_477_v69_ = _dafny.MultiSet([d_358_v1_])
            d_478_v70_: _dafny.MultiSet
            d_478_v70_ = _dafny.MultiSet([self.f18, p2, (0) - (p2), (d_477_v69_).cardinality])
            d_479_v71_: _dafny.Seq
            d_479_v71_ = _dafny.SeqWithoutIsStrInference([d_478_v70_])
            index83_ = default__.safeIndex(589, (d_474_v67_).length(0))
            nw81_ = _dafny.Array(None, 15)
            nw81_[int(0)] = (d_479_v71_)[default__.safeIndex(p2, len(d_479_v71_))]
            nw81_[int(1)] = ((d_470_v64_).fm6(len(d_472_v66_), d_358_v1_, globalState)) - (d_478_v70_)
            nw81_[int(2)] = d_478_v70_
            nw81_[int(3)] = _dafny.MultiSet(((d_472_v66_).set(default__.safeIndex((0) - (p2), len(d_472_v66_)), 706)) + (d_472_v66_))
            nw81_[int(4)] = d_478_v70_
            nw81_[int(5)] = d_478_v70_
            nw81_[int(6)] = (d_478_v70_).intersection(d_478_v70_)
            nw81_[int(7)] = d_478_v70_
            nw81_[int(8)] = (d_478_v70_).set(p2, default__.abs(p2))
            nw81_[int(9)] = d_478_v70_
            nw81_[int(10)] = d_478_v70_
            nw81_[int(11)] = (d_478_v70_) - (d_478_v70_)
            nw81_[int(12)] = d_478_v70_
            nw81_[int(13)] = _dafny.MultiSet([p2, self.f18])
            nw81_[int(14)] = (d_478_v70_) | ((_dafny.MultiSet([p2, self.f18])).set(p2, default__.abs(self.f18)))
            (d_474_v67_)[index83_] = nw81_
            d_480_v72_: _dafny.Seq
            d_480_v72_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_480_v72_ = d_480_v72_
        if (d_358_v1_) or (not (d_358_v1_) or (d_358_v1_)):
            d_481_v73_: _dafny.Set
            d_481_v73_ = _dafny.Set({True})
            d_482_v74_: str
            d_482_v74_ = _dafny.CodePoint('i')
            d_483_v75_: _dafny.Seq
            d_483_v75_ = _dafny.SeqWithoutIsStrInference([d_481_v73_, default__.fm12(d_482_v74_, self.f18, globalState)])
            d_484_v76_: _dafny.Map
            d_484_v76_ = _dafny.Map({len((d_483_v75_)[default__.safeIndex(p2, len(d_483_v75_))]): d_357_v0_})
            d_485_v77_: D6
            d_485_v77_ = D6_DC13(d_484_v76_)
            source5_ = D0_DC3(p2, default__.fm1(d_358_v1_, d_358_v1_, len((d_485_v77_).cf27), globalState), p2, p2)
            if source5_.is_DC1:
                d_486___mcc_h14_ = source5_.cf1
                d_487___mcc_h15_ = source5_.cf2
                d_488___mcc_h16_ = source5_.cf3
                d_489___mcc_h17_ = source5_.cf4
                d_490___mcc_h18_ = source5_.cf5
                d_491_cf5_ = d_490___mcc_h18_
                d_492_cf4_ = d_489___mcc_h17_
                d_493_cf3_ = d_488___mcc_h16_
                d_494_cf2_ = d_487___mcc_h15_
                d_495_cf1_ = d_486___mcc_h14_
                index84_ = default__.safeIndex(463, (d_357_v0_).length(0))
                (d_357_v0_)[index84_] = p2
                d_496_v78_: _dafny.MultiSet
                d_496_v78_ = _dafny.MultiSet([d_358_v1_, d_358_v1_])
                d_497_v79_: _dafny.Seq
                d_497_v79_ = _dafny.SeqWithoutIsStrInference([d_358_v1_, d_358_v1_, d_358_v1_, d_358_v1_, d_358_v1_])
                index85_ = default__.safeIndex(463, (d_357_v0_).length(0))
                (d_357_v0_)[index85_] = ((d_494_cf2_) - (((d_496_v78_)[d_358_v1_] if (d_358_v1_) in (d_496_v78_) else p2))) * (len(default__.fm13(d_358_v1_, d_497_v79_, globalState)))
                d_498_v80_: C0
                nw82_ = C0()
                nw82_.ctor__()
                d_498_v80_ = nw82_
                d_499_v81_: C0
                nw83_ = C0()
                nw83_.ctor__()
                d_499_v81_ = nw83_
                d_500_v82_: _dafny.Map
                d_500_v82_ = _dafny.Map({d_498_v80_: d_494_cf2_})
                d_501_v83_: D3
                d_501_v83_ = D3_DC7(d_500_v82_)
                pat_let_tv5_ = d_500_v82_
                def iife32_(_pat_let7_0):
                    def iife33_(d_502_dt__update__tmp_h1_):
                        def iife34_(_pat_let8_0):
                            def iife35_(d_503_dt__update_hcf15_h0_):
                                return D3_DC7(d_503_dt__update_hcf15_h0_)
                            return iife35_(_pat_let8_0)
                        return iife34_(pat_let_tv5_)
                    return iife33_(_pat_let7_0)
                d_501_v83_ = iife32_((d_501_v83_ if d_358_v1_ else d_501_v83_))
            elif source5_.is_DC2:
                d_504___mcc_h19_ = source5_.cf6
                d_505_cf6_ = d_504___mcc_h19_
                d_506_v84_: C0
                nw84_ = C0()
                nw84_.ctor__()
                d_506_v84_ = nw84_
                d_507_v85_: _dafny.Map
                d_507_v85_ = _dafny.Map({not(d_505_cf6_): d_506_v84_})
                d_507_v85_ = (d_507_v85_).set(d_358_v1_, d_506_v84_)
                index86_ = default__.safeIndex(849, (p0).length(0))
                (p0)[index86_] = d_505_cf6_
                d_508_v86_: _dafny.Map
                d_508_v86_ = _dafny.Map({d_505_cf6_: d_358_v1_})
                d_509_v87_: _dafny.Seq
                d_509_v87_ = _dafny.SeqWithoutIsStrInference([p2, len(d_508_v86_)])
                d_510_v88_: _dafny.MultiSet
                d_510_v88_ = _dafny.MultiSet([p2])
                index87_ = default__.safeIndex(849, (p0).length(0))
                (p0)[index87_] = ((_dafny.MultiSet(d_509_v87_)).intersection(d_510_v88_)).ispropersubset(d_510_v88_)
                d_511_v89_: _dafny.Map
                d_511_v89_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p2 for d_512_i18_ in range(default__.abs(79))]): self.f18})
                d_511_v89_ = (d_511_v89_).set(d_509_v87_, (p2) - (p2))
                d_513_v90_: _dafny.Array
                nw85_ = _dafny.Array(None, 18)
                nw85_[int(0)] = (p0)[default__.safeIndex(849, (p0).length(0))]
                nw85_[int(1)] = not(d_358_v1_)
                nw85_[int(2)] = (p0)[default__.safeIndex(849, (p0).length(0))]
                nw85_[int(3)] = False
                nw85_[int(4)] = (p0)[default__.safeIndex(849, (p0).length(0))]
                nw85_[int(5)] = False
                nw85_[int(6)] = False
                nw85_[int(7)] = d_505_cf6_
                nw85_[int(8)] = d_505_cf6_
                nw85_[int(9)] = (p0)[default__.safeIndex(849, (p0).length(0))]
                nw85_[int(10)] = (p0)[default__.safeIndex(849, (p0).length(0))]
                nw85_[int(11)] = d_505_cf6_
                nw85_[int(12)] = (p0)[default__.safeIndex(849, (p0).length(0))]
                nw85_[int(13)] = d_505_cf6_
                nw85_[int(14)] = d_358_v1_
                nw85_[int(15)] = d_505_cf6_
                nw85_[int(16)] = d_505_cf6_
                nw85_[int(17)] = True
                d_513_v90_ = nw85_
                d_514_v91_: _dafny.MultiSet
                d_514_v91_ = _dafny.MultiSet([d_513_v90_])
                (self).f18 = len((d_509_v87_) + (_dafny.SeqWithoutIsStrInference([p2, (d_514_v91_).cardinality])))
            elif source5_.is_DC3:
                d_515___mcc_h20_ = source5_.cf7
                d_516___mcc_h21_ = source5_.cf8
                d_517___mcc_h22_ = source5_.cf9
                d_518___mcc_h23_ = source5_.cf10
                d_519_cf10_ = d_518___mcc_h23_
                d_520_cf9_ = d_517___mcc_h22_
                d_521_cf8_ = d_516___mcc_h21_
                d_522_cf7_ = d_515___mcc_h20_
                d_523_v92_: _dafny.Seq
                d_523_v92_ = _dafny.SeqWithoutIsStrInference([d_358_v1_, d_358_v1_])
                d_524_v93_: _dafny.MultiSet
                d_524_v93_ = _dafny.MultiSet([default__.fm1(d_358_v1_, d_521_cf8_, self.f18, globalState), d_358_v1_, d_358_v1_, d_358_v1_, d_358_v1_])
                d_525_v94_: _dafny.Seq
                d_525_v94_ = _dafny.SeqWithoutIsStrInference([(p2) + (len(d_523_v92_)), ((d_524_v93_)[d_521_cf8_] if (d_521_cf8_) in (d_524_v93_) else p2), d_519_cf10_])
                d_525_v94_ = (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_522_cf7_})) for d_526_i19_ in range(default__.abs(429))]) if d_358_v1_ else d_525_v94_)
                d_527_v95_: _dafny.MultiSet
                d_527_v95_ = _dafny.MultiSet([d_482_v74_])
                d_528_v96_: _dafny.Map
                d_528_v96_ = _dafny.Map({default__.safeDivisionInt(default__.fm0(globalState), p2): d_527_v95_})
                d_528_v96_ = (d_528_v96_).set(d_520_cf9_, d_527_v95_)
                nw86_ = _dafny.Array(int(0), 12)
                d_357_v0_ = nw86_
                (globalState).f10 = self.f18
            elif True:
                d_529___mcc_h24_ = source5_.cf0
                d_530_cf0_ = d_529___mcc_h24_
                d_531_v97_: _dafny.Map
                d_531_v97_ = _dafny.Map({d_358_v1_: d_358_v1_})
                rhs82_ = default__.fm0(globalState)
                rhs83_ = (d_530_cf0_) >= (len((d_531_v97_) | (d_531_v97_)))
                rhs84_ = d_357_v0_
                lhs66_ = globalState
                lhs67_ = globalState
                lhs66_.f11 = rhs82_
                lhs67_.f12 = rhs83_
                d_357_v0_ = rhs84_
                d_532_v98_: _dafny.MultiSet
                d_532_v98_ = _dafny.MultiSet([d_358_v1_])
                d_532_v98_ = ((d_532_v98_) - (d_532_v98_)) | (default__.fm14(self.f18, D1_DC5(d_358_v1_, d_530_cf0_), globalState))
                (globalState).f12 = d_358_v1_
                d_533_v99_: C0
                nw87_ = C0()
                nw87_.ctor__()
                d_533_v99_ = nw87_
            d_357_v0_ = d_357_v0_
            d_534_v100_: _dafny.Set
            d_534_v100_ = _dafny.Set({p2})
            d_534_v100_ = d_534_v100_
            d_358_v1_ = (p1) == (p1)
            d_535_v101_: _dafny.Map
            d_535_v101_ = _dafny.Map({d_358_v1_: self.f18})
            d_536_v102_: D0
            d_536_v102_ = D0_DC3(429, d_358_v1_, len(d_535_v101_), 366)
            d_537_v103_: _dafny.Map
            d_537_v103_ = _dafny.Map({p1: (d_536_v102_).cf8})
            d_537_v103_ = d_537_v103_
        elif True:
            (globalState).f10 = self.f18
            d_538_v104_: _dafny.Array
            def lambda22_(d_539_i20_):
                return _dafny.CodePoint('t')

            init12_ = lambda22_
            nw88_ = _dafny.Array(None, 11)
            for i0_12_ in range(nw88_.length(0)):
                nw88_[i0_12_] = init12_(i0_12_)
            d_538_v104_ = nw88_
            d_538_v104_ = d_538_v104_
            d_540_v105_: str
            d_540_v105_ = _dafny.CodePoint('l')
            d_541_v106_: _dafny.Set
            d_541_v106_ = _dafny.Set({d_540_v105_, default__.fm9(d_358_v1_, not(False), globalState), d_540_v105_, d_540_v105_})
            d_542_v107_: D0
            d_542_v107_ = D0_DC2((p2) >= (len(d_541_v106_)))
            source6_ = d_542_v107_
            if source6_.is_DC1:
                d_543___mcc_h25_ = source6_.cf1
                d_544___mcc_h26_ = source6_.cf2
                d_545___mcc_h27_ = source6_.cf3
                d_546___mcc_h28_ = source6_.cf4
                d_547___mcc_h29_ = source6_.cf5
                d_548_cf5_ = d_547___mcc_h29_
                d_549_cf4_ = d_546___mcc_h28_
                d_550_cf3_ = d_545___mcc_h27_
                d_551_cf2_ = d_544___mcc_h26_
                d_552_cf1_ = d_543___mcc_h25_
                index88_ = default__.safeIndex(173, (d_548_cf5_).length(0))
                (d_548_cf5_)[index88_] = (default__.fm1(d_358_v1_, d_358_v1_, default__.fm0(globalState), globalState)) and (d_358_v1_)
                d_553_v108_: _dafny.Seq
                d_553_v108_ = _dafny.SeqWithoutIsStrInference([d_552_cf1_])
                d_554_v109_: _dafny.Map
                d_554_v109_ = _dafny.Map({d_357_v0_: d_358_v1_})
                d_555_v110_: _dafny.MultiSet
                d_555_v110_ = _dafny.MultiSet([786])
                index89_ = default__.safeIndex(173, (d_548_cf5_).length(0))
                (d_548_cf5_)[index89_] = ((_dafny.MultiSet(d_553_v108_)).set(d_550_cf3_, default__.abs(len(d_554_v109_)))).isdisjoint(d_555_v110_)
                d_556_v111_: _dafny.Seq
                d_556_v111_ = _dafny.SeqWithoutIsStrInference([(d_548_cf5_)[default__.safeIndex(173, (d_548_cf5_).length(0))]])
                d_556_v111_ = d_556_v111_
                d_557_v112_: C0
                nw89_ = C0()
                nw89_.ctor__()
                d_557_v112_ = nw89_
                d_358_v1_ = (d_556_v111_)[default__.safeIndex(default__.fm0(globalState), len(d_556_v111_))]
            elif source6_.is_DC2:
                d_558___mcc_h30_ = source6_.cf6
                d_559_cf6_ = d_558___mcc_h30_
                d_560_v113_: D6
                d_560_v113_ = D6_DC14(False, d_540_v105_, d_538_v104_)
                d_560_v113_ = d_560_v113_
                d_561_v114_: _dafny.Array
                nw90_ = _dafny.Array(_dafny.Seq({}), 18)
                d_561_v114_ = nw90_
                d_562_v115_: D4
                d_562_v115_ = D4_DC10(d_540_v105_, d_561_v114_, self.f18)
                d_563_v116_: _dafny.Set
                d_563_v116_ = _dafny.Set({d_562_v115_})
                d_564_v117_: _dafny.Map
                d_564_v117_ = _dafny.Map({(_dafny.Set({d_562_v115_})).issubset(d_563_v116_): d_358_v1_})
                d_565_v118_: _dafny.Array
                def lambda23_(d_566_v107_):
                    def lambda24_(d_567_i21_):
                        return d_566_v107_

                    return lambda24_

                init13_ = lambda23_(d_542_v107_)
                nw91_ = _dafny.Array(None, 7)
                for i0_13_ in range(nw91_.length(0)):
                    nw91_[i0_13_] = init13_(i0_13_)
                d_565_v118_ = nw91_
                index90_ = default__.safeIndex(94, (d_565_v118_).length(0))
                (d_565_v118_)[index90_] = d_542_v107_
                index91_ = default__.safeIndex(94, (d_565_v118_).length(0))
                rhs85_ = d_564_v117_
                rhs86_ = _dafny.SeqWithoutIsStrInference([d_540_v105_])
                rhs87_ = d_542_v107_
                rhs88_ = d_559_cf6_
                lhs68_ = globalState
                lhs69_ = d_565_v118_
                lhs70_ = default__.safeIndex(94, (d_565_v118_).length(0))
                lhs71_ = globalState
                d_564_v117_ = rhs85_
                lhs68_.f3 = rhs86_
                lhs69_[lhs70_] = rhs87_
                lhs71_.f12 = rhs88_
                d_568_v119_: C0
                nw92_ = C0()
                nw92_.ctor__()
                d_568_v119_ = nw92_
                d_569_v120_: C0
                nw93_ = C0()
                nw93_.ctor__()
                d_569_v120_ = nw93_
            elif source6_.is_DC3:
                d_570___mcc_h31_ = source6_.cf7
                d_571___mcc_h32_ = source6_.cf8
                d_572___mcc_h33_ = source6_.cf9
                d_573___mcc_h34_ = source6_.cf10
                d_574_cf10_ = d_573___mcc_h34_
                d_575_cf9_ = d_572___mcc_h33_
                d_576_cf8_ = d_571___mcc_h32_
                d_577_cf7_ = d_570___mcc_h31_
                d_578_v121_: C0
                nw94_ = C0()
                nw94_.ctor__()
                d_578_v121_ = nw94_
                d_579_v122_: _dafny.Map
                d_579_v122_ = _dafny.Map({d_540_v105_: d_577_cf7_})
                (globalState).f11 = len((d_579_v122_) | (d_579_v122_))
                d_580_v123_: C0
                nw95_ = C0()
                nw95_.ctor__()
                d_580_v123_ = nw95_
                d_581_v124_: _dafny.Array
                nw96_ = _dafny.Array(None, 28)
                d_581_v124_ = nw96_
                index92_ = default__.safeIndex(371, (d_581_v124_).length(0))
                (d_581_v124_)[index92_] = d_578_v121_
                index93_ = default__.safeIndex(371, (d_581_v124_).length(0))
                (d_581_v124_)[index93_] = d_580_v123_
            elif True:
                d_582___mcc_h35_ = source6_.cf0
                d_583_cf0_ = d_582___mcc_h35_
                d_584_v125_: _dafny.Set
                d_584_v125_ = _dafny.Set({self.f18})
                (globalState).f12 = (_dafny.Set({d_583_cf0_, d_583_cf0_})).issubset(d_584_v125_)
                d_585_v126_: _dafny.Map
                d_585_v126_ = _dafny.Map({self.f18: p2})
                d_586_v127_: _dafny.Seq
                d_586_v127_ = _dafny.SeqWithoutIsStrInference([d_358_v1_])
                d_587_v131_: C0
                nw97_ = C0()
                nw97_.ctor__()
                d_587_v131_ = nw97_
                d_588_v132_: _dafny.Set
                d_588_v132_ = _dafny.Set({d_587_v131_})
                d_589_v133_: _dafny.Array
                nw98_ = _dafny.Array(None, 21)
                nw98_[int(0)] = (d_585_v126_) | (d_585_v126_)
                nw98_[int(1)] = _dafny.Map({self.f18: self.f18})
                nw98_[int(2)] = d_585_v126_
                nw98_[int(3)] = (d_585_v126_ if (d_586_v127_)[default__.safeIndex(self.f18, len(d_586_v127_))] else d_585_v126_)
                nw98_[int(4)] = d_585_v126_
                nw98_[int(5)] = d_585_v126_
                def iife36_():
                    coll18_ = _dafny.Map()
                    compr_18_: int
                    for compr_18_ in _dafny.IntegerRange(241, 378):
                        d_590_v128_: int = compr_18_
                        if ((241) <= (d_590_v128_)) and ((d_590_v128_) < (378)):
                            coll18_[default__.safeModuloInt(d_590_v128_, (0) - (len(d_586_v127_)))] = self.f18
                    return _dafny.Map(coll18_)
                nw98_[int(6)] = iife36_()
                
                nw98_[int(7)] = d_585_v126_
                def iife37_():
                    coll19_ = _dafny.Map()
                    compr_19_: int
                    for compr_19_ in (d_585_v126_).keys.Elements:
                        d_591_v129_: int = compr_19_
                        if (d_591_v129_) in (d_585_v126_):
                            coll19_[(d_591_v129_) + ((0) - (p2))] = self.f18
                    return _dafny.Map(coll19_)
                def iife38_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(735, 507):
                        d_592_v130_: int = compr_20_
                        if ((735) <= (d_592_v130_)) and ((d_592_v130_) < (507)):
                            coll20_[default__.safeDivisionInt(d_592_v130_, d_583_cf0_)] = (_dafny.MultiSet([d_583_cf0_])).cardinality
                    return _dafny.Map(coll20_)
                nw98_[int(8)] = (iife37_()
                ) | (iife38_()
                )
                nw98_[int(9)] = default__.fm15(len(d_588_v132_), d_358_v1_, d_583_cf0_, globalState)
                nw98_[int(10)] = d_585_v126_
                nw98_[int(11)] = d_585_v126_
                nw98_[int(12)] = _dafny.Map({d_583_cf0_: 650})
                nw98_[int(13)] = d_585_v126_
                nw98_[int(14)] = (d_585_v126_) | (d_585_v126_)
                nw98_[int(15)] = d_585_v126_
                nw98_[int(16)] = _dafny.Map({self.f18: d_583_cf0_})
                nw98_[int(17)] = d_585_v126_
                nw98_[int(18)] = d_585_v126_
                nw98_[int(19)] = d_585_v126_
                nw98_[int(20)] = d_585_v126_
                d_589_v133_ = nw98_
                nw99_ = _dafny.Array(_dafny.Map({}), 20)
                d_589_v133_ = nw99_
                d_593_v134_: _dafny.Array
                nw100_ = _dafny.Array(D4.default()(), 14)
                d_593_v134_ = nw100_
                d_594_v135_: _dafny.Array
                nw101_ = _dafny.Array(_dafny.Seq({}), 26)
                d_594_v135_ = nw101_
                d_595_v136_: D4
                d_595_v136_ = D4_DC10(d_540_v105_, d_594_v135_, (0) - (self.f18))
                index94_ = default__.safeIndex(875, (d_593_v134_).length(0))
                (d_593_v134_)[index94_] = d_595_v136_
                index95_ = default__.safeIndex(875, (d_593_v134_).length(0))
                (d_593_v134_)[index95_] = d_595_v136_
                d_596_v137_: _dafny.MultiSet
                d_596_v137_ = _dafny.MultiSet([p1])
                d_585_v126_ = (d_585_v126_).set((d_596_v137_).cardinality, p2)
            d_597_v138_: _dafny.Map
            d_597_v138_ = _dafny.Map({d_358_v1_: d_358_v1_})
            d_598_v139_: _dafny.Array
            def lambda25_(d_599_v1_):
                def lambda26_(d_600_i22_):
                    return _dafny.Set({d_599_v1_})

                return lambda26_

            init14_ = lambda25_(d_358_v1_)
            nw102_ = _dafny.Array(None, 9)
            for i0_14_ in range(nw102_.length(0)):
                nw102_[i0_14_] = init14_(i0_14_)
            d_598_v139_ = nw102_
            d_601_v140_: _dafny.Array
            d_601_v140_ = d_598_v139_
            source7_ = (d_598_v139_ if ((d_597_v138_)[d_358_v1_] if (d_358_v1_) in (d_597_v138_) else d_358_v1_) else d_601_v140_)
            d_602___mcc_h36_ = source7_
            d_603_cf26_ = d_602___mcc_h36_
            index96_ = default__.safeIndex(125, (p0).length(0))
            (p0)[index96_] = (self.f18) != (p2)
            index97_ = default__.safeIndex(125, (p0).length(0))
            (p0)[index97_] = d_358_v1_
            d_604_v141_: C0
            nw103_ = C0()
            nw103_.ctor__()
            d_604_v141_ = nw103_
            (globalState).f11 = self.f18
            d_605_v142_: _dafny.Seq
            d_605_v142_ = _dafny.SeqWithoutIsStrInference([d_358_v1_, d_358_v1_, (p0)[default__.safeIndex(125, (p0).length(0))]])
            d_606_v143_: _dafny.Seq
            d_606_v143_ = _dafny.SeqWithoutIsStrInference([d_538_v104_])
            d_607_v144_: D6
            d_607_v144_ = D6_DC14(False, d_540_v105_, (d_606_v143_)[default__.safeIndex(p2, len(d_606_v143_))])
            d_608_v145_: _dafny.Map
            d_608_v145_ = _dafny.Map({(d_605_v142_) + (d_605_v142_): (True if (p0)[default__.safeIndex(125, (p0).length(0))] else (d_607_v144_).cf28)})
            d_608_v145_ = (d_608_v145_).set(_dafny.SeqWithoutIsStrInference([d_358_v1_]), not(d_358_v1_))
            if not (d_358_v1_) or (not((d_358_v1_ if d_358_v1_ else d_358_v1_))):
                d_609_v146_: C0
                nw104_ = C0()
                nw104_.ctor__()
                d_609_v146_ = nw104_
                d_610_v147_: _dafny.Array
                out12_: _dafny.Array
                out12_ = default__.m0(p2, globalState)
                d_610_v147_ = out12_
                index98_ = default__.safeIndex(146, (d_357_v0_).length(0))
                (d_357_v0_)[index98_] = (self.f18 if default__.fm1(not(d_358_v1_), True, self.f18, globalState) else self.f18)
                index99_ = default__.safeIndex(146, (d_357_v0_).length(0))
                (d_357_v0_)[index99_] = (((0) - (self.f18)) - (self.f18)) * ((0) - ((p2) * (self.f18)))
                d_611_v148_: _dafny.Map
                d_611_v148_ = _dafny.Map({d_358_v1_: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cudbn"))))})
                (globalState).f3 = (d_609_v146_).fm7(d_358_v1_, len(_dafny.SeqWithoutIsStrInference([d_540_v105_ for d_612_i23_ in range(default__.abs(554))])), d_611_v148_, p1, globalState)
                rhs89_ = d_540_v105_
                rhs90_ = self.f18
                lhs72_ = globalState
                d_540_v105_ = rhs89_
                lhs72_.f10 = rhs90_
            elif True:
                d_613_v149_: C0
                nw105_ = C0()
                nw105_.ctor__()
                d_613_v149_ = nw105_
                d_614_v151_: _dafny.Array
                def lambda27_(d_615_p2_):
                    def lambda28_(d_616_i24_):
                        def iife39_():
                            coll21_ = _dafny.Map()
                            compr_21_: int
                            for compr_21_ in _dafny.IntegerRange(640, -60):
                                d_617_v150_: int = compr_21_
                                if ((640) <= (d_617_v150_)) and ((d_617_v150_) < (-60)):
                                    coll21_[(d_617_v150_) * (d_615_p2_)] = self.f18
                            return _dafny.Map(coll21_)
                        return _dafny.SeqWithoutIsStrInference([self.f18, len(iife39_()
                        )])

                    return lambda28_

                init15_ = lambda27_(p2)
                nw106_ = _dafny.Array(None, 4)
                for i0_15_ in range(nw106_.length(0)):
                    nw106_[i0_15_] = init15_(i0_15_)
                d_614_v151_ = nw106_
                d_618_v152_: D0
                d_618_v152_ = D0_DC3(self.f18, False, self.f18, len(d_541_v106_))
                d_619_v153_: D4
                d_619_v153_ = D4_DC10(default__.fm9(d_358_v1_, d_358_v1_, globalState), d_614_v151_, len(_dafny.Set({d_618_v152_})))
                d_619_v153_ = d_619_v153_
                d_620_v154_: _dafny.Seq
                d_620_v154_ = _dafny.SeqWithoutIsStrInference([d_357_v0_, d_357_v0_])
                d_358_v1_ = not(default__.fm1(d_358_v1_, (d_357_v0_) in (d_620_v154_), (p2) + (p2), globalState))
                (globalState).f10 = (p2 if d_358_v1_ else (self.f18) - (p2))
                (globalState).f10 = self.f18
        d_621_v155_: _dafny.Map
        d_621_v155_ = _dafny.Map({default__.fm0(globalState): not(d_358_v1_)})
        d_622_v156_: _dafny.Map
        d_622_v156_ = _dafny.Map({len((d_621_v155_ if d_358_v1_ else d_621_v155_)): p2})
        d_622_v156_ = (d_622_v156_).set((0) - (self.f18), (0) - (len(p1)))

