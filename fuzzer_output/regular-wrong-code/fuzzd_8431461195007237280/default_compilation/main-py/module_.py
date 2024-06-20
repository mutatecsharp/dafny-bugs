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
        return (0) - ((0) - ((default__.safeModuloInt(4, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False})))) + (592)))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return not(not(((_dafny.Set({_dafny.CodePoint('t'), _dafny.CodePoint('s'), _dafny.CodePoint('f')}) if not(True) else _dafny.Set({_dafny.CodePoint('r'), _dafny.CodePoint('k')}))).ispropersubset(_dafny.Set({_dafny.CodePoint('k')}))))

    @staticmethod
    def fm2(globalState):
        if not((_dafny.CodePoint('x')) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anivmhual")))):
            return _dafny.Set({False})
        elif True:
            return _dafny.Set({True, not(True), True})

    @staticmethod
    def fm5(p0, globalState):
        if (_dafny.CodePoint('w')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujvkkcrk"))):
            return _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_0_i0_ in range(default__.abs(16))])), len(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpe"))))]))}))})
        elif True:
            return _dafny.Set({233})

    @staticmethod
    def fm7(p0, p1, globalState):
        def iife0_():
            def iife2_():
                coll2_ = _dafny.Set()
                compr_2_: str
                for compr_2_ in (_dafny.Map({_dafny.CodePoint('p'): len(_dafny.Map({(_dafny.MultiSet([len(_dafny.Set({24}))])).cardinality: True}))})).keys.Elements:
                    d_3_v1_: str = compr_2_
                    if (d_3_v1_) in (_dafny.Map({_dafny.CodePoint('p'): len(_dafny.Map({(_dafny.MultiSet([len(_dafny.Set({24}))])).cardinality: True}))})):
                        coll2_ = coll2_.union(_dafny.Set([d_3_v1_]))
                return _dafny.Set(coll2_)
            coll0_ = _dafny.Map()
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: str
                for compr_1_ in (_dafny.Map({_dafny.CodePoint('p'): len(_dafny.Map({(_dafny.MultiSet([len(_dafny.Set({24}))])).cardinality: True}))})).keys.Elements:
                    d_1_v1_: str = compr_1_
                    if (d_1_v1_) in (_dafny.Map({_dafny.CodePoint('p'): len(_dafny.Map({(_dafny.MultiSet([len(_dafny.Set({24}))])).cardinality: True}))})):
                        coll1_ = coll1_.union(_dafny.Set([d_1_v1_]))
                return _dafny.Set(coll1_)
            compr_0_: str
            for compr_0_ in (iife1_()
            ).Elements:
                d_2_v0_: str = compr_0_
                if (d_2_v0_) in (iife2_()
                ):
                    coll0_[d_2_v0_] = len(_dafny.Map({False: _dafny.CodePoint('v')}))
            return _dafny.Map(coll0_)
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: str
                for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('o')])).Elements:
                    d_5_v3_: str = compr_5_
                    if (d_5_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('o')])):
                        coll5_[d_5_v3_] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])), _dafny.MultiSet([False]), _dafny.MultiSet([True])])), len(_dafny.SeqWithoutIsStrInference([629, 834]))]))
                return _dafny.Map(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: str
                for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('o')])).Elements:
                    d_5_v3_: str = compr_4_
                    if (d_5_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('o')])):
                        coll4_[d_5_v3_] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])), _dafny.MultiSet([False]), _dafny.MultiSet([True])])), len(_dafny.SeqWithoutIsStrInference([629, 834]))]))
                return _dafny.Map(coll4_)
            compr_3_: str
            for compr_3_ in (iife4_()
            ).keys.Elements:
                d_6_v2_: str = compr_3_
                if (d_6_v2_) in (iife5_()
                ):
                    def iife6_():
                        coll6_ = _dafny.Map()
                        compr_6_: int
                        for compr_6_ in _dafny.IntegerRange(309, 772):
                            d_7_v4_: int = compr_6_
                            if ((309) <= (d_7_v4_)) and ((d_7_v4_) < (772)):
                                coll6_[default__.safeDivisionInt(d_7_v4_, -942)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wg"))))
                        return _dafny.Map(coll6_)
                    coll3_[d_6_v2_] = len(iife6_()
                    )
            return _dafny.Map(coll3_)
        return ((iife0_()
        ) | (_dafny.Map({_dafny.CodePoint('s'): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_4_i0_ in range(default__.abs(508))]))}))) | (iife3_()
        )

    @staticmethod
    def fm8(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kh"))

    @staticmethod
    def fm9(globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])])) | ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_8_i0_ in range(default__.abs(722))]))))

    @staticmethod
    def fm10(globalState):
        return _dafny.CodePoint('a')

    @staticmethod
    def fm11(globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Seq
            for compr_7_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))})).Elements:
                d_10_v0_: _dafny.Seq = compr_7_
                if (d_10_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))})):
                    coll7_[d_10_v0_] = True
            return _dafny.Map(coll7_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_9_i0_ in range(default__.abs(35))]): False})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsov")): False}))) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")): not(True)})) | (iife7_()
        ))

    @staticmethod
    def fm14(p0, globalState):
        return _dafny.Map({_dafny.CodePoint('t'): (len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pess")))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pj"))), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))).cardinality, len(_dafny.Map({501: True}))}))) * (143)})

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        if (D1_DC6(not(False), False)).cf8:
            return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddchqe")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_11_i0_ in range(default__.abs(788))])])))
        elif True:
            return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_12_i1_ in range(default__.abs(597))])]))

    @staticmethod
    def fm16(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True, False, False, True, True])): 698})])

    @staticmethod
    def fm17(globalState):
        return _dafny.Set({(472) - (-750)})

    @staticmethod
    def fm18(p0, globalState):
        return (_dafny.Map({(D12_DC31(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wyhk")), 716, True)).cf47: (D5_DC14(not(False), 560)).cf17})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgxmchpc")): 411}))

    @staticmethod
    def fm19(p0, p1, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(-378, -241):
                d_13_v0_: int = compr_8_
                if ((-378) <= (d_13_v0_)) and ((d_13_v0_) < (-241)):
                    coll8_[(d_13_v0_) * (-530)] = (len(_dafny.SeqWithoutIsStrInference([690 for d_14_i0_ in range(default__.abs(249))]))) - (len(_dafny.SeqWithoutIsStrInference([False, True])))
            return _dafny.Map(coll8_)
        return iife8_()
        

    @staticmethod
    def fm20(p0, globalState):
        source0_ = D3_DC10(not(False))
        if source0_.is_DC9:
            d_15___mcc_h0_ = source0_.cf11
            d_16___mcc_h1_ = source0_.cf12
            d_17_cf12_ = d_16___mcc_h1_
            d_18_cf11_ = d_15___mcc_h0_
            return _dafny.CodePoint('u')
        elif source0_.is_DC10:
            d_19___mcc_h2_ = source0_.cf13
            d_20_cf13_ = d_19___mcc_h2_
            return _dafny.CodePoint('e')
        elif True:
            d_21___mcc_h3_ = source0_.cf10
            d_22_cf10_ = d_21___mcc_h3_
            return d_22_cf10_

    @staticmethod
    def fm21(globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihvn"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvbkgws")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjxpor")))

    @staticmethod
    def fm22(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 282}), _dafny.Map({not(True): 266}), _dafny.Map({True: len(_dafny.Map({False: -107}))}), _dafny.Map({True: len(_dafny.Map({False: False}))})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 571})]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({True: (0) - (-604)})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 453})])))

    @staticmethod
    def fm23(p0, p1, globalState):
        return ((_dafny.Map({False: not(True)})) | (_dafny.Map({False: False}))) | ((_dafny.Map({True: True})) | (_dafny.Map({not(not(True)): True})))

    @staticmethod
    def fm24(globalState):
        return D3_DC10(not(True))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return D0_DC1()

    @staticmethod
    def fm26(globalState):
        return ((_dafny.Map({not(True): -304})) | (_dafny.Map({False: 597}))) | (_dafny.Map({False: -128}))

    @staticmethod
    def fm27(globalState):
        return ((_dafny.SeqWithoutIsStrInference([False, True]) if True else _dafny.SeqWithoutIsStrInference([False, True, True]))) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(-876, 398):
                d_23_v0_: int = compr_9_
                if ((-876) <= (d_23_v0_)) and ((d_23_v0_) < (398)):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_23_v0_, 378)]))
            return _dafny.Set(coll9_)
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-633, len(iife9_()
        ), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_24_i0_ in range(default__.abs(80))])))]))) - (_dafny.MultiSet([52]))) | ((_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: True}) for d_25_i1_ in range(default__.abs(466))])): _dafny.Set({168})})), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({_dafny.Set({len(_dafny.Set({_dafny.CodePoint('a')}))}): False})): 569}))])), len(_dafny.Map({True: True}))]) if False else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({976}))]))))

    @staticmethod
    def fm29(globalState):
        return D1_DC4(295, default__.safeModuloInt((_dafny.MultiSet([not(True)])).cardinality, 420), (151) <= (-521))

    @staticmethod
    def fm30(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([-734])

    @staticmethod
    def fm31(p0, p1, globalState):
        source1_ = D12_DC29(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([595])), 463}))
        if source1_.is_DC30:
            d_26___mcc_h0_ = source1_.cf44
            d_27___mcc_h1_ = source1_.cf45
            d_28_cf45_ = d_27___mcc_h1_
            d_29_cf44_ = d_26___mcc_h0_
            return D1_DC5(993)
        elif source1_.is_DC31:
            d_30___mcc_h2_ = source1_.cf46
            d_31___mcc_h3_ = source1_.cf47
            d_32___mcc_h4_ = source1_.cf48
            d_33___mcc_h5_ = source1_.cf49
            d_34_cf49_ = d_33___mcc_h5_
            d_35_cf48_ = d_32___mcc_h4_
            d_36_cf47_ = d_31___mcc_h3_
            d_37_cf46_ = d_30___mcc_h2_
            return D1_DC5((0) - (d_35_cf48_))
        elif source1_.is_DC29:
            d_38___mcc_h6_ = source1_.cf43
            d_39_cf43_ = d_38___mcc_h6_
            return D1_DC5(215)
        elif True:
            d_40___mcc_h7_ = source1_.cf50
            d_41_cf50_ = d_40___mcc_h7_
            if False:
                return D1_DC5(-264)
            elif True:
                return D1_DC5(len(_dafny.Map({False: len(_dafny.Set({784}))})))

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        r2: int = int(0)
        if not(p1):
            d_42_v0_: str
            d_42_v0_ = _dafny.CodePoint('d')
            d_43_v1_: _dafny.Seq
            d_43_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srfvaskfl"))
            d_44_v2_: D0
            d_44_v2_ = D0_DC0(p2)
            d_45_v3_: _dafny.Map
            d_45_v3_ = _dafny.Map({d_43_v1_: (d_44_v2_).cf0})
            d_46_v4_: C3
            nw0_ = C3()
            nw0_.ctor__(d_42_v0_, d_45_v3_)
            d_46_v4_ = nw0_
            d_47_v5_: _dafny.Set
            d_47_v5_ = _dafny.Set({d_46_v4_, d_46_v4_})
            d_48_v6_: _dafny.Seq
            d_48_v6_ = _dafny.SeqWithoutIsStrInference([d_47_v5_])
            d_49_v7_: D1
            d_49_v7_ = D1_DC6(p1, p1)
            d_50_v8_: _dafny.Seq
            d_50_v8_ = _dafny.SeqWithoutIsStrInference([p1])
            d_51_v9_: _dafny.MultiSet
            d_51_v9_ = _dafny.MultiSet([p2, p0, (0) - (p2), (0) - (p0), p0])
            d_52_v10_: _dafny.Array
            nw1_ = _dafny.Array(None, 17)
            nw1_[int(0)] = p1
            nw1_[int(1)] = p1
            nw1_[int(2)] = p1
            nw1_[int(3)] = ((d_48_v6_)[default__.safeIndex(p0, len(d_48_v6_))]).ispropersubset(d_47_v5_)
            nw1_[int(4)] = (d_49_v7_).cf8
            nw1_[int(5)] = p1
            nw1_[int(6)] = not(((0) - (p0)) < (len(default__.fm21(globalState))))
            nw1_[int(7)] = p1
            nw1_[int(8)] = not(False)
            nw1_[int(9)] = (d_50_v8_)[default__.safeIndex((d_51_v9_).cardinality, len(d_50_v8_))]
            nw1_[int(10)] = p1
            nw1_[int(11)] = p1
            nw1_[int(12)] = p1
            nw1_[int(13)] = p1
            nw1_[int(14)] = (p0) <= (p0)
            nw1_[int(15)] = p1
            nw1_[int(16)] = p1
            d_52_v10_ = nw1_
            index0_ = default__.safeIndex(517, (d_52_v10_).length(0))
            (d_52_v10_)[index0_] = not(p1)
            index1_ = default__.safeIndex(517, (d_52_v10_).length(0))
            (d_52_v10_)[index1_] = p1
            d_53_v11_: _dafny.Map
            d_53_v11_ = _dafny.Map({(d_52_v10_)[default__.safeIndex(517, (d_52_v10_).length(0))]: d_42_v0_})
            d_53_v11_ = (d_53_v11_).set(p1, d_42_v0_)
            index2_ = default__.safeIndex(517, (d_52_v10_).length(0))
            (d_52_v10_)[index2_] = True
            (globalState).f6 = len(d_43_v1_)
            index3_ = default__.safeIndex(517, (d_52_v10_).length(0))
            (d_52_v10_)[index3_] = not ((d_52_v10_)[default__.safeIndex(517, (d_52_v10_).length(0))]) or (p1)
        elif True:
            d_54_v12_: str
            d_54_v12_ = _dafny.CodePoint('j')
            d_55_v13_: _dafny.Seq
            d_55_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqatmsp"))
            d_56_v14_: _dafny.Array
            nw2_ = _dafny.Array(None, 6)
            nw2_[int(0)] = p0
            nw2_[int(1)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_57_i0_ in range(default__.abs(986))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_58_i0_ in range(default__.abs(986))]))), d_54_v12_))
            nw2_[int(2)] = default__.safeDivisionInt(p0, (0) - (len(d_55_v13_)))
            nw2_[int(3)] = p0
            nw2_[int(4)] = -689
            nw2_[int(5)] = p2
            d_56_v14_ = nw2_
            index4_ = default__.safeIndex(69, (d_56_v14_).length(0))
            (d_56_v14_)[index4_] = p0
            d_59_v15_: _dafny.MultiSet
            d_59_v15_ = _dafny.MultiSet([p0])
            d_60_v16_: _dafny.Array
            nw3_ = _dafny.Array(None, 4)
            nw3_[int(0)] = d_59_v15_
            nw3_[int(1)] = (d_59_v15_) - (d_59_v15_)
            nw3_[int(2)] = d_59_v15_
            nw3_[int(3)] = (d_59_v15_).intersection(d_59_v15_)
            d_60_v16_ = nw3_
            index5_ = default__.safeIndex(93, (d_60_v16_).length(0))
            (d_60_v16_)[index5_] = default__.fm28(p1, p2, p1, globalState)
            d_61_v17_: _dafny.MultiSet
            d_61_v17_ = _dafny.MultiSet([p1, p1, p1, not(not(p1))])
            d_62_v18_: _dafny.Array
            nw4_ = _dafny.Array(None, 17)
            nw4_[int(0)] = p1
            nw4_[int(1)] = p1
            nw4_[int(2)] = p1
            nw4_[int(3)] = True
            nw4_[int(4)] = p1
            nw4_[int(5)] = p1
            nw4_[int(6)] = p1
            nw4_[int(7)] = p1
            nw4_[int(8)] = p1
            nw4_[int(9)] = (p1 if p1 else p1)
            nw4_[int(10)] = p1
            nw4_[int(11)] = True
            nw4_[int(12)] = default__.fm1(p1, d_54_v12_, 245, globalState)
            nw4_[int(13)] = not(not(p1))
            nw4_[int(14)] = (d_61_v17_).isdisjoint(d_61_v17_)
            nw4_[int(15)] = default__.fm1(True, d_54_v12_, p2, globalState)
            nw4_[int(16)] = p1
            d_62_v18_ = nw4_
            d_63_v19_: _dafny.Map
            d_63_v19_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjijcedg")): p2})
            d_64_v20_: C2
            nw5_ = C2()
            nw5_.ctor__(d_54_v12_, d_63_v19_)
            d_64_v20_ = nw5_
            d_65_v21_: _dafny.Set
            d_65_v21_ = _dafny.Set({d_64_v20_})
            index6_ = default__.safeIndex(69, (d_56_v14_).length(0))
            index7_ = default__.safeIndex(93, (d_60_v16_).length(0))
            rhs0_ = default__.safeDivisionInt(p2, p2)
            rhs1_ = (d_59_v15_) - (d_59_v15_)
            rhs2_ = d_62_v18_
            rhs3_ = 664
            rhs4_ = not((d_64_v20_) not in (d_65_v21_))
            lhs0_ = d_56_v14_
            lhs1_ = default__.safeIndex(69, (d_56_v14_).length(0))
            lhs2_ = d_60_v16_
            lhs3_ = default__.safeIndex(93, (d_60_v16_).length(0))
            lhs4_ = globalState
            lhs0_[lhs1_] = rhs0_
            lhs2_[lhs3_] = rhs1_
            r0 = rhs2_
            lhs4_.f6 = rhs3_
            r1 = rhs4_
            d_66_v22_: C3
            nw6_ = C3()
            nw6_.ctor__(d_54_v12_, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umuxo")): (d_56_v14_)[default__.safeIndex(69, (d_56_v14_).length(0))]}))
            d_66_v22_ = nw6_
            d_67_v23_: _dafny.Array
            def lambda0_(d_68_v13_):
                def lambda1_(d_69_i1_):
                    return d_68_v13_

                return lambda1_

            init0_ = lambda0_(d_55_v13_)
            nw7_ = _dafny.Array(None, 7)
            for i0_0_ in range(nw7_.length(0)):
                nw7_[i0_0_] = init0_(i0_0_)
            d_67_v23_ = nw7_
            index8_ = default__.safeIndex(671, (d_67_v23_).length(0))
            (d_67_v23_)[index8_] = d_55_v13_
            index9_ = default__.safeIndex(671, (d_67_v23_).length(0))
            rhs5_ = p1
            rhs6_ = ((_dafny.SeqWithoutIsStrInference([d_54_v12_ for d_70_i2_ in range(default__.abs(38))])) + (d_55_v13_)) + (((d_55_v13_).set(default__.safeIndex((d_56_v14_)[default__.safeIndex(69, (d_56_v14_).length(0))], len(d_55_v13_)), d_54_v12_)).set(default__.safeIndex(p2, len((d_55_v13_).set(default__.safeIndex((d_56_v14_)[default__.safeIndex(69, (d_56_v14_).length(0))], len(d_55_v13_)), d_54_v12_))), _dafny.CodePoint('q')))
            rhs7_ = d_66_v22_
            rhs8_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "estdt"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfn")))
            lhs5_ = d_67_v23_
            lhs6_ = default__.safeIndex(671, (d_67_v23_).length(0))
            r1 = rhs5_
            d_55_v13_ = rhs6_
            d_66_v22_ = rhs7_
            lhs5_[lhs6_] = rhs8_
            d_71_v24_: D0
            d_71_v24_ = D0_DC2(D0_DC0(p2))
            d_72_v25_: D11
            d_72_v25_ = D11_DC27(d_63_v19_)
            d_73_v26_: C1
            nw8_ = C1()
            nw8_.ctor__(687, p1, default__.fm20(default__.fm1(p1, d_54_v12_, (d_59_v15_).cardinality, globalState), globalState), ((d_72_v25_).cf39) | ((d_63_v19_).set((d_67_v23_)[default__.safeIndex(671, (d_67_v23_).length(0))], (d_56_v14_)[default__.safeIndex(69, (d_56_v14_).length(0))])))
            d_73_v26_ = nw8_
            d_74_v27_: _dafny.Map
            d_74_v27_ = _dafny.Map({d_54_v12_: p0})
            d_75_v28_: T0
            nw9_ = C0()
            nw9_.ctor__((d_67_v23_)[default__.safeIndex(671, (d_67_v23_).length(0))], d_74_v27_, d_54_v12_, _dafny.Map({d_55_v13_: len(_dafny.Map({default__.fm0(globalState): (d_56_v14_)[default__.safeIndex(69, (d_56_v14_).length(0))]}))}))
            d_75_v28_ = nw9_
            d_76_v29_: _dafny.Map
            d_76_v29_ = _dafny.Map({d_75_v28_: p2})
            rhs9_ = d_62_v18_
            rhs10_ = d_71_v24_
            rhs11_ = d_73_v26_
            rhs12_ = ((d_76_v29_)[d_75_v28_] if (d_75_v28_) in (d_76_v29_) else ((d_56_v14_)[default__.safeIndex(69, (d_56_v14_).length(0))]) - (p2))
            rhs13_ = d_62_v18_
            d_62_v18_ = rhs9_
            d_71_v24_ = rhs10_
            d_73_v26_ = rhs11_
            r2 = rhs12_
            d_62_v18_ = rhs13_
            (d_75_v28_).f18 = d_54_v12_
            index10_ = default__.safeIndex(297, (d_62_v18_).length(0))
            (d_62_v18_)[index10_] = (d_73_v26_).f23
            index11_ = default__.safeIndex(297, (d_62_v18_).length(0))
            (d_62_v18_)[index11_] = (316) in (_dafny.SeqWithoutIsStrInference([d_73_v26_.f22]))
        d_77_v30_: _dafny.Seq
        d_77_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgcj"))
        hi0_ = len((d_77_v30_) + (default__.fm21(globalState)))
        for d_78_i3_ in range(p0, hi0_):
            d_79_v31_: _dafny.Map
            d_79_v31_ = _dafny.Map({False: d_78_i3_})
            (globalState).f6 = ((d_79_v31_)[False] if (False) in (d_79_v31_) else d_78_i3_)
            d_80_v32_: _dafny.Map
            d_80_v32_ = _dafny.Map({not(p1): (d_77_v30_) + (d_77_v30_)})
            d_81_v33_: str
            d_81_v33_ = _dafny.CodePoint('y')
            d_80_v32_ = (d_80_v32_).set(default__.fm1(p1, d_81_v33_, len((_dafny.SeqWithoutIsStrInference([d_81_v33_ for d_82_i4_ in range(default__.abs(538))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_81_v33_ for d_83_i4_ in range(default__.abs(538))]))), default__.fm20(default__.fm1(p1, d_81_v33_, p2, globalState), globalState))), globalState), d_77_v30_)
            d_84_v34_: _dafny.Map
            d_84_v34_ = _dafny.Map({d_81_v33_: p2})
            d_85_v35_: _dafny.Map
            d_85_v35_ = _dafny.Map({d_77_v30_: p0})
            d_86_v36_: C0
            nw10_ = C0()
            nw10_.ctor__((d_77_v30_).set(default__.safeIndex(d_78_i3_, len(d_77_v30_)), d_81_v33_), (d_84_v34_) | (d_84_v34_), _dafny.CodePoint('r'), (d_85_v35_) | (d_85_v35_))
            d_86_v36_ = nw10_
            d_87_v37_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.MultiSet({}), 3)
            d_87_v37_ = nw11_
            d_88_v38_: _dafny.Map
            d_88_v38_ = _dafny.Map({p0: d_87_v37_})
            d_89_v39_: _dafny.Array
            nw12_ = _dafny.Array(None, 24)
            nw12_[int(0)] = d_87_v37_
            nw12_[int(1)] = d_87_v37_
            nw12_[int(2)] = d_87_v37_
            nw12_[int(3)] = d_87_v37_
            nw12_[int(4)] = d_87_v37_
            nw12_[int(5)] = d_87_v37_
            nw12_[int(6)] = d_87_v37_
            nw12_[int(7)] = d_87_v37_
            nw12_[int(8)] = d_87_v37_
            nw12_[int(9)] = d_87_v37_
            nw12_[int(10)] = ((d_88_v38_)[p0] if (p0) in (d_88_v38_) else d_87_v37_)
            nw12_[int(11)] = d_87_v37_
            nw12_[int(12)] = d_87_v37_
            nw12_[int(13)] = d_87_v37_
            nw12_[int(14)] = d_87_v37_
            nw12_[int(15)] = d_87_v37_
            nw12_[int(16)] = d_87_v37_
            nw12_[int(17)] = d_87_v37_
            nw12_[int(18)] = d_87_v37_
            nw12_[int(19)] = d_87_v37_
            nw12_[int(20)] = d_87_v37_
            nw12_[int(21)] = d_87_v37_
            nw12_[int(22)] = d_87_v37_
            nw12_[int(23)] = d_87_v37_
            d_89_v39_ = nw12_
            index12_ = default__.safeIndex(181, (d_89_v39_).length(0))
            (d_89_v39_)[index12_] = d_87_v37_
            index13_ = default__.safeIndex(181, (d_89_v39_).length(0))
            (d_89_v39_)[index13_] = d_87_v37_
        d_90_v40_: _dafny.Array
        nw13_ = _dafny.Array(int(0), 2)
        d_90_v40_ = nw13_
        index14_ = default__.safeIndex(338, (d_90_v40_).length(0))
        (d_90_v40_)[index14_] = default__.safeDivisionInt(p2, p0)
        d_91_v41_: _dafny.Array
        nw14_ = _dafny.Array(False, 6)
        d_91_v41_ = nw14_
        index15_ = default__.safeIndex(360, (d_91_v41_).length(0))
        (d_91_v41_)[index15_] = ((0) - ((0) - (p2))) >= ((0) - (len(_dafny.SeqWithoutIsStrInference([p2]))))
        d_92_v42_: str
        d_92_v42_ = _dafny.CodePoint('j')
        index16_ = default__.safeIndex(338, (d_90_v40_).length(0))
        index17_ = default__.safeIndex(360, (d_91_v41_).length(0))
        rhs14_ = 715
        rhs15_ = not(p1)
        rhs16_ = (_dafny.SeqWithoutIsStrInference([d_92_v42_, default__.fm10(globalState)])) + (default__.fm8(globalState))
        lhs7_ = d_90_v40_
        lhs8_ = default__.safeIndex(338, (d_90_v40_).length(0))
        lhs9_ = d_91_v41_
        lhs10_ = default__.safeIndex(360, (d_91_v41_).length(0))
        lhs7_[lhs8_] = rhs14_
        lhs9_[lhs10_] = rhs15_
        d_77_v30_ = rhs16_
        d_93_v43_: _dafny.Map
        d_93_v43_ = _dafny.Map({d_92_v42_: (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]})
        d_94_v45_: _dafny.Map
        d_94_v45_ = _dafny.Map({d_92_v42_: 979})
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: str
            for compr_10_ in (d_94_v45_).keys.Elements:
                d_95_v44_: str = compr_10_
                if (d_95_v44_) in (d_94_v45_):
                    coll10_[d_95_v44_] = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
            return _dafny.Map(coll10_)
        d_93_v43_ = (iife10_()
        ) | ((d_93_v43_) | (_dafny.Map({d_92_v42_: (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]})))
        d_96_v46_: _dafny.Map
        d_96_v46_ = _dafny.Map({not(p1): p1})
        if ((0) - (default__.safeDivisionInt((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], len(d_96_v46_)))) < (p2):
            d_97_v47_: _dafny.Map
            d_97_v47_ = _dafny.Map({p0: True})
            d_98_v48_: _dafny.Seq
            d_98_v48_ = _dafny.SeqWithoutIsStrInference([d_97_v47_])
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(-992, 954):
                    d_99_v49_: int = compr_11_
                    if ((-992) <= (d_99_v49_)) and ((d_99_v49_) < (954)):
                        coll11_[default__.safeDivisionInt(d_99_v49_, p2)] = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                return _dafny.Map(coll11_)
            if (d_98_v48_) != (_dafny.SeqWithoutIsStrInference([iife11_()
 for d_100_i5_ in range(default__.abs(472))])):
                d_101_v50_: D9
                d_101_v50_ = D9_DC23(_dafny.SeqWithoutIsStrInference([_dafny.Set({p0}) for d_102_i6_ in range(default__.abs(865))]), p2, p1, p2)
                index18_ = default__.safeIndex(360, (d_91_v41_).length(0))
                index19_ = default__.safeIndex(338, (d_90_v40_).length(0))
                index20_ = default__.safeIndex(360, (d_91_v41_).length(0))
                rhs17_ = (_dafny.SeqWithoutIsStrInference([(d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]])).set(default__.safeIndex((0) - (p0), len(_dafny.SeqWithoutIsStrInference([(d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]]))), (d_101_v50_).cf31)
                rhs18_ = ((d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))] if p1 else not(p1))
                rhs19_ = default__.safeModuloInt((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], p0)
                rhs20_ = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                rhs21_ = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                lhs11_ = globalState
                lhs12_ = d_91_v41_
                lhs13_ = default__.safeIndex(360, (d_91_v41_).length(0))
                lhs14_ = d_90_v40_
                lhs15_ = default__.safeIndex(338, (d_90_v40_).length(0))
                lhs16_ = d_91_v41_
                lhs17_ = default__.safeIndex(360, (d_91_v41_).length(0))
                lhs11_.f1 = rhs17_
                lhs12_[lhs13_] = rhs18_
                lhs14_[lhs15_] = rhs19_
                lhs16_[lhs17_] = rhs20_
                r1 = rhs21_
                d_103_v51_: _dafny.Map
                d_103_v51_ = _dafny.Map({d_77_v30_: p0})
                d_104_v52_: _dafny.Set
                d_104_v52_ = _dafny.Set({True, True, not(p1)})
                d_105_v53_: D0
                d_105_v53_ = D0_DC0(p0)
                d_103_v51_ = (d_103_v51_).set((d_77_v30_).set(default__.safeIndex((default__.fm31(d_104_v52_, p0, globalState)).cf6, len(d_77_v30_)), d_92_v42_), (d_105_v53_).cf0)
                index21_ = default__.safeIndex(360, (d_91_v41_).length(0))
                (d_91_v41_)[index21_] = p1
                d_106_v54_: _dafny.Map
                d_106_v54_ = _dafny.Map({(d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]: p2})
                d_106_v54_ = (d_106_v54_).set((d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))], (0) - (default__.safeDivisionInt((0) - (p2), 119)))
                index22_ = default__.safeIndex(338, (d_90_v40_).length(0))
                (d_90_v40_)[index22_] = -852
            elif True:
                d_107_v55_: D10
                d_107_v55_ = D10_DC25(d_90_v40_)
                d_108_v56_: _dafny.Seq
                d_108_v56_ = _dafny.SeqWithoutIsStrInference([(d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]])
                pat_let_tv0_ = d_90_v40_
                d_109_v57_: _dafny.Map
                def iife12_(_pat_let0_0):
                    def iife13_(d_110_dt__update__tmp_h0_):
                        def iife14_(_pat_let1_0):
                            def iife15_(d_111_dt__update_hcf33_h0_):
                                return D10_DC25(d_111_dt__update_hcf33_h0_)
                            return iife15_(_pat_let1_0)
                        return iife14_(pat_let_tv0_)
                    return iife13_(_pat_let0_0)
                d_109_v57_ = _dafny.Map({(iife12_(d_107_v55_)).cf33: d_108_v56_})
                d_109_v57_ = d_109_v57_
                d_112_v58_: _dafny.Seq
                d_112_v58_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                index23_ = default__.safeIndex(338, (d_90_v40_).length(0))
                index24_ = default__.safeIndex(360, (d_91_v41_).length(0))
                rhs22_ = len(_dafny.SeqWithoutIsStrInference([d_92_v42_ for d_113_i7_ in range(default__.abs(79))]))
                rhs23_ = default__.safeModuloInt(default__.safeDivisionInt((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], (0) - (p2)), p2)
                rhs24_ = d_112_v58_
                rhs25_ = False
                lhs18_ = globalState
                lhs19_ = d_90_v40_
                lhs20_ = default__.safeIndex(338, (d_90_v40_).length(0))
                lhs21_ = globalState
                lhs22_ = d_91_v41_
                lhs23_ = default__.safeIndex(360, (d_91_v41_).length(0))
                lhs18_.f6 = rhs22_
                lhs19_[lhs20_] = rhs23_
                lhs21_.f1 = rhs24_
                lhs22_[lhs23_] = rhs25_
                r1 = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                d_92_v42_ = d_92_v42_
                index25_ = default__.safeIndex(360, (d_91_v41_).length(0))
                (d_91_v41_)[index25_] = not(p1)
            if True:
                d_114_v59_: _dafny.MultiSet
                d_114_v59_ = _dafny.MultiSet([p1])
                index26_ = default__.safeIndex(360, (d_91_v41_).length(0))
                (d_91_v41_)[index26_] = (p1 if (d_114_v59_).ispropersubset(d_114_v59_) else (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))])
                d_115_v60_: _dafny.Seq
                d_115_v60_ = _dafny.SeqWithoutIsStrInference([(d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]])
                d_116_v61_: _dafny.Map
                d_116_v61_ = _dafny.Map({d_77_v30_: len(d_115_v60_)})
                d_117_v62_: C1
                nw15_ = C1()
                nw15_.ctor__(p2, (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))], d_92_v42_, (d_116_v61_).set(d_77_v30_, p0))
                d_117_v62_ = nw15_
                d_118_v63_: _dafny.MultiSet
                d_118_v63_ = _dafny.MultiSet([d_117_v62_, d_117_v62_, d_117_v62_])
                d_119_v64_: _dafny.Map
                d_119_v64_ = _dafny.Map({(d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]: d_118_v63_})
                d_120_v65_: _dafny.MultiSet
                d_120_v65_ = _dafny.MultiSet([len(d_119_v64_)])
                index27_ = default__.safeIndex(360, (d_91_v41_).length(0))
                (d_91_v41_)[index27_] = not(((0) - ((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))])) in (d_120_v65_))
                d_121_v66_: _dafny.Map
                d_121_v66_ = _dafny.Map({p0: d_90_v40_})
                d_122_v67_: _dafny.Set
                d_122_v67_ = _dafny.Set({p0})
                d_123_v68_: _dafny.MultiSet
                d_123_v68_ = _dafny.MultiSet([d_90_v40_, d_90_v40_, (((d_121_v66_)[len(d_122_v67_)] if (len(d_122_v67_)) in (d_121_v66_) else d_90_v40_) if (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))] else d_90_v40_)])
                d_123_v68_ = d_123_v68_
                index28_ = default__.safeIndex(338, (d_90_v40_).length(0))
                (d_90_v40_)[index28_] = default__.safeModuloInt((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], (p0) * ((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]))
                d_124_v69_: C0
                nw16_ = C0()
                nw16_.ctor__(d_77_v30_, default__.fm14((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], globalState), d_92_v42_, d_116_v61_)
                d_124_v69_ = nw16_
            elif True:
                d_125_v70_: _dafny.Set
                d_125_v70_ = _dafny.Set({p2})
                d_126_v71_: _dafny.Seq
                d_126_v71_ = _dafny.SeqWithoutIsStrInference([(d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]])
                d_127_v72_: _dafny.MultiSet
                d_127_v72_ = _dafny.MultiSet([d_77_v30_])
                d_128_v73_: _dafny.Seq
                d_128_v73_ = _dafny.SeqWithoutIsStrInference([(d_126_v71_)[default__.safeIndex(len(_dafny.Map({len((d_77_v30_).set(default__.safeIndex((d_127_v72_).cardinality, len(d_77_v30_)), d_92_v42_)): ((d_97_v47_)[p2] if (p2) in (d_97_v47_) else p1)})), len(d_126_v71_))]])
                index29_ = default__.safeIndex(360, (d_91_v41_).length(0))
                rhs26_ = ((d_125_v70_).intersection(d_125_v70_)).isdisjoint((d_125_v70_) | (d_125_v70_))
                rhs27_ = (d_128_v73_)[default__.safeIndex((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], len(d_128_v73_))]
                lhs24_ = d_91_v41_
                lhs25_ = default__.safeIndex(360, (d_91_v41_).length(0))
                r1 = rhs26_
                lhs24_[lhs25_] = rhs27_
                index30_ = default__.safeIndex(338, (d_90_v40_).length(0))
                (d_90_v40_)[index30_] = default__.safeModuloInt(len(d_77_v30_), p0)
                index31_ = default__.safeIndex(338, (d_90_v40_).length(0))
                (d_90_v40_)[index31_] = p2
                d_92_v42_ = d_92_v42_
                d_129_v74_: D1
                d_129_v74_ = D1_DC3(p1)
                d_130_v75_: D0
                d_130_v75_ = D0_DC1()
                d_131_v76_: _dafny.Array
                def lambda2_(d_132_v71_):
                    def lambda3_(d_133_i8_):
                        return d_132_v71_

                    return lambda3_

                init1_ = lambda2_(d_126_v71_)
                nw17_ = _dafny.Array(None, 25)
                for i0_1_ in range(nw17_.length(0)):
                    nw17_[i0_1_] = init1_(i0_1_)
                d_131_v76_ = nw17_
                index32_ = default__.safeIndex(765, (d_131_v76_).length(0))
                (d_131_v76_)[index32_] = d_126_v71_
                d_134_v77_: _dafny.Array
                nw18_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_134_v77_ = nw18_
                d_135_v78_: _dafny.Array
                nw19_ = _dafny.Array(D1.default()(), 17)
                d_135_v78_ = nw19_
                index33_ = default__.safeIndex(952, (d_134_v77_).length(0))
                (d_134_v77_)[index33_] = d_135_v78_
                index34_ = default__.safeIndex(765, (d_131_v76_).length(0))
                index35_ = default__.safeIndex(952, (d_134_v77_).length(0))
                rhs28_ = d_129_v74_
                rhs29_ = -121
                rhs30_ = (d_130_v75_ if p1 else d_130_v75_)
                rhs31_ = d_126_v71_
                rhs32_ = d_135_v78_
                lhs26_ = globalState
                lhs27_ = d_131_v76_
                lhs28_ = default__.safeIndex(765, (d_131_v76_).length(0))
                lhs29_ = d_134_v77_
                lhs30_ = default__.safeIndex(952, (d_134_v77_).length(0))
                d_129_v74_ = rhs28_
                lhs26_.f6 = rhs29_
                d_130_v75_ = rhs30_
                lhs27_[lhs28_] = rhs31_
                lhs29_[lhs30_] = rhs32_
            d_136_v79_: D6
            d_136_v79_ = D6_DC15(d_94_v45_)
            source2_ = d_136_v79_
            if source2_.is_DC16:
                d_137___mcc_h0_ = source2_.cf19
                d_138___mcc_h1_ = source2_.cf20
                d_139_cf20_ = d_138___mcc_h1_
                d_140_cf19_ = d_137___mcc_h0_
                d_141_v80_: _dafny.Map
                d_141_v80_ = _dafny.Map({613: d_77_v30_})
                d_141_v80_ = ((d_141_v80_) | (d_141_v80_)) | ((d_141_v80_) | (d_141_v80_))
                d_142_v82_: D9
                d_142_v82_ = D9_DC22(d_77_v30_)
                d_143_v83_: _dafny.Seq
                d_143_v83_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qribx"))).set(default__.safeIndex(73, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qribx")))), d_92_v42_), (d_142_v82_).cf27])
                d_144_v84_: C2
                nw20_ = C2()
                def iife16_():
                    coll12_ = _dafny.Map()
                    compr_12_: _dafny.Seq
                    for compr_12_ in (d_143_v83_).Elements:
                        d_145_v81_: _dafny.Seq = compr_12_
                        if (d_145_v81_) in (d_143_v83_):
                            coll12_[d_145_v81_] = len(_dafny.Map({len(_dafny.Set({p2, p0})): p1}))
                    return _dafny.Map(coll12_)
                nw20_.ctor__(d_92_v42_, iife16_()
                )
                d_144_v84_ = nw20_
                d_146_v85_: _dafny.Set
                d_146_v85_ = _dafny.Set({d_97_v47_})
                d_147_v87_: _dafny.MultiSet
                def iife17_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in _dafny.IntegerRange(763, -498):
                        d_148_v86_: int = compr_13_
                        if ((763) <= (d_148_v86_)) and ((d_148_v86_) < (-498)):
                            coll13_[(d_148_v86_) * (len(d_97_v47_))] = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                    return _dafny.Map(coll13_)
                d_147_v87_ = _dafny.MultiSet([iife17_()
                ])
                def iife18_():
                    coll14_ = _dafny.Set()
                    compr_14_: _dafny.Map
                    for compr_14_ in (d_147_v87_).Elements:
                        d_149_v88_: _dafny.Map = compr_14_
                        if (d_149_v88_) in (d_147_v87_):
                            coll14_ = coll14_.union(_dafny.Set([d_149_v88_]))
                    return _dafny.Set(coll14_)
                d_146_v85_ = ((d_146_v85_) - (iife18_()
                )) - (d_146_v85_)
                d_150_v89_: _dafny.Map
                d_150_v89_ = _dafny.Map({d_77_v30_: p2})
                d_151_v90_: C1
                nw21_ = C1()
                nw21_.ctor__(p2, (default__.fm30(p1, globalState)) <= (_dafny.SeqWithoutIsStrInference([p2, (d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]])), _dafny.CodePoint('u'), d_150_v89_)
                d_151_v90_ = nw21_
                d_152_v91_: C3
                nw22_ = C3()
                nw22_.ctor__(d_92_v42_, (d_150_v89_) | (d_150_v89_))
                d_152_v91_ = nw22_
                d_153_v92_: _dafny.Array
                nw23_ = _dafny.Array(None, 16)
                d_153_v92_ = nw23_
                index36_ = default__.safeIndex(439, (d_153_v92_).length(0))
                (d_153_v92_)[index36_] = d_152_v91_
                d_154_v93_: _dafny.MultiSet
                d_154_v93_ = _dafny.MultiSet([(d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))], True])
                pat_let_tv1_ = d_90_v40_
                pat_let_tv2_ = d_90_v40_
                pat_let_tv3_ = p1
                index37_ = default__.safeIndex(338, (d_90_v40_).length(0))
                index38_ = default__.safeIndex(439, (d_153_v92_).length(0))
                def iife19_(_pat_let2_0):
                    def iife20_(d_155_dt__update__tmp_h1_):
                        def iife21_(_pat_let3_0):
                            def iife22_(d_156_dt__update_hcf38_h0_):
                                def iife23_(_pat_let4_0):
                                    def iife24_(d_157_dt__update_hcf37_h0_):
                                        def iife25_(_pat_let5_0):
                                            def iife26_(d_158_dt__update_hcf34_h0_):
                                                return D10_DC26(d_158_dt__update_hcf34_h0_, (d_155_dt__update__tmp_h1_).cf35, (d_155_dt__update__tmp_h1_).cf36, d_157_dt__update_hcf37_h0_, d_156_dt__update_hcf38_h0_)
                                            return iife26_(_pat_let5_0)
                                        return iife25_(-203)
                                    return iife24_(_pat_let4_0)
                                return iife23_(not(pat_let_tv3_))
                            return iife22_(_pat_let3_0)
                        return iife21_((pat_let_tv2_)[default__.safeIndex(338, (pat_let_tv1_).length(0))])
                    return iife20_(_pat_let2_0)
                rhs33_ = d_151_v90_
                rhs34_ = d_152_v91_
                rhs35_ = ((len(d_77_v30_)) * ((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))])) + ((p2) - ((d_154_v93_).cardinality))
                rhs36_ = (iife19_(D10_DC26(p2, d_91_v41_, (d_151_v90_).f23, (d_151_v90_).f23, p2))).cf34
                rhs37_ = d_152_v91_
                lhs31_ = globalState
                lhs32_ = d_90_v40_
                lhs33_ = default__.safeIndex(338, (d_90_v40_).length(0))
                lhs34_ = d_153_v92_
                lhs35_ = default__.safeIndex(439, (d_153_v92_).length(0))
                d_151_v90_ = rhs33_
                d_152_v91_ = rhs34_
                lhs31_.f6 = rhs35_
                lhs32_[lhs33_] = rhs36_
                lhs34_[lhs35_] = rhs37_
            elif source2_.is_DC15:
                d_159___mcc_h2_ = source2_.cf18
                d_160_cf18_ = d_159___mcc_h2_
                index39_ = default__.safeIndex(360, (d_91_v41_).length(0))
                (d_91_v41_)[index39_] = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                d_77_v30_ = ((d_77_v30_) + (d_77_v30_)) + (d_77_v30_)
                d_161_v94_: _dafny.Array
                nw24_ = _dafny.Array(None, 26)
                d_161_v94_ = nw24_
                d_162_v95_: C3
                nw25_ = C3()
                nw25_.ctor__(d_92_v42_, _dafny.Map({d_77_v30_: p2}))
                d_162_v95_ = nw25_
                index40_ = default__.safeIndex(44, (d_161_v94_).length(0))
                (d_161_v94_)[index40_] = d_162_v95_
                d_163_v97_: _dafny.Map
                d_163_v97_ = _dafny.Map({d_77_v30_: len((D12_DC29(_dafny.Set({p2, p2}))).cf43)})
                d_164_v98_: C0
                nw26_ = C0()
                def iife27_():
                    coll15_ = _dafny.Map()
                    compr_15_: str
                    for compr_15_ in (d_77_v30_).Elements:
                        d_165_v96_: str = compr_15_
                        if (d_165_v96_) in (d_77_v30_):
                            coll15_[d_165_v96_] = p0
                    return _dafny.Map(coll15_)
                nw26_.ctor__(d_77_v30_, iife27_()
                , d_92_v42_, (d_163_v97_) | (d_163_v97_))
                d_164_v98_ = nw26_
                d_166_v99_: _dafny.Array
                def lambda4_(d_167_v47_):
                    def lambda5_(d_168_i9_):
                        return d_167_v47_

                    return lambda5_

                init2_ = lambda4_(d_97_v47_)
                nw27_ = _dafny.Array(None, 24)
                for i0_2_ in range(nw27_.length(0)):
                    nw27_[i0_2_] = init2_(i0_2_)
                d_166_v99_ = nw27_
                index41_ = default__.safeIndex(44, (d_161_v94_).length(0))
                rhs38_ = d_162_v95_
                rhs39_ = d_164_v98_
                rhs40_ = d_166_v99_
                lhs36_ = d_161_v94_
                lhs37_ = default__.safeIndex(44, (d_161_v94_).length(0))
                lhs36_[lhs37_] = rhs38_
                d_164_v98_ = rhs39_
                d_166_v99_ = rhs40_
                d_169_v100_: _dafny.Set
                d_169_v100_ = _dafny.Set({p1})
                d_170_v101_: _dafny.Map
                d_170_v101_ = _dafny.Map({p2: p2})
                d_171_v102_: _dafny.Map
                d_171_v102_ = _dafny.Map({(d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]: d_170_v101_})
                (globalState).f6 = len((d_170_v101_ if (d_169_v100_).issubset(d_169_v100_) else (((d_171_v102_)[p1] if (p1) in (d_171_v102_) else d_170_v101_)) | (_dafny.Map({p0: p0}))))
            elif True:
                d_172___mcc_h3_ = source2_.cf21
                d_173_cf21_ = d_172___mcc_h3_
                d_174_v103_: _dafny.Map
                d_174_v103_ = _dafny.Map({(p1 if True else p1): (d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]})
                d_174_v103_ = d_174_v103_
                index42_ = default__.safeIndex(360, (d_91_v41_).length(0))
                (d_91_v41_)[index42_] = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                index43_ = default__.safeIndex(360, (d_91_v41_).length(0))
                (d_91_v41_)[index43_] = (p2) < (p2)
                (globalState).f6 = p0
            rhs41_ = (default__.safeModuloInt(p0, 362) if (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))] else (p0) + (p2))
            rhs42_ = d_91_v41_
            lhs38_ = globalState
            lhs38_.f6 = rhs41_
            d_91_v41_ = rhs42_
            d_175_v104_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_175_v104_ = nw28_
            d_176_v105_: D6
            d_176_v105_ = D6_DC16(p1, d_175_v104_)
            d_177_v106_: D6
            d_177_v106_ = D6_DC17(d_176_v105_)
            d_178_v107_: D6
            d_178_v107_ = D6_DC17((d_176_v105_ if p1 else d_177_v106_))
            d_178_v107_ = d_178_v107_
        elif True:
            d_179_v108_: _dafny.Map
            d_179_v108_ = _dafny.Map({d_77_v30_: (d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]})
            d_180_v109_: C2
            nw29_ = C2()
            nw29_.ctor__(_dafny.CodePoint('f'), d_179_v108_)
            d_180_v109_ = nw29_
            d_181_v110_: _dafny.Seq
            d_181_v110_ = _dafny.SeqWithoutIsStrInference([d_180_v109_])
            d_182_v111_: _dafny.Set
            d_182_v111_ = _dafny.Set({len(d_181_v110_), (d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]})
            if (d_182_v111_).isdisjoint(d_182_v111_):
                d_183_v112_: _dafny.Map
                d_183_v112_ = _dafny.Map({((d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]) == ((d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]): p0})
                d_184_v113_: T0
                nw30_ = C0()
                nw30_.ctor__(_dafny.SeqWithoutIsStrInference([d_92_v42_ for d_185_i10_ in range(default__.abs(525))]), d_94_v45_, d_92_v42_, default__.fm18((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], globalState))
                d_184_v113_ = nw30_
                d_186_v114_: _dafny.Seq
                d_186_v114_ = _dafny.SeqWithoutIsStrInference([d_184_v113_])
                d_183_v112_ = (d_183_v112_).set((d_186_v114_) != (d_186_v114_), default__.safeDivisionInt((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], p0))
                d_187_v115_: _dafny.Seq
                d_187_v115_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_188_v116_: _dafny.Seq
                d_188_v116_ = _dafny.SeqWithoutIsStrInference([p1, (d_187_v115_)[default__.safeIndex(-484, len(d_187_v115_))], p1])
                d_189_v117_: _dafny.Map
                d_189_v117_ = _dafny.Map({len(d_187_v115_): p2})
                d_189_v117_ = d_189_v117_
                d_190_v118_: C3
                nw31_ = C3()
                nw31_.ctor__(_dafny.CodePoint('p'), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfumgxqqe")): p0}))
                d_190_v118_ = nw31_
                d_191_v119_: D9
                d_191_v119_ = D9_DC24(False)
                d_192_v120_: _dafny.Map
                d_192_v120_ = _dafny.Map({d_190_v118_: (p1) == (not((d_191_v119_).cf32))})
                d_192_v120_ = (d_192_v120_).set(d_190_v118_, p1)
                index44_ = default__.safeIndex(360, (d_91_v41_).length(0))
                rhs43_ = (d_180_v109_).fm3(p1, len((d_77_v30_) + (d_77_v30_)), globalState)
                rhs44_ = d_92_v42_
                rhs45_ = default__.safeModuloInt(default__.safeDivisionInt(p2, len(_dafny.SeqWithoutIsStrInference([d_91_v41_, d_91_v41_]))), (d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))])
                lhs39_ = d_91_v41_
                lhs40_ = default__.safeIndex(360, (d_91_v41_).length(0))
                lhs41_ = d_184_v113_
                lhs42_ = globalState
                lhs39_[lhs40_] = rhs43_
                lhs41_.f18 = rhs44_
                lhs42_.f6 = rhs45_
                d_193_v122_: _dafny.Seq
                def iife28_():
                    coll16_ = _dafny.Set()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(-685, 204):
                        d_194_v121_: int = compr_16_
                        if ((-685) <= (d_194_v121_)) and ((d_194_v121_) < (204)):
                            coll16_ = coll16_.union(_dafny.Set([default__.safeModuloInt(d_194_v121_, len(_dafny.Set({(d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]})))]))
                    return _dafny.Set(coll16_)
                d_193_v122_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0}), (iife28_()
                ).intersection(d_182_v111_)])
                d_193_v122_ = (_dafny.SeqWithoutIsStrInference([d_182_v111_])) + (_dafny.SeqWithoutIsStrInference([d_182_v111_, d_182_v111_]))
            elif True:
                r2 = (0) - (default__.safeDivisionInt(((0) - (p0)) - (p2), (d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]))
                d_195_v123_: _dafny.Array
                def lambda6_(d_196_v30_, d_197_p2_):
                    def lambda7_(d_198_i11_):
                        return (d_196_v30_)[default__.safeIndex(d_197_p2_, len(d_196_v30_))]

                    return lambda7_

                init3_ = lambda6_(d_77_v30_, p2)
                nw32_ = _dafny.Array(None, 7)
                for i0_3_ in range(nw32_.length(0)):
                    nw32_[i0_3_] = init3_(i0_3_)
                d_195_v123_ = nw32_
                index45_ = default__.safeIndex(967, (d_195_v123_).length(0))
                (d_195_v123_)[index45_] = d_92_v42_
                index46_ = default__.safeIndex(967, (d_195_v123_).length(0))
                (d_195_v123_)[index46_] = (D3_DC8(d_92_v42_)).cf10
                d_199_v124_: _dafny.Map
                d_199_v124_ = _dafny.Map({d_182_v111_: _dafny.SeqWithoutIsStrInference([p1])})
                d_200_v125_: _dafny.Seq
                d_200_v125_ = _dafny.SeqWithoutIsStrInference([d_182_v111_])
                d_201_v126_: D9
                d_201_v126_ = D9_DC23(d_200_v125_, p2, p1, (d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))])
                d_202_v127_: _dafny.Seq
                d_202_v127_ = _dafny.SeqWithoutIsStrInference([(d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))], (d_201_v126_).cf30])
                d_203_v128_: _dafny.Seq
                d_203_v128_ = _dafny.SeqWithoutIsStrInference([d_202_v127_])
                d_204_v129_: _dafny.Map
                d_204_v129_ = _dafny.Map({d_202_v127_: (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]})
                d_205_v130_: _dafny.Map
                d_205_v130_ = _dafny.Map({d_204_v129_: d_202_v127_})
                d_206_v132_: _dafny.Array
                nw33_ = _dafny.Array(None, 6)
                nw33_[int(0)] = (((d_199_v124_)[d_182_v111_] if (d_182_v111_) in (d_199_v124_) else d_202_v127_)) + (d_202_v127_)
                nw33_[int(1)] = d_202_v127_
                nw33_[int(2)] = d_202_v127_
                nw33_[int(3)] = (d_203_v128_)[default__.safeIndex((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], len(d_203_v128_))]
                nw33_[int(4)] = d_202_v127_
                def iife29_():
                    coll17_ = _dafny.Map()
                    compr_17_: _dafny.Seq
                    for compr_17_ in (_dafny.SeqWithoutIsStrInference([d_202_v127_ for d_207_i12_ in range(default__.abs(-229))])).Elements:
                        d_208_v131_: _dafny.Seq = compr_17_
                        if (d_208_v131_) in (_dafny.SeqWithoutIsStrInference([d_202_v127_ for d_207_i12_ in range(default__.abs(-229))])):
                            coll17_[d_208_v131_] = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                    return _dafny.Map(coll17_)
                def iife30_():
                    coll18_ = _dafny.Map()
                    compr_18_: _dafny.Seq
                    for compr_18_ in (_dafny.SeqWithoutIsStrInference([d_202_v127_ for d_209_i12_ in range(default__.abs(-229))])).Elements:
                        d_210_v131_: _dafny.Seq = compr_18_
                        if (d_210_v131_) in (_dafny.SeqWithoutIsStrInference([d_202_v127_ for d_209_i12_ in range(default__.abs(-229))])):
                            coll18_[d_210_v131_] = (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]
                    return _dafny.Map(coll18_)
                nw33_[int(5)] = ((d_205_v130_)[iife29_()
                ] if (iife30_()
                ) in (d_205_v130_) else _dafny.SeqWithoutIsStrInference([p1, (d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))]]))
                d_206_v132_ = nw33_
                d_206_v132_ = d_206_v132_
                d_211_v133_: _dafny.Set
                d_211_v133_ = _dafny.Set({p1})
                d_212_v134_: _dafny.Seq
                d_212_v134_ = _dafny.SeqWithoutIsStrInference([len(d_211_v133_)])
                d_213_v135_: _dafny.Seq
                d_213_v135_ = _dafny.SeqWithoutIsStrInference([d_179_v108_])
                d_214_v136_: C1
                nw34_ = C1()
                nw34_.ctor__((0) - ((d_212_v134_)[default__.safeIndex((0) - (p2), len(d_212_v134_))]), p1, (d_195_v123_)[default__.safeIndex(967, (d_195_v123_).length(0))], (d_213_v135_)[default__.safeIndex(len(d_202_v127_), len(d_213_v135_))])
                d_214_v136_ = nw34_
                r2 = p2
            def iife31_():
                coll19_ = _dafny.Set()
                compr_19_: int
                for compr_19_ in (d_182_v111_).Elements:
                    d_215_v137_: int = compr_19_
                    if (d_215_v137_) in (d_182_v111_):
                        coll19_ = coll19_.union(_dafny.Set([default__.safeModuloInt(d_215_v137_, 36)]))
                return _dafny.Set(coll19_)
            r1 = (iife31_()
            ).isdisjoint(d_182_v111_)
            d_216_v138_: _dafny.Array
            nw35_ = _dafny.Array(_dafny.MultiSet({}), 11)
            d_216_v138_ = nw35_
            d_217_v139_: _dafny.MultiSet
            d_217_v139_ = _dafny.MultiSet([d_91_v41_])
            d_218_v140_: _dafny.Map
            d_218_v140_ = _dafny.Map({True: d_217_v139_})
            index47_ = default__.safeIndex(166, (d_216_v138_).length(0))
            (d_216_v138_)[index47_] = ((d_218_v140_)[p1] if (p1) in (d_218_v140_) else d_217_v139_)
            index48_ = default__.safeIndex(166, (d_216_v138_).length(0))
            (d_216_v138_)[index48_] = d_217_v139_
            d_219_v141_: _dafny.Seq
            d_219_v141_ = _dafny.SeqWithoutIsStrInference([(d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))], default__.fm0(globalState)])
            (globalState).f1 = ((d_219_v141_ if (d_180_v109_).fm3(p1, p0, globalState) else d_219_v141_)) + (d_219_v141_)
            d_220_v142_: _dafny.Set
            d_220_v142_ = _dafny.Set({p1})
            d_221_v143_: D12
            d_221_v143_ = D12_DC31(p1, _dafny.SeqWithoutIsStrInference([d_92_v42_ for d_222_i13_ in range(default__.abs(758))]), len(default__.fm30((d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))], globalState)), p1)
            rhs46_ = ((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))] if not((d_220_v142_).isdisjoint(d_220_v142_)) else len(((d_221_v143_).cf47) + (d_77_v30_)))
            rhs47_ = (((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]) * (len(_dafny.Map({(0) - ((d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]): d_91_v41_}))) if p1 else p0)
            lhs43_ = globalState
            lhs43_.f6 = rhs46_
            r2 = rhs47_
        d_223_v144_: _dafny.Map
        d_223_v144_ = _dafny.Map({((d_77_v30_).set(default__.safeIndex(p0, len(d_77_v30_)), d_92_v42_)).set(default__.safeIndex(p2, len((d_77_v30_).set(default__.safeIndex(p0, len(d_77_v30_)), d_92_v42_))), d_92_v42_): p2})
        d_224_v145_: C2
        nw36_ = C2()
        nw36_.ctor__(d_92_v42_, d_223_v144_)
        d_224_v145_ = nw36_
        r0 = d_91_v41_
        r1 = ((d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))] if p1 else p1)
        d_225_v146_: _dafny.Set
        d_225_v146_ = _dafny.Set({(d_90_v40_)[default__.safeIndex(338, (d_90_v40_).length(0))]})
        r2 = (d_224_v145_).fm4((d_91_v41_)[default__.safeIndex(360, (d_91_v41_).length(0))], d_92_v42_, default__.fm10(globalState), (d_225_v146_).ispropersubset(d_225_v146_), globalState)
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_226_v0_: int
        d_226_v0_ = 140
        d_227_v1_: _dafny.Seq
        d_227_v1_ = _dafny.SeqWithoutIsStrInference([(0) - (d_226_v0_)])
        d_228_v2_: bool
        d_228_v2_ = True
        d_229_v3_: _dafny.Map
        d_229_v3_ = _dafny.Map({not(True): d_228_v2_})
        d_230_v4_: str
        d_230_v4_ = _dafny.CodePoint('q')
        d_231_v5_: _dafny.Map
        d_231_v5_ = _dafny.Map({d_230_v4_: (0) - (d_226_v0_)})
        d_232_v7_: _dafny.Array
        def lambda8_(d_233_v0_):
            def lambda9_(d_234_i0_):
                def iife32_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(399, -703):
                        d_235_v6_: int = compr_20_
                        if ((399) <= (d_235_v6_)) and ((d_235_v6_) < (-703)):
                            coll20_[default__.safeModuloInt(d_235_v6_, 668)] = d_233_v0_
                    return _dafny.Map(coll20_)
                return iife32_()
                

            return lambda9_

        init4_ = lambda8_(d_226_v0_)
        nw37_ = _dafny.Array(None, 10)
        for i0_4_ in range(nw37_.length(0)):
            nw37_[i0_4_] = init4_(i0_4_)
        d_232_v7_ = nw37_
        d_236_v8_: _dafny.Set
        d_236_v8_ = _dafny.Set({d_228_v2_})
        d_237_v9_: _dafny.Array
        nw38_ = _dafny.Array(None, 12)
        nw38_[int(0)] = d_226_v0_
        nw38_[int(1)] = d_226_v0_
        nw38_[int(2)] = d_226_v0_
        nw38_[int(3)] = d_226_v0_
        nw38_[int(4)] = d_226_v0_
        nw38_[int(5)] = 44
        nw38_[int(6)] = d_226_v0_
        nw38_[int(7)] = d_226_v0_
        nw38_[int(8)] = d_226_v0_
        nw38_[int(9)] = d_226_v0_
        nw38_[int(10)] = d_226_v0_
        nw38_[int(11)] = len(d_236_v8_)
        d_237_v9_ = nw38_
        d_238_v10_: _dafny.Array
        nw39_ = _dafny.Array(False, 5)
        d_238_v10_ = nw39_
        d_239_v11_: _dafny.Map
        d_239_v11_ = _dafny.Map({d_237_v9_: d_238_v10_})
        d_240_v12_: _dafny.Seq
        d_240_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oalh"))
        d_241_v13_: _dafny.Seq
        d_241_v13_ = _dafny.SeqWithoutIsStrInference([d_237_v9_])
        d_242_v14_: _dafny.Seq
        d_242_v14_ = _dafny.SeqWithoutIsStrInference([False, d_228_v2_])
        d_243_v15_: _dafny.Map
        d_243_v15_ = _dafny.Map({d_226_v0_: d_228_v2_})
        d_244_v16_: _dafny.Map
        d_244_v16_ = _dafny.Map({d_226_v0_: ((d_243_v15_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgoah")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgoah")))) in (d_243_v15_) else d_228_v2_)})
        d_245_v17_: _dafny.Seq
        d_245_v17_ = _dafny.SeqWithoutIsStrInference([d_244_v16_])
        d_246_globalState_: GlobalState
        nw40_ = GlobalState()
        nw40_.ctor__(515, d_227_v1_, d_227_v1_, d_229_v3_, 523, d_231_v5_, 914, d_232_v7_, d_239_v11_, d_240_v12_, _dafny.Map({d_238_v10_: d_228_v2_}), (_dafny.SeqWithoutIsStrInference([d_237_v9_, d_237_v9_, d_237_v9_, d_237_v9_, d_237_v9_])) + (d_241_v13_), False, d_242_v14_, False, 682, (_dafny.SeqWithoutIsStrInference([d_244_v16_])) + (d_245_v17_), True)
        d_246_globalState_ = nw40_
        hi1_ = (0) - (default__.safeModuloInt(902, d_226_v0_))
        for d_247_i1_ in range(d_226_v0_, hi1_):
            d_228_v2_ = ((d_247_i1_ if d_228_v2_ else (0) - (d_247_i1_))) < (((0) - (d_226_v0_)) * (d_226_v0_))
            d_248_v18_: _dafny.Seq
            d_248_v18_ = _dafny.SeqWithoutIsStrInference([d_240_v12_, d_240_v12_])
            d_248_v18_ = (_dafny.SeqWithoutIsStrInference([d_240_v12_])) + (d_248_v18_)
            d_249_v19_: _dafny.Set
            d_249_v19_ = _dafny.Set({d_237_v9_})
            rhs48_ = d_228_v2_
            rhs49_ = d_228_v2_
            rhs50_ = d_227_v1_
            rhs51_ = default__.safeDivisionInt(len((d_249_v19_).intersection(d_249_v19_)), len(d_227_v1_))
            lhs44_ = d_246_globalState_
            d_228_v2_ = rhs48_
            d_228_v2_ = rhs49_
            lhs44_.f1 = rhs50_
            d_226_v0_ = rhs51_
            index49_ = default__.safeIndex(545, (d_238_v10_).length(0))
            (d_238_v10_)[index49_] = d_228_v2_
            d_250_v20_: _dafny.Map
            d_250_v20_ = _dafny.Map({default__.fm0(d_246_globalState_): (d_230_v4_ if d_228_v2_ else d_230_v4_)})
            d_251_v21_: D0
            d_251_v21_ = D0_DC0(default__.fm0(d_246_globalState_))
            index50_ = default__.safeIndex(545, (d_238_v10_).length(0))
            rhs52_ = ((d_250_v20_)[288] if (288) in (d_250_v20_) else d_230_v4_)
            rhs53_ = ((d_251_v21_).cf0) - (d_226_v0_)
            rhs54_ = d_228_v2_
            rhs55_ = (False if False else d_228_v2_)
            rhs56_ = (d_251_v21_).cf0
            lhs45_ = d_246_globalState_
            lhs46_ = d_238_v10_
            lhs47_ = default__.safeIndex(545, (d_238_v10_).length(0))
            d_230_v4_ = rhs52_
            lhs45_.f6 = rhs53_
            lhs46_[lhs47_] = rhs54_
            d_228_v2_ = rhs55_
            d_226_v0_ = rhs56_
        d_252_i2_: int
        d_252_i2_ = 0
        with _dafny.label("0"):
            while d_228_v2_:
                with _dafny.c_label("0"):
                    if (d_252_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_252_i2_ = (d_252_i2_) + (1)
                    d_229_v3_ = (d_229_v3_).set(not((904) <= (d_226_v0_)), d_228_v2_)
                    rhs57_ = d_230_v4_
                    rhs58_ = d_226_v0_
                    lhs48_ = d_246_globalState_
                    d_230_v4_ = rhs57_
                    lhs48_.f6 = rhs58_
                    d_253_v22_: _dafny.Map
                    d_253_v22_ = _dafny.Map({d_238_v10_: True})
                    (d_246_globalState_).f6 = (d_226_v0_ if d_228_v2_ else ((0) - (-533)) * (len(d_253_v22_)))
                    index51_ = default__.safeIndex(307, (d_237_v9_).length(0))
                    (d_237_v9_)[index51_] = d_226_v0_
                    index52_ = default__.safeIndex(307, (d_237_v9_).length(0))
                    (d_237_v9_)[index52_] = (-29) - (d_226_v0_)
                    pass
            pass
        index53_ = default__.safeIndex(465, (d_237_v9_).length(0))
        (d_237_v9_)[index53_] = d_226_v0_
        index54_ = default__.safeIndex(465, (d_237_v9_).length(0))
        (d_237_v9_)[index54_] = default__.safeModuloInt(d_226_v0_, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxertem")))))
        d_254_v23_: _dafny.Array
        d_255_v24_: bool
        d_256_v25_: int
        out0_: _dafny.Array
        out1_: bool
        out2_: int
        out0_, out1_, out2_ = default__.m0((0) - ((d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]), default__.fm1(((d_243_v15_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjgtu")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjgtu")))) in (d_243_v15_) else d_228_v2_), d_230_v4_, -261, d_246_globalState_), (len((d_240_v12_).set(default__.safeIndex(d_226_v0_, len(d_240_v12_)), d_230_v4_))) * ((0) - (d_226_v0_)), d_246_globalState_)
        d_254_v23_ = out0_
        d_255_v24_ = out1_
        d_256_v25_ = out2_
        d_240_v12_ = d_240_v12_
        d_257_v26_: _dafny.Map
        d_257_v26_ = _dafny.Map({d_255_v24_: d_242_v14_})
        index55_ = default__.safeIndex(465, (d_237_v9_).length(0))
        rhs59_ = default__.fm0(d_246_globalState_)
        rhs60_ = (((d_257_v26_)[d_255_v24_] if (d_255_v24_) in (d_257_v26_) else d_242_v14_)) + (d_242_v14_)
        lhs49_ = d_237_v9_
        lhs50_ = default__.safeIndex(465, (d_237_v9_).length(0))
        lhs49_[lhs50_] = rhs59_
        d_242_v14_ = rhs60_
        d_255_v24_ = d_228_v2_
        d_258_v27_: _dafny.Array
        def lambda10_(d_259_v4_):
            def lambda11_(d_260_i3_):
                return d_259_v4_

            return lambda11_

        init5_ = lambda10_(d_230_v4_)
        nw41_ = _dafny.Array(None, 1)
        for i0_5_ in range(nw41_.length(0)):
            nw41_[i0_5_] = init5_(i0_5_)
        d_258_v27_ = nw41_
        d_261_v28_: _dafny.Map
        d_261_v28_ = _dafny.Map({(d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]: d_258_v27_})
        if (_dafny.Set({default__.fm1(d_228_v2_, d_230_v4_, len(d_261_v28_), d_246_globalState_), d_255_v24_, ((d_229_v3_)[d_255_v24_] if (d_255_v24_) in (d_229_v3_) else d_255_v24_)})).issubset(default__.fm2(d_246_globalState_)):
            d_262_v29_: D0
            d_262_v29_ = D0_DC1()
            d_262_v29_ = d_262_v29_
            d_263_v30_: _dafny.Map
            d_263_v30_ = _dafny.Map({d_228_v2_: 531})
            d_264_v31_: C3
            nw42_ = C3()
            nw42_.ctor__(default__.fm20(False, d_246_globalState_), _dafny.Map({d_240_v12_: len(d_263_v30_)}))
            d_264_v31_ = nw42_
            def iife33_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(584, 334):
                    d_265_v32_: int = compr_21_
                    if ((584) <= (d_265_v32_)) and ((d_265_v32_) < (334)):
                        coll21_[(d_265_v32_) - (d_226_v0_)] = D5_DC12(_dafny.MultiSet([d_228_v2_]))
                return _dafny.Map(coll21_)
            (d_246_globalState_).f6 = (len(iife33_()
            )) * ((d_256_v25_) + ((d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]))
            index56_ = default__.safeIndex(790, (d_254_v23_).length(0))
            (d_254_v23_)[index56_] = not((d_255_v24_ if d_255_v24_ else d_228_v2_))
            d_266_v33_: _dafny.Set
            d_266_v33_ = _dafny.Set({d_256_v25_})
            d_267_v34_: _dafny.Seq
            d_267_v34_ = _dafny.SeqWithoutIsStrInference([d_266_v33_])
            d_268_v35_: _dafny.Map
            d_268_v35_ = _dafny.Map({d_226_v0_: d_267_v34_})
            index57_ = default__.safeIndex(790, (d_254_v23_).length(0))
            rhs61_ = not(d_255_v24_)
            rhs62_ = d_228_v2_
            rhs63_ = default__.fm25(d_229_v3_, D9_DC23(((d_268_v35_)[d_226_v0_] if (d_226_v0_) in (d_268_v35_) else d_267_v34_), (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))], d_228_v2_, d_226_v0_), default__.fm1(False, d_230_v4_, d_226_v0_, d_246_globalState_), (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))], d_246_globalState_)
            rhs64_ = not (d_228_v2_) or ((d_256_v25_) < (len(default__.fm26(d_246_globalState_))))
            rhs65_ = not(d_255_v24_)
            lhs51_ = d_254_v23_
            lhs52_ = default__.safeIndex(790, (d_254_v23_).length(0))
            lhs51_[lhs52_] = rhs61_
            d_228_v2_ = rhs62_
            d_262_v29_ = rhs63_
            d_228_v2_ = rhs64_
            d_255_v24_ = rhs65_
            d_228_v2_ = (d_256_v25_) != (d_256_v25_)
        elif True:
            d_269_v36_: _dafny.Array
            nw43_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
            d_269_v36_ = nw43_
            d_270_v37_: _dafny.Map
            d_270_v37_ = _dafny.Map({d_269_v36_: d_238_v10_})
            d_270_v37_ = (d_270_v37_).set(d_269_v36_, d_254_v23_)
            d_271_v38_: _dafny.Array
            def lambda12_(d_272_v14_):
                def lambda13_(d_273_i4_):
                    return (d_272_v14_) + (d_272_v14_)

                return lambda13_

            init6_ = lambda12_(d_242_v14_)
            nw44_ = _dafny.Array(None, 16)
            for i0_6_ in range(nw44_.length(0)):
                nw44_[i0_6_] = init6_(i0_6_)
            d_271_v38_ = nw44_
            d_271_v38_ = d_271_v38_
            d_274_v39_: _dafny.Array
            nw45_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_274_v39_ = nw45_
            d_275_v40_: D6
            d_275_v40_ = D6_DC16(d_228_v2_, d_274_v39_)
            source3_ = d_275_v40_
            if source3_.is_DC16:
                d_276___mcc_h0_ = source3_.cf19
                d_277___mcc_h1_ = source3_.cf20
                d_278_cf20_ = d_277___mcc_h1_
                d_279_cf19_ = d_276___mcc_h0_
                d_280_v41_: D7
                d_280_v41_ = D7_DC18(d_236_v8_)
                index58_ = default__.safeIndex(465, (d_237_v9_).length(0))
                (d_237_v9_)[index58_] = (d_256_v25_) - ((len((d_280_v41_).cf22)) * (len(_dafny.Map({d_279_cf19_: d_256_v25_}))))
                d_279_cf19_ = d_255_v24_
                d_281_v42_: _dafny.MultiSet
                d_281_v42_ = _dafny.MultiSet([(d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]])
                index59_ = default__.safeIndex(284, (d_238_v10_).length(0))
                (d_238_v10_)[index59_] = not((len(d_240_v12_)) not in (d_281_v42_))
                index60_ = default__.safeIndex(284, (d_238_v10_).length(0))
                (d_238_v10_)[index60_] = d_228_v2_
                d_226_v0_ = default__.fm0(d_246_globalState_)
            elif source3_.is_DC15:
                d_282___mcc_h2_ = source3_.cf18
                d_283_cf18_ = d_282___mcc_h2_
                d_284_v43_: _dafny.Array
                d_285_v44_: bool
                d_286_v45_: int
                out3_: _dafny.Array
                out4_: bool
                out5_: int
                out3_, out4_, out5_ = default__.m0(len(d_240_v12_), (default__.fm0(d_246_globalState_)) in (_dafny.SeqWithoutIsStrInference([629 for d_287_i5_ in range(default__.abs(158))])), len(d_240_v12_), d_246_globalState_)
                d_284_v43_ = out3_
                d_285_v44_ = out4_
                d_286_v45_ = out5_
                d_288_v46_: C3
                nw46_ = C3()
                nw46_.ctor__(d_230_v4_, _dafny.Map({d_240_v12_: d_226_v0_}))
                d_288_v46_ = nw46_
                d_288_v46_ = d_288_v46_
                d_289_v47_: _dafny.MultiSet
                d_289_v47_ = _dafny.MultiSet([d_285_v44_, True])
                d_290_v48_: D3
                d_290_v48_ = D3_DC10((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_228_v2_]))).ispropersubset(d_289_v47_))
                d_290_v48_ = d_290_v48_
                d_291_v49_: _dafny.Array
                nw47_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_291_v49_ = nw47_
                d_292_v50_: _dafny.Map
                d_292_v50_ = _dafny.Map({d_228_v2_: d_269_v36_})
                index61_ = default__.safeIndex(151, (d_291_v49_).length(0))
                (d_291_v49_)[index61_] = ((d_292_v50_)[d_285_v44_] if (d_285_v44_) in (d_292_v50_) else d_269_v36_)
                index62_ = default__.safeIndex(151, (d_291_v49_).length(0))
                (d_291_v49_)[index62_] = d_269_v36_
            elif True:
                d_293___mcc_h3_ = source3_.cf21
                d_294_cf21_ = d_293___mcc_h3_
                d_295_v51_: _dafny.Map
                d_295_v51_ = _dafny.Map({d_240_v12_: len((D9_DC22(_dafny.SeqWithoutIsStrInference([d_230_v4_ for d_296_i6_ in range(default__.abs(932))]))).cf27)})
                d_297_v52_: C3
                nw48_ = C3()
                nw48_.ctor__(d_230_v4_, d_295_v51_)
                d_297_v52_ = nw48_
                d_297_v52_ = d_297_v52_
                d_298_v53_: _dafny.Map
                d_298_v53_ = _dafny.Map({d_255_v24_: (d_237_v9_ if d_228_v2_ else (d_241_v13_)[default__.safeIndex(d_226_v0_, len(d_241_v13_))])})
                d_237_v9_ = ((d_298_v53_)[(d_297_v52_).fm3(d_228_v2_, (0) - (d_226_v0_), d_246_globalState_)] if ((d_297_v52_).fm3(d_228_v2_, (0) - (d_226_v0_), d_246_globalState_)) in (d_298_v53_) else d_237_v9_)
                d_299_v54_: C3
                nw49_ = C3()
                nw49_.ctor__(d_230_v4_, (d_295_v51_) | (d_295_v51_))
                d_299_v54_ = nw49_
                d_300_v55_: D0
                d_301_v56_: int
                d_302_v57_: bool
                d_303_v58_: bool
                out6_: D0
                out7_: int
                out8_: bool
                out9_: bool
                out6_, out7_, out8_, out9_ = (d_297_v52_).m1(default__.fm1(d_228_v2_, _dafny.CodePoint('m'), len(d_298_v53_), d_246_globalState_), d_246_globalState_)
                d_300_v55_ = out6_
                d_301_v56_ = out7_
                d_302_v57_ = out8_
                d_303_v58_ = out9_
            d_304_v59_: _dafny.Map
            d_304_v59_ = _dafny.Map({d_240_v12_: d_226_v0_})
            d_305_v60_: C1
            nw50_ = C1()
            nw50_.ctor__((d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))], d_255_v24_, d_230_v4_, d_304_v59_)
            d_305_v60_ = nw50_
            d_306_v61_: _dafny.Map
            d_306_v61_ = _dafny.Map({d_305_v60_: (0) - (default__.safeModuloInt(d_226_v0_, d_256_v25_))})
            d_306_v61_ = (d_306_v61_).set(d_305_v60_, d_256_v25_)
            d_307_v62_: _dafny.MultiSet
            d_307_v62_ = _dafny.MultiSet([d_255_v24_, True])
            d_308_v63_: _dafny.Map
            d_308_v63_ = _dafny.Map({d_230_v4_: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_305_v60_).f23]))).intersection(d_307_v62_)})
            d_308_v63_ = (d_308_v63_).set(d_230_v4_, d_307_v62_)
        d_309_v64_: _dafny.Array
        nw51_ = _dafny.Array(D7.default()(), 18)
        d_309_v64_ = nw51_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_309_v64_).length(0)):
            d_310_i7_: int = guard_loop_0_
            if (True) and (((0) <= (d_310_i7_)) and ((d_310_i7_) < ((d_309_v64_).length(0)))):
                (d_309_v64_)[(d_310_i7_)] = D7_DC20(d_228_v2_, d_228_v2_)
        d_311_v65_: _dafny.Map
        d_311_v65_ = _dafny.Map({d_256_v25_: 458})
        d_312_i8_: int
        d_312_i8_ = 0
        with _dafny.label("1"):
            while (((0) - (len(d_311_v65_))) * (len(d_240_v12_))) >= (d_256_v25_):
                with _dafny.c_label("1"):
                    if (d_312_i8_) >= (100):
                        raise _dafny.Break("1")
                    d_312_i8_ = (d_312_i8_) + (1)
                    d_256_v25_ = len(((d_227_v1_).set(default__.safeIndex(104, len(d_227_v1_)), (0) - (d_256_v25_))).set(default__.safeIndex(default__.safeDivisionInt(d_256_v25_, d_226_v0_), len((d_227_v1_).set(default__.safeIndex(104, len(d_227_v1_)), (0) - (d_256_v25_)))), 123))
                    d_313_v66_: _dafny.Set
                    d_313_v66_ = _dafny.Set({(d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))], (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]})
                    d_314_v67_: _dafny.Seq
                    d_314_v67_ = _dafny.SeqWithoutIsStrInference([d_313_v66_, d_313_v66_])
                    def iife34_():
                        coll22_ = _dafny.Set()
                        compr_22_: int
                        for compr_22_ in _dafny.IntegerRange(-240, 336):
                            d_315_v68_: int = compr_22_
                            if ((-240) <= (d_315_v68_)) and ((d_315_v68_) < (336)):
                                coll22_ = coll22_.union(_dafny.Set([(d_315_v68_) - (len(d_243_v15_))]))
                        return _dafny.Set(coll22_)
                    d_314_v67_ = (_dafny.SeqWithoutIsStrInference([d_313_v66_, d_313_v66_, d_313_v66_, d_313_v66_, _dafny.Set({d_226_v0_})]) if d_228_v2_ else _dafny.SeqWithoutIsStrInference([iife34_()
                    , _dafny.Set({90}), d_313_v66_, d_313_v66_, d_313_v66_]))
                    d_316_v70_: _dafny.Seq
                    d_316_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_230_v4_ for d_317_i9_ in range(default__.abs(216))])])
                    d_318_v71_: C2
                    nw52_ = C2()
                    def iife35_():
                        coll23_ = _dafny.Map()
                        compr_23_: _dafny.Seq
                        for compr_23_ in (d_316_v70_).Elements:
                            d_319_v69_: _dafny.Seq = compr_23_
                            if (d_319_v69_) in (d_316_v70_):
                                coll23_[d_319_v69_] = (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]
                        return _dafny.Map(coll23_)
                    nw52_.ctor__(d_230_v4_, iife35_()
                    )
                    d_318_v71_ = nw52_
                    d_311_v65_ = (d_311_v65_).set((0) - ((d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]), default__.safeDivisionInt(378, d_256_v25_))
                    pass
            pass
        d_320_v72_: _dafny.Array
        nw53_ = _dafny.Array(None, 27)
        d_320_v72_ = nw53_
        d_321_v73_: _dafny.Map
        d_321_v73_ = _dafny.Map({d_320_v72_: d_227_v1_})
        d_321_v73_ = (d_321_v73_).set(d_320_v72_, d_227_v1_)
        d_322_v74_: D1
        d_322_v74_ = D1_DC6((d_240_v12_) <= (d_240_v12_), d_255_v24_)
        source4_ = d_322_v74_
        if source4_.is_DC4:
            d_323___mcc_h4_ = source4_.cf3
            d_324___mcc_h5_ = source4_.cf4
            d_325___mcc_h6_ = source4_.cf5
            d_326_cf5_ = d_325___mcc_h6_
            d_327_cf4_ = d_324___mcc_h5_
            d_328_cf3_ = d_323___mcc_h4_
            d_329_v75_: _dafny.Map
            d_329_v75_ = _dafny.Map({d_240_v12_: d_327_cf4_})
            d_330_v76_: T0
            nw54_ = C0()
            nw54_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikrmgew")), _dafny.Map({_dafny.CodePoint('o'): (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]}), d_230_v4_, d_329_v75_)
            d_330_v76_ = nw54_
            d_331_v77_: _dafny.MultiSet
            d_331_v77_ = _dafny.MultiSet([d_330_v76_])
            d_332_v78_: _dafny.Seq
            d_332_v78_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_330_v76_]), _dafny.MultiSet([d_330_v76_]), (d_331_v77_).set(d_330_v76_, default__.abs(default__.fm0(d_246_globalState_))), _dafny.MultiSet([d_330_v76_, d_330_v76_])])
            d_332_v78_ = d_332_v78_
            d_333_v79_: _dafny.Set
            d_333_v79_ = _dafny.Set({788, d_256_v25_, 979, d_328_cf3_, (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]})
            d_334_v80_: _dafny.Seq
            d_334_v80_ = _dafny.SeqWithoutIsStrInference([d_333_v79_, d_333_v79_])
            d_335_v81_: _dafny.Array
            d_336_v82_: bool
            d_337_v83_: int
            out10_: _dafny.Array
            out11_: bool
            out12_: int
            out10_, out11_, out12_ = default__.m0((0) - ((D9_DC23(d_334_v80_, d_256_v25_, d_228_v2_, 277)).cf31), d_228_v2_, d_226_v0_, d_246_globalState_)
            d_335_v81_ = out10_
            d_336_v82_ = out11_
            d_337_v83_ = out12_
            if d_228_v2_:
                d_336_v82_ = not((d_326_cf5_) == (default__.fm1(d_255_v24_, d_330_v76_.f18, d_226_v0_, d_246_globalState_)))
                (d_246_globalState_).f6 = d_226_v0_
                d_338_v86_: _dafny.Seq
                d_338_v86_ = _dafny.SeqWithoutIsStrInference([d_238_v10_])
                d_339_v87_: _dafny.Array
                nw55_ = _dafny.Array(None, 29)
                nw55_[int(0)] = d_335_v81_
                nw55_[int(1)] = d_335_v81_
                nw55_[int(2)] = d_254_v23_
                nw55_[int(3)] = d_238_v10_
                nw55_[int(4)] = d_335_v81_
                nw55_[int(5)] = d_335_v81_
                nw55_[int(6)] = d_238_v10_
                nw55_[int(7)] = d_238_v10_
                nw55_[int(8)] = d_335_v81_
                nw55_[int(9)] = d_335_v81_
                nw55_[int(10)] = d_335_v81_
                nw55_[int(11)] = d_238_v10_
                nw55_[int(12)] = d_254_v23_
                nw55_[int(13)] = d_254_v23_
                nw55_[int(14)] = d_254_v23_
                nw55_[int(15)] = d_335_v81_
                nw55_[int(16)] = d_335_v81_
                nw55_[int(17)] = d_238_v10_
                nw55_[int(18)] = d_238_v10_
                nw55_[int(19)] = d_254_v23_
                nw55_[int(20)] = d_254_v23_
                nw55_[int(21)] = d_238_v10_
                nw55_[int(22)] = d_254_v23_
                nw55_[int(23)] = d_254_v23_
                nw55_[int(24)] = d_238_v10_
                nw55_[int(25)] = d_254_v23_
                nw55_[int(26)] = d_254_v23_
                nw55_[int(27)] = d_335_v81_
                nw55_[int(28)] = (d_338_v86_)[default__.safeIndex(len(d_236_v8_), len(d_338_v86_))]
                d_339_v87_ = nw55_
                d_340_v88_: D6
                d_340_v88_ = D6_DC16(d_326_cf5_, d_339_v87_)
                d_341_v89_: D6
                d_341_v89_ = D6_DC17(d_340_v88_)
                d_342_v90_: _dafny.Map
                d_342_v90_ = _dafny.Map({d_341_v89_: d_326_cf5_})
                d_343_v93_: _dafny.Array
                nw56_ = _dafny.Array(None, 18)
                nw56_[int(0)] = _dafny.Map({(0) - (len(d_240_v12_)): d_228_v2_})
                nw56_[int(1)] = d_243_v15_
                nw56_[int(2)] = d_244_v16_
                nw56_[int(3)] = d_244_v16_
                nw56_[int(4)] = d_244_v16_
                nw56_[int(5)] = _dafny.Map({len(d_236_v8_): default__.fm1(True, d_230_v4_, d_337_v83_, d_246_globalState_)})
                nw56_[int(6)] = _dafny.Map({d_327_cf4_: True})
                def iife36_():
                    coll24_ = _dafny.Map()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(919, 570):
                        d_344_v84_: int = compr_24_
                        if ((919) <= (d_344_v84_)) and ((d_344_v84_) < (570)):
                            coll24_[default__.safeModuloInt(d_344_v84_, d_226_v0_)] = d_326_cf5_
                    return _dafny.Map(coll24_)
                nw56_[int(7)] = iife36_()
                
                nw56_[int(8)] = d_244_v16_
                nw56_[int(9)] = d_243_v15_
                nw56_[int(10)] = d_244_v16_
                def iife37_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in _dafny.IntegerRange(699, 764):
                        d_345_v85_: int = compr_25_
                        if ((699) <= (d_345_v85_)) and ((d_345_v85_) < (764)):
                            coll25_[(d_345_v85_) + (d_256_v25_)] = d_326_cf5_
                    return _dafny.Map(coll25_)
                nw56_[int(11)] = (iife37_()
                ) | (d_243_v15_)
                nw56_[int(12)] = (_dafny.Map({d_226_v0_: not(((d_342_v90_)[d_341_v89_] if (d_341_v89_) in (d_342_v90_) else d_228_v2_))})) | (d_244_v16_)
                def iife38_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(58, 424):
                        d_346_v91_: int = compr_26_
                        if ((58) <= (d_346_v91_)) and ((d_346_v91_) < (424)):
                            coll26_[(d_346_v91_) - (335)] = True
                    return _dafny.Map(coll26_)
                nw56_[int(13)] = iife38_()
                
                nw56_[int(14)] = d_243_v15_
                nw56_[int(15)] = d_244_v16_
                def iife39_():
                    coll27_ = _dafny.Map()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(230, 64):
                        d_347_v92_: int = compr_27_
                        if ((230) <= (d_347_v92_)) and ((d_347_v92_) < (64)):
                            coll27_[(d_347_v92_) - (d_226_v0_)] = d_255_v24_
                    return _dafny.Map(coll27_)
                nw56_[int(16)] = iife39_()
                
                nw56_[int(17)] = d_244_v16_
                d_343_v93_ = nw56_
                index63_ = default__.safeIndex(567, (d_343_v93_).length(0))
                (d_343_v93_)[index63_] = _dafny.Map({len(d_240_v12_): not(d_255_v24_)})
                index64_ = default__.safeIndex(567, (d_343_v93_).length(0))
                (d_343_v93_)[index64_] = (d_243_v15_ if default__.fm1(d_336_v82_, d_330_v76_.f18, (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))], d_246_globalState_) else d_244_v16_)
                d_348_v94_: D10
                d_348_v94_ = D10_DC25(d_237_v9_)
                d_349_v95_: _dafny.Array
                nw57_ = _dafny.Array(None, 9)
                nw57_[int(0)] = d_237_v9_
                nw57_[int(1)] = d_237_v9_
                nw57_[int(2)] = d_237_v9_
                nw57_[int(3)] = (d_348_v94_).cf33
                nw57_[int(4)] = d_237_v9_
                nw57_[int(5)] = d_237_v9_
                nw57_[int(6)] = d_237_v9_
                nw57_[int(7)] = (d_241_v13_)[default__.safeIndex(d_337_v83_, len(d_241_v13_))]
                nw57_[int(8)] = (d_237_v9_ if d_255_v24_ else d_237_v9_)
                d_349_v95_ = nw57_
                index65_ = default__.safeIndex(323, (d_349_v95_).length(0))
                (d_349_v95_)[index65_] = d_237_v9_
                index66_ = default__.safeIndex(323, (d_349_v95_).length(0))
                index67_ = default__.safeIndex(465, (d_237_v9_).length(0))
                rhs66_ = d_227_v1_
                rhs67_ = d_237_v9_
                rhs68_ = d_256_v25_
                lhs53_ = d_246_globalState_
                lhs54_ = d_349_v95_
                lhs55_ = default__.safeIndex(323, (d_349_v95_).length(0))
                lhs56_ = d_237_v9_
                lhs57_ = default__.safeIndex(465, (d_237_v9_).length(0))
                lhs53_.f1 = rhs66_
                lhs54_[lhs55_] = rhs67_
                lhs56_[lhs57_] = rhs68_
                d_255_v24_ = d_336_v82_
            elif True:
                d_350_v96_: _dafny.Array
                d_351_v97_: bool
                d_352_v98_: int
                out13_: _dafny.Array
                out14_: bool
                out15_: int
                out13_, out14_, out15_ = default__.m0(d_328_cf3_, (d_256_v25_) == (len(default__.fm2(d_246_globalState_))), len(d_333_v79_), d_246_globalState_)
                d_350_v96_ = out13_
                d_351_v97_ = out14_
                d_352_v98_ = out15_
                d_255_v24_ = False
                d_353_v99_: T1
                nw58_ = C3()
                nw58_.ctor__(d_330_v76_.f18, _dafny.Map({d_240_v12_: d_256_v25_}))
                d_353_v99_ = nw58_
                d_354_v100_: C2
                nw59_ = C2()
                nw59_.ctor__(d_330_v76_.f18, d_329_v75_)
                d_354_v100_ = nw59_
                d_355_v101_: _dafny.Map
                d_355_v101_ = _dafny.Map({d_353_v99_: d_354_v100_})
                d_356_v102_: _dafny.MultiSet
                d_356_v102_ = _dafny.MultiSet([(d_355_v101_) | (d_355_v101_)])
                d_357_v103_: _dafny.Seq
                d_357_v103_ = _dafny.SeqWithoutIsStrInference([d_355_v101_])
                d_356_v102_ = (_dafny.MultiSet(d_357_v103_)) - (d_356_v102_)
                d_337_v83_ = d_226_v0_
                d_226_v0_ = d_352_v98_
            d_230_v4_ = d_230_v4_
        elif source4_.is_DC5:
            d_358___mcc_h7_ = source4_.cf6
            d_359_cf6_ = d_358___mcc_h7_
            d_236_v8_ = ((default__.fm2(d_246_globalState_) if d_228_v2_ else d_236_v8_) if (not(d_255_v24_)) and (d_228_v2_) else d_236_v8_)
            d_360_v104_: T0
            nw60_ = C0()
            nw60_.ctor__(d_240_v12_, d_231_v5_, d_230_v4_, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vf")): (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]}))
            d_360_v104_ = nw60_
            d_360_v104_ = d_360_v104_
            d_361_v105_: _dafny.Array
            nw61_ = _dafny.Array(_dafny.Seq({}), 13)
            d_361_v105_ = nw61_
            d_361_v105_ = d_361_v105_
            d_362_v106_: _dafny.MultiSet
            d_362_v106_ = _dafny.MultiSet([((d_229_v3_)[False] if (False) in (d_229_v3_) else d_228_v2_), not(d_228_v2_)])
            d_363_v107_: D1
            d_363_v107_ = D1_DC4((0) - ((d_362_v106_).cardinality), (0) - (d_256_v25_), d_228_v2_)
            d_364_v108_: _dafny.Seq
            d_364_v108_ = _dafny.SeqWithoutIsStrInference([d_236_v8_])
            index68_ = default__.safeIndex(465, (d_237_v9_).length(0))
            rhs69_ = (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]
            rhs70_ = D1_DC4(d_226_v0_, d_256_v25_, True)
            rhs71_ = d_364_v108_
            rhs72_ = default__.fm27(d_246_globalState_)
            lhs58_ = d_237_v9_
            lhs59_ = default__.safeIndex(465, (d_237_v9_).length(0))
            lhs60_ = d_246_globalState_
            lhs58_[lhs59_] = rhs69_
            d_363_v107_ = rhs70_
            d_364_v108_ = rhs71_
            lhs60_.f13 = rhs72_
        elif source4_.is_DC6:
            d_365___mcc_h8_ = source4_.cf7
            d_366___mcc_h9_ = source4_.cf8
            d_367_cf8_ = d_366___mcc_h9_
            d_368_cf7_ = d_365___mcc_h8_
            d_369_v109_: _dafny.Array
            d_370_v110_: bool
            d_371_v111_: int
            out16_: _dafny.Array
            out17_: bool
            out18_: int
            out16_, out17_, out18_ = default__.m0(d_226_v0_, (d_368_cf7_) and (d_228_v2_), d_226_v0_, d_246_globalState_)
            d_369_v109_ = out16_
            d_370_v110_ = out17_
            d_371_v111_ = out18_
            d_372_v112_: _dafny.Map
            d_372_v112_ = _dafny.Map({d_240_v12_: (0) - (len(_dafny.Map({d_256_v25_: d_368_cf7_})))})
            d_373_v113_: _dafny.Map
            d_373_v113_ = _dafny.Map({(0) - (d_256_v25_): d_372_v112_})
            d_374_v114_: C1
            nw62_ = C1()
            nw62_.ctor__((0) - (default__.safeDivisionInt(d_256_v25_, d_371_v111_)), True, d_230_v4_, ((d_373_v113_)[d_226_v0_] if (d_226_v0_) in (d_373_v113_) else d_372_v112_))
            d_374_v114_ = nw62_
            d_375_v115_: int
            out19_: int
            out19_ = (d_374_v114_).m2(not(d_367_cf8_), d_374_v114_.f22, (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))], d_371_v111_, d_246_globalState_)
            d_375_v115_ = out19_
            d_376_v116_: C0
            nw63_ = C0()
            nw63_.ctor__(_dafny.SeqWithoutIsStrInference([d_230_v4_ for d_377_i10_ in range(default__.abs(102))]), d_231_v5_, d_230_v4_, d_372_v112_)
            d_376_v116_ = nw63_
        elif True:
            d_378___mcc_h10_ = source4_.cf2
            d_379_cf2_ = d_378___mcc_h10_
            d_380_v117_: _dafny.Map
            d_380_v117_ = _dafny.Map({d_240_v12_: (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]})
            d_381_v118_: _dafny.MultiSet
            d_381_v118_ = _dafny.MultiSet([d_230_v4_])
            d_382_v119_: C0
            nw64_ = C0()
            nw64_.ctor__((d_240_v12_) + (d_240_v12_), d_231_v5_, d_230_v4_, (d_380_v117_) | ((_dafny.Map({d_240_v12_: ((d_381_v118_).set(d_230_v4_, default__.abs(d_256_v25_))).cardinality})).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtoijma")), (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))])))
            d_382_v119_ = nw64_
            d_383_v120_: D5
            d_383_v120_ = D5_DC14(d_255_v24_, (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))])
            d_384_v121_: _dafny.Map
            d_384_v121_ = _dafny.Map({_dafny.Map({d_383_v120_: d_228_v2_}): d_256_v25_})
            d_385_v122_: _dafny.Map
            d_385_v122_ = _dafny.Map({d_383_v120_: d_255_v24_})
            d_384_v121_ = _dafny.Map({d_385_v122_: ((d_311_v65_)[(d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]] if ((d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]) in (d_311_v65_) else d_256_v25_)})
            rhs73_ = d_230_v4_
            rhs74_ = ((d_229_v3_)[(d_242_v14_)[default__.safeIndex(d_256_v25_, len(d_242_v14_))]] if ((d_242_v14_)[default__.safeIndex(d_256_v25_, len(d_242_v14_))]) in (d_229_v3_) else d_255_v24_)
            d_230_v4_ = rhs73_
            d_379_cf2_ = rhs74_
            d_386_v123_: _dafny.Array
            nw65_ = _dafny.Array(None, 12)
            nw65_[int(0)] = d_240_v12_
            nw65_[int(1)] = _dafny.SeqWithoutIsStrInference([d_230_v4_ for d_387_i11_ in range(default__.abs(416))])
            nw65_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
            nw65_[int(3)] = d_240_v12_
            nw65_[int(4)] = default__.fm8(d_246_globalState_)
            nw65_[int(5)] = (d_382_v119_).f20
            nw65_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_230_v4_ for d_388_i12_ in range(default__.abs(49))]) if not(d_255_v24_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))
            nw65_[int(7)] = _dafny.SeqWithoutIsStrInference([d_230_v4_ for d_389_i13_ in range(default__.abs(466))])
            nw65_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))) + ((d_382_v119_).f20)
            nw65_[int(9)] = d_240_v12_
            nw65_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "art"))
            nw65_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhfctskoe"))
            d_386_v123_ = nw65_
            index69_ = default__.safeIndex(463, (d_386_v123_).length(0))
            (d_386_v123_)[index69_] = (d_240_v12_) + ((d_240_v12_).set(default__.safeIndex(d_226_v0_, len(d_240_v12_)), d_230_v4_))
            index70_ = default__.safeIndex(463, (d_386_v123_).length(0))
            (d_386_v123_)[index70_] = d_240_v12_
        d_390_i14_: int
        d_390_i14_ = 0
        with _dafny.label("2"):
            while False:
                with _dafny.c_label("2"):
                    if (d_390_i14_) >= (100):
                        raise _dafny.Break("2")
                    d_390_i14_ = (d_390_i14_) + (1)
                    d_391_v124_: C0
                    nw66_ = C0()
                    nw66_.ctor__(d_240_v12_, d_231_v5_, d_230_v4_, _dafny.Map({d_240_v12_: len(_dafny.Set({len(d_242_v14_)}))}))
                    d_391_v124_ = nw66_
                    d_391_v124_ = d_391_v124_
                    d_392_v125_: _dafny.Map
                    d_392_v125_ = _dafny.Map({d_240_v12_: len(d_240_v12_)})
                    d_393_v126_: C0
                    nw67_ = C0()
                    nw67_.ctor__(d_240_v12_, (d_391_v124_).f21, (d_230_v4_ if d_255_v24_ else d_230_v4_), (d_392_v125_) | (d_392_v125_))
                    d_393_v126_ = nw67_
                    d_394_v127_: _dafny.Array
                    nw68_ = _dafny.Array(_dafny.Seq({}), 22)
                    d_394_v127_ = nw68_
                    index71_ = default__.safeIndex(279, (d_394_v127_).length(0))
                    (d_394_v127_)[index71_] = (_dafny.SeqWithoutIsStrInference([d_255_v24_])) + (d_242_v14_)
                    index72_ = default__.safeIndex(279, (d_394_v127_).length(0))
                    (d_394_v127_)[index72_] = d_242_v14_
                    d_395_v128_: _dafny.Map
                    d_395_v128_ = _dafny.Map({d_237_v9_: d_392_v125_})
                    d_396_v129_: C3
                    nw69_ = C3()
                    nw69_.ctor__(d_230_v4_, (((d_395_v128_)[d_237_v9_] if (d_237_v9_) in (d_395_v128_) else d_392_v125_)) | (d_392_v125_))
                    d_396_v129_ = nw69_
                    pass
            pass
        d_397_v130_: _dafny.Map
        d_397_v130_ = _dafny.Map({d_255_v24_: d_226_v0_})
        hi2_ = default__.fm0(d_246_globalState_)
        for d_398_i15_ in range(len(d_397_v130_), hi2_):
            d_228_v2_ = (not(d_228_v2_) if (d_226_v0_) >= (d_398_i15_) else d_255_v24_)
            d_399_v131_: _dafny.Map
            d_399_v131_ = _dafny.Map({d_398_i15_: _dafny.Map({((d_397_v130_)[d_228_v2_] if (d_228_v2_) in (d_397_v130_) else d_398_i15_): len(d_236_v8_)})})
            d_399_v131_ = d_399_v131_
            d_229_v3_ = d_229_v3_
            d_237_v9_ = d_237_v9_
        d_256_v25_ = d_226_v0_
        if d_255_v24_:
            d_400_v132_: D3
            d_400_v132_ = D3_DC8(default__.fm20(d_255_v24_, d_246_globalState_))
            d_401_v134_: _dafny.Seq
            d_401_v134_ = _dafny.SeqWithoutIsStrInference([d_240_v12_])
            d_402_v135_: C3
            nw70_ = C3()
            def iife40_():
                coll28_ = _dafny.Map()
                compr_28_: _dafny.Seq
                for compr_28_ in (d_401_v134_).Elements:
                    d_403_v133_: _dafny.Seq = compr_28_
                    if (d_403_v133_) in (d_401_v134_):
                        coll28_[d_403_v133_] = d_256_v25_
                return _dafny.Map(coll28_)
            nw70_.ctor__((d_400_v132_).cf10, iife40_()
            )
            d_402_v135_ = nw70_
            d_404_v136_: _dafny.Map
            d_404_v136_ = _dafny.Map({d_402_v135_: (d_228_v2_ if d_228_v2_ else not((d_402_v135_).fm3(d_255_v24_, len(d_240_v12_), d_246_globalState_)))})
            d_405_v137_: _dafny.Seq
            d_405_v137_ = _dafny.SeqWithoutIsStrInference([default__.fm28(d_228_v2_, (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))], d_255_v24_, d_246_globalState_)])
            d_404_v136_ = (d_404_v136_).set(d_402_v135_, (((d_405_v137_)[default__.safeIndex(d_256_v25_, len(d_405_v137_))]).cardinality) != (d_226_v0_))
            (d_246_globalState_).f6 = (d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]
            d_406_v138_: _dafny.Map
            d_406_v138_ = _dafny.Map({d_240_v12_: default__.fm0(d_246_globalState_)})
            d_407_v139_: C1
            nw71_ = C1()
            nw71_.ctor__(d_256_v25_, d_228_v2_, d_230_v4_, d_406_v138_)
            d_407_v139_ = nw71_
            d_408_v140_: _dafny.Array
            nw72_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_408_v140_ = nw72_
            index73_ = default__.safeIndex(399, (d_408_v140_).length(0))
            (d_408_v140_)[index73_] = d_237_v9_
            index74_ = default__.safeIndex(399, (d_408_v140_).length(0))
            (d_408_v140_)[index74_] = d_237_v9_
            d_240_v12_ = default__.fm8(d_246_globalState_)
        elif True:
            (d_246_globalState_).f6 = d_256_v25_
            d_409_v141_: _dafny.Array
            nw73_ = _dafny.Array(None, 10)
            nw73_[int(0)] = d_237_v9_
            nw73_[int(1)] = d_237_v9_
            nw73_[int(2)] = d_237_v9_
            nw73_[int(3)] = d_237_v9_
            nw73_[int(4)] = d_237_v9_
            nw73_[int(5)] = d_237_v9_
            nw73_[int(6)] = d_237_v9_
            nw73_[int(7)] = d_237_v9_
            nw73_[int(8)] = d_237_v9_
            nw73_[int(9)] = d_237_v9_
            d_409_v141_ = nw73_
            index75_ = default__.safeIndex(146, (d_409_v141_).length(0))
            (d_409_v141_)[index75_] = d_237_v9_
            index76_ = default__.safeIndex(146, (d_409_v141_).length(0))
            (d_409_v141_)[index76_] = d_237_v9_
            d_410_v142_: D5
            d_410_v142_ = D5_DC14(False, 51)
            rhs75_ = (default__.fm29(d_246_globalState_)).cf5
            rhs76_ = (((0) - ((d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))])) - (len(d_240_v12_))) - ((d_410_v142_).cf17)
            rhs77_ = ((d_237_v9_)[default__.safeIndex(465, (d_237_v9_).length(0))]) < ((0) - (d_256_v25_))
            d_255_v24_ = rhs75_
            d_226_v0_ = rhs76_
            d_228_v2_ = rhs77_
            d_411_v143_: _dafny.MultiSet
            d_411_v143_ = _dafny.MultiSet([d_255_v24_, d_255_v24_, d_228_v2_])
            d_412_v144_: _dafny.Array
            d_413_v145_: bool
            d_414_v146_: int
            out20_: _dafny.Array
            out21_: bool
            out22_: int
            out20_, out21_, out22_ = default__.m0((d_411_v143_).cardinality, d_255_v24_, (0) - ((d_256_v25_) * (len(default__.fm30(d_228_v2_, d_246_globalState_)))), d_246_globalState_)
            d_412_v144_ = out20_
            d_413_v145_ = out21_
            d_414_v146_ = out22_
            d_413_v145_ = d_255_v24_
        _dafny.print(_dafny.string_of(d_226_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v1_) == (_dafny.SeqWithoutIsStrInference([-140]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_228_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v3_) == (_dafny.Map({False: True, True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_230_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v5_) == (_dafny.Map({_dafny.CodePoint('q'): -140}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[0]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[1]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[2]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[3]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[4]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[5]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[6]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[7]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[8]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v7_)[9]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v8_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_237_v9_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_239_v11_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_240_v12_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_241_v13_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v14_) == (_dafny.SeqWithoutIsStrInference([False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_243_v15_) == (_dafny.Map({140: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v16_) == (_dafny.Map({140: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v17_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({140: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_globalState_.f1) == (_dafny.SeqWithoutIsStrInference([-140]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_246_globalState_).f2) == (_dafny.SeqWithoutIsStrInference([-140]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_246_globalState_).f3) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_246_globalState_).f5) == (_dafny.Map({_dafny.CodePoint('q'): -140}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_246_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[0]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[1]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[2]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[3]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[4]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[5]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[6]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[7]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[8]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_246_globalState_).f7)[9]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_246_globalState_).f8)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_246_globalState_).f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_246_globalState_).f10)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_246_globalState_).f11)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_globalState_.f13) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_246_globalState_).f16) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({140: True}), _dafny.Map({140: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_252_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v23_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v23_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_255_v24_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_256_v25_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_257_v26_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False, True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v27_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_261_v28_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[0]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[0]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[1]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[1]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[2]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[2]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[3]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[3]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[4]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[4]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[5]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[5]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[6]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[6]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[7]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[7]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[8]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[8]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[9]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[9]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[10]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[10]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[11]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[11]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[12]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[12]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[13]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[13]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[14]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[14]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[15]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[15]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[16]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[16]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[17]).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_v64_)[17]).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v65_) == (_dafny.Map({-976: 458, -592: 378}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_312_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_321_v73_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v74_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v74_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_390_i14_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_397_v130_) == (_dafny.Map({False: 140}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC1(D0, NamedTuple('DC1', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC7(D2, NamedTuple('DC7', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC11(D4, NamedTuple('DC11', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC16(D6, NamedTuple('DC16', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC19(D7, NamedTuple('DC19', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)

class D8_DC21(D8, NamedTuple('DC21', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23(_dafny.Seq({}), int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)

class D9_DC23(D9, NamedTuple('DC23', [('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({self.cf27.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC26(int(0), _dafny.Array(None, 0), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)

class D10_DC26(D10, NamedTuple('DC26', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC25(D10, NamedTuple('DC25', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC28(None, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)

class D11_DC28(D11, NamedTuple('DC28', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC30(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)

class D12_DC30(D12, NamedTuple('DC30', [('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf46)}, {self.cf47.VerbatimString(True)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC29(D12, NamedTuple('DC29', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)

class D13_DC33(D13, NamedTuple('DC33', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value

class T1:
    pass
    def m1(self, p0, globalState):
        pass

    def m2(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: _dafny.Seq = _dafny.Seq({})
        self.f6: int = int(0)
        self.f13: _dafny.Seq = _dafny.Seq({})
        self._f0: int = int(0)
        self._f2: _dafny.Seq = _dafny.Seq({})
        self._f3: _dafny.Map = _dafny.Map({})
        self._f4: int = int(0)
        self._f5: _dafny.Map = _dafny.Map({})
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        self._f8: _dafny.Map = _dafny.Map({})
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Map = _dafny.Map({})
        self._f11: _dafny.Seq = _dafny.Seq({})
        self._f12: bool = False
        self._f14: bool = False
        self._f15: int = int(0)
        self._f16: _dafny.Seq = _dafny.Seq({})
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17

    @property
    def f0(self):
        return self._f0
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
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17

class C0(T0):
    def  __init__(self):
        self._f18: str = _dafny.CodePoint('D')
        self._f19: _dafny.Map = _dafny.Map({})
        self._f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f21: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value
    @property
    def f19(self):
        return self._f19
    def ctor__(self, f20, f21, f18, f19):
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f18 = f18
        (self)._f19 = f19

    def fm6(self, p0, p1, globalState):
        return ((len(_dafny.Map({False: False}))) - (len((self).f20))) != ((-574) + (924))

    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21

class C1(T1, T0):
    def  __init__(self):
        self._f18: str = _dafny.CodePoint('D')
        self._f19: _dafny.Map = _dafny.Map({})
        self.f22: int = int(0)
        self._f23: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value
    @property
    def f19(self):
        return self._f19
    def ctor__(self, f22, f23, f18, f19):
        (self).f22 = f22
        (self)._f23 = f23
        (self).f18 = f18
        (self)._f19 = f19

    def fm3(self, p0, p1, globalState):
        return (self).f23

    def fm4(self, p0, p1, p2, p3, globalState):
        return (390) * ((0) - ((0) - ((851 if not((self).f23) else 366))))

    def fm12(self, globalState):
        return (self).f23

    def fm13(self, p0, p1, p2, p3, globalState):
        if not(not((self).f23)):
            return self.f18
        elif True:
            return _dafny.CodePoint('b')

    def m1(self, p0, globalState):
        r0: D0 = D0.default()()
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_415_v0_: _dafny.Map
        d_415_v0_ = _dafny.Map({(self).f23: self.f22})
        d_416_v1_: _dafny.Set
        d_416_v1_ = _dafny.Set({d_415_v0_, d_415_v0_})
        r3 = ((_dafny.Set({d_415_v0_})) | (d_416_v1_)).issubset(d_416_v1_)
        (globalState).f6 = (self.f22) + (self.f22)
        d_417_v2_: _dafny.MultiSet
        d_417_v2_ = _dafny.MultiSet([self.f22, self.f22, self.f22])
        d_418_i0_: int
        d_418_i0_ = 0
        with _dafny.label("3"):
            while ((d_417_v2_).set(self.f22, default__.abs(self.f22))).issubset((d_417_v2_ if p0 else d_417_v2_)):
                with _dafny.c_label("3"):
                    if (d_418_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_418_i0_ = (d_418_i0_) + (1)
                    d_419_v3_: D3
                    d_419_v3_ = D3_DC10(p0)
                    source5_ = d_419_v3_
                    if source5_.is_DC9:
                        d_420___mcc_h0_ = source5_.cf11
                        d_421___mcc_h1_ = source5_.cf12
                        d_422_cf12_ = d_421___mcc_h1_
                        d_423_cf11_ = d_420___mcc_h0_
                        d_424_v4_: _dafny.Seq
                        d_424_v4_ = _dafny.SeqWithoutIsStrInference([not(d_423_cf11_)])
                        d_425_v5_: _dafny.MultiSet
                        d_425_v5_ = _dafny.MultiSet([(self).f23, (self).f23, p0])
                        d_426_v6_: _dafny.Map
                        d_426_v6_ = _dafny.Map({d_422_cf12_: (d_425_v5_).cardinality})
                        d_427_v7_: D3
                        d_427_v7_ = D3_DC8(self.f18)
                        d_428_v8_: _dafny.Seq
                        d_428_v8_ = _dafny.SeqWithoutIsStrInference([self.f22])
                        d_429_v9_: _dafny.Set
                        d_429_v9_ = _dafny.Set({d_423_cf11_})
                        d_430_v10_: _dafny.Set
                        d_430_v10_ = _dafny.Set({d_422_cf12_})
                        d_431_v11_: _dafny.Map
                        d_431_v11_ = _dafny.Map({d_423_cf11_: d_428_v8_})
                        d_432_v12_: _dafny.Array
                        nw74_ = _dafny.Array(None, 21)
                        nw74_[int(0)] = p0
                        nw74_[int(1)] = default__.fm1(p0, self.f18, self.f22, globalState)
                        nw74_[int(2)] = d_423_cf11_
                        nw74_[int(3)] = False
                        nw74_[int(4)] = d_423_cf11_
                        nw74_[int(5)] = not((self).f23)
                        nw74_[int(6)] = (self).f23
                        nw74_[int(7)] = default__.fm1((self).f23, self.f18, self.f22, globalState)
                        nw74_[int(8)] = (d_424_v4_) < (_dafny.SeqWithoutIsStrInference([(self).f23, p0, p0, d_423_cf11_]))
                        nw74_[int(9)] = (d_422_cf12_) <= (len(d_426_v6_))
                        nw74_[int(10)] = (self.f22) != (self.f22)
                        nw74_[int(11)] = d_423_cf11_
                        nw74_[int(12)] = default__.fm1((self).f23, (d_427_v7_).cf10, (d_428_v8_)[default__.safeIndex(d_422_cf12_, len(d_428_v8_))], globalState)
                        nw74_[int(13)] = (default__.fm2(globalState)).issubset(d_429_v9_)
                        nw74_[int(14)] = not(False)
                        nw74_[int(15)] = p0
                        nw74_[int(16)] = (len(d_430_v10_)) in (((d_431_v11_)[(self).f23] if ((self).f23) in (d_431_v11_) else d_428_v8_))
                        nw74_[int(17)] = (d_422_cf12_) == (d_422_cf12_)
                        nw74_[int(18)] = (self).f23
                        nw74_[int(19)] = d_423_cf11_
                        nw74_[int(20)] = default__.fm1(False, self.f18, -279, globalState)
                        d_432_v12_ = nw74_
                        d_432_v12_ = d_432_v12_
                        d_433_v13_: D0
                        d_433_v13_ = D0_DC2(D0_DC1())
                        d_434_v14_: D0
                        d_434_v14_ = D0_DC2(d_433_v13_)
                        d_434_v14_ = d_434_v14_
                        d_435_v15_: _dafny.Seq
                        d_435_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkdt"))
                        d_436_v16_: _dafny.Set
                        d_436_v16_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylwbg")), d_435_v15_})
                        d_436_v16_ = (d_436_v16_) | (_dafny.Set({d_435_v15_, d_435_v15_, d_435_v15_}))
                        d_437_v17_: _dafny.Map
                        d_437_v17_ = _dafny.Map({self.f18: self.f22})
                        d_438_v19_: _dafny.Seq
                        d_438_v19_ = _dafny.SeqWithoutIsStrInference([d_435_v15_, d_435_v15_, d_435_v15_])
                        d_439_v20_: _dafny.Seq
                        def iife41_():
                            coll29_ = _dafny.Map()
                            compr_29_: _dafny.Seq
                            for compr_29_ in (d_438_v19_).Elements:
                                d_441_v18_: _dafny.Seq = compr_29_
                                if (d_441_v18_) in (d_438_v19_):
                                    coll29_[d_441_v18_] = self.f22
                            return _dafny.Map(coll29_)
                        d_439_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f18 for d_440_i1_ in range(default__.abs(170))]): d_422_cf12_}), (self).f19, iife41_()
                        , (self).f19])
                        d_442_v21_: C0
                        nw75_ = C0()
                        nw75_.ctor__(d_435_v15_, d_437_v17_, _dafny.CodePoint('x'), (d_439_v20_)[default__.safeIndex(self.f22, len(d_439_v20_))])
                        d_442_v21_ = nw75_
                    elif source5_.is_DC10:
                        d_443___mcc_h2_ = source5_.cf13
                        d_444_cf13_ = d_443___mcc_h2_
                        d_444_cf13_ = ((D3_DC9(d_444_cf13_, default__.fm0(globalState))).cf11) == (d_444_cf13_)
                        d_445_v22_: _dafny.Array
                        nw76_ = _dafny.Array(False, 13)
                        d_445_v22_ = nw76_
                        d_446_v23_: _dafny.Map
                        d_446_v23_ = _dafny.Map({d_445_v22_: self.f22})
                        d_447_v24_: _dafny.MultiSet
                        d_447_v24_ = _dafny.MultiSet([(self).f23])
                        d_448_v25_: _dafny.Set
                        d_448_v25_ = _dafny.Set({False})
                        d_449_v26_: _dafny.Array
                        nw77_ = _dafny.Array(None, 8)
                        nw77_[int(0)] = (self.f22) + (self.f22)
                        nw77_[int(1)] = (((d_446_v23_)[d_445_v22_] if (d_445_v22_) in (d_446_v23_) else (d_447_v24_).cardinality)) * (self.f22)
                        nw77_[int(2)] = self.f22
                        nw77_[int(3)] = self.f22
                        nw77_[int(4)] = self.f22
                        nw77_[int(5)] = len(d_448_v25_)
                        nw77_[int(6)] = (0) - (self.f22)
                        nw77_[int(7)] = (251) - (self.f22)
                        d_449_v26_ = nw77_
                        index77_ = default__.safeIndex(392, (d_449_v26_).length(0))
                        (d_449_v26_)[index77_] = default__.safeDivisionInt(self.f22, len(default__.fm14(default__.fm0(globalState), globalState)))
                        d_450_v27_: _dafny.Seq
                        d_450_v27_ = _dafny.SeqWithoutIsStrInference([default__.fm0(globalState)])
                        index78_ = default__.safeIndex(392, (d_449_v26_).length(0))
                        (d_449_v26_)[index78_] = default__.safeDivisionInt((_dafny.MultiSet(d_450_v27_)).cardinality, self.f22)
                        d_451_v28_: _dafny.Array
                        nw78_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                        d_451_v28_ = nw78_
                        index79_ = default__.safeIndex(156, (d_451_v28_).length(0))
                        (d_451_v28_)[index79_] = _dafny.SeqWithoutIsStrInference([self.f18 for d_452_i2_ in range(default__.abs(132))])
                        d_453_v29_: _dafny.Seq
                        d_453_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vphuxrsd"))
                        index80_ = default__.safeIndex(156, (d_451_v28_).length(0))
                        (d_451_v28_)[index80_] = (d_453_v29_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))
                        index81_ = default__.safeIndex(392, (d_449_v26_).length(0))
                        (d_449_v26_)[index81_] = (249) * ((0) - (self.f22))
                    elif True:
                        d_454___mcc_h3_ = source5_.cf10
                        d_455_cf10_ = d_454___mcc_h3_
                        d_456_v30_: _dafny.MultiSet
                        d_456_v30_ = _dafny.MultiSet([(self).f23, p0, p0, (self).f23, (self).f23])
                        d_457_v31_: D5
                        d_457_v31_ = D5_DC12(d_456_v30_)
                        d_458_v32_: _dafny.Map
                        d_458_v32_ = _dafny.Map({self.f22: (d_457_v31_).cf15})
                        d_459_v33_: _dafny.Seq
                        d_459_v33_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_460_v34_: _dafny.Array
                        nw79_ = _dafny.Array(None, 17)
                        nw79_[int(0)] = d_456_v30_
                        nw79_[int(1)] = d_456_v30_
                        nw79_[int(2)] = ((d_458_v32_)[(_dafny.MultiSet([p0])).cardinality] if ((_dafny.MultiSet([p0])).cardinality) in (d_458_v32_) else _dafny.MultiSet([(self).f23]))
                        nw79_[int(3)] = d_456_v30_
                        nw79_[int(4)] = d_456_v30_
                        nw79_[int(5)] = d_456_v30_
                        nw79_[int(6)] = d_456_v30_
                        nw79_[int(7)] = (d_456_v30_).set((self).f23, default__.abs(self.f22))
                        nw79_[int(8)] = _dafny.MultiSet(d_459_v33_)
                        nw79_[int(9)] = (d_456_v30_).intersection(d_456_v30_)
                        nw79_[int(10)] = d_456_v30_
                        nw79_[int(11)] = d_456_v30_
                        nw79_[int(12)] = d_456_v30_
                        nw79_[int(13)] = d_456_v30_
                        nw79_[int(14)] = d_456_v30_
                        nw79_[int(15)] = d_456_v30_
                        nw79_[int(16)] = (d_456_v30_) - (_dafny.MultiSet(d_459_v33_))
                        d_460_v34_ = nw79_
                        index82_ = default__.safeIndex(184, (d_460_v34_).length(0))
                        (d_460_v34_)[index82_] = (d_456_v30_) | (d_456_v30_)
                        d_461_v35_: _dafny.Seq
                        d_461_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hoka"))
                        d_462_v36_: _dafny.MultiSet
                        d_462_v36_ = _dafny.MultiSet([(d_461_v35_ if p0 else d_461_v35_), (d_461_v35_) + (d_461_v35_), _dafny.SeqWithoutIsStrInference([self.f18 for d_463_i3_ in range(default__.abs(874))]), d_461_v35_, (d_461_v35_).set(default__.safeIndex(self.f22, len(d_461_v35_)), d_455_cf10_)])
                        d_464_v37_: _dafny.Map
                        d_464_v37_ = _dafny.Map({self.f22: (d_456_v30_).cardinality})
                        index83_ = default__.safeIndex(184, (d_460_v34_).length(0))
                        rhs78_ = p0
                        rhs79_ = (d_456_v30_) | (d_456_v30_)
                        rhs80_ = default__.fm15(self.f18, (0) - (self.f22), d_419_v3_, globalState)
                        rhs81_ = (0) - (((self.f22) - (self.f22)) * (((d_464_v37_)[len(d_415_v0_)] if (len(d_415_v0_)) in (d_464_v37_) else self.f22)))
                        lhs61_ = d_460_v34_
                        lhs62_ = default__.safeIndex(184, (d_460_v34_).length(0))
                        lhs63_ = globalState
                        r3 = rhs78_
                        lhs61_[lhs62_] = rhs79_
                        d_462_v36_ = rhs80_
                        lhs63_.f6 = rhs81_
                        d_465_v38_: _dafny.Map
                        d_465_v38_ = _dafny.Map({p0: (self).f23})
                        d_466_v39_: _dafny.Set
                        d_466_v39_ = _dafny.Set({self.f22, self.f22, 295})
                        d_467_v40_: _dafny.Seq
                        d_467_v40_ = _dafny.SeqWithoutIsStrInference([d_466_v39_])
                        d_465_v38_ = (d_465_v38_).set((d_466_v39_).isdisjoint((d_467_v40_)[default__.safeIndex(-737, len(d_467_v40_))]), p0)
                        d_468_v41_: _dafny.Array
                        def lambda14_(d_469_i4_):
                            return (d_469_i4_) * (self.f22)

                        init7_ = lambda14_
                        nw80_ = _dafny.Array(None, 21)
                        for i0_7_ in range(nw80_.length(0)):
                            nw80_[i0_7_] = init7_(i0_7_)
                        d_468_v41_ = nw80_
                        index84_ = default__.safeIndex(162, (d_468_v41_).length(0))
                        (d_468_v41_)[index84_] = self.f22
                        index85_ = default__.safeIndex(162, (d_468_v41_).length(0))
                        (d_468_v41_)[index85_] = (0) - (self.f22)
                        d_470_v42_: D5
                        d_470_v42_ = D5_DC14((p0 if p0 else (self).f23), (d_468_v41_)[default__.safeIndex(162, (d_468_v41_).length(0))])
                        d_471_v43_: _dafny.Seq
                        d_471_v43_ = _dafny.SeqWithoutIsStrInference([len(d_461_v35_)])
                        pat_let_tv4_ = d_468_v41_
                        pat_let_tv5_ = d_468_v41_
                        pat_let_tv6_ = d_468_v41_
                        pat_let_tv7_ = d_468_v41_
                        pat_let_tv8_ = d_471_v43_
                        pat_let_tv9_ = d_471_v43_
                        def iife44_(_pat_let8_0):
                            def iife45_(d_472_dt__update__tmp_h0_):
                                def iife46_(_pat_let9_0):
                                    def iife47_(d_473_dt__update_hcf17_h0_):
                                        return D5_DC14((d_472_dt__update__tmp_h0_).cf16, d_473_dt__update_hcf17_h0_)
                                    return iife47_(_pat_let9_0)
                                return iife46_((pat_let_tv5_)[default__.safeIndex(162, (pat_let_tv4_).length(0))])
                            return iife45_(_pat_let8_0)
                        def iife43_(_pat_let7_0):
                            def iife48_(d_474_dt__update__tmp_h1_):
                                def iife49_(_pat_let10_0):
                                    def iife50_(d_475_dt__update_hcf17_h1_):
                                        return D5_DC14((d_474_dt__update__tmp_h1_).cf16, d_475_dt__update_hcf17_h1_)
                                    return iife50_(_pat_let10_0)
                                return iife49_((pat_let_tv7_)[default__.safeIndex(162, (pat_let_tv6_).length(0))])
                            return iife48_(_pat_let7_0)
                        def iife42_(_pat_let6_0):
                            def iife51_(d_476_dt__update__tmp_h2_):
                                def iife52_(_pat_let11_0):
                                    def iife53_(d_477_dt__update_hcf17_h2_):
                                        return D5_DC14((d_476_dt__update__tmp_h2_).cf16, d_477_dt__update_hcf17_h2_)
                                    return iife53_(_pat_let11_0)
                                return iife52_(len((pat_let_tv8_) + (pat_let_tv9_)))
                            return iife51_(_pat_let6_0)
                        d_470_v42_ = iife42_(iife43_(iife44_(d_470_v42_)))
                    d_478_v44_: _dafny.Seq
                    d_478_v44_ = _dafny.SeqWithoutIsStrInference([(self).f23])
                    d_479_v45_: _dafny.Map
                    d_479_v45_ = _dafny.Map({self.f22: self.f22})
                    d_480_v46_: _dafny.Seq
                    d_480_v46_ = _dafny.SeqWithoutIsStrInference([d_479_v45_, d_479_v45_])
                    d_481_v47_: _dafny.Seq
                    d_481_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gr"))
                    d_482_v48_: _dafny.Seq
                    d_482_v48_ = d_478_v44_
                    d_483_v49_: _dafny.Array
                    nw81_ = _dafny.Array(None, 15)
                    nw81_[int(0)] = False
                    nw81_[int(1)] = (self).f23
                    nw81_[int(2)] = (d_478_v44_)[default__.safeIndex(446, len(d_478_v44_))]
                    nw81_[int(3)] = (self).f23
                    nw81_[int(4)] = (d_480_v46_) < (default__.fm16(globalState))
                    nw81_[int(5)] = (self).f23
                    nw81_[int(6)] = p0
                    nw81_[int(7)] = (self.f18) in (d_481_v47_)
                    nw81_[int(8)] = ((d_482_v48_)) == (d_478_v44_)
                    nw81_[int(9)] = p0
                    nw81_[int(10)] = p0
                    nw81_[int(11)] = not ((self).f23) or (p0)
                    nw81_[int(12)] = not(p0)
                    nw81_[int(13)] = default__.fm1(p0, self.f18, self.f22, globalState)
                    nw81_[int(14)] = p0
                    d_483_v49_ = nw81_
                    index86_ = default__.safeIndex(732, (d_483_v49_).length(0))
                    (d_483_v49_)[index86_] = p0
                    index87_ = default__.safeIndex(732, (d_483_v49_).length(0))
                    (d_483_v49_)[index87_] = default__.fm1(p0, (self.f18 if not(p0) else self.f18), default__.fm0(globalState), globalState)
                    d_484_v50_: C0
                    nw82_ = C0()
                    nw82_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))) + (d_481_v47_), _dafny.Map({_dafny.CodePoint('d'): self.f22}), self.f18, ((self).f19) | ((self).f19))
                    d_484_v50_ = nw82_
                    d_485_v51_: _dafny.Array
                    nw83_ = _dafny.Array(int(0), 12)
                    d_485_v51_ = nw83_
                    index88_ = default__.safeIndex(194, (d_485_v51_).length(0))
                    (d_485_v51_)[index88_] = len((d_484_v50_).f20)
                    index89_ = default__.safeIndex(194, (d_485_v51_).length(0))
                    (d_485_v51_)[index89_] = self.f22
                    pass
            pass
        hi3_ = self.f22
        for d_486_i5_ in range(self.f22, hi3_):
            d_487_v52_: _dafny.Seq
            d_487_v52_ = _dafny.SeqWithoutIsStrInference([d_486_i5_, -684])
            d_488_v53_: _dafny.Seq
            d_488_v53_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_486_i5_]), d_487_v52_])
            (globalState).f6 = (d_486_i5_) * ((len(d_488_v53_)) * (d_486_i5_))
            d_489_v54_: _dafny.Seq
            d_489_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isnmbr"))
            rhs82_ = (self).f23
            rhs83_ = (0) - (len(d_489_v54_))
            r2 = rhs82_
            r1 = rhs83_
            r3 = p0
            r2 = not(p0)
        d_490_v55_: _dafny.Seq
        d_490_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk"))
        d_491_i6_: int
        d_491_i6_ = 0
        with _dafny.label("4"):
            while ((d_490_v55_) != (d_490_v55_) if (self).f23 else (self).f23):
                with _dafny.c_label("4"):
                    if (d_491_i6_) >= (100):
                        raise _dafny.Break("4")
                    d_491_i6_ = (d_491_i6_) + (1)
                    d_492_v56_: _dafny.Map
                    d_492_v56_ = _dafny.Map({self.f18: self.f22})
                    d_493_v57_: C0
                    nw84_ = C0()
                    nw84_.ctor__(d_490_v55_, d_492_v56_, (self.f18 if (self).f23 else self.f18), ((self).f19) | ((self).f19))
                    d_493_v57_ = nw84_
                    d_494_v58_: _dafny.Seq
                    d_494_v58_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_495_v59_: _dafny.Map
                    d_495_v59_ = _dafny.Map({self.f22: d_494_v58_})
                    d_496_v60_: _dafny.Seq
                    d_496_v60_ = (((d_495_v59_)[self.f22] if (self.f22) in (d_495_v59_) else _dafny.SeqWithoutIsStrInference([(self).f23]))) + (d_494_v58_)
                    d_497_v61_: _dafny.Set
                    d_497_v61_ = _dafny.Set({p0, default__.fm1(not(not(p0)), self.f18, 169, globalState), (self).f23})
                    d_498_v62_: _dafny.Map
                    d_498_v62_ = _dafny.Map({self.f22: (self).f23})
                    d_499_v63_: _dafny.Array
                    nw85_ = _dafny.Array(None, 4)
                    nw85_[int(0)] = True
                    nw85_[int(1)] = (d_497_v61_).isdisjoint(_dafny.Set({(self).f23}))
                    nw85_[int(2)] = (self).f23
                    nw85_[int(3)] = ((d_498_v62_)[self.f22] if (self.f22) in (d_498_v62_) else (self).f23)
                    d_499_v63_ = nw85_
                    index90_ = default__.safeIndex(342, (d_499_v63_).length(0))
                    (d_499_v63_)[index90_] = p0
                    d_500_v64_: _dafny.Seq
                    d_500_v64_ = _dafny.SeqWithoutIsStrInference([d_498_v62_, d_498_v62_])
                    index91_ = default__.safeIndex(342, (d_499_v63_).length(0))
                    rhs84_ = self.f22
                    rhs85_ = d_496_v60_
                    rhs86_ = True
                    rhs87_ = self.f22
                    rhs88_ = (d_498_v62_) == ((d_500_v64_)[default__.safeIndex(len((d_493_v57_).f20), len(d_500_v64_))])
                    lhs64_ = globalState
                    lhs65_ = globalState
                    lhs66_ = d_499_v63_
                    lhs67_ = default__.safeIndex(342, (d_499_v63_).length(0))
                    lhs64_.f6 = rhs84_
                    d_496_v60_ = rhs85_
                    r3 = rhs86_
                    lhs65_.f6 = rhs87_
                    lhs66_[lhs67_] = rhs88_
                    index92_ = default__.safeIndex(342, (d_499_v63_).length(0))
                    (d_499_v63_)[index92_] = (d_494_v58_)[default__.safeIndex(386, len(d_494_v58_))]
                    d_501_v65_: D3
                    d_501_v65_ = D3_DC8((self.f18 if False else self.f18))
                    source6_ = d_501_v65_
                    if source6_.is_DC9:
                        d_502___mcc_h4_ = source6_.cf11
                        d_503___mcc_h5_ = source6_.cf12
                        d_504_cf12_ = d_503___mcc_h5_
                        d_505_cf11_ = d_502___mcc_h4_
                        d_506_v66_: C0
                        nw86_ = C0()
                        nw86_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isrngtcpi")), ((d_493_v57_).f21 if d_505_cf11_ else _dafny.Map({self.f18: len(_dafny.SeqWithoutIsStrInference([self.f22, d_504_cf12_]))})), self.f18, (self).f19)
                        d_506_v66_ = nw86_
                        d_507_v67_: _dafny.Array
                        nw87_ = _dafny.Array(int(0), 26)
                        d_507_v67_ = nw87_
                        index93_ = default__.safeIndex(341, (d_507_v67_).length(0))
                        (d_507_v67_)[index93_] = (self.f22) + (d_504_cf12_)
                        index94_ = default__.safeIndex(341, (d_507_v67_).length(0))
                        (d_507_v67_)[index94_] = d_504_cf12_
                        d_508_v68_: _dafny.Array
                        nw88_ = _dafny.Array(None, 11)
                        nw88_[int(0)] = self.f18
                        nw88_[int(1)] = _dafny.CodePoint('w')
                        nw88_[int(2)] = (self).fm13(d_504_cf12_, self.f22, (_dafny.MultiSet([not(p0)])).cardinality, (self).f23, globalState)
                        nw88_[int(3)] = self.f18
                        nw88_[int(4)] = self.f18
                        nw88_[int(5)] = (self.f18 if d_505_cf11_ else self.f18)
                        nw88_[int(6)] = (self).fm13((d_507_v67_)[default__.safeIndex(341, (d_507_v67_).length(0))], d_504_cf12_, self.f22, not(False), globalState)
                        nw88_[int(7)] = self.f18
                        nw88_[int(8)] = self.f18
                        nw88_[int(9)] = self.f18
                        nw88_[int(10)] = self.f18
                        d_508_v68_ = nw88_
                        d_508_v68_ = d_508_v68_
                        r3 = d_505_cf11_
                    elif source6_.is_DC10:
                        d_509___mcc_h6_ = source6_.cf13
                        d_510_cf13_ = d_509___mcc_h6_
                        d_511_v69_: D5
                        d_511_v69_ = D5_DC14(((d_499_v63_)[default__.safeIndex(342, (d_499_v63_).length(0))]) == (d_510_cf13_), (self.f22) - (self.f22))
                        d_511_v69_ = d_511_v69_
                        r1 = (0) - (self.f22)
                        d_512_v70_: _dafny.Array
                        nw89_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                        d_512_v70_ = nw89_
                        index95_ = default__.safeIndex(367, (d_512_v70_).length(0))
                        (d_512_v70_)[index95_] = (d_493_v57_).f20
                        index96_ = default__.safeIndex(367, (d_512_v70_).length(0))
                        (d_512_v70_)[index96_] = (d_493_v57_).f20
                        d_513_v71_: _dafny.Array
                        nw90_ = _dafny.Array(_dafny.Seq({}), 14)
                        d_513_v71_ = nw90_
                        def lambda15_(d_514_v60_):
                            def lambda16_(d_515_i7_):
                                return d_514_v60_

                            return lambda16_

                        init8_ = lambda15_(d_496_v60_)
                        nw91_ = _dafny.Array(None, 9)
                        for i0_8_ in range(nw91_.length(0)):
                            nw91_[i0_8_] = init8_(i0_8_)
                        d_513_v71_ = nw91_
                    elif True:
                        d_516___mcc_h7_ = source6_.cf10
                        d_517_cf10_ = d_516___mcc_h7_
                        d_517_cf10_ = d_517_cf10_
                        (globalState).f6 = self.f22
                        d_518_v72_: _dafny.Set
                        d_518_v72_ = _dafny.Set({self.f22, 513})
                        d_519_v73_: _dafny.Seq
                        d_519_v73_ = _dafny.SeqWithoutIsStrInference([len(d_498_v62_), self.f22])
                        d_518_v72_ = (_dafny.Set({(d_519_v73_)[default__.safeIndex(self.f22, len(d_519_v73_))]})).intersection(d_518_v72_)
                        d_520_v74_: _dafny.Seq
                        d_520_v74_ = _dafny.SeqWithoutIsStrInference([d_417_v2_])
                        d_521_v75_: _dafny.Set
                        d_521_v75_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([836, self.f22])) + (_dafny.SeqWithoutIsStrInference([self.f22 for d_522_i8_ in range(default__.abs(264))])), _dafny.SeqWithoutIsStrInference([len(d_520_v74_), len(d_519_v73_)]), d_519_v73_})
                        def iife54_():
                            coll30_ = _dafny.Set()
                            compr_30_: _dafny.Seq
                            for compr_30_ in (_dafny.SeqWithoutIsStrInference([d_519_v73_])).Elements:
                                d_523_v76_: _dafny.Seq = compr_30_
                                if (d_523_v76_) in (_dafny.SeqWithoutIsStrInference([d_519_v73_])):
                                    coll30_ = coll30_.union(_dafny.Set([d_523_v76_]))
                            return _dafny.Set(coll30_)
                        d_521_v75_ = (iife54_()
                        ) | (d_521_v75_)
                    pass
            pass
        d_524_v77_: _dafny.Array
        nw92_ = _dafny.Array(int(0), 2)
        d_524_v77_ = nw92_
        d_525_v78_: D3
        d_525_v78_ = D3_DC8(_dafny.CodePoint('r'))
        d_526_v79_: _dafny.MultiSet
        d_526_v79_ = _dafny.MultiSet([d_490_v55_, _dafny.SeqWithoutIsStrInference([self.f18 for d_527_i9_ in range(default__.abs(-325))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdamu"))])
        d_528_v80_: _dafny.Array
        nw93_ = _dafny.Array(False, 13)
        d_528_v80_ = nw93_
        d_529_v81_: _dafny.Seq
        d_529_v81_ = _dafny.SeqWithoutIsStrInference([self.f22, self.f22, self.f22, self.f22, self.f22])
        d_530_v82_: _dafny.Map
        d_530_v82_ = _dafny.Map({d_528_v80_: (d_529_v81_)[default__.safeIndex(self.f22, len(d_529_v81_))]})
        d_531_v83_: _dafny.Seq
        d_531_v83_ = _dafny.SeqWithoutIsStrInference([(self).f23])
        rhs89_ = d_524_v77_
        rhs90_ = d_525_v78_
        rhs91_ = d_526_v79_
        rhs92_ = _dafny.Map({d_528_v80_: self.f22})
        rhs93_ = d_531_v83_
        lhs68_ = globalState
        d_524_v77_ = rhs89_
        d_525_v78_ = rhs90_
        d_526_v79_ = rhs91_
        d_530_v82_ = rhs92_
        lhs68_.f13 = rhs93_
        r0 = D0_DC0(self.f22)
        r1 = self.f22
        r2 = not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjptbhmhv"))) != (_dafny.SeqWithoutIsStrInference([self.f18 for d_532_i10_ in range(default__.abs(423))])))
        r3 = p0
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_533_v0_: _dafny.Array
        nw94_ = _dafny.Array(None, 15)
        nw94_[int(0)] = self.f18
        nw94_[int(1)] = _dafny.CodePoint('q')
        nw94_[int(2)] = self.f18
        nw94_[int(3)] = self.f18
        nw94_[int(4)] = (self).fm13(p3, p2, 968, p0, globalState)
        nw94_[int(5)] = self.f18
        nw94_[int(6)] = self.f18
        nw94_[int(7)] = self.f18
        nw94_[int(8)] = self.f18
        nw94_[int(9)] = self.f18
        nw94_[int(10)] = self.f18
        nw94_[int(11)] = self.f18
        nw94_[int(12)] = self.f18
        nw94_[int(13)] = self.f18
        nw94_[int(14)] = self.f18
        d_533_v0_ = nw94_
        d_534_v1_: _dafny.Set
        d_534_v1_ = _dafny.Set({d_533_v0_, d_533_v0_})
        (globalState).f6 = (self).fm4(not((self).f23), (self.f18 if (self).f23 else self.f18), self.f18, (d_533_v0_) in (d_534_v1_), globalState)
        d_535_v2_: _dafny.Seq
        d_535_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdgvd"))
        d_536_v3_: _dafny.Map
        d_536_v3_ = _dafny.Map({((d_535_v2_) + (d_535_v2_)).set(default__.safeIndex(p2, len((d_535_v2_) + (d_535_v2_))), self.f18): not(default__.fm1((self).f23, self.f18, len(d_535_v2_), globalState))})
        d_537_v4_: _dafny.Map
        d_537_v4_ = _dafny.Map({self.f22: d_535_v2_})
        d_536_v3_ = (d_536_v3_).set((((d_537_v4_)[p1] if (p1) in (d_537_v4_) else d_535_v2_)) + (d_535_v2_), False)
        d_538_i0_: int
        d_538_i0_ = 0
        with _dafny.label("5"):
            while (p1) >= (p2):
                with _dafny.c_label("5"):
                    if (d_538_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_538_i0_ = (d_538_i0_) + (1)
                    (globalState).f6 = (p2) - (p1)
                    d_539_v5_: _dafny.Map
                    d_539_v5_ = _dafny.Map({p3: self.f18})
                    d_540_v6_: D1
                    d_540_v6_ = D1_DC4(p1, p2, (self).f23)
                    r0 = (self).fm4((self).fm12(globalState), self.f18, ((d_539_v5_)[(d_540_v6_).cf4] if ((d_540_v6_).cf4) in (d_539_v5_) else _dafny.CodePoint('y')), (self).f23, globalState)
                    (self).f22 = p1
                    d_541_v7_: _dafny.Array
                    nw95_ = _dafny.Array(int(0), 7)
                    d_541_v7_ = nw95_
                    index97_ = default__.safeIndex(352, (d_541_v7_).length(0))
                    (d_541_v7_)[index97_] = default__.safeModuloInt(p2, len(d_539_v5_))
                    index98_ = default__.safeIndex(352, (d_541_v7_).length(0))
                    (d_541_v7_)[index98_] = 365
                    pass
            pass
        d_542_v8_: _dafny.Seq
        d_542_v8_ = _dafny.SeqWithoutIsStrInference([p0, False])
        if (p1) >= (default__.safeModuloInt(p1, len(d_542_v8_))):
            d_543_v9_: _dafny.MultiSet
            d_543_v9_ = _dafny.MultiSet([False, (self).f23])
            d_543_v9_ = (d_543_v9_) | (d_543_v9_)
            d_542_v8_ = (d_542_v8_) + (d_542_v8_)
            d_544_v10_: _dafny.Set
            d_544_v10_ = _dafny.Set({299})
            d_544_v10_ = default__.fm17(globalState)
            d_545_v11_: _dafny.Map
            d_545_v11_ = _dafny.Map({self.f22: p3})
            d_546_v12_: _dafny.Map
            d_546_v12_ = _dafny.Map({_dafny.CodePoint('s'): len(d_545_v11_)})
            d_547_v14_: _dafny.Seq
            d_547_v14_ = _dafny.SeqWithoutIsStrInference([d_535_v2_, d_535_v2_])
            d_548_v15_: C0
            nw96_ = C0()
            def iife55_():
                coll31_ = _dafny.Map()
                compr_31_: _dafny.Seq
                for compr_31_ in ((d_547_v14_).set(default__.safeIndex(p3, len(d_547_v14_)), d_535_v2_)).Elements:
                    d_549_v13_: _dafny.Seq = compr_31_
                    if (d_549_v13_) in ((d_547_v14_).set(default__.safeIndex(p3, len(d_547_v14_)), d_535_v2_)):
                        coll31_[d_549_v13_] = (_dafny.MultiSet([self.f18, self.f18, _dafny.CodePoint('f'), _dafny.CodePoint('t'), self.f18])).cardinality
                return _dafny.Map(coll31_)
            nw96_.ctor__(d_535_v2_, d_546_v12_, self.f18, iife55_()
            )
            d_548_v15_ = nw96_
            d_550_v16_: _dafny.Map
            d_550_v16_ = _dafny.Map({d_548_v15_: d_548_v15_})
            d_551_v17_: _dafny.Map
            d_551_v17_ = _dafny.Map({(self).f23: ((d_550_v16_)[d_548_v15_] if (d_548_v15_) in (d_550_v16_) else d_548_v15_)})
            d_551_v17_ = (d_551_v17_).set((d_542_v8_) != (d_542_v8_), d_548_v15_)
            d_552_v18_: D6
            d_552_v18_ = D6_DC15((d_548_v15_).f21)
            d_553_v19_: _dafny.MultiSet
            d_553_v19_ = _dafny.MultiSet([(d_546_v12_) | (d_546_v12_), (d_552_v18_).cf18])
            d_553_v19_ = d_553_v19_
        elif True:
            r0 = p1
            if (self).f23:
                d_554_v20_: _dafny.Map
                d_554_v20_ = _dafny.Map({(self).f23: self.f18})
                d_554_v20_ = (d_554_v20_).set(not((self).f23), self.f18)
                d_555_v21_: _dafny.Map
                d_555_v21_ = _dafny.Map({D1_DC6(not(not(False)), True): p0})
                d_556_v22_: D1
                d_556_v22_ = D1_DC6((self).f23, p0)
                d_555_v21_ = (d_555_v21_).set(d_556_v22_, (self).f23)
                (self).f22 = p2
                (globalState).f6 = (0) - (self.f22)
                d_557_v23_: _dafny.Map
                d_557_v23_ = _dafny.Map({(self).f23: p0})
                d_558_v24_: _dafny.Array
                nw97_ = _dafny.Array(None, 11)
                nw97_[int(0)] = (len(d_557_v23_)) - (len(d_542_v8_))
                nw97_[int(1)] = (744) + (p1)
                nw97_[int(2)] = p1
                nw97_[int(3)] = self.f22
                nw97_[int(4)] = (self).fm4((self).f23, self.f18, self.f18, p0, globalState)
                nw97_[int(5)] = default__.fm0(globalState)
                nw97_[int(6)] = p1
                nw97_[int(7)] = p3
                nw97_[int(8)] = -798
                nw97_[int(9)] = p2
                nw97_[int(10)] = p2
                d_558_v24_ = nw97_
                d_559_v25_: _dafny.Seq
                d_559_v25_ = _dafny.SeqWithoutIsStrInference([len(d_542_v8_)])
                d_560_v26_: _dafny.MultiSet
                d_560_v26_ = _dafny.MultiSet([(self).f23, p0])
                index99_ = default__.safeIndex(250, (d_558_v24_).length(0))
                (d_558_v24_)[index99_] = default__.safeModuloInt((d_559_v25_)[default__.safeIndex(p2, len(d_559_v25_))], (d_560_v26_).cardinality)
                index100_ = default__.safeIndex(250, (d_558_v24_).length(0))
                (d_558_v24_)[index100_] = p3
            elif True:
                d_561_v27_: T0
                nw98_ = C0()
                nw98_.ctor__(d_535_v2_, default__.fm14(356, globalState), self.f18, (self).f19)
                d_561_v27_ = nw98_
                d_562_v28_: _dafny.Seq
                d_562_v28_ = _dafny.SeqWithoutIsStrInference([d_561_v27_])
                d_563_v29_: _dafny.Map
                d_563_v29_ = _dafny.Map({(self).f23: (d_562_v28_) + (d_562_v28_)})
                d_564_v30_: _dafny.Array
                nw99_ = _dafny.Array(False, 25)
                d_564_v30_ = nw99_
                d_565_v31_: _dafny.Map
                d_565_v31_ = _dafny.Map({_dafny.CodePoint('i'): self.f22})
                d_566_v32_: C0
                nw100_ = C0()
                nw100_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")), d_565_v31_, self.f18, (d_561_v27_).f19)
                d_566_v32_ = nw100_
                d_567_v33_: _dafny.Map
                d_567_v33_ = _dafny.Map({d_566_v32_: not((self).f23)})
                index101_ = default__.safeIndex(163, (d_564_v30_).length(0))
                (d_564_v30_)[index101_] = ((d_567_v33_)[d_566_v32_] if (d_566_v32_) in (d_567_v33_) else p0)
                index102_ = default__.safeIndex(163, (d_564_v30_).length(0))
                rhs94_ = (d_563_v29_).set((self).f23, d_562_v28_)
                rhs95_ = (self).f23
                lhs69_ = d_564_v30_
                lhs70_ = default__.safeIndex(163, (d_564_v30_).length(0))
                d_563_v29_ = rhs94_
                lhs69_[lhs70_] = rhs95_
                d_568_v34_: _dafny.Map
                d_568_v34_ = _dafny.Map({(self).f23: ((self).f23) == ((self).f23)})
                d_568_v34_ = (d_568_v34_).set((_dafny.SeqWithoutIsStrInference([d_561_v27_.f18 for d_569_i1_ in range(default__.abs(42))])) <= (d_535_v2_), (self).f23)
                d_570_v35_: _dafny.Seq
                d_570_v35_ = _dafny.SeqWithoutIsStrInference([self.f22, p1, p2])
                d_571_v36_: _dafny.Set
                d_571_v36_ = _dafny.Set({304, len(d_570_v35_)})
                index103_ = default__.safeIndex(163, (d_564_v30_).length(0))
                (d_564_v30_)[index103_] = default__.fm1(p0, d_561_v27_.f18, (len(d_571_v36_)) - ((d_570_v35_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_561_v27_.f18 for d_572_i2_ in range(default__.abs(-813))])), len(d_570_v35_))]), globalState)
                d_573_v37_: _dafny.Seq
                d_573_v37_ = _dafny.SeqWithoutIsStrInference([(d_566_v32_).f20, d_535_v2_, (d_535_v2_) + (d_535_v2_)])
                rhs96_ = (d_570_v35_)[default__.safeIndex(len(d_536_v3_), len(d_570_v35_))]
                rhs97_ = d_573_v37_
                rhs98_ = default__.safeModuloInt((D5_DC14(not(False), p1)).cf17, default__.safeDivisionInt(-221, 274))
                lhs71_ = globalState
                lhs72_ = globalState
                lhs71_.f6 = rhs96_
                d_573_v37_ = rhs97_
                lhs72_.f6 = rhs98_
                d_535_v2_ = ((d_566_v32_).f20).set(default__.safeIndex(p1, len((d_566_v32_).f20)), d_561_v27_.f18)
            if (self).fm3(not((p0) and ((self).f23)), (len(d_535_v2_)) + (len(d_535_v2_)), globalState):
                (self).f18 = self.f18
                d_574_v38_: bool
                d_574_v38_ = False
                d_574_v38_ = not(p0)
                (self).f22 = p2
                d_575_v39_: D1
                d_575_v39_ = D1_DC4(p3, p1, d_574_v38_)
                d_576_v40_: _dafny.Map
                d_576_v40_ = _dafny.Map({(d_575_v39_).cf5: (self).f23})
                d_576_v40_ = (d_576_v40_).set((self).f23, (self).f23)
                (self).f22 = (0) - (p2)
            elif True:
                d_577_v41_: _dafny.MultiSet
                d_577_v41_ = _dafny.MultiSet([p1, p1, p1, p1])
                d_578_v42_: D3
                d_578_v42_ = D3_DC10((d_577_v41_).issubset(d_577_v41_))
                d_578_v42_ = (D3_DC10((self).f23) if (self).f23 else d_578_v42_)
                d_579_v43_: _dafny.Array
                def lambda17_(d_580_p3_):
                    def lambda18_(d_581_i3_):
                        return default__.safeModuloInt(d_581_i3_, d_580_p3_)

                    return lambda18_

                init9_ = lambda17_(p3)
                nw101_ = _dafny.Array(None, 10)
                for i0_9_ in range(nw101_.length(0)):
                    nw101_[i0_9_] = init9_(i0_9_)
                d_579_v43_ = nw101_
                index104_ = default__.safeIndex(753, (d_579_v43_).length(0))
                (d_579_v43_)[index104_] = p3
                d_582_v44_: _dafny.Set
                d_582_v44_ = _dafny.Set({p2, p3})
                index105_ = default__.safeIndex(753, (d_579_v43_).length(0))
                (d_579_v43_)[index105_] = (p2) + ((693) * (len(d_582_v44_)))
                d_583_v45_: _dafny.Seq
                d_583_v45_ = _dafny.SeqWithoutIsStrInference([d_533_v0_, d_533_v0_, d_533_v0_, d_533_v0_, d_533_v0_])
                d_584_v46_: _dafny.Seq
                d_584_v46_ = _dafny.SeqWithoutIsStrInference([d_583_v45_])
                d_585_v47_: _dafny.Set
                d_585_v47_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqhnwiyw")), d_535_v2_, d_535_v2_})
                d_586_v48_: _dafny.Map
                d_586_v48_ = _dafny.Map({d_535_v2_: (d_584_v46_)[default__.safeIndex(len(d_585_v47_), len(d_584_v46_))]})
                d_586_v48_ = (d_586_v48_).set((d_535_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrv"))), d_583_v45_)
                d_587_v49_: bool
                d_587_v49_ = True
                d_588_v50_: _dafny.MultiSet
                d_588_v50_ = _dafny.MultiSet([(d_535_v2_)[default__.safeIndex(p2, len(d_535_v2_))], self.f18])
                d_587_v49_ = ((d_542_v8_)[default__.safeIndex(((d_588_v50_)[self.f18] if (self.f18) in (d_588_v50_) else (d_577_v41_).cardinality), len(d_542_v8_))] if d_587_v49_ else ((d_588_v50_).cardinality) <= ((d_579_v43_)[default__.safeIndex(753, (d_579_v43_).length(0))]))
                d_589_v51_: _dafny.Map
                d_589_v51_ = _dafny.Map({d_587_v49_: self.f22})
                d_589_v51_ = (d_589_v51_).set(p0, (843) - (p2))
            if (d_542_v8_)[default__.safeIndex(self.f22, len(d_542_v8_))]:
                d_590_v52_: _dafny.Array
                nw102_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_590_v52_ = nw102_
                d_591_v53_: D6
                d_591_v53_ = D6_DC16(not((self).f23), d_590_v52_)
                d_592_v54_: D6
                d_592_v54_ = D6_DC17(d_591_v53_)
                d_593_v55_: _dafny.Map
                d_593_v55_ = _dafny.Map({d_592_v54_: _dafny.CodePoint('y')})
                d_593_v55_ = (d_593_v55_).set(d_592_v54_, self.f18)
                (self).f22 = len(_dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(p3, len(_dafny.Map({False: p0}))) for d_594_i4_ in range(default__.abs(671))]))
                d_595_v56_: _dafny.MultiSet
                d_595_v56_ = _dafny.MultiSet([527])
                d_596_v57_: _dafny.Set
                d_596_v57_ = _dafny.Set({_dafny.Map({p2: self.f18})})
                d_597_v58_: _dafny.Set
                d_597_v58_ = _dafny.Set({(self).f23})
                d_598_v59_: _dafny.Array
                nw103_ = _dafny.Array(None, 27)
                nw103_[int(0)] = p0
                nw103_[int(1)] = (self).f23
                nw103_[int(2)] = (self).f23
                nw103_[int(3)] = not (True) or (not(default__.fm1(p0, self.f18, (d_595_v56_).cardinality, globalState)))
                nw103_[int(4)] = not((self).f23)
                nw103_[int(5)] = p0
                nw103_[int(6)] = True
                nw103_[int(7)] = (self).fm3(False, len(_dafny.SeqWithoutIsStrInference([self.f22])), globalState)
                nw103_[int(8)] = p0
                nw103_[int(9)] = not((self).f23)
                nw103_[int(10)] = (self).f23
                nw103_[int(11)] = (self).f23
                nw103_[int(12)] = True
                nw103_[int(13)] = (d_542_v8_)[default__.safeIndex(p2, len(d_542_v8_))]
                nw103_[int(14)] = p0
                nw103_[int(15)] = (self.f22) >= (p2)
                nw103_[int(16)] = (d_596_v57_).ispropersubset(d_596_v57_)
                nw103_[int(17)] = p0
                nw103_[int(18)] = True
                nw103_[int(19)] = p0
                nw103_[int(20)] = not((self).f23)
                nw103_[int(21)] = (self).f23
                nw103_[int(22)] = (self).f23
                nw103_[int(23)] = (d_597_v58_).isdisjoint(d_597_v58_)
                nw103_[int(24)] = p0
                nw103_[int(25)] = p0
                nw103_[int(26)] = (self).f23
                d_598_v59_ = nw103_
                index106_ = default__.safeIndex(600, (d_598_v59_).length(0))
                (d_598_v59_)[index106_] = p0
                index107_ = default__.safeIndex(600, (d_598_v59_).length(0))
                (d_598_v59_)[index107_] = (p0 if (p1) <= (len(d_535_v2_)) else (p2) != (p2))
                d_599_v60_: D1
                d_599_v60_ = D1_DC4(p2, 192, (self).f23)
                d_600_v61_: _dafny.Seq
                d_600_v61_ = _dafny.SeqWithoutIsStrInference([(self).fm4(default__.fm1((self).f23, self.f18, p3, globalState), self.f18, _dafny.CodePoint('d'), not(p0), globalState)])
                d_601_v62_: _dafny.Set
                d_601_v62_ = _dafny.Set({p3, self.f22, 241, p2, (d_595_v56_).cardinality})
                d_602_v63_: _dafny.Array
                nw104_ = _dafny.Array(None, 24)
                nw104_[int(0)] = d_599_v60_
                nw104_[int(1)] = d_599_v60_
                nw104_[int(2)] = D1_DC4(p1, p1, True)
                nw104_[int(3)] = d_599_v60_
                nw104_[int(4)] = D1_DC4(p1, p3, (self).f23)
                nw104_[int(5)] = d_599_v60_
                nw104_[int(6)] = d_599_v60_
                nw104_[int(7)] = D1_DC4(len(d_600_v61_), len(d_601_v62_), False)
                nw104_[int(8)] = d_599_v60_
                nw104_[int(9)] = d_599_v60_
                nw104_[int(10)] = d_599_v60_
                nw104_[int(11)] = d_599_v60_
                nw104_[int(12)] = d_599_v60_
                nw104_[int(13)] = d_599_v60_
                nw104_[int(14)] = D1_DC4(p3, p3, (self).f23)
                nw104_[int(15)] = d_599_v60_
                nw104_[int(16)] = d_599_v60_
                nw104_[int(17)] = d_599_v60_
                nw104_[int(18)] = d_599_v60_
                nw104_[int(19)] = d_599_v60_
                nw104_[int(20)] = D1_DC4(self.f22, p2, p0)
                nw104_[int(21)] = d_599_v60_
                nw104_[int(22)] = D1_DC4(len(d_542_v8_), self.f22, (d_542_v8_)[default__.safeIndex(p3, len(d_542_v8_))])
                nw104_[int(23)] = d_599_v60_
                d_602_v63_ = nw104_
                d_603_v64_: _dafny.MultiSet
                d_603_v64_ = _dafny.MultiSet([d_602_v63_, d_602_v63_, d_602_v63_])
                d_603_v64_ = d_603_v64_
                d_604_v65_: _dafny.Map
                d_604_v65_ = _dafny.Map({self.f18: p3})
                d_605_v66_: C0
                nw105_ = C0()
                nw105_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scd")), d_604_v65_, self.f18, (self).f19)
                d_605_v66_ = nw105_
            elif True:
                (globalState).f6 = p3
                d_606_v67_: _dafny.MultiSet
                d_606_v67_ = _dafny.MultiSet([len(d_542_v8_), p3, 542])
                rhs99_ = p2
                rhs100_ = d_606_v67_
                lhs73_ = self
                lhs73_.f22 = rhs99_
                d_606_v67_ = rhs100_
                d_607_v68_: D3
                d_607_v68_ = D3_DC8((self).fm13((0) - (p2), p1, p2, (self).f23, globalState))
                d_608_v69_: _dafny.Map
                d_608_v69_ = _dafny.Map({d_607_v68_: (self).f23})
                d_609_v70_: _dafny.Set
                d_609_v70_ = _dafny.Set({p0, not(p0), p0})
                def iife56_(_pat_let12_0):
                    def iife57_(d_610_dt__update__tmp_h0_):
                        def iife58_(_pat_let13_0):
                            def iife59_(d_611_dt__update_hcf10_h0_):
                                return D3_DC8(d_611_dt__update_hcf10_h0_)
                            return iife59_(_pat_let13_0)
                        return iife58_(_dafny.CodePoint('a'))
                    return iife57_(_pat_let12_0)
                d_608_v69_ = (d_608_v69_).set(iife56_(d_607_v68_), (_dafny.Set({p0})).ispropersubset(d_609_v70_))
                d_612_v71_: bool
                d_612_v71_ = True
                d_612_v71_ = d_612_v71_
                d_613_v72_: _dafny.Array
                nw106_ = _dafny.Array(int(0), 23)
                d_613_v72_ = nw106_
                index108_ = default__.safeIndex(849, (d_613_v72_).length(0))
                (d_613_v72_)[index108_] = default__.safeModuloInt(p3, p3)
                d_614_v73_: _dafny.Map
                d_614_v73_ = _dafny.Map({self.f22: 931})
                index109_ = default__.safeIndex(849, (d_613_v72_).length(0))
                rhs101_ = self.f22
                rhs102_ = default__.safeDivisionInt(((d_606_v67_)[p1] if (p1) in (d_606_v67_) else ((d_614_v73_)[self.f22] if (self.f22) in (d_614_v73_) else 698)), (0) - ((p1 if d_612_v71_ else self.f22)))
                rhs103_ = (p0) and (False)
                lhs74_ = d_613_v72_
                lhs75_ = default__.safeIndex(849, (d_613_v72_).length(0))
                lhs76_ = globalState
                lhs74_[lhs75_] = rhs101_
                lhs76_.f6 = rhs102_
                d_612_v71_ = rhs103_
            d_615_v74_: bool
            d_615_v74_ = True
            d_615_v74_ = default__.fm1((d_534_v1_).issubset(d_534_v1_), _dafny.CodePoint('y'), p3, globalState)
        (globalState).f6 = -101
        d_616_i5_: int
        d_616_i5_ = 0
        with _dafny.label("6"):
            while (d_542_v8_)[default__.safeIndex(297, len(d_542_v8_))]:
                with _dafny.c_label("6"):
                    if (d_616_i5_) >= (100):
                        raise _dafny.Break("6")
                    d_616_i5_ = (d_616_i5_) + (1)
                    d_617_v75_: _dafny.Array
                    def lambda19_(d_618_p3_):
                        def lambda20_(d_619_i6_):
                            return (d_619_i6_) + (d_618_p3_)

                        return lambda20_

                    init10_ = lambda19_(p3)
                    nw107_ = _dafny.Array(None, 4)
                    for i0_10_ in range(nw107_.length(0)):
                        nw107_[i0_10_] = init10_(i0_10_)
                    d_617_v75_ = nw107_
                    d_620_v76_: _dafny.Map
                    d_620_v76_ = _dafny.Map({(self).f23: False})
                    d_621_v77_: _dafny.Map
                    d_621_v77_ = _dafny.Map({p1: (0) - (len(d_620_v76_))})
                    index110_ = default__.safeIndex(240, (d_617_v75_).length(0))
                    (d_617_v75_)[index110_] = default__.safeDivisionInt(p1, len(d_621_v77_))
                    index111_ = default__.safeIndex(240, (d_617_v75_).length(0))
                    (d_617_v75_)[index111_] = p1
                    d_622_v78_: _dafny.MultiSet
                    d_622_v78_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([p0, (self).f23])), p1])
                    d_623_v79_: D3
                    d_623_v79_ = D3_DC9(p0, (d_622_v78_).cardinality)
                    d_624_v80_: D1
                    d_624_v80_ = D1_DC4((d_623_v79_).cf12, self.f22, (self).f23)
                    d_624_v80_ = d_624_v80_
                    (globalState).f6 = self.f22
                    d_617_v75_ = d_617_v75_
                    pass
            pass
        r0 = (p1) + (default__.safeModuloInt(599, self.f22))
        return r0

    @property
    def f23(self):
        return self._f23

class C2(T1, T0):
    def  __init__(self):
        self._f18: str = _dafny.CodePoint('D')
        self._f19: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value
    @property
    def f19(self):
        return self._f19
    def ctor__(self, f18, f19):
        (self).f18 = f18
        (self)._f19 = f19

    def fm3(self, p0, p1, globalState):
        if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fywgn")))) <= (896):
            return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnebdr")))) <= (329)
        elif True:
            return True

    def fm4(self, p0, p1, p2, p3, globalState):
        return (0) - ((len((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False, False, not(True), False])))) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, not(not(False)), True, not(not(True)), not(False)])) for d_625_i0_ in range(default__.abs(971))]))))

    def m1(self, p0, globalState):
        r0: D0 = D0.default()()
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_626_v0_: _dafny.Array
        def lambda21_(d_627_i0_):
            return (d_627_i0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fugltux"))))

        init11_ = lambda21_
        nw108_ = _dafny.Array(None, 15)
        for i0_11_ in range(nw108_.length(0)):
            nw108_[i0_11_] = init11_(i0_11_)
        d_626_v0_ = nw108_
        d_628_v1_: int
        d_628_v1_ = -605
        index112_ = default__.safeIndex(438, (d_626_v0_).length(0))
        (d_626_v0_)[index112_] = d_628_v1_
        d_629_v2_: _dafny.Array
        def lambda22_(d_630_p0_, d_631_v1_):
            def lambda23_(d_632_i1_):
                return (_dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxgl")): d_630_p0_})): D0_DC0(len(_dafny.SeqWithoutIsStrInference([d_631_v1_ for d_633_i2_ in range(default__.abs(-553))])))})) | (_dafny.Map({d_631_v1_: D0_DC0(d_631_v1_)}))

            return lambda23_

        init12_ = lambda22_(p0, d_628_v1_)
        nw109_ = _dafny.Array(None, 28)
        for i0_12_ in range(nw109_.length(0)):
            nw109_[i0_12_] = init12_(i0_12_)
        d_629_v2_ = nw109_
        index113_ = default__.safeIndex(969, (d_626_v0_).length(0))
        (d_626_v0_)[index113_] = d_628_v1_
        d_634_v3_: _dafny.Seq
        d_634_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuruhxs"))
        d_635_v4_: _dafny.Map
        d_635_v4_ = _dafny.Map({d_628_v1_: d_634_v3_})
        d_636_v5_: _dafny.Array
        nw110_ = _dafny.Array(False, 5)
        d_636_v5_ = nw110_
        d_637_v6_: _dafny.MultiSet
        d_637_v6_ = _dafny.MultiSet([d_636_v5_, d_636_v5_])
        d_638_v7_: _dafny.Seq
        d_638_v7_ = _dafny.SeqWithoutIsStrInference([not(p0)])
        d_639_v8_: _dafny.Seq
        d_639_v8_ = d_638_v7_
        pat_let_tv10_ = d_628_v1_
        pat_let_tv11_ = d_628_v1_
        index114_ = default__.safeIndex(438, (d_626_v0_).length(0))
        index115_ = default__.safeIndex(969, (d_626_v0_).length(0))
        def lambda24_(source7_):
            d_640___mcc_h0_ = source7_
            d_641_cf9_ = d_640___mcc_h0_
            return (pat_let_tv10_) * (pat_let_tv11_)

        rhs104_ = (len(d_635_v4_)) <= ((579) * ((d_637_v6_).cardinality))
        rhs105_ = self.f18
        rhs106_ = d_628_v1_
        rhs107_ = d_629_v2_
        rhs108_ = lambda24_(d_639_v8_)
        lhs77_ = self
        lhs78_ = d_626_v0_
        lhs79_ = default__.safeIndex(438, (d_626_v0_).length(0))
        lhs80_ = d_626_v0_
        lhs81_ = default__.safeIndex(969, (d_626_v0_).length(0))
        r2 = rhs104_
        lhs77_.f18 = rhs105_
        lhs78_[lhs79_] = rhs106_
        d_629_v2_ = rhs107_
        lhs80_[lhs81_] = rhs108_
        d_642_v9_: _dafny.Map
        d_642_v9_ = _dafny.Map({self.f18: (d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]})
        d_643_v10_: C0
        nw111_ = C0()
        nw111_.ctor__(d_634_v3_, d_642_v9_, self.f18, (self).f19)
        d_643_v10_ = nw111_
        d_628_v1_ = (d_628_v1_) * (988)
        d_644_v11_: _dafny.Map
        d_644_v11_ = _dafny.Map({(d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]: (d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]})
        d_645_v12_: _dafny.Seq
        d_645_v12_ = _dafny.SeqWithoutIsStrInference([(len(d_644_v11_)) + (d_628_v1_), (d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))], (d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]])
        r1 = (d_645_v12_)[default__.safeIndex(d_628_v1_, len(d_645_v12_))]
        hi4_ = d_628_v1_
        for d_646_i3_ in range((d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))], hi4_):
            d_647_v13_: bool
            out23_: bool
            out23_ = (self).m4((d_646_i3_) * (d_628_v1_), default__.fm0(globalState), d_636_v5_, globalState)
            d_647_v13_ = out23_
            r3 = d_647_v13_
            d_648_v14_: D0
            d_648_v14_ = D0_DC1()
            source8_ = d_648_v14_
            if source8_.is_DC1:
                d_649_v15_: _dafny.MultiSet
                d_649_v15_ = _dafny.MultiSet([d_647_v13_])
                index116_ = default__.safeIndex(438, (d_626_v0_).length(0))
                (d_626_v0_)[index116_] = (d_649_v15_).cardinality
                (globalState).f13 = ((d_638_v7_) + (d_638_v7_)) + (_dafny.SeqWithoutIsStrInference([d_647_v13_]))
                index117_ = default__.safeIndex(264, (d_636_v5_).length(0))
                (d_636_v5_)[index117_] = False
                index118_ = default__.safeIndex(56, (d_636_v5_).length(0))
                (d_636_v5_)[index118_] = d_647_v13_
                d_650_v16_: _dafny.MultiSet
                d_650_v16_ = _dafny.MultiSet([d_626_v0_, d_626_v0_, d_626_v0_])
                d_651_v17_: _dafny.Seq
                d_651_v17_ = _dafny.SeqWithoutIsStrInference([d_650_v16_, d_650_v16_])
                index119_ = default__.safeIndex(264, (d_636_v5_).length(0))
                index120_ = default__.safeIndex(56, (d_636_v5_).length(0))
                rhs109_ = d_647_v13_
                rhs110_ = default__.fm0(globalState)
                rhs111_ = ((d_651_v17_)[default__.safeIndex(d_628_v1_, len(d_651_v17_))]) != ((d_650_v16_).set(d_626_v0_, default__.abs(len((d_643_v10_).f20))))
                rhs112_ = (d_646_i3_) >= (default__.safeModuloInt((d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))], d_628_v1_))
                rhs113_ = p0
                lhs82_ = globalState
                lhs83_ = d_636_v5_
                lhs84_ = default__.safeIndex(264, (d_636_v5_).length(0))
                lhs85_ = d_636_v5_
                lhs86_ = default__.safeIndex(56, (d_636_v5_).length(0))
                d_647_v13_ = rhs109_
                lhs82_.f6 = rhs110_
                d_647_v13_ = rhs111_
                lhs83_[lhs84_] = rhs112_
                lhs85_[lhs86_] = rhs113_
                d_647_v13_ = d_647_v13_
            elif source8_.is_DC0:
                d_652___mcc_h1_ = source8_.cf0
                d_653_cf0_ = d_652___mcc_h1_
                d_654_v18_: _dafny.Set
                d_654_v18_ = _dafny.Set({d_647_v13_})
                d_654_v18_ = d_654_v18_
                (globalState).f6 = (d_628_v1_) * (d_628_v1_)
                d_655_v19_: _dafny.Array
                nw112_ = _dafny.Array(_dafny.Seq({}), 20)
                d_655_v19_ = nw112_
                index121_ = default__.safeIndex(883, (d_655_v19_).length(0))
                (d_655_v19_)[index121_] = _dafny.SeqWithoutIsStrInference([d_647_v13_])
                index122_ = default__.safeIndex(883, (d_655_v19_).length(0))
                (d_655_v19_)[index122_] = d_639_v8_
                d_656_v20_: _dafny.Map
                d_656_v20_ = _dafny.Map({d_634_v3_: not(p0)})
                d_657_v21_: _dafny.Map
                d_657_v21_ = _dafny.Map({(-196 if d_647_v13_ else d_646_i3_): (_dafny.Map({(d_643_v10_).f20: d_647_v13_})) | (d_656_v20_)})
                d_657_v21_ = (d_657_v21_).set((d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))], (default__.fm11(globalState) if d_647_v13_ else (d_656_v20_).set((d_643_v10_).f20, True)))
            elif True:
                d_658___mcc_h2_ = source8_.cf1
                d_659_cf1_ = d_658___mcc_h2_
                d_660_v22_: T0
                nw113_ = C0()
                nw113_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfwpi")), _dafny.Map({self.f18: d_628_v1_}), _dafny.CodePoint('m'), _dafny.Map({((d_643_v10_).f20).set(default__.safeIndex((d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))], len((d_643_v10_).f20)), self.f18): (d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]}))
                d_660_v22_ = nw113_
                d_661_v23_: T0
                d_661_v23_ = d_660_v22_
                d_662_v24_: T0
                d_662_v24_ = (d_661_v23_)
                d_660_v22_ = (d_660_v22_)
                d_663_v26_: _dafny.MultiSet
                d_663_v26_ = _dafny.MultiSet([d_634_v3_])
                d_664_v27_: C0
                nw114_ = C0()
                def iife60_():
                    coll32_ = _dafny.Map()
                    compr_32_: _dafny.Seq
                    for compr_32_ in (d_663_v26_).Elements:
                        d_665_v25_: _dafny.Seq = compr_32_
                        if (d_665_v25_) in (d_663_v26_):
                            coll32_[d_665_v25_] = len((d_643_v10_).f20)
                    return _dafny.Map(coll32_)
                nw114_.ctor__((d_643_v10_).f20, d_642_v9_, _dafny.CodePoint('l'), iife60_()
                )
                d_664_v27_ = nw114_
                d_666_v28_: _dafny.Map
                d_666_v28_ = _dafny.Map({d_643_v10_: (0) - (d_628_v1_)})
                rhs114_ = (((d_666_v28_)[d_664_v27_] if (d_664_v27_) in (d_666_v28_) else (d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))])) < (d_628_v1_)
                rhs115_ = 6
                d_647_v13_ = rhs114_
                d_628_v1_ = rhs115_
                d_667_v29_: T1
                nw115_ = C1()
                nw115_.ctor__(d_628_v1_, d_647_v13_, self.f18, default__.fm18(len((d_643_v10_).f20), globalState))
                d_667_v29_ = nw115_
                d_668_v30_: _dafny.Map
                d_668_v30_ = _dafny.Map({d_638_v7_: d_667_v29_})
                d_669_v31_: _dafny.Map
                d_669_v31_ = _dafny.Map({False: _dafny.Map({d_638_v7_: d_667_v29_})})
                rhs116_ = (((d_669_v31_)[p0] if (p0) in (d_669_v31_) else d_668_v30_)) | (d_668_v30_)
                rhs117_ = (d_628_v1_) - (len(_dafny.SeqWithoutIsStrInference([d_667_v29_.f18 for d_670_i4_ in range(default__.abs(357))])))
                d_668_v30_ = rhs116_
                d_628_v1_ = rhs117_
            r2 = not(p0)
        index123_ = default__.safeIndex(496, (d_636_v5_).length(0))
        (d_636_v5_)[index123_] = p0
        index124_ = default__.safeIndex(438, (d_626_v0_).length(0))
        index125_ = default__.safeIndex(496, (d_636_v5_).length(0))
        rhs118_ = (d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]
        rhs119_ = ((d_643_v10_).f20) == ((d_643_v10_).f20)
        rhs120_ = self.f18
        rhs121_ = d_628_v1_
        lhs87_ = d_626_v0_
        lhs88_ = default__.safeIndex(438, (d_626_v0_).length(0))
        lhs89_ = d_636_v5_
        lhs90_ = default__.safeIndex(496, (d_636_v5_).length(0))
        lhs91_ = self
        lhs87_[lhs88_] = rhs118_
        lhs89_[lhs90_] = rhs119_
        lhs91_.f18 = rhs120_
        d_628_v1_ = rhs121_
        r0 = D0_DC0(((d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]) - ((self).fm4(p0, self.f18, _dafny.CodePoint('r'), (d_636_v5_)[default__.safeIndex(496, (d_636_v5_).length(0))], globalState)))
        r1 = ((d_644_v11_)[(d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]] if ((d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]) in (d_644_v11_) else len(_dafny.SeqWithoutIsStrInference([d_644_v11_ for d_671_i5_ in range(default__.abs(380))])))
        r2 = (p0 if (d_636_v5_)[default__.safeIndex(496, (d_636_v5_).length(0))] else ((d_626_v0_)[default__.safeIndex(438, (d_626_v0_).length(0))]) < (d_628_v1_))
        d_672_v32_: _dafny.Set
        d_672_v32_ = _dafny.Set({not(False)})
        r3 = (d_672_v32_).issubset(d_672_v32_)
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r0 = p2
        if p0:
            r0 = p3
            r0 = p3
            d_673_v0_: _dafny.Array
            nw116_ = _dafny.Array(False, 20)
            d_673_v0_ = nw116_
            d_673_v0_ = d_673_v0_
            index126_ = default__.safeIndex(256, (d_673_v0_).length(0))
            (d_673_v0_)[index126_] = p0
            index127_ = default__.safeIndex(256, (d_673_v0_).length(0))
            (d_673_v0_)[index127_] = (p3) > ((0) - (p3))
            index128_ = default__.safeIndex(256, (d_673_v0_).length(0))
            (d_673_v0_)[index128_] = (True) == ((p2) > (p3))
        elif True:
            d_674_v1_: bool
            d_674_v1_ = False
            d_675_v2_: _dafny.Set
            d_675_v2_ = _dafny.Set({d_674_v1_})
            d_674_v1_ = (d_675_v2_).issubset(d_675_v2_)
            d_676_v3_: _dafny.Seq
            d_676_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrx"))
            d_677_v4_: _dafny.Map
            d_677_v4_ = _dafny.Map({d_674_v1_: not(d_674_v1_)})
            d_678_v5_: _dafny.Map
            d_678_v5_ = _dafny.Map({(d_676_v3_)[default__.safeIndex(len(d_677_v4_), len(d_676_v3_))]: p2})
            d_679_v6_: D6
            d_679_v6_ = D6_DC15((d_678_v5_) | (d_678_v5_))
            d_680_v7_: _dafny.Array
            nw117_ = _dafny.Array(None, 8)
            nw117_[int(0)] = d_676_v3_
            nw117_[int(1)] = d_676_v3_
            nw117_[int(2)] = d_676_v3_
            nw117_[int(3)] = (d_676_v3_).set(default__.safeIndex(p2, len(d_676_v3_)), _dafny.CodePoint('p'))
            nw117_[int(4)] = d_676_v3_
            nw117_[int(5)] = d_676_v3_
            nw117_[int(6)] = d_676_v3_
            nw117_[int(7)] = d_676_v3_
            d_680_v7_ = nw117_
            d_681_v8_: _dafny.Array
            nw118_ = _dafny.Array(False, 17)
            d_681_v8_ = nw118_
            d_682_v9_: _dafny.MultiSet
            d_682_v9_ = _dafny.MultiSet([d_674_v1_, not(False), d_674_v1_, default__.fm1(d_674_v1_, _dafny.CodePoint('r'), p2, globalState)])
            index129_ = default__.safeIndex(138, (d_681_v8_).length(0))
            (d_681_v8_)[index129_] = (d_682_v9_).issubset(d_682_v9_)
            d_683_v10_: D0
            d_683_v10_ = D0_DC0((0) - (p1))
            index130_ = default__.safeIndex(138, (d_681_v8_).length(0))
            rhs122_ = D6_DC15((d_678_v5_) | (d_678_v5_))
            rhs123_ = d_680_v7_
            rhs124_ = default__.fm1(d_674_v1_, self.f18, -405, globalState)
            rhs125_ = D0_DC0(p2)
            rhs126_ = self.f18
            lhs92_ = d_681_v8_
            lhs93_ = default__.safeIndex(138, (d_681_v8_).length(0))
            lhs94_ = self
            d_679_v6_ = rhs122_
            d_680_v7_ = rhs123_
            lhs92_[lhs93_] = rhs124_
            d_683_v10_ = rhs125_
            lhs94_.f18 = rhs126_
            d_684_v11_: D1
            d_684_v11_ = D1_DC6(True, p0)
            index131_ = default__.safeIndex(138, (d_681_v8_).length(0))
            (d_681_v8_)[index131_] = (d_684_v11_).cf7
            d_685_v12_: _dafny.Map
            d_685_v12_ = _dafny.Map({p1: not((d_681_v8_)[default__.safeIndex(138, (d_681_v8_).length(0))])})
            d_686_v13_: _dafny.Map
            d_686_v13_ = _dafny.Map({len(d_676_v3_): ((d_685_v12_)[p3] if (p3) in (d_685_v12_) else p0)})
            d_686_v13_ = (d_686_v13_).set(p1, d_674_v1_)
            (globalState).f6 = default__.fm0(globalState)
        d_687_v14_: bool
        d_687_v14_ = True
        d_688_v15_: _dafny.Set
        d_688_v15_ = _dafny.Set({p0})
        d_689_v16_: D7
        d_689_v16_ = D7_DC18(d_688_v15_)
        d_687_v14_ = (d_688_v15_) == ((d_689_v16_).cf22)
        d_690_v17_: _dafny.Map
        d_690_v17_ = _dafny.Map({d_687_v14_: p0})
        d_691_v18_: _dafny.Map
        d_691_v18_ = _dafny.Map({len(((d_690_v17_).set(p0, d_687_v14_)) | (d_690_v17_)): p2})
        d_692_v19_: _dafny.Seq
        d_692_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oiyohkc"))
        d_693_v20_: _dafny.Array
        nw119_ = _dafny.Array(int(0), 15)
        d_693_v20_ = nw119_
        index132_ = default__.safeIndex(605, (d_693_v20_).length(0))
        (d_693_v20_)[index132_] = p1
        d_694_v21_: _dafny.MultiSet
        d_694_v21_ = _dafny.MultiSet([d_687_v14_, d_687_v14_])
        d_695_v22_: D5
        d_695_v22_ = D5_DC12(d_694_v21_)
        pat_let_tv12_ = p2
        pat_let_tv13_ = p1
        index133_ = default__.safeIndex(605, (d_693_v20_).length(0))
        def lambda25_(source9_):
            if source9_.is_DC13:
                return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))])) + ((_dafny.SeqWithoutIsStrInference([pat_let_tv12_]))))).cardinality
            elif source9_.is_DC14:
                d_696___mcc_h0_ = source9_.cf16
                d_697___mcc_h1_ = source9_.cf17
                d_698_cf17_ = d_697___mcc_h1_
                d_699_cf16_ = d_696___mcc_h0_
                return (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_700_i0_ in range(default__.abs(378))])))
            elif True:
                d_701___mcc_h2_ = source9_.cf15
                d_702_cf15_ = d_701___mcc_h2_
                return pat_let_tv13_

        rhs127_ = (default__.fm19(p0, p0, globalState)) | (_dafny.Map({p2: p2}))
        rhs128_ = d_692_v19_
        rhs129_ = lambda25_(d_695_v22_)
        lhs95_ = d_693_v20_
        lhs96_ = default__.safeIndex(605, (d_693_v20_).length(0))
        d_691_v18_ = rhs127_
        d_692_v19_ = rhs128_
        lhs95_[lhs96_] = rhs129_
        d_703_v23_: _dafny.Map
        d_703_v23_ = _dafny.Map({d_687_v14_: p3})
        d_703_v23_ = (d_703_v23_).set(p0, p2)
        d_704_v24_: _dafny.MultiSet
        d_704_v24_ = _dafny.MultiSet([296])
        d_705_v25_: C1
        nw120_ = C1()
        nw120_.ctor__(p3, ((d_693_v20_)[default__.safeIndex(605, (d_693_v20_).length(0))]) in (d_704_v24_), _dafny.CodePoint('y'), ((self).f19) | ((self).f19))
        d_705_v25_ = nw120_
        d_706_v26_: _dafny.Seq
        d_706_v26_ = _dafny.SeqWithoutIsStrInference([d_705_v25_])
        d_705_v25_ = (d_706_v26_)[default__.safeIndex(p3, len(d_706_v26_))]
        r0 = p1
        return r0

    def m4(self, p0, p1, p2, globalState):
        r0: bool = False
        hi5_ = p1
        for d_707_i0_ in range((p1) + (p1), hi5_):
            d_708_v0_: bool
            d_708_v0_ = False
            d_709_v1_: _dafny.Map
            d_709_v1_ = _dafny.Map({d_708_v0_: p0})
            d_710_v2_: _dafny.MultiSet
            d_710_v2_ = _dafny.MultiSet([d_708_v0_])
            d_711_v3_: _dafny.Set
            d_711_v3_ = _dafny.Set({len(d_709_v1_), (d_710_v2_).cardinality})
            d_712_v4_: _dafny.Map
            d_712_v4_ = _dafny.Map({d_711_v3_: (p1) * (p0)})
            d_712_v4_ = (d_712_v4_).set(d_711_v3_, default__.fm0(globalState))
            d_713_v5_: _dafny.Seq
            d_713_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtyc"))
            (globalState).f6 = (p1) - (len(_dafny.SeqWithoutIsStrInference([d_713_v5_, d_713_v5_])))
            d_708_v0_ = True
            d_714_v6_: _dafny.Map
            d_714_v6_ = _dafny.Map({d_713_v5_: d_708_v0_})
            index134_ = default__.safeIndex(135, (p2).length(0))
            (p2)[index134_] = ((d_714_v6_)[d_713_v5_] if (d_713_v5_) in (d_714_v6_) else d_708_v0_)
            d_715_v7_: D5
            d_715_v7_ = D5_DC13()
            d_716_v8_: _dafny.Map
            d_716_v8_ = _dafny.Map({d_715_v7_: d_707_i0_})
            index135_ = default__.safeIndex(135, (p2).length(0))
            rhs130_ = (d_708_v0_) not in (default__.fm2(globalState))
            rhs131_ = d_716_v8_
            rhs132_ = default__.fm1((d_708_v0_) or (d_708_v0_), _dafny.CodePoint('v'), (0) - (p1), globalState)
            lhs97_ = p2
            lhs98_ = default__.safeIndex(135, (p2).length(0))
            lhs97_[lhs98_] = rhs130_
            d_716_v8_ = rhs131_
            d_708_v0_ = rhs132_
        d_717_v9_: bool
        d_717_v9_ = False
        d_718_i1_: int
        d_718_i1_ = 0
        with _dafny.label("7"):
            while d_717_v9_:
                with _dafny.c_label("7"):
                    if (d_718_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_718_i1_ = (d_718_i1_) + (1)
                    (globalState).f6 = (p0) - ((0) - ((0) - (p1)))
                    d_719_v10_: _dafny.Seq
                    d_719_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yesur"))
                    d_720_v11_: _dafny.Map
                    d_720_v11_ = _dafny.Map({d_717_v9_: (d_719_v10_)[default__.safeIndex(p1, len(d_719_v10_))]})
                    d_721_v12_: _dafny.Seq
                    d_721_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_717_v9_: _dafny.CodePoint('c')})])
                    d_722_v13_: _dafny.Map
                    d_722_v13_ = _dafny.Map({p0: d_719_v10_})
                    d_720_v11_ = (d_721_v12_)[default__.safeIndex(len(((d_722_v13_)[-768] if (-768) in (d_722_v13_) else _dafny.SeqWithoutIsStrInference([self.f18 for d_723_i2_ in range(default__.abs(175))]))), len(d_721_v12_))]
                    r0 = default__.fm1(d_717_v9_, self.f18, p1, globalState)
                    d_724_v14_: _dafny.Map
                    d_724_v14_ = _dafny.Map({not(d_717_v9_): d_717_v9_})
                    d_725_v15_: _dafny.Seq
                    d_725_v15_ = _dafny.SeqWithoutIsStrInference([d_717_v9_])
                    d_726_v16_: _dafny.Map
                    d_726_v16_ = _dafny.Map({-623: self.f18})
                    d_727_v17_: _dafny.Map
                    d_727_v17_ = _dafny.Map({d_725_v15_: d_726_v16_})
                    d_728_v19_: _dafny.Set
                    d_728_v19_ = _dafny.Set({d_725_v15_, d_725_v15_, d_725_v15_})
                    d_729_v20_: _dafny.Array
                    nw121_ = _dafny.Array(None, 21)
                    nw121_[int(0)] = d_717_v9_
                    nw121_[int(1)] = d_717_v9_
                    nw121_[int(2)] = (True if d_717_v9_ else not(((d_724_v14_)[d_717_v9_] if (d_717_v9_) in (d_724_v14_) else d_717_v9_)))
                    nw121_[int(3)] = d_717_v9_
                    nw121_[int(4)] = d_717_v9_
                    nw121_[int(5)] = (d_717_v9_) in (d_725_v15_)
                    nw121_[int(6)] = d_717_v9_
                    nw121_[int(7)] = True
                    nw121_[int(8)] = True
                    nw121_[int(9)] = not((d_717_v9_ if d_717_v9_ else True))
                    nw121_[int(10)] = d_717_v9_
                    nw121_[int(11)] = (d_725_v15_)[default__.safeIndex(len(_dafny.Map({p0: True})), len(d_725_v15_))]
                    nw121_[int(12)] = d_717_v9_
                    nw121_[int(13)] = d_717_v9_
                    nw121_[int(14)] = not (d_717_v9_) or (d_717_v9_)
                    nw121_[int(15)] = d_717_v9_
                    def iife61_():
                        coll33_ = _dafny.Set()
                        compr_33_: _dafny.Seq
                        for compr_33_ in (d_727_v17_).keys.Elements:
                            d_730_v18_: _dafny.Seq = compr_33_
                            if (d_730_v18_) in (d_727_v17_):
                                coll33_ = coll33_.union(_dafny.Set([d_730_v18_]))
                        return _dafny.Set(coll33_)
                    nw121_[int(16)] = (iife61_()
                    ).isdisjoint(d_728_v19_)
                    nw121_[int(17)] = d_717_v9_
                    nw121_[int(18)] = d_717_v9_
                    nw121_[int(19)] = d_717_v9_
                    nw121_[int(20)] = d_717_v9_
                    d_729_v20_ = nw121_
                    d_731_v21_: _dafny.Map
                    d_731_v21_ = _dafny.Map({(d_719_v10_)[default__.safeIndex(316, len(d_719_v10_))]: p1})
                    d_732_v22_: T0
                    nw122_ = C0()
                    nw122_.ctor__(_dafny.SeqWithoutIsStrInference([self.f18 for d_733_i3_ in range(default__.abs(214))]), d_731_v21_, self.f18, (self).f19)
                    d_732_v22_ = nw122_
                    d_734_v23_: _dafny.Map
                    d_734_v23_ = _dafny.Map({d_717_v9_: d_729_v20_})
                    rhs133_ = p2
                    rhs134_ = len((d_734_v23_ if d_717_v9_ else (d_734_v23_) | (d_734_v23_)))
                    rhs135_ = (d_732_v22_ if d_717_v9_ else d_732_v22_)
                    rhs136_ = default__.fm20(d_717_v9_, globalState)
                    lhs99_ = globalState
                    lhs100_ = self
                    d_729_v20_ = rhs133_
                    lhs99_.f6 = rhs134_
                    d_732_v22_ = rhs135_
                    lhs100_.f18 = rhs136_
                    pass
            pass
        d_735_v24_: _dafny.Array
        nw123_ = _dafny.Array(_dafny.Array(None, 0), 28)
        d_735_v24_ = nw123_
        index136_ = default__.safeIndex(400, (d_735_v24_).length(0))
        (d_735_v24_)[index136_] = p2
        index137_ = default__.safeIndex(400, (d_735_v24_).length(0))
        (d_735_v24_)[index137_] = p2
        d_736_v25_: _dafny.MultiSet
        d_736_v25_ = _dafny.MultiSet([_dafny.Set({-174, p1})])
        d_737_v26_: _dafny.Map
        d_737_v26_ = _dafny.Map({d_717_v9_: p0})
        d_738_v27_: _dafny.Seq
        d_738_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcyiogjr"))
        d_739_v28_: _dafny.Set
        d_739_v28_ = _dafny.Set({len(d_737_v26_), len(d_738_v27_)})
        d_740_v29_: _dafny.Seq
        d_740_v29_ = _dafny.SeqWithoutIsStrInference([p0])
        def iife62_():
            coll34_ = _dafny.Set()
            compr_34_: int
            for compr_34_ in (d_740_v29_).Elements:
                d_741_v30_: int = compr_34_
                if (d_741_v30_) in (d_740_v29_):
                    def iife63_():
                        def iife65_():
                            coll37_ = _dafny.Map()
                            compr_37_: int
                            for compr_37_ in _dafny.IntegerRange(814, 478):
                                d_742_v31_: int = compr_37_
                                if ((814) <= (d_742_v31_)) and ((d_742_v31_) < (478)):
                                    coll37_[(d_742_v31_) + (-657)] = 846
                            return _dafny.Map(coll37_)
                        coll35_ = _dafny.Set()
                        def iife64_():
                            coll36_ = _dafny.Map()
                            compr_36_: int
                            for compr_36_ in _dafny.IntegerRange(814, 478):
                                d_742_v31_: int = compr_36_
                                if ((814) <= (d_742_v31_)) and ((d_742_v31_) < (478)):
                                    coll36_[(d_742_v31_) + (-657)] = 846
                            return _dafny.Map(coll36_)
                        compr_35_: int
                        for compr_35_ in (iife64_()
                        ).keys.Elements:
                            d_743_v32_: int = compr_35_
                            if (d_743_v32_) in (iife65_()
                            ):
                                coll35_ = coll35_.union(_dafny.Set([(d_743_v32_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_744_i4_ in range(default__.abs(883))])))]))
                        return _dafny.Set(coll35_)
                    coll34_ = coll34_.union(_dafny.Set([default__.safeDivisionInt(d_741_v30_, len(iife63_()
))]))
            return _dafny.Set(coll34_)
        d_736_v25_ = (_dafny.MultiSet([d_739_v28_])).set((d_739_v28_) | (iife62_()
        ), default__.abs(default__.fm0(globalState)))
        d_745_i5_: int
        d_745_i5_ = 0
        with _dafny.label("8"):
            while (d_739_v28_).isdisjoint(d_739_v28_):
                with _dafny.c_label("8"):
                    if (d_745_i5_) >= (100):
                        raise _dafny.Break("8")
                    d_745_i5_ = (d_745_i5_) + (1)
                    (globalState).f6 = p0
                    (globalState).f6 = p0
                    d_746_v33_: _dafny.Map
                    d_746_v33_ = _dafny.Map({self.f18: p0})
                    d_747_v34_: T0
                    nw124_ = C0()
                    nw124_.ctor__(d_738_v27_, d_746_v33_, self.f18, (self).f19)
                    d_747_v34_ = nw124_
                    d_748_v35_: T0
                    d_748_v35_ = d_747_v34_
                    source10_ = d_748_v35_
                    d_749___mcc_h0_ = source10_
                    d_750_cf14_ = d_749___mcc_h0_
                    d_751_v36_: _dafny.Seq
                    d_751_v36_ = _dafny.SeqWithoutIsStrInference([d_717_v9_])
                    d_738_v27_ = (((default__.fm21(globalState)).set(default__.safeIndex(p1, len(default__.fm21(globalState))), self.f18)).set(default__.safeIndex(p0, len((default__.fm21(globalState)).set(default__.safeIndex(p1, len(default__.fm21(globalState))), self.f18))), d_750_cf14_.f18)) + (((d_738_v27_) + (d_738_v27_)).set(default__.safeIndex(len(d_751_v36_), len((d_738_v27_) + (d_738_v27_))), self.f18))
                    d_752_v37_: _dafny.Seq
                    d_752_v37_ = _dafny.SeqWithoutIsStrInference([d_737_v26_])
                    d_753_v38_: _dafny.Map
                    d_753_v38_ = _dafny.Map({p0: default__.fm22(globalState)})
                    d_754_v39_: _dafny.Seq
                    d_754_v39_ = _dafny.SeqWithoutIsStrInference([((d_753_v38_)[p1] if (p1) in (d_753_v38_) else d_752_v37_)])
                    pat_let_tv14_ = d_717_v9_
                    pat_let_tv15_ = p0
                    def iife67_(_pat_let15_0):
                        def iife68_(d_755_dt__update__tmp_h0_):
                            def iife69_(_pat_let16_0):
                                def iife70_(d_756_dt__update_hcf16_h0_):
                                    return D5_DC14(d_756_dt__update_hcf16_h0_, (d_755_dt__update__tmp_h0_).cf17)
                                return iife70_(_pat_let16_0)
                            return iife69_(not(pat_let_tv14_))
                        return iife68_(_pat_let15_0)
                    def iife66_(_pat_let14_0):
                        def iife71_(d_757_dt__update__tmp_h1_):
                            def iife72_(_pat_let17_0):
                                def iife73_(d_758_dt__update_hcf17_h0_):
                                    return D5_DC14((d_757_dt__update__tmp_h1_).cf16, d_758_dt__update_hcf17_h0_)
                                return iife73_(_pat_let17_0)
                            return iife72_(pat_let_tv15_)
                        return iife71_(_pat_let14_0)
                    d_752_v37_ = (d_754_v39_)[default__.safeIndex((iife66_(iife67_(D5_DC14(d_717_v9_, p1)))).cf17, len(d_754_v39_))]
                    d_759_v40_: _dafny.Map
                    d_759_v40_ = _dafny.Map({d_738_v27_: d_717_v9_})
                    d_759_v40_ = (d_759_v40_).set(d_738_v27_, d_717_v9_)
                    (globalState).f6 = (0) - (default__.fm0(globalState))
                    d_760_v41_: _dafny.Seq
                    d_760_v41_ = d_740_v29_
                    d_760_v41_ = d_760_v41_
                    pass
            pass
        d_761_v42_: _dafny.Array
        nw125_ = _dafny.Array(int(0), 9)
        d_761_v42_ = nw125_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_761_v42_).length(0)):
            d_762_i6_: int = guard_loop_1_
            if (True) and (((0) <= (d_762_i6_)) and ((d_762_i6_) < ((d_761_v42_).length(0)))):
                (d_761_v42_)[(d_762_i6_)] = (d_762_i6_) + (default__.safeModuloInt(p0, p1))
        r0 = not(d_717_v9_)
        return r0


