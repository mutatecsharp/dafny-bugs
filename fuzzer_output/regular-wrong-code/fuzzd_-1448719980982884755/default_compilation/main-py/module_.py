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
        return (D1_DC2(121)).cf2

    @staticmethod
    def fm2(globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(686, 805):
                d_0_v0_: int = compr_0_
                if ((686) <= (d_0_v0_)) and ((d_0_v0_) < (805)):
                    coll0_ = coll0_.union(_dafny.Set([(d_0_v0_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({323: True})])))]))
            return _dafny.Set(coll0_)
        return (not(not(not((_dafny.MultiSet([-776])).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([iife0_()
]))).cardinality for d_1_i0_ in range(default__.abs(448))]))))))) and (not(not((False if False else False))))

    @staticmethod
    def fm3(p0, globalState):
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifnkwt"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hammgsnta"))), ((0) - (-352)) <= ((0) - (len(_dafny.Set({False})))), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False]))])).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([277 for d_2_i0_ in range(default__.abs(240))]))), False, True})

    @staticmethod
    def fm4(p0, p1, globalState):
        source0_ = D4_DC11(823)
        if source0_.is_DC9:
            d_3___mcc_h0_ = source0_.cf10
            d_4___mcc_h1_ = source0_.cf11
            d_5___mcc_h2_ = source0_.cf12
            d_6_cf12_ = d_5___mcc_h2_
            d_7_cf11_ = d_4___mcc_h1_
            d_8_cf10_ = d_3___mcc_h0_
            return _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqlkl")))})
        elif source0_.is_DC10:
            d_9___mcc_h3_ = source0_.cf13
            d_10___mcc_h4_ = source0_.cf14
            d_11___mcc_h5_ = source0_.cf15
            d_12___mcc_h6_ = source0_.cf16
            d_13___mcc_h7_ = source0_.cf17
            d_14_cf17_ = d_13___mcc_h7_
            d_15_cf16_ = d_12___mcc_h6_
            d_16_cf15_ = d_11___mcc_h5_
            d_17_cf14_ = d_10___mcc_h4_
            d_18_cf13_ = d_9___mcc_h3_
            return (_dafny.Set({d_16_cf15_})).intersection(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_17_cf14_ for d_19_i0_ in range(default__.abs(611))])), d_16_cf15_}))
        elif source0_.is_DC11:
            d_20___mcc_h8_ = source0_.cf18
            d_21_cf18_ = d_20___mcc_h8_
            return _dafny.Set({d_21_cf18_, d_21_cf18_, d_21_cf18_, (0) - (d_21_cf18_), (0) - ((0) - (d_21_cf18_))})
        elif source0_.is_DC8:
            d_22___mcc_h9_ = source0_.cf9
            d_23_cf9_ = d_22___mcc_h9_
            return _dafny.Set({-944, 396, len(_dafny.Map({d_23_cf9_: 228}))})
        elif True:
            d_24___mcc_h10_ = source0_.cf19
            d_25_cf19_ = d_24___mcc_h10_
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(73, -142):
                    d_26_v0_: int = compr_1_
                    if ((73) <= (d_26_v0_)) and ((d_26_v0_) < (-142)):
                        coll1_ = coll1_.union(_dafny.Set([default__.safeModuloInt(d_26_v0_, 972)]))
                return _dafny.Set(coll1_)
            return iife1_()
            

    @staticmethod
    def fm5(globalState):
        if not((len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([True, True, True, True, True]))), 207]))) >= (350)):
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(75, 606):
                    d_27_v0_: int = compr_2_
                    if ((75) <= (d_27_v0_)) and ((d_27_v0_) < (606)):
                        coll2_[default__.safeDivisionInt(d_27_v0_, 436)] = True
                return _dafny.Map(coll2_)
            return (_dafny.Map({True: iife2_()
            })) | (_dafny.Map({False: _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkytbt"))): False})}))
        elif True:
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in (_dafny.MultiSet([(_dafny.MultiSet([-230, len(_dafny.Map({True: -856}))])).cardinality, len(_dafny.Map({514: False}))])).Elements:
                    d_28_v1_: int = compr_3_
                    if (d_28_v1_) in (_dafny.MultiSet([(_dafny.MultiSet([-230, len(_dafny.Map({True: -856}))])).cardinality, len(_dafny.Map({514: False}))])):
                        coll3_[default__.safeModuloInt(d_28_v1_, (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([340]))]))))] = False
                return _dafny.Map(coll3_)
            return (_dafny.Map({True: iife3_()
            })) | (_dafny.Map({False: _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qevan"))): False})}))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('y') if not(False) else _dafny.CodePoint('g')) for d_29_i0_ in range(default__.abs(165))])

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([914, (0) - (-417)])) + ((_dafny.SeqWithoutIsStrInference([400])) + (_dafny.SeqWithoutIsStrInference([553])))

    @staticmethod
    def fm8(p0, p1, globalState):
        return D3_DC6(522)

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        if False:
            return _dafny.CodePoint('u')
        elif True:
            return _dafny.CodePoint('q')

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return D2_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gywvd")))

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(_dafny.MultiSet([False])).cardinality: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({817: len(_dafny.SeqWithoutIsStrInference([147 for d_30_i0_ in range(default__.abs(583))]))}))]))}))])) + (_dafny.SeqWithoutIsStrInference([-387, (_dafny.MultiSet([len(_dafny.Map({False: _dafny.CodePoint('x')})), 868, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oendfkfg")))]))])).cardinality])), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htm")))])) for d_31_i1_ in range(default__.abs(559))]), _dafny.SeqWithoutIsStrInference([766])]))

    @staticmethod
    def fm12(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True]))) + (_dafny.SeqWithoutIsStrInference([not(not(True))]))

    @staticmethod
    def fm13(p0, globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): True})).keys.Elements:
                d_32_v0_: _dafny.Seq = compr_4_
                if (d_32_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): True})):
                    coll4_ = coll4_.union(_dafny.Set([d_32_v0_]))
            return _dafny.Set(coll4_)
        return (((D8_DC19(_dafny.Set({_dafny.SeqWithoutIsStrInference([True, False, False, True]), _dafny.SeqWithoutIsStrInference([True])}))).cf26) | (iife4_()
        )).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([False, True, True, True, not(True)]), _dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([True, True, False])}))

    @staticmethod
    def fm14(globalState):
        return ((_dafny.Map({True: False})) | (_dafny.Map({True: False}))) | (_dafny.Map({True: not(False)}))

    @staticmethod
    def m0(globalState):
        d_33_v0_: bool
        d_33_v0_ = False
        d_34_v1_: C0
        nw0_ = C0()
        nw0_.ctor__(d_33_v0_)
        d_34_v1_ = nw0_
        d_35_v2_: _dafny.Array
        def lambda0_(d_36_i0_):
            return _dafny.MultiSet([497, len(_dafny.Set({-320, -68, 69})), -332])

        init0_ = lambda0_
        nw1_ = _dafny.Array(None, 5)
        for i0_0_ in range(nw1_.length(0)):
            nw1_[i0_0_] = init0_(i0_0_)
        d_35_v2_ = nw1_
        d_37_v3_: _dafny.Map
        d_37_v3_ = _dafny.Map({d_34_v1_: d_35_v2_})
        d_37_v3_ = (d_37_v3_).set(d_34_v1_, d_35_v2_)
        d_38_v4_: int
        d_38_v4_ = 570
        d_39_v5_: str
        d_39_v5_ = _dafny.CodePoint('h')
        pat_let_tv0_ = d_34_v1_
        pat_let_tv1_ = d_34_v1_
        def lambda1_(source1_):
            if source1_.is_DC4:
                d_40___mcc_h0_ = source1_.cf4
                d_41___mcc_h1_ = source1_.cf5
                d_42_cf5_ = d_41___mcc_h1_
                d_43_cf4_ = d_40___mcc_h0_
                return pat_let_tv0_.f15
            elif True:
                d_44___mcc_h2_ = source1_.cf3
                d_45_cf3_ = d_44___mcc_h2_
                return not(pat_let_tv1_.f15)

        if lambda1_(default__.fm10(d_38_v4_, d_39_v5_, d_33_v0_, globalState)):
            d_46_v6_: _dafny.Array
            def lambda2_(d_47_i1_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))

            init1_ = lambda2_
            nw2_ = _dafny.Array(None, 20)
            for i0_1_ in range(nw2_.length(0)):
                nw2_[i0_1_] = init1_(i0_1_)
            d_46_v6_ = nw2_
            d_46_v6_ = d_46_v6_
            d_48_v7_: _dafny.Seq
            d_48_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxcmd"))
            d_49_v8_: _dafny.Array
            def lambda3_(d_50_i3_):
                return True

            init2_ = lambda3_
            nw3_ = _dafny.Array(None, 27)
            for i0_2_ in range(nw3_.length(0)):
                nw3_[i0_2_] = init2_(i0_2_)
            d_49_v8_ = nw3_
            d_51_v9_: D4
            d_51_v9_ = D4_DC9(d_49_v8_, d_34_v1_, d_33_v0_)
            d_52_v10_: _dafny.Seq
            d_52_v10_ = _dafny.SeqWithoutIsStrInference([d_34_v1_.f15, d_34_v1_.f15, d_34_v1_.f15])
            d_53_v11_: _dafny.MultiSet
            d_53_v11_ = _dafny.MultiSet([d_34_v1_.f15])
            d_54_v12_: _dafny.Set
            d_54_v12_ = _dafny.Set({d_34_v1_, d_34_v1_})
            d_55_v13_: _dafny.Set
            d_55_v13_ = _dafny.Set({d_34_v1_.f15})
            pat_let_tv2_ = d_34_v1_
            pat_let_tv3_ = d_49_v8_
            d_56_v14_: _dafny.Array
            nw4_ = _dafny.Array(None, 28)
            nw4_[int(0)] = (d_48_v7_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_57_i2_ in range(default__.abs(391))]))
            nw4_[int(1)] = d_34_v1_.f15
            nw4_[int(2)] = d_34_v1_.f15
            def iife5_(_pat_let0_0):
                def iife6_(d_58_dt__update__tmp_h0_):
                    def iife7_(_pat_let1_0):
                        def iife8_(d_59_dt__update_hcf11_h0_):
                            def iife9_(_pat_let2_0):
                                def iife10_(d_60_dt__update_hcf10_h0_):
                                    return D4_DC9(d_60_dt__update_hcf10_h0_, d_59_dt__update_hcf11_h0_, (d_58_dt__update__tmp_h0_).cf12)
                                return iife10_(_pat_let2_0)
                            return iife9_(pat_let_tv3_)
                        return iife8_(_pat_let1_0)
                    return iife7_(pat_let_tv2_)
                return iife6_(_pat_let0_0)
            nw4_[int(3)] = (iife5_(d_51_v9_)).cf12
            nw4_[int(4)] = ((0) - (len(_dafny.Set({d_33_v0_, d_34_v1_.f15, d_34_v1_.f15, d_34_v1_.f15, d_34_v1_.f15})))) == (len(d_48_v7_))
            nw4_[int(5)] = (d_39_v5_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "difl")))
            nw4_[int(6)] = True
            nw4_[int(7)] = d_34_v1_.f15
            nw4_[int(8)] = d_34_v1_.f15
            nw4_[int(9)] = (d_34_v1_.f15) and ((d_52_v10_)[default__.safeIndex(default__.fm0(d_33_v0_, d_38_v4_, d_34_v1_.f15, globalState), len(d_52_v10_))])
            nw4_[int(10)] = (d_53_v11_).issubset((d_53_v11_).set(d_33_v0_, default__.abs(d_38_v4_)))
            nw4_[int(11)] = d_33_v0_
            nw4_[int(12)] = d_34_v1_.f15
            nw4_[int(13)] = d_34_v1_.f15
            nw4_[int(14)] = (not(d_33_v0_)) or (True)
            nw4_[int(15)] = default__.fm2(globalState)
            nw4_[int(16)] = d_34_v1_.f15
            nw4_[int(17)] = d_33_v0_
            nw4_[int(18)] = (d_54_v12_).isdisjoint(d_54_v12_)
            nw4_[int(19)] = d_33_v0_
            nw4_[int(20)] = (d_55_v13_) == (d_55_v13_)
            nw4_[int(21)] = not (d_34_v1_.f15) or (d_34_v1_.f15)
            nw4_[int(22)] = d_34_v1_.f15
            nw4_[int(23)] = d_34_v1_.f15
            nw4_[int(24)] = d_34_v1_.f15
            nw4_[int(25)] = d_33_v0_
            nw4_[int(26)] = d_34_v1_.f15
            nw4_[int(27)] = d_34_v1_.f15
            d_56_v14_ = nw4_
            index0_ = default__.safeIndex(71, (d_49_v8_).length(0))
            (d_49_v8_)[index0_] = d_33_v0_
            index1_ = default__.safeIndex(71, (d_49_v8_).length(0))
            (d_49_v8_)[index1_] = (True if not((332) == (d_38_v4_)) else (d_34_v1_.f15) or (d_33_v0_))
            index2_ = default__.safeIndex(71, (d_49_v8_).length(0))
            (d_49_v8_)[index2_] = default__.fm2(globalState)
            d_38_v4_ = d_38_v4_
            d_61_v15_: _dafny.Array
            nw5_ = _dafny.Array(int(0), 28)
            d_61_v15_ = nw5_
            index3_ = default__.safeIndex(842, (d_61_v15_).length(0))
            (d_61_v15_)[index3_] = d_38_v4_
            d_62_v16_: _dafny.Seq
            d_62_v16_ = _dafny.SeqWithoutIsStrInference([d_34_v1_])
            d_63_v17_: _dafny.MultiSet
            d_63_v17_ = _dafny.MultiSet([d_62_v16_, d_62_v16_, d_62_v16_])
            d_64_v18_: _dafny.Seq
            d_64_v18_ = _dafny.SeqWithoutIsStrInference([d_62_v16_])
            index4_ = default__.safeIndex(842, (d_61_v15_).length(0))
            rhs0_ = d_61_v15_
            rhs1_ = d_38_v4_
            rhs2_ = ((d_63_v17_).intersection(_dafny.MultiSet(d_64_v18_))).issubset((d_63_v17_).intersection(d_63_v17_))
            rhs3_ = d_38_v4_
            lhs0_ = d_61_v15_
            lhs1_ = default__.safeIndex(842, (d_61_v15_).length(0))
            lhs2_ = globalState
            lhs3_ = globalState
            d_61_v15_ = rhs0_
            lhs0_[lhs1_] = rhs1_
            lhs2_.f8 = rhs2_
            lhs3_.f0 = rhs3_
        elif True:
            d_65_v19_: _dafny.Map
            d_65_v19_ = _dafny.Map({d_33_v0_: default__.fm2(globalState)})
            d_66_v20_: _dafny.Seq
            d_66_v20_ = _dafny.SeqWithoutIsStrInference([d_65_v19_])
            d_67_v22_: D1
            d_67_v22_ = D1_DC2(d_38_v4_)
            def iife11_():
                coll5_ = _dafny.Set()
                compr_5_: _dafny.Map
                for compr_5_ in (d_66_v20_).Elements:
                    d_68_v21_: _dafny.Map = compr_5_
                    if (d_68_v21_) in (d_66_v20_):
                        coll5_ = coll5_.union(_dafny.Set([d_68_v21_]))
                return _dafny.Set(coll5_)
            (globalState).f1 = (len(iife11_()
            )) <= ((0) - (len(default__.fm6(d_38_v4_, d_34_v1_.f15, d_67_v22_, False, globalState))))
            d_69_v23_: _dafny.MultiSet
            d_69_v23_ = _dafny.MultiSet([False])
            d_70_v24_: _dafny.Seq
            d_70_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kv"))
            (globalState).f10 = (((d_69_v23_)[d_33_v0_] if (d_33_v0_) in (d_69_v23_) else len(d_70_v24_))) == (717)
            d_71_v25_: _dafny.Map
            d_71_v25_ = _dafny.Map({d_39_v5_: d_34_v1_})
            d_72_v26_: _dafny.Map
            d_72_v26_ = _dafny.Map({d_71_v25_: d_33_v0_})
            d_73_v27_: D5
            d_73_v27_ = D5_DC13(d_72_v26_)
            d_72_v26_ = (d_73_v27_).cf20
            d_34_v1_ = d_34_v1_
            d_65_v19_ = (d_65_v19_).set(True, d_34_v1_.f15)
        d_74_i4_: int
        d_74_i4_ = 0
        with _dafny.label("0"):
            while not(d_33_v0_):
                with _dafny.c_label("0"):
                    if (d_74_i4_) >= (100):
                        raise _dafny.Break("0")
                    d_74_i4_ = (d_74_i4_) + (1)
                    d_75_v28_: _dafny.Array
                    nw6_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                    d_75_v28_ = nw6_
                    d_76_v29_: _dafny.Seq
                    d_76_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brhkdmqxe"))
                    index5_ = default__.safeIndex(983, (d_75_v28_).length(0))
                    (d_75_v28_)[index5_] = d_76_v29_
                    index6_ = default__.safeIndex(983, (d_75_v28_).length(0))
                    (d_75_v28_)[index6_] = d_76_v29_
                    d_77_v30_: _dafny.Map
                    d_77_v30_ = _dafny.Map({d_33_v0_: d_39_v5_})
                    d_78_v31_: _dafny.Array
                    nw7_ = _dafny.Array(None, 22)
                    nw7_[int(0)] = d_39_v5_
                    nw7_[int(1)] = d_39_v5_
                    nw7_[int(2)] = d_39_v5_
                    nw7_[int(3)] = d_39_v5_
                    nw7_[int(4)] = default__.fm9(398, (d_75_v28_)[default__.safeIndex(983, (d_75_v28_).length(0))], d_38_v4_, globalState)
                    nw7_[int(5)] = d_39_v5_
                    nw7_[int(6)] = d_39_v5_
                    nw7_[int(7)] = d_39_v5_
                    nw7_[int(8)] = d_39_v5_
                    nw7_[int(9)] = d_39_v5_
                    nw7_[int(10)] = d_39_v5_
                    nw7_[int(11)] = d_39_v5_
                    nw7_[int(12)] = (_dafny.CodePoint('s') if d_33_v0_ else _dafny.CodePoint('l'))
                    nw7_[int(13)] = d_39_v5_
                    nw7_[int(14)] = d_39_v5_
                    nw7_[int(15)] = (_dafny.CodePoint('v') if d_33_v0_ else d_39_v5_)
                    nw7_[int(16)] = _dafny.CodePoint('r')
                    nw7_[int(17)] = _dafny.CodePoint('o')
                    nw7_[int(18)] = d_39_v5_
                    nw7_[int(19)] = (d_76_v29_)[default__.safeIndex(d_38_v4_, len(d_76_v29_))]
                    nw7_[int(20)] = d_39_v5_
                    nw7_[int(21)] = ((d_77_v30_)[d_33_v0_] if (d_33_v0_) in (d_77_v30_) else d_39_v5_)
                    d_78_v31_ = nw7_
                    index7_ = default__.safeIndex(452, (d_78_v31_).length(0))
                    (d_78_v31_)[index7_] = d_39_v5_
                    index8_ = default__.safeIndex(452, (d_78_v31_).length(0))
                    (d_78_v31_)[index8_] = d_39_v5_
                    d_79_v32_: C0
                    nw8_ = C0()
                    nw8_.ctor__(d_33_v0_)
                    d_79_v32_ = nw8_
                    index9_ = default__.safeIndex(452, (d_78_v31_).length(0))
                    (d_78_v31_)[index9_] = d_39_v5_
                    pass
            pass
        if d_33_v0_:
            d_80_v33_: _dafny.Array
            def lambda4_(d_81_v1_):
                def lambda5_(d_82_i5_):
                    return d_81_v1_.f15

                return lambda5_

            init3_ = lambda4_(d_34_v1_)
            nw9_ = _dafny.Array(None, 16)
            for i0_3_ in range(nw9_.length(0)):
                nw9_[i0_3_] = init3_(i0_3_)
            d_80_v33_ = nw9_
            d_83_v34_: _dafny.Map
            d_83_v34_ = _dafny.Map({d_38_v4_: (_dafny.SeqWithoutIsStrInference([d_34_v1_.f15, d_33_v0_, d_34_v1_.f15])).set(default__.safeIndex((0) - (d_38_v4_), len(_dafny.SeqWithoutIsStrInference([d_34_v1_.f15, d_33_v0_, d_34_v1_.f15]))), d_33_v0_)})
            d_84_v35_: _dafny.Seq
            d_84_v35_ = _dafny.SeqWithoutIsStrInference([d_33_v0_, d_34_v1_.f15])
            d_85_v36_: _dafny.Set
            d_85_v36_ = _dafny.Set({d_84_v35_})
            index10_ = default__.safeIndex(800, (d_80_v33_).length(0))
            (d_80_v33_)[index10_] = (((d_83_v34_)[(_dafny.MultiSet([not(d_34_v1_.f15), d_33_v0_])).cardinality] if ((_dafny.MultiSet([not(d_34_v1_.f15), d_33_v0_])).cardinality) in (d_83_v34_) else d_84_v35_)) not in (d_85_v36_)
            index11_ = default__.safeIndex(800, (d_80_v33_).length(0))
            (d_80_v33_)[index11_] = d_33_v0_
            d_86_v37_: _dafny.Seq
            d_86_v37_ = _dafny.SeqWithoutIsStrInference([d_38_v4_])
            d_87_v38_: _dafny.MultiSet
            d_87_v38_ = _dafny.MultiSet([d_86_v37_, d_86_v37_])
            d_88_v39_: _dafny.Seq
            d_88_v39_ = _dafny.SeqWithoutIsStrInference([d_86_v37_])
            d_89_v40_: _dafny.Map
            d_89_v40_ = _dafny.Map({d_34_v1_.f15: d_38_v4_})
            d_90_v41_: _dafny.Map
            d_90_v41_ = _dafny.Map({d_38_v4_: d_38_v4_})
            d_91_v42_: _dafny.Array
            nw10_ = _dafny.Array(None, 13)
            nw10_[int(0)] = d_87_v38_
            nw10_[int(1)] = d_87_v38_
            nw10_[int(2)] = _dafny.MultiSet(d_88_v39_)
            nw10_[int(3)] = d_87_v38_
            nw10_[int(4)] = d_87_v38_
            nw10_[int(5)] = d_87_v38_
            nw10_[int(6)] = _dafny.MultiSet(d_88_v39_)
            nw10_[int(7)] = default__.fm11(d_38_v4_, d_38_v4_, d_33_v0_, d_33_v0_, globalState)
            nw10_[int(8)] = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(d_89_v40_), len(d_90_v41_)]), d_86_v37_, _dafny.SeqWithoutIsStrInference([d_38_v4_]), (_dafny.SeqWithoutIsStrInference([d_38_v4_])).set(default__.safeIndex(d_38_v4_, len(_dafny.SeqWithoutIsStrInference([d_38_v4_]))), default__.fm0((d_80_v33_)[default__.safeIndex(800, (d_80_v33_).length(0))], d_38_v4_, d_34_v1_.f15, globalState)), d_86_v37_])
            nw10_[int(9)] = (d_87_v38_) - (d_87_v38_)
            nw10_[int(10)] = d_87_v38_
            nw10_[int(11)] = _dafny.MultiSet([d_86_v37_])
            nw10_[int(12)] = (d_87_v38_ if d_34_v1_.f15 else d_87_v38_)
            d_91_v42_ = nw10_
            index12_ = default__.safeIndex(616, (d_91_v42_).length(0))
            (d_91_v42_)[index12_] = default__.fm11(d_38_v4_, d_38_v4_, not(d_34_v1_.f15), not(False), globalState)
            index13_ = default__.safeIndex(616, (d_91_v42_).length(0))
            (d_91_v42_)[index13_] = (d_87_v38_) - ((d_87_v38_) | (d_87_v38_))
            (globalState).f4 = d_38_v4_
            (globalState).f1 = (d_80_v33_)[default__.safeIndex(800, (d_80_v33_).length(0))]
            (globalState).f0 = d_38_v4_
        elif True:
            d_92_v43_: _dafny.Seq
            d_92_v43_ = _dafny.SeqWithoutIsStrInference([d_38_v4_, d_38_v4_])
            d_93_v44_: _dafny.Map
            d_93_v44_ = _dafny.Map({d_38_v4_: _dafny.Map({d_39_v5_: d_92_v43_})})
            d_94_v45_: _dafny.Map
            d_94_v45_ = _dafny.Map({len(d_92_v43_): d_38_v4_})
            d_95_v46_: _dafny.Seq
            d_95_v46_ = _dafny.SeqWithoutIsStrInference([d_34_v1_.f15])
            d_93_v44_ = (d_93_v44_).set(d_38_v4_, _dafny.Map({d_39_v5_: _dafny.SeqWithoutIsStrInference([len(d_94_v45_), (0) - ((0) - (len(d_95_v46_))), d_38_v4_, d_38_v4_, d_38_v4_])}))
            d_96_v47_: _dafny.Array
            nw11_ = _dafny.Array(False, 25)
            d_96_v47_ = nw11_
            index14_ = default__.safeIndex(157, (d_96_v47_).length(0))
            (d_96_v47_)[index14_] = False
            index15_ = default__.safeIndex(157, (d_96_v47_).length(0))
            (d_96_v47_)[index15_] = not (True) or (False)
            d_97_v48_: _dafny.Array
            nw12_ = _dafny.Array(int(0), 27)
            d_97_v48_ = nw12_
            index16_ = default__.safeIndex(259, (d_97_v48_).length(0))
            (d_97_v48_)[index16_] = default__.fm0((d_96_v47_)[default__.safeIndex(157, (d_96_v47_).length(0))], 604, (d_96_v47_)[default__.safeIndex(157, (d_96_v47_).length(0))], globalState)
            index17_ = default__.safeIndex(259, (d_97_v48_).length(0))
            (d_97_v48_)[index17_] = default__.safeModuloInt((0) - (d_38_v4_), d_38_v4_)
            d_98_v49_: _dafny.Map
            d_98_v49_ = _dafny.Map({(d_96_v47_)[default__.safeIndex(157, (d_96_v47_).length(0))]: d_34_v1_.f15})
            (globalState).f0 = len((d_95_v46_).set(default__.safeIndex((d_38_v4_ if ((d_98_v49_)[d_33_v0_] if (d_33_v0_) in (d_98_v49_) else d_34_v1_.f15) else (0) - (d_38_v4_)), len(d_95_v46_)), d_34_v1_.f15))
            if d_33_v0_:
                rhs4_ = d_39_v5_
                rhs5_ = d_94_v45_
                d_39_v5_ = rhs4_
                d_94_v45_ = rhs5_
                d_99_v50_: _dafny.Set
                d_99_v50_ = _dafny.Set({d_38_v4_})
                d_100_v51_: _dafny.Map
                d_100_v51_ = _dafny.Map({default__.fm12(len(d_99_v50_), globalState): _dafny.Set({(d_96_v47_)[default__.safeIndex(157, (d_96_v47_).length(0))]})})
                d_101_v52_: _dafny.Set
                d_101_v52_ = _dafny.Set({True})
                d_100_v51_ = (d_100_v51_).set(d_95_v46_, d_101_v52_)
                d_102_v53_: C0
                nw13_ = C0()
                nw13_.ctor__((367) == (643))
                d_102_v53_ = nw13_
                d_95_v46_ = (d_95_v46_) + (d_95_v46_)
                (globalState).f11 = d_96_v47_
            elif True:
                d_103_v54_: _dafny.Map
                d_103_v54_ = _dafny.Map({(d_97_v48_)[default__.safeIndex(259, (d_97_v48_).length(0))]: d_34_v1_})
                d_104_v55_: _dafny.Set
                d_104_v55_ = _dafny.Set({d_34_v1_.f15, not(d_34_v1_.f15)})
                d_105_v56_: _dafny.Map
                d_105_v56_ = _dafny.Map({d_104_v55_: (d_97_v48_)[default__.safeIndex(259, (d_97_v48_).length(0))]})
                d_106_v57_: _dafny.Set
                d_106_v57_ = _dafny.Set({d_34_v1_, ((d_103_v54_)[len(d_105_v56_)] if (len(d_105_v56_)) in (d_103_v54_) else d_34_v1_), d_34_v1_})
                d_107_v58_: _dafny.MultiSet
                d_107_v58_ = _dafny.MultiSet([(d_96_v47_)[default__.safeIndex(157, (d_96_v47_).length(0))], (d_96_v47_)[default__.safeIndex(157, (d_96_v47_).length(0))]])
                rhs6_ = (0) - ((d_97_v48_)[default__.safeIndex(259, (d_97_v48_).length(0))])
                rhs7_ = ((d_103_v54_)[default__.safeDivisionInt(d_38_v4_, (d_107_v58_).cardinality)] if (default__.safeDivisionInt(d_38_v4_, (d_107_v58_).cardinality)) in (d_103_v54_) else d_34_v1_)
                rhs8_ = d_38_v4_
                rhs9_ = ((d_106_v57_) | (d_106_v57_)) | (d_106_v57_)
                lhs4_ = globalState
                lhs5_ = globalState
                lhs4_.f4 = rhs6_
                d_34_v1_ = rhs7_
                lhs5_.f6 = rhs8_
                d_106_v57_ = rhs9_
                d_108_v59_: _dafny.Map
                d_108_v59_ = _dafny.Map({d_96_v47_: (d_97_v48_)[default__.safeIndex(259, (d_97_v48_).length(0))]})
                d_108_v59_ = (d_108_v59_).set(d_96_v47_, 282)
                d_109_v60_: _dafny.Map
                d_109_v60_ = _dafny.Map({(d_38_v4_) <= (d_38_v4_): d_96_v47_})
                d_109_v60_ = (d_109_v60_).set(default__.fm2(globalState), d_96_v47_)
                (globalState).f6 = default__.fm0(d_34_v1_.f15, (d_92_v43_)[default__.safeIndex(-243, len(d_92_v43_))], d_34_v1_.f15, globalState)
                d_110_v61_: _dafny.Seq
                d_110_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkynuw"))
                d_94_v45_ = (d_94_v45_).set((d_97_v48_)[default__.safeIndex(259, (d_97_v48_).length(0))], len((d_110_v61_) + (d_110_v61_)))
        if d_34_v1_.f15:
            d_111_v62_: _dafny.Array
            def lambda6_(d_112_v4_):
                def lambda7_(d_113_i6_):
                    return (d_113_i6_) - (d_112_v4_)

                return lambda7_

            init4_ = lambda6_(d_38_v4_)
            nw14_ = _dafny.Array(None, 10)
            for i0_4_ in range(nw14_.length(0)):
                nw14_[i0_4_] = init4_(i0_4_)
            d_111_v62_ = nw14_
            d_114_v63_: _dafny.MultiSet
            d_114_v63_ = _dafny.MultiSet([d_39_v5_])
            d_115_v64_: _dafny.Set
            d_115_v64_ = _dafny.Set({d_38_v4_})
            d_116_v65_: _dafny.Seq
            d_116_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
            d_117_v66_: D1
            d_117_v66_ = D1_DC2(d_38_v4_)
            d_118_v67_: _dafny.Map
            d_118_v67_ = _dafny.Map({d_33_v0_: d_34_v1_.f15})
            d_119_v68_: _dafny.Array
            nw15_ = _dafny.Array(None, 25)
            nw15_[int(0)] = (d_116_v65_) + (d_116_v65_)
            nw15_[int(1)] = default__.fm6(d_38_v4_, d_34_v1_.f15, d_117_v66_, True, globalState)
            nw15_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kov"))
            nw15_[int(3)] = d_116_v65_
            nw15_[int(4)] = d_116_v65_
            nw15_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndxuxbq"))
            nw15_[int(6)] = d_116_v65_
            nw15_[int(7)] = d_116_v65_
            nw15_[int(8)] = d_116_v65_
            nw15_[int(9)] = d_116_v65_
            nw15_[int(10)] = d_116_v65_
            nw15_[int(11)] = default__.fm6(d_38_v4_, d_34_v1_.f15, D1_DC2(len(d_116_v65_)), d_34_v1_.f15, globalState)
            nw15_[int(12)] = _dafny.SeqWithoutIsStrInference([d_39_v5_ for d_120_i7_ in range(default__.abs(372))])
            nw15_[int(13)] = d_116_v65_
            nw15_[int(14)] = d_116_v65_
            nw15_[int(15)] = d_116_v65_
            nw15_[int(16)] = (d_116_v65_ if not(d_34_v1_.f15) else d_116_v65_)
            nw15_[int(17)] = d_116_v65_
            nw15_[int(18)] = d_116_v65_
            nw15_[int(19)] = (_dafny.SeqWithoutIsStrInference([d_39_v5_ for d_121_i8_ in range(default__.abs(123))])).set(default__.safeIndex(d_38_v4_, len(_dafny.SeqWithoutIsStrInference([d_39_v5_ for d_122_i8_ in range(default__.abs(123))]))), d_39_v5_)
            nw15_[int(20)] = d_116_v65_
            nw15_[int(21)] = d_116_v65_
            nw15_[int(22)] = default__.fm6(d_38_v4_, default__.fm2(globalState), D1_DC2(len(d_118_v67_)), d_34_v1_.f15, globalState)
            nw15_[int(23)] = default__.fm6(d_38_v4_, True, d_117_v66_, d_34_v1_.f15, globalState)
            nw15_[int(24)] = d_116_v65_
            d_119_v68_ = nw15_
            index18_ = default__.safeIndex(985, (d_119_v68_).length(0))
            (d_119_v68_)[index18_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ju"))
            index19_ = default__.safeIndex(132, (d_111_v62_).length(0))
            (d_111_v62_)[index19_] = d_38_v4_
            index20_ = default__.safeIndex(985, (d_119_v68_).length(0))
            index21_ = default__.safeIndex(132, (d_111_v62_).length(0))
            rhs10_ = d_111_v62_
            rhs11_ = d_114_v63_
            rhs12_ = (_dafny.Set({d_38_v4_})) | (d_115_v64_)
            rhs13_ = d_116_v65_
            rhs14_ = (d_38_v4_) + (d_38_v4_)
            lhs6_ = d_119_v68_
            lhs7_ = default__.safeIndex(985, (d_119_v68_).length(0))
            lhs8_ = d_111_v62_
            lhs9_ = default__.safeIndex(132, (d_111_v62_).length(0))
            d_111_v62_ = rhs10_
            d_114_v63_ = rhs11_
            d_115_v64_ = rhs12_
            lhs6_[lhs7_] = rhs13_
            lhs8_[lhs9_] = rhs14_
            d_123_v69_: _dafny.Array
            nw16_ = _dafny.Array(_dafny.Seq({}), 23)
            d_123_v69_ = nw16_
            d_124_v70_: _dafny.Seq
            d_124_v70_ = _dafny.SeqWithoutIsStrInference([(d_111_v62_)[default__.safeIndex(132, (d_111_v62_).length(0))], default__.fm0(d_34_v1_.f15, d_38_v4_, not(d_33_v0_), globalState), d_38_v4_, (d_111_v62_)[default__.safeIndex(132, (d_111_v62_).length(0))], (d_111_v62_)[default__.safeIndex(132, (d_111_v62_).length(0))]])
            index22_ = default__.safeIndex(550, (d_123_v69_).length(0))
            (d_123_v69_)[index22_] = d_124_v70_
            index23_ = default__.safeIndex(550, (d_123_v69_).length(0))
            (d_123_v69_)[index23_] = d_124_v70_
            d_125_v71_: _dafny.Array
            nw17_ = _dafny.Array(False, 6)
            d_125_v71_ = nw17_
            index24_ = default__.safeIndex(460, (d_125_v71_).length(0))
            (d_125_v71_)[index24_] = not(d_34_v1_.f15)
            index25_ = default__.safeIndex(460, (d_125_v71_).length(0))
            (d_125_v71_)[index25_] = d_33_v0_
            d_126_v72_: _dafny.Map
            d_126_v72_ = _dafny.Map({len((d_119_v68_)[default__.safeIndex(985, (d_119_v68_).length(0))]): d_34_v1_.f15})
            d_127_v73_: _dafny.Seq
            d_127_v73_ = _dafny.SeqWithoutIsStrInference([d_111_v62_, d_111_v62_])
            d_128_v74_: D6
            d_128_v74_ = D6_DC15(d_111_v62_)
            d_129_v75_: _dafny.Array
            nw18_ = _dafny.Array(None, 16)
            nw18_[int(0)] = d_111_v62_
            nw18_[int(1)] = d_111_v62_
            nw18_[int(2)] = d_111_v62_
            nw18_[int(3)] = (d_111_v62_ if ((d_126_v72_)[d_38_v4_] if (d_38_v4_) in (d_126_v72_) else d_34_v1_.f15) else (d_127_v73_)[default__.safeIndex((d_111_v62_)[default__.safeIndex(132, (d_111_v62_).length(0))], len(d_127_v73_))])
            nw18_[int(4)] = d_111_v62_
            nw18_[int(5)] = d_111_v62_
            nw18_[int(6)] = d_111_v62_
            nw18_[int(7)] = d_111_v62_
            nw18_[int(8)] = d_111_v62_
            nw18_[int(9)] = (d_111_v62_ if (d_125_v71_)[default__.safeIndex(460, (d_125_v71_).length(0))] else (d_128_v74_).cf21)
            nw18_[int(10)] = d_111_v62_
            nw18_[int(11)] = d_111_v62_
            nw18_[int(12)] = d_111_v62_
            nw18_[int(13)] = d_111_v62_
            nw18_[int(14)] = d_111_v62_
            nw18_[int(15)] = d_111_v62_
            d_129_v75_ = nw18_
            index26_ = default__.safeIndex(181, (d_129_v75_).length(0))
            (d_129_v75_)[index26_] = d_111_v62_
            d_130_v77_: _dafny.Array
            def lambda8_(d_131_v1_, d_132_v62_, d_133_v5_):
                def lambda9_(d_134_i9_):
                    def iife12_():
                        coll6_ = _dafny.Set()
                        compr_6_: _dafny.Seq
                        for compr_6_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_131_v1_.f15, (D4_DC10((d_132_v62_)[default__.safeIndex(132, (d_132_v62_).length(0))], d_133_v5_, (d_132_v62_)[default__.safeIndex(132, (d_132_v62_).length(0))], d_131_v1_.f15, _dafny.SeqWithoutIsStrInference([D3_DC6((d_132_v62_)[default__.safeIndex(132, (d_132_v62_).length(0))])]))).cf16])])).Elements:
                            d_135_v76_: _dafny.Seq = compr_6_
                            if (d_135_v76_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_131_v1_.f15, (D4_DC10((d_132_v62_)[default__.safeIndex(132, (d_132_v62_).length(0))], d_133_v5_, (d_132_v62_)[default__.safeIndex(132, (d_132_v62_).length(0))], d_131_v1_.f15, _dafny.SeqWithoutIsStrInference([D3_DC6((d_132_v62_)[default__.safeIndex(132, (d_132_v62_).length(0))])]))).cf16])])):
                                coll6_ = coll6_.union(_dafny.Set([d_135_v76_]))
                        return _dafny.Set(coll6_)
                    return iife12_()
                    

                return lambda9_

            init5_ = lambda8_(d_34_v1_, d_111_v62_, d_39_v5_)
            nw19_ = _dafny.Array(None, 1)
            for i0_5_ in range(nw19_.length(0)):
                nw19_[i0_5_] = init5_(i0_5_)
            d_130_v77_ = nw19_
            d_136_v78_: _dafny.Seq
            d_136_v78_ = _dafny.SeqWithoutIsStrInference([True, d_34_v1_.f15, (d_125_v71_)[default__.safeIndex(460, (d_125_v71_).length(0))]])
            d_137_v79_: _dafny.Set
            d_137_v79_ = _dafny.Set({d_136_v78_, d_136_v78_, d_136_v78_})
            index27_ = default__.safeIndex(552, (d_130_v77_).length(0))
            (d_130_v77_)[index27_] = d_137_v79_
            index28_ = default__.safeIndex(181, (d_129_v75_).length(0))
            index29_ = default__.safeIndex(552, (d_130_v77_).length(0))
            index30_ = default__.safeIndex(460, (d_125_v71_).length(0))
            rhs15_ = (d_128_v74_).cf21
            rhs16_ = d_34_v1_.f15
            rhs17_ = ((d_111_v62_)[default__.safeIndex(132, (d_111_v62_).length(0))] if d_34_v1_.f15 else (d_111_v62_)[default__.safeIndex(132, (d_111_v62_).length(0))])
            rhs18_ = (default__.fm13((d_111_v62_)[default__.safeIndex(132, (d_111_v62_).length(0))], globalState)) - (d_137_v79_)
            rhs19_ = (d_125_v71_)[default__.safeIndex(460, (d_125_v71_).length(0))]
            lhs10_ = d_129_v75_
            lhs11_ = default__.safeIndex(181, (d_129_v75_).length(0))
            lhs12_ = globalState
            lhs13_ = globalState
            lhs14_ = d_130_v77_
            lhs15_ = default__.safeIndex(552, (d_130_v77_).length(0))
            lhs16_ = d_125_v71_
            lhs17_ = default__.safeIndex(460, (d_125_v71_).length(0))
            lhs10_[lhs11_] = rhs15_
            lhs12_.f10 = rhs16_
            lhs13_.f0 = rhs17_
            lhs14_[lhs15_] = rhs18_
            lhs16_[lhs17_] = rhs19_
            (globalState).f10 = (d_125_v71_)[default__.safeIndex(460, (d_125_v71_).length(0))]
        elif True:
            (d_34_v1_).f15 = d_34_v1_.f15
            if d_34_v1_.f15:
                d_138_v80_: _dafny.Array
                nw20_ = _dafny.Array(None, 1)
                nw20_[int(0)] = d_38_v4_
                d_138_v80_ = nw20_
                d_139_v81_: _dafny.Map
                d_139_v81_ = _dafny.Map({d_38_v4_: d_138_v80_})
                d_140_v82_: _dafny.Array
                nw21_ = _dafny.Array(None, 25)
                nw21_[int(0)] = d_34_v1_.f15
                nw21_[int(1)] = d_34_v1_.f15
                nw21_[int(2)] = d_34_v1_.f15
                nw21_[int(3)] = False
                nw21_[int(4)] = d_34_v1_.f15
                nw21_[int(5)] = d_34_v1_.f15
                nw21_[int(6)] = d_34_v1_.f15
                nw21_[int(7)] = d_34_v1_.f15
                nw21_[int(8)] = d_34_v1_.f15
                nw21_[int(9)] = d_33_v0_
                nw21_[int(10)] = d_34_v1_.f15
                nw21_[int(11)] = d_33_v0_
                nw21_[int(12)] = d_34_v1_.f15
                nw21_[int(13)] = d_34_v1_.f15
                nw21_[int(14)] = True
                nw21_[int(15)] = d_33_v0_
                nw21_[int(16)] = d_34_v1_.f15
                nw21_[int(17)] = d_33_v0_
                nw21_[int(18)] = d_34_v1_.f15
                nw21_[int(19)] = d_33_v0_
                nw21_[int(20)] = d_34_v1_.f15
                nw21_[int(21)] = d_34_v1_.f15
                nw21_[int(22)] = d_33_v0_
                nw21_[int(23)] = d_33_v0_
                nw21_[int(24)] = False
                d_140_v82_ = nw21_
                d_141_v83_: _dafny.Array
                d_141_v83_ = d_140_v82_
                d_142_v84_: _dafny.Array
                nw22_ = _dafny.Array(None, 18)
                nw22_[int(0)] = d_138_v80_
                nw22_[int(1)] = d_138_v80_
                nw22_[int(2)] = ((d_139_v81_)[len(_dafny.Map({d_141_v83_: d_34_v1_.f15}))] if (len(_dafny.Map({d_141_v83_: d_34_v1_.f15}))) in (d_139_v81_) else d_138_v80_)
                nw22_[int(3)] = d_138_v80_
                nw22_[int(4)] = d_138_v80_
                nw22_[int(5)] = d_138_v80_
                nw22_[int(6)] = d_138_v80_
                nw22_[int(7)] = d_138_v80_
                nw22_[int(8)] = d_138_v80_
                nw22_[int(9)] = d_138_v80_
                nw22_[int(10)] = d_138_v80_
                nw22_[int(11)] = d_138_v80_
                nw22_[int(12)] = d_138_v80_
                nw22_[int(13)] = d_138_v80_
                nw22_[int(14)] = d_138_v80_
                nw22_[int(15)] = d_138_v80_
                nw22_[int(16)] = d_138_v80_
                nw22_[int(17)] = d_138_v80_
                d_142_v84_ = nw22_
                d_143_v85_: _dafny.Seq
                d_143_v85_ = _dafny.SeqWithoutIsStrInference([d_142_v84_])
                d_144_v86_: _dafny.Seq
                d_144_v86_ = _dafny.SeqWithoutIsStrInference([d_34_v1_.f15, d_34_v1_.f15])
                d_145_v87_: _dafny.Map
                d_145_v87_ = _dafny.Map({(d_143_v85_)[default__.safeIndex(d_38_v4_, len(d_143_v85_))]: (d_144_v86_) + (d_144_v86_)})
                d_145_v87_ = (d_145_v87_).set(d_142_v84_, d_144_v86_)
                d_33_v0_ = (d_38_v4_) == (d_38_v4_)
                d_146_v88_: _dafny.Map
                d_146_v88_ = _dafny.Map({d_34_v1_.f15: d_34_v1_.f15})
                d_146_v88_ = (d_146_v88_).set(d_34_v1_.f15, d_34_v1_.f15)
                d_147_v89_: D3
                d_147_v89_ = D3_DC5(d_144_v86_)
                (globalState).f0 = len(_dafny.Map({d_144_v86_: D3_DC7(d_147_v89_)}))
                index31_ = default__.safeIndex(630, (d_140_v82_).length(0))
                (d_140_v82_)[index31_] = d_34_v1_.f15
                d_148_v90_: D1
                d_148_v90_ = D1_DC2(d_38_v4_)
                d_149_v91_: _dafny.Map
                d_149_v91_ = _dafny.Map({249: d_38_v4_})
                d_150_v92_: _dafny.Array
                nw23_ = _dafny.Array(None, 2)
                nw23_[int(0)] = d_148_v90_
                nw23_[int(1)] = D1_DC2(len(d_149_v91_))
                d_150_v92_ = nw23_
                d_151_v93_: _dafny.MultiSet
                d_151_v93_ = _dafny.MultiSet([d_150_v92_])
                index32_ = default__.safeIndex(630, (d_140_v82_).length(0))
                rhs20_ = d_38_v4_
                rhs21_ = d_34_v1_.f15
                rhs22_ = ((_dafny.MultiSet([d_150_v92_]) if True else d_151_v93_)).isdisjoint(d_151_v93_)
                rhs23_ = d_138_v80_
                lhs18_ = globalState
                lhs19_ = d_140_v82_
                lhs20_ = default__.safeIndex(630, (d_140_v82_).length(0))
                lhs21_ = d_34_v1_
                lhs18_.f4 = rhs20_
                lhs19_[lhs20_] = rhs21_
                lhs21_.f15 = rhs22_
                d_138_v80_ = rhs23_
            elif True:
                d_152_v94_: _dafny.Map
                d_152_v94_ = _dafny.Map({d_34_v1_.f15: d_34_v1_})
                d_153_v95_: _dafny.Array
                nw24_ = _dafny.Array(None, 14)
                nw24_[int(0)] = ((d_152_v94_)[not(True)] if (not(True)) in (d_152_v94_) else d_34_v1_)
                nw24_[int(1)] = d_34_v1_
                nw24_[int(2)] = d_34_v1_
                nw24_[int(3)] = d_34_v1_
                nw24_[int(4)] = d_34_v1_
                nw24_[int(5)] = d_34_v1_
                nw24_[int(6)] = d_34_v1_
                nw24_[int(7)] = d_34_v1_
                nw24_[int(8)] = d_34_v1_
                nw24_[int(9)] = ((d_152_v94_)[d_34_v1_.f15] if (d_34_v1_.f15) in (d_152_v94_) else d_34_v1_)
                nw24_[int(10)] = d_34_v1_
                nw24_[int(11)] = d_34_v1_
                nw24_[int(12)] = d_34_v1_
                nw24_[int(13)] = d_34_v1_
                d_153_v95_ = nw24_
                d_154_v96_: _dafny.Seq
                d_154_v96_ = _dafny.SeqWithoutIsStrInference([d_34_v1_, d_34_v1_])
                index33_ = default__.safeIndex(101, (d_153_v95_).length(0))
                (d_153_v95_)[index33_] = (d_154_v96_)[default__.safeIndex(d_38_v4_, len(d_154_v96_))]
                d_155_v97_: _dafny.Seq
                d_155_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iaufjsuky"))
                d_156_v98_: _dafny.Seq
                d_156_v98_ = _dafny.SeqWithoutIsStrInference([d_155_v97_, d_155_v97_])
                d_157_v99_: D2
                d_157_v99_ = D2_DC3(d_155_v97_)
                d_158_v100_: _dafny.Array
                nw25_ = _dafny.Array(None, 5)
                nw25_[int(0)] = (d_156_v98_)[default__.safeIndex(len(d_155_v97_), len(d_156_v98_))]
                nw25_[int(1)] = d_155_v97_
                nw25_[int(2)] = d_155_v97_
                nw25_[int(3)] = ((_dafny.SeqWithoutIsStrInference([d_39_v5_ for d_159_i10_ in range(default__.abs(214))])) + (d_155_v97_)).set(default__.safeIndex(d_38_v4_, len((_dafny.SeqWithoutIsStrInference([d_39_v5_ for d_160_i10_ in range(default__.abs(214))])) + (d_155_v97_))), d_39_v5_)
                nw25_[int(4)] = ((d_157_v99_).cf3) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtb")))
                d_158_v100_ = nw25_
                index34_ = default__.safeIndex(222, (d_158_v100_).length(0))
                (d_158_v100_)[index34_] = d_155_v97_
                index35_ = default__.safeIndex(101, (d_153_v95_).length(0))
                index36_ = default__.safeIndex(222, (d_158_v100_).length(0))
                rhs24_ = d_34_v1_.f15
                rhs25_ = (default__.fm0(d_34_v1_.f15, d_38_v4_, d_34_v1_.f15, globalState)) - ((0) - (d_38_v4_))
                rhs26_ = d_34_v1_
                rhs27_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjfq"))) + (d_155_v97_)
                lhs22_ = d_34_v1_
                lhs23_ = globalState
                lhs24_ = d_153_v95_
                lhs25_ = default__.safeIndex(101, (d_153_v95_).length(0))
                lhs26_ = d_158_v100_
                lhs27_ = default__.safeIndex(222, (d_158_v100_).length(0))
                lhs22_.f15 = rhs24_
                lhs23_.f4 = rhs25_
                lhs24_[lhs25_] = rhs26_
                lhs26_[lhs27_] = rhs27_
                d_161_v101_: _dafny.Seq
                d_161_v101_ = _dafny.SeqWithoutIsStrInference([d_34_v1_.f15, d_34_v1_.f15, default__.fm2(globalState)])
                d_161_v101_ = d_161_v101_
                rhs28_ = ((d_34_v1_.f15) or (d_33_v0_)) or (d_34_v1_.f15)
                rhs29_ = (default__.safeDivisionInt(len((d_158_v100_)[default__.safeIndex(222, (d_158_v100_).length(0))]), d_38_v4_) if d_34_v1_.f15 else d_38_v4_)
                lhs28_ = globalState
                lhs28_.f1 = rhs28_
                d_38_v4_ = rhs29_
                d_162_v102_: _dafny.Seq
                d_162_v102_ = _dafny.SeqWithoutIsStrInference([d_38_v4_])
                d_163_v103_: _dafny.Set
                d_163_v103_ = _dafny.Set({not (False) or (False), (_dafny.SeqWithoutIsStrInference([d_38_v4_ for d_164_i11_ in range(default__.abs(-681))])) < (d_162_v102_)})
                d_165_v104_: _dafny.Set
                d_165_v104_ = _dafny.Set({d_38_v4_, 741, default__.safeDivisionInt(d_38_v4_, len((d_158_v100_)[default__.safeIndex(222, (d_158_v100_).length(0))]))})
                rhs30_ = (0) - (len(d_163_v103_))
                rhs31_ = (0) - (d_38_v4_)
                rhs32_ = len(d_165_v104_)
                lhs29_ = globalState
                lhs30_ = globalState
                lhs31_ = globalState
                lhs29_.f6 = rhs30_
                lhs30_.f0 = rhs31_
                lhs31_.f0 = rhs32_
                d_166_v105_: _dafny.Array
                def lambda10_(d_167_v4_):
                    def lambda11_(d_168_i12_):
                        return _dafny.Map({d_167_v4_: d_167_v4_})

                    return lambda11_

                init6_ = lambda10_(d_38_v4_)
                nw26_ = _dafny.Array(None, 15)
                for i0_6_ in range(nw26_.length(0)):
                    nw26_[i0_6_] = init6_(i0_6_)
                d_166_v105_ = nw26_
                d_169_v106_: _dafny.Map
                d_169_v106_ = _dafny.Map({d_34_v1_.f15: d_34_v1_.f15})
                d_170_v107_: D1
                d_170_v107_ = D1_DC2(d_38_v4_)
                rhs33_ = d_166_v105_
                rhs34_ = (default__.fm0(d_34_v1_.f15, (0) - ((0) - (d_38_v4_)), d_34_v1_.f15, globalState)) * ((0) - (len(default__.fm6(len(d_169_v106_), d_34_v1_.f15, d_170_v107_, d_34_v1_.f15, globalState))))
                rhs35_ = d_38_v4_
                rhs36_ = d_38_v4_
                lhs32_ = globalState
                lhs33_ = globalState
                lhs34_ = globalState
                d_166_v105_ = rhs33_
                lhs32_.f0 = rhs34_
                lhs33_.f6 = rhs35_
                lhs34_.f4 = rhs36_
            d_171_v108_: _dafny.Array
            nw27_ = _dafny.Array(False, 25)
            d_171_v108_ = nw27_
            d_172_v109_: _dafny.Array
            d_172_v109_ = d_171_v108_
            source2_ = d_172_v109_
            d_173___mcc_h3_ = source2_
            d_174_cf0_ = d_173___mcc_h3_
            d_175_v110_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
            d_175_v110_ = nw28_
            d_176_v111_: _dafny.Seq
            d_176_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mftvud"))
            index37_ = default__.safeIndex(547, (d_175_v110_).length(0))
            (d_175_v110_)[index37_] = (d_176_v111_ if d_34_v1_.f15 else d_176_v111_)
            index38_ = default__.safeIndex(144, (d_174_cf0_).length(0))
            (d_174_cf0_)[index38_] = d_34_v1_.f15
            index39_ = default__.safeIndex(547, (d_175_v110_).length(0))
            index40_ = default__.safeIndex(144, (d_174_cf0_).length(0))
            rhs37_ = d_176_v111_
            rhs38_ = d_38_v4_
            rhs39_ = default__.safeDivisionInt(d_38_v4_, default__.safeDivisionInt(d_38_v4_, d_38_v4_))
            rhs40_ = not(not(d_34_v1_.f15))
            lhs35_ = d_175_v110_
            lhs36_ = default__.safeIndex(547, (d_175_v110_).length(0))
            lhs37_ = globalState
            lhs38_ = d_174_cf0_
            lhs39_ = default__.safeIndex(144, (d_174_cf0_).length(0))
            lhs35_[lhs36_] = rhs37_
            d_38_v4_ = rhs38_
            lhs37_.f0 = rhs39_
            lhs38_[lhs39_] = rhs40_
            d_177_v112_: _dafny.Map
            d_177_v112_ = _dafny.Map({(d_175_v110_ if not(d_33_v0_) else d_175_v110_): (0) - (d_38_v4_)})
            d_177_v112_ = (d_177_v112_).set(d_175_v110_, default__.fm0(False, d_38_v4_, d_33_v0_, globalState))
            d_178_v113_: C0
            nw29_ = C0()
            nw29_.ctor__(not(not (d_34_v1_.f15) or (d_34_v1_.f15)))
            d_178_v113_ = nw29_
            d_179_v114_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_179_v114_ = nw30_
            d_180_v115_: _dafny.Array
            nw31_ = _dafny.Array(int(0), 21)
            d_180_v115_ = nw31_
            index41_ = default__.safeIndex(611, (d_179_v114_).length(0))
            (d_179_v114_)[index41_] = d_180_v115_
            index42_ = default__.safeIndex(611, (d_179_v114_).length(0))
            (d_179_v114_)[index42_] = d_180_v115_
            if d_34_v1_.f15:
                d_33_v0_ = not(d_34_v1_.f15)
                (globalState).f4 = d_38_v4_
                d_181_v116_: _dafny.Map
                d_181_v116_ = _dafny.Map({d_34_v1_.f15: d_38_v4_})
                d_182_v117_: _dafny.Seq
                d_182_v117_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shhs"))
                d_183_v118_: D1
                d_183_v118_ = D1_DC1(d_38_v4_)
                d_184_v119_: _dafny.Set
                d_184_v119_ = _dafny.Set({d_34_v1_.f15, d_34_v1_.f15})
                d_185_v120_: _dafny.Array
                nw32_ = _dafny.Array(None, 23)
                nw32_[int(0)] = 115
                nw32_[int(1)] = len(d_181_v116_)
                nw32_[int(2)] = len(d_182_v117_)
                nw32_[int(3)] = d_38_v4_
                nw32_[int(4)] = len(d_182_v117_)
                nw32_[int(5)] = (d_183_v118_).cf1
                nw32_[int(6)] = len(d_182_v117_)
                nw32_[int(7)] = (0) - (d_38_v4_)
                nw32_[int(8)] = d_38_v4_
                nw32_[int(9)] = d_38_v4_
                nw32_[int(10)] = d_38_v4_
                nw32_[int(11)] = d_38_v4_
                nw32_[int(12)] = len(default__.fm4(d_34_v1_.f15, d_38_v4_, globalState))
                nw32_[int(13)] = len(d_184_v119_)
                nw32_[int(14)] = d_38_v4_
                nw32_[int(15)] = d_38_v4_
                nw32_[int(16)] = default__.fm0(d_34_v1_.f15, d_38_v4_, d_34_v1_.f15, globalState)
                nw32_[int(17)] = d_38_v4_
                nw32_[int(18)] = d_38_v4_
                nw32_[int(19)] = (0) - (d_38_v4_)
                nw32_[int(20)] = d_38_v4_
                nw32_[int(21)] = d_38_v4_
                nw32_[int(22)] = d_38_v4_
                d_185_v120_ = nw32_
                d_186_v121_: _dafny.Array
                nw33_ = _dafny.Array(None, 17)
                nw33_[int(0)] = d_185_v120_
                nw33_[int(1)] = d_185_v120_
                nw33_[int(2)] = d_185_v120_
                nw33_[int(3)] = d_185_v120_
                nw33_[int(4)] = d_185_v120_
                nw33_[int(5)] = d_185_v120_
                nw33_[int(6)] = d_185_v120_
                nw33_[int(7)] = d_185_v120_
                nw33_[int(8)] = d_185_v120_
                nw33_[int(9)] = d_185_v120_
                nw33_[int(10)] = d_185_v120_
                nw33_[int(11)] = d_185_v120_
                nw33_[int(12)] = d_185_v120_
                nw33_[int(13)] = d_185_v120_
                nw33_[int(14)] = d_185_v120_
                nw33_[int(15)] = d_185_v120_
                nw33_[int(16)] = d_185_v120_
                d_186_v121_ = nw33_
                d_187_v122_: _dafny.Seq
                d_187_v122_ = _dafny.SeqWithoutIsStrInference([d_186_v121_, d_186_v121_, d_186_v121_])
                d_188_v123_: D7
                d_188_v123_ = D7_DC17(d_186_v121_)
                d_189_v124_: _dafny.Array
                nw34_ = _dafny.Array(None, 16)
                nw34_[int(0)] = d_186_v121_
                nw34_[int(1)] = d_186_v121_
                nw34_[int(2)] = (d_187_v122_)[default__.safeIndex(d_38_v4_, len(d_187_v122_))]
                nw34_[int(3)] = d_186_v121_
                nw34_[int(4)] = d_186_v121_
                nw34_[int(5)] = d_186_v121_
                nw34_[int(6)] = d_186_v121_
                nw34_[int(7)] = d_186_v121_
                nw34_[int(8)] = d_186_v121_
                nw34_[int(9)] = d_186_v121_
                nw34_[int(10)] = d_186_v121_
                nw34_[int(11)] = d_186_v121_
                nw34_[int(12)] = d_186_v121_
                nw34_[int(13)] = (d_188_v123_).cf25
                nw34_[int(14)] = d_186_v121_
                nw34_[int(15)] = d_186_v121_
                d_189_v124_ = nw34_
                index43_ = default__.safeIndex(50, (d_189_v124_).length(0))
                (d_189_v124_)[index43_] = d_186_v121_
                d_190_v125_: _dafny.Array
                nw35_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_190_v125_ = nw35_
                index44_ = default__.safeIndex(50, (d_189_v124_).length(0))
                rhs41_ = d_39_v5_
                rhs42_ = d_186_v121_
                rhs43_ = d_190_v125_
                rhs44_ = d_38_v4_
                rhs45_ = (D2_DC3(d_182_v117_)).cf3
                lhs40_ = d_189_v124_
                lhs41_ = default__.safeIndex(50, (d_189_v124_).length(0))
                lhs42_ = globalState
                d_39_v5_ = rhs41_
                lhs40_[lhs41_] = rhs42_
                d_190_v125_ = rhs43_
                lhs42_.f6 = rhs44_
                d_182_v117_ = rhs45_
                d_191_v126_: _dafny.Seq
                d_191_v126_ = _dafny.SeqWithoutIsStrInference([d_33_v0_])
                d_192_v127_: D3
                d_192_v127_ = D3_DC6(d_38_v4_)
                d_193_v128_: _dafny.Seq
                d_193_v128_ = _dafny.SeqWithoutIsStrInference([d_192_v127_])
                d_194_v129_: D4
                d_194_v129_ = D4_DC10(d_38_v4_, d_39_v5_, d_38_v4_, d_34_v1_.f15, d_193_v128_)
                d_195_v130_: D3
                d_195_v130_ = D3_DC7(D3_DC6(d_38_v4_))
                d_196_v131_: D3
                d_196_v131_ = D3_DC7(d_195_v130_)
                d_197_v132_: D3
                d_197_v132_ = D3_DC7(d_195_v130_)
                pat_let_tv4_ = d_195_v130_
                d_198_v133_: _dafny.Set
                def iife13_(_pat_let3_0):
                    def iife14_(d_199_dt__update__tmp_h1_):
                        def iife15_(_pat_let4_0):
                            def iife16_(d_200_dt__update_hcf8_h0_):
                                return D3_DC7(d_200_dt__update_hcf8_h0_)
                            return iife16_(_pat_let4_0)
                        return iife15_(pat_let_tv4_)
                    return iife14_(_pat_let3_0)
                d_198_v133_ = _dafny.Set({d_197_v132_, iife13_(d_197_v132_), d_197_v132_, d_197_v132_})
                rhs46_ = (d_194_v129_).cf16
                rhs47_ = default__.safeModuloInt(45, (d_38_v4_) - (len(d_182_v117_)))
                rhs48_ = (d_182_v117_).set(default__.safeIndex(len((d_198_v133_) | (d_198_v133_)), len(d_182_v117_)), d_39_v5_)
                rhs49_ = d_34_v1_
                rhs50_ = d_191_v126_
                lhs43_ = globalState
                lhs44_ = globalState
                lhs43_.f8 = rhs46_
                lhs44_.f0 = rhs47_
                d_182_v117_ = rhs48_
                d_34_v1_ = rhs49_
                d_191_v126_ = rhs50_
                d_201_v134_: D4
                d_201_v134_ = D4_DC8(d_33_v0_)
                (globalState).f10 = (d_201_v134_).cf9
            elif True:
                d_202_v135_: C0
                nw36_ = C0()
                nw36_.ctor__(False)
                d_202_v135_ = nw36_
                d_203_v136_: _dafny.Map
                d_203_v136_ = _dafny.Map({default__.fm2(globalState): d_34_v1_.f15})
                d_204_v137_: _dafny.Set
                d_204_v137_ = _dafny.Set({d_39_v5_})
                d_203_v136_ = (d_203_v136_).set(d_202_v135_.f15, (d_204_v137_).ispropersubset(d_204_v137_))
                d_205_v138_: _dafny.Map
                d_205_v138_ = _dafny.Map({d_38_v4_: d_202_v135_})
                d_206_v139_: _dafny.Seq
                d_206_v139_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "medr"))
                d_207_v140_: _dafny.Map
                d_207_v140_ = _dafny.Map({(d_38_v4_ if not(d_34_v1_.f15) else d_38_v4_): ((d_205_v138_)[len(d_206_v139_)] if (len(d_206_v139_)) in (d_205_v138_) else d_34_v1_)})
                d_205_v138_ = (d_205_v138_).set(d_38_v4_, d_34_v1_)
                d_208_v141_: _dafny.Seq
                d_208_v141_ = _dafny.SeqWithoutIsStrInference([d_202_v135_.f15])
                index45_ = default__.safeIndex(536, (d_171_v108_).length(0))
                (d_171_v108_)[index45_] = (d_208_v141_) <= (d_208_v141_)
                index46_ = default__.safeIndex(536, (d_171_v108_).length(0))
                rhs51_ = d_206_v139_
                rhs52_ = 959
                rhs53_ = d_34_v1_.f15
                lhs45_ = globalState
                lhs46_ = d_171_v108_
                lhs47_ = default__.safeIndex(536, (d_171_v108_).length(0))
                d_206_v139_ = rhs51_
                lhs45_.f0 = rhs52_
                lhs46_[lhs47_] = rhs53_
                d_209_v142_: _dafny.Set
                d_209_v142_ = _dafny.Set({d_33_v0_})
                d_210_v143_: _dafny.Array
                nw37_ = _dafny.Array(None, 8)
                nw37_[int(0)] = d_38_v4_
                nw37_[int(1)] = default__.fm0(True, d_38_v4_, (d_208_v141_)[default__.safeIndex(default__.fm0(d_202_v135_.f15, d_38_v4_, (d_208_v141_)[default__.safeIndex(len(d_209_v142_), len(d_208_v141_))], globalState), len(d_208_v141_))], globalState)
                nw37_[int(2)] = (d_38_v4_) + (len(d_206_v139_))
                nw37_[int(3)] = d_38_v4_
                nw37_[int(4)] = default__.safeModuloInt(d_38_v4_, d_38_v4_)
                nw37_[int(5)] = d_38_v4_
                nw37_[int(6)] = (d_38_v4_ if d_34_v1_.f15 else d_38_v4_)
                nw37_[int(7)] = d_38_v4_
                d_210_v143_ = nw37_
                index47_ = default__.safeIndex(789, (d_210_v143_).length(0))
                (d_210_v143_)[index47_] = default__.safeDivisionInt(d_38_v4_, d_38_v4_)
                d_211_v144_: _dafny.Map
                d_211_v144_ = _dafny.Map({d_38_v4_: d_38_v4_})
                d_212_v146_: D4
                d_212_v146_ = D4_DC8(True)
                index48_ = default__.safeIndex(789, (d_210_v143_).length(0))
                def iife17_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(-450, 868):
                        d_213_v145_: int = compr_7_
                        if ((-450) <= (d_213_v145_)) and ((d_213_v145_) < (868)):
                            coll7_[default__.safeModuloInt(d_213_v145_, d_38_v4_)] = (_dafny.MultiSet([d_38_v4_, d_38_v4_])).cardinality
                    return _dafny.Map(coll7_)
                rhs54_ = (d_211_v144_) != (iife17_()
                )
                rhs55_ = default__.safeDivisionInt(d_38_v4_, default__.safeModuloInt(d_38_v4_, d_38_v4_))
                rhs56_ = (d_34_v1_.f15) or ((d_212_v146_).cf9)
                rhs57_ = default__.safeModuloInt(d_38_v4_, len(default__.fm14(globalState)))
                lhs48_ = globalState
                lhs49_ = d_210_v143_
                lhs50_ = default__.safeIndex(789, (d_210_v143_).length(0))
                lhs51_ = globalState
                lhs52_ = globalState
                lhs48_.f1 = rhs54_
                lhs49_[lhs50_] = rhs55_
                lhs51_.f8 = rhs56_
                lhs52_.f6 = rhs57_
            d_214_v147_: _dafny.Array
            nw38_ = _dafny.Array(None, 2)
            d_214_v147_ = nw38_
            index49_ = default__.safeIndex(141, (d_214_v147_).length(0))
            (d_214_v147_)[index49_] = d_34_v1_
            index50_ = default__.safeIndex(141, (d_214_v147_).length(0))
            (d_214_v147_)[index50_] = d_34_v1_
        d_215_v148_: _dafny.Seq
        d_215_v148_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usstt"))
        d_38_v4_ = (d_38_v4_) + (len((d_215_v148_) + (d_215_v148_)))

    @staticmethod
    def Main(noArgsParameter__):
        d_216_v0_: _dafny.Array
        def lambda12_(d_217_i0_):
            return not(not(not(True)))

        init7_ = lambda12_
        nw39_ = _dafny.Array(None, 18)
        for i0_7_ in range(nw39_.length(0)):
            nw39_[i0_7_] = init7_(i0_7_)
        d_216_v0_ = nw39_
        d_218_v1_: _dafny.Array
        d_218_v1_ = d_216_v0_
        d_219_v2_: bool
        d_219_v2_ = True
        d_220_v3_: _dafny.Map
        d_220_v3_ = _dafny.Map({d_219_v2_: True})
        d_221_globalState_: GlobalState
        nw40_ = GlobalState()
        nw40_.ctor__(527, False, True, 871, 272, 844, 508, True, True, d_216_v0_, False, (d_218_v1_), (d_220_v3_) | (d_220_v3_), 357, False)
        d_221_globalState_ = nw40_
        d_222_v4_: int
        d_222_v4_ = -669
        (d_221_globalState_).f6 = default__.safeDivisionInt(d_222_v4_, -539)
        d_223_v5_: _dafny.Map
        d_223_v5_ = _dafny.Map({d_219_v2_: default__.fm0(True, d_222_v4_, d_219_v2_, d_221_globalState_)})
        if (d_219_v2_) == ((d_219_v2_) in (d_223_v5_)):
            d_219_v2_ = d_219_v2_
            if d_219_v2_:
                d_224_v6_: _dafny.Seq
                d_224_v6_ = _dafny.SeqWithoutIsStrInference([d_219_v2_])
                d_219_v2_ = (d_219_v2_) in (d_224_v6_)
                d_225_v7_: _dafny.Seq
                d_225_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_216_v0_: d_222_v4_})])
                d_226_v8_: _dafny.Map
                d_226_v8_ = _dafny.Map({(d_225_v7_)[default__.safeIndex(d_222_v4_, len(d_225_v7_))]: False})
                d_227_v9_: str
                d_227_v9_ = _dafny.CodePoint('h')
                rhs58_ = d_219_v2_
                rhs59_ = d_226_v8_
                rhs60_ = d_227_v9_
                rhs61_ = d_222_v4_
                lhs53_ = d_221_globalState_
                lhs54_ = d_221_globalState_
                lhs53_.f1 = rhs58_
                d_226_v8_ = rhs59_
                d_227_v9_ = rhs60_
                lhs54_.f6 = rhs61_
                d_228_v10_: C0
                nw41_ = C0()
                nw41_.ctor__(default__.fm2(d_221_globalState_))
                d_228_v10_ = nw41_
                d_229_v11_: _dafny.Seq
                d_229_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
                d_230_v12_: _dafny.Set
                d_230_v12_ = _dafny.Set({893, (len(d_229_v11_)) + (d_222_v4_), d_222_v4_, d_222_v4_, d_222_v4_})
                d_231_v13_: _dafny.Map
                d_231_v13_ = _dafny.Map({d_228_v10_: len(_dafny.SeqWithoutIsStrInference([d_227_v9_ for d_232_i1_ in range(default__.abs(588))]))})
                d_230_v12_ = _dafny.Set({len(d_231_v13_), d_222_v4_})
                (d_221_globalState_).f0 = d_222_v4_
            elif True:
                d_233_v14_: D1
                d_233_v14_ = D1_DC1(-329)
                (d_221_globalState_).f6 = (d_233_v14_).cf1
                d_234_v15_: _dafny.Seq
                d_234_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwxho"))
                d_235_v16_: _dafny.Map
                d_235_v16_ = _dafny.Map({d_222_v4_: d_219_v2_})
                d_236_v17_: D2
                d_236_v17_ = D2_DC3((d_234_v15_).set(default__.safeIndex(d_222_v4_, len(d_234_v15_)), _dafny.CodePoint('b')))
                d_237_v18_: _dafny.Seq
                d_237_v18_ = _dafny.SeqWithoutIsStrInference([d_219_v2_])
                rhs62_ = (d_236_v17_).cf3
                rhs63_ = len((d_237_v18_) + (d_237_v18_))
                rhs64_ = (not(d_219_v2_)) == (not(d_219_v2_))
                rhs65_ = d_235_v16_
                lhs55_ = d_221_globalState_
                lhs56_ = d_221_globalState_
                d_234_v15_ = rhs62_
                lhs55_.f6 = rhs63_
                lhs56_.f8 = rhs64_
                d_235_v16_ = rhs65_
                d_238_v19_: _dafny.Seq
                d_238_v19_ = _dafny.SeqWithoutIsStrInference([d_222_v4_])
                d_239_v20_: str
                d_239_v20_ = _dafny.CodePoint('h')
                d_240_v21_: _dafny.Map
                d_240_v21_ = _dafny.Map({d_239_v20_: d_222_v4_})
                d_241_v22_: _dafny.Array
                nw42_ = _dafny.Array(int(0), 27)
                d_241_v22_ = nw42_
                d_242_v23_: _dafny.Map
                d_242_v23_ = _dafny.Map({d_241_v22_: d_222_v4_})
                d_238_v19_ = (d_238_v19_) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_243_i2_ in range(default__.abs(943))])), len(_dafny.SeqWithoutIsStrInference([len(d_240_v21_), len(d_220_v3_)])), len(d_242_v23_)])) + (d_238_v19_))
                d_244_v24_: C0
                nw43_ = C0()
                nw43_.ctor__(not (d_219_v2_) or (d_219_v2_))
                d_244_v24_ = nw43_
                (d_221_globalState_).f8 = d_244_v24_.f15
            if d_219_v2_:
                d_245_v25_: _dafny.Seq
                d_245_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xygbbejn"))
                d_246_v26_: _dafny.Map
                d_246_v26_ = _dafny.Map({d_245_v25_: d_219_v2_})
                d_247_v27_: C0
                nw44_ = C0()
                nw44_.ctor__(((d_246_v26_)[d_245_v25_] if (d_245_v25_) in (d_246_v26_) else d_219_v2_))
                d_247_v27_ = nw44_
                d_248_v28_: _dafny.Map
                d_248_v28_ = _dafny.Map({d_247_v27_: d_247_v27_.f15})
                d_248_v28_ = (d_248_v28_).set(d_247_v27_, default__.fm2(d_221_globalState_))
                default__.m0(d_221_globalState_)
                d_249_v29_: _dafny.Set
                d_249_v29_ = _dafny.Set({not (d_219_v2_) or (not(d_219_v2_))})
                d_249_v29_ = (default__.fm3(d_222_v4_, d_221_globalState_)) - ((d_249_v29_) | (_dafny.Set({d_247_v27_.f15})))
                (d_221_globalState_).f8 = d_247_v27_.f15
                d_250_v30_: _dafny.Array
                def lambda13_(d_251_v4_):
                    def lambda14_(d_252_i3_):
                        return (d_252_i3_) + (d_251_v4_)

                    return lambda14_

                init8_ = lambda13_(d_222_v4_)
                nw45_ = _dafny.Array(None, 14)
                for i0_8_ in range(nw45_.length(0)):
                    nw45_[i0_8_] = init8_(i0_8_)
                d_250_v30_ = nw45_
                d_253_v31_: _dafny.MultiSet
                d_253_v31_ = _dafny.MultiSet([d_250_v30_])
                d_254_v32_: _dafny.Seq
                d_254_v32_ = _dafny.SeqWithoutIsStrInference([d_253_v31_, d_253_v31_])
                d_255_v33_: _dafny.Seq
                d_255_v33_ = _dafny.SeqWithoutIsStrInference([d_222_v4_, d_222_v4_, d_222_v4_, ((d_254_v32_)[default__.safeIndex(d_222_v4_, len(d_254_v32_))]).cardinality, d_222_v4_])
                d_255_v33_ = d_255_v33_
            elif True:
                d_256_v34_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_256_v34_ = nw46_
                index51_ = default__.safeIndex(202, (d_256_v34_).length(0))
                (d_256_v34_)[index51_] = d_216_v0_
                d_257_v35_: _dafny.Set
                d_257_v35_ = _dafny.Set({True})
                d_258_v36_: _dafny.Set
                d_258_v36_ = _dafny.Set({default__.fm0(d_219_v2_, d_222_v4_, d_219_v2_, d_221_globalState_), (d_222_v4_) * (len(d_257_v35_)), d_222_v4_, (0) - (default__.safeModuloInt(len(d_223_v5_), d_222_v4_))})
                d_259_v37_: _dafny.Seq
                d_259_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xd"))
                index52_ = default__.safeIndex(202, (d_256_v34_).length(0))
                rhs66_ = d_216_v0_
                rhs67_ = d_222_v4_
                rhs68_ = (default__.fm4(d_219_v2_, d_222_v4_, d_221_globalState_)) | ((d_258_v36_).intersection(d_258_v36_))
                rhs69_ = d_219_v2_
                rhs70_ = (d_219_v2_ if (d_259_v37_) < (d_259_v37_) else d_219_v2_)
                lhs57_ = d_256_v34_
                lhs58_ = default__.safeIndex(202, (d_256_v34_).length(0))
                lhs59_ = d_221_globalState_
                lhs60_ = d_221_globalState_
                lhs61_ = d_221_globalState_
                lhs57_[lhs58_] = rhs66_
                lhs59_.f0 = rhs67_
                d_258_v36_ = rhs68_
                lhs60_.f10 = rhs69_
                lhs61_.f8 = rhs70_
                d_260_v38_: C0
                nw47_ = C0()
                nw47_.ctor__(d_219_v2_)
                d_260_v38_ = nw47_
                d_261_v39_: _dafny.Array
                nw48_ = _dafny.Array(int(0), 23)
                d_261_v39_ = nw48_
                d_262_v40_: _dafny.Map
                d_262_v40_ = _dafny.Map({d_222_v4_: d_222_v4_})
                index53_ = default__.safeIndex(726, (d_261_v39_).length(0))
                (d_261_v39_)[index53_] = len(d_262_v40_)
                d_263_v41_: D2
                d_263_v41_ = D2_DC4(613, d_222_v4_)
                d_264_v42_: _dafny.Seq
                d_264_v42_ = _dafny.SeqWithoutIsStrInference([d_219_v2_, d_219_v2_, d_260_v38_.f15])
                d_265_v43_: _dafny.Map
                d_265_v43_ = _dafny.Map({d_263_v41_: d_264_v42_})
                d_266_v44_: _dafny.MultiSet
                d_266_v44_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_260_v38_.f15, d_260_v38_.f15, d_260_v38_.f15]), d_264_v42_])
                index54_ = default__.safeIndex(726, (d_261_v39_).length(0))
                rhs71_ = d_260_v38_
                rhs72_ = d_261_v39_
                rhs73_ = ((d_266_v44_).cardinality) - (d_222_v4_)
                rhs74_ = d_265_v43_
                lhs62_ = d_261_v39_
                lhs63_ = default__.safeIndex(726, (d_261_v39_).length(0))
                d_260_v38_ = rhs71_
                d_261_v39_ = rhs72_
                lhs62_[lhs63_] = rhs73_
                d_265_v43_ = rhs74_
                (d_221_globalState_).f1 = default__.fm2(d_221_globalState_)
                default__.m0(d_221_globalState_)
                d_267_v45_: _dafny.Array
                def lambda15_(d_268_v42_, d_269_v39_, d_270_v38_):
                    def lambda16_(d_271_i4_):
                        return ((d_268_v42_) + (d_268_v42_)).set(default__.safeIndex((d_269_v39_)[default__.safeIndex(726, (d_269_v39_).length(0))], len((d_268_v42_) + (d_268_v42_))), d_270_v38_.f15)

                    return lambda16_

                init9_ = lambda15_(d_264_v42_, d_261_v39_, d_260_v38_)
                nw49_ = _dafny.Array(None, 10)
                for i0_9_ in range(nw49_.length(0)):
                    nw49_[i0_9_] = init9_(i0_9_)
                d_267_v45_ = nw49_
                index55_ = default__.safeIndex(25, (d_267_v45_).length(0))
                (d_267_v45_)[index55_] = d_264_v42_
                index56_ = default__.safeIndex(25, (d_267_v45_).length(0))
                (d_267_v45_)[index56_] = d_264_v42_
            d_272_v46_: D2
            d_272_v46_ = D2_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_273_i5_ in range(default__.abs(707))]))
            d_274_v47_: _dafny.Seq
            d_274_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqrgy"))
            pat_let_tv5_ = d_274_v47_
            def iife18_(_pat_let5_0):
                def iife19_(d_275_dt__update__tmp_h0_):
                    def iife20_(_pat_let6_0):
                        def iife21_(d_276_dt__update_hcf3_h0_):
                            return D2_DC3(d_276_dt__update_hcf3_h0_)
                        return iife21_(_pat_let6_0)
                    return iife20_((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))) + (pat_let_tv5_))
                return iife19_(_pat_let5_0)
            source3_ = iife18_(d_272_v46_)
            if source3_.is_DC4:
                d_277___mcc_h0_ = source3_.cf4
                d_278___mcc_h1_ = source3_.cf5
                d_279_cf5_ = d_278___mcc_h1_
                d_280_cf4_ = d_277___mcc_h0_
                d_281_v48_: str
                d_281_v48_ = _dafny.CodePoint('d')
                rhs75_ = default__.safeDivisionInt(654, -318)
                rhs76_ = (d_281_v48_) in (d_274_v47_)
                rhs77_ = not((d_222_v4_) <= (d_280_cf4_))
                rhs78_ = d_219_v2_
                rhs79_ = d_280_cf4_
                lhs64_ = d_221_globalState_
                lhs65_ = d_221_globalState_
                lhs66_ = d_221_globalState_
                lhs67_ = d_221_globalState_
                lhs68_ = d_221_globalState_
                lhs64_.f4 = rhs75_
                lhs65_.f1 = rhs76_
                lhs66_.f10 = rhs77_
                lhs67_.f1 = rhs78_
                lhs68_.f6 = rhs79_
                d_282_v49_: _dafny.Seq
                d_282_v49_ = _dafny.SeqWithoutIsStrInference([d_216_v0_])
                d_282_v49_ = _dafny.SeqWithoutIsStrInference([d_216_v0_, (d_218_v1_), d_216_v0_])
                (d_221_globalState_).f6 = d_280_cf4_
                d_283_v50_: D1
                d_283_v50_ = D1_DC2(d_280_cf4_)
                d_284_v51_: _dafny.Map
                d_284_v51_ = _dafny.Map({len(d_274_v47_): len(d_274_v47_)})
                d_285_v52_: _dafny.Map
                d_285_v52_ = _dafny.Map({len(_dafny.Set({d_283_v50_})): d_284_v51_})
                d_286_v53_: _dafny.Set
                d_286_v53_ = _dafny.Set({d_219_v2_, d_219_v2_})
                d_287_v54_: _dafny.Array
                nw50_ = _dafny.Array(None, 14)
                nw50_[int(0)] = d_279_cf5_
                nw50_[int(1)] = -32
                nw50_[int(2)] = default__.safeDivisionInt(893, d_279_cf5_)
                nw50_[int(3)] = d_280_cf4_
                nw50_[int(4)] = default__.fm0(d_219_v2_, d_279_cf5_, default__.fm2(d_221_globalState_), d_221_globalState_)
                nw50_[int(5)] = d_279_cf5_
                nw50_[int(6)] = default__.fm0(d_219_v2_, len(_dafny.SeqWithoutIsStrInference([d_279_cf5_ for d_288_i6_ in range(default__.abs(230))])), d_219_v2_, d_221_globalState_)
                nw50_[int(7)] = d_280_cf4_
                nw50_[int(8)] = d_222_v4_
                nw50_[int(9)] = default__.safeDivisionInt(d_222_v4_, d_222_v4_)
                nw50_[int(10)] = len(d_285_v52_)
                nw50_[int(11)] = d_222_v4_
                nw50_[int(12)] = (d_279_cf5_) + (len(d_286_v53_))
                nw50_[int(13)] = d_222_v4_
                d_287_v54_ = nw50_
                index57_ = default__.safeIndex(331, (d_287_v54_).length(0))
                (d_287_v54_)[index57_] = d_222_v4_
                index58_ = default__.safeIndex(331, (d_287_v54_).length(0))
                (d_287_v54_)[index58_] = (d_280_cf4_) * (d_280_cf4_)
            elif True:
                d_289___mcc_h2_ = source3_.cf3
                d_290_cf3_ = d_289___mcc_h2_
                (d_221_globalState_).f4 = d_222_v4_
                d_291_v55_: _dafny.Map
                d_291_v55_ = _dafny.Map({default__.fm0(d_219_v2_, d_222_v4_, d_219_v2_, d_221_globalState_): d_219_v2_})
                d_292_v56_: _dafny.Map
                d_292_v56_ = _dafny.Map({d_219_v2_: d_291_v55_})
                d_293_v57_: _dafny.Array
                nw51_ = _dafny.Array(None, 17)
                nw51_[int(0)] = (_dafny.Map({default__.fm2(d_221_globalState_): _dafny.Map({d_222_v4_: not(d_219_v2_)})})) | (d_292_v56_)
                nw51_[int(1)] = d_292_v56_
                nw51_[int(2)] = d_292_v56_
                nw51_[int(3)] = d_292_v56_
                nw51_[int(4)] = d_292_v56_
                nw51_[int(5)] = d_292_v56_
                nw51_[int(6)] = (d_292_v56_) | (d_292_v56_)
                nw51_[int(7)] = d_292_v56_
                nw51_[int(8)] = d_292_v56_
                nw51_[int(9)] = d_292_v56_
                nw51_[int(10)] = d_292_v56_
                nw51_[int(11)] = default__.fm5(d_221_globalState_)
                nw51_[int(12)] = (d_292_v56_) | (_dafny.Map({d_219_v2_: d_291_v55_}))
                nw51_[int(13)] = d_292_v56_
                nw51_[int(14)] = (default__.fm5(d_221_globalState_)) | (_dafny.Map({d_219_v2_: d_291_v55_}))
                nw51_[int(15)] = d_292_v56_
                nw51_[int(16)] = d_292_v56_
                d_293_v57_ = nw51_
                index59_ = default__.safeIndex(746, (d_293_v57_).length(0))
                (d_293_v57_)[index59_] = d_292_v56_
                index60_ = default__.safeIndex(746, (d_293_v57_).length(0))
                (d_293_v57_)[index60_] = ((d_292_v56_) | (d_292_v56_)) | (d_292_v56_)
                (d_221_globalState_).f0 = -334
                d_294_v58_: C0
                nw52_ = C0()
                nw52_.ctor__(d_219_v2_)
                d_294_v58_ = nw52_
            d_295_v59_: D1
            d_295_v59_ = D1_DC2(d_222_v4_)
            (d_221_globalState_).f8 = ((d_274_v47_) + (default__.fm6(d_222_v4_, d_219_v2_, d_295_v59_, d_219_v2_, d_221_globalState_))) <= (d_274_v47_)
        elif True:
            index61_ = default__.safeIndex(865, (d_216_v0_).length(0))
            (d_216_v0_)[index61_] = d_219_v2_
            d_296_v60_: _dafny.Seq
            d_296_v60_ = _dafny.SeqWithoutIsStrInference([d_219_v2_])
            index62_ = default__.safeIndex(865, (d_216_v0_).length(0))
            (d_216_v0_)[index62_] = ((d_296_v60_ if d_219_v2_ else _dafny.SeqWithoutIsStrInference([default__.fm2(d_221_globalState_), d_219_v2_]))) <= (d_296_v60_)
            (d_221_globalState_).f6 = (default__.safeModuloInt(-760, (0) - (d_222_v4_))) * ((d_222_v4_ if (d_216_v0_)[default__.safeIndex(865, (d_216_v0_).length(0))] else d_222_v4_))
            d_297_v61_: _dafny.Seq
            d_297_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqluto"))
            d_297_v61_ = d_297_v61_
            d_298_v62_: _dafny.Seq
            d_298_v62_ = _dafny.SeqWithoutIsStrInference([d_222_v4_])
            d_299_v63_: _dafny.Map
            d_299_v63_ = _dafny.Map({(d_298_v62_)[default__.safeIndex(d_222_v4_, len(d_298_v62_))]: (d_216_v0_)[default__.safeIndex(865, (d_216_v0_).length(0))]})
            d_300_v64_: D3
            d_300_v64_ = D3_DC5(d_296_v60_)
            d_299_v63_ = (d_299_v63_).set((0) - ((d_222_v4_) + ((0) - (d_222_v4_))), (d_296_v60_) <= ((d_300_v64_).cf6))
            (d_221_globalState_).f4 = (d_222_v4_) - (d_222_v4_)
        (d_221_globalState_).f1 = d_219_v2_
        d_301_i7_: int
        d_301_i7_ = 0
        with _dafny.label("1"):
            while (d_222_v4_) != (439):
                with _dafny.c_label("1"):
                    if (d_301_i7_) >= (100):
                        raise _dafny.Break("1")
                    d_301_i7_ = (d_301_i7_) + (1)
                    d_302_v65_: _dafny.Seq
                    d_302_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                    (d_221_globalState_).f8 = (d_302_v65_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avcf")))
                    d_303_v66_: C0
                    nw53_ = C0()
                    nw53_.ctor__((d_219_v2_) or (d_219_v2_))
                    d_303_v66_ = nw53_
                    d_303_v66_ = d_303_v66_
                    d_220_v3_ = (d_220_v3_).set(d_219_v2_, d_303_v66_.f15)
                    d_304_v67_: _dafny.Array
                    def lambda17_(d_305_v4_):
                        def lambda18_(d_306_i8_):
                            return (_dafny.SeqWithoutIsStrInference([d_305_v4_])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('o'), _dafny.CodePoint('t'), _dafny.CodePoint('v'), _dafny.CodePoint('n'), _dafny.CodePoint('m')])).cardinality for d_307_i9_ in range(default__.abs(159))]))

                        return lambda18_

                    init10_ = lambda17_(d_222_v4_)
                    nw54_ = _dafny.Array(None, 10)
                    for i0_10_ in range(nw54_.length(0)):
                        nw54_[i0_10_] = init10_(i0_10_)
                    d_304_v67_ = nw54_
                    d_308_v68_: _dafny.Seq
                    d_308_v68_ = _dafny.SeqWithoutIsStrInference([d_222_v4_, d_222_v4_, d_222_v4_, d_222_v4_, d_222_v4_])
                    index63_ = default__.safeIndex(30, (d_304_v67_).length(0))
                    (d_304_v67_)[index63_] = d_308_v68_
                    d_309_v69_: _dafny.Seq
                    d_309_v69_ = _dafny.SeqWithoutIsStrInference([d_303_v66_.f15, d_303_v66_.f15, default__.fm2(d_221_globalState_), d_219_v2_])
                    index64_ = default__.safeIndex(30, (d_304_v67_).length(0))
                    (d_304_v67_)[index64_] = default__.fm7((d_222_v4_) + (d_222_v4_), True, d_309_v69_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhvsj"))), d_221_globalState_)
                    pass
            pass
        d_310_v70_: C0
        nw55_ = C0()
        nw55_.ctor__(False)
        d_310_v70_ = nw55_
        d_311_v71_: D4
        d_311_v71_ = D4_DC9(d_216_v0_, d_310_v70_, not(d_310_v70_.f15))
        d_312_v72_: C0
        nw56_ = C0()
        nw56_.ctor__((d_311_v71_).cf12)
        d_312_v72_ = nw56_
        d_313_v73_: _dafny.Array
        def lambda19_(d_314_v4_):
            def lambda20_(d_315_i10_):
                return default__.safeModuloInt(d_315_i10_, d_314_v4_)

            return lambda20_

        init11_ = lambda19_(d_222_v4_)
        nw57_ = _dafny.Array(None, 21)
        for i0_11_ in range(nw57_.length(0)):
            nw57_[i0_11_] = init11_(i0_11_)
        d_313_v73_ = nw57_
        index65_ = default__.safeIndex(626, (d_313_v73_).length(0))
        (d_313_v73_)[index65_] = d_222_v4_
        index66_ = default__.safeIndex(626, (d_313_v73_).length(0))
        (d_313_v73_)[index66_] = default__.fm0((False) or (d_310_v70_.f15), d_222_v4_, d_312_v72_.f15, d_221_globalState_)
        d_316_v74_: str
        d_316_v74_ = _dafny.CodePoint('f')
        d_317_v75_: _dafny.Map
        d_317_v75_ = _dafny.Map({d_312_v72_.f15: d_316_v74_})
        pat_let_tv6_ = d_313_v73_
        pat_let_tv7_ = d_313_v73_
        d_318_v76_: D4
        def iife22_(_pat_let7_0):
            def iife23_(d_319_dt__update__tmp_h1_):
                def iife24_(_pat_let8_0):
                    def iife25_(d_320_dt__update_hcf7_h0_):
                        return D3_DC6(d_320_dt__update_hcf7_h0_)
                    return iife25_(_pat_let8_0)
                return iife24_((pat_let_tv7_)[default__.safeIndex(626, (pat_let_tv6_).length(0))])
            return iife23_(_pat_let7_0)
        d_318_v76_ = D4_DC10(d_222_v4_, d_316_v74_, d_222_v4_, d_219_v2_, _dafny.SeqWithoutIsStrInference([iife22_(D3_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjuqikydf"))))), default__.fm8(len(_dafny.Map({len(_dafny.Map({d_310_v70_.f15: (d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]})): default__.fm2(d_221_globalState_)})), d_219_v2_, d_221_globalState_)]))
        d_321_v77_: D3
        d_321_v77_ = D3_DC5(_dafny.SeqWithoutIsStrInference([d_312_v72_.f15]))
        pat_let_tv8_ = d_313_v73_
        pat_let_tv9_ = d_313_v73_
        pat_let_tv10_ = d_222_v4_
        pat_let_tv11_ = d_222_v4_
        index67_ = default__.safeIndex(626, (d_313_v73_).length(0))
        index68_ = default__.safeIndex(626, (d_313_v73_).length(0))
        index69_ = default__.safeIndex(626, (d_313_v73_).length(0))
        def lambda21_(source4_):
            if source4_.is_DC6:
                d_322___mcc_h3_ = source4_.cf7
                d_323_cf7_ = d_322___mcc_h3_
                return ((pat_let_tv9_)[default__.safeIndex(626, (pat_let_tv8_).length(0))]) * (d_323_cf7_)
            elif source4_.is_DC5:
                d_324___mcc_h4_ = source4_.cf6
                d_325_cf6_ = d_324___mcc_h4_
                return pat_let_tv10_
            elif True:
                d_326___mcc_h5_ = source4_.cf8
                d_327_cf8_ = d_326___mcc_h5_
                return (0) - (pat_let_tv11_)

        rhs80_ = (d_222_v4_) + (len(_dafny.Map({not(True): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_218_v1_, d_218_v1_, d_218_v1_]))})))
        rhs81_ = (d_222_v4_) - ((d_222_v4_) + (len(d_317_v75_)))
        rhs82_ = d_312_v72_.f15
        rhs83_ = (d_318_v76_).cf13
        rhs84_ = lambda21_(d_321_v77_)
        lhs69_ = d_313_v73_
        lhs70_ = default__.safeIndex(626, (d_313_v73_).length(0))
        lhs71_ = d_313_v73_
        lhs72_ = default__.safeIndex(626, (d_313_v73_).length(0))
        lhs73_ = d_221_globalState_
        lhs74_ = d_313_v73_
        lhs75_ = default__.safeIndex(626, (d_313_v73_).length(0))
        lhs69_[lhs70_] = rhs80_
        lhs71_[lhs72_] = rhs81_
        lhs73_.f1 = rhs82_
        d_222_v4_ = rhs83_
        lhs74_[lhs75_] = rhs84_
        (d_221_globalState_).f10 = (d_222_v4_) >= (d_222_v4_)
        (d_221_globalState_).f6 = ((d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]) + ((0) - (len(_dafny.Map({(d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]: d_222_v4_}))))
        hi0_ = d_222_v4_
        for d_328_i11_ in range(685, hi0_):
            d_329_v78_: _dafny.Map
            d_329_v78_ = _dafny.Map({d_328_i11_: -16})
            d_330_v79_: _dafny.Map
            d_330_v79_ = _dafny.Map({d_316_v74_: d_310_v70_.f15})
            d_331_v80_: _dafny.MultiSet
            d_331_v80_ = _dafny.MultiSet([(d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))], d_222_v4_, d_222_v4_, len(d_330_v79_)])
            d_332_v81_: _dafny.Map
            d_332_v81_ = _dafny.Map({d_310_v70_: 159})
            d_333_v82_: _dafny.Array
            nw58_ = _dafny.Array(_dafny.CodePoint('D'), 17)
            d_333_v82_ = nw58_
            d_334_v83_: _dafny.Seq
            d_334_v83_ = _dafny.SeqWithoutIsStrInference([d_312_v72_.f15, d_310_v70_.f15])
            d_335_v84_: _dafny.Map
            d_335_v84_ = _dafny.Map({d_333_v82_: len(d_334_v83_)})
            d_336_v85_: _dafny.Map
            d_336_v85_ = _dafny.Map({((d_329_v78_)[d_222_v4_] if (d_222_v4_) in (d_329_v78_) else ((d_331_v80_)[8] if (8) in (d_331_v80_) else d_328_i11_)): ((d_332_v81_).set(d_310_v70_, (0) - (len(d_335_v84_)))) | (d_332_v81_)})
            d_336_v85_ = (d_336_v85_).set((0) - ((d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]), (d_332_v81_ if d_310_v70_.f15 else d_332_v81_))
            default__.m0(d_221_globalState_)
            d_331_v80_ = d_331_v80_
            d_337_v86_: _dafny.Seq
            d_337_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_316_v74_ for d_338_i13_ in range(default__.abs(824))])])
            d_339_v87_: _dafny.Seq
            d_339_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqs"))
            d_340_v88_: _dafny.Set
            d_340_v88_ = _dafny.Set({d_219_v2_, d_310_v70_.f15, d_310_v70_.f15, d_312_v72_.f15, d_310_v70_.f15})
            d_341_v89_: D1
            d_341_v89_ = D1_DC2(len(d_340_v88_))
            d_342_v90_: _dafny.Array
            nw59_ = _dafny.Array(None, 11)
            nw59_[int(0)] = _dafny.SeqWithoutIsStrInference([d_316_v74_ for d_343_i12_ in range(default__.abs(687))])
            nw59_[int(1)] = ((d_337_v86_)[default__.safeIndex(d_328_i11_, len(d_337_v86_))]) + (d_339_v87_)
            nw59_[int(2)] = d_339_v87_
            nw59_[int(3)] = d_339_v87_
            nw59_[int(4)] = d_339_v87_
            nw59_[int(5)] = d_339_v87_
            nw59_[int(6)] = _dafny.SeqWithoutIsStrInference([d_316_v74_ for d_344_i14_ in range(default__.abs(847))])
            nw59_[int(7)] = (d_339_v87_) + ((d_339_v87_).set(default__.safeIndex(d_222_v4_, len(d_339_v87_)), d_316_v74_))
            nw59_[int(8)] = (d_339_v87_) + (d_339_v87_)
            nw59_[int(9)] = d_339_v87_
            nw59_[int(10)] = default__.fm6(d_222_v4_, (d_311_v71_).cf12, d_341_v89_, d_219_v2_, d_221_globalState_)
            d_342_v90_ = nw59_
            index70_ = default__.safeIndex(219, (d_342_v90_).length(0))
            (d_342_v90_)[index70_] = d_339_v87_
            index71_ = default__.safeIndex(219, (d_342_v90_).length(0))
            rhs85_ = d_310_v70_.f15
            rhs86_ = d_310_v70_.f15
            rhs87_ = (d_219_v2_) or (not((d_328_i11_) <= (d_328_i11_)))
            rhs88_ = d_339_v87_
            rhs89_ = (((d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]) + ((d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))])) * (d_222_v4_)
            lhs76_ = d_312_v72_
            lhs77_ = d_221_globalState_
            lhs78_ = d_221_globalState_
            lhs79_ = d_342_v90_
            lhs80_ = default__.safeIndex(219, (d_342_v90_).length(0))
            lhs81_ = d_221_globalState_
            lhs76_.f15 = rhs85_
            lhs77_.f10 = rhs86_
            lhs78_.f10 = rhs87_
            lhs79_[lhs80_] = rhs88_
            lhs81_.f0 = rhs89_
        d_345_v91_: _dafny.Seq
        d_345_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymky"))
        hi1_ = default__.fm0(d_310_v70_.f15, len(d_345_v91_), default__.fm2(d_221_globalState_), d_221_globalState_)
        for d_346_i15_ in range(default__.safeModuloInt(d_222_v4_, (d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]), hi1_):
            d_347_v92_: _dafny.Map
            d_347_v92_ = _dafny.Map({d_219_v2_: d_312_v72_})
            d_348_v93_: _dafny.Set
            d_348_v93_ = _dafny.Set({(d_347_v92_).set(d_310_v70_.f15, d_312_v72_), d_347_v92_})
            (d_221_globalState_).f6 = len(d_348_v93_)
            d_313_v73_ = d_313_v73_
            default__.m0(d_221_globalState_)
            d_349_v94_: _dafny.Seq
            d_349_v94_ = _dafny.SeqWithoutIsStrInference([d_310_v70_.f15])
            d_350_v95_: _dafny.Seq
            d_350_v95_ = _dafny.SeqWithoutIsStrInference([(d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]])
            d_351_v96_: D1
            d_351_v96_ = D1_DC2((_dafny.MultiSet(d_350_v95_)).cardinality)
            pat_let_tv12_ = d_312_v72_
            pat_let_tv13_ = d_222_v4_
            pat_let_tv14_ = d_312_v72_
            pat_let_tv15_ = d_221_globalState_
            def iife26_(_pat_let9_0):
                def iife27_(d_352_dt__update__tmp_h2_):
                    def iife28_(_pat_let10_0):
                        def iife29_(d_353_dt__update_hcf2_h0_):
                            return D1_DC2(d_353_dt__update_hcf2_h0_)
                        return iife29_(_pat_let10_0)
                    return iife28_(default__.fm0(pat_let_tv12_.f15, pat_let_tv13_, pat_let_tv14_.f15, pat_let_tv15_))
                return iife27_(_pat_let9_0)
            d_345_v91_ = (default__.fm6(882, (d_349_v94_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(d_345_v91_), (d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))], (d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))], len(d_223_v5_)])), len(d_349_v94_))], iife26_(d_351_v96_), d_219_v2_, d_221_globalState_)) + (d_345_v91_)
        index72_ = default__.safeIndex(626, (d_313_v73_).length(0))
        rhs90_ = (d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]
        rhs91_ = not (d_219_v2_) or (d_310_v70_.f15)
        lhs82_ = d_313_v73_
        lhs83_ = default__.safeIndex(626, (d_313_v73_).length(0))
        lhs82_[lhs83_] = rhs90_
        d_219_v2_ = rhs91_
        d_354_v97_: _dafny.MultiSet
        d_354_v97_ = _dafny.MultiSet([d_310_v70_.f15, False])
        d_355_v98_: C0
        nw60_ = C0()
        nw60_.ctor__((d_354_v97_).issubset(d_354_v97_))
        d_355_v98_ = nw60_
        d_356_v99_: _dafny.Seq
        d_356_v99_ = _dafny.SeqWithoutIsStrInference([d_312_v72_])
        d_357_v100_: _dafny.Seq
        d_357_v100_ = _dafny.SeqWithoutIsStrInference([len(d_356_v99_)])
        d_316_v74_ = default__.fm9(default__.safeModuloInt(len(d_357_v100_), (d_313_v73_)[default__.safeIndex(626, (d_313_v73_).length(0))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxgmqx")), d_222_v4_, d_221_globalState_)
        (d_221_globalState_).f10 = not(not(d_355_v98_.f15))
        d_345_v91_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbo"))) + (d_345_v91_)).set(default__.safeIndex(d_222_v4_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbo"))) + (d_345_v91_))), _dafny.CodePoint('e'))) + (d_345_v91_)
        _dafny.print(_dafny.string_of((d_216_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_v1_))[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_219_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_220_v3_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_221_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_221_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_221_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_221_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_221_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f9)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_221_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_.f11)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_globalState_).f12) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_222_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_223_v5_) == (_dafny.Map({True: 121}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_301_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_310_v70_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v71_).cf10)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v71_).cf11.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v71_).cf12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_312_v72_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v73_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_316_v74_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_317_v75_) == (_dafny.Map({True: _dafny.CodePoint('f')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v76_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v76_).cf14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v76_).cf15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v76_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_318_v76_).cf17) == (_dafny.SeqWithoutIsStrInference([D3_DC6(121), D3_DC6(522)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_321_v77_).cf6) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_345_v91_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_354_v97_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_355_v98_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_356_v99_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v100_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
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
        return lambda: D1_DC2(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
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
        return lambda: D2_DC4(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)

class D2_DC4(D2, NamedTuple('DC4', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({self.cf3.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC6(int(0))
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

class D3_DC6(D3, NamedTuple('DC6', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC5(D3, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC9(_dafny.Array(None, 0), None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC9(D4, NamedTuple('DC9', [('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC14(D5, NamedTuple('DC14', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(False, None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC18(D7, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC20(D8, NamedTuple('DC20', [('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f1: bool = False
        self.f4: int = int(0)
        self.f6: int = int(0)
        self.f8: bool = False
        self.f9: _dafny.Array = _dafny.Array(None, 0)
        self.f10: bool = False
        self.f11: _dafny.Array = _dafny.Array(None, 0)
        self._f2: bool = False
        self._f3: int = int(0)
        self._f5: int = int(0)
        self._f7: bool = False
        self._f12: _dafny.Map = _dafny.Map({})
        self._f13: int = int(0)
        self._f14: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14):
        (self).f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self).f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14

    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f5(self):
        return self._f5
    @property
    def f7(self):
        return self._f7
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14

class C0:
    def  __init__(self):
        self.f15: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f15):
        (self).f15 = f15

    def fm1(self, p0, p1, globalState):
        return _dafny.Map({(len(_dafny.Map({self.f15: -178}))) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbuyqaqah")))): self.f15})

