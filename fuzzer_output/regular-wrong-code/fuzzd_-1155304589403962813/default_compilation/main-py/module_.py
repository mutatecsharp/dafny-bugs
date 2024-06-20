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
    def fm0(p0, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, False])) if False else _dafny.MultiSet([True]))).ispropersubset((_dafny.MultiSet([True, False, not(True)]) if False else _dafny.MultiSet([not(not(True)), False])))

    @staticmethod
    def fm1(p0, globalState):
        return default__.safeDivisionInt((0) - (-606), 0)

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm3(globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rm")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlxxw"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_0_i0_ in range(default__.abs(992))])))

    @staticmethod
    def fm4(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sh")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "siqr")) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovpwyj"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amlqk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "reultjxs"))])

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([836]))).Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([836]))):
                    coll0_ = coll0_.union(_dafny.Set([(d_1_v0_) * (248)]))
            return _dafny.Set(coll0_)
        return (_dafny.MultiSet([len(_dafny.Map({119: True})), 921, len(_dafny.Map({len(iife0_()
        ): False}))])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: 839})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2_i0_ in range(default__.abs(42))])), len(_dafny.Set({False, False, not(False), False}))])))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iffq")): (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3_i0_ in range(default__.abs(609))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxl")))})

    @staticmethod
    def fm7(globalState):
        source0_ = D5_DC10(_dafny.MultiSet([False]))
        if source0_.is_DC11:
            d_4___mcc_h0_ = source0_.cf17
            d_5_cf17_ = d_4___mcc_h0_
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(-39, 780):
                    d_6_v0_: int = compr_1_
                    if ((-39) <= (d_6_v0_)) and ((d_6_v0_) < (780)):
                        coll1_[default__.safeModuloInt(d_6_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nshammh"))))] = True
                return _dafny.Map(coll1_)
            return _dafny.Map({len(_dafny.Map({True: len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csjiecvhl")): len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(True): len(_dafny.Map({_dafny.CodePoint('a'): len(iife1_()
            )}))}))]))}))})): D4_DC8(_dafny.Map({True: 564}))})
        elif source0_.is_DC12:
            d_7___mcc_h1_ = source0_.cf18
            d_8___mcc_h2_ = source0_.cf19
            d_9___mcc_h3_ = source0_.cf20
            d_10___mcc_h4_ = source0_.cf21
            d_11___mcc_h5_ = source0_.cf22
            d_12_cf22_ = d_11___mcc_h5_
            d_13_cf21_ = d_10___mcc_h4_
            d_14_cf20_ = d_9___mcc_h3_
            d_15_cf19_ = d_8___mcc_h2_
            d_16_cf18_ = d_7___mcc_h1_
            return _dafny.Map({d_12_cf22_: D4_DC8(_dafny.Map({d_14_cf20_: d_16_cf18_}))})
        elif True:
            d_17___mcc_h6_ = source0_.cf16
            d_18_cf16_ = d_17___mcc_h6_
            return (_dafny.Map({-694: D4_DC8(_dafny.Map({False: len(_dafny.Set({False}))}))})) | (_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality: D4_DC8(_dafny.Map({False: -749}))}))

    @staticmethod
    def fm8(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - (-863) for d_19_i0_ in range(default__.abs(755))]) if True else _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhedjju"))) for d_20_i1_ in range(default__.abs(-749))]))) + ((_dafny.SeqWithoutIsStrInference([227]) if not(True) else _dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, False, False, not(False)})), 215])))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return _dafny.Map({len((_dafny.SeqWithoutIsStrInference([221])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({True: True})): _dafny.MultiSet([True])})) for d_21_i0_ in range(default__.abs(6))]))): not((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwneg")))) >= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjycrdgy")))))})

    @staticmethod
    def fm10(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), (_dafny.SeqWithoutIsStrInference([not(True), not(True)])) + (_dafny.SeqWithoutIsStrInference([not(False)]))])

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([not(False)])))

    @staticmethod
    def fm12(globalState):
        if False:
            if True:
                return _dafny.CodePoint('k')
            elif True:
                return _dafny.CodePoint('t')
        elif False:
            return _dafny.CodePoint('l')
        elif True:
            return _dafny.CodePoint('h')

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        return (D7_DC14(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsiewudbc")): _dafny.Set({len(_dafny.SeqWithoutIsStrInference([926, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_22_i0_ in range(default__.abs(-434))])), -827, 851, 812])), 380})}))).cf24

    @staticmethod
    def m0(p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_23_v0_: bool
        d_23_v0_ = True
        d_24_v1_: _dafny.Seq
        d_24_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtfyrctyq"))
        d_25_v2_: _dafny.Seq
        d_25_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plyeo"))
        source1_ = (d_24_v1_ if d_23_v0_ else d_25_v2_)
        d_26___mcc_h0_ = source1_
        d_27_cf4_ = d_26___mcc_h0_
        d_28_v3_: str
        d_28_v3_ = _dafny.CodePoint('m')
        d_29_v4_: _dafny.MultiSet
        d_29_v4_ = _dafny.MultiSet([d_28_v3_, d_28_v3_, d_28_v3_])
        d_30_v5_: _dafny.Map
        d_30_v5_ = _dafny.Map({(d_29_v4_).cardinality: d_23_v0_})
        d_31_v6_: int
        d_31_v6_ = 847
        d_30_v5_ = (d_30_v5_).set(d_31_v6_, (d_31_v6_) >= (-326))
        if d_23_v0_:
            d_32_v7_: _dafny.Seq
            d_32_v7_ = _dafny.SeqWithoutIsStrInference([d_27_cf4_, _dafny.SeqWithoutIsStrInference([d_28_v3_ for d_33_i0_ in range(default__.abs(180))])])
            d_34_v8_: _dafny.Array
            nw0_ = _dafny.Array(None, 6)
            nw0_[int(0)] = (d_32_v7_) + (d_32_v7_)
            nw0_[int(1)] = d_32_v7_
            nw0_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsajg"))])
            nw0_[int(3)] = (d_32_v7_) + (d_32_v7_)
            nw0_[int(4)] = (d_32_v7_) + (_dafny.SeqWithoutIsStrInference([d_24_v1_, d_27_cf4_]))
            nw0_[int(5)] = (d_32_v7_) + (d_32_v7_)
            d_34_v8_ = nw0_
            index0_ = default__.safeIndex(9, (d_34_v8_).length(0))
            (d_34_v8_)[index0_] = _dafny.SeqWithoutIsStrInference([d_27_cf4_ for d_35_i1_ in range(default__.abs(-648))])
            index1_ = default__.safeIndex(9, (d_34_v8_).length(0))
            (d_34_v8_)[index1_] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_28_v3_ for d_36_i2_ in range(default__.abs(822))])])
            d_37_v9_: _dafny.Array
            def lambda0_(d_38_v6_):
                def lambda1_(d_39_i3_):
                    return (_dafny.SeqWithoutIsStrInference([d_38_v6_, d_38_v6_])) + (_dafny.SeqWithoutIsStrInference([d_38_v6_]))

                return lambda1_

            init0_ = lambda0_(d_31_v6_)
            nw1_ = _dafny.Array(None, 18)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_37_v9_ = nw1_
            index2_ = default__.safeIndex(573, (d_37_v9_).length(0))
            (d_37_v9_)[index2_] = default__.fm8(True, d_29_v4_, globalState)
            d_40_v10_: _dafny.Seq
            d_40_v10_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_31_v6_))])
            d_41_v11_: _dafny.Seq
            d_41_v11_ = _dafny.SeqWithoutIsStrInference([(d_40_v10_) + (default__.fm8(d_23_v0_, d_29_v4_, globalState))])
            index3_ = default__.safeIndex(573, (d_37_v9_).length(0))
            (d_37_v9_)[index3_] = (d_41_v11_)[default__.safeIndex(default__.safeModuloInt(d_31_v6_, d_31_v6_), len(d_41_v11_))]
            d_42_v12_: _dafny.Map
            d_42_v12_ = _dafny.Map({842: d_30_v5_})
            d_43_v13_: _dafny.Seq
            d_43_v13_ = _dafny.SeqWithoutIsStrInference([((d_42_v12_)[(0) - (d_31_v6_)] if ((0) - (d_31_v6_)) in (d_42_v12_) else _dafny.Map({d_31_v6_: d_23_v0_})), d_30_v5_])
            d_43_v13_ = (_dafny.SeqWithoutIsStrInference([(d_30_v5_).set(d_31_v6_, d_23_v0_), default__.fm9(d_23_v0_, d_23_v0_, d_23_v0_, d_23_v0_, globalState), d_30_v5_, d_30_v5_])) + (d_43_v13_)
            d_44_v14_: _dafny.Array
            def lambda2_(d_45_v0_):
                def lambda3_(d_46_i4_):
                    return d_45_v0_

                return lambda3_

            init1_ = lambda2_(d_23_v0_)
            nw2_ = _dafny.Array(None, 21)
            for i0_1_ in range(nw2_.length(0)):
                nw2_[i0_1_] = init1_(i0_1_)
            d_44_v14_ = nw2_
            d_47_v15_: _dafny.Map
            d_47_v15_ = _dafny.Map({d_23_v0_: d_44_v14_})
            d_48_v16_: _dafny.Map
            d_48_v16_ = _dafny.Map({(d_31_v6_) - (d_31_v6_): d_31_v6_})
            d_49_v17_: _dafny.Array
            def lambda4_(d_50_v6_):
                def lambda5_(d_51_i5_):
                    return (d_51_i5_) * (d_50_v6_)

                return lambda5_

            init2_ = lambda4_(d_31_v6_)
            nw3_ = _dafny.Array(None, 9)
            for i0_2_ in range(nw3_.length(0)):
                nw3_[i0_2_] = init2_(i0_2_)
            d_49_v17_ = nw3_
            rhs0_ = (d_47_v15_) | ((d_47_v15_) | (d_47_v15_))
            rhs1_ = (d_31_v6_) + (d_31_v6_)
            rhs2_ = (_dafny.Map({d_31_v6_: d_31_v6_})) | (d_48_v16_)
            rhs3_ = (d_31_v6_) != (d_31_v6_)
            rhs4_ = d_49_v17_
            lhs0_ = globalState
            d_47_v15_ = rhs0_
            d_31_v6_ = rhs1_
            d_48_v16_ = rhs2_
            lhs0_.f1 = rhs3_
            d_49_v17_ = rhs4_
            d_52_v18_: _dafny.Seq
            d_52_v18_ = _dafny.SeqWithoutIsStrInference([d_23_v0_])
            d_53_v19_: _dafny.MultiSet
            d_53_v19_ = _dafny.MultiSet([(0) - (len(d_52_v18_))])
            d_54_v20_: _dafny.Map
            d_54_v20_ = _dafny.Map({d_23_v0_: d_40_v10_})
            index4_ = default__.safeIndex(397, (d_44_v14_).length(0))
            def iife2_():
                coll2_ = _dafny.Set()
                compr_2_: str
                for compr_2_ in (_dafny.Map({d_28_v3_: len(d_54_v20_)})).keys.Elements:
                    d_55_v21_: str = compr_2_
                    if (d_55_v21_) in (_dafny.Map({d_28_v3_: len(d_54_v20_)})):
                        coll2_ = coll2_.union(_dafny.Set([d_55_v21_]))
                return _dafny.Set(coll2_)
            def iife3_():
                coll3_ = _dafny.Set()
                compr_3_: str
                for compr_3_ in (_dafny.Map({d_28_v3_: len(d_54_v20_)})).keys.Elements:
                    d_56_v21_: str = compr_3_
                    if (d_56_v21_) in (_dafny.Map({d_28_v3_: len(d_54_v20_)})):
                        coll3_ = coll3_.union(_dafny.Set([d_56_v21_]))
                return _dafny.Set(coll3_)
            (d_44_v14_)[index4_] = default__.fm0(((d_53_v19_)[len(iife2_()
            )] if (len(iife3_()
            )) in (d_53_v19_) else 589), globalState)
            d_57_v22_: _dafny.Map
            d_57_v22_ = _dafny.Map({d_28_v3_: d_49_v17_})
            index5_ = default__.safeIndex(397, (d_44_v14_).length(0))
            (d_44_v14_)[index5_] = ((d_57_v22_ if d_23_v0_ else _dafny.Map({d_28_v3_: d_49_v17_}))) != ((d_57_v22_) | (d_57_v22_))
        elif True:
            d_58_v23_: _dafny.Map
            d_58_v23_ = _dafny.Map({d_23_v0_: d_27_cf4_})
            d_58_v23_ = ((_dafny.Map({d_23_v0_: _dafny.SeqWithoutIsStrInference([d_28_v3_ for d_59_i6_ in range(default__.abs(269))])})) | (d_58_v23_)) | (d_58_v23_)
            d_60_v24_: C0
            nw4_ = C0()
            nw4_.ctor__()
            d_60_v24_ = nw4_
            d_61_v25_: _dafny.Array
            nw5_ = _dafny.Array(_dafny.Seq({}), 16)
            d_61_v25_ = nw5_
            d_62_v26_: _dafny.Set
            d_62_v26_ = _dafny.Set({d_31_v6_})
            d_63_v27_: _dafny.MultiSet
            d_63_v27_ = _dafny.MultiSet([default__.fm11(d_23_v0_, d_23_v0_, len(d_62_v26_), globalState)])
            index6_ = default__.safeIndex(150, (d_61_v25_).length(0))
            (d_61_v25_)[index6_] = default__.fm10((d_63_v27_).cardinality, globalState)
            d_64_v28_: _dafny.Array
            nw6_ = _dafny.Array(None, 24)
            nw6_[int(0)] = d_27_cf4_
            nw6_[int(1)] = (d_24_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knesulqj")))
            nw6_[int(2)] = d_27_cf4_
            nw6_[int(3)] = d_27_cf4_
            nw6_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxthifwc"))
            nw6_[int(5)] = ((d_58_v23_)[True] if (True) in (d_58_v23_) else d_24_v1_)
            nw6_[int(6)] = d_27_cf4_
            nw6_[int(7)] = (d_24_v1_) + (d_27_cf4_)
            nw6_[int(8)] = d_27_cf4_
            nw6_[int(9)] = d_24_v1_
            nw6_[int(10)] = d_27_cf4_
            nw6_[int(11)] = (d_24_v1_) + ((d_24_v1_).set(default__.safeIndex(len(d_24_v1_), len(d_24_v1_)), d_28_v3_))
            nw6_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqmctbicw"))
            nw6_[int(13)] = d_24_v1_
            nw6_[int(14)] = d_24_v1_
            nw6_[int(15)] = d_27_cf4_
            nw6_[int(16)] = d_27_cf4_
            nw6_[int(17)] = (d_24_v1_) + (_dafny.SeqWithoutIsStrInference([d_28_v3_ for d_65_i7_ in range(default__.abs(-182))]))
            nw6_[int(18)] = (d_27_cf4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkykgdr")))
            nw6_[int(19)] = d_24_v1_
            nw6_[int(20)] = (d_24_v1_) + (_dafny.SeqWithoutIsStrInference([(d_24_v1_)[default__.safeIndex(d_31_v6_, len(d_24_v1_))] for d_66_i8_ in range(default__.abs(18))]))
            nw6_[int(21)] = ((d_27_cf4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyohbsaxk")))).set(default__.safeIndex(d_31_v6_, len((d_27_cf4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyohbsaxk"))))), d_28_v3_)
            nw6_[int(22)] = d_24_v1_
            nw6_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krh"))
            d_64_v28_ = nw6_
            index7_ = default__.safeIndex(140, (d_64_v28_).length(0))
            (d_64_v28_)[index7_] = (d_27_cf4_) + (d_27_cf4_)
            d_67_v29_: _dafny.Set
            d_67_v29_ = _dafny.Set({not(True), d_23_v0_, d_23_v0_, d_23_v0_})
            d_68_v30_: _dafny.Seq
            d_68_v30_ = _dafny.SeqWithoutIsStrInference([d_23_v0_])
            d_69_v31_: _dafny.Seq
            d_69_v31_ = _dafny.SeqWithoutIsStrInference([d_68_v30_])
            index8_ = default__.safeIndex(150, (d_61_v25_).length(0))
            index9_ = default__.safeIndex(140, (d_64_v28_).length(0))
            rhs5_ = default__.fm0(len(d_67_v29_), globalState)
            rhs6_ = d_69_v31_
            rhs7_ = (d_24_v1_) + ((d_27_cf4_).set(default__.safeIndex(d_31_v6_, len(d_27_cf4_)), (d_24_v1_)[default__.safeIndex(d_31_v6_, len(d_24_v1_))]))
            rhs8_ = (0) - (d_31_v6_)
            rhs9_ = d_60_v24_
            lhs1_ = d_61_v25_
            lhs2_ = default__.safeIndex(150, (d_61_v25_).length(0))
            lhs3_ = d_64_v28_
            lhs4_ = default__.safeIndex(140, (d_64_v28_).length(0))
            d_23_v0_ = rhs5_
            lhs1_[lhs2_] = rhs6_
            lhs3_[lhs4_] = rhs7_
            r0 = rhs8_
            d_60_v24_ = rhs9_
            d_25_v2_ = d_25_v2_
            d_24_v1_ = d_27_cf4_
        r1 = default__.safeDivisionInt(d_31_v6_, len(d_27_cf4_))
        d_31_v6_ = d_31_v6_
        d_70_v32_: _dafny.Seq
        d_70_v32_ = _dafny.SeqWithoutIsStrInference([d_23_v0_, True, d_23_v0_])
        d_71_v33_: int
        d_71_v33_ = 514
        d_72_i9_: int
        d_72_i9_ = 0
        with _dafny.label("0"):
            while (d_70_v32_)[default__.safeIndex(d_71_v33_, len(d_70_v32_))]:
                with _dafny.c_label("0"):
                    if (d_72_i9_) >= (100):
                        raise _dafny.Break("0")
                    d_72_i9_ = (d_72_i9_) + (1)
                    d_73_v34_: str
                    d_73_v34_ = _dafny.CodePoint('b')
                    d_74_v35_: _dafny.Map
                    d_74_v35_ = _dafny.Map({(d_23_v0_) == (d_23_v0_): d_25_v2_})
                    d_75_v36_: C0
                    nw7_ = C0()
                    nw7_.ctor__()
                    d_75_v36_ = nw7_
                    d_76_v37_: _dafny.Seq
                    d_76_v37_ = _dafny.SeqWithoutIsStrInference([d_75_v36_, d_75_v36_])
                    d_77_v38_: _dafny.Map
                    d_77_v38_ = _dafny.Map({(d_76_v37_)[default__.safeIndex(d_71_v33_, len(d_76_v37_))]: d_71_v33_})
                    d_78_v39_: _dafny.MultiSet
                    d_78_v39_ = _dafny.MultiSet([d_71_v33_, d_71_v33_, d_71_v33_])
                    rhs10_ = ((d_77_v38_)[d_75_v36_] if (d_75_v36_) in (d_77_v38_) else ((d_78_v39_)[925] if (925) in (d_78_v39_) else d_71_v33_))
                    rhs11_ = d_73_v34_
                    rhs12_ = d_74_v35_
                    r1 = rhs10_
                    d_73_v34_ = rhs11_
                    d_74_v35_ = rhs12_
                    d_79_v40_: _dafny.Map
                    d_79_v40_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_73_v34_ for d_80_i10_ in range(default__.abs(867))])): default__.fm0(d_71_v33_, globalState)})
                    d_81_v41_: _dafny.Array
                    nw8_ = _dafny.Array(None, 11)
                    nw8_[int(0)] = d_23_v0_
                    nw8_[int(1)] = d_23_v0_
                    nw8_[int(2)] = d_23_v0_
                    nw8_[int(3)] = d_23_v0_
                    nw8_[int(4)] = d_23_v0_
                    nw8_[int(5)] = ((d_79_v40_)[d_71_v33_] if (d_71_v33_) in (d_79_v40_) else d_23_v0_)
                    nw8_[int(6)] = not(d_23_v0_)
                    nw8_[int(7)] = d_23_v0_
                    nw8_[int(8)] = d_23_v0_
                    nw8_[int(9)] = d_23_v0_
                    nw8_[int(10)] = d_23_v0_
                    d_81_v41_ = nw8_
                    index10_ = default__.safeIndex(760, (d_81_v41_).length(0))
                    (d_81_v41_)[index10_] = (False if d_23_v0_ else d_23_v0_)
                    d_82_v42_: D5
                    d_82_v42_ = D5_DC12(d_71_v33_, default__.fm12(globalState), d_23_v0_, d_24_v1_, d_71_v33_)
                    index11_ = default__.safeIndex(760, (d_81_v41_).length(0))
                    (d_81_v41_)[index11_] = not((d_23_v0_) or ((d_82_v42_).cf20))
                    rhs13_ = d_71_v33_
                    rhs14_ = (d_81_v41_)[default__.safeIndex(760, (d_81_v41_).length(0))]
                    rhs15_ = (d_81_v41_)[default__.safeIndex(760, (d_81_v41_).length(0))]
                    rhs16_ = (d_71_v33_) + (d_71_v33_)
                    lhs5_ = globalState
                    r0 = rhs13_
                    d_23_v0_ = rhs14_
                    lhs5_.f1 = rhs15_
                    r1 = rhs16_
                    d_83_v43_: _dafny.MultiSet
                    d_83_v43_ = _dafny.MultiSet([d_23_v0_])
                    d_83_v43_ = d_83_v43_
                    pass
            pass
        d_84_v44_: _dafny.MultiSet
        d_84_v44_ = _dafny.MultiSet([d_23_v0_])
        d_85_v45_: _dafny.Set
        d_85_v45_ = _dafny.Set({((d_84_v44_)[d_23_v0_] if (d_23_v0_) in (d_84_v44_) else d_71_v33_), len(d_70_v32_)})
        d_86_v46_: _dafny.Seq
        d_86_v46_ = _dafny.SeqWithoutIsStrInference([d_85_v45_])
        d_87_v47_: _dafny.Map
        d_87_v47_ = _dafny.Map({d_24_v1_: d_85_v45_})
        d_88_v49_: _dafny.Seq
        d_88_v49_ = _dafny.SeqWithoutIsStrInference([d_24_v1_, d_24_v1_, d_24_v1_])
        d_89_v50_: str
        d_89_v50_ = _dafny.CodePoint('l')
        d_90_v51_: D7
        d_90_v51_ = D7_DC14(d_87_v47_)
        d_91_v54_: _dafny.MultiSet
        d_91_v54_ = _dafny.MultiSet([d_24_v1_])
        d_92_v55_: _dafny.Array
        nw9_ = _dafny.Array(None, 18)
        nw9_[int(0)] = d_87_v47_
        nw9_[int(1)] = d_87_v47_
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: _dafny.Seq
            for compr_4_ in (d_88_v49_).Elements:
                d_93_v48_: _dafny.Seq = compr_4_
                if (d_93_v48_) in (d_88_v49_):
                    coll4_[d_93_v48_] = d_85_v45_
            return _dafny.Map(coll4_)
        nw9_[int(2)] = iife4_()
        
        nw9_[int(3)] = default__.fm13(d_23_v0_, False, d_24_v1_, globalState)
        nw9_[int(4)] = _dafny.Map({d_24_v1_: d_85_v45_})
        nw9_[int(5)] = d_87_v47_
        nw9_[int(6)] = _dafny.Map({d_24_v1_: d_85_v45_})
        nw9_[int(7)] = (d_87_v47_) | (d_87_v47_)
        nw9_[int(8)] = _dafny.Map({((d_25_v2_)).set(default__.safeIndex(d_71_v33_, len((d_25_v2_))), d_89_v50_): d_85_v45_})
        nw9_[int(9)] = (d_87_v47_) | ((d_90_v51_).cf24)
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Seq
            for compr_5_ in (d_88_v49_).Elements:
                d_94_v52_: _dafny.Seq = compr_5_
                if (d_94_v52_) in (d_88_v49_):
                    coll5_[d_94_v52_] = d_85_v45_
            return _dafny.Map(coll5_)
        nw9_[int(10)] = (d_87_v47_ if d_23_v0_ else iife5_()
        )
        nw9_[int(11)] = d_87_v47_
        nw9_[int(12)] = (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrfel")): _dafny.Set({d_71_v33_, d_71_v33_})})).set(d_24_v1_, d_85_v45_)
        nw9_[int(13)] = d_87_v47_
        nw9_[int(14)] = _dafny.Map({d_24_v1_: d_85_v45_})
        nw9_[int(15)] = (_dafny.Map({d_24_v1_: d_85_v45_})) | (d_87_v47_)
        nw9_[int(16)] = _dafny.Map({d_24_v1_: d_85_v45_})
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: _dafny.Seq
            for compr_6_ in (d_91_v54_).Elements:
                d_95_v53_: _dafny.Seq = compr_6_
                if (d_95_v53_) in (d_91_v54_):
                    coll6_[d_95_v53_] = d_85_v45_
            return _dafny.Map(coll6_)
        nw9_[int(17)] = iife6_()
        
        d_92_v55_ = nw9_
        index12_ = default__.safeIndex(5, (d_92_v55_).length(0))
        (d_92_v55_)[index12_] = d_87_v47_
        d_96_v56_: _dafny.Array
        nw10_ = _dafny.Array(int(0), 4)
        d_96_v56_ = nw10_
        d_97_v57_: _dafny.Array
        def lambda6_(d_98_v33_):
            def lambda7_(d_99_i11_):
                return D2_DC3(_dafny.MultiSet([d_98_v33_]))

            return lambda7_

        init3_ = lambda6_(d_71_v33_)
        nw11_ = _dafny.Array(None, 23)
        for i0_3_ in range(nw11_.length(0)):
            nw11_[i0_3_] = init3_(i0_3_)
        d_97_v57_ = nw11_
        d_100_v58_: D2
        d_100_v58_ = D2_DC3(_dafny.MultiSet([d_71_v33_, d_71_v33_]))
        index13_ = default__.safeIndex(306, (d_97_v57_).length(0))
        (d_97_v57_)[index13_] = d_100_v58_
        d_101_v59_: D4
        d_101_v59_ = D4_DC9(d_96_v56_, len(d_88_v49_), d_96_v56_)
        d_102_v60_: _dafny.MultiSet
        d_102_v60_ = _dafny.MultiSet([d_71_v33_, d_71_v33_])
        index14_ = default__.safeIndex(5, (d_92_v55_).length(0))
        index15_ = default__.safeIndex(306, (d_97_v57_).length(0))
        rhs17_ = d_86_v46_
        rhs18_ = d_87_v47_
        rhs19_ = ((d_101_v59_).cf13 if (default__.fm0(d_71_v33_, globalState) if d_23_v0_ else d_23_v0_) else d_96_v56_)
        rhs20_ = ((0) - (((d_102_v60_)[353] if (353) in (d_102_v60_) else d_71_v33_))) >= (d_71_v33_)
        rhs21_ = d_100_v58_
        lhs6_ = d_92_v55_
        lhs7_ = default__.safeIndex(5, (d_92_v55_).length(0))
        lhs8_ = d_97_v57_
        lhs9_ = default__.safeIndex(306, (d_97_v57_).length(0))
        d_86_v46_ = rhs17_
        lhs6_[lhs7_] = rhs18_
        d_96_v56_ = rhs19_
        r2 = rhs20_
        lhs8_[lhs9_] = rhs21_
        d_103_v61_: C0
        nw12_ = C0()
        nw12_.ctor__()
        d_103_v61_ = nw12_
        hi0_ = d_71_v33_
        for d_104_i12_ in range(d_71_v33_, hi0_):
            d_105_v62_: _dafny.Array
            nw13_ = _dafny.Array(False, 9)
            d_105_v62_ = nw13_
            index16_ = default__.safeIndex(379, (d_105_v62_).length(0))
            (d_105_v62_)[index16_] = True
            index17_ = default__.safeIndex(379, (d_105_v62_).length(0))
            (d_105_v62_)[index17_] = d_23_v0_
            d_106_v63_: _dafny.Array
            nw14_ = _dafny.Array(_dafny.MultiSet({}), 20)
            d_106_v63_ = nw14_
            d_106_v63_ = d_106_v63_
            d_107_v64_: C0
            nw15_ = C0()
            nw15_.ctor__()
            d_107_v64_ = nw15_
            (d_103_v61_).m1(d_104_i12_, not(d_23_v0_), globalState)
        d_108_v65_: D0
        d_108_v65_ = D0_DC0((d_71_v33_) != (len(_dafny.SeqWithoutIsStrInference([d_71_v33_ for d_109_i13_ in range(default__.abs(89))]))))
        pat_let_tv0_ = d_102_v60_
        pat_let_tv1_ = d_102_v60_
        def iife7_(_pat_let0_0):
            def iife8_(d_110_dt__update__tmp_h0_):
                def iife9_(_pat_let1_0):
                    def iife10_(d_111_dt__update_hcf0_h0_):
                        return D0_DC0(d_111_dt__update_hcf0_h0_)
                    return iife10_(_pat_let1_0)
                return iife9_((pat_let_tv0_).issubset(pat_let_tv1_))
            return iife8_(_pat_let0_0)
        d_108_v65_ = iife7_(d_108_v65_)
        r0 = d_71_v33_
        r1 = -928
        r2 = d_23_v0_
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_112_globalState_: GlobalState
        nw16_ = GlobalState()
        nw16_.ctor__(_dafny.SeqWithoutIsStrInference([879]), True)
        d_112_globalState_ = nw16_
        d_113_v0_: _dafny.Array
        def lambda8_(d_114_i1_):
            return (_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference([645, 936, len(_dafny.SeqWithoutIsStrInference([901, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_115_i2_ in range(default__.abs(945))])), 989, 532, 440]))]))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_116_i3_ in range(default__.abs(274))]))])), (0) - (len(_dafny.SeqWithoutIsStrInference([467 for d_117_i4_ in range(default__.abs(872))])))])) != (_dafny.MultiSet([874]))

        init4_ = lambda8_
        nw17_ = _dafny.Array(None, 12)
        for i0_4_ in range(nw17_.length(0)):
            nw17_[i0_4_] = init4_(i0_4_)
        d_113_v0_ = nw17_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_113_v0_).length(0)):
            d_118_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_118_i0_)) and ((d_118_i0_) < ((d_113_v0_).length(0)))):
                (d_113_v0_)[(d_118_i0_)] = (D0_DC0(True)).cf0
        d_119_v1_: D0
        d_119_v1_ = D0_DC0(not(True))
        source2_ = d_119_v1_
        if source2_.is_DC1:
            d_120___mcc_h0_ = source2_.cf1
            d_121___mcc_h1_ = source2_.cf2
            d_122___mcc_h2_ = source2_.cf3
            d_123_cf3_ = d_122___mcc_h2_
            d_124_cf2_ = d_121___mcc_h1_
            d_125_cf1_ = d_120___mcc_h0_
            d_126_v2_: _dafny.Seq
            d_126_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkd"))
            d_127_v3_: _dafny.Map
            d_127_v3_ = _dafny.Map({d_125_cf1_: d_124_cf2_})
            d_128_v4_: _dafny.Seq
            d_128_v4_ = _dafny.SeqWithoutIsStrInference([892, len(d_127_v3_)])
            d_129_v5_: _dafny.Map
            d_129_v5_ = _dafny.Map({d_123_cf3_: d_128_v4_})
            if (len(d_126_v2_)) not in ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_130_i6_ in range(default__.abs(488))])) for d_131_i5_ in range(default__.abs(308))])) + (((d_129_v5_)[d_123_cf3_] if (d_123_cf3_) in (d_129_v5_) else d_128_v4_))):
                d_125_cf1_ = default__.safeModuloInt(len(_dafny.Map({d_123_cf3_: _dafny.SeqWithoutIsStrInference([d_124_cf2_ for d_132_i7_ in range(default__.abs(669))])})), d_124_cf2_)
                d_133_v6_: _dafny.MultiSet
                d_133_v6_ = _dafny.MultiSet([(len(d_126_v2_)) <= (d_125_cf1_), d_123_cf3_, d_123_cf3_, not (d_123_cf3_) or (d_123_cf3_), not (d_123_cf3_) or (default__.fm0(d_125_cf1_, d_112_globalState_))])
                d_133_v6_ = (d_133_v6_).intersection(d_133_v6_)
                d_134_v7_: int
                d_135_v8_: int
                d_136_v9_: bool
                out0_: int
                out1_: int
                out2_: bool
                out0_, out1_, out2_ = default__.m0(_dafny.MultiSet([d_128_v4_, d_128_v4_]), d_112_globalState_)
                d_134_v7_ = out0_
                d_135_v8_ = out1_
                d_136_v9_ = out2_
                d_137_v10_: _dafny.Map
                d_137_v10_ = _dafny.Map({d_136_v9_: not(not (d_136_v9_) or (not((d_119_v1_).cf0)))})
                d_138_v11_: _dafny.Seq
                d_138_v11_ = _dafny.SeqWithoutIsStrInference([d_123_cf3_])
                d_137_v10_ = (d_137_v10_).set(d_136_v9_, (d_138_v11_)[default__.safeIndex(d_134_v7_, len(d_138_v11_))])
                (d_112_globalState_).f1 = d_136_v9_
            elif True:
                d_139_v12_: int
                d_140_v13_: int
                d_141_v14_: bool
                out3_: int
                out4_: int
                out5_: bool
                out3_, out4_, out5_ = default__.m0(_dafny.MultiSet([d_128_v4_, d_128_v4_]), d_112_globalState_)
                d_139_v12_ = out3_
                d_140_v13_ = out4_
                d_141_v14_ = out5_
                (d_112_globalState_).f1 = (default__.safeModuloInt((0) - (d_125_cf1_), default__.fm1(d_124_cf2_, d_112_globalState_))) <= (d_139_v12_)
                d_126_v2_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ne"))) + (d_126_v2_)) + ((d_126_v2_) + (d_126_v2_))
                (d_112_globalState_).f1 = False
                rhs22_ = d_141_v14_
                rhs23_ = ((395) - (d_125_cf1_)) == (d_124_cf2_)
                d_141_v14_ = rhs22_
                d_123_cf3_ = rhs23_
            if not (not(d_123_cf3_)) or ((d_119_v1_).cf0):
                d_142_v15_: _dafny.Array
                nw18_ = _dafny.Array(_dafny.Map({}), 9)
                d_142_v15_ = nw18_
                nw19_ = _dafny.Array(_dafny.Map({}), 19)
                d_142_v15_ = nw19_
                index18_ = default__.safeIndex(47, (d_113_v0_).length(0))
                (d_113_v0_)[index18_] = default__.fm0(d_125_cf1_, d_112_globalState_)
                index19_ = default__.safeIndex(47, (d_113_v0_).length(0))
                (d_113_v0_)[index19_] = not (d_123_cf3_) or (d_123_cf3_)
                (d_112_globalState_).f0 = d_128_v4_
                d_125_cf1_ = default__.fm1((len(d_128_v4_)) + (d_125_cf1_), d_112_globalState_)
                d_124_cf2_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_143_i8_ in range(default__.abs(332))])), d_124_cf2_)
            elif True:
                d_144_v16_: _dafny.Array
                def lambda9_(d_145_cf1_):
                    def lambda10_(d_146_i9_):
                        return (d_146_i9_) + (d_145_cf1_)

                    return lambda10_

                init5_ = lambda9_(d_125_cf1_)
                nw20_ = _dafny.Array(None, 26)
                for i0_5_ in range(nw20_.length(0)):
                    nw20_[i0_5_] = init5_(i0_5_)
                d_144_v16_ = nw20_
                def lambda11_(d_147_cf1_, d_148_v3_, d_149_cf2_):
                    def lambda12_(d_150_i10_):
                        return (d_150_i10_) + (((d_148_v3_)[d_147_cf1_] if (d_147_cf1_) in (d_148_v3_) else (0) - (d_149_cf2_)))

                    return lambda12_

                init6_ = lambda11_(d_125_cf1_, d_127_v3_, d_124_cf2_)
                nw21_ = _dafny.Array(None, 22)
                for i0_6_ in range(nw21_.length(0)):
                    nw21_[i0_6_] = init6_(i0_6_)
                d_144_v16_ = nw21_
                d_151_v17_: _dafny.Map
                d_151_v17_ = _dafny.Map({d_124_cf2_: d_123_cf3_})
                d_152_v18_: _dafny.Seq
                d_152_v18_ = _dafny.SeqWithoutIsStrInference([d_151_v17_, d_151_v17_])
                rhs24_ = d_123_cf3_
                rhs25_ = d_152_v18_
                lhs10_ = d_112_globalState_
                lhs10_.f1 = rhs24_
                d_152_v18_ = rhs25_
                (d_112_globalState_).f1 = d_123_cf3_
                index20_ = default__.safeIndex(77, (d_113_v0_).length(0))
                (d_113_v0_)[index20_] = d_123_cf3_
                index21_ = default__.safeIndex(77, (d_113_v0_).length(0))
                rhs26_ = d_123_cf3_
                rhs27_ = (d_125_cf1_) - (653)
                lhs11_ = d_113_v0_
                lhs12_ = default__.safeIndex(77, (d_113_v0_).length(0))
                lhs11_[lhs12_] = rhs26_
                d_124_cf2_ = rhs27_
                d_153_v19_: _dafny.Map
                d_153_v19_ = _dafny.Map({(d_113_v0_)[default__.safeIndex(77, (d_113_v0_).length(0))]: d_125_cf1_})
                d_154_v20_: _dafny.MultiSet
                d_154_v20_ = _dafny.MultiSet([d_128_v4_])
                d_155_v21_: int
                d_156_v22_: int
                d_157_v23_: bool
                out6_: int
                out7_: int
                out8_: bool
                out6_, out7_, out8_ = default__.m0((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_124_cf2_, d_125_cf1_]), d_128_v4_, _dafny.SeqWithoutIsStrInference([d_124_cf2_, len(d_153_v19_)]), d_128_v4_, d_128_v4_])).intersection(d_154_v20_), d_112_globalState_)
                d_155_v21_ = out6_
                d_156_v22_ = out7_
                d_157_v23_ = out8_
            d_124_cf2_ = (0) - (d_124_cf2_)
            d_158_v24_: _dafny.Map
            d_158_v24_ = _dafny.Map({d_123_cf3_: True})
            d_158_v24_ = (d_158_v24_).set((d_125_cf1_) > ((0) - ((0) - (d_125_cf1_))), d_123_cf3_)
        elif True:
            d_159___mcc_h3_ = source2_.cf0
            d_160_cf0_ = d_159___mcc_h3_
            d_161_v25_: int
            d_161_v25_ = 136
            d_162_v26_: _dafny.Array
            def lambda13_(d_163_v25_):
                def lambda14_(d_164_i11_):
                    return default__.safeModuloInt(d_164_i11_, d_163_v25_)

                return lambda14_

            init7_ = lambda13_(d_161_v25_)
            nw22_ = _dafny.Array(None, 2)
            for i0_7_ in range(nw22_.length(0)):
                nw22_[i0_7_] = init7_(i0_7_)
            d_162_v26_ = nw22_
            index22_ = default__.safeIndex(660, (d_162_v26_).length(0))
            (d_162_v26_)[index22_] = default__.safeDivisionInt(d_161_v25_, d_161_v25_)
            d_165_v27_: _dafny.MultiSet
            d_165_v27_ = _dafny.MultiSet([d_160_cf0_])
            d_166_v28_: _dafny.Seq
            d_166_v28_ = _dafny.SeqWithoutIsStrInference([d_165_v27_])
            d_167_v29_: D0
            d_167_v29_ = D0_DC1(d_161_v25_, d_161_v25_, d_160_cf0_)
            d_168_v30_: _dafny.Map
            d_168_v30_ = _dafny.Map({not(d_160_cf0_): -24})
            index23_ = default__.safeIndex(660, (d_162_v26_).length(0))
            rhs28_ = (((d_166_v28_)[default__.safeIndex(d_161_v25_, len(d_166_v28_))] if d_160_cf0_ else (_dafny.MultiSet([(d_167_v29_).cf3, d_160_cf0_])).set(default__.fm0(len(d_168_v30_), d_112_globalState_), default__.abs(d_161_v25_)))).cardinality
            rhs29_ = d_161_v25_
            rhs30_ = (len(_dafny.Map({d_160_cf0_: d_160_cf0_}))) - (d_161_v25_)
            rhs31_ = ((d_168_v30_)[True] if (True) in (d_168_v30_) else d_161_v25_)
            rhs32_ = default__.fm0(d_161_v25_, d_112_globalState_)
            lhs13_ = d_162_v26_
            lhs14_ = default__.safeIndex(660, (d_162_v26_).length(0))
            d_161_v25_ = rhs28_
            d_161_v25_ = rhs29_
            d_161_v25_ = rhs30_
            lhs13_[lhs14_] = rhs31_
            d_160_cf0_ = rhs32_
            d_169_v31_: _dafny.Set
            d_169_v31_ = _dafny.Set({d_168_v30_, d_168_v30_})
            d_170_v32_: _dafny.Set
            d_170_v32_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wop"))})
            index24_ = default__.safeIndex(634, (d_113_v0_).length(0))
            (d_113_v0_)[index24_] = (d_170_v32_) != (d_170_v32_)
            d_171_v33_: _dafny.Seq
            d_171_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqinws"))
            index25_ = default__.safeIndex(634, (d_113_v0_).length(0))
            rhs33_ = d_169_v31_
            rhs34_ = (d_171_v33_) != ((d_171_v33_).set(default__.safeIndex(d_161_v25_, len(d_171_v33_)), _dafny.CodePoint('p')))
            rhs35_ = not (d_160_cf0_) or (((0) - (d_161_v25_)) <= (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_172_i12_ in range(default__.abs(786))]))))
            rhs36_ = (0) - (default__.fm1((0) - (d_161_v25_), d_112_globalState_))
            rhs37_ = d_160_cf0_
            lhs15_ = d_112_globalState_
            lhs16_ = d_113_v0_
            lhs17_ = default__.safeIndex(634, (d_113_v0_).length(0))
            d_169_v31_ = rhs33_
            lhs15_.f1 = rhs34_
            d_160_cf0_ = rhs35_
            d_161_v25_ = rhs36_
            lhs16_[lhs17_] = rhs37_
            source3_ = d_167_v29_
            if source3_.is_DC1:
                d_173___mcc_h4_ = source3_.cf1
                d_174___mcc_h5_ = source3_.cf2
                d_175___mcc_h6_ = source3_.cf3
                d_176_cf3_ = d_175___mcc_h6_
                d_177_cf2_ = d_174___mcc_h5_
                d_178_cf1_ = d_173___mcc_h4_
                d_179_v34_: C0
                nw23_ = C0()
                nw23_.ctor__()
                d_179_v34_ = nw23_
                d_180_v35_: _dafny.Seq
                d_180_v35_ = _dafny.SeqWithoutIsStrInference([(d_113_v0_)[default__.safeIndex(634, (d_113_v0_).length(0))]])
                d_181_v36_: _dafny.Seq
                d_181_v36_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_113_v0_)[default__.safeIndex(634, (d_113_v0_).length(0))]: d_161_v25_}))])
                d_182_v37_: _dafny.Map
                d_182_v37_ = _dafny.Map({496: d_179_v34_})
                rhs38_ = D0_DC0((d_165_v27_) == (d_165_v27_))
                rhs39_ = _dafny.SeqWithoutIsStrInference([False, (d_113_v0_)[default__.safeIndex(634, (d_113_v0_).length(0))], d_160_cf0_])
                rhs40_ = ((d_181_v36_) + (d_181_v36_)) + (_dafny.SeqWithoutIsStrInference([len(d_180_v35_), d_177_cf2_, d_177_cf2_]))
                rhs41_ = ((0) - (len(d_171_v33_))) + (len((d_182_v37_) | (d_182_v37_)))
                rhs42_ = d_179_v34_
                lhs18_ = d_112_globalState_
                d_119_v1_ = rhs38_
                d_180_v35_ = rhs39_
                lhs18_.f0 = rhs40_
                d_177_cf2_ = rhs41_
                d_179_v34_ = rhs42_
                d_183_v38_: C0
                nw24_ = C0()
                nw24_.ctor__()
                d_183_v38_ = nw24_
                (d_183_v38_).m1(d_178_cf1_, d_176_cf3_, d_112_globalState_)
            elif True:
                d_184___mcc_h7_ = source3_.cf0
                d_185_cf0_ = d_184___mcc_h7_
                d_186_v39_: C0
                nw25_ = C0()
                nw25_.ctor__()
                d_186_v39_ = nw25_
                d_187_v40_: _dafny.Set
                d_187_v40_ = _dafny.Set({d_186_v39_, d_186_v39_, d_186_v39_, d_186_v39_})
                d_187_v40_ = ((d_187_v40_).intersection(d_187_v40_)) - (_dafny.Set({d_186_v39_}))
                index26_ = default__.safeIndex(660, (d_162_v26_).length(0))
                (d_162_v26_)[index26_] = d_161_v25_
                d_188_v41_: _dafny.Seq
                d_188_v41_ = _dafny.SeqWithoutIsStrInference([(d_162_v26_)[default__.safeIndex(660, (d_162_v26_).length(0))]])
                d_189_v42_: _dafny.MultiSet
                d_189_v42_ = _dafny.MultiSet([d_188_v41_])
                d_190_v43_: int
                d_191_v44_: int
                d_192_v45_: bool
                out9_: int
                out10_: int
                out11_: bool
                out9_, out10_, out11_ = default__.m0(d_189_v42_, d_112_globalState_)
                d_190_v43_ = out9_
                d_191_v44_ = out10_
                d_192_v45_ = out11_
                d_171_v33_ = d_171_v33_
            d_193_v46_: C0
            nw26_ = C0()
            nw26_.ctor__()
            d_193_v46_ = nw26_
            d_194_v47_: str
            d_194_v47_ = _dafny.CodePoint('j')
            d_195_v48_: _dafny.Map
            d_195_v48_ = _dafny.Map({d_194_v47_: d_193_v46_})
            d_193_v46_ = ((d_195_v48_)[d_194_v47_] if (d_194_v47_) in (d_195_v48_) else d_193_v46_)
        d_196_v49_: C0
        nw27_ = C0()
        nw27_.ctor__()
        d_196_v49_ = nw27_
        d_197_v50_: bool
        d_197_v50_ = False
        if d_197_v50_:
            d_198_v51_: str
            d_198_v51_ = _dafny.CodePoint('i')
            d_198_v51_ = d_198_v51_
            d_199_v52_: int
            d_199_v52_ = 296
            d_200_v53_: _dafny.Set
            d_200_v53_ = _dafny.Set({(0) - (d_199_v52_), (0) - (len(_dafny.SeqWithoutIsStrInference([d_197_v50_, d_197_v50_])))})
            d_199_v52_ = default__.fm1(len((d_200_v53_ if d_197_v50_ else d_200_v53_)), d_112_globalState_)
            d_198_v51_ = d_198_v51_
            d_201_v54_: C0
            nw28_ = C0()
            nw28_.ctor__()
            d_201_v54_ = nw28_
            if (default__.safeModuloInt(d_199_v52_, d_199_v52_)) <= (d_199_v52_):
                (d_196_v49_).m1((d_199_v52_) + (d_199_v52_), d_197_v50_, d_112_globalState_)
                index27_ = default__.safeIndex(720, (d_113_v0_).length(0))
                (d_113_v0_)[index27_] = d_197_v50_
                index28_ = default__.safeIndex(720, (d_113_v0_).length(0))
                (d_113_v0_)[index28_] = d_197_v50_
                (d_112_globalState_).f1 = not(not((d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))]))
                d_202_v55_: _dafny.Array
                def lambda15_(d_203_i13_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvo"))

                init8_ = lambda15_
                nw29_ = _dafny.Array(None, 27)
                for i0_8_ in range(nw29_.length(0)):
                    nw29_[i0_8_] = init8_(i0_8_)
                d_202_v55_ = nw29_
                d_204_v56_: _dafny.Seq
                d_204_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdpbjuu"))
                index29_ = default__.safeIndex(674, (d_202_v55_).length(0))
                (d_202_v55_)[index29_] = (d_204_v56_) + (d_204_v56_)
                d_205_v57_: D0
                d_205_v57_ = D0_DC1(692, d_199_v52_, (d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))])
                d_206_v58_: _dafny.MultiSet
                d_206_v58_ = _dafny.MultiSet([(d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))], (d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))], d_197_v50_, (d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))], not((d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))])])
                index30_ = default__.safeIndex(674, (d_202_v55_).length(0))
                (d_202_v55_)[index30_] = (_dafny.SeqWithoutIsStrInference([d_198_v51_ for d_207_i14_ in range(default__.abs(196))]) if not((default__.fm5((d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))], (d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))], d_205_v57_, d_112_globalState_)).isdisjoint((D2_DC3(_dafny.MultiSet([d_199_v52_, ((d_206_v58_).set((d_113_v0_)[default__.safeIndex(720, (d_113_v0_).length(0))], default__.abs(d_199_v52_))).cardinality]))).cf5)) else (_dafny.SeqWithoutIsStrInference([d_198_v51_ for d_208_i15_ in range(default__.abs(-556))])) + (d_204_v56_))
                d_209_v59_: _dafny.Seq
                d_209_v59_ = _dafny.SeqWithoutIsStrInference([d_199_v52_, (d_206_v58_).cardinality, (0) - (d_199_v52_)])
                d_199_v52_ = ((d_209_v59_)[default__.safeIndex(341, len(d_209_v59_))]) * ((d_199_v52_) * (d_199_v52_))
            elif True:
                d_210_v60_: _dafny.Array
                def lambda16_(d_211_i16_):
                    return (d_211_i16_) + (len(_dafny.SeqWithoutIsStrInference([False])))

                init9_ = lambda16_
                nw30_ = _dafny.Array(None, 24)
                for i0_9_ in range(nw30_.length(0)):
                    nw30_[i0_9_] = init9_(i0_9_)
                d_210_v60_ = nw30_
                index31_ = default__.safeIndex(922, (d_210_v60_).length(0))
                (d_210_v60_)[index31_] = (0) - (d_199_v52_)
                index32_ = default__.safeIndex(922, (d_210_v60_).length(0))
                (d_210_v60_)[index32_] = d_199_v52_
                d_212_v61_: _dafny.Map
                d_212_v61_ = _dafny.Map({d_198_v51_: d_199_v52_})
                d_213_v62_: _dafny.Map
                d_213_v62_ = _dafny.Map({(d_210_v60_)[default__.safeIndex(922, (d_210_v60_).length(0))]: (d_210_v60_)[default__.safeIndex(922, (d_210_v60_).length(0))]})
                d_212_v61_ = (d_212_v61_).set(d_198_v51_, default__.safeDivisionInt(len(d_213_v62_), 196))
                d_214_v63_: _dafny.MultiSet
                d_214_v63_ = _dafny.MultiSet([d_198_v51_, _dafny.CodePoint('w'), d_198_v51_, d_198_v51_, d_198_v51_])
                d_215_v64_: _dafny.Seq
                d_215_v64_ = _dafny.SeqWithoutIsStrInference([d_198_v51_, d_198_v51_, _dafny.CodePoint('a'), d_198_v51_, d_198_v51_])
                d_216_v65_: _dafny.Seq
                d_216_v65_ = _dafny.SeqWithoutIsStrInference([d_214_v63_, d_214_v63_, (d_214_v63_).set(d_198_v51_, default__.abs(d_199_v52_)), _dafny.MultiSet(d_215_v64_), d_214_v63_])
                d_214_v63_ = (d_216_v65_)[default__.safeIndex(d_199_v52_, len(d_216_v65_))]
                d_199_v52_ = default__.safeDivisionInt(default__.safeDivisionInt(d_199_v52_, d_199_v52_), d_199_v52_)
                (d_196_v49_).m1((d_210_v60_)[default__.safeIndex(922, (d_210_v60_).length(0))], (not(d_197_v50_) if d_197_v50_ else d_197_v50_), d_112_globalState_)
        elif True:
            d_217_v66_: _dafny.MultiSet
            d_217_v66_ = _dafny.MultiSet([not(True)])
            d_218_v67_: int
            d_218_v67_ = 399
            rhs43_ = d_217_v66_
            rhs44_ = d_218_v67_
            d_217_v66_ = rhs43_
            d_218_v67_ = rhs44_
            d_218_v67_ = 644
            d_196_v49_ = d_196_v49_
            d_219_v68_: _dafny.Array
            nw31_ = _dafny.Array(int(0), 7)
            d_219_v68_ = nw31_
            index33_ = default__.safeIndex(54, (d_219_v68_).length(0))
            (d_219_v68_)[index33_] = default__.fm1(d_218_v67_, d_112_globalState_)
            index34_ = default__.safeIndex(54, (d_219_v68_).length(0))
            (d_219_v68_)[index34_] = d_218_v67_
            d_220_v69_: _dafny.Seq
            d_220_v69_ = _dafny.SeqWithoutIsStrInference([d_197_v50_])
            d_221_v70_: _dafny.MultiSet
            d_221_v70_ = _dafny.MultiSet([d_220_v69_])
            d_222_v71_: _dafny.Map
            d_222_v71_ = _dafny.Map({(d_219_v68_)[default__.safeIndex(54, (d_219_v68_).length(0))]: d_221_v70_})
            d_197_v50_ = (True if (((d_222_v71_)[d_218_v67_] if (d_218_v67_) in (d_222_v71_) else d_221_v70_)).isdisjoint(d_221_v70_) else d_197_v50_)
        d_223_v72_: int
        d_223_v72_ = 694
        d_223_v72_ = d_223_v72_
        d_224_v73_: _dafny.Array
        def lambda17_(d_225_v50_):
            def lambda18_(d_226_i18_):
                return _dafny.SeqWithoutIsStrInference([True, d_225_v50_])

            return lambda18_

        init10_ = lambda17_(d_197_v50_)
        nw32_ = _dafny.Array(None, 3)
        for i0_10_ in range(nw32_.length(0)):
            nw32_[i0_10_] = init10_(i0_10_)
        d_224_v73_ = nw32_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_224_v73_).length(0)):
            d_227_i17_: int = guard_loop_1_
            if (True) and (((0) <= (d_227_i17_)) and ((d_227_i17_) < ((d_224_v73_).length(0)))):
                (d_224_v73_)[(d_227_i17_)] = _dafny.SeqWithoutIsStrInference([d_197_v50_, d_197_v50_, d_197_v50_])
        d_228_v74_: C0
        nw33_ = C0()
        nw33_.ctor__()
        d_228_v74_ = nw33_
        if default__.fm0((d_223_v72_) - (d_223_v72_), d_112_globalState_):
            d_229_v75_: str
            d_229_v75_ = _dafny.CodePoint('o')
            d_230_v76_: _dafny.Map
            d_230_v76_ = _dafny.Map({d_197_v50_: d_229_v75_})
            d_231_v77_: _dafny.Map
            d_231_v77_ = _dafny.Map({d_197_v50_: d_230_v76_})
            d_232_v78_: _dafny.Map
            d_232_v78_ = ((d_231_v77_)[d_197_v50_] if (d_197_v50_) in (d_231_v77_) else d_230_v76_)
            d_233_v79_: _dafny.Seq
            d_233_v79_ = _dafny.SeqWithoutIsStrInference([len((d_230_v76_ if d_197_v50_ else (d_232_v78_)))])
            rhs45_ = d_223_v72_
            rhs46_ = d_233_v79_
            lhs19_ = d_112_globalState_
            d_223_v72_ = rhs45_
            lhs19_.f0 = rhs46_
            d_234_v80_: _dafny.MultiSet
            d_234_v80_ = _dafny.MultiSet([d_223_v72_, d_223_v72_])
            d_235_v81_: _dafny.Map
            d_235_v81_ = _dafny.Map({d_223_v72_: d_197_v50_})
            d_236_v82_: D2
            d_236_v82_ = D2_DC5(d_223_v72_, d_223_v72_)
            d_237_v84_: _dafny.Array
            nw34_ = _dafny.Array(int(0), 26)
            d_237_v84_ = nw34_
            d_238_v85_: _dafny.MultiSet
            d_238_v85_ = _dafny.MultiSet([d_237_v84_])
            d_239_v86_: _dafny.Array
            nw35_ = _dafny.Array(None, 9)
            nw35_[int(0)] = (d_234_v80_).cardinality
            nw35_[int(1)] = d_223_v72_
            nw35_[int(2)] = len(default__.fm6(d_197_v50_, d_235_v81_, (D0_DC0(d_197_v50_)).cf0, d_223_v72_, d_112_globalState_))
            nw35_[int(3)] = (0) - (d_223_v72_)
            nw35_[int(4)] = (d_236_v82_).cf8
            def iife11_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(144, 803):
                    d_240_v83_: int = compr_7_
                    if ((144) <= (d_240_v83_)) and ((d_240_v83_) < (803)):
                        coll7_[(d_240_v83_) + (d_223_v72_)] = 785
                return _dafny.Map(coll7_)
            nw35_[int(5)] = default__.fm1(len(iife11_()
            ), d_112_globalState_)
            nw35_[int(6)] = (d_238_v85_).cardinality
            nw35_[int(7)] = d_223_v72_
            nw35_[int(8)] = d_223_v72_
            d_239_v86_ = nw35_
            d_241_v87_: _dafny.Seq
            d_241_v87_ = _dafny.SeqWithoutIsStrInference([d_237_v84_, d_237_v84_])
            d_242_v88_: _dafny.Set
            d_242_v88_ = _dafny.Set({len(d_241_v87_)})
            d_242_v88_ = _dafny.Set({(0) - (len((d_233_v79_).set(default__.safeIndex(d_223_v72_, len(d_233_v79_)), d_223_v72_))), d_223_v72_})
            d_243_v89_: _dafny.Seq
            d_243_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlrf"))
            d_244_v90_: _dafny.Map
            d_244_v90_ = _dafny.Map({(0) - (d_223_v72_): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iff")) if not(not(d_197_v50_)) else d_243_v89_)})
            d_243_v89_ = ((d_244_v90_)[d_223_v72_] if (d_223_v72_) in (d_244_v90_) else d_243_v89_)
            index35_ = default__.safeIndex(71, (d_239_v86_).length(0))
            (d_239_v86_)[index35_] = len(_dafny.Map({d_197_v50_: d_223_v72_}))
            index36_ = default__.safeIndex(71, (d_239_v86_).length(0))
            (d_239_v86_)[index36_] = default__.safeModuloInt(len((d_243_v89_) + (d_243_v89_)), (d_233_v79_)[default__.safeIndex(456, len(d_233_v79_))])
            d_196_v49_ = d_228_v74_
        elif True:
            d_245_v91_: C0
            nw36_ = C0()
            nw36_.ctor__()
            d_245_v91_ = nw36_
            d_196_v49_ = d_228_v74_
            (d_228_v74_).m1(d_223_v72_, True, d_112_globalState_)
            d_246_v92_: _dafny.Seq
            d_246_v92_ = _dafny.SeqWithoutIsStrInference([(256) < (d_223_v72_)])
            d_197_v50_ = (d_246_v92_)[default__.safeIndex(-242, len(d_246_v92_))]
            d_247_v93_: _dafny.Seq
            d_247_v93_ = _dafny.SeqWithoutIsStrInference([(0) - (d_223_v72_), (d_223_v72_ if d_197_v50_ else d_223_v72_), d_223_v72_])
            d_223_v72_ = (d_247_v93_)[default__.safeIndex(-252, len(d_247_v93_))]
        d_248_v94_: _dafny.MultiSet
        d_248_v94_ = _dafny.MultiSet([d_228_v74_, d_228_v74_, d_228_v74_])
        d_248_v94_ = d_248_v94_
        hi1_ = d_223_v72_
        for d_249_i19_ in range(762, hi1_):
            (d_112_globalState_).f1 = d_197_v50_
            (d_112_globalState_).f1 = d_197_v50_
            d_223_v72_ = d_223_v72_
            d_250_v95_: _dafny.Array
            def lambda19_(d_251_v72_):
                def lambda20_(d_252_i20_):
                    return default__.safeDivisionInt(d_252_i20_, d_251_v72_)

                return lambda20_

            init11_ = lambda19_(d_223_v72_)
            nw37_ = _dafny.Array(None, 25)
            for i0_11_ in range(nw37_.length(0)):
                nw37_[i0_11_] = init11_(i0_11_)
            d_250_v95_ = nw37_
            index37_ = default__.safeIndex(670, (d_250_v95_).length(0))
            (d_250_v95_)[index37_] = (0) - (d_249_i19_)
            index38_ = default__.safeIndex(670, (d_250_v95_).length(0))
            (d_250_v95_)[index38_] = d_249_i19_
        hi2_ = d_223_v72_
        for d_253_i21_ in range(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))), hi2_):
            d_254_v96_: _dafny.MultiSet
            d_254_v96_ = _dafny.MultiSet([d_253_i21_])
            d_223_v72_ = (((d_254_v96_)[247] if (247) in (d_254_v96_) else d_253_i21_)) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_255_i22_ in range(default__.abs(92))]))))
            d_256_v97_: str
            d_256_v97_ = _dafny.CodePoint('f')
            d_256_v97_ = d_256_v97_
            if False:
                d_257_v98_: _dafny.Array
                nw38_ = _dafny.Array(int(0), 10)
                d_257_v98_ = nw38_
                index39_ = default__.safeIndex(273, (d_257_v98_).length(0))
                (d_257_v98_)[index39_] = (default__.fm1(-790, d_112_globalState_)) - (d_223_v72_)
                d_258_v99_: _dafny.Seq
                d_258_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ensayrry"))
                index40_ = default__.safeIndex(273, (d_257_v98_).length(0))
                (d_257_v98_)[index40_] = len(((default__.fm3(d_112_globalState_)) + (d_258_v99_)) + (d_258_v99_))
                d_259_v100_: D4
                d_259_v100_ = D4_DC8(_dafny.Map({default__.fm0(d_223_v72_, d_112_globalState_): (d_257_v98_)[default__.safeIndex(273, (d_257_v98_).length(0))]}))
                d_260_v101_: _dafny.Map
                d_260_v101_ = _dafny.Map({False: (d_257_v98_)[default__.safeIndex(273, (d_257_v98_).length(0))]})
                d_261_v102_: _dafny.Map
                d_261_v102_ = _dafny.Map({((d_259_v100_).cf12) | (d_260_v101_): d_256_v97_})
                d_261_v102_ = ((d_261_v102_).set(d_260_v101_, d_256_v97_)) | (d_261_v102_)
                d_256_v97_ = _dafny.CodePoint('k')
                d_262_v103_: C0
                nw39_ = C0()
                nw39_.ctor__()
                d_262_v103_ = nw39_
                d_263_v104_: _dafny.Map
                d_263_v104_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([False]): d_223_v72_})
                index41_ = default__.safeIndex(273, (d_257_v98_).length(0))
                (d_257_v98_)[index41_] = ((0) - (-723)) - (default__.safeDivisionInt(len(d_263_v104_), (d_257_v98_)[default__.safeIndex(273, (d_257_v98_).length(0))]))
            elif True:
                d_264_v105_: C0
                nw40_ = C0()
                nw40_.ctor__()
                d_264_v105_ = nw40_
                d_197_v50_ = True
                d_265_v106_: _dafny.Seq
                d_265_v106_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxn"))
                d_265_v106_ = d_265_v106_
                d_266_v107_: _dafny.MultiSet
                d_266_v107_ = _dafny.MultiSet([d_197_v50_])
                d_267_v108_: D5
                d_267_v108_ = D5_DC10(d_266_v107_)
                (d_112_globalState_).f1 = ((d_266_v107_) | ((d_267_v108_).cf16)).ispropersubset((d_266_v107_) - (d_266_v107_))
                d_265_v106_ = (((d_265_v106_).set(default__.safeIndex(d_253_i21_, len(d_265_v106_)), d_256_v97_)) + (default__.fm3(d_112_globalState_))) + (d_265_v106_)
            d_268_v109_: _dafny.Set
            d_268_v109_ = _dafny.Set({d_223_v72_})
            d_269_v110_: _dafny.Seq
            d_269_v110_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybjqinu"))
            d_223_v72_ = ((len(d_268_v109_) if d_197_v50_ else len(d_269_v110_))) + ((d_223_v72_) - (d_223_v72_))
        d_270_v111_: _dafny.MultiSet
        d_270_v111_ = _dafny.MultiSet([d_197_v50_])
        (d_196_v49_).m1((d_270_v111_).cardinality, d_197_v50_, d_112_globalState_)
        (d_228_v74_).m1(d_223_v72_, d_197_v50_, d_112_globalState_)
        hi3_ = 285
        for d_271_i23_ in range(188, hi3_):
            d_272_v112_: _dafny.Seq
            d_272_v112_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imxv"))
            d_273_v113_: _dafny.Seq
            d_273_v113_ = _dafny.SeqWithoutIsStrInference([d_271_i23_, (0) - (d_271_i23_), (d_223_v72_) * (default__.fm1(len(d_272_v112_), d_112_globalState_))])
            d_223_v72_ = (d_273_v113_)[default__.safeIndex(((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_274_i24_ in range(default__.abs(612))])))) + (default__.fm1(d_271_i23_, d_112_globalState_)), len(d_273_v113_))]
            d_224_v73_ = d_224_v73_
            if ((d_223_v72_) * (-552)) != (len(default__.fm7(d_112_globalState_))):
                d_275_v114_: D2
                d_275_v114_ = D2_DC5((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_223_v72_ for d_276_i25_ in range(default__.abs(-94))]))).cardinality, 902)
                d_277_v115_: _dafny.MultiSet
                d_277_v115_ = _dafny.MultiSet([d_275_v114_, d_275_v114_])
                d_278_v116_: _dafny.Set
                d_278_v116_ = _dafny.Set({d_271_i23_, 914})
                d_223_v72_ = ((d_277_v115_)[d_275_v114_] if (d_275_v114_) in (d_277_v115_) else (d_223_v72_) * (len(d_278_v116_)))
                d_279_v117_: str
                d_279_v117_ = _dafny.CodePoint('b')
                d_280_v118_: _dafny.MultiSet
                d_280_v118_ = _dafny.MultiSet([d_279_v117_, _dafny.CodePoint('p')])
                d_281_v119_: _dafny.Seq
                d_281_v119_ = _dafny.SeqWithoutIsStrInference([d_197_v50_, default__.fm0(d_271_i23_, d_112_globalState_)])
                d_282_v120_: _dafny.Array
                nw41_ = _dafny.Array(None, 5)
                nw41_[int(0)] = d_223_v72_
                nw41_[int(1)] = (d_273_v113_)[default__.safeIndex(((d_280_v118_)[d_279_v117_] if (d_279_v117_) in (d_280_v118_) else len(d_281_v119_)), len(d_273_v113_))]
                nw41_[int(2)] = d_223_v72_
                nw41_[int(3)] = default__.safeModuloInt(d_223_v72_, len(d_281_v119_))
                nw41_[int(4)] = (154 if d_197_v50_ else d_271_i23_)
                d_282_v120_ = nw41_
                index42_ = default__.safeIndex(621, (d_282_v120_).length(0))
                (d_282_v120_)[index42_] = d_223_v72_
                index43_ = default__.safeIndex(621, (d_282_v120_).length(0))
                (d_282_v120_)[index43_] = d_271_i23_
                (d_112_globalState_).f1 = d_197_v50_
                d_283_v121_: _dafny.Map
                d_283_v121_ = _dafny.Map({d_197_v50_: True})
                d_284_v122_: _dafny.Array
                def lambda21_(d_285_i26_):
                    return _dafny.CodePoint('e')

                init12_ = lambda21_
                nw42_ = _dafny.Array(None, 17)
                for i0_12_ in range(nw42_.length(0)):
                    nw42_[i0_12_] = init12_(i0_12_)
                d_284_v122_ = nw42_
                index44_ = default__.safeIndex(120, (d_284_v122_).length(0))
                (d_284_v122_)[index44_] = d_279_v117_
                d_286_v123_: _dafny.Array
                def lambda22_(d_287_v111_):
                    def lambda23_(d_288_i27_):
                        return d_287_v111_

                    return lambda23_

                init13_ = lambda22_(d_270_v111_)
                nw43_ = _dafny.Array(None, 9)
                for i0_13_ in range(nw43_.length(0)):
                    nw43_[i0_13_] = init13_(i0_13_)
                d_286_v123_ = nw43_
                d_289_v124_: D5
                d_289_v124_ = D5_DC10(d_270_v111_)
                index45_ = default__.safeIndex(621, (d_282_v120_).length(0))
                index46_ = default__.safeIndex(120, (d_284_v122_).length(0))
                rhs47_ = (d_270_v111_) | ((d_289_v124_).cf16)
                rhs48_ = d_283_v121_
                rhs49_ = (0) - ((0) - (d_271_i23_))
                rhs50_ = d_279_v117_
                rhs51_ = d_286_v123_
                lhs20_ = d_282_v120_
                lhs21_ = default__.safeIndex(621, (d_282_v120_).length(0))
                lhs22_ = d_284_v122_
                lhs23_ = default__.safeIndex(120, (d_284_v122_).length(0))
                d_270_v111_ = rhs47_
                d_283_v121_ = rhs48_
                lhs20_[lhs21_] = rhs49_
                lhs22_[lhs23_] = rhs50_
                d_286_v123_ = rhs51_
                (d_112_globalState_).f1 = default__.fm0(default__.fm1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehbbvaf"))), d_112_globalState_), d_112_globalState_)
            elif True:
                d_223_v72_ = d_223_v72_
                d_290_v125_: D0
                d_290_v125_ = D0_DC1(d_271_i23_, d_223_v72_, (d_223_v72_) not in (_dafny.Map({d_271_i23_: d_197_v50_})))
                d_290_v125_ = d_290_v125_
                d_291_v127_: _dafny.Array
                def lambda24_(d_292_v72_):
                    def lambda25_(d_293_i28_):
                        def iife12_():
                            coll8_ = _dafny.Map()
                            compr_8_: int
                            for compr_8_ in _dafny.IntegerRange(374, 471):
                                d_294_v126_: int = compr_8_
                                if ((374) <= (d_294_v126_)) and ((d_294_v126_) < (471)):
                                    coll8_[(d_294_v126_) * (-80)] = d_292_v72_
                            return _dafny.Map(coll8_)
                        return default__.safeModuloInt(d_293_i28_, len(iife12_()
                        ))

                    return lambda25_

                init14_ = lambda24_(d_223_v72_)
                nw44_ = _dafny.Array(None, 24)
                for i0_14_ in range(nw44_.length(0)):
                    nw44_[i0_14_] = init14_(i0_14_)
                d_291_v127_ = nw44_
                index47_ = default__.safeIndex(907, (d_291_v127_).length(0))
                (d_291_v127_)[index47_] = d_271_i23_
                d_295_v128_: str
                d_295_v128_ = _dafny.CodePoint('o')
                d_296_v129_: _dafny.Seq
                d_296_v129_ = _dafny.SeqWithoutIsStrInference([(d_272_v112_).set(default__.safeIndex(525, len(d_272_v112_)), d_295_v128_)])
                d_297_v130_: _dafny.Array
                nw45_ = _dafny.Array(None, 18)
                nw45_[int(0)] = d_272_v112_
                nw45_[int(1)] = d_272_v112_
                nw45_[int(2)] = d_272_v112_
                nw45_[int(3)] = (d_272_v112_) + (d_272_v112_)
                nw45_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cscllebh"))
                nw45_[int(5)] = d_272_v112_
                nw45_[int(6)] = (((d_272_v112_).set(default__.safeIndex(d_271_i23_, len(d_272_v112_)), d_295_v128_)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_298_i29_ in range(default__.abs(860))]))).set(default__.safeIndex(d_271_i23_, len(((d_272_v112_).set(default__.safeIndex(d_271_i23_, len(d_272_v112_)), d_295_v128_)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_299_i29_ in range(default__.abs(860))])))), d_295_v128_)
                nw45_[int(7)] = (d_296_v129_)[default__.safeIndex((d_273_v113_)[default__.safeIndex(d_271_i23_, len(d_273_v113_))], len(d_296_v129_))]
                nw45_[int(8)] = d_272_v112_
                nw45_[int(9)] = (d_272_v112_) + (d_272_v112_)
                nw45_[int(10)] = (d_272_v112_) + (d_272_v112_)
                nw45_[int(11)] = d_272_v112_
                nw45_[int(12)] = d_272_v112_
                nw45_[int(13)] = d_272_v112_
                nw45_[int(14)] = d_272_v112_
                nw45_[int(15)] = d_272_v112_
                nw45_[int(16)] = (d_272_v112_ if False else d_272_v112_)
                nw45_[int(17)] = d_272_v112_
                d_297_v130_ = nw45_
                d_300_v131_: _dafny.Map
                d_300_v131_ = _dafny.Map({d_271_i23_: d_223_v72_})
                d_301_v132_: _dafny.Map
                d_301_v132_ = _dafny.Map({d_197_v50_: d_228_v74_})
                d_302_v133_: _dafny.MultiSet
                d_302_v133_ = _dafny.MultiSet([(0) - (((d_300_v131_)[d_271_i23_] if (d_271_i23_) in (d_300_v131_) else len(d_301_v132_)))])
                index48_ = default__.safeIndex(907, (d_291_v127_).length(0))
                rhs52_ = not(d_197_v50_)
                rhs53_ = (d_223_v72_) * (((d_302_v133_) | (d_302_v133_)).cardinality)
                rhs54_ = d_297_v130_
                rhs55_ = (0) - ((d_271_i23_) + (d_271_i23_))
                lhs24_ = d_291_v127_
                lhs25_ = default__.safeIndex(907, (d_291_v127_).length(0))
                d_197_v50_ = rhs52_
                lhs24_[lhs25_] = rhs53_
                d_297_v130_ = rhs54_
                d_223_v72_ = rhs55_
                d_291_v127_ = d_291_v127_
                d_197_v50_ = not (d_197_v50_) or (default__.fm0(340, d_112_globalState_))
            (d_112_globalState_).f1 = d_197_v50_
        if not(default__.fm0(default__.safeModuloInt(d_223_v72_, d_223_v72_), d_112_globalState_)):
            d_303_v134_: _dafny.Seq
            d_303_v134_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hceiynyhv"))
            d_303_v134_ = d_303_v134_
            d_304_v135_: C0
            nw46_ = C0()
            nw46_.ctor__()
            d_304_v135_ = nw46_
            d_305_v136_: _dafny.Array
            nw47_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_305_v136_ = nw47_
            d_306_v137_: _dafny.MultiSet
            d_306_v137_ = _dafny.MultiSet([d_223_v72_, (0) - (d_223_v72_), len(d_303_v134_)])
            d_307_v138_: _dafny.Map
            d_307_v138_ = _dafny.Map({d_197_v50_: d_223_v72_})
            d_308_v139_: _dafny.Map
            d_308_v139_ = _dafny.Map({(d_305_v136_ if True else d_305_v136_): ((d_306_v137_) | ((d_306_v137_).set(d_223_v72_, default__.abs(len(d_307_v138_))))).cardinality})
            d_308_v139_ = (d_308_v139_).set(d_305_v136_, d_223_v72_)
            (d_112_globalState_).f1 = d_197_v50_
            d_309_v140_: _dafny.Map
            d_309_v140_ = _dafny.Map({d_223_v72_: 573})
            d_310_v141_: _dafny.MultiSet
            d_310_v141_ = _dafny.MultiSet([d_309_v140_, d_309_v140_])
            d_311_v142_: _dafny.Map
            d_311_v142_ = _dafny.Map({d_197_v50_: d_197_v50_})
            if not (((d_311_v142_)[False] if (False) in (d_311_v142_) else d_197_v50_)) or ((d_310_v141_).issubset(d_310_v141_)):
                d_312_v143_: C0
                nw48_ = C0()
                nw48_.ctor__()
                d_312_v143_ = nw48_
                (d_112_globalState_).f1 = (d_197_v50_ if d_197_v50_ else d_197_v50_)
                d_313_v144_: _dafny.Seq
                d_313_v144_ = _dafny.SeqWithoutIsStrInference([d_223_v72_])
                index49_ = default__.safeIndex(244, (d_113_v0_).length(0))
                (d_113_v0_)[index49_] = (d_313_v144_) == (d_313_v144_)
                d_314_v145_: D0
                d_314_v145_ = D0_DC1(d_223_v72_, len(_dafny.Map({-55: d_223_v72_})), d_197_v50_)
                d_315_v146_: _dafny.Map
                d_315_v146_ = _dafny.Map({default__.fm0(-268, d_112_globalState_): d_313_v144_})
                index50_ = default__.safeIndex(244, (d_113_v0_).length(0))
                (d_113_v0_)[index50_] = (default__.fm1(d_223_v72_, d_112_globalState_)) in ((_dafny.SeqWithoutIsStrInference([d_223_v72_, (0) - ((default__.fm5(d_197_v50_, d_197_v50_, d_314_v145_, d_112_globalState_)).cardinality)])) + (((d_315_v146_)[d_197_v50_] if (d_197_v50_) in (d_315_v146_) else d_313_v144_)))
                d_197_v50_ = (d_197_v50_) and (True)
                d_223_v72_ = default__.safeModuloInt((d_223_v72_) + (d_223_v72_), d_223_v72_)
            elif True:
                d_223_v72_ = ((len(d_303_v134_)) - (d_223_v72_)) * (d_223_v72_)
                d_316_v147_: _dafny.Array
                nw49_ = _dafny.Array(_dafny.CodePoint('D'), 25)
                d_316_v147_ = nw49_
                d_317_v148_: _dafny.Array
                nw50_ = _dafny.Array(int(0), 3)
                d_317_v148_ = nw50_
                d_318_v149_: _dafny.Map
                d_318_v149_ = _dafny.Map({d_113_v0_: d_316_v147_})
                d_319_v150_: _dafny.Seq
                d_319_v150_ = _dafny.SeqWithoutIsStrInference([((d_318_v149_)[d_113_v0_] if (d_113_v0_) in (d_318_v149_) else d_316_v147_), d_316_v147_, d_316_v147_, d_316_v147_])
                rhs56_ = (d_319_v150_)[default__.safeIndex(default__.fm1(d_223_v72_, d_112_globalState_), len(d_319_v150_))]
                rhs57_ = d_223_v72_
                rhs58_ = not ((d_197_v50_) and (d_197_v50_)) or (d_197_v50_)
                rhs59_ = d_317_v148_
                lhs26_ = d_112_globalState_
                d_316_v147_ = rhs56_
                d_223_v72_ = rhs57_
                lhs26_.f1 = rhs58_
                d_317_v148_ = rhs59_
                d_320_v151_: _dafny.Map
                d_320_v151_ = _dafny.Map({len(_dafny.Set({d_197_v50_})): d_197_v50_})
                (d_112_globalState_).f1 = default__.fm0((d_223_v72_) - (len(d_320_v151_)), d_112_globalState_)
                d_321_v152_: str
                d_321_v152_ = _dafny.CodePoint('h')
                d_322_v153_: _dafny.Map
                d_322_v153_ = _dafny.Map({(default__.fm8(d_197_v50_, _dafny.MultiSet([d_321_v152_]), d_112_globalState_)) + (_dafny.SeqWithoutIsStrInference([d_223_v72_, d_223_v72_])): d_304_v135_})
                d_323_v154_: _dafny.Seq
                d_323_v154_ = _dafny.SeqWithoutIsStrInference([d_223_v72_, d_223_v72_])
                d_304_v135_ = ((d_322_v153_)[d_323_v154_] if (d_323_v154_) in (d_322_v153_) else (d_304_v135_ if d_197_v50_ else d_304_v135_))
                d_324_v155_: _dafny.Array
                nw51_ = _dafny.Array(None, 8)
                d_324_v155_ = nw51_
                d_324_v155_ = d_324_v155_
        elif True:
            (d_112_globalState_).f1 = d_197_v50_
            d_325_v156_: _dafny.Seq
            d_325_v156_ = _dafny.SeqWithoutIsStrInference([d_223_v72_])
            d_326_v157_: _dafny.Map
            d_326_v157_ = _dafny.Map({(len((d_325_v156_).set(default__.safeIndex((d_270_v111_).cardinality, len(d_325_v156_)), d_223_v72_))) - (d_223_v72_): d_197_v50_})
            (d_112_globalState_).f1 = ((d_326_v157_)[d_223_v72_] if (d_223_v72_) in (d_326_v157_) else d_197_v50_)
            d_327_v158_: _dafny.Seq
            d_327_v158_ = _dafny.SeqWithoutIsStrInference([d_197_v50_, d_197_v50_])
            d_327_v158_ = _dafny.SeqWithoutIsStrInference([d_197_v50_, d_197_v50_])
            (d_112_globalState_).f1 = ((d_270_v111_).intersection(d_270_v111_)).ispropersubset(d_270_v111_)
            d_328_v159_: _dafny.Array
            nw52_ = _dafny.Array(int(0), 8)
            d_328_v159_ = nw52_
            d_329_v160_: _dafny.MultiSet
            d_329_v160_ = _dafny.MultiSet([d_328_v159_])
            d_330_v161_: _dafny.MultiSet
            d_330_v161_ = d_329_v160_
            d_331_v162_: _dafny.Map
            d_331_v162_ = _dafny.Map({d_223_v72_: (((d_330_v161_)).intersection(d_329_v160_)).cardinality})
            d_332_v163_: str
            d_332_v163_ = _dafny.CodePoint('p')
            d_333_v164_: _dafny.Map
            d_333_v164_ = _dafny.Map({d_197_v50_: d_332_v163_})
            d_334_v165_: _dafny.Map
            d_334_v165_ = _dafny.Map({d_333_v164_: 133})
            d_331_v162_ = (d_331_v162_).set(d_223_v72_, ((d_334_v165_)[d_333_v164_] if (d_333_v164_) in (d_334_v165_) else d_223_v72_))
        d_335_v166_: _dafny.Seq
        d_335_v166_ = _dafny.SeqWithoutIsStrInference([True])
        d_197_v50_ = (d_335_v166_) <= (d_335_v166_)
        _dafny.print(_dafny.string_of((d_112_globalState_.f0) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_112_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v1_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_197_v50_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_223_v72_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_224_v73_)[0]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_224_v73_)[1]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_224_v73_)[2]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v94_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v111_) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_v166_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), False)
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
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({self.cf4.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC4(D2, NamedTuple('DC4', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

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
        return lambda: D4_DC9(_dafny.Array(None, 0), int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)

class D4_DC9(D4, NamedTuple('DC9', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D5_DC11(None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)

class D5_DC11(D5, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {self.cf21.VerbatimString(True)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC10(D5, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC13(D6, NamedTuple('DC13', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC15(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)

class D7_DC15(D7, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({self.cf25.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Seq = _dafny.Seq({})
        self.f1: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1):
        (self).f0 = f0
        (self).f1 = f1


class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def m1(self, p0, p1, globalState):
        hi4_ = p0
        for d_336_i0_ in range(p0, hi4_):
            d_337_v0_: int
            d_337_v0_ = -15
            d_338_v2_: str
            d_338_v2_ = _dafny.CodePoint('e')
            d_339_v3_: _dafny.Map
            d_339_v3_ = _dafny.Map({len(_dafny.Set({d_338_v2_})): p0})
            def iife13_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(666, -533):
                    d_340_v1_: int = compr_9_
                    if ((666) <= (d_340_v1_)) and ((d_340_v1_) < (-533)):
                        coll9_[default__.safeModuloInt(d_340_v1_, d_337_v0_)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dchcbikbk"))))
                return _dafny.Map(coll9_)
            d_337_v0_ = (d_337_v0_) * (len((iife13_()
            ) | (d_339_v3_)))
            d_341_v4_: _dafny.Seq
            d_341_v4_ = _dafny.SeqWithoutIsStrInference([d_336_i0_])
            d_342_v5_: _dafny.Seq
            d_342_v5_ = _dafny.SeqWithoutIsStrInference([d_341_v4_])
            d_343_v6_: _dafny.Map
            d_343_v6_ = _dafny.Map({d_337_v0_: not(p1)})
            (globalState).f0 = (d_342_v5_)[default__.safeIndex(len((_dafny.Map({p0: p1})) | (d_343_v6_)), len(d_342_v5_))]
            (globalState).f1 = (len((_dafny.SeqWithoutIsStrInference([p1])).set(default__.safeIndex(d_336_i0_, len(_dafny.SeqWithoutIsStrInference([p1]))), p1))) > ((D0_DC1(d_337_v0_, default__.fm1(d_336_i0_, globalState), p1)).cf1)
            (globalState).f1 = p1
        (globalState).f1 = p1
        d_344_v7_: int
        d_344_v7_ = 436
        d_345_v8_: _dafny.MultiSet
        d_345_v8_ = _dafny.MultiSet([p1, p1])
        d_344_v7_ = default__.safeDivisionInt(default__.fm1(p0, globalState), (803) + ((d_345_v8_).cardinality))
        d_346_v9_: _dafny.Seq
        d_346_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hb"))
        d_347_v10_: _dafny.Seq
        d_347_v10_ = _dafny.SeqWithoutIsStrInference([p1])
        d_348_v11_: _dafny.Map
        d_348_v11_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([default__.fm0(d_344_v7_, globalState)])): p1})
        d_349_v12_: _dafny.Array
        nw53_ = _dafny.Array(None, 20)
        nw53_[int(0)] = _dafny.SeqWithoutIsStrInference([True, p1])
        nw53_[int(1)] = d_347_v10_
        nw53_[int(2)] = ((d_347_v10_).set(default__.safeIndex(d_344_v7_, len(d_347_v10_)), p1)) + (d_347_v10_)
        nw53_[int(3)] = (d_347_v10_) + (d_347_v10_)
        nw53_[int(4)] = _dafny.SeqWithoutIsStrInference([p1, p1])
        nw53_[int(5)] = d_347_v10_
        nw53_[int(6)] = _dafny.SeqWithoutIsStrInference([p1, p1])
        nw53_[int(7)] = (d_347_v10_).set(default__.safeIndex(312, len(d_347_v10_)), p1)
        nw53_[int(8)] = d_347_v10_
        nw53_[int(9)] = d_347_v10_
        nw53_[int(10)] = (d_347_v10_) + (d_347_v10_)
        nw53_[int(11)] = d_347_v10_
        nw53_[int(12)] = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
        nw53_[int(13)] = d_347_v10_
        nw53_[int(14)] = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
        nw53_[int(15)] = default__.fm2(p1, p1, d_348_v11_, p1, globalState)
        nw53_[int(16)] = d_347_v10_
        nw53_[int(17)] = (d_347_v10_ if p1 else d_347_v10_)
        nw53_[int(18)] = d_347_v10_
        nw53_[int(19)] = _dafny.SeqWithoutIsStrInference([True, p1])
        d_349_v12_ = nw53_
        index51_ = default__.safeIndex(468, (d_349_v12_).length(0))
        (d_349_v12_)[index51_] = (d_347_v10_) + (d_347_v10_)
        index52_ = default__.safeIndex(468, (d_349_v12_).length(0))
        rhs60_ = (default__.fm3(globalState)) + (d_346_v9_)
        rhs61_ = d_347_v10_
        rhs62_ = (default__.fm4(globalState)) < (_dafny.SeqWithoutIsStrInference([(d_346_v9_) for d_350_i1_ in range(default__.abs(857))]))
        rhs63_ = p1
        lhs27_ = d_349_v12_
        lhs28_ = default__.safeIndex(468, (d_349_v12_).length(0))
        lhs29_ = globalState
        lhs30_ = globalState
        d_346_v9_ = rhs60_
        lhs27_[lhs28_] = rhs61_
        lhs29_.f1 = rhs62_
        lhs30_.f1 = rhs63_
        d_344_v7_ = d_344_v7_
        d_351_v13_: _dafny.Map
        d_351_v13_ = _dafny.Map({d_344_v7_: d_344_v7_})
        d_352_v14_: _dafny.Map
        d_352_v14_ = _dafny.Map({(p0) == (((d_351_v13_)[d_344_v7_] if (d_344_v7_) in (d_351_v13_) else p0)): -600})
        rhs64_ = ((p0) + ((0) - (len(d_351_v13_)))) * (p0)
        rhs65_ = ((d_352_v14_)[p1] if (p1) in (d_352_v14_) else (0) - (d_344_v7_))
        d_344_v7_ = rhs64_
        d_344_v7_ = rhs65_

