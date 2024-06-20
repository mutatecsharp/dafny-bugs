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
    def fm2(p0, p1, globalState):
        source0_ = D2_DC6(len(_dafny.SeqWithoutIsStrInference([True, True, False, True])))
        if source0_.is_DC5:
            d_0___mcc_h0_ = source0_.cf6
            d_1___mcc_h1_ = source0_.cf7
            d_2___mcc_h2_ = source0_.cf8
            d_3_cf8_ = d_2___mcc_h2_
            d_4_cf7_ = d_1___mcc_h1_
            d_5_cf6_ = d_0___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([False])
        elif source0_.is_DC6:
            d_6___mcc_h3_ = source0_.cf9
            d_7_cf9_ = d_6___mcc_h3_
            return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True]))
        elif source0_.is_DC7:
            d_8___mcc_h4_ = source0_.cf10
            d_9___mcc_h5_ = source0_.cf11
            d_10_cf11_ = d_9___mcc_h5_
            d_11_cf10_ = d_8___mcc_h4_
            return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True]))
        elif True:
            d_12___mcc_h6_ = source0_.cf5
            d_13_cf5_ = d_12___mcc_h6_
            return _dafny.SeqWithoutIsStrInference([True, False, True])

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(385, 991):
                d_14_v0_: int = compr_0_
                if ((385) <= (d_14_v0_)) and ((d_14_v0_) < (991)):
                    coll0_[(d_14_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlcha"))))] = 726
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(476, 250):
                d_15_v1_: int = compr_1_
                if ((476) <= (d_15_v1_)) and ((d_15_v1_) < (250)):
                    coll1_[(d_15_v1_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ge"))))] = (0) - (len(_dafny.SeqWithoutIsStrInference([True])))
            return _dafny.Map(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-390, 638):
                d_16_v2_: int = compr_2_
                if ((-390) <= (d_16_v2_)) and ((d_16_v2_) < (638)):
                    coll2_[(d_16_v2_) * (581)] = _dafny.CodePoint('u')
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: _dafny.Seq
            for compr_3_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (-420)])]))).Elements:
                d_17_v3_: _dafny.Seq = compr_3_
                if (d_17_v3_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (-420)])]))):
                    coll3_ = coll3_.union(_dafny.Set([d_17_v3_]))
            return _dafny.Set(coll3_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Set({-960, 274})), 714])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.Map({False: True}))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbbbp")))])}))) | ((_dafny.Set({_dafny.SeqWithoutIsStrInference([311, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yckvpg"))), len(_dafny.Set({D6_DC19(True, len(_dafny.Map({(0) - (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eshn")))}))): 869})), 101, len(_dafny.Map({False: -108})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))), D6_DC19(True, len(_dafny.Map({True: 950})), 121, len(_dafny.Set({len(iife0_()
)})), len(_dafny.Map({False: len(_dafny.Map({365: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnjtc"))): len(_dafny.Map({not(True): len(_dafny.Map({not(False): False}))}))}))}))}))), D6_DC19(False, len(_dafny.Map({True: True})), len(iife1_()
), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, False]))])), len(_dafny.Set({len(_dafny.Map({len(iife2_()
): (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.Set({True})), 903, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mi"))), -912])).cardinality, 272])).cardinality})), len(_dafny.Map({False: False}))}))])), 948)})), 863])})).intersection(iife3_()
        ))

    @staticmethod
    def fm6(p0, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rf")))])).cardinality])): True})

    @staticmethod
    def fm7(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_18_i0_ in range(default__.abs(-444))])).Elements:
                d_19_v0_: int = compr_4_
                if (d_19_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_18_i0_ in range(default__.abs(-444))])):
                    coll4_[(d_19_v0_) + (len(_dafny.Set({True})))] = 390
            return _dafny.Map(coll4_)
        return (0) - ((0) - ((default__.safeDivisionInt(951, len(_dafny.SeqWithoutIsStrInference([len(iife4_()
        )])))) * (((D5_DC16(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmlkxdin"))), True, True)).cf22) * (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))): _dafny.CodePoint('f')}))))))

    @staticmethod
    def fm8(globalState):
        source1_ = D5_DC15(True, False)
        if source1_.is_DC15:
            d_20___mcc_h0_ = source1_.cf20
            d_21___mcc_h1_ = source1_.cf21
            d_22_cf21_ = d_21___mcc_h1_
            d_23_cf20_ = d_20___mcc_h0_
            return D1_DC2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpt"))), d_23_cf20_)
        elif source1_.is_DC16:
            d_24___mcc_h2_ = source1_.cf22
            d_25___mcc_h3_ = source1_.cf23
            d_26___mcc_h4_ = source1_.cf24
            d_27_cf24_ = d_26___mcc_h4_
            d_28_cf23_ = d_25___mcc_h3_
            d_29_cf22_ = d_24___mcc_h2_
            return D1_DC2(len(_dafny.Set({d_27_cf24_, d_28_cf23_})), d_27_cf24_)
        elif True:
            d_30___mcc_h5_ = source1_.cf19
            d_31_cf19_ = d_30___mcc_h5_
            return D1_DC2(-957, True)

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggyklnkw"))), 133])

    @staticmethod
    def fm10(p0, globalState):
        return (_dafny.Set({False})) - ((_dafny.Set({not(not(not(False)))})).intersection(_dafny.Set({False, True, not(not(False)), True, True})))

    @staticmethod
    def fm11(p0, globalState):
        if not(not((True if False else False))):
            return not(True)
        elif True:
            return False

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: _dafny.Seq
            for compr_5_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjljvpj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imkj")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_32_i0_ in range(default__.abs(208))])])).Elements:
                d_33_v0_: _dafny.Seq = compr_5_
                if (d_33_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjljvpj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imkj")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_32_i0_ in range(default__.abs(208))])])):
                    coll5_ = coll5_.union(_dafny.Set([d_33_v0_]))
            return _dafny.Set(coll5_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suuupmma")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogyyc"))})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgrjnfjl"))}))).intersection(iife5_()
        )

    @staticmethod
    def fm13(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(762) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdm"))))])

    @staticmethod
    def fm14(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_34_i0_ in range(default__.abs(-952))])

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        def iife6_():
            def iife8_():
                coll8_ = _dafny.Set()
                compr_8_: str
                for compr_8_ in (_dafny.Map({_dafny.CodePoint('d'): len(_dafny.Set({(0) - (len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D2_DC5(True, False, _dafny.Set({-109, 518}))).cf6]))).cardinality: (0) - (-434)})))}))})).keys.Elements:
                    d_37_v0_: str = compr_8_
                    if (d_37_v0_) in (_dafny.Map({_dafny.CodePoint('d'): len(_dafny.Set({(0) - (len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D2_DC5(True, False, _dafny.Set({-109, 518}))).cf6]))).cardinality: (0) - (-434)})))}))})):
                        coll8_ = coll8_.union(_dafny.Set([d_37_v0_]))
                return _dafny.Set(coll8_)
            coll6_ = _dafny.Set()
            def iife7_():
                coll7_ = _dafny.Set()
                compr_7_: str
                for compr_7_ in (_dafny.Map({_dafny.CodePoint('d'): len(_dafny.Set({(0) - (len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D2_DC5(True, False, _dafny.Set({-109, 518}))).cf6]))).cardinality: (0) - (-434)})))}))})).keys.Elements:
                    d_35_v0_: str = compr_7_
                    if (d_35_v0_) in (_dafny.Map({_dafny.CodePoint('d'): len(_dafny.Set({(0) - (len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D2_DC5(True, False, _dafny.Set({-109, 518}))).cf6]))).cardinality: (0) - (-434)})))}))})):
                        coll7_ = coll7_.union(_dafny.Set([d_35_v0_]))
                return _dafny.Set(coll7_)
            compr_6_: str
            for compr_6_ in (iife7_()
            ).Elements:
                d_36_v1_: str = compr_6_
                if (d_36_v1_) in (iife8_()
                ):
                    coll6_ = coll6_.union(_dafny.Set([d_36_v1_]))
            return _dafny.Set(coll6_)
        return _dafny.Map({(len(_dafny.Map({(D5_DC16(len(_dafny.Map({False: len(iife6_()
)})), True, False)).cf22: True}))) > (-981): (_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([35, 349]))})).ispropersubset(_dafny.Set({_dafny.MultiSet([837, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcdsylhn")))])}))})

    @staticmethod
    def m0(p0, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        d_38_v0_: int
        d_38_v0_ = 353
        d_39_v1_: str
        d_39_v1_ = _dafny.CodePoint('d')
        d_40_v2_: _dafny.Map
        d_40_v2_ = _dafny.Map({default__.fm7(136, d_38_v0_, globalState): d_39_v1_})
        d_40_v2_ = (d_40_v2_).set(d_38_v0_, d_39_v1_)
        d_41_v3_: D2
        d_41_v3_ = D2_DC6(d_38_v0_)
        d_42_v4_: C0
        nw0_ = C0()
        nw0_.ctor__((d_41_v3_).cf9)
        d_42_v4_ = nw0_
        d_43_v5_: _dafny.Map
        d_43_v5_ = _dafny.Map({default__.fm11(641, globalState): d_42_v4_})
        d_44_v6_: _dafny.Array
        nw1_ = _dafny.Array(_dafny.Array(None, 0), 10)
        d_44_v6_ = nw1_
        d_45_v8_: _dafny.Array
        def lambda0_(d_46_v1_):
            def lambda1_(d_47_i0_):
                def iife9_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(826, 649):
                        d_48_v7_: int = compr_9_
                        if ((826) <= (d_48_v7_)) and ((d_48_v7_) < (649)):
                            coll9_[default__.safeModuloInt(d_48_v7_, len(_dafny.Map({True: d_46_v1_})))] = False
                    return _dafny.Map(coll9_)
                return iife9_()
                

            return lambda1_

        init0_ = lambda0_(d_39_v1_)
        nw2_ = _dafny.Array(None, 16)
        for i0_0_ in range(nw2_.length(0)):
            nw2_[i0_0_] = init0_(i0_0_)
        d_45_v8_ = nw2_
        index0_ = default__.safeIndex(37, (d_44_v6_).length(0))
        (d_44_v6_)[index0_] = d_45_v8_
        d_49_v9_: bool
        d_49_v9_ = True
        d_50_v10_: _dafny.Seq
        d_50_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kngri"))
        d_51_v11_: D4
        d_51_v11_ = D4_DC10(d_45_v8_)
        pat_let_tv0_ = d_45_v8_
        index1_ = default__.safeIndex(37, (d_44_v6_).length(0))
        def iife10_(_pat_let0_0):
            def iife11_(d_52_dt__update__tmp_h0_):
                def iife12_(_pat_let1_0):
                    def iife13_(d_53_dt__update_hcf14_h0_):
                        return D4_DC10(d_53_dt__update_hcf14_h0_)
                    return iife13_(_pat_let1_0)
                return iife12_(pat_let_tv0_)
            return iife11_(_pat_let0_0)
        rhs0_ = (len(_dafny.Set({not(d_49_v9_), d_49_v9_, d_49_v9_}))) - ((d_38_v0_) - (default__.fm7((0) - ((d_42_v4_).f22), (0) - (d_38_v0_), globalState)))
        rhs1_ = d_43_v5_
        rhs2_ = default__.safeModuloInt(default__.fm7((d_42_v4_).f22, d_38_v0_, globalState), len((d_50_v10_) + (d_50_v10_)))
        rhs3_ = (iife10_(d_51_v11_)).cf14
        rhs4_ = d_49_v9_
        lhs0_ = globalState
        lhs1_ = d_44_v6_
        lhs2_ = default__.safeIndex(37, (d_44_v6_).length(0))
        lhs3_ = globalState
        d_38_v0_ = rhs0_
        d_43_v5_ = rhs1_
        lhs0_.f3 = rhs2_
        lhs1_[lhs2_] = rhs3_
        lhs3_.f14 = rhs4_
        d_50_v10_ = d_50_v10_
        d_54_v12_: _dafny.Map
        d_54_v12_ = _dafny.Map({d_38_v0_: d_49_v9_})
        d_55_v13_: _dafny.Map
        d_55_v13_ = d_54_v12_
        d_56_v14_: _dafny.Map
        d_56_v14_ = _dafny.Map({d_49_v9_: len((d_55_v13_))})
        d_57_v15_: _dafny.Set
        d_57_v15_ = _dafny.Set({len(default__.fm12((d_42_v4_).f22, d_38_v0_, d_50_v10_, globalState)), (d_42_v4_).f22, ((d_56_v14_)[d_49_v9_] if (d_49_v9_) in (d_56_v14_) else (d_42_v4_).f22), d_38_v0_})
        d_58_i1_: int
        d_58_i1_ = 0
        with _dafny.label("0"):
            while not((d_57_v15_).ispropersubset(_dafny.Set({d_38_v0_}))):
                with _dafny.c_label("0"):
                    if (d_58_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_58_i1_ = (d_58_i1_) + (1)
                    d_59_v16_: _dafny.Seq
                    d_59_v16_ = _dafny.SeqWithoutIsStrInference([d_43_v5_, d_43_v5_])
                    (globalState).f14 = (len(((d_59_v16_)[default__.safeIndex(607, len(d_59_v16_))]) | ((d_43_v5_).set(False, d_42_v4_)))) >= ((d_42_v4_).f22)
                    d_60_v17_: _dafny.Seq
                    d_60_v17_ = _dafny.SeqWithoutIsStrInference([d_49_v9_, d_49_v9_])
                    d_61_v18_: C1
                    nw3_ = C1()
                    nw3_.ctor__(929, d_60_v17_)
                    d_61_v18_ = nw3_
                    d_62_v19_: _dafny.Set
                    d_62_v19_ = _dafny.Set({d_61_v18_, d_61_v18_})
                    d_62_v19_ = d_62_v19_
                    d_63_v20_: _dafny.Map
                    d_63_v20_ = _dafny.Map({default__.fm13(d_39_v1_, d_49_v9_, globalState): default__.safeDivisionInt(552, (d_42_v4_).f22)})
                    d_64_v21_: _dafny.Seq
                    d_64_v21_ = _dafny.SeqWithoutIsStrInference([(d_42_v4_).f22])
                    d_63_v20_ = (d_63_v20_).set(d_64_v21_, (d_42_v4_).f22)
                    d_65_v22_: _dafny.Array
                    nw4_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                    d_65_v22_ = nw4_
                    index2_ = default__.safeIndex(484, (d_65_v22_).length(0))
                    (d_65_v22_)[index2_] = d_50_v10_
                    index3_ = default__.safeIndex(484, (d_65_v22_).length(0))
                    (d_65_v22_)[index3_] = d_50_v10_
                    pass
            pass
        if not(d_49_v9_):
            if d_49_v9_:
                d_66_v23_: _dafny.Seq
                d_66_v23_ = _dafny.SeqWithoutIsStrInference([d_49_v9_])
                d_67_v24_: C1
                nw5_ = C1()
                nw5_.ctor__((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_68_i2_ in range(default__.abs(674))]))) + ((d_42_v4_).f22), d_66_v23_)
                d_67_v24_ = nw5_
                rhs5_ = ((d_42_v4_).f22 if d_49_v9_ else (d_42_v4_).f22)
                rhs6_ = (d_42_v4_).f22
                lhs4_ = globalState
                lhs4_.f3 = rhs5_
                d_38_v0_ = rhs6_
                d_69_v25_: D3
                d_69_v25_ = D3_DC8((D3_DC8(d_67_v24_)).cf12)
                d_70_v26_: _dafny.MultiSet
                d_70_v26_ = _dafny.MultiSet([d_69_v25_])
                d_71_v27_: _dafny.Array
                nw6_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_71_v27_ = nw6_
                d_72_v28_: _dafny.Array
                def lambda2_(d_73_v9_):
                    def lambda3_(d_74_i3_):
                        return d_73_v9_

                    return lambda3_

                init1_ = lambda2_(d_49_v9_)
                nw7_ = _dafny.Array(None, 12)
                for i0_1_ in range(nw7_.length(0)):
                    nw7_[i0_1_] = init1_(i0_1_)
                d_72_v28_ = nw7_
                index4_ = default__.safeIndex(532, (d_71_v27_).length(0))
                (d_71_v27_)[index4_] = (d_72_v28_ if d_49_v9_ else d_72_v28_)
                index5_ = default__.safeIndex(532, (d_71_v27_).length(0))
                rhs7_ = _dafny.MultiSet([d_69_v25_])
                rhs8_ = d_72_v28_
                rhs9_ = not (d_49_v9_) or (not(d_49_v9_))
                rhs10_ = (d_39_v1_) not in (d_50_v10_)
                lhs5_ = d_71_v27_
                lhs6_ = default__.safeIndex(532, (d_71_v27_).length(0))
                lhs7_ = globalState
                lhs8_ = globalState
                d_70_v26_ = rhs7_
                lhs5_[lhs6_] = rhs8_
                lhs7_.f14 = rhs9_
                lhs8_.f14 = rhs10_
                (globalState).f7 = d_49_v9_
                d_75_v30_: D2
                def iife14_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in _dafny.IntegerRange(667, 551):
                        d_76_v29_: int = compr_10_
                        if ((667) <= (d_76_v29_)) and ((d_76_v29_) < (551)):
                            coll10_[default__.safeDivisionInt(d_76_v29_, (d_42_v4_).f22)] = d_49_v9_
                    return _dafny.Map(coll10_)
                d_75_v30_ = D2_DC7((0) - ((d_42_v4_).f22), len(iife14_()
))
                d_77_v31_: _dafny.Map
                d_77_v31_ = _dafny.Map({(_dafny.Map({d_49_v9_: d_39_v1_})).set(d_49_v9_, d_39_v1_): d_66_v23_})
                d_78_v32_: _dafny.Set
                d_78_v32_ = _dafny.Set({d_49_v9_})
                d_79_v33_: _dafny.Map
                d_79_v33_ = _dafny.Map({d_49_v9_: d_78_v32_})
                d_80_v34_: _dafny.Array
                nw8_ = _dafny.Array(None, 5)
                nw8_[int(0)] = (d_75_v30_).cf10
                nw8_[int(1)] = len((default__.fm14(globalState)) + (d_50_v10_))
                nw8_[int(2)] = ((d_42_v4_).f22) * ((d_42_v4_).f22)
                nw8_[int(3)] = ((d_67_v24_).fm1(d_77_v31_, globalState)) * (351)
                nw8_[int(4)] = len(d_79_v33_)
                d_80_v34_ = nw8_
                nw9_ = _dafny.Array(int(0), 12)
                d_80_v34_ = nw9_
            elif True:
                d_81_v35_: _dafny.Seq
                d_81_v35_ = _dafny.SeqWithoutIsStrInference([d_49_v9_])
                d_82_v36_: C1
                nw10_ = C1()
                nw10_.ctor__((d_42_v4_).f22, d_81_v35_)
                d_82_v36_ = nw10_
                d_83_v37_: _dafny.Seq
                d_83_v37_ = _dafny.SeqWithoutIsStrInference([d_82_v36_])
                d_84_v38_: _dafny.Seq
                d_84_v38_ = _dafny.SeqWithoutIsStrInference([(d_82_v36_).f20, (d_42_v4_).f22, (d_42_v4_).f22, (d_42_v4_).f22])
                d_85_v39_: _dafny.Seq
                d_85_v39_ = _dafny.SeqWithoutIsStrInference([d_84_v38_, d_84_v38_])
                d_86_v40_: _dafny.Seq
                d_86_v40_ = _dafny.SeqWithoutIsStrInference([d_50_v10_])
                d_87_v41_: _dafny.Array
                nw11_ = _dafny.Array(None, 15)
                nw11_[int(0)] = len(d_50_v10_)
                nw11_[int(1)] = d_38_v0_
                nw11_[int(2)] = -302
                nw11_[int(3)] = len(d_83_v37_)
                nw11_[int(4)] = (d_82_v36_).f20
                nw11_[int(5)] = (d_42_v4_).f22
                nw11_[int(6)] = (d_42_v4_).f22
                nw11_[int(7)] = (d_42_v4_).f22
                nw11_[int(8)] = ((d_42_v4_).f22) - (531)
                nw11_[int(9)] = d_38_v0_
                nw11_[int(10)] = d_38_v0_
                nw11_[int(11)] = (0) - (len(d_81_v35_))
                nw11_[int(12)] = (d_42_v4_).f22
                nw11_[int(13)] = (len(d_85_v39_)) - (len(d_86_v40_))
                nw11_[int(14)] = (0) - (((0) - ((d_42_v4_).f22)) - (366))
                d_87_v41_ = nw11_
                index6_ = default__.safeIndex(294, (d_87_v41_).length(0))
                (d_87_v41_)[index6_] = len(d_50_v10_)
                index7_ = default__.safeIndex(294, (d_87_v41_).length(0))
                (d_87_v41_)[index7_] = (0) - (default__.fm7(default__.safeDivisionInt((d_82_v36_).f20, (d_42_v4_).f22), (d_42_v4_).f22, globalState))
                d_88_v42_: _dafny.Array
                nw12_ = _dafny.Array(False, 5)
                d_88_v42_ = nw12_
                index8_ = default__.safeIndex(573, (d_88_v42_).length(0))
                (d_88_v42_)[index8_] = d_49_v9_
                index9_ = default__.safeIndex(573, (d_88_v42_).length(0))
                (d_88_v42_)[index9_] = (d_82_v36_).fm0((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsnskptul"))) + (_dafny.SeqWithoutIsStrInference([d_39_v1_ for d_89_i4_ in range(default__.abs(893))])), (d_42_v4_).f22, (d_42_v4_).f22, 994, globalState)
                index10_ = default__.safeIndex(573, (d_88_v42_).length(0))
                (d_88_v42_)[index10_] = not(d_49_v9_)
                d_90_v43_: _dafny.Map
                d_90_v43_ = _dafny.Map({d_49_v9_: (d_88_v42_)[default__.safeIndex(573, (d_88_v42_).length(0))]})
                index11_ = default__.safeIndex(294, (d_87_v41_).length(0))
                (d_87_v41_)[index11_] = default__.safeModuloInt(default__.safeModuloInt((d_41_v3_).cf9, (d_82_v36_).f20), default__.safeDivisionInt((d_82_v36_).f20, len(d_90_v43_)))
                (globalState).f11 = (d_81_v35_)[default__.safeIndex(d_38_v0_, len(d_81_v35_))]
            d_91_v44_: _dafny.Array
            nw13_ = _dafny.Array(int(0), 14)
            d_91_v44_ = nw13_
            index12_ = default__.safeIndex(577, (d_91_v44_).length(0))
            (d_91_v44_)[index12_] = d_38_v0_
            d_92_v45_: _dafny.MultiSet
            d_92_v45_ = _dafny.MultiSet([d_49_v9_])
            index13_ = default__.safeIndex(577, (d_91_v44_).length(0))
            (d_91_v44_)[index13_] = ((d_42_v4_).f22 if d_49_v9_ else default__.safeDivisionInt((d_92_v45_).cardinality, d_38_v0_))
            (globalState).f14 = d_49_v9_
            (globalState).f3 = (d_91_v44_)[default__.safeIndex(577, (d_91_v44_).length(0))]
            index14_ = default__.safeIndex(577, (d_91_v44_).length(0))
            (d_91_v44_)[index14_] = ((d_42_v4_).f22) + ((d_91_v44_)[default__.safeIndex(577, (d_91_v44_).length(0))])
        elif True:
            d_93_v46_: D1
            d_93_v46_ = D1_DC2((d_42_v4_).f22, d_49_v9_)
            (globalState).f5 = ((d_93_v46_).cf3 if d_49_v9_ else d_49_v9_)
            (globalState).f14 = default__.fm11((d_42_v4_).f22, globalState)
            d_94_v47_: _dafny.Seq
            d_94_v47_ = _dafny.SeqWithoutIsStrInference([True, d_49_v9_])
            d_95_v48_: C1
            nw14_ = C1()
            nw14_.ctor__(d_38_v0_, d_94_v47_)
            d_95_v48_ = nw14_
            d_96_v49_: _dafny.Map
            d_96_v49_ = _dafny.Map({d_95_v48_: _dafny.SeqWithoutIsStrInference([(d_95_v48_).f20, (d_42_v4_).f22])})
            d_97_v50_: _dafny.Seq
            d_97_v50_ = _dafny.SeqWithoutIsStrInference([-91, (d_42_v4_).f22])
            if (((d_95_v48_).f20) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihliafrys")))) if (d_38_v0_) not in (((d_96_v49_)[d_95_v48_] if (d_95_v48_) in (d_96_v49_) else d_97_v50_)) else d_49_v9_):
                d_98_v51_: _dafny.Map
                d_98_v51_ = _dafny.Map({d_42_v4_: (d_95_v48_).f20})
                d_38_v0_ = len(d_98_v51_)
                d_99_v52_: _dafny.Seq
                d_99_v52_ = _dafny.SeqWithoutIsStrInference([d_50_v10_, d_50_v10_])
                d_100_v53_: _dafny.MultiSet
                d_100_v53_ = _dafny.MultiSet([d_39_v1_])
                d_101_v54_: D6
                d_101_v54_ = D6_DC17(d_39_v1_)
                rhs11_ = (0) - (((d_42_v4_).f22) + (830))
                rhs12_ = ((D5_DC14((d_99_v52_)[default__.safeIndex((d_100_v53_).cardinality, len(d_99_v52_))])).cf19) + ((d_50_v10_).set(default__.safeIndex(d_38_v0_, len(d_50_v10_)), (d_101_v54_).cf25))
                rhs13_ = (d_42_v4_).f22
                lhs9_ = globalState
                lhs10_ = globalState
                lhs9_.f3 = rhs11_
                d_50_v10_ = rhs12_
                lhs10_.f3 = rhs13_
                d_102_v55_: _dafny.MultiSet
                d_102_v55_ = _dafny.MultiSet([d_42_v4_, d_42_v4_])
                d_103_v56_: _dafny.Map
                d_103_v56_ = _dafny.Map({d_49_v9_: True})
                d_104_v58_: _dafny.Seq
                d_104_v58_ = _dafny.SeqWithoutIsStrInference([d_94_v47_])
                d_105_v59_: _dafny.Map
                def iife15_():
                    coll11_ = _dafny.Map()
                    compr_11_: _dafny.Seq
                    for compr_11_ in (d_104_v58_).Elements:
                        d_106_v57_: _dafny.Seq = compr_11_
                        if (d_106_v57_) in (d_104_v58_):
                            coll11_[d_106_v57_] = (d_95_v48_).f20
                    return _dafny.Map(coll11_)
                d_105_v59_ = _dafny.Map({d_103_v56_: iife15_()
                })
                d_107_v60_: _dafny.Map
                d_107_v60_ = _dafny.Map({(d_95_v48_).f21: -797})
                d_108_v61_: _dafny.MultiSet
                d_108_v61_ = _dafny.MultiSet([((d_105_v59_)[default__.fm15((d_42_v4_).f22, d_50_v10_, False, (d_42_v4_).f22, globalState)] if (default__.fm15((d_42_v4_).f22, d_50_v10_, False, (d_42_v4_).f22, globalState)) in (d_105_v59_) else d_107_v60_), d_107_v60_])
                d_109_v62_: _dafny.Array
                nw15_ = _dafny.Array(None, 26)
                nw15_[int(0)] = (d_95_v48_).f20
                nw15_[int(1)] = len(d_50_v10_)
                nw15_[int(2)] = (d_102_v55_).cardinality
                nw15_[int(3)] = len(d_50_v10_)
                nw15_[int(4)] = (0) - ((d_95_v48_).f20)
                nw15_[int(5)] = -971
                nw15_[int(6)] = (0) - (d_38_v0_)
                nw15_[int(7)] = (d_42_v4_).f22
                nw15_[int(8)] = (d_95_v48_).f20
                nw15_[int(9)] = (d_42_v4_).f22
                nw15_[int(10)] = 570
                nw15_[int(11)] = -802
                nw15_[int(12)] = (d_95_v48_).f20
                nw15_[int(13)] = (d_42_v4_).f22
                nw15_[int(14)] = (d_95_v48_).f20
                nw15_[int(15)] = (d_42_v4_).f22
                nw15_[int(16)] = (d_42_v4_).f22
                nw15_[int(17)] = ((d_108_v61_)[_dafny.Map({d_94_v47_: 980})] if (_dafny.Map({d_94_v47_: 980})) in (d_108_v61_) else (d_42_v4_).f22)
                nw15_[int(18)] = (0) - ((d_42_v4_).f22)
                nw15_[int(19)] = (d_42_v4_).f22
                nw15_[int(20)] = (d_95_v48_).f20
                nw15_[int(21)] = (d_95_v48_).f20
                nw15_[int(22)] = len(d_56_v14_)
                nw15_[int(23)] = (d_42_v4_).f22
                nw15_[int(24)] = (d_42_v4_).f22
                nw15_[int(25)] = (0) - (((d_56_v14_)[d_49_v9_] if (d_49_v9_) in (d_56_v14_) else (d_95_v48_).f20))
                d_109_v62_ = nw15_
                d_110_v63_: _dafny.Map
                d_110_v63_ = _dafny.Map({d_49_v9_: d_109_v62_})
                d_110_v63_ = (d_110_v63_).set(d_49_v9_, d_109_v62_)
                d_111_v64_: _dafny.MultiSet
                d_111_v64_ = _dafny.MultiSet([(d_95_v48_).f20])
                d_111_v64_ = _dafny.MultiSet([d_38_v0_, (d_42_v4_).f22, (d_95_v48_).f20, (d_95_v48_).f20, d_38_v0_])
                d_112_v65_: _dafny.Seq
                d_112_v65_ = _dafny.SeqWithoutIsStrInference([d_42_v4_, d_42_v4_, d_42_v4_, d_42_v4_])
                d_113_v66_: C0
                d_113_v66_ = d_42_v4_
                d_114_v67_: _dafny.Array
                nw16_ = _dafny.Array(None, 23)
                nw16_[int(0)] = (d_112_v65_)[default__.safeIndex(len(d_50_v10_), len(d_112_v65_))]
                nw16_[int(1)] = (d_113_v66_)
                nw16_[int(2)] = d_42_v4_
                nw16_[int(3)] = d_42_v4_
                nw16_[int(4)] = d_42_v4_
                nw16_[int(5)] = d_42_v4_
                nw16_[int(6)] = d_42_v4_
                nw16_[int(7)] = d_42_v4_
                nw16_[int(8)] = (d_42_v4_ if d_49_v9_ else d_42_v4_)
                nw16_[int(9)] = d_42_v4_
                nw16_[int(10)] = d_42_v4_
                nw16_[int(11)] = d_42_v4_
                nw16_[int(12)] = d_42_v4_
                nw16_[int(13)] = d_42_v4_
                nw16_[int(14)] = d_42_v4_
                nw16_[int(15)] = d_42_v4_
                nw16_[int(16)] = (d_112_v65_)[default__.safeIndex((d_42_v4_).f22, len(d_112_v65_))]
                nw16_[int(17)] = d_42_v4_
                nw16_[int(18)] = d_42_v4_
                nw16_[int(19)] = d_42_v4_
                nw16_[int(20)] = d_42_v4_
                nw16_[int(21)] = d_42_v4_
                nw16_[int(22)] = d_42_v4_
                d_114_v67_ = nw16_
                index15_ = default__.safeIndex(918, (d_114_v67_).length(0))
                (d_114_v67_)[index15_] = d_42_v4_
                index16_ = default__.safeIndex(918, (d_114_v67_).length(0))
                (d_114_v67_)[index16_] = d_42_v4_
            elif True:
                d_115_v68_: _dafny.Array
                nw17_ = _dafny.Array(False, 13)
                d_115_v68_ = nw17_
                index17_ = default__.safeIndex(108, (d_115_v68_).length(0))
                (d_115_v68_)[index17_] = d_49_v9_
                index18_ = default__.safeIndex(108, (d_115_v68_).length(0))
                (d_115_v68_)[index18_] = (d_38_v0_) < ((0) - ((d_95_v48_).f20))
                d_116_v69_: _dafny.Array
                nw18_ = _dafny.Array(None, 1)
                nw18_[int(0)] = (0) - (-659)
                d_116_v69_ = nw18_
                d_117_v70_: _dafny.MultiSet
                d_117_v70_ = _dafny.MultiSet([(d_115_v68_)[default__.safeIndex(108, (d_115_v68_).length(0))], (d_115_v68_)[default__.safeIndex(108, (d_115_v68_).length(0))]])
                d_118_v71_: _dafny.Map
                d_118_v71_ = _dafny.Map({d_49_v9_: (d_42_v4_).fm4((d_117_v70_).cardinality, globalState)})
                d_119_v72_: _dafny.Map
                d_119_v72_ = _dafny.Map({d_118_v71_: (d_95_v48_).f21})
                index19_ = default__.safeIndex(28, (d_116_v69_).length(0))
                (d_116_v69_)[index19_] = (d_95_v48_).fm1(d_119_v72_, globalState)
                index20_ = default__.safeIndex(28, (d_116_v69_).length(0))
                (d_116_v69_)[index20_] = (d_95_v48_).f20
                (globalState).f3 = default__.safeDivisionInt((-49 if d_49_v9_ else 290), (d_42_v4_).f22)
                d_120_v73_: D5
                d_120_v73_ = D5_DC14(d_50_v10_)
                d_121_v74_: _dafny.Map
                d_121_v74_ = _dafny.Map({(d_115_v68_)[default__.safeIndex(108, (d_115_v68_).length(0))]: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btkawurs"))})
                d_122_v75_: _dafny.Array
                nw19_ = _dafny.Array(None, 19)
                nw19_[int(0)] = d_50_v10_
                nw19_[int(1)] = d_50_v10_
                nw19_[int(2)] = d_50_v10_
                nw19_[int(3)] = ((d_120_v73_).cf19) + (d_50_v10_)
                nw19_[int(4)] = (d_50_v10_) + (d_50_v10_)
                nw19_[int(5)] = ((d_50_v10_) + (d_50_v10_)).set(default__.safeIndex((d_97_v50_)[default__.safeIndex(988, len(d_97_v50_))], len((d_50_v10_) + (d_50_v10_))), d_39_v1_)
                nw19_[int(6)] = d_50_v10_
                nw19_[int(7)] = d_50_v10_
                nw19_[int(8)] = d_50_v10_
                nw19_[int(9)] = d_50_v10_
                nw19_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kspjtwcah"))) + (default__.fm14(globalState))
                nw19_[int(11)] = d_50_v10_
                nw19_[int(12)] = (d_50_v10_).set(default__.safeIndex((d_42_v4_).f22, len(d_50_v10_)), _dafny.CodePoint('j'))
                nw19_[int(13)] = d_50_v10_
                nw19_[int(14)] = d_50_v10_
                nw19_[int(15)] = d_50_v10_
                nw19_[int(16)] = d_50_v10_
                nw19_[int(17)] = d_50_v10_
                nw19_[int(18)] = ((d_121_v74_)[d_49_v9_] if (d_49_v9_) in (d_121_v74_) else d_50_v10_)
                d_122_v75_ = nw19_
                index21_ = default__.safeIndex(56, (d_122_v75_).length(0))
                (d_122_v75_)[index21_] = d_50_v10_
                index22_ = default__.safeIndex(56, (d_122_v75_).length(0))
                (d_122_v75_)[index22_] = (d_50_v10_ if False else d_50_v10_)
                d_123_v76_: _dafny.Map
                d_123_v76_ = _dafny.Map({d_116_v69_: d_49_v9_})
                d_124_v77_: _dafny.Map
                d_124_v77_ = _dafny.Map({default__.safeDivisionInt((d_116_v69_)[default__.safeIndex(28, (d_116_v69_).length(0))], (d_42_v4_).f22): d_123_v76_})
                d_124_v77_ = (d_124_v77_).set((d_95_v48_).f20, d_123_v76_)
            d_125_v78_: D8
            d_125_v78_ = D8_DC21((d_95_v48_).f21)
            d_126_v79_: _dafny.Map
            d_126_v79_ = _dafny.Map({not(d_49_v9_): d_49_v9_})
            d_127_v80_: _dafny.Set
            d_127_v80_ = _dafny.Set({d_42_v4_})
            d_128_v81_: _dafny.Map
            d_128_v81_ = _dafny.Map({((d_126_v79_)[default__.fm11(len(d_127_v80_), globalState)] if (default__.fm11(len(d_127_v80_), globalState)) in (d_126_v79_) else not(d_49_v9_)): d_57_v15_})
            d_129_v82_: _dafny.Map
            d_129_v82_ = _dafny.Map({(d_125_v78_).cf34: ((d_128_v81_)[d_49_v9_] if (d_49_v9_) in (d_128_v81_) else _dafny.Set({len(d_94_v47_)}))})
            d_129_v82_ = (d_129_v82_).set((d_94_v47_).set(default__.safeIndex((d_42_v4_).f22, len(d_94_v47_)), ((d_126_v79_)[d_49_v9_] if (d_49_v9_) in (d_126_v79_) else d_49_v9_)), d_57_v15_)
            d_130_v83_: _dafny.Array
            def lambda4_(d_131_v15_):
                def lambda5_(d_132_i5_):
                    return default__.safeModuloInt(d_132_i5_, len(d_131_v15_))

                return lambda5_

            init2_ = lambda4_(d_57_v15_)
            nw20_ = _dafny.Array(None, 3)
            for i0_2_ in range(nw20_.length(0)):
                nw20_[i0_2_] = init2_(i0_2_)
            d_130_v83_ = nw20_
            index23_ = default__.safeIndex(959, (d_130_v83_).length(0))
            (d_130_v83_)[index23_] = (d_42_v4_).f22
            index24_ = default__.safeIndex(959, (d_130_v83_).length(0))
            rhs14_ = default__.safeDivisionInt(default__.safeModuloInt((d_42_v4_).f22, (d_42_v4_).f22), 5)
            rhs15_ = (d_95_v48_).f20
            lhs11_ = globalState
            lhs12_ = d_130_v83_
            lhs13_ = default__.safeIndex(959, (d_130_v83_).length(0))
            lhs11_.f3 = rhs14_
            lhs12_[lhs13_] = rhs15_
        d_133_v84_: _dafny.MultiSet
        d_133_v84_ = _dafny.MultiSet([d_50_v10_, d_50_v10_])
        hi0_ = (d_42_v4_).f22
        for d_134_i6_ in range((d_133_v84_).cardinality, hi0_):
            d_135_v85_: _dafny.Array
            nw21_ = _dafny.Array(int(0), 20)
            d_135_v85_ = nw21_
            index25_ = default__.safeIndex(12, (d_135_v85_).length(0))
            (d_135_v85_)[index25_] = (d_42_v4_).f22
            index26_ = default__.safeIndex(12, (d_135_v85_).length(0))
            (d_135_v85_)[index26_] = default__.safeModuloInt(d_134_i6_, (d_42_v4_).f22)
            d_136_v86_: _dafny.Seq
            d_136_v86_ = _dafny.SeqWithoutIsStrInference([d_134_i6_])
            d_137_v87_: _dafny.Map
            d_137_v87_ = _dafny.Map({d_49_v9_: d_136_v86_})
            d_138_v88_: C0
            nw22_ = C0()
            nw22_.ctor__(default__.safeModuloInt(len(d_137_v87_), d_134_i6_))
            d_138_v88_ = nw22_
            d_139_v89_: _dafny.Array
            def lambda6_(d_140_i7_):
                return True

            init3_ = lambda6_
            nw23_ = _dafny.Array(None, 25)
            for i0_3_ in range(nw23_.length(0)):
                nw23_[i0_3_] = init3_(i0_3_)
            d_139_v89_ = nw23_
            index27_ = default__.safeIndex(64, (d_139_v89_).length(0))
            (d_139_v89_)[index27_] = d_49_v9_
            index28_ = default__.safeIndex(64, (d_139_v89_).length(0))
            (d_139_v89_)[index28_] = (d_49_v9_) or (d_49_v9_)
            d_141_v90_: C1
            nw24_ = C1()
            nw24_.ctor__(d_134_i6_, _dafny.SeqWithoutIsStrInference([True]))
            d_141_v90_ = nw24_
            d_142_v91_: _dafny.Map
            d_142_v91_ = _dafny.Map({D3_DC8(d_141_v90_): True})
            d_143_v92_: D3
            d_143_v92_ = D3_DC8(d_141_v90_)
            d_142_v91_ = (d_142_v91_).set(d_143_v92_, (True if (d_139_v89_)[default__.safeIndex(64, (d_139_v89_).length(0))] else (d_139_v89_)[default__.safeIndex(64, (d_139_v89_).length(0))]))
        d_144_v93_: _dafny.Set
        d_144_v93_ = _dafny.Set({not(d_49_v9_), d_49_v9_})
        r0 = d_144_v93_
        d_145_v94_: _dafny.Array
        nw25_ = _dafny.Array(int(0), 9)
        d_145_v94_ = nw25_
        d_146_v95_: _dafny.Seq
        d_146_v95_ = _dafny.SeqWithoutIsStrInference([580])
        d_147_v96_: _dafny.Array
        nw26_ = _dafny.Array(_dafny.Map({}), 10)
        d_147_v96_ = nw26_
        d_148_v97_: D8
        d_148_v97_ = D8_DC22(d_145_v94_, d_146_v95_, d_38_v0_, d_147_v96_, d_49_v9_)
        r1 = (d_49_v9_ if (d_148_v97_).cf39 else d_49_v9_)
        r2 = d_54_v12_
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_149_v0_: bool
        d_149_v0_ = False
        d_150_v1_: _dafny.Array
        nw27_ = _dafny.Array(_dafny.Array(None, 0), 17)
        d_150_v1_ = nw27_
        d_151_v2_: int
        d_151_v2_ = 204
        d_152_v3_: _dafny.Array
        nw28_ = _dafny.Array(False, 22)
        d_152_v3_ = nw28_
        d_153_v4_: _dafny.Map
        d_153_v4_ = _dafny.Map({d_151_v2_: d_152_v3_})
        d_154_v5_: _dafny.Seq
        d_154_v5_ = _dafny.SeqWithoutIsStrInference([True])
        d_155_v6_: _dafny.Map
        d_155_v6_ = _dafny.Map({d_149_v0_: d_151_v2_})
        d_156_v7_: _dafny.Seq
        d_156_v7_ = _dafny.SeqWithoutIsStrInference([d_149_v0_, (d_154_v5_)[default__.safeIndex(len(d_155_v6_), len(d_154_v5_))]])
        d_157_v8_: _dafny.Set
        d_157_v8_ = _dafny.Set({d_151_v2_})
        d_158_v9_: _dafny.Set
        d_158_v9_ = _dafny.Set({d_149_v0_})
        d_159_v10_: _dafny.Array
        nw29_ = _dafny.Array(_dafny.Array(None, 0), 16)
        d_159_v10_ = nw29_
        d_160_globalState_: GlobalState
        nw30_ = GlobalState()
        nw30_.ctor__((d_150_v1_ if d_149_v0_ else d_150_v1_), 872, False, 91, 751, True, -494, True, 204, d_153_v4_, False, True, d_156_v7_, d_157_v8_, False, -442, d_158_v9_, 341, 699, d_159_v10_)
        d_160_globalState_ = nw30_
        (d_160_globalState_).f3 = len(((d_158_v9_ if d_149_v0_ else d_158_v9_)).intersection(d_158_v9_))
        d_161_v11_: _dafny.Array
        nw31_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
        d_161_v11_ = nw31_
        d_162_v12_: _dafny.Seq
        d_162_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skhwrlhb"))
        index29_ = default__.safeIndex(804, (d_161_v11_).length(0))
        (d_161_v11_)[index29_] = (d_162_v12_) + (_dafny.SeqWithoutIsStrInference([(d_162_v12_)[default__.safeIndex(d_151_v2_, len(d_162_v12_))] for d_163_i0_ in range(default__.abs(-478))]))
        index30_ = default__.safeIndex(804, (d_161_v11_).length(0))
        (d_161_v11_)[index30_] = d_162_v12_
        if d_149_v0_:
            d_164_v13_: _dafny.Array
            def lambda7_(d_165_i1_):
                return (d_165_i1_) * (553)

            init4_ = lambda7_
            nw32_ = _dafny.Array(None, 21)
            for i0_4_ in range(nw32_.length(0)):
                nw32_[i0_4_] = init4_(i0_4_)
            d_164_v13_ = nw32_
            d_166_v14_: _dafny.Seq
            d_166_v14_ = _dafny.SeqWithoutIsStrInference([d_162_v12_])
            rhs16_ = d_164_v13_
            rhs17_ = (0) - (default__.safeModuloInt(len(d_166_v14_), (d_151_v2_) + (d_151_v2_)))
            rhs18_ = not(True)
            lhs14_ = d_160_globalState_
            d_164_v13_ = rhs16_
            lhs14_.f3 = rhs17_
            d_149_v0_ = rhs18_
            d_167_v15_: _dafny.MultiSet
            d_167_v15_ = _dafny.MultiSet([d_151_v2_])
            d_168_v16_: _dafny.Set
            d_169_v17_: bool
            d_170_v18_: _dafny.Map
            out0_: _dafny.Set
            out1_: bool
            out2_: _dafny.Map
            out0_, out1_, out2_ = default__.m0(d_167_v15_, d_160_globalState_)
            d_168_v16_ = out0_
            d_169_v17_ = out1_
            d_170_v18_ = out2_
            d_171_v19_: _dafny.Seq
            d_171_v19_ = _dafny.SeqWithoutIsStrInference([d_151_v2_])
            d_172_v20_: _dafny.Map
            d_172_v20_ = _dafny.Map({d_171_v19_: d_151_v2_})
            d_173_v21_: _dafny.Map
            d_173_v21_ = _dafny.Map({(32) + (-306): ((d_172_v20_)[_dafny.SeqWithoutIsStrInference([d_151_v2_])] if (_dafny.SeqWithoutIsStrInference([d_151_v2_])) in (d_172_v20_) else d_151_v2_)})
            d_173_v21_ = (d_173_v21_).set((0) - (d_151_v2_), len((d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))]))
            d_174_v22_: _dafny.Map
            d_174_v22_ = _dafny.Map({not(d_149_v0_): d_149_v0_})
            d_155_v6_ = (d_155_v6_).set((d_151_v2_) < (len(d_174_v22_)), d_151_v2_)
            d_175_v23_: C1
            nw33_ = C1()
            nw33_.ctor__(362, d_156_v7_)
            d_175_v23_ = nw33_
        elif True:
            d_151_v2_ = d_151_v2_
            d_176_v24_: C1
            nw34_ = C1()
            nw34_.ctor__(d_151_v2_, _dafny.SeqWithoutIsStrInference([d_149_v0_]))
            d_176_v24_ = nw34_
            d_177_v25_: _dafny.Map
            d_177_v25_ = _dafny.Map({d_176_v24_: -39})
            d_178_v26_: _dafny.MultiSet
            d_178_v26_ = _dafny.MultiSet([len(d_177_v25_), len(d_158_v9_), (d_176_v24_).f20, d_151_v2_])
            d_179_v27_: _dafny.Set
            d_180_v28_: bool
            d_181_v29_: _dafny.Map
            out3_: _dafny.Set
            out4_: bool
            out5_: _dafny.Map
            out3_, out4_, out5_ = default__.m0(d_178_v26_, d_160_globalState_)
            d_179_v27_ = out3_
            d_180_v28_ = out4_
            d_181_v29_ = out5_
            d_151_v2_ = default__.safeModuloInt(default__.safeModuloInt((d_176_v24_).f20, (d_176_v24_).f20), len(d_158_v9_))
            (d_160_globalState_).f11 = (default__.safeDivisionInt(d_151_v2_, (d_176_v24_).f20)) != ((d_176_v24_).f20)
            d_182_v30_: _dafny.Array
            def lambda8_(d_183_i2_):
                return (d_183_i2_) * (-99)

            init5_ = lambda8_
            nw35_ = _dafny.Array(None, 21)
            for i0_5_ in range(nw35_.length(0)):
                nw35_[i0_5_] = init5_(i0_5_)
            d_182_v30_ = nw35_
            (d_160_globalState_).f7 = (d_176_v24_).fm0(d_162_v12_, (d_176_v24_).f20, default__.safeDivisionInt(d_151_v2_, len(_dafny.Set({d_182_v30_, d_182_v30_, d_182_v30_, d_182_v30_, d_182_v30_}))), d_151_v2_, d_160_globalState_)
        d_184_v31_: _dafny.MultiSet
        d_184_v31_ = _dafny.MultiSet([d_149_v0_, d_149_v0_, d_149_v0_])
        d_185_i3_: int
        d_185_i3_ = 0
        with _dafny.label("1"):
            while (d_154_v5_)[default__.safeIndex(((d_184_v31_)[d_149_v0_] if (d_149_v0_) in (d_184_v31_) else d_151_v2_), len(d_154_v5_))]:
                with _dafny.c_label("1"):
                    if (d_185_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_185_i3_ = (d_185_i3_) + (1)
                    d_186_v33_: _dafny.Seq
                    d_186_v33_ = _dafny.SeqWithoutIsStrInference([d_151_v2_, d_151_v2_])
                    d_187_v34_: _dafny.Map
                    def iife16_():
                        coll12_ = _dafny.Map()
                        compr_12_: int
                        for compr_12_ in (d_186_v33_).Elements:
                            d_188_v32_: int = compr_12_
                            if (d_188_v32_) in (d_186_v33_):
                                coll12_[(d_188_v32_) + (d_151_v2_)] = d_149_v0_
                        return _dafny.Map(coll12_)
                    d_187_v34_ = iife16_()
                    
                    d_187_v34_ = default__.fm6((0) - (default__.fm7(((d_184_v31_)[d_149_v0_] if (d_149_v0_) in (d_184_v31_) else d_151_v2_), d_151_v2_, d_160_globalState_)), d_160_globalState_)
                    d_189_v35_: _dafny.Map
                    d_189_v35_ = _dafny.Map({d_187_v34_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "syahspiq"))})
                    d_189_v35_ = (d_189_v35_).set(d_187_v34_, d_162_v12_)
                    d_190_v36_: _dafny.Array
                    nw36_ = _dafny.Array(_dafny.Array(None, 0), 2)
                    d_190_v36_ = nw36_
                    d_190_v36_ = d_190_v36_
                    (d_160_globalState_).f14 = (default__.safeModuloInt(d_151_v2_, (0) - (d_151_v2_))) == ((default__.fm8(d_160_globalState_)).cf2)
                    pass
            pass
        d_191_v37_: _dafny.Array
        nw37_ = _dafny.Array(int(0), 7)
        d_191_v37_ = nw37_
        index31_ = default__.safeIndex(792, (d_191_v37_).length(0))
        (d_191_v37_)[index31_] = d_151_v2_
        d_192_v38_: _dafny.Seq
        d_192_v38_ = _dafny.SeqWithoutIsStrInference([d_162_v12_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsf")), ((d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))]) + (d_162_v12_), d_162_v12_, d_162_v12_])
        index32_ = default__.safeIndex(804, (d_161_v11_).length(0))
        index33_ = default__.safeIndex(792, (d_191_v37_).length(0))
        rhs19_ = (d_192_v38_)[default__.safeIndex((0) - (d_151_v2_), len(d_192_v38_))]
        rhs20_ = default__.safeDivisionInt(d_151_v2_, d_151_v2_)
        lhs15_ = d_161_v11_
        lhs16_ = default__.safeIndex(804, (d_161_v11_).length(0))
        lhs17_ = d_191_v37_
        lhs18_ = default__.safeIndex(792, (d_191_v37_).length(0))
        lhs15_[lhs16_] = rhs19_
        lhs17_[lhs18_] = rhs20_
        d_193_v39_: D1
        d_193_v39_ = D1_DC2(d_151_v2_, d_149_v0_)
        pat_let_tv1_ = d_149_v0_
        pat_let_tv2_ = d_191_v37_
        pat_let_tv3_ = d_191_v37_
        pat_let_tv4_ = d_151_v2_
        pat_let_tv5_ = d_151_v2_
        pat_let_tv6_ = d_191_v37_
        pat_let_tv7_ = d_191_v37_
        def lambda9_(source2_):
            if source2_.is_DC2:
                d_194___mcc_h0_ = source2_.cf2
                d_195___mcc_h1_ = source2_.cf3
                d_196_cf3_ = d_195___mcc_h1_
                d_197_cf2_ = d_194___mcc_h0_
                return not (pat_let_tv1_) or (d_196_cf3_)
            elif source2_.is_DC1:
                d_198___mcc_h2_ = source2_.cf1
                d_199_cf1_ = d_198___mcc_h2_
                return ((pat_let_tv3_)[default__.safeIndex(792, (pat_let_tv2_).length(0))]) <= (pat_let_tv4_)
            elif True:
                d_200___mcc_h3_ = source2_.cf4
                d_201_cf4_ = d_200___mcc_h3_
                return (pat_let_tv5_) >= ((pat_let_tv7_)[default__.safeIndex(792, (pat_let_tv6_).length(0))])

        (d_160_globalState_).f7 = lambda9_(d_193_v39_)
        d_202_v40_: _dafny.Map
        d_202_v40_ = _dafny.Map({(d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))]: d_149_v0_})
        d_203_v41_: _dafny.Map
        d_203_v41_ = d_202_v40_
        pat_let_tv8_ = d_151_v2_
        def lambda10_(source3_):
            d_204___mcc_h4_ = source3_
            d_205_cf0_ = d_204___mcc_h4_
            return pat_let_tv8_

        (d_160_globalState_).f3 = lambda10_(d_203_v41_)
        d_206_v42_: str
        d_206_v42_ = _dafny.CodePoint('d')
        d_207_v43_: _dafny.Set
        d_207_v43_ = _dafny.Set({d_206_v42_, d_206_v42_})
        d_208_v44_: _dafny.Map
        d_208_v44_ = _dafny.Map({d_151_v2_: d_207_v43_})
        d_209_v45_: _dafny.MultiSet
        d_209_v45_ = _dafny.MultiSet([(d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))], len(d_158_v9_), (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))], 26, len((d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))])])
        (d_160_globalState_).f3 = (default__.safeModuloInt(d_151_v2_, (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))])) * (default__.fm7(len(((d_208_v44_)[((d_184_v31_)[d_149_v0_] if (d_149_v0_) in (d_184_v31_) else (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))])] if (((d_184_v31_)[d_149_v0_] if (d_149_v0_) in (d_184_v31_) else (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))])) in (d_208_v44_) else d_207_v43_)), ((d_209_v45_)[len((d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))])] if (len((d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))])) in (d_209_v45_) else (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))]), d_160_globalState_))
        d_210_v46_: _dafny.Map
        d_210_v46_ = _dafny.Map({d_149_v0_: d_209_v45_})
        d_211_v47_: _dafny.Seq
        d_211_v47_ = _dafny.SeqWithoutIsStrInference([d_151_v2_])
        d_212_v48_: D2
        d_212_v48_ = D2_DC4(_dafny.MultiSet(d_211_v47_))
        pat_let_tv9_ = d_151_v2_
        d_213_v49_: _dafny.Seq
        def iife17_(_pat_let2_0):
            def iife18_(d_214_dt__update__tmp_h0_):
                def iife19_(_pat_let3_0):
                    def iife20_(d_215_dt__update_hcf5_h0_):
                        return D2_DC4(d_215_dt__update_hcf5_h0_)
                    return iife20_(_pat_let3_0)
                return iife19_(_dafny.MultiSet([pat_let_tv9_, 765]))
            return iife18_(_pat_let2_0)
        d_213_v49_ = _dafny.SeqWithoutIsStrInference([(iife17_(d_212_v48_)).cf5, d_209_v45_])
        d_216_v50_: _dafny.Map
        d_216_v50_ = _dafny.Map({False: d_162_v12_})
        (d_160_globalState_).f14 = ((d_209_v45_) | (d_209_v45_)) != (((d_210_v46_)[d_149_v0_] if (d_149_v0_) in (d_210_v46_) else (d_213_v49_)[default__.safeIndex(len(((d_216_v50_)[d_149_v0_] if (d_149_v0_) in (d_216_v50_) else (d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))])), len(d_213_v49_))]))
        hi1_ = ((d_155_v6_)[d_149_v0_] if (d_149_v0_) in (d_155_v6_) else d_151_v2_)
        for d_217_i4_ in range((d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))], hi1_):
            d_218_v51_: _dafny.Map
            d_218_v51_ = _dafny.Map({d_202_v40_: (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))]})
            d_219_v52_: _dafny.Seq
            d_219_v52_ = _dafny.SeqWithoutIsStrInference([d_218_v51_])
            d_220_v53_: _dafny.Map
            d_220_v53_ = _dafny.Map({d_151_v2_: d_151_v2_})
            d_221_v54_: _dafny.Set
            d_222_v55_: bool
            d_223_v56_: _dafny.Map
            out6_: _dafny.Set
            out7_: bool
            out8_: _dafny.Map
            out6_, out7_, out8_ = default__.m0(default__.fm9(False, (d_162_v12_)[default__.safeIndex((d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))], len(d_162_v12_))], (d_219_v52_)[default__.safeIndex(d_151_v2_, len(d_219_v52_))], d_220_v53_, d_160_globalState_), d_160_globalState_)
            d_221_v54_ = out6_
            d_222_v55_ = out7_
            d_223_v56_ = out8_
            (d_160_globalState_).f3 = (0) - (default__.fm7(d_151_v2_, len((d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))]), d_160_globalState_))
            (d_160_globalState_).f7 = not(d_149_v0_)
            d_224_v57_: _dafny.Map
            d_224_v57_ = _dafny.Map({True: (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skjwiqufa")))) != ((d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))])})
            d_224_v57_ = (d_224_v57_).set(d_149_v0_, d_149_v0_)
        d_225_v58_: D1
        d_225_v58_ = D1_DC2(default__.fm7(d_151_v2_, d_151_v2_, d_160_globalState_), (d_154_v5_)[default__.safeIndex((d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))], len(d_154_v5_))])
        d_226_v59_: D1
        d_226_v59_ = D1_DC3(d_225_v58_)
        d_227_v60_: _dafny.Map
        d_227_v60_ = _dafny.Map({d_226_v59_: d_206_v42_})
        d_228_v61_: _dafny.Map
        d_228_v61_ = _dafny.Map({d_227_v60_: (d_149_v0_) or (True)})
        d_228_v61_ = (d_228_v61_).set(d_227_v60_, (d_149_v0_ if d_149_v0_ else d_149_v0_))
        d_229_v62_: _dafny.Map
        d_229_v62_ = _dafny.Map({d_206_v42_: d_191_v37_})
        hi2_ = len(d_229_v62_)
        for d_230_i5_ in range(d_151_v2_, hi2_):
            d_231_v63_: C1
            nw38_ = C1()
            nw38_.ctor__(d_151_v2_, (d_156_v7_) + (d_156_v7_))
            d_231_v63_ = nw38_
            (d_160_globalState_).f3 = (d_230_i5_) * ((d_231_v63_).f20)
            (d_160_globalState_).f11 = d_149_v0_
            d_232_v64_: D1
            d_232_v64_ = D1_DC1(d_230_i5_)
            source4_ = d_232_v64_
            if source4_.is_DC2:
                d_233___mcc_h5_ = source4_.cf2
                d_234___mcc_h6_ = source4_.cf3
                d_235_cf3_ = d_234___mcc_h6_
                d_236_cf2_ = d_233___mcc_h5_
                d_237_v65_: D2
                d_237_v65_ = D2_DC6((d_231_v63_).f20)
                d_238_v66_: _dafny.Map
                d_238_v66_ = _dafny.Map({d_231_v63_: (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))]})
                d_239_v67_: _dafny.Seq
                d_239_v67_ = _dafny.SeqWithoutIsStrInference([d_238_v66_, d_238_v66_])
                d_240_v68_: _dafny.Map
                d_240_v68_ = _dafny.Map({(0) - (d_151_v2_): d_206_v42_})
                d_241_v69_: _dafny.Seq
                d_241_v69_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_239_v67_): d_206_v42_}), d_240_v68_, d_240_v68_, _dafny.Map({(d_231_v63_).f20: d_206_v42_})])
                pat_let_tv10_ = d_241_v69_
                d_242_v70_: _dafny.Array
                nw39_ = _dafny.Array(None, 18)
                nw39_[int(0)] = d_158_v9_
                nw39_[int(1)] = d_158_v9_
                nw39_[int(2)] = d_158_v9_
                nw39_[int(3)] = d_158_v9_
                nw39_[int(4)] = d_158_v9_
                nw39_[int(5)] = default__.fm10(d_237_v65_, d_160_globalState_)
                nw39_[int(6)] = d_158_v9_
                nw39_[int(7)] = (d_158_v9_) - (d_158_v9_)
                def iife21_(_pat_let4_0):
                    def iife22_(d_243_dt__update__tmp_h1_):
                        def iife23_(_pat_let5_0):
                            def iife24_(d_244_dt__update_hcf9_h0_):
                                return D2_DC6(d_244_dt__update_hcf9_h0_)
                            return iife24_(_pat_let5_0)
                        return iife23_((_dafny.MultiSet(pat_let_tv10_)).cardinality)
                    return iife22_(_pat_let4_0)
                nw39_[int(8)] = default__.fm10(iife21_(d_237_v65_), d_160_globalState_)
                nw39_[int(9)] = _dafny.Set({d_149_v0_, d_149_v0_, d_149_v0_})
                nw39_[int(10)] = (d_158_v9_) - (d_158_v9_)
                nw39_[int(11)] = d_158_v9_
                nw39_[int(12)] = d_158_v9_
                nw39_[int(13)] = d_158_v9_
                nw39_[int(14)] = d_158_v9_
                nw39_[int(15)] = d_158_v9_
                nw39_[int(16)] = d_158_v9_
                nw39_[int(17)] = default__.fm10(d_237_v65_, d_160_globalState_)
                d_242_v70_ = nw39_
                index34_ = default__.safeIndex(658, (d_242_v70_).length(0))
                (d_242_v70_)[index34_] = (_dafny.Set({not(d_235_cf3_)})) - (d_158_v9_)
                d_245_v71_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_245_v71_ = nw40_
                index35_ = default__.safeIndex(735, (d_245_v71_).length(0))
                (d_245_v71_)[index35_] = d_161_v11_
                index36_ = default__.safeIndex(792, (d_191_v37_).length(0))
                index37_ = default__.safeIndex(658, (d_242_v70_).length(0))
                index38_ = default__.safeIndex(735, (d_245_v71_).length(0))
                rhs21_ = (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))]
                rhs22_ = ((0) - ((d_231_v63_).f20)) + (107)
                rhs23_ = d_158_v9_
                rhs24_ = d_161_v11_
                lhs19_ = d_160_globalState_
                lhs20_ = d_191_v37_
                lhs21_ = default__.safeIndex(792, (d_191_v37_).length(0))
                lhs22_ = d_242_v70_
                lhs23_ = default__.safeIndex(658, (d_242_v70_).length(0))
                lhs24_ = d_245_v71_
                lhs25_ = default__.safeIndex(735, (d_245_v71_).length(0))
                lhs19_.f3 = rhs21_
                lhs20_[lhs21_] = rhs22_
                lhs22_[lhs23_] = rhs23_
                lhs24_[lhs25_] = rhs24_
                d_155_v6_ = ((d_155_v6_).set(d_149_v0_, d_236_cf2_)) | (d_155_v6_)
                (d_160_globalState_).f7 = d_149_v0_
                (d_160_globalState_).f14 = (d_231_v63_).fm0((d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))], (17) * (d_151_v2_), d_151_v2_, ((d_209_v45_).cardinality) + ((d_231_v63_).f20), d_160_globalState_)
            elif source4_.is_DC1:
                d_246___mcc_h7_ = source4_.cf1
                d_247_cf1_ = d_246___mcc_h7_
                d_212_v48_ = d_212_v48_
                d_248_v72_: _dafny.Map
                d_248_v72_ = _dafny.Map({D1_DC1(len(_dafny.SeqWithoutIsStrInference([d_152_v3_, d_152_v3_]))): d_211_v47_})
                d_248_v72_ = (d_248_v72_).set(d_232_v64_, (d_211_v47_) + (d_211_v47_))
                (d_160_globalState_).f11 = ((d_158_v9_).intersection(d_158_v9_)) != ((d_158_v9_) | (d_158_v9_))
                index39_ = default__.safeIndex(804, (d_161_v11_).length(0))
                (d_161_v11_)[index39_] = (d_161_v11_)[default__.safeIndex(804, (d_161_v11_).length(0))]
            elif True:
                d_249___mcc_h8_ = source4_.cf4
                d_250_cf4_ = d_249___mcc_h8_
                d_202_v40_ = d_202_v40_
                index40_ = default__.safeIndex(792, (d_191_v37_).length(0))
                rhs25_ = d_149_v0_
                rhs26_ = d_230_i5_
                rhs27_ = (d_149_v0_) and ((d_206_v42_) not in (d_162_v12_))
                rhs28_ = (0) - ((len(d_202_v40_)) * (len(d_162_v12_)))
                lhs26_ = d_160_globalState_
                lhs27_ = d_191_v37_
                lhs28_ = default__.safeIndex(792, (d_191_v37_).length(0))
                lhs29_ = d_160_globalState_
                lhs26_.f7 = rhs25_
                lhs27_[lhs28_] = rhs26_
                d_149_v0_ = rhs27_
                lhs29_.f3 = rhs28_
                d_156_v7_ = d_154_v5_
                d_156_v7_ = d_156_v7_
        d_251_v73_: _dafny.Array
        nw41_ = _dafny.Array(_dafny.CodePoint('D'), 23)
        d_251_v73_ = nw41_
        d_251_v73_ = d_251_v73_
        (d_160_globalState_).f3 = ((d_184_v31_).intersection((d_184_v31_) - (d_184_v31_))).cardinality
        d_252_v74_: _dafny.Set
        d_253_v75_: bool
        d_254_v76_: _dafny.Map
        out9_: _dafny.Set
        out10_: bool
        out11_: _dafny.Map
        out9_, out10_, out11_ = default__.m0(_dafny.MultiSet([857, (0) - ((d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))]), (d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))], len(d_192_v38_), d_151_v2_]), d_160_globalState_)
        d_252_v74_ = out9_
        d_253_v75_ = out10_
        d_254_v76_ = out11_
        d_255_v77_: C1
        nw42_ = C1()
        nw42_.ctor__(default__.fm7((d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))], -177, d_160_globalState_), d_154_v5_)
        d_255_v77_ = nw42_
        d_256_v78_: D3
        d_256_v78_ = D3_DC8(d_255_v77_)
        d_257_v79_: _dafny.Set
        d_257_v79_ = _dafny.Set({(d_256_v78_).cf12})
        rhs29_ = True
        rhs30_ = ((d_191_v37_)[default__.safeIndex(792, (d_191_v37_).length(0))]) + (d_151_v2_)
        rhs31_ = (0) - (((d_184_v31_)[d_149_v0_] if (d_149_v0_) in (d_184_v31_) else len(d_257_v79_)))
        lhs30_ = d_160_globalState_
        d_149_v0_ = rhs29_
        d_151_v2_ = rhs30_
        lhs30_.f3 = rhs31_
        _dafny.print(_dafny.string_of(d_149_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_151_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_153_v4_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v5_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_) == (_dafny.Map({False: 204}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v7_) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v8_) == (_dafny.Set({204}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v9_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_160_globalState_).f9)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_160_globalState_).f12) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_160_globalState_).f13) == (_dafny.Set({204}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_160_globalState_).f16) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_161_v11_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_162_v12_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v31_) == (_dafny.MultiSet([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_185_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_191_v37_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_192_v38_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skhwrlhb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skhwrlhbskhwrlhb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skhwrlhb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skhwrlhb"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v39_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v39_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_202_v40_) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_203_v41_)) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_206_v42_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_207_v43_) == (_dafny.Set({_dafny.CodePoint('d')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v44_) == (_dafny.Map({0: _dafny.Set({_dafny.CodePoint('d')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v45_) == (_dafny.MultiSet([0, 0, 1, 26, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_v46_) == (_dafny.Map({False: _dafny.MultiSet([0, 0, 1, 26, 8])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v47_) == (_dafny.SeqWithoutIsStrInference([0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_212_v48_).cf5) == (_dafny.MultiSet([0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v49_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([0, 765]), _dafny.MultiSet([0, 0, 1, 26, 8])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v50_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skhwrlhb"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v58_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v58_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_v59_).cf4).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_v59_).cf4).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v60_) == (_dafny.Map({D1_DC3(D1_DC2(7608, True)): _dafny.CodePoint('d')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v61_) == (_dafny.Map({_dafny.Map({D1_DC3(D1_DC2(7608, True)): _dafny.CodePoint('d')}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_229_v62_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v74_) == (_dafny.Set({False, True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_253_v75_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v76_) == (_dafny.Map({7257: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_v77_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_v77_).f21) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v78_).cf12).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_256_v78_).cf12).f21) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_257_v79_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
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
        return lambda: D1_DC2(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(False, False, _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC11(D4, NamedTuple('DC11', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC15(D5, NamedTuple('DC15', [('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({self.cf19.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC18(D6, NamedTuple('DC18', [('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)

class D7_DC20(D7, NamedTuple('DC20', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC22(_dafny.Array(None, 0), _dafny.Seq({}), int(0), _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC22(D8, NamedTuple('DC22', [('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f3: int = int(0)
        self.f5: bool = False
        self.f7: bool = False
        self.f11: bool = False
        self.f14: bool = False
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: int = int(0)
        self._f2: bool = False
        self._f4: int = int(0)
        self._f6: int = int(0)
        self._f8: int = int(0)
        self._f9: _dafny.Map = _dafny.Map({})
        self._f10: bool = False
        self._f12: _dafny.Seq = _dafny.Seq({})
        self._f13: _dafny.Set = _dafny.Set({})
        self._f15: int = int(0)
        self._f16: _dafny.Set = _dafny.Set({})
        self._f17: int = int(0)
        self._f18: int = int(0)
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19

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
    def f6(self):
        return self._f6
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
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19

class C0:
    def  __init__(self):
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f22):
        (self)._f22 = f22

    def fm3(self, p0, globalState):
        return _dafny.Map({(self).f22: (False) in (_dafny.MultiSet([True, not(not(True))]))})

    def fm4(self, p0, globalState):
        return _dafny.CodePoint('g')

    @property
    def f22(self):
        return self._f22

class C1:
    def  __init__(self):
        self._f20: int = int(0)
        self._f21: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f20, f21):
        (self)._f20 = f20
        (self)._f21 = f21

    def fm0(self, p0, p1, p2, p3, globalState):
        return not (False) or (((self).f21)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_258_i0_ in range(default__.abs(564))])), len((self).f21))])

    def fm1(self, p0, globalState):
        return ((self).f20) - ((self).f20)

    def m1(self, globalState):
        d_259_v0_: bool
        d_259_v0_ = False
        d_260_v1_: _dafny.Map
        d_260_v1_ = _dafny.Map({(self).f20: d_259_v0_})
        d_261_v3_: _dafny.Seq
        d_261_v3_ = _dafny.SeqWithoutIsStrInference([50])
        d_262_v5_: _dafny.Map
        d_262_v5_ = _dafny.Map({len((self).f21): (self).f20})
        d_263_v6_: _dafny.Map
        def iife25_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in (d_262_v5_).keys.Elements:
                d_264_v4_: int = compr_13_
                if (d_264_v4_) in (d_262_v5_):
                    coll13_[default__.safeDivisionInt(d_264_v4_, (self).f20)] = d_259_v0_
            return _dafny.Map(coll13_)
        d_263_v6_ = iife25_()
        
        d_265_v7_: _dafny.Map
        d_265_v7_ = (d_263_v6_)
        d_266_v9_: _dafny.Seq
        def iife26_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(207, 401):
                d_267_v8_: int = compr_14_
                if ((207) <= (d_267_v8_)) and ((d_267_v8_) < (401)):
                    coll14_[default__.safeDivisionInt(d_267_v8_, (self).f20)] = d_259_v0_
            return _dafny.Map(coll14_)
        d_266_v9_ = _dafny.SeqWithoutIsStrInference([iife26_()
        ])
        d_268_v10_: _dafny.Array
        nw43_ = _dafny.Array(None, 20)
        nw43_[int(0)] = (d_260_v1_) | (d_260_v1_)
        nw43_[int(1)] = (d_260_v1_) | (d_260_v1_)
        def iife27_():
            coll15_ = _dafny.Map()
            compr_15_: int
            for compr_15_ in (d_261_v3_).Elements:
                d_269_v2_: int = compr_15_
                if (d_269_v2_) in (d_261_v3_):
                    coll15_[(d_269_v2_) + ((self).f20)] = d_259_v0_
            return _dafny.Map(coll15_)
        nw43_[int(2)] = iife27_()
        
        nw43_[int(3)] = (d_260_v1_) | ((d_260_v1_).set((self).f20, d_259_v0_))
        nw43_[int(4)] = d_260_v1_
        nw43_[int(5)] = d_260_v1_
        nw43_[int(6)] = d_260_v1_
        nw43_[int(7)] = (d_260_v1_).set((self).f20, True)
        nw43_[int(8)] = d_260_v1_
        nw43_[int(9)] = _dafny.Map({69: d_259_v0_})
        nw43_[int(10)] = _dafny.Map({(self).f20: d_259_v0_})
        nw43_[int(11)] = d_260_v1_
        nw43_[int(12)] = d_260_v1_
        nw43_[int(13)] = (d_265_v7_)
        nw43_[int(14)] = (d_266_v9_)[default__.safeIndex((self).f20, len(d_266_v9_))]
        nw43_[int(15)] = d_260_v1_
        nw43_[int(16)] = d_260_v1_
        nw43_[int(17)] = d_260_v1_
        nw43_[int(18)] = (d_260_v1_) | (d_260_v1_)
        nw43_[int(19)] = d_260_v1_
        d_268_v10_ = nw43_
        d_270_v11_: _dafny.Seq
        d_270_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjg"))
        rhs32_ = d_268_v10_
        rhs33_ = (self).f20
        rhs34_ = not((self).fm0(d_270_v11_, (self).f20, (self).f20, (self).f20, globalState))
        lhs31_ = globalState
        lhs32_ = globalState
        d_268_v10_ = rhs32_
        lhs31_.f3 = rhs33_
        lhs32_.f5 = rhs34_
        d_271_v12_: _dafny.Array
        def lambda11_(d_272_v0_):
            def lambda12_(d_273_i0_):
                return _dafny.Map({d_272_v0_: d_272_v0_})

            return lambda12_

        init6_ = lambda11_(d_259_v0_)
        nw44_ = _dafny.Array(None, 29)
        for i0_6_ in range(nw44_.length(0)):
            nw44_[i0_6_] = init6_(i0_6_)
        d_271_v12_ = nw44_
        d_274_v13_: _dafny.Seq
        d_274_v13_ = _dafny.SeqWithoutIsStrInference([d_271_v12_])
        d_275_v14_: _dafny.Map
        d_275_v14_ = _dafny.Map({d_259_v0_: 468})
        d_276_v15_: _dafny.Map
        d_276_v15_ = _dafny.Map({(d_274_v13_)[default__.safeIndex(len(_dafny.Map({d_259_v0_: d_259_v0_})), len(d_274_v13_))]: (d_275_v14_) | (d_275_v14_)})
        d_276_v15_ = (d_276_v15_).set(d_271_v12_, d_275_v14_)
        d_277_i1_: int
        d_277_i1_ = 0
        with _dafny.label("2"):
            while d_259_v0_:
                with _dafny.c_label("2"):
                    if (d_277_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_277_i1_ = (d_277_i1_) + (1)
                    (globalState).f3 = (0) - ((self).f20)
                    (globalState).f7 = d_259_v0_
                    d_278_v16_: str
                    d_278_v16_ = _dafny.CodePoint('w')
                    (globalState).f5 = not ((d_278_v16_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axeqcyxa")))) or (d_259_v0_)
                    d_279_v17_: _dafny.MultiSet
                    d_279_v17_ = _dafny.MultiSet([d_259_v0_, True])
                    (globalState).f3 = ((0) - (len(d_270_v11_)) if d_259_v0_ else ((d_279_v17_)[True] if (True) in (d_279_v17_) else (self).f20))
                    pass
            pass
        d_280_v18_: _dafny.MultiSet
        d_280_v18_ = _dafny.MultiSet([(self).f20])
        d_281_v19_: _dafny.Set
        d_282_v20_: bool
        d_283_v21_: _dafny.Map
        out12_: _dafny.Set
        out13_: bool
        out14_: _dafny.Map
        out12_, out13_, out14_ = default__.m0(d_280_v18_, globalState)
        d_281_v19_ = out12_
        d_282_v20_ = out13_
        d_283_v21_ = out14_
        if d_282_v20_:
            d_284_v22_: _dafny.Map
            d_284_v22_ = _dafny.Map({d_282_v20_: d_259_v0_})
            index41_ = default__.safeIndex(192, (d_271_v12_).length(0))
            (d_271_v12_)[index41_] = d_284_v22_
            index42_ = default__.safeIndex(192, (d_271_v12_).length(0))
            (d_271_v12_)[index42_] = ((d_284_v22_).set(d_282_v20_, d_282_v20_)) | (d_284_v22_)
            d_265_v7_ = d_263_v6_
            d_285_v23_: str
            d_285_v23_ = _dafny.CodePoint('s')
            d_286_v24_: _dafny.Map
            d_286_v24_ = _dafny.Map({d_259_v0_: d_285_v23_})
            (globalState).f3 = (self).fm1(_dafny.Map({d_286_v24_: default__.fm2(d_259_v0_, len(d_275_v14_), globalState)}), globalState)
            source5_ = d_263_v6_
            d_287___mcc_h0_ = source5_
            d_288_cf0_ = d_287___mcc_h0_
            d_270_v11_ = d_270_v11_
            d_289_v25_: C0
            nw45_ = C0()
            nw45_.ctor__((self).f20)
            d_289_v25_ = nw45_
            d_290_v26_: _dafny.Array
            nw46_ = _dafny.Array(False, 25)
            d_290_v26_ = nw46_
            index43_ = default__.safeIndex(144, (d_290_v26_).length(0))
            (d_290_v26_)[index43_] = d_282_v20_
            index44_ = default__.safeIndex(144, (d_290_v26_).length(0))
            (d_290_v26_)[index44_] = not((((d_289_v25_).f22) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_291_i2_ in range(default__.abs(311))])))) > (-390))
            d_292_v27_: _dafny.Seq
            d_292_v27_ = _dafny.SeqWithoutIsStrInference([d_289_v25_, d_289_v25_, d_289_v25_])
            rhs35_ = (d_292_v27_)[default__.safeIndex(((d_289_v25_).f22) - ((d_289_v25_).f22), len(d_292_v27_))]
            rhs36_ = not((self).fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdmtla")), ((d_289_v25_).f22) + ((0) - ((d_289_v25_).f22)), (self).f20, -502, globalState))
            rhs37_ = True
            rhs38_ = d_259_v0_
            lhs33_ = globalState
            lhs34_ = globalState
            d_289_v25_ = rhs35_
            lhs33_.f7 = rhs36_
            lhs34_.f5 = rhs37_
            d_282_v20_ = rhs38_
            def iife28_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(99, -193):
                    d_293_v28_: int = compr_16_
                    if ((99) <= (d_293_v28_)) and ((d_293_v28_) < (-193)):
                        coll16_[(d_293_v28_) + ((self).f20)] = d_259_v0_
                return _dafny.Map(coll16_)
            d_283_v21_ = (iife28_()
            ) | ((d_266_v9_)[default__.safeIndex((self).f20, len(d_266_v9_))])
        elif True:
            d_282_v20_ = d_259_v0_
            if not (d_282_v20_) or ((True if d_282_v20_ else d_282_v20_)):
                d_294_v29_: _dafny.Array
                nw47_ = _dafny.Array(int(0), 25)
                d_294_v29_ = nw47_
                nw48_ = _dafny.Array(int(0), 1)
                d_294_v29_ = nw48_
                (globalState).f11 = not (d_259_v0_) or (((self).f21)[default__.safeIndex(len(default__.fm5((self).f20, d_282_v20_, (self).f20, globalState)), len((self).f21))])
                d_294_v29_ = d_294_v29_
                (globalState).f3 = (self).f20
                d_295_v30_: str
                d_295_v30_ = _dafny.CodePoint('f')
                d_296_v31_: _dafny.Map
                d_296_v31_ = _dafny.Map({d_295_v30_: d_282_v20_})
                (globalState).f11 = ((d_296_v31_)[d_295_v30_] if (d_295_v30_) in (d_296_v31_) else (d_280_v18_).isdisjoint(d_280_v18_))
            elif True:
                (globalState).f3 = (self).f20
                d_297_v32_: C0
                nw49_ = C0()
                nw49_.ctor__((self).f20)
                d_297_v32_ = nw49_
                (globalState).f14 = ((d_281_v19_).intersection(d_281_v19_)).issubset(d_281_v19_)
                d_298_v33_: str
                d_298_v33_ = _dafny.CodePoint('s')
                d_298_v33_ = d_298_v33_
                d_270_v11_ = ((d_270_v11_).set(default__.safeIndex((self).f20, len(d_270_v11_)), d_298_v33_)) + (d_270_v11_)
            (globalState).f5 = d_282_v20_
            d_299_v34_: _dafny.Seq
            d_299_v34_ = _dafny.SeqWithoutIsStrInference([d_263_v6_])
            rhs39_ = d_282_v20_
            rhs40_ = ((d_265_v7_) in (d_299_v34_)) and (d_282_v20_)
            lhs35_ = globalState
            d_259_v0_ = rhs39_
            lhs35_.f7 = rhs40_
            (globalState).f3 = default__.safeModuloInt((self).f20, (self).f20)
        source6_ = d_260_v1_
        d_300___mcc_h1_ = source6_
        d_301_cf0_ = d_300___mcc_h1_
        d_302_v35_: C0
        nw50_ = C0()
        nw50_.ctor__((0) - ((self).f20))
        d_302_v35_ = nw50_
        d_303_v36_: _dafny.Array
        def lambda13_(d_304_i3_):
            return default__.safeDivisionInt(d_304_i3_, -681)

        init7_ = lambda13_
        nw51_ = _dafny.Array(None, 17)
        for i0_7_ in range(nw51_.length(0)):
            nw51_[i0_7_] = init7_(i0_7_)
        d_303_v36_ = nw51_
        d_305_v37_: _dafny.Seq
        d_305_v37_ = _dafny.SeqWithoutIsStrInference([d_303_v36_])
        d_306_v38_: _dafny.Map
        d_306_v38_ = _dafny.Map({d_280_v18_: (d_305_v37_)[default__.safeIndex((self).f20, len(d_305_v37_))]})
        d_307_v39_: _dafny.Array
        nw52_ = _dafny.Array(None, 8)
        nw52_[int(0)] = d_303_v36_
        nw52_[int(1)] = ((d_306_v38_)[_dafny.MultiSet(d_261_v3_)] if (_dafny.MultiSet(d_261_v3_)) in (d_306_v38_) else d_303_v36_)
        nw52_[int(2)] = d_303_v36_
        nw52_[int(3)] = d_303_v36_
        nw52_[int(4)] = d_303_v36_
        nw52_[int(5)] = d_303_v36_
        nw52_[int(6)] = d_303_v36_
        nw52_[int(7)] = (d_305_v37_)[default__.safeIndex((d_302_v35_).f22, len(d_305_v37_))]
        d_307_v39_ = nw52_
        index45_ = default__.safeIndex(926, (d_307_v39_).length(0))
        (d_307_v39_)[index45_] = d_303_v36_
        index46_ = default__.safeIndex(926, (d_307_v39_).length(0))
        rhs41_ = d_303_v36_
        rhs42_ = d_265_v7_
        rhs43_ = (0) - (len(d_261_v3_))
        lhs36_ = d_307_v39_
        lhs37_ = default__.safeIndex(926, (d_307_v39_).length(0))
        lhs38_ = globalState
        lhs36_[lhs37_] = rhs41_
        d_263_v6_ = rhs42_
        lhs38_.f3 = rhs43_
        d_308_v40_: _dafny.Array
        nw53_ = _dafny.Array(_dafny.Map({}), 14)
        d_308_v40_ = nw53_
        index47_ = default__.safeIndex(673, (d_308_v40_).length(0))
        (d_308_v40_)[index47_] = d_263_v6_
        index48_ = default__.safeIndex(673, (d_308_v40_).length(0))
        (d_308_v40_)[index48_] = d_263_v6_
        arr0_ = (d_307_v39_)[default__.safeIndex(926, (d_307_v39_).length(0))]
        index49_ = default__.safeIndex(929, ((d_307_v39_)[default__.safeIndex(926, (d_307_v39_).length(0))]).length(0))
        arr0_[index49_] = (self).f20
        d_309_v41_: _dafny.Set
        d_309_v41_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_259_v0_]))})
        d_310_v42_: _dafny.Map
        d_310_v42_ = _dafny.Map({d_263_v6_: d_309_v41_})
        arr1_ = (d_307_v39_)[default__.safeIndex(926, (d_307_v39_).length(0))]
        index50_ = default__.safeIndex(929, ((d_307_v39_)[default__.safeIndex(926, (d_307_v39_).length(0))]).length(0))
        rhs44_ = (self).f20
        rhs45_ = (0) - ((0) - (default__.safeDivisionInt(((self).f20) * ((self).f20), (d_302_v35_).f22)))
        rhs46_ = d_310_v42_
        rhs47_ = d_309_v41_
        lhs39_ = globalState
        lhs40_ = (d_307_v39_)[default__.safeIndex(926, (d_307_v39_).length(0))]
        lhs41_ = default__.safeIndex(929, ((d_307_v39_)[default__.safeIndex(926, (d_307_v39_).length(0))]).length(0))
        lhs39_.f3 = rhs44_
        lhs40_[lhs41_] = rhs45_
        d_310_v42_ = rhs46_
        d_309_v41_ = rhs47_

    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
