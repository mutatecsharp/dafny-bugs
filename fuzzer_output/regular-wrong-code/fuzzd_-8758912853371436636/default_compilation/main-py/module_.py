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
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_0_i0_ in range(default__.abs(617))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocmloxxh")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwbfxj"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctmkds"))))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        def iife0_():
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(907, 423):
                    d_1_v0_: int = compr_2_
                    if ((907) <= (d_1_v0_)) and ((d_1_v0_) < (423)):
                        coll2_[default__.safeModuloInt(d_1_v0_, (_dafny.MultiSet([849, 974])).cardinality)] = _dafny.CodePoint('v')
                return _dafny.Map(coll2_)
            coll0_ = _dafny.Set()
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(907, 423):
                    d_1_v0_: int = compr_1_
                    if ((907) <= (d_1_v0_)) and ((d_1_v0_) < (423)):
                        coll1_[default__.safeModuloInt(d_1_v0_, (_dafny.MultiSet([849, 974])).cardinality)] = _dafny.CodePoint('v')
                return _dafny.Map(coll1_)
            compr_0_: int
            for compr_0_ in (iife1_()
            ).keys.Elements:
                d_2_v1_: int = compr_0_
                if (d_2_v1_) in (iife2_()
                ):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_2_v1_, 47)]))
            return _dafny.Set(coll0_)
        return (len((_dafny.Map({428: True}) if False else _dafny.Map({-698: not(False)})))) == (len((iife0_()
        ) | (_dafny.Set({956, 701}))))

    @staticmethod
    def fm3(p0, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: str
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_3_i0_ in range(default__.abs(502))])).Elements:
                d_4_v0_: str = compr_3_
                if (d_4_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_3_i0_ in range(default__.abs(502))])):
                    coll3_[d_4_v0_] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vaujf")))]))
            return _dafny.Map(coll3_)
        if (len(iife3_()
        )) == ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uat")))))):
            return _dafny.CodePoint('d')
        elif False:
            return _dafny.CodePoint('u')
        elif True:
            return _dafny.CodePoint('w')

    @staticmethod
    def fm4(p0, p1, globalState):
        return _dafny.MultiSet([(len(_dafny.Map({len(_dafny.Map({True: not(not(False))})): _dafny.MultiSet([len(_dafny.Map({_dafny.CodePoint('i'): not(False)})), len(_dafny.SeqWithoutIsStrInference([-421, len(_dafny.Map({532: False})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnwxeunx")))])), -846, 246, 938])}))) >= (-324)])

    @staticmethod
    def fm5(p0, globalState):
        return len(((_dafny.SeqWithoutIsStrInference([True, False, True, True, False]) if not(True) else _dafny.SeqWithoutIsStrInference([True]))) + (_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def m0(p0, p1, globalState):
        d_5_i0_: int
        d_5_i0_ = 0
        with _dafny.label("0"):
            while p1:
                with _dafny.c_label("0"):
                    if (d_5_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_5_i0_ = (d_5_i0_) + (1)
                    d_6_v0_: _dafny.Array
                    nw0_ = _dafny.Array(False, 12)
                    d_6_v0_ = nw0_
                    d_6_v0_ = d_6_v0_
                    d_7_v1_: int
                    d_7_v1_ = 35
                    d_8_v2_: _dafny.Array
                    def lambda0_(d_9_i1_):
                        return _dafny.CodePoint('j')

                    init0_ = lambda0_
                    nw1_ = _dafny.Array(None, 9)
                    for i0_0_ in range(nw1_.length(0)):
                        nw1_[i0_0_] = init0_(i0_0_)
                    d_8_v2_ = nw1_
                    d_10_v3_: _dafny.Map
                    d_10_v3_ = _dafny.Map({d_7_v1_: _dafny.SeqWithoutIsStrInference([d_8_v2_])})
                    d_11_v4_: _dafny.Map
                    d_11_v4_ = _dafny.Map({d_7_v1_: p1})
                    d_12_v5_: _dafny.Seq
                    d_12_v5_ = _dafny.SeqWithoutIsStrInference([d_8_v2_, d_8_v2_])
                    d_13_v6_: C0
                    nw2_ = C0()
                    nw2_.ctor__(((d_10_v3_)[(0) - (default__.fm5(((d_11_v4_)[(0) - (d_7_v1_)] if ((0) - (d_7_v1_)) in (d_11_v4_) else p0), globalState))] if ((0) - (default__.fm5(((d_11_v4_)[(0) - (d_7_v1_)] if ((0) - (d_7_v1_)) in (d_11_v4_) else p0), globalState))) in (d_10_v3_) else d_12_v5_))
                    d_13_v6_ = nw2_
                    d_14_v7_: bool
                    d_14_v7_ = True
                    d_15_v8_: _dafny.Seq
                    d_15_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                    d_14_v7_ = not((d_7_v1_) >= (len(d_15_v8_)))
                    d_14_v7_ = p1
                    pass
            pass
        d_16_v9_: _dafny.Seq
        d_16_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "peut"))
        d_17_v10_: _dafny.Array
        nw3_ = _dafny.Array(_dafny.CodePoint('D'), 22)
        d_17_v10_ = nw3_
        d_18_v11_: _dafny.Seq
        d_18_v11_ = _dafny.SeqWithoutIsStrInference([d_17_v10_, d_17_v10_, d_17_v10_, d_17_v10_, d_17_v10_])
        d_19_v12_: C0
        nw4_ = C0()
        nw4_.ctor__(d_18_v11_)
        d_19_v12_ = nw4_
        d_20_v13_: _dafny.Map
        d_20_v13_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_21_i2_ in range(default__.abs(-877))])) != (d_16_v9_): d_19_v12_})
        d_20_v13_ = (d_20_v13_).set(p0, d_19_v12_)
        d_22_v14_: D2
        d_22_v14_ = D2_DC9(p1)
        pat_let_tv0_ = d_16_v9_
        pat_let_tv1_ = d_16_v9_
        pat_let_tv2_ = d_16_v9_
        pat_let_tv3_ = d_16_v9_
        def lambda1_(source0_):
            if source0_.is_DC7:
                d_23___mcc_h0_ = source0_.cf11
                d_24___mcc_h1_ = source0_.cf12
                d_25___mcc_h2_ = source0_.cf13
                d_26___mcc_h3_ = source0_.cf14
                d_27_cf14_ = d_26___mcc_h3_
                d_28_cf13_ = d_25___mcc_h2_
                d_29_cf12_ = d_24___mcc_h1_
                d_30_cf11_ = d_23___mcc_h0_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ba"))
            elif source0_.is_DC8:
                d_31___mcc_h4_ = source0_.cf15
                d_32___mcc_h5_ = source0_.cf16
                d_33_cf16_ = d_32___mcc_h5_
                d_34_cf15_ = d_31___mcc_h4_
                return pat_let_tv0_
            elif source0_.is_DC9:
                d_35___mcc_h6_ = source0_.cf17
                d_36_cf17_ = d_35___mcc_h6_
                return (pat_let_tv1_) + (pat_let_tv2_)
            elif source0_.is_DC6:
                d_37___mcc_h7_ = source0_.cf10
                d_38_cf10_ = d_37___mcc_h7_
                return (d_38_cf10_) + (pat_let_tv3_)
            elif True:
                d_39___mcc_h8_ = source0_.cf18
                d_40_cf18_ = d_39___mcc_h8_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eohfqvsl"))

        d_16_v9_ = lambda1_(d_22_v14_)
        d_16_v9_ = d_16_v9_
        if not (not (p0) or (p0)) or (p0):
            d_41_v15_: bool
            d_41_v15_ = True
            d_42_v16_: int
            d_42_v16_ = 348
            d_43_v17_: _dafny.MultiSet
            d_43_v17_ = _dafny.MultiSet([d_42_v16_])
            d_41_v15_ = ((len(d_16_v9_)) * (d_42_v16_)) in (d_43_v17_)
            d_41_v15_ = p1
            d_41_v15_ = d_41_v15_
            d_44_v18_: T0
            nw5_ = C0()
            nw5_.ctor__(d_18_v11_)
            d_44_v18_ = nw5_
            d_45_v19_: D1
            d_45_v19_ = D1_DC3(d_44_v18_)
            d_46_v20_: D1
            d_46_v20_ = D1_DC5(d_45_v19_)
            d_47_v21_: _dafny.Map
            d_47_v21_ = _dafny.Map({d_41_v15_: _dafny.SeqWithoutIsStrInference([d_46_v20_])})
            d_48_v22_: _dafny.Seq
            d_48_v22_ = _dafny.SeqWithoutIsStrInference([D1_DC5(d_45_v19_), d_46_v20_, d_46_v20_, D1_DC5(d_45_v19_), d_46_v20_])
            d_49_v23_: _dafny.Seq
            d_49_v23_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_50_v24_: _dafny.Seq
            d_50_v24_ = _dafny.SeqWithoutIsStrInference([((d_47_v21_)[not(p0)] if (not(p0)) in (d_47_v21_) else d_48_v22_), (_dafny.SeqWithoutIsStrInference([d_46_v20_])) + (d_48_v22_), (d_48_v22_ if (d_49_v23_)[default__.safeIndex(d_42_v16_, len(d_49_v23_))] else d_48_v22_)])
            d_42_v16_ = (0) - ((0) - (len((d_50_v24_)[default__.safeIndex(default__.fm5(d_41_v15_, globalState), len(d_50_v24_))])))
            d_44_v18_ = d_44_v18_
        elif True:
            d_51_v25_: int
            d_51_v25_ = 534
            d_52_v26_: bool
            d_52_v26_ = True
            d_53_v27_: _dafny.Map
            d_53_v27_ = _dafny.Map({not(p0): p0})
            rhs0_ = d_51_v25_
            rhs1_ = (d_51_v25_ if (p0) or (((d_53_v27_)[d_52_v26_] if (d_52_v26_) in (d_53_v27_) else p0)) else d_51_v25_)
            rhs2_ = True
            d_51_v25_ = rhs0_
            d_51_v25_ = rhs1_
            d_52_v26_ = rhs2_
            if (not (p1) or (p1) if d_52_v26_ else (default__.fm5(d_52_v26_, globalState)) < (d_51_v25_)):
                d_54_v28_: _dafny.Array
                nw6_ = _dafny.Array(None, 22)
                d_54_v28_ = nw6_
                index0_ = default__.safeIndex(184, (d_54_v28_).length(0))
                (d_54_v28_)[index0_] = d_19_v12_
                index1_ = default__.safeIndex(184, (d_54_v28_).length(0))
                (d_54_v28_)[index1_] = d_19_v12_
                d_55_v29_: _dafny.Seq
                d_55_v29_ = _dafny.SeqWithoutIsStrInference([d_20_v13_, d_20_v13_])
                d_51_v25_ = len(d_55_v29_)
                d_51_v25_ = d_51_v25_
                d_52_v26_ = p0
                d_56_v30_: _dafny.Array
                def lambda2_(d_57_v25_):
                    def lambda3_(d_58_i3_):
                        return default__.safeDivisionInt(d_58_i3_, d_57_v25_)

                    return lambda3_

                init1_ = lambda2_(d_51_v25_)
                nw7_ = _dafny.Array(None, 3)
                for i0_1_ in range(nw7_.length(0)):
                    nw7_[i0_1_] = init1_(i0_1_)
                d_56_v30_ = nw7_
                d_56_v30_ = d_56_v30_
            elif True:
                d_52_v26_ = p0
                d_59_v31_: C0
                nw8_ = C0()
                nw8_.ctor__((d_18_v11_) + (d_18_v11_))
                d_59_v31_ = nw8_
                d_60_v32_: str
                d_60_v32_ = _dafny.CodePoint('u')
                d_61_v33_: _dafny.MultiSet
                d_61_v33_ = _dafny.MultiSet([d_60_v32_])
                d_51_v25_ = ((d_61_v33_)[d_60_v32_] if (d_60_v32_) in (d_61_v33_) else (d_59_v31_).fm1(p0, d_52_v26_, True, p1, globalState))
                d_19_v12_ = d_19_v12_
                d_62_v34_: T0
                nw9_ = C0()
                nw9_.ctor__(_dafny.SeqWithoutIsStrInference([d_17_v10_]))
                d_62_v34_ = nw9_
                d_62_v34_ = d_62_v34_
            d_63_v35_: _dafny.Set
            d_63_v35_ = _dafny.Set({False})
            d_64_v36_: D2
            d_64_v36_ = D2_DC7(d_16_v9_, len(d_63_v35_), d_51_v25_, not(True))
            d_65_v37_: _dafny.Map
            d_65_v37_ = _dafny.Map({d_64_v36_: d_52_v26_})
            d_65_v37_ = _dafny.Map({d_64_v36_: d_52_v26_})
            d_66_v38_: str
            d_66_v38_ = _dafny.CodePoint('t')
            d_67_v39_: D0
            d_67_v39_ = D0_DC1(d_52_v26_, default__.fm5(d_52_v26_, globalState), d_51_v25_, False, d_51_v25_)
            d_68_v40_: _dafny.Map
            d_68_v40_ = _dafny.Map({d_66_v38_: (d_67_v39_).cf3})
            d_68_v40_ = (d_68_v40_).set(_dafny.CodePoint('h'), d_51_v25_)
            if d_52_v26_:
                d_51_v25_ = (d_51_v25_) * ((len(d_16_v9_)) - (436))
                d_66_v38_ = d_66_v38_
                d_69_v41_: C0
                nw10_ = C0()
                nw10_.ctor__(d_18_v11_)
                d_69_v41_ = nw10_
                d_70_v42_: _dafny.Set
                d_70_v42_ = _dafny.Set({d_51_v25_})
                d_71_v43_: _dafny.Seq
                d_71_v43_ = _dafny.SeqWithoutIsStrInference([d_69_v41_, d_69_v41_])
                d_72_v44_: _dafny.Array
                nw11_ = _dafny.Array(None, 15)
                nw11_[int(0)] = (0) - (default__.safeModuloInt(d_51_v25_, (0) - (d_51_v25_)))
                nw11_[int(1)] = (len(d_70_v42_)) - (d_51_v25_)
                nw11_[int(2)] = d_51_v25_
                nw11_[int(3)] = d_51_v25_
                nw11_[int(4)] = len(d_16_v9_)
                nw11_[int(5)] = d_51_v25_
                nw11_[int(6)] = len(d_71_v43_)
                nw11_[int(7)] = d_51_v25_
                nw11_[int(8)] = (d_64_v36_).cf13
                nw11_[int(9)] = (0) - (len(d_16_v9_))
                nw11_[int(10)] = d_51_v25_
                nw11_[int(11)] = (d_64_v36_).cf13
                nw11_[int(12)] = d_51_v25_
                nw11_[int(13)] = d_51_v25_
                nw11_[int(14)] = d_51_v25_
                d_72_v44_ = nw11_
                index2_ = default__.safeIndex(780, (d_72_v44_).length(0))
                (d_72_v44_)[index2_] = d_51_v25_
                d_73_v45_: _dafny.Seq
                d_73_v45_ = _dafny.SeqWithoutIsStrInference([d_52_v26_, p0])
                d_74_v46_: D4
                d_74_v46_ = D4_DC14(d_73_v45_, _dafny.CodePoint('l'), d_51_v25_, d_51_v25_)
                d_75_v47_: _dafny.Set
                d_75_v47_ = _dafny.Set({(d_74_v46_).cf22, d_66_v38_, d_66_v38_})
                d_76_v48_: D3
                d_76_v48_ = D3_DC11(d_75_v47_)
                pat_let_tv4_ = d_75_v47_
                pat_let_tv5_ = d_75_v47_
                index3_ = default__.safeIndex(780, (d_72_v44_).length(0))
                def iife4_(_pat_let0_0):
                    def iife5_(d_77_dt__update__tmp_h0_):
                        def iife7_():
                            coll4_ = _dafny.Set()
                            compr_4_: str
                            for compr_4_ in (pat_let_tv4_).Elements:
                                d_78_v49_: str = compr_4_
                                if (d_78_v49_) in (pat_let_tv5_):
                                    coll4_ = coll4_.union(_dafny.Set([d_78_v49_]))
                            return _dafny.Set(coll4_)
                        def iife6_(_pat_let1_0):
                            def iife8_(d_79_dt__update_hcf19_h0_):
                                return D3_DC11(d_79_dt__update_hcf19_h0_)
                            return iife8_(_pat_let1_0)
                        return iife6_(iife7_()
                        )
                    return iife5_(_pat_let0_0)
                rhs3_ = not (p0) or (True)
                rhs4_ = len((iife4_(d_76_v48_)).cf19)
                lhs0_ = d_72_v44_
                lhs1_ = default__.safeIndex(780, (d_72_v44_).length(0))
                d_52_v26_ = rhs3_
                lhs0_[lhs1_] = rhs4_
                d_73_v45_ = d_73_v45_
            elif True:
                d_51_v25_ = d_51_v25_
                d_80_v50_: T0
                nw12_ = C0()
                nw12_.ctor__(d_18_v11_)
                d_80_v50_ = nw12_
                d_81_v51_: D2
                d_81_v51_ = D2_DC8(p0, d_80_v50_)
                d_52_v26_ = ((d_81_v51_ if True else d_81_v51_)).cf15
                d_82_v52_: _dafny.MultiSet
                d_82_v52_ = _dafny.MultiSet([_dafny.CodePoint('v'), d_66_v38_])
                d_83_v53_: _dafny.Map
                d_83_v53_ = _dafny.Map({(d_67_v39_).cf2: d_16_v9_})
                d_84_v54_: _dafny.Seq
                d_84_v54_ = _dafny.SeqWithoutIsStrInference([(d_82_v52_).issubset(_dafny.MultiSet([d_66_v38_, d_66_v38_])), not(False), (d_51_v25_) <= (len(((d_83_v53_)[d_51_v25_] if (d_51_v25_) in (d_83_v53_) else d_16_v9_)))])
                d_84_v54_ = (d_84_v54_) + ((d_84_v54_).set(default__.safeIndex(len(d_16_v9_), len(d_84_v54_)), p1))
                d_51_v25_ = (d_51_v25_) + (d_51_v25_)
                d_51_v25_ = d_51_v25_
        d_85_v55_: int
        d_85_v55_ = 115
        d_86_v56_: _dafny.Set
        d_86_v56_ = _dafny.Set({d_85_v55_, d_85_v55_})
        d_87_v57_: _dafny.Map
        d_87_v57_ = _dafny.Map({d_86_v56_: d_16_v9_})
        d_87_v57_ = (d_87_v57_) | (d_87_v57_)

    @staticmethod
    def Main(noArgsParameter__):
        d_88_v0_: int
        d_88_v0_ = 159
        d_89_v1_: _dafny.Seq
        d_89_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqomatru"))
        d_90_v2_: _dafny.Seq
        d_90_v2_ = _dafny.SeqWithoutIsStrInference([d_88_v0_, (0) - ((0) - (d_88_v0_)), d_88_v0_, len(d_89_v1_)])
        d_91_globalState_: GlobalState
        nw13_ = GlobalState()
        def iife9_():
            coll5_ = _dafny.Set()
            compr_5_: int
            for compr_5_ in (d_90_v2_).Elements:
                d_92_v3_: int = compr_5_
                if (d_92_v3_) in (d_90_v2_):
                    coll5_ = coll5_.union(_dafny.Set([(d_92_v3_) + (477)]))
            return _dafny.Set(coll5_)
        nw13_.ctor__(862, 771, 785, iife9_()
        , -185, _dafny.CodePoint('u'), False, 385, True)
        d_91_globalState_ = nw13_
        d_93_v4_: D0
        d_93_v4_ = D0_DC0(len((default__.fm0(d_91_globalState_)) + (d_89_v1_)))
        source1_ = d_93_v4_
        if source1_.is_DC0:
            d_94___mcc_h0_ = source1_.cf0
            d_95_cf0_ = d_94___mcc_h0_
            d_88_v0_ = d_95_cf0_
            d_96_v5_: bool
            d_96_v5_ = False
            d_96_v5_ = d_96_v5_
            d_97_v6_: _dafny.Set
            d_97_v6_ = _dafny.Set({d_89_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erdcbq")), (d_89_v1_) + (d_89_v1_)})
            d_98_v7_: _dafny.Array
            nw14_ = _dafny.Array(int(0), 13)
            d_98_v7_ = nw14_
            index4_ = default__.safeIndex(895, (d_98_v7_).length(0))
            (d_98_v7_)[index4_] = d_88_v0_
            index5_ = default__.safeIndex(895, (d_98_v7_).length(0))
            rhs5_ = (d_97_v6_) - (_dafny.Set({d_89_v1_}))
            rhs6_ = d_88_v0_
            lhs2_ = d_98_v7_
            lhs3_ = default__.safeIndex(895, (d_98_v7_).length(0))
            d_97_v6_ = rhs5_
            lhs2_[lhs3_] = rhs6_
            d_99_v8_: _dafny.Seq
            d_99_v8_ = _dafny.SeqWithoutIsStrInference([d_96_v5_, (d_96_v5_) and (d_96_v5_)])
            d_96_v5_ = (d_99_v8_)[default__.safeIndex(d_95_cf0_, len(d_99_v8_))]
        elif source1_.is_DC1:
            d_100___mcc_h1_ = source1_.cf1
            d_101___mcc_h2_ = source1_.cf2
            d_102___mcc_h3_ = source1_.cf3
            d_103___mcc_h4_ = source1_.cf4
            d_104___mcc_h5_ = source1_.cf5
            d_105_cf5_ = d_104___mcc_h5_
            d_106_cf4_ = d_103___mcc_h4_
            d_107_cf3_ = d_102___mcc_h3_
            d_108_cf2_ = d_101___mcc_h2_
            d_109_cf1_ = d_100___mcc_h1_
            d_110_v9_: _dafny.Set
            d_110_v9_ = _dafny.Set({False})
            d_111_v10_: _dafny.Set
            d_111_v10_ = _dafny.Set({d_110_v9_, d_110_v9_})
            d_106_cf4_ = not((d_111_v10_).issubset((d_111_v10_) | (d_111_v10_)))
            default__.m0(d_109_cf1_, d_109_cf1_, d_91_globalState_)
            d_112_v11_: _dafny.Array
            nw15_ = _dafny.Array(False, 23)
            d_112_v11_ = nw15_
            index6_ = default__.safeIndex(611, (d_112_v11_).length(0))
            (d_112_v11_)[index6_] = d_109_cf1_
            index7_ = default__.safeIndex(611, (d_112_v11_).length(0))
            (d_112_v11_)[index7_] = ((d_88_v0_) > (d_108_cf2_) if d_109_cf1_ else not (d_106_cf4_) or (False))
            d_113_v12_: _dafny.Array
            nw16_ = _dafny.Array(None, 1)
            nw16_[int(0)] = _dafny.CodePoint('o')
            d_113_v12_ = nw16_
            d_114_v13_: _dafny.Seq
            d_114_v13_ = _dafny.SeqWithoutIsStrInference([d_113_v12_])
            d_115_v14_: C0
            nw17_ = C0()
            nw17_.ctor__((d_114_v13_).set(default__.safeIndex(d_105_cf5_, len(d_114_v13_)), d_113_v12_))
            d_115_v14_ = nw17_
        elif True:
            d_116___mcc_h6_ = source1_.cf6
            d_117_cf6_ = d_116___mcc_h6_
            d_118_v15_: bool
            d_118_v15_ = True
            d_119_v16_: _dafny.Map
            d_119_v16_ = _dafny.Map({(d_118_v15_) == (d_118_v15_): d_88_v0_})
            d_119_v16_ = (d_119_v16_).set(d_118_v15_, d_88_v0_)
            d_120_v17_: str
            d_120_v17_ = _dafny.CodePoint('d')
            d_118_v15_ = default__.fm2(not(d_118_v15_), d_120_v17_, d_88_v0_, (d_118_v15_ if not(d_118_v15_) else d_118_v15_), d_91_globalState_)
            d_121_v18_: _dafny.Map
            d_121_v18_ = _dafny.Map({(0) - (d_88_v0_): (d_118_v15_) or (d_118_v15_)})
            d_121_v18_ = (d_121_v18_).set(-558, (d_118_v15_) or (False))
            d_122_v19_: _dafny.Array
            nw18_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_122_v19_ = nw18_
            d_123_v20_: _dafny.Array
            nw19_ = _dafny.Array(D0.default()(), 9)
            d_123_v20_ = nw19_
            index8_ = default__.safeIndex(937, (d_122_v19_).length(0))
            (d_122_v19_)[index8_] = d_123_v20_
            index9_ = default__.safeIndex(937, (d_122_v19_).length(0))
            (d_122_v19_)[index9_] = d_123_v20_
        d_124_v21_: _dafny.Array
        nw20_ = _dafny.Array(_dafny.Set({}), 14)
        d_124_v21_ = nw20_
        d_125_v22_: str
        d_125_v22_ = _dafny.CodePoint('h')
        d_126_v23_: _dafny.Array
        nw21_ = _dafny.Array(None, 12)
        nw21_[int(0)] = d_125_v22_
        nw21_[int(1)] = d_125_v22_
        nw21_[int(2)] = d_125_v22_
        nw21_[int(3)] = d_125_v22_
        nw21_[int(4)] = d_125_v22_
        nw21_[int(5)] = d_125_v22_
        nw21_[int(6)] = d_125_v22_
        nw21_[int(7)] = _dafny.CodePoint('q')
        nw21_[int(8)] = d_125_v22_
        nw21_[int(9)] = d_125_v22_
        nw21_[int(10)] = d_125_v22_
        nw21_[int(11)] = d_125_v22_
        d_126_v23_ = nw21_
        d_127_v24_: _dafny.Seq
        d_127_v24_ = _dafny.SeqWithoutIsStrInference([d_126_v23_])
        d_128_v25_: C0
        nw22_ = C0()
        nw22_.ctor__(d_127_v24_)
        d_128_v25_ = nw22_
        d_129_v26_: _dafny.Set
        d_129_v26_ = _dafny.Set({d_128_v25_, d_128_v25_})
        index10_ = default__.safeIndex(611, (d_124_v21_).length(0))
        (d_124_v21_)[index10_] = d_129_v26_
        index11_ = default__.safeIndex(611, (d_124_v21_).length(0))
        (d_124_v21_)[index11_] = d_129_v26_
        source2_ = d_93_v4_
        if source2_.is_DC0:
            d_130___mcc_h7_ = source2_.cf0
            d_131_cf0_ = d_130___mcc_h7_
            d_89_v1_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rllcb"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))).set(default__.safeIndex(d_88_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))), d_125_v22_))
            d_132_v27_: T0
            nw23_ = C0()
            nw23_.ctor__(_dafny.SeqWithoutIsStrInference([d_126_v23_, d_126_v23_]))
            d_132_v27_ = nw23_
            d_132_v27_ = d_132_v27_
            d_88_v0_ = d_88_v0_
            d_133_v28_: _dafny.Array
            def lambda4_(d_134_i0_):
                return (False) == (True)

            init2_ = lambda4_
            nw24_ = _dafny.Array(None, 9)
            for i0_2_ in range(nw24_.length(0)):
                nw24_[i0_2_] = init2_(i0_2_)
            d_133_v28_ = nw24_
            d_135_v29_: bool
            d_135_v29_ = True
            index12_ = default__.safeIndex(208, (d_133_v28_).length(0))
            (d_133_v28_)[index12_] = d_135_v29_
            d_136_v30_: _dafny.Set
            d_136_v30_ = _dafny.Set({d_135_v29_})
            index13_ = default__.safeIndex(208, (d_133_v28_).length(0))
            (d_133_v28_)[index13_] = (not (d_135_v29_) or ((D0_DC1(d_135_v29_, len(d_129_v26_), (0) - (d_88_v0_), d_135_v29_, 556)).cf4) if d_135_v29_ else (len(d_136_v30_)) <= (d_88_v0_))
        elif source2_.is_DC1:
            d_137___mcc_h8_ = source2_.cf1
            d_138___mcc_h9_ = source2_.cf2
            d_139___mcc_h10_ = source2_.cf3
            d_140___mcc_h11_ = source2_.cf4
            d_141___mcc_h12_ = source2_.cf5
            d_142_cf5_ = d_141___mcc_h12_
            d_143_cf4_ = d_140___mcc_h11_
            d_144_cf3_ = d_139___mcc_h10_
            d_145_cf2_ = d_138___mcc_h9_
            d_146_cf1_ = d_137___mcc_h8_
            d_147_v31_: _dafny.Array
            nw25_ = _dafny.Array(int(0), 22)
            d_147_v31_ = nw25_
            d_148_v32_: _dafny.Array
            nw26_ = _dafny.Array(False, 14)
            d_148_v32_ = nw26_
            d_149_v33_: _dafny.Map
            d_149_v33_ = _dafny.Map({d_147_v31_: d_148_v32_})
            d_149_v33_ = (d_149_v33_).set(d_147_v31_, d_148_v32_)
            d_128_v25_ = d_128_v25_
            d_89_v1_ = (d_89_v1_) + (d_89_v1_)
            if d_143_cf4_:
                d_150_v34_: D0
                d_150_v34_ = D0_DC1(False, d_142_cf5_, d_88_v0_, d_146_cf1_, d_145_cf2_)
                d_144_cf3_ = (d_150_v34_).cf2
                d_151_v35_: _dafny.Seq
                d_151_v35_ = _dafny.SeqWithoutIsStrInference([d_148_v32_, d_148_v32_])
                d_152_v36_: _dafny.Seq
                d_152_v36_ = _dafny.SeqWithoutIsStrInference([(d_151_v35_)[default__.safeIndex(d_144_cf3_, len(d_151_v35_))], d_148_v32_, d_148_v32_])
                d_153_v37_: _dafny.Map
                d_153_v37_ = _dafny.Map({d_146_cf1_: len((d_151_v35_) + (d_151_v35_))})
                d_154_v38_: _dafny.Seq
                d_154_v38_ = _dafny.SeqWithoutIsStrInference([d_146_cf1_])
                d_155_v39_: _dafny.Map
                d_155_v39_ = _dafny.Map({d_89_v1_: d_154_v38_})
                d_153_v37_ = (d_153_v37_).set((_dafny.Map({d_89_v1_: d_154_v38_})) == (d_155_v39_), d_88_v0_)
                index14_ = default__.safeIndex(327, (d_147_v31_).length(0))
                (d_147_v31_)[index14_] = (d_145_cf2_ if d_143_cf4_ else (0) - ((0) - (d_88_v0_)))
                d_156_v40_: _dafny.Set
                d_156_v40_ = _dafny.Set({(d_128_v25_).fm1(default__.fm2(d_146_cf1_, _dafny.CodePoint('w'), len(d_90_v2_), not(d_143_cf4_), d_91_globalState_), d_146_cf1_, d_146_cf1_, d_146_cf1_, d_91_globalState_), d_144_cf3_, d_145_cf2_})
                index15_ = default__.safeIndex(327, (d_147_v31_).length(0))
                (d_147_v31_)[index15_] = (d_88_v0_) + (len(d_156_v40_))
                d_146_cf1_ = (d_150_v34_).cf1
                d_146_cf1_ = not(not(default__.fm2(d_143_cf4_, (d_125_v22_ if d_146_cf1_ else default__.fm3((d_150_v34_).cf3, d_91_globalState_)), (d_144_cf3_) - ((0) - (d_88_v0_)), d_146_cf1_, d_91_globalState_)))
            elif True:
                default__.m0((d_143_cf4_) == (False), d_146_cf1_, d_91_globalState_)
                d_88_v0_ = ((d_90_v2_)[default__.safeIndex(d_142_cf5_, len(d_90_v2_))]) + (-475)
                index16_ = default__.safeIndex(154, (d_148_v32_).length(0))
                (d_148_v32_)[index16_] = d_146_cf1_
                index17_ = default__.safeIndex(154, (d_148_v32_).length(0))
                (d_148_v32_)[index17_] = (d_128_v25_) in (d_129_v26_)
                d_157_v41_: _dafny.Array
                nw27_ = _dafny.Array(_dafny.Map({}), 19)
                d_157_v41_ = nw27_
                d_157_v41_ = d_157_v41_
                d_158_v42_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_158_v42_ = nw28_
                d_158_v42_ = d_158_v42_
        elif True:
            d_159___mcc_h13_ = source2_.cf6
            d_160_cf6_ = d_159___mcc_h13_
            d_161_v43_: bool
            d_161_v43_ = False
            if not(d_161_v43_):
                d_88_v0_ = d_88_v0_
                d_162_v44_: _dafny.Seq
                d_162_v44_ = _dafny.SeqWithoutIsStrInference([d_161_v43_, d_161_v43_, d_161_v43_, d_161_v43_, d_161_v43_])
                d_163_v45_: _dafny.MultiSet
                d_163_v45_ = _dafny.MultiSet([default__.fm2(d_161_v43_, d_125_v22_, d_88_v0_, not(d_161_v43_), d_91_globalState_)])
                d_164_v46_: _dafny.Map
                d_164_v46_ = _dafny.Map({d_88_v0_: d_88_v0_})
                d_165_v47_: D0
                d_165_v47_ = D0_DC1(d_161_v43_, (d_163_v45_).cardinality, len(d_164_v46_), d_161_v43_, d_88_v0_)
                d_166_v48_: T0
                nw29_ = C0()
                nw29_.ctor__(d_127_v24_)
                d_166_v48_ = nw29_
                d_167_v49_: _dafny.Map
                d_167_v49_ = _dafny.Map({d_166_v48_: d_161_v43_})
                d_168_v50_: _dafny.MultiSet
                d_168_v50_ = _dafny.MultiSet([d_88_v0_, len(_dafny.SeqWithoutIsStrInference([d_125_v22_ for d_169_i1_ in range(default__.abs(529))]))])
                d_170_v51_: D0
                d_170_v51_ = D0_DC1((d_165_v47_).cf4, d_88_v0_, len((d_167_v49_).set(d_166_v48_, d_161_v43_)), d_161_v43_, ((d_168_v50_)[d_88_v0_] if (d_88_v0_) in (d_168_v50_) else (d_90_v2_)[default__.safeIndex(d_88_v0_, len(d_90_v2_))]))
                d_171_v52_: _dafny.MultiSet
                d_171_v52_ = _dafny.MultiSet([d_128_v25_])
                d_172_v53_: _dafny.Array
                nw30_ = _dafny.Array(None, 12)
                nw30_[int(0)] = d_88_v0_
                nw30_[int(1)] = len(d_162_v44_)
                nw30_[int(2)] = d_88_v0_
                nw30_[int(3)] = 916
                nw30_[int(4)] = (d_165_v47_).cf5
                nw30_[int(5)] = d_88_v0_
                nw30_[int(6)] = default__.safeDivisionInt((d_171_v52_).cardinality, (0) - ((d_166_v48_).fm1(d_161_v43_, d_161_v43_, d_161_v43_, False, d_91_globalState_)))
                nw30_[int(7)] = len(d_89_v1_)
                nw30_[int(8)] = d_88_v0_
                nw30_[int(9)] = 810
                nw30_[int(10)] = d_88_v0_
                nw30_[int(11)] = d_88_v0_
                d_172_v53_ = nw30_
                index18_ = default__.safeIndex(266, (d_172_v53_).length(0))
                (d_172_v53_)[index18_] = d_88_v0_
                index19_ = default__.safeIndex(266, (d_172_v53_).length(0))
                (d_172_v53_)[index19_] = (d_166_v48_).fm1((d_88_v0_) < (d_88_v0_), d_161_v43_, d_161_v43_, (default__.fm2(d_161_v43_, d_125_v22_, len(d_89_v1_), d_161_v43_, d_91_globalState_) if default__.fm2(default__.fm2(d_161_v43_, d_125_v22_, d_88_v0_, d_161_v43_, d_91_globalState_), d_125_v22_, d_88_v0_, d_161_v43_, d_91_globalState_) else False), d_91_globalState_)
                index20_ = default__.safeIndex(266, (d_172_v53_).length(0))
                (d_172_v53_)[index20_] = (d_128_v25_).fm1(d_161_v43_, default__.fm2(d_161_v43_, d_125_v22_, -962, d_161_v43_, d_91_globalState_), d_161_v43_, not((d_168_v50_).issubset(d_168_v50_)), d_91_globalState_)
                d_88_v0_ = len(d_89_v1_)
                d_161_v43_ = not(d_161_v43_)
            elif True:
                d_173_v54_: T0
                nw31_ = C0()
                nw31_.ctor__(d_127_v24_)
                d_173_v54_ = nw31_
                d_174_v55_: _dafny.Seq
                d_174_v55_ = _dafny.SeqWithoutIsStrInference([d_173_v54_, d_173_v54_, d_173_v54_, d_173_v54_])
                d_175_v56_: D1
                d_175_v56_ = D1_DC3((d_174_v55_)[default__.safeIndex(d_88_v0_, len(d_174_v55_))])
                d_176_v57_: _dafny.Array
                nw32_ = _dafny.Array(None, 21)
                nw32_[int(0)] = d_173_v54_
                nw32_[int(1)] = (D1_DC3(d_173_v54_)).cf7
                nw32_[int(2)] = d_173_v54_
                nw32_[int(3)] = (d_175_v56_).cf7
                nw32_[int(4)] = d_173_v54_
                nw32_[int(5)] = d_173_v54_
                nw32_[int(6)] = d_173_v54_
                nw32_[int(7)] = d_173_v54_
                nw32_[int(8)] = d_173_v54_
                nw32_[int(9)] = d_173_v54_
                nw32_[int(10)] = d_173_v54_
                nw32_[int(11)] = d_173_v54_
                nw32_[int(12)] = d_173_v54_
                nw32_[int(13)] = d_173_v54_
                nw32_[int(14)] = d_173_v54_
                nw32_[int(15)] = d_173_v54_
                nw32_[int(16)] = d_173_v54_
                nw32_[int(17)] = d_173_v54_
                nw32_[int(18)] = d_173_v54_
                nw32_[int(19)] = d_173_v54_
                nw32_[int(20)] = d_173_v54_
                d_176_v57_ = nw32_
                index21_ = default__.safeIndex(216, (d_176_v57_).length(0))
                (d_176_v57_)[index21_] = d_173_v54_
                index22_ = default__.safeIndex(216, (d_176_v57_).length(0))
                (d_176_v57_)[index22_] = (d_173_v54_ if d_161_v43_ else (d_173_v54_ if d_161_v43_ else d_173_v54_))
                d_177_v58_: _dafny.Map
                d_177_v58_ = _dafny.Map({d_161_v43_: False})
                d_177_v58_ = (d_177_v58_).set(d_161_v43_, d_161_v43_)
                d_178_v59_: _dafny.Array
                nw33_ = _dafny.Array(False, 14)
                d_178_v59_ = nw33_
                d_178_v59_ = d_178_v59_
                d_179_v60_: _dafny.Array
                nw34_ = _dafny.Array(_dafny.Map({}), 24)
                d_179_v60_ = nw34_
                nw35_ = _dafny.Array(_dafny.Map({}), 17)
                d_179_v60_ = nw35_
                d_180_v61_: _dafny.Seq
                d_180_v61_ = _dafny.SeqWithoutIsStrInference([d_161_v43_, d_161_v43_, d_161_v43_])
                d_180_v61_ = d_180_v61_
            d_181_v62_: _dafny.Array
            nw36_ = _dafny.Array(int(0), 16)
            d_181_v62_ = nw36_
            index23_ = default__.safeIndex(992, (d_181_v62_).length(0))
            (d_181_v62_)[index23_] = d_88_v0_
            index24_ = default__.safeIndex(992, (d_181_v62_).length(0))
            (d_181_v62_)[index24_] = d_88_v0_
            d_161_v43_ = d_161_v43_
            d_88_v0_ = d_88_v0_
        d_182_v63_: bool
        d_182_v63_ = True
        if d_182_v63_:
            d_182_v63_ = d_182_v63_
            d_183_v64_: _dafny.Map
            d_183_v64_ = _dafny.Map({d_88_v0_: False})
            d_88_v0_ = ((len(d_183_v64_)) * (d_88_v0_)) + (d_88_v0_)
            d_89_v1_ = default__.fm0(d_91_globalState_)
            if not(d_182_v63_):
                d_182_v63_ = d_182_v63_
                d_184_v65_: _dafny.Array
                nw37_ = _dafny.Array(False, 20)
                d_184_v65_ = nw37_
                d_185_v66_: _dafny.Seq
                d_185_v66_ = _dafny.SeqWithoutIsStrInference([d_182_v63_])
                d_186_v67_: _dafny.Set
                d_186_v67_ = _dafny.Set({len(d_185_v66_), d_88_v0_})
                d_187_v68_: D0
                d_187_v68_ = D0_DC1(d_182_v63_, d_88_v0_, d_88_v0_, d_182_v63_, len(d_186_v67_))
                index25_ = default__.safeIndex(698, (d_184_v65_).length(0))
                (d_184_v65_)[index25_] = (d_187_v68_).cf1
                index26_ = default__.safeIndex(698, (d_184_v65_).length(0))
                (d_184_v65_)[index26_] = True
                index27_ = default__.safeIndex(698, (d_184_v65_).length(0))
                (d_184_v65_)[index27_] = d_182_v63_
                d_188_v69_: _dafny.Map
                d_188_v69_ = _dafny.Map({(d_184_v65_)[default__.safeIndex(698, (d_184_v65_).length(0))]: d_185_v66_})
                d_188_v69_ = d_188_v69_
                d_189_v70_: T0
                nw38_ = C0()
                nw38_.ctor__(d_127_v24_)
                d_189_v70_ = nw38_
                d_190_v71_: _dafny.Map
                d_190_v71_ = _dafny.Map({d_189_v70_: d_88_v0_})
                d_190_v71_ = (d_190_v71_).set(d_189_v70_, d_88_v0_)
            elif True:
                d_191_v72_: _dafny.Map
                d_191_v72_ = _dafny.Map({d_182_v63_: _dafny.CodePoint('m')})
                d_191_v72_ = d_191_v72_
                d_192_v73_: _dafny.MultiSet
                d_192_v73_ = _dafny.MultiSet([d_182_v63_])
                d_193_v74_: _dafny.Seq
                d_193_v74_ = _dafny.SeqWithoutIsStrInference([d_182_v63_])
                d_194_v75_: _dafny.Array
                nw39_ = _dafny.Array(None, 26)
                nw39_[int(0)] = d_192_v73_
                nw39_[int(1)] = default__.fm4(False, d_182_v63_, d_91_globalState_)
                nw39_[int(2)] = d_192_v73_
                nw39_[int(3)] = _dafny.MultiSet([False, d_182_v63_, d_182_v63_, d_182_v63_, d_182_v63_])
                nw39_[int(4)] = d_192_v73_
                nw39_[int(5)] = d_192_v73_
                nw39_[int(6)] = d_192_v73_
                nw39_[int(7)] = _dafny.MultiSet([d_182_v63_])
                nw39_[int(8)] = d_192_v73_
                nw39_[int(9)] = d_192_v73_
                nw39_[int(10)] = default__.fm4(d_182_v63_, False, d_91_globalState_)
                nw39_[int(11)] = (d_192_v73_).set(d_182_v63_, default__.abs(d_88_v0_))
                nw39_[int(12)] = d_192_v73_
                nw39_[int(13)] = d_192_v73_
                nw39_[int(14)] = d_192_v73_
                nw39_[int(15)] = d_192_v73_
                nw39_[int(16)] = d_192_v73_
                nw39_[int(17)] = d_192_v73_
                nw39_[int(18)] = _dafny.MultiSet([d_182_v63_])
                nw39_[int(19)] = d_192_v73_
                nw39_[int(20)] = _dafny.MultiSet([d_182_v63_])
                nw39_[int(21)] = d_192_v73_
                nw39_[int(22)] = d_192_v73_
                nw39_[int(23)] = default__.fm4(default__.fm2(d_182_v63_, (d_89_v1_)[default__.safeIndex(len(d_183_v64_), len(d_89_v1_))], d_88_v0_, d_182_v63_, d_91_globalState_), (d_193_v74_)[default__.safeIndex(96, len(d_193_v74_))], d_91_globalState_)
                nw39_[int(24)] = d_192_v73_
                nw39_[int(25)] = d_192_v73_
                d_194_v75_ = nw39_
                d_195_v76_: _dafny.Map
                d_195_v76_ = _dafny.Map({d_194_v75_: not (d_182_v63_) or (d_182_v63_)})
                d_182_v63_ = ((d_195_v76_)[d_194_v75_] if (d_194_v75_) in (d_195_v76_) else d_182_v63_)
                d_88_v0_ = d_88_v0_
                index28_ = default__.safeIndex(611, (d_124_v21_).length(0))
                (d_124_v21_)[index28_] = ((d_124_v21_)[default__.safeIndex(611, (d_124_v21_).length(0))]) - (d_129_v26_)
                d_196_v77_: _dafny.Array
                def lambda5_(d_197_v1_):
                    def lambda6_(d_198_i2_):
                        return d_197_v1_

                    return lambda6_

                init3_ = lambda5_(d_89_v1_)
                nw40_ = _dafny.Array(None, 12)
                for i0_3_ in range(nw40_.length(0)):
                    nw40_[i0_3_] = init3_(i0_3_)
                d_196_v77_ = nw40_
                index29_ = default__.safeIndex(809, (d_196_v77_).length(0))
                (d_196_v77_)[index29_] = d_89_v1_
                d_199_v78_: _dafny.Array
                nw41_ = _dafny.Array(False, 23)
                d_199_v78_ = nw41_
                index30_ = default__.safeIndex(965, (d_199_v78_).length(0))
                (d_199_v78_)[index30_] = not(d_182_v63_)
                index31_ = default__.safeIndex(809, (d_196_v77_).length(0))
                index32_ = default__.safeIndex(965, (d_199_v78_).length(0))
                rhs7_ = d_89_v1_
                rhs8_ = (not(d_182_v63_)) and (d_182_v63_)
                lhs4_ = d_196_v77_
                lhs5_ = default__.safeIndex(809, (d_196_v77_).length(0))
                lhs6_ = d_199_v78_
                lhs7_ = default__.safeIndex(965, (d_199_v78_).length(0))
                lhs4_[lhs5_] = rhs7_
                lhs6_[lhs7_] = rhs8_
            d_200_v79_: _dafny.Seq
            d_200_v79_ = _dafny.SeqWithoutIsStrInference([d_182_v63_])
            d_182_v63_ = not (default__.fm2(d_182_v63_, d_125_v22_, len(d_90_v2_), default__.fm2(d_182_v63_, d_125_v22_, len(d_200_v79_), d_182_v63_, d_91_globalState_), d_91_globalState_)) or (not(not(d_182_v63_)))
        elif True:
            index33_ = default__.safeIndex(18, (d_126_v23_).length(0))
            (d_126_v23_)[index33_] = (d_89_v1_)[default__.safeIndex(d_88_v0_, len(d_89_v1_))]
            index34_ = default__.safeIndex(18, (d_126_v23_).length(0))
            rhs9_ = False
            rhs10_ = d_88_v0_
            rhs11_ = d_125_v22_
            rhs12_ = (d_88_v0_ if d_182_v63_ else d_88_v0_)
            rhs13_ = True
            lhs8_ = d_126_v23_
            lhs9_ = default__.safeIndex(18, (d_126_v23_).length(0))
            d_182_v63_ = rhs9_
            d_88_v0_ = rhs10_
            lhs8_[lhs9_] = rhs11_
            d_88_v0_ = rhs12_
            d_182_v63_ = rhs13_
            d_201_v80_: _dafny.Array
            def lambda7_(d_202_v2_, d_203_v0_):
                def lambda8_(d_204_i3_):
                    return (d_202_v2_) + (_dafny.SeqWithoutIsStrInference([d_203_v0_, d_203_v0_, d_203_v0_]))

                return lambda8_

            init4_ = lambda7_(d_90_v2_, d_88_v0_)
            nw42_ = _dafny.Array(None, 28)
            for i0_4_ in range(nw42_.length(0)):
                nw42_[i0_4_] = init4_(i0_4_)
            d_201_v80_ = nw42_
            d_201_v80_ = d_201_v80_
            d_205_v81_: _dafny.Map
            d_205_v81_ = _dafny.Map({d_90_v2_: d_89_v1_})
            d_205_v81_ = d_205_v81_
            d_182_v63_ = d_182_v63_
            index35_ = default__.safeIndex(18, (d_126_v23_).length(0))
            (d_126_v23_)[index35_] = (d_126_v23_)[default__.safeIndex(18, (d_126_v23_).length(0))]
        default__.m0(d_182_v63_, (d_182_v63_) == (d_182_v63_), d_91_globalState_)
        d_182_v63_ = d_182_v63_
        d_182_v63_ = d_182_v63_
        d_206_v82_: _dafny.Array
        def lambda9_(d_207_v2_):
            def lambda10_(d_208_i4_):
                return d_207_v2_

            return lambda10_

        init5_ = lambda9_(d_90_v2_)
        nw43_ = _dafny.Array(None, 1)
        for i0_5_ in range(nw43_.length(0)):
            nw43_[i0_5_] = init5_(i0_5_)
        d_206_v82_ = nw43_
        d_209_v83_: _dafny.MultiSet
        d_209_v83_ = _dafny.MultiSet([d_182_v63_, d_182_v63_])
        d_210_v84_: _dafny.Map
        d_210_v84_ = _dafny.Map({d_206_v82_: d_209_v83_})
        d_211_v85_: _dafny.Seq
        d_211_v85_ = _dafny.SeqWithoutIsStrInference([((d_210_v84_)[d_206_v82_] if (d_206_v82_) in (d_210_v84_) else d_209_v83_), d_209_v83_, d_209_v83_])
        d_211_v85_ = d_211_v85_
        d_212_v86_: C0
        nw44_ = C0()
        nw44_.ctor__(d_127_v24_)
        d_212_v86_ = nw44_
        default__.m0(d_182_v63_, d_182_v63_, d_91_globalState_)
        d_213_v87_: _dafny.Array
        def lambda11_(d_214_v2_, d_215_v0_):
            def lambda12_(d_216_i5_):
                return default__.safeModuloInt(d_216_i5_, (d_214_v2_)[default__.safeIndex(d_215_v0_, len(d_214_v2_))])

            return lambda12_

        init6_ = lambda11_(d_90_v2_, d_88_v0_)
        nw45_ = _dafny.Array(None, 16)
        for i0_6_ in range(nw45_.length(0)):
            nw45_[i0_6_] = init6_(i0_6_)
        d_213_v87_ = nw45_
        index36_ = default__.safeIndex(544, (d_213_v87_).length(0))
        (d_213_v87_)[index36_] = d_88_v0_
        d_217_v88_: _dafny.Map
        d_217_v88_ = _dafny.Map({d_182_v63_: 211})
        d_218_v89_: _dafny.Set
        d_218_v89_ = _dafny.Set({len(d_90_v2_)})
        index37_ = default__.safeIndex(544, (d_213_v87_).length(0))
        (d_213_v87_)[index37_] = (0) - (default__.safeModuloInt(((d_217_v88_)[d_182_v63_] if (d_182_v63_) in (d_217_v88_) else d_88_v0_), len(_dafny.Map({len(d_218_v89_): (d_128_v25_).fm1(False, d_182_v63_, not(d_182_v63_), not(d_182_v63_), d_91_globalState_)}))))
        index38_ = default__.safeIndex(544, (d_213_v87_).length(0))
        (d_213_v87_)[index38_] = (0) - (default__.safeModuloInt(d_88_v0_, (d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))]))
        index39_ = default__.safeIndex(544, (d_213_v87_).length(0))
        (d_213_v87_)[index39_] = (d_88_v0_) - ((d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))])
        d_219_v90_: _dafny.Array
        nw46_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
        d_219_v90_ = nw46_
        index40_ = default__.safeIndex(918, (d_219_v90_).length(0))
        (d_219_v90_)[index40_] = d_89_v1_
        index41_ = default__.safeIndex(918, (d_219_v90_).length(0))
        (d_219_v90_)[index41_] = d_89_v1_
        d_220_v91_: _dafny.Map
        d_220_v91_ = _dafny.Map({(d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))]: (d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))]})
        def iife10_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in (d_90_v2_).Elements:
                d_221_v92_: int = compr_6_
                if (d_221_v92_) in (d_90_v2_):
                    coll6_[(d_221_v92_) * (d_88_v0_)] = (d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))]
            return _dafny.Map(coll6_)
        d_220_v91_ = (iife10_()
         if d_182_v63_ else d_220_v91_)
        d_222_v93_: _dafny.Set
        d_222_v93_ = _dafny.Set({d_209_v83_, (d_209_v83_).set(d_182_v63_, default__.abs((d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))])), _dafny.MultiSet([d_182_v63_]), d_209_v83_})
        if (_dafny.Set({d_209_v83_, d_209_v83_})).isdisjoint(d_222_v93_):
            d_223_v94_: _dafny.Seq
            d_223_v94_ = _dafny.SeqWithoutIsStrInference([d_89_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqfejscgi")), d_89_v1_, _dafny.SeqWithoutIsStrInference([d_125_v22_ for d_224_i6_ in range(default__.abs(-358))])])
            index42_ = default__.safeIndex(918, (d_219_v90_).length(0))
            index43_ = default__.safeIndex(544, (d_213_v87_).length(0))
            rhs14_ = ((d_223_v94_)[default__.safeIndex((d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))], len(d_223_v94_))]) + ((d_89_v1_) + ((d_219_v90_)[default__.safeIndex(918, (d_219_v90_).length(0))]))
            rhs15_ = ((d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))]) - ((d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))])
            lhs10_ = d_219_v90_
            lhs11_ = default__.safeIndex(918, (d_219_v90_).length(0))
            lhs12_ = d_213_v87_
            lhs13_ = default__.safeIndex(544, (d_213_v87_).length(0))
            lhs10_[lhs11_] = rhs14_
            lhs12_[lhs13_] = rhs15_
            d_225_v95_: _dafny.MultiSet
            d_225_v95_ = _dafny.MultiSet([d_88_v0_])
            d_226_v97_: _dafny.Map
            d_226_v97_ = _dafny.Map({d_225_v95_: d_218_v89_})
            def iife11_():
                coll7_ = _dafny.Set()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(-62, 728):
                    d_227_v96_: int = compr_7_
                    if ((-62) <= (d_227_v96_)) and ((d_227_v96_) < (728)):
                        coll7_ = coll7_.union(_dafny.Set([(d_227_v96_) - (d_88_v0_)]))
                return _dafny.Set(coll7_)
            default__.m0((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_212_v86_).fm1(d_182_v63_, d_182_v63_, d_182_v63_, d_182_v63_, d_91_globalState_), d_88_v0_]))).issubset(d_225_v95_), (iife11_()
            ).isdisjoint(((d_226_v97_)[_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbacwiua")))])] if (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbacwiua")))])) in (d_226_v97_) else d_218_v89_)), d_91_globalState_)
            index44_ = default__.safeIndex(544, (d_213_v87_).length(0))
            (d_213_v87_)[index44_] = (0) - (d_88_v0_)
            d_228_v98_: D2
            d_228_v98_ = D2_DC7((d_219_v90_)[default__.safeIndex(918, (d_219_v90_).length(0))], len(d_89_v1_), (d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))], d_182_v63_)
            d_229_v99_: _dafny.MultiSet
            d_229_v99_ = _dafny.MultiSet([(d_228_v98_).cf11, d_89_v1_, (d_219_v90_)[default__.safeIndex(918, (d_219_v90_).length(0))]])
            d_230_v100_: _dafny.Map
            d_230_v100_ = _dafny.Map({d_229_v99_: d_182_v63_})
            d_230_v100_ = (d_230_v100_).set((_dafny.MultiSet([d_89_v1_])) - (d_229_v99_), False)
            d_231_v101_: _dafny.Map
            d_231_v101_ = _dafny.Map({(d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))]: d_182_v63_})
            d_231_v101_ = _dafny.Map({(d_212_v86_).fm1(d_182_v63_, d_182_v63_, d_182_v63_, d_182_v63_, d_91_globalState_): d_182_v63_})
        elif True:
            d_182_v63_ = d_182_v63_
            d_232_v102_: T0
            nw47_ = C0()
            nw47_.ctor__((d_127_v24_) + (d_127_v24_))
            d_232_v102_ = nw47_
            d_233_v103_: _dafny.Seq
            d_233_v103_ = _dafny.SeqWithoutIsStrInference([d_182_v63_, d_182_v63_, False, d_182_v63_])
            d_234_v104_: D2
            d_234_v104_ = D2_DC7(((d_219_v90_)[default__.safeIndex(918, (d_219_v90_).length(0))]).set(default__.safeIndex(d_88_v0_, len((d_219_v90_)[default__.safeIndex(918, (d_219_v90_).length(0))])), d_125_v22_), d_88_v0_, (0) - (len(_dafny.Set({d_182_v63_}))), d_182_v63_)
            d_235_v105_: _dafny.Map
            d_235_v105_ = _dafny.Map({len(d_89_v1_): d_182_v63_})
            d_236_v106_: D1
            d_236_v106_ = D1_DC4(d_213_v87_)
            d_237_v107_: _dafny.Seq
            d_237_v107_ = _dafny.SeqWithoutIsStrInference([d_213_v87_, (d_236_v106_).cf8])
            rhs16_ = (d_90_v2_)[default__.safeIndex((d_232_v102_).fm1(d_182_v63_, (d_233_v103_)[default__.safeIndex((d_234_v104_).cf13, len(d_233_v103_))], d_182_v63_, not(not(d_182_v63_)), d_91_globalState_), len(d_90_v2_))]
            rhs17_ = d_232_v102_
            rhs18_ = (not(((d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))]) > (len(d_237_v107_))) if not(default__.fm2(not(d_182_v63_), default__.fm3(len(d_235_v105_), d_91_globalState_), (d_213_v87_)[default__.safeIndex(544, (d_213_v87_).length(0))], d_182_v63_, d_91_globalState_)) else d_182_v63_)
            rhs19_ = d_232_v102_
            d_88_v0_ = rhs16_
            d_232_v102_ = rhs17_
            d_182_v63_ = rhs18_
            d_232_v102_ = rhs19_
            d_125_v22_ = (d_125_v22_ if (d_182_v63_) == (d_182_v63_) else d_125_v22_)
            d_128_v25_ = d_128_v25_
            d_220_v91_ = (d_220_v91_).set(d_88_v0_, (d_232_v102_).fm1(d_182_v63_, d_182_v63_, True, d_182_v63_, d_91_globalState_))
        _dafny.print(_dafny.string_of(d_88_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_89_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_v2_) == (_dafny.SeqWithoutIsStrInference([159, 159, 159, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_globalState_).f3) == (_dafny.Set({636, 485}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_93_v4_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_124_v21_)[9])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_125_v22_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v23_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_127_v24_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_129_v26_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_182_v63_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_206_v82_)[0]) == (_dafny.SeqWithoutIsStrInference([159, 159, 159, 8]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v83_) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_210_v84_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v85_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, True]), _dafny.MultiSet([True, True]), _dafny.MultiSet([True, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v87_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v88_) == (_dafny.Map({True: 211}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v89_) == (_dafny.Set({4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_219_v90_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_220_v91_) == (_dafny.Map({205110: 2580, 10320: 2580, 159: 1105}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_v93_) == (_dafny.Set({_dafny.MultiSet([True, True]), _dafny.MultiSet([True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True]), _dafny.MultiSet([True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0))
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

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

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


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D2_DC10)

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({self.cf11.VerbatimString(True)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({self.cf10.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC10(D2, NamedTuple('DC10', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC10({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC10) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC12()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC12(D3, NamedTuple('DC12', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(_dafny.Seq({}), _dafny.CodePoint('D'), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC14(D4, NamedTuple('DC14', [('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value

class GlobalState:
    def  __init__(self):
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f2: int = int(0)
        self._f3: _dafny.Set = _dafny.Set({})
        self._f4: int = int(0)
        self._f5: str = _dafny.CodePoint('D')
        self._f6: bool = False
        self._f7: int = int(0)
        self._f8: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8

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
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8

class C0(T0):
    def  __init__(self):
        self._f9: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    def ctor__(self, f9):
        (self).f9 = f9

    def fm1(self, p0, p1, p2, p3, globalState):
        return (len((_dafny.SeqWithoutIsStrInference([141 for d_238_i0_ in range(default__.abs(134))])) + (_dafny.SeqWithoutIsStrInference([81 for d_239_i1_ in range(default__.abs(-967))])))) - ((0) - (len((_dafny.SeqWithoutIsStrInference([True, not(False), (D0_DC1(True, 486, len(_dafny.SeqWithoutIsStrInference([True, True])), True, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([308]))).cardinality)).cf4])) + (_dafny.SeqWithoutIsStrInference([True])))))

