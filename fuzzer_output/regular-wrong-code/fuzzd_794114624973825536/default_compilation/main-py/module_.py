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
    def fm2(p0, p1, p2, globalState):
        return ((0) - ((len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): _dafny.CodePoint('n')})) if True else 98))) <= ((0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdayefsjg"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_0_i0_ in range(default__.abs(543))])))))

    @staticmethod
    def fm3(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({True: not(True)}) for d_1_i0_ in range(default__.abs(988))])

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return D1_DC1(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}) for d_2_i0_ in range(default__.abs(61))]))

    @staticmethod
    def fm5(p0, globalState):
        return _dafny.CodePoint('o')

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([False, not((_dafny.Map({True: True})) == (_dafny.Map({True: True}))), not(False)])

    @staticmethod
    def fm8(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_3_i0_ in range(default__.abs(32))])

    @staticmethod
    def fm9(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(468, -404):
                d_4_v0_: int = compr_0_
                if ((468) <= (d_4_v0_)) and ((d_4_v0_) < (-404)):
                    coll0_[(d_4_v0_) * (203)] = 415
            return _dafny.Map(coll0_)
        return ((0) - (default__.safeDivisionInt((_dafny.MultiSet([len(_dafny.Set({True}))])).cardinality, len(iife0_()
        )))) + (665)

    @staticmethod
    def fm10(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False, True])))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjyj")))), -123])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmifknqvs")))]))) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([781])) + (_dafny.SeqWithoutIsStrInference([406 for d_5_i0_ in range(default__.abs(-166))]))))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxkcy")))) for d_6_i0_ in range(default__.abs(684))])) + (_dafny.SeqWithoutIsStrInference([-110 for d_7_i1_ in range(default__.abs(173))]))) + (_dafny.SeqWithoutIsStrInference([194]))

    @staticmethod
    def fm15(p0, p1, globalState):
        return ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([57, 825]), _dafny.SeqWithoutIsStrInference([897 for d_8_i0_ in range(default__.abs(13))]), _dafny.SeqWithoutIsStrInference([-853, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "inrv")))])])), -354, -817])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([4])))) | ((_dafny.MultiSet([len(_dafny.Map({True: 5})), (0) - ((_dafny.MultiSet([False])).cardinality), len(_dafny.Map({True: False})), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvdonglnj"))): False}))])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([527 for d_9_i1_ in range(default__.abs(577))]))])))

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        return ((_dafny.Set({False, True, False})) - (_dafny.Set({not(True)}))).intersection((_dafny.Set({False, True, (D4_DC9(293, _dafny.CodePoint('i'), _dafny.SeqWithoutIsStrInference([767 for d_10_i0_ in range(default__.abs(726))]), 37, False)).cf20, True}) if True else _dafny.Set({True, True})))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        return ((0) - ((190) - ((0) - (-831)))) * ((0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([False, True])), len(_dafny.Map({not(False): _dafny.CodePoint('r')})))))

    @staticmethod
    def fm18(p0, p1, globalState):
        return D1_DC2(503, False, (len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([True, not(True)]))): True}))) + (829))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        return (len((_dafny.SeqWithoutIsStrInference([not(not(True)), False, False])) + (_dafny.SeqWithoutIsStrInference([True])))) < ((0) - (-948))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return D2_DC3((_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([925 for d_11_i0_ in range(default__.abs(884))])): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqwbp"))) for d_12_i1_ in range(default__.abs(78))]))})) | (_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([837 for d_13_i2_ in range(default__.abs(769))]))).cardinality, 471])): 819})))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return (len((_dafny.SeqWithoutIsStrInference([394])) + (_dafny.SeqWithoutIsStrInference([904])))) >= ((239) - (156))

    @staticmethod
    def fm24(p0, globalState):
        return (0) - (((938) * (len(_dafny.SeqWithoutIsStrInference([False])))) * ((634) * (619)))

    @staticmethod
    def fm25(p0, p1, globalState):
        source0_ = D13_DC25((0) - ((0) - ((0) - (len(_dafny.Map({True: False}))))), 787, False)
        if source0_.is_DC25:
            d_14___mcc_h0_ = source0_.cf42
            d_15___mcc_h1_ = source0_.cf43
            d_16___mcc_h2_ = source0_.cf44
            d_17_cf44_ = d_16___mcc_h2_
            d_18_cf43_ = d_15___mcc_h1_
            d_19_cf42_ = d_14___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "li"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggkk")))
        elif True:
            d_20___mcc_h3_ = source0_.cf41
            d_21_cf41_ = d_20___mcc_h3_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnno"))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return _dafny.CodePoint('v')

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hls"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhuwclrv"))):
            return _dafny.SeqWithoutIsStrInference([(0) - (-387) for d_22_i0_ in range(default__.abs(472))])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([((D14_DC26(_dafny.MultiSet([False]))).cf45).cardinality])) + ((D4_DC9(317, _dafny.CodePoint('l'), _dafny.SeqWithoutIsStrInference([111, len(_dafny.Map({True: _dafny.CodePoint('v')}))]), 447, not(True))).cf18)

    @staticmethod
    def fm28(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(24, 239):
                d_23_v0_: int = compr_1_
                if ((24) <= (d_23_v0_)) and ((d_23_v0_) < (239)):
                    coll1_[(d_23_v0_) * (758)] = 627
            return _dafny.Map(coll1_)
        return ((_dafny.Map({len(_dafny.Map({len(_dafny.Set({True})): len(_dafny.SeqWithoutIsStrInference([True, False]))})): 729})) | (iife1_()
        )) | (_dafny.Map({167: len(_dafny.SeqWithoutIsStrInference([True]))}))

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        if not(not(True)):
            return D2_DC5(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True)])), -659])), not(True), True, True, not(False))
        elif True:
            return D2_DC5(63, True, not(False), False, False)

    @staticmethod
    def fm31(p0, p1, p2, globalState):
        return (_dafny.Set({False, False})) | ((_dafny.Set({False, True})) | (_dafny.Set({True})))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False, not(not(not(True)))])) + (_dafny.SeqWithoutIsStrInference([True, True]))) + (_dafny.SeqWithoutIsStrInference([not(True)]))

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, False])]))

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        return ((_dafny.Map({(D3_DC7(397, len(_dafny.SeqWithoutIsStrInference([not(True)])))).cf14: False})) | (_dafny.Map({len(_dafny.Set({False})): True}))) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([637])).cardinality, 462])): False})) | (_dafny.Map({449: not(True)})))

    @staticmethod
    def fm35(globalState):
        if not (False) or (True):
            return _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_24_i0_ in range(default__.abs(397))])})
        elif True:
            return _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_25_i1_ in range(default__.abs(-615))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))})

    @staticmethod
    def fm36(p0, globalState):
        return D7_DC14((814) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqkmgu")))))

    @staticmethod
    def fm37(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.Map({True: False})) | (_dafny.Map({False: False}))])

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        return (_dafny.Set({len(_dafny.Set({len(_dafny.Map({not(False): (0) - (-689)})), -952, 228}))})) | ((_dafny.Set({621})) - (_dafny.Set({-276})))

    @staticmethod
    def fm39(p0, globalState):
        return ((_dafny.Map({not(False): True})) | (_dafny.Map({False: True}))) | ((_dafny.Map({False: not(True)})) | (_dafny.Map({False: not(True)})))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([not(not(True)), False])).intersection(_dafny.MultiSet([True]))).intersection((_dafny.MultiSet([False])).intersection(_dafny.MultiSet([False])))

    @staticmethod
    def fm41(p0, p1, globalState):
        source1_ = D1_DC2(len(_dafny.SeqWithoutIsStrInference([True, True])), True, len(_dafny.Map({True: True})))
        if source1_.is_DC2:
            d_26___mcc_h0_ = source1_.cf2
            d_27___mcc_h1_ = source1_.cf3
            d_28___mcc_h2_ = source1_.cf4
            d_29_cf4_ = d_28___mcc_h2_
            d_30_cf3_ = d_27___mcc_h1_
            d_31_cf2_ = d_26___mcc_h0_
            return D8_DC16(d_29_cf4_, d_30_cf3_, d_30_cf3_, d_31_cf2_)
        elif True:
            d_32___mcc_h3_ = source1_.cf1
            d_33_cf1_ = d_32___mcc_h3_
            return D8_DC16(len(_dafny.Set({False})), False, True, len(_dafny.Set({not(False)})))

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        def iife2_():
            def iife6_():
                def iife8_():
                    coll8_ = _dafny.Map()
                    compr_8_: _dafny.Seq
                    for compr_8_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prgm"))])).Elements:
                        d_34_v1_: _dafny.Seq = compr_8_
                        if (d_34_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prgm"))])):
                            coll8_[d_34_v1_] = False
                    return _dafny.Map(coll8_)
                coll6_ = _dafny.Set()
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: _dafny.Seq
                    for compr_7_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prgm"))])).Elements:
                        d_34_v1_: _dafny.Seq = compr_7_
                        if (d_34_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prgm"))])):
                            coll7_[d_34_v1_] = False
                    return _dafny.Map(coll7_)
                compr_6_: _dafny.Seq
                for compr_6_ in (iife7_()
                ).keys.Elements:
                    d_38_v2_: _dafny.Seq = compr_6_
                    if (d_38_v2_) in (iife8_()
                    ):
                        coll6_ = coll6_.union(_dafny.Set([d_38_v2_]))
                return _dafny.Set(coll6_)
            coll2_ = _dafny.Map()
            def iife3_():
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: _dafny.Seq
                    for compr_5_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prgm"))])).Elements:
                        d_34_v1_: _dafny.Seq = compr_5_
                        if (d_34_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prgm"))])):
                            coll5_[d_34_v1_] = False
                    return _dafny.Map(coll5_)
                coll3_ = _dafny.Set()
                def iife4_():
                    coll4_ = _dafny.Map()
                    compr_4_: _dafny.Seq
                    for compr_4_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prgm"))])).Elements:
                        d_34_v1_: _dafny.Seq = compr_4_
                        if (d_34_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prgm"))])):
                            coll4_[d_34_v1_] = False
                    return _dafny.Map(coll4_)
                compr_3_: _dafny.Seq
                for compr_3_ in (iife4_()
                ).keys.Elements:
                    d_35_v2_: _dafny.Seq = compr_3_
                    if (d_35_v2_) in (iife5_()
                    ):
                        coll3_ = coll3_.union(_dafny.Set([d_35_v2_]))
                return _dafny.Set(coll3_)
            compr_2_: _dafny.Seq
            for compr_2_ in ((iife3_()
            ) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_36_i0_ in range(default__.abs(580))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbeim"))}))).Elements:
                d_37_v0_: _dafny.Seq = compr_2_
                if (d_37_v0_) in ((iife6_()
                ) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_36_i0_ in range(default__.abs(580))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbeim"))}))):
                    coll2_[d_37_v0_] = (_dafny.Map({_dafny.CodePoint('w'): True})) | (_dafny.Map({_dafny.CodePoint('i'): False}))
            return _dafny.Map(coll2_)
        return iife2_()
        

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return ((_dafny.Map({_dafny.MultiSet([False]): _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snxgcwim"))), 430])})) | (_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False])): _dafny.MultiSet([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejalkcx")))])).cardinality])}))) | (_dafny.Map({_dafny.MultiSet([not(False)]): _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('r'), _dafny.CodePoint('y'), _dafny.CodePoint('p'), _dafny.CodePoint('r')]))])}))

    @staticmethod
    def fm44(globalState):
        return (_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm45(p0, p1, globalState):
        return D9_DC20(not(not((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_39_i0_ in range(default__.abs(429))])) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fc"))))))

    @staticmethod
    def m0(p0, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_40_v0_: bool
        d_40_v0_ = True
        d_41_v1_: str
        d_41_v1_ = _dafny.CodePoint('r')
        d_42_v2_: _dafny.Seq
        d_42_v2_ = _dafny.SeqWithoutIsStrInference([d_41_v1_, d_41_v1_])
        d_43_v3_: _dafny.Seq
        d_43_v3_ = _dafny.SeqWithoutIsStrInference([d_42_v2_])
        d_44_v4_: _dafny.Array
        nw0_ = _dafny.Array(False, 13)
        d_44_v4_ = nw0_
        d_45_v5_: _dafny.Map
        d_45_v5_ = _dafny.Map({d_40_v0_: d_44_v4_})
        d_46_v6_: _dafny.Array
        nw1_ = _dafny.Array(_dafny.Map({}), 13)
        d_46_v6_ = nw1_
        d_47_v7_: C4
        nw2_ = C4()
        nw2_.ctor__((d_40_v0_) == (d_40_v0_), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (default__.fm17(d_43_v3_, len(d_45_v5_), p0, len(d_42_v2_), globalState))])), d_46_v6_)
        d_47_v7_ = nw2_
        d_40_v0_ = (d_47_v7_).f7
        d_48_v8_: _dafny.Seq
        d_48_v8_ = _dafny.SeqWithoutIsStrInference([(d_47_v7_).f7, (d_47_v7_).f7, d_40_v0_])
        d_49_v9_: _dafny.Map
        d_49_v9_ = _dafny.Map({(d_47_v7_).f7: (_dafny.MultiSet(d_48_v8_)).cardinality})
        d_50_v10_: C6
        nw3_ = C6()
        nw3_.ctor__((d_49_v9_) | ((_dafny.Map({d_40_v0_: p0})).set((d_47_v7_).f7, p0)), p0)
        d_50_v10_ = nw3_
        d_51_v11_: C5
        nw4_ = C5()
        nw4_.ctor__(_dafny.CodePoint('m'))
        d_51_v11_ = nw4_
        d_52_v12_: _dafny.Map
        d_52_v12_ = _dafny.Map({d_51_v11_: d_50_v10_.f5})
        d_52_v12_ = (d_52_v12_).set(d_51_v11_, 821)
        (d_50_v10_).f5 = default__.safeDivisionInt(p0, d_50_v10_.f5)
        d_53_v13_: C2
        nw5_ = C2()
        nw5_.ctor__(d_40_v0_, (0) - (p0), d_46_v6_)
        d_53_v13_ = nw5_
        d_54_v14_: _dafny.Array
        nw6_ = _dafny.Array(None, 20)
        nw6_[int(0)] = d_53_v13_
        nw6_[int(1)] = d_53_v13_
        nw6_[int(2)] = d_53_v13_
        nw6_[int(3)] = d_53_v13_
        nw6_[int(4)] = d_53_v13_
        nw6_[int(5)] = d_53_v13_
        nw6_[int(6)] = d_53_v13_
        nw6_[int(7)] = d_53_v13_
        nw6_[int(8)] = d_53_v13_
        nw6_[int(9)] = d_53_v13_
        nw6_[int(10)] = d_53_v13_
        nw6_[int(11)] = d_53_v13_
        nw6_[int(12)] = d_53_v13_
        nw6_[int(13)] = d_53_v13_
        nw6_[int(14)] = d_53_v13_
        nw6_[int(15)] = d_53_v13_
        nw6_[int(16)] = d_53_v13_
        nw6_[int(17)] = d_53_v13_
        nw6_[int(18)] = d_53_v13_
        nw6_[int(19)] = d_53_v13_
        d_54_v14_ = nw6_
        index0_ = default__.safeIndex(813, (d_54_v14_).length(0))
        (d_54_v14_)[index0_] = d_53_v13_
        index1_ = default__.safeIndex(813, (d_54_v14_).length(0))
        (d_54_v14_)[index1_] = d_53_v13_
        r0 = False
        r1 = ((d_53_v13_).f12) < (p0)
        d_55_v15_: D6
        d_55_v15_ = D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tw")))
        pat_let_tv0_ = d_42_v2_
        pat_let_tv1_ = d_42_v2_
        def lambda0_(source2_):
            if source2_.is_DC12:
                d_56___mcc_h0_ = source2_.cf23
                d_57___mcc_h1_ = source2_.cf24
                d_58_cf24_ = d_57___mcc_h1_
                d_59_cf23_ = d_56___mcc_h0_
                return pat_let_tv0_
            elif True:
                d_60___mcc_h2_ = source2_.cf22
                d_61_cf22_ = d_60___mcc_h2_
                return pat_let_tv1_

        r2 = len(lambda0_(d_55_v15_))
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_62_v0_: bool
        d_62_v0_ = False
        d_63_v1_: _dafny.Seq
        d_63_v1_ = _dafny.SeqWithoutIsStrInference([d_62_v0_])
        d_64_v2_: _dafny.Seq
        d_64_v2_ = d_63_v1_
        d_65_v3_: int
        d_65_v3_ = 124
        d_66_v4_: _dafny.Array
        nw7_ = _dafny.Array(None, 11)
        nw7_[int(0)] = d_63_v1_
        nw7_[int(1)] = d_63_v1_
        nw7_[int(2)] = (d_64_v2_)
        nw7_[int(3)] = d_63_v1_
        nw7_[int(4)] = (d_63_v1_).set(default__.safeIndex(d_65_v3_, len(d_63_v1_)), d_62_v0_)
        nw7_[int(5)] = d_63_v1_
        nw7_[int(6)] = d_63_v1_
        nw7_[int(7)] = d_63_v1_
        nw7_[int(8)] = d_63_v1_
        nw7_[int(9)] = _dafny.SeqWithoutIsStrInference([d_62_v0_])
        nw7_[int(10)] = d_63_v1_
        d_66_v4_ = nw7_
        d_67_globalState_: GlobalState
        nw8_ = GlobalState()
        nw8_.ctor__(d_66_v4_, False, 47)
        d_67_globalState_ = nw8_
        d_68_v5_: _dafny.Set
        d_68_v5_ = _dafny.Set({d_65_v3_})
        if (d_68_v5_).ispropersubset(d_68_v5_):
            d_69_v6_: bool
            d_70_v7_: bool
            d_71_v8_: int
            out0_: bool
            out1_: bool
            out2_: int
            out0_, out1_, out2_ = default__.m0(d_65_v3_, d_67_globalState_)
            d_69_v6_ = out0_
            d_70_v7_ = out1_
            d_71_v8_ = out2_
            d_62_v0_ = d_62_v0_
            if not(d_70_v7_):
                d_72_v9_: _dafny.Map
                d_72_v9_ = _dafny.Map({-674: False})
                d_73_v11_: _dafny.Seq
                d_73_v11_ = _dafny.SeqWithoutIsStrInference([d_65_v3_])
                d_74_v12_: _dafny.Map
                d_74_v12_ = d_72_v9_
                d_75_v13_: str
                d_75_v13_ = _dafny.CodePoint('v')
                d_76_v14_: _dafny.Seq
                d_76_v14_ = _dafny.SeqWithoutIsStrInference([d_75_v13_])
                d_77_v17_: _dafny.Array
                nw9_ = _dafny.Array(None, 17)
                nw9_[int(0)] = d_72_v9_
                nw9_[int(1)] = d_72_v9_
                nw9_[int(2)] = d_72_v9_
                nw9_[int(3)] = d_72_v9_
                nw9_[int(4)] = d_72_v9_
                def iife9_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in (d_73_v11_).Elements:
                        d_78_v10_: int = compr_9_
                        if (d_78_v10_) in (d_73_v11_):
                            coll9_[(d_78_v10_) * (789)] = d_70_v7_
                    return _dafny.Map(coll9_)
                nw9_[int(5)] = iife9_()
                
                nw9_[int(6)] = d_72_v9_
                nw9_[int(7)] = (d_74_v12_)
                nw9_[int(8)] = default__.fm34(d_62_v0_, d_75_v13_, d_73_v11_, d_67_globalState_)
                nw9_[int(9)] = _dafny.Map({default__.fm17(_dafny.SeqWithoutIsStrInference([d_76_v14_, d_76_v14_]), d_71_v8_, d_71_v8_, d_71_v8_, d_67_globalState_): d_69_v6_})
                nw9_[int(10)] = d_72_v9_
                nw9_[int(11)] = d_72_v9_
                def iife10_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in (_dafny.Map({d_71_v8_: d_71_v8_})).keys.Elements:
                        d_79_v15_: int = compr_10_
                        if (d_79_v15_) in (_dafny.Map({d_71_v8_: d_71_v8_})):
                            coll10_[(d_79_v15_) + (d_65_v3_)] = d_62_v0_
                    return _dafny.Map(coll10_)
                nw9_[int(12)] = iife10_()
                
                def iife11_():
                    coll11_ = _dafny.Map()
                    compr_11_: int
                    for compr_11_ in (d_73_v11_).Elements:
                        d_80_v16_: int = compr_11_
                        if (d_80_v16_) in (d_73_v11_):
                            coll11_[(d_80_v16_) - (d_71_v8_)] = d_70_v7_
                    return _dafny.Map(coll11_)
                nw9_[int(13)] = iife11_()
                
                nw9_[int(14)] = _dafny.Map({d_65_v3_: d_70_v7_})
                nw9_[int(15)] = d_72_v9_
                nw9_[int(16)] = d_72_v9_
                d_77_v17_ = nw9_
                d_81_v18_: C2
                nw10_ = C2()
                nw10_.ctor__(d_69_v6_, (0) - (d_65_v3_), d_77_v17_)
                d_81_v18_ = nw10_
                d_82_v19_: D6
                d_82_v19_ = D6_DC11(d_76_v14_)
                d_83_v20_: _dafny.Array
                nw11_ = _dafny.Array(None, 7)
                nw11_[int(0)] = d_82_v19_
                nw11_[int(1)] = d_82_v19_
                nw11_[int(2)] = d_82_v19_
                nw11_[int(3)] = d_82_v19_
                nw11_[int(4)] = d_82_v19_
                nw11_[int(5)] = d_82_v19_
                nw11_[int(6)] = d_82_v19_
                d_83_v20_ = nw11_
                rhs0_ = (d_81_v18_).f12
                rhs1_ = d_65_v3_
                rhs2_ = (d_83_v20_ if d_62_v0_ else d_83_v20_)
                rhs3_ = (d_65_v3_) + (d_65_v3_)
                lhs0_ = d_67_globalState_
                lhs1_ = d_67_globalState_
                lhs0_.f2 = rhs0_
                lhs1_.f2 = rhs1_
                d_83_v20_ = rhs2_
                d_65_v3_ = rhs3_
                d_69_v6_ = True
                d_84_v21_: _dafny.Map
                d_84_v21_ = _dafny.Map({(d_81_v18_).f11: (d_81_v18_).f12})
                d_85_v22_: _dafny.Map
                d_85_v22_ = _dafny.Map({not(d_69_v6_): ((d_84_v21_)[(d_81_v18_).f11] if ((d_81_v18_).f11) in (d_84_v21_) else d_71_v8_)})
                d_71_v8_ = ((d_85_v22_)[True] if (True) in (d_85_v22_) else 325)
                d_85_v22_ = (d_85_v22_).set((_dafny.SeqWithoutIsStrInference([d_75_v13_ for d_86_i0_ in range(default__.abs(218))])) < (d_76_v14_), d_71_v8_)
            elif True:
                d_87_v23_: _dafny.Array
                nw12_ = _dafny.Array(None, 6)
                d_87_v23_ = nw12_
                d_87_v23_ = d_87_v23_
                d_88_v24_: _dafny.Seq
                d_88_v24_ = _dafny.SeqWithoutIsStrInference([d_65_v3_, d_71_v8_, d_71_v8_])
                d_89_v25_: _dafny.Seq
                d_89_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ro"))
                d_90_v26_: str
                d_90_v26_ = _dafny.CodePoint('j')
                d_91_v27_: _dafny.Array
                nw13_ = _dafny.Array(None, 9)
                nw13_[int(0)] = (d_70_v7_ if d_62_v0_ else d_62_v0_)
                nw13_[int(1)] = d_69_v6_
                nw13_[int(2)] = default__.fm21(len(d_89_v25_), d_90_v26_, False, d_70_v7_, d_67_globalState_)
                nw13_[int(3)] = d_70_v7_
                nw13_[int(4)] = False
                nw13_[int(5)] = (default__.fm21(d_71_v8_, d_90_v26_, d_70_v7_, not(d_70_v7_), d_67_globalState_) if d_69_v6_ else d_69_v6_)
                nw13_[int(6)] = d_62_v0_
                nw13_[int(7)] = not((not(default__.fm2(_dafny.CodePoint('f'), _dafny.SeqWithoutIsStrInference([_dafny.Map({d_70_v7_: d_70_v7_}) for d_92_i1_ in range(default__.abs(745))]), d_70_v7_, d_67_globalState_)) if d_70_v7_ else d_70_v7_))
                nw13_[int(8)] = not (True) or (d_70_v7_)
                d_91_v27_ = nw13_
                rhs4_ = d_88_v24_
                rhs5_ = ((d_69_v6_) == (d_70_v7_) if d_69_v6_ else d_62_v0_)
                rhs6_ = d_91_v27_
                rhs7_ = (d_88_v24_) != (d_88_v24_)
                rhs8_ = d_65_v3_
                d_88_v24_ = rhs4_
                d_69_v6_ = rhs5_
                d_91_v27_ = rhs6_
                d_69_v6_ = rhs7_
                d_71_v8_ = rhs8_
                d_93_v28_: _dafny.Array
                nw14_ = _dafny.Array(_dafny.Map({}), 13)
                d_93_v28_ = nw14_
                d_94_v29_: C1
                nw15_ = C1()
                nw15_.ctor__(d_71_v8_, d_93_v28_)
                d_94_v29_ = nw15_
                d_95_v30_: _dafny.Array
                nw16_ = _dafny.Array(None, 14)
                nw16_[int(0)] = d_64_v2_
                nw16_[int(1)] = d_64_v2_
                nw16_[int(2)] = d_64_v2_
                nw16_[int(3)] = d_64_v2_
                nw16_[int(4)] = default__.fm44(d_67_globalState_)
                nw16_[int(5)] = default__.fm44(d_67_globalState_)
                nw16_[int(6)] = d_64_v2_
                nw16_[int(7)] = d_64_v2_
                nw16_[int(8)] = d_64_v2_
                nw16_[int(9)] = default__.fm44(d_67_globalState_)
                nw16_[int(10)] = d_64_v2_
                nw16_[int(11)] = d_64_v2_
                nw16_[int(12)] = d_64_v2_
                nw16_[int(13)] = d_64_v2_
                d_95_v30_ = nw16_
                index2_ = default__.safeIndex(154, (d_95_v30_).length(0))
                (d_95_v30_)[index2_] = d_64_v2_
                index3_ = default__.safeIndex(154, (d_95_v30_).length(0))
                rhs9_ = d_64_v2_
                rhs10_ = len((d_89_v25_) + (d_89_v25_))
                rhs11_ = (d_94_v29_).f14
                rhs12_ = (d_63_v1_)[default__.safeIndex((d_94_v29_).f14, len(d_63_v1_))]
                rhs13_ = d_90_v26_
                lhs2_ = d_95_v30_
                lhs3_ = default__.safeIndex(154, (d_95_v30_).length(0))
                lhs4_ = d_67_globalState_
                lhs2_[lhs3_] = rhs9_
                d_65_v3_ = rhs10_
                lhs4_.f2 = rhs11_
                d_70_v7_ = rhs12_
                d_90_v26_ = rhs13_
                d_96_v31_: _dafny.Seq
                d_96_v31_ = _dafny.SeqWithoutIsStrInference([d_89_v25_])
                d_97_v32_: _dafny.Set
                d_97_v32_ = _dafny.Set({d_91_v27_, d_91_v27_})
                d_98_v33_: _dafny.Set
                d_98_v33_ = _dafny.Set({d_94_v29_})
                d_69_v6_ = ((len(d_89_v25_)) - (default__.fm17(d_96_v31_, (d_88_v24_)[default__.safeIndex((d_94_v29_).f14, len(d_88_v24_))], len(d_97_v32_), (d_94_v29_).f14, d_67_globalState_))) <= (default__.fm17(_dafny.SeqWithoutIsStrInference([d_89_v25_, d_89_v25_, (d_89_v25_).set(default__.safeIndex(653, len(d_89_v25_)), _dafny.CodePoint('b'))]), (d_94_v29_).f14, len(_dafny.Map({d_65_v3_: len(d_98_v33_)})), d_65_v3_, d_67_globalState_))
            d_62_v0_ = d_70_v7_
            d_99_v34_: bool
            d_100_v35_: bool
            d_101_v36_: int
            out3_: bool
            out4_: bool
            out5_: int
            out3_, out4_, out5_ = default__.m0((0) - (d_65_v3_), d_67_globalState_)
            d_99_v34_ = out3_
            d_100_v35_ = out4_
            d_101_v36_ = out5_
        elif True:
            d_62_v0_ = d_62_v0_
            d_102_v37_: _dafny.Map
            d_102_v37_ = _dafny.Map({d_62_v0_: not (d_62_v0_) or (d_62_v0_)})
            d_103_v38_: str
            d_103_v38_ = _dafny.CodePoint('b')
            d_102_v37_ = (d_102_v37_).set(default__.fm21(d_65_v3_, d_103_v38_, d_62_v0_, False, d_67_globalState_), d_62_v0_)
            d_104_v39_: _dafny.Seq
            d_104_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgobb"))
            d_104_v39_ = (_dafny.SeqWithoutIsStrInference([d_103_v38_ for d_105_i2_ in range(default__.abs(440))])) + (d_104_v39_)
            d_62_v0_ = ((d_68_v5_).isdisjoint(d_68_v5_) if d_62_v0_ else d_62_v0_)
            if d_62_v0_:
                d_62_v0_ = (False) and (d_62_v0_)
                d_106_v40_: _dafny.Map
                d_106_v40_ = _dafny.Map({d_65_v3_: d_65_v3_})
                d_107_v41_: bool
                d_108_v42_: bool
                d_109_v43_: int
                out6_: bool
                out7_: bool
                out8_: int
                out6_, out7_, out8_ = default__.m0(((d_106_v40_)[d_65_v3_] if (d_65_v3_) in (d_106_v40_) else d_65_v3_), d_67_globalState_)
                d_107_v41_ = out6_
                d_108_v42_ = out7_
                d_109_v43_ = out8_
                d_107_v41_ = d_62_v0_
                d_110_v44_: _dafny.MultiSet
                d_110_v44_ = _dafny.MultiSet([d_107_v41_])
                d_111_v45_: _dafny.Map
                d_111_v45_ = _dafny.Map({len(d_104_v39_): (d_110_v44_).set(d_107_v41_, default__.abs(d_65_v3_))})
                d_111_v45_ = (d_111_v45_).set(d_65_v3_, d_110_v44_)
                d_112_v46_: _dafny.MultiSet
                d_112_v46_ = _dafny.MultiSet([d_109_v43_, d_65_v3_, d_109_v43_, d_65_v3_])
                d_113_v47_: C3
                nw17_ = C3()
                nw17_.ctor__(d_112_v46_, d_107_v41_)
                d_113_v47_ = nw17_
                d_114_v48_: _dafny.Seq
                d_114_v48_ = _dafny.SeqWithoutIsStrInference([d_113_v47_])
                d_115_v49_: _dafny.Set
                d_115_v49_ = _dafny.Set({d_106_v40_})
                d_116_v50_: _dafny.Array
                nw18_ = _dafny.Array(None, 6)
                nw18_[int(0)] = d_113_v47_
                nw18_[int(1)] = d_113_v47_
                nw18_[int(2)] = (d_114_v48_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(d_115_v49_), d_109_v43_])), len(d_114_v48_))]
                nw18_[int(3)] = d_113_v47_
                nw18_[int(4)] = d_113_v47_
                nw18_[int(5)] = d_113_v47_
                d_116_v50_ = nw18_
                d_116_v50_ = d_116_v50_
            elif True:
                d_117_v51_: _dafny.Array
                nw19_ = _dafny.Array(int(0), 7)
                d_117_v51_ = nw19_
                d_118_v52_: _dafny.Array
                nw20_ = _dafny.Array(None, 10)
                nw20_[int(0)] = d_62_v0_
                nw20_[int(1)] = not(d_62_v0_)
                nw20_[int(2)] = d_62_v0_
                nw20_[int(3)] = default__.fm21(d_65_v3_, d_103_v38_, d_62_v0_, d_62_v0_, d_67_globalState_)
                nw20_[int(4)] = d_62_v0_
                nw20_[int(5)] = d_62_v0_
                nw20_[int(6)] = default__.fm23(d_104_v39_, d_65_v3_, d_65_v3_, d_62_v0_, d_67_globalState_)
                nw20_[int(7)] = False
                nw20_[int(8)] = d_62_v0_
                nw20_[int(9)] = d_62_v0_
                d_118_v52_ = nw20_
                d_119_v53_: _dafny.Set
                d_119_v53_ = _dafny.Set({d_118_v52_, d_118_v52_, d_118_v52_, d_118_v52_, d_118_v52_})
                index4_ = default__.safeIndex(699, (d_117_v51_).length(0))
                (d_117_v51_)[index4_] = len(d_119_v53_)
                d_120_v54_: _dafny.Seq
                d_120_v54_ = _dafny.SeqWithoutIsStrInference([default__.fm25(d_62_v0_, d_65_v3_, d_67_globalState_)])
                d_121_v55_: _dafny.Map
                d_121_v55_ = _dafny.Map({d_62_v0_: d_65_v3_})
                d_122_v56_: _dafny.Seq
                d_122_v56_ = _dafny.SeqWithoutIsStrInference([d_121_v55_])
                d_123_v57_: D4
                d_123_v57_ = D4_DC9(len(d_122_v56_), d_103_v38_, _dafny.SeqWithoutIsStrInference([d_65_v3_]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muudfjey"))), default__.fm23(d_104_v39_, d_65_v3_, (_dafny.MultiSet([default__.fm17(d_120_v54_, d_65_v3_, -758, d_65_v3_, d_67_globalState_)])).cardinality, False, d_67_globalState_))
                index5_ = default__.safeIndex(699, (d_117_v51_).length(0))
                (d_117_v51_)[index5_] = default__.fm17(d_120_v54_, len(d_104_v39_), (d_123_v57_).cf19, d_65_v3_, d_67_globalState_)
                index6_ = default__.safeIndex(689, (d_118_v52_).length(0))
                (d_118_v52_)[index6_] = d_62_v0_
                index7_ = default__.safeIndex(689, (d_118_v52_).length(0))
                (d_118_v52_)[index7_] = d_62_v0_
                d_121_v55_ = (d_121_v55_).set(d_62_v0_, (d_117_v51_)[default__.safeIndex(699, (d_117_v51_).length(0))])
                d_124_v58_: _dafny.Map
                d_124_v58_ = _dafny.Map({d_103_v38_: (len(d_104_v39_)) + (d_65_v3_)})
                d_124_v58_ = (d_124_v58_).set(d_103_v38_, (d_117_v51_)[default__.safeIndex(699, (d_117_v51_).length(0))])
                d_125_v59_: _dafny.Array
                def lambda1_(d_126_v3_, d_127_v0_):
                    def lambda2_(d_128_i3_):
                        return _dafny.Map({d_126_v3_: d_127_v0_})

                    return lambda2_

                init0_ = lambda1_(d_65_v3_, d_62_v0_)
                nw21_ = _dafny.Array(None, 1)
                for i0_0_ in range(nw21_.length(0)):
                    nw21_[i0_0_] = init0_(i0_0_)
                d_125_v59_ = nw21_
                d_129_v60_: C7
                nw22_ = C7()
                nw22_.ctor__(d_125_v59_)
                d_129_v60_ = nw22_
                d_129_v60_ = d_129_v60_
        hi0_ = (d_65_v3_) + (d_65_v3_)
        for d_130_i4_ in range(d_65_v3_, hi0_):
            d_131_v61_: str
            d_131_v61_ = _dafny.CodePoint('k')
            d_132_v62_: _dafny.MultiSet
            d_132_v62_ = _dafny.MultiSet([d_131_v61_, default__.fm5(d_65_v3_, d_67_globalState_)])
            d_133_v63_: D1
            d_133_v63_ = D1_DC2((d_132_v62_).cardinality, d_62_v0_, 502)
            d_134_v64_: _dafny.Set
            d_134_v64_ = _dafny.Set({False, False})
            d_135_v65_: _dafny.Seq
            d_135_v65_ = _dafny.SeqWithoutIsStrInference([d_134_v64_, d_134_v64_])
            d_136_v66_: _dafny.Array
            nw23_ = _dafny.Array(None, 1)
            nw23_[int(0)] = len((d_135_v65_)[default__.safeIndex((0) - ((0) - (d_65_v3_)), len(d_135_v65_))])
            d_136_v66_ = nw23_
            d_137_v67_: _dafny.Seq
            d_137_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qm"))
            d_138_v68_: _dafny.Map
            d_138_v68_ = _dafny.Map({d_136_v66_: len(d_137_v67_)})
            d_139_v69_: _dafny.Array
            nw24_ = _dafny.Array(None, 14)
            nw24_[int(0)] = d_62_v0_
            nw24_[int(1)] = (d_62_v0_ if d_62_v0_ else d_62_v0_)
            nw24_[int(2)] = d_62_v0_
            nw24_[int(3)] = (d_133_v63_).cf3
            nw24_[int(4)] = d_62_v0_
            nw24_[int(5)] = d_62_v0_
            nw24_[int(6)] = d_62_v0_
            nw24_[int(7)] = d_62_v0_
            nw24_[int(8)] = d_62_v0_
            nw24_[int(9)] = d_62_v0_
            nw24_[int(10)] = False
            nw24_[int(11)] = (d_136_v66_) not in (d_138_v68_)
            nw24_[int(12)] = True
            nw24_[int(13)] = d_62_v0_
            d_139_v69_ = nw24_
            index8_ = default__.safeIndex(422, (d_139_v69_).length(0))
            (d_139_v69_)[index8_] = d_62_v0_
            index9_ = default__.safeIndex(422, (d_139_v69_).length(0))
            (d_139_v69_)[index9_] = d_62_v0_
            d_62_v0_ = not(d_62_v0_)
            d_140_v70_: _dafny.MultiSet
            d_140_v70_ = _dafny.MultiSet([len(d_137_v67_)])
            d_141_v71_: C3
            nw25_ = C3()
            nw25_.ctor__(d_140_v70_, True)
            d_141_v71_ = nw25_
            d_142_v72_: _dafny.Map
            d_142_v72_ = _dafny.Map({d_62_v0_: (d_141_v71_).f10})
            index10_ = default__.safeIndex(422, (d_139_v69_).length(0))
            (d_139_v69_)[index10_] = ((d_142_v72_)[d_62_v0_] if (d_62_v0_) in (d_142_v72_) else (d_139_v69_)[default__.safeIndex(422, (d_139_v69_).length(0))])
        d_62_v0_ = d_62_v0_
        d_143_v73_: _dafny.Array
        nw26_ = _dafny.Array(False, 29)
        d_143_v73_ = nw26_
        d_143_v73_ = d_143_v73_
        d_144_v74_: _dafny.MultiSet
        d_144_v74_ = _dafny.MultiSet([d_62_v0_])
        d_145_i5_: int
        d_145_i5_ = 0
        with _dafny.label("0"):
            while ((d_144_v74_ if d_62_v0_ else d_144_v74_)) != ((d_144_v74_) | (d_144_v74_)):
                with _dafny.c_label("0"):
                    if (d_145_i5_) >= (100):
                        raise _dafny.Break("0")
                    d_145_i5_ = (d_145_i5_) + (1)
                    d_146_v75_: _dafny.Seq
                    d_146_v75_ = _dafny.SeqWithoutIsStrInference([d_68_v5_])
                    d_147_v76_: _dafny.Map
                    d_147_v76_ = _dafny.Map({default__.safeModuloInt(len(d_146_v75_), len(d_63_v1_)): _dafny.SeqWithoutIsStrInference([d_65_v3_, d_65_v3_])})
                    d_147_v76_ = d_147_v76_
                    d_148_v77_: _dafny.Map
                    d_148_v77_ = _dafny.Map({d_62_v0_: d_65_v3_})
                    d_149_v78_: _dafny.Seq
                    d_149_v78_ = _dafny.SeqWithoutIsStrInference([742])
                    d_150_v79_: _dafny.Map
                    d_150_v79_ = _dafny.Map({len(d_149_v78_): d_62_v0_})
                    d_151_v80_: _dafny.Map
                    d_151_v80_ = d_150_v79_
                    d_152_v81_: _dafny.MultiSet
                    d_152_v81_ = _dafny.MultiSet([d_151_v80_, d_151_v80_])
                    d_153_v82_: _dafny.Map
                    d_153_v82_ = _dafny.Map({d_149_v78_: d_152_v81_})
                    d_154_v83_: C6
                    nw27_ = C6()
                    nw27_.ctor__(d_148_v77_, len((d_153_v82_) | (d_153_v82_)))
                    d_154_v83_ = nw27_
                    d_155_v84_: str
                    d_155_v84_ = _dafny.CodePoint('m')
                    d_155_v84_ = d_155_v84_
                    d_63_v1_ = d_63_v1_
                    pass
            pass
        d_156_v85_: str
        d_156_v85_ = _dafny.CodePoint('y')
        d_157_v86_: C5
        nw28_ = C5()
        nw28_.ctor__(d_156_v85_)
        d_157_v86_ = nw28_
        d_158_i6_: int
        d_158_i6_ = 0
        with _dafny.label("1"):
            while not(d_62_v0_):
                with _dafny.c_label("1"):
                    if (d_158_i6_) >= (100):
                        raise _dafny.Break("1")
                    d_158_i6_ = (d_158_i6_) + (1)
                    index11_ = default__.safeIndex(724, (d_143_v73_).length(0))
                    (d_143_v73_)[index11_] = d_62_v0_
                    index12_ = default__.safeIndex(724, (d_143_v73_).length(0))
                    (d_143_v73_)[index12_] = d_62_v0_
                    (d_67_globalState_).f2 = d_65_v3_
                    d_159_v87_: _dafny.Array
                    def lambda3_(d_160_v3_, d_161_v0_):
                        def lambda4_(d_162_i7_):
                            return _dafny.Map({d_160_v3_: d_161_v0_})

                        return lambda4_

                    init1_ = lambda3_(d_65_v3_, d_62_v0_)
                    nw29_ = _dafny.Array(None, 11)
                    for i0_1_ in range(nw29_.length(0)):
                        nw29_[i0_1_] = init1_(i0_1_)
                    d_159_v87_ = nw29_
                    d_163_v88_: T0
                    nw30_ = C2()
                    nw30_.ctor__(d_62_v0_, (0) - ((0) - (d_65_v3_)), d_159_v87_)
                    d_163_v88_ = nw30_
                    d_164_v89_: _dafny.Array
                    def lambda5_(d_165_v3_):
                        def lambda6_(d_166_i8_):
                            return (d_166_i8_) * (d_165_v3_)

                        return lambda6_

                    init2_ = lambda5_(d_65_v3_)
                    nw31_ = _dafny.Array(None, 10)
                    for i0_2_ in range(nw31_.length(0)):
                        nw31_[i0_2_] = init2_(i0_2_)
                    d_164_v89_ = nw31_
                    d_167_v90_: _dafny.Map
                    d_167_v90_ = _dafny.Map({d_65_v3_: d_164_v89_})
                    d_168_v91_: D2
                    d_168_v91_ = D2_DC5(len(d_167_v90_), (d_143_v73_)[default__.safeIndex(724, (d_143_v73_).length(0))], d_62_v0_, not(d_62_v0_), (d_143_v73_)[default__.safeIndex(724, (d_143_v73_).length(0))])
                    d_169_v92_: _dafny.Seq
                    d_169_v92_ = _dafny.SeqWithoutIsStrInference([d_168_v91_])
                    d_170_v93_: D9
                    d_170_v93_ = D9_DC19(d_163_v88_, (_dafny.MultiSet([d_169_v92_])).cardinality, d_65_v3_)
                    d_171_v94_: _dafny.Array
                    nw32_ = _dafny.Array(None, 5)
                    nw32_[int(0)] = d_163_v88_
                    nw32_[int(1)] = d_163_v88_
                    nw32_[int(2)] = (d_170_v93_).cf34
                    nw32_[int(3)] = d_163_v88_
                    nw32_[int(4)] = d_163_v88_
                    d_171_v94_ = nw32_
                    d_172_v95_: _dafny.Seq
                    d_172_v95_ = _dafny.SeqWithoutIsStrInference([d_163_v88_, d_163_v88_])
                    index13_ = default__.safeIndex(160, (d_171_v94_).length(0))
                    (d_171_v94_)[index13_] = (d_172_v95_)[default__.safeIndex(d_65_v3_, len(d_172_v95_))]
                    index14_ = default__.safeIndex(160, (d_171_v94_).length(0))
                    (d_171_v94_)[index14_] = (d_170_v93_).cf34
                    d_173_v96_: _dafny.Seq
                    d_173_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfj"))
                    d_173_v96_ = (d_173_v96_) + ((d_173_v96_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wye"))))
                    pass
            pass
        d_174_v97_: _dafny.Set
        d_174_v97_ = _dafny.Set({d_62_v0_, d_62_v0_, d_62_v0_, False, d_62_v0_})
        d_175_v98_: _dafny.Set
        d_175_v98_ = d_174_v97_
        d_176_v99_: _dafny.Map
        d_176_v99_ = _dafny.Map({d_62_v0_: d_175_v98_})
        rhs14_ = (_dafny.Map({d_62_v0_: d_175_v98_})) | (_dafny.Map({d_62_v0_: d_175_v98_}))
        rhs15_ = d_65_v3_
        lhs5_ = d_67_globalState_
        d_176_v99_ = rhs14_
        lhs5_.f2 = rhs15_
        d_177_v100_: _dafny.Map
        d_177_v100_ = _dafny.Map({d_65_v3_: d_62_v0_})
        d_178_v101_: _dafny.Seq
        d_178_v101_ = _dafny.SeqWithoutIsStrInference([d_156_v85_])
        d_179_v102_: _dafny.MultiSet
        d_179_v102_ = _dafny.MultiSet([d_65_v3_, len(d_178_v101_), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_180_i9_ in range(default__.abs(652))])), d_65_v3_])
        d_181_v103_: _dafny.Seq
        d_181_v103_ = _dafny.SeqWithoutIsStrInference([(d_179_v102_).cardinality, len(_dafny.Set({len(_dafny.Map({d_156_v85_: d_179_v102_}))}))])
        d_182_v106_: _dafny.Map
        d_182_v106_ = d_177_v100_
        d_183_v107_: _dafny.Map
        d_183_v107_ = _dafny.Map({d_62_v0_: d_65_v3_})
        d_184_v108_: C6
        nw33_ = C6()
        nw33_.ctor__(d_183_v107_, d_65_v3_)
        d_184_v108_ = nw33_
        d_185_v109_: C0
        nw34_ = C0()
        nw34_.ctor__(default__.fm9(d_65_v3_, d_67_globalState_))
        d_185_v109_ = nw34_
        d_186_v110_: _dafny.Seq
        d_186_v110_ = _dafny.SeqWithoutIsStrInference([d_185_v109_, d_185_v109_])
        d_187_v111_: _dafny.Map
        d_187_v111_ = _dafny.Map({d_184_v108_: (0) - (len(d_186_v110_))})
        d_188_v112_: _dafny.Array
        nw35_ = _dafny.Array(None, 29)
        nw35_[int(0)] = d_177_v100_
        nw35_[int(1)] = d_177_v100_
        nw35_[int(2)] = d_177_v100_
        nw35_[int(3)] = d_177_v100_
        nw35_[int(4)] = (_dafny.Map({d_65_v3_: d_62_v0_})).set(147, d_62_v0_)
        nw35_[int(5)] = d_177_v100_
        nw35_[int(6)] = (_dafny.Map({d_65_v3_: d_62_v0_})).set(d_65_v3_, d_62_v0_)
        nw35_[int(7)] = default__.fm34(d_62_v0_, (d_157_v86_).f6, d_181_v103_, d_67_globalState_)
        nw35_[int(8)] = _dafny.Map({len(d_178_v101_): d_62_v0_})
        nw35_[int(9)] = d_177_v100_
        nw35_[int(10)] = d_177_v100_
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in (d_179_v102_).Elements:
                d_189_v104_: int = compr_12_
                if (d_189_v104_) in (d_179_v102_):
                    coll12_[(d_189_v104_) + (d_65_v3_)] = d_62_v0_
            return _dafny.Map(coll12_)
        nw35_[int(11)] = iife12_()
        
        nw35_[int(12)] = d_177_v100_
        nw35_[int(13)] = d_177_v100_
        nw35_[int(14)] = (d_177_v100_).set(d_65_v3_, d_62_v0_)
        nw35_[int(15)] = d_177_v100_
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in _dafny.IntegerRange(80, 355):
                d_190_v105_: int = compr_13_
                if ((80) <= (d_190_v105_)) and ((d_190_v105_) < (355)):
                    coll13_[(d_190_v105_) + (d_65_v3_)] = d_62_v0_
            return _dafny.Map(coll13_)
        nw35_[int(16)] = iife13_()
        
        nw35_[int(17)] = default__.fm34(d_62_v0_, d_156_v85_, d_181_v103_, d_67_globalState_)
        nw35_[int(18)] = (d_182_v106_)
        nw35_[int(19)] = d_177_v100_
        nw35_[int(20)] = _dafny.Map({d_65_v3_: d_62_v0_})
        nw35_[int(21)] = _dafny.Map({((d_187_v111_)[d_184_v108_] if (d_184_v108_) in (d_187_v111_) else 684): d_62_v0_})
        nw35_[int(22)] = d_177_v100_
        nw35_[int(23)] = d_177_v100_
        nw35_[int(24)] = d_177_v100_
        nw35_[int(25)] = _dafny.Map({(_dafny.MultiSet([d_65_v3_])).cardinality: d_62_v0_})
        nw35_[int(26)] = d_177_v100_
        nw35_[int(27)] = d_177_v100_
        nw35_[int(28)] = d_177_v100_
        d_188_v112_ = nw35_
        d_191_v113_: C7
        nw36_ = C7()
        nw36_.ctor__(d_188_v112_)
        d_191_v113_ = nw36_
        d_192_v114_: T1
        nw37_ = C7()
        nw37_.ctor__(d_188_v112_)
        d_192_v114_ = nw37_
        d_193_v115_: _dafny.Seq
        d_193_v115_ = _dafny.SeqWithoutIsStrInference([d_192_v114_, d_192_v114_])
        d_194_v116_: C2
        nw38_ = C2()
        nw38_.ctor__(d_62_v0_, (d_181_v103_)[default__.safeIndex(len(d_193_v115_), len(d_181_v103_))], d_188_v112_)
        d_194_v116_ = nw38_
        d_195_v117_: int
        d_196_v118_: str
        out9_: int
        out10_: str
        out9_, out10_ = (d_191_v113_).m1(_dafny.SeqWithoutIsStrInference([d_65_v3_ for d_197_i10_ in range(default__.abs(119))]), ((d_185_v109_).f13) == ((d_194_v116_).f12), d_178_v101_, d_67_globalState_)
        d_195_v117_ = out9_
        d_196_v118_ = out10_
        index15_ = default__.safeIndex(902, (d_143_v73_).length(0))
        (d_143_v73_)[index15_] = (_dafny.Set({(d_185_v109_).f13})).issubset(d_68_v5_)
        index16_ = default__.safeIndex(902, (d_143_v73_).length(0))
        (d_143_v73_)[index16_] = default__.fm21(default__.safeDivisionInt(71, d_195_v117_), d_196_v118_, not ((d_194_v116_).f11) or (d_62_v0_), d_62_v0_, d_67_globalState_)
        d_62_v0_ = not(((d_194_v116_).f11) == ((d_144_v74_).issubset(d_144_v74_)))
        d_198_v119_: D8
        d_198_v119_ = D8_DC15(d_143_v73_)
        d_199_v120_: _dafny.MultiSet
        d_199_v120_ = _dafny.MultiSet([d_198_v119_])
        d_200_v121_: _dafny.Seq
        d_200_v121_ = _dafny.SeqWithoutIsStrInference([d_179_v102_, d_179_v102_])
        d_201_v122_: _dafny.Map
        d_201_v122_ = _dafny.Map({(d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))]: d_62_v0_})
        d_202_v123_: _dafny.MultiSet
        d_202_v123_ = _dafny.MultiSet([d_181_v103_])
        d_203_v124_: _dafny.Map
        d_203_v124_ = _dafny.Map({d_178_v101_: (0) - ((d_144_v74_).cardinality)})
        d_204_v125_: _dafny.Array
        nw39_ = _dafny.Array(None, 9)
        nw39_[int(0)] = ((d_199_v120_ if (d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))] else d_199_v120_)).cardinality
        nw39_[int(1)] = len(((d_200_v121_).set(default__.safeIndex((d_194_v116_).f12, len(d_200_v121_)), _dafny.MultiSet([d_195_v117_, (d_194_v116_).f12]))) + (d_200_v121_))
        nw39_[int(2)] = ((d_194_v116_).f12 if d_62_v0_ else ((d_179_v102_).set(len(d_201_v122_), default__.abs(d_195_v117_))).cardinality)
        nw39_[int(3)] = ((d_202_v123_)[_dafny.SeqWithoutIsStrInference([d_65_v3_])] if (_dafny.SeqWithoutIsStrInference([d_65_v3_])) in (d_202_v123_) else d_65_v3_)
        nw39_[int(4)] = ((d_203_v124_)[d_178_v101_] if (d_178_v101_) in (d_203_v124_) else d_184_v108_.f5)
        nw39_[int(5)] = (d_65_v3_) * ((d_181_v103_)[default__.safeIndex(len(d_178_v101_), len(d_181_v103_))])
        nw39_[int(6)] = (d_185_v109_).f13
        nw39_[int(7)] = 689
        nw39_[int(8)] = (d_185_v109_).f13
        d_204_v125_ = nw39_
        index17_ = default__.safeIndex(288, (d_204_v125_).length(0))
        (d_204_v125_)[index17_] = (d_185_v109_).f13
        d_205_v126_: _dafny.Map
        d_205_v126_ = _dafny.Map({d_204_v125_: (d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))]})
        index18_ = default__.safeIndex(288, (d_204_v125_).length(0))
        (d_204_v125_)[index18_] = len(d_205_v126_)
        source3_ = D3_DC7(len(d_63_v1_), default__.safeDivisionInt(d_184_v108_.f5, (d_204_v125_)[default__.safeIndex(288, (d_204_v125_).length(0))]))
        if source3_.is_DC7:
            d_206___mcc_h0_ = source3_.cf13
            d_207___mcc_h1_ = source3_.cf14
            d_208_cf14_ = d_207___mcc_h1_
            d_209_cf13_ = d_206___mcc_h0_
            d_210_v127_: int
            d_211_v128_: str
            out11_: int
            out12_: str
            out11_, out12_ = (d_191_v113_).m1(default__.fm12(not((d_194_v116_).f11), d_184_v108_.f5, d_174_v97_, _dafny.Map({(d_194_v116_).f12: len(d_68_v5_)}), d_67_globalState_), (d_194_v116_).f11, d_178_v101_, d_67_globalState_)
            d_210_v127_ = out11_
            d_211_v128_ = out12_
            index19_ = default__.safeIndex(902, (d_143_v73_).length(0))
            (d_143_v73_)[index19_] = (d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))]
            d_212_v130_: _dafny.Map
            d_212_v130_ = _dafny.Map({d_208_cf14_: d_209_cf13_})
            d_213_v131_: _dafny.Map
            def iife14_():
                coll14_ = _dafny.Set()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(676, 163):
                    d_214_v129_: int = compr_14_
                    if ((676) <= (d_214_v129_)) and ((d_214_v129_) < (163)):
                        coll14_ = coll14_.union(_dafny.Set([default__.safeModuloInt(d_214_v129_, len(d_174_v97_))]))
                return _dafny.Set(coll14_)
            d_213_v131_ = _dafny.Map({(default__.fm24(iife14_()
            , d_67_globalState_)) * (len(d_212_v130_)): (d_185_v109_).f13})
            d_212_v130_ = d_213_v131_
            d_215_v132_: _dafny.Array
            nw40_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
            d_215_v132_ = nw40_
            d_216_v133_: D9
            d_216_v133_ = D9_DC18(d_215_v132_)
            index20_ = default__.safeIndex(288, (d_204_v125_).length(0))
            rhs16_ = d_216_v133_
            rhs17_ = default__.safeModuloInt(len(d_178_v101_), ((d_194_v116_).f12) + (d_208_cf14_))
            rhs18_ = d_143_v73_
            rhs19_ = ((332) + (len(d_178_v101_))) + (d_184_v108_.f5)
            lhs6_ = d_204_v125_
            lhs7_ = default__.safeIndex(288, (d_204_v125_).length(0))
            d_216_v133_ = rhs16_
            d_209_cf13_ = rhs17_
            d_143_v73_ = rhs18_
            lhs6_[lhs7_] = rhs19_
        elif True:
            d_217___mcc_h2_ = source3_.cf12
            d_218_cf12_ = d_217___mcc_h2_
            d_179_v102_ = (d_179_v102_).intersection((d_179_v102_) | (d_179_v102_))
            d_68_v5_ = (d_68_v5_ if (d_194_v116_).f11 else d_68_v5_)
            d_65_v3_ = (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_194_v116_])), d_184_v108_.f5) if d_62_v0_ else (d_194_v116_).f12)
            d_62_v0_ = True
        d_219_v134_: D4
        d_219_v134_ = D4_DC8(d_185_v109_)
        source4_ = d_219_v134_
        if source4_.is_DC9:
            d_220___mcc_h3_ = source4_.cf16
            d_221___mcc_h4_ = source4_.cf17
            d_222___mcc_h5_ = source4_.cf18
            d_223___mcc_h6_ = source4_.cf19
            d_224___mcc_h7_ = source4_.cf20
            d_225_cf20_ = d_224___mcc_h7_
            d_226_cf19_ = d_223___mcc_h6_
            d_227_cf18_ = d_222___mcc_h5_
            d_228_cf17_ = d_221___mcc_h4_
            d_229_cf16_ = d_220___mcc_h3_
            d_62_v0_ = d_225_cf20_
            d_195_v117_ = (d_191_v113_).fm0((d_185_v109_).f13, d_67_globalState_)
            d_230_v135_: int
            d_231_v136_: str
            out13_: int
            out14_: str
            out13_, out14_ = (d_194_v116_).m1(_dafny.SeqWithoutIsStrInference([len(d_63_v1_)]), default__.fm23(d_178_v101_, (d_179_v102_).cardinality, (d_194_v116_).f12, d_225_cf20_, d_67_globalState_), default__.fm25(True, d_184_v108_.f5, d_67_globalState_), d_67_globalState_)
            d_230_v135_ = out13_
            d_231_v136_ = out14_
            if d_62_v0_:
                d_183_v107_ = (d_183_v107_).set((d_194_v116_).f11, d_65_v3_)
                d_232_v137_: _dafny.Array
                def lambda7_(d_233_cf18_):
                    def lambda8_(d_234_i11_):
                        return d_233_cf18_

                    return lambda8_

                init3_ = lambda7_(d_227_cf18_)
                nw41_ = _dafny.Array(None, 16)
                for i0_3_ in range(nw41_.length(0)):
                    nw41_[i0_3_] = init3_(i0_3_)
                d_232_v137_ = nw41_
                index21_ = default__.safeIndex(642, (d_232_v137_).length(0))
                (d_232_v137_)[index21_] = d_181_v103_
                index22_ = default__.safeIndex(642, (d_232_v137_).length(0))
                index23_ = default__.safeIndex(902, (d_143_v73_).length(0))
                rhs20_ = (((d_184_v108_).f4) | (d_183_v107_)) | ((d_184_v108_).f4)
                rhs21_ = d_195_v117_
                rhs22_ = (d_185_v109_).f13
                rhs23_ = d_227_cf18_
                rhs24_ = not(((_dafny.MultiSet(d_63_v1_)).set((d_194_v116_).f11, default__.abs(len(((d_201_v122_).set(False, (d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))])).set(d_62_v0_, (d_194_v116_).f11))))) != (_dafny.MultiSet(d_63_v1_)))
                lhs8_ = d_67_globalState_
                lhs9_ = d_232_v137_
                lhs10_ = default__.safeIndex(642, (d_232_v137_).length(0))
                lhs11_ = d_143_v73_
                lhs12_ = default__.safeIndex(902, (d_143_v73_).length(0))
                d_183_v107_ = rhs20_
                lhs8_.f2 = rhs21_
                d_65_v3_ = rhs22_
                lhs9_[lhs10_] = rhs23_
                lhs11_[lhs12_] = rhs24_
                d_235_v138_: _dafny.Map
                d_235_v138_ = _dafny.Map({d_191_v113_: (d_232_v137_)[default__.safeIndex(642, (d_232_v137_).length(0))]})
                d_236_v139_: _dafny.Seq
                d_236_v139_ = _dafny.SeqWithoutIsStrInference([(d_232_v137_)[default__.safeIndex(642, (d_232_v137_).length(0))], (d_232_v137_)[default__.safeIndex(642, (d_232_v137_).length(0))]])
                d_235_v138_ = ((_dafny.Map({d_191_v113_: d_181_v103_})).set(d_191_v113_, (d_236_v139_)[default__.safeIndex((d_185_v109_).f13, len(d_236_v139_))])) | (d_235_v138_)
                d_237_v140_: D9
                d_237_v140_ = D9_DC20((d_194_v116_).f11)
                d_237_v140_ = default__.fm45((d_144_v74_).ispropersubset(_dafny.MultiSet([default__.fm21(d_230_v135_, _dafny.CodePoint('f'), (d_194_v116_).f11, not((d_194_v116_).f11), d_67_globalState_)])), (d_194_v116_).f12, d_67_globalState_)
                d_238_v141_: _dafny.Map
                d_238_v141_ = _dafny.Map({d_195_v117_: d_178_v101_})
                d_204_v125_ = (d_204_v125_ if (d_157_v86_).fm7(380, (d_194_v116_).f11, d_238_v141_, d_67_globalState_) else d_204_v125_)
            elif True:
                d_239_v142_: _dafny.Seq
                d_239_v142_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, ((d_177_v100_)[(d_194_v116_).f12] if ((d_194_v116_).f12) in (d_177_v100_) else d_225_cf20_)]), (d_63_v1_ if (d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))] else d_63_v1_), d_63_v1_])
                d_240_v143_: T0
                nw42_ = C2()
                nw42_.ctor__((d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))], d_184_v108_.f5, d_188_v112_)
                d_240_v143_ = nw42_
                d_241_v144_: _dafny.Map
                d_241_v144_ = _dafny.Map({d_240_v143_: d_239_v142_})
                d_239_v142_ = ((d_241_v144_)[d_240_v143_] if (d_240_v143_) in (d_241_v144_) else d_239_v142_)
                d_231_v136_ = (d_157_v86_).f6
                d_186_v110_ = d_186_v110_
                d_191_v113_ = (d_191_v113_ if (d_179_v102_).ispropersubset(d_179_v102_) else d_191_v113_)
                d_242_v145_: _dafny.Array
                nw43_ = _dafny.Array(None, 12)
                nw43_[int(0)] = (d_178_v101_)[default__.safeIndex((d_185_v109_).f13, len(d_178_v101_))]
                nw43_[int(1)] = d_156_v85_
                nw43_[int(2)] = (d_178_v101_)[default__.safeIndex(len(d_178_v101_), len(d_178_v101_))]
                nw43_[int(3)] = d_156_v85_
                nw43_[int(4)] = (d_157_v86_).f6
                nw43_[int(5)] = d_196_v118_
                nw43_[int(6)] = ((D3_DC6(d_196_v118_)).cf12 if not(True) else (d_157_v86_).f6)
                nw43_[int(7)] = (d_157_v86_).f6
                nw43_[int(8)] = (d_157_v86_).f6
                nw43_[int(9)] = d_228_cf17_
                nw43_[int(10)] = (d_157_v86_).f6
                nw43_[int(11)] = (d_157_v86_).f6
                d_242_v145_ = nw43_
                index24_ = default__.safeIndex(825, (d_242_v145_).length(0))
                (d_242_v145_)[index24_] = (d_157_v86_).f6
                d_243_v146_: _dafny.Map
                d_243_v146_ = _dafny.Map({d_62_v0_: d_174_v97_})
                d_244_v147_: _dafny.Map
                d_244_v147_ = (_dafny.Map({len(d_178_v101_): len(((d_243_v146_)[d_225_cf20_] if (d_225_cf20_) in (d_243_v146_) else d_174_v97_))})).set(d_184_v108_.f5, d_65_v3_)
                index25_ = default__.safeIndex(131, (d_242_v145_).length(0))
                (d_242_v145_)[index25_] = (d_231_v136_ if d_225_cf20_ else (d_157_v86_).f6)
                index26_ = default__.safeIndex(825, (d_242_v145_).length(0))
                index27_ = default__.safeIndex(131, (d_242_v145_).length(0))
                rhs25_ = (d_194_v116_).f11
                rhs26_ = d_196_v118_
                rhs27_ = d_244_v147_
                rhs28_ = (d_178_v101_)[default__.safeIndex(default__.safeDivisionInt(-851, len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbovk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kydcrfrb"))}))), len(d_178_v101_))]
                lhs13_ = d_242_v145_
                lhs14_ = default__.safeIndex(825, (d_242_v145_).length(0))
                lhs15_ = d_242_v145_
                lhs16_ = default__.safeIndex(131, (d_242_v145_).length(0))
                d_62_v0_ = rhs25_
                lhs13_[lhs14_] = rhs26_
                d_244_v147_ = rhs27_
                lhs15_[lhs16_] = rhs28_
        elif True:
            d_245___mcc_h8_ = source4_.cf15
            d_246_cf15_ = d_245___mcc_h8_
            if ((d_246_cf15_).f13) == ((-485) - ((d_185_v109_).f13)):
                d_200_v121_ = d_200_v121_
                (d_67_globalState_).f2 = (0) - (default__.safeDivisionInt(default__.fm9(836, d_67_globalState_), (d_204_v125_)[default__.safeIndex(288, (d_204_v125_).length(0))]))
                (d_184_v108_).f5 = ((d_183_v107_)[((d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))] if True else default__.fm2(_dafny.CodePoint('r'), _dafny.SeqWithoutIsStrInference([d_201_v122_]), False, d_67_globalState_))] if (((d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))] if True else default__.fm2(_dafny.CodePoint('r'), _dafny.SeqWithoutIsStrInference([d_201_v122_]), False, d_67_globalState_))) in (d_183_v107_) else (d_195_v117_) + ((d_246_cf15_).f13))
                (d_67_globalState_).f2 = d_195_v117_
                d_65_v3_ = default__.fm9(((d_246_cf15_).f13) * ((d_194_v116_).f12), d_67_globalState_)
            elif True:
                index28_ = default__.safeIndex(902, (d_143_v73_).length(0))
                (d_143_v73_)[index28_] = True
                d_68_v5_ = (d_68_v5_) - (d_68_v5_)
                d_247_v148_: _dafny.Array
                nw44_ = _dafny.Array(None, 7)
                d_247_v148_ = nw44_
                d_247_v148_ = d_247_v148_
                d_248_v149_: _dafny.Seq
                d_248_v149_ = _dafny.SeqWithoutIsStrInference([d_194_v116_, d_194_v116_])
                d_249_v150_: D13
                d_249_v150_ = D13_DC24(d_194_v116_)
                d_250_v151_: _dafny.Array
                nw45_ = _dafny.Array(None, 23)
                nw45_[int(0)] = d_194_v116_
                nw45_[int(1)] = d_194_v116_
                nw45_[int(2)] = d_194_v116_
                nw45_[int(3)] = (d_248_v149_)[default__.safeIndex((d_194_v116_).f12, len(d_248_v149_))]
                nw45_[int(4)] = d_194_v116_
                nw45_[int(5)] = d_194_v116_
                nw45_[int(6)] = d_194_v116_
                nw45_[int(7)] = d_194_v116_
                nw45_[int(8)] = (d_194_v116_ if (d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))] else d_194_v116_)
                nw45_[int(9)] = d_194_v116_
                nw45_[int(10)] = d_194_v116_
                nw45_[int(11)] = d_194_v116_
                nw45_[int(12)] = d_194_v116_
                nw45_[int(13)] = d_194_v116_
                nw45_[int(14)] = d_194_v116_
                nw45_[int(15)] = (d_194_v116_ if (d_194_v116_).f11 else (d_249_v150_).cf41)
                nw45_[int(16)] = d_194_v116_
                nw45_[int(17)] = (d_248_v149_)[default__.safeIndex((d_204_v125_)[default__.safeIndex(288, (d_204_v125_).length(0))], len(d_248_v149_))]
                nw45_[int(18)] = d_194_v116_
                nw45_[int(19)] = d_194_v116_
                nw45_[int(20)] = d_194_v116_
                nw45_[int(21)] = d_194_v116_
                nw45_[int(22)] = d_194_v116_
                d_250_v151_ = nw45_
                index29_ = default__.safeIndex(385, (d_250_v151_).length(0))
                (d_250_v151_)[index29_] = d_194_v116_
                index30_ = default__.safeIndex(385, (d_250_v151_).length(0))
                (d_250_v151_)[index30_] = d_194_v116_
                d_251_v152_: C7
                nw46_ = C7()
                nw46_.ctor__(d_188_v112_)
                d_251_v152_ = nw46_
            d_252_v153_: C4
            nw47_ = C4()
            nw47_.ctor__((d_143_v73_)[default__.safeIndex(902, (d_143_v73_).length(0))], d_179_v102_, d_188_v112_)
            d_252_v153_ = nw47_
            d_253_v154_: _dafny.Map
            d_253_v154_ = _dafny.Map({d_252_v153_: d_62_v0_})
            d_178_v101_ = (((d_178_v101_) + (d_178_v101_)).set(default__.safeIndex(len(d_253_v154_), len((d_178_v101_) + (d_178_v101_))), (d_157_v86_).f6)) + (d_178_v101_)
            d_254_v155_: _dafny.Seq
            d_254_v155_ = _dafny.SeqWithoutIsStrInference([d_178_v101_])
            d_255_v156_: _dafny.Map
            d_255_v156_ = _dafny.Map({d_195_v117_: default__.fm17(d_254_v155_, default__.fm9((d_185_v109_).f13, d_67_globalState_), (d_194_v116_).f12, (d_194_v116_).f12, d_67_globalState_)})
            index31_ = default__.safeIndex(288, (d_204_v125_).length(0))
            (d_204_v125_)[index31_] = ((d_255_v156_)[(d_185_v109_).f13] if ((d_185_v109_).f13) in (d_255_v156_) else -30)
            d_256_v157_: int
            d_257_v158_: bool
            out15_: int
            out16_: bool
            out15_, out16_ = (d_184_v108_).m2((d_194_v116_).f11, d_67_globalState_)
            d_256_v157_ = out15_
            d_257_v158_ = out16_
        _dafny.print(_dafny.string_of(d_62_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v1_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_64_v2_)) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_65_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[0]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[1]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[2]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[3]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[4]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[5]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[6]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[7]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[8]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[9]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_66_v4_)[10]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[0]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[1]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[2]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[3]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[4]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[5]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[6]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[7]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[8]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[9]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_67_globalState_).f0)[10]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_67_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_67_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_68_v5_) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v73_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v73_)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_144_v74_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_156_v85_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v86_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_158_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v97_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v98_)) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v99_) == (_dafny.Map({False: _dafny.Set({False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_v100_) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_v101_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('y'), _dafny.CodePoint('y')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v102_) == (_dafny.MultiSet([124, 124, 1, 652]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_v103_) == (_dafny.SeqWithoutIsStrInference([4, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_182_v106_)) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v107_) == (_dafny.Map({False: 124}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_184_v108_).f4) == (_dafny.Map({False: 124}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_184_v108_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_v109_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_186_v110_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_187_v111_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[0]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[1]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[2]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[3]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[4]) == (_dafny.Map({124: False, 147: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[5]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[6]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[7]) == (_dafny.Map({1: True, 2: False, 449: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[8]) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[9]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[10]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[11]) == (_dafny.Map({248: False, 125: False, 776: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[12]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[13]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[14]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[15]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[16]) == (_dafny.Map({204: False, 205: False, 206: False, 207: False, 208: False, 209: False, 210: False, 211: False, 212: False, 213: False, 214: False, 215: False, 216: False, 217: False, 218: False, 219: False, 220: False, 221: False, 222: False, 223: False, 224: False, 225: False, 226: False, 227: False, 228: False, 229: False, 230: False, 231: False, 232: False, 233: False, 234: False, 235: False, 236: False, 237: False, 238: False, 239: False, 240: False, 241: False, 242: False, 243: False, 244: False, 245: False, 246: False, 247: False, 248: False, 249: False, 250: False, 251: False, 252: False, 253: False, 254: False, 255: False, 256: False, 257: False, 258: False, 259: False, 260: False, 261: False, 262: False, 263: False, 264: False, 265: False, 266: False, 267: False, 268: False, 269: False, 270: False, 271: False, 272: False, 273: False, 274: False, 275: False, 276: False, 277: False, 278: False, 279: False, 280: False, 281: False, 282: False, 283: False, 284: False, 285: False, 286: False, 287: False, 288: False, 289: False, 290: False, 291: False, 292: False, 293: False, 294: False, 295: False, 296: False, 297: False, 298: False, 299: False, 300: False, 301: False, 302: False, 303: False, 304: False, 305: False, 306: False, 307: False, 308: False, 309: False, 310: False, 311: False, 312: False, 313: False, 314: False, 315: False, 316: False, 317: False, 318: False, 319: False, 320: False, 321: False, 322: False, 323: False, 324: False, 325: False, 326: False, 327: False, 328: False, 329: False, 330: False, 331: False, 332: False, 333: False, 334: False, 335: False, 336: False, 337: False, 338: False, 339: False, 340: False, 341: False, 342: False, 343: False, 344: False, 345: False, 346: False, 347: False, 348: False, 349: False, 350: False, 351: False, 352: False, 353: False, 354: False, 355: False, 356: False, 357: False, 358: False, 359: False, 360: False, 361: False, 362: False, 363: False, 364: False, 365: False, 366: False, 367: False, 368: False, 369: False, 370: False, 371: False, 372: False, 373: False, 374: False, 375: False, 376: False, 377: False, 378: False, 379: False, 380: False, 381: False, 382: False, 383: False, 384: False, 385: False, 386: False, 387: False, 388: False, 389: False, 390: False, 391: False, 392: False, 393: False, 394: False, 395: False, 396: False, 397: False, 398: False, 399: False, 400: False, 401: False, 402: False, 403: False, 404: False, 405: False, 406: False, 407: False, 408: False, 409: False, 410: False, 411: False, 412: False, 413: False, 414: False, 415: False, 416: False, 417: False, 418: False, 419: False, 420: False, 421: False, 422: False, 423: False, 424: False, 425: False, 426: False, 427: False, 428: False, 429: False, 430: False, 431: False, 432: False, 433: False, 434: False, 435: False, 436: False, 437: False, 438: False, 439: False, 440: False, 441: False, 442: False, 443: False, 444: False, 445: False, 446: False, 447: False, 448: False, 449: False, 450: False, 451: False, 452: False, 453: False, 454: False, 455: False, 456: False, 457: False, 458: False, 459: False, 460: False, 461: False, 462: False, 463: False, 464: False, 465: False, 466: False, 467: False, 468: False, 469: False, 470: False, 471: False, 472: False, 473: False, 474: False, 475: False, 476: False, 477: False, 478: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[17]) == (_dafny.Map({1: True, 2: False, 449: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[18]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[19]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[20]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[21]) == (_dafny.Map({-2: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[22]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[23]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[24]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[25]) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[26]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[27]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_188_v112_)[28]) == (_dafny.Map({124: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_193_v115_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v116_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v116_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_v117_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_196_v118_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_v119_).cf27)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_v119_).cf27)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v120_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_200_v121_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([124, 124, 1, 652]), _dafny.MultiSet([124, 124, 1, 652])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_201_v122_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_202_v123_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([4, 1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_v124_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')]): -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v125_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_205_v126_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_219_v134_).cf15).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
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
        return lambda: D1_DC2(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
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
        return lambda: D2_DC4(int(0))
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

class D2_DC4(D2, NamedTuple('DC4', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC7(D3, NamedTuple('DC7', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC9(int(0), _dafny.CodePoint('D'), _dafny.Seq({}), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)

class D4_DC9(D4, NamedTuple('DC9', [('cf16', Any), ('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)

class D5_DC10(D5, NamedTuple('DC10', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC12(False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)

class D6_DC12(D6, NamedTuple('DC12', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC11(D6, NamedTuple('DC11', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({self.cf22.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC14(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D7_DC13)

class D7_DC14(D7, NamedTuple('DC14', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC13(D7, NamedTuple('DC13', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC13({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC13) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC16(int(0), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D8_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D8_DC15)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)

class D8_DC16(D8, NamedTuple('DC16', [('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC16({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC16) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC15(D8, NamedTuple('DC15', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC15({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC15) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC17(D8, NamedTuple('DC17', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC19(None, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D9_DC18)

class D9_DC19(D9, NamedTuple('DC19', [('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC18(D9, NamedTuple('DC18', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC18({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC18) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)

class D10_DC21(D10, NamedTuple('DC21', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D11_DC22)

class D11_DC22(D11, NamedTuple('DC22', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC22({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC22) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D12_DC23)

class D12_DC23(D12, NamedTuple('DC23', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC23({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC23) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC25(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D13_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D13_DC24)

class D13_DC25(D13, NamedTuple('DC25', [('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC25({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC25) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC24(D13, NamedTuple('DC24', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC24({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC24) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC27(int(0), False, False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D14_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D14_DC28)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D14_DC26)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D14_DC29)

class D14_DC27(D14, NamedTuple('DC27', [('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC27({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC27) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC28(D14, NamedTuple('DC28', [('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC28({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC28) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC26(D14, NamedTuple('DC26', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC26({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC26) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC29(D14, NamedTuple('DC29', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC29({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC29) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m1(self, p0, p1, p2, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1

class C0:
    def  __init__(self):
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f13):
        (self)._f13 = f13

    @property
    def f13(self):
        return self._f13

class C1(T0):
    def  __init__(self):
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f14: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f14, f3):
        (self)._f14 = f14
        (self)._f3 = f3

    def fm29(self, p0, p1, p2, globalState):
        return _dafny.CodePoint('y')

    @property
    def f14(self):
        return self._f14

class C2(T0, T1):
    def  __init__(self):
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f11: bool = False
        self._f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f11, f12, f3):
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f3 = f3

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: str = _dafny.CodePoint('D')
        d_258_v0_: bool
        d_258_v0_ = True
        d_258_v0_ = (self).f11
        d_259_v1_: _dafny.Map
        d_259_v1_ = _dafny.Map({(self).f11: default__.fm23(p2, (self).f12, (self).f12, True, globalState)})
        d_260_v2_: _dafny.Seq
        d_260_v2_ = _dafny.SeqWithoutIsStrInference([d_259_v1_, d_259_v1_])
        d_261_v3_: _dafny.Map
        d_261_v3_ = _dafny.Map({(p0)[default__.safeIndex(default__.fm24(_dafny.Set({(self).f12, (self).f12, len(p2)}), globalState), len(p0))]: d_260_v2_})
        d_262_v4_: D1
        d_262_v4_ = D1_DC1((d_260_v2_) + (((d_261_v3_)[(self).f12] if ((self).f12) in (d_261_v3_) else _dafny.SeqWithoutIsStrInference([d_259_v1_]))))
        source5_ = d_262_v4_
        if source5_.is_DC2:
            d_263___mcc_h0_ = source5_.cf2
            d_264___mcc_h1_ = source5_.cf3
            d_265___mcc_h2_ = source5_.cf4
            d_266_cf4_ = d_265___mcc_h2_
            d_267_cf3_ = d_264___mcc_h1_
            d_268_cf2_ = d_263___mcc_h0_
            d_269_v5_: _dafny.Array
            nw48_ = _dafny.Array(False, 22)
            d_269_v5_ = nw48_
            index32_ = default__.safeIndex(309, (d_269_v5_).length(0))
            (d_269_v5_)[index32_] = d_258_v0_
            d_270_v6_: _dafny.Set
            d_270_v6_ = _dafny.Set({d_268_cf2_, len(p2), -72})
            index33_ = default__.safeIndex(309, (d_269_v5_).length(0))
            (d_269_v5_)[index33_] = (d_270_v6_).issubset((d_270_v6_) - (d_270_v6_))
            d_270_v6_ = d_270_v6_
            if d_267_cf3_:
                d_271_v7_: _dafny.Seq
                d_271_v7_ = _dafny.SeqWithoutIsStrInference([True, d_258_v0_])
                d_272_v8_: C0
                nw49_ = C0()
                nw49_.ctor__((d_266_cf4_) - (len(d_271_v7_)))
                d_272_v8_ = nw49_
                d_273_v9_: _dafny.Seq
                d_273_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdgtom"))
                d_273_v9_ = (d_273_v9_ if not(False) else d_273_v9_)
                d_274_v10_: _dafny.Set
                d_274_v10_ = _dafny.Set({(self).f11, p1})
                d_275_v11_: _dafny.Array
                nw50_ = _dafny.Array(None, 14)
                nw50_[int(0)] = _dafny.Set({d_258_v0_})
                nw50_[int(1)] = (d_274_v10_) | (d_274_v10_)
                nw50_[int(2)] = d_274_v10_
                nw50_[int(3)] = _dafny.Set({(d_269_v5_)[default__.safeIndex(309, (d_269_v5_).length(0))], True})
                nw50_[int(4)] = d_274_v10_
                nw50_[int(5)] = d_274_v10_
                nw50_[int(6)] = d_274_v10_
                nw50_[int(7)] = d_274_v10_
                nw50_[int(8)] = (d_274_v10_ if True else d_274_v10_)
                nw50_[int(9)] = (_dafny.Set({d_258_v0_})) - (d_274_v10_)
                nw50_[int(10)] = d_274_v10_
                nw50_[int(11)] = d_274_v10_
                nw50_[int(12)] = d_274_v10_
                nw50_[int(13)] = d_274_v10_
                d_275_v11_ = nw50_
                d_275_v11_ = d_275_v11_
                d_276_v12_: _dafny.MultiSet
                d_276_v12_ = _dafny.MultiSet([(self).f12, 634, (d_272_v8_).f13])
                d_277_v13_: _dafny.Map
                d_277_v13_ = _dafny.Map({d_276_v12_: d_266_cf4_})
                d_278_v14_: D2
                d_278_v14_ = D2_DC3(d_277_v13_)
                d_279_v15_: _dafny.Map
                d_279_v15_ = _dafny.Map({d_278_v14_: p1})
                d_280_v16_: _dafny.Map
                d_280_v16_ = _dafny.Map({d_268_cf2_: (len(d_270_v6_)) >= (len(d_279_v15_))})
                d_280_v16_ = (d_280_v16_).set((d_272_v8_).f13, (d_269_v5_)[default__.safeIndex(309, (d_269_v5_).length(0))])
                d_281_v17_: _dafny.Array
                nw51_ = _dafny.Array(int(0), 26)
                d_281_v17_ = nw51_
                d_282_v18_: _dafny.Seq
                d_282_v18_ = _dafny.SeqWithoutIsStrInference([d_270_v6_])
                d_283_v19_: _dafny.Map
                d_283_v19_ = _dafny.Map({d_281_v17_: len((d_282_v18_) + (_dafny.SeqWithoutIsStrInference([d_270_v6_, d_270_v6_, d_270_v6_])))})
                d_283_v19_ = (d_283_v19_).set(d_281_v17_, (self).f12)
            elif True:
                d_258_v0_ = not(d_267_cf3_)
                d_284_v20_: bool
                d_285_v21_: bool
                d_286_v22_: int
                out17_: bool
                out18_: bool
                out19_: int
                out17_, out18_, out19_ = default__.m0(d_266_cf4_, globalState)
                d_284_v20_ = out17_
                d_285_v21_ = out18_
                d_286_v22_ = out19_
                d_287_v23_: _dafny.Seq
                d_287_v23_ = _dafny.SeqWithoutIsStrInference([d_267_cf3_, p1])
                d_288_v25_: _dafny.Seq
                def iife15_():
                    coll15_ = _dafny.Set()
                    compr_15_: int
                    for compr_15_ in (p0).Elements:
                        d_289_v24_: int = compr_15_
                        if (d_289_v24_) in (p0):
                            coll15_ = coll15_.union(_dafny.Set([(d_289_v24_) * (495)]))
                    return _dafny.Set(coll15_)
                d_288_v25_ = _dafny.SeqWithoutIsStrInference([iife15_()
                ])
                d_290_v26_: _dafny.Map
                d_290_v26_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([(d_287_v23_)[default__.safeIndex(896, len(d_287_v23_))], d_267_cf3_]))) * (default__.fm24((d_288_v25_)[default__.safeIndex((p0)[default__.safeIndex(d_286_v22_, len(p0))], len(d_288_v25_))], globalState)): d_258_v0_})
                d_290_v26_ = (d_290_v26_).set((self).f12, not (not(d_284_v20_)) or (d_258_v0_))
                d_291_v27_: _dafny.Array
                nw52_ = _dafny.Array(None, 1)
                nw52_[int(0)] = (self).f12
                d_291_v27_ = nw52_
                d_292_v28_: _dafny.Map
                d_292_v28_ = _dafny.Map({d_258_v0_: d_270_v6_})
                d_293_v29_: _dafny.MultiSet
                d_293_v29_ = _dafny.MultiSet([not(False), d_267_cf3_])
                d_294_v30_: _dafny.Map
                d_294_v30_ = _dafny.Map({(self).f12: ((d_293_v29_)[d_284_v20_] if (d_284_v20_) in (d_293_v29_) else 293)})
                index34_ = default__.safeIndex(976, (d_291_v27_).length(0))
                (d_291_v27_)[index34_] = default__.safeModuloInt(len(((d_292_v28_)[((d_259_v1_)[p1] if (p1) in (d_259_v1_) else False)] if (((d_259_v1_)[p1] if (p1) in (d_259_v1_) else False)) in (d_292_v28_) else (d_288_v25_)[default__.safeIndex(len(d_294_v30_), len(d_288_v25_))])), (0) - (default__.fm24(d_270_v6_, globalState)))
                index35_ = default__.safeIndex(976, (d_291_v27_).length(0))
                (d_291_v27_)[index35_] = (d_286_v22_) - (d_268_cf2_)
                d_295_v31_: C0
                nw53_ = C0()
                nw53_.ctor__((0) - (-817))
                d_295_v31_ = nw53_
            (globalState).f2 = (self).f12
        elif True:
            d_296___mcc_h3_ = source5_.cf1
            d_297_cf1_ = d_296___mcc_h3_
            d_298_v32_: _dafny.Array
            def lambda9_(d_299_i0_):
                return (self).f11

            init4_ = lambda9_
            nw54_ = _dafny.Array(None, 10)
            for i0_4_ in range(nw54_.length(0)):
                nw54_[i0_4_] = init4_(i0_4_)
            d_298_v32_ = nw54_
            index36_ = default__.safeIndex(478, (d_298_v32_).length(0))
            (d_298_v32_)[index36_] = d_258_v0_
            d_300_v33_: D2
            d_300_v33_ = D2_DC5((self).f12, (self).f11, (self).f11, d_258_v0_, d_258_v0_)
            index37_ = default__.safeIndex(478, (d_298_v32_).length(0))
            (d_298_v32_)[index37_] = (d_300_v33_).cf11
            d_301_v34_: _dafny.MultiSet
            d_301_v34_ = _dafny.MultiSet([d_298_v32_])
            index38_ = default__.safeIndex(478, (d_298_v32_).length(0))
            (d_298_v32_)[index38_] = ((d_301_v34_) - (d_301_v34_)).issubset(d_301_v34_)
            d_302_v35_: _dafny.Set
            d_302_v35_ = _dafny.Set({(self).f11, d_258_v0_, (d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))]})
            d_303_v36_: _dafny.Map
            d_303_v36_ = _dafny.Map({d_302_v35_: d_302_v35_})
            if (d_302_v35_).isdisjoint(((d_303_v36_)[_dafny.Set({d_258_v0_, d_258_v0_})] if (_dafny.Set({d_258_v0_, d_258_v0_})) in (d_303_v36_) else _dafny.Set({d_258_v0_}))):
                d_304_v37_: _dafny.Map
                d_304_v37_ = _dafny.Map({(d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))]: p0})
                d_305_v38_: _dafny.Seq
                d_305_v38_ = _dafny.SeqWithoutIsStrInference([(d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))], (self).f11, (d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))], False])
                d_306_v39_: _dafny.Map
                d_306_v39_ = _dafny.Map({(d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))]: (self).f12})
                d_307_v40_: _dafny.Array
                nw55_ = _dafny.Array(None, 6)
                nw55_[int(0)] = p0
                nw55_[int(1)] = (((d_304_v37_)[not(default__.fm23(p2, (self).f12, 784, (d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))], globalState))] if (not(default__.fm23(p2, (self).f12, 784, (d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))], globalState))) in (d_304_v37_) else p0)) + ((p0).set(default__.safeIndex(len(d_305_v38_), len(p0)), len(d_306_v39_)))
                nw55_[int(2)] = _dafny.SeqWithoutIsStrInference([(self).f12, (self).f12])
                nw55_[int(3)] = p0
                nw55_[int(4)] = p0
                nw55_[int(5)] = p0
                d_307_v40_ = nw55_
                d_308_v41_: _dafny.Map
                d_308_v41_ = _dafny.Map({(self).f12: d_307_v40_})
                d_309_v42_: _dafny.Seq
                d_309_v42_ = _dafny.SeqWithoutIsStrInference([((d_308_v41_)[(self).f12] if ((self).f12) in (d_308_v41_) else d_307_v40_)])
                d_307_v40_ = (d_309_v42_)[default__.safeIndex((self).f12, len(d_309_v42_))]
                d_310_v43_: str
                d_310_v43_ = _dafny.CodePoint('l')
                r1 = d_310_v43_
                index39_ = default__.safeIndex(478, (d_298_v32_).length(0))
                (d_298_v32_)[index39_] = (not((p2) in ((_dafny.SeqWithoutIsStrInference([p2])).set(default__.safeIndex((self).f12, len(_dafny.SeqWithoutIsStrInference([p2]))), p2)))) and (not(default__.fm23(p2, (self).f12, (self).f12, (d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))], globalState)))
                d_311_v44_: _dafny.Array
                nw56_ = _dafny.Array(None, 3)
                d_311_v44_ = nw56_
                d_312_v45_: C0
                nw57_ = C0()
                nw57_.ctor__((self).f12)
                d_312_v45_ = nw57_
                index40_ = default__.safeIndex(607, (d_311_v44_).length(0))
                (d_311_v44_)[index40_] = d_312_v45_
                index41_ = default__.safeIndex(607, (d_311_v44_).length(0))
                (d_311_v44_)[index41_] = d_312_v45_
                index42_ = default__.safeIndex(567, (d_307_v40_).length(0))
                (d_307_v40_)[index42_] = p0
                d_313_v46_: _dafny.Seq
                d_313_v46_ = _dafny.SeqWithoutIsStrInference([d_305_v38_])
                d_314_v47_: _dafny.MultiSet
                d_314_v47_ = _dafny.MultiSet([d_305_v38_, (d_313_v46_)[default__.safeIndex((d_312_v45_).f13, len(d_313_v46_))]])
                d_315_v48_: _dafny.MultiSet
                d_315_v48_ = _dafny.MultiSet([d_258_v0_])
                d_316_v49_: _dafny.Map
                d_316_v49_ = _dafny.Map({d_314_v47_: ((p0).set(default__.safeIndex((p0)[default__.safeIndex(((d_315_v48_)[(d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))]] if ((d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))]) in (d_315_v48_) else len(p2)), len(p0))], len(p0)), (self).f12)) + (p0)})
                index43_ = default__.safeIndex(567, (d_307_v40_).length(0))
                (d_307_v40_)[index43_] = ((d_316_v49_)[d_314_v47_] if (d_314_v47_) in (d_316_v49_) else _dafny.SeqWithoutIsStrInference([-644 for d_317_i1_ in range(default__.abs(651))]))
            elif True:
                d_318_v50_: _dafny.Set
                d_318_v50_ = _dafny.Set({(self).f12})
                d_319_v51_: _dafny.Seq
                d_319_v51_ = _dafny.SeqWithoutIsStrInference([True, p1])
                d_320_v52_: _dafny.MultiSet
                d_320_v52_ = _dafny.MultiSet([p1, p1])
                d_321_v53_: _dafny.Map
                d_321_v53_ = _dafny.Map({(self).f11: (self).f12})
                d_322_v54_: _dafny.Array
                nw58_ = _dafny.Array(None, 24)
                nw58_[int(0)] = 413
                nw58_[int(1)] = (self).f12
                nw58_[int(2)] = (self).f12
                nw58_[int(3)] = (self).f12
                nw58_[int(4)] = (self).f12
                nw58_[int(5)] = (self).f12
                nw58_[int(6)] = (0) - ((self).f12)
                nw58_[int(7)] = default__.fm24(d_318_v50_, globalState)
                nw58_[int(8)] = (0) - (len(d_319_v51_))
                nw58_[int(9)] = len(p2)
                nw58_[int(10)] = (0) - ((self).f12)
                nw58_[int(11)] = (0) - ((self).f12)
                nw58_[int(12)] = (self).f12
                nw58_[int(13)] = (self).f12
                nw58_[int(14)] = 825
                nw58_[int(15)] = (self).f12
                nw58_[int(16)] = (0) - (len(p2))
                nw58_[int(17)] = (self).f12
                nw58_[int(18)] = (self).f12
                nw58_[int(19)] = ((d_320_v52_)[(self).f11] if ((self).f11) in (d_320_v52_) else len(d_321_v53_))
                nw58_[int(20)] = (self).f12
                nw58_[int(21)] = (d_320_v52_).cardinality
                nw58_[int(22)] = (self).f12
                nw58_[int(23)] = 241
                d_322_v54_ = nw58_
                d_323_v55_: _dafny.Seq
                d_323_v55_ = _dafny.SeqWithoutIsStrInference([d_322_v54_, d_322_v54_])
                d_324_v56_: C0
                nw59_ = C0()
                nw59_.ctor__((self).f12)
                d_324_v56_ = nw59_
                rhs29_ = _dafny.SeqWithoutIsStrInference([d_322_v54_, d_322_v54_, d_322_v54_])
                rhs30_ = d_324_v56_
                d_323_v55_ = rhs29_
                d_324_v56_ = rhs30_
                d_325_v57_: _dafny.Seq
                d_325_v57_ = _dafny.SeqWithoutIsStrInference([d_318_v50_, _dafny.Set({(self).f12}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p2: d_258_v0_})) for d_326_i2_ in range(default__.abs(23))]))})])
                d_327_v58_: _dafny.MultiSet
                d_327_v58_ = _dafny.MultiSet([848, default__.fm24((d_325_v57_)[default__.safeIndex((self).f12, len(d_325_v57_))], globalState), len(_dafny.Set({d_324_v56_, d_324_v56_, d_324_v56_}))])
                d_328_v59_: _dafny.Map
                d_328_v59_ = _dafny.Map({len(d_259_v1_): (self).f12})
                (globalState).f2 = ((d_327_v58_)[725] if (725) in (d_327_v58_) else len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_329_i3_ in range(default__.abs(300))]) if default__.fm23(p2, (d_324_v56_).f13, len(d_328_v59_), (d_298_v32_)[default__.safeIndex(478, (d_298_v32_).length(0))], globalState) else p2)))
                d_330_v60_: _dafny.Seq
                d_330_v60_ = _dafny.SeqWithoutIsStrInference([d_319_v51_])
                d_258_v0_ = (d_330_v60_) == (d_330_v60_)
                (globalState).f2 = ((d_300_v33_ if not(d_258_v0_) else d_300_v33_)).cf7
                r0 = ((self).f12) - ((d_324_v56_).f13)
            d_331_v61_: C0
            nw60_ = C0()
            nw60_.ctor__((733) * ((self).f12))
            d_331_v61_ = nw60_
        d_332_v62_: _dafny.Map
        d_332_v62_ = _dafny.Map({p0: p2})
        d_332_v62_ = (d_332_v62_).set(p0, p2)
        d_333_v63_: _dafny.Array
        def lambda10_(d_334_i5_):
            return (d_334_i5_) - (-892)

        init5_ = lambda10_
        nw61_ = _dafny.Array(None, 1)
        for i0_5_ in range(nw61_.length(0)):
            nw61_[i0_5_] = init5_(i0_5_)
        d_333_v63_ = nw61_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_333_v63_).length(0)):
            d_335_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_335_i4_)) and ((d_335_i4_) < ((d_333_v63_).length(0)))):
                (d_333_v63_)[(d_335_i4_)] = default__.safeModuloInt(d_335_i4_, (self).f12)
        r0 = (self).f12
        d_336_v64_: _dafny.Set
        d_336_v64_ = _dafny.Set({_dafny.CodePoint('y')})
        hi1_ = len(d_336_v64_)
        for d_337_i6_ in range((0) - ((self).f12), hi1_):
            d_338_v65_: D3
            d_338_v65_ = D3_DC7((d_337_i6_) + (d_337_i6_), (self).f12)
            source6_ = d_338_v65_
            if source6_.is_DC7:
                d_339___mcc_h4_ = source6_.cf13
                d_340___mcc_h5_ = source6_.cf14
                d_341_cf14_ = d_340___mcc_h5_
                d_342_cf13_ = d_339___mcc_h4_
                d_343_v66_: _dafny.Map
                d_343_v66_ = _dafny.Map({d_342_cf13_: not((self).f11)})
                (globalState).f2 = (d_342_cf13_) + (len(d_343_v66_))
                d_258_v0_ = not((p2) <= (p2))
                d_344_v67_: C0
                nw62_ = C0()
                nw62_.ctor__(d_341_cf14_)
                d_344_v67_ = nw62_
                d_258_v0_ = True
            elif True:
                d_345___mcc_h6_ = source6_.cf12
                d_346_cf12_ = d_345___mcc_h6_
                d_258_v0_ = ((d_337_i6_) - (d_337_i6_)) > (len(p2))
                d_347_v68_: _dafny.Seq
                d_347_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chue"))
                d_348_v69_: _dafny.Array
                nw63_ = _dafny.Array(False, 25)
                d_348_v69_ = nw63_
                d_349_v70_: _dafny.Set
                d_349_v70_ = _dafny.Set({d_348_v69_, d_348_v69_})
                d_350_v71_: _dafny.MultiSet
                d_350_v71_ = _dafny.MultiSet([(d_337_i6_ if True else len(d_349_v70_))])
                rhs31_ = p2
                rhs32_ = (0) - ((d_337_i6_ if (True) or ((self).f11) else (self).f12))
                rhs33_ = (d_350_v71_) - ((_dafny.MultiSet(p0)).intersection(d_350_v71_))
                d_347_v68_ = rhs31_
                r0 = rhs32_
                d_350_v71_ = rhs33_
                d_348_v69_ = d_348_v69_
                d_262_v4_ = d_262_v4_
            d_351_v72_: C0
            nw64_ = C0()
            nw64_.ctor__((self).f12)
            d_351_v72_ = nw64_
            (globalState).f2 = 99
            d_352_v73_: _dafny.MultiSet
            d_352_v73_ = _dafny.MultiSet([p1, (self).f11])
            d_353_v74_: _dafny.Map
            d_353_v74_ = _dafny.Map({d_352_v73_: not(p1)})
            d_354_v75_: _dafny.Map
            d_354_v75_ = _dafny.Map({(_dafny.Map({d_352_v73_: p1})) | (d_353_v74_): default__.fm25(p1, (d_351_v72_).f13, globalState)})
            d_354_v75_ = (d_354_v75_).set((d_353_v74_) | (_dafny.Map({d_352_v73_: p1})), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_355_i7_ in range(default__.abs(147))]))
        r0 = (0) - (default__.safeDivisionInt((self).f12, (D2_DC4((self).f12)).cf6))
        d_356_v76_: str
        d_356_v76_ = _dafny.CodePoint('q')
        r1 = d_356_v76_
        return r0, r1

    def m6(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        r2 = ((self).f12) != (((self).f12) * ((self).f12))
        d_357_v0_: _dafny.Map
        d_357_v0_ = _dafny.Map({(self).f11: True})
        d_358_v1_: _dafny.Map
        d_358_v1_ = _dafny.Map({not(True): d_357_v0_})
        d_359_v2_: _dafny.Set
        d_359_v2_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_360_i1_ in range(default__.abs(760))])), (self).f12, len(d_358_v1_)})
        hi2_ = default__.fm24(d_359_v2_, globalState)
        for d_361_i0_ in range((self).f12, hi2_):
            r2 = ((self).f11 if (self).f11 else (self).f11)
            (globalState).f2 = d_361_i0_
            d_362_v3_: _dafny.Seq
            d_362_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            if not(default__.fm23(d_362_v3_, default__.safeDivisionInt(d_361_i0_, -836), d_361_i0_, (self).f11, globalState)):
                d_362_v3_ = (((d_362_v3_) + (d_362_v3_)) + (d_362_v3_)).set(default__.safeIndex((self).f12, len(((d_362_v3_) + (d_362_v3_)) + (d_362_v3_))), _dafny.CodePoint('f'))
                d_363_v4_: _dafny.Array
                nw65_ = _dafny.Array(False, 16)
                d_363_v4_ = nw65_
                d_364_v5_: D2
                d_364_v5_ = D2_DC5(d_361_i0_, (self).f11, (self).f11, (self).f11, (self).f11)
                index44_ = default__.safeIndex(561, (d_363_v4_).length(0))
                (d_363_v4_)[index44_] = ((d_364_v5_).cf9) or ((self).f11)
                d_365_v6_: _dafny.Array
                def lambda11_(d_366_v5_):
                    def lambda12_(d_367_i2_):
                        return d_366_v5_

                    return lambda12_

                init6_ = lambda11_(d_364_v5_)
                nw66_ = _dafny.Array(None, 22)
                for i0_6_ in range(nw66_.length(0)):
                    nw66_[i0_6_] = init6_(i0_6_)
                d_365_v6_ = nw66_
                d_368_v7_: _dafny.Seq
                d_368_v7_ = _dafny.SeqWithoutIsStrInference([not((self).f11), (self).f11, True, (self).f11, not((self).f11)])
                d_369_v8_: _dafny.MultiSet
                d_369_v8_ = _dafny.MultiSet([(self).f11])
                index45_ = default__.safeIndex(561, (d_363_v4_).length(0))
                rhs34_ = ((self).f12) != (len(d_368_v7_))
                rhs35_ = ((d_369_v8_).set(False, default__.abs(d_361_i0_))) != ((d_369_v8_) - (d_369_v8_))
                rhs36_ = (self).f11
                rhs37_ = d_365_v6_
                lhs17_ = d_363_v4_
                lhs18_ = default__.safeIndex(561, (d_363_v4_).length(0))
                r2 = rhs34_
                lhs17_[lhs18_] = rhs35_
                r2 = rhs36_
                d_365_v6_ = rhs37_
                d_370_v9_: str
                d_370_v9_ = _dafny.CodePoint('j')
                d_371_v10_: D1
                d_371_v10_ = D1_DC2((self).f12, (d_363_v4_)[default__.safeIndex(561, (d_363_v4_).length(0))], -552)
                d_372_v11_: _dafny.Array
                nw67_ = _dafny.Array(None, 4)
                nw67_[int(0)] = d_368_v7_
                nw67_[int(1)] = (d_368_v7_) + (d_368_v7_)
                nw67_[int(2)] = d_368_v7_
                nw67_[int(3)] = (d_368_v7_ if (d_371_v10_).cf3 else _dafny.SeqWithoutIsStrInference([(d_363_v4_)[default__.safeIndex(561, (d_363_v4_).length(0))]]))
                d_372_v11_ = nw67_
                index46_ = default__.safeIndex(807, (d_372_v11_).length(0))
                (d_372_v11_)[index46_] = d_368_v7_
                d_373_v12_: _dafny.Map
                d_373_v12_ = _dafny.Map({(self).f11: d_370_v9_})
                d_374_v13_: _dafny.Set
                d_374_v13_ = _dafny.Set({((d_373_v12_)[(d_363_v4_)[default__.safeIndex(561, (d_363_v4_).length(0))]] if ((d_363_v4_)[default__.safeIndex(561, (d_363_v4_).length(0))]) in (d_373_v12_) else _dafny.CodePoint('w')), d_370_v9_})
                d_375_v14_: _dafny.Array
                nw68_ = _dafny.Array(None, 7)
                nw68_[int(0)] = len(d_362_v3_)
                nw68_[int(1)] = len(d_374_v13_)
                nw68_[int(2)] = (self).f12
                nw68_[int(3)] = d_361_i0_
                nw68_[int(4)] = (self).f12
                nw68_[int(5)] = d_361_i0_
                nw68_[int(6)] = (self).f12
                d_375_v14_ = nw68_
                d_376_v15_: _dafny.Seq
                d_376_v15_ = _dafny.SeqWithoutIsStrInference([d_375_v14_, d_375_v14_, d_375_v14_])
                d_377_v16_: _dafny.Array
                nw69_ = _dafny.Array(None, 3)
                nw69_[int(0)] = (d_376_v15_)[default__.safeIndex(d_361_i0_, len(d_376_v15_))]
                nw69_[int(1)] = d_375_v14_
                nw69_[int(2)] = d_375_v14_
                d_377_v16_ = nw69_
                index47_ = default__.safeIndex(113, (d_377_v16_).length(0))
                (d_377_v16_)[index47_] = d_375_v14_
                d_378_v17_: _dafny.Set
                d_378_v17_ = _dafny.Set({(d_363_v4_)[default__.safeIndex(561, (d_363_v4_).length(0))], (d_363_v4_)[default__.safeIndex(561, (d_363_v4_).length(0))]})
                index48_ = default__.safeIndex(807, (d_372_v11_).length(0))
                index49_ = default__.safeIndex(113, (d_377_v16_).length(0))
                rhs38_ = default__.fm26(default__.safeDivisionInt(d_361_i0_, d_361_i0_), len((_dafny.Set({(self).f11})) | (d_378_v17_)), (d_363_v4_)[default__.safeIndex(561, (d_363_v4_).length(0))], globalState)
                rhs39_ = (d_368_v7_) + (d_368_v7_)
                rhs40_ = (0) - ((len((d_368_v7_) + (d_368_v7_)) if (len(d_362_v3_)) <= (d_361_i0_) else (self).f12))
                rhs41_ = d_375_v14_
                rhs42_ = (d_368_v7_).set(default__.safeIndex((self).f12, len(d_368_v7_)), ((self).f11 if True else (d_363_v4_)[default__.safeIndex(561, (d_363_v4_).length(0))]))
                lhs19_ = d_372_v11_
                lhs20_ = default__.safeIndex(807, (d_372_v11_).length(0))
                lhs21_ = globalState
                lhs22_ = d_377_v16_
                lhs23_ = default__.safeIndex(113, (d_377_v16_).length(0))
                d_370_v9_ = rhs38_
                lhs19_[lhs20_] = rhs39_
                lhs21_.f2 = rhs40_
                lhs22_[lhs23_] = rhs41_
                d_368_v7_ = rhs42_
                (globalState).f2 = d_361_i0_
                index50_ = default__.safeIndex(561, (d_363_v4_).length(0))
                (d_363_v4_)[index50_] = (self).f11
            elif True:
                d_379_v18_: _dafny.Array
                def lambda13_(d_380_i3_):
                    return (d_380_i3_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kisi"))))

                init7_ = lambda13_
                nw70_ = _dafny.Array(None, 7)
                for i0_7_ in range(nw70_.length(0)):
                    nw70_[i0_7_] = init7_(i0_7_)
                d_379_v18_ = nw70_
                d_381_v19_: _dafny.MultiSet
                d_381_v19_ = _dafny.MultiSet([d_379_v18_])
                (globalState).f2 = ((d_381_v19_)[d_379_v18_] if (d_379_v18_) in (d_381_v19_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "styupjoxw"))))
                d_382_v20_: _dafny.Seq
                d_382_v20_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_383_v21_: C0
                nw71_ = C0()
                nw71_.ctor__(d_361_i0_)
                d_383_v21_ = nw71_
                d_384_v22_: D4
                d_384_v22_ = D4_DC8(d_383_v21_)
                d_385_v23_: _dafny.Map
                d_385_v23_ = _dafny.Map({(d_383_v21_).f13: d_383_v21_})
                d_386_v24_: _dafny.Array
                nw72_ = _dafny.Array(None, 27)
                nw72_[int(0)] = d_383_v21_
                nw72_[int(1)] = d_383_v21_
                nw72_[int(2)] = d_383_v21_
                nw72_[int(3)] = d_383_v21_
                nw72_[int(4)] = d_383_v21_
                nw72_[int(5)] = d_383_v21_
                nw72_[int(6)] = d_383_v21_
                nw72_[int(7)] = d_383_v21_
                nw72_[int(8)] = d_383_v21_
                nw72_[int(9)] = d_383_v21_
                nw72_[int(10)] = (d_384_v22_).cf15
                nw72_[int(11)] = d_383_v21_
                nw72_[int(12)] = ((d_385_v23_)[(d_383_v21_).f13] if ((d_383_v21_).f13) in (d_385_v23_) else d_383_v21_)
                nw72_[int(13)] = d_383_v21_
                nw72_[int(14)] = d_383_v21_
                nw72_[int(15)] = d_383_v21_
                nw72_[int(16)] = d_383_v21_
                nw72_[int(17)] = d_383_v21_
                nw72_[int(18)] = d_383_v21_
                nw72_[int(19)] = d_383_v21_
                nw72_[int(20)] = d_383_v21_
                nw72_[int(21)] = d_383_v21_
                nw72_[int(22)] = d_383_v21_
                nw72_[int(23)] = ((d_385_v23_)[(d_383_v21_).f13] if ((d_383_v21_).f13) in (d_385_v23_) else d_383_v21_)
                nw72_[int(24)] = d_383_v21_
                nw72_[int(25)] = d_383_v21_
                nw72_[int(26)] = d_383_v21_
                d_386_v24_ = nw72_
                d_387_v25_: _dafny.Map
                d_387_v25_ = _dafny.Map({d_386_v24_: (d_383_v21_).f13})
                d_388_v26_: _dafny.Set
                d_388_v26_ = _dafny.Set({(self).f11, (len(d_382_v20_)) == (len(d_387_v25_)), (self).f11})
                d_389_v27_: _dafny.Set
                d_389_v27_ = d_388_v26_
                d_388_v26_ = (d_388_v26_).intersection((d_389_v27_))
                d_390_v28_: _dafny.Seq
                d_390_v28_ = _dafny.SeqWithoutIsStrInference([d_361_i0_])
                (globalState).f2 = ((self).f12 if not((d_390_v28_) < (_dafny.SeqWithoutIsStrInference([885 for d_391_i4_ in range(default__.abs(297))]))) else (self).f12)
                d_392_v29_: _dafny.MultiSet
                d_392_v29_ = _dafny.MultiSet([(self).f11])
                (globalState).f2 = ((d_392_v29_)[True] if (True) in (d_392_v29_) else d_361_i0_)
                r2 = (self).f11
            d_393_v30_: _dafny.Array
            def lambda14_(d_394_i5_):
                return (d_394_i5_) * ((self).f12)

            init8_ = lambda14_
            nw73_ = _dafny.Array(None, 12)
            for i0_8_ in range(nw73_.length(0)):
                nw73_[i0_8_] = init8_(i0_8_)
            d_393_v30_ = nw73_
            d_393_v30_ = d_393_v30_
        d_395_v31_: _dafny.Array
        nw74_ = _dafny.Array(int(0), 5)
        d_395_v31_ = nw74_
        d_396_v32_: _dafny.Array
        nw75_ = _dafny.Array(None, 22)
        nw75_[int(0)] = d_395_v31_
        nw75_[int(1)] = d_395_v31_
        nw75_[int(2)] = d_395_v31_
        nw75_[int(3)] = d_395_v31_
        nw75_[int(4)] = d_395_v31_
        nw75_[int(5)] = d_395_v31_
        nw75_[int(6)] = d_395_v31_
        nw75_[int(7)] = d_395_v31_
        nw75_[int(8)] = d_395_v31_
        nw75_[int(9)] = d_395_v31_
        nw75_[int(10)] = d_395_v31_
        nw75_[int(11)] = d_395_v31_
        nw75_[int(12)] = d_395_v31_
        nw75_[int(13)] = d_395_v31_
        nw75_[int(14)] = d_395_v31_
        nw75_[int(15)] = d_395_v31_
        nw75_[int(16)] = d_395_v31_
        nw75_[int(17)] = d_395_v31_
        nw75_[int(18)] = d_395_v31_
        nw75_[int(19)] = d_395_v31_
        nw75_[int(20)] = d_395_v31_
        nw75_[int(21)] = d_395_v31_
        d_396_v32_ = nw75_
        index51_ = default__.safeIndex(923, (d_396_v32_).length(0))
        (d_396_v32_)[index51_] = d_395_v31_
        d_397_v33_: _dafny.Array
        def lambda15_(d_398_i6_):
            return (D6_DC11((D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwqyvo")))).cf22)).cf22

        init9_ = lambda15_
        nw76_ = _dafny.Array(None, 24)
        for i0_9_ in range(nw76_.length(0)):
            nw76_[i0_9_] = init9_(i0_9_)
        d_397_v33_ = nw76_
        d_399_v34_: _dafny.Seq
        d_399_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
        d_400_v35_: str
        d_400_v35_ = _dafny.CodePoint('x')
        d_401_v36_: D6
        d_401_v36_ = D6_DC11((d_399_v34_).set(default__.safeIndex(261, len(d_399_v34_)), d_400_v35_))
        index52_ = default__.safeIndex(772, (d_397_v33_).length(0))
        (d_397_v33_)[index52_] = (d_401_v36_).cf22
        index53_ = default__.safeIndex(923, (d_396_v32_).length(0))
        index54_ = default__.safeIndex(772, (d_397_v33_).length(0))
        rhs43_ = (d_395_v31_ if (self).f11 else d_395_v31_)
        rhs44_ = ((_dafny.SeqWithoutIsStrInference([d_400_v35_ for d_402_i7_ in range(default__.abs(48))])) + (d_399_v34_)) + ((_dafny.SeqWithoutIsStrInference([d_400_v35_ for d_403_i8_ in range(default__.abs(569))])) + (d_399_v34_))
        rhs45_ = (self).f12
        rhs46_ = (self).f11
        lhs24_ = d_396_v32_
        lhs25_ = default__.safeIndex(923, (d_396_v32_).length(0))
        lhs26_ = d_397_v33_
        lhs27_ = default__.safeIndex(772, (d_397_v33_).length(0))
        lhs24_[lhs25_] = rhs43_
        lhs26_[lhs27_] = rhs44_
        r1 = rhs45_
        r2 = rhs46_
        d_404_v37_: bool
        d_405_v38_: bool
        d_406_v39_: int
        out20_: bool
        out21_: bool
        out22_: int
        out20_, out21_, out22_ = default__.m0((self).f12, globalState)
        d_404_v37_ = out20_
        d_405_v38_ = out21_
        d_406_v39_ = out22_
        d_407_v40_: _dafny.Seq
        d_407_v40_ = _dafny.SeqWithoutIsStrInference([(self).f12])
        hi3_ = len(d_407_v40_)
        for d_408_i9_ in range(d_406_v39_, hi3_):
            d_409_v41_: _dafny.Seq
            d_409_v41_ = _dafny.SeqWithoutIsStrInference([default__.fm27(d_407_v40_, not(d_404_v37_), d_404_v37_, globalState)])
            r0 = len(d_409_v41_)
            arr0_ = (d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]
            index55_ = default__.safeIndex(816, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]).length(0))
            arr0_[index55_] = 483
            d_410_v42_: _dafny.Seq
            d_410_v42_ = _dafny.SeqWithoutIsStrInference([(d_357_v0_).set(d_405_v38_, False)])
            d_411_v43_: D1
            d_411_v43_ = D1_DC1(d_410_v42_)
            arr1_ = (d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]
            index56_ = default__.safeIndex(816, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]).length(0))
            rhs47_ = _dafny.CodePoint('r')
            rhs48_ = 707
            rhs49_ = d_411_v43_
            lhs28_ = (d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]
            lhs29_ = default__.safeIndex(816, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]).length(0))
            d_400_v35_ = rhs47_
            lhs28_[lhs29_] = rhs48_
            d_411_v43_ = rhs49_
            d_406_v39_ = default__.safeModuloInt(((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))])[default__.safeIndex(816, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]).length(0))], (((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))])[default__.safeIndex(816, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]).length(0))]) * (len(d_407_v40_)))
            d_412_v44_: _dafny.Set
            d_412_v44_ = _dafny.Set({d_395_v31_})
            d_413_v45_: D1
            d_413_v45_ = D1_DC2(((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))])[default__.safeIndex(816, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]).length(0))], d_404_v37_, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))])[default__.safeIndex(816, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]).length(0))])
            arr2_ = (d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]
            index57_ = default__.safeIndex(816, ((d_396_v32_)[default__.safeIndex(923, (d_396_v32_).length(0))]).length(0))
            arr2_[index57_] = (len(d_412_v44_)) + ((d_413_v45_).cf4)
        pat_let_tv2_ = d_406_v39_
        pat_let_tv3_ = d_406_v39_
        pat_let_tv4_ = d_406_v39_
        pat_let_tv5_ = d_406_v39_
        pat_let_tv6_ = d_406_v39_
        pat_let_tv7_ = d_406_v39_
        pat_let_tv8_ = d_406_v39_
        pat_let_tv9_ = d_406_v39_
        pat_let_tv10_ = d_406_v39_
        def lambda16_(source8_):
            if source8_.is_DC2:
                d_414___mcc_h7_ = source8_.cf2
                d_415___mcc_h8_ = source8_.cf3
                d_416___mcc_h9_ = source8_.cf4
                d_417_cf4_ = d_416___mcc_h9_
                d_418_cf3_ = d_415___mcc_h8_
                d_419_cf2_ = d_414___mcc_h7_
                def iife16_():
                    coll16_ = _dafny.Map()
                    compr_16_: _dafny.MultiSet
                    for compr_16_ in (_dafny.Set({_dafny.MultiSet([pat_let_tv2_]), _dafny.MultiSet([pat_let_tv3_, pat_let_tv4_, (self).f12, pat_let_tv5_])})).Elements:
                        d_420_v46_: _dafny.MultiSet = compr_16_
                        if (d_420_v46_) in (_dafny.Set({_dafny.MultiSet([pat_let_tv6_]), _dafny.MultiSet([pat_let_tv7_, pat_let_tv8_, (self).f12, pat_let_tv9_])})):
                            coll16_[d_420_v46_] = d_419_cf2_
                    return _dafny.Map(coll16_)
                return D2_DC3(iife16_()
)
            elif True:
                d_421___mcc_h10_ = source8_.cf1
                d_422_cf1_ = d_421___mcc_h10_
                return D2_DC3(_dafny.Map({_dafny.MultiSet([-146]): pat_let_tv10_}))

        source7_ = lambda16_(D1_DC2(d_406_v39_, (self).f11, (self).f12))
        if source7_.is_DC4:
            d_423___mcc_h0_ = source7_.cf6
            d_424_cf6_ = d_423___mcc_h0_
            d_404_v37_ = (d_404_v37_) == (False)
            d_425_v47_: C0
            nw77_ = C0()
            nw77_.ctor__(d_424_cf6_)
            d_425_v47_ = nw77_
            index58_ = default__.safeIndex(311, (d_395_v31_).length(0))
            (d_395_v31_)[index58_] = 809
            index59_ = default__.safeIndex(311, (d_395_v31_).length(0))
            (d_395_v31_)[index59_] = (d_406_v39_) * ((d_424_cf6_) * (d_424_cf6_))
            d_424_cf6_ = (0) - ((d_425_v47_).f13)
        elif source7_.is_DC5:
            d_426___mcc_h1_ = source7_.cf7
            d_427___mcc_h2_ = source7_.cf8
            d_428___mcc_h3_ = source7_.cf9
            d_429___mcc_h4_ = source7_.cf10
            d_430___mcc_h5_ = source7_.cf11
            d_431_cf11_ = d_430___mcc_h5_
            d_432_cf10_ = d_429___mcc_h4_
            d_433_cf9_ = d_428___mcc_h3_
            d_434_cf8_ = d_427___mcc_h2_
            d_435_cf7_ = d_426___mcc_h1_
            d_436_v48_: _dafny.Seq
            d_436_v48_ = _dafny.SeqWithoutIsStrInference([d_399_v34_])
            d_437_v49_: _dafny.Map
            d_437_v49_ = _dafny.Map({False: d_400_v35_})
            d_438_v50_: _dafny.Array
            nw78_ = _dafny.Array(None, 12)
            nw78_[int(0)] = d_433_cf9_
            nw78_[int(1)] = not(True)
            nw78_[int(2)] = d_404_v37_
            nw78_[int(3)] = d_432_cf10_
            nw78_[int(4)] = d_405_v38_
            nw78_[int(5)] = d_404_v37_
            nw78_[int(6)] = not(d_431_cf11_)
            nw78_[int(7)] = d_404_v37_
            nw78_[int(8)] = d_434_cf8_
            nw78_[int(9)] = d_432_cf10_
            nw78_[int(10)] = (d_436_v48_) != (d_436_v48_)
            nw78_[int(11)] = (((d_437_v49_)[((d_357_v0_)[d_431_cf11_] if (d_431_cf11_) in (d_357_v0_) else d_432_cf10_)] if (((d_357_v0_)[d_431_cf11_] if (d_431_cf11_) in (d_357_v0_) else d_432_cf10_)) in (d_437_v49_) else d_400_v35_)) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odq")))
            d_438_v50_ = nw78_
            index60_ = default__.safeIndex(510, (d_438_v50_).length(0))
            (d_438_v50_)[index60_] = d_431_cf11_
            d_439_v52_: _dafny.Map
            d_439_v52_ = _dafny.Map({len((d_397_v33_)[default__.safeIndex(772, (d_397_v33_).length(0))]): d_405_v38_})
            index61_ = default__.safeIndex(510, (d_438_v50_).length(0))
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(865, 185):
                    d_440_v51_: int = compr_17_
                    if ((865) <= (d_440_v51_)) and ((d_440_v51_) < (185)):
                        coll17_[default__.safeDivisionInt(d_440_v51_, 235)] = True
                return _dafny.Map(coll17_)
            rhs50_ = d_435_cf7_
            rhs51_ = (0) - (len(((iife17_()
            ) | (((d_439_v52_).set((self).f12, d_431_cf11_)).set(len(d_407_v40_), default__.fm23((d_397_v33_)[default__.safeIndex(772, (d_397_v33_).length(0))], (self).f12, len(_dafny.SeqWithoutIsStrInference([d_400_v35_ for d_441_i10_ in range(default__.abs(145))])), d_431_cf11_, globalState)))) | ((d_439_v52_).set((self).f12, d_431_cf11_))))
            rhs52_ = d_434_cf8_
            rhs53_ = True
            lhs30_ = d_438_v50_
            lhs31_ = default__.safeIndex(510, (d_438_v50_).length(0))
            r1 = rhs50_
            r1 = rhs51_
            lhs30_[lhs31_] = rhs52_
            d_433_cf9_ = rhs53_
            d_442_v53_: D3
            d_442_v53_ = D3_DC6(d_400_v35_)
            d_400_v35_ = (d_442_v53_).cf12
            r2 = (d_438_v50_)[default__.safeIndex(510, (d_438_v50_).length(0))]
            d_443_v54_: D2
            d_443_v54_ = D2_DC5(d_435_cf7_, d_431_cf11_, (self).f11, d_431_cf11_, d_434_cf8_)
            pat_let_tv11_ = d_431_cf11_
            pat_let_tv12_ = d_438_v50_
            pat_let_tv13_ = d_438_v50_
            def iife18_(_pat_let0_0):
                def iife19_(d_444_dt__update__tmp_h0_):
                    def iife20_(_pat_let1_0):
                        def iife21_(d_445_dt__update_hcf9_h0_):
                            def iife22_(_pat_let2_0):
                                def iife23_(d_446_dt__update_hcf8_h0_):
                                    def iife24_(_pat_let3_0):
                                        def iife25_(d_447_dt__update_hcf11_h0_):
                                            return D2_DC5((d_444_dt__update__tmp_h0_).cf7, d_446_dt__update_hcf8_h0_, d_445_dt__update_hcf9_h0_, (d_444_dt__update__tmp_h0_).cf10, d_447_dt__update_hcf11_h0_)
                                        return iife25_(_pat_let3_0)
                                    return iife24_((self).f11)
                                return iife23_(_pat_let2_0)
                            return iife22_((pat_let_tv13_)[default__.safeIndex(510, (pat_let_tv12_).length(0))])
                        return iife21_(_pat_let1_0)
                    return iife20_(not(not(pat_let_tv11_)))
                return iife19_(_pat_let0_0)
            d_434_cf8_ = (iife18_(d_443_v54_)).cf11
        elif True:
            d_448___mcc_h6_ = source7_.cf5
            d_449_cf5_ = d_448___mcc_h6_
            r1 = d_406_v39_
            d_450_v55_: C0
            nw79_ = C0()
            nw79_.ctor__((self).f12)
            d_450_v55_ = nw79_
            d_451_v56_: _dafny.Map
            d_451_v56_ = _dafny.Map({(self).f12: (d_450_v55_).f13})
            d_405_v38_ = ((d_406_v39_) in (d_451_v56_)) and ((self).f11)
            r1 = (self).f12
        r0 = d_406_v39_
        d_452_v57_: _dafny.MultiSet
        d_452_v57_ = _dafny.MultiSet([True])
        d_453_v58_: _dafny.Seq
        d_453_v58_ = _dafny.SeqWithoutIsStrInference([(self).f11, d_404_v37_])
        r1 = ((d_406_v39_) - ((d_452_v57_).cardinality)) * (len(((d_397_v33_)[default__.safeIndex(772, (d_397_v33_).length(0))]).set(default__.safeIndex(len(d_453_v58_), len((d_397_v33_)[default__.safeIndex(772, (d_397_v33_).length(0))])), d_400_v35_)))
        r2 = d_404_v37_
        return r0, r1, r2

    def m7(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: int = int(0)
        d_454_v0_: _dafny.Array
        nw80_ = _dafny.Array(False, 13)
        d_454_v0_ = nw80_
        d_455_v1_: _dafny.Seq
        d_455_v1_ = _dafny.SeqWithoutIsStrInference([(self).f12])
        d_456_v2_: _dafny.Map
        d_456_v2_ = _dafny.Map({d_454_v0_: d_455_v1_})
        d_456_v2_ = ((d_456_v2_) | (d_456_v2_)) | ((d_456_v2_) | (d_456_v2_))
        if p1:
            d_457_v3_: _dafny.Seq
            d_457_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmjjfn"))
            r1 = (d_457_v3_) + (d_457_v3_)
            d_458_v4_: _dafny.Array
            nw81_ = _dafny.Array(None, 6)
            nw81_[int(0)] = d_454_v0_
            nw81_[int(1)] = d_454_v0_
            nw81_[int(2)] = d_454_v0_
            nw81_[int(3)] = d_454_v0_
            nw81_[int(4)] = d_454_v0_
            nw81_[int(5)] = d_454_v0_
            d_458_v4_ = nw81_
            index62_ = default__.safeIndex(881, (d_458_v4_).length(0))
            (d_458_v4_)[index62_] = d_454_v0_
            d_459_v5_: _dafny.Map
            d_459_v5_ = _dafny.Map({((self).f12) + ((self).f12): d_454_v0_})
            index63_ = default__.safeIndex(881, (d_458_v4_).length(0))
            (d_458_v4_)[index63_] = ((d_459_v5_)[(self).f12] if ((self).f12) in (d_459_v5_) else d_454_v0_)
            d_460_v6_: _dafny.Map
            d_460_v6_ = _dafny.Map({not(p0): p1})
            d_461_v7_: _dafny.Set
            d_461_v7_ = _dafny.Set({d_460_v6_, d_460_v6_, d_460_v6_})
            d_461_v7_ = d_461_v7_
            d_462_v8_: bool
            d_462_v8_ = True
            d_463_v9_: _dafny.Seq
            d_463_v9_ = _dafny.SeqWithoutIsStrInference([d_460_v6_])
            d_462_v8_ = default__.fm23(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_464_i0_ in range(default__.abs(140))]), len(d_463_v9_), (self).f12, p0, globalState)
            d_462_v8_ = ((self).f12) < (((self).f12 if p0 else 941))
        elif True:
            d_465_v10_: _dafny.Set
            d_465_v10_ = _dafny.Set({(self).f12})
            (globalState).f2 = default__.fm24((d_465_v10_) - (d_465_v10_), globalState)
            d_466_v11_: C0
            nw82_ = C0()
            nw82_.ctor__((self).f12)
            d_466_v11_ = nw82_
            d_467_v12_: _dafny.Seq
            d_467_v12_ = _dafny.SeqWithoutIsStrInference([d_466_v11_])
            d_468_v13_: _dafny.Array
            nw83_ = _dafny.Array(None, 6)
            nw83_[int(0)] = len(d_467_v12_)
            nw83_[int(1)] = default__.safeModuloInt((d_466_v11_).f13, (d_466_v11_).f13)
            nw83_[int(2)] = ((self).f12) - ((0) - ((d_466_v11_).f13))
            nw83_[int(3)] = len(((d_455_v1_) + (_dafny.SeqWithoutIsStrInference([(d_466_v11_).f13, len(d_455_v1_), (self).f12]))).set(default__.safeIndex((d_466_v11_).f13, len((d_455_v1_) + (_dafny.SeqWithoutIsStrInference([(d_466_v11_).f13, len(d_455_v1_), (self).f12])))), (d_466_v11_).f13))
            nw83_[int(4)] = (self).f12
            nw83_[int(5)] = (self).f12
            d_468_v13_ = nw83_
            index64_ = default__.safeIndex(352, (d_468_v13_).length(0))
            (d_468_v13_)[index64_] = (self).f12
            d_469_v14_: D3
            d_469_v14_ = D3_DC7((self).f12, (316) * ((self).f12))
            d_470_v15_: _dafny.MultiSet
            d_470_v15_ = _dafny.MultiSet([(d_466_v11_).f13, (self).f12])
            d_471_v16_: _dafny.Map
            d_471_v16_ = _dafny.Map({d_470_v15_: -719})
            d_472_v17_: D2
            d_472_v17_ = D2_DC3(d_471_v16_)
            d_473_v18_: D2
            d_473_v18_ = D2_DC3((d_472_v17_).cf5)
            d_474_v19_: D4
            d_474_v19_ = D4_DC9((self).f12, _dafny.CodePoint('e'), d_455_v1_, (d_466_v11_).f13, (self).f11)
            d_475_v20_: _dafny.Seq
            d_475_v20_ = _dafny.SeqWithoutIsStrInference([d_474_v19_])
            d_476_v21_: _dafny.Seq
            d_476_v21_ = _dafny.SeqWithoutIsStrInference([d_475_v20_])
            index65_ = default__.safeIndex(352, (d_468_v13_).length(0))
            rhs54_ = default__.safeModuloInt((d_466_v11_).f13, (d_466_v11_).f13)
            rhs55_ = len(d_476_v21_)
            rhs56_ = d_469_v14_
            rhs57_ = d_472_v17_
            lhs32_ = d_468_v13_
            lhs33_ = default__.safeIndex(352, (d_468_v13_).length(0))
            r0 = rhs54_
            lhs32_[lhs33_] = rhs55_
            d_469_v14_ = rhs56_
            d_472_v17_ = rhs57_
            d_477_v22_: int
            d_478_v23_: int
            d_479_v24_: bool
            out23_: int
            out24_: int
            out25_: bool
            out23_, out24_, out25_ = (self).m6(globalState)
            d_477_v22_ = out23_
            d_478_v23_ = out24_
            d_479_v24_ = out25_
            d_479_v24_ = not(((self).f11 if p1 else True))
            d_480_v25_: _dafny.Array
            nw84_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_480_v25_ = nw84_
            index66_ = default__.safeIndex(966, (d_480_v25_).length(0))
            (d_480_v25_)[index66_] = d_468_v13_
            index67_ = default__.safeIndex(251, (d_454_v0_).length(0))
            (d_454_v0_)[index67_] = not(not(False))
            index68_ = default__.safeIndex(352, (d_468_v13_).length(0))
            index69_ = default__.safeIndex(966, (d_480_v25_).length(0))
            index70_ = default__.safeIndex(251, (d_454_v0_).length(0))
            rhs58_ = 365
            rhs59_ = d_468_v13_
            rhs60_ = p1
            lhs34_ = d_468_v13_
            lhs35_ = default__.safeIndex(352, (d_468_v13_).length(0))
            lhs36_ = d_480_v25_
            lhs37_ = default__.safeIndex(966, (d_480_v25_).length(0))
            lhs38_ = d_454_v0_
            lhs39_ = default__.safeIndex(251, (d_454_v0_).length(0))
            lhs34_[lhs35_] = rhs58_
            lhs36_[lhs37_] = rhs59_
            lhs38_[lhs39_] = rhs60_
        d_481_v26_: _dafny.Seq
        d_481_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtttnk"))
        d_482_v28_: _dafny.Array
        def lambda17_(d_483_i1_):
            def iife26_():
                coll18_ = _dafny.Map()
                compr_18_: _dafny.MultiSet
                for compr_18_ in (_dafny.MultiSet([_dafny.MultiSet([91]), _dafny.MultiSet([(self).f12])])).Elements:
                    d_484_v27_: _dafny.MultiSet = compr_18_
                    if (d_484_v27_) in (_dafny.MultiSet([_dafny.MultiSet([91]), _dafny.MultiSet([(self).f12])])):
                        coll18_[d_484_v27_] = (self).f12
                return _dafny.Map(coll18_)
            return D2_DC3(iife26_()
)

        init10_ = lambda17_
        nw85_ = _dafny.Array(None, 24)
        for i0_10_ in range(nw85_.length(0)):
            nw85_[i0_10_] = init10_(i0_10_)
        d_482_v28_ = nw85_
        d_485_v29_: _dafny.Map
        d_485_v29_ = _dafny.Map({d_482_v28_: d_481_v26_})
        d_486_v30_: str
        d_486_v30_ = _dafny.CodePoint('u')
        r1 = (((default__.fm25(not((self).f11), (self).f12, globalState)) + (d_481_v26_)) + (((d_485_v29_)[d_482_v28_] if (d_482_v28_) in (d_485_v29_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))))).set(default__.safeIndex((self).f12, len(((default__.fm25(not((self).f11), (self).f12, globalState)) + (d_481_v26_)) + (((d_485_v29_)[d_482_v28_] if (d_482_v28_) in (d_485_v29_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))))), d_486_v30_)
        (globalState).f2 = 238
        d_487_v31_: _dafny.Map
        d_487_v31_ = _dafny.Map({p1: d_454_v0_})
        d_488_v32_: _dafny.Map
        d_488_v32_ = _dafny.Map({(self).f11: ((d_487_v31_)[(self).f11] if ((self).f11) in (d_487_v31_) else d_454_v0_)})
        d_487_v31_ = d_488_v32_
        d_489_v33_: _dafny.Set
        d_489_v33_ = _dafny.Set({616, (self).f12})
        d_490_v35_: D1
        def iife27_():
            coll19_ = _dafny.Set()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(835, 567):
                d_491_v34_: int = compr_19_
                if ((835) <= (d_491_v34_)) and ((d_491_v34_) < (567)):
                    coll19_ = coll19_.union(_dafny.Set([(d_491_v34_) - (97)]))
            return _dafny.Set(coll19_)
        d_490_v35_ = D1_DC2(((self).f12) + (611), not((d_489_v33_) == (iife27_()
)), (self).f12)
        source9_ = d_490_v35_
        if source9_.is_DC2:
            d_492___mcc_h0_ = source9_.cf2
            d_493___mcc_h1_ = source9_.cf3
            d_494___mcc_h2_ = source9_.cf4
            d_495_cf4_ = d_494___mcc_h2_
            d_496_cf3_ = d_493___mcc_h1_
            d_497_cf2_ = d_492___mcc_h0_
            d_496_cf3_ = (self).f11
            d_498_v36_: C0
            nw86_ = C0()
            nw86_.ctor__(-580)
            d_498_v36_ = nw86_
            d_499_v37_: _dafny.Set
            d_499_v37_ = _dafny.Set({d_486_v30_, d_486_v30_})
            d_500_v38_: D7
            d_500_v38_ = D7_DC13(d_499_v37_)
            d_499_v37_ = (d_499_v37_) | ((d_500_v38_).cf25)
            d_501_v39_: _dafny.Map
            d_501_v39_ = _dafny.Map({(self).f12: len(d_481_v26_)})
            d_502_v40_: _dafny.MultiSet
            d_502_v40_ = _dafny.MultiSet([d_501_v39_])
            d_503_v41_: D2
            d_503_v41_ = D2_DC4(((d_502_v40_)[default__.fm28(d_497_cf2_, globalState)] if (default__.fm28(d_497_cf2_, globalState)) in (d_502_v40_) else d_497_cf2_))
            d_504_v42_: _dafny.Seq
            d_504_v42_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
            d_505_v43_: _dafny.Map
            d_505_v43_ = _dafny.Map({D7_DC13(d_499_v37_): -23})
            d_503_v41_ = ((d_503_v41_ if p1 else D2_DC4((self).f12)) if (d_504_v42_)[default__.safeIndex(len(d_505_v43_), len(d_504_v42_))] else d_503_v41_)
        elif True:
            d_506___mcc_h3_ = source9_.cf1
            d_507_cf1_ = d_506___mcc_h3_
            if p0:
                d_508_v44_: _dafny.Set
                d_508_v44_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), d_481_v26_})
                d_509_v45_: D6
                d_509_v45_ = D6_DC12((self).f11, default__.fm26((self).f12, (self).f12, (self).f11, globalState))
                d_481_v26_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xramcle"))).set(default__.safeIndex(len(d_508_v44_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xramcle")))), (d_509_v45_).cf24)) + (d_481_v26_)
                d_510_v46_: C0
                nw87_ = C0()
                nw87_.ctor__((self).f12)
                d_510_v46_ = nw87_
                d_511_v47_: bool
                d_511_v47_ = False
                d_512_v48_: _dafny.Seq
                d_512_v48_ = _dafny.SeqWithoutIsStrInference([d_510_v46_, d_510_v46_, d_510_v46_])
                d_513_v49_: _dafny.Map
                d_513_v49_ = _dafny.Map({(d_510_v46_).f13: d_481_v26_})
                index71_ = default__.safeIndex(451, (d_454_v0_).length(0))
                (d_454_v0_)[index71_] = p1
                d_514_v50_: D8
                d_514_v50_ = D8_DC15(d_454_v0_)
                d_515_v51_: _dafny.Map
                d_515_v51_ = _dafny.Map({(d_514_v50_).cf27: p1})
                index72_ = default__.safeIndex(451, (d_454_v0_).length(0))
                rhs61_ = d_511_v47_
                rhs62_ = (d_512_v48_) + ((d_512_v48_).set(default__.safeIndex((self).f12, len(d_512_v48_)), (d_512_v48_)[default__.safeIndex((d_510_v46_).f13, len(d_512_v48_))]))
                rhs63_ = _dafny.Map({(self).f12: default__.fm25(((d_515_v51_)[d_454_v0_] if (d_454_v0_) in (d_515_v51_) else p0), 185, globalState)})
                rhs64_ = True
                lhs40_ = d_454_v0_
                lhs41_ = default__.safeIndex(451, (d_454_v0_).length(0))
                d_511_v47_ = rhs61_
                d_512_v48_ = rhs62_
                d_513_v49_ = rhs63_
                lhs40_[lhs41_] = rhs64_
                r0 = ((self).f12) + (len((d_481_v26_) + (d_481_v26_)))
                d_516_v52_: _dafny.MultiSet
                d_516_v52_ = _dafny.MultiSet([(d_510_v46_).f13])
                d_517_v53_: _dafny.Seq
                d_517_v53_ = _dafny.SeqWithoutIsStrInference([(d_516_v52_).isdisjoint(_dafny.MultiSet([(d_510_v46_).f13])), p1, p0, (self).f11])
                d_517_v53_ = (d_517_v53_ if False else _dafny.SeqWithoutIsStrInference([p1, d_511_v47_, p1]))
            elif True:
                d_518_v54_: _dafny.Map
                d_518_v54_ = _dafny.Map({(self).f12: (d_481_v26_) < (d_481_v26_)})
                d_519_v55_: _dafny.MultiSet
                d_519_v55_ = _dafny.MultiSet([d_455_v1_, d_455_v1_, _dafny.SeqWithoutIsStrInference([-84])])
                d_518_v54_ = (d_518_v54_).set(((d_519_v55_)[d_455_v1_] if (d_455_v1_) in (d_519_v55_) else (self).f12), (self).f11)
                d_520_v56_: bool
                d_521_v57_: bool
                d_522_v58_: int
                out26_: bool
                out27_: bool
                out28_: int
                out26_, out27_, out28_ = default__.m0((self).f12, globalState)
                d_520_v56_ = out26_
                d_521_v57_ = out27_
                d_522_v58_ = out28_
                d_520_v56_ = ((d_489_v33_ if d_521_v57_ else d_489_v33_)).ispropersubset(d_489_v33_)
                d_454_v0_ = d_454_v0_
                d_523_v59_: bool
                d_524_v60_: bool
                d_525_v61_: int
                out29_: bool
                out30_: bool
                out31_: int
                out29_, out30_, out31_ = default__.m0((self).f12, globalState)
                d_523_v59_ = out29_
                d_524_v60_ = out30_
                d_525_v61_ = out31_
            d_526_v62_: D8
            d_526_v62_ = D8_DC15(d_454_v0_)
            pat_let_tv14_ = d_454_v0_
            def iife28_(_pat_let4_0):
                def iife29_(d_527_dt__update__tmp_h0_):
                    def iife30_(_pat_let5_0):
                        def iife31_(d_528_dt__update_hcf27_h0_):
                            return D8_DC15(d_528_dt__update_hcf27_h0_)
                        return iife31_(_pat_let5_0)
                    return iife30_(pat_let_tv14_)
                return iife29_(_pat_let4_0)
            source10_ = iife28_(d_526_v62_)
            if source10_.is_DC16:
                d_529___mcc_h4_ = source10_.cf28
                d_530___mcc_h5_ = source10_.cf29
                d_531___mcc_h6_ = source10_.cf30
                d_532___mcc_h7_ = source10_.cf31
                d_533_cf31_ = d_532___mcc_h7_
                d_534_cf30_ = d_531___mcc_h6_
                d_535_cf29_ = d_530___mcc_h5_
                d_536_cf28_ = d_529___mcc_h4_
                index73_ = default__.safeIndex(406, (d_454_v0_).length(0))
                (d_454_v0_)[index73_] = (d_489_v33_).issubset(d_489_v33_)
                index74_ = default__.safeIndex(406, (d_454_v0_).length(0))
                (d_454_v0_)[index74_] = False
                d_537_v63_: _dafny.MultiSet
                d_537_v63_ = _dafny.MultiSet([d_533_cf31_, len(d_481_v26_)])
                (globalState).f2 = ((d_536_cf28_) + ((d_537_v63_).cardinality)) * (d_533_cf31_)
                d_538_v64_: _dafny.Set
                d_538_v64_ = _dafny.Set({p0})
                d_539_v65_: _dafny.Array
                nw88_ = _dafny.Array(None, 2)
                nw88_[int(0)] = d_533_cf31_
                nw88_[int(1)] = default__.safeModuloInt(d_536_cf28_, len(d_538_v64_))
                d_539_v65_ = nw88_
                index75_ = default__.safeIndex(202, (d_539_v65_).length(0))
                (d_539_v65_)[index75_] = d_536_cf28_
                index76_ = default__.safeIndex(202, (d_539_v65_).length(0))
                (d_539_v65_)[index76_] = -366
                d_540_v66_: D8
                d_540_v66_ = D8_DC16(d_533_cf31_, p0, (self).f11, d_536_cf28_)
                d_535_cf29_ = (d_540_v66_).cf29
            elif source10_.is_DC15:
                d_541___mcc_h8_ = source10_.cf27
                d_542_cf27_ = d_541___mcc_h8_
                d_543_v67_: T0
                nw89_ = C1()
                nw89_.ctor__((d_455_v1_)[default__.safeIndex((self).f12, len(d_455_v1_))], (self).f3)
                d_543_v67_ = nw89_
                d_543_v67_ = d_543_v67_
                d_544_v68_: int
                d_545_v69_: int
                d_546_v70_: bool
                out32_: int
                out33_: int
                out34_: bool
                out32_, out33_, out34_ = (self).m6(globalState)
                d_544_v68_ = out32_
                d_545_v69_ = out33_
                d_546_v70_ = out34_
                r1 = (d_481_v26_) + ((d_481_v26_) + (d_481_v26_))
                r1 = d_481_v26_
            elif True:
                d_547___mcc_h9_ = source10_.cf32
                d_548_cf32_ = d_547___mcc_h9_
                d_549_v71_: bool
                d_549_v71_ = False
                d_550_v72_: C0
                nw90_ = C0()
                nw90_.ctor__((self).f12)
                d_550_v72_ = nw90_
                d_551_v73_: D4
                d_551_v73_ = D4_DC8(d_550_v72_)
                rhs65_ = p0
                rhs66_ = p0
                rhs67_ = d_551_v73_
                d_549_v71_ = rhs65_
                d_549_v71_ = rhs66_
                d_551_v73_ = rhs67_
                d_552_v74_: _dafny.MultiSet
                d_552_v74_ = _dafny.MultiSet([d_549_v71_])
                d_553_v75_: _dafny.Array
                nw91_ = _dafny.Array(int(0), 23)
                d_553_v75_ = nw91_
                index77_ = default__.safeIndex(685, (d_553_v75_).length(0))
                (d_553_v75_)[index77_] = (self).f12
                d_554_v76_: _dafny.Map
                d_554_v76_ = _dafny.Map({d_549_v71_: d_552_v74_})
                d_555_v77_: _dafny.Map
                d_555_v77_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeosvtnfb")): ((d_554_v76_)[(self).f11] if ((self).f11) in (d_554_v76_) else d_552_v74_)})
                index78_ = default__.safeIndex(685, (d_553_v75_).length(0))
                rhs68_ = len(((d_481_v26_) + (d_481_v26_)) + (((d_481_v26_) + (d_481_v26_)).set(default__.safeIndex((d_550_v72_).f13, len((d_481_v26_) + (d_481_v26_))), d_486_v30_)))
                rhs69_ = default__.safeModuloInt((0) - ((d_550_v72_).f13), (d_550_v72_).f13)
                rhs70_ = ((d_555_v77_)[(d_481_v26_ if default__.fm23(d_481_v26_, (d_550_v72_).f13, (d_550_v72_).f13, d_549_v71_, globalState) else (d_481_v26_).set(default__.safeIndex((default__.fm30((d_550_v72_).f13, (self).f12, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkitllpy")), globalState)).cf7, len(d_481_v26_)), _dafny.CodePoint('f')))] if ((d_481_v26_ if default__.fm23(d_481_v26_, (d_550_v72_).f13, (d_550_v72_).f13, d_549_v71_, globalState) else (d_481_v26_).set(default__.safeIndex((default__.fm30((d_550_v72_).f13, (self).f12, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkitllpy")), globalState)).cf7, len(d_481_v26_)), _dafny.CodePoint('f')))) in (d_555_v77_) else d_552_v74_)
                rhs71_ = (self).f12
                lhs42_ = globalState
                lhs43_ = globalState
                lhs44_ = d_553_v75_
                lhs45_ = default__.safeIndex(685, (d_553_v75_).length(0))
                lhs42_.f2 = rhs68_
                lhs43_.f2 = rhs69_
                d_552_v74_ = rhs70_
                lhs44_[lhs45_] = rhs71_
                d_556_v78_: C0
                nw92_ = C0()
                nw92_.ctor__((d_455_v1_)[default__.safeIndex((d_550_v72_).f13, len(d_455_v1_))])
                d_556_v78_ = nw92_
                d_557_v79_: _dafny.Seq
                d_557_v79_ = _dafny.SeqWithoutIsStrInference([not(d_549_v71_)])
                d_558_v80_: _dafny.Seq
                d_558_v80_ = d_557_v79_
                d_559_v81_: _dafny.Map
                d_559_v81_ = _dafny.Map({d_558_v80_: p1})
                d_559_v81_ = (d_559_v81_).set(d_558_v80_, (self).f11)
            d_560_v82_: _dafny.Set
            d_560_v82_ = _dafny.Set({False})
            d_560_v82_ = (default__.fm31(p1, (self).f12, (self).f11, globalState)) - (d_560_v82_)
            d_454_v0_ = d_454_v0_
        r0 = (self).f12
        d_561_v83_: _dafny.Map
        d_561_v83_ = _dafny.Map({(self).f12: p0})
        r1 = (default__.fm25(((d_561_v83_)[(self).f12] if ((self).f12) in (d_561_v83_) else p0), len(default__.fm32(True, p1, (self).f12, (self).f11, globalState)), globalState)) + (d_481_v26_)
        r2 = len(d_481_v26_)
        return r0, r1, r2

    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12

class C3:
    def  __init__(self):
        self._f9: _dafny.MultiSet = _dafny.MultiSet({})
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f9, f10):
        (self)._f9 = f9
        (self)._f10 = f10

    def fm19(self, p0, p1, globalState):
        if not(False):
            return len(_dafny.SeqWithoutIsStrInference([550, (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([-925, 484, len(_dafny.Set({True}))])).cardinality, -771]))))]))
        elif True:
            return (0) - (-312)

    def fm20(self, p0, p1, p2, p3, globalState):
        return 483

    def m5(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: str = _dafny.CodePoint('D')
        r3: bool = False
        d_562_v0_: _dafny.Array
        nw93_ = _dafny.Array(_dafny.Map({}), 5)
        d_562_v0_ = nw93_
        d_563_v1_: _dafny.Map
        d_563_v1_ = _dafny.Map({p1: p1})
        d_564_v2_: _dafny.Array
        nw94_ = _dafny.Array(None, 3)
        nw94_[int(0)] = p0
        nw94_[int(1)] = p0
        nw94_[int(2)] = len(d_563_v1_)
        d_564_v2_ = nw94_
        d_565_v3_: _dafny.Map
        d_565_v3_ = _dafny.Map({d_564_v2_: d_564_v2_})
        index79_ = default__.safeIndex(803, (d_562_v0_).length(0))
        (d_562_v0_)[index79_] = (d_565_v3_) | (d_565_v3_)
        index80_ = default__.safeIndex(803, (d_562_v0_).length(0))
        (d_562_v0_)[index80_] = d_565_v3_
        d_566_v5_: _dafny.Seq
        d_566_v5_ = _dafny.SeqWithoutIsStrInference([(self).f9])
        d_567_v6_: D2
        def iife32_():
            coll20_ = _dafny.Map()
            compr_20_: _dafny.MultiSet
            for compr_20_ in (d_566_v5_).Elements:
                d_568_v4_: _dafny.MultiSet = compr_20_
                if (d_568_v4_) in (d_566_v5_):
                    coll20_[d_568_v4_] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0, p0]))).cardinality
            return _dafny.Map(coll20_)
        d_567_v6_ = D2_DC3(iife32_()
)
        source11_ = d_567_v6_
        if source11_.is_DC4:
            d_569___mcc_h0_ = source11_.cf6
            d_570_cf6_ = d_569___mcc_h0_
            d_571_v7_: _dafny.Map
            d_571_v7_ = _dafny.Map({(p1) and (p1): (p0) - (p0)})
            d_571_v7_ = (d_571_v7_).set(p1, d_570_cf6_)
            d_572_v8_: _dafny.Seq
            d_572_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfvoqiom"))
            d_573_v9_: _dafny.Set
            d_573_v9_ = _dafny.Set({p0, p0, p0})
            d_574_v10_: _dafny.Seq
            d_574_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(0) - (len(d_572_v8_)), p0, 30}), d_573_v9_])
            d_575_v11_: str
            d_575_v11_ = _dafny.CodePoint('k')
            d_576_v12_: _dafny.Map
            d_576_v12_ = _dafny.Map({p0: d_564_v2_})
            d_577_v13_: _dafny.Seq
            d_577_v13_ = _dafny.SeqWithoutIsStrInference([d_570_cf6_])
            d_578_v14_: _dafny.Seq
            d_578_v14_ = _dafny.SeqWithoutIsStrInference([(self).f10, p1])
            d_579_v15_: D2
            d_579_v15_ = D2_DC5(p0, (self).f10, p1, (self).f10, (self).f10)
            d_580_v17_: _dafny.Array
            nw95_ = _dafny.Array(None, 27)
            nw95_[int(0)] = default__.fm21(len(d_574_v10_), d_575_v11_, default__.fm21((0) - (len((d_576_v12_).set((self).fm19(d_575_v11_, (d_577_v13_)[default__.safeIndex(len(d_578_v14_), len(d_577_v13_))], globalState), d_564_v2_))), d_575_v11_, p1, (self).f10, globalState), p1, globalState)
            nw95_[int(1)] = (d_573_v9_).issubset(d_573_v9_)
            nw95_[int(2)] = ((self).f9).ispropersubset(_dafny.MultiSet(d_577_v13_))
            nw95_[int(3)] = (self).f10
            nw95_[int(4)] = p1
            nw95_[int(5)] = default__.fm21(len(d_572_v8_), d_575_v11_, (self).f10, False, globalState)
            nw95_[int(6)] = p1
            nw95_[int(7)] = (self).f10
            nw95_[int(8)] = (p1 if False else (self).f10)
            nw95_[int(9)] = p1
            nw95_[int(10)] = p1
            nw95_[int(11)] = (d_579_v15_).cf11
            nw95_[int(12)] = (True if (self).f10 else (self).f10)
            nw95_[int(13)] = not((self).f10)
            nw95_[int(14)] = (self).f10
            nw95_[int(15)] = p1
            nw95_[int(16)] = False
            nw95_[int(17)] = (self).f10
            nw95_[int(18)] = ((self).f10) == (not((self).f10))
            nw95_[int(19)] = p1
            nw95_[int(20)] = (p1) and ((self).f10)
            nw95_[int(21)] = p1
            nw95_[int(22)] = False
            def iife33_():
                coll21_ = _dafny.Set()
                compr_21_: int
                for compr_21_ in (d_577_v13_).Elements:
                    d_581_v16_: int = compr_21_
                    if (d_581_v16_) in (d_577_v13_):
                        coll21_ = coll21_.union(_dafny.Set([default__.safeModuloInt(d_581_v16_, len(_dafny.SeqWithoutIsStrInference([121 for d_582_i0_ in range(default__.abs(994))])))]))
                return _dafny.Set(coll21_)
            nw95_[int(23)] = (_dafny.Set({p0, d_570_cf6_})).ispropersubset(iife33_()
            )
            nw95_[int(24)] = p1
            nw95_[int(25)] = (d_577_v13_) == (d_577_v13_)
            nw95_[int(26)] = default__.fm21(d_570_cf6_, d_575_v11_, ((d_563_v1_)[True] if (True) in (d_563_v1_) else (self).f10), p1, globalState)
            d_580_v17_ = nw95_
            index81_ = default__.safeIndex(725, (d_580_v17_).length(0))
            (d_580_v17_)[index81_] = p1
            index82_ = default__.safeIndex(725, (d_580_v17_).length(0))
            (d_580_v17_)[index82_] = (self).f10
            d_564_v2_ = d_564_v2_
            d_580_v17_ = d_580_v17_
        elif source11_.is_DC5:
            d_583___mcc_h1_ = source11_.cf7
            d_584___mcc_h2_ = source11_.cf8
            d_585___mcc_h3_ = source11_.cf9
            d_586___mcc_h4_ = source11_.cf10
            d_587___mcc_h5_ = source11_.cf11
            d_588_cf11_ = d_587___mcc_h5_
            d_589_cf10_ = d_586___mcc_h4_
            d_590_cf9_ = d_585___mcc_h3_
            d_591_cf8_ = d_584___mcc_h2_
            d_592_cf7_ = d_583___mcc_h1_
            d_588_cf11_ = (p0) != ((d_592_cf7_) - (d_592_cf7_))
            d_593_v18_: _dafny.Seq
            d_593_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsovhibx"))
            d_594_v19_: _dafny.Seq
            d_594_v19_ = _dafny.SeqWithoutIsStrInference([270, len(d_593_v18_)])
            d_595_v20_: str
            d_595_v20_ = _dafny.CodePoint('q')
            d_596_v21_: _dafny.Seq
            d_596_v21_ = _dafny.SeqWithoutIsStrInference([d_591_cf8_])
            r1 = (default__.fm21(p0, d_595_v20_, (d_596_v21_)[default__.safeIndex(d_592_cf7_, len(d_596_v21_))], d_589_cf10_, globalState) if (default__.fm22(d_594_v19_, d_592_cf7_, d_595_v20_, globalState)) not in (_dafny.Map({d_567_v6_: p1})) else default__.fm21(d_592_cf7_, d_595_v20_, (self).f10, d_589_cf10_, globalState))
            d_597_v22_: _dafny.Map
            d_597_v22_ = _dafny.Map({d_592_cf7_: d_564_v2_})
            d_597_v22_ = ((d_597_v22_) | ((d_597_v22_).set(963, d_564_v2_))) | (_dafny.Map({p0: d_564_v2_}))
            d_598_v23_: _dafny.Map
            d_598_v23_ = _dafny.Map({d_589_cf10_: _dafny.Map({d_588_cf11_: -126})})
            d_599_v24_: _dafny.Array
            nw96_ = _dafny.Array(_dafny.Map({}), 4)
            d_599_v24_ = nw96_
            d_600_v25_: C1
            nw97_ = C1()
            nw97_.ctor__((0) - (len(d_598_v23_)), d_599_v24_)
            d_600_v25_ = nw97_
        elif True:
            d_601___mcc_h6_ = source11_.cf5
            d_602_cf5_ = d_601___mcc_h6_
            d_603_v26_: _dafny.Seq
            d_603_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnqpg"))
            d_604_v27_: str
            d_604_v27_ = _dafny.CodePoint('b')
            d_605_v28_: D4
            d_605_v28_ = D4_DC9(991, d_604_v27_, _dafny.SeqWithoutIsStrInference([141 for d_606_i1_ in range(default__.abs(119))]), p0, (self).f10)
            d_607_v29_: _dafny.Seq
            d_607_v29_ = _dafny.SeqWithoutIsStrInference([p0, -848])
            d_608_v30_: _dafny.Seq
            d_608_v30_ = _dafny.SeqWithoutIsStrInference([len((d_603_v26_) + (default__.fm25((d_605_v28_).cf20, len(d_607_v29_), globalState)))])
            (globalState).f2 = (d_607_v29_)[default__.safeIndex(p0, len(d_607_v29_))]
            (globalState).f2 = p0
            d_609_v31_: _dafny.Array
            nw98_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
            d_609_v31_ = nw98_
            index83_ = default__.safeIndex(842, (d_609_v31_).length(0))
            (d_609_v31_)[index83_] = default__.fm25((self).f10, 444, globalState)
            index84_ = default__.safeIndex(842, (d_609_v31_).length(0))
            (d_609_v31_)[index84_] = _dafny.SeqWithoutIsStrInference([d_604_v27_ for d_610_i2_ in range(default__.abs(641))])
            r3 = p1
        d_611_v32_: _dafny.Array
        nw99_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
        d_611_v32_ = nw99_
        d_612_v33_: _dafny.Map
        d_612_v33_ = _dafny.Map({d_611_v32_: p1})
        r0 = ((d_612_v33_)[d_611_v32_] if (d_611_v32_) in (d_612_v33_) else (self).f10)
        d_613_v34_: _dafny.Array
        def lambda18_(d_614_i3_):
            return _dafny.Map({_dafny.CodePoint('r'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvkrqon"))})

        init11_ = lambda18_
        nw100_ = _dafny.Array(None, 6)
        for i0_11_ in range(nw100_.length(0)):
            nw100_[i0_11_] = init11_(i0_11_)
        d_613_v34_ = nw100_
        d_615_v35_: str
        d_615_v35_ = _dafny.CodePoint('t')
        d_616_v36_: _dafny.Seq
        d_616_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aewdkgg"))
        d_617_v37_: _dafny.Map
        d_617_v37_ = _dafny.Map({d_615_v35_: default__.fm25(False, len(d_616_v36_), globalState)})
        index85_ = default__.safeIndex(975, (d_613_v34_).length(0))
        (d_613_v34_)[index85_] = d_617_v37_
        d_618_v38_: _dafny.Map
        d_618_v38_ = _dafny.Map({(self).f9: (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f10, p1, not((self).f10), (self).f10, p1])))})
        pat_let_tv15_ = d_617_v37_
        pat_let_tv16_ = d_615_v35_
        pat_let_tv17_ = d_616_v36_
        pat_let_tv18_ = d_617_v37_
        pat_let_tv19_ = d_617_v37_
        index86_ = default__.safeIndex(975, (d_613_v34_).length(0))
        def lambda19_(source12_):
            if source12_.is_DC4:
                d_619___mcc_h7_ = source12_.cf6
                d_620_cf6_ = d_619___mcc_h7_
                return pat_let_tv15_
            elif source12_.is_DC5:
                d_621___mcc_h8_ = source12_.cf7
                d_622___mcc_h9_ = source12_.cf8
                d_623___mcc_h10_ = source12_.cf9
                d_624___mcc_h11_ = source12_.cf10
                d_625___mcc_h12_ = source12_.cf11
                d_626_cf11_ = d_625___mcc_h12_
                d_627_cf10_ = d_624___mcc_h11_
                d_628_cf9_ = d_623___mcc_h10_
                d_629_cf8_ = d_622___mcc_h9_
                d_630_cf7_ = d_621___mcc_h8_
                return (_dafny.Map({pat_let_tv16_: pat_let_tv17_})) | (pat_let_tv18_)
            elif True:
                d_631___mcc_h13_ = source12_.cf5
                d_632_cf5_ = d_631___mcc_h13_
                return pat_let_tv19_

        (d_613_v34_)[index86_] = lambda19_(D2_DC3(d_618_v38_))
        d_633_v40_: _dafny.Array
        def lambda20_(d_634_p0_):
            def lambda21_(d_635_i4_):
                def iife34_():
                    coll22_ = _dafny.Map()
                    compr_22_: int
                    for compr_22_ in (_dafny.Set({d_634_p0_})).Elements:
                        d_636_v39_: int = compr_22_
                        if (d_636_v39_) in (_dafny.Set({d_634_p0_})):
                            coll22_[(d_636_v39_) - (d_634_p0_)] = (self).f10
                    return _dafny.Map(coll22_)
                return iife34_()
                

            return lambda21_

        init12_ = lambda20_(p0)
        nw101_ = _dafny.Array(None, 15)
        for i0_12_ in range(nw101_.length(0)):
            nw101_[i0_12_] = init12_(i0_12_)
        d_633_v40_ = nw101_
        d_637_v41_: C2
        nw102_ = C2()
        nw102_.ctor__((self).f10, p0, d_633_v40_)
        d_637_v41_ = nw102_
        d_638_v42_: _dafny.Array
        nw103_ = _dafny.Array(_dafny.Array(None, 0), 10)
        d_638_v42_ = nw103_
        d_638_v42_ = d_638_v42_
        r0 = p1
        r1 = (d_637_v41_).f11
        d_639_v43_: _dafny.Array
        def lambda22_(d_640_p0_, d_641_v41_):
            def lambda23_(d_642_i5_):
                return _dafny.SeqWithoutIsStrInference([d_640_p0_, len(_dafny.SeqWithoutIsStrInference([(d_641_v41_).f11, (self).f10])), (d_641_v41_).f12])

            return lambda23_

        init13_ = lambda22_(p0, d_637_v41_)
        nw104_ = _dafny.Array(None, 22)
        for i0_13_ in range(nw104_.length(0)):
            nw104_[i0_13_] = init13_(i0_13_)
        d_639_v43_ = nw104_
        d_643_v44_: _dafny.Map
        d_643_v44_ = _dafny.Map({d_639_v43_: (d_637_v41_).f12})
        d_644_v45_: _dafny.Map
        d_644_v45_ = _dafny.Map({p1: (d_637_v41_).f12})
        d_645_v46_: _dafny.Seq
        d_645_v46_ = _dafny.SeqWithoutIsStrInference([p1, (d_637_v41_).f11, (d_637_v41_).f11, (self).f10, (d_637_v41_).f11])
        d_646_v47_: _dafny.Seq
        d_646_v47_ = _dafny.SeqWithoutIsStrInference([len(d_645_v46_), p0])
        d_647_v48_: D4
        d_647_v48_ = D4_DC9(((d_643_v44_)[d_639_v43_] if (d_639_v43_) in (d_643_v44_) else ((d_644_v45_)[(d_637_v41_).f11] if ((d_637_v41_).f11) in (d_644_v45_) else 273)), d_615_v35_, d_646_v47_, p0, (d_637_v41_).f11)
        r2 = (d_647_v48_).cf17
        r3 = ((self).fm19(d_615_v35_, (d_637_v41_).f12, globalState)) <= (636)
        return r0, r1, r2, r3

    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10

class C4(T1, T0):
    def  __init__(self):
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f7: bool = False
        self._f8: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f7, f8, f3):
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f3 = f3

    def fm13(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Set({196, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yapel"))), len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruu")))])]))}) if True else _dafny.Set({532, 424}))).intersection(_dafny.Set({566}))

    def fm14(self, p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aljqjwu"))

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: str = _dafny.CodePoint('D')
        d_648_v0_: int
        d_648_v0_ = 671
        d_649_v1_: D1
        d_649_v1_ = D1_DC2(d_648_v0_, p1, d_648_v0_)
        d_650_v2_: _dafny.Array
        nw105_ = _dafny.Array(None, 10)
        nw105_[int(0)] = (self).f7
        nw105_[int(1)] = (d_648_v0_) >= (d_648_v0_)
        nw105_[int(2)] = (d_649_v1_).cf3
        nw105_[int(3)] = not(p1)
        nw105_[int(4)] = (len(p0)) > (d_648_v0_)
        nw105_[int(5)] = p1
        nw105_[int(6)] = (d_649_v1_).cf3
        nw105_[int(7)] = (d_648_v0_) <= (d_648_v0_)
        nw105_[int(8)] = (self).f7
        nw105_[int(9)] = (self).f7
        d_650_v2_ = nw105_
        d_651_v3_: _dafny.Array
        nw106_ = _dafny.Array(int(0), 22)
        d_651_v3_ = nw106_
        index87_ = default__.safeIndex(414, (d_651_v3_).length(0))
        (d_651_v3_)[index87_] = d_648_v0_
        d_652_v6_: _dafny.Set
        d_652_v6_ = _dafny.Set({d_648_v0_, len(p2)})
        d_653_v7_: _dafny.Map
        def iife35_():
            coll23_ = _dafny.Map()
            compr_23_: int
            for compr_23_ in (d_652_v6_).Elements:
                d_654_v5_: int = compr_23_
                if (d_654_v5_) in (d_652_v6_):
                    coll23_[(d_654_v5_) * (len(_dafny.Map({_dafny.CodePoint('t'): (self).f7})))] = True
            return _dafny.Map(coll23_)
        d_653_v7_ = _dafny.Map({iife35_()
        : (self).f7})
        d_655_v8_: _dafny.Map
        d_655_v8_ = _dafny.Map({d_648_v0_: p1})
        index88_ = default__.safeIndex(414, (d_651_v3_).length(0))
        def iife36_():
            coll24_ = _dafny.Map()
            compr_24_: _dafny.Map
            for compr_24_ in (d_653_v7_).keys.Elements:
                d_656_v4_: _dafny.Map = compr_24_
                if (d_656_v4_) in (d_653_v7_):
                    coll24_[d_656_v4_] = d_648_v0_
            return _dafny.Map(coll24_)
        rhs72_ = d_650_v2_
        rhs73_ = d_648_v0_
        rhs74_ = len((iife36_()
        ).set(d_655_v8_, (d_648_v0_) * (d_648_v0_)))
        lhs46_ = d_651_v3_
        lhs47_ = default__.safeIndex(414, (d_651_v3_).length(0))
        lhs48_ = globalState
        d_650_v2_ = rhs72_
        lhs46_[lhs47_] = rhs73_
        lhs48_.f2 = rhs74_
        d_657_i0_: int
        d_657_i0_ = 0
        with _dafny.label("2"):
            while p1:
                with _dafny.c_label("2"):
                    if (d_657_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_657_i0_ = (d_657_i0_) + (1)
                    d_658_v9_: _dafny.Map
                    d_658_v9_ = _dafny.Map({default__.fm15(not(p1), True, globalState): d_648_v0_})
                    d_659_v10_: D2
                    d_659_v10_ = D2_DC3(d_658_v9_)
                    d_660_v11_: _dafny.Map
                    d_660_v11_ = _dafny.Map({(0) - (d_648_v0_): 226})
                    d_661_v12_: _dafny.Seq
                    d_661_v12_ = _dafny.SeqWithoutIsStrInference([d_650_v2_])
                    d_662_v13_: _dafny.Array
                    nw107_ = _dafny.Array(None, 8)
                    nw107_[int(0)] = D2_DC3(d_658_v9_)
                    nw107_[int(1)] = d_659_v10_
                    nw107_[int(2)] = D2_DC3((_dafny.Map({(self).f8: len(_dafny.SeqWithoutIsStrInference([((_dafny.MultiSet([(d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], -947])).set(len(p2), default__.abs(d_648_v0_))).cardinality for d_663_i1_ in range(default__.abs(538))]))})).set((self).f8, ((d_660_v11_)[len(d_661_v12_)] if (len(d_661_v12_)) in (d_660_v11_) else d_648_v0_)))
                    nw107_[int(3)] = d_659_v10_
                    nw107_[int(4)] = d_659_v10_
                    nw107_[int(5)] = d_659_v10_
                    nw107_[int(6)] = d_659_v10_
                    nw107_[int(7)] = d_659_v10_
                    d_662_v13_ = nw107_
                    index89_ = default__.safeIndex(780, (d_662_v13_).length(0))
                    (d_662_v13_)[index89_] = d_659_v10_
                    index90_ = default__.safeIndex(780, (d_662_v13_).length(0))
                    (d_662_v13_)[index90_] = d_659_v10_
                    d_664_v14_: _dafny.Seq
                    d_664_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
                    d_664_v14_ = (d_664_v14_) + ((p2) + (p2))
                    if (self).f7:
                        d_665_v15_: bool
                        d_665_v15_ = False
                        def iife37_():
                            coll25_ = _dafny.Set()
                            compr_25_: int
                            for compr_25_ in (d_652_v6_).Elements:
                                d_666_v16_: int = compr_25_
                                if (d_666_v16_) in (d_652_v6_):
                                    coll25_ = coll25_.union(_dafny.Set([(d_666_v16_) - (713)]))
                            return _dafny.Set(coll25_)
                        d_665_v15_ = ((d_652_v6_).intersection(iife37_()
                        )).isdisjoint(d_652_v6_)
                        d_665_v15_ = p1
                        d_667_v17_: _dafny.Seq
                        d_667_v17_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                        d_668_v18_: _dafny.Map
                        d_668_v18_ = _dafny.Map({(0) - (d_648_v0_): d_667_v17_})
                        d_668_v18_ = (d_668_v18_).set((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], (_dafny.SeqWithoutIsStrInference([(self).f7, d_665_v15_])) + (d_667_v17_))
                        d_665_v15_ = d_665_v15_
                        index91_ = default__.safeIndex(414, (d_651_v3_).length(0))
                        (d_651_v3_)[index91_] = d_648_v0_
                    elif True:
                        d_669_v19_: D3
                        d_669_v19_ = D3_DC6(_dafny.CodePoint('l'))
                        d_648_v0_ = (p0)[default__.safeIndex(len(default__.fm16((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], (d_669_v19_).cf12, (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], globalState)), len(p0))]
                        d_670_v20_: _dafny.Seq
                        d_670_v20_ = _dafny.SeqWithoutIsStrInference([p2, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')])])
                        (globalState).f2 = (default__.fm17(_dafny.SeqWithoutIsStrInference([d_664_v14_]), (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], len(p0), (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], globalState)) + (default__.fm17(d_670_v20_, d_648_v0_, (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], d_648_v0_, globalState))
                        d_671_v21_: _dafny.Array
                        nw108_ = _dafny.Array(D2.default()(), 22)
                        d_671_v21_ = nw108_
                        d_671_v21_ = d_671_v21_
                        d_672_v22_: _dafny.Map
                        d_672_v22_ = _dafny.Map({False: True})
                        d_673_v23_: _dafny.Seq
                        d_673_v23_ = _dafny.SeqWithoutIsStrInference([(d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], (len(d_672_v22_)) * ((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]), (0) - (-212), default__.safeModuloInt(571, d_648_v0_)])
                        d_674_v24_: _dafny.Set
                        d_674_v24_ = _dafny.Set({d_650_v2_})
                        rhs75_ = (0) - (default__.safeModuloInt((0) - (len(d_664_v14_)), d_648_v0_))
                        rhs76_ = d_648_v0_
                        rhs77_ = default__.fm18(len(d_674_v24_), len(p0), globalState)
                        rhs78_ = (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]
                        rhs79_ = p0
                        lhs49_ = globalState
                        lhs50_ = globalState
                        d_648_v0_ = rhs75_
                        lhs49_.f2 = rhs76_
                        d_649_v1_ = rhs77_
                        lhs50_.f2 = rhs78_
                        d_673_v23_ = rhs79_
                        d_675_v25_: _dafny.Array
                        nw109_ = _dafny.Array(_dafny.Array(None, 0), 16)
                        d_675_v25_ = nw109_
                        index92_ = default__.safeIndex(306, (d_675_v25_).length(0))
                        (d_675_v25_)[index92_] = d_651_v3_
                        index93_ = default__.safeIndex(306, (d_675_v25_).length(0))
                        (d_675_v25_)[index93_] = d_651_v3_
                    d_676_v26_: _dafny.Set
                    d_676_v26_ = _dafny.Set({True})
                    d_677_v27_: _dafny.Map
                    d_677_v27_ = _dafny.Map({d_676_v26_: len(d_664_v14_)})
                    d_677_v27_ = (d_677_v27_).set(d_676_v26_, 680)
                    pass
            pass
        d_678_i2_: int
        d_678_i2_ = 0
        with _dafny.label("3"):
            while p1:
                with _dafny.c_label("3"):
                    if (d_678_i2_) >= (100):
                        raise _dafny.Break("3")
                    d_678_i2_ = (d_678_i2_) + (1)
                    (globalState).f2 = (0) - (len(((d_655_v8_ if p1 else (d_655_v8_).set((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], (self).f7))) | (d_655_v8_)))
                    d_679_v28_: C0
                    nw110_ = C0()
                    nw110_.ctor__((default__.fm17(_dafny.SeqWithoutIsStrInference([p2 for d_680_i3_ in range(default__.abs(-339))]), (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))], d_648_v0_, globalState) if p1 else (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]))
                    d_679_v28_ = nw110_
                    d_681_v29_: C3
                    nw111_ = C3()
                    nw111_.ctor__((self).f8, (self).f7)
                    d_681_v29_ = nw111_
                    (globalState).f2 = d_648_v0_
                    pass
            pass
        d_682_v30_: bool
        d_682_v30_ = False
        d_682_v30_ = p1
        d_683_i4_: int
        d_683_i4_ = 0
        with _dafny.label("4"):
            while (p0) != (p0):
                with _dafny.c_label("4"):
                    if (d_683_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_683_i4_ = (d_683_i4_) + (1)
                    d_684_v31_: str
                    d_684_v31_ = _dafny.CodePoint('x')
                    rhs80_ = (d_684_v31_ if d_682_v30_ else (p2)[default__.safeIndex(d_648_v0_, len(p2))])
                    rhs81_ = (d_648_v0_) <= (default__.safeModuloInt((0) - (d_648_v0_), d_648_v0_))
                    rhs82_ = (((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]) * (d_648_v0_)) == ((d_648_v0_ if p1 else d_648_v0_))
                    r1 = rhs80_
                    d_682_v30_ = rhs81_
                    d_682_v30_ = rhs82_
                    d_685_v32_: _dafny.Array
                    def lambda24_(d_686_v30_, d_687_p1_):
                        def lambda25_(d_688_i5_):
                            return (_dafny.Set({d_686_v30_, not((self).f7), d_687_p1_})) - (_dafny.Set({d_687_p1_}))

                        return lambda25_

                    init14_ = lambda24_(d_682_v30_, p1)
                    nw112_ = _dafny.Array(None, 16)
                    for i0_14_ in range(nw112_.length(0)):
                        nw112_[i0_14_] = init14_(i0_14_)
                    d_685_v32_ = nw112_
                    index94_ = default__.safeIndex(670, (d_685_v32_).length(0))
                    (d_685_v32_)[index94_] = _dafny.Set({True, p1, (self).f7, not(d_682_v30_)})
                    d_689_v33_: _dafny.Set
                    d_689_v33_ = _dafny.Set({(self).f7})
                    index95_ = default__.safeIndex(670, (d_685_v32_).length(0))
                    (d_685_v32_)[index95_] = d_689_v33_
                    d_682_v30_ = ((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]) > (d_648_v0_)
                    d_690_v34_: D2
                    d_690_v34_ = D2_DC4(default__.safeModuloInt(d_648_v0_, d_648_v0_))
                    source13_ = d_690_v34_
                    if source13_.is_DC4:
                        d_691___mcc_h0_ = source13_.cf6
                        d_692_cf6_ = d_691___mcc_h0_
                        d_684_v31_ = d_684_v31_
                        d_651_v3_ = d_651_v3_
                        d_693_v35_: _dafny.Map
                        d_693_v35_ = _dafny.Map({(d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]: 169})
                        d_694_v36_: _dafny.Map
                        d_694_v36_ = _dafny.Map({(((d_693_v35_)[d_648_v0_] if (d_648_v0_) in (d_693_v35_) else default__.fm24(d_652_v6_, globalState))) + ((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]): (d_692_cf6_) * (d_692_cf6_)})
                        d_695_v37_: _dafny.MultiSet
                        d_695_v37_ = _dafny.MultiSet([(self).fm13(d_648_v0_, 495, p1, (0) - (d_692_cf6_), globalState)])
                        d_693_v35_ = (d_693_v35_).set(d_692_cf6_, ((d_695_v37_)[_dafny.Set({d_692_cf6_})] if (_dafny.Set({d_692_cf6_})) in (d_695_v37_) else len(p2)))
                        d_682_v30_ = default__.fm23((p2 if not(p1) else p2), d_648_v0_, d_692_cf6_, d_682_v30_, globalState)
                    elif source13_.is_DC5:
                        d_696___mcc_h1_ = source13_.cf7
                        d_697___mcc_h2_ = source13_.cf8
                        d_698___mcc_h3_ = source13_.cf9
                        d_699___mcc_h4_ = source13_.cf10
                        d_700___mcc_h5_ = source13_.cf11
                        d_701_cf11_ = d_700___mcc_h5_
                        d_702_cf10_ = d_699___mcc_h4_
                        d_703_cf9_ = d_698___mcc_h3_
                        d_704_cf8_ = d_697___mcc_h2_
                        d_705_cf7_ = d_696___mcc_h1_
                        d_706_v38_: _dafny.Array
                        def lambda26_(d_707_v31_):
                            def lambda27_(d_708_i6_):
                                return d_707_v31_

                            return lambda27_

                        init15_ = lambda26_(d_684_v31_)
                        nw113_ = _dafny.Array(None, 21)
                        for i0_15_ in range(nw113_.length(0)):
                            nw113_[i0_15_] = init15_(i0_15_)
                        d_706_v38_ = nw113_
                        d_709_v39_: _dafny.Seq
                        d_709_v39_ = _dafny.SeqWithoutIsStrInference([d_706_v38_, d_706_v38_])
                        d_648_v0_ = ((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]) - (len(d_709_v39_))
                        d_710_v40_: C0
                        nw114_ = C0()
                        nw114_.ctor__(d_705_cf7_)
                        d_710_v40_ = nw114_
                        d_711_v41_: _dafny.Seq
                        d_711_v41_ = _dafny.SeqWithoutIsStrInference([d_704_cf8_])
                        d_712_v42_: _dafny.MultiSet
                        d_712_v42_ = _dafny.MultiSet([(d_711_v41_) + (d_711_v41_), (d_711_v41_).set(default__.safeIndex(720, len(d_711_v41_)), not(True)), _dafny.SeqWithoutIsStrInference([d_704_cf8_]), d_711_v41_])
                        d_713_v43_: _dafny.Map
                        d_713_v43_ = _dafny.Map({d_702_cf10_: default__.fm33(d_702_cf10_, _dafny.SeqWithoutIsStrInference([(0) - ((d_710_v40_).f13)]), d_703_cf9_, d_684_v31_, globalState)})
                        d_712_v42_ = ((d_713_v43_)[not(d_682_v30_)] if (not(d_682_v30_)) in (d_713_v43_) else d_712_v42_)
                        d_714_v44_: _dafny.Array
                        def lambda28_(d_715_cf11_):
                            def lambda29_(d_716_i7_):
                                return _dafny.SeqWithoutIsStrInference([d_715_cf11_, d_715_cf11_])

                            return lambda29_

                        init16_ = lambda28_(d_701_cf11_)
                        nw115_ = _dafny.Array(None, 2)
                        for i0_16_ in range(nw115_.length(0)):
                            nw115_[i0_16_] = init16_(i0_16_)
                        d_714_v44_ = nw115_
                        index96_ = default__.safeIndex(426, (d_714_v44_).length(0))
                        (d_714_v44_)[index96_] = d_711_v41_
                        index97_ = default__.safeIndex(426, (d_714_v44_).length(0))
                        (d_714_v44_)[index97_] = d_711_v41_
                    elif True:
                        d_717___mcc_h6_ = source13_.cf5
                        d_718_cf5_ = d_717___mcc_h6_
                        d_719_v45_: C2
                        nw116_ = C2()
                        nw116_.ctor__((self).f7, 296, ((self).f3 if p1 else (self).f3))
                        d_719_v45_ = nw116_
                        d_720_v46_: _dafny.MultiSet
                        d_720_v46_ = _dafny.MultiSet([p1, not((d_719_v45_).f11), True])
                        d_721_v47_: _dafny.Map
                        d_721_v47_ = _dafny.Map({False: d_651_v3_})
                        d_722_v48_: _dafny.Map
                        d_722_v48_ = _dafny.Map({len(p2): d_651_v3_})
                        d_723_v49_: _dafny.Array
                        nw117_ = _dafny.Array(None, 18)
                        nw117_[int(0)] = d_651_v3_
                        nw117_[int(1)] = d_651_v3_
                        nw117_[int(2)] = ((d_721_v47_)[(d_719_v45_).f11] if ((d_719_v45_).f11) in (d_721_v47_) else d_651_v3_)
                        nw117_[int(3)] = d_651_v3_
                        nw117_[int(4)] = d_651_v3_
                        nw117_[int(5)] = d_651_v3_
                        nw117_[int(6)] = d_651_v3_
                        nw117_[int(7)] = d_651_v3_
                        nw117_[int(8)] = d_651_v3_
                        nw117_[int(9)] = d_651_v3_
                        nw117_[int(10)] = d_651_v3_
                        nw117_[int(11)] = d_651_v3_
                        nw117_[int(12)] = d_651_v3_
                        nw117_[int(13)] = d_651_v3_
                        nw117_[int(14)] = d_651_v3_
                        nw117_[int(15)] = d_651_v3_
                        nw117_[int(16)] = d_651_v3_
                        nw117_[int(17)] = ((d_722_v48_)[(0) - ((d_719_v45_).f12)] if ((0) - ((d_719_v45_).f12)) in (d_722_v48_) else d_651_v3_)
                        d_723_v49_ = nw117_
                        index98_ = default__.safeIndex(914, (d_723_v49_).length(0))
                        (d_723_v49_)[index98_] = d_651_v3_
                        d_724_v50_: _dafny.Map
                        d_724_v50_ = _dafny.Map({(self).f7: (_dafny.MultiSet([p1])).intersection(d_720_v46_)})
                        index99_ = default__.safeIndex(914, (d_723_v49_).length(0))
                        rhs83_ = ((d_724_v50_)[p1] if (p1) in (d_724_v50_) else d_720_v46_)
                        rhs84_ = p1
                        rhs85_ = ((d_722_v48_)[((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]) * ((d_719_v45_).f12)] if (((d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]) * ((d_719_v45_).f12)) in (d_722_v48_) else d_651_v3_)
                        lhs51_ = d_723_v49_
                        lhs52_ = default__.safeIndex(914, (d_723_v49_).length(0))
                        d_720_v46_ = rhs83_
                        d_682_v30_ = rhs84_
                        lhs51_[lhs52_] = rhs85_
                        d_725_v51_: _dafny.Map
                        d_725_v51_ = _dafny.Map({d_648_v0_: (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]})
                        d_726_v52_: _dafny.Map
                        d_726_v52_ = _dafny.Map({d_725_v51_: ((self).f8).intersection(_dafny.MultiSet([(d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))]]))})
                        d_726_v52_ = (d_726_v52_).set(((d_725_v51_).set((d_719_v45_).f12, (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))])).set(((self).f8).cardinality, (0) - ((_dafny.MultiSet([(self).f7])).cardinality)), (self).f8)
                        rhs86_ = True
                        rhs87_ = (0) - ((default__.safeDivisionInt((d_719_v45_).f12, (d_651_v3_)[default__.safeIndex(414, (d_651_v3_).length(0))])) * (d_648_v0_))
                        rhs88_ = (d_719_v45_).f12
                        rhs89_ = (self).f7
                        rhs90_ = (default__.safeModuloInt(d_648_v0_, d_648_v0_)) > ((d_719_v45_).f12)
                        lhs53_ = globalState
                        d_682_v30_ = rhs86_
                        r0 = rhs87_
                        lhs53_.f2 = rhs88_
                        d_682_v30_ = rhs89_
                        d_682_v30_ = rhs90_
                    pass
            pass
        d_682_v30_ = ((D2_DC5(d_648_v0_, d_682_v30_, (self).f7, not(d_682_v30_), d_682_v30_)).cf7) == (len(p2))
        r0 = d_648_v0_
        d_727_v53_: str
        d_727_v53_ = _dafny.CodePoint('u')
        r1 = d_727_v53_
        return r0, r1

    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8

class C5:
    def  __init__(self):
        self._f6: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm7(self, p0, p1, p2, globalState):
        return ((976 if False else 852)) <= (len((_dafny.Map({not(False): (_dafny.MultiSet([(D1_DC2(840, True, -480)).cf3, True, False])).cardinality})) | (_dafny.Map({not(False): len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False, False])}))}))))

    def m4(self, p0, p1, globalState):
        d_728_v0_: bool
        d_728_v0_ = False
        if not(d_728_v0_):
            d_729_v1_: _dafny.Seq
            d_729_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otnenv"))
            d_728_v0_ = (d_729_v1_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sypxlsr")))
            d_730_v2_: _dafny.Array
            nw118_ = _dafny.Array(_dafny.Map({}), 11)
            d_730_v2_ = nw118_
            index100_ = default__.safeIndex(439, (d_730_v2_).length(0))
            (d_730_v2_)[index100_] = _dafny.Map({p1: p1})
            d_731_v3_: _dafny.Map
            d_731_v3_ = _dafny.Map({p0: d_728_v0_})
            d_732_v4_: _dafny.MultiSet
            d_732_v4_ = _dafny.MultiSet([p0, p0])
            d_733_v5_: _dafny.Map
            d_733_v5_ = _dafny.Map({d_732_v4_: p1})
            d_734_v6_: D2
            d_734_v6_ = D2_DC3(d_733_v5_)
            d_735_v7_: _dafny.Seq
            d_735_v7_ = _dafny.SeqWithoutIsStrInference([len((d_734_v6_).cf5), p0])
            d_736_v8_: _dafny.Map
            d_736_v8_ = _dafny.Map({(0) - ((p0) * (len(d_731_v3_))): len(d_735_v7_)})
            index101_ = default__.safeIndex(439, (d_730_v2_).length(0))
            (d_730_v2_)[index101_] = d_736_v8_
            d_737_v9_: _dafny.Map
            d_737_v9_ = _dafny.Map({not(d_728_v0_): -76})
            d_738_v10_: _dafny.MultiSet
            d_738_v10_ = _dafny.MultiSet([not(d_728_v0_)])
            d_737_v9_ = (d_737_v9_).set(not(d_728_v0_), (d_738_v10_).cardinality)
            d_739_v11_: _dafny.Array
            def lambda30_(d_740_v3_, d_741_p1_, d_742_v0_):
                def lambda31_(d_743_i0_):
                    return _dafny.Set({len((d_740_v3_).set(d_741_p1_, d_742_v0_))})

                return lambda31_

            init17_ = lambda30_(d_731_v3_, p1, d_728_v0_)
            nw119_ = _dafny.Array(None, 2)
            for i0_17_ in range(nw119_.length(0)):
                nw119_[i0_17_] = init17_(i0_17_)
            d_739_v11_ = nw119_
            d_744_v12_: _dafny.Array
            def lambda32_(d_745_v0_):
                def lambda33_(d_746_i1_):
                    return d_745_v0_

                return lambda33_

            init18_ = lambda32_(d_728_v0_)
            nw120_ = _dafny.Array(None, 22)
            for i0_18_ in range(nw120_.length(0)):
                nw120_[i0_18_] = init18_(i0_18_)
            d_744_v12_ = nw120_
            rhs91_ = not((d_728_v0_ if not(d_728_v0_) else d_728_v0_))
            rhs92_ = d_739_v11_
            rhs93_ = not (d_728_v0_) or (d_728_v0_)
            rhs94_ = d_728_v0_
            rhs95_ = d_744_v12_
            d_728_v0_ = rhs91_
            d_739_v11_ = rhs92_
            d_728_v0_ = rhs93_
            d_728_v0_ = rhs94_
            d_744_v12_ = rhs95_
            d_729_v1_ = ((d_729_v1_) + (default__.fm8(p1, globalState))) + (d_729_v1_)
        elif True:
            d_728_v0_ = d_728_v0_
            d_747_v13_: _dafny.Map
            d_747_v13_ = _dafny.Map({d_728_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqv"))})
            d_748_v14_: _dafny.Seq
            d_748_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "keousjicc"))
            d_749_v15_: _dafny.Seq
            d_749_v15_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_747_v13_)), len(d_748_v14_), p0, p0, (0) - (p1)])
            d_750_v16_: _dafny.Map
            d_750_v16_ = _dafny.Map({(p0 if False else p0): d_749_v15_})
            d_751_v17_: _dafny.Array
            def lambda34_(d_752_p0_, d_753_v0_, d_754_p1_):
                def lambda35_(d_755_i2_):
                    return D1_DC2(d_752_p0_, d_753_v0_, d_754_p1_)

                return lambda35_

            init19_ = lambda34_(p0, d_728_v0_, p1)
            nw121_ = _dafny.Array(None, 20)
            for i0_19_ in range(nw121_.length(0)):
                nw121_[i0_19_] = init19_(i0_19_)
            d_751_v17_ = nw121_
            d_756_v18_: D1
            d_756_v18_ = D1_DC2(p0, True, p1)
            index102_ = default__.safeIndex(445, (d_751_v17_).length(0))
            (d_751_v17_)[index102_] = d_756_v18_
            d_757_v19_: _dafny.Map
            d_757_v19_ = _dafny.Map({p0: d_748_v14_})
            d_758_v20_: _dafny.Set
            d_758_v20_ = _dafny.Set({(self).fm7(p0, d_728_v0_, d_757_v19_, globalState), False})
            d_759_v21_: _dafny.Map
            d_759_v21_ = _dafny.Map({d_728_v0_: p1})
            d_760_v22_: _dafny.Map
            d_760_v22_ = _dafny.Map({d_759_v21_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "grdlrs"))})
            d_761_v23_: _dafny.Array
            nw122_ = _dafny.Array(False, 4)
            d_761_v23_ = nw122_
            d_762_v24_: _dafny.Map
            d_762_v24_ = _dafny.Map({d_761_v23_: d_728_v0_})
            d_763_v25_: _dafny.Array
            nw123_ = _dafny.Array(None, 25)
            nw123_[int(0)] = (0) - (p1)
            nw123_[int(1)] = p1
            nw123_[int(2)] = p0
            nw123_[int(3)] = p0
            nw123_[int(4)] = p1
            nw123_[int(5)] = p1
            nw123_[int(6)] = (0) - (len(d_758_v20_))
            nw123_[int(7)] = 519
            nw123_[int(8)] = 118
            nw123_[int(9)] = p0
            nw123_[int(10)] = (d_749_v15_)[default__.safeIndex(p1, len(d_749_v15_))]
            nw123_[int(11)] = p1
            nw123_[int(12)] = p0
            nw123_[int(13)] = default__.fm9((0) - (p1), globalState)
            nw123_[int(14)] = p0
            nw123_[int(15)] = p0
            nw123_[int(16)] = (p1) * ((0) - (p1))
            nw123_[int(17)] = len(((d_760_v22_)[d_759_v21_] if (d_759_v21_) in (d_760_v22_) else _dafny.SeqWithoutIsStrInference([(self).f6 for d_764_i3_ in range(default__.abs(220))])))
            nw123_[int(18)] = ((d_759_v21_)[d_728_v0_] if (d_728_v0_) in (d_759_v21_) else p1)
            nw123_[int(19)] = default__.safeDivisionInt(default__.fm9(p0, globalState), len(d_748_v14_))
            nw123_[int(20)] = p1
            nw123_[int(21)] = default__.safeDivisionInt((0) - (p0), p0)
            nw123_[int(22)] = len((d_762_v24_) | (_dafny.Map({d_761_v23_: d_728_v0_})))
            nw123_[int(23)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngqqsqvhs")))
            nw123_[int(24)] = p1
            d_763_v25_ = nw123_
            d_765_v26_: _dafny.Seq
            d_765_v26_ = _dafny.SeqWithoutIsStrInference([d_728_v0_, d_728_v0_, d_728_v0_])
            d_766_v27_: _dafny.Seq
            d_766_v27_ = _dafny.SeqWithoutIsStrInference([default__.fm10(d_728_v0_, globalState), d_765_v26_])
            d_767_v28_: _dafny.Seq
            d_767_v28_ = (d_766_v27_)[default__.safeIndex(62, len(d_766_v27_))]
            index103_ = default__.safeIndex(351, (d_763_v25_).length(0))
            (d_763_v25_)[index103_] = len((d_767_v28_))
            d_768_v29_: _dafny.MultiSet
            d_768_v29_ = _dafny.MultiSet([p0, (len(d_749_v15_)) * (p0)])
            d_769_v30_: _dafny.Map
            d_769_v30_ = _dafny.Map({len(d_759_v21_): p1})
            d_770_v31_: _dafny.MultiSet
            d_770_v31_ = _dafny.MultiSet([default__.fm12(d_728_v0_, p0, d_758_v20_, d_769_v30_, globalState)])
            index104_ = default__.safeIndex(445, (d_751_v17_).length(0))
            index105_ = default__.safeIndex(351, (d_763_v25_).length(0))
            rhs96_ = (d_750_v16_) | (d_750_v16_)
            rhs97_ = d_756_v18_
            rhs98_ = p1
            rhs99_ = default__.fm11(d_728_v0_, d_728_v0_, D2_DC5(p0, d_728_v0_, d_728_v0_, d_728_v0_, d_728_v0_), globalState)
            rhs100_ = (0) - ((((d_770_v31_)[d_749_v15_] if (d_749_v15_) in (d_770_v31_) else len(d_765_v26_))) + (p1))
            lhs54_ = d_751_v17_
            lhs55_ = default__.safeIndex(445, (d_751_v17_).length(0))
            lhs56_ = d_763_v25_
            lhs57_ = default__.safeIndex(351, (d_763_v25_).length(0))
            lhs58_ = globalState
            d_750_v16_ = rhs96_
            lhs54_[lhs55_] = rhs97_
            lhs56_[lhs57_] = rhs98_
            d_768_v29_ = rhs99_
            lhs58_.f2 = rhs100_
            if d_728_v0_:
                index106_ = default__.safeIndex(351, (d_763_v25_).length(0))
                (d_763_v25_)[index106_] = (d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))]
                index107_ = default__.safeIndex(86, (d_761_v23_).length(0))
                (d_761_v23_)[index107_] = d_728_v0_
                d_771_v32_: _dafny.Map
                d_771_v32_ = _dafny.Map({d_728_v0_: d_728_v0_})
                d_772_v33_: _dafny.Set
                d_772_v33_ = _dafny.Set({(0) - ((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))]), len(d_769_v30_)})
                d_773_v34_: _dafny.Seq
                d_773_v34_ = _dafny.SeqWithoutIsStrInference([d_759_v21_])
                d_774_v35_: D2
                d_774_v35_ = D2_DC5((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))], d_728_v0_, d_728_v0_, ((d_771_v32_)[d_728_v0_] if (d_728_v0_) in (d_771_v32_) else d_728_v0_), (self).fm7(p0, d_728_v0_, d_757_v19_, globalState))
                d_775_v36_: _dafny.Map
                d_775_v36_ = _dafny.Map({d_761_v23_: d_774_v35_})
                d_776_v37_: _dafny.Map
                d_776_v37_ = _dafny.Map({d_775_v36_: d_772_v33_})
                index108_ = default__.safeIndex(86, (d_761_v23_).length(0))
                rhs101_ = not(True)
                rhs102_ = (((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))]) + (p0)) + ((p1) + (p0))
                rhs103_ = _dafny.SeqWithoutIsStrInference([(self).f6 for d_777_i4_ in range(default__.abs(285))])
                rhs104_ = (d_771_v32_).set((d_759_v21_) == ((d_773_v34_)[default__.safeIndex(p1, len(d_773_v34_))]), ((self).f6) in (d_748_v14_))
                rhs105_ = ((d_776_v37_)[d_775_v36_] if (d_775_v36_) in (d_776_v37_) else d_772_v33_)
                lhs59_ = d_761_v23_
                lhs60_ = default__.safeIndex(86, (d_761_v23_).length(0))
                lhs61_ = globalState
                lhs59_[lhs60_] = rhs101_
                lhs61_.f2 = rhs102_
                d_748_v14_ = rhs103_
                d_771_v32_ = rhs104_
                d_772_v33_ = rhs105_
                d_778_v38_: _dafny.Map
                d_778_v38_ = _dafny.Map({(d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))]: (d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))]})
                d_779_v41_: _dafny.Map
                d_779_v41_ = _dafny.Map({_dafny.CodePoint('o'): (d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))]})
                d_780_v42_: _dafny.Array
                nw124_ = _dafny.Array(None, 14)
                nw124_[int(0)] = d_778_v38_
                nw124_[int(1)] = d_778_v38_
                def iife38_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(305, 57):
                        d_781_v39_: int = compr_26_
                        if ((305) <= (d_781_v39_)) and ((d_781_v39_) < (57)):
                            coll26_[(d_781_v39_) - (p1)] = d_728_v0_
                    return _dafny.Map(coll26_)
                nw124_[int(2)] = iife38_()
                
                nw124_[int(3)] = d_778_v38_
                nw124_[int(4)] = _dafny.Map({(d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))]: (d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))]})
                nw124_[int(5)] = d_778_v38_
                nw124_[int(6)] = d_778_v38_
                def iife39_():
                    coll27_ = _dafny.Map()
                    compr_27_: int
                    for compr_27_ in (d_778_v38_).keys.Elements:
                        d_782_v40_: int = compr_27_
                        if (d_782_v40_) in (d_778_v38_):
                            coll27_[default__.safeModuloInt(d_782_v40_, p0)] = (d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))]
                    return _dafny.Map(coll27_)
                nw124_[int(7)] = iife39_()
                
                nw124_[int(8)] = default__.fm34((d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))], _dafny.CodePoint('k'), d_749_v15_, globalState)
                nw124_[int(9)] = d_778_v38_
                nw124_[int(10)] = d_778_v38_
                nw124_[int(11)] = _dafny.Map({len(d_769_v30_): ((d_779_v41_)[(self).f6] if ((self).f6) in (d_779_v41_) else ((d_771_v32_)[(d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))]] if ((d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))]) in (d_771_v32_) else d_728_v0_))})
                nw124_[int(12)] = d_778_v38_
                nw124_[int(13)] = _dafny.Map({(d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))]: d_728_v0_})
                d_780_v42_ = nw124_
                d_783_v43_: T0
                nw125_ = C2()
                nw125_.ctor__(True, (d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))], d_780_v42_)
                d_783_v43_ = nw125_
                d_784_v44_: _dafny.Map
                d_784_v44_ = _dafny.Map({(self).fm7((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))], d_728_v0_, d_757_v19_, globalState): d_783_v43_})
                d_785_v45_: _dafny.Seq
                d_785_v45_ = _dafny.SeqWithoutIsStrInference([d_783_v43_])
                d_784_v44_ = (d_784_v44_).set(True, (d_785_v45_)[default__.safeIndex((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))], len(d_785_v45_))])
                d_786_v46_: D3
                d_786_v46_ = D3_DC7(len(d_748_v14_), p0)
                d_787_v47_: _dafny.Map
                d_787_v47_ = _dafny.Map({d_786_v46_: d_763_v25_})
                d_788_v48_: _dafny.Map
                d_788_v48_ = _dafny.Map({d_728_v0_: d_763_v25_})
                d_789_v49_: _dafny.Array
                nw126_ = _dafny.Array(None, 17)
                nw126_[int(0)] = ((d_787_v47_)[d_786_v46_] if (d_786_v46_) in (d_787_v47_) else d_763_v25_)
                nw126_[int(1)] = d_763_v25_
                nw126_[int(2)] = d_763_v25_
                nw126_[int(3)] = d_763_v25_
                nw126_[int(4)] = d_763_v25_
                nw126_[int(5)] = d_763_v25_
                nw126_[int(6)] = d_763_v25_
                nw126_[int(7)] = ((d_788_v48_)[(d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))]] if ((d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))]) in (d_788_v48_) else d_763_v25_)
                nw126_[int(8)] = d_763_v25_
                nw126_[int(9)] = d_763_v25_
                nw126_[int(10)] = d_763_v25_
                nw126_[int(11)] = d_763_v25_
                nw126_[int(12)] = d_763_v25_
                nw126_[int(13)] = d_763_v25_
                nw126_[int(14)] = d_763_v25_
                nw126_[int(15)] = d_763_v25_
                nw126_[int(16)] = d_763_v25_
                d_789_v49_ = nw126_
                nw127_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_789_v49_ = nw127_
                d_771_v32_ = (d_771_v32_).set((d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))], (d_761_v23_)[default__.safeIndex(86, (d_761_v23_).length(0))])
            elif True:
                d_790_v50_: D3
                d_790_v50_ = D3_DC7((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))], p0)
                d_791_v51_: _dafny.Seq
                d_791_v51_ = _dafny.SeqWithoutIsStrInference([d_790_v50_, d_790_v50_, d_790_v50_])
                index109_ = default__.safeIndex(140, (d_761_v23_).length(0))
                (d_761_v23_)[index109_] = d_728_v0_
                d_792_v52_: D4
                d_792_v52_ = D4_DC9(p0, (self).f6, d_749_v15_, len(d_748_v14_), True)
                d_793_v53_: _dafny.Map
                d_793_v53_ = _dafny.Map({d_792_v52_: d_761_v23_})
                index110_ = default__.safeIndex(140, (d_761_v23_).length(0))
                rhs106_ = d_791_v51_
                rhs107_ = (d_748_v14_).set(default__.safeIndex((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))], len(d_748_v14_)), (self).f6)
                rhs108_ = default__.fm17(_dafny.SeqWithoutIsStrInference([d_748_v14_]), p1, default__.safeModuloInt(p0, -351), (d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))], globalState)
                rhs109_ = ((d_793_v53_)[d_792_v52_] if (d_792_v52_) in (d_793_v53_) else d_761_v23_)
                rhs110_ = (d_728_v0_ if (d_765_v26_)[default__.safeIndex((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))], len(d_765_v26_))] else not((len(d_748_v14_)) > ((d_763_v25_)[default__.safeIndex(351, (d_763_v25_).length(0))])))
                lhs62_ = globalState
                lhs63_ = d_761_v23_
                lhs64_ = default__.safeIndex(140, (d_761_v23_).length(0))
                d_791_v51_ = rhs106_
                d_748_v14_ = rhs107_
                lhs62_.f2 = rhs108_
                d_761_v23_ = rhs109_
                lhs63_[lhs64_] = rhs110_
                index111_ = default__.safeIndex(140, (d_761_v23_).length(0))
                (d_761_v23_)[index111_] = default__.fm23((d_748_v14_) + ((d_748_v14_).set(default__.safeIndex(len(d_749_v15_), len(d_748_v14_)), (self).f6)), p1, p1, not(False), globalState)
                d_794_v54_: _dafny.Seq
                d_794_v54_ = _dafny.SeqWithoutIsStrInference([d_767_v28_])
                d_728_v0_ = (d_794_v54_) <= (d_794_v54_)
                index112_ = default__.safeIndex(351, (d_763_v25_).length(0))
                (d_763_v25_)[index112_] = p0
                d_759_v21_ = d_759_v21_
            index113_ = default__.safeIndex(351, (d_763_v25_).length(0))
            rhs111_ = len(_dafny.SeqWithoutIsStrInference([(self).f6 for d_795_i5_ in range(default__.abs(356))]))
            rhs112_ = _dafny.Map({True: (d_748_v14_) + (_dafny.SeqWithoutIsStrInference([(self).f6 for d_796_i6_ in range(default__.abs(121))]))})
            lhs65_ = d_763_v25_
            lhs66_ = default__.safeIndex(351, (d_763_v25_).length(0))
            lhs65_[lhs66_] = rhs111_
            d_747_v13_ = rhs112_
            d_797_v55_: _dafny.Array
            nw128_ = _dafny.Array(_dafny.Map({}), 2)
            d_797_v55_ = nw128_
            d_798_v56_: C2
            nw129_ = C2()
            nw129_.ctor__((d_765_v26_)[default__.safeIndex(len(d_748_v14_), len(d_765_v26_))], len(d_765_v26_), d_797_v55_)
            d_798_v56_ = nw129_
        d_799_v57_: D6
        d_799_v57_ = D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))
        source14_ = d_799_v57_
        if source14_.is_DC12:
            d_800___mcc_h0_ = source14_.cf23
            d_801___mcc_h1_ = source14_.cf24
            d_802_cf24_ = d_801___mcc_h1_
            d_803_cf23_ = d_800___mcc_h0_
            if d_803_cf23_:
                d_804_v58_: _dafny.Seq
                d_804_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgyx"))
                d_805_v59_: _dafny.Array
                nw130_ = _dafny.Array(None, 6)
                nw130_[int(0)] = d_804_v58_
                nw130_[int(1)] = d_804_v58_
                nw130_[int(2)] = _dafny.SeqWithoutIsStrInference([d_802_cf24_ for d_806_i7_ in range(default__.abs(536))])
                nw130_[int(3)] = d_804_v58_
                nw130_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ef"))
                nw130_[int(5)] = d_804_v58_
                d_805_v59_ = nw130_
                d_807_v60_: D9
                d_807_v60_ = D9_DC18(d_805_v59_)
                d_808_v61_: _dafny.Array
                nw131_ = _dafny.Array(None, 18)
                nw131_[int(0)] = d_805_v59_
                nw131_[int(1)] = d_805_v59_
                nw131_[int(2)] = d_805_v59_
                nw131_[int(3)] = d_805_v59_
                nw131_[int(4)] = d_805_v59_
                nw131_[int(5)] = d_805_v59_
                nw131_[int(6)] = d_805_v59_
                nw131_[int(7)] = d_805_v59_
                nw131_[int(8)] = d_805_v59_
                nw131_[int(9)] = d_805_v59_
                nw131_[int(10)] = d_805_v59_
                nw131_[int(11)] = (d_807_v60_).cf33
                nw131_[int(12)] = d_805_v59_
                nw131_[int(13)] = d_805_v59_
                nw131_[int(14)] = d_805_v59_
                nw131_[int(15)] = (d_805_v59_ if d_803_cf23_ else d_805_v59_)
                nw131_[int(16)] = d_805_v59_
                nw131_[int(17)] = d_805_v59_
                d_808_v61_ = nw131_
                index114_ = default__.safeIndex(959, (d_808_v61_).length(0))
                (d_808_v61_)[index114_] = d_805_v59_
                index115_ = default__.safeIndex(959, (d_808_v61_).length(0))
                (d_808_v61_)[index115_] = d_805_v59_
                d_809_v62_: _dafny.Array
                nw132_ = _dafny.Array(False, 8)
                d_809_v62_ = nw132_
                d_809_v62_ = d_809_v62_
                d_803_cf23_ = not(default__.fm21((p1) * (73), _dafny.CodePoint('i'), (d_804_v58_) < (d_804_v58_), d_803_cf23_, globalState))
                d_810_v63_: _dafny.Set
                d_810_v63_ = _dafny.Set({not(d_728_v0_)})
                d_811_v64_: _dafny.Map
                d_811_v64_ = _dafny.Map({d_810_v63_: d_804_v58_})
                d_811_v64_ = (d_811_v64_).set(d_810_v63_, _dafny.SeqWithoutIsStrInference([(self).f6 for d_812_i8_ in range(default__.abs(981))]))
                d_813_v66_: D2
                def iife40_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(345, -615):
                        d_814_v65_: int = compr_28_
                        if ((345) <= (d_814_v65_)) and ((d_814_v65_) < (-615)):
                            coll28_[(d_814_v65_) * (p0)] = p1
                    return _dafny.Map(coll28_)
                d_813_v66_ = D2_DC5(len(iife40_()
), d_728_v0_, d_803_cf23_, d_803_cf23_, d_803_cf23_)
                pat_let_tv20_ = d_803_cf23_
                pat_let_tv21_ = d_803_cf23_
                def iife41_(_pat_let6_0):
                    def iife42_(d_815_dt__update__tmp_h0_):
                        def iife43_(_pat_let7_0):
                            def iife44_(d_816_dt__update_hcf8_h0_):
                                def iife45_(_pat_let8_0):
                                    def iife46_(d_817_dt__update_hcf10_h0_):
                                        return D2_DC5((d_815_dt__update__tmp_h0_).cf7, d_816_dt__update_hcf8_h0_, (d_815_dt__update__tmp_h0_).cf9, d_817_dt__update_hcf10_h0_, (d_815_dt__update__tmp_h0_).cf11)
                                    return iife46_(_pat_let8_0)
                                return iife45_(pat_let_tv21_)
                            return iife44_(_pat_let7_0)
                        return iife43_(pat_let_tv20_)
                    return iife42_(_pat_let6_0)
                d_813_v66_ = iife41_(d_813_v66_)
            elif True:
                d_818_v67_: _dafny.Map
                d_818_v67_ = _dafny.Map({(d_728_v0_) and (d_803_cf23_): p1})
                d_818_v67_ = (d_818_v67_).set(d_803_cf23_, p0)
                d_819_v68_: _dafny.Seq
                d_819_v68_ = _dafny.SeqWithoutIsStrInference([291, p0])
                d_820_v69_: _dafny.Seq
                d_820_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulgwrp"))
                d_821_v70_: _dafny.Array
                nw133_ = _dafny.Array(_dafny.Map({}), 7)
                d_821_v70_ = nw133_
                d_822_v71_: T0
                nw134_ = C2()
                nw134_.ctor__(True, p1, d_821_v70_)
                d_822_v71_ = nw134_
                d_823_v72_: _dafny.Set
                d_823_v72_ = _dafny.Set({p0})
                d_824_v73_: C0
                nw135_ = C0()
                nw135_.ctor__(p0)
                d_824_v73_ = nw135_
                d_825_v74_: _dafny.MultiSet
                d_825_v74_ = _dafny.MultiSet([d_824_v73_, d_824_v73_])
                d_826_v75_: _dafny.Array
                nw136_ = _dafny.Array(None, 15)
                nw136_[int(0)] = p0
                nw136_[int(1)] = len(d_820_v69_)
                nw136_[int(2)] = p0
                nw136_[int(3)] = p1
                nw136_[int(4)] = p1
                nw136_[int(5)] = p1
                nw136_[int(6)] = len(d_819_v68_)
                nw136_[int(7)] = len(_dafny.Map({d_822_v71_: d_803_cf23_}))
                nw136_[int(8)] = len(d_823_v72_)
                nw136_[int(9)] = p1
                nw136_[int(10)] = p1
                nw136_[int(11)] = p1
                nw136_[int(12)] = len(d_819_v68_)
                nw136_[int(13)] = (d_825_v74_).cardinality
                nw136_[int(14)] = -423
                d_826_v75_ = nw136_
                d_827_v76_: _dafny.Map
                d_827_v76_ = _dafny.Map({d_826_v75_: d_728_v0_})
                d_828_v77_: D8
                d_828_v77_ = D8_DC16((d_819_v68_)[default__.safeIndex(p1, len(d_819_v68_))], ((d_827_v76_)[d_826_v75_] if (d_826_v75_) in (d_827_v76_) else d_803_cf23_), d_803_cf23_, p1)
                d_728_v0_ = (d_828_v77_).cf29
                d_829_v78_: _dafny.Seq
                d_829_v78_ = _dafny.SeqWithoutIsStrInference([d_820_v69_])
                d_820_v69_ = (d_820_v69_) + (((d_820_v69_) + (_dafny.SeqWithoutIsStrInference([(self).f6 for d_830_i9_ in range(default__.abs(13))]))).set(default__.safeIndex((d_819_v68_)[default__.safeIndex(default__.fm17(d_829_v78_, p0, len(d_820_v69_), len(d_823_v72_), globalState), len(d_819_v68_))], len((d_820_v69_) + (_dafny.SeqWithoutIsStrInference([(self).f6 for d_831_i9_ in range(default__.abs(13))])))), d_802_cf24_))
                index116_ = default__.safeIndex(309, (d_826_v75_).length(0))
                (d_826_v75_)[index116_] = p0
                d_832_v79_: _dafny.Map
                d_832_v79_ = _dafny.Map({default__.fm35(globalState): d_803_cf23_})
                d_833_v80_: _dafny.Set
                d_833_v80_ = _dafny.Set({d_820_v69_})
                index117_ = default__.safeIndex(309, (d_826_v75_).length(0))
                rhs113_ = p1
                rhs114_ = d_799_v57_
                rhs115_ = p1
                rhs116_ = ((d_832_v79_)[d_833_v80_] if (d_833_v80_) in (d_832_v79_) else not(d_728_v0_))
                rhs117_ = d_823_v72_
                lhs67_ = globalState
                lhs68_ = d_826_v75_
                lhs69_ = default__.safeIndex(309, (d_826_v75_).length(0))
                lhs67_.f2 = rhs113_
                d_799_v57_ = rhs114_
                lhs68_[lhs69_] = rhs115_
                d_728_v0_ = rhs116_
                d_823_v72_ = rhs117_
                d_834_v81_: _dafny.Map
                d_834_v81_ = _dafny.Map({d_826_v75_: p0})
                d_834_v81_ = (d_834_v81_).set(d_826_v75_, 657)
            (globalState).f2 = p1
            d_835_v82_: _dafny.Seq
            d_835_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_835_v82_ = (d_835_v82_) + ((d_835_v82_ if d_728_v0_ else d_835_v82_))
            if default__.fm23(_dafny.SeqWithoutIsStrInference([d_802_cf24_ for d_836_i10_ in range(default__.abs(608))]), 296, default__.safeModuloInt(p0, len(d_835_v82_)), d_803_cf23_, globalState):
                d_837_v83_: _dafny.Array
                def lambda36_(d_838_p0_):
                    def lambda37_(d_839_i11_):
                        return (d_839_i11_) - (d_838_p0_)

                    return lambda37_

                init20_ = lambda36_(p0)
                nw137_ = _dafny.Array(None, 4)
                for i0_20_ in range(nw137_.length(0)):
                    nw137_[i0_20_] = init20_(i0_20_)
                d_837_v83_ = nw137_
                index118_ = default__.safeIndex(915, (d_837_v83_).length(0))
                (d_837_v83_)[index118_] = p0
                d_840_v84_: _dafny.Seq
                d_840_v84_ = _dafny.SeqWithoutIsStrInference([d_835_v82_, _dafny.SeqWithoutIsStrInference([d_802_cf24_ for d_841_i12_ in range(default__.abs(775))])])
                d_842_v85_: _dafny.Seq
                d_842_v85_ = _dafny.SeqWithoutIsStrInference([(d_840_v84_)[default__.safeIndex(p0, len(d_840_v84_))], d_835_v82_, default__.fm25(d_728_v0_, p0, globalState)])
                d_843_v86_: _dafny.Map
                d_843_v86_ = _dafny.Map({p0: d_835_v82_})
                index119_ = default__.safeIndex(915, (d_837_v83_).length(0))
                rhs118_ = default__.fm17((d_842_v85_ if d_728_v0_ else d_842_v85_), p1, 82, p0, globalState)
                rhs119_ = (self).fm7(p0, d_803_cf23_, (d_843_v86_).set(p0, d_835_v82_), globalState)
                lhs70_ = d_837_v83_
                lhs71_ = default__.safeIndex(915, (d_837_v83_).length(0))
                lhs70_[lhs71_] = rhs118_
                d_803_cf23_ = rhs119_
                d_844_v87_: _dafny.Array
                def lambda38_(d_845_v0_, d_846_cf23_):
                    def lambda39_(d_847_i13_):
                        return (_dafny.SeqWithoutIsStrInference([d_845_v0_, d_846_cf23_])) + (_dafny.SeqWithoutIsStrInference([not(True)]))

                    return lambda39_

                init21_ = lambda38_(d_728_v0_, d_803_cf23_)
                nw138_ = _dafny.Array(None, 15)
                for i0_21_ in range(nw138_.length(0)):
                    nw138_[i0_21_] = init21_(i0_21_)
                d_844_v87_ = nw138_
                d_848_v88_: _dafny.Seq
                d_848_v88_ = _dafny.SeqWithoutIsStrInference([d_803_cf23_])
                index120_ = default__.safeIndex(45, (d_844_v87_).length(0))
                (d_844_v87_)[index120_] = (d_848_v88_) + (d_848_v88_)
                index121_ = default__.safeIndex(45, (d_844_v87_).length(0))
                (d_844_v87_)[index121_] = (_dafny.SeqWithoutIsStrInference([d_803_cf23_])) + (d_848_v88_)
                d_849_v89_: _dafny.Seq
                d_849_v89_ = _dafny.SeqWithoutIsStrInference([((d_837_v83_)[default__.safeIndex(915, (d_837_v83_).length(0))]) + (p0)])
                d_850_v90_: D4
                d_850_v90_ = D4_DC9((0) - ((d_837_v83_)[default__.safeIndex(915, (d_837_v83_).length(0))]), (self).f6, d_849_v89_, p0, d_803_cf23_)
                d_851_v91_: _dafny.Set
                d_851_v91_ = _dafny.Set({d_728_v0_, d_803_cf23_})
                d_852_v92_: _dafny.Map
                d_852_v92_ = _dafny.Map({p0: p0})
                d_849_v89_ = (d_849_v89_ if d_803_cf23_ else ((d_850_v90_).cf18) + (default__.fm12(True, p0, d_851_v91_, d_852_v92_, globalState)))
                d_853_v93_: _dafny.Map
                d_853_v93_ = _dafny.Map({False: p0})
                d_853_v93_ = (d_853_v93_).set(d_803_cf23_, (0) - (p1))
                d_728_v0_ = ((834) == (p1) if (_dafny.SeqWithoutIsStrInference([(d_837_v83_)[default__.safeIndex(915, (d_837_v83_).length(0))] for d_854_i14_ in range(default__.abs(900))])) <= (_dafny.SeqWithoutIsStrInference([p0, p0])) else d_728_v0_)
            elif True:
                d_855_v94_: _dafny.Array
                def lambda40_(d_856_p1_):
                    def lambda41_(d_857_i15_):
                        return (d_857_i15_) + (d_856_p1_)

                    return lambda41_

                init22_ = lambda40_(p1)
                nw139_ = _dafny.Array(None, 26)
                for i0_22_ in range(nw139_.length(0)):
                    nw139_[i0_22_] = init22_(i0_22_)
                d_855_v94_ = nw139_
                index122_ = default__.safeIndex(457, (d_855_v94_).length(0))
                (d_855_v94_)[index122_] = p1
                index123_ = default__.safeIndex(457, (d_855_v94_).length(0))
                (d_855_v94_)[index123_] = len((((d_835_v82_).set(default__.safeIndex(p1, len(d_835_v82_)), d_802_cf24_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lojxgqr"))))
                d_858_v95_: C0
                nw140_ = C0()
                nw140_.ctor__(p0)
                d_858_v95_ = nw140_
                d_859_v96_: D4
                d_859_v96_ = D4_DC8(d_858_v95_)
                d_860_v97_: _dafny.MultiSet
                d_860_v97_ = _dafny.MultiSet([d_859_v96_, d_859_v96_])
                d_861_v98_: D7
                d_861_v98_ = D7_DC14(((d_860_v97_)[d_859_v96_] if (d_859_v96_) in (d_860_v97_) else p0))
                pat_let_tv22_ = d_855_v94_
                pat_let_tv23_ = d_855_v94_
                def iife47_(_pat_let9_0):
                    def iife48_(d_862_dt__update__tmp_h1_):
                        def iife49_(_pat_let10_0):
                            def iife50_(d_863_dt__update_hcf26_h0_):
                                return D7_DC14(d_863_dt__update_hcf26_h0_)
                            return iife50_(_pat_let10_0)
                        return iife49_((pat_let_tv23_)[default__.safeIndex(457, (pat_let_tv22_).length(0))])
                    return iife48_(_pat_let9_0)
                d_861_v98_ = iife47_(default__.fm36(d_728_v0_, globalState))
                d_864_v99_: _dafny.Array
                d_864_v99_ = d_855_v94_
                d_865_v100_: _dafny.MultiSet
                d_865_v100_ = _dafny.MultiSet([(d_864_v99_), d_855_v94_, d_855_v94_])
                d_865_v100_ = _dafny.MultiSet([d_855_v94_, (d_855_v94_ if d_728_v0_ else d_855_v94_), d_855_v94_, d_855_v94_, d_855_v94_])
                d_865_v100_ = d_865_v100_
                d_803_cf23_ = d_803_cf23_
        elif True:
            d_866___mcc_h2_ = source14_.cf22
            d_867_cf22_ = d_866___mcc_h2_
            d_868_v101_: _dafny.Seq
            d_868_v101_ = _dafny.SeqWithoutIsStrInference([d_867_cf22_, d_867_cf22_])
            d_869_v102_: _dafny.Map
            d_869_v102_ = _dafny.Map({((self).f6) in ((d_799_v57_).cf22): default__.fm17(d_868_v101_, p1, 803, p0, globalState)})
            d_869_v102_ = (d_869_v102_).set(d_728_v0_, (p1 if d_728_v0_ else len(d_867_cf22_)))
            d_870_v103_: _dafny.Seq
            d_870_v103_ = _dafny.SeqWithoutIsStrInference([p0, 192])
            d_871_v104_: D2
            d_871_v104_ = D2_DC4((d_870_v103_)[default__.safeIndex(p0, len(d_870_v103_))])
            d_872_v105_: _dafny.Seq
            d_872_v105_ = _dafny.SeqWithoutIsStrInference([d_871_v104_])
            d_872_v105_ = (_dafny.SeqWithoutIsStrInference([d_871_v104_])) + ((d_872_v105_) + (d_872_v105_))
            d_873_v106_: _dafny.Array
            nw141_ = _dafny.Array(False, 3)
            d_873_v106_ = nw141_
            d_874_v107_: _dafny.Seq
            d_874_v107_ = _dafny.SeqWithoutIsStrInference([d_873_v106_])
            d_875_v108_: _dafny.Seq
            d_875_v108_ = _dafny.SeqWithoutIsStrInference([(d_874_v107_)[default__.safeIndex(p1, len(d_874_v107_))], d_873_v106_])
            d_876_v109_: D8
            d_876_v109_ = D8_DC15((d_875_v108_)[default__.safeIndex(p0, len(d_875_v108_))])
            source15_ = d_876_v109_
            if source15_.is_DC16:
                d_877___mcc_h3_ = source15_.cf28
                d_878___mcc_h4_ = source15_.cf29
                d_879___mcc_h5_ = source15_.cf30
                d_880___mcc_h6_ = source15_.cf31
                d_881_cf31_ = d_880___mcc_h6_
                d_882_cf30_ = d_879___mcc_h5_
                d_883_cf29_ = d_878___mcc_h4_
                d_884_cf28_ = d_877___mcc_h3_
                d_728_v0_ = d_728_v0_
                d_882_cf30_ = d_882_cf30_
                d_885_v110_: _dafny.Array
                nw142_ = _dafny.Array(int(0), 8)
                d_885_v110_ = nw142_
                d_886_v111_: _dafny.MultiSet
                d_886_v111_ = _dafny.MultiSet([p1])
                index124_ = default__.safeIndex(150, (d_885_v110_).length(0))
                (d_885_v110_)[index124_] = default__.safeModuloInt(((d_886_v111_)[d_881_cf31_] if (d_881_cf31_) in (d_886_v111_) else len(_dafny.Map({d_883_cf29_: p0}))), (d_871_v104_).cf6)
                index125_ = default__.safeIndex(150, (d_885_v110_).length(0))
                (d_885_v110_)[index125_] = d_884_cf28_
                d_887_v112_: str
                d_887_v112_ = _dafny.CodePoint('m')
                d_887_v112_ = d_887_v112_
            elif source15_.is_DC15:
                d_888___mcc_h7_ = source15_.cf27
                d_889_cf27_ = d_888___mcc_h7_
                d_890_v113_: _dafny.Seq
                d_890_v113_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_728_v0_: d_728_v0_})])
                d_891_v114_: _dafny.Map
                d_891_v114_ = _dafny.Map({p1: (d_890_v113_)[default__.safeIndex(p0, len(d_890_v113_))]})
                d_892_v115_: _dafny.Set
                d_892_v115_ = _dafny.Set({p1})
                d_893_v116_: _dafny.Map
                d_893_v116_ = _dafny.Map({d_728_v0_: d_728_v0_})
                d_894_v117_: _dafny.MultiSet
                d_894_v117_ = _dafny.MultiSet([_dafny.Map({d_728_v0_: d_728_v0_}), ((d_891_v114_)[default__.fm24(d_892_v115_, globalState)] if (default__.fm24(d_892_v115_, globalState)) in (d_891_v114_) else d_893_v116_)])
                d_895_v118_: _dafny.Seq
                d_895_v118_ = _dafny.SeqWithoutIsStrInference([d_894_v117_])
                d_896_v119_: _dafny.Seq
                d_896_v119_ = _dafny.SeqWithoutIsStrInference([d_728_v0_, d_728_v0_, d_728_v0_, d_728_v0_])
                d_728_v0_ = (((d_895_v118_)[default__.safeIndex(len(d_896_v119_), len(d_895_v118_))]) | (d_894_v117_)).issubset(_dafny.MultiSet((d_890_v113_) + (default__.fm37(D1_DC2(p1, False, p0), d_728_v0_, globalState))))
                d_728_v0_ = d_728_v0_
                (globalState).f2 = 14
                d_897_v121_: _dafny.Map
                d_897_v121_ = _dafny.Map({d_867_cf22_: 746})
                d_898_v122_: _dafny.MultiSet
                def iife51_():
                    coll29_ = _dafny.Map()
                    compr_29_: _dafny.Seq
                    for compr_29_ in (d_897_v121_).keys.Elements:
                        d_899_v120_: _dafny.Seq = compr_29_
                        if (d_899_v120_) in (d_897_v121_):
                            coll29_[d_899_v120_] = d_728_v0_
                    return _dafny.Map(coll29_)
                d_898_v122_ = _dafny.MultiSet([p1, len(iife51_()
                ), p1])
                d_900_v123_: C3
                nw143_ = C3()
                nw143_.ctor__(d_898_v122_, d_728_v0_)
                d_900_v123_ = nw143_
                nw144_ = C3()
                nw144_.ctor__((d_900_v123_).f9, True)
                d_900_v123_ = nw144_
            elif True:
                d_901___mcc_h8_ = source15_.cf32
                d_902_cf32_ = d_901___mcc_h8_
                (globalState).f2 = p0
                d_903_v124_: _dafny.MultiSet
                d_903_v124_ = _dafny.MultiSet([p1, p1])
                d_904_v125_: _dafny.Array
                nw145_ = _dafny.Array(_dafny.Map({}), 21)
                d_904_v125_ = nw145_
                d_905_v126_: C4
                nw146_ = C4()
                nw146_.ctor__(d_728_v0_, d_903_v124_, d_904_v125_)
                d_905_v126_ = nw146_
                d_906_v127_: _dafny.Array
                def lambda42_(d_907_i16_):
                    return default__.safeDivisionInt(d_907_i16_, 778)

                init23_ = lambda42_
                nw147_ = _dafny.Array(None, 28)
                for i0_23_ in range(nw147_.length(0)):
                    nw147_[i0_23_] = init23_(i0_23_)
                d_906_v127_ = nw147_
                d_908_v128_: _dafny.Map
                d_908_v128_ = _dafny.Map({_dafny.CodePoint('e'): d_906_v127_})
                d_909_v129_: _dafny.Array
                d_909_v129_ = d_906_v127_
                d_910_v130_: _dafny.Map
                d_910_v130_ = _dafny.Map({d_728_v0_: d_906_v127_})
                d_911_v131_: _dafny.Array
                nw148_ = _dafny.Array(None, 17)
                nw148_[int(0)] = (d_906_v127_ if d_728_v0_ else d_906_v127_)
                nw148_[int(1)] = d_906_v127_
                nw148_[int(2)] = d_906_v127_
                nw148_[int(3)] = d_906_v127_
                nw148_[int(4)] = d_906_v127_
                nw148_[int(5)] = ((d_908_v128_)[(self).f6] if ((self).f6) in (d_908_v128_) else d_906_v127_)
                nw148_[int(6)] = d_906_v127_
                nw148_[int(7)] = d_906_v127_
                nw148_[int(8)] = (d_909_v129_)
                nw148_[int(9)] = d_906_v127_
                nw148_[int(10)] = d_906_v127_
                nw148_[int(11)] = d_906_v127_
                nw148_[int(12)] = d_906_v127_
                nw148_[int(13)] = d_906_v127_
                nw148_[int(14)] = ((d_910_v130_)[(d_905_v126_).f7] if ((d_905_v126_).f7) in (d_910_v130_) else d_906_v127_)
                nw148_[int(15)] = d_906_v127_
                nw148_[int(16)] = d_906_v127_
                d_911_v131_ = nw148_
                index126_ = default__.safeIndex(536, (d_911_v131_).length(0))
                (d_911_v131_)[index126_] = d_906_v127_
                index127_ = default__.safeIndex(536, (d_911_v131_).length(0))
                (d_911_v131_)[index127_] = d_906_v127_
                rhs120_ = d_728_v0_
                rhs121_ = len(((d_867_cf22_) + (d_867_cf22_) if (p0) == (439) else d_867_cf22_))
                lhs72_ = globalState
                d_728_v0_ = rhs120_
                lhs72_.f2 = rhs121_
            d_912_v132_: _dafny.Array
            nw149_ = _dafny.Array(int(0), 9)
            d_912_v132_ = nw149_
            d_913_v133_: _dafny.MultiSet
            d_913_v133_ = _dafny.MultiSet([p1, p0, p0, p1])
            d_914_v134_: _dafny.Map
            d_914_v134_ = _dafny.Map({not(d_728_v0_): d_728_v0_})
            index128_ = default__.safeIndex(50, (d_912_v132_).length(0))
            (d_912_v132_)[index128_] = ((d_913_v133_)[default__.fm9(len(d_867_cf22_), globalState)] if (default__.fm9(len(d_867_cf22_), globalState)) in (d_913_v133_) else (0) - (len(d_914_v134_)))
            d_915_v135_: _dafny.Map
            d_915_v135_ = _dafny.Map({_dafny.CodePoint('b'): (d_867_cf22_) + (d_867_cf22_)})
            d_916_v136_: _dafny.Map
            d_916_v136_ = _dafny.Map({p1: d_912_v132_})
            d_917_v138_: _dafny.Set
            d_917_v138_ = _dafny.Set({(self).f6})
            index129_ = default__.safeIndex(50, (d_912_v132_).length(0))
            def iife52_():
                coll30_ = _dafny.Map()
                compr_30_: str
                for compr_30_ in (d_917_v138_).Elements:
                    d_918_v137_: str = compr_30_
                    if (d_918_v137_) in (d_917_v138_):
                        coll30_[d_918_v137_] = d_867_cf22_
                return _dafny.Map(coll30_)
            rhs122_ = (p0 if not(True) else p0)
            rhs123_ = (d_915_v135_ if d_728_v0_ else (iife52_()
            ).set(_dafny.CodePoint('f'), d_867_cf22_))
            rhs124_ = default__.safeModuloInt((p1) + (p1), p1)
            rhs125_ = p1
            rhs126_ = _dafny.Map({p1: d_912_v132_})
            lhs73_ = d_912_v132_
            lhs74_ = default__.safeIndex(50, (d_912_v132_).length(0))
            lhs75_ = globalState
            lhs76_ = globalState
            lhs73_[lhs74_] = rhs122_
            d_915_v135_ = rhs123_
            lhs75_.f2 = rhs124_
            lhs76_.f2 = rhs125_
            d_916_v136_ = rhs126_
        d_919_v139_: _dafny.Seq
        d_919_v139_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
        d_920_v140_: _dafny.Seq
        d_920_v140_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_921_v141_: _dafny.Map
        d_921_v141_ = _dafny.Map({len(d_920_v140_): not(d_728_v0_)})
        d_922_v142_: _dafny.Array
        nw150_ = _dafny.Array(None, 11)
        nw150_[int(0)] = d_921_v141_
        nw150_[int(1)] = d_921_v141_
        nw150_[int(2)] = d_921_v141_
        nw150_[int(3)] = d_921_v141_
        nw150_[int(4)] = d_921_v141_
        nw150_[int(5)] = d_921_v141_
        nw150_[int(6)] = d_921_v141_
        nw150_[int(7)] = d_921_v141_
        nw150_[int(8)] = d_921_v141_
        nw150_[int(9)] = d_921_v141_
        nw150_[int(10)] = d_921_v141_
        d_922_v142_ = nw150_
        d_923_v143_: T0
        nw151_ = C2()
        nw151_.ctor__(d_728_v0_, len(d_919_v139_), d_922_v142_)
        d_923_v143_ = nw151_
        d_924_v144_: _dafny.Map
        d_924_v144_ = _dafny.Map({p1: d_923_v143_})
        d_925_v145_: D2
        d_925_v145_ = D2_DC5(p1, False, d_728_v0_, d_728_v0_, d_728_v0_)
        source16_ = D2_DC5(len(d_924_v144_), (d_925_v145_).cf8, d_728_v0_, d_728_v0_, ((self).f6) not in (d_919_v139_))
        if source16_.is_DC4:
            d_926___mcc_h9_ = source16_.cf6
            d_927_cf6_ = d_926___mcc_h9_
            d_928_v146_: C2
            nw152_ = C2()
            nw152_.ctor__(False, len((d_920_v140_ if not(d_728_v0_) else d_920_v140_)), (d_923_v143_).f3)
            d_928_v146_ = nw152_
            d_929_v147_: _dafny.Array
            nw153_ = _dafny.Array(None, 26)
            nw153_[int(0)] = (d_928_v146_).f11
            nw153_[int(1)] = (d_928_v146_).f11
            nw153_[int(2)] = d_728_v0_
            nw153_[int(3)] = d_728_v0_
            nw153_[int(4)] = (d_928_v146_).f11
            nw153_[int(5)] = (d_928_v146_).f11
            nw153_[int(6)] = default__.fm23(default__.fm25(default__.fm23(d_919_v139_, 924, (d_928_v146_).f12, d_728_v0_, globalState), p1, globalState), 379, d_927_cf6_, (d_928_v146_).f11, globalState)
            nw153_[int(7)] = not(d_728_v0_)
            nw153_[int(8)] = d_728_v0_
            nw153_[int(9)] = d_728_v0_
            nw153_[int(10)] = d_728_v0_
            nw153_[int(11)] = d_728_v0_
            nw153_[int(12)] = d_728_v0_
            nw153_[int(13)] = False
            nw153_[int(14)] = d_728_v0_
            nw153_[int(15)] = (d_928_v146_).f11
            nw153_[int(16)] = False
            nw153_[int(17)] = True
            nw153_[int(18)] = (d_928_v146_).f11
            nw153_[int(19)] = not((d_928_v146_).f11)
            nw153_[int(20)] = d_728_v0_
            nw153_[int(21)] = (d_928_v146_).f11
            nw153_[int(22)] = (d_928_v146_).f11
            nw153_[int(23)] = (d_928_v146_).f11
            nw153_[int(24)] = (d_928_v146_).f11
            nw153_[int(25)] = d_728_v0_
            d_929_v147_ = nw153_
            d_930_v148_: _dafny.MultiSet
            d_930_v148_ = _dafny.MultiSet([len(d_921_v141_)])
            d_931_v149_: C3
            nw154_ = C3()
            nw154_.ctor__(d_930_v148_, d_728_v0_)
            d_931_v149_ = nw154_
            d_932_v150_: _dafny.Map
            d_932_v150_ = _dafny.Map({d_929_v147_: d_931_v149_})
            d_728_v0_ = ((d_929_v147_) not in (d_932_v150_) if d_728_v0_ else (d_931_v149_).f10)
            d_933_v151_: _dafny.MultiSet
            d_933_v151_ = _dafny.MultiSet([(d_931_v149_).f10, (d_931_v149_).f10, (d_928_v146_).f11])
            d_934_v152_: _dafny.Seq
            d_934_v152_ = _dafny.SeqWithoutIsStrInference([(d_933_v151_).isdisjoint(d_933_v151_), (d_931_v149_).f10, True, (d_928_v146_).f11, (d_931_v149_).f10])
            if (d_934_v152_)[default__.safeIndex(p0, len(d_934_v152_))]:
                d_935_v154_: _dafny.Map
                d_935_v154_ = _dafny.Map({(d_928_v146_).f12: (0) - ((d_928_v146_).f12)})
                d_936_v155_: _dafny.Set
                def iife53_():
                    coll31_ = _dafny.Map()
                    compr_31_: int
                    for compr_31_ in (d_935_v154_).keys.Elements:
                        d_937_v153_: int = compr_31_
                        if (d_937_v153_) in (d_935_v154_):
                            coll31_[default__.safeModuloInt(d_937_v153_, p1)] = (d_928_v146_).f11
                    return _dafny.Map(coll31_)
                d_936_v155_ = _dafny.Set({len(iife53_()
                )})
                (globalState).f2 = (0) - (default__.fm24((d_936_v155_ if True else d_936_v155_), globalState))
                d_938_v156_: _dafny.MultiSet
                d_938_v156_ = _dafny.MultiSet([d_929_v147_, d_929_v147_])
                d_939_v157_: _dafny.Seq
                d_939_v157_ = _dafny.SeqWithoutIsStrInference([d_938_v156_])
                d_940_v158_: _dafny.Seq
                d_940_v158_ = _dafny.SeqWithoutIsStrInference([d_929_v147_, d_929_v147_, d_929_v147_])
                d_938_v156_ = ((d_939_v157_)[default__.safeIndex(d_927_cf6_, len(d_939_v157_))]) - ((_dafny.MultiSet([d_929_v147_, (d_940_v158_)[default__.safeIndex(-888, len(d_940_v158_))], d_929_v147_])) | (d_938_v156_))
                d_941_v159_: _dafny.Seq
                d_941_v159_ = _dafny.SeqWithoutIsStrInference([(d_931_v149_).f10])
                (globalState).f2 = len((d_941_v159_))
                d_927_cf6_ = (d_928_v146_).f12
                d_919_v139_ = default__.fm8(826, globalState)
            elif True:
                (globalState).f2 = (0) - ((d_928_v146_).f12)
                d_942_v160_: _dafny.Map
                d_942_v160_ = _dafny.Map({(d_931_v149_).f10: (d_919_v139_) == ((d_799_v57_).cf22)})
                d_728_v0_ = ((d_942_v160_)[(d_928_v146_).f11] if ((d_928_v146_).f11) in (d_942_v160_) else not(not((((d_921_v141_)[(0) - ((d_928_v146_).f12)] if ((0) - ((d_928_v146_).f12)) in (d_921_v141_) else (d_931_v149_).f10)) and ((d_931_v149_).f10))))
                d_943_v161_: _dafny.Array
                nw155_ = _dafny.Array(int(0), 18)
                d_943_v161_ = nw155_
                d_943_v161_ = d_943_v161_
                d_942_v160_ = (d_942_v160_).set(((self).f6) in (d_919_v139_), (d_928_v146_).f11)
                d_944_v162_: _dafny.Map
                d_944_v162_ = _dafny.Map({(d_928_v146_).f12: p0})
                d_945_v164_: _dafny.Map
                d_945_v164_ = d_944_v162_
                d_946_v165_: _dafny.Array
                nw156_ = _dafny.Array(None, 5)
                nw156_[int(0)] = (_dafny.Map({(d_928_v146_).f12: len(d_919_v139_)})) | (_dafny.Map({57: p1}))
                nw156_[int(1)] = d_944_v162_
                def iife54_():
                    coll32_ = _dafny.Map()
                    compr_32_: int
                    for compr_32_ in _dafny.IntegerRange(572, 990):
                        d_947_v163_: int = compr_32_
                        if ((572) <= (d_947_v163_)) and ((d_947_v163_) < (990)):
                            coll32_[(d_947_v163_) * (p1)] = p0
                    return _dafny.Map(coll32_)
                nw156_[int(2)] = iife54_()
                
                nw156_[int(3)] = (d_944_v162_) | ((d_945_v164_))
                nw156_[int(4)] = d_944_v162_
                d_946_v165_ = nw156_
                index130_ = default__.safeIndex(473, (d_946_v165_).length(0))
                (d_946_v165_)[index130_] = d_944_v162_
                index131_ = default__.safeIndex(473, (d_946_v165_).length(0))
                (d_946_v165_)[index131_] = (d_944_v162_).set(p1, (len(d_934_v152_) if (d_931_v149_).f10 else -935))
            if (d_931_v149_).f10:
                d_948_v166_: int
                d_949_v167_: str
                out35_: int
                out36_: str
                out35_, out36_ = (d_928_v146_).m1(d_920_v140_, ((d_928_v146_).f11) or ((d_928_v146_).f11), d_919_v139_, globalState)
                d_948_v166_ = out35_
                d_949_v167_ = out36_
                (globalState).f2 = (p0) - ((d_928_v146_).f12)
                d_950_v168_: bool
                d_951_v169_: bool
                d_952_v170_: str
                d_953_v171_: bool
                out37_: bool
                out38_: bool
                out39_: str
                out40_: bool
                out37_, out38_, out39_, out40_ = (d_931_v149_).m5(p1, (d_931_v149_).f10, globalState)
                d_950_v168_ = out37_
                d_951_v169_ = out38_
                d_952_v170_ = out39_
                d_953_v171_ = out40_
                d_954_v172_: D9
                d_954_v172_ = D9_DC19(d_923_v143_, (0) - (((d_930_v148_)[(d_928_v146_).f12] if ((d_928_v146_).f12) in (d_930_v148_) else (d_933_v151_).cardinality)), default__.fm9(len(d_920_v140_), globalState))
                d_954_v172_ = d_954_v172_
                d_955_v173_: _dafny.Array
                nw157_ = _dafny.Array(None, 2)
                nw157_[int(0)] = d_933_v151_
                nw157_[int(1)] = d_933_v151_
                d_955_v173_ = nw157_
                d_956_v174_: C0
                nw158_ = C0()
                nw158_.ctor__(p1)
                d_956_v174_ = nw158_
                d_957_v175_: _dafny.Set
                d_957_v175_ = _dafny.Set({(d_928_v146_).f11})
                rhs127_ = (d_949_v167_ if d_953_v171_ else d_952_v170_)
                rhs128_ = ((d_957_v175_) | (d_957_v175_)) == (d_957_v175_)
                rhs129_ = d_955_v173_
                rhs130_ = d_956_v174_
                d_949_v167_ = rhs127_
                d_950_v168_ = rhs128_
                d_955_v173_ = rhs129_
                d_956_v174_ = rhs130_
            elif True:
                d_958_v176_: _dafny.Set
                d_958_v176_ = _dafny.Set({(d_928_v146_).f12, (d_920_v140_)[default__.safeIndex((0) - ((0) - (p1)), len(d_920_v140_))]})
                d_959_v177_: _dafny.Map
                d_959_v177_ = _dafny.Map({(d_928_v146_).f11: d_958_v176_})
                d_960_v178_: _dafny.Map
                d_960_v178_ = _dafny.Map({(d_920_v140_) == (d_920_v140_): (default__.fm38(len(_dafny.Map({(d_931_v149_).f10: (d_934_v152_)[default__.safeIndex(492, len(d_934_v152_))]})), (d_934_v152_)[default__.safeIndex(p1, len(d_934_v152_))], (d_931_v149_).f10, (d_931_v149_).f10, globalState)) | (((d_959_v177_)[(d_931_v149_).f10] if ((d_931_v149_).f10) in (d_959_v177_) else _dafny.Set({(d_928_v146_).f12, p1})))})
                def iife55_():
                    coll33_ = _dafny.Set()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(121, 74):
                        d_961_v179_: int = compr_33_
                        if ((121) <= (d_961_v179_)) and ((d_961_v179_) < (74)):
                            def iife56_():
                                coll34_ = _dafny.Map()
                                compr_34_: int
                                for compr_34_ in _dafny.IntegerRange(124, 935):
                                    d_962_v180_: int = compr_34_
                                    if ((124) <= (d_962_v180_)) and ((d_962_v180_) < (935)):
                                        coll34_[(d_962_v180_) - (d_927_cf6_)] = p1
                                return _dafny.Map(coll34_)
                            coll33_ = coll33_.union(_dafny.Set([default__.safeModuloInt(d_961_v179_, len(iife56_()
))]))
                    return _dafny.Set(coll33_)
                d_959_v177_ = (d_959_v177_).set(not((d_927_cf6_) not in (d_958_v176_)), (iife55_()
                 if False else d_958_v176_))
                d_963_v181_: _dafny.Set
                d_963_v181_ = _dafny.Set({(self).f6})
                d_964_v182_: _dafny.Set
                d_964_v182_ = _dafny.Set({(d_931_v149_).f10, (d_963_v181_).isdisjoint(d_963_v181_), False})
                d_964_v182_ = d_964_v182_
                (globalState).f2 = d_927_cf6_
                d_958_v176_ = d_958_v176_
                d_965_v183_: _dafny.Set
                d_965_v183_ = d_964_v182_
                d_965_v183_ = (_dafny.Set({(d_931_v149_).f10})) - (d_964_v182_)
        elif source16_.is_DC5:
            d_966___mcc_h10_ = source16_.cf7
            d_967___mcc_h11_ = source16_.cf8
            d_968___mcc_h12_ = source16_.cf9
            d_969___mcc_h13_ = source16_.cf10
            d_970___mcc_h14_ = source16_.cf11
            d_971_cf11_ = d_970___mcc_h14_
            d_972_cf10_ = d_969___mcc_h13_
            d_973_cf9_ = d_968___mcc_h12_
            d_974_cf8_ = d_967___mcc_h11_
            d_975_cf7_ = d_966___mcc_h10_
            d_976_v184_: _dafny.Array
            nw159_ = _dafny.Array(_dafny.MultiSet({}), 16)
            d_976_v184_ = nw159_
            d_977_v185_: _dafny.Array
            nw160_ = _dafny.Array(int(0), 11)
            d_977_v185_ = nw160_
            index132_ = default__.safeIndex(389, (d_976_v184_).length(0))
            (d_976_v184_)[index132_] = (_dafny.MultiSet([d_977_v185_, d_977_v185_, d_977_v185_])).set(d_977_v185_, default__.abs(p1))
            d_978_v186_: _dafny.MultiSet
            d_978_v186_ = _dafny.MultiSet([d_977_v185_, d_977_v185_, d_977_v185_, d_977_v185_])
            index133_ = default__.safeIndex(389, (d_976_v184_).length(0))
            (d_976_v184_)[index133_] = ((d_978_v186_) - (d_978_v186_) if d_974_cf8_ else (d_978_v186_) | (d_978_v186_))
            d_979_v187_: _dafny.MultiSet
            d_979_v187_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "polxqwadi"))])
            d_979_v187_ = ((d_979_v187_) | (d_979_v187_)).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uiiokhs")), default__.abs(-980))
            d_919_v139_ = (d_799_v57_).cf22
            d_980_v188_: _dafny.Array
            def lambda43_(d_981_cf8_):
                def lambda44_(d_982_i17_):
                    return _dafny.Set({d_981_cf8_})

                return lambda44_

            init24_ = lambda43_(d_974_cf8_)
            nw161_ = _dafny.Array(None, 5)
            for i0_24_ in range(nw161_.length(0)):
                nw161_[i0_24_] = init24_(i0_24_)
            d_980_v188_ = nw161_
            index134_ = default__.safeIndex(560, (d_980_v188_).length(0))
            (d_980_v188_)[index134_] = default__.fm16(p0, _dafny.CodePoint('l'), d_975_cf7_, globalState)
            d_983_v189_: _dafny.Set
            d_983_v189_ = _dafny.Set({p1})
            d_984_v190_: _dafny.Seq
            d_984_v190_ = _dafny.SeqWithoutIsStrInference([not(False)])
            d_985_v191_: _dafny.Set
            d_985_v191_ = _dafny.Set({(d_983_v189_).issubset(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([p1 for d_986_i18_ in range(default__.abs(544))]))})), d_972_cf10_, (d_984_v190_)[default__.safeIndex(p0, len(d_984_v190_))]})
            index135_ = default__.safeIndex(560, (d_980_v188_).length(0))
            (d_980_v188_)[index135_] = d_985_v191_
        elif True:
            d_987___mcc_h15_ = source16_.cf5
            d_988_cf5_ = d_987___mcc_h15_
            d_989_v192_: C2
            nw162_ = C2()
            nw162_.ctor__(d_728_v0_, default__.safeDivisionInt(p0, 130), d_922_v142_)
            d_989_v192_ = nw162_
            if (d_989_v192_).f11:
                d_728_v0_ = d_728_v0_
                d_990_v193_: _dafny.Array
                nw163_ = _dafny.Array(None, 20)
                nw163_[int(0)] = False
                nw163_[int(1)] = (d_989_v192_).f11
                nw163_[int(2)] = d_728_v0_
                nw163_[int(3)] = d_728_v0_
                nw163_[int(4)] = (d_989_v192_).f11
                nw163_[int(5)] = (d_989_v192_).f11
                nw163_[int(6)] = (d_989_v192_).f11
                nw163_[int(7)] = (d_989_v192_).f11
                nw163_[int(8)] = (d_989_v192_).f11
                nw163_[int(9)] = (d_989_v192_).f11
                nw163_[int(10)] = False
                nw163_[int(11)] = (d_989_v192_).f11
                nw163_[int(12)] = (d_989_v192_).f11
                nw163_[int(13)] = (d_989_v192_).f11
                nw163_[int(14)] = True
                nw163_[int(15)] = d_728_v0_
                nw163_[int(16)] = d_728_v0_
                nw163_[int(17)] = d_728_v0_
                nw163_[int(18)] = (d_989_v192_).f11
                nw163_[int(19)] = (d_989_v192_).f11
                d_990_v193_ = nw163_
                d_991_v194_: _dafny.Map
                d_991_v194_ = _dafny.Map({D8_DC15(d_990_v193_): D9_DC19(d_923_v143_, p0, (d_989_v192_).f12)})
                d_921_v141_ = (d_921_v141_).set((0) - ((len(d_991_v194_)) + ((d_989_v192_).f12)), (d_989_v192_).f11)
                d_992_v195_: D8
                d_992_v195_ = D8_DC15(d_990_v193_)
                pat_let_tv24_ = d_990_v193_
                def iife57_(_pat_let11_0):
                    def iife58_(d_993_dt__update__tmp_h3_):
                        def iife59_(_pat_let12_0):
                            def iife60_(d_994_dt__update_hcf27_h0_):
                                return D8_DC15(d_994_dt__update_hcf27_h0_)
                            return iife60_(_pat_let12_0)
                        return iife59_(pat_let_tv24_)
                    return iife58_(_pat_let11_0)
                d_992_v195_ = iife57_(d_992_v195_)
                d_728_v0_ = (p0) <= ((d_989_v192_).f12)
                d_728_v0_ = ((d_989_v192_).f11 if ((0) - ((d_920_v140_)[default__.safeIndex(p0, len(d_920_v140_))])) >= (len(d_919_v139_)) else ((d_989_v192_).f11) == (default__.fm23(d_919_v139_, (_dafny.MultiSet([(d_989_v192_).f12])).cardinality, p1, (d_989_v192_).f11, globalState)))
            elif True:
                rhs131_ = not(not(not((d_989_v192_).f11)))
                rhs132_ = d_919_v139_
                d_728_v0_ = rhs131_
                d_919_v139_ = rhs132_
                def iife61_():
                    coll35_ = _dafny.Set()
                    compr_35_: int
                    for compr_35_ in (_dafny.SeqWithoutIsStrInference([p1])).Elements:
                        d_995_v196_: int = compr_35_
                        if (d_995_v196_) in (_dafny.SeqWithoutIsStrInference([p1])):
                            coll35_ = coll35_.union(_dafny.Set([(d_995_v196_) * (489)]))
                    return _dafny.Set(coll35_)
                (globalState).f2 = (len((_dafny.Set({36, len(d_919_v139_)})).intersection(iife61_()
                ))) - (p1)
                d_728_v0_ = d_728_v0_
                d_996_v197_: _dafny.Map
                d_996_v197_ = _dafny.Map({(d_989_v192_).f12: len(d_921_v141_)})
                (globalState).f2 = default__.safeDivisionInt(((d_996_v197_)[(d_989_v192_).f12] if ((d_989_v192_).f12) in (d_996_v197_) else (d_989_v192_).f12), p1)
                d_997_v198_: str
                d_997_v198_ = _dafny.CodePoint('e')
                d_997_v198_ = (self).f6
            (globalState).f2 = (0) - (((d_989_v192_).f12) + (len(_dafny.Set({False}))))
            (globalState).f2 = ((p1) + (p1)) * (((0) - (-916)) - (p0))
        d_998_v199_: _dafny.Array
        nw164_ = _dafny.Array(False, 7)
        d_998_v199_ = nw164_
        index136_ = default__.safeIndex(840, (d_998_v199_).length(0))
        (d_998_v199_)[index136_] = (_dafny.SeqWithoutIsStrInference([(self).f6 for d_999_i19_ in range(default__.abs(-512))])) < (d_919_v139_)
        d_1000_v200_: _dafny.Map
        d_1000_v200_ = _dafny.Map({p1: p0})
        d_1001_v201_: D3
        d_1001_v201_ = D3_DC6((self).f6)
        d_1002_v202_: _dafny.Array
        nw165_ = _dafny.Array(int(0), 11)
        d_1002_v202_ = nw165_
        index137_ = default__.safeIndex(52, (d_1002_v202_).length(0))
        (d_1002_v202_)[index137_] = p1
        index138_ = default__.safeIndex(840, (d_998_v199_).length(0))
        index139_ = default__.safeIndex(52, (d_1002_v202_).length(0))
        rhs133_ = d_728_v0_
        rhs134_ = d_1000_v200_
        rhs135_ = ((p0) - (p1)) - (((0) - (p1)) - (468))
        rhs136_ = d_1001_v201_
        rhs137_ = p1
        lhs77_ = d_998_v199_
        lhs78_ = default__.safeIndex(840, (d_998_v199_).length(0))
        lhs79_ = globalState
        lhs80_ = d_1002_v202_
        lhs81_ = default__.safeIndex(52, (d_1002_v202_).length(0))
        lhs77_[lhs78_] = rhs133_
        d_1000_v200_ = rhs134_
        lhs79_.f2 = rhs135_
        d_1001_v201_ = rhs136_
        lhs80_[lhs81_] = rhs137_
        index140_ = default__.safeIndex(840, (d_998_v199_).length(0))
        (d_998_v199_)[index140_] = not(True)
        d_1003_v203_: _dafny.MultiSet
        d_1003_v203_ = _dafny.MultiSet([p0])
        d_1004_v204_: C3
        nw166_ = C3()
        nw166_.ctor__(d_1003_v203_, d_728_v0_)
        d_1004_v204_ = nw166_
        d_1005_v205_: _dafny.Map
        d_1005_v205_ = _dafny.Map({d_1004_v204_: (self).f6})
        index141_ = default__.safeIndex(840, (d_998_v199_).length(0))
        (d_998_v199_)[index141_] = (d_1005_v205_) != (_dafny.Map({d_1004_v204_: (D3_DC6((self).f6)).cf12}))

    @property
    def f6(self):
        return self._f6

class C6:
    def  __init__(self):
        self.f5: int = int(0)
        self._f4: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f4, f5):
        (self)._f4 = f4
        (self).f5 = f5

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1006_i0_: int
        d_1006_i0_ = 0
        with _dafny.label("5"):
            while p0:
                with _dafny.c_label("5"):
                    if (d_1006_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_1006_i0_ = (d_1006_i0_) + (1)
                    d_1007_v0_: _dafny.Array
                    def lambda45_(d_1008_p0_):
                        def lambda46_(d_1009_i1_):
                            return D1_DC2(self.f5, d_1008_p0_, self.f5)

                        return lambda46_

                    init25_ = lambda45_(p0)
                    nw167_ = _dafny.Array(None, 10)
                    for i0_25_ in range(nw167_.length(0)):
                        nw167_[i0_25_] = init25_(i0_25_)
                    d_1007_v0_ = nw167_
                    d_1010_v1_: D1
                    d_1010_v1_ = D1_DC2(self.f5, p0, self.f5)
                    index142_ = default__.safeIndex(620, (d_1007_v0_).length(0))
                    (d_1007_v0_)[index142_] = d_1010_v1_
                    index143_ = default__.safeIndex(620, (d_1007_v0_).length(0))
                    (d_1007_v0_)[index143_] = d_1010_v1_
                    d_1011_v2_: str
                    d_1011_v2_ = _dafny.CodePoint('x')
                    d_1012_v3_: C5
                    nw168_ = C5()
                    nw168_.ctor__(d_1011_v2_)
                    d_1012_v3_ = nw168_
                    d_1013_v4_: _dafny.Map
                    d_1013_v4_ = _dafny.Map({p0: False})
                    d_1014_v5_: _dafny.Seq
                    d_1014_v5_ = _dafny.SeqWithoutIsStrInference([(0) - (self.f5), self.f5])
                    d_1015_v6_: _dafny.Map
                    d_1015_v6_ = _dafny.Map({d_1013_v4_: d_1014_v5_})
                    d_1015_v6_ = (d_1015_v6_).set(default__.fm39(False, globalState), d_1014_v5_)
                    d_1016_v7_: _dafny.Array
                    nw169_ = _dafny.Array(False, 21)
                    d_1016_v7_ = nw169_
                    d_1017_v8_: _dafny.Set
                    d_1017_v8_ = _dafny.Set({d_1016_v7_, d_1016_v7_, d_1016_v7_, d_1016_v7_, d_1016_v7_})
                    d_1018_v9_: _dafny.Seq
                    d_1018_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eueyat"))
                    d_1019_v10_: D6
                    d_1019_v10_ = D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuxdxnku")))
                    d_1020_v11_: _dafny.Map
                    d_1020_v11_ = _dafny.Map({d_1017_v8_: len((d_1018_v9_) + ((d_1019_v10_).cf22))})
                    d_1021_v12_: _dafny.Seq
                    d_1021_v12_ = _dafny.SeqWithoutIsStrInference([d_1012_v3_, d_1012_v3_])
                    d_1022_v13_: _dafny.MultiSet
                    d_1022_v13_ = _dafny.MultiSet([len(d_1021_v12_), self.f5])
                    d_1023_v14_: _dafny.Seq
                    d_1023_v14_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_1020_v11_ = (d_1020_v11_).set((d_1017_v8_) - (d_1017_v8_), ((d_1022_v13_).cardinality if p0 else len(_dafny.Map({(d_1014_v5_)[default__.safeIndex(len(d_1023_v14_), len(d_1014_v5_))]: d_1013_v4_}))))
                    pass
            pass
        d_1024_v15_: _dafny.MultiSet
        d_1024_v15_ = _dafny.MultiSet([True, not(p0), p0, p0])
        d_1025_v16_: str
        d_1025_v16_ = _dafny.CodePoint('x')
        d_1026_v17_: _dafny.Map
        d_1026_v17_ = _dafny.Map({884: default__.fm21(self.f5, d_1025_v16_, p0, False, globalState)})
        d_1027_v18_: _dafny.Seq
        d_1027_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbwc"))
        d_1028_v19_: _dafny.MultiSet
        d_1028_v19_ = _dafny.MultiSet([((d_1024_v15_).cardinality) * (self.f5), self.f5, default__.safeModuloInt(len(d_1026_v17_), len(d_1027_v18_))])
        d_1029_v20_: _dafny.Array
        nw170_ = _dafny.Array(D2.default()(), 19)
        d_1029_v20_ = nw170_
        d_1030_v21_: _dafny.Seq
        d_1030_v21_ = _dafny.SeqWithoutIsStrInference([d_1029_v20_])
        rhs138_ = (d_1028_v19_).intersection(d_1028_v19_)
        rhs139_ = self.f5
        rhs140_ = _dafny.SeqWithoutIsStrInference([d_1029_v20_])
        lhs82_ = self
        d_1028_v19_ = rhs138_
        lhs82_.f5 = rhs139_
        d_1030_v21_ = rhs140_
        d_1031_v22_: _dafny.Array
        nw171_ = _dafny.Array(_dafny.Array(None, 0), 12)
        d_1031_v22_ = nw171_
        d_1032_v23_: _dafny.Array
        nw172_ = _dafny.Array(False, 28)
        d_1032_v23_ = nw172_
        index144_ = default__.safeIndex(678, (d_1031_v22_).length(0))
        (d_1031_v22_)[index144_] = d_1032_v23_
        index145_ = default__.safeIndex(678, (d_1031_v22_).length(0))
        (d_1031_v22_)[index145_] = d_1032_v23_
        hi4_ = self.f5
        for d_1033_i2_ in range(((0) - (self.f5)) - (self.f5), hi4_):
            r1 = p0
            d_1024_v15_ = d_1024_v15_
            d_1034_v24_: _dafny.Array
            def lambda47_(d_1035_v19_):
                def lambda48_(d_1036_i3_):
                    return (d_1035_v19_) - (d_1035_v19_)

                return lambda48_

            init26_ = lambda47_(d_1028_v19_)
            nw173_ = _dafny.Array(None, 2)
            for i0_26_ in range(nw173_.length(0)):
                nw173_[i0_26_] = init26_(i0_26_)
            d_1034_v24_ = nw173_
            d_1034_v24_ = d_1034_v24_
            d_1037_v25_: _dafny.Seq
            d_1037_v25_ = _dafny.SeqWithoutIsStrInference([len((self).f4), d_1033_i2_, d_1033_i2_, self.f5, len(d_1027_v18_)])
            d_1038_v26_: _dafny.Set
            d_1038_v26_ = _dafny.Set({False, p0})
            d_1039_v27_: _dafny.Map
            d_1039_v27_ = _dafny.Map({self.f5: self.f5})
            d_1040_v28_: _dafny.Array
            nw174_ = _dafny.Array(None, 21)
            nw174_[int(0)] = (d_1037_v25_) + (default__.fm12(p0, len(d_1037_v25_), d_1038_v26_, d_1039_v27_, globalState))
            nw174_[int(1)] = _dafny.SeqWithoutIsStrInference([d_1033_i2_ for d_1041_i4_ in range(default__.abs(699))])
            nw174_[int(2)] = default__.fm12(p0, len(d_1027_v18_), d_1038_v26_, d_1039_v27_, globalState)
            nw174_[int(3)] = (d_1037_v25_) + (d_1037_v25_)
            nw174_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1033_i2_ for d_1042_i5_ in range(default__.abs(240))])
            nw174_[int(5)] = _dafny.SeqWithoutIsStrInference([self.f5])
            nw174_[int(6)] = d_1037_v25_
            nw174_[int(7)] = ((_dafny.SeqWithoutIsStrInference([self.f5, d_1033_i2_])).set(default__.safeIndex(d_1033_i2_, len(_dafny.SeqWithoutIsStrInference([self.f5, d_1033_i2_]))), self.f5)) + (d_1037_v25_)
            nw174_[int(8)] = d_1037_v25_
            nw174_[int(9)] = d_1037_v25_
            nw174_[int(10)] = d_1037_v25_
            nw174_[int(11)] = d_1037_v25_
            nw174_[int(12)] = d_1037_v25_
            nw174_[int(13)] = d_1037_v25_
            nw174_[int(14)] = d_1037_v25_
            nw174_[int(15)] = _dafny.SeqWithoutIsStrInference([d_1033_i2_])
            nw174_[int(16)] = _dafny.SeqWithoutIsStrInference([len(d_1037_v25_), d_1033_i2_, d_1033_i2_, d_1033_i2_])
            nw174_[int(17)] = d_1037_v25_
            nw174_[int(18)] = d_1037_v25_
            nw174_[int(19)] = (d_1037_v25_) + (_dafny.SeqWithoutIsStrInference([729]))
            nw174_[int(20)] = (d_1037_v25_) + (d_1037_v25_)
            d_1040_v28_ = nw174_
            d_1043_v29_: _dafny.Map
            d_1043_v29_ = _dafny.Map({(d_1037_v25_)[default__.safeIndex(d_1033_i2_, len(d_1037_v25_))]: d_1040_v28_})
            d_1040_v28_ = ((d_1043_v29_)[self.f5] if (self.f5) in (d_1043_v29_) else d_1040_v28_)
        d_1044_i6_: int
        d_1044_i6_ = 0
        with _dafny.label("6"):
            while not(p0):
                with _dafny.c_label("6"):
                    if (d_1044_i6_) >= (100):
                        raise _dafny.Break("6")
                    d_1044_i6_ = (d_1044_i6_) + (1)
                    d_1045_v30_: _dafny.Map
                    d_1045_v30_ = _dafny.Map({p0: d_1032_v23_})
                    d_1046_v31_: _dafny.Array
                    nw175_ = _dafny.Array(_dafny.Map({}), 3)
                    d_1046_v31_ = nw175_
                    d_1047_v32_: C2
                    nw176_ = C2()
                    nw176_.ctor__(default__.fm21(self.f5, d_1025_v16_, p0, p0, globalState), self.f5, d_1046_v31_)
                    d_1047_v32_ = nw176_
                    d_1048_v33_: _dafny.Seq
                    d_1048_v33_ = _dafny.SeqWithoutIsStrInference([len((d_1045_v30_).set(p0, (d_1031_v22_)[default__.safeIndex(678, (d_1031_v22_).length(0))])), len(_dafny.SeqWithoutIsStrInference([d_1047_v32_])), self.f5, (d_1047_v32_).f12])
                    d_1049_v34_: _dafny.Seq
                    d_1049_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1048_v33_)])
                    d_1050_v35_: C3
                    nw177_ = C3()
                    nw177_.ctor__((_dafny.MultiSet([(d_1028_v19_).cardinality, self.f5])) | ((d_1049_v34_)[default__.safeIndex((d_1047_v32_).f12, len(d_1049_v34_))]), p0)
                    d_1050_v35_ = nw177_
                    d_1051_v36_: _dafny.Seq
                    d_1051_v36_ = _dafny.SeqWithoutIsStrInference([(d_1047_v32_).f11])
                    d_1052_v37_: D4
                    d_1052_v37_ = D4_DC9((_dafny.MultiSet(d_1051_v36_)).cardinality, d_1025_v16_, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(d_1047_v32_).f12, self.f5]))]), self.f5, p0)
                    d_1025_v16_ = (d_1052_v37_).cf17
                    d_1053_v38_: _dafny.Seq
                    d_1053_v38_ = _dafny.SeqWithoutIsStrInference([default__.fm40((d_1050_v35_).f10, (d_1047_v32_).f12, d_1025_v16_, (d_1047_v32_).f11, globalState), (_dafny.MultiSet([(d_1047_v32_).f11])).set((d_1047_v32_).f11, default__.abs(self.f5)), d_1024_v15_])
                    d_1054_v39_: _dafny.Set
                    d_1054_v39_ = _dafny.Set({(d_1047_v32_).f12})
                    d_1055_v40_: _dafny.Seq
                    d_1055_v40_ = _dafny.SeqWithoutIsStrInference([d_1054_v39_, d_1054_v39_, d_1054_v39_, d_1054_v39_])
                    d_1056_v41_: _dafny.Array
                    nw178_ = _dafny.Array(None, 5)
                    nw178_[int(0)] = (0) - ((d_1047_v32_).f12)
                    nw178_[int(1)] = ((d_1024_v15_).intersection((d_1053_v38_)[default__.safeIndex(308, len(d_1053_v38_))])).cardinality
                    nw178_[int(2)] = len(d_1055_v40_)
                    nw178_[int(3)] = self.f5
                    nw178_[int(4)] = (d_1047_v32_).f12
                    d_1056_v41_ = nw178_
                    index146_ = default__.safeIndex(999, (d_1056_v41_).length(0))
                    (d_1056_v41_)[index146_] = self.f5
                    d_1057_v42_: _dafny.Set
                    d_1057_v42_ = _dafny.Set({p0})
                    index147_ = default__.safeIndex(999, (d_1056_v41_).length(0))
                    rhs141_ = (d_1047_v32_).f12
                    rhs142_ = len(((d_1057_v42_).intersection(d_1057_v42_)) - (_dafny.Set({True, False})))
                    rhs143_ = len((_dafny.SeqWithoutIsStrInference([((d_1047_v32_).f12) * (len(d_1026_v17_)) for d_1058_i7_ in range(default__.abs(748))])).set(default__.safeIndex(self.f5, len(_dafny.SeqWithoutIsStrInference([((d_1047_v32_).f12) * (len(d_1026_v17_)) for d_1059_i7_ in range(default__.abs(748))]))), (0) - (((0) - (self.f5)) * ((d_1047_v32_).f12))))
                    rhs144_ = (d_1050_v35_).f10
                    lhs83_ = globalState
                    lhs84_ = self
                    lhs85_ = d_1056_v41_
                    lhs86_ = default__.safeIndex(999, (d_1056_v41_).length(0))
                    lhs83_.f2 = rhs141_
                    lhs84_.f5 = rhs142_
                    lhs85_[lhs86_] = rhs143_
                    r1 = rhs144_
                    index148_ = default__.safeIndex(848, (d_1032_v23_).length(0))
                    (d_1032_v23_)[index148_] = (self.f5) != ((d_1056_v41_)[default__.safeIndex(999, (d_1056_v41_).length(0))])
                    index149_ = default__.safeIndex(848, (d_1032_v23_).length(0))
                    (d_1032_v23_)[index149_] = ((d_1047_v32_).f11 if ((d_1047_v32_).f12) > (self.f5) else (d_1025_v16_) in (d_1027_v18_))
                    pass
            pass
        d_1060_v43_: _dafny.Seq
        d_1060_v43_ = _dafny.SeqWithoutIsStrInference([(self.f5) + (self.f5)])
        d_1061_v44_: _dafny.Map
        d_1061_v44_ = _dafny.Map({self.f5: d_1060_v43_})
        d_1060_v43_ = ((d_1060_v43_) + (((d_1061_v44_)[self.f5] if (self.f5) in (d_1061_v44_) else d_1060_v43_))) + (d_1060_v43_)
        d_1062_v45_: _dafny.Map
        d_1062_v45_ = _dafny.Map({len(d_1027_v18_): d_1025_v16_})
        d_1063_v46_: _dafny.MultiSet
        d_1063_v46_ = _dafny.MultiSet([d_1062_v45_])
        r0 = default__.fm9(((d_1063_v46_).set(d_1062_v45_, default__.abs(self.f5))).cardinality, globalState)
        r1 = (len((d_1027_v18_) + (d_1027_v18_))) <= (self.f5)
        return r0, r1

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        d_1064_v0_: _dafny.Array
        nw179_ = _dafny.Array(False, 18)
        d_1064_v0_ = nw179_
        d_1065_v1_: D8
        d_1065_v1_ = D8_DC15(d_1064_v0_)
        source17_ = d_1065_v1_
        if source17_.is_DC16:
            d_1066___mcc_h0_ = source17_.cf28
            d_1067___mcc_h1_ = source17_.cf29
            d_1068___mcc_h2_ = source17_.cf30
            d_1069___mcc_h3_ = source17_.cf31
            d_1070_cf31_ = d_1069___mcc_h3_
            d_1071_cf30_ = d_1068___mcc_h2_
            d_1072_cf29_ = d_1067___mcc_h1_
            d_1073_cf28_ = d_1066___mcc_h0_
            d_1074_v2_: _dafny.Array
            nw180_ = _dafny.Array(int(0), 28)
            d_1074_v2_ = nw180_
            d_1075_v3_: _dafny.Set
            d_1075_v3_ = _dafny.Set({d_1072_cf29_})
            index150_ = default__.safeIndex(584, (d_1074_v2_).length(0))
            (d_1074_v2_)[index150_] = len((d_1075_v3_ if d_1071_cf30_ else d_1075_v3_))
            index151_ = default__.safeIndex(584, (d_1074_v2_).length(0))
            (d_1074_v2_)[index151_] = d_1070_cf31_
            d_1073_cf28_ = d_1073_cf28_
            d_1076_v4_: str
            d_1076_v4_ = _dafny.CodePoint('x')
            d_1077_v5_: D3
            d_1077_v5_ = D3_DC6(d_1076_v4_)
            d_1078_v6_: C5
            nw181_ = C5()
            nw181_.ctor__((d_1077_v5_).cf12)
            d_1078_v6_ = nw181_
            d_1079_v7_: _dafny.MultiSet
            d_1079_v7_ = _dafny.MultiSet([(d_1074_v2_)[default__.safeIndex(584, (d_1074_v2_).length(0))], p1])
            index152_ = default__.safeIndex(584, (d_1074_v2_).length(0))
            (d_1074_v2_)[index152_] = default__.safeModuloInt(p1, (self.f5) * (((d_1079_v7_)[self.f5] if (self.f5) in (d_1079_v7_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))))))
        elif source17_.is_DC15:
            d_1080___mcc_h4_ = source17_.cf27
            d_1081_cf27_ = d_1080___mcc_h4_
            d_1082_v8_: bool
            d_1082_v8_ = True
            d_1082_v8_ = (default__.fm41(self.f5, self.f5, globalState)).cf30
            d_1083_v9_: _dafny.Map
            d_1083_v9_ = _dafny.Map({p3: p2})
            d_1084_v10_: _dafny.Map
            d_1084_v10_ = _dafny.Map({p1: p3})
            d_1085_v11_: _dafny.Array
            def lambda49_(d_1086_v9_):
                def lambda50_(d_1087_i0_):
                    return d_1086_v9_

                return lambda50_

            init27_ = lambda49_(d_1083_v9_)
            nw182_ = _dafny.Array(None, 26)
            for i0_27_ in range(nw182_.length(0)):
                nw182_[i0_27_] = init27_(i0_27_)
            d_1085_v11_ = nw182_
            d_1088_v12_: C2
            nw183_ = C2()
            nw183_.ctor__(((d_1083_v9_)[(((self).f4)[not(d_1082_v8_)] if (not(d_1082_v8_)) in ((self).f4) else p1)] if ((((self).f4)[not(d_1082_v8_)] if (not(d_1082_v8_)) in ((self).f4) else p1)) in (d_1083_v9_) else p2), len(d_1084_v10_), d_1085_v11_)
            d_1088_v12_ = nw183_
            d_1089_v13_: str
            d_1089_v13_ = _dafny.CodePoint('l')
            d_1090_v14_: _dafny.Seq
            d_1090_v14_ = _dafny.SeqWithoutIsStrInference([d_1089_v13_])
            d_1091_v15_: _dafny.Seq
            d_1091_v15_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1089_v13_]), d_1090_v14_])
            (globalState).f2 = (0) - (default__.safeModuloInt((0) - (default__.fm17(d_1091_v15_, (0) - (p1), len(d_1083_v9_), p1, globalState)), p1))
            d_1092_v16_: _dafny.Array
            nw184_ = _dafny.Array(None, 5)
            nw184_[int(0)] = (d_1088_v12_).f12
            nw184_[int(1)] = p3
            nw184_[int(2)] = (p3 if d_1082_v8_ else (0) - (len(d_1090_v14_)))
            nw184_[int(3)] = (p3) + (p1)
            nw184_[int(4)] = p3
            d_1092_v16_ = nw184_
            def lambda51_(d_1093_i1_):
                return (d_1093_i1_) + (self.f5)

            init28_ = lambda51_
            nw185_ = _dafny.Array(None, 21)
            for i0_28_ in range(nw185_.length(0)):
                nw185_[i0_28_] = init28_(i0_28_)
            d_1092_v16_ = nw185_
        elif True:
            d_1094___mcc_h5_ = source17_.cf32
            d_1095_cf32_ = d_1094___mcc_h5_
            d_1096_v17_: C5
            nw186_ = C5()
            nw186_.ctor__(default__.fm26(self.f5, (0) - ((0) - ((((self).f4)[p2] if (p2) in ((self).f4) else p1))), p2, globalState))
            d_1096_v17_ = nw186_
            d_1097_v18_: _dafny.Map
            d_1097_v18_ = _dafny.Map({p3: (d_1096_v17_).f6})
            d_1098_v19_: _dafny.Seq
            d_1098_v19_ = _dafny.SeqWithoutIsStrInference([p1, -934, len(d_1097_v18_)])
            d_1099_v20_: _dafny.Set
            d_1099_v20_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([self.f5, p3, len(_dafny.SeqWithoutIsStrInference([228 for d_1100_i2_ in range(default__.abs(467))])), p1])) + (d_1098_v19_)})
            d_1099_v20_ = d_1099_v20_
            d_1101_v21_: _dafny.Array
            nw187_ = _dafny.Array(_dafny.MultiSet({}), 29)
            d_1101_v21_ = nw187_
            d_1102_v22_: _dafny.Seq
            d_1102_v22_ = _dafny.SeqWithoutIsStrInference([d_1101_v21_])
            d_1103_v23_: _dafny.Array
            nw188_ = _dafny.Array(None, 18)
            nw188_[int(0)] = d_1101_v21_
            nw188_[int(1)] = d_1101_v21_
            nw188_[int(2)] = d_1101_v21_
            nw188_[int(3)] = d_1101_v21_
            nw188_[int(4)] = d_1101_v21_
            nw188_[int(5)] = d_1101_v21_
            nw188_[int(6)] = d_1101_v21_
            nw188_[int(7)] = d_1101_v21_
            nw188_[int(8)] = (d_1101_v21_ if p2 else d_1101_v21_)
            nw188_[int(9)] = (d_1102_v22_)[default__.safeIndex(p3, len(d_1102_v22_))]
            nw188_[int(10)] = d_1101_v21_
            nw188_[int(11)] = d_1101_v21_
            nw188_[int(12)] = d_1101_v21_
            nw188_[int(13)] = (d_1102_v22_)[default__.safeIndex(self.f5, len(d_1102_v22_))]
            nw188_[int(14)] = d_1101_v21_
            nw188_[int(15)] = (d_1101_v21_ if True else d_1101_v21_)
            nw188_[int(16)] = d_1101_v21_
            nw188_[int(17)] = d_1101_v21_
            d_1103_v23_ = nw188_
            index153_ = default__.safeIndex(460, (d_1103_v23_).length(0))
            (d_1103_v23_)[index153_] = d_1101_v21_
            index154_ = default__.safeIndex(460, (d_1103_v23_).length(0))
            (d_1103_v23_)[index154_] = d_1101_v21_
            d_1104_v24_: _dafny.Map
            d_1104_v24_ = _dafny.Map({p3: p2})
            if ((p2) or (p2)) == (((d_1104_v24_)[self.f5] if (self.f5) in (d_1104_v24_) else p2)):
                d_1105_v25_: bool
                d_1105_v25_ = False
                d_1105_v25_ = p2
                d_1106_v26_: C0
                nw189_ = C0()
                nw189_.ctor__((self.f5) - (self.f5))
                d_1106_v26_ = nw189_
                d_1107_v27_: _dafny.Seq
                d_1107_v27_ = _dafny.SeqWithoutIsStrInference([default__.fm25(False, len(d_1098_v19_), globalState)])
                d_1108_v28_: _dafny.Set
                d_1108_v28_ = _dafny.Set({p1})
                rhs145_ = default__.fm17(d_1107_v27_, (d_1106_v26_).f13, default__.fm9(len(d_1108_v28_), globalState), (0) - (p1), globalState)
                rhs146_ = d_1106_v26_
                lhs87_ = globalState
                lhs87_.f2 = rhs145_
                d_1106_v26_ = rhs146_
                d_1105_v25_ = d_1105_v25_
                d_1109_v29_: _dafny.Map
                d_1109_v29_ = _dafny.Map({d_1105_v25_: d_1105_v25_})
                d_1110_v30_: _dafny.Set
                d_1110_v30_ = _dafny.Set({((d_1109_v29_)[d_1105_v25_] if (d_1105_v25_) in (d_1109_v29_) else d_1105_v25_)})
                d_1105_v25_ = (p2) not in (d_1110_v30_)
                d_1111_v31_: _dafny.Seq
                d_1111_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myiqun"))
                d_1112_v32_: D6
                d_1112_v32_ = D6_DC11(d_1111_v31_)
                d_1113_v33_: _dafny.Seq
                d_1113_v33_ = _dafny.SeqWithoutIsStrInference([d_1112_v32_])
                d_1114_v34_: _dafny.Array
                nw190_ = _dafny.Array(None, 20)
                d_1114_v34_ = nw190_
                d_1115_v35_: _dafny.Array
                nw191_ = _dafny.Array(int(0), 28)
                d_1115_v35_ = nw191_
                d_1116_v36_: _dafny.MultiSet
                d_1116_v36_ = _dafny.MultiSet([d_1115_v35_, d_1115_v35_])
                d_1117_v37_: _dafny.Map
                d_1117_v37_ = _dafny.Map({_dafny.Map({d_1113_v33_: d_1114_v34_}): d_1116_v36_})
                d_1118_v38_: _dafny.Map
                d_1118_v38_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1112_v32_]): d_1114_v34_})
                d_1117_v37_ = (d_1117_v37_).set(d_1118_v38_, d_1116_v36_)
            elif True:
                d_1119_v39_: _dafny.Array
                def lambda52_(d_1120_p1_, d_1121_p2_):
                    def lambda53_(d_1122_i3_):
                        return (d_1122_i3_) * (len(_dafny.Map({d_1120_p1_: _dafny.SeqWithoutIsStrInference([d_1121_p2_, d_1121_p2_])})))

                    return lambda53_

                init29_ = lambda52_(p1, p2)
                nw192_ = _dafny.Array(None, 26)
                for i0_29_ in range(nw192_.length(0)):
                    nw192_[i0_29_] = init29_(i0_29_)
                d_1119_v39_ = nw192_
                index155_ = default__.safeIndex(999, (d_1119_v39_).length(0))
                (d_1119_v39_)[index155_] = self.f5
                index156_ = default__.safeIndex(999, (d_1119_v39_).length(0))
                (d_1119_v39_)[index156_] = default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference([p2, p2])).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference([p2, p2]))), p2)), len(_dafny.SeqWithoutIsStrInference([d_1064_v0_])))
                d_1123_v40_: int
                d_1124_v41_: bool
                out41_: int
                out42_: bool
                out41_, out42_ = (self).m2(p2, globalState)
                d_1123_v40_ = out41_
                d_1124_v41_ = out42_
                d_1125_v42_: _dafny.Seq
                d_1125_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxohwdduo"))
                d_1126_v43_: _dafny.Map
                d_1126_v43_ = _dafny.Map({(d_1096_v17_).f6: (0) - (len(d_1125_v42_))})
                d_1126_v43_ = (d_1126_v43_).set((d_1096_v17_).f6, 435)
                d_1124_v41_ = default__.fm21((d_1119_v39_)[default__.safeIndex(999, (d_1119_v39_).length(0))], (d_1096_v17_).f6, d_1124_v41_, (False if d_1124_v41_ else d_1124_v41_), globalState)
                d_1127_v44_: _dafny.Map
                d_1127_v44_ = _dafny.Map({p2: p1})
                d_1127_v44_ = (d_1127_v44_).set(((d_1119_v39_)[default__.safeIndex(999, (d_1119_v39_).length(0))]) != ((d_1119_v39_)[default__.safeIndex(999, (d_1119_v39_).length(0))]), (d_1123_v40_) + ((d_1119_v39_)[default__.safeIndex(999, (d_1119_v39_).length(0))]))
        if p2:
            d_1128_v45_: _dafny.Map
            d_1128_v45_ = _dafny.Map({(p3) * (p3): d_1064_v0_})
            d_1129_v47_: _dafny.MultiSet
            def iife62_():
                coll36_ = _dafny.Set()
                compr_36_: int
                for compr_36_ in (_dafny.MultiSet([-165, 396])).Elements:
                    d_1130_v46_: int = compr_36_
                    if (d_1130_v46_) in (_dafny.MultiSet([-165, 396])):
                        coll36_ = coll36_.union(_dafny.Set([(d_1130_v46_) * (-395)]))
                return _dafny.Set(coll36_)
            d_1129_v47_ = _dafny.MultiSet([(0) - (default__.fm24(iife62_()
            , globalState)), p1])
            d_1131_v48_: str
            d_1131_v48_ = _dafny.CodePoint('d')
            d_1132_v49_: _dafny.Seq
            d_1132_v49_ = _dafny.SeqWithoutIsStrInference([p3])
            d_1133_v50_: _dafny.Seq
            d_1133_v50_ = _dafny.SeqWithoutIsStrInference([d_1132_v49_])
            d_1134_v51_: _dafny.Seq
            d_1134_v51_ = _dafny.SeqWithoutIsStrInference([len(d_1133_v50_), p3])
            d_1135_v52_: D4
            d_1135_v52_ = D4_DC9((d_1129_v47_).cardinality, d_1131_v48_, d_1132_v49_, self.f5, p2)
            d_1128_v45_ = (d_1128_v45_).set((d_1135_v52_).cf19, d_1064_v0_)
            d_1136_v53_: bool
            d_1136_v53_ = False
            d_1137_v54_: D8
            d_1137_v54_ = D8_DC16(len(d_1134_v51_), d_1136_v53_, False, self.f5)
            rhs147_ = (d_1137_v54_).cf30
            rhs148_ = p2
            rhs149_ = p3
            rhs150_ = p1
            lhs88_ = self
            d_1136_v53_ = rhs147_
            d_1136_v53_ = rhs148_
            r1 = rhs149_
            lhs88_.f5 = rhs150_
            d_1138_v55_: _dafny.Seq
            d_1138_v55_ = _dafny.SeqWithoutIsStrInference([d_1136_v53_])
            d_1139_v56_: _dafny.Map
            d_1139_v56_ = _dafny.Map({p2: d_1138_v55_})
            d_1140_v57_: _dafny.Seq
            d_1140_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isekkdev"))
            d_1141_v58_: _dafny.MultiSet
            d_1141_v58_ = _dafny.MultiSet([not(d_1136_v53_)])
            d_1142_v59_: _dafny.Seq
            d_1142_v59_ = _dafny.SeqWithoutIsStrInference([d_1140_v57_, d_1140_v57_, ((d_1140_v57_).set(default__.safeIndex(self.f5, len(d_1140_v57_)), d_1131_v48_)).set(default__.safeIndex((0) - ((d_1141_v58_).cardinality), len((d_1140_v57_).set(default__.safeIndex(self.f5, len(d_1140_v57_)), d_1131_v48_))), d_1131_v48_)])
            d_1143_v60_: _dafny.Map
            d_1143_v60_ = _dafny.Map({(d_1129_v47_).cardinality: d_1136_v53_})
            d_1144_v61_: _dafny.Seq
            d_1144_v61_ = _dafny.SeqWithoutIsStrInference([d_1143_v60_, d_1143_v60_])
            d_1145_v63_: D6
            d_1145_v63_ = D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwxjc")))
            d_1146_v65_: _dafny.Array
            nw193_ = _dafny.Array(_dafny.Map({}), 13)
            d_1146_v65_ = nw193_
            d_1147_v66_: T0
            nw194_ = C4()
            nw194_.ctor__(d_1136_v53_, d_1129_v47_, d_1146_v65_)
            d_1147_v66_ = nw194_
            d_1148_v67_: D9
            d_1148_v67_ = D9_DC19(d_1147_v66_, 80, (0) - (self.f5))
            d_1149_v68_: _dafny.Map
            d_1149_v68_ = _dafny.Map({(d_1132_v49_)[default__.safeIndex(395, len(d_1132_v49_))]: d_1131_v48_})
            d_1150_v69_: _dafny.Array
            nw195_ = _dafny.Array(None, 27)
            nw195_[int(0)] = p1
            nw195_[int(1)] = len((_dafny.Map({p2: d_1138_v55_})) | (d_1139_v56_))
            nw195_[int(2)] = (p3) + (self.f5)
            nw195_[int(3)] = p3
            nw195_[int(4)] = len(d_1142_v59_)
            nw195_[int(5)] = p1
            def iife63_():
                coll37_ = _dafny.Set()
                compr_37_: int
                for compr_37_ in ((d_1144_v61_)[default__.safeIndex(p1, len(d_1144_v61_))]).keys.Elements:
                    d_1151_v62_: int = compr_37_
                    if (d_1151_v62_) in ((d_1144_v61_)[default__.safeIndex(p1, len(d_1144_v61_))]):
                        coll37_ = coll37_.union(_dafny.Set([(d_1151_v62_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([559, 505]))))]))
                return _dafny.Set(coll37_)
            nw195_[int(6)] = default__.fm24(iife63_()
            , globalState)
            nw195_[int(7)] = ((d_1129_v47_)[p1] if (p1) in (d_1129_v47_) else (D2_DC4(p3)).cf6)
            nw195_[int(8)] = (len(d_1140_v57_)) - ((D7_DC14(p1)).cf26)
            nw195_[int(9)] = default__.safeModuloInt(len((d_1145_v63_).cf22), (d_1134_v51_)[default__.safeIndex(p1, len(d_1134_v51_))])
            def iife64_():
                coll38_ = _dafny.Set()
                compr_38_: int
                for compr_38_ in _dafny.IntegerRange(190, 666):
                    d_1152_v64_: int = compr_38_
                    if ((190) <= (d_1152_v64_)) and ((d_1152_v64_) < (666)):
                        coll38_ = coll38_.union(_dafny.Set([(d_1152_v64_) - ((d_1141_v58_).cardinality)]))
                return _dafny.Set(coll38_)
            nw195_[int(10)] = default__.fm24(iife64_()
            , globalState)
            nw195_[int(11)] = default__.safeDivisionInt(len(d_1134_v51_), p1)
            nw195_[int(12)] = p1
            nw195_[int(13)] = p3
            nw195_[int(14)] = self.f5
            nw195_[int(15)] = self.f5
            nw195_[int(16)] = (p1) - (self.f5)
            nw195_[int(17)] = p3
            nw195_[int(18)] = self.f5
            nw195_[int(19)] = (d_1148_v67_).cf35
            nw195_[int(20)] = (p1) - (self.f5)
            nw195_[int(21)] = self.f5
            nw195_[int(22)] = self.f5
            nw195_[int(23)] = len(_dafny.SeqWithoutIsStrInference([_dafny.Map({225: _dafny.CodePoint('l')}), d_1149_v68_, d_1149_v68_, d_1149_v68_]))
            nw195_[int(24)] = p1
            nw195_[int(25)] = p3
            nw195_[int(26)] = p3
            d_1150_v69_ = nw195_
            index157_ = default__.safeIndex(963, (d_1150_v69_).length(0))
            (d_1150_v69_)[index157_] = p3
            d_1153_v70_: _dafny.Array
            nw196_ = _dafny.Array(_dafny.CodePoint('D'), 22)
            d_1153_v70_ = nw196_
            index158_ = default__.safeIndex(31, (d_1064_v0_).length(0))
            (d_1064_v0_)[index158_] = (p2 if p2 else d_1136_v53_)
            d_1154_v72_: _dafny.Map
            d_1154_v72_ = _dafny.Map({d_1140_v57_: d_1140_v57_})
            index159_ = default__.safeIndex(963, (d_1150_v69_).length(0))
            index160_ = default__.safeIndex(31, (d_1064_v0_).length(0))
            def iife65_():
                coll39_ = _dafny.Map()
                compr_39_: _dafny.Seq
                for compr_39_ in (d_1154_v72_).keys.Elements:
                    d_1155_v71_: _dafny.Seq = compr_39_
                    if (d_1155_v71_) in (d_1154_v72_):
                        coll39_[d_1155_v71_] = d_1140_v57_
                return _dafny.Map(coll39_)
            rhs151_ = default__.fm17(d_1142_v59_, len(iife65_()
            ), p1, self.f5, globalState)
            rhs152_ = d_1153_v70_
            rhs153_ = default__.fm21((((self).f4)[(d_1137_v54_).cf30] if ((d_1137_v54_).cf30) in ((self).f4) else self.f5), d_1131_v48_, p2, p2, globalState)
            lhs89_ = d_1150_v69_
            lhs90_ = default__.safeIndex(963, (d_1150_v69_).length(0))
            lhs91_ = d_1064_v0_
            lhs92_ = default__.safeIndex(31, (d_1064_v0_).length(0))
            lhs89_[lhs90_] = rhs151_
            d_1153_v70_ = rhs152_
            lhs91_[lhs92_] = rhs153_
            d_1156_v73_: _dafny.MultiSet
            d_1156_v73_ = _dafny.MultiSet([d_1138_v55_])
            d_1157_v74_: _dafny.Map
            d_1157_v74_ = _dafny.Map({d_1136_v53_: p2})
            d_1158_v75_: _dafny.Array
            d_1158_v75_ = d_1150_v69_
            d_1159_v76_: _dafny.Map
            d_1159_v76_ = _dafny.Map({default__.fm8(((d_1156_v73_)[d_1138_v55_] if (d_1138_v55_) in (d_1156_v73_) else len(d_1157_v74_)), globalState): d_1158_v75_})
            d_1159_v76_ = (d_1159_v76_).set(d_1140_v57_, d_1158_v75_)
            d_1160_v77_: _dafny.Map
            d_1160_v77_ = _dafny.Map({p3: self.f5})
            (globalState).f2 = default__.safeDivisionInt(p3, (len(d_1160_v77_) if p2 else len(d_1140_v57_)))
        elif True:
            d_1161_v78_: bool
            d_1161_v78_ = True
            d_1161_v78_ = (True) or (p2)
            d_1162_v79_: _dafny.Set
            d_1162_v79_ = _dafny.Set({self.f5})
            if (p3) not in ((default__.fm38(773, True, not(p2), d_1161_v78_, globalState)) | (d_1162_v79_)):
                d_1163_v80_: str
                d_1163_v80_ = _dafny.CodePoint('v')
                d_1163_v80_ = default__.fm26(len(_dafny.Map({d_1161_v78_: d_1161_v78_})), len(d_1162_v79_), p2, globalState)
                d_1164_v81_: C5
                nw197_ = C5()
                nw197_.ctor__(d_1163_v80_)
                d_1164_v81_ = nw197_
                d_1163_v80_ = (d_1164_v81_).f6
                r0 = p1
                d_1165_v82_: _dafny.Seq
                d_1165_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvmfvuww"))
                d_1166_v83_: _dafny.Seq
                d_1166_v83_ = _dafny.SeqWithoutIsStrInference([p2, d_1161_v78_, d_1161_v78_])
                d_1167_v84_: _dafny.Map
                d_1167_v84_ = _dafny.Map({d_1165_v82_: d_1166_v83_})
                d_1168_v85_: _dafny.Map
                d_1168_v85_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oaxbduxb"))})
                d_1167_v84_ = (d_1167_v84_).set(((((d_1168_v85_)[229] if (229) in (d_1168_v85_) else _dafny.SeqWithoutIsStrInference([(d_1164_v81_).f6 for d_1169_i4_ in range(default__.abs(-136))]))).set(default__.safeIndex(self.f5, len(((d_1168_v85_)[229] if (229) in (d_1168_v85_) else _dafny.SeqWithoutIsStrInference([(d_1164_v81_).f6 for d_1170_i4_ in range(default__.abs(-136))])))), d_1163_v80_)) + (d_1165_v82_), d_1166_v83_)
            elif True:
                d_1171_v86_: _dafny.MultiSet
                d_1171_v86_ = _dafny.MultiSet([p3])
                d_1172_v87_: C3
                nw198_ = C3()
                nw198_.ctor__((d_1171_v86_) - (d_1171_v86_), not(False))
                d_1172_v87_ = nw198_
                d_1173_v88_: _dafny.Seq
                d_1173_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tplrn"))
                d_1161_v78_ = not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjpsq"))) < ((d_1173_v88_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "diviu")))))
                d_1174_v89_: _dafny.Map
                d_1174_v89_ = _dafny.Map({d_1161_v78_: d_1161_v78_})
                d_1174_v89_ = (d_1174_v89_).set(True, p2)
                d_1175_v90_: _dafny.Array
                def lambda54_(d_1176_p3_):
                    def lambda55_(d_1177_i5_):
                        return (d_1177_i5_) * (d_1176_p3_)

                    return lambda55_

                init30_ = lambda54_(p3)
                nw199_ = _dafny.Array(None, 20)
                for i0_30_ in range(nw199_.length(0)):
                    nw199_[i0_30_] = init30_(i0_30_)
                d_1175_v90_ = nw199_
                index161_ = default__.safeIndex(949, (d_1175_v90_).length(0))
                (d_1175_v90_)[index161_] = p1
                d_1178_v91_: str
                d_1178_v91_ = _dafny.CodePoint('h')
                index162_ = default__.safeIndex(949, (d_1175_v90_).length(0))
                (d_1175_v90_)[index162_] = (d_1172_v87_).fm19(d_1178_v91_, 272, globalState)
                (globalState).f2 = default__.safeDivisionInt(default__.safeDivisionInt(74, (0) - (((d_1171_v86_)[self.f5] if (self.f5) in (d_1171_v86_) else self.f5))), p1)
            (self).f5 = default__.fm24(d_1162_v79_, globalState)
            d_1179_v92_: int
            d_1180_v93_: bool
            out43_: int
            out44_: bool
            out43_, out44_ = (self).m2(p2, globalState)
            d_1179_v92_ = out43_
            d_1180_v93_ = out44_
            d_1181_v94_: _dafny.Map
            d_1181_v94_ = _dafny.Map({-249: d_1180_v93_})
            d_1181_v94_ = (d_1181_v94_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpbxb"))), p2)
        (self).f5 = (p1) * (p3)
        d_1182_i6_: int
        d_1182_i6_ = 0
        with _dafny.label("7"):
            while p2:
                with _dafny.c_label("7"):
                    if (d_1182_i6_) >= (100):
                        raise _dafny.Break("7")
                    d_1182_i6_ = (d_1182_i6_) + (1)
                    d_1183_v95_: _dafny.MultiSet
                    d_1183_v95_ = _dafny.MultiSet([p3])
                    d_1184_v96_: _dafny.Seq
                    d_1184_v96_ = _dafny.SeqWithoutIsStrInference([d_1183_v95_, d_1183_v95_, d_1183_v95_])
                    (globalState).f2 = default__.safeModuloInt((0) - (default__.safeModuloInt(p1, len(d_1184_v96_))), self.f5)
                    d_1185_v97_: _dafny.Array
                    nw200_ = _dafny.Array(None, 7)
                    d_1185_v97_ = nw200_
                    d_1185_v97_ = d_1185_v97_
                    d_1186_v98_: bool
                    d_1186_v98_ = False
                    d_1187_v99_: _dafny.Seq
                    d_1187_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vm"))
                    d_1188_v100_: _dafny.Map
                    d_1188_v100_ = _dafny.Map({self.f5: (d_1187_v99_) < (d_1187_v99_)})
                    d_1189_v101_: _dafny.Set
                    d_1189_v101_ = _dafny.Set({d_1064_v0_})
                    rhs154_ = not(((d_1188_v100_)[(len(d_1189_v101_)) - (self.f5)] if ((len(d_1189_v101_)) - (self.f5)) in (d_1188_v100_) else (p2 if p2 else False)))
                    rhs155_ = not(d_1186_v98_)
                    d_1186_v98_ = rhs154_
                    d_1186_v98_ = rhs155_
                    (globalState).f2 = p3
                    pass
            pass
        d_1190_v102_: _dafny.Seq
        d_1190_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tp"))
        d_1191_v103_: _dafny.Map
        d_1191_v103_ = _dafny.Map({_dafny.CodePoint('q'): (d_1190_v102_) + (d_1190_v102_)})
        d_1192_v104_: str
        d_1192_v104_ = _dafny.CodePoint('b')
        d_1191_v103_ = (d_1191_v103_).set(d_1192_v104_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eltjfofh")))
        d_1193_v105_: _dafny.Set
        d_1193_v105_ = _dafny.Set({False})
        d_1194_v106_: _dafny.Map
        d_1194_v106_ = _dafny.Map({p2: False})
        d_1195_v107_: _dafny.Map
        d_1195_v107_ = _dafny.Map({p3: len(d_1194_v106_)})
        d_1196_v108_: _dafny.Map
        d_1196_v108_ = _dafny.Map({(_dafny.Set({p2, p2, p2})).ispropersubset(d_1193_v105_): len((d_1195_v107_).set(p1, self.f5))})
        d_1196_v108_ = (d_1196_v108_).set(p2, self.f5)
        r0 = 883
        r1 = (0) - (p1)
        d_1197_v109_: D2
        d_1197_v109_ = D2_DC5(p3, p2, False, p2, p2)
        def iife66_():
            coll40_ = _dafny.Set()
            compr_40_: _dafny.Seq
            for compr_40_ in (default__.fm42(False, p3, p2, globalState)).keys.Elements:
                d_1198_v110_: _dafny.Seq = compr_40_
                if (d_1198_v110_) in (default__.fm42(False, p3, p2, globalState)):
                    coll40_ = coll40_.union(_dafny.Set([d_1198_v110_]))
            return _dafny.Set(coll40_)
        r2 = (iife66_()
         if not((d_1197_v109_).cf10) else default__.fm35(globalState))
        return r0, r1, r2

    @property
    def f4(self):
        return self._f4

class C7(T0, T1):
    def  __init__(self):
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f3(self):
        return self._f3
    def ctor__(self, f3):
        (self)._f3 = f3

    def fm0(self, p0, globalState):
        return (default__.safeDivisionInt(len(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference([297]))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuwqkmw"))))) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1199_i0_ in range(default__.abs(679))])))

    def fm1(self, p0, p1, p2, p3, globalState):
        return 133

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: str = _dafny.CodePoint('D')
        d_1200_v0_: int
        d_1200_v0_ = -681
        r0 = (d_1200_v0_ if p1 else d_1200_v0_)
        if p1:
            d_1201_v1_: _dafny.Seq
            d_1201_v1_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1202_v2_: _dafny.Map
            d_1202_v2_ = _dafny.Map({(d_1201_v1_).set(default__.safeIndex(d_1200_v0_, len(d_1201_v1_)), p1): (p0).set(default__.safeIndex(d_1200_v0_, len(p0)), len(p2))})
            d_1203_v3_: str
            d_1203_v3_ = _dafny.CodePoint('t')
            d_1204_v4_: _dafny.Map
            d_1204_v4_ = _dafny.Map({p1: (d_1201_v1_)[default__.safeIndex(303, len(d_1201_v1_))]})
            d_1205_v5_: _dafny.Seq
            d_1205_v5_ = _dafny.SeqWithoutIsStrInference([d_1204_v4_, d_1204_v4_])
            d_1206_v6_: D1
            d_1206_v6_ = D1_DC1(d_1205_v5_)
            d_1207_v7_: _dafny.MultiSet
            d_1207_v7_ = _dafny.MultiSet([p1])
            d_1208_v8_: _dafny.Array
            nw201_ = _dafny.Array(None, 29)
            nw201_[int(0)] = (d_1201_v1_) in (d_1202_v2_)
            nw201_[int(1)] = p1
            nw201_[int(2)] = p1
            nw201_[int(3)] = (p1) or (False)
            nw201_[int(4)] = not (p1) or (p1)
            nw201_[int(5)] = p1
            nw201_[int(6)] = p1
            nw201_[int(7)] = (d_1200_v0_) < (d_1200_v0_)
            nw201_[int(8)] = p1
            nw201_[int(9)] = p1
            nw201_[int(10)] = not (p1) or (p1)
            nw201_[int(11)] = p1
            nw201_[int(12)] = (d_1200_v0_) > (d_1200_v0_)
            nw201_[int(13)] = not(default__.fm2(d_1203_v3_, (d_1206_v6_).cf1, p1, globalState))
            nw201_[int(14)] = p1
            nw201_[int(15)] = (-190) <= (len(_dafny.Map({d_1200_v0_: True})))
            nw201_[int(16)] = p1
            nw201_[int(17)] = default__.fm2(d_1203_v3_, default__.fm3(globalState), False, globalState)
            nw201_[int(18)] = (d_1200_v0_) <= (-359)
            nw201_[int(19)] = p1
            nw201_[int(20)] = p1
            nw201_[int(21)] = p1
            nw201_[int(22)] = p1
            nw201_[int(23)] = p1
            nw201_[int(24)] = p1
            nw201_[int(25)] = not((True) in (d_1207_v7_))
            nw201_[int(26)] = p1
            nw201_[int(27)] = p1
            nw201_[int(28)] = p1
            d_1208_v8_ = nw201_
            index163_ = default__.safeIndex(667, (d_1208_v8_).length(0))
            (d_1208_v8_)[index163_] = p1
            index164_ = default__.safeIndex(667, (d_1208_v8_).length(0))
            (d_1208_v8_)[index164_] = p1
            d_1209_v9_: _dafny.Array
            def lambda56_(d_1210_v0_, d_1211_v1_, d_1212_v8_, d_1213_p1_):
                def lambda57_(d_1214_i0_):
                    return (_dafny.Map({D1_DC2(-156, False, d_1210_v0_): d_1211_v1_})) | (_dafny.Map({D1_DC2(len(_dafny.Map({(d_1212_v8_)[default__.safeIndex(667, (d_1212_v8_).length(0))]: d_1210_v0_})), d_1213_p1_, d_1210_v0_): d_1211_v1_}))

                return lambda57_

            init31_ = lambda56_(d_1200_v0_, d_1201_v1_, d_1208_v8_, p1)
            nw202_ = _dafny.Array(None, 15)
            for i0_31_ in range(nw202_.length(0)):
                nw202_[i0_31_] = init31_(i0_31_)
            d_1209_v9_ = nw202_
            d_1215_v10_: _dafny.MultiSet
            d_1215_v10_ = _dafny.MultiSet([d_1200_v0_])
            d_1216_v11_: D1
            d_1216_v11_ = D1_DC2((d_1215_v10_).cardinality, (d_1208_v8_)[default__.safeIndex(667, (d_1208_v8_).length(0))], d_1200_v0_)
            d_1217_v12_: _dafny.Map
            d_1217_v12_ = _dafny.Map({d_1216_v11_: d_1201_v1_})
            index165_ = default__.safeIndex(119, (d_1209_v9_).length(0))
            (d_1209_v9_)[index165_] = (d_1217_v12_).set(d_1216_v11_, d_1201_v1_)
            d_1218_v13_: _dafny.Set
            d_1218_v13_ = _dafny.Set({p1})
            index166_ = default__.safeIndex(119, (d_1209_v9_).length(0))
            (d_1209_v9_)[index166_] = ((d_1217_v12_).set(D1_DC2(d_1200_v0_, (d_1208_v8_)[default__.safeIndex(667, (d_1208_v8_).length(0))], len(d_1218_v13_)), d_1201_v1_)) | (_dafny.Map({d_1216_v11_: d_1201_v1_}))
            index167_ = default__.safeIndex(667, (d_1208_v8_).length(0))
            (d_1208_v8_)[index167_] = (d_1216_v11_).cf3
            d_1215_v10_ = d_1215_v10_
            if default__.fm2(d_1203_v3_, d_1205_v5_, p1, globalState):
                d_1206_v6_ = default__.fm4(p1, d_1203_v3_, d_1200_v0_, default__.fm5(d_1200_v0_, globalState), globalState)
                d_1219_v14_: _dafny.Array
                def lambda58_(d_1220_p2_):
                    def lambda59_(d_1221_i1_):
                        return d_1220_p2_

                    return lambda59_

                init32_ = lambda58_(p2)
                nw203_ = _dafny.Array(None, 29)
                for i0_32_ in range(nw203_.length(0)):
                    nw203_[i0_32_] = init32_(i0_32_)
                d_1219_v14_ = nw203_
                index168_ = default__.safeIndex(515, (d_1219_v14_).length(0))
                (d_1219_v14_)[index168_] = p2
                index169_ = default__.safeIndex(515, (d_1219_v14_).length(0))
                (d_1219_v14_)[index169_] = p2
                index170_ = default__.safeIndex(667, (d_1208_v8_).length(0))
                (d_1208_v8_)[index170_] = (d_1208_v8_)[default__.safeIndex(667, (d_1208_v8_).length(0))]
                d_1222_v15_: _dafny.Map
                d_1222_v15_ = _dafny.Map({p2: _dafny.CodePoint('n')})
                d_1222_v15_ = (d_1222_v15_).set(p2, default__.fm5(len(d_1201_v1_), globalState))
                r0 = d_1200_v0_
            elif True:
                index171_ = default__.safeIndex(667, (d_1208_v8_).length(0))
                (d_1208_v8_)[index171_] = (d_1208_v8_)[default__.safeIndex(667, (d_1208_v8_).length(0))]
                d_1223_v16_: _dafny.Array
                def lambda60_(d_1224_v11_):
                    def lambda61_(d_1225_i2_):
                        return d_1224_v11_

                    return lambda61_

                init33_ = lambda60_(d_1216_v11_)
                nw204_ = _dafny.Array(None, 10)
                for i0_33_ in range(nw204_.length(0)):
                    nw204_[i0_33_] = init33_(i0_33_)
                d_1223_v16_ = nw204_
                d_1226_v17_: _dafny.Seq
                d_1226_v17_ = _dafny.SeqWithoutIsStrInference([d_1208_v8_, d_1208_v8_, d_1208_v8_, d_1208_v8_])
                d_1227_v18_: _dafny.Map
                d_1227_v18_ = _dafny.Map({d_1223_v16_: (d_1226_v17_) + (d_1226_v17_)})
                d_1227_v18_ = (d_1227_v18_).set(d_1223_v16_, _dafny.SeqWithoutIsStrInference([d_1208_v8_, d_1208_v8_]))
                r0 = 597
                d_1228_v19_: _dafny.Seq
                d_1228_v19_ = _dafny.SeqWithoutIsStrInference([d_1201_v1_])
                d_1229_v20_: _dafny.Seq
                d_1229_v20_ = d_1201_v1_
                d_1230_v21_: _dafny.Map
                d_1230_v21_ = _dafny.Map({(d_1216_v11_).cf3: ((d_1228_v19_)[default__.safeIndex(d_1200_v0_, len(d_1228_v19_))]) + (default__.fm6(len(p2), (self).fm1((d_1201_v1_)[default__.safeIndex(-124, len(d_1201_v1_))], d_1203_v3_, d_1229_v20_, (d_1208_v8_)[default__.safeIndex(667, (d_1208_v8_).length(0))], globalState), d_1203_v3_, globalState))})
                d_1230_v21_ = (d_1230_v21_).set((d_1208_v8_)[default__.safeIndex(667, (d_1208_v8_).length(0))], d_1201_v1_)
                d_1231_v22_: _dafny.Array
                def lambda62_(d_1232_p2_):
                    def lambda63_(d_1233_i3_):
                        return (d_1233_i3_) + (len(d_1232_p2_))

                    return lambda63_

                init34_ = lambda62_(p2)
                nw205_ = _dafny.Array(None, 20)
                for i0_34_ in range(nw205_.length(0)):
                    nw205_[i0_34_] = init34_(i0_34_)
                d_1231_v22_ = nw205_
                d_1234_v23_: _dafny.Seq
                d_1234_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgdm"))
                index172_ = default__.safeIndex(667, (d_1208_v8_).length(0))
                rhs156_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khi"))) + (d_1234_v23_))
                rhs157_ = d_1231_v22_
                rhs158_ = not ((default__.fm2(d_1203_v3_, d_1205_v5_, not(not(True)), globalState)) and ((d_1208_v8_)[default__.safeIndex(667, (d_1208_v8_).length(0))])) or ((d_1200_v0_) == (d_1200_v0_))
                rhs159_ = p2
                lhs93_ = globalState
                lhs94_ = d_1208_v8_
                lhs95_ = default__.safeIndex(667, (d_1208_v8_).length(0))
                lhs93_.f2 = rhs156_
                d_1231_v22_ = rhs157_
                lhs94_[lhs95_] = rhs158_
                d_1234_v23_ = rhs159_
        elif True:
            (globalState).f2 = d_1200_v0_
            d_1235_v24_: bool
            d_1235_v24_ = False
            d_1235_v24_ = d_1235_v24_
            d_1236_v25_: _dafny.Array
            nw206_ = _dafny.Array(False, 11)
            d_1236_v25_ = nw206_
            d_1237_v26_: str
            d_1237_v26_ = _dafny.CodePoint('i')
            d_1238_v27_: _dafny.Map
            d_1238_v27_ = _dafny.Map({d_1235_v24_: p1})
            d_1239_v28_: _dafny.Seq
            d_1239_v28_ = _dafny.SeqWithoutIsStrInference([d_1238_v27_])
            index173_ = default__.safeIndex(845, (d_1236_v25_).length(0))
            (d_1236_v25_)[index173_] = default__.fm2(d_1237_v26_, d_1239_v28_, d_1235_v24_, globalState)
            index174_ = default__.safeIndex(845, (d_1236_v25_).length(0))
            (d_1236_v25_)[index174_] = d_1235_v24_
            d_1238_v27_ = (d_1238_v27_) | (d_1238_v27_)
            d_1240_v29_: _dafny.MultiSet
            d_1240_v29_ = _dafny.MultiSet([d_1200_v0_, d_1200_v0_, d_1200_v0_])
            d_1241_v30_: _dafny.Map
            d_1241_v30_ = _dafny.Map({d_1200_v0_: d_1200_v0_})
            d_1242_v31_: _dafny.MultiSet
            d_1242_v31_ = _dafny.MultiSet([p1])
            d_1243_v32_: _dafny.Array
            nw207_ = _dafny.Array(None, 18)
            nw207_[int(0)] = (d_1240_v29_).cardinality
            nw207_[int(1)] = d_1200_v0_
            nw207_[int(2)] = d_1200_v0_
            nw207_[int(3)] = d_1200_v0_
            nw207_[int(4)] = d_1200_v0_
            nw207_[int(5)] = d_1200_v0_
            nw207_[int(6)] = d_1200_v0_
            nw207_[int(7)] = 297
            nw207_[int(8)] = d_1200_v0_
            nw207_[int(9)] = (D1_DC2(591, (d_1236_v25_)[default__.safeIndex(845, (d_1236_v25_).length(0))], d_1200_v0_)).cf4
            nw207_[int(10)] = d_1200_v0_
            nw207_[int(11)] = len(d_1241_v30_)
            nw207_[int(12)] = d_1200_v0_
            nw207_[int(13)] = d_1200_v0_
            nw207_[int(14)] = d_1200_v0_
            nw207_[int(15)] = len(p2)
            nw207_[int(16)] = d_1200_v0_
            nw207_[int(17)] = ((d_1242_v31_)[p1] if (p1) in (d_1242_v31_) else d_1200_v0_)
            d_1243_v32_ = nw207_
            d_1244_v33_: _dafny.Map
            d_1244_v33_ = _dafny.Map({(len(_dafny.Set({p1, (d_1236_v25_)[default__.safeIndex(845, (d_1236_v25_).length(0))]}))) >= (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1245_i4_ in range(default__.abs(-718))]))): d_1243_v32_})
            d_1244_v33_ = (d_1244_v33_) | (d_1244_v33_)
        d_1246_v34_: _dafny.Array
        nw208_ = _dafny.Array(None, 2)
        nw208_[int(0)] = p1
        nw208_[int(1)] = p1
        d_1246_v34_ = nw208_
        index175_ = default__.safeIndex(649, (d_1246_v34_).length(0))
        (d_1246_v34_)[index175_] = not (False) or (p1)
        d_1247_v35_: _dafny.Map
        d_1247_v35_ = _dafny.Map({p1: True})
        d_1248_v36_: _dafny.Seq
        d_1248_v36_ = _dafny.SeqWithoutIsStrInference([d_1247_v35_])
        index176_ = default__.safeIndex(649, (d_1246_v34_).length(0))
        (d_1246_v34_)[index176_] = not (not(p1)) or (default__.fm2(default__.fm5(d_1200_v0_, globalState), d_1248_v36_, p1, globalState))
        d_1249_v37_: C0
        nw209_ = C0()
        nw209_.ctor__(d_1200_v0_)
        d_1249_v37_ = nw209_
        d_1250_v38_: _dafny.MultiSet
        d_1250_v38_ = _dafny.MultiSet([d_1200_v0_, len(p2), (d_1249_v37_).f13])
        hi5_ = d_1200_v0_
        for d_1251_i5_ in range(default__.safeDivisionInt((d_1250_v38_).cardinality, d_1200_v0_), hi5_):
            d_1247_v35_ = (d_1247_v35_).set(p1, True)
            if p1:
                d_1252_v39_: _dafny.Seq
                d_1252_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')])
                d_1252_v39_ = p2
                d_1253_v40_: _dafny.Seq
                d_1253_v40_ = _dafny.SeqWithoutIsStrInference([d_1250_v38_, d_1250_v38_, d_1250_v38_])
                d_1254_v41_: _dafny.Set
                d_1254_v41_ = _dafny.Set({(d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))]})
                d_1255_v42_: _dafny.Seq
                d_1255_v42_ = _dafny.SeqWithoutIsStrInference([(d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))]])
                d_1256_v43_: _dafny.Array
                nw210_ = _dafny.Array(None, 16)
                nw210_[int(0)] = _dafny.MultiSet([(d_1249_v37_).f13])
                nw210_[int(1)] = d_1250_v38_
                nw210_[int(2)] = d_1250_v38_
                nw210_[int(3)] = (d_1253_v40_)[default__.safeIndex(384, len(d_1253_v40_))]
                nw210_[int(4)] = d_1250_v38_
                nw210_[int(5)] = d_1250_v38_
                nw210_[int(6)] = _dafny.MultiSet([(d_1249_v37_).f13])
                nw210_[int(7)] = _dafny.MultiSet([(d_1249_v37_).f13, 907])
                nw210_[int(8)] = d_1250_v38_
                nw210_[int(9)] = _dafny.MultiSet([d_1251_i5_, (d_1249_v37_).f13])
                nw210_[int(10)] = d_1250_v38_
                nw210_[int(11)] = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffaybi"))), len(d_1254_v41_), d_1251_i5_, (0) - ((d_1249_v37_).f13)])
                nw210_[int(12)] = d_1250_v38_
                nw210_[int(13)] = _dafny.MultiSet([len(d_1255_v42_)])
                nw210_[int(14)] = d_1250_v38_
                nw210_[int(15)] = d_1250_v38_
                d_1256_v43_ = nw210_
                d_1257_v44_: _dafny.Map
                d_1257_v44_ = _dafny.Map({d_1256_v43_: d_1200_v0_})
                d_1257_v44_ = (d_1257_v44_).set(d_1256_v43_, (d_1249_v37_).f13)
                index177_ = default__.safeIndex(649, (d_1246_v34_).length(0))
                (d_1246_v34_)[index177_] = not(p1)
                index178_ = default__.safeIndex(649, (d_1246_v34_).length(0))
                (d_1246_v34_)[index178_] = False
                (globalState).f2 = 804
            elif True:
                d_1258_v45_: _dafny.MultiSet
                d_1258_v45_ = _dafny.MultiSet([p1, (d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))]])
                d_1259_v46_: _dafny.Seq
                d_1259_v46_ = _dafny.SeqWithoutIsStrInference([(d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))], p1, p1])
                d_1260_v47_: _dafny.MultiSet
                d_1260_v47_ = _dafny.MultiSet([(d_1258_v45_) - (_dafny.MultiSet(d_1259_v46_))])
                d_1260_v47_ = ((d_1260_v47_ if p1 else d_1260_v47_)) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p1]), _dafny.MultiSet([p1, default__.fm23(p2, (d_1249_v37_).f13, (d_1249_v37_).f13, False, globalState)]), d_1258_v45_, d_1258_v45_])))
                d_1261_v48_: D1
                d_1261_v48_ = D1_DC2(len(p2), p1, (d_1249_v37_).f13)
                index179_ = default__.safeIndex(649, (d_1246_v34_).length(0))
                (d_1246_v34_)[index179_] = ((d_1261_v48_).cf4) >= ((d_1249_v37_).f13)
                d_1262_v49_: _dafny.Array
                nw211_ = _dafny.Array(_dafny.Map({}), 24)
                d_1262_v49_ = nw211_
                index180_ = default__.safeIndex(303, (d_1262_v49_).length(0))
                (d_1262_v49_)[index180_] = default__.fm43(_dafny.MultiSet(p0), d_1251_i5_, d_1251_i5_, globalState)
                d_1263_v50_: _dafny.Map
                d_1263_v50_ = _dafny.Map({p2: (0) - ((d_1249_v37_).f13)})
                d_1264_v51_: _dafny.Map
                d_1264_v51_ = _dafny.Map({d_1258_v45_: _dafny.MultiSet([(d_1249_v37_).f13])})
                d_1265_v52_: _dafny.Seq
                d_1265_v52_ = _dafny.SeqWithoutIsStrInference([d_1263_v50_])
                d_1266_v53_: _dafny.Set
                d_1266_v53_ = _dafny.Set({(d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))], p1})
                d_1267_v54_: _dafny.Seq
                d_1267_v54_ = _dafny.SeqWithoutIsStrInference([d_1263_v50_, (d_1265_v52_)[default__.safeIndex(len(d_1266_v53_), len(d_1265_v52_))]])
                d_1268_v56_: _dafny.Map
                d_1268_v56_ = _dafny.Map({p2: (d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))]})
                index181_ = default__.safeIndex(303, (d_1262_v49_).length(0))
                index182_ = default__.safeIndex(649, (d_1246_v34_).length(0))
                def iife67_():
                    coll41_ = _dafny.Map()
                    compr_41_: _dafny.Seq
                    for compr_41_ in (d_1268_v56_).keys.Elements:
                        d_1269_v55_: _dafny.Seq = compr_41_
                        if (d_1269_v55_) in (d_1268_v56_):
                            coll41_[d_1269_v55_] = (d_1249_v37_).f13
                    return _dafny.Map(coll41_)
                rhs160_ = (d_1264_v51_) | (d_1264_v51_)
                rhs161_ = ((d_1263_v50_) | ((d_1267_v54_)[default__.safeIndex(d_1251_i5_, len(d_1267_v54_))])) | (iife67_()
                )
                rhs162_ = ((d_1249_v37_).f13) > (d_1251_i5_)
                lhs96_ = d_1262_v49_
                lhs97_ = default__.safeIndex(303, (d_1262_v49_).length(0))
                lhs98_ = d_1246_v34_
                lhs99_ = default__.safeIndex(649, (d_1246_v34_).length(0))
                lhs96_[lhs97_] = rhs160_
                d_1263_v50_ = rhs161_
                lhs98_[lhs99_] = rhs162_
                d_1200_v0_ = (0) - ((0) - (default__.safeModuloInt((d_1249_v37_).f13, default__.safeDivisionInt(d_1251_i5_, 803))))
                d_1270_v57_: _dafny.Map
                d_1270_v57_ = _dafny.Map({d_1251_i5_: p1})
                d_1271_v58_: _dafny.Map
                d_1271_v58_ = _dafny.Map({d_1270_v57_: -541})
                index183_ = default__.safeIndex(649, (d_1246_v34_).length(0))
                (d_1246_v34_)[index183_] = (_dafny.Map({(0) - (d_1251_i5_): p1})) not in (d_1271_v58_)
            d_1272_v59_: D1
            d_1272_v59_ = D1_DC2((d_1249_v37_).f13, (d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))], d_1251_i5_)
            index184_ = default__.safeIndex(649, (d_1246_v34_).length(0))
            (d_1246_v34_)[index184_] = (d_1272_v59_).cf3
            d_1273_v60_: D2
            d_1273_v60_ = D2_DC4((d_1250_v38_).cardinality)
            d_1273_v60_ = d_1273_v60_
        if p1:
            d_1274_v61_: _dafny.Seq
            d_1274_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyoeph"))
            d_1274_v61_ = d_1274_v61_
            d_1275_v62_: str
            d_1275_v62_ = _dafny.CodePoint('c')
            r1 = d_1275_v62_
            d_1276_v63_: _dafny.Array
            nw212_ = _dafny.Array(None, 8)
            nw212_[int(0)] = (d_1249_v37_).f13
            nw212_[int(1)] = (d_1249_v37_).f13
            nw212_[int(2)] = d_1200_v0_
            nw212_[int(3)] = (d_1249_v37_).f13
            nw212_[int(4)] = 159
            nw212_[int(5)] = (p0)[default__.safeIndex(d_1200_v0_, len(p0))]
            nw212_[int(6)] = (d_1249_v37_).f13
            nw212_[int(7)] = (_dafny.MultiSet(p2)).cardinality
            d_1276_v63_ = nw212_
            d_1277_v64_: _dafny.Array
            d_1277_v64_ = d_1276_v63_
            d_1278_v65_: _dafny.Map
            d_1278_v65_ = _dafny.Map({(d_1249_v37_).f13: d_1277_v64_})
            d_1279_v66_: _dafny.Set
            d_1279_v66_ = _dafny.Set({p1})
            d_1280_v67_: _dafny.Map
            d_1280_v67_ = _dafny.Map({d_1200_v0_: (d_1249_v37_).f13})
            d_1278_v65_ = (d_1278_v65_).set(len((default__.fm12(not(p1), d_1200_v0_, d_1279_v66_, d_1280_v67_, globalState)) + (p0)), d_1277_v64_)
            d_1281_v68_: C2
            nw213_ = C2()
            nw213_.ctor__((False) and ((d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))]), default__.safeModuloInt(d_1200_v0_, (d_1249_v37_).f13), (self).f3)
            d_1281_v68_ = nw213_
            d_1282_v69_: _dafny.Map
            d_1282_v69_ = _dafny.Map({p1: d_1281_v68_})
            d_1281_v68_ = ((d_1282_v69_)[(d_1281_v68_).f11] if ((d_1281_v68_).f11) in (d_1282_v69_) else d_1281_v68_)
            d_1283_v70_: D4
            d_1283_v70_ = D4_DC8(d_1249_v37_)
            pat_let_tv25_ = d_1249_v37_
            def iife68_(_pat_let13_0):
                def iife69_(d_1284_dt__update__tmp_h0_):
                    def iife70_(_pat_let14_0):
                        def iife71_(d_1285_dt__update_hcf15_h0_):
                            return D4_DC8(d_1285_dt__update_hcf15_h0_)
                        return iife71_(_pat_let14_0)
                    return iife70_(pat_let_tv25_)
                return iife69_(_pat_let13_0)
            d_1283_v70_ = iife68_(d_1283_v70_)
        elif True:
            d_1286_v71_: T1
            nw214_ = C2()
            nw214_.ctor__(p1, (d_1249_v37_).f13, (self).f3)
            d_1286_v71_ = nw214_
            d_1287_v72_: _dafny.Array
            nw215_ = _dafny.Array(None, 2)
            nw215_[int(0)] = d_1286_v71_
            nw215_[int(1)] = d_1286_v71_
            d_1287_v72_ = nw215_
            index185_ = default__.safeIndex(721, (d_1287_v72_).length(0))
            (d_1287_v72_)[index185_] = d_1286_v71_
            index186_ = default__.safeIndex(721, (d_1287_v72_).length(0))
            (d_1287_v72_)[index186_] = d_1286_v71_
            d_1288_v73_: _dafny.MultiSet
            d_1288_v73_ = _dafny.MultiSet([(d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))], (d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))]])
            d_1289_v74_: _dafny.Seq
            d_1289_v74_ = _dafny.SeqWithoutIsStrInference([((d_1288_v73_)[(d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))]] if ((d_1246_v34_)[default__.safeIndex(649, (d_1246_v34_).length(0))]) in (d_1288_v73_) else (d_1249_v37_).f13), (0) - ((d_1249_v37_).f13), d_1200_v0_, (d_1249_v37_).f13, (d_1249_v37_).f13])
            d_1289_v74_ = d_1289_v74_
            d_1290_v75_: _dafny.Array
            nw216_ = _dafny.Array(_dafny.Map({}), 8)
            d_1290_v75_ = nw216_
            nw217_ = _dafny.Array(_dafny.Map({}), 20)
            d_1290_v75_ = nw217_
            d_1288_v73_ = d_1288_v73_
            d_1291_v76_: _dafny.Map
            d_1291_v76_ = _dafny.Map({(d_1249_v37_).f13: (d_1249_v37_).f13})
            d_1292_v77_: D1
            d_1292_v77_ = D1_DC2(((d_1291_v76_)[(d_1249_v37_).f13] if ((d_1249_v37_).f13) in (d_1291_v76_) else d_1200_v0_), (p2) <= (p2), d_1200_v0_)
            d_1292_v77_ = d_1292_v77_
        r0 = d_1200_v0_
        d_1293_v78_: str
        d_1293_v78_ = _dafny.CodePoint('n')
        r1 = d_1293_v78_
        return r0, r1

