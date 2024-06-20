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
        return default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False]))), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rakhhgk"))) if not(not(False)) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ry")))))

    @staticmethod
    def fm1(p0, p1, globalState):
        return (_dafny.CodePoint('d')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkmmphev")))

    @staticmethod
    def fm6(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_0_i0_ in range(default__.abs(110))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olpkpd"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1_i1_ in range(default__.abs(775))])))

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([D1_DC2(61, True, -978, 52) for d_2_i0_ in range(default__.abs(981))])) + ((_dafny.SeqWithoutIsStrInference([D1_DC2(len(_dafny.Set({False})), True, len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([False]))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ropqb")))), D1_DC2(422, False, -728, -268), D1_DC2(157, False, 267, -998)])) + (_dafny.SeqWithoutIsStrInference([D1_DC2(179, True, 207, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xv"))) for d_3_i1_ in range(default__.abs(-541))])))])))

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        if (False if not(True) else False):
            return _dafny.SeqWithoutIsStrInference([-925])
        elif True:
            return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([420, len(_dafny.Map({False: -368})), (0) - (-659), -693, len(_dafny.Map({False: 548}))])), len(_dafny.Map({(0) - ((_dafny.MultiSet([False, True, True])).cardinality): _dafny.Map({(0) - ((0) - (len(_dafny.Map({_dafny.CodePoint('p'): not(True)})))): 916})}))])

    @staticmethod
    def fm10(p0, globalState):
        return (_dafny.Set({True})) | (_dafny.Set({not(False), not(True)}))

    @staticmethod
    def fm11(p0, globalState):
        source0_ = (D16_DC45(_dafny.Map({len(_dafny.Set({True})): 142})) if True else D16_DC45(_dafny.Map({len(_dafny.Map({87: False})): 220})))
        if source0_.is_DC46:
            d_4___mcc_h0_ = source0_.cf79
            d_5_cf79_ = d_4___mcc_h0_
            return _dafny.CodePoint('d')
        elif source0_.is_DC45:
            d_6___mcc_h1_ = source0_.cf78
            d_7_cf78_ = d_6___mcc_h1_
            return _dafny.CodePoint('g')
        elif True:
            d_8___mcc_h2_ = source0_.cf80
            d_9_cf80_ = d_8___mcc_h2_
            return _dafny.CodePoint('a')

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gighne"))), 946})

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([855, -260, (0) - ((0) - (len(_dafny.Map({310: len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnvebqbov")))}))})))), len(_dafny.Map({509: 851})), 965])).Elements:
                d_11_v0_: int = compr_0_
                if (d_11_v0_) in (_dafny.SeqWithoutIsStrInference([855, -260, (0) - ((0) - (len(_dafny.Map({310: len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnvebqbov")))}))})))), len(_dafny.Map({509: 851})), 965])):
                    coll0_[(d_11_v0_) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(not(False))]))])))] = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))])
            return _dafny.Map(coll0_)
        return (_dafny.Map({756: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_10_i0_ in range(default__.abs(289))]))}))]))})) | (iife0_()
        )

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        return _dafny.Set({default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([False])), (0) - ((_dafny.MultiSet([False])).cardinality)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahpb"))), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_12_i0_ in range(default__.abs(202))]))), (0) - ((0) - (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), False]))).cardinality]))).intersection(_dafny.MultiSet([904]))).cardinality))})

    @staticmethod
    def fm15(p0, globalState):
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('u'), _dafny.CodePoint('f')])): True}))) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([627 for d_13_i0_ in range(default__.abs(816))])): False}))

    @staticmethod
    def fm16(p0, p1, globalState):
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([814, (0) - ((0) - (-505)), (0) - (len(_dafny.Map({704: 455}))), (0) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_14_i0_ in range(default__.abs(-308))])])).cardinality), len(_dafny.Map({not(False): False}))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Set({893})), 646])}))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: D3
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([D3_DC8()])).Elements:
                d_15_v0_: D3 = compr_1_
                if (d_15_v0_) in (_dafny.SeqWithoutIsStrInference([D3_DC8()])):
                    coll1_[d_15_v0_] = ((0) - (len(_dafny.Set({True})))) + (558)
            return _dafny.Map(coll1_)
        return iife1_()
        

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return D1_DC1(230)

    @staticmethod
    def fm19(p0, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxerujrc"))]) if False else _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwiu")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fghlrjlu"))]))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efts"))]))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm21(p0, p1, globalState):
        return ((_dafny.MultiSet([True])).intersection(_dafny.MultiSet([True]))) | ((_dafny.MultiSet([True])) - (_dafny.MultiSet([False])))

    @staticmethod
    def fm22(globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.MultiSet
            for compr_2_ in (_dafny.Map({_dafny.MultiSet([len(_dafny.Set({558, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_16_i0_ in range(default__.abs(417))]))}))]): len(_dafny.SeqWithoutIsStrInference([False, False]))})).keys.Elements:
                d_17_v0_: _dafny.MultiSet = compr_2_
                if (d_17_v0_) in (_dafny.Map({_dafny.MultiSet([len(_dafny.Set({558, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_16_i0_ in range(default__.abs(417))]))}))]): len(_dafny.SeqWithoutIsStrInference([False, False]))})):
                    coll2_[d_17_v0_] = not(True)
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.MultiSet
            for compr_3_ in (_dafny.Map({_dafny.MultiSet([len(_dafny.Map({57: True})), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}))]): 656})).keys.Elements:
                d_18_v1_: _dafny.MultiSet = compr_3_
                if (d_18_v1_) in (_dafny.Map({_dafny.MultiSet([len(_dafny.Map({57: True})), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}))]): 656})):
                    coll3_[d_18_v1_] = True
            return _dafny.Map(coll3_)
        return (iife2_()
        ) | (iife3_()
        )

    @staticmethod
    def fm23(p0, p1, p2, globalState):
        return (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([False]))})) | ((_dafny.Map({not(False): 911})) | (_dafny.Map({False: 559})))

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return (_dafny.Set({not(True)})).intersection(_dafny.Set({True, False, False, False}))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return D1_DC2(len(_dafny.Map({False: 957})), (not(False)) and (False), (568) - (800), 154)

    @staticmethod
    def fm26(globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in (_dafny.Set({len(_dafny.Map({False: 837}))})).Elements:
                d_19_v0_: int = compr_4_
                if (d_19_v0_) in (_dafny.Set({len(_dafny.Map({False: 837}))})):
                    coll4_ = coll4_.union(_dafny.Set([default__.safeModuloInt(d_19_v0_, 866)]))
            return _dafny.Set(coll4_)
        return ((_dafny.SeqWithoutIsStrInference([D1_DC2(len(_dafny.Map({True: _dafny.CodePoint('j')})), False, len(_dafny.Set({True, True, not(False), False, True})), 15), D1_DC2(len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xc")): len(_dafny.Set({True}))})), not(True), 679, 163)])) + (_dafny.SeqWithoutIsStrInference([D1_DC2(len(_dafny.Set({iife4_()
, _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xydlmxbg")))})})), False, 683, (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxrqqpsh"))])).cardinality), D1_DC2(len(_dafny.SeqWithoutIsStrInference([31, 106])), True, 860, -392), D1_DC2(116, True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mn"))), -325)]))) + (_dafny.SeqWithoutIsStrInference([D1_DC2(627, not(False), 837, 429)]))

    @staticmethod
    def fm27(globalState):
        return D3_DC10()

    @staticmethod
    def fm28(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfu"))

    @staticmethod
    def fm29(p0, p1, globalState):
        return _dafny.MultiSet([default__.safeModuloInt(897, -821)])

    @staticmethod
    def fm30(globalState):
        if True:
            return (_dafny.Map({False: True})) | (_dafny.Map({False: False}))
        elif True:
            return _dafny.Map({True: not(False)})

    @staticmethod
    def fm31(globalState):
        source1_ = D12_DC32(-862, True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knta"))))
        if source1_.is_DC32:
            d_20___mcc_h0_ = source1_.cf53
            d_21___mcc_h1_ = source1_.cf54
            d_22___mcc_h2_ = source1_.cf55
            d_23_cf55_ = d_22___mcc_h2_
            d_24_cf54_ = d_21___mcc_h1_
            d_25_cf53_ = d_20___mcc_h0_
            return _dafny.CodePoint('h')
        elif source1_.is_DC31:
            d_26___mcc_h3_ = source1_.cf52
            d_27_cf52_ = d_26___mcc_h3_
            return _dafny.CodePoint('p')
        elif True:
            d_28___mcc_h4_ = source1_.cf56
            d_29_cf56_ = d_28___mcc_h4_
            return _dafny.CodePoint('c')

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opsm"))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmbalwqgj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uetnpshyt"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_30_i0_ in range(default__.abs(-573))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))])))

    @staticmethod
    def fm35(globalState):
        if (True if True else not(not(True))):
            return _dafny.CodePoint('c')
        elif True:
            return _dafny.CodePoint('b')

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return (_dafny.Map({_dafny.Map({True: len(_dafny.Map({901: -213}))}): -297})) | (_dafny.Map({_dafny.Map({False: -323}): len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxungteu"))]))}))

    @staticmethod
    def fm37(p0, p1, globalState):
        return _dafny.Set({_dafny.CodePoint('l'), _dafny.CodePoint('t')})

    @staticmethod
    def fm38(p0, p1, globalState):
        return D1_DC2((0) - (len((_dafny.Map({False: -290})) | (_dafny.Map({not(not(not(not(False)))): 519})))), True, len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pntnad")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vo"))})), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvbi"))), 448])).cardinality)

    @staticmethod
    def fm39(globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(37, 25):
                d_31_v0_: int = compr_5_
                if ((37) <= (d_31_v0_)) and ((d_31_v0_) < (25)):
                    coll5_ = coll5_.union(_dafny.Set([default__.safeDivisionInt(d_31_v0_, (_dafny.MultiSet([207])).cardinality)]))
            return _dafny.Set(coll5_)
        return (_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D12_DC32(len(_dafny.Map({False: _dafny.Map({False: True})})), False, len(iife5_()
))).cf54, False]))})).intersection(_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))}))

    @staticmethod
    def fm40(globalState):
        return _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_32_i0_ in range(default__.abs(213))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rp")))])

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([False])

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        return (_dafny.MultiSet([False, True, False]))

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([_dafny.Set({True, False, not(False)})]) if not(False) else _dafny.MultiSet([_dafny.Set({False, True})]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({True, True, False, True}), _dafny.Set({not(False)}), _dafny.Set({False, False, False, not(False)}), _dafny.Set({True}), _dafny.Set({not(not(False))})])))

    @staticmethod
    def fm44(globalState):
        source2_ = D14_DC38(_dafny.CodePoint('p'), not(True), _dafny.SeqWithoutIsStrInference([(D7_DC20(True, _dafny.CodePoint('k'))).cf36]))
        if source2_.is_DC38:
            d_33___mcc_h0_ = source2_.cf67
            d_34___mcc_h1_ = source2_.cf68
            d_35___mcc_h2_ = source2_.cf69
            d_36_cf69_ = d_35___mcc_h2_
            d_37_cf68_ = d_34___mcc_h1_
            d_38_cf67_ = d_33___mcc_h0_
            return D7_DC20(False, d_38_cf67_)
        elif source2_.is_DC39:
            d_39___mcc_h3_ = source2_.cf70
            d_40___mcc_h4_ = source2_.cf71
            d_41_cf71_ = d_40___mcc_h4_
            d_42_cf70_ = d_39___mcc_h3_
            if False:
                return D7_DC20(True, _dafny.CodePoint('h'))
            elif True:
                return D7_DC20(False, _dafny.CodePoint('t'))
        elif True:
            d_43___mcc_h5_ = source2_.cf66
            d_44_cf66_ = d_43___mcc_h5_
            return D7_DC20(False, _dafny.CodePoint('g'))

    @staticmethod
    def fm45(globalState):
        return ((_dafny.SeqWithoutIsStrInference([D1_DC3(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwhk")), -798)])) + (_dafny.SeqWithoutIsStrInference([D1_DC3(not(False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cc")), (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsv"))))))]))) + (_dafny.SeqWithoutIsStrInference([D1_DC3(not(True), (D13_DC36(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_45_i0_ in range(default__.abs(0))]), 369, len(_dafny.SeqWithoutIsStrInference([False])), (_dafny.MultiSet([True, False])).cardinality)).cf62, 246), D1_DC3(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")), 89)]))

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([945 for d_46_i0_ in range(default__.abs(703))])).Elements:
                d_47_v0_: int = compr_6_
                if (d_47_v0_) in (_dafny.SeqWithoutIsStrInference([945 for d_46_i0_ in range(default__.abs(703))])):
                    coll6_[default__.safeModuloInt(d_47_v0_, len(_dafny.SeqWithoutIsStrInference([655, 235])))] = (0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwex"))): 493})))
            return _dafny.Map(coll6_)
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(706, 867):
                d_49_v1_: int = compr_7_
                if ((706) <= (d_49_v1_)) and ((d_49_v1_) < (867)):
                    coll7_[(d_49_v1_) * ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([941]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True)]))).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_50_i2_ in range(default__.abs(420))]))]), _dafny.SeqWithoutIsStrInference([284])]))).cardinality)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkrufmho"))
            return _dafny.Map(coll7_)
        return ((_dafny.Map({len(iife6_()
        ): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_48_i1_ in range(default__.abs(673))])})) | (iife7_()
        )) | (_dafny.Map({(_dafny.MultiSet([True, False, False])).cardinality: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_51_i3_ in range(default__.abs(492))])}))

    @staticmethod
    def fm47(globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: _dafny.Seq
            for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, False]))])).cardinality])])).Elements:
                d_52_v0_: _dafny.Seq = compr_8_
                if (d_52_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, False]))])).cardinality])])):
                    coll8_[d_52_v0_] = 180
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in (_dafny.Set({len(_dafny.Map({531: True}))})).Elements:
                d_54_v1_: int = compr_9_
                if (d_54_v1_) in (_dafny.Set({len(_dafny.Map({531: True}))})):
                    coll9_ = coll9_.union(_dafny.Set([(d_54_v1_) * (801)]))
            return _dafny.Set(coll9_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([936]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ti")))})) | (iife8_()
        )) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([347, len(_dafny.Set({-530}))]): len(_dafny.Set({_dafny.MultiSet([True])}))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_53_i0_ in range(default__.abs(100))])), len(_dafny.Map({(_dafny.MultiSet([len(iife9_()
        )])).cardinality: not(False)}))]): 660})))

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ks"))): (0) - ((len(_dafny.Set({-433, 647, 919}))) - (852))})

    @staticmethod
    def m0(p0, globalState):
        d_55_v1_: int
        d_55_v1_ = -50
        d_56_v2_: _dafny.Map
        d_56_v2_ = _dafny.Map({d_55_v1_: d_55_v1_})
        d_57_v3_: _dafny.Map
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in (_dafny.SeqWithoutIsStrInference([d_55_v1_])).Elements:
                d_58_v0_: int = compr_10_
                if (d_58_v0_) in (_dafny.SeqWithoutIsStrInference([d_55_v1_])):
                    coll10_[(d_58_v0_) * (d_55_v1_)] = p0
            return _dafny.Map(coll10_)
        d_57_v3_ = _dafny.Map({iife10_()
        : (d_56_v2_) | ((d_56_v2_).set(d_55_v1_, d_55_v1_))})
        d_59_v4_: _dafny.MultiSet
        d_59_v4_ = _dafny.MultiSet([d_55_v1_])
        d_60_v5_: _dafny.Map
        d_60_v5_ = _dafny.Map({d_55_v1_: default__.fm1(d_59_v4_, d_55_v1_, globalState)})
        d_61_v6_: D18
        d_61_v6_ = D18_DC52(d_57_v3_)
        d_57_v3_ = ((_dafny.Map({d_60_v5_: d_56_v2_})) | ((d_61_v6_).cf88)) | (((d_57_v3_).set(d_60_v5_, d_56_v2_)) | (d_57_v3_))
        d_55_v1_ = d_55_v1_
        if p0:
            d_62_v7_: _dafny.Seq
            d_62_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlxanf"))
            d_55_v1_ = default__.safeDivisionInt(((d_56_v2_)[(0) - ((_dafny.MultiSet([p0, p0])).cardinality)] if ((0) - ((_dafny.MultiSet([p0, p0])).cardinality)) in (d_56_v2_) else d_55_v1_), len(d_62_v7_))
            d_63_v8_: _dafny.Seq
            d_63_v8_ = _dafny.SeqWithoutIsStrInference([d_55_v1_])
            d_55_v1_ = ((_dafny.MultiSet(d_63_v8_)).cardinality) * (d_55_v1_)
            d_64_v9_: _dafny.Map
            d_64_v9_ = _dafny.Map({p0: (d_55_v1_) != (d_55_v1_)})
            (globalState).f2 = ((d_64_v9_)[p0] if (p0) in (d_64_v9_) else default__.fm1(d_59_v4_, d_55_v1_, globalState))
            d_65_v10_: _dafny.Seq
            d_65_v10_ = _dafny.SeqWithoutIsStrInference([p0])
            d_65_v10_ = d_65_v10_
            d_66_v11_: str
            d_66_v11_ = _dafny.CodePoint('f')
            d_67_v12_: T0
            nw0_ = C1()
            nw0_.ctor__(_dafny.SeqWithoutIsStrInference([D1_DC2((0) - (len(_dafny.Map({p0: p0}))), p0, len(_dafny.Set({(d_59_v4_).cardinality})), d_55_v1_) for d_68_i0_ in range(default__.abs(644))]))
            d_67_v12_ = nw0_
            d_69_v13_: D17
            d_69_v13_ = D17_DC50(d_66_v11_, p0, d_62_v7_, d_55_v1_, d_67_v12_)
            d_62_v7_ = ((d_69_v13_).cf84).set(default__.safeIndex(d_55_v1_, len((d_69_v13_).cf84)), d_66_v11_)
        elif True:
            d_70_v14_: D1
            d_70_v14_ = D1_DC2(d_55_v1_, (d_55_v1_) <= (d_55_v1_), d_55_v1_, d_55_v1_)
            d_71_v15_: _dafny.MultiSet
            d_71_v15_ = _dafny.MultiSet([p0, p0])
            d_72_v16_: C4
            nw1_ = C4()
            nw1_.ctor__(p0)
            d_72_v16_ = nw1_
            d_73_v17_: _dafny.Map
            d_73_v17_ = _dafny.Map({d_72_v16_: p0})
            d_70_v14_ = D1_DC2((d_55_v1_) * (d_55_v1_), p0, d_55_v1_, ((d_71_v15_).cardinality) + (len(d_73_v17_)))
            (globalState).f2 = default__.fm1(d_59_v4_, default__.safeModuloInt(d_55_v1_, d_55_v1_), globalState)
            d_74_v18_: _dafny.Seq
            d_74_v18_ = _dafny.SeqWithoutIsStrInference([d_70_v14_])
            d_75_v19_: C1
            nw2_ = C1()
            nw2_.ctor__(d_74_v18_)
            d_75_v19_ = nw2_
            d_76_v20_: _dafny.Seq
            d_76_v20_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([p0, p0, p0, p0])).cardinality, d_55_v1_, 419, len((_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex(d_55_v1_, len(_dafny.SeqWithoutIsStrInference([p0]))), True))])
            d_77_v21_: _dafny.Map
            d_77_v21_ = _dafny.Map({(d_76_v20_) < (d_76_v20_): d_55_v1_})
            d_78_v22_: _dafny.Seq
            d_78_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqwwfaqw"))
            d_79_v23_: C5
            nw3_ = C5()
            nw3_.ctor__(d_55_v1_, p0)
            d_79_v23_ = nw3_
            d_80_v24_: _dafny.Seq
            d_80_v24_ = _dafny.SeqWithoutIsStrInference([d_79_v23_, d_79_v23_, d_79_v23_])
            rhs0_ = p0
            rhs1_ = d_77_v21_
            rhs2_ = p0
            rhs3_ = d_78_v22_
            rhs4_ = len(d_80_v24_)
            lhs0_ = globalState
            lhs1_ = globalState
            lhs0_.f2 = rhs0_
            d_77_v21_ = rhs1_
            lhs1_.f2 = rhs2_
            d_78_v22_ = rhs3_
            d_55_v1_ = rhs4_
            d_81_v25_: _dafny.Set
            d_81_v25_ = _dafny.Set({d_76_v20_})
            d_82_v26_: _dafny.Seq
            d_82_v26_ = _dafny.SeqWithoutIsStrInference([p0, p0, False, p0, p0])
            d_83_v27_: _dafny.Seq
            d_83_v27_ = _dafny.SeqWithoutIsStrInference([not(p0), (d_82_v26_)[default__.safeIndex((d_79_v23_).f10, len(d_82_v26_))], True])
            d_84_v28_: D11
            d_84_v28_ = D11_DC30()
            if not((d_79_v23_).fm5((d_81_v25_) - (d_81_v25_), _dafny.Map({(d_79_v23_).fm4(p0, p0, p0, p0, globalState): d_82_v26_}), len((_dafny.SeqWithoutIsStrInference([d_84_v28_]) if p0 else _dafny.SeqWithoutIsStrInference([d_84_v28_]))), globalState)):
                d_85_v29_: str
                d_85_v29_ = _dafny.CodePoint('x')
                d_86_v30_: _dafny.Map
                d_86_v30_ = _dafny.Map({p0: d_85_v29_})
                d_87_v31_: _dafny.Array
                nw4_ = _dafny.Array(None, 23)
                nw4_[int(0)] = d_85_v29_
                nw4_[int(1)] = ((d_86_v30_)[p0] if (p0) in (d_86_v30_) else d_85_v29_)
                nw4_[int(2)] = d_85_v29_
                nw4_[int(3)] = d_85_v29_
                nw4_[int(4)] = d_85_v29_
                nw4_[int(5)] = (d_85_v29_ if not(not(not(p0))) else d_85_v29_)
                nw4_[int(6)] = d_85_v29_
                nw4_[int(7)] = d_85_v29_
                nw4_[int(8)] = d_85_v29_
                nw4_[int(9)] = d_85_v29_
                nw4_[int(10)] = d_85_v29_
                nw4_[int(11)] = d_85_v29_
                nw4_[int(12)] = d_85_v29_
                nw4_[int(13)] = (d_85_v29_ if p0 else d_85_v29_)
                nw4_[int(14)] = d_85_v29_
                nw4_[int(15)] = _dafny.CodePoint('s')
                nw4_[int(16)] = _dafny.CodePoint('d')
                nw4_[int(17)] = d_85_v29_
                nw4_[int(18)] = _dafny.CodePoint('i')
                nw4_[int(19)] = _dafny.CodePoint('r')
                nw4_[int(20)] = d_85_v29_
                nw4_[int(21)] = d_85_v29_
                nw4_[int(22)] = (d_78_v22_)[default__.safeIndex((d_79_v23_).f10, len(d_78_v22_))]
                d_87_v31_ = nw4_
                index0_ = default__.safeIndex(554, (d_87_v31_).length(0))
                (d_87_v31_)[index0_] = d_85_v29_
                index1_ = default__.safeIndex(554, (d_87_v31_).length(0))
                (d_87_v31_)[index1_] = d_85_v29_
                d_76_v20_ = d_76_v20_
                d_88_v32_: C2
                nw5_ = C2()
                nw5_.ctor__((d_79_v23_).f10, not (p0) or (p0))
                d_88_v32_ = nw5_
                d_89_v33_: _dafny.Array
                nw6_ = _dafny.Array(False, 3)
                d_89_v33_ = nw6_
                index2_ = default__.safeIndex(153, (d_89_v33_).length(0))
                (d_89_v33_)[index2_] = ((d_76_v20_)[default__.safeIndex((d_79_v23_).f10, len(d_76_v20_))]) >= ((d_79_v23_).f10)
                index3_ = default__.safeIndex(153, (d_89_v33_).length(0))
                (d_89_v33_)[index3_] = False
                (globalState).f2 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jb"))) == (default__.fm6(d_83_v27_, globalState))
            elif True:
                d_90_v34_: _dafny.Array
                nw7_ = _dafny.Array(int(0), 12)
                d_90_v34_ = nw7_
                d_91_v35_: _dafny.Set
                d_91_v35_ = _dafny.Set({d_90_v34_, d_90_v34_})
                d_92_v36_: str
                d_92_v36_ = _dafny.CodePoint('m')
                d_93_v37_: _dafny.Map
                d_93_v37_ = _dafny.Map({(_dafny.Set({d_90_v34_})) | (d_91_v35_): (d_78_v22_).set(default__.safeIndex(d_55_v1_, len(d_78_v22_)), d_92_v36_)})
                d_93_v37_ = (d_93_v37_).set((d_91_v35_).intersection(d_91_v35_), d_78_v22_)
                d_94_v38_: C0
                nw8_ = C0()
                nw8_.ctor__(d_81_v25_, ((d_79_v23_).f10) - ((d_79_v23_).f10))
                d_94_v38_ = nw8_
                d_95_v39_: _dafny.Seq
                d_95_v39_ = _dafny.SeqWithoutIsStrInference([(d_90_v34_)])
                rhs5_ = p0
                rhs6_ = (d_95_v39_)[default__.safeIndex((d_79_v23_).f10, len(d_95_v39_))]
                rhs7_ = d_76_v20_
                rhs8_ = d_94_v38_
                rhs9_ = -96
                lhs2_ = globalState
                lhs3_ = d_94_v38_
                lhs2_.f2 = rhs5_
                d_90_v34_ = rhs6_
                d_76_v20_ = rhs7_
                d_94_v38_ = rhs8_
                lhs3_.f12 = rhs9_
                (globalState).f2 = p0
                rhs10_ = ((d_79_v23_).f10) - ((603) + (d_94_v38_.f12))
                rhs11_ = d_55_v1_
                lhs4_ = d_94_v38_
                d_55_v1_ = rhs10_
                lhs4_.f12 = rhs11_
                d_55_v1_ = (0) - ((d_79_v23_).f10)
        d_96_v40_: _dafny.Set
        d_96_v40_ = _dafny.Set({-54, d_55_v1_, d_55_v1_, d_55_v1_, d_55_v1_})
        d_97_i1_: int
        d_97_i1_ = 0
        with _dafny.label("0"):
            while default__.fm1(d_59_v4_, default__.safeModuloInt(len(d_96_v40_), d_55_v1_), globalState):
                with _dafny.c_label("0"):
                    if (d_97_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_97_i1_ = (d_97_i1_) + (1)
                    d_98_v41_: _dafny.Array
                    def lambda0_(d_99_p0_):
                        def lambda1_(d_100_i2_):
                            return d_99_p0_

                        return lambda1_

                    init0_ = lambda0_(p0)
                    nw9_ = _dafny.Array(None, 15)
                    for i0_0_ in range(nw9_.length(0)):
                        nw9_[i0_0_] = init0_(i0_0_)
                    d_98_v41_ = nw9_
                    index4_ = default__.safeIndex(901, (d_98_v41_).length(0))
                    (d_98_v41_)[index4_] = p0
                    index5_ = default__.safeIndex(901, (d_98_v41_).length(0))
                    (d_98_v41_)[index5_] = not(p0)
                    (globalState).f2 = p0
                    d_55_v1_ = (d_55_v1_) + (d_55_v1_)
                    index6_ = default__.safeIndex(901, (d_98_v41_).length(0))
                    (d_98_v41_)[index6_] = p0
                    pass
            pass
        if not (p0) or (p0):
            d_101_v42_: _dafny.Set
            d_101_v42_ = d_96_v40_
            d_102_v44_: _dafny.Array
            nw10_ = _dafny.Array(None, 4)
            nw10_[int(0)] = d_96_v40_
            nw10_[int(1)] = (d_101_v42_)
            nw10_[int(2)] = d_96_v40_
            def iife11_():
                coll11_ = _dafny.Set()
                compr_11_: int
                for compr_11_ in (d_56_v2_).keys.Elements:
                    d_103_v43_: int = compr_11_
                    if (d_103_v43_) in (d_56_v2_):
                        coll11_ = coll11_.union(_dafny.Set([(d_103_v43_) * (len(_dafny.Set({len(_dafny.Set({92, (_dafny.MultiSet([False, False])).cardinality})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oaq"))), len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, False})), len(_dafny.Set({125}))]))): len(_dafny.Map({True: 290}))})), 623, len(_dafny.Map({False: True}))})))]))
                return _dafny.Set(coll11_)
            nw10_[int(3)] = iife11_()
            
            d_102_v44_ = nw10_
            d_104_v45_: C3
            nw11_ = C3()
            nw11_.ctor__(p0, d_102_v44_, not(p0))
            d_104_v45_ = nw11_
            d_55_v1_ = (0) - ((0) - (default__.fm0(globalState)))
            d_55_v1_ = (d_55_v1_) * (d_55_v1_)
            (globalState).f2 = (not (p0) or (p0) if p0 else p0)
            (globalState).f2 = not((d_104_v45_).f15)
        elif True:
            d_105_v46_: _dafny.Map
            d_105_v46_ = _dafny.Map({d_96_v40_: d_55_v1_})
            d_105_v46_ = (d_105_v46_) | (d_105_v46_)
            d_106_v47_: _dafny.Seq
            d_106_v47_ = _dafny.SeqWithoutIsStrInference([p0])
            d_107_v48_: _dafny.Map
            d_107_v48_ = _dafny.Map({p0: d_106_v47_})
            d_108_v49_: _dafny.MultiSet
            d_108_v49_ = _dafny.MultiSet([p0])
            d_109_v50_: _dafny.Array
            nw12_ = _dafny.Array(None, 10)
            nw12_[int(0)] = _dafny.MultiSet(((d_107_v48_)[p0] if (p0) in (d_107_v48_) else _dafny.SeqWithoutIsStrInference([p0])))
            nw12_[int(1)] = (d_108_v49_) - (_dafny.MultiSet([p0]))
            nw12_[int(2)] = (d_108_v49_).set(p0, default__.abs(d_55_v1_))
            nw12_[int(3)] = (d_108_v49_).intersection(d_108_v49_)
            nw12_[int(4)] = d_108_v49_
            nw12_[int(5)] = d_108_v49_
            nw12_[int(6)] = (d_108_v49_) | (d_108_v49_)
            nw12_[int(7)] = _dafny.MultiSet((d_106_v47_ if p0 else d_106_v47_))
            nw12_[int(8)] = d_108_v49_
            nw12_[int(9)] = d_108_v49_
            d_109_v50_ = nw12_
            index7_ = default__.safeIndex(402, (d_109_v50_).length(0))
            (d_109_v50_)[index7_] = d_108_v49_
            index8_ = default__.safeIndex(402, (d_109_v50_).length(0))
            (d_109_v50_)[index8_] = d_108_v49_
            d_110_v51_: _dafny.Set
            d_110_v51_ = _dafny.Set({p0, p0})
            d_111_v52_: D1
            d_111_v52_ = D1_DC2(d_55_v1_, p0, len(d_110_v51_), d_55_v1_)
            d_112_v53_: _dafny.Seq
            d_112_v53_ = _dafny.SeqWithoutIsStrInference([d_111_v52_])
            d_113_v54_: T0
            nw13_ = C1()
            nw13_.ctor__(d_112_v53_)
            d_113_v54_ = nw13_
            d_114_v55_: _dafny.Set
            d_114_v55_ = _dafny.Set({d_113_v54_})
            d_55_v1_ = len(d_114_v55_)
            d_115_v56_: D10
            d_115_v56_ = D10_DC27(_dafny.Map({p0: d_55_v1_}))
            d_116_v57_: _dafny.Map
            d_116_v57_ = _dafny.Map({True: d_55_v1_})
            pat_let_tv0_ = d_116_v57_
            pat_let_tv1_ = d_116_v57_
            def iife12_(_pat_let0_0):
                def iife13_(d_117_dt__update__tmp_h0_):
                    def iife14_(_pat_let1_0):
                        def iife15_(d_118_dt__update_hcf50_h0_):
                            return D10_DC27(d_118_dt__update_hcf50_h0_)
                        return iife15_(_pat_let1_0)
                    return iife14_((pat_let_tv0_) | (pat_let_tv1_))
                return iife13_(_pat_let0_0)
            source3_ = iife12_(d_115_v56_)
            if source3_.is_DC28:
                d_119_v58_: _dafny.Array
                def lambda2_(d_120_i3_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvtkev"))

                init1_ = lambda2_
                nw14_ = _dafny.Array(None, 14)
                for i0_1_ in range(nw14_.length(0)):
                    nw14_[i0_1_] = init1_(i0_1_)
                d_119_v58_ = nw14_
                index9_ = default__.safeIndex(90, (d_119_v58_).length(0))
                (d_119_v58_)[index9_] = (D13_DC36(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_121_i4_ in range(default__.abs(579))]), 327, d_55_v1_, d_55_v1_)).cf62
                d_122_v59_: str
                d_122_v59_ = _dafny.CodePoint('m')
                d_123_v60_: _dafny.Map
                d_123_v60_ = _dafny.Map({d_55_v1_: d_113_v54_})
                d_124_v61_: _dafny.Map
                d_124_v61_ = _dafny.Map({len(d_123_v60_): d_122_v59_})
                index10_ = default__.safeIndex(90, (d_119_v58_).length(0))
                rhs12_ = (len((_dafny.Map({d_55_v1_: d_122_v59_})) | (d_124_v61_)) if False else d_55_v1_)
                rhs13_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_125_i5_ in range(default__.abs(158))])
                rhs14_ = default__.fm1(d_59_v4_, ((d_59_v4_).cardinality if ((d_60_v5_)[d_55_v1_] if (d_55_v1_) in (d_60_v5_) else p0) else d_55_v1_), globalState)
                lhs5_ = d_119_v58_
                lhs6_ = default__.safeIndex(90, (d_119_v58_).length(0))
                lhs7_ = globalState
                d_55_v1_ = rhs12_
                lhs5_[lhs6_] = rhs13_
                lhs7_.f2 = rhs14_
                d_55_v1_ = d_55_v1_
                d_110_v51_ = d_110_v51_
                (globalState).f2 = (not(p0) if p0 else p0)
            elif True:
                d_126___mcc_h0_ = source3_.cf50
                d_127_cf50_ = d_126___mcc_h0_
                d_128_v62_: C4
                nw15_ = C4()
                nw15_.ctor__(True)
                d_128_v62_ = nw15_
                d_129_v63_: int
                d_130_v64_: D4
                d_131_v65_: _dafny.MultiSet
                d_132_v66_: bool
                out0_: int
                out1_: D4
                out2_: _dafny.MultiSet
                out3_: bool
                out0_, out1_, out2_, out3_ = (d_128_v62_).m7(d_55_v1_, d_55_v1_, False, globalState)
                d_129_v63_ = out0_
                d_130_v64_ = out1_
                d_131_v65_ = out2_
                d_132_v66_ = out3_
                (globalState).f2 = (-737) != ((d_108_v49_).cardinality)
                d_133_v67_: int
                d_134_v68_: D4
                d_135_v69_: _dafny.MultiSet
                d_136_v70_: bool
                out4_: int
                out5_: D4
                out6_: _dafny.MultiSet
                out7_: bool
                out4_, out5_, out6_, out7_ = (d_128_v62_).m7(d_129_v63_, d_55_v1_, p0, globalState)
                d_133_v67_ = out4_
                d_134_v68_ = out5_
                d_135_v69_ = out6_
                d_136_v70_ = out7_
            d_137_v71_: D16
            d_137_v71_ = D16_DC46(d_55_v1_)
            d_138_v72_: _dafny.Map
            d_138_v72_ = _dafny.Map({d_137_v71_: p0})
            rhs15_ = (0) - ((0) - (d_55_v1_))
            rhs16_ = (d_55_v1_) * ((default__.fm0(globalState) if p0 else d_55_v1_))
            rhs17_ = default__.fm1(d_59_v4_, d_55_v1_, globalState)
            rhs18_ = (d_138_v72_ if (default__.fm0(globalState)) <= (d_55_v1_) else d_138_v72_)
            lhs8_ = globalState
            d_55_v1_ = rhs15_
            d_55_v1_ = rhs16_
            lhs8_.f2 = rhs17_
            d_138_v72_ = rhs18_
        d_139_i6_: int
        d_139_i6_ = 0
        with _dafny.label("1"):
            while p0:
                with _dafny.c_label("1"):
                    if (d_139_i6_) >= (100):
                        raise _dafny.Break("1")
                    d_139_i6_ = (d_139_i6_) + (1)
                    d_140_v73_: _dafny.Seq
                    d_140_v73_ = _dafny.SeqWithoutIsStrInference([False, p0])
                    d_141_v74_: T2
                    nw16_ = C5()
                    nw16_.ctor__(d_55_v1_, p0)
                    d_141_v74_ = nw16_
                    d_142_v75_: _dafny.Map
                    d_142_v75_ = _dafny.Map({len(d_140_v73_): d_141_v74_})
                    d_143_v76_: _dafny.Seq
                    d_143_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mivytavqf"))
                    d_144_v77_: C4
                    nw17_ = C4()
                    nw17_.ctor__(p0)
                    d_144_v77_ = nw17_
                    d_145_v78_: _dafny.Seq
                    d_145_v78_ = _dafny.SeqWithoutIsStrInference([d_144_v77_, d_144_v77_])
                    d_146_v79_: _dafny.Seq
                    d_146_v79_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_144_v77_]), d_145_v78_, _dafny.SeqWithoutIsStrInference([d_144_v77_, d_144_v77_, d_144_v77_])])
                    d_147_v80_: _dafny.Map
                    d_147_v80_ = _dafny.Map({d_143_v76_: d_55_v1_})
                    d_148_v81_: _dafny.Seq
                    d_148_v81_ = _dafny.SeqWithoutIsStrInference([d_143_v76_])
                    d_149_v82_: _dafny.Array
                    nw18_ = _dafny.Array(None, 29)
                    nw18_[int(0)] = len(_dafny.Set({d_140_v73_, d_140_v73_}))
                    nw18_[int(1)] = 616
                    nw18_[int(2)] = (d_55_v1_) * (d_55_v1_)
                    nw18_[int(3)] = len(((d_142_v75_).set((d_59_v4_).cardinality, d_141_v74_)) | (d_142_v75_))
                    nw18_[int(4)] = d_55_v1_
                    nw18_[int(5)] = (0) - ((0) - ((d_55_v1_) * (d_55_v1_)))
                    nw18_[int(6)] = (d_55_v1_) + (d_55_v1_)
                    nw18_[int(7)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_150_i7_ in range(default__.abs(569))])) + (d_143_v76_))
                    nw18_[int(8)] = default__.safeDivisionInt(len(default__.fm6(d_140_v73_, globalState)), len(d_146_v79_))
                    nw18_[int(9)] = d_55_v1_
                    nw18_[int(10)] = d_55_v1_
                    nw18_[int(11)] = d_55_v1_
                    nw18_[int(12)] = len(d_143_v76_)
                    nw18_[int(13)] = (d_55_v1_) - (d_55_v1_)
                    nw18_[int(14)] = 989
                    nw18_[int(15)] = 912
                    nw18_[int(16)] = (d_55_v1_) - (((d_147_v80_)[(d_148_v81_)[default__.safeIndex(d_55_v1_, len(d_148_v81_))]] if ((d_148_v81_)[default__.safeIndex(d_55_v1_, len(d_148_v81_))]) in (d_147_v80_) else 967))
                    nw18_[int(17)] = d_55_v1_
                    nw18_[int(18)] = 530
                    nw18_[int(19)] = d_55_v1_
                    nw18_[int(20)] = 322
                    nw18_[int(21)] = d_55_v1_
                    nw18_[int(22)] = (len(default__.fm20(p0, not((d_141_v74_).f9), -541, d_143_v76_, globalState))) * (d_55_v1_)
                    nw18_[int(23)] = default__.fm0(globalState)
                    nw18_[int(24)] = (0) - (d_55_v1_)
                    nw18_[int(25)] = (d_55_v1_) * (d_55_v1_)
                    nw18_[int(26)] = d_55_v1_
                    nw18_[int(27)] = d_55_v1_
                    nw18_[int(28)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))
                    d_149_v82_ = nw18_
                    d_151_v83_: str
                    d_151_v83_ = _dafny.CodePoint('a')
                    d_152_v84_: _dafny.Set
                    d_152_v84_ = _dafny.Set({d_151_v83_, d_151_v83_})
                    d_153_v85_: _dafny.Map
                    d_153_v85_ = _dafny.Map({p0: d_152_v84_})
                    rhs19_ = d_149_v82_
                    rhs20_ = (d_153_v85_) == (_dafny.Map({not((d_141_v74_).f9): d_152_v84_}))
                    lhs9_ = globalState
                    d_149_v82_ = rhs19_
                    lhs9_.f2 = rhs20_
                    d_154_v86_: _dafny.Map
                    d_154_v86_ = _dafny.Map({(d_141_v74_).f9: d_141_v74_})
                    d_55_v1_ = (len((d_154_v86_) | (d_154_v86_))) - (d_55_v1_)
                    d_155_v87_: _dafny.MultiSet
                    d_155_v87_ = _dafny.MultiSet([d_56_v2_, d_56_v2_])
                    d_55_v1_ = default__.safeModuloInt(default__.safeDivisionInt(d_55_v1_, len(d_143_v76_)), ((d_155_v87_)[d_56_v2_] if (d_56_v2_) in (d_155_v87_) else 540))
                    d_156_v88_: _dafny.Array
                    nw19_ = _dafny.Array(_dafny.Map({}), 10)
                    d_156_v88_ = nw19_
                    index11_ = default__.safeIndex(640, (d_156_v88_).length(0))
                    (d_156_v88_)[index11_] = (d_60_v5_) | (default__.fm15((d_141_v74_).f9, globalState))
                    index12_ = default__.safeIndex(640, (d_156_v88_).length(0))
                    def iife16_():
                        coll12_ = _dafny.Map()
                        compr_12_: int
                        for compr_12_ in (d_60_v5_).keys.Elements:
                            d_157_v89_: int = compr_12_
                            if (d_157_v89_) in (d_60_v5_):
                                coll12_[(d_157_v89_) + (d_55_v1_)] = p0
                        return _dafny.Map(coll12_)
                    (d_156_v88_)[index12_] = (iife16_()
                    ) | (d_60_v5_)
                    pass
            pass

    @staticmethod
    def Main(noArgsParameter__):
        d_158_v0_: bool
        d_158_v0_ = True
        d_159_v1_: _dafny.Array
        nw20_ = _dafny.Array(_dafny.Map({}), 18)
        d_159_v1_ = nw20_
        d_160_v2_: _dafny.Seq
        d_160_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pa"))
        d_161_globalState_: GlobalState
        nw21_ = GlobalState()
        nw21_.ctor__(True, 162, True, False, (d_159_v1_ if d_158_v0_ else d_159_v1_), True, 253, d_160_v2_, False)
        d_161_globalState_ = nw21_
        d_162_v3_: _dafny.Array
        def lambda3_(d_163_v0_):
            def lambda4_(d_164_i1_):
                return (d_164_i1_) + ((0) - (len(_dafny.Map({d_163_v0_: 882}))))

            return lambda4_

        init2_ = lambda3_(d_158_v0_)
        nw22_ = _dafny.Array(None, 7)
        for i0_2_ in range(nw22_.length(0)):
            nw22_[i0_2_] = init2_(i0_2_)
        d_162_v3_ = nw22_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_162_v3_).length(0)):
            d_165_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_165_i0_)) and ((d_165_i0_) < ((d_162_v3_).length(0)))):
                (d_162_v3_)[(d_165_i0_)] = default__.safeDivisionInt(d_165_i0_, len((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_158_v0_, d_158_v0_]) for d_166_i2_ in range(default__.abs(105))]) if d_158_v0_ else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_158_v0_, d_158_v0_])), _dafny.MultiSet([d_158_v0_, d_158_v0_, d_158_v0_])]))))
        d_167_v4_: int
        d_167_v4_ = 490
        hi0_ = default__.fm0(d_161_globalState_)
        for d_168_i3_ in range(d_167_v4_, hi0_):
            default__.m0(d_158_v0_, d_161_globalState_)
            index13_ = default__.safeIndex(487, (d_162_v3_).length(0))
            (d_162_v3_)[index13_] = (0) - ((d_168_i3_) * (d_168_i3_))
            d_169_v5_: _dafny.Array
            def lambda5_(d_170_i4_):
                return False

            init3_ = lambda5_
            nw23_ = _dafny.Array(None, 9)
            for i0_3_ in range(nw23_.length(0)):
                nw23_[i0_3_] = init3_(i0_3_)
            d_169_v5_ = nw23_
            d_171_v6_: str
            d_171_v6_ = _dafny.CodePoint('g')
            d_172_v7_: _dafny.Map
            d_172_v7_ = _dafny.Map({d_169_v5_: d_171_v6_})
            index14_ = default__.safeIndex(487, (d_162_v3_).length(0))
            rhs21_ = len(d_172_v7_)
            rhs22_ = d_168_i3_
            lhs10_ = d_162_v3_
            lhs11_ = default__.safeIndex(487, (d_162_v3_).length(0))
            d_167_v4_ = rhs21_
            lhs10_[lhs11_] = rhs22_
            d_173_v8_: _dafny.Map
            d_173_v8_ = _dafny.Map({d_167_v4_: (d_162_v3_)[default__.safeIndex(487, (d_162_v3_).length(0))]})
            d_174_v9_: _dafny.Seq
            d_174_v9_ = _dafny.SeqWithoutIsStrInference([d_169_v5_, d_169_v5_])
            d_167_v4_ = default__.safeDivisionInt(len(d_173_v8_), len((d_174_v9_) + (d_174_v9_)))
            d_175_v10_: _dafny.MultiSet
            d_175_v10_ = _dafny.MultiSet([(d_162_v3_)[default__.safeIndex(487, (d_162_v3_).length(0))], (d_162_v3_)[default__.safeIndex(487, (d_162_v3_).length(0))]])
            d_158_v0_ = default__.fm1((d_175_v10_ if d_158_v0_ else d_175_v10_), (d_162_v3_)[default__.safeIndex(487, (d_162_v3_).length(0))], d_161_globalState_)
        d_176_i5_: int
        d_176_i5_ = 0
        with _dafny.label("2"):
            while d_158_v0_:
                with _dafny.c_label("2"):
                    if (d_176_i5_) >= (100):
                        raise _dafny.Break("2")
                    d_176_i5_ = (d_176_i5_) + (1)
                    d_177_v11_: str
                    d_177_v11_ = _dafny.CodePoint('j')
                    d_178_v12_: _dafny.Map
                    d_178_v12_ = _dafny.Map({_dafny.Map({d_177_v11_: d_167_v4_}): d_158_v0_})
                    d_179_v13_: _dafny.Seq
                    d_179_v13_ = _dafny.SeqWithoutIsStrInference([d_160_v2_, _dafny.SeqWithoutIsStrInference([d_177_v11_ for d_180_i6_ in range(default__.abs(655))]), d_160_v2_, d_160_v2_, d_160_v2_])
                    d_181_v15_: _dafny.Map
                    def iife17_():
                        coll13_ = _dafny.Set()
                        compr_13_: _dafny.Seq
                        for compr_13_ in (d_179_v13_).Elements:
                            d_182_v14_: _dafny.Seq = compr_13_
                            if (d_182_v14_) in (d_179_v13_):
                                coll13_ = coll13_.union(_dafny.Set([d_182_v14_]))
                        return _dafny.Set(coll13_)
                    d_181_v15_ = _dafny.Map({d_177_v11_: len(iife17_()
                    )})
                    d_178_v12_ = (d_178_v12_).set((d_181_v15_) | (d_181_v15_), d_158_v0_)
                    d_183_v16_: C4
                    nw24_ = C4()
                    nw24_.ctor__(d_158_v0_)
                    d_183_v16_ = nw24_
                    (d_161_globalState_).f2 = (((default__.fm28(d_167_v4_, not(d_158_v0_), d_161_globalState_)) + (d_160_v2_)).set(default__.safeIndex(d_167_v4_, len((default__.fm28(d_167_v4_, not(d_158_v0_), d_161_globalState_)) + (d_160_v2_))), d_177_v11_)) <= (_dafny.SeqWithoutIsStrInference([d_177_v11_]))
                    if not (d_158_v0_) or (d_158_v0_):
                        d_184_v17_: _dafny.Seq
                        d_184_v17_ = _dafny.SeqWithoutIsStrInference([d_158_v0_])
                        d_185_v18_: _dafny.MultiSet
                        d_185_v18_ = _dafny.MultiSet([(_dafny.MultiSet([d_167_v4_])).cardinality, len(d_184_v17_)])
                        d_186_v19_: _dafny.Set
                        d_186_v19_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_185_v18_).cardinality]), _dafny.SeqWithoutIsStrInference([d_167_v4_ for d_187_i7_ in range(default__.abs(-165))])})
                        d_188_v20_: C0
                        nw25_ = C0()
                        nw25_.ctor__(d_186_v19_, (0) - (d_167_v4_))
                        d_188_v20_ = nw25_
                        d_189_v21_: D1
                        d_189_v21_ = D1_DC2(d_188_v20_.f12, (d_188_v20_.f12) == (d_167_v4_), default__.safeModuloInt(d_167_v4_, d_188_v20_.f12), d_167_v4_)
                        d_189_v21_ = d_189_v21_
                        d_190_v22_: _dafny.Map
                        d_190_v22_ = _dafny.Map({d_167_v4_: d_162_v3_})
                        rhs23_ = (d_190_v22_) | ((d_190_v22_) | (d_190_v22_))
                        rhs24_ = d_159_v1_
                        d_190_v22_ = rhs23_
                        d_159_v1_ = rhs24_
                        d_191_v23_: C2
                        nw26_ = C2()
                        nw26_.ctor__(d_167_v4_, d_158_v0_)
                        d_191_v23_ = nw26_
                        d_191_v23_ = d_191_v23_
                        d_158_v0_ = (len(d_160_v2_)) >= (d_167_v4_)
                    elif True:
                        d_192_v24_: _dafny.Array
                        nw27_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                        d_192_v24_ = nw27_
                        index15_ = default__.safeIndex(589, (d_192_v24_).length(0))
                        (d_192_v24_)[index15_] = (d_177_v11_ if d_158_v0_ else _dafny.CodePoint('o'))
                        index16_ = default__.safeIndex(589, (d_192_v24_).length(0))
                        (d_192_v24_)[index16_] = d_177_v11_
                        rhs25_ = d_167_v4_
                        rhs26_ = d_160_v2_
                        d_167_v4_ = rhs25_
                        d_160_v2_ = rhs26_
                        d_193_v25_: _dafny.Seq
                        d_193_v25_ = _dafny.SeqWithoutIsStrInference([d_158_v0_])
                        d_194_v26_: C5
                        nw28_ = C5()
                        nw28_.ctor__(len(d_193_v25_), (d_158_v0_) or (not(d_158_v0_)))
                        d_194_v26_ = nw28_
                        d_195_v27_: D1
                        d_195_v27_ = D1_DC2((d_194_v26_).f10, True, d_167_v4_, (0) - ((d_194_v26_).f10))
                        d_196_v28_: _dafny.Seq
                        d_196_v28_ = _dafny.SeqWithoutIsStrInference([d_195_v27_])
                        d_197_v29_: T0
                        nw29_ = C1()
                        nw29_.ctor__(d_196_v28_)
                        d_197_v29_ = nw29_
                        d_198_v30_: _dafny.Map
                        d_198_v30_ = _dafny.Map({d_197_v29_: d_167_v4_})
                        d_198_v30_ = (d_198_v30_).set(d_197_v29_, 324)
                        d_199_v31_: _dafny.Seq
                        d_199_v31_ = _dafny.SeqWithoutIsStrInference([(d_194_v26_).f10])
                        d_200_v32_: _dafny.Set
                        d_200_v32_ = _dafny.Set({d_199_v31_})
                        d_201_v33_: _dafny.Map
                        d_201_v33_ = _dafny.Map({(d_194_v26_).f10: d_167_v4_})
                        d_202_v34_: _dafny.MultiSet
                        d_202_v34_ = _dafny.MultiSet([d_162_v3_, d_162_v3_])
                        d_203_v35_: C0
                        nw30_ = C0()
                        nw30_.ctor__(d_200_v32_, ((d_201_v33_)[901] if (901) in (d_201_v33_) else ((d_202_v34_)[d_162_v3_] if (d_162_v3_) in (d_202_v34_) else (0) - (d_167_v4_))))
                        d_203_v35_ = nw30_
                    pass
            pass
        hi1_ = d_167_v4_
        for d_204_i8_ in range(len(d_160_v2_), hi1_):
            d_205_v36_: str
            d_205_v36_ = _dafny.CodePoint('n')
            d_206_v37_: _dafny.MultiSet
            d_206_v37_ = _dafny.MultiSet([d_205_v36_, d_205_v36_, d_205_v36_])
            d_207_v38_: _dafny.Seq
            d_207_v38_ = _dafny.SeqWithoutIsStrInference([D1_DC2(d_204_i8_, d_158_v0_, d_167_v4_, (d_206_v37_).cardinality)])
            d_208_v39_: T0
            nw31_ = C1()
            nw31_.ctor__((d_207_v38_) + (d_207_v38_))
            d_208_v39_ = nw31_
            d_209_v41_: D17
            d_209_v41_ = D17_DC48(d_208_v39_)
            def iife18_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(-582, 489):
                    d_210_v40_: int = compr_14_
                    if ((-582) <= (d_210_v40_)) and ((d_210_v40_) < (489)):
                        coll14_[default__.safeModuloInt(d_210_v40_, d_167_v4_)] = d_160_v2_
                return _dafny.Map(coll14_)
            rhs27_ = (default__.fm46(d_158_v0_, d_158_v0_, not(not(d_158_v0_)), d_161_globalState_)) == (iife18_()
            )
            rhs28_ = d_167_v4_
            rhs29_ = (d_209_v41_).cf81
            lhs12_ = d_161_globalState_
            lhs12_.f2 = rhs27_
            d_167_v4_ = rhs28_
            d_208_v39_ = rhs29_
            default__.m0(d_158_v0_, d_161_globalState_)
            (d_161_globalState_).f2 = (d_204_i8_) != (d_167_v4_)
            d_211_v42_: _dafny.Seq
            d_211_v42_ = _dafny.SeqWithoutIsStrInference([d_158_v0_])
            rhs30_ = d_204_i8_
            rhs31_ = d_158_v0_
            rhs32_ = not (d_158_v0_) or (((d_211_v42_)[default__.safeIndex(d_167_v4_, len(d_211_v42_))] if d_158_v0_ else d_158_v0_))
            lhs13_ = d_161_globalState_
            d_167_v4_ = rhs30_
            lhs13_.f2 = rhs31_
            d_158_v0_ = rhs32_
        default__.m0(d_158_v0_, d_161_globalState_)
        d_212_i9_: int
        d_212_i9_ = 0
        with _dafny.label("3"):
            while d_158_v0_:
                with _dafny.c_label("3"):
                    if (d_212_i9_) >= (100):
                        raise _dafny.Break("3")
                    d_212_i9_ = (d_212_i9_) + (1)
                    d_213_v43_: C2
                    nw32_ = C2()
                    nw32_.ctor__(98, not(d_158_v0_))
                    d_213_v43_ = nw32_
                    d_214_v44_: _dafny.Seq
                    d_214_v44_ = _dafny.SeqWithoutIsStrInference([d_213_v43_, d_213_v43_, d_213_v43_, d_213_v43_, d_213_v43_])
                    d_214_v44_ = d_214_v44_
                    d_215_v45_: _dafny.MultiSet
                    d_215_v45_ = _dafny.MultiSet([d_167_v4_])
                    d_216_v47_: _dafny.Array
                    def lambda6_(d_217_v43_):
                        def lambda7_(d_218_i10_):
                            def iife19_():
                                coll15_ = _dafny.Set()
                                compr_15_: int
                                for compr_15_ in _dafny.IntegerRange(809, -805):
                                    d_219_v46_: int = compr_15_
                                    if ((809) <= (d_219_v46_)) and ((d_219_v46_) < (-805)):
                                        coll15_ = coll15_.union(_dafny.Set([(d_219_v46_) - ((d_217_v43_).f14)]))
                                return _dafny.Set(coll15_)
                            return iife19_()
                            

                        return lambda7_

                    init4_ = lambda6_(d_213_v43_)
                    nw33_ = _dafny.Array(None, 11)
                    for i0_4_ in range(nw33_.length(0)):
                        nw33_[i0_4_] = init4_(i0_4_)
                    d_216_v47_ = nw33_
                    d_220_v48_: C3
                    nw34_ = C3()
                    nw34_.ctor__(False, d_216_v47_, d_158_v0_)
                    d_220_v48_ = nw34_
                    d_221_v49_: _dafny.Map
                    d_221_v49_ = _dafny.Map({(d_213_v43_).f14: d_220_v48_})
                    d_222_v50_: _dafny.Map
                    d_222_v50_ = _dafny.Map({d_158_v0_: d_160_v2_})
                    d_223_v51_: _dafny.Seq
                    d_223_v51_ = _dafny.SeqWithoutIsStrInference([len(((d_222_v50_)[(d_220_v48_).f15] if ((d_220_v48_).f15) in (d_222_v50_) else d_160_v2_)), d_167_v4_])
                    d_224_v52_: _dafny.Seq
                    d_224_v52_ = _dafny.SeqWithoutIsStrInference([len(d_223_v51_)])
                    d_225_v53_: D9
                    d_225_v53_ = D9_DC24(default__.fm1(d_215_v45_, len(d_221_v49_), d_161_globalState_), len(d_223_v51_), d_160_v2_)
                    default__.m0((d_225_v53_).cf41, d_161_globalState_)
                    d_160_v2_ = d_160_v2_
                    d_226_v54_: C3
                    nw35_ = C3()
                    nw35_.ctor__((d_220_v48_).f15, d_220_v48_.f16, d_158_v0_)
                    d_226_v54_ = nw35_
                    pass
            pass
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_162_v3_).length(0)):
            d_227_i11_: int = guard_loop_1_
            if (True) and (((0) <= (d_227_i11_)) and ((d_227_i11_) < ((d_162_v3_).length(0)))):
                (d_162_v3_)[(d_227_i11_)] = default__.safeDivisionInt(d_227_i11_, d_167_v4_)
        index17_ = default__.safeIndex(547, (d_162_v3_).length(0))
        (d_162_v3_)[index17_] = 77
        d_228_v55_: str
        d_228_v55_ = _dafny.CodePoint('m')
        d_229_v56_: D1
        d_229_v56_ = D1_DC2(d_167_v4_, d_158_v0_, d_167_v4_, len(d_160_v2_))
        d_230_v57_: _dafny.Seq
        d_230_v57_ = _dafny.SeqWithoutIsStrInference([d_229_v56_])
        d_231_v58_: T0
        nw36_ = C1()
        nw36_.ctor__(d_230_v57_)
        d_231_v58_ = nw36_
        d_232_v59_: D17
        d_232_v59_ = D17_DC50(d_228_v55_, d_158_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awbtnvw")), d_167_v4_, d_231_v58_)
        index18_ = default__.safeIndex(547, (d_162_v3_).length(0))
        rhs33_ = d_167_v4_
        rhs34_ = d_232_v59_
        lhs14_ = d_162_v3_
        lhs15_ = default__.safeIndex(547, (d_162_v3_).length(0))
        lhs14_[lhs15_] = rhs33_
        d_232_v59_ = rhs34_
        d_233_v60_: C1
        nw37_ = C1()
        nw37_.ctor__(d_230_v57_)
        d_233_v60_ = nw37_
        d_234_v61_: D13
        d_234_v61_ = D13_DC35(d_158_v0_, d_228_v55_, d_233_v60_, d_158_v0_)
        d_235_v62_: _dafny.Array
        nw38_ = _dafny.Array(None, 25)
        nw38_[int(0)] = d_233_v60_
        nw38_[int(1)] = d_233_v60_
        nw38_[int(2)] = d_233_v60_
        nw38_[int(3)] = d_233_v60_
        nw38_[int(4)] = d_233_v60_
        nw38_[int(5)] = d_233_v60_
        nw38_[int(6)] = d_233_v60_
        nw38_[int(7)] = d_233_v60_
        nw38_[int(8)] = d_233_v60_
        nw38_[int(9)] = d_233_v60_
        nw38_[int(10)] = d_233_v60_
        nw38_[int(11)] = d_233_v60_
        nw38_[int(12)] = d_233_v60_
        nw38_[int(13)] = d_233_v60_
        nw38_[int(14)] = d_233_v60_
        nw38_[int(15)] = (d_233_v60_ if False else d_233_v60_)
        nw38_[int(16)] = d_233_v60_
        nw38_[int(17)] = d_233_v60_
        nw38_[int(18)] = (d_234_v61_).cf60
        nw38_[int(19)] = d_233_v60_
        nw38_[int(20)] = d_233_v60_
        nw38_[int(21)] = d_233_v60_
        nw38_[int(22)] = d_233_v60_
        nw38_[int(23)] = d_233_v60_
        nw38_[int(24)] = d_233_v60_
        d_235_v62_ = nw38_
        d_235_v62_ = d_235_v62_
        d_236_v63_: _dafny.Seq
        d_236_v63_ = _dafny.SeqWithoutIsStrInference([d_162_v3_])
        d_237_i12_: int
        d_237_i12_ = 0
        with _dafny.label("4"):
            while (False if d_158_v0_ else (len(d_236_v63_)) <= (default__.fm0(d_161_globalState_))):
                with _dafny.c_label("4"):
                    if (d_237_i12_) >= (100):
                        raise _dafny.Break("4")
                    d_237_i12_ = (d_237_i12_) + (1)
                    d_238_v64_: _dafny.Array
                    nw39_ = _dafny.Array(_dafny.Seq({}), 22)
                    d_238_v64_ = nw39_
                    d_238_v64_ = d_238_v64_
                    d_239_v65_: D16
                    d_239_v65_ = D16_DC46(d_167_v4_)
                    d_240_v66_: D16
                    d_240_v66_ = D16_DC47(d_239_v65_)
                    d_241_v67_: D16
                    d_241_v67_ = D16_DC47(d_240_v66_)
                    d_242_v68_: _dafny.Map
                    d_242_v68_ = _dafny.Map({d_167_v4_: d_241_v67_})
                    d_242_v68_ = (d_242_v68_).set(d_167_v4_, d_241_v67_)
                    if not(d_158_v0_):
                        d_243_v69_: _dafny.Map
                        d_243_v69_ = _dafny.Map({d_158_v0_: d_158_v0_})
                        d_243_v69_ = (d_243_v69_).set(d_158_v0_, d_158_v0_)
                        d_244_v70_: _dafny.Seq
                        d_244_v70_ = _dafny.SeqWithoutIsStrInference([(d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]])
                        d_245_v71_: _dafny.Set
                        d_245_v71_ = _dafny.Set({d_244_v70_, _dafny.SeqWithoutIsStrInference([948, (d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]]), d_244_v70_})
                        d_246_v72_: C0
                        nw40_ = C0()
                        nw40_.ctor__(d_245_v71_, ((d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]) + (len(d_243_v69_)))
                        d_246_v72_ = nw40_
                        d_247_v73_: _dafny.Map
                        d_247_v73_ = _dafny.Map({_dafny.MultiSet([default__.fm0(d_161_globalState_)]): d_158_v0_})
                        d_248_v74_: D7
                        d_248_v74_ = D7_DC18(d_247_v73_)
                        d_249_v75_: _dafny.Map
                        d_249_v75_ = _dafny.Map({d_167_v4_: d_248_v74_})
                        d_249_v75_ = (d_249_v75_).set((d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))], d_248_v74_)
                        pat_let_tv2_ = d_161_globalState_
                        def iife20_(_pat_let2_0):
                            def iife21_(d_250_dt__update__tmp_h0_):
                                def iife22_(_pat_let3_0):
                                    def iife23_(d_251_dt__update_hcf32_h0_):
                                        return D7_DC18(d_251_dt__update_hcf32_h0_)
                                    return iife23_(_pat_let3_0)
                                return iife22_(default__.fm22(pat_let_tv2_))
                            return iife21_(_pat_let2_0)
                        d_248_v74_ = iife20_(d_248_v74_)
                        (d_246_v72_).f12 = d_167_v4_
                    elif True:
                        d_158_v0_ = d_158_v0_
                        d_252_v76_: _dafny.Array
                        def lambda8_(d_253_v3_, d_254_v4_):
                            def lambda9_(d_255_i13_):
                                return (_dafny.Map({_dafny.SeqWithoutIsStrInference([(d_253_v3_)[default__.safeIndex(547, (d_253_v3_).length(0))]]): (d_253_v3_)[default__.safeIndex(547, (d_253_v3_).length(0))]})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_254_v4_ for d_256_i14_ in range(default__.abs(694))]): len(_dafny.Set({(d_253_v3_)[default__.safeIndex(547, (d_253_v3_).length(0))]}))}))

                            return lambda9_

                        init5_ = lambda8_(d_162_v3_, d_167_v4_)
                        nw41_ = _dafny.Array(None, 28)
                        for i0_5_ in range(nw41_.length(0)):
                            nw41_[i0_5_] = init5_(i0_5_)
                        d_252_v76_ = nw41_
                        index19_ = default__.safeIndex(880, (d_252_v76_).length(0))
                        (d_252_v76_)[index19_] = default__.fm47(d_161_globalState_)
                        d_257_v77_: _dafny.Seq
                        d_257_v77_ = _dafny.SeqWithoutIsStrInference([d_167_v4_, (d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]])
                        d_258_v78_: _dafny.Seq
                        d_258_v78_ = d_257_v77_
                        d_259_v79_: _dafny.Map
                        d_259_v79_ = _dafny.Map({d_258_v78_: d_167_v4_})
                        index20_ = default__.safeIndex(880, (d_252_v76_).length(0))
                        (d_252_v76_)[index20_] = ((_dafny.Map({d_258_v78_: d_167_v4_})) | ((d_259_v79_).set(d_257_v77_, len(d_257_v77_)))).set(d_258_v78_, d_167_v4_)
                        d_260_v80_: _dafny.MultiSet
                        d_260_v80_ = _dafny.MultiSet([(d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]])
                        d_261_v81_: _dafny.Array
                        d_261_v81_ = d_162_v3_
                        rhs35_ = d_158_v0_
                        rhs36_ = d_158_v0_
                        rhs37_ = ((d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]) <= ((d_260_v80_).cardinality)
                        rhs38_ = d_158_v0_
                        rhs39_ = (d_261_v81_)
                        lhs16_ = d_161_globalState_
                        lhs17_ = d_161_globalState_
                        lhs18_ = d_161_globalState_
                        lhs16_.f2 = rhs35_
                        lhs17_.f2 = rhs36_
                        d_158_v0_ = rhs37_
                        lhs18_.f2 = rhs38_
                        d_162_v3_ = rhs39_
                        d_262_v82_: D15
                        d_262_v82_ = D15_DC43()
                        d_263_v83_: D15
                        d_263_v83_ = D15_DC44(d_262_v82_)
                        d_264_v84_: T1
                        nw42_ = C2()
                        nw42_.ctor__(len(d_160_v2_), not(True))
                        d_264_v84_ = nw42_
                        d_265_v85_: _dafny.Seq
                        d_265_v85_ = _dafny.SeqWithoutIsStrInference([d_264_v84_, d_264_v84_])
                        d_266_v86_: C2
                        nw43_ = C2()
                        nw43_.ctor__(default__.fm0(d_161_globalState_), True)
                        d_266_v86_ = nw43_
                        d_267_v87_: _dafny.Map
                        d_267_v87_ = _dafny.Map({len((d_265_v85_) + (d_265_v85_)): d_266_v86_})
                        d_268_v88_: _dafny.Array
                        def lambda10_(d_269_v2_):
                            def lambda11_(d_270_i15_):
                                return _dafny.Set({(0) - (len(d_269_v2_))})

                            return lambda11_

                        init6_ = lambda10_(d_160_v2_)
                        nw44_ = _dafny.Array(None, 3)
                        for i0_6_ in range(nw44_.length(0)):
                            nw44_[i0_6_] = init6_(i0_6_)
                        d_268_v88_ = nw44_
                        d_271_v89_: T2
                        nw45_ = C3()
                        nw45_.ctor__(False, d_268_v88_, (d_264_v84_).f9)
                        d_271_v89_ = nw45_
                        d_272_v90_: D15
                        d_272_v90_ = D15_DC42(d_158_v0_, (d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))], d_271_v89_, d_167_v4_)
                        d_273_v91_: _dafny.Map
                        d_273_v91_ = _dafny.Map({d_272_v90_: ((d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]) < ((0) - (d_167_v4_))})
                        pat_let_tv3_ = d_262_v82_
                        def iife24_(_pat_let4_0):
                            def iife25_(d_274_dt__update__tmp_h2_):
                                def iife26_(_pat_let5_0):
                                    def iife27_(d_275_dt__update_hcf77_h0_):
                                        return D15_DC44(d_275_dt__update_hcf77_h0_)
                                    return iife27_(_pat_let5_0)
                                return iife26_(pat_let_tv3_)
                            return iife25_(_pat_let4_0)
                        rhs40_ = (0) - ((0) - (default__.safeModuloInt(d_167_v4_, ((d_266_v86_).f14) * ((d_266_v86_).f14))))
                        rhs41_ = iife24_(d_263_v83_)
                        rhs42_ = ((d_273_v91_)[d_272_v90_] if (d_272_v90_) in (d_273_v91_) else (d_271_v89_).f9)
                        rhs43_ = (d_267_v87_) | (d_267_v87_)
                        lhs19_ = d_161_globalState_
                        d_167_v4_ = rhs40_
                        d_263_v83_ = rhs41_
                        lhs19_.f2 = rhs42_
                        d_267_v87_ = rhs43_
                        d_276_v92_: _dafny.Array
                        nw46_ = _dafny.Array(None, 2)
                        nw46_[int(0)] = (d_271_v89_).f9
                        nw46_[int(1)] = (d_271_v89_).f9
                        d_276_v92_ = nw46_
                        d_277_v93_: bool
                        d_278_v94_: bool
                        out8_: bool
                        out9_: bool
                        out8_, out9_ = (d_233_v60_).m6((d_231_v58_ if d_158_v0_ else d_231_v58_), d_276_v92_, (d_264_v84_).f9, d_161_globalState_)
                        d_277_v93_ = out8_
                        d_278_v94_ = out9_
                    d_167_v4_ = (d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]
                    pass
            pass
        d_279_v95_: _dafny.Map
        d_279_v95_ = _dafny.Map({d_167_v4_: (len(default__.fm23(d_158_v0_, d_158_v0_, d_167_v4_, d_161_globalState_))) + (d_167_v4_)})
        d_280_v96_: _dafny.Map
        d_280_v96_ = _dafny.Map({d_167_v4_: d_158_v0_})
        d_281_v97_: D4
        d_281_v97_ = D4_DC13(d_167_v4_)
        d_282_v98_: C2
        nw47_ = C2()
        nw47_.ctor__((d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))], d_158_v0_)
        d_282_v98_ = nw47_
        d_283_v99_: _dafny.MultiSet
        d_283_v99_ = _dafny.MultiSet([d_282_v98_])
        d_284_v100_: _dafny.Seq
        d_284_v100_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_158_v0_: (d_282_v98_).f14})])
        d_285_v101_: _dafny.Seq
        d_285_v101_ = _dafny.SeqWithoutIsStrInference([(d_284_v100_)[default__.safeIndex((d_282_v98_).f14, len(d_284_v100_))]])
        d_286_v102_: _dafny.Map
        d_286_v102_ = _dafny.Map({d_158_v0_: len(_dafny.Set({d_158_v0_, d_158_v0_}))})
        d_287_v103_: _dafny.Map
        d_287_v103_ = _dafny.Map({d_158_v0_: d_158_v0_})
        d_288_v104_: D9
        d_288_v104_ = D9_DC24(d_158_v0_, (d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqe")))
        pat_let_tv4_ = d_162_v3_
        pat_let_tv5_ = d_162_v3_
        index21_ = default__.safeIndex(547, (d_162_v3_).length(0))
        def iife28_(_pat_let6_0):
            def iife29_(d_289_dt__update__tmp_h3_):
                def iife30_(_pat_let7_0):
                    def iife31_(d_290_dt__update_hcf25_h0_):
                        return D4_DC13(d_290_dt__update_hcf25_h0_)
                    return iife31_(_pat_let7_0)
                return iife30_((pat_let_tv5_)[default__.safeIndex(547, (pat_let_tv4_).length(0))])
            return iife29_(_pat_let6_0)
        rhs44_ = (d_279_v95_) | ((default__.fm48(d_158_v0_, default__.fm11(False, d_161_globalState_), d_158_v0_, ((d_279_v95_)[(0) - ((d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))])] if ((0) - ((d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))])) in (d_279_v95_) else len(d_280_v96_)), d_161_globalState_)).set(-614, len(_dafny.Map({iife28_(d_281_v97_): d_231_v58_}))))
        rhs45_ = (d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]
        rhs46_ = d_228_v55_
        rhs47_ = ((d_283_v99_)[d_282_v98_] if (d_282_v98_) in (d_283_v99_) else len(((d_284_v100_)[default__.safeIndex(d_167_v4_, len(d_284_v100_))] if False else d_286_v102_)))
        rhs48_ = not ((False if d_158_v0_ else not((d_288_v104_).cf41))) or ((len(d_287_v103_)) > ((d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]))
        lhs20_ = d_162_v3_
        lhs21_ = default__.safeIndex(547, (d_162_v3_).length(0))
        lhs22_ = d_161_globalState_
        d_279_v95_ = rhs44_
        d_167_v4_ = rhs45_
        d_228_v55_ = rhs46_
        lhs20_[lhs21_] = rhs47_
        lhs22_.f2 = rhs48_
        d_291_v105_: _dafny.Array
        nw48_ = _dafny.Array(False, 7)
        d_291_v105_ = nw48_
        index22_ = default__.safeIndex(547, (d_162_v3_).length(0))
        rhs49_ = d_291_v105_
        rhs50_ = (d_162_v3_)[default__.safeIndex(547, (d_162_v3_).length(0))]
        rhs51_ = D4_DC13(default__.safeModuloInt(len(d_160_v2_), (d_282_v98_).f14))
        rhs52_ = not(d_158_v0_)
        lhs23_ = d_162_v3_
        lhs24_ = default__.safeIndex(547, (d_162_v3_).length(0))
        d_291_v105_ = rhs49_
        lhs23_[lhs24_] = rhs50_
        d_281_v97_ = rhs51_
        d_158_v0_ = rhs52_
        d_292_v106_: _dafny.Array
        nw49_ = _dafny.Array(_dafny.Map({}), 19)
        d_292_v106_ = nw49_
        d_292_v106_ = d_292_v106_
        d_293_v107_: D15
        d_293_v107_ = D15_DC41()
        d_294_v108_: D15
        d_294_v108_ = D15_DC44(d_293_v107_)
        d_295_v109_: _dafny.Set
        d_295_v109_ = _dafny.Set({d_158_v0_, d_158_v0_, d_158_v0_, d_158_v0_})
        d_296_v110_: _dafny.Map
        d_296_v110_ = _dafny.Map({d_294_v108_: (d_295_v109_) - (d_295_v109_)})
        d_296_v110_ = (d_296_v110_) | (d_296_v110_)
        d_297_v111_: D17
        d_297_v111_ = D17_DC48((d_231_v58_ if d_158_v0_ else d_231_v58_))
        d_297_v111_ = d_297_v111_
        index23_ = default__.safeIndex(547, (d_162_v3_).length(0))
        (d_162_v3_)[index23_] = d_167_v4_
        _dafny.print(_dafny.string_of(d_158_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_160_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_161_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_161_globalState_).f7).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_167_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_176_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_212_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_228_v55_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v56_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v56_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v56_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v56_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_230_v57_) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v59_).cf82))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v59_).cf83))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_232_v59_).cf84).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v59_).cf85))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v60_.f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v61_).cf58))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v61_).cf59))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v61_).cf60.f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v61_).cf61))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[0].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[1].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[2].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[3].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[4].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[5].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[6].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[7].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[8].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[9].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[10].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[11].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[12].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[13].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[14].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[15].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[16].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[17].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[18].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[19].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[20].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[21].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[22].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[23].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v62_)[24].f13) == (_dafny.SeqWithoutIsStrInference([D1_DC2(489, True, 489, 2)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_236_v63_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_237_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_279_v95_) == (_dafny.Map({489: 491, 2: 849, -614: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v96_) == (_dafny.Map({489: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_281_v97_).cf25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_v98_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_283_v99_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v100_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 489})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v101_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 489})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_286_v102_) == (_dafny.Map({True: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_287_v103_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_288_v104_).cf41))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_288_v104_).cf42))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_288_v104_).cf43).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v109_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_296_v110_)))
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
        return lambda: D1_DC2(int(0), False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)}, {self.cf7.VerbatimString(True)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC8(D3, NamedTuple('DC8', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(False, False, int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC12(D4, NamedTuple('DC12', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC16(D5, NamedTuple('DC16', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC17(D6, NamedTuple('DC17', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(_dafny.Array(None, 0), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
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
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC19(D7, NamedTuple('DC19', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {self.cf35.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC22(D8, NamedTuple('DC22', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24(False, int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC24(D9, NamedTuple('DC24', [('cf41', Any), ('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)}, {self.cf43.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf44', Any), ('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC28()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)

class D10_DC28(D10, NamedTuple('DC28', [])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28)
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC27(D10, NamedTuple('DC27', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC30()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)

class D11_DC30(D11, NamedTuple('DC30', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC32(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)

class D12_DC32(D12, NamedTuple('DC32', [('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC35(False, _dafny.CodePoint('D'), None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC35(D13, NamedTuple('DC35', [('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC36(D13, NamedTuple('DC36', [('cf62', Any), ('cf63', Any), ('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({self.cf62.VerbatimString(True)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC38(_dafny.CodePoint('D'), False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)

class D14_DC38(D14, NamedTuple('DC38', [('cf67', Any), ('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC39(D14, NamedTuple('DC39', [('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC37(D14, NamedTuple('DC37', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC41()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D15_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)

class D15_DC41(D15, NamedTuple('DC41', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC41'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC41)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf73', Any), ('cf74', Any), ('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC43(D15, NamedTuple('DC43', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC40(D15, NamedTuple('DC40', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC44(D15, NamedTuple('DC44', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC46(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D16_DC47)

class D16_DC46(D16, NamedTuple('DC46', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC45(D16, NamedTuple('DC45', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC47(D16, NamedTuple('DC47', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC47({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC47) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC49()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D17_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D17_DC50)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D17_DC51)

class D17_DC49(D17, NamedTuple('DC49', [])):
    def __dafnystr__(self) -> str:
        return f'D17.DC49'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC49)
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC50(D17, NamedTuple('DC50', [('cf82', Any), ('cf83', Any), ('cf84', Any), ('cf85', Any), ('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC50({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {self.cf84.VerbatimString(True)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC50) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC48(D17, NamedTuple('DC48', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC51(D17, NamedTuple('DC51', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC51({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC51) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC53(int(0), None, _dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D18_DC53)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D18_DC52)

class D18_DC53(D18, NamedTuple('DC53', [('cf89', Any), ('cf90', Any), ('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC53({_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC53) and self.cf89 == __o.cf89 and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC52(D18, NamedTuple('DC52', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC52({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC52) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D19_DC54)

class D19_DC54(D19, NamedTuple('DC54', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC54({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC54) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D20_DC55)

class D20_DC55(D20, NamedTuple('DC55', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC55({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC55) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m1(self, p0, p1, p2, p3, globalState):
        pass


class T2:
    pass
    def m2(self, p0, p1, p2, p3, globalState):
        pass

    def m3(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: bool = False
        self._f0: bool = False
        self._f1: int = int(0)
        self._f3: bool = False
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: bool = False
        self._f6: int = int(0)
        self._f7: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f8: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
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

class C0:
    def  __init__(self):
        self.f11: _dafny.Set = _dafny.Set({})
        self.f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f11, f12):
        (self).f11 = f11
        (self).f12 = f12

    def fm8(self, p0, p1, p2, globalState):
        source4_ = D1_DC2(self.f12, not(not(True)), self.f12, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_298_i0_ in range(default__.abs(67))])))
        if source4_.is_DC2:
            d_299___mcc_h0_ = source4_.cf2
            d_300___mcc_h1_ = source4_.cf3
            d_301___mcc_h2_ = source4_.cf4
            d_302___mcc_h3_ = source4_.cf5
            d_303_cf5_ = d_302___mcc_h3_
            d_304_cf4_ = d_301___mcc_h2_
            d_305_cf3_ = d_300___mcc_h1_
            d_306_cf2_ = d_299___mcc_h0_
            return d_305_cf3_
        elif source4_.is_DC3:
            d_307___mcc_h4_ = source4_.cf6
            d_308___mcc_h5_ = source4_.cf7
            d_309___mcc_h6_ = source4_.cf8
            d_310_cf8_ = d_309___mcc_h6_
            d_311_cf7_ = d_308___mcc_h5_
            d_312_cf6_ = d_307___mcc_h4_
            if d_312_cf6_:
                return d_312_cf6_
            elif True:
                return d_312_cf6_
        elif source4_.is_DC4:
            d_313___mcc_h7_ = source4_.cf9
            d_314___mcc_h8_ = source4_.cf10
            d_315___mcc_h9_ = source4_.cf11
            d_316___mcc_h10_ = source4_.cf12
            d_317___mcc_h11_ = source4_.cf13
            d_318_cf13_ = d_317___mcc_h11_
            d_319_cf12_ = d_316___mcc_h10_
            d_320_cf11_ = d_315___mcc_h9_
            d_321_cf10_ = d_314___mcc_h8_
            d_322_cf9_ = d_313___mcc_h7_
            return True
        elif source4_.is_DC1:
            d_323___mcc_h12_ = source4_.cf1
            d_324_cf1_ = d_323___mcc_h12_
            return (self.f12) >= (d_324_cf1_)
        elif True:
            d_325___mcc_h13_ = source4_.cf14
            d_326_cf14_ = d_325___mcc_h13_
            return (D1_DC3((D1_DC4(self.f12, False, True, False, False)).cf11, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfyipkfqa")), self.f12)).cf6


class C1(T0):
    def  __init__(self):
        self.f13: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f13):
        (self).f13 = f13

    def fm2(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eiomh")))]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([628]))])).cardinality: False}))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_327_i0_ in range(default__.abs(649))]))]))) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, not(not(not(True))), not(False), not(not((D1_DC4(len(_dafny.SeqWithoutIsStrInference([976])), False, False, False, False)).cf12))})) for d_328_i1_ in range(default__.abs(56))]))

    def m6(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        d_329_v0_: int
        d_329_v0_ = 226
        d_330_v1_: _dafny.Seq
        d_330_v1_ = _dafny.SeqWithoutIsStrInference([d_329_v0_])
        d_331_v2_: _dafny.Seq
        d_331_v2_ = _dafny.SeqWithoutIsStrInference([d_330_v1_])
        d_332_v3_: _dafny.MultiSet
        d_332_v3_ = _dafny.MultiSet([d_329_v0_])
        d_333_v4_: _dafny.Seq
        d_333_v4_ = _dafny.SeqWithoutIsStrInference([not(p2), p2, p2])
        d_334_v5_: _dafny.Array
        nw50_ = _dafny.Array(None, 23)
        nw50_[int(0)] = d_329_v0_
        nw50_[int(1)] = len(d_331_v2_)
        nw50_[int(2)] = d_329_v0_
        nw50_[int(3)] = d_329_v0_
        nw50_[int(4)] = d_329_v0_
        nw50_[int(5)] = d_329_v0_
        nw50_[int(6)] = d_329_v0_
        nw50_[int(7)] = d_329_v0_
        nw50_[int(8)] = d_329_v0_
        nw50_[int(9)] = (0) - ((d_332_v3_).cardinality)
        nw50_[int(10)] = d_329_v0_
        nw50_[int(11)] = d_329_v0_
        nw50_[int(12)] = 893
        nw50_[int(13)] = d_329_v0_
        nw50_[int(14)] = d_329_v0_
        nw50_[int(15)] = (0) - (d_329_v0_)
        nw50_[int(16)] = len(d_333_v4_)
        nw50_[int(17)] = d_329_v0_
        nw50_[int(18)] = d_329_v0_
        nw50_[int(19)] = d_329_v0_
        nw50_[int(20)] = d_329_v0_
        nw50_[int(21)] = d_329_v0_
        nw50_[int(22)] = (0) - (default__.fm0(globalState))
        d_334_v5_ = nw50_
        d_335_v6_: _dafny.Map
        d_335_v6_ = _dafny.Map({d_334_v5_: d_334_v5_})
        d_335_v6_ = (d_335_v6_).set(d_334_v5_, d_334_v5_)
        d_336_v7_: _dafny.Array
        nw51_ = _dafny.Array(_dafny.Set({}), 11)
        d_336_v7_ = nw51_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_336_v7_).length(0)):
            d_337_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_337_i0_)) and ((d_337_i0_) < ((d_336_v7_).length(0)))):
                (d_336_v7_)[(d_337_i0_)] = ((_dafny.Set({p2})) | (_dafny.Set({p2, p2, p2, p2, p2}))) - (_dafny.Set({p2, p2}))
        d_329_v0_ = (default__.fm0(globalState) if p2 else default__.fm0(globalState))
        if ((0) - (d_329_v0_)) >= ((d_329_v0_) * (d_329_v0_)):
            d_338_v8_: _dafny.MultiSet
            d_338_v8_ = _dafny.MultiSet([p1, p1, p1])
            d_338_v8_ = ((D3_DC7(_dafny.MultiSet([p1]))).cf16) - (d_338_v8_)
            index24_ = default__.safeIndex(713, (d_334_v5_).length(0))
            (d_334_v5_)[index24_] = ((0) - (d_329_v0_)) + (d_329_v0_)
            d_339_v9_: _dafny.Map
            d_339_v9_ = _dafny.Map({d_329_v0_: p2})
            d_340_v10_: _dafny.Map
            d_340_v10_ = _dafny.Map({d_329_v0_: len(d_339_v9_)})
            index25_ = default__.safeIndex(713, (d_334_v5_).length(0))
            (d_334_v5_)[index25_] = (d_329_v0_) - (((d_340_v10_)[180] if (180) in (d_340_v10_) else (d_330_v1_)[default__.safeIndex(d_329_v0_, len(d_330_v1_))]))
            d_341_v11_: _dafny.Set
            d_341_v11_ = _dafny.Set({d_330_v1_})
            d_342_v12_: C0
            nw52_ = C0()
            nw52_.ctor__((d_341_v11_) | (_dafny.Set({d_330_v1_, d_330_v1_})), default__.fm0(globalState))
            d_342_v12_ = nw52_
            d_343_v13_: str
            d_343_v13_ = _dafny.CodePoint('j')
            d_344_v14_: D1
            d_344_v14_ = D1_DC3(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgnqr")), (d_334_v5_)[default__.safeIndex(713, (d_334_v5_).length(0))])
            d_345_v15_: _dafny.Map
            d_345_v15_ = _dafny.Map({d_343_v13_: ((0) - (d_342_v12_.f12)) != ((D1_DC4(d_329_v0_, (d_344_v14_).cf6, p2, p2, p2)).cf9)})
            d_346_v16_: _dafny.Set
            d_346_v16_ = _dafny.Set({d_342_v12_, d_342_v12_, d_342_v12_})
            d_347_v17_: _dafny.Map
            d_347_v17_ = _dafny.Map({(_dafny.Set({d_342_v12_})).isdisjoint(d_346_v16_): (d_342_v12_ if p2 else d_342_v12_)})
            index26_ = default__.safeIndex(713, (d_334_v5_).length(0))
            rhs53_ = ((d_347_v17_)[p2] if (p2) in (d_347_v17_) else d_342_v12_)
            rhs54_ = d_342_v12_.f12
            rhs55_ = p2
            rhs56_ = _dafny.Map({d_343_v13_: not(p2)})
            lhs25_ = d_334_v5_
            lhs26_ = default__.safeIndex(713, (d_334_v5_).length(0))
            lhs27_ = globalState
            d_342_v12_ = rhs53_
            lhs25_[lhs26_] = rhs54_
            lhs27_.f2 = rhs55_
            d_345_v15_ = rhs56_
            r1 = p2
        elif True:
            index27_ = default__.safeIndex(602, (p1).length(0))
            (p1)[index27_] = p2
            d_348_v18_: _dafny.MultiSet
            d_348_v18_ = _dafny.MultiSet([p2, not(p2)])
            index28_ = default__.safeIndex(602, (p1).length(0))
            rhs57_ = p2
            rhs58_ = p2
            rhs59_ = (((d_348_v18_) | (_dafny.MultiSet(d_333_v4_))).intersection(d_348_v18_)).cardinality
            rhs60_ = ((d_332_v3_).cardinality) * ((d_329_v0_) * (d_329_v0_))
            rhs61_ = d_329_v0_
            lhs28_ = globalState
            lhs29_ = p1
            lhs30_ = default__.safeIndex(602, (p1).length(0))
            lhs28_.f2 = rhs57_
            lhs29_[lhs30_] = rhs58_
            d_329_v0_ = rhs59_
            d_329_v0_ = rhs60_
            d_329_v0_ = rhs61_
            rhs62_ = (p1)[default__.safeIndex(602, (p1).length(0))]
            rhs63_ = d_329_v0_
            rhs64_ = d_332_v3_
            lhs31_ = globalState
            lhs31_.f2 = rhs62_
            d_329_v0_ = rhs63_
            d_332_v3_ = rhs64_
            d_349_v20_: _dafny.Set
            d_349_v20_ = _dafny.Set({p2})
            d_350_v21_: _dafny.MultiSet
            d_350_v21_ = _dafny.MultiSet([d_349_v20_, d_349_v20_, d_349_v20_, d_349_v20_])
            def iife32_():
                coll16_ = _dafny.Map()
                compr_16_: _dafny.Set
                for compr_16_ in (d_350_v21_).Elements:
                    d_351_v19_: _dafny.Set = compr_16_
                    if (d_351_v19_) in (d_350_v21_):
                        coll16_[d_351_v19_] = (p1)[default__.safeIndex(602, (p1).length(0))]
                return _dafny.Map(coll16_)
            d_329_v0_ = (len(iife32_()
            )) * (default__.safeDivisionInt(47, d_329_v0_))
            d_352_v22_: _dafny.Set
            d_352_v22_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_329_v0_])})
            d_353_v23_: C0
            nw53_ = C0()
            nw53_.ctor__(d_352_v22_, d_329_v0_)
            d_353_v23_ = nw53_
            d_354_v24_: _dafny.Map
            d_354_v24_ = _dafny.Map({d_353_v23_: (p1)[default__.safeIndex(602, (p1).length(0))]})
            d_354_v24_ = (d_354_v24_).set(d_353_v23_, not((d_353_v23_.f12) >= ((0) - (d_329_v0_))))
            (d_353_v23_).f12 = d_353_v23_.f12
        d_355_v25_: _dafny.Seq
        d_355_v25_ = d_330_v1_
        source5_ = d_355_v25_
        d_356___mcc_h0_ = source5_
        d_357_cf0_ = d_356___mcc_h0_
        d_358_v26_: str
        d_358_v26_ = _dafny.CodePoint('l')
        d_359_v27_: _dafny.Map
        d_359_v27_ = _dafny.Map({d_329_v0_: d_358_v26_})
        d_329_v0_ = (0) - (default__.safeDivisionInt((len(d_359_v27_)) * ((0) - (d_329_v0_)), d_329_v0_))
        rhs65_ = (False) or (True)
        rhs66_ = (d_329_v0_ if not(p2) else (0) - (len(((self).fm2(d_329_v0_, p2, p2, len(d_333_v4_), globalState)) + (d_357_cf0_))))
        rhs67_ = default__.safeModuloInt(default__.fm0(globalState), (0) - (d_329_v0_))
        rhs68_ = (not(p2)) and (p2)
        lhs32_ = globalState
        lhs33_ = globalState
        lhs32_.f2 = rhs65_
        d_329_v0_ = rhs66_
        d_329_v0_ = rhs67_
        lhs33_.f2 = rhs68_
        r0 = p2
        d_360_v28_: D3
        d_360_v28_ = D3_DC10()
        d_361_v29_: _dafny.Map
        d_361_v29_ = _dafny.Map({p2: d_360_v28_})
        d_362_v30_: _dafny.Seq
        d_362_v30_ = _dafny.SeqWithoutIsStrInference([d_361_v29_])
        d_363_v31_: _dafny.Map
        d_363_v31_ = _dafny.Map({d_329_v0_: d_329_v0_})
        r1 = (p2) not in ((d_362_v30_)[default__.safeIndex(len(d_363_v31_), len(d_362_v30_))])
        if True:
            d_364_v32_: str
            d_364_v32_ = _dafny.CodePoint('t')
            d_365_v33_: _dafny.Map
            d_365_v33_ = _dafny.Map({d_364_v32_: p2})
            index29_ = default__.safeIndex(7, (p1).length(0))
            (p1)[index29_] = ((d_365_v33_)[d_364_v32_] if (d_364_v32_) in (d_365_v33_) else p2)
            d_366_v34_: _dafny.Set
            d_366_v34_ = _dafny.Set({len((p0).fm2((0) - (d_329_v0_), True, p2, d_329_v0_, globalState))})
            index30_ = default__.safeIndex(7, (p1).length(0))
            (p1)[index30_] = ((_dafny.Set({d_329_v0_, (0) - (d_329_v0_)})) - (d_366_v34_)).issubset(d_366_v34_)
            d_367_v35_: _dafny.Seq
            d_367_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcyxdyb"))
            d_368_v36_: _dafny.Map
            d_368_v36_ = _dafny.Map({d_329_v0_: True})
            d_369_v37_: _dafny.Array
            nw54_ = _dafny.Array(None, 2)
            nw54_[int(0)] = p2
            nw54_[int(1)] = (_dafny.Set({d_329_v0_, len(d_330_v1_)})).isdisjoint(default__.fm14(d_329_v0_, d_329_v0_, len(d_368_v36_), globalState))
            d_369_v37_ = nw54_
            rhs69_ = ((d_367_v35_) + (_dafny.SeqWithoutIsStrInference([d_364_v32_ for d_370_i1_ in range(default__.abs(404))]))) + (d_367_v35_)
            rhs70_ = d_369_v37_
            d_367_v35_ = rhs69_
            d_369_v37_ = rhs70_
            if (p1)[default__.safeIndex(7, (p1).length(0))]:
                d_371_v38_: _dafny.Set
                d_371_v38_ = _dafny.Set({d_330_v1_, _dafny.SeqWithoutIsStrInference([d_329_v0_]), d_330_v1_, d_330_v1_})
                d_372_v39_: C0
                nw55_ = C0()
                nw55_.ctor__((d_371_v38_).intersection(d_371_v38_), (d_329_v0_) + (73))
                d_372_v39_ = nw55_
                d_373_v40_: _dafny.Map
                d_373_v40_ = _dafny.Map({(p1)[default__.safeIndex(7, (p1).length(0))]: p2})
                (d_372_v39_).f12 = (default__.safeModuloInt(144, len(_dafny.SeqWithoutIsStrInference([d_372_v39_.f12]))) if p2 else (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aajbnx"))).set(default__.safeIndex(len(d_373_v40_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aajbnx")))), d_364_v32_))) - (d_329_v0_))
                index31_ = default__.safeIndex(7, (p1).length(0))
                (p1)[index31_] = False
                index32_ = default__.safeIndex(7, (p1).length(0))
                (p1)[index32_] = default__.fm1(d_332_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgnkdwmjc"))), globalState)
                d_374_v41_: _dafny.Map
                d_374_v41_ = _dafny.Map({d_372_v39_.f12: d_367_v35_})
                d_374_v41_ = (d_374_v41_).set(d_372_v39_.f12, d_367_v35_)
            elif True:
                d_375_v42_: _dafny.Array
                nw56_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_375_v42_ = nw56_
                d_376_v43_: _dafny.Array
                nw57_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_376_v43_ = nw57_
                index33_ = default__.safeIndex(417, (d_375_v42_).length(0))
                (d_375_v42_)[index33_] = d_376_v43_
                index34_ = default__.safeIndex(417, (d_375_v42_).length(0))
                (d_375_v42_)[index34_] = d_376_v43_
                index35_ = default__.safeIndex(343, (d_334_v5_).length(0))
                (d_334_v5_)[index35_] = len(_dafny.SeqWithoutIsStrInference([p2, p2]))
                index36_ = default__.safeIndex(343, (d_334_v5_).length(0))
                (d_334_v5_)[index36_] = (len(d_330_v1_) if default__.fm1(d_332_v3_, len(d_366_v34_), globalState) else d_329_v0_)
                d_377_v44_: _dafny.Array
                def lambda12_(d_378_v35_):
                    def lambda13_(d_379_i2_):
                        return (d_379_i2_) - ((D1_DC1(len(d_378_v35_))).cf1)

                    return lambda13_

                init7_ = lambda12_(d_367_v35_)
                nw58_ = _dafny.Array(None, 10)
                for i0_7_ in range(nw58_.length(0)):
                    nw58_[i0_7_] = init7_(i0_7_)
                d_377_v44_ = nw58_
                d_380_v45_: _dafny.Map
                d_380_v45_ = _dafny.Map({(p1)[default__.safeIndex(7, (p1).length(0))]: d_377_v44_})
                d_380_v45_ = (d_380_v45_).set(default__.fm1(d_332_v3_, (d_334_v5_)[default__.safeIndex(343, (d_334_v5_).length(0))], globalState), ((d_380_v45_)[not(p2)] if (not(p2)) in (d_380_v45_) else d_377_v44_))
                d_381_v47_: _dafny.Set
                d_381_v47_ = _dafny.Set({d_330_v1_})
                d_382_v48_: C0
                nw59_ = C0()
                nw59_.ctor__(d_381_v47_, (d_334_v5_)[default__.safeIndex(343, (d_334_v5_).length(0))])
                d_382_v48_ = nw59_
                d_383_v49_: _dafny.Map
                d_383_v49_ = _dafny.Map({d_382_v48_: p2})
                index37_ = default__.safeIndex(7, (p1).length(0))
                index38_ = default__.safeIndex(343, (d_334_v5_).length(0))
                index39_ = default__.safeIndex(343, (d_334_v5_).length(0))
                def iife33_():
                    coll17_ = _dafny.Map()
                    compr_17_: int
                    for compr_17_ in _dafny.IntegerRange(-724, 374):
                        d_384_v46_: int = compr_17_
                        if ((-724) <= (d_384_v46_)) and ((d_384_v46_) < (374)):
                            coll17_[default__.safeDivisionInt(d_384_v46_, d_329_v0_)] = (p1)[default__.safeIndex(7, (p1).length(0))]
                    return _dafny.Map(coll17_)
                rhs71_ = (default__.fm15(p2, globalState) if True else (d_368_v36_) | (iife33_()
                ))
                rhs72_ = (d_333_v4_)[default__.safeIndex(len((d_383_v49_) | (_dafny.Map({d_382_v48_: (p1)[default__.safeIndex(7, (p1).length(0))]}))), len(d_333_v4_))]
                rhs73_ = (d_382_v48_.f12) * ((d_382_v48_.f12) + ((d_334_v5_)[default__.safeIndex(343, (d_334_v5_).length(0))]))
                rhs74_ = (d_334_v5_)[default__.safeIndex(343, (d_334_v5_).length(0))]
                lhs34_ = p1
                lhs35_ = default__.safeIndex(7, (p1).length(0))
                lhs36_ = d_334_v5_
                lhs37_ = default__.safeIndex(343, (d_334_v5_).length(0))
                lhs38_ = d_334_v5_
                lhs39_ = default__.safeIndex(343, (d_334_v5_).length(0))
                d_368_v36_ = rhs71_
                lhs34_[lhs35_] = rhs72_
                lhs36_[lhs37_] = rhs73_
                lhs38_[lhs39_] = rhs74_
                d_385_v50_: _dafny.Map
                d_385_v50_ = _dafny.Map({(p1)[default__.safeIndex(7, (p1).length(0))]: 521})
                r1 = (True if (d_385_v50_) != (d_385_v50_) else (p1)[default__.safeIndex(7, (p1).length(0))])
            d_329_v0_ = default__.safeModuloInt(d_329_v0_, (0) - (d_329_v0_))
            d_329_v0_ = default__.safeModuloInt(d_329_v0_, d_329_v0_)
        elif True:
            d_386_v51_: _dafny.Array
            def lambda14_(d_387_i3_):
                return D3_DC8()

            init8_ = lambda14_
            nw60_ = _dafny.Array(None, 9)
            for i0_8_ in range(nw60_.length(0)):
                nw60_[i0_8_] = init8_(i0_8_)
            d_386_v51_ = nw60_
            d_386_v51_ = d_386_v51_
            d_388_v52_: _dafny.Map
            d_388_v52_ = _dafny.Map({p2: d_329_v0_})
            d_388_v52_ = d_388_v52_
            d_389_v53_: _dafny.Array
            def lambda15_(d_390_i4_):
                return _dafny.CodePoint('v')

            init9_ = lambda15_
            nw61_ = _dafny.Array(None, 22)
            for i0_9_ in range(nw61_.length(0)):
                nw61_[i0_9_] = init9_(i0_9_)
            d_389_v53_ = nw61_
            d_391_v54_: str
            d_391_v54_ = _dafny.CodePoint('v')
            index40_ = default__.safeIndex(286, (d_389_v53_).length(0))
            (d_389_v53_)[index40_] = d_391_v54_
            index41_ = default__.safeIndex(286, (d_389_v53_).length(0))
            (d_389_v53_)[index41_] = d_391_v54_
            d_392_v55_: _dafny.MultiSet
            d_392_v55_ = _dafny.MultiSet([p2, p2, p2, True, p2])
            d_393_v56_: D1
            d_393_v56_ = D1_DC4((d_392_v55_).cardinality, p2, p2, p2, p2)
            d_394_v57_: D1
            d_394_v57_ = D1_DC5(d_393_v56_)
            d_395_v58_: _dafny.Map
            d_395_v58_ = _dafny.Map({d_391_v54_: p2})
            d_396_v59_: C0
            nw62_ = C0()
            nw62_.ctor__(default__.fm16(d_395_v58_, p2, globalState), d_329_v0_)
            d_396_v59_ = nw62_
            d_397_v60_: _dafny.Map
            d_397_v60_ = _dafny.Map({(d_394_v57_ if default__.fm1(d_332_v3_, d_329_v0_, globalState) else d_394_v57_): d_396_v59_})
            d_397_v60_ = d_397_v60_
            if not((not(p2)) == (p2)):
                d_398_v61_: _dafny.Map
                d_398_v61_ = _dafny.Map({p1: d_396_v59_.f12})
                d_399_v62_: _dafny.Seq
                d_399_v62_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_400_v63_: _dafny.Seq
                d_400_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuorgwewj"))
                d_401_v64_: _dafny.Map
                d_401_v64_ = _dafny.Map({d_396_v59_.f12: p2})
                d_402_v65_: _dafny.Map
                d_402_v65_ = _dafny.Map({(_dafny.MultiSet([p2, (d_396_v59_).fm8(p2, p2, d_401_v64_, globalState), p2, p2, p2])).cardinality: p2})
                d_403_v66_: _dafny.Seq
                d_403_v66_ = _dafny.SeqWithoutIsStrInference([d_401_v64_])
                d_404_v67_: _dafny.Map
                d_404_v67_ = _dafny.Map({p1: p2})
                d_405_v68_: D1
                d_405_v68_ = D1_DC4(len(d_330_v1_), p2, ((d_404_v67_)[p1] if (p1) in (d_404_v67_) else p2), p2, (d_396_v59_).fm8(p2, p2, d_402_v65_, globalState))
                rhs75_ = len((d_398_v61_) | (_dafny.Map({(d_399_v62_)[default__.safeIndex(len(d_400_v63_), len(d_399_v62_))]: d_396_v59_.f12})))
                rhs76_ = (p2 if p2 else (len(d_403_v66_)) != (len(_dafny.SeqWithoutIsStrInference([d_405_v68_, d_405_v68_]))))
                rhs77_ = d_334_v5_
                lhs40_ = globalState
                d_329_v0_ = rhs75_
                lhs40_.f2 = rhs76_
                d_334_v5_ = rhs77_
                d_406_v69_: _dafny.Map
                d_406_v69_ = _dafny.Map({(d_333_v4_).set(default__.safeIndex(((d_332_v3_)[d_329_v0_] if (d_329_v0_) in (d_332_v3_) else d_396_v59_.f12), len(d_333_v4_)), p2): not(p2)})
                d_406_v69_ = (d_406_v69_).set((_dafny.SeqWithoutIsStrInference([p2])) + (d_333_v4_), (p2) and (p2))
                d_407_v70_: _dafny.Set
                d_407_v70_ = _dafny.Set({p1, p1, p1})
                (globalState).f2 = not (p2) or ((d_396_v59_.f12) >= ((0) - (len(d_407_v70_))))
                (d_396_v59_).f12 = (0) - (d_329_v0_)
                d_408_v71_: _dafny.Map
                d_408_v71_ = _dafny.Map({d_330_v1_: d_396_v59_.f12})
                d_409_v72_: _dafny.Map
                d_409_v72_ = _dafny.Map({d_329_v0_: d_396_v59_.f12})
                d_408_v71_ = (_dafny.Map({d_330_v1_: d_396_v59_.f12})) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([d_396_v59_.f12, ((d_409_v72_)[len(default__.fm15(p2, globalState))] if (len(default__.fm15(p2, globalState))) in (d_409_v72_) else d_329_v0_)]): d_329_v0_})) | (d_408_v71_))
            elif True:
                d_410_v73_: _dafny.Set
                d_410_v73_ = _dafny.Set({d_396_v59_.f12, d_329_v0_})
                d_411_v74_: _dafny.Set
                d_411_v74_ = d_410_v73_
                d_412_v75_: _dafny.Map
                d_412_v75_ = _dafny.Map({D3_DC8(): d_329_v0_})
                d_413_v76_: _dafny.Array
                def lambda16_(d_414_v73_):
                    def lambda17_(d_415_i5_):
                        return d_414_v73_

                    return lambda17_

                init10_ = lambda16_(d_410_v73_)
                nw63_ = _dafny.Array(None, 12)
                for i0_10_ in range(nw63_.length(0)):
                    nw63_[i0_10_] = init10_(i0_10_)
                d_413_v76_ = nw63_
                index42_ = default__.safeIndex(404, (d_413_v76_).length(0))
                (d_413_v76_)[index42_] = d_410_v73_
                d_416_v77_: _dafny.Seq
                d_416_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrhtmhjuf"))
                d_417_v78_: D1
                d_417_v78_ = D1_DC3(not(default__.fm1(d_332_v3_, d_396_v59_.f12, globalState)), d_416_v77_, 404)
                d_418_v79_: _dafny.Map
                d_418_v79_ = _dafny.Map({d_416_v77_: p2})
                d_419_v80_: _dafny.Map
                d_419_v80_ = _dafny.Map({d_396_v59_.f12: d_416_v77_})
                index43_ = default__.safeIndex(404, (d_413_v76_).length(0))
                rhs78_ = d_411_v74_
                rhs79_ = default__.fm17(d_396_v59_.f12, d_418_v79_, p2, ((d_392_v55_)[p2] if (p2) in (d_392_v55_) else len(d_419_v80_)), globalState)
                rhs80_ = (0) - ((len((d_331_v2_)[default__.safeIndex(d_396_v59_.f12, len(d_331_v2_))])) - ((d_396_v59_.f12 if p2 else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))))))
                rhs81_ = (d_410_v73_).intersection(d_410_v73_)
                rhs82_ = d_417_v78_
                lhs41_ = d_396_v59_
                lhs42_ = d_413_v76_
                lhs43_ = default__.safeIndex(404, (d_413_v76_).length(0))
                d_411_v74_ = rhs78_
                d_412_v75_ = rhs79_
                lhs41_.f12 = rhs80_
                lhs42_[lhs43_] = rhs81_
                d_417_v78_ = rhs82_
                d_330_v1_ = _dafny.SeqWithoutIsStrInference([d_396_v59_.f12])
                d_420_v81_: _dafny.Map
                d_420_v81_ = _dafny.Map({len(_dafny.Map({not(p2): False})): p2})
                default__.m0(((d_420_v81_)[701] if (701) in (d_420_v81_) else p2), globalState)
                rhs83_ = (217) * (default__.fm0(globalState))
                rhs84_ = p2
                rhs85_ = -297
                lhs44_ = d_396_v59_
                lhs44_.f12 = rhs83_
                r0 = rhs84_
                d_329_v0_ = rhs85_
                d_421_v82_: _dafny.Map
                d_421_v82_ = _dafny.Map({False: d_396_v59_})
                d_421_v82_ = _dafny.Map({p2: d_396_v59_})
        d_422_v83_: str
        d_422_v83_ = _dafny.CodePoint('x')
        d_423_v84_: _dafny.MultiSet
        d_423_v84_ = _dafny.MultiSet([d_422_v83_, d_422_v83_])
        def iife34_():
            coll18_ = _dafny.Set()
            compr_18_: str
            for compr_18_ in (d_423_v84_).Elements:
                d_424_v85_: str = compr_18_
                if (d_424_v85_) in (d_423_v84_):
                    coll18_ = coll18_.union(_dafny.Set([d_424_v85_]))
            return _dafny.Set(coll18_)
        r0 = (p2 if not((565) >= (d_329_v0_)) else (len(iife34_()
        )) != (d_329_v0_))
        r1 = p2
        return r0, r1


class C2(T1, T0):
    def  __init__(self):
        self._f9: bool = False
        self._f14: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f9(self):
        return self._f9
    def ctor__(self, f14, f9):
        (self)._f14 = f14
        (self)._f9 = f9

    def fm2(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([(self).f14])) + (_dafny.SeqWithoutIsStrInference([(self).f14, (self).f14, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_425_i0_ in range(default__.abs(460))]))]))

    def m1(self, p0, p1, p2, p3, globalState):
        d_426_v0_: int
        d_426_v0_ = 884
        d_427_v1_: _dafny.Seq
        d_427_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqsey"))
        d_428_v2_: _dafny.Seq
        d_428_v2_ = _dafny.SeqWithoutIsStrInference([p0])
        d_429_v3_: _dafny.Set
        d_429_v3_ = _dafny.Set({d_428_v2_})
        d_430_v4_: C0
        nw64_ = C0()
        nw64_.ctor__(d_429_v3_, default__.fm0(globalState))
        d_430_v4_ = nw64_
        d_431_v5_: _dafny.Seq
        d_431_v5_ = _dafny.SeqWithoutIsStrInference([d_430_v4_, d_430_v4_])
        d_432_v6_: str
        d_432_v6_ = _dafny.CodePoint('o')
        d_433_v7_: _dafny.MultiSet
        d_433_v7_ = _dafny.MultiSet([len(((d_427_v1_).set(default__.safeIndex((0) - (len(((d_431_v5_).set(default__.safeIndex(p0, len(d_431_v5_)), d_430_v4_)).set(default__.safeIndex(d_430_v4_.f12, len((d_431_v5_).set(default__.safeIndex(p0, len(d_431_v5_)), d_430_v4_))), d_430_v4_))), len(d_427_v1_)), d_432_v6_)).set(default__.safeIndex(len(_dafny.Map({(self).f14: p1})), len((d_427_v1_).set(default__.safeIndex((0) - (len(((d_431_v5_).set(default__.safeIndex(p0, len(d_431_v5_)), d_430_v4_)).set(default__.safeIndex(d_430_v4_.f12, len((d_431_v5_).set(default__.safeIndex(p0, len(d_431_v5_)), d_430_v4_))), d_430_v4_))), len(d_427_v1_)), d_432_v6_))), d_432_v6_))])
        d_426_v0_ = (((d_433_v7_)[(self).f14] if ((self).f14) in (d_433_v7_) else (self).f14) if p2 else len(d_427_v1_))
        (globalState).f2 = not(p2)
        if p1:
            d_434_v8_: _dafny.Array
            nw65_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_434_v8_ = nw65_
            d_435_v9_: _dafny.Array
            d_435_v9_ = d_434_v8_
            d_436_v10_: _dafny.Array
            nw66_ = _dafny.Array(False, 25)
            d_436_v10_ = nw66_
            index44_ = default__.safeIndex(992, (d_436_v10_).length(0))
            (d_436_v10_)[index44_] = (self).f9
            d_437_v11_: _dafny.Map
            d_437_v11_ = _dafny.Map({p1: d_432_v6_})
            index45_ = default__.safeIndex(992, (d_436_v10_).length(0))
            rhs86_ = d_435_v9_
            rhs87_ = (d_437_v11_) != (d_437_v11_)
            lhs45_ = d_436_v10_
            lhs46_ = default__.safeIndex(992, (d_436_v10_).length(0))
            d_435_v9_ = rhs86_
            lhs45_[lhs46_] = rhs87_
            d_438_v12_: _dafny.Map
            d_438_v12_ = _dafny.Map({(self).f14: ((self).f9 if default__.fm1(d_433_v7_, 620, globalState) else (d_436_v10_)[default__.safeIndex(992, (d_436_v10_).length(0))])})
            d_438_v12_ = d_438_v12_
            d_439_v13_: _dafny.Set
            d_439_v13_ = _dafny.Set({d_426_v0_})
            (globalState).f2 = (d_439_v13_) == (d_439_v13_)
            if True:
                (d_430_v4_).f12 = d_430_v4_.f12
                d_440_v14_: _dafny.Set
                d_440_v14_ = _dafny.Set({d_427_v1_})
                d_441_v15_: _dafny.Array
                nw67_ = _dafny.Array(None, 2)
                nw67_[int(0)] = len(d_440_v14_)
                nw67_[int(1)] = default__.safeModuloInt(d_430_v4_.f12, 679)
                d_441_v15_ = nw67_
                index46_ = default__.safeIndex(181, (d_441_v15_).length(0))
                (d_441_v15_)[index46_] = len(_dafny.SeqWithoutIsStrInference([(0) - (d_430_v4_.f12), d_430_v4_.f12, len(d_428_v2_), d_430_v4_.f12, d_430_v4_.f12]))
                index47_ = default__.safeIndex(181, (d_441_v15_).length(0))
                (d_441_v15_)[index47_] = d_430_v4_.f12
                d_442_v16_: D1
                d_442_v16_ = D1_DC2((self).f14, p2, d_430_v4_.f12, d_430_v4_.f12)
                d_442_v16_ = D1_DC2(718, (d_436_v10_)[default__.safeIndex(992, (d_436_v10_).length(0))], d_430_v4_.f12, (d_430_v4_.f12) * ((0) - (d_426_v0_)))
                d_443_v17_: _dafny.Array
                nw68_ = _dafny.Array(None, 23)
                d_443_v17_ = nw68_
                index48_ = default__.safeIndex(246, (d_443_v17_).length(0))
                (d_443_v17_)[index48_] = d_430_v4_
                d_444_v18_: _dafny.MultiSet
                d_444_v18_ = _dafny.MultiSet([p1, (d_427_v1_) == (d_427_v1_)])
                d_445_v19_: _dafny.Map
                d_445_v19_ = _dafny.Map({d_433_v7_: (d_436_v10_)[default__.safeIndex(992, (d_436_v10_).length(0))]})
                d_446_v20_: D7
                d_446_v20_ = D7_DC18(d_445_v19_)
                d_447_v21_: _dafny.Seq
                d_447_v21_ = _dafny.SeqWithoutIsStrInference([(d_436_v10_)[default__.safeIndex(992, (d_436_v10_).length(0))], p1, p1, (self).f9, False])
                index49_ = default__.safeIndex(992, (d_436_v10_).length(0))
                index50_ = default__.safeIndex(181, (d_441_v15_).length(0))
                index51_ = default__.safeIndex(246, (d_443_v17_).length(0))
                rhs88_ = (d_436_v10_)[default__.safeIndex(992, (d_436_v10_).length(0))]
                rhs89_ = d_430_v4_.f12
                rhs90_ = d_430_v4_
                rhs91_ = ((d_444_v18_).intersection(d_444_v18_)) | ((_dafny.MultiSet(d_447_v21_)) - (_dafny.MultiSet([(self).f9])))
                rhs92_ = d_446_v20_
                lhs47_ = d_436_v10_
                lhs48_ = default__.safeIndex(992, (d_436_v10_).length(0))
                lhs49_ = d_441_v15_
                lhs50_ = default__.safeIndex(181, (d_441_v15_).length(0))
                lhs51_ = d_443_v17_
                lhs52_ = default__.safeIndex(246, (d_443_v17_).length(0))
                lhs47_[lhs48_] = rhs88_
                lhs49_[lhs50_] = rhs89_
                lhs51_[lhs52_] = rhs90_
                d_444_v18_ = rhs91_
                d_446_v20_ = rhs92_
                d_447_v21_ = d_447_v21_
            elif True:
                d_448_v22_: _dafny.Map
                d_448_v22_ = _dafny.Map({d_427_v1_: d_436_v10_})
                d_448_v22_ = d_448_v22_
                d_449_v23_: _dafny.Array
                nw69_ = _dafny.Array(_dafny.Map({}), 2)
                d_449_v23_ = nw69_
                nw70_ = _dafny.Array(_dafny.Map({}), 1)
                d_449_v23_ = nw70_
                d_450_v24_: _dafny.Set
                d_450_v24_ = _dafny.Set({not(True), (_dafny.MultiSet([(0) - (d_430_v4_.f12)])).isdisjoint(d_433_v7_), p1, p1, p1})
                d_450_v24_ = d_450_v24_
                (globalState).f2 = p1
                d_451_v25_: _dafny.Map
                d_451_v25_ = _dafny.Map({p1: p2})
                d_452_v26_: _dafny.Array
                nw71_ = _dafny.Array(None, 21)
                nw71_[int(0)] = len((d_428_v2_).set(default__.safeIndex(d_430_v4_.f12, len(d_428_v2_)), len(d_428_v2_)))
                nw71_[int(1)] = d_426_v0_
                nw71_[int(2)] = -78
                nw71_[int(3)] = d_430_v4_.f12
                nw71_[int(4)] = d_426_v0_
                nw71_[int(5)] = p0
                nw71_[int(6)] = p0
                nw71_[int(7)] = d_430_v4_.f12
                nw71_[int(8)] = default__.fm0(globalState)
                nw71_[int(9)] = d_430_v4_.f12
                nw71_[int(10)] = d_426_v0_
                nw71_[int(11)] = d_430_v4_.f12
                nw71_[int(12)] = (self).f14
                nw71_[int(13)] = len(d_451_v25_)
                nw71_[int(14)] = d_430_v4_.f12
                nw71_[int(15)] = (self).f14
                nw71_[int(16)] = p0
                nw71_[int(17)] = p0
                nw71_[int(18)] = d_426_v0_
                nw71_[int(19)] = d_430_v4_.f12
                nw71_[int(20)] = 976
                d_452_v26_ = nw71_
                d_453_v27_: _dafny.MultiSet
                d_453_v27_ = _dafny.MultiSet([d_452_v26_, d_452_v26_])
                (self).m8(not((self).f9), d_453_v27_, globalState)
            (d_430_v4_).f12 = d_430_v4_.f12
        elif True:
            d_454_v28_: _dafny.Seq
            d_454_v28_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_455_v29_: _dafny.Array
            nw72_ = _dafny.Array(None, 21)
            nw72_[int(0)] = (self).f14
            nw72_[int(1)] = p0
            nw72_[int(2)] = (0) - (len(d_454_v28_))
            nw72_[int(3)] = d_430_v4_.f12
            nw72_[int(4)] = d_430_v4_.f12
            nw72_[int(5)] = (self).f14
            nw72_[int(6)] = p0
            nw72_[int(7)] = d_430_v4_.f12
            nw72_[int(8)] = len(d_427_v1_)
            nw72_[int(9)] = d_430_v4_.f12
            nw72_[int(10)] = (self).f14
            nw72_[int(11)] = len(d_428_v2_)
            nw72_[int(12)] = d_426_v0_
            nw72_[int(13)] = d_430_v4_.f12
            nw72_[int(14)] = default__.fm0(globalState)
            nw72_[int(15)] = 958
            nw72_[int(16)] = d_430_v4_.f12
            nw72_[int(17)] = len(_dafny.SeqWithoutIsStrInference([p0]))
            nw72_[int(18)] = d_430_v4_.f12
            nw72_[int(19)] = len(d_427_v1_)
            nw72_[int(20)] = 55
            d_455_v29_ = nw72_
            d_456_v30_: _dafny.Seq
            d_456_v30_ = _dafny.SeqWithoutIsStrInference([d_455_v29_])
            (d_430_v4_).f12 = (len(default__.fm24(d_430_v4_.f12, not(True), (self).f9, p1, globalState)) if (self).f9 else (len(d_456_v30_)) * (d_430_v4_.f12))
            (globalState).f2 = (self).f9
            d_457_v31_: _dafny.MultiSet
            d_457_v31_ = _dafny.MultiSet([p2, p2])
            d_458_v32_: D1
            d_458_v32_ = D1_DC3((self).f9, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doecrfd")), p0)
            rhs93_ = d_427_v1_
            rhs94_ = ((d_428_v2_) + ((self).fm2(d_430_v4_.f12, p1, (self).f9, (d_457_v31_).cardinality, globalState))).set(default__.safeIndex((len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))).set(default__.safeIndex(533, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))), d_432_v6_)).set(default__.safeIndex(-52, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))).set(default__.safeIndex(533, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))), d_432_v6_))), d_432_v6_))) - (d_430_v4_.f12), len((d_428_v2_) + ((self).fm2(d_430_v4_.f12, p1, (self).f9, (d_457_v31_).cardinality, globalState)))), (self).f14)
            rhs95_ = not(p1)
            rhs96_ = ((d_458_v32_).cf7) + (d_427_v1_)
            lhs53_ = globalState
            d_427_v1_ = rhs93_
            d_428_v2_ = rhs94_
            lhs53_.f2 = rhs95_
            d_427_v1_ = rhs96_
            d_459_v33_: D4
            d_459_v33_ = D4_DC12((self).f9, p1, (self).f14, len(d_428_v2_), (self).f9)
            d_460_v34_: D1
            d_460_v34_ = D1_DC2(default__.safeModuloInt((self).f14, d_430_v4_.f12), (d_454_v28_)[default__.safeIndex(default__.fm0(globalState), len(d_454_v28_))], ((0) - (d_430_v4_.f12) if p1 else d_426_v0_), ((self).f14) - ((d_459_v33_).cf23))
            d_461_v35_: _dafny.Map
            d_461_v35_ = _dafny.Map({p0: p1})
            d_462_v36_: _dafny.Set
            d_462_v36_ = _dafny.Set({len(d_461_v35_)})
            d_463_v37_: _dafny.MultiSet
            d_463_v37_ = _dafny.MultiSet([d_455_v29_, d_455_v29_])
            pat_let_tv6_ = d_463_v37_
            pat_let_tv7_ = d_463_v37_
            pat_let_tv8_ = p1
            def iife35_(_pat_let8_0):
                def iife36_(d_464_dt__update__tmp_h0_):
                    def iife37_(_pat_let9_0):
                        def iife38_(d_465_dt__update_hcf3_h0_):
                            def iife39_(_pat_let10_0):
                                def iife40_(d_466_dt__update_hcf4_h0_):
                                    return D1_DC2((d_464_dt__update__tmp_h0_).cf2, d_465_dt__update_hcf3_h0_, d_466_dt__update_hcf4_h0_, (d_464_dt__update__tmp_h0_).cf5)
                                return iife40_(_pat_let10_0)
                            return iife39_(len(_dafny.SeqWithoutIsStrInference([(self).f9, True, pat_let_tv8_])))
                        return iife38_(_pat_let9_0)
                    return iife37_((pat_let_tv6_).ispropersubset(pat_let_tv7_))
                return iife36_(_pat_let8_0)
            d_460_v34_ = iife35_(default__.fm25(len(d_462_v36_), p1, p1, d_433_v7_, globalState))
            d_467_v38_: _dafny.Map
            d_467_v38_ = _dafny.Map({(p2) == (not(p1)): (self).f14})
            d_467_v38_ = (d_467_v38_).set(p1, default__.fm0(globalState))
        d_468_v39_: _dafny.Array
        def lambda18_(d_469_p2_):
            def lambda19_(d_470_i0_):
                return d_469_p2_

            return lambda19_

        init11_ = lambda18_(p2)
        nw73_ = _dafny.Array(None, 2)
        for i0_11_ in range(nw73_.length(0)):
            nw73_[i0_11_] = init11_(i0_11_)
        d_468_v39_ = nw73_
        index52_ = default__.safeIndex(459, (d_468_v39_).length(0))
        (d_468_v39_)[index52_] = p1
        index53_ = default__.safeIndex(459, (d_468_v39_).length(0))
        (d_468_v39_)[index53_] = not((self).f9)
        d_471_v40_: _dafny.Array
        nw74_ = _dafny.Array(None, 1)
        d_471_v40_ = nw74_
        d_472_v41_: D4
        d_472_v41_ = D4_DC11(d_471_v40_)
        d_473_v42_: _dafny.Array
        nw75_ = _dafny.Array(None, 7)
        nw75_[int(0)] = d_471_v40_
        nw75_[int(1)] = (d_472_v41_).cf19
        nw75_[int(2)] = d_471_v40_
        nw75_[int(3)] = (d_472_v41_).cf19
        nw75_[int(4)] = d_471_v40_
        nw75_[int(5)] = d_471_v40_
        nw75_[int(6)] = d_471_v40_
        d_473_v42_ = nw75_
        source6_ = d_473_v42_
        d_474___mcc_h0_ = source6_
        d_475_cf30_ = d_474___mcc_h0_
        (d_430_v4_).f12 = d_426_v0_
        d_476_v43_: D4
        d_476_v43_ = D4_DC12(not((self).f9), p2, (0) - (d_430_v4_.f12), (self).f14, False)
        (d_430_v4_).f12 = ((d_476_v43_).cf23) + (default__.safeModuloInt(d_430_v4_.f12, len(d_427_v1_)))
        d_477_v44_: _dafny.Seq
        d_477_v44_ = _dafny.SeqWithoutIsStrInference([d_430_v4_.f11, d_430_v4_.f11])
        d_478_v46_: C0
        nw76_ = C0()
        def iife41_():
            coll19_ = _dafny.Set()
            compr_19_: _dafny.Seq
            for compr_19_ in ((d_477_v44_)[default__.safeIndex(len(d_427_v1_), len(d_477_v44_))]).Elements:
                d_479_v45_: _dafny.Seq = compr_19_
                if (d_479_v45_) in ((d_477_v44_)[default__.safeIndex(len(d_427_v1_), len(d_477_v44_))]):
                    coll19_ = coll19_.union(_dafny.Set([d_479_v45_]))
            return _dafny.Set(coll19_)
        nw76_.ctor__(iife41_()
        , d_430_v4_.f12)
        d_478_v46_ = nw76_
        d_480_v47_: _dafny.Map
        d_480_v47_ = _dafny.Map({d_432_v6_: p2})
        d_481_v48_: _dafny.Set
        d_481_v48_ = _dafny.Set({d_480_v47_, d_480_v47_})
        index54_ = default__.safeIndex(459, (d_468_v39_).length(0))
        (d_468_v39_)[index54_] = not(((d_481_v48_) - (d_481_v48_)).issubset(d_481_v48_))
        d_482_v49_: _dafny.Set
        d_482_v49_ = _dafny.Set({(self).f14, (self).f14})
        d_483_v50_: _dafny.Seq
        d_483_v50_ = _dafny.SeqWithoutIsStrInference([(self).f9])
        d_484_v51_: _dafny.MultiSet
        d_484_v51_ = _dafny.MultiSet([(d_468_v39_)[default__.safeIndex(459, (d_468_v39_).length(0))]])
        d_485_v52_: _dafny.Array
        nw77_ = _dafny.Array(None, 17)
        nw77_[int(0)] = (d_426_v0_) + (len(d_483_v50_))
        nw77_[int(1)] = p0
        nw77_[int(2)] = len(d_427_v1_)
        nw77_[int(3)] = default__.safeModuloInt((d_484_v51_).cardinality, len(d_428_v2_))
        nw77_[int(4)] = p0
        nw77_[int(5)] = default__.safeModuloInt((self).f14, p0)
        nw77_[int(6)] = d_426_v0_
        nw77_[int(7)] = d_426_v0_
        nw77_[int(8)] = (d_426_v0_) + ((self).f14)
        nw77_[int(9)] = d_430_v4_.f12
        nw77_[int(10)] = (self).f14
        nw77_[int(11)] = d_426_v0_
        nw77_[int(12)] = default__.fm0(globalState)
        nw77_[int(13)] = d_426_v0_
        nw77_[int(14)] = default__.safeModuloInt(d_430_v4_.f12, d_430_v4_.f12)
        nw77_[int(15)] = d_426_v0_
        nw77_[int(16)] = ((self).f14 if (d_468_v39_)[default__.safeIndex(459, (d_468_v39_).length(0))] else (0) - (default__.fm0(globalState)))
        d_485_v52_ = nw77_
        d_486_v53_: C1
        nw78_ = C1()
        nw78_.ctor__(default__.fm26(globalState))
        d_486_v53_ = nw78_
        d_487_v54_: _dafny.Map
        d_487_v54_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_432_v6_ for d_488_i1_ in range(default__.abs(414))]): _dafny.Map({d_486_v53_: d_430_v4_.f12})})
        d_489_v55_: _dafny.Map
        d_489_v55_ = _dafny.Map({d_486_v53_: p0})
        index55_ = default__.safeIndex(751, (d_485_v52_).length(0))
        (d_485_v52_)[index55_] = len(((d_487_v54_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nukgadch"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nukgadch"))) in (d_487_v54_) else d_489_v55_))
        d_490_v56_: _dafny.Seq
        d_490_v56_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_426_v0_}), d_482_v49_, d_482_v49_])
        index56_ = default__.safeIndex(751, (d_485_v52_).length(0))
        rhs97_ = (d_482_v49_) - ((d_490_v56_)[default__.safeIndex(d_426_v0_, len(d_490_v56_))])
        rhs98_ = p0
        lhs54_ = d_485_v52_
        lhs55_ = default__.safeIndex(751, (d_485_v52_).length(0))
        d_482_v49_ = rhs97_
        lhs54_[lhs55_] = rhs98_

    def m8(self, p0, p1, globalState):
        d_491_v0_: _dafny.MultiSet
        d_491_v0_ = _dafny.MultiSet([(self).f14, (self).f14])
        d_492_v1_: _dafny.Seq
        d_492_v1_ = _dafny.SeqWithoutIsStrInference([p0])
        d_493_v2_: _dafny.Set
        d_493_v2_ = _dafny.Set({(self).f14})
        d_494_v3_: _dafny.Map
        d_494_v3_ = _dafny.Map({p0: (self).f9})
        d_495_v4_: _dafny.MultiSet
        d_495_v4_ = _dafny.MultiSet([((d_494_v3_)[(self).f9] if ((self).f9) in (d_494_v3_) else (self).f9), (self).f9])
        d_496_v5_: _dafny.Array
        nw79_ = _dafny.Array(None, 3)
        nw79_[int(0)] = (d_491_v0_).ispropersubset(d_491_v0_)
        nw79_[int(1)] = ((self).f14) >= ((self).f14)
        nw79_[int(2)] = ((_dafny.MultiSet((d_492_v1_).set(default__.safeIndex((self).f14, len(d_492_v1_)), (self).f9))).set((self).f9, default__.abs(len(d_493_v2_)))).isdisjoint(d_495_v4_)
        d_496_v5_ = nw79_
        index57_ = default__.safeIndex(557, (d_496_v5_).length(0))
        (d_496_v5_)[index57_] = (self).f9
        d_497_v6_: _dafny.Array
        nw80_ = _dafny.Array(int(0), 7)
        d_497_v6_ = nw80_
        index58_ = default__.safeIndex(781, (d_497_v6_).length(0))
        (d_497_v6_)[index58_] = -89
        d_498_v7_: _dafny.Array
        nw81_ = _dafny.Array(_dafny.MultiSet({}), 12)
        d_498_v7_ = nw81_
        d_499_v8_: _dafny.Seq
        d_499_v8_ = _dafny.SeqWithoutIsStrInference([d_498_v7_])
        d_500_v9_: _dafny.Array
        nw82_ = _dafny.Array(None, 16)
        nw82_[int(0)] = d_498_v7_
        nw82_[int(1)] = d_498_v7_
        nw82_[int(2)] = d_498_v7_
        nw82_[int(3)] = d_498_v7_
        nw82_[int(4)] = d_498_v7_
        nw82_[int(5)] = d_498_v7_
        nw82_[int(6)] = d_498_v7_
        nw82_[int(7)] = d_498_v7_
        nw82_[int(8)] = (d_499_v8_)[default__.safeIndex((self).f14, len(d_499_v8_))]
        nw82_[int(9)] = d_498_v7_
        nw82_[int(10)] = d_498_v7_
        nw82_[int(11)] = d_498_v7_
        nw82_[int(12)] = d_498_v7_
        nw82_[int(13)] = d_498_v7_
        nw82_[int(14)] = d_498_v7_
        nw82_[int(15)] = d_498_v7_
        d_500_v9_ = nw82_
        index59_ = default__.safeIndex(218, (d_500_v9_).length(0))
        (d_500_v9_)[index59_] = d_498_v7_
        d_501_v10_: _dafny.Seq
        d_501_v10_ = _dafny.SeqWithoutIsStrInference([(self).f14])
        index60_ = default__.safeIndex(557, (d_496_v5_).length(0))
        index61_ = default__.safeIndex(781, (d_497_v6_).length(0))
        index62_ = default__.safeIndex(218, (d_500_v9_).length(0))
        rhs99_ = ((self).f9 if (self).f9 else p0)
        rhs100_ = default__.fm1(d_491_v0_, (0) - ((self).f14), globalState)
        rhs101_ = default__.safeDivisionInt((self).f14, (0) - ((len(d_501_v10_)) - ((self).f14)))
        rhs102_ = d_498_v7_
        lhs56_ = globalState
        lhs57_ = d_496_v5_
        lhs58_ = default__.safeIndex(557, (d_496_v5_).length(0))
        lhs59_ = d_497_v6_
        lhs60_ = default__.safeIndex(781, (d_497_v6_).length(0))
        lhs61_ = d_500_v9_
        lhs62_ = default__.safeIndex(218, (d_500_v9_).length(0))
        lhs56_.f2 = rhs99_
        lhs57_[lhs58_] = rhs100_
        lhs59_[lhs60_] = rhs101_
        lhs61_[lhs62_] = rhs102_
        index63_ = default__.safeIndex(781, (d_497_v6_).length(0))
        (d_497_v6_)[index63_] = (d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))]
        d_502_v11_: D1
        d_502_v11_ = D1_DC5(D1_DC4((self).f14, (self).f9, True, (D3_DC9((self).f9, _dafny.SeqWithoutIsStrInference([p0, p0]))).cf17, (self).f9))
        d_503_v12_: D1
        d_503_v12_ = D1_DC5(d_502_v11_)
        pat_let_tv9_ = d_502_v11_
        pat_let_tv10_ = d_502_v11_
        d_504_v13_: _dafny.Map
        def iife42_(_pat_let11_0):
            def iife43_(d_505_dt__update__tmp_h0_):
                def iife44_(_pat_let12_0):
                    def iife45_(d_506_dt__update_hcf14_h0_):
                        return D1_DC5(d_506_dt__update_hcf14_h0_)
                    return iife45_(_pat_let12_0)
                return iife44_(pat_let_tv9_)
            return iife43_(_pat_let11_0)
        def iife46_(_pat_let13_0):
            def iife47_(d_507_dt__update__tmp_h1_):
                def iife48_(_pat_let14_0):
                    def iife49_(d_508_dt__update_hcf14_h1_):
                        return D1_DC5(d_508_dt__update_hcf14_h1_)
                    return iife49_(_pat_let14_0)
                return iife48_(pat_let_tv10_)
            return iife47_(_pat_let13_0)
        d_504_v13_ = _dafny.Map({_dafny.MultiSet([iife42_(d_503_v12_), d_503_v12_, iife46_(d_503_v12_)]): d_491_v0_})
        arr0_ = (d_500_v9_)[default__.safeIndex(218, (d_500_v9_).length(0))]
        index64_ = default__.safeIndex(912, ((d_500_v9_)[default__.safeIndex(218, (d_500_v9_).length(0))]).length(0))
        arr0_[index64_] = ((d_504_v13_)[_dafny.MultiSet([d_503_v12_, d_503_v12_, d_503_v12_])] if (_dafny.MultiSet([d_503_v12_, d_503_v12_, d_503_v12_])) in (d_504_v13_) else _dafny.MultiSet([(self).f14, (d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))]]))
        arr1_ = (d_500_v9_)[default__.safeIndex(218, (d_500_v9_).length(0))]
        index65_ = default__.safeIndex(912, ((d_500_v9_)[default__.safeIndex(218, (d_500_v9_).length(0))]).length(0))
        arr1_[index65_] = (_dafny.MultiSet((d_501_v10_) + (_dafny.SeqWithoutIsStrInference([(self).f14])))) - (d_491_v0_)
        d_509_i0_: int
        d_509_i0_ = 0
        with _dafny.label("5"):
            while (d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))]:
                with _dafny.c_label("5"):
                    if (d_509_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_509_i0_ = (d_509_i0_) + (1)
                    d_510_v14_: _dafny.Seq
                    d_510_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                    d_511_v15_: D1
                    d_511_v15_ = D1_DC4(len(d_510_v14_), not(p0), p0, (d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))], (d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))])
                    if default__.fm1((((d_500_v9_)[default__.safeIndex(218, (d_500_v9_).length(0))])[default__.safeIndex(912, ((d_500_v9_)[default__.safeIndex(218, (d_500_v9_).length(0))]).length(0))]) | (d_491_v0_), default__.safeDivisionInt((d_511_v15_).cf9, (self).f14), globalState):
                        d_512_v16_: _dafny.Map
                        d_512_v16_ = _dafny.Map({d_491_v0_: (self).f9})
                        index66_ = default__.safeIndex(781, (d_497_v6_).length(0))
                        (d_497_v6_)[index66_] = ((d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))]) + ((0) - ((len(d_510_v14_)) - (len(d_512_v16_))))
                        d_513_v17_: D3
                        d_513_v17_ = D3_DC9((self).f9, d_492_v1_)
                        d_514_v18_: _dafny.MultiSet
                        d_514_v18_ = _dafny.MultiSet([d_496_v5_])
                        d_515_v19_: D3
                        d_515_v19_ = D3_DC7(d_514_v18_)
                        d_516_v20_: _dafny.Map
                        d_516_v20_ = _dafny.Map({(d_510_v14_) + (d_510_v14_): p0})
                        d_517_v21_: D9
                        d_517_v21_ = D9_DC23(d_516_v20_)
                        pat_let_tv11_ = d_514_v18_
                        index67_ = default__.safeIndex(781, (d_497_v6_).length(0))
                        def iife50_(_pat_let15_0):
                            def iife51_(d_518_dt__update__tmp_h2_):
                                def iife52_(_pat_let16_0):
                                    def iife53_(d_519_dt__update_hcf16_h0_):
                                        return D3_DC7(d_519_dt__update_hcf16_h0_)
                                    return iife53_(_pat_let16_0)
                                return iife52_(pat_let_tv11_)
                            return iife51_(_pat_let15_0)
                        rhs103_ = ((d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))]) * ((d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))])
                        rhs104_ = d_513_v17_
                        rhs105_ = iife50_(d_515_v19_)
                        rhs106_ = (d_517_v21_).cf40
                        lhs63_ = d_497_v6_
                        lhs64_ = default__.safeIndex(781, (d_497_v6_).length(0))
                        lhs63_[lhs64_] = rhs103_
                        d_513_v17_ = rhs104_
                        d_515_v19_ = rhs105_
                        d_516_v20_ = rhs106_
                        d_520_v22_: str
                        d_520_v22_ = _dafny.CodePoint('e')
                        d_521_v23_: _dafny.Seq
                        d_521_v23_ = _dafny.SeqWithoutIsStrInference([(d_510_v14_).set(default__.safeIndex((self).f14, len(d_510_v14_)), d_520_v22_)])
                        d_522_v24_: _dafny.Map
                        d_522_v24_ = _dafny.Map({(d_491_v0_).set((d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))], default__.abs(len(d_521_v23_))): (0) - ((d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))])})
                        d_523_v25_: _dafny.Map
                        d_523_v25_ = _dafny.Map({d_522_v24_: (self).f9})
                        d_523_v25_ = (d_523_v25_).set(_dafny.Map({((d_500_v9_)[default__.safeIndex(218, (d_500_v9_).length(0))])[default__.safeIndex(912, ((d_500_v9_)[default__.safeIndex(218, (d_500_v9_).length(0))]).length(0))]: (d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))]}), (self).f9)
                        d_524_v26_: _dafny.Array
                        nw83_ = _dafny.Array(_dafny.Set({}), 13)
                        d_524_v26_ = nw83_
                        d_525_v27_: _dafny.Map
                        d_525_v27_ = _dafny.Map({(d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))]: d_520_v22_})
                        d_526_v28_: _dafny.Set
                        d_526_v28_ = _dafny.Set({d_525_v27_, d_525_v27_, d_525_v27_})
                        index68_ = default__.safeIndex(435, (d_524_v26_).length(0))
                        (d_524_v26_)[index68_] = d_526_v28_
                        d_527_v29_: _dafny.Array
                        def lambda20_(d_528_v14_):
                            def lambda21_(d_529_i1_):
                                return d_528_v14_

                            return lambda21_

                        init12_ = lambda20_(d_510_v14_)
                        nw84_ = _dafny.Array(None, 15)
                        for i0_12_ in range(nw84_.length(0)):
                            nw84_[i0_12_] = init12_(i0_12_)
                        d_527_v29_ = nw84_
                        index69_ = default__.safeIndex(840, (d_527_v29_).length(0))
                        (d_527_v29_)[index69_] = d_510_v14_
                        index70_ = default__.safeIndex(435, (d_524_v26_).length(0))
                        index71_ = default__.safeIndex(840, (d_527_v29_).length(0))
                        index72_ = default__.safeIndex(557, (d_496_v5_).length(0))
                        rhs107_ = (d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))]
                        rhs108_ = (d_526_v28_).intersection((d_526_v28_).intersection(d_526_v28_))
                        rhs109_ = d_510_v14_
                        rhs110_ = (_dafny.CodePoint('r') if (self).f9 else d_520_v22_)
                        rhs111_ = not((p0 if (d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))] else (d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))]))
                        lhs65_ = globalState
                        lhs66_ = d_524_v26_
                        lhs67_ = default__.safeIndex(435, (d_524_v26_).length(0))
                        lhs68_ = d_527_v29_
                        lhs69_ = default__.safeIndex(840, (d_527_v29_).length(0))
                        lhs70_ = d_496_v5_
                        lhs71_ = default__.safeIndex(557, (d_496_v5_).length(0))
                        lhs65_.f2 = rhs107_
                        lhs66_[lhs67_] = rhs108_
                        lhs68_[lhs69_] = rhs109_
                        d_520_v22_ = rhs110_
                        lhs70_[lhs71_] = rhs111_
                        (globalState).f2 = (d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))]
                    elif True:
                        d_510_v14_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_530_i2_ in range(default__.abs(337))])) + (d_510_v14_)
                        d_531_v30_: str
                        d_531_v30_ = _dafny.CodePoint('w')
                        (globalState).f2 = not(((((d_510_v14_).set(default__.safeIndex(len(d_510_v14_), len(d_510_v14_)), d_531_v30_)) + (_dafny.SeqWithoutIsStrInference([d_531_v30_ for d_532_i3_ in range(default__.abs(766))]))).set(default__.safeIndex((self).f14, len(((d_510_v14_).set(default__.safeIndex(len(d_510_v14_), len(d_510_v14_)), d_531_v30_)) + (_dafny.SeqWithoutIsStrInference([d_531_v30_ for d_533_i3_ in range(default__.abs(766))])))), d_531_v30_)) == ((d_510_v14_) + (d_510_v14_)))
                        d_534_v31_: _dafny.Array
                        d_534_v31_ = d_497_v6_
                        d_535_v32_: _dafny.Map
                        d_535_v32_ = _dafny.Map({(d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))]: (self).f14})
                        index73_ = default__.safeIndex(781, (d_497_v6_).length(0))
                        rhs112_ = ((d_535_v32_)[p0] if (p0) in (d_535_v32_) else 738)
                        rhs113_ = d_534_v31_
                        lhs72_ = d_497_v6_
                        lhs73_ = default__.safeIndex(781, (d_497_v6_).length(0))
                        lhs72_[lhs73_] = rhs112_
                        d_534_v31_ = rhs113_
                        d_536_v33_: T0
                        nw85_ = C1()
                        nw85_.ctor__(_dafny.SeqWithoutIsStrInference([D1_DC2(len(_dafny.Set({p0})), p0, (d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))], (self).f14) for d_537_i4_ in range(default__.abs(324))]))
                        d_536_v33_ = nw85_
                        d_538_v34_: _dafny.MultiSet
                        d_538_v34_ = _dafny.MultiSet([d_536_v33_])
                        index74_ = default__.safeIndex(781, (d_497_v6_).length(0))
                        (d_497_v6_)[index74_] = ((self).f14 if (d_538_v34_).issubset(d_538_v34_) else (self).f14)
                        d_539_v35_: _dafny.Set
                        d_539_v35_ = _dafny.Set({d_497_v6_})
                        d_540_v36_: _dafny.Map
                        d_540_v36_ = _dafny.Map({(d_539_v35_).intersection(d_539_v35_): (self).f14})
                        index75_ = default__.safeIndex(781, (d_497_v6_).length(0))
                        (d_497_v6_)[index75_] = ((d_540_v36_)[d_539_v35_] if (d_539_v35_) in (d_540_v36_) else (d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))])
                    d_497_v6_ = d_497_v6_
                    d_541_v37_: D3
                    d_541_v37_ = D3_DC10()
                    d_541_v37_ = default__.fm27(globalState)
                    (globalState).f2 = (self).f9
                    pass
            pass
        d_542_v38_: str
        d_542_v38_ = _dafny.CodePoint('k')
        d_543_v39_: _dafny.Set
        d_543_v39_ = _dafny.Set({d_542_v38_, _dafny.CodePoint('o'), d_542_v38_, d_542_v38_, d_542_v38_})
        d_543_v39_ = d_543_v39_
        d_544_v40_: D4
        d_544_v40_ = D4_DC14((d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))], False, (d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))])
        d_545_v41_: _dafny.MultiSet
        d_545_v41_ = _dafny.MultiSet([d_542_v38_, d_542_v38_])
        d_546_v42_: D1
        d_546_v42_ = D1_DC3((self).f9, default__.fm28((d_545_v41_).cardinality, True, globalState), (d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))])
        d_547_v43_: _dafny.MultiSet
        d_547_v43_ = _dafny.MultiSet([d_546_v42_])
        index76_ = default__.safeIndex(557, (d_496_v5_).length(0))
        rhs114_ = D4_DC14(((d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))] if (self).f9 else not(p0)), (d_496_v5_)[default__.safeIndex(557, (d_496_v5_).length(0))], (d_497_v6_)[default__.safeIndex(781, (d_497_v6_).length(0))])
        rhs115_ = (_dafny.MultiSet([d_546_v42_])).set(d_546_v42_, default__.abs((self).f14))
        rhs116_ = p0
        lhs74_ = d_496_v5_
        lhs75_ = default__.safeIndex(557, (d_496_v5_).length(0))
        d_544_v40_ = rhs114_
        d_547_v43_ = rhs115_
        lhs74_[lhs75_] = rhs116_

    @property
    def f14(self):
        return self._f14

class C3(T2, T1, T0):
    def  __init__(self):
        self._f9: bool = False
        self.f16: _dafny.Array = _dafny.Array(None, 0)
        self._f15: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f9(self):
        return self._f9
    def ctor__(self, f15, f16, f9):
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f9 = f9

    def fm3(self, globalState):
        def iife54_():
            coll20_ = _dafny.Set()
            compr_20_: _dafny.Set
            for compr_20_ in (_dafny.MultiSet([_dafny.Set({True, (self).f9}), _dafny.Set({False})])).Elements:
                d_548_v0_: _dafny.Set = compr_20_
                if (d_548_v0_) in (_dafny.MultiSet([_dafny.Set({True, (self).f9}), _dafny.Set({False})])):
                    coll20_ = coll20_.union(_dafny.Set([d_548_v0_]))
            return _dafny.Set(coll20_)
        return (len((_dafny.Map({(self).f15: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kthtybr")))})) | (_dafny.Map({(self).f15: len(iife54_()
        )})))) >= (default__.safeModuloInt((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_549_i0_ in range(default__.abs(-746))]))), len(_dafny.Map({D10_DC28(): (self).f9}))))

    def fm2(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({True})))])) + ((_dafny.SeqWithoutIsStrInference([-817 for d_550_i0_ in range(default__.abs(286))])) + (_dafny.SeqWithoutIsStrInference([611])))

    def fm32(self, p0, p1, p2, p3, globalState):
        return not(not((self).f9))

    def fm33(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chkqh"))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: _dafny.Set = _dafny.Set({})
        d_551_v0_: int
        d_551_v0_ = 418
        d_552_v1_: _dafny.MultiSet
        d_552_v1_ = _dafny.MultiSet([d_551_v0_])
        d_553_v2_: _dafny.Map
        d_553_v2_ = _dafny.Map({d_552_v1_: (self).f9})
        d_554_v3_: D7
        d_554_v3_ = D7_DC18(d_553_v2_)
        d_554_v3_ = d_554_v3_
        d_555_v4_: _dafny.Array
        nw86_ = _dafny.Array(None, 11)
        nw86_[int(0)] = (self).f15
        nw86_[int(1)] = (self).f15
        nw86_[int(2)] = (self).f15
        nw86_[int(3)] = (self).f9
        nw86_[int(4)] = (self).f15
        nw86_[int(5)] = (self).f15
        nw86_[int(6)] = (self).f9
        nw86_[int(7)] = (self).f15
        nw86_[int(8)] = (self).f9
        nw86_[int(9)] = (self).f15
        nw86_[int(10)] = (self).f9
        d_555_v4_ = nw86_
        d_556_v5_: _dafny.MultiSet
        d_556_v5_ = _dafny.MultiSet([p1, p1])
        index77_ = default__.safeIndex(394, (d_555_v4_).length(0))
        (d_555_v4_)[index77_] = not(not((p1) not in (d_556_v5_)))
        index78_ = default__.safeIndex(394, (d_555_v4_).length(0))
        (d_555_v4_)[index78_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrugtgog"))) < (_dafny.SeqWithoutIsStrInference([p2 for d_557_i0_ in range(default__.abs(882))]))
        index79_ = default__.safeIndex(394, (d_555_v4_).length(0))
        (d_555_v4_)[index79_] = (self).f15
        d_558_v6_: _dafny.Set
        d_558_v6_ = _dafny.Set({d_551_v0_, d_551_v0_})
        d_559_i1_: int
        d_559_i1_ = 0
        with _dafny.label("6"):
            while (d_558_v6_).ispropersubset(d_558_v6_):
                with _dafny.c_label("6"):
                    if (d_559_i1_) >= (100):
                        raise _dafny.Break("6")
                    d_559_i1_ = (d_559_i1_) + (1)
                    d_560_v7_: _dafny.Array
                    def lambda22_(d_561_v0_):
                        def lambda23_(d_562_i2_):
                            return (d_562_i2_) - (d_561_v0_)

                        return lambda23_

                    init13_ = lambda22_(d_551_v0_)
                    nw87_ = _dafny.Array(None, 15)
                    for i0_13_ in range(nw87_.length(0)):
                        nw87_[i0_13_] = init13_(i0_13_)
                    d_560_v7_ = nw87_
                    index80_ = default__.safeIndex(527, (d_560_v7_).length(0))
                    (d_560_v7_)[index80_] = d_551_v0_
                    index81_ = default__.safeIndex(527, (d_560_v7_).length(0))
                    (d_560_v7_)[index81_] = d_551_v0_
                    d_563_v8_: D9
                    d_563_v8_ = D9_DC25(d_551_v0_, (d_560_v7_)[default__.safeIndex(527, (d_560_v7_).length(0))], d_552_v1_, _dafny.Map({(d_560_v7_)[default__.safeIndex(527, (d_560_v7_).length(0))]: p2}), -777)
                    d_564_v9_: _dafny.Map
                    d_564_v9_ = _dafny.Map({(d_560_v7_)[default__.safeIndex(527, (d_560_v7_).length(0))]: d_563_v8_})
                    d_564_v9_ = d_564_v9_
                    d_565_v10_: D9
                    d_565_v10_ = D9_DC24(False, (d_560_v7_)[default__.safeIndex(527, (d_560_v7_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prdj")))
                    d_566_v11_: _dafny.Set
                    d_566_v11_ = _dafny.Set({p1, p1, (d_565_v10_).cf43})
                    d_567_v12_: _dafny.Seq
                    d_567_v12_ = _dafny.SeqWithoutIsStrInference([len(d_566_v11_), len(_dafny.SeqWithoutIsStrInference([p2, p2]))])
                    d_568_v13_: _dafny.Map
                    d_568_v13_ = _dafny.Map({(self).f15: d_551_v0_})
                    d_569_v14_: _dafny.Seq
                    d_569_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_568_v13_)]), p3, p3, p3, _dafny.SeqWithoutIsStrInference([d_551_v0_ for d_570_i3_ in range(default__.abs(775))])])
                    d_567_v12_ = ((d_569_v14_)[default__.safeIndex(d_551_v0_, len(d_569_v14_))]) + (p3)
                    d_571_v15_: _dafny.Seq
                    d_571_v15_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_572_v16_: _dafny.Map
                    d_572_v16_ = _dafny.Map({not ((self).f15) or ((self).f15): d_571_v15_})
                    d_573_v17_: _dafny.Seq
                    d_573_v17_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                    d_574_v18_: _dafny.Map
                    d_574_v18_ = _dafny.Map({d_551_v0_: 859})
                    d_572_v16_ = (d_572_v16_).set((d_573_v17_)[default__.safeIndex(len(d_574_v18_), len(d_573_v17_))], (default__.fm34((0) - (d_551_v0_), (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))], (self).f15, globalState)).set(default__.safeIndex(d_551_v0_, len(default__.fm34((0) - (d_551_v0_), (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))], (self).f15, globalState))), ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sooucaqvp"))).set(default__.safeIndex(d_551_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sooucaqvp")))), p2)).set(default__.safeIndex(d_551_v0_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sooucaqvp"))).set(default__.safeIndex(d_551_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sooucaqvp")))), p2))), p2)))
                    pass
            pass
        if (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]:
            d_575_v19_: _dafny.Seq
            d_575_v19_ = _dafny.SeqWithoutIsStrInference([default__.fm0(globalState)])
            d_576_v20_: str
            d_576_v20_ = _dafny.CodePoint('s')
            d_577_v21_: _dafny.Seq
            d_577_v21_ = _dafny.SeqWithoutIsStrInference([True])
            rhs117_ = (_dafny.SeqWithoutIsStrInference([d_551_v0_])) + (p3)
            rhs118_ = p2
            rhs119_ = d_577_v21_
            d_575_v19_ = rhs117_
            d_576_v20_ = rhs118_
            d_577_v21_ = rhs119_
            if (self).f9:
                d_578_v22_: _dafny.Map
                d_578_v22_ = _dafny.Map({(0) - ((d_551_v0_) - (d_551_v0_)): (self).f15})
                d_578_v22_ = (d_578_v22_).set(d_551_v0_, (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))])
                d_555_v4_ = d_555_v4_
                d_579_v23_: _dafny.Array
                nw88_ = _dafny.Array(int(0), 11)
                d_579_v23_ = nw88_
                index82_ = default__.safeIndex(233, (d_579_v23_).length(0))
                (d_579_v23_)[index82_] = d_551_v0_
                index83_ = default__.safeIndex(233, (d_579_v23_).length(0))
                (d_579_v23_)[index83_] = d_551_v0_
                (globalState).f2 = (self).f9
                r2 = d_551_v0_
            elif True:
                d_580_v24_: _dafny.Map
                d_580_v24_ = _dafny.Map({default__.fm1(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_551_v0_ for d_581_i4_ in range(default__.abs(-632))])), 308, globalState): (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]})
                r2 = len(d_580_v24_)
                d_582_v25_: _dafny.Array
                nw89_ = _dafny.Array(None, 22)
                nw89_[int(0)] = p1
                nw89_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvqafp"))
                nw89_[int(2)] = (p1).set(default__.safeIndex(d_551_v0_, len(p1)), d_576_v20_)
                nw89_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                nw89_[int(4)] = p1
                nw89_[int(5)] = p1
                nw89_[int(6)] = p1
                nw89_[int(7)] = p1
                nw89_[int(8)] = p1
                nw89_[int(9)] = p1
                nw89_[int(10)] = p1
                nw89_[int(11)] = _dafny.SeqWithoutIsStrInference([p2 for d_583_i5_ in range(default__.abs(51))])
                nw89_[int(12)] = (self).fm33(globalState)
                nw89_[int(13)] = p1
                nw89_[int(14)] = (self).fm33(globalState)
                nw89_[int(15)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osyhtuwn"))) + (p1)
                nw89_[int(16)] = ((p1) + (p1)).set(default__.safeIndex(d_551_v0_, len((p1) + (p1))), p2)
                nw89_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqiqgyo"))
                nw89_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amy"))
                nw89_[int(19)] = (p1) + ((self).fm33(globalState))
                nw89_[int(20)] = p1
                nw89_[int(21)] = (p1) + (p1)
                d_582_v25_ = nw89_
                d_582_v25_ = d_582_v25_
                d_584_v26_: _dafny.Set
                d_584_v26_ = _dafny.Set({False, True})
                (globalState).f2 = (d_584_v26_).issubset((d_584_v26_).intersection(d_584_v26_))
                d_585_v27_: _dafny.Array
                nw90_ = _dafny.Array(int(0), 22)
                d_585_v27_ = nw90_
                d_585_v27_ = d_585_v27_
                index84_ = default__.safeIndex(394, (d_555_v4_).length(0))
                (d_555_v4_)[index84_] = False
            d_586_v28_: _dafny.Map
            d_586_v28_ = _dafny.Map({d_551_v0_: 986})
            d_587_v29_: _dafny.Map
            d_587_v29_ = _dafny.Map({(self).f15: d_551_v0_})
            d_588_v30_: _dafny.Array
            nw91_ = _dafny.Array(None, 9)
            nw91_[int(0)] = ((d_586_v28_)[d_551_v0_] if (d_551_v0_) in (d_586_v28_) else d_551_v0_)
            nw91_[int(1)] = (0) - ((0) - ((0) - ((len(_dafny.SeqWithoutIsStrInference([len(d_587_v29_)])) if (self).f15 else d_551_v0_))))
            nw91_[int(2)] = d_551_v0_
            nw91_[int(3)] = d_551_v0_
            nw91_[int(4)] = d_551_v0_
            nw91_[int(5)] = len(_dafny.SeqWithoutIsStrInference([(D1_DC3((self).f9, p1, 973)).cf7 for d_589_i6_ in range(default__.abs(899))]))
            nw91_[int(6)] = d_551_v0_
            nw91_[int(7)] = d_551_v0_
            nw91_[int(8)] = default__.fm0(globalState)
            d_588_v30_ = nw91_
            nw92_ = _dafny.Array(None, 1)
            nw92_[int(0)] = default__.safeDivisionInt(d_551_v0_, len(_dafny.Map({d_551_v0_: (self).f15})))
            d_588_v30_ = nw92_
            r0 = not(not ((d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]) or (not((self).f9)))
            (globalState).f2 = (((p1)[default__.safeIndex(d_551_v0_, len(p1))] if (self).f9 else p2)) not in ((p1) + ((self).fm33(globalState)))
        elif True:
            if not(False):
                r1 = ((d_551_v0_) * (d_551_v0_)) <= (d_551_v0_)
                d_590_v31_: _dafny.Map
                d_590_v31_ = _dafny.Map({d_555_v4_: d_551_v0_})
                d_591_v32_: _dafny.Array
                nw93_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_591_v32_ = nw93_
                d_592_v33_: _dafny.Array
                d_592_v33_ = d_591_v32_
                d_593_v34_: _dafny.Set
                d_593_v34_ = _dafny.Set({d_592_v33_})
                d_594_v35_: _dafny.Array
                nw94_ = _dafny.Array(None, 9)
                nw94_[int(0)] = d_551_v0_
                nw94_[int(1)] = len(d_593_v34_)
                nw94_[int(2)] = d_551_v0_
                nw94_[int(3)] = d_551_v0_
                nw94_[int(4)] = d_551_v0_
                nw94_[int(5)] = d_551_v0_
                nw94_[int(6)] = d_551_v0_
                nw94_[int(7)] = 323
                nw94_[int(8)] = d_551_v0_
                d_594_v35_ = nw94_
                d_595_v36_: C2
                nw95_ = C2()
                nw95_.ctor__(d_551_v0_, (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))])
                d_595_v36_ = nw95_
                d_596_v37_: _dafny.Seq
                d_596_v37_ = _dafny.SeqWithoutIsStrInference([d_595_v36_, d_595_v36_, d_595_v36_])
                d_597_v38_: _dafny.Map
                d_597_v38_ = _dafny.Map({d_594_v35_: len(d_596_v37_)})
                d_598_v39_: D11
                d_598_v39_ = D11_DC29(d_597_v38_)
                rhs120_ = ((_dafny.Map({d_555_v4_: d_551_v0_})) | (d_590_v31_)) | ((_dafny.Map({d_555_v4_: (d_595_v36_).f14}) if (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))] else d_590_v31_))
                rhs121_ = (d_598_v39_).cf51
                rhs122_ = ((d_595_v36_).f14 if ((d_595_v36_).f14) >= ((d_595_v36_).f14) else (d_595_v36_).f14)
                d_590_v31_ = rhs120_
                d_597_v38_ = rhs121_
                d_551_v0_ = rhs122_
                d_599_v40_: _dafny.Map
                d_599_v40_ = _dafny.Map({(0) - (d_551_v0_): d_551_v0_})
                d_600_v41_: _dafny.Seq
                d_600_v41_ = _dafny.SeqWithoutIsStrInference([d_599_v40_, d_599_v40_, d_599_v40_])
                d_601_v42_: _dafny.Map
                d_601_v42_ = _dafny.Map({(0) - (len(d_600_v41_)): (self).f15})
                (globalState).f2 = (((d_601_v42_)[22] if (22) in (d_601_v42_) else (self).f9) if ((d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]) == ((self).f9) else not((self).f15))
                d_602_v43_: _dafny.Array
                nw96_ = _dafny.Array(None, 13)
                nw96_[int(0)] = p1
                nw96_[int(1)] = p1
                nw96_[int(2)] = p1
                nw96_[int(3)] = p1
                nw96_[int(4)] = p1
                nw96_[int(5)] = p1
                nw96_[int(6)] = (((p1).set(default__.safeIndex((d_595_v36_).f14, len(p1)), p2)) + (p1)).set(default__.safeIndex(d_551_v0_, len(((p1).set(default__.safeIndex((d_595_v36_).f14, len(p1)), p2)) + (p1))), p2)
                nw96_[int(7)] = p1
                nw96_[int(8)] = p1
                nw96_[int(9)] = p1
                nw96_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofiikthw"))
                nw96_[int(11)] = p1
                nw96_[int(12)] = (p1).set(default__.safeIndex(d_551_v0_, len(p1)), p2)
                d_602_v43_ = nw96_
                d_602_v43_ = d_602_v43_
                d_552_v1_ = d_552_v1_
            elif True:
                index85_ = default__.safeIndex(394, (d_555_v4_).length(0))
                (d_555_v4_)[index85_] = (d_552_v1_).issubset(d_552_v1_)
                index86_ = default__.safeIndex(394, (d_555_v4_).length(0))
                (d_555_v4_)[index86_] = (self).f15
                d_603_v44_: _dafny.Map
                d_603_v44_ = _dafny.Map({(d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]: d_555_v4_})
                r2 = default__.safeDivisionInt(len((d_603_v44_) | (d_603_v44_)), (d_551_v0_) * ((0) - (d_551_v0_)))
                d_604_v45_: _dafny.Seq
                d_604_v45_ = _dafny.SeqWithoutIsStrInference([(p1) + (p1), p1, p1, p1, p1])
                r2 = len((d_604_v45_)[default__.safeIndex(((d_552_v1_)[len(p3)] if (len(p3)) in (d_552_v1_) else 275), len(d_604_v45_))])
                d_605_v46_: D1
                d_605_v46_ = D1_DC2(d_551_v0_, (self).f9, len(p1), d_551_v0_)
                d_606_v47_: _dafny.Seq
                d_606_v47_ = _dafny.SeqWithoutIsStrInference([d_605_v46_, d_605_v46_, d_605_v46_, d_605_v46_])
                d_607_v48_: D12
                d_607_v48_ = D12_DC31(d_606_v47_)
                d_608_v49_: C1
                nw97_ = C1()
                nw97_.ctor__(((d_607_v48_).cf52).set(default__.safeIndex(d_551_v0_, len((d_607_v48_).cf52)), d_605_v46_))
                d_608_v49_ = nw97_
            d_609_v50_: str
            d_609_v50_ = _dafny.CodePoint('n')
            rhs123_ = d_551_v0_
            rhs124_ = (self).f9
            rhs125_ = (self).f9
            rhs126_ = (p1)[default__.safeIndex(d_551_v0_, len(p1))]
            d_551_v0_ = rhs123_
            r0 = rhs124_
            r1 = rhs125_
            d_609_v50_ = rhs126_
            r2 = (d_551_v0_) * (d_551_v0_)
            default__.m0((d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))], globalState)
            r0 = ((self).f15 if ((self).f15 if (self).f15 else (self).f9) else (self).f15)
        d_610_v51_: _dafny.Seq
        d_610_v51_ = _dafny.SeqWithoutIsStrInference([(self).f9, (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]])
        d_611_v52_: _dafny.Set
        d_611_v52_ = _dafny.Set({(d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))], (d_610_v51_)[default__.safeIndex(d_551_v0_, len(d_610_v51_))], (self).f15, (True if (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))] else (self).f15), (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]})
        d_611_v52_ = (d_611_v52_) | (d_611_v52_)
        r0 = (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]
        r1 = (d_555_v4_)[default__.safeIndex(394, (d_555_v4_).length(0))]
        r2 = default__.safeDivisionInt(d_551_v0_, 186)
        r3 = (d_558_v6_) - (d_558_v6_)
        return r0, r1, r2, r3

    def m3(self, p0, globalState):
        r0: bool = False
        d_612_v0_: C2
        nw98_ = C2()
        nw98_.ctor__(660, (self).f9)
        d_612_v0_ = nw98_
        (globalState).f2 = (self).f15
        d_613_v1_: _dafny.Set
        d_613_v1_ = _dafny.Set({-456, (d_612_v0_).f14})
        d_614_v2_: _dafny.Seq
        d_614_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amua"))
        if (self).fm32(_dafny.SeqWithoutIsStrInference([d_613_v1_]), True, (d_612_v0_).f14, (len(d_614_v2_)) <= ((d_612_v0_).f14), globalState):
            d_615_v3_: int
            d_615_v3_ = -903
            d_615_v3_ = (d_612_v0_).f14
            d_616_v4_: str
            d_616_v4_ = _dafny.CodePoint('d')
            d_617_v5_: _dafny.Map
            d_617_v5_ = _dafny.Map({369: d_616_v4_})
            d_618_v6_: _dafny.Map
            d_618_v6_ = _dafny.Map({(self).f9: ((d_617_v5_)[d_615_v3_] if (d_615_v3_) in (d_617_v5_) else default__.fm35(globalState))})
            d_619_v7_: _dafny.Seq
            d_619_v7_ = _dafny.SeqWithoutIsStrInference([d_618_v6_])
            d_619_v7_ = _dafny.SeqWithoutIsStrInference([d_618_v6_ for d_620_i0_ in range(default__.abs(369))])
            d_621_v8_: _dafny.Array
            nw99_ = _dafny.Array(int(0), 2)
            d_621_v8_ = nw99_
            d_621_v8_ = d_621_v8_
            d_622_v9_: _dafny.Array
            nw100_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
            d_622_v9_ = nw100_
            index87_ = default__.safeIndex(496, (d_622_v9_).length(0))
            (d_622_v9_)[index87_] = d_614_v2_
            index88_ = default__.safeIndex(496, (d_622_v9_).length(0))
            (d_622_v9_)[index88_] = d_614_v2_
            d_623_v10_: _dafny.Map
            d_623_v10_ = _dafny.Map({(self).f9: (d_612_v0_).f14})
            d_624_v11_: _dafny.MultiSet
            d_624_v11_ = _dafny.MultiSet([992])
            d_625_v12_: _dafny.Map
            d_625_v12_ = _dafny.Map({d_623_v10_: (D9_DC25(d_615_v3_, (d_612_v0_).f14, (d_624_v11_).set((d_624_v11_).cardinality, default__.abs(d_615_v3_)), d_617_v5_, (d_612_v0_).f14)).cf45})
            d_615_v3_ = (len((default__.fm36((self).f15, (self).f15, p0, globalState)) | (d_625_v12_))) - (default__.safeDivisionInt((0) - (d_615_v3_), d_615_v3_))
        elif True:
            (globalState).f2 = ((self).f15) or (((d_612_v0_).f14) < ((d_612_v0_).f14))
            d_626_v13_: C2
            nw101_ = C2()
            nw101_.ctor__(((0) - ((d_612_v0_).f14)) * ((d_612_v0_).f14), not(((0) - ((d_612_v0_).f14)) == ((d_612_v0_).f14)))
            d_626_v13_ = nw101_
            d_627_v14_: int
            d_627_v14_ = 678
            d_627_v14_ = (d_612_v0_).f14
            d_628_v15_: _dafny.Map
            d_628_v15_ = _dafny.Map({(self).f15: (self).f9})
            d_629_v16_: _dafny.Map
            d_629_v16_ = _dafny.Map({(d_612_v0_).f14: not ((self).f15) or (((d_628_v15_)[not(True)] if (not(True)) in (d_628_v15_) else (self).f15))})
            d_630_v17_: _dafny.Map
            d_630_v17_ = _dafny.Map({(d_612_v0_).f14: len(d_629_v16_)})
            d_631_v18_: _dafny.Map
            d_631_v18_ = _dafny.Map({p0: d_630_v17_})
            if ((d_629_v16_)[d_627_v14_] if (d_627_v14_) in (d_629_v16_) else ((self).f9) not in (d_631_v18_)):
                (globalState).f2 = False
                (globalState).f2 = ((self).f9) == (not ((self).f9) or (p0))
                d_632_v19_: _dafny.MultiSet
                d_632_v19_ = _dafny.MultiSet([(self).f9])
                d_633_v20_: _dafny.Array
                nw102_ = _dafny.Array(None, 13)
                nw102_[int(0)] = d_627_v14_
                nw102_[int(1)] = -248
                nw102_[int(2)] = -829
                nw102_[int(3)] = len(_dafny.SeqWithoutIsStrInference([(d_632_v19_).cardinality]))
                nw102_[int(4)] = d_627_v14_
                nw102_[int(5)] = (d_626_v13_).f14
                nw102_[int(6)] = d_627_v14_
                nw102_[int(7)] = (d_626_v13_).f14
                nw102_[int(8)] = (d_626_v13_).f14
                nw102_[int(9)] = (d_612_v0_).f14
                nw102_[int(10)] = (d_612_v0_).f14
                nw102_[int(11)] = d_627_v14_
                nw102_[int(12)] = (d_626_v13_).f14
                d_633_v20_ = nw102_
                d_634_v21_: _dafny.MultiSet
                d_634_v21_ = _dafny.MultiSet([d_633_v20_, d_633_v20_, d_633_v20_])
                (d_612_v0_).m8((self).f9, (d_634_v21_) | (d_634_v21_), globalState)
                d_635_v22_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_635_v22_ = nw103_
                d_636_v23_: _dafny.Array
                nw104_ = _dafny.Array(None, 14)
                nw104_[int(0)] = d_635_v22_
                nw104_[int(1)] = d_635_v22_
                nw104_[int(2)] = (d_635_v22_ if (self).f9 else d_635_v22_)
                nw104_[int(3)] = d_635_v22_
                nw104_[int(4)] = d_635_v22_
                nw104_[int(5)] = d_635_v22_
                nw104_[int(6)] = d_635_v22_
                nw104_[int(7)] = d_635_v22_
                nw104_[int(8)] = d_635_v22_
                nw104_[int(9)] = d_635_v22_
                nw104_[int(10)] = d_635_v22_
                nw104_[int(11)] = d_635_v22_
                nw104_[int(12)] = d_635_v22_
                nw104_[int(13)] = d_635_v22_
                d_636_v23_ = nw104_
                index89_ = default__.safeIndex(387, (d_636_v23_).length(0))
                (d_636_v23_)[index89_] = d_635_v22_
                d_637_v24_: D13
                d_637_v24_ = D13_DC34(d_635_v22_)
                d_638_v25_: _dafny.Seq
                d_638_v25_ = _dafny.SeqWithoutIsStrInference([d_635_v22_, d_635_v22_, (d_637_v24_).cf57, d_635_v22_, d_635_v22_])
                index90_ = default__.safeIndex(387, (d_636_v23_).length(0))
                (d_636_v23_)[index90_] = (d_638_v25_)[default__.safeIndex((d_612_v0_).f14, len(d_638_v25_))]
                d_627_v14_ = ((d_626_v13_).f14) - (d_627_v14_)
            elif True:
                d_639_v26_: str
                d_639_v26_ = _dafny.CodePoint('q')
                d_639_v26_ = d_639_v26_
                d_640_v27_: _dafny.Set
                d_640_v27_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_626_v13_).f14, (d_612_v0_).f14]), _dafny.SeqWithoutIsStrInference([d_627_v14_])})
                d_641_v28_: C0
                nw105_ = C0()
                nw105_.ctor__(d_640_v27_, (d_626_v13_).f14)
                d_641_v28_ = nw105_
                d_642_v29_: _dafny.MultiSet
                d_642_v29_ = _dafny.MultiSet([d_641_v28_])
                d_643_v30_: _dafny.Map
                d_643_v30_ = _dafny.Map({(self).f15: default__.fm37(p0, (self).f9, globalState)})
                (globalState).f2 = ((d_642_v29_).isdisjoint(_dafny.MultiSet([d_641_v28_]))) not in (d_643_v30_)
                d_644_v31_: _dafny.Array
                nw106_ = _dafny.Array(_dafny.Seq({}), 20)
                d_644_v31_ = nw106_
                d_645_v32_: _dafny.Seq
                d_645_v32_ = _dafny.SeqWithoutIsStrInference([d_641_v28_, d_641_v28_, d_641_v28_])
                index91_ = default__.safeIndex(878, (d_644_v31_).length(0))
                (d_644_v31_)[index91_] = d_645_v32_
                index92_ = default__.safeIndex(878, (d_644_v31_).length(0))
                (d_644_v31_)[index92_] = d_645_v32_
                d_646_v33_: _dafny.Array
                nw107_ = _dafny.Array(int(0), 19)
                d_646_v33_ = nw107_
                index93_ = default__.safeIndex(861, (d_646_v33_).length(0))
                (d_646_v33_)[index93_] = (d_612_v0_).f14
                d_647_v34_: _dafny.Seq
                d_647_v34_ = _dafny.SeqWithoutIsStrInference([(self).f15, p0])
                d_648_v35_: D1
                d_648_v35_ = D1_DC3((self).f15, d_614_v2_, (d_626_v13_).f14)
                d_649_v36_: D4
                d_649_v36_ = D4_DC14(not((self).f9), (self).f15, (d_612_v0_).f14)
                index94_ = default__.safeIndex(861, (d_646_v33_).length(0))
                rhs127_ = (len(d_647_v34_)) - ((d_626_v13_).f14)
                rhs128_ = default__.fm35(globalState)
                rhs129_ = default__.safeDivisionInt((d_612_v0_).f14, 928)
                rhs130_ = not((len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyablh"))) + ((d_648_v35_).cf7))) >= (((d_612_v0_).f14) + ((d_649_v36_).cf28)))
                rhs131_ = d_639_v26_
                lhs76_ = d_641_v28_
                lhs77_ = d_646_v33_
                lhs78_ = default__.safeIndex(861, (d_646_v33_).length(0))
                lhs79_ = globalState
                lhs76_.f12 = rhs127_
                d_639_v26_ = rhs128_
                lhs77_[lhs78_] = rhs129_
                lhs79_.f2 = rhs130_
                d_639_v26_ = rhs131_
                (globalState).f2 = not((self).f9)
            d_627_v14_ = 335
        r0 = not((p0) == ((self).f15))
        rhs132_ = p0
        rhs133_ = not((self).f9)
        rhs134_ = (self).f9
        rhs135_ = d_614_v2_
        lhs80_ = globalState
        lhs81_ = globalState
        lhs82_ = globalState
        lhs80_.f2 = rhs132_
        lhs81_.f2 = rhs133_
        lhs82_.f2 = rhs134_
        d_614_v2_ = rhs135_
        hi2_ = default__.safeDivisionInt(default__.fm0(globalState), (d_612_v0_).f14)
        for d_650_i1_ in range((d_612_v0_).f14, hi2_):
            d_651_v37_: _dafny.Map
            d_651_v37_ = _dafny.Map({(self).f9: not((self).f9)})
            d_652_v38_: _dafny.MultiSet
            d_652_v38_ = _dafny.MultiSet([(d_612_v0_).f14])
            d_653_v39_: _dafny.Map
            d_653_v39_ = _dafny.Map({(self).f15: (d_612_v0_).f14})
            d_651_v37_ = (d_651_v37_).set((d_652_v38_).isdisjoint(d_652_v38_), (d_653_v39_) == (d_653_v39_))
            d_654_v40_: int
            d_654_v40_ = -700
            d_654_v40_ = d_654_v40_
            d_655_v41_: _dafny.MultiSet
            d_655_v41_ = _dafny.MultiSet([p0, (self).f9, (self).f9])
            d_656_v42_: _dafny.Seq
            d_656_v42_ = _dafny.SeqWithoutIsStrInference([(0) - ((530) * (d_650_i1_)), default__.safeDivisionInt(len(d_651_v37_), (d_612_v0_).f14), (d_655_v41_).cardinality])
            d_654_v40_ = len(d_656_v42_)
            d_657_v43_: _dafny.Map
            d_657_v43_ = _dafny.Map({default__.safeModuloInt(378, d_650_i1_): not((self).f9)})
            rhs136_ = (self).f9
            rhs137_ = ((d_657_v43_ if True else d_657_v43_)).set(d_650_i1_, (d_654_v40_) > (d_650_i1_))
            lhs83_ = globalState
            lhs83_.f2 = rhs136_
            d_657_v43_ = rhs137_
        r0 = (self).f15
        return r0

    def m1(self, p0, p1, p2, p3, globalState):
        d_658_v0_: int
        d_658_v0_ = 605
        d_659_v1_: _dafny.Map
        d_659_v1_ = _dafny.Map({p2: d_658_v0_})
        d_660_v2_: _dafny.Seq
        d_660_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p2: p0}), d_659_v1_, d_659_v1_, d_659_v1_])
        d_661_v3_: D14
        d_661_v3_ = D14_DC37(d_660_v2_)
        d_662_v4_: _dafny.Set
        d_662_v4_ = _dafny.Set({d_658_v0_})
        pat_let_tv12_ = d_660_v2_
        pat_let_tv13_ = d_658_v0_
        pat_let_tv14_ = d_660_v2_
        pat_let_tv15_ = d_659_v1_
        pat_let_tv16_ = d_662_v4_
        pat_let_tv17_ = d_659_v1_
        def iife55_(_pat_let17_0):
            def iife56_(d_663_dt__update__tmp_h0_):
                def iife57_(_pat_let18_0):
                    def iife58_(d_664_dt__update_hcf66_h0_):
                        return D14_DC37(d_664_dt__update_hcf66_h0_)
                    return iife58_(_pat_let18_0)
                return iife57_(_dafny.SeqWithoutIsStrInference([(pat_let_tv12_)[default__.safeIndex(pat_let_tv13_, len(pat_let_tv14_))], (pat_let_tv15_).set((self).f9, len(pat_let_tv16_)), pat_let_tv17_]))
            return iife56_(_pat_let17_0)
        d_658_v0_ = len((d_660_v2_) + ((iife55_(d_661_v3_)).cf66))
        d_665_v5_: _dafny.Seq
        d_665_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thqlb"))
        d_666_v6_: _dafny.MultiSet
        d_666_v6_ = _dafny.MultiSet([d_658_v0_])
        d_667_v7_: D1
        d_667_v7_ = D1_DC2(937, p1, (d_666_v6_).cardinality, d_658_v0_)
        d_668_v8_: C1
        nw108_ = C1()
        nw108_.ctor__(_dafny.SeqWithoutIsStrInference([default__.fm38((self).f15, len(d_665_v5_), globalState), d_667_v7_, d_667_v7_]))
        d_668_v8_ = nw108_
        d_669_v9_: _dafny.Seq
        d_669_v9_ = _dafny.SeqWithoutIsStrInference([d_668_v8_])
        d_670_v10_: _dafny.Seq
        d_670_v10_ = _dafny.SeqWithoutIsStrInference([p1])
        d_671_v11_: _dafny.Seq
        d_671_v11_ = _dafny.SeqWithoutIsStrInference([p0])
        d_672_v12_: _dafny.Array
        nw109_ = _dafny.Array(None, 15)
        nw109_[int(0)] = p2
        nw109_[int(1)] = True
        nw109_[int(2)] = (self).f15
        nw109_[int(3)] = not((self).f15)
        nw109_[int(4)] = (self).f9
        nw109_[int(5)] = (self).f15
        nw109_[int(6)] = (self).f15
        nw109_[int(7)] = p2
        nw109_[int(8)] = (self).f9
        nw109_[int(9)] = (self).f15
        nw109_[int(10)] = (self).f9
        nw109_[int(11)] = (self).f9
        nw109_[int(12)] = (self).f15
        nw109_[int(13)] = (self).f9
        nw109_[int(14)] = p1
        d_672_v12_ = nw109_
        d_673_v13_: _dafny.MultiSet
        d_673_v13_ = _dafny.MultiSet([d_672_v12_, d_672_v12_])
        d_674_v14_: _dafny.Map
        d_674_v14_ = _dafny.Map({d_658_v0_: p1})
        d_675_v15_: _dafny.Array
        nw110_ = _dafny.Array(None, 25)
        nw110_[int(0)] = (d_658_v0_ if (self).f9 else d_658_v0_)
        nw110_[int(1)] = len(d_669_v9_)
        nw110_[int(2)] = len(d_670_v10_)
        nw110_[int(3)] = len(d_671_v11_)
        nw110_[int(4)] = d_658_v0_
        nw110_[int(5)] = (p0) + (d_658_v0_)
        nw110_[int(6)] = p0
        nw110_[int(7)] = p0
        nw110_[int(8)] = (d_673_v13_).cardinality
        nw110_[int(9)] = len((d_665_v5_) + (d_665_v5_))
        nw110_[int(10)] = d_658_v0_
        nw110_[int(11)] = d_658_v0_
        nw110_[int(12)] = d_658_v0_
        nw110_[int(13)] = p0
        nw110_[int(14)] = d_658_v0_
        nw110_[int(15)] = d_658_v0_
        nw110_[int(16)] = len(_dafny.SeqWithoutIsStrInference([len((d_668_v8_).fm2(d_658_v0_, True, p1, p0, globalState))]))
        nw110_[int(17)] = (0) - (d_658_v0_)
        nw110_[int(18)] = ((0) - (len(d_674_v14_))) * (-316)
        nw110_[int(19)] = p0
        nw110_[int(20)] = 44
        nw110_[int(21)] = 660
        nw110_[int(22)] = len(default__.fm39(globalState))
        nw110_[int(23)] = p0
        nw110_[int(24)] = d_658_v0_
        d_675_v15_ = nw110_
        index95_ = default__.safeIndex(897, (d_675_v15_).length(0))
        (d_675_v15_)[index95_] = p0
        d_676_v16_: _dafny.MultiSet
        d_676_v16_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afneutgkf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eaxwwu")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_677_i0_ in range(default__.abs(-209))])])
        d_678_v17_: _dafny.Array
        def lambda24_(d_679_p2_):
            def lambda25_(d_680_i1_):
                return _dafny.Map({d_679_p2_: True})

            return lambda25_

        init14_ = lambda24_(p2)
        nw111_ = _dafny.Array(None, 4)
        for i0_14_ in range(nw111_.length(0)):
            nw111_[i0_14_] = init14_(i0_14_)
        d_678_v17_ = nw111_
        d_681_v18_: D7
        d_681_v18_ = D7_DC19(d_678_v17_, p1, d_665_v5_)
        index96_ = default__.safeIndex(897, (d_675_v15_).length(0))
        rhs138_ = default__.safeDivisionInt((d_671_v11_)[default__.safeIndex(d_658_v0_, len(d_671_v11_))], p0)
        rhs139_ = (_dafny.MultiSet([d_665_v5_, (d_681_v18_).cf35, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_682_i2_ in range(default__.abs(651))]), d_665_v5_])) | (default__.fm40(globalState))
        rhs140_ = (self).f15
        rhs141_ = (d_658_v0_) - ((d_671_v11_)[default__.safeIndex(p0, len(d_671_v11_))])
        lhs84_ = d_675_v15_
        lhs85_ = default__.safeIndex(897, (d_675_v15_).length(0))
        lhs86_ = globalState
        lhs84_[lhs85_] = rhs138_
        d_676_v16_ = rhs139_
        lhs86_.f2 = rhs140_
        d_658_v0_ = rhs141_
        d_683_v19_: _dafny.Array
        def lambda26_(d_684_i3_):
            return _dafny.CodePoint('d')

        init15_ = lambda26_
        nw112_ = _dafny.Array(None, 26)
        for i0_15_ in range(nw112_.length(0)):
            nw112_[i0_15_] = init15_(i0_15_)
        d_683_v19_ = nw112_
        d_685_v20_: str
        d_685_v20_ = _dafny.CodePoint('m')
        index97_ = default__.safeIndex(966, (d_683_v19_).length(0))
        (d_683_v19_)[index97_] = d_685_v20_
        index98_ = default__.safeIndex(966, (d_683_v19_).length(0))
        (d_683_v19_)[index98_] = d_685_v20_
        def lambda27_(d_686_v15_):
            def lambda28_(d_687_i4_):
                return default__.safeModuloInt(d_687_i4_, len(_dafny.Map({(d_686_v15_)[default__.safeIndex(897, (d_686_v15_).length(0))]: (d_686_v15_)[default__.safeIndex(897, (d_686_v15_).length(0))]})))

            return lambda28_

        init16_ = lambda27_(d_675_v15_)
        nw113_ = _dafny.Array(None, 6)
        for i0_16_ in range(nw113_.length(0)):
            nw113_[i0_16_] = init16_(i0_16_)
        d_675_v15_ = nw113_
        d_688_v21_: D10
        d_688_v21_ = D10_DC27(_dafny.Map({p1: len((self).fm2((0) - (len(d_665_v5_)), (self).f9, (self).f9, d_658_v0_, globalState))}))
        d_689_v22_: D14
        d_689_v22_ = D14_DC38((d_683_v19_)[default__.safeIndex(966, (d_683_v19_).length(0))], (self).f15, default__.fm41(d_688_v21_, d_658_v0_, d_671_v11_, globalState))
        pat_let_tv18_ = p1
        pat_let_tv19_ = d_670_v10_
        d_690_v23_: _dafny.Set
        def iife59_(_pat_let19_0):
            def iife60_(d_691_dt__update__tmp_h1_):
                def iife61_(_pat_let20_0):
                    def iife62_(d_692_dt__update_hcf68_h0_):
                        def iife63_(_pat_let21_0):
                            def iife64_(d_693_dt__update_hcf69_h0_):
                                return D14_DC38((d_691_dt__update__tmp_h1_).cf67, d_692_dt__update_hcf68_h0_, d_693_dt__update_hcf69_h0_)
                            return iife64_(_pat_let21_0)
                        return iife63_(pat_let_tv19_)
                    return iife62_(_pat_let20_0)
                return iife61_(pat_let_tv18_)
            return iife60_(_pat_let19_0)
        d_690_v23_ = _dafny.Set({d_689_v22_, d_689_v22_, d_689_v22_, d_689_v22_, iife59_(d_689_v22_)})
        d_690_v23_ = (d_690_v23_).intersection((d_690_v23_) - (d_690_v23_))
        d_694_v24_: _dafny.Seq
        d_694_v24_ = _dafny.SeqWithoutIsStrInference([d_666_v6_, _dafny.MultiSet(d_671_v11_), d_666_v6_])
        d_695_v25_: _dafny.MultiSet
        d_695_v25_ = (d_694_v24_)[default__.safeIndex((d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))], len(d_694_v24_))]
        def lambda29_(source7_):
            d_696___mcc_h0_ = source7_
            d_697_cf31_ = d_696___mcc_h0_
            return False

        if lambda29_(d_695_v25_):
            (globalState).f2 = ((d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))]) >= (d_658_v0_)
            if (_dafny.SeqWithoutIsStrInference([(d_689_v22_).cf67 for d_698_i5_ in range(default__.abs(-978))])) <= (d_665_v5_):
                d_699_v26_: _dafny.Map
                d_699_v26_ = _dafny.Map({_dafny.CodePoint('k'): p1})
                d_700_v27_: _dafny.Set
                d_700_v27_ = _dafny.Set({not(p1)})
                d_701_v28_: D15
                d_701_v28_ = D15_DC40(_dafny.Set({(self).f9}))
                d_702_v29_: _dafny.Seq
                d_702_v29_ = _dafny.SeqWithoutIsStrInference([d_700_v27_, d_700_v27_, (d_701_v28_).cf72, d_700_v27_])
                index99_ = default__.safeIndex(897, (d_675_v15_).length(0))
                (d_675_v15_)[index99_] = default__.safeDivisionInt(len(d_699_v26_), len((d_702_v29_)[default__.safeIndex(d_658_v0_, len(d_702_v29_))]))
                d_703_v30_: _dafny.Map
                d_703_v30_ = _dafny.Map({p1: d_685_v20_})
                (globalState).f2 = not((p2) or ((((d_703_v30_)[p1] if (p1) in (d_703_v30_) else (d_683_v19_)[default__.safeIndex(966, (d_683_v19_).length(0))])) in ((self).fm33(globalState))))
                d_658_v0_ = (d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))]
                d_704_v31_: _dafny.Map
                d_704_v31_ = _dafny.Map({d_665_v5_: len(_dafny.SeqWithoutIsStrInference([p0 for d_705_i6_ in range(default__.abs(511))]))})
                index100_ = default__.safeIndex(897, (d_675_v15_).length(0))
                (d_675_v15_)[index100_] = (0) - (((d_704_v31_)[(d_665_v5_) + ((self).fm33(globalState))] if ((d_665_v5_) + ((self).fm33(globalState))) in (d_704_v31_) else (default__.fm42(p2, p1, len(d_670_v10_), globalState)).cardinality))
                nw114_ = _dafny.Array(int(0), 27)
                d_675_v15_ = nw114_
            elif True:
                d_706_v32_: _dafny.Map
                d_706_v32_ = _dafny.Map({not(p1): p1})
                index101_ = default__.safeIndex(897, (d_675_v15_).length(0))
                (d_675_v15_)[index101_] = len(d_706_v32_)
                index102_ = default__.safeIndex(897, (d_675_v15_).length(0))
                (d_675_v15_)[index102_] = p0
                d_707_v33_: _dafny.Set
                d_707_v33_ = _dafny.Set({(d_658_v0_) > ((d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))])})
                index103_ = default__.safeIndex(897, (d_675_v15_).length(0))
                rhs142_ = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([p1, p1])), d_658_v0_)
                rhs143_ = d_707_v33_
                rhs144_ = d_675_v15_
                rhs145_ = (self).f9
                rhs146_ = (p0) <= (d_658_v0_)
                lhs87_ = d_675_v15_
                lhs88_ = default__.safeIndex(897, (d_675_v15_).length(0))
                lhs89_ = globalState
                lhs90_ = globalState
                lhs87_[lhs88_] = rhs142_
                d_707_v33_ = rhs143_
                d_675_v15_ = rhs144_
                lhs89_.f2 = rhs145_
                lhs90_.f2 = rhs146_
                d_708_v34_: _dafny.Array
                nw115_ = _dafny.Array(D3.default()(), 29)
                d_708_v34_ = nw115_
                d_709_v35_: D3
                d_709_v35_ = D3_DC8()
                index104_ = default__.safeIndex(147, (d_708_v34_).length(0))
                (d_708_v34_)[index104_] = d_709_v35_
                index105_ = default__.safeIndex(897, (d_675_v15_).length(0))
                index106_ = default__.safeIndex(147, (d_708_v34_).length(0))
                rhs147_ = (default__.safeModuloInt((d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))], p0)) >= (p0)
                rhs148_ = d_706_v32_
                rhs149_ = (p0 if p1 else (d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))])
                rhs150_ = (d_709_v35_ if p1 else d_709_v35_)
                lhs91_ = globalState
                lhs92_ = d_675_v15_
                lhs93_ = default__.safeIndex(897, (d_675_v15_).length(0))
                lhs94_ = d_708_v34_
                lhs95_ = default__.safeIndex(147, (d_708_v34_).length(0))
                lhs91_.f2 = rhs147_
                d_706_v32_ = rhs148_
                lhs92_[lhs93_] = rhs149_
                lhs94_[lhs95_] = rhs150_
                index107_ = default__.safeIndex(897, (d_675_v15_).length(0))
                (d_675_v15_)[index107_] = (0) - (default__.fm0(globalState))
            d_710_v36_: D4
            d_710_v36_ = D4_DC13(d_658_v0_)
            d_658_v0_ = (0) - ((d_710_v36_).cf25)
            d_711_v37_: _dafny.Array
            nw116_ = _dafny.Array(None, 8)
            nw116_[int(0)] = d_665_v5_
            nw116_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmco"))
            nw116_[int(2)] = (self).fm33(globalState)
            nw116_[int(3)] = d_665_v5_
            nw116_[int(4)] = d_665_v5_
            nw116_[int(5)] = d_665_v5_
            nw116_[int(6)] = d_665_v5_
            nw116_[int(7)] = d_665_v5_
            d_711_v37_ = nw116_
            d_712_v38_: _dafny.Seq
            d_712_v38_ = _dafny.SeqWithoutIsStrInference([d_711_v37_, d_711_v37_, d_711_v37_])
            d_713_v39_: _dafny.Map
            d_713_v39_ = _dafny.Map({_dafny.Map({d_668_v8_: 590}): (d_712_v38_)[default__.safeIndex((d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))], len(d_712_v38_))]})
            d_714_v40_: _dafny.Map
            d_714_v40_ = _dafny.Map({d_668_v8_: p0})
            d_713_v39_ = (d_713_v39_).set(d_714_v40_, d_711_v37_)
            index108_ = default__.safeIndex(897, (d_675_v15_).length(0))
            (d_675_v15_)[index108_] = d_658_v0_
        elif True:
            index109_ = default__.safeIndex(897, (d_675_v15_).length(0))
            (d_675_v15_)[index109_] = p0
            d_715_v41_: _dafny.Map
            d_715_v41_ = _dafny.Map({(d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))]: (d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))]})
            d_716_v42_: _dafny.Map
            d_716_v42_ = _dafny.Map({((D16_DC45(d_715_v41_)).cf78) | (d_715_v41_): ((self).f9) or ((self).f15)})
            (globalState).f2 = ((d_716_v42_)[d_715_v41_] if (d_715_v41_) in (d_716_v42_) else p1)
            d_658_v0_ = (d_675_v15_)[default__.safeIndex(897, (d_675_v15_).length(0))]
            d_658_v0_ = default__.fm0(globalState)
            d_671_v11_ = d_671_v11_

    @property
    def f15(self):
        return self._f15

class C4(T1, T0):
    def  __init__(self):
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f9(self):
        return self._f9
    def ctor__(self, f9):
        (self)._f9 = f9

    def fm2(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({222})), (_dafny.MultiSet([not((self).f9)])).cardinality])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('v')])).cardinality]))) + ((_dafny.SeqWithoutIsStrInference([725 for d_717_i0_ in range(default__.abs(413))])) + (_dafny.SeqWithoutIsStrInference([381])))

    def m1(self, p0, p1, p2, p3, globalState):
        d_718_v0_: _dafny.MultiSet
        d_718_v0_ = _dafny.MultiSet([p0])
        d_719_i0_: int
        d_719_i0_ = 0
        with _dafny.label("7"):
            while (((d_718_v0_).intersection(d_718_v0_)).cardinality) <= (p0):
                with _dafny.c_label("7"):
                    if (d_719_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_719_i0_ = (d_719_i0_) + (1)
                    d_720_v1_: _dafny.Array
                    def lambda30_(d_721_i1_):
                        return _dafny.CodePoint('x')

                    init17_ = lambda30_
                    nw117_ = _dafny.Array(None, 14)
                    for i0_17_ in range(nw117_.length(0)):
                        nw117_[i0_17_] = init17_(i0_17_)
                    d_720_v1_ = nw117_
                    d_720_v1_ = d_720_v1_
                    d_722_v2_: _dafny.Array
                    def lambda31_(d_723_p2_):
                        def lambda32_(d_724_i2_):
                            return d_723_p2_

                        return lambda32_

                    init18_ = lambda31_(p2)
                    nw118_ = _dafny.Array(None, 26)
                    for i0_18_ in range(nw118_.length(0)):
                        nw118_[i0_18_] = init18_(i0_18_)
                    d_722_v2_ = nw118_
                    d_725_v3_: _dafny.Map
                    d_725_v3_ = _dafny.Map({p0: p1})
                    d_726_v4_: _dafny.Seq
                    d_726_v4_ = _dafny.SeqWithoutIsStrInference([len(d_725_v3_)])
                    d_727_v5_: _dafny.Seq
                    d_727_v5_ = _dafny.SeqWithoutIsStrInference([(p0) + (p0), p0, (p0) - (len(_dafny.Map({p0: d_722_v2_}))), (d_726_v4_)[default__.safeIndex(p0, len(d_726_v4_))], len(_dafny.Map({p2: (self).f9}))])
                    d_726_v4_ = (d_726_v4_) + (d_727_v5_)
                    d_728_v6_: int
                    d_728_v6_ = 442
                    d_728_v6_ = (len(d_726_v4_)) * (p0)
                    d_729_v7_: _dafny.Array
                    def lambda33_(d_730_i3_):
                        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwxxvgrr"))

                    init19_ = lambda33_
                    nw119_ = _dafny.Array(None, 16)
                    for i0_19_ in range(nw119_.length(0)):
                        nw119_[i0_19_] = init19_(i0_19_)
                    d_729_v7_ = nw119_
                    d_731_v8_: D1
                    d_731_v8_ = D1_DC3(p2, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_732_i4_ in range(default__.abs(388))]), d_728_v6_)
                    index110_ = default__.safeIndex(653, (d_729_v7_).length(0))
                    (d_729_v7_)[index110_] = (d_731_v8_).cf7
                    index111_ = default__.safeIndex(653, (d_729_v7_).length(0))
                    (d_729_v7_)[index111_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_733_i5_ in range(default__.abs(997))])
                    pass
            pass
        (globalState).f2 = not(not(True))
        d_734_v9_: _dafny.Seq
        d_734_v9_ = _dafny.SeqWithoutIsStrInference([p0])
        d_735_v10_: _dafny.MultiSet
        d_735_v10_ = _dafny.MultiSet([d_734_v9_])
        d_736_v11_: D1
        d_736_v11_ = D1_DC2(p0, (self).f9, p0, (d_735_v10_).cardinality)
        d_737_v12_: _dafny.Seq
        d_737_v12_ = _dafny.SeqWithoutIsStrInference([d_736_v11_])
        d_738_v13_: C1
        nw120_ = C1()
        nw120_.ctor__((d_737_v12_) + (d_737_v12_))
        d_738_v13_ = nw120_
        d_739_v14_: _dafny.Map
        d_739_v14_ = _dafny.Map({True: p0})
        d_740_v15_: str
        d_740_v15_ = _dafny.CodePoint('o')
        d_741_v16_: _dafny.Seq
        d_741_v16_ = _dafny.SeqWithoutIsStrInference([p2, False])
        d_742_v17_: _dafny.Map
        d_742_v17_ = _dafny.Map({_dafny.CodePoint('d'): d_741_v16_})
        d_743_v18_: _dafny.Array
        nw121_ = _dafny.Array(None, 21)
        nw121_[int(0)] = p2
        nw121_[int(1)] = (self).f9
        nw121_[int(2)] = p1
        nw121_[int(3)] = (p2 if p2 else p2)
        nw121_[int(4)] = p2
        nw121_[int(5)] = p1
        nw121_[int(6)] = not((self).f9)
        nw121_[int(7)] = (p0) != (868)
        nw121_[int(8)] = (p0) not in ((d_718_v0_).set(p0, default__.abs((0) - (len(d_739_v14_)))))
        nw121_[int(9)] = (d_740_v15_) in (_dafny.Map({d_740_v15_: _dafny.MultiSet(((d_742_v17_)[d_740_v15_] if (d_740_v15_) in (d_742_v17_) else d_741_v16_))}))
        nw121_[int(10)] = (p1) or (p1)
        nw121_[int(11)] = False
        nw121_[int(12)] = (p0) >= (p0)
        nw121_[int(13)] = (True) == (not(p1))
        nw121_[int(14)] = default__.fm1(d_718_v0_, p0, globalState)
        nw121_[int(15)] = (self).f9
        nw121_[int(16)] = (self).f9
        nw121_[int(17)] = (self).f9
        nw121_[int(18)] = p2
        nw121_[int(19)] = (d_741_v16_)[default__.safeIndex(p0, len(d_741_v16_))]
        nw121_[int(20)] = p1
        d_743_v18_ = nw121_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_743_v18_).length(0)):
            d_744_i6_: int = guard_loop_3_
            if (True) and (((0) <= (d_744_i6_)) and ((d_744_i6_) < ((d_743_v18_).length(0)))):
                (d_743_v18_)[(d_744_i6_)] = False
        pat_let_tv20_ = p0
        pat_let_tv21_ = p0
        pat_let_tv22_ = p0
        pat_let_tv23_ = p0
        pat_let_tv24_ = p0
        def lambda34_(source9_):
            if source9_.is_DC8:
                return D1_DC1(pat_let_tv20_)
            elif source9_.is_DC9:
                d_745___mcc_h14_ = source9_.cf17
                d_746___mcc_h15_ = source9_.cf18
                d_747_cf18_ = d_746___mcc_h15_
                d_748_cf17_ = d_745___mcc_h14_
                if d_748_cf17_:
                    return D1_DC1((0) - (pat_let_tv21_))
                elif True:
                    return D1_DC1(pat_let_tv22_)
            elif source9_.is_DC10:
                return D1_DC1(pat_let_tv23_)
            elif True:
                d_749___mcc_h16_ = source9_.cf16
                d_750_cf16_ = d_749___mcc_h16_
                return D1_DC1(pat_let_tv24_)

        source8_ = lambda34_(D3_DC8())
        if source8_.is_DC2:
            d_751___mcc_h0_ = source8_.cf2
            d_752___mcc_h1_ = source8_.cf3
            d_753___mcc_h2_ = source8_.cf4
            d_754___mcc_h3_ = source8_.cf5
            d_755_cf5_ = d_754___mcc_h3_
            d_756_cf4_ = d_753___mcc_h2_
            d_757_cf3_ = d_752___mcc_h1_
            d_758_cf2_ = d_751___mcc_h0_
            d_759_v19_: _dafny.Map
            d_759_v19_ = _dafny.Map({p0: False})
            d_760_v20_: D1
            d_760_v20_ = D1_DC4(d_755_cf5_, p2, ((d_759_v19_)[d_756_cf4_] if (d_756_cf4_) in (d_759_v19_) else p2), True, d_757_cf3_)
            d_761_v21_: D1
            d_761_v21_ = D1_DC5(d_760_v20_)
            source10_ = d_761_v21_
            if source10_.is_DC2:
                d_762___mcc_h17_ = source10_.cf2
                d_763___mcc_h18_ = source10_.cf3
                d_764___mcc_h19_ = source10_.cf4
                d_765___mcc_h20_ = source10_.cf5
                d_766_cf5_ = d_765___mcc_h20_
                d_767_cf4_ = d_764___mcc_h19_
                d_768_cf3_ = d_763___mcc_h18_
                d_769_cf2_ = d_762___mcc_h17_
                d_770_v22_: _dafny.Array
                def lambda35_(d_771_cf4_):
                    def lambda36_(d_772_i7_):
                        return (d_772_i7_) - ((0) - (d_771_cf4_))

                    return lambda36_

                init20_ = lambda35_(d_767_cf4_)
                nw122_ = _dafny.Array(None, 18)
                for i0_20_ in range(nw122_.length(0)):
                    nw122_[i0_20_] = init20_(i0_20_)
                d_770_v22_ = nw122_
                index112_ = default__.safeIndex(731, (d_770_v22_).length(0))
                (d_770_v22_)[index112_] = len(_dafny.SeqWithoutIsStrInference([d_766_cf5_]))
                index113_ = default__.safeIndex(731, (d_770_v22_).length(0))
                (d_770_v22_)[index113_] = d_769_cf2_
                d_773_v23_: T0
                nw123_ = C1()
                nw123_.ctor__(d_738_v13_.f13)
                d_773_v23_ = nw123_
                d_774_v24_: bool
                d_775_v25_: bool
                out10_: bool
                out11_: bool
                out10_, out11_ = (d_738_v13_).m6(d_773_v23_, d_743_v18_, True, globalState)
                d_774_v24_ = out10_
                d_775_v25_ = out11_
                d_776_v26_: _dafny.Array
                nw124_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                d_776_v26_ = nw124_
                d_777_v27_: _dafny.Array
                def lambda37_(d_778_v0_):
                    def lambda38_(d_779_i8_):
                        return _dafny.Map({d_778_v0_: True})

                    return lambda38_

                init21_ = lambda37_(d_718_v0_)
                nw125_ = _dafny.Array(None, 5)
                for i0_21_ in range(nw125_.length(0)):
                    nw125_[i0_21_] = init21_(i0_21_)
                d_777_v27_ = nw125_
                d_780_v28_: _dafny.Map
                d_780_v28_ = _dafny.Map({d_718_v0_: d_775_v25_})
                index114_ = default__.safeIndex(146, (d_777_v27_).length(0))
                (d_777_v27_)[index114_] = d_780_v28_
                d_781_v29_: _dafny.Seq
                d_781_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_774_v24_: p0})) for d_782_i9_ in range(default__.abs(499))])])
                d_783_v31_: D4
                d_783_v31_ = D4_DC12(d_774_v24_, d_775_v25_, p0, (d_770_v22_)[default__.safeIndex(731, (d_770_v22_).length(0))], (self).f9)
                d_784_v32_: D7
                def iife65_():
                    coll21_ = _dafny.Map()
                    compr_21_: _dafny.MultiSet
                    for compr_21_ in (_dafny.SeqWithoutIsStrInference([d_718_v0_, d_718_v0_, _dafny.MultiSet([(0) - ((d_783_v31_).cf22), d_767_cf4_])])).Elements:
                        d_785_v30_: _dafny.MultiSet = compr_21_
                        if (d_785_v30_) in (_dafny.SeqWithoutIsStrInference([d_718_v0_, d_718_v0_, _dafny.MultiSet([(0) - ((d_783_v31_).cf22), d_767_cf4_])])):
                            coll21_[d_785_v30_] = True
                    return _dafny.Map(coll21_)
                d_784_v32_ = D7_DC18(iife65_()
)
                pat_let_tv25_ = d_760_v20_
                index115_ = default__.safeIndex(146, (d_777_v27_).length(0))
                def iife66_(_pat_let22_0):
                    def iife67_(d_786_dt__update__tmp_h0_):
                        def iife68_(_pat_let23_0):
                            def iife69_(d_787_dt__update_hcf14_h0_):
                                return D1_DC5(d_787_dt__update_hcf14_h0_)
                            return iife69_(_pat_let23_0)
                        return iife68_(pat_let_tv25_)
                    return iife67_(_pat_let22_0)
                rhs151_ = iife66_(D1_DC5(d_760_v20_))
                rhs152_ = d_776_v26_
                rhs153_ = ((d_781_v29_) != (d_781_v29_) if not((self).f9) else not(d_774_v24_))
                rhs154_ = (d_784_v32_).cf32
                lhs96_ = d_777_v27_
                lhs97_ = default__.safeIndex(146, (d_777_v27_).length(0))
                d_761_v21_ = rhs151_
                d_776_v26_ = rhs152_
                d_768_cf3_ = rhs153_
                lhs96_[lhs97_] = rhs154_
                d_766_cf5_ = (d_769_cf2_) * ((d_718_v0_).cardinality)
            elif source10_.is_DC3:
                d_788___mcc_h21_ = source10_.cf6
                d_789___mcc_h22_ = source10_.cf7
                d_790___mcc_h23_ = source10_.cf8
                d_791_cf8_ = d_790___mcc_h23_
                d_792_cf7_ = d_789___mcc_h22_
                d_793_cf6_ = d_788___mcc_h21_
                default__.m0(p1, globalState)
                d_740_v15_ = d_740_v15_
                d_794_v33_: _dafny.Set
                d_794_v33_ = _dafny.Set({d_734_v9_, _dafny.SeqWithoutIsStrInference([561 for d_795_i10_ in range(default__.abs(881))])})
                d_796_v34_: C0
                nw126_ = C0()
                nw126_.ctor__(d_794_v33_, len(_dafny.Map({p1: d_757_cf3_})))
                d_796_v34_ = nw126_
                d_797_v35_: _dafny.Map
                d_797_v35_ = _dafny.Map({d_793_cf6_: d_796_v34_})
                d_798_v36_: _dafny.Seq
                d_798_v36_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p2: d_796_v34_}), (d_797_v35_) | (d_797_v35_), (d_797_v35_) | (_dafny.Map({d_793_cf6_: d_796_v34_}))])
                d_799_v37_: _dafny.MultiSet
                d_799_v37_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_740_v15_ for d_800_i11_ in range(default__.abs(720))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unmdd"))])
                d_801_v38_: _dafny.Seq
                d_801_v38_ = _dafny.SeqWithoutIsStrInference([d_799_v37_])
                d_802_v39_: _dafny.Seq
                d_802_v39_ = _dafny.SeqWithoutIsStrInference([d_792_cf7_, d_792_cf7_])
                d_803_v40_: _dafny.Array
                nw127_ = _dafny.Array(None, 21)
                nw127_[int(0)] = (d_799_v37_) | (d_799_v37_)
                nw127_[int(1)] = (d_799_v37_) | ((d_801_v38_)[default__.safeIndex(d_758_cf2_, len(d_801_v38_))])
                nw127_[int(2)] = d_799_v37_
                nw127_[int(3)] = d_799_v37_
                nw127_[int(4)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_792_cf7_ for d_804_i12_ in range(default__.abs(954))]))) | (d_799_v37_)
                nw127_[int(5)] = (_dafny.MultiSet(d_802_v39_)) - (d_799_v37_)
                nw127_[int(6)] = default__.fm19(p0, globalState)
                nw127_[int(7)] = d_799_v37_
                nw127_[int(8)] = (_dafny.MultiSet(d_802_v39_)).intersection(d_799_v37_)
                nw127_[int(9)] = d_799_v37_
                nw127_[int(10)] = d_799_v37_
                nw127_[int(11)] = d_799_v37_
                nw127_[int(12)] = d_799_v37_
                nw127_[int(13)] = default__.fm19(d_756_cf4_, globalState)
                nw127_[int(14)] = d_799_v37_
                nw127_[int(15)] = d_799_v37_
                nw127_[int(16)] = (d_799_v37_).set(d_792_cf7_, default__.abs(d_796_v34_.f12))
                nw127_[int(17)] = d_799_v37_
                nw127_[int(18)] = d_799_v37_
                nw127_[int(19)] = d_799_v37_
                nw127_[int(20)] = d_799_v37_
                d_803_v40_ = nw127_
                index116_ = default__.safeIndex(442, (d_803_v40_).length(0))
                (d_803_v40_)[index116_] = d_799_v37_
                d_805_v41_: _dafny.Map
                d_805_v41_ = _dafny.Map({d_796_v34_.f12: d_796_v34_})
                d_806_v42_: _dafny.Array
                def lambda39_(d_807_p1_, d_808_cf3_):
                    def lambda40_(d_809_i13_):
                        return _dafny.Map({d_807_p1_: d_808_cf3_})

                    return lambda40_

                init22_ = lambda39_(p1, d_757_cf3_)
                nw128_ = _dafny.Array(None, 27)
                for i0_22_ in range(nw128_.length(0)):
                    nw128_[i0_22_] = init22_(i0_22_)
                d_806_v42_ = nw128_
                d_810_v43_: D7
                d_810_v43_ = D7_DC19(d_806_v42_, p1, d_792_cf7_)
                index117_ = default__.safeIndex(442, (d_803_v40_).length(0))
                rhs155_ = -627
                rhs156_ = ((d_798_v36_) + (d_798_v36_)).set(default__.safeIndex(d_755_cf5_, len((d_798_v36_) + (d_798_v36_))), d_797_v35_)
                rhs157_ = d_799_v37_
                rhs158_ = d_805_v41_
                rhs159_ = len(((d_810_v43_).cf35) + (d_792_cf7_))
                lhs98_ = d_803_v40_
                lhs99_ = default__.safeIndex(442, (d_803_v40_).length(0))
                d_756_cf4_ = rhs155_
                d_798_v36_ = rhs156_
                lhs98_[lhs99_] = rhs157_
                d_805_v41_ = rhs158_
                d_756_cf4_ = rhs159_
                d_811_v44_: D1
                d_811_v44_ = D1_DC4(d_756_cf4_, True, not(not(d_757_cf3_)), (self).f9, False)
                d_812_v45_: C0
                nw129_ = C0()
                nw129_.ctor__(d_794_v33_, (d_811_v44_).cf9)
                d_812_v45_ = nw129_
            elif source10_.is_DC4:
                d_813___mcc_h24_ = source10_.cf9
                d_814___mcc_h25_ = source10_.cf10
                d_815___mcc_h26_ = source10_.cf11
                d_816___mcc_h27_ = source10_.cf12
                d_817___mcc_h28_ = source10_.cf13
                d_818_cf13_ = d_817___mcc_h28_
                d_819_cf12_ = d_816___mcc_h27_
                d_820_cf11_ = d_815___mcc_h26_
                d_821_cf10_ = d_814___mcc_h25_
                d_822_cf9_ = d_813___mcc_h24_
                d_823_v46_: _dafny.Map
                d_823_v46_ = _dafny.Map({(d_718_v0_) - (d_718_v0_): not((p1) and (p2))})
                d_823_v46_ = (d_823_v46_).set((d_718_v0_).intersection(_dafny.MultiSet([d_755_cf5_])), not((d_741_v16_)[default__.safeIndex(d_822_cf9_, len(d_741_v16_))]))
                index118_ = default__.safeIndex(191, (d_743_v18_).length(0))
                (d_743_v18_)[index118_] = (d_757_cf3_) == (d_757_cf3_)
                index119_ = default__.safeIndex(191, (d_743_v18_).length(0))
                (d_743_v18_)[index119_] = d_819_cf12_
                index120_ = default__.safeIndex(191, (d_743_v18_).length(0))
                (d_743_v18_)[index120_] = (d_743_v18_)[default__.safeIndex(191, (d_743_v18_).length(0))]
                (globalState).f2 = p2
            elif source10_.is_DC1:
                d_824___mcc_h29_ = source10_.cf1
                d_825_cf1_ = d_824___mcc_h29_
                d_826_v47_: _dafny.Set
                d_826_v47_ = _dafny.Set({d_734_v9_})
                d_827_v48_: C0
                nw130_ = C0()
                nw130_.ctor__((d_826_v47_).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([(0) - (d_755_cf5_) for d_828_i14_ in range(default__.abs(534))])})), 197)
                d_827_v48_ = nw130_
                d_827_v48_ = d_827_v48_
                d_829_v49_: D1
                d_829_v49_ = D1_DC3((d_757_cf3_) == (p2), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srewo")), d_756_cf4_)
                d_829_v49_ = D1_DC3(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrcaagrw")), d_825_cf1_)
                d_830_v50_: _dafny.Array
                def lambda41_(d_831_p0_):
                    def lambda42_(d_832_i15_):
                        return (d_832_i15_) * (d_831_p0_)

                    return lambda42_

                init23_ = lambda41_(p0)
                nw131_ = _dafny.Array(None, 5)
                for i0_23_ in range(nw131_.length(0)):
                    nw131_[i0_23_] = init23_(i0_23_)
                d_830_v50_ = nw131_
                d_830_v50_ = d_830_v50_
            elif True:
                d_833___mcc_h30_ = source10_.cf14
                d_834_cf14_ = d_833___mcc_h30_
                d_835_v51_: T0
                nw132_ = C1()
                nw132_.ctor__(d_737_v12_)
                d_835_v51_ = nw132_
                d_836_v52_: bool
                d_837_v53_: bool
                out12_: bool
                out13_: bool
                out12_, out13_ = (d_738_v13_).m6(d_835_v51_, d_743_v18_, p1, globalState)
                d_836_v52_ = out12_
                d_837_v53_ = out13_
                d_838_v54_: _dafny.Seq
                d_838_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wslaqc"))
                d_838_v54_ = (d_838_v54_).set(default__.safeIndex(d_758_cf2_, len(d_838_v54_)), d_740_v15_)
                d_740_v15_ = _dafny.CodePoint('r')
                d_838_v54_ = d_838_v54_
            rhs160_ = not(True)
            rhs161_ = (0) - ((default__.safeModuloInt((0) - (d_756_cf4_), len(_dafny.Map({len(d_734_v9_): p0})))) * (d_755_cf5_))
            lhs100_ = globalState
            lhs100_.f2 = rhs160_
            d_756_cf4_ = rhs161_
            d_757_cf3_ = p1
            d_756_cf4_ = (p0) + (d_758_cf2_)
        elif source8_.is_DC3:
            d_839___mcc_h4_ = source8_.cf6
            d_840___mcc_h5_ = source8_.cf7
            d_841___mcc_h6_ = source8_.cf8
            d_842_cf8_ = d_841___mcc_h6_
            d_843_cf7_ = d_840___mcc_h5_
            d_844_cf6_ = d_839___mcc_h4_
            d_845_v55_: _dafny.Array
            nw133_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
            d_845_v55_ = nw133_
            index121_ = default__.safeIndex(87, (d_845_v55_).length(0))
            (d_845_v55_)[index121_] = d_843_cf7_
            d_846_v56_: _dafny.Array
            def lambda43_(d_847_p0_):
                def lambda44_(d_848_i16_):
                    return default__.safeDivisionInt(d_848_i16_, d_847_p0_)

                return lambda44_

            init24_ = lambda43_(p0)
            nw134_ = _dafny.Array(None, 29)
            for i0_24_ in range(nw134_.length(0)):
                nw134_[i0_24_] = init24_(i0_24_)
            d_846_v56_ = nw134_
            d_849_v57_: _dafny.MultiSet
            d_849_v57_ = _dafny.MultiSet([p2])
            d_850_v58_: _dafny.Map
            d_850_v58_ = _dafny.Map({(self).f9: _dafny.SeqWithoutIsStrInference([p0, (d_849_v57_).cardinality])})
            index122_ = default__.safeIndex(520, (d_846_v56_).length(0))
            (d_846_v56_)[index122_] = default__.safeModuloInt(len(d_739_v14_), len(d_850_v58_))
            index123_ = default__.safeIndex(87, (d_845_v55_).length(0))
            index124_ = default__.safeIndex(520, (d_846_v56_).length(0))
            rhs162_ = (_dafny.SeqWithoutIsStrInference([d_740_v15_ for d_851_i17_ in range(default__.abs(491))])) + (d_843_cf7_)
            rhs163_ = p0
            rhs164_ = default__.fm20(not(False), p1, default__.fm0(globalState), _dafny.SeqWithoutIsStrInference([d_740_v15_ for d_852_i18_ in range(default__.abs(51))]), globalState)
            rhs165_ = 812
            lhs101_ = d_845_v55_
            lhs102_ = default__.safeIndex(87, (d_845_v55_).length(0))
            lhs103_ = d_846_v56_
            lhs104_ = default__.safeIndex(520, (d_846_v56_).length(0))
            lhs101_[lhs102_] = rhs162_
            lhs103_[lhs104_] = rhs163_
            d_741_v16_ = rhs164_
            d_842_cf8_ = rhs165_
            d_853_v59_: _dafny.Set
            d_853_v59_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_844_cf6_]), d_741_v16_})
            if (d_741_v16_) not in (d_853_v59_):
                index125_ = default__.safeIndex(520, (d_846_v56_).length(0))
                rhs166_ = d_842_cf8_
                rhs167_ = d_846_v56_
                lhs105_ = d_846_v56_
                lhs106_ = default__.safeIndex(520, (d_846_v56_).length(0))
                lhs105_[lhs106_] = rhs166_
                d_846_v56_ = rhs167_
                index126_ = default__.safeIndex(561, (d_743_v18_).length(0))
                (d_743_v18_)[index126_] = False
                index127_ = default__.safeIndex(561, (d_743_v18_).length(0))
                (d_743_v18_)[index127_] = p2
                d_842_cf8_ = (d_846_v56_)[default__.safeIndex(520, (d_846_v56_).length(0))]
                d_854_v60_: D1
                d_854_v60_ = D1_DC4(p0, False, p1, not(d_844_cf6_), (d_743_v18_)[default__.safeIndex(561, (d_743_v18_).length(0))])
                d_854_v60_ = d_854_v60_
                d_844_cf6_ = (((d_739_v14_)[(d_743_v18_)[default__.safeIndex(561, (d_743_v18_).length(0))]] if ((d_743_v18_)[default__.safeIndex(561, (d_743_v18_).length(0))]) in (d_739_v14_) else (d_846_v56_)[default__.safeIndex(520, (d_846_v56_).length(0))])) != (default__.safeDivisionInt((d_846_v56_)[default__.safeIndex(520, (d_846_v56_).length(0))], d_842_cf8_))
            elif True:
                index128_ = default__.safeIndex(520, (d_846_v56_).length(0))
                (d_846_v56_)[index128_] = (default__.safeModuloInt(p0, d_842_cf8_)) * ((d_846_v56_)[default__.safeIndex(520, (d_846_v56_).length(0))])
                default__.m0(d_844_cf6_, globalState)
                d_855_v61_: _dafny.Set
                d_855_v61_ = _dafny.Set({d_734_v9_})
                d_856_v62_: C0
                nw135_ = C0()
                nw135_.ctor__((d_855_v61_ if d_844_cf6_ else _dafny.Set({(self).fm2(d_842_cf8_, (self).f9, d_844_cf6_, d_842_cf8_, globalState), d_734_v9_, d_734_v9_})), ((d_739_v14_)[p1] if (p1) in (d_739_v14_) else len((d_845_v55_)[default__.safeIndex(87, (d_845_v55_).length(0))])))
                d_856_v62_ = nw135_
                d_857_v63_: _dafny.Map
                d_857_v63_ = _dafny.Map({d_844_cf6_: d_846_v56_})
                d_858_v64_: _dafny.Array
                d_858_v64_ = ((d_857_v63_)[False] if (False) in (d_857_v63_) else d_846_v56_)
                rhs168_ = (d_858_v64_)
                rhs169_ = (self).f9
                lhs107_ = globalState
                d_846_v56_ = rhs168_
                lhs107_.f2 = rhs169_
                d_859_v65_: _dafny.Map
                d_859_v65_ = _dafny.Map({d_842_cf8_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfasfg"))})
                d_859_v65_ = (d_859_v65_).set(568, (d_843_cf7_) + (d_843_cf7_))
            d_842_cf8_ = (d_846_v56_)[default__.safeIndex(520, (d_846_v56_).length(0))]
            index129_ = default__.safeIndex(520, (d_846_v56_).length(0))
            (d_846_v56_)[index129_] = ((d_849_v57_) | ((d_849_v57_) - (default__.fm21(len(d_734_v9_), d_844_cf6_, globalState)))).cardinality
        elif source8_.is_DC4:
            d_860___mcc_h7_ = source8_.cf9
            d_861___mcc_h8_ = source8_.cf10
            d_862___mcc_h9_ = source8_.cf11
            d_863___mcc_h10_ = source8_.cf12
            d_864___mcc_h11_ = source8_.cf13
            d_865_cf13_ = d_864___mcc_h11_
            d_866_cf12_ = d_863___mcc_h10_
            d_867_cf11_ = d_862___mcc_h9_
            d_868_cf10_ = d_861___mcc_h8_
            d_869_cf9_ = d_860___mcc_h7_
            d_870_v66_: _dafny.Map
            d_870_v66_ = _dafny.Map({d_743_v18_: d_869_cf9_})
            if (len(d_870_v66_)) != (len((d_734_v9_) + (d_734_v9_))):
                d_871_v67_: _dafny.Array
                def lambda45_(d_872_v9_):
                    def lambda46_(d_873_i19_):
                        return d_872_v9_

                    return lambda46_

                init25_ = lambda45_(d_734_v9_)
                nw136_ = _dafny.Array(None, 19)
                for i0_25_ in range(nw136_.length(0)):
                    nw136_[i0_25_] = init25_(i0_25_)
                d_871_v67_ = nw136_
                d_874_v68_: _dafny.Seq
                d_874_v68_ = d_734_v9_
                index130_ = default__.safeIndex(3, (d_871_v67_).length(0))
                (d_871_v67_)[index130_] = d_874_v68_
                index131_ = default__.safeIndex(3, (d_871_v67_).length(0))
                (d_871_v67_)[index131_] = d_874_v68_
                d_739_v14_ = (d_739_v14_).set(d_867_cf11_, p0)
                d_875_v69_: _dafny.Set
                d_875_v69_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_869_cf9_])})
                d_876_v70_: C0
                nw137_ = C0()
                nw137_.ctor__(d_875_v69_, 663)
                d_876_v70_ = nw137_
                default__.m0(p1, globalState)
                (d_876_v70_).f12 = p0
            elif True:
                d_877_v71_: _dafny.Seq
                d_877_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
                d_877_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hick"))
                d_878_v72_: _dafny.Array
                nw138_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_878_v72_ = nw138_
                d_879_v73_: _dafny.Array
                nw139_ = _dafny.Array(_dafny.Seq({}), 9)
                d_879_v73_ = nw139_
                index132_ = default__.safeIndex(698, (d_878_v72_).length(0))
                (d_878_v72_)[index132_] = d_879_v73_
                d_880_v74_: _dafny.Seq
                d_880_v74_ = _dafny.SeqWithoutIsStrInference([d_879_v73_])
                index133_ = default__.safeIndex(698, (d_878_v72_).length(0))
                rhs170_ = (d_879_v73_ if default__.fm1(d_718_v0_, d_869_cf9_, globalState) else (d_880_v74_)[default__.safeIndex(522, len(d_880_v74_))])
                rhs171_ = d_743_v18_
                rhs172_ = (self).f9
                lhs108_ = d_878_v72_
                lhs109_ = default__.safeIndex(698, (d_878_v72_).length(0))
                lhs108_[lhs109_] = rhs170_
                d_743_v18_ = rhs171_
                d_865_cf13_ = rhs172_
                d_881_v75_: D7
                d_881_v75_ = D7_DC18((_dafny.Map({d_718_v0_: p1})) | (default__.fm22(globalState)))
                d_882_v76_: _dafny.Map
                d_882_v76_ = _dafny.Map({d_718_v0_: d_865_cf13_})
                d_881_v75_ = D7_DC18(d_882_v76_)
                d_883_v77_: _dafny.Map
                d_883_v77_ = _dafny.Map({p0: d_869_cf9_})
                d_884_v78_: _dafny.MultiSet
                d_884_v78_ = _dafny.MultiSet([d_883_v77_])
                d_885_v79_: _dafny.Map
                d_885_v79_ = _dafny.Map({d_877_v71_: d_869_cf9_})
                d_884_v78_ = ((d_884_v78_) - (d_884_v78_)).set(d_883_v77_, default__.abs(((d_734_v9_)[default__.safeIndex(d_869_cf9_, len(d_734_v9_))]) * (len(d_885_v79_))))
                d_868_cf10_ = (d_868_cf10_) == (p1)
            d_886_v80_: _dafny.Set
            d_886_v80_ = _dafny.Set({d_867_cf11_, (self).f9, d_865_cf13_, p2, d_868_cf10_})
            d_887_v81_: _dafny.Map
            d_887_v81_ = _dafny.Map({d_886_v80_: d_868_cf10_})
            d_869_cf9_ = (p0) - ((0) - (len(d_887_v81_)))
            d_888_v82_: _dafny.Seq
            d_888_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
            d_889_v83_: _dafny.Seq
            d_889_v83_ = _dafny.SeqWithoutIsStrInference([d_888_v82_, d_888_v82_])
            d_869_cf9_ = len((_dafny.SeqWithoutIsStrInference([d_888_v82_])) + (d_889_v83_))
            d_890_v84_: _dafny.Array
            def lambda47_(d_891_cf9_):
                def lambda48_(d_892_i20_):
                    return (_dafny.SeqWithoutIsStrInference([d_891_cf9_ for d_893_i21_ in range(default__.abs(-154))]))

                return lambda48_

            init26_ = lambda47_(d_869_cf9_)
            nw140_ = _dafny.Array(None, 27)
            for i0_26_ in range(nw140_.length(0)):
                nw140_[i0_26_] = init26_(i0_26_)
            d_890_v84_ = nw140_
            d_890_v84_ = d_890_v84_
        elif source8_.is_DC1:
            d_894___mcc_h12_ = source8_.cf1
            d_895_cf1_ = d_894___mcc_h12_
            d_896_v85_: _dafny.Array
            nw141_ = _dafny.Array(int(0), 22)
            d_896_v85_ = nw141_
            d_897_v86_: _dafny.Map
            d_897_v86_ = _dafny.Map({p2: not(p2)})
            index134_ = default__.safeIndex(431, (d_896_v85_).length(0))
            (d_896_v85_)[index134_] = len(d_897_v86_)
            index135_ = default__.safeIndex(431, (d_896_v85_).length(0))
            (d_896_v85_)[index135_] = d_895_cf1_
            index136_ = default__.safeIndex(431, (d_896_v85_).length(0))
            (d_896_v85_)[index136_] = d_895_cf1_
            rhs173_ = (self).f9
            rhs174_ = not(not(default__.fm1(d_718_v0_, default__.safeDivisionInt(-397, (d_896_v85_)[default__.safeIndex(431, (d_896_v85_).length(0))]), globalState)))
            lhs110_ = globalState
            lhs111_ = globalState
            lhs110_.f2 = rhs173_
            lhs111_.f2 = rhs174_
            (globalState).f2 = p2
        elif True:
            d_898___mcc_h13_ = source8_.cf14
            d_899_cf14_ = d_898___mcc_h13_
            d_900_v87_: _dafny.Array
            nw142_ = _dafny.Array(int(0), 8)
            d_900_v87_ = nw142_
            index137_ = default__.safeIndex(175, (d_900_v87_).length(0))
            (d_900_v87_)[index137_] = p0
            index138_ = default__.safeIndex(175, (d_900_v87_).length(0))
            (d_900_v87_)[index138_] = p0
            d_901_v88_: _dafny.Seq
            d_901_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
            d_902_v90_: _dafny.Set
            d_902_v90_ = _dafny.Set({-638})
            index139_ = default__.safeIndex(175, (d_900_v87_).length(0))
            def iife70_():
                coll22_ = _dafny.Set()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(319, 877):
                    d_903_v89_: int = compr_22_
                    if ((319) <= (d_903_v89_)) and ((d_903_v89_) < (877)):
                        coll22_ = coll22_.union(_dafny.Set([default__.safeModuloInt(d_903_v89_, (d_900_v87_)[default__.safeIndex(175, (d_900_v87_).length(0))])]))
                return _dafny.Set(coll22_)
            rhs175_ = (0) - (len(default__.fm23(not((iife70_()
            ).issubset(d_902_v90_)), (d_901_v88_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnvrovbj"))), p0, globalState)))
            rhs176_ = default__.fm1(_dafny.MultiSet([(d_900_v87_)[default__.safeIndex(175, (d_900_v87_).length(0))], 844]), default__.safeDivisionInt(p0, p0), globalState)
            rhs177_ = ((self).f9) and ((self).f9)
            rhs178_ = d_901_v88_
            rhs179_ = d_901_v88_
            lhs112_ = d_900_v87_
            lhs113_ = default__.safeIndex(175, (d_900_v87_).length(0))
            lhs114_ = globalState
            lhs115_ = globalState
            lhs112_[lhs113_] = rhs175_
            lhs114_.f2 = rhs176_
            lhs115_.f2 = rhs177_
            d_901_v88_ = rhs178_
            d_901_v88_ = rhs179_
            d_904_v91_: _dafny.Array
            nw143_ = _dafny.Array(None, 12)
            d_904_v91_ = nw143_
            d_905_v92_: _dafny.Map
            d_905_v92_ = _dafny.Map({d_734_v9_: d_718_v0_})
            d_906_v93_: T1
            nw144_ = C2()
            nw144_.ctor__((((d_905_v92_)[d_734_v9_] if (d_734_v9_) in (d_905_v92_) else d_718_v0_)).cardinality, p1)
            d_906_v93_ = nw144_
            index140_ = default__.safeIndex(142, (d_904_v91_).length(0))
            (d_904_v91_)[index140_] = d_906_v93_
            index141_ = default__.safeIndex(142, (d_904_v91_).length(0))
            (d_904_v91_)[index141_] = d_906_v93_
            index142_ = default__.safeIndex(175, (d_900_v87_).length(0))
            (d_900_v87_)[index142_] = (d_900_v87_)[default__.safeIndex(175, (d_900_v87_).length(0))]
        d_734_v9_ = d_734_v9_

    def m7(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: D4 = D4.default()()
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        r3: bool = False
        if not ((self).f9) or (not((self).f9)):
            d_907_v0_: _dafny.MultiSet
            d_907_v0_ = _dafny.MultiSet([230, len(_dafny.Map({(_dafny.MultiSet([(self).f9])).cardinality: p0}))])
            d_908_v1_: _dafny.Seq
            d_908_v1_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_909_v2_: _dafny.Map
            d_909_v2_ = _dafny.Map({(self).f9: (self).f9})
            d_910_v3_: _dafny.MultiSet
            d_910_v3_ = _dafny.MultiSet([(self).f9])
            d_911_v4_: _dafny.Map
            d_911_v4_ = _dafny.Map({(self).f9: p0})
            d_912_v5_: _dafny.Map
            d_912_v5_ = _dafny.Map({p1: p0})
            d_913_v6_: _dafny.Set
            d_913_v6_ = _dafny.Set({d_912_v5_})
            d_914_v7_: _dafny.Array
            nw145_ = _dafny.Array(None, 28)
            nw145_[int(0)] = (261) * (p0)
            nw145_[int(1)] = default__.safeDivisionInt(295, p1)
            nw145_[int(2)] = p1
            nw145_[int(3)] = p1
            nw145_[int(4)] = p0
            nw145_[int(5)] = (0) - (default__.safeDivisionInt(((d_907_v0_)[p0] if (p0) in (d_907_v0_) else p0), default__.fm0(globalState)))
            nw145_[int(6)] = p1
            nw145_[int(7)] = p0
            nw145_[int(8)] = default__.safeModuloInt(len((d_908_v1_).set(default__.safeIndex(p0, len(d_908_v1_)), p2)), (0) - (p0))
            nw145_[int(9)] = default__.safeModuloInt(p0, p0)
            nw145_[int(10)] = p0
            nw145_[int(11)] = p0
            nw145_[int(12)] = p1
            nw145_[int(13)] = p0
            nw145_[int(14)] = 898
            nw145_[int(15)] = p1
            nw145_[int(16)] = p1
            nw145_[int(17)] = default__.safeModuloInt(len(d_909_v2_), (d_910_v3_).cardinality)
            nw145_[int(18)] = p0
            nw145_[int(19)] = (0) - (p1)
            nw145_[int(20)] = p1
            nw145_[int(21)] = len((d_911_v4_) | (default__.fm23((self).f9, p2, p0, globalState)))
            nw145_[int(22)] = 899
            nw145_[int(23)] = (p1) - (944)
            nw145_[int(24)] = (229 if (self).f9 else p0)
            nw145_[int(25)] = p0
            nw145_[int(26)] = p1
            nw145_[int(27)] = (p1) - (len(d_913_v6_))
            d_914_v7_ = nw145_
            index143_ = default__.safeIndex(262, (d_914_v7_).length(0))
            (d_914_v7_)[index143_] = p1
            d_915_v8_: _dafny.Map
            d_915_v8_ = _dafny.Map({_dafny.MultiSet([p0, p1, p0]): p2})
            d_916_v9_: D7
            d_916_v9_ = D7_DC18((d_915_v8_).set(d_907_v0_, p2))
            d_917_v10_: _dafny.Array
            def lambda49_(d_918_i0_):
                return D7_DC20((self).f9, _dafny.CodePoint('q'))

            init27_ = lambda49_
            nw146_ = _dafny.Array(None, 9)
            for i0_27_ in range(nw146_.length(0)):
                nw146_[i0_27_] = init27_(i0_27_)
            d_917_v10_ = nw146_
            d_919_v11_: _dafny.Seq
            d_919_v11_ = _dafny.SeqWithoutIsStrInference([(d_912_v5_) | (d_912_v5_)])
            d_920_v12_: _dafny.Seq
            d_920_v12_ = _dafny.SeqWithoutIsStrInference([p1])
            d_921_v13_: _dafny.Seq
            d_921_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcofbj"))
            index144_ = default__.safeIndex(262, (d_914_v7_).length(0))
            rhs180_ = len(d_919_v11_)
            rhs181_ = (default__.fm20((p1) != (len(d_908_v1_)), default__.fm1(default__.fm29(p2, (self).f9, globalState), (d_920_v12_)[default__.safeIndex(133, len(d_920_v12_))], globalState), (p0) + (len(d_908_v1_)), d_921_v13_, globalState)).set(default__.safeIndex((0) - (p1), len(default__.fm20((p1) != (len(d_908_v1_)), default__.fm1(default__.fm29(p2, (self).f9, globalState), (d_920_v12_)[default__.safeIndex(133, len(d_920_v12_))], globalState), (p0) + (len(d_908_v1_)), d_921_v13_, globalState))), p2)
            rhs182_ = D7_DC18(d_915_v8_)
            rhs183_ = p0
            rhs184_ = d_917_v10_
            lhs116_ = d_914_v7_
            lhs117_ = default__.safeIndex(262, (d_914_v7_).length(0))
            lhs116_[lhs117_] = rhs180_
            d_908_v1_ = rhs181_
            d_916_v9_ = rhs182_
            r0 = rhs183_
            d_917_v10_ = rhs184_
            d_922_v14_: _dafny.Seq
            d_922_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('c')])])
            d_923_v15_: D1
            d_923_v15_ = D1_DC2(len(d_921_v13_), (self).f9, 503, p0)
            d_924_v16_: _dafny.Seq
            def iife71_(_pat_let24_0):
                def iife72_(d_925_dt__update__tmp_h0_):
                    def iife73_(_pat_let25_0):
                        def iife74_(d_926_dt__update_hcf3_h0_):
                            return D1_DC2((d_925_dt__update__tmp_h0_).cf2, d_926_dt__update_hcf3_h0_, (d_925_dt__update__tmp_h0_).cf4, (d_925_dt__update__tmp_h0_).cf5)
                        return iife74_(_pat_let25_0)
                    return iife73_((self).f9)
                return iife72_(_pat_let24_0)
            d_924_v16_ = _dafny.SeqWithoutIsStrInference([iife71_(d_923_v15_), d_923_v15_])
            d_927_v17_: str
            d_927_v17_ = _dafny.CodePoint('e')
            pat_let_tv26_ = d_914_v7_
            pat_let_tv27_ = d_914_v7_
            pat_let_tv28_ = p0
            pat_let_tv29_ = d_914_v7_
            pat_let_tv30_ = d_914_v7_
            pat_let_tv31_ = p0
            pat_let_tv32_ = p1
            pat_let_tv33_ = d_914_v7_
            pat_let_tv34_ = d_914_v7_
            pat_let_tv35_ = p0
            pat_let_tv36_ = p0
            pat_let_tv37_ = d_914_v7_
            pat_let_tv38_ = d_914_v7_
            pat_let_tv39_ = p0
            pat_let_tv40_ = p1
            pat_let_tv41_ = d_914_v7_
            pat_let_tv42_ = d_914_v7_
            pat_let_tv43_ = p1
            pat_let_tv44_ = d_927_v17_
            pat_let_tv45_ = d_907_v0_
            pat_let_tv46_ = globalState
            d_928_v18_: C1
            nw147_ = C1()
            def iife75_(_pat_let26_0):
                def iife76_(d_929_dt__update__tmp_h1_):
                    def iife77_(_pat_let27_0):
                        def iife78_(d_930_dt__update_hcf5_h0_):
                            def iife79_(_pat_let28_0):
                                def iife80_(d_933_dt__update_hcf4_h0_):
                                    def iife81_(_pat_let29_0):
                                        def iife82_(d_935_dt__update_hcf3_h1_):
                                            return D1_DC2((d_929_dt__update__tmp_h1_).cf2, d_935_dt__update_hcf3_h1_, d_933_dt__update_hcf4_h0_, d_930_dt__update_hcf5_h0_)
                                        return iife82_(_pat_let29_0)
                                    return iife81_(default__.fm1(pat_let_tv45_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_934_i1_ in range(default__.abs(-302))])), pat_let_tv46_))
                                return iife80_(_pat_let28_0)
                            return iife79_(len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_931_i2_ in range(default__.abs(949))])).set(default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([pat_let_tv28_, (pat_let_tv30_)[default__.safeIndex(262, (pat_let_tv29_).length(0))], pat_let_tv31_, pat_let_tv32_, (pat_let_tv34_)[default__.safeIndex(262, (pat_let_tv33_).length(0))]])).set(default__.safeIndex(pat_let_tv35_, len(_dafny.SeqWithoutIsStrInference([pat_let_tv36_, (pat_let_tv38_)[default__.safeIndex(262, (pat_let_tv37_).length(0))], pat_let_tv39_, pat_let_tv40_, (pat_let_tv42_)[default__.safeIndex(262, (pat_let_tv41_).length(0))]]))), pat_let_tv43_)), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_932_i2_ in range(default__.abs(949))]))), pat_let_tv44_)))
                        return iife78_(_pat_let27_0)
                    return iife77_((pat_let_tv27_)[default__.safeIndex(262, (pat_let_tv26_).length(0))])
                return iife76_(_pat_let26_0)
            nw147_.ctor__((d_924_v16_ if default__.fm1(_dafny.MultiSet([637, ((d_922_v14_)[default__.safeIndex(p0, len(d_922_v14_))]).cardinality, (0) - ((d_914_v7_)[default__.safeIndex(262, (d_914_v7_).length(0))])]), p0, globalState) else _dafny.SeqWithoutIsStrInference([iife75_(d_923_v15_), d_923_v15_, d_923_v15_, d_923_v15_, d_923_v15_])))
            d_928_v18_ = nw147_
            d_936_v19_: _dafny.Seq
            d_936_v19_ = _dafny.SeqWithoutIsStrInference([default__.fm30(globalState)])
            d_909_v2_ = (d_936_v19_)[default__.safeIndex((d_914_v7_)[default__.safeIndex(262, (d_914_v7_).length(0))], len(d_936_v19_))]
            d_937_v20_: _dafny.Map
            d_937_v20_ = _dafny.Map({(0) - ((d_914_v7_)[default__.safeIndex(262, (d_914_v7_).length(0))]): d_927_v17_})
            d_938_v21_: _dafny.Seq
            d_938_v21_ = _dafny.SeqWithoutIsStrInference([d_911_v4_, d_911_v4_, d_911_v4_])
            d_939_v22_: _dafny.Set
            d_939_v22_ = _dafny.Set({p2, (self).f9, p2})
            d_940_v23_: _dafny.Seq
            d_940_v23_ = d_920_v12_
            d_941_v24_: _dafny.Array
            nw148_ = _dafny.Array(None, 19)
            nw148_[int(0)] = (d_910_v3_).isdisjoint(_dafny.MultiSet([(self).f9]))
            nw148_[int(1)] = p2
            nw148_[int(2)] = not((self).f9)
            nw148_[int(3)] = p2
            nw148_[int(4)] = (p1) == (len(d_937_v20_))
            nw148_[int(5)] = (self).f9
            nw148_[int(6)] = (self).f9
            nw148_[int(7)] = (self).f9
            nw148_[int(8)] = p2
            nw148_[int(9)] = default__.fm1((_dafny.MultiSet([p1])).set(p0, default__.abs(len((d_938_v21_)[default__.safeIndex((d_914_v7_)[default__.safeIndex(262, (d_914_v7_).length(0))], len(d_938_v21_))]))), len(d_939_v22_), globalState)
            nw148_[int(10)] = (self).f9
            nw148_[int(11)] = p2
            nw148_[int(12)] = p2
            nw148_[int(13)] = (self).f9
            nw148_[int(14)] = (559) == ((0) - (p1))
            nw148_[int(15)] = False
            nw148_[int(16)] = (p2) == (not(p2))
            nw148_[int(17)] = False
            nw148_[int(18)] = (p0) in ((_dafny.SeqWithoutIsStrInference([(d_914_v7_)[default__.safeIndex(262, (d_914_v7_).length(0))], (d_914_v7_)[default__.safeIndex(262, (d_914_v7_).length(0))]])))
            d_941_v24_ = nw148_
            index145_ = default__.safeIndex(745, (d_941_v24_).length(0))
            (d_941_v24_)[index145_] = (self).f9
            index146_ = default__.safeIndex(91, (d_941_v24_).length(0))
            (d_941_v24_)[index146_] = not((d_927_v17_) not in (_dafny.SeqWithoutIsStrInference([d_927_v17_ for d_942_i3_ in range(default__.abs(-568))])))
            d_943_v25_: _dafny.Map
            d_943_v25_ = _dafny.Map({False: d_927_v17_})
            index147_ = default__.safeIndex(745, (d_941_v24_).length(0))
            index148_ = default__.safeIndex(91, (d_941_v24_).length(0))
            rhs185_ = p1
            rhs186_ = (p2) not in (d_910_v3_)
            rhs187_ = (p0) == (((d_914_v7_)[default__.safeIndex(262, (d_914_v7_).length(0))]) * (len(d_943_v25_)))
            rhs188_ = not (p2) or (False)
            rhs189_ = (((d_921_v13_) + (default__.fm28(p1, not((self).f9), globalState))).set(default__.safeIndex(p1, len((d_921_v13_) + (default__.fm28(p1, not((self).f9), globalState)))), d_927_v17_)) + (d_921_v13_)
            lhs118_ = globalState
            lhs119_ = d_941_v24_
            lhs120_ = default__.safeIndex(745, (d_941_v24_).length(0))
            lhs121_ = d_941_v24_
            lhs122_ = default__.safeIndex(91, (d_941_v24_).length(0))
            r0 = rhs185_
            lhs118_.f2 = rhs186_
            lhs119_[lhs120_] = rhs187_
            lhs121_[lhs122_] = rhs188_
            d_921_v13_ = rhs189_
            d_944_v26_: D10
            d_944_v26_ = D10_DC27(d_911_v4_)
            d_911_v4_ = (((d_944_v26_).cf50) | (d_911_v4_)) | (d_911_v4_)
        elif True:
            r0 = (p1) * (p1)
            d_945_v27_: _dafny.Seq
            d_945_v27_ = _dafny.SeqWithoutIsStrInference([p2, True])
            d_946_v28_: _dafny.Set
            d_946_v28_ = _dafny.Set({p2})
            d_947_v29_: _dafny.Array
            nw149_ = _dafny.Array(_dafny.Map({}), 27)
            d_947_v29_ = nw149_
            d_948_v30_: _dafny.Seq
            d_948_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_949_v31_: D7
            d_949_v31_ = D7_DC19(d_947_v29_, (self).f9, d_948_v30_)
            d_950_v32_: D7
            d_950_v32_ = D7_DC21(d_949_v31_)
            d_951_v33_: _dafny.Set
            d_951_v33_ = _dafny.Set({d_950_v32_, d_950_v32_})
            d_952_v34_: _dafny.MultiSet
            d_952_v34_ = _dafny.MultiSet([(0) - (p0)])
            d_953_v35_: _dafny.Array
            nw150_ = _dafny.Array(None, 19)
            nw150_[int(0)] = (d_945_v27_)[default__.safeIndex(p1, len(d_945_v27_))]
            nw150_[int(1)] = (self).f9
            nw150_[int(2)] = p2
            nw150_[int(3)] = p2
            nw150_[int(4)] = (d_946_v28_).issubset(d_946_v28_)
            nw150_[int(5)] = (d_951_v33_).ispropersubset(d_951_v33_)
            nw150_[int(6)] = True
            nw150_[int(7)] = default__.fm1(d_952_v34_, (0) - (p0), globalState)
            nw150_[int(8)] = (self).f9
            nw150_[int(9)] = (self).f9
            nw150_[int(10)] = (self).f9
            nw150_[int(11)] = p2
            nw150_[int(12)] = not ((self).f9) or (True)
            nw150_[int(13)] = not(False)
            nw150_[int(14)] = not(False)
            nw150_[int(15)] = not((d_946_v28_).ispropersubset(d_946_v28_))
            nw150_[int(16)] = (self).f9
            nw150_[int(17)] = p2
            nw150_[int(18)] = p2
            d_953_v35_ = nw150_
            index149_ = default__.safeIndex(877, (d_953_v35_).length(0))
            (d_953_v35_)[index149_] = (self).f9
            index150_ = default__.safeIndex(877, (d_953_v35_).length(0))
            (d_953_v35_)[index150_] = (self).f9
            d_954_v36_: _dafny.Set
            d_954_v36_ = _dafny.Set({len(d_948_v30_), p1})
            d_955_v37_: _dafny.Map
            d_955_v37_ = _dafny.Map({d_954_v36_: p2})
            index151_ = default__.safeIndex(877, (d_953_v35_).length(0))
            (d_953_v35_)[index151_] = not(((d_955_v37_)[d_954_v36_] if (d_954_v36_) in (d_955_v37_) else not (not(not((d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))]))) or (True)))
            d_956_v38_: str
            d_956_v38_ = _dafny.CodePoint('y')
            d_957_v39_: D1
            d_957_v39_ = D1_DC3((self).f9, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwjn"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwjn")))), d_956_v38_), p0)
            d_958_v40_: _dafny.Set
            d_958_v40_ = _dafny.Set({d_956_v38_, _dafny.CodePoint('j'), d_956_v38_, d_956_v38_})
            d_959_v41_: _dafny.Seq
            d_959_v41_ = _dafny.SeqWithoutIsStrInference([p1])
            d_960_v42_: _dafny.Map
            d_960_v42_ = _dafny.Map({p2: d_948_v30_})
            pat_let_tv47_ = d_948_v30_
            pat_let_tv48_ = p1
            pat_let_tv49_ = p1
            pat_let_tv50_ = p1
            pat_let_tv51_ = d_958_v40_
            pat_let_tv52_ = p1
            pat_let_tv53_ = p2
            pat_let_tv54_ = d_948_v30_
            pat_let_tv55_ = p1
            pat_let_tv56_ = p1
            pat_let_tv57_ = p1
            pat_let_tv58_ = d_958_v40_
            pat_let_tv59_ = p1
            pat_let_tv60_ = p2
            d_961_v43_: _dafny.Array
            nw151_ = _dafny.Array(None, 18)
            nw151_[int(0)] = (d_948_v30_) + (d_948_v30_)
            def iife84_(_pat_let31_0):
                def iife85_(d_962_dt__update__tmp_h3_):
                    def iife86_(_pat_let32_0):
                        def iife87_(d_963_dt__update_hcf7_h0_):
                            def iife88_(_pat_let33_0):
                                def iife89_(d_964_dt__update_hcf6_h0_):
                                    return D1_DC3(d_964_dt__update_hcf6_h0_, d_963_dt__update_hcf7_h0_, (d_962_dt__update__tmp_h3_).cf8)
                                return iife89_(_pat_let33_0)
                            return iife88_(True)
                        return iife87_(_pat_let32_0)
                    return iife86_(pat_let_tv47_)
                return iife85_(_pat_let31_0)
            def iife83_(_pat_let30_0):
                def iife90_(d_965_dt__update__tmp_h4_):
                    def iife91_(_pat_let34_0):
                        def iife92_(d_966_dt__update_hcf8_h0_):
                            def iife93_(_pat_let35_0):
                                def iife94_(d_967_dt__update_hcf6_h1_):
                                    return D1_DC3(d_967_dt__update_hcf6_h1_, (d_965_dt__update__tmp_h4_).cf7, d_966_dt__update_hcf8_h0_)
                                return iife94_(_pat_let35_0)
                            return iife93_(pat_let_tv53_)
                        return iife92_(_pat_let34_0)
                    return iife91_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv48_, pat_let_tv49_, pat_let_tv50_, len(pat_let_tv51_), pat_let_tv52_])))
                return iife90_(_pat_let30_0)
            def iife96_(_pat_let37_0):
                def iife97_(d_968_dt__update__tmp_h5_):
                    def iife98_(_pat_let38_0):
                        def iife99_(d_969_dt__update_hcf7_h1_):
                            def iife100_(_pat_let39_0):
                                def iife101_(d_970_dt__update_hcf6_h2_):
                                    return D1_DC3(d_970_dt__update_hcf6_h2_, d_969_dt__update_hcf7_h1_, (d_968_dt__update__tmp_h5_).cf8)
                                return iife101_(_pat_let39_0)
                            return iife100_(True)
                        return iife99_(_pat_let38_0)
                    return iife98_(pat_let_tv54_)
                return iife97_(_pat_let37_0)
            def iife95_(_pat_let36_0):
                def iife102_(d_971_dt__update__tmp_h6_):
                    def iife103_(_pat_let40_0):
                        def iife104_(d_972_dt__update_hcf8_h1_):
                            def iife105_(_pat_let41_0):
                                def iife106_(d_973_dt__update_hcf6_h3_):
                                    return D1_DC3(d_973_dt__update_hcf6_h3_, (d_971_dt__update__tmp_h6_).cf7, d_972_dt__update_hcf8_h1_)
                                return iife106_(_pat_let41_0)
                            return iife105_(pat_let_tv60_)
                        return iife104_(_pat_let40_0)
                    return iife103_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv55_, pat_let_tv56_, pat_let_tv57_, len(pat_let_tv58_), pat_let_tv59_])))
                return iife102_(_pat_let36_0)
            nw151_[int(1)] = (((iife83_(iife84_(d_957_v39_))).cf7).set(default__.safeIndex(len(d_959_v41_), len((iife95_(iife96_(d_957_v39_))).cf7)), d_956_v38_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltwca")))
            nw151_[int(2)] = d_948_v30_
            nw151_[int(3)] = d_948_v30_
            nw151_[int(4)] = d_948_v30_
            nw151_[int(5)] = (d_948_v30_ if (d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))] else (d_948_v30_).set(default__.safeIndex(p1, len(d_948_v30_)), d_956_v38_))
            nw151_[int(6)] = d_948_v30_
            nw151_[int(7)] = (d_948_v30_) + (d_948_v30_)
            nw151_[int(8)] = default__.fm28(p1, (d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))], globalState)
            nw151_[int(9)] = d_948_v30_
            nw151_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrwvirwn"))
            nw151_[int(11)] = (d_948_v30_) + (d_948_v30_)
            nw151_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "silymvy"))
            nw151_[int(13)] = d_948_v30_
            nw151_[int(14)] = ((d_960_v42_)[not((d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))])] if (not((d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))])) in (d_960_v42_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmvllgg")))
            nw151_[int(15)] = d_948_v30_
            nw151_[int(16)] = (_dafny.SeqWithoutIsStrInference([d_956_v38_ for d_974_i4_ in range(default__.abs(-884))])) + (d_948_v30_)
            nw151_[int(17)] = d_948_v30_
            d_961_v43_ = nw151_
            index152_ = default__.safeIndex(370, (d_961_v43_).length(0))
            (d_961_v43_)[index152_] = d_948_v30_
            d_975_v44_: _dafny.Map
            d_975_v44_ = _dafny.Map({p0: False})
            d_976_v45_: D9
            d_976_v45_ = D9_DC24(((d_975_v44_)[p0] if (p0) in (d_975_v44_) else p2), p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))
            index153_ = default__.safeIndex(370, (d_961_v43_).length(0))
            (d_961_v43_)[index153_] = (d_976_v45_).cf43
            if (d_945_v27_)[default__.safeIndex(p0, len(d_945_v27_))]:
                d_956_v38_ = default__.fm31(globalState)
                d_977_v46_: _dafny.Map
                d_977_v46_ = _dafny.Map({_dafny.MultiSet([470, p1, p1, p0]): (self).f9})
                d_978_v47_: D7
                d_978_v47_ = D7_DC18((_dafny.Map({_dafny.MultiSet(d_959_v41_): p2})) | (d_977_v46_))
                index154_ = default__.safeIndex(877, (d_953_v35_).length(0))
                rhs190_ = (0) - ((0) - (p1))
                rhs191_ = p2
                rhs192_ = d_978_v47_
                lhs123_ = d_953_v35_
                lhs124_ = default__.safeIndex(877, (d_953_v35_).length(0))
                r0 = rhs190_
                lhs123_[lhs124_] = rhs191_
                d_978_v47_ = rhs192_
                r2 = default__.fm21(len((d_945_v27_).set(default__.safeIndex(p0, len(d_945_v27_)), (self).f9)), default__.fm1(d_952_v34_, p1, globalState), globalState)
                d_979_v48_: D4
                d_979_v48_ = D4_DC14(p2, p2, p1)
                r0 = (d_979_v48_).cf28
                (globalState).f2 = False
            elif True:
                d_980_v49_: _dafny.Array
                nw152_ = _dafny.Array(int(0), 11)
                d_980_v49_ = nw152_
                index155_ = default__.safeIndex(555, (d_980_v49_).length(0))
                (d_980_v49_)[index155_] = default__.safeDivisionInt(p0, len(_dafny.SeqWithoutIsStrInference([(d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))]])))
                index156_ = default__.safeIndex(555, (d_980_v49_).length(0))
                (d_980_v49_)[index156_] = default__.safeModuloInt(-642, (0) - (p0))
                d_948_v30_ = d_948_v30_
                d_981_v50_: C2
                nw153_ = C2()
                nw153_.ctor__(len((d_948_v30_) + (d_948_v30_)), (d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))])
                d_981_v50_ = nw153_
                index157_ = default__.safeIndex(370, (d_961_v43_).length(0))
                (d_961_v43_)[index157_] = d_948_v30_
                d_982_v51_: _dafny.MultiSet
                d_982_v51_ = _dafny.MultiSet([p2])
                d_983_v52_: D1
                d_983_v52_ = D1_DC4(385, (d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))], p2, (d_953_v35_)[default__.safeIndex(877, (d_953_v35_).length(0))], p2)
                r2 = (d_982_v51_).set((d_983_v52_).cf12, default__.abs(len(d_945_v27_)))
        source11_ = D3_DC10()
        if source11_.is_DC8:
            d_984_v53_: _dafny.Seq
            d_984_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujap"))
            r0 = len(d_984_v53_)
            d_985_v54_: _dafny.Array
            nw154_ = _dafny.Array(int(0), 13)
            d_985_v54_ = nw154_
            d_986_v55_: _dafny.Array
            d_986_v55_ = d_985_v54_
            d_985_v54_ = (d_986_v55_)
            source12_ = D1_DC3(p2, d_984_v53_, p1)
            if source12_.is_DC2:
                d_987___mcc_h3_ = source12_.cf2
                d_988___mcc_h4_ = source12_.cf3
                d_989___mcc_h5_ = source12_.cf4
                d_990___mcc_h6_ = source12_.cf5
                d_991_cf5_ = d_990___mcc_h6_
                d_992_cf4_ = d_989___mcc_h5_
                d_993_cf3_ = d_988___mcc_h4_
                d_994_cf2_ = d_987___mcc_h3_
                d_992_cf4_ = (d_991_cf5_) * (d_991_cf5_)
                d_995_v56_: _dafny.Map
                d_995_v56_ = _dafny.Map({(self).f9: d_993_cf3_})
                d_996_v57_: _dafny.MultiSet
                d_996_v57_ = _dafny.MultiSet([p0, len(d_995_v56_)])
                d_997_v58_: str
                d_997_v58_ = _dafny.CodePoint('h')
                d_998_v59_: D7
                d_998_v59_ = D7_DC20(False, d_997_v58_)
                d_999_v60_: D4
                d_999_v60_ = D4_DC15(D4_DC12(default__.fm1(d_996_v57_, d_992_cf4_, globalState), False, d_994_cf2_, d_994_cf2_, (d_998_v59_).cf36))
                d_999_v60_ = d_999_v60_
                d_1000_v61_: _dafny.Array
                def lambda50_(d_1001_v53_, d_1002_p0_, d_1003_v58_, d_1004_p1_):
                    def lambda51_(d_1005_i5_):
                        return not((len((d_1001_v53_).set(default__.safeIndex((0) - (d_1002_p0_), len(d_1001_v53_)), d_1003_v58_))) < (d_1004_p1_))

                    return lambda51_

                init28_ = lambda50_(d_984_v53_, p0, d_997_v58_, p1)
                nw155_ = _dafny.Array(None, 24)
                for i0_28_ in range(nw155_.length(0)):
                    nw155_[i0_28_] = init28_(i0_28_)
                d_1000_v61_ = nw155_
                d_1000_v61_ = d_1000_v61_
                d_1006_v62_: C2
                nw156_ = C2()
                nw156_.ctor__(d_994_cf2_, p2)
                d_1006_v62_ = nw156_
                d_1007_v63_: _dafny.Map
                d_1007_v63_ = _dafny.Map({d_1006_v62_: d_993_cf3_})
                d_1008_v64_: _dafny.Array
                nw157_ = _dafny.Array(_dafny.Set({}), 4)
                d_1008_v64_ = nw157_
                d_1009_v65_: T2
                nw158_ = C3()
                nw158_.ctor__(((d_1007_v63_)[d_1006_v62_] if (d_1006_v62_) in (d_1007_v63_) else not(False)), d_1008_v64_, d_993_cf3_)
                d_1009_v65_ = nw158_
                d_1010_v66_: _dafny.Seq
                d_1010_v66_ = _dafny.SeqWithoutIsStrInference([d_1009_v65_, d_1009_v65_])
                r3 = not(default__.fm1(d_996_v57_, len(d_1010_v66_), globalState))
            elif source12_.is_DC3:
                d_1011___mcc_h7_ = source12_.cf6
                d_1012___mcc_h8_ = source12_.cf7
                d_1013___mcc_h9_ = source12_.cf8
                d_1014_cf8_ = d_1013___mcc_h9_
                d_1015_cf7_ = d_1012___mcc_h8_
                d_1016_cf6_ = d_1011___mcc_h7_
                (globalState).f2 = d_1016_cf6_
                d_1017_v67_: _dafny.Array
                def lambda52_(d_1018_cf6_, d_1019_p2_):
                    def lambda53_(d_1020_i6_):
                        return (d_1018_cf6_ if not(d_1018_cf6_) else d_1019_p2_)

                    return lambda53_

                init29_ = lambda52_(d_1016_cf6_, p2)
                nw159_ = _dafny.Array(None, 14)
                for i0_29_ in range(nw159_.length(0)):
                    nw159_[i0_29_] = init29_(i0_29_)
                d_1017_v67_ = nw159_
                d_1021_v68_: _dafny.Seq
                d_1021_v68_ = _dafny.SeqWithoutIsStrInference([p1, d_1014_cf8_])
                d_1022_v69_: _dafny.Map
                d_1022_v69_ = _dafny.Map({d_1021_v68_: p0})
                d_1023_v71_: C0
                nw160_ = C0()
                def iife107_():
                    coll23_ = _dafny.Set()
                    compr_23_: _dafny.Seq
                    for compr_23_ in (d_1022_v69_).keys.Elements:
                        d_1024_v70_: _dafny.Seq = compr_23_
                        if (d_1024_v70_) in (d_1022_v69_):
                            coll23_ = coll23_.union(_dafny.Set([d_1024_v70_]))
                    return _dafny.Set(coll23_)
                nw160_.ctor__(iife107_()
                , d_1014_cf8_)
                d_1023_v71_ = nw160_
                d_1025_v72_: _dafny.MultiSet
                d_1025_v72_ = _dafny.MultiSet([d_1023_v71_, d_1023_v71_, d_1023_v71_])
                index158_ = default__.safeIndex(497, (d_1017_v67_).length(0))
                (d_1017_v67_)[index158_] = default__.fm1(_dafny.MultiSet([p1, p0]), (d_1025_v72_).cardinality, globalState)
                index159_ = default__.safeIndex(497, (d_1017_v67_).length(0))
                (d_1017_v67_)[index159_] = d_1016_cf6_
                d_1026_v73_: _dafny.Map
                d_1026_v73_ = _dafny.Map({d_1023_v71_.f12: (d_1023_v71_.f12) + ((0) - ((0) - (p0)))})
                d_1026_v73_ = (d_1026_v73_).set((0) - (p0), default__.fm0(globalState))
                d_1027_v74_: C2
                nw161_ = C2()
                nw161_.ctor__(d_1014_cf8_, d_1016_cf6_)
                d_1027_v74_ = nw161_
                d_1028_v75_: _dafny.Seq
                d_1028_v75_ = _dafny.SeqWithoutIsStrInference([d_1027_v74_, d_1027_v74_])
                d_1029_v76_: _dafny.Map
                d_1029_v76_ = _dafny.Map({d_1028_v75_: d_1017_v67_})
                d_1030_v77_: _dafny.Array
                nw162_ = _dafny.Array(None, 4)
                nw162_[int(0)] = ((d_1029_v76_)[d_1028_v75_] if (d_1028_v75_) in (d_1029_v76_) else d_1017_v67_)
                nw162_[int(1)] = d_1017_v67_
                nw162_[int(2)] = d_1017_v67_
                nw162_[int(3)] = d_1017_v67_
                d_1030_v77_ = nw162_
                index160_ = default__.safeIndex(469, (d_1030_v77_).length(0))
                (d_1030_v77_)[index160_] = d_1017_v67_
                d_1031_v78_: _dafny.Map
                d_1031_v78_ = _dafny.Map({len(d_984_v53_): d_1017_v67_})
                index161_ = default__.safeIndex(469, (d_1030_v77_).length(0))
                (d_1030_v77_)[index161_] = ((d_1031_v78_)[d_1023_v71_.f12] if (d_1023_v71_.f12) in (d_1031_v78_) else d_1017_v67_)
            elif source12_.is_DC4:
                d_1032___mcc_h10_ = source12_.cf9
                d_1033___mcc_h11_ = source12_.cf10
                d_1034___mcc_h12_ = source12_.cf11
                d_1035___mcc_h13_ = source12_.cf12
                d_1036___mcc_h14_ = source12_.cf13
                d_1037_cf13_ = d_1036___mcc_h14_
                d_1038_cf12_ = d_1035___mcc_h13_
                d_1039_cf11_ = d_1034___mcc_h12_
                d_1040_cf10_ = d_1033___mcc_h11_
                d_1041_cf9_ = d_1032___mcc_h10_
                d_1042_v79_: str
                d_1042_v79_ = _dafny.CodePoint('d')
                d_1042_v79_ = d_1042_v79_
                d_1043_v80_: _dafny.Array
                nw163_ = _dafny.Array(_dafny.Set({}), 21)
                d_1043_v80_ = nw163_
                d_1043_v80_ = d_1043_v80_
                (globalState).f2 = True
                (globalState).f2 = (True) and ((self).f9)
            elif source12_.is_DC1:
                d_1044___mcc_h15_ = source12_.cf1
                d_1045_cf1_ = d_1044___mcc_h15_
                d_1045_cf1_ = default__.safeModuloInt(p1, p0)
                d_1046_v81_: str
                d_1046_v81_ = _dafny.CodePoint('x')
                d_984_v53_ = ((d_984_v53_) + (d_984_v53_)) + (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jor"))).set(default__.safeIndex(len(d_984_v53_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jor")))), d_1046_v81_)).set(default__.safeIndex(d_1045_cf1_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jor"))).set(default__.safeIndex(len(d_984_v53_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jor")))), d_1046_v81_))), default__.fm31(globalState)))
                r3 = True
                r3 = p2
            elif True:
                d_1047___mcc_h16_ = source12_.cf14
                d_1048_cf14_ = d_1047___mcc_h16_
                d_1049_v82_: _dafny.Map
                d_1049_v82_ = _dafny.Map({not(p2): p2})
                r0 = (p0) * ((p1) - (len(d_1049_v82_)))
                d_1050_v83_: _dafny.Array
                def lambda54_(d_1051_i7_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmokwm"))

                init30_ = lambda54_
                nw164_ = _dafny.Array(None, 17)
                for i0_30_ in range(nw164_.length(0)):
                    nw164_[i0_30_] = init30_(i0_30_)
                d_1050_v83_ = nw164_
                d_1050_v83_ = d_1050_v83_
                r0 = default__.safeDivisionInt(p1, len((d_1049_v82_) | (d_1049_v82_)))
                d_1052_v84_: D1
                d_1052_v84_ = D1_DC2(p1, True, p1, default__.fm0(globalState))
                d_1053_v85_: _dafny.Seq
                d_1053_v85_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                d_1054_v86_: C1
                nw165_ = C1()
                nw165_.ctor__((_dafny.SeqWithoutIsStrInference([d_1052_v84_, d_1052_v84_, d_1052_v84_, d_1052_v84_, D1_DC2(-578, (self).f9, len(d_1053_v85_), p0)])) + (_dafny.SeqWithoutIsStrInference([d_1052_v84_, d_1052_v84_, d_1052_v84_, d_1052_v84_, D1_DC2(default__.fm0(globalState), p2, p0, p0)])))
                d_1054_v86_ = nw165_
            d_1055_v87_: _dafny.Map
            d_1055_v87_ = _dafny.Map({d_985_v54_: (self).f9})
            d_1055_v87_ = (d_1055_v87_).set(d_985_v54_, not((self).f9))
        elif source11_.is_DC9:
            d_1056___mcc_h0_ = source11_.cf17
            d_1057___mcc_h1_ = source11_.cf18
            d_1058_cf18_ = d_1057___mcc_h1_
            d_1059_cf17_ = d_1056___mcc_h0_
            d_1060_v88_: _dafny.Array
            def lambda55_(d_1061_p1_):
                def lambda56_(d_1062_i8_):
                    return default__.safeDivisionInt(d_1062_i8_, d_1061_p1_)

                return lambda56_

            init31_ = lambda55_(p1)
            nw166_ = _dafny.Array(None, 11)
            for i0_31_ in range(nw166_.length(0)):
                nw166_[i0_31_] = init31_(i0_31_)
            d_1060_v88_ = nw166_
            index162_ = default__.safeIndex(663, (d_1060_v88_).length(0))
            (d_1060_v88_)[index162_] = len(d_1058_cf18_)
            index163_ = default__.safeIndex(663, (d_1060_v88_).length(0))
            (d_1060_v88_)[index163_] = p1
            d_1063_v89_: str
            d_1063_v89_ = _dafny.CodePoint('v')
            d_1064_v90_: _dafny.Map
            d_1064_v90_ = _dafny.Map({d_1063_v89_: d_1063_v89_})
            d_1065_v91_: _dafny.Seq
            d_1065_v91_ = _dafny.SeqWithoutIsStrInference([d_1063_v89_, d_1063_v89_])
            d_1064_v90_ = ((d_1064_v90_) | (d_1064_v90_)) | ((d_1064_v90_) | (_dafny.Map({_dafny.CodePoint('u'): (d_1065_v91_)[default__.safeIndex((d_1060_v88_)[default__.safeIndex(663, (d_1060_v88_).length(0))], len(d_1065_v91_))]})))
            d_1066_v92_: D10
            d_1066_v92_ = D10_DC28()
            d_1067_v93_: _dafny.Map
            d_1067_v93_ = _dafny.Map({p2: p1})
            d_1068_v94_: _dafny.Map
            d_1068_v94_ = _dafny.Map({d_1066_v92_: d_1067_v93_})
            d_1069_v95_: _dafny.MultiSet
            d_1069_v95_ = _dafny.MultiSet([len(d_1068_v94_)])
            rhs193_ = (0) - (((d_1060_v88_)[default__.safeIndex(663, (d_1060_v88_).length(0))]) * ((0) - (p0)))
            rhs194_ = default__.fm1(d_1069_v95_, (d_1060_v88_)[default__.safeIndex(663, (d_1060_v88_).length(0))], globalState)
            rhs195_ = p1
            r0 = rhs193_
            d_1059_cf17_ = rhs194_
            r0 = rhs195_
            d_1070_v96_: D1
            d_1070_v96_ = D1_DC2(p0, False, p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chnn"))))
            d_1071_v97_: _dafny.Seq
            d_1071_v97_ = _dafny.SeqWithoutIsStrInference([d_1070_v96_])
            d_1072_v98_: C1
            nw167_ = C1()
            nw167_.ctor__((d_1071_v97_) + (d_1071_v97_))
            d_1072_v98_ = nw167_
            d_1073_v99_: C2
            nw168_ = C2()
            nw168_.ctor__((d_1060_v88_)[default__.safeIndex(663, (d_1060_v88_).length(0))], not(p2))
            d_1073_v99_ = nw168_
            d_1074_v100_: D4
            d_1074_v100_ = D4_DC13((d_1073_v99_).f14)
            d_1075_v101_: _dafny.Set
            d_1075_v101_ = _dafny.Set({d_1059_cf17_, p2})
            pat_let_tv61_ = d_1069_v95_
            pat_let_tv62_ = d_1060_v88_
            pat_let_tv63_ = d_1060_v88_
            pat_let_tv64_ = p1
            d_1076_v102_: _dafny.Array
            nw169_ = _dafny.Array(None, 25)
            nw169_[int(0)] = d_1074_v100_
            nw169_[int(1)] = d_1074_v100_
            nw169_[int(2)] = d_1074_v100_
            nw169_[int(3)] = d_1074_v100_
            nw169_[int(4)] = d_1074_v100_
            nw169_[int(5)] = d_1074_v100_
            nw169_[int(6)] = d_1074_v100_
            def iife108_(_pat_let42_0):
                def iife109_(d_1077_dt__update__tmp_h7_):
                    def iife110_(_pat_let43_0):
                        def iife111_(d_1078_dt__update_hcf25_h0_):
                            return D4_DC13(d_1078_dt__update_hcf25_h0_)
                        return iife111_(_pat_let43_0)
                    return iife110_((pat_let_tv61_).cardinality)
                return iife109_(_pat_let42_0)
            nw169_[int(7)] = iife108_(d_1074_v100_)
            nw169_[int(8)] = d_1074_v100_
            nw169_[int(9)] = d_1074_v100_
            def iife112_(_pat_let44_0):
                def iife113_(d_1079_dt__update__tmp_h8_):
                    def iife114_(_pat_let45_0):
                        def iife115_(d_1080_dt__update_hcf25_h1_):
                            return D4_DC13(d_1080_dt__update_hcf25_h1_)
                        return iife115_(_pat_let45_0)
                    return iife114_((pat_let_tv63_)[default__.safeIndex(663, (pat_let_tv62_).length(0))])
                return iife113_(_pat_let44_0)
            nw169_[int(10)] = iife112_(d_1074_v100_)
            nw169_[int(11)] = d_1074_v100_
            nw169_[int(12)] = d_1074_v100_
            nw169_[int(13)] = d_1074_v100_
            nw169_[int(14)] = d_1074_v100_
            def iife116_(_pat_let46_0):
                def iife117_(d_1081_dt__update__tmp_h9_):
                    def iife118_(_pat_let47_0):
                        def iife119_(d_1082_dt__update_hcf25_h2_):
                            return D4_DC13(d_1082_dt__update_hcf25_h2_)
                        return iife119_(_pat_let47_0)
                    return iife118_(pat_let_tv64_)
                return iife117_(_pat_let46_0)
            nw169_[int(15)] = iife116_(d_1074_v100_)
            nw169_[int(16)] = d_1074_v100_
            nw169_[int(17)] = d_1074_v100_
            nw169_[int(18)] = d_1074_v100_
            nw169_[int(19)] = d_1074_v100_
            nw169_[int(20)] = D4_DC13((0) - (len(d_1075_v101_)))
            nw169_[int(21)] = d_1074_v100_
            nw169_[int(22)] = D4_DC13(p0)
            nw169_[int(23)] = d_1074_v100_
            nw169_[int(24)] = d_1074_v100_
            d_1076_v102_ = nw169_
            d_1083_v103_: _dafny.Array
            nw170_ = _dafny.Array(None, 13)
            nw170_[int(0)] = d_1076_v102_
            nw170_[int(1)] = d_1076_v102_
            nw170_[int(2)] = d_1076_v102_
            nw170_[int(3)] = d_1076_v102_
            nw170_[int(4)] = d_1076_v102_
            nw170_[int(5)] = d_1076_v102_
            nw170_[int(6)] = d_1076_v102_
            nw170_[int(7)] = d_1076_v102_
            nw170_[int(8)] = d_1076_v102_
            nw170_[int(9)] = d_1076_v102_
            nw170_[int(10)] = d_1076_v102_
            nw170_[int(11)] = d_1076_v102_
            nw170_[int(12)] = d_1076_v102_
            d_1083_v103_ = nw170_
            index164_ = default__.safeIndex(663, (d_1060_v88_).length(0))
            rhs196_ = d_1072_v98_
            rhs197_ = 416
            rhs198_ = d_1073_v99_
            rhs199_ = d_1083_v103_
            lhs125_ = d_1060_v88_
            lhs126_ = default__.safeIndex(663, (d_1060_v88_).length(0))
            d_1072_v98_ = rhs196_
            lhs125_[lhs126_] = rhs197_
            d_1073_v99_ = rhs198_
            d_1083_v103_ = rhs199_
        elif source11_.is_DC10:
            r3 = (self).f9
            d_1084_v104_: str
            d_1084_v104_ = _dafny.CodePoint('k')
            d_1085_v105_: _dafny.Set
            d_1085_v105_ = _dafny.Set({d_1084_v104_, d_1084_v104_})
            d_1086_v106_: _dafny.Seq
            d_1086_v106_ = _dafny.SeqWithoutIsStrInference([True])
            d_1087_v107_: _dafny.Set
            d_1087_v107_ = _dafny.Set({p0})
            d_1088_v108_: _dafny.MultiSet
            d_1088_v108_ = _dafny.MultiSet([p1])
            d_1089_v109_: _dafny.Array
            nw171_ = _dafny.Array(None, 29)
            nw171_[int(0)] = (self).f9
            nw171_[int(1)] = p2
            nw171_[int(2)] = (self).f9
            nw171_[int(3)] = p2
            nw171_[int(4)] = (d_1086_v106_)[default__.safeIndex(len(d_1087_v107_), len(d_1086_v106_))]
            nw171_[int(5)] = p2
            nw171_[int(6)] = (self).f9
            nw171_[int(7)] = False
            nw171_[int(8)] = p2
            nw171_[int(9)] = (self).f9
            nw171_[int(10)] = True
            nw171_[int(11)] = (self).f9
            nw171_[int(12)] = (self).f9
            nw171_[int(13)] = (self).f9
            nw171_[int(14)] = not(p2)
            nw171_[int(15)] = default__.fm1(d_1088_v108_, p1, globalState)
            nw171_[int(16)] = (self).f9
            nw171_[int(17)] = p2
            nw171_[int(18)] = (self).f9
            nw171_[int(19)] = True
            nw171_[int(20)] = (self).f9
            nw171_[int(21)] = default__.fm1(d_1088_v108_, (0) - (p1), globalState)
            nw171_[int(22)] = (self).f9
            nw171_[int(23)] = (self).f9
            nw171_[int(24)] = (self).f9
            nw171_[int(25)] = (self).f9
            nw171_[int(26)] = (self).f9
            nw171_[int(27)] = not(p2)
            nw171_[int(28)] = not((self).f9)
            d_1089_v109_ = nw171_
            d_1090_v110_: _dafny.Map
            d_1090_v110_ = _dafny.Map({(d_1085_v105_) == (d_1085_v105_): d_1089_v109_})
            d_1091_v111_: _dafny.Seq
            d_1091_v111_ = _dafny.SeqWithoutIsStrInference([d_1090_v110_])
            d_1092_v112_: C2
            nw172_ = C2()
            nw172_.ctor__(-876, p2)
            d_1092_v112_ = nw172_
            d_1093_v113_: _dafny.Seq
            d_1093_v113_ = _dafny.SeqWithoutIsStrInference([d_1092_v112_])
            d_1090_v110_ = (((d_1091_v111_)[default__.safeIndex(len(d_1093_v113_), len(d_1091_v111_))]) | (d_1090_v110_)).set(default__.fm1(_dafny.MultiSet([p1, p0]), len(_dafny.Set({(self).f9})), globalState), d_1089_v109_)
            r0 = (d_1092_v112_).f14
            r0 = (0) - ((p1) + (default__.fm0(globalState)))
        elif True:
            d_1094___mcc_h2_ = source11_.cf16
            d_1095_cf16_ = d_1094___mcc_h2_
            d_1096_v114_: D4
            d_1096_v114_ = D4_DC13((0) - (p0))
            d_1097_v115_: str
            d_1097_v115_ = _dafny.CodePoint('y')
            d_1098_v116_: C1
            nw173_ = C1()
            nw173_.ctor__(_dafny.SeqWithoutIsStrInference([D1_DC2(len(_dafny.Map({(self).f9: p0})), p2, p0, p0) for d_1099_i9_ in range(default__.abs(670))]))
            d_1098_v116_ = nw173_
            d_1100_v117_: _dafny.MultiSet
            d_1100_v117_ = _dafny.MultiSet([(D13_DC35(p2, d_1097_v115_, d_1098_v116_, (self).f9)).cf61, not(False)])
            d_1101_v118_: _dafny.Map
            d_1101_v118_ = _dafny.Map({d_1097_v115_: ((d_1100_v117_)[(self).f9] if ((self).f9) in (d_1100_v117_) else p0)})
            d_1102_v119_: _dafny.Map
            d_1102_v119_ = _dafny.Map({d_1096_v114_: ((d_1101_v118_)[d_1097_v115_] if (d_1097_v115_) in (d_1101_v118_) else p1)})
            d_1102_v119_ = (d_1102_v119_).set(d_1096_v114_, p1)
            d_1103_v120_: _dafny.Array
            nw174_ = _dafny.Array(_dafny.Map({}), 26)
            d_1103_v120_ = nw174_
            d_1104_v121_: _dafny.Map
            d_1104_v121_ = _dafny.Map({p1: (self).f9})
            index165_ = default__.safeIndex(348, (d_1103_v120_).length(0))
            (d_1103_v120_)[index165_] = (d_1104_v121_) | (d_1104_v121_)
            d_1105_v123_: _dafny.MultiSet
            d_1105_v123_ = _dafny.MultiSet([p1])
            index166_ = default__.safeIndex(348, (d_1103_v120_).length(0))
            def iife120_():
                coll24_ = _dafny.Map()
                compr_24_: int
                for compr_24_ in (d_1105_v123_).Elements:
                    d_1106_v122_: int = compr_24_
                    if (d_1106_v122_) in (d_1105_v123_):
                        coll24_[default__.safeModuloInt(d_1106_v122_, p0)] = False
                return _dafny.Map(coll24_)
            (d_1103_v120_)[index166_] = (d_1104_v121_) | (iife120_()
            )
            d_1107_v124_: _dafny.Seq
            d_1107_v124_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_1108_v125_: D4
            d_1108_v125_ = D4_DC12(p2, p2, len(d_1107_v124_), p1, (self).f9)
            d_1109_v126_: D4
            d_1109_v126_ = D4_DC15((d_1108_v125_ if p2 else d_1108_v125_))
            source13_ = d_1109_v126_
            if source13_.is_DC12:
                d_1110___mcc_h17_ = source13_.cf20
                d_1111___mcc_h18_ = source13_.cf21
                d_1112___mcc_h19_ = source13_.cf22
                d_1113___mcc_h20_ = source13_.cf23
                d_1114___mcc_h21_ = source13_.cf24
                d_1115_cf24_ = d_1114___mcc_h21_
                d_1116_cf23_ = d_1113___mcc_h20_
                d_1117_cf22_ = d_1112___mcc_h19_
                d_1118_cf21_ = d_1111___mcc_h18_
                d_1119_cf20_ = d_1110___mcc_h17_
                d_1120_v127_: _dafny.Array
                def lambda57_(d_1121_cf23_):
                    def lambda58_(d_1122_i10_):
                        return default__.safeDivisionInt(d_1122_i10_, d_1121_cf23_)

                    return lambda58_

                init32_ = lambda57_(d_1116_cf23_)
                nw175_ = _dafny.Array(None, 6)
                for i0_32_ in range(nw175_.length(0)):
                    nw175_[i0_32_] = init32_(i0_32_)
                d_1120_v127_ = nw175_
                r0 = len(_dafny.Map({d_1120_v127_: d_1120_v127_}))
                d_1120_v127_ = d_1120_v127_
                d_1123_v128_: C1
                nw176_ = C1()
                nw176_.ctor__(d_1098_v116_.f13)
                d_1123_v128_ = nw176_
                d_1124_v129_: _dafny.Seq
                d_1124_v129_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1117_cf22_: p0}))])
                d_1125_v130_: _dafny.Set
                d_1125_v130_ = _dafny.Set({p0})
                d_1126_v131_: _dafny.Map
                d_1126_v131_ = _dafny.Map({d_1118_cf21_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acs")))})
                d_1127_v132_: _dafny.Array
                def lambda59_(d_1128_cf20_):
                    def lambda60_(d_1129_i11_):
                        return d_1128_cf20_

                    return lambda60_

                init33_ = lambda59_(d_1119_cf20_)
                nw177_ = _dafny.Array(None, 2)
                for i0_33_ in range(nw177_.length(0)):
                    nw177_[i0_33_] = init33_(i0_33_)
                d_1127_v132_ = nw177_
                d_1130_v133_: _dafny.Map
                d_1130_v133_ = _dafny.Map({d_1127_v132_: d_1117_cf22_})
                d_1131_v134_: _dafny.Array
                nw178_ = _dafny.Array(None, 19)
                nw178_[int(0)] = d_1125_v130_
                nw178_[int(1)] = d_1125_v130_
                nw178_[int(2)] = d_1125_v130_
                nw178_[int(3)] = d_1125_v130_
                nw178_[int(4)] = d_1125_v130_
                nw178_[int(5)] = d_1125_v130_
                nw178_[int(6)] = _dafny.Set({p0})
                nw178_[int(7)] = d_1125_v130_
                nw178_[int(8)] = _dafny.Set({len(d_1126_v131_)})
                nw178_[int(9)] = d_1125_v130_
                nw178_[int(10)] = d_1125_v130_
                nw178_[int(11)] = d_1125_v130_
                nw178_[int(12)] = d_1125_v130_
                nw178_[int(13)] = d_1125_v130_
                nw178_[int(14)] = _dafny.Set({len(d_1130_v133_), d_1116_cf23_, 748})
                nw178_[int(15)] = d_1125_v130_
                nw178_[int(16)] = d_1125_v130_
                nw178_[int(17)] = d_1125_v130_
                nw178_[int(18)] = d_1125_v130_
                d_1131_v134_ = nw178_
                d_1132_v135_: T2
                nw179_ = C3()
                nw179_.ctor__(d_1118_cf21_, d_1131_v134_, d_1119_cf20_)
                d_1132_v135_ = nw179_
                d_1133_v136_: _dafny.Seq
                d_1133_v136_ = _dafny.SeqWithoutIsStrInference([d_1132_v135_, d_1132_v135_])
                d_1134_v137_: _dafny.MultiSet
                d_1134_v137_ = _dafny.MultiSet([d_1097_v115_])
                d_1135_v138_: _dafny.Array
                nw180_ = _dafny.Array(None, 24)
                nw180_[int(0)] = (62) != ((0) - (d_1117_cf22_))
                nw180_[int(1)] = not(default__.fm1(d_1105_v123_, 186, globalState))
                nw180_[int(2)] = (d_1097_v115_) in (default__.fm28((d_1124_v129_)[default__.safeIndex((d_1105_v123_).cardinality, len(d_1124_v129_))], p2, globalState))
                nw180_[int(3)] = (len(d_1133_v136_)) > (-521)
                nw180_[int(4)] = not(not(d_1119_cf20_))
                nw180_[int(5)] = True
                nw180_[int(6)] = (d_1132_v135_).f9
                nw180_[int(7)] = (self).f9
                nw180_[int(8)] = (d_1134_v137_).issubset(d_1134_v137_)
                nw180_[int(9)] = (d_1132_v135_).f9
                nw180_[int(10)] = True
                nw180_[int(11)] = not (p2) or (not(d_1119_cf20_))
                nw180_[int(12)] = (self).f9
                nw180_[int(13)] = d_1115_cf24_
                nw180_[int(14)] = (self).f9
                nw180_[int(15)] = p2
                nw180_[int(16)] = ((self).fm2(d_1116_cf23_, (d_1132_v135_).f9, not(((d_1104_v121_)[p0] if (p0) in (d_1104_v121_) else (d_1132_v135_).f9)), default__.fm0(globalState), globalState)) <= (d_1124_v129_)
                nw180_[int(17)] = False
                nw180_[int(18)] = d_1118_cf21_
                nw180_[int(19)] = p2
                nw180_[int(20)] = d_1118_cf21_
                nw180_[int(21)] = d_1119_cf20_
                nw180_[int(22)] = d_1119_cf20_
                nw180_[int(23)] = d_1119_cf20_
                d_1135_v138_ = nw180_
                index167_ = default__.safeIndex(235, (d_1127_v132_).length(0))
                (d_1127_v132_)[index167_] = True
                index168_ = default__.safeIndex(665, (d_1120_v127_).length(0))
                (d_1120_v127_)[index168_] = d_1117_cf22_
                d_1136_v139_: _dafny.Seq
                d_1136_v139_ = _dafny.SeqWithoutIsStrInference([d_1126_v131_, d_1126_v131_, d_1126_v131_, d_1126_v131_])
                index169_ = default__.safeIndex(633, (d_1120_v127_).length(0))
                (d_1120_v127_)[index169_] = len(d_1136_v139_)
                d_1137_v140_: _dafny.MultiSet
                d_1137_v140_ = _dafny.MultiSet([d_1120_v127_, d_1120_v127_, d_1120_v127_])
                index170_ = default__.safeIndex(235, (d_1127_v132_).length(0))
                index171_ = default__.safeIndex(665, (d_1120_v127_).length(0))
                index172_ = default__.safeIndex(633, (d_1120_v127_).length(0))
                rhs200_ = (d_1137_v140_).isdisjoint(d_1137_v140_)
                rhs201_ = p1
                rhs202_ = ((d_1126_v131_)[(d_1100_v117_).ispropersubset(d_1100_v117_)] if ((d_1100_v117_).ispropersubset(d_1100_v117_)) in (d_1126_v131_) else d_1117_cf22_)
                rhs203_ = (0) - (d_1116_cf23_)
                rhs204_ = True
                lhs127_ = d_1127_v132_
                lhs128_ = default__.safeIndex(235, (d_1127_v132_).length(0))
                lhs129_ = d_1120_v127_
                lhs130_ = default__.safeIndex(665, (d_1120_v127_).length(0))
                lhs131_ = d_1120_v127_
                lhs132_ = default__.safeIndex(633, (d_1120_v127_).length(0))
                lhs133_ = globalState
                lhs127_[lhs128_] = rhs200_
                lhs129_[lhs130_] = rhs201_
                lhs131_[lhs132_] = rhs202_
                d_1117_cf22_ = rhs203_
                lhs133_.f2 = rhs204_
            elif source13_.is_DC13:
                d_1138___mcc_h22_ = source13_.cf25
                d_1139_cf25_ = d_1138___mcc_h22_
                d_1140_v141_: _dafny.Array
                def lambda61_(d_1141_cf25_):
                    def lambda62_(d_1142_i12_):
                        return (d_1142_i12_) + (d_1141_cf25_)

                    return lambda62_

                init34_ = lambda61_(d_1139_cf25_)
                nw181_ = _dafny.Array(None, 12)
                for i0_34_ in range(nw181_.length(0)):
                    nw181_[i0_34_] = init34_(i0_34_)
                d_1140_v141_ = nw181_
                index173_ = default__.safeIndex(39, (d_1140_v141_).length(0))
                (d_1140_v141_)[index173_] = (p1 if (self).f9 else d_1139_cf25_)
                index174_ = default__.safeIndex(39, (d_1140_v141_).length(0))
                (d_1140_v141_)[index174_] = default__.safeModuloInt(d_1139_cf25_, len(_dafny.SeqWithoutIsStrInference([d_1097_v115_ for d_1143_i13_ in range(default__.abs(34))])))
                d_1144_v142_: D3
                d_1144_v142_ = D3_DC9(default__.fm1(_dafny.MultiSet([527, d_1139_cf25_, p1, d_1139_cf25_]), p0, globalState), _dafny.SeqWithoutIsStrInference([not(True)]))
                d_1145_v143_: _dafny.Map
                d_1145_v143_ = _dafny.Map({((0) - (p1)) < ((0) - (p1)): (d_1144_v142_).cf17})
                d_1145_v143_ = (d_1145_v143_).set(p2, not (False) or (p2))
                d_1146_v144_: _dafny.Array
                nw182_ = _dafny.Array(False, 3)
                d_1146_v144_ = nw182_
                d_1147_v145_: _dafny.Map
                d_1147_v145_ = _dafny.Map({d_1146_v144_: (self).f9})
                d_1148_v146_: _dafny.Seq
                d_1148_v146_ = _dafny.SeqWithoutIsStrInference([d_1146_v144_, d_1146_v144_])
                d_1149_v147_: _dafny.Seq
                d_1149_v147_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                d_1150_v148_: _dafny.MultiSet
                d_1150_v148_ = _dafny.MultiSet([_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_1139_cf25_])).set(default__.safeIndex((d_1140_v141_)[default__.safeIndex(39, (d_1140_v141_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_1139_cf25_]))), d_1139_cf25_)), d_1105_v123_])
                d_1147_v145_ = (d_1147_v145_).set((d_1148_v146_)[default__.safeIndex(len(d_1149_v147_), len(d_1148_v146_))], (d_1105_v123_) not in (d_1150_v148_))
                r0 = default__.safeDivisionInt(p1, 84)
            elif source13_.is_DC14:
                d_1151___mcc_h23_ = source13_.cf26
                d_1152___mcc_h24_ = source13_.cf27
                d_1153___mcc_h25_ = source13_.cf28
                d_1154_cf28_ = d_1153___mcc_h25_
                d_1155_cf27_ = d_1152___mcc_h24_
                d_1156_cf26_ = d_1151___mcc_h23_
                d_1157_v149_: _dafny.Array
                nw183_ = _dafny.Array(int(0), 29)
                d_1157_v149_ = nw183_
                index175_ = default__.safeIndex(857, (d_1157_v149_).length(0))
                (d_1157_v149_)[index175_] = p1
                d_1158_v150_: _dafny.Seq
                d_1158_v150_ = _dafny.SeqWithoutIsStrInference([d_1105_v123_])
                d_1159_v151_: _dafny.Map
                d_1159_v151_ = _dafny.Map({d_1155_cf27_: (d_1158_v150_)[default__.safeIndex(p1, len(d_1158_v150_))]})
                d_1160_v152_: _dafny.Seq
                d_1160_v152_ = _dafny.SeqWithoutIsStrInference([p1, len(d_1107_v124_)])
                index176_ = default__.safeIndex(857, (d_1157_v149_).length(0))
                rhs205_ = (((d_1159_v151_)[not((self).f9)] if (not((self).f9)) in (d_1159_v151_) else _dafny.MultiSet(d_1160_v152_))) | (d_1105_v123_)
                rhs206_ = p1
                lhs134_ = d_1157_v149_
                lhs135_ = default__.safeIndex(857, (d_1157_v149_).length(0))
                d_1105_v123_ = rhs205_
                lhs134_[lhs135_] = rhs206_
                d_1161_v153_: _dafny.MultiSet
                d_1161_v153_ = _dafny.MultiSet(d_1160_v152_)
                d_1162_v154_: _dafny.MultiSet
                d_1162_v154_ = _dafny.MultiSet([d_1161_v153_])
                d_1154_cf28_ = ((d_1162_v154_)[d_1161_v153_] if (d_1161_v153_) in (d_1162_v154_) else p0)
                d_1163_v155_: _dafny.Seq
                d_1163_v155_ = _dafny.SeqWithoutIsStrInference([d_1107_v124_, d_1107_v124_, d_1107_v124_])
                d_1164_v156_: _dafny.Seq
                d_1164_v156_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifap"))
                d_1165_v157_: _dafny.Array
                nw184_ = _dafny.Array(None, 12)
                nw184_[int(0)] = d_1107_v124_
                nw184_[int(1)] = (d_1163_v155_)[default__.safeIndex(p0, len(d_1163_v155_))]
                nw184_[int(2)] = d_1107_v124_
                nw184_[int(3)] = default__.fm20((self).f9, d_1155_cf27_, (d_1157_v149_)[default__.safeIndex(857, (d_1157_v149_).length(0))], d_1164_v156_, globalState)
                nw184_[int(4)] = d_1107_v124_
                nw184_[int(5)] = (_dafny.SeqWithoutIsStrInference([d_1156_cf26_])) + (d_1107_v124_)
                nw184_[int(6)] = d_1107_v124_
                nw184_[int(7)] = d_1107_v124_
                nw184_[int(8)] = (d_1107_v124_) + (d_1107_v124_)
                nw184_[int(9)] = (d_1107_v124_).set(default__.safeIndex(d_1154_cf28_, len(d_1107_v124_)), True)
                nw184_[int(10)] = (d_1107_v124_) + (d_1107_v124_)
                nw184_[int(11)] = (d_1107_v124_).set(default__.safeIndex(p1, len(d_1107_v124_)), d_1155_cf27_)
                d_1165_v157_ = nw184_
                index177_ = default__.safeIndex(554, (d_1165_v157_).length(0))
                (d_1165_v157_)[index177_] = (d_1107_v124_) + (d_1107_v124_)
                d_1166_v158_: _dafny.Map
                d_1166_v158_ = _dafny.Map({d_1156_cf26_: len(default__.fm28(284, default__.fm1(d_1105_v123_, p0, globalState), globalState))})
                d_1167_v159_: D10
                d_1167_v159_ = D10_DC27(d_1166_v158_)
                index178_ = default__.safeIndex(554, (d_1165_v157_).length(0))
                (d_1165_v157_)[index178_] = ((d_1107_v124_).set(default__.safeIndex(d_1154_cf28_, len(d_1107_v124_)), p2)) + (default__.fm41(d_1167_v159_, (d_1157_v149_)[default__.safeIndex(857, (d_1157_v149_).length(0))], d_1160_v152_, globalState))
                d_1156_cf26_ = d_1155_cf27_
            elif source13_.is_DC11:
                d_1168___mcc_h26_ = source13_.cf19
                d_1169_cf19_ = d_1168___mcc_h26_
                d_1170_v160_: _dafny.Array
                def lambda63_(d_1171_p2_):
                    def lambda64_(d_1172_i14_):
                        return d_1171_p2_

                    return lambda64_

                init35_ = lambda63_(p2)
                nw185_ = _dafny.Array(None, 11)
                for i0_35_ in range(nw185_.length(0)):
                    nw185_[i0_35_] = init35_(i0_35_)
                d_1170_v160_ = nw185_
                d_1173_v161_: _dafny.Map
                d_1173_v161_ = _dafny.Map({p0: d_1170_v160_})
                d_1174_v162_: _dafny.Map
                d_1174_v162_ = _dafny.Map({(d_1173_v161_) | (d_1173_v161_): (self).f9})
                d_1175_v163_: D9
                d_1175_v163_ = D9_DC24(not(p2), p0, default__.fm28(p1, p2, globalState))
                d_1174_v162_ = (d_1174_v162_).set((_dafny.Map({p0: d_1170_v160_})) | (d_1173_v161_), default__.fm1(d_1105_v123_, (d_1175_v163_).cf42, globalState))
                d_1176_v164_: T1
                nw186_ = C2()
                nw186_.ctor__(p1, not(p2))
                d_1176_v164_ = nw186_
                d_1176_v164_ = (d_1176_v164_ if p2 else d_1176_v164_)
                d_1177_v165_: _dafny.Seq
                d_1177_v165_ = _dafny.SeqWithoutIsStrInference([d_1170_v160_])
                (globalState).f2 = ((self).f9 if (d_1100_v117_) != (d_1100_v117_) else (d_1177_v165_) < (d_1177_v165_))
                r0 = len(d_1107_v124_)
            elif True:
                d_1178___mcc_h27_ = source13_.cf29
                d_1179_cf29_ = d_1178___mcc_h27_
                r0 = p1
                d_1180_v166_: _dafny.Set
                d_1180_v166_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgsxcsdpe"))})
                d_1180_v166_ = d_1180_v166_
                d_1181_v167_: _dafny.Array
                def lambda65_(d_1182_p0_):
                    def lambda66_(d_1183_i15_):
                        return (d_1183_i15_) - (d_1182_p0_)

                    return lambda66_

                init36_ = lambda65_(p0)
                nw187_ = _dafny.Array(None, 2)
                for i0_36_ in range(nw187_.length(0)):
                    nw187_[i0_36_] = init36_(i0_36_)
                d_1181_v167_ = nw187_
                index179_ = default__.safeIndex(168, (d_1181_v167_).length(0))
                (d_1181_v167_)[index179_] = 688
                index180_ = default__.safeIndex(168, (d_1181_v167_).length(0))
                (d_1181_v167_)[index180_] = p1
                d_1184_v168_: _dafny.Seq
                d_1184_v168_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                r0 = len(d_1184_v168_)
            if (self).f9:
                d_1105_v123_ = d_1105_v123_
                rhs207_ = d_1100_v117_
                rhs208_ = (0) - ((0) - (p1))
                rhs209_ = (D14_DC38(d_1097_v115_, p2, d_1107_v124_)).cf67
                d_1100_v117_ = rhs207_
                r0 = rhs208_
                d_1097_v115_ = rhs209_
                d_1185_v169_: _dafny.Seq
                d_1185_v169_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqbsdfpq"))
                r3 = (d_1097_v115_) not in (d_1185_v169_)
                d_1186_v170_: _dafny.Set
                d_1186_v170_ = _dafny.Set({d_1098_v116_})
                (globalState).f2 = (True if not((self).f9) else (d_1186_v170_) == (d_1186_v170_))
                d_1187_v171_: _dafny.Array
                nw188_ = _dafny.Array(_dafny.Set({}), 10)
                d_1187_v171_ = nw188_
                d_1188_v172_: _dafny.Array
                nw189_ = _dafny.Array(None, 11)
                nw189_[int(0)] = p2
                nw189_[int(1)] = p2
                nw189_[int(2)] = True
                nw189_[int(3)] = default__.fm1(_dafny.MultiSet([p1]), p0, globalState)
                nw189_[int(4)] = p2
                nw189_[int(5)] = True
                nw189_[int(6)] = p2
                nw189_[int(7)] = default__.fm1(d_1105_v123_, p0, globalState)
                nw189_[int(8)] = (self).f9
                nw189_[int(9)] = False
                nw189_[int(10)] = default__.fm1(d_1105_v123_, (_dafny.MultiSet([(self).f9])).cardinality, globalState)
                d_1188_v172_ = nw189_
                index181_ = default__.safeIndex(610, (d_1187_v171_).length(0))
                (d_1187_v171_)[index181_] = _dafny.Set({d_1188_v172_})
                d_1189_v173_: _dafny.Set
                d_1189_v173_ = _dafny.Set({d_1188_v172_, d_1188_v172_, d_1188_v172_})
                index182_ = default__.safeIndex(610, (d_1187_v171_).length(0))
                (d_1187_v171_)[index182_] = d_1189_v173_
            elif True:
                (globalState).f2 = (self).f9
                d_1190_v174_: _dafny.Seq
                d_1190_v174_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ie"))
                d_1191_v175_: _dafny.Array
                def lambda67_(d_1192_i17_):
                    return _dafny.Map({(self).f9: (self).f9})

                init37_ = lambda67_
                nw190_ = _dafny.Array(None, 4)
                for i0_37_ in range(nw190_.length(0)):
                    nw190_[i0_37_] = init37_(i0_37_)
                d_1191_v175_ = nw190_
                d_1193_v176_: D7
                d_1193_v176_ = D7_DC19(d_1191_v175_, (self).f9, d_1190_v174_)
                d_1194_v177_: _dafny.Seq
                d_1194_v177_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1195_i16_ in range(default__.abs(153))]), d_1190_v174_, (d_1193_v176_).cf35])
                rhs210_ = d_1194_v177_
                rhs211_ = p0
                rhs212_ = ((self).f9) in (d_1107_v124_)
                rhs213_ = d_1097_v115_
                d_1194_v177_ = rhs210_
                r0 = rhs211_
                r3 = rhs212_
                d_1097_v115_ = rhs213_
                d_1196_v178_: _dafny.Array
                nw191_ = _dafny.Array(D3.default()(), 24)
                d_1196_v178_ = nw191_
                pat_let_tv65_ = d_1107_v124_
                pat_let_tv66_ = p1
                pat_let_tv67_ = d_1107_v124_
                pat_let_tv68_ = p2
                index183_ = default__.safeIndex(789, (d_1196_v178_).length(0))
                def iife121_(_pat_let48_0):
                    def iife122_(d_1197_dt__update__tmp_h10_):
                        def iife123_(_pat_let49_0):
                            def iife124_(d_1198_dt__update_hcf18_h0_):
                                return D3_DC9((d_1197_dt__update__tmp_h10_).cf17, d_1198_dt__update_hcf18_h0_)
                            return iife124_(_pat_let49_0)
                        return iife123_((pat_let_tv65_).set(default__.safeIndex(pat_let_tv66_, len(pat_let_tv67_)), pat_let_tv68_))
                    return iife122_(_pat_let48_0)
                (d_1196_v178_)[index183_] = iife121_(D3_DC9((self).f9, d_1107_v124_))
                d_1199_v179_: D3
                d_1199_v179_ = D3_DC9((657) > (default__.fm0(globalState)), _dafny.SeqWithoutIsStrInference([(self).f9, p2]))
                index184_ = default__.safeIndex(789, (d_1196_v178_).length(0))
                (d_1196_v178_)[index184_] = d_1199_v179_
                d_1200_v180_: _dafny.MultiSet
                d_1200_v180_ = _dafny.MultiSet([d_1097_v115_])
                def iife125_():
                    coll25_ = _dafny.Set()
                    compr_25_: str
                    for compr_25_ in ((d_1200_v180_) | (d_1200_v180_)).Elements:
                        d_1201_v181_: str = compr_25_
                        if (d_1201_v181_) in ((d_1200_v180_) | (d_1200_v180_)):
                            coll25_ = coll25_.union(_dafny.Set([d_1201_v181_]))
                    return _dafny.Set(coll25_)
                r0 = len(iife125_()
                )
                d_1202_v182_: _dafny.Array
                def lambda68_(d_1203_v115_, d_1204_v117_):
                    def lambda69_(d_1205_i18_):
                        return _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_1203_v115_ for d_1206_i19_ in range(default__.abs(846))])), (d_1204_v117_).cardinality})

                    return lambda69_

                init38_ = lambda68_(d_1097_v115_, d_1100_v117_)
                nw192_ = _dafny.Array(None, 12)
                for i0_38_ in range(nw192_.length(0)):
                    nw192_[i0_38_] = init38_(i0_38_)
                d_1202_v182_ = nw192_
                d_1207_v183_: _dafny.Seq
                d_1207_v183_ = _dafny.SeqWithoutIsStrInference([d_1202_v182_, d_1202_v182_])
                d_1208_v184_: _dafny.Seq
                d_1208_v184_ = _dafny.SeqWithoutIsStrInference([(d_1207_v183_)[default__.safeIndex(default__.fm0(globalState), len(d_1207_v183_))]])
                d_1209_v185_: C3
                nw193_ = C3()
                nw193_.ctor__(True, (d_1208_v184_)[default__.safeIndex(p0, len(d_1208_v184_))], p2)
                d_1209_v185_ = nw193_
                d_1210_v186_: _dafny.Map
                d_1210_v186_ = _dafny.Map({(self).f9: d_1209_v185_})
                d_1211_v187_: _dafny.Seq
                d_1211_v187_ = _dafny.SeqWithoutIsStrInference([p0, p1, (0) - (len(d_1210_v186_))])
                d_1211_v187_ = d_1211_v187_
        d_1212_v188_: _dafny.Seq
        d_1212_v188_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "canqxbdox"))
        d_1213_v189_: D1
        d_1213_v189_ = D1_DC2(p1, True, len(d_1212_v188_), p0)
        d_1214_v190_: C1
        nw194_ = C1()
        nw194_.ctor__(_dafny.SeqWithoutIsStrInference([d_1213_v189_]))
        d_1214_v190_ = nw194_
        d_1215_v191_: _dafny.Set
        d_1215_v191_ = _dafny.Set({d_1214_v190_, d_1214_v190_})
        hi3_ = (default__.fm0(globalState)) - (p0)
        for d_1216_i20_ in range(len(d_1215_v191_), hi3_):
            d_1217_v192_: C1
            nw195_ = C1()
            nw195_.ctor__((d_1214_v190_.f13) + (d_1214_v190_.f13))
            d_1217_v192_ = nw195_
            d_1218_v193_: _dafny.Seq
            d_1218_v193_ = _dafny.SeqWithoutIsStrInference([(self).f9, not ((self).f9) or (not((self).f9)), True])
            r0 = len(d_1218_v193_)
            d_1219_v194_: _dafny.Set
            d_1219_v194_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([p0])})
            d_1220_v195_: _dafny.MultiSet
            d_1220_v195_ = _dafny.MultiSet([p0, (0) - ((0) - ((d_1216_i20_) * (len(d_1219_v194_))))])
            d_1220_v195_ = d_1220_v195_
            r0 = len(_dafny.SeqWithoutIsStrInference([(d_1212_v188_)[default__.safeIndex(p0, len(d_1212_v188_))] for d_1221_i21_ in range(default__.abs(218))]))
        d_1222_v196_: str
        d_1222_v196_ = _dafny.CodePoint('b')
        d_1223_v197_: D7
        d_1223_v197_ = D7_DC20((self).f9, d_1222_v196_)
        d_1224_v198_: D7
        d_1224_v198_ = D7_DC21(d_1223_v197_)
        d_1225_v199_: D7
        d_1225_v199_ = D7_DC21(d_1223_v197_)
        d_1226_v200_: D7
        d_1226_v200_ = D7_DC21(d_1224_v198_)
        d_1227_v201_: _dafny.Seq
        d_1227_v201_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lo")))])
        rhs214_ = ((self).f9) == ((self).f9)
        rhs215_ = p2
        rhs216_ = d_1226_v200_
        rhs217_ = (p0) == ((p0 if (self).f9 else len(d_1227_v201_)))
        rhs218_ = default__.safeModuloInt(p0, p0)
        lhs136_ = globalState
        lhs137_ = globalState
        lhs138_ = globalState
        lhs136_.f2 = rhs214_
        lhs137_.f2 = rhs215_
        d_1226_v200_ = rhs216_
        lhs138_.f2 = rhs217_
        r0 = rhs218_
        d_1228_v202_: _dafny.Array
        nw196_ = _dafny.Array(D15.default()(), 19)
        d_1228_v202_ = nw196_
        d_1229_v203_: _dafny.Set
        d_1229_v203_ = _dafny.Set({p0, p1})
        d_1230_v204_: _dafny.Seq
        d_1230_v204_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1})])
        d_1231_v206_: _dafny.Array
        nw197_ = _dafny.Array(None, 15)
        nw197_[int(0)] = d_1229_v203_
        nw197_[int(1)] = d_1229_v203_
        nw197_[int(2)] = d_1229_v203_
        nw197_[int(3)] = d_1229_v203_
        nw197_[int(4)] = d_1229_v203_
        nw197_[int(5)] = d_1229_v203_
        nw197_[int(6)] = d_1229_v203_
        nw197_[int(7)] = (d_1230_v204_)[default__.safeIndex(default__.fm0(globalState), len(d_1230_v204_))]
        nw197_[int(8)] = _dafny.Set({p0, -854})
        nw197_[int(9)] = d_1229_v203_
        nw197_[int(10)] = d_1229_v203_
        nw197_[int(11)] = d_1229_v203_
        nw197_[int(12)] = d_1229_v203_
        nw197_[int(13)] = d_1229_v203_
        def iife126_():
            coll26_ = _dafny.Set()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(96, 923):
                d_1232_v205_: int = compr_26_
                if ((96) <= (d_1232_v205_)) and ((d_1232_v205_) < (923)):
                    coll26_ = coll26_.union(_dafny.Set([default__.safeDivisionInt(d_1232_v205_, p1)]))
            return _dafny.Set(coll26_)
        nw197_[int(14)] = iife126_()
        
        d_1231_v206_ = nw197_
        d_1233_v207_: T2
        nw198_ = C3()
        nw198_.ctor__((self).f9, d_1231_v206_, default__.fm1(default__.fm29(not(p2), False, globalState), len(d_1212_v188_), globalState))
        d_1233_v207_ = nw198_
        d_1234_v208_: D15
        d_1234_v208_ = D15_DC42((self).f9, p1, d_1233_v207_, len(d_1229_v203_))
        index185_ = default__.safeIndex(708, (d_1228_v202_).length(0))
        (d_1228_v202_)[index185_] = d_1234_v208_
        index186_ = default__.safeIndex(708, (d_1228_v202_).length(0))
        (d_1228_v202_)[index186_] = d_1234_v208_
        (globalState).f2 = p2
        d_1235_v209_: _dafny.MultiSet
        d_1235_v209_ = _dafny.MultiSet([p1])
        d_1236_v210_: _dafny.Map
        d_1236_v210_ = _dafny.Map({(self).f9: (d_1233_v207_).f9})
        d_1237_v211_: D1
        d_1237_v211_ = D1_DC4(p1, (d_1233_v207_).f9, (self).f9, p2, default__.fm1(d_1235_v209_, len(d_1236_v210_), globalState))
        r0 = (0) - (((d_1237_v211_).cf9) - (p1))
        d_1238_v212_: D4
        d_1238_v212_ = D4_DC13(p0)
        r1 = d_1238_v212_
        d_1239_v213_: _dafny.MultiSet
        d_1239_v213_ = _dafny.MultiSet([(d_1212_v188_) < ((d_1212_v188_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f9])).cardinality])), len(d_1212_v188_)), d_1222_v196_))])
        r2 = d_1239_v213_
        r3 = p2
        return r0, r1, r2, r3


class C5(T1, T2, T0):
    def  __init__(self):
        self._f9: bool = False
        self._f10: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f9(self):
        return self._f9
    def ctor__(self, f10, f9):
        (self)._f10 = f10
        (self)._f9 = f9

    def fm2(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f9, (self).f9]))]))

    def fm3(self, globalState):
        return (self).f9

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((self).f10) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10]): (0) - ((self).f10)})))

    def fm5(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([(0) - ((self).f10), (self).f10])).issubset(_dafny.MultiSet([(self).f10, 884]))

    def m1(self, p0, p1, p2, p3, globalState):
        d_1240_v0_: _dafny.Seq
        d_1240_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
        d_1241_v1_: _dafny.Seq
        d_1241_v1_ = _dafny.SeqWithoutIsStrInference([d_1240_v0_, d_1240_v0_])
        hi4_ = (_dafny.MultiSet((d_1241_v1_) + (_dafny.SeqWithoutIsStrInference([d_1240_v0_ for d_1242_i1_ in range(default__.abs(379))])))).cardinality
        for d_1243_i0_ in range((self).f10, hi4_):
            d_1244_v2_: str
            d_1244_v2_ = _dafny.CodePoint('x')
            rhs219_ = p2
            rhs220_ = _dafny.CodePoint('c')
            lhs139_ = globalState
            lhs139_.f2 = rhs219_
            d_1244_v2_ = rhs220_
            d_1245_v3_: int
            d_1245_v3_ = 504
            d_1246_v4_: _dafny.Array
            nw199_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_1246_v4_ = nw199_
            d_1247_v5_: _dafny.Array
            nw200_ = _dafny.Array(False, 29)
            d_1247_v5_ = nw200_
            index187_ = default__.safeIndex(129, (d_1246_v4_).length(0))
            (d_1246_v4_)[index187_] = d_1247_v5_
            index188_ = default__.safeIndex(129, (d_1246_v4_).length(0))
            rhs221_ = p2
            rhs222_ = p1
            rhs223_ = default__.safeDivisionInt((d_1243_i0_ if p2 else d_1243_i0_), len(d_1240_v0_))
            rhs224_ = d_1247_v5_
            rhs225_ = (self).f9
            lhs140_ = globalState
            lhs141_ = globalState
            lhs142_ = d_1246_v4_
            lhs143_ = default__.safeIndex(129, (d_1246_v4_).length(0))
            lhs144_ = globalState
            lhs140_.f2 = rhs221_
            lhs141_.f2 = rhs222_
            d_1245_v3_ = rhs223_
            lhs142_[lhs143_] = rhs224_
            lhs144_.f2 = rhs225_
            d_1248_v6_: _dafny.Seq
            d_1248_v6_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_1249_v7_: _dafny.Map
            d_1249_v7_ = _dafny.Map({p2: (d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]})
            index189_ = default__.safeIndex(129, (d_1246_v4_).length(0))
            rhs226_ = ((d_1249_v7_)[p2] if (p2) in (d_1249_v7_) else (d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))])
            rhs227_ = p0
            rhs228_ = d_1248_v6_
            lhs145_ = d_1246_v4_
            lhs146_ = default__.safeIndex(129, (d_1246_v4_).length(0))
            lhs145_[lhs146_] = rhs226_
            d_1245_v3_ = rhs227_
            d_1248_v6_ = rhs228_
            d_1250_v8_: _dafny.MultiSet
            d_1250_v8_ = _dafny.MultiSet([d_1243_i0_])
            d_1251_v9_: _dafny.MultiSet
            d_1251_v9_ = _dafny.MultiSet([p1, p1, p1])
            d_1252_v10_: _dafny.Map
            d_1252_v10_ = _dafny.Map({(self).f9: default__.fm1(d_1250_v8_, (d_1251_v9_).cardinality, globalState)})
            d_1253_v11_: _dafny.MultiSet
            d_1253_v11_ = _dafny.MultiSet([d_1252_v10_])
            if (_dafny.MultiSet([d_1252_v10_, d_1252_v10_])).ispropersubset(d_1253_v11_):
                d_1254_v12_: _dafny.Array
                def lambda70_(d_1255_i0_):
                    def lambda71_(d_1256_i2_):
                        return (d_1256_i2_) - (d_1255_i0_)

                    return lambda71_

                init39_ = lambda70_(d_1243_i0_)
                nw201_ = _dafny.Array(None, 18)
                for i0_39_ in range(nw201_.length(0)):
                    nw201_[i0_39_] = init39_(i0_39_)
                d_1254_v12_ = nw201_
                index190_ = default__.safeIndex(852, (d_1254_v12_).length(0))
                (d_1254_v12_)[index190_] = p0
                index191_ = default__.safeIndex(852, (d_1254_v12_).length(0))
                (d_1254_v12_)[index191_] = d_1243_i0_
                d_1245_v3_ = p0
                (globalState).f2 = not((True) == ((self).f9))
                d_1252_v10_ = (d_1252_v10_).set(p1, (self).f9)
                d_1252_v10_ = (d_1252_v10_).set(not(p2), p2)
            elif True:
                arr2_ = (d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]
                index192_ = default__.safeIndex(931, ((d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]).length(0))
                arr2_[index192_] = (self).f9
                d_1257_v13_: _dafny.Array
                def lambda72_(d_1258_p0_):
                    def lambda73_(d_1259_i3_):
                        return _dafny.Map({(self).f9: d_1258_p0_})

                    return lambda73_

                init40_ = lambda72_(p0)
                nw202_ = _dafny.Array(None, 9)
                for i0_40_ in range(nw202_.length(0)):
                    nw202_[i0_40_] = init40_(i0_40_)
                d_1257_v13_ = nw202_
                d_1260_v14_: _dafny.Map
                d_1260_v14_ = _dafny.Map({((default__.fm6(_dafny.SeqWithoutIsStrInference([p2]), globalState)) + (d_1240_v0_)).set(default__.safeIndex(p0, len((default__.fm6(_dafny.SeqWithoutIsStrInference([p2]), globalState)) + (d_1240_v0_))), d_1244_v2_): d_1257_v13_})
                arr3_ = (d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]
                index193_ = default__.safeIndex(931, ((d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]).length(0))
                rhs229_ = d_1244_v2_
                rhs230_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srhyeyrp"))) + (_dafny.SeqWithoutIsStrInference([d_1244_v2_ for d_1261_i4_ in range(default__.abs(574))]))) <= (d_1240_v0_)
                rhs231_ = ((d_1260_v14_)[(d_1240_v0_) + (d_1240_v0_)] if ((d_1240_v0_) + (d_1240_v0_)) in (d_1260_v14_) else d_1257_v13_)
                lhs147_ = (d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]
                lhs148_ = default__.safeIndex(931, ((d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]).length(0))
                d_1244_v2_ = rhs229_
                lhs147_[lhs148_] = rhs230_
                d_1257_v13_ = rhs231_
                d_1262_v15_: _dafny.Map
                d_1262_v15_ = _dafny.Map({(d_1248_v6_)[default__.safeIndex(p0, len(d_1248_v6_))]: d_1244_v2_})
                d_1262_v15_ = (d_1262_v15_).set((d_1250_v8_).ispropersubset(d_1250_v8_), d_1244_v2_)
                d_1263_v16_: _dafny.Map
                d_1263_v16_ = _dafny.Map({d_1252_v10_: d_1243_i0_})
                d_1264_v17_: _dafny.Seq
                d_1264_v17_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                d_1245_v3_ = (0) - ((((d_1263_v16_)[d_1252_v10_] if (d_1252_v10_) in (d_1263_v16_) else (self).f10) if (self).f9 else (_dafny.MultiSet((d_1264_v17_).set(default__.safeIndex(p0, len(d_1264_v17_)), p0))).cardinality))
                d_1265_v18_: D1
                d_1265_v18_ = D1_DC1(d_1245_v3_)
                d_1266_v19_: _dafny.Map
                d_1266_v19_ = _dafny.Map({d_1240_v0_: p2})
                d_1267_v20_: _dafny.Map
                d_1267_v20_ = _dafny.Map({len((d_1264_v17_) + ((self).fm2(len(_dafny.SeqWithoutIsStrInference([-758, -196])), p2, p1, (d_1265_v18_).cf1, globalState))): _dafny.Map({len(d_1266_v19_): default__.fm0(globalState)})})
                d_1268_v21_: _dafny.Set
                d_1268_v21_ = _dafny.Set({(d_1248_v6_)[default__.safeIndex(len(_dafny.Map({((d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))])[default__.safeIndex(931, ((d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]).length(0))]: p2})), len(d_1248_v6_))]})
                d_1269_v22_: _dafny.Array
                nw203_ = _dafny.Array(int(0), 19)
                d_1269_v22_ = nw203_
                d_1270_v23_: _dafny.Map
                d_1270_v23_ = _dafny.Map({d_1269_v22_: d_1243_i0_})
                d_1271_v24_: _dafny.Map
                d_1271_v24_ = _dafny.Map({d_1268_v21_: p1})
                d_1272_v25_: _dafny.MultiSet
                d_1272_v25_ = _dafny.MultiSet([d_1271_v24_])
                d_1273_v26_: _dafny.Array
                nw204_ = _dafny.Array(None, 27)
                nw204_[int(0)] = ((d_1270_v23_)[d_1269_v22_] if (d_1269_v22_) in (d_1270_v23_) else 245)
                nw204_[int(1)] = (d_1250_v8_).cardinality
                nw204_[int(2)] = p0
                nw204_[int(3)] = (d_1251_v9_).cardinality
                nw204_[int(4)] = default__.safeDivisionInt(d_1245_v3_, 457)
                nw204_[int(5)] = default__.fm0(globalState)
                nw204_[int(6)] = p0
                nw204_[int(7)] = ((self).f10) - (d_1243_i0_)
                nw204_[int(8)] = d_1243_i0_
                nw204_[int(9)] = (d_1243_i0_) + (len(d_1240_v0_))
                nw204_[int(10)] = (self).f10
                nw204_[int(11)] = ((d_1263_v16_)[d_1252_v10_] if (d_1252_v10_) in (d_1263_v16_) else len(d_1240_v0_))
                nw204_[int(12)] = (d_1243_i0_) - ((d_1250_v8_).cardinality)
                nw204_[int(13)] = len(d_1240_v0_)
                nw204_[int(14)] = (self).f10
                nw204_[int(15)] = (d_1251_v9_).cardinality
                nw204_[int(16)] = p0
                nw204_[int(17)] = (d_1245_v3_) * ((self).f10)
                nw204_[int(18)] = (d_1245_v3_) - (d_1245_v3_)
                nw204_[int(19)] = 845
                nw204_[int(20)] = default__.safeModuloInt(d_1245_v3_, d_1243_i0_)
                nw204_[int(21)] = (0) - (p0)
                nw204_[int(22)] = d_1245_v3_
                nw204_[int(23)] = (self).f10
                nw204_[int(24)] = (d_1243_i0_) + (p0)
                nw204_[int(25)] = (len(d_1240_v0_)) * ((0) - (((d_1272_v25_)[d_1271_v24_] if (d_1271_v24_) in (d_1272_v25_) else p0)))
                nw204_[int(26)] = (len(d_1240_v0_) if p1 else 678)
                d_1273_v26_ = nw204_
                d_1274_v27_: _dafny.Map
                d_1274_v27_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1244_v2_ for d_1275_i5_ in range(default__.abs(-168))])): ((d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))])[default__.safeIndex(931, ((d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]).length(0))]})
                index194_ = default__.safeIndex(106, (d_1273_v26_).length(0))
                (d_1273_v26_)[index194_] = len((d_1274_v27_ if p1 else d_1274_v27_))
                index195_ = default__.safeIndex(106, (d_1273_v26_).length(0))
                def iife127_():
                    coll27_ = _dafny.Map()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(644, 494):
                        d_1276_v28_: int = compr_27_
                        if ((644) <= (d_1276_v28_)) and ((d_1276_v28_) < (494)):
                            coll27_[default__.safeModuloInt(d_1276_v28_, p0)] = d_1245_v3_
                    return _dafny.Map(coll27_)
                rhs232_ = (self).f10
                rhs233_ = 854
                rhs234_ = ((d_1267_v20_) | (d_1267_v20_)) | (_dafny.Map({len(d_1240_v0_): iife127_()
                }))
                rhs235_ = d_1268_v21_
                rhs236_ = d_1243_i0_
                lhs149_ = d_1273_v26_
                lhs150_ = default__.safeIndex(106, (d_1273_v26_).length(0))
                d_1245_v3_ = rhs232_
                d_1245_v3_ = rhs233_
                d_1267_v20_ = rhs234_
                d_1268_v21_ = rhs235_
                lhs149_[lhs150_] = rhs236_
                arr4_ = (d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]
                index196_ = default__.safeIndex(931, ((d_1246_v4_)[default__.safeIndex(129, (d_1246_v4_).length(0))]).length(0))
                arr4_[index196_] = (self).fm3(globalState)
        d_1277_v29_: _dafny.MultiSet
        d_1277_v29_ = _dafny.MultiSet([p0])
        rhs237_ = ((d_1277_v29_ if (self).f9 else d_1277_v29_)).isdisjoint(d_1277_v29_)
        rhs238_ = d_1240_v0_
        lhs151_ = globalState
        lhs151_.f2 = rhs237_
        d_1240_v0_ = rhs238_
        d_1278_v30_: D1
        d_1278_v30_ = D1_DC2((self).f10, (self).f9, p0, (D1_DC4(347, (self).f9, p1, False, p2)).cf9)
        d_1279_v31_: _dafny.Seq
        d_1279_v31_ = _dafny.SeqWithoutIsStrInference([p0, (d_1278_v30_).cf2])
        (globalState).f2 = ((len(d_1279_v31_)) * (p0)) > (p0)
        hi5_ = (self).f10
        for d_1280_i6_ in range(p0, hi5_):
            d_1281_v32_: bool
            d_1282_v33_: int
            d_1283_v34_: int
            out14_: bool
            out15_: int
            out16_: int
            out14_, out15_, out16_ = (self).m5(globalState)
            d_1281_v32_ = out14_
            d_1282_v33_ = out15_
            d_1283_v34_ = out16_
            d_1284_v35_: _dafny.Map
            d_1284_v35_ = _dafny.Map({d_1281_v32_: p1})
            d_1285_v36_: _dafny.Seq
            d_1285_v36_ = _dafny.SeqWithoutIsStrInference([(self).f9, not(d_1281_v32_)])
            d_1286_v37_: _dafny.Map
            d_1286_v37_ = _dafny.Map({d_1284_v35_: d_1285_v36_})
            d_1287_v38_: _dafny.Set
            d_1287_v38_ = _dafny.Set({d_1279_v31_})
            d_1288_v39_: _dafny.Map
            d_1288_v39_ = _dafny.Map({d_1283_v34_: d_1285_v36_})
            d_1282_v33_ = len(default__.fm7(len(d_1286_v37_), p2, D1_DC4(d_1282_v33_, p1, not((self).fm5(d_1287_v38_, d_1288_v39_, len(d_1240_v0_), globalState)), not(d_1281_v32_), False), (len(d_1285_v36_)) > (d_1280_i6_), globalState))
            d_1289_v40_: _dafny.Map
            d_1289_v40_ = _dafny.Map({len((self).fm2(len((d_1279_v31_).set(default__.safeIndex(d_1283_v34_, len(d_1279_v31_)), 263)), (self).f9, False, (self).f10, globalState)): D1_DC1(794)})
            d_1290_v41_: D1
            d_1290_v41_ = D1_DC1(d_1283_v34_)
            d_1289_v40_ = _dafny.Map({default__.safeDivisionInt(d_1280_i6_, d_1283_v34_): d_1290_v41_})
            d_1291_v42_: C0
            nw205_ = C0()
            nw205_.ctor__(d_1287_v38_, d_1282_v33_)
            d_1291_v42_ = nw205_
        d_1292_v43_: D1
        d_1292_v43_ = D1_DC4(p0, (self).f9, True, (self).f9, (self).f9)
        d_1293_v44_: _dafny.Seq
        d_1293_v44_ = _dafny.SeqWithoutIsStrInference([d_1292_v43_, d_1292_v43_, D1_DC4((self).f10, not(p1), not(p1), p2, p2), d_1292_v43_, d_1292_v43_])
        d_1294_v45_: _dafny.Seq
        d_1294_v45_ = _dafny.SeqWithoutIsStrInference([d_1293_v44_, d_1293_v44_, d_1293_v44_])
        d_1293_v44_ = ((d_1294_v45_)[default__.safeIndex((self).f10, len(d_1294_v45_))]) + ((d_1294_v45_)[default__.safeIndex((self).f10, len(d_1294_v45_))])
        d_1295_v46_: int
        d_1295_v46_ = 38
        d_1296_v47_: _dafny.Seq
        d_1296_v47_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1295_v46_: (self).f9})) for d_1297_i7_ in range(default__.abs(884))])
        d_1298_v48_: _dafny.Set
        d_1298_v48_ = _dafny.Set({d_1279_v31_, d_1279_v31_})
        d_1299_v49_: C0
        nw206_ = C0()
        nw206_.ctor__(d_1298_v48_, d_1295_v46_)
        d_1299_v49_ = nw206_
        d_1300_v50_: _dafny.Map
        d_1300_v50_ = _dafny.Map({(self).f9: d_1299_v49_})
        rhs239_ = (p0) != ((0) - (d_1295_v46_))
        rhs240_ = default__.safeModuloInt(len(d_1279_v31_), ((self).f10) * (len(d_1300_v50_)))
        rhs241_ = default__.fm9(default__.safeDivisionInt((self).f10, p0), d_1295_v46_, False, globalState)
        rhs242_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuoebdhit"))
        lhs152_ = globalState
        lhs152_.f2 = rhs239_
        d_1295_v46_ = rhs240_
        d_1296_v47_ = rhs241_
        d_1240_v0_ = rhs242_

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: _dafny.Set = _dafny.Set({})
        d_1301_i0_: int
        d_1301_i0_ = 0
        with _dafny.label("8"):
            while (self).f9:
                with _dafny.c_label("8"):
                    if (d_1301_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_1301_i0_ = (d_1301_i0_) + (1)
                    d_1302_v0_: _dafny.Set
                    d_1302_v0_ = _dafny.Set({(p3).set(default__.safeIndex((self).f10, len(p3)), (self).f10)})
                    d_1303_v1_: C0
                    nw207_ = C0()
                    nw207_.ctor__(d_1302_v0_, ((p3)[default__.safeIndex((self).f10, len(p3))]) - ((self).f10))
                    d_1303_v1_ = nw207_
                    d_1304_v2_: _dafny.Array
                    nw208_ = _dafny.Array(int(0), 15)
                    d_1304_v2_ = nw208_
                    index197_ = default__.safeIndex(77, (d_1304_v2_).length(0))
                    (d_1304_v2_)[index197_] = default__.safeDivisionInt(d_1303_v1_.f12, d_1303_v1_.f12)
                    index198_ = default__.safeIndex(77, (d_1304_v2_).length(0))
                    (d_1304_v2_)[index198_] = d_1303_v1_.f12
                    r1 = (self).f9
                    index199_ = default__.safeIndex(77, (d_1304_v2_).length(0))
                    (d_1304_v2_)[index199_] = default__.safeDivisionInt((d_1304_v2_)[default__.safeIndex(77, (d_1304_v2_).length(0))], (d_1304_v2_)[default__.safeIndex(77, (d_1304_v2_).length(0))])
                    pass
            pass
        d_1305_v3_: _dafny.Seq
        d_1305_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgopc"))
        d_1306_v4_: _dafny.Set
        d_1306_v4_ = _dafny.Set({(self).f10})
        d_1307_v5_: _dafny.Map
        d_1307_v5_ = _dafny.Map({(self).f9: (d_1306_v4_).isdisjoint(d_1306_v4_)})
        d_1308_v6_: _dafny.Array
        nw209_ = _dafny.Array(int(0), 19)
        d_1308_v6_ = nw209_
        index200_ = default__.safeIndex(716, (d_1308_v6_).length(0))
        (d_1308_v6_)[index200_] = ((self).f10) - ((self).f10)
        index201_ = default__.safeIndex(716, (d_1308_v6_).length(0))
        rhs243_ = p1
        rhs244_ = d_1307_v5_
        rhs245_ = (default__.fm0(globalState)) in (_dafny.Set({(self).f10, 280, (self).f10, (self).f10}))
        rhs246_ = (self).f10
        rhs247_ = not ((self).f9) or (((self).f9) and (False))
        lhs153_ = d_1308_v6_
        lhs154_ = default__.safeIndex(716, (d_1308_v6_).length(0))
        d_1305_v3_ = rhs243_
        d_1307_v5_ = rhs244_
        r1 = rhs245_
        lhs153_[lhs154_] = rhs246_
        r1 = rhs247_
        if (self).f9:
            d_1309_v7_: _dafny.Map
            d_1309_v7_ = _dafny.Map({(self).f10: True})
            d_1310_v8_: D1
            d_1310_v8_ = D1_DC3((default__.fm0(globalState)) not in (d_1309_v7_), p1, ((self).f10) * ((self).f10))
            d_1311_v9_: _dafny.Set
            d_1311_v9_ = _dafny.Set({p2, p2, p2, (p2 if (self).f9 else p2)})
            d_1312_v10_: _dafny.MultiSet
            d_1312_v10_ = _dafny.MultiSet([(d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))]])
            d_1313_v11_: _dafny.Array
            nw210_ = _dafny.Array(None, 21)
            nw210_[int(0)] = (len(default__.fm10((self).f9, globalState))) <= ((self).f10)
            nw210_[int(1)] = not((self).f9)
            nw210_[int(2)] = (self).f9
            nw210_[int(3)] = (self).f9
            nw210_[int(4)] = (self).f9
            nw210_[int(5)] = (d_1310_v8_).cf6
            nw210_[int(6)] = (self).f9
            nw210_[int(7)] = (self).f9
            nw210_[int(8)] = (self).f9
            nw210_[int(9)] = (self).f9
            nw210_[int(10)] = (p2) not in ((d_1305_v3_).set(default__.safeIndex((self).f10, len(d_1305_v3_)), default__.fm11((self).f9, globalState)))
            nw210_[int(11)] = (self).f9
            nw210_[int(12)] = not((self).f9)
            nw210_[int(13)] = ((D1_DC1((self).f10)).cf1) >= ((self).f10)
            nw210_[int(14)] = (self).f9
            nw210_[int(15)] = (self).f9
            nw210_[int(16)] = (self).f9
            nw210_[int(17)] = default__.fm1(_dafny.MultiSet(p3), (d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], globalState)
            nw210_[int(18)] = (d_1312_v10_).ispropersubset(d_1312_v10_)
            nw210_[int(19)] = (self).f9
            nw210_[int(20)] = (self).f9
            d_1313_v11_ = nw210_
            index202_ = default__.safeIndex(341, (d_1313_v11_).length(0))
            (d_1313_v11_)[index202_] = (default__.fm12((d_1310_v8_).cf6, len(d_1305_v3_), (self).f9, len(d_1305_v3_), globalState)).isdisjoint(default__.fm12(True, (self).f10, (self).f9, (d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], globalState))
            d_1314_v12_: _dafny.Seq
            d_1314_v12_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_1315_v13_: D1
            d_1315_v13_ = D1_DC4(144, (self).f9, (self).f9, (self).f9, (d_1314_v12_)[default__.safeIndex((self).f10, len(d_1314_v12_))])
            pat_let_tv69_ = d_1308_v6_
            pat_let_tv70_ = d_1308_v6_
            pat_let_tv71_ = d_1305_v3_
            index203_ = default__.safeIndex(341, (d_1313_v11_).length(0))
            def iife128_(_pat_let50_0):
                def iife129_(d_1316_dt__update__tmp_h0_):
                    def iife130_(_pat_let51_0):
                        def iife131_(d_1318_dt__update_hcf6_h0_):
                            return D1_DC3(d_1318_dt__update_hcf6_h0_, (d_1316_dt__update__tmp_h0_).cf7, (d_1316_dt__update__tmp_h0_).cf8)
                        return iife131_(_pat_let51_0)
                    return iife130_((_dafny.SeqWithoutIsStrInference([(0) - ((pat_let_tv70_)[default__.safeIndex(716, (pat_let_tv69_).length(0))]) for d_1317_i1_ in range(default__.abs(68))])) < (_dafny.SeqWithoutIsStrInference([len(pat_let_tv71_)])))
                return iife129_(_pat_let50_0)
            rhs248_ = iife128_(d_1310_v8_)
            rhs249_ = d_1311_v9_
            rhs250_ = (self).f9
            rhs251_ = (0) - ((d_1315_v13_).cf9)
            rhs252_ = (self).f10
            lhs155_ = d_1313_v11_
            lhs156_ = default__.safeIndex(341, (d_1313_v11_).length(0))
            d_1310_v8_ = rhs248_
            d_1311_v9_ = rhs249_
            lhs155_[lhs156_] = rhs250_
            r2 = rhs251_
            r2 = rhs252_
            if ((d_1313_v11_)[default__.safeIndex(341, (d_1313_v11_).length(0))]) or ((self).f9):
                d_1319_v14_: _dafny.Set
                d_1319_v14_ = _dafny.Set({p3})
                d_1320_v15_: _dafny.Map
                d_1320_v15_ = _dafny.Map({(d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))]: d_1314_v12_})
                (globalState).f2 = (self).fm5(d_1319_v14_, d_1320_v15_, ((self).f10) * ((self).f10), globalState)
                r0 = (d_1313_v11_)[default__.safeIndex(341, (d_1313_v11_).length(0))]
                d_1321_v16_: _dafny.Map
                d_1321_v16_ = _dafny.Map({(self).f10: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjurujc"))})
                d_1305_v3_ = (((d_1321_v16_)[(0) - ((d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))])] if ((0) - ((d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))])) in (d_1321_v16_) else p1)) + (p1)
                d_1322_v17_: _dafny.MultiSet
                d_1322_v17_ = _dafny.MultiSet([(p1) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfyqtatu")))])
                d_1322_v17_ = d_1322_v17_
                d_1312_v10_ = d_1312_v10_
            elif True:
                index204_ = default__.safeIndex(716, (d_1308_v6_).length(0))
                rhs253_ = default__.safeModuloInt((self).f10, (self).f10)
                rhs254_ = (self).f10
                rhs255_ = not((self).f9)
                rhs256_ = (self).f10
                lhs157_ = globalState
                lhs158_ = d_1308_v6_
                lhs159_ = default__.safeIndex(716, (d_1308_v6_).length(0))
                r2 = rhs253_
                r2 = rhs254_
                lhs157_.f2 = rhs255_
                lhs158_[lhs159_] = rhs256_
                d_1323_v18_: _dafny.Map
                d_1323_v18_ = _dafny.Map({(self).f10: p2})
                d_1324_v20_: _dafny.Seq
                def iife132_():
                    coll28_ = _dafny.Set()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(545, 560):
                        d_1325_v19_: int = compr_28_
                        if ((545) <= (d_1325_v19_)) and ((d_1325_v19_) < (560)):
                            coll28_ = coll28_.union(_dafny.Set([(d_1325_v19_) + ((self).f10)]))
                    return _dafny.Set(coll28_)
                d_1324_v20_ = _dafny.SeqWithoutIsStrInference([iife132_()
                , (d_1306_v4_), d_1306_v4_])
                d_1326_v21_: _dafny.Map
                d_1326_v21_ = _dafny.Map({p2: p2})
                d_1327_v22_: _dafny.Array
                nw211_ = _dafny.Array(None, 8)
                nw211_[int(0)] = p2
                nw211_[int(1)] = p2
                nw211_[int(2)] = p2
                nw211_[int(3)] = ((d_1323_v18_)[len(d_1324_v20_)] if (len(d_1324_v20_)) in (d_1323_v18_) else p2)
                nw211_[int(4)] = p2
                nw211_[int(5)] = p2
                nw211_[int(6)] = p2
                nw211_[int(7)] = ((d_1326_v21_)[p2] if (p2) in (d_1326_v21_) else p2)
                d_1327_v22_ = nw211_
                index205_ = default__.safeIndex(675, (d_1327_v22_).length(0))
                (d_1327_v22_)[index205_] = p2
                d_1328_v23_: _dafny.Seq
                d_1328_v23_ = _dafny.SeqWithoutIsStrInference([len((p1) + (p1)), ((self).f10) * ((d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))])])
                d_1329_v24_: C0
                nw212_ = C0()
                nw212_.ctor__(_dafny.Set({d_1328_v23_}), (self).f10)
                d_1329_v24_ = nw212_
                d_1330_v25_: _dafny.Seq
                d_1330_v25_ = _dafny.SeqWithoutIsStrInference([d_1329_v24_])
                index206_ = default__.safeIndex(675, (d_1327_v22_).length(0))
                rhs257_ = p2
                rhs258_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lg"))
                rhs259_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lf"))
                rhs260_ = d_1328_v23_
                rhs261_ = d_1330_v25_
                lhs160_ = d_1327_v22_
                lhs161_ = default__.safeIndex(675, (d_1327_v22_).length(0))
                lhs160_[lhs161_] = rhs257_
                d_1305_v3_ = rhs258_
                d_1305_v3_ = rhs259_
                d_1328_v23_ = rhs260_
                d_1330_v25_ = rhs261_
                (globalState).f2 = ((self).f10) >= (len(p3))
                d_1331_v26_: _dafny.Array
                def lambda74_(d_1332_v6_, d_1333_v10_):
                    def lambda75_(d_1334_i2_):
                        return _dafny.Map({(d_1332_v6_)[default__.safeIndex(716, (d_1332_v6_).length(0))]: d_1333_v10_})

                    return lambda75_

                init41_ = lambda74_(d_1308_v6_, d_1312_v10_)
                nw213_ = _dafny.Array(None, 8)
                for i0_41_ in range(nw213_.length(0)):
                    nw213_[i0_41_] = init41_(i0_41_)
                d_1331_v26_ = nw213_
                d_1335_v27_: _dafny.Map
                d_1335_v27_ = _dafny.Map({len(d_1306_v4_): _dafny.MultiSet(p3)})
                index207_ = default__.safeIndex(131, (d_1331_v26_).length(0))
                (d_1331_v26_)[index207_] = (_dafny.Map({(d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))]: d_1312_v10_})) | (d_1335_v27_)
                d_1336_v28_: D1
                d_1336_v28_ = D1_DC2((self).f10, False, (self).f10, (0) - ((self).f10))
                d_1337_v29_: _dafny.Seq
                d_1337_v29_ = _dafny.SeqWithoutIsStrInference([d_1336_v28_, d_1336_v28_])
                d_1338_v30_: T0
                nw214_ = C1()
                nw214_.ctor__(d_1337_v29_)
                d_1338_v30_ = nw214_
                d_1339_v31_: _dafny.Map
                d_1339_v31_ = _dafny.Map({True: d_1338_v30_})
                d_1340_v32_: D1
                d_1340_v32_ = D1_DC2((d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], (self).f9, len(d_1339_v31_), (self).f10)
                index208_ = default__.safeIndex(131, (d_1331_v26_).length(0))
                (d_1331_v26_)[index208_] = (d_1335_v27_) | ((default__.fm13((self).fm4((d_1313_v11_)[default__.safeIndex(341, (d_1313_v11_).length(0))], (self).f9, (self).f9, (d_1340_v32_).cf3, globalState), (d_1313_v11_)[default__.safeIndex(341, (d_1313_v11_).length(0))], (self).f9, globalState)) | (d_1335_v27_))
                d_1341_v33_: C0
                nw215_ = C0()
                nw215_.ctor__(d_1329_v24_.f11, len(_dafny.SeqWithoutIsStrInference([p2 for d_1342_i3_ in range(default__.abs(839))])))
                d_1341_v33_ = nw215_
            d_1343_v34_: _dafny.Array
            nw216_ = _dafny.Array(_dafny.CodePoint('D'), 2)
            d_1343_v34_ = nw216_
            d_1344_v35_: _dafny.Seq
            d_1344_v35_ = _dafny.SeqWithoutIsStrInference([d_1343_v34_])
            index209_ = default__.safeIndex(341, (d_1313_v11_).length(0))
            (d_1313_v11_)[index209_] = (d_1343_v34_) in (d_1344_v35_)
            default__.m0((self).f9, globalState)
            d_1345_v36_: D1
            d_1345_v36_ = D1_DC2((self).f10, (self).f9, (d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], (self).f10)
            d_1346_v37_: _dafny.Seq
            d_1346_v37_ = _dafny.SeqWithoutIsStrInference([d_1345_v36_, d_1345_v36_])
            d_1347_v38_: C1
            nw217_ = C1()
            nw217_.ctor__(d_1346_v37_)
            d_1347_v38_ = nw217_
        elif True:
            d_1348_v39_: _dafny.Array
            nw218_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_1348_v39_ = nw218_
            index210_ = default__.safeIndex(191, (d_1348_v39_).length(0))
            (d_1348_v39_)[index210_] = d_1305_v3_
            index211_ = default__.safeIndex(191, (d_1348_v39_).length(0))
            (d_1348_v39_)[index211_] = d_1305_v3_
            d_1349_v40_: _dafny.Array
            nw219_ = _dafny.Array(None, 21)
            d_1349_v40_ = nw219_
            d_1350_v41_: D4
            d_1350_v41_ = D4_DC11(d_1349_v40_)
            d_1351_v42_: _dafny.Array
            nw220_ = _dafny.Array(None, 19)
            nw220_[int(0)] = d_1349_v40_
            nw220_[int(1)] = d_1349_v40_
            nw220_[int(2)] = d_1349_v40_
            nw220_[int(3)] = d_1349_v40_
            nw220_[int(4)] = d_1349_v40_
            nw220_[int(5)] = d_1349_v40_
            nw220_[int(6)] = d_1349_v40_
            nw220_[int(7)] = d_1349_v40_
            nw220_[int(8)] = d_1349_v40_
            nw220_[int(9)] = d_1349_v40_
            nw220_[int(10)] = d_1349_v40_
            nw220_[int(11)] = d_1349_v40_
            nw220_[int(12)] = d_1349_v40_
            nw220_[int(13)] = d_1349_v40_
            nw220_[int(14)] = d_1349_v40_
            nw220_[int(15)] = d_1349_v40_
            nw220_[int(16)] = d_1349_v40_
            nw220_[int(17)] = (d_1350_v41_).cf19
            nw220_[int(18)] = d_1349_v40_
            d_1351_v42_ = nw220_
            d_1352_v43_: _dafny.Array
            d_1352_v43_ = d_1351_v42_
            d_1351_v42_ = (d_1352_v43_)
            index212_ = default__.safeIndex(716, (d_1308_v6_).length(0))
            (d_1308_v6_)[index212_] = (self).f10
            d_1353_v44_: D1
            d_1353_v44_ = D1_DC2((self).f10, (self).f9, (d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], 247)
            d_1354_v45_: C1
            nw221_ = C1()
            nw221_.ctor__(_dafny.SeqWithoutIsStrInference([d_1353_v44_, d_1353_v44_]))
            d_1354_v45_ = nw221_
            d_1355_v46_: _dafny.Map
            d_1355_v46_ = _dafny.Map({(self).f9: (0) - ((d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))])})
            d_1355_v46_ = (d_1355_v46_) | (d_1355_v46_)
        r2 = ((d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))]) - ((d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))])
        d_1356_v47_: _dafny.Array
        nw222_ = _dafny.Array(False, 23)
        d_1356_v47_ = nw222_
        index213_ = default__.safeIndex(305, (d_1356_v47_).length(0))
        (d_1356_v47_)[index213_] = (self).f9
        d_1357_v49_: _dafny.Seq
        d_1357_v49_ = _dafny.SeqWithoutIsStrInference([d_1356_v47_])
        d_1358_v50_: _dafny.MultiSet
        d_1358_v50_ = _dafny.MultiSet([len(p1)])
        index214_ = default__.safeIndex(305, (d_1356_v47_).length(0))
        def iife133_():
            coll29_ = _dafny.Set()
            compr_29_: int
            for compr_29_ in (d_1306_v4_).Elements:
                d_1359_v48_: int = compr_29_
                if (d_1359_v48_) in (d_1306_v4_):
                    coll29_ = coll29_.union(_dafny.Set([(d_1359_v48_) * (78)]))
            return _dafny.Set(coll29_)
        rhs262_ = not((default__.fm12((self).f9, (d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], (self).f9, (self).f10, globalState)).issubset(iife133_()
        ))
        rhs263_ = d_1356_v47_
        rhs264_ = (d_1357_v49_)[default__.safeIndex((self).f10, len(d_1357_v49_))]
        rhs265_ = default__.fm1(d_1358_v50_, (d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], globalState)
        lhs162_ = d_1356_v47_
        lhs163_ = default__.safeIndex(305, (d_1356_v47_).length(0))
        lhs162_[lhs163_] = rhs262_
        d_1356_v47_ = rhs263_
        d_1356_v47_ = rhs264_
        r1 = rhs265_
        d_1360_v51_: D3
        d_1360_v51_ = D3_DC10()
        d_1360_v51_ = (d_1360_v51_ if (self).f9 else d_1360_v51_)
        r0 = (d_1356_v47_)[default__.safeIndex(305, (d_1356_v47_).length(0))]
        d_1361_v52_: _dafny.Seq
        d_1361_v52_ = _dafny.SeqWithoutIsStrInference([p3])
        d_1362_v54_: _dafny.Seq
        d_1362_v54_ = _dafny.SeqWithoutIsStrInference([(self).f9])
        d_1363_v55_: D3
        d_1363_v55_ = D3_DC9((self).f9, d_1362_v54_)
        d_1364_v56_: _dafny.Map
        d_1364_v56_ = _dafny.Map({(d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))]: _dafny.SeqWithoutIsStrInference([len((d_1363_v55_).cf18), (d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], (self).f10])})
        d_1365_v57_: _dafny.Set
        d_1365_v57_ = _dafny.Set({((d_1364_v56_)[(self).f10] if ((self).f10) in (d_1364_v56_) else p3)})
        d_1366_v58_: _dafny.Map
        d_1366_v58_ = _dafny.Map({(self).f10: d_1362_v54_})
        def iife134_():
            coll30_ = _dafny.Set()
            compr_30_: _dafny.Seq
            for compr_30_ in (d_1361_v52_).Elements:
                d_1367_v53_: _dafny.Seq = compr_30_
                if (d_1367_v53_) in (d_1361_v52_):
                    coll30_ = coll30_.union(_dafny.Set([d_1367_v53_]))
            return _dafny.Set(coll30_)
        r1 = (self).fm5((iife134_()
        ) | (d_1365_v57_), (d_1366_v58_) | (d_1366_v58_), (d_1308_v6_)[default__.safeIndex(716, (d_1308_v6_).length(0))], globalState)
        r2 = (self).f10
        r3 = (d_1306_v4_).intersection(d_1306_v4_)
        return r0, r1, r2, r3

    def m3(self, p0, globalState):
        r0: bool = False
        hi6_ = (self).f10
        for d_1368_i0_ in range((self).f10, hi6_):
            if p0:
                d_1369_v0_: _dafny.Set
                d_1369_v0_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_1368_i0_ for d_1370_i1_ in range(default__.abs(-810))])})
                d_1371_v1_: C0
                nw223_ = C0()
                nw223_.ctor__(d_1369_v0_, ((self).f10) + (d_1368_i0_))
                d_1371_v1_ = nw223_
                d_1372_v2_: _dafny.MultiSet
                d_1372_v2_ = _dafny.MultiSet([(self).f10, (self).f10])
                d_1373_v3_: D1
                d_1373_v3_ = D1_DC4((0) - ((0) - (((self).f10) * (d_1371_v1_.f12))), not((d_1372_v2_).ispropersubset(d_1372_v2_)), (self).f9, (self).f9, not(p0))
                d_1373_v3_ = D1_DC4(d_1368_i0_, (self).f9, p0, (self).f9, p0)
                d_1374_v4_: _dafny.Seq
                d_1374_v4_ = _dafny.SeqWithoutIsStrInference([(self).f10, 267])
                d_1375_v5_: _dafny.Seq
                d_1375_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdwey"))
                d_1376_v6_: D1
                d_1376_v6_ = D1_DC1(len(d_1375_v5_))
                pat_let_tv72_ = globalState
                d_1377_v7_: _dafny.Array
                nw224_ = _dafny.Array(None, 27)
                nw224_[int(0)] = default__.fm18(len(d_1374_v4_), p0, p0, globalState)
                nw224_[int(1)] = d_1376_v6_
                nw224_[int(2)] = d_1376_v6_
                nw224_[int(3)] = default__.fm18(d_1368_i0_, p0, p0, globalState)
                nw224_[int(4)] = (d_1376_v6_ if p0 else d_1376_v6_)
                nw224_[int(5)] = D1_DC1((0) - (d_1368_i0_))
                nw224_[int(6)] = d_1376_v6_
                nw224_[int(7)] = d_1376_v6_
                nw224_[int(8)] = d_1376_v6_
                nw224_[int(9)] = d_1376_v6_
                nw224_[int(10)] = d_1376_v6_
                def iife135_(_pat_let52_0):
                    def iife136_(d_1378_dt__update__tmp_h0_):
                        def iife137_(_pat_let53_0):
                            def iife138_(d_1379_dt__update_hcf1_h0_):
                                return D1_DC1(d_1379_dt__update_hcf1_h0_)
                            return iife138_(_pat_let53_0)
                        return iife137_(d_1368_i0_)
                    return iife136_(_pat_let52_0)
                nw224_[int(11)] = iife135_(d_1376_v6_)
                nw224_[int(12)] = d_1376_v6_
                nw224_[int(13)] = D1_DC1(d_1368_i0_)
                nw224_[int(14)] = (d_1376_v6_ if p0 else d_1376_v6_)
                nw224_[int(15)] = d_1376_v6_
                nw224_[int(16)] = d_1376_v6_
                nw224_[int(17)] = d_1376_v6_
                nw224_[int(18)] = d_1376_v6_
                nw224_[int(19)] = d_1376_v6_
                nw224_[int(20)] = d_1376_v6_
                nw224_[int(21)] = (d_1376_v6_ if (self).f9 else d_1376_v6_)
                nw224_[int(22)] = d_1376_v6_
                nw224_[int(23)] = d_1376_v6_
                def iife139_(_pat_let54_0):
                    def iife140_(d_1380_dt__update__tmp_h1_):
                        def iife141_(_pat_let55_0):
                            def iife142_(d_1381_dt__update_hcf1_h1_):
                                return D1_DC1(d_1381_dt__update_hcf1_h1_)
                            return iife142_(_pat_let55_0)
                        return iife141_(default__.fm0(pat_let_tv72_))
                    return iife140_(_pat_let54_0)
                nw224_[int(24)] = iife139_(d_1376_v6_)
                nw224_[int(25)] = d_1376_v6_
                nw224_[int(26)] = d_1376_v6_
                d_1377_v7_ = nw224_
                index215_ = default__.safeIndex(74, (d_1377_v7_).length(0))
                (d_1377_v7_)[index215_] = d_1376_v6_
                pat_let_tv73_ = d_1371_v1_
                pat_let_tv74_ = d_1371_v1_
                index216_ = default__.safeIndex(74, (d_1377_v7_).length(0))
                def iife143_(_pat_let56_0):
                    def iife144_(d_1382_dt__update__tmp_h2_):
                        def iife145_(_pat_let57_0):
                            def iife146_(d_1383_dt__update_hcf1_h2_):
                                return D1_DC1(d_1383_dt__update_hcf1_h2_)
                            return iife146_(_pat_let57_0)
                        return iife145_(default__.safeDivisionInt(pat_let_tv73_.f12, pat_let_tv74_.f12))
                    return iife144_(_pat_let56_0)
                (d_1377_v7_)[index216_] = iife143_(d_1376_v6_)
                d_1384_v8_: _dafny.Seq
                d_1384_v8_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                (d_1371_v1_).f12 = (0) - (((d_1368_i0_) * (len(d_1384_v8_))) + (d_1368_i0_))
                d_1385_v9_: C1
                nw225_ = C1()
                nw225_.ctor__(_dafny.SeqWithoutIsStrInference([D1_DC2((self).f10, p0, d_1368_i0_, len(d_1374_v4_)) for d_1386_i2_ in range(default__.abs(468))]))
                d_1385_v9_ = nw225_
                d_1387_v10_: _dafny.Map
                d_1387_v10_ = _dafny.Map({d_1371_v1_.f12: d_1385_v9_})
                d_1385_v9_ = ((d_1387_v10_)[(self).f10] if ((self).f10) in (d_1387_v10_) else d_1385_v9_)
            elif True:
                d_1388_v11_: _dafny.Seq
                d_1388_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                d_1389_v12_: _dafny.Set
                d_1389_v12_ = default__.fm14(d_1368_i0_, len(d_1388_v11_), (self).f10, globalState)
                d_1389_v12_ = (d_1389_v12_ if (self).f9 else d_1389_v12_)
                d_1390_v13_: _dafny.MultiSet
                d_1390_v13_ = _dafny.MultiSet([(self).f9])
                d_1391_v14_: _dafny.Seq
                d_1391_v14_ = _dafny.SeqWithoutIsStrInference([d_1388_v11_])
                d_1392_v15_: _dafny.Seq
                d_1392_v15_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])
                d_1393_v16_: _dafny.Seq
                d_1393_v16_ = _dafny.SeqWithoutIsStrInference([d_1368_i0_, d_1368_i0_, len(d_1392_v15_), (self).f10])
                d_1394_v17_: _dafny.Seq
                d_1394_v17_ = _dafny.SeqWithoutIsStrInference([default__.fm0(globalState), d_1368_i0_, len(d_1393_v16_)])
                d_1395_v18_: _dafny.Map
                d_1395_v18_ = _dafny.Map({len(d_1388_v11_): d_1394_v17_})
                d_1396_v19_: _dafny.MultiSet
                d_1396_v19_ = _dafny.MultiSet([d_1368_i0_])
                d_1397_v20_: _dafny.Map
                d_1397_v20_ = _dafny.Map({len(d_1388_v11_): d_1368_i0_})
                d_1398_v21_: _dafny.Map
                d_1398_v21_ = _dafny.Map({(self).f9: _dafny.MultiSet([(self).f10, 756])})
                d_1399_v22_: _dafny.Set
                d_1399_v22_ = _dafny.Set({d_1396_v19_, _dafny.MultiSet([(self).f10, len(d_1397_v20_)]), ((d_1398_v21_)[True] if (True) in (d_1398_v21_) else d_1396_v19_)})
                d_1400_v23_: _dafny.Array
                nw226_ = _dafny.Array(None, 21)
                nw226_[int(0)] = (self).fm4(p0, p0, (self).f9, p0, globalState)
                nw226_[int(1)] = ((d_1390_v13_).set(not((self).f9), default__.abs((self).f10))).cardinality
                nw226_[int(2)] = (self).f10
                nw226_[int(3)] = default__.safeModuloInt(d_1368_i0_, (self).f10)
                nw226_[int(4)] = (len(d_1391_v14_)) + ((0) - ((self).f10))
                nw226_[int(5)] = (self).f10
                nw226_[int(6)] = d_1368_i0_
                nw226_[int(7)] = default__.safeDivisionInt((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0]))).cardinality, len(d_1395_v18_))
                nw226_[int(8)] = (self).f10
                nw226_[int(9)] = (0) - (default__.safeModuloInt((d_1394_v17_)[default__.safeIndex(923, len(d_1394_v17_))], (self).f10))
                nw226_[int(10)] = (self).f10
                nw226_[int(11)] = d_1368_i0_
                nw226_[int(12)] = len(d_1392_v15_)
                nw226_[int(13)] = (self).fm4((self).f9, p0, not(False), not(p0), globalState)
                nw226_[int(14)] = (0) - (default__.safeDivisionInt(807, 9))
                nw226_[int(15)] = (self).f10
                nw226_[int(16)] = (self).f10
                nw226_[int(17)] = (self).f10
                nw226_[int(18)] = len(d_1399_v22_)
                nw226_[int(19)] = d_1368_i0_
                nw226_[int(20)] = d_1368_i0_
                d_1400_v23_ = nw226_
                index217_ = default__.safeIndex(405, (d_1400_v23_).length(0))
                (d_1400_v23_)[index217_] = d_1368_i0_
                index218_ = default__.safeIndex(405, (d_1400_v23_).length(0))
                (d_1400_v23_)[index218_] = (0) - ((self).f10)
                d_1401_v24_: _dafny.Array
                nw227_ = _dafny.Array(False, 19)
                d_1401_v24_ = nw227_
                d_1402_v25_: _dafny.Seq
                d_1402_v25_ = _dafny.SeqWithoutIsStrInference([d_1401_v24_])
                d_1401_v24_ = (d_1402_v25_)[default__.safeIndex((d_1400_v23_)[default__.safeIndex(405, (d_1400_v23_).length(0))], len(d_1402_v25_))]
                d_1388_v11_ = d_1388_v11_
                d_1403_v27_: _dafny.MultiSet
                d_1403_v27_ = d_1396_v19_
                def iife147_():
                    coll31_ = _dafny.Map()
                    compr_31_: int
                    for compr_31_ in ((d_1403_v27_)).Elements:
                        d_1404_v26_: int = compr_31_
                        if (d_1404_v26_) in ((d_1403_v27_)):
                            coll31_[(d_1404_v26_) - (d_1368_i0_)] = (d_1400_v23_)[default__.safeIndex(405, (d_1400_v23_).length(0))]
                    return _dafny.Map(coll31_)
                d_1397_v20_ = iife147_()
                
            d_1405_v28_: _dafny.Array
            nw228_ = _dafny.Array(None, 7)
            nw228_[int(0)] = (d_1368_i0_) * (d_1368_i0_)
            nw228_[int(1)] = 584
            nw228_[int(2)] = (self).f10
            nw228_[int(3)] = d_1368_i0_
            nw228_[int(4)] = (self).f10
            nw228_[int(5)] = (0) - ((self).f10)
            nw228_[int(6)] = ((self).f10) * (d_1368_i0_)
            d_1405_v28_ = nw228_
            index219_ = default__.safeIndex(790, (d_1405_v28_).length(0))
            (d_1405_v28_)[index219_] = d_1368_i0_
            d_1406_v29_: T1
            nw229_ = C2()
            nw229_.ctor__(d_1368_i0_, p0)
            d_1406_v29_ = nw229_
            d_1407_v30_: _dafny.Set
            d_1407_v30_ = _dafny.Set({(d_1406_v29_ if p0 else d_1406_v29_), d_1406_v29_, d_1406_v29_, d_1406_v29_, d_1406_v29_})
            d_1408_v31_: int
            d_1408_v31_ = 485
            index220_ = default__.safeIndex(790, (d_1405_v28_).length(0))
            rhs266_ = ((d_1408_v31_) + ((self).f10)) + ((d_1408_v31_) + (114))
            rhs267_ = ((d_1407_v30_) | (_dafny.Set({d_1406_v29_})) if True else d_1407_v30_)
            rhs268_ = default__.safeDivisionInt((self).f10, (self).f10)
            lhs164_ = d_1405_v28_
            lhs165_ = default__.safeIndex(790, (d_1405_v28_).length(0))
            lhs164_[lhs165_] = rhs266_
            d_1407_v30_ = rhs267_
            d_1408_v31_ = rhs268_
            d_1409_v32_: _dafny.Seq
            d_1409_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jk"))
            d_1410_v33_: str
            d_1410_v33_ = _dafny.CodePoint('q')
            d_1411_v34_: _dafny.MultiSet
            d_1411_v34_ = _dafny.MultiSet([d_1410_v33_, d_1410_v33_])
            d_1412_v35_: _dafny.Seq
            d_1412_v35_ = _dafny.SeqWithoutIsStrInference([(d_1411_v34_).cardinality, d_1408_v31_])
            d_1413_v36_: _dafny.Set
            d_1413_v36_ = _dafny.Set({d_1412_v35_})
            d_1414_v38_: _dafny.Array
            nw230_ = _dafny.Array(_dafny.Set({}), 21)
            d_1414_v38_ = nw230_
            d_1415_v39_: T2
            nw231_ = C3()
            nw231_.ctor__(p0, d_1414_v38_, (self).f9)
            d_1415_v39_ = nw231_
            d_1416_v40_: _dafny.Seq
            d_1416_v40_ = _dafny.SeqWithoutIsStrInference([d_1415_v39_])
            d_1417_v41_: _dafny.Seq
            d_1417_v41_ = _dafny.SeqWithoutIsStrInference([d_1416_v40_])
            d_1418_v42_: _dafny.MultiSet
            d_1418_v42_ = _dafny.MultiSet([len(d_1417_v41_), (self).f10])
            d_1419_v43_: _dafny.Seq
            def iife148_():
                coll32_ = _dafny.Map()
                compr_32_: int
                for compr_32_ in (d_1418_v42_).Elements:
                    d_1421_v37_: int = compr_32_
                    if (d_1421_v37_) in (d_1418_v42_):
                        coll32_[(d_1421_v37_) * (173)] = _dafny.SeqWithoutIsStrInference([(d_1415_v39_).f9])
                return _dafny.Map(coll32_)
            d_1419_v43_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1420_i3_ in range(default__.abs(96))]), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))).set(default__.safeIndex((self).f10, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))), d_1410_v33_), (d_1409_v32_ if (self).fm5(d_1413_v36_, iife148_()
            , (0) - (len(d_1409_v32_)), globalState) else d_1409_v32_)])
            d_1422_v44_: _dafny.Set
            d_1422_v44_ = _dafny.Set({not((self).f9), (d_1415_v39_).f9, (d_1406_v29_).f9})
            d_1423_v45_: _dafny.Map
            d_1423_v45_ = _dafny.Map({(d_1415_v39_).f9: (d_1410_v33_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "goklh")))})
            d_1424_v46_: _dafny.Seq
            d_1424_v46_ = _dafny.SeqWithoutIsStrInference([(d_1406_v29_).f9])
            rhs269_ = (d_1419_v43_)[default__.safeIndex((d_1412_v35_)[default__.safeIndex(len(d_1422_v44_), len(d_1412_v35_))], len(d_1419_v43_))]
            rhs270_ = _dafny.SeqWithoutIsStrInference([d_1410_v33_ for d_1425_i4_ in range(default__.abs(-945))])
            rhs271_ = not(((d_1423_v45_)[(d_1424_v46_)[default__.safeIndex((d_1405_v28_)[default__.safeIndex(790, (d_1405_v28_).length(0))], len(d_1424_v46_))]] if ((d_1424_v46_)[default__.safeIndex((d_1405_v28_)[default__.safeIndex(790, (d_1405_v28_).length(0))], len(d_1424_v46_))]) in (d_1423_v45_) else (self).f9))
            rhs272_ = (d_1415_v39_).f9
            lhs166_ = globalState
            lhs167_ = globalState
            d_1409_v32_ = rhs269_
            d_1409_v32_ = rhs270_
            lhs166_.f2 = rhs271_
            lhs167_.f2 = rhs272_
            d_1426_v47_: _dafny.Map
            d_1426_v47_ = _dafny.Map({False: (self).f10})
            d_1427_v48_: D4
            d_1427_v48_ = D4_DC12(not((self).f9), (d_1406_v29_).f9, len(d_1426_v47_), (self).f10, not((d_1406_v29_).f9))
            d_1428_v49_: D4
            d_1428_v49_ = D4_DC15(d_1427_v48_)
            source14_ = d_1428_v49_
            if source14_.is_DC12:
                d_1429___mcc_h0_ = source14_.cf20
                d_1430___mcc_h1_ = source14_.cf21
                d_1431___mcc_h2_ = source14_.cf22
                d_1432___mcc_h3_ = source14_.cf23
                d_1433___mcc_h4_ = source14_.cf24
                d_1434_cf24_ = d_1433___mcc_h4_
                d_1435_cf23_ = d_1432___mcc_h3_
                d_1436_cf22_ = d_1431___mcc_h2_
                d_1437_cf21_ = d_1430___mcc_h1_
                d_1438_cf20_ = d_1429___mcc_h0_
                d_1439_v50_: _dafny.Array
                def lambda76_(d_1440_i5_):
                    return False

                init42_ = lambda76_
                nw232_ = _dafny.Array(None, 24)
                for i0_42_ in range(nw232_.length(0)):
                    nw232_[i0_42_] = init42_(i0_42_)
                d_1439_v50_ = nw232_
                index221_ = default__.safeIndex(949, (d_1439_v50_).length(0))
                (d_1439_v50_)[index221_] = (d_1406_v29_).f9
                index222_ = default__.safeIndex(949, (d_1439_v50_).length(0))
                (d_1439_v50_)[index222_] = (d_1415_v39_).f9
                d_1435_cf23_ = (d_1412_v35_)[default__.safeIndex((d_1405_v28_)[default__.safeIndex(790, (d_1405_v28_).length(0))], len(d_1412_v35_))]
                index223_ = default__.safeIndex(790, (d_1405_v28_).length(0))
                (d_1405_v28_)[index223_] = (d_1412_v35_)[default__.safeIndex(d_1368_i0_, len(d_1412_v35_))]
                d_1441_v51_: C2
                nw233_ = C2()
                nw233_.ctor__((0) - (len(d_1422_v44_)), (d_1439_v50_)[default__.safeIndex(949, (d_1439_v50_).length(0))])
                d_1441_v51_ = nw233_
            elif source14_.is_DC13:
                d_1442___mcc_h5_ = source14_.cf25
                d_1443_cf25_ = d_1442___mcc_h5_
                d_1444_v52_: _dafny.MultiSet
                d_1444_v52_ = _dafny.MultiSet([(d_1424_v46_)[default__.safeIndex((d_1405_v28_)[default__.safeIndex(790, (d_1405_v28_).length(0))], len(d_1424_v46_))]])
                d_1444_v52_ = (_dafny.MultiSet(d_1424_v46_) if not((d_1415_v39_).f9) else (d_1444_v52_).set(p0, default__.abs((d_1405_v28_)[default__.safeIndex(790, (d_1405_v28_).length(0))])))
                d_1445_v53_: _dafny.Array
                nw234_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_1445_v53_ = nw234_
                d_1446_v54_: D14
                d_1446_v54_ = D14_DC38(d_1410_v33_, (d_1406_v29_).f9, d_1424_v46_)
                d_1447_v55_: _dafny.Array
                nw235_ = _dafny.Array(None, 15)
                nw235_[int(0)] = not(p0)
                nw235_[int(1)] = (d_1406_v29_).f9
                nw235_[int(2)] = p0
                nw235_[int(3)] = (self).f9
                nw235_[int(4)] = (d_1406_v29_).f9
                nw235_[int(5)] = (self).f9
                nw235_[int(6)] = (d_1446_v54_).cf68
                nw235_[int(7)] = (d_1406_v29_).f9
                nw235_[int(8)] = (d_1415_v39_).f9
                nw235_[int(9)] = not((d_1406_v29_).f9)
                nw235_[int(10)] = (d_1415_v39_).f9
                nw235_[int(11)] = (d_1415_v39_).f9
                nw235_[int(12)] = (d_1406_v29_).f9
                nw235_[int(13)] = (d_1415_v39_).f9
                nw235_[int(14)] = (self).f9
                d_1447_v55_ = nw235_
                index224_ = default__.safeIndex(616, (d_1445_v53_).length(0))
                (d_1445_v53_)[index224_] = d_1447_v55_
                index225_ = default__.safeIndex(616, (d_1445_v53_).length(0))
                (d_1445_v53_)[index225_] = d_1447_v55_
                d_1448_v56_: C2
                nw236_ = C2()
                nw236_.ctor__((self).f10, ((d_1415_v39_).f9 if (d_1406_v29_).f9 else p0))
                d_1448_v56_ = nw236_
                d_1449_v57_: _dafny.Seq
                d_1449_v57_ = _dafny.SeqWithoutIsStrInference([d_1448_v56_, d_1448_v56_, d_1448_v56_, d_1448_v56_])
                d_1448_v56_ = (d_1449_v57_)[default__.safeIndex(d_1408_v31_, len(d_1449_v57_))]
                d_1450_v58_: _dafny.MultiSet
                d_1450_v58_ = _dafny.MultiSet([d_1422_v44_])
                d_1451_v59_: D3
                d_1451_v59_ = D3_DC9(True, d_1424_v46_)
                d_1450_v58_ = (default__.fm43(_dafny.Set({len((d_1451_v59_).cf18), d_1368_i0_}), len(d_1426_v47_), default__.fm0(globalState), globalState)).set(d_1422_v44_, default__.abs(d_1443_cf25_))
            elif source14_.is_DC14:
                d_1452___mcc_h6_ = source14_.cf26
                d_1453___mcc_h7_ = source14_.cf27
                d_1454___mcc_h8_ = source14_.cf28
                d_1455_cf28_ = d_1454___mcc_h8_
                d_1456_cf27_ = d_1453___mcc_h7_
                d_1457_cf26_ = d_1452___mcc_h6_
                d_1458_v60_: _dafny.Array
                def lambda77_(d_1459_cf28_, d_1460_v33_, d_1461_v32_):
                    def lambda78_(d_1462_i6_):
                        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hksmc"))).set(default__.safeIndex(d_1459_cf28_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hksmc")))), d_1460_v33_)) + (d_1461_v32_)

                    return lambda78_

                init43_ = lambda77_(d_1455_cf28_, d_1410_v33_, d_1409_v32_)
                nw237_ = _dafny.Array(None, 9)
                for i0_43_ in range(nw237_.length(0)):
                    nw237_[i0_43_] = init43_(i0_43_)
                d_1458_v60_ = nw237_
                index226_ = default__.safeIndex(790, (d_1405_v28_).length(0))
                rhs273_ = d_1412_v35_
                rhs274_ = default__.fm0(globalState)
                rhs275_ = d_1458_v60_
                lhs168_ = d_1405_v28_
                lhs169_ = default__.safeIndex(790, (d_1405_v28_).length(0))
                d_1412_v35_ = rhs273_
                lhs168_[lhs169_] = rhs274_
                d_1458_v60_ = rhs275_
                d_1463_v61_: _dafny.Map
                d_1463_v61_ = _dafny.Map({(0) - (d_1408_v31_): (d_1406_v29_).f9})
                d_1464_v62_: _dafny.Set
                d_1464_v62_ = _dafny.Set({(self).f10, (self).f10})
                d_1465_v63_: _dafny.Set
                d_1465_v63_ = d_1464_v62_
                d_1466_v64_: _dafny.Map
                d_1466_v64_ = _dafny.Map({d_1465_v63_: d_1455_cf28_})
                d_1467_v65_: _dafny.Set
                d_1467_v65_ = _dafny.Set({d_1466_v64_})
                rhs276_ = (p0 if False else ((d_1463_v61_)[(d_1405_v28_)[default__.safeIndex(790, (d_1405_v28_).length(0))]] if ((d_1405_v28_)[default__.safeIndex(790, (d_1405_v28_).length(0))]) in (d_1463_v61_) else default__.fm1(_dafny.MultiSet([(0) - (d_1455_cf28_), (self).f10]), d_1455_cf28_, globalState)))
                rhs277_ = not(((d_1467_v65_ if not((d_1415_v39_).f9) else d_1467_v65_)).ispropersubset(d_1467_v65_))
                lhs170_ = globalState
                r0 = rhs276_
                lhs170_.f2 = rhs277_
                d_1409_v32_ = d_1409_v32_
                d_1468_v66_: _dafny.Map
                d_1468_v66_ = _dafny.Map({_dafny.MultiSet([d_1408_v31_]): (d_1406_v29_).f9})
                d_1469_v67_: D7
                d_1469_v67_ = D7_DC18(d_1468_v66_)
                pat_let_tv75_ = d_1468_v66_
                d_1470_v68_: _dafny.Array
                nw238_ = _dafny.Array(None, 9)
                nw238_[int(0)] = d_1469_v67_
                nw238_[int(1)] = d_1469_v67_
                nw238_[int(2)] = d_1469_v67_
                nw238_[int(3)] = d_1469_v67_
                nw238_[int(4)] = D7_DC18(d_1468_v66_)
                nw238_[int(5)] = d_1469_v67_
                nw238_[int(6)] = d_1469_v67_
                def iife149_(_pat_let58_0):
                    def iife150_(d_1471_dt__update__tmp_h3_):
                        def iife151_(_pat_let59_0):
                            def iife152_(d_1472_dt__update_hcf32_h0_):
                                return D7_DC18(d_1472_dt__update_hcf32_h0_)
                            return iife152_(_pat_let59_0)
                        return iife151_(pat_let_tv75_)
                    return iife150_(_pat_let58_0)
                nw238_[int(7)] = iife149_(d_1469_v67_)
                nw238_[int(8)] = D7_DC18(d_1468_v66_)
                d_1470_v68_ = nw238_
                index227_ = default__.safeIndex(341, (d_1470_v68_).length(0))
                (d_1470_v68_)[index227_] = d_1469_v67_
                index228_ = default__.safeIndex(341, (d_1470_v68_).length(0))
                (d_1470_v68_)[index228_] = d_1469_v67_
            elif source14_.is_DC11:
                d_1473___mcc_h9_ = source14_.cf19
                d_1474_cf19_ = d_1473___mcc_h9_
                d_1475_v69_: C4
                nw239_ = C4()
                nw239_.ctor__((d_1406_v29_).f9)
                d_1475_v69_ = nw239_
                d_1476_v70_: _dafny.MultiSet
                d_1476_v70_ = _dafny.MultiSet([default__.fm1(default__.fm29((self).f9, (d_1415_v39_).f9, globalState), d_1368_i0_, globalState), (d_1415_v39_).f9, (d_1406_v29_).f9])
                d_1408_v31_ = (0) - (((d_1476_v70_)[(d_1422_v44_).isdisjoint(d_1422_v44_)] if ((d_1422_v44_).isdisjoint(d_1422_v44_)) in (d_1476_v70_) else (0) - (d_1368_i0_)))
                d_1410_v33_ = d_1410_v33_
                d_1408_v31_ = default__.fm0(globalState)
            elif True:
                d_1477___mcc_h10_ = source14_.cf29
                d_1478_cf29_ = d_1477___mcc_h10_
                d_1479_v71_: C2
                nw240_ = C2()
                nw240_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xf"))), (self).f9)
                d_1479_v71_ = nw240_
                d_1480_v72_: _dafny.Seq
                d_1480_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_1415_v39_).f9, True])])
                index229_ = default__.safeIndex(790, (d_1405_v28_).length(0))
                (d_1405_v28_)[index229_] = len(d_1480_v72_)
                index230_ = default__.safeIndex(790, (d_1405_v28_).length(0))
                (d_1405_v28_)[index230_] = (len(d_1426_v47_)) + ((0) - ((d_1479_v71_).f14))
                index231_ = default__.safeIndex(790, (d_1405_v28_).length(0))
                (d_1405_v28_)[index231_] = d_1368_i0_
        d_1481_v73_: D1
        d_1481_v73_ = D1_DC4((self).f10, not(p0), p0, (self).f9, False)
        d_1482_v74_: C1
        nw241_ = C1()
        nw241_.ctor__((_dafny.SeqWithoutIsStrInference([D1_DC2(828, (self).f9, (self).f10, (self).f10) for d_1483_i7_ in range(default__.abs(207))])) + (default__.fm7((self).f10, (self).f9, d_1481_v73_, True, globalState)))
        d_1482_v74_ = nw241_
        d_1482_v74_ = d_1482_v74_
        d_1484_v75_: _dafny.Array
        nw242_ = _dafny.Array(_dafny.Set({}), 18)
        d_1484_v75_ = nw242_
        d_1485_v76_: C3
        nw243_ = C3()
        nw243_.ctor__(p0, d_1484_v75_, True)
        d_1485_v76_ = nw243_
        d_1486_v78_: _dafny.Map
        d_1486_v78_ = _dafny.Map({(self).f10: p0})
        d_1487_v79_: _dafny.Seq
        d_1487_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urpoc"))
        d_1488_v81_: _dafny.Seq
        d_1488_v81_ = _dafny.SeqWithoutIsStrInference([d_1486_v78_, d_1486_v78_])
        d_1489_v82_: _dafny.Array
        nw244_ = _dafny.Array(None, 27)
        def iife153_():
            coll33_ = _dafny.Map()
            compr_33_: int
            for compr_33_ in _dafny.IntegerRange(294, 229):
                d_1490_v77_: int = compr_33_
                if ((294) <= (d_1490_v77_)) and ((d_1490_v77_) < (229)):
                    coll33_[default__.safeModuloInt(d_1490_v77_, (self).f10)] = p0
            return _dafny.Map(coll33_)
        nw244_[int(0)] = (iife153_()
        ).set((0) - ((self).f10), (d_1485_v76_).f15)
        nw244_[int(1)] = (default__.fm15((d_1485_v76_).f15, globalState)) | (d_1486_v78_)
        nw244_[int(2)] = d_1486_v78_
        nw244_[int(3)] = _dafny.Map({(self).f10: (d_1485_v76_).f15})
        nw244_[int(4)] = d_1486_v78_
        nw244_[int(5)] = d_1486_v78_
        nw244_[int(6)] = _dafny.Map({(self).f10: p0})
        nw244_[int(7)] = d_1486_v78_
        nw244_[int(8)] = d_1486_v78_
        nw244_[int(9)] = ((d_1486_v78_).set((self).f10, not((d_1485_v76_).f15))) | (d_1486_v78_)
        nw244_[int(10)] = d_1486_v78_
        nw244_[int(11)] = d_1486_v78_
        nw244_[int(12)] = d_1486_v78_
        nw244_[int(13)] = d_1486_v78_
        nw244_[int(14)] = d_1486_v78_
        nw244_[int(15)] = d_1486_v78_
        nw244_[int(16)] = _dafny.Map({len(d_1487_v79_): p0})
        nw244_[int(17)] = d_1486_v78_
        nw244_[int(18)] = d_1486_v78_
        nw244_[int(19)] = (d_1486_v78_) | (d_1486_v78_)
        nw244_[int(20)] = default__.fm15((self).f9, globalState)
        nw244_[int(21)] = _dafny.Map({(self).f10: p0})
        nw244_[int(22)] = d_1486_v78_
        def iife154_():
            coll34_ = _dafny.Map()
            compr_34_: int
            for compr_34_ in _dafny.IntegerRange(670, -668):
                d_1491_v80_: int = compr_34_
                if ((670) <= (d_1491_v80_)) and ((d_1491_v80_) < (-668)):
                    coll34_[(d_1491_v80_) * ((self).f10)] = (self).f9
            return _dafny.Map(coll34_)
        nw244_[int(23)] = iife154_()
        
        nw244_[int(24)] = (default__.fm15(p0, globalState)) | ((d_1488_v81_)[default__.safeIndex((self).f10, len(d_1488_v81_))])
        nw244_[int(25)] = d_1486_v78_
        nw244_[int(26)] = (_dafny.Map({(self).f10: p0})) | (d_1486_v78_)
        d_1489_v82_ = nw244_
        index232_ = default__.safeIndex(713, (d_1489_v82_).length(0))
        (d_1489_v82_)[index232_] = d_1486_v78_
        d_1492_v83_: D1
        d_1492_v83_ = D1_DC2((self).f10, True, (self).f10, (self).f10)
        pat_let_tv76_ = d_1486_v78_
        pat_let_tv77_ = d_1486_v78_
        pat_let_tv78_ = d_1488_v81_
        pat_let_tv79_ = d_1488_v81_
        pat_let_tv80_ = d_1486_v78_
        pat_let_tv81_ = d_1486_v78_
        pat_let_tv82_ = d_1486_v78_
        index233_ = default__.safeIndex(713, (d_1489_v82_).length(0))
        def lambda79_(source15_):
            if source15_.is_DC2:
                d_1493___mcc_h11_ = source15_.cf2
                d_1494___mcc_h12_ = source15_.cf3
                d_1495___mcc_h13_ = source15_.cf4
                d_1496___mcc_h14_ = source15_.cf5
                d_1497_cf5_ = d_1496___mcc_h14_
                d_1498_cf4_ = d_1495___mcc_h13_
                d_1499_cf3_ = d_1494___mcc_h12_
                d_1500_cf2_ = d_1493___mcc_h11_
                return (pat_let_tv76_) | (pat_let_tv77_)
            elif source15_.is_DC3:
                d_1501___mcc_h15_ = source15_.cf6
                d_1502___mcc_h16_ = source15_.cf7
                d_1503___mcc_h17_ = source15_.cf8
                d_1504_cf8_ = d_1503___mcc_h17_
                d_1505_cf7_ = d_1502___mcc_h16_
                d_1506_cf6_ = d_1501___mcc_h15_
                return (pat_let_tv78_)[default__.safeIndex((self).f10, len(pat_let_tv79_))]
            elif source15_.is_DC4:
                d_1507___mcc_h18_ = source15_.cf9
                d_1508___mcc_h19_ = source15_.cf10
                d_1509___mcc_h20_ = source15_.cf11
                d_1510___mcc_h21_ = source15_.cf12
                d_1511___mcc_h22_ = source15_.cf13
                d_1512_cf13_ = d_1511___mcc_h22_
                d_1513_cf12_ = d_1510___mcc_h21_
                d_1514_cf11_ = d_1509___mcc_h20_
                d_1515_cf10_ = d_1508___mcc_h19_
                d_1516_cf9_ = d_1507___mcc_h18_
                return pat_let_tv80_
            elif source15_.is_DC1:
                d_1517___mcc_h23_ = source15_.cf1
                d_1518_cf1_ = d_1517___mcc_h23_
                return pat_let_tv81_
            elif True:
                d_1519___mcc_h24_ = source15_.cf14
                d_1520_cf14_ = d_1519___mcc_h24_
                return pat_let_tv82_

        (d_1489_v82_)[index233_] = lambda79_(d_1492_v83_)
        d_1521_v84_: _dafny.Array
        def lambda80_(d_1522_i8_):
            return (self).f9

        init44_ = lambda80_
        nw245_ = _dafny.Array(None, 22)
        for i0_44_ in range(nw245_.length(0)):
            nw245_[i0_44_] = init44_(i0_44_)
        d_1521_v84_ = nw245_
        index234_ = default__.safeIndex(515, (d_1521_v84_).length(0))
        (d_1521_v84_)[index234_] = True
        d_1523_v85_: _dafny.Seq
        d_1523_v85_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1524_v86_: _dafny.MultiSet
        d_1524_v86_ = _dafny.MultiSet([p0, (d_1485_v76_).f15])
        d_1525_v87_: _dafny.MultiSet
        d_1525_v87_ = _dafny.MultiSet([(self).f10, (d_1524_v86_).cardinality, 771, (self).f10, (self).f10])
        index235_ = default__.safeIndex(515, (d_1521_v84_).length(0))
        (d_1521_v84_)[index235_] = ((d_1523_v85_)[default__.safeIndex(((d_1525_v87_)[(self).f10] if ((self).f10) in (d_1525_v87_) else (self).f10), len(d_1523_v85_))] if True else False)
        if ((self).f10) >= ((self).f10):
            d_1526_v88_: _dafny.Seq
            d_1526_v88_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({len(d_1487_v79_), (self).f10})])
            d_1527_v89_: _dafny.Map
            d_1527_v89_ = _dafny.Map({(d_1485_v76_).fm32(d_1526_v88_, not((self).f9), (self).f10, False, globalState): d_1487_v79_})
            d_1528_v90_: _dafny.Seq
            d_1528_v90_ = _dafny.SeqWithoutIsStrInference([(self).f10, 668])
            d_1529_v91_: str
            d_1529_v91_ = _dafny.CodePoint('e')
            d_1527_v89_ = (d_1527_v89_).set((d_1487_v79_) <= (d_1487_v79_), ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgiaxc"))).set(default__.safeIndex((d_1528_v90_)[default__.safeIndex((self).f10, len(d_1528_v90_))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgiaxc")))), d_1529_v91_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehohxsxbk"))))
            d_1530_v92_: int
            d_1530_v92_ = 256
            d_1530_v92_ = (self).f10
            index236_ = default__.safeIndex(515, (d_1521_v84_).length(0))
            (d_1521_v84_)[index236_] = (self).f9
            d_1530_v92_ = d_1530_v92_
            d_1531_v93_: _dafny.Array
            def lambda81_(d_1532_v79_):
                def lambda82_(d_1533_i9_):
                    return d_1532_v79_

                return lambda82_

            init45_ = lambda81_(d_1487_v79_)
            nw246_ = _dafny.Array(None, 27)
            for i0_45_ in range(nw246_.length(0)):
                nw246_[i0_45_] = init45_(i0_45_)
            d_1531_v93_ = nw246_
            index237_ = default__.safeIndex(599, (d_1531_v93_).length(0))
            (d_1531_v93_)[index237_] = ((d_1487_v79_).set(default__.safeIndex(d_1530_v92_, len(d_1487_v79_)), (D13_DC35((d_1485_v76_).f15, d_1529_v91_, d_1482_v74_, (d_1485_v76_).f15)).cf59)).set(default__.safeIndex((self).f10, len((d_1487_v79_).set(default__.safeIndex(d_1530_v92_, len(d_1487_v79_)), (D13_DC35((d_1485_v76_).f15, d_1529_v91_, d_1482_v74_, (d_1485_v76_).f15)).cf59))), d_1529_v91_)
            index238_ = default__.safeIndex(599, (d_1531_v93_).length(0))
            (d_1531_v93_)[index238_] = (d_1487_v79_) + ((d_1487_v79_).set(default__.safeIndex((self).f10, len(d_1487_v79_)), d_1529_v91_))
        elif True:
            d_1534_v94_: C4
            nw247_ = C4()
            nw247_.ctor__(True)
            d_1534_v94_ = nw247_
            d_1535_v95_: _dafny.Array
            nw248_ = _dafny.Array(_dafny.Seq({}), 27)
            d_1535_v95_ = nw248_
            (d_1485_v76_).m1((0) - (-225), p0, p0, d_1535_v95_, globalState)
            d_1536_v96_: _dafny.Seq
            d_1536_v96_ = _dafny.SeqWithoutIsStrInference([D4_DC13(923)])
            d_1537_v97_: D4
            d_1537_v97_ = D4_DC15((d_1536_v96_)[default__.safeIndex((self).f10, len(d_1536_v96_))])
            d_1538_v98_: _dafny.Set
            d_1538_v98_ = _dafny.Set({D4_DC15(D4_DC13(863)), d_1537_v97_, d_1537_v97_})
            d_1538_v98_ = d_1538_v98_
            d_1539_v99_: _dafny.Seq
            d_1539_v99_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            d_1540_v100_: _dafny.Set
            d_1540_v100_ = _dafny.Set({d_1539_v99_})
            d_1541_v101_: C0
            nw249_ = C0()
            nw249_.ctor__((d_1540_v100_).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([-421 for d_1542_i10_ in range(default__.abs(905))]), _dafny.SeqWithoutIsStrInference([(self).f10])})), (self).f10)
            d_1541_v101_ = nw249_
            source16_ = default__.fm44(globalState)
            if source16_.is_DC19:
                d_1543___mcc_h25_ = source16_.cf33
                d_1544___mcc_h26_ = source16_.cf34
                d_1545___mcc_h27_ = source16_.cf35
                d_1546_cf35_ = d_1545___mcc_h27_
                d_1547_cf34_ = d_1544___mcc_h26_
                d_1548_cf33_ = d_1543___mcc_h25_
                (d_1541_v101_).f12 = default__.safeDivisionInt(((self).f10 if d_1547_cf34_ else len(d_1523_v85_)), d_1541_v101_.f12)
                d_1549_v102_: C3
                nw250_ = C3()
                nw250_.ctor__(True, d_1484_v75_, (d_1485_v76_).f15)
                d_1549_v102_ = nw250_
                d_1550_v103_: _dafny.Seq
                d_1550_v103_ = _dafny.SeqWithoutIsStrInference([d_1521_v84_, d_1521_v84_, d_1521_v84_, d_1521_v84_, d_1521_v84_])
                d_1521_v84_ = (d_1550_v103_)[default__.safeIndex(d_1541_v101_.f12, len(d_1550_v103_))]
                d_1551_v104_: _dafny.Map
                d_1551_v104_ = _dafny.Map({D1_DC1(len(_dafny.SeqWithoutIsStrInference([p0]))): (d_1524_v86_).set(False, default__.abs((0) - (d_1541_v101_.f12)))})
                pat_let_tv83_ = d_1541_v101_
                def iife155_(_pat_let60_0):
                    def iife156_(d_1552_dt__update__tmp_h4_):
                        def iife157_(_pat_let61_0):
                            def iife158_(d_1553_dt__update_hcf1_h3_):
                                return D1_DC1(d_1553_dt__update_hcf1_h3_)
                            return iife158_(_pat_let61_0)
                        return iife157_(pat_let_tv83_.f12)
                    return iife156_(_pat_let60_0)
                d_1551_v104_ = (d_1551_v104_).set(iife155_(D1_DC1(((d_1524_v86_)[(d_1485_v76_).f15] if ((d_1485_v76_).f15) in (d_1524_v86_) else d_1541_v101_.f12))), d_1524_v86_)
            elif source16_.is_DC20:
                d_1554___mcc_h28_ = source16_.cf36
                d_1555___mcc_h29_ = source16_.cf37
                d_1556_cf37_ = d_1555___mcc_h29_
                d_1557_cf36_ = d_1554___mcc_h28_
                d_1558_v105_: _dafny.Map
                d_1558_v105_ = _dafny.Map({default__.fm0(globalState): d_1541_v101_.f12})
                d_1557_cf36_ = not(((self).f10) not in (d_1558_v105_))
                d_1559_v106_: _dafny.Map
                d_1559_v106_ = _dafny.Map({False: (0) - ((len(d_1539_v99_)) - (d_1541_v101_.f12))})
                d_1559_v106_ = (d_1559_v106_).set(not(not((d_1521_v84_)[default__.safeIndex(515, (d_1521_v84_).length(0))])), 414)
                d_1560_v107_: _dafny.Map
                d_1560_v107_ = _dafny.Map({(d_1485_v76_).f15: d_1557_cf36_})
                d_1561_v108_: _dafny.Set
                d_1561_v108_ = _dafny.Set({p0, p0, p0})
                index239_ = default__.safeIndex(515, (d_1521_v84_).length(0))
                rhs278_ = (d_1541_v101_.f12) - ((self).f10)
                rhs279_ = (d_1524_v86_).issubset(_dafny.MultiSet([d_1557_cf36_]))
                rhs280_ = d_1524_v86_
                rhs281_ = ((d_1560_v107_)[(_dafny.Set({(d_1485_v76_).f15})).isdisjoint(d_1561_v108_)] if ((_dafny.Set({(d_1485_v76_).f15})).isdisjoint(d_1561_v108_)) in (d_1560_v107_) else False)
                lhs171_ = d_1541_v101_
                lhs172_ = globalState
                lhs173_ = d_1521_v84_
                lhs174_ = default__.safeIndex(515, (d_1521_v84_).length(0))
                lhs171_.f12 = rhs278_
                lhs172_.f2 = rhs279_
                d_1524_v86_ = rhs280_
                lhs173_[lhs174_] = rhs281_
                (d_1541_v101_).f12 = default__.fm0(globalState)
            elif source16_.is_DC18:
                d_1562___mcc_h30_ = source16_.cf32
                d_1563_cf32_ = d_1562___mcc_h30_
                d_1564_v109_: D15
                d_1564_v109_ = D15_DC40(_dafny.Set({(d_1485_v76_).f15}))
                d_1565_v110_: _dafny.Seq
                d_1565_v110_ = _dafny.SeqWithoutIsStrInference([(d_1564_v109_).cf72, (d_1564_v109_).cf72])
                d_1566_v111_: str
                d_1566_v111_ = _dafny.CodePoint('s')
                d_1567_v112_: _dafny.Map
                d_1567_v112_ = _dafny.Map({d_1541_v101_.f12: d_1566_v111_})
                d_1568_v113_: _dafny.Map
                d_1568_v113_ = _dafny.Map({(len(d_1565_v110_)) + (len(d_1487_v79_)): d_1567_v112_})
                d_1568_v113_ = (d_1568_v113_).set(d_1541_v101_.f12, d_1567_v112_)
                d_1569_v114_: C1
                nw251_ = C1()
                nw251_.ctor__((d_1482_v74_.f13) + (d_1482_v74_.f13))
                d_1569_v114_ = nw251_
                d_1570_v116_: D14
                def iife159_():
                    coll35_ = _dafny.Map()
                    compr_35_: D1
                    for compr_35_ in (default__.fm45(globalState)).Elements:
                        d_1571_v115_: D1 = compr_35_
                        if (d_1571_v115_) in (default__.fm45(globalState)):
                            coll35_[d_1571_v115_] = (self).f9
                    return _dafny.Map(coll35_)
                d_1570_v116_ = D14_DC38(d_1566_v111_, (d_1523_v85_)[default__.safeIndex(len(iife159_()
), len(d_1523_v85_))], d_1523_v85_)
                (globalState).f2 = (d_1570_v116_).cf68
                d_1572_v117_: _dafny.Array
                def lambda83_(d_1573_v101_):
                    def lambda84_(d_1574_i11_):
                        return (d_1574_i11_) + (d_1573_v101_.f12)

                    return lambda84_

                init46_ = lambda83_(d_1541_v101_)
                nw252_ = _dafny.Array(None, 18)
                for i0_46_ in range(nw252_.length(0)):
                    nw252_[i0_46_] = init46_(i0_46_)
                d_1572_v117_ = nw252_
                index240_ = default__.safeIndex(734, (d_1572_v117_).length(0))
                (d_1572_v117_)[index240_] = len((d_1487_v79_) + (d_1487_v79_))
                index241_ = default__.safeIndex(734, (d_1572_v117_).length(0))
                (d_1572_v117_)[index241_] = (self).f10
            elif True:
                d_1575___mcc_h31_ = source16_.cf38
                d_1576_cf38_ = d_1575___mcc_h31_
                (globalState).f2 = (d_1521_v84_)[default__.safeIndex(515, (d_1521_v84_).length(0))]
                (d_1541_v101_).f12 = (self).f10
                (globalState).f2 = (d_1485_v76_).f15
                index242_ = default__.safeIndex(89, (d_1535_v95_).length(0))
                (d_1535_v95_)[index242_] = d_1523_v85_
                d_1577_v118_: _dafny.Map
                d_1577_v118_ = _dafny.Map({(self).f9: 653})
                d_1578_v119_: _dafny.Map
                d_1578_v119_ = _dafny.Map({d_1485_v76_: d_1487_v79_})
                index243_ = default__.safeIndex(89, (d_1535_v95_).length(0))
                (d_1535_v95_)[index243_] = default__.fm41(D10_DC27(d_1577_v118_), (self).f10, (d_1539_v99_).set(default__.safeIndex(d_1541_v101_.f12, len(d_1539_v99_)), len(d_1578_v119_)), globalState)
        r0 = (d_1521_v84_)[default__.safeIndex(515, (d_1521_v84_).length(0))]
        return r0

    def m4(self, p0, p1, p2, globalState):
        d_1579_i0_: int
        d_1579_i0_ = 0
        with _dafny.label("9"):
            while not((self).f9):
                with _dafny.c_label("9"):
                    if (d_1579_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1579_i0_ = (d_1579_i0_) + (1)
                    d_1580_v0_: int
                    d_1580_v0_ = -100
                    d_1581_v1_: str
                    d_1581_v1_ = _dafny.CodePoint('r')
                    d_1582_v2_: _dafny.Seq
                    d_1582_v2_ = _dafny.SeqWithoutIsStrInference([d_1581_v1_])
                    d_1583_v3_: _dafny.MultiSet
                    d_1583_v3_ = _dafny.MultiSet([d_1581_v1_, d_1581_v1_])
                    d_1584_v4_: _dafny.Map
                    d_1584_v4_ = _dafny.Map({(self).f10: (_dafny.MultiSet(d_1582_v2_)).isdisjoint(d_1583_v3_)})
                    d_1580_v0_ = len(d_1584_v4_)
                    d_1585_v5_: _dafny.Array
                    nw253_ = _dafny.Array(False, 6)
                    d_1585_v5_ = nw253_
                    index244_ = default__.safeIndex(242, (d_1585_v5_).length(0))
                    (d_1585_v5_)[index244_] = p2
                    index245_ = default__.safeIndex(242, (d_1585_v5_).length(0))
                    (d_1585_v5_)[index245_] = (self).f9
                    d_1580_v0_ = default__.safeDivisionInt((self).f10, d_1580_v0_)
                    d_1586_v6_: _dafny.Seq
                    d_1586_v6_ = _dafny.SeqWithoutIsStrInference([(d_1585_v5_)[default__.safeIndex(242, (d_1585_v5_).length(0))]])
                    d_1587_v7_: _dafny.Map
                    d_1587_v7_ = _dafny.Map({(self).f9: len(d_1586_v6_)})
                    d_1588_v8_: _dafny.Set
                    d_1588_v8_ = _dafny.Set({(0) - (len(d_1587_v7_)), (self).f10, d_1580_v0_, -171})
                    index246_ = default__.safeIndex(242, (d_1585_v5_).length(0))
                    rhs282_ = True
                    rhs283_ = (default__.safeModuloInt(d_1580_v0_, d_1580_v0_)) + (len((_dafny.Set({(self).f10, -206})) - (d_1588_v8_)))
                    rhs284_ = False
                    rhs285_ = (d_1580_v0_) in (_dafny.Map({d_1580_v0_: (self).f9}))
                    rhs286_ = (self).f10
                    lhs175_ = d_1585_v5_
                    lhs176_ = default__.safeIndex(242, (d_1585_v5_).length(0))
                    lhs177_ = globalState
                    lhs178_ = globalState
                    lhs175_[lhs176_] = rhs282_
                    d_1580_v0_ = rhs283_
                    lhs177_.f2 = rhs284_
                    lhs178_.f2 = rhs285_
                    d_1580_v0_ = rhs286_
                    pass
            pass
        d_1589_v9_: _dafny.Seq
        d_1589_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ua"))
        d_1589_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iatft"))
        d_1590_v10_: int
        d_1590_v10_ = 415
        d_1591_v11_: _dafny.Array
        def lambda85_(d_1592_i1_):
            return (self).f9

        init47_ = lambda85_
        nw254_ = _dafny.Array(None, 1)
        for i0_47_ in range(nw254_.length(0)):
            nw254_[i0_47_] = init47_(i0_47_)
        d_1591_v11_ = nw254_
        d_1593_v12_: _dafny.Map
        d_1593_v12_ = _dafny.Map({d_1591_v11_: (p1) < (len(_dafny.SeqWithoutIsStrInference([p1])))})
        d_1594_v13_: D11
        d_1594_v13_ = D11_DC30()
        d_1595_v14_: _dafny.Seq
        d_1595_v14_ = _dafny.SeqWithoutIsStrInference([D11_DC30(), d_1594_v13_, d_1594_v13_, d_1594_v13_, d_1594_v13_])
        d_1596_v15_: _dafny.Map
        d_1596_v15_ = _dafny.Map({(self).f9: d_1595_v14_})
        rhs287_ = (0) - ((p1) + ((0) - (p1)))
        rhs288_ = len((d_1596_v15_) | ((_dafny.Map({(self).f9: (d_1595_v14_).set(default__.safeIndex((self).f10, len(d_1595_v14_)), d_1594_v13_)})) | (d_1596_v15_)))
        rhs289_ = d_1593_v12_
        d_1590_v10_ = rhs287_
        d_1590_v10_ = rhs288_
        d_1593_v12_ = rhs289_
        d_1597_v16_: _dafny.Map
        d_1597_v16_ = _dafny.Map({d_1589_v9_: (self).f9})
        d_1598_v17_: D9
        d_1598_v17_ = D9_DC23(d_1597_v16_)
        pat_let_tv84_ = d_1597_v16_
        d_1599_v18_: _dafny.Map
        def iife160_(_pat_let62_0):
            def iife161_(d_1600_dt__update__tmp_h0_):
                def iife162_(_pat_let63_0):
                    def iife163_(d_1601_dt__update_hcf40_h0_):
                        return D9_DC23(d_1601_dt__update_hcf40_h0_)
                    return iife163_(_pat_let63_0)
                return iife162_(pat_let_tv84_)
            return iife161_(_pat_let62_0)
        d_1599_v18_ = _dafny.Map({(iife160_(d_1598_v17_) if (self).f9 else d_1598_v17_): d_1589_v9_})
        d_1599_v18_ = (d_1599_v18_).set(d_1598_v17_, d_1589_v9_)
        d_1602_v19_: str
        d_1602_v19_ = _dafny.CodePoint('j')
        d_1602_v19_ = d_1602_v19_
        d_1603_v20_: _dafny.Array
        nw255_ = _dafny.Array(_dafny.Array(None, 0), 10)
        d_1603_v20_ = nw255_
        index247_ = default__.safeIndex(821, (d_1603_v20_).length(0))
        (d_1603_v20_)[index247_] = p0
        d_1604_v21_: _dafny.Array
        nw256_ = _dafny.Array(None, 18)
        d_1604_v21_ = nw256_
        d_1605_v22_: _dafny.Seq
        d_1605_v22_ = _dafny.SeqWithoutIsStrInference([d_1604_v21_])
        index248_ = default__.safeIndex(821, (d_1603_v20_).length(0))
        rhs290_ = p0
        rhs291_ = (d_1605_v22_)[default__.safeIndex((self).f10, len(d_1605_v22_))]
        lhs179_ = d_1603_v20_
        lhs180_ = default__.safeIndex(821, (d_1603_v20_).length(0))
        lhs179_[lhs180_] = rhs290_
        d_1604_v21_ = rhs291_

    def m5(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        d_1606_i0_: int
        d_1606_i0_ = 0
        with _dafny.label("10"):
            while ((self).f9 if (self).f9 else (self).f9):
                with _dafny.c_label("10"):
                    if (d_1606_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1606_i0_ = (d_1606_i0_) + (1)
                    d_1607_v0_: _dafny.Array
                    def lambda86_(d_1608_i1_):
                        return _dafny.CodePoint('x')

                    init48_ = lambda86_
                    nw257_ = _dafny.Array(None, 16)
                    for i0_48_ in range(nw257_.length(0)):
                        nw257_[i0_48_] = init48_(i0_48_)
                    d_1607_v0_ = nw257_
                    d_1609_v1_: D13
                    d_1609_v1_ = D13_DC34(d_1607_v0_)
                    d_1610_v2_: _dafny.Map
                    d_1610_v2_ = _dafny.Map({(self).f9: d_1609_v1_})
                    d_1611_v3_: _dafny.Map
                    d_1611_v3_ = _dafny.Map({(self).f9: len(d_1610_v2_)})
                    d_1611_v3_ = (d_1611_v3_) | ((d_1611_v3_ if False else (d_1611_v3_).set((self).f9, (self).f10)))
                    d_1612_v4_: _dafny.Seq
                    d_1612_v4_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])
                    d_1613_v5_: _dafny.Set
                    d_1613_v5_ = _dafny.Set({len((d_1612_v4_) + (d_1612_v4_)), (self).f10, (len(_dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])) if (self).f9 else (self).f10), 827})
                    d_1613_v5_ = d_1613_v5_
                    d_1614_v6_: _dafny.Array
                    def lambda87_(d_1615_i2_):
                        return _dafny.MultiSet([430, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))), (self).f10, (self).f10, (0) - ((self).f10)])

                    init49_ = lambda87_
                    nw258_ = _dafny.Array(None, 16)
                    for i0_49_ in range(nw258_.length(0)):
                        nw258_[i0_49_] = init49_(i0_49_)
                    d_1614_v6_ = nw258_
                    d_1614_v6_ = d_1614_v6_
                    d_1616_v7_: _dafny.Array
                    nw259_ = _dafny.Array(D12.default()(), 13)
                    d_1616_v7_ = nw259_
                    d_1617_v8_: _dafny.Seq
                    d_1617_v8_ = _dafny.SeqWithoutIsStrInference([d_1616_v7_, d_1616_v7_])
                    r0 = ((self).f9) or ((len(d_1617_v8_)) >= ((self).f10))
                    pass
            pass
        d_1618_v9_: _dafny.Array
        nw260_ = _dafny.Array(_dafny.Set({}), 10)
        d_1618_v9_ = nw260_
        d_1619_v10_: C3
        nw261_ = C3()
        nw261_.ctor__((self).f9, d_1618_v9_, (self).f9)
        d_1619_v10_ = nw261_
        d_1620_v11_: _dafny.Array
        nw262_ = _dafny.Array(None, 12)
        nw262_[int(0)] = d_1619_v10_
        nw262_[int(1)] = d_1619_v10_
        nw262_[int(2)] = d_1619_v10_
        nw262_[int(3)] = d_1619_v10_
        nw262_[int(4)] = d_1619_v10_
        nw262_[int(5)] = d_1619_v10_
        nw262_[int(6)] = d_1619_v10_
        nw262_[int(7)] = d_1619_v10_
        nw262_[int(8)] = d_1619_v10_
        nw262_[int(9)] = d_1619_v10_
        nw262_[int(10)] = d_1619_v10_
        nw262_[int(11)] = d_1619_v10_
        d_1620_v11_ = nw262_
        index249_ = default__.safeIndex(407, (d_1620_v11_).length(0))
        (d_1620_v11_)[index249_] = d_1619_v10_
        index250_ = default__.safeIndex(407, (d_1620_v11_).length(0))
        (d_1620_v11_)[index250_] = d_1619_v10_
        d_1621_v12_: _dafny.Seq
        d_1621_v12_ = _dafny.SeqWithoutIsStrInference([(d_1619_v10_).f15, (self).fm3(globalState)])
        hi7_ = len(d_1621_v12_)
        for d_1622_i3_ in range((0) - ((self).f10), hi7_):
            d_1623_v13_: D4
            d_1623_v13_ = D4_DC13((self).f10)
            d_1624_v14_: _dafny.Seq
            d_1624_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_1623_v13_).cf25})])
            d_1625_v15_: _dafny.Seq
            d_1625_v15_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            d_1626_v16_: _dafny.Map
            d_1626_v16_ = _dafny.Map({default__.fm1(_dafny.MultiSet(d_1625_v15_), d_1622_i3_, globalState): (d_1619_v10_).f15})
            (globalState).f2 = ((d_1619_v10_).f15) or ((d_1619_v10_).fm32(d_1624_v14_, (d_1619_v10_).f15, len(d_1626_v16_), (self).f9, globalState))
            d_1627_v17_: T1
            nw263_ = C2()
            nw263_.ctor__(d_1622_i3_, (self).f9)
            d_1627_v17_ = nw263_
            d_1628_v18_: _dafny.Map
            d_1628_v18_ = _dafny.Map({(d_1619_v10_).f15: d_1627_v17_})
            d_1629_v19_: _dafny.Set
            d_1629_v19_ = _dafny.Set({d_1622_i3_, len(d_1628_v18_)})
            d_1630_v20_: _dafny.Set
            d_1630_v20_ = d_1629_v19_
            source17_ = d_1630_v20_
            d_1631___mcc_h0_ = source17_
            d_1632_cf15_ = d_1631___mcc_h0_
            d_1633_v21_: D13
            d_1633_v21_ = D13_DC36(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpwdus")), d_1622_i3_, (self).f10, (self).f10)
            r2 = ((d_1633_v21_).cf64) * (d_1622_i3_)
            d_1634_v22_: _dafny.Array
            nw264_ = _dafny.Array(False, 13)
            d_1634_v22_ = nw264_
            index251_ = default__.safeIndex(698, (d_1634_v22_).length(0))
            (d_1634_v22_)[index251_] = (d_1619_v10_).f15
            d_1635_v23_: _dafny.Map
            d_1635_v23_ = _dafny.Map({(d_1627_v17_).f9: d_1622_i3_})
            d_1636_v25_: _dafny.Set
            d_1636_v25_ = _dafny.Set({d_1625_v15_})
            d_1637_v26_: C0
            nw265_ = C0()
            nw265_.ctor__(d_1636_v25_, d_1622_i3_)
            d_1637_v26_ = nw265_
            d_1638_v27_: _dafny.Map
            d_1638_v27_ = _dafny.Map({d_1637_v26_: (self).f10})
            d_1639_v28_: _dafny.Map
            d_1639_v28_ = _dafny.Map({-4: d_1638_v27_})
            d_1640_v29_: _dafny.MultiSet
            def iife164_():
                coll36_ = _dafny.Set()
                compr_36_: int
                for compr_36_ in _dafny.IntegerRange(-117, 410):
                    d_1641_v24_: int = compr_36_
                    if ((-117) <= (d_1641_v24_)) and ((d_1641_v24_) < (410)):
                        coll36_ = coll36_.union(_dafny.Set([default__.safeDivisionInt(d_1641_v24_, -521)]))
                return _dafny.Set(coll36_)
            d_1640_v29_ = _dafny.MultiSet([((d_1635_v23_)[(d_1619_v10_).f15] if ((d_1619_v10_).f15) in (d_1635_v23_) else len(iife164_()
            )), d_1622_i3_, len(((d_1639_v28_)[(self).f10] if ((self).f10) in (d_1639_v28_) else d_1638_v27_))])
            index252_ = default__.safeIndex(698, (d_1634_v22_).length(0))
            (d_1634_v22_)[index252_] = ((d_1619_v10_).f15 if (d_1619_v10_).f15 else not((default__.fm29((self).f9, (self).f9, globalState)).issubset(d_1640_v29_)))
            r1 = 342
            (d_1637_v26_).f12 = (0) - (d_1622_i3_)
            d_1642_v30_: _dafny.Set
            d_1642_v30_ = _dafny.Set({(d_1627_v17_).f9, (d_1619_v10_).f15, (d_1619_v10_).f15})
            rhs292_ = ((self).f10) != ((self).f10)
            rhs293_ = d_1642_v30_
            rhs294_ = default__.safeModuloInt(d_1622_i3_, default__.safeModuloInt((d_1625_v15_)[default__.safeIndex((0) - ((self).f10), len(d_1625_v15_))], d_1622_i3_))
            r0 = rhs292_
            d_1642_v30_ = rhs293_
            r2 = rhs294_
            d_1643_v31_: C2
            nw266_ = C2()
            nw266_.ctor__(d_1622_i3_, not((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyy")))) > (default__.fm0(globalState))))
            d_1643_v31_ = nw266_
        d_1644_v32_: _dafny.Set
        d_1644_v32_ = _dafny.Set({(d_1619_v10_).f15})
        d_1645_v33_: D15
        d_1645_v33_ = D15_DC40(d_1644_v32_)
        source18_ = d_1645_v33_
        if source18_.is_DC41:
            r2 = (self).f10
            d_1646_v34_: _dafny.Array
            def lambda88_(d_1647_v10_):
                def lambda89_(d_1648_i4_):
                    return (d_1647_v10_).f15

                return lambda89_

            init50_ = lambda88_(d_1619_v10_)
            nw267_ = _dafny.Array(None, 15)
            for i0_50_ in range(nw267_.length(0)):
                nw267_[i0_50_] = init50_(i0_50_)
            d_1646_v34_ = nw267_
            index253_ = default__.safeIndex(472, (d_1646_v34_).length(0))
            (d_1646_v34_)[index253_] = (self).f9
            index254_ = default__.safeIndex(472, (d_1646_v34_).length(0))
            (d_1646_v34_)[index254_] = (d_1619_v10_).f15
            d_1649_v35_: _dafny.Seq
            d_1649_v35_ = _dafny.SeqWithoutIsStrInference([len(d_1644_v32_), (self).f10])
            d_1650_v36_: _dafny.Set
            d_1650_v36_ = _dafny.Set({d_1649_v35_})
            d_1651_v37_: _dafny.Array
            def lambda90_(d_1652_i5_):
                return default__.safeDivisionInt(d_1652_i5_, (self).f10)

            init51_ = lambda90_
            nw268_ = _dafny.Array(None, 11)
            for i0_51_ in range(nw268_.length(0)):
                nw268_[i0_51_] = init51_(i0_51_)
            d_1651_v37_ = nw268_
            d_1653_v38_: _dafny.Set
            d_1653_v38_ = _dafny.Set({d_1651_v37_, d_1651_v37_, d_1651_v37_})
            d_1654_v39_: _dafny.Map
            d_1654_v39_ = _dafny.Map({(self).fm5(d_1650_v36_, _dafny.Map({len(d_1653_v38_): d_1621_v12_}), len(_dafny.SeqWithoutIsStrInference([(self).f9])), globalState): (self).f10})
            d_1655_v40_: _dafny.MultiSet
            d_1655_v40_ = _dafny.MultiSet([len(d_1654_v39_), (self).f10])
            d_1656_v41_: _dafny.Map
            d_1656_v41_ = _dafny.Map({(d_1619_v10_).f15: default__.fm1(d_1655_v40_, (d_1649_v35_)[default__.safeIndex((self).f10, len(d_1649_v35_))], globalState)})
            d_1656_v41_ = (d_1656_v41_).set((d_1619_v10_).f15, not((d_1646_v34_)[default__.safeIndex(472, (d_1646_v34_).length(0))]))
            d_1657_v42_: _dafny.MultiSet
            d_1657_v42_ = _dafny.MultiSet([_dafny.CodePoint('x')])
            d_1658_v43_: C0
            nw269_ = C0()
            nw269_.ctor__(d_1650_v36_, (d_1657_v42_).cardinality)
            d_1658_v43_ = nw269_
            d_1658_v43_ = d_1658_v43_
        elif source18_.is_DC42:
            d_1659___mcc_h1_ = source18_.cf73
            d_1660___mcc_h2_ = source18_.cf74
            d_1661___mcc_h3_ = source18_.cf75
            d_1662___mcc_h4_ = source18_.cf76
            d_1663_cf76_ = d_1662___mcc_h4_
            d_1664_cf75_ = d_1661___mcc_h3_
            d_1665_cf74_ = d_1660___mcc_h2_
            d_1666_cf73_ = d_1659___mcc_h1_
            d_1667_v45_: _dafny.Seq
            d_1667_v45_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_1663_cf76_, (self).f10])).cardinality])
            d_1668_v46_: _dafny.MultiSet
            d_1668_v46_ = _dafny.MultiSet([48, d_1663_cf76_])
            d_1669_v47_: _dafny.MultiSet
            d_1669_v47_ = _dafny.MultiSet([_dafny.MultiSet(d_1667_v45_), d_1668_v46_])
            d_1670_v48_: str
            d_1670_v48_ = _dafny.CodePoint('h')
            d_1671_v49_: _dafny.Map
            d_1671_v49_ = _dafny.Map({_dafny.MultiSet([(self).f10, d_1665_cf74_]): d_1670_v48_})
            d_1672_v50_: _dafny.Map
            def iife165_():
                coll37_ = _dafny.Map()
                compr_37_: _dafny.MultiSet
                for compr_37_ in (d_1669_v47_).Elements:
                    d_1673_v44_: _dafny.MultiSet = compr_37_
                    if (d_1673_v44_) in (d_1669_v47_):
                        coll37_[d_1673_v44_] = _dafny.CodePoint('r')
                return _dafny.Map(coll37_)
            d_1672_v50_ = _dafny.Map({(iife165_()
             if not(not(d_1666_cf73_)) else d_1671_v49_): default__.safeModuloInt((self).f10, len(d_1667_v45_))})
            d_1672_v50_ = (d_1672_v50_).set(d_1671_v49_, (self).f10)
            d_1674_v51_: _dafny.Array
            nw270_ = _dafny.Array(D4.default()(), 9)
            d_1674_v51_ = nw270_
            d_1675_v52_: _dafny.Set
            d_1675_v52_ = _dafny.Set({(d_1674_v51_ if (self).f9 else d_1674_v51_)})
            d_1675_v52_ = d_1675_v52_
            d_1665_cf74_ = d_1665_cf74_
            d_1676_v53_: _dafny.Set
            d_1676_v53_ = _dafny.Set({(self).f10})
            d_1677_v54_: _dafny.Seq
            d_1677_v54_ = _dafny.SeqWithoutIsStrInference([d_1676_v53_, _dafny.Set({d_1663_cf76_})])
            d_1678_v55_: _dafny.MultiSet
            d_1678_v55_ = _dafny.MultiSet([(d_1619_v10_).f15])
            d_1679_v56_: _dafny.Seq
            d_1679_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxndbo"))
            d_1680_v57_: D1
            d_1680_v57_ = D1_DC2(((d_1678_v55_).set(True, default__.abs(len(d_1679_v56_)))).cardinality, d_1666_cf73_, len(d_1679_v56_), d_1663_cf76_)
            d_1681_v58_: _dafny.Map
            d_1681_v58_ = _dafny.Map({d_1670_v48_: d_1668_v46_})
            d_1682_v59_: _dafny.Array
            nw271_ = _dafny.Array(None, 25)
            nw271_[int(0)] = (d_1619_v10_).fm32(d_1677_v54_, (self).fm3(globalState), 328, (d_1680_v57_).cf3, globalState)
            nw271_[int(1)] = (True if (d_1619_v10_).f15 else (d_1664_cf75_).f9)
            nw271_[int(2)] = (self).f9
            nw271_[int(3)] = not(((d_1664_cf75_).f9 if (d_1664_cf75_).f9 else (d_1664_cf75_).f9))
            nw271_[int(4)] = d_1666_cf73_
            nw271_[int(5)] = False
            nw271_[int(6)] = (d_1664_cf75_).f9
            nw271_[int(7)] = (d_1619_v10_).f15
            nw271_[int(8)] = (self).f9
            nw271_[int(9)] = (self).f9
            nw271_[int(10)] = d_1666_cf73_
            nw271_[int(11)] = not((d_1664_cf75_).fm3(globalState))
            nw271_[int(12)] = (d_1665_cf74_) <= ((self).f10)
            nw271_[int(13)] = (d_1664_cf75_).f9
            nw271_[int(14)] = (self).f9
            nw271_[int(15)] = (d_1619_v10_).f15
            nw271_[int(16)] = not ((d_1619_v10_).f15) or (False)
            nw271_[int(17)] = (d_1664_cf75_).fm3(globalState)
            nw271_[int(18)] = d_1666_cf73_
            nw271_[int(19)] = ((self).f10) > (d_1663_cf76_)
            nw271_[int(20)] = (d_1664_cf75_).f9
            nw271_[int(21)] = (d_1619_v10_).f15
            nw271_[int(22)] = ((self).f10) < (d_1665_cf74_)
            nw271_[int(23)] = not(default__.fm1(((d_1681_v58_)[(d_1679_v56_)[default__.safeIndex(299, len(d_1679_v56_))]] if ((d_1679_v56_)[default__.safeIndex(299, len(d_1679_v56_))]) in (d_1681_v58_) else d_1668_v46_), (self).f10, globalState))
            nw271_[int(24)] = (d_1619_v10_).f15
            d_1682_v59_ = nw271_
            d_1683_v60_: C1
            nw272_ = C1()
            nw272_.ctor__(_dafny.SeqWithoutIsStrInference([d_1680_v57_, d_1680_v57_]))
            d_1683_v60_ = nw272_
            d_1684_v61_: D13
            d_1684_v61_ = D13_DC35(d_1666_cf73_, d_1670_v48_, d_1683_v60_, d_1666_cf73_)
            index255_ = default__.safeIndex(10, (d_1682_v59_).length(0))
            (d_1682_v59_)[index255_] = (d_1684_v61_).cf58
            d_1685_v62_: _dafny.Map
            d_1685_v62_ = _dafny.Map({d_1665_cf74_: d_1665_cf74_})
            index256_ = default__.safeIndex(10, (d_1682_v59_).length(0))
            rhs295_ = (d_1621_v12_)[default__.safeIndex(default__.safeModuloInt(len(d_1685_v62_), d_1665_cf74_), len(d_1621_v12_))]
            rhs296_ = d_1621_v12_
            rhs297_ = d_1666_cf73_
            lhs181_ = d_1682_v59_
            lhs182_ = default__.safeIndex(10, (d_1682_v59_).length(0))
            lhs183_ = globalState
            lhs181_[lhs182_] = rhs295_
            d_1621_v12_ = rhs296_
            lhs183_.f2 = rhs297_
        elif source18_.is_DC43:
            d_1686_v63_: _dafny.Seq
            d_1686_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nf"))
            d_1687_v64_: D1
            d_1687_v64_ = D1_DC2((self).f10, (d_1619_v10_).f15, (self).f10, len(d_1686_v63_))
            d_1688_v65_: _dafny.Seq
            d_1688_v65_ = _dafny.SeqWithoutIsStrInference([d_1687_v64_])
            d_1689_v66_: C1
            nw273_ = C1()
            nw273_.ctor__(d_1688_v65_)
            d_1689_v66_ = nw273_
            d_1690_v67_: _dafny.MultiSet
            d_1690_v67_ = _dafny.MultiSet([d_1689_v66_, d_1689_v66_])
            d_1691_v68_: _dafny.Map
            d_1691_v68_ = _dafny.Map({d_1690_v67_: (d_1619_v10_).f15})
            d_1691_v68_ = (d_1691_v68_).set((d_1690_v67_).set(d_1689_v66_, default__.abs(677)), (d_1619_v10_).f15)
            d_1692_v69_: _dafny.Set
            d_1692_v69_ = _dafny.Set({len(d_1686_v63_)})
            d_1693_v70_: _dafny.Seq
            d_1693_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f10, (self).f10}), d_1692_v69_])
            d_1694_v71_: D3
            d_1694_v71_ = D3_DC9((self).f9, d_1621_v12_)
            d_1695_v72_: _dafny.Seq
            d_1695_v72_ = _dafny.SeqWithoutIsStrInference([-45, default__.safeModuloInt((self).fm4(False, (d_1619_v10_).fm32(d_1693_v70_, (d_1694_v71_).cf17, (self).f10, (self).fm3(globalState), globalState), (d_1619_v10_).f15, (d_1619_v10_).f15, globalState), (self).f10)])
            d_1695_v72_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f10) for d_1696_i6_ in range(default__.abs(-392))])
            r0 = (d_1621_v12_)[default__.safeIndex((self).f10, len(d_1621_v12_))]
            d_1697_v73_: _dafny.Map
            d_1697_v73_ = _dafny.Map({d_1695_v72_: d_1686_v63_})
            d_1698_v75_: _dafny.Map
            d_1698_v75_ = _dafny.Map({(self).f10: _dafny.SeqWithoutIsStrInference([(d_1619_v10_).f15])})
            d_1699_v76_: _dafny.Array
            nw274_ = _dafny.Array(None, 11)
            nw274_[int(0)] = (_dafny.SeqWithoutIsStrInference([False])) < (d_1621_v12_)
            nw274_[int(1)] = (self).f9
            nw274_[int(2)] = (self).f9
            def iife166_():
                coll38_ = _dafny.Set()
                compr_38_: _dafny.Seq
                for compr_38_ in (d_1697_v73_).keys.Elements:
                    d_1700_v74_: _dafny.Seq = compr_38_
                    if (d_1700_v74_) in (d_1697_v73_):
                        coll38_ = coll38_.union(_dafny.Set([d_1700_v74_]))
                return _dafny.Set(coll38_)
            nw274_[int(3)] = (self).fm5(iife166_()
            , d_1698_v75_, (self).f10, globalState)
            nw274_[int(4)] = (d_1619_v10_).f15
            nw274_[int(5)] = (d_1619_v10_).f15
            nw274_[int(6)] = (self).f9
            nw274_[int(7)] = (d_1619_v10_).f15
            nw274_[int(8)] = False
            nw274_[int(9)] = (d_1619_v10_).fm3(globalState)
            nw274_[int(10)] = True
            d_1699_v76_ = nw274_
            index257_ = default__.safeIndex(130, (d_1699_v76_).length(0))
            (d_1699_v76_)[index257_] = (d_1619_v10_).f15
            index258_ = default__.safeIndex(130, (d_1699_v76_).length(0))
            (d_1699_v76_)[index258_] = not((d_1619_v10_).f15)
        elif source18_.is_DC40:
            d_1701___mcc_h5_ = source18_.cf72
            d_1702_cf72_ = d_1701___mcc_h5_
            d_1703_v77_: _dafny.Seq
            d_1703_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lf"))
            d_1704_v78_: D1
            d_1704_v78_ = D1_DC2((self).f10, (d_1619_v10_).f15, len(d_1703_v77_), (self).f10)
            d_1705_v79_: _dafny.Seq
            d_1705_v79_ = _dafny.SeqWithoutIsStrInference([d_1704_v78_, d_1704_v78_])
            d_1706_v80_: C1
            nw275_ = C1()
            nw275_.ctor__(d_1705_v79_)
            d_1706_v80_ = nw275_
            d_1707_v81_: _dafny.Seq
            d_1707_v81_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            d_1707_v81_ = d_1707_v81_
            r1 = (self).f10
            d_1707_v81_ = d_1707_v81_
        elif True:
            d_1708___mcc_h6_ = source18_.cf77
            d_1709_cf77_ = d_1708___mcc_h6_
            d_1710_v82_: _dafny.Seq
            d_1710_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxjy"))
            d_1711_v83_: _dafny.Map
            d_1711_v83_ = _dafny.Map({d_1710_v82_: (0) - ((self).f10)})
            d_1712_v84_: _dafny.Set
            d_1712_v84_ = _dafny.Set({((0) - (len(d_1711_v83_))) - ((self).f10), (self).f10, (self).f10})
            d_1713_v85_: str
            d_1713_v85_ = _dafny.CodePoint('l')
            d_1714_v86_: _dafny.Map
            d_1714_v86_ = _dafny.Map({(D7_DC20(True, d_1713_v85_)).cf36: not(True)})
            d_1712_v84_ = ((d_1712_v84_) - (_dafny.Set({(self).f10, len(d_1621_v12_), len(d_1714_v86_), (self).f10}))) | (d_1712_v84_)
            d_1715_v87_: T2
            nw276_ = C3()
            nw276_.ctor__((d_1619_v10_).f15, d_1618_v9_, (d_1619_v10_).f15)
            d_1715_v87_ = nw276_
            d_1716_v88_: _dafny.Array
            nw277_ = _dafny.Array(None, 28)
            nw277_[int(0)] = d_1715_v87_
            nw277_[int(1)] = d_1715_v87_
            nw277_[int(2)] = d_1715_v87_
            nw277_[int(3)] = d_1715_v87_
            nw277_[int(4)] = d_1715_v87_
            nw277_[int(5)] = d_1715_v87_
            nw277_[int(6)] = d_1715_v87_
            nw277_[int(7)] = d_1715_v87_
            nw277_[int(8)] = d_1715_v87_
            nw277_[int(9)] = d_1715_v87_
            nw277_[int(10)] = d_1715_v87_
            nw277_[int(11)] = d_1715_v87_
            nw277_[int(12)] = d_1715_v87_
            nw277_[int(13)] = d_1715_v87_
            nw277_[int(14)] = d_1715_v87_
            nw277_[int(15)] = d_1715_v87_
            nw277_[int(16)] = d_1715_v87_
            nw277_[int(17)] = d_1715_v87_
            nw277_[int(18)] = d_1715_v87_
            nw277_[int(19)] = d_1715_v87_
            nw277_[int(20)] = d_1715_v87_
            nw277_[int(21)] = d_1715_v87_
            nw277_[int(22)] = d_1715_v87_
            nw277_[int(23)] = d_1715_v87_
            nw277_[int(24)] = d_1715_v87_
            nw277_[int(25)] = d_1715_v87_
            nw277_[int(26)] = d_1715_v87_
            nw277_[int(27)] = d_1715_v87_
            d_1716_v88_ = nw277_
            index259_ = default__.safeIndex(388, (d_1716_v88_).length(0))
            (d_1716_v88_)[index259_] = d_1715_v87_
            index260_ = default__.safeIndex(388, (d_1716_v88_).length(0))
            (d_1716_v88_)[index260_] = d_1715_v87_
            d_1710_v82_ = d_1710_v82_
            if not((((d_1619_v10_).f15 if (d_1619_v10_).f15 else (d_1715_v87_).f9) if ((self).f10) <= ((self).f10) else (self).f9)):
                d_1717_v89_: _dafny.Seq
                d_1717_v89_ = _dafny.SeqWithoutIsStrInference([((self).f10) - ((self).f10)])
                d_1718_v90_: _dafny.Array
                nw278_ = _dafny.Array(False, 18)
                d_1718_v90_ = nw278_
                index261_ = default__.safeIndex(590, (d_1718_v90_).length(0))
                (d_1718_v90_)[index261_] = True
                d_1719_v91_: _dafny.MultiSet
                d_1719_v91_ = _dafny.MultiSet([(self).f10])
                index262_ = default__.safeIndex(590, (d_1718_v90_).length(0))
                rhs298_ = (d_1717_v89_) + ((_dafny.SeqWithoutIsStrInference([(self).f10])) + (d_1717_v89_))
                rhs299_ = default__.fm1(d_1719_v91_, (self).f10, globalState)
                lhs184_ = d_1718_v90_
                lhs185_ = default__.safeIndex(590, (d_1718_v90_).length(0))
                d_1717_v89_ = rhs298_
                lhs184_[lhs185_] = rhs299_
                d_1720_v92_: _dafny.Map
                d_1720_v92_ = _dafny.Map({242: d_1710_v82_})
                d_1721_v93_: _dafny.Map
                d_1721_v93_ = _dafny.Map({(d_1715_v87_).f9: ((d_1720_v92_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lm")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lm")))) in (d_1720_v92_) else d_1710_v82_)})
                d_1722_v94_: _dafny.Seq
                d_1722_v94_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1713_v85_ for d_1723_i9_ in range(default__.abs(-116))])])
                d_1724_v95_: _dafny.Array
                nw279_ = _dafny.Array(None, 25)
                nw279_[int(0)] = d_1710_v82_
                nw279_[int(1)] = (d_1710_v82_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouulicrh")))
                nw279_[int(2)] = d_1710_v82_
                nw279_[int(3)] = d_1710_v82_
                nw279_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cc")) if (d_1619_v10_).f15 else d_1710_v82_)
                nw279_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1713_v85_ for d_1725_i7_ in range(default__.abs(652))])
                nw279_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                nw279_[int(7)] = d_1710_v82_
                nw279_[int(8)] = d_1710_v82_
                nw279_[int(9)] = d_1710_v82_
                nw279_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anxityp"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))
                nw279_[int(11)] = d_1710_v82_
                nw279_[int(12)] = ((d_1721_v93_)[(d_1715_v87_).f9] if ((d_1715_v87_).f9) in (d_1721_v93_) else d_1710_v82_)
                nw279_[int(13)] = d_1710_v82_
                nw279_[int(14)] = default__.fm28((self).f10, not((self).f9), globalState)
                nw279_[int(15)] = d_1710_v82_
                nw279_[int(16)] = d_1710_v82_
                nw279_[int(17)] = _dafny.SeqWithoutIsStrInference([d_1713_v85_ for d_1726_i8_ in range(default__.abs(839))])
                nw279_[int(18)] = (d_1722_v94_)[default__.safeIndex((self).f10, len(d_1722_v94_))]
                nw279_[int(19)] = _dafny.SeqWithoutIsStrInference([d_1713_v85_ for d_1727_i10_ in range(default__.abs(859))])
                nw279_[int(20)] = _dafny.SeqWithoutIsStrInference([d_1713_v85_ for d_1728_i11_ in range(default__.abs(975))])
                nw279_[int(21)] = d_1710_v82_
                nw279_[int(22)] = d_1710_v82_
                nw279_[int(23)] = d_1710_v82_
                nw279_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fc"))
                d_1724_v95_ = nw279_
                d_1729_v96_: _dafny.Map
                d_1729_v96_ = _dafny.Map({d_1720_v92_: d_1724_v95_})
                d_1724_v95_ = ((d_1729_v96_)[(d_1720_v92_) | (d_1720_v92_)] if ((d_1720_v92_) | (d_1720_v92_)) in (d_1729_v96_) else d_1724_v95_)
                d_1730_v97_: _dafny.MultiSet
                d_1730_v97_ = _dafny.MultiSet([(d_1718_v90_)[default__.safeIndex(590, (d_1718_v90_).length(0))]])
                d_1730_v97_ = d_1730_v97_
                d_1731_v98_: C2
                nw280_ = C2()
                nw280_.ctor__((self).f10, (d_1715_v87_).f9)
                d_1731_v98_ = nw280_
                d_1732_v99_: _dafny.Map
                d_1732_v99_ = _dafny.Map({not(True): ((d_1731_v98_).f14) + ((self).f10)})
                d_1732_v99_ = (d_1732_v99_).set(False, (d_1731_v98_).f14)
            elif True:
                r1 = (self).f10
                d_1733_v100_: _dafny.Seq
                d_1733_v100_ = _dafny.SeqWithoutIsStrInference([(self).f10, ((self).f10 if (d_1619_v10_).f15 else (self).f10), (self).f10, (self).f10, len(d_1710_v82_)])
                d_1733_v100_ = (_dafny.SeqWithoutIsStrInference([(self).f10])).set(default__.safeIndex((self).f10, len(_dafny.SeqWithoutIsStrInference([(self).f10]))), (self).f10)
                d_1734_v101_: C2
                nw281_ = C2()
                nw281_.ctor__((0) - ((self).f10), (d_1619_v10_).f15)
                d_1734_v101_ = nw281_
                d_1735_v102_: _dafny.MultiSet
                d_1735_v102_ = _dafny.MultiSet([(d_1715_v87_).f9, (d_1715_v87_).f9])
                d_1736_v103_: bool
                out17_: bool
                out17_ = (d_1619_v10_).m3(((_dafny.MultiSet([(d_1619_v10_).f15])).set((d_1621_v12_)[default__.safeIndex((self).f10, len(d_1621_v12_))], default__.abs((self).f10))).isdisjoint(d_1735_v102_), globalState)
                d_1736_v103_ = out17_
                r0 = (len(d_1621_v12_)) > ((0) - ((d_1734_v101_).f14))
        d_1737_v104_: _dafny.Array
        nw282_ = _dafny.Array(int(0), 17)
        d_1737_v104_ = nw282_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1737_v104_).length(0)):
            d_1738_i12_: int = guard_loop_4_
            if (True) and (((0) <= (d_1738_i12_)) and ((d_1738_i12_) < ((d_1737_v104_).length(0)))):
                (d_1737_v104_)[(d_1738_i12_)] = (d_1738_i12_) + ((self).f10)
        d_1739_v105_: D16
        d_1739_v105_ = D16_DC46(((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1740_i13_ in range(default__.abs(-44))])))) - ((self).f10))
        d_1739_v105_ = d_1739_v105_
        r0 = (d_1619_v10_).f15
        r1 = (self).f10
        r2 = (self).f10
        return r0, r1, r2

    @property
    def f10(self):
        return self._f10
