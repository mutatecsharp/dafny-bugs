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
            for compr_0_ in (_dafny.Set({(0) - (-229), len(_dafny.SeqWithoutIsStrInference([True]))})).Elements:
                d_0_v0_: int = compr_0_
                if (d_0_v0_) in (_dafny.Set({(0) - (-229), len(_dafny.SeqWithoutIsStrInference([True]))})):
                    coll0_[default__.safeModuloInt(d_0_v0_, -693)] = len(_dafny.Map({-902: _dafny.CodePoint('d')}))
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(430, 54):
                d_1_v1_: int = compr_1_
                if ((430) <= (d_1_v1_)) and ((d_1_v1_) < (54)):
                    coll1_ = coll1_.union(_dafny.Set([default__.safeModuloInt(d_1_v1_, 738)]))
            return _dafny.Set(coll1_)
        return (0) - ((0) - (((581 if not(not(not(False))) else len(_dafny.Map({-3: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhl")))]))}))) if True else ((_dafny.MultiSet([981, -818])) - (_dafny.MultiSet([len(iife0_()
        ), (_dafny.MultiSet([(0) - (len(iife1_()
        ))])).cardinality]))).cardinality)))

    @staticmethod
    def fm1(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2_i0_ in range(default__.abs(669))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3_i1_ in range(default__.abs(681))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjfujd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnhbwhio"))))

    @staticmethod
    def fm2(p0, globalState):
        return D0_DC1(977, False, -930)

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return (_dafny.MultiSet([True, not(False)])) | ((_dafny.MultiSet([False, True, False])) - (_dafny.MultiSet([True])))

    @staticmethod
    def fm4(p0, p1, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([953, len(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference([True, True]))})), -771])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))))])))

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txsljff"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cuu")) for d_4_i0_ in range(default__.abs(302))]))

    @staticmethod
    def fm6(p0, p1, globalState):
        source0_ = D0_DC1(369, False, 448)
        if source0_.is_DC1:
            d_5___mcc_h0_ = source0_.cf1
            d_6___mcc_h1_ = source0_.cf2
            d_7___mcc_h2_ = source0_.cf3
            d_8_cf3_ = d_7___mcc_h2_
            d_9_cf2_ = d_6___mcc_h1_
            d_10_cf1_ = d_5___mcc_h0_
            return (d_8_cf3_) < (d_10_cf1_)
        elif source0_.is_DC2:
            return (False) == (True)
        elif source0_.is_DC3:
            return not((True) not in (_dafny.Map({not(False): not(False)})))
        elif True:
            d_11___mcc_h3_ = source0_.cf0
            d_12_cf0_ = d_11___mcc_h3_
            return ((0) - (d_12_cf0_)) > (917)

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({_dafny.CodePoint('h')})) | (_dafny.Set({_dafny.CodePoint('q')}))) | ((D11_DC29(_dafny.Set({_dafny.CodePoint('m'), _dafny.CodePoint('u')}))).cf39)

    @staticmethod
    def fm8(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([437 for d_13_i0_ in range(default__.abs(636))])) + (_dafny.SeqWithoutIsStrInference([-774]))

    @staticmethod
    def fm9(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: str
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')])).Elements:
                d_14_v0_: str = compr_2_
                if (d_14_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')])):
                    coll2_[d_14_v0_] = D9_DC23(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtgcdb"))))
            return _dafny.Map(coll2_)
        return (_dafny.Set({(_dafny.MultiSet([not(False)])).cardinality, 768, 863})) - (_dafny.Set({len(iife2_()
        )}))

    @staticmethod
    def fm10(globalState):
        return (_dafny.SeqWithoutIsStrInference([False])) + ((_dafny.SeqWithoutIsStrInference([True, not(False), True])) + (_dafny.SeqWithoutIsStrInference([True, not(True)])))

    @staticmethod
    def m0(p0, globalState):
        d_15_v0_: _dafny.Array
        def lambda0_(d_16_i0_):
            return default__.safeDivisionInt(d_16_i0_, -52)

        init0_ = lambda0_
        nw0_ = _dafny.Array(None, 28)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_15_v0_ = nw0_
        d_17_v1_: int
        d_17_v1_ = 473
        index0_ = default__.safeIndex(723, (d_15_v0_).length(0))
        (d_15_v0_)[index0_] = default__.fm0(d_17_v1_, d_17_v1_, globalState)
        d_18_v2_: _dafny.Seq
        d_18_v2_ = _dafny.SeqWithoutIsStrInference([p0])
        d_19_v3_: _dafny.Seq
        d_19_v3_ = _dafny.SeqWithoutIsStrInference([d_17_v1_, len(d_18_v2_)])
        d_20_v4_: _dafny.Seq
        d_20_v4_ = _dafny.SeqWithoutIsStrInference([d_19_v3_])
        index1_ = default__.safeIndex(723, (d_15_v0_).length(0))
        (d_15_v0_)[index1_] = ((0) - (default__.safeDivisionInt(d_17_v1_, d_17_v1_))) + (len(d_20_v4_))
        d_21_i1_: int
        d_21_i1_ = 0
        with _dafny.label("0"):
            while (d_18_v2_)[default__.safeIndex((d_17_v1_) - (d_17_v1_), len(d_18_v2_))]:
                with _dafny.c_label("0"):
                    if (d_21_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_21_i1_ = (d_21_i1_) + (1)
                    d_22_v5_: _dafny.Array
                    def lambda1_(d_23_i2_):
                        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvewhxik"))

                    init1_ = lambda1_
                    nw1_ = _dafny.Array(None, 26)
                    for i0_1_ in range(nw1_.length(0)):
                        nw1_[i0_1_] = init1_(i0_1_)
                    d_22_v5_ = nw1_
                    d_24_v6_: _dafny.Seq
                    d_24_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqwfqaq"))
                    index2_ = default__.safeIndex(542, (d_22_v5_).length(0))
                    (d_22_v5_)[index2_] = (d_24_v6_) + (d_24_v6_)
                    d_25_v7_: _dafny.Array
                    nw2_ = _dafny.Array(_dafny.Map({}), 2)
                    d_25_v7_ = nw2_
                    d_26_v8_: D0
                    d_26_v8_ = D0_DC0(d_17_v1_)
                    d_27_v9_: _dafny.MultiSet
                    d_27_v9_ = _dafny.MultiSet([p0, p0, p0, True, p0])
                    index3_ = default__.safeIndex(907, (d_25_v7_).length(0))
                    (d_25_v7_)[index3_] = _dafny.Map({d_26_v8_: d_27_v9_})
                    d_28_v10_: _dafny.MultiSet
                    d_28_v10_ = _dafny.MultiSet([607])
                    d_29_v11_: _dafny.Map
                    d_29_v11_ = _dafny.Map({d_26_v8_: default__.fm3(p0, d_28_v10_, p0, globalState)})
                    index4_ = default__.safeIndex(542, (d_22_v5_).length(0))
                    index5_ = default__.safeIndex(907, (d_25_v7_).length(0))
                    rhs0_ = d_24_v6_
                    rhs1_ = (106) + (default__.safeDivisionInt((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]))
                    rhs2_ = (d_29_v11_) | (d_29_v11_)
                    rhs3_ = not((d_27_v9_).ispropersubset(d_27_v9_))
                    lhs0_ = d_22_v5_
                    lhs1_ = default__.safeIndex(542, (d_22_v5_).length(0))
                    lhs2_ = globalState
                    lhs3_ = d_25_v7_
                    lhs4_ = default__.safeIndex(907, (d_25_v7_).length(0))
                    lhs5_ = globalState
                    lhs0_[lhs1_] = rhs0_
                    lhs2_.f4 = rhs1_
                    lhs3_[lhs4_] = rhs2_
                    lhs5_.f2 = rhs3_
                    index6_ = default__.safeIndex(723, (d_15_v0_).length(0))
                    (d_15_v0_)[index6_] = d_17_v1_
                    d_30_v13_: _dafny.Array
                    def lambda2_(d_31_v1_, d_32_v10_, d_33_p0_, d_34_v0_):
                        def lambda3_(d_35_i3_):
                            def iife3_():
                                coll3_ = _dafny.Set()
                                compr_3_: int
                                for compr_3_ in _dafny.IntegerRange(671, -131):
                                    d_36_v12_: int = compr_3_
                                    if ((671) <= (d_36_v12_)) and ((d_36_v12_) < (-131)):
                                        coll3_ = coll3_.union(_dafny.Set([(d_36_v12_) * ((D0_DC1((d_32_v10_).cardinality, d_33_p0_, (d_34_v0_)[default__.safeIndex(723, (d_34_v0_).length(0))])).cf3)]))
                                return _dafny.Set(coll3_)
                            return (iife3_()
                            ).isdisjoint(_dafny.Set({d_31_v1_}))

                        return lambda3_

                    init2_ = lambda2_(d_17_v1_, d_28_v10_, p0, d_15_v0_)
                    nw3_ = _dafny.Array(None, 23)
                    for i0_2_ in range(nw3_.length(0)):
                        nw3_[i0_2_] = init2_(i0_2_)
                    d_30_v13_ = nw3_
                    index7_ = default__.safeIndex(484, (d_30_v13_).length(0))
                    (d_30_v13_)[index7_] = True
                    index8_ = default__.safeIndex(484, (d_30_v13_).length(0))
                    (d_30_v13_)[index8_] = not(default__.fm6((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_37_i4_ in range(default__.abs(698))])) != ((d_22_v5_)[default__.safeIndex(542, (d_22_v5_).length(0))]), (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], globalState))
                    d_38_v14_: D0
                    d_38_v14_ = D0_DC1(d_17_v1_, p0, (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))])
                    d_39_v15_: _dafny.Map
                    d_39_v15_ = _dafny.Map({d_38_v14_: (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]})
                    d_40_v16_: C0
                    nw4_ = C0()
                    nw4_.ctor__(d_39_v15_)
                    d_40_v16_ = nw4_
                    pass
            pass
        hi0_ = default__.safeDivisionInt((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))])
        for d_41_i5_ in range((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], hi0_):
            d_42_v17_: _dafny.Set
            d_42_v17_ = _dafny.Set({p0, p0})
            d_43_v18_: str
            d_43_v18_ = _dafny.CodePoint('y')
            d_44_v19_: _dafny.Map
            d_44_v19_ = _dafny.Map({d_42_v17_: d_43_v18_})
            d_45_v20_: D9
            d_45_v20_ = D9_DC22(d_42_v17_)
            d_44_v19_ = (d_44_v19_).set((d_45_v20_).cf25, d_43_v18_)
            d_46_v21_: D0
            d_46_v21_ = D0_DC1(d_41_i5_, p0, d_17_v1_)
            d_47_v22_: _dafny.Map
            d_47_v22_ = _dafny.Map({d_46_v21_: -29})
            d_48_v23_: C0
            nw5_ = C0()
            nw5_.ctor__(d_47_v22_)
            d_48_v23_ = nw5_
            (globalState).f3 = d_41_i5_
            d_49_v24_: _dafny.Seq
            d_49_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfxeijna"))
            d_50_v25_: _dafny.MultiSet
            d_50_v25_ = _dafny.MultiSet([d_49_v24_, d_49_v24_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mv")), d_49_v24_])
            d_51_v26_: _dafny.MultiSet
            d_51_v26_ = _dafny.MultiSet([d_41_i5_, (d_19_v3_)[default__.safeIndex((d_50_v25_).cardinality, len(d_19_v3_))]])
            d_52_v27_: _dafny.Map
            d_52_v27_ = _dafny.Map({(d_17_v1_) == (((d_51_v26_)[d_17_v1_] if (d_17_v1_) in (d_51_v26_) else d_17_v1_)): default__.safeModuloInt(d_41_i5_, d_17_v1_)})
            d_52_v27_ = (d_52_v27_).set(p0, d_17_v1_)
        if p0:
            d_53_v29_: _dafny.Map
            d_53_v29_ = _dafny.Map({d_17_v1_: p0})
            d_54_v30_: D0
            d_54_v30_ = D0_DC1((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], ((d_53_v29_)[(d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]] if ((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]) in (d_53_v29_) else p0), d_17_v1_)
            d_55_v31_: _dafny.Map
            d_55_v31_ = _dafny.Map({d_54_v30_: (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]})
            d_56_v32_: C0
            nw6_ = C0()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: D0
                for compr_4_ in (d_55_v31_).keys.Elements:
                    d_57_v28_: D0 = compr_4_
                    if (d_57_v28_) in (d_55_v31_):
                        coll4_[d_57_v28_] = d_17_v1_
                return _dafny.Map(coll4_)
            nw6_.ctor__(iife4_()
            )
            d_56_v32_ = nw6_
            d_58_v33_: str
            d_58_v33_ = _dafny.CodePoint('l')
            d_59_v34_: D9
            d_59_v34_ = D9_DC24(default__.fm0((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], globalState), d_56_v32_, d_58_v33_, d_17_v1_)
            d_60_v35_: _dafny.Map
            d_60_v35_ = _dafny.Map({d_59_v34_: (0) - (d_17_v1_)})
            d_61_v36_: _dafny.MultiSet
            d_61_v36_ = _dafny.MultiSet([p0])
            d_62_v37_: _dafny.Seq
            d_62_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaoscxjb"))
            d_60_v35_ = (d_60_v35_).set(d_59_v34_, default__.safeModuloInt(((d_61_v36_)[p0] if (p0) in (d_61_v36_) else len(d_62_v37_)), (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]))
            (globalState).f2 = p0
            (globalState).f2 = p0
            d_63_v38_: _dafny.Map
            d_63_v38_ = _dafny.Map({p0: d_58_v33_})
            d_64_v39_: D8
            d_64_v39_ = D8_DC20(d_58_v33_)
            d_65_v40_: _dafny.Array
            nw7_ = _dafny.Array(None, 26)
            nw7_[int(0)] = _dafny.CodePoint('q')
            nw7_[int(1)] = d_58_v33_
            nw7_[int(2)] = d_58_v33_
            nw7_[int(3)] = d_58_v33_
            nw7_[int(4)] = (D9_DC24((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], d_56_v32_, _dafny.CodePoint('g'), d_17_v1_)).cf29
            nw7_[int(5)] = d_58_v33_
            nw7_[int(6)] = d_58_v33_
            nw7_[int(7)] = d_58_v33_
            nw7_[int(8)] = _dafny.CodePoint('x')
            nw7_[int(9)] = ((d_63_v38_)[p0] if (p0) in (d_63_v38_) else d_58_v33_)
            nw7_[int(10)] = _dafny.CodePoint('s')
            nw7_[int(11)] = d_58_v33_
            nw7_[int(12)] = d_58_v33_
            nw7_[int(13)] = d_58_v33_
            nw7_[int(14)] = d_58_v33_
            nw7_[int(15)] = _dafny.CodePoint('w')
            nw7_[int(16)] = d_58_v33_
            nw7_[int(17)] = d_58_v33_
            nw7_[int(18)] = (d_64_v39_).cf22
            nw7_[int(19)] = d_58_v33_
            nw7_[int(20)] = d_58_v33_
            nw7_[int(21)] = d_58_v33_
            nw7_[int(22)] = d_58_v33_
            nw7_[int(23)] = _dafny.CodePoint('k')
            nw7_[int(24)] = d_58_v33_
            nw7_[int(25)] = d_58_v33_
            d_65_v40_ = nw7_
            index9_ = default__.safeIndex(167, (d_65_v40_).length(0))
            (d_65_v40_)[index9_] = d_58_v33_
            d_66_v41_: _dafny.Set
            d_66_v41_ = _dafny.Set({(d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]})
            index10_ = default__.safeIndex(723, (d_15_v0_).length(0))
            index11_ = default__.safeIndex(167, (d_65_v40_).length(0))
            rhs4_ = (-33 if p0 else len(d_66_v41_))
            rhs5_ = d_58_v33_
            rhs6_ = d_17_v1_
            lhs6_ = d_15_v0_
            lhs7_ = default__.safeIndex(723, (d_15_v0_).length(0))
            lhs8_ = d_65_v40_
            lhs9_ = default__.safeIndex(167, (d_65_v40_).length(0))
            lhs10_ = globalState
            lhs6_[lhs7_] = rhs4_
            lhs8_[lhs9_] = rhs5_
            lhs10_.f3 = rhs6_
            d_67_v42_: _dafny.Map
            d_67_v42_ = _dafny.Map({(d_62_v37_ if True else _dafny.SeqWithoutIsStrInference([(d_65_v40_)[default__.safeIndex(167, (d_65_v40_).length(0))] for d_68_i6_ in range(default__.abs(-146))])): ((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]) - (476)})
            d_67_v42_ = d_67_v42_
        elif True:
            if p0:
                d_69_v43_: _dafny.Seq
                d_69_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnu"))
                d_70_v44_: str
                d_70_v44_ = _dafny.CodePoint('n')
                d_71_v45_: D0
                d_71_v45_ = D0_DC1((343) - (d_17_v1_), (d_69_v43_) <= ((d_69_v43_).set(default__.safeIndex(335, len(d_69_v43_)), d_70_v44_)), len(d_18_v2_))
                d_71_v45_ = d_71_v45_
                (globalState).f4 = (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]
                d_72_v46_: D0
                d_72_v46_ = D0_DC3()
                d_73_v47_: _dafny.Map
                d_73_v47_ = _dafny.Map({d_72_v46_: d_15_v0_})
                d_73_v47_ = (d_73_v47_).set(d_72_v46_, d_15_v0_)
                d_74_v48_: _dafny.Map
                d_74_v48_ = _dafny.Map({p0: p0})
                d_75_v50_: _dafny.Map
                d_75_v50_ = _dafny.Map({d_71_v45_: d_19_v3_})
                d_76_v51_: C0
                nw8_ = C0()
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: D0
                    for compr_5_ in (d_75_v50_).keys.Elements:
                        d_77_v49_: D0 = compr_5_
                        if (d_77_v49_) in (d_75_v50_):
                            coll5_[d_77_v49_] = len(d_69_v43_)
                    return _dafny.Map(coll5_)
                nw8_.ctor__(iife5_()
                )
                d_76_v51_ = nw8_
                d_78_v52_: D10
                d_78_v52_ = D10_DC26(_dafny.Map({d_74_v48_: d_76_v51_}))
                d_17_v1_ = (len((d_78_v52_).cf31)) * (d_17_v1_)
                d_70_v44_ = d_70_v44_
            elif True:
                d_79_v53_: _dafny.Seq
                d_79_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqlo"))
                d_80_v54_: _dafny.Array
                def lambda4_(d_81_v2_):
                    def lambda5_(d_82_i7_):
                        return (d_81_v2_) + (d_81_v2_)

                    return lambda5_

                init3_ = lambda4_(d_18_v2_)
                nw9_ = _dafny.Array(None, 1)
                for i0_3_ in range(nw9_.length(0)):
                    nw9_[i0_3_] = init3_(i0_3_)
                d_80_v54_ = nw9_
                index12_ = default__.safeIndex(619, (d_80_v54_).length(0))
                (d_80_v54_)[index12_] = _dafny.SeqWithoutIsStrInference([not(True), p0, p0, p0, p0])
                index13_ = default__.safeIndex(619, (d_80_v54_).length(0))
                index14_ = default__.safeIndex(723, (d_15_v0_).length(0))
                rhs7_ = (d_79_v53_ if p0 else d_79_v53_)
                rhs8_ = (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]
                rhs9_ = d_18_v2_
                rhs10_ = (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]
                lhs11_ = globalState
                lhs12_ = d_80_v54_
                lhs13_ = default__.safeIndex(619, (d_80_v54_).length(0))
                lhs14_ = d_15_v0_
                lhs15_ = default__.safeIndex(723, (d_15_v0_).length(0))
                d_79_v53_ = rhs7_
                lhs11_.f4 = rhs8_
                lhs12_[lhs13_] = rhs9_
                lhs14_[lhs15_] = rhs10_
                d_83_v55_: _dafny.Map
                d_83_v55_ = _dafny.Map({d_17_v1_: (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]})
                d_84_v57_: _dafny.Seq
                d_84_v57_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({default__.fm0((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], 923, globalState): (_dafny.MultiSet([d_17_v1_, -135])).cardinality})])
                d_85_v58_: _dafny.MultiSet
                def iife6_():
                    coll6_ = _dafny.Map()
                    compr_6_: int
                    for compr_6_ in _dafny.IntegerRange(686, 758):
                        d_86_v56_: int = compr_6_
                        if ((686) <= (d_86_v56_)) and ((d_86_v56_) < (758)):
                            coll6_[(d_86_v56_) + (387)] = d_17_v1_
                    return _dafny.Map(coll6_)
                d_85_v58_ = _dafny.MultiSet([d_83_v55_, d_83_v55_, iife6_()
                , (d_84_v57_)[default__.safeIndex(d_17_v1_, len(d_84_v57_))], d_83_v55_])
                (globalState).f2 = ((d_85_v58_) | (d_85_v58_)).ispropersubset(d_85_v58_)
                d_87_v59_: _dafny.Array
                nw10_ = _dafny.Array(False, 18)
                d_87_v59_ = nw10_
                d_88_v60_: D7
                d_88_v60_ = D7_DC17(d_15_v0_)
                d_89_v61_: _dafny.Map
                d_89_v61_ = _dafny.Map({451: (d_88_v60_).cf19})
                index15_ = default__.safeIndex(637, (d_87_v59_).length(0))
                (d_87_v59_)[index15_] = (d_89_v61_) != (d_89_v61_)
                index16_ = default__.safeIndex(637, (d_87_v59_).length(0))
                (d_87_v59_)[index16_] = p0
                (globalState).f2 = (d_87_v59_)[default__.safeIndex(637, (d_87_v59_).length(0))]
                (globalState).f2 = ((0) - ((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))])) != (621)
            d_19_v3_ = _dafny.SeqWithoutIsStrInference([494])
            d_15_v0_ = d_15_v0_
            d_90_v62_: D0
            d_90_v62_ = D0_DC1((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))], (d_18_v2_)[default__.safeIndex(d_17_v1_, len(d_18_v2_))], (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))])
            d_91_v63_: _dafny.Map
            d_91_v63_ = _dafny.Map({d_90_v62_: d_17_v1_})
            d_92_v64_: C0
            nw11_ = C0()
            nw11_.ctor__(d_91_v63_)
            d_92_v64_ = nw11_
            d_93_v65_: str
            d_93_v65_ = _dafny.CodePoint('k')
            d_94_v66_: D9
            d_94_v66_ = D9_DC24(d_17_v1_, d_92_v64_, d_93_v65_, (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))])
            d_95_v67_: _dafny.Map
            d_95_v67_ = _dafny.Map({d_94_v66_: p0})
            d_96_v68_: _dafny.Array
            nw12_ = _dafny.Array(None, 9)
            nw12_[int(0)] = d_18_v2_
            nw12_[int(1)] = d_18_v2_
            nw12_[int(2)] = d_18_v2_
            nw12_[int(3)] = _dafny.SeqWithoutIsStrInference([p0, p0])
            nw12_[int(4)] = (d_18_v2_ if p0 else d_18_v2_)
            nw12_[int(5)] = _dafny.SeqWithoutIsStrInference([not(False), ((d_95_v67_)[d_94_v66_] if (d_94_v66_) in (d_95_v67_) else p0)])
            nw12_[int(6)] = _dafny.SeqWithoutIsStrInference([p0, p0])
            nw12_[int(7)] = d_18_v2_
            nw12_[int(8)] = d_18_v2_
            d_96_v68_ = nw12_
            d_97_v69_: _dafny.Array
            nw13_ = _dafny.Array(D6.default()(), 21)
            d_97_v69_ = nw13_
            d_98_v70_: D6
            d_98_v70_ = D6_DC15(d_17_v1_, d_17_v1_)
            index17_ = default__.safeIndex(512, (d_97_v69_).length(0))
            (d_97_v69_)[index17_] = d_98_v70_
            pat_let_tv0_ = d_17_v1_
            index18_ = default__.safeIndex(512, (d_97_v69_).length(0))
            def iife7_(_pat_let0_0):
                def iife8_(d_99_dt__update__tmp_h0_):
                    def iife9_(_pat_let1_0):
                        def iife10_(d_100_dt__update_hcf17_h0_):
                            return D6_DC15((d_99_dt__update__tmp_h0_).cf16, d_100_dt__update_hcf17_h0_)
                        return iife10_(_pat_let1_0)
                    return iife9_(pat_let_tv0_)
                return iife8_(_pat_let0_0)
            rhs11_ = d_96_v68_
            rhs12_ = iife7_(d_98_v70_)
            lhs16_ = d_97_v69_
            lhs17_ = default__.safeIndex(512, (d_97_v69_).length(0))
            d_96_v68_ = rhs11_
            lhs16_[lhs17_] = rhs12_
            d_101_v71_: _dafny.Array
            nw14_ = _dafny.Array(None, 26)
            d_101_v71_ = nw14_
            d_102_v72_: _dafny.Map
            d_102_v72_ = _dafny.Map({d_101_v71_: p0})
            index19_ = default__.safeIndex(723, (d_15_v0_).length(0))
            index20_ = default__.safeIndex(723, (d_15_v0_).length(0))
            rhs13_ = (0) - ((((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]) * ((0) - (d_17_v1_))) - ((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]))
            rhs14_ = (p0) == (not(((d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]) <= (len(d_102_v72_))))
            rhs15_ = (d_15_v0_)[default__.safeIndex(723, (d_15_v0_).length(0))]
            rhs16_ = default__.fm0(d_17_v1_, d_17_v1_, globalState)
            lhs18_ = globalState
            lhs19_ = globalState
            lhs20_ = d_15_v0_
            lhs21_ = default__.safeIndex(723, (d_15_v0_).length(0))
            lhs22_ = d_15_v0_
            lhs23_ = default__.safeIndex(723, (d_15_v0_).length(0))
            lhs18_.f4 = rhs13_
            lhs19_.f2 = rhs14_
            lhs20_[lhs21_] = rhs15_
            lhs22_[lhs23_] = rhs16_
        d_103_v73_: str
        d_103_v73_ = _dafny.CodePoint('f')
        d_104_v74_: _dafny.Seq
        d_104_v74_ = _dafny.SeqWithoutIsStrInference([d_103_v73_])
        hi1_ = d_17_v1_
        for d_105_i8_ in range((d_17_v1_ if p0 else len(d_104_v74_)), hi1_):
            (globalState).f2 = not (p0) or (p0)
            d_106_v75_: _dafny.Set
            d_106_v75_ = _dafny.Set({(p0 if p0 else p0)})
            d_106_v75_ = d_106_v75_
            d_107_v76_: _dafny.MultiSet
            d_107_v76_ = _dafny.MultiSet([False])
            d_107_v76_ = (_dafny.MultiSet([p0, p0, p0, p0, True])) | ((d_107_v76_).intersection(d_107_v76_))
            d_108_v77_: _dafny.Set
            d_108_v77_ = _dafny.Set({d_17_v1_, d_105_i8_})
            d_109_v78_: D0
            d_109_v78_ = D0_DC1(d_105_i8_, p0, len(d_108_v77_))
            d_110_v79_: _dafny.Map
            def iife11_(_pat_let2_0):
                def iife12_(d_111_dt__update__tmp_h1_):
                    def iife13_(_pat_let3_0):
                        def iife14_(d_112_dt__update_hcf3_h0_):
                            return D0_DC1((d_111_dt__update__tmp_h1_).cf1, (d_111_dt__update__tmp_h1_).cf2, d_112_dt__update_hcf3_h0_)
                        return iife14_(_pat_let3_0)
                    return iife13_(d_105_i8_)
                return iife12_(_pat_let2_0)
            d_110_v79_ = _dafny.Map({iife11_(d_109_v78_): d_17_v1_})
            d_113_v80_: T0
            nw15_ = C0()
            nw15_.ctor__(d_110_v79_)
            d_113_v80_ = nw15_
            d_114_v81_: _dafny.Array
            nw16_ = _dafny.Array(None, 2)
            nw16_[int(0)] = d_113_v80_
            nw16_[int(1)] = d_113_v80_
            d_114_v81_ = nw16_
            d_114_v81_ = (d_114_v81_ if p0 else d_114_v81_)
        d_115_v82_: _dafny.Set
        d_115_v82_ = _dafny.Set({d_19_v3_})
        (globalState).f2 = (p0 if (d_115_v82_).ispropersubset(d_115_v82_) else (d_17_v1_) > (d_17_v1_))

    @staticmethod
    def Main(noArgsParameter__):
        d_116_v0_: int
        d_116_v0_ = 877
        d_117_v1_: _dafny.Set
        d_117_v1_ = _dafny.Set({(0) - (d_116_v0_)})
        d_118_v2_: _dafny.Seq
        d_118_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxsvjmbqp"))
        d_119_v3_: _dafny.Set
        d_119_v3_ = _dafny.Set({d_118_v2_, d_118_v2_})
        d_120_v4_: _dafny.Array
        nw17_ = _dafny.Array(None, 20)
        nw17_[int(0)] = d_116_v0_
        nw17_[int(1)] = d_116_v0_
        nw17_[int(2)] = d_116_v0_
        nw17_[int(3)] = d_116_v0_
        nw17_[int(4)] = d_116_v0_
        nw17_[int(5)] = 782
        nw17_[int(6)] = d_116_v0_
        nw17_[int(7)] = d_116_v0_
        nw17_[int(8)] = d_116_v0_
        nw17_[int(9)] = (D0_DC0(d_116_v0_)).cf0
        nw17_[int(10)] = d_116_v0_
        nw17_[int(11)] = d_116_v0_
        nw17_[int(12)] = len(d_117_v1_)
        nw17_[int(13)] = (0) - (len(d_119_v3_))
        nw17_[int(14)] = d_116_v0_
        nw17_[int(15)] = d_116_v0_
        nw17_[int(16)] = d_116_v0_
        nw17_[int(17)] = (0) - (d_116_v0_)
        nw17_[int(18)] = d_116_v0_
        nw17_[int(19)] = 217
        d_120_v4_ = nw17_
        d_121_v5_: _dafny.Array
        def lambda6_(d_122_v0_):
            def lambda7_(d_123_i0_):
                return (D0_DC1(d_122_v0_, True, d_122_v0_)).cf2

            return lambda7_

        init4_ = lambda6_(d_116_v0_)
        nw18_ = _dafny.Array(None, 12)
        for i0_4_ in range(nw18_.length(0)):
            nw18_[i0_4_] = init4_(i0_4_)
        d_121_v5_ = nw18_
        d_124_globalState_: GlobalState
        nw19_ = GlobalState()
        nw19_.ctor__(False, True, True, -701, 757, 943, d_120_v4_, True, d_120_v4_, d_121_v5_, 685, False, d_118_v2_, 826, True)
        d_124_globalState_ = nw19_
        d_125_i1_: int
        d_125_i1_ = 0
        with _dafny.label("1"):
            while (d_116_v0_) == (d_116_v0_):
                with _dafny.c_label("1"):
                    if (d_125_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_125_i1_ = (d_125_i1_) + (1)
                    d_126_v6_: _dafny.Map
                    d_126_v6_ = _dafny.Map({d_116_v0_: d_116_v0_})
                    index21_ = default__.safeIndex(621, (d_120_v4_).length(0))
                    (d_120_v4_)[index21_] = len((d_126_v6_) | (d_126_v6_))
                    d_127_v7_: bool
                    d_127_v7_ = True
                    d_128_v8_: str
                    d_128_v8_ = _dafny.CodePoint('i')
                    index22_ = default__.safeIndex(621, (d_120_v4_).length(0))
                    (d_120_v4_)[index22_] = default__.fm0((d_116_v0_ if d_127_v7_ else -650), len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_129_i2_ in range(default__.abs(438))]) if True else d_118_v2_)).set(default__.safeIndex(d_116_v0_, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_130_i2_ in range(default__.abs(438))]) if True else d_118_v2_))), d_128_v8_)), d_124_globalState_)
                    index23_ = default__.safeIndex(621, (d_120_v4_).length(0))
                    (d_120_v4_)[index23_] = default__.fm0(d_116_v0_, d_116_v0_, d_124_globalState_)
                    if (d_116_v0_) < (d_116_v0_):
                        d_131_v9_: _dafny.MultiSet
                        d_131_v9_ = _dafny.MultiSet([d_116_v0_])
                        d_132_v10_: _dafny.Seq
                        d_132_v10_ = _dafny.SeqWithoutIsStrInference([d_118_v2_])
                        d_133_v11_: _dafny.MultiSet
                        d_133_v11_ = _dafny.MultiSet([default__.fm0((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], (d_131_v9_).cardinality, d_124_globalState_), (d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], len((d_132_v10_)[default__.safeIndex(((d_131_v9_)[d_116_v0_] if (d_116_v0_) in (d_131_v9_) else d_116_v0_), len(d_132_v10_))]), d_116_v0_])
                        d_134_v12_: _dafny.Array
                        def lambda8_(d_135_v2_):
                            def lambda9_(d_136_i3_):
                                return d_135_v2_

                            return lambda9_

                        init5_ = lambda8_(d_118_v2_)
                        nw20_ = _dafny.Array(None, 20)
                        for i0_5_ in range(nw20_.length(0)):
                            nw20_[i0_5_] = init5_(i0_5_)
                        d_134_v12_ = nw20_
                        index24_ = default__.safeIndex(607, (d_134_v12_).length(0))
                        (d_134_v12_)[index24_] = d_118_v2_
                        d_137_v13_: _dafny.Seq
                        d_137_v13_ = d_118_v2_
                        index25_ = default__.safeIndex(621, (d_120_v4_).length(0))
                        index26_ = default__.safeIndex(607, (d_134_v12_).length(0))
                        rhs17_ = d_116_v0_
                        rhs18_ = d_133_v11_
                        rhs19_ = d_118_v2_
                        rhs20_ = d_127_v7_
                        rhs21_ = (d_137_v13_)
                        lhs24_ = d_120_v4_
                        lhs25_ = default__.safeIndex(621, (d_120_v4_).length(0))
                        lhs26_ = d_124_globalState_
                        lhs27_ = d_134_v12_
                        lhs28_ = default__.safeIndex(607, (d_134_v12_).length(0))
                        lhs24_[lhs25_] = rhs17_
                        d_131_v9_ = rhs18_
                        d_118_v2_ = rhs19_
                        lhs26_.f2 = rhs20_
                        lhs27_[lhs28_] = rhs21_
                        d_138_v14_: _dafny.Seq
                        d_138_v14_ = _dafny.SeqWithoutIsStrInference([(d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], (d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], (d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], d_116_v0_])
                        (d_124_globalState_).f2 = (d_116_v0_) not in (d_138_v14_)
                        d_139_v15_: _dafny.MultiSet
                        d_139_v15_ = _dafny.MultiSet([d_118_v2_, (d_134_v12_)[default__.safeIndex(607, (d_134_v12_).length(0))]])
                        index27_ = default__.safeIndex(607, (d_134_v12_).length(0))
                        rhs22_ = ((d_118_v2_).set(default__.safeIndex((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], len(d_118_v2_)), d_128_v8_)) + ((d_137_v13_))
                        rhs23_ = (d_116_v0_) != ((d_116_v0_) * ((d_138_v14_)[default__.safeIndex(d_116_v0_, len(d_138_v14_))]))
                        rhs24_ = (d_139_v15_).ispropersubset(d_139_v15_)
                        rhs25_ = d_116_v0_
                        lhs29_ = d_134_v12_
                        lhs30_ = default__.safeIndex(607, (d_134_v12_).length(0))
                        lhs31_ = d_124_globalState_
                        lhs32_ = d_124_globalState_
                        lhs33_ = d_124_globalState_
                        lhs29_[lhs30_] = rhs22_
                        lhs31_.f2 = rhs23_
                        lhs32_.f2 = rhs24_
                        lhs33_.f4 = rhs25_
                        d_140_v16_: D0
                        d_140_v16_ = D0_DC1((0) - ((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]), d_127_v7_, len(d_118_v2_))
                        d_141_v17_: C0
                        nw21_ = C0()
                        nw21_.ctor__(_dafny.Map({d_140_v16_: (d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]}))
                        d_141_v17_ = nw21_
                        (d_124_globalState_).f3 = ((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]) - ((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))])
                    elif True:
                        default__.m0((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkmnn"))) != (d_118_v2_), d_124_globalState_)
                        d_142_v18_: _dafny.Array
                        nw22_ = _dafny.Array(None, 3)
                        nw22_[int(0)] = (d_121_v5_ if d_127_v7_ else d_121_v5_)
                        nw22_[int(1)] = d_121_v5_
                        nw22_[int(2)] = d_121_v5_
                        d_142_v18_ = nw22_
                        index28_ = default__.safeIndex(506, (d_142_v18_).length(0))
                        (d_142_v18_)[index28_] = d_121_v5_
                        index29_ = default__.safeIndex(506, (d_142_v18_).length(0))
                        (d_142_v18_)[index29_] = d_121_v5_
                        d_143_v19_: D0
                        d_143_v19_ = D0_DC2()
                        d_143_v19_ = d_143_v19_
                        d_144_v20_: _dafny.Array
                        def lambda10_(d_145_v7_, d_146_v8_, d_147_v2_):
                            def lambda11_(d_148_i4_):
                                return (_dafny.SeqWithoutIsStrInference([d_146_v8_ for d_149_i5_ in range(default__.abs(918))]) if d_145_v7_ else d_147_v2_)

                            return lambda11_

                        init6_ = lambda10_(d_127_v7_, d_128_v8_, d_118_v2_)
                        nw23_ = _dafny.Array(None, 4)
                        for i0_6_ in range(nw23_.length(0)):
                            nw23_[i0_6_] = init6_(i0_6_)
                        d_144_v20_ = nw23_
                        d_150_v21_: _dafny.Map
                        d_150_v21_ = _dafny.Map({not(d_127_v7_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcnuuyv"))})
                        index30_ = default__.safeIndex(931, (d_144_v20_).length(0))
                        (d_144_v20_)[index30_] = ((d_150_v21_)[d_127_v7_] if (d_127_v7_) in (d_150_v21_) else default__.fm1(d_124_globalState_))
                        index31_ = default__.safeIndex(931, (d_144_v20_).length(0))
                        (d_144_v20_)[index31_] = (d_118_v2_) + (d_118_v2_)
                        d_151_v22_: _dafny.MultiSet
                        d_151_v22_ = _dafny.MultiSet([(d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], (d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]])
                        default__.m0((577) not in (d_151_v22_), d_124_globalState_)
                    if (d_116_v0_) == ((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]):
                        (d_124_globalState_).f3 = (0) - ((0) - ((0) - ((default__.safeModuloInt((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], d_116_v0_)) + (d_116_v0_))))
                        (d_124_globalState_).f4 = d_116_v0_
                        index32_ = default__.safeIndex(621, (d_120_v4_).length(0))
                        (d_120_v4_)[index32_] = (0) - (d_116_v0_)
                        d_152_v23_: _dafny.Array
                        def lambda12_(d_153_v7_, d_154_v4_):
                            def lambda13_(d_155_i6_):
                                return (_dafny.Map({False: -2})) | (_dafny.Map({d_153_v7_: len(_dafny.Map({(d_154_v4_)[default__.safeIndex(621, (d_154_v4_).length(0))]: d_153_v7_}))}))

                            return lambda13_

                        init7_ = lambda12_(d_127_v7_, d_120_v4_)
                        nw24_ = _dafny.Array(None, 13)
                        for i0_7_ in range(nw24_.length(0)):
                            nw24_[i0_7_] = init7_(i0_7_)
                        d_152_v23_ = nw24_
                        d_156_v24_: _dafny.Map
                        d_156_v24_ = _dafny.Map({d_127_v7_: d_116_v0_})
                        index33_ = default__.safeIndex(845, (d_152_v23_).length(0))
                        (d_152_v23_)[index33_] = d_156_v24_
                        index34_ = default__.safeIndex(845, (d_152_v23_).length(0))
                        (d_152_v23_)[index34_] = (d_156_v24_).set(d_127_v7_, 410)
                        d_157_v25_: _dafny.Map
                        d_157_v25_ = _dafny.Map({d_116_v0_: _dafny.MultiSet([300, 58])})
                        d_158_v26_: _dafny.Seq
                        d_158_v26_ = _dafny.SeqWithoutIsStrInference([d_116_v0_, d_116_v0_])
                        d_159_v27_: D2
                        d_159_v27_ = D2_DC5(_dafny.MultiSet(d_158_v26_))
                        d_157_v25_ = (d_157_v25_).set(d_116_v0_, (d_159_v27_).cf5)
                    elif True:
                        d_160_v29_: _dafny.MultiSet
                        d_160_v29_ = _dafny.MultiSet([d_127_v7_])
                        d_161_v30_: D0
                        d_161_v30_ = D0_DC1((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], d_127_v7_, (d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))])
                        d_162_v31_: _dafny.Map
                        d_162_v31_ = _dafny.Map({d_161_v30_: d_116_v0_})
                        d_163_v32_: T0
                        nw25_ = C0()
                        nw25_.ctor__(d_162_v31_)
                        d_163_v32_ = nw25_
                        d_164_v33_: _dafny.MultiSet
                        d_164_v33_ = _dafny.MultiSet([d_163_v32_, d_163_v32_])
                        d_165_v34_: D0
                        d_165_v34_ = D0_DC1((d_164_v33_).cardinality, d_127_v7_, d_116_v0_)
                        d_166_v35_: _dafny.MultiSet
                        d_166_v35_ = _dafny.MultiSet([D0_DC1((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))], d_127_v7_, ((d_160_v29_)[d_127_v7_] if (d_127_v7_) in (d_160_v29_) else d_116_v0_)), d_165_v34_])
                        d_167_v36_: T0
                        nw26_ = C0()
                        def iife15_():
                            coll7_ = _dafny.Map()
                            compr_7_: D0
                            for compr_7_ in (d_166_v35_).Elements:
                                d_168_v28_: D0 = compr_7_
                                if (d_168_v28_) in (d_166_v35_):
                                    coll7_[d_168_v28_] = (d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]
                            return _dafny.Map(coll7_)
                        nw26_.ctor__(iife15_()
                        )
                        d_167_v36_ = nw26_
                        d_169_v37_: _dafny.Array
                        nw27_ = _dafny.Array(int(0), 11)
                        d_169_v37_ = nw27_
                        d_170_v38_: _dafny.Map
                        d_170_v38_ = _dafny.Map({d_163_v32_: d_169_v37_})
                        d_171_v39_: C0
                        nw28_ = C0()
                        nw28_.ctor__(d_162_v31_)
                        d_171_v39_ = nw28_
                        d_172_v40_: _dafny.Set
                        d_172_v40_ = _dafny.Set({d_171_v39_})
                        rhs26_ = d_127_v7_
                        rhs27_ = d_127_v7_
                        rhs28_ = ((d_172_v40_) - (d_172_v40_)).isdisjoint((d_172_v40_).intersection(_dafny.Set({d_171_v39_, d_171_v39_, d_171_v39_})))
                        rhs29_ = d_127_v7_
                        rhs30_ = (d_170_v38_) | ((d_170_v38_))
                        lhs34_ = d_124_globalState_
                        lhs35_ = d_124_globalState_
                        d_127_v7_ = rhs26_
                        lhs34_.f2 = rhs27_
                        d_127_v7_ = rhs28_
                        lhs35_.f2 = rhs29_
                        d_170_v38_ = rhs30_
                        index35_ = default__.safeIndex(621, (d_120_v4_).length(0))
                        rhs31_ = d_127_v7_
                        rhs32_ = ((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]) + ((0) - ((d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]))
                        lhs36_ = d_124_globalState_
                        lhs37_ = d_120_v4_
                        lhs38_ = default__.safeIndex(621, (d_120_v4_).length(0))
                        lhs36_.f2 = rhs31_
                        lhs37_[lhs38_] = rhs32_
                        d_173_v41_: _dafny.Set
                        d_173_v41_ = _dafny.Set({d_127_v7_, d_127_v7_, d_127_v7_, d_127_v7_, d_127_v7_})
                        d_174_v42_: _dafny.Seq
                        d_174_v42_ = _dafny.SeqWithoutIsStrInference([d_173_v41_, d_173_v41_, d_173_v41_, d_173_v41_])
                        d_175_v43_: _dafny.Map
                        d_175_v43_ = _dafny.Map({d_127_v7_: _dafny.MultiSet(d_174_v42_)})
                        d_176_v44_: _dafny.MultiSet
                        d_176_v44_ = _dafny.MultiSet([d_173_v41_, d_173_v41_, d_173_v41_])
                        d_177_v45_: _dafny.Map
                        d_177_v45_ = _dafny.Map({len(d_173_v41_): d_176_v44_})
                        d_175_v43_ = (d_175_v43_).set(d_127_v7_, ((d_177_v45_)[d_116_v0_] if (d_116_v0_) in (d_177_v45_) else d_176_v44_))
                        d_178_v46_: _dafny.Array
                        nw29_ = _dafny.Array(_dafny.MultiSet({}), 4)
                        d_178_v46_ = nw29_
                        d_178_v46_ = d_178_v46_
                        d_179_v47_: C0
                        nw30_ = C0()
                        nw30_.ctor__(_dafny.Map({default__.fm2(d_127_v7_, d_124_globalState_): (d_120_v4_)[default__.safeIndex(621, (d_120_v4_).length(0))]}))
                        d_179_v47_ = nw30_
                    pass
            pass
        d_180_v48_: bool
        d_180_v48_ = False
        d_181_v49_: _dafny.Seq
        d_181_v49_ = _dafny.SeqWithoutIsStrInference([d_180_v48_])
        if (d_180_v48_) == ((d_181_v49_) < (_dafny.SeqWithoutIsStrInference([True]))):
            d_182_v50_: _dafny.MultiSet
            d_182_v50_ = _dafny.MultiSet([d_118_v2_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eueosg"))) + (d_118_v2_), d_118_v2_])
            d_116_v0_ = (d_182_v50_).cardinality
            d_183_v51_: str
            d_183_v51_ = _dafny.CodePoint('y')
            d_183_v51_ = d_183_v51_
            d_184_v52_: D0
            d_184_v52_ = D0_DC1(d_116_v0_, d_180_v48_, default__.fm0(d_116_v0_, d_116_v0_, d_124_globalState_))
            d_185_v54_: _dafny.Map
            d_185_v54_ = _dafny.Map({d_184_v52_: d_116_v0_})
            d_186_v55_: C0
            nw31_ = C0()
            def iife16_():
                coll8_ = _dafny.Map()
                compr_8_: D0
                for compr_8_ in (d_185_v54_).keys.Elements:
                    d_187_v53_: D0 = compr_8_
                    if (d_187_v53_) in (d_185_v54_):
                        coll8_[d_187_v53_] = (d_184_v52_).cf3
                return _dafny.Map(coll8_)
            nw31_.ctor__((_dafny.Map({d_184_v52_: d_116_v0_})) | (iife16_()
            ))
            d_186_v55_ = nw31_
            d_188_v56_: _dafny.Seq
            d_188_v56_ = _dafny.SeqWithoutIsStrInference([d_116_v0_])
            if (d_116_v0_) not in (d_188_v56_):
                index36_ = default__.safeIndex(151, (d_121_v5_).length(0))
                (d_121_v5_)[index36_] = d_180_v48_
                d_189_v57_: _dafny.Map
                d_189_v57_ = _dafny.Map({d_116_v0_: d_116_v0_})
                d_190_v58_: T0
                nw32_ = C0()
                nw32_.ctor__(d_185_v54_)
                d_190_v58_ = nw32_
                d_191_v59_: _dafny.Seq
                d_191_v59_ = _dafny.SeqWithoutIsStrInference([d_190_v58_, d_190_v58_])
                d_192_v60_: _dafny.Map
                d_192_v60_ = _dafny.Map({default__.safeModuloInt(len(d_189_v57_), d_116_v0_): ((D4_DC9(d_191_v59_)).cf11) + (d_191_v59_)})
                d_193_v61_: _dafny.Map
                d_193_v61_ = _dafny.Map({d_190_v58_: d_116_v0_})
                d_194_v62_: _dafny.Map
                d_194_v62_ = _dafny.Map({d_190_v58_: (d_193_v61_ if d_180_v48_ else d_193_v61_)})
                index37_ = default__.safeIndex(151, (d_121_v5_).length(0))
                rhs33_ = d_180_v48_
                rhs34_ = (_dafny.Map({d_116_v0_: d_191_v59_})) | (d_192_v60_)
                rhs35_ = d_194_v62_
                lhs39_ = d_121_v5_
                lhs40_ = default__.safeIndex(151, (d_121_v5_).length(0))
                lhs39_[lhs40_] = rhs33_
                d_192_v60_ = rhs34_
                d_194_v62_ = rhs35_
                d_195_v63_: _dafny.Map
                d_195_v63_ = d_185_v54_
                d_196_v64_: C0
                nw33_ = C0()
                nw33_.ctor__((d_195_v63_))
                d_196_v64_ = nw33_
                d_196_v64_ = d_196_v64_
                index38_ = default__.safeIndex(151, (d_121_v5_).length(0))
                (d_121_v5_)[index38_] = ((d_121_v5_)[default__.safeIndex(151, (d_121_v5_).length(0))] if (d_116_v0_) == ((_dafny.MultiSet([d_116_v0_])).cardinality) else d_180_v48_)
                nw34_ = C0()
                nw34_.ctor__((d_186_v55_).f15)
                d_190_v58_ = nw34_
                index39_ = default__.safeIndex(151, (d_121_v5_).length(0))
                (d_121_v5_)[index39_] = (False if (d_183_v51_) not in (d_118_v2_) else (d_121_v5_)[default__.safeIndex(151, (d_121_v5_).length(0))])
            elif True:
                index40_ = default__.safeIndex(553, (d_120_v4_).length(0))
                (d_120_v4_)[index40_] = len(d_188_v56_)
                d_197_v65_: _dafny.MultiSet
                d_197_v65_ = _dafny.MultiSet([d_180_v48_])
                d_198_v66_: _dafny.Map
                d_198_v66_ = _dafny.Map({d_116_v0_: _dafny.MultiSet([d_180_v48_])})
                d_199_v67_: _dafny.MultiSet
                d_199_v67_ = _dafny.MultiSet([d_116_v0_])
                d_200_v68_: _dafny.Array
                nw35_ = _dafny.Array(None, 22)
                nw35_[int(0)] = _dafny.MultiSet([not(True), d_180_v48_])
                nw35_[int(1)] = d_197_v65_
                nw35_[int(2)] = d_197_v65_
                nw35_[int(3)] = d_197_v65_
                nw35_[int(4)] = d_197_v65_
                nw35_[int(5)] = _dafny.MultiSet(d_181_v49_)
                nw35_[int(6)] = d_197_v65_
                nw35_[int(7)] = d_197_v65_
                nw35_[int(8)] = d_197_v65_
                nw35_[int(9)] = _dafny.MultiSet(d_181_v49_)
                nw35_[int(10)] = ((d_198_v66_)[len((_dafny.SeqWithoutIsStrInference([d_116_v0_ for d_201_i7_ in range(default__.abs(892))])).set(default__.safeIndex(d_116_v0_, len(_dafny.SeqWithoutIsStrInference([d_116_v0_ for d_202_i7_ in range(default__.abs(892))]))), d_116_v0_))] if (len((_dafny.SeqWithoutIsStrInference([d_116_v0_ for d_203_i7_ in range(default__.abs(892))])).set(default__.safeIndex(d_116_v0_, len(_dafny.SeqWithoutIsStrInference([d_116_v0_ for d_204_i7_ in range(default__.abs(892))]))), d_116_v0_))) in (d_198_v66_) else d_197_v65_)
                nw35_[int(11)] = d_197_v65_
                nw35_[int(12)] = d_197_v65_
                nw35_[int(13)] = d_197_v65_
                nw35_[int(14)] = d_197_v65_
                nw35_[int(15)] = d_197_v65_
                nw35_[int(16)] = d_197_v65_
                nw35_[int(17)] = _dafny.MultiSet(d_181_v49_)
                nw35_[int(18)] = d_197_v65_
                nw35_[int(19)] = default__.fm3(True, d_199_v67_, d_180_v48_, d_124_globalState_)
                nw35_[int(20)] = _dafny.MultiSet(d_181_v49_)
                nw35_[int(21)] = d_197_v65_
                d_200_v68_ = nw35_
                d_205_v69_: _dafny.Map
                d_205_v69_ = _dafny.Map({d_180_v48_: d_121_v5_})
                d_206_v70_: _dafny.Map
                d_206_v70_ = _dafny.Map({d_200_v68_: ((d_205_v69_)[d_180_v48_] if (d_180_v48_) in (d_205_v69_) else d_121_v5_)})
                d_207_v71_: _dafny.Seq
                d_207_v71_ = _dafny.SeqWithoutIsStrInference([d_121_v5_])
                d_208_v72_: _dafny.Array
                nw36_ = _dafny.Array(None, 25)
                nw36_[int(0)] = d_121_v5_
                nw36_[int(1)] = d_121_v5_
                nw36_[int(2)] = ((d_206_v70_)[d_200_v68_] if (d_200_v68_) in (d_206_v70_) else d_121_v5_)
                nw36_[int(3)] = d_121_v5_
                nw36_[int(4)] = d_121_v5_
                nw36_[int(5)] = d_121_v5_
                nw36_[int(6)] = d_121_v5_
                nw36_[int(7)] = d_121_v5_
                nw36_[int(8)] = d_121_v5_
                nw36_[int(9)] = d_121_v5_
                nw36_[int(10)] = d_121_v5_
                nw36_[int(11)] = (d_207_v71_)[default__.safeIndex(d_116_v0_, len(d_207_v71_))]
                nw36_[int(12)] = d_121_v5_
                nw36_[int(13)] = d_121_v5_
                nw36_[int(14)] = d_121_v5_
                nw36_[int(15)] = d_121_v5_
                nw36_[int(16)] = (D6_DC14(d_121_v5_)).cf15
                nw36_[int(17)] = d_121_v5_
                nw36_[int(18)] = (d_121_v5_ if d_180_v48_ else d_121_v5_)
                nw36_[int(19)] = d_121_v5_
                nw36_[int(20)] = d_121_v5_
                nw36_[int(21)] = d_121_v5_
                nw36_[int(22)] = (d_121_v5_ if d_180_v48_ else d_121_v5_)
                nw36_[int(23)] = d_121_v5_
                nw36_[int(24)] = d_121_v5_
                d_208_v72_ = nw36_
                index41_ = default__.safeIndex(378, (d_208_v72_).length(0))
                (d_208_v72_)[index41_] = d_121_v5_
                d_209_v73_: _dafny.Map
                d_209_v73_ = _dafny.Map({d_118_v2_: d_116_v0_})
                index42_ = default__.safeIndex(553, (d_120_v4_).length(0))
                index43_ = default__.safeIndex(378, (d_208_v72_).length(0))
                rhs36_ = default__.safeDivisionInt(default__.safeDivisionInt((0) - (d_116_v0_), d_116_v0_), ((d_209_v73_)[d_118_v2_] if (d_118_v2_) in (d_209_v73_) else d_116_v0_))
                rhs37_ = d_118_v2_
                rhs38_ = (209) - (602)
                rhs39_ = d_121_v5_
                lhs41_ = d_124_globalState_
                lhs42_ = d_120_v4_
                lhs43_ = default__.safeIndex(553, (d_120_v4_).length(0))
                lhs44_ = d_208_v72_
                lhs45_ = default__.safeIndex(378, (d_208_v72_).length(0))
                lhs41_.f3 = rhs36_
                d_118_v2_ = rhs37_
                lhs42_[lhs43_] = rhs38_
                lhs44_[lhs45_] = rhs39_
                d_210_v74_: C0
                nw37_ = C0()
                nw37_.ctor__(_dafny.Map({d_184_v52_: d_116_v0_}))
                d_210_v74_ = nw37_
                d_211_v75_: _dafny.Map
                d_211_v75_ = _dafny.Map({d_116_v0_: d_118_v2_})
                d_211_v75_ = (d_211_v75_).set((d_120_v4_)[default__.safeIndex(553, (d_120_v4_).length(0))], _dafny.SeqWithoutIsStrInference([d_183_v51_ for d_212_i8_ in range(default__.abs(644))]))
                (d_124_globalState_).f4 = (0) - (((d_120_v4_)[default__.safeIndex(553, (d_120_v4_).length(0))]) + (820))
                index44_ = default__.safeIndex(553, (d_120_v4_).length(0))
                rhs40_ = d_188_v56_
                rhs41_ = (default__.fm4(not(d_180_v48_), d_180_v48_, d_124_globalState_)).cardinality
                lhs46_ = d_120_v4_
                lhs47_ = default__.safeIndex(553, (d_120_v4_).length(0))
                d_188_v56_ = rhs40_
                lhs46_[lhs47_] = rhs41_
            (d_124_globalState_).f2 = d_180_v48_
        elif True:
            d_213_v76_: str
            d_213_v76_ = _dafny.CodePoint('x')
            d_214_v77_: _dafny.Map
            d_214_v77_ = _dafny.Map({d_213_v76_: (d_118_v2_) == (d_118_v2_)})
            d_215_v78_: _dafny.Map
            d_215_v78_ = _dafny.Map({d_116_v0_: d_180_v48_})
            d_216_v79_: D0
            d_216_v79_ = D0_DC1(116, d_180_v48_, d_116_v0_)
            d_217_v80_: _dafny.Map
            d_217_v80_ = _dafny.Map({d_216_v79_: d_116_v0_})
            d_218_v81_: C0
            nw38_ = C0()
            nw38_.ctor__((d_217_v80_).set(d_216_v79_, d_116_v0_))
            d_218_v81_ = nw38_
            d_219_v82_: _dafny.Map
            d_219_v82_ = _dafny.Map({d_218_v81_: d_218_v81_})
            d_220_v83_: _dafny.Map
            d_220_v83_ = _dafny.Map({_dafny.Map({d_218_v81_: d_218_v81_}): d_116_v0_})
            d_221_v84_: _dafny.Seq
            d_221_v84_ = _dafny.SeqWithoutIsStrInference([len(d_181_v49_)])
            def iife17_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(17, 599):
                    d_222_v86_: int = compr_9_
                    if ((17) <= (d_222_v86_)) and ((d_222_v86_) < (599)):
                        coll9_ = coll9_.union(_dafny.Set([(d_222_v86_) * (904)]))
                return _dafny.Set(coll9_)
            def iife18_():
                coll10_ = _dafny.Set()
                compr_10_: int
                for compr_10_ in (_dafny.Set({(d_221_v84_)[default__.safeIndex(d_116_v0_, len(d_221_v84_))]})).Elements:
                    d_223_v85_: int = compr_10_
                    if (d_223_v85_) in (_dafny.Set({(d_221_v84_)[default__.safeIndex(d_116_v0_, len(d_221_v84_))]})):
                        coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_223_v85_, -93)]))
                return _dafny.Set(coll10_)
            rhs42_ = (d_214_v77_ if ((d_219_v82_).set(d_218_v81_, d_218_v81_)) not in (d_220_v83_) else d_214_v77_)
            rhs43_ = (d_215_v78_).set(d_116_v0_, (iife17_()
            ).ispropersubset(iife18_()
            ))
            d_214_v77_ = rhs42_
            d_215_v78_ = rhs43_
            d_224_v87_: _dafny.Map
            d_224_v87_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_118_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlvlgk"))]): d_120_v4_})
            d_224_v87_ = (d_224_v87_).set(default__.fm5(_dafny.Map({d_180_v48_: d_116_v0_}), d_180_v48_, d_116_v0_, d_116_v0_, d_124_globalState_), d_120_v4_)
            default__.m0(d_180_v48_, d_124_globalState_)
            d_225_v88_: _dafny.Map
            d_225_v88_ = _dafny.Map({d_120_v4_: d_116_v0_})
            d_225_v88_ = d_225_v88_
            index45_ = default__.safeIndex(557, (d_120_v4_).length(0))
            (d_120_v4_)[index45_] = (d_116_v0_) * (d_116_v0_)
            d_226_v89_: _dafny.Array
            def lambda14_(d_227_v84_):
                def lambda15_(d_228_i9_):
                    return d_227_v84_

                return lambda15_

            init8_ = lambda14_(d_221_v84_)
            nw39_ = _dafny.Array(None, 1)
            for i0_8_ in range(nw39_.length(0)):
                nw39_[i0_8_] = init8_(i0_8_)
            d_226_v89_ = nw39_
            d_229_v90_: _dafny.Map
            d_229_v90_ = _dafny.Map({default__.fm6(default__.fm6(d_180_v48_, len(_dafny.Set({d_120_v4_, d_120_v4_, d_120_v4_, (D7_DC17(d_120_v4_)).cf19, d_120_v4_})), d_124_globalState_), d_116_v0_, d_124_globalState_): _dafny.SeqWithoutIsStrInference([(0) - (d_116_v0_), d_116_v0_])})
            d_230_v91_: _dafny.Map
            d_230_v91_ = _dafny.Map({d_180_v48_: not(d_180_v48_)})
            d_231_v92_: D2
            d_231_v92_ = D2_DC6(d_230_v91_, d_180_v48_)
            index46_ = default__.safeIndex(76, (d_226_v89_).length(0))
            (d_226_v89_)[index46_] = ((d_229_v90_)[not((d_231_v92_).cf7)] if (not((d_231_v92_).cf7)) in (d_229_v90_) else _dafny.SeqWithoutIsStrInference([d_116_v0_]))
            index47_ = default__.safeIndex(557, (d_120_v4_).length(0))
            index48_ = default__.safeIndex(76, (d_226_v89_).length(0))
            rhs44_ = (d_116_v0_) * (d_116_v0_)
            rhs45_ = ((d_116_v0_) * (d_116_v0_)) + (d_116_v0_)
            rhs46_ = d_221_v84_
            lhs48_ = d_120_v4_
            lhs49_ = default__.safeIndex(557, (d_120_v4_).length(0))
            lhs50_ = d_124_globalState_
            lhs51_ = d_226_v89_
            lhs52_ = default__.safeIndex(76, (d_226_v89_).length(0))
            lhs48_[lhs49_] = rhs44_
            lhs50_.f3 = rhs45_
            lhs51_[lhs52_] = rhs46_
        d_232_i10_: int
        d_232_i10_ = 0
        with _dafny.label("2"):
            while False:
                with _dafny.c_label("2"):
                    if (d_232_i10_) >= (100):
                        raise _dafny.Break("2")
                    d_232_i10_ = (d_232_i10_) + (1)
                    d_233_v93_: D0
                    d_233_v93_ = D0_DC1(d_116_v0_, d_180_v48_, d_116_v0_)
                    d_234_v94_: _dafny.Map
                    d_234_v94_ = _dafny.Map({d_233_v93_: (0) - (d_116_v0_)})
                    d_235_v95_: _dafny.Map
                    d_235_v95_ = d_234_v94_
                    source1_ = d_235_v95_
                    d_236___mcc_h0_ = source1_
                    d_237_cf14_ = d_236___mcc_h0_
                    d_238_v96_: _dafny.Map
                    d_238_v96_ = _dafny.Map({d_180_v48_: d_180_v48_})
                    d_180_v48_ = ((len(d_238_v96_)) <= (d_116_v0_)) == (d_180_v48_)
                    default__.m0(d_180_v48_, d_124_globalState_)
                    d_116_v0_ = (d_116_v0_ if not(d_180_v48_) else d_116_v0_)
                    d_239_v97_: D2
                    d_239_v97_ = D2_DC6(d_238_v96_, d_180_v48_)
                    index49_ = default__.safeIndex(371, (d_121_v5_).length(0))
                    (d_121_v5_)[index49_] = ((d_239_v97_).cf7) == (d_180_v48_)
                    index50_ = default__.safeIndex(371, (d_121_v5_).length(0))
                    (d_121_v5_)[index50_] = d_180_v48_
                    (d_124_globalState_).f2 = not(d_180_v48_)
                    rhs47_ = len(d_118_v2_)
                    rhs48_ = d_116_v0_
                    lhs53_ = d_124_globalState_
                    d_116_v0_ = rhs47_
                    lhs53_.f4 = rhs48_
                    if d_180_v48_:
                        d_240_v98_: D0
                        d_240_v98_ = D0_DC0(d_116_v0_)
                        (d_124_globalState_).f2 = ((d_240_v98_).cf0) > (d_116_v0_)
                        d_118_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuglv"))
                        d_241_v99_: C0
                        nw40_ = C0()
                        nw40_.ctor__(d_234_v94_)
                        d_241_v99_ = nw40_
                        (d_124_globalState_).f4 = (len(default__.fm1(d_124_globalState_))) + (759)
                        d_242_v100_: _dafny.Seq
                        d_242_v100_ = _dafny.SeqWithoutIsStrInference([d_116_v0_, 609, d_116_v0_])
                        index51_ = default__.safeIndex(27, (d_120_v4_).length(0))
                        (d_120_v4_)[index51_] = (d_116_v0_ if d_180_v48_ else len(d_242_v100_))
                        index52_ = default__.safeIndex(27, (d_120_v4_).length(0))
                        (d_120_v4_)[index52_] = d_116_v0_
                    elif True:
                        d_243_v101_: _dafny.Map
                        d_243_v101_ = _dafny.Map({(d_116_v0_) * (d_116_v0_): (d_116_v0_) - (d_116_v0_)})
                        d_243_v101_ = d_243_v101_
                        (d_124_globalState_).f3 = d_116_v0_
                        default__.m0((d_181_v49_) <= (_dafny.SeqWithoutIsStrInference([d_180_v48_])), d_124_globalState_)
                        (d_124_globalState_).f3 = d_116_v0_
                        pat_let_tv1_ = d_116_v0_
                        d_244_v102_: C0
                        nw41_ = C0()
                        def iife19_(_pat_let4_0):
                            def iife20_(d_245_dt__update__tmp_h0_):
                                def iife21_(_pat_let5_0):
                                    def iife22_(d_246_dt__update_hcf3_h0_):
                                        return D0_DC1((d_245_dt__update__tmp_h0_).cf1, (d_245_dt__update__tmp_h0_).cf2, d_246_dt__update_hcf3_h0_)
                                    return iife22_(_pat_let5_0)
                                return iife21_(pat_let_tv1_)
                            return iife20_(_pat_let4_0)
                        nw41_.ctor__(_dafny.Map({iife19_(d_233_v93_): d_116_v0_}))
                        d_244_v102_ = nw41_
                    pass
            pass
        source2_ = D0_DC0(d_116_v0_)
        if source2_.is_DC1:
            d_247___mcc_h1_ = source2_.cf1
            d_248___mcc_h2_ = source2_.cf2
            d_249___mcc_h3_ = source2_.cf3
            d_250_cf3_ = d_249___mcc_h3_
            d_251_cf2_ = d_248___mcc_h2_
            d_252_cf1_ = d_247___mcc_h1_
            d_253_v103_: str
            d_253_v103_ = _dafny.CodePoint('o')
            d_253_v103_ = _dafny.CodePoint('u')
            d_254_v104_: _dafny.Array
            nw42_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_254_v104_ = nw42_
            index53_ = default__.safeIndex(490, (d_254_v104_).length(0))
            (d_254_v104_)[index53_] = d_120_v4_
            index54_ = default__.safeIndex(490, (d_254_v104_).length(0))
            nw43_ = _dafny.Array(int(0), 22)
            (d_254_v104_)[index54_] = nw43_
            index55_ = default__.safeIndex(207, (d_121_v5_).length(0))
            (d_121_v5_)[index55_] = d_180_v48_
            index56_ = default__.safeIndex(207, (d_121_v5_).length(0))
            (d_121_v5_)[index56_] = d_180_v48_
            d_252_cf1_ = default__.fm0(default__.safeModuloInt(d_250_cf3_, d_252_cf1_), default__.fm0(d_250_cf3_, d_250_cf3_, d_124_globalState_), d_124_globalState_)
        elif source2_.is_DC2:
            (d_124_globalState_).f2 = d_180_v48_
            index57_ = default__.safeIndex(690, (d_121_v5_).length(0))
            (d_121_v5_)[index57_] = False
            d_255_v105_: _dafny.Map
            d_255_v105_ = _dafny.Map({d_180_v48_: d_180_v48_})
            d_256_v106_: _dafny.MultiSet
            d_256_v106_ = _dafny.MultiSet([len(d_255_v105_)])
            d_257_v107_: D2
            d_257_v107_ = D2_DC5(d_256_v106_)
            d_258_v108_: _dafny.Map
            d_258_v108_ = _dafny.Map({(D2_DC5(d_256_v106_) if d_180_v48_ else d_257_v107_): (d_116_v0_) == (953)})
            index58_ = default__.safeIndex(690, (d_121_v5_).length(0))
            rhs49_ = not(not((956) < ((0) - (d_116_v0_))))
            rhs50_ = True
            rhs51_ = d_116_v0_
            rhs52_ = ((d_258_v108_)[d_257_v107_] if (d_257_v107_) in (d_258_v108_) else d_180_v48_)
            lhs54_ = d_124_globalState_
            lhs55_ = d_121_v5_
            lhs56_ = default__.safeIndex(690, (d_121_v5_).length(0))
            lhs57_ = d_124_globalState_
            lhs58_ = d_124_globalState_
            lhs54_.f2 = rhs49_
            lhs55_[lhs56_] = rhs50_
            lhs57_.f4 = rhs51_
            lhs58_.f2 = rhs52_
            d_116_v0_ = d_116_v0_
            d_259_v109_: D2
            d_259_v109_ = D2_DC6(d_255_v105_, (d_121_v5_)[default__.safeIndex(690, (d_121_v5_).length(0))])
            d_260_v110_: str
            d_260_v110_ = _dafny.CodePoint('x')
            d_261_v111_: _dafny.Map
            d_261_v111_ = _dafny.Map({d_259_v109_: d_260_v110_})
            d_262_v112_: _dafny.Set
            d_262_v112_ = _dafny.Set({((d_261_v111_)[d_259_v109_] if (d_259_v109_) in (d_261_v111_) else d_260_v110_)})
            d_263_v113_: _dafny.Seq
            d_263_v113_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_116_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfgpdshm"))), d_124_globalState_), d_116_v0_])
            d_264_v114_: _dafny.Seq
            d_264_v114_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_263_v113_)).cardinality])
            rhs53_ = default__.fm7(d_260_v110_, len((default__.fm8(d_116_v0_, d_124_globalState_)) + (d_264_v114_)), d_116_v0_, d_180_v48_, d_124_globalState_)
            rhs54_ = (d_121_v5_)[default__.safeIndex(690, (d_121_v5_).length(0))]
            rhs55_ = d_260_v110_
            lhs59_ = d_124_globalState_
            d_262_v112_ = rhs53_
            lhs59_.f2 = rhs54_
            d_260_v110_ = rhs55_
        elif source2_.is_DC3:
            (d_124_globalState_).f2 = d_180_v48_
            d_265_v115_: D0
            d_265_v115_ = D0_DC0(d_116_v0_)
            d_265_v115_ = d_265_v115_
            (d_124_globalState_).f2 = d_180_v48_
            if d_180_v48_:
                d_266_v116_: D0
                d_266_v116_ = D0_DC1(len(d_118_v2_), not(d_180_v48_), d_116_v0_)
                d_267_v117_: _dafny.Map
                d_267_v117_ = _dafny.Map({d_266_v116_: len(d_181_v49_)})
                d_268_v118_: C0
                nw44_ = C0()
                nw44_.ctor__(d_267_v117_)
                d_268_v118_ = nw44_
                d_118_v2_ = (d_118_v2_) + ((d_118_v2_) + (d_118_v2_))
                d_269_v119_: C0
                nw45_ = C0()
                nw45_.ctor__((d_268_v118_).f15)
                d_269_v119_ = nw45_
                d_269_v119_ = d_268_v118_
                index59_ = default__.safeIndex(125, (d_121_v5_).length(0))
                (d_121_v5_)[index59_] = (d_116_v0_) < (d_116_v0_)
                index60_ = default__.safeIndex(125, (d_121_v5_).length(0))
                (d_121_v5_)[index60_] = default__.fm6(False, 603, d_124_globalState_)
            elif True:
                d_118_v2_ = ((d_118_v2_) + (d_118_v2_)) + (d_118_v2_)
                d_270_v120_: D0
                d_270_v120_ = D0_DC1((0) - ((0) - (d_116_v0_)), d_180_v48_, d_116_v0_)
                d_271_v121_: _dafny.Map
                d_271_v121_ = _dafny.Map({d_270_v120_: d_116_v0_})
                d_272_v122_: C0
                nw46_ = C0()
                nw46_.ctor__(d_271_v121_)
                d_272_v122_ = nw46_
                d_273_v124_: _dafny.Seq
                d_273_v124_ = _dafny.SeqWithoutIsStrInference([d_270_v120_, d_270_v120_])
                d_274_v125_: T0
                nw47_ = C0()
                def iife23_():
                    coll11_ = _dafny.Map()
                    compr_11_: D0
                    for compr_11_ in (d_273_v124_).Elements:
                        d_275_v123_: D0 = compr_11_
                        if (d_275_v123_) in (d_273_v124_):
                            coll11_[d_275_v123_] = d_116_v0_
                    return _dafny.Map(coll11_)
                nw47_.ctor__(iife23_()
                )
                d_274_v125_ = nw47_
                index61_ = default__.safeIndex(339, (d_120_v4_).length(0))
                (d_120_v4_)[index61_] = d_116_v0_
                d_276_v126_: _dafny.Map
                d_276_v126_ = _dafny.Map({d_180_v48_: d_180_v48_})
                d_277_v127_: D2
                d_277_v127_ = D2_DC6(d_276_v126_, (d_180_v48_) and (d_180_v48_))
                index62_ = default__.safeIndex(339, (d_120_v4_).length(0))
                rhs56_ = d_274_v125_
                rhs57_ = default__.safeModuloInt(d_116_v0_, d_116_v0_)
                rhs58_ = d_180_v48_
                rhs59_ = d_116_v0_
                rhs60_ = d_277_v127_
                lhs60_ = d_120_v4_
                lhs61_ = default__.safeIndex(339, (d_120_v4_).length(0))
                lhs62_ = d_124_globalState_
                lhs63_ = d_124_globalState_
                d_274_v125_ = rhs56_
                lhs60_[lhs61_] = rhs57_
                lhs62_.f2 = rhs58_
                lhs63_.f3 = rhs59_
                d_277_v127_ = rhs60_
                d_278_v128_: str
                d_278_v128_ = _dafny.CodePoint('c')
                d_279_v129_: D8
                d_279_v129_ = D8_DC20(d_278_v128_)
                d_278_v128_ = (d_279_v129_).cf22
                d_280_v130_: _dafny.Array
                nw48_ = _dafny.Array(None, 6)
                d_280_v130_ = nw48_
                index63_ = default__.safeIndex(44, (d_280_v130_).length(0))
                (d_280_v130_)[index63_] = d_274_v125_
                d_281_v131_: _dafny.Map
                d_281_v131_ = _dafny.Map({d_180_v48_: d_274_v125_})
                index64_ = default__.safeIndex(44, (d_280_v130_).length(0))
                (d_280_v130_)[index64_] = ((d_281_v131_)[d_180_v48_] if (d_180_v48_) in (d_281_v131_) else d_274_v125_)
        elif True:
            d_282___mcc_h4_ = source2_.cf0
            d_283_cf0_ = d_282___mcc_h4_
            d_284_v132_: str
            d_284_v132_ = _dafny.CodePoint('i')
            d_285_v133_: _dafny.MultiSet
            d_285_v133_ = _dafny.MultiSet([d_284_v132_, d_284_v132_, d_284_v132_, d_284_v132_])
            (d_124_globalState_).f2 = default__.fm6(d_180_v48_, (d_285_v133_).cardinality, d_124_globalState_)
            index65_ = default__.safeIndex(591, (d_121_v5_).length(0))
            (d_121_v5_)[index65_] = (d_181_v49_)[default__.safeIndex(d_116_v0_, len(d_181_v49_))]
            d_286_v134_: _dafny.Seq
            d_286_v134_ = _dafny.SeqWithoutIsStrInference([d_116_v0_])
            index66_ = default__.safeIndex(591, (d_121_v5_).length(0))
            (d_121_v5_)[index66_] = (default__.safeModuloInt(694, len(_dafny.SeqWithoutIsStrInference([d_180_v48_])))) <= ((d_286_v134_)[default__.safeIndex(d_283_cf0_, len(d_286_v134_))])
            d_287_v135_: _dafny.Set
            d_287_v135_ = _dafny.Set({d_180_v48_})
            d_288_v136_: _dafny.Map
            d_288_v136_ = _dafny.Map({D0_DC1(len(d_287_v135_), d_180_v48_, d_116_v0_): (0) - (d_283_cf0_)})
            d_289_v137_: C0
            nw49_ = C0()
            nw49_.ctor__(d_288_v136_)
            d_289_v137_ = nw49_
            d_290_v138_: _dafny.Map
            d_290_v138_ = _dafny.Map({d_180_v48_: d_289_v137_})
            if (d_117_v1_).issubset(default__.fm9(len(d_290_v138_), default__.fm0(d_283_cf0_, d_116_v0_, d_124_globalState_), d_124_globalState_)):
                (d_124_globalState_).f4 = (default__.safeModuloInt(d_283_cf0_, d_283_cf0_)) - (-167)
                d_291_v139_: _dafny.Map
                d_291_v139_ = _dafny.Map({_dafny.CodePoint('v'): (d_121_v5_)[default__.safeIndex(591, (d_121_v5_).length(0))]})
                default__.m0(((d_291_v139_)[d_284_v132_] if (d_284_v132_) in (d_291_v139_) else True), d_124_globalState_)
                d_289_v137_ = d_289_v137_
                d_292_v141_: _dafny.MultiSet
                d_292_v141_ = _dafny.MultiSet([d_283_cf0_])
                d_293_v142_: D0
                d_293_v142_ = D0_DC1(((d_292_v141_)[d_283_cf0_] if (d_283_cf0_) in (d_292_v141_) else d_116_v0_), (d_121_v5_)[default__.safeIndex(591, (d_121_v5_).length(0))], d_283_cf0_)
                d_294_v143_: _dafny.Map
                d_294_v143_ = _dafny.Map({d_293_v142_: (d_121_v5_)[default__.safeIndex(591, (d_121_v5_).length(0))]})
                d_295_v144_: C0
                nw50_ = C0()
                def iife24_():
                    coll12_ = _dafny.Map()
                    compr_12_: D0
                    for compr_12_ in (d_294_v143_).keys.Elements:
                        d_296_v140_: D0 = compr_12_
                        if (d_296_v140_) in (d_294_v143_):
                            coll12_[d_296_v140_] = d_116_v0_
                    return _dafny.Map(coll12_)
                nw50_.ctor__((iife24_()
                ) | (_dafny.Map({d_293_v142_: d_116_v0_})))
                d_295_v144_ = nw50_
                (d_124_globalState_).f2 = (d_284_v132_) not in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_297_i11_ in range(default__.abs(750))]))
            elif True:
                d_284_v132_ = (D8_DC20(d_284_v132_)).cf22
                default__.m0((d_121_v5_)[default__.safeIndex(591, (d_121_v5_).length(0))], d_124_globalState_)
                (d_124_globalState_).f3 = d_116_v0_
                d_298_v145_: _dafny.Map
                d_298_v145_ = _dafny.Map({(d_121_v5_)[default__.safeIndex(591, (d_121_v5_).length(0))]: 189})
                d_299_v146_: _dafny.Map
                d_299_v146_ = _dafny.Map({d_116_v0_: ((d_298_v145_)[True] if (True) in (d_298_v145_) else d_283_cf0_)})
                d_300_v147_: _dafny.Map
                d_300_v147_ = _dafny.Map({((d_299_v146_)[d_283_cf0_] if (d_283_cf0_) in (d_299_v146_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yeoger")))): d_283_cf0_})
                d_301_v148_: _dafny.Seq
                d_301_v148_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({796: 480})])
                d_300_v147_ = (((d_301_v148_)[default__.safeIndex(d_283_cf0_, len(d_301_v148_))]) | (_dafny.Map({len(d_118_v2_): 503}))) | (d_300_v147_)
                d_118_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfpkjomp"))
            d_302_v149_: _dafny.Map
            d_302_v149_ = _dafny.Map({default__.fm10(d_124_globalState_): d_289_v137_})
            d_283_cf0_ = len(d_302_v149_)
        index67_ = default__.safeIndex(642, (d_120_v4_).length(0))
        (d_120_v4_)[index67_] = d_116_v0_
        d_303_v150_: _dafny.Set
        d_303_v150_ = _dafny.Set({d_121_v5_, d_121_v5_, d_121_v5_, d_121_v5_})
        index68_ = default__.safeIndex(642, (d_120_v4_).length(0))
        (d_120_v4_)[index68_] = default__.safeModuloInt(len(d_303_v150_), 328)
        d_304_v151_: _dafny.Map
        d_304_v151_ = _dafny.Map({d_116_v0_: default__.fm6(d_180_v48_, (d_120_v4_)[default__.safeIndex(642, (d_120_v4_).length(0))], d_124_globalState_)})
        index69_ = default__.safeIndex(359, (d_121_v5_).length(0))
        (d_121_v5_)[index69_] = d_180_v48_
        index70_ = default__.safeIndex(642, (d_120_v4_).length(0))
        index71_ = default__.safeIndex(359, (d_121_v5_).length(0))
        rhs61_ = ((d_120_v4_)[default__.safeIndex(642, (d_120_v4_).length(0))]) + ((0) - (d_116_v0_))
        rhs62_ = d_180_v48_
        rhs63_ = d_304_v151_
        rhs64_ = d_120_v4_
        rhs65_ = (956) != (d_116_v0_)
        lhs64_ = d_120_v4_
        lhs65_ = default__.safeIndex(642, (d_120_v4_).length(0))
        lhs66_ = d_121_v5_
        lhs67_ = default__.safeIndex(359, (d_121_v5_).length(0))
        lhs64_[lhs65_] = rhs61_
        d_180_v48_ = rhs62_
        d_304_v151_ = rhs63_
        d_120_v4_ = rhs64_
        lhs66_[lhs67_] = rhs65_
        d_305_v152_: _dafny.MultiSet
        d_305_v152_ = _dafny.MultiSet([d_117_v1_])
        d_306_v153_: _dafny.Seq
        d_306_v153_ = _dafny.SeqWithoutIsStrInference([d_305_v152_])
        d_307_i12_: int
        d_307_i12_ = 0
        with _dafny.label("3"):
            while (((d_306_v153_)[default__.safeIndex(d_116_v0_, len(d_306_v153_))]) | (d_305_v152_)).ispropersubset((d_305_v152_) | (d_305_v152_)):
                with _dafny.c_label("3"):
                    if (d_307_i12_) >= (100):
                        raise _dafny.Break("3")
                    d_307_i12_ = (d_307_i12_) + (1)
                    d_308_v154_: D0
                    d_308_v154_ = D0_DC1((d_120_v4_)[default__.safeIndex(642, (d_120_v4_).length(0))], (d_121_v5_)[default__.safeIndex(359, (d_121_v5_).length(0))], (d_120_v4_)[default__.safeIndex(642, (d_120_v4_).length(0))])
                    d_309_v155_: _dafny.Map
                    d_309_v155_ = _dafny.Map({d_308_v154_: d_116_v0_})
                    d_310_v156_: C0
                    nw51_ = C0()
                    nw51_.ctor__(d_309_v155_)
                    d_310_v156_ = nw51_
                    d_311_v157_: _dafny.Map
                    d_311_v157_ = _dafny.Map({d_310_v156_: d_118_v2_})
                    index72_ = default__.safeIndex(359, (d_121_v5_).length(0))
                    index73_ = default__.safeIndex(359, (d_121_v5_).length(0))
                    rhs66_ = (d_121_v5_)[default__.safeIndex(359, (d_121_v5_).length(0))]
                    rhs67_ = ((d_311_v157_)[d_310_v156_] if (d_310_v156_) in (d_311_v157_) else d_118_v2_)
                    rhs68_ = (True) and (not(d_180_v48_))
                    lhs68_ = d_121_v5_
                    lhs69_ = default__.safeIndex(359, (d_121_v5_).length(0))
                    lhs70_ = d_121_v5_
                    lhs71_ = default__.safeIndex(359, (d_121_v5_).length(0))
                    lhs68_[lhs69_] = rhs66_
                    d_118_v2_ = rhs67_
                    lhs70_[lhs71_] = rhs68_
                    d_312_v158_: str
                    d_312_v158_ = _dafny.CodePoint('d')
                    d_313_v159_: D8
                    d_313_v159_ = D8_DC20(d_312_v158_)
                    d_180_v48_ = (d_313_v159_) not in (_dafny.SeqWithoutIsStrInference([d_313_v159_ for d_314_i13_ in range(default__.abs(982))]))
                    d_310_v156_ = d_310_v156_
                    d_315_v160_: _dafny.Map
                    d_315_v160_ = _dafny.Map({not((d_121_v5_)[default__.safeIndex(359, (d_121_v5_).length(0))]): d_310_v156_})
                    d_315_v160_ = d_315_v160_
                    pass
            pass
        rhs69_ = d_180_v48_
        rhs70_ = _dafny.SeqWithoutIsStrInference([d_180_v48_])
        lhs72_ = d_124_globalState_
        lhs72_.f2 = rhs69_
        d_181_v49_ = rhs70_
        (d_124_globalState_).f2 = d_180_v48_
        d_316_v161_: _dafny.Array
        def lambda16_(d_317_v48_, d_318_v0_):
            def lambda17_(d_319_i15_):
                return _dafny.Map({d_317_v48_: d_318_v0_})

            return lambda17_

        init9_ = lambda16_(d_180_v48_, d_116_v0_)
        nw52_ = _dafny.Array(None, 4)
        for i0_9_ in range(nw52_.length(0)):
            nw52_[i0_9_] = init9_(i0_9_)
        d_316_v161_ = nw52_
        _ingredientsd_0 = []
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_316_v161_).length(0)):
            d_320_i14_: int = guard_loop_0_
            if (True) and (((0) <= (d_320_i14_)) and ((d_320_i14_) < ((d_316_v161_).length(0)))):
                _ingredientsd_0.append((d_316_v161_, int(d_320_i14_), (_dafny.Map({False: d_116_v0_})) | (_dafny.Map({(d_121_v5_)[default__.safeIndex(359, (d_121_v5_).length(0))]: (d_120_v4_)[default__.safeIndex(642, (d_120_v4_).length(0))]}))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        default__.m0(not (d_180_v48_) or ((d_121_v5_)[default__.safeIndex(359, (d_121_v5_).length(0))]), d_124_globalState_)
        d_321_v162_: _dafny.Seq
        d_321_v162_ = (d_118_v2_)
        pat_let_tv2_ = d_116_v0_
        pat_let_tv3_ = d_120_v4_
        pat_let_tv4_ = d_120_v4_
        index74_ = default__.safeIndex(642, (d_120_v4_).length(0))
        def lambda18_(source3_):
            d_322___mcc_h5_ = source3_
            d_323_cf4_ = d_322___mcc_h5_
            return (0) - ((pat_let_tv2_) * ((_dafny.MultiSet([(0) - ((pat_let_tv4_)[default__.safeIndex(642, (pat_let_tv3_).length(0))])])).cardinality))

        (d_120_v4_)[index74_] = lambda18_(d_321_v162_)
        d_116_v0_ = d_116_v0_
        d_324_v163_: str
        d_324_v163_ = _dafny.CodePoint('e')
        d_325_v164_: _dafny.Map
        d_325_v164_ = _dafny.Map({d_324_v163_: d_180_v48_})
        d_326_v165_: _dafny.Map
        d_326_v165_ = _dafny.Map({d_180_v48_: len(d_325_v164_)})
        d_327_v166_: D0
        d_327_v166_ = D0_DC1(d_116_v0_, d_180_v48_, (0) - (len(d_326_v165_)))
        d_328_v167_: _dafny.Map
        d_328_v167_ = _dafny.Map({d_327_v166_: d_116_v0_})
        d_329_v168_: C0
        nw53_ = C0()
        nw53_.ctor__(d_328_v167_)
        d_329_v168_ = nw53_
        d_330_v169_: _dafny.MultiSet
        d_330_v169_ = _dafny.MultiSet([_dafny.Map({d_329_v168_: d_329_v168_})])
        d_331_v170_: D8
        d_331_v170_ = D8_DC21(d_116_v0_, d_329_v168_)
        d_332_v171_: _dafny.Map
        d_332_v171_ = _dafny.Map({(d_331_v170_).cf24: d_329_v168_})
        d_330_v169_ = (d_330_v169_).set(d_332_v171_, default__.abs(d_116_v0_))
        d_333_i16_: int
        d_333_i16_ = 0
        with _dafny.label("4"):
            while ((d_121_v5_)[default__.safeIndex(359, (d_121_v5_).length(0))]) or (d_180_v48_):
                with _dafny.c_label("4"):
                    if (d_333_i16_) >= (100):
                        raise _dafny.Break("4")
                    d_333_i16_ = (d_333_i16_) + (1)
                    d_326_v165_ = (d_326_v165_).set(d_180_v48_, d_116_v0_)
                    d_116_v0_ = (-967) + (len((d_117_v1_) | (d_117_v1_)))
                    d_118_v2_ = d_118_v2_
                    d_326_v165_ = (d_326_v165_).set(d_180_v48_, (d_120_v4_)[default__.safeIndex(642, (d_120_v4_).length(0))])
                    pass
            pass
        d_334_v172_: _dafny.Array
        nw54_ = _dafny.Array(_dafny.Array(None, 0), 3)
        d_334_v172_ = nw54_
        d_335_v173_: _dafny.Map
        d_335_v173_ = _dafny.Map({d_334_v172_: d_180_v48_})
        d_335_v173_ = (d_335_v173_).set(d_334_v172_, d_180_v48_)
        _dafny.print(_dafny.string_of(d_116_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_117_v1_) == (_dafny.Set({-877}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_118_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v3_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxsvjmbqp"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v4_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_124_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_124_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_124_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f6)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f8)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_124_globalState_).f9)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_124_globalState_).f12).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_125_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_180_v48_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_v49_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_232_i10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_303_v150_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v151_) == (_dafny.Map({3: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v152_) == (_dafny.MultiSet([_dafny.Set({-877})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v153_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.Set({-877})])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_307_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_316_v161_)[0]) == (_dafny.Map({False: 3, True: -2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_316_v161_)[1]) == (_dafny.Map({False: 3, True: -2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_316_v161_)[2]) == (_dafny.Map({False: 3, True: -2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_316_v161_)[3]) == (_dafny.Map({False: 3, True: -2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_321_v162_)).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_324_v163_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_325_v164_) == (_dafny.Map({_dafny.CodePoint('e'): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_326_v165_) == (_dafny.Map({False: -3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v166_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v166_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v166_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_328_v167_) == (_dafny.Map({D0_DC1(3, False, -1): 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_329_v168_).f15) == (_dafny.Map({D0_DC1(3, False, -1): 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_330_v169_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_331_v170_).cf23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_331_v170_).cf24).f15) == (_dafny.Map({D0_DC1(3, False, -1): 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_332_v171_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_333_i16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_335_v173_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), False, int(0))
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

class D0_DC2(D0, NamedTuple('DC2', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3)
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
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({self.cf4.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(_dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC8(D3, NamedTuple('DC8', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC10(D4, NamedTuple('DC10', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC13(D5, NamedTuple('DC13', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC15(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC15(D6, NamedTuple('DC15', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC18(D7, NamedTuple('DC18', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC21(D8, NamedTuple('DC21', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)

class D9_DC23(D9, NamedTuple('DC23', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf27', Any), ('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC27(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)

class D10_DC27(D10, NamedTuple('DC27', [('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC26(D10, NamedTuple('DC26', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC30(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)

class D11_DC30(D11, NamedTuple('DC30', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f2: bool = False
        self.f3: int = int(0)
        self.f4: int = int(0)
        self._f0: bool = False
        self._f1: bool = False
        self._f5: int = int(0)
        self._f6: _dafny.Array = _dafny.Array(None, 0)
        self._f7: bool = False
        self._f8: _dafny.Array = _dafny.Array(None, 0)
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f10: int = int(0)
        self._f11: bool = False
        self._f12: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f13: int = int(0)
        self._f14: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
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
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14

class C0(T0):
    def  __init__(self):
        self._f15: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f15):
        (self)._f15 = f15

    @property
    def f15(self):
        return self._f15
