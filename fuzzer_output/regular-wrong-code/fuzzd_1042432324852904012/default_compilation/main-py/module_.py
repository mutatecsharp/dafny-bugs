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
        source0_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - (-702))])
        d_0___mcc_h0_ = source0_
        d_1_cf0_ = d_0___mcc_h0_
        return True

    @staticmethod
    def fm6(p0, p1, globalState):
        if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2_i0_ in range(default__.abs(790))]))) >= (len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urqk")): 865}))):
            return True
        elif True:
            return True

    @staticmethod
    def fm7(p0, p1, globalState):
        if not(not((False) == (True))):
            return _dafny.Set({True})
        elif True:
            return (_dafny.Set({False, False, False, not(True)})).intersection(_dafny.Set({True, (D16_DC45(False)).cf61, False}))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True]))) + (_dafny.SeqWithoutIsStrInference([False, True]))

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return ((_dafny.Map({True: _dafny.Map({592: True})}))) | ((_dafny.Map({True: _dafny.Map({(_dafny.MultiSet([227])).cardinality: True})})))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([128 for d_3_i0_ in range(default__.abs(732))])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([411, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsxnajks")))}))]))), 437]))

    @staticmethod
    def fm11(p0, p1, globalState):
        if (_dafny.MultiSet([_dafny.CodePoint('x')])).isdisjoint(_dafny.MultiSet([_dafny.CodePoint('f'), _dafny.CodePoint('r'), _dafny.CodePoint('n'), (D13_DC35(_dafny.CodePoint('c'))).cf45, _dafny.CodePoint('w')])):
            return (_dafny.MultiSet([405, 994])).intersection(_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality]))
        elif True:
            return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qo")))])

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrq"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_4_i0_ in range(default__.abs(337))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhpmme")))

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return (_dafny.Map({not(False): False}))

    @staticmethod
    def fm17(p0, globalState):
        return _dafny.CodePoint('w')

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return (_dafny.Set({True})) | (_dafny.Set({False}))

    @staticmethod
    def fm19(p0, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: (D11_DC30(False, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, not(False), True, True])))).cf40}))) | ((_dafny.Map({True: True})) | (_dafny.Map({True: not(True)})))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([D7_DC16(_dafny.Set({False})), D7_DC16(_dafny.Set({True, False})), D7_DC16(_dafny.Set({not(False), True, not(True), False}))])) + (_dafny.SeqWithoutIsStrInference([D7_DC16(_dafny.Set({False}))]))

    @staticmethod
    def fm21(globalState):
        return _dafny.SeqWithoutIsStrInference([not(not(True))])

    @staticmethod
    def fm22(p0, globalState):
        source1_ = not(not(True))
        d_5___mcc_h0_ = source1_
        d_6_cf1_ = d_5___mcc_h0_
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaete"))): len(_dafny.SeqWithoutIsStrInference([687]))})

    @staticmethod
    def fm23(globalState):
        return (_dafny.MultiSet([not(False), False, False])) - (_dafny.MultiSet([False, not(not(False)), not(False)]))

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return D4_DC11(-410, not((_dafny.Set({False})).ispropersubset(_dafny.Set({True, False}))), (not(False)) == (not(False)))

    @staticmethod
    def fm25(globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(528, -787):
                d_7_v0_: int = compr_0_
                if ((528) <= (d_7_v0_)) and ((d_7_v0_) < (-787)):
                    coll0_ = coll0_.union(_dafny.Set([(d_7_v0_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({298})])))]))
            return _dafny.Set(coll0_)
        return (iife0_()
        ).intersection(_dafny.Set({949}))

    @staticmethod
    def fm26(p0, globalState):
        if False:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fix"))
        elif True:
            return (D6_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))).cf21

    @staticmethod
    def fm27(p0, globalState):
        return (_dafny.Map({True: 634})) | ((_dafny.Map({False: 945}) if (D4_DC10(_dafny.MultiSet([False]), True)).cf16 else _dafny.Map({True: -925})))

    @staticmethod
    def fm28(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([277]))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False, True, not(True), False, False]): False}))]))

    @staticmethod
    def fm29(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(6, 99):
                d_8_v0_: int = compr_1_
                if ((6) <= (d_8_v0_)) and ((d_8_v0_) < (99)):
                    coll1_[(d_8_v0_) + ((_dafny.MultiSet([False])).cardinality)] = True
            return _dafny.Map(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(46, 717):
                d_9_v1_: int = compr_2_
                if ((46) <= (d_9_v1_)) and ((d_9_v1_) < (717)):
                    coll2_[default__.safeModuloInt(d_9_v1_, 132)] = (False)
            return _dafny.Map(coll2_)
        return (_dafny.Map({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([-614, (0) - (len(iife1_()
        )), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnsrwmjof")))])))): True})) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilip"))): True})) | (iife2_()
        ))

    @staticmethod
    def fm30(globalState):
        source2_ = D12_DC33()
        if source2_.is_DC33:
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('d')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('v')]) for d_10_i0_ in range(default__.abs(349))]))
        elif source2_.is_DC32:
            d_11___mcc_h0_ = source2_.cf43
            d_12_cf43_ = d_11___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('d')])])
        elif True:
            d_13___mcc_h1_ = source2_.cf44
            d_14_cf44_ = d_13___mcc_h1_
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('g'), _dafny.CodePoint('y'), _dafny.CodePoint('d'), _dafny.CodePoint('u'), _dafny.CodePoint('j')]), _dafny.MultiSet([_dafny.CodePoint('t'), _dafny.CodePoint('u')]), _dafny.MultiSet([_dafny.CodePoint('i'), _dafny.CodePoint('b')])])

    @staticmethod
    def fm31(globalState):
        return (_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('n')])), _dafny.MultiSet([_dafny.CodePoint('g'), _dafny.CodePoint('w')]), _dafny.MultiSet([_dafny.CodePoint('q')])])) - (_dafny.MultiSet([_dafny.MultiSet([_dafny.CodePoint('f'), _dafny.CodePoint('s'), _dafny.CodePoint('l')])]))

    @staticmethod
    def fm32(globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('b')])) - (_dafny.MultiSet([_dafny.CodePoint('k')]))) - (_dafny.MultiSet([_dafny.CodePoint('n')]))

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qe")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmhgftehd")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yatjnok")), (D6_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebufvm")))).cf21, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytlonbky"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lw"))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oojfql"))]))

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return D3_DC7(_dafny.SeqWithoutIsStrInference([_dafny.Set({not(False)}) for d_15_i0_ in range(default__.abs(760))]), (_dafny.CodePoint('p')) not in (_dafny.Map({_dafny.CodePoint('b'): len(_dafny.SeqWithoutIsStrInference([True]))})))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        return _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([_dafny.Map({945: True}) for d_16_i0_ in range(default__.abs(36))]))) >= (703): (_dafny.Set({not(False)})) - ((D7_DC16(_dafny.Set({True}))).cf24)})

    @staticmethod
    def fm36(globalState):
        return D8_DC20()

    @staticmethod
    def fm37(p0, globalState):
        return ((_dafny.MultiSet([730, 70, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thphighp")))]))).cardinality, -40, 563])) | (_dafny.MultiSet([210, 532, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, not(True)])) for d_17_i0_ in range(default__.abs(302))])), 561]))) | (_dafny.MultiSet([len(_dafny.Map({163: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({290, len(_dafny.Map({False: 320}))}))]))}))]))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(-678, -417):
                d_18_v0_: int = compr_3_
                if ((-678) <= (d_18_v0_)) and ((d_18_v0_) < (-417)):
                    coll3_[(d_18_v0_) + (len(_dafny.Map({(0) - (len(_dafny.Set({not(True)}))): not((D7_DC17(True)).cf25)})))] = False
            return _dafny.Map(coll3_)
        return (_dafny.Map({iife3_()
        : 235})) | ((_dafny.Map({_dafny.Map({len(_dafny.Set({len(_dafny.Map({998: 447}))})): not(True)}): 951})) | (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))): False}): 709})))

    @staticmethod
    def fm39(globalState):
        return D10_DC25(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False]))]))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return _dafny.Map({_dafny.CodePoint('w'): (_dafny.Map({not(True): 457})) | (_dafny.Map({True: 619}))})

    @staticmethod
    def fm41(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])).Elements:
                d_19_v0_: _dafny.Seq = compr_4_
                if (d_19_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])):
                    coll4_ = coll4_.union(_dafny.Set([d_19_v0_]))
            return _dafny.Set(coll4_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([(D16_DC45(not(True))).cf61])})).intersection(iife4_()
        )

    @staticmethod
    def fm42(p0, p1, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(410, 632):
                d_20_v0_: int = compr_5_
                if ((410) <= (d_20_v0_)) and ((d_20_v0_) < (632)):
                    coll5_[(d_20_v0_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])))] = 78
            return _dafny.Map(coll5_)
        if not(((0) - (len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvsgx")))): False})))) >= (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tanewc"))): (0) - (len(_dafny.SeqWithoutIsStrInference([iife5_()
        ])))})))):
            def iife6_():
                def iife8_():
                    coll8_ = _dafny.Set()
                    compr_8_: int
                    for compr_8_ in (_dafny.MultiSet([879, 452])).Elements:
                        d_23_v2_: int = compr_8_
                        if (d_23_v2_) in (_dafny.MultiSet([879, 452])):
                            coll8_ = coll8_.union(_dafny.Set([default__.safeModuloInt(d_23_v2_, -222)]))
                    return _dafny.Set(coll8_)
                coll6_ = _dafny.Map()
                def iife7_():
                    coll7_ = _dafny.Set()
                    compr_7_: int
                    for compr_7_ in (_dafny.MultiSet([879, 452])).Elements:
                        d_21_v2_: int = compr_7_
                        if (d_21_v2_) in (_dafny.MultiSet([879, 452])):
                            coll7_ = coll7_.union(_dafny.Set([default__.safeModuloInt(d_21_v2_, -222)]))
                    return _dafny.Set(coll7_)
                compr_6_: int
                for compr_6_ in (iife7_()
                ).Elements:
                    d_22_v1_: int = compr_6_
                    if (d_22_v1_) in (iife8_()
                    ):
                        coll6_[default__.safeModuloInt(d_22_v1_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))]))).cardinality)] = True
                return _dafny.Map(coll6_)
            return D3_DC6((0) - (len(iife6_()
)))
        elif True:
            def iife9_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in (_dafny.MultiSet([162])).Elements:
                    d_24_v3_: int = compr_9_
                    if (d_24_v3_) in (_dafny.MultiSet([162])):
                        coll9_ = coll9_.union(_dafny.Set([(d_24_v3_) * ((0) - (-187))]))
                return _dafny.Set(coll9_)
            return D3_DC6(len(iife9_()
))

    @staticmethod
    def Main(noArgsParameter__):
        d_25_v0_: _dafny.Seq
        d_25_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrqsc"))
        d_26_v1_: int
        d_26_v1_ = -60
        d_27_v2_: _dafny.Array
        nw0_ = _dafny.Array(False, 9)
        d_27_v2_ = nw0_
        d_28_v3_: _dafny.Map
        d_28_v3_ = _dafny.Map({d_26_v1_: d_27_v2_})
        d_29_v4_: _dafny.Set
        d_29_v4_ = _dafny.Set({d_26_v1_})
        d_30_v5_: _dafny.Map
        d_30_v5_ = _dafny.Map({d_26_v1_: d_29_v4_})
        d_31_v6_: _dafny.Map
        d_31_v6_ = _dafny.Map({False: len(((d_30_v5_)[d_26_v1_] if (d_26_v1_) in (d_30_v5_) else d_29_v4_))})
        d_32_v7_: bool
        d_32_v7_ = True
        d_33_v8_: _dafny.Map
        d_33_v8_ = _dafny.Map({d_26_v1_: d_26_v1_})
        d_34_v9_: _dafny.Set
        d_34_v9_ = _dafny.Set({d_32_v7_})
        d_35_v10_: _dafny.Map
        d_35_v10_ = _dafny.Map({len(d_34_v9_): d_32_v7_})
        d_36_v12_: _dafny.Map
        d_36_v12_ = _dafny.Map({not(d_32_v7_): d_33_v8_})
        d_37_v13_: _dafny.Map
        d_37_v13_ = _dafny.Map({(0) - (d_26_v1_): d_33_v8_})
        d_38_v16_: _dafny.Seq
        d_38_v16_ = _dafny.SeqWithoutIsStrInference([d_26_v1_, d_26_v1_])
        d_39_v17_: _dafny.Map
        d_39_v17_ = _dafny.Map({len(d_25_v0_): d_38_v16_})
        d_40_v18_: _dafny.Array
        nw1_ = _dafny.Array(None, 19)
        nw1_[int(0)] = (d_33_v8_).set(len((_dafny.SeqWithoutIsStrInference([d_32_v7_, d_32_v7_, True])).set(default__.safeIndex(d_26_v1_, len(_dafny.SeqWithoutIsStrInference([d_32_v7_, d_32_v7_, True]))), d_32_v7_)), len(d_35_v10_))
        nw1_[int(1)] = d_33_v8_
        nw1_[int(2)] = d_33_v8_
        nw1_[int(3)] = d_33_v8_
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(419, 570):
                d_41_v11_: int = compr_10_
                if ((419) <= (d_41_v11_)) and ((d_41_v11_) < (570)):
                    coll10_[default__.safeDivisionInt(d_41_v11_, d_26_v1_)] = 898
            return _dafny.Map(coll10_)
        nw1_[int(4)] = iife10_()
        
        nw1_[int(5)] = d_33_v8_
        nw1_[int(6)] = d_33_v8_
        nw1_[int(7)] = ((d_36_v12_)[d_32_v7_] if (d_32_v7_) in (d_36_v12_) else d_33_v8_)
        nw1_[int(8)] = d_33_v8_
        nw1_[int(9)] = d_33_v8_
        nw1_[int(10)] = d_33_v8_
        nw1_[int(11)] = d_33_v8_
        nw1_[int(12)] = ((d_37_v13_)[d_26_v1_] if (d_26_v1_) in (d_37_v13_) else d_33_v8_)
        nw1_[int(13)] = d_33_v8_
        nw1_[int(14)] = d_33_v8_
        nw1_[int(15)] = ((d_37_v13_)[d_26_v1_] if (d_26_v1_) in (d_37_v13_) else d_33_v8_)
        nw1_[int(16)] = d_33_v8_
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in (d_29_v4_).Elements:
                d_42_v14_: int = compr_11_
                if (d_42_v14_) in (d_29_v4_):
                    coll11_[default__.safeDivisionInt(d_42_v14_, d_26_v1_)] = d_26_v1_
            return _dafny.Map(coll11_)
        nw1_[int(17)] = iife11_()
        
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in (d_39_v17_).keys.Elements:
                d_43_v15_: int = compr_12_
                if (d_43_v15_) in (d_39_v17_):
                    coll12_[(d_43_v15_) - (d_26_v1_)] = d_26_v1_
            return _dafny.Map(coll12_)
        nw1_[int(18)] = iife12_()
        
        d_40_v18_ = nw1_
        d_44_v19_: _dafny.Seq
        d_44_v19_ = _dafny.SeqWithoutIsStrInference([d_25_v0_])
        d_45_v20_: _dafny.Map
        d_45_v20_ = _dafny.Map({d_27_v2_: len(d_44_v19_)})
        d_46_v21_: str
        d_46_v21_ = _dafny.CodePoint('m')
        d_47_globalState_: GlobalState
        nw2_ = GlobalState()
        nw2_.ctor__(687, (d_25_v0_) + (d_25_v0_), True, False, False, 351, ((d_28_v3_)[((d_31_v6_)[d_32_v7_] if (d_32_v7_) in (d_31_v6_) else d_26_v1_)] if (((d_31_v6_)[d_32_v7_] if (d_32_v7_) in (d_31_v6_) else d_26_v1_)) in (d_28_v3_) else d_27_v2_), False, d_40_v18_, 825, False, -252, False, 318, True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_48_i0_ in range(default__.abs(774))]), (d_31_v6_) | (d_31_v6_), 124, False, ((d_25_v0_).set(default__.safeIndex(len(d_45_v20_), len(d_25_v0_)), d_46_v21_)) + (d_25_v0_))
        d_47_globalState_ = nw2_
        d_49_v22_: _dafny.Array
        nw3_ = _dafny.Array(int(0), 28)
        d_49_v22_ = nw3_
        d_50_v23_: C1
        nw4_ = C1()
        nw4_.ctor__(d_26_v1_, d_49_v22_, 542, d_32_v7_)
        d_50_v23_ = nw4_
        d_34_v9_ = d_34_v9_
        d_51_v24_: _dafny.Map
        d_51_v24_ = _dafny.Map({d_32_v7_: d_34_v9_})
        d_34_v9_ = ((((d_51_v24_)[d_32_v7_] if (d_32_v7_) in (d_51_v24_) else d_34_v9_)).intersection(d_34_v9_)) | ((d_34_v9_) - (d_34_v9_))
        d_52_v25_: _dafny.Seq
        d_53_v26_: bool
        out0_: _dafny.Seq
        out1_: bool
        out0_, out1_ = (d_50_v23_).m0(default__.safeModuloInt(d_26_v1_, (0) - (d_26_v1_)), d_32_v7_, d_47_globalState_)
        d_52_v25_ = out0_
        d_53_v26_ = out1_
        (d_47_globalState_).f3 = d_32_v7_
        d_54_v27_: _dafny.Array
        def lambda0_(d_55_v10_, d_56_v23_):
            def lambda1_(d_57_i1_):
                return _dafny.Map({d_55_v10_: d_56_v23_.f27})

            return lambda1_

        init0_ = lambda0_(d_35_v10_, d_50_v23_)
        nw5_ = _dafny.Array(None, 26)
        for i0_0_ in range(nw5_.length(0)):
            nw5_[i0_0_] = init0_(i0_0_)
        d_54_v27_ = nw5_
        d_58_v28_: _dafny.Map
        d_58_v28_ = _dafny.Map({d_32_v7_: _dafny.Map({d_50_v23_.f27: d_53_v26_})})
        d_59_v29_: _dafny.Map
        d_59_v29_ = _dafny.Map({((d_58_v28_)[d_32_v7_] if (d_32_v7_) in (d_58_v28_) else d_35_v10_): d_50_v23_.f27})
        index0_ = default__.safeIndex(864, (d_54_v27_).length(0))
        (d_54_v27_)[index0_] = (d_59_v29_) | (_dafny.Map({d_35_v10_: d_50_v23_.f27}))
        d_60_v30_: C5
        nw6_ = C5()
        nw6_.ctor__(d_53_v26_, d_26_v1_, d_53_v26_)
        d_60_v30_ = nw6_
        d_61_v31_: _dafny.Map
        d_61_v31_ = _dafny.Map({d_53_v26_: d_60_v30_})
        d_62_v34_: _dafny.Map
        d_62_v34_ = _dafny.Map({d_53_v26_: d_60_v30_.f23})
        index1_ = default__.safeIndex(864, (d_54_v27_).length(0))
        def iife13_():
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: _dafny.Map
                for compr_15_ in ((d_59_v29_).set(d_35_v10_, len(d_62_v34_))).keys.Elements:
                    d_63_v33_: _dafny.Map = compr_15_
                    if (d_63_v33_) in ((d_59_v29_).set(d_35_v10_, len(d_62_v34_))):
                        coll15_[d_63_v33_] = False
                return _dafny.Map(coll15_)
            coll13_ = _dafny.Map()
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: _dafny.Map
                for compr_14_ in ((d_59_v29_).set(d_35_v10_, len(d_62_v34_))).keys.Elements:
                    d_63_v33_: _dafny.Map = compr_14_
                    if (d_63_v33_) in ((d_59_v29_).set(d_35_v10_, len(d_62_v34_))):
                        coll14_[d_63_v33_] = False
                return _dafny.Map(coll14_)
            compr_13_: _dafny.Map
            for compr_13_ in ((iife14_()
            ).set(d_35_v10_, d_60_v30_.f23)).keys.Elements:
                d_64_v32_: _dafny.Map = compr_13_
                if (d_64_v32_) in ((iife15_()
                ).set(d_35_v10_, d_60_v30_.f23)):
                    coll13_[d_64_v32_] = d_50_v23_.f27
            return _dafny.Map(coll13_)
        (d_54_v27_)[index1_] = ((iife13_()
        ) | (d_59_v29_) if (d_32_v7_) not in (d_61_v31_) else _dafny.Map({d_35_v10_: d_50_v23_.f27}))
        hi0_ = len(_dafny.SeqWithoutIsStrInference([d_46_v21_]))
        for d_65_i2_ in range((676) - (d_50_v23_.f27), hi0_):
            d_66_v35_: _dafny.Array
            d_66_v35_ = d_49_v22_
            d_67_v36_: _dafny.Map
            d_67_v36_ = _dafny.Map({(default__.fm17(d_53_v26_, d_47_globalState_) if d_60_v30_.f23 else d_46_v21_): d_66_v35_})
            d_68_v37_: _dafny.Array
            def lambda2_(d_69_v23_):
                def lambda3_(d_70_i3_):
                    return _dafny.MultiSet([d_69_v23_.f27, 195])

                return lambda3_

            init1_ = lambda2_(d_50_v23_)
            nw7_ = _dafny.Array(None, 20)
            for i0_1_ in range(nw7_.length(0)):
                nw7_[i0_1_] = init1_(i0_1_)
            d_68_v37_ = nw7_
            d_71_v38_: _dafny.MultiSet
            d_71_v38_ = _dafny.MultiSet([d_26_v1_])
            index2_ = default__.safeIndex(15, (d_68_v37_).length(0))
            (d_68_v37_)[index2_] = (d_71_v38_) - ((d_71_v38_).set(d_50_v23_.f27, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))))))
            d_72_v39_: _dafny.Array
            nw8_ = _dafny.Array(_dafny.Seq({}), 24)
            d_72_v39_ = nw8_
            d_73_v40_: D6
            d_73_v40_ = D6_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ip")))
            d_74_v41_: _dafny.Seq
            d_74_v41_ = _dafny.SeqWithoutIsStrInference([d_62_v34_])
            index3_ = default__.safeIndex(15, (d_68_v37_).length(0))
            rhs0_ = not((d_50_v23_).fm0((d_33_v8_) | (d_33_v8_), (d_73_v40_).cf21, (d_74_v41_) + (((d_74_v41_).set(default__.safeIndex(d_50_v23_.f27, len(d_74_v41_)), d_62_v34_)).set(default__.safeIndex(d_50_v23_.f27, len((d_74_v41_).set(default__.safeIndex(d_50_v23_.f27, len(d_74_v41_)), d_62_v34_))), d_62_v34_)), d_47_globalState_))
            rhs1_ = d_67_v36_
            rhs2_ = ((d_71_v38_).set(14, default__.abs(d_50_v23_.f27))) | ((d_71_v38_) | (d_71_v38_))
            rhs3_ = d_53_v26_
            rhs4_ = d_72_v39_
            lhs0_ = d_47_globalState_
            lhs1_ = d_68_v37_
            lhs2_ = default__.safeIndex(15, (d_68_v37_).length(0))
            lhs3_ = d_47_globalState_
            lhs0_.f14 = rhs0_
            d_67_v36_ = rhs1_
            lhs1_[lhs2_] = rhs2_
            lhs3_.f14 = rhs3_
            d_72_v39_ = rhs4_
            (d_50_v23_).f27 = d_65_i2_
            d_49_v22_ = d_49_v22_
            (d_47_globalState_).f3 = (d_50_v23_.f27) != (214)
        (d_47_globalState_).f17 = d_50_v23_.f27
        d_75_v42_: D13
        d_75_v42_ = D13_DC36(d_50_v23_.f27, (_dafny.Set({d_50_v23_.f27, d_50_v23_.f27})).ispropersubset(d_29_v4_), d_32_v7_)
        source3_ = d_75_v42_
        if source3_.is_DC36:
            d_76___mcc_h0_ = source3_.cf46
            d_77___mcc_h1_ = source3_.cf47
            d_78___mcc_h2_ = source3_.cf48
            d_79_cf48_ = d_78___mcc_h2_
            d_80_cf47_ = d_77___mcc_h1_
            d_81_cf46_ = d_76___mcc_h0_
            (d_47_globalState_).f4 = d_60_v30_.f23
            d_82_v43_: bool
            d_82_v43_ = d_32_v7_
            d_83_v44_: C2
            nw9_ = C2()
            nw9_.ctor__(default__.safeDivisionInt(d_50_v23_.f27, d_50_v23_.f27), (d_82_v43_))
            d_83_v44_ = nw9_
            d_84_v45_: C4
            nw10_ = C4()
            nw10_.ctor__(not(d_79_cf48_), d_50_v23_.f27, d_32_v7_)
            d_84_v45_ = nw10_
            d_85_v46_: _dafny.MultiSet
            d_85_v46_ = _dafny.MultiSet([len(_dafny.Map({len(d_25_v0_): d_26_v1_})), d_26_v1_, d_81_cf46_])
            d_86_v47_: _dafny.Map
            d_86_v47_ = _dafny.Map({d_84_v45_: d_85_v46_})
            d_87_v48_: _dafny.Seq
            d_87_v48_ = _dafny.SeqWithoutIsStrInference([d_53_v26_])
            d_88_v49_: _dafny.Map
            d_88_v49_ = _dafny.Map({(((d_86_v47_)[d_84_v45_] if (d_84_v45_) in (d_86_v47_) else _dafny.MultiSet([d_26_v1_, (0) - (d_26_v1_), d_50_v23_.f27, d_26_v1_, len(d_87_v48_)]))).isdisjoint(d_85_v46_): d_25_v0_})
            d_81_cf46_ = len(d_88_v49_)
            (d_60_v30_).f23 = ((d_35_v10_)[d_50_v23_.f27] if (d_50_v23_.f27) in (d_35_v10_) else (d_60_v30_.f23) == (d_32_v7_))
        elif source3_.is_DC37:
            d_89___mcc_h3_ = source3_.cf49
            d_90___mcc_h4_ = source3_.cf50
            d_91_cf50_ = d_90___mcc_h4_
            d_92_cf49_ = d_89___mcc_h3_
            (d_47_globalState_).f18 = d_92_cf49_
            d_93_v50_: _dafny.Map
            d_93_v50_ = _dafny.Map({(d_29_v4_).intersection(d_29_v4_): _dafny.Set({not(d_60_v30_.f23)})})
            d_91_cf50_ = len(d_93_v50_)
            d_46_v21_ = _dafny.CodePoint('t')
            d_94_v51_: _dafny.Seq
            d_94_v51_ = _dafny.SeqWithoutIsStrInference([d_32_v7_])
            if (len(d_94_v51_)) <= ((0) - (default__.safeModuloInt(d_26_v1_, d_91_cf50_))):
                d_95_v52_: _dafny.Seq
                d_95_v52_ = _dafny.SeqWithoutIsStrInference([d_27_v2_, d_27_v2_])
                d_95_v52_ = (_dafny.SeqWithoutIsStrInference([d_27_v2_, d_27_v2_, d_27_v2_])).set(default__.safeIndex(default__.safeDivisionInt(d_26_v1_, d_26_v1_), len(_dafny.SeqWithoutIsStrInference([d_27_v2_, d_27_v2_, d_27_v2_]))), d_27_v2_)
                d_33_v8_ = (d_33_v8_).set((d_50_v23_.f27 if d_60_v30_.f23 else (0) - (len(d_33_v8_))), d_50_v23_.f27)
                d_96_v53_: T0
                nw11_ = C5()
                nw11_.ctor__(d_60_v30_.f23, d_50_v23_.f27, d_32_v7_)
                d_96_v53_ = nw11_
                d_96_v53_ = d_96_v53_
                d_46_v21_ = d_46_v21_
                d_97_v54_: _dafny.Map
                d_97_v54_ = _dafny.Map({d_46_v21_: d_29_v4_})
                d_97_v54_ = d_97_v54_
            elif True:
                (d_47_globalState_).f5 = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), d_46_v21_])) + (d_25_v0_))
                d_31_v6_ = (d_31_v6_).set(d_60_v30_.f23, d_91_cf50_)
                d_98_v55_: _dafny.Array
                nw12_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_98_v55_ = nw12_
                d_99_v56_: D18
                d_99_v56_ = D18_DC50(d_98_v55_)
                d_100_v57_: _dafny.Array
                nw13_ = _dafny.Array(None, 17)
                nw13_[int(0)] = d_98_v55_
                nw13_[int(1)] = d_98_v55_
                nw13_[int(2)] = d_98_v55_
                nw13_[int(3)] = d_98_v55_
                nw13_[int(4)] = d_98_v55_
                nw13_[int(5)] = (d_99_v56_).cf69
                nw13_[int(6)] = d_98_v55_
                nw13_[int(7)] = d_98_v55_
                nw13_[int(8)] = d_98_v55_
                nw13_[int(9)] = d_98_v55_
                nw13_[int(10)] = d_98_v55_
                nw13_[int(11)] = d_98_v55_
                nw13_[int(12)] = d_98_v55_
                nw13_[int(13)] = d_98_v55_
                nw13_[int(14)] = ((D18_DC50(d_98_v55_)).cf69 if d_53_v26_ else d_98_v55_)
                nw13_[int(15)] = d_98_v55_
                nw13_[int(16)] = d_98_v55_
                d_100_v57_ = nw13_
                index4_ = default__.safeIndex(528, (d_100_v57_).length(0))
                (d_100_v57_)[index4_] = d_98_v55_
                index5_ = default__.safeIndex(528, (d_100_v57_).length(0))
                (d_100_v57_)[index5_] = d_98_v55_
                d_101_v58_: _dafny.Seq
                d_101_v58_ = _dafny.SeqWithoutIsStrInference([d_34_v9_, d_34_v9_, d_34_v9_])
                d_102_v59_: D3
                d_102_v59_ = D3_DC7(d_101_v58_, d_60_v30_.f23)
                d_103_v60_: C4
                nw14_ = C4()
                nw14_.ctor__(True, d_50_v23_.f27, d_32_v7_)
                d_103_v60_ = nw14_
                d_104_v61_: _dafny.Map
                d_104_v61_ = _dafny.Map({d_103_v60_: d_50_v23_.f27})
                d_105_v62_: D19
                d_105_v62_ = D19_DC54(d_104_v61_)
                d_106_v63_: C3
                nw15_ = C3()
                nw15_.ctor__(d_50_v23_.f27, (d_102_v59_).cf12, (0) - (d_50_v23_.f27), (d_103_v60_) in ((d_105_v62_).cf72))
                d_106_v63_ = nw15_
                (d_60_v30_).f23 = ((d_34_v9_) | (d_34_v9_)).isdisjoint((d_34_v9_) | (_dafny.Set({(d_103_v60_).f24})))
        elif True:
            d_107___mcc_h5_ = source3_.cf45
            d_108_cf45_ = d_107___mcc_h5_
            d_109_v64_: C2
            nw16_ = C2()
            nw16_.ctor__(d_50_v23_.f27, d_32_v7_)
            d_109_v64_ = nw16_
            d_109_v64_ = d_109_v64_
            d_110_v65_: _dafny.Array
            d_110_v65_ = d_49_v22_
            d_111_v66_: _dafny.Map
            d_111_v66_ = _dafny.Map({d_60_v30_.f23: (d_50_v23_).f28})
            d_112_v67_: _dafny.Array
            nw17_ = _dafny.Array(None, 16)
            nw17_[int(0)] = d_49_v22_
            nw17_[int(1)] = d_110_v65_
            nw17_[int(2)] = d_110_v65_
            nw17_[int(3)] = d_110_v65_
            nw17_[int(4)] = d_110_v65_
            nw17_[int(5)] = (((d_111_v66_)[d_60_v30_.f23] if (d_60_v30_.f23) in (d_111_v66_) else (d_50_v23_).f28) if False else (d_50_v23_).f28)
            nw17_[int(6)] = d_110_v65_
            nw17_[int(7)] = d_110_v65_
            nw17_[int(8)] = (d_50_v23_).f28
            nw17_[int(9)] = d_110_v65_
            nw17_[int(10)] = d_110_v65_
            nw17_[int(11)] = (d_50_v23_).f28
            nw17_[int(12)] = d_110_v65_
            nw17_[int(13)] = (d_50_v23_).f28
            nw17_[int(14)] = d_110_v65_
            nw17_[int(15)] = d_110_v65_
            d_112_v67_ = nw17_
            index6_ = default__.safeIndex(567, (d_112_v67_).length(0))
            (d_112_v67_)[index6_] = (d_50_v23_).f28
            index7_ = default__.safeIndex(567, (d_112_v67_).length(0))
            rhs5_ = d_110_v65_
            rhs6_ = (d_38_v16_) + (d_38_v16_)
            rhs7_ = d_50_v23_.f27
            lhs4_ = d_112_v67_
            lhs5_ = default__.safeIndex(567, (d_112_v67_).length(0))
            lhs6_ = d_47_globalState_
            lhs4_[lhs5_] = rhs5_
            d_38_v16_ = rhs6_
            lhs6_.f17 = rhs7_
            (d_47_globalState_).f4 = (len(d_62_v34_)) < (d_26_v1_)
            d_113_v68_: C5
            nw18_ = C5()
            nw18_.ctor__(not(d_60_v30_.f23), 651, d_32_v7_)
            d_113_v68_ = nw18_
        (d_47_globalState_).f17 = d_26_v1_
        (d_47_globalState_).f5 = d_26_v1_
        d_114_v69_: D17
        d_114_v69_ = D17_DC47(default__.safeModuloInt(len(default__.fm41((0) - (d_26_v1_), ((d_31_v6_)[d_32_v7_] if (d_32_v7_) in (d_31_v6_) else (d_50_v23_).fm1(d_26_v1_, d_26_v1_, d_60_v30_.f23, d_47_globalState_)), d_47_globalState_)), (0) - (len(d_25_v0_))), d_60_v30_.f23, (d_26_v1_) + (len(d_33_v8_)))
        d_114_v69_ = d_114_v69_
        d_115_v70_: C0
        nw19_ = C0()
        nw19_.ctor__()
        d_115_v70_ = nw19_
        d_116_v71_: _dafny.Array
        nw20_ = _dafny.Array(None, 12)
        nw20_[int(0)] = d_115_v70_
        nw20_[int(1)] = d_115_v70_
        nw20_[int(2)] = d_115_v70_
        nw20_[int(3)] = d_115_v70_
        nw20_[int(4)] = d_115_v70_
        nw20_[int(5)] = d_115_v70_
        nw20_[int(6)] = d_115_v70_
        nw20_[int(7)] = (d_115_v70_ if d_60_v30_.f23 else d_115_v70_)
        nw20_[int(8)] = d_115_v70_
        nw20_[int(9)] = d_115_v70_
        nw20_[int(10)] = d_115_v70_
        nw20_[int(11)] = d_115_v70_
        d_116_v71_ = nw20_
        d_117_v72_: _dafny.Seq
        d_117_v72_ = _dafny.SeqWithoutIsStrInference([d_115_v70_])
        index8_ = default__.safeIndex(663, (d_116_v71_).length(0))
        (d_116_v71_)[index8_] = (d_117_v72_)[default__.safeIndex(d_50_v23_.f27, len(d_117_v72_))]
        index9_ = default__.safeIndex(663, (d_116_v71_).length(0))
        (d_116_v71_)[index9_] = d_115_v70_
        d_118_v73_: _dafny.Array
        def lambda4_(d_119_v9_):
            def lambda5_(d_120_i5_):
                return d_119_v9_

            return lambda5_

        init2_ = lambda4_(d_34_v9_)
        nw21_ = _dafny.Array(None, 9)
        for i0_2_ in range(nw21_.length(0)):
            nw21_[i0_2_] = init2_(i0_2_)
        d_118_v73_ = nw21_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_118_v73_).length(0)):
            d_121_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_121_i4_)) and ((d_121_i4_) < ((d_118_v73_).length(0)))):
                (d_118_v73_)[(d_121_i4_)] = (D7_DC16(d_34_v9_)).cf24
        d_46_v21_ = default__.fm17(d_53_v26_, d_47_globalState_)
        d_122_v74_: D12
        d_122_v74_ = D12_DC33()
        d_123_i6_: int
        d_123_i6_ = 0
        with _dafny.label("0"):
            pat_let_tv0_ = d_50_v23_
            pat_let_tv1_ = d_50_v23_
            pat_let_tv2_ = d_60_v30_
            pat_let_tv3_ = d_32_v7_
            def lambda8_(source4_):
                if source4_.is_DC33:
                    return (pat_let_tv0_.f27) != (pat_let_tv1_.f27)
                elif source4_.is_DC32:
                    d_133___mcc_h6_ = source4_.cf43
                    d_134_cf43_ = d_133___mcc_h6_
                    return pat_let_tv2_.f23
                elif True:
                    d_135___mcc_h7_ = source4_.cf44
                    d_136_cf44_ = d_135___mcc_h7_
                    return pat_let_tv3_

            while lambda8_(d_122_v74_):
                with _dafny.c_label("0"):
                    if (d_123_i6_) >= (100):
                        raise _dafny.Break("0")
                    d_123_i6_ = (d_123_i6_) + (1)
                    index10_ = default__.safeIndex(744, (d_27_v2_).length(0))
                    (d_27_v2_)[index10_] = d_60_v30_.f23
                    index11_ = default__.safeIndex(937, (d_27_v2_).length(0))
                    (d_27_v2_)[index11_] = d_32_v7_
                    d_124_v75_: _dafny.Array
                    nw22_ = _dafny.Array(None, 25)
                    d_124_v75_ = nw22_
                    d_125_v76_: _dafny.Seq
                    d_125_v76_ = _dafny.SeqWithoutIsStrInference([d_53_v26_])
                    d_126_v77_: C6
                    nw23_ = C6()
                    nw23_.ctor__(d_25_v0_, len(d_31_v6_), (d_125_v76_)[default__.safeIndex(d_26_v1_, len(d_125_v76_))])
                    d_126_v77_ = nw23_
                    index12_ = default__.safeIndex(200, (d_124_v75_).length(0))
                    (d_124_v75_)[index12_] = d_126_v77_
                    index13_ = default__.safeIndex(677, (d_27_v2_).length(0))
                    (d_27_v2_)[index13_] = d_53_v26_
                    d_127_v78_: _dafny.MultiSet
                    d_127_v78_ = _dafny.MultiSet([(d_126_v77_).f22])
                    d_128_v79_: D16
                    d_128_v79_ = D16_DC43(d_26_v1_, d_127_v78_)
                    index14_ = default__.safeIndex(744, (d_27_v2_).length(0))
                    index15_ = default__.safeIndex(937, (d_27_v2_).length(0))
                    index16_ = default__.safeIndex(200, (d_124_v75_).length(0))
                    index17_ = default__.safeIndex(677, (d_27_v2_).length(0))
                    rhs8_ = d_60_v30_.f23
                    rhs9_ = d_53_v26_
                    rhs10_ = (((d_60_v30_).fm1((d_50_v23_).fm1((d_128_v79_).cf55, d_26_v1_, d_53_v26_, d_47_globalState_), len(default__.fm21(d_47_globalState_)), d_60_v30_.f23, d_47_globalState_)) - (d_26_v1_) if d_32_v7_ else (len((d_126_v77_).f22)) * (d_50_v23_.f27))
                    rhs11_ = (d_126_v77_ if d_53_v26_ else d_126_v77_)
                    rhs12_ = False
                    lhs7_ = d_27_v2_
                    lhs8_ = default__.safeIndex(744, (d_27_v2_).length(0))
                    lhs9_ = d_27_v2_
                    lhs10_ = default__.safeIndex(937, (d_27_v2_).length(0))
                    lhs11_ = d_47_globalState_
                    lhs12_ = d_124_v75_
                    lhs13_ = default__.safeIndex(200, (d_124_v75_).length(0))
                    lhs14_ = d_27_v2_
                    lhs15_ = default__.safeIndex(677, (d_27_v2_).length(0))
                    lhs7_[lhs8_] = rhs8_
                    lhs9_[lhs10_] = rhs9_
                    lhs11_.f5 = rhs10_
                    lhs12_[lhs13_] = rhs11_
                    lhs14_[lhs15_] = rhs12_
                    (d_47_globalState_).f5 = (d_50_v23_.f27) + (d_26_v1_)
                    d_129_v80_: D13
                    d_129_v80_ = D13_DC35(d_46_v21_)
                    d_129_v80_ = (d_129_v80_ if (D17_DC48(d_60_v30_.f23, (d_52_v25_)[default__.safeIndex((default__.fm42(d_53_v26_, (d_27_v2_)[default__.safeIndex(744, (d_27_v2_).length(0))], d_47_globalState_)).cf10, len(d_52_v25_))])).cf66 else D13_DC35(d_46_v21_))
                    d_130_v81_: _dafny.Array
                    def lambda6_(d_131_v2_):
                        def lambda7_(d_132_i7_):
                            return (d_131_v2_)[default__.safeIndex(744, (d_131_v2_).length(0))]

                        return lambda7_

                    init3_ = lambda6_(d_27_v2_)
                    nw24_ = _dafny.Array(None, 8)
                    for i0_3_ in range(nw24_.length(0)):
                        nw24_[i0_3_] = init3_(i0_3_)
                    d_130_v81_ = nw24_
                    d_28_v3_ = (d_28_v3_).set(d_50_v23_.f27, d_130_v81_)
                    pass
            pass
        _dafny.print((d_25_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_26_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_28_v3_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_29_v4_) == (_dafny.Set({-60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_30_v5_) == (_dafny.Map({-60: _dafny.Set({-60})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_31_v6_) == (_dafny.Map({False: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_32_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_33_v8_) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_34_v9_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_35_v10_) == (_dafny.Map({1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_36_v12_) == (_dafny.Map({False: _dafny.Map({-60: -60})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_37_v13_) == (_dafny.Map({60: _dafny.Map({-60: -60})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_38_v16_) == (_dafny.SeqWithoutIsStrInference([-60, -60]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_39_v17_) == (_dafny.Map({5: _dafny.SeqWithoutIsStrInference([-60, -60])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[0]) == (_dafny.Map({-60: -60, 3: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[1]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[2]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[3]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[4]) == (_dafny.Map({-6: 898, -7: 898, -8: 898, -9: 898}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[5]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[6]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[7]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[8]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[9]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[10]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[11]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[12]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[13]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[14]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[15]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[16]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[17]) == (_dafny.Map({1: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_40_v18_)[18]) == (_dafny.Map({65: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_44_v19_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrqsc"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_45_v20_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_46_v21_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_47_globalState_).f1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_47_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_47_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_47_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[0]) == (_dafny.Map({-60: -60, 3: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[1]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[2]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[3]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[4]) == (_dafny.Map({-6: 898, -7: 898, -8: 898, -9: 898}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[5]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[6]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[7]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[8]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[9]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[10]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[11]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[12]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[13]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[14]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[15]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[16]) == (_dafny.Map({-60: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[17]) == (_dafny.Map({1: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_.f8)[18]) == (_dafny.Map({65: -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_47_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_47_globalState_.f15).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_globalState_).f16) == (_dafny.Map({False: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_47_globalState_.f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_47_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_47_globalState_).f19).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_50_v23_.f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v24_) == (_dafny.Map({True: _dafny.Set({True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_52_v25_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_53_v26_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[0]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[1]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[2]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[3]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[4]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[5]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[6]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[7]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[8]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[9]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[10]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[11]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[12]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[13]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[14]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[15]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[16]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[17]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[18]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[19]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[20]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[21]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[22]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[23]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[24]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_54_v27_)[25]) == (_dafny.Map({_dafny.Map({1: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_58_v28_) == (_dafny.Map({True: _dafny.Map({-60: True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_v29_) == (_dafny.Map({_dafny.Map({-60: True}): -60}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_60_v30_.f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_61_v31_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_62_v34_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_v42_).cf46))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_v42_).cf47))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_v42_).cf48))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_v69_).cf63))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_v69_).cf64))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_v69_).cf65))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_117_v72_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[0]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[1]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[2]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[3]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[4]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[5]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[6]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[7]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_118_v73_)[8]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_123_i6_))
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
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

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
        return lambda: D2_DC3(int(0), False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D2_DC2)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC3(D2, NamedTuple('DC3', [('cf3', Any), ('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC2(D2, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf8 == __o.cf8
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
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC6(D3, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC5(D3, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(_dafny.MultiSet({}), False)
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

class D4_DC10(D4, NamedTuple('DC10', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf14 == __o.cf14
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

class D5_DC12(D5, NamedTuple('DC12', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC14(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC14(D6, NamedTuple('DC14', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({self.cf21.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC17(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC17(D7, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20()
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

class D8_DC20(D8, NamedTuple('DC20', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)

class D9_DC23(D9, NamedTuple('DC23', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC26(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)

class D10_DC26(D10, NamedTuple('DC26', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC27(D10, NamedTuple('DC27', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC25(D10, NamedTuple('DC25', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC29()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)

class D11_DC29(D11, NamedTuple('DC29', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC28(D11, NamedTuple('DC28', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC33(D12, NamedTuple('DC33', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC36(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)

class D13_DC36(D13, NamedTuple('DC36', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC37(D13, NamedTuple('DC37', [('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC35(D13, NamedTuple('DC35', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC39()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)

class D14_DC39(D14, NamedTuple('DC39', [])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39)
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC38(D14, NamedTuple('DC38', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC40(D14, NamedTuple('DC40', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D15_DC41)

class D15_DC41(D15, NamedTuple('DC41', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC41({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC41) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC43(int(0), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D16_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)

class D16_DC43(D16, NamedTuple('DC43', [('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC45(D16, NamedTuple('DC45', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC42(D16, NamedTuple('DC42', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC47(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D17_DC49)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)

class D17_DC47(D17, NamedTuple('DC47', [('cf63', Any), ('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC48(D17, NamedTuple('DC48', [('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC49(D17, NamedTuple('DC49', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC49({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC49) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC46(D17, NamedTuple('DC46', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC51()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D18_DC51)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D18_DC52)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D18_DC53)

class D18_DC51(D18, NamedTuple('DC51', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC52(D18, NamedTuple('DC52', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC52({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC52) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC50(D18, NamedTuple('DC50', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC53(D18, NamedTuple('DC53', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC53({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC53) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC55(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D19_DC55)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D19_DC54)

class D19_DC55(D19, NamedTuple('DC55', [('cf73', Any), ('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC55({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC55) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC54(D19, NamedTuple('DC54', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC54({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC54) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D20_DC56)

class D20_DC56(D20, NamedTuple('DC56', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC56({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC56) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)

class D21_DC57(D21, NamedTuple('DC57', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f3: bool = False
        self.f4: bool = False
        self.f5: int = int(0)
        self.f8: _dafny.Array = _dafny.Array(None, 0)
        self.f14: bool = False
        self.f15: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f17: int = int(0)
        self.f18: bool = False
        self._f0: int = int(0)
        self._f1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f2: bool = False
        self._f6: _dafny.Array = _dafny.Array(None, 0)
        self._f7: bool = False
        self._f9: int = int(0)
        self._f10: bool = False
        self._f11: int = int(0)
        self._f12: bool = False
        self._f13: int = int(0)
        self._f16: _dafny.Map = _dafny.Map({})
        self._f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self).f17 = f17
        (self).f18 = f18
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
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
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
    def f16(self):
        return self._f16
    @property
    def f19(self):
        return self._f19

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm3(self, p0, p1, globalState):
        return not(False)

    def fm4(self, p0, p1, p2, globalState):
        return not(not((True)))


class C1(T0):
    def  __init__(self):
        self._f20: int = int(0)
        self._f21: bool = False
        self.f27: int = int(0)
        self._f28: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    def ctor__(self, f27, f28, f20, f21):
        (self).f27 = f27
        (self)._f28 = f28
        (self)._f20 = f20
        (self)._f21 = f21

    def fm0(self, p0, p1, p2, globalState):
        return (self).f21

    def fm1(self, p0, p1, p2, globalState):
        return self.f27

    def fm16(self, p0, p1, globalState):
        return ((_dafny.Set({(self).f21, (self).f21, (self).f21, (self).f21, (self).f21}) if (self).f21 else _dafny.Set({(self).f21, (self).f21, (self).f21}))).ispropersubset((_dafny.Set({(self).f21})) - (_dafny.Set({(self).f21})))

    def m0(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_137_v0_: _dafny.MultiSet
        d_137_v0_ = _dafny.MultiSet([self.f27])
        d_137_v0_ = _dafny.MultiSet([self.f27])
        d_138_i0_: int
        d_138_i0_ = 0
        with _dafny.label("1"):
            while not (p1) or ((self.f27) <= ((self).f20)):
                with _dafny.c_label("1"):
                    if (d_138_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_138_i0_ = (d_138_i0_) + (1)
                    (globalState).f18 = p1
                    d_139_v1_: _dafny.Map
                    d_139_v1_ = _dafny.Map({p0: (self).f28})
                    d_140_v2_: _dafny.Map
                    d_140_v2_ = _dafny.Map({(self).f21: not(p1)})
                    d_141_v3_: _dafny.MultiSet
                    d_141_v3_ = _dafny.MultiSet([((d_140_v2_)[p1] if (p1) in (d_140_v2_) else (self).f21)])
                    d_139_v1_ = (d_139_v1_).set(((d_141_v3_).cardinality) * ((self).f20), (self).f28)
                    (globalState).f5 = (self.f27) + (p0)
                    (globalState).f5 = ((self).f20 if (self).f21 else p0)
                    pass
            pass
        d_142_v4_: _dafny.MultiSet
        d_142_v4_ = _dafny.MultiSet([(self).f21])
        d_143_v5_: _dafny.Seq
        d_143_v5_ = _dafny.SeqWithoutIsStrInference([p0, (d_137_v0_).cardinality])
        d_144_v6_: _dafny.Map
        d_144_v6_ = _dafny.Map({self.f27: d_143_v5_})
        if (self).fm16(d_142_v4_, ((d_144_v6_)[824] if (824) in (d_144_v6_) else _dafny.SeqWithoutIsStrInference([self.f27 for d_145_i1_ in range(default__.abs(-961))])), globalState):
            d_146_v7_: D4
            d_146_v7_ = D4_DC11((self).f20, (self).f21, (self).f21)
            source5_ = d_146_v7_
            if source5_.is_DC10:
                d_147___mcc_h0_ = source5_.cf15
                d_148___mcc_h1_ = source5_.cf16
                d_149_cf16_ = d_148___mcc_h1_
                d_150_cf15_ = d_147___mcc_h0_
                d_151_v8_: _dafny.Array
                nw25_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_151_v8_ = nw25_
                index18_ = default__.safeIndex(810, (d_151_v8_).length(0))
                (d_151_v8_)[index18_] = (self).f28
                index19_ = default__.safeIndex(810, (d_151_v8_).length(0))
                (d_151_v8_)[index19_] = (self).f28
                d_152_v9_: C0
                nw26_ = C0()
                nw26_.ctor__()
                d_152_v9_ = nw26_
                (globalState).f3 = p1
                (globalState).f14 = ((self).f20) <= (self.f27)
            elif source5_.is_DC11:
                d_153___mcc_h2_ = source5_.cf17
                d_154___mcc_h3_ = source5_.cf18
                d_155___mcc_h4_ = source5_.cf19
                d_156_cf19_ = d_155___mcc_h4_
                d_157_cf18_ = d_154___mcc_h3_
                d_158_cf17_ = d_153___mcc_h2_
                (globalState).f14 = p1
                d_159_v10_: C0
                nw27_ = C0()
                nw27_.ctor__()
                d_159_v10_ = nw27_
                (globalState).f18 = not(d_156_cf19_)
                (globalState).f18 = (self).f21
            elif True:
                d_160___mcc_h5_ = source5_.cf14
                d_161_cf14_ = d_160___mcc_h5_
                d_162_v11_: _dafny.Map
                d_162_v11_ = _dafny.Map({self.f27: (self).f28})
                d_163_v12_: _dafny.Seq
                d_163_v12_ = _dafny.SeqWithoutIsStrInference([((d_162_v11_)[35] if (35) in (d_162_v11_) else (self).f28)])
                d_163_v12_ = d_163_v12_
                index20_ = default__.safeIndex(699, (d_161_cf14_).length(0))
                (d_161_cf14_)[index20_] = p1
                index21_ = default__.safeIndex(699, (d_161_cf14_).length(0))
                (d_161_cf14_)[index21_] = not ((self).f21) or (p1)
                d_164_v13_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_164_v13_ = nw28_
                d_165_v14_: _dafny.Seq
                d_165_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgaj"))
                index22_ = default__.safeIndex(348, (d_164_v13_).length(0))
                (d_164_v13_)[index22_] = d_165_v14_
                d_166_v15_: str
                d_166_v15_ = _dafny.CodePoint('k')
                index23_ = default__.safeIndex(348, (d_164_v13_).length(0))
                (d_164_v13_)[index23_] = (d_165_v14_ if True else (d_165_v14_).set(default__.safeIndex(self.f27, len(d_165_v14_)), d_166_v15_))
                d_167_v16_: D6
                d_167_v16_ = D6_DC13(d_165_v14_)
                d_165_v14_ = (d_167_v16_).cf21
            d_168_v17_: _dafny.Array
            nw29_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_168_v17_ = nw29_
            d_168_v17_ = d_168_v17_
            (globalState).f3 = p1
            (globalState).f4 = not (False) or (p1)
            d_169_v18_: _dafny.Map
            d_169_v18_ = _dafny.Map({(self).f28: (p0) * (self.f27)})
            d_170_v19_: _dafny.Seq
            d_170_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
            d_169_v18_ = (d_169_v18_).set((self).f28, (D2_DC3(((d_137_v0_)[p0] if (p0) in (d_137_v0_) else (0) - (len(d_170_v19_))), (self).f21, not((self).f21), (0) - ((self).f20), p1)).cf6)
        elif True:
            (globalState).f18 = (self).f21
            d_171_v20_: _dafny.Array
            nw30_ = _dafny.Array(False, 14)
            d_171_v20_ = nw30_
            d_172_v21_: _dafny.Seq
            d_172_v21_ = _dafny.SeqWithoutIsStrInference([d_171_v20_, d_171_v20_])
            d_173_v22_: _dafny.Array
            nw31_ = _dafny.Array(None, 7)
            nw31_[int(0)] = (d_171_v20_ if True else d_171_v20_)
            nw31_[int(1)] = d_171_v20_
            nw31_[int(2)] = d_171_v20_
            nw31_[int(3)] = d_171_v20_
            nw31_[int(4)] = d_171_v20_
            nw31_[int(5)] = d_171_v20_
            nw31_[int(6)] = (d_172_v21_)[default__.safeIndex(self.f27, len(d_172_v21_))]
            d_173_v22_ = nw31_
            d_174_v23_: _dafny.Seq
            d_174_v23_ = _dafny.SeqWithoutIsStrInference([d_173_v22_])
            d_173_v22_ = (d_174_v23_)[default__.safeIndex(p0, len(d_174_v23_))]
            (self).f27 = self.f27
            (globalState).f17 = p0
            d_175_v24_: str
            d_175_v24_ = _dafny.CodePoint('j')
            d_176_v25_: _dafny.MultiSet
            d_176_v25_ = _dafny.MultiSet([d_175_v24_, _dafny.CodePoint('l'), d_175_v24_, d_175_v24_])
            d_176_v25_ = _dafny.MultiSet([d_175_v24_, d_175_v24_, default__.fm17(not((self).f21), globalState), d_175_v24_, d_175_v24_])
        d_177_v26_: _dafny.Seq
        d_177_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogosrhy"))
        (globalState).f15 = (d_177_v26_) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_178_i2_ in range(default__.abs(531))])) + (d_177_v26_))
        d_179_v27_: _dafny.Array
        nw32_ = _dafny.Array(False, 15)
        d_179_v27_ = nw32_
        d_180_v28_: _dafny.Set
        d_180_v28_ = _dafny.Set({(self).f21, (self).f21})
        d_181_v29_: _dafny.Seq
        d_181_v29_ = _dafny.SeqWithoutIsStrInference([d_180_v28_, d_180_v28_])
        d_182_v30_: _dafny.Map
        d_182_v30_ = _dafny.Map({p1: self.f27})
        d_183_v31_: str
        d_183_v31_ = _dafny.CodePoint('m')
        d_184_v32_: _dafny.Array
        nw33_ = _dafny.Array(None, 25)
        nw33_[int(0)] = d_180_v28_
        nw33_[int(1)] = d_180_v28_
        nw33_[int(2)] = d_180_v28_
        nw33_[int(3)] = (d_181_v29_)[default__.safeIndex(self.f27, len(d_181_v29_))]
        nw33_[int(4)] = d_180_v28_
        nw33_[int(5)] = _dafny.Set({not(p1), not((self).f21)})
        nw33_[int(6)] = d_180_v28_
        nw33_[int(7)] = default__.fm18(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))), ((d_182_v30_)[p1] if (p1) in (d_182_v30_) else 221), d_183_v31_, globalState)
        nw33_[int(8)] = d_180_v28_
        nw33_[int(9)] = d_180_v28_
        nw33_[int(10)] = d_180_v28_
        nw33_[int(11)] = (D7_DC16(d_180_v28_)).cf24
        nw33_[int(12)] = d_180_v28_
        nw33_[int(13)] = _dafny.Set({p1, (self).f21})
        nw33_[int(14)] = d_180_v28_
        nw33_[int(15)] = d_180_v28_
        nw33_[int(16)] = d_180_v28_
        nw33_[int(17)] = d_180_v28_
        nw33_[int(18)] = d_180_v28_
        nw33_[int(19)] = d_180_v28_
        nw33_[int(20)] = _dafny.Set({True, p1})
        nw33_[int(21)] = d_180_v28_
        nw33_[int(22)] = d_180_v28_
        nw33_[int(23)] = d_180_v28_
        nw33_[int(24)] = default__.fm18(p0, p0, d_183_v31_, globalState)
        d_184_v32_ = nw33_
        d_185_v33_: _dafny.Map
        d_185_v33_ = _dafny.Map({d_179_v27_: d_184_v32_})
        d_185_v33_ = (d_185_v33_).set(d_179_v27_, d_184_v32_)
        index24_ = default__.safeIndex(776, (d_179_v27_).length(0))
        (d_179_v27_)[index24_] = p1
        d_186_v34_: _dafny.Seq
        d_186_v34_ = _dafny.SeqWithoutIsStrInference([p1])
        index25_ = default__.safeIndex(776, (d_179_v27_).length(0))
        (d_179_v27_)[index25_] = ((d_186_v34_)[default__.safeIndex(len(d_182_v30_), len(d_186_v34_))]) and ((self).f21)
        d_187_v35_: _dafny.Seq
        d_187_v35_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f28])
        d_188_v36_: _dafny.Seq
        d_188_v36_ = _dafny.SeqWithoutIsStrInference([d_187_v35_])
        r0 = ((d_188_v36_)[default__.safeIndex((d_143_v5_)[default__.safeIndex(len(d_180_v28_), len(d_143_v5_))], len(d_188_v36_))]) + (d_187_v35_)
        r1 = p1
        return r0, r1

    @property
    def f28(self):
        return self._f28

class C2(T0):
    def  __init__(self):
        self._f20: int = int(0)
        self._f21: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    def ctor__(self, f20, f21):
        (self)._f20 = f20
        (self)._f21 = f21

    def fm0(self, p0, p1, p2, globalState):
        return not(((0) - ((len(_dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, True, (self).f21, (self).f21]))) + (len(_dafny.SeqWithoutIsStrInference([(self).f20 for d_189_i0_ in range(default__.abs(189))]))))) in (_dafny.SeqWithoutIsStrInference([(self).f20, (self).f20])))

    def fm1(self, p0, p1, p2, globalState):
        return (self).f20

    def m0(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_190_v0_: _dafny.Seq
        d_190_v0_ = _dafny.SeqWithoutIsStrInference([True])
        d_191_v1_: int
        d_192_v2_: bool
        d_193_v3_: int
        d_194_v4_: int
        out2_: int
        out3_: bool
        out4_: int
        out5_: int
        out2_, out3_, out4_, out5_ = (self).m3((d_190_v0_) == (d_190_v0_), globalState)
        d_191_v1_ = out2_
        d_192_v2_ = out3_
        d_193_v3_ = out4_
        d_194_v4_ = out5_
        d_195_v5_: _dafny.Set
        d_195_v5_ = _dafny.Set({len(_dafny.Map({True: d_191_v1_})), (self).f20})
        d_196_v6_: _dafny.Map
        d_196_v6_ = _dafny.Map({p1: len(d_195_v5_)})
        d_196_v6_ = (d_196_v6_).set(p1, (self).f20)
        if True:
            d_197_v7_: _dafny.Seq
            d_197_v7_ = _dafny.SeqWithoutIsStrInference([p0, (0) - ((self).f20)])
            d_197_v7_ = (_dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((0) - ((self).f20), p0) for d_198_i0_ in range(default__.abs(418))])).set(default__.safeIndex((self).f20, len(_dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((0) - ((self).f20), p0) for d_199_i0_ in range(default__.abs(418))]))), p0)
            d_200_v8_: str
            d_200_v8_ = _dafny.CodePoint('v')
            d_200_v8_ = d_200_v8_
            d_201_v9_: _dafny.MultiSet
            d_201_v9_ = _dafny.MultiSet([(self).f21])
            d_193_v3_ = (self).fm1(d_191_v1_, (self).fm1(d_194_v4_, d_194_v4_, d_192_v2_, globalState), (d_201_v9_).issubset(d_201_v9_), globalState)
            d_202_v10_: C0
            nw34_ = C0()
            nw34_.ctor__()
            d_202_v10_ = nw34_
            d_203_v11_: C0
            nw35_ = C0()
            nw35_.ctor__()
            d_203_v11_ = nw35_
        elif True:
            d_204_v12_: _dafny.Array
            nw36_ = _dafny.Array(False, 1)
            d_204_v12_ = nw36_
            d_205_v13_: _dafny.Seq
            d_205_v13_ = _dafny.SeqWithoutIsStrInference([d_204_v12_])
            d_206_v14_: D3
            d_206_v14_ = D3_DC5(d_205_v13_)
            pat_let_tv4_ = d_205_v13_
            pat_let_tv5_ = d_205_v13_
            def iife17_(_pat_let1_0):
                def iife18_(d_207_dt__update__tmp_h0_):
                    def iife19_(_pat_let2_0):
                        def iife20_(d_208_dt__update_hcf9_h0_):
                            return D3_DC5(d_208_dt__update_hcf9_h0_)
                        return iife20_(_pat_let2_0)
                    return iife19_(pat_let_tv4_)
                return iife18_(_pat_let1_0)
            def iife16_(_pat_let0_0):
                def iife21_(d_209_dt__update__tmp_h1_):
                    def iife22_(_pat_let3_0):
                        def iife23_(d_210_dt__update_hcf9_h1_):
                            return D3_DC5(d_210_dt__update_hcf9_h1_)
                        return iife23_(_pat_let3_0)
                    return iife22_(pat_let_tv5_)
                return iife21_(_pat_let0_0)
            source6_ = iife16_((d_206_v14_ if (self).f21 else iife17_(d_206_v14_)))
            if source6_.is_DC6:
                d_211___mcc_h0_ = source6_.cf10
                d_212_cf10_ = d_211___mcc_h0_
                d_213_v15_: _dafny.Array
                def lambda9_(d_214_v3_):
                    def lambda10_(d_215_i1_):
                        return (d_215_i1_) - (d_214_v3_)

                    return lambda10_

                init4_ = lambda9_(d_193_v3_)
                nw37_ = _dafny.Array(None, 16)
                for i0_4_ in range(nw37_.length(0)):
                    nw37_[i0_4_] = init4_(i0_4_)
                d_213_v15_ = nw37_
                d_216_v16_: T0
                nw38_ = C1()
                nw38_.ctor__(len(_dafny.Set({d_192_v2_})), d_213_v15_, d_194_v4_, True)
                d_216_v16_ = nw38_
                d_217_v17_: _dafny.Map
                d_217_v17_ = _dafny.Map({p0: d_216_v16_})
                d_218_v18_: _dafny.Set
                d_218_v18_ = _dafny.Set({d_192_v2_, p1, d_192_v2_, d_192_v2_})
                d_219_v19_: _dafny.Map
                d_219_v19_ = _dafny.Map({((self).f20) + (len(d_217_v17_)): d_218_v18_})
                d_220_v20_: str
                d_220_v20_ = _dafny.CodePoint('f')
                index26_ = default__.safeIndex(612, (d_204_v12_).length(0))
                (d_204_v12_)[index26_] = (d_218_v18_).isdisjoint(default__.fm18(d_212_cf10_, d_194_v4_, d_220_v20_, globalState))
                d_221_v21_: _dafny.Map
                d_221_v21_ = _dafny.Map({(self).f20: (self).f20})
                d_222_v22_: _dafny.Seq
                d_222_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlhaskp"))
                d_223_v23_: _dafny.Map
                d_223_v23_ = _dafny.Map({p1: (d_190_v0_)[default__.safeIndex(d_193_v3_, len(d_190_v0_))]})
                d_224_v24_: _dafny.Seq
                d_224_v24_ = _dafny.SeqWithoutIsStrInference([default__.fm19(True, globalState), d_223_v23_])
                index27_ = default__.safeIndex(612, (d_204_v12_).length(0))
                rhs13_ = d_219_v19_
                rhs14_ = (self).f20
                rhs15_ = (d_216_v16_).fm0(d_221_v21_, d_222_v22_, d_224_v24_, globalState)
                rhs16_ = not(p1)
                lhs16_ = globalState
                lhs17_ = d_204_v12_
                lhs18_ = default__.safeIndex(612, (d_204_v12_).length(0))
                d_219_v19_ = rhs13_
                d_212_cf10_ = rhs14_
                lhs16_.f18 = rhs15_
                lhs17_[lhs18_] = rhs16_
                d_190_v0_ = (d_190_v0_) + (d_190_v0_)
                d_225_v25_: _dafny.Array
                d_225_v25_ = d_213_v15_
                d_226_v26_: _dafny.Seq
                d_226_v26_ = _dafny.SeqWithoutIsStrInference([d_213_v15_])
                d_227_v27_: _dafny.Map
                d_227_v27_ = _dafny.Map({d_225_v25_: (d_226_v26_)[default__.safeIndex(931, len(d_226_v26_))]})
                d_228_v28_: _dafny.Set
                d_228_v28_ = _dafny.Set({((d_227_v27_)[d_225_v25_] if (d_225_v25_) in (d_227_v27_) else d_213_v15_), d_213_v15_, d_213_v15_})
                index28_ = default__.safeIndex(612, (d_204_v12_).length(0))
                (d_204_v12_)[index28_] = not(not(not((d_228_v28_).issubset(d_228_v28_))))
                (globalState).f3 = (self).f21
            elif source6_.is_DC7:
                d_229___mcc_h1_ = source6_.cf11
                d_230___mcc_h2_ = source6_.cf12
                d_231_cf12_ = d_230___mcc_h2_
                d_232_cf11_ = d_229___mcc_h1_
                d_191_v1_ = (self).f20
                d_233_v29_: _dafny.Array
                nw39_ = _dafny.Array(None, 20)
                d_233_v29_ = nw39_
                d_234_v30_: _dafny.Map
                d_234_v30_ = _dafny.Map({d_192_v2_: d_233_v29_})
                d_235_v31_: _dafny.Seq
                d_235_v31_ = _dafny.SeqWithoutIsStrInference([d_234_v30_, d_234_v30_, _dafny.Map({False: d_233_v29_})])
                d_192_v2_ = (len((d_235_v31_)[default__.safeIndex(p0, len(d_235_v31_))])) == (p0)
                d_236_v32_: _dafny.Seq
                d_236_v32_ = _dafny.SeqWithoutIsStrInference([d_194_v4_])
                d_193_v3_ = ((d_193_v3_) + ((d_236_v32_)[default__.safeIndex(d_194_v4_, len(d_236_v32_))])) - (p0)
                d_237_v33_: _dafny.Map
                d_237_v33_ = _dafny.Map({d_192_v2_: p1})
                d_238_v34_: _dafny.Seq
                d_238_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uj"))
                d_239_v35_: _dafny.Map
                d_239_v35_ = _dafny.Map({518: d_191_v1_})
                d_240_v36_: bool
                d_240_v36_ = p1
                d_241_v37_: _dafny.Map
                d_241_v37_ = _dafny.Map({d_240_v36_: False})
                d_242_v38_: _dafny.Map
                d_242_v38_ = _dafny.Map({d_204_v12_: False})
                d_243_v39_: _dafny.Array
                nw40_ = _dafny.Array(None, 28)
                nw40_[int(0)] = len(d_237_v33_)
                nw40_[int(1)] = d_191_v1_
                nw40_[int(2)] = len(_dafny.Map({d_238_v34_: len(d_239_v35_)}))
                nw40_[int(3)] = d_193_v3_
                nw40_[int(4)] = d_194_v4_
                nw40_[int(5)] = (self).fm1(d_191_v1_, len(d_239_v35_), d_231_cf12_, globalState)
                nw40_[int(6)] = 388
                nw40_[int(7)] = d_191_v1_
                nw40_[int(8)] = d_193_v3_
                nw40_[int(9)] = 135
                nw40_[int(10)] = default__.safeDivisionInt((_dafny.MultiSet([p0])).cardinality, d_194_v4_)
                nw40_[int(11)] = ((0) - (len(d_238_v34_))) - (len(_dafny.SeqWithoutIsStrInference([(self).f20 for d_244_i2_ in range(default__.abs(-625))])))
                nw40_[int(12)] = ((self).f20 if d_231_cf12_ else len(d_241_v37_))
                nw40_[int(13)] = len(d_242_v38_)
                nw40_[int(14)] = (self).f20
                nw40_[int(15)] = d_194_v4_
                nw40_[int(16)] = (self).f20
                nw40_[int(17)] = d_191_v1_
                nw40_[int(18)] = 730
                nw40_[int(19)] = d_191_v1_
                nw40_[int(20)] = p0
                nw40_[int(21)] = (self).f20
                nw40_[int(22)] = (self).f20
                nw40_[int(23)] = d_193_v3_
                nw40_[int(24)] = d_193_v3_
                nw40_[int(25)] = p0
                nw40_[int(26)] = d_191_v1_
                nw40_[int(27)] = 525
                d_243_v39_ = nw40_
                index29_ = default__.safeIndex(370, (d_243_v39_).length(0))
                (d_243_v39_)[index29_] = d_193_v3_
                d_245_v40_: _dafny.Set
                d_245_v40_ = _dafny.Set({p1})
                d_246_v41_: D7
                d_246_v41_ = D7_DC16(d_245_v40_)
                index30_ = default__.safeIndex(370, (d_243_v39_).length(0))
                (d_243_v39_)[index30_] = len((_dafny.SeqWithoutIsStrInference([d_246_v41_, d_246_v41_])) + ((_dafny.SeqWithoutIsStrInference([d_246_v41_ for d_247_i3_ in range(default__.abs(374))])) + (default__.fm20((self).f20, d_193_v3_, d_193_v3_, globalState))))
            elif source6_.is_DC5:
                d_248___mcc_h3_ = source6_.cf9
                d_249_cf9_ = d_248___mcc_h3_
                d_250_v42_: _dafny.Map
                d_250_v42_ = _dafny.Map({d_194_v4_: True})
                d_251_v43_: _dafny.Map
                d_251_v43_ = _dafny.Map({D7_DC16(_dafny.Set({((d_250_v42_)[636] if (636) in (d_250_v42_) else (self).f21), (self).f21})): p1})
                d_252_v45_: _dafny.Seq
                d_252_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnfdr"))
                d_253_v46_: _dafny.Map
                d_253_v46_ = _dafny.Map({(self).f21: p1})
                d_254_v47_: _dafny.Seq
                d_254_v47_ = _dafny.SeqWithoutIsStrInference([d_253_v46_])
                d_255_v48_: _dafny.Map
                def iife24_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(926, 940):
                        d_256_v44_: int = compr_16_
                        if ((926) <= (d_256_v44_)) and ((d_256_v44_) < (940)):
                            coll16_[(d_256_v44_) + (len(d_190_v0_))] = len(_dafny.Set({True, p1, d_192_v2_}))
                    return _dafny.Map(coll16_)
                d_255_v48_ = _dafny.Map({not(d_192_v2_): (self).fm0(iife24_()
                , d_252_v45_, d_254_v47_, globalState)})
                d_257_v49_: _dafny.Set
                d_257_v49_ = _dafny.Set({((d_253_v46_)[d_192_v2_] if (d_192_v2_) in (d_253_v46_) else True)})
                d_258_v50_: D7
                d_258_v50_ = D7_DC16(d_257_v49_)
                (globalState).f3 = ((d_251_v43_)[d_258_v50_] if (d_258_v50_) in (d_251_v43_) else not(p1))
                d_259_v51_: _dafny.Map
                d_259_v51_ = _dafny.Map({len(d_252_v45_): (self).f20})
                d_260_v52_: _dafny.Map
                d_260_v52_ = _dafny.Map({(0) - (d_193_v3_): d_254_v47_})
                d_261_v53_: _dafny.MultiSet
                d_261_v53_ = _dafny.MultiSet([(self).fm0(d_259_v51_, d_252_v45_, ((d_260_v52_)[p0] if (p0) in (d_260_v52_) else d_254_v47_), globalState), d_192_v2_])
                d_262_v54_: _dafny.MultiSet
                d_262_v54_ = _dafny.MultiSet([((self).f21) == ((self).f21), (d_261_v53_).ispropersubset(_dafny.MultiSet([not(not(p1)), (self).f21])), p1])
                d_263_v55_: _dafny.Seq
                d_263_v55_ = _dafny.SeqWithoutIsStrInference([d_262_v54_])
                d_261_v53_ = ((d_263_v55_)[default__.safeIndex(d_194_v4_, len(d_263_v55_))]) - ((_dafny.MultiSet([not(p1), True])) - (d_261_v53_))
                d_264_v56_: _dafny.Seq
                d_264_v56_ = _dafny.SeqWithoutIsStrInference([p0, d_193_v3_])
                d_265_v57_: D7
                d_265_v57_ = D7_DC18(d_264_v56_, p0)
                d_266_v58_: _dafny.Array
                nw41_ = _dafny.Array(None, 9)
                nw41_[int(0)] = d_191_v1_
                nw41_[int(1)] = len(d_252_v45_)
                nw41_[int(2)] = d_191_v1_
                nw41_[int(3)] = d_191_v1_
                nw41_[int(4)] = (self).f20
                nw41_[int(5)] = d_194_v4_
                nw41_[int(6)] = p0
                nw41_[int(7)] = (d_265_v57_).cf27
                nw41_[int(8)] = (self).f20
                d_266_v58_ = nw41_
                d_267_v59_: _dafny.Array
                d_267_v59_ = d_266_v58_
                d_268_v60_: _dafny.MultiSet
                d_268_v60_ = _dafny.MultiSet([d_267_v59_, d_267_v59_])
                d_269_v61_: _dafny.Map
                d_269_v61_ = _dafny.Map({d_194_v4_: d_268_v60_})
                d_270_v62_: D6
                d_270_v62_ = D6_DC14(764)
                d_271_v63_: _dafny.Seq
                d_271_v63_ = _dafny.SeqWithoutIsStrInference([d_266_v58_])
                d_272_v64_: _dafny.Map
                d_272_v64_ = _dafny.Map({_dafny.CodePoint('w'): (self).f21})
                d_273_v65_: D8
                d_273_v65_ = D8_DC19(d_259_v51_)
                d_274_v66_: _dafny.Array
                nw42_ = _dafny.Array(None, 26)
                nw42_[int(0)] = d_268_v60_
                nw42_[int(1)] = d_268_v60_
                nw42_[int(2)] = (d_268_v60_) - (d_268_v60_)
                nw42_[int(3)] = d_268_v60_
                nw42_[int(4)] = (((d_269_v61_)[(d_270_v62_).cf22] if ((d_270_v62_).cf22) in (d_269_v61_) else ((d_269_v61_)[d_191_v1_] if (d_191_v1_) in (d_269_v61_) else d_268_v60_))) | (d_268_v60_)
                nw42_[int(5)] = (d_268_v60_).intersection(_dafny.MultiSet([d_267_v59_]))
                nw42_[int(6)] = d_268_v60_
                nw42_[int(7)] = d_268_v60_
                nw42_[int(8)] = (d_268_v60_).set(d_267_v59_, default__.abs(d_193_v3_))
                nw42_[int(9)] = d_268_v60_
                nw42_[int(10)] = d_268_v60_
                nw42_[int(11)] = (d_268_v60_) | (d_268_v60_)
                nw42_[int(12)] = d_268_v60_
                nw42_[int(13)] = _dafny.MultiSet(d_271_v63_)
                nw42_[int(14)] = ((d_268_v60_).set(d_267_v59_, default__.abs(len(d_272_v64_)))).intersection(d_268_v60_)
                nw42_[int(15)] = d_268_v60_
                nw42_[int(16)] = (d_268_v60_).set(d_267_v59_, default__.abs(d_193_v3_))
                nw42_[int(17)] = (_dafny.MultiSet([d_266_v58_, d_266_v58_, d_267_v59_])).set(d_267_v59_, default__.abs(((d_259_v51_)[len((d_273_v65_).cf28)] if (len((d_273_v65_).cf28)) in (d_259_v51_) else d_194_v4_)))
                nw42_[int(18)] = d_268_v60_
                nw42_[int(19)] = d_268_v60_
                nw42_[int(20)] = d_268_v60_
                nw42_[int(21)] = d_268_v60_
                nw42_[int(22)] = (d_268_v60_).intersection(d_268_v60_)
                nw42_[int(23)] = d_268_v60_
                nw42_[int(24)] = _dafny.MultiSet([d_267_v59_])
                nw42_[int(25)] = d_268_v60_
                d_274_v66_ = nw42_
                index31_ = default__.safeIndex(585, (d_274_v66_).length(0))
                (d_274_v66_)[index31_] = d_268_v60_
                index32_ = default__.safeIndex(585, (d_274_v66_).length(0))
                rhs17_ = not(((self).f21 if p1 else p1))
                rhs18_ = (0) - ((self).f20)
                rhs19_ = d_268_v60_
                lhs19_ = globalState
                lhs20_ = globalState
                lhs21_ = d_274_v66_
                lhs22_ = default__.safeIndex(585, (d_274_v66_).length(0))
                lhs19_.f4 = rhs17_
                lhs20_.f5 = rhs18_
                lhs21_[lhs22_] = rhs19_
                d_275_v67_: _dafny.Seq
                d_275_v67_ = _dafny.SeqWithoutIsStrInference([d_190_v0_])
                d_276_v68_: str
                d_276_v68_ = _dafny.CodePoint('s')
                d_277_v69_: _dafny.Set
                d_277_v69_ = _dafny.Set({d_276_v68_})
                d_278_v70_: D9
                d_278_v70_ = D9_DC22(d_190_v0_)
                d_279_v71_: D9
                d_279_v71_ = D9_DC23(d_193_v3_, d_192_v2_)
                d_280_v72_: _dafny.Seq
                d_280_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p1: len(_dafny.SeqWithoutIsStrInference([d_276_v68_ for d_281_i4_ in range(default__.abs(783))]))})])
                d_282_v73_: _dafny.Map
                d_282_v73_ = _dafny.Map({d_280_v72_: p1})
                d_283_v74_: _dafny.Array
                nw43_ = _dafny.Array(None, 22)
                nw43_[int(0)] = (_dafny.SeqWithoutIsStrInference([(self).f21, d_192_v2_])) + (d_190_v0_)
                nw43_[int(1)] = (default__.fm21(globalState)).set(default__.safeIndex(-21, len(default__.fm21(globalState))), (self).f21)
                nw43_[int(2)] = (d_190_v0_ if p1 else _dafny.SeqWithoutIsStrInference([p1, p1]))
                nw43_[int(3)] = (d_275_v67_)[default__.safeIndex(len(d_277_v69_), len(d_275_v67_))]
                nw43_[int(4)] = (d_190_v0_).set(default__.safeIndex(d_194_v4_, len(d_190_v0_)), p1)
                nw43_[int(5)] = d_190_v0_
                nw43_[int(6)] = default__.fm21(globalState)
                nw43_[int(7)] = _dafny.SeqWithoutIsStrInference([p1])
                nw43_[int(8)] = d_190_v0_
                nw43_[int(9)] = (d_190_v0_).set(default__.safeIndex(d_194_v4_, len(d_190_v0_)), d_192_v2_)
                nw43_[int(10)] = (((d_190_v0_) + (d_190_v0_)).set(default__.safeIndex(112, len((d_190_v0_) + (d_190_v0_))), True)).set(default__.safeIndex(d_194_v4_, len(((d_190_v0_) + (d_190_v0_)).set(default__.safeIndex(112, len((d_190_v0_) + (d_190_v0_))), True))), d_192_v2_)
                nw43_[int(11)] = d_190_v0_
                nw43_[int(12)] = d_190_v0_
                nw43_[int(13)] = d_190_v0_
                nw43_[int(14)] = ((_dafny.SeqWithoutIsStrInference([p1, d_192_v2_])) + ((d_278_v70_).cf30)).set(default__.safeIndex(d_194_v4_, len((_dafny.SeqWithoutIsStrInference([p1, d_192_v2_])) + ((d_278_v70_).cf30))), (self).f21)
                nw43_[int(15)] = default__.fm21(globalState)
                nw43_[int(16)] = (d_190_v0_ if (d_279_v71_).cf32 else d_190_v0_)
                nw43_[int(17)] = _dafny.SeqWithoutIsStrInference([(self).f21])
                nw43_[int(18)] = (d_190_v0_) + (_dafny.SeqWithoutIsStrInference([p1, (self).fm0(default__.fm22(d_193_v3_, globalState), d_252_v45_, _dafny.SeqWithoutIsStrInference([d_255_v48_]), globalState)]))
                nw43_[int(19)] = default__.fm21(globalState)
                nw43_[int(20)] = (d_190_v0_) + (_dafny.SeqWithoutIsStrInference([(self).f21, ((d_282_v73_)[d_280_v72_] if (d_280_v72_) in (d_282_v73_) else (self).f21), (d_190_v0_)[default__.safeIndex((0) - ((self).f20), len(d_190_v0_))], d_192_v2_]))
                nw43_[int(21)] = (d_190_v0_) + ((_dafny.SeqWithoutIsStrInference([(self).f21, p1])).set(default__.safeIndex((0) - (d_194_v4_), len(_dafny.SeqWithoutIsStrInference([(self).f21, p1]))), (self).f21))
                d_283_v74_ = nw43_
                index33_ = default__.safeIndex(823, (d_283_v74_).length(0))
                (d_283_v74_)[index33_] = d_190_v0_
                index34_ = default__.safeIndex(823, (d_283_v74_).length(0))
                (d_283_v74_)[index34_] = (_dafny.SeqWithoutIsStrInference([d_192_v2_])) + (d_190_v0_)
            elif True:
                d_284___mcc_h4_ = source6_.cf13
                d_285_cf13_ = d_284___mcc_h4_
                d_286_v75_: _dafny.Seq
                d_286_v75_ = _dafny.SeqWithoutIsStrInference([(self).f20])
                d_287_v76_: D7
                d_287_v76_ = D7_DC18(d_286_v75_, (self).f20)
                d_288_v77_: _dafny.Map
                d_288_v77_ = _dafny.Map({d_192_v2_: p1})
                d_289_v78_: _dafny.Seq
                d_289_v78_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_192_v2_: p1})).set(d_192_v2_, p1), d_288_v77_, d_288_v77_])
                d_290_v79_: _dafny.Map
                d_290_v79_ = _dafny.Map({d_191_v1_: (self).f21})
                rhs20_ = (d_194_v4_) * (len((((d_287_v76_).cf26).set(default__.safeIndex(len((d_289_v78_)[default__.safeIndex(d_194_v4_, len(d_289_v78_))]), len((d_287_v76_).cf26)), d_193_v3_)).set(default__.safeIndex(p0, len(((d_287_v76_).cf26).set(default__.safeIndex(len((d_289_v78_)[default__.safeIndex(d_194_v4_, len(d_289_v78_))]), len((d_287_v76_).cf26)), d_193_v3_))), len(d_290_v79_))))
                rhs21_ = p0
                lhs23_ = globalState
                lhs23_.f17 = rhs20_
                d_194_v4_ = rhs21_
                d_291_v80_: _dafny.Seq
                d_291_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmfshk"))
                d_292_v81_: _dafny.Map
                d_292_v81_ = _dafny.Map({d_291_v80_: 170})
                d_292_v81_ = (d_292_v81_).set((d_291_v80_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_293_i5_ in range(default__.abs(74))])), d_191_v1_)
                d_194_v4_ = (self).fm1(d_194_v4_, len((_dafny.SeqWithoutIsStrInference([840 for d_294_i6_ in range(default__.abs(-454))])) + (_dafny.SeqWithoutIsStrInference([p0 for d_295_i7_ in range(default__.abs(85))]))), d_192_v2_, globalState)
                d_296_v82_: str
                d_296_v82_ = _dafny.CodePoint('m')
                d_296_v82_ = d_296_v82_
            d_297_v83_: _dafny.Seq
            d_297_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoj"))
            (globalState).f15 = (d_297_v83_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsyle")))
            d_298_v84_: _dafny.Seq
            d_298_v84_ = _dafny.SeqWithoutIsStrInference([d_191_v1_])
            d_299_v85_: _dafny.MultiSet
            d_299_v85_ = _dafny.MultiSet([d_192_v2_])
            d_300_v86_: _dafny.MultiSet
            d_300_v86_ = _dafny.MultiSet([(default__.fm23(globalState)).set((self).f21, default__.abs(d_194_v4_)), d_299_v85_, d_299_v85_])
            d_301_v87_: _dafny.Array
            nw44_ = _dafny.Array(None, 23)
            nw44_[int(0)] = (self).f20
            nw44_[int(1)] = (self).f20
            nw44_[int(2)] = ((self).f20 if (self).f21 else (0) - (p0))
            nw44_[int(3)] = d_191_v1_
            nw44_[int(4)] = (self).f20
            nw44_[int(5)] = d_191_v1_
            nw44_[int(6)] = d_194_v4_
            nw44_[int(7)] = len(d_195_v5_)
            nw44_[int(8)] = default__.safeDivisionInt(d_194_v4_, d_191_v1_)
            nw44_[int(9)] = p0
            nw44_[int(10)] = default__.safeDivisionInt(d_193_v3_, d_194_v4_)
            nw44_[int(11)] = d_194_v4_
            nw44_[int(12)] = d_193_v3_
            nw44_[int(13)] = d_194_v4_
            nw44_[int(14)] = (0) - (d_191_v1_)
            nw44_[int(15)] = (0) - (d_194_v4_)
            nw44_[int(16)] = (d_298_v84_)[default__.safeIndex(len(_dafny.Set({p1})), len(d_298_v84_))]
            nw44_[int(17)] = p0
            nw44_[int(18)] = (self).f20
            nw44_[int(19)] = default__.safeModuloInt(d_194_v4_, (self).f20)
            nw44_[int(20)] = ((self).f20) * (d_194_v4_)
            nw44_[int(21)] = p0
            nw44_[int(22)] = (0) - ((0) - ((d_300_v86_).cardinality))
            d_301_v87_ = nw44_
            index35_ = default__.safeIndex(164, (d_301_v87_).length(0))
            (d_301_v87_)[index35_] = p0
            index36_ = default__.safeIndex(164, (d_301_v87_).length(0))
            (d_301_v87_)[index36_] = ((p0) * (len(d_297_v83_))) - (((self).f20) + ((self).fm1(d_194_v4_, d_193_v3_, (self).f21, globalState)))
            (globalState).f5 = d_191_v1_
            d_302_v89_: _dafny.MultiSet
            d_302_v89_ = _dafny.MultiSet([(0) - ((self).f20)])
            d_303_v90_: D8
            def iife25_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in (d_302_v89_).Elements:
                    d_304_v88_: int = compr_17_
                    if (d_304_v88_) in (d_302_v89_):
                        coll17_[(d_304_v88_) * (len(d_297_v83_))] = d_193_v3_
                return _dafny.Map(coll17_)
            d_303_v90_ = D8_DC19(iife25_()
)
            d_305_v91_: _dafny.Map
            d_305_v91_ = _dafny.Map({(d_302_v89_).cardinality: 409})
            d_303_v90_ = D8_DC19(d_305_v91_)
        d_306_v92_: D4
        d_306_v92_ = D4_DC11(d_191_v1_, (d_190_v0_)[default__.safeIndex((0) - (p0), len(d_190_v0_))], p1)
        d_307_v93_: _dafny.Set
        d_307_v93_ = _dafny.Set({p1})
        pat_let_tv6_ = d_192_v2_
        pat_let_tv7_ = d_193_v3_
        d_308_v94_: _dafny.Array
        nw45_ = _dafny.Array(None, 28)
        nw45_[int(0)] = d_306_v92_
        nw45_[int(1)] = d_306_v92_
        nw45_[int(2)] = D4_DC11((self).f20, p1, d_192_v2_)
        nw45_[int(3)] = d_306_v92_
        nw45_[int(4)] = d_306_v92_
        nw45_[int(5)] = d_306_v92_
        nw45_[int(6)] = default__.fm24(d_191_v1_, d_192_v2_, d_307_v93_, p1, globalState)
        nw45_[int(7)] = d_306_v92_
        nw45_[int(8)] = d_306_v92_
        nw45_[int(9)] = d_306_v92_
        nw45_[int(10)] = d_306_v92_
        nw45_[int(11)] = d_306_v92_
        nw45_[int(12)] = d_306_v92_
        nw45_[int(13)] = d_306_v92_
        nw45_[int(14)] = D4_DC11(d_191_v1_, (self).f21, d_192_v2_)
        nw45_[int(15)] = d_306_v92_
        nw45_[int(16)] = d_306_v92_
        nw45_[int(17)] = default__.fm24(127, p1, d_307_v93_, d_192_v2_, globalState)
        nw45_[int(18)] = d_306_v92_
        nw45_[int(19)] = d_306_v92_
        nw45_[int(20)] = d_306_v92_
        def iife26_(_pat_let4_0):
            def iife27_(d_309_dt__update__tmp_h2_):
                def iife28_(_pat_let5_0):
                    def iife29_(d_310_dt__update_hcf19_h0_):
                        return D4_DC11((d_309_dt__update__tmp_h2_).cf17, (d_309_dt__update__tmp_h2_).cf18, d_310_dt__update_hcf19_h0_)
                    return iife29_(_pat_let5_0)
                return iife28_((self).f21)
            return iife27_(_pat_let4_0)
        nw45_[int(21)] = iife26_(d_306_v92_)
        def iife30_(_pat_let6_0):
            def iife31_(d_311_dt__update__tmp_h3_):
                def iife32_(_pat_let7_0):
                    def iife33_(d_312_dt__update_hcf18_h0_):
                        return D4_DC11((d_311_dt__update__tmp_h3_).cf17, d_312_dt__update_hcf18_h0_, (d_311_dt__update__tmp_h3_).cf19)
                    return iife33_(_pat_let7_0)
                return iife32_(pat_let_tv6_)
            return iife31_(_pat_let6_0)
        nw45_[int(22)] = iife30_(d_306_v92_)
        nw45_[int(23)] = d_306_v92_
        nw45_[int(24)] = d_306_v92_
        nw45_[int(25)] = d_306_v92_
        def iife34_(_pat_let8_0):
            def iife35_(d_313_dt__update__tmp_h4_):
                def iife36_(_pat_let9_0):
                    def iife37_(d_314_dt__update_hcf18_h1_):
                        def iife38_(_pat_let10_0):
                            def iife39_(d_315_dt__update_hcf17_h0_):
                                return D4_DC11(d_315_dt__update_hcf17_h0_, d_314_dt__update_hcf18_h1_, (d_313_dt__update__tmp_h4_).cf19)
                            return iife39_(_pat_let10_0)
                        return iife38_(pat_let_tv7_)
                    return iife37_(_pat_let9_0)
                return iife36_((self).f21)
            return iife35_(_pat_let8_0)
        nw45_[int(26)] = iife34_(d_306_v92_)
        nw45_[int(27)] = d_306_v92_
        d_308_v94_ = nw45_
        index37_ = default__.safeIndex(748, (d_308_v94_).length(0))
        (d_308_v94_)[index37_] = D4_DC11((self).f20, not((self).f21), (self).f21)
        pat_let_tv8_ = d_191_v1_
        pat_let_tv9_ = d_191_v1_
        index38_ = default__.safeIndex(748, (d_308_v94_).length(0))
        def iife40_(_pat_let11_0):
            def iife41_(d_316_dt__update__tmp_h5_):
                def iife42_(_pat_let12_0):
                    def iife43_(d_317_dt__update_hcf19_h1_):
                        return D4_DC11((d_316_dt__update__tmp_h5_).cf17, (d_316_dt__update__tmp_h5_).cf18, d_317_dt__update_hcf19_h1_)
                    return iife43_(_pat_let12_0)
                return iife42_((pat_let_tv8_) < (pat_let_tv9_))
            return iife41_(_pat_let11_0)
        (d_308_v94_)[index38_] = iife40_(d_306_v92_)
        d_318_v95_: str
        d_318_v95_ = _dafny.CodePoint('w')
        d_319_v96_: _dafny.Map
        d_319_v96_ = _dafny.Map({d_318_v95_: p1})
        if (d_192_v2_) and (not (((d_319_v96_)[default__.fm17(d_192_v2_, globalState)] if (default__.fm17(d_192_v2_, globalState)) in (d_319_v96_) else (d_190_v0_)[default__.safeIndex(d_191_v1_, len(d_190_v0_))])) or (p1)):
            d_320_v97_: _dafny.Seq
            d_320_v97_ = _dafny.SeqWithoutIsStrInference([d_193_v3_, d_191_v1_])
            d_321_v98_: _dafny.Map
            d_321_v98_ = _dafny.Map({d_191_v1_: len(d_320_v97_)})
            d_321_v98_ = (d_321_v98_) | (default__.fm22(398, globalState))
            d_322_v99_: _dafny.Seq
            d_322_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqqgtwun"))
            d_323_v100_: _dafny.Map
            d_323_v100_ = _dafny.Map({d_322_v99_: d_318_v95_})
            d_324_v101_: _dafny.Seq
            d_324_v101_ = _dafny.SeqWithoutIsStrInference([d_320_v97_])
            d_320_v97_ = ((_dafny.SeqWithoutIsStrInference([len(d_323_v100_)])) + ((d_324_v101_)[default__.safeIndex(d_194_v4_, len(d_324_v101_))])) + (d_320_v97_)
            d_325_v102_: _dafny.Array
            nw46_ = _dafny.Array(None, 4)
            nw46_[int(0)] = (self).f21
            nw46_[int(1)] = False
            nw46_[int(2)] = False
            nw46_[int(3)] = p1
            d_325_v102_ = nw46_
            d_326_v103_: _dafny.Seq
            d_326_v103_ = _dafny.SeqWithoutIsStrInference([d_325_v102_, d_325_v102_, d_325_v102_, d_325_v102_, d_325_v102_])
            d_327_v104_: D3
            d_327_v104_ = D3_DC5(d_326_v103_)
            d_327_v104_ = d_327_v104_
            d_328_v105_: _dafny.Seq
            d_328_v105_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({len(d_196_v6_), (self).f20})) - (default__.fm25(globalState)), d_195_v5_, d_195_v5_])
            d_328_v105_ = _dafny.SeqWithoutIsStrInference([d_195_v5_, d_195_v5_, d_195_v5_])
            d_192_v2_ = p1
        elif True:
            d_329_v106_: _dafny.Seq
            d_329_v106_ = _dafny.SeqWithoutIsStrInference([d_194_v4_])
            (globalState).f5 = ((d_329_v106_)[default__.safeIndex((self).f20, len(d_329_v106_))]) - (d_191_v1_)
            d_330_v107_: C0
            nw47_ = C0()
            nw47_.ctor__()
            d_330_v107_ = nw47_
            d_331_v108_: _dafny.Seq
            d_331_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssmu"))
            d_332_v109_: _dafny.MultiSet
            d_332_v109_ = _dafny.MultiSet([len((d_331_v108_).set(default__.safeIndex(d_194_v4_, len(d_331_v108_)), default__.fm17(p1, globalState))), default__.safeDivisionInt(d_194_v4_, p0), (0) - ((self).f20), (len(d_331_v108_)) - (d_194_v4_), (self).fm1(d_191_v1_, d_193_v3_, d_192_v2_, globalState)])
            d_332_v109_ = (d_332_v109_) | (d_332_v109_)
            d_333_v110_: C0
            nw48_ = C0()
            nw48_.ctor__()
            d_333_v110_ = nw48_
            d_334_v111_: D8
            d_334_v111_ = D8_DC20()
            d_334_v111_ = D8_DC20()
        d_335_v112_: _dafny.Map
        d_335_v112_ = _dafny.Map({(self).f20: d_194_v4_})
        d_336_v113_: _dafny.Seq
        d_336_v113_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ac"))
        d_335_v112_ = (d_335_v112_).set(((self).f20) + (len(d_336_v113_)), p0)
        d_337_v114_: _dafny.Array
        nw49_ = _dafny.Array(int(0), 26)
        d_337_v114_ = nw49_
        d_338_v115_: _dafny.Seq
        d_338_v115_ = _dafny.SeqWithoutIsStrInference([d_337_v114_, d_337_v114_, d_337_v114_])
        r0 = (d_338_v115_ if d_192_v2_ else d_338_v115_)
        r1 = p1
        return r0, r1

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        hi1_ = (self).f20
        for d_339_i0_ in range(len(default__.fm26((self).f21, globalState)), hi1_):
            (globalState).f3 = False
            d_340_v0_: _dafny.Seq
            d_340_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfbohx"))
            d_341_v1_: str
            d_341_v1_ = _dafny.CodePoint('r')
            (globalState).f15 = ((d_340_v0_).set(default__.safeIndex(d_339_i0_, len(d_340_v0_)), d_341_v1_)) + (d_340_v0_)
            (globalState).f15 = d_340_v0_
            (globalState).f3 = (self).f21
        d_342_v2_: _dafny.Seq
        d_342_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fahojwr"))
        (globalState).f5 = (self).fm1((self).f20, len(d_342_v2_), (self).f21, globalState)
        d_343_v3_: _dafny.Array
        def lambda11_(d_344_i1_):
            return default__.safeModuloInt(d_344_i1_, (self).f20)

        init5_ = lambda11_
        nw50_ = _dafny.Array(None, 4)
        for i0_5_ in range(nw50_.length(0)):
            nw50_[i0_5_] = init5_(i0_5_)
        d_343_v3_ = nw50_
        index39_ = default__.safeIndex(768, (d_343_v3_).length(0))
        (d_343_v3_)[index39_] = len(_dafny.SeqWithoutIsStrInference([(self).f20, (self).f20]))
        d_345_v4_: str
        d_345_v4_ = _dafny.CodePoint('g')
        d_346_v5_: _dafny.Seq
        d_346_v5_ = _dafny.SeqWithoutIsStrInference([len((d_342_v2_).set(default__.safeIndex((self).f20, len(d_342_v2_)), d_345_v4_))])
        d_347_v6_: _dafny.Map
        d_347_v6_ = _dafny.Map({p0: len(d_346_v5_)})
        index40_ = default__.safeIndex(768, (d_343_v3_).length(0))
        (d_343_v3_)[index40_] = ((self).f20) - (((d_347_v6_)[(self).f21] if ((self).f21) in (d_347_v6_) else (self).f20))
        d_348_v7_: _dafny.Map
        d_348_v7_ = _dafny.Map({d_343_v3_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpkgyo")))})
        d_349_v8_: _dafny.MultiSet
        d_349_v8_ = _dafny.MultiSet([_dafny.CodePoint('a')])
        d_350_v9_: D7
        d_350_v9_ = D7_DC18(d_346_v5_, (d_349_v8_).cardinality)
        hi2_ = default__.safeModuloInt(((d_348_v7_)[d_343_v3_] if (d_343_v3_) in (d_348_v7_) else (d_350_v9_).cf27), (self).f20)
        for d_351_i2_ in range((len(d_342_v2_)) + ((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]), hi2_):
            d_352_v10_: D3
            d_352_v10_ = D3_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksptncyh"))))
            d_353_v11_: _dafny.Map
            d_353_v11_ = _dafny.Map({d_351_i2_: (0) - (d_351_i2_)})
            d_354_v12_: _dafny.Map
            d_354_v12_ = _dafny.Map({(self).f21: not((self).f21)})
            d_355_v13_: _dafny.Seq
            d_355_v13_ = _dafny.SeqWithoutIsStrInference([d_354_v12_, _dafny.Map({p0: False}), d_354_v12_, d_354_v12_])
            d_356_v14_: _dafny.Seq
            d_356_v14_ = _dafny.SeqWithoutIsStrInference([d_354_v12_, (d_355_v13_)[default__.safeIndex((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], len(d_355_v13_))], d_354_v12_, _dafny.Map({(self).f21: (self).f21})])
            index41_ = default__.safeIndex(768, (d_343_v3_).length(0))
            rhs22_ = d_345_v4_
            rhs23_ = (d_352_v10_).cf10
            rhs24_ = (self).f21
            rhs25_ = not ((self).f21) or (not((self).fm0(d_353_v11_, d_342_v2_, d_356_v14_, globalState)))
            lhs24_ = d_343_v3_
            lhs25_ = default__.safeIndex(768, (d_343_v3_).length(0))
            lhs26_ = globalState
            lhs27_ = globalState
            d_345_v4_ = rhs22_
            lhs24_[lhs25_] = rhs23_
            lhs26_.f14 = rhs24_
            lhs27_.f18 = rhs25_
            d_357_v15_: _dafny.Seq
            d_357_v15_ = _dafny.SeqWithoutIsStrInference([(self).fm0(d_353_v11_, d_342_v2_, d_356_v14_, globalState), p0, (self).fm0(d_353_v11_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgw")), d_355_v13_, globalState)])
            d_358_v16_: D4
            d_358_v16_ = D4_DC11((self).f20, (self).f21, not((self).f21))
            if (d_357_v15_)[default__.safeIndex((d_358_v16_).cf17, len(d_357_v15_))]:
                d_359_v17_: _dafny.Array
                nw51_ = _dafny.Array(False, 10)
                d_359_v17_ = nw51_
                index42_ = default__.safeIndex(120, (d_359_v17_).length(0))
                def iife44_():
                    coll18_ = _dafny.Map()
                    compr_18_: int
                    for compr_18_ in _dafny.IntegerRange(-127, 39):
                        d_360_v18_: int = compr_18_
                        if ((-127) <= (d_360_v18_)) and ((d_360_v18_) < (39)):
                            coll18_[(d_360_v18_) - ((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])] = d_351_i2_
                    return _dafny.Map(coll18_)
                (d_359_v17_)[index42_] = (self).fm0(iife44_()
                , d_342_v2_, d_356_v14_, globalState)
                index43_ = default__.safeIndex(120, (d_359_v17_).length(0))
                (d_359_v17_)[index43_] = (self).f21
                d_361_v19_: D4
                d_361_v19_ = D4_DC9(d_359_v17_)
                d_362_v20_: _dafny.Map
                d_362_v20_ = _dafny.Map({(d_359_v17_)[default__.safeIndex(120, (d_359_v17_).length(0))]: (d_361_v19_).cf14})
                index44_ = default__.safeIndex(768, (d_343_v3_).length(0))
                (d_343_v3_)[index44_] = (self).fm1(d_351_i2_, len(d_362_v20_), p0, globalState)
                (globalState).f17 = (len(d_342_v2_)) + ((self).f20)
                index45_ = default__.safeIndex(120, (d_359_v17_).length(0))
                rhs26_ = (d_351_i2_) in (_dafny.SeqWithoutIsStrInference([297, d_351_i2_]))
                rhs27_ = (0) - ((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])
                lhs28_ = d_359_v17_
                lhs29_ = default__.safeIndex(120, (d_359_v17_).length(0))
                lhs30_ = globalState
                lhs28_[lhs29_] = rhs26_
                lhs30_.f5 = rhs27_
                d_363_v21_: _dafny.Map
                d_363_v21_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqpn")): ((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]) * (d_351_i2_)})
                d_364_v22_: _dafny.Map
                d_364_v22_ = _dafny.Map({len(d_342_v2_): (self).f21})
                d_365_v23_: D9
                d_365_v23_ = D9_DC23(d_351_i2_, p0)
                d_363_v21_ = (d_363_v21_).set(default__.fm26(not(((d_364_v22_)[(d_365_v23_).cf31] if ((d_365_v23_).cf31) in (d_364_v22_) else True)), globalState), (d_351_i2_) - (d_351_i2_))
            elif True:
                d_346_v5_ = d_346_v5_
                index46_ = default__.safeIndex(768, (d_343_v3_).length(0))
                (d_343_v3_)[index46_] = (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]
                d_366_v24_: _dafny.Array
                nw52_ = _dafny.Array(_dafny.Set({}), 4)
                d_366_v24_ = nw52_
                rhs28_ = d_366_v24_
                rhs29_ = default__.safeDivisionInt((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], (0) - (((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]) + (d_351_i2_)))
                rhs30_ = (0) - ((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])
                rhs31_ = (((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]) + (193)) == ((0) - ((self).f20))
                lhs31_ = globalState
                lhs32_ = globalState
                d_366_v24_ = rhs28_
                r3 = rhs29_
                lhs31_.f17 = rhs30_
                lhs32_.f4 = rhs31_
                d_367_v25_: _dafny.Map
                d_367_v25_ = _dafny.Map({p0: d_357_v15_})
                (globalState).f5 = ((0) - ((self).f20)) - (len((_dafny.Map({False: d_357_v15_})) | (d_367_v25_)))
                d_368_v26_: _dafny.Array
                nw53_ = _dafny.Array(None, 1)
                nw53_[int(0)] = (self).fm0(_dafny.Map({(0) - ((self).f20): (self).fm1((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], p0, globalState)}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbcdjta")), d_355_v13_, globalState)
                d_368_v26_ = nw53_
                d_369_v27_: D8
                d_369_v27_ = D8_DC20()
                rhs32_ = (self).f21
                rhs33_ = d_368_v26_
                rhs34_ = D8_DC20()
                lhs33_ = globalState
                lhs33_.f3 = rhs32_
                d_368_v26_ = rhs33_
                d_369_v27_ = rhs34_
            rhs35_ = (self).f21
            rhs36_ = 948
            lhs34_ = globalState
            lhs35_ = globalState
            lhs34_.f4 = rhs35_
            lhs35_.f5 = rhs36_
            d_370_v28_: _dafny.MultiSet
            d_370_v28_ = _dafny.MultiSet([not(p0)])
            d_371_v29_: C1
            nw54_ = C1()
            nw54_.ctor__((self).f20, d_343_v3_, default__.safeModuloInt((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], (self).f20), (default__.fm23(globalState)).issubset(d_370_v28_))
            d_371_v29_ = nw54_
        if (self).f21:
            r0 = (self).f20
            d_372_v30_: _dafny.Map
            d_372_v30_ = _dafny.Map({d_345_v4_: (self).f21})
            d_373_v31_: _dafny.Map
            d_373_v31_ = _dafny.Map({d_372_v30_: (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]})
            d_373_v31_ = (d_373_v31_).set(d_372_v30_, (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])
            d_374_v32_: _dafny.Map
            d_374_v32_ = _dafny.Map({(self).f20: (self).f20})
            d_375_v33_: _dafny.Map
            d_375_v33_ = _dafny.Map({True: True})
            d_376_v34_: _dafny.Set
            d_376_v34_ = _dafny.Set({p0, p0, (self).f21, (self).f21})
            r1 = (d_376_v34_).ispropersubset((_dafny.Set({(self).fm0(d_374_v32_, d_342_v2_, (_dafny.SeqWithoutIsStrInference([_dafny.Map({p0: p0})])).set(default__.safeIndex((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], len(_dafny.SeqWithoutIsStrInference([_dafny.Map({p0: p0})]))), d_375_v33_), globalState), p0, p0})) | (_dafny.Set({p0})))
            d_377_v35_: _dafny.Array
            nw55_ = _dafny.Array(None, 10)
            nw55_[int(0)] = d_345_v4_
            nw55_[int(1)] = (d_345_v4_ if p0 else _dafny.CodePoint('t'))
            nw55_[int(2)] = d_345_v4_
            nw55_[int(3)] = default__.fm17(p0, globalState)
            nw55_[int(4)] = d_345_v4_
            nw55_[int(5)] = d_345_v4_
            nw55_[int(6)] = d_345_v4_
            nw55_[int(7)] = d_345_v4_
            nw55_[int(8)] = d_345_v4_
            nw55_[int(9)] = _dafny.CodePoint('k')
            d_377_v35_ = nw55_
            d_378_v36_: _dafny.MultiSet
            d_378_v36_ = _dafny.MultiSet([True, False, (self).f21])
            d_379_v37_: D3
            d_379_v37_ = D3_DC6((d_378_v36_).cardinality)
            d_380_v38_: _dafny.MultiSet
            d_380_v38_ = _dafny.MultiSet([(_dafny.Set({(self).f21, (self).f21, (self).f21})) - (d_376_v34_)])
            index47_ = default__.safeIndex(768, (d_343_v3_).length(0))
            rhs37_ = d_377_v35_
            rhs38_ = d_379_v37_
            rhs39_ = (self).f20
            rhs40_ = (len(d_342_v2_) if ((self).f20) >= ((self).f20) else (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])
            rhs41_ = d_380_v38_
            lhs36_ = d_343_v3_
            lhs37_ = default__.safeIndex(768, (d_343_v3_).length(0))
            d_377_v35_ = rhs37_
            d_379_v37_ = rhs38_
            r0 = rhs39_
            lhs36_[lhs37_] = rhs40_
            d_380_v38_ = rhs41_
            d_381_v39_: _dafny.Seq
            d_381_v39_ = _dafny.SeqWithoutIsStrInference([d_375_v33_])
            (globalState).f3 = (self).fm0(d_374_v32_, default__.fm26(p0, globalState), (d_381_v39_) + (d_381_v39_), globalState)
        elif True:
            (globalState).f14 = (self).f21
            (globalState).f5 = (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]
            (globalState).f17 = (self).f20
            d_382_v40_: _dafny.Map
            d_382_v40_ = _dafny.Map({d_342_v2_: (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]})
            index48_ = default__.safeIndex(768, (d_343_v3_).length(0))
            (d_343_v3_)[index48_] = default__.safeModuloInt(len(d_382_v40_), (self).f20)
            d_383_v41_: bool
            d_383_v41_ = p0
            d_383_v41_ = d_383_v41_
        d_384_v42_: _dafny.Map
        d_384_v42_ = _dafny.Map({(d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]: (self).f20})
        d_385_v43_: _dafny.Map
        d_385_v43_ = _dafny.Map({p0: p0})
        d_386_v44_: _dafny.Seq
        d_386_v44_ = _dafny.SeqWithoutIsStrInference([d_385_v43_])
        d_387_v45_: _dafny.Map
        d_387_v45_ = _dafny.Map({(self).f21: (self).fm0(d_384_v42_, d_342_v2_, d_386_v44_, globalState)})
        d_388_v46_: _dafny.Map
        d_388_v46_ = _dafny.Map({(self).f20: d_385_v43_})
        if not(((d_385_v43_) | (d_387_v45_)) == ((((d_388_v46_)[240] if (240) in (d_388_v46_) else default__.fm19((self).f21, globalState))) | (d_387_v45_))):
            d_389_v47_: _dafny.Set
            d_389_v47_ = _dafny.Set({d_345_v4_, d_345_v4_})
            d_390_v48_: _dafny.Seq
            d_390_v48_ = _dafny.SeqWithoutIsStrInference([d_346_v5_, d_346_v5_])
            d_387_v45_ = (d_387_v45_).set((len(d_389_v47_)) < (len(d_390_v48_)), True)
            d_391_v49_: _dafny.Array
            def lambda12_(d_392_i3_):
                return (self).f21

            init6_ = lambda12_
            nw56_ = _dafny.Array(None, 18)
            for i0_6_ in range(nw56_.length(0)):
                nw56_[i0_6_] = init6_(i0_6_)
            d_391_v49_ = nw56_
            d_393_v50_: D4
            d_393_v50_ = D4_DC9(d_391_v49_)
            pat_let_tv10_ = d_391_v49_
            d_394_v51_: _dafny.MultiSet
            def iife45_(_pat_let13_0):
                def iife46_(d_395_dt__update__tmp_h0_):
                    def iife47_(_pat_let14_0):
                        def iife48_(d_396_dt__update_hcf14_h0_):
                            return D4_DC9(d_396_dt__update_hcf14_h0_)
                        return iife48_(_pat_let14_0)
                    return iife47_(pat_let_tv10_)
                return iife46_(_pat_let13_0)
            d_394_v51_ = _dafny.MultiSet([iife45_(d_393_v50_), d_393_v50_, D4_DC9(d_391_v49_), d_393_v50_])
            d_397_v52_: _dafny.MultiSet
            d_397_v52_ = _dafny.MultiSet([(self).f20, (0) - ((self).f20), ((d_394_v51_)[d_393_v50_] if (d_393_v50_) in (d_394_v51_) else (0) - ((self).f20)), 969])
            d_398_v53_: _dafny.Map
            d_398_v53_ = _dafny.Map({d_397_v52_: d_350_v9_})
            pat_let_tv11_ = d_346_v5_
            def iife49_(_pat_let15_0):
                def iife50_(d_399_dt__update__tmp_h1_):
                    def iife51_(_pat_let16_0):
                        def iife52_(d_400_dt__update_hcf26_h0_):
                            return D7_DC18(d_400_dt__update_hcf26_h0_, (d_399_dt__update__tmp_h1_).cf27)
                        return iife52_(_pat_let16_0)
                    return iife51_(pat_let_tv11_)
                return iife50_(_pat_let15_0)
            d_398_v53_ = (d_398_v53_).set(_dafny.MultiSet([(self).f20, (self).f20, -398, (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]]), iife49_(D7_DC18(d_346_v5_, (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])))
            d_401_v54_: _dafny.MultiSet
            d_401_v54_ = _dafny.MultiSet([(self).f21])
            d_402_v55_: _dafny.Map
            d_402_v55_ = _dafny.Map({(d_401_v54_).cardinality: d_343_v3_})
            d_403_v56_: _dafny.Set
            d_403_v56_ = _dafny.Set({(self).f21})
            d_404_v57_: C1
            nw57_ = C1()
            nw57_.ctor__((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], ((d_402_v55_)[(d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]] if ((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]) in (d_402_v55_) else d_343_v3_), default__.safeDivisionInt((self).f20, (0) - (len(d_403_v56_))), not((self).f21))
            d_404_v57_ = nw57_
            rhs42_ = 632
            rhs43_ = d_404_v57_
            r3 = rhs42_
            d_404_v57_ = rhs43_
            d_405_v58_: _dafny.Map
            d_405_v58_ = _dafny.Map({d_391_v49_: (0) - ((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])})
            d_405_v58_ = (d_405_v58_).set(d_391_v49_, (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])
            index49_ = default__.safeIndex(768, (d_343_v3_).length(0))
            (d_343_v3_)[index49_] = (0) - (d_404_v57_.f27)
        elif True:
            d_406_v59_: _dafny.Array
            def lambda13_(d_407_p0_):
                def lambda14_(d_408_i4_):
                    return d_407_p0_

                return lambda14_

            init7_ = lambda13_(p0)
            nw58_ = _dafny.Array(None, 9)
            for i0_7_ in range(nw58_.length(0)):
                nw58_[i0_7_] = init7_(i0_7_)
            d_406_v59_ = nw58_
            d_409_v60_: D4
            d_409_v60_ = D4_DC9(d_406_v59_)
            d_410_v61_: _dafny.Seq
            d_410_v61_ = _dafny.SeqWithoutIsStrInference([not(((self).f20) != ((d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])), (self).f21])
            rhs44_ = d_409_v60_
            rhs45_ = d_410_v61_
            rhs46_ = (self).f21
            rhs47_ = (self).f20
            lhs38_ = globalState
            d_409_v60_ = rhs44_
            d_410_v61_ = rhs45_
            r1 = rhs46_
            lhs38_.f17 = rhs47_
            d_411_v62_: _dafny.Map
            d_411_v62_ = _dafny.Map({(d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]: p0})
            d_384_v42_ = (d_384_v42_).set(((self).f20) * (len(d_411_v62_)), (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))])
            d_412_v63_: _dafny.Set
            d_412_v63_ = _dafny.Set({(self).f21, (self).f21})
            d_413_v64_: _dafny.MultiSet
            d_413_v64_ = _dafny.MultiSet([d_412_v63_])
            index50_ = default__.safeIndex(228, (d_406_v59_).length(0))
            (d_406_v59_)[index50_] = ((self).f20) >= ((d_413_v64_).cardinality)
            index51_ = default__.safeIndex(228, (d_406_v59_).length(0))
            (d_406_v59_)[index51_] = ((self).f21) and (p0)
            (globalState).f4 = (d_406_v59_)[default__.safeIndex(228, (d_406_v59_).length(0))]
            d_414_v65_: _dafny.Seq
            d_414_v65_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))], (self).f20, (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]]), d_346_v5_])
            d_415_v66_: _dafny.Seq
            d_415_v66_ = _dafny.SeqWithoutIsStrInference([(d_346_v5_) + ((d_414_v65_)[default__.safeIndex((self).f20, len(d_414_v65_))]), d_346_v5_])
            d_414_v65_ = d_415_v66_
        r0 = (d_343_v3_)[default__.safeIndex(768, (d_343_v3_).length(0))]
        r1 = (self).fm0(d_384_v42_, d_342_v2_, d_386_v44_, globalState)
        r2 = -558
        d_416_v67_: _dafny.Map
        d_416_v67_ = _dafny.Map({len(d_346_v5_): p0})
        d_417_v68_: _dafny.Map
        d_417_v68_ = _dafny.Map({d_416_v67_: p0})
        r3 = default__.safeModuloInt(len(d_417_v68_), (self).f20)
        return r0, r1, r2, r3

    def m4(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        hi3_ = (self).f20
        for d_418_i0_ in range((0) - (p0), hi3_):
            d_419_v0_: _dafny.Array
            nw59_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_419_v0_ = nw59_
            d_420_v1_: _dafny.Array
            nw60_ = _dafny.Array(False, 26)
            d_420_v1_ = nw60_
            d_421_v2_: D3
            d_421_v2_ = D3_DC6(d_418_i0_)
            d_422_v3_: _dafny.Seq
            d_422_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvejr"))
            rhs48_ = ((d_422_v3_).set(default__.safeIndex(-478, len(d_422_v3_)), _dafny.CodePoint('r'))) + (d_422_v3_)
            rhs49_ = d_419_v0_
            rhs50_ = d_420_v1_
            rhs51_ = D3_DC6(d_418_i0_)
            rhs52_ = p0
            lhs39_ = globalState
            lhs40_ = globalState
            lhs39_.f15 = rhs48_
            d_419_v0_ = rhs49_
            d_420_v1_ = rhs50_
            d_421_v2_ = rhs51_
            lhs40_.f5 = rhs52_
            index52_ = default__.safeIndex(278, (d_419_v0_).length(0))
            (d_419_v0_)[index52_] = d_420_v1_
            index53_ = default__.safeIndex(278, (d_419_v0_).length(0))
            (d_419_v0_)[index53_] = d_420_v1_
            index54_ = default__.safeIndex(685, (d_420_v1_).length(0))
            (d_420_v1_)[index54_] = p1
            arr0_ = (d_419_v0_)[default__.safeIndex(278, (d_419_v0_).length(0))]
            index55_ = default__.safeIndex(969, ((d_419_v0_)[default__.safeIndex(278, (d_419_v0_).length(0))]).length(0))
            arr0_[index55_] = (self).f21
            index56_ = default__.safeIndex(685, (d_420_v1_).length(0))
            arr1_ = (d_419_v0_)[default__.safeIndex(278, (d_419_v0_).length(0))]
            index57_ = default__.safeIndex(969, ((d_419_v0_)[default__.safeIndex(278, (d_419_v0_).length(0))]).length(0))
            rhs53_ = (self).f21
            rhs54_ = p1
            lhs41_ = d_420_v1_
            lhs42_ = default__.safeIndex(685, (d_420_v1_).length(0))
            lhs43_ = (d_419_v0_)[default__.safeIndex(278, (d_419_v0_).length(0))]
            lhs44_ = default__.safeIndex(969, ((d_419_v0_)[default__.safeIndex(278, (d_419_v0_).length(0))]).length(0))
            lhs41_[lhs42_] = rhs53_
            lhs43_[lhs44_] = rhs54_
            d_423_v4_: _dafny.Set
            d_423_v4_ = _dafny.Set({(d_420_v1_)[default__.safeIndex(685, (d_420_v1_).length(0))]})
            d_424_v5_: _dafny.Array
            nw61_ = _dafny.Array(None, 6)
            nw61_[int(0)] = d_418_i0_
            nw61_[int(1)] = (self).f20
            nw61_[int(2)] = -412
            nw61_[int(3)] = (0) - (d_418_i0_)
            nw61_[int(4)] = d_418_i0_
            nw61_[int(5)] = len(d_423_v4_)
            d_424_v5_ = nw61_
            d_425_v6_: C1
            nw62_ = C1()
            nw62_.ctor__(p0, d_424_v5_, d_418_i0_, p1)
            d_425_v6_ = nw62_
        d_426_i1_: int
        d_426_i1_ = 0
        with _dafny.label("2"):
            while (self).f21:
                with _dafny.c_label("2"):
                    if (d_426_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_426_i1_ = (d_426_i1_) + (1)
                    d_427_v7_: _dafny.Seq
                    d_427_v7_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_428_v8_: _dafny.Seq
                    d_428_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogxg"))
                    d_429_v9_: _dafny.Seq
                    d_429_v9_ = _dafny.SeqWithoutIsStrInference([(self).f21, p1])
                    d_430_v10_: _dafny.Array
                    nw63_ = _dafny.Array(None, 24)
                    nw63_[int(0)] = (d_427_v7_)[default__.safeIndex(p0, len(d_427_v7_))]
                    nw63_[int(1)] = len(d_428_v8_)
                    nw63_[int(2)] = (self).f20
                    nw63_[int(3)] = (self).f20
                    nw63_[int(4)] = (d_427_v7_)[default__.safeIndex((self).f20, len(d_427_v7_))]
                    nw63_[int(5)] = (self).f20
                    nw63_[int(6)] = p0
                    nw63_[int(7)] = len(d_428_v8_)
                    nw63_[int(8)] = 551
                    nw63_[int(9)] = p0
                    nw63_[int(10)] = (self).fm1(-469, (self).f20, (self).f21, globalState)
                    nw63_[int(11)] = 158
                    nw63_[int(12)] = p0
                    nw63_[int(13)] = p0
                    nw63_[int(14)] = p0
                    nw63_[int(15)] = (self).f20
                    nw63_[int(16)] = len(d_428_v8_)
                    nw63_[int(17)] = len(d_429_v9_)
                    nw63_[int(18)] = p0
                    nw63_[int(19)] = len(_dafny.Set({True, p1, (self).f21, (self).f21}))
                    nw63_[int(20)] = len(_dafny.Map({d_427_v7_: p1}))
                    nw63_[int(21)] = (0) - ((self).fm1(len(d_428_v8_), (self).f20, p1, globalState))
                    nw63_[int(22)] = p0
                    nw63_[int(23)] = p0
                    d_430_v10_ = nw63_
                    d_431_v11_: C1
                    nw64_ = C1()
                    nw64_.ctor__(-261, d_430_v10_, 556, (self).f21)
                    d_431_v11_ = nw64_
                    (globalState).f14 = not(not(p1))
                    if (d_429_v9_)[default__.safeIndex(d_431_v11_.f27, len(d_429_v9_))]:
                        r1 = False
                        d_432_v12_: str
                        d_432_v12_ = _dafny.CodePoint('w')
                        d_432_v12_ = _dafny.CodePoint('i')
                        (globalState).f14 = (p0) == (p0)
                        d_433_v13_: _dafny.Array
                        def lambda15_(d_434_v8_):
                            def lambda16_(d_435_i2_):
                                return d_434_v8_

                            return lambda16_

                        init8_ = lambda15_(d_428_v8_)
                        nw65_ = _dafny.Array(None, 1)
                        for i0_8_ in range(nw65_.length(0)):
                            nw65_[i0_8_] = init8_(i0_8_)
                        d_433_v13_ = nw65_
                        index58_ = default__.safeIndex(60, (d_433_v13_).length(0))
                        (d_433_v13_)[index58_] = d_428_v8_
                        index59_ = default__.safeIndex(60, (d_433_v13_).length(0))
                        (d_433_v13_)[index59_] = d_428_v8_
                        (globalState).f17 = (d_431_v11_.f27) + (d_431_v11_.f27)
                    elif True:
                        d_436_v14_: int
                        d_437_v15_: bool
                        d_438_v16_: int
                        d_439_v17_: int
                        out6_: int
                        out7_: bool
                        out8_: int
                        out9_: int
                        out6_, out7_, out8_, out9_ = (self).m3(((self).f21 if p1 else (self).f21), globalState)
                        d_436_v14_ = out6_
                        d_437_v15_ = out7_
                        d_438_v16_ = out8_
                        d_439_v17_ = out9_
                        d_440_v18_: _dafny.Set
                        d_440_v18_ = _dafny.Set({d_428_v8_, (d_428_v8_) + (d_428_v8_), d_428_v8_, d_428_v8_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_441_i3_ in range(default__.abs(24))])})
                        d_440_v18_ = d_440_v18_
                        (globalState).f5 = d_438_v16_
                        d_442_v19_: _dafny.MultiSet
                        d_442_v19_ = _dafny.MultiSet([p1, (self).f21, d_437_v15_])
                        d_443_v20_: _dafny.Seq
                        d_443_v20_ = _dafny.SeqWithoutIsStrInference([d_431_v11_.f27, d_438_v16_])
                        d_444_v21_: C1
                        nw66_ = C1()
                        nw66_.ctor__(d_431_v11_.f27, (d_431_v11_).f28, (d_431_v11_).fm1(len(d_429_v9_), d_431_v11_.f27, False, globalState), (d_431_v11_).fm16(d_442_v19_, d_443_v20_, globalState))
                        d_444_v21_ = nw66_
                        d_445_v22_: D7
                        d_445_v22_ = D7_DC17(not(p1))
                        (globalState).f4 = (d_445_v22_).cf25
                    if not ((d_431_v11_.f27) != (p0)) or (True):
                        index60_ = default__.safeIndex(851, ((d_431_v11_).f28).length(0))
                        ((d_431_v11_).f28)[index60_] = d_431_v11_.f27
                        index61_ = default__.safeIndex(851, ((d_431_v11_).f28).length(0))
                        ((d_431_v11_).f28)[index61_] = d_431_v11_.f27
                        d_446_v23_: _dafny.Set
                        d_446_v23_ = _dafny.Set({p0, p0, d_431_v11_.f27, d_431_v11_.f27, p0})
                        d_447_v24_: int
                        d_448_v25_: bool
                        d_449_v26_: int
                        d_450_v27_: int
                        out10_: int
                        out11_: bool
                        out12_: int
                        out13_: int
                        out10_, out11_, out12_, out13_ = (self).m3((len(d_446_v23_)) < (d_431_v11_.f27), globalState)
                        d_447_v24_ = out10_
                        d_448_v25_ = out11_
                        d_449_v26_ = out12_
                        d_450_v27_ = out13_
                        d_451_v28_: _dafny.Array
                        nw67_ = _dafny.Array(False, 29)
                        d_451_v28_ = nw67_
                        d_452_v29_: _dafny.MultiSet
                        d_452_v29_ = _dafny.MultiSet([False])
                        index62_ = default__.safeIndex(638, (d_451_v28_).length(0))
                        (d_451_v28_)[index62_] = (d_452_v29_).issubset(_dafny.MultiSet([(self).f21, d_448_v25_]))
                        index63_ = default__.safeIndex(638, (d_451_v28_).length(0))
                        (d_451_v28_)[index63_] = (((d_431_v11_).f28)[default__.safeIndex(851, ((d_431_v11_).f28).length(0))]) < (d_447_v24_)
                        d_453_v30_: C1
                        nw68_ = C1()
                        nw68_.ctor__(d_431_v11_.f27, (d_431_v11_).f28, default__.safeDivisionInt((0) - (((d_431_v11_).f28)[default__.safeIndex(851, ((d_431_v11_).f28).length(0))]), d_431_v11_.f27), (((d_431_v11_).f28)[default__.safeIndex(851, ((d_431_v11_).f28).length(0))]) == (992))
                        d_453_v30_ = nw68_
                        d_454_v31_: _dafny.MultiSet
                        d_454_v31_ = _dafny.MultiSet([default__.safeDivisionInt((d_431_v11_).fm1(d_449_v26_, len(d_446_v23_), True, globalState), (0) - (len(d_428_v8_))), len(d_428_v8_)])
                        d_455_v32_: D10
                        d_455_v32_ = D10_DC25(d_454_v31_)
                        index64_ = default__.safeIndex(851, ((d_431_v11_).f28).length(0))
                        rhs55_ = (self).f21
                        rhs56_ = ((d_455_v32_).cf34).intersection(d_454_v31_)
                        rhs57_ = (0) - (d_447_v24_)
                        rhs58_ = p1
                        lhs45_ = globalState
                        lhs46_ = (d_431_v11_).f28
                        lhs47_ = default__.safeIndex(851, ((d_431_v11_).f28).length(0))
                        lhs48_ = globalState
                        lhs45_.f18 = rhs55_
                        d_454_v31_ = rhs56_
                        lhs46_[lhs47_] = rhs57_
                        lhs48_.f18 = rhs58_
                    elif True:
                        (d_431_v11_).f27 = len(d_429_v9_)
                        d_456_v33_: bool
                        d_456_v33_ = not((d_428_v8_) <= (d_428_v8_))
                        d_456_v33_ = d_456_v33_
                        d_457_v34_: _dafny.Map
                        d_457_v34_ = _dafny.Map({150: p0})
                        d_458_v35_: _dafny.Map
                        d_458_v35_ = _dafny.Map({p1: (self).f21})
                        d_459_v36_: _dafny.Seq
                        d_459_v36_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f21: (self).f21}), d_458_v35_, d_458_v35_])
                        (globalState).f3 = (not((self).fm0(d_457_v34_, d_428_v8_, d_459_v36_, globalState)) if (self).f21 else (self).f21)
                        d_460_v37_: _dafny.Map
                        d_460_v37_ = _dafny.Map({(self).f21: d_431_v11_.f27})
                        d_460_v37_ = ((default__.fm27((self).f20, globalState)) | (d_460_v37_)) | (p2)
                        d_461_v38_: C1
                        nw69_ = C1()
                        nw69_.ctor__((d_431_v11_.f27) * (687), (d_431_v11_).f28, p0, (self).f21)
                        d_461_v38_ = nw69_
                    pass
            pass
        r1 = ((self).f21 if not(not((self).f21)) else p1)
        d_462_v39_: _dafny.Map
        d_462_v39_ = _dafny.Map({(self).f20: (self).f21})
        d_463_v40_: str
        d_463_v40_ = _dafny.CodePoint('k')
        d_464_v41_: _dafny.Seq
        d_464_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuposfy"))
        d_465_v42_: _dafny.MultiSet
        d_465_v42_ = _dafny.MultiSet([p0])
        d_466_v43_: _dafny.Seq
        d_466_v43_ = _dafny.SeqWithoutIsStrInference([p0])
        d_467_v44_: _dafny.Set
        d_467_v44_ = _dafny.Set({862, p0})
        d_468_v45_: _dafny.Map
        d_468_v45_ = _dafny.Map({d_463_v40_: len(d_467_v44_)})
        d_469_v46_: _dafny.Array
        def lambda17_(d_470_i4_):
            return (d_470_i4_) + ((self).f20)

        init9_ = lambda17_
        nw70_ = _dafny.Array(None, 10)
        for i0_9_ in range(nw70_.length(0)):
            nw70_[i0_9_] = init9_(i0_9_)
        d_469_v46_ = nw70_
        d_471_v47_: C1
        nw71_ = C1()
        nw71_.ctor__(len(d_468_v45_), d_469_v46_, p0, p1)
        d_471_v47_ = nw71_
        d_472_v48_: _dafny.Set
        d_472_v48_ = _dafny.Set({d_471_v47_, d_471_v47_})
        d_473_v49_: _dafny.Map
        d_473_v49_ = _dafny.Map({p1: (self).f21})
        d_474_v50_: _dafny.Array
        nw72_ = _dafny.Array(None, 28)
        nw72_[int(0)] = len(d_464_v41_)
        nw72_[int(1)] = p0
        nw72_[int(2)] = (d_465_v42_).cardinality
        nw72_[int(3)] = (self).f20
        nw72_[int(4)] = (d_466_v43_)[default__.safeIndex(p0, len(d_466_v43_))]
        nw72_[int(5)] = p0
        nw72_[int(6)] = (self).f20
        nw72_[int(7)] = p0
        nw72_[int(8)] = (self).f20
        nw72_[int(9)] = len(d_462_v39_)
        nw72_[int(10)] = p0
        nw72_[int(11)] = len(_dafny.SeqWithoutIsStrInference([len(d_472_v48_)]))
        nw72_[int(12)] = p0
        nw72_[int(13)] = d_471_v47_.f27
        nw72_[int(14)] = len(d_473_v49_)
        nw72_[int(15)] = d_471_v47_.f27
        nw72_[int(16)] = d_471_v47_.f27
        nw72_[int(17)] = len(d_464_v41_)
        nw72_[int(18)] = (self).f20
        nw72_[int(19)] = d_471_v47_.f27
        nw72_[int(20)] = (self).f20
        nw72_[int(21)] = d_471_v47_.f27
        nw72_[int(22)] = d_471_v47_.f27
        nw72_[int(23)] = (self).f20
        nw72_[int(24)] = p0
        nw72_[int(25)] = (self).f20
        nw72_[int(26)] = len(d_468_v45_)
        nw72_[int(27)] = len(d_466_v43_)
        d_474_v50_ = nw72_
        d_475_v51_: C1
        nw73_ = C1()
        nw73_.ctor__((D10_DC26(545)).cf35, d_469_v46_, p0, p1)
        d_475_v51_ = nw73_
        d_476_v52_: _dafny.Map
        d_476_v52_ = _dafny.Map({d_463_v40_: d_471_v47_})
        d_477_v53_: _dafny.Seq
        d_477_v53_ = _dafny.SeqWithoutIsStrInference([d_476_v52_, d_476_v52_, d_476_v52_])
        d_478_v54_: D10
        d_478_v54_ = D10_DC27((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([len(d_462_v39_)])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([len(d_462_v39_)]))), p0))).cardinality, d_477_v53_, default__.fm17((self).f21, globalState))
        d_479_v55_: _dafny.Seq
        d_479_v55_ = d_466_v43_
        pat_let_tv12_ = d_475_v51_
        pat_let_tv13_ = p1
        pat_let_tv14_ = d_479_v55_
        pat_let_tv15_ = globalState
        pat_let_tv16_ = globalState
        def iife53_(_pat_let17_0):
            def iife54_(d_480_dt__update__tmp_h0_):
                def iife55_(_pat_let18_0):
                    def iife56_(d_481_dt__update_hcf36_h0_):
                        def iife57_(_pat_let19_0):
                            def iife58_(d_482_dt__update_hcf38_h0_):
                                return D10_DC27(d_481_dt__update_hcf36_h0_, (d_480_dt__update__tmp_h0_).cf37, d_482_dt__update_hcf38_h0_)
                            return iife58_(_pat_let19_0)
                        return iife57_(default__.fm17((pat_let_tv12_).fm16(_dafny.MultiSet([True, not(pat_let_tv13_)]), pat_let_tv14_, pat_let_tv15_), pat_let_tv16_))
                    return iife56_(_pat_let18_0)
                return iife55_((self).f20)
            return iife54_(_pat_let17_0)
        source7_ = iife53_(d_478_v54_)
        if source7_.is_DC26:
            d_483___mcc_h0_ = source7_.cf35
            d_484_cf35_ = d_483___mcc_h0_
            (globalState).f18 = (350) <= (p0)
            d_485_v56_: _dafny.Array
            nw74_ = _dafny.Array(False, 12)
            d_485_v56_ = nw74_
            d_486_v57_: _dafny.Map
            d_486_v57_ = _dafny.Map({(0) - (d_471_v47_.f27): d_471_v47_.f27})
            d_487_v58_: _dafny.Seq
            d_487_v58_ = _dafny.SeqWithoutIsStrInference([d_473_v49_])
            index65_ = default__.safeIndex(441, (d_485_v56_).length(0))
            (d_485_v56_)[index65_] = (self).fm0(d_486_v57_, d_464_v41_, (d_487_v58_).set(default__.safeIndex(d_471_v47_.f27, len(d_487_v58_)), _dafny.Map({(self).f21: p1})), globalState)
            d_488_v59_: _dafny.Set
            d_488_v59_ = _dafny.Set({(self).f21, p1})
            index66_ = default__.safeIndex(254, ((d_475_v51_).f28).length(0))
            ((d_475_v51_).f28)[index66_] = (default__.fm24(d_471_v47_.f27, not(p1), d_488_v59_, p1, globalState)).cf17
            index67_ = default__.safeIndex(292, (d_485_v56_).length(0))
            (d_485_v56_)[index67_] = p1
            index68_ = default__.safeIndex(89, (d_469_v46_).length(0))
            (d_469_v46_)[index68_] = d_475_v51_.f27
            d_489_v60_: _dafny.MultiSet
            d_489_v60_ = _dafny.MultiSet([d_485_v56_, d_485_v56_])
            d_490_v61_: _dafny.Map
            d_490_v61_ = _dafny.Map({d_467_v44_: d_475_v51_})
            index69_ = default__.safeIndex(441, (d_485_v56_).length(0))
            index70_ = default__.safeIndex(254, ((d_475_v51_).f28).length(0))
            index71_ = default__.safeIndex(292, (d_485_v56_).length(0))
            index72_ = default__.safeIndex(89, (d_469_v46_).length(0))
            rhs59_ = (self).f21
            rhs60_ = d_484_cf35_
            rhs61_ = (d_490_v61_) != (d_490_v61_)
            rhs62_ = d_475_v51_.f27
            rhs63_ = ((d_489_v60_) | (d_489_v60_)).intersection(d_489_v60_)
            lhs49_ = d_485_v56_
            lhs50_ = default__.safeIndex(441, (d_485_v56_).length(0))
            lhs51_ = (d_475_v51_).f28
            lhs52_ = default__.safeIndex(254, ((d_475_v51_).f28).length(0))
            lhs53_ = d_485_v56_
            lhs54_ = default__.safeIndex(292, (d_485_v56_).length(0))
            lhs55_ = d_469_v46_
            lhs56_ = default__.safeIndex(89, (d_469_v46_).length(0))
            lhs49_[lhs50_] = rhs59_
            lhs51_[lhs52_] = rhs60_
            lhs53_[lhs54_] = rhs61_
            lhs55_[lhs56_] = rhs62_
            d_489_v60_ = rhs63_
            d_491_v62_: _dafny.Array
            nw75_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
            d_491_v62_ = nw75_
            index73_ = default__.safeIndex(390, (d_491_v62_).length(0))
            (d_491_v62_)[index73_] = d_464_v41_
            index74_ = default__.safeIndex(390, (d_491_v62_).length(0))
            index75_ = default__.safeIndex(254, ((d_475_v51_).f28).length(0))
            rhs64_ = d_464_v41_
            rhs65_ = default__.safeModuloInt(d_471_v47_.f27, d_471_v47_.f27)
            rhs66_ = not((self).f21)
            lhs57_ = d_491_v62_
            lhs58_ = default__.safeIndex(390, (d_491_v62_).length(0))
            lhs59_ = (d_475_v51_).f28
            lhs60_ = default__.safeIndex(254, ((d_475_v51_).f28).length(0))
            lhs61_ = globalState
            lhs57_[lhs58_] = rhs64_
            lhs59_[lhs60_] = rhs65_
            lhs61_.f3 = rhs66_
            if ((d_473_v49_)[(d_485_v56_)[default__.safeIndex(441, (d_485_v56_).length(0))]] if ((d_485_v56_)[default__.safeIndex(441, (d_485_v56_).length(0))]) in (d_473_v49_) else p1):
                d_492_v64_: _dafny.Array
                def lambda18_(d_493_v44_, d_494_p2_, d_495_v57_):
                    def lambda19_(d_496_i5_):
                        def iife59_():
                            coll19_ = _dafny.Set()
                            compr_19_: int
                            for compr_19_ in (d_495_v57_).keys.Elements:
                                d_497_v63_: int = compr_19_
                                if (d_497_v63_) in (d_495_v57_):
                                    coll19_ = coll19_.union(_dafny.Set([(d_497_v63_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afyfapfi"))))]))
                            return _dafny.Set(coll19_)
                        return _dafny.SeqWithoutIsStrInference([d_493_v44_, iife59_()
                        , _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ml"))), ((d_494_p2_)[(self).f21] if ((self).f21) in (d_494_p2_) else (self).f20)})])

                    return lambda19_

                init10_ = lambda18_(d_467_v44_, p2, d_486_v57_)
                nw76_ = _dafny.Array(None, 8)
                for i0_10_ in range(nw76_.length(0)):
                    nw76_[i0_10_] = init10_(i0_10_)
                d_492_v64_ = nw76_
                d_498_v65_: _dafny.MultiSet
                d_498_v65_ = _dafny.MultiSet([p1, True])
                index76_ = default__.safeIndex(666, (d_492_v64_).length(0))
                (d_492_v64_)[index76_] = _dafny.SeqWithoutIsStrInference([_dafny.Set({((d_498_v65_)[(self).f21] if ((self).f21) in (d_498_v65_) else 684), (d_475_v51_).fm1(len(_dafny.SeqWithoutIsStrInference([d_484_cf35_ for d_499_i6_ in range(default__.abs(-418))])), -357, (d_485_v56_)[default__.safeIndex(441, (d_485_v56_).length(0))], globalState)}), d_467_v44_])
                d_500_v66_: _dafny.Seq
                d_500_v66_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({len(p2), (0) - (d_484_cf35_), ((d_475_v51_).f28)[default__.safeIndex(254, ((d_475_v51_).f28).length(0))], len(_dafny.SeqWithoutIsStrInference([d_463_v40_ for d_501_i7_ in range(default__.abs(748))]))})])
                d_502_v67_: _dafny.Seq
                d_502_v67_ = _dafny.SeqWithoutIsStrInference([d_500_v66_, d_500_v66_, d_500_v66_, _dafny.SeqWithoutIsStrInference([(d_500_v66_)[default__.safeIndex(d_475_v51_.f27, len(d_500_v66_))]]), d_500_v66_])
                index77_ = default__.safeIndex(666, (d_492_v64_).length(0))
                (d_492_v64_)[index77_] = (((d_502_v67_)[default__.safeIndex(d_475_v51_.f27, len(d_502_v67_))]).set(default__.safeIndex(d_471_v47_.f27, len((d_502_v67_)[default__.safeIndex(d_475_v51_.f27, len(d_502_v67_))])), _dafny.Set({len((d_491_v62_)[default__.safeIndex(390, (d_491_v62_).length(0))])}))) + (d_500_v66_)
                d_503_v68_: _dafny.Map
                d_503_v68_ = _dafny.Map({d_485_v56_: (d_484_cf35_) + (len(d_467_v44_))})
                d_503_v68_ = (d_503_v68_).set(d_485_v56_, d_471_v47_.f27)
                index78_ = default__.safeIndex(390, (d_491_v62_).length(0))
                (d_491_v62_)[index78_] = d_464_v41_
                d_504_v69_: _dafny.MultiSet
                d_504_v69_ = _dafny.MultiSet([default__.fm17(p1, globalState)])
                (globalState).f5 = (d_471_v47_.f27) - ((len(d_462_v39_)) * (((d_504_v69_)[d_463_v40_] if (d_463_v40_) in (d_504_v69_) else (self).f20)))
                d_505_v70_: _dafny.Seq
                d_505_v70_ = _dafny.SeqWithoutIsStrInference([p1, not((d_475_v51_.f27) > (d_475_v51_.f27)), (d_498_v65_).issubset(d_498_v65_)])
                (globalState).f18 = (d_505_v70_)[default__.safeIndex(-639, len(d_505_v70_))]
            elif True:
                (globalState).f5 = (d_466_v43_)[default__.safeIndex(((d_469_v46_)[default__.safeIndex(89, (d_469_v46_).length(0))]) - (d_484_cf35_), len(d_466_v43_))]
                d_506_v71_: _dafny.Seq
                d_506_v71_ = _dafny.SeqWithoutIsStrInference([(d_471_v47_).f28])
                index79_ = default__.safeIndex(441, (d_485_v56_).length(0))
                (d_485_v56_)[index79_] = (len(d_506_v71_)) == (d_475_v51_.f27)
                index80_ = default__.safeIndex(89, (d_469_v46_).length(0))
                (d_469_v46_)[index80_] = (d_475_v51_).fm1(d_475_v51_.f27, d_484_cf35_, not(p1), globalState)
                d_507_v72_: D9
                d_507_v72_ = D9_DC23((d_469_v46_)[default__.safeIndex(89, (d_469_v46_).length(0))], p1)
                d_508_v73_: D9
                d_508_v73_ = D9_DC24(d_507_v72_)
                d_509_v74_: D9
                d_509_v74_ = D9_DC24(d_508_v73_)
                d_509_v74_ = d_509_v74_
                d_510_v75_: C1
                nw77_ = C1()
                nw77_.ctor__((0) - (d_471_v47_.f27), (d_475_v51_).f28, len(((d_491_v62_)[default__.safeIndex(390, (d_491_v62_).length(0))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ien")))), (d_485_v56_)[default__.safeIndex(441, (d_485_v56_).length(0))])
                d_510_v75_ = nw77_
        elif source7_.is_DC27:
            d_511___mcc_h1_ = source7_.cf36
            d_512___mcc_h2_ = source7_.cf37
            d_513___mcc_h3_ = source7_.cf38
            d_514_cf38_ = d_513___mcc_h3_
            d_515_cf37_ = d_512___mcc_h2_
            d_516_cf36_ = d_511___mcc_h1_
            (globalState).f18 = p1
            d_517_v76_: _dafny.Array
            nw78_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
            d_517_v76_ = nw78_
            index81_ = default__.safeIndex(35, (d_517_v76_).length(0))
            (d_517_v76_)[index81_] = ((d_464_v41_).set(default__.safeIndex(len(d_464_v41_), len(d_464_v41_)), _dafny.CodePoint('a'))) + (d_464_v41_)
            d_518_v77_: D6
            d_518_v77_ = D6_DC13(d_464_v41_)
            index82_ = default__.safeIndex(35, (d_517_v76_).length(0))
            (d_517_v76_)[index82_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndkkvypt"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dat")))) + ((d_518_v77_).cf21)
            d_519_v78_: _dafny.Seq
            d_519_v78_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, (self).f21])
            d_520_v79_: _dafny.MultiSet
            d_520_v79_ = _dafny.MultiSet([not((d_519_v78_)[default__.safeIndex(len(d_466_v43_), len(d_519_v78_))]), (self).f21])
            d_521_v80_: _dafny.Seq
            d_521_v80_ = _dafny.SeqWithoutIsStrInference([(d_475_v51_).fm16(d_520_v79_, _dafny.SeqWithoutIsStrInference([d_516_cf36_, d_475_v51_.f27]), globalState), p1])
            d_522_v81_: _dafny.Seq
            d_522_v81_ = _dafny.SeqWithoutIsStrInference([d_473_v49_, (d_473_v49_).set((self).f21, p1), (d_473_v49_ if not((d_521_v80_)[default__.safeIndex(856, len(d_521_v80_))]) else d_473_v49_)])
            d_523_v82_: D7
            d_523_v82_ = D7_DC18(d_466_v43_, d_471_v47_.f27)
            d_524_v83_: _dafny.Map
            d_524_v83_ = _dafny.Map({d_523_v82_: (True) and (p1)})
            rhs67_ = ((_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f21: (self).f21}), d_473_v49_])) + (d_522_v81_)).set(default__.safeIndex((0) - (d_471_v47_.f27), len((_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f21: (self).f21}), d_473_v49_])) + (d_522_v81_))), d_473_v49_)
            rhs68_ = ((d_524_v83_)[d_523_v82_] if (d_523_v82_) in (d_524_v83_) else p1)
            rhs69_ = len(d_466_v43_)
            lhs62_ = globalState
            lhs63_ = d_471_v47_
            d_522_v81_ = rhs67_
            lhs62_.f3 = rhs68_
            lhs63_.f27 = rhs69_
            d_525_v84_: _dafny.Array
            def lambda20_(d_526_i8_):
                return (self).f21

            init11_ = lambda20_
            nw79_ = _dafny.Array(None, 17)
            for i0_11_ in range(nw79_.length(0)):
                nw79_[i0_11_] = init11_(i0_11_)
            d_525_v84_ = nw79_
            index83_ = default__.safeIndex(676, (d_525_v84_).length(0))
            (d_525_v84_)[index83_] = (self).f21
            d_527_v85_: _dafny.Map
            d_527_v85_ = _dafny.Map({(self).f20: d_471_v47_.f27})
            index84_ = default__.safeIndex(676, (d_525_v84_).length(0))
            (d_525_v84_)[index84_] = (d_475_v51_).fm0(d_527_v85_, d_464_v41_, d_522_v81_, globalState)
        elif True:
            d_528___mcc_h4_ = source7_.cf34
            d_529_cf34_ = d_528___mcc_h4_
            d_530_v86_: _dafny.Array
            def lambda21_(d_531_i9_):
                return (self).f21

            init12_ = lambda21_
            nw80_ = _dafny.Array(None, 16)
            for i0_12_ in range(nw80_.length(0)):
                nw80_[i0_12_] = init12_(i0_12_)
            d_530_v86_ = nw80_
            d_532_v87_: _dafny.MultiSet
            d_532_v87_ = _dafny.MultiSet([d_530_v86_, d_530_v86_])
            source8_ = D2_DC2((_dafny.MultiSet([d_530_v86_])) - (d_532_v87_))
            if source8_.is_DC3:
                d_533___mcc_h5_ = source8_.cf3
                d_534___mcc_h6_ = source8_.cf4
                d_535___mcc_h7_ = source8_.cf5
                d_536___mcc_h8_ = source8_.cf6
                d_537___mcc_h9_ = source8_.cf7
                d_538_cf7_ = d_537___mcc_h9_
                d_539_cf6_ = d_536___mcc_h8_
                d_540_cf5_ = d_535___mcc_h7_
                d_541_cf4_ = d_534___mcc_h6_
                d_542_cf3_ = d_533___mcc_h5_
                (globalState).f4 = d_541_cf4_
                (d_471_v47_).f27 = d_475_v51_.f27
                d_539_cf6_ = (d_471_v47_.f27) + (default__.safeDivisionInt(d_475_v51_.f27, d_471_v47_.f27))
                d_543_v88_: C0
                nw81_ = C0()
                nw81_.ctor__()
                d_543_v88_ = nw81_
            elif source8_.is_DC2:
                d_544___mcc_h10_ = source8_.cf2
                d_545_cf2_ = d_544___mcc_h10_
                d_546_v89_: _dafny.Seq
                d_546_v89_ = _dafny.SeqWithoutIsStrInference([p1])
                d_547_v90_: _dafny.Seq
                d_547_v90_ = _dafny.SeqWithoutIsStrInference([d_473_v49_])
                d_548_v91_: _dafny.MultiSet
                d_548_v91_ = _dafny.MultiSet([(self).fm0(_dafny.Map({(d_475_v51_).fm1((_dafny.MultiSet([d_471_v47_.f27])).cardinality, (self).fm1(len(d_546_v89_), d_471_v47_.f27, (self).f21, globalState), p1, globalState): len(d_546_v89_)}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "folerbj")), d_547_v90_, globalState), (self).f21])
                rhs70_ = ((self).f20) * ((0) - (((d_548_v91_)[p1] if (p1) in (d_548_v91_) else d_475_v51_.f27)))
                rhs71_ = p1
                lhs64_ = d_471_v47_
                lhs64_.f27 = rhs70_
                r1 = rhs71_
                d_473_v49_ = (d_473_v49_).set((p1 if p1 else False), (self).f21)
                (globalState).f18 = not(not(p1))
                d_549_v92_: C1
                nw82_ = C1()
                nw82_.ctor__((self).f20, (d_471_v47_).f28, default__.safeDivisionInt((self).f20, (self).f20), p1)
                d_549_v92_ = nw82_
            elif True:
                d_550___mcc_h11_ = source8_.cf8
                d_551_cf8_ = d_550___mcc_h11_
                d_552_v93_: C1
                nw83_ = C1()
                nw83_.ctor__(d_475_v51_.f27, d_469_v46_, (d_475_v51_.f27) * ((self).f20), (self).f21)
                d_552_v93_ = nw83_
                d_553_v94_: _dafny.Array
                nw84_ = _dafny.Array(None, 28)
                nw84_[int(0)] = (d_471_v47_).f28
                nw84_[int(1)] = (d_552_v93_).f28
                nw84_[int(2)] = d_474_v50_
                nw84_[int(3)] = (d_475_v51_).f28
                nw84_[int(4)] = (d_475_v51_).f28
                nw84_[int(5)] = d_474_v50_
                nw84_[int(6)] = d_474_v50_
                nw84_[int(7)] = (d_552_v93_).f28
                nw84_[int(8)] = (d_552_v93_).f28
                nw84_[int(9)] = (d_471_v47_).f28
                nw84_[int(10)] = (d_471_v47_).f28
                nw84_[int(11)] = (d_552_v93_).f28
                nw84_[int(12)] = d_469_v46_
                nw84_[int(13)] = (d_475_v51_).f28
                nw84_[int(14)] = d_474_v50_
                nw84_[int(15)] = (d_471_v47_).f28
                nw84_[int(16)] = (d_471_v47_).f28
                nw84_[int(17)] = (d_552_v93_).f28
                nw84_[int(18)] = d_474_v50_
                nw84_[int(19)] = (d_552_v93_).f28
                nw84_[int(20)] = d_474_v50_
                nw84_[int(21)] = (d_475_v51_).f28
                nw84_[int(22)] = d_474_v50_
                nw84_[int(23)] = d_469_v46_
                nw84_[int(24)] = (d_471_v47_).f28
                nw84_[int(25)] = (d_471_v47_).f28
                nw84_[int(26)] = d_469_v46_
                nw84_[int(27)] = (d_475_v51_).f28
                d_553_v94_ = nw84_
                d_554_v95_: _dafny.Seq
                d_554_v95_ = _dafny.SeqWithoutIsStrInference([d_553_v94_])
                d_555_v96_: _dafny.Array
                nw85_ = _dafny.Array(None, 29)
                nw85_[int(0)] = (d_554_v95_)[default__.safeIndex(((d_529_cf34_)[d_471_v47_.f27] if (d_471_v47_.f27) in (d_529_cf34_) else 592), len(d_554_v95_))]
                nw85_[int(1)] = d_553_v94_
                nw85_[int(2)] = d_553_v94_
                nw85_[int(3)] = d_553_v94_
                nw85_[int(4)] = d_553_v94_
                nw85_[int(5)] = d_553_v94_
                nw85_[int(6)] = d_553_v94_
                nw85_[int(7)] = d_553_v94_
                nw85_[int(8)] = d_553_v94_
                nw85_[int(9)] = d_553_v94_
                nw85_[int(10)] = d_553_v94_
                nw85_[int(11)] = (d_553_v94_ if p1 else d_553_v94_)
                nw85_[int(12)] = d_553_v94_
                nw85_[int(13)] = d_553_v94_
                nw85_[int(14)] = d_553_v94_
                nw85_[int(15)] = d_553_v94_
                nw85_[int(16)] = (d_554_v95_)[default__.safeIndex(d_471_v47_.f27, len(d_554_v95_))]
                nw85_[int(17)] = d_553_v94_
                nw85_[int(18)] = d_553_v94_
                nw85_[int(19)] = d_553_v94_
                nw85_[int(20)] = d_553_v94_
                nw85_[int(21)] = d_553_v94_
                nw85_[int(22)] = d_553_v94_
                nw85_[int(23)] = d_553_v94_
                nw85_[int(24)] = d_553_v94_
                nw85_[int(25)] = d_553_v94_
                nw85_[int(26)] = d_553_v94_
                nw85_[int(27)] = d_553_v94_
                nw85_[int(28)] = d_553_v94_
                d_555_v96_ = nw85_
                index85_ = default__.safeIndex(411, (d_555_v96_).length(0))
                (d_555_v96_)[index85_] = d_553_v94_
                index86_ = default__.safeIndex(411, (d_555_v96_).length(0))
                (d_555_v96_)[index86_] = d_553_v94_
                d_556_v97_: D7
                d_556_v97_ = D7_DC17((p1))
                d_556_v97_ = d_556_v97_
                d_557_v98_: _dafny.Array
                def lambda22_(d_558_v41_):
                    def lambda23_(d_559_i10_):
                        return d_558_v41_

                    return lambda23_

                init13_ = lambda22_(d_464_v41_)
                nw86_ = _dafny.Array(None, 22)
                for i0_13_ in range(nw86_.length(0)):
                    nw86_[i0_13_] = init13_(i0_13_)
                d_557_v98_ = nw86_
                index87_ = default__.safeIndex(433, (d_557_v98_).length(0))
                (d_557_v98_)[index87_] = d_464_v41_
                index88_ = default__.safeIndex(433, (d_557_v98_).length(0))
                (d_557_v98_)[index88_] = d_464_v41_
            (globalState).f5 = (self).f20
            if not ((self).f21) or ((self).f21):
                d_560_v99_: _dafny.Array
                nw87_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_560_v99_ = nw87_
                d_561_v100_: _dafny.Array
                nw88_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_561_v100_ = nw88_
                index89_ = default__.safeIndex(374, (d_560_v99_).length(0))
                (d_560_v99_)[index89_] = d_561_v100_
                index90_ = default__.safeIndex(374, (d_560_v99_).length(0))
                (d_560_v99_)[index90_] = d_561_v100_
                d_562_v101_: D8
                d_562_v101_ = D8_DC20()
                index91_ = default__.safeIndex(986, (d_561_v100_).length(0))
                (d_561_v100_)[index91_] = (d_475_v51_).f28
                d_563_v102_: _dafny.Seq
                d_563_v102_ = _dafny.SeqWithoutIsStrInference([(d_475_v51_).fm16(_dafny.MultiSet([p1]), d_479_v55_, globalState), not(p1)])
                index92_ = default__.safeIndex(986, (d_561_v100_).length(0))
                rhs72_ = d_562_v101_
                rhs73_ = p1
                rhs74_ = (d_471_v47_).f28
                rhs75_ = (self).f21
                rhs76_ = d_563_v102_
                lhs65_ = globalState
                lhs66_ = d_561_v100_
                lhs67_ = default__.safeIndex(986, (d_561_v100_).length(0))
                lhs68_ = globalState
                d_562_v101_ = rhs72_
                lhs65_.f18 = rhs73_
                lhs66_[lhs67_] = rhs74_
                lhs68_.f3 = rhs75_
                d_563_v102_ = rhs76_
                index93_ = default__.safeIndex(921, (d_469_v46_).length(0))
                (d_469_v46_)[index93_] = (d_475_v51_.f27) + (d_471_v47_.f27)
                d_564_v103_: D7
                d_564_v103_ = D7_DC17(not((self).f21))
                index94_ = default__.safeIndex(921, (d_469_v46_).length(0))
                (d_469_v46_)[index94_] = default__.safeDivisionInt(d_471_v47_.f27, (d_475_v51_).fm1((self).f20, d_471_v47_.f27, not((d_564_v103_).cf25), globalState))
                (globalState).f5 = d_475_v51_.f27
                pat_let_tv17_ = d_477_v53_
                d_565_v104_: _dafny.Map
                def iife60_(_pat_let20_0):
                    def iife61_(d_566_dt__update__tmp_h1_):
                        def iife62_(_pat_let21_0):
                            def iife63_(d_567_dt__update_hcf37_h0_):
                                return D10_DC27((d_566_dt__update__tmp_h1_).cf36, d_567_dt__update_hcf37_h0_, (d_566_dt__update__tmp_h1_).cf38)
                            return iife63_(_pat_let21_0)
                        return iife62_(pat_let_tv17_)
                    return iife61_(_pat_let20_0)
                d_565_v104_ = _dafny.Map({iife60_(d_478_v54_): not(not((self).f21))})
                (d_475_v51_).f27 = (d_471_v47_.f27 if (self).f21 else len(d_565_v104_))
            elif True:
                (globalState).f4 = (self).f21
                d_568_v105_: _dafny.Array
                nw89_ = _dafny.Array(None, 10)
                nw89_[int(0)] = default__.fm28(d_475_v51_.f27, globalState)
                nw89_[int(1)] = d_466_v43_
                nw89_[int(2)] = d_466_v43_
                nw89_[int(3)] = d_466_v43_
                nw89_[int(4)] = d_466_v43_
                nw89_[int(5)] = (d_466_v43_).set(default__.safeIndex(d_475_v51_.f27, len(d_466_v43_)), d_471_v47_.f27)
                nw89_[int(6)] = d_466_v43_
                nw89_[int(7)] = (d_466_v43_) + (_dafny.SeqWithoutIsStrInference([d_475_v51_.f27 for d_569_i11_ in range(default__.abs(739))]))
                nw89_[int(8)] = d_466_v43_
                nw89_[int(9)] = _dafny.SeqWithoutIsStrInference([p0, (self).f20, d_471_v47_.f27])
                d_568_v105_ = nw89_
                d_570_v106_: _dafny.Map
                d_570_v106_ = _dafny.Map({d_475_v51_.f27: len(_dafny.SeqWithoutIsStrInference([d_463_v40_ for d_571_i12_ in range(default__.abs(512))]))})
                d_572_v107_: _dafny.Seq
                d_572_v107_ = _dafny.SeqWithoutIsStrInference([d_473_v49_, d_473_v49_])
                d_573_v108_: _dafny.Map
                d_573_v108_ = _dafny.Map({not((d_471_v47_).fm0(d_570_v106_, d_464_v41_, d_572_v107_, globalState)): _dafny.CodePoint('q')})
                rhs77_ = (d_475_v51_).fm1(d_475_v51_.f27, d_475_v51_.f27, p1, globalState)
                rhs78_ = default__.safeDivisionInt((0) - (p0), d_475_v51_.f27)
                rhs79_ = (p1 if (self).f21 else ((self).f20) >= (d_471_v47_.f27))
                rhs80_ = (d_568_v105_ if p1 else d_568_v105_)
                rhs81_ = d_573_v108_
                lhs69_ = globalState
                lhs70_ = d_475_v51_
                lhs71_ = globalState
                lhs69_.f17 = rhs77_
                lhs70_.f27 = rhs78_
                lhs71_.f18 = rhs79_
                d_568_v105_ = rhs80_
                d_573_v108_ = rhs81_
                (globalState).f18 = p1
                d_574_v109_: _dafny.Map
                d_574_v109_ = _dafny.Map({(self).f21: d_473_v49_})
                (globalState).f5 = ((len(d_574_v109_)) - (len(default__.fm26(p1, globalState)))) - (d_471_v47_.f27)
                d_575_v110_: _dafny.Array
                def lambda24_(d_576_p2_, d_577_p1_, d_578_v40_):
                    def lambda25_(d_579_i13_):
                        return (d_576_p2_).set(d_577_p1_, len(_dafny.SeqWithoutIsStrInference([d_578_v40_, d_578_v40_])))

                    return lambda25_

                init14_ = lambda24_(p2, p1, d_463_v40_)
                nw90_ = _dafny.Array(None, 13)
                for i0_14_ in range(nw90_.length(0)):
                    nw90_[i0_14_] = init14_(i0_14_)
                d_575_v110_ = nw90_
                def lambda26_(d_580_p2_):
                    def lambda27_(d_581_i14_):
                        return (d_580_p2_) | (d_580_p2_)

                    return lambda27_

                init15_ = lambda26_(p2)
                nw91_ = _dafny.Array(None, 13)
                for i0_15_ in range(nw91_.length(0)):
                    nw91_[i0_15_] = init15_(i0_15_)
                d_575_v110_ = nw91_
            (d_475_v51_).f27 = len(d_467_v44_)
        index95_ = default__.safeIndex(719, (d_469_v46_).length(0))
        (d_469_v46_)[index95_] = p0
        d_582_v111_: D6
        d_582_v111_ = D6_DC13(d_464_v41_)
        d_583_v112_: _dafny.Array
        nw92_ = _dafny.Array(None, 22)
        nw92_[int(0)] = (d_464_v41_).set(default__.safeIndex(d_475_v51_.f27, len(d_464_v41_)), d_463_v40_)
        nw92_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "miga"))
        nw92_[int(2)] = d_464_v41_
        nw92_[int(3)] = d_464_v41_
        nw92_[int(4)] = d_464_v41_
        nw92_[int(5)] = d_464_v41_
        nw92_[int(6)] = d_464_v41_
        nw92_[int(7)] = (d_464_v41_) + (d_464_v41_)
        nw92_[int(8)] = d_464_v41_
        nw92_[int(9)] = _dafny.SeqWithoutIsStrInference([d_463_v40_ for d_584_i15_ in range(default__.abs(807))])
        nw92_[int(10)] = d_464_v41_
        nw92_[int(11)] = d_464_v41_
        nw92_[int(12)] = default__.fm26(True, globalState)
        nw92_[int(13)] = d_464_v41_
        nw92_[int(14)] = d_464_v41_
        nw92_[int(15)] = (d_582_v111_).cf21
        nw92_[int(16)] = d_464_v41_
        nw92_[int(17)] = d_464_v41_
        nw92_[int(18)] = (d_464_v41_) + (d_464_v41_)
        nw92_[int(19)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_585_i16_ in range(default__.abs(619))])
        nw92_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qifnsfgu"))
        nw92_[int(21)] = _dafny.SeqWithoutIsStrInference([d_463_v40_ for d_586_i17_ in range(default__.abs(922))])
        d_583_v112_ = nw92_
        index96_ = default__.safeIndex(963, (d_583_v112_).length(0))
        (d_583_v112_)[index96_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwsydd"))) + (d_464_v41_)
        d_587_v113_: _dafny.Array
        nw93_ = _dafny.Array(False, 12)
        d_587_v113_ = nw93_
        index97_ = default__.safeIndex(524, (d_587_v113_).length(0))
        (d_587_v113_)[index97_] = (p0) == (d_475_v51_.f27)
        d_588_v114_: C0
        nw94_ = C0()
        nw94_.ctor__()
        d_588_v114_ = nw94_
        d_589_v115_: _dafny.MultiSet
        d_589_v115_ = _dafny.MultiSet([d_588_v114_])
        d_590_v116_: _dafny.Seq
        d_590_v116_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f21: (d_589_v115_).cardinality}), _dafny.Map({p1: d_475_v51_.f27})])
        d_591_v117_: _dafny.Seq
        d_591_v117_ = _dafny.SeqWithoutIsStrInference([d_464_v41_])
        d_592_v118_: D11
        d_592_v118_ = D11_DC28(d_583_v112_)
        d_593_v119_: _dafny.Map
        d_593_v119_ = _dafny.Map({144: len(d_464_v41_)})
        d_594_v120_: _dafny.Seq
        d_594_v120_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_471_v47_).fm16(_dafny.MultiSet([(self).f21]), d_479_v55_, globalState): p1}), d_473_v49_, d_473_v49_])
        index98_ = default__.safeIndex(719, (d_469_v46_).length(0))
        index99_ = default__.safeIndex(963, (d_583_v112_).length(0))
        index100_ = default__.safeIndex(524, (d_587_v113_).length(0))
        rhs82_ = len((d_590_v116_) + ((d_590_v116_) + (d_590_v116_)))
        rhs83_ = ((d_591_v117_)[default__.safeIndex(len(d_464_v41_), len(d_591_v117_))]) + ((d_464_v41_) + (d_464_v41_))
        rhs84_ = (d_592_v118_).cf39
        rhs85_ = (self).fm0(((d_593_v119_).set((self).f20, d_471_v47_.f27)) | (d_593_v119_), d_464_v41_, (d_594_v120_).set(default__.safeIndex(166, len(d_594_v120_)), d_473_v49_), globalState)
        lhs72_ = d_469_v46_
        lhs73_ = default__.safeIndex(719, (d_469_v46_).length(0))
        lhs74_ = d_583_v112_
        lhs75_ = default__.safeIndex(963, (d_583_v112_).length(0))
        lhs76_ = d_587_v113_
        lhs77_ = default__.safeIndex(524, (d_587_v113_).length(0))
        lhs72_[lhs73_] = rhs82_
        lhs74_[lhs75_] = rhs83_
        d_583_v112_ = rhs84_
        lhs76_[lhs77_] = rhs85_
        d_595_v122_: _dafny.Seq
        d_595_v122_ = _dafny.SeqWithoutIsStrInference([not((d_587_v113_)[default__.safeIndex(524, (d_587_v113_).length(0))]), p1])
        d_596_v123_: _dafny.Map
        d_596_v123_ = _dafny.Map({(d_595_v122_)[default__.safeIndex((self).fm1(d_471_v47_.f27, (_dafny.MultiSet(d_591_v117_)).cardinality, p1, globalState), len(d_595_v122_))]: d_593_v119_})
        d_597_v124_: _dafny.Array
        nw95_ = _dafny.Array(None, 11)
        nw95_[int(0)] = d_593_v119_
        nw95_[int(1)] = d_593_v119_
        nw95_[int(2)] = d_593_v119_
        nw95_[int(3)] = (d_593_v119_ if p1 else d_593_v119_)
        nw95_[int(4)] = (d_593_v119_).set(-384, d_471_v47_.f27)
        nw95_[int(5)] = d_593_v119_
        nw95_[int(6)] = (d_593_v119_) | (d_593_v119_)
        nw95_[int(7)] = d_593_v119_
        def iife64_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in _dafny.IntegerRange(-508, -87):
                d_598_v121_: int = compr_20_
                if ((-508) <= (d_598_v121_)) and ((d_598_v121_) < (-87)):
                    coll20_[default__.safeModuloInt(d_598_v121_, 392)] = d_475_v51_.f27
            return _dafny.Map(coll20_)
        nw95_[int(8)] = (d_593_v119_ if False else iife64_()
        )
        nw95_[int(9)] = (d_593_v119_) | (((d_596_v123_)[False] if (False) in (d_596_v123_) else d_593_v119_))
        nw95_[int(10)] = d_593_v119_
        d_597_v124_ = nw95_
        (globalState).f8 = d_597_v124_
        def iife65_():
            coll21_ = _dafny.Set()
            compr_21_: int
            for compr_21_ in _dafny.IntegerRange(72, 341):
                d_599_v125_: int = compr_21_
                if ((72) <= (d_599_v125_)) and ((d_599_v125_) < (341)):
                    coll21_ = coll21_.union(_dafny.Set([(d_599_v125_) * ((self).f20)]))
            return _dafny.Set(coll21_)
        r0 = ((d_467_v44_).ispropersubset(iife65_()
        )) and ((d_465_v42_) == (d_465_v42_))
        def iife66_():
            coll22_ = _dafny.Set()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(-389, -64):
                d_600_v126_: int = compr_22_
                if ((-389) <= (d_600_v126_)) and ((d_600_v126_) < (-64)):
                    coll22_ = coll22_.union(_dafny.Set([(d_600_v126_) - (d_471_v47_.f27)]))
            return _dafny.Set(coll22_)
        r1 = (d_467_v44_).ispropersubset((d_467_v44_).intersection(iife66_()
        ))
        return r0, r1


class C3(T0):
    def  __init__(self):
        self._f20: int = int(0)
        self._f21: bool = False
        self._f25: int = int(0)
        self._f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    def ctor__(self, f25, f26, f20, f21):
        (self)._f25 = f25
        (self)._f26 = f26
        (self)._f20 = f20
        (self)._f21 = f21

    def fm0(self, p0, p1, p2, globalState):
        return (self).f26

    def fm1(self, p0, p1, p2, globalState):
        return (self).f25

    def m0(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_601_v0_: D3
        d_601_v0_ = D3_DC6((self).f25)
        pat_let_tv18_ = p1
        pat_let_tv19_ = p1
        pat_let_tv20_ = p1
        pat_let_tv21_ = p1
        def lambda28_(source10_):
            if source10_.is_DC6:
                d_602___mcc_h6_ = source10_.cf10
                d_603_cf10_ = d_602___mcc_h6_
                return D4_DC10(_dafny.MultiSet([(False)]), (self).f21)
            elif source10_.is_DC7:
                d_604___mcc_h7_ = source10_.cf11
                d_605___mcc_h8_ = source10_.cf12
                d_606_cf12_ = d_605___mcc_h8_
                d_607_cf11_ = d_604___mcc_h7_
                return D4_DC10(_dafny.MultiSet([d_606_cf12_, (self).f26, pat_let_tv18_]), (self).f21)
            elif source10_.is_DC5:
                d_608___mcc_h9_ = source10_.cf9
                d_609_cf9_ = d_608___mcc_h9_
                return D4_DC10(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, (self).f26, pat_let_tv19_, False])), pat_let_tv20_)
            elif True:
                d_610___mcc_h10_ = source10_.cf13
                d_611_cf13_ = d_610___mcc_h10_
                return D4_DC10(_dafny.MultiSet([not(pat_let_tv21_)]), (self).f26)

        source9_ = lambda28_(d_601_v0_)
        if source9_.is_DC10:
            d_612___mcc_h0_ = source9_.cf15
            d_613___mcc_h1_ = source9_.cf16
            d_614_cf16_ = d_613___mcc_h1_
            d_615_cf15_ = d_612___mcc_h0_
            d_616_v1_: _dafny.Seq
            d_616_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dknejhjps"))
            d_617_v2_: _dafny.Array
            nw96_ = _dafny.Array(None, 16)
            nw96_[int(0)] = default__.safeModuloInt((self).f25, (self).f20)
            nw96_[int(1)] = (self).f25
            nw96_[int(2)] = (self).fm1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbras"))), (self).f20, (self).f26, globalState)
            nw96_[int(3)] = p0
            nw96_[int(4)] = default__.safeDivisionInt((self).f25, (self).f20)
            nw96_[int(5)] = p0
            nw96_[int(6)] = 205
            nw96_[int(7)] = ((self).f20 if (self).f21 else (self).f25)
            nw96_[int(8)] = ((self).f20) - (len(d_616_v1_))
            nw96_[int(9)] = (self).f25
            nw96_[int(10)] = (self).f25
            nw96_[int(11)] = default__.safeModuloInt((self).f20, p0)
            nw96_[int(12)] = (self).fm1(p0, 978, (self).f21, globalState)
            nw96_[int(13)] = (self).f25
            nw96_[int(14)] = (0) - (default__.safeDivisionInt((self).f20, (self).f20))
            nw96_[int(15)] = (self).f20
            d_617_v2_ = nw96_
            index101_ = default__.safeIndex(877, (d_617_v2_).length(0))
            (d_617_v2_)[index101_] = (self).fm1(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_618_i0_ in range(default__.abs(759))])), (self).f20, p1, globalState)
            d_619_v3_: _dafny.Array
            def lambda29_(d_620_p0_):
                def lambda30_(d_621_i1_):
                    return D4_DC11(d_620_p0_, (self).f21, False)

                return lambda30_

            init16_ = lambda29_(p0)
            nw97_ = _dafny.Array(None, 22)
            for i0_16_ in range(nw97_.length(0)):
                nw97_[i0_16_] = init16_(i0_16_)
            d_619_v3_ = nw97_
            d_622_v4_: _dafny.Array
            d_622_v4_ = d_617_v2_
            d_623_v6_: _dafny.Map
            d_623_v6_ = _dafny.Map({not((self).f21): (self).f21})
            d_624_v7_: _dafny.Seq
            d_624_v7_ = _dafny.SeqWithoutIsStrInference([d_623_v6_])
            d_625_v8_: _dafny.Seq
            d_625_v8_ = _dafny.SeqWithoutIsStrInference([(self).f20, p0])
            d_626_v9_: _dafny.Array
            nw98_ = _dafny.Array(None, 25)
            nw98_[int(0)] = (d_622_v4_)
            nw98_[int(1)] = d_617_v2_
            nw98_[int(2)] = d_617_v2_
            nw98_[int(3)] = (d_617_v2_ if (self).f21 else d_617_v2_)
            nw98_[int(4)] = d_617_v2_
            nw98_[int(5)] = d_617_v2_
            nw98_[int(6)] = d_617_v2_
            def iife67_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(82, -571):
                    d_627_v5_: int = compr_23_
                    if ((82) <= (d_627_v5_)) and ((d_627_v5_) < (-571)):
                        coll23_[default__.safeModuloInt(d_627_v5_, 939)] = (self).f20
                return _dafny.Map(coll23_)
            nw98_[int(7)] = (d_617_v2_ if (self).fm0(iife67_()
            , d_616_v1_, (d_624_v7_).set(default__.safeIndex((d_625_v8_)[default__.safeIndex((self).f20, len(d_625_v8_))], len(d_624_v7_)), _dafny.Map({(self).f21: d_614_cf16_})), globalState) else d_617_v2_)
            nw98_[int(8)] = d_617_v2_
            nw98_[int(9)] = d_617_v2_
            nw98_[int(10)] = d_617_v2_
            nw98_[int(11)] = d_617_v2_
            nw98_[int(12)] = d_617_v2_
            nw98_[int(13)] = d_617_v2_
            nw98_[int(14)] = d_617_v2_
            nw98_[int(15)] = d_617_v2_
            nw98_[int(16)] = d_617_v2_
            nw98_[int(17)] = d_617_v2_
            nw98_[int(18)] = d_617_v2_
            nw98_[int(19)] = d_617_v2_
            nw98_[int(20)] = d_617_v2_
            nw98_[int(21)] = d_617_v2_
            nw98_[int(22)] = d_617_v2_
            nw98_[int(23)] = d_617_v2_
            nw98_[int(24)] = d_617_v2_
            d_626_v9_ = nw98_
            index102_ = default__.safeIndex(339, (d_626_v9_).length(0))
            (d_626_v9_)[index102_] = d_617_v2_
            d_628_v10_: _dafny.Array
            def lambda31_(d_629_v6_, d_630_p1_):
                def lambda32_(d_631_i2_):
                    return (d_629_v6_).set((self).f21, d_630_p1_)

                return lambda32_

            init17_ = lambda31_(d_623_v6_, p1)
            nw99_ = _dafny.Array(None, 26)
            for i0_17_ in range(nw99_.length(0)):
                nw99_[i0_17_] = init17_(i0_17_)
            d_628_v10_ = nw99_
            index103_ = default__.safeIndex(561, (d_628_v10_).length(0))
            (d_628_v10_)[index103_] = d_623_v6_
            d_632_v11_: str
            d_632_v11_ = _dafny.CodePoint('b')
            d_633_v12_: _dafny.MultiSet
            d_633_v12_ = _dafny.MultiSet([d_632_v11_])
            index104_ = default__.safeIndex(877, (d_617_v2_).length(0))
            index105_ = default__.safeIndex(339, (d_626_v9_).length(0))
            index106_ = default__.safeIndex(561, (d_628_v10_).length(0))
            rhs86_ = ((0) - ((p0) - ((d_633_v12_).cardinality))) - (-180)
            rhs87_ = d_619_v3_
            rhs88_ = d_617_v2_
            rhs89_ = ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f20: (self).f20}) for d_634_i3_ in range(default__.abs(851))])))) != ((self).f20)
            rhs90_ = ((d_623_v6_).set((self).f26, (self).f21) if (self).f21 else default__.fm15((self).f20, (self).f25, (self).f25, -399, globalState))
            lhs78_ = d_617_v2_
            lhs79_ = default__.safeIndex(877, (d_617_v2_).length(0))
            lhs80_ = d_626_v9_
            lhs81_ = default__.safeIndex(339, (d_626_v9_).length(0))
            lhs82_ = globalState
            lhs83_ = d_628_v10_
            lhs84_ = default__.safeIndex(561, (d_628_v10_).length(0))
            lhs78_[lhs79_] = rhs86_
            d_619_v3_ = rhs87_
            lhs80_[lhs81_] = rhs88_
            lhs82_.f3 = rhs89_
            lhs83_[lhs84_] = rhs90_
            if ((self).f20) == ((self).fm1(len(d_625_v8_), (d_617_v2_)[default__.safeIndex(877, (d_617_v2_).length(0))], False, globalState)):
                (globalState).f15 = _dafny.SeqWithoutIsStrInference([d_632_v11_ for d_635_i4_ in range(default__.abs(348))])
                d_636_v13_: _dafny.MultiSet
                d_636_v13_ = _dafny.MultiSet([(0) - (((self).f20) + ((self).f25))])
                d_636_v13_ = d_636_v13_
                d_637_v14_: _dafny.Array
                nw100_ = _dafny.Array(False, 28)
                d_637_v14_ = nw100_
                index107_ = default__.safeIndex(51, (d_637_v14_).length(0))
                (d_637_v14_)[index107_] = (self).f21
                d_638_v15_: D4
                d_638_v15_ = D4_DC10(d_615_cf15_, p1)
                index108_ = default__.safeIndex(51, (d_637_v14_).length(0))
                (d_637_v14_)[index108_] = (d_638_v15_).cf16
                d_639_v16_: _dafny.Map
                d_639_v16_ = _dafny.Map({(d_617_v2_)[default__.safeIndex(877, (d_617_v2_).length(0))]: (self).f25})
                d_640_v17_: _dafny.Map
                d_640_v17_ = _dafny.Map({d_614_cf16_: len(d_616_v1_)})
                d_641_v18_: _dafny.Map
                d_641_v18_ = _dafny.Map({d_639_v16_: len(d_640_v17_)})
                (globalState).f17 = ((d_641_v18_)[d_639_v16_] if (d_639_v16_) in (d_641_v18_) else -292)
                d_642_v19_: C0
                nw101_ = C0()
                nw101_.ctor__()
                d_642_v19_ = nw101_
            elif True:
                d_643_v20_: _dafny.Set
                d_643_v20_ = _dafny.Set({d_615_cf15_})
                d_644_v21_: _dafny.Seq
                d_644_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_615_cf15_, d_615_cf15_, d_615_cf15_}), _dafny.Set({d_615_cf15_, _dafny.MultiSet([False, p1]), d_615_cf15_})])
                d_643_v20_ = (d_644_v21_)[default__.safeIndex(((self).f20) + ((self).f25), len(d_644_v21_))]
                d_645_v22_: _dafny.Array
                nw102_ = _dafny.Array(None, 6)
                nw102_[int(0)] = d_622_v4_
                nw102_[int(1)] = d_622_v4_
                nw102_[int(2)] = d_622_v4_
                nw102_[int(3)] = d_622_v4_
                nw102_[int(4)] = d_622_v4_
                nw102_[int(5)] = d_622_v4_
                d_645_v22_ = nw102_
                d_646_v23_: _dafny.Map
                d_646_v23_ = _dafny.Map({d_645_v22_: p1})
                d_646_v23_ = (d_646_v23_).set(d_645_v22_, (self).f26)
                index109_ = default__.safeIndex(877, (d_617_v2_).length(0))
                (d_617_v2_)[index109_] = p0
                d_647_v24_: C0
                nw103_ = C0()
                nw103_.ctor__()
                d_647_v24_ = nw103_
                d_648_v25_: _dafny.Map
                d_648_v25_ = _dafny.Map({d_647_v24_: (0) - (((self).f25) * ((d_617_v2_)[default__.safeIndex(877, (d_617_v2_).length(0))]))})
                d_649_v26_: _dafny.Seq
                d_649_v26_ = _dafny.SeqWithoutIsStrInference([d_647_v24_])
                d_648_v25_ = (d_648_v25_).set((d_649_v26_)[default__.safeIndex((d_617_v2_)[default__.safeIndex(877, (d_617_v2_).length(0))], len(d_649_v26_))], (self).f25)
                d_650_v27_: _dafny.Seq
                d_650_v27_ = _dafny.SeqWithoutIsStrInference([(d_622_v4_), (d_626_v9_)[default__.safeIndex(339, (d_626_v9_).length(0))], (d_626_v9_)[default__.safeIndex(339, (d_626_v9_).length(0))]])
                d_651_v28_: T0
                nw104_ = C1()
                nw104_.ctor__(((d_625_v8_)[default__.safeIndex((self).f20, len(d_625_v8_))]) + (p0), (d_650_v27_)[default__.safeIndex(len(default__.fm29((self).f25, (self).f26, globalState)), len(d_650_v27_))], len(d_616_v1_), (self).f21)
                d_651_v28_ = nw104_
                d_652_v29_: _dafny.Map
                d_652_v29_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_628_v10_)[default__.safeIndex(561, (d_628_v10_).length(0))]])): (d_651_v28_).f21})
                d_653_v30_: D9
                d_653_v30_ = D9_DC23((d_617_v2_)[default__.safeIndex(877, (d_617_v2_).length(0))], (self).f26)
                d_654_v31_: _dafny.Array
                def lambda33_(d_655_v1_):
                    def lambda34_(d_656_i5_):
                        return D6_DC13(d_655_v1_)

                    return lambda34_

                init18_ = lambda33_(d_616_v1_)
                nw105_ = _dafny.Array(None, 16)
                for i0_18_ in range(nw105_.length(0)):
                    nw105_[i0_18_] = init18_(i0_18_)
                d_654_v31_ = nw105_
                d_657_v32_: _dafny.Map
                d_657_v32_ = _dafny.Map({(self).f21: d_654_v31_})
                d_658_v33_: _dafny.Array
                nw106_ = _dafny.Array(False, 10)
                d_658_v33_ = nw106_
                d_659_v34_: _dafny.Seq
                d_659_v34_ = _dafny.SeqWithoutIsStrInference([d_658_v33_, d_658_v33_])
                d_660_v35_: D3
                d_660_v35_ = D3_DC5(d_659_v34_)
                d_661_v36_: _dafny.MultiSet
                d_661_v36_ = _dafny.MultiSet([d_660_v35_])
                d_662_v37_: _dafny.Map
                d_662_v37_ = _dafny.Map({(self).f20: ((d_661_v36_)[d_660_v35_] if (d_660_v35_) in (d_661_v36_) else len(_dafny.SeqWithoutIsStrInference([d_632_v11_ for d_663_i6_ in range(default__.abs(454))])))})
                d_664_v38_: _dafny.Seq
                d_664_v38_ = _dafny.SeqWithoutIsStrInference([d_662_v37_])
                d_665_v39_: C1
                nw107_ = C1()
                nw107_.ctor__(p0, d_617_v2_, (d_651_v28_).f20, d_614_cf16_)
                d_665_v39_ = nw107_
                d_666_v40_: _dafny.Seq
                d_666_v40_ = _dafny.SeqWithoutIsStrInference([d_665_v39_])
                d_667_v41_: _dafny.Array
                nw108_ = _dafny.Array(None, 19)
                nw108_[int(0)] = p1
                nw108_[int(1)] = (d_651_v28_).f21
                nw108_[int(2)] = ((d_652_v29_)[(d_651_v28_).f20] if ((d_651_v28_).f20) in (d_652_v29_) else True)
                nw108_[int(3)] = (d_653_v30_).cf32
                nw108_[int(4)] = (self).f21
                nw108_[int(5)] = (self).f26
                nw108_[int(6)] = (d_651_v28_).f21
                nw108_[int(7)] = d_614_cf16_
                nw108_[int(8)] = False
                nw108_[int(9)] = ((d_651_v28_).f21) == ((self).f21)
                nw108_[int(10)] = ((self).f21 if d_614_cf16_ else d_614_cf16_)
                nw108_[int(11)] = not(((self).f21 if d_614_cf16_ else d_614_cf16_))
                nw108_[int(12)] = (d_647_v24_).fm4(len(d_657_v32_), d_664_v38_, (self).f20, globalState)
                nw108_[int(13)] = (self).f21
                nw108_[int(14)] = (d_665_v39_) in (d_666_v40_)
                nw108_[int(15)] = (d_665_v39_).fm0(d_662_v37_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "safjb")), d_624_v7_, globalState)
                nw108_[int(16)] = p1
                nw108_[int(17)] = (self).f21
                nw108_[int(18)] = (self).f26
                d_667_v41_ = nw108_
                index110_ = default__.safeIndex(780, (d_658_v33_).length(0))
                (d_658_v33_)[index110_] = p1
                d_668_v42_: _dafny.MultiSet
                d_668_v42_ = _dafny.MultiSet([286, (self).f25])
                d_669_v44_: _dafny.Set
                d_669_v44_ = _dafny.Set({(self).f26, (self).f21, (self).f26})
                d_670_v45_: _dafny.Seq
                d_670_v45_ = _dafny.SeqWithoutIsStrInference([len(d_669_v44_)])
                d_671_v46_: _dafny.Map
                d_671_v46_ = _dafny.Map({(self).f21: -665})
                index111_ = default__.safeIndex(780, (d_658_v33_).length(0))
                def iife68_():
                    coll24_ = _dafny.Set()
                    compr_24_: int
                    for compr_24_ in (d_668_v42_).Elements:
                        d_672_v43_: int = compr_24_
                        if (d_672_v43_) in (d_668_v42_):
                            coll24_ = coll24_.union(_dafny.Set([default__.safeDivisionInt(d_672_v43_, 145)]))
                    return _dafny.Set(coll24_)
                rhs91_ = (default__.safeDivisionInt(-584, len(iife68_()
                )) if (d_651_v28_).f21 else (self).f25)
                rhs92_ = d_651_v28_
                rhs93_ = not((d_651_v28_).f21)
                rhs94_ = _dafny.SeqWithoutIsStrInference([(self).f20 for d_673_i7_ in range(default__.abs(846))])
                rhs95_ = not (not ((d_665_v39_).fm16(d_615_cf15_, d_670_v45_, globalState)) or (p1)) or (((self).f26) not in (d_671_v46_))
                lhs85_ = globalState
                lhs86_ = globalState
                lhs87_ = d_658_v33_
                lhs88_ = default__.safeIndex(780, (d_658_v33_).length(0))
                lhs85_.f17 = rhs91_
                d_651_v28_ = rhs92_
                lhs86_.f3 = rhs93_
                d_625_v8_ = rhs94_
                lhs87_[lhs88_] = rhs95_
            d_674_v47_: _dafny.MultiSet
            d_674_v47_ = _dafny.MultiSet([512])
            d_675_v48_: _dafny.Map
            d_675_v48_ = _dafny.Map({d_674_v47_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auwfsukpw"))})
            d_675_v48_ = (d_675_v48_).set(d_674_v47_, (d_616_v1_) + (d_616_v1_))
            d_676_v49_: _dafny.Seq
            d_676_v49_ = _dafny.SeqWithoutIsStrInference([d_674_v47_])
            d_676_v49_ = d_676_v49_
        elif source9_.is_DC11:
            d_677___mcc_h2_ = source9_.cf17
            d_678___mcc_h3_ = source9_.cf18
            d_679___mcc_h4_ = source9_.cf19
            d_680_cf19_ = d_679___mcc_h4_
            d_681_cf18_ = d_678___mcc_h3_
            d_682_cf17_ = d_677___mcc_h2_
            if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_683_i8_ in range(default__.abs(372))])) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_684_i9_ in range(default__.abs(332))])):
                d_685_v50_: _dafny.Seq
                d_685_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iagyloo"))
                d_686_v51_: _dafny.Array
                def lambda35_(d_687_i10_):
                    return default__.safeDivisionInt(d_687_i10_, (0) - ((self).f25))

                init19_ = lambda35_
                nw109_ = _dafny.Array(None, 2)
                for i0_19_ in range(nw109_.length(0)):
                    nw109_[i0_19_] = init19_(i0_19_)
                d_686_v51_ = nw109_
                d_688_v52_: C1
                nw110_ = C1()
                nw110_.ctor__((len(d_685_v50_)) + ((self).f20), d_686_v51_, p0, d_681_cf18_)
                d_688_v52_ = nw110_
                d_689_v53_: C0
                nw111_ = C0()
                nw111_.ctor__()
                d_689_v53_ = nw111_
                d_690_v54_: _dafny.Array
                nw112_ = _dafny.Array(None, 14)
                nw112_[int(0)] = d_681_cf18_
                nw112_[int(1)] = (self).f21
                nw112_[int(2)] = (self).f21
                nw112_[int(3)] = (self).f26
                nw112_[int(4)] = d_680_cf19_
                nw112_[int(5)] = False
                nw112_[int(6)] = (self).f21
                nw112_[int(7)] = (self).f26
                nw112_[int(8)] = d_680_cf19_
                nw112_[int(9)] = (self).f26
                nw112_[int(10)] = d_681_cf18_
                nw112_[int(11)] = not(d_681_cf18_)
                nw112_[int(12)] = d_681_cf18_
                nw112_[int(13)] = d_680_cf19_
                d_690_v54_ = nw112_
                d_691_v55_: _dafny.Set
                d_691_v55_ = _dafny.Set({d_690_v54_, d_690_v54_, d_690_v54_})
                r1 = (d_690_v54_) in (d_691_v55_)
                d_692_v56_: str
                d_692_v56_ = _dafny.CodePoint('j')
                d_693_v57_: _dafny.MultiSet
                d_693_v57_ = _dafny.MultiSet([d_692_v56_, d_692_v56_])
                d_694_v58_: _dafny.MultiSet
                d_694_v58_ = _dafny.MultiSet([d_693_v57_])
                d_695_v59_: _dafny.Seq
                d_695_v59_ = _dafny.SeqWithoutIsStrInference([d_694_v58_, _dafny.MultiSet([d_693_v57_, _dafny.MultiSet([d_692_v56_])])])
                d_696_v60_: _dafny.Seq
                d_696_v60_ = _dafny.SeqWithoutIsStrInference([d_694_v58_, d_694_v58_, _dafny.MultiSet([d_693_v57_]), (d_695_v59_)[default__.safeIndex(d_682_cf17_, len(d_695_v59_))], default__.fm31(globalState)])
                d_680_cf19_ = not((_dafny.MultiSet(default__.fm30(globalState))) == ((d_695_v59_)[default__.safeIndex((self).f20, len(d_695_v59_))]))
                (globalState).f4 = (self).f21
            elif True:
                d_697_v61_: _dafny.MultiSet
                d_697_v61_ = _dafny.MultiSet([p0, d_682_cf17_])
                d_698_v62_: _dafny.Map
                d_698_v62_ = _dafny.Map({False: (self).fm1(p0, (d_697_v61_).cardinality, d_680_cf19_, globalState)})
                d_699_v63_: _dafny.Map
                d_699_v63_ = _dafny.Map({((self).f26 if not(True) else (self).f21): (self).fm1(((d_698_v62_)[(self).f26] if ((self).f26) in (d_698_v62_) else p0), p0, (self).f21, globalState)})
                d_699_v63_ = (d_699_v63_).set((self).f26, -402)
                d_700_v64_: _dafny.Seq
                d_700_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvfhbcfbh"))
                d_701_v65_: _dafny.Seq
                d_701_v65_ = _dafny.SeqWithoutIsStrInference([default__.fm26(d_681_cf18_, globalState), d_700_v64_])
                d_701_v65_ = d_701_v65_
                d_702_v66_: _dafny.Map
                d_702_v66_ = _dafny.Map({p1: False})
                d_703_v67_: _dafny.Seq
                d_703_v67_ = _dafny.SeqWithoutIsStrInference([((d_702_v66_)[d_681_cf18_] if (d_681_cf18_) in (d_702_v66_) else d_680_cf19_)])
                d_704_v68_: _dafny.MultiSet
                d_704_v68_ = _dafny.MultiSet([d_703_v67_, d_703_v67_])
                d_705_v69_: _dafny.Seq
                d_705_v69_ = _dafny.SeqWithoutIsStrInference([594, (d_704_v68_).cardinality, (self).fm1(d_682_cf17_, len(d_702_v66_), d_681_cf18_, globalState)])
                d_706_v70_: _dafny.Set
                d_706_v70_ = _dafny.Set({(0) - (d_682_cf17_)})
                d_707_v71_: _dafny.Seq
                d_707_v71_ = _dafny.SeqWithoutIsStrInference([d_706_v70_, d_706_v70_])
                d_708_v72_: _dafny.Seq
                d_708_v72_ = _dafny.SeqWithoutIsStrInference([d_682_cf17_, (self).f25, (D7_DC18(d_705_v69_, d_682_cf17_)).cf27, (self).f25, (D8_DC21(len((d_707_v71_)[default__.safeIndex((self).f20, len(d_707_v71_))]))).cf29])
                d_709_v73_: _dafny.Array
                nw113_ = _dafny.Array(int(0), 23)
                d_709_v73_ = nw113_
                d_710_v74_: _dafny.Seq
                d_710_v74_ = _dafny.SeqWithoutIsStrInference([d_709_v73_])
                rhs96_ = d_681_cf18_
                rhs97_ = ((0) - (d_682_cf17_) if not(True) else (0) - (p0))
                rhs98_ = (d_708_v72_) < (d_708_v72_)
                rhs99_ = d_710_v74_
                lhs89_ = globalState
                lhs90_ = globalState
                lhs91_ = globalState
                lhs89_.f18 = rhs96_
                lhs90_.f5 = rhs97_
                lhs91_.f14 = rhs98_
                r0 = rhs99_
                d_711_v75_: _dafny.Map
                d_711_v75_ = _dafny.Map({p0: len(d_705_v69_)})
                d_712_v76_: _dafny.Map
                d_712_v76_ = _dafny.Map({d_711_v75_: (self).f20})
                d_712_v76_ = (d_712_v76_).set(d_711_v75_, p0)
                d_713_v77_: str
                d_713_v77_ = _dafny.CodePoint('v')
                d_714_v78_: C1
                nw114_ = C1()
                nw114_.ctor__(d_682_cf17_, d_709_v73_, (self).f25, d_681_cf18_)
                d_714_v78_ = nw114_
                d_715_v79_: _dafny.Map
                d_715_v79_ = _dafny.Map({d_713_v77_: d_714_v78_})
                d_716_v80_: _dafny.Seq
                d_716_v80_ = _dafny.SeqWithoutIsStrInference([d_715_v79_, d_715_v79_])
                d_717_v81_: D10
                d_717_v81_ = D10_DC27(d_682_cf17_, d_716_v80_, d_713_v77_)
                d_718_v82_: _dafny.MultiSet
                d_718_v82_ = _dafny.MultiSet([(d_717_v81_).cf38])
                d_719_v83_: _dafny.Seq
                d_719_v83_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_700_v64_), d_718_v82_])
                d_720_v84_: _dafny.Array
                nw115_ = _dafny.Array(None, 17)
                nw115_[int(0)] = (d_718_v82_) | (default__.fm32(globalState))
                nw115_[int(1)] = d_718_v82_
                nw115_[int(2)] = d_718_v82_
                nw115_[int(3)] = (d_719_v83_)[default__.safeIndex((self).f20, len(d_719_v83_))]
                nw115_[int(4)] = d_718_v82_
                nw115_[int(5)] = d_718_v82_
                nw115_[int(6)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), d_713_v77_]))
                nw115_[int(7)] = (_dafny.MultiSet([d_713_v77_, _dafny.CodePoint('v')]) if d_680_cf19_ else d_718_v82_)
                nw115_[int(8)] = (d_718_v82_) | ((d_719_v83_)[default__.safeIndex(d_714_v78_.f27, len(d_719_v83_))])
                nw115_[int(9)] = _dafny.MultiSet([d_713_v77_, d_713_v77_, d_713_v77_, d_713_v77_, d_713_v77_])
                nw115_[int(10)] = d_718_v82_
                nw115_[int(11)] = d_718_v82_
                nw115_[int(12)] = _dafny.MultiSet([d_713_v77_, d_713_v77_])
                nw115_[int(13)] = _dafny.MultiSet([d_713_v77_])
                nw115_[int(14)] = d_718_v82_
                nw115_[int(15)] = (d_718_v82_) - (d_718_v82_)
                nw115_[int(16)] = d_718_v82_
                d_720_v84_ = nw115_
                d_720_v84_ = d_720_v84_
            (globalState).f17 = (self).f25
            d_721_v85_: bool
            d_721_v85_ = d_680_cf19_
            if (d_721_v85_):
                d_722_v86_: _dafny.Array
                nw116_ = _dafny.Array(int(0), 20)
                d_722_v86_ = nw116_
                index112_ = default__.safeIndex(969, (d_722_v86_).length(0))
                (d_722_v86_)[index112_] = (self).f25
                d_723_v87_: D7
                d_723_v87_ = D7_DC17(d_681_cf18_)
                d_724_v88_: _dafny.Set
                d_724_v88_ = _dafny.Set({False})
                d_725_v89_: _dafny.Seq
                d_725_v89_ = _dafny.SeqWithoutIsStrInference([d_724_v88_, d_724_v88_])
                d_726_v90_: _dafny.Map
                d_726_v90_ = _dafny.Map({d_723_v87_: d_725_v89_})
                d_727_v91_: D3
                d_727_v91_ = D3_DC7(((d_726_v90_)[d_723_v87_] if (d_723_v87_) in (d_726_v90_) else d_725_v89_), d_680_cf19_)
                d_728_v92_: _dafny.MultiSet
                d_728_v92_ = _dafny.MultiSet([(self).f20, (self).fm1(-252, (0) - (d_682_cf17_), (d_727_v91_).cf12, globalState), (self).f25])
                index113_ = default__.safeIndex(969, (d_722_v86_).length(0))
                (d_722_v86_)[index113_] = ((d_728_v92_)[(self).f20] if ((self).f20) in (d_728_v92_) else (self).f25)
                d_729_v93_: _dafny.Array
                def lambda36_(d_730_p1_):
                    def lambda37_(d_731_i11_):
                        return d_730_p1_

                    return lambda37_

                init20_ = lambda36_(p1)
                nw117_ = _dafny.Array(None, 12)
                for i0_20_ in range(nw117_.length(0)):
                    nw117_[i0_20_] = init20_(i0_20_)
                d_729_v93_ = nw117_
                nw118_ = _dafny.Array(None, 9)
                nw118_[int(0)] = p1
                nw118_[int(1)] = not (d_681_cf18_) or (p1)
                nw118_[int(2)] = True
                nw118_[int(3)] = p1
                nw118_[int(4)] = (self).f21
                nw118_[int(5)] = d_681_cf18_
                nw118_[int(6)] = p1
                nw118_[int(7)] = not((not(p1)) and (d_681_cf18_))
                nw118_[int(8)] = True
                d_729_v93_ = nw118_
                d_732_v94_: _dafny.Array
                nw119_ = _dafny.Array(_dafny.Seq({}), 26)
                d_732_v94_ = nw119_
                d_733_v95_: _dafny.Seq
                d_733_v95_ = _dafny.SeqWithoutIsStrInference([d_722_v86_, d_722_v86_])
                d_734_v96_: _dafny.Seq
                d_734_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysgax"))
                rhs100_ = d_732_v94_
                rhs101_ = 785
                rhs102_ = len(_dafny.SeqWithoutIsStrInference([(d_733_v95_)[default__.safeIndex(p0, len(d_733_v95_))], d_722_v86_]))
                rhs103_ = d_680_cf19_
                rhs104_ = default__.safeModuloInt(p0, (p0) - (len(d_734_v96_)))
                lhs92_ = globalState
                lhs93_ = globalState
                lhs94_ = globalState
                d_732_v94_ = rhs100_
                lhs92_.f17 = rhs101_
                d_682_cf17_ = rhs102_
                lhs93_.f14 = rhs103_
                lhs94_.f5 = rhs104_
                d_735_v97_: _dafny.Seq
                d_735_v97_ = _dafny.SeqWithoutIsStrInference([(self).f21, d_681_cf18_])
                d_736_v98_: _dafny.Map
                d_736_v98_ = _dafny.Map({(0) - (p0): d_735_v97_})
                (globalState).f18 = (len(d_736_v98_)) == ((0) - (len((d_734_v96_).set(default__.safeIndex(-800, len(d_734_v96_)), _dafny.CodePoint('f')))))
                (globalState).f4 = (d_681_cf18_) or (False)
            elif True:
                d_737_v99_: _dafny.Array
                nw120_ = _dafny.Array(D6.default()(), 25)
                d_737_v99_ = nw120_
                d_738_v100_: _dafny.Seq
                d_738_v100_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f25, (self).f20, 81, p0])])
                d_739_v101_: _dafny.Seq
                d_739_v101_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                rhs105_ = (d_737_v99_ if p1 else d_737_v99_)
                rhs106_ = d_738_v100_
                rhs107_ = (d_739_v101_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvtiyhvfl"))), len(d_739_v101_))]
                rhs108_ = (self).f20
                rhs109_ = d_682_cf17_
                lhs95_ = globalState
                d_737_v99_ = rhs105_
                d_738_v100_ = rhs106_
                lhs95_.f14 = rhs107_
                d_682_cf17_ = rhs108_
                d_682_cf17_ = rhs109_
                d_740_v102_: _dafny.Array
                def lambda38_(d_741_i12_):
                    return (_dafny.Set({(self).f21})).ispropersubset(_dafny.Set({(self).f26}))

                init21_ = lambda38_
                nw121_ = _dafny.Array(None, 27)
                for i0_21_ in range(nw121_.length(0)):
                    nw121_[i0_21_] = init21_(i0_21_)
                d_740_v102_ = nw121_
                index114_ = default__.safeIndex(278, (d_740_v102_).length(0))
                (d_740_v102_)[index114_] = (d_681_cf18_) == (d_680_cf19_)
                index115_ = default__.safeIndex(278, (d_740_v102_).length(0))
                (d_740_v102_)[index115_] = (self).f26
                d_742_v103_: _dafny.Seq
                d_742_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etpbgjoh"))
                d_743_v104_: D6
                d_743_v104_ = D6_DC13(d_742_v103_)
                (globalState).f15 = (d_743_v104_).cf21
                d_744_v105_: _dafny.Map
                d_744_v105_ = _dafny.Map({p0: d_742_v103_})
                d_745_v106_: _dafny.Seq
                d_745_v106_ = _dafny.SeqWithoutIsStrInference([d_742_v103_, d_742_v103_, d_742_v103_, d_742_v103_])
                d_744_v105_ = (d_744_v105_).set(p0, (d_745_v106_)[default__.safeIndex(p0, len(d_745_v106_))])
                d_746_v107_: D9
                d_746_v107_ = D9_DC23(75, p1)
                d_747_v108_: D9
                d_747_v108_ = D9_DC24(D9_DC24(d_746_v107_))
                d_748_v109_: _dafny.Array
                nw122_ = _dafny.Array(None, 11)
                nw122_[int(0)] = D9_DC24(d_746_v107_)
                nw122_[int(1)] = d_747_v108_
                nw122_[int(2)] = d_747_v108_
                nw122_[int(3)] = d_747_v108_
                nw122_[int(4)] = d_747_v108_
                nw122_[int(5)] = d_747_v108_
                nw122_[int(6)] = D9_DC24(D9_DC23(d_682_cf17_, d_681_cf18_))
                nw122_[int(7)] = D9_DC24(D9_DC24(D9_DC23(722, d_681_cf18_)))
                nw122_[int(8)] = d_747_v108_
                nw122_[int(9)] = d_747_v108_
                nw122_[int(10)] = d_747_v108_
                d_748_v109_ = nw122_
                index116_ = default__.safeIndex(490, (d_748_v109_).length(0))
                (d_748_v109_)[index116_] = d_747_v108_
                d_749_v110_: _dafny.Map
                d_749_v110_ = _dafny.Map({(self).f26: (d_740_v102_)[default__.safeIndex(278, (d_740_v102_).length(0))]})
                d_750_v111_: _dafny.Seq
                d_750_v111_ = _dafny.SeqWithoutIsStrInference([d_749_v110_])
                index117_ = default__.safeIndex(490, (d_748_v109_).length(0))
                (d_748_v109_)[index117_] = ((d_747_v108_ if not(p1) else d_747_v108_) if (self).fm0(_dafny.Map({(self).f20: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmql")))}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbvsweaf")), d_750_v111_, globalState) else d_747_v108_)
            d_751_v112_: _dafny.Array
            nw123_ = _dafny.Array(_dafny.Seq({}), 21)
            d_751_v112_ = nw123_
            d_752_v113_: _dafny.Seq
            d_752_v113_ = _dafny.SeqWithoutIsStrInference([p0, (self).f20])
            index118_ = default__.safeIndex(19, (d_751_v112_).length(0))
            (d_751_v112_)[index118_] = d_752_v113_
            index119_ = default__.safeIndex(19, (d_751_v112_).length(0))
            (d_751_v112_)[index119_] = d_752_v113_
        elif True:
            d_753___mcc_h5_ = source9_.cf14
            d_754_cf14_ = d_753___mcc_h5_
            d_755_v114_: _dafny.Seq
            d_755_v114_ = _dafny.SeqWithoutIsStrInference([(self).f25, p0])
            d_756_v115_: _dafny.Set
            d_756_v115_ = _dafny.Set({(d_755_v114_ if p1 else d_755_v114_)})
            d_756_v115_ = d_756_v115_
            d_757_v116_: C2
            nw124_ = C2()
            nw124_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngcig"))), (self).f26)
            d_757_v116_ = nw124_
            d_758_v117_: str
            d_758_v117_ = _dafny.CodePoint('a')
            d_758_v117_ = d_758_v117_
            d_759_v118_: _dafny.Set
            d_759_v118_ = _dafny.Set({p1})
            d_760_v119_: _dafny.Seq
            d_760_v119_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1, (self).f26})])
            d_761_v120_: D3
            d_761_v120_ = D3_DC7(((_dafny.SeqWithoutIsStrInference([d_759_v118_, d_759_v118_])).set(default__.safeIndex((self).f20, len(_dafny.SeqWithoutIsStrInference([d_759_v118_, d_759_v118_]))), d_759_v118_)) + (d_760_v119_), p1)
            source11_ = d_761_v120_
            if source11_.is_DC6:
                d_762___mcc_h11_ = source11_.cf10
                d_763_cf10_ = d_762___mcc_h11_
                d_764_v121_: _dafny.Seq
                d_764_v121_ = _dafny.SeqWithoutIsStrInference([p1])
                d_765_v122_: _dafny.Array
                def lambda39_(d_766_i13_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bty"))

                init22_ = lambda39_
                nw125_ = _dafny.Array(None, 15)
                for i0_22_ in range(nw125_.length(0)):
                    nw125_[i0_22_] = init22_(i0_22_)
                d_765_v122_ = nw125_
                d_767_v123_: _dafny.Map
                d_767_v123_ = _dafny.Map({(self).f25: p0})
                d_768_v124_: _dafny.Map
                d_768_v124_ = _dafny.Map({(self).f21: True})
                d_769_v125_: _dafny.Seq
                d_769_v125_ = _dafny.SeqWithoutIsStrInference([d_768_v124_])
                d_770_v126_: _dafny.Seq
                d_770_v126_ = _dafny.SeqWithoutIsStrInference([d_769_v125_])
                rhs110_ = (332) - ((p0 if (self).fm0(d_767_v123_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivwxw")), (d_770_v126_)[default__.safeIndex((self).f20, len(d_770_v126_))], globalState) else -46))
                rhs111_ = d_764_v121_
                rhs112_ = d_765_v122_
                rhs113_ = (self).f26
                lhs96_ = globalState
                lhs97_ = globalState
                lhs96_.f17 = rhs110_
                d_764_v121_ = rhs111_
                d_765_v122_ = rhs112_
                lhs97_.f4 = rhs113_
                d_771_v128_: _dafny.Map
                d_771_v128_ = _dafny.Map({p0: p1})
                d_772_v129_: _dafny.Map
                d_772_v129_ = _dafny.Map({len(d_755_v114_): ((d_771_v128_)[len(_dafny.Map({(self).f26: (self).f20}))] if (len(_dafny.Map({(self).f26: (self).f20}))) in (d_771_v128_) else p1)})
                d_773_v131_: _dafny.Seq
                d_773_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdacsl"))
                d_774_v132_: _dafny.Map
                def iife69_():
                    coll25_ = _dafny.Set()
                    compr_25_: int
                    for compr_25_ in (d_772_v129_).keys.Elements:
                        d_775_v130_: int = compr_25_
                        if (d_775_v130_) in (d_772_v129_):
                            coll25_ = coll25_.union(_dafny.Set([default__.safeModuloInt(d_775_v130_, 486)]))
                    return _dafny.Set(coll25_)
                d_774_v132_ = _dafny.Map({iife69_()
                : d_773_v131_})
                d_776_v133_: _dafny.Array
                nw126_ = _dafny.Array(None, 14)
                nw126_[int(0)] = default__.safeModuloInt((self).f20, p0)
                nw126_[int(1)] = len((d_768_v124_) | (d_768_v124_))
                nw126_[int(2)] = len(d_759_v118_)
                nw126_[int(3)] = (self).f20
                nw126_[int(4)] = (self).f25
                nw126_[int(5)] = d_763_cf10_
                nw126_[int(6)] = (self).f25
                nw126_[int(7)] = p0
                nw126_[int(8)] = len(d_767_v123_)
                nw126_[int(9)] = (self).f20
                def iife70_():
                    coll26_ = _dafny.Map()
                    compr_26_: _dafny.Set
                    for compr_26_ in (d_774_v132_).keys.Elements:
                        d_777_v127_: _dafny.Set = compr_26_
                        if (d_777_v127_) in (d_774_v132_):
                            coll26_[d_777_v127_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))
                    return _dafny.Map(coll26_)
                nw126_[int(10)] = (len(iife70_()
                )) + ((self).f25)
                nw126_[int(11)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pttigqy")))
                nw126_[int(12)] = p0
                nw126_[int(13)] = p0
                d_776_v133_ = nw126_
                index120_ = default__.safeIndex(417, (d_776_v133_).length(0))
                (d_776_v133_)[index120_] = d_763_cf10_
                index121_ = default__.safeIndex(417, (d_776_v133_).length(0))
                rhs114_ = not((d_759_v118_).isdisjoint(d_759_v118_))
                rhs115_ = (d_757_v116_).fm1(d_763_cf10_, (self).f20, (self).f21, globalState)
                lhs98_ = globalState
                lhs99_ = d_776_v133_
                lhs100_ = default__.safeIndex(417, (d_776_v133_).length(0))
                lhs98_.f18 = rhs114_
                lhs99_[lhs100_] = rhs115_
                rhs116_ = (_dafny.Map({p1: (self).f21})).set(p1, not((self).f26))
                rhs117_ = (0) - ((d_776_v133_)[default__.safeIndex(417, (d_776_v133_).length(0))])
                rhs118_ = d_773_v131_
                rhs119_ = p1
                rhs120_ = (_dafny.Map({False: True})) != (default__.fm15((d_776_v133_)[default__.safeIndex(417, (d_776_v133_).length(0))], (d_776_v133_)[default__.safeIndex(417, (d_776_v133_).length(0))], 953, len(d_773_v131_), globalState))
                lhs101_ = globalState
                lhs102_ = globalState
                d_768_v124_ = rhs116_
                d_763_cf10_ = rhs117_
                d_773_v131_ = rhs118_
                lhs101_.f3 = rhs119_
                lhs102_.f3 = rhs120_
                index122_ = default__.safeIndex(417, (d_776_v133_).length(0))
                (d_776_v133_)[index122_] = (0) - (p0)
            elif source11_.is_DC7:
                d_778___mcc_h12_ = source11_.cf11
                d_779___mcc_h13_ = source11_.cf12
                d_780_cf12_ = d_779___mcc_h13_
                d_781_cf11_ = d_778___mcc_h12_
                d_782_v134_: _dafny.Array
                nw127_ = _dafny.Array(int(0), 17)
                d_782_v134_ = nw127_
                index123_ = default__.safeIndex(870, (d_782_v134_).length(0))
                (d_782_v134_)[index123_] = (self).f20
                d_783_v135_: _dafny.Array
                def lambda40_(d_784_p0_, d_785_cf12_):
                    def lambda41_(d_786_i14_):
                        return D2_DC3(d_784_p0_, (self).f26, d_785_cf12_, 593, (self).f26)

                    return lambda41_

                init23_ = lambda40_(p0, d_780_cf12_)
                nw128_ = _dafny.Array(None, 19)
                for i0_23_ in range(nw128_.length(0)):
                    nw128_[i0_23_] = init23_(i0_23_)
                d_783_v135_ = nw128_
                d_787_v136_: D2
                d_787_v136_ = D2_DC3((self).f25, p1, False, (self).f25, (self).f21)
                index124_ = default__.safeIndex(835, (d_783_v135_).length(0))
                (d_783_v135_)[index124_] = d_787_v136_
                d_788_v137_: _dafny.Map
                d_788_v137_ = _dafny.Map({p1: d_757_v116_})
                index125_ = default__.safeIndex(870, (d_782_v134_).length(0))
                index126_ = default__.safeIndex(835, (d_783_v135_).length(0))
                rhs121_ = len(d_788_v137_)
                rhs122_ = d_787_v136_
                rhs123_ = p1
                rhs124_ = (d_755_v114_)[default__.safeIndex(29, len(d_755_v114_))]
                lhs103_ = d_782_v134_
                lhs104_ = default__.safeIndex(870, (d_782_v134_).length(0))
                lhs105_ = d_783_v135_
                lhs106_ = default__.safeIndex(835, (d_783_v135_).length(0))
                lhs107_ = globalState
                lhs108_ = globalState
                lhs103_[lhs104_] = rhs121_
                lhs105_[lhs106_] = rhs122_
                lhs107_.f3 = rhs123_
                lhs108_.f17 = rhs124_
                d_789_v138_: _dafny.Seq
                d_789_v138_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgfualrx"))
                (globalState).f5 = (len((d_789_v138_) + (d_789_v138_))) + (len(((_dafny.SeqWithoutIsStrInference([756]) if d_780_cf12_ else d_755_v114_)).set(default__.safeIndex((d_755_v114_)[default__.safeIndex((self).f20, len(d_755_v114_))], len((_dafny.SeqWithoutIsStrInference([756]) if d_780_cf12_ else d_755_v114_))), (d_782_v134_)[default__.safeIndex(870, (d_782_v134_).length(0))])))
                d_790_v139_: _dafny.Map
                d_790_v139_ = _dafny.Map({(self).f25: (self).f26})
                d_791_v140_: _dafny.Map
                d_791_v140_ = _dafny.Map({d_789_v138_: d_790_v139_})
                d_792_v141_: _dafny.Seq
                d_792_v141_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_793_v142_: _dafny.Map
                d_793_v142_ = _dafny.Map({d_791_v140_: (d_792_v141_) <= (d_792_v141_)})
                d_793_v142_ = (d_793_v142_).set(_dafny.Map({d_789_v138_: d_790_v139_}), (self).f26)
                index127_ = default__.safeIndex(485, (d_754_cf14_).length(0))
                (d_754_cf14_)[index127_] = d_780_cf12_
                d_794_v143_: _dafny.Set
                d_794_v143_ = _dafny.Set({d_758_v117_, d_758_v117_})
                index128_ = default__.safeIndex(485, (d_754_cf14_).length(0))
                index129_ = default__.safeIndex(870, (d_782_v134_).length(0))
                rhs125_ = d_780_cf12_
                rhs126_ = ((d_782_v134_)[default__.safeIndex(870, (d_782_v134_).length(0))]) <= ((d_782_v134_)[default__.safeIndex(870, (d_782_v134_).length(0))])
                rhs127_ = (0) - ((0) - (((d_782_v134_)[default__.safeIndex(870, (d_782_v134_).length(0))]) * (len(((d_789_v138_) + (d_789_v138_)).set(default__.safeIndex(len(d_794_v143_), len((d_789_v138_) + (d_789_v138_))), d_758_v117_)))))
                lhs109_ = globalState
                lhs110_ = d_754_cf14_
                lhs111_ = default__.safeIndex(485, (d_754_cf14_).length(0))
                lhs112_ = d_782_v134_
                lhs113_ = default__.safeIndex(870, (d_782_v134_).length(0))
                lhs109_.f14 = rhs125_
                lhs110_[lhs111_] = rhs126_
                lhs112_[lhs113_] = rhs127_
            elif source11_.is_DC5:
                d_795___mcc_h14_ = source11_.cf9
                d_796_cf9_ = d_795___mcc_h14_
                d_797_v144_: _dafny.Map
                d_797_v144_ = _dafny.Map({p1: d_755_v114_})
                d_797_v144_ = (d_797_v144_).set(not(p1), _dafny.SeqWithoutIsStrInference([p0, p0, (self).f25]))
                (globalState).f5 = (0) - ((self).f25)
                d_798_v145_: _dafny.Seq
                d_798_v145_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pd"))
                d_799_v146_: _dafny.Array
                nw129_ = _dafny.Array(None, 2)
                nw129_[int(0)] = (self).f20
                nw129_[int(1)] = (self).f20
                d_799_v146_ = nw129_
                d_800_v147_: _dafny.Map
                d_800_v147_ = _dafny.Map({d_798_v145_: d_799_v146_})
                d_801_v148_: C1
                nw130_ = C1()
                nw130_.ctor__(p0, ((d_800_v147_)[d_798_v145_] if (d_798_v145_) in (d_800_v147_) else d_799_v146_), default__.safeDivisionInt((self).f25, 656), p1)
                d_801_v148_ = nw130_
                d_802_v149_: C1
                nw131_ = C1()
                nw131_.ctor__(((0) - ((self).f25)) * (p0), (d_801_v148_).f28, ((self).f25) + (316), (self).f26)
                d_802_v149_ = nw131_
            elif True:
                d_803___mcc_h15_ = source11_.cf13
                d_804_cf13_ = d_803___mcc_h15_
                rhs128_ = p0
                rhs129_ = True
                rhs130_ = (d_757_v116_).fm1(p0, (self).f25, (self).f26, globalState)
                lhs114_ = globalState
                lhs115_ = globalState
                lhs116_ = globalState
                lhs114_.f17 = rhs128_
                lhs115_.f18 = rhs129_
                lhs116_.f17 = rhs130_
                d_805_v151_: _dafny.Array
                def lambda42_(d_806_p0_):
                    def lambda43_(d_807_i15_):
                        def iife71_():
                            coll27_ = _dafny.Set()
                            compr_27_: int
                            for compr_27_ in _dafny.IntegerRange(439, 582):
                                d_808_v150_: int = compr_27_
                                if ((439) <= (d_808_v150_)) and ((d_808_v150_) < (582)):
                                    coll27_ = coll27_.union(_dafny.Set([(d_808_v150_) + (d_806_p0_)]))
                            return _dafny.Set(coll27_)
                        return iife71_()
                        

                    return lambda43_

                init24_ = lambda42_(p0)
                nw132_ = _dafny.Array(None, 3)
                for i0_24_ in range(nw132_.length(0)):
                    nw132_[i0_24_] = init24_(i0_24_)
                d_805_v151_ = nw132_
                d_809_v152_: D8
                d_809_v152_ = D8_DC19(_dafny.Map({(self).f20: (self).f20}))
                d_810_v153_: _dafny.Seq
                d_810_v153_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvwgksww"))
                d_811_v154_: _dafny.Map
                d_811_v154_ = _dafny.Map({not(p1): True})
                d_812_v155_: _dafny.Map
                d_812_v155_ = _dafny.Map({(_dafny.Map({p1: p1})).set((d_757_v116_).fm0((d_809_v152_).cf28, (d_810_v153_).set(default__.safeIndex((d_757_v116_).fm1((self).f20, p0, p1, globalState), len(d_810_v153_)), d_758_v117_), _dafny.SeqWithoutIsStrInference([d_811_v154_, d_811_v154_, d_811_v154_]), globalState), (self).f21): _dafny.Map({len(d_755_v114_): (self).f26})})
                rhs131_ = d_805_v151_
                rhs132_ = (d_812_v155_) | (d_812_v155_)
                rhs133_ = (self).f21
                lhs117_ = globalState
                d_805_v151_ = rhs131_
                d_812_v155_ = rhs132_
                lhs117_.f3 = rhs133_
                d_813_v156_: _dafny.Map
                d_813_v156_ = _dafny.Map({d_754_cf14_: (0) - ((self).f25)})
                d_814_v157_: _dafny.Seq
                d_814_v157_ = _dafny.SeqWithoutIsStrInference([d_813_v156_, d_813_v156_, (d_813_v156_) | (d_813_v156_), (d_813_v156_) | (d_813_v156_)])
                d_813_v156_ = (d_814_v157_)[default__.safeIndex(-318, len(d_814_v157_))]
                d_815_v158_: D4
                d_815_v158_ = D4_DC11(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqebepvf"))), p1, p1)
                d_816_v159_: _dafny.Map
                d_816_v159_ = _dafny.Map({d_759_v118_: p0})
                d_817_v160_: _dafny.Array
                nw133_ = _dafny.Array(_dafny.Seq({}), 28)
                d_817_v160_ = nw133_
                d_818_v161_: _dafny.MultiSet
                d_818_v161_ = _dafny.MultiSet([d_817_v160_])
                d_819_v162_: _dafny.Array
                nw134_ = _dafny.Array(None, 12)
                nw134_[int(0)] = (d_815_v158_).cf17
                nw134_[int(1)] = (self).fm1((self).f20, (self).f25, p1, globalState)
                nw134_[int(2)] = p0
                nw134_[int(3)] = (((d_816_v159_)[d_759_v118_] if (d_759_v118_) in (d_816_v159_) else ((d_818_v161_)[d_817_v160_] if (d_817_v160_) in (d_818_v161_) else (self).f25))) * ((self).f25)
                nw134_[int(4)] = 489
                nw134_[int(5)] = p0
                nw134_[int(6)] = p0
                nw134_[int(7)] = default__.safeModuloInt((self).f20, len(d_811_v154_))
                nw134_[int(8)] = -12
                nw134_[int(9)] = len(_dafny.SeqWithoutIsStrInference([(self).f20, (self).f25]))
                nw134_[int(10)] = ((self).f25) + ((0) - (p0))
                nw134_[int(11)] = p0
                d_819_v162_ = nw134_
                index130_ = default__.safeIndex(305, (d_819_v162_).length(0))
                (d_819_v162_)[index130_] = p0
                index131_ = default__.safeIndex(305, (d_819_v162_).length(0))
                (d_819_v162_)[index131_] = (self).f20
        d_820_v163_: _dafny.Seq
        d_820_v163_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uftkj"))
        d_821_v164_: _dafny.Seq
        d_821_v164_ = _dafny.SeqWithoutIsStrInference([d_820_v163_, (d_820_v163_ if (self).f21 else d_820_v163_)])
        d_822_v165_: _dafny.MultiSet
        d_822_v165_ = _dafny.MultiSet([(self).f25])
        rhs134_ = default__.safeModuloInt(default__.safeDivisionInt((d_822_v165_).cardinality, (self).f25), (self).f20)
        rhs135_ = (0) - ((((self).f20) * ((self).f25)) * ((self).f20))
        rhs136_ = d_821_v164_
        rhs137_ = (self).f26
        rhs138_ = p0
        lhs118_ = globalState
        lhs119_ = globalState
        lhs120_ = globalState
        lhs121_ = globalState
        lhs118_.f5 = rhs134_
        lhs119_.f17 = rhs135_
        d_821_v164_ = rhs136_
        lhs120_.f14 = rhs137_
        lhs121_.f5 = rhs138_
        d_823_v166_: str
        d_823_v166_ = _dafny.CodePoint('w')
        d_824_v167_: _dafny.MultiSet
        d_824_v167_ = _dafny.MultiSet([p1])
        d_825_v168_: _dafny.Seq
        d_825_v168_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21])
        d_826_v169_: _dafny.Set
        d_826_v169_ = _dafny.Set({(self).f26})
        d_827_v170_: _dafny.Seq
        d_827_v170_ = _dafny.SeqWithoutIsStrInference([(self).f20, (self).f20])
        d_828_v171_: D6
        d_828_v171_ = D6_DC13(d_820_v163_)
        d_829_v172_: _dafny.Map
        d_829_v172_ = _dafny.Map({-993: len((d_828_v171_).cf21)})
        d_830_v174_: D8
        def iife72_():
            coll28_ = _dafny.Map()
            compr_28_: int
            for compr_28_ in _dafny.IntegerRange(377, 479):
                d_831_v173_: int = compr_28_
                if ((377) <= (d_831_v173_)) and ((d_831_v173_) < (479)):
                    coll28_[default__.safeDivisionInt(d_831_v173_, (self).f20)] = len(_dafny.Map({(self).f25: (self).f20}))
            return _dafny.Map(coll28_)
        d_830_v174_ = D8_DC19(iife72_()
)
        d_832_v175_: _dafny.Array
        nw135_ = _dafny.Array(None, 27)
        nw135_[int(0)] = (self).f25
        nw135_[int(1)] = (self).f25
        nw135_[int(2)] = (self).f25
        nw135_[int(3)] = (self).f20
        nw135_[int(4)] = ((d_824_v167_)[(self).f21] if ((self).f21) in (d_824_v167_) else p0)
        nw135_[int(5)] = p0
        nw135_[int(6)] = (self).f20
        nw135_[int(7)] = 489
        nw135_[int(8)] = len(default__.fm18((self).f20, (self).f20, d_823_v166_, globalState))
        nw135_[int(9)] = 248
        nw135_[int(10)] = (self).f25
        nw135_[int(11)] = len(d_825_v168_)
        nw135_[int(12)] = len(d_826_v169_)
        nw135_[int(13)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esaxrsb")))
        nw135_[int(14)] = 268
        nw135_[int(15)] = p0
        nw135_[int(16)] = (self).f25
        nw135_[int(17)] = (self).f20
        nw135_[int(18)] = (self).f25
        nw135_[int(19)] = p0
        nw135_[int(20)] = len(d_827_v170_)
        nw135_[int(21)] = (self).f20
        nw135_[int(22)] = 48
        nw135_[int(23)] = (self).f25
        nw135_[int(24)] = len(d_829_v172_)
        nw135_[int(25)] = len(d_825_v168_)
        nw135_[int(26)] = len((d_830_v174_).cf28)
        d_832_v175_ = nw135_
        d_833_v176_: C1
        nw136_ = C1()
        nw136_.ctor__((self).f20, d_832_v175_, (self).f20, (self).f21)
        d_833_v176_ = nw136_
        d_834_v177_: _dafny.Map
        d_834_v177_ = _dafny.Map({d_823_v166_: d_833_v176_})
        d_835_v178_: _dafny.Seq
        d_835_v178_ = _dafny.SeqWithoutIsStrInference([d_834_v177_])
        d_836_v179_: D10
        d_836_v179_ = D10_DC27(len(d_820_v163_), d_835_v178_, _dafny.CodePoint('p'))
        d_837_v180_: _dafny.Array
        nw137_ = _dafny.Array(None, 16)
        nw137_[int(0)] = d_823_v166_
        nw137_[int(1)] = d_823_v166_
        nw137_[int(2)] = d_823_v166_
        nw137_[int(3)] = d_823_v166_
        nw137_[int(4)] = d_823_v166_
        nw137_[int(5)] = d_823_v166_
        nw137_[int(6)] = d_823_v166_
        nw137_[int(7)] = (d_836_v179_).cf38
        nw137_[int(8)] = d_823_v166_
        nw137_[int(9)] = d_823_v166_
        nw137_[int(10)] = d_823_v166_
        nw137_[int(11)] = d_823_v166_
        nw137_[int(12)] = d_823_v166_
        nw137_[int(13)] = d_823_v166_
        nw137_[int(14)] = d_823_v166_
        nw137_[int(15)] = _dafny.CodePoint('e')
        d_837_v180_ = nw137_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_837_v180_).length(0)):
            d_838_i16_: int = guard_loop_1_
            if (True) and (((0) <= (d_838_i16_)) and ((d_838_i16_) < ((d_837_v180_).length(0)))):
                (d_837_v180_)[(d_838_i16_)] = (d_820_v163_)[default__.safeIndex(len(_dafny.Map({d_827_v170_: d_833_v176_.f27})), len(d_820_v163_))]
        d_821_v164_ = default__.fm33(d_833_v176_.f27, (self).f21, (d_820_v163_) + (d_820_v163_), globalState)
        d_839_i17_: int
        d_839_i17_ = 0
        with _dafny.label("3"):
            while (self).f26:
                with _dafny.c_label("3"):
                    if (d_839_i17_) >= (100):
                        raise _dafny.Break("3")
                    d_839_i17_ = (d_839_i17_) + (1)
                    d_840_v181_: _dafny.Map
                    d_840_v181_ = _dafny.Map({(self).f21: (self).f21})
                    d_841_v182_: _dafny.Map
                    d_841_v182_ = _dafny.Map({not ((self).f26) or (((d_840_v181_)[False] if (False) in (d_840_v181_) else (self).f21)): (290) - ((self).f25)})
                    (d_833_v176_).f27 = ((d_841_v182_)[(self).f26] if ((self).f26) in (d_841_v182_) else ((d_829_v172_)[(self).f20] if ((self).f20) in (d_829_v172_) else (self).f20))
                    (globalState).f5 = default__.safeModuloInt(d_833_v176_.f27, d_833_v176_.f27)
                    d_842_v183_: _dafny.Map
                    d_842_v183_ = _dafny.Map({d_822_v165_: p1})
                    (globalState).f5 = default__.safeModuloInt((d_833_v176_).fm1(d_833_v176_.f27, p0, ((d_842_v183_)[d_822_v165_] if (d_822_v165_) in (d_842_v183_) else (self).f26), globalState), len((_dafny.Map({(self).f21: 793})) | (d_841_v182_)))
                    d_843_v184_: _dafny.Map
                    d_843_v184_ = _dafny.Map({(self).f25: (self).f26})
                    d_843_v184_ = (d_843_v184_).set(p0, ((d_843_v184_)[d_833_v176_.f27] if (d_833_v176_.f27) in (d_843_v184_) else (self).f21))
                    pass
            pass
        (globalState).f14 = False
        d_844_v185_: _dafny.Seq
        d_844_v185_ = _dafny.SeqWithoutIsStrInference([(d_833_v176_).f28])
        d_845_v186_: _dafny.Seq
        d_845_v186_ = _dafny.SeqWithoutIsStrInference([d_844_v185_, d_844_v185_])
        r0 = (d_845_v186_)[default__.safeIndex(((self).f20) - (463), len(d_845_v186_))]
        r1 = True
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: bool = False
        d_846_v0_: _dafny.Set
        d_846_v0_ = _dafny.Set({(self).f26, (self).f26})
        d_847_v1_: _dafny.Map
        d_847_v1_ = _dafny.Map({len(d_846_v0_): 856})
        d_848_v2_: _dafny.Map
        d_848_v2_ = _dafny.Map({(0) - ((0) - ((self).f20)): True})
        d_849_v3_: _dafny.Seq
        d_849_v3_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f20), len(d_848_v2_), (self).f20, (self).f20])
        d_847_v1_ = (d_847_v1_).set(len(d_849_v3_), (0) - ((self).f25))
        d_850_v4_: _dafny.Seq
        d_850_v4_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f26])
        if (d_850_v4_)[default__.safeIndex((self).f20, len(d_850_v4_))]:
            d_851_v5_: _dafny.Array
            def lambda44_(d_852_i0_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjsrayy")) if (self).f21 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asut")))

            init25_ = lambda44_
            nw138_ = _dafny.Array(None, 9)
            for i0_25_ in range(nw138_.length(0)):
                nw138_[i0_25_] = init25_(i0_25_)
            d_851_v5_ = nw138_
            d_853_v6_: _dafny.Seq
            d_853_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxyux"))
            index132_ = default__.safeIndex(949, (d_851_v5_).length(0))
            (d_851_v5_)[index132_] = d_853_v6_
            d_854_v7_: str
            d_854_v7_ = _dafny.CodePoint('l')
            index133_ = default__.safeIndex(949, (d_851_v5_).length(0))
            rhs139_ = (0) - ((len((_dafny.SeqWithoutIsStrInference([(self).f25 for d_855_i1_ in range(default__.abs(402))])) + (_dafny.SeqWithoutIsStrInference([(self).f25])))) * (821))
            rhs140_ = (((d_853_v6_) + (d_853_v6_)).set(default__.safeIndex((self).f20, len((d_853_v6_) + (d_853_v6_))), d_854_v7_)) + (d_853_v6_)
            rhs141_ = (self).f25
            lhs122_ = globalState
            lhs123_ = d_851_v5_
            lhs124_ = default__.safeIndex(949, (d_851_v5_).length(0))
            lhs125_ = globalState
            lhs122_.f17 = rhs139_
            lhs123_[lhs124_] = rhs140_
            lhs125_.f17 = rhs141_
            (globalState).f18 = (self).f26
            d_856_v8_: D2
            d_856_v8_ = D2_DC3((self).f20, (self).f26, not(not((self).f21)), (self).f20, (self).f26)
            d_857_v9_: _dafny.Seq
            d_857_v9_ = _dafny.SeqWithoutIsStrInference([d_846_v0_])
            d_858_v10_: D3
            d_858_v10_ = D3_DC7(d_857_v9_, (self).f26)
            pat_let_tv22_ = d_857_v9_
            def iife73_(_pat_let22_0):
                def iife74_(d_859_dt__update__tmp_h0_):
                    def iife75_(_pat_let23_0):
                        def iife76_(d_860_dt__update_hcf11_h0_):
                            return D3_DC7(d_860_dt__update_hcf11_h0_, (d_859_dt__update__tmp_h0_).cf12)
                        return iife76_(_pat_let23_0)
                    return iife75_(pat_let_tv22_)
                return iife74_(_pat_let22_0)
            source12_ = iife73_((d_858_v10_ if (d_856_v8_).cf7 else default__.fm34(_dafny.CodePoint('v'), (self).f21, (self).f20, (self).f25, globalState)))
            if source12_.is_DC6:
                d_861___mcc_h0_ = source12_.cf10
                d_862_cf10_ = d_861___mcc_h0_
                d_854_v7_ = d_854_v7_
                index134_ = default__.safeIndex(342, (p1).length(0))
                (p1)[index134_] = ((d_847_v1_)[d_862_cf10_] if (d_862_cf10_) in (d_847_v1_) else d_862_cf10_)
                d_863_v11_: _dafny.Set
                d_863_v11_ = _dafny.Set({(self).f25})
                index135_ = default__.safeIndex(342, (p1).length(0))
                (p1)[index135_] = (0) - (len((default__.fm25(globalState)).intersection(d_863_v11_)))
                d_864_v12_: _dafny.Array
                nw139_ = _dafny.Array(None, 4)
                nw139_[int(0)] = (self).f21
                nw139_[int(1)] = (self).f26
                nw139_[int(2)] = (self).f21
                nw139_[int(3)] = True
                d_864_v12_ = nw139_
                d_865_v13_: _dafny.Seq
                d_865_v13_ = _dafny.SeqWithoutIsStrInference([d_864_v12_])
                d_866_v14_: _dafny.MultiSet
                d_866_v14_ = _dafny.MultiSet([(d_865_v13_)[default__.safeIndex((D2_DC3((p1)[default__.safeIndex(342, (p1).length(0))], not((self).f21), False, (self).f20, (self).f21)).cf3, len(d_865_v13_))], d_864_v12_])
                (globalState).f4 = (d_866_v14_).issubset(d_866_v14_)
                (globalState).f5 = (self).f20
            elif source12_.is_DC7:
                d_867___mcc_h1_ = source12_.cf11
                d_868___mcc_h2_ = source12_.cf12
                d_869_cf12_ = d_868___mcc_h2_
                d_870_cf11_ = d_867___mcc_h1_
                index136_ = default__.safeIndex(382, (p1).length(0))
                (p1)[index136_] = (self).f20
                d_871_v15_: C0
                nw140_ = C0()
                nw140_.ctor__()
                d_871_v15_ = nw140_
                d_872_v16_: _dafny.Set
                d_872_v16_ = _dafny.Set({(self).f25})
                d_873_v17_: D4
                d_873_v17_ = D4_DC11((self).f25, d_869_cf12_, (self).f26)
                d_874_v18_: _dafny.Map
                d_874_v18_ = _dafny.Map({(self).f20: d_873_v17_})
                d_875_v19_: _dafny.Map
                d_875_v19_ = _dafny.Map({(self).f25: d_874_v18_})
                index137_ = default__.safeIndex(382, (p1).length(0))
                rhs142_ = not((d_871_v15_).fm3((self).f20, (self).f21, globalState))
                rhs143_ = len((d_875_v19_) | (d_875_v19_))
                rhs144_ = d_871_v15_
                rhs145_ = (d_872_v16_) | ((_dafny.Set({(self).f25})) | (d_872_v16_))
                rhs146_ = d_869_cf12_
                lhs126_ = globalState
                lhs127_ = p1
                lhs128_ = default__.safeIndex(382, (p1).length(0))
                lhs129_ = globalState
                lhs126_.f18 = rhs142_
                lhs127_[lhs128_] = rhs143_
                d_871_v15_ = rhs144_
                d_872_v16_ = rhs145_
                lhs129_.f3 = rhs146_
                (globalState).f17 = (self).f20
                (globalState).f17 = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysfj")))) + ((self).f25)
                index138_ = default__.safeIndex(368, (p0).length(0))
                (p0)[index138_] = d_847_v1_
                index139_ = default__.safeIndex(368, (p0).length(0))
                (p0)[index139_] = d_847_v1_
            elif source12_.is_DC5:
                d_876___mcc_h3_ = source12_.cf9
                d_877_cf9_ = d_876___mcc_h3_
                index140_ = default__.safeIndex(925, (p1).length(0))
                (p1)[index140_] = (self).f25
                index141_ = default__.safeIndex(925, (p1).length(0))
                (p1)[index141_] = len(d_849_v3_)
                (globalState).f17 = -305
                d_878_v20_: C0
                nw141_ = C0()
                nw141_.ctor__()
                d_878_v20_ = nw141_
                d_879_v21_: _dafny.Map
                d_879_v21_ = _dafny.Map({(self).f20: d_878_v20_})
                d_880_v22_: _dafny.Set
                d_880_v22_ = _dafny.Set({d_878_v20_, ((d_879_v21_)[(self).f25] if ((self).f25) in (d_879_v21_) else d_878_v20_), d_878_v20_})
                rhs147_ = d_880_v22_
                rhs148_ = (self).f21
                rhs149_ = ((d_847_v1_)[(self).f20] if ((self).f20) in (d_847_v1_) else (self).f25)
                rhs150_ = (self).f25
                rhs151_ = ((d_849_v3_)[default__.safeIndex((p1)[default__.safeIndex(925, (p1).length(0))], len(d_849_v3_))] if (self).f21 else ((self).f25) - ((p1)[default__.safeIndex(925, (p1).length(0))]))
                lhs130_ = globalState
                lhs131_ = globalState
                lhs132_ = globalState
                lhs133_ = globalState
                d_880_v22_ = rhs147_
                lhs130_.f3 = rhs148_
                lhs131_.f5 = rhs149_
                lhs132_.f17 = rhs150_
                lhs133_.f17 = rhs151_
                (globalState).f4 = (self).f26
            elif True:
                d_881___mcc_h4_ = source12_.cf13
                d_882_cf13_ = d_881___mcc_h4_
                (globalState).f17 = (0) - (len(d_847_v1_))
                index142_ = default__.safeIndex(949, (d_851_v5_).length(0))
                (d_851_v5_)[index142_] = (_dafny.SeqWithoutIsStrInference([default__.fm17((self).f26, globalState)])) + ((_dafny.SeqWithoutIsStrInference([d_854_v7_, _dafny.CodePoint('u')])) + (_dafny.SeqWithoutIsStrInference([d_854_v7_ for d_883_i2_ in range(default__.abs(-627))])))
                d_884_v23_: _dafny.Array
                nw142_ = _dafny.Array(False, 3)
                d_884_v23_ = nw142_
                def lambda45_(d_885_i3_):
                    return ((self).f20) > ((self).f25)

                init26_ = lambda45_
                nw143_ = _dafny.Array(None, 21)
                for i0_26_ in range(nw143_.length(0)):
                    nw143_[i0_26_] = init26_(i0_26_)
                d_884_v23_ = nw143_
                index143_ = default__.safeIndex(428, (p1).length(0))
                (p1)[index143_] = (self).f20
                index144_ = default__.safeIndex(428, (p1).length(0))
                (p1)[index144_] = (self).fm1((0) - ((self).fm1((self).f20, (self).f25, not((self).f26), globalState)), 900, False, globalState)
            (globalState).f17 = (self).f25
            d_853_v6_ = (d_851_v5_)[default__.safeIndex(949, (d_851_v5_).length(0))]
        elif True:
            d_886_v24_: _dafny.Array
            nw144_ = _dafny.Array(int(0), 4)
            d_886_v24_ = nw144_
            d_887_v25_: _dafny.Array
            nw145_ = _dafny.Array(False, 15)
            d_887_v25_ = nw145_
            index145_ = default__.safeIndex(412, (d_887_v25_).length(0))
            (d_887_v25_)[index145_] = (self).f26
            d_888_v26_: _dafny.MultiSet
            d_888_v26_ = _dafny.MultiSet([(self).f21])
            d_889_v27_: D11
            d_889_v27_ = D11_DC31(D11_DC30((self).f26, (d_888_v26_).set(True, default__.abs((self).f20))))
            d_890_v28_: _dafny.Map
            d_890_v28_ = _dafny.Map({d_889_v27_: len(_dafny.SeqWithoutIsStrInference([(self).f25]))})
            d_891_v29_: _dafny.Seq
            d_891_v29_ = _dafny.SeqWithoutIsStrInference([(d_890_v28_) | (d_890_v28_), d_890_v28_])
            d_892_v30_: _dafny.Seq
            d_892_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uielcygj"))
            index146_ = default__.safeIndex(412, (d_887_v25_).length(0))
            rhs152_ = d_886_v24_
            rhs153_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ie"))) + ((d_892_v30_) + (d_892_v30_))
            rhs154_ = False
            rhs155_ = d_891_v29_
            rhs156_ = d_892_v30_
            lhs134_ = globalState
            lhs135_ = d_887_v25_
            lhs136_ = default__.safeIndex(412, (d_887_v25_).length(0))
            lhs137_ = globalState
            d_886_v24_ = rhs152_
            lhs134_.f15 = rhs153_
            lhs135_[lhs136_] = rhs154_
            d_891_v29_ = rhs155_
            lhs137_.f15 = rhs156_
            (globalState).f4 = (d_887_v25_)[default__.safeIndex(412, (d_887_v25_).length(0))]
            index147_ = default__.safeIndex(297, (p1).length(0))
            (p1)[index147_] = (self).f25
            index148_ = default__.safeIndex(297, (p1).length(0))
            (p1)[index148_] = (self).f25
            d_893_v31_: _dafny.Array
            nw146_ = _dafny.Array(None, 27)
            d_893_v31_ = nw146_
            d_894_v32_: T0
            nw147_ = C2()
            nw147_.ctor__(len(d_892_v30_), (d_887_v25_)[default__.safeIndex(412, (d_887_v25_).length(0))])
            d_894_v32_ = nw147_
            index149_ = default__.safeIndex(866, (d_893_v31_).length(0))
            (d_893_v31_)[index149_] = d_894_v32_
            index150_ = default__.safeIndex(866, (d_893_v31_).length(0))
            (d_893_v31_)[index150_] = d_894_v32_
            d_895_v33_: _dafny.Set
            d_895_v33_ = _dafny.Set({d_847_v1_, _dafny.Map({(d_894_v32_).f20: (self).f20}), d_847_v1_, d_847_v1_})
            d_896_v34_: C1
            nw148_ = C1()
            nw148_.ctor__((self).f25, d_886_v24_, len(d_895_v33_), (self).f26)
            d_896_v34_ = nw148_
        d_897_v35_: str
        d_897_v35_ = _dafny.CodePoint('b')
        d_898_v36_: _dafny.Map
        d_898_v36_ = _dafny.Map({(not(not((self).f26))) and ((self).f21): d_897_v35_})
        d_898_v36_ = (d_898_v36_).set((self).f21, d_897_v35_)
        d_899_v37_: _dafny.Seq
        d_899_v37_ = _dafny.SeqWithoutIsStrInference([p1, p1, (p1 if (self).f26 else p1), p1, p1])
        d_899_v37_ = d_899_v37_
        d_900_v38_: _dafny.MultiSet
        d_900_v38_ = _dafny.MultiSet([(self).f21, (self).f26])
        source13_ = D11_DC30((self).f21, d_900_v38_)
        if source13_.is_DC29:
            d_901_v39_: _dafny.Seq
            d_901_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yakhfedov"))
            d_902_v40_: _dafny.Map
            d_902_v40_ = _dafny.Map({(self).f26: False})
            d_903_v41_: _dafny.Seq
            d_903_v41_ = _dafny.SeqWithoutIsStrInference([d_902_v40_, d_902_v40_, _dafny.Map({(self).f26: False}), _dafny.Map({True: (self).f21})])
            rhs157_ = (self).f21
            rhs158_ = (self).fm0(d_847_v1_, d_901_v39_, (d_903_v41_) + (d_903_v41_), globalState)
            lhs138_ = globalState
            r0 = rhs157_
            lhs138_.f4 = rhs158_
            (globalState).f17 = (self).f25
            d_904_v42_: _dafny.Array
            nw149_ = _dafny.Array(_dafny.CodePoint('D'), 23)
            d_904_v42_ = nw149_
            d_905_v43_: _dafny.Array
            nw150_ = _dafny.Array(None, 2)
            nw150_[int(0)] = d_904_v42_
            nw150_[int(1)] = d_904_v42_
            d_905_v43_ = nw150_
            index151_ = default__.safeIndex(293, (d_905_v43_).length(0))
            (d_905_v43_)[index151_] = d_904_v42_
            d_906_v44_: _dafny.Map
            d_906_v44_ = _dafny.Map({(self).f26: 495})
            d_907_v45_: _dafny.Map
            d_907_v45_ = _dafny.Map({d_906_v44_: d_904_v42_})
            d_908_v46_: _dafny.Seq
            d_908_v46_ = _dafny.SeqWithoutIsStrInference([d_904_v42_, ((d_907_v45_)[d_906_v44_] if (d_906_v44_) in (d_907_v45_) else d_904_v42_)])
            index152_ = default__.safeIndex(293, (d_905_v43_).length(0))
            (d_905_v43_)[index152_] = (d_908_v46_)[default__.safeIndex(default__.safeDivisionInt(len(d_901_v39_), (self).f20), len(d_908_v46_))]
            d_909_v47_: _dafny.MultiSet
            d_909_v47_ = _dafny.MultiSet([((d_906_v44_)[(self).f26] if ((self).f26) in (d_906_v44_) else (self).f25)])
            if ((self).f20) not in (d_909_v47_):
                d_910_v48_: _dafny.MultiSet
                d_910_v48_ = _dafny.MultiSet([d_897_v35_])
                d_911_v49_: D7
                d_911_v49_ = D7_DC18(_dafny.SeqWithoutIsStrInference([len(d_901_v39_)]), (self).f20)
                d_912_v50_: C1
                nw151_ = C1()
                nw151_.ctor__((d_910_v48_).cardinality, p1, len((d_911_v49_).cf26), (d_901_v39_) != (d_901_v39_))
                d_912_v50_ = nw151_
                d_897_v35_ = d_897_v35_
                d_913_v51_: _dafny.Set
                d_913_v51_ = _dafny.Set({(self).f20, 438, (self).f25, -926, (self).f20})
                index153_ = default__.safeIndex(791, (p1).length(0))
                (p1)[index153_] = (self).f20
                index154_ = default__.safeIndex(791, (p1).length(0))
                rhs159_ = len(default__.fm35((self).f26, (self).f21, (self).fm1((0) - (d_912_v50_.f27), 981, (self).f26, globalState), globalState))
                rhs160_ = (((self).f26) == ((d_850_v4_)[default__.safeIndex((self).f20, len(d_850_v4_))])) or ((self).f21)
                rhs161_ = (d_913_v51_) | (default__.fm25(globalState))
                rhs162_ = d_912_v50_.f27
                rhs163_ = ((self).f25) > (((d_912_v50_).fm1(d_912_v50_.f27, (self).f25, (self).f21, globalState)) + ((self).f25))
                lhs139_ = globalState
                lhs140_ = p1
                lhs141_ = default__.safeIndex(791, (p1).length(0))
                lhs142_ = globalState
                lhs139_.f5 = rhs159_
                r0 = rhs160_
                d_913_v51_ = rhs161_
                lhs140_[lhs141_] = rhs162_
                lhs142_.f3 = rhs163_
                r0 = ((self).f21) and ((d_901_v39_) != (d_901_v39_))
                d_914_v52_: C0
                nw152_ = C0()
                nw152_.ctor__()
                d_914_v52_ = nw152_
            elif True:
                d_915_v53_: _dafny.Array
                nw153_ = _dafny.Array(False, 16)
                d_915_v53_ = nw153_
                index155_ = default__.safeIndex(93, (d_915_v53_).length(0))
                (d_915_v53_)[index155_] = (self).f26
                index156_ = default__.safeIndex(93, (d_915_v53_).length(0))
                (d_915_v53_)[index156_] = (((d_901_v39_) + (d_901_v39_)).set(default__.safeIndex((self).f20, len((d_901_v39_) + (d_901_v39_))), d_897_v35_)) <= ((d_901_v39_).set(default__.safeIndex((self).f25, len(d_901_v39_)), d_897_v35_))
                (globalState).f17 = default__.safeModuloInt((-536) + ((self).f20), (self).f25)
                d_916_v54_: _dafny.Seq
                d_916_v54_ = _dafny.SeqWithoutIsStrInference([d_915_v53_])
                d_916_v54_ = d_916_v54_
                index157_ = default__.safeIndex(979, (p1).length(0))
                (p1)[index157_] = ((self).f20) - ((self).f20)
                index158_ = default__.safeIndex(979, (p1).length(0))
                rhs164_ = (self).f26
                rhs165_ = ((self).f20) == ((self).f20)
                rhs166_ = d_846_v0_
                rhs167_ = ((d_847_v1_)[((d_849_v3_)[default__.safeIndex((self).f25, len(d_849_v3_))]) * ((self).f20)] if (((d_849_v3_)[default__.safeIndex((self).f25, len(d_849_v3_))]) * ((self).f20)) in (d_847_v1_) else 961)
                lhs143_ = globalState
                lhs144_ = globalState
                lhs145_ = p1
                lhs146_ = default__.safeIndex(979, (p1).length(0))
                lhs143_.f3 = rhs164_
                lhs144_.f4 = rhs165_
                d_846_v0_ = rhs166_
                lhs145_[lhs146_] = rhs167_
                d_849_v3_ = (d_849_v3_ if (d_915_v53_)[default__.safeIndex(93, (d_915_v53_).length(0))] else ((d_849_v3_).set(default__.safeIndex((p1)[default__.safeIndex(979, (p1).length(0))], len(d_849_v3_)), (self).f25)) + (d_849_v3_))
        elif source13_.is_DC30:
            d_917___mcc_h5_ = source13_.cf40
            d_918___mcc_h6_ = source13_.cf41
            d_919_cf41_ = d_918___mcc_h6_
            d_920_cf40_ = d_917___mcc_h5_
            d_921_v55_: _dafny.Array
            def lambda46_(d_922_i4_):
                return (self).f26

            init27_ = lambda46_
            nw154_ = _dafny.Array(None, 21)
            for i0_27_ in range(nw154_.length(0)):
                nw154_[i0_27_] = init27_(i0_27_)
            d_921_v55_ = nw154_
            d_923_v56_: _dafny.Map
            d_923_v56_ = _dafny.Map({(self).f21: d_921_v55_})
            d_924_v57_: D8
            d_924_v57_ = D8_DC21(len((d_923_v56_) | (d_923_v56_)))
            d_924_v57_ = d_924_v57_
            d_925_v58_: _dafny.MultiSet
            d_925_v58_ = _dafny.MultiSet([(self).f25])
            d_847_v1_ = (d_847_v1_).set((d_925_v58_).cardinality, (self).f25)
            if (self).f21:
                d_926_v59_: _dafny.Map
                d_926_v59_ = _dafny.Map({d_897_v35_: (self).f25})
                index159_ = default__.safeIndex(526, (p1).length(0))
                (p1)[index159_] = ((self).f20) + (((d_926_v59_)[d_897_v35_] if (d_897_v35_) in (d_926_v59_) else (self).f20))
                index160_ = default__.safeIndex(526, (p1).length(0))
                (p1)[index160_] = ((self).f20) * (((self).f25) * ((self).f20))
                d_927_v60_: _dafny.Set
                d_927_v60_ = _dafny.Set({(p1)[default__.safeIndex(526, (p1).length(0))], (self).f20, (self).f25, (self).f25, (0) - ((p1)[default__.safeIndex(526, (p1).length(0))])})
                d_928_v61_: _dafny.Seq
                d_928_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwv"))
                d_927_v60_ = _dafny.Set({(((d_925_v58_)[len(d_928_v61_)] if (len(d_928_v61_)) in (d_925_v58_) else len(d_847_v1_))) - ((self).f20), 221, (self).f20, 123, (self).f20})
                d_929_v62_: D8
                d_929_v62_ = D8_DC20()
                d_929_v62_ = default__.fm36(globalState)
                index161_ = default__.safeIndex(526, (p1).length(0))
                (p1)[index161_] = default__.safeDivisionInt((0) - (default__.safeModuloInt((self).f25, len(_dafny.Map({(self).f26: (p1)[default__.safeIndex(526, (p1).length(0))]})))), default__.safeModuloInt((p1)[default__.safeIndex(526, (p1).length(0))], len(d_850_v4_)))
                d_930_v63_: C2
                nw155_ = C2()
                nw155_.ctor__(((self).f25) * ((p1)[default__.safeIndex(526, (p1).length(0))]), ((d_848_v2_)[(self).fm1((p1)[default__.safeIndex(526, (p1).length(0))], (self).f20, True, globalState)] if ((self).fm1((p1)[default__.safeIndex(526, (p1).length(0))], (self).f20, True, globalState)) in (d_848_v2_) else d_920_cf40_))
                d_930_v63_ = nw155_
            elif True:
                (globalState).f17 = ((self).f25) - (default__.safeModuloInt((self).f25, (self).f20))
                d_931_v64_: _dafny.Set
                d_931_v64_ = _dafny.Set({(self).f20, (self).f20, (self).f25, default__.safeModuloInt((self).f20, 752), (129) + ((self).f25)})
                d_932_v65_: _dafny.Array
                nw156_ = _dafny.Array(_dafny.MultiSet({}), 15)
                d_932_v65_ = nw156_
                index162_ = default__.safeIndex(453, (d_932_v65_).length(0))
                (d_932_v65_)[index162_] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f25]))
                index163_ = default__.safeIndex(453, (d_932_v65_).length(0))
                def iife77_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in _dafny.IntegerRange(632, -38):
                        d_933_v66_: int = compr_29_
                        if ((632) <= (d_933_v66_)) and ((d_933_v66_) < (-38)):
                            coll29_ = coll29_.union(_dafny.Set([default__.safeDivisionInt(d_933_v66_, len(d_850_v4_))]))
                    return _dafny.Set(coll29_)
                rhs168_ = iife77_()

                rhs169_ = d_925_v58_
                rhs170_ = -424
                lhs147_ = d_932_v65_
                lhs148_ = default__.safeIndex(453, (d_932_v65_).length(0))
                lhs149_ = globalState
                d_931_v64_ = rhs168_
                lhs147_[lhs148_] = rhs169_
                lhs149_.f17 = rhs170_
                d_934_v67_: _dafny.Seq
                d_934_v67_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f20, 860])])
                d_934_v67_ = d_934_v67_
                d_935_v68_: _dafny.Map
                d_935_v68_ = _dafny.Map({d_850_v4_: (self).f20})
                d_935_v68_ = (d_935_v68_).set(_dafny.SeqWithoutIsStrInference([((d_848_v2_)[len(d_849_v3_)] if (len(d_849_v3_)) in (d_848_v2_) else (self).f21), (self).f26]), len(_dafny.SeqWithoutIsStrInference([d_920_cf40_, (self).f26, (self).f26, False])))
                d_936_v69_: _dafny.MultiSet
                d_936_v69_ = _dafny.MultiSet([d_850_v4_, (d_850_v4_) + (d_850_v4_)])
                d_937_v70_: _dafny.Seq
                d_937_v70_ = _dafny.SeqWithoutIsStrInference([d_850_v4_, d_850_v4_])
                d_938_v71_: _dafny.Seq
                d_938_v71_ = _dafny.SeqWithoutIsStrInference([d_850_v4_, d_850_v4_, (d_937_v70_)[default__.safeIndex((self).f25, len(d_937_v70_))], d_850_v4_])
                d_939_v72_: D12
                d_939_v72_ = D12_DC32(d_938_v71_)
                d_936_v69_ = _dafny.MultiSet((d_939_v72_).cf43)
            d_848_v2_ = (d_848_v2_).set((self).f20, not(((self).f25) != ((self).f25)))
        elif source13_.is_DC28:
            d_940___mcc_h7_ = source13_.cf39
            d_941_cf39_ = d_940___mcc_h7_
            d_942_v73_: _dafny.Map
            d_942_v73_ = _dafny.Map({(self).f26: (self).f25})
            d_943_v74_: C1
            nw157_ = C1()
            nw157_.ctor__(default__.safeModuloInt((self).f25, (self).f20), p1, (0) - (((d_942_v73_)[(self).f26] if ((self).f26) in (d_942_v73_) else (self).f20)), ((self).f26) or ((self).f26))
            d_943_v74_ = nw157_
            d_943_v74_ = d_943_v74_
            d_944_v75_: _dafny.Array
            d_944_v75_ = (d_943_v74_).f28
            d_945_v76_: _dafny.Map
            d_945_v76_ = _dafny.Map({len(d_850_v4_): d_944_v75_})
            d_946_v77_: _dafny.Seq
            d_946_v77_ = _dafny.SeqWithoutIsStrInference([d_945_v76_])
            (d_943_v74_).f27 = (default__.safeModuloInt(d_943_v74_.f27, (0) - (len((d_946_v77_)[default__.safeIndex((self).f25, len(d_946_v77_))])))) + (default__.safeDivisionInt(d_943_v74_.f27, -289))
            d_947_v78_: _dafny.Array
            nw158_ = _dafny.Array(_dafny.MultiSet({}), 16)
            d_947_v78_ = nw158_
            index164_ = default__.safeIndex(945, (d_947_v78_).length(0))
            (d_947_v78_)[index164_] = d_900_v38_
            index165_ = default__.safeIndex(945, (d_947_v78_).length(0))
            (d_947_v78_)[index165_] = _dafny.MultiSet([not((self).f21)])
        elif True:
            d_948___mcc_h8_ = source13_.cf42
            d_949_cf42_ = d_948___mcc_h8_
            d_950_v79_: _dafny.Array
            nw159_ = _dafny.Array(_dafny.Seq({}), 23)
            d_950_v79_ = nw159_
            index166_ = default__.safeIndex(542, (d_950_v79_).length(0))
            (d_950_v79_)[index166_] = d_850_v4_
            d_951_v80_: _dafny.Set
            d_951_v80_ = _dafny.Set({(self).f25})
            d_952_v81_: _dafny.Array
            nw160_ = _dafny.Array(None, 12)
            nw160_[int(0)] = d_897_v35_
            nw160_[int(1)] = d_897_v35_
            nw160_[int(2)] = default__.fm17((self).f26, globalState)
            nw160_[int(3)] = d_897_v35_
            nw160_[int(4)] = d_897_v35_
            nw160_[int(5)] = d_897_v35_
            nw160_[int(6)] = d_897_v35_
            nw160_[int(7)] = d_897_v35_
            nw160_[int(8)] = d_897_v35_
            nw160_[int(9)] = d_897_v35_
            nw160_[int(10)] = d_897_v35_
            nw160_[int(11)] = _dafny.CodePoint('f')
            d_952_v81_ = nw160_
            d_953_v82_: _dafny.Map
            d_953_v82_ = _dafny.Map({_dafny.CodePoint('a'): (self).f25})
            d_954_v83_: _dafny.Map
            d_954_v83_ = _dafny.Map({d_952_v81_: d_953_v82_})
            index167_ = default__.safeIndex(542, (d_950_v79_).length(0))
            rhs171_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, False, (d_951_v80_).ispropersubset(d_951_v80_), (self).f26])
            rhs172_ = ((d_849_v3_)[default__.safeIndex(len(d_849_v3_), len(d_849_v3_))]) > (((self).f20) * ((self).f20))
            rhs173_ = (len(d_954_v83_) if (self).f21 else 153)
            rhs174_ = (self).f25
            rhs175_ = default__.safeModuloInt((self).f25, (0) - ((self).f25))
            lhs150_ = d_950_v79_
            lhs151_ = default__.safeIndex(542, (d_950_v79_).length(0))
            lhs152_ = globalState
            lhs153_ = globalState
            lhs154_ = globalState
            lhs155_ = globalState
            lhs150_[lhs151_] = rhs171_
            lhs152_.f14 = rhs172_
            lhs153_.f17 = rhs173_
            lhs154_.f5 = rhs174_
            lhs155_.f17 = rhs175_
            d_955_v85_: _dafny.Seq
            d_955_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dr"))
            def iife78_():
                coll30_ = _dafny.Map()
                compr_30_: int
                for compr_30_ in _dafny.IntegerRange(587, -160):
                    d_956_v84_: int = compr_30_
                    if ((587) <= (d_956_v84_)) and ((d_956_v84_) < (-160)):
                        coll30_[(d_956_v84_) * ((self).f20)] = (self).f25
                return _dafny.Map(coll30_)
            if not(not((self).fm0((iife78_()
            ) | (d_847_v1_), d_955_v85_, _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f26: (self).f26}) for d_957_i5_ in range(default__.abs(-466))]), globalState))):
                (globalState).f17 = 88
                d_958_v86_: D6
                d_958_v86_ = D6_DC14((self).fm1((self).f20, (self).f20, (self).f21, globalState))
                d_959_v87_: D2
                d_959_v87_ = D2_DC3(len((d_955_v85_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nw"))), len(d_955_v85_)), d_897_v35_)), (self).f26, (self).f26, (self).f20, (self).f21)
                d_960_v88_: _dafny.Map
                d_960_v88_ = _dafny.Map({(self).f26: (self).f21})
                d_961_v89_: _dafny.Seq
                d_961_v89_ = _dafny.SeqWithoutIsStrInference([d_960_v88_])
                d_962_v90_: _dafny.Map
                d_962_v90_ = _dafny.Map({(self).fm0(_dafny.Map({len(d_850_v4_): (self).f25}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwox")), d_961_v89_, globalState): 798})
                d_963_v91_: _dafny.Array
                nw161_ = _dafny.Array(None, 23)
                nw161_[int(0)] = (0) - (len(d_850_v4_))
                nw161_[int(1)] = default__.safeModuloInt((0) - ((self).f25), -804)
                nw161_[int(2)] = (d_958_v86_).cf22
                nw161_[int(3)] = default__.safeDivisionInt((self).f20, (self).f20)
                nw161_[int(4)] = (self).f25
                nw161_[int(5)] = (self).f20
                nw161_[int(6)] = (self).f25
                nw161_[int(7)] = ((self).f25 if (self).f21 else (self).f20)
                nw161_[int(8)] = (self).fm1(len(d_955_v85_), (self).f25, (self).f26, globalState)
                nw161_[int(9)] = (self).f20
                nw161_[int(10)] = (self).f25
                nw161_[int(11)] = default__.safeModuloInt((self).f25, (self).f20)
                nw161_[int(12)] = (self).f25
                nw161_[int(13)] = (d_959_v87_).cf6
                nw161_[int(14)] = (self).f20
                nw161_[int(15)] = (self).f25
                nw161_[int(16)] = ((d_962_v90_)[(self).f26] if ((self).f26) in (d_962_v90_) else (self).f20)
                nw161_[int(17)] = (self).f20
                nw161_[int(18)] = len(d_960_v88_)
                nw161_[int(19)] = (0) - ((self).f25)
                nw161_[int(20)] = (self).f25
                nw161_[int(21)] = (len(d_846_v0_) if (self).f26 else (self).f20)
                nw161_[int(22)] = (self).f25
                d_963_v91_ = nw161_
                rhs176_ = not((self).f21)
                rhs177_ = p1
                lhs156_ = globalState
                lhs156_.f18 = rhs176_
                d_963_v91_ = rhs177_
                (globalState).f3 = ((d_900_v38_) | (default__.fm23(globalState))).isdisjoint((d_900_v38_).intersection(_dafny.MultiSet(default__.fm21(globalState))))
                nw162_ = _dafny.Array(None, 2)
                nw162_[int(0)] = (self).f20
                nw162_[int(1)] = (self).f20
                d_963_v91_ = nw162_
                d_964_v92_: C0
                nw163_ = C0()
                nw163_.ctor__()
                d_964_v92_ = nw163_
            elif True:
                (globalState).f5 = (self).f25
                d_955_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbh"))
                d_965_v93_: _dafny.Array
                d_965_v93_ = p1
                d_965_v93_ = d_965_v93_
                (globalState).f5 = (0) - (len(d_849_v3_))
                d_966_v94_: T0
                nw164_ = C2()
                nw164_.ctor__(340, (self).f21)
                d_966_v94_ = nw164_
                d_967_v95_: _dafny.Map
                d_967_v95_ = _dafny.Map({(d_966_v94_ if (self).f26 else d_966_v94_): (0) - (((self).fm1(((d_847_v1_)[(self).f20] if ((self).f20) in (d_847_v1_) else (self).f20), (self).f25, True, globalState)) * ((d_966_v94_).f20))})
                d_967_v95_ = (d_967_v95_).set((d_966_v94_ if (self).f21 else d_966_v94_), (d_966_v94_).f20)
            d_968_v96_: D8
            d_968_v96_ = D8_DC20()
            d_968_v96_ = d_968_v96_
            d_969_v97_: _dafny.MultiSet
            d_969_v97_ = _dafny.MultiSet([(0) - ((0) - ((self).f25))])
            index168_ = default__.safeIndex(598, (p1).length(0))
            (p1)[index168_] = (((d_969_v97_).set((self).f20, default__.abs(len(d_848_v2_))) if (self).f21 else d_969_v97_)).cardinality
            index169_ = default__.safeIndex(598, (p1).length(0))
            (p1)[index169_] = (-521) * ((self).f25)
        d_970_v98_: _dafny.Array
        def lambda47_(d_971_i6_):
            return D8_DC20()

        init28_ = lambda47_
        nw165_ = _dafny.Array(None, 16)
        for i0_28_ in range(nw165_.length(0)):
            nw165_[i0_28_] = init28_(i0_28_)
        d_970_v98_ = nw165_
        d_972_v99_: _dafny.Array
        def lambda48_(d_973_v35_):
            def lambda49_(d_974_i7_):
                return (D13_DC35(d_973_v35_)).cf45

            return lambda49_

        init29_ = lambda48_(d_897_v35_)
        nw166_ = _dafny.Array(None, 14)
        for i0_29_ in range(nw166_.length(0)):
            nw166_[i0_29_] = init29_(i0_29_)
        d_972_v99_ = nw166_
        d_975_v100_: _dafny.Seq
        d_975_v100_ = _dafny.SeqWithoutIsStrInference([d_972_v99_])
        d_976_v101_: _dafny.Map
        d_976_v101_ = _dafny.Map({p1: (self).f20})
        d_977_v102_: _dafny.Map
        d_977_v102_ = _dafny.Map({(self).f20: d_849_v3_})
        d_978_v103_: _dafny.Seq
        d_978_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cntlqc"))
        rhs178_ = ((d_975_v100_) + (_dafny.SeqWithoutIsStrInference([d_972_v99_, d_972_v99_, d_972_v99_, d_972_v99_, d_972_v99_]))) == (d_975_v100_)
        rhs179_ = default__.safeModuloInt(((d_976_v101_)[p1] if (p1) in (d_976_v101_) else 862), (self).f20)
        rhs180_ = d_970_v98_
        rhs181_ = ((D8_DC21((self).f25)).cf29) + (default__.safeModuloInt(len(d_977_v102_), (0) - ((self).f20)))
        rhs182_ = not ((self).f26) or ((len(d_978_v103_)) == ((self).f20))
        lhs157_ = globalState
        lhs158_ = globalState
        lhs159_ = globalState
        lhs157_.f14 = rhs178_
        lhs158_.f5 = rhs179_
        d_970_v98_ = rhs180_
        lhs159_.f5 = rhs181_
        r0 = rhs182_
        d_979_v104_: D8
        d_979_v104_ = D8_DC19(d_847_v1_)
        d_980_v105_: _dafny.Map
        d_980_v105_ = _dafny.Map({(self).f26: (self).f21})
        d_981_v106_: _dafny.Seq
        d_981_v106_ = _dafny.SeqWithoutIsStrInference([d_980_v105_])
        r0 = (self).fm0((d_979_v104_).cf28, d_978_v103_, d_981_v106_, globalState)
        return r0

    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26

class C4(T0):
    def  __init__(self):
        self._f20: int = int(0)
        self._f21: bool = False
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    def ctor__(self, f24, f20, f21):
        (self)._f24 = f24
        (self)._f20 = f20
        (self)._f21 = f21

    def fm0(self, p0, p1, p2, globalState):
        return (self).f24

    def fm1(self, p0, p1, p2, globalState):
        return ((len(_dafny.Set({(self).f24}))) * ((self).f20)) * ((D4_DC11((self).f20, (self).f21, (self).f24)).cf17)

    def fm13(self, p0, globalState):
        return ((len(_dafny.Map({(self).f24: (self).f21}))) + ((self).f20)) == ((0) - ((self).f20))

    def fm14(self, p0, p1, p2, p3, globalState):
        return (0) - (default__.safeModuloInt((self).f20, (self).f20))

    def m0(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_982_v0_: _dafny.Array
        nw167_ = _dafny.Array(_dafny.Seq({}), 9)
        d_982_v0_ = nw167_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_982_v0_).length(0)):
            d_983_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_983_i0_)) and ((d_983_i0_) < ((d_982_v0_).length(0)))):
                (d_982_v0_)[(d_983_i0_)] = (_dafny.SeqWithoutIsStrInference([(self).f21, (self).f21])) + ((_dafny.SeqWithoutIsStrInference([(self).f24, p1])) + (_dafny.SeqWithoutIsStrInference([(self).f24])))
        d_984_v1_: D4
        d_984_v1_ = D4_DC10(_dafny.MultiSet([(self).f21]), (self).f24)
        def lambda50_(source14_):
            if source14_.is_DC10:
                d_985___mcc_h0_ = source14_.cf15
                d_986___mcc_h1_ = source14_.cf16
                d_987_cf16_ = d_986___mcc_h1_
                d_988_cf15_ = d_985___mcc_h0_
                return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqfaswbj")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_989_i1_ in range(default__.abs(-778))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_990_i2_ in range(default__.abs(640))])])
            elif source14_.is_DC11:
                d_991___mcc_h2_ = source14_.cf17
                d_992___mcc_h3_ = source14_.cf18
                d_993___mcc_h4_ = source14_.cf19
                d_994_cf19_ = d_993___mcc_h4_
                d_995_cf18_ = d_992___mcc_h3_
                d_996_cf17_ = d_991___mcc_h2_
                return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlweiv")) for d_997_i3_ in range(default__.abs(388))])
            elif True:
                d_998___mcc_h5_ = source14_.cf14
                d_999_cf14_ = d_998___mcc_h5_
                return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnq"))])

        (globalState).f5 = len(lambda50_(d_984_v1_))
        d_1000_v2_: D10
        d_1000_v2_ = D10_DC26(p0)
        d_1001_v3_: _dafny.Set
        d_1001_v3_ = _dafny.Set({d_1000_v2_})
        d_1002_v4_: _dafny.Seq
        d_1002_v4_ = _dafny.SeqWithoutIsStrInference([not(False)])
        d_1003_v5_: _dafny.Array
        nw168_ = _dafny.Array(None, 9)
        nw168_[int(0)] = (self).f20
        nw168_[int(1)] = len(d_1002_v4_)
        nw168_[int(2)] = (self).f20
        nw168_[int(3)] = (self).f20
        nw168_[int(4)] = (self).f20
        nw168_[int(5)] = p0
        nw168_[int(6)] = (self).f20
        nw168_[int(7)] = p0
        nw168_[int(8)] = len(_dafny.Set({p0, (self).f20}))
        d_1003_v5_ = nw168_
        d_1004_v6_: T0
        nw169_ = C1()
        nw169_.ctor__(default__.safeModuloInt(749, len(d_1001_v3_)), d_1003_v5_, p0, ((self).f24 if p1 else (self).f24))
        d_1004_v6_ = nw169_
        d_1004_v6_ = d_1004_v6_
        if False:
            d_1005_v7_: _dafny.Seq
            d_1005_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxdrfss"))
            d_1006_v8_: D6
            d_1006_v8_ = D6_DC13((default__.fm26((self).f21, globalState)) + (d_1005_v7_))
            source15_ = d_1006_v8_
            if source15_.is_DC14:
                d_1007___mcc_h6_ = source15_.cf22
                d_1008_cf22_ = d_1007___mcc_h6_
                d_1009_v9_: C1
                nw170_ = C1()
                nw170_.ctor__((0) - (default__.safeModuloInt((0) - ((self).f20), (d_1004_v6_).f20)), d_1003_v5_, len(d_1005_v7_), (d_1004_v6_).f21)
                d_1009_v9_ = nw170_
                (globalState).f5 = (0) - ((d_1004_v6_).f20)
                d_1010_v10_: _dafny.Array
                d_1011_v11_: int
                out14_: _dafny.Array
                out15_: int
                out14_, out15_ = (self).m1(d_1008_cf22_, (d_1004_v6_).f21, d_1004_v6_, globalState)
                d_1010_v10_ = out14_
                d_1011_v11_ = out15_
                (d_1009_v9_).f27 = (d_1004_v6_).f20
            elif source15_.is_DC15:
                d_1012___mcc_h7_ = source15_.cf23
                d_1013_cf23_ = d_1012___mcc_h7_
                d_1014_v12_: _dafny.MultiSet
                d_1014_v12_ = _dafny.MultiSet([(self).f20])
                d_1015_v13_: _dafny.Seq
                d_1015_v13_ = _dafny.SeqWithoutIsStrInference([(d_1004_v6_).f20, len(d_1005_v7_), len(d_1005_v7_), p0, (d_1004_v6_).f20])
                d_1016_v14_: _dafny.MultiSet
                d_1016_v14_ = _dafny.MultiSet([p1])
                d_1017_v15_: _dafny.Seq
                d_1017_v15_ = _dafny.SeqWithoutIsStrInference([(d_1004_v6_).f20, (self).f20, (self).fm14((d_1004_v6_).f21, (d_1004_v6_).f20, ((d_1014_v12_)[(d_1015_v13_)[default__.safeIndex(((d_1014_v12_)[(d_1016_v14_).cardinality] if ((d_1016_v14_).cardinality) in (d_1014_v12_) else (d_1004_v6_).f20), len(d_1015_v13_))]] if ((d_1015_v13_)[default__.safeIndex(((d_1014_v12_)[(d_1016_v14_).cardinality] if ((d_1016_v14_).cardinality) in (d_1014_v12_) else (d_1004_v6_).f20), len(d_1015_v13_))]) in (d_1014_v12_) else (self).f20), p1, globalState)])
                d_1017_v15_ = d_1017_v15_
                d_1018_v16_: C0
                nw171_ = C0()
                nw171_.ctor__()
                d_1018_v16_ = nw171_
                (globalState).f3 = not ((self).f21) or (((d_1004_v6_).f21) and ((D13_DC36(p0, not(True), (self).f24)).cf48))
                d_1019_v17_: C2
                nw172_ = C2()
                nw172_.ctor__(default__.safeModuloInt((d_1004_v6_).f20, (d_1004_v6_).f20), (p0) != (len(d_1005_v7_)))
                d_1019_v17_ = nw172_
            elif True:
                d_1020___mcc_h8_ = source15_.cf21
                d_1021_cf21_ = d_1020___mcc_h8_
                d_1022_v18_: _dafny.Array
                def lambda51_(d_1023_i4_):
                    return (self).f21

                init30_ = lambda51_
                nw173_ = _dafny.Array(None, 22)
                for i0_30_ in range(nw173_.length(0)):
                    nw173_[i0_30_] = init30_(i0_30_)
                d_1022_v18_ = nw173_
                index170_ = default__.safeIndex(292, (d_1022_v18_).length(0))
                (d_1022_v18_)[index170_] = False
                index171_ = default__.safeIndex(292, (d_1022_v18_).length(0))
                rhs183_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1024_i5_ in range(default__.abs(542))])) == (d_1021_cf21_)
                rhs184_ = not((self).f24)
                lhs160_ = d_1022_v18_
                lhs161_ = default__.safeIndex(292, (d_1022_v18_).length(0))
                lhs162_ = globalState
                lhs160_[lhs161_] = rhs183_
                lhs162_.f14 = rhs184_
                d_1003_v5_ = d_1003_v5_
                d_1025_v19_: _dafny.Map
                d_1025_v19_ = _dafny.Map({((d_1004_v6_).f21) == ((self).f24): p0})
                d_1025_v19_ = (d_1025_v19_).set((d_1004_v6_).f21, p0)
                d_1026_v20_: _dafny.Seq
                d_1026_v20_ = _dafny.SeqWithoutIsStrInference([p0, 478, -570])
                index172_ = default__.safeIndex(292, (d_1022_v18_).length(0))
                rhs185_ = (d_1004_v6_).f21
                rhs186_ = len(d_1026_v20_)
                rhs187_ = (p0) + (len(d_1026_v20_))
                rhs188_ = (self).f21
                rhs189_ = (d_1004_v6_).f21
                lhs163_ = globalState
                lhs164_ = globalState
                lhs165_ = globalState
                lhs166_ = d_1022_v18_
                lhs167_ = default__.safeIndex(292, (d_1022_v18_).length(0))
                lhs168_ = globalState
                lhs163_.f18 = rhs185_
                lhs164_.f5 = rhs186_
                lhs165_.f5 = rhs187_
                lhs166_[lhs167_] = rhs188_
                lhs168_.f14 = rhs189_
            d_1027_v21_: _dafny.Array
            def lambda52_(d_1028_v7_):
                def lambda53_(d_1029_i6_):
                    return d_1028_v7_

                return lambda53_

            init31_ = lambda52_(d_1005_v7_)
            nw174_ = _dafny.Array(None, 20)
            for i0_31_ in range(nw174_.length(0)):
                nw174_[i0_31_] = init31_(i0_31_)
            d_1027_v21_ = nw174_
            index173_ = default__.safeIndex(433, (d_1027_v21_).length(0))
            (d_1027_v21_)[index173_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1030_i7_ in range(default__.abs(46))])
            d_1031_v22_: _dafny.Seq
            d_1031_v22_ = _dafny.SeqWithoutIsStrInference([p0, p0, (d_1004_v6_).f20])
            d_1032_v23_: _dafny.Seq
            d_1032_v23_ = _dafny.SeqWithoutIsStrInference([d_1031_v22_, d_1031_v22_, d_1031_v22_])
            d_1033_v24_: D9
            d_1033_v24_ = D9_DC22(d_1002_v4_)
            d_1034_v25_: D9
            d_1034_v25_ = D9_DC24(d_1033_v24_)
            d_1035_v26_: D9
            d_1035_v26_ = D9_DC24(d_1034_v25_)
            d_1036_v27_: str
            d_1036_v27_ = _dafny.CodePoint('j')
            d_1037_v28_: _dafny.Map
            d_1037_v28_ = _dafny.Map({(d_1004_v6_).f20: (d_1004_v6_).f20})
            d_1038_v29_: _dafny.Map
            d_1038_v29_ = _dafny.Map({True: (self).f21})
            d_1039_v30_: _dafny.Seq
            d_1039_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f24: (d_1004_v6_).f21}), d_1038_v29_])
            d_1040_v31_: C3
            nw175_ = C3()
            nw175_.ctor__(382, (self).fm0(d_1037_v28_, d_1005_v7_, d_1039_v30_, globalState), (d_1004_v6_).f20, (self).fm13((d_1004_v6_).f20, globalState))
            d_1040_v31_ = nw175_
            d_1041_v32_: _dafny.Map
            d_1041_v32_ = _dafny.Map({d_1040_v31_: d_1032_v23_})
            index174_ = default__.safeIndex(433, (d_1027_v21_).length(0))
            rhs190_ = (d_1005_v7_).set(default__.safeIndex((d_1004_v6_).f20, len(d_1005_v7_)), d_1036_v27_)
            rhs191_ = default__.safeModuloInt(657, p0)
            rhs192_ = ((d_1032_v23_) + (((d_1041_v32_)[d_1040_v31_] if (d_1040_v31_) in (d_1041_v32_) else d_1032_v23_))) + (d_1032_v23_)
            rhs193_ = d_1035_v26_
            rhs194_ = d_1027_v21_
            lhs169_ = d_1027_v21_
            lhs170_ = default__.safeIndex(433, (d_1027_v21_).length(0))
            lhs171_ = globalState
            lhs169_[lhs170_] = rhs190_
            lhs171_.f5 = rhs191_
            d_1032_v23_ = rhs192_
            d_1035_v26_ = rhs193_
            d_1027_v21_ = rhs194_
            d_1042_v33_: _dafny.Array
            nw176_ = _dafny.Array(False, 15)
            d_1042_v33_ = nw176_
            index175_ = default__.safeIndex(620, (d_1042_v33_).length(0))
            def iife79_():
                coll31_ = _dafny.Set()
                compr_31_: int
                for compr_31_ in _dafny.IntegerRange(841, 77):
                    d_1043_v34_: int = compr_31_
                    if ((841) <= (d_1043_v34_)) and ((d_1043_v34_) < (77)):
                        coll31_ = coll31_.union(_dafny.Set([(d_1043_v34_) * ((d_1004_v6_).f20)]))
                return _dafny.Set(coll31_)
            (d_1042_v33_)[index175_] = (iife79_()
            ).isdisjoint(_dafny.Set({987}))
            index176_ = default__.safeIndex(620, (d_1042_v33_).length(0))
            (d_1042_v33_)[index176_] = (len(d_1005_v7_)) > (len(d_1005_v7_))
            d_1044_v35_: _dafny.Map
            d_1044_v35_ = _dafny.Map({(default__.fm34(d_1036_v27_, (self).f21, (d_1004_v6_).f20, p0, globalState)).cf12: -844})
            d_1045_v36_: _dafny.MultiSet
            d_1045_v36_ = _dafny.MultiSet([p0])
            d_1046_v37_: _dafny.Seq
            d_1046_v37_ = _dafny.SeqWithoutIsStrInference([d_1044_v35_, (_dafny.Map({(d_1042_v33_)[default__.safeIndex(620, (d_1042_v33_).length(0))]: (d_1045_v36_).cardinality})) | (d_1044_v35_), (_dafny.Map({(d_1042_v33_)[default__.safeIndex(620, (d_1042_v33_).length(0))]: p0}) if (self).f21 else d_1044_v35_), (d_1044_v35_) | (default__.fm27((d_1004_v6_).f20, globalState)), d_1044_v35_])
            d_1046_v37_ = d_1046_v37_
            d_1047_v38_: C1
            nw177_ = C1()
            nw177_.ctor__(p0, d_1003_v5_, 342, (d_1040_v31_).f26)
            d_1047_v38_ = nw177_
            d_1048_v39_: _dafny.Map
            d_1048_v39_ = _dafny.Map({(d_1004_v6_).f20: d_1047_v38_})
            d_1049_v40_: C1
            nw178_ = C1()
            nw178_.ctor__((0) - (p0), d_1003_v5_, len(d_1048_v39_), (d_1047_v38_.f27) != ((d_1004_v6_).f20))
            d_1049_v40_ = nw178_
        elif True:
            (globalState).f18 = (d_1004_v6_).f21
            rhs195_ = p0
            rhs196_ = (d_1004_v6_).f20
            rhs197_ = (p1) and (((d_1004_v6_).f20) <= (p0))
            lhs172_ = globalState
            lhs173_ = globalState
            lhs174_ = globalState
            lhs172_.f17 = rhs195_
            lhs173_.f17 = rhs196_
            lhs174_.f14 = rhs197_
            d_1050_v41_: _dafny.Map
            d_1050_v41_ = _dafny.Map({p0: not((d_1004_v6_).f21)})
            d_1051_v42_: _dafny.Map
            d_1051_v42_ = _dafny.Map({not(((d_1050_v41_)[(self).f20] if ((self).f20) in (d_1050_v41_) else (self).f21)): -741})
            d_1051_v42_ = (d_1051_v42_).set((p1 if (d_1004_v6_).f21 else (d_1004_v6_).f21), (self).f20)
            d_1003_v5_ = d_1003_v5_
            (globalState).f17 = default__.safeModuloInt((self).f20, (d_1004_v6_).f20)
        d_1052_v43_: _dafny.Set
        d_1052_v43_ = _dafny.Set({(self).f21, p1})
        d_1053_v44_: _dafny.Set
        d_1053_v44_ = _dafny.Set({len(d_1052_v43_)})
        d_1054_v45_: _dafny.MultiSet
        d_1054_v45_ = _dafny.MultiSet([len(d_1053_v44_)])
        if (d_1054_v45_).ispropersubset((d_1054_v45_) - (d_1054_v45_)):
            d_1055_v46_: D13
            d_1055_v46_ = D13_DC36((self).f20, not(not(((d_1004_v6_).f21) and ((d_1004_v6_).f21))), (d_1002_v4_)[default__.safeIndex((d_1004_v6_).f20, len(d_1002_v4_))])
            d_1056_v47_: _dafny.Map
            d_1056_v47_ = _dafny.Map({(self).f20: (self).f20})
            d_1057_v48_: _dafny.Seq
            d_1057_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hadgdwudf"))
            rhs198_ = (default__.safeModuloInt((d_1004_v6_).f20, (self).f20)) + (((d_1056_v47_)[p0] if (p0) in (d_1056_v47_) else (self).f20))
            rhs199_ = len(d_1057_v48_)
            rhs200_ = d_1055_v46_
            lhs175_ = globalState
            lhs176_ = globalState
            lhs175_.f5 = rhs198_
            lhs176_.f5 = rhs199_
            d_1055_v46_ = rhs200_
            d_1058_v49_: _dafny.Map
            d_1058_v49_ = _dafny.Map({len(d_1057_v48_): d_1004_v6_})
            if ((d_1004_v6_).f20) > (len(d_1058_v49_)):
                d_1003_v5_ = d_1003_v5_
                (globalState).f5 = (d_1004_v6_).f20
                d_1059_v50_: _dafny.MultiSet
                d_1059_v50_ = _dafny.MultiSet([p1])
                d_1060_v51_: _dafny.Seq
                d_1060_v51_ = _dafny.SeqWithoutIsStrInference([((d_1059_v50_)[(self).f21] if ((self).f21) in (d_1059_v50_) else (d_1004_v6_).fm1((d_1004_v6_).f20, len(_dafny.Set({(self).f20, p0, 809})), (d_1004_v6_).f21, globalState)), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1061_i8_ in range(default__.abs(225))]))])
                d_1062_v52_: str
                d_1062_v52_ = _dafny.CodePoint('v')
                (globalState).f18 = (d_1057_v48_) <= ((d_1057_v48_ if (d_1004_v6_).f21 else (d_1057_v48_).set(default__.safeIndex(len(d_1060_v51_), len(d_1057_v48_)), d_1062_v52_)))
                d_1063_v53_: _dafny.Array
                nw179_ = _dafny.Array(False, 1)
                d_1063_v53_ = nw179_
                d_1064_v56_: _dafny.Map
                def iife80_():
                    coll32_ = _dafny.Map()
                    compr_32_: str
                    for compr_32_ in (_dafny.SeqWithoutIsStrInference([d_1062_v52_ for d_1065_i9_ in range(default__.abs(367))])).Elements:
                        d_1066_v55_: str = compr_32_
                        if (d_1066_v55_) in (_dafny.SeqWithoutIsStrInference([d_1062_v52_ for d_1065_i9_ in range(default__.abs(367))])):
                            coll32_[d_1066_v55_] = (d_1004_v6_).f20
                    return _dafny.Map(coll32_)
                d_1064_v56_ = _dafny.Map({default__.fm37(len(d_1057_v48_), globalState): len(iife80_()
                )})
                index177_ = default__.safeIndex(919, (d_1063_v53_).length(0))
                def iife81_():
                    coll33_ = _dafny.Map()
                    compr_33_: _dafny.MultiSet
                    for compr_33_ in (d_1064_v56_).keys.Elements:
                        d_1067_v54_: _dafny.MultiSet = compr_33_
                        if (d_1067_v54_) in (d_1064_v56_):
                            coll33_[d_1067_v54_] = len(_dafny.SeqWithoutIsStrInference([d_1062_v52_ for d_1068_i10_ in range(default__.abs(608))]))
                    return _dafny.Map(coll33_)
                (d_1063_v53_)[index177_] = (iife81_()
                ) == (_dafny.Map({_dafny.MultiSet([(self).f20]): (self).f20}))
                index178_ = default__.safeIndex(919, (d_1063_v53_).length(0))
                (d_1063_v53_)[index178_] = (d_1004_v6_).f21
                index179_ = default__.safeIndex(919, (d_1063_v53_).length(0))
                (d_1063_v53_)[index179_] = not ((d_1063_v53_)[default__.safeIndex(919, (d_1063_v53_).length(0))]) or ((p0) == (208))
            elif True:
                d_1069_v57_: C3
                nw180_ = C3()
                nw180_.ctor__(((d_1004_v6_).f20) - ((d_1004_v6_).f20), (p0) >= ((d_1004_v6_).f20), 582, ((d_1004_v6_).f21) == ((d_1004_v6_).f21))
                d_1069_v57_ = nw180_
                d_1070_v58_: _dafny.Seq
                d_1070_v58_ = _dafny.SeqWithoutIsStrInference([251])
                d_1071_v59_: _dafny.MultiSet
                d_1071_v59_ = _dafny.MultiSet([d_1053_v44_])
                d_1072_v60_: _dafny.Map
                d_1072_v60_ = _dafny.Map({(d_1004_v6_).f21: (d_1069_v57_).f25})
                d_1073_v61_: str
                d_1073_v61_ = _dafny.CodePoint('m')
                d_1074_v62_: _dafny.Map
                d_1074_v62_ = _dafny.Map({d_1072_v60_: d_1073_v61_})
                d_1075_v63_: _dafny.Array
                nw181_ = _dafny.Array(None, 11)
                nw181_[int(0)] = d_1070_v58_
                nw181_[int(1)] = ((d_1070_v58_).set(default__.safeIndex(((d_1071_v59_)[d_1053_v44_] if (d_1053_v44_) in (d_1071_v59_) else len(d_1057_v48_)), len(d_1070_v58_)), len(d_1057_v48_))).set(default__.safeIndex(len(d_1074_v62_), len((d_1070_v58_).set(default__.safeIndex(((d_1071_v59_)[d_1053_v44_] if (d_1053_v44_) in (d_1071_v59_) else len(d_1057_v48_)), len(d_1070_v58_)), len(d_1057_v48_)))), (d_1004_v6_).f20)
                nw181_[int(2)] = (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxcytgdfb"))) for d_1076_i11_ in range(default__.abs(73))])) + (d_1070_v58_)
                nw181_[int(3)] = _dafny.SeqWithoutIsStrInference([(d_1004_v6_).f20])
                nw181_[int(4)] = _dafny.SeqWithoutIsStrInference([p0 for d_1077_i12_ in range(default__.abs(636))])
                nw181_[int(5)] = default__.fm28((d_1004_v6_).f20, globalState)
                nw181_[int(6)] = d_1070_v58_
                nw181_[int(7)] = d_1070_v58_
                nw181_[int(8)] = d_1070_v58_
                nw181_[int(9)] = d_1070_v58_
                nw181_[int(10)] = d_1070_v58_
                d_1075_v63_ = nw181_
                index180_ = default__.safeIndex(391, (d_1075_v63_).length(0))
                (d_1075_v63_)[index180_] = d_1070_v58_
                d_1078_v64_: _dafny.MultiSet
                d_1078_v64_ = _dafny.MultiSet([True, (self).f24])
                index181_ = default__.safeIndex(391, (d_1075_v63_).length(0))
                rhs201_ = (0) - ((((d_1078_v64_).cardinality) + (908)) + ((d_1004_v6_).f20))
                rhs202_ = ((d_1070_v58_).set(default__.safeIndex(p0, len(d_1070_v58_)), len(d_1002_v4_))) + (default__.fm28((d_1004_v6_).f20, globalState))
                lhs177_ = globalState
                lhs178_ = d_1075_v63_
                lhs179_ = default__.safeIndex(391, (d_1075_v63_).length(0))
                lhs177_.f17 = rhs201_
                lhs178_[lhs179_] = rhs202_
                (globalState).f18 = not ((d_1004_v6_).f21) or (not(not((d_1004_v6_).f21)))
                d_1003_v5_ = d_1003_v5_
                d_1079_v65_: _dafny.Array
                nw182_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_1079_v65_ = nw182_
                d_1080_v66_: _dafny.Map
                d_1080_v66_ = _dafny.Map({d_1079_v65_: (d_1004_v6_).f21})
                d_1081_v67_: _dafny.Map
                d_1081_v67_ = _dafny.Map({_dafny.Map({(d_1004_v6_).f20: p1}): (d_1069_v57_).f25})
                d_1082_v68_: _dafny.Map
                d_1082_v68_ = _dafny.Map({(d_1004_v6_).f20: (self).f21})
                (globalState).f17 = len(((default__.fm38((0) - ((d_1069_v57_).f25), ((d_1080_v66_)[d_1079_v65_] if (d_1079_v65_) in (d_1080_v66_) else p1), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrgiiat"))), (d_1004_v6_).f21, globalState) if (self).f21 else d_1081_v67_)).set(d_1082_v68_, (p0) - ((d_1069_v57_).f25)))
            (globalState).f5 = (d_1004_v6_).f20
            d_1083_v69_: _dafny.Seq
            d_1083_v69_ = _dafny.SeqWithoutIsStrInference([d_1057_v48_])
            d_1084_v70_: _dafny.Array
            nw183_ = _dafny.Array(None, 13)
            nw183_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1085_i13_ in range(default__.abs(-767))])
            nw183_[int(1)] = d_1057_v48_
            nw183_[int(2)] = default__.fm26((d_1004_v6_).f21, globalState)
            nw183_[int(3)] = d_1057_v48_
            nw183_[int(4)] = (default__.fm26((self).f24, globalState)) + (d_1057_v48_)
            nw183_[int(5)] = (d_1057_v48_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gntqqvui")))
            nw183_[int(6)] = d_1057_v48_
            nw183_[int(7)] = ((d_1083_v69_)[default__.safeIndex((self).f20, len(d_1083_v69_))]) + (d_1057_v48_)
            nw183_[int(8)] = d_1057_v48_
            nw183_[int(9)] = d_1057_v48_
            nw183_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxfbmf"))
            nw183_[int(11)] = d_1057_v48_
            nw183_[int(12)] = d_1057_v48_
            d_1084_v70_ = nw183_
            d_1084_v70_ = d_1084_v70_
            if not(True):
                (globalState).f4 = (d_1004_v6_).f21
                (globalState).f15 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhpdkduv"))
                d_1086_v71_: str
                d_1086_v71_ = _dafny.CodePoint('r')
                d_1087_v72_: _dafny.Array
                def lambda54_(d_1088_i14_):
                    return (self).f21

                init32_ = lambda54_
                nw184_ = _dafny.Array(None, 16)
                for i0_32_ in range(nw184_.length(0)):
                    nw184_[i0_32_] = init32_(i0_32_)
                d_1087_v72_ = nw184_
                index182_ = default__.safeIndex(757, (d_1087_v72_).length(0))
                (d_1087_v72_)[index182_] = (d_1004_v6_).f21
                d_1089_v73_: _dafny.Map
                d_1089_v73_ = _dafny.Map({(d_1004_v6_).f21: (self).f21})
                d_1090_v74_: _dafny.Map
                d_1090_v74_ = _dafny.Map({not(True): ((d_1089_v73_).set((self).f24, (d_1004_v6_).f21)) | (d_1089_v73_)})
                d_1091_v75_: _dafny.Array
                nw185_ = _dafny.Array(None, 23)
                nw185_[int(0)] = d_1003_v5_
                nw185_[int(1)] = d_1003_v5_
                nw185_[int(2)] = d_1003_v5_
                nw185_[int(3)] = d_1003_v5_
                nw185_[int(4)] = d_1003_v5_
                nw185_[int(5)] = d_1003_v5_
                nw185_[int(6)] = d_1003_v5_
                nw185_[int(7)] = d_1003_v5_
                nw185_[int(8)] = d_1003_v5_
                nw185_[int(9)] = d_1003_v5_
                nw185_[int(10)] = d_1003_v5_
                nw185_[int(11)] = d_1003_v5_
                nw185_[int(12)] = d_1003_v5_
                nw185_[int(13)] = d_1003_v5_
                nw185_[int(14)] = d_1003_v5_
                nw185_[int(15)] = d_1003_v5_
                nw185_[int(16)] = d_1003_v5_
                nw185_[int(17)] = d_1003_v5_
                nw185_[int(18)] = d_1003_v5_
                nw185_[int(19)] = d_1003_v5_
                nw185_[int(20)] = (d_1003_v5_ if True else d_1003_v5_)
                nw185_[int(21)] = (d_1003_v5_ if (self).f24 else d_1003_v5_)
                nw185_[int(22)] = (d_1003_v5_ if (d_1004_v6_).f21 else d_1003_v5_)
                d_1091_v75_ = nw185_
                index183_ = default__.safeIndex(63, (d_1091_v75_).length(0))
                (d_1091_v75_)[index183_] = d_1003_v5_
                index184_ = default__.safeIndex(757, (d_1087_v72_).length(0))
                index185_ = default__.safeIndex(63, (d_1091_v75_).length(0))
                rhs203_ = d_1086_v71_
                rhs204_ = (d_1004_v6_).f21
                rhs205_ = (d_1004_v6_).f21
                rhs206_ = d_1090_v74_
                rhs207_ = d_1003_v5_
                lhs180_ = globalState
                lhs181_ = d_1087_v72_
                lhs182_ = default__.safeIndex(757, (d_1087_v72_).length(0))
                lhs183_ = d_1091_v75_
                lhs184_ = default__.safeIndex(63, (d_1091_v75_).length(0))
                d_1086_v71_ = rhs203_
                lhs180_.f14 = rhs204_
                lhs181_[lhs182_] = rhs205_
                d_1090_v74_ = rhs206_
                lhs183_[lhs184_] = rhs207_
                d_1092_v76_: _dafny.MultiSet
                d_1092_v76_ = _dafny.MultiSet([p1])
                d_1093_v77_: _dafny.Map
                d_1093_v77_ = _dafny.Map({(d_1004_v6_).f21: (d_1004_v6_).f20})
                d_1094_v78_: _dafny.Seq
                d_1094_v78_ = _dafny.SeqWithoutIsStrInference([(self).f20, (d_1004_v6_).f20, (self).f20, len(d_1093_v77_)])
                rhs208_ = (((d_1092_v76_).set((d_1004_v6_).f21, default__.abs(len(_dafny.SeqWithoutIsStrInference([((d_1054_v45_)[(d_1004_v6_).f20] if ((d_1004_v6_).f20) in (d_1054_v45_) else (self).f20), (self).f20, len(d_1094_v78_), (d_1004_v6_).f20]))))) | (_dafny.MultiSet(d_1002_v4_))).ispropersubset(d_1092_v76_)
                rhs209_ = default__.safeDivisionInt((d_1004_v6_).f20, len(((d_1083_v69_)[default__.safeIndex((d_1004_v6_).f20, len(d_1083_v69_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwncqvclf")))))
                rhs210_ = (d_1004_v6_).f20
                lhs185_ = globalState
                lhs186_ = globalState
                lhs187_ = globalState
                lhs185_.f4 = rhs208_
                lhs186_.f5 = rhs209_
                lhs187_.f17 = rhs210_
                d_1095_v79_: _dafny.Seq
                d_1095_v79_ = _dafny.SeqWithoutIsStrInference([d_1054_v45_, (d_1054_v45_) | (d_1054_v45_)])
                arr2_ = (d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]
                index186_ = default__.safeIndex(754, ((d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]).length(0))
                arr2_[index186_] = default__.safeModuloInt((self).f20, (d_1004_v6_).f20)
                index187_ = default__.safeIndex(623, (d_1084_v70_).length(0))
                (d_1084_v70_)[index187_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                arr3_ = (d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]
                index188_ = default__.safeIndex(400, ((d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]).length(0))
                arr3_[index188_] = ((0) - ((d_1004_v6_).f20)) * (p0)
                arr4_ = (d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]
                index189_ = default__.safeIndex(754, ((d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]).length(0))
                index190_ = default__.safeIndex(623, (d_1084_v70_).length(0))
                arr5_ = (d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]
                index191_ = default__.safeIndex(400, ((d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]).length(0))
                rhs211_ = _dafny.SeqWithoutIsStrInference([d_1054_v45_ for d_1096_i15_ in range(default__.abs(185))])
                rhs212_ = (0) - ((self).fm14((d_1092_v76_).issubset(_dafny.MultiSet([not((d_1004_v6_).f21), (d_1004_v6_).f21, not((d_1004_v6_).f21), (d_1004_v6_).f21])), (d_1004_v6_).f20, (d_1004_v6_).f20, (d_1004_v6_).f21, globalState))
                rhs213_ = (d_1004_v6_).f20
                rhs214_ = (d_1057_v48_) + (d_1057_v48_)
                rhs215_ = (((d_1092_v76_)[(d_1004_v6_).f21] if ((d_1004_v6_).f21) in (d_1092_v76_) else (0) - ((d_1004_v6_).f20))) + (default__.safeDivisionInt((d_1004_v6_).f20, (d_1004_v6_).f20))
                lhs188_ = (d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]
                lhs189_ = default__.safeIndex(754, ((d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]).length(0))
                lhs190_ = globalState
                lhs191_ = d_1084_v70_
                lhs192_ = default__.safeIndex(623, (d_1084_v70_).length(0))
                lhs193_ = (d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]
                lhs194_ = default__.safeIndex(400, ((d_1091_v75_)[default__.safeIndex(63, (d_1091_v75_).length(0))]).length(0))
                d_1095_v79_ = rhs211_
                lhs188_[lhs189_] = rhs212_
                lhs190_.f5 = rhs213_
                lhs191_[lhs192_] = rhs214_
                lhs193_[lhs194_] = rhs215_
            elif True:
                (globalState).f17 = (d_1004_v6_).fm1((d_1004_v6_).f20, p0, (self).f21, globalState)
                d_1097_v80_: _dafny.Seq
                d_1097_v80_ = _dafny.SeqWithoutIsStrInference([(0) - (p0), (d_1004_v6_).f20, (self).f20, (d_1004_v6_).f20])
                d_1097_v80_ = (d_1097_v80_) + (d_1097_v80_)
                d_1098_v81_: C2
                nw186_ = C2()
                nw186_.ctor__(default__.safeDivisionInt((d_1004_v6_).f20, (d_1004_v6_).f20), (self).fm13((d_1004_v6_).f20, globalState))
                d_1098_v81_ = nw186_
                d_1099_v82_: _dafny.Array
                def lambda55_(d_1100_v6_):
                    def lambda56_(d_1101_i16_):
                        return not((d_1100_v6_).f21)

                    return lambda56_

                init33_ = lambda55_(d_1004_v6_)
                nw187_ = _dafny.Array(None, 15)
                for i0_33_ in range(nw187_.length(0)):
                    nw187_[i0_33_] = init33_(i0_33_)
                d_1099_v82_ = nw187_
                d_1102_v83_: _dafny.MultiSet
                d_1102_v83_ = _dafny.MultiSet([d_1099_v82_, d_1099_v82_])
                d_1102_v83_ = _dafny.MultiSet([d_1099_v82_])
                index192_ = default__.safeIndex(528, (d_1099_v82_).length(0))
                (d_1099_v82_)[index192_] = (self).f24
                index193_ = default__.safeIndex(902, (d_1099_v82_).length(0))
                (d_1099_v82_)[index193_] = (self).f24
                d_1103_v84_: _dafny.MultiSet
                d_1103_v84_ = _dafny.MultiSet([(d_1057_v48_)[default__.safeIndex(p0, len(d_1057_v48_))]])
                d_1104_v85_: bool
                d_1104_v85_ = (d_1004_v6_).f21
                d_1105_v86_: _dafny.Seq
                d_1105_v86_ = d_1097_v80_
                index194_ = default__.safeIndex(528, (d_1099_v82_).length(0))
                index195_ = default__.safeIndex(902, (d_1099_v82_).length(0))
                rhs216_ = not(((d_1103_v84_).cardinality) <= ((d_1004_v6_).f20))
                rhs217_ = ((self).f21 if (d_1004_v6_).f21 else (d_1104_v85_))
                rhs218_ = (d_1004_v6_).f21
                rhs219_ = (d_1004_v6_).f21
                rhs220_ = (0) - (len((d_1097_v80_) + ((d_1097_v80_ if (self).f24 else (d_1105_v86_)))))
                lhs195_ = d_1099_v82_
                lhs196_ = default__.safeIndex(528, (d_1099_v82_).length(0))
                lhs197_ = globalState
                lhs198_ = d_1099_v82_
                lhs199_ = default__.safeIndex(902, (d_1099_v82_).length(0))
                lhs200_ = globalState
                r1 = rhs216_
                lhs195_[lhs196_] = rhs217_
                lhs197_.f4 = rhs218_
                lhs198_[lhs199_] = rhs219_
                lhs200_.f17 = rhs220_
        elif True:
            (globalState).f17 = default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1106_i17_ in range(default__.abs(302))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))), 409)
            d_1107_v88_: _dafny.Map
            d_1107_v88_ = _dafny.Map({(d_1004_v6_).f21: (self).f21})
            d_1108_v89_: _dafny.Map
            d_1108_v89_ = _dafny.Map({d_1107_v88_: (d_1004_v6_).f21})
            d_1109_v90_: _dafny.Map
            d_1109_v90_ = _dafny.Map({len(d_1052_v43_): (d_1004_v6_).f21})
            d_1110_v91_: _dafny.Map
            d_1110_v91_ = _dafny.Map({p0: d_1109_v90_})
            d_1111_v92_: _dafny.Map
            d_1111_v92_ = _dafny.Map({((d_1110_v91_)[p0] if (p0) in (d_1110_v91_) else d_1109_v90_): (self).f24})
            def iife82_():
                coll34_ = _dafny.Map()
                compr_34_: int
                for compr_34_ in (_dafny.MultiSet([len(d_1108_v89_)])).Elements:
                    d_1112_v87_: int = compr_34_
                    if (d_1112_v87_) in (_dafny.MultiSet([len(d_1108_v89_)])):
                        coll34_[(d_1112_v87_) * ((d_1004_v6_).f20)] = (d_984_v1_).cf16
                return _dafny.Map(coll34_)
            (globalState).f14 = (iife82_()
            ) in (d_1111_v92_)
            if p1:
                d_1113_v93_: _dafny.Array
                nw188_ = _dafny.Array(None, 5)
                nw188_[int(0)] = d_1054_v45_
                nw188_[int(1)] = d_1054_v45_
                nw188_[int(2)] = d_1054_v45_
                nw188_[int(3)] = (d_1054_v45_) - (d_1054_v45_)
                nw188_[int(4)] = _dafny.MultiSet([993])
                d_1113_v93_ = nw188_
                d_1114_v94_: _dafny.Seq
                d_1114_v94_ = _dafny.SeqWithoutIsStrInference([(d_1004_v6_).f20])
                index196_ = default__.safeIndex(812, (d_1113_v93_).length(0))
                (d_1113_v93_)[index196_] = (_dafny.MultiSet(d_1114_v94_)).intersection(d_1054_v45_)
                index197_ = default__.safeIndex(812, (d_1113_v93_).length(0))
                (d_1113_v93_)[index197_] = d_1054_v45_
                r1 = ((d_1052_v43_) - (d_1052_v43_)) == ((d_1052_v43_) - (d_1052_v43_))
                d_1115_v95_: str
                d_1115_v95_ = _dafny.CodePoint('u')
                d_1116_v96_: C1
                nw189_ = C1()
                nw189_.ctor__((955) + ((0) - (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1117_i18_ in range(default__.abs(605))])).set(default__.safeIndex(((d_1113_v93_)[default__.safeIndex(812, (d_1113_v93_).length(0))]).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1118_i18_ in range(default__.abs(605))]))), d_1115_v95_)))), (d_1003_v5_ if (self).f21 else d_1003_v5_), (d_1004_v6_).f20, (len((_dafny.SeqWithoutIsStrInference([d_1115_v95_ for d_1119_i19_ in range(default__.abs(420))])).set(default__.safeIndex((d_1004_v6_).f20, len(_dafny.SeqWithoutIsStrInference([d_1115_v95_ for d_1120_i19_ in range(default__.abs(420))]))), d_1115_v95_))) < ((d_1004_v6_).f20))
                d_1116_v96_ = nw189_
                d_1121_v97_: _dafny.Seq
                d_1121_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                d_1122_v98_: _dafny.Map
                d_1122_v98_ = _dafny.Map({453: (0) - (len(d_1121_v97_))})
                (globalState).f18 = (d_1004_v6_).fm0(d_1122_v98_, d_1121_v97_, (_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f24: (d_1004_v6_).f21})])).set(default__.safeIndex(len(d_1114_v94_), len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f24: (d_1004_v6_).f21})]))), d_1107_v88_), globalState)
                (globalState).f18 = (d_1116_v96_.f27) > ((d_1004_v6_).f20)
            elif True:
                d_1123_v99_: C1
                nw190_ = C1()
                nw190_.ctor__(p0, d_1003_v5_, p0, not(not(not (True) or ((self).f21))))
                d_1123_v99_ = nw190_
                d_1124_v100_: _dafny.Array
                nw191_ = _dafny.Array(False, 26)
                d_1124_v100_ = nw191_
                d_1124_v100_ = d_1124_v100_
                index198_ = default__.safeIndex(662, (d_1124_v100_).length(0))
                (d_1124_v100_)[index198_] = (self).f24
                index199_ = default__.safeIndex(662, (d_1124_v100_).length(0))
                (d_1124_v100_)[index199_] = p1
                index200_ = default__.safeIndex(355, (d_1003_v5_).length(0))
                (d_1003_v5_)[index200_] = 753
                d_1125_v101_: _dafny.Map
                d_1125_v101_ = _dafny.Map({(self).f24: d_1002_v4_})
                index201_ = default__.safeIndex(355, (d_1003_v5_).length(0))
                (d_1003_v5_)[index201_] = len((d_1125_v101_) | (d_1125_v101_))
                d_1053_v44_ = (d_1053_v44_) | ((d_1053_v44_) - (d_1053_v44_))
            (globalState).f5 = default__.safeDivisionInt((d_1004_v6_).f20, p0)
            (globalState).f14 = (self).f21
        d_1126_v102_: _dafny.MultiSet
        d_1126_v102_ = _dafny.MultiSet([p1])
        if (not((_dafny.MultiSet([(d_1004_v6_).f21])).isdisjoint(d_1126_v102_))) and ((self).f24):
            d_1127_v103_: _dafny.Seq
            d_1127_v103_ = _dafny.SeqWithoutIsStrInference([p0, p0, (self).f20])
            d_1128_v105_: _dafny.Map
            d_1128_v105_ = _dafny.Map({p0: (self).f20})
            d_1129_v106_: _dafny.Map
            d_1129_v106_ = _dafny.Map({(self).f21: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "keo"))})
            d_1130_v107_: _dafny.Map
            d_1130_v107_ = _dafny.Map({p0: True})
            d_1131_v108_: _dafny.Seq
            d_1131_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rupn"))
            d_1132_v109_: _dafny.Map
            d_1132_v109_ = _dafny.Map({(d_1004_v6_).f21: (d_1004_v6_).f21})
            def iife83_():
                coll35_ = _dafny.Map()
                compr_35_: int
                for compr_35_ in _dafny.IntegerRange(297, 305):
                    d_1133_v104_: int = compr_35_
                    if ((297) <= (d_1133_v104_)) and ((d_1133_v104_) < (305)):
                        coll35_[(d_1133_v104_) * ((d_1004_v6_).f20)] = (self).f20
                return _dafny.Map(coll35_)
            rhs221_ = (self).fm0((iife83_()
            ) | (d_1128_v105_), ((d_1129_v106_)[((d_1130_v107_)[(d_1004_v6_).f20] if ((d_1004_v6_).f20) in (d_1130_v107_) else p1)] if (((d_1130_v107_)[(d_1004_v6_).f20] if ((d_1004_v6_).f20) in (d_1130_v107_) else p1)) in (d_1129_v106_) else d_1131_v108_), _dafny.SeqWithoutIsStrInference([d_1132_v109_]), globalState)
            rhs222_ = (d_1127_v103_) + (d_1127_v103_)
            lhs201_ = globalState
            lhs201_.f4 = rhs221_
            d_1127_v103_ = rhs222_
            (globalState).f5 = default__.safeModuloInt(p0, (d_1004_v6_).f20)
            (globalState).f5 = (0) - ((d_1004_v6_).f20)
            if False:
                d_1134_v110_: _dafny.Map
                d_1134_v110_ = _dafny.Map({(d_1004_v6_).f21: default__.fm17((self).f21, globalState)})
                index202_ = default__.safeIndex(642, (d_1003_v5_).length(0))
                (d_1003_v5_)[index202_] = len((d_1134_v110_) | (d_1134_v110_))
                index203_ = default__.safeIndex(642, (d_1003_v5_).length(0))
                rhs223_ = d_1126_v102_
                rhs224_ = (0) - ((self).f20)
                rhs225_ = (d_1126_v102_).cardinality
                rhs226_ = default__.safeModuloInt(((d_1004_v6_).f20) * ((self).f20), (self).f20)
                lhs202_ = globalState
                lhs203_ = d_1003_v5_
                lhs204_ = default__.safeIndex(642, (d_1003_v5_).length(0))
                lhs205_ = globalState
                d_1126_v102_ = rhs223_
                lhs202_.f17 = rhs224_
                lhs203_[lhs204_] = rhs225_
                lhs205_.f5 = rhs226_
                d_1135_v111_: D10
                d_1135_v111_ = D10_DC25(d_1054_v45_)
                d_1136_v112_: _dafny.Seq
                d_1136_v112_ = _dafny.SeqWithoutIsStrInference([(d_1135_v111_).cf34])
                d_1137_v113_: _dafny.Array
                nw192_ = _dafny.Array(False, 6)
                d_1137_v113_ = nw192_
                d_1138_v114_: _dafny.Seq
                d_1138_v114_ = _dafny.SeqWithoutIsStrInference([d_1137_v113_])
                d_1139_v115_: _dafny.Map
                d_1139_v115_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_1054_v45_ for d_1140_i20_ in range(default__.abs(776))])) + (d_1136_v112_): (d_1138_v114_) + (d_1138_v114_)})
                d_1141_v116_: D3
                d_1141_v116_ = D3_DC5(d_1138_v114_)
                d_1139_v115_ = (d_1139_v115_).set(_dafny.SeqWithoutIsStrInference([d_1054_v45_]), (d_1141_v116_).cf9)
                d_1142_v117_: _dafny.Seq
                d_1142_v117_ = _dafny.SeqWithoutIsStrInference([d_1002_v4_, (d_1002_v4_).set(default__.safeIndex(len(d_1132_v109_), len(d_1002_v4_)), (d_1004_v6_).f21), d_1002_v4_, d_1002_v4_, d_1002_v4_])
                d_1143_v118_: _dafny.Map
                d_1143_v118_ = _dafny.Map({d_1142_v117_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbrjrlv"))})
                d_1144_v119_: str
                d_1144_v119_ = _dafny.CodePoint('t')
                rhs227_ = (self).f21
                rhs228_ = (d_1004_v6_).f21
                rhs229_ = ((d_1131_v108_) + (((d_1143_v118_)[d_1142_v117_] if (d_1142_v117_) in (d_1143_v118_) else d_1131_v108_))).set(default__.safeIndex((self).f20, len((d_1131_v108_) + (((d_1143_v118_)[d_1142_v117_] if (d_1142_v117_) in (d_1143_v118_) else d_1131_v108_)))), d_1144_v119_)
                lhs206_ = globalState
                lhs207_ = globalState
                lhs208_ = globalState
                lhs206_.f14 = rhs227_
                lhs207_.f14 = rhs228_
                lhs208_.f15 = rhs229_
                index204_ = default__.safeIndex(642, (d_1003_v5_).length(0))
                (d_1003_v5_)[index204_] = (self).fm14(True, (d_1004_v6_).f20, (d_1004_v6_).f20, (d_1131_v108_) != (default__.fm26((self).f21, globalState)), globalState)
                d_1132_v109_ = (d_1132_v109_).set(True, ((d_1130_v107_)[len(d_1132_v109_)] if (len(d_1132_v109_)) in (d_1130_v107_) else (d_1004_v6_).f21))
            elif True:
                d_1145_v120_: _dafny.Map
                d_1145_v120_ = _dafny.Map({(self).f24: 93})
                (globalState).f5 = (0) - ((((d_1145_v120_)[(d_1004_v6_).f21] if ((d_1004_v6_).f21) in (d_1145_v120_) else (d_1004_v6_).f20)) - ((d_1004_v6_).f20))
                d_1146_v121_: _dafny.Array
                nw193_ = _dafny.Array(False, 21)
                d_1146_v121_ = nw193_
                index205_ = default__.safeIndex(212, (d_1146_v121_).length(0))
                (d_1146_v121_)[index205_] = (self).f21
                index206_ = default__.safeIndex(212, (d_1146_v121_).length(0))
                (d_1146_v121_)[index206_] = (d_1002_v4_)[default__.safeIndex((d_1004_v6_).f20, len(d_1002_v4_))]
                index207_ = default__.safeIndex(212, (d_1146_v121_).length(0))
                (d_1146_v121_)[index207_] = (self).fm13(p0, globalState)
                (globalState).f5 = (len((d_1128_v105_) | (d_1128_v105_))) + ((d_1004_v6_).f20)
                (globalState).f17 = (self).f20
            d_1147_v122_: _dafny.Array
            nw194_ = _dafny.Array(False, 22)
            d_1147_v122_ = nw194_
            d_1147_v122_ = d_1147_v122_
        elif True:
            d_1003_v5_ = (d_1003_v5_ if (self).f21 else d_1003_v5_)
            d_1148_v123_: _dafny.Seq
            d_1148_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_1149_v124_: _dafny.Map
            d_1149_v124_ = _dafny.Map({len(d_1148_v123_): len(d_1002_v4_)})
            d_1150_v125_: _dafny.Map
            d_1150_v125_ = _dafny.Map({(d_1004_v6_).f20: ((d_1149_v124_)[(self).f20] if ((self).f20) in (d_1149_v124_) else (self).f20)})
            (globalState).f4 = (d_1004_v6_).fm0(d_1150_v125_, d_1148_v123_, _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1004_v6_).f21: (d_1004_v6_).f21}) for d_1151_i21_ in range(default__.abs(-561))]), globalState)
            d_1152_v126_: _dafny.Array
            d_1153_v127_: int
            out16_: _dafny.Array
            out17_: int
            out16_, out17_ = (self).m1(len(d_1148_v123_), (self).f24, d_1004_v6_, globalState)
            d_1152_v126_ = out16_
            d_1153_v127_ = out17_
            d_1154_v128_: _dafny.Map
            d_1154_v128_ = _dafny.Map({p0: (self).f21})
            d_1155_v129_: _dafny.Seq
            d_1155_v129_ = _dafny.SeqWithoutIsStrInference([(d_1004_v6_).f20])
            d_1156_v130_: str
            d_1156_v130_ = _dafny.CodePoint('n')
            d_1157_v131_: _dafny.Map
            d_1157_v131_ = _dafny.Map({(d_1004_v6_).f21: (d_1004_v6_).f21})
            d_1158_v132_: _dafny.Seq
            d_1158_v132_ = _dafny.SeqWithoutIsStrInference([d_1157_v131_])
            d_1159_v133_: _dafny.Map
            d_1159_v133_ = _dafny.Map({False: (d_1004_v6_).fm0(d_1149_v124_, d_1148_v123_, d_1158_v132_, globalState)})
            d_1154_v128_ = (d_1154_v128_).set(len((d_1155_v129_) + ((d_1155_v129_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1156_v130_])), len(d_1155_v129_)), (d_1054_v45_).cardinality))), ((d_1159_v133_)[(d_1004_v6_).f21] if ((d_1004_v6_).f21) in (d_1159_v133_) else (self).f21))
            d_1160_v134_: _dafny.Seq
            d_1160_v134_ = _dafny.SeqWithoutIsStrInference([d_1152_v126_, d_1152_v126_, d_1152_v126_, d_1003_v5_, d_1152_v126_])
            r0 = ((d_1160_v134_) + (d_1160_v134_)) + ((_dafny.SeqWithoutIsStrInference([d_1152_v126_, d_1152_v126_]) if (self).f21 else d_1160_v134_))
        d_1161_v135_: _dafny.Seq
        d_1161_v135_ = _dafny.SeqWithoutIsStrInference([d_1003_v5_])
        r0 = (d_1161_v135_) + (d_1161_v135_)
        d_1162_v138_: _dafny.Map
        d_1162_v138_ = _dafny.Map({(self).f24: (self).f21})
        d_1163_v139_: _dafny.Map
        d_1163_v139_ = _dafny.Map({p1: p0})
        d_1164_v140_: _dafny.Seq
        d_1164_v140_ = _dafny.SeqWithoutIsStrInference([d_1162_v138_, default__.fm15(522, (d_1004_v6_).f20, ((d_1163_v139_)[(self).f24] if ((self).f24) in (d_1163_v139_) else p0), (d_1004_v6_).f20, globalState), d_1162_v138_, _dafny.Map({(d_1004_v6_).f21: (self).f21}), _dafny.Map({(self).f24: (self).f24})])
        def iife84_():
            def iife85_():
                coll37_ = _dafny.Map()
                compr_37_: int
                for compr_37_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({_dafny.CodePoint('t'): p1})))])).Elements:
                    d_1166_v137_: int = compr_37_
                    if (d_1166_v137_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({_dafny.CodePoint('t'): p1})))])):
                        coll37_[(d_1166_v137_) - ((d_1004_v6_).f20)] = len(d_1053_v44_)
                return _dafny.Map(coll37_)
            coll36_ = _dafny.Map()
            compr_36_: int
            for compr_36_ in _dafny.IntegerRange(702, 756):
                d_1165_v136_: int = compr_36_
                if ((702) <= (d_1165_v136_)) and ((d_1165_v136_) < (756)):
                    coll36_[(d_1165_v136_) * (len(iife85_()
                    ))] = (0) - ((d_1004_v6_).f20)
            return _dafny.Map(coll36_)
        r1 = (self).fm0(iife84_()
        , _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1167_i22_ in range(default__.abs(172))]), d_1164_v140_, globalState)
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        (globalState).f17 = p0
        if (self).f24:
            d_1168_v0_: _dafny.Map
            d_1168_v0_ = _dafny.Map({(self).f20: (p2).f20})
            d_1169_v1_: _dafny.Map
            d_1169_v1_ = _dafny.Map({(p2).f20: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyiwlx"))})
            d_1170_v2_: _dafny.Map
            d_1170_v2_ = _dafny.Map({True: (self).f21})
            d_1171_v3_: _dafny.Seq
            d_1171_v3_ = _dafny.SeqWithoutIsStrInference([d_1170_v2_, d_1170_v2_, _dafny.Map({(self).f21: (D7_DC17((self).f24)).cf25})])
            d_1172_v4_: C3
            nw195_ = C3()
            nw195_.ctor__((0) - (len(_dafny.Map({(self).fm0(d_1168_v0_, ((d_1169_v1_)[(p2).f20] if ((p2).f20) in (d_1169_v1_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1173_i0_ in range(default__.abs(-688))])), d_1171_v3_, globalState): True}))), (self).f21, p0, (self).f21)
            d_1172_v4_ = nw195_
            nw196_ = C3()
            nw196_.ctor__(365, (self).f24, 599, (p2).f21)
            d_1172_v4_ = nw196_
            if not((d_1172_v4_).f26):
                d_1174_v5_: _dafny.Array
                nw197_ = _dafny.Array(False, 6)
                d_1174_v5_ = nw197_
                d_1175_v6_: _dafny.Seq
                d_1175_v6_ = _dafny.SeqWithoutIsStrInference([d_1174_v5_])
                d_1176_v7_: D3
                d_1176_v7_ = D3_DC5(d_1175_v6_)
                pat_let_tv23_ = d_1175_v6_
                d_1177_v8_: _dafny.Array
                nw198_ = _dafny.Array(None, 15)
                nw198_[int(0)] = d_1176_v7_
                nw198_[int(1)] = d_1176_v7_
                nw198_[int(2)] = d_1176_v7_
                nw198_[int(3)] = d_1176_v7_
                nw198_[int(4)] = d_1176_v7_
                nw198_[int(5)] = d_1176_v7_
                nw198_[int(6)] = d_1176_v7_
                nw198_[int(7)] = d_1176_v7_
                def iife86_(_pat_let24_0):
                    def iife87_(d_1178_dt__update__tmp_h0_):
                        def iife88_(_pat_let25_0):
                            def iife89_(d_1179_dt__update_hcf9_h0_):
                                return D3_DC5(d_1179_dt__update_hcf9_h0_)
                            return iife89_(_pat_let25_0)
                        return iife88_(pat_let_tv23_)
                    return iife87_(_pat_let24_0)
                nw198_[int(8)] = iife86_(d_1176_v7_)
                nw198_[int(9)] = d_1176_v7_
                nw198_[int(10)] = D3_DC5(d_1175_v6_)
                nw198_[int(11)] = D3_DC5(d_1175_v6_)
                nw198_[int(12)] = d_1176_v7_
                nw198_[int(13)] = d_1176_v7_
                nw198_[int(14)] = d_1176_v7_
                d_1177_v8_ = nw198_
                d_1180_v9_: _dafny.Array
                nw199_ = _dafny.Array(None, 23)
                nw199_[int(0)] = d_1177_v8_
                nw199_[int(1)] = d_1177_v8_
                nw199_[int(2)] = d_1177_v8_
                nw199_[int(3)] = d_1177_v8_
                nw199_[int(4)] = d_1177_v8_
                nw199_[int(5)] = d_1177_v8_
                nw199_[int(6)] = (D14_DC38(d_1177_v8_)).cf51
                nw199_[int(7)] = d_1177_v8_
                nw199_[int(8)] = d_1177_v8_
                nw199_[int(9)] = d_1177_v8_
                nw199_[int(10)] = d_1177_v8_
                nw199_[int(11)] = d_1177_v8_
                nw199_[int(12)] = d_1177_v8_
                nw199_[int(13)] = d_1177_v8_
                nw199_[int(14)] = d_1177_v8_
                nw199_[int(15)] = d_1177_v8_
                nw199_[int(16)] = d_1177_v8_
                nw199_[int(17)] = d_1177_v8_
                nw199_[int(18)] = d_1177_v8_
                nw199_[int(19)] = d_1177_v8_
                nw199_[int(20)] = d_1177_v8_
                nw199_[int(21)] = d_1177_v8_
                nw199_[int(22)] = d_1177_v8_
                d_1180_v9_ = nw199_
                index208_ = default__.safeIndex(497, (d_1180_v9_).length(0))
                (d_1180_v9_)[index208_] = d_1177_v8_
                d_1181_v10_: _dafny.Seq
                d_1181_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpcrmxv"))
                d_1182_v11_: _dafny.Set
                d_1182_v11_ = _dafny.Set({(self).f24, not(not((d_1172_v4_).f26))})
                index209_ = default__.safeIndex(497, (d_1180_v9_).length(0))
                rhs230_ = default__.safeModuloInt(((d_1172_v4_).f25) * ((d_1172_v4_).f25), (d_1172_v4_).fm1(len(d_1181_v10_), len(_dafny.Map({(self).f20: p1})), False, globalState))
                rhs231_ = ((self).f20) <= ((len(d_1182_v11_)) * ((p2).f20))
                rhs232_ = (len(d_1170_v2_)) - (len((d_1181_v10_) + (d_1181_v10_)))
                rhs233_ = d_1177_v8_
                lhs209_ = globalState
                lhs210_ = globalState
                lhs211_ = globalState
                lhs212_ = d_1180_v9_
                lhs213_ = default__.safeIndex(497, (d_1180_v9_).length(0))
                lhs209_.f17 = rhs230_
                lhs210_.f4 = rhs231_
                lhs211_.f5 = rhs232_
                lhs212_[lhs213_] = rhs233_
                d_1183_v12_: _dafny.Seq
                d_1183_v12_ = _dafny.SeqWithoutIsStrInference([(p2).f21, (p2).f21])
                d_1184_v13_: _dafny.Seq
                d_1184_v13_ = _dafny.SeqWithoutIsStrInference([(d_1183_v12_)[default__.safeIndex((d_1172_v4_).f25, len(d_1183_v12_))]])
                d_1185_v14_: _dafny.Seq
                d_1185_v14_ = _dafny.SeqWithoutIsStrInference([d_1183_v12_, _dafny.SeqWithoutIsStrInference([(p2).f21])])
                d_1186_v15_: C2
                nw200_ = C2()
                nw200_.ctor__(len(d_1185_v14_), not((self).fm13((p2).f20, globalState)))
                d_1186_v15_ = nw200_
                d_1187_v18_: _dafny.Map
                d_1187_v18_ = _dafny.Map({p0: (self).fm0(d_1168_v0_, d_1181_v10_, _dafny.SeqWithoutIsStrInference([(d_1171_v3_)[default__.safeIndex((d_1172_v4_).f25, len(d_1171_v3_))] for d_1188_i1_ in range(default__.abs(167))]), globalState)})
                d_1189_v19_: _dafny.Set
                d_1189_v19_ = _dafny.Set({_dafny.Map({(p2).f20: (self).f24}), d_1187_v18_, d_1187_v18_})
                d_1190_v20_: str
                d_1190_v20_ = _dafny.CodePoint('w')
                d_1191_v21_: _dafny.Map
                def iife90_():
                    def iife92_():
                        coll40_ = _dafny.Map()
                        compr_40_: _dafny.Map
                        for compr_40_ in (d_1189_v19_).Elements:
                            d_1192_v17_: _dafny.Map = compr_40_
                            if (d_1192_v17_) in (d_1189_v19_):
                                coll40_[d_1192_v17_] = (self).f24
                        return _dafny.Map(coll40_)
                    coll38_ = _dafny.Map()
                    def iife91_():
                        coll39_ = _dafny.Map()
                        compr_39_: _dafny.Map
                        for compr_39_ in (d_1189_v19_).Elements:
                            d_1192_v17_: _dafny.Map = compr_39_
                            if (d_1192_v17_) in (d_1189_v19_):
                                coll39_[d_1192_v17_] = (self).f24
                        return _dafny.Map(coll39_)
                    compr_38_: _dafny.Map
                    for compr_38_ in (iife91_()
                    ).keys.Elements:
                        d_1193_v16_: _dafny.Map = compr_38_
                        if (d_1193_v16_) in (iife92_()
                        ):
                            coll38_[d_1193_v16_] = p1
                    return _dafny.Map(coll38_)
                d_1191_v21_ = _dafny.Map({iife90_()
                : d_1190_v20_})
                d_1194_v22_: bool
                d_1194_v22_ = (self).f24
                d_1195_v23_: _dafny.Map
                d_1195_v23_ = _dafny.Map({d_1187_v18_: d_1194_v22_})
                d_1196_v24_: _dafny.Map
                d_1196_v24_ = _dafny.Map({(self).f24: d_1190_v20_})
                d_1191_v21_ = (d_1191_v21_).set(d_1195_v23_, ((d_1196_v24_)[(d_1172_v4_).f26] if ((d_1172_v4_).f26) in (d_1196_v24_) else d_1190_v20_))
                d_1197_v26_: _dafny.Seq
                d_1197_v26_ = _dafny.SeqWithoutIsStrInference([(d_1172_v4_).f25])
                pat_let_tv24_ = d_1184_v13_
                pat_let_tv25_ = d_1184_v13_
                d_1198_v27_: D8
                def iife93_():
                    def iife98_(_pat_let28_0):
                        def iife99_(d_1202_dt__update__tmp_h1_):
                            def iife100_(_pat_let29_0):
                                def iife101_(d_1203_dt__update_hcf0_h0_):
                                    return d_1203_dt__update_hcf0_h0_
                                return iife101_(_pat_let29_0)
                            return iife100_(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(pat_let_tv25_)).cardinality]))
                        return iife99_(_pat_let28_0)
                    coll41_ = _dafny.Map()
                    def iife94_(_pat_let26_0):
                        def iife95_(d_1199_dt__update__tmp_h1_):
                            def iife96_(_pat_let27_0):
                                def iife97_(d_1200_dt__update_hcf0_h0_):
                                    return d_1200_dt__update_hcf0_h0_
                                return iife97_(_pat_let27_0)
                            return iife96_(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(pat_let_tv24_)).cardinality]))
                        return iife95_(_pat_let26_0)
                    compr_41_: int
                    for compr_41_ in ((iife94_(d_1197_v26_))).Elements:
                        d_1201_v25_: int = compr_41_
                        if (d_1201_v25_) in ((iife98_(d_1197_v26_))):
                            coll41_[(d_1201_v25_) * (p0)] = (p2).f20
                    return _dafny.Map(coll41_)
                d_1198_v27_ = D8_DC19(iife93_()
)
                d_1204_v28_: _dafny.Seq
                d_1204_v28_ = _dafny.SeqWithoutIsStrInference([(d_1198_v27_).cf28, (d_1168_v0_).set((d_1172_v4_).f25, p0)])
                (globalState).f18 = (default__.safeModuloInt((p2).f20, len((d_1204_v28_)[default__.safeIndex((p2).f20, len(d_1204_v28_))]))) != ((p2).f20)
                d_1205_v29_: D13
                d_1205_v29_ = D13_DC37(p1, (self).f20)
                d_1206_v30_: D4
                d_1206_v30_ = D4_DC11((p2).f20, (d_1205_v29_).cf49, (d_1172_v4_).f26)
                d_1207_v31_: _dafny.Map
                d_1207_v31_ = _dafny.Map({(p2).f21: d_1206_v30_})
                (globalState).f4 = (d_1183_v12_)[default__.safeIndex(len((d_1207_v31_).set((d_1172_v4_).f26, d_1206_v30_)), len(d_1183_v12_))]
            elif True:
                (globalState).f14 = (p2).f21
                d_1208_v32_: _dafny.Array
                def lambda57_(d_1209_v4_):
                    def lambda58_(d_1210_i2_):
                        return (d_1210_i2_) * ((d_1209_v4_).f25)

                    return lambda58_

                init34_ = lambda57_(d_1172_v4_)
                nw201_ = _dafny.Array(None, 29)
                for i0_34_ in range(nw201_.length(0)):
                    nw201_[i0_34_] = init34_(i0_34_)
                d_1208_v32_ = nw201_
                d_1211_v33_: _dafny.Set
                d_1211_v33_ = _dafny.Set({d_1208_v32_})
                d_1212_v34_: C1
                nw202_ = C1()
                nw202_.ctor__(675, d_1208_v32_, (p2).f20, (d_1172_v4_).f26)
                d_1212_v34_ = nw202_
                d_1213_v35_: _dafny.Array
                nw203_ = _dafny.Array(False, 4)
                d_1213_v35_ = nw203_
                d_1214_v36_: _dafny.Seq
                d_1214_v36_ = _dafny.SeqWithoutIsStrInference([d_1213_v35_, d_1213_v35_])
                d_1215_v37_: D3
                d_1215_v37_ = D3_DC5(d_1214_v36_)
                d_1216_v38_: D3
                d_1216_v38_ = D3_DC8(d_1215_v37_)
                d_1217_v39_: _dafny.Map
                d_1217_v39_ = _dafny.Map({d_1212_v34_: d_1216_v38_})
                d_1218_v40_: _dafny.Set
                d_1218_v40_ = _dafny.Set({d_1212_v34_.f27, -407, len(_dafny.SeqWithoutIsStrInference([p1]))})
                d_1219_v41_: _dafny.Seq
                d_1219_v41_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                d_1220_v42_: _dafny.Map
                d_1220_v42_ = _dafny.Map({d_1217_v39_: (d_1218_v40_) == (_dafny.Set({(self).f20, len(d_1219_v41_), (d_1212_v34_).fm1((p2).f20, -317, p1, globalState)}))})
                d_1221_v43_: _dafny.Seq
                d_1221_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klgw"))
                rhs234_ = d_1211_v33_
                rhs235_ = d_1220_v42_
                rhs236_ = (d_1212_v34_).fm0(d_1168_v0_, d_1221_v43_, d_1171_v3_, globalState)
                lhs214_ = globalState
                d_1211_v33_ = rhs234_
                d_1220_v42_ = rhs235_
                lhs214_.f4 = rhs236_
                d_1222_v44_: _dafny.Map
                d_1222_v44_ = _dafny.Map({(p2).f20: d_1168_v0_})
                (globalState).f4 = (p2).fm0(((d_1222_v44_)[(p2).f20] if ((p2).f20) in (d_1222_v44_) else d_1168_v0_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtqfdj")), d_1171_v3_, globalState)
                d_1223_v45_: _dafny.MultiSet
                d_1223_v45_ = _dafny.MultiSet([not((self).f24)])
                d_1223_v45_ = _dafny.MultiSet(d_1219_v41_)
                (globalState).f4 = ((default__.fm37((p2).f20, globalState)).ispropersubset(_dafny.MultiSet([(self).f20])) if (d_1172_v4_).f26 else (p2).f21)
            d_1224_v46_: _dafny.Seq
            d_1224_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyhopkq"))
            d_1225_v47_: D8
            d_1225_v47_ = D8_DC19(d_1168_v0_)
            d_1226_v48_: _dafny.Map
            d_1226_v48_ = _dafny.Map({d_1224_v46_: True})
            d_1227_v49_: _dafny.Set
            d_1227_v49_ = _dafny.Set({not((p2).f21)})
            (globalState).f4 = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwcov"))) + (d_1224_v46_))) > (len(((d_1225_v47_).cf28).set(len(d_1226_v48_), len(d_1227_v49_))))
            (globalState).f15 = ((d_1224_v46_ if (self).fm13((0) - ((p2).f20), globalState) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1228_i3_ in range(default__.abs(514))]))) + (d_1224_v46_)
            d_1229_v50_: _dafny.Array
            nw204_ = _dafny.Array(int(0), 18)
            d_1229_v50_ = nw204_
            d_1230_v51_: T0
            nw205_ = C1()
            nw205_.ctor__((self).f20, (d_1229_v50_ if (self).f24 else d_1229_v50_), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndkax")))) * (447), (p2).f21)
            d_1230_v51_ = nw205_
            d_1230_v51_ = d_1230_v51_
        elif True:
            (globalState).f17 = (0) - ((p2).f20)
            d_1231_v52_: _dafny.Seq
            d_1231_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucx"))
            d_1232_v53_: _dafny.Seq
            d_1232_v53_ = _dafny.SeqWithoutIsStrInference([(self).f20])
            d_1233_v54_: _dafny.Map
            d_1233_v54_ = _dafny.Map({(0) - (p0): len(d_1232_v53_)})
            d_1234_v55_: _dafny.Map
            d_1234_v55_ = _dafny.Map({(self).f21: p1})
            d_1235_v56_: _dafny.Seq
            d_1235_v56_ = _dafny.SeqWithoutIsStrInference([d_1234_v55_, d_1234_v55_, d_1234_v55_, d_1234_v55_])
            d_1236_v57_: _dafny.Seq
            d_1236_v57_ = _dafny.SeqWithoutIsStrInference([(self).f24, not((self).fm0(d_1233_v54_, d_1231_v52_, d_1235_v56_, globalState))])
            d_1237_v58_: _dafny.Array
            nw206_ = _dafny.Array(None, 8)
            nw206_[int(0)] = -892
            nw206_[int(1)] = (_dafny.MultiSet([p1, (self).f21, (self).f24, not(p1), (p2).f21])).cardinality
            nw206_[int(2)] = len(d_1231_v52_)
            nw206_[int(3)] = (p2).f20
            nw206_[int(4)] = 120
            nw206_[int(5)] = len(d_1236_v57_)
            nw206_[int(6)] = p0
            nw206_[int(7)] = (self).f20
            d_1237_v58_ = nw206_
            d_1238_v59_: C1
            nw207_ = C1()
            nw207_.ctor__((p2).f20, d_1237_v58_, (0) - ((p2).f20), (p2).f21)
            d_1238_v59_ = nw207_
            d_1238_v59_ = d_1238_v59_
            d_1239_v60_: _dafny.Seq
            d_1239_v60_ = _dafny.SeqWithoutIsStrInference([d_1236_v57_])
            d_1236_v57_ = ((d_1239_v60_)[default__.safeIndex((self).f20, len(d_1239_v60_))]) + (_dafny.SeqWithoutIsStrInference([(p2).f21, (p2).f21]))
            d_1240_v61_: _dafny.Array
            nw208_ = _dafny.Array(None, 21)
            d_1240_v61_ = nw208_
            index210_ = default__.safeIndex(108, (d_1240_v61_).length(0))
            (d_1240_v61_)[index210_] = d_1238_v59_
            d_1241_v62_: C1
            d_1241_v62_ = d_1238_v59_
            index211_ = default__.safeIndex(108, (d_1240_v61_).length(0))
            (d_1240_v61_)[index211_] = (d_1241_v62_)
            (globalState).f18 = (p2).f21
        d_1242_v63_: _dafny.Map
        d_1242_v63_ = _dafny.Map({p0: (self).f24})
        d_1243_i4_: int
        d_1243_i4_ = 0
        with _dafny.label("4"):
            while (((d_1242_v63_)[-116] if (-116) in (d_1242_v63_) else (self).f24)) and ((p2).f21):
                with _dafny.c_label("4"):
                    if (d_1243_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_1243_i4_ = (d_1243_i4_) + (1)
                    d_1244_v64_: _dafny.Seq
                    d_1244_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqwrp"))
                    d_1245_v65_: _dafny.Seq
                    d_1245_v65_ = _dafny.SeqWithoutIsStrInference([(p2).f20])
                    d_1246_v66_: str
                    d_1246_v66_ = _dafny.CodePoint('x')
                    (globalState).f3 = (((d_1244_v64_).set(default__.safeIndex(len(d_1245_v65_), len(d_1244_v64_)), d_1246_v66_)) + (d_1244_v64_)) <= ((d_1244_v64_) + ((d_1244_v64_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1246_v66_, d_1246_v66_, d_1246_v66_, d_1246_v66_, _dafny.CodePoint('f')])), len(d_1244_v64_)), d_1246_v66_)))
                    d_1247_v67_: _dafny.Array
                    def lambda59_(d_1248_i5_):
                        return (d_1248_i5_) - ((self).f20)

                    init35_ = lambda59_
                    nw209_ = _dafny.Array(None, 17)
                    for i0_35_ in range(nw209_.length(0)):
                        nw209_[i0_35_] = init35_(i0_35_)
                    d_1247_v67_ = nw209_
                    d_1249_v68_: C1
                    nw210_ = C1()
                    nw210_.ctor__((self).f20, d_1247_v67_, 284, not((self).f21))
                    d_1249_v68_ = nw210_
                    if (p2).f21:
                        (globalState).f17 = (p2).f20
                        d_1250_v70_: _dafny.Map
                        d_1250_v70_ = _dafny.Map({(p2).f21: (p2).f21})
                        d_1251_v71_: D2
                        def iife102_():
                            coll42_ = _dafny.Map()
                            compr_42_: int
                            for compr_42_ in _dafny.IntegerRange(260, 408):
                                d_1252_v69_: int = compr_42_
                                if ((260) <= (d_1252_v69_)) and ((d_1252_v69_) < (408)):
                                    coll42_[default__.safeModuloInt(d_1252_v69_, len(d_1244_v64_))] = d_1249_v68_.f27
                            return _dafny.Map(coll42_)
                        d_1251_v71_ = D2_DC3((p2).f20, not((self).f24), p1, (self).f20, (d_1249_v68_).fm0(iife102_()
, d_1244_v64_, _dafny.SeqWithoutIsStrInference([d_1250_v70_, d_1250_v70_]), globalState))
                        d_1253_v72_: _dafny.Seq
                        d_1253_v72_ = _dafny.SeqWithoutIsStrInference([not((d_1251_v71_).cf7)])
                        d_1254_v73_: C0
                        nw211_ = C0()
                        nw211_.ctor__()
                        d_1254_v73_ = nw211_
                        d_1255_v74_: _dafny.Map
                        d_1255_v74_ = _dafny.Map({d_1254_v73_: p2})
                        d_1256_v75_: _dafny.Set
                        d_1256_v75_ = _dafny.Set({((d_1255_v74_)[d_1254_v73_] if (d_1254_v73_) in (d_1255_v74_) else p2)})
                        d_1257_v76_: _dafny.Array
                        nw212_ = _dafny.Array(None, 25)
                        nw212_[int(0)] = (p2).f21
                        nw212_[int(1)] = (p2).f21
                        nw212_[int(2)] = p1
                        nw212_[int(3)] = p1
                        nw212_[int(4)] = (p2).f21
                        nw212_[int(5)] = (d_1244_v64_) <= (default__.fm26((p2).f21, globalState))
                        nw212_[int(6)] = not(not(not((p2).f21)))
                        nw212_[int(7)] = (d_1242_v63_) == (d_1242_v63_)
                        nw212_[int(8)] = (p2).f21
                        nw212_[int(9)] = (self).f24
                        nw212_[int(10)] = (self).f21
                        nw212_[int(11)] = p1
                        nw212_[int(12)] = (p0) >= ((p2).f20)
                        nw212_[int(13)] = (p2).f21
                        nw212_[int(14)] = (p1) or ((self).f24)
                        nw212_[int(15)] = (d_1253_v72_)[default__.safeIndex(769, len(d_1253_v72_))]
                        nw212_[int(16)] = (p2).f21
                        nw212_[int(17)] = not((d_1256_v75_).ispropersubset(d_1256_v75_))
                        nw212_[int(18)] = (self).f24
                        nw212_[int(19)] = (p2).f21
                        nw212_[int(20)] = (d_1249_v68_.f27) <= (-465)
                        nw212_[int(21)] = not((p2).f21)
                        nw212_[int(22)] = (not((p2).f21) if (p2).f21 else (self).f24)
                        nw212_[int(23)] = True
                        nw212_[int(24)] = ((self).f21) == ((p2).f21)
                        d_1257_v76_ = nw212_
                        index212_ = default__.safeIndex(452, (d_1257_v76_).length(0))
                        (d_1257_v76_)[index212_] = (p2).f21
                        index213_ = default__.safeIndex(452, (d_1257_v76_).length(0))
                        (d_1257_v76_)[index213_] = p1
                        index214_ = default__.safeIndex(261, ((d_1249_v68_).f28).length(0))
                        ((d_1249_v68_).f28)[index214_] = p0
                        index215_ = default__.safeIndex(261, ((d_1249_v68_).f28).length(0))
                        rhs237_ = (0) - (d_1249_v68_.f27)
                        rhs238_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfdxlwohg"))
                        lhs215_ = (d_1249_v68_).f28
                        lhs216_ = default__.safeIndex(261, ((d_1249_v68_).f28).length(0))
                        lhs215_[lhs216_] = rhs237_
                        d_1244_v64_ = rhs238_
                        index216_ = default__.safeIndex(452, (d_1257_v76_).length(0))
                        (d_1257_v76_)[index216_] = (p2).f21
                        (globalState).f5 = ((p2).f20) + (86)
                    elif True:
                        d_1258_v77_: _dafny.Array
                        nw213_ = _dafny.Array(False, 11)
                        d_1258_v77_ = nw213_
                        d_1259_v78_: _dafny.Map
                        d_1259_v78_ = _dafny.Map({(p2).f21: d_1258_v77_})
                        d_1260_v79_: _dafny.Map
                        d_1260_v79_ = _dafny.Map({len(d_1244_v64_): len(d_1259_v78_)})
                        d_1261_v80_: _dafny.MultiSet
                        d_1261_v80_ = _dafny.MultiSet([(self).f24, (self).f24])
                        d_1262_v81_: _dafny.Seq
                        d_1262_v81_ = d_1245_v65_
                        d_1263_v82_: _dafny.Map
                        d_1263_v82_ = _dafny.Map({(p2).f21: (d_1249_v68_).fm16(d_1261_v80_, d_1262_v81_, globalState)})
                        d_1264_v83_: _dafny.Seq
                        d_1264_v83_ = _dafny.SeqWithoutIsStrInference([d_1263_v82_, d_1263_v82_, d_1263_v82_, d_1263_v82_])
                        d_1265_v84_: _dafny.Map
                        d_1265_v84_ = _dafny.Map({(p2).f21: d_1244_v64_})
                        (globalState).f15 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyibl")) if not ((d_1249_v68_).fm0(d_1260_v79_, (d_1244_v64_).set(default__.safeIndex(p0, len(d_1244_v64_)), d_1246_v66_), d_1264_v83_, globalState)) or ((p2).f21) else ((d_1265_v84_)[(p2).f21] if ((p2).f21) in (d_1265_v84_) else d_1244_v64_))
                        d_1266_v85_: C0
                        nw214_ = C0()
                        nw214_.ctor__()
                        d_1266_v85_ = nw214_
                        (globalState).f18 = (d_1245_v65_) == ((d_1245_v65_) + (_dafny.SeqWithoutIsStrInference([(self).f20 for d_1267_i6_ in range(default__.abs(169))])))
                        d_1268_v86_: _dafny.MultiSet
                        d_1268_v86_ = _dafny.MultiSet([(p2).fm1((p2).f20, len(default__.fm28((p2).f20, globalState)), (D9_DC23(p0, (p2).f21)).cf32, globalState), (0) - ((p2).fm1(len(d_1242_v63_), (0) - ((p2).f20), p1, globalState)), d_1249_v68_.f27, 855])
                        d_1269_v87_: D10
                        d_1269_v87_ = D10_DC25(d_1268_v86_)
                        d_1270_v88_: _dafny.Array
                        nw215_ = _dafny.Array(_dafny.Array(None, 0), 15)
                        d_1270_v88_ = nw215_
                        index217_ = default__.safeIndex(825, (d_1270_v88_).length(0))
                        (d_1270_v88_)[index217_] = (d_1258_v77_ if (self).f21 else d_1258_v77_)
                        index218_ = default__.safeIndex(6, (d_1258_v77_).length(0))
                        (d_1258_v77_)[index218_] = (self).f21
                        d_1271_v89_: _dafny.Seq
                        d_1271_v89_ = _dafny.SeqWithoutIsStrInference([d_1260_v79_])
                        d_1272_v90_: _dafny.Map
                        d_1272_v90_ = _dafny.Map({(d_1249_v68_).f28: False})
                        d_1273_v91_: _dafny.Seq
                        d_1273_v91_ = _dafny.SeqWithoutIsStrInference([d_1258_v77_])
                        index219_ = default__.safeIndex(825, (d_1270_v88_).length(0))
                        index220_ = default__.safeIndex(6, (d_1258_v77_).length(0))
                        rhs239_ = (p0) == ((0) - (default__.safeDivisionInt((self).f20, (self).f20)))
                        rhs240_ = default__.fm39(globalState)
                        rhs241_ = not((((d_1266_v85_).fm4((d_1261_v80_).cardinality, d_1271_v89_, (0) - (len(d_1272_v90_)), globalState)) and ((p2).f21)) and ((self).f24))
                        rhs242_ = (d_1273_v91_)[default__.safeIndex((p2).f20, len(d_1273_v91_))]
                        rhs243_ = (p2).f21
                        lhs217_ = globalState
                        lhs218_ = globalState
                        lhs219_ = d_1270_v88_
                        lhs220_ = default__.safeIndex(825, (d_1270_v88_).length(0))
                        lhs221_ = d_1258_v77_
                        lhs222_ = default__.safeIndex(6, (d_1258_v77_).length(0))
                        lhs217_.f18 = rhs239_
                        d_1269_v87_ = rhs240_
                        lhs218_.f3 = rhs241_
                        lhs219_[lhs220_] = rhs242_
                        lhs221_[lhs222_] = rhs243_
                        (globalState).f14 = (p2).f21
                    d_1274_v92_: _dafny.MultiSet
                    d_1274_v92_ = _dafny.MultiSet([(self).f21])
                    d_1275_v93_: D9
                    d_1275_v93_ = D9_DC23(455, p1)
                    d_1276_v94_: C0
                    nw216_ = C0()
                    nw216_.ctor__()
                    d_1276_v94_ = nw216_
                    d_1277_v95_: _dafny.MultiSet
                    d_1277_v95_ = _dafny.MultiSet([d_1276_v94_])
                    d_1278_v96_: _dafny.Map
                    d_1278_v96_ = _dafny.Map({(p2).f21: (p2).f21})
                    d_1279_v97_: _dafny.MultiSet
                    d_1279_v97_ = _dafny.MultiSet([(0) - ((p2).f20), (p2).f20])
                    d_1280_v98_: _dafny.Array
                    nw217_ = _dafny.Array(None, 17)
                    nw217_[int(0)] = False
                    nw217_[int(1)] = p1
                    nw217_[int(2)] = (p2).f21
                    nw217_[int(3)] = (d_1274_v92_).issubset(d_1274_v92_)
                    nw217_[int(4)] = not((d_1275_v93_).cf32)
                    nw217_[int(5)] = (d_1244_v64_) != (d_1244_v64_)
                    nw217_[int(6)] = p1
                    nw217_[int(7)] = p1
                    nw217_[int(8)] = (d_1276_v94_) in (d_1277_v95_)
                    nw217_[int(9)] = (p2).f21
                    nw217_[int(10)] = (d_1249_v68_).fm0(_dafny.Map({d_1249_v68_.f27: (p2).f20}), d_1244_v64_, _dafny.SeqWithoutIsStrInference([d_1278_v96_, d_1278_v96_]), globalState)
                    nw217_[int(11)] = (p2).f21
                    nw217_[int(12)] = (self).f21
                    nw217_[int(13)] = (self).f24
                    nw217_[int(14)] = ((p2).f20) <= ((d_1279_v97_).cardinality)
                    nw217_[int(15)] = (self).f21
                    nw217_[int(16)] = p1
                    d_1280_v98_ = nw217_
                    d_1280_v98_ = d_1280_v98_
                    pass
            pass
        d_1242_v63_ = (d_1242_v63_).set(605, (p2).f21)
        (globalState).f17 = (p2).f20
        d_1281_v99_: _dafny.Seq
        d_1281_v99_ = _dafny.SeqWithoutIsStrInference([False, (self).f24, (self).f24])
        d_1282_v100_: _dafny.Set
        d_1282_v100_ = _dafny.Set({(self).f20, (self).f20, len(d_1281_v99_), p0})
        d_1283_v101_: C3
        nw218_ = C3()
        nw218_.ctor__(default__.safeDivisionInt(p0, (self).f20), False, (self).f20, (d_1282_v100_).isdisjoint(d_1282_v100_))
        d_1283_v101_ = nw218_
        d_1284_v102_: _dafny.MultiSet
        d_1284_v102_ = _dafny.MultiSet([p1])
        d_1285_v104_: _dafny.Seq
        def iife103_():
            coll43_ = _dafny.Map()
            compr_43_: int
            for compr_43_ in (d_1242_v63_).keys.Elements:
                d_1286_v103_: int = compr_43_
                if (d_1286_v103_) in (d_1242_v63_):
                    coll43_[default__.safeModuloInt(d_1286_v103_, (d_1283_v101_).f25)] = (d_1283_v101_).f26
            return _dafny.Map(coll43_)
        d_1285_v104_ = _dafny.SeqWithoutIsStrInference([iife103_()
        ])
        d_1287_v105_: _dafny.Seq
        d_1287_v105_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rukj"))
        d_1288_v106_: _dafny.Seq
        d_1288_v106_ = _dafny.SeqWithoutIsStrInference([len(d_1287_v105_)])
        d_1289_v107_: _dafny.Array
        nw219_ = _dafny.Array(None, 23)
        nw219_[int(0)] = (p2).f20
        nw219_[int(1)] = (d_1283_v101_).f25
        nw219_[int(2)] = ((d_1283_v101_).f25) - (p0)
        nw219_[int(3)] = (d_1284_v102_).cardinality
        nw219_[int(4)] = len(d_1285_v104_)
        nw219_[int(5)] = (325) + ((d_1283_v101_).f25)
        nw219_[int(6)] = default__.safeModuloInt((d_1283_v101_).f25, (self).fm1((p2).f20, (d_1283_v101_).f25, (d_1283_v101_).f26, globalState))
        nw219_[int(7)] = (self).fm1(p0, (d_1283_v101_).f25, (self).f24, globalState)
        nw219_[int(8)] = len(d_1287_v105_)
        nw219_[int(9)] = (d_1283_v101_).f25
        nw219_[int(10)] = p0
        nw219_[int(11)] = (p2).fm1(len(d_1281_v99_), (self).f20, (p2).f21, globalState)
        nw219_[int(12)] = (d_1283_v101_).f25
        nw219_[int(13)] = p0
        nw219_[int(14)] = (d_1283_v101_).f25
        nw219_[int(15)] = (p2).fm1(p0, p0, (self).f24, globalState)
        nw219_[int(16)] = (d_1283_v101_).f25
        nw219_[int(17)] = len(d_1288_v106_)
        nw219_[int(18)] = (self).f20
        nw219_[int(19)] = (-384) * ((self).f20)
        nw219_[int(20)] = (d_1283_v101_).f25
        nw219_[int(21)] = (p2).f20
        nw219_[int(22)] = (0) - ((self).f20)
        d_1289_v107_ = nw219_
        r0 = d_1289_v107_
        r1 = default__.safeModuloInt(default__.safeModuloInt(500, p0), (0) - ((self).f20))
        return r0, r1

    @property
    def f24(self):
        return self._f24

class C5(T0):
    def  __init__(self):
        self._f20: int = int(0)
        self._f21: bool = False
        self.f23: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    def ctor__(self, f23, f20, f21):
        (self).f23 = f23
        (self)._f20 = f20
        (self)._f21 = f21

    def fm0(self, p0, p1, p2, globalState):
        return not((self).f21)

    def fm1(self, p0, p1, p2, globalState):
        def iife104_():
            coll44_ = _dafny.Map()
            compr_44_: int
            for compr_44_ in _dafny.IntegerRange(-451, -31):
                d_1290_v0_: int = compr_44_
                if ((-451) <= (d_1290_v0_)) and ((d_1290_v0_) < (-31)):
                    coll44_[(d_1290_v0_) + ((self).f20)] = self.f23
            return _dafny.Map(coll44_)
        return (len(_dafny.Set({(self).f20, (self).f20}))) + (len((iife104_()
        ) | (_dafny.Map({(self).f20: (self).f21}))))

    def fm5(self, p0, globalState):
        return self.f23

    def m0(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_1291_v0_: _dafny.Seq
        d_1291_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjgjhukoy"))
        d_1292_v1_: _dafny.Map
        d_1292_v1_ = _dafny.Map({len(d_1291_v0_): (self).f21})
        d_1293_v2_: _dafny.MultiSet
        d_1293_v2_ = _dafny.MultiSet([d_1292_v1_])
        d_1294_v3_: _dafny.Seq
        d_1294_v3_ = _dafny.SeqWithoutIsStrInference([default__.fm6(False, d_1293_v2_, globalState)])
        d_1295_v4_: _dafny.Map
        d_1295_v4_ = _dafny.Map({default__.fm7((self).f20, p0, globalState): p1})
        d_1296_v5_: _dafny.Map
        d_1296_v5_ = _dafny.Map({d_1294_v3_: (p0) + (len(d_1295_v4_))})
        d_1297_v6_: _dafny.Set
        d_1297_v6_ = _dafny.Set({(self).fm5(not(p1), globalState), (self).f21})
        d_1298_v7_: _dafny.Seq
        d_1298_v7_ = _dafny.SeqWithoutIsStrInference([(self).f20, (0) - (len(d_1297_v6_))])
        d_1296_v5_ = (d_1296_v5_).set((d_1294_v3_ if (self).f21 else default__.fm8((self).f20, (self).f21, d_1298_v7_, self.f23, globalState)), 820)
        d_1299_v9_: _dafny.Seq
        d_1299_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f21: p1})])
        d_1300_v10_: str
        d_1300_v10_ = _dafny.CodePoint('p')
        d_1301_v11_: _dafny.Map
        d_1301_v11_ = _dafny.Map({(self).f20: 121})
        d_1302_v12_: _dafny.Seq
        d_1302_v12_ = _dafny.SeqWithoutIsStrInference([(self).fm0(d_1301_v11_, _dafny.SeqWithoutIsStrInference([d_1300_v10_ for d_1303_i0_ in range(default__.abs(384))]), d_1299_v9_, globalState), self.f23, not(self.f23)])
        d_1304_v13_: _dafny.Map
        d_1304_v13_ = _dafny.Map({self.f23: p1})
        d_1305_v14_: _dafny.Array
        nw220_ = _dafny.Array(None, 22)
        nw220_[int(0)] = not((self).f21)
        nw220_[int(1)] = p1
        nw220_[int(2)] = p1
        nw220_[int(3)] = p1
        nw220_[int(4)] = not((p1) == (self.f23))
        nw220_[int(5)] = True
        nw220_[int(6)] = ((self).f20) <= ((self).f20)
        nw220_[int(7)] = p1
        nw220_[int(8)] = (self).f21
        nw220_[int(9)] = False
        nw220_[int(10)] = p1
        def iife105_():
            coll45_ = _dafny.Map()
            compr_45_: int
            for compr_45_ in (d_1298_v7_).Elements:
                d_1306_v8_: int = compr_45_
                if (d_1306_v8_) in (d_1298_v7_):
                    coll45_[default__.safeDivisionInt(d_1306_v8_, p0)] = (self).f20
            return _dafny.Map(coll45_)
        nw220_[int(11)] = (self).fm0(iife105_()
        , _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), d_1299_v9_, globalState)
        nw220_[int(12)] = (d_1300_v10_) not in (d_1291_v0_)
        nw220_[int(13)] = ((self).f20) >= ((self).f20)
        nw220_[int(14)] = self.f23
        nw220_[int(15)] = self.f23
        nw220_[int(16)] = (d_1302_v12_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eghu"))), len(d_1302_v12_))]
        nw220_[int(17)] = self.f23
        nw220_[int(18)] = ((self).f20) == (len(d_1304_v13_))
        nw220_[int(19)] = (self).f21
        nw220_[int(20)] = not ((self).f21) or (self.f23)
        nw220_[int(21)] = (p0) >= ((self).f20)
        d_1305_v14_ = nw220_
        index221_ = default__.safeIndex(886, (d_1305_v14_).length(0))
        (d_1305_v14_)[index221_] = p1
        index222_ = default__.safeIndex(886, (d_1305_v14_).length(0))
        (d_1305_v14_)[index222_] = not(self.f23)
        d_1307_i1_: int
        d_1307_i1_ = 0
        with _dafny.label("5"):
            while (len(d_1291_v0_)) <= ((self).f20):
                with _dafny.c_label("5"):
                    if (d_1307_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_1307_i1_ = (d_1307_i1_) + (1)
                    d_1308_v15_: _dafny.Array
                    def lambda60_(d_1309_p0_):
                        def lambda61_(d_1310_i2_):
                            return default__.safeModuloInt(d_1310_i2_, d_1309_p0_)

                        return lambda61_

                    init36_ = lambda60_(p0)
                    nw221_ = _dafny.Array(None, 15)
                    for i0_36_ in range(nw221_.length(0)):
                        nw221_[i0_36_] = init36_(i0_36_)
                    d_1308_v15_ = nw221_
                    index223_ = default__.safeIndex(166, (d_1308_v15_).length(0))
                    (d_1308_v15_)[index223_] = p0
                    d_1311_v16_: _dafny.MultiSet
                    d_1311_v16_ = _dafny.MultiSet([(self).f20, (self).f20])
                    index224_ = default__.safeIndex(166, (d_1308_v15_).length(0))
                    (d_1308_v15_)[index224_] = (p0) - ((d_1311_v16_).cardinality)
                    (globalState).f18 = self.f23
                    d_1312_v17_: _dafny.Map
                    d_1312_v17_ = _dafny.Map({d_1291_v0_: (d_1308_v15_)[default__.safeIndex(166, (d_1308_v15_).length(0))]})
                    index225_ = default__.safeIndex(166, (d_1308_v15_).length(0))
                    (d_1308_v15_)[index225_] = len(((d_1312_v17_) | (d_1312_v17_)) | (d_1312_v17_))
                    if ((self).f20) > ((self).f20):
                        d_1313_v18_: _dafny.Array
                        nw222_ = _dafny.Array(_dafny.Seq({}), 1)
                        d_1313_v18_ = nw222_
                        d_1314_v19_: _dafny.Map
                        d_1314_v19_ = _dafny.Map({d_1313_v18_: not ((self).f21) or (self.f23)})
                        index226_ = default__.safeIndex(166, (d_1308_v15_).length(0))
                        rhs244_ = d_1305_v14_
                        rhs245_ = 273
                        rhs246_ = (self).fm0(d_1301_v11_, (_dafny.SeqWithoutIsStrInference([d_1300_v10_ for d_1315_i3_ in range(default__.abs(916))])) + (d_1291_v0_), d_1299_v9_, globalState)
                        rhs247_ = ((d_1314_v19_)[d_1313_v18_] if (d_1313_v18_) in (d_1314_v19_) else (d_1305_v14_)[default__.safeIndex(886, (d_1305_v14_).length(0))])
                        lhs223_ = d_1308_v15_
                        lhs224_ = default__.safeIndex(166, (d_1308_v15_).length(0))
                        lhs225_ = globalState
                        lhs226_ = globalState
                        d_1305_v14_ = rhs244_
                        lhs223_[lhs224_] = rhs245_
                        lhs225_.f18 = rhs246_
                        lhs226_.f3 = rhs247_
                        d_1316_v20_: _dafny.Array
                        def lambda62_(d_1317_i4_):
                            return (self).f21

                        init37_ = lambda62_
                        nw223_ = _dafny.Array(None, 21)
                        for i0_37_ in range(nw223_.length(0)):
                            nw223_[i0_37_] = init37_(i0_37_)
                        d_1316_v20_ = nw223_
                        d_1318_v21_: _dafny.Seq
                        d_1318_v21_ = _dafny.SeqWithoutIsStrInference([d_1316_v20_])
                        d_1318_v21_ = _dafny.SeqWithoutIsStrInference([d_1316_v20_, d_1316_v20_])
                        (globalState).f18 = (p0) > (len(d_1302_v12_))
                        index227_ = default__.safeIndex(886, (d_1305_v14_).length(0))
                        (d_1305_v14_)[index227_] = (self).f21
                        d_1319_v22_: _dafny.Seq
                        d_1319_v22_ = _dafny.SeqWithoutIsStrInference([d_1305_v14_, d_1305_v14_])
                        d_1320_v23_: _dafny.Map
                        d_1320_v23_ = _dafny.Map({p0: d_1319_v22_})
                        d_1320_v23_ = (d_1320_v23_).set((self).f20, d_1319_v22_)
                    elif True:
                        (globalState).f4 = (self).f21
                        d_1321_v24_: _dafny.MultiSet
                        d_1321_v24_ = _dafny.MultiSet([d_1305_v14_, d_1305_v14_])
                        d_1322_v25_: _dafny.Seq
                        d_1322_v25_ = _dafny.SeqWithoutIsStrInference([d_1321_v24_])
                        d_1323_v26_: _dafny.Seq
                        d_1323_v26_ = _dafny.SeqWithoutIsStrInference([(d_1298_v7_).set(default__.safeIndex(p0, len(d_1298_v7_)), (self).f20)])
                        d_1324_v27_: _dafny.Seq
                        d_1324_v27_ = (d_1323_v26_)[default__.safeIndex((self).f20, len(d_1323_v26_))]
                        d_1325_v28_: _dafny.Map
                        d_1325_v28_ = _dafny.Map({(d_1308_v15_)[default__.safeIndex(166, (d_1308_v15_).length(0))]: d_1321_v24_})
                        d_1326_v29_: D2
                        d_1326_v29_ = D2_DC2(d_1321_v24_)
                        d_1327_v30_: _dafny.Map
                        d_1327_v30_ = _dafny.Map({p1: (self).f20})
                        d_1328_v31_: _dafny.Array
                        nw224_ = _dafny.Array(None, 21)
                        nw224_[int(0)] = d_1321_v24_
                        nw224_[int(1)] = d_1321_v24_
                        nw224_[int(2)] = (d_1322_v25_)[default__.safeIndex(len(default__.fm9(len(d_1291_v0_), d_1324_v27_, (self).f20, globalState)), len(d_1322_v25_))]
                        nw224_[int(3)] = d_1321_v24_
                        nw224_[int(4)] = (d_1322_v25_)[default__.safeIndex((d_1308_v15_)[default__.safeIndex(166, (d_1308_v15_).length(0))], len(d_1322_v25_))]
                        nw224_[int(5)] = d_1321_v24_
                        nw224_[int(6)] = ((d_1325_v28_)[(self).f20] if ((self).f20) in (d_1325_v28_) else d_1321_v24_)
                        nw224_[int(7)] = d_1321_v24_
                        nw224_[int(8)] = d_1321_v24_
                        nw224_[int(9)] = d_1321_v24_
                        nw224_[int(10)] = (d_1321_v24_).set(d_1305_v14_, default__.abs(p0))
                        nw224_[int(11)] = (d_1326_v29_).cf2
                        nw224_[int(12)] = (_dafny.MultiSet([d_1305_v14_])) - (d_1321_v24_)
                        nw224_[int(13)] = d_1321_v24_
                        nw224_[int(14)] = d_1321_v24_
                        nw224_[int(15)] = d_1321_v24_
                        nw224_[int(16)] = (d_1321_v24_).set(d_1305_v14_, default__.abs((self).fm1((self).f20, len(d_1327_v30_), True, globalState)))
                        nw224_[int(17)] = (d_1321_v24_).intersection(_dafny.MultiSet([d_1305_v14_, d_1305_v14_]))
                        nw224_[int(18)] = d_1321_v24_
                        nw224_[int(19)] = d_1321_v24_
                        nw224_[int(20)] = d_1321_v24_
                        d_1328_v31_ = nw224_
                        index228_ = default__.safeIndex(764, (d_1328_v31_).length(0))
                        (d_1328_v31_)[index228_] = d_1321_v24_
                        index229_ = default__.safeIndex(764, (d_1328_v31_).length(0))
                        (d_1328_v31_)[index229_] = (d_1322_v25_)[default__.safeIndex((self).f20, len(d_1322_v25_))]
                        (globalState).f3 = not((d_1305_v14_)[default__.safeIndex(886, (d_1305_v14_).length(0))])
                        (globalState).f18 = not((self).f21)
                        d_1329_v32_: C0
                        nw225_ = C0()
                        nw225_.ctor__()
                        d_1329_v32_ = nw225_
                    pass
            pass
        if self.f23:
            d_1330_v33_: _dafny.Array
            nw226_ = _dafny.Array(int(0), 15)
            d_1330_v33_ = nw226_
            index230_ = default__.safeIndex(234, (d_1330_v33_).length(0))
            (d_1330_v33_)[index230_] = ((self).f20 if (self).f21 else (self).f20)
            index231_ = default__.safeIndex(234, (d_1330_v33_).length(0))
            (d_1330_v33_)[index231_] = default__.safeDivisionInt((self).f20, (self).f20)
            index232_ = default__.safeIndex(234, (d_1330_v33_).length(0))
            (d_1330_v33_)[index232_] = len((d_1301_v11_) | (_dafny.Map({(self).f20: 204})))
            d_1300_v10_ = d_1300_v10_
            d_1331_v34_: _dafny.Seq
            d_1331_v34_ = _dafny.SeqWithoutIsStrInference([d_1305_v14_, d_1305_v14_])
            d_1332_v35_: D3
            d_1332_v35_ = D3_DC5(d_1331_v34_)
            d_1333_v36_: _dafny.MultiSet
            d_1333_v36_ = _dafny.MultiSet([_dafny.CodePoint('q'), d_1300_v10_, d_1300_v10_])
            index233_ = default__.safeIndex(234, (d_1330_v33_).length(0))
            rhs248_ = (self).f20
            rhs249_ = ((d_1332_v35_ if True else d_1332_v35_)).cf9
            rhs250_ = (_dafny.MultiSet([d_1300_v10_])).issubset(d_1333_v36_)
            lhs227_ = d_1330_v33_
            lhs228_ = default__.safeIndex(234, (d_1330_v33_).length(0))
            lhs227_[lhs228_] = rhs248_
            d_1331_v34_ = rhs249_
            r1 = rhs250_
            index234_ = default__.safeIndex(886, (d_1305_v14_).length(0))
            index235_ = default__.safeIndex(234, (d_1330_v33_).length(0))
            rhs251_ = not((714) > ((self).f20))
            rhs252_ = _dafny.SeqWithoutIsStrInference([(self).f20 for d_1334_i5_ in range(default__.abs(955))])
            rhs253_ = default__.safeModuloInt((p0) * (len(d_1302_v12_)), p0)
            lhs229_ = d_1305_v14_
            lhs230_ = default__.safeIndex(886, (d_1305_v14_).length(0))
            lhs231_ = d_1330_v33_
            lhs232_ = default__.safeIndex(234, (d_1330_v33_).length(0))
            lhs229_[lhs230_] = rhs251_
            d_1298_v7_ = rhs252_
            lhs231_[lhs232_] = rhs253_
        elif True:
            d_1335_v37_: _dafny.MultiSet
            d_1335_v37_ = _dafny.MultiSet([d_1305_v14_])
            d_1336_v38_: D2
            d_1336_v38_ = D2_DC2((d_1335_v37_).intersection((_dafny.MultiSet([d_1305_v14_, d_1305_v14_, d_1305_v14_, d_1305_v14_, d_1305_v14_])).set(d_1305_v14_, default__.abs(438))))
            d_1336_v38_ = d_1336_v38_
            d_1337_v39_: _dafny.Seq
            d_1337_v39_ = d_1298_v7_
            d_1338_v40_: _dafny.Set
            d_1338_v40_ = _dafny.Set({d_1337_v39_})
            d_1339_v41_: _dafny.MultiSet
            d_1339_v41_ = _dafny.MultiSet([default__.fm10(True, (self).f21, len(d_1291_v0_), globalState)])
            def iife106_():
                coll46_ = _dafny.Set()
                compr_46_: _dafny.Seq
                for compr_46_ in (d_1339_v41_).Elements:
                    d_1340_v42_: _dafny.Seq = compr_46_
                    if (d_1340_v42_) in (d_1339_v41_):
                        coll46_ = coll46_.union(_dafny.Set([d_1340_v42_]))
                return _dafny.Set(coll46_)
            (globalState).f5 = len((d_1338_v40_).intersection(iife106_()
            ))
            d_1341_v43_: C0
            nw227_ = C0()
            nw227_.ctor__()
            d_1341_v43_ = nw227_
            d_1342_v44_: _dafny.Array
            nw228_ = _dafny.Array(int(0), 26)
            d_1342_v44_ = nw228_
            index236_ = default__.safeIndex(337, (d_1342_v44_).length(0))
            (d_1342_v44_)[index236_] = p0
            index237_ = default__.safeIndex(337, (d_1342_v44_).length(0))
            rhs254_ = (0) - (p0)
            rhs255_ = (self).f20
            rhs256_ = ((self).f20) + (default__.safeModuloInt((self).f20, len(d_1291_v0_)))
            lhs233_ = d_1342_v44_
            lhs234_ = default__.safeIndex(337, (d_1342_v44_).length(0))
            lhs235_ = globalState
            lhs236_ = globalState
            lhs233_[lhs234_] = rhs254_
            lhs235_.f17 = rhs255_
            lhs236_.f17 = rhs256_
            (globalState).f18 = not (False) or (self.f23)
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1305_v14_).length(0)):
            d_1343_i6_: int = guard_loop_3_
            if (True) and (((0) <= (d_1343_i6_)) and ((d_1343_i6_) < ((d_1305_v14_).length(0)))):
                (d_1305_v14_)[(d_1343_i6_)] = not(p1)
        d_1344_v45_: _dafny.MultiSet
        d_1344_v45_ = _dafny.MultiSet([self.f23, self.f23])
        d_1345_v46_: D2
        d_1345_v46_ = D2_DC3(len(d_1302_v12_), (self.f23) not in (d_1344_v45_), (self).f21, p0, ((self).fm5((d_1305_v14_)[default__.safeIndex(886, (d_1305_v14_).length(0))], globalState)) or ((self).f21))
        source16_ = d_1345_v46_
        if source16_.is_DC3:
            d_1346___mcc_h0_ = source16_.cf3
            d_1347___mcc_h1_ = source16_.cf4
            d_1348___mcc_h2_ = source16_.cf5
            d_1349___mcc_h3_ = source16_.cf6
            d_1350___mcc_h4_ = source16_.cf7
            d_1351_cf7_ = d_1350___mcc_h4_
            d_1352_cf6_ = d_1349___mcc_h3_
            d_1353_cf5_ = d_1348___mcc_h2_
            d_1354_cf4_ = d_1347___mcc_h1_
            d_1355_cf3_ = d_1346___mcc_h0_
            d_1356_v47_: _dafny.Array
            nw229_ = _dafny.Array(D2.default()(), 6)
            d_1356_v47_ = nw229_
            d_1357_v48_: _dafny.MultiSet
            d_1357_v48_ = _dafny.MultiSet([d_1305_v14_, (D4_DC9(d_1305_v14_)).cf14, d_1305_v14_])
            d_1358_v49_: D2
            d_1358_v49_ = D2_DC2(d_1357_v48_)
            index238_ = default__.safeIndex(501, (d_1356_v47_).length(0))
            (d_1356_v47_)[index238_] = d_1358_v49_
            index239_ = default__.safeIndex(501, (d_1356_v47_).length(0))
            (d_1356_v47_)[index239_] = d_1358_v49_
            index240_ = default__.safeIndex(886, (d_1305_v14_).length(0))
            (d_1305_v14_)[index240_] = (self).f21
            (globalState).f17 = p0
            d_1359_v50_: _dafny.Seq
            d_1359_v50_ = _dafny.SeqWithoutIsStrInference([d_1291_v0_])
            d_1360_v51_: _dafny.Seq
            d_1360_v51_ = _dafny.SeqWithoutIsStrInference([d_1291_v0_, d_1291_v0_, (d_1291_v0_) + (d_1291_v0_), d_1291_v0_, (d_1359_v50_)[default__.safeIndex(d_1352_cf6_, len(d_1359_v50_))]])
            d_1359_v50_ = d_1360_v51_
        elif source16_.is_DC2:
            d_1361___mcc_h5_ = source16_.cf2
            d_1362_cf2_ = d_1361___mcc_h5_
            (globalState).f5 = len(_dafny.SeqWithoutIsStrInference([d_1300_v10_ for d_1363_i7_ in range(default__.abs(696))]))
            (self).f23 = (d_1305_v14_)[default__.safeIndex(886, (d_1305_v14_).length(0))]
            (globalState).f17 = (self).f20
            d_1364_v52_: _dafny.Array
            def lambda63_(d_1365_v0_):
                def lambda64_(d_1366_i8_):
                    return (d_1365_v0_) + (d_1365_v0_)

                return lambda64_

            init38_ = lambda63_(d_1291_v0_)
            nw230_ = _dafny.Array(None, 20)
            for i0_38_ in range(nw230_.length(0)):
                nw230_[i0_38_] = init38_(i0_38_)
            d_1364_v52_ = nw230_
            index241_ = default__.safeIndex(618, (d_1364_v52_).length(0))
            (d_1364_v52_)[index241_] = d_1291_v0_
            index242_ = default__.safeIndex(618, (d_1364_v52_).length(0))
            (d_1364_v52_)[index242_] = d_1291_v0_
        elif True:
            d_1367___mcc_h6_ = source16_.cf8
            d_1368_cf8_ = d_1367___mcc_h6_
            d_1369_v53_: C0
            nw231_ = C0()
            nw231_.ctor__()
            d_1369_v53_ = nw231_
            d_1370_v54_: _dafny.Array
            def lambda65_(d_1371_v10_):
                def lambda66_(d_1372_i9_):
                    return _dafny.SeqWithoutIsStrInference([d_1371_v10_ for d_1373_i10_ in range(default__.abs(-378))])

                return lambda66_

            init39_ = lambda65_(d_1300_v10_)
            nw232_ = _dafny.Array(None, 22)
            for i0_39_ in range(nw232_.length(0)):
                nw232_[i0_39_] = init39_(i0_39_)
            d_1370_v54_ = nw232_
            index243_ = default__.safeIndex(730, (d_1370_v54_).length(0))
            (d_1370_v54_)[index243_] = d_1291_v0_
            index244_ = default__.safeIndex(730, (d_1370_v54_).length(0))
            (d_1370_v54_)[index244_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iby"))) + (d_1291_v0_)
            if (p0) > ((self).f20):
                (globalState).f17 = (self).fm1((self).f20, (self).f20, (p1 if p1 else (d_1305_v14_)[default__.safeIndex(886, (d_1305_v14_).length(0))]), globalState)
                d_1374_v55_: _dafny.Map
                d_1374_v55_ = _dafny.Map({(d_1305_v14_)[default__.safeIndex(886, (d_1305_v14_).length(0))]: d_1370_v54_})
                d_1375_v56_: _dafny.MultiSet
                d_1375_v56_ = _dafny.MultiSet([(self).f20, p0, (self).f20])
                d_1370_v54_ = ((d_1374_v55_)[(d_1375_v56_).ispropersubset(d_1375_v56_)] if ((d_1375_v56_).ispropersubset(d_1375_v56_)) in (d_1374_v55_) else d_1370_v54_)
                d_1376_v57_: _dafny.Map
                d_1376_v57_ = _dafny.Map({(d_1345_v46_).cf6: d_1375_v56_})
                d_1376_v57_ = (d_1376_v57_).set(797, default__.fm11(len(d_1298_v7_), (self).f21, globalState))
                (self).f23 = (default__.fm12(p0, d_1300_v10_, (self).f21, ((d_1301_v11_)[(self).f20] if ((self).f20) in (d_1301_v11_) else (self).f20), globalState)) <= (((d_1370_v54_)[default__.safeIndex(730, (d_1370_v54_).length(0))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avxghk"))))
                d_1298_v7_ = (d_1298_v7_ if p1 else d_1298_v7_)
            elif True:
                d_1377_v58_: C0
                nw233_ = C0()
                nw233_.ctor__()
                d_1377_v58_ = nw233_
                d_1378_v59_: _dafny.Map
                d_1378_v59_ = _dafny.Map({(self).f21: d_1292_v1_})
                d_1379_v60_: _dafny.Seq
                d_1379_v60_ = _dafny.SeqWithoutIsStrInference([d_1301_v11_])
                d_1380_v61_: _dafny.Map
                d_1380_v61_ = _dafny.Map({D2_DC3(p0, self.f23, (self).f21, p0, True): (d_1369_v53_).fm4(p0, d_1379_v60_, (self).f20, globalState)})
                d_1381_v62_: _dafny.Seq
                d_1381_v62_ = _dafny.SeqWithoutIsStrInference([d_1305_v14_])
                d_1382_v63_: D3
                d_1382_v63_ = D3_DC6((self).f20)
                d_1383_v64_: _dafny.Set
                d_1383_v64_ = _dafny.Set({p0, (self).f20, len((d_1370_v54_)[default__.safeIndex(730, (d_1370_v54_).length(0))])})
                d_1384_v65_: D4
                d_1384_v65_ = D4_DC11(len(d_1302_v12_), (self).f21, p1)
                d_1385_v66_: _dafny.Array
                nw234_ = _dafny.Array(None, 25)
                nw234_[int(0)] = (0) - (p0)
                nw234_[int(1)] = len(d_1378_v59_)
                nw234_[int(2)] = p0
                nw234_[int(3)] = p0
                nw234_[int(4)] = len(d_1302_v12_)
                nw234_[int(5)] = len((_dafny.Map({d_1345_v46_: (self).f21})) | (d_1380_v61_))
                nw234_[int(6)] = len(d_1381_v62_)
                nw234_[int(7)] = p0
                nw234_[int(8)] = (p0) * (418)
                nw234_[int(9)] = default__.safeDivisionInt(p0, (self).f20)
                nw234_[int(10)] = ((self).f20) - ((d_1382_v63_).cf10)
                nw234_[int(11)] = len(d_1383_v64_)
                nw234_[int(12)] = (0) - (p0)
                nw234_[int(13)] = (0) - (((self).f20) * (p0))
                nw234_[int(14)] = len(default__.fm12((self).f20, _dafny.CodePoint('n'), (d_1369_v53_).fm3(len(_dafny.SeqWithoutIsStrInference([d_1300_v10_ for d_1386_i11_ in range(default__.abs(431))])), False, globalState), (self).fm1(-961, 442, (self).f21, globalState), globalState))
                nw234_[int(15)] = 390
                nw234_[int(16)] = (self).f20
                nw234_[int(17)] = len(d_1298_v7_)
                nw234_[int(18)] = (self).f20
                nw234_[int(19)] = (p0) + ((d_1384_v65_).cf17)
                nw234_[int(20)] = p0
                nw234_[int(21)] = (self).f20
                nw234_[int(22)] = (self).f20
                nw234_[int(23)] = p0
                nw234_[int(24)] = p0
                d_1385_v66_ = nw234_
                index245_ = default__.safeIndex(471, (d_1385_v66_).length(0))
                (d_1385_v66_)[index245_] = (0) - (p0)
                d_1387_v67_: _dafny.MultiSet
                d_1387_v67_ = _dafny.MultiSet([len(d_1304_v13_)])
                d_1388_v68_: _dafny.Seq
                d_1388_v68_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p0]), d_1387_v67_, d_1387_v67_])
                index246_ = default__.safeIndex(471, (d_1385_v66_).length(0))
                (d_1385_v66_)[index246_] = (0) - (len((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p0]) for d_1389_i12_ in range(default__.abs(498))])) + (d_1388_v68_)))
                d_1390_v69_: C0
                nw235_ = C0()
                nw235_.ctor__()
                d_1390_v69_ = nw235_
                d_1391_v70_: C0
                nw236_ = C0()
                nw236_.ctor__()
                d_1391_v70_ = nw236_
                d_1385_v66_ = d_1385_v66_
            def iife107_():
                coll47_ = _dafny.Map()
                compr_47_: _dafny.Seq
                for compr_47_ in (_dafny.SeqWithoutIsStrInference([d_1298_v7_])).Elements:
                    d_1392_v71_: _dafny.Seq = compr_47_
                    if (d_1392_v71_) in (_dafny.SeqWithoutIsStrInference([d_1298_v7_])):
                        coll47_[d_1392_v71_] = p1
                return _dafny.Map(coll47_)
            (globalState).f5 = len(iife107_()
            )
        d_1393_v72_: _dafny.Map
        d_1393_v72_ = _dafny.Map({d_1305_v14_: (self).f21})
        d_1394_v73_: C0
        nw237_ = C0()
        nw237_.ctor__()
        d_1394_v73_ = nw237_
        d_1395_v74_: _dafny.Set
        d_1395_v74_ = _dafny.Set({d_1394_v73_})
        d_1396_v75_: T0
        nw238_ = C4()
        nw238_.ctor__(p1, 471, True)
        d_1396_v75_ = nw238_
        d_1397_v76_: _dafny.Map
        d_1397_v76_ = _dafny.Map({(self).f20: d_1396_v75_})
        d_1398_v77_: _dafny.MultiSet
        d_1398_v77_ = _dafny.MultiSet([(0) - (p0), (self).f20])
        d_1399_v78_: _dafny.Map
        d_1399_v78_ = _dafny.Map({(self).f20: d_1398_v77_})
        d_1400_v79_: _dafny.Map
        d_1400_v79_ = _dafny.Map({d_1396_v75_: p0})
        d_1401_v80_: _dafny.Map
        d_1401_v80_ = _dafny.Map({self.f23: d_1344_v45_})
        d_1402_v81_: _dafny.Array
        nw239_ = _dafny.Array(None, 16)
        nw239_[int(0)] = (self).f20
        nw239_[int(1)] = p0
        nw239_[int(2)] = len(d_1395_v74_)
        nw239_[int(3)] = 815
        nw239_[int(4)] = p0
        nw239_[int(5)] = len(d_1397_v76_)
        nw239_[int(6)] = (self).f20
        nw239_[int(7)] = (d_1396_v75_).f20
        nw239_[int(8)] = (self).f20
        nw239_[int(9)] = (d_1396_v75_).f20
        nw239_[int(10)] = (self).f20
        nw239_[int(11)] = p0
        nw239_[int(12)] = (((d_1399_v78_)[len(d_1400_v79_)] if (len(d_1400_v79_)) in (d_1399_v78_) else d_1398_v77_)).cardinality
        nw239_[int(13)] = (0) - ((d_1344_v45_).cardinality)
        nw239_[int(14)] = len(d_1301_v11_)
        nw239_[int(15)] = len(d_1401_v80_)
        d_1402_v81_ = nw239_
        d_1403_v82_: _dafny.Seq
        d_1403_v82_ = _dafny.SeqWithoutIsStrInference([d_1402_v81_, d_1402_v81_])
        d_1404_v83_: _dafny.Map
        d_1404_v83_ = _dafny.Map({len(d_1393_v72_): (d_1403_v82_).set(default__.safeIndex((d_1298_v7_)[default__.safeIndex(244, len(d_1298_v7_))], len(d_1403_v82_)), d_1402_v81_)})
        r0 = (((d_1404_v83_)[p0] if (p0) in (d_1404_v83_) else d_1403_v82_)) + (d_1403_v82_)
        r1 = (d_1305_v14_)[default__.safeIndex(886, (d_1305_v14_).length(0))]
        return r0, r1


class C6(T0):
    def  __init__(self):
        self._f20: int = int(0)
        self._f21: bool = False
        self._f22: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    def ctor__(self, f22, f20, f21):
        (self)._f22 = f22
        (self)._f20 = f20
        (self)._f21 = f21

    def fm0(self, p0, p1, p2, globalState):
        return (self).f21

    def fm1(self, p0, p1, p2, globalState):
        return ((0) - ((self).f20)) - ((self).f20)

    def m0(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_1405_i0_: int
        d_1405_i0_ = 0
        with _dafny.label("6"):
            while (self).f21:
                with _dafny.c_label("6"):
                    if (d_1405_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1405_i0_ = (d_1405_i0_) + (1)
                    d_1406_v0_: _dafny.Seq
                    d_1406_v0_ = _dafny.SeqWithoutIsStrInference([p0 for d_1407_i1_ in range(default__.abs(851))])
                    d_1408_v1_: _dafny.MultiSet
                    d_1408_v1_ = _dafny.MultiSet([p0])
                    d_1409_v2_: _dafny.Seq
                    d_1409_v2_ = _dafny.SeqWithoutIsStrInference([d_1408_v1_, d_1408_v1_, d_1408_v1_])
                    d_1410_v3_: _dafny.MultiSet
                    d_1410_v3_ = _dafny.MultiSet([p0, len(d_1409_v2_)])
                    d_1411_v4_: _dafny.Seq
                    d_1411_v4_ = _dafny.SeqWithoutIsStrInference([((d_1410_v3_)[p0] if (p0) in (d_1410_v3_) else (self).f20)])
                    d_1412_v5_: _dafny.Map
                    d_1412_v5_ = _dafny.Map({(True) == ((self).f21): (d_1411_v4_)})
                    d_1413_v6_: _dafny.Map
                    d_1413_v6_ = _dafny.Map({(self).f20: -218})
                    d_1414_v7_: _dafny.Map
                    d_1414_v7_ = _dafny.Map({p1: p1})
                    d_1415_v8_: bool
                    d_1415_v8_ = (self).f21
                    d_1416_v9_: _dafny.Seq
                    d_1416_v9_ = _dafny.SeqWithoutIsStrInference([d_1414_v7_, _dafny.Map({(self).f21: (d_1415_v8_)})])
                    d_1412_v5_ = (d_1412_v5_).set((self).fm0(d_1413_v6_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1417_i2_ in range(default__.abs(787))]), d_1416_v9_, globalState), d_1411_v4_)
                    source17_ = d_1406_v0_
                    d_1418___mcc_h0_ = source17_
                    d_1419_cf0_ = d_1418___mcc_h0_
                    d_1420_v10_: _dafny.Set
                    d_1420_v10_ = _dafny.Set({p1, (self).f21})
                    (globalState).f3 = (_dafny.Set({(self).f21})) != (d_1420_v10_)
                    d_1421_v11_: _dafny.MultiSet
                    d_1421_v11_ = _dafny.MultiSet([(self).f21])
                    d_1422_v12_: _dafny.Array
                    nw240_ = _dafny.Array(None, 8)
                    nw240_[int(0)] = p0
                    nw240_[int(1)] = (d_1421_v11_).cardinality
                    nw240_[int(2)] = (self).fm1(p0, -743, p1, globalState)
                    nw240_[int(3)] = p0
                    nw240_[int(4)] = (p0) * (p0)
                    nw240_[int(5)] = 296
                    nw240_[int(6)] = (d_1419_cf0_)[default__.safeIndex((0) - ((self).f20), len(d_1419_cf0_))]
                    nw240_[int(7)] = (p0) * (p0)
                    d_1422_v12_ = nw240_
                    d_1423_v13_: _dafny.Map
                    d_1423_v13_ = _dafny.Map({d_1411_v4_: d_1422_v12_})
                    d_1422_v12_ = ((d_1423_v13_)[(d_1406_v0_ if (self).f21 else d_1406_v0_)] if ((d_1406_v0_ if (self).f21 else d_1406_v0_)) in (d_1423_v13_) else d_1422_v12_)
                    d_1424_v14_: _dafny.Array
                    def lambda67_(d_1425_v8_):
                        def lambda68_(d_1426_i3_):
                            return d_1425_v8_

                        return lambda68_

                    init40_ = lambda67_(d_1415_v8_)
                    nw241_ = _dafny.Array(None, 28)
                    for i0_40_ in range(nw241_.length(0)):
                        nw241_[i0_40_] = init40_(i0_40_)
                    d_1424_v14_ = nw241_
                    index247_ = default__.safeIndex(546, (d_1424_v14_).length(0))
                    (d_1424_v14_)[index247_] = d_1415_v8_
                    index248_ = default__.safeIndex(546, (d_1424_v14_).length(0))
                    rhs257_ = d_1415_v8_
                    rhs258_ = len(d_1420_v10_)
                    rhs259_ = ((p1 if (self).f21 else d_1415_v8_))
                    rhs260_ = ((_dafny.SeqWithoutIsStrInference([len((self).f22)])) + (d_1411_v4_)) + (d_1419_cf0_)
                    rhs261_ = p1
                    lhs237_ = d_1424_v14_
                    lhs238_ = default__.safeIndex(546, (d_1424_v14_).length(0))
                    lhs239_ = globalState
                    lhs240_ = globalState
                    lhs241_ = globalState
                    lhs237_[lhs238_] = rhs257_
                    lhs239_.f5 = rhs258_
                    lhs240_.f14 = rhs259_
                    d_1419_cf0_ = rhs260_
                    lhs241_.f14 = rhs261_
                    index249_ = default__.safeIndex(326, (d_1422_v12_).length(0))
                    (d_1422_v12_)[index249_] = default__.safeDivisionInt(p0, (self).fm1(p0, (self).f20, p1, globalState))
                    index250_ = default__.safeIndex(326, (d_1422_v12_).length(0))
                    (d_1422_v12_)[index250_] = 157
                    d_1414_v7_ = (d_1414_v7_).set(not((self).f21), p1)
                    d_1427_v15_: _dafny.Array
                    nw242_ = _dafny.Array(int(0), 21)
                    d_1427_v15_ = nw242_
                    d_1428_v16_: _dafny.Map
                    d_1428_v16_ = _dafny.Map({p1: d_1427_v15_})
                    d_1428_v16_ = (d_1428_v16_).set((self).f21, d_1427_v15_)
                    pass
            pass
        hi4_ = p0
        for d_1429_i4_ in range((0) - (-898), hi4_):
            d_1430_v17_: str
            d_1430_v17_ = _dafny.CodePoint('x')
            d_1431_v18_: _dafny.Map
            d_1431_v18_ = _dafny.Map({d_1430_v17_: default__.fm2(not((self).f21), (self).f20, p1, globalState)})
            d_1432_v19_: bool
            d_1432_v19_ = (self).f21
            d_1431_v18_ = (d_1431_v18_).set(d_1430_v17_, d_1432_v19_)
            (globalState).f5 = (self).f20
            d_1433_v20_: _dafny.Array
            nw243_ = _dafny.Array(int(0), 21)
            d_1433_v20_ = nw243_
            index251_ = default__.safeIndex(383, (d_1433_v20_).length(0))
            (d_1433_v20_)[index251_] = (p0) * (len((self).f22))
            index252_ = default__.safeIndex(383, (d_1433_v20_).length(0))
            rhs262_ = (0) - ((self).f20)
            rhs263_ = (self).f22
            lhs242_ = d_1433_v20_
            lhs243_ = default__.safeIndex(383, (d_1433_v20_).length(0))
            lhs244_ = globalState
            lhs242_[lhs243_] = rhs262_
            lhs244_.f15 = rhs263_
            d_1434_v21_: _dafny.Seq
            d_1434_v21_ = _dafny.SeqWithoutIsStrInference([p0])
            source18_ = d_1434_v21_
            d_1435___mcc_h1_ = source18_
            d_1436_cf0_ = d_1435___mcc_h1_
            (globalState).f18 = (self).f21
            d_1437_v22_: C0
            nw244_ = C0()
            nw244_.ctor__()
            d_1437_v22_ = nw244_
            d_1438_v23_: _dafny.Seq
            d_1438_v23_ = d_1436_cf0_
            d_1438_v23_ = d_1438_v23_
            d_1439_v24_: _dafny.Map
            d_1439_v24_ = _dafny.Map({d_1432_v19_: p1})
            d_1440_v25_: _dafny.Map
            d_1440_v25_ = _dafny.Map({(self).f21: (self).f21})
            d_1441_v26_: _dafny.Map
            d_1441_v26_ = _dafny.Map({p0: (d_1433_v20_)[default__.safeIndex(383, (d_1433_v20_).length(0))]})
            d_1442_v27_: _dafny.Seq
            d_1442_v27_ = _dafny.SeqWithoutIsStrInference([d_1440_v25_])
            d_1439_v24_ = (d_1439_v24_).set(d_1432_v19_, (not((self).f21) if ((d_1440_v25_)[(self).fm0(d_1441_v26_, (self).f22, d_1442_v27_, globalState)] if ((self).fm0(d_1441_v26_, (self).f22, d_1442_v27_, globalState)) in (d_1440_v25_) else p1) else (self).f21))
        d_1443_v28_: _dafny.Array
        nw245_ = _dafny.Array(_dafny.Array(None, 0), 26)
        d_1443_v28_ = nw245_
        d_1444_v29_: _dafny.Array
        def lambda69_(d_1445_p0_):
            def lambda70_(d_1446_i5_):
                return (d_1446_i5_) + (d_1445_p0_)

            return lambda70_

        init41_ = lambda69_(p0)
        nw246_ = _dafny.Array(None, 19)
        for i0_41_ in range(nw246_.length(0)):
            nw246_[i0_41_] = init41_(i0_41_)
        d_1444_v29_ = nw246_
        d_1447_v30_: T0
        nw247_ = C1()
        nw247_.ctor__((self).f20, d_1444_v29_, (self).fm1((self).f20, p0, p1, globalState), False)
        d_1447_v30_ = nw247_
        d_1448_v31_: _dafny.Map
        d_1448_v31_ = _dafny.Map({d_1447_v30_: (self).f21})
        d_1449_v32_: _dafny.Map
        d_1449_v32_ = _dafny.Map({d_1448_v31_: d_1443_v28_})
        d_1450_v33_: _dafny.Map
        d_1450_v33_ = _dafny.Map({p0: (d_1447_v30_).f20})
        d_1451_v34_: _dafny.Map
        d_1451_v34_ = _dafny.Map({not(True): True})
        d_1452_v35_: _dafny.Seq
        d_1452_v35_ = _dafny.SeqWithoutIsStrInference([d_1451_v34_])
        d_1453_v36_: _dafny.Array
        nw248_ = _dafny.Array(None, 13)
        nw248_[int(0)] = d_1443_v28_
        nw248_[int(1)] = d_1443_v28_
        nw248_[int(2)] = d_1443_v28_
        nw248_[int(3)] = d_1443_v28_
        nw248_[int(4)] = d_1443_v28_
        nw248_[int(5)] = d_1443_v28_
        nw248_[int(6)] = (d_1443_v28_ if p1 else d_1443_v28_)
        nw248_[int(7)] = d_1443_v28_
        nw248_[int(8)] = ((d_1449_v32_)[_dafny.Map({d_1447_v30_: (self).fm0(d_1450_v33_, (self).f22, d_1452_v35_, globalState)})] if (_dafny.Map({d_1447_v30_: (self).fm0(d_1450_v33_, (self).f22, d_1452_v35_, globalState)})) in (d_1449_v32_) else d_1443_v28_)
        nw248_[int(9)] = d_1443_v28_
        nw248_[int(10)] = d_1443_v28_
        nw248_[int(11)] = d_1443_v28_
        nw248_[int(12)] = d_1443_v28_
        d_1453_v36_ = nw248_
        index253_ = default__.safeIndex(485, (d_1453_v36_).length(0))
        (d_1453_v36_)[index253_] = d_1443_v28_
        d_1454_v37_: str
        d_1454_v37_ = _dafny.CodePoint('i')
        index254_ = default__.safeIndex(485, (d_1453_v36_).length(0))
        rhs264_ = d_1443_v28_
        rhs265_ = ((len(default__.fm12((d_1447_v30_).f20, d_1454_v37_, p1, (self).f20, globalState))) + ((d_1447_v30_).f20) if False else default__.safeModuloInt((d_1447_v30_).f20, p0))
        rhs266_ = ((d_1454_v37_ if p1 else d_1454_v37_)) in (_dafny.SeqWithoutIsStrInference([d_1454_v37_ for d_1455_i6_ in range(default__.abs(469))]))
        lhs245_ = d_1453_v36_
        lhs246_ = default__.safeIndex(485, (d_1453_v36_).length(0))
        lhs247_ = globalState
        lhs248_ = globalState
        lhs245_[lhs246_] = rhs264_
        lhs247_.f17 = rhs265_
        lhs248_.f4 = rhs266_
        d_1447_v30_ = d_1447_v30_
        hi5_ = 585
        for d_1456_i7_ in range(((self).f20 if (d_1447_v30_).f21 else ((d_1450_v33_)[(d_1447_v30_).f20] if ((d_1447_v30_).f20) in (d_1450_v33_) else (self).f20)), hi5_):
            d_1457_v38_: bool
            d_1457_v38_ = p1
            d_1457_v38_ = d_1457_v38_
            d_1458_v39_: C1
            nw249_ = C1()
            nw249_.ctor__((d_1447_v30_).f20, d_1444_v29_, p0, (d_1447_v30_).f21)
            d_1458_v39_ = nw249_
            d_1458_v39_ = d_1458_v39_
            d_1459_v40_: _dafny.Set
            d_1459_v40_ = _dafny.Set({len(d_1451_v34_), (self).f20})
            d_1460_v41_: _dafny.Seq
            d_1460_v41_ = _dafny.SeqWithoutIsStrInference([(d_1459_v40_) - (d_1459_v40_)])
            d_1460_v41_ = d_1460_v41_
            if (d_1447_v30_).f21:
                (d_1458_v39_).f27 = p0
                d_1461_v42_: D8
                d_1461_v42_ = D8_DC21(d_1458_v39_.f27)
                d_1462_v43_: _dafny.MultiSet
                d_1462_v43_ = _dafny.MultiSet([((d_1461_v42_).cf29) > ((self).f20)])
                d_1462_v43_ = (d_1462_v43_) | (d_1462_v43_)
                d_1463_v44_: _dafny.Seq
                d_1463_v44_ = _dafny.SeqWithoutIsStrInference([len((self).f22)])
                d_1463_v44_ = (_dafny.SeqWithoutIsStrInference([(self).f20 for d_1464_i8_ in range(default__.abs(472))])) + ((d_1463_v44_) + (d_1463_v44_))
                (globalState).f15 = (self).f22
                d_1444_v29_ = (d_1458_v39_).f28
            elif True:
                d_1465_v45_: _dafny.Seq
                d_1465_v45_ = _dafny.SeqWithoutIsStrInference([d_1458_v39_.f27, (d_1447_v30_).f20])
                d_1466_v47_: _dafny.Set
                def iife108_():
                    coll48_ = _dafny.Set()
                    compr_48_: int
                    for compr_48_ in (_dafny.MultiSet(d_1465_v45_)).Elements:
                        d_1467_v46_: int = compr_48_
                        if (d_1467_v46_) in (_dafny.MultiSet(d_1465_v45_)):
                            coll48_ = coll48_.union(_dafny.Set([(d_1467_v46_) * ((_dafny.MultiSet([not(False)])).cardinality)]))
                    return _dafny.Set(coll48_)
                d_1466_v47_ = _dafny.Set({d_1459_v40_, d_1459_v40_, d_1459_v40_, iife108_()
                , d_1459_v40_})
                rhs267_ = (d_1447_v30_).f20
                rhs268_ = (self).f21
                rhs269_ = (d_1466_v47_) - (d_1466_v47_)
                lhs249_ = globalState
                lhs250_ = globalState
                lhs249_.f17 = rhs267_
                lhs250_.f4 = rhs268_
                d_1466_v47_ = rhs269_
                (globalState).f5 = p0
                d_1468_v48_: D13
                d_1468_v48_ = D13_DC37((d_1447_v30_).f21, (self).f20)
                (globalState).f17 = ((d_1447_v30_).f20) + ((d_1468_v48_).cf50)
                d_1469_v49_: _dafny.Array
                nw250_ = _dafny.Array(False, 18)
                d_1469_v49_ = nw250_
                d_1470_v50_: _dafny.Set
                d_1470_v50_ = _dafny.Set({(self).f21})
                d_1471_v51_: _dafny.Seq
                d_1471_v51_ = _dafny.SeqWithoutIsStrInference([d_1470_v50_])
                d_1472_v52_: D3
                d_1472_v52_ = D3_DC7(d_1471_v51_, p1)
                index255_ = default__.safeIndex(738, (d_1469_v49_).length(0))
                (d_1469_v49_)[index255_] = (d_1472_v52_).cf12
                index256_ = default__.safeIndex(738, (d_1469_v49_).length(0))
                (d_1469_v49_)[index256_] = p1
                d_1473_v53_: _dafny.Array
                d_1473_v53_ = d_1444_v29_
                d_1474_v54_: _dafny.Set
                d_1474_v54_ = _dafny.Set({d_1473_v53_})
                d_1475_v55_: _dafny.Seq
                d_1475_v55_ = _dafny.SeqWithoutIsStrInference([(d_1447_v30_).f21])
                d_1476_v56_: _dafny.Map
                d_1476_v56_ = _dafny.Map({(self).f21: d_1475_v55_})
                d_1477_v57_: _dafny.Map
                d_1477_v57_ = _dafny.Map({d_1474_v54_: (0) - (len(((d_1476_v56_)[(d_1447_v30_).f21] if ((d_1447_v30_).f21) in (d_1476_v56_) else d_1475_v55_)))})
                d_1477_v57_ = d_1477_v57_
        d_1478_v58_: _dafny.Set
        d_1478_v58_ = _dafny.Set({d_1444_v29_})
        hi6_ = len(d_1478_v58_)
        for d_1479_i9_ in range(default__.safeModuloInt((self).f20, p0), hi6_):
            d_1480_v59_: _dafny.Map
            d_1480_v59_ = _dafny.Map({(self).f20: (self).f21})
            d_1481_v60_: _dafny.Map
            d_1481_v60_ = _dafny.Map({(self).f20: d_1480_v59_})
            d_1482_v61_: D16
            d_1482_v61_ = D16_DC42(default__.fm40((self).f21, (d_1447_v30_).f21, (d_1447_v30_).f20, len(d_1481_v60_), globalState))
            (globalState).f17 = len((d_1482_v61_).cf54)
            (globalState).f15 = ((self).f22) + ((self).f22)
            d_1451_v34_ = d_1451_v34_
            if not(((d_1454_v37_) in ((self).f22)) == (not((d_1447_v30_).f21))):
                (globalState).f17 = (self).f20
                (globalState).f3 = True
                (globalState).f4 = (self).fm0(d_1450_v33_, ((self).f22) + ((self).f22), d_1452_v35_, globalState)
                d_1483_v62_: _dafny.Array
                nw251_ = _dafny.Array(_dafny.Seq({}), 12)
                d_1483_v62_ = nw251_
                d_1484_v63_: _dafny.Seq
                d_1484_v63_ = _dafny.SeqWithoutIsStrInference([(d_1447_v30_).f21])
                index257_ = default__.safeIndex(983, (d_1483_v62_).length(0))
                (d_1483_v62_)[index257_] = d_1484_v63_
                index258_ = default__.safeIndex(983, (d_1483_v62_).length(0))
                rhs270_ = (d_1484_v63_) + (_dafny.SeqWithoutIsStrInference([p1, (self).f21, (d_1447_v30_).f21, (d_1447_v30_).f21]))
                rhs271_ = ((d_1450_v33_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eopxabkhd")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eopxabkhd")))) in (d_1450_v33_) else (255 if p1 else (d_1447_v30_).f20))
                lhs251_ = d_1483_v62_
                lhs252_ = default__.safeIndex(983, (d_1483_v62_).length(0))
                lhs253_ = globalState
                lhs251_[lhs252_] = rhs270_
                lhs253_.f5 = rhs271_
                d_1485_v64_: _dafny.Array
                nw252_ = _dafny.Array(None, 19)
                nw252_[int(0)] = (self).f22
                nw252_[int(1)] = (self).f22
                nw252_[int(2)] = (self).f22
                nw252_[int(3)] = (self).f22
                nw252_[int(4)] = (self).f22
                nw252_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "diwso"))
                nw252_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkkwmspk"))
                nw252_[int(7)] = (self).f22
                nw252_[int(8)] = (self).f22
                nw252_[int(9)] = _dafny.SeqWithoutIsStrInference([d_1454_v37_ for d_1486_i10_ in range(default__.abs(629))])
                nw252_[int(10)] = (self).f22
                nw252_[int(11)] = (self).f22
                nw252_[int(12)] = _dafny.SeqWithoutIsStrInference([d_1454_v37_ for d_1487_i11_ in range(default__.abs(-284))])
                nw252_[int(13)] = (self).f22
                nw252_[int(14)] = (self).f22
                nw252_[int(15)] = (self).f22
                nw252_[int(16)] = (self).f22
                nw252_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikkiep"))
                nw252_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvy"))
                d_1485_v64_ = nw252_
                d_1488_v65_: _dafny.Map
                d_1488_v65_ = _dafny.Map({D11_DC28(d_1485_v64_): (d_1447_v30_).f21})
                d_1489_v66_: _dafny.Seq
                d_1489_v66_ = _dafny.SeqWithoutIsStrInference([d_1488_v65_])
                d_1490_v67_: D2
                d_1490_v67_ = D2_DC3((d_1447_v30_).f20, (self).f21, (d_1447_v30_).f21, (d_1447_v30_).f20, (d_1447_v30_).f21)
                d_1491_v68_: _dafny.Seq
                d_1491_v68_ = _dafny.SeqWithoutIsStrInference([d_1484_v63_])
                d_1492_v69_: _dafny.Array
                def lambda71_(d_1493_p1_):
                    def lambda72_(d_1494_i12_):
                        return d_1493_p1_

                    return lambda72_

                init42_ = lambda71_(p1)
                nw253_ = _dafny.Array(None, 3)
                for i0_42_ in range(nw253_.length(0)):
                    nw253_[i0_42_] = init42_(i0_42_)
                d_1492_v69_ = nw253_
                d_1495_v70_: _dafny.Set
                d_1495_v70_ = _dafny.Set({(d_1447_v30_).f20, (d_1447_v30_).f20})
                d_1496_v71_: _dafny.Map
                d_1496_v71_ = _dafny.Map({d_1492_v69_: d_1495_v70_})
                d_1497_v72_: _dafny.Set
                d_1497_v72_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([((d_1451_v34_)[(d_1447_v30_).f21] if ((d_1447_v30_).f21) in (d_1451_v34_) else not((d_1490_v67_).cf4))]), (d_1491_v68_)[default__.safeIndex(len((d_1484_v63_).set(default__.safeIndex(len(d_1496_v71_), len(d_1484_v63_)), (d_1447_v30_).f21)), len(d_1491_v68_))]})
                d_1498_v73_: _dafny.MultiSet
                d_1498_v73_ = _dafny.MultiSet([not((d_1447_v30_).f21)])
                d_1499_v74_: D4
                d_1499_v74_ = D4_DC10(d_1498_v73_, (self).f21)
                d_1500_v75_: _dafny.Seq
                d_1500_v75_ = _dafny.SeqWithoutIsStrInference([(d_1447_v30_).f20])
                rhs272_ = -116
                rhs273_ = (d_1447_v30_).fm1((0) - (p0), ((d_1447_v30_).f20) * ((d_1447_v30_).f20), (d_1499_v74_).cf16, globalState)
                rhs274_ = d_1489_v66_
                rhs275_ = d_1497_v72_
                rhs276_ = default__.safeDivisionInt((d_1447_v30_).f20, len(((d_1500_v75_).set(default__.safeIndex((self).f20, len(d_1500_v75_)), d_1479_i9_)) + (_dafny.SeqWithoutIsStrInference([(self).f20]))))
                lhs254_ = globalState
                lhs255_ = globalState
                lhs256_ = globalState
                lhs254_.f17 = rhs272_
                lhs255_.f17 = rhs273_
                d_1489_v66_ = rhs274_
                d_1497_v72_ = rhs275_
                lhs256_.f17 = rhs276_
            elif True:
                d_1480_v59_ = (d_1480_v59_).set(p0, True)
                (globalState).f18 = (d_1454_v37_) not in (((self).f22) + ((self).f22))
                d_1501_v76_: _dafny.Array
                nw254_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                d_1501_v76_ = nw254_
                d_1502_v77_: D11
                d_1502_v77_ = D11_DC28(d_1501_v76_)
                d_1502_v77_ = D11_DC28(d_1501_v76_)
                d_1503_v78_: _dafny.Array
                nw255_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_1503_v78_ = nw255_
                d_1504_v79_: _dafny.MultiSet
                d_1504_v79_ = _dafny.MultiSet([(d_1447_v30_).f21, (d_1447_v30_).f21])
                d_1505_v81_: D11
                d_1505_v81_ = D11_DC30(False, d_1504_v79_)
                d_1506_v82_: _dafny.Map
                d_1506_v82_ = _dafny.Map({(d_1447_v30_).f20: d_1505_v81_})
                d_1507_v83_: _dafny.Array
                nw256_ = _dafny.Array(None, 13)
                nw256_[int(0)] = (self).f21
                nw256_[int(1)] = (d_1447_v30_).f21
                nw256_[int(2)] = p1
                nw256_[int(3)] = (self).fm0(_dafny.Map({(d_1504_v79_).cardinality: (d_1447_v30_).f20}), (self).f22, _dafny.SeqWithoutIsStrInference([d_1451_v34_ for d_1508_i13_ in range(default__.abs(452))]), globalState)
                nw256_[int(4)] = (d_1447_v30_).f21
                nw256_[int(5)] = (self).fm0(d_1450_v33_, (self).f22, d_1452_v35_, globalState)
                nw256_[int(6)] = (self).f21
                nw256_[int(7)] = ((d_1480_v59_)[len((self).f22)] if (len((self).f22)) in (d_1480_v59_) else (d_1447_v30_).f21)
                nw256_[int(8)] = (d_1447_v30_).f21
                nw256_[int(9)] = (self).f21
                nw256_[int(10)] = (d_1447_v30_).f21
                nw256_[int(11)] = True
                def iife109_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in (d_1506_v82_).keys.Elements:
                        d_1509_v80_: int = compr_49_
                        if (d_1509_v80_) in (d_1506_v82_):
                            coll49_[default__.safeModuloInt(d_1509_v80_, (d_1447_v30_).f20)] = p0
                    return _dafny.Map(coll49_)
                nw256_[int(12)] = (self).fm0(iife109_()
                , (self).f22, d_1452_v35_, globalState)
                d_1507_v83_ = nw256_
                index259_ = default__.safeIndex(188, (d_1503_v78_).length(0))
                (d_1503_v78_)[index259_] = d_1507_v83_
                index260_ = default__.safeIndex(188, (d_1503_v78_).length(0))
                (d_1503_v78_)[index260_] = d_1507_v83_
                (globalState).f5 = default__.safeDivisionInt(787, (self).f20)
        d_1510_v84_: _dafny.Seq
        d_1510_v84_ = _dafny.SeqWithoutIsStrInference([d_1444_v29_])
        d_1511_v85_: D17
        d_1511_v85_ = D17_DC46(d_1510_v84_)
        r0 = ((d_1511_v85_).cf62) + ((d_1511_v85_).cf62)
        d_1512_v86_: D6
        d_1512_v86_ = D6_DC13(_dafny.SeqWithoutIsStrInference([d_1454_v37_ for d_1513_i14_ in range(default__.abs(-363))]))
        r1 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhqis"))) == (((d_1512_v86_).cf21) + ((self).f22))
        return r0, r1

    @property
    def f22(self):
        return self._f22