class C3(T1, T0):
    def  __init__(self):
        self._f18: str = _dafny.CodePoint('D')
        self._f19: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value
    @property
    def f19(self):
        return self._f19
    def ctor__(self, f18, f19):
        (self).f18 = f18
        (self)._f19 = f19

    def fm3(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([937, 678])) for d_763_i0_ in range(default__.abs(514))])) + (_dafny.SeqWithoutIsStrInference([307 for d_764_i1_ in range(default__.abs(880))]))) != (_dafny.SeqWithoutIsStrInference([-125 for d_765_i2_ in range(default__.abs(647))]))

    def fm4(self, p0, p1, p2, p3, globalState):
        return default__.safeModuloInt(((_dafny.MultiSet([265])).intersection(_dafny.MultiSet([373, 858]))).cardinality, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkdhj"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtgkshg")))))

    def m1(self, p0, globalState):
        r0: D0 = D0.default()()
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_766_i0_: int
        d_766_i0_ = 0
        with _dafny.label("9"):
            while p0:
                with _dafny.c_label("9"):
                    if (d_766_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_766_i0_ = (d_766_i0_) + (1)
                    d_767_v0_: int
                    d_767_v0_ = 652
                    d_768_v1_: _dafny.MultiSet
                    d_768_v1_ = _dafny.MultiSet([d_767_v0_])
                    d_769_v2_: _dafny.Seq
                    d_769_v2_ = _dafny.SeqWithoutIsStrInference([self.f18])
                    d_770_v3_: _dafny.Set
                    d_770_v3_ = _dafny.Set({(0) - (d_767_v0_), d_767_v0_})
                    d_771_v4_: _dafny.Seq
                    d_771_v4_ = _dafny.SeqWithoutIsStrInference([False, p0])
                    d_772_v5_: _dafny.Seq
                    d_772_v5_ = d_771_v4_
                    d_773_v6_: _dafny.Map
                    d_773_v6_ = _dafny.Map({(d_772_v5_): d_767_v0_})
                    d_774_v7_: _dafny.Map
                    d_774_v7_ = _dafny.Map({len(d_773_v6_): p0})
                    d_775_v8_: _dafny.Map
                    d_775_v8_ = _dafny.Map({len(d_774_v7_): p0})
                    d_776_v9_: _dafny.Array
                    def lambda26_(d_777_v0_):
                        def lambda27_(d_778_i1_):
                            return default__.safeModuloInt(d_778_i1_, d_777_v0_)

                        return lambda27_

                    init13_ = lambda26_(d_767_v0_)
                    nw126_ = _dafny.Array(None, 21)
                    for i0_13_ in range(nw126_.length(0)):
                        nw126_[i0_13_] = init13_(i0_13_)
                    d_776_v9_ = nw126_
                    d_779_v10_: _dafny.Map
                    d_779_v10_ = _dafny.Map({d_776_v9_: (self).fm3(p0, d_767_v0_, globalState)})
                    d_780_v11_: _dafny.Array
                    nw127_ = _dafny.Array(None, 23)
                    nw127_[int(0)] = (p0) or (p0)
                    nw127_[int(1)] = (d_768_v1_).issubset(d_768_v1_)
                    nw127_[int(2)] = False
                    nw127_[int(3)] = (D1_DC4(d_767_v0_, d_767_v0_, True)).cf5
                    nw127_[int(4)] = (d_769_v2_) == (d_769_v2_)
                    nw127_[int(5)] = False
                    nw127_[int(6)] = p0
                    nw127_[int(7)] = p0
                    nw127_[int(8)] = (p0) or (True)
                    nw127_[int(9)] = (p0 if p0 else p0)
                    nw127_[int(10)] = p0
                    nw127_[int(11)] = p0
                    nw127_[int(12)] = not(p0)
                    nw127_[int(13)] = not((default__.fm5(d_767_v0_, globalState)).isdisjoint(d_770_v3_))
                    nw127_[int(14)] = (len(d_769_v2_)) > (len(d_774_v7_))
                    nw127_[int(15)] = (d_767_v0_) < (d_767_v0_)
                    nw127_[int(16)] = p0
                    nw127_[int(17)] = p0
                    nw127_[int(18)] = ((d_779_v10_)[d_776_v9_] if (d_776_v9_) in (d_779_v10_) else p0)
                    nw127_[int(19)] = (d_767_v0_) > (d_767_v0_)
                    nw127_[int(20)] = (d_768_v1_).isdisjoint(d_768_v1_)
                    nw127_[int(21)] = p0
                    nw127_[int(22)] = (p0) or (p0)
                    d_780_v11_ = nw127_
                    index138_ = default__.safeIndex(994, (d_780_v11_).length(0))
                    (d_780_v11_)[index138_] = p0
                    index139_ = default__.safeIndex(977, (d_776_v9_).length(0))
                    (d_776_v9_)[index139_] = d_767_v0_
                    index140_ = default__.safeIndex(994, (d_780_v11_).length(0))
                    index141_ = default__.safeIndex(977, (d_776_v9_).length(0))
                    rhs137_ = (d_767_v0_) > (d_767_v0_)
                    rhs138_ = (d_767_v0_) * (d_767_v0_)
                    rhs139_ = p0
                    rhs140_ = 620
                    lhs101_ = globalState
                    lhs102_ = d_780_v11_
                    lhs103_ = default__.safeIndex(994, (d_780_v11_).length(0))
                    lhs104_ = d_776_v9_
                    lhs105_ = default__.safeIndex(977, (d_776_v9_).length(0))
                    r3 = rhs137_
                    lhs101_.f6 = rhs138_
                    lhs102_[lhs103_] = rhs139_
                    lhs104_[lhs105_] = rhs140_
                    d_781_v12_: _dafny.Map
                    d_781_v12_ = _dafny.Map({len(d_769_v2_): d_767_v0_})
                    d_774_v7_ = (_dafny.Map({d_767_v0_: p0})).set((0) - (((d_776_v9_)[default__.safeIndex(977, (d_776_v9_).length(0))]) * (len(d_781_v12_))), (d_780_v11_)[default__.safeIndex(994, (d_780_v11_).length(0))])
                    d_782_v13_: C0
                    nw128_ = C0()
                    nw128_.ctor__(((d_769_v2_ if (D1_DC3(p0)).cf2 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "as")))).set(default__.safeIndex((d_776_v9_)[default__.safeIndex(977, (d_776_v9_).length(0))], len((d_769_v2_ if (D1_DC3(p0)).cf2 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "as"))))), self.f18), default__.fm7((d_780_v11_)[default__.safeIndex(994, (d_780_v11_).length(0))], d_767_v0_, globalState), self.f18, ((self).f19) | ((self).f19))
                    d_782_v13_ = nw128_
                    d_783_v14_: C0
                    nw129_ = C0()
                    nw129_.ctor__((d_782_v13_).f20, (d_782_v13_).f21, self.f18, (self).f19)
                    d_783_v14_ = nw129_
                    pass
            pass
        d_784_v15_: int
        d_784_v15_ = 854
        d_785_v16_: _dafny.MultiSet
        d_785_v16_ = _dafny.MultiSet([d_784_v15_])
        d_786_v17_: _dafny.Array
        nw130_ = _dafny.Array(None, 21)
        nw130_[int(0)] = d_784_v15_
        nw130_[int(1)] = d_784_v15_
        nw130_[int(2)] = d_784_v15_
        nw130_[int(3)] = d_784_v15_
        nw130_[int(4)] = 200
        nw130_[int(5)] = d_784_v15_
        nw130_[int(6)] = d_784_v15_
        nw130_[int(7)] = d_784_v15_
        nw130_[int(8)] = (0) - (d_784_v15_)
        nw130_[int(9)] = 497
        nw130_[int(10)] = d_784_v15_
        nw130_[int(11)] = d_784_v15_
        nw130_[int(12)] = d_784_v15_
        nw130_[int(13)] = (d_785_v16_).cardinality
        nw130_[int(14)] = d_784_v15_
        nw130_[int(15)] = d_784_v15_
        nw130_[int(16)] = d_784_v15_
        nw130_[int(17)] = d_784_v15_
        nw130_[int(18)] = d_784_v15_
        nw130_[int(19)] = 192
        nw130_[int(20)] = -38
        d_786_v17_ = nw130_
        d_787_v18_: _dafny.Map
        d_787_v18_ = _dafny.Map({d_786_v17_: self.f18})
        d_788_v19_: _dafny.Seq
        d_788_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbgkabcen"))
        hi6_ = (len(d_788_v19_)) - (d_784_v15_)
        for d_789_i2_ in range(len(d_787_v18_), hi6_):
            d_790_v20_: D3
            d_790_v20_ = D3_DC8(self.f18)
            d_791_v21_: _dafny.Map
            d_791_v21_ = _dafny.Map({p0: p0})
            d_792_v22_: _dafny.Seq
            d_792_v22_ = _dafny.SeqWithoutIsStrInference([p0])
            d_793_v23_: _dafny.Map
            d_793_v23_ = _dafny.Map({d_791_v21_: len(d_792_v22_)})
            d_794_v24_: _dafny.Seq
            d_794_v24_ = _dafny.SeqWithoutIsStrInference([p0, not (p0) or (default__.fm1(p0, (d_790_v20_).cf10, ((d_793_v23_)[d_791_v21_] if (d_791_v21_) in (d_793_v23_) else d_789_i2_), globalState)), p0])
            r1 = len(d_792_v22_)
            (globalState).f6 = d_784_v15_
            (globalState).f6 = d_784_v15_
            r1 = (d_784_v15_ if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqpigsw"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlp"))) else d_784_v15_)
        hi7_ = d_784_v15_
        for d_795_i3_ in range(d_784_v15_, hi7_):
            d_796_v25_: _dafny.Array
            def lambda28_(d_797_v19_):
                def lambda29_(d_798_i4_):
                    return d_797_v19_

                return lambda29_

            init14_ = lambda28_(d_788_v19_)
            nw131_ = _dafny.Array(None, 17)
            for i0_14_ in range(nw131_.length(0)):
                nw131_[i0_14_] = init14_(i0_14_)
            d_796_v25_ = nw131_
            index142_ = default__.safeIndex(962, (d_796_v25_).length(0))
            (d_796_v25_)[index142_] = d_788_v19_
            index143_ = default__.safeIndex(962, (d_796_v25_).length(0))
            (d_796_v25_)[index143_] = (d_788_v19_) + (d_788_v19_)
            d_786_v17_ = (d_786_v17_ if False else d_786_v17_)
            d_799_v26_: _dafny.Map
            d_799_v26_ = _dafny.Map({d_795_i3_: p0})
            d_800_v27_: _dafny.Set
            d_800_v27_ = _dafny.Set({p0})
            d_801_v28_: _dafny.Seq
            d_801_v28_ = _dafny.SeqWithoutIsStrInference([((d_799_v26_)[(d_785_v16_).cardinality] if ((d_785_v16_).cardinality) in (d_799_v26_) else p0), p0, p0, default__.fm1(p0, self.f18, len(d_800_v27_), globalState), p0])
            d_802_v29_: _dafny.Seq
            d_802_v29_ = _dafny.SeqWithoutIsStrInference([True, (d_801_v28_)[default__.safeIndex(d_784_v15_, len(d_801_v28_))]])
            d_803_v30_: _dafny.Map
            d_803_v30_ = _dafny.Map({_dafny.CodePoint('j'): len(d_788_v19_)})
            d_804_v31_: T0
            nw132_ = C0()
            nw132_.ctor__((d_796_v25_)[default__.safeIndex(962, (d_796_v25_).length(0))], d_803_v30_, self.f18, (self).f19)
            d_804_v31_ = nw132_
            d_805_v32_: _dafny.Set
            d_805_v32_ = _dafny.Set({d_804_v31_, d_804_v31_, d_804_v31_})
            rhs141_ = (d_784_v15_ if (d_801_v28_)[default__.safeIndex(d_784_v15_, len(d_801_v28_))] else d_795_i3_)
            rhs142_ = not(((d_805_v32_) | (d_805_v32_)).ispropersubset(d_805_v32_))
            lhs106_ = globalState
            lhs106_.f6 = rhs141_
            r3 = rhs142_
            (globalState).f6 = d_795_i3_
        d_806_v33_: _dafny.Array
        def lambda30_(d_807_p0_):
            def lambda31_(d_808_i5_):
                return d_807_p0_

            return lambda31_

        init15_ = lambda30_(p0)
        nw133_ = _dafny.Array(None, 9)
        for i0_15_ in range(nw133_.length(0)):
            nw133_[i0_15_] = init15_(i0_15_)
        d_806_v33_ = nw133_
        d_806_v33_ = d_806_v33_
        d_809_v34_: D1
        d_809_v34_ = D1_DC3(False)
        r2 = (not((d_809_v34_).cf2)) and (p0)
        d_810_v35_: C0
        nw134_ = C0()
        nw134_.ctor__((d_788_v19_ if p0 else d_788_v19_), default__.fm7(p0, d_784_v15_, globalState), _dafny.CodePoint('g'), (self).f19)
        d_810_v35_ = nw134_
        d_811_v36_: D0
        d_811_v36_ = D0_DC0((0) - (d_784_v15_))
        r0 = d_811_v36_
        d_812_v37_: _dafny.Seq
        d_812_v37_ = _dafny.SeqWithoutIsStrInference([d_784_v15_, d_784_v15_, d_784_v15_])
        r1 = (d_784_v15_) * ((len(d_812_v37_)) - (d_784_v15_))
        r2 = p0
        r3 = p0
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_813_v0_: _dafny.Array
        def lambda32_(d_814_p3_):
            def lambda33_(d_815_i0_):
                return (d_815_i0_) + (d_814_p3_)

            return lambda33_

        init16_ = lambda32_(p3)
        nw135_ = _dafny.Array(None, 8)
        for i0_16_ in range(nw135_.length(0)):
            nw135_[i0_16_] = init16_(i0_16_)
        d_813_v0_ = nw135_
        index144_ = default__.safeIndex(974, (d_813_v0_).length(0))
        (d_813_v0_)[index144_] = p3
        d_816_v1_: bool
        d_816_v1_ = True
        d_817_v2_: _dafny.Set
        d_817_v2_ = _dafny.Set({d_816_v1_})
        d_818_v3_: _dafny.Map
        d_818_v3_ = _dafny.Map({p3: d_817_v2_})
        d_819_v4_: D3
        d_819_v4_ = D3_DC9(d_816_v1_, default__.fm0(globalState))
        pat_let_tv16_ = p2
        pat_let_tv17_ = p2
        pat_let_tv18_ = p3
        index145_ = default__.safeIndex(974, (d_813_v0_).length(0))
        def lambda34_(source11_):
            if source11_.is_DC9:
                d_820___mcc_h0_ = source11_.cf11
                d_821___mcc_h1_ = source11_.cf12
                d_822_cf12_ = d_821___mcc_h1_
                d_823_cf11_ = d_820___mcc_h0_
                return d_822_cf12_
            elif source11_.is_DC10:
                d_824___mcc_h2_ = source11_.cf13
                d_825_cf13_ = d_824___mcc_h2_
                return default__.safeModuloInt(pat_let_tv16_, pat_let_tv17_)
            elif True:
                d_826___mcc_h3_ = source11_.cf10
                d_827_cf10_ = d_826___mcc_h3_
                return pat_let_tv18_

        rhs143_ = (0) - (lambda34_(d_819_v4_))
        rhs144_ = False
        rhs145_ = d_818_v3_
        lhs107_ = d_813_v0_
        lhs108_ = default__.safeIndex(974, (d_813_v0_).length(0))
        lhs107_[lhs108_] = rhs143_
        d_816_v1_ = rhs144_
        d_818_v3_ = rhs145_
        d_828_i1_: int
        d_828_i1_ = 0
        with _dafny.label("10"):
            while d_816_v1_:
                with _dafny.c_label("10"):
                    if (d_828_i1_) >= (100):
                        raise _dafny.Break("10")
                    d_828_i1_ = (d_828_i1_) + (1)
                    (globalState).f6 = ((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]) * (p3)
                    d_829_v5_: _dafny.Seq
                    d_829_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxtvrh"))
                    d_830_v6_: _dafny.Map
                    d_830_v6_ = _dafny.Map({self.f18: 308})
                    d_831_v7_: C0
                    nw136_ = C0()
                    nw136_.ctor__(d_829_v5_, d_830_v6_, self.f18, ((self).f19).set(d_829_v5_, (d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]))
                    d_831_v7_ = nw136_
                    index146_ = default__.safeIndex(974, (d_813_v0_).length(0))
                    (d_813_v0_)[index146_] = default__.fm0(globalState)
                    if ((p3) - (p2)) <= ((0) - ((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))])):
                        d_832_v8_: _dafny.Seq
                        d_832_v8_ = _dafny.SeqWithoutIsStrInference([(d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]])
                        (globalState).f1 = (d_832_v8_) + (_dafny.SeqWithoutIsStrInference([p2]))
                        d_833_v9_: _dafny.Map
                        d_833_v9_ = _dafny.Map({default__.safeModuloInt(p2, p1): _dafny.Map({p1: p0})})
                        d_833_v9_ = (d_833_v9_).set((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], _dafny.Map({default__.fm0(globalState): p0}))
                        d_834_v10_: D1
                        d_834_v10_ = D1_DC3(True)
                        d_835_v11_: _dafny.Map
                        d_835_v11_ = _dafny.Map({p3: d_834_v10_})
                        d_836_v12_: T0
                        nw137_ = C0()
                        nw137_.ctor__(default__.fm8(globalState), (d_831_v7_).f21, self.f18, (_dafny.Map({(d_831_v7_).f20: len(d_835_v11_)})).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnfqve")), (0) - (default__.fm0(globalState))))
                        d_836_v12_ = nw137_
                        d_837_v13_: _dafny.Array
                        def lambda35_(d_838_p0_):
                            def lambda36_(d_839_i2_):
                                return d_838_p0_

                            return lambda36_

                        init17_ = lambda35_(p0)
                        nw138_ = _dafny.Array(None, 23)
                        for i0_17_ in range(nw138_.length(0)):
                            nw138_[i0_17_] = init17_(i0_17_)
                        d_837_v13_ = nw138_
                        d_840_v14_: _dafny.Map
                        d_840_v14_ = _dafny.Map({d_836_v12_: d_837_v13_})
                        d_840_v14_ = (d_840_v14_).set(d_836_v12_, d_837_v13_)
                        d_841_v15_: _dafny.Seq
                        d_841_v15_ = _dafny.SeqWithoutIsStrInference([True, d_816_v1_])
                        d_842_v16_: _dafny.Map
                        d_842_v16_ = _dafny.Map({p1: True})
                        d_843_v17_: _dafny.Seq
                        d_843_v17_ = (d_841_v15_).set(default__.safeIndex(len(d_842_v16_), len(d_841_v15_)), p0)
                        d_844_v18_: _dafny.MultiSet
                        d_844_v18_ = _dafny.MultiSet([d_843_v17_])
                        d_845_v19_: _dafny.Seq
                        d_845_v19_ = _dafny.SeqWithoutIsStrInference([d_841_v15_])
                        d_846_v20_: _dafny.Seq
                        d_846_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_845_v19_)])
                        d_847_v21_: _dafny.Array
                        nw139_ = _dafny.Array(None, 11)
                        nw139_[int(0)] = (d_844_v18_).intersection(d_844_v18_)
                        nw139_[int(1)] = _dafny.MultiSet([d_843_v17_, d_843_v17_, d_843_v17_, d_841_v15_, _dafny.SeqWithoutIsStrInference([p0, p0, False, d_816_v1_, p0])])
                        nw139_[int(2)] = d_844_v18_
                        nw139_[int(3)] = (d_844_v18_) - (d_844_v18_)
                        nw139_[int(4)] = _dafny.MultiSet([d_843_v17_, d_843_v17_])
                        nw139_[int(5)] = default__.fm9(globalState)
                        nw139_[int(6)] = d_844_v18_
                        nw139_[int(7)] = (d_846_v20_)[default__.safeIndex((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], len(d_846_v20_))]
                        nw139_[int(8)] = (d_846_v20_)[default__.safeIndex((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], len(d_846_v20_))]
                        nw139_[int(9)] = d_844_v18_
                        nw139_[int(10)] = (_dafny.MultiSet([d_843_v17_])).set(d_843_v17_, default__.abs((self).fm4(d_816_v1_, self.f18, _dafny.CodePoint('u'), d_816_v1_, globalState)))
                        d_847_v21_ = nw139_
                        index147_ = default__.safeIndex(790, (d_847_v21_).length(0))
                        (d_847_v21_)[index147_] = d_844_v18_
                        d_848_v22_: _dafny.MultiSet
                        d_848_v22_ = _dafny.MultiSet([d_816_v1_])
                        index148_ = default__.safeIndex(790, (d_847_v21_).length(0))
                        (d_847_v21_)[index148_] = ((d_844_v18_).set(d_843_v17_, default__.abs((d_848_v22_).cardinality))).intersection(d_844_v18_)
                        d_849_v23_: _dafny.Set
                        d_849_v23_ = _dafny.Set({-192, (d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], (d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], (d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]})
                        d_816_v1_ = (d_849_v23_).issubset((d_849_v23_) | (d_849_v23_))
                    elif True:
                        d_850_v24_: _dafny.Map
                        d_850_v24_ = _dafny.Map({(default__.fm8(globalState)) + ((d_831_v7_).f20): d_816_v1_})
                        d_816_v1_ = ((d_850_v24_)[(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvox"))) + ((d_831_v7_).f20)] if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvox"))) + ((d_831_v7_).f20)) in (d_850_v24_) else d_816_v1_)
                        d_851_v25_: _dafny.Array
                        nw140_ = _dafny.Array(_dafny.Seq({}), 19)
                        d_851_v25_ = nw140_
                        d_852_v26_: _dafny.Set
                        d_852_v26_ = _dafny.Set({p3})
                        d_853_v27_: _dafny.Seq
                        d_853_v27_ = _dafny.SeqWithoutIsStrInference([p1, (d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], p1, len(d_852_v26_), p1])
                        index149_ = default__.safeIndex(462, (d_851_v25_).length(0))
                        (d_851_v25_)[index149_] = d_853_v27_
                        d_854_v28_: _dafny.Array
                        def lambda37_(d_855_v2_):
                            def lambda38_(d_856_i3_):
                                return d_855_v2_

                            return lambda38_

                        init18_ = lambda37_(d_817_v2_)
                        nw141_ = _dafny.Array(None, 22)
                        for i0_18_ in range(nw141_.length(0)):
                            nw141_[i0_18_] = init18_(i0_18_)
                        d_854_v28_ = nw141_
                        index150_ = default__.safeIndex(562, (d_854_v28_).length(0))
                        (d_854_v28_)[index150_] = _dafny.Set({d_816_v1_, d_816_v1_, d_816_v1_, p0, not(d_816_v1_)})
                        d_857_v29_: _dafny.Array
                        nw142_ = _dafny.Array(False, 19)
                        d_857_v29_ = nw142_
                        index151_ = default__.safeIndex(615, (d_857_v29_).length(0))
                        (d_857_v29_)[index151_] = d_816_v1_
                        d_858_v30_: _dafny.Seq
                        d_858_v30_ = _dafny.SeqWithoutIsStrInference([not(d_816_v1_)])
                        index152_ = default__.safeIndex(462, (d_851_v25_).length(0))
                        index153_ = default__.safeIndex(562, (d_854_v28_).length(0))
                        index154_ = default__.safeIndex(615, (d_857_v29_).length(0))
                        rhs146_ = (d_858_v30_)[default__.safeIndex(p1, len(d_858_v30_))]
                        rhs147_ = D3_DC9(p0, p3)
                        rhs148_ = ((d_853_v27_) + (d_853_v27_)) + (d_853_v27_)
                        rhs149_ = ((d_817_v2_) | (d_817_v2_)).intersection(_dafny.Set({p0}))
                        rhs150_ = ((d_853_v27_ if p0 else d_853_v27_)) < ((d_853_v27_) + (d_853_v27_))
                        lhs109_ = d_851_v25_
                        lhs110_ = default__.safeIndex(462, (d_851_v25_).length(0))
                        lhs111_ = d_854_v28_
                        lhs112_ = default__.safeIndex(562, (d_854_v28_).length(0))
                        lhs113_ = d_857_v29_
                        lhs114_ = default__.safeIndex(615, (d_857_v29_).length(0))
                        d_816_v1_ = rhs146_
                        d_819_v4_ = rhs147_
                        lhs109_[lhs110_] = rhs148_
                        lhs111_[lhs112_] = rhs149_
                        lhs113_[lhs114_] = rhs150_
                        index155_ = default__.safeIndex(974, (d_813_v0_).length(0))
                        (d_813_v0_)[index155_] = 600
                        index156_ = default__.safeIndex(615, (d_857_v29_).length(0))
                        (d_857_v29_)[index156_] = (d_858_v30_)[default__.safeIndex(p1, len(d_858_v30_))]
                        d_857_v29_ = d_857_v29_
                    pass
            pass
        d_859_v31_: _dafny.Map
        d_859_v31_ = _dafny.Map({p3: p0})
        d_860_v32_: _dafny.Seq
        d_860_v32_ = _dafny.SeqWithoutIsStrInference([d_859_v31_, d_859_v31_, d_859_v31_])
        index157_ = default__.safeIndex(974, (d_813_v0_).length(0))
        (d_813_v0_)[index157_] = (0) - (((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))] if not (p0) or (not(d_816_v1_)) else (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.Map({(_dafny.MultiSet([self.f18, self.f18, self.f18])).cardinality: d_816_v1_}) for d_861_i4_ in range(default__.abs(-147))])) + (d_860_v32_))).cardinality))
        d_862_v33_: _dafny.Map
        d_862_v33_ = _dafny.Map({(d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]: 135})
        hi8_ = len(d_862_v33_)
        for d_863_i5_ in range(p3, hi8_):
            d_864_v34_: D1
            d_864_v34_ = D1_DC5(p1)
            d_865_v35_: _dafny.Array
            nw143_ = _dafny.Array(None, 6)
            nw143_[int(0)] = d_864_v34_
            nw143_[int(1)] = (d_864_v34_ if d_816_v1_ else D1_DC5(717))
            nw143_[int(2)] = d_864_v34_
            nw143_[int(3)] = d_864_v34_
            nw143_[int(4)] = d_864_v34_
            nw143_[int(5)] = d_864_v34_
            d_865_v35_ = nw143_
            index158_ = default__.safeIndex(63, (d_865_v35_).length(0))
            (d_865_v35_)[index158_] = d_864_v34_
            index159_ = default__.safeIndex(63, (d_865_v35_).length(0))
            (d_865_v35_)[index159_] = D1_DC5(p2)
            d_816_v1_ = (p0) and (d_816_v1_)
            d_866_v36_: _dafny.Array
            nw144_ = _dafny.Array(False, 10)
            d_866_v36_ = nw144_
            index160_ = default__.safeIndex(991, (d_866_v36_).length(0))
            (d_866_v36_)[index160_] = not(p0)
            d_867_v37_: _dafny.MultiSet
            d_867_v37_ = _dafny.MultiSet([(0) - (p1), p3, p1])
            d_868_v38_: _dafny.Map
            d_868_v38_ = _dafny.Map({d_816_v1_: p0})
            d_869_v39_: T1
            nw145_ = C2()
            nw145_.ctor__(self.f18, (self).f19)
            d_869_v39_ = nw145_
            d_870_v40_: _dafny.Map
            d_870_v40_ = _dafny.Map({d_869_v39_: d_816_v1_})
            d_871_v41_: D1
            d_871_v41_ = D1_DC4(p1, len(d_870_v40_), d_816_v1_)
            index161_ = default__.safeIndex(991, (d_866_v36_).length(0))
            rhs151_ = ((self).fm3(d_816_v1_, ((d_867_v37_)[(0) - (p3)] if ((0) - (p3)) in (d_867_v37_) else len(d_868_v38_)), globalState)) and ((default__.fm1(p0, default__.fm10(globalState), d_863_i5_, globalState) if p0 else d_816_v1_))
            rhs152_ = d_817_v2_
            rhs153_ = (d_871_v41_).cf5
            lhs115_ = d_866_v36_
            lhs116_ = default__.safeIndex(991, (d_866_v36_).length(0))
            lhs115_[lhs116_] = rhs151_
            d_817_v2_ = rhs152_
            d_816_v1_ = rhs153_
            d_872_v42_: _dafny.Array
            nw146_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_872_v42_ = nw146_
            d_873_v43_: _dafny.Array
            nw147_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_873_v43_ = nw147_
            index162_ = default__.safeIndex(390, (d_872_v42_).length(0))
            (d_872_v42_)[index162_] = d_873_v43_
            d_874_v44_: _dafny.Seq
            d_874_v44_ = _dafny.SeqWithoutIsStrInference([d_816_v1_])
            d_875_v45_: _dafny.Set
            d_875_v45_ = _dafny.Set({((d_874_v44_) + (_dafny.SeqWithoutIsStrInference([(d_866_v36_)[default__.safeIndex(991, (d_866_v36_).length(0))], d_816_v1_, d_816_v1_, True, not(p0)]))).set(default__.safeIndex(p3, len((d_874_v44_) + (_dafny.SeqWithoutIsStrInference([(d_866_v36_)[default__.safeIndex(991, (d_866_v36_).length(0))], d_816_v1_, d_816_v1_, True, not(p0)])))), d_816_v1_)})
            index163_ = default__.safeIndex(390, (d_872_v42_).length(0))
            rhs154_ = (d_866_v36_)[default__.safeIndex(991, (d_866_v36_).length(0))]
            rhs155_ = d_873_v43_
            rhs156_ = d_875_v45_
            lhs117_ = d_872_v42_
            lhs118_ = default__.safeIndex(390, (d_872_v42_).length(0))
            d_816_v1_ = rhs154_
            lhs117_[lhs118_] = rhs155_
            d_875_v45_ = rhs156_
        d_876_v46_: C2
        nw148_ = C2()
        nw148_.ctor__(_dafny.CodePoint('q'), (self).f19)
        d_876_v46_ = nw148_
        d_877_v47_: _dafny.Map
        d_877_v47_ = _dafny.Map({d_876_v46_: p0})
        d_878_v48_: _dafny.Set
        d_878_v48_ = _dafny.Set({d_877_v47_, d_877_v47_, d_877_v47_, d_877_v47_})
        d_879_v49_: _dafny.Seq
        d_879_v49_ = _dafny.SeqWithoutIsStrInference([d_878_v48_])
        d_880_v50_: _dafny.MultiSet
        d_880_v50_ = _dafny.MultiSet([p0])
        if not(((d_879_v49_)[default__.safeIndex((d_880_v50_).cardinality, len(d_879_v49_))]) != ((d_878_v48_) - (d_878_v48_))):
            d_881_v51_: _dafny.Seq
            d_881_v51_ = _dafny.SeqWithoutIsStrInference([p0])
            d_882_v52_: _dafny.Array
            nw149_ = _dafny.Array(None, 29)
            nw149_[int(0)] = p0
            nw149_[int(1)] = p0
            nw149_[int(2)] = True
            nw149_[int(3)] = p0
            nw149_[int(4)] = d_816_v1_
            nw149_[int(5)] = d_816_v1_
            nw149_[int(6)] = d_816_v1_
            nw149_[int(7)] = (d_881_v51_)[default__.safeIndex(p1, len(d_881_v51_))]
            nw149_[int(8)] = d_816_v1_
            nw149_[int(9)] = p0
            nw149_[int(10)] = d_816_v1_
            nw149_[int(11)] = p0
            nw149_[int(12)] = d_816_v1_
            nw149_[int(13)] = d_816_v1_
            nw149_[int(14)] = d_816_v1_
            nw149_[int(15)] = d_816_v1_
            nw149_[int(16)] = p0
            nw149_[int(17)] = p0
            nw149_[int(18)] = d_816_v1_
            nw149_[int(19)] = True
            nw149_[int(20)] = d_816_v1_
            nw149_[int(21)] = p0
            nw149_[int(22)] = d_816_v1_
            nw149_[int(23)] = p0
            nw149_[int(24)] = p0
            nw149_[int(25)] = d_816_v1_
            nw149_[int(26)] = (d_881_v51_)[default__.safeIndex((0) - (len(_dafny.Map({d_816_v1_: d_813_v0_}))), len(d_881_v51_))]
            nw149_[int(27)] = p0
            nw149_[int(28)] = p0
            d_882_v52_ = nw149_
            d_883_v53_: _dafny.Map
            d_883_v53_ = _dafny.Map({p0: d_882_v52_})
            d_884_v54_: _dafny.Array
            nw150_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
            d_884_v54_ = nw150_
            d_885_v55_: _dafny.Map
            d_885_v55_ = _dafny.Map({((d_883_v53_)[d_816_v1_] if (d_816_v1_) in (d_883_v53_) else d_882_v52_): d_884_v54_})
            d_886_v56_: _dafny.Map
            d_886_v56_ = _dafny.Map({((d_885_v55_)[d_882_v52_] if (d_882_v52_) in (d_885_v55_) else d_884_v54_): p0})
            d_886_v56_ = (d_886_v56_).set(d_884_v54_, p0)
            d_887_v57_: _dafny.Map
            d_887_v57_ = _dafny.Map({p0: p0})
            d_888_v58_: _dafny.Seq
            d_888_v58_ = _dafny.SeqWithoutIsStrInference([d_887_v57_])
            d_889_v59_: _dafny.MultiSet
            d_889_v59_ = _dafny.MultiSet([d_887_v57_, d_887_v57_, d_887_v57_, (d_887_v57_) | ((d_888_v58_)[default__.safeIndex(p3, len(d_888_v58_))]), default__.fm23(p1, d_816_v1_, globalState)])
            d_889_v59_ = d_889_v59_
            d_890_v60_: C1
            nw151_ = C1()
            nw151_.ctor__((p2) + (p2), p0, self.f18, default__.fm18((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], globalState))
            d_890_v60_ = nw151_
            d_891_v61_: _dafny.Seq
            d_891_v61_ = _dafny.SeqWithoutIsStrInference([p2])
            d_892_v62_: _dafny.Map
            d_892_v62_ = _dafny.Map({d_891_v61_: d_816_v1_})
            d_816_v1_ = not (d_816_v1_) or ((True) or (((d_892_v62_)[_dafny.SeqWithoutIsStrInference([131 for d_893_i6_ in range(default__.abs(-491))])] if (_dafny.SeqWithoutIsStrInference([131 for d_894_i6_ in range(default__.abs(-491))])) in (d_892_v62_) else False)))
            index164_ = default__.safeIndex(974, (d_813_v0_).length(0))
            (d_813_v0_)[index164_] = default__.safeModuloInt(len((d_881_v51_).set(default__.safeIndex(len(d_817_v2_), len(d_881_v51_)), not(d_816_v1_))), len(d_859_v31_))
        elif True:
            rhs157_ = p2
            rhs158_ = p2
            rhs159_ = (True if (d_817_v2_).isdisjoint(d_817_v2_) else not(False))
            lhs119_ = globalState
            lhs119_.f6 = rhs157_
            r0 = rhs158_
            d_816_v1_ = rhs159_
            d_895_v63_: _dafny.Seq
            d_895_v63_ = _dafny.SeqWithoutIsStrInference([d_813_v0_, d_813_v0_, d_813_v0_, d_813_v0_, d_813_v0_])
            d_896_v64_: D1
            d_896_v64_ = D1_DC4((0) - ((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]), (d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], d_816_v1_)
            d_897_v65_: _dafny.Map
            d_897_v65_ = _dafny.Map({p0: (d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]})
            d_898_v66_: _dafny.MultiSet
            d_898_v66_ = _dafny.MultiSet([p1])
            d_899_v67_: _dafny.Seq
            d_899_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wyt"))
            d_900_v68_: _dafny.Array
            nw152_ = _dafny.Array(None, 28)
            nw152_[int(0)] = p0
            nw152_[int(1)] = d_816_v1_
            nw152_[int(2)] = d_816_v1_
            nw152_[int(3)] = d_816_v1_
            nw152_[int(4)] = not(True)
            nw152_[int(5)] = p0
            nw152_[int(6)] = d_816_v1_
            nw152_[int(7)] = (d_813_v0_) not in (d_895_v63_)
            nw152_[int(8)] = d_816_v1_
            nw152_[int(9)] = not(p0)
            nw152_[int(10)] = ((d_859_v31_)[(d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]] if ((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))]) in (d_859_v31_) else (d_896_v64_).cf5)
            nw152_[int(11)] = (p0) in (d_897_v65_)
            nw152_[int(12)] = d_816_v1_
            nw152_[int(13)] = (d_898_v66_).issubset(d_898_v66_)
            nw152_[int(14)] = d_816_v1_
            nw152_[int(15)] = ((d_819_v4_).cf11) == (p0)
            nw152_[int(16)] = ((d_859_v31_)[p2] if (p2) in (d_859_v31_) else not(p0))
            nw152_[int(17)] = True
            nw152_[int(18)] = True
            nw152_[int(19)] = not (d_816_v1_) or (p0)
            nw152_[int(20)] = not(p0)
            nw152_[int(21)] = (d_899_v67_) != (d_899_v67_)
            nw152_[int(22)] = p0
            nw152_[int(23)] = (False if d_816_v1_ else d_816_v1_)
            nw152_[int(24)] = d_816_v1_
            nw152_[int(25)] = not(True)
            nw152_[int(26)] = (p0 if not(d_816_v1_) else d_816_v1_)
            nw152_[int(27)] = p0
            d_900_v68_ = nw152_
            index165_ = default__.safeIndex(459, (d_900_v68_).length(0))
            (d_900_v68_)[index165_] = d_816_v1_
            d_901_v69_: _dafny.Map
            d_901_v69_ = _dafny.Map({d_880_v50_: p2})
            d_902_v71_: _dafny.Set
            d_902_v71_ = _dafny.Set({d_880_v50_, d_880_v50_})
            d_903_v72_: _dafny.Map
            d_903_v72_ = _dafny.Map({d_880_v50_: p0})
            index166_ = default__.safeIndex(459, (d_900_v68_).length(0))
            def iife74_():
                coll38_ = _dafny.Set()
                compr_38_: _dafny.MultiSet
                for compr_38_ in ((d_903_v72_).set(d_880_v50_, p0)).keys.Elements:
                    d_904_v73_: _dafny.MultiSet = compr_38_
                    if (d_904_v73_) in ((d_903_v72_).set(d_880_v50_, p0)):
                        coll38_ = coll38_.union(_dafny.Set([d_904_v73_]))
                return _dafny.Set(coll38_)
            def iife75_():
                coll39_ = _dafny.Set()
                compr_39_: _dafny.MultiSet
                for compr_39_ in (d_901_v69_).keys.Elements:
                    d_905_v70_: _dafny.MultiSet = compr_39_
                    if (d_905_v70_) in (d_901_v69_):
                        coll39_ = coll39_.union(_dafny.Set([d_905_v70_]))
                return _dafny.Set(coll39_)
            rhs160_ = (iife74_()
            ).ispropersubset((iife75_()
            ) | (d_902_v71_))
            rhs161_ = d_816_v1_
            rhs162_ = d_899_v67_
            lhs120_ = d_900_v68_
            lhs121_ = default__.safeIndex(459, (d_900_v68_).length(0))
            lhs120_[lhs121_] = rhs160_
            d_816_v1_ = rhs161_
            d_899_v67_ = rhs162_
            if True:
                d_816_v1_ = not ((d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))]) or ((False if False else p0))
                d_862_v33_ = d_862_v33_
                d_906_v75_: _dafny.Seq
                d_906_v75_ = _dafny.SeqWithoutIsStrInference([(d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))]])
                d_907_v76_: _dafny.Seq
                d_907_v76_ = d_906_v75_
                d_908_v77_: _dafny.Seq
                d_908_v77_ = _dafny.SeqWithoutIsStrInference([d_907_v76_, d_906_v75_])
                d_909_v78_: _dafny.Seq
                d_909_v78_ = _dafny.SeqWithoutIsStrInference([d_908_v77_])
                d_910_v79_: _dafny.Map
                def iife76_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(304, -945):
                        d_911_v74_: int = compr_40_
                        if ((304) <= (d_911_v74_)) and ((d_911_v74_) < (-945)):
                            coll40_[(d_911_v74_) - (p1)] = p0
                    return _dafny.Map(coll40_)
                d_910_v79_ = _dafny.Map({iife76_()
                : (d_909_v78_)[default__.safeIndex(p3, len(d_909_v78_))]})
                d_912_v80_: _dafny.Map
                d_912_v80_ = _dafny.Map({d_817_v2_: d_910_v79_})
                d_912_v80_ = (d_912_v80_).set((d_817_v2_ if (d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))] else _dafny.Set({(d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))], p0, not((d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))])})), d_910_v79_)
                index167_ = default__.safeIndex(459, (d_900_v68_).length(0))
                (d_900_v68_)[index167_] = p0
                r0 = p3
            elif True:
                d_913_v81_: D3
                d_913_v81_ = D3_DC8(_dafny.CodePoint('d'))
                d_914_v82_: _dafny.Map
                d_914_v82_ = _dafny.Map({(d_913_v81_).cf10: p3})
                d_915_v83_: C0
                nw153_ = C0()
                nw153_.ctor__(d_899_v67_, d_914_v82_, self.f18, ((self).f19) | ((self).f19))
                d_915_v83_ = nw153_
                d_916_v84_: C0
                nw154_ = C0()
                nw154_.ctor__(default__.fm8(globalState), (d_915_v83_).f21, self.f18, (self).f19)
                d_916_v84_ = nw154_
                d_916_v84_ = d_915_v83_
                d_917_v85_: D1
                d_917_v85_ = D1_DC6(d_816_v1_, (d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))])
                d_918_v86_: _dafny.Map
                d_918_v86_ = _dafny.Map({p3: d_917_v85_})
                d_918_v86_ = (d_918_v86_).set((d_813_v0_)[default__.safeIndex(974, (d_813_v0_).length(0))], d_917_v85_)
                d_919_v87_: C0
                nw155_ = C0()
                nw155_.ctor__(((d_916_v84_).f20 if not(not(True)) else d_899_v67_), (d_914_v82_) | (d_914_v82_), self.f18, (self).f19)
                d_919_v87_ = nw155_
                (globalState).f6 = default__.safeModuloInt((p3) - (p1), 259)
            d_920_v88_: _dafny.Map
            d_920_v88_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference([d_816_v1_, p0])})
            d_921_v89_: D3
            d_921_v89_ = D3_DC10(not(p0))
            d_922_v90_: _dafny.Map
            d_922_v90_ = _dafny.Map({d_920_v88_: d_921_v89_})
            d_922_v90_ = (d_922_v90_).set(d_920_v88_, default__.fm24(globalState))
            d_923_v92_: _dafny.Map
            d_923_v92_ = _dafny.Map({self.f18: d_816_v1_})
            d_924_v93_: D6
            def iife77_():
                coll41_ = _dafny.Map()
                compr_41_: str
                for compr_41_ in (d_923_v92_).keys.Elements:
                    d_925_v91_: str = compr_41_
                    if (d_925_v91_) in (d_923_v92_):
                        coll41_[d_925_v91_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))
                return _dafny.Map(coll41_)
            d_924_v93_ = D6_DC15(iife77_()
)
            d_926_v94_: D6
            d_926_v94_ = D6_DC17(d_924_v93_)
            d_927_v95_: _dafny.Map
            d_927_v95_ = _dafny.Map({d_926_v94_: (d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))]})
            d_928_v96_: _dafny.Seq
            d_928_v96_ = _dafny.SeqWithoutIsStrInference([d_927_v95_])
            d_929_v97_: _dafny.Array
            nw156_ = _dafny.Array(None, 29)
            nw156_[int(0)] = d_927_v95_
            nw156_[int(1)] = (d_927_v95_) | ((d_927_v95_).set(d_926_v94_, p0))
            nw156_[int(2)] = d_927_v95_
            nw156_[int(3)] = d_927_v95_
            nw156_[int(4)] = (d_927_v95_) | (d_927_v95_)
            nw156_[int(5)] = d_927_v95_
            nw156_[int(6)] = d_927_v95_
            nw156_[int(7)] = d_927_v95_
            nw156_[int(8)] = d_927_v95_
            nw156_[int(9)] = d_927_v95_
            nw156_[int(10)] = (d_928_v96_)[default__.safeIndex(default__.fm0(globalState), len(d_928_v96_))]
            nw156_[int(11)] = d_927_v95_
            nw156_[int(12)] = d_927_v95_
            nw156_[int(13)] = d_927_v95_
            nw156_[int(14)] = d_927_v95_
            nw156_[int(15)] = d_927_v95_
            nw156_[int(16)] = d_927_v95_
            nw156_[int(17)] = (d_927_v95_ if d_816_v1_ else d_927_v95_)
            nw156_[int(18)] = d_927_v95_
            nw156_[int(19)] = d_927_v95_
            nw156_[int(20)] = (d_927_v95_).set(d_926_v94_, (self).fm3(not((d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))]), ((d_880_v50_)[(d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))]] if ((d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))]) in (d_880_v50_) else len(d_899_v67_)), globalState))
            nw156_[int(21)] = d_927_v95_
            nw156_[int(22)] = d_927_v95_
            nw156_[int(23)] = d_927_v95_
            nw156_[int(24)] = (d_927_v95_) | ((d_927_v95_).set(d_926_v94_, (d_900_v68_)[default__.safeIndex(459, (d_900_v68_).length(0))]))
            nw156_[int(25)] = (d_927_v95_) | (d_927_v95_)
            nw156_[int(26)] = d_927_v95_
            nw156_[int(27)] = _dafny.Map({D6_DC17(D6_DC15(_dafny.Map({self.f18: p3}))): d_816_v1_})
            nw156_[int(28)] = d_927_v95_
            d_929_v97_ = nw156_
            d_929_v97_ = d_929_v97_
        d_930_v98_: _dafny.Seq
        d_930_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dv"))
        d_931_v99_: D9
        d_931_v99_ = D9_DC22(d_930_v98_)
        d_932_v100_: _dafny.Array
        nw157_ = _dafny.Array(None, 16)
        nw157_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_933_i8_ in range(default__.abs(670))])
        nw157_[int(1)] = d_930_v98_
        nw157_[int(2)] = (d_931_v99_).cf27
        nw157_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ude"))
        nw157_[int(4)] = d_930_v98_
        nw157_[int(5)] = d_930_v98_
        nw157_[int(6)] = d_930_v98_
        nw157_[int(7)] = (d_930_v98_) + (d_930_v98_)
        nw157_[int(8)] = _dafny.SeqWithoutIsStrInference([self.f18 for d_934_i9_ in range(default__.abs(168))])
        nw157_[int(9)] = (d_930_v98_) + (d_930_v98_)
        nw157_[int(10)] = d_930_v98_
        nw157_[int(11)] = (d_931_v99_).cf27
        nw157_[int(12)] = (d_930_v98_) + (d_930_v98_)
        nw157_[int(13)] = d_930_v98_
        nw157_[int(14)] = d_930_v98_
        nw157_[int(15)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))) + (d_930_v98_)
        d_932_v100_ = nw157_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_932_v100_).length(0)):
            d_935_i7_: int = guard_loop_2_
            if (True) and (((0) <= (d_935_i7_)) and ((d_935_i7_) < ((d_932_v100_).length(0)))):
                (d_932_v100_)[(d_935_i7_)] = d_930_v98_
        r0 = p3
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_936_v0_: _dafny.Seq
        d_936_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfbxyfb"))
        d_937_v1_: _dafny.Map
        d_937_v1_ = _dafny.Map({(p2).cardinality: p1})
        d_938_v2_: _dafny.Seq
        d_938_v2_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1, p1])), p0, p0, p0, len(d_936_v0_)])
        hi9_ = (len(d_937_v1_)) - ((d_938_v2_)[default__.safeIndex(764, len(d_938_v2_))])
        for d_939_i0_ in range((0) - (len(d_936_v0_)), hi9_):
            d_936_v0_ = d_936_v0_
            d_940_v3_: _dafny.Array
            nw158_ = _dafny.Array(_dafny.Map({}), 6)
            d_940_v3_ = nw158_
            d_941_v4_: _dafny.Map
            d_941_v4_ = _dafny.Map({p0: (p2).cardinality})
            index168_ = default__.safeIndex(973, (d_940_v3_).length(0))
            (d_940_v3_)[index168_] = d_941_v4_
            index169_ = default__.safeIndex(973, (d_940_v3_).length(0))
            (d_940_v3_)[index169_] = (d_941_v4_ if (p1) and (p1) else d_941_v4_)
            d_942_v5_: _dafny.Array
            d_943_v6_: bool
            d_944_v7_: int
            out24_: _dafny.Array
            out25_: bool
            out26_: int
            out24_, out25_, out26_ = default__.m0((self).fm4(p1, _dafny.CodePoint('e'), self.f18, p1, globalState), p1, d_939_i0_, globalState)
            d_942_v5_ = out24_
            d_943_v6_ = out25_
            d_944_v7_ = out26_
            d_945_v8_: _dafny.MultiSet
            d_945_v8_ = _dafny.MultiSet([p2, p2, _dafny.MultiSet([d_944_v7_, default__.fm0(globalState), (0) - (d_944_v7_), p0, len(d_936_v0_)]), (_dafny.MultiSet([d_939_i0_])) | (p2)])
            d_946_v9_: _dafny.Set
            d_946_v9_ = _dafny.Set({default__.fm0(globalState), len(d_936_v0_), len(d_936_v0_)})
            rhs163_ = d_939_i0_
            rhs164_ = ((d_945_v8_)[p2] if (p2) in (d_945_v8_) else d_944_v7_)
            rhs165_ = default__.safeDivisionInt(len(d_946_v9_), default__.fm0(globalState))
            rhs166_ = default__.fm8(globalState)
            lhs122_ = globalState
            lhs123_ = globalState
            r0 = rhs163_
            lhs122_.f6 = rhs164_
            lhs123_.f6 = rhs165_
            d_936_v0_ = rhs166_
        d_947_v10_: C2
        nw159_ = C2()
        nw159_.ctor__(self.f18, (self).f19)
        d_947_v10_ = nw159_
        d_948_v11_: _dafny.Seq
        d_948_v11_ = _dafny.SeqWithoutIsStrInference([d_947_v10_])
        d_949_v12_: _dafny.Map
        d_949_v12_ = _dafny.Map({p1: (d_948_v11_ if p1 else d_948_v11_)})
        d_949_v12_ = ((d_949_v12_) | (_dafny.Map({False: d_948_v11_}))) | (_dafny.Map({p1: d_948_v11_}))
        d_950_v13_: _dafny.MultiSet
        d_950_v13_ = _dafny.MultiSet([False, False, not(p1), p1, (d_936_v0_) <= (d_936_v0_)])
        d_950_v13_ = d_950_v13_
        d_951_v14_: bool
        d_951_v14_ = True
        d_951_v14_ = not(d_951_v14_)
        d_951_v14_ = not(p1)
        d_952_i1_: int
        d_952_i1_ = 0
        with _dafny.label("11"):
            while d_951_v14_:
                with _dafny.c_label("11"):
                    if (d_952_i1_) >= (100):
                        raise _dafny.Break("11")
                    d_952_i1_ = (d_952_i1_) + (1)
                    d_951_v14_ = p1
                    (self).f18 = (d_936_v0_)[default__.safeIndex(len((d_936_v0_) + (d_936_v0_)), len(d_936_v0_))]
                    d_951_v14_ = d_951_v14_
                    d_953_v15_: _dafny.Set
                    d_953_v15_ = _dafny.Set({(self).fm4(p1, (d_936_v0_)[default__.safeIndex(p0, len(d_936_v0_))], self.f18, d_951_v14_, globalState)})
                    d_951_v14_ = (d_953_v15_).issubset(d_953_v15_)
                    pass
            pass
        r0 = p0
        return r0

