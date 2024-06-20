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
        return ((_dafny.MultiSet([49])) | (_dafny.MultiSet([-112]))) in (_dafny.Map({_dafny.MultiSet([-229]): not(True)}))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return (default__.safeDivisionInt(len(_dafny.Set({True})), (0) - (-576))) - (len((_dafny.Map({604: True})) | (_dafny.Map({471: True}))))

    @staticmethod
    def fm2(p0, p1, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byj"))): -962})), -999])) + (_dafny.SeqWithoutIsStrInference([227]))) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: 599}))])))

    @staticmethod
    def fm3(p0, p1, globalState):
        return _dafny.CodePoint('w')

    @staticmethod
    def fm9(p0, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: True}))) | (_dafny.Map({not(False): True}))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyepokid"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gl"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djdolp"))))

    @staticmethod
    def fm11(p0, globalState):
        source0_ = (D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([822, 444]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True]))])])) if True else D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([634])])))
        if source0_.is_DC1:
            d_0___mcc_h0_ = source0_.cf1
            d_1_cf1_ = d_0___mcc_h0_
            if False:
                return D3_DC11(_dafny.CodePoint('q'), not(True), not(True))
            elif True:
                return D3_DC11(_dafny.CodePoint('m'), not(True), True)
        elif source0_.is_DC2:
            d_2___mcc_h1_ = source0_.cf2
            d_3___mcc_h2_ = source0_.cf3
            d_4___mcc_h3_ = source0_.cf4
            d_5___mcc_h4_ = source0_.cf5
            d_6___mcc_h5_ = source0_.cf6
            d_7_cf6_ = d_6___mcc_h5_
            d_8_cf5_ = d_5___mcc_h4_
            d_9_cf4_ = d_4___mcc_h3_
            d_10_cf3_ = d_3___mcc_h2_
            d_11_cf2_ = d_2___mcc_h1_
            return D3_DC11(_dafny.CodePoint('a'), d_9_cf4_, d_11_cf2_)
        elif source0_.is_DC3:
            d_12___mcc_h6_ = source0_.cf7
            d_13_cf7_ = d_12___mcc_h6_
            return D3_DC11(_dafny.CodePoint('f'), False, True)
        elif True:
            d_14___mcc_h7_ = source0_.cf0
            d_15_cf0_ = d_14___mcc_h7_
            return D3_DC11(_dafny.CodePoint('t'), False, True)

    @staticmethod
    def fm12(p0, p1, globalState):
        def iife0_():
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(-198, -371):
                    d_16_v1_: int = compr_2_
                    if ((-198) <= (d_16_v1_)) and ((d_16_v1_) < (-371)):
                        coll2_[(d_16_v1_) + ((_dafny.MultiSet([-110])).cardinality)] = not(False)
                return _dafny.Map(coll2_)
            coll0_ = _dafny.Map()
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(-198, -371):
                    d_16_v1_: int = compr_1_
                    if ((-198) <= (d_16_v1_)) and ((d_16_v1_) < (-371)):
                        coll1_[(d_16_v1_) + ((_dafny.MultiSet([-110])).cardinality)] = not(False)
                return _dafny.Map(coll1_)
            compr_0_: int
            for compr_0_ in (iife1_()
            ).keys.Elements:
                d_17_v0_: int = compr_0_
                if (d_17_v0_) in (iife2_()
                ):
                    coll0_[(d_17_v0_) - (873)] = True
            return _dafny.Map(coll0_)
        return ((_dafny.SeqWithoutIsStrInference([len(iife0_()
) for d_18_i0_ in range(default__.abs(312))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([987])) for d_19_i1_ in range(default__.abs(-962))]))) + (_dafny.SeqWithoutIsStrInference([(0) - (-915)]))

    @staticmethod
    def fm13(globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(302, 656):
                d_20_v0_: int = compr_3_
                if ((302) <= (d_20_v0_)) and ((d_20_v0_) < (656)):
                    coll3_[(d_20_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkdjgk"))))] = (_dafny.MultiSet([(D2_DC6(False)).cf10])).cardinality
            return _dafny.Map(coll3_)
        return (_dafny.SeqWithoutIsStrInference([iife3_()
        ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.Map({_dafny.Set({False}): not(False)}))})]))

    @staticmethod
    def fm14(globalState):
        return D0_DC0((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([587]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([921 for d_21_i0_ in range(default__.abs(-336))]))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([531, 263]), _dafny.MultiSet([245, 950])])))

    @staticmethod
    def fm15(globalState):
        return _dafny.Set({(0) - ((0) - ((0) - (-122)))})

    @staticmethod
    def fm16(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_22_i0_ in range(default__.abs(178))]))])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet((D7_DC19(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([151])): _dafny.CodePoint('v')})): True}))]))).cf34)).cardinality]))), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqbthl")))])) | ((D4_DC14(_dafny.MultiSet([119, -456]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ty"))), D2_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehtttxx"))))).cf28), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lotvixcay")))]), _dafny.MultiSet([508, -985])])

    @staticmethod
    def fm17(globalState):
        source1_ = D7_DC21(True, True)
        if source1_.is_DC20:
            d_23___mcc_h0_ = source1_.cf35
            d_24___mcc_h1_ = source1_.cf36
            d_25___mcc_h2_ = source1_.cf37
            d_26___mcc_h3_ = source1_.cf38
            d_27___mcc_h4_ = source1_.cf39
            d_28_cf39_ = d_27___mcc_h4_
            d_29_cf38_ = d_26___mcc_h3_
            d_30_cf37_ = d_25___mcc_h2_
            d_31_cf36_ = d_24___mcc_h1_
            d_32_cf35_ = d_23___mcc_h0_
            return D2_DC6(d_32_cf35_)
        elif source1_.is_DC21:
            d_33___mcc_h5_ = source1_.cf40
            d_34___mcc_h6_ = source1_.cf41
            d_35_cf41_ = d_34___mcc_h6_
            d_36_cf40_ = d_33___mcc_h5_
            if d_35_cf41_:
                return D2_DC6(False)
            elif True:
                return D2_DC6(d_36_cf40_)
        elif True:
            d_37___mcc_h7_ = source1_.cf34
            d_38_cf34_ = d_37___mcc_h7_
            return D2_DC6(False)

    @staticmethod
    def fm18(p0, p1, globalState):
        return ((_dafny.Map({_dafny.CodePoint('x'): len(_dafny.Map({897: False}))})) | (_dafny.Map({_dafny.CodePoint('d'): -664}))) | (_dafny.Map({_dafny.CodePoint('a'): 441}))

    @staticmethod
    def fm19(globalState):
        return _dafny.MultiSet([D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([500 for d_39_i0_ in range(default__.abs(426))]))])), D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([410, 803]))])), (D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True]))])])) if False else D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.Map({False: 723}))])])))])

    @staticmethod
    def fm20(p0, globalState):
        return ((_dafny.Map({len(_dafny.Map({not(True): _dafny.MultiSet([False])})): -497}) if True else _dafny.Map({130: -588}))) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_40_i0_ in range(default__.abs(811))])): (0) - ((_dafny.MultiSet([len(_dafny.Map({True: _dafny.Map({False: True})})), len(_dafny.SeqWithoutIsStrInference([True]))])).cardinality)})))

    @staticmethod
    def fm21(p0, p1, globalState):
        return _dafny.Set({False, True, False})

    @staticmethod
    def m1(p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_41_v0_: _dafny.Map
        d_41_v0_ = _dafny.Map({((0) - (p1)) > (p1): p0})
        d_41_v0_ = (d_41_v0_).set(p0, p0)
        d_42_v1_: _dafny.Seq
        d_42_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0])])
        d_43_v2_: D0
        d_43_v2_ = D0_DC3((d_42_v1_)[default__.safeIndex(472, len(d_42_v1_))])
        rhs0_ = (d_43_v2_ if p0 else (d_43_v2_ if p0 else d_43_v2_))
        rhs1_ = (p1) * (p1)
        lhs0_ = globalState
        d_43_v2_ = rhs0_
        lhs0_.f2 = rhs1_
        d_44_v3_: _dafny.Array
        nw0_ = _dafny.Array(int(0), 22)
        d_44_v3_ = nw0_
        d_45_v4_: _dafny.Set
        d_45_v4_ = _dafny.Set({d_44_v3_})
        d_46_v5_: C0
        nw1_ = C0()
        nw1_.ctor__()
        d_46_v5_ = nw1_
        d_47_v6_: _dafny.Seq
        d_47_v6_ = _dafny.SeqWithoutIsStrInference([d_46_v5_])
        d_48_v7_: str
        d_48_v7_ = _dafny.CodePoint('o')
        d_49_v8_: _dafny.MultiSet
        d_49_v8_ = _dafny.MultiSet([d_48_v7_, d_48_v7_, d_48_v7_])
        d_50_v9_: _dafny.Map
        d_50_v9_ = _dafny.Map({d_49_v8_: p1})
        d_51_v10_: D0
        d_51_v10_ = D0_DC2(p0, d_45_v4_, not(not((d_47_v6_) < (d_47_v6_))), p0, len(d_50_v9_))
        source2_ = d_51_v10_
        if source2_.is_DC1:
            d_52___mcc_h0_ = source2_.cf1
            d_53_cf1_ = d_52___mcc_h0_
            d_54_v11_: _dafny.Array
            nw2_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_54_v11_ = nw2_
            d_55_v12_: D6
            d_55_v12_ = D6_DC16(d_44_v3_)
            index0_ = default__.safeIndex(549, (d_54_v11_).length(0))
            (d_54_v11_)[index0_] = (d_55_v12_).cf32
            d_56_v13_: _dafny.Seq
            d_56_v13_ = _dafny.SeqWithoutIsStrInference([d_44_v3_])
            index1_ = default__.safeIndex(549, (d_54_v11_).length(0))
            (d_54_v11_)[index1_] = (d_44_v3_ if p0 else (d_56_v13_)[default__.safeIndex(686, len(d_56_v13_))])
            d_57_v14_: _dafny.Array
            def lambda0_(d_58_p0_):
                def lambda1_(d_59_i0_):
                    return not(d_58_p0_)

                return lambda1_

            init0_ = lambda0_(p0)
            nw3_ = _dafny.Array(None, 25)
            for i0_0_ in range(nw3_.length(0)):
                nw3_[i0_0_] = init0_(i0_0_)
            d_57_v14_ = nw3_
            index2_ = default__.safeIndex(398, (d_57_v14_).length(0))
            (d_57_v14_)[index2_] = p0
            index3_ = default__.safeIndex(398, (d_57_v14_).length(0))
            (d_57_v14_)[index3_] = p0
            if (d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]:
                d_60_v15_: C0
                nw4_ = C0()
                nw4_.ctor__()
                d_60_v15_ = nw4_
                (globalState).f8 = d_53_cf1_
                d_61_v16_: _dafny.Set
                d_61_v16_ = _dafny.Set({(d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]})
                d_62_v17_: C1
                nw5_ = C1()
                nw5_.ctor__(d_61_v16_, d_53_cf1_)
                d_62_v17_ = nw5_
                d_63_v18_: _dafny.Array
                nw6_ = _dafny.Array(None, 15)
                nw6_[int(0)] = d_62_v17_
                nw6_[int(1)] = d_62_v17_
                nw6_[int(2)] = d_62_v17_
                nw6_[int(3)] = d_62_v17_
                nw6_[int(4)] = d_62_v17_
                nw6_[int(5)] = d_62_v17_
                nw6_[int(6)] = d_62_v17_
                nw6_[int(7)] = d_62_v17_
                nw6_[int(8)] = d_62_v17_
                nw6_[int(9)] = d_62_v17_
                nw6_[int(10)] = d_62_v17_
                nw6_[int(11)] = d_62_v17_
                nw6_[int(12)] = d_62_v17_
                nw6_[int(13)] = d_62_v17_
                nw6_[int(14)] = d_62_v17_
                d_63_v18_ = nw6_
                d_64_v19_: _dafny.Array
                nw7_ = _dafny.Array(None, 8)
                nw7_[int(0)] = d_63_v18_
                nw7_[int(1)] = d_63_v18_
                nw7_[int(2)] = d_63_v18_
                nw7_[int(3)] = d_63_v18_
                nw7_[int(4)] = d_63_v18_
                nw7_[int(5)] = d_63_v18_
                nw7_[int(6)] = d_63_v18_
                nw7_[int(7)] = d_63_v18_
                d_64_v19_ = nw7_
                index4_ = default__.safeIndex(835, (d_64_v19_).length(0))
                (d_64_v19_)[index4_] = d_63_v18_
                d_65_v20_: _dafny.Seq
                d_65_v20_ = _dafny.SeqWithoutIsStrInference([p0])
                d_66_v21_: _dafny.Set
                d_66_v21_ = _dafny.Set({d_65_v20_})
                d_67_v22_: _dafny.Map
                d_67_v22_ = _dafny.Map({p0: d_63_v18_})
                d_68_v23_: _dafny.Seq
                d_68_v23_ = _dafny.SeqWithoutIsStrInference([d_63_v18_, ((d_67_v22_)[(d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]] if ((d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]) in (d_67_v22_) else d_63_v18_)])
                d_69_v24_: _dafny.MultiSet
                d_69_v24_ = _dafny.MultiSet([d_62_v17_.f27])
                d_70_v25_: _dafny.Seq
                d_70_v25_ = _dafny.SeqWithoutIsStrInference([(d_68_v23_)[default__.safeIndex((D3_DC10(p0, (d_69_v24_).cardinality, p0)).cf17, len(d_68_v23_))], d_63_v18_, d_63_v18_])
                d_71_v26_: T0
                nw8_ = C0()
                nw8_.ctor__()
                d_71_v26_ = nw8_
                d_72_v27_: _dafny.Seq
                d_72_v27_ = _dafny.SeqWithoutIsStrInference([d_71_v26_])
                d_73_v28_: _dafny.MultiSet
                d_73_v28_ = _dafny.MultiSet([(d_72_v27_)[default__.safeIndex(p1, len(d_72_v27_))], d_71_v26_])
                index5_ = default__.safeIndex(398, (d_57_v14_).length(0))
                index6_ = default__.safeIndex(835, (d_64_v19_).length(0))
                rhs2_ = (d_66_v21_).ispropersubset(_dafny.Set({d_65_v20_}))
                rhs3_ = (d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]
                rhs4_ = (d_70_v25_)[default__.safeIndex((d_73_v28_).cardinality, len(d_70_v25_))]
                lhs1_ = globalState
                lhs2_ = d_57_v14_
                lhs3_ = default__.safeIndex(398, (d_57_v14_).length(0))
                lhs4_ = d_64_v19_
                lhs5_ = default__.safeIndex(835, (d_64_v19_).length(0))
                lhs1_.f23 = rhs2_
                lhs2_[lhs3_] = rhs3_
                lhs4_[lhs5_] = rhs4_
                (globalState).f2 = p1
                (globalState).f23 = not((d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))])
            elif True:
                d_74_v29_: _dafny.Seq
                d_74_v29_ = _dafny.SeqWithoutIsStrInference([(d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))], (d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]])
                d_75_v30_: _dafny.Seq
                d_75_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                (globalState).f19 = default__.fm1((d_74_v29_)[default__.safeIndex(len(d_75_v30_), len(d_74_v29_))], d_48_v7_, len(d_75_v30_), globalState)
                d_76_v31_: C0
                nw9_ = C0()
                nw9_.ctor__()
                d_76_v31_ = nw9_
                (globalState).f23 = (474) >= ((p1) * (572))
                d_77_v32_: _dafny.MultiSet
                d_77_v32_ = _dafny.MultiSet([p1])
                d_78_v33_: _dafny.Seq
                d_78_v33_ = _dafny.SeqWithoutIsStrInference([d_77_v32_, d_77_v32_])
                d_79_v34_: D0
                d_79_v34_ = D0_DC0(d_78_v33_)
                d_80_v35_: _dafny.MultiSet
                d_80_v35_ = _dafny.MultiSet([d_79_v34_, d_79_v34_])
                d_81_v36_: _dafny.MultiSet
                d_81_v36_ = d_80_v35_
                d_82_v37_: _dafny.Set
                d_82_v37_ = _dafny.Set({(d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]})
                d_83_v38_: T0
                nw10_ = C1()
                nw10_.ctor__(d_82_v37_, p1)
                d_83_v38_ = nw10_
                d_84_v39_: _dafny.Map
                d_84_v39_ = _dafny.Map({(d_54_v11_)[default__.safeIndex(549, (d_54_v11_).length(0))]: len(d_74_v29_)})
                d_85_v40_: _dafny.Map
                d_85_v40_ = _dafny.Map({d_83_v38_: ((d_84_v39_)[(d_54_v11_)[default__.safeIndex(549, (d_54_v11_).length(0))]] if ((d_54_v11_)[default__.safeIndex(549, (d_54_v11_).length(0))]) in (d_84_v39_) else d_53_cf1_)})
                d_86_v41_: _dafny.Seq
                d_86_v41_ = _dafny.SeqWithoutIsStrInference([d_85_v40_])
                index7_ = default__.safeIndex(398, (d_57_v14_).length(0))
                rhs5_ = (d_80_v35_).intersection(d_80_v35_)
                rhs6_ = (d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]
                rhs7_ = len(d_86_v41_)
                rhs8_ = d_48_v7_
                lhs6_ = d_57_v14_
                lhs7_ = default__.safeIndex(398, (d_57_v14_).length(0))
                lhs8_ = globalState
                d_81_v36_ = rhs5_
                lhs6_[lhs7_] = rhs6_
                lhs8_.f10 = rhs7_
                d_48_v7_ = rhs8_
                d_87_v42_: _dafny.Map
                d_87_v42_ = _dafny.Map({p1: ((d_41_v0_)[p0] if (p0) in (d_41_v0_) else p0)})
                d_88_v43_: _dafny.Seq
                d_88_v43_ = _dafny.SeqWithoutIsStrInference([default__.fm20(((d_87_v42_)[p1] if (p1) in (d_87_v42_) else (d_57_v14_)[default__.safeIndex(398, (d_57_v14_).length(0))]), globalState)])
                d_89_v44_: _dafny.Map
                d_89_v44_ = _dafny.Map({d_44_v3_: p0})
                d_90_v45_: _dafny.Map
                d_90_v45_ = _dafny.Map({len((d_88_v43_)[default__.safeIndex(len(d_89_v44_), len(d_88_v43_))]): (d_54_v11_)[default__.safeIndex(549, (d_54_v11_).length(0))]})
                (globalState).f11 = ((d_90_v45_)[default__.safeDivisionInt((d_76_v31_).fm8(d_82_v37_, globalState), p1)] if (default__.safeDivisionInt((d_76_v31_).fm8(d_82_v37_, globalState), p1)) in (d_90_v45_) else d_44_v3_)
            d_91_v46_: _dafny.MultiSet
            d_91_v46_ = _dafny.MultiSet([(d_54_v11_)[default__.safeIndex(549, (d_54_v11_).length(0))], d_44_v3_])
            index8_ = default__.safeIndex(398, (d_57_v14_).length(0))
            (d_57_v14_)[index8_] = (d_91_v46_).issubset(d_91_v46_)
        elif source2_.is_DC2:
            d_92___mcc_h1_ = source2_.cf2
            d_93___mcc_h2_ = source2_.cf3
            d_94___mcc_h3_ = source2_.cf4
            d_95___mcc_h4_ = source2_.cf5
            d_96___mcc_h5_ = source2_.cf6
            d_97_cf6_ = d_96___mcc_h5_
            d_98_cf5_ = d_95___mcc_h4_
            d_99_cf4_ = d_94___mcc_h3_
            d_100_cf3_ = d_93___mcc_h2_
            d_101_cf2_ = d_92___mcc_h1_
            r1 = d_99_cf4_
            d_102_v47_: T0
            nw11_ = C1()
            nw11_.ctor__(default__.fm21(d_98_cf5_, True, globalState), d_97_cf6_)
            d_102_v47_ = nw11_
            d_103_v48_: _dafny.Map
            d_103_v48_ = _dafny.Map({d_48_v7_: d_98_cf5_})
            d_104_v49_: _dafny.Seq
            d_104_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpfwbpp"))
            d_105_v50_: _dafny.Map
            d_105_v50_ = _dafny.Map({d_103_v48_: d_104_v49_})
            d_102_v47_ = (d_102_v47_ if (d_48_v7_) not in (((d_105_v50_)[d_103_v48_] if (d_103_v48_) in (d_105_v50_) else d_104_v49_)) else d_102_v47_)
            d_49_v8_ = _dafny.MultiSet(d_104_v49_)
            if not(d_101_cf2_):
                (globalState).f19 = p1
                d_106_v51_: _dafny.Set
                d_106_v51_ = _dafny.Set({d_99_cf4_, not(d_98_cf5_), d_99_cf4_, p0})
                d_107_v52_: _dafny.Seq
                d_107_v52_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_101_cf2_}), d_106_v51_, d_106_v51_])
                d_108_v53_: _dafny.Map
                d_108_v53_ = _dafny.Map({p0: len(d_107_v52_)})
                d_109_v54_: _dafny.Map
                d_109_v54_ = _dafny.Map({d_97_cf6_: len(d_108_v53_)})
                d_110_v55_: _dafny.Seq
                d_110_v55_ = _dafny.SeqWithoutIsStrInference([206, ((d_109_v54_)[p1] if (p1) in (d_109_v54_) else p1)])
                d_110_v55_ = d_110_v55_
                d_48_v7_ = d_48_v7_
                index9_ = default__.safeIndex(668, (d_44_v3_).length(0))
                (d_44_v3_)[index9_] = d_97_cf6_
                d_111_v56_: _dafny.MultiSet
                d_111_v56_ = _dafny.MultiSet([d_101_cf2_])
                index10_ = default__.safeIndex(668, (d_44_v3_).length(0))
                rhs9_ = d_97_cf6_
                rhs10_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjpyri")))) * (len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_112_i1_ in range(default__.abs(816))])) + (d_104_v49_)).set(default__.safeIndex((0) - (((d_111_v56_)[d_98_cf5_] if (d_98_cf5_) in (d_111_v56_) else d_97_cf6_)), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_113_i1_ in range(default__.abs(816))])) + (d_104_v49_))), d_48_v7_)))
                rhs11_ = True
                rhs12_ = (0) - ((d_97_cf6_) + ((0) - (d_97_cf6_)))
                lhs9_ = d_44_v3_
                lhs10_ = default__.safeIndex(668, (d_44_v3_).length(0))
                lhs11_ = globalState
                lhs12_ = globalState
                lhs9_[lhs10_] = rhs9_
                lhs11_.f5 = rhs10_
                r1 = rhs11_
                lhs12_.f10 = rhs12_
                d_114_v57_: C0
                nw12_ = C0()
                nw12_.ctor__()
                d_114_v57_ = nw12_
            elif True:
                index11_ = default__.safeIndex(250, (d_44_v3_).length(0))
                (d_44_v3_)[index11_] = p1
                index12_ = default__.safeIndex(250, (d_44_v3_).length(0))
                (d_44_v3_)[index12_] = p1
                d_115_v58_: _dafny.MultiSet
                d_115_v58_ = _dafny.MultiSet([d_104_v49_])
                index13_ = default__.safeIndex(250, (d_44_v3_).length(0))
                (d_44_v3_)[index13_] = ((d_115_v58_)[(d_104_v49_).set(default__.safeIndex(d_97_cf6_, len(d_104_v49_)), d_48_v7_)] if ((d_104_v49_).set(default__.safeIndex(d_97_cf6_, len(d_104_v49_)), d_48_v7_)) in (d_115_v58_) else p1)
                d_116_v59_: _dafny.Array
                nw13_ = _dafny.Array(False, 21)
                d_116_v59_ = nw13_
                d_117_v60_: _dafny.Map
                d_117_v60_ = _dafny.Map({d_116_v59_: p1})
                d_118_v61_: _dafny.Map
                d_118_v61_ = _dafny.Map({d_99_cf4_: d_117_v60_})
                (globalState).f23 = ((d_41_v0_)[(p1) < (d_97_cf6_)] if ((p1) < (d_97_cf6_)) in (d_41_v0_) else (d_117_v60_) == (((d_118_v61_)[d_98_cf5_] if (d_98_cf5_) in (d_118_v61_) else d_117_v60_)))
                d_119_v62_: _dafny.Set
                d_119_v62_ = _dafny.Set({d_98_cf5_, d_101_cf2_, not(d_98_cf5_)})
                d_120_v63_: _dafny.MultiSet
                d_120_v63_ = _dafny.MultiSet([len(d_104_v49_)])
                d_121_v64_: _dafny.Seq
                d_121_v64_ = _dafny.SeqWithoutIsStrInference([d_120_v63_])
                d_122_v65_: D0
                d_122_v65_ = D0_DC0(d_121_v64_)
                d_123_v66_: _dafny.MultiSet
                d_123_v66_ = _dafny.MultiSet([d_122_v65_])
                d_124_v67_: _dafny.MultiSet
                d_124_v67_ = d_123_v66_
                d_125_v68_: _dafny.Map
                d_125_v68_ = _dafny.Map({((d_46_v5_).fm8(d_119_v62_, globalState)) * ((0) - ((d_44_v3_)[default__.safeIndex(250, (d_44_v3_).length(0))])): d_124_v67_})
                d_125_v68_ = d_125_v68_
                (globalState).f22 = d_99_cf4_
        elif source2_.is_DC3:
            d_126___mcc_h6_ = source2_.cf7
            d_127_cf7_ = d_126___mcc_h6_
            d_128_v69_: _dafny.MultiSet
            d_128_v69_ = _dafny.MultiSet([d_44_v3_])
            d_129_v70_: _dafny.Seq
            d_129_v70_ = _dafny.SeqWithoutIsStrInference([p1])
            d_130_v71_: _dafny.Seq
            d_130_v71_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt((d_128_v69_).cardinality, (d_129_v70_)[default__.safeIndex(p1, len(d_129_v70_))]), p1])
            d_129_v70_ = (_dafny.SeqWithoutIsStrInference([p1])) + (d_130_v71_)
            index14_ = default__.safeIndex(955, (d_44_v3_).length(0))
            (d_44_v3_)[index14_] = p1
            index15_ = default__.safeIndex(955, (d_44_v3_).length(0))
            (d_44_v3_)[index15_] = p1
            d_131_v72_: _dafny.MultiSet
            d_131_v72_ = _dafny.MultiSet([default__.fm0(p0, d_41_v0_, globalState)])
            d_132_v73_: D2
            d_132_v73_ = D2_DC8(p1, p0, ((d_131_v72_)[p0] if (p0) in (d_131_v72_) else (d_44_v3_)[default__.safeIndex(955, (d_44_v3_).length(0))]))
            d_132_v73_ = d_132_v73_
            d_46_v5_ = d_46_v5_
        elif True:
            d_133___mcc_h7_ = source2_.cf0
            d_134_cf0_ = d_133___mcc_h7_
            (globalState).f25 = p0
            rhs13_ = not(p0)
            rhs14_ = p1
            lhs13_ = globalState
            r1 = rhs13_
            lhs13_.f2 = rhs14_
            d_135_v74_: D0
            d_135_v74_ = D0_DC1(p1)
            d_136_v75_: T0
            nw14_ = C0()
            nw14_.ctor__()
            d_136_v75_ = nw14_
            d_137_v76_: D3
            d_137_v76_ = D3_DC12(p0, d_135_v74_, p0, p0, d_136_v75_)
            d_138_v77_: _dafny.Map
            d_138_v77_ = _dafny.Map({d_137_v76_: p1})
            d_138_v77_ = (d_138_v77_).set(d_137_v76_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))))
            d_139_v78_: _dafny.Seq
            d_139_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akpsj"))
            d_139_v78_ = (d_139_v78_) + (d_139_v78_)
        hi0_ = p1
        for d_140_i2_ in range(len((d_41_v0_).set(p0, p0)), hi0_):
            d_141_v79_: D3
            d_141_v79_ = D3_DC10(p0, d_140_i2_, p0)
            r1 = default__.fm0((d_141_v79_).cf16, d_41_v0_, globalState)
            if (106) != (d_140_i2_):
                d_142_v80_: _dafny.MultiSet
                d_142_v80_ = _dafny.MultiSet([p0, p0])
                d_143_v81_: _dafny.Seq
                d_143_v81_ = _dafny.SeqWithoutIsStrInference([p0])
                rhs15_ = not(not(default__.fm0(not(p0), _dafny.Map({False: default__.fm0(True, d_41_v0_, globalState)}), globalState)))
                rhs16_ = d_140_i2_
                rhs17_ = d_142_v80_
                rhs18_ = (d_143_v81_)[default__.safeIndex(p1, len(d_143_v81_))]
                lhs14_ = globalState
                lhs15_ = globalState
                lhs16_ = globalState
                lhs17_ = globalState
                lhs14_.f25 = rhs15_
                lhs15_.f19 = rhs16_
                lhs16_.f16 = rhs17_
                lhs17_.f23 = rhs18_
                d_144_v82_: _dafny.Set
                d_144_v82_ = _dafny.Set({p0, p0, p0, p0, p0})
                d_145_v84_: _dafny.Set
                d_145_v84_ = _dafny.Set({d_140_i2_, p1, d_140_i2_})
                d_146_v85_: _dafny.Map
                def iife4_():
                    coll4_ = _dafny.Set()
                    compr_4_: int
                    for compr_4_ in _dafny.IntegerRange(20, 792):
                        d_147_v83_: int = compr_4_
                        if ((20) <= (d_147_v83_)) and ((d_147_v83_) < (792)):
                            coll4_ = coll4_.union(_dafny.Set([(d_147_v83_) + (p1)]))
                    return _dafny.Set(coll4_)
                d_146_v85_ = _dafny.Map({d_144_v82_: (iife4_()
                ) | (d_145_v84_)})
                d_148_v86_: T0
                nw15_ = C1()
                nw15_.ctor__(_dafny.Set({p0}), d_140_i2_)
                d_148_v86_ = nw15_
                d_149_v87_: _dafny.MultiSet
                d_149_v87_ = _dafny.MultiSet([d_148_v86_, d_148_v86_])
                rhs19_ = (_dafny.MultiSet([d_148_v86_, d_148_v86_])).issubset((d_149_v87_).intersection(d_149_v87_))
                rhs20_ = _dafny.Map({d_144_v82_: (d_145_v84_) - (d_145_v84_)})
                rhs21_ = default__.fm0(default__.fm0(False, d_41_v0_, globalState), d_41_v0_, globalState)
                rhs22_ = p1
                lhs18_ = globalState
                lhs19_ = globalState
                lhs20_ = globalState
                lhs18_.f22 = rhs19_
                d_146_v85_ = rhs20_
                lhs19_.f22 = rhs21_
                lhs20_.f15 = rhs22_
                d_150_v88_: _dafny.Map
                d_150_v88_ = _dafny.Map({d_148_v86_: (d_148_v86_ if p0 else d_148_v86_)})
                d_150_v88_ = (d_150_v88_).set(d_148_v86_, d_148_v86_)
                d_151_v89_: D0
                d_151_v89_ = D0_DC1(p1)
                d_152_v90_: _dafny.Seq
                d_152_v90_ = _dafny.SeqWithoutIsStrInference([800, d_140_i2_])
                d_153_v91_: _dafny.Map
                d_153_v91_ = _dafny.Map({(p0) and (not(default__.fm0(True, _dafny.Map({p0: p0}), globalState))): len((default__.fm12(len(d_145_v84_), d_151_v89_, globalState)) + (d_152_v90_))})
                d_153_v91_ = (d_153_v91_).set((True if not(p0) else default__.fm0(p0, d_41_v0_, globalState)), default__.safeDivisionInt(len(default__.fm21(p0, p0, globalState)), p1))
                d_154_v92_: _dafny.Seq
                d_154_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldritmxh"))
                d_154_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oocr"))
            elif True:
                (globalState).f5 = (d_140_i2_ if p0 else d_140_i2_)
                d_155_v93_: _dafny.Map
                d_155_v93_ = _dafny.Map({p1: p0})
                d_156_v95_: _dafny.Seq
                d_156_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_157_v96_: _dafny.Map
                d_157_v96_ = _dafny.Map({len(d_156_v95_): p1})
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: int
                    for compr_5_ in (d_157_v96_).keys.Elements:
                        d_158_v94_: int = compr_5_
                        if (d_158_v94_) in (d_157_v96_):
                            coll5_[(d_158_v94_) - (p1)] = p0
                    return _dafny.Map(coll5_)
                rhs23_ = (iife5_()
                ) | (d_155_v93_)
                rhs24_ = d_48_v7_
                d_155_v93_ = rhs23_
                d_48_v7_ = rhs24_
                (globalState).f22 = p0
                index16_ = default__.safeIndex(45, (d_44_v3_).length(0))
                (d_44_v3_)[index16_] = (p1) * (695)
                index17_ = default__.safeIndex(45, (d_44_v3_).length(0))
                (d_44_v3_)[index17_] = d_140_i2_
                d_159_v97_: _dafny.Set
                d_159_v97_ = _dafny.Set({p0})
                d_160_v98_: C1
                nw16_ = C1()
                nw16_.ctor__(d_159_v97_, p1)
                d_160_v98_ = nw16_
                d_161_v99_: _dafny.Map
                d_161_v99_ = _dafny.Map({d_160_v98_: len(d_156_v95_)})
                d_162_v100_: _dafny.Map
                d_162_v100_ = _dafny.Map({d_160_v98_: False})
                d_163_v101_: _dafny.Seq
                d_163_v101_ = _dafny.SeqWithoutIsStrInference([d_140_i2_, len(d_162_v100_), (0) - (p1)])
                d_164_v102_: _dafny.Seq
                d_164_v102_ = _dafny.SeqWithoutIsStrInference([False, p0])
                d_165_v103_: D0
                d_165_v103_ = D0_DC1((_dafny.MultiSet(d_164_v102_)).cardinality)
                d_166_v104_: _dafny.Map
                d_166_v104_ = _dafny.Map({d_156_v95_: p0})
                pat_let_tv0_ = p0
                pat_let_tv1_ = d_48_v7_
                pat_let_tv2_ = d_44_v3_
                pat_let_tv3_ = d_44_v3_
                pat_let_tv4_ = globalState
                d_167_v105_: _dafny.Array
                nw17_ = _dafny.Array(None, 20)
                nw17_[int(0)] = (d_44_v3_)[default__.safeIndex(45, (d_44_v3_).length(0))]
                nw17_[int(1)] = d_140_i2_
                nw17_[int(2)] = (952 if True else d_140_i2_)
                nw17_[int(3)] = (d_44_v3_)[default__.safeIndex(45, (d_44_v3_).length(0))]
                nw17_[int(4)] = (len(d_161_v99_)) + ((0) - ((d_44_v3_)[default__.safeIndex(45, (d_44_v3_).length(0))]))
                nw17_[int(5)] = p1
                nw17_[int(6)] = (d_163_v101_)[default__.safeIndex(p1, len(d_163_v101_))]
                nw17_[int(7)] = (0) - (d_160_v98_.f27)
                nw17_[int(8)] = (0) - (d_160_v98_.f27)
                nw17_[int(9)] = (0) - (d_160_v98_.f27)
                nw17_[int(10)] = (d_44_v3_)[default__.safeIndex(45, (d_44_v3_).length(0))]
                nw17_[int(11)] = ((d_44_v3_)[default__.safeIndex(45, (d_44_v3_).length(0))] if not(default__.fm0(p0, _dafny.Map({p0: p0}), globalState)) else (_dafny.MultiSet([(d_164_v102_)[default__.safeIndex((0) - (d_140_i2_), len(d_164_v102_))], p0, p0, p0])).cardinality)
                nw17_[int(12)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrsul")))
                nw17_[int(13)] = (d_44_v3_)[default__.safeIndex(45, (d_44_v3_).length(0))]
                nw17_[int(14)] = len(d_156_v95_)
                nw17_[int(15)] = (d_160_v98_.f27 if p0 else default__.fm1(p0, d_48_v7_, d_160_v98_.f27, globalState))
                nw17_[int(16)] = p1
                def iife6_(_pat_let0_0):
                    def iife7_(d_168_dt__update__tmp_h2_):
                        def iife8_(_pat_let1_0):
                            def iife9_(d_169_dt__update_hcf1_h0_):
                                return D0_DC1(d_169_dt__update_hcf1_h0_)
                            return iife9_(_pat_let1_0)
                        return iife8_(default__.fm1(pat_let_tv0_, pat_let_tv1_, (pat_let_tv3_)[default__.safeIndex(45, (pat_let_tv2_).length(0))], pat_let_tv4_))
                    return iife7_(_pat_let0_0)
                nw17_[int(17)] = ((iife6_(d_165_v103_)).cf1) + (((_dafny.MultiSet([p0])).set(((d_166_v104_)[d_156_v95_] if (d_156_v95_) in (d_166_v104_) else p0), default__.abs(d_140_i2_))).cardinality)
                nw17_[int(18)] = (0) - (d_140_i2_)
                nw17_[int(19)] = (d_44_v3_)[default__.safeIndex(45, (d_44_v3_).length(0))]
                d_167_v105_ = nw17_
                (globalState).f11 = d_167_v105_
            r0 = p1
            d_170_v106_: C0
            nw18_ = C0()
            nw18_.ctor__()
            d_170_v106_ = nw18_
        d_171_v107_: C1
        nw19_ = C1()
        nw19_.ctor__(default__.fm21(p0, not(p0), globalState), p1)
        d_171_v107_ = nw19_
        d_172_v108_: _dafny.Set
        d_172_v108_ = _dafny.Set({d_171_v107_, d_171_v107_})
        d_172_v108_ = (_dafny.Set({d_171_v107_})).intersection(d_172_v108_)
        d_173_v109_: int
        d_174_v110_: int
        out0_: int
        out1_: int
        out0_, out1_ = (d_171_v107_).m0(p0, globalState)
        d_173_v109_ = out0_
        d_174_v110_ = out1_
        d_175_v111_: _dafny.MultiSet
        d_175_v111_ = _dafny.MultiSet([True])
        r0 = default__.safeDivisionInt((0) - ((default__.fm1(p0, d_48_v7_, d_173_v109_, globalState) if p0 else ((d_175_v111_)[p0] if (p0) in (d_175_v111_) else p1))), d_173_v109_)
        r1 = p0
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_176_v0_: bool
        d_176_v0_ = True
        d_177_v1_: int
        d_177_v1_ = 17
        d_178_v2_: _dafny.Seq
        d_178_v2_ = _dafny.SeqWithoutIsStrInference([d_177_v1_])
        d_179_v3_: _dafny.Map
        d_179_v3_ = _dafny.Map({d_176_v0_: d_178_v2_})
        d_180_v4_: _dafny.Array
        def lambda2_(d_181_v1_):
            def lambda3_(d_182_i1_):
                return (d_182_i1_) + (d_181_v1_)

            return lambda3_

        init1_ = lambda2_(d_177_v1_)
        nw20_ = _dafny.Array(None, 8)
        for i0_1_ in range(nw20_.length(0)):
            nw20_[i0_1_] = init1_(i0_1_)
        d_180_v4_ = nw20_
        d_183_v5_: _dafny.Seq
        d_183_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cb"))
        d_184_v6_: _dafny.MultiSet
        d_184_v6_ = _dafny.MultiSet([d_183_v5_])
        d_185_v8_: _dafny.MultiSet
        d_185_v8_ = _dafny.MultiSet([d_176_v0_, d_176_v0_, d_176_v0_])
        d_186_v9_: _dafny.Array
        nw21_ = _dafny.Array(None, 11)
        nw21_[int(0)] = d_176_v0_
        nw21_[int(1)] = False
        nw21_[int(2)] = d_176_v0_
        nw21_[int(3)] = d_176_v0_
        nw21_[int(4)] = d_176_v0_
        nw21_[int(5)] = d_176_v0_
        nw21_[int(6)] = d_176_v0_
        nw21_[int(7)] = d_176_v0_
        nw21_[int(8)] = d_176_v0_
        nw21_[int(9)] = d_176_v0_
        nw21_[int(10)] = d_176_v0_
        d_186_v9_ = nw21_
        d_187_v10_: str
        d_187_v10_ = _dafny.CodePoint('u')
        d_188_v11_: _dafny.MultiSet
        d_188_v11_ = _dafny.MultiSet([d_187_v10_, d_187_v10_, _dafny.CodePoint('f'), d_187_v10_])
        d_189_v12_: _dafny.Map
        d_189_v12_ = _dafny.Map({d_177_v1_: _dafny.MultiSet([d_187_v10_])})
        d_190_globalState_: GlobalState
        nw22_ = GlobalState()
        def iife10_():
            coll6_ = _dafny.Set()
            compr_6_: _dafny.Seq
            for compr_6_ in (d_184_v6_).Elements:
                d_192_v7_: _dafny.Seq = compr_6_
                if (d_192_v7_) in (d_184_v6_):
                    coll6_ = coll6_.union(_dafny.Set([d_192_v7_]))
            return _dafny.Set(coll6_)
        nw22_.ctor__(d_179_v3_, 201, -912, -807, False, 990, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_191_i0_ in range(default__.abs(144))]), 150, 944, True, 334, d_180_v4_, iife10_()
        , (d_183_v5_) + (d_183_v5_), True, 850, d_185_v8_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_193_i2_ in range(default__.abs(20))]), d_186_v9_, -870, d_183_v5_, (_dafny.Map({d_177_v1_: d_188_v11_})) | (d_189_v12_), False, False, d_183_v5_, True)
        d_190_globalState_ = nw22_
        d_194_v13_: _dafny.Seq
        d_194_v13_ = _dafny.SeqWithoutIsStrInference([d_180_v4_, d_180_v4_, d_180_v4_, d_180_v4_])
        d_194_v13_ = d_194_v13_
        d_195_v14_: _dafny.Map
        d_195_v14_ = _dafny.Map({(0) - (d_177_v1_): d_176_v0_})
        d_196_v15_: _dafny.Map
        d_196_v15_ = _dafny.Map({d_176_v0_: d_176_v0_})
        if not(default__.fm0(((d_195_v14_)[d_177_v1_] if (d_177_v1_) in (d_195_v14_) else d_176_v0_), (d_196_v15_) | (d_196_v15_), d_190_globalState_)):
            (d_190_globalState_).f10 = (0) - (d_177_v1_)
            (d_190_globalState_).f22 = not (d_176_v0_) or (d_176_v0_)
            index18_ = default__.safeIndex(884, (d_180_v4_).length(0))
            (d_180_v4_)[index18_] = 945
            index19_ = default__.safeIndex(884, (d_180_v4_).length(0))
            (d_180_v4_)[index19_] = default__.safeModuloInt(669, d_177_v1_)
            if d_176_v0_:
                (d_190_globalState_).f2 = ((d_180_v4_)[default__.safeIndex(884, (d_180_v4_).length(0))]) * ((d_180_v4_)[default__.safeIndex(884, (d_180_v4_).length(0))])
                d_197_v16_: _dafny.Seq
                d_197_v16_ = _dafny.SeqWithoutIsStrInference([d_176_v0_])
                d_198_v17_: _dafny.MultiSet
                d_198_v17_ = _dafny.MultiSet([len(d_178_v2_), len(d_197_v16_), d_177_v1_])
                d_199_v18_: _dafny.MultiSet
                d_199_v18_ = _dafny.MultiSet([d_198_v17_])
                d_200_v19_: _dafny.MultiSet
                d_200_v19_ = _dafny.MultiSet([((d_199_v18_)[d_198_v17_] if (d_198_v17_) in (d_199_v18_) else d_177_v1_)])
                d_201_v20_: _dafny.Map
                d_201_v20_ = _dafny.Map({d_176_v0_: (d_180_v4_)[default__.safeIndex(884, (d_180_v4_).length(0))]})
                d_176_v0_ = (((d_185_v8_)[d_176_v0_] if (d_176_v0_) in (d_185_v8_) else (d_180_v4_)[default__.safeIndex(884, (d_180_v4_).length(0))])) <= (default__.safeDivisionInt(((d_198_v17_)[default__.fm1(d_176_v0_, d_187_v10_, len(_dafny.Map({(_dafny.MultiSet(d_178_v2_)).cardinality: (d_180_v4_)[default__.safeIndex(884, (d_180_v4_).length(0))]})), d_190_globalState_)] if (default__.fm1(d_176_v0_, d_187_v10_, len(_dafny.Map({(_dafny.MultiSet(d_178_v2_)).cardinality: (d_180_v4_)[default__.safeIndex(884, (d_180_v4_).length(0))]})), d_190_globalState_)) in (d_198_v17_) else (d_180_v4_)[default__.safeIndex(884, (d_180_v4_).length(0))]), (0) - (len(_dafny.Map({len(d_201_v20_): False})))))
                (d_190_globalState_).f2 = default__.safeModuloInt(d_177_v1_, (d_180_v4_)[default__.safeIndex(884, (d_180_v4_).length(0))])
                d_202_v21_: _dafny.Seq
                d_202_v21_ = _dafny.SeqWithoutIsStrInference([(d_200_v19_) | (d_198_v17_), (_dafny.MultiSet(d_178_v2_)).intersection(default__.fm2(default__.fm1(d_176_v0_, default__.fm3(d_177_v1_, d_176_v0_, d_190_globalState_), d_177_v1_, d_190_globalState_), d_176_v0_, d_190_globalState_))])
                d_203_v22_: D0
                d_203_v22_ = D0_DC0(d_202_v21_)
                d_202_v21_ = (d_203_v22_).cf0
                d_183_v5_ = d_183_v5_
            elif True:
                d_183_v5_ = d_183_v5_
                d_204_v23_: C0
                nw23_ = C0()
                nw23_.ctor__()
                d_204_v23_ = nw23_
                d_205_v24_: _dafny.Array
                def lambda4_(d_206_v10_, d_207_v0_):
                    def lambda5_(d_208_i3_):
                        return D3_DC11(d_206_v10_, d_207_v0_, True)

                    return lambda5_

                init2_ = lambda4_(d_187_v10_, d_176_v0_)
                nw24_ = _dafny.Array(None, 25)
                for i0_2_ in range(nw24_.length(0)):
                    nw24_[i0_2_] = init2_(i0_2_)
                d_205_v24_ = nw24_
                d_209_v25_: D3
                d_209_v25_ = D3_DC11(d_187_v10_, d_176_v0_, d_176_v0_)
                index20_ = default__.safeIndex(210, (d_205_v24_).length(0))
                (d_205_v24_)[index20_] = d_209_v25_
                pat_let_tv5_ = d_180_v4_
                pat_let_tv6_ = d_180_v4_
                pat_let_tv7_ = d_176_v0_
                pat_let_tv8_ = d_190_globalState_
                index21_ = default__.safeIndex(210, (d_205_v24_).length(0))
                def iife11_(_pat_let2_0):
                    def iife12_(d_210_dt__update__tmp_h0_):
                        def iife13_(_pat_let3_0):
                            def iife14_(d_211_dt__update_hcf19_h0_):
                                return D3_DC11(d_211_dt__update_hcf19_h0_, (d_210_dt__update__tmp_h0_).cf20, (d_210_dt__update__tmp_h0_).cf21)
                            return iife14_(_pat_let3_0)
                        return iife13_(default__.fm3((pat_let_tv6_)[default__.safeIndex(884, (pat_let_tv5_).length(0))], pat_let_tv7_, pat_let_tv8_))
                    return iife12_(_pat_let2_0)
                (d_205_v24_)[index21_] = iife11_(d_209_v25_)
                d_212_v26_: str
                d_212_v26_ = d_187_v10_
                d_212_v26_ = d_187_v10_
                (d_190_globalState_).f19 = ((d_185_v8_)[(False) and (not(d_176_v0_))] if ((False) and (not(d_176_v0_))) in (d_185_v8_) else d_177_v1_)
            d_213_v27_: _dafny.Seq
            d_213_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_176_v0_, d_176_v0_}), _dafny.Set({d_176_v0_, d_176_v0_})])
            d_214_v28_: C1
            nw25_ = C1()
            nw25_.ctor__((d_213_v27_)[default__.safeIndex(d_177_v1_, len(d_213_v27_))], d_177_v1_)
            d_214_v28_ = nw25_
            d_215_v29_: _dafny.Seq
            d_215_v29_ = _dafny.SeqWithoutIsStrInference([d_176_v0_, not(d_176_v0_), (D3_DC11(d_187_v10_, d_176_v0_, d_176_v0_)).cf20])
            d_216_v30_: _dafny.Map
            d_216_v30_ = _dafny.Map({d_215_v29_: d_214_v28_})
            index22_ = default__.safeIndex(884, (d_180_v4_).length(0))
            rhs25_ = default__.safeModuloInt(len(d_178_v2_), d_177_v1_)
            rhs26_ = ((d_216_v30_)[((d_215_v29_ if False else _dafny.SeqWithoutIsStrInference([d_176_v0_, d_176_v0_, d_176_v0_]))).set(default__.safeIndex((0) - (d_177_v1_), len((d_215_v29_ if False else _dafny.SeqWithoutIsStrInference([d_176_v0_, d_176_v0_, d_176_v0_])))), d_176_v0_)] if (((d_215_v29_ if False else _dafny.SeqWithoutIsStrInference([d_176_v0_, d_176_v0_, d_176_v0_]))).set(default__.safeIndex((0) - (d_177_v1_), len((d_215_v29_ if False else _dafny.SeqWithoutIsStrInference([d_176_v0_, d_176_v0_, d_176_v0_])))), d_176_v0_)) in (d_216_v30_) else d_214_v28_)
            lhs21_ = d_180_v4_
            lhs22_ = default__.safeIndex(884, (d_180_v4_).length(0))
            lhs21_[lhs22_] = rhs25_
            d_214_v28_ = rhs26_
        elif True:
            if (d_177_v1_) > ((d_185_v8_).cardinality):
                d_217_v31_: T0
                nw26_ = C0()
                nw26_.ctor__()
                d_217_v31_ = nw26_
                d_218_v32_: _dafny.Seq
                d_218_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_177_v1_, d_177_v1_])])
                d_219_v33_: D0
                d_219_v33_ = D0_DC0((d_218_v32_) + (d_218_v32_))
                d_220_v34_: _dafny.Seq
                d_220_v34_ = _dafny.SeqWithoutIsStrInference([True])
                rhs27_ = (d_177_v1_ if False else (d_177_v1_ if (d_220_v34_)[default__.safeIndex(d_177_v1_, len(d_220_v34_))] else d_177_v1_))
                rhs28_ = (d_177_v1_) + (d_177_v1_)
                rhs29_ = d_217_v31_
                rhs30_ = d_176_v0_
                rhs31_ = d_219_v33_
                lhs23_ = d_190_globalState_
                lhs24_ = d_190_globalState_
                lhs25_ = d_190_globalState_
                lhs23_.f10 = rhs27_
                lhs24_.f19 = rhs28_
                d_217_v31_ = rhs29_
                lhs25_.f23 = rhs30_
                d_219_v33_ = rhs31_
                d_217_v31_ = d_217_v31_
                d_195_v14_ = (d_195_v14_).set(d_177_v1_, d_176_v0_)
                d_221_v35_: _dafny.Set
                d_221_v35_ = _dafny.Set({d_176_v0_, d_176_v0_})
                d_222_v36_: C1
                nw27_ = C1()
                nw27_.ctor__(d_221_v35_, (d_177_v1_) * (107))
                d_222_v36_ = nw27_
                d_223_v37_: _dafny.Map
                d_223_v37_ = _dafny.Map({d_222_v36_: d_187_v10_})
                d_223_v37_ = (d_223_v37_).set(d_222_v36_, _dafny.CodePoint('x'))
            elif True:
                d_224_v38_: _dafny.Array
                def lambda6_(d_225_v2_):
                    def lambda7_(d_226_i4_):
                        return d_225_v2_

                    return lambda7_

                init3_ = lambda6_(d_178_v2_)
                nw28_ = _dafny.Array(None, 4)
                for i0_3_ in range(nw28_.length(0)):
                    nw28_[i0_3_] = init3_(i0_3_)
                d_224_v38_ = nw28_
                d_227_v39_: _dafny.Set
                d_227_v39_ = _dafny.Set({default__.safeDivisionInt(d_177_v1_, d_177_v1_)})
                d_228_v40_: _dafny.Set
                d_228_v40_ = _dafny.Set({d_176_v0_, d_176_v0_, d_176_v0_, d_176_v0_})
                d_229_v41_: C0
                nw29_ = C0()
                nw29_.ctor__()
                d_229_v41_ = nw29_
                d_230_v42_: _dafny.Map
                d_230_v42_ = _dafny.Map({d_229_v41_: d_176_v0_})
                d_231_v43_: C1
                nw30_ = C1()
                nw30_.ctor__((d_228_v40_) | (_dafny.Set({((d_230_v42_)[d_229_v41_] if (d_229_v41_) in (d_230_v42_) else d_176_v0_), d_176_v0_, not(d_176_v0_)})), (d_229_v41_).fm4(d_177_v1_, d_176_v0_, d_190_globalState_))
                d_231_v43_ = nw30_
                d_232_v44_: _dafny.Seq
                d_232_v44_ = _dafny.SeqWithoutIsStrInference([d_231_v43_])
                rhs32_ = d_224_v38_
                rhs33_ = _dafny.Set({d_231_v43_.f27, default__.safeModuloInt(len(d_232_v44_), (d_178_v2_)[default__.safeIndex(len(d_183_v5_), len(d_178_v2_))]), d_177_v1_})
                rhs34_ = d_231_v43_
                rhs35_ = d_231_v43_.f27
                lhs26_ = d_190_globalState_
                d_224_v38_ = rhs32_
                d_227_v39_ = rhs33_
                d_231_v43_ = rhs34_
                lhs26_.f19 = rhs35_
                d_233_v45_: _dafny.Array
                def lambda8_(d_234_v8_):
                    def lambda9_(d_235_i5_):
                        return D4_DC13(d_234_v8_)

                    return lambda9_

                init4_ = lambda8_(d_185_v8_)
                nw31_ = _dafny.Array(None, 25)
                for i0_4_ in range(nw31_.length(0)):
                    nw31_[i0_4_] = init4_(i0_4_)
                d_233_v45_ = nw31_
                d_236_v46_: T0
                nw32_ = C1()
                nw32_.ctor__((d_231_v43_).f26, d_177_v1_)
                d_236_v46_ = nw32_
                d_237_v47_: _dafny.Seq
                d_237_v47_ = _dafny.SeqWithoutIsStrInference([d_236_v46_, d_236_v46_])
                d_238_v48_: _dafny.Map
                d_238_v48_ = _dafny.Map({(d_237_v47_)[default__.safeIndex((d_178_v2_)[default__.safeIndex(d_177_v1_, len(d_178_v2_))], len(d_237_v47_))]: d_177_v1_})
                d_239_v49_: _dafny.Map
                d_239_v49_ = _dafny.Map({d_176_v0_: d_233_v45_})
                rhs36_ = ((d_239_v49_)[(False if d_176_v0_ else d_176_v0_)] if ((False if d_176_v0_ else d_176_v0_)) in (d_239_v49_) else d_233_v45_)
                rhs37_ = d_238_v48_
                rhs38_ = d_231_v43_.f27
                lhs27_ = d_190_globalState_
                d_233_v45_ = rhs36_
                d_238_v48_ = rhs37_
                lhs27_.f15 = rhs38_
                index23_ = default__.safeIndex(323, (d_180_v4_).length(0))
                (d_180_v4_)[index23_] = d_231_v43_.f27
                index24_ = default__.safeIndex(323, (d_180_v4_).length(0))
                (d_180_v4_)[index24_] = d_231_v43_.f27
                d_240_v50_: _dafny.Array
                nw33_ = _dafny.Array(_dafny.Map({}), 18)
                d_240_v50_ = nw33_
                d_241_v51_: _dafny.Map
                d_241_v51_ = _dafny.Map({d_176_v0_: d_177_v1_})
                index25_ = default__.safeIndex(841, (d_240_v50_).length(0))
                (d_240_v50_)[index25_] = d_241_v51_
                index26_ = default__.safeIndex(841, (d_240_v50_).length(0))
                (d_240_v50_)[index26_] = d_241_v51_
                d_242_v52_: D0
                d_242_v52_ = D0_DC1((d_180_v4_)[default__.safeIndex(323, (d_180_v4_).length(0))])
                d_243_v53_: D3
                d_243_v53_ = D3_DC12(False, d_242_v52_, d_176_v0_, d_176_v0_, d_236_v46_)
                d_244_v54_: _dafny.Array
                nw34_ = _dafny.Array(None, 19)
                nw34_[int(0)] = d_236_v46_
                nw34_[int(1)] = d_236_v46_
                nw34_[int(2)] = d_236_v46_
                nw34_[int(3)] = d_236_v46_
                nw34_[int(4)] = d_236_v46_
                nw34_[int(5)] = d_236_v46_
                nw34_[int(6)] = d_236_v46_
                nw34_[int(7)] = (d_243_v53_).cf26
                nw34_[int(8)] = d_236_v46_
                nw34_[int(9)] = d_236_v46_
                nw34_[int(10)] = d_236_v46_
                nw34_[int(11)] = d_236_v46_
                nw34_[int(12)] = ((D3_DC12(d_176_v0_, d_242_v52_, d_176_v0_, d_176_v0_, d_236_v46_)).cf26 if d_176_v0_ else d_236_v46_)
                nw34_[int(13)] = d_236_v46_
                nw34_[int(14)] = d_236_v46_
                nw34_[int(15)] = d_236_v46_
                nw34_[int(16)] = d_236_v46_
                nw34_[int(17)] = d_236_v46_
                nw34_[int(18)] = d_236_v46_
                d_244_v54_ = nw34_
                index27_ = default__.safeIndex(39, (d_244_v54_).length(0))
                (d_244_v54_)[index27_] = d_236_v46_
                index28_ = default__.safeIndex(39, (d_244_v54_).length(0))
                (d_244_v54_)[index28_] = d_236_v46_
            d_245_v55_: D2
            d_245_v55_ = D2_DC8(len(d_183_v5_), d_176_v0_, d_177_v1_)
            if (d_245_v55_).cf13:
                d_246_v56_: D3
                d_246_v56_ = D3_DC11(d_187_v10_, d_176_v0_, d_176_v0_)
                pat_let_tv9_ = d_176_v0_
                d_247_v57_: _dafny.Seq
                def iife15_(_pat_let4_0):
                    def iife16_(d_248_dt__update__tmp_h2_):
                        def iife17_(_pat_let5_0):
                            def iife18_(d_249_dt__update_hcf21_h0_):
                                return D3_DC11((d_248_dt__update__tmp_h2_).cf19, (d_248_dt__update__tmp_h2_).cf20, d_249_dt__update_hcf21_h0_)
                            return iife18_(_pat_let5_0)
                        return iife17_(pat_let_tv9_)
                    return iife16_(_pat_let4_0)
                d_247_v57_ = _dafny.SeqWithoutIsStrInference([iife15_(d_246_v56_), default__.fm11(d_177_v1_, d_190_globalState_), D3_DC11(d_187_v10_, d_176_v0_, d_176_v0_)])
                (d_190_globalState_).f5 = len((d_247_v57_) + (_dafny.SeqWithoutIsStrInference([D3_DC11(d_187_v10_, d_176_v0_, d_176_v0_) for d_250_i6_ in range(default__.abs(876))])))
                (d_190_globalState_).f25 = d_176_v0_
                (d_190_globalState_).f2 = (d_177_v1_) + (default__.safeDivisionInt(d_177_v1_, d_177_v1_))
                d_251_v58_: _dafny.Array
                def lambda10_(d_252_v10_, d_253_v0_):
                    def lambda11_(d_254_i7_):
                        return D3_DC11(d_252_v10_, not(d_253_v0_), not(not(not(d_253_v0_))))

                    return lambda11_

                init5_ = lambda10_(d_187_v10_, d_176_v0_)
                nw35_ = _dafny.Array(None, 3)
                for i0_5_ in range(nw35_.length(0)):
                    nw35_[i0_5_] = init5_(i0_5_)
                d_251_v58_ = nw35_
                index29_ = default__.safeIndex(610, (d_251_v58_).length(0))
                (d_251_v58_)[index29_] = d_246_v56_
                index30_ = default__.safeIndex(610, (d_251_v58_).length(0))
                (d_251_v58_)[index30_] = default__.fm11(d_177_v1_, d_190_globalState_)
                d_255_v59_: D0
                d_255_v59_ = D0_DC1(d_177_v1_)
                d_256_v60_: _dafny.Set
                d_256_v60_ = _dafny.Set({d_176_v0_, d_176_v0_})
                d_257_v61_: T0
                nw36_ = C1()
                nw36_.ctor__(d_256_v60_, d_177_v1_)
                d_257_v61_ = nw36_
                d_258_v62_: D3
                d_258_v62_ = D3_DC12(d_176_v0_, d_255_v59_, d_176_v0_, d_176_v0_, d_257_v61_)
                d_259_v63_: _dafny.Map
                d_259_v63_ = _dafny.Map({d_258_v62_: d_177_v1_})
                d_260_v64_: _dafny.Map
                d_260_v64_ = _dafny.Map({d_177_v1_: d_259_v63_})
                (d_190_globalState_).f22 = not((_dafny.Map({d_258_v62_: d_177_v1_})) == (((d_260_v64_)[(0) - (len(d_178_v2_))] if ((0) - (len(d_178_v2_))) in (d_260_v64_) else d_259_v63_)))
            elif True:
                d_261_v65_: _dafny.Set
                d_261_v65_ = _dafny.Set({d_176_v0_})
                d_262_v66_: C1
                nw37_ = C1()
                nw37_.ctor__(d_261_v65_, d_177_v1_)
                d_262_v66_ = nw37_
                d_263_v67_: D0
                d_263_v67_ = D0_DC1(-492)
                d_264_v68_: _dafny.Array
                nw38_ = _dafny.Array(None, 13)
                nw38_[int(0)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwxonbmhb"))) for d_265_i8_ in range(default__.abs(28))])
                nw38_[int(1)] = (d_178_v2_) + (default__.fm12(d_177_v1_, d_263_v67_, d_190_globalState_))
                nw38_[int(2)] = d_178_v2_
                nw38_[int(3)] = _dafny.SeqWithoutIsStrInference([d_262_v66_.f27])
                nw38_[int(4)] = d_178_v2_
                nw38_[int(5)] = d_178_v2_
                nw38_[int(6)] = (d_178_v2_) + (_dafny.SeqWithoutIsStrInference([((d_185_v8_).set(d_176_v0_, default__.abs((0) - (d_262_v66_.f27)))).cardinality for d_266_i9_ in range(default__.abs(762))]))
                nw38_[int(7)] = d_178_v2_
                nw38_[int(8)] = d_178_v2_
                nw38_[int(9)] = (d_178_v2_ if d_176_v0_ else d_178_v2_)
                nw38_[int(10)] = d_178_v2_
                nw38_[int(11)] = d_178_v2_
                nw38_[int(12)] = d_178_v2_
                d_264_v68_ = nw38_
                index31_ = default__.safeIndex(273, (d_264_v68_).length(0))
                (d_264_v68_)[index31_] = (d_178_v2_) + (d_178_v2_)
                d_267_v69_: _dafny.Set
                d_267_v69_ = _dafny.Set({d_262_v66_.f27, d_177_v1_})
                d_268_v70_: _dafny.Seq
                d_268_v70_ = _dafny.SeqWithoutIsStrInference([((d_178_v2_).set(default__.safeIndex(len(d_267_v69_), len(d_178_v2_)), d_262_v66_.f27)).set(default__.safeIndex(d_262_v66_.f27, len((d_178_v2_).set(default__.safeIndex(len(d_267_v69_), len(d_178_v2_)), d_262_v66_.f27))), d_177_v1_), _dafny.SeqWithoutIsStrInference([d_177_v1_ for d_269_i10_ in range(default__.abs(-788))]), d_178_v2_, (_dafny.SeqWithoutIsStrInference([d_262_v66_.f27 for d_270_i11_ in range(default__.abs(377))])) + (d_178_v2_)])
                index32_ = default__.safeIndex(273, (d_264_v68_).length(0))
                (d_264_v68_)[index32_] = (d_268_v70_)[default__.safeIndex(d_177_v1_, len(d_268_v70_))]
                (d_190_globalState_).f15 = d_262_v66_.f27
                d_271_v71_: D2
                d_271_v71_ = D2_DC7(d_176_v0_)
                rhs39_ = d_271_v71_
                rhs40_ = (d_262_v66_.f27) * (d_177_v1_)
                lhs28_ = d_190_globalState_
                d_271_v71_ = rhs39_
                lhs28_.f8 = rhs40_
                (d_190_globalState_).f22 = (d_183_v5_) <= (d_183_v5_)
            (d_190_globalState_).f10 = default__.fm1(d_176_v0_, d_187_v10_, d_177_v1_, d_190_globalState_)
            (d_190_globalState_).f19 = d_177_v1_
            if (d_185_v8_).isdisjoint(_dafny.MultiSet([d_176_v0_, d_176_v0_, True])):
                d_272_v72_: _dafny.Map
                d_272_v72_ = _dafny.Map({len(d_183_v5_): d_177_v1_})
                (d_190_globalState_).f23 = (d_272_v72_) not in (default__.fm13(d_190_globalState_))
                d_273_v73_: _dafny.Map
                d_273_v73_ = _dafny.Map({d_185_v8_: 363})
                d_273_v73_ = (d_273_v73_).set(_dafny.MultiSet([d_176_v0_]), d_177_v1_)
                index33_ = default__.safeIndex(451, (d_180_v4_).length(0))
                (d_180_v4_)[index33_] = len(((d_183_v5_) + (d_183_v5_)).set(default__.safeIndex(d_177_v1_, len((d_183_v5_) + (d_183_v5_))), d_187_v10_))
                index34_ = default__.safeIndex(451, (d_180_v4_).length(0))
                (d_180_v4_)[index34_] = d_177_v1_
                d_274_v74_: _dafny.Map
                d_274_v74_ = _dafny.Map({_dafny.Map({(d_180_v4_)[default__.safeIndex(451, (d_180_v4_).length(0))]: (0) - (d_177_v1_)}): d_176_v0_})
                rhs41_ = (d_274_v74_) | (d_274_v74_)
                rhs42_ = 465
                rhs43_ = ((d_178_v2_) + (d_178_v2_)) + (_dafny.SeqWithoutIsStrInference([d_177_v1_ for d_275_i12_ in range(default__.abs(23))]))
                rhs44_ = default__.fm0((d_176_v0_) == (d_176_v0_), _dafny.Map({d_176_v0_: default__.fm0(d_176_v0_, d_196_v15_, d_190_globalState_)}), d_190_globalState_)
                rhs45_ = (d_180_v4_)[default__.safeIndex(451, (d_180_v4_).length(0))]
                lhs29_ = d_190_globalState_
                lhs30_ = d_190_globalState_
                lhs31_ = d_190_globalState_
                d_274_v74_ = rhs41_
                lhs29_.f15 = rhs42_
                d_178_v2_ = rhs43_
                lhs30_.f25 = rhs44_
                lhs31_.f2 = rhs45_
                d_276_v75_: _dafny.Set
                d_276_v75_ = _dafny.Set({d_176_v0_})
                d_277_v76_: _dafny.Array
                nw39_ = _dafny.Array(int(0), 16)
                d_277_v76_ = nw39_
                d_278_v77_: _dafny.Set
                d_278_v77_ = _dafny.Set({d_277_v76_})
                d_279_v78_: _dafny.Seq
                d_279_v78_ = _dafny.SeqWithoutIsStrInference([d_176_v0_])
                d_280_v79_: D0
                d_280_v79_ = D0_DC2(d_176_v0_, d_278_v77_, not(d_176_v0_), d_176_v0_, len(d_279_v78_))
                d_281_v80_: C1
                nw40_ = C1()
                nw40_.ctor__(d_276_v75_, (d_280_v79_).cf6)
                d_281_v80_ = nw40_
                d_282_v81_: _dafny.Set
                d_282_v81_ = _dafny.Set({d_281_v80_})
                (d_190_globalState_).f25 = (len(d_282_v81_)) <= (549)
            elif True:
                d_178_v2_ = _dafny.SeqWithoutIsStrInference([d_177_v1_ for d_283_i13_ in range(default__.abs(762))])
                d_284_v82_: _dafny.Set
                d_284_v82_ = _dafny.Set({d_176_v0_, d_176_v0_, d_176_v0_})
                d_285_v83_: str
                d_285_v83_ = d_187_v10_
                d_286_v84_: _dafny.Seq
                d_286_v84_ = _dafny.SeqWithoutIsStrInference([d_176_v0_, d_176_v0_])
                d_287_v85_: _dafny.Set
                d_287_v85_ = _dafny.Set({d_177_v1_})
                rhs46_ = d_180_v4_
                rhs47_ = default__.fm0((True) == (default__.fm0(d_176_v0_, d_196_v15_, d_190_globalState_)), (d_196_v15_) | (d_196_v15_), d_190_globalState_)
                rhs48_ = (d_185_v8_).issubset((_dafny.MultiSet(d_286_v84_)) | (d_185_v8_))
                rhs49_ = _dafny.Set({d_176_v0_, d_176_v0_, not(d_176_v0_), (d_286_v84_)[default__.safeIndex(len(d_287_v85_), len(d_286_v84_))]})
                rhs50_ = d_187_v10_
                lhs32_ = d_190_globalState_
                lhs33_ = d_190_globalState_
                lhs32_.f11 = rhs46_
                lhs33_.f23 = rhs47_
                d_176_v0_ = rhs48_
                d_284_v82_ = rhs49_
                d_285_v83_ = rhs50_
                d_288_v86_: _dafny.Array
                nw41_ = _dafny.Array(None, 9)
                nw41_[int(0)] = d_187_v10_
                nw41_[int(1)] = d_187_v10_
                nw41_[int(2)] = d_187_v10_
                nw41_[int(3)] = d_187_v10_
                nw41_[int(4)] = d_187_v10_
                nw41_[int(5)] = d_187_v10_
                nw41_[int(6)] = _dafny.CodePoint('k')
                nw41_[int(7)] = d_187_v10_
                nw41_[int(8)] = _dafny.CodePoint('u')
                d_288_v86_ = nw41_
                index35_ = default__.safeIndex(28, (d_288_v86_).length(0))
                (d_288_v86_)[index35_] = _dafny.CodePoint('h')
                index36_ = default__.safeIndex(28, (d_288_v86_).length(0))
                (d_288_v86_)[index36_] = d_187_v10_
                (d_190_globalState_).f22 = d_176_v0_
                d_289_v87_: _dafny.MultiSet
                d_289_v87_ = _dafny.MultiSet([d_177_v1_])
                d_290_v88_: _dafny.Seq
                d_290_v88_ = _dafny.SeqWithoutIsStrInference([d_289_v87_, d_289_v87_, d_289_v87_])
                d_291_v89_: D0
                d_291_v89_ = D0_DC0(d_290_v88_)
                d_291_v89_ = default__.fm14(d_190_globalState_)
        (d_190_globalState_).f22 = ((d_196_v15_ if False else d_196_v15_)) != (d_196_v15_)
        d_292_v90_: _dafny.Set
        d_292_v90_ = _dafny.Set({d_177_v1_, d_177_v1_})
        if (d_292_v90_).issubset((d_292_v90_) - (d_292_v90_)):
            if not (d_176_v0_) or (d_176_v0_):
                d_293_v91_: _dafny.MultiSet
                d_293_v91_ = _dafny.MultiSet([d_177_v1_])
                d_294_v92_: D2
                d_294_v92_ = D2_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvlnsoaai")))
                d_295_v93_: D4
                d_295_v93_ = D4_DC14(d_293_v91_, d_177_v1_, d_294_v92_)
                (d_190_globalState_).f2 = (d_177_v1_ if d_176_v0_ else (d_295_v93_).cf29)
                (d_190_globalState_).f22 = default__.fm0(d_176_v0_, d_196_v15_, d_190_globalState_)
                d_296_v94_: C0
                nw42_ = C0()
                nw42_.ctor__()
                d_296_v94_ = nw42_
                d_297_v95_: _dafny.Map
                d_297_v95_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_187_v10_ for d_298_i14_ in range(default__.abs(-221))]): (d_183_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkvceykyp")))})
                d_297_v95_ = (d_297_v95_).set(d_183_v5_, d_183_v5_)
                d_299_v96_: _dafny.Map
                d_299_v96_ = _dafny.Map({d_177_v1_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgejkifdl"))})
                d_300_v97_: _dafny.Map
                d_300_v97_ = _dafny.Map({d_299_v96_: d_177_v1_})
                d_300_v97_ = (d_300_v97_).set(d_299_v96_, d_177_v1_)
            elif True:
                d_187_v10_ = _dafny.CodePoint('w')
                (d_190_globalState_).f8 = default__.safeDivisionInt(d_177_v1_, d_177_v1_)
                index37_ = default__.safeIndex(876, (d_186_v9_).length(0))
                (d_186_v9_)[index37_] = d_176_v0_
                index38_ = default__.safeIndex(873, (d_186_v9_).length(0))
                (d_186_v9_)[index38_] = not (d_176_v0_) or (d_176_v0_)
                index39_ = default__.safeIndex(876, (d_186_v9_).length(0))
                index40_ = default__.safeIndex(873, (d_186_v9_).length(0))
                rhs51_ = d_176_v0_
                rhs52_ = d_187_v10_
                rhs53_ = d_176_v0_
                rhs54_ = d_176_v0_
                lhs34_ = d_190_globalState_
                lhs35_ = d_186_v9_
                lhs36_ = default__.safeIndex(876, (d_186_v9_).length(0))
                lhs37_ = d_186_v9_
                lhs38_ = default__.safeIndex(873, (d_186_v9_).length(0))
                lhs34_.f23 = rhs51_
                d_187_v10_ = rhs52_
                lhs35_[lhs36_] = rhs53_
                lhs37_[lhs38_] = rhs54_
                d_301_v98_: int
                d_302_v99_: bool
                out2_: int
                out3_: bool
                out2_, out3_ = default__.m1(((d_196_v15_)[(d_186_v9_)[default__.safeIndex(876, (d_186_v9_).length(0))]] if ((d_186_v9_)[default__.safeIndex(876, (d_186_v9_).length(0))]) in (d_196_v15_) else d_176_v0_), default__.safeDivisionInt(d_177_v1_, 140), d_190_globalState_)
                d_301_v98_ = out2_
                d_302_v99_ = out3_
                index41_ = default__.safeIndex(876, (d_186_v9_).length(0))
                (d_186_v9_)[index41_] = (default__.fm15(d_190_globalState_)).isdisjoint(_dafny.Set({d_301_v98_, d_177_v1_, d_177_v1_, (0) - (d_177_v1_)}))
            d_303_v100_: C0
            nw43_ = C0()
            nw43_.ctor__()
            d_303_v100_ = nw43_
            d_304_v101_: _dafny.Seq
            d_304_v101_ = _dafny.SeqWithoutIsStrInference([(d_176_v0_ if d_176_v0_ else d_176_v0_), d_176_v0_, not(d_176_v0_)])
            d_305_v102_: _dafny.Map
            d_305_v102_ = _dafny.Map({d_177_v1_: d_177_v1_})
            rhs55_ = (d_303_v100_).fm4(((d_305_v102_)[d_177_v1_] if (d_177_v1_) in (d_305_v102_) else d_177_v1_), (d_177_v1_) != (len(d_304_v101_)), d_190_globalState_)
            rhs56_ = (0) - (d_177_v1_)
            rhs57_ = (d_304_v101_).set(default__.safeIndex(d_177_v1_, len(d_304_v101_)), (d_177_v1_) not in (d_195_v14_))
            rhs58_ = d_177_v1_
            lhs39_ = d_190_globalState_
            lhs40_ = d_190_globalState_
            lhs41_ = d_190_globalState_
            lhs39_.f10 = rhs55_
            lhs40_.f8 = rhs56_
            d_304_v101_ = rhs57_
            lhs41_.f19 = rhs58_
            d_306_v103_: _dafny.Set
            d_306_v103_ = _dafny.Set({d_303_v100_, d_303_v100_, d_303_v100_, d_303_v100_, d_303_v100_})
            d_307_v104_: _dafny.MultiSet
            d_307_v104_ = _dafny.MultiSet([d_306_v103_])
            d_308_v105_: _dafny.Seq
            d_308_v105_ = _dafny.SeqWithoutIsStrInference([d_306_v103_])
            d_307_v104_ = _dafny.MultiSet((d_308_v105_).set(default__.safeIndex(d_177_v1_, len(d_308_v105_)), d_306_v103_))
            (d_190_globalState_).f8 = len(d_183_v5_)
        elif True:
            d_183_v5_ = d_183_v5_
            d_309_v106_: D3
            d_309_v106_ = D3_DC10((len(_dafny.SeqWithoutIsStrInference([d_177_v1_]))) <= (d_177_v1_), (0) - (d_177_v1_), d_176_v0_)
            d_309_v106_ = d_309_v106_
            index42_ = default__.safeIndex(924, (d_186_v9_).length(0))
            (d_186_v9_)[index42_] = d_176_v0_
            d_310_v107_: D2
            d_310_v107_ = D2_DC8(d_177_v1_, d_176_v0_, d_177_v1_)
            index43_ = default__.safeIndex(924, (d_186_v9_).length(0))
            rhs59_ = (d_176_v0_) == (d_176_v0_)
            rhs60_ = default__.fm1((d_176_v0_ if d_176_v0_ else d_176_v0_), d_187_v10_, default__.safeModuloInt(len(d_183_v5_), d_177_v1_), d_190_globalState_)
            rhs61_ = (d_310_v107_).cf12
            lhs42_ = d_186_v9_
            lhs43_ = default__.safeIndex(924, (d_186_v9_).length(0))
            lhs44_ = d_190_globalState_
            lhs42_[lhs43_] = rhs59_
            lhs44_.f2 = rhs60_
            d_177_v1_ = rhs61_
            d_311_v108_: _dafny.Map
            d_311_v108_ = _dafny.Map({685: d_180_v4_})
            d_311_v108_ = (d_311_v108_).set((len(d_183_v5_)) - (d_177_v1_), d_180_v4_)
            d_312_v109_: _dafny.Set
            d_312_v109_ = _dafny.Set({(d_186_v9_)[default__.safeIndex(924, (d_186_v9_).length(0))]})
            d_313_v110_: C1
            nw44_ = C1()
            nw44_.ctor__(d_312_v109_, d_177_v1_)
            d_313_v110_ = nw44_
        d_314_v111_: int
        d_315_v112_: bool
        out4_: int
        out5_: bool
        out4_, out5_ = default__.m1(False, 292, d_190_globalState_)
        d_314_v111_ = out4_
        d_315_v112_ = out5_
        if d_176_v0_:
            (d_190_globalState_).f23 = d_176_v0_
            d_316_v113_: _dafny.MultiSet
            d_316_v113_ = _dafny.MultiSet([d_177_v1_, d_177_v1_])
            d_316_v113_ = (d_316_v113_) | (_dafny.MultiSet([d_314_v111_, d_177_v1_]))
            if (d_292_v90_).ispropersubset(d_292_v90_):
                index44_ = default__.safeIndex(641, (d_180_v4_).length(0))
                (d_180_v4_)[index44_] = d_314_v111_
                index45_ = default__.safeIndex(641, (d_180_v4_).length(0))
                (d_180_v4_)[index45_] = (d_177_v1_) - ((714 if d_176_v0_ else (0) - (d_314_v111_)))
                index46_ = default__.safeIndex(761, (d_186_v9_).length(0))
                def iife19_():
                    coll7_ = _dafny.Set()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(722, 354):
                        d_317_v114_: int = compr_7_
                        if ((722) <= (d_317_v114_)) and ((d_317_v114_) < (354)):
                            coll7_ = coll7_.union(_dafny.Set([default__.safeModuloInt(d_317_v114_, -455)]))
                    return _dafny.Set(coll7_)
                (d_186_v9_)[index46_] = (default__.fm15(d_190_globalState_)).ispropersubset(iife19_()
                )
                index47_ = default__.safeIndex(761, (d_186_v9_).length(0))
                (d_186_v9_)[index47_] = d_315_v112_
                d_196_v15_ = (d_196_v15_).set(default__.fm0(d_315_v112_, d_196_v15_, d_190_globalState_), d_176_v0_)
                d_318_v115_: _dafny.Set
                d_318_v115_ = _dafny.Set({d_315_v112_})
                d_319_v116_: C1
                nw45_ = C1()
                nw45_.ctor__(d_318_v115_, len(d_318_v115_))
                d_319_v116_ = nw45_
                d_320_v117_: _dafny.MultiSet
                d_320_v117_ = _dafny.MultiSet([d_319_v116_])
                d_321_v118_: C1
                nw46_ = C1()
                nw46_.ctor__(_dafny.Set({d_176_v0_}), ((d_320_v117_)[d_319_v116_] if (d_319_v116_) in (d_320_v117_) else d_177_v1_))
                d_321_v118_ = nw46_
                (d_190_globalState_).f2 = d_319_v116_.f27
            elif True:
                (d_190_globalState_).f23 = d_176_v0_
                d_186_v9_ = d_186_v9_
                d_322_v119_: C0
                nw47_ = C0()
                nw47_.ctor__()
                d_322_v119_ = nw47_
                d_186_v9_ = d_186_v9_
                d_323_v120_: int
                d_324_v121_: bool
                out6_: int
                out7_: bool
                out6_, out7_ = default__.m1(d_176_v0_, (-429) + (d_314_v111_), d_190_globalState_)
                d_323_v120_ = out6_
                d_324_v121_ = out7_
            (d_190_globalState_).f23 = d_315_v112_
            d_325_v122_: int
            d_326_v123_: bool
            out8_: int
            out9_: bool
            out8_, out9_ = default__.m1(d_176_v0_, (_dafny.MultiSet([d_314_v111_])).cardinality, d_190_globalState_)
            d_325_v122_ = out8_
            d_326_v123_ = out9_
        elif True:
            d_327_v124_: D0
            d_327_v124_ = D0_DC1(d_177_v1_)
            (d_190_globalState_).f5 = default__.fm1(d_315_v112_, d_187_v10_, (d_327_v124_).cf1, d_190_globalState_)
            d_195_v14_ = (d_195_v14_).set(438, d_176_v0_)
            d_328_v125_: _dafny.Map
            d_328_v125_ = _dafny.Map({d_180_v4_: (0) - (((d_185_v8_).intersection(d_185_v8_)).cardinality)})
            d_328_v125_ = (d_328_v125_).set(d_180_v4_, d_314_v111_)
            d_186_v9_ = d_186_v9_
            d_329_v126_: int
            d_330_v127_: bool
            out10_: int
            out11_: bool
            out10_, out11_ = default__.m1(d_315_v112_, len((default__.fm10(d_177_v1_, 288, d_176_v0_, d_176_v0_, d_190_globalState_) if d_176_v0_ else d_183_v5_)), d_190_globalState_)
            d_329_v126_ = out10_
            d_330_v127_ = out11_
        d_331_v128_: _dafny.Seq
        d_331_v128_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_176_v0_, d_315_v112_, d_176_v0_, d_315_v112_, d_315_v112_])])
        d_332_v129_: _dafny.Set
        d_332_v129_ = _dafny.Set({d_185_v8_, d_185_v8_, d_185_v8_, (d_331_v128_)[default__.safeIndex(d_177_v1_, len(d_331_v128_))], d_185_v8_})
        d_333_v130_: _dafny.Map
        d_333_v130_ = _dafny.Map({d_185_v8_: (d_314_v111_) == (len(d_332_v129_))})
        d_333_v130_ = (d_333_v130_).set(d_185_v8_, d_176_v0_)
        d_183_v5_ = d_183_v5_
        d_334_v131_: D3
        d_334_v131_ = D3_DC10(d_176_v0_, d_177_v1_, False)
        d_335_v132_: _dafny.Seq
        d_335_v132_ = _dafny.SeqWithoutIsStrInference([True])
        d_336_v133_: int
        d_337_v134_: bool
        out12_: int
        out13_: bool
        out12_, out13_ = default__.m1(not((d_334_v131_).cf16), ((_dafny.MultiSet(d_335_v132_)).intersection(d_185_v8_)).cardinality, d_190_globalState_)
        d_336_v133_ = out12_
        d_337_v134_ = out13_
        d_337_v134_ = d_337_v134_
        d_338_v135_: _dafny.MultiSet
        d_338_v135_ = _dafny.MultiSet([60])
        d_339_v136_: _dafny.Map
        d_339_v136_ = _dafny.Map({d_338_v135_: d_336_v133_})
        rhs62_ = d_339_v136_
        rhs63_ = default__.safeModuloInt((d_177_v1_) * (d_314_v111_), (d_314_v111_) + (d_314_v111_))
        rhs64_ = d_336_v133_
        rhs65_ = (0) - (d_177_v1_)
        lhs45_ = d_190_globalState_
        lhs46_ = d_190_globalState_
        d_339_v136_ = rhs62_
        d_314_v111_ = rhs63_
        lhs45_.f19 = rhs64_
        lhs46_.f5 = rhs65_
        if d_337_v134_:
            d_340_v137_: _dafny.Set
            d_340_v137_ = _dafny.Set({d_176_v0_})
            d_341_v138_: C1
            nw48_ = C1()
            nw48_.ctor__((d_340_v137_) | (d_340_v137_), d_336_v133_)
            d_341_v138_ = nw48_
            (d_190_globalState_).f25 = default__.fm0(d_315_v112_, (d_196_v15_) | (d_196_v15_), d_190_globalState_)
            d_342_v139_: _dafny.Array
            nw49_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_342_v139_ = nw49_
            d_343_v140_: _dafny.Array
            nw50_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_343_v140_ = nw50_
            index48_ = default__.safeIndex(116, (d_342_v139_).length(0))
            (d_342_v139_)[index48_] = d_343_v140_
            index49_ = default__.safeIndex(116, (d_342_v139_).length(0))
            (d_342_v139_)[index49_] = d_343_v140_
            d_344_v141_: int
            d_345_v142_: int
            out14_: int
            out15_: int
            out14_, out15_ = (d_341_v138_).m0((d_176_v0_ if not(d_337_v134_) else d_176_v0_), d_190_globalState_)
            d_344_v141_ = out14_
            d_345_v142_ = out15_
            (d_190_globalState_).f23 = d_337_v134_
        elif True:
            (d_190_globalState_).f5 = d_336_v133_
            d_346_v143_: D2
            d_346_v143_ = D2_DC8(d_336_v133_, d_315_v112_, len(d_183_v5_))
            d_347_v144_: _dafny.Map
            d_347_v144_ = _dafny.Map({d_315_v112_: d_177_v1_})
            rhs66_ = d_336_v133_
            rhs67_ = _dafny.SeqWithoutIsStrInference([d_337_v134_, d_315_v112_, (d_337_v134_ if default__.fm0(d_337_v134_, d_196_v15_, d_190_globalState_) else not(d_337_v134_)), (d_315_v112_ if (d_346_v143_).cf13 else d_315_v112_), (d_336_v133_) < (default__.fm1(d_315_v112_, _dafny.CodePoint('w'), len(d_347_v144_), d_190_globalState_))])
            rhs68_ = d_176_v0_
            lhs47_ = d_190_globalState_
            lhs47_.f5 = rhs66_
            d_335_v132_ = rhs67_
            d_176_v0_ = rhs68_
            d_348_v145_: _dafny.Set
            d_348_v145_ = _dafny.Set({False, True})
            d_349_v146_: T0
            nw51_ = C1()
            nw51_.ctor__(d_348_v145_, d_177_v1_)
            d_349_v146_ = nw51_
            pat_let_tv10_ = d_183_v5_
            d_350_v147_: D3
            def iife20_(_pat_let6_0):
                def iife21_(d_351_dt__update__tmp_h4_):
                    def iife22_(_pat_let7_0):
                        def iife23_(d_352_dt__update_hcf1_h0_):
                            return D0_DC1(d_352_dt__update_hcf1_h0_)
                        return iife23_(_pat_let7_0)
                    return iife22_(len(pat_let_tv10_))
                return iife21_(_pat_let6_0)
            d_350_v147_ = D3_DC12(d_315_v112_, iife20_(D0_DC1(d_336_v133_)), (d_334_v131_).cf18, False, d_349_v146_)
            d_315_v112_ = (d_350_v147_).cf22
            d_353_v149_: _dafny.Map
            def iife24_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in _dafny.IntegerRange(967, 310):
                    d_354_v148_: int = compr_8_
                    if ((967) <= (d_354_v148_)) and ((d_354_v148_) < (310)):
                        coll8_[(d_354_v148_) + (d_336_v133_)] = d_337_v134_
                return _dafny.Map(coll8_)
            d_353_v149_ = _dafny.Map({(0) - (default__.safeDivisionInt(d_177_v1_, d_336_v133_)): len(iife24_()
            )})
            d_355_v150_: _dafny.Seq
            d_355_v150_ = _dafny.SeqWithoutIsStrInference([d_183_v5_])
            d_353_v149_ = (d_353_v149_).set(d_177_v1_, len(d_355_v150_))
            d_356_v151_: C1
            nw52_ = C1()
            nw52_.ctor__((d_348_v145_ if True else _dafny.Set({d_337_v134_, d_176_v0_})), d_177_v1_)
            d_356_v151_ = nw52_
        d_357_v152_: D0
        d_357_v152_ = D0_DC0(default__.fm16(d_177_v1_, d_190_globalState_))
        pat_let_tv11_ = d_187_v10_
        pat_let_tv12_ = d_187_v10_
        pat_let_tv13_ = d_187_v10_
        pat_let_tv14_ = d_315_v112_
        pat_let_tv15_ = d_187_v10_
        pat_let_tv16_ = d_187_v10_
        def lambda12_(source4_):
            if source4_.is_DC1:
                d_358___mcc_h1_ = source4_.cf1
                d_359_cf1_ = d_358___mcc_h1_
                return pat_let_tv11_
            elif source4_.is_DC2:
                d_360___mcc_h2_ = source4_.cf2
                d_361___mcc_h3_ = source4_.cf3
                d_362___mcc_h4_ = source4_.cf4
                d_363___mcc_h5_ = source4_.cf5
                d_364___mcc_h6_ = source4_.cf6
                d_365_cf6_ = d_364___mcc_h6_
                d_366_cf5_ = d_363___mcc_h5_
                d_367_cf4_ = d_362___mcc_h4_
                d_368_cf3_ = d_361___mcc_h3_
                d_369_cf2_ = d_360___mcc_h2_
                return pat_let_tv12_
            elif source4_.is_DC3:
                d_370___mcc_h7_ = source4_.cf7
                d_371_cf7_ = d_370___mcc_h7_
                return pat_let_tv13_
            elif True:
                d_372___mcc_h8_ = source4_.cf0
                d_373_cf0_ = d_372___mcc_h8_
                if pat_let_tv14_:
                    return pat_let_tv15_
                elif True:
                    return pat_let_tv16_

        source3_ = lambda12_(d_357_v152_)
        d_374___mcc_h0_ = source3_
        d_375_cf8_ = d_374___mcc_h0_
        d_376_v153_: _dafny.Map
        d_376_v153_ = _dafny.Map({_dafny.CodePoint('i'): d_315_v112_})
        d_376_v153_ = (d_376_v153_).set(d_375_cf8_, not (True) or (d_176_v0_))
        d_377_v154_: _dafny.Set
        d_377_v154_ = _dafny.Set({(D2_DC7(d_337_v134_)).cf11})
        d_378_v155_: C1
        nw53_ = C1()
        nw53_.ctor__(d_377_v154_, d_177_v1_)
        d_378_v155_ = nw53_
        d_379_v156_: int
        d_380_v157_: bool
        out16_: int
        out17_: bool
        out16_, out17_ = default__.m1(d_315_v112_, d_314_v111_, d_190_globalState_)
        d_379_v156_ = out16_
        d_380_v157_ = out17_
        (d_190_globalState_).f5 = d_336_v133_
        if True:
            d_381_v158_: D2
            d_381_v158_ = D2_DC5(d_183_v5_)
            d_382_v159_: _dafny.Seq
            d_382_v159_ = _dafny.SeqWithoutIsStrInference([d_183_v5_, d_183_v5_, d_183_v5_])
            d_383_v160_: _dafny.Array
            nw54_ = _dafny.Array(None, 12)
            nw54_[int(0)] = d_183_v5_
            nw54_[int(1)] = d_183_v5_
            nw54_[int(2)] = d_183_v5_
            nw54_[int(3)] = _dafny.SeqWithoutIsStrInference([d_187_v10_ for d_384_i15_ in range(default__.abs(929))])
            nw54_[int(4)] = (d_183_v5_) + (d_183_v5_)
            nw54_[int(5)] = d_183_v5_
            nw54_[int(6)] = d_183_v5_
            nw54_[int(7)] = d_183_v5_
            nw54_[int(8)] = (((d_381_v158_).cf9) + (_dafny.SeqWithoutIsStrInference([d_187_v10_ for d_385_i16_ in range(default__.abs(879))]))).set(default__.safeIndex(d_336_v133_, len(((d_381_v158_).cf9) + (_dafny.SeqWithoutIsStrInference([d_187_v10_ for d_386_i16_ in range(default__.abs(879))])))), _dafny.CodePoint('w'))
            nw54_[int(9)] = (d_183_v5_) + (d_183_v5_)
            nw54_[int(10)] = (((d_382_v159_)[default__.safeIndex(d_314_v111_, len(d_382_v159_))]).set(default__.safeIndex(len(d_178_v2_), len((d_382_v159_)[default__.safeIndex(d_314_v111_, len(d_382_v159_))])), d_187_v10_)).set(default__.safeIndex(d_177_v1_, len(((d_382_v159_)[default__.safeIndex(d_314_v111_, len(d_382_v159_))]).set(default__.safeIndex(len(d_178_v2_), len((d_382_v159_)[default__.safeIndex(d_314_v111_, len(d_382_v159_))])), d_187_v10_))), d_187_v10_)
            nw54_[int(11)] = (d_183_v5_) + (d_183_v5_)
            d_383_v160_ = nw54_
            d_383_v160_ = d_383_v160_
            d_387_v161_: _dafny.Array
            nw55_ = _dafny.Array(None, 23)
            d_387_v161_ = nw55_
            index50_ = default__.safeIndex(184, (d_186_v9_).length(0))
            (d_186_v9_)[index50_] = d_337_v134_
            index51_ = default__.safeIndex(184, (d_186_v9_).length(0))
            rhs69_ = d_387_v161_
            rhs70_ = True
            rhs71_ = (d_183_v5_).set(default__.safeIndex(d_336_v133_, len(d_183_v5_)), d_187_v10_)
            rhs72_ = ((d_183_v5_) + (d_183_v5_)) + (d_183_v5_)
            lhs48_ = d_186_v9_
            lhs49_ = default__.safeIndex(184, (d_186_v9_).length(0))
            d_387_v161_ = rhs69_
            lhs48_[lhs49_] = rhs70_
            d_183_v5_ = rhs71_
            d_183_v5_ = rhs72_
            d_388_v162_: _dafny.Array
            def lambda13_(d_389_v2_):
                def lambda14_(d_390_i17_):
                    return (d_389_v2_) + (d_389_v2_)

                return lambda14_

            init6_ = lambda13_(d_178_v2_)
            nw56_ = _dafny.Array(None, 20)
            for i0_6_ in range(nw56_.length(0)):
                nw56_[i0_6_] = init6_(i0_6_)
            d_388_v162_ = nw56_
            index52_ = default__.safeIndex(630, (d_388_v162_).length(0))
            (d_388_v162_)[index52_] = _dafny.SeqWithoutIsStrInference([(0) - (d_177_v1_) for d_391_i18_ in range(default__.abs(70))])
            d_392_v163_: _dafny.Set
            d_392_v163_ = _dafny.Set({d_337_v134_, d_337_v134_})
            d_393_v164_: C1
            nw57_ = C1()
            nw57_.ctor__(d_392_v163_, d_177_v1_)
            d_393_v164_ = nw57_
            d_394_v165_: _dafny.Map
            d_394_v165_ = _dafny.Map({d_314_v111_: d_393_v164_})
            d_395_v166_: _dafny.Seq
            d_395_v166_ = _dafny.SeqWithoutIsStrInference([d_394_v165_])
            index53_ = default__.safeIndex(630, (d_388_v162_).length(0))
            (d_388_v162_)[index53_] = _dafny.SeqWithoutIsStrInference([d_314_v111_, len(d_395_v166_)])
            d_393_v164_ = d_393_v164_
            if ((d_183_v5_) + (d_183_v5_)) < (d_183_v5_):
                d_396_v167_: int
                d_397_v168_: bool
                out18_: int
                out19_: bool
                out18_, out19_ = default__.m1((d_337_v134_) or (d_176_v0_), d_177_v1_, d_190_globalState_)
                d_396_v167_ = out18_
                d_397_v168_ = out19_
                d_337_v134_ = (((d_388_v162_)[default__.safeIndex(630, (d_388_v162_).length(0))]) + (d_178_v2_)) == (((_dafny.SeqWithoutIsStrInference([d_396_v167_, 270])) + (d_178_v2_)).set(default__.safeIndex(len(d_335_v132_), len((_dafny.SeqWithoutIsStrInference([d_396_v167_, 270])) + (d_178_v2_))), d_314_v111_))
                d_398_v169_: T0
                nw58_ = C0()
                nw58_.ctor__()
                d_398_v169_ = nw58_
                d_399_v170_: _dafny.Seq
                d_399_v170_ = _dafny.SeqWithoutIsStrInference([d_398_v169_])
                d_399_v170_ = d_399_v170_
                (d_190_globalState_).f19 = (0) - ((d_177_v1_) - (254))
                (d_190_globalState_).f25 = d_176_v0_
            elif True:
                d_400_v171_: _dafny.Map
                d_400_v171_ = _dafny.Map({d_176_v0_: (d_335_v132_).set(default__.safeIndex(115, len(d_335_v132_)), d_315_v112_)})
                d_401_v172_: _dafny.Map
                d_401_v172_ = _dafny.Map({len(d_178_v2_): (d_400_v171_) | (d_400_v171_)})
                d_401_v172_ = (d_401_v172_).set(len((_dafny.Map({d_314_v111_: not((d_186_v9_)[default__.safeIndex(184, (d_186_v9_).length(0))])})).set(d_177_v1_, d_337_v134_)), _dafny.Map({d_337_v134_: d_335_v132_}))
                index54_ = default__.safeIndex(184, (d_186_v9_).length(0))
                (d_186_v9_)[index54_] = not((d_393_v164_.f27) <= (d_336_v133_))
                d_402_v173_: _dafny.Array
                nw59_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_402_v173_ = nw59_
                d_402_v173_ = d_402_v173_
                d_403_v174_: C0
                nw60_ = C0()
                nw60_.ctor__()
                d_403_v174_ = nw60_
                d_404_v175_: _dafny.Seq
                d_404_v175_ = _dafny.SeqWithoutIsStrInference([d_403_v174_])
                d_405_v176_: _dafny.Array
                nw61_ = _dafny.Array(None, 25)
                nw61_[int(0)] = d_403_v174_
                nw61_[int(1)] = d_403_v174_
                nw61_[int(2)] = d_403_v174_
                nw61_[int(3)] = d_403_v174_
                nw61_[int(4)] = d_403_v174_
                nw61_[int(5)] = d_403_v174_
                nw61_[int(6)] = d_403_v174_
                nw61_[int(7)] = d_403_v174_
                nw61_[int(8)] = d_403_v174_
                nw61_[int(9)] = d_403_v174_
                nw61_[int(10)] = d_403_v174_
                nw61_[int(11)] = d_403_v174_
                nw61_[int(12)] = d_403_v174_
                nw61_[int(13)] = d_403_v174_
                nw61_[int(14)] = d_403_v174_
                nw61_[int(15)] = d_403_v174_
                nw61_[int(16)] = (d_403_v174_ if d_337_v134_ else d_403_v174_)
                nw61_[int(17)] = d_403_v174_
                nw61_[int(18)] = (d_404_v175_)[default__.safeIndex(len(d_335_v132_), len(d_404_v175_))]
                nw61_[int(19)] = d_403_v174_
                nw61_[int(20)] = d_403_v174_
                nw61_[int(21)] = d_403_v174_
                nw61_[int(22)] = d_403_v174_
                nw61_[int(23)] = d_403_v174_
                nw61_[int(24)] = d_403_v174_
                d_405_v176_ = nw61_
                index55_ = default__.safeIndex(123, (d_405_v176_).length(0))
                (d_405_v176_)[index55_] = d_403_v174_
                index56_ = default__.safeIndex(123, (d_405_v176_).length(0))
                rhs73_ = d_314_v111_
                rhs74_ = (d_336_v133_) * (len(_dafny.SeqWithoutIsStrInference([len(d_404_v175_), len(d_178_v2_), -511, d_393_v164_.f27, d_314_v111_])))
                rhs75_ = _dafny.SeqWithoutIsStrInference([d_187_v10_ for d_406_i19_ in range(default__.abs(563))])
                rhs76_ = d_403_v174_
                rhs77_ = d_337_v134_
                lhs50_ = d_190_globalState_
                lhs51_ = d_393_v164_
                lhs52_ = d_405_v176_
                lhs53_ = default__.safeIndex(123, (d_405_v176_).length(0))
                lhs54_ = d_190_globalState_
                lhs50_.f8 = rhs73_
                lhs51_.f27 = rhs74_
                d_183_v5_ = rhs75_
                lhs52_[lhs53_] = rhs76_
                lhs54_.f25 = rhs77_
                d_407_v177_: int
                d_408_v178_: bool
                out20_: int
                out21_: bool
                out20_, out21_ = default__.m1(not((d_185_v8_).ispropersubset(d_185_v8_)), (d_177_v1_) + (d_393_v164_.f27), d_190_globalState_)
                d_407_v177_ = out20_
                d_408_v178_ = out21_
        elif True:
            d_409_v179_: _dafny.Set
            d_409_v179_ = _dafny.Set({not(True), d_176_v0_, d_176_v0_})
            d_410_v180_: C1
            nw62_ = C1()
            nw62_.ctor__(d_409_v179_, d_314_v111_)
            d_410_v180_ = nw62_
            d_411_v181_: _dafny.MultiSet
            d_411_v181_ = _dafny.MultiSet([d_410_v180_, d_410_v180_])
            if ((d_411_v181_) - (d_411_v181_)).isdisjoint(d_411_v181_):
                d_412_v182_: _dafny.Seq
                d_412_v182_ = _dafny.SeqWithoutIsStrInference([(d_338_v135_).intersection(d_338_v135_)])
                d_412_v182_ = d_412_v182_
                d_413_v183_: C0
                nw63_ = C0()
                nw63_.ctor__()
                d_413_v183_ = nw63_
                d_186_v9_ = d_186_v9_
                d_414_v184_: _dafny.Map
                d_414_v184_ = _dafny.Map({(d_183_v5_).set(default__.safeIndex(d_410_v180_.f27, len(d_183_v5_)), d_187_v10_): d_176_v0_})
                (d_190_globalState_).f22 = ((d_414_v184_)[d_183_v5_] if (d_183_v5_) in (d_414_v184_) else d_315_v112_)
                d_415_v185_: C0
                nw64_ = C0()
                nw64_.ctor__()
                d_415_v185_ = nw64_
            elif True:
                d_416_v186_: _dafny.Map
                d_416_v186_ = _dafny.Map({default__.safeDivisionInt(625, ((d_338_v135_)[len((d_410_v180_).f26)] if (len((d_410_v180_).f26)) in (d_338_v135_) else d_314_v111_)): (d_183_v5_).set(default__.safeIndex(d_314_v111_, len(d_183_v5_)), d_187_v10_)})
                d_416_v186_ = d_416_v186_
                d_417_v187_: _dafny.Array
                def lambda15_(d_418_v8_):
                    def lambda16_(d_419_i20_):
                        return d_418_v8_

                    return lambda16_

                init7_ = lambda15_(d_185_v8_)
                nw65_ = _dafny.Array(None, 8)
                for i0_7_ in range(nw65_.length(0)):
                    nw65_[i0_7_] = init7_(i0_7_)
                d_417_v187_ = nw65_
                d_420_v188_: _dafny.Map
                d_420_v188_ = _dafny.Map({d_417_v187_: d_410_v180_.f27})
                d_420_v188_ = (d_420_v188_).set(d_417_v187_, (d_314_v111_) + (d_314_v111_))
                d_183_v5_ = ((d_183_v5_) + (d_183_v5_)) + (d_183_v5_)
                d_421_v189_: D2
                d_421_v189_ = D2_DC6((d_410_v180_.f27) == (d_177_v1_))
                rhs78_ = d_176_v0_
                rhs79_ = default__.fm17(d_190_globalState_)
                lhs55_ = d_190_globalState_
                lhs55_.f23 = rhs78_
                d_421_v189_ = rhs79_
                index57_ = default__.safeIndex(742, (d_180_v4_).length(0))
                (d_180_v4_)[index57_] = d_336_v133_
                index58_ = default__.safeIndex(742, (d_180_v4_).length(0))
                (d_180_v4_)[index58_] = (0) - (default__.fm1(True, d_187_v10_, d_336_v133_, d_190_globalState_))
            rhs80_ = d_177_v1_
            rhs81_ = (d_314_v111_) * (d_314_v111_)
            lhs56_ = d_190_globalState_
            lhs57_ = d_410_v180_
            lhs56_.f8 = rhs80_
            lhs57_.f27 = rhs81_
            d_422_v190_: _dafny.Map
            d_422_v190_ = _dafny.Map({_dafny.CodePoint('o'): (d_410_v180_.f27) + (d_336_v133_)})
            d_423_v191_: _dafny.Map
            d_423_v191_ = _dafny.Map({not(d_176_v0_): d_314_v111_})
            d_424_v192_: _dafny.Seq
            d_424_v192_ = _dafny.SeqWithoutIsStrInference([d_410_v180_, d_410_v180_])
            d_422_v190_ = default__.fm18(((d_423_v191_)[d_176_v0_] if (d_176_v0_) in (d_423_v191_) else d_410_v180_.f27), (_dafny.SeqWithoutIsStrInference([d_410_v180_])) < (d_424_v192_), d_190_globalState_)
            d_425_v193_: D0
            d_425_v193_ = D0_DC1(len(d_195_v14_))
            if not (not((len(default__.fm12(len(d_183_v5_), d_425_v193_, d_190_globalState_))) >= (len(d_196_v15_)))) or (d_337_v134_):
                d_426_v194_: D2
                d_426_v194_ = D2_DC5(d_183_v5_)
                d_427_v195_: D4
                d_427_v195_ = D4_DC14(d_338_v135_, d_410_v180_.f27, d_426_v194_)
                (d_190_globalState_).f23 = (d_314_v111_) == ((0) - (default__.safeModuloInt(605, (d_427_v195_).cf29)))
                d_428_v196_: int
                d_429_v197_: bool
                out22_: int
                out23_: bool
                out22_, out23_ = default__.m1(d_337_v134_, d_410_v180_.f27, d_190_globalState_)
                d_428_v196_ = out22_
                d_429_v197_ = out23_
                d_430_v198_: T0
                nw66_ = C0()
                nw66_.ctor__()
                d_430_v198_ = nw66_
                rhs82_ = (0) - ((0) - (default__.safeDivisionInt(((d_338_v135_)[d_336_v133_] if (d_336_v133_) in (d_338_v135_) else d_410_v180_.f27), d_314_v111_)))
                rhs83_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfeampgsi"))) + (d_183_v5_)) + ((_dafny.SeqWithoutIsStrInference([d_187_v10_ for d_431_i21_ in range(default__.abs(573))])) + ((d_426_v194_).cf9))
                rhs84_ = (d_177_v1_) + (-503)
                rhs85_ = d_430_v198_
                rhs86_ = d_430_v198_
                lhs58_ = d_190_globalState_
                d_177_v1_ = rhs82_
                d_183_v5_ = rhs83_
                lhs58_.f19 = rhs84_
                d_430_v198_ = rhs85_
                d_430_v198_ = rhs86_
                (d_190_globalState_).f22 = (d_410_v180_).fm6(_dafny.SeqWithoutIsStrInference([d_357_v152_]), (d_183_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))), d_185_v8_, d_190_globalState_)
                index59_ = default__.safeIndex(941, (d_186_v9_).length(0))
                (d_186_v9_)[index59_] = not((d_187_v10_) not in (d_183_v5_))
                d_432_v199_: _dafny.Map
                d_432_v199_ = _dafny.Map({d_338_v135_: d_178_v2_})
                d_433_v200_: D2
                d_433_v200_ = D2_DC8(d_314_v111_, d_176_v0_, d_336_v133_)
                index60_ = default__.safeIndex(941, (d_186_v9_).length(0))
                rhs87_ = d_187_v10_
                rhs88_ = d_314_v111_
                rhs89_ = len((d_432_v199_ if d_315_v112_ else d_432_v199_))
                rhs90_ = (d_433_v200_).cf13
                rhs91_ = d_177_v1_
                lhs59_ = d_190_globalState_
                lhs60_ = d_186_v9_
                lhs61_ = default__.safeIndex(941, (d_186_v9_).length(0))
                lhs62_ = d_190_globalState_
                d_187_v10_ = rhs87_
                d_314_v111_ = rhs88_
                lhs59_.f2 = rhs89_
                lhs60_[lhs61_] = rhs90_
                lhs62_.f10 = rhs91_
            elif True:
                rhs92_ = d_410_v180_
                rhs93_ = d_331_v128_
                d_410_v180_ = rhs92_
                d_331_v128_ = rhs93_
                d_434_v201_: D2
                d_434_v201_ = D2_DC8(d_336_v133_, d_315_v112_, d_410_v180_.f27)
                d_435_v202_: int
                d_436_v203_: bool
                out24_: int
                out25_: bool
                out24_, out25_ = default__.m1(False, (d_434_v201_).cf12, d_190_globalState_)
                d_435_v202_ = out24_
                d_436_v203_ = out25_
                (d_190_globalState_).f2 = len(d_183_v5_)
                index61_ = default__.safeIndex(60, (d_186_v9_).length(0))
                (d_186_v9_)[index61_] = d_176_v0_
                index62_ = default__.safeIndex(60, (d_186_v9_).length(0))
                (d_186_v9_)[index62_] = d_436_v203_
                d_195_v14_ = (d_195_v14_) | (d_195_v14_)
            d_437_v204_: _dafny.MultiSet
            d_437_v204_ = _dafny.MultiSet([d_357_v152_])
            d_438_v205_: _dafny.Map
            d_438_v205_ = _dafny.Map({d_410_v180_: d_437_v204_})
            d_439_v206_: _dafny.Seq
            d_439_v206_ = _dafny.SeqWithoutIsStrInference([((_dafny.MultiSet(d_178_v2_)).set(d_410_v180_.f27, default__.abs(d_314_v111_))).set(d_410_v180_.f27, default__.abs(154)), d_338_v135_])
            pat_let_tv17_ = d_439_v206_
            d_440_v207_: _dafny.MultiSet
            def iife25_(_pat_let8_0):
                def iife26_(d_441_dt__update__tmp_h5_):
                    def iife27_(_pat_let9_0):
                        def iife28_(d_442_dt__update_hcf0_h0_):
                            return D0_DC0(d_442_dt__update_hcf0_h0_)
                        return iife28_(_pat_let9_0)
                    return iife27_(pat_let_tv17_)
                return iife26_(_pat_let8_0)
            d_440_v207_ = _dafny.MultiSet([iife25_(d_357_v152_), d_357_v152_])
            d_443_v208_: _dafny.Seq
            d_443_v208_ = _dafny.SeqWithoutIsStrInference([default__.fm19(d_190_globalState_), (d_440_v207_)])
            d_438_v205_ = (d_438_v205_).set(d_410_v180_, (d_443_v208_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbifwaa"))), len(d_443_v208_))])
        d_444_v209_: T0
        nw67_ = C0()
        nw67_.ctor__()
        d_444_v209_ = nw67_
        d_444_v209_ = d_444_v209_
        d_445_v210_: D3
        d_445_v210_ = D3_DC11(_dafny.CodePoint('r'), False, d_176_v0_)
        hi1_ = d_177_v1_
        for d_446_i22_ in range((default__.fm1(not((d_445_v210_).cf20), d_187_v10_, (_dafny.MultiSet([d_336_v133_])).cardinality, d_190_globalState_)) - (d_314_v111_), hi1_):
            d_447_v211_: C0
            nw68_ = C0()
            nw68_.ctor__()
            d_447_v211_ = nw68_
            d_447_v211_ = d_447_v211_
            d_448_v212_: int
            d_449_v213_: bool
            out26_: int
            out27_: bool
            out26_, out27_ = default__.m1(d_176_v0_, (d_334_v131_).cf17, d_190_globalState_)
            d_448_v212_ = out26_
            d_449_v213_ = out27_
            d_450_v214_: _dafny.Set
            d_450_v214_ = _dafny.Set({d_449_v213_})
            d_451_v215_: C1
            nw69_ = C1()
            nw69_.ctor__(d_450_v214_, len(d_183_v5_))
            d_451_v215_ = nw69_
            d_336_v133_ = (0) - ((578) + (d_177_v1_))
        _dafny.print(_dafny.string_of(d_176_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_177_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_v2_) == (_dafny.SeqWithoutIsStrInference([17]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v3_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([17])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v4_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v4_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v4_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v4_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v4_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v4_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v4_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v4_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_183_v5_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v6_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cb"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_v8_) == (_dafny.MultiSet([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_187_v10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v11_) == (_dafny.MultiSet([_dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('f')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_189_v12_) == (_dafny.Map({17: _dafny.MultiSet([_dafny.CodePoint('u')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f0) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([17])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f6) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f11)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f12) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cb"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_190_globalState_).f13).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_globalState_.f16) == (_dafny.MultiSet([False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f17) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f18)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_190_globalState_).f20).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_190_globalState_).f21) == (_dafny.Map({17: _dafny.MultiSet([_dafny.CodePoint('u')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_190_globalState_).f24).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_190_globalState_.f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_194_v13_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v14_) == (_dafny.Map({-17: True, 438: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v15_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v90_) == (_dafny.Set({17}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_314_v111_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_315_v112_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_331_v128_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False, False, False, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_332_v129_) == (_dafny.Set({_dafny.MultiSet([True, True, True]), _dafny.MultiSet([False, False, False, False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v130_) == (_dafny.Map({_dafny.MultiSet([True, True, True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_v131_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_v131_).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_v131_).cf18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_v132_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_336_v133_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_337_v134_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_338_v135_) == (_dafny.MultiSet([60]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v136_) == (_dafny.Map({_dafny.MultiSet([60]): 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v152_).cf0) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([]), _dafny.MultiSet([6, 119, -456]), _dafny.MultiSet([9]), _dafny.MultiSet([508, -985])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_445_v210_).cf19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_445_v210_).cf20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_445_v210_).cf21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf7 == __o.cf7
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
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({self.cf9.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf22', Any), ('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(_dafny.MultiSet({}), int(0), D2.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC14(D4, NamedTuple('DC14', [('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC15(D5, NamedTuple('DC15', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)

class D6_DC17(D6, NamedTuple('DC17', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC20(D7, NamedTuple('DC20', [('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf35)}, {self.cf36.VerbatimString(True)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC22(D8, NamedTuple('DC22', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self.f5: int = int(0)
        self.f8: int = int(0)
        self.f10: int = int(0)
        self.f11: _dafny.Array = _dafny.Array(None, 0)
        self.f15: int = int(0)
        self.f16: _dafny.MultiSet = _dafny.MultiSet({})
        self.f19: int = int(0)
        self.f22: bool = False
        self.f23: bool = False
        self.f25: bool = False
        self._f0: _dafny.Map = _dafny.Map({})
        self._f1: int = int(0)
        self._f3: int = int(0)
        self._f4: bool = False
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f7: int = int(0)
        self._f9: bool = False
        self._f12: _dafny.Set = _dafny.Set({})
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f14: bool = False
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f18: _dafny.Array = _dafny.Array(None, 0)
        self._f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f21: _dafny.Map = _dafny.Map({})
        self._f24: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f22 = f22
        (self).f23 = f23
        (self)._f24 = f24
        (self).f25 = f25

    @property
    def f0(self):
        return self._f0
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
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f9(self):
        return self._f9
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    @property
    def f24(self):
        return self._f24

class C0(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm4(self, p0, p1, globalState):
        return (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_452_i0_ in range(default__.abs(152))])), 860)) - (661)

    def fm7(self, p0, p1, p2, globalState):
        def iife29_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in ((_dafny.SeqWithoutIsStrInference([-108])) + (_dafny.SeqWithoutIsStrInference([(D0_DC1(160)).cf1, 51, len(_dafny.Map({True: 574})), len(_dafny.Map({True: False}))]))).Elements:
                d_453_v0_: int = compr_9_
                if (d_453_v0_) in ((_dafny.SeqWithoutIsStrInference([-108])) + (_dafny.SeqWithoutIsStrInference([(D0_DC1(160)).cf1, 51, len(_dafny.Map({True: 574})), len(_dafny.Map({True: False}))]))):
                    coll9_[(d_453_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhtgllpi"))))] = not(((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_454_i0_ in range(default__.abs(742))])))) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sntppqvt")))))
            return _dafny.Map(coll9_)
        return iife29_()
        

    def fm8(self, p0, globalState):
        return (len((_dafny.SeqWithoutIsStrInference([-934, 520]) if True else _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "papx")))): False}))])))) + ((0) - (((0) - (len(_dafny.SeqWithoutIsStrInference([D0_DC1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ua")))) for d_455_i0_ in range(default__.abs(783))])))) - (826)))


class C1(T0):
    def  __init__(self):
        self.f27: int = int(0)
        self._f26: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f26, f27):
        (self)._f26 = f26
        (self).f27 = f27

    def fm4(self, p0, p1, globalState):
        return self.f27

    def fm5(self, globalState):
        return (((_dafny.CodePoint('s')) if True else _dafny.CodePoint('h')))

    def fm6(self, p0, p1, p2, globalState):
        return ((self.f27) + (self.f27)) <= (self.f27)

    def m0(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_456_v0_: str
        d_456_v0_ = _dafny.CodePoint('x')
        d_457_v1_: _dafny.Map
        d_457_v1_ = _dafny.Map({self.f27: d_456_v0_})
        d_457_v1_ = (d_457_v1_).set(default__.fm1(True, d_456_v0_, self.f27, globalState), d_456_v0_)
        d_458_v2_: _dafny.Array
        nw70_ = _dafny.Array(int(0), 10)
        d_458_v2_ = nw70_
        d_459_v3_: _dafny.Set
        d_459_v3_ = _dafny.Set({d_458_v2_, d_458_v2_})
        if ((d_459_v3_).intersection(d_459_v3_)).issubset(d_459_v3_):
            d_460_v4_: _dafny.MultiSet
            d_460_v4_ = _dafny.MultiSet([860])
            d_461_v5_: _dafny.Seq
            d_461_v5_ = _dafny.SeqWithoutIsStrInference([d_460_v4_, d_460_v4_])
            d_462_v6_: D0
            d_462_v6_ = D0_DC0(d_461_v5_)
            d_463_v7_: _dafny.Seq
            d_463_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypmgf"))
            d_464_v8_: _dafny.MultiSet
            d_464_v8_ = _dafny.MultiSet([p0, p0])
            d_465_v9_: _dafny.Map
            d_465_v9_ = _dafny.Map({p0: p0})
            d_466_v10_: _dafny.Map
            d_466_v10_ = _dafny.Map({(self).fm4(self.f27, default__.fm0((self).fm6(_dafny.SeqWithoutIsStrInference([d_462_v6_]), d_463_v7_, d_464_v8_, globalState), d_465_v9_, globalState), globalState): p0})
            index63_ = default__.safeIndex(892, (d_458_v2_).length(0))
            (d_458_v2_)[index63_] = self.f27
            index64_ = default__.safeIndex(892, (d_458_v2_).length(0))
            rhs94_ = not ((931) not in (_dafny.Set({self.f27}))) or (p0)
            rhs95_ = (self.f27 if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f27, self.f27, (d_464_v8_).cardinality]))) == (d_460_v4_) else self.f27)
            rhs96_ = _dafny.Map({self.f27: p0})
            rhs97_ = -345
            lhs63_ = globalState
            lhs64_ = globalState
            lhs65_ = d_458_v2_
            lhs66_ = default__.safeIndex(892, (d_458_v2_).length(0))
            lhs63_.f23 = rhs94_
            lhs64_.f19 = rhs95_
            d_466_v10_ = rhs96_
            lhs65_[lhs66_] = rhs97_
            d_467_v11_: _dafny.Array
            nw71_ = _dafny.Array(None, 5)
            nw71_[int(0)] = d_463_v7_
            nw71_[int(1)] = d_463_v7_
            nw71_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtybqap"))).set(default__.safeIndex(self.f27, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtybqap")))), d_456_v0_)
            nw71_[int(3)] = (d_463_v7_) + ((d_463_v7_).set(default__.safeIndex((d_458_v2_)[default__.safeIndex(892, (d_458_v2_).length(0))], len(d_463_v7_)), d_456_v0_))
            nw71_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmqrsdga"))
            d_467_v11_ = nw71_
            d_467_v11_ = d_467_v11_
            d_463_v7_ = d_463_v7_
            (globalState).f23 = p0
            d_468_v12_: _dafny.Map
            d_468_v12_ = _dafny.Map({self.f27: len(d_463_v7_)})
            d_468_v12_ = (d_468_v12_).set(len(_dafny.SeqWithoutIsStrInference([d_456_v0_ for d_469_i0_ in range(default__.abs(154))])), len((self).f26))
        elif True:
            d_470_v13_: D0
            d_470_v13_ = D0_DC2(p0, d_459_v3_, p0, p0, self.f27)
            (globalState).f22 = (len(_dafny.Map({(d_470_v13_).cf6: self.f27}))) == ((0) - (self.f27))
            d_471_v14_: _dafny.Array
            nw72_ = _dafny.Array(False, 20)
            d_471_v14_ = nw72_
            index65_ = default__.safeIndex(990, (d_471_v14_).length(0))
            (d_471_v14_)[index65_] = p0
            d_472_v15_: _dafny.Seq
            d_472_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "np"))
            index66_ = default__.safeIndex(990, (d_471_v14_).length(0))
            rhs98_ = d_471_v14_
            rhs99_ = (len(d_472_v15_)) >= (len(d_459_v3_))
            lhs67_ = d_471_v14_
            lhs68_ = default__.safeIndex(990, (d_471_v14_).length(0))
            d_471_v14_ = rhs98_
            lhs67_[lhs68_] = rhs99_
            (self).f27 = default__.safeDivisionInt((0) - (default__.fm1((d_471_v14_)[default__.safeIndex(990, (d_471_v14_).length(0))], d_456_v0_, self.f27, globalState)), self.f27)
            d_473_v16_: _dafny.Array
            nw73_ = _dafny.Array(None, 13)
            nw73_[int(0)] = d_456_v0_
            nw73_[int(1)] = d_456_v0_
            nw73_[int(2)] = d_456_v0_
            nw73_[int(3)] = d_456_v0_
            nw73_[int(4)] = (self).fm5(globalState)
            nw73_[int(5)] = d_456_v0_
            nw73_[int(6)] = d_456_v0_
            nw73_[int(7)] = d_456_v0_
            nw73_[int(8)] = d_456_v0_
            nw73_[int(9)] = d_456_v0_
            nw73_[int(10)] = d_456_v0_
            nw73_[int(11)] = d_456_v0_
            nw73_[int(12)] = d_456_v0_
            d_473_v16_ = nw73_
            index67_ = default__.safeIndex(866, (d_473_v16_).length(0))
            (d_473_v16_)[index67_] = d_456_v0_
            index68_ = default__.safeIndex(866, (d_473_v16_).length(0))
            (d_473_v16_)[index68_] = (_dafny.CodePoint('l') if p0 else d_456_v0_)
            d_474_v17_: _dafny.MultiSet
            d_474_v17_ = _dafny.MultiSet([not(p0), (d_471_v14_)[default__.safeIndex(990, (d_471_v14_).length(0))], p0])
            (globalState).f15 = (self.f27 if (True) or (p0) else ((d_474_v17_)[not((d_471_v14_)[default__.safeIndex(990, (d_471_v14_).length(0))])] if (not((d_471_v14_)[default__.safeIndex(990, (d_471_v14_).length(0))])) in (d_474_v17_) else len((d_472_v15_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sc"))), len(d_472_v15_)), d_456_v0_))))
        if p0:
            d_475_v18_: _dafny.Map
            d_475_v18_ = _dafny.Map({p0: _dafny.Set({_dafny.SeqWithoutIsStrInference([(0) - ((0) - (self.f27)) for d_476_i1_ in range(default__.abs(-390))])})})
            d_477_v19_: _dafny.Seq
            d_477_v19_ = _dafny.SeqWithoutIsStrInference([48])
            d_478_v20_: _dafny.Set
            d_478_v20_ = _dafny.Set({d_477_v19_})
            d_475_v18_ = (d_475_v18_).set(p0, d_478_v20_)
            d_479_v21_: _dafny.Seq
            d_479_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uflnveram"))
            (globalState).f8 = (self).fm4(len(d_479_v21_), p0, globalState)
            d_480_v22_: _dafny.MultiSet
            d_480_v22_ = _dafny.MultiSet([self.f27])
            d_481_v23_: _dafny.MultiSet
            d_481_v23_ = _dafny.MultiSet([((d_480_v22_)[self.f27] if (self.f27) in (d_480_v22_) else len(d_479_v21_))])
            d_482_v24_: _dafny.Seq
            d_482_v24_ = _dafny.SeqWithoutIsStrInference([d_481_v23_])
            d_483_v25_: D0
            d_483_v25_ = D0_DC0((d_482_v24_) + (_dafny.SeqWithoutIsStrInference([(d_482_v24_)[default__.safeIndex(self.f27, len(d_482_v24_))]])))
            source5_ = d_483_v25_
            if source5_.is_DC1:
                d_484___mcc_h0_ = source5_.cf1
                d_485_cf1_ = d_484___mcc_h0_
                d_486_v26_: _dafny.Array
                nw74_ = _dafny.Array(False, 1)
                d_486_v26_ = nw74_
                index69_ = default__.safeIndex(126, (d_486_v26_).length(0))
                (d_486_v26_)[index69_] = p0
                index70_ = default__.safeIndex(126, (d_486_v26_).length(0))
                (d_486_v26_)[index70_] = p0
                index71_ = default__.safeIndex(126, (d_486_v26_).length(0))
                rhs100_ = not((d_486_v26_)[default__.safeIndex(126, (d_486_v26_).length(0))])
                rhs101_ = (d_486_v26_)[default__.safeIndex(126, (d_486_v26_).length(0))]
                lhs69_ = d_486_v26_
                lhs70_ = default__.safeIndex(126, (d_486_v26_).length(0))
                lhs71_ = globalState
                lhs69_[lhs70_] = rhs100_
                lhs71_.f25 = rhs101_
                index72_ = default__.safeIndex(126, (d_486_v26_).length(0))
                (d_486_v26_)[index72_] = ((d_486_v26_)[default__.safeIndex(126, (d_486_v26_).length(0))] if (p0) and ((d_486_v26_)[default__.safeIndex(126, (d_486_v26_).length(0))]) else p0)
                d_487_v27_: T0
                nw75_ = C0()
                nw75_.ctor__()
                d_487_v27_ = nw75_
                d_487_v27_ = d_487_v27_
            elif source5_.is_DC2:
                d_488___mcc_h1_ = source5_.cf2
                d_489___mcc_h2_ = source5_.cf3
                d_490___mcc_h3_ = source5_.cf4
                d_491___mcc_h4_ = source5_.cf5
                d_492___mcc_h5_ = source5_.cf6
                d_493_cf6_ = d_492___mcc_h5_
                d_494_cf5_ = d_491___mcc_h4_
                d_495_cf4_ = d_490___mcc_h3_
                d_496_cf3_ = d_489___mcc_h2_
                d_497_cf2_ = d_488___mcc_h1_
                d_498_v28_: _dafny.Map
                d_498_v28_ = _dafny.Map({p0: d_479_v21_})
                d_499_v29_: _dafny.Map
                d_499_v29_ = _dafny.Map({d_497_cf2_: len(((d_498_v28_)[d_494_cf5_] if (d_494_cf5_) in (d_498_v28_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))))})
                d_500_v30_: _dafny.Map
                d_500_v30_ = _dafny.Map({(0) - ((0) - (len(((_dafny.SeqWithoutIsStrInference([d_456_v0_ for d_501_i2_ in range(default__.abs(879))])).set(default__.safeIndex(len(d_499_v29_), len(_dafny.SeqWithoutIsStrInference([d_456_v0_ for d_502_i2_ in range(default__.abs(879))]))), d_456_v0_)) + (d_479_v21_)))): d_493_cf6_})
                d_500_v30_ = (d_500_v30_).set((611) - (len(d_477_v19_)), -674)
                d_503_v31_: _dafny.Seq
                d_503_v31_ = _dafny.SeqWithoutIsStrInference([d_483_v25_, d_483_v25_])
                d_504_v32_: _dafny.MultiSet
                d_504_v32_ = _dafny.MultiSet([False, d_494_cf5_, False])
                d_505_v33_: _dafny.Map
                d_505_v33_ = _dafny.Map({d_497_cf2_: (self).fm6(d_503_v31_, d_479_v21_, d_504_v32_, globalState)})
                d_506_v34_: _dafny.Seq
                d_506_v34_ = _dafny.SeqWithoutIsStrInference([d_458_v2_])
                (globalState).f2 = default__.safeModuloInt((0) - (len(d_505_v33_)), (self.f27) * (len(d_506_v34_)))
                (globalState).f10 = self.f27
                (globalState).f22 = default__.fm0(not((d_493_cf6_) == (d_493_cf6_)), default__.fm9(836, globalState), globalState)
            elif source5_.is_DC3:
                d_507___mcc_h6_ = source5_.cf7
                d_508_cf7_ = d_507___mcc_h6_
                (globalState).f2 = 744
                (globalState).f2 = self.f27
                d_509_v35_: C0
                nw76_ = C0()
                nw76_.ctor__()
                d_509_v35_ = nw76_
                d_510_v36_: _dafny.Seq
                d_510_v36_ = _dafny.SeqWithoutIsStrInference([d_477_v19_, d_477_v19_, (d_477_v19_ if not(p0) else _dafny.SeqWithoutIsStrInference([self.f27])), d_477_v19_])
                (globalState).f5 = len(d_510_v36_)
            elif True:
                d_511___mcc_h7_ = source5_.cf0
                d_512_cf0_ = d_511___mcc_h7_
                d_513_v37_: C0
                nw77_ = C0()
                nw77_.ctor__()
                d_513_v37_ = nw77_
                d_514_v38_: _dafny.Map
                d_514_v38_ = _dafny.Map({p0: p0})
                d_514_v38_ = (d_514_v38_).set(True, p0)
                d_515_v39_: str
                d_515_v39_ = _dafny.CodePoint('p')
                d_516_v40_: _dafny.Map
                d_516_v40_ = _dafny.Map({p0: d_456_v0_})
                rhs102_ = ((d_516_v40_)[p0] if (p0) in (d_516_v40_) else d_456_v0_)
                rhs103_ = d_480_v22_
                d_515_v39_ = rhs102_
                d_481_v23_ = rhs103_
                d_517_v41_: _dafny.MultiSet
                d_517_v41_ = _dafny.MultiSet([d_456_v0_, d_456_v0_, d_456_v0_, d_456_v0_, d_456_v0_])
                rhs104_ = p0
                rhs105_ = p0
                rhs106_ = d_517_v41_
                lhs72_ = globalState
                lhs73_ = globalState
                lhs72_.f25 = rhs104_
                lhs73_.f25 = rhs105_
                d_517_v41_ = rhs106_
            d_518_v42_: _dafny.Array
            nw78_ = _dafny.Array(False, 8)
            d_518_v42_ = nw78_
            d_519_v43_: _dafny.Set
            d_519_v43_ = _dafny.Set({d_518_v42_})
            (globalState).f23 = ((d_519_v43_) | (d_519_v43_)).isdisjoint(d_519_v43_)
            d_520_v44_: _dafny.MultiSet
            d_520_v44_ = _dafny.MultiSet([p0])
            d_521_v45_: _dafny.Map
            d_521_v45_ = _dafny.Map({p0: ((d_520_v44_).cardinality if p0 else self.f27)})
            d_521_v45_ = (d_521_v45_).set((d_480_v22_).isdisjoint(d_480_v22_), self.f27)
        elif True:
            if p0:
                d_522_v46_: _dafny.Array
                nw79_ = _dafny.Array(None, 4)
                d_522_v46_ = nw79_
                d_523_v47_: _dafny.Map
                d_523_v47_ = _dafny.Map({p0: d_522_v46_})
                d_524_v48_: _dafny.MultiSet
                d_524_v48_ = _dafny.MultiSet([len(d_523_v47_)])
                d_524_v48_ = (d_524_v48_) | (d_524_v48_)
                d_525_v49_: C0
                nw80_ = C0()
                nw80_.ctor__()
                d_525_v49_ = nw80_
                d_526_v50_: _dafny.Seq
                d_526_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                d_526_v50_ = (_dafny.SeqWithoutIsStrInference([d_456_v0_ for d_527_i3_ in range(default__.abs(416))])) + ((D2_DC5(d_526_v50_)).cf9)
                d_528_v51_: C0
                nw81_ = C0()
                nw81_.ctor__()
                d_528_v51_ = nw81_
                d_529_v52_: _dafny.Map
                d_529_v52_ = _dafny.Map({p0: p0})
                d_530_v53_: _dafny.Map
                d_530_v53_ = _dafny.Map({len((d_529_v52_) | (d_529_v52_)): (self.f27) + (self.f27)})
                d_530_v53_ = (d_530_v53_).set(792, self.f27)
            elif True:
                d_531_v54_: _dafny.MultiSet
                d_531_v54_ = _dafny.MultiSet([self.f27, self.f27, self.f27])
                (globalState).f25 = default__.fm0((d_531_v54_).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f27]))), _dafny.Map({p0: p0}), globalState)
                d_532_v55_: _dafny.Array
                nw82_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_532_v55_ = nw82_
                d_533_v56_: _dafny.Array
                nw83_ = _dafny.Array(False, 3)
                d_533_v56_ = nw83_
                index73_ = default__.safeIndex(99, (d_532_v55_).length(0))
                (d_532_v55_)[index73_] = d_533_v56_
                index74_ = default__.safeIndex(99, (d_532_v55_).length(0))
                (d_532_v55_)[index74_] = d_533_v56_
                d_534_v57_: _dafny.MultiSet
                d_534_v57_ = _dafny.MultiSet([p0])
                (globalState).f16 = ((_dafny.MultiSet([False, False])) - (d_534_v57_)) | (d_534_v57_)
                (globalState).f22 = p0
                d_535_v58_: _dafny.Seq
                d_535_v58_ = _dafny.SeqWithoutIsStrInference([p0])
                d_536_v59_: _dafny.Seq
                d_536_v59_ = _dafny.SeqWithoutIsStrInference([self.f27])
                arr0_ = (d_532_v55_)[default__.safeIndex(99, (d_532_v55_).length(0))]
                index75_ = default__.safeIndex(691, ((d_532_v55_)[default__.safeIndex(99, (d_532_v55_).length(0))]).length(0))
                arr0_[index75_] = (d_535_v58_)[default__.safeIndex((d_536_v59_)[default__.safeIndex(self.f27, len(d_536_v59_))], len(d_535_v58_))]
                d_537_v60_: _dafny.Map
                d_537_v60_ = _dafny.Map({p0: d_456_v0_})
                arr1_ = (d_532_v55_)[default__.safeIndex(99, (d_532_v55_).length(0))]
                index76_ = default__.safeIndex(691, ((d_532_v55_)[default__.safeIndex(99, (d_532_v55_).length(0))]).length(0))
                arr1_[index76_] = (d_537_v60_) != ((D3_DC9(d_537_v60_)).cf15)
            d_538_v61_: _dafny.MultiSet
            d_538_v61_ = _dafny.MultiSet([p0, p0, p0, p0])
            d_539_v62_: D4
            d_539_v62_ = D4_DC13(d_538_v61_)
            (globalState).f16 = (d_539_v62_).cf27
            d_540_v63_: _dafny.Seq
            d_540_v63_ = _dafny.SeqWithoutIsStrInference([p0])
            d_541_v64_: _dafny.Seq
            d_541_v64_ = _dafny.SeqWithoutIsStrInference([(d_540_v63_)[default__.safeIndex(self.f27, len(d_540_v63_))], p0, p0])
            d_542_v65_: D2
            d_542_v65_ = D2_DC7(p0)
            d_541_v64_ = (((d_540_v63_).set(default__.safeIndex(default__.fm1(p0, d_456_v0_, self.f27, globalState), len(d_540_v63_)), False)).set(default__.safeIndex(self.f27, len((d_540_v63_).set(default__.safeIndex(default__.fm1(p0, d_456_v0_, self.f27, globalState), len(d_540_v63_)), False))), p0) if (d_542_v65_).cf11 else d_541_v64_)
            d_543_v66_: _dafny.Set
            d_543_v66_ = _dafny.Set({self.f27})
            d_544_v67_: _dafny.Map
            d_544_v67_ = _dafny.Map({p0: p0})
            d_545_v68_: _dafny.Seq
            d_545_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eelid"))
            d_546_v69_: _dafny.Seq
            d_546_v69_ = _dafny.SeqWithoutIsStrInference([d_545_v68_])
            d_547_v70_: _dafny.Seq
            d_547_v70_ = _dafny.SeqWithoutIsStrInference([len((d_546_v69_)[default__.safeIndex(self.f27, len(d_546_v69_))])])
            d_548_v71_: _dafny.Map
            d_548_v71_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([self.f27, default__.fm1(p0, d_456_v0_, self.f27, globalState), len(d_543_v66_), -710, len(default__.fm10(self.f27, self.f27, p0, default__.fm0(p0, d_544_v67_, globalState), globalState))])) + (d_547_v70_): D3_DC10(p0, self.f27, p0)})
            d_548_v71_ = d_548_v71_
            (globalState).f8 = (0) - (self.f27)
        d_549_v72_: _dafny.Array
        def lambda17_(d_550_p0_):
            def lambda18_(d_551_i4_):
                return _dafny.Map({806: d_550_p0_})

            return lambda18_

        init8_ = lambda17_(p0)
        nw84_ = _dafny.Array(None, 7)
        for i0_8_ in range(nw84_.length(0)):
            nw84_[i0_8_] = init8_(i0_8_)
        d_549_v72_ = nw84_
        d_552_v73_: _dafny.Map
        d_552_v73_ = _dafny.Map({self.f27: p0})
        index77_ = default__.safeIndex(106, (d_549_v72_).length(0))
        (d_549_v72_)[index77_] = d_552_v73_
        index78_ = default__.safeIndex(106, (d_549_v72_).length(0))
        (d_549_v72_)[index78_] = (d_552_v73_) | ((d_552_v73_) | (d_552_v73_))
        d_553_v74_: C0
        nw85_ = C0()
        nw85_.ctor__()
        d_553_v74_ = nw85_
        (globalState).f23 = p0
        r0 = (0) - (default__.safeModuloInt(self.f27, self.f27))
        r1 = self.f27
        return r0, r1

    @property
    def f26(self):
        return self._f26
