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
        source0_ = D0_DC1(801, (_dafny.MultiSet([672, 917, (_dafny.MultiSet([False])).cardinality])).cardinality, 458)
        if source0_.is_DC1:
            d_0___mcc_h0_ = source0_.cf1
            d_1___mcc_h1_ = source0_.cf2
            d_2___mcc_h2_ = source0_.cf3
            d_3_cf3_ = d_2___mcc_h2_
            d_4_cf2_ = d_1___mcc_h1_
            d_5_cf1_ = d_0___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bielrkk"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))
        elif source0_.is_DC2:
            d_6___mcc_h3_ = source0_.cf4
            d_7___mcc_h4_ = source0_.cf5
            d_8___mcc_h5_ = source0_.cf6
            d_9_cf6_ = d_8___mcc_h5_
            d_10_cf5_ = d_7___mcc_h4_
            d_11_cf4_ = d_6___mcc_h3_
            return not(False)
        elif source0_.is_DC0:
            d_12___mcc_h6_ = source0_.cf0
            d_13_cf0_ = d_12___mcc_h6_
            return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_14_i0_ in range(default__.abs(263))]))) < (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phiihuc"))))
        elif True:
            d_15___mcc_h7_ = source0_.cf7
            d_16_cf7_ = d_15___mcc_h7_
            return True

    @staticmethod
    def fm1(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('p') if True else _dafny.CodePoint('n')) for d_17_i0_ in range(default__.abs(-879))])

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        return ((D5_DC13(_dafny.Map({False: _dafny.Map({641: False})})) if True else D5_DC13(_dafny.Map({True: _dafny.Map({219: False})})))).cf20

    @staticmethod
    def fm3(p0, p1, globalState):
        if not (True) or (False):
            if False:
                return _dafny.CodePoint('j')
            elif True:
                return _dafny.CodePoint('g')
        elif True:
            return _dafny.CodePoint('y')

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        return ((_dafny.Set({True})) | (_dafny.Set({not(not(not(not(True))))}))) | (_dafny.Set({False}))

    @staticmethod
    def fm5(p0, globalState):
        return 521

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return (_dafny.Map({not(not((D1_DC5(not(False), not(True), len(_dafny.Map({len(_dafny.Set({True})): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))).cardinality})))).cf10)): True})) | ((_dafny.Map({False: False}) if False else _dafny.Map({True: not(True)})))

    @staticmethod
    def fm7(p0, globalState):
        return _dafny.Map({(523) - (len(_dafny.Set({185}))): len(_dafny.Set({(0) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, not(True), not(False)])])).cardinality)}))})

    @staticmethod
    def fm8(p0, globalState):
        source1_ = D0_DC0(False)
        if source1_.is_DC1:
            d_18___mcc_h0_ = source1_.cf1
            d_19___mcc_h1_ = source1_.cf2
            d_20___mcc_h2_ = source1_.cf3
            d_21_cf3_ = d_20___mcc_h2_
            d_22_cf2_ = d_19___mcc_h1_
            d_23_cf1_ = d_18___mcc_h0_
            def iife0_():
                coll0_ = _dafny.Set()
                compr_0_: D0
                for compr_0_ in (_dafny.SeqWithoutIsStrInference([D0_DC0(True), D0_DC0(True), D0_DC0(False)])).Elements:
                    d_24_v0_: D0 = compr_0_
                    if (d_24_v0_) in (_dafny.SeqWithoutIsStrInference([D0_DC0(True), D0_DC0(True), D0_DC0(False)])):
                        coll0_ = coll0_.union(_dafny.Set([d_24_v0_]))
                return _dafny.Set(coll0_)
            return iife0_()
            
        elif source1_.is_DC2:
            d_25___mcc_h3_ = source1_.cf4
            d_26___mcc_h4_ = source1_.cf5
            d_27___mcc_h5_ = source1_.cf6
            d_28_cf6_ = d_27___mcc_h5_
            d_29_cf5_ = d_26___mcc_h4_
            d_30_cf4_ = d_25___mcc_h3_
            if not(False):
                def iife1_():
                    coll1_ = _dafny.Set()
                    compr_1_: D0
                    for compr_1_ in (_dafny.Set({D0_DC0(True)})).Elements:
                        d_31_v1_: D0 = compr_1_
                        if (d_31_v1_) in (_dafny.Set({D0_DC0(True)})):
                            coll1_ = coll1_.union(_dafny.Set([d_31_v1_]))
                    return _dafny.Set(coll1_)
                return iife1_()
                
            elif True:
                def iife2_():
                    coll2_ = _dafny.Set()
                    compr_2_: D0
                    for compr_2_ in (_dafny.SeqWithoutIsStrInference([D0_DC0(False), D0_DC0(False)])).Elements:
                        d_32_v2_: D0 = compr_2_
                        if (d_32_v2_) in (_dafny.SeqWithoutIsStrInference([D0_DC0(False), D0_DC0(False)])):
                            coll2_ = coll2_.union(_dafny.Set([d_32_v2_]))
                    return _dafny.Set(coll2_)
                return iife2_()
                
        elif source1_.is_DC0:
            d_33___mcc_h6_ = source1_.cf0
            d_34_cf0_ = d_33___mcc_h6_
            def iife3_():
                coll3_ = _dafny.Set()
                compr_3_: D0
                for compr_3_ in (_dafny.SeqWithoutIsStrInference([D0_DC0(d_34_cf0_) for d_35_i0_ in range(default__.abs(226))])).Elements:
                    d_36_v3_: D0 = compr_3_
                    if (d_36_v3_) in (_dafny.SeqWithoutIsStrInference([D0_DC0(d_34_cf0_) for d_35_i0_ in range(default__.abs(226))])):
                        coll3_ = coll3_.union(_dafny.Set([d_36_v3_]))
                return _dafny.Set(coll3_)
            return iife3_()
            
        elif True:
            d_37___mcc_h7_ = source1_.cf7
            d_38_cf7_ = d_37___mcc_h7_
            return _dafny.Set({D0_DC0(False), D0_DC0(not(True))})

    @staticmethod
    def fm9(globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: str
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_39_i0_ in range(default__.abs(732))])).Elements:
                d_40_v0_: str = compr_4_
                if (d_40_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_39_i0_ in range(default__.abs(732))])):
                    coll4_[d_40_v0_] = (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.Set({D5_DC13(_dafny.Map({True: _dafny.Map({808: False})}))}))])).cardinality, 720])).cardinality
            return _dafny.Map(coll4_)
        return (_dafny.Map({_dafny.CodePoint('x'): 697})) | (iife4_()
        )

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: str
            for compr_5_ in (_dafny.Set({_dafny.CodePoint('t')})).Elements:
                d_41_v0_: str = compr_5_
                if (d_41_v0_) in (_dafny.Set({_dafny.CodePoint('t')})):
                    coll5_[d_41_v0_] = 357
            return _dafny.Map(coll5_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('r'): 498})])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('j'): 330})])) + (_dafny.SeqWithoutIsStrInference([iife5_()
 for d_42_i0_ in range(default__.abs(515))])))

    @staticmethod
    def fm11(p0, globalState):
        return D2_DC8((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_43_i0_ in range(default__.abs(328))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_44_i1_ in range(default__.abs(387))])), (0) - (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('x'), _dafny.CodePoint('c'), _dafny.CodePoint('u')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('f')])))))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yo")))])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-126])))

    @staticmethod
    def fm13(p0, globalState):
        source2_ = D1_DC6()
        if source2_.is_DC5:
            d_45___mcc_h0_ = source2_.cf9
            d_46___mcc_h1_ = source2_.cf10
            d_47___mcc_h2_ = source2_.cf11
            d_48_cf11_ = d_47___mcc_h2_
            d_49_cf10_ = d_46___mcc_h1_
            d_50_cf9_ = d_45___mcc_h0_
            if d_50_cf9_:
                return D1_DC6()
            elif True:
                return D1_DC6()
        elif source2_.is_DC6:
            return D1_DC6()
        elif True:
            d_51___mcc_h3_ = source2_.cf8
            d_52_cf8_ = d_51___mcc_h3_
            return D1_DC6()

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return D4_DC10(_dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "givcgweu")), -147, 324))]))

    @staticmethod
    def Main(noArgsParameter__):
        d_53_globalState_: GlobalState
        nw0_ = GlobalState()
        nw0_.ctor__(825, 1, True)
        d_53_globalState_ = nw0_
        d_54_v0_: bool
        d_54_v0_ = True
        d_55_v1_: _dafny.Array
        nw1_ = _dafny.Array(int(0), 22)
        d_55_v1_ = nw1_
        d_56_v2_: int
        d_56_v2_ = 891
        d_57_v3_: _dafny.Map
        d_57_v3_ = _dafny.Map({(d_55_v1_ if d_54_v0_ else d_55_v1_): (0) - (d_56_v2_)})
        d_57_v3_ = d_57_v3_
        d_54_v0_ = not(default__.fm0(d_56_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "if"))), d_53_globalState_))
        d_58_v4_: _dafny.Array
        def lambda0_(d_59_v2_, d_60_v0_):
            def lambda1_(d_61_i0_):
                return (_dafny.Map({d_59_v2_: d_60_v0_})) | (_dafny.Map({d_59_v2_: d_60_v0_}))

            return lambda1_

        init0_ = lambda0_(d_56_v2_, d_54_v0_)
        nw2_ = _dafny.Array(None, 23)
        for i0_0_ in range(nw2_.length(0)):
            nw2_[i0_0_] = init0_(i0_0_)
        d_58_v4_ = nw2_
        d_62_v5_: _dafny.Map
        d_62_v5_ = _dafny.Map({d_56_v2_: d_54_v0_})
        index0_ = default__.safeIndex(111, (d_58_v4_).length(0))
        (d_58_v4_)[index0_] = d_62_v5_
        index1_ = default__.safeIndex(111, (d_58_v4_).length(0))
        (d_58_v4_)[index1_] = d_62_v5_
        d_63_v6_: _dafny.MultiSet
        d_63_v6_ = _dafny.MultiSet([False])
        d_54_v0_ = not(((d_56_v2_) - (d_56_v2_)) > ((d_56_v2_) + ((d_63_v6_).cardinality)))
        d_64_v7_: D0
        d_64_v7_ = D0_DC0(d_54_v0_)
        d_65_v8_: C0
        nw3_ = C0()
        nw3_.ctor__(d_64_v7_, _dafny.Map({d_54_v0_: d_55_v1_}), d_55_v1_, d_56_v2_)
        d_65_v8_ = nw3_
        d_66_v9_: str
        d_66_v9_ = _dafny.CodePoint('q')
        d_67_v10_: T0
        nw4_ = C0()
        nw4_.ctor__((d_65_v8_).f7, (d_65_v8_).f8, d_55_v1_, d_56_v2_)
        d_67_v10_ = nw4_
        d_68_v11_: _dafny.Map
        d_68_v11_ = _dafny.Map({_dafny.CodePoint('w'): d_67_v10_})
        d_69_v12_: _dafny.Map
        d_69_v12_ = _dafny.Map({d_62_v5_: not((d_66_v9_) not in (d_68_v11_))})
        d_69_v12_ = d_69_v12_
        d_70_v13_: _dafny.Array
        nw5_ = _dafny.Array(_dafny.Set({}), 16)
        d_70_v13_ = nw5_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_70_v13_).length(0)):
            d_71_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_71_i1_)) and ((d_71_i1_) < ((d_70_v13_).length(0)))):
                (d_70_v13_)[(d_71_i1_)] = ((_dafny.Set({False, d_54_v0_})).intersection(_dafny.Set({d_54_v0_, True, d_54_v0_}))) | (_dafny.Set({d_54_v0_}))
        rhs0_ = 551
        rhs1_ = -189
        lhs0_ = d_53_globalState_
        lhs0_.f1 = rhs0_
        d_56_v2_ = rhs1_
        d_54_v0_ = d_54_v0_
        index2_ = default__.safeIndex(634, (d_55_v1_).length(0))
        (d_55_v1_)[index2_] = d_56_v2_
        index3_ = default__.safeIndex(634, (d_55_v1_).length(0))
        (d_55_v1_)[index3_] = ((d_63_v6_)[(d_56_v2_) < (d_56_v2_)] if ((d_56_v2_) < (d_56_v2_)) in (d_63_v6_) else (d_67_v10_).f4)
        d_72_v14_: _dafny.Array
        nw6_ = _dafny.Array(None, 7)
        d_72_v14_ = nw6_
        d_73_v15_: _dafny.Map
        d_73_v15_ = _dafny.Map({d_72_v14_: d_54_v0_})
        d_74_v16_: _dafny.Map
        d_74_v16_ = _dafny.Map({(d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]: (d_67_v10_).f4})
        d_75_v17_: _dafny.Seq
        d_75_v17_ = _dafny.SeqWithoutIsStrInference([d_54_v0_])
        d_76_v18_: D0
        d_76_v18_ = D0_DC1(len(d_75_v17_), (d_67_v10_).f4, len(_dafny.SeqWithoutIsStrInference([d_56_v2_])))
        d_77_v19_: _dafny.Map
        d_77_v19_ = _dafny.Map({d_67_v10_: d_76_v18_})
        d_73_v15_ = (d_73_v15_).set(d_72_v14_, default__.fm0(len(d_74_v16_), len(d_77_v19_), d_53_globalState_))
        d_78_v20_: _dafny.Array
        nw7_ = _dafny.Array(None, 25)
        d_78_v20_ = nw7_
        d_79_v21_: _dafny.Map
        d_79_v21_ = _dafny.Map({d_67_v10_: d_78_v20_})
        d_79_v21_ = (d_79_v21_).set(d_67_v10_, d_78_v20_)
        d_80_v22_: D1
        d_80_v22_ = D1_DC4(d_55_v1_)
        d_80_v22_ = d_80_v22_
        d_81_v23_: _dafny.Seq
        d_81_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpuqdrl"))
        d_82_v24_: _dafny.Array
        nw8_ = _dafny.Array(False, 2)
        d_82_v24_ = nw8_
        d_83_v25_: C1
        nw9_ = C1()
        nw9_.ctor__((_dafny.CodePoint('y')), d_82_v24_, (d_67_v10_).f3, d_56_v2_)
        d_83_v25_ = nw9_
        rhs2_ = (d_81_v23_) + (d_81_v23_)
        rhs3_ = d_83_v25_
        d_81_v23_ = rhs2_
        d_83_v25_ = rhs3_
        d_84_v26_: D0
        d_84_v26_ = D0_DC3(D0_DC2(d_81_v23_, default__.fm5(False, d_53_globalState_), d_56_v2_))
        source3_ = d_84_v26_
        if source3_.is_DC1:
            d_85___mcc_h0_ = source3_.cf1
            d_86___mcc_h1_ = source3_.cf2
            d_87___mcc_h2_ = source3_.cf3
            d_88_cf3_ = d_87___mcc_h2_
            d_89_cf2_ = d_86___mcc_h1_
            d_90_cf1_ = d_85___mcc_h0_
            d_91_v27_: _dafny.MultiSet
            d_91_v27_ = _dafny.MultiSet([d_88_cf3_, (d_67_v10_).f4])
            if (d_91_v27_).issubset(_dafny.MultiSet([(d_67_v10_).f4])):
                d_92_v28_: _dafny.Map
                d_92_v28_ = _dafny.Map({d_88_cf3_: d_83_v25_})
                d_93_v29_: _dafny.Map
                d_93_v29_ = _dafny.Map({(d_92_v28_) | (_dafny.Map({(d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]: d_83_v25_})): not(d_54_v0_)})
                d_93_v29_ = d_93_v29_
                d_94_v30_: _dafny.Seq
                d_94_v30_ = _dafny.SeqWithoutIsStrInference([d_84_v26_])
                d_95_v31_: D4
                d_95_v31_ = D4_DC10(_dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC1((d_67_v10_).f4, (d_67_v10_).f4, len(_dafny.Map({d_54_v0_: (d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]})))) for d_96_i2_ in range(default__.abs(-50))]))
                d_97_v32_: _dafny.Map
                d_97_v32_ = _dafny.Map({d_54_v0_: d_90_cf1_})
                d_98_v33_: _dafny.Seq
                d_98_v33_ = _dafny.SeqWithoutIsStrInference([len(d_81_v23_)])
                d_99_v34_: _dafny.MultiSet
                d_99_v34_ = _dafny.MultiSet([(d_94_v30_) + (((d_95_v31_).cf16).set(default__.safeIndex(((d_97_v32_)[d_54_v0_] if (d_54_v0_) in (d_97_v32_) else len(d_98_v33_)), len((d_95_v31_).cf16)), d_84_v26_)), d_94_v30_])
                d_100_v35_: D2
                d_100_v35_ = D2_DC7(d_75_v17_)
                rhs4_ = d_54_v0_
                rhs5_ = (d_67_v10_).f4
                rhs6_ = d_99_v34_
                rhs7_ = _dafny.MultiSet((d_100_v35_).cf12)
                d_54_v0_ = rhs4_
                d_89_cf2_ = rhs5_
                d_99_v34_ = rhs6_
                d_63_v6_ = rhs7_
                d_54_v0_ = d_54_v0_
                d_101_v36_: bool
                d_102_v37_: bool
                d_103_v38_: bool
                out0_: bool
                out1_: bool
                out2_: bool
                out0_, out1_, out2_ = (d_83_v25_).m0(d_53_globalState_)
                d_101_v36_ = out0_
                d_102_v37_ = out1_
                d_103_v38_ = out2_
                d_103_v38_ = d_101_v36_
            elif True:
                d_104_v39_: _dafny.Seq
                d_104_v39_ = _dafny.SeqWithoutIsStrInference([d_67_v10_, d_67_v10_])
                d_67_v10_ = ((d_104_v39_)[default__.safeIndex(-846, len(d_104_v39_))] if d_54_v0_ else d_67_v10_)
                d_105_v40_: _dafny.Seq
                d_105_v40_ = _dafny.SeqWithoutIsStrInference([(d_91_v27_).cardinality])
                d_106_v41_: _dafny.Seq
                d_106_v41_ = _dafny.SeqWithoutIsStrInference([d_91_v27_, _dafny.MultiSet([len(d_105_v40_)])])
                d_107_v42_: _dafny.Map
                d_107_v42_ = _dafny.Map({(0) - (d_89_cf2_): (d_106_v41_)[default__.safeIndex(737, len(d_106_v41_))]})
                d_91_v27_ = ((d_107_v42_)[855] if (855) in (d_107_v42_) else d_91_v27_)
                d_108_v43_: D2
                d_108_v43_ = D2_DC8(d_81_v23_, d_90_cf1_)
                pat_let_tv0_ = d_81_v23_
                index4_ = default__.safeIndex(634, (d_55_v1_).length(0))
                def iife6_(_pat_let0_0):
                    def iife7_(d_109_dt__update__tmp_h0_):
                        def iife8_(_pat_let1_0):
                            def iife9_(d_110_dt__update_hcf13_h0_):
                                return D2_DC8(d_110_dt__update_hcf13_h0_, (d_109_dt__update__tmp_h0_).cf14)
                            return iife9_(_pat_let1_0)
                        return iife8_(pat_let_tv0_)
                    return iife7_(_pat_let0_0)
                rhs8_ = default__.fm0((d_67_v10_).f4, d_88_cf3_, d_53_globalState_)
                rhs9_ = iife6_(default__.fm11(d_54_v0_, d_53_globalState_))
                rhs10_ = default__.fm3(d_66_v9_, d_90_cf1_, d_53_globalState_)
                rhs11_ = (d_67_v10_).f4
                lhs1_ = d_55_v1_
                lhs2_ = default__.safeIndex(634, (d_55_v1_).length(0))
                d_54_v0_ = rhs8_
                d_108_v43_ = rhs9_
                d_66_v9_ = rhs10_
                lhs1_[lhs2_] = rhs11_
                d_82_v24_ = (d_83_v25_).f6
                index5_ = default__.safeIndex(634, (d_55_v1_).length(0))
                (d_55_v1_)[index5_] = (0) - ((d_67_v10_).f4)
            index6_ = default__.safeIndex(634, (d_55_v1_).length(0))
            (d_55_v1_)[index6_] = ((d_67_v10_).f4) - ((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))])
            if (d_54_v0_) not in (((_dafny.SeqWithoutIsStrInference([d_54_v0_])).set(default__.safeIndex(d_56_v2_, len(_dafny.SeqWithoutIsStrInference([d_54_v0_]))), d_54_v0_)).set(default__.safeIndex((d_67_v10_).f4, len((_dafny.SeqWithoutIsStrInference([d_54_v0_])).set(default__.safeIndex(d_56_v2_, len(_dafny.SeqWithoutIsStrInference([d_54_v0_]))), d_54_v0_))), d_54_v0_)):
                index7_ = default__.safeIndex(631, ((d_83_v25_).f6).length(0))
                ((d_83_v25_).f6)[index7_] = d_54_v0_
                index8_ = default__.safeIndex(631, ((d_83_v25_).f6).length(0))
                ((d_83_v25_).f6)[index8_] = default__.fm0(d_89_cf2_, ((d_67_v10_).f4) - (d_89_cf2_), d_53_globalState_)
                d_74_v16_ = (d_74_v16_).set(len((d_58_v4_)[default__.safeIndex(111, (d_58_v4_).length(0))]), d_90_cf1_)
                index9_ = default__.safeIndex(634, (d_55_v1_).length(0))
                (d_55_v1_)[index9_] = 656
                d_111_v44_: _dafny.Seq
                out3_: _dafny.Seq
                out3_ = (d_83_v25_).m1(False, d_56_v2_, default__.fm5(not(d_54_v0_), d_53_globalState_), d_53_globalState_)
                d_111_v44_ = out3_
                d_112_v45_: _dafny.MultiSet
                d_112_v45_ = _dafny.MultiSet([(d_67_v10_).f3])
                index10_ = default__.safeIndex(631, ((d_83_v25_).f6).length(0))
                index11_ = default__.safeIndex(631, ((d_83_v25_).f6).length(0))
                index12_ = default__.safeIndex(634, (d_55_v1_).length(0))
                rhs12_ = not(default__.fm0(len(d_81_v23_), (d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], d_53_globalState_))
                rhs13_ = not(((d_67_v10_).f3) not in ((d_112_v45_) - (d_112_v45_)))
                rhs14_ = len(d_111_v44_)
                lhs3_ = (d_83_v25_).f6
                lhs4_ = default__.safeIndex(631, ((d_83_v25_).f6).length(0))
                lhs5_ = (d_83_v25_).f6
                lhs6_ = default__.safeIndex(631, ((d_83_v25_).f6).length(0))
                lhs7_ = d_55_v1_
                lhs8_ = default__.safeIndex(634, (d_55_v1_).length(0))
                lhs3_[lhs4_] = rhs12_
                lhs5_[lhs6_] = rhs13_
                lhs7_[lhs8_] = rhs14_
            elif True:
                index13_ = default__.safeIndex(634, (d_55_v1_).length(0))
                (d_55_v1_)[index13_] = default__.safeDivisionInt(d_88_cf3_, (d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))])
                d_54_v0_ = d_54_v0_
                d_113_v46_: _dafny.Seq
                d_113_v46_ = _dafny.SeqWithoutIsStrInference([d_81_v23_, d_81_v23_, d_81_v23_, d_81_v23_])
                d_114_v47_: _dafny.Map
                d_114_v47_ = _dafny.Map({d_54_v0_: (d_67_v10_).f4})
                d_115_v48_: _dafny.Array
                nw10_ = _dafny.Array(None, 18)
                nw10_[int(0)] = (_dafny.SeqWithoutIsStrInference([(d_83_v25_).f5 for d_116_i3_ in range(default__.abs(-279))]) if d_54_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avwjo")))
                nw10_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntbyypb"))
                nw10_[int(2)] = d_81_v23_
                nw10_[int(3)] = (d_113_v46_)[default__.safeIndex(((d_114_v47_)[d_54_v0_] if (d_54_v0_) in (d_114_v47_) else default__.fm5(d_54_v0_, d_53_globalState_)), len(d_113_v46_))]
                nw10_[int(4)] = default__.fm1(d_54_v0_, d_54_v0_, d_53_globalState_)
                nw10_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dddi"))
                nw10_[int(6)] = (d_81_v23_) + (d_81_v23_)
                nw10_[int(7)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnfds"))) + (d_81_v23_)
                nw10_[int(8)] = _dafny.SeqWithoutIsStrInference([(d_83_v25_).f5 for d_117_i4_ in range(default__.abs(-336))])
                nw10_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
                nw10_[int(10)] = (d_81_v23_) + (d_81_v23_)
                nw10_[int(11)] = (d_81_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddtjgujl")))
                nw10_[int(12)] = d_81_v23_
                nw10_[int(13)] = _dafny.SeqWithoutIsStrInference([(d_83_v25_).f5 for d_118_i5_ in range(default__.abs(308))])
                nw10_[int(14)] = (d_81_v23_) + (default__.fm1(d_54_v0_, not(d_54_v0_), d_53_globalState_))
                nw10_[int(15)] = d_81_v23_
                nw10_[int(16)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "my"))) + (d_81_v23_)
                nw10_[int(17)] = d_81_v23_
                d_115_v48_ = nw10_
                index14_ = default__.safeIndex(636, (d_115_v48_).length(0))
                (d_115_v48_)[index14_] = (d_81_v23_) + (d_81_v23_)
                index15_ = default__.safeIndex(636, (d_115_v48_).length(0))
                (d_115_v48_)[index15_] = ((d_81_v23_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_119_i6_ in range(default__.abs(237))]))).set(default__.safeIndex(d_89_cf2_, len((d_81_v23_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_120_i6_ in range(default__.abs(237))])))), (d_83_v25_).f5)
                d_67_v10_ = d_67_v10_
                d_121_v49_: _dafny.Seq
                d_121_v49_ = _dafny.SeqWithoutIsStrInference([d_90_cf1_])
                d_122_v50_: _dafny.Seq
                out4_: _dafny.Seq
                out4_ = (d_83_v25_).m1(d_54_v0_, (0) - ((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]), (d_121_v49_)[default__.safeIndex(d_89_cf2_, len(d_121_v49_))], d_53_globalState_)
                d_122_v50_ = out4_
            d_123_v51_: _dafny.Map
            d_123_v51_ = _dafny.Map({d_75_v17_: (d_81_v23_) != (_dafny.SeqWithoutIsStrInference([(d_81_v23_)[default__.safeIndex(d_88_cf3_, len(d_81_v23_))]]))})
            pat_let_tv1_ = d_54_v0_
            index16_ = default__.safeIndex(634, (d_55_v1_).length(0))
            def iife10_(_pat_let2_0):
                def iife11_(d_124_dt__update__tmp_h1_):
                    def iife12_(_pat_let3_0):
                        def iife13_(d_125_dt__update_hcf0_h0_):
                            return D0_DC0(d_125_dt__update_hcf0_h0_)
                        return iife13_(_pat_let3_0)
                    return iife12_(pat_let_tv1_)
                return iife11_(_pat_let2_0)
            rhs15_ = ((d_123_v51_)[_dafny.SeqWithoutIsStrInference([d_54_v0_])] if (_dafny.SeqWithoutIsStrInference([d_54_v0_])) in (d_123_v51_) else (iife10_(d_64_v7_)).cf0)
            rhs16_ = default__.fm5(d_54_v0_, d_53_globalState_)
            lhs9_ = d_55_v1_
            lhs10_ = default__.safeIndex(634, (d_55_v1_).length(0))
            d_54_v0_ = rhs15_
            lhs9_[lhs10_] = rhs16_
        elif source3_.is_DC2:
            d_126___mcc_h3_ = source3_.cf4
            d_127___mcc_h4_ = source3_.cf5
            d_128___mcc_h5_ = source3_.cf6
            d_129_cf6_ = d_128___mcc_h5_
            d_130_cf5_ = d_127___mcc_h4_
            d_131_cf4_ = d_126___mcc_h3_
            d_132_v52_: _dafny.Seq
            d_132_v52_ = _dafny.SeqWithoutIsStrInference([(d_67_v10_).f3])
            d_55_v1_ = (((d_65_v8_).f8)[not(d_54_v0_)] if (not(d_54_v0_)) in ((d_65_v8_).f8) else (d_132_v52_)[default__.safeIndex((0) - ((d_67_v10_).f4), len(d_132_v52_))])
            if d_54_v0_:
                d_133_v53_: C1
                nw11_ = C1()
                nw11_.ctor__((d_83_v25_).f5, (d_83_v25_).f6, (d_67_v10_).f3, len(d_75_v17_))
                d_133_v53_ = nw11_
                d_55_v1_ = (d_67_v10_).f3
                d_131_cf4_ = d_81_v23_
                d_134_v54_: _dafny.Array
                def lambda2_(d_135_v25_):
                    def lambda3_(d_136_i7_):
                        return _dafny.SeqWithoutIsStrInference([(d_135_v25_).f5])

                    return lambda3_

                init1_ = lambda2_(d_83_v25_)
                nw12_ = _dafny.Array(None, 27)
                for i0_1_ in range(nw12_.length(0)):
                    nw12_[i0_1_] = init1_(i0_1_)
                d_134_v54_ = nw12_
                rhs17_ = (0) - (d_130_cf5_)
                rhs18_ = d_134_v54_
                rhs19_ = (d_67_v10_).f4
                rhs20_ = d_54_v0_
                rhs21_ = ((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]) < ((d_67_v10_).f4)
                lhs11_ = d_53_globalState_
                lhs11_.f1 = rhs17_
                d_134_v54_ = rhs18_
                d_56_v2_ = rhs19_
                d_54_v0_ = rhs20_
                d_54_v0_ = rhs21_
                d_137_v55_: D2
                d_137_v55_ = D2_DC8(d_81_v23_, d_130_cf5_)
                d_138_v56_: _dafny.MultiSet
                d_138_v56_ = _dafny.MultiSet([(d_137_v55_).cf14])
                d_139_v57_: _dafny.Map
                d_139_v57_ = _dafny.Map({d_133_v53_: d_138_v56_})
                d_54_v0_ = not (d_54_v0_) or ((_dafny.Set({len(d_139_v57_)})).isdisjoint(_dafny.Set({(d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], d_130_cf5_})))
            elif True:
                d_67_v10_ = d_67_v10_
                d_75_v17_ = (d_75_v17_) + (d_75_v17_)
                d_140_v58_: bool
                d_141_v59_: bool
                d_142_v60_: bool
                out5_: bool
                out6_: bool
                out7_: bool
                out5_, out6_, out7_ = (d_83_v25_).m0(d_53_globalState_)
                d_140_v58_ = out5_
                d_141_v59_ = out6_
                d_142_v60_ = out7_
                d_142_v60_ = ((0) - (d_56_v2_)) == ((d_129_cf6_ if default__.fm0((d_67_v10_).f4, (d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], d_53_globalState_) else (d_67_v10_).f4))
                d_143_v61_: _dafny.Seq
                d_143_v61_ = _dafny.SeqWithoutIsStrInference([(d_67_v10_).f4, (d_67_v10_).f4, d_130_cf5_])
                d_144_v62_: _dafny.Map
                d_144_v62_ = _dafny.Map({d_143_v61_: d_129_cf6_})
                d_145_v64_: _dafny.Array
                nw13_ = _dafny.Array(None, 16)
                def iife14_():
                    coll6_ = _dafny.Set()
                    compr_6_: _dafny.Seq
                    for compr_6_ in (d_144_v62_).keys.Elements:
                        d_146_v63_: _dafny.Seq = compr_6_
                        if (d_146_v63_) in (d_144_v62_):
                            coll6_ = coll6_.union(_dafny.Set([d_146_v63_]))
                    return _dafny.Set(coll6_)
                nw13_[int(0)] = default__.fm3(d_66_v9_, len(iife14_()
                ), d_53_globalState_)
                nw13_[int(1)] = d_66_v9_
                nw13_[int(2)] = (d_83_v25_).f5
                nw13_[int(3)] = (d_83_v25_).f5
                nw13_[int(4)] = (d_83_v25_).f5
                nw13_[int(5)] = d_66_v9_
                nw13_[int(6)] = (d_83_v25_).f5
                nw13_[int(7)] = _dafny.CodePoint('q')
                nw13_[int(8)] = d_66_v9_
                nw13_[int(9)] = (d_83_v25_).f5
                nw13_[int(10)] = ((d_83_v25_).f5 if d_141_v59_ else _dafny.CodePoint('o'))
                nw13_[int(11)] = (d_83_v25_).f5
                nw13_[int(12)] = d_66_v9_
                nw13_[int(13)] = (d_83_v25_).f5
                nw13_[int(14)] = (d_83_v25_).f5
                nw13_[int(15)] = default__.fm3(_dafny.CodePoint('o'), d_56_v2_, d_53_globalState_)
                d_145_v64_ = nw13_
                index17_ = default__.safeIndex(179, (d_145_v64_).length(0))
                (d_145_v64_)[index17_] = _dafny.CodePoint('t')
                index18_ = default__.safeIndex(179, (d_145_v64_).length(0))
                (d_145_v64_)[index18_] = d_66_v9_
            if False:
                d_147_v65_: _dafny.Map
                d_147_v65_ = _dafny.Map({default__.fm0((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], d_130_cf5_, d_53_globalState_): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjmglnlgb")))})
                d_129_cf6_ = ((0) - (len(d_132_v52_)) if (d_54_v0_ if d_54_v0_ else True) else (((d_147_v65_)[d_54_v0_] if (d_54_v0_) in (d_147_v65_) else d_129_cf6_)) + ((d_67_v10_).f4))
                d_148_v66_: _dafny.Seq
                d_148_v66_ = _dafny.SeqWithoutIsStrInference([(d_67_v10_).f4, (d_67_v10_).f4])
                d_149_v67_: _dafny.MultiSet
                d_149_v67_ = _dafny.MultiSet([len(d_148_v66_)])
                d_54_v0_ = ((d_149_v67_) - (d_149_v67_)) != ((d_149_v67_) - (d_149_v67_))
                d_150_v68_: str
                d_150_v68_ = (d_83_v25_).f5
                d_151_v69_: _dafny.Map
                d_151_v69_ = _dafny.Map({d_150_v68_: True})
                d_151_v69_ = (d_151_v69_).set(d_66_v9_, d_54_v0_)
                d_54_v0_ = (default__.fm0(d_130_cf5_, (d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], d_53_globalState_)) == (d_54_v0_)
                d_152_v70_: D1
                d_152_v70_ = D1_DC6()
                d_153_v71_: _dafny.Map
                d_153_v71_ = _dafny.Map({(d_81_v23_)[default__.safeIndex((d_67_v10_).f4, len(d_81_v23_))]: d_65_v8_})
                d_154_v72_: _dafny.Seq
                d_154_v72_ = _dafny.SeqWithoutIsStrInference([d_65_v8_, d_65_v8_, ((d_153_v71_)[(d_83_v25_).f5] if ((d_83_v25_).f5) in (d_153_v71_) else d_65_v8_), d_65_v8_, d_65_v8_])
                rhs22_ = d_152_v70_
                rhs23_ = d_154_v72_
                rhs24_ = d_54_v0_
                d_152_v70_ = rhs22_
                d_154_v72_ = rhs23_
                d_54_v0_ = rhs24_
            elif True:
                d_155_v73_: C0
                nw14_ = C0()
                nw14_.ctor__(d_64_v7_, _dafny.Map({d_54_v0_: d_55_v1_}), d_55_v1_, (d_67_v10_).f4)
                d_155_v73_ = nw14_
                d_156_v74_: C0
                nw15_ = C0()
                nw15_.ctor__((d_65_v8_).f7, (d_65_v8_).f8, (d_67_v10_).f3, 931)
                d_156_v74_ = nw15_
                d_157_v75_: _dafny.Map
                d_157_v75_ = _dafny.Map({d_63_v6_: d_54_v0_})
                d_158_v76_: _dafny.Map
                d_158_v76_ = _dafny.Map({d_54_v0_: d_54_v0_})
                d_157_v75_ = (d_157_v75_).set(d_63_v6_, (len(d_158_v76_)) != ((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]))
                d_54_v0_ = not(default__.fm0((d_67_v10_).f4, d_129_cf6_, d_53_globalState_))
                d_159_v77_: C1
                nw16_ = C1()
                nw16_.ctor__(d_66_v9_, (d_83_v25_).f6, (d_132_v52_)[default__.safeIndex(d_56_v2_, len(d_132_v52_))], (d_67_v10_).f4)
                d_159_v77_ = nw16_
            d_81_v23_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mo"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_160_i8_ in range(default__.abs(-834))]))).set(default__.safeIndex((d_67_v10_).f4, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mo"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_161_i8_ in range(default__.abs(-834))])))), (d_83_v25_).f5)
        elif source3_.is_DC0:
            d_162___mcc_h6_ = source3_.cf0
            d_163_cf0_ = d_162___mcc_h6_
            d_164_v78_: C0
            nw17_ = C0()
            nw17_.ctor__((d_65_v8_).f7, ((d_65_v8_).f8) | (_dafny.Map({d_163_cf0_: (d_67_v10_).f3})), (d_67_v10_).f3, len(d_81_v23_))
            d_164_v78_ = nw17_
            d_165_v79_: _dafny.MultiSet
            d_165_v79_ = _dafny.MultiSet([-708, 399, d_56_v2_, (d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], 153])
            d_166_v80_: _dafny.Seq
            d_166_v80_ = _dafny.SeqWithoutIsStrInference([(d_67_v10_).f4])
            if ((_dafny.MultiSet(d_166_v80_)).intersection(d_165_v79_)).issubset(d_165_v79_):
                d_167_v81_: _dafny.Seq
                out8_: _dafny.Seq
                out8_ = (d_83_v25_).m1(d_54_v0_, (d_67_v10_).f4, (0) - ((d_67_v10_).f4), d_53_globalState_)
                d_167_v81_ = out8_
                (d_53_globalState_).f1 = 90
                d_168_v82_: _dafny.Seq
                d_168_v82_ = _dafny.SeqWithoutIsStrInference([d_84_v26_])
                d_169_v83_: D4
                d_169_v83_ = D4_DC10(d_168_v82_)
                d_170_v84_: _dafny.Seq
                d_170_v84_ = _dafny.SeqWithoutIsStrInference([d_169_v83_, D4_DC10(d_168_v82_), d_169_v83_])
                d_171_v85_: _dafny.Map
                d_171_v85_ = _dafny.Map({d_165_v79_: (d_170_v84_) + (d_170_v84_)})
                d_171_v85_ = (d_171_v85_).set(default__.fm12((default__.fm12(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijxylmpsk"))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dk"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijxylmpsk")))), (d_83_v25_).f5)), (d_67_v10_).f4, (d_67_v10_).f4, (d_167_v81_).set(default__.safeIndex((d_67_v10_).f4, len(d_167_v81_)), (d_83_v25_).f5), d_53_globalState_)).cardinality, d_56_v2_, (d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], d_167_v81_, d_53_globalState_), _dafny.SeqWithoutIsStrInference([d_169_v83_]))
                d_172_v86_: D0
                d_172_v86_ = D0_DC2(d_81_v23_, d_56_v2_, 785)
                rhs25_ = d_56_v2_
                rhs26_ = d_172_v86_
                lhs12_ = d_53_globalState_
                lhs12_.f1 = rhs25_
                d_172_v86_ = rhs26_
                d_163_cf0_ = d_54_v0_
            elif True:
                (d_53_globalState_).f1 = default__.safeModuloInt((default__.fm5(d_54_v0_, d_53_globalState_)) + ((d_67_v10_).f4), default__.safeDivisionInt((d_67_v10_).f4, (d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]))
                d_173_v87_: _dafny.Seq
                d_173_v87_ = _dafny.SeqWithoutIsStrInference([d_81_v23_, d_81_v23_, (d_81_v23_).set(default__.safeIndex((d_67_v10_).f4, len(d_81_v23_)), d_66_v9_)])
                d_81_v23_ = (d_173_v87_)[default__.safeIndex((d_67_v10_).f4, len(d_173_v87_))]
                d_55_v1_ = (d_67_v10_).f3
                d_174_v88_: bool
                d_175_v89_: bool
                d_176_v90_: bool
                out9_: bool
                out10_: bool
                out11_: bool
                out9_, out10_, out11_ = (d_83_v25_).m0(d_53_globalState_)
                d_174_v88_ = out9_
                d_175_v89_ = out10_
                d_176_v90_ = out11_
                d_177_v91_: D1
                d_177_v91_ = D1_DC6()
                d_178_v92_: _dafny.Array
                nw18_ = _dafny.Array(None, 18)
                nw18_[int(0)] = d_177_v91_
                nw18_[int(1)] = d_177_v91_
                nw18_[int(2)] = d_177_v91_
                nw18_[int(3)] = default__.fm13((0) - ((d_67_v10_).f4), d_53_globalState_)
                nw18_[int(4)] = d_177_v91_
                nw18_[int(5)] = d_177_v91_
                nw18_[int(6)] = (d_177_v91_ if d_175_v89_ else d_177_v91_)
                nw18_[int(7)] = d_177_v91_
                nw18_[int(8)] = d_177_v91_
                nw18_[int(9)] = D1_DC6()
                nw18_[int(10)] = d_177_v91_
                nw18_[int(11)] = d_177_v91_
                nw18_[int(12)] = d_177_v91_
                nw18_[int(13)] = default__.fm13(len(d_74_v16_), d_53_globalState_)
                nw18_[int(14)] = d_177_v91_
                nw18_[int(15)] = d_177_v91_
                nw18_[int(16)] = d_177_v91_
                nw18_[int(17)] = d_177_v91_
                d_178_v92_ = nw18_
                index19_ = default__.safeIndex(506, (d_178_v92_).length(0))
                (d_178_v92_)[index19_] = d_177_v91_
                index20_ = default__.safeIndex(506, (d_178_v92_).length(0))
                (d_178_v92_)[index20_] = d_177_v91_
            d_179_v93_: C0
            nw19_ = C0()
            nw19_.ctor__((d_65_v8_).f7, ((d_164_v78_).f8) | (_dafny.Map({default__.fm0((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], d_56_v2_, d_53_globalState_): (d_67_v10_).f3})), d_55_v1_, (len(d_81_v23_)) + ((0) - (len(_dafny.SeqWithoutIsStrInference([(d_83_v25_).f5 for d_180_i9_ in range(default__.abs(-931))])))))
            d_179_v93_ = nw19_
            d_181_v94_: _dafny.Map
            d_181_v94_ = _dafny.Map({(d_83_v25_).f5: d_163_cf0_})
            source4_ = default__.fm14((d_181_v94_ if d_54_v0_ else d_181_v94_), ((d_67_v10_).f4) * ((d_67_v10_).f4), default__.safeModuloInt(197, (0) - ((d_67_v10_).f4)), d_76_v18_, d_53_globalState_)
            if source4_.is_DC11:
                d_182___mcc_h8_ = source4_.cf17
                d_183_cf17_ = d_182___mcc_h8_
                d_184_v95_: _dafny.Array
                def lambda4_(d_185_v25_):
                    def lambda5_(d_186_i10_):
                        return _dafny.SeqWithoutIsStrInference([(d_185_v25_).f5 for d_187_i11_ in range(default__.abs(813))])

                    return lambda5_

                init2_ = lambda4_(d_83_v25_)
                nw20_ = _dafny.Array(None, 9)
                for i0_2_ in range(nw20_.length(0)):
                    nw20_[i0_2_] = init2_(i0_2_)
                d_184_v95_ = nw20_
                index21_ = default__.safeIndex(617, (d_184_v95_).length(0))
                (d_184_v95_)[index21_] = (_dafny.SeqWithoutIsStrInference([(d_66_v9_) for d_188_i12_ in range(default__.abs(38))])) + (d_81_v23_)
                index22_ = default__.safeIndex(617, (d_184_v95_).length(0))
                index23_ = default__.safeIndex(634, (d_55_v1_).length(0))
                rhs27_ = 849
                rhs28_ = d_81_v23_
                rhs29_ = ((0) - ((d_67_v10_).f4)) * (default__.safeDivisionInt((d_67_v10_).f4, (d_67_v10_).f4))
                rhs30_ = ((0) - (((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]) + ((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))]))) - ((d_67_v10_).f4)
                lhs13_ = d_53_globalState_
                lhs14_ = d_184_v95_
                lhs15_ = default__.safeIndex(617, (d_184_v95_).length(0))
                lhs16_ = d_53_globalState_
                lhs17_ = d_55_v1_
                lhs18_ = default__.safeIndex(634, (d_55_v1_).length(0))
                lhs13_.f1 = rhs27_
                lhs14_[lhs15_] = rhs28_
                lhs16_.f1 = rhs29_
                lhs17_[lhs18_] = rhs30_
                d_189_v96_: _dafny.Array
                nw21_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_189_v96_ = nw21_
                index24_ = default__.safeIndex(703, (d_189_v96_).length(0))
                (d_189_v96_)[index24_] = (d_83_v25_).f6
                index25_ = default__.safeIndex(703, (d_189_v96_).length(0))
                nw22_ = _dafny.Array(None, 9)
                nw22_[int(0)] = False
                nw22_[int(1)] = d_163_cf0_
                nw22_[int(2)] = d_54_v0_
                nw22_[int(3)] = d_54_v0_
                nw22_[int(4)] = False
                nw22_[int(5)] = d_54_v0_
                nw22_[int(6)] = d_163_cf0_
                nw22_[int(7)] = d_163_cf0_
                nw22_[int(8)] = d_54_v0_
                (d_189_v96_)[index25_] = nw22_
                pat_let_tv2_ = d_163_cf0_
                d_190_v97_: C0
                nw23_ = C0()
                def iife15_(_pat_let4_0):
                    def iife16_(d_191_dt__update__tmp_h2_):
                        def iife17_(_pat_let5_0):
                            def iife18_(d_192_dt__update_hcf0_h1_):
                                return D0_DC0(d_192_dt__update_hcf0_h1_)
                            return iife18_(_pat_let5_0)
                        return iife17_(pat_let_tv2_)
                    return iife16_(_pat_let4_0)
                nw23_.ctor__(iife15_((d_179_v93_).f7), _dafny.Map({d_54_v0_: (d_67_v10_).f3}), (d_67_v10_).f3, len(d_166_v80_))
                d_190_v97_ = nw23_
                index26_ = default__.safeIndex(634, (d_55_v1_).length(0))
                (d_55_v1_)[index26_] = (d_67_v10_).f4
            elif source4_.is_DC12:
                d_193___mcc_h9_ = source4_.cf18
                d_194___mcc_h10_ = source4_.cf19
                d_195_cf19_ = d_194___mcc_h10_
                d_196_cf18_ = d_193___mcc_h9_
                d_197_v98_: _dafny.Set
                d_197_v98_ = _dafny.Set({d_163_cf0_})
                (d_53_globalState_).f1 = default__.safeDivisionInt(default__.fm5(not(default__.fm0((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], (d_67_v10_).f4, d_53_globalState_)), d_53_globalState_), len(d_197_v98_))
                index27_ = default__.safeIndex(634, (d_55_v1_).length(0))
                (d_55_v1_)[index27_] = ((d_67_v10_).f4) * (d_56_v2_)
                d_198_v99_: _dafny.Set
                d_198_v99_ = _dafny.Set({d_81_v23_, (d_81_v23_) + (_dafny.SeqWithoutIsStrInference([(d_81_v23_)[default__.safeIndex(d_56_v2_, len(d_81_v23_))] for d_199_i13_ in range(default__.abs(696))])), _dafny.SeqWithoutIsStrInference([(d_83_v25_).f5 for d_200_i14_ in range(default__.abs(974))]), d_81_v23_, d_81_v23_})
                d_198_v99_ = (_dafny.Set({d_81_v23_})).intersection(d_198_v99_)
                d_165_v79_ = d_165_v79_
            elif True:
                d_201___mcc_h11_ = source4_.cf16
                d_202_cf16_ = d_201___mcc_h11_
                d_203_v100_: D0
                d_203_v100_ = D0_DC2(d_81_v23_, (d_67_v10_).f4, (d_67_v10_).f4)
                (d_53_globalState_).f1 = (d_203_v100_).cf5
                d_179_v93_ = d_65_v8_
                index28_ = default__.safeIndex(634, (d_55_v1_).length(0))
                (d_55_v1_)[index28_] = d_56_v2_
                d_204_v101_: C1
                nw24_ = C1()
                nw24_.ctor__((d_83_v25_).f5, d_82_v24_, (d_67_v10_).f3, default__.safeDivisionInt((d_67_v10_).f4, d_56_v2_))
                d_204_v101_ = nw24_
        elif True:
            d_205___mcc_h7_ = source3_.cf7
            d_206_cf7_ = d_205___mcc_h7_
            (d_53_globalState_).f1 = (d_67_v10_).f4
            d_207_v102_: bool
            d_208_v103_: bool
            d_209_v104_: bool
            out12_: bool
            out13_: bool
            out14_: bool
            out12_, out13_, out14_ = (d_83_v25_).m0(d_53_globalState_)
            d_207_v102_ = out12_
            d_208_v103_ = out13_
            d_209_v104_ = out14_
            d_210_v105_: _dafny.Seq
            d_210_v105_ = _dafny.SeqWithoutIsStrInference([len(d_81_v23_), (d_67_v10_).f4])
            d_54_v0_ = default__.fm0((918) - (default__.fm5(default__.fm0(len(d_210_v105_), d_56_v2_, d_53_globalState_), d_53_globalState_)), (d_56_v2_) * (-630), d_53_globalState_)
            d_211_v106_: _dafny.Seq
            d_211_v106_ = _dafny.SeqWithoutIsStrInference([d_67_v10_, d_67_v10_])
            rhs31_ = not(default__.fm0((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], (d_67_v10_).f4, d_53_globalState_))
            rhs32_ = (0) - (d_56_v2_)
            rhs33_ = d_81_v23_
            rhs34_ = (d_211_v106_)[default__.safeIndex(d_56_v2_, len(d_211_v106_))]
            rhs35_ = (D1_DC5(d_209_v104_, d_209_v104_, (d_67_v10_).f4)).cf10
            lhs19_ = d_53_globalState_
            d_54_v0_ = rhs31_
            lhs19_.f1 = rhs32_
            d_81_v23_ = rhs33_
            d_67_v10_ = rhs34_
            d_208_v103_ = rhs35_
        d_212_v107_: _dafny.Map
        d_212_v107_ = _dafny.Map({d_67_v10_: d_54_v0_})
        d_213_v108_: _dafny.Seq
        d_213_v108_ = _dafny.SeqWithoutIsStrInference([len(d_212_v107_)])
        index29_ = default__.safeIndex(634, (d_55_v1_).length(0))
        (d_55_v1_)[index29_] = default__.safeDivisionInt((d_55_v1_)[default__.safeIndex(634, (d_55_v1_).length(0))], default__.safeModuloInt(d_56_v2_, (0) - ((d_213_v108_)[default__.safeIndex(d_56_v2_, len(d_213_v108_))])))
        _dafny.print(_dafny.string_of((d_53_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_53_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_54_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_55_v1_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_55_v1_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_55_v1_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_55_v1_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_56_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_57_v3_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[0]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[1]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[2]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[3]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[4]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[5]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[6]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[7]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[8]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[9]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[10]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[11]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[12]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[13]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[14]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[15]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[16]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[17]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[18]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[19]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[20]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[21]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_58_v4_)[22]) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_62_v5_) == (_dafny.Map({891: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v6_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_64_v7_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_65_v8_).f7).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_65_v8_).f8)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_66_v9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_67_v10_).f3)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_67_v10_).f3)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_67_v10_).f3)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_67_v10_).f3)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_67_v10_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_68_v11_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v12_) == (_dafny.Map({_dafny.Map({891: True}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[0]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[1]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[2]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[3]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[4]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[5]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[6]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[7]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[8]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[9]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[10]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[11]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[12]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[13]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[14]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_70_v13_)[15]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_73_v15_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_74_v16_) == (_dafny.Map({1: 891}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_v17_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v18_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v18_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v18_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_77_v19_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_79_v21_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v22_).cf8)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v22_).cf8)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v22_).cf8)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v22_).cf8)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_81_v23_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v24_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v24_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v25_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_83_v25_).f6)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_83_v25_).f6)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_84_v26_).cf7).cf4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_v26_).cf7).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_v26_).cf7).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_212_v107_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v108_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({self.cf4.VerbatimString(True)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D2_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({self.cf13.VerbatimString(True)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

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
        return lambda: D4_DC11(_dafny.CodePoint('D'))
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

class D4_DC11(D4, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(int(0), int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC14(D5, NamedTuple('DC14', [('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self._f0: int = int(0)
        self._f2: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2

class C0(T0):
    def  __init__(self):
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f4: int = int(0)
        self._f7: D0 = D0.default()()
        self._f8: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f7, f8, f3, f4):
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f3 = f3
        (self)._f4 = f4

    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8

class C1(T0):
    def  __init__(self):
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f4: int = int(0)
        self._f5: str = _dafny.CodePoint('D')
        self._f6: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f5, f6, f3, f4):
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f3 = f3
        (self)._f4 = f4

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        d_214_v0_: _dafny.Seq
        d_214_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwcdnhtm"))
        d_215_v1_: _dafny.Set
        d_215_v1_ = _dafny.Set({d_214_v0_})
        d_216_v2_: bool
        d_216_v2_ = False
        d_217_v3_: D0
        d_217_v3_ = D0_DC0(d_216_v2_)
        d_218_v4_: _dafny.Map
        d_218_v4_ = _dafny.Map({d_215_v1_: (d_217_v3_).cf0})
        d_218_v4_ = (d_218_v4_).set((d_215_v1_) | (d_215_v1_), d_216_v2_)
        d_219_v5_: _dafny.Seq
        d_219_v5_ = _dafny.SeqWithoutIsStrInference([(self).f4, 930, (self).f4, (self).f4, (self).f4])
        d_220_v6_: _dafny.Map
        d_220_v6_ = _dafny.Map({d_216_v2_: (0) - ((self).f4)})
        d_221_v7_: _dafny.MultiSet
        d_221_v7_ = _dafny.MultiSet([(self).f4, len(d_220_v6_), len(d_214_v0_), (self).f4])
        (globalState).f1 = ((0) - (len(d_219_v5_)) if not(not(d_216_v2_)) else default__.safeModuloInt((self).f4, (d_221_v7_).cardinality))
        index30_ = default__.safeIndex(592, ((self).f6).length(0))
        ((self).f6)[index30_] = d_216_v2_
        d_222_v8_: _dafny.Map
        d_222_v8_ = _dafny.Map({d_216_v2_: not(d_216_v2_)})
        index31_ = default__.safeIndex(592, ((self).f6).length(0))
        ((self).f6)[index31_] = (len((_dafny.Map({d_216_v2_: d_216_v2_}) if d_216_v2_ else d_222_v8_))) not in (d_219_v5_)
        d_223_v9_: D0
        d_223_v9_ = D0_DC2(d_214_v0_, len(_dafny.Map({((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]: _dafny.Map({True: d_216_v2_})})), len(d_214_v0_))
        d_224_v10_: _dafny.Seq
        d_224_v10_ = _dafny.SeqWithoutIsStrInference([(d_223_v9_).cf4, d_214_v0_])
        hi0_ = len((d_224_v10_)[default__.safeIndex((self).f4, len(d_224_v10_))])
        for d_225_i0_ in range(-768, hi0_):
            d_226_v11_: _dafny.Seq
            d_226_v11_ = _dafny.SeqWithoutIsStrInference([d_216_v2_])
            index32_ = default__.safeIndex(159, ((self).f3).length(0))
            ((self).f3)[index32_] = ((d_221_v7_).cardinality) + (len(d_226_v11_))
            index33_ = default__.safeIndex(159, ((self).f3).length(0))
            rhs36_ = (d_223_v9_).cf6
            rhs37_ = (default__.safeModuloInt((self).f4, (self).f4)) == (288)
            rhs38_ = d_216_v2_
            rhs39_ = len(d_214_v0_)
            lhs20_ = globalState
            lhs21_ = (self).f3
            lhs22_ = default__.safeIndex(159, ((self).f3).length(0))
            lhs20_.f1 = rhs36_
            r1 = rhs37_
            r1 = rhs38_
            lhs21_[lhs22_] = rhs39_
            (globalState).f1 = default__.safeModuloInt((-468) - (((self).f3)[default__.safeIndex(159, ((self).f3).length(0))]), (self).f4)
            d_227_v12_: _dafny.Map
            d_227_v12_ = _dafny.Map({((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]: (default__.fm1(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], d_216_v2_, globalState)).set(default__.safeIndex(((self).f3)[default__.safeIndex(159, ((self).f3).length(0))], len(default__.fm1(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], d_216_v2_, globalState))), (self).f5)})
            index34_ = default__.safeIndex(592, ((self).f6).length(0))
            ((self).f6)[index34_] = ((d_222_v8_)[((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]] if (((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]) in (d_222_v8_) else not((d_216_v2_) in (d_227_v12_)))
            d_228_v13_: _dafny.Array
            nw25_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_228_v13_ = nw25_
            d_229_v14_: _dafny.Array
            def lambda6_(d_230_v3_):
                def lambda7_(d_231_i1_):
                    return _dafny.SeqWithoutIsStrInference([(d_230_v3_).cf0])

                return lambda7_

            init3_ = lambda6_(d_217_v3_)
            nw26_ = _dafny.Array(None, 5)
            for i0_3_ in range(nw26_.length(0)):
                nw26_[i0_3_] = init3_(i0_3_)
            d_229_v14_ = nw26_
            index35_ = default__.safeIndex(11, (d_228_v13_).length(0))
            (d_228_v13_)[index35_] = d_229_v14_
            index36_ = default__.safeIndex(11, (d_228_v13_).length(0))
            (d_228_v13_)[index36_] = d_229_v14_
        if not(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]):
            r2 = False
            d_232_v15_: _dafny.Map
            d_232_v15_ = _dafny.Map({((self).f4) + ((self).f4): ((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]})
            d_232_v15_ = (d_232_v15_).set((self).f4, ((self).f4) <= ((self).f4))
            d_233_v16_: _dafny.Seq
            d_233_v16_ = _dafny.SeqWithoutIsStrInference([((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], d_216_v2_])
            d_234_v17_: _dafny.Seq
            d_234_v17_ = _dafny.SeqWithoutIsStrInference([(d_233_v16_ if (d_233_v16_)[default__.safeIndex(((d_220_v6_)[((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]] if (((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]) in (d_220_v6_) else (self).f4), len(d_233_v16_))] else _dafny.SeqWithoutIsStrInference([False]))])
            d_234_v17_ = ((d_234_v17_) + (d_234_v17_)) + ((d_234_v17_).set(default__.safeIndex((self).f4, len(d_234_v17_)), d_233_v16_))
            (globalState).f1 = (self).f4
            if d_216_v2_:
                (globalState).f1 = (self).f4
                d_235_v18_: _dafny.Map
                d_235_v18_ = _dafny.Map({((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]: d_232_v15_})
                d_235_v18_ = (d_235_v18_) | ((d_235_v18_) | (default__.fm2((self).f4, not(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]), d_216_v2_, globalState)))
                d_221_v7_ = d_221_v7_
                index37_ = default__.safeIndex(821, ((self).f3).length(0))
                ((self).f3)[index37_] = (self).f4
                index38_ = default__.safeIndex(821, ((self).f3).length(0))
                ((self).f3)[index38_] = (self).f4
                (globalState).f1 = ((self).f3)[default__.safeIndex(821, ((self).f3).length(0))]
            elif True:
                d_236_v19_: _dafny.Set
                d_236_v19_ = _dafny.Set({d_216_v2_})
                index39_ = default__.safeIndex(592, ((self).f6).length(0))
                ((self).f6)[index39_] = (_dafny.Set({False})).ispropersubset(d_236_v19_)
                d_237_v20_: _dafny.Array
                nw27_ = _dafny.Array(_dafny.MultiSet({}), 24)
                d_237_v20_ = nw27_
                d_238_v21_: _dafny.MultiSet
                d_238_v21_ = _dafny.MultiSet([D0_DC0(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))])])
                index40_ = default__.safeIndex(836, (d_237_v20_).length(0))
                (d_237_v20_)[index40_] = d_238_v21_
                d_239_v22_: _dafny.Map
                d_239_v22_ = _dafny.Map({(self).f4: d_232_v15_})
                d_240_v23_: _dafny.MultiSet
                d_240_v23_ = _dafny.MultiSet([default__.fm3((self).f5, (self).f4, globalState), default__.fm3((self).f5, len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_241_i2_ in range(default__.abs(601))])), globalState), (self).f5, (self).f5])
                d_242_v24_: _dafny.Map
                d_242_v24_ = _dafny.Map({(d_239_v22_) == (d_239_v22_): d_240_v23_})
                d_243_v25_: _dafny.Map
                d_243_v25_ = _dafny.Map({(self).f4: (self).f4})
                index41_ = default__.safeIndex(836, (d_237_v20_).length(0))
                rhs40_ = (0) - ((((d_242_v24_)[not(d_216_v2_)] if (not(d_216_v2_)) in (d_242_v24_) else d_240_v23_)).cardinality)
                rhs41_ = ((d_243_v25_)[(self).f4] if ((self).f4) in (d_243_v25_) else default__.safeDivisionInt((self).f4, -768))
                rhs42_ = d_238_v21_
                lhs23_ = globalState
                lhs24_ = globalState
                lhs25_ = d_237_v20_
                lhs26_ = default__.safeIndex(836, (d_237_v20_).length(0))
                lhs23_.f1 = rhs40_
                lhs24_.f1 = rhs41_
                lhs25_[lhs26_] = rhs42_
                (globalState).f1 = len(d_214_v0_)
                d_244_v26_: _dafny.Array
                nw28_ = _dafny.Array(False, 21)
                d_244_v26_ = nw28_
                d_244_v26_ = d_244_v26_
                d_245_v27_: _dafny.Map
                d_245_v27_ = _dafny.Map({((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]: (self).f3})
                d_246_v28_: C0
                nw29_ = C0()
                nw29_.ctor__(D0_DC0(True), (d_245_v27_) | (d_245_v27_), (self).f3, (self).f4)
                d_246_v28_ = nw29_
        elif True:
            (globalState).f1 = (default__.safeDivisionInt((self).f4, len(_dafny.SeqWithoutIsStrInference([d_216_v2_])))) - ((0) - ((self).f4))
            d_214_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "blty"))
            r1 = (True) and (not((d_217_v3_).cf0))
            d_247_v29_: _dafny.Map
            d_247_v29_ = _dafny.Map({((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]: (self).f3})
            d_248_v30_: C0
            nw30_ = C0()
            nw30_.ctor__(d_217_v3_, d_247_v29_, (self).f3, (self).f4)
            d_248_v30_ = nw30_
            r0 = ((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]
        hi1_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hki")))
        for d_249_i3_ in range(len(default__.fm4(((d_222_v8_)[d_216_v2_] if (d_216_v2_) in (d_222_v8_) else d_216_v2_), len(_dafny.SeqWithoutIsStrInference([d_216_v2_, d_216_v2_, d_216_v2_])), (self).f4, globalState)), hi1_):
            d_250_v31_: _dafny.Array
            nw31_ = _dafny.Array(False, 1)
            d_250_v31_ = nw31_
            d_251_v32_: _dafny.Seq
            d_251_v32_ = _dafny.SeqWithoutIsStrInference([d_250_v31_])
            rhs43_ = not(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))])
            rhs44_ = d_214_v0_
            rhs45_ = (d_251_v32_) + (d_251_v32_)
            r0 = rhs43_
            d_214_v0_ = rhs44_
            d_251_v32_ = rhs45_
            d_252_v33_: _dafny.Set
            d_252_v33_ = _dafny.Set({d_250_v31_, d_250_v31_, d_250_v31_, d_250_v31_, d_250_v31_})
            d_253_v34_: _dafny.Array
            def lambda8_(d_254_v0_):
                def lambda9_(d_255_i4_):
                    return d_254_v0_

                return lambda9_

            init4_ = lambda8_(d_214_v0_)
            nw32_ = _dafny.Array(None, 16)
            for i0_4_ in range(nw32_.length(0)):
                nw32_[i0_4_] = init4_(i0_4_)
            d_253_v34_ = nw32_
            index42_ = default__.safeIndex(767, (d_253_v34_).length(0))
            (d_253_v34_)[index42_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcbn"))
            index43_ = default__.safeIndex(767, (d_253_v34_).length(0))
            rhs46_ = d_216_v2_
            rhs47_ = d_252_v33_
            rhs48_ = ((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]
            rhs49_ = d_214_v0_
            lhs27_ = d_253_v34_
            lhs28_ = default__.safeIndex(767, (d_253_v34_).length(0))
            r0 = rhs46_
            d_252_v33_ = rhs47_
            r1 = rhs48_
            lhs27_[lhs28_] = rhs49_
            r0 = d_216_v2_
            source5_ = d_223_v9_
            if source5_.is_DC1:
                d_256___mcc_h0_ = source5_.cf1
                d_257___mcc_h1_ = source5_.cf2
                d_258___mcc_h2_ = source5_.cf3
                d_259_cf3_ = d_258___mcc_h2_
                d_260_cf2_ = d_257___mcc_h1_
                d_261_cf1_ = d_256___mcc_h0_
                d_262_v35_: _dafny.Array
                nw33_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_262_v35_ = nw33_
                index44_ = default__.safeIndex(762, (d_262_v35_).length(0))
                (d_262_v35_)[index44_] = d_250_v31_
                index45_ = default__.safeIndex(592, ((self).f6).length(0))
                index46_ = default__.safeIndex(762, (d_262_v35_).length(0))
                rhs50_ = ((self).f5) in (d_214_v0_)
                rhs51_ = (d_251_v32_)[default__.safeIndex((self).f4, len(d_251_v32_))]
                rhs52_ = d_216_v2_
                lhs29_ = (self).f6
                lhs30_ = default__.safeIndex(592, ((self).f6).length(0))
                lhs31_ = d_262_v35_
                lhs32_ = default__.safeIndex(762, (d_262_v35_).length(0))
                lhs29_[lhs30_] = rhs50_
                lhs31_[lhs32_] = rhs51_
                r2 = rhs52_
                d_263_v36_: _dafny.Set
                d_263_v36_ = _dafny.Set({(self).f4, d_260_cf2_})
                d_263_v36_ = (d_263_v36_) | (d_263_v36_)
                d_264_v37_: _dafny.Seq
                d_264_v37_ = _dafny.SeqWithoutIsStrInference([d_217_v3_, d_217_v3_, D0_DC0(d_216_v2_)])
                d_265_v38_: str
                d_265_v38_ = _dafny.CodePoint('h')
                d_266_v39_: D1
                d_266_v39_ = D1_DC4((self).f3)
                d_267_v40_: _dafny.Seq
                d_267_v40_ = _dafny.SeqWithoutIsStrInference([(self).f3])
                d_268_v41_: _dafny.Array
                nw34_ = _dafny.Array(None, 28)
                nw34_[int(0)] = (self).f3
                nw34_[int(1)] = (d_266_v39_).cf8
                nw34_[int(2)] = (self).f3
                nw34_[int(3)] = (self).f3
                nw34_[int(4)] = (self).f3
                nw34_[int(5)] = (self).f3
                nw34_[int(6)] = (self).f3
                nw34_[int(7)] = (self).f3
                nw34_[int(8)] = (self).f3
                nw34_[int(9)] = (self).f3
                nw34_[int(10)] = (self).f3
                nw34_[int(11)] = (self).f3
                nw34_[int(12)] = (self).f3
                nw34_[int(13)] = (self).f3
                nw34_[int(14)] = (self).f3
                nw34_[int(15)] = (self).f3
                nw34_[int(16)] = (self).f3
                nw34_[int(17)] = (self).f3
                nw34_[int(18)] = (self).f3
                nw34_[int(19)] = (self).f3
                nw34_[int(20)] = (self).f3
                nw34_[int(21)] = (self).f3
                nw34_[int(22)] = (d_267_v40_)[default__.safeIndex(347, len(d_267_v40_))]
                nw34_[int(23)] = (self).f3
                nw34_[int(24)] = (self).f3
                nw34_[int(25)] = (self).f3
                nw34_[int(26)] = (d_267_v40_)[default__.safeIndex(len(d_219_v5_), len(d_267_v40_))]
                nw34_[int(27)] = (self).f3
                d_268_v41_ = nw34_
                d_269_v42_: _dafny.Seq
                d_269_v42_ = _dafny.SeqWithoutIsStrInference([d_268_v41_, d_268_v41_, d_268_v41_, d_268_v41_, d_268_v41_])
                d_270_v43_: _dafny.Array
                nw35_ = _dafny.Array(None, 18)
                nw35_[int(0)] = d_268_v41_
                nw35_[int(1)] = d_268_v41_
                nw35_[int(2)] = d_268_v41_
                nw35_[int(3)] = d_268_v41_
                nw35_[int(4)] = d_268_v41_
                nw35_[int(5)] = (d_268_v41_ if default__.fm0((self).f4, d_260_cf2_, globalState) else d_268_v41_)
                nw35_[int(6)] = d_268_v41_
                nw35_[int(7)] = d_268_v41_
                nw35_[int(8)] = (d_269_v42_)[default__.safeIndex(d_260_cf2_, len(d_269_v42_))]
                nw35_[int(9)] = d_268_v41_
                nw35_[int(10)] = d_268_v41_
                nw35_[int(11)] = d_268_v41_
                nw35_[int(12)] = d_268_v41_
                nw35_[int(13)] = d_268_v41_
                nw35_[int(14)] = (d_268_v41_ if ((self).f6)[default__.safeIndex(592, ((self).f6).length(0))] else d_268_v41_)
                nw35_[int(15)] = d_268_v41_
                nw35_[int(16)] = d_268_v41_
                nw35_[int(17)] = d_268_v41_
                d_270_v43_ = nw35_
                index47_ = default__.safeIndex(787, (d_270_v43_).length(0))
                (d_270_v43_)[index47_] = d_268_v41_
                index48_ = default__.safeIndex(787, (d_270_v43_).length(0))
                def iife19_(_pat_let6_0):
                    def iife20_(d_271_dt__update__tmp_h0_):
                        def iife21_(_pat_let7_0):
                            def iife22_(d_272_dt__update_hcf0_h0_):
                                return D0_DC0(d_272_dt__update_hcf0_h0_)
                            return iife22_(_pat_let7_0)
                        return iife21_(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))])
                    return iife20_(_pat_let6_0)
                def iife23_(_pat_let8_0):
                    def iife24_(d_273_dt__update__tmp_h1_):
                        def iife25_(_pat_let9_0):
                            def iife26_(d_274_dt__update_hcf0_h1_):
                                return D0_DC0(d_274_dt__update_hcf0_h1_)
                            return iife26_(_pat_let9_0)
                        return iife25_(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))])
                    return iife24_(_pat_let8_0)
                rhs53_ = (_dafny.SeqWithoutIsStrInference([D0_DC0(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]), iife19_(D0_DC0(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]))])).set(default__.safeIndex(d_261_cf1_, len(_dafny.SeqWithoutIsStrInference([D0_DC0(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]), iife23_(D0_DC0(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]))]))), d_217_v3_)
                rhs54_ = d_265_v38_
                rhs55_ = d_268_v41_
                lhs33_ = d_270_v43_
                lhs34_ = default__.safeIndex(787, (d_270_v43_).length(0))
                d_264_v37_ = rhs53_
                d_265_v38_ = rhs54_
                lhs33_[lhs34_] = rhs55_
                index49_ = default__.safeIndex(499, ((self).f3).length(0))
                ((self).f3)[index49_] = d_249_i3_
                index50_ = default__.safeIndex(499, ((self).f3).length(0))
                ((self).f3)[index50_] = (0) - (d_261_cf1_)
            elif source5_.is_DC2:
                d_275___mcc_h3_ = source5_.cf4
                d_276___mcc_h4_ = source5_.cf5
                d_277___mcc_h5_ = source5_.cf6
                d_278_cf6_ = d_277___mcc_h5_
                d_279_cf5_ = d_276___mcc_h4_
                d_280_cf4_ = d_275___mcc_h3_
                (globalState).f1 = d_279_cf5_
                d_281_v44_: _dafny.Map
                d_281_v44_ = _dafny.Map({d_216_v2_: (self).f3})
                d_282_v45_: C0
                nw36_ = C0()
                nw36_.ctor__(d_217_v3_, d_281_v44_, (self).f3, default__.safeDivisionInt((self).f4, d_249_i3_))
                d_282_v45_ = nw36_
                d_282_v45_ = d_282_v45_
                d_283_v46_: _dafny.Seq
                out15_: _dafny.Seq
                out15_ = (self).m1(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], 484, (0) - (d_249_i3_), globalState)
                d_283_v46_ = out15_
                d_284_v47_: C0
                nw37_ = C0()
                nw37_.ctor__((d_282_v45_).f7, (d_281_v44_) | ((d_282_v45_).f8), (self).f3, d_249_i3_)
                d_284_v47_ = nw37_
            elif source5_.is_DC0:
                d_285___mcc_h6_ = source5_.cf0
                d_286_cf0_ = d_285___mcc_h6_
                d_287_v48_: _dafny.Map
                d_287_v48_ = _dafny.Map({d_286_cf0_: (self).f3})
                d_288_v49_: C0
                nw38_ = C0()
                nw38_.ctor__(d_217_v3_, d_287_v48_, (self).f3, (len(d_214_v0_) if d_216_v2_ else (self).f4))
                d_288_v49_ = nw38_
                d_289_v50_: D1
                d_289_v50_ = D1_DC6()
                d_290_v51_: D0
                d_290_v51_ = D0_DC1(len(_dafny.Map({d_216_v2_: d_249_i3_})), len(d_219_v5_), 852)
                d_291_v52_: _dafny.Map
                d_291_v52_ = _dafny.Map({d_289_v50_: d_290_v51_})
                d_291_v52_ = (d_291_v52_).set(D1_DC6(), d_290_v51_)
                d_292_v53_: _dafny.Map
                d_292_v53_ = _dafny.Map({len((d_253_v34_)[default__.safeIndex(767, (d_253_v34_).length(0))]): True})
                d_222_v8_ = (d_222_v8_).set((((d_292_v53_)[d_249_i3_] if (d_249_i3_) in (d_292_v53_) else default__.fm0((0) - (default__.fm5(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], globalState)), (self).f4, globalState)) if ((d_222_v8_)[d_286_cf0_] if (d_286_cf0_) in (d_222_v8_) else d_286_cf0_) else d_216_v2_), d_286_cf0_)
                d_293_v54_: _dafny.Map
                d_293_v54_ = _dafny.Map({(self).f5: len(_dafny.Set({d_216_v2_, d_216_v2_}))})
                d_294_v55_: _dafny.Map
                d_294_v55_ = _dafny.Map({d_216_v2_: d_222_v8_})
                rhs56_ = default__.fm5(False, globalState)
                rhs57_ = d_293_v54_
                rhs58_ = ((len(default__.fm6(True, ((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], d_249_i3_, globalState))) >= (d_249_i3_)) or (d_216_v2_)
                rhs59_ = (d_294_v55_ if ((self).f4) <= (d_249_i3_) else (d_294_v55_) | (d_294_v55_))
                rhs60_ = d_217_v3_
                lhs35_ = globalState
                lhs35_.f1 = rhs56_
                d_293_v54_ = rhs57_
                d_286_cf0_ = rhs58_
                d_294_v55_ = rhs59_
                d_217_v3_ = rhs60_
            elif True:
                d_295___mcc_h7_ = source5_.cf7
                d_296_cf7_ = d_295___mcc_h7_
                d_297_v56_: _dafny.Array
                nw39_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_297_v56_ = nw39_
                d_298_v57_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_298_v57_ = nw40_
                index51_ = default__.safeIndex(27, (d_297_v56_).length(0))
                (d_297_v56_)[index51_] = d_298_v57_
                index52_ = default__.safeIndex(27, (d_297_v56_).length(0))
                (d_297_v56_)[index52_] = d_298_v57_
                d_299_v58_: str
                d_299_v58_ = _dafny.CodePoint('r')
                d_299_v58_ = d_299_v58_
                (globalState).f1 = default__.safeDivisionInt((self).f4, len(default__.fm7(((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], globalState)))
                (globalState).f1 = -569
        d_300_v59_: _dafny.Seq
        d_300_v59_ = _dafny.SeqWithoutIsStrInference([((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], ((self).f6)[default__.safeIndex(592, ((self).f6).length(0))], d_216_v2_])
        r0 = (d_300_v59_)[default__.safeIndex((self).f4, len(d_300_v59_))]
        r1 = d_216_v2_
        r2 = ((self).f6)[default__.safeIndex(592, ((self).f6).length(0))]
        return r0, r1, r2

    def m1(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_301_i0_: int
        d_301_i0_ = 0
        with _dafny.label("0"):
            while p0:
                with _dafny.c_label("0"):
                    if (d_301_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_301_i0_ = (d_301_i0_) + (1)
                    d_302_v0_: _dafny.Array
                    nw41_ = _dafny.Array(_dafny.Map({}), 19)
                    d_302_v0_ = nw41_
                    d_302_v0_ = d_302_v0_
                    d_303_v1_: _dafny.Map
                    d_303_v1_ = _dafny.Map({(self).f4: (self).f4})
                    d_304_v2_: _dafny.Seq
                    d_304_v2_ = _dafny.SeqWithoutIsStrInference([d_303_v1_, d_303_v1_, d_303_v1_])
                    d_305_v3_: _dafny.Map
                    d_305_v3_ = _dafny.Map({(0) - (-227): (d_304_v2_)[default__.safeIndex((self).f4, len(d_304_v2_))]})
                    index53_ = default__.safeIndex(803, ((self).f3).length(0))
                    ((self).f3)[index53_] = len(((d_305_v3_)[p1] if (p1) in (d_305_v3_) else d_303_v1_))
                    index54_ = default__.safeIndex(439, ((self).f6).length(0))
                    ((self).f6)[index54_] = p0
                    index55_ = default__.safeIndex(803, ((self).f3).length(0))
                    index56_ = default__.safeIndex(439, ((self).f6).length(0))
                    rhs61_ = p1
                    rhs62_ = not(p0)
                    lhs36_ = (self).f3
                    lhs37_ = default__.safeIndex(803, ((self).f3).length(0))
                    lhs38_ = (self).f6
                    lhs39_ = default__.safeIndex(439, ((self).f6).length(0))
                    lhs36_[lhs37_] = rhs61_
                    lhs38_[lhs39_] = rhs62_
                    d_306_v4_: _dafny.Seq
                    d_306_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                    d_307_v5_: _dafny.Seq
                    d_307_v5_ = _dafny.SeqWithoutIsStrInference([d_306_v4_, d_306_v4_])
                    d_308_v6_: D0
                    d_308_v6_ = D0_DC0(p0)
                    d_309_v7_: _dafny.Array
                    def lambda10_(d_310_i1_):
                        return (d_310_i1_) * ((self).f4)

                    init5_ = lambda10_
                    nw42_ = _dafny.Array(None, 19)
                    for i0_5_ in range(nw42_.length(0)):
                        nw42_[i0_5_] = init5_(i0_5_)
                    d_309_v7_ = nw42_
                    d_311_v8_: _dafny.Map
                    d_311_v8_ = _dafny.Map({((self).f6)[default__.safeIndex(439, ((self).f6).length(0))]: d_309_v7_})
                    d_312_v9_: _dafny.Seq
                    d_312_v9_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_313_v10_: T0
                    nw43_ = C0()
                    nw43_.ctor__(d_308_v6_, d_311_v8_, d_309_v7_, len(d_312_v9_))
                    d_313_v10_ = nw43_
                    d_314_v11_: _dafny.Map
                    d_314_v11_ = _dafny.Map({((d_307_v5_) + (d_307_v5_)).set(default__.safeIndex(((self).f3)[default__.safeIndex(803, ((self).f3).length(0))], len((d_307_v5_) + (d_307_v5_))), d_306_v4_): d_313_v10_})
                    d_315_v12_: _dafny.Map
                    d_315_v12_ = _dafny.Map({p2: d_307_v5_})
                    d_316_v13_: _dafny.Map
                    d_316_v13_ = _dafny.Map({p0: d_313_v10_})
                    d_314_v11_ = (d_314_v11_).set(((d_315_v12_)[(d_313_v10_).f4] if ((d_313_v10_).f4) in (d_315_v12_) else d_307_v5_), ((d_316_v13_)[not(p0)] if (not(p0)) in (d_316_v13_) else d_313_v10_))
                    d_317_v14_: _dafny.Map
                    d_317_v14_ = _dafny.Map({d_306_v4_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcftnkrjh"))})
                    d_317_v14_ = (d_317_v14_).set(d_306_v4_, d_306_v4_)
                    pass
            pass
        d_318_v15_: _dafny.Array
        nw44_ = _dafny.Array(_dafny.Seq({}), 16)
        d_318_v15_ = nw44_
        d_318_v15_ = d_318_v15_
        d_319_v16_: _dafny.Array
        def lambda11_(d_320_i3_):
            return default__.safeDivisionInt(d_320_i3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "noju"))))

        init6_ = lambda11_
        nw45_ = _dafny.Array(None, 17)
        for i0_6_ in range(nw45_.length(0)):
            nw45_[i0_6_] = init6_(i0_6_)
        d_319_v16_ = nw45_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_319_v16_).length(0)):
            d_321_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_321_i2_)) and ((d_321_i2_) < ((d_319_v16_).length(0)))):
                (d_319_v16_)[(d_321_i2_)] = (d_321_i2_) + ((0) - ((self).f4))
        if p0:
            d_322_v17_: D1
            d_322_v17_ = D1_DC4((self).f3)
            d_323_v18_: _dafny.MultiSet
            d_323_v18_ = _dafny.MultiSet([(d_322_v17_).cf8])
            d_324_v19_: _dafny.Seq
            d_324_v19_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_319_v16_]), d_323_v18_, d_323_v18_])
            d_325_v20_: D0
            d_325_v20_ = D0_DC0(p0)
            d_326_v21_: _dafny.Map
            d_326_v21_ = _dafny.Map({p0: d_319_v16_})
            d_327_v22_: C0
            nw46_ = C0()
            nw46_.ctor__(d_325_v20_, d_326_v21_, (self).f3, p2)
            d_327_v22_ = nw46_
            d_328_v23_: _dafny.Map
            d_328_v23_ = _dafny.Map({(_dafny.MultiSet([(d_322_v17_).cf8])).ispropersubset((d_324_v19_)[default__.safeIndex(p1, len(d_324_v19_))]): d_327_v22_})
            d_329_v24_: bool
            d_329_v24_ = False
            d_330_v25_: str
            d_330_v25_ = _dafny.CodePoint('p')
            d_331_v26_: _dafny.Map
            d_331_v26_ = _dafny.Map({d_330_v25_: (self).f5})
            d_332_v27_: _dafny.Seq
            d_332_v27_ = _dafny.SeqWithoutIsStrInference([((d_331_v26_)[d_330_v25_] if (d_330_v25_) in (d_331_v26_) else d_330_v25_)])
            d_333_v28_: _dafny.Set
            d_333_v28_ = _dafny.Set({p1})
            d_334_v29_: _dafny.Array
            nw47_ = _dafny.Array(None, 23)
            nw47_[int(0)] = d_332_v27_
            nw47_[int(1)] = d_332_v27_
            nw47_[int(2)] = ((d_332_v27_ if d_329_v24_ else _dafny.SeqWithoutIsStrInference([d_330_v25_, (self).f5]))).set(default__.safeIndex(p2, len((d_332_v27_ if d_329_v24_ else _dafny.SeqWithoutIsStrInference([d_330_v25_, (self).f5])))), _dafny.CodePoint('l'))
            nw47_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_330_v25_ for d_335_i4_ in range(default__.abs(940))])) + (_dafny.SeqWithoutIsStrInference([d_330_v25_ for d_336_i5_ in range(default__.abs(92))]))
            nw47_[int(4)] = d_332_v27_
            nw47_[int(5)] = ((d_332_v27_) + (d_332_v27_)).set(default__.safeIndex((0) - (p2), len((d_332_v27_) + (d_332_v27_))), (d_332_v27_)[default__.safeIndex(p1, len(d_332_v27_))])
            nw47_[int(6)] = d_332_v27_
            nw47_[int(7)] = d_332_v27_
            nw47_[int(8)] = d_332_v27_
            nw47_[int(9)] = d_332_v27_
            nw47_[int(10)] = d_332_v27_
            nw47_[int(11)] = _dafny.SeqWithoutIsStrInference([(self).f5 for d_337_i6_ in range(default__.abs(924))])
            nw47_[int(12)] = d_332_v27_
            nw47_[int(13)] = d_332_v27_
            nw47_[int(14)] = d_332_v27_
            nw47_[int(15)] = (d_332_v27_) + (d_332_v27_)
            nw47_[int(16)] = d_332_v27_
            nw47_[int(17)] = (_dafny.SeqWithoutIsStrInference([(self).f5 for d_338_i7_ in range(default__.abs(938))]) if p0 else _dafny.SeqWithoutIsStrInference([(self).f5]))
            nw47_[int(18)] = (d_332_v27_) + (d_332_v27_)
            nw47_[int(19)] = ((d_332_v27_).set(default__.safeIndex(len(d_333_v28_), len(d_332_v27_)), (self).f5)) + (d_332_v27_)
            nw47_[int(20)] = _dafny.SeqWithoutIsStrInference([d_330_v25_ for d_339_i8_ in range(default__.abs(516))])
            nw47_[int(21)] = d_332_v27_
            nw47_[int(22)] = _dafny.SeqWithoutIsStrInference([(self).f5])
            d_334_v29_ = nw47_
            d_340_v30_: _dafny.Seq
            d_340_v30_ = _dafny.SeqWithoutIsStrInference([d_332_v27_])
            index57_ = default__.safeIndex(266, (d_334_v29_).length(0))
            (d_334_v29_)[index57_] = (d_340_v30_)[default__.safeIndex(p2, len(d_340_v30_))]
            d_341_v31_: _dafny.Map
            d_341_v31_ = _dafny.Map({True: p1})
            index58_ = default__.safeIndex(381, ((self).f3).length(0))
            ((self).f3)[index58_] = ((d_341_v31_)[p0] if (p0) in (d_341_v31_) else p1)
            d_342_v32_: _dafny.Seq
            d_342_v32_ = _dafny.SeqWithoutIsStrInference([d_329_v24_, p0])
            index59_ = default__.safeIndex(266, (d_334_v29_).length(0))
            index60_ = default__.safeIndex(381, ((self).f3).length(0))
            rhs63_ = (d_328_v23_).set(d_329_v24_, d_327_v22_)
            rhs64_ = p0
            rhs65_ = d_330_v25_
            rhs66_ = default__.fm1(p0, (default__.fm5(d_329_v24_, globalState)) >= (len(d_342_v32_)), globalState)
            rhs67_ = 157
            lhs40_ = d_334_v29_
            lhs41_ = default__.safeIndex(266, (d_334_v29_).length(0))
            lhs42_ = (self).f3
            lhs43_ = default__.safeIndex(381, ((self).f3).length(0))
            d_328_v23_ = rhs63_
            d_329_v24_ = rhs64_
            d_330_v25_ = rhs65_
            lhs40_[lhs41_] = rhs66_
            lhs42_[lhs43_] = rhs67_
            if p0:
                d_329_v24_ = d_329_v24_
                d_343_v33_: _dafny.Array
                def lambda12_(d_344_p0_, d_345_v24_):
                    def lambda13_(d_346_i9_):
                        return (d_344_p0_ if d_344_p0_ else d_345_v24_)

                    return lambda13_

                init7_ = lambda12_(p0, d_329_v24_)
                nw48_ = _dafny.Array(None, 17)
                for i0_7_ in range(nw48_.length(0)):
                    nw48_[i0_7_] = init7_(i0_7_)
                d_343_v33_ = nw48_
                d_347_v34_: _dafny.Array
                def lambda14_(d_348_v24_):
                    def lambda15_(d_349_i10_):
                        return _dafny.SeqWithoutIsStrInference([d_348_v24_])

                    return lambda15_

                init8_ = lambda14_(d_329_v24_)
                nw49_ = _dafny.Array(None, 16)
                for i0_8_ in range(nw49_.length(0)):
                    nw49_[i0_8_] = init8_(i0_8_)
                d_347_v34_ = nw49_
                index61_ = default__.safeIndex(266, (d_334_v29_).length(0))
                rhs68_ = (self).f6
                rhs69_ = (self).f5
                rhs70_ = d_347_v34_
                rhs71_ = (d_334_v29_)[default__.safeIndex(266, (d_334_v29_).length(0))]
                rhs72_ = default__.fm0(p2, p2, globalState)
                lhs44_ = d_334_v29_
                lhs45_ = default__.safeIndex(266, (d_334_v29_).length(0))
                d_343_v33_ = rhs68_
                d_330_v25_ = rhs69_
                d_347_v34_ = rhs70_
                lhs44_[lhs45_] = rhs71_
                d_329_v24_ = rhs72_
                index62_ = default__.safeIndex(381, ((self).f3).length(0))
                ((self).f3)[index62_] = (default__.fm5(not(d_329_v24_), globalState) if (p0) or (d_329_v24_) else 469)
                (globalState).f1 = 683
                d_350_v35_: _dafny.MultiSet
                d_350_v35_ = _dafny.MultiSet([p0])
                d_351_v36_: _dafny.Map
                d_351_v36_ = _dafny.Map({default__.fm0((self).f4, ((self).f3)[default__.safeIndex(381, ((self).f3).length(0))], globalState): d_350_v35_})
                d_351_v36_ = _dafny.Map({d_329_v24_: d_350_v35_})
            elif True:
                (globalState).f1 = default__.fm5(p0, globalState)
                d_352_v37_: _dafny.Map
                d_352_v37_ = _dafny.Map({p0: p0})
                index63_ = default__.safeIndex(696, ((self).f6).length(0))
                ((self).f6)[index63_] = ((d_352_v37_)[p0] if (p0) in (d_352_v37_) else p0)
                index64_ = default__.safeIndex(696, ((self).f6).length(0))
                ((self).f6)[index64_] = (175) == (p1)
                index65_ = default__.safeIndex(381, ((self).f3).length(0))
                ((self).f3)[index65_] = p1
                nw50_ = C0()
                nw50_.ctor__(d_325_v20_, ((d_327_v22_).f8).set(((self).f6)[default__.safeIndex(696, ((self).f6).length(0))], d_319_v16_), d_319_v16_, -771)
                d_327_v22_ = nw50_
                (globalState).f1 = (0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")) if ((self).f6)[default__.safeIndex(696, ((self).f6).length(0))] else d_332_v27_)))
            if p0:
                d_353_v38_: D0
                d_353_v38_ = D0_DC2(d_332_v27_, ((self).f3)[default__.safeIndex(381, ((self).f3).length(0))], p1)
                index66_ = default__.safeIndex(266, (d_334_v29_).length(0))
                (d_334_v29_)[index66_] = (((d_334_v29_)[default__.safeIndex(266, (d_334_v29_).length(0))]) + ((d_353_v38_).cf4)) + ((d_334_v29_)[default__.safeIndex(266, (d_334_v29_).length(0))])
                d_354_v39_: C0
                nw51_ = C0()
                nw51_.ctor__(D0_DC0(d_329_v24_), ((d_327_v22_).f8) | ((d_327_v22_).f8), d_319_v16_, (-897) - (p1))
                d_354_v39_ = nw51_
                d_355_v40_: C0
                nw52_ = C0()
                nw52_.ctor__(D0_DC0(d_329_v24_), (d_354_v39_).f8, d_319_v16_, default__.safeDivisionInt(p2, p1))
                d_355_v40_ = nw52_
                (globalState).f1 = p1
                index67_ = default__.safeIndex(381, ((self).f3).length(0))
                ((self).f3)[index67_] = p1
            elif True:
                d_341_v31_ = ((d_341_v31_) | (d_341_v31_)) | (d_341_v31_)
                d_356_v41_: _dafny.Map
                d_356_v41_ = _dafny.Map({d_329_v24_: d_329_v24_})
                d_357_v42_: _dafny.Set
                d_357_v42_ = _dafny.Set({(d_327_v22_).f7, d_325_v20_, (d_327_v22_).f7})
                d_358_v43_: _dafny.Array
                nw53_ = _dafny.Array(None, 19)
                nw53_[int(0)] = p0
                nw53_[int(1)] = d_329_v24_
                nw53_[int(2)] = p0
                nw53_[int(3)] = d_329_v24_
                nw53_[int(4)] = p0
                nw53_[int(5)] = d_329_v24_
                nw53_[int(6)] = d_329_v24_
                nw53_[int(7)] = (((self).f3)[default__.safeIndex(381, ((self).f3).length(0))]) < ((self).f4)
                nw53_[int(8)] = ((self).f4) <= (p2)
                nw53_[int(9)] = p0
                nw53_[int(10)] = not(d_329_v24_)
                nw53_[int(11)] = d_329_v24_
                nw53_[int(12)] = ((d_356_v41_)[p0] if (p0) in (d_356_v41_) else p0)
                nw53_[int(13)] = (d_357_v42_).ispropersubset(default__.fm8(d_329_v24_, globalState))
                nw53_[int(14)] = p0
                nw53_[int(15)] = d_329_v24_
                nw53_[int(16)] = p0
                nw53_[int(17)] = (((self).f3)[default__.safeIndex(381, ((self).f3).length(0))]) < (((d_341_v31_)[default__.fm0(((self).f3)[default__.safeIndex(381, ((self).f3).length(0))], len(d_333_v28_), globalState)] if (default__.fm0(((self).f3)[default__.safeIndex(381, ((self).f3).length(0))], len(d_333_v28_), globalState)) in (d_341_v31_) else p1))
                nw53_[int(18)] = not(True)
                d_358_v43_ = nw53_
                def lambda16_(d_359_p0_):
                    def lambda17_(d_360_i11_):
                        return d_359_p0_

                    return lambda17_

                init9_ = lambda16_(p0)
                nw54_ = _dafny.Array(None, 1)
                for i0_9_ in range(nw54_.length(0)):
                    nw54_[i0_9_] = init9_(i0_9_)
                d_358_v43_ = nw54_
                d_361_v44_: _dafny.Seq
                d_361_v44_ = _dafny.SeqWithoutIsStrInference([247, (self).f4])
                d_362_v45_: _dafny.Set
                d_362_v45_ = _dafny.Set({not(not(d_329_v24_)), True})
                d_363_v46_: _dafny.Map
                d_363_v46_ = _dafny.Map({default__.fm5(p0, globalState): d_362_v45_})
                d_364_v47_: _dafny.Map
                d_364_v47_ = _dafny.Map({(d_332_v27_) + (d_332_v27_): (d_361_v44_) + ((d_361_v44_).set(default__.safeIndex(len(d_363_v46_), len(d_361_v44_)), 32))})
                rhs73_ = d_329_v24_
                rhs74_ = (d_341_v31_).set(not(p0), ((self).f3)[default__.safeIndex(381, ((self).f3).length(0))])
                rhs75_ = ((d_364_v47_)[d_332_v27_] if (d_332_v27_) in (d_364_v47_) else d_361_v44_)
                d_329_v24_ = rhs73_
                d_341_v31_ = rhs74_
                d_361_v44_ = rhs75_
                index68_ = default__.safeIndex(381, ((self).f3).length(0))
                ((self).f3)[index68_] = ((self).f4) + (p2)
                d_365_v48_: D1
                d_365_v48_ = D1_DC5((p0) and (d_329_v24_), True, ((self).f3)[default__.safeIndex(381, ((self).f3).length(0))])
                pat_let_tv3_ = d_365_v48_
                pat_let_tv4_ = d_329_v24_
                def iife27_(_pat_let10_0):
                    def iife28_(d_366_dt__update__tmp_h0_):
                        def iife29_(_pat_let11_0):
                            def iife30_(d_367_dt__update_hcf9_h0_):
                                def iife31_(_pat_let12_0):
                                    def iife32_(d_368_dt__update_hcf10_h0_):
                                        return D1_DC5(d_367_dt__update_hcf9_h0_, d_368_dt__update_hcf10_h0_, (d_366_dt__update__tmp_h0_).cf11)
                                    return iife32_(_pat_let12_0)
                                return iife31_(pat_let_tv4_)
                            return iife30_(_pat_let11_0)
                        return iife29_((pat_let_tv3_).cf9)
                    return iife28_(_pat_let10_0)
                d_365_v48_ = iife27_(D1_DC5(p0, p0, -408))
            d_329_v24_ = (-255) >= (p2)
            d_332_v27_ = (d_334_v29_)[default__.safeIndex(266, (d_334_v29_).length(0))]
        elif True:
            d_369_v49_: _dafny.Seq
            d_369_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uaymukwrs"))
            d_370_v50_: _dafny.Array
            nw55_ = _dafny.Array(None, 13)
            nw55_[int(0)] = (self).f5
            nw55_[int(1)] = (d_369_v49_)[default__.safeIndex((self).f4, len(d_369_v49_))]
            nw55_[int(2)] = (self).f5
            nw55_[int(3)] = (self).f5
            nw55_[int(4)] = default__.fm3((self).f5, p2, globalState)
            nw55_[int(5)] = (self).f5
            nw55_[int(6)] = (self).f5
            nw55_[int(7)] = (self).f5
            nw55_[int(8)] = (self).f5
            nw55_[int(9)] = (self).f5
            nw55_[int(10)] = (self).f5
            nw55_[int(11)] = _dafny.CodePoint('x')
            nw55_[int(12)] = (self).f5
            d_370_v50_ = nw55_
            index69_ = default__.safeIndex(496, (d_370_v50_).length(0))
            (d_370_v50_)[index69_] = (self).f5
            index70_ = default__.safeIndex(496, (d_370_v50_).length(0))
            (d_370_v50_)[index70_] = ((self).f5 if p0 else (self).f5)
            d_371_v51_: bool
            d_371_v51_ = False
            d_371_v51_ = p0
            d_372_v52_: _dafny.Map
            d_372_v52_ = _dafny.Map({553: d_369_v49_})
            d_373_v53_: _dafny.Set
            d_373_v53_ = _dafny.Set({len(d_372_v52_)})
            d_374_v54_: _dafny.Seq
            d_374_v54_ = _dafny.SeqWithoutIsStrInference([d_371_v51_, d_371_v51_])
            d_375_v55_: D0
            d_375_v55_ = D0_DC2(d_369_v49_, len(d_373_v53_), (0) - (len(d_374_v54_)))
            d_376_v56_: D0
            d_376_v56_ = D0_DC3(d_375_v55_)
            d_377_v57_: D0
            d_377_v57_ = D0_DC3(d_376_v56_)
            source6_ = d_377_v57_
            if source6_.is_DC1:
                d_378___mcc_h0_ = source6_.cf1
                d_379___mcc_h1_ = source6_.cf2
                d_380___mcc_h2_ = source6_.cf3
                d_381_cf3_ = d_380___mcc_h2_
                d_382_cf2_ = d_379___mcc_h1_
                d_383_cf1_ = d_378___mcc_h0_
                d_384_v58_: _dafny.Seq
                d_384_v58_ = _dafny.SeqWithoutIsStrInference([D0_DC2(_dafny.SeqWithoutIsStrInference([(self).f5 for d_385_i12_ in range(default__.abs(783))]), (self).f4, p2)])
                d_386_v59_: _dafny.Seq
                d_386_v59_ = _dafny.SeqWithoutIsStrInference([len(d_384_v58_)])
                d_387_v60_: _dafny.Array
                nw56_ = _dafny.Array(None, 10)
                nw56_[int(0)] = (d_374_v54_) + (d_374_v54_)
                nw56_[int(1)] = d_374_v54_
                nw56_[int(2)] = d_374_v54_
                nw56_[int(3)] = d_374_v54_
                nw56_[int(4)] = _dafny.SeqWithoutIsStrInference([p0, d_371_v51_, p0, True])
                nw56_[int(5)] = d_374_v54_
                nw56_[int(6)] = d_374_v54_
                nw56_[int(7)] = (d_374_v54_ if p0 else d_374_v54_)
                nw56_[int(8)] = _dafny.SeqWithoutIsStrInference([p0])
                nw56_[int(9)] = d_374_v54_
                d_387_v60_ = nw56_
                d_388_v61_: _dafny.MultiSet
                d_388_v61_ = _dafny.MultiSet([d_386_v59_])
                rhs76_ = (d_386_v59_ if True else (d_386_v59_).set(default__.safeIndex(-229, len(d_386_v59_)), (d_388_v61_).cardinality))
                rhs77_ = d_381_cf3_
                rhs78_ = d_387_v60_
                rhs79_ = (0) - (p1)
                lhs46_ = globalState
                d_386_v59_ = rhs76_
                d_381_cf3_ = rhs77_
                d_387_v60_ = rhs78_
                lhs46_.f1 = rhs79_
                d_389_v62_: _dafny.Map
                d_389_v62_ = _dafny.Map({False: (self).f3})
                d_390_v63_: C0
                nw57_ = C0()
                nw57_.ctor__(D0_DC0(False), d_389_v62_, d_319_v16_, d_381_cf3_)
                d_390_v63_ = nw57_
                d_391_v64_: _dafny.Seq
                d_391_v64_ = _dafny.SeqWithoutIsStrInference([d_390_v63_])
                d_369_v49_ = (d_369_v49_).set(default__.safeIndex(len((d_391_v64_) + (d_391_v64_)), len(d_369_v49_)), _dafny.CodePoint('b'))
                d_377_v57_ = d_377_v57_
                nw58_ = _dafny.Array(None, 13)
                nw58_[int(0)] = (d_369_v49_)[default__.safeIndex(d_381_cf3_, len(d_369_v49_))]
                nw58_[int(1)] = (self).f5
                nw58_[int(2)] = (self).f5
                nw58_[int(3)] = (self).f5
                nw58_[int(4)] = (self).f5
                nw58_[int(5)] = (self).f5
                nw58_[int(6)] = (self).f5
                nw58_[int(7)] = (d_370_v50_)[default__.safeIndex(496, (d_370_v50_).length(0))]
                nw58_[int(8)] = (d_370_v50_)[default__.safeIndex(496, (d_370_v50_).length(0))]
                nw58_[int(9)] = (d_370_v50_)[default__.safeIndex(496, (d_370_v50_).length(0))]
                nw58_[int(10)] = (d_370_v50_)[default__.safeIndex(496, (d_370_v50_).length(0))]
                nw58_[int(11)] = _dafny.CodePoint('b')
                nw58_[int(12)] = (d_370_v50_)[default__.safeIndex(496, (d_370_v50_).length(0))]
                d_370_v50_ = nw58_
            elif source6_.is_DC2:
                d_392___mcc_h3_ = source6_.cf4
                d_393___mcc_h4_ = source6_.cf5
                d_394___mcc_h5_ = source6_.cf6
                d_395_cf6_ = d_394___mcc_h5_
                d_396_cf5_ = d_393___mcc_h4_
                d_397_cf4_ = d_392___mcc_h3_
                d_398_v65_: _dafny.Map
                d_398_v65_ = _dafny.Map({p0: (self).f3})
                d_399_v66_: C0
                nw59_ = C0()
                nw59_.ctor__(D0_DC0(p0), d_398_v65_, d_319_v16_, p1)
                d_399_v66_ = nw59_
                d_400_v67_: _dafny.Array
                def lambda18_(d_401_v57_):
                    def lambda19_(d_402_i13_):
                        return (_dafny.SeqWithoutIsStrInference([d_401_v57_ for d_403_i14_ in range(default__.abs(10))]) if True else _dafny.SeqWithoutIsStrInference([d_401_v57_, d_401_v57_]))

                    return lambda19_

                init10_ = lambda18_(d_377_v57_)
                nw60_ = _dafny.Array(None, 13)
                for i0_10_ in range(nw60_.length(0)):
                    nw60_[i0_10_] = init10_(i0_10_)
                d_400_v67_ = nw60_
                d_404_v68_: D1
                d_404_v68_ = D1_DC5(d_371_v51_, d_371_v51_, p1)
                rhs80_ = p1
                rhs81_ = d_400_v67_
                rhs82_ = default__.safeDivisionInt((0) - (p1), len((d_397_cf4_) + ((d_397_cf4_).set(default__.safeIndex(p2, len(d_397_cf4_)), (d_370_v50_)[default__.safeIndex(496, (d_370_v50_).length(0))]))))
                rhs83_ = (default__.safeModuloInt(p1, 856)) >= ((d_404_v68_).cf11)
                lhs47_ = globalState
                lhs48_ = globalState
                lhs47_.f1 = rhs80_
                d_400_v67_ = rhs81_
                lhs48_.f1 = rhs82_
                d_371_v51_ = rhs83_
                d_405_v69_: _dafny.Seq
                d_405_v69_ = _dafny.SeqWithoutIsStrInference([default__.fm5(d_371_v51_, globalState)])
                d_406_v70_: _dafny.Map
                d_406_v70_ = _dafny.Map({default__.fm5(True, globalState): -217})
                d_407_v71_: _dafny.Array
                nw61_ = _dafny.Array(None, 19)
                nw61_[int(0)] = d_371_v51_
                nw61_[int(1)] = ((self).f4) < (d_395_cf6_)
                nw61_[int(2)] = d_371_v51_
                nw61_[int(3)] = (d_405_v69_) <= (d_405_v69_)
                nw61_[int(4)] = (p1) <= (len(d_373_v53_))
                nw61_[int(5)] = p0
                nw61_[int(6)] = (d_373_v53_).ispropersubset(d_373_v53_)
                nw61_[int(7)] = p0
                nw61_[int(8)] = (d_371_v51_) == (d_371_v51_)
                nw61_[int(9)] = p0
                nw61_[int(10)] = p0
                nw61_[int(11)] = d_371_v51_
                nw61_[int(12)] = (d_374_v54_)[default__.safeIndex(d_395_cf6_, len(d_374_v54_))]
                nw61_[int(13)] = True
                nw61_[int(14)] = d_371_v51_
                nw61_[int(15)] = (((d_406_v70_)[(0) - (d_395_cf6_)] if ((0) - (d_395_cf6_)) in (d_406_v70_) else (0) - ((self).f4))) != ((self).f4)
                nw61_[int(16)] = (_dafny.SeqWithoutIsStrInference([(d_370_v50_)[default__.safeIndex(496, (d_370_v50_).length(0))] for d_408_i15_ in range(default__.abs(752))])) <= (d_369_v49_)
                nw61_[int(17)] = not(d_371_v51_)
                nw61_[int(18)] = p0
                d_407_v71_ = nw61_
                d_407_v71_ = d_407_v71_
                (globalState).f1 = p1
            elif source6_.is_DC0:
                d_409___mcc_h6_ = source6_.cf0
                d_410_cf0_ = d_409___mcc_h6_
                d_411_v72_: D1
                d_411_v72_ = D1_DC4(d_319_v16_)
                d_412_v73_: _dafny.Map
                d_412_v73_ = _dafny.Map({not(d_410_cf0_): default__.fm5(True, globalState)})
                rhs84_ = (((0) - (len(d_412_v73_)) if p0 else (self).f4)) + ((self).f4)
                rhs85_ = d_411_v72_
                lhs49_ = globalState
                lhs49_.f1 = rhs84_
                d_411_v72_ = rhs85_
                d_413_v74_: D0
                d_413_v74_ = D0_DC0(d_371_v51_)
                d_414_v75_: _dafny.Map
                d_414_v75_ = _dafny.Map({d_410_cf0_: (self).f3})
                d_415_v76_: C0
                nw62_ = C0()
                nw62_.ctor__(d_413_v74_, d_414_v75_, (self).f3, p2)
                d_415_v76_ = nw62_
                (globalState).f1 = (self).f4
                rhs86_ = p1
                rhs87_ = default__.safeModuloInt(p1, p1)
                lhs50_ = globalState
                lhs51_ = globalState
                lhs50_.f1 = rhs86_
                lhs51_.f1 = rhs87_
            elif True:
                d_416___mcc_h7_ = source6_.cf7
                d_417_cf7_ = d_416___mcc_h7_
                (globalState).f1 = p2
                d_418_v77_: _dafny.Map
                d_418_v77_ = _dafny.Map({(self).f4: (self).f4})
                d_419_v78_: _dafny.Map
                d_419_v78_ = _dafny.Map({(d_371_v51_) and (False): (d_369_v49_)[default__.safeIndex(len(d_418_v77_), len(d_369_v49_))]})
                d_419_v78_ = (d_419_v78_).set(d_371_v51_, (self).f5)
                d_420_v79_: _dafny.Map
                d_420_v79_ = _dafny.Map({(self).f3: p1})
                d_420_v79_ = (d_420_v79_).set(d_319_v16_, p2)
                d_421_v80_: _dafny.Array
                nw63_ = _dafny.Array(False, 18)
                d_421_v80_ = nw63_
                d_421_v80_ = (self).f6
            d_422_v81_: _dafny.Seq
            d_422_v81_ = _dafny.SeqWithoutIsStrInference([d_370_v50_, d_370_v50_])
            rhs88_ = (d_422_v81_) < ((d_422_v81_) + (_dafny.SeqWithoutIsStrInference([(d_422_v81_)[default__.safeIndex(len(d_374_v54_), len(d_422_v81_))], d_370_v50_, d_370_v50_])))
            rhs89_ = (self).f4
            rhs90_ = not((d_369_v49_) != (d_369_v49_))
            lhs52_ = globalState
            d_371_v51_ = rhs88_
            lhs52_.f1 = rhs89_
            d_371_v51_ = rhs90_
            d_371_v51_ = d_371_v51_
        d_423_v82_: _dafny.Map
        d_423_v82_ = _dafny.Map({p0: (0) - (p2)})
        d_424_v83_: D0
        d_424_v83_ = D0_DC1((self).f4, (0) - (p2), len(d_423_v82_))
        d_425_v84_: _dafny.Seq
        d_425_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uohkfjswt"))
        hi2_ = default__.safeModuloInt((self).f4, default__.fm5(p0, globalState))
        for d_426_i16_ in range(default__.safeModuloInt((d_424_v83_).cf2, len(d_425_v84_)), hi2_):
            d_427_v85_: _dafny.Seq
            d_427_v85_ = _dafny.SeqWithoutIsStrInference([p1])
            d_428_v86_: _dafny.Array
            nw64_ = _dafny.Array(None, 22)
            nw64_[int(0)] = default__.fm3((self).f5, (self).f4, globalState)
            nw64_[int(1)] = _dafny.CodePoint('f')
            nw64_[int(2)] = (self).f5
            nw64_[int(3)] = (self).f5
            nw64_[int(4)] = default__.fm3((self).f5, p1, globalState)
            nw64_[int(5)] = (self).f5
            nw64_[int(6)] = (self).f5
            nw64_[int(7)] = (self).f5
            nw64_[int(8)] = _dafny.CodePoint('a')
            nw64_[int(9)] = (self).f5
            nw64_[int(10)] = (self).f5
            nw64_[int(11)] = (self).f5
            nw64_[int(12)] = (self).f5
            nw64_[int(13)] = (self).f5
            nw64_[int(14)] = _dafny.CodePoint('j')
            nw64_[int(15)] = (self).f5
            nw64_[int(16)] = (self).f5
            nw64_[int(17)] = (self).f5
            nw64_[int(18)] = (d_425_v84_)[default__.safeIndex(len(d_427_v85_), len(d_425_v84_))]
            nw64_[int(19)] = (self).f5
            nw64_[int(20)] = (self).f5
            nw64_[int(21)] = (self).f5
            d_428_v86_ = nw64_
            d_429_v87_: _dafny.Seq
            d_429_v87_ = _dafny.SeqWithoutIsStrInference([d_428_v86_, d_428_v86_])
            d_430_v88_: _dafny.Set
            d_430_v88_ = _dafny.Set({(self).f5})
            source7_ = D0_DC2(d_425_v84_, len(d_429_v87_), default__.safeDivisionInt(p2, len(d_430_v88_)))
            if source7_.is_DC1:
                d_431___mcc_h8_ = source7_.cf1
                d_432___mcc_h9_ = source7_.cf2
                d_433___mcc_h10_ = source7_.cf3
                d_434_cf3_ = d_433___mcc_h10_
                d_435_cf2_ = d_432___mcc_h9_
                d_436_cf1_ = d_431___mcc_h8_
                d_437_v89_: bool
                d_437_v89_ = False
                d_437_v89_ = not(not(not(d_437_v89_)))
                d_438_v90_: _dafny.Set
                d_438_v90_ = _dafny.Set({d_437_v89_, d_437_v89_, True})
                d_438_v90_ = (d_438_v90_) | (d_438_v90_)
                d_439_v91_: D0
                d_439_v91_ = D0_DC0(p0)
                d_440_v92_: _dafny.Map
                d_440_v92_ = _dafny.Map({d_437_v89_: d_319_v16_})
                d_441_v93_: T0
                nw65_ = C0()
                nw65_.ctor__(d_439_v91_, d_440_v92_, d_319_v16_, d_434_cf3_)
                d_441_v93_ = nw65_
                d_442_v94_: _dafny.Array
                nw66_ = _dafny.Array(None, 21)
                nw66_[int(0)] = d_441_v93_
                nw66_[int(1)] = d_441_v93_
                nw66_[int(2)] = d_441_v93_
                nw66_[int(3)] = d_441_v93_
                nw66_[int(4)] = d_441_v93_
                nw66_[int(5)] = d_441_v93_
                nw66_[int(6)] = d_441_v93_
                nw66_[int(7)] = d_441_v93_
                nw66_[int(8)] = d_441_v93_
                nw66_[int(9)] = d_441_v93_
                nw66_[int(10)] = d_441_v93_
                nw66_[int(11)] = d_441_v93_
                nw66_[int(12)] = d_441_v93_
                nw66_[int(13)] = d_441_v93_
                nw66_[int(14)] = d_441_v93_
                nw66_[int(15)] = d_441_v93_
                nw66_[int(16)] = d_441_v93_
                nw66_[int(17)] = d_441_v93_
                nw66_[int(18)] = d_441_v93_
                nw66_[int(19)] = d_441_v93_
                nw66_[int(20)] = d_441_v93_
                d_442_v94_ = nw66_
                index71_ = default__.safeIndex(351, (d_442_v94_).length(0))
                (d_442_v94_)[index71_] = d_441_v93_
                index72_ = default__.safeIndex(351, (d_442_v94_).length(0))
                (d_442_v94_)[index72_] = d_441_v93_
                d_443_v95_: _dafny.Map
                d_443_v95_ = _dafny.Map({d_437_v89_: d_437_v89_})
                d_444_v96_: C0
                nw67_ = C0()
                nw67_.ctor__(d_439_v91_, d_440_v92_, (self).f3, len(d_443_v95_))
                d_444_v96_ = nw67_
                d_445_v97_: _dafny.Map
                d_445_v97_ = _dafny.Map({d_444_v96_: (d_441_v93_).f3})
                d_446_v98_: C0
                nw68_ = C0()
                nw68_.ctor__(D0_DC0(p0), d_440_v92_, ((d_445_v97_)[d_444_v96_] if (d_444_v96_) in (d_445_v97_) else d_319_v16_), -693)
                d_446_v98_ = nw68_
            elif source7_.is_DC2:
                d_447___mcc_h11_ = source7_.cf4
                d_448___mcc_h12_ = source7_.cf5
                d_449___mcc_h13_ = source7_.cf6
                d_450_cf6_ = d_449___mcc_h13_
                d_451_cf5_ = d_448___mcc_h12_
                d_452_cf4_ = d_447___mcc_h11_
                d_453_v99_: _dafny.Array
                def lambda20_(d_454_i17_):
                    return _dafny.MultiSet((D2_DC7(_dafny.SeqWithoutIsStrInference([True]))).cf12)

                init11_ = lambda20_
                nw69_ = _dafny.Array(None, 27)
                for i0_11_ in range(nw69_.length(0)):
                    nw69_[i0_11_] = init11_(i0_11_)
                d_453_v99_ = nw69_
                d_455_v100_: _dafny.Map
                d_455_v100_ = _dafny.Map({(0) - (default__.safeModuloInt(d_450_cf6_, d_451_cf5_)): d_453_v99_})
                d_455_v100_ = (d_455_v100_).set(d_450_cf6_, d_453_v99_)
                d_456_v101_: _dafny.MultiSet
                d_456_v101_ = _dafny.MultiSet([(len(d_452_cf4_) if not(p0) else d_451_cf5_), default__.safeModuloInt(d_426_i16_, d_451_cf5_), d_450_cf6_, (d_451_cf5_) + (p2)])
                (globalState).f1 = (d_456_v101_).cardinality
                d_457_v102_: bool
                d_457_v102_ = True
                d_458_v103_: _dafny.Map
                d_458_v103_ = _dafny.Map({(self).f5: d_457_v102_})
                d_457_v102_ = not(not(((d_458_v103_) | (d_458_v103_)) != ((d_458_v103_) | (_dafny.Map({_dafny.CodePoint('o'): p0})))))
                d_459_v104_: D1
                d_459_v104_ = D1_DC4((self).f3)
                d_460_v105_: _dafny.Seq
                d_460_v105_ = _dafny.SeqWithoutIsStrInference([d_457_v102_])
                d_461_v106_: D1
                d_461_v106_ = D1_DC5(False, p0, len(d_460_v105_))
                d_459_v104_ = D1_DC4((d_319_v16_ if (d_461_v106_).cf10 else d_319_v16_))
            elif source7_.is_DC0:
                d_462___mcc_h14_ = source7_.cf0
                d_463_cf0_ = d_462___mcc_h14_
                d_463_cf0_ = True
                (globalState).f1 = d_426_i16_
                d_425_v84_ = d_425_v84_
                d_464_v107_: _dafny.Map
                d_464_v107_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f5 for d_465_i18_ in range(default__.abs(993))]): d_426_i16_})
                d_464_v107_ = (d_464_v107_).set(d_425_v84_, default__.fm5(p0, globalState))
            elif True:
                d_466___mcc_h15_ = source7_.cf7
                d_467_cf7_ = d_466___mcc_h15_
                d_468_v108_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.Seq({}), 8)
                d_468_v108_ = nw70_
                d_469_v109_: D1
                d_469_v109_ = D1_DC4(d_319_v16_)
                d_470_v110_: _dafny.Seq
                d_470_v110_ = _dafny.SeqWithoutIsStrInference([(d_469_v109_).cf8])
                index73_ = default__.safeIndex(211, (d_468_v108_).length(0))
                (d_468_v108_)[index73_] = d_470_v110_
                d_471_v111_: bool
                d_471_v111_ = False
                d_472_v112_: _dafny.Seq
                d_472_v112_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, (self).f6, (self).f6, (self).f6])
                index74_ = default__.safeIndex(211, (d_468_v108_).length(0))
                rhs91_ = 957
                rhs92_ = len((d_472_v112_).set(default__.safeIndex(p2, len(d_472_v112_)), (self).f6))
                rhs93_ = d_470_v110_
                rhs94_ = not((default__.fm0(len(d_427_v85_), d_426_i16_, globalState) if not (d_471_v111_) or (d_471_v111_) else d_471_v111_))
                lhs53_ = globalState
                lhs54_ = globalState
                lhs55_ = d_468_v108_
                lhs56_ = default__.safeIndex(211, (d_468_v108_).length(0))
                lhs53_.f1 = rhs91_
                lhs54_.f1 = rhs92_
                lhs55_[lhs56_] = rhs93_
                d_471_v111_ = rhs94_
                d_471_v111_ = (d_471_v111_) == ((D1_DC5(d_471_v111_, d_471_v111_, len(d_423_v82_))).cf9)
                d_471_v111_ = d_471_v111_
                d_473_v113_: str
                d_473_v113_ = _dafny.CodePoint('j')
                d_473_v113_ = (self).f5
            d_474_v114_: D0
            d_474_v114_ = D0_DC0(p0)
            d_475_v115_: _dafny.Map
            d_475_v115_ = _dafny.Map({p0: d_319_v16_})
            d_476_v116_: _dafny.Seq
            d_476_v116_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            pat_let_tv5_ = p0
            d_477_v117_: C0
            nw71_ = C0()
            def iife33_(_pat_let13_0):
                def iife34_(d_478_dt__update__tmp_h1_):
                    def iife35_(_pat_let14_0):
                        def iife36_(d_479_dt__update_hcf0_h0_):
                            return D0_DC0(d_479_dt__update_hcf0_h0_)
                        return iife36_(_pat_let14_0)
                    return iife35_(pat_let_tv5_)
                return iife34_(_pat_let13_0)
            nw71_.ctor__(iife33_(d_474_v114_), ((d_475_v115_).set(p0, d_319_v16_)) | ((d_475_v115_).set((d_476_v116_)[default__.safeIndex(default__.fm5(p0, globalState), len(d_476_v116_))], (self).f3)), (self).f3, len(d_425_v84_))
            d_477_v117_ = nw71_
            d_480_v118_: bool
            d_480_v118_ = True
            d_481_v119_: _dafny.Seq
            d_481_v119_ = _dafny.SeqWithoutIsStrInference([d_425_v84_])
            d_480_v118_ = ((d_481_v119_)[default__.safeIndex(198, len(d_481_v119_))]) == ((d_425_v84_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjsywnkxd"))))
            d_482_v120_: _dafny.Map
            d_482_v120_ = _dafny.Map({_dafny.CodePoint('c'): d_426_i16_})
            d_483_v121_: _dafny.MultiSet
            d_483_v121_ = _dafny.MultiSet([d_482_v120_, default__.fm9(globalState), d_482_v120_])
            d_483_v121_ = _dafny.MultiSet((default__.fm10(D1_DC6(), (d_425_v84_)[default__.safeIndex(d_426_i16_, len(d_425_v84_))], (0) - (d_426_i16_), not(d_480_v118_), globalState)).set(default__.safeIndex(p2, len(default__.fm10(D1_DC6(), (d_425_v84_)[default__.safeIndex(d_426_i16_, len(d_425_v84_))], (0) - (d_426_i16_), not(d_480_v118_), globalState))), d_482_v120_))
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_319_v16_).length(0)):
            d_484_i19_: int = guard_loop_2_
            if (True) and (((0) <= (d_484_i19_)) and ((d_484_i19_) < ((d_319_v16_).length(0)))):
                (d_319_v16_)[(d_484_i19_)] = default__.safeModuloInt(d_484_i19_, ((self).f4 if p0 else p2))
        d_485_v122_: _dafny.Seq
        d_485_v122_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f5 for d_486_i20_ in range(default__.abs(870))])])
        d_487_v123_: D0
        d_487_v123_ = D0_DC2(d_425_v84_, (self).f4, (0) - (p2))
        r0 = ((d_485_v122_)[default__.safeIndex((self).f4, len(d_485_v122_))]) + ((d_487_v123_).cf4)
        return r0

    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
