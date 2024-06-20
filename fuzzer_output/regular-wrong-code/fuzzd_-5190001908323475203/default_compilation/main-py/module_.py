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
        return default__.safeDivisionInt(len((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sppkdbw"))): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')])})) | (_dafny.Map({-658: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), (D23_DC66(True, False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tv"))), _dafny.CodePoint('n'))).cf115, _dafny.CodePoint('y')])}))), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "av"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iysixoa")))))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return ((True) and (False)) and ((False) not in (_dafny.SeqWithoutIsStrInference([True, True, True, not(False)])))

    @staticmethod
    def fm2(p0, p1, globalState):
        return _dafny.Set({(-694) - (827), len((_dafny.Map({True: False})))})

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkry"))])) + ((D29_DC80(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhavv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anqfsyd"))]))).cf137)

    @staticmethod
    def fm4(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(21, 743):
                d_1_v0_: int = compr_0_
                if ((21) <= (d_1_v0_)) and ((d_1_v0_) < (743)):
                    coll0_[(d_1_v0_) + (186)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))
            return _dafny.Map(coll0_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([940 for d_0_i0_ in range(default__.abs(162))]): False})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([559, 10, 822, (_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(iife0_()
        )}))]))).cardinality])).cardinality, 614]): True}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: (D12_DC32(False)).cf55})) for d_2_i1_ in range(default__.abs(-27))]): not(True)}))

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return _dafny.MultiSet([264, ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emmbf")))])).cardinality) + ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u'), _dafny.CodePoint('w')])))), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmkep")))) - (-486), 520, (712) - ((0) - (len(_dafny.SeqWithoutIsStrInference([(0) - (-864)]))))])

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({False: False}))) | ((_dafny.Map({False: True})) | (_dafny.Map({True: True})))

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnrfyvd"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtttsx"))):
            return _dafny.CodePoint('a')
        elif True:
            return _dafny.CodePoint('c')

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(515, -320):
                d_3_v0_: int = compr_1_
                if ((515) <= (d_3_v0_)) and ((d_3_v0_) < (-320)):
                    coll1_[(d_3_v0_) * (17)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_4_i0_ in range(default__.abs(327))]))
            return _dafny.Map(coll1_)
        return D1_DC5(not(False), (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([-930])).cardinality])) + (_dafny.SeqWithoutIsStrInference([260, len(_dafny.Map({False: len(iife1_()
)})), (0) - (len(_dafny.SeqWithoutIsStrInference([False])))])), 86)

    @staticmethod
    def fm18(globalState):
        return _dafny.CodePoint('o')

    @staticmethod
    def fm19(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-636, 620):
                d_5_v0_: int = compr_2_
                if ((-636) <= (d_5_v0_)) and ((d_5_v0_) < (620)):
                    def iife3_():
                        coll3_ = _dafny.Map()
                        compr_3_: int
                        for compr_3_ in _dafny.IntegerRange(759, -371):
                            d_6_v1_: int = compr_3_
                            if ((759) <= (d_6_v1_)) and ((d_6_v1_) < (-371)):
                                coll3_[(d_6_v1_) * (-511)] = False
                        return _dafny.Map(coll3_)
                    coll2_[(d_5_v0_) * ((_dafny.MultiSet([False])).cardinality)] = len(iife3_()
                    )
            return _dafny.Map(coll2_)
        return (_dafny.Map({len(_dafny.Map({True: (0) - (len(_dafny.Map({147: D21_DC58(462)})))})): len(_dafny.Map({647: _dafny.CodePoint('a')}))})) | ((_dafny.Map({277: 106})) | (iife2_()
        ))

    @staticmethod
    def fm20(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dftv")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_7_i0_ in range(default__.abs(-393))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxmd"))))

    @staticmethod
    def fm21(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([(0) - (-961), -525, 796])) + ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhati"))) for d_8_i1_ in range(default__.abs(196))]))) for d_9_i0_ in range(default__.abs(-308))])) + (_dafny.SeqWithoutIsStrInference([521, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltildgwl")))]))])))

    @staticmethod
    def fm22(globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: _dafny.MultiSet
            for compr_4_ in ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(True), False])]) if False else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, (D14_DC40(D6_DC18(D6_DC15(_dafny.CodePoint('l'))), False, True)).cf66])]))).Elements:
                d_10_v0_: _dafny.MultiSet = compr_4_
                if (d_10_v0_) in ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(True), False])]) if False else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, (D14_DC40(D6_DC18(D6_DC15(_dafny.CodePoint('l'))), False, True)).cf66])]))):
                    coll4_ = coll4_.union(_dafny.Set([d_10_v0_]))
            return _dafny.Set(coll4_)
        return iife4_()
        

    @staticmethod
    def fm23(globalState):
        return D1_DC5(not((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_11_i0_ in range(default__.abs(673))]))) != (len(_dafny.SeqWithoutIsStrInference([False])))), (_dafny.SeqWithoutIsStrInference([567])) + (_dafny.SeqWithoutIsStrInference([731, 368])), len(_dafny.SeqWithoutIsStrInference([True, True, True])))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcestcjpt"))): True}))])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, False])), (_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('w')])).cardinality, (0) - (len(_dafny.SeqWithoutIsStrInference([False, False]))), 860]))).intersection(_dafny.MultiSet([572, -2, 633]))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return _dafny.CodePoint('q')

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({True, False})) - (_dafny.Set({not(True)}))) - ((_dafny.Set({not(not(not(False)))}) if True else _dafny.Set({False})))

    @staticmethod
    def fm28(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elpv"))

    @staticmethod
    def fm29(p0, p1, globalState):
        return _dafny.Map({(594) * (437): (-159) + (716)})

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        return (_dafny.Set({_dafny.CodePoint('r')})).intersection(_dafny.Set({_dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('x')}))

    @staticmethod
    def fm33(p0, p1, globalState):
        return _dafny.MultiSet([-809, (len(_dafny.Map({(D23_DC66(True, True, 442, _dafny.CodePoint('f'))).cf112: 721}))) * (335), (len(_dafny.Map({True: False}))) - (881), 889])

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return (_dafny.Map({not(True): True})) | (_dafny.Map({True: not(True)}))

    @staticmethod
    def fm35(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phtihs"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aq")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vw")))

    @staticmethod
    def fm36(p0, p1, globalState):
        return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({453: not(False)}))])): not(False)})), 137])})) | (_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([367, (_dafny.MultiSet([995])).cardinality]))])}))

    @staticmethod
    def fm37(p0, p1, globalState):
        if (False) == (False):
            return _dafny.CodePoint('u')
        elif not(True):
            return _dafny.CodePoint('s')
        elif True:
            return _dafny.CodePoint('e')

    @staticmethod
    def fm38(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Set({not(True)})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, not(True)}) for d_12_i0_ in range(default__.abs(772))]))

    @staticmethod
    def fm39(p0, p1, globalState):
        return D1_DC5(not(True), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False}))}) for d_13_i0_ in range(default__.abs(907))])), (D6_DC17(True, -786, _dafny.CodePoint('d'), _dafny.CodePoint('r'))).cf32, 598]), 259)

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        return D9_DC25()

    @staticmethod
    def fm41(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([not(True)])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False]) for d_14_i0_ in range(default__.abs(-967))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, True, True, False]) for d_15_i1_ in range(default__.abs(502))])))

    @staticmethod
    def fm42(p0, p1, globalState):
        if (_dafny.MultiSet([True, False, not(True)])) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])])):
            if not(True):
                return D6_DC17(False, len(_dafny.Set({len(_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-748, (_dafny.MultiSet([False, True])).cardinality])): True}))})), _dafny.CodePoint('x'), _dafny.CodePoint('m'))
            elif True:
                return D6_DC17(False, 409, _dafny.CodePoint('c'), _dafny.CodePoint('r'))
        elif True:
            return D6_DC17(not(True), -333, _dafny.CodePoint('t'), _dafny.CodePoint('c'))

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return _dafny.Map({True: len((_dafny.Map({800: True})) | (_dafny.Map({993: not(False)})))})

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        if (True if True else False):
            return _dafny.Set({_dafny.CodePoint('p'), _dafny.CodePoint('e')})
        elif True:
            return _dafny.Set({_dafny.CodePoint('t')})

    @staticmethod
    def fm45(p0, p1, globalState):
        return _dafny.Map({(522) + (len(_dafny.SeqWithoutIsStrInference([not(True), False]))): True})

    @staticmethod
    def fm46(p0, globalState):
        return ((_dafny.Map({_dafny.CodePoint('j'): True})) | (_dafny.Map({_dafny.CodePoint('q'): True}))) | (_dafny.Map({_dafny.CodePoint('n'): False}))

    @staticmethod
    def fm47(p0, p1, globalState):
        if (-838) <= (146):
            return D6_DC18(D6_DC17(not(False), 16, _dafny.CodePoint('i'), _dafny.CodePoint('i')))
        elif True:
            return D6_DC18(D6_DC16(True))

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        return D12_DC29((_dafny.MultiSet([239])) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([521]))])))

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        return (_dafny.MultiSet([False, True])) | (_dafny.MultiSet([True]))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([False])) + ((_dafny.SeqWithoutIsStrInference([(D1_DC5(not(False), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gr"))), len(_dafny.Set({(_dafny.MultiSet([len(_dafny.Set({True})), 807])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "re")))})), -637, len(_dafny.Map({not(True): 726})), len(_dafny.SeqWithoutIsStrInference([-441]))]), 323)).cf9])) + (_dafny.SeqWithoutIsStrInference([False, True, False, True, False])))

    @staticmethod
    def fm51(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([D0_DC0(466)])

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return D12_DC32((False if False else False))

    @staticmethod
    def fm53(p0, p1, globalState):
        return D2_DC8(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tb")))

    @staticmethod
    def fm54(globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(14, -272):
                d_16_v0_: int = compr_5_
                if ((14) <= (d_16_v0_)) and ((d_16_v0_) < (-272)):
                    coll5_ = coll5_.union(_dafny.Set([(d_16_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfbs"))))]))
            return _dafny.Set(coll5_)
        return _dafny.SeqWithoutIsStrInference([(iife5_()
        ).intersection(_dafny.Set({545}))])

    @staticmethod
    def fm55(p0, globalState):
        return D8_DC22((-47) - (897), D1_DC3(_dafny.Set({True, False, not(not(False)), not(False)})), D0_DC0(856))

    @staticmethod
    def fm56(globalState):
        return D0_DC0(((_dafny.MultiSet([400, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sa")))])) - (_dafny.MultiSet([626]))).cardinality)

    @staticmethod
    def fm57(p0, p1, p2, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: str
            for compr_6_ in (_dafny.Map({_dafny.CodePoint('a'): _dafny.SeqWithoutIsStrInference([918])})).keys.Elements:
                d_17_v0_: str = compr_6_
                if (d_17_v0_) in (_dafny.Map({_dafny.CodePoint('a'): _dafny.SeqWithoutIsStrInference([918])})):
                    coll6_[d_17_v0_] = _dafny.Map({False: 713})
            return _dafny.Map(coll6_)
        return (_dafny.Map({_dafny.CodePoint('i'): _dafny.Map({False: (0) - (-776)})})) | ((iife6_()
        ) | (_dafny.Map({_dafny.CodePoint('h'): _dafny.Map({False: 804})})))

    @staticmethod
    def fm58(p0, p1, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(-502, 829):
                d_18_v0_: int = compr_7_
                if ((-502) <= (d_18_v0_)) and ((d_18_v0_) < (829)):
                    coll7_[(d_18_v0_) * (358)] = -191
            return _dafny.Map(coll7_)
        return D0_DC1(len((_dafny.Map({(0) - (len(_dafny.Map({False: 432}))): 324})) | (iife7_()
)))

    @staticmethod
    def fm59(globalState):
        return _dafny.SeqWithoutIsStrInference([D15_DC44(D15_DC44(D15_DC44(D15_DC43(len(_dafny.SeqWithoutIsStrInference([-529 for d_19_i0_ in range(default__.abs(260))])))))), D15_DC44(D15_DC44(D15_DC42(_dafny.MultiSet([True, False, not(False)])))), D15_DC44(D15_DC43(len(_dafny.SeqWithoutIsStrInference([686, (D2_DC9(len(_dafny.Map({115: False})), -426, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_20_i1_ in range(default__.abs(103))]), False, _dafny.Map({D0_DC0(len(_dafny.Map({_dafny.CodePoint('r'): False}))): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_21_i2_ in range(default__.abs(739))]))}))).cf17])))), D15_DC44(D15_DC42(_dafny.MultiSet([True]))), D15_DC44(D15_DC44(D15_DC43((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True}))]))).cardinality)))])

    @staticmethod
    def fm60(globalState):
        if not((_dafny.Set({-364, 513})).ispropersubset(_dafny.Set({(0) - (-76)}))):
            return _dafny.Map({_dafny.Set({False, not(True)}): len(_dafny.Set({True}))})
        elif True:
            return _dafny.Map({_dafny.Set({False, not(False), True, True}): 435})

    @staticmethod
    def fm61(p0, p1, globalState):
        def iife8_():
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(319, 45):
                    d_26_v0_: int = compr_10_
                    if ((319) <= (d_26_v0_)) and ((d_26_v0_) < (45)):
                        coll10_[default__.safeDivisionInt(d_26_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_27_i4_ in range(default__.abs(322))])))] = len(_dafny.Set({True}))
                return _dafny.Map(coll10_)
            coll8_ = _dafny.Set()
            def iife9_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(319, 45):
                    d_26_v0_: int = compr_9_
                    if ((319) <= (d_26_v0_)) and ((d_26_v0_) < (45)):
                        coll9_[default__.safeDivisionInt(d_26_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_27_i4_ in range(default__.abs(322))])))] = len(_dafny.Set({True}))
                return _dafny.Map(coll9_)
            compr_8_: int
            for compr_8_ in (_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('k')])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_25_i3_ in range(default__.abs(-841))]))])): len(iife9_()
            )})): not(not(not(False)))})).keys.Elements:
                d_28_v1_: int = compr_8_
                if (d_28_v1_) in (_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('k')])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_25_i3_ in range(default__.abs(-841))]))])): len(iife10_()
                )})): not(not(not(False)))})):
                    coll8_ = coll8_.union(_dafny.Set([default__.safeDivisionInt(d_28_v1_, len(_dafny.SeqWithoutIsStrInference([True])))]))
            return _dafny.Set(coll8_)
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False})) for d_22_i1_ in range(default__.abs(51))]) for d_23_i0_ in range(default__.abs(896))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-910 for d_24_i2_ in range(default__.abs(-203))]), _dafny.SeqWithoutIsStrInference([533, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([-670, 717]))})), len(_dafny.Map({-282: len(iife8_()
        )})), len(_dafny.Map({False: False}))])])))).intersection((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([722, len(_dafny.SeqWithoutIsStrInference([663]))])])) - ((D30_DC82(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([817, 778]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "neutv"))) for d_29_i5_ in range(default__.abs(642))]))]), _dafny.SeqWithoutIsStrInference([574, len(_dafny.SeqWithoutIsStrInference([False]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({5, len(_dafny.SeqWithoutIsStrInference([True]))})) for d_30_i6_ in range(default__.abs(674))])), 197])])))).cf141))

    @staticmethod
    def fm62(p0, p1, globalState):
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: _dafny.Seq
            for compr_11_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_31_i0_ in range(default__.abs(133))])})).Elements:
                d_32_v0_: _dafny.Seq = compr_11_
                if (d_32_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_31_i0_ in range(default__.abs(133))])})):
                    coll11_ = coll11_.union(_dafny.Set([d_32_v0_]))
            return _dafny.Set(coll11_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldysnsn")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tddexouk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yicgsp"))})) | ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iuli")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wobeyci"))})).intersection(iife11_()
        ))

    @staticmethod
    def fm63(p0, p1, p2, globalState):
        return D17_DC48(not((len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_33_i0_ in range(default__.abs(554))])), len(_dafny.Set({not(True), not(False), False}))])).cardinality]))) != (-705)), (True) and (not(True)), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wivlux"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nw"))))

    @staticmethod
    def fm64(globalState):
        if False:
            return D2_DC7(_dafny.Set({-133}))
        elif True:
            return D2_DC7(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jalxahvpr"))])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_34_i0_ in range(default__.abs(7))])), 570}))

    @staticmethod
    def m0(p0, p1, globalState):
        d_35_v0_: bool
        d_35_v0_ = True
        d_36_i0_: int
        d_36_i0_ = 0
        with _dafny.label("0"):
            while d_35_v0_:
                with _dafny.c_label("0"):
                    if (d_36_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_36_i0_ = (d_36_i0_) + (1)
                    d_37_v1_: _dafny.Map
                    d_37_v1_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvrx"))): d_35_v0_})
                    if (d_35_v0_) or ((d_35_v0_ if ((d_37_v1_)[p0] if (p0) in (d_37_v1_) else d_35_v0_) else d_35_v0_)):
                        d_38_v2_: int
                        d_38_v2_ = 519
                        d_38_v2_ = default__.safeModuloInt((d_38_v2_) - (d_38_v2_), p0)
                        d_38_v2_ = default__.safeModuloInt(default__.fm0(globalState), len(d_37_v1_))
                        d_39_v3_: C3
                        nw0_ = C3()
                        nw0_.ctor__(d_38_v2_, d_35_v0_)
                        d_39_v3_ = nw0_
                        d_40_v4_: C5
                        nw1_ = C5()
                        nw1_.ctor__(False)
                        d_40_v4_ = nw1_
                        d_41_v5_: _dafny.Map
                        d_41_v5_ = _dafny.Map({(d_40_v4_).f10: d_40_v4_})
                        d_42_v6_: _dafny.Map
                        d_42_v6_ = _dafny.Map({d_38_v2_: d_40_v4_})
                        d_43_v7_: str
                        d_43_v7_ = _dafny.CodePoint('h')
                        d_44_v8_: _dafny.Array
                        nw2_ = _dafny.Array(None, 19)
                        nw2_[int(0)] = d_40_v4_
                        nw2_[int(1)] = d_40_v4_
                        nw2_[int(2)] = d_40_v4_
                        nw2_[int(3)] = d_40_v4_
                        nw2_[int(4)] = d_40_v4_
                        nw2_[int(5)] = d_40_v4_
                        nw2_[int(6)] = d_40_v4_
                        nw2_[int(7)] = d_40_v4_
                        nw2_[int(8)] = d_40_v4_
                        nw2_[int(9)] = (d_40_v4_ if d_35_v0_ else d_40_v4_)
                        nw2_[int(10)] = d_40_v4_
                        nw2_[int(11)] = d_40_v4_
                        nw2_[int(12)] = d_40_v4_
                        nw2_[int(13)] = d_40_v4_
                        nw2_[int(14)] = ((d_41_v5_)[(d_40_v4_).f10] if ((d_40_v4_).f10) in (d_41_v5_) else d_40_v4_)
                        nw2_[int(15)] = d_40_v4_
                        nw2_[int(16)] = ((d_42_v6_)[(d_39_v3_).fm9(True, p0, d_43_v7_, globalState)] if ((d_39_v3_).fm9(True, p0, d_43_v7_, globalState)) in (d_42_v6_) else d_40_v4_)
                        nw2_[int(17)] = d_40_v4_
                        nw2_[int(18)] = d_40_v4_
                        d_44_v8_ = nw2_
                        index0_ = default__.safeIndex(350, (d_44_v8_).length(0))
                        (d_44_v8_)[index0_] = d_40_v4_
                        d_45_v9_: _dafny.Array
                        nw3_ = _dafny.Array(None, 2)
                        nw3_[int(0)] = (d_40_v4_).f10
                        nw3_[int(1)] = (d_40_v4_).f10
                        d_45_v9_ = nw3_
                        d_46_v10_: _dafny.MultiSet
                        d_46_v10_ = _dafny.MultiSet([(d_40_v4_).f10])
                        d_47_v11_: _dafny.Map
                        d_47_v11_ = _dafny.Map({((d_46_v10_)[d_35_v0_] if (d_35_v0_) in (d_46_v10_) else (d_39_v3_).f14): default__.safeModuloInt(274, p0)})
                        index1_ = default__.safeIndex(350, (d_44_v8_).length(0))
                        rhs0_ = (0) - (p0)
                        rhs1_ = d_40_v4_
                        rhs2_ = p1
                        rhs3_ = d_45_v9_
                        rhs4_ = default__.fm29(default__.safeDivisionInt((d_39_v3_).f14, p1), (d_46_v10_) != (_dafny.MultiSet([(d_40_v4_).f10])), globalState)
                        lhs0_ = d_44_v8_
                        lhs1_ = default__.safeIndex(350, (d_44_v8_).length(0))
                        d_38_v2_ = rhs0_
                        lhs0_[lhs1_] = rhs1_
                        d_38_v2_ = rhs2_
                        d_45_v9_ = rhs3_
                        d_47_v11_ = rhs4_
                        d_35_v0_ = (d_40_v4_).f10
                    elif True:
                        d_48_v12_: T1
                        nw4_ = C3()
                        nw4_.ctor__(len(d_37_v1_), d_35_v0_)
                        d_48_v12_ = nw4_
                        d_48_v12_ = d_48_v12_
                        d_49_v13_: _dafny.Array
                        nw5_ = _dafny.Array(False, 19)
                        d_49_v13_ = nw5_
                        index2_ = default__.safeIndex(890, (d_49_v13_).length(0))
                        (d_49_v13_)[index2_] = not (d_35_v0_) or (False)
                        index3_ = default__.safeIndex(890, (d_49_v13_).length(0))
                        (d_49_v13_)[index3_] = not(d_48_v12_.f7)
                        d_49_v13_ = d_49_v13_
                        d_50_v14_: int
                        d_50_v14_ = 929
                        d_51_v15_: _dafny.Set
                        d_51_v15_ = _dafny.Set({d_48_v12_, d_48_v12_})
                        d_50_v14_ = (p0) + (default__.safeModuloInt(694, len(d_51_v15_)))
                        d_52_v16_: _dafny.Array
                        nw6_ = _dafny.Array(_dafny.Array(None, 0), 4)
                        d_52_v16_ = nw6_
                        index4_ = default__.safeIndex(385, (d_52_v16_).length(0))
                        (d_52_v16_)[index4_] = d_49_v13_
                        index5_ = default__.safeIndex(385, (d_52_v16_).length(0))
                        (d_52_v16_)[index5_] = d_49_v13_
                    d_35_v0_ = (p0) != (default__.safeModuloInt(p1, p1))
                    if d_35_v0_:
                        d_53_v17_: _dafny.MultiSet
                        d_53_v17_ = _dafny.MultiSet([d_35_v0_, d_35_v0_])
                        d_54_v18_: _dafny.Seq
                        d_54_v18_ = _dafny.SeqWithoutIsStrInference([d_53_v17_])
                        d_55_v19_: _dafny.MultiSet
                        d_55_v19_ = _dafny.MultiSet([d_53_v17_, d_53_v17_, (d_54_v18_)[default__.safeIndex(p1, len(d_54_v18_))], d_53_v17_])
                        d_56_v20_: C3
                        nw7_ = C3()
                        nw7_.ctor__(((d_55_v19_)[_dafny.MultiSet([d_35_v0_])] if (_dafny.MultiSet([d_35_v0_])) in (d_55_v19_) else p0), d_35_v0_)
                        d_56_v20_ = nw7_
                        d_57_v21_: int
                        d_57_v21_ = 195
                        d_57_v21_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwk")))
                        d_58_v22_: str
                        d_58_v22_ = _dafny.CodePoint('c')
                        d_58_v22_ = d_58_v22_
                        d_35_v0_ = d_35_v0_
                        d_59_v23_: D1
                        d_59_v23_ = D1_DC3(default__.fm27(d_35_v0_, d_58_v22_, d_35_v0_, d_35_v0_, globalState))
                        d_60_v24_: D1
                        d_60_v24_ = D1_DC6(d_59_v23_)
                        d_60_v24_ = d_60_v24_
                    elif True:
                        d_61_v25_: _dafny.Array
                        nw8_ = _dafny.Array(False, 19)
                        d_61_v25_ = nw8_
                        d_62_v26_: _dafny.Array
                        nw9_ = _dafny.Array(None, 21)
                        d_62_v26_ = nw9_
                        d_63_v27_: _dafny.Seq
                        d_63_v27_ = _dafny.SeqWithoutIsStrInference([d_62_v26_, d_62_v26_])
                        d_64_v28_: _dafny.Seq
                        d_64_v28_ = _dafny.SeqWithoutIsStrInference([(d_63_v27_)[default__.safeIndex(p1, len(d_63_v27_))], d_62_v26_])
                        index6_ = default__.safeIndex(748, (d_61_v25_).length(0))
                        (d_61_v25_)[index6_] = (p0) >= (len(d_64_v28_))
                        index7_ = default__.safeIndex(748, (d_61_v25_).length(0))
                        (d_61_v25_)[index7_] = True
                        d_65_v29_: int
                        d_65_v29_ = 816
                        d_65_v29_ = p0
                        d_65_v29_ = p1
                        d_66_v30_: _dafny.Seq
                        d_66_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knmujqefl"))
                        d_67_v31_: _dafny.Set
                        d_67_v31_ = _dafny.Set({d_61_v25_})
                        d_68_v32_: str
                        d_68_v32_ = _dafny.CodePoint('a')
                        d_69_v33_: C2
                        nw10_ = C2()
                        nw10_.ctor__(d_67_v31_, d_68_v32_, d_35_v0_)
                        d_69_v33_ = nw10_
                        d_70_v34_: _dafny.Seq
                        d_70_v34_ = _dafny.SeqWithoutIsStrInference([d_69_v33_, d_69_v33_, d_69_v33_])
                        d_71_v35_: _dafny.MultiSet
                        d_71_v35_ = _dafny.MultiSet([d_61_v25_, d_61_v25_, d_61_v25_, d_61_v25_])
                        d_72_v36_: T1
                        nw11_ = C6()
                        nw11_.ctor__((d_61_v25_)[default__.safeIndex(748, (d_61_v25_).length(0))], d_35_v0_)
                        d_72_v36_ = nw11_
                        d_73_v37_: _dafny.Set
                        d_73_v37_ = _dafny.Set({d_72_v36_})
                        d_74_v38_: D8
                        d_74_v38_ = D8_DC21(p1, d_67_v31_, len(d_73_v37_))
                        pat_let_tv0_ = d_67_v31_
                        d_75_v39_: _dafny.MultiSet
                        def iife12_(_pat_let0_0):
                            def iife13_(d_76_dt__update__tmp_h0_):
                                def iife14_(_pat_let1_0):
                                    def iife15_(d_77_dt__update_hcf39_h0_):
                                        return D8_DC21((d_76_dt__update__tmp_h0_).cf38, d_77_dt__update_hcf39_h0_, (d_76_dt__update__tmp_h0_).cf40)
                                    return iife15_(_pat_let1_0)
                                return iife14_(pat_let_tv0_)
                            return iife13_(_pat_let0_0)
                        d_75_v39_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([iife12_(d_74_v38_)]))])
                        d_78_v40_: _dafny.Seq
                        d_78_v40_ = _dafny.SeqWithoutIsStrInference([d_35_v0_, d_35_v0_])
                        d_79_v41_: _dafny.Set
                        d_79_v41_ = _dafny.Set({(d_61_v25_)[default__.safeIndex(748, (d_61_v25_).length(0))]})
                        d_80_v42_: _dafny.Map
                        d_80_v42_ = _dafny.Map({True: d_65_v29_})
                        d_81_v43_: _dafny.Set
                        d_81_v43_ = _dafny.Set({d_65_v29_, p1, p0, p0})
                        d_82_v44_: _dafny.MultiSet
                        d_82_v44_ = _dafny.MultiSet([(d_69_v33_).f16])
                        d_83_v45_: _dafny.Array
                        nw12_ = _dafny.Array(None, 28)
                        nw12_[int(0)] = d_65_v29_
                        nw12_[int(1)] = d_65_v29_
                        nw12_[int(2)] = p1
                        nw12_[int(3)] = d_65_v29_
                        nw12_[int(4)] = (d_75_v39_).cardinality
                        nw12_[int(5)] = p1
                        nw12_[int(6)] = len(d_80_v42_)
                        nw12_[int(7)] = p1
                        nw12_[int(8)] = p1
                        nw12_[int(9)] = (0) - (p1)
                        nw12_[int(10)] = d_65_v29_
                        nw12_[int(11)] = len(d_66_v30_)
                        nw12_[int(12)] = p1
                        nw12_[int(13)] = p1
                        nw12_[int(14)] = len(d_81_v43_)
                        nw12_[int(15)] = len(d_66_v30_)
                        nw12_[int(16)] = (d_82_v44_).cardinality
                        nw12_[int(17)] = p0
                        nw12_[int(18)] = len(d_66_v30_)
                        nw12_[int(19)] = d_65_v29_
                        nw12_[int(20)] = p1
                        nw12_[int(21)] = d_65_v29_
                        nw12_[int(22)] = d_65_v29_
                        nw12_[int(23)] = p1
                        nw12_[int(24)] = p0
                        nw12_[int(25)] = p1
                        nw12_[int(26)] = d_65_v29_
                        nw12_[int(27)] = d_65_v29_
                        d_83_v45_ = nw12_
                        d_84_v46_: _dafny.Map
                        d_84_v46_ = _dafny.Map({d_83_v45_: d_81_v43_})
                        d_85_v47_: D13
                        d_85_v47_ = D13_DC34(d_84_v46_)
                        d_86_v48_: _dafny.Seq
                        d_86_v48_ = _dafny.SeqWithoutIsStrInference([d_85_v47_])
                        d_87_v49_: D17
                        d_87_v49_ = D17_DC49(len(d_79_v41_), d_86_v48_, (0) - (len(d_78_v40_)), (d_61_v25_)[default__.safeIndex(748, (d_61_v25_).length(0))])
                        d_88_v50_: D2
                        d_88_v50_ = D2_DC7(d_81_v43_)
                        d_89_v51_: _dafny.Array
                        nw13_ = _dafny.Array(None, 6)
                        nw13_[int(0)] = d_88_v50_
                        nw13_[int(1)] = d_88_v50_
                        nw13_[int(2)] = d_88_v50_
                        nw13_[int(3)] = d_88_v50_
                        nw13_[int(4)] = d_88_v50_
                        nw13_[int(5)] = D2_DC7(d_81_v43_)
                        d_89_v51_ = nw13_
                        d_90_v52_: C4
                        nw14_ = C4()
                        nw14_.ctor__(d_89_v51_, (d_61_v25_)[default__.safeIndex(748, (d_61_v25_).length(0))])
                        d_90_v52_ = nw14_
                        d_91_v53_: _dafny.MultiSet
                        d_91_v53_ = _dafny.MultiSet([_dafny.Map({d_72_v36_.f7: d_90_v52_}), _dafny.Map({d_72_v36_.f7: d_90_v52_})])
                        d_92_v55_: _dafny.Array
                        nw15_ = _dafny.Array(None, 21)
                        nw15_[int(0)] = d_65_v29_
                        nw15_[int(1)] = (0) - ((p0) - (761))
                        nw15_[int(2)] = (p1) + (len(d_66_v30_))
                        nw15_[int(3)] = len(d_70_v34_)
                        nw15_[int(4)] = ((d_71_v35_).cardinality) * (d_65_v29_)
                        nw15_[int(5)] = (0) - (p1)
                        nw15_[int(6)] = (d_75_v39_).cardinality
                        nw15_[int(7)] = (d_75_v39_).cardinality
                        nw15_[int(8)] = len(d_78_v40_)
                        nw15_[int(9)] = default__.fm0(globalState)
                        nw15_[int(10)] = (0) - (d_65_v29_)
                        nw15_[int(11)] = len(d_66_v30_)
                        nw15_[int(12)] = (d_87_v49_).cf82
                        nw15_[int(13)] = ((_dafny.MultiSet(d_78_v40_)).cardinality) * ((d_91_v53_).cardinality)
                        nw15_[int(14)] = d_65_v29_
                        nw15_[int(15)] = (0) - ((0) - ((d_72_v36_).fm9((d_61_v25_)[default__.safeIndex(748, (d_61_v25_).length(0))], p0, (d_69_v33_).f16, globalState)))
                        nw15_[int(16)] = p0
                        nw15_[int(17)] = d_65_v29_
                        nw15_[int(18)] = p1
                        nw15_[int(19)] = d_65_v29_
                        def iife16_():
                            coll12_ = _dafny.Set()
                            compr_12_: int
                            for compr_12_ in _dafny.IntegerRange(744, -968):
                                d_93_v54_: int = compr_12_
                                if ((744) <= (d_93_v54_)) and ((d_93_v54_) < (-968)):
                                    coll12_ = coll12_.union(_dafny.Set([default__.safeModuloInt(d_93_v54_, 450)]))
                            return _dafny.Set(coll12_)
                        nw15_[int(20)] = len((d_81_v43_).intersection(iife16_()
                        ))
                        d_92_v55_ = nw15_
                        index8_ = default__.safeIndex(386, (d_92_v55_).length(0))
                        (d_92_v55_)[index8_] = p0
                        index9_ = default__.safeIndex(386, (d_92_v55_).length(0))
                        index10_ = default__.safeIndex(748, (d_61_v25_).length(0))
                        rhs5_ = (p1) + (default__.safeModuloInt(p1, len(d_79_v41_)))
                        rhs6_ = len(d_79_v41_)
                        rhs7_ = d_35_v0_
                        lhs2_ = d_92_v55_
                        lhs3_ = default__.safeIndex(386, (d_92_v55_).length(0))
                        lhs4_ = d_61_v25_
                        lhs5_ = default__.safeIndex(748, (d_61_v25_).length(0))
                        d_65_v29_ = rhs5_
                        lhs2_[lhs3_] = rhs6_
                        lhs4_[lhs5_] = rhs7_
                        d_94_v56_: _dafny.Map
                        d_94_v56_ = _dafny.Map({True: d_72_v36_.f7})
                        d_94_v56_ = (d_94_v56_).set((d_72_v36_).fm8(p0, globalState), (True) or (False))
                    d_95_v57_: D27
                    d_95_v57_ = D27_DC75(d_37_v1_)
                    d_96_v58_: T1
                    nw16_ = C3()
                    nw16_.ctor__(p0, ((d_95_v57_).cf130) != (d_37_v1_))
                    d_96_v58_ = nw16_
                    d_97_v59_: _dafny.Seq
                    d_97_v59_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_98_v60_: C7
                    nw17_ = C7()
                    nw17_.ctor__(_dafny.Map({True: d_97_v59_}), d_35_v0_)
                    d_98_v60_ = nw17_
                    d_99_v61_: _dafny.Array
                    def lambda0_(d_100_p1_):
                        def lambda1_(d_101_i1_):
                            return (d_101_i1_) * (d_100_p1_)

                        return lambda1_

                    init0_ = lambda0_(p1)
                    nw18_ = _dafny.Array(None, 29)
                    for i0_0_ in range(nw18_.length(0)):
                        nw18_[i0_0_] = init0_(i0_0_)
                    d_99_v61_ = nw18_
                    d_102_v62_: _dafny.Array
                    nw19_ = _dafny.Array(None, 26)
                    nw19_[int(0)] = d_99_v61_
                    nw19_[int(1)] = d_99_v61_
                    nw19_[int(2)] = d_99_v61_
                    nw19_[int(3)] = d_99_v61_
                    nw19_[int(4)] = d_99_v61_
                    nw19_[int(5)] = d_99_v61_
                    nw19_[int(6)] = d_99_v61_
                    nw19_[int(7)] = d_99_v61_
                    nw19_[int(8)] = d_99_v61_
                    nw19_[int(9)] = (d_99_v61_ if d_35_v0_ else d_99_v61_)
                    nw19_[int(10)] = d_99_v61_
                    nw19_[int(11)] = d_99_v61_
                    nw19_[int(12)] = d_99_v61_
                    nw19_[int(13)] = d_99_v61_
                    nw19_[int(14)] = d_99_v61_
                    nw19_[int(15)] = (d_99_v61_ if d_96_v58_.f7 else d_99_v61_)
                    nw19_[int(16)] = d_99_v61_
                    nw19_[int(17)] = d_99_v61_
                    nw19_[int(18)] = d_99_v61_
                    nw19_[int(19)] = d_99_v61_
                    nw19_[int(20)] = d_99_v61_
                    nw19_[int(21)] = d_99_v61_
                    nw19_[int(22)] = d_99_v61_
                    nw19_[int(23)] = d_99_v61_
                    nw19_[int(24)] = d_99_v61_
                    nw19_[int(25)] = d_99_v61_
                    d_102_v62_ = nw19_
                    index11_ = default__.safeIndex(861, (d_102_v62_).length(0))
                    (d_102_v62_)[index11_] = (d_99_v61_ if d_35_v0_ else d_99_v61_)
                    d_103_v63_: D18
                    d_103_v63_ = D18_DC50(d_96_v58_)
                    index12_ = default__.safeIndex(861, (d_102_v62_).length(0))
                    rhs8_ = d_96_v58_
                    rhs9_ = d_98_v60_
                    rhs10_ = d_35_v0_
                    rhs11_ = d_99_v61_
                    rhs12_ = d_103_v63_
                    lhs6_ = d_102_v62_
                    lhs7_ = default__.safeIndex(861, (d_102_v62_).length(0))
                    d_96_v58_ = rhs8_
                    d_98_v60_ = rhs9_
                    d_35_v0_ = rhs10_
                    lhs6_[lhs7_] = rhs11_
                    d_103_v63_ = rhs12_
                    pass
            pass
        d_104_v64_: _dafny.Seq
        d_104_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ed"))
        d_105_v65_: _dafny.Seq
        d_105_v65_ = _dafny.SeqWithoutIsStrInference([d_104_v64_])
        d_106_v66_: _dafny.Map
        d_106_v66_ = _dafny.Map({d_35_v0_: p0})
        d_107_v67_: _dafny.MultiSet
        d_107_v67_ = _dafny.MultiSet([p1])
        d_108_v68_: _dafny.MultiSet
        d_108_v68_ = _dafny.MultiSet([p1, p1, (0) - ((d_107_v67_).cardinality), p0, p0])
        d_109_v69_: _dafny.MultiSet
        d_109_v69_ = _dafny.MultiSet([False])
        d_110_v70_: _dafny.Seq
        d_110_v70_ = _dafny.SeqWithoutIsStrInference([d_109_v69_, d_109_v69_])
        d_111_v71_: _dafny.Array
        nw20_ = _dafny.Array(None, 27)
        nw20_[int(0)] = d_104_v64_
        nw20_[int(1)] = d_104_v64_
        nw20_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxnpy"))
        nw20_[int(3)] = d_104_v64_
        nw20_[int(4)] = default__.fm28(d_35_v0_, p0, globalState)
        nw20_[int(5)] = (d_105_v65_)[default__.safeIndex(p1, len(d_105_v65_))]
        nw20_[int(6)] = d_104_v64_
        nw20_[int(7)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_112_i2_ in range(default__.abs(616))])
        nw20_[int(8)] = default__.fm35(730, p0, globalState)
        nw20_[int(9)] = d_104_v64_
        nw20_[int(10)] = default__.fm28(d_35_v0_, (0) - (len(d_106_v66_)), globalState)
        nw20_[int(11)] = (d_104_v64_) + (d_104_v64_)
        nw20_[int(12)] = (d_104_v64_) + (d_104_v64_)
        nw20_[int(13)] = (d_104_v64_) + (d_104_v64_)
        nw20_[int(14)] = default__.fm28(d_35_v0_, len(default__.fm34((d_107_v67_).cardinality, d_35_v0_, d_35_v0_, ((d_110_v70_)[default__.safeIndex(876, len(d_110_v70_))]).cardinality, globalState)), globalState)
        nw20_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hccoqdanf"))
        nw20_[int(16)] = d_104_v64_
        nw20_[int(17)] = d_104_v64_
        nw20_[int(18)] = d_104_v64_
        nw20_[int(19)] = (d_104_v64_) + (d_104_v64_)
        nw20_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlpfou"))
        nw20_[int(21)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wniyk"))) + (d_104_v64_)
        nw20_[int(22)] = d_104_v64_
        nw20_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vclldfv"))
        nw20_[int(24)] = d_104_v64_
        nw20_[int(25)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vimrxtn"))
        nw20_[int(26)] = d_104_v64_
        d_111_v71_ = nw20_
        d_113_v72_: D28
        d_113_v72_ = D28_DC77(d_111_v71_)
        d_111_v71_ = (d_113_v72_).cf134
        d_114_v73_: int
        d_114_v73_ = 999
        d_115_v74_: _dafny.Map
        d_115_v74_ = _dafny.Map({d_114_v73_: d_114_v73_})
        d_114_v73_ = ((d_115_v74_)[p1] if (p1) in (d_115_v74_) else (p1) - (p1))
        d_116_v75_: _dafny.Seq
        d_116_v75_ = _dafny.SeqWithoutIsStrInference([d_35_v0_])
        d_117_v76_: _dafny.Map
        d_117_v76_ = _dafny.Map({D0_DC1(p0): d_116_v75_})
        d_114_v73_ = len(((d_117_v76_)[D0_DC1(p0)] if (D0_DC1(p0)) in (d_117_v76_) else (d_116_v75_) + (d_116_v75_)))
        d_118_v77_: _dafny.Map
        d_118_v77_ = _dafny.Map({d_104_v64_: d_35_v0_})
        d_35_v0_ = ((d_118_v77_)[(d_104_v64_) + (d_104_v64_)] if ((d_104_v64_) + (d_104_v64_)) in (d_118_v77_) else d_35_v0_)
        d_119_v78_: _dafny.Set
        d_119_v78_ = _dafny.Set({535})
        d_120_v79_: D2
        d_120_v79_ = D2_DC7(d_119_v78_)
        pat_let_tv1_ = d_119_v78_
        pat_let_tv2_ = d_35_v0_
        pat_let_tv3_ = globalState
        d_121_v81_: _dafny.Array
        nw21_ = _dafny.Array(None, 25)
        def iife17_(_pat_let2_0):
            def iife18_(d_122_dt__update__tmp_h1_):
                def iife19_(_pat_let3_0):
                    def iife20_(d_123_dt__update_hcf13_h0_):
                        return D2_DC7(d_123_dt__update_hcf13_h0_)
                    return iife20_(_pat_let3_0)
                return iife19_(pat_let_tv1_)
            return iife18_(_pat_let2_0)
        nw21_[int(0)] = iife17_(d_120_v79_)
        nw21_[int(1)] = d_120_v79_
        nw21_[int(2)] = d_120_v79_
        nw21_[int(3)] = D2_DC7(d_119_v78_)
        def iife21_(_pat_let4_0):
            def iife22_(d_124_dt__update__tmp_h2_):
                def iife23_(_pat_let5_0):
                    def iife24_(d_125_dt__update_hcf13_h1_):
                        return D2_DC7(d_125_dt__update_hcf13_h1_)
                    return iife24_(_pat_let5_0)
                return iife23_(default__.fm2(644, pat_let_tv2_, pat_let_tv3_))
            return iife22_(_pat_let4_0)
        nw21_[int(4)] = iife21_(default__.fm64(globalState))
        nw21_[int(5)] = d_120_v79_
        nw21_[int(6)] = d_120_v79_
        nw21_[int(7)] = d_120_v79_
        nw21_[int(8)] = d_120_v79_
        nw21_[int(9)] = d_120_v79_
        nw21_[int(10)] = d_120_v79_
        nw21_[int(11)] = D2_DC7(_dafny.Set({d_114_v73_}))
        nw21_[int(12)] = d_120_v79_
        nw21_[int(13)] = d_120_v79_
        nw21_[int(14)] = d_120_v79_
        nw21_[int(15)] = d_120_v79_
        nw21_[int(16)] = d_120_v79_
        def iife25_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in _dafny.IntegerRange(75, 633):
                d_126_v80_: int = compr_13_
                if ((75) <= (d_126_v80_)) and ((d_126_v80_) < (633)):
                    coll13_[default__.safeModuloInt(d_126_v80_, p0)] = 779
            return _dafny.Map(coll13_)
        nw21_[int(17)] = D2_DC7(_dafny.Set({d_114_v73_, len(iife25_()
)}))
        nw21_[int(18)] = d_120_v79_
        nw21_[int(19)] = d_120_v79_
        nw21_[int(20)] = d_120_v79_
        nw21_[int(21)] = D2_DC7(_dafny.Set({d_114_v73_}))
        nw21_[int(22)] = d_120_v79_
        nw21_[int(23)] = D2_DC7(d_119_v78_)
        nw21_[int(24)] = d_120_v79_
        d_121_v81_ = nw21_
        d_127_v82_: C4
        nw22_ = C4()
        nw22_.ctor__(d_121_v81_, (d_114_v73_) >= (p1))
        d_127_v82_ = nw22_

    @staticmethod
    def Main(noArgsParameter__):
        d_128_v0_: int
        d_128_v0_ = -231
        d_129_v1_: _dafny.Seq
        d_129_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_128_v0_})])
        d_130_v2_: bool
        d_130_v2_ = True
        d_131_v3_: _dafny.Seq
        d_131_v3_ = _dafny.SeqWithoutIsStrInference([d_130_v2_])
        d_132_v4_: _dafny.MultiSet
        d_132_v4_ = _dafny.MultiSet([True, (d_131_v3_)[default__.safeIndex(d_128_v0_, len(d_131_v3_))]])
        d_133_v5_: _dafny.Array
        def lambda2_(d_134_v3_):
            def lambda3_(d_135_i0_):
                return d_134_v3_

            return lambda3_

        init1_ = lambda2_(d_131_v3_)
        nw23_ = _dafny.Array(None, 29)
        for i0_1_ in range(nw23_.length(0)):
            nw23_[i0_1_] = init1_(i0_1_)
        d_133_v5_ = nw23_
        d_136_globalState_: GlobalState
        nw24_ = GlobalState()
        nw24_.ctor__(False, False, d_129_v1_, d_132_v4_, 290, d_133_v5_)
        d_136_globalState_ = nw24_
        if d_130_v2_:
            d_137_v6_: _dafny.Map
            d_137_v6_ = _dafny.Map({d_130_v2_: default__.fm0(d_136_globalState_)})
            d_137_v6_ = (d_137_v6_).set(d_130_v2_, (d_128_v0_ if d_130_v2_ else d_128_v0_))
            d_138_v7_: _dafny.Array
            nw25_ = _dafny.Array(_dafny.Map({}), 24)
            d_138_v7_ = nw25_
            d_138_v7_ = d_138_v7_
            d_137_v6_ = (d_137_v6_).set((d_130_v2_) or (d_130_v2_), d_128_v0_)
            d_139_v8_: _dafny.Seq
            d_139_v8_ = _dafny.SeqWithoutIsStrInference([d_128_v0_, d_128_v0_])
            d_128_v0_ = (len(d_139_v8_)) * (d_128_v0_)
            if False:
                d_128_v0_ = (d_128_v0_) - (default__.safeModuloInt(d_128_v0_, d_128_v0_))
                d_140_v9_: _dafny.Array
                nw26_ = _dafny.Array(int(0), 8)
                d_140_v9_ = nw26_
                d_141_v10_: _dafny.Set
                d_141_v10_ = _dafny.Set({d_140_v9_, d_140_v9_})
                d_142_v11_: D0
                d_142_v11_ = D0_DC2(d_130_v2_, d_128_v0_, d_141_v10_, 558)
                pat_let_tv4_ = d_128_v0_
                pat_let_tv5_ = d_141_v10_
                def iife26_(_pat_let6_0):
                    def iife27_(d_143_dt__update__tmp_h0_):
                        def iife28_(_pat_let7_0):
                            def iife29_(d_144_dt__update_hcf5_h0_):
                                def iife30_(_pat_let8_0):
                                    def iife31_(d_145_dt__update_hcf4_h0_):
                                        return D0_DC2((d_143_dt__update__tmp_h0_).cf2, (d_143_dt__update__tmp_h0_).cf3, d_145_dt__update_hcf4_h0_, d_144_dt__update_hcf5_h0_)
                                    return iife31_(_pat_let8_0)
                                return iife30_(pat_let_tv5_)
                            return iife29_(_pat_let7_0)
                        return iife28_(pat_let_tv4_)
                    return iife27_(_pat_let6_0)
                d_128_v0_ = (iife26_(d_142_v11_)).cf3
                d_130_v2_ = d_130_v2_
                d_128_v0_ = default__.safeDivisionInt(d_128_v0_, d_128_v0_)
                d_146_v12_: _dafny.Seq
                d_146_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
                d_147_v13_: _dafny.Set
                d_147_v13_ = _dafny.Set({(d_131_v3_)[default__.safeIndex(d_128_v0_, len(d_131_v3_))]})
                d_148_v14_: str
                d_148_v14_ = _dafny.CodePoint('w')
                d_149_v15_: _dafny.Set
                d_149_v15_ = _dafny.Set({d_148_v14_})
                d_150_v16_: _dafny.Set
                d_150_v16_ = _dafny.Set({d_128_v0_, d_128_v0_})
                d_151_v17_: _dafny.Map
                d_151_v17_ = _dafny.Map({d_150_v16_: d_130_v2_})
                d_152_v18_: _dafny.Array
                nw27_ = _dafny.Array(None, 20)
                nw27_[int(0)] = d_130_v2_
                nw27_[int(1)] = d_130_v2_
                nw27_[int(2)] = (d_128_v0_) >= (d_128_v0_)
                nw27_[int(3)] = (d_146_v12_) != (d_146_v12_)
                nw27_[int(4)] = not((d_130_v2_) and (d_130_v2_))
                nw27_[int(5)] = d_130_v2_
                nw27_[int(6)] = (_dafny.Set({d_130_v2_})) == (d_147_v13_)
                nw27_[int(7)] = True
                nw27_[int(8)] = (d_147_v13_).ispropersubset(d_147_v13_)
                nw27_[int(9)] = d_130_v2_
                nw27_[int(10)] = not(default__.fm1(_dafny.SeqWithoutIsStrInference([D0_DC0(d_128_v0_)]), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_153_i1_ in range(default__.abs(973))])), len(d_137_v6_), d_136_globalState_))
                nw27_[int(11)] = (d_130_v2_) in (d_132_v4_)
                nw27_[int(12)] = d_130_v2_
                nw27_[int(13)] = (d_128_v0_) > (d_128_v0_)
                nw27_[int(14)] = d_130_v2_
                nw27_[int(15)] = (_dafny.Set({d_130_v2_, False})).ispropersubset(d_147_v13_)
                nw27_[int(16)] = (d_149_v15_).ispropersubset(d_149_v15_)
                nw27_[int(17)] = (d_130_v2_) and (d_130_v2_)
                nw27_[int(18)] = not((default__.fm2(d_128_v0_, d_130_v2_, d_136_globalState_)) in (d_151_v17_))
                nw27_[int(19)] = d_130_v2_
                d_152_v18_ = nw27_
                d_152_v18_ = d_152_v18_
            elif True:
                d_154_v19_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.CodePoint('D'), 11)
                d_154_v19_ = nw28_
                d_155_v20_: str
                d_155_v20_ = _dafny.CodePoint('i')
                index13_ = default__.safeIndex(404, (d_154_v19_).length(0))
                (d_154_v19_)[index13_] = d_155_v20_
                index14_ = default__.safeIndex(404, (d_154_v19_).length(0))
                (d_154_v19_)[index14_] = d_155_v20_
                d_156_v21_: _dafny.Array
                nw29_ = _dafny.Array(_dafny.MultiSet({}), 7)
                d_156_v21_ = nw29_
                d_157_v22_: _dafny.MultiSet
                d_157_v22_ = _dafny.MultiSet([340])
                index15_ = default__.safeIndex(967, (d_156_v21_).length(0))
                (d_156_v21_)[index15_] = d_157_v22_
                index16_ = default__.safeIndex(967, (d_156_v21_).length(0))
                rhs13_ = d_130_v2_
                rhs14_ = (0) - (d_128_v0_)
                rhs15_ = d_157_v22_
                lhs8_ = d_156_v21_
                lhs9_ = default__.safeIndex(967, (d_156_v21_).length(0))
                d_130_v2_ = rhs13_
                d_128_v0_ = rhs14_
                lhs8_[lhs9_] = rhs15_
                d_158_v23_: _dafny.Seq
                d_158_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdayhkro"))
                d_158_v23_ = (d_158_v23_) + (_dafny.SeqWithoutIsStrInference([(d_154_v19_)[default__.safeIndex(404, (d_154_v19_).length(0))] for d_159_i2_ in range(default__.abs(702))]))
                default__.m0(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvuwlpyo"))) + (d_158_v23_)), d_128_v0_, d_136_globalState_)
                default__.m0((0) - (default__.fm0(d_136_globalState_)), d_128_v0_, d_136_globalState_)
        elif True:
            d_160_v24_: D0
            d_160_v24_ = D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqkta"))))
            d_160_v24_ = d_160_v24_
            d_161_v25_: _dafny.Seq
            d_161_v25_ = _dafny.SeqWithoutIsStrInference([d_160_v24_])
            d_162_v26_: _dafny.Map
            d_162_v26_ = _dafny.Map({d_130_v2_: d_128_v0_})
            d_163_v27_: _dafny.Array
            nw30_ = _dafny.Array(None, 28)
            nw30_[int(0)] = d_130_v2_
            nw30_[int(1)] = not(d_130_v2_)
            nw30_[int(2)] = default__.fm1(d_161_v25_, d_128_v0_, d_128_v0_, d_136_globalState_)
            nw30_[int(3)] = d_130_v2_
            nw30_[int(4)] = d_130_v2_
            nw30_[int(5)] = d_130_v2_
            nw30_[int(6)] = d_130_v2_
            nw30_[int(7)] = d_130_v2_
            nw30_[int(8)] = d_130_v2_
            nw30_[int(9)] = d_130_v2_
            nw30_[int(10)] = d_130_v2_
            nw30_[int(11)] = not(d_130_v2_)
            nw30_[int(12)] = d_130_v2_
            nw30_[int(13)] = (d_131_v3_)[default__.safeIndex(d_128_v0_, len(d_131_v3_))]
            nw30_[int(14)] = d_130_v2_
            nw30_[int(15)] = d_130_v2_
            nw30_[int(16)] = d_130_v2_
            nw30_[int(17)] = d_130_v2_
            nw30_[int(18)] = d_130_v2_
            nw30_[int(19)] = default__.fm1(d_161_v25_, len(d_162_v26_), d_128_v0_, d_136_globalState_)
            nw30_[int(20)] = d_130_v2_
            nw30_[int(21)] = d_130_v2_
            nw30_[int(22)] = d_130_v2_
            nw30_[int(23)] = False
            nw30_[int(24)] = d_130_v2_
            nw30_[int(25)] = d_130_v2_
            nw30_[int(26)] = d_130_v2_
            nw30_[int(27)] = d_130_v2_
            d_163_v27_ = nw30_
            d_164_v28_: _dafny.Seq
            d_164_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mre"))
            d_165_v29_: _dafny.Map
            d_165_v29_ = _dafny.Map({_dafny.Map({d_163_v27_: d_164_v28_}): d_128_v0_})
            d_166_v30_: _dafny.Map
            d_166_v30_ = _dafny.Map({d_163_v27_: d_164_v28_})
            d_165_v29_ = (d_165_v29_).set(d_166_v30_, d_128_v0_)
            d_167_v31_: _dafny.Array
            nw31_ = _dafny.Array(int(0), 26)
            d_167_v31_ = nw31_
            d_168_v32_: _dafny.MultiSet
            d_168_v32_ = _dafny.MultiSet([d_167_v31_])
            d_169_v33_: _dafny.Map
            d_169_v33_ = _dafny.Map({d_128_v0_: ((d_168_v32_)[d_167_v31_] if (d_167_v31_) in (d_168_v32_) else d_128_v0_)})
            default__.m0(d_128_v0_, len(d_169_v33_), d_136_globalState_)
            d_170_v34_: _dafny.Map
            d_170_v34_ = _dafny.Map({default__.fm3(d_130_v2_, (d_132_v4_).cardinality, d_128_v0_, len(d_164_v28_), d_136_globalState_): d_130_v2_})
            d_171_v35_: str
            d_171_v35_ = _dafny.CodePoint('t')
            d_170_v34_ = (d_170_v34_).set(_dafny.SeqWithoutIsStrInference([(d_164_v28_).set(default__.safeIndex(d_128_v0_, len(d_164_v28_)), d_171_v35_), _dafny.SeqWithoutIsStrInference([d_171_v35_ for d_172_i3_ in range(default__.abs(-651))])]), d_130_v2_)
            nw32_ = _dafny.Array(None, 9)
            nw32_[int(0)] = d_130_v2_
            nw32_[int(1)] = (d_132_v4_).ispropersubset(d_132_v4_)
            nw32_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_128_v0_ for d_173_i4_ in range(default__.abs(222))])) < (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_171_v35_ for d_174_i5_ in range(default__.abs(958))]))]))
            nw32_[int(3)] = (d_128_v0_) <= (-697)
            nw32_[int(4)] = d_130_v2_
            nw32_[int(5)] = not(d_130_v2_)
            nw32_[int(6)] = d_130_v2_
            nw32_[int(7)] = (d_128_v0_) < (len(d_164_v28_))
            nw32_[int(8)] = d_130_v2_
            d_163_v27_ = nw32_
        d_175_v36_: _dafny.MultiSet
        d_175_v36_ = _dafny.MultiSet([d_128_v0_, d_128_v0_, 826])
        hi0_ = (d_175_v36_).cardinality
        for d_176_i6_ in range(d_128_v0_, hi0_):
            d_175_v36_ = d_175_v36_
            if d_130_v2_:
                d_177_v37_: _dafny.Array
                def lambda4_(d_178_v36_, d_179_v3_, d_180_v0_):
                    def lambda5_(d_181_i7_):
                        return (_dafny.MultiSet([len(d_179_v3_), d_180_v0_])).ispropersubset(d_178_v36_)

                    return lambda5_

                init2_ = lambda4_(d_175_v36_, d_131_v3_, d_128_v0_)
                nw33_ = _dafny.Array(None, 7)
                for i0_2_ in range(nw33_.length(0)):
                    nw33_[i0_2_] = init2_(i0_2_)
                d_177_v37_ = nw33_
                index17_ = default__.safeIndex(505, (d_177_v37_).length(0))
                (d_177_v37_)[index17_] = d_130_v2_
                index18_ = default__.safeIndex(505, (d_177_v37_).length(0))
                (d_177_v37_)[index18_] = d_130_v2_
                index19_ = default__.safeIndex(505, (d_177_v37_).length(0))
                (d_177_v37_)[index19_] = not((d_177_v37_)[default__.safeIndex(505, (d_177_v37_).length(0))])
                d_182_v38_: _dafny.Seq
                d_182_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lragd"))
                d_182_v38_ = (d_182_v38_) + (d_182_v38_)
                d_130_v2_ = (d_177_v37_)[default__.safeIndex(505, (d_177_v37_).length(0))]
                def iife32_():
                    coll14_ = _dafny.Map()
                    compr_14_: _dafny.Seq
                    for compr_14_ in (default__.fm4(d_176_i6_, d_136_globalState_)).keys.Elements:
                        d_183_v39_: _dafny.Seq = compr_14_
                        if (d_183_v39_) in (default__.fm4(d_176_i6_, d_136_globalState_)):
                            coll14_[d_183_v39_] = (d_177_v37_)[default__.safeIndex(505, (d_177_v37_).length(0))]
                    return _dafny.Map(coll14_)
                default__.m0(d_176_i6_, (len(iife32_()
                )) * (650), d_136_globalState_)
            elif True:
                d_184_v40_: D0
                d_184_v40_ = D0_DC0(d_176_i6_)
                d_185_v41_: _dafny.Seq
                def iife33_(_pat_let9_0):
                    def iife34_(d_186_dt__update__tmp_h1_):
                        def iife35_(_pat_let10_0):
                            def iife36_(d_187_dt__update_hcf0_h0_):
                                return D0_DC0(d_187_dt__update_hcf0_h0_)
                            return iife36_(_pat_let10_0)
                        return iife35_(d_176_i6_)
                    return iife34_(_pat_let9_0)
                d_185_v41_ = _dafny.SeqWithoutIsStrInference([d_184_v40_, iife33_(D0_DC0(d_176_i6_)), d_184_v40_, d_184_v40_])
                d_188_v42_: _dafny.Seq
                d_188_v42_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_128_v0_ for d_189_i8_ in range(default__.abs(633))])), d_176_i6_])
                d_130_v2_ = default__.fm1((d_185_v41_) + (d_185_v41_), d_176_i6_, len((d_188_v42_) + (d_188_v42_)), d_136_globalState_)
                d_190_v43_: _dafny.Seq
                d_190_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mb"))
                d_190_v43_ = d_190_v43_
                default__.m0(d_128_v0_, (d_188_v42_)[default__.safeIndex(d_128_v0_, len(d_188_v42_))], d_136_globalState_)
                d_128_v0_ = d_128_v0_
                d_191_v44_: _dafny.Map
                d_191_v44_ = _dafny.Map({d_130_v2_: d_130_v2_})
                d_192_v45_: _dafny.MultiSet
                d_192_v45_ = _dafny.MultiSet([d_191_v44_, d_191_v44_, d_191_v44_, d_191_v44_])
                d_130_v2_ = (d_192_v45_).ispropersubset(_dafny.MultiSet([d_191_v44_, d_191_v44_]))
            d_193_v46_: _dafny.Seq
            d_193_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxbij"))
            d_193_v46_ = (d_193_v46_) + (d_193_v46_)
            if True:
                d_194_v47_: D0
                d_194_v47_ = D0_DC0(d_128_v0_)
                d_195_v48_: _dafny.Array
                nw34_ = _dafny.Array(None, 6)
                nw34_[int(0)] = default__.fm1(_dafny.SeqWithoutIsStrInference([d_194_v47_, d_194_v47_]), d_128_v0_, d_128_v0_, d_136_globalState_)
                nw34_[int(1)] = (-499) != (d_128_v0_)
                nw34_[int(2)] = d_130_v2_
                nw34_[int(3)] = not (d_130_v2_) or (d_130_v2_)
                nw34_[int(4)] = not(d_130_v2_)
                nw34_[int(5)] = (False if not(d_130_v2_) else not(d_130_v2_))
                d_195_v48_ = nw34_
                index20_ = default__.safeIndex(152, (d_195_v48_).length(0))
                (d_195_v48_)[index20_] = d_130_v2_
                index21_ = default__.safeIndex(152, (d_195_v48_).length(0))
                (d_195_v48_)[index21_] = (d_128_v0_) <= (d_176_i6_)
                d_128_v0_ = d_128_v0_
                d_196_v49_: _dafny.Array
                def lambda6_(d_197_v48_):
                    def lambda7_(d_198_i9_):
                        return (d_198_i9_) * (len(_dafny.Map({(d_197_v48_)[default__.safeIndex(152, (d_197_v48_).length(0))]: True})))

                    return lambda7_

                init3_ = lambda6_(d_195_v48_)
                nw35_ = _dafny.Array(None, 12)
                for i0_3_ in range(nw35_.length(0)):
                    nw35_[i0_3_] = init3_(i0_3_)
                d_196_v49_ = nw35_
                index22_ = default__.safeIndex(383, (d_196_v49_).length(0))
                (d_196_v49_)[index22_] = 543
                index23_ = default__.safeIndex(383, (d_196_v49_).length(0))
                (d_196_v49_)[index23_] = default__.safeDivisionInt((0) - (d_176_i6_), (d_176_i6_) + (d_128_v0_))
                d_199_v50_: _dafny.Seq
                d_199_v50_ = _dafny.SeqWithoutIsStrInference([d_196_v49_])
                default__.m0(len(d_199_v50_), (len(d_193_v46_)) + (d_176_i6_), d_136_globalState_)
                d_200_v51_: _dafny.Map
                d_200_v51_ = _dafny.Map({d_130_v2_: d_176_i6_})
                d_194_v47_ = D0_DC0(len(d_200_v51_))
            elif True:
                d_201_v52_: D0
                d_201_v52_ = D0_DC0((0) - (len(_dafny.Map({not(d_130_v2_): d_130_v2_}))))
                d_202_v53_: D0
                d_202_v53_ = D0_DC1(d_128_v0_)
                d_203_v54_: _dafny.Set
                d_203_v54_ = _dafny.Set({d_176_i6_, d_128_v0_})
                d_204_v55_: _dafny.Array
                nw36_ = _dafny.Array(None, 13)
                nw36_[int(0)] = (d_201_v52_).cf0
                nw36_[int(1)] = len(d_193_v46_)
                nw36_[int(2)] = d_176_i6_
                nw36_[int(3)] = d_176_i6_
                nw36_[int(4)] = (d_202_v53_).cf1
                nw36_[int(5)] = 49
                nw36_[int(6)] = len(d_203_v54_)
                nw36_[int(7)] = d_128_v0_
                nw36_[int(8)] = d_128_v0_
                nw36_[int(9)] = d_128_v0_
                nw36_[int(10)] = d_176_i6_
                nw36_[int(11)] = 895
                nw36_[int(12)] = d_176_i6_
                d_204_v55_ = nw36_
                d_205_v56_: _dafny.Map
                d_205_v56_ = _dafny.Map({d_176_i6_: d_204_v55_})
                d_205_v56_ = (d_205_v56_).set(d_128_v0_, d_204_v55_)
                d_206_v57_: _dafny.Array
                nw37_ = _dafny.Array(None, 5)
                nw37_[int(0)] = d_130_v2_
                nw37_[int(1)] = d_130_v2_
                nw37_[int(2)] = d_130_v2_
                nw37_[int(3)] = d_130_v2_
                nw37_[int(4)] = d_130_v2_
                d_206_v57_ = nw37_
                d_207_v58_: _dafny.Map
                d_207_v58_ = _dafny.Map({d_206_v57_: d_130_v2_})
                d_128_v0_ = len(d_207_v58_)
                index24_ = default__.safeIndex(578, (d_206_v57_).length(0))
                (d_206_v57_)[index24_] = (d_176_i6_) == (d_128_v0_)
                index25_ = default__.safeIndex(578, (d_206_v57_).length(0))
                (d_206_v57_)[index25_] = (d_130_v2_) == (d_130_v2_)
                d_208_v59_: _dafny.Array
                nw38_ = _dafny.Array(D0.default()(), 4)
                d_208_v59_ = nw38_
                d_209_v60_: _dafny.Seq
                d_209_v60_ = _dafny.SeqWithoutIsStrInference([d_208_v59_])
                d_209_v60_ = (d_209_v60_).set(default__.safeIndex(d_128_v0_, len(d_209_v60_)), d_208_v59_)
                index26_ = default__.safeIndex(578, (d_206_v57_).length(0))
                (d_206_v57_)[index26_] = (d_206_v57_)[default__.safeIndex(578, (d_206_v57_).length(0))]
        d_130_v2_ = d_130_v2_
        d_210_v61_: D0
        d_210_v61_ = D0_DC0(d_128_v0_)
        d_211_v62_: _dafny.Map
        d_211_v62_ = _dafny.Map({d_210_v61_: (d_130_v2_ if d_130_v2_ else d_130_v2_)})
        d_212_v63_: _dafny.MultiSet
        d_212_v63_ = _dafny.MultiSet([d_175_v36_])
        if ((d_211_v62_)[d_210_v61_] if (d_210_v61_) in (d_211_v62_) else (d_212_v63_).ispropersubset(d_212_v63_)):
            d_213_v64_: _dafny.Array
            nw39_ = _dafny.Array(None, 2)
            nw39_[int(0)] = not(d_130_v2_)
            nw39_[int(1)] = (d_132_v4_).ispropersubset(d_132_v4_)
            d_213_v64_ = nw39_
            index27_ = default__.safeIndex(870, (d_213_v64_).length(0))
            (d_213_v64_)[index27_] = d_130_v2_
            index28_ = default__.safeIndex(870, (d_213_v64_).length(0))
            (d_213_v64_)[index28_] = True
            d_214_v65_: _dafny.Map
            d_214_v65_ = _dafny.Map({(d_132_v4_).cardinality: d_128_v0_})
            d_214_v65_ = (d_214_v65_).set(d_128_v0_, (d_128_v0_) * (default__.fm0(d_136_globalState_)))
            d_215_v66_: _dafny.Map
            d_215_v66_ = _dafny.Map({not (d_130_v2_) or (default__.fm1(_dafny.SeqWithoutIsStrInference([d_210_v61_]), d_128_v0_, d_128_v0_, d_136_globalState_)): d_132_v4_})
            d_216_v67_: _dafny.Array
            def lambda8_(d_217_v0_):
                def lambda9_(d_218_i10_):
                    return (d_218_i10_) + (d_217_v0_)

                return lambda9_

            init4_ = lambda8_(d_128_v0_)
            nw40_ = _dafny.Array(None, 26)
            for i0_4_ in range(nw40_.length(0)):
                nw40_[i0_4_] = init4_(i0_4_)
            d_216_v67_ = nw40_
            d_219_v68_: _dafny.Set
            d_219_v68_ = _dafny.Set({d_216_v67_, d_216_v67_})
            d_220_v69_: D0
            d_220_v69_ = D0_DC2(True, d_128_v0_, d_219_v68_, d_128_v0_)
            d_221_v70_: _dafny.Seq
            d_221_v70_ = _dafny.SeqWithoutIsStrInference([d_219_v68_])
            d_222_v71_: D0
            d_222_v71_ = D0_DC2((d_220_v69_).cf2, d_128_v0_, (d_221_v70_)[default__.safeIndex(d_128_v0_, len(d_221_v70_))], d_128_v0_)
            d_215_v66_ = (d_215_v66_).set(not((d_222_v71_).cf2), d_132_v4_)
            d_223_v72_: _dafny.Set
            d_223_v72_ = _dafny.Set({True})
            d_224_v73_: D1
            d_224_v73_ = D1_DC3(d_223_v72_)
            d_223_v72_ = (d_224_v73_).cf6
            d_225_v74_: _dafny.Set
            d_225_v74_ = _dafny.Set({d_213_v64_})
            d_226_v75_: _dafny.Map
            d_226_v75_ = _dafny.Map({d_128_v0_: d_130_v2_})
            d_227_v76_: str
            d_227_v76_ = _dafny.CodePoint('g')
            d_228_v77_: C2
            nw41_ = C2()
            nw41_.ctor__((D8_DC21(d_128_v0_, d_225_v74_, len(d_226_v75_))).cf39, d_227_v76_, (d_213_v64_)[default__.safeIndex(870, (d_213_v64_).length(0))])
            d_228_v77_ = nw41_
        elif True:
            d_229_v78_: _dafny.Array
            def lambda10_(d_230_v2_):
                def lambda11_(d_231_i11_):
                    return d_230_v2_

                return lambda11_

            init5_ = lambda10_(d_130_v2_)
            nw42_ = _dafny.Array(None, 13)
            for i0_5_ in range(nw42_.length(0)):
                nw42_[i0_5_] = init5_(i0_5_)
            d_229_v78_ = nw42_
            index29_ = default__.safeIndex(962, (d_229_v78_).length(0))
            (d_229_v78_)[index29_] = d_130_v2_
            d_232_v79_: _dafny.Seq
            d_232_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkcb"))
            d_233_v80_: _dafny.Seq
            d_233_v80_ = _dafny.SeqWithoutIsStrInference([d_232_v79_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_234_i12_ in range(default__.abs(494))])])
            index30_ = default__.safeIndex(962, (d_229_v78_).length(0))
            (d_229_v78_)[index30_] = ((d_233_v80_)[default__.safeIndex(d_128_v0_, len(d_233_v80_))]) <= (d_232_v79_)
            default__.m0(d_128_v0_, d_128_v0_, d_136_globalState_)
            d_235_v81_: _dafny.Map
            d_235_v81_ = _dafny.Map({d_130_v2_: d_130_v2_})
            d_236_v82_: _dafny.Seq
            d_236_v82_ = _dafny.SeqWithoutIsStrInference([D0_DC0(d_128_v0_)])
            d_235_v81_ = (d_235_v81_).set(d_130_v2_, not (default__.fm1(d_236_v82_, d_128_v0_, (0) - (d_128_v0_), d_136_globalState_)) or ((d_229_v78_)[default__.safeIndex(962, (d_229_v78_).length(0))]))
            d_237_v83_: _dafny.Map
            d_237_v83_ = _dafny.Map({(d_229_v78_)[default__.safeIndex(962, (d_229_v78_).length(0))]: d_128_v0_})
            d_238_v84_: _dafny.Array
            def lambda12_(d_239_v0_):
                def lambda13_(d_240_i13_):
                    return default__.safeDivisionInt(d_240_i13_, d_239_v0_)

                return lambda13_

            init6_ = lambda12_(d_128_v0_)
            nw43_ = _dafny.Array(None, 19)
            for i0_6_ in range(nw43_.length(0)):
                nw43_[i0_6_] = init6_(i0_6_)
            d_238_v84_ = nw43_
            d_241_v85_: D13
            d_241_v85_ = D13_DC34(_dafny.Map({d_238_v84_: default__.fm2(540, d_130_v2_, d_136_globalState_)}))
            d_242_v86_: _dafny.Seq
            d_242_v86_ = _dafny.SeqWithoutIsStrInference([d_241_v85_])
            d_243_v87_: D17
            d_243_v87_ = D17_DC49(len(d_237_v83_), d_242_v86_, len(d_237_v83_), d_130_v2_)
            d_243_v87_ = d_243_v87_
            d_244_v88_: D6
            d_244_v88_ = D6_DC16((d_229_v78_)[default__.safeIndex(962, (d_229_v78_).length(0))])
            d_245_v89_: _dafny.Map
            d_245_v89_ = _dafny.Map({(d_210_v61_).cf0: (d_244_v88_).cf30})
            d_245_v89_ = (d_245_v89_).set(d_128_v0_, (588) <= (d_128_v0_))
        d_246_v90_: _dafny.Seq
        d_246_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpq"))
        d_246_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mg"))
        if d_130_v2_:
            d_128_v0_ = d_128_v0_
            d_247_v91_: C1
            nw44_ = C1()
            nw44_.ctor__(d_130_v2_)
            d_247_v91_ = nw44_
            d_248_v92_: _dafny.Seq
            d_248_v92_ = _dafny.SeqWithoutIsStrInference([d_247_v91_])
            d_249_v93_: _dafny.Array
            nw45_ = _dafny.Array(None, 27)
            nw45_[int(0)] = d_247_v91_
            nw45_[int(1)] = d_247_v91_
            nw45_[int(2)] = d_247_v91_
            nw45_[int(3)] = d_247_v91_
            nw45_[int(4)] = d_247_v91_
            nw45_[int(5)] = d_247_v91_
            nw45_[int(6)] = d_247_v91_
            nw45_[int(7)] = d_247_v91_
            nw45_[int(8)] = d_247_v91_
            nw45_[int(9)] = d_247_v91_
            nw45_[int(10)] = d_247_v91_
            nw45_[int(11)] = d_247_v91_
            nw45_[int(12)] = d_247_v91_
            nw45_[int(13)] = d_247_v91_
            nw45_[int(14)] = d_247_v91_
            nw45_[int(15)] = d_247_v91_
            nw45_[int(16)] = (d_248_v92_)[default__.safeIndex(d_128_v0_, len(d_248_v92_))]
            nw45_[int(17)] = d_247_v91_
            nw45_[int(18)] = (d_248_v92_)[default__.safeIndex(d_128_v0_, len(d_248_v92_))]
            nw45_[int(19)] = d_247_v91_
            nw45_[int(20)] = d_247_v91_
            nw45_[int(21)] = d_247_v91_
            nw45_[int(22)] = d_247_v91_
            nw45_[int(23)] = d_247_v91_
            nw45_[int(24)] = d_247_v91_
            nw45_[int(25)] = d_247_v91_
            nw45_[int(26)] = d_247_v91_
            d_249_v93_ = nw45_
            d_250_v94_: _dafny.Map
            d_250_v94_ = _dafny.Map({d_249_v93_: d_130_v2_})
            d_251_v95_: _dafny.Seq
            d_251_v95_ = _dafny.SeqWithoutIsStrInference([d_249_v93_])
            d_252_v96_: _dafny.Map
            d_252_v96_ = _dafny.Map({d_210_v61_: d_128_v0_})
            d_253_v97_: D2
            d_253_v97_ = D2_DC9(d_128_v0_, d_128_v0_, d_246_v90_, d_130_v2_, d_252_v96_)
            d_254_v98_: _dafny.Map
            d_254_v98_ = _dafny.Map({d_253_v97_: d_130_v2_})
            d_255_v99_: _dafny.Array
            nw46_ = _dafny.Array(None, 25)
            nw46_[int(0)] = d_250_v94_
            nw46_[int(1)] = d_250_v94_
            nw46_[int(2)] = d_250_v94_
            nw46_[int(3)] = d_250_v94_
            nw46_[int(4)] = (_dafny.Map({d_249_v93_: d_130_v2_})).set((d_251_v95_)[default__.safeIndex(d_128_v0_, len(d_251_v95_))], d_130_v2_)
            nw46_[int(5)] = _dafny.Map({d_249_v93_: ((d_254_v98_)[D2_DC9((d_175_v36_).cardinality, 94, d_246_v90_, d_130_v2_, d_252_v96_)] if (D2_DC9((d_175_v36_).cardinality, 94, d_246_v90_, d_130_v2_, d_252_v96_)) in (d_254_v98_) else d_130_v2_)})
            nw46_[int(6)] = (d_250_v94_ if d_130_v2_ else d_250_v94_)
            nw46_[int(7)] = d_250_v94_
            nw46_[int(8)] = _dafny.Map({d_249_v93_: d_130_v2_})
            nw46_[int(9)] = d_250_v94_
            nw46_[int(10)] = d_250_v94_
            nw46_[int(11)] = (d_250_v94_) | (d_250_v94_)
            nw46_[int(12)] = _dafny.Map({d_249_v93_: (d_131_v3_)[default__.safeIndex(d_128_v0_, len(d_131_v3_))]})
            nw46_[int(13)] = _dafny.Map({d_249_v93_: d_130_v2_})
            nw46_[int(14)] = d_250_v94_
            nw46_[int(15)] = (_dafny.Map({d_249_v93_: False})).set(d_249_v93_, d_130_v2_)
            nw46_[int(16)] = (d_250_v94_).set((D24_DC67(d_249_v93_)).cf116, d_130_v2_)
            nw46_[int(17)] = d_250_v94_
            nw46_[int(18)] = d_250_v94_
            nw46_[int(19)] = _dafny.Map({d_249_v93_: d_130_v2_})
            nw46_[int(20)] = d_250_v94_
            nw46_[int(21)] = d_250_v94_
            nw46_[int(22)] = d_250_v94_
            nw46_[int(23)] = d_250_v94_
            nw46_[int(24)] = d_250_v94_
            d_255_v99_ = nw46_
            index31_ = default__.safeIndex(147, (d_255_v99_).length(0))
            (d_255_v99_)[index31_] = d_250_v94_
            d_256_v100_: D24
            d_256_v100_ = D24_DC67(d_249_v93_)
            index32_ = default__.safeIndex(147, (d_255_v99_).length(0))
            (d_255_v99_)[index32_] = (d_250_v94_).set((d_256_v100_).cf116, d_130_v2_)
            d_257_v101_: _dafny.Array
            def lambda14_(d_258_v90_):
                def lambda15_(d_259_i14_):
                    return (d_259_i14_) + (len(d_258_v90_))

                return lambda15_

            init7_ = lambda14_(d_246_v90_)
            nw47_ = _dafny.Array(None, 20)
            for i0_7_ in range(nw47_.length(0)):
                nw47_[i0_7_] = init7_(i0_7_)
            d_257_v101_ = nw47_
            d_257_v101_ = d_257_v101_
            d_260_v102_: _dafny.Map
            d_260_v102_ = _dafny.Map({(d_131_v3_)[default__.safeIndex(d_128_v0_, len(d_131_v3_))]: d_249_v93_})
            d_261_v103_: D22
            d_261_v103_ = D22_DC62(not(d_130_v2_), d_128_v0_)
            d_262_v104_: _dafny.Array
            nw48_ = _dafny.Array(None, 26)
            nw48_[int(0)] = d_249_v93_
            nw48_[int(1)] = d_249_v93_
            nw48_[int(2)] = d_249_v93_
            nw48_[int(3)] = d_249_v93_
            nw48_[int(4)] = d_249_v93_
            nw48_[int(5)] = d_249_v93_
            nw48_[int(6)] = d_249_v93_
            nw48_[int(7)] = d_249_v93_
            nw48_[int(8)] = d_249_v93_
            nw48_[int(9)] = d_249_v93_
            nw48_[int(10)] = d_249_v93_
            nw48_[int(11)] = (d_251_v95_)[default__.safeIndex(d_128_v0_, len(d_251_v95_))]
            nw48_[int(12)] = d_249_v93_
            nw48_[int(13)] = d_249_v93_
            nw48_[int(14)] = d_249_v93_
            nw48_[int(15)] = d_249_v93_
            nw48_[int(16)] = d_249_v93_
            nw48_[int(17)] = (d_249_v93_ if d_130_v2_ else d_249_v93_)
            nw48_[int(18)] = (d_256_v100_).cf116
            nw48_[int(19)] = d_249_v93_
            nw48_[int(20)] = d_249_v93_
            nw48_[int(21)] = d_249_v93_
            nw48_[int(22)] = d_249_v93_
            nw48_[int(23)] = d_249_v93_
            nw48_[int(24)] = d_249_v93_
            nw48_[int(25)] = ((d_260_v102_)[(d_261_v103_).cf106] if ((d_261_v103_).cf106) in (d_260_v102_) else d_249_v93_)
            d_262_v104_ = nw48_
            index33_ = default__.safeIndex(898, (d_262_v104_).length(0))
            (d_262_v104_)[index33_] = (d_249_v93_ if d_130_v2_ else d_249_v93_)
            d_263_v105_: _dafny.Array
            def lambda16_(d_264_v0_):
                def lambda17_(d_265_i15_):
                    return D2_DC7(_dafny.Set({d_264_v0_, d_264_v0_, d_264_v0_, d_264_v0_}))

                return lambda17_

            init8_ = lambda16_(d_128_v0_)
            nw49_ = _dafny.Array(None, 7)
            for i0_8_ in range(nw49_.length(0)):
                nw49_[i0_8_] = init8_(i0_8_)
            d_263_v105_ = nw49_
            d_266_v106_: D22
            d_266_v106_ = D22_DC61(d_263_v105_)
            d_267_v107_: D17
            d_267_v107_ = D17_DC48(d_130_v2_, not(d_130_v2_), d_130_v2_)
            pat_let_tv6_ = d_130_v2_
            d_268_v108_: _dafny.Array
            nw50_ = _dafny.Array(None, 9)
            nw50_[int(0)] = d_267_v107_
            nw50_[int(1)] = d_267_v107_
            nw50_[int(2)] = d_267_v107_
            def iife37_(_pat_let11_0):
                def iife38_(d_269_dt__update__tmp_h2_):
                    def iife39_(_pat_let12_0):
                        def iife40_(d_270_dt__update_hcf77_h0_):
                            return D17_DC48(d_270_dt__update_hcf77_h0_, (d_269_dt__update__tmp_h2_).cf78, (d_269_dt__update__tmp_h2_).cf79)
                        return iife40_(_pat_let12_0)
                    return iife39_(pat_let_tv6_)
                return iife38_(_pat_let11_0)
            nw50_[int(3)] = iife37_(d_267_v107_)
            nw50_[int(4)] = d_267_v107_
            nw50_[int(5)] = d_267_v107_
            nw50_[int(6)] = (d_267_v107_ if d_130_v2_ else d_267_v107_)
            nw50_[int(7)] = (d_267_v107_ if not(d_130_v2_) else d_267_v107_)
            nw50_[int(8)] = D17_DC48(d_130_v2_, d_130_v2_, d_130_v2_)
            d_268_v108_ = nw50_
            index34_ = default__.safeIndex(921, (d_268_v108_).length(0))
            (d_268_v108_)[index34_] = d_267_v107_
            d_271_v109_: str
            d_271_v109_ = _dafny.CodePoint('m')
            d_272_v110_: _dafny.Set
            d_272_v110_ = _dafny.Set({d_130_v2_})
            d_273_v111_: _dafny.Map
            d_273_v111_ = _dafny.Map({d_272_v110_: d_128_v0_})
            d_274_v112_: _dafny.Array
            nw51_ = _dafny.Array(False, 23)
            d_274_v112_ = nw51_
            d_275_v113_: D3
            d_275_v113_ = D3_DC10(d_274_v112_)
            pat_let_tv7_ = d_274_v112_
            d_276_v114_: _dafny.Set
            def iife41_(_pat_let13_0):
                def iife42_(d_277_dt__update__tmp_h3_):
                    def iife43_(_pat_let14_0):
                        def iife44_(d_278_dt__update_hcf21_h0_):
                            return D3_DC10(d_278_dt__update_hcf21_h0_)
                        return iife44_(_pat_let14_0)
                    return iife43_(pat_let_tv7_)
                return iife42_(_pat_let13_0)
            d_276_v114_ = _dafny.Set({(iife41_(d_275_v113_)).cf21})
            d_279_v115_: _dafny.Seq
            d_279_v115_ = _dafny.SeqWithoutIsStrInference([((d_273_v111_)[d_272_v110_] if (d_272_v110_) in (d_273_v111_) else -10), len(d_276_v114_)])
            d_280_v116_: _dafny.Map
            d_280_v116_ = _dafny.Map({False: d_128_v0_})
            d_281_v117_: D24
            d_281_v117_ = D24_DC68(79, d_130_v2_, d_280_v116_)
            pat_let_tv8_ = d_130_v2_
            index35_ = default__.safeIndex(898, (d_262_v104_).length(0))
            index36_ = default__.safeIndex(921, (d_268_v108_).length(0))
            def iife45_(_pat_let15_0):
                def iife46_(d_282_dt__update__tmp_h4_):
                    def iife47_(_pat_let16_0):
                        def iife48_(d_283_dt__update_hcf118_h0_):
                            return D24_DC68((d_282_dt__update__tmp_h4_).cf117, d_283_dt__update_hcf118_h0_, (d_282_dt__update__tmp_h4_).cf119)
                        return iife48_(_pat_let16_0)
                    return iife47_(pat_let_tv8_)
                return iife46_(_pat_let15_0)
            rhs16_ = d_249_v93_
            rhs17_ = d_266_v106_
            rhs18_ = default__.fm63((d_279_v115_).set(default__.safeIndex(d_128_v0_, len(d_279_v115_)), d_128_v0_), (iife45_(d_281_v117_)).cf118, d_130_v2_, d_136_globalState_)
            rhs19_ = d_271_v109_
            lhs10_ = d_262_v104_
            lhs11_ = default__.safeIndex(898, (d_262_v104_).length(0))
            lhs12_ = d_268_v108_
            lhs13_ = default__.safeIndex(921, (d_268_v108_).length(0))
            lhs10_[lhs11_] = rhs16_
            d_266_v106_ = rhs17_
            lhs12_[lhs13_] = rhs18_
            d_271_v109_ = rhs19_
            d_128_v0_ = default__.safeDivisionInt(d_128_v0_, (d_279_v115_)[default__.safeIndex(319, len(d_279_v115_))])
        elif True:
            default__.m0(d_128_v0_, d_128_v0_, d_136_globalState_)
            d_284_v118_: C6
            nw52_ = C6()
            nw52_.ctor__(d_130_v2_, d_130_v2_)
            d_284_v118_ = nw52_
            d_131_v3_ = d_131_v3_
            d_285_v119_: C6
            nw53_ = C6()
            nw53_.ctor__(d_130_v2_, d_130_v2_)
            d_285_v119_ = nw53_
            d_286_v120_: _dafny.Seq
            d_286_v120_ = _dafny.SeqWithoutIsStrInference([d_128_v0_])
            d_287_v121_: _dafny.Array
            nw54_ = _dafny.Array(None, 2)
            nw54_[int(0)] = (d_286_v120_)[default__.safeIndex((0) - (d_128_v0_), len(d_286_v120_))]
            nw54_[int(1)] = len(d_286_v120_)
            d_287_v121_ = nw54_
            index37_ = default__.safeIndex(729, (d_287_v121_).length(0))
            (d_287_v121_)[index37_] = d_128_v0_
            index38_ = default__.safeIndex(729, (d_287_v121_).length(0))
            (d_287_v121_)[index38_] = (len(d_246_v90_) if (d_131_v3_)[default__.safeIndex(d_128_v0_, len(d_131_v3_))] else d_128_v0_)
        d_288_v122_: _dafny.Map
        d_288_v122_ = _dafny.Map({d_130_v2_: d_128_v0_})
        d_289_v123_: C0
        nw55_ = C0()
        nw55_.ctor__(d_130_v2_, d_130_v2_)
        d_289_v123_ = nw55_
        d_290_v124_: _dafny.Array
        nw56_ = _dafny.Array(None, 5)
        nw56_[int(0)] = d_289_v123_
        nw56_[int(1)] = d_289_v123_
        nw56_[int(2)] = d_289_v123_
        nw56_[int(3)] = d_289_v123_
        nw56_[int(4)] = d_289_v123_
        d_290_v124_ = nw56_
        d_291_v125_: _dafny.Map
        d_291_v125_ = _dafny.Map({d_290_v124_: (d_289_v123_).f12})
        d_292_v126_: _dafny.Seq
        d_292_v126_ = _dafny.SeqWithoutIsStrInference([-882])
        d_293_v127_: _dafny.Seq
        d_293_v127_ = _dafny.SeqWithoutIsStrInference([len(d_292_v126_)])
        d_294_v128_: _dafny.Map
        d_294_v128_ = _dafny.Map({D0_DC0(len(d_131_v3_)): len(d_293_v127_)})
        d_295_v129_: D2
        d_295_v129_ = D2_DC9(d_128_v0_, len(d_291_v125_), d_246_v90_, d_289_v123_.f13, d_294_v128_)
        d_288_v122_ = (d_288_v122_).set((d_295_v129_).cf19, d_128_v0_)
        d_296_v130_: _dafny.Array
        nw57_ = _dafny.Array(None, 3)
        nw57_[int(0)] = d_289_v123_.f13
        nw57_[int(1)] = d_289_v123_.f13
        nw57_[int(2)] = d_289_v123_.f13
        d_296_v130_ = nw57_
        d_297_v131_: _dafny.Array
        def lambda18_(d_298_i16_):
            return default__.safeModuloInt(d_298_i16_, len(_dafny.Set({_dafny.CodePoint('c')})))

        init9_ = lambda18_
        nw58_ = _dafny.Array(None, 12)
        for i0_9_ in range(nw58_.length(0)):
            nw58_[i0_9_] = init9_(i0_9_)
        d_297_v131_ = nw58_
        d_299_v132_: D21
        d_299_v132_ = D21_DC59(d_296_v130_, d_246_v90_, d_297_v131_, d_130_v2_)
        d_300_v133_: _dafny.Set
        d_300_v133_ = _dafny.Set({(d_299_v132_).cf103})
        d_301_v134_: _dafny.Seq
        d_301_v134_ = _dafny.SeqWithoutIsStrInference([d_300_v133_])
        if (d_301_v134_) < (d_301_v134_):
            index39_ = default__.safeIndex(708, (d_296_v130_).length(0))
            (d_296_v130_)[index39_] = d_289_v123_.f13
            index40_ = default__.safeIndex(708, (d_296_v130_).length(0))
            (d_296_v130_)[index40_] = ((d_289_v123_).f12) == (not((d_289_v123_).f12))
            d_128_v0_ = d_128_v0_
            if ((d_289_v123_.f13 if d_130_v2_ else (d_296_v130_)[default__.safeIndex(708, (d_296_v130_).length(0))])) or (d_289_v123_.f13):
                d_302_v135_: _dafny.Array
                def lambda19_(d_303_v0_):
                    def lambda20_(d_304_i17_):
                        return D2_DC7((D2_DC7(_dafny.Set({d_303_v0_}))).cf13)

                    return lambda20_

                init10_ = lambda19_(d_128_v0_)
                nw59_ = _dafny.Array(None, 29)
                for i0_10_ in range(nw59_.length(0)):
                    nw59_[i0_10_] = init10_(i0_10_)
                d_302_v135_ = nw59_
                d_305_v136_: C4
                nw60_ = C4()
                nw60_.ctor__(d_302_v135_, False)
                d_305_v136_ = nw60_
                d_306_v137_: _dafny.Map
                d_306_v137_ = _dafny.Map({(d_289_v123_).f12: d_305_v136_})
                d_306_v137_ = (d_306_v137_).set((d_296_v130_)[default__.safeIndex(708, (d_296_v130_).length(0))], d_305_v136_)
                d_130_v2_ = (d_131_v3_)[default__.safeIndex(d_128_v0_, len(d_131_v3_))]
                index41_ = default__.safeIndex(373, (d_297_v131_).length(0))
                (d_297_v131_)[index41_] = d_128_v0_
                index42_ = default__.safeIndex(373, (d_297_v131_).length(0))
                (d_297_v131_)[index42_] = default__.safeDivisionInt(d_128_v0_, d_128_v0_)
                d_307_v138_: str
                d_307_v138_ = _dafny.CodePoint('j')
                (d_289_v123_).f13 = (D6_DC17(d_289_v123_.f13, 642, d_307_v138_, d_307_v138_)).cf31
                index43_ = default__.safeIndex(708, (d_296_v130_).length(0))
                index44_ = default__.safeIndex(373, (d_297_v131_).length(0))
                rhs20_ = (d_289_v123_).f12
                rhs21_ = not(not(not((((0) - ((d_297_v131_)[default__.safeIndex(373, (d_297_v131_).length(0))])) + ((d_305_v136_).fm15(d_136_globalState_))) <= (d_128_v0_))))
                rhs22_ = (((0) - ((d_297_v131_)[default__.safeIndex(373, (d_297_v131_).length(0))])) - (485)) != (d_128_v0_)
                rhs23_ = d_128_v0_
                lhs14_ = d_296_v130_
                lhs15_ = default__.safeIndex(708, (d_296_v130_).length(0))
                lhs16_ = d_289_v123_
                lhs17_ = d_297_v131_
                lhs18_ = default__.safeIndex(373, (d_297_v131_).length(0))
                lhs14_[lhs15_] = rhs20_
                d_130_v2_ = rhs21_
                lhs16_.f13 = rhs22_
                lhs17_[lhs18_] = rhs23_
            elif True:
                d_308_v139_: _dafny.Map
                d_308_v139_ = _dafny.Map({d_128_v0_: d_128_v0_})
                (d_289_v123_).f13 = (d_175_v36_) == (_dafny.MultiSet([d_128_v0_, d_128_v0_, 334, ((d_308_v139_)[(0) - (d_128_v0_)] if ((0) - (d_128_v0_)) in (d_308_v139_) else d_128_v0_)]))
                d_309_v140_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_309_v140_ = nw61_
                d_310_v141_: _dafny.Array
                def lambda21_(d_311_v123_):
                    def lambda22_(d_312_i18_):
                        return (d_311_v123_).f12

                    return lambda22_

                init11_ = lambda21_(d_289_v123_)
                nw62_ = _dafny.Array(None, 14)
                for i0_11_ in range(nw62_.length(0)):
                    nw62_[i0_11_] = init11_(i0_11_)
                d_310_v141_ = nw62_
                index45_ = default__.safeIndex(765, (d_309_v140_).length(0))
                (d_309_v140_)[index45_] = d_310_v141_
                d_313_v142_: D15
                d_313_v142_ = D15_DC43(d_128_v0_)
                d_314_v143_: _dafny.Array
                nw63_ = _dafny.Array(D2.default()(), 13)
                d_314_v143_ = nw63_
                d_315_v144_: T0
                nw64_ = C4()
                nw64_.ctor__(d_314_v143_, not((d_296_v130_)[default__.safeIndex(708, (d_296_v130_).length(0))]))
                d_315_v144_ = nw64_
                d_316_v145_: _dafny.Seq
                d_316_v145_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_315_v144_: d_315_v144_.f7})])
                index46_ = default__.safeIndex(765, (d_309_v140_).length(0))
                rhs24_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqlveg"))) + ((d_246_v90_) + (d_246_v90_)))
                rhs25_ = default__.safeModuloInt((d_313_v142_).cf70, (0) - (len((d_316_v145_) + (d_316_v145_))))
                rhs26_ = d_310_v141_
                rhs27_ = (d_289_v123_).f12
                lhs19_ = d_309_v140_
                lhs20_ = default__.safeIndex(765, (d_309_v140_).length(0))
                lhs21_ = d_289_v123_
                d_128_v0_ = rhs24_
                d_128_v0_ = rhs25_
                lhs19_[lhs20_] = rhs26_
                lhs21_.f13 = rhs27_
                d_128_v0_ = (0) - (d_128_v0_)
                d_130_v2_ = (d_130_v2_) or ((d_128_v0_) > (d_128_v0_))
                d_317_v146_: str
                d_317_v146_ = _dafny.CodePoint('v')
                rhs28_ = d_317_v146_
                rhs29_ = d_128_v0_
                rhs30_ = d_297_v131_
                d_317_v146_ = rhs28_
                d_128_v0_ = rhs29_
                d_297_v131_ = rhs30_
            d_318_v147_: _dafny.Seq
            d_318_v147_ = _dafny.SeqWithoutIsStrInference([d_210_v61_, d_210_v61_])
            d_319_v148_: _dafny.Array
            nw65_ = _dafny.Array(None, 5)
            nw65_[int(0)] = d_130_v2_
            nw65_[int(1)] = d_289_v123_.f13
            nw65_[int(2)] = False
            nw65_[int(3)] = (d_289_v123_).f12
            nw65_[int(4)] = d_130_v2_
            d_319_v148_ = nw65_
            d_320_v149_: _dafny.Set
            d_320_v149_ = _dafny.Set({d_319_v148_})
            d_321_v150_: str
            d_321_v150_ = _dafny.CodePoint('e')
            d_322_v151_: C2
            nw66_ = C2()
            nw66_.ctor__(d_320_v149_, d_321_v150_, False)
            d_322_v151_ = nw66_
            d_323_v152_: _dafny.Seq
            d_323_v152_ = _dafny.SeqWithoutIsStrInference([d_322_v151_])
            index47_ = default__.safeIndex(708, (d_296_v130_).length(0))
            (d_296_v130_)[index47_] = default__.fm1(d_318_v147_, default__.safeModuloInt(len(d_323_v152_), d_128_v0_), (861) - (d_128_v0_), d_136_globalState_)
            d_324_v153_: _dafny.Array
            nw67_ = _dafny.Array(D2.default()(), 21)
            d_324_v153_ = nw67_
            d_325_v154_: C4
            nw68_ = C4()
            nw68_.ctor__(d_324_v153_, d_289_v123_.f13)
            d_325_v154_ = nw68_
        elif True:
            d_326_v155_: C6
            nw69_ = C6()
            nw69_.ctor__(d_289_v123_.f13, (d_289_v123_.f13) or ((d_289_v123_).f12))
            d_326_v155_ = nw69_
            nw70_ = C6()
            nw70_.ctor__(d_130_v2_, (d_300_v133_).issubset(_dafny.Set({d_289_v123_.f13})))
            d_326_v155_ = nw70_
            index48_ = default__.safeIndex(941, (d_296_v130_).length(0))
            (d_296_v130_)[index48_] = True
            d_327_v156_: _dafny.Set
            d_327_v156_ = _dafny.Set({d_296_v130_, d_296_v130_})
            index49_ = default__.safeIndex(941, (d_296_v130_).length(0))
            rhs31_ = not((d_128_v0_) != ((d_128_v0_ if d_289_v123_.f13 else d_128_v0_)))
            rhs32_ = (_dafny.Set({d_296_v130_})).ispropersubset(d_327_v156_)
            lhs22_ = d_289_v123_
            lhs23_ = d_296_v130_
            lhs24_ = default__.safeIndex(941, (d_296_v130_).length(0))
            lhs22_.f13 = rhs31_
            lhs23_[lhs24_] = rhs32_
            d_131_v3_ = (d_131_v3_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivahwxl"))), len(d_131_v3_)), (d_296_v130_)[default__.safeIndex(941, (d_296_v130_).length(0))])
            if ((d_246_v90_) + (d_246_v90_)) <= ((d_246_v90_) + (d_246_v90_)):
                rhs33_ = (0) - (((0) - (d_128_v0_)) + (d_128_v0_))
                rhs34_ = d_289_v123_.f13
                rhs35_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_328_i19_ in range(default__.abs(772))])) + ((d_246_v90_) + (d_246_v90_))
                d_128_v0_ = rhs33_
                d_130_v2_ = rhs34_
                d_246_v90_ = rhs35_
                d_329_v157_: D12
                d_329_v157_ = D12_DC30(d_128_v0_, (d_326_v155_).f9, d_292_v126_, d_128_v0_)
                d_330_v158_: _dafny.Seq
                d_330_v158_ = _dafny.SeqWithoutIsStrInference([d_329_v157_, d_329_v157_, d_329_v157_, d_329_v157_, d_329_v157_])
                d_331_v159_: _dafny.Map
                d_331_v159_ = _dafny.Map({d_330_v158_: d_289_v123_.f13})
                d_332_v160_: _dafny.Seq
                d_332_v160_ = _dafny.SeqWithoutIsStrInference([d_288_v122_, d_288_v122_])
                d_331_v159_ = (d_331_v159_).set(_dafny.SeqWithoutIsStrInference([d_329_v157_]), (d_332_v160_) < (d_332_v160_))
                d_333_v161_: _dafny.Array
                def lambda23_(d_334_v3_, d_335_v0_, d_336_v130_, d_337_v123_):
                    def lambda24_(d_338_i20_):
                        return ((d_336_v130_)[default__.safeIndex(941, (d_336_v130_).length(0))] if (d_334_v3_)[default__.safeIndex(d_335_v0_, len(d_334_v3_))] else (d_337_v123_).f12)

                    return lambda24_

                init12_ = lambda23_(d_131_v3_, d_128_v0_, d_296_v130_, d_289_v123_)
                nw71_ = _dafny.Array(None, 10)
                for i0_12_ in range(nw71_.length(0)):
                    nw71_[i0_12_] = init12_(i0_12_)
                d_333_v161_ = nw71_
                d_339_v162_: _dafny.Map
                d_339_v162_ = _dafny.Map({d_128_v0_: d_128_v0_})
                rhs36_ = len((d_339_v162_) | (d_339_v162_))
                rhs37_ = d_333_v161_
                d_128_v0_ = rhs36_
                d_333_v161_ = rhs37_
                (d_289_v123_).f13 = True
                d_340_v163_: D13
                d_340_v163_ = D13_DC35()
                d_341_v164_: _dafny.Map
                d_341_v164_ = _dafny.Map({539: (d_340_v163_ if d_130_v2_ else d_340_v163_)})
                d_342_v165_: _dafny.Map
                d_342_v165_ = _dafny.Map({d_289_v123_.f13: d_292_v126_})
                d_343_v166_: T1
                nw72_ = C7()
                nw72_.ctor__(d_342_v165_, (d_289_v123_).f12)
                d_343_v166_ = nw72_
                d_344_v167_: _dafny.Array
                nw73_ = _dafny.Array(None, 2)
                nw73_[int(0)] = d_343_v166_
                nw73_[int(1)] = d_343_v166_
                d_344_v167_ = nw73_
                d_345_v168_: _dafny.Seq
                d_345_v168_ = _dafny.SeqWithoutIsStrInference([d_344_v167_, d_344_v167_, d_344_v167_])
                d_346_v169_: C1
                nw74_ = C1()
                nw74_.ctor__(d_343_v166_.f7)
                d_346_v169_ = nw74_
                d_347_v170_: _dafny.MultiSet
                d_347_v170_ = _dafny.MultiSet([d_346_v169_])
                d_348_v171_: _dafny.Map
                d_348_v171_ = _dafny.Map({d_289_v123_.f13: d_346_v169_})
                d_349_v172_: _dafny.Array
                nw75_ = _dafny.Array(None, 28)
                nw75_[int(0)] = d_344_v167_
                nw75_[int(1)] = d_344_v167_
                nw75_[int(2)] = d_344_v167_
                nw75_[int(3)] = d_344_v167_
                nw75_[int(4)] = d_344_v167_
                nw75_[int(5)] = d_344_v167_
                nw75_[int(6)] = d_344_v167_
                nw75_[int(7)] = (d_345_v168_)[default__.safeIndex(((d_347_v170_)[((d_348_v171_)[(d_326_v155_).f9] if ((d_326_v155_).f9) in (d_348_v171_) else d_346_v169_)] if (((d_348_v171_)[(d_326_v155_).f9] if ((d_326_v155_).f9) in (d_348_v171_) else d_346_v169_)) in (d_347_v170_) else -16), len(d_345_v168_))]
                nw75_[int(8)] = d_344_v167_
                nw75_[int(9)] = d_344_v167_
                nw75_[int(10)] = (d_345_v168_)[default__.safeIndex(((d_339_v162_)[d_128_v0_] if (d_128_v0_) in (d_339_v162_) else d_128_v0_), len(d_345_v168_))]
                nw75_[int(11)] = (d_344_v167_ if (d_289_v123_).f12 else d_344_v167_)
                nw75_[int(12)] = d_344_v167_
                nw75_[int(13)] = d_344_v167_
                nw75_[int(14)] = d_344_v167_
                nw75_[int(15)] = (d_344_v167_ if (d_326_v155_).f9 else d_344_v167_)
                nw75_[int(16)] = d_344_v167_
                nw75_[int(17)] = d_344_v167_
                nw75_[int(18)] = d_344_v167_
                nw75_[int(19)] = d_344_v167_
                nw75_[int(20)] = d_344_v167_
                nw75_[int(21)] = d_344_v167_
                nw75_[int(22)] = d_344_v167_
                nw75_[int(23)] = (d_344_v167_ if (d_326_v155_).f9 else d_344_v167_)
                nw75_[int(24)] = d_344_v167_
                nw75_[int(25)] = d_344_v167_
                nw75_[int(26)] = d_344_v167_
                nw75_[int(27)] = d_344_v167_
                d_349_v172_ = nw75_
                index50_ = default__.safeIndex(785, (d_349_v172_).length(0))
                (d_349_v172_)[index50_] = d_344_v167_
                d_350_v174_: _dafny.Set
                d_350_v174_ = _dafny.Set({48, d_128_v0_})
                d_351_v177_: str
                d_351_v177_ = _dafny.CodePoint('x')
                d_352_v178_: _dafny.Map
                def iife49_():
                    def iife51_():
                        coll17_ = _dafny.Map()
                        compr_17_: str
                        for compr_17_ in ((d_246_v90_).set(default__.safeIndex(len(d_131_v3_), len(d_246_v90_)), d_351_v177_)).Elements:
                            d_353_v176_: str = compr_17_
                            if (d_353_v176_) in ((d_246_v90_).set(default__.safeIndex(len(d_131_v3_), len(d_246_v90_)), d_351_v177_)):
                                coll17_[d_353_v176_] = (d_289_v123_).f12
                        return _dafny.Map(coll17_)
                    coll15_ = _dafny.Map()
                    def iife50_():
                        coll16_ = _dafny.Map()
                        compr_16_: str
                        for compr_16_ in ((d_246_v90_).set(default__.safeIndex(len(d_131_v3_), len(d_246_v90_)), d_351_v177_)).Elements:
                            d_353_v176_: str = compr_16_
                            if (d_353_v176_) in ((d_246_v90_).set(default__.safeIndex(len(d_131_v3_), len(d_246_v90_)), d_351_v177_)):
                                coll16_[d_353_v176_] = (d_289_v123_).f12
                        return _dafny.Map(coll16_)
                    compr_15_: str
                    for compr_15_ in (iife50_()
                    ).keys.Elements:
                        d_354_v175_: str = compr_15_
                        if (d_354_v175_) in (iife51_()
                        ):
                            coll15_[d_354_v175_] = d_128_v0_
                    return _dafny.Map(coll15_)
                d_352_v178_ = _dafny.Map({d_343_v166_: len((D25_DC70(iife49_()
)).cf125)})
                index51_ = default__.safeIndex(785, (d_349_v172_).length(0))
                def iife52_():
                    coll18_ = _dafny.Map()
                    compr_18_: int
                    for compr_18_ in (d_350_v174_).Elements:
                        d_355_v173_: int = compr_18_
                        if (d_355_v173_) in (d_350_v174_):
                            coll18_[(d_355_v173_) - (d_128_v0_)] = D13_DC35()
                    return _dafny.Map(coll18_)
                rhs38_ = (d_128_v0_) - (d_128_v0_)
                rhs39_ = ((d_341_v164_) | (iife52_()
                )).set(len(d_292_v126_), d_340_v163_)
                rhs40_ = d_344_v167_
                rhs41_ = (d_343_v166_) not in ((d_352_v178_) | (d_352_v178_))
                lhs25_ = d_349_v172_
                lhs26_ = default__.safeIndex(785, (d_349_v172_).length(0))
                lhs27_ = d_289_v123_
                d_128_v0_ = rhs38_
                d_341_v164_ = rhs39_
                lhs25_[lhs26_] = rhs40_
                lhs27_.f13 = rhs41_
            elif True:
                d_300_v133_ = ((d_300_v133_) - (d_300_v133_)) - (d_300_v133_)
                d_356_v179_: _dafny.Array
                def lambda25_(d_357_v123_, d_358_v155_):
                    def lambda26_(d_359_i21_):
                        return _dafny.Set({d_357_v123_.f13, (d_358_v155_).f9})

                    return lambda26_

                init13_ = lambda25_(d_289_v123_, d_326_v155_)
                nw76_ = _dafny.Array(None, 15)
                for i0_13_ in range(nw76_.length(0)):
                    nw76_[i0_13_] = init13_(i0_13_)
                d_356_v179_ = nw76_
                d_360_v180_: _dafny.Array
                nw77_ = _dafny.Array(None, 8)
                d_360_v180_ = nw77_
                d_361_v181_: C5
                nw78_ = C5()
                nw78_.ctor__(d_130_v2_)
                d_361_v181_ = nw78_
                d_362_v182_: D26
                d_362_v182_ = D26_DC73(d_361_v181_)
                d_363_v183_: str
                d_363_v183_ = _dafny.CodePoint('r')
                d_364_v184_: C2
                nw79_ = C2()
                nw79_.ctor__(d_327_v156_, d_363_v183_, (d_296_v130_)[default__.safeIndex(941, (d_296_v130_).length(0))])
                d_364_v184_ = nw79_
                d_365_v185_: _dafny.Map
                d_365_v185_ = _dafny.Map({(d_362_v182_).cf129: d_364_v184_})
                index52_ = default__.safeIndex(715, (d_360_v180_).length(0))
                (d_360_v180_)[index52_] = ((d_365_v185_)[d_361_v181_] if (d_361_v181_) in (d_365_v185_) else d_364_v184_)
                index53_ = default__.safeIndex(941, (d_296_v130_).length(0))
                index54_ = default__.safeIndex(715, (d_360_v180_).length(0))
                rhs42_ = d_356_v179_
                rhs43_ = False
                rhs44_ = d_128_v0_
                rhs45_ = d_364_v184_
                lhs28_ = d_296_v130_
                lhs29_ = default__.safeIndex(941, (d_296_v130_).length(0))
                lhs30_ = d_360_v180_
                lhs31_ = default__.safeIndex(715, (d_360_v180_).length(0))
                d_356_v179_ = rhs42_
                lhs28_[lhs29_] = rhs43_
                d_128_v0_ = rhs44_
                lhs30_[lhs31_] = rhs45_
                index55_ = default__.safeIndex(250, (d_297_v131_).length(0))
                (d_297_v131_)[index55_] = len(d_246_v90_)
                d_366_v186_: _dafny.Map
                d_366_v186_ = _dafny.Map({d_128_v0_: d_128_v0_})
                index56_ = default__.safeIndex(250, (d_297_v131_).length(0))
                (d_297_v131_)[index56_] = len((d_366_v186_) | (d_366_v186_))
                d_367_v187_: D15
                d_367_v187_ = D15_DC43(len((d_246_v90_).set(default__.safeIndex((d_326_v155_).fm12((d_289_v123_).f12, d_128_v0_, len(d_292_v126_), (d_289_v123_).f12, d_136_globalState_), len(d_246_v90_)), (d_364_v184_).f16)))
                index57_ = default__.safeIndex(250, (d_297_v131_).length(0))
                index58_ = default__.safeIndex(250, (d_297_v131_).length(0))
                index59_ = default__.safeIndex(250, (d_297_v131_).length(0))
                index60_ = default__.safeIndex(250, (d_297_v131_).length(0))
                rhs46_ = d_128_v0_
                rhs47_ = ((d_297_v131_)[default__.safeIndex(250, (d_297_v131_).length(0))]) - ((d_297_v131_)[default__.safeIndex(250, (d_297_v131_).length(0))])
                rhs48_ = (d_367_v187_).cf70
                rhs49_ = not((d_326_v155_).f9)
                rhs50_ = default__.fm0(d_136_globalState_)
                lhs32_ = d_297_v131_
                lhs33_ = default__.safeIndex(250, (d_297_v131_).length(0))
                lhs34_ = d_297_v131_
                lhs35_ = default__.safeIndex(250, (d_297_v131_).length(0))
                lhs36_ = d_297_v131_
                lhs37_ = default__.safeIndex(250, (d_297_v131_).length(0))
                lhs38_ = d_289_v123_
                lhs39_ = d_297_v131_
                lhs40_ = default__.safeIndex(250, (d_297_v131_).length(0))
                lhs32_[lhs33_] = rhs46_
                lhs34_[lhs35_] = rhs47_
                lhs36_[lhs37_] = rhs48_
                lhs38_.f13 = rhs49_
                lhs39_[lhs40_] = rhs50_
                d_368_v188_: D13
                d_368_v188_ = D13_DC35()
                d_368_v188_ = d_368_v188_
            d_297_v131_ = d_297_v131_
        d_369_v189_: C6
        nw80_ = C6()
        nw80_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpio"))) <= (d_246_v90_), (d_289_v123_).f12)
        d_369_v189_ = nw80_
        d_128_v0_ = default__.safeDivisionInt((d_128_v0_) - (d_128_v0_), len(d_246_v90_))
        d_128_v0_ = d_128_v0_
        d_130_v2_ = (d_131_v3_)[default__.safeIndex((d_292_v126_)[default__.safeIndex(d_128_v0_, len(d_292_v126_))], len(d_131_v3_))]
        d_370_v190_: _dafny.Map
        d_370_v190_ = _dafny.Map({d_128_v0_: d_246_v90_})
        (d_289_v123_).f13 = ((((d_370_v190_)[d_128_v0_] if (d_128_v0_) in (d_370_v190_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjbvvw")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwgsuo")))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_371_i22_ in range(default__.abs(636))]))
        index61_ = default__.safeIndex(514, (d_297_v131_).length(0))
        (d_297_v131_)[index61_] = d_128_v0_
        index62_ = default__.safeIndex(514, (d_297_v131_).length(0))
        (d_297_v131_)[index62_] = d_128_v0_
        d_372_v191_: _dafny.Array
        nw81_ = _dafny.Array(None, 13)
        nw81_[int(0)] = d_296_v130_
        nw81_[int(1)] = d_296_v130_
        nw81_[int(2)] = d_296_v130_
        nw81_[int(3)] = d_296_v130_
        nw81_[int(4)] = d_296_v130_
        nw81_[int(5)] = d_296_v130_
        nw81_[int(6)] = d_296_v130_
        nw81_[int(7)] = d_296_v130_
        nw81_[int(8)] = (d_296_v130_ if d_130_v2_ else d_296_v130_)
        nw81_[int(9)] = d_296_v130_
        nw81_[int(10)] = d_296_v130_
        nw81_[int(11)] = d_296_v130_
        nw81_[int(12)] = d_296_v130_
        d_372_v191_ = nw81_
        d_372_v191_ = d_372_v191_
        rhs51_ = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('k') if (d_289_v123_).f12 else _dafny.CodePoint('x')) for d_373_i23_ in range(default__.abs(781))])
        rhs52_ = _dafny.MultiSet((d_293_v127_) + (d_293_v127_))
        rhs53_ = d_297_v131_
        d_246_v90_ = rhs51_
        d_175_v36_ = rhs52_
        d_297_v131_ = rhs53_
        _dafny.print(_dafny.string_of(d_128_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v1_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({-231})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_130_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v3_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v4_) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[0]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[1]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[2]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[3]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[4]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[5]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[6]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[7]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[8]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[9]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[10]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[11]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[12]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[13]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[14]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[15]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[16]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[17]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[18]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[19]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[20]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[21]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[22]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[23]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[24]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[25]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[26]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[27]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v5_)[28]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_globalState_).f2) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({-231})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_globalState_).f3) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[0]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[1]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[2]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[3]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[4]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[5]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[6]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[7]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[8]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[9]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[10]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[11]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[12]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[13]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[14]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[15]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[16]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[17]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[18]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[19]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[20]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[21]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[22]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[23]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[24]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[25]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[26]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[27]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_136_globalState_).f5)[28]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v36_) == (_dafny.MultiSet([1, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_v61_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v62_) == (_dafny.Map({D0_DC0(462): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v63_) == (_dafny.MultiSet([_dafny.MultiSet([462, 462, 826])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_v90_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_288_v122_) == (_dafny.Map({True: 462}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_289_v123_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_289_v123_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_290_v124_)[0]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v124_)[0].f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_290_v124_)[1]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v124_)[1].f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_290_v124_)[2]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v124_)[2].f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_290_v124_)[3]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v124_)[3].f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_290_v124_)[4]).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v124_)[4].f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_291_v125_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v126_) == (_dafny.SeqWithoutIsStrInference([-882]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v127_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v128_) == (_dafny.Map({D0_DC0(1): 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v129_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v129_).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_295_v129_).cf18).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v129_).cf19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_295_v129_).cf20) == (_dafny.Map({D0_DC0(1): 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v130_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v130_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v130_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v131_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf100)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf100)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf100)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_299_v132_).cf101).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v132_).cf102)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_299_v132_).cf103))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_v133_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_301_v134_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_369_v189_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_370_v190_) == (_dafny.Map({0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffmgmg"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[1])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[1])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[1])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[2])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[2])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[2])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[3])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[3])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[3])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[4])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[4])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[4])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[5])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[5])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[5])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[6])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[6])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[6])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[7])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[7])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[7])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[8])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[8])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[8])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[9])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[9])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[9])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[10])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[10])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[10])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[11])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[11])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[11])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[12])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[12])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v191_)[12])[2]))
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
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
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
        return lambda: D1_DC4(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf14)}, {self.cf15.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf16', Any), ('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {self.cf18.VerbatimString(True)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC11(int(0), int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)

class D3_DC11(D3, NamedTuple('DC11', [('cf22', Any), ('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

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
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC14(D5, NamedTuple('DC14', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)

class D6_DC16(D6, NamedTuple('DC16', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC19(D7, NamedTuple('DC19', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(int(0), _dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC21(D8, NamedTuple('DC21', [('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf41', Any), ('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)

class D9_DC24(D9, NamedTuple('DC24', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)

class D10_DC27(D10, NamedTuple('DC27', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)

class D11_DC28(D11, NamedTuple('DC28', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC30(int(0), False, _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)

class D12_DC30(D12, NamedTuple('DC30', [('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC29(D12, NamedTuple('DC29', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf48 == __o.cf48
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
        return lambda: D13_DC35()
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
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)

class D13_DC35(D13, NamedTuple('DC35', [])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35)
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC36(D13, NamedTuple('DC36', [('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC37(D13, NamedTuple('DC37', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC39(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)

class D14_DC39(D14, NamedTuple('DC39', [('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({self.cf63.VerbatimString(True)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC40(D14, NamedTuple('DC40', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC38(D14, NamedTuple('DC38', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC41(D14, NamedTuple('DC41', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC43(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)

class D15_DC43(D15, NamedTuple('DC43', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC44(D15, NamedTuple('DC44', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC46(int(0), _dafny.Array(None, 0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)

class D16_DC46(D16, NamedTuple('DC46', [('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC45(D16, NamedTuple('DC45', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC48(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D17_DC49)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)

class D17_DC48(D17, NamedTuple('DC48', [('cf77', Any), ('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC49(D17, NamedTuple('DC49', [('cf80', Any), ('cf81', Any), ('cf82', Any), ('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC49({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC49) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82 and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC47(D17, NamedTuple('DC47', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC51(_dafny.Seq({}), int(0), False)
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

class D18_DC51(D18, NamedTuple('DC51', [('cf85', Any), ('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51({_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51) and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC52(D18, NamedTuple('DC52', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC52'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC52)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC50(D18, NamedTuple('DC50', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D19_DC53)

class D19_DC53(D19, NamedTuple('DC53', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC53({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC53) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC55(False, int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D20_DC55)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D20_DC56)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)

class D20_DC55(D20, NamedTuple('DC55', [('cf90', Any), ('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC55({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC55) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC56(D20, NamedTuple('DC56', [('cf93', Any), ('cf94', Any), ('cf95', Any), ('cf96', Any), ('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC56({_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC56) and self.cf93 == __o.cf93 and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC54(D20, NamedTuple('DC54', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC58(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D21_DC58)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D21_DC59)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D21_DC60)

class D21_DC58(D21, NamedTuple('DC58', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC58({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC58) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC59(D21, NamedTuple('DC59', [('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC59({_dafny.string_of(self.cf100)}, {self.cf101.VerbatimString(True)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC59) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC57(D21, NamedTuple('DC57', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC60(D21, NamedTuple('DC60', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC60({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC60) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC62(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D22_DC62)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D22_DC61)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D22_DC63)

class D22_DC62(D22, NamedTuple('DC62', [('cf106', Any), ('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC62({_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC62) and self.cf106 == __o.cf106 and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC61(D22, NamedTuple('DC61', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC61({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC61) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC63(D22, NamedTuple('DC63', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC63({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC63) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC65(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D23_DC65)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D23_DC66)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D23_DC64)

class D23_DC65(D23, NamedTuple('DC65', [('cf110', Any), ('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC65({_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC65) and self.cf110 == __o.cf110 and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC66(D23, NamedTuple('DC66', [('cf112', Any), ('cf113', Any), ('cf114', Any), ('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC66({_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)}, {_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC66) and self.cf112 == __o.cf112 and self.cf113 == __o.cf113 and self.cf114 == __o.cf114 and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC64(D23, NamedTuple('DC64', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC64({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC64) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC68(int(0), False, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D24_DC68)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D24_DC69)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D24_DC67)

class D24_DC68(D24, NamedTuple('DC68', [('cf117', Any), ('cf118', Any), ('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC68({_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC68) and self.cf117 == __o.cf117 and self.cf118 == __o.cf118 and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC69(D24, NamedTuple('DC69', [('cf120', Any), ('cf121', Any), ('cf122', Any), ('cf123', Any), ('cf124', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC69({_dafny.string_of(self.cf120)}, {_dafny.string_of(self.cf121)}, {_dafny.string_of(self.cf122)}, {_dafny.string_of(self.cf123)}, {_dafny.string_of(self.cf124)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC69) and self.cf120 == __o.cf120 and self.cf121 == __o.cf121 and self.cf122 == __o.cf122 and self.cf123 == __o.cf123 and self.cf124 == __o.cf124
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC67(D24, NamedTuple('DC67', [('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC67({_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC67) and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC71(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D25_DC71)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D25_DC70)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D25_DC72)

class D25_DC71(D25, NamedTuple('DC71', [('cf126', Any), ('cf127', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC71({_dafny.string_of(self.cf126)}, {_dafny.string_of(self.cf127)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC71) and self.cf126 == __o.cf126 and self.cf127 == __o.cf127
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC70(D25, NamedTuple('DC70', [('cf125', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC70({_dafny.string_of(self.cf125)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC70) and self.cf125 == __o.cf125
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC72(D25, NamedTuple('DC72', [('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC72({_dafny.string_of(self.cf128)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC72) and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC74()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D26_DC74)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D26_DC73)

class D26_DC74(D26, NamedTuple('DC74', [])):
    def __dafnystr__(self) -> str:
        return f'D26.DC74'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC74)
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC73(D26, NamedTuple('DC73', [('cf129', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC73({_dafny.string_of(self.cf129)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC73) and self.cf129 == __o.cf129
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC76(_dafny.Set({}), _dafny.CodePoint('D'), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D27_DC76)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D27_DC75)

class D27_DC76(D27, NamedTuple('DC76', [('cf131', Any), ('cf132', Any), ('cf133', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC76({_dafny.string_of(self.cf131)}, {_dafny.string_of(self.cf132)}, {_dafny.string_of(self.cf133)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC76) and self.cf131 == __o.cf131 and self.cf132 == __o.cf132 and self.cf133 == __o.cf133
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC75(D27, NamedTuple('DC75', [('cf130', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC75({_dafny.string_of(self.cf130)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC75) and self.cf130 == __o.cf130
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC78(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D28_DC78)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D28_DC77)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D28_DC79)

class D28_DC78(D28, NamedTuple('DC78', [('cf135', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC78({_dafny.string_of(self.cf135)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC78) and self.cf135 == __o.cf135
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC77(D28, NamedTuple('DC77', [('cf134', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC77({_dafny.string_of(self.cf134)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC77) and self.cf134 == __o.cf134
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC79(D28, NamedTuple('DC79', [('cf136', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC79({_dafny.string_of(self.cf136)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC79) and self.cf136 == __o.cf136
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC81(_dafny.CodePoint('D'), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D29_DC81)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D29_DC80)

class D29_DC81(D29, NamedTuple('DC81', [('cf138', Any), ('cf139', Any), ('cf140', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC81({_dafny.string_of(self.cf138)}, {_dafny.string_of(self.cf139)}, {_dafny.string_of(self.cf140)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC81) and self.cf138 == __o.cf138 and self.cf139 == __o.cf139 and self.cf140 == __o.cf140
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC80(D29, NamedTuple('DC80', [('cf137', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC80({_dafny.string_of(self.cf137)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC80) and self.cf137 == __o.cf137
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: D30_DC83(_dafny.Map({}), int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D30_DC83)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D30_DC84)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D30_DC82)

class D30_DC83(D30, NamedTuple('DC83', [('cf142', Any), ('cf143', Any), ('cf144', Any), ('cf145', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC83({_dafny.string_of(self.cf142)}, {_dafny.string_of(self.cf143)}, {_dafny.string_of(self.cf144)}, {_dafny.string_of(self.cf145)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC83) and self.cf142 == __o.cf142 and self.cf143 == __o.cf143 and self.cf144 == __o.cf144 and self.cf145 == __o.cf145
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC84(D30, NamedTuple('DC84', [('cf146', Any), ('cf147', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC84({_dafny.string_of(self.cf146)}, {_dafny.string_of(self.cf147)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC84) and self.cf146 == __o.cf146 and self.cf147 == __o.cf147
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC82(D30, NamedTuple('DC82', [('cf141', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC82({_dafny.string_of(self.cf141)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC82) and self.cf141 == __o.cf141
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def m2(self, p0, p1, p2, globalState):
        pass

    def m3(self, globalState):
        pass


class T1:
    pass
    def m4(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self._f0: bool = False
        self._f1: bool = False
        self._f2: _dafny.Seq = _dafny.Seq({})
        self._f3: _dafny.MultiSet = _dafny.MultiSet({})
        self._f4: int = int(0)
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5

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

class C0:
    def  __init__(self):
        self.f13: bool = False
        self._f12: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f12, f13):
        (self)._f12 = f12
        (self).f13 = f13

    def fm16(self, p0, globalState):
        return self.f13

    def fm17(self, p0, p1, globalState):
        source0_ = D0_DC1(-379)
        if source0_.is_DC1:
            d_374___mcc_h0_ = source0_.cf1
            d_375_cf1_ = d_374___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([(self).f12])
        elif source0_.is_DC2:
            d_376___mcc_h1_ = source0_.cf2
            d_377___mcc_h2_ = source0_.cf3
            d_378___mcc_h3_ = source0_.cf4
            d_379___mcc_h4_ = source0_.cf5
            d_380_cf5_ = d_379___mcc_h4_
            d_381_cf4_ = d_378___mcc_h3_
            d_382_cf3_ = d_377___mcc_h2_
            d_383_cf2_ = d_376___mcc_h1_
            return (_dafny.SeqWithoutIsStrInference([d_383_cf2_, False])) + (_dafny.SeqWithoutIsStrInference([d_383_cf2_]))
        elif True:
            d_384___mcc_h5_ = source0_.cf0
            d_385_cf0_ = d_384___mcc_h5_
            return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([self.f13]))

    @property
    def f12(self):
        return self._f12

class C1(T0):
    def  __init__(self):
        self._f7: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def ctor__(self, f7):
        (self).f7 = f7

    def fm31(self, p0, p1, p2, p3, globalState):
        return _dafny.Set({self.f7, True, False})

    def fm32(self, p0, p1, p2, p3, globalState):
        return (D2_DC9(len(_dafny.Map({self.f7: self.f7})), (D1_DC5(self.f7, _dafny.SeqWithoutIsStrInference([-556 for d_386_i0_ in range(default__.abs(-515))]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywsvi"))))).cf11, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nemcnlivt")), self.f7, _dafny.Map({D0_DC0((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({self.f7: D1_DC5(self.f7, _dafny.SeqWithoutIsStrInference([781, 663]), -754)}))])))): (D3_DC11(87, (_dafny.MultiSet([self.f7])).cardinality, 546, self.f7)).cf23}))).cf19

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Set = _dafny.Set({})
        d_387_v0_: D1
        d_387_v0_ = D1_DC4(self.f7, self.f7)
        d_388_v1_: _dafny.Map
        d_388_v1_ = _dafny.Map({d_387_v0_: p1})
        d_389_v2_: _dafny.Set
        d_389_v2_ = _dafny.Set({-86, -766})
        def iife53_(_pat_let17_0):
            def iife54_(d_390_dt__update__tmp_h0_):
                def iife55_(_pat_let18_0):
                    def iife56_(d_391_dt__update_hcf8_h0_):
                        return D1_DC4((d_390_dt__update__tmp_h0_).cf7, d_391_dt__update_hcf8_h0_)
                    return iife56_(_pat_let18_0)
                return iife55_(False)
            return iife54_(_pat_let17_0)
        d_388_v1_ = (d_388_v1_).set(iife53_(D1_DC4(not(p2), p2)), len((_dafny.Set({p1, p1, p1})) | (d_389_v2_)))
        d_392_v3_: C0
        nw82_ = C0()
        nw82_.ctor__(not((p2) == (p2)), (self.f7 if self.f7 else p0))
        d_392_v3_ = nw82_
        d_393_v4_: _dafny.Array
        nw83_ = _dafny.Array(False, 24)
        d_393_v4_ = nw83_
        index63_ = default__.safeIndex(500, (d_393_v4_).length(0))
        (d_393_v4_)[index63_] = not (False) or (not(d_392_v3_.f13))
        d_394_v5_: _dafny.Map
        d_394_v5_ = _dafny.Map({p1: self.f7})
        index64_ = default__.safeIndex(500, (d_393_v4_).length(0))
        (d_393_v4_)[index64_] = ((d_394_v5_)[p1] if (p1) in (d_394_v5_) else (p0 if (d_392_v3_).f12 else (d_392_v3_).f12))
        d_395_i0_: int
        d_395_i0_ = 0
        with _dafny.label("1"):
            while not(True):
                with _dafny.c_label("1"):
                    if (d_395_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_395_i0_ = (d_395_i0_) + (1)
                    d_396_v6_: _dafny.Map
                    d_396_v6_ = _dafny.Map({p1: p1})
                    d_397_v7_: _dafny.Seq
                    d_397_v7_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                    d_398_v8_: _dafny.Seq
                    d_398_v8_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_396_v6_)]), d_397_v7_])
                    d_399_v9_: _dafny.Seq
                    d_399_v9_ = _dafny.SeqWithoutIsStrInference([d_398_v8_, _dafny.SeqWithoutIsStrInference([(d_397_v7_).set(default__.safeIndex(-573, len(d_397_v7_)), p1) for d_400_i1_ in range(default__.abs(-313))]), d_398_v8_])
                    d_401_v10_: _dafny.Seq
                    d_401_v10_ = _dafny.SeqWithoutIsStrInference([(d_393_v4_)[default__.safeIndex(500, (d_393_v4_).length(0))]])
                    d_402_v11_: _dafny.Map
                    d_402_v11_ = _dafny.Map({d_401_v10_: d_401_v10_})
                    d_403_v12_: _dafny.Map
                    d_403_v12_ = _dafny.Map({default__.fm0(globalState): ((d_399_v9_)[default__.safeIndex(p1, len(d_399_v9_))]).set(default__.safeIndex(len(d_402_v11_), len((d_399_v9_)[default__.safeIndex(p1, len(d_399_v9_))])), d_397_v7_)})
                    d_403_v12_ = (d_403_v12_).set(p1, d_398_v8_)
                    d_404_v13_: int
                    d_404_v13_ = -438
                    d_405_v14_: _dafny.MultiSet
                    d_405_v14_ = _dafny.MultiSet([p1])
                    d_406_v15_: str
                    d_406_v15_ = _dafny.CodePoint('n')
                    d_407_v16_: _dafny.Seq
                    d_407_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrennfvri"))
                    d_408_v17_: _dafny.Map
                    d_408_v17_ = _dafny.Map({d_407_v16_: p1})
                    d_404_v13_ = (p1 if (d_405_v14_).ispropersubset(default__.fm33(p1, d_406_v15_, globalState)) else default__.safeDivisionInt(len(d_408_v17_), d_404_v13_))
                    d_404_v13_ = (0) - (p1)
                    d_404_v13_ = default__.safeModuloInt(697, (p1) - (d_404_v13_))
                    pass
            pass
        d_409_v18_: C0
        nw84_ = C0()
        nw84_.ctor__(p2, self.f7)
        d_409_v18_ = nw84_
        d_410_v19_: _dafny.Seq
        d_410_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppm"))
        d_410_v19_ = d_410_v19_
        d_411_v20_: _dafny.Map
        d_411_v20_ = _dafny.Map({d_409_v18_.f13: p2})
        r0 = (d_411_v20_) | (d_411_v20_)
        d_412_v21_: _dafny.Seq
        d_412_v21_ = _dafny.SeqWithoutIsStrInference([self.f7])
        d_413_v22_: _dafny.Set
        d_413_v22_ = _dafny.Set({p2, d_409_v18_.f13, (d_392_v3_).fm16(False, globalState), (d_412_v21_)[default__.safeIndex(p1, len(d_412_v21_))], (d_392_v3_).f12})
        d_414_v23_: _dafny.Map
        d_414_v23_ = _dafny.Map({(d_393_v4_)[default__.safeIndex(500, (d_393_v4_).length(0))]: (d_413_v22_).intersection(d_413_v22_)})
        r1 = ((d_414_v23_)[p2] if (p2) in (d_414_v23_) else d_413_v22_)
        return r0, r1

    def m3(self, globalState):
        source1_ = D9_DC24()
        if source1_.is_DC24:
            d_415_v0_: int
            d_415_v0_ = 456
            d_416_v1_: _dafny.Set
            d_416_v1_ = _dafny.Set({d_415_v0_, d_415_v0_})
            d_417_v2_: _dafny.Seq
            d_417_v2_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_418_v3_: _dafny.Map
            d_418_v3_ = _dafny.Map({self.f7: len(d_417_v2_)})
            d_419_v4_: _dafny.Array
            nw85_ = _dafny.Array(False, 2)
            d_419_v4_ = nw85_
            d_420_v5_: _dafny.Set
            d_420_v5_ = _dafny.Set({d_419_v4_, d_419_v4_})
            d_421_v6_: _dafny.MultiSet
            d_421_v6_ = _dafny.MultiSet([d_415_v0_, d_415_v0_])
            d_422_v7_: D8
            d_422_v7_ = D8_DC21(len(default__.fm34(default__.fm0(globalState), self.f7, (self).fm32(self.f7, d_416_v1_, len(d_418_v3_), d_415_v0_, globalState), d_415_v0_, globalState)), d_420_v5_, (d_421_v6_).cardinality)
            d_415_v0_ = (d_415_v0_) + ((d_422_v7_).cf40)
            d_415_v0_ = d_415_v0_
            (self).f7 = self.f7
            (self).f7 = self.f7
        elif source1_.is_DC25:
            d_423_v8_: int
            d_423_v8_ = 319
            d_424_v9_: _dafny.Map
            d_424_v9_ = _dafny.Map({default__.safeModuloInt(d_423_v8_, d_423_v8_): self.f7})
            d_425_v10_: _dafny.Set
            d_425_v10_ = _dafny.Set({self.f7, True})
            d_426_v11_: str
            d_426_v11_ = _dafny.CodePoint('k')
            d_427_v12_: _dafny.Map
            d_427_v12_ = _dafny.Map({d_423_v8_: d_426_v11_})
            d_428_v13_: _dafny.Set
            d_428_v13_ = _dafny.Set({len(default__.fm35(len(d_425_v10_), d_423_v8_, globalState)), len(d_427_v12_)})
            d_424_v9_ = (_dafny.Map({d_423_v8_: self.f7})) | (_dafny.Map({default__.fm0(globalState): (self).fm32(False, d_428_v13_, d_423_v8_, d_423_v8_, globalState)}))
            d_429_v14_: _dafny.Seq
            d_429_v14_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_430_v15_: _dafny.Seq
            d_430_v15_ = _dafny.SeqWithoutIsStrInference([d_423_v8_, len(d_429_v14_)])
            d_431_v16_: _dafny.Map
            d_431_v16_ = _dafny.Map({self.f7: d_430_v15_})
            d_432_v17_: _dafny.Seq
            d_432_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkse"))
            d_433_v18_: _dafny.MultiSet
            d_433_v18_ = _dafny.MultiSet([self.f7, self.f7, self.f7, self.f7, self.f7])
            d_434_v19_: C0
            nw86_ = C0()
            nw86_.ctor__(self.f7, self.f7)
            d_434_v19_ = nw86_
            d_435_v20_: _dafny.Seq
            d_435_v20_ = _dafny.SeqWithoutIsStrInference([d_434_v19_, d_434_v19_])
            d_436_v21_: _dafny.Seq
            d_436_v21_ = _dafny.SeqWithoutIsStrInference([((d_429_v14_).set(default__.safeIndex(((d_433_v18_)[self.f7] if (self.f7) in (d_433_v18_) else len(d_435_v20_)), len(d_429_v14_)), self.f7)) + (d_429_v14_), d_429_v14_])
            rhs54_ = len(d_436_v21_)
            rhs55_ = default__.fm36(not(self.f7), (self.f7) and (self.f7), globalState)
            rhs56_ = ((d_432_v17_ if False else d_432_v17_) if (_dafny.Set({972})).isdisjoint(d_428_v13_) else d_432_v17_)
            rhs57_ = ((d_432_v17_).set(default__.safeIndex(d_423_v8_, len(d_432_v17_)), d_426_v11_)) + (_dafny.SeqWithoutIsStrInference([d_426_v11_ for d_437_i0_ in range(default__.abs(349))]))
            rhs58_ = default__.safeModuloInt(len(d_432_v17_), d_423_v8_)
            d_423_v8_ = rhs54_
            d_431_v16_ = rhs55_
            d_432_v17_ = rhs56_
            d_432_v17_ = rhs57_
            d_423_v8_ = rhs58_
            d_438_v22_: _dafny.Map
            d_438_v22_ = _dafny.Map({d_423_v8_: d_433_v18_})
            d_424_v9_ = (d_424_v9_).set(len((d_438_v22_)), (d_423_v8_) < ((0) - (d_423_v8_)))
            if (d_434_v19_).f12:
                (self).f7 = True
                d_432_v17_ = d_432_v17_
                d_439_v23_: _dafny.Map
                d_439_v23_ = _dafny.Map({(d_434_v19_).fm16(False, globalState): (0) - (d_423_v8_)})
                (d_434_v19_).f13 = not (((d_434_v19_).f12 if d_434_v19_.f13 else (d_434_v19_).f12)) or ((d_434_v19_.f13) not in ((d_439_v23_).set(not(not((d_429_v14_)[default__.safeIndex(d_423_v8_, len(d_429_v14_))])), d_423_v8_)))
                d_440_v24_: C0
                nw87_ = C0()
                nw87_.ctor__((d_434_v19_).f12, (d_434_v19_).f12)
                d_440_v24_ = nw87_
                (d_434_v19_).f13 = (d_434_v19_).f12
            elif True:
                d_441_v25_: _dafny.Array
                def lambda27_(d_442_v15_):
                    def lambda28_(d_443_i1_):
                        return d_442_v15_

                    return lambda28_

                init14_ = lambda27_(d_430_v15_)
                nw88_ = _dafny.Array(None, 13)
                for i0_14_ in range(nw88_.length(0)):
                    nw88_[i0_14_] = init14_(i0_14_)
                d_441_v25_ = nw88_
                index65_ = default__.safeIndex(105, (d_441_v25_).length(0))
                (d_441_v25_)[index65_] = _dafny.SeqWithoutIsStrInference([d_423_v8_])
                d_444_v26_: _dafny.Map
                d_444_v26_ = _dafny.Map({d_423_v8_: d_428_v13_})
                d_445_v28_: _dafny.Set
                def iife57_():
                    coll19_ = _dafny.Set()
                    compr_19_: int
                    for compr_19_ in (_dafny.Set({default__.fm0(globalState), d_423_v8_})).Elements:
                        d_446_v27_: int = compr_19_
                        if (d_446_v27_) in (_dafny.Set({default__.fm0(globalState), d_423_v8_})):
                            coll19_ = coll19_.union(_dafny.Set([(d_446_v27_) * (len(_dafny.Map({-79: True})))]))
                    return _dafny.Set(coll19_)
                d_445_v28_ = _dafny.Set({d_428_v13_, ((d_444_v26_)[d_423_v8_] if (d_423_v8_) in (d_444_v26_) else iife57_()
                ), d_428_v13_, (d_428_v13_) - (d_428_v13_)})
                d_447_v29_: _dafny.Array
                def lambda29_(d_448_i2_):
                    return False

                init15_ = lambda29_
                nw89_ = _dafny.Array(None, 23)
                for i0_15_ in range(nw89_.length(0)):
                    nw89_[i0_15_] = init15_(i0_15_)
                d_447_v29_ = nw89_
                d_449_v30_: _dafny.Map
                d_449_v30_ = _dafny.Map({d_428_v13_: d_434_v19_.f13})
                d_450_v32_: _dafny.Seq
                d_450_v32_ = _dafny.SeqWithoutIsStrInference([d_447_v29_, d_447_v29_, d_447_v29_, d_447_v29_])
                index66_ = default__.safeIndex(105, (d_441_v25_).length(0))
                def iife58_():
                    coll20_ = _dafny.Set()
                    compr_20_: _dafny.Set
                    for compr_20_ in (d_449_v30_).keys.Elements:
                        d_451_v31_: _dafny.Set = compr_20_
                        if (d_451_v31_) in (d_449_v30_):
                            coll20_ = coll20_.union(_dafny.Set([d_451_v31_]))
                    return _dafny.Set(coll20_)
                rhs59_ = d_430_v15_
                rhs60_ = d_426_v11_
                rhs61_ = ((d_445_v28_) | (iife58_()
                )).intersection(_dafny.Set({d_428_v13_}))
                rhs62_ = d_426_v11_
                rhs63_ = (d_450_v32_)[default__.safeIndex(d_423_v8_, len(d_450_v32_))]
                lhs41_ = d_441_v25_
                lhs42_ = default__.safeIndex(105, (d_441_v25_).length(0))
                lhs41_[lhs42_] = rhs59_
                d_426_v11_ = rhs60_
                d_445_v28_ = rhs61_
                d_426_v11_ = rhs62_
                d_447_v29_ = rhs63_
                d_452_v33_: _dafny.Map
                d_452_v33_ = _dafny.Map({d_426_v11_: d_434_v19_.f13})
                d_453_v34_: D0
                d_453_v34_ = D0_DC0(d_423_v8_)
                rhs64_ = (d_434_v19_.f13) == (self.f7)
                rhs65_ = d_434_v19_.f13
                rhs66_ = default__.fm1(_dafny.SeqWithoutIsStrInference([d_453_v34_]), d_423_v8_, default__.fm0(globalState), globalState)
                rhs67_ = d_452_v33_
                rhs68_ = default__.fm0(globalState)
                lhs43_ = self
                lhs44_ = d_434_v19_
                lhs45_ = d_434_v19_
                lhs43_.f7 = rhs64_
                lhs44_.f13 = rhs65_
                lhs45_.f13 = rhs66_
                d_452_v33_ = rhs67_
                d_423_v8_ = rhs68_
                index67_ = default__.safeIndex(91, (d_447_v29_).length(0))
                (d_447_v29_)[index67_] = self.f7
                index68_ = default__.safeIndex(91, (d_447_v29_).length(0))
                (d_447_v29_)[index68_] = self.f7
                index69_ = default__.safeIndex(91, (d_447_v29_).length(0))
                (d_447_v29_)[index69_] = (d_432_v17_) <= (d_432_v17_)
                d_454_v35_: _dafny.Array
                nw90_ = _dafny.Array(int(0), 23)
                d_454_v35_ = nw90_
                index70_ = default__.safeIndex(217, (d_454_v35_).length(0))
                (d_454_v35_)[index70_] = (default__.fm0(globalState)) * (d_423_v8_)
                index71_ = default__.safeIndex(217, (d_454_v35_).length(0))
                (d_454_v35_)[index71_] = d_423_v8_
        elif source1_.is_DC26:
            d_455___mcc_h0_ = source1_.cf45
            d_456_cf45_ = d_455___mcc_h0_
            d_456_cf45_ = d_456_cf45_
            d_457_v36_: _dafny.Map
            d_457_v36_ = _dafny.Map({self.f7: not(self.f7)})
            d_458_v37_: int
            d_458_v37_ = 941
            d_457_v36_ = (d_457_v36_).set((d_458_v37_) != (d_458_v37_), self.f7)
            d_459_v38_: _dafny.Map
            d_459_v38_ = _dafny.Map({(617) + (d_458_v37_): d_458_v37_})
            d_459_v38_ = _dafny.Map({d_458_v37_: 566})
            d_460_v39_: _dafny.Array
            nw91_ = _dafny.Array(D1.default()(), 6)
            d_460_v39_ = nw91_
            d_460_v39_ = d_460_v39_
        elif True:
            d_461___mcc_h1_ = source1_.cf44
            d_462_cf44_ = d_461___mcc_h1_
            d_463_v40_: int
            d_463_v40_ = 458
            d_464_v41_: _dafny.Map
            d_464_v41_ = _dafny.Map({d_463_v40_: d_463_v40_})
            d_465_v42_: _dafny.Seq
            d_465_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "befgtwleg"))
            d_466_v43_: C0
            nw92_ = C0()
            nw92_.ctor__((((d_464_v41_)[d_463_v40_] if (d_463_v40_) in (d_464_v41_) else len(d_465_v42_))) == (d_463_v40_), self.f7)
            d_466_v43_ = nw92_
            index72_ = default__.safeIndex(404, (d_462_cf44_).length(0))
            (d_462_cf44_)[index72_] = d_463_v40_
            index73_ = default__.safeIndex(404, (d_462_cf44_).length(0))
            (d_462_cf44_)[index73_] = ((0) - ((_dafny.MultiSet([d_463_v40_, d_463_v40_])).cardinality)) - (d_463_v40_)
            d_467_v44_: _dafny.Array
            nw93_ = _dafny.Array(D0.default()(), 21)
            d_467_v44_ = nw93_
            d_468_v45_: D0
            d_468_v45_ = D0_DC0((d_462_cf44_)[default__.safeIndex(404, (d_462_cf44_).length(0))])
            index74_ = default__.safeIndex(643, (d_467_v44_).length(0))
            (d_467_v44_)[index74_] = d_468_v45_
            pat_let_tv9_ = d_462_cf44_
            pat_let_tv10_ = d_462_cf44_
            index75_ = default__.safeIndex(643, (d_467_v44_).length(0))
            def iife59_(_pat_let19_0):
                def iife60_(d_469_dt__update__tmp_h0_):
                    def iife61_(_pat_let20_0):
                        def iife62_(d_470_dt__update_hcf0_h0_):
                            return D0_DC0(d_470_dt__update_hcf0_h0_)
                        return iife62_(_pat_let20_0)
                    return iife61_((pat_let_tv10_)[default__.safeIndex(404, (pat_let_tv9_).length(0))])
                return iife60_(_pat_let19_0)
            rhs69_ = iife59_(d_468_v45_)
            rhs70_ = default__.safeModuloInt((0) - (d_463_v40_), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edflhdb")))) + (len(d_465_v42_)))
            lhs46_ = d_467_v44_
            lhs47_ = default__.safeIndex(643, (d_467_v44_).length(0))
            lhs46_[lhs47_] = rhs69_
            d_463_v40_ = rhs70_
            d_471_v46_: str
            d_471_v46_ = _dafny.CodePoint('w')
            d_472_v47_: _dafny.Set
            d_472_v47_ = _dafny.Set({d_471_v46_, d_471_v46_})
            d_473_v48_: _dafny.Set
            d_473_v48_ = d_472_v47_
            d_474_v49_: _dafny.Map
            d_474_v49_ = _dafny.Map({(d_473_v48_): d_466_v43_.f13})
            d_475_v50_: _dafny.Map
            d_475_v50_ = _dafny.Map({(d_466_v43_.f13) and (True): d_474_v49_})
            d_475_v50_ = (d_475_v50_).set(self.f7, d_474_v49_)
        d_476_v51_: C0
        nw94_ = C0()
        nw94_.ctor__((False) == (False), self.f7)
        d_476_v51_ = nw94_
        d_477_v52_: _dafny.Array
        nw95_ = _dafny.Array(int(0), 14)
        d_477_v52_ = nw95_
        d_478_v53_: _dafny.Map
        d_478_v53_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gca"))): self.f7})
        d_479_v54_: int
        d_479_v54_ = 259
        d_480_v55_: _dafny.Map
        d_480_v55_ = _dafny.Map({default__.fm37(True, D9_DC25(), globalState): ((d_478_v53_)[d_479_v54_] if (d_479_v54_) in (d_478_v53_) else self.f7)})
        index76_ = default__.safeIndex(138, (d_477_v52_).length(0))
        (d_477_v52_)[index76_] = (len(d_480_v55_)) * (d_479_v54_)
        index77_ = default__.safeIndex(138, (d_477_v52_).length(0))
        (d_477_v52_)[index77_] = (0) - ((d_479_v54_) - (d_479_v54_))
        d_481_v56_: _dafny.Array
        def lambda30_(d_482_v51_):
            def lambda31_(d_483_i3_):
                return _dafny.SeqWithoutIsStrInference([_dafny.Set({d_482_v51_.f13}), _dafny.Set({d_482_v51_.f13, self.f7, (d_482_v51_).f12, True}), _dafny.Set({(d_482_v51_).f12})])

            return lambda31_

        init16_ = lambda30_(d_476_v51_)
        nw96_ = _dafny.Array(None, 6)
        for i0_16_ in range(nw96_.length(0)):
            nw96_[i0_16_] = init16_(i0_16_)
        d_481_v56_ = nw96_
        index78_ = default__.safeIndex(929, (d_481_v56_).length(0))
        (d_481_v56_)[index78_] = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_476_v51_.f13, self.f7, (d_476_v51_).f12}), _dafny.Set({False, (d_476_v51_).f12})])
        index79_ = default__.safeIndex(929, (d_481_v56_).length(0))
        (d_481_v56_)[index79_] = default__.fm38((d_476_v51_).f12, globalState)
        d_484_v57_: _dafny.Seq
        d_484_v57_ = _dafny.SeqWithoutIsStrInference([d_479_v54_])
        if ((d_484_v57_)[default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_484_v57_))]) <= (d_479_v54_):
            d_485_v58_: _dafny.Map
            d_485_v58_ = _dafny.Map({(d_476_v51_).f12: 238})
            d_486_v59_: _dafny.MultiSet
            d_486_v59_ = _dafny.MultiSet([d_485_v58_])
            d_487_v60_: _dafny.MultiSet
            d_487_v60_ = _dafny.MultiSet([(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]])
            d_488_v61_: _dafny.Seq
            d_488_v61_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({self.f7: (d_487_v60_).cardinality}), _dafny.Map({d_476_v51_.f13: (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]})])
            d_489_v62_: _dafny.Seq
            d_489_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_488_v61_)])
            (self).f7 = (d_486_v59_).isdisjoint(((d_489_v62_)[default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_489_v62_))]) - (d_486_v59_))
            index80_ = default__.safeIndex(138, (d_477_v52_).length(0))
            (d_477_v52_)[index80_] = d_479_v54_
            d_490_v63_: _dafny.Seq
            d_490_v63_ = _dafny.SeqWithoutIsStrInference([not(self.f7)])
            (self).f7 = (d_490_v63_) <= ((d_490_v63_ if self.f7 else d_490_v63_))
            d_491_v64_: _dafny.Seq
            d_491_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwepdcn"))
            d_492_v65_: str
            d_492_v65_ = _dafny.CodePoint('g')
            rhs71_ = (d_479_v54_) <= (d_479_v54_)
            rhs72_ = not((d_476_v51_.f13 if d_476_v51_.f13 else d_476_v51_.f13))
            rhs73_ = (d_476_v51_.f13 if d_476_v51_.f13 else d_476_v51_.f13)
            rhs74_ = ((d_491_v64_).set(default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_491_v64_)), d_492_v65_)) + (d_491_v64_)
            rhs75_ = (d_492_v65_ if self.f7 else d_492_v65_)
            lhs48_ = self
            lhs49_ = d_476_v51_
            lhs50_ = d_476_v51_
            lhs48_.f7 = rhs71_
            lhs49_.f13 = rhs72_
            lhs50_.f13 = rhs73_
            d_491_v64_ = rhs74_
            d_492_v65_ = rhs75_
            d_479_v54_ = default__.fm0(globalState)
        elif True:
            (d_476_v51_).f13 = (d_476_v51_).f12
            d_493_v66_: C0
            nw97_ = C0()
            nw97_.ctor__(d_476_v51_.f13, d_476_v51_.f13)
            d_493_v66_ = nw97_
            d_494_v67_: D3
            d_494_v67_ = D3_DC11(-973, 984, d_479_v54_, d_493_v66_.f13)
            d_495_v68_: D3
            d_495_v68_ = D3_DC12(d_494_v67_)
            pat_let_tv11_ = d_494_v67_
            def iife63_(_pat_let21_0):
                def iife64_(d_496_dt__update__tmp_h1_):
                    def iife65_(_pat_let22_0):
                        def iife66_(d_497_dt__update_hcf26_h0_):
                            return D3_DC12(d_497_dt__update_hcf26_h0_)
                        return iife66_(_pat_let22_0)
                    return iife65_(pat_let_tv11_)
                return iife64_(_pat_let21_0)
            d_495_v68_ = iife63_(d_495_v68_)
            d_498_v69_: str
            d_498_v69_ = _dafny.CodePoint('s')
            d_499_v70_: _dafny.Seq
            d_499_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnehap"))
            if (d_498_v69_) in (d_499_v70_):
                d_500_v71_: _dafny.Set
                d_500_v71_ = _dafny.Set({-964})
                d_501_v72_: _dafny.Set
                d_501_v72_ = _dafny.Set({d_500_v71_})
                d_502_v73_: _dafny.Map
                d_502_v73_ = _dafny.Map({len(d_501_v72_): d_479_v54_})
                d_503_v74_: _dafny.Map
                d_503_v74_ = _dafny.Map({(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]: d_502_v73_})
                d_503_v74_ = (d_503_v74_) | (_dafny.Map({d_479_v54_: d_502_v73_}))
                d_504_v75_: _dafny.Array
                nw98_ = _dafny.Array(False, 21)
                d_504_v75_ = nw98_
                d_505_v76_: _dafny.Seq
                d_505_v76_ = _dafny.SeqWithoutIsStrInference([d_504_v75_, d_504_v75_])
                d_506_v77_: _dafny.Array
                nw99_ = _dafny.Array(None, 8)
                nw99_[int(0)] = d_493_v66_.f13
                nw99_[int(1)] = d_476_v51_.f13
                nw99_[int(2)] = (d_476_v51_).fm16((d_476_v51_).f12, globalState)
                nw99_[int(3)] = d_476_v51_.f13
                nw99_[int(4)] = (d_476_v51_).fm16((d_476_v51_).f12, globalState)
                nw99_[int(5)] = (d_476_v51_).f12
                nw99_[int(6)] = d_493_v66_.f13
                nw99_[int(7)] = (d_504_v75_) not in (d_505_v76_)
                d_506_v77_ = nw99_
                d_504_v75_ = d_506_v77_
                d_507_v78_: _dafny.MultiSet
                d_507_v78_ = _dafny.MultiSet([d_499_v70_])
                d_508_v79_: D0
                d_508_v79_ = D0_DC0((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                d_509_v80_: _dafny.Map
                d_509_v80_ = _dafny.Map({d_508_v79_: d_479_v54_})
                d_510_v81_: _dafny.Map
                d_510_v81_ = _dafny.Map({True: d_498_v69_})
                d_511_v82_: _dafny.MultiSet
                d_511_v82_ = _dafny.MultiSet([_dafny.Map({(D2_DC9(len(_dafny.SeqWithoutIsStrInference([d_479_v54_ for d_512_i4_ in range(default__.abs(-963))])), (d_507_v78_).cardinality, d_499_v70_, (d_476_v51_).f12, d_509_v80_)).cf19: d_498_v69_}), d_510_v81_])
                d_513_v83_: C0
                nw100_ = C0()
                nw100_.ctor__((d_511_v82_).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_510_v81_]))), not((d_476_v51_).f12))
                d_513_v83_ = nw100_
                d_514_v84_: _dafny.Set
                d_514_v84_ = _dafny.Set({d_477_v52_, d_477_v52_})
                d_515_v85_: D0
                d_515_v85_ = D0_DC2(((d_493_v66_).f12) or ((self).fm32(not((d_476_v51_).f12), d_500_v71_, d_479_v54_, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], globalState)), default__.safeModuloInt((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]), d_514_v84_, len(d_500_v71_))
                d_515_v85_ = d_515_v85_
                d_516_v86_: _dafny.Set
                d_516_v86_ = _dafny.Set({(d_476_v51_).f12, False})
                d_517_v87_: _dafny.Seq
                d_517_v87_ = _dafny.SeqWithoutIsStrInference([not(d_493_v66_.f13), d_493_v66_.f13, d_476_v51_.f13])
                d_479_v54_ = (len(d_516_v86_) if self.f7 else (d_479_v54_) + (len(_dafny.Set({d_479_v54_, 340, len(d_517_v87_)}))))
            elif True:
                (self).f7 = self.f7
                d_518_v88_: D1
                d_518_v88_ = D1_DC4(self.f7, d_476_v51_.f13)
                d_519_v89_: _dafny.Map
                d_519_v89_ = _dafny.Map({self.f7: d_476_v51_.f13})
                d_520_v90_: C0
                nw101_ = C0()
                nw101_.ctor__((d_476_v51_).f12, (_dafny.Map({(d_518_v88_).cf7: d_476_v51_.f13})) != (d_519_v89_))
                d_520_v90_ = nw101_
                index81_ = default__.safeIndex(138, (d_477_v52_).length(0))
                rhs76_ = True
                rhs77_ = d_479_v54_
                lhs51_ = self
                lhs52_ = d_477_v52_
                lhs53_ = default__.safeIndex(138, (d_477_v52_).length(0))
                lhs51_.f7 = rhs76_
                lhs52_[lhs53_] = rhs77_
                d_521_v91_: _dafny.Set
                d_521_v91_ = _dafny.Set({d_476_v51_.f13, d_476_v51_.f13})
                d_522_v92_: C0
                nw102_ = C0()
                nw102_.ctor__(((d_476_v51_).f12) not in (d_521_v91_), self.f7)
                d_522_v92_ = nw102_
                d_523_v93_: _dafny.Seq
                d_523_v93_ = _dafny.SeqWithoutIsStrInference([d_499_v70_])
                d_524_v94_: _dafny.Set
                d_524_v94_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cr")), (d_523_v93_)[default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_523_v93_))]})
                index82_ = default__.safeIndex(138, (d_477_v52_).length(0))
                rhs78_ = (0) - (((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) * ((d_479_v54_) * ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])))
                rhs79_ = (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]
                rhs80_ = ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) + (d_479_v54_)
                rhs81_ = (d_524_v94_).isdisjoint(_dafny.Set({d_499_v70_, d_499_v70_, d_499_v70_, d_499_v70_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ct"))}))
                lhs54_ = d_477_v52_
                lhs55_ = default__.safeIndex(138, (d_477_v52_).length(0))
                lhs56_ = d_493_v66_
                d_479_v54_ = rhs78_
                lhs54_[lhs55_] = rhs79_
                d_479_v54_ = rhs80_
                lhs56_.f13 = rhs81_
            if not(((d_476_v51_).f12) == ((d_493_v66_).fm16(d_476_v51_.f13, globalState))):
                d_525_v95_: _dafny.Array
                def lambda32_(d_526_v69_):
                    def lambda33_(d_527_i5_):
                        return d_526_v69_

                    return lambda33_

                init17_ = lambda32_(d_498_v69_)
                nw103_ = _dafny.Array(None, 29)
                for i0_17_ in range(nw103_.length(0)):
                    nw103_[i0_17_] = init17_(i0_17_)
                d_525_v95_ = nw103_
                index83_ = default__.safeIndex(748, (d_525_v95_).length(0))
                (d_525_v95_)[index83_] = d_498_v69_
                index84_ = default__.safeIndex(748, (d_525_v95_).length(0))
                (d_525_v95_)[index84_] = d_498_v69_
                d_528_v96_: _dafny.Set
                d_528_v96_ = _dafny.Set({d_476_v51_.f13, d_476_v51_.f13})
                d_529_v98_: _dafny.MultiSet
                def iife67_():
                    coll21_ = _dafny.Set()
                    compr_21_: int
                    for compr_21_ in _dafny.IntegerRange(168, 983):
                        d_530_v97_: int = compr_21_
                        if ((168) <= (d_530_v97_)) and ((d_530_v97_) < (983)):
                            coll21_ = coll21_.union(_dafny.Set([default__.safeDivisionInt(d_530_v97_, d_479_v54_)]))
                    return _dafny.Set(coll21_)
                d_529_v98_ = _dafny.MultiSet([len(d_528_v96_), d_479_v54_, len(iife67_()
                ), d_479_v54_])
                d_531_v99_: _dafny.Map
                d_531_v99_ = _dafny.Map({(d_476_v51_).f12: d_529_v98_})
                d_532_v100_: D0
                d_532_v100_ = D0_DC0(498)
                d_533_v101_: _dafny.Seq
                d_533_v101_ = _dafny.SeqWithoutIsStrInference([d_532_v100_])
                d_534_v102_: _dafny.Set
                d_534_v102_ = _dafny.Set({d_477_v52_})
                d_535_v103_: D0
                d_535_v103_ = D0_DC2(d_476_v51_.f13, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], d_534_v102_, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                d_536_v104_: _dafny.Array
                nw104_ = _dafny.Array(None, 11)
                nw104_[int(0)] = d_476_v51_.f13
                nw104_[int(1)] = self.f7
                nw104_[int(2)] = not(d_476_v51_.f13)
                nw104_[int(3)] = not((d_529_v98_).isdisjoint(((d_531_v99_)[d_493_v66_.f13] if (d_493_v66_.f13) in (d_531_v99_) else d_529_v98_)))
                nw104_[int(4)] = not(d_476_v51_.f13)
                nw104_[int(5)] = (d_476_v51_).f12
                nw104_[int(6)] = (d_493_v66_).f12
                nw104_[int(7)] = default__.fm1(d_533_v101_, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], -545, globalState)
                nw104_[int(8)] = (d_476_v51_.f13 if (d_535_v103_).cf2 else not(d_493_v66_.f13))
                nw104_[int(9)] = self.f7
                nw104_[int(10)] = (d_476_v51_).f12
                d_536_v104_ = nw104_
                d_537_v105_: _dafny.Seq
                d_537_v105_ = _dafny.SeqWithoutIsStrInference([True])
                index85_ = default__.safeIndex(630, (d_536_v104_).length(0))
                (d_536_v104_)[index85_] = (d_537_v105_)[default__.safeIndex(d_479_v54_, len(d_537_v105_))]
                d_538_v106_: _dafny.Array
                def lambda34_(d_539_v57_, d_540_v52_, d_541_v54_):
                    def lambda35_(d_542_i6_):
                        return (d_539_v57_).set(default__.safeIndex((d_540_v52_)[default__.safeIndex(138, (d_540_v52_).length(0))], len(d_539_v57_)), d_541_v54_)

                    return lambda35_

                init18_ = lambda34_(d_484_v57_, d_477_v52_, d_479_v54_)
                nw105_ = _dafny.Array(None, 19)
                for i0_18_ in range(nw105_.length(0)):
                    nw105_[i0_18_] = init18_(i0_18_)
                d_538_v106_ = nw105_
                d_543_v107_: _dafny.Map
                d_543_v107_ = _dafny.Map({(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]: d_499_v70_})
                index86_ = default__.safeIndex(488, (d_538_v106_).length(0))
                (d_538_v106_)[index86_] = ((d_484_v57_) + (d_484_v57_)).set(default__.safeIndex((0) - (len(((d_543_v107_)[d_479_v54_] if (d_479_v54_) in (d_543_v107_) else (d_499_v70_).set(default__.safeIndex(d_479_v54_, len(d_499_v70_)), (d_525_v95_)[default__.safeIndex(748, (d_525_v95_).length(0))])))), len((d_484_v57_) + (d_484_v57_))), default__.fm0(globalState))
                index87_ = default__.safeIndex(138, (d_477_v52_).length(0))
                index88_ = default__.safeIndex(630, (d_536_v104_).length(0))
                index89_ = default__.safeIndex(488, (d_538_v106_).length(0))
                rhs82_ = len(d_499_v70_)
                rhs83_ = ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) <= (d_479_v54_)
                rhs84_ = (d_484_v57_) + (d_484_v57_)
                lhs57_ = d_477_v52_
                lhs58_ = default__.safeIndex(138, (d_477_v52_).length(0))
                lhs59_ = d_536_v104_
                lhs60_ = default__.safeIndex(630, (d_536_v104_).length(0))
                lhs61_ = d_538_v106_
                lhs62_ = default__.safeIndex(488, (d_538_v106_).length(0))
                lhs57_[lhs58_] = rhs82_
                lhs59_[lhs60_] = rhs83_
                lhs61_[lhs62_] = rhs84_
                d_544_v108_: D3
                d_544_v108_ = D3_DC10(d_536_v104_)
                d_536_v104_ = (d_544_v108_).cf21
                d_545_v109_: _dafny.Array
                def lambda36_(d_546_v105_):
                    def lambda37_(d_547_i7_):
                        return d_546_v105_

                    return lambda37_

                init19_ = lambda36_(d_537_v105_)
                nw106_ = _dafny.Array(None, 2)
                for i0_19_ in range(nw106_.length(0)):
                    nw106_[i0_19_] = init19_(i0_19_)
                d_545_v109_ = nw106_
                index90_ = default__.safeIndex(950, (d_545_v109_).length(0))
                (d_545_v109_)[index90_] = (d_476_v51_).fm17(len(d_499_v70_), (d_493_v66_).f12, globalState)
                index91_ = default__.safeIndex(950, (d_545_v109_).length(0))
                (d_545_v109_)[index91_] = (d_493_v66_).fm17((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], d_476_v51_.f13, globalState)
                d_479_v54_ = d_479_v54_
            elif True:
                d_548_v110_: _dafny.Seq
                d_548_v110_ = _dafny.SeqWithoutIsStrInference([False])
                d_549_v111_: _dafny.Map
                d_549_v111_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwb")): d_548_v110_})
                d_479_v54_ = len(((d_549_v111_).set(d_499_v70_, d_548_v110_)) | (d_549_v111_))
                d_550_v112_: _dafny.Array
                nw107_ = _dafny.Array(False, 16)
                d_550_v112_ = nw107_
                d_551_v113_: _dafny.Set
                d_551_v113_ = _dafny.Set({d_550_v112_, d_550_v112_, d_550_v112_, d_550_v112_, d_550_v112_})
                d_552_v114_: _dafny.Map
                d_552_v114_ = _dafny.Map({d_498_v69_: d_551_v113_})
                d_552_v114_ = (d_552_v114_).set(d_498_v69_, (d_551_v113_) | (d_551_v113_))
                d_479_v54_ = ((len(d_484_v57_)) * ((0) - ((0) - ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])))) * (d_479_v54_)
                index92_ = default__.safeIndex(138, (d_477_v52_).length(0))
                (d_477_v52_)[index92_] = ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) - (d_479_v54_)
                default__.m0((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], d_479_v54_, globalState)
        d_553_v115_: D0
        d_553_v115_ = D0_DC0((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
        d_554_v116_: _dafny.Seq
        d_554_v116_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "to"))
        d_555_v117_: _dafny.Seq
        d_555_v117_ = _dafny.SeqWithoutIsStrInference([d_553_v115_, D0_DC0(len(d_554_v116_)), d_553_v115_])
        d_556_v118_: str
        d_556_v118_ = _dafny.CodePoint('e')
        d_557_v119_: _dafny.Set
        d_557_v119_ = _dafny.Set({d_556_v118_, d_556_v118_, d_556_v118_, d_556_v118_, d_556_v118_})
        d_558_v120_: _dafny.Set
        d_558_v120_ = d_557_v119_
        d_559_v121_: _dafny.Map
        d_559_v121_ = _dafny.Map({d_558_v120_: (d_476_v51_).f12})
        if default__.fm1((d_555_v117_) + (d_555_v117_), len(d_559_v121_), (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], globalState):
            d_560_v122_: _dafny.Seq
            d_560_v122_ = _dafny.SeqWithoutIsStrInference([d_554_v116_, d_554_v116_])
            if ((d_560_v122_) + (d_560_v122_)) < (_dafny.SeqWithoutIsStrInference([d_554_v116_ for d_561_i8_ in range(default__.abs(275))])):
                d_562_v123_: _dafny.Map
                d_562_v123_ = _dafny.Map({d_553_v115_: d_479_v54_})
                d_563_v124_: D2
                d_563_v124_ = D2_DC9((d_479_v54_) * (d_479_v54_), d_479_v54_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")), (d_476_v51_).f12, d_562_v123_)
                index93_ = default__.safeIndex(138, (d_477_v52_).length(0))
                rhs85_ = default__.safeDivisionInt(d_479_v54_, (0) - (default__.safeDivisionInt((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])))
                rhs86_ = d_476_v51_.f13
                rhs87_ = d_563_v124_
                lhs63_ = d_477_v52_
                lhs64_ = default__.safeIndex(138, (d_477_v52_).length(0))
                lhs65_ = d_476_v51_
                lhs63_[lhs64_] = rhs85_
                lhs65_.f13 = rhs86_
                d_563_v124_ = rhs87_
                d_564_v125_: _dafny.Map
                d_564_v125_ = _dafny.Map({((d_476_v51_).f12) == ((d_476_v51_).f12): (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]})
                d_479_v54_ = ((d_564_v125_)[d_476_v51_.f13] if (d_476_v51_.f13) in (d_564_v125_) else (default__.fm0(globalState) if d_476_v51_.f13 else (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]))
                (self).f7 = self.f7
                d_565_v126_: D2
                d_565_v126_ = D2_DC8((d_476_v51_).f12, d_554_v116_)
                rhs88_ = default__.safeDivisionInt(d_479_v54_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "giukdrl"))) + (d_554_v116_)))
                rhs89_ = d_565_v126_
                rhs90_ = d_476_v51_.f13
                rhs91_ = ((d_554_v116_) + (d_554_v116_)) + (((d_554_v116_).set(default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_554_v116_)), d_556_v118_)) + (d_554_v116_))
                lhs66_ = d_476_v51_
                d_479_v54_ = rhs88_
                d_565_v126_ = rhs89_
                lhs66_.f13 = rhs90_
                d_554_v116_ = rhs91_
                d_566_v127_: D3
                d_566_v127_ = D3_DC11((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], d_479_v54_, d_479_v54_, d_476_v51_.f13)
                d_567_v128_: _dafny.Seq
                d_567_v128_ = _dafny.SeqWithoutIsStrInference([(d_566_v127_).cf25])
                d_568_v129_: _dafny.MultiSet
                d_568_v129_ = _dafny.MultiSet([d_479_v54_])
                d_569_v130_: _dafny.Seq
                d_569_v130_ = _dafny.SeqWithoutIsStrInference([d_568_v129_])
                d_570_v131_: _dafny.Array
                nw108_ = _dafny.Array(None, 24)
                nw108_[int(0)] = True
                nw108_[int(1)] = (d_476_v51_).f12
                nw108_[int(2)] = (d_476_v51_).f12
                nw108_[int(3)] = d_476_v51_.f13
                nw108_[int(4)] = (d_567_v128_)[default__.safeIndex(d_479_v54_, len(d_567_v128_))]
                nw108_[int(5)] = (d_476_v51_).f12
                nw108_[int(6)] = (d_476_v51_).f12
                nw108_[int(7)] = d_476_v51_.f13
                nw108_[int(8)] = ((d_569_v130_)[default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_569_v130_))]).ispropersubset((_dafny.MultiSet(d_484_v57_)).set((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], default__.abs((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])))
                nw108_[int(9)] = (d_476_v51_).f12
                nw108_[int(10)] = (d_476_v51_).f12
                nw108_[int(11)] = self.f7
                nw108_[int(12)] = ((default__.fm39((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], (d_476_v51_).f12, globalState)).cf10) == (_dafny.SeqWithoutIsStrInference([default__.fm0(globalState)]))
                nw108_[int(13)] = (d_479_v54_) >= (d_479_v54_)
                nw108_[int(14)] = (d_476_v51_).f12
                nw108_[int(15)] = self.f7
                nw108_[int(16)] = (d_476_v51_).f12
                nw108_[int(17)] = ((d_476_v51_).f12 if self.f7 else (d_476_v51_).f12)
                nw108_[int(18)] = d_476_v51_.f13
                nw108_[int(19)] = (d_476_v51_).f12
                nw108_[int(20)] = (len(d_554_v116_)) <= (d_479_v54_)
                nw108_[int(21)] = (d_476_v51_).f12
                nw108_[int(22)] = False
                nw108_[int(23)] = (d_568_v129_).ispropersubset(_dafny.MultiSet([d_479_v54_]))
                d_570_v131_ = nw108_
                index94_ = default__.safeIndex(783, (d_570_v131_).length(0))
                (d_570_v131_)[index94_] = (d_476_v51_).f12
                index95_ = default__.safeIndex(783, (d_570_v131_).length(0))
                (d_570_v131_)[index95_] = ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) > (d_479_v54_)
            elif True:
                index96_ = default__.safeIndex(138, (d_477_v52_).length(0))
                (d_477_v52_)[index96_] = (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]
                d_571_v132_: _dafny.Seq
                d_571_v132_ = _dafny.SeqWithoutIsStrInference([(d_476_v51_).f12])
                d_572_v133_: _dafny.MultiSet
                d_572_v133_ = _dafny.MultiSet([d_571_v132_])
                d_573_v134_: D2
                d_573_v134_ = D2_DC8(False, _dafny.SeqWithoutIsStrInference([d_556_v118_ for d_574_i9_ in range(default__.abs(368))]))
                d_575_v135_: _dafny.Map
                d_575_v135_ = _dafny.Map({(((default__.fm35((d_572_v133_).cardinality, -39, globalState)) + ((d_573_v134_).cf15)).set(default__.safeIndex(default__.fm0(globalState), len((default__.fm35((d_572_v133_).cardinality, -39, globalState)) + ((d_573_v134_).cf15))), _dafny.CodePoint('h'))).set(default__.safeIndex(d_479_v54_, len(((default__.fm35((d_572_v133_).cardinality, -39, globalState)) + ((d_573_v134_).cf15)).set(default__.safeIndex(default__.fm0(globalState), len((default__.fm35((d_572_v133_).cardinality, -39, globalState)) + ((d_573_v134_).cf15))), _dafny.CodePoint('h')))), d_556_v118_): d_484_v57_})
                d_575_v135_ = ((d_575_v135_) | (d_575_v135_) if d_476_v51_.f13 else d_575_v135_)
                d_576_v136_: C0
                nw109_ = C0()
                nw109_.ctor__(self.f7, (d_484_v57_) < (d_484_v57_))
                d_576_v136_ = nw109_
                d_479_v54_ = (0) - ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                index97_ = default__.safeIndex(138, (d_477_v52_).length(0))
                (d_477_v52_)[index97_] = ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))] if (d_576_v136_).f12 else (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
            if (d_476_v51_).f12:
                d_577_v137_: _dafny.Map
                d_577_v137_ = _dafny.Map({True: (d_476_v51_).f12})
                d_577_v137_ = (d_577_v137_).set(self.f7, default__.fm1(d_555_v117_, d_479_v54_, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], globalState))
                d_578_v138_: D6
                d_578_v138_ = D6_DC17((d_476_v51_).f12, (0) - ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]), d_556_v118_, d_556_v118_)
                d_579_v139_: D6
                d_579_v139_ = D6_DC18(d_578_v138_)
                d_580_v140_: _dafny.Map
                d_580_v140_ = _dafny.Map({(d_476_v51_).f12: default__.fm40(((d_478_v53_)[(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]] if ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) in (d_478_v53_) else self.f7), d_579_v139_, d_479_v54_, globalState)})
                d_581_v141_: D9
                d_581_v141_ = D9_DC25()
                d_580_v140_ = (d_580_v140_).set(((d_478_v53_)[len(_dafny.Set({d_554_v116_}))] if (len(_dafny.Set({d_554_v116_}))) in (d_478_v53_) else d_476_v51_.f13), d_581_v141_)
                (self).f7 = False
                d_582_v142_: _dafny.Array
                nw110_ = _dafny.Array(False, 16)
                d_582_v142_ = nw110_
                d_583_v143_: D3
                d_583_v143_ = D3_DC10(d_582_v142_)
                d_584_v144_: D3
                d_584_v144_ = D3_DC12(d_583_v143_)
                pat_let_tv12_ = d_583_v143_
                def iife68_(_pat_let23_0):
                    def iife69_(d_585_dt__update__tmp_h2_):
                        def iife70_(_pat_let24_0):
                            def iife71_(d_586_dt__update_hcf26_h1_):
                                return D3_DC12(d_586_dt__update_hcf26_h1_)
                            return iife71_(_pat_let24_0)
                        return iife70_(pat_let_tv12_)
                    return iife69_(_pat_let23_0)
                d_584_v144_ = iife68_(d_584_v144_)
                d_587_v145_: _dafny.Array
                nw111_ = _dafny.Array(None, 1)
                nw111_[int(0)] = d_554_v116_
                d_587_v145_ = nw111_
                index98_ = default__.safeIndex(29, (d_587_v145_).length(0))
                (d_587_v145_)[index98_] = d_554_v116_
                index99_ = default__.safeIndex(314, (d_582_v142_).length(0))
                (d_582_v142_)[index99_] = (default__.fm2((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], (d_476_v51_).f12, globalState)).issubset(_dafny.Set({(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]}))
                index100_ = default__.safeIndex(29, (d_587_v145_).length(0))
                index101_ = default__.safeIndex(314, (d_582_v142_).length(0))
                rhs92_ = d_554_v116_
                rhs93_ = d_476_v51_.f13
                lhs67_ = d_587_v145_
                lhs68_ = default__.safeIndex(29, (d_587_v145_).length(0))
                lhs69_ = d_582_v142_
                lhs70_ = default__.safeIndex(314, (d_582_v142_).length(0))
                lhs67_[lhs68_] = rhs92_
                lhs69_[lhs70_] = rhs93_
            elif True:
                d_588_v146_: _dafny.Array
                nw112_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_588_v146_ = nw112_
                d_589_v147_: _dafny.Array
                def lambda38_(d_590_v51_):
                    def lambda39_(d_591_i10_):
                        return _dafny.Map({d_590_v51_.f13: (d_590_v51_).f12})

                    return lambda39_

                init20_ = lambda38_(d_476_v51_)
                nw113_ = _dafny.Array(None, 19)
                for i0_20_ in range(nw113_.length(0)):
                    nw113_[i0_20_] = init20_(i0_20_)
                d_589_v147_ = nw113_
                d_592_v148_: _dafny.Array
                d_592_v148_ = d_589_v147_
                d_593_v149_: _dafny.Array
                nw114_ = _dafny.Array(None, 1)
                nw114_[int(0)] = d_592_v148_
                d_593_v149_ = nw114_
                index102_ = default__.safeIndex(985, (d_588_v146_).length(0))
                (d_588_v146_)[index102_] = d_593_v149_
                index103_ = default__.safeIndex(985, (d_588_v146_).length(0))
                (d_588_v146_)[index103_] = d_593_v149_
                d_594_v150_: _dafny.Array
                def lambda40_(d_595_v51_):
                    def lambda41_(d_596_i11_):
                        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_595_v51_.f13, d_595_v51_.f13, d_595_v51_.f13, (d_595_v51_).f12, True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_595_v51_).f12]) for d_597_i12_ in range(default__.abs(92))]))

                    return lambda41_

                init21_ = lambda40_(d_476_v51_)
                nw115_ = _dafny.Array(None, 25)
                for i0_21_ in range(nw115_.length(0)):
                    nw115_[i0_21_] = init21_(i0_21_)
                d_594_v150_ = nw115_
                index104_ = default__.safeIndex(783, (d_594_v150_).length(0))
                (d_594_v150_)[index104_] = default__.fm41(d_476_v51_.f13, globalState)
                d_598_v151_: _dafny.Seq
                d_598_v151_ = _dafny.SeqWithoutIsStrInference([(d_476_v51_).f12, True, d_476_v51_.f13])
                d_599_v152_: _dafny.Seq
                d_599_v152_ = _dafny.SeqWithoutIsStrInference([d_598_v151_, d_598_v151_, (d_598_v151_).set(default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_598_v151_)), True), (d_598_v151_) + (((d_598_v151_).set(default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_598_v151_)), (d_476_v51_).f12)).set(default__.safeIndex(d_479_v54_, len((d_598_v151_).set(default__.safeIndex((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], len(d_598_v151_)), (d_476_v51_).f12))), d_476_v51_.f13)), (d_598_v151_) + (d_598_v151_)])
                index105_ = default__.safeIndex(783, (d_594_v150_).length(0))
                (d_594_v150_)[index105_] = d_599_v152_
                d_600_v153_: _dafny.Array
                def lambda42_(d_601_i13_):
                    return True

                init22_ = lambda42_
                nw116_ = _dafny.Array(None, 11)
                for i0_22_ in range(nw116_.length(0)):
                    nw116_[i0_22_] = init22_(i0_22_)
                d_600_v153_ = nw116_
                index106_ = default__.safeIndex(691, (d_600_v153_).length(0))
                (d_600_v153_)[index106_] = (d_476_v51_).f12
                index107_ = default__.safeIndex(691, (d_600_v153_).length(0))
                (d_600_v153_)[index107_] = self.f7
                index108_ = default__.safeIndex(138, (d_477_v52_).length(0))
                (d_477_v52_)[index108_] = (d_484_v57_)[default__.safeIndex(default__.safeModuloInt((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], d_479_v54_), len(d_484_v57_))]
                d_479_v54_ = (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]
            d_602_v154_: _dafny.Map
            d_602_v154_ = _dafny.Map({d_477_v52_: d_479_v54_})
            d_603_v155_: _dafny.Seq
            d_603_v155_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_602_v154_ = (d_602_v154_ if self.f7 else _dafny.Map({d_477_v52_: len((d_603_v155_).set(default__.safeIndex(d_479_v54_, len(d_603_v155_)), d_476_v51_.f13))}))
            if self.f7:
                d_554_v116_ = d_554_v116_
                d_604_v156_: D1
                d_604_v156_ = D1_DC4(self.f7, self.f7)
                pat_let_tv13_ = d_476_v51_
                d_605_v157_: _dafny.Map
                def iife72_(_pat_let25_0):
                    def iife73_(d_606_dt__update__tmp_h3_):
                        def iife74_(_pat_let26_0):
                            def iife75_(d_607_dt__update_hcf7_h0_):
                                return D1_DC4(d_607_dt__update_hcf7_h0_, (d_606_dt__update__tmp_h3_).cf8)
                            return iife75_(_pat_let26_0)
                        return iife74_(pat_let_tv13_.f13)
                    return iife73_(_pat_let25_0)
                d_605_v157_ = _dafny.Map({iife72_(d_604_v156_): self.f7})
                d_608_v159_: _dafny.Set
                d_608_v159_ = _dafny.Set({d_604_v156_})
                def iife76_():
                    coll22_ = _dafny.Map()
                    compr_22_: D1
                    for compr_22_ in (d_608_v159_).Elements:
                        d_609_v158_: D1 = compr_22_
                        if (d_609_v158_) in (d_608_v159_):
                            coll22_[d_609_v158_] = self.f7
                    return _dafny.Map(coll22_)
                rhs94_ = iife76_()

                rhs95_ = self.f7
                lhs71_ = d_476_v51_
                d_605_v157_ = rhs94_
                lhs71_.f13 = rhs95_
                d_556_v118_ = d_556_v118_
                d_610_v160_: _dafny.Set
                d_610_v160_ = _dafny.Set({(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]})
                d_611_v161_: _dafny.MultiSet
                d_611_v161_ = _dafny.MultiSet([d_477_v52_, d_477_v52_, d_477_v52_, d_477_v52_, d_477_v52_])
                rhs96_ = d_479_v54_
                rhs97_ = (d_610_v160_) - (d_610_v160_)
                rhs98_ = d_611_v161_
                d_479_v54_ = rhs96_
                d_610_v160_ = rhs97_
                d_611_v161_ = rhs98_
                d_612_v162_: D1
                d_612_v162_ = D1_DC5(d_476_v51_.f13, _dafny.SeqWithoutIsStrInference([d_479_v54_ for d_613_i14_ in range(default__.abs(409))]), 394)
                index109_ = default__.safeIndex(138, (d_477_v52_).length(0))
                rhs99_ = ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) >= ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                rhs100_ = d_476_v51_.f13
                rhs101_ = default__.fm0(globalState)
                rhs102_ = (d_612_v162_).cf9
                lhs72_ = self
                lhs73_ = d_476_v51_
                lhs74_ = d_477_v52_
                lhs75_ = default__.safeIndex(138, (d_477_v52_).length(0))
                lhs76_ = d_476_v51_
                lhs72_.f7 = rhs99_
                lhs73_.f13 = rhs100_
                lhs74_[lhs75_] = rhs101_
                lhs76_.f13 = rhs102_
            elif True:
                (d_476_v51_).f13 = d_476_v51_.f13
                d_614_v163_: _dafny.Map
                d_614_v163_ = _dafny.Map({d_476_v51_: d_477_v52_})
                d_615_v164_: _dafny.MultiSet
                d_615_v164_ = _dafny.MultiSet([d_477_v52_, d_477_v52_, ((d_614_v163_)[d_476_v51_] if (d_476_v51_) in (d_614_v163_) else d_477_v52_)])
                d_615_v164_ = d_615_v164_
                d_616_v165_: _dafny.Map
                d_616_v165_ = _dafny.Map({(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]: (d_476_v51_).fm17(d_479_v54_, d_476_v51_.f13, globalState)})
                d_616_v165_ = (d_616_v165_).set(d_479_v54_, d_603_v155_)
                d_617_v166_: _dafny.Map
                d_617_v166_ = _dafny.Map({(d_476_v51_).f12: len(_dafny.Map({d_476_v51_.f13: (d_476_v51_).f12}))})
                d_617_v166_ = (d_617_v166_).set((d_476_v51_).f12, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                d_618_v167_: _dafny.Array
                nw117_ = _dafny.Array(None, 18)
                nw117_[int(0)] = self.f7
                nw117_[int(1)] = (d_476_v51_).f12
                nw117_[int(2)] = (self).fm32(d_476_v51_.f13, _dafny.Set({(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]}), (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], globalState)
                nw117_[int(3)] = (d_476_v51_).f12
                nw117_[int(4)] = (d_476_v51_).f12
                nw117_[int(5)] = (d_476_v51_).fm16(True, globalState)
                nw117_[int(6)] = self.f7
                nw117_[int(7)] = default__.fm1(d_555_v117_, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], d_479_v54_, globalState)
                nw117_[int(8)] = self.f7
                nw117_[int(9)] = ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) != ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                nw117_[int(10)] = (d_476_v51_).f12
                nw117_[int(11)] = (d_476_v51_).f12
                nw117_[int(12)] = not (not(d_476_v51_.f13)) or (self.f7)
                nw117_[int(13)] = (d_476_v51_.f13) == ((d_476_v51_).f12)
                nw117_[int(14)] = d_476_v51_.f13
                nw117_[int(15)] = d_476_v51_.f13
                nw117_[int(16)] = not(((d_476_v51_).f12 if self.f7 else (d_476_v51_).f12))
                nw117_[int(17)] = not(self.f7)
                d_618_v167_ = nw117_
                rhs103_ = (d_554_v116_) <= (_dafny.SeqWithoutIsStrInference([d_556_v118_ for d_619_i15_ in range(default__.abs(356))]))
                rhs104_ = d_618_v167_
                lhs77_ = self
                lhs77_.f7 = rhs103_
                d_618_v167_ = rhs104_
            d_620_v168_: _dafny.Map
            d_620_v168_ = _dafny.Map({self.f7: (d_476_v51_.f13 if (d_476_v51_).f12 else d_476_v51_.f13)})
            (self).f7 = ((d_620_v168_)[self.f7] if (self.f7) in (d_620_v168_) else d_476_v51_.f13)
        elif True:
            if False:
                d_479_v54_ = d_479_v54_
                d_621_v169_: C0
                nw118_ = C0()
                nw118_.ctor__(self.f7, d_476_v51_.f13)
                d_621_v169_ = nw118_
                (d_476_v51_).f13 = d_476_v51_.f13
                d_622_v170_: _dafny.Map
                d_622_v170_ = _dafny.Map({d_476_v51_.f13: _dafny.CodePoint('e')})
                d_623_v171_: _dafny.Map
                d_623_v171_ = _dafny.Map({((d_476_v51_).f12 if True else True): len((_dafny.Map({self.f7: d_556_v118_})) | (d_622_v170_))})
                d_623_v171_ = (d_623_v171_).set((d_476_v51_).f12, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                d_624_v172_: _dafny.MultiSet
                d_624_v172_ = _dafny.MultiSet([d_621_v169_.f13, False])
                default__.m0(((d_624_v172_)[(d_476_v51_).f12] if ((d_476_v51_).f12) in (d_624_v172_) else (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]), d_479_v54_, globalState)
            elif True:
                d_479_v54_ = default__.safeDivisionInt(d_479_v54_, d_479_v54_)
                index110_ = default__.safeIndex(138, (d_477_v52_).length(0))
                rhs105_ = ((d_479_v54_) + (d_479_v54_)) - ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                rhs106_ = d_554_v116_
                lhs78_ = d_477_v52_
                lhs79_ = default__.safeIndex(138, (d_477_v52_).length(0))
                lhs78_[lhs79_] = rhs105_
                d_554_v116_ = rhs106_
                d_625_v173_: _dafny.Array
                nw119_ = _dafny.Array(None, 1)
                nw119_[int(0)] = (d_476_v51_).f12
                d_625_v173_ = nw119_
                d_626_v174_: _dafny.Set
                d_626_v174_ = _dafny.Set({d_625_v173_, d_625_v173_})
                d_627_v175_: D8
                d_627_v175_ = D8_DC21(len(d_554_v116_), d_626_v174_, (0) - ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]))
                default__.m0((d_479_v54_ if self.f7 else (d_627_v175_).cf38), (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], globalState)
                d_479_v54_ = (d_479_v54_) * (d_479_v54_)
                d_628_v176_: _dafny.MultiSet
                d_628_v176_ = _dafny.MultiSet([d_479_v54_])
                (d_476_v51_).f13 = not(((d_628_v176_) - (_dafny.MultiSet([(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))], d_479_v54_]))).ispropersubset(d_628_v176_))
            d_629_v177_: _dafny.Map
            d_629_v177_ = _dafny.Map({(d_476_v51_).f12: (0) - (default__.safeModuloInt(163, d_479_v54_))})
            d_629_v177_ = (d_629_v177_).set(d_476_v51_.f13, ((d_629_v177_)[(d_476_v51_).fm16((d_476_v51_).f12, globalState)] if ((d_476_v51_).fm16((d_476_v51_).f12, globalState)) in (d_629_v177_) else d_479_v54_))
            d_630_v178_: _dafny.Set
            d_630_v178_ = _dafny.Set({d_477_v52_})
            d_631_v179_: D0
            d_631_v179_ = D0_DC2(d_476_v51_.f13, 856, d_630_v178_, d_479_v54_)
            d_632_v180_: _dafny.Map
            d_632_v180_ = _dafny.Map({(d_476_v51_).fm16(d_476_v51_.f13, globalState): d_631_v179_})
            default__.m0((0) - (len(d_632_v180_)), d_479_v54_, globalState)
            rhs107_ = (d_476_v51_).f12
            rhs108_ = (d_484_v57_) + (d_484_v57_)
            lhs80_ = d_476_v51_
            lhs80_.f13 = rhs107_
            d_484_v57_ = rhs108_
            if (d_476_v51_).f12:
                d_633_v181_: _dafny.Set
                d_633_v181_ = _dafny.Set({d_476_v51_.f13})
                d_634_v182_: _dafny.Set
                d_634_v182_ = _dafny.Set({len(d_633_v181_)})
                d_635_v183_: _dafny.Array
                nw120_ = _dafny.Array(None, 11)
                nw120_[int(0)] = d_476_v51_.f13
                nw120_[int(1)] = d_476_v51_.f13
                nw120_[int(2)] = d_476_v51_.f13
                nw120_[int(3)] = (d_476_v51_).f12
                nw120_[int(4)] = True
                nw120_[int(5)] = d_476_v51_.f13
                nw120_[int(6)] = (self).fm32(d_476_v51_.f13, d_634_v182_, 893, (0) - ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]), globalState)
                nw120_[int(7)] = d_476_v51_.f13
                nw120_[int(8)] = d_476_v51_.f13
                nw120_[int(9)] = False
                nw120_[int(10)] = (d_476_v51_).f12
                d_635_v183_ = nw120_
                d_636_v184_: _dafny.Set
                d_636_v184_ = _dafny.Set({d_635_v183_})
                (d_476_v51_).f13 = (d_635_v183_) in (d_636_v184_)
                d_637_v185_: _dafny.Map
                d_637_v185_ = _dafny.Map({d_635_v183_: (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]})
                d_637_v185_ = (d_637_v185_).set(d_635_v183_, (d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                d_638_v186_: D1
                d_638_v186_ = D1_DC4(d_476_v51_.f13, (d_476_v51_).f12)
                index111_ = default__.safeIndex(312, (d_635_v183_).length(0))
                (d_635_v183_)[index111_] = (len(d_633_v181_)) >= (d_479_v54_)
                index112_ = default__.safeIndex(138, (d_477_v52_).length(0))
                index113_ = default__.safeIndex(312, (d_635_v183_).length(0))
                rhs109_ = 301
                rhs110_ = d_638_v186_
                rhs111_ = ((d_476_v51_).f12) and (self.f7)
                rhs112_ = (True if self.f7 else (d_476_v51_).f12)
                lhs81_ = d_477_v52_
                lhs82_ = default__.safeIndex(138, (d_477_v52_).length(0))
                lhs83_ = d_635_v183_
                lhs84_ = default__.safeIndex(312, (d_635_v183_).length(0))
                lhs85_ = d_476_v51_
                lhs81_[lhs82_] = rhs109_
                d_638_v186_ = rhs110_
                lhs83_[lhs84_] = rhs111_
                lhs85_.f13 = rhs112_
                d_639_v187_: _dafny.Map
                d_639_v187_ = _dafny.Map({(d_476_v51_).f12: d_484_v57_})
                d_640_v188_: _dafny.Seq
                d_640_v188_ = _dafny.SeqWithoutIsStrInference([True])
                d_639_v187_ = (d_639_v187_).set(not ((d_640_v188_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))] for d_641_i16_ in range(default__.abs(254))])), len(d_640_v188_))]) or (d_476_v51_.f13), d_484_v57_)
                (self).f7 = (((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))]) - ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])) == ((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
            elif True:
                d_642_v189_: _dafny.Set
                d_642_v189_ = _dafny.Set({d_479_v54_})
                d_643_v190_: _dafny.Set
                d_643_v190_ = _dafny.Set({d_476_v51_.f13})
                (self).f7 = not((not(not((d_642_v189_).issubset(d_642_v189_)))) not in (d_643_v190_))
                d_644_v191_: D0
                d_644_v191_ = D0_DC1((d_477_v52_)[default__.safeIndex(138, (d_477_v52_).length(0))])
                d_645_v192_: _dafny.Array
                def lambda43_(d_646_v51_):
                    def lambda44_(d_647_i17_):
                        return (d_646_v51_).f12

                    return lambda44_

                init23_ = lambda43_(d_476_v51_)
                nw121_ = _dafny.Array(None, 22)
                for i0_23_ in range(nw121_.length(0)):
                    nw121_[i0_23_] = init23_(i0_23_)
                d_645_v192_ = nw121_
                d_648_v193_: _dafny.Array
                nw122_ = _dafny.Array(None, 3)
                nw122_[int(0)] = d_645_v192_
                nw122_[int(1)] = d_645_v192_
                nw122_[int(2)] = d_645_v192_
                d_648_v193_ = nw122_
                rhs113_ = d_644_v191_
                rhs114_ = default__.safeDivisionInt(default__.fm0(globalState), d_479_v54_)
                rhs115_ = d_476_v51_.f13
                rhs116_ = _dafny.SeqWithoutIsStrInference([d_556_v118_ for d_649_i18_ in range(default__.abs(147))])
                rhs117_ = d_648_v193_
                lhs86_ = d_476_v51_
                d_644_v191_ = rhs113_
                d_479_v54_ = rhs114_
                lhs86_.f13 = rhs115_
                d_554_v116_ = rhs116_
                d_648_v193_ = rhs117_
                (d_476_v51_).f13 = d_476_v51_.f13
                d_650_v194_: _dafny.Seq
                d_650_v194_ = _dafny.SeqWithoutIsStrInference([d_478_v53_, _dafny.Map({(0) - (d_479_v54_): (d_476_v51_).f12}), (d_478_v53_) | (d_478_v53_), d_478_v53_, d_478_v53_])
                d_650_v194_ = d_650_v194_
                d_651_v195_: C0
                nw123_ = C0()
                nw123_.ctor__(self.f7, self.f7)
                d_651_v195_ = nw123_


class C2(T0):
    def  __init__(self):
        self._f7: bool = False
        self._f15: _dafny.Set = _dafny.Set({})
        self._f16: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def ctor__(self, f15, f16, f7):
        (self)._f15 = f15
        (self)._f16 = f16
        (self).f7 = f7

    def fm24(self, p0, p1, p2, globalState):
        return (_dafny.Map({self.f7: self.f7})) | (_dafny.Map({self.f7: self.f7}))

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Set = _dafny.Set({})
        if (_dafny.SeqWithoutIsStrInference([(self).f16 for d_652_i0_ in range(default__.abs(85))])) not in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nohtpuim")) for d_653_i1_ in range(default__.abs(23))])):
            d_654_v0_: _dafny.Seq
            d_654_v0_ = _dafny.SeqWithoutIsStrInference([p1])
            d_655_v1_: _dafny.Seq
            d_655_v1_ = _dafny.SeqWithoutIsStrInference([p0])
            d_656_v2_: D1
            d_656_v2_ = D1_DC5((p2) == (p0), ((d_654_v0_).set(default__.safeIndex(p1, len(d_654_v0_)), p1)) + (d_654_v0_), len(d_655_v1_))
            source2_ = d_656_v2_
            if source2_.is_DC4:
                d_657___mcc_h0_ = source2_.cf7
                d_658___mcc_h1_ = source2_.cf8
                d_659_cf8_ = d_658___mcc_h1_
                d_660_cf7_ = d_657___mcc_h0_
                d_661_v3_: int
                d_661_v3_ = 528
                d_661_v3_ = default__.fm0(globalState)
                d_661_v3_ = d_661_v3_
                rhs118_ = d_661_v3_
                rhs119_ = default__.safeDivisionInt(p1, d_661_v3_)
                rhs120_ = (p1) < (d_661_v3_)
                d_661_v3_ = rhs118_
                d_661_v3_ = rhs119_
                d_660_cf7_ = rhs120_
                d_661_v3_ = (p1) - (d_661_v3_)
            elif source2_.is_DC5:
                d_662___mcc_h2_ = source2_.cf9
                d_663___mcc_h3_ = source2_.cf10
                d_664___mcc_h4_ = source2_.cf11
                d_665_cf11_ = d_664___mcc_h4_
                d_666_cf10_ = d_663___mcc_h3_
                d_667_cf9_ = d_662___mcc_h2_
                d_668_v4_: _dafny.Array
                def lambda45_(d_669_p1_):
                    def lambda46_(d_670_i2_):
                        return default__.safeDivisionInt(d_670_i2_, d_669_p1_)

                    return lambda46_

                init24_ = lambda45_(p1)
                nw124_ = _dafny.Array(None, 11)
                for i0_24_ in range(nw124_.length(0)):
                    nw124_[i0_24_] = init24_(i0_24_)
                d_668_v4_ = nw124_
                d_668_v4_ = d_668_v4_
                d_671_v5_: _dafny.Array
                def lambda47_(d_672_i3_):
                    return self.f7

                init25_ = lambda47_
                nw125_ = _dafny.Array(None, 2)
                for i0_25_ in range(nw125_.length(0)):
                    nw125_[i0_25_] = init25_(i0_25_)
                d_671_v5_ = nw125_
                d_673_v6_: _dafny.MultiSet
                d_673_v6_ = _dafny.MultiSet([d_671_v5_, d_671_v5_, d_671_v5_, d_671_v5_])
                d_673_v6_ = d_673_v6_
                d_674_v7_: _dafny.Set
                d_674_v7_ = _dafny.Set({p0})
                d_675_v8_: _dafny.Map
                d_675_v8_ = _dafny.Map({p0: len(d_674_v7_)})
                d_676_v9_: _dafny.MultiSet
                d_676_v9_ = _dafny.MultiSet([len(d_675_v8_)])
                d_677_v10_: _dafny.Seq
                d_677_v10_ = _dafny.SeqWithoutIsStrInference([d_676_v9_, d_676_v9_])
                d_678_v11_: D1
                d_678_v11_ = D1_DC3(d_674_v7_)
                d_679_v12_: D0
                d_679_v12_ = D0_DC0(d_665_cf11_)
                d_680_v13_: _dafny.Seq
                d_680_v13_ = _dafny.SeqWithoutIsStrInference([d_679_v12_, d_679_v12_])
                d_681_v14_: _dafny.Array
                nw126_ = _dafny.Array(None, 13)
                nw126_[int(0)] = _dafny.CodePoint('a')
                nw126_[int(1)] = (self).f16
                nw126_[int(2)] = _dafny.CodePoint('c')
                nw126_[int(3)] = (self).f16
                nw126_[int(4)] = default__.fm26(d_667_cf9_, p1, d_665_cf11_, globalState)
                nw126_[int(5)] = (self).f16
                nw126_[int(6)] = (self).f16
                nw126_[int(7)] = (self).f16
                nw126_[int(8)] = (self).f16
                nw126_[int(9)] = (self).f16
                nw126_[int(10)] = (self).f16
                nw126_[int(11)] = _dafny.CodePoint('p')
                nw126_[int(12)] = _dafny.CodePoint('y')
                d_681_v14_ = nw126_
                (self).m13((_dafny.MultiSet(d_677_v10_)).isdisjoint(_dafny.MultiSet([d_676_v9_, default__.fm25(d_665_cf11_, (self).f16, len(_dafny.SeqWithoutIsStrInference([p2, p0])), d_678_v11_, globalState)])), p0, default__.fm1(d_680_v13_, d_665_cf11_, d_665_cf11_, globalState), d_681_v14_, globalState)
                d_682_v15_: _dafny.Seq
                d_682_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wawknevf"))
                d_682_v15_ = d_682_v15_
            elif source2_.is_DC3:
                d_683___mcc_h5_ = source2_.cf6
                d_684_cf6_ = d_683___mcc_h5_
                d_685_v16_: _dafny.Map
                d_685_v16_ = _dafny.Map({p1: default__.safeModuloInt(p1, p1)})
                d_685_v16_ = (d_685_v16_).set(p1, default__.fm0(globalState))
                d_686_v17_: _dafny.Array
                nw127_ = _dafny.Array(int(0), 1)
                d_686_v17_ = nw127_
                index114_ = default__.safeIndex(452, (d_686_v17_).length(0))
                (d_686_v17_)[index114_] = (0) - (p1)
                index115_ = default__.safeIndex(452, (d_686_v17_).length(0))
                (d_686_v17_)[index115_] = default__.fm0(globalState)
                (self).f7 = p0
                index116_ = default__.safeIndex(452, (d_686_v17_).length(0))
                (d_686_v17_)[index116_] = (p1) + ((0) - ((d_686_v17_)[default__.safeIndex(452, (d_686_v17_).length(0))]))
            elif True:
                d_687___mcc_h6_ = source2_.cf12
                d_688_cf12_ = d_687___mcc_h6_
                d_689_v18_: _dafny.Map
                d_689_v18_ = _dafny.Map({p0: not(p2)})
                r0 = ((self).fm24(False, p0, p1, globalState)) | (d_689_v18_)
                d_690_v19_: int
                d_690_v19_ = 410
                d_691_v20_: _dafny.Array
                nw128_ = _dafny.Array(False, 28)
                d_691_v20_ = nw128_
                d_692_v21_: _dafny.Seq
                d_692_v21_ = _dafny.SeqWithoutIsStrInference([D0_DC0(-726)])
                index117_ = default__.safeIndex(733, (d_691_v20_).length(0))
                (d_691_v20_)[index117_] = not(default__.fm1(d_692_v21_, d_690_v19_, 891, globalState))
                d_693_v22_: C0
                nw129_ = C0()
                nw129_.ctor__(self.f7, not(p0))
                d_693_v22_ = nw129_
                index118_ = default__.safeIndex(733, (d_691_v20_).length(0))
                rhs121_ = (d_690_v19_) * (d_690_v19_)
                rhs122_ = p2
                rhs123_ = d_693_v22_
                rhs124_ = p1
                lhs87_ = d_691_v20_
                lhs88_ = default__.safeIndex(733, (d_691_v20_).length(0))
                d_690_v19_ = rhs121_
                lhs87_[lhs88_] = rhs122_
                d_693_v22_ = rhs123_
                d_690_v19_ = rhs124_
                d_694_v23_: _dafny.Map
                d_694_v23_ = _dafny.Map({p1: d_654_v0_})
                d_694_v23_ = (d_694_v23_).set(p1, d_654_v0_)
                d_695_v24_: _dafny.Array
                def lambda48_(d_696_i4_):
                    return (_dafny.SeqWithoutIsStrInference([(self).f16 for d_697_i5_ in range(default__.abs(761))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwiju")))

                init26_ = lambda48_
                nw130_ = _dafny.Array(None, 26)
                for i0_26_ in range(nw130_.length(0)):
                    nw130_[i0_26_] = init26_(i0_26_)
                d_695_v24_ = nw130_
                d_698_v25_: _dafny.Seq
                d_698_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnjpk"))
                index119_ = default__.safeIndex(808, (d_695_v24_).length(0))
                (d_695_v24_)[index119_] = d_698_v25_
                d_699_v26_: _dafny.Array
                nw131_ = _dafny.Array(_dafny.Map({}), 4)
                d_699_v26_ = nw131_
                d_700_v27_: _dafny.Map
                d_700_v27_ = _dafny.Map({(d_693_v22_).f12: d_690_v19_})
                index120_ = default__.safeIndex(673, (d_699_v26_).length(0))
                (d_699_v26_)[index120_] = d_700_v27_
                d_701_v28_: str
                d_701_v28_ = _dafny.CodePoint('d')
                d_702_v29_: _dafny.MultiSet
                d_702_v29_ = _dafny.MultiSet([d_690_v19_])
                d_703_v30_: _dafny.MultiSet
                d_703_v30_ = _dafny.MultiSet([_dafny.MultiSet([80, p1]), d_702_v29_, d_702_v29_, d_702_v29_])
                d_704_v31_: _dafny.Set
                d_704_v31_ = _dafny.Set({self.f7, ((d_689_v18_)[p2] if (p2) in (d_689_v18_) else (d_691_v20_)[default__.safeIndex(733, (d_691_v20_).length(0))]), (d_698_v25_) == (d_698_v25_), (d_703_v30_).isdisjoint(d_703_v30_), (d_691_v20_)[default__.safeIndex(733, (d_691_v20_).length(0))]})
                index121_ = default__.safeIndex(808, (d_695_v24_).length(0))
                index122_ = default__.safeIndex(673, (d_699_v26_).length(0))
                rhs125_ = d_698_v25_
                rhs126_ = (d_700_v27_) | (d_700_v27_)
                rhs127_ = d_704_v31_
                rhs128_ = default__.fm26(not(self.f7), default__.safeDivisionInt(p1, (d_654_v0_)[default__.safeIndex(p1, len(d_654_v0_))]), 236, globalState)
                lhs89_ = d_695_v24_
                lhs90_ = default__.safeIndex(808, (d_695_v24_).length(0))
                lhs91_ = d_699_v26_
                lhs92_ = default__.safeIndex(673, (d_699_v26_).length(0))
                lhs89_[lhs90_] = rhs125_
                lhs91_[lhs92_] = rhs126_
                r1 = rhs127_
                d_701_v28_ = rhs128_
            d_705_v32_: _dafny.Set
            d_705_v32_ = _dafny.Set({p2})
            d_706_v33_: _dafny.Array
            def lambda49_(d_707_p1_):
                def lambda50_(d_708_i6_):
                    return default__.safeDivisionInt(d_708_i6_, d_707_p1_)

                return lambda50_

            init27_ = lambda49_(p1)
            nw132_ = _dafny.Array(None, 12)
            for i0_27_ in range(nw132_.length(0)):
                nw132_[i0_27_] = init27_(i0_27_)
            d_706_v33_ = nw132_
            d_709_v34_: _dafny.Map
            d_709_v34_ = _dafny.Map({d_705_v32_: d_706_v33_})
            d_710_v35_: _dafny.Array
            def lambda51_(d_711_p0_, d_712_p2_):
                def lambda52_(d_713_i7_):
                    return (D1_DC4(d_711_p0_, d_712_p2_)).cf8

                return lambda52_

            init28_ = lambda51_(p0, p2)
            nw133_ = _dafny.Array(None, 20)
            for i0_28_ in range(nw133_.length(0)):
                nw133_[i0_28_] = init28_(i0_28_)
            d_710_v35_ = nw133_
            d_714_v36_: _dafny.Map
            d_714_v36_ = _dafny.Map({d_709_v34_: d_710_v35_})
            d_714_v36_ = (d_714_v36_).set((d_709_v34_) | (d_709_v34_), d_710_v35_)
            d_715_v37_: int
            d_715_v37_ = 584
            d_716_v38_: _dafny.Map
            d_716_v38_ = _dafny.Map({not((p0) or (True)): default__.safeDivisionInt(len(d_654_v0_), 428)})
            d_715_v37_ = ((d_716_v38_)[p0] if (p0) in (d_716_v38_) else p1)
            index123_ = default__.safeIndex(632, (d_710_v35_).length(0))
            (d_710_v35_)[index123_] = p2
            index124_ = default__.safeIndex(632, (d_710_v35_).length(0))
            (d_710_v35_)[index124_] = p2
            default__.m0(default__.safeModuloInt(p1, d_715_v37_), d_715_v37_, globalState)
        elif True:
            d_717_v39_: int
            d_717_v39_ = -291
            d_718_v40_: _dafny.Seq
            d_718_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uk"))
            d_717_v39_ = (0) - (len((_dafny.Map({p0: d_717_v39_})) | (_dafny.Map({not(self.f7): len(d_718_v40_)}))))
            d_719_v41_: D0
            d_719_v41_ = D0_DC0(p1)
            d_720_v42_: _dafny.Seq
            def iife77_(_pat_let27_0):
                def iife78_(d_721_dt__update__tmp_h0_):
                    def iife79_(_pat_let28_0):
                        def iife80_(d_722_dt__update_hcf0_h0_):
                            return D0_DC0(d_722_dt__update_hcf0_h0_)
                        return iife80_(_pat_let28_0)
                    return iife79_(778)
                return iife78_(_pat_let27_0)
            d_720_v42_ = _dafny.SeqWithoutIsStrInference([iife77_(D0_DC0(p1))])
            d_723_v43_: D6
            d_723_v43_ = D6_DC17(default__.fm1(d_720_v42_, 472, p1, globalState), d_717_v39_, (self).f16, _dafny.CodePoint('u'))
            pat_let_tv14_ = d_717_v39_
            d_724_v44_: _dafny.Array
            nw134_ = _dafny.Array(None, 19)
            nw134_[int(0)] = self.f7
            nw134_[int(1)] = p2
            nw134_[int(2)] = default__.fm1(_dafny.SeqWithoutIsStrInference([d_719_v41_]), (d_723_v43_).cf32, p1, globalState)
            nw134_[int(3)] = not(p2)
            nw134_[int(4)] = p0
            def iife81_(_pat_let29_0):
                def iife82_(d_725_dt__update__tmp_h1_):
                    def iife83_(_pat_let30_0):
                        def iife84_(d_726_dt__update_hcf0_h1_):
                            return D0_DC0(d_726_dt__update_hcf0_h1_)
                        return iife84_(_pat_let30_0)
                    return iife83_(pat_let_tv14_)
                return iife82_(_pat_let29_0)
            nw134_[int(5)] = default__.fm1(_dafny.SeqWithoutIsStrInference([d_719_v41_, iife81_(d_719_v41_), d_719_v41_, d_719_v41_]), d_717_v39_, d_717_v39_, globalState)
            nw134_[int(6)] = self.f7
            nw134_[int(7)] = self.f7
            nw134_[int(8)] = self.f7
            nw134_[int(9)] = p2
            nw134_[int(10)] = self.f7
            nw134_[int(11)] = p0
            nw134_[int(12)] = p2
            nw134_[int(13)] = p0
            nw134_[int(14)] = p0
            nw134_[int(15)] = self.f7
            nw134_[int(16)] = self.f7
            nw134_[int(17)] = p0
            nw134_[int(18)] = True
            d_724_v44_ = nw134_
            d_727_v45_: D8
            d_727_v45_ = D8_DC21(d_717_v39_, (self).f15, d_717_v39_)
            if ((d_727_v45_).cf39).ispropersubset(_dafny.Set({d_724_v44_})):
                d_728_v46_: _dafny.Map
                d_728_v46_ = _dafny.Map({default__.fm0(globalState): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wl")))})
                d_729_v47_: _dafny.MultiSet
                d_729_v47_ = _dafny.MultiSet([p1, p1])
                d_717_v39_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xngciqxxs")))) - (((d_728_v46_)[d_717_v39_] if (d_717_v39_) in (d_728_v46_) else (0) - (((d_729_v47_)[p1] if (p1) in (d_729_v47_) else p1))))
                d_730_v48_: D0
                d_730_v48_ = D0_DC1((p1) - (d_717_v39_))
                d_730_v48_ = d_730_v48_
                d_731_v49_: _dafny.Array
                nw135_ = _dafny.Array(None, 14)
                nw135_[int(0)] = -717
                nw135_[int(1)] = d_717_v39_
                nw135_[int(2)] = d_717_v39_
                nw135_[int(3)] = d_717_v39_
                nw135_[int(4)] = default__.fm0(globalState)
                nw135_[int(5)] = p1
                nw135_[int(6)] = d_717_v39_
                nw135_[int(7)] = p1
                nw135_[int(8)] = p1
                nw135_[int(9)] = d_717_v39_
                nw135_[int(10)] = d_717_v39_
                nw135_[int(11)] = p1
                nw135_[int(12)] = p1
                nw135_[int(13)] = p1
                d_731_v49_ = nw135_
                d_732_v50_: _dafny.Map
                d_732_v50_ = _dafny.Map({d_731_v49_: self.f7})
                d_732_v50_ = (d_732_v50_).set(d_731_v49_, p0)
                d_733_v51_: _dafny.Map
                d_733_v51_ = _dafny.Map({(p0 if self.f7 else self.f7): p0})
                r0 = d_733_v51_
                (self).f7 = ((default__.fm0(globalState)) * (p1)) == ((0) - (p1))
            elif True:
                d_734_v52_: _dafny.Array
                nw136_ = _dafny.Array(None, 20)
                nw136_[int(0)] = (self).f16
                nw136_[int(1)] = (self).f16
                nw136_[int(2)] = (self).f16
                nw136_[int(3)] = (self).f16
                nw136_[int(4)] = (self).f16
                nw136_[int(5)] = (self).f16
                nw136_[int(6)] = (self).f16
                nw136_[int(7)] = (self).f16
                nw136_[int(8)] = (self).f16
                nw136_[int(9)] = _dafny.CodePoint('p')
                nw136_[int(10)] = (self).f16
                nw136_[int(11)] = _dafny.CodePoint('h')
                nw136_[int(12)] = (self).f16
                nw136_[int(13)] = (self).f16
                nw136_[int(14)] = (self).f16
                nw136_[int(15)] = (self).f16
                nw136_[int(16)] = (self).f16
                nw136_[int(17)] = (self).f16
                nw136_[int(18)] = (self).f16
                nw136_[int(19)] = _dafny.CodePoint('w')
                d_734_v52_ = nw136_
                (self).m13(p2, p0, not (p0) or (self.f7), d_734_v52_, globalState)
                d_717_v39_ = (d_717_v39_) + (d_717_v39_)
                d_735_v53_: _dafny.MultiSet
                d_735_v53_ = _dafny.MultiSet([not(p0)])
                d_736_v54_: _dafny.Map
                d_736_v54_ = _dafny.Map({self.f7: p1})
                d_737_v55_: _dafny.Map
                d_737_v55_ = _dafny.Map({_dafny.CodePoint('h'): self.f7})
                d_738_v56_: _dafny.Seq
                d_738_v56_ = _dafny.SeqWithoutIsStrInference([((d_737_v55_)[_dafny.CodePoint('y')] if (_dafny.CodePoint('y')) in (d_737_v55_) else p0)])
                d_739_v57_: _dafny.Seq
                d_739_v57_ = _dafny.SeqWithoutIsStrInference([d_717_v39_])
                d_740_v58_: _dafny.Map
                d_740_v58_ = _dafny.Map({p1: False})
                d_741_v59_: _dafny.Array
                nw137_ = _dafny.Array(None, 27)
                nw137_[int(0)] = d_717_v39_
                nw137_[int(1)] = ((d_735_v53_)[p0] if (p0) in (d_735_v53_) else p1)
                nw137_[int(2)] = p1
                nw137_[int(3)] = p1
                nw137_[int(4)] = ((d_736_v54_)[p0] if (p0) in (d_736_v54_) else d_717_v39_)
                nw137_[int(5)] = len(d_738_v56_)
                nw137_[int(6)] = p1
                nw137_[int(7)] = p1
                nw137_[int(8)] = d_717_v39_
                nw137_[int(9)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_742_i8_ in range(default__.abs(467))])).set(default__.safeIndex(104, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_743_i8_ in range(default__.abs(467))]))), (self).f16))
                nw137_[int(10)] = p1
                nw137_[int(11)] = (d_739_v57_)[default__.safeIndex(p1, len(d_739_v57_))]
                nw137_[int(12)] = p1
                nw137_[int(13)] = d_717_v39_
                nw137_[int(14)] = d_717_v39_
                nw137_[int(15)] = p1
                nw137_[int(16)] = d_717_v39_
                nw137_[int(17)] = default__.fm0(globalState)
                nw137_[int(18)] = -396
                nw137_[int(19)] = (d_739_v57_)[default__.safeIndex(d_717_v39_, len(d_739_v57_))]
                nw137_[int(20)] = -866
                nw137_[int(21)] = len(_dafny.SeqWithoutIsStrInference([(self).f16 for d_744_i9_ in range(default__.abs(72))]))
                nw137_[int(22)] = d_717_v39_
                nw137_[int(23)] = p1
                nw137_[int(24)] = (0) - (p1)
                nw137_[int(25)] = len(d_740_v58_)
                nw137_[int(26)] = p1
                d_741_v59_ = nw137_
                d_745_v60_: _dafny.Seq
                d_745_v60_ = _dafny.SeqWithoutIsStrInference([d_741_v59_])
                d_746_v61_: _dafny.Seq
                d_746_v61_ = _dafny.SeqWithoutIsStrInference([d_745_v60_, d_745_v60_, d_745_v60_, (d_745_v60_) + (d_745_v60_), d_745_v60_])
                d_717_v39_ = len(((d_746_v61_)[default__.safeIndex((0) - (d_717_v39_), len(d_746_v61_))]).set(default__.safeIndex(len(default__.fm27(self.f7, _dafny.CodePoint('c'), self.f7, (d_723_v43_).cf31, globalState)), len((d_746_v61_)[default__.safeIndex((0) - (d_717_v39_), len(d_746_v61_))])), d_741_v59_))
                (self).f7 = (p1) < (p1)
                index125_ = default__.safeIndex(377, (d_741_v59_).length(0))
                (d_741_v59_)[index125_] = len(d_718_v40_)
                index126_ = default__.safeIndex(377, (d_741_v59_).length(0))
                (d_741_v59_)[index126_] = p1
            (self).f7 = p2
            (self).f7 = ((d_718_v40_) + (d_718_v40_)) == ((d_718_v40_).set(default__.safeIndex((0) - (p1), len(d_718_v40_)), (self).f16))
            d_747_v62_: _dafny.Array
            def lambda53_(d_748_v40_):
                def lambda54_(d_749_i10_):
                    return d_748_v40_

                return lambda54_

            init29_ = lambda53_(d_718_v40_)
            nw138_ = _dafny.Array(None, 3)
            for i0_29_ in range(nw138_.length(0)):
                nw138_[i0_29_] = init29_(i0_29_)
            d_747_v62_ = nw138_
            index127_ = default__.safeIndex(49, (d_747_v62_).length(0))
            (d_747_v62_)[index127_] = d_718_v40_
            d_750_v63_: _dafny.Seq
            d_750_v63_ = _dafny.SeqWithoutIsStrInference([938])
            d_751_v64_: _dafny.MultiSet
            d_751_v64_ = _dafny.MultiSet([len((_dafny.SeqWithoutIsStrInference([(self).f16 for d_752_i11_ in range(default__.abs(662))])).set(default__.safeIndex(72, len(_dafny.SeqWithoutIsStrInference([(self).f16 for d_753_i11_ in range(default__.abs(662))]))), (self).f16)), (0) - ((0) - (len(d_750_v63_))), d_717_v39_])
            d_754_v65_: _dafny.Map
            d_754_v65_ = _dafny.Map({d_717_v39_: _dafny.MultiSet([default__.fm0(globalState), p1])})
            index128_ = default__.safeIndex(49, (d_747_v62_).length(0))
            rhs129_ = not((d_751_v64_).isdisjoint(((d_754_v65_)[p1] if (p1) in (d_754_v65_) else d_751_v64_)))
            rhs130_ = p1
            rhs131_ = self.f7
            rhs132_ = d_718_v40_
            lhs93_ = self
            lhs94_ = self
            lhs95_ = d_747_v62_
            lhs96_ = default__.safeIndex(49, (d_747_v62_).length(0))
            lhs93_.f7 = rhs129_
            d_717_v39_ = rhs130_
            lhs94_.f7 = rhs131_
            lhs95_[lhs96_] = rhs132_
        if (D3_DC11(p1, p1, p1, self.f7)).cf25:
            d_755_v66_: _dafny.Map
            d_755_v66_ = _dafny.Map({self.f7: p1})
            d_756_v67_: _dafny.Seq
            d_756_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_757_v68_: C0
            nw139_ = C0()
            nw139_.ctor__(self.f7, (len(d_755_v66_)) == (len(d_756_v67_)))
            d_757_v68_ = nw139_
            (self).f7 = (len((d_756_v67_) + (d_756_v67_))) <= (p1)
            d_758_v69_: D0
            d_758_v69_ = D0_DC0(p1)
            d_759_v70_: _dafny.Seq
            d_759_v70_ = _dafny.SeqWithoutIsStrInference([d_758_v69_])
            d_760_v71_: _dafny.Map
            d_760_v71_ = _dafny.Map({p1: default__.fm1(d_759_v70_, p1, p1, globalState)})
            (self).f7 = ((d_760_v71_)[default__.safeDivisionInt((0) - (p1), p1)] if (default__.safeDivisionInt((0) - (p1), p1)) in (d_760_v71_) else p2)
            d_761_v72_: int
            d_761_v72_ = 408
            d_761_v72_ = d_761_v72_
            d_762_v73_: _dafny.Map
            d_762_v73_ = _dafny.Map({d_757_v68_.f13: (self).f16})
            d_763_v74_: _dafny.Seq
            d_763_v74_ = _dafny.SeqWithoutIsStrInference([318, d_761_v72_])
            d_764_v75_: _dafny.Array
            nw140_ = _dafny.Array(None, 25)
            nw140_[int(0)] = _dafny.CodePoint('s')
            nw140_[int(1)] = (self).f16
            nw140_[int(2)] = (self).f16
            nw140_[int(3)] = (self).f16
            nw140_[int(4)] = ((d_762_v73_)[not(p2)] if (not(p2)) in (d_762_v73_) else (self).f16)
            nw140_[int(5)] = (self).f16
            nw140_[int(6)] = (d_756_v67_)[default__.safeIndex(len(d_763_v74_), len(d_756_v67_))]
            nw140_[int(7)] = (self).f16
            nw140_[int(8)] = (self).f16
            nw140_[int(9)] = (self).f16
            nw140_[int(10)] = (self).f16
            nw140_[int(11)] = (self).f16
            nw140_[int(12)] = (self).f16
            nw140_[int(13)] = (self).f16
            nw140_[int(14)] = (self).f16
            nw140_[int(15)] = (self).f16
            nw140_[int(16)] = _dafny.CodePoint('d')
            nw140_[int(17)] = (self).f16
            nw140_[int(18)] = (self).f16
            nw140_[int(19)] = _dafny.CodePoint('x')
            nw140_[int(20)] = (self).f16
            nw140_[int(21)] = (self).f16
            nw140_[int(22)] = (self).f16
            nw140_[int(23)] = _dafny.CodePoint('u')
            nw140_[int(24)] = (self).f16
            d_764_v75_ = nw140_
            (self).m13(p2, self.f7, (p0) == ((d_757_v68_).f12), d_764_v75_, globalState)
        elif True:
            if p2:
                d_765_v76_: _dafny.Array
                nw141_ = _dafny.Array(_dafny.Set({}), 6)
                d_765_v76_ = nw141_
                d_766_v77_: _dafny.Set
                d_766_v77_ = _dafny.Set({p0})
                index129_ = default__.safeIndex(795, (d_765_v76_).length(0))
                (d_765_v76_)[index129_] = d_766_v77_
                index130_ = default__.safeIndex(795, (d_765_v76_).length(0))
                (d_765_v76_)[index130_] = default__.fm27(p0, (self).f16, p2, p2, globalState)
                d_767_v78_: _dafny.Array
                nw142_ = _dafny.Array(int(0), 6)
                d_767_v78_ = nw142_
                index131_ = default__.safeIndex(394, (d_767_v78_).length(0))
                (d_767_v78_)[index131_] = p1
                d_768_v79_: _dafny.Seq
                d_768_v79_ = _dafny.SeqWithoutIsStrInference([False])
                index132_ = default__.safeIndex(394, (d_767_v78_).length(0))
                (d_767_v78_)[index132_] = (default__.safeModuloInt(len(d_768_v79_), p1)) + ((p1) * (p1))
                d_769_v80_: _dafny.Map
                d_769_v80_ = _dafny.Map({d_767_v78_: d_767_v78_})
                d_769_v80_ = (d_769_v80_).set(d_767_v78_, d_767_v78_)
                (self).f7 = p2
                d_770_v81_: str
                d_770_v81_ = _dafny.CodePoint('o')
                d_771_v82_: _dafny.Map
                d_771_v82_ = _dafny.Map({self.f7: self.f7})
                d_772_v83_: _dafny.Map
                d_772_v83_ = _dafny.Map({p1: d_771_v82_})
                d_770_v81_ = (default__.fm26(p0, len(d_772_v83_), (d_767_v78_)[default__.safeIndex(394, (d_767_v78_).length(0))], globalState) if p2 else (self).f16)
            elif True:
                d_773_v84_: _dafny.Seq
                d_773_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymragmkba"))
                d_773_v84_ = d_773_v84_
                d_774_v85_: _dafny.Map
                d_774_v85_ = _dafny.Map({p2: p0})
                d_774_v85_ = (d_774_v85_).set(not (p0) or (p2), p2)
                d_775_v86_: int
                d_775_v86_ = 718
                d_776_v87_: _dafny.Array
                def lambda55_(d_777_p0_):
                    def lambda56_(d_778_i12_):
                        return d_777_p0_

                    return lambda56_

                init30_ = lambda55_(p0)
                nw143_ = _dafny.Array(None, 22)
                for i0_30_ in range(nw143_.length(0)):
                    nw143_[i0_30_] = init30_(i0_30_)
                d_776_v87_ = nw143_
                d_779_v88_: _dafny.MultiSet
                d_779_v88_ = _dafny.MultiSet([d_776_v87_])
                d_780_v89_: _dafny.Map
                d_780_v89_ = _dafny.Map({(self).f16: ((_dafny.MultiSet([d_776_v87_, d_776_v87_, d_776_v87_, d_776_v87_, d_776_v87_])) - (d_779_v88_)).cardinality})
                d_781_v90_: _dafny.Seq
                d_781_v90_ = _dafny.SeqWithoutIsStrInference([False, p0])
                d_782_v91_: _dafny.MultiSet
                d_782_v91_ = _dafny.MultiSet([p0, p0])
                rhs133_ = len(d_780_v89_)
                rhs134_ = not ((d_773_v84_) != (d_773_v84_)) or ((d_782_v91_).issubset(_dafny.MultiSet(d_781_v90_)))
                lhs97_ = self
                d_775_v86_ = rhs133_
                lhs97_.f7 = rhs134_
                d_783_v92_: _dafny.Array
                def lambda57_(d_784_i13_):
                    return _dafny.CodePoint('l')

                init31_ = lambda57_
                nw144_ = _dafny.Array(None, 3)
                for i0_31_ in range(nw144_.length(0)):
                    nw144_[i0_31_] = init31_(i0_31_)
                d_783_v92_ = nw144_
                (self).m13(p0, (p0) == (self.f7), not (p2) or (p2), d_783_v92_, globalState)
                d_785_v93_: str
                d_785_v93_ = _dafny.CodePoint('b')
                d_785_v93_ = (self).f16
            if p0:
                (self).f7 = p0
                d_786_v94_: _dafny.Seq
                d_786_v94_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cksvuxoeh"))
                d_787_v95_: _dafny.Map
                d_787_v95_ = _dafny.Map({d_786_v94_: self.f7})
                d_788_v96_: _dafny.Seq
                d_788_v96_ = _dafny.SeqWithoutIsStrInference([self.f7])
                d_789_v97_: _dafny.MultiSet
                d_789_v97_ = _dafny.MultiSet([False, False, (d_788_v96_)[default__.safeIndex(p1, len(d_788_v96_))], self.f7, self.f7])
                d_790_v98_: _dafny.Array
                nw145_ = _dafny.Array(False, 1)
                d_790_v98_ = nw145_
                d_791_v99_: _dafny.Map
                d_791_v99_ = _dafny.Map({p0: d_790_v98_})
                d_792_v100_: _dafny.Map
                d_792_v100_ = _dafny.Map({d_791_v99_: d_788_v96_})
                d_793_v101_: _dafny.Map
                d_793_v101_ = _dafny.Map({p0: (D1_DC4(p2, True)).cf7})
                d_794_v102_: _dafny.Array
                nw146_ = _dafny.Array(None, 22)
                nw146_[int(0)] = _dafny.SeqWithoutIsStrInference([p2])
                nw146_[int(1)] = _dafny.SeqWithoutIsStrInference([((d_787_v95_)[default__.fm28(p2, len(_dafny.SeqWithoutIsStrInference([p0])), globalState)] if (default__.fm28(p2, len(_dafny.SeqWithoutIsStrInference([p0])), globalState)) in (d_787_v95_) else p2), p0, p0, self.f7, self.f7])
                nw146_[int(2)] = (d_788_v96_) + (d_788_v96_)
                nw146_[int(3)] = d_788_v96_
                nw146_[int(4)] = d_788_v96_
                nw146_[int(5)] = d_788_v96_
                nw146_[int(6)] = d_788_v96_
                nw146_[int(7)] = d_788_v96_
                nw146_[int(8)] = d_788_v96_
                nw146_[int(9)] = d_788_v96_
                nw146_[int(10)] = (d_788_v96_).set(default__.safeIndex((d_789_v97_).cardinality, len(d_788_v96_)), self.f7)
                nw146_[int(11)] = d_788_v96_
                nw146_[int(12)] = d_788_v96_
                nw146_[int(13)] = (d_788_v96_) + (d_788_v96_)
                nw146_[int(14)] = ((d_792_v100_)[d_791_v99_] if (d_791_v99_) in (d_792_v100_) else d_788_v96_)
                nw146_[int(15)] = d_788_v96_
                nw146_[int(16)] = (_dafny.SeqWithoutIsStrInference([((d_793_v101_)[p0] if (p0) in (d_793_v101_) else p0), p0, p0])) + (d_788_v96_)
                nw146_[int(17)] = d_788_v96_
                nw146_[int(18)] = _dafny.SeqWithoutIsStrInference([p2, self.f7])
                nw146_[int(19)] = _dafny.SeqWithoutIsStrInference([p2, True])
                nw146_[int(20)] = d_788_v96_
                nw146_[int(21)] = d_788_v96_
                d_794_v102_ = nw146_
                index133_ = default__.safeIndex(947, (d_794_v102_).length(0))
                (d_794_v102_)[index133_] = _dafny.SeqWithoutIsStrInference([not(p0)])
                index134_ = default__.safeIndex(947, (d_794_v102_).length(0))
                (d_794_v102_)[index134_] = d_788_v96_
                d_795_v103_: int
                d_795_v103_ = -377
                d_795_v103_ = d_795_v103_
                d_796_v104_: _dafny.Array
                nw147_ = _dafny.Array(int(0), 20)
                d_796_v104_ = nw147_
                d_797_v105_: D9
                d_797_v105_ = D9_DC23(d_796_v104_)
                d_796_v104_ = (d_797_v105_).cf44
                d_798_v107_: _dafny.Set
                d_798_v107_ = _dafny.Set({p1, d_795_v103_})
                def iife85_():
                    coll23_ = _dafny.Set()
                    compr_23_: int
                    for compr_23_ in _dafny.IntegerRange(140, 998):
                        d_799_v106_: int = compr_23_
                        if ((140) <= (d_799_v106_)) and ((d_799_v106_) < (998)):
                            coll23_ = coll23_.union(_dafny.Set([(d_799_v106_) * (p1)]))
                    return _dafny.Set(coll23_)
                (self).f7 = not((p2 if self.f7 else (iife85_()
                ).isdisjoint(d_798_v107_)))
            elif True:
                (self).f7 = self.f7
                d_800_v108_: _dafny.Array
                def lambda58_(d_801_p1_):
                    def lambda59_(d_802_i14_):
                        return default__.safeModuloInt(d_802_i14_, d_801_p1_)

                    return lambda59_

                init32_ = lambda58_(p1)
                nw148_ = _dafny.Array(None, 7)
                for i0_32_ in range(nw148_.length(0)):
                    nw148_[i0_32_] = init32_(i0_32_)
                d_800_v108_ = nw148_
                d_803_v109_: _dafny.Map
                d_803_v109_ = _dafny.Map({d_800_v108_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guk"))})
                d_804_v110_: _dafny.Seq
                d_804_v110_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoh"))
                d_803_v109_ = (d_803_v109_).set(d_800_v108_, d_804_v110_)
                d_805_v111_: _dafny.Array
                nw149_ = _dafny.Array(_dafny.Map({}), 11)
                d_805_v111_ = nw149_
                d_806_v112_: _dafny.Array
                nw150_ = _dafny.Array(None, 12)
                nw150_[int(0)] = d_800_v108_
                nw150_[int(1)] = d_800_v108_
                nw150_[int(2)] = d_800_v108_
                nw150_[int(3)] = d_800_v108_
                nw150_[int(4)] = d_800_v108_
                nw150_[int(5)] = d_800_v108_
                nw150_[int(6)] = d_800_v108_
                nw150_[int(7)] = d_800_v108_
                nw150_[int(8)] = d_800_v108_
                nw150_[int(9)] = d_800_v108_
                nw150_[int(10)] = d_800_v108_
                nw150_[int(11)] = d_800_v108_
                d_806_v112_ = nw150_
                index135_ = default__.safeIndex(554, (d_806_v112_).length(0))
                (d_806_v112_)[index135_] = d_800_v108_
                d_807_v113_: int
                d_807_v113_ = -572
                d_808_v114_: _dafny.MultiSet
                d_808_v114_ = _dafny.MultiSet([(0) - (d_807_v113_)])
                index136_ = default__.safeIndex(554, (d_806_v112_).length(0))
                rhs135_ = p2
                rhs136_ = d_800_v108_
                rhs137_ = d_805_v111_
                rhs138_ = d_800_v108_
                rhs139_ = ((d_808_v114_)[(d_807_v113_) - (p1)] if ((d_807_v113_) - (p1)) in (d_808_v114_) else p1)
                lhs98_ = self
                lhs99_ = d_806_v112_
                lhs100_ = default__.safeIndex(554, (d_806_v112_).length(0))
                lhs98_.f7 = rhs135_
                d_800_v108_ = rhs136_
                d_805_v111_ = rhs137_
                lhs99_[lhs100_] = rhs138_
                d_807_v113_ = rhs139_
                d_809_v115_: _dafny.Array
                nw151_ = _dafny.Array(_dafny.Seq({}), 25)
                d_809_v115_ = nw151_
                d_809_v115_ = d_809_v115_
                d_810_v116_: _dafny.Map
                d_810_v116_ = _dafny.Map({self.f7: 855})
                arr0_ = (d_806_v112_)[default__.safeIndex(554, (d_806_v112_).length(0))]
                index137_ = default__.safeIndex(692, ((d_806_v112_)[default__.safeIndex(554, (d_806_v112_).length(0))]).length(0))
                arr0_[index137_] = len(d_810_v116_)
                arr1_ = (d_806_v112_)[default__.safeIndex(554, (d_806_v112_).length(0))]
                index138_ = default__.safeIndex(692, ((d_806_v112_)[default__.safeIndex(554, (d_806_v112_).length(0))]).length(0))
                arr1_[index138_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
            d_811_v117_: _dafny.Map
            d_811_v117_ = _dafny.Map({p2: self.f7})
            d_812_v118_: D0
            d_812_v118_ = D0_DC0(p1)
            d_813_v119_: _dafny.Seq
            d_813_v119_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yher"))
            r0 = ((d_811_v117_) | (_dafny.Map({default__.fm1(_dafny.SeqWithoutIsStrInference([d_812_v118_, d_812_v118_]), p1, len(d_813_v119_), globalState): p0}))) | (d_811_v117_)
            d_814_v120_: int
            d_814_v120_ = 120
            d_815_v121_: _dafny.Seq
            d_815_v121_ = _dafny.SeqWithoutIsStrInference([p1])
            d_814_v120_ = ((0) - (len(d_815_v121_))) * (p1)
            d_816_v122_: _dafny.Array
            nw152_ = _dafny.Array(int(0), 10)
            d_816_v122_ = nw152_
            d_817_v123_: _dafny.MultiSet
            d_817_v123_ = _dafny.MultiSet([d_814_v120_])
            index139_ = default__.safeIndex(565, (d_816_v122_).length(0))
            (d_816_v122_)[index139_] = (0) - (((d_817_v123_)[(d_815_v121_)[default__.safeIndex(p1, len(d_815_v121_))]] if ((d_815_v121_)[default__.safeIndex(p1, len(d_815_v121_))]) in (d_817_v123_) else len(d_813_v119_)))
            index140_ = default__.safeIndex(565, (d_816_v122_).length(0))
            (d_816_v122_)[index140_] = default__.safeModuloInt(p1, default__.safeModuloInt(p1, len(d_813_v119_)))
        hi1_ = (p1) + (p1)
        for d_818_i15_ in range(p1, hi1_):
            (self).f7 = not(p2)
            d_819_v124_: _dafny.Array
            nw153_ = _dafny.Array(False, 17)
            d_819_v124_ = nw153_
            index141_ = default__.safeIndex(522, (d_819_v124_).length(0))
            (d_819_v124_)[index141_] = self.f7
            d_820_v125_: _dafny.Set
            d_820_v125_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwpthm"))})
            index142_ = default__.safeIndex(522, (d_819_v124_).length(0))
            (d_819_v124_)[index142_] = (d_820_v125_).isdisjoint(d_820_v125_)
            d_821_v126_: int
            d_821_v126_ = -90
            d_822_v127_: _dafny.Seq
            d_822_v127_ = _dafny.SeqWithoutIsStrInference([(d_819_v124_)[default__.safeIndex(522, (d_819_v124_).length(0))], (d_819_v124_)[default__.safeIndex(522, (d_819_v124_).length(0))], p0])
            d_821_v126_ = (len(_dafny.SeqWithoutIsStrInference([(d_822_v127_)[default__.safeIndex(d_821_v126_, len(d_822_v127_))], p0, p0]))) + (len(d_822_v127_))
            d_823_v128_: D6
            d_823_v128_ = D6_DC16(p0)
            d_823_v128_ = d_823_v128_
        d_824_i16_: int
        d_824_i16_ = 0
        with _dafny.label("2"):
            while (default__.fm0(globalState)) == (p1):
                with _dafny.c_label("2"):
                    if (d_824_i16_) >= (100):
                        raise _dafny.Break("2")
                    d_824_i16_ = (d_824_i16_) + (1)
                    d_825_v129_: int
                    d_825_v129_ = -698
                    d_825_v129_ = (0) - ((default__.safeDivisionInt(d_825_v129_, d_825_v129_)) * (default__.safeModuloInt(d_825_v129_, 198)))
                    d_826_v130_: str
                    d_826_v130_ = _dafny.CodePoint('p')
                    d_827_v131_: _dafny.Seq
                    d_827_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmp"))
                    d_826_v130_ = ((d_827_v131_)[default__.safeIndex(p1, len(d_827_v131_))] if p2 else (self).f16)
                    d_828_v132_: _dafny.Array
                    nw154_ = _dafny.Array(False, 4)
                    d_828_v132_ = nw154_
                    d_829_v133_: _dafny.Map
                    d_829_v133_ = _dafny.Map({d_828_v132_: d_827_v131_})
                    d_830_v134_: D3
                    d_830_v134_ = D3_DC10(d_828_v132_)
                    d_829_v133_ = (d_829_v133_).set((d_830_v134_).cf21, d_827_v131_)
                    d_831_v135_: _dafny.Seq
                    d_831_v135_ = _dafny.SeqWithoutIsStrInference([D0_DC0(d_825_v129_)])
                    d_832_v136_: D0
                    d_832_v136_ = D0_DC0((0) - (p1))
                    d_833_v137_: _dafny.Map
                    d_833_v137_ = _dafny.Map({d_828_v132_: p1})
                    d_834_v138_: _dafny.Map
                    d_834_v138_ = _dafny.Map({default__.fm1(((d_831_v135_).set(default__.safeIndex(d_825_v129_, len(d_831_v135_)), d_832_v136_)).set(default__.safeIndex((0) - (p1), len((d_831_v135_).set(default__.safeIndex(d_825_v129_, len(d_831_v135_)), d_832_v136_))), d_832_v136_), (0) - (d_825_v129_), len(d_833_v137_), globalState): _dafny.SeqWithoutIsStrInference([d_827_v131_ for d_835_i17_ in range(default__.abs(485))])})
                    d_836_v139_: _dafny.Seq
                    d_836_v139_ = _dafny.SeqWithoutIsStrInference([self.f7])
                    d_837_v140_: _dafny.Map
                    d_837_v140_ = _dafny.Map({True: (D3_DC11(len(d_836_v139_), p1, p1, self.f7)).cf25})
                    d_838_v141_: _dafny.Seq
                    d_838_v141_ = _dafny.SeqWithoutIsStrInference([d_827_v131_])
                    d_834_v138_ = (d_834_v138_).set((d_837_v140_) == (d_837_v140_), ((d_838_v141_).set(default__.safeIndex(d_825_v129_, len(d_838_v141_)), d_827_v131_)) + (d_838_v141_))
                    pass
            pass
        d_839_v142_: D0
        d_839_v142_ = D0_DC0(p1)
        d_840_v143_: _dafny.Array
        nw155_ = _dafny.Array(None, 2)
        nw155_[int(0)] = p1
        nw155_[int(1)] = (d_839_v142_).cf0
        d_840_v143_ = nw155_
        index143_ = default__.safeIndex(210, (d_840_v143_).length(0))
        (d_840_v143_)[index143_] = p1
        index144_ = default__.safeIndex(210, (d_840_v143_).length(0))
        (d_840_v143_)[index144_] = p1
        d_841_v144_: D6
        d_841_v144_ = D6_DC16(True)
        d_842_v145_: _dafny.Map
        d_842_v145_ = _dafny.Map({(d_841_v144_).cf30: (d_840_v143_)[default__.safeIndex(210, (d_840_v143_).length(0))]})
        d_842_v145_ = (d_842_v145_).set(self.f7, p1)
        d_843_v146_: D0
        d_843_v146_ = D0_DC1(-88)
        pat_let_tv15_ = p0
        pat_let_tv16_ = p2
        pat_let_tv17_ = p0
        pat_let_tv18_ = p2
        pat_let_tv19_ = p2
        pat_let_tv20_ = p2
        pat_let_tv21_ = p2
        pat_let_tv22_ = p2
        pat_let_tv23_ = p0
        def lambda60_(source3_):
            if source3_.is_DC1:
                d_844___mcc_h7_ = source3_.cf1
                d_845_cf1_ = d_844___mcc_h7_
                return (((_dafny.Map({pat_let_tv15_: self.f7})).set(True, pat_let_tv16_)).set(pat_let_tv17_, self.f7)) | (_dafny.Map({pat_let_tv18_: pat_let_tv19_}))
            elif source3_.is_DC2:
                d_846___mcc_h8_ = source3_.cf2
                d_847___mcc_h9_ = source3_.cf3
                d_848___mcc_h10_ = source3_.cf4
                d_849___mcc_h11_ = source3_.cf5
                d_850_cf5_ = d_849___mcc_h11_
                d_851_cf4_ = d_848___mcc_h10_
                d_852_cf3_ = d_847___mcc_h9_
                d_853_cf2_ = d_846___mcc_h8_
                return _dafny.Map({d_853_cf2_: pat_let_tv20_})
            elif True:
                d_854___mcc_h12_ = source3_.cf0
                d_855_cf0_ = d_854___mcc_h12_
                return (_dafny.Map({pat_let_tv21_: self.f7})).set((D6_DC17(pat_let_tv22_, 742, (self).f16, (self).f16)).cf31, pat_let_tv23_)

        r0 = lambda60_(d_843_v146_)
        d_856_v147_: _dafny.Set
        d_856_v147_ = _dafny.Set({self.f7})
        r1 = (d_856_v147_).intersection((d_856_v147_).intersection(d_856_v147_))
        return r0, r1

    def m3(self, globalState):
        d_857_v0_: str
        d_857_v0_ = _dafny.CodePoint('d')
        d_857_v0_ = d_857_v0_
        d_858_v1_: _dafny.Seq
        d_858_v1_ = _dafny.SeqWithoutIsStrInference([False])
        d_859_v2_: _dafny.Seq
        d_859_v2_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_860_v3_: int
        d_860_v3_ = 504
        d_861_v4_: D8
        d_861_v4_ = D8_DC21(len(d_858_v1_), (d_859_v2_)[default__.safeIndex(d_860_v3_, len(d_859_v2_))], d_860_v3_)
        source4_ = d_861_v4_
        if source4_.is_DC21:
            d_862___mcc_h0_ = source4_.cf38
            d_863___mcc_h1_ = source4_.cf39
            d_864___mcc_h2_ = source4_.cf40
            d_865_cf40_ = d_864___mcc_h2_
            d_866_cf39_ = d_863___mcc_h1_
            d_867_cf38_ = d_862___mcc_h0_
            d_868_v5_: C0
            nw156_ = C0()
            nw156_.ctor__(not(True), self.f7)
            d_868_v5_ = nw156_
            d_869_v6_: _dafny.Array
            nw157_ = _dafny.Array(_dafny.Map({}), 1)
            d_869_v6_ = nw157_
            d_870_v7_: _dafny.Array
            d_870_v7_ = d_869_v6_
            d_871_v8_: _dafny.Seq
            d_871_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwqfqdd"))
            d_872_v9_: _dafny.Map
            d_872_v9_ = _dafny.Map({d_870_v7_: d_871_v8_})
            d_873_v10_: _dafny.Set
            d_873_v10_ = _dafny.Set({d_868_v5_.f13, self.f7, (d_868_v5_).f12, self.f7, d_868_v5_.f13})
            d_874_v11_: D1
            d_874_v11_ = D1_DC3(d_873_v10_)
            d_875_v12_: D0
            d_875_v12_ = D0_DC0(d_860_v3_)
            d_876_v13_: D8
            d_876_v13_ = D8_DC22(d_865_cf40_, d_874_v11_, d_875_v12_)
            d_877_v14_: _dafny.Map
            d_877_v14_ = _dafny.Map({d_876_v13_: d_867_cf38_})
            d_878_v15_: D8
            d_878_v15_ = D8_DC22((len(((d_872_v9_)[d_869_v6_] if (d_869_v6_) in (d_872_v9_) else d_871_v8_))) + (((d_877_v14_)[d_876_v13_] if (d_876_v13_) in (d_877_v14_) else d_867_cf38_)), D1_DC3(d_873_v10_), d_875_v12_)
            source5_ = d_878_v15_
            if source5_.is_DC21:
                d_879___mcc_h7_ = source5_.cf38
                d_880___mcc_h8_ = source5_.cf39
                d_881___mcc_h9_ = source5_.cf40
                d_882_cf40_ = d_881___mcc_h9_
                d_883_cf39_ = d_880___mcc_h8_
                d_884_cf38_ = d_879___mcc_h7_
                d_885_v16_: C0
                nw158_ = C0()
                nw158_.ctor__(True, False)
                d_885_v16_ = nw158_
                d_857_v0_ = (self).f16
                (d_868_v5_).f13 = ((d_868_v5_).f12 if self.f7 else True)
                d_886_v17_: _dafny.Map
                d_886_v17_ = _dafny.Map({d_882_cf40_: len(d_871_v8_)})
                d_887_v18_: _dafny.Seq
                d_887_v18_ = _dafny.SeqWithoutIsStrInference([len(d_873_v10_), len(d_886_v17_)])
                d_888_v19_: D1
                d_888_v19_ = D1_DC5((d_868_v5_).f12, d_887_v18_, d_882_cf40_)
                d_889_v20_: C0
                nw159_ = C0()
                nw159_.ctor__((d_860_v3_) >= (d_884_cf38_), (d_888_v19_).cf9)
                d_889_v20_ = nw159_
            elif source5_.is_DC22:
                d_890___mcc_h10_ = source5_.cf41
                d_891___mcc_h11_ = source5_.cf42
                d_892___mcc_h12_ = source5_.cf43
                d_893_cf43_ = d_892___mcc_h12_
                d_894_cf42_ = d_891___mcc_h11_
                d_895_cf41_ = d_890___mcc_h10_
                d_896_v22_: _dafny.Array
                def lambda61_(d_897_v8_, d_898_cf41_, d_899_cf40_):
                    def lambda62_(d_900_i0_):
                        def iife86_():
                            coll24_ = _dafny.Map()
                            compr_24_: int
                            for compr_24_ in _dafny.IntegerRange(261, 923):
                                d_901_v21_: int = compr_24_
                                if ((261) <= (d_901_v21_)) and ((d_901_v21_) < (923)):
                                    coll24_[default__.safeModuloInt(d_901_v21_, (0) - (d_898_cf41_))] = d_899_cf40_
                            return _dafny.Map(coll24_)
                        return _dafny.SeqWithoutIsStrInference([454, len(iife86_()
                        ), len(d_897_v8_)])

                    return lambda62_

                init33_ = lambda61_(d_871_v8_, d_895_cf41_, d_865_cf40_)
                nw160_ = _dafny.Array(None, 17)
                for i0_33_ in range(nw160_.length(0)):
                    nw160_[i0_33_] = init33_(i0_33_)
                d_896_v22_ = nw160_
                d_896_v22_ = (d_896_v22_ if d_868_v5_.f13 else d_896_v22_)
                d_895_cf41_ = d_895_cf41_
                d_858_v1_ = d_858_v1_
                d_860_v3_ = (d_895_cf41_) * (d_895_cf41_)
            elif True:
                d_902___mcc_h13_ = source5_.cf37
                d_903_cf37_ = d_902___mcc_h13_
                d_904_v23_: _dafny.Map
                d_904_v23_ = _dafny.Map({d_860_v3_: d_868_v5_.f13})
                d_905_v24_: _dafny.Map
                d_905_v24_ = _dafny.Map({default__.fm29(d_865_cf40_, self.f7, globalState): d_904_v23_})
                rhs140_ = d_905_v24_
                rhs141_ = 342
                d_905_v24_ = rhs140_
                d_865_cf40_ = rhs141_
                d_906_v25_: _dafny.Array
                nw161_ = _dafny.Array(D2.default()(), 26)
                d_906_v25_ = nw161_
                d_907_v26_: _dafny.Set
                d_907_v26_ = _dafny.Set({984})
                d_908_v27_: D2
                d_908_v27_ = D2_DC7(d_907_v26_)
                index145_ = default__.safeIndex(674, (d_906_v25_).length(0))
                (d_906_v25_)[index145_] = d_908_v27_
                index146_ = default__.safeIndex(674, (d_906_v25_).length(0))
                (d_906_v25_)[index146_] = d_908_v27_
                d_909_v28_: _dafny.Array
                nw162_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_909_v28_ = nw162_
                d_910_v29_: _dafny.Seq
                d_910_v29_ = _dafny.SeqWithoutIsStrInference([(d_909_v28_ if (d_858_v1_)[default__.safeIndex(d_865_cf40_, len(d_858_v1_))] else d_909_v28_), d_909_v28_])
                d_909_v28_ = (d_910_v29_)[default__.safeIndex((d_865_cf40_) * (d_867_cf38_), len(d_910_v29_))]
                index147_ = default__.safeIndex(675, (d_909_v28_).length(0))
                (d_909_v28_)[index147_] = d_857_v0_
                index148_ = default__.safeIndex(675, (d_909_v28_).length(0))
                (d_909_v28_)[index148_] = (self).f16
            (self).f7 = not (self.f7) or ((self.f7) == ((d_868_v5_).f12))
            d_911_v30_: _dafny.MultiSet
            d_911_v30_ = _dafny.MultiSet([d_860_v3_, len((d_858_v1_).set(default__.safeIndex(d_867_cf38_, len(d_858_v1_)), self.f7))])
            d_912_v31_: _dafny.Array
            nw163_ = _dafny.Array(None, 3)
            nw163_[int(0)] = ((d_911_v30_)[d_865_cf40_] if (d_865_cf40_) in (d_911_v30_) else d_865_cf40_)
            nw163_[int(1)] = (d_860_v3_) * (d_867_cf38_)
            nw163_[int(2)] = len((d_871_v8_) + (d_871_v8_))
            d_912_v31_ = nw163_
            index149_ = default__.safeIndex(868, (d_912_v31_).length(0))
            (d_912_v31_)[index149_] = (d_860_v3_) + (d_860_v3_)
            index150_ = default__.safeIndex(868, (d_912_v31_).length(0))
            (d_912_v31_)[index150_] = (0) - (d_865_cf40_)
        elif source4_.is_DC22:
            d_913___mcc_h3_ = source4_.cf41
            d_914___mcc_h4_ = source4_.cf42
            d_915___mcc_h5_ = source4_.cf43
            d_916_cf43_ = d_915___mcc_h5_
            d_917_cf42_ = d_914___mcc_h4_
            d_918_cf41_ = d_913___mcc_h3_
            d_919_v32_: _dafny.Set
            d_919_v32_ = _dafny.Set({(self).f16})
            (self).f7 = (default__.fm30(self.f7, d_860_v3_, self.f7, globalState)).isdisjoint(d_919_v32_)
            d_920_v33_: _dafny.Map
            d_920_v33_ = _dafny.Map({d_916_cf43_: 437})
            d_921_v34_: D2
            d_921_v34_ = D2_DC9(532, d_918_cf41_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuln")), self.f7, d_920_v33_)
            (self).f7 = (d_921_v34_).cf19
            if (d_858_v1_)[default__.safeIndex(d_860_v3_, len(d_858_v1_))]:
                d_922_v35_: _dafny.Seq
                d_922_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "reopw"))
                rhs142_ = default__.safeModuloInt(d_918_cf41_, len(_dafny.SeqWithoutIsStrInference([self.f7])))
                rhs143_ = d_922_v35_
                d_860_v3_ = rhs142_
                d_922_v35_ = rhs143_
                d_923_v36_: T0
                nw164_ = C1()
                nw164_.ctor__(self.f7)
                d_923_v36_ = nw164_
                d_924_v37_: _dafny.Map
                d_924_v37_ = _dafny.Map({d_923_v36_: d_860_v3_})
                d_925_v38_: _dafny.MultiSet
                d_925_v38_ = _dafny.MultiSet([len(d_924_v37_), d_860_v3_])
                (self).f7 = (self.f7) == ((((d_925_v38_)[d_918_cf41_] if (d_918_cf41_) in (d_925_v38_) else d_918_cf41_)) >= (d_860_v3_))
                d_918_cf41_ = default__.safeModuloInt(d_860_v3_, d_860_v3_)
                d_926_v39_: _dafny.Map
                d_926_v39_ = _dafny.Map({(0) - ((d_860_v3_) + (len(_dafny.SeqWithoutIsStrInference([d_857_v0_ for d_927_i1_ in range(default__.abs(128))])))): d_860_v3_})
                d_926_v39_ = (d_926_v39_).set(d_860_v3_, d_918_cf41_)
                d_928_v40_: _dafny.Map
                d_928_v40_ = _dafny.Map({not(self.f7): d_918_cf41_})
                d_860_v3_ = (d_860_v3_) + (((d_928_v40_)[d_923_v36_.f7] if (d_923_v36_.f7) in (d_928_v40_) else d_918_cf41_))
            elif True:
                d_929_v41_: _dafny.Array
                nw165_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                d_929_v41_ = nw165_
                d_930_v42_: _dafny.Seq
                d_930_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yt"))
                index151_ = default__.safeIndex(343, (d_929_v41_).length(0))
                (d_929_v41_)[index151_] = (d_930_v42_) + (d_930_v42_)
                index152_ = default__.safeIndex(343, (d_929_v41_).length(0))
                (d_929_v41_)[index152_] = d_930_v42_
                d_931_v43_: _dafny.Set
                d_931_v43_ = _dafny.Set({self.f7, self.f7})
                d_932_v44_: _dafny.Set
                d_932_v44_ = _dafny.Set({263, len(d_931_v43_)})
                d_933_v45_: _dafny.Seq
                d_933_v45_ = _dafny.SeqWithoutIsStrInference([len(d_932_v44_)])
                d_934_v46_: D9
                d_934_v46_ = D9_DC26(d_933_v45_)
                d_934_v46_ = d_934_v46_
                d_935_v47_: _dafny.Array
                nw166_ = _dafny.Array(int(0), 15)
                d_935_v47_ = nw166_
                index153_ = default__.safeIndex(370, (d_935_v47_).length(0))
                (d_935_v47_)[index153_] = d_918_cf41_
                index154_ = default__.safeIndex(370, (d_935_v47_).length(0))
                (d_935_v47_)[index154_] = d_918_cf41_
                d_860_v3_ = d_918_cf41_
                d_936_v48_: _dafny.Map
                d_936_v48_ = _dafny.Map({self.f7: len((d_929_v41_)[default__.safeIndex(343, (d_929_v41_).length(0))])})
                d_936_v48_ = (d_936_v48_).set(self.f7, d_918_cf41_)
            d_937_v49_: _dafny.Seq
            d_937_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_938_v50_: _dafny.Map
            d_938_v50_ = _dafny.Map({d_918_cf41_: d_937_v49_})
            d_937_v49_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ix"))) + (((d_938_v50_)[d_860_v3_] if (d_860_v3_) in (d_938_v50_) else d_937_v49_))
        elif True:
            d_939___mcc_h6_ = source4_.cf37
            d_940_cf37_ = d_939___mcc_h6_
            d_941_v51_: _dafny.Map
            d_941_v51_ = _dafny.Map({self.f7: self.f7})
            d_941_v51_ = (d_941_v51_).set((d_860_v3_) != (-684), self.f7)
            d_942_v52_: C0
            nw167_ = C0()
            nw167_.ctor__((len(_dafny.SeqWithoutIsStrInference([(self).f16 for d_943_i2_ in range(default__.abs(863))]))) >= (d_860_v3_), (d_860_v3_) < (d_860_v3_))
            d_942_v52_ = nw167_
            d_944_v53_: _dafny.Map
            d_944_v53_ = _dafny.Map({d_860_v3_: (d_942_v52_).f12})
            d_945_v54_: _dafny.Array
            nw168_ = _dafny.Array(None, 3)
            nw168_[int(0)] = (len(_dafny.SeqWithoutIsStrInference([d_857_v0_ for d_946_i3_ in range(default__.abs(420))]))) not in (d_944_v53_)
            nw168_[int(1)] = d_942_v52_.f13
            nw168_[int(2)] = d_942_v52_.f13
            d_945_v54_ = nw168_
            index155_ = default__.safeIndex(275, (d_945_v54_).length(0))
            (d_945_v54_)[index155_] = (d_942_v52_).f12
            index156_ = default__.safeIndex(275, (d_945_v54_).length(0))
            (d_945_v54_)[index156_] = not((d_942_v52_).f12)
            d_947_v55_: _dafny.Set
            d_947_v55_ = _dafny.Set({(d_944_v53_) | (d_944_v53_), d_944_v53_})
            d_948_v56_: _dafny.Seq
            d_948_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rck"))
            d_949_v58_: _dafny.Map
            d_949_v58_ = _dafny.Map({d_860_v3_: default__.fm28(True, (0) - (-784), globalState)})
            def iife87_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(667, 221):
                    d_950_v57_: int = compr_25_
                    if ((667) <= (d_950_v57_)) and ((d_950_v57_) < (221)):
                        coll25_[default__.safeDivisionInt(d_950_v57_, d_860_v3_)] = d_942_v52_.f13
                return _dafny.Map(coll25_)
            rhs144_ = (d_947_v55_) - (_dafny.Set({iife87_()
            , d_944_v53_}))
            rhs145_ = ((d_949_v58_)[(d_860_v3_ if self.f7 else d_860_v3_)] if ((d_860_v3_ if self.f7 else d_860_v3_)) in (d_949_v58_) else d_948_v56_)
            rhs146_ = default__.fm0(globalState)
            rhs147_ = d_857_v0_
            d_947_v55_ = rhs144_
            d_948_v56_ = rhs145_
            d_860_v3_ = rhs146_
            d_857_v0_ = rhs147_
        d_951_v59_: C1
        nw169_ = C1()
        nw169_.ctor__(self.f7)
        d_951_v59_ = nw169_
        d_952_v60_: _dafny.Array
        nw170_ = _dafny.Array(False, 3)
        d_952_v60_ = nw170_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_952_v60_).length(0)):
            d_953_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_953_i4_)) and ((d_953_i4_) < ((d_952_v60_).length(0)))):
                (d_952_v60_)[(d_953_i4_)] = ((_dafny.MultiSet([d_860_v3_])) | (_dafny.MultiSet([len(_dafny.Map({self.f7: d_860_v3_}))]))).issubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltjmuktis")))]))
        d_954_i5_: int
        d_954_i5_ = 0
        with _dafny.label("3"):
            while ((0) - (len(d_858_v1_))) != (373):
                with _dafny.c_label("3"):
                    if (d_954_i5_) >= (100):
                        raise _dafny.Break("3")
                    d_954_i5_ = (d_954_i5_) + (1)
                    d_955_v61_: _dafny.Array
                    nw171_ = _dafny.Array(_dafny.CodePoint('D'), 9)
                    d_955_v61_ = nw171_
                    index157_ = default__.safeIndex(23, (d_955_v61_).length(0))
                    (d_955_v61_)[index157_] = d_857_v0_
                    index158_ = default__.safeIndex(23, (d_955_v61_).length(0))
                    (d_955_v61_)[index158_] = d_857_v0_
                    index159_ = default__.safeIndex(23, (d_955_v61_).length(0))
                    (d_955_v61_)[index159_] = (d_857_v0_ if self.f7 else default__.fm37(self.f7, D9_DC25(), globalState))
                    d_956_v62_: _dafny.Map
                    d_956_v62_ = _dafny.Map({d_860_v3_: self.f7})
                    d_957_v63_: _dafny.Map
                    d_957_v63_ = _dafny.Map({d_860_v3_: d_860_v3_})
                    d_956_v62_ = (d_956_v62_).set(((d_957_v63_)[d_860_v3_] if (d_860_v3_) in (d_957_v63_) else d_860_v3_), self.f7)
                    index160_ = default__.safeIndex(23, (d_955_v61_).length(0))
                    (d_955_v61_)[index160_] = d_857_v0_
                    pass
            pass
        d_958_v64_: _dafny.Array
        def lambda63_(d_959_v3_, d_960_v0_):
            def lambda64_(d_961_i6_):
                return D6_DC17(self.f7, d_959_v3_, (self).f16, d_960_v0_)

            return lambda64_

        init34_ = lambda63_(d_860_v3_, d_857_v0_)
        nw172_ = _dafny.Array(None, 16)
        for i0_34_ in range(nw172_.length(0)):
            nw172_[i0_34_] = init34_(i0_34_)
        d_958_v64_ = nw172_
        d_962_v65_: D6
        d_962_v65_ = D6_DC17(self.f7, d_860_v3_, (self).f16, _dafny.CodePoint('b'))
        index161_ = default__.safeIndex(681, (d_958_v64_).length(0))
        (d_958_v64_)[index161_] = d_962_v65_
        index162_ = default__.safeIndex(681, (d_958_v64_).length(0))
        (d_958_v64_)[index162_] = default__.fm42(len(_dafny.Map({self.f7: _dafny.CodePoint('j')})), (d_860_v3_) <= (d_860_v3_), globalState)

    def m13(self, p0, p1, p2, p3, globalState):
        d_963_v0_: _dafny.Seq
        d_963_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_964_v1_: int
        d_964_v1_ = 608
        d_965_v2_: _dafny.Seq
        d_965_v2_ = _dafny.SeqWithoutIsStrInference([(d_963_v0_).set(default__.safeIndex(d_964_v1_, len(d_963_v0_)), self.f7), d_963_v0_])
        (self).f7 = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f7]) for d_966_i0_ in range(default__.abs(44))])) <= (d_965_v2_)
        d_967_v3_: _dafny.Map
        d_967_v3_ = _dafny.Map({d_964_v1_: d_964_v1_})
        d_968_v4_: _dafny.Map
        d_968_v4_ = _dafny.Map({p2: d_964_v1_})
        hi2_ = len(d_968_v4_)
        for d_969_i1_ in range(len(d_967_v3_), hi2_):
            d_970_v5_: _dafny.Seq
            d_970_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_969_i1_: default__.fm0(globalState)}), (d_967_v3_) | (default__.fm29(default__.fm0(globalState), p0, globalState))])
            d_971_v6_: _dafny.Seq
            d_971_v6_ = _dafny.SeqWithoutIsStrInference([d_969_i1_, -871])
            rhs148_ = (d_971_v6_)[default__.safeIndex(d_964_v1_, len(d_971_v6_))]
            rhs149_ = p1
            rhs150_ = d_964_v1_
            rhs151_ = d_970_v5_
            lhs101_ = self
            d_964_v1_ = rhs148_
            lhs101_.f7 = rhs149_
            d_964_v1_ = rhs150_
            d_970_v5_ = rhs151_
            d_964_v1_ = d_969_i1_
            (self).f7 = p0
            d_972_v7_: D6
            d_972_v7_ = D6_DC16(p0)
            (self).f7 = (d_972_v7_).cf30
        d_973_v8_: _dafny.Map
        d_973_v8_ = _dafny.Map({self.f7: self.f7})
        d_974_v9_: _dafny.Map
        d_974_v9_ = _dafny.Map({d_964_v1_: p1})
        d_975_i2_: int
        d_975_i2_ = 0
        with _dafny.label("4"):
            while ((d_973_v8_)[((d_974_v9_)[d_964_v1_] if (d_964_v1_) in (d_974_v9_) else p1)] if (((d_974_v9_)[d_964_v1_] if (d_964_v1_) in (d_974_v9_) else p1)) in (d_973_v8_) else not(False)):
                with _dafny.c_label("4"):
                    if (d_975_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_975_i2_ = (d_975_i2_) + (1)
                    d_976_v10_: _dafny.Array
                    nw173_ = _dafny.Array(_dafny.Array(None, 0), 21)
                    d_976_v10_ = nw173_
                    d_977_v11_: _dafny.Set
                    d_977_v11_ = _dafny.Set({d_964_v1_, 551})
                    d_978_v12_: _dafny.Seq
                    d_978_v12_ = _dafny.SeqWithoutIsStrInference([(self).f16])
                    d_979_v13_: _dafny.Set
                    d_979_v13_ = _dafny.Set({p1, self.f7})
                    d_980_v14_: D1
                    d_980_v14_ = D1_DC3(d_979_v13_)
                    d_981_v15_: D0
                    d_981_v15_ = D0_DC0(len(d_968_v4_))
                    d_982_v16_: D8
                    d_982_v16_ = D8_DC22(d_964_v1_, d_980_v14_, d_981_v15_)
                    d_983_v17_: _dafny.Array
                    nw174_ = _dafny.Array(None, 21)
                    nw174_[int(0)] = d_964_v1_
                    nw174_[int(1)] = d_964_v1_
                    nw174_[int(2)] = default__.fm0(globalState)
                    nw174_[int(3)] = d_964_v1_
                    nw174_[int(4)] = d_964_v1_
                    nw174_[int(5)] = d_964_v1_
                    nw174_[int(6)] = len(d_977_v11_)
                    nw174_[int(7)] = d_964_v1_
                    nw174_[int(8)] = d_964_v1_
                    nw174_[int(9)] = len(_dafny.Map({730: len(d_963_v0_)}))
                    nw174_[int(10)] = d_964_v1_
                    nw174_[int(11)] = d_964_v1_
                    nw174_[int(12)] = len(d_963_v0_)
                    nw174_[int(13)] = d_964_v1_
                    nw174_[int(14)] = len(d_978_v12_)
                    nw174_[int(15)] = 800
                    nw174_[int(16)] = d_964_v1_
                    nw174_[int(17)] = d_964_v1_
                    nw174_[int(18)] = default__.fm0(globalState)
                    nw174_[int(19)] = (d_982_v16_).cf41
                    nw174_[int(20)] = d_964_v1_
                    d_983_v17_ = nw174_
                    d_984_v18_: _dafny.Seq
                    d_984_v18_ = _dafny.SeqWithoutIsStrInference([d_983_v17_, d_983_v17_])
                    index163_ = default__.safeIndex(773, (d_976_v10_).length(0))
                    (d_976_v10_)[index163_] = (d_984_v18_)[default__.safeIndex(d_964_v1_, len(d_984_v18_))]
                    index164_ = default__.safeIndex(773, (d_976_v10_).length(0))
                    (d_976_v10_)[index164_] = d_983_v17_
                    if p2:
                        d_985_v19_: str
                        d_985_v19_ = _dafny.CodePoint('b')
                        d_985_v19_ = (self).f16
                        d_985_v19_ = (self).f16
                        d_986_v20_: _dafny.Array
                        def lambda65_(d_987_v12_):
                            def lambda66_(d_988_i3_):
                                return (_dafny.SeqWithoutIsStrInference([803])) + (_dafny.SeqWithoutIsStrInference([len(d_987_v12_) for d_989_i4_ in range(default__.abs(429))]))

                            return lambda66_

                        init35_ = lambda65_(d_978_v12_)
                        nw175_ = _dafny.Array(None, 3)
                        for i0_35_ in range(nw175_.length(0)):
                            nw175_[i0_35_] = init35_(i0_35_)
                        d_986_v20_ = nw175_
                        def lambda67_(d_990_v4_):
                            def lambda68_(d_991_i5_):
                                return _dafny.SeqWithoutIsStrInference([len(d_990_v4_)])

                            return lambda68_

                        init36_ = lambda67_(d_968_v4_)
                        nw176_ = _dafny.Array(None, 14)
                        for i0_36_ in range(nw176_.length(0)):
                            nw176_[i0_36_] = init36_(i0_36_)
                        d_986_v20_ = nw176_
                        d_964_v1_ = d_964_v1_
                        d_978_v12_ = d_978_v12_
                    elif True:
                        d_992_v21_: _dafny.Seq
                        d_992_v21_ = _dafny.SeqWithoutIsStrInference([d_981_v15_])
                        d_993_v22_: _dafny.Set
                        d_993_v22_ = _dafny.Set({d_973_v8_, d_973_v8_, _dafny.Map({default__.fm1((d_992_v21_).set(default__.safeIndex(len(default__.fm43((self).f16, self.f7, p2, globalState)), len(d_992_v21_)), D0_DC0(d_964_v1_)), d_964_v1_, d_964_v1_, globalState): not(not(True))})})
                        rhs152_ = len(((d_978_v12_) + ((d_978_v12_) + (d_978_v12_))).set(default__.safeIndex(d_964_v1_, len((d_978_v12_) + ((d_978_v12_) + (d_978_v12_)))), (self).f16))
                        rhs153_ = d_993_v22_
                        d_964_v1_ = rhs152_
                        d_993_v22_ = rhs153_
                        index165_ = default__.safeIndex(773, (d_976_v10_).length(0))
                        (d_976_v10_)[index165_] = (d_976_v10_)[default__.safeIndex(773, (d_976_v10_).length(0))]
                        d_994_v23_: _dafny.Array
                        nw177_ = _dafny.Array(_dafny.Set({}), 17)
                        d_994_v23_ = nw177_
                        d_995_v24_: _dafny.Set
                        d_995_v24_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f16 for d_996_i6_ in range(default__.abs(322))])})
                        index166_ = default__.safeIndex(836, (d_994_v23_).length(0))
                        (d_994_v23_)[index166_] = d_995_v24_
                        index167_ = default__.safeIndex(836, (d_994_v23_).length(0))
                        (d_994_v23_)[index167_] = d_995_v24_
                        rhs154_ = d_964_v1_
                        rhs155_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qr"))) + (d_978_v12_)
                        d_964_v1_ = rhs154_
                        d_978_v12_ = rhs155_
                        d_997_v25_: _dafny.Array
                        nw178_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
                        d_997_v25_ = nw178_
                        index168_ = default__.safeIndex(673, (d_997_v25_).length(0))
                        (d_997_v25_)[index168_] = _dafny.SeqWithoutIsStrInference([(self).f16 for d_998_i7_ in range(default__.abs(772))])
                        index169_ = default__.safeIndex(673, (d_997_v25_).length(0))
                        (d_997_v25_)[index169_] = (d_978_v12_) + (d_978_v12_)
                    arr2_ = (d_976_v10_)[default__.safeIndex(773, (d_976_v10_).length(0))]
                    index170_ = default__.safeIndex(934, ((d_976_v10_)[default__.safeIndex(773, (d_976_v10_).length(0))]).length(0))
                    arr2_[index170_] = default__.fm0(globalState)
                    arr3_ = (d_976_v10_)[default__.safeIndex(773, (d_976_v10_).length(0))]
                    index171_ = default__.safeIndex(934, ((d_976_v10_)[default__.safeIndex(773, (d_976_v10_).length(0))]).length(0))
                    arr3_[index171_] = d_964_v1_
                    d_999_v26_: D6
                    d_999_v26_ = D6_DC16(True)
                    (self).f7 = (d_999_v26_).cf30
                    pass
            pass
        d_1000_v27_: str
        d_1000_v27_ = _dafny.CodePoint('a')
        d_1000_v27_ = d_1000_v27_
        d_1001_v28_: D0
        d_1001_v28_ = D0_DC0(d_964_v1_)
        d_1002_v29_: _dafny.Seq
        d_1002_v29_ = _dafny.SeqWithoutIsStrInference([d_1001_v28_, d_1001_v28_])
        if (default__.fm1(d_1002_v29_, d_964_v1_, len(d_973_v8_), globalState) if not((d_964_v1_) == (889)) else not (self.f7) or (True)):
            d_964_v1_ = d_964_v1_
            d_1003_v30_: C0
            nw179_ = C0()
            nw179_.ctor__(p0, not((p1 if p0 else p2)))
            d_1003_v30_ = nw179_
            d_1004_v31_: C0
            nw180_ = C0()
            nw180_.ctor__(self.f7, d_1003_v30_.f13)
            d_1004_v31_ = nw180_
            d_1005_v32_: _dafny.Seq
            d_1005_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpqblmqyy"))
            d_1006_v33_: _dafny.Seq
            d_1006_v33_ = _dafny.SeqWithoutIsStrInference([len(d_1005_v32_), d_964_v1_])
            d_1007_v34_: _dafny.Map
            d_1007_v34_ = _dafny.Map({d_1006_v33_: not((d_1003_v30_).f12)})
            d_1007_v34_ = (d_1007_v34_).set(d_1006_v33_, self.f7)
            d_964_v1_ = default__.fm0(globalState)
        elif True:
            d_964_v1_ = (0) - (d_964_v1_)
            d_1008_v35_: _dafny.MultiSet
            d_1008_v35_ = _dafny.MultiSet([d_964_v1_])
            d_1009_v36_: _dafny.MultiSet
            d_1009_v36_ = _dafny.MultiSet([d_964_v1_, (0) - ((d_1008_v35_).cardinality)])
            d_1010_v37_: _dafny.Map
            d_1010_v37_ = _dafny.Map({d_964_v1_: d_1009_v36_})
            d_1011_v38_: _dafny.Seq
            d_1011_v38_ = _dafny.SeqWithoutIsStrInference([d_967_v3_, d_967_v3_, d_967_v3_, d_967_v3_, _dafny.Map({d_964_v1_: d_964_v1_})])
            d_1012_v39_: _dafny.Seq
            d_1012_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvugstq"))
            source6_ = default__.fm44(default__.safeModuloInt(d_964_v1_, d_964_v1_), (_dafny.Map({d_964_v1_: d_1009_v36_})) | (d_1010_v37_), d_1000_v27_, (d_1011_v38_)[default__.safeIndex(len(default__.fm45(len(d_1012_v39_), len(d_1012_v39_), globalState)), len(d_1011_v38_))], globalState)
            d_1013___mcc_h0_ = source6_
            d_1014_cf47_ = d_1013___mcc_h0_
            d_1012_v39_ = d_1012_v39_
            d_1015_v40_: C1
            nw181_ = C1()
            nw181_.ctor__(((d_974_v9_)[868] if (868) in (d_974_v9_) else p1))
            d_1015_v40_ = nw181_
            d_1016_v41_: _dafny.Array
            nw182_ = _dafny.Array(int(0), 23)
            d_1016_v41_ = nw182_
            d_1017_v42_: _dafny.Map
            d_1017_v42_ = _dafny.Map({d_1016_v41_: _dafny.CodePoint('e')})
            d_1018_v43_: _dafny.Set
            d_1018_v43_ = _dafny.Set({self.f7})
            d_1019_v44_: _dafny.Map
            d_1019_v44_ = _dafny.Map({d_1018_v43_: d_967_v3_})
            rhs156_ = d_1017_v42_
            rhs157_ = (len((d_1019_v44_) | (d_1019_v44_))) - ((816) - (d_964_v1_))
            rhs158_ = not(p2)
            lhs102_ = self
            d_1017_v42_ = rhs156_
            d_964_v1_ = rhs157_
            lhs102_.f7 = rhs158_
            d_1020_v45_: T0
            nw183_ = C1()
            nw183_.ctor__(p2)
            d_1020_v45_ = nw183_
            d_1020_v45_ = d_1020_v45_
            d_1021_v46_: D1
            d_1021_v46_ = D1_DC4(p1, p2)
            d_1022_v47_: D1
            d_1022_v47_ = D1_DC6(d_1021_v46_)
            d_1022_v47_ = d_1022_v47_
            d_974_v9_ = (d_974_v9_).set(d_964_v1_, p1)
            if (default__.fm1(d_1002_v29_, 5, (0) - (len((self).f15)), globalState)) == (self.f7):
                d_964_v1_ = d_964_v1_
                d_1023_v48_: _dafny.Array
                nw184_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_1023_v48_ = nw184_
                d_1024_v49_: _dafny.Set
                d_1024_v49_ = _dafny.Set({d_1012_v39_})
                d_1025_v50_: _dafny.Seq
                d_1025_v50_ = _dafny.SeqWithoutIsStrInference([d_1009_v36_])
                d_1026_v51_: _dafny.Seq
                d_1026_v51_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1024_v49_)), ((d_1025_v50_)[default__.safeIndex(d_964_v1_, len(d_1025_v50_))]).cardinality])
                d_1027_v52_: D2
                d_1027_v52_ = D2_DC8(self.f7, d_1012_v39_)
                d_1028_v53_: _dafny.MultiSet
                d_1028_v53_ = _dafny.MultiSet([d_1027_v52_])
                d_1029_v54_: _dafny.Array
                nw185_ = _dafny.Array(None, 9)
                nw185_[int(0)] = d_964_v1_
                nw185_[int(1)] = d_964_v1_
                nw185_[int(2)] = d_964_v1_
                nw185_[int(3)] = len(d_968_v4_)
                nw185_[int(4)] = (d_1026_v51_)[default__.safeIndex(d_964_v1_, len(d_1026_v51_))]
                nw185_[int(5)] = d_964_v1_
                nw185_[int(6)] = d_964_v1_
                nw185_[int(7)] = d_964_v1_
                nw185_[int(8)] = (d_1028_v53_).cardinality
                d_1029_v54_ = nw185_
                d_1030_v55_: _dafny.Map
                d_1030_v55_ = _dafny.Map({p2: d_1029_v54_})
                nw186_ = _dafny.Array(None, 12)
                nw186_[int(0)] = d_1029_v54_
                nw186_[int(1)] = d_1029_v54_
                nw186_[int(2)] = d_1029_v54_
                nw186_[int(3)] = d_1029_v54_
                nw186_[int(4)] = d_1029_v54_
                nw186_[int(5)] = d_1029_v54_
                nw186_[int(6)] = ((d_1030_v55_)[p1] if (p1) in (d_1030_v55_) else d_1029_v54_)
                nw186_[int(7)] = d_1029_v54_
                nw186_[int(8)] = d_1029_v54_
                nw186_[int(9)] = d_1029_v54_
                nw186_[int(10)] = d_1029_v54_
                nw186_[int(11)] = d_1029_v54_
                d_1023_v48_ = nw186_
                d_1012_v39_ = d_1012_v39_
                d_1031_v56_: D6
                d_1031_v56_ = D6_DC17(self.f7, (d_1026_v51_)[default__.safeIndex(d_964_v1_, len(d_1026_v51_))], _dafny.CodePoint('b'), (self).f16)
                rhs159_ = (self).f16
                rhs160_ = ((d_1031_v56_).cf31 if p1 else p1)
                rhs161_ = _dafny.MultiSet([d_964_v1_])
                rhs162_ = (0) - (len((_dafny.Set({d_964_v1_, d_964_v1_})) | (default__.fm2(d_964_v1_, self.f7, globalState))))
                rhs163_ = (self).f16
                lhs103_ = self
                d_1000_v27_ = rhs159_
                lhs103_.f7 = rhs160_
                d_1009_v36_ = rhs161_
                d_964_v1_ = rhs162_
                d_1000_v27_ = rhs163_
                d_1032_v57_: _dafny.Array
                def lambda69_(d_1033_v0_, d_1034_v9_):
                    def lambda70_(d_1035_i8_):
                        return (d_1033_v0_)[default__.safeIndex(len(d_1034_v9_), len(d_1033_v0_))]

                    return lambda70_

                init37_ = lambda69_(d_963_v0_, d_974_v9_)
                nw187_ = _dafny.Array(None, 6)
                for i0_37_ in range(nw187_.length(0)):
                    nw187_[i0_37_] = init37_(i0_37_)
                d_1032_v57_ = nw187_
                d_1032_v57_ = d_1032_v57_
            elif True:
                d_1036_v58_: _dafny.Seq
                d_1036_v58_ = _dafny.SeqWithoutIsStrInference([-595])
                d_1037_v59_: _dafny.Set
                d_1037_v59_ = _dafny.Set({d_964_v1_, len(d_1036_v58_), d_964_v1_, (0) - ((d_1036_v58_)[default__.safeIndex((0) - (d_964_v1_), len(d_1036_v58_))]), (d_1009_v36_).cardinality})
                d_1038_v60_: C0
                nw188_ = C0()
                nw188_.ctor__(p1, (d_1037_v59_).issubset(d_1037_v59_))
                d_1038_v60_ = nw188_
                d_1039_v62_: _dafny.Set
                def iife88_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in (d_1009_v36_).Elements:
                        d_1040_v61_: int = compr_26_
                        if (d_1040_v61_) in (d_1009_v36_):
                            coll26_[default__.safeModuloInt(d_1040_v61_, d_964_v1_)] = d_964_v1_
                    return _dafny.Map(coll26_)
                d_1039_v62_ = _dafny.Set({iife88_()
                })
                d_973_v8_ = (d_973_v8_).set((d_1039_v62_).issubset(d_1039_v62_), self.f7)
                (self).f7 = (d_1038_v60_).f12
                default__.m0(len(d_968_v4_), (0) - (default__.safeModuloInt((0) - (d_964_v1_), d_964_v1_)), globalState)
                d_964_v1_ = (0) - (((d_968_v4_)[(d_963_v0_)[default__.safeIndex(d_964_v1_, len(d_963_v0_))]] if ((d_963_v0_)[default__.safeIndex(d_964_v1_, len(d_963_v0_))]) in (d_968_v4_) else d_964_v1_))
        d_1041_v63_: _dafny.Array
        nw189_ = _dafny.Array(False, 28)
        d_1041_v63_ = nw189_
        index172_ = default__.safeIndex(334, (d_1041_v63_).length(0))
        (d_1041_v63_)[index172_] = not(p1)
        index173_ = default__.safeIndex(334, (d_1041_v63_).length(0))
        (d_1041_v63_)[index173_] = p0

    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16

class C3(T1, T0):
    def  __init__(self):
        self._f7: bool = False
        self._f14: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def ctor__(self, f14, f7):
        (self)._f14 = f14
        (self).f7 = f7

    def fm8(self, p0, globalState):
        return self.f7

    def fm9(self, p0, p1, p2, globalState):
        return (self).f14

    def m4(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        d_1042_v0_: C0
        nw190_ = C0()
        nw190_.ctor__(self.f7, not (self.f7) or (not(self.f7)))
        d_1042_v0_ = nw190_
        if not(self.f7):
            (self).f7 = (d_1042_v0_).f12
            (d_1042_v0_).f13 = (d_1042_v0_).f12
            d_1043_v1_: _dafny.MultiSet
            d_1043_v1_ = _dafny.MultiSet([D3_DC11(273, p0, 187, False), D3_DC11(159, (self).f14, (self).f14, self.f7)])
            d_1044_v2_: D3
            d_1044_v2_ = D3_DC11((self).f14, p0, (self).f14, d_1042_v0_.f13)
            d_1043_v1_ = _dafny.MultiSet([d_1044_v2_, d_1044_v2_])
            d_1045_v3_: _dafny.Set
            d_1045_v3_ = _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([(self).f14])))})
            d_1046_v4_: C0
            nw191_ = C0()
            nw191_.ctor__(self.f7, (d_1045_v3_).ispropersubset(d_1045_v3_))
            d_1046_v4_ = nw191_
            d_1047_v5_: _dafny.Set
            d_1047_v5_ = _dafny.Set({d_1042_v0_.f13})
            d_1048_v6_: _dafny.Map
            d_1048_v6_ = _dafny.Map({(d_1042_v0_).f12: len(d_1047_v5_)})
            d_1049_v7_: _dafny.Map
            d_1049_v7_ = d_1048_v6_
            d_1050_v8_: _dafny.Set
            d_1050_v8_ = _dafny.Set({d_1048_v6_, _dafny.Map({d_1042_v0_.f13: (self).f14})})
            d_1051_v9_: _dafny.MultiSet
            d_1051_v9_ = _dafny.MultiSet([True, self.f7])
            d_1052_v10_: _dafny.Seq
            d_1052_v10_ = _dafny.SeqWithoutIsStrInference([False, (d_1050_v8_).issubset(_dafny.Set({((d_1049_v7_)).set((d_1042_v0_).f12, 949), d_1048_v6_})), (d_1046_v4_).f12, ((d_1042_v0_).f12) in (d_1051_v9_)])
            d_1052_v10_ = d_1052_v10_
        elif True:
            if self.f7:
                d_1053_v11_: _dafny.Seq
                d_1053_v11_ = _dafny.SeqWithoutIsStrInference([d_1042_v0_.f13])
                d_1054_v12_: str
                d_1054_v12_ = _dafny.CodePoint('t')
                r1 = (self).fm9(d_1042_v0_.f13, ((self).f14) - (len((d_1053_v11_).set(default__.safeIndex(p0, len(d_1053_v11_)), d_1042_v0_.f13))), d_1054_v12_, globalState)
                (d_1042_v0_).f13 = not(not(d_1042_v0_.f13))
                d_1055_v13_: _dafny.Array
                nw192_ = _dafny.Array(False, 10)
                d_1055_v13_ = nw192_
                d_1056_v14_: D3
                d_1056_v14_ = D3_DC10(d_1055_v13_)
                d_1057_v15_: D3
                d_1057_v15_ = D3_DC12(d_1056_v14_)
                rhs164_ = p0
                rhs165_ = d_1057_v15_
                r1 = rhs164_
                d_1057_v15_ = rhs165_
                d_1058_v16_: _dafny.MultiSet
                d_1058_v16_ = _dafny.MultiSet([d_1042_v0_, d_1042_v0_])
                d_1059_v17_: _dafny.Map
                d_1059_v17_ = _dafny.Map({(d_1058_v16_).cardinality: (d_1042_v0_).f12})
                d_1060_v18_: C0
                nw193_ = C0()
                nw193_.ctor__(((d_1059_v17_)[p0] if (p0) in (d_1059_v17_) else d_1042_v0_.f13), not(not((d_1042_v0_).f12)))
                d_1060_v18_ = nw193_
                d_1061_v19_: _dafny.Array
                def lambda71_(d_1062_p0_):
                    def lambda72_(d_1063_i0_):
                        return default__.safeDivisionInt(d_1063_i0_, d_1062_p0_)

                    return lambda72_

                init38_ = lambda71_(p0)
                nw194_ = _dafny.Array(None, 20)
                for i0_38_ in range(nw194_.length(0)):
                    nw194_[i0_38_] = init38_(i0_38_)
                d_1061_v19_ = nw194_
                index174_ = default__.safeIndex(917, (d_1061_v19_).length(0))
                (d_1061_v19_)[index174_] = ((self).f14) * (p0)
                index175_ = default__.safeIndex(917, (d_1061_v19_).length(0))
                rhs166_ = p0
                rhs167_ = ((0) - (p0)) + (-334)
                rhs168_ = d_1060_v18_
                lhs104_ = d_1061_v19_
                lhs105_ = default__.safeIndex(917, (d_1061_v19_).length(0))
                lhs104_[lhs105_] = rhs166_
                r3 = rhs167_
                d_1060_v18_ = rhs168_
            elif True:
                d_1064_v20_: _dafny.Array
                nw195_ = _dafny.Array(int(0), 20)
                d_1064_v20_ = nw195_
                index176_ = default__.safeIndex(456, (d_1064_v20_).length(0))
                (d_1064_v20_)[index176_] = p0
                d_1065_v21_: str
                d_1065_v21_ = _dafny.CodePoint('e')
                index177_ = default__.safeIndex(456, (d_1064_v20_).length(0))
                rhs169_ = (self).fm9(d_1042_v0_.f13, (self).f14, (D6_DC17((d_1042_v0_).f12, p0, d_1065_v21_, d_1065_v21_)).cf33, globalState)
                rhs170_ = self.f7
                rhs171_ = p0
                rhs172_ = 811
                lhs106_ = d_1064_v20_
                lhs107_ = default__.safeIndex(456, (d_1064_v20_).length(0))
                lhs108_ = d_1042_v0_
                lhs106_[lhs107_] = rhs169_
                lhs108_.f13 = rhs170_
                r3 = rhs171_
                r1 = rhs172_
                d_1066_v22_: _dafny.Seq
                d_1066_v22_ = _dafny.SeqWithoutIsStrInference([(self).f14, p0, p0])
                d_1067_v23_: _dafny.MultiSet
                d_1067_v23_ = _dafny.MultiSet([d_1066_v22_, default__.fm21(d_1042_v0_.f13, globalState)])
                d_1067_v23_ = _dafny.MultiSet([d_1066_v22_])
                d_1068_v24_: _dafny.Seq
                d_1068_v24_ = _dafny.SeqWithoutIsStrInference([(d_1042_v0_).fm16(d_1042_v0_.f13, globalState)])
                d_1069_v25_: C0
                nw196_ = C0()
                nw196_.ctor__((d_1068_v24_)[default__.safeIndex((d_1064_v20_)[default__.safeIndex(456, (d_1064_v20_).length(0))], len(d_1068_v24_))], d_1042_v0_.f13)
                d_1069_v25_ = nw196_
                d_1070_v26_: _dafny.Array
                nw197_ = _dafny.Array(_dafny.Seq({}), 19)
                d_1070_v26_ = nw197_
                d_1070_v26_ = d_1070_v26_
                d_1071_v27_: _dafny.Map
                d_1071_v27_ = _dafny.Map({769: self.f7})
                d_1072_v28_: _dafny.MultiSet
                d_1072_v28_ = _dafny.MultiSet([False])
                rhs173_ = d_1065_v21_
                rhs174_ = default__.safeModuloInt((d_1072_v28_).cardinality, p0)
                rhs175_ = _dafny.Map({(d_1066_v22_)[default__.safeIndex(len(_dafny.Set({(d_1064_v20_)[default__.safeIndex(456, (d_1064_v20_).length(0))]})), len(d_1066_v22_))]: d_1042_v0_.f13})
                d_1065_v21_ = rhs173_
                r3 = rhs174_
                d_1071_v27_ = rhs175_
            d_1073_v29_: _dafny.Array
            def lambda73_(d_1074_v0_):
                def lambda74_(d_1075_i1_):
                    return (d_1074_v0_).f12

                return lambda74_

            init39_ = lambda73_(d_1042_v0_)
            nw198_ = _dafny.Array(None, 19)
            for i0_39_ in range(nw198_.length(0)):
                nw198_[i0_39_] = init39_(i0_39_)
            d_1073_v29_ = nw198_
            d_1076_v30_: _dafny.Seq
            d_1076_v30_ = _dafny.SeqWithoutIsStrInference([d_1073_v29_])
            d_1076_v30_ = d_1076_v30_
            d_1077_v31_: str
            d_1077_v31_ = _dafny.CodePoint('x')
            d_1077_v31_ = d_1077_v31_
            if (d_1042_v0_).f12:
                d_1078_v32_: _dafny.Array
                nw199_ = _dafny.Array(None, 12)
                nw199_[int(0)] = ((self).f14) + (-954)
                nw199_[int(1)] = len((_dafny.SeqWithoutIsStrInference([d_1077_v31_ for d_1079_i2_ in range(default__.abs(836))]) if self.f7 else p1))
                nw199_[int(2)] = p0
                nw199_[int(3)] = default__.safeDivisionInt(p0, 919)
                nw199_[int(4)] = 566
                nw199_[int(5)] = (0) - (p0)
                nw199_[int(6)] = (self).f14
                nw199_[int(7)] = p0
                nw199_[int(8)] = (self).f14
                nw199_[int(9)] = (self).f14
                nw199_[int(10)] = p0
                nw199_[int(11)] = (self).f14
                d_1078_v32_ = nw199_
                index178_ = default__.safeIndex(436, (d_1078_v32_).length(0))
                (d_1078_v32_)[index178_] = len(p1)
                d_1080_v33_: _dafny.Seq
                d_1080_v33_ = _dafny.SeqWithoutIsStrInference([self.f7, not((d_1042_v0_).f12)])
                index179_ = default__.safeIndex(998, (d_1078_v32_).length(0))
                (d_1078_v32_)[index179_] = len((_dafny.SeqWithoutIsStrInference([(d_1042_v0_).f12])) + (d_1080_v33_))
                index180_ = default__.safeIndex(478, (d_1073_v29_).length(0))
                (d_1073_v29_)[index180_] = (d_1042_v0_).f12
                index181_ = default__.safeIndex(436, (d_1078_v32_).length(0))
                index182_ = default__.safeIndex(998, (d_1078_v32_).length(0))
                index183_ = default__.safeIndex(478, (d_1073_v29_).length(0))
                rhs176_ = p0
                rhs177_ = (0) - (p0)
                rhs178_ = d_1042_v0_.f13
                rhs179_ = d_1077_v31_
                lhs109_ = d_1078_v32_
                lhs110_ = default__.safeIndex(436, (d_1078_v32_).length(0))
                lhs111_ = d_1078_v32_
                lhs112_ = default__.safeIndex(998, (d_1078_v32_).length(0))
                lhs113_ = d_1073_v29_
                lhs114_ = default__.safeIndex(478, (d_1073_v29_).length(0))
                lhs109_[lhs110_] = rhs176_
                lhs111_[lhs112_] = rhs177_
                lhs113_[lhs114_] = rhs178_
                d_1077_v31_ = rhs179_
                d_1081_v34_: _dafny.Seq
                d_1081_v34_ = _dafny.SeqWithoutIsStrInference([(d_1078_v32_)[default__.safeIndex(436, (d_1078_v32_).length(0))]])
                index184_ = default__.safeIndex(436, (d_1078_v32_).length(0))
                (d_1078_v32_)[index184_] = (len(_dafny.Map({d_1042_v0_.f13: 33}))) * ((d_1081_v34_)[default__.safeIndex(p0, len(d_1081_v34_))])
                d_1082_v35_: _dafny.Seq
                d_1082_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wloks"))
                d_1083_v36_: _dafny.Array
                nw200_ = _dafny.Array(_dafny.MultiSet({}), 17)
                d_1083_v36_ = nw200_
                rhs180_ = p1
                rhs181_ = d_1083_v36_
                d_1082_v35_ = rhs180_
                d_1083_v36_ = rhs181_
                r3 = (len(d_1082_v35_)) * ((self).f14)
                r1 = 626
            elif True:
                r3 = p0
                (d_1042_v0_).f13 = not(d_1042_v0_.f13)
                r3 = (0) - (((self).f14) + ((self).f14))
                d_1084_v37_: _dafny.Seq
                d_1084_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyjc"))
                d_1084_v37_ = d_1084_v37_
                d_1085_v38_: _dafny.Array
                nw201_ = _dafny.Array(int(0), 10)
                d_1085_v38_ = nw201_
                index185_ = default__.safeIndex(568, (d_1085_v38_).length(0))
                (d_1085_v38_)[index185_] = p0
                index186_ = default__.safeIndex(568, (d_1085_v38_).length(0))
                (d_1085_v38_)[index186_] = p0
            d_1086_v39_: bool
            d_1087_v40_: bool
            out0_: bool
            out1_: bool
            out0_, out1_ = (self).m11(globalState)
            d_1086_v39_ = out0_
            d_1087_v40_ = out1_
        d_1088_v41_: _dafny.Seq
        d_1088_v41_ = _dafny.SeqWithoutIsStrInference([self.f7, d_1042_v0_.f13])
        d_1089_v42_: _dafny.Set
        d_1089_v42_ = _dafny.Set({(d_1042_v0_).f12})
        d_1090_v43_: _dafny.Set
        d_1090_v43_ = _dafny.Set({len(d_1089_v42_)})
        hi3_ = len(((d_1088_v41_).set(default__.safeIndex(len(d_1090_v43_), len(d_1088_v41_)), d_1042_v0_.f13)).set(default__.safeIndex(len(_dafny.Map({(self).f14: (d_1042_v0_).f12})), len((d_1088_v41_).set(default__.safeIndex(len(d_1090_v43_), len(d_1088_v41_)), d_1042_v0_.f13))), (d_1042_v0_).f12))
        for d_1091_i3_ in range(p0, hi3_):
            (d_1042_v0_).f13 = d_1042_v0_.f13
            r1 = (p0) - ((self).f14)
            d_1092_v44_: C0
            nw202_ = C0()
            nw202_.ctor__(self.f7, ((self).f14) >= (d_1091_i3_))
            d_1092_v44_ = nw202_
            d_1093_v45_: D3
            d_1093_v45_ = D3_DC11(d_1091_i3_, (self).f14, p0, (d_1042_v0_).f12)
            nw203_ = C0()
            nw203_.ctor__(self.f7, (d_1093_v45_).cf25)
            d_1092_v44_ = nw203_
            d_1094_v46_: _dafny.Array
            def lambda75_(d_1095_v42_, d_1096_v44_):
                def lambda76_(d_1097_i4_):
                    return (_dafny.Set({(d_1096_v44_).f12, d_1096_v44_.f13})).ispropersubset(d_1095_v42_)

                return lambda76_

            init40_ = lambda75_(d_1089_v42_, d_1092_v44_)
            nw204_ = _dafny.Array(None, 25)
            for i0_40_ in range(nw204_.length(0)):
                nw204_[i0_40_] = init40_(i0_40_)
            d_1094_v46_ = nw204_
            index187_ = default__.safeIndex(360, (d_1094_v46_).length(0))
            (d_1094_v46_)[index187_] = d_1042_v0_.f13
            index188_ = default__.safeIndex(360, (d_1094_v46_).length(0))
            rhs182_ = d_1091_i3_
            rhs183_ = not (d_1042_v0_.f13) or ((d_1092_v44_).f12)
            lhs115_ = d_1094_v46_
            lhs116_ = default__.safeIndex(360, (d_1094_v46_).length(0))
            r1 = rhs182_
            lhs115_[lhs116_] = rhs183_
        r3 = len((p1) + (p1))
        d_1098_v47_: _dafny.Array
        nw205_ = _dafny.Array(False, 1)
        d_1098_v47_ = nw205_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1098_v47_).length(0)):
            d_1099_i5_: int = guard_loop_1_
            if (True) and (((0) <= (d_1099_i5_)) and ((d_1099_i5_) < ((d_1098_v47_).length(0)))):
                (d_1098_v47_)[(d_1099_i5_)] = d_1042_v0_.f13
        d_1100_v48_: _dafny.MultiSet
        d_1100_v48_ = _dafny.MultiSet([(self).f14])
        if ((0) - (((d_1100_v48_)[len(p1)] if (len(p1)) in (d_1100_v48_) else (self).f14))) not in (d_1100_v48_):
            d_1101_v49_: _dafny.Map
            d_1101_v49_ = _dafny.Map({(self).fm8(869, globalState): d_1042_v0_.f13})
            d_1102_v50_: _dafny.Array
            nw206_ = _dafny.Array(None, 20)
            nw206_[int(0)] = _dafny.Map({d_1042_v0_.f13: (d_1042_v0_).f12})
            nw206_[int(1)] = d_1101_v49_
            nw206_[int(2)] = d_1101_v49_
            nw206_[int(3)] = d_1101_v49_
            nw206_[int(4)] = (d_1101_v49_).set((d_1042_v0_).f12, d_1042_v0_.f13)
            nw206_[int(5)] = d_1101_v49_
            nw206_[int(6)] = d_1101_v49_
            nw206_[int(7)] = d_1101_v49_
            nw206_[int(8)] = _dafny.Map({not(self.f7): d_1042_v0_.f13})
            nw206_[int(9)] = d_1101_v49_
            nw206_[int(10)] = d_1101_v49_
            nw206_[int(11)] = d_1101_v49_
            nw206_[int(12)] = d_1101_v49_
            nw206_[int(13)] = d_1101_v49_
            nw206_[int(14)] = d_1101_v49_
            nw206_[int(15)] = d_1101_v49_
            nw206_[int(16)] = d_1101_v49_
            nw206_[int(17)] = d_1101_v49_
            nw206_[int(18)] = d_1101_v49_
            nw206_[int(19)] = d_1101_v49_
            d_1102_v50_ = nw206_
            d_1103_v51_: _dafny.Array
            d_1103_v51_ = d_1102_v50_
            d_1104_v52_: _dafny.Array
            nw207_ = _dafny.Array(None, 4)
            nw207_[int(0)] = d_1103_v51_
            nw207_[int(1)] = d_1102_v50_
            nw207_[int(2)] = d_1102_v50_
            nw207_[int(3)] = d_1103_v51_
            d_1104_v52_ = nw207_
            d_1104_v52_ = d_1104_v52_
            r1 = default__.fm0(globalState)
            d_1105_v53_: str
            d_1105_v53_ = _dafny.CodePoint('b')
            r2 = (d_1105_v53_) in (_dafny.SeqWithoutIsStrInference([d_1105_v53_ for d_1106_i6_ in range(default__.abs(748))]))
            if ((d_1042_v0_).f12) not in (d_1101_v49_):
                r1 = (self).f14
                d_1107_v54_: _dafny.Map
                d_1107_v54_ = _dafny.Map({(self).f14: d_1098_v47_})
                d_1108_v55_: _dafny.Map
                d_1108_v55_ = _dafny.Map({len(d_1107_v54_): (0) - (((_dafny.MultiSet([len(p1), p0, len(d_1090_v43_)])).set(len(_dafny.SeqWithoutIsStrInference([(self).f14])), default__.abs(len(d_1090_v43_)))).cardinality)})
                d_1109_v56_: _dafny.Seq
                d_1109_v56_ = _dafny.SeqWithoutIsStrInference([d_1108_v55_])
                d_1109_v56_ = d_1109_v56_
                d_1108_v55_ = (d_1108_v55_).set((self).f14, 710)
                d_1110_v57_: _dafny.Seq
                d_1110_v57_ = _dafny.SeqWithoutIsStrInference([39, default__.fm0(globalState), len(p1), (self).f14])
                d_1111_v58_: C0
                nw208_ = C0()
                nw208_.ctor__((d_1042_v0_).f12, (p0) in (d_1110_v57_))
                d_1111_v58_ = nw208_
                r1 = (self).f14
            elif True:
                r2 = (d_1042_v0_).f12
                d_1105_v53_ = d_1105_v53_
                d_1112_v59_: _dafny.Seq
                d_1112_v59_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1113_v60_: D1
                d_1113_v60_ = D1_DC5(self.f7, (d_1112_v59_).set(default__.safeIndex((self).f14, len(d_1112_v59_)), p0), len(d_1112_v59_))
                d_1114_v61_: _dafny.MultiSet
                d_1114_v61_ = _dafny.MultiSet([(d_1042_v0_).f12])
                d_1115_v62_: _dafny.Seq
                d_1115_v62_ = _dafny.SeqWithoutIsStrInference([(d_1114_v61_).set((d_1042_v0_).f12, default__.abs(p0))])
                d_1116_v64_: C0
                nw209_ = C0()
                def iife89_():
                    coll27_ = _dafny.Set()
                    compr_27_: _dafny.MultiSet
                    for compr_27_ in (d_1115_v62_).Elements:
                        d_1117_v63_: _dafny.MultiSet = compr_27_
                        if (d_1117_v63_) in (d_1115_v62_):
                            coll27_ = coll27_.union(_dafny.Set([d_1117_v63_]))
                    return _dafny.Set(coll27_)
                nw209_.ctor__((d_1042_v0_).f12, (default__.fm22(globalState)).issubset(iife89_()
                ))
                d_1116_v64_ = nw209_
                rhs184_ = (d_1116_v64_).fm16((d_1042_v0_).f12, globalState)
                rhs185_ = default__.fm23(globalState)
                rhs186_ = (d_1090_v43_) | (d_1090_v43_)
                rhs187_ = d_1042_v0_
                lhs117_ = d_1042_v0_
                lhs117_.f13 = rhs184_
                d_1113_v60_ = rhs185_
                d_1090_v43_ = rhs186_
                d_1116_v64_ = rhs187_
                (self).f7 = (d_1112_v59_) == (d_1112_v59_)
                (d_1042_v0_).f13 = (d_1088_v41_)[default__.safeIndex((self).f14, len(d_1088_v41_))]
            d_1118_v65_: _dafny.Seq
            d_1118_v65_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1119_v66_: _dafny.Seq
            d_1119_v66_ = _dafny.SeqWithoutIsStrInference([p0, (self).f14])
            d_1120_v67_: D3
            d_1120_v67_ = D3_DC11(len((d_1118_v65_ if (d_1042_v0_).f12 else d_1118_v65_)), (len(d_1119_v66_)) * (p0), p0, (d_1042_v0_).f12)
            source7_ = d_1120_v67_
            if source7_.is_DC11:
                d_1121___mcc_h0_ = source7_.cf22
                d_1122___mcc_h1_ = source7_.cf23
                d_1123___mcc_h2_ = source7_.cf24
                d_1124___mcc_h3_ = source7_.cf25
                d_1125_cf25_ = d_1124___mcc_h3_
                d_1126_cf24_ = d_1123___mcc_h2_
                d_1127_cf23_ = d_1122___mcc_h1_
                d_1128_cf22_ = d_1121___mcc_h0_
                (self).f7 = d_1042_v0_.f13
                index189_ = default__.safeIndex(453, (d_1098_v47_).length(0))
                (d_1098_v47_)[index189_] = (d_1042_v0_).f12
                index190_ = default__.safeIndex(453, (d_1098_v47_).length(0))
                (d_1098_v47_)[index190_] = (d_1042_v0_).f12
                d_1127_cf23_ = (0) - ((d_1119_v66_)[default__.safeIndex(p0, len(d_1119_v66_))])
                d_1118_v65_ = _dafny.SeqWithoutIsStrInference([p1 for d_1129_i7_ in range(default__.abs(471))])
            elif source7_.is_DC10:
                d_1130___mcc_h4_ = source7_.cf21
                d_1131_cf21_ = d_1130___mcc_h4_
                (self).f7 = (self).fm8((self).f14, globalState)
                (d_1042_v0_).f13 = (d_1042_v0_).f12
                d_1132_v68_: _dafny.Array
                def lambda77_(d_1133_v53_):
                    def lambda78_(d_1134_i8_):
                        return d_1133_v53_

                    return lambda78_

                init41_ = lambda77_(d_1105_v53_)
                nw210_ = _dafny.Array(None, 13)
                for i0_41_ in range(nw210_.length(0)):
                    nw210_[i0_41_] = init41_(i0_41_)
                d_1132_v68_ = nw210_
                d_1135_v69_: _dafny.MultiSet
                d_1135_v69_ = _dafny.MultiSet([d_1132_v68_])
                d_1135_v69_ = d_1135_v69_
                d_1136_v70_: _dafny.Map
                d_1136_v70_ = _dafny.Map({d_1100_v48_: (d_1042_v0_).f12})
                d_1136_v70_ = (d_1136_v70_).set((d_1100_v48_) - (d_1100_v48_), ((d_1042_v0_).f12 if (d_1042_v0_).f12 else d_1042_v0_.f13))
            elif True:
                d_1137___mcc_h5_ = source7_.cf26
                d_1138_cf26_ = d_1137___mcc_h5_
                d_1139_v71_: _dafny.Array
                nw211_ = _dafny.Array(_dafny.Map({}), 15)
                d_1139_v71_ = nw211_
                d_1140_v72_: _dafny.Map
                d_1140_v72_ = _dafny.Map({d_1042_v0_.f13: _dafny.MultiSet([d_1042_v0_.f13])})
                index191_ = default__.safeIndex(982, (d_1139_v71_).length(0))
                (d_1139_v71_)[index191_] = d_1140_v72_
                index192_ = default__.safeIndex(982, (d_1139_v71_).length(0))
                rhs188_ = d_1042_v0_.f13
                rhs189_ = (self).fm8(p0, globalState)
                rhs190_ = len(_dafny.SeqWithoutIsStrInference([(d_1100_v48_).issubset(_dafny.MultiSet([(self).f14]))]))
                rhs191_ = (d_1140_v72_) | (d_1140_v72_)
                lhs118_ = d_1042_v0_
                lhs119_ = self
                lhs120_ = d_1139_v71_
                lhs121_ = default__.safeIndex(982, (d_1139_v71_).length(0))
                lhs118_.f13 = rhs188_
                lhs119_.f7 = rhs189_
                r3 = rhs190_
                lhs120_[lhs121_] = rhs191_
                index193_ = default__.safeIndex(693, (d_1098_v47_).length(0))
                (d_1098_v47_)[index193_] = d_1042_v0_.f13
                d_1141_v73_: _dafny.Map
                d_1141_v73_ = _dafny.Map({d_1105_v53_: self.f7})
                index194_ = default__.safeIndex(693, (d_1098_v47_).length(0))
                (d_1098_v47_)[index194_] = (((d_1141_v73_)[d_1105_v53_] if (d_1105_v53_) in (d_1141_v73_) else (d_1042_v0_).f12) if (self).fm8((self).f14, globalState) else (d_1042_v0_).f12)
                (d_1042_v0_).f13 = self.f7
                d_1142_v74_: _dafny.Array
                def lambda79_(d_1143_i9_):
                    return _dafny.MultiSet([(self).f14])

                init42_ = lambda79_
                nw212_ = _dafny.Array(None, 13)
                for i0_42_ in range(nw212_.length(0)):
                    nw212_[i0_42_] = init42_(i0_42_)
                d_1142_v74_ = nw212_
                d_1142_v74_ = (d_1142_v74_)
        elif True:
            r3 = (self).f14
            d_1144_v75_: _dafny.Array
            nw213_ = _dafny.Array(int(0), 5)
            d_1144_v75_ = nw213_
            index195_ = default__.safeIndex(450, (d_1144_v75_).length(0))
            (d_1144_v75_)[index195_] = len((d_1042_v0_).fm17((self).f14, self.f7, globalState))
            index196_ = default__.safeIndex(450, (d_1144_v75_).length(0))
            (d_1144_v75_)[index196_] = p0
            d_1145_v76_: _dafny.Seq
            d_1145_v76_ = _dafny.SeqWithoutIsStrInference([p0, p0, (self).f14])
            d_1146_v77_: _dafny.Seq
            d_1146_v77_ = _dafny.SeqWithoutIsStrInference([(len(d_1145_v76_)) * ((self).f14)])
            index197_ = default__.safeIndex(450, (d_1144_v75_).length(0))
            (d_1144_v75_)[index197_] = (d_1145_v76_)[default__.safeIndex(p0, len(d_1145_v76_))]
            d_1147_v78_: _dafny.Array
            nw214_ = _dafny.Array(_dafny.MultiSet({}), 27)
            d_1147_v78_ = nw214_
            rhs192_ = d_1147_v78_
            rhs193_ = (self).f14
            rhs194_ = p0
            d_1147_v78_ = rhs192_
            r1 = rhs193_
            r3 = rhs194_
            d_1148_v79_: _dafny.Set
            d_1148_v79_ = _dafny.Set({d_1098_v47_})
            d_1149_v80_: str
            d_1149_v80_ = _dafny.CodePoint('v')
            d_1150_v81_: T0
            nw215_ = C2()
            nw215_.ctor__(d_1148_v79_, d_1149_v80_, (d_1042_v0_).f12)
            d_1150_v81_ = nw215_
            d_1151_v82_: _dafny.Map
            d_1151_v82_ = _dafny.Map({d_1150_v81_: (self).f14})
            d_1152_v83_: D3
            d_1152_v83_ = D3_DC11(p0, (self).f14, ((d_1151_v82_)[d_1150_v81_] if (d_1150_v81_) in (d_1151_v82_) else (self).f14), d_1150_v81_.f7)
            d_1152_v83_ = d_1152_v83_
        d_1153_v84_: D3
        d_1153_v84_ = D3_DC10(d_1098_v47_)
        d_1154_v85_: _dafny.Seq
        d_1154_v85_ = _dafny.SeqWithoutIsStrInference([d_1098_v47_])
        d_1155_v86_: _dafny.Seq
        d_1155_v86_ = _dafny.SeqWithoutIsStrInference([d_1098_v47_, (d_1154_v85_)[default__.safeIndex(-433, len(d_1154_v85_))]])
        d_1156_v87_: _dafny.Array
        nw216_ = _dafny.Array(None, 27)
        nw216_[int(0)] = d_1098_v47_
        nw216_[int(1)] = (d_1153_v84_).cf21
        nw216_[int(2)] = d_1098_v47_
        nw216_[int(3)] = d_1098_v47_
        nw216_[int(4)] = d_1098_v47_
        nw216_[int(5)] = d_1098_v47_
        nw216_[int(6)] = d_1098_v47_
        nw216_[int(7)] = d_1098_v47_
        nw216_[int(8)] = d_1098_v47_
        nw216_[int(9)] = d_1098_v47_
        nw216_[int(10)] = d_1098_v47_
        nw216_[int(11)] = d_1098_v47_
        nw216_[int(12)] = d_1098_v47_
        nw216_[int(13)] = d_1098_v47_
        nw216_[int(14)] = (d_1098_v47_ if (d_1042_v0_).f12 else d_1098_v47_)
        nw216_[int(15)] = d_1098_v47_
        nw216_[int(16)] = d_1098_v47_
        nw216_[int(17)] = d_1098_v47_
        nw216_[int(18)] = (d_1155_v86_)[default__.safeIndex(default__.fm0(globalState), len(d_1155_v86_))]
        nw216_[int(19)] = d_1098_v47_
        nw216_[int(20)] = d_1098_v47_
        nw216_[int(21)] = d_1098_v47_
        nw216_[int(22)] = d_1098_v47_
        nw216_[int(23)] = d_1098_v47_
        nw216_[int(24)] = d_1098_v47_
        nw216_[int(25)] = d_1098_v47_
        nw216_[int(26)] = d_1098_v47_
        d_1156_v87_ = nw216_
        r0 = d_1156_v87_
        r1 = (self).f14
        d_1157_v88_: D0
        d_1157_v88_ = D0_DC0(p0)
        pat_let_tv24_ = d_1100_v48_
        pat_let_tv25_ = p0
        pat_let_tv26_ = d_1090_v43_
        pat_let_tv27_ = d_1042_v0_
        pat_let_tv28_ = d_1042_v0_
        def lambda80_(source8_):
            if source8_.is_DC1:
                d_1158___mcc_h6_ = source8_.cf1
                d_1159_cf1_ = d_1158___mcc_h6_
                return ((D12_DC29(pat_let_tv24_)).cf48).issubset(_dafny.MultiSet([pat_let_tv25_, len(pat_let_tv26_)]))
            elif source8_.is_DC2:
                d_1160___mcc_h7_ = source8_.cf2
                d_1161___mcc_h8_ = source8_.cf3
                d_1162___mcc_h9_ = source8_.cf4
                d_1163___mcc_h10_ = source8_.cf5
                d_1164_cf5_ = d_1163___mcc_h10_
                d_1165_cf4_ = d_1162___mcc_h9_
                d_1166_cf3_ = d_1161___mcc_h8_
                d_1167_cf2_ = d_1160___mcc_h7_
                return (pat_let_tv27_).f12
            elif True:
                d_1168___mcc_h11_ = source8_.cf0
                d_1169_cf0_ = d_1168___mcc_h11_
                return not((pat_let_tv28_).f12)

        r2 = lambda80_(d_1157_v88_)
        r3 = (len(d_1088_v41_)) * (((self).f14) * (p0))
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Set = _dafny.Set({})
        d_1170_v0_: _dafny.Array
        def lambda81_(d_1171_i1_):
            return not(self.f7)

        init43_ = lambda81_
        nw217_ = _dafny.Array(None, 28)
        for i0_43_ in range(nw217_.length(0)):
            nw217_[i0_43_] = init43_(i0_43_)
        d_1170_v0_ = nw217_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1170_v0_).length(0)):
            d_1172_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1172_i0_)) and ((d_1172_i0_) < ((d_1170_v0_).length(0)))):
                (d_1170_v0_)[(d_1172_i0_)] = self.f7
        d_1173_v1_: bool
        d_1174_v2_: _dafny.MultiSet
        d_1175_v3_: bool
        d_1176_v4_: _dafny.MultiSet
        out2_: bool
        out3_: _dafny.MultiSet
        out4_: bool
        out5_: _dafny.MultiSet
        out2_, out3_, out4_, out5_ = (self).m12(globalState)
        d_1173_v1_ = out2_
        d_1174_v2_ = out3_
        d_1175_v3_ = out4_
        d_1176_v4_ = out5_
        d_1177_i2_: int
        d_1177_i2_ = 0
        with _dafny.label("5"):
            while (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1184_i3_ in range(default__.abs(344))]))) > (p1):
                with _dafny.c_label("5"):
                    if (d_1177_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_1177_i2_ = (d_1177_i2_) + (1)
                    d_1178_v5_: int
                    d_1178_v5_ = -13
                    d_1179_v6_: _dafny.Set
                    d_1179_v6_ = _dafny.Set({p2})
                    rhs195_ = ((d_1179_v6_).intersection(d_1179_v6_)) != ((d_1179_v6_).intersection(d_1179_v6_))
                    rhs196_ = (d_1178_v5_) + (679)
                    lhs122_ = self
                    lhs122_.f7 = rhs195_
                    d_1178_v5_ = rhs196_
                    r1 = d_1179_v6_
                    d_1180_v7_: _dafny.Seq
                    d_1180_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqvfuxln"))
                    d_1181_v8_: _dafny.Set
                    d_1181_v8_ = _dafny.Set({len(_dafny.Map({len(d_1180_v7_): p2}))})
                    d_1173_v1_ = ((d_1181_v8_).intersection(d_1181_v8_)).isdisjoint((d_1181_v8_) | (d_1181_v8_))
                    d_1182_v9_: str
                    d_1182_v9_ = _dafny.CodePoint('x')
                    d_1183_v10_: _dafny.Map
                    d_1183_v10_ = _dafny.Map({d_1178_v5_: d_1182_v9_})
                    d_1183_v10_ = (d_1183_v10_).set(d_1178_v5_, d_1182_v9_)
                    pass
            pass
        index198_ = default__.safeIndex(645, (d_1170_v0_).length(0))
        (d_1170_v0_)[index198_] = p0
        d_1185_v11_: _dafny.Map
        d_1185_v11_ = _dafny.Map({not(p0): False})
        index199_ = default__.safeIndex(645, (d_1170_v0_).length(0))
        (d_1170_v0_)[index199_] = (d_1173_v1_) not in (d_1185_v11_)
        d_1186_v12_: D3
        d_1186_v12_ = D3_DC10(d_1170_v0_)
        pat_let_tv29_ = d_1170_v0_
        def iife90_(_pat_let31_0):
            def iife91_(d_1187_dt__update__tmp_h0_):
                def iife92_(_pat_let32_0):
                    def iife93_(d_1188_dt__update_hcf21_h0_):
                        return D3_DC10(d_1188_dt__update_hcf21_h0_)
                    return iife93_(_pat_let32_0)
                return iife92_(pat_let_tv29_)
            return iife91_(_pat_let31_0)
        d_1186_v12_ = iife90_(d_1186_v12_)
        d_1189_v13_: _dafny.MultiSet
        d_1189_v13_ = _dafny.MultiSet([(self).f14, p1, p1])
        d_1190_v14_: int
        d_1190_v14_ = 484
        rhs197_ = d_1189_v13_
        rhs198_ = (p1) - (default__.safeDivisionInt((0) - ((0) - ((0) - ((self).f14))), d_1190_v14_))
        rhs199_ = (self).f14
        d_1189_v13_ = rhs197_
        d_1190_v14_ = rhs198_
        d_1190_v14_ = rhs199_
        d_1191_v15_: _dafny.Set
        d_1191_v15_ = _dafny.Set({d_1176_v4_, d_1174_v2_})
        r0 = (d_1185_v11_).set(p2, (d_1190_v14_) > (len(d_1191_v15_)))
        d_1192_v16_: _dafny.Seq
        d_1192_v16_ = _dafny.SeqWithoutIsStrInference([self.f7])
        d_1193_v17_: str
        d_1193_v17_ = _dafny.CodePoint('s')
        d_1194_v18_: _dafny.Set
        d_1194_v18_ = _dafny.Set({self.f7, True})
        r1 = (default__.fm27((d_1192_v16_)[default__.safeIndex((self).f14, len(d_1192_v16_))], d_1193_v17_, d_1175_v3_, p0, globalState)) - ((d_1194_v18_) - (d_1194_v18_))
        return r0, r1

    def m3(self, globalState):
        d_1195_v1_: _dafny.Map
        d_1195_v1_ = _dafny.Map({(0) - ((self).f14): (self).f14})
        d_1196_v2_: _dafny.Seq
        def iife94_():
            coll28_ = _dafny.Map()
            compr_28_: int
            for compr_28_ in (d_1195_v1_).keys.Elements:
                d_1197_v0_: int = compr_28_
                if (d_1197_v0_) in (d_1195_v1_):
                    coll28_[(d_1197_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exnsnbf"))))] = (self).f14
            return _dafny.Map(coll28_)
        d_1196_v2_ = _dafny.SeqWithoutIsStrInference([len(iife94_()
        )])
        if not((d_1196_v2_) < (d_1196_v2_)):
            d_1198_v3_: int
            d_1198_v3_ = 980
            d_1199_v4_: _dafny.Seq
            d_1199_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kd"))
            d_1198_v3_ = len(((d_1199_v4_) + (d_1199_v4_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aa"))))
            d_1200_v5_: _dafny.Map
            d_1200_v5_ = _dafny.Map({self.f7: d_1199_v4_})
            d_1198_v3_ = len(((d_1200_v5_ if self.f7 else _dafny.Map({self.f7: d_1199_v4_}))) | (_dafny.Map({self.f7: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1201_i0_ in range(default__.abs(760))])})))
            d_1198_v3_ = (d_1196_v2_)[default__.safeIndex(d_1198_v3_, len(d_1196_v2_))]
            d_1202_v6_: D1
            d_1202_v6_ = D1_DC6(D1_DC4(self.f7, self.f7))
            d_1203_v7_: _dafny.Seq
            d_1203_v7_ = _dafny.SeqWithoutIsStrInference([True])
            d_1204_v8_: _dafny.Seq
            d_1204_v8_ = _dafny.SeqWithoutIsStrInference([(d_1203_v7_).set(default__.safeIndex(d_1198_v3_, len(d_1203_v7_)), self.f7), (d_1203_v7_) + (d_1203_v7_)])
            rhs200_ = d_1202_v6_
            rhs201_ = (self.f7 if self.f7 else (d_1198_v3_) >= (d_1198_v3_))
            rhs202_ = (self).f14
            rhs203_ = _dafny.SeqWithoutIsStrInference([(d_1203_v7_) + (d_1203_v7_) for d_1205_i1_ in range(default__.abs(-714))])
            lhs123_ = self
            d_1202_v6_ = rhs200_
            lhs123_.f7 = rhs201_
            d_1198_v3_ = rhs202_
            d_1204_v8_ = rhs203_
            d_1206_v9_: _dafny.Map
            d_1206_v9_ = _dafny.Map({(self).f14: True})
            d_1207_v10_: _dafny.Map
            d_1207_v10_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))) <= (d_1199_v4_): ((d_1206_v9_)[(self).f14] if ((self).f14) in (d_1206_v9_) else self.f7)})
            d_1208_v11_: D2
            d_1208_v11_ = D2_DC8(self.f7, d_1199_v4_)
            (self).f7 = ((d_1207_v10_)[not (self.f7) or (self.f7)] if (not (self.f7) or (self.f7)) in (d_1207_v10_) else (d_1208_v11_).cf14)
        elif True:
            d_1209_v12_: _dafny.Seq
            d_1209_v12_ = _dafny.SeqWithoutIsStrInference([self.f7, self.f7, True, self.f7])
            if ((d_1209_v12_)[default__.safeIndex((self).f14, len(d_1209_v12_))]) or (self.f7):
                d_1210_v13_: D6
                d_1210_v13_ = D6_DC16(not((not(self.f7) if (d_1209_v12_)[default__.safeIndex(len(d_1196_v2_), len(d_1209_v12_))] else self.f7)))
                def iife95_(_pat_let33_0):
                    def iife96_(d_1211_dt__update__tmp_h0_):
                        def iife97_(_pat_let34_0):
                            def iife98_(d_1212_dt__update_hcf30_h0_):
                                return D6_DC16(d_1212_dt__update_hcf30_h0_)
                            return iife98_(_pat_let34_0)
                        return iife97_(self.f7)
                    return iife96_(_pat_let33_0)
                d_1210_v13_ = iife95_(d_1210_v13_)
                d_1213_v14_: int
                d_1213_v14_ = 305
                d_1214_v15_: _dafny.Array
                nw218_ = _dafny.Array(False, 11)
                d_1214_v15_ = nw218_
                d_1215_v16_: _dafny.Set
                d_1215_v16_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qasccg"))), (_dafny.MultiSet([d_1214_v15_, d_1214_v15_, d_1214_v15_])).cardinality, d_1213_v14_})
                d_1213_v14_ = (d_1213_v14_) * ((d_1213_v14_) - (len(d_1215_v16_)))
                d_1216_v17_: _dafny.Map
                d_1216_v17_ = _dafny.Map({self.f7: _dafny.CodePoint('d')})
                d_1217_v18_: str
                d_1217_v18_ = _dafny.CodePoint('g')
                d_1218_v19_: D0
                d_1218_v19_ = D0_DC0((self).f14)
                d_1219_v20_: _dafny.Seq
                d_1219_v20_ = _dafny.SeqWithoutIsStrInference([d_1218_v19_])
                d_1220_v21_: _dafny.MultiSet
                d_1220_v21_ = _dafny.MultiSet([self.f7, default__.fm1(d_1219_v20_, (self).f14, d_1213_v14_, globalState)])
                d_1221_v22_: _dafny.Map
                d_1221_v22_ = _dafny.Map({((d_1216_v17_)[self.f7] if (self.f7) in (d_1216_v17_) else d_1217_v18_): (d_1220_v21_).cardinality})
                d_1222_v23_: C0
                nw219_ = C0()
                nw219_.ctor__(not((len(d_1221_v22_)) not in (d_1195_v1_)), self.f7)
                d_1222_v23_ = nw219_
                d_1223_v24_: _dafny.Seq
                d_1223_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twgvgcqtp"))
                d_1224_v25_: _dafny.Array
                nw220_ = _dafny.Array(None, 2)
                nw220_[int(0)] = d_1213_v14_
                nw220_[int(1)] = (0) - (len(d_1223_v24_))
                d_1224_v25_ = nw220_
                d_1225_v26_: D9
                d_1225_v26_ = D9_DC23(d_1224_v25_)
                d_1226_v27_: _dafny.Map
                d_1226_v27_ = _dafny.Map({d_1225_v26_: d_1224_v25_})
                d_1227_v28_: _dafny.MultiSet
                d_1227_v28_ = _dafny.MultiSet([d_1226_v27_, (d_1226_v27_).set(d_1225_v26_, d_1224_v25_), d_1226_v27_])
                default__.m0(((d_1227_v28_)[d_1226_v27_] if (d_1226_v27_) in (d_1227_v28_) else (self).f14), ((self).f14) * (d_1213_v14_), globalState)
                d_1228_v29_: D3
                d_1228_v29_ = D3_DC10(d_1214_v15_)
                d_1229_v30_: D3
                d_1229_v30_ = D3_DC12(d_1228_v29_)
                pat_let_tv30_ = globalState
                pat_let_tv31_ = d_1213_v14_
                pat_let_tv32_ = d_1222_v23_
                d_1230_v31_: _dafny.MultiSet
                def iife99_(_pat_let35_0):
                    def iife100_(d_1231_dt__update__tmp_h1_):
                        def iife101_(_pat_let36_0):
                            def iife102_(d_1232_dt__update_hcf26_h0_):
                                return D3_DC12(d_1232_dt__update_hcf26_h0_)
                            return iife102_(_pat_let36_0)
                        return iife101_(D3_DC11(562, len(default__.fm46((self).f14, pat_let_tv30_)), pat_let_tv31_, (pat_let_tv32_).f12))
                    return iife100_(_pat_let35_0)
                d_1230_v31_ = _dafny.MultiSet([iife99_(d_1229_v30_)])
                d_1230_v31_ = d_1230_v31_
            elif True:
                d_1233_v32_: _dafny.Array
                def lambda82_(d_1234_i2_):
                    return (d_1234_i2_) + ((0) - ((self).f14))

                init44_ = lambda82_
                nw221_ = _dafny.Array(None, 9)
                for i0_44_ in range(nw221_.length(0)):
                    nw221_[i0_44_] = init44_(i0_44_)
                d_1233_v32_ = nw221_
                d_1235_v33_: _dafny.MultiSet
                d_1235_v33_ = _dafny.MultiSet([d_1233_v32_, d_1233_v32_, d_1233_v32_])
                index200_ = default__.safeIndex(504, (d_1233_v32_).length(0))
                (d_1233_v32_)[index200_] = (d_1235_v33_).cardinality
                index201_ = default__.safeIndex(504, (d_1233_v32_).length(0))
                (d_1233_v32_)[index201_] = len(d_1196_v2_)
                d_1236_v34_: str
                d_1236_v34_ = _dafny.CodePoint('b')
                d_1237_v35_: _dafny.Set
                d_1237_v35_ = _dafny.Set({d_1236_v34_})
                d_1238_v36_: _dafny.Map
                d_1238_v36_ = _dafny.Map({self.f7: d_1236_v34_})
                d_1239_v38_: _dafny.Map
                def iife103_():
                    coll29_ = _dafny.Set()
                    compr_29_: str
                    for compr_29_ in (_dafny.Set({d_1236_v34_})).Elements:
                        d_1240_v37_: str = compr_29_
                        if (d_1240_v37_) in (_dafny.Set({d_1236_v34_})):
                            coll29_ = coll29_.union(_dafny.Set([d_1240_v37_]))
                    return _dafny.Set(coll29_)
                d_1239_v38_ = _dafny.Map({d_1238_v36_: iife103_()
                })
                d_1241_v39_: _dafny.Seq
                d_1241_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1236_v34_: (self).f14})])
                d_1242_v41_: _dafny.Array
                nw222_ = _dafny.Array(None, 12)
                nw222_[int(0)] = _dafny.Set({_dafny.CodePoint('a'), d_1236_v34_, d_1236_v34_})
                nw222_[int(1)] = d_1237_v35_
                nw222_[int(2)] = _dafny.Set({d_1236_v34_, d_1236_v34_, d_1236_v34_})
                nw222_[int(3)] = (d_1237_v35_) - (d_1237_v35_)
                nw222_[int(4)] = d_1237_v35_
                nw222_[int(5)] = ((d_1239_v38_)[d_1238_v36_] if (d_1238_v36_) in (d_1239_v38_) else d_1237_v35_)
                nw222_[int(6)] = d_1237_v35_
                nw222_[int(7)] = (d_1237_v35_) | (d_1237_v35_)
                nw222_[int(8)] = (d_1237_v35_).intersection(d_1237_v35_)
                def iife104_():
                    coll30_ = _dafny.Set()
                    compr_30_: str
                    for compr_30_ in ((d_1241_v39_)[default__.safeIndex((0) - ((d_1233_v32_)[default__.safeIndex(504, (d_1233_v32_).length(0))]), len(d_1241_v39_))]).keys.Elements:
                        d_1243_v40_: str = compr_30_
                        if (d_1243_v40_) in ((d_1241_v39_)[default__.safeIndex((0) - ((d_1233_v32_)[default__.safeIndex(504, (d_1233_v32_).length(0))]), len(d_1241_v39_))]):
                            coll30_ = coll30_.union(_dafny.Set([d_1243_v40_]))
                    return _dafny.Set(coll30_)
                nw222_[int(9)] = iife104_()
                
                nw222_[int(10)] = (d_1237_v35_) | (d_1237_v35_)
                nw222_[int(11)] = d_1237_v35_
                d_1242_v41_ = nw222_
                index202_ = default__.safeIndex(500, (d_1242_v41_).length(0))
                (d_1242_v41_)[index202_] = d_1237_v35_
                index203_ = default__.safeIndex(500, (d_1242_v41_).length(0))
                (d_1242_v41_)[index203_] = ((d_1237_v35_) - (d_1237_v35_)).intersection(d_1237_v35_)
                d_1244_v42_: _dafny.Set
                d_1244_v42_ = _dafny.Set({(d_1233_v32_)[default__.safeIndex(504, (d_1233_v32_).length(0))], (self).f14})
                d_1245_v43_: _dafny.MultiSet
                d_1245_v43_ = _dafny.MultiSet([((self).f14) * ((self).f14), (0) - ((d_1233_v32_)[default__.safeIndex(504, (d_1233_v32_).length(0))]), default__.safeDivisionInt((self).f14, len(d_1244_v42_)), (d_1233_v32_)[default__.safeIndex(504, (d_1233_v32_).length(0))]])
                d_1245_v43_ = (default__.fm33(257, d_1236_v34_, globalState)) - (d_1245_v43_)
                d_1246_v44_: D0
                d_1246_v44_ = D0_DC0((self).f14)
                d_1247_v45_: _dafny.Seq
                d_1247_v45_ = _dafny.SeqWithoutIsStrInference([d_1246_v44_])
                d_1248_v46_: _dafny.Seq
                d_1248_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htrgmixse"))
                d_1249_v47_: _dafny.Seq
                d_1249_v47_ = _dafny.SeqWithoutIsStrInference([default__.fm21(default__.fm1(d_1247_v45_, len(d_1248_v46_), (d_1233_v32_)[default__.safeIndex(504, (d_1233_v32_).length(0))], globalState), globalState)])
                d_1196_v2_ = (d_1249_v47_)[default__.safeIndex(((d_1233_v32_)[default__.safeIndex(504, (d_1233_v32_).length(0))]) + ((d_1233_v32_)[default__.safeIndex(504, (d_1233_v32_).length(0))]), len(d_1249_v47_))]
                d_1250_v48_: _dafny.Array
                nw223_ = _dafny.Array(_dafny.Seq({}), 20)
                d_1250_v48_ = nw223_
                index204_ = default__.safeIndex(811, (d_1250_v48_).length(0))
                (d_1250_v48_)[index204_] = d_1196_v2_
                index205_ = default__.safeIndex(811, (d_1250_v48_).length(0))
                rhs204_ = d_1196_v2_
                rhs205_ = d_1248_v46_
                lhs124_ = d_1250_v48_
                lhs125_ = default__.safeIndex(811, (d_1250_v48_).length(0))
                lhs124_[lhs125_] = rhs204_
                d_1248_v46_ = rhs205_
            (self).f7 = True
            d_1251_v49_: _dafny.Set
            d_1251_v49_ = _dafny.Set({(self).f14})
            (self).f7 = ((d_1251_v49_) | (d_1251_v49_)).ispropersubset(d_1251_v49_)
            d_1252_v50_: _dafny.MultiSet
            d_1252_v50_ = _dafny.MultiSet([self.f7, self.f7, self.f7, True])
            if ((self).f14) >= (default__.safeDivisionInt((d_1252_v50_).cardinality, (self).f14)):
                (self).f7 = (D1_DC4(not(True), self.f7)).cf7
                (self).f7 = self.f7
                (self).f7 = self.f7
                d_1253_v51_: _dafny.Seq
                d_1253_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tu"))
                d_1254_v52_: _dafny.Map
                d_1254_v52_ = _dafny.Map({(d_1253_v51_) + (d_1253_v51_): self.f7})
                d_1254_v52_ = (d_1254_v52_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1255_i3_ in range(default__.abs(203))]), self.f7)
                (self).f7 = self.f7
            elif True:
                d_1256_v53_: _dafny.Seq
                d_1256_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsnf"))
                d_1257_v54_: _dafny.Array
                nw224_ = _dafny.Array(None, 21)
                nw224_[int(0)] = self.f7
                nw224_[int(1)] = self.f7
                nw224_[int(2)] = (self).fm8(default__.fm0(globalState), globalState)
                nw224_[int(3)] = self.f7
                nw224_[int(4)] = self.f7
                nw224_[int(5)] = self.f7
                nw224_[int(6)] = self.f7
                nw224_[int(7)] = self.f7
                nw224_[int(8)] = self.f7
                nw224_[int(9)] = self.f7
                nw224_[int(10)] = self.f7
                nw224_[int(11)] = self.f7
                nw224_[int(12)] = self.f7
                nw224_[int(13)] = self.f7
                nw224_[int(14)] = self.f7
                nw224_[int(15)] = self.f7
                nw224_[int(16)] = self.f7
                nw224_[int(17)] = self.f7
                nw224_[int(18)] = self.f7
                nw224_[int(19)] = (self).fm8(len(d_1256_v53_), globalState)
                nw224_[int(20)] = self.f7
                d_1257_v54_ = nw224_
                d_1258_v55_: _dafny.Set
                d_1258_v55_ = _dafny.Set({d_1257_v54_})
                d_1259_v56_: str
                d_1259_v56_ = _dafny.CodePoint('c')
                d_1260_v57_: C2
                nw225_ = C2()
                nw225_.ctor__(d_1258_v55_, d_1259_v56_, self.f7)
                d_1260_v57_ = nw225_
                d_1260_v57_ = d_1260_v57_
                d_1261_v58_: _dafny.Map
                d_1261_v58_ = _dafny.Map({(d_1252_v50_).issubset(d_1252_v50_): self.f7})
                d_1261_v58_ = (d_1261_v58_).set(not (self.f7) or (self.f7), (d_1256_v53_) == ((d_1256_v53_).set(default__.safeIndex(732, len(d_1256_v53_)), (d_1260_v57_).f16)))
                d_1262_v59_: int
                d_1262_v59_ = -744
                d_1262_v59_ = (self).f14
                default__.m0(d_1262_v59_, d_1262_v59_, globalState)
                d_1263_v60_: _dafny.Array
                nw226_ = _dafny.Array(_dafny.CodePoint('D'), 13)
                d_1263_v60_ = nw226_
                d_1264_v61_: _dafny.Set
                d_1264_v61_ = _dafny.Set({self.f7, self.f7})
                index206_ = default__.safeIndex(186, (d_1257_v54_).length(0))
                (d_1257_v54_)[index206_] = (d_1264_v61_).ispropersubset(d_1264_v61_)
                d_1265_v62_: _dafny.Seq
                d_1265_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acfbl")), _dafny.SeqWithoutIsStrInference([(d_1260_v57_).f16 for d_1266_i4_ in range(default__.abs(-775))]), d_1256_v53_])
                index207_ = default__.safeIndex(186, (d_1257_v54_).length(0))
                rhs206_ = ((d_1265_v62_)[default__.safeIndex(len(d_1251_v49_), len(d_1265_v62_))]) + (d_1256_v53_)
                rhs207_ = d_1263_v60_
                rhs208_ = self.f7
                rhs209_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acvrlugbv"))
                lhs126_ = d_1257_v54_
                lhs127_ = default__.safeIndex(186, (d_1257_v54_).length(0))
                d_1256_v53_ = rhs206_
                d_1263_v60_ = rhs207_
                lhs126_[lhs127_] = rhs208_
                d_1256_v53_ = rhs209_
            d_1267_v63_: _dafny.Set
            d_1267_v63_ = _dafny.Set({self.f7, self.f7, self.f7, self.f7, self.f7})
            d_1268_v64_: D1
            d_1268_v64_ = D1_DC3(d_1267_v63_)
            source9_ = d_1268_v64_
            if source9_.is_DC4:
                d_1269___mcc_h0_ = source9_.cf7
                d_1270___mcc_h1_ = source9_.cf8
                d_1271_cf8_ = d_1270___mcc_h1_
                d_1272_cf7_ = d_1269___mcc_h0_
                d_1273_v65_: D0
                d_1273_v65_ = D0_DC0((self).f14)
                d_1274_v66_: _dafny.Map
                d_1274_v66_ = _dafny.Map({d_1267_v63_: False})
                d_1275_v67_: _dafny.Map
                d_1275_v67_ = _dafny.Map({default__.fm1(_dafny.SeqWithoutIsStrInference([d_1273_v65_, d_1273_v65_]), (self).f14, 449, globalState): d_1274_v66_})
                d_1275_v67_ = (d_1275_v67_).set(self.f7, d_1274_v66_)
                d_1276_v68_: T0
                nw227_ = C1()
                nw227_.ctor__(d_1271_cf8_)
                d_1276_v68_ = nw227_
                d_1277_v69_: _dafny.Map
                d_1277_v69_ = _dafny.Map({d_1271_cf8_: d_1276_v68_})
                d_1278_v70_: str
                d_1278_v70_ = _dafny.CodePoint('h')
                d_1279_v71_: _dafny.MultiSet
                d_1279_v71_ = _dafny.MultiSet([((self).f14) + (len((d_1277_v69_).set(d_1272_cf7_, d_1276_v68_))), (self).f14, (0) - (len(_dafny.Map({d_1278_v70_: d_1276_v68_.f7})))])
                d_1280_v72_: _dafny.Map
                d_1280_v72_ = _dafny.Map({d_1272_cf7_: (self).f14})
                d_1279_v71_ = (((_dafny.MultiSet([(self).f14])).set((self).f14, default__.abs((self).f14)) if not(d_1272_cf7_) else d_1279_v71_)).intersection(_dafny.MultiSet([(self).f14, ((d_1280_v72_)[d_1271_cf8_] if (d_1271_cf8_) in (d_1280_v72_) else (self).f14)]))
                d_1281_v73_: _dafny.Seq
                d_1281_v73_ = _dafny.SeqWithoutIsStrInference([d_1278_v70_, d_1278_v70_])
                d_1282_v74_: _dafny.Map
                d_1282_v74_ = _dafny.Map({d_1196_v2_: (d_1281_v73_)[default__.safeIndex((self).f14, len(d_1281_v73_))]})
                d_1282_v74_ = d_1282_v74_
                (d_1276_v68_).f7 = not (d_1271_cf8_) or (d_1276_v68_.f7)
            elif source9_.is_DC5:
                d_1283___mcc_h2_ = source9_.cf9
                d_1284___mcc_h3_ = source9_.cf10
                d_1285___mcc_h4_ = source9_.cf11
                d_1286_cf11_ = d_1285___mcc_h4_
                d_1287_cf10_ = d_1284___mcc_h3_
                d_1288_cf9_ = d_1283___mcc_h2_
                d_1195_v1_ = (d_1195_v1_).set((self).f14, (d_1286_cf11_) + (d_1286_cf11_))
                d_1289_v75_: str
                d_1289_v75_ = _dafny.CodePoint('a')
                d_1290_v76_: D6
                d_1290_v76_ = D6_DC17(d_1288_cf9_, 579, _dafny.CodePoint('f'), d_1289_v75_)
                d_1291_v77_: D6
                d_1291_v77_ = D6_DC18(d_1290_v76_)
                d_1292_v78_: _dafny.Array
                def lambda83_(d_1293_i5_):
                    return False

                init45_ = lambda83_
                nw228_ = _dafny.Array(None, 19)
                for i0_45_ in range(nw228_.length(0)):
                    nw228_[i0_45_] = init45_(i0_45_)
                d_1292_v78_ = nw228_
                d_1294_v79_: D12
                d_1294_v79_ = D12_DC30(335, d_1288_cf9_, d_1196_v2_, (self).f14)
                pat_let_tv33_ = d_1196_v2_
                def iife105_(_pat_let37_0):
                    def iife106_(d_1295_dt__update__tmp_h2_):
                        def iife107_(_pat_let38_0):
                            def iife108_(d_1296_dt__update_hcf49_h0_):
                                return D12_DC30(d_1296_dt__update_hcf49_h0_, (d_1295_dt__update__tmp_h2_).cf50, (d_1295_dt__update__tmp_h2_).cf51, (d_1295_dt__update__tmp_h2_).cf52)
                            return iife108_(_pat_let38_0)
                        return iife107_(len(pat_let_tv33_))
                    return iife106_(_pat_let37_0)
                rhs210_ = default__.fm47((iife105_(d_1294_v79_)).cf51, default__.safeModuloInt(d_1286_cf11_, d_1286_cf11_), globalState)
                rhs211_ = d_1292_v78_
                rhs212_ = d_1292_v78_
                d_1291_v77_ = rhs210_
                d_1292_v78_ = rhs211_
                d_1292_v78_ = rhs212_
                d_1209_v12_ = d_1209_v12_
                d_1297_v80_: _dafny.Array
                def lambda84_(d_1298_i6_):
                    return (d_1298_i6_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdgr"))))

                init46_ = lambda84_
                nw229_ = _dafny.Array(None, 27)
                for i0_46_ in range(nw229_.length(0)):
                    nw229_[i0_46_] = init46_(i0_46_)
                d_1297_v80_ = nw229_
                d_1297_v80_ = d_1297_v80_
            elif source9_.is_DC3:
                d_1299___mcc_h5_ = source9_.cf6
                d_1300_cf6_ = d_1299___mcc_h5_
                d_1301_v81_: C1
                nw230_ = C1()
                nw230_.ctor__(self.f7)
                d_1301_v81_ = nw230_
                (self).f7 = self.f7
                d_1302_v82_: _dafny.Seq
                d_1302_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                d_1303_v83_: str
                d_1303_v83_ = _dafny.CodePoint('s')
                d_1304_v84_: _dafny.MultiSet
                d_1304_v84_ = _dafny.MultiSet([(self).f14])
                d_1305_v85_: D12
                d_1305_v85_ = D12_DC29(d_1304_v84_)
                d_1306_v86_: _dafny.Array
                nw231_ = _dafny.Array(int(0), 24)
                d_1306_v86_ = nw231_
                d_1307_v87_: _dafny.Map
                d_1307_v87_ = _dafny.Map({(self).f14: d_1306_v86_})
                d_1308_v88_: _dafny.Map
                d_1308_v88_ = _dafny.Map({self.f7: d_1306_v86_})
                d_1309_v89_: _dafny.MultiSet
                d_1309_v89_ = _dafny.MultiSet([((d_1307_v87_)[(0) - ((self).f14)] if ((0) - ((self).f14)) in (d_1307_v87_) else ((d_1308_v88_)[self.f7] if (self.f7) in (d_1308_v88_) else d_1306_v86_))])
                d_1310_v90_: _dafny.Array
                nw232_ = _dafny.Array(False, 2)
                d_1310_v90_ = nw232_
                d_1311_v91_: _dafny.Map
                d_1311_v91_ = _dafny.Map({d_1303_v83_: d_1310_v90_})
                d_1312_v92_: D0
                d_1312_v92_ = D0_DC0((self).f14)
                d_1313_v93_: _dafny.Map
                d_1313_v93_ = _dafny.Map({d_1306_v86_: (D2_DC9(-590, (self).f14, d_1302_v82_, self.f7, _dafny.Map({d_1312_v92_: (self).f14}))).cf16})
                d_1314_v94_: _dafny.Map
                d_1314_v94_ = _dafny.Map({d_1312_v92_: (0) - (len(d_1195_v1_))})
                d_1315_v95_: D2
                d_1315_v95_ = D2_DC9((self).f14, len(d_1209_v12_), d_1302_v82_, self.f7, d_1314_v94_)
                d_1316_v96_: _dafny.Seq
                def iife109_(_pat_let39_0):
                    def iife110_(d_1317_dt__update__tmp_h3_):
                        def iife111_(_pat_let40_0):
                            def iife112_(d_1318_dt__update_hcf17_h0_):
                                def iife113_(_pat_let41_0):
                                    def iife114_(d_1319_dt__update_hcf16_h0_):
                                        return D2_DC9(d_1319_dt__update_hcf16_h0_, d_1318_dt__update_hcf17_h0_, (d_1317_dt__update__tmp_h3_).cf18, (d_1317_dt__update__tmp_h3_).cf19, (d_1317_dt__update__tmp_h3_).cf20)
                                    return iife114_(_pat_let41_0)
                                return iife113_((self).f14)
                            return iife112_(_pat_let40_0)
                        return iife111_(-22)
                    return iife110_(_pat_let39_0)
                d_1316_v96_ = _dafny.SeqWithoutIsStrInference([iife109_(d_1315_v95_)])
                d_1320_v97_: _dafny.Array
                nw233_ = _dafny.Array(None, 27)
                nw233_[int(0)] = len((d_1302_v82_) + (d_1302_v82_))
                nw233_[int(1)] = (0) - (len(d_1196_v2_))
                nw233_[int(2)] = (self).fm9(self.f7, (_dafny.MultiSet([self.f7])).cardinality, d_1303_v83_, globalState)
                nw233_[int(3)] = ((self).f14) - ((d_1304_v84_).cardinality)
                nw233_[int(4)] = default__.fm0(globalState)
                nw233_[int(5)] = ((self).f14) - (((d_1305_v85_).cf48).cardinality)
                nw233_[int(6)] = (self).fm9(self.f7, (self).f14, d_1303_v83_, globalState)
                nw233_[int(7)] = (d_1309_v89_).cardinality
                nw233_[int(8)] = 638
                nw233_[int(9)] = len((d_1302_v82_).set(default__.safeIndex(len(d_1311_v91_), len(d_1302_v82_)), _dafny.CodePoint('l')))
                nw233_[int(10)] = len(_dafny.Set({self.f7}))
                nw233_[int(11)] = 483
                nw233_[int(12)] = (0) - ((self).f14)
                nw233_[int(13)] = ((self).f14) - (709)
                nw233_[int(14)] = default__.fm0(globalState)
                nw233_[int(15)] = ((d_1313_v93_)[d_1306_v86_] if (d_1306_v86_) in (d_1313_v93_) else 787)
                nw233_[int(16)] = (self).f14
                nw233_[int(17)] = (978 if self.f7 else (self).f14)
                nw233_[int(18)] = -411
                nw233_[int(19)] = (self).f14
                nw233_[int(20)] = len((d_1302_v82_) + ((d_1302_v82_).set(default__.safeIndex(len(d_1209_v12_), len(d_1302_v82_)), d_1303_v83_)))
                nw233_[int(21)] = (self).f14
                nw233_[int(22)] = default__.safeDivisionInt(len(d_1316_v96_), (self).f14)
                nw233_[int(23)] = ((self).f14 if self.f7 else (d_1252_v50_).cardinality)
                nw233_[int(24)] = (self).f14
                nw233_[int(25)] = -18
                nw233_[int(26)] = (self).f14
                d_1320_v97_ = nw233_
                index208_ = default__.safeIndex(600, (d_1306_v86_).length(0))
                (d_1306_v86_)[index208_] = (self).f14
                index209_ = default__.safeIndex(600, (d_1306_v86_).length(0))
                (d_1306_v86_)[index209_] = (0) - ((self).f14)
                d_1321_v98_: D0
                d_1321_v98_ = D0_DC1((self).f14)
                pat_let_tv34_ = d_1303_v83_
                def iife115_(_pat_let42_0):
                    def iife116_(d_1322_dt__update__tmp_h4_):
                        def iife117_(_pat_let43_0):
                            def iife118_(d_1324_dt__update_hcf1_h0_):
                                return D0_DC1(d_1324_dt__update_hcf1_h0_)
                            return iife118_(_pat_let43_0)
                        return iife117_((len(_dafny.SeqWithoutIsStrInference([pat_let_tv34_ for d_1323_i7_ in range(default__.abs(576))]))) * (725))
                    return iife116_(_pat_let42_0)
                d_1321_v98_ = iife115_(d_1321_v98_)
            elif True:
                d_1325___mcc_h6_ = source9_.cf12
                d_1326_cf12_ = d_1325___mcc_h6_
                d_1327_v99_: int
                d_1327_v99_ = 365
                d_1327_v99_ = (self).f14
                d_1327_v99_ = ((0) - (d_1327_v99_)) - ((self).f14)
                d_1328_v100_: _dafny.Array
                def lambda85_(d_1329_i8_):
                    return not(True)

                init47_ = lambda85_
                nw234_ = _dafny.Array(None, 25)
                for i0_47_ in range(nw234_.length(0)):
                    nw234_[i0_47_] = init47_(i0_47_)
                d_1328_v100_ = nw234_
                d_1330_v101_: _dafny.Set
                d_1330_v101_ = _dafny.Set({d_1328_v100_})
                d_1331_v102_: str
                d_1331_v102_ = _dafny.CodePoint('d')
                d_1332_v103_: _dafny.Seq
                d_1332_v103_ = _dafny.SeqWithoutIsStrInference([D0_DC0((self).f14)])
                d_1333_v104_: C2
                nw235_ = C2()
                nw235_.ctor__(d_1330_v101_, d_1331_v102_, (default__.fm1(d_1332_v103_, d_1327_v99_, d_1327_v99_, globalState)) == (self.f7))
                d_1333_v104_ = nw235_
                d_1334_v105_: _dafny.Array
                nw236_ = _dafny.Array(None, 2)
                nw236_[int(0)] = d_1327_v99_
                nw236_[int(1)] = d_1327_v99_
                d_1334_v105_ = nw236_
                d_1335_v106_: _dafny.Map
                d_1335_v106_ = _dafny.Map({self.f7: len(d_1195_v1_)})
                d_1336_v107_: _dafny.Map
                d_1336_v107_ = _dafny.Map({d_1327_v99_: self.f7})
                d_1337_v108_: _dafny.Map
                d_1337_v108_ = _dafny.Map({d_1335_v106_: (default__.fm48((0) - (d_1327_v99_), d_1327_v99_, ((d_1336_v107_)[d_1327_v99_] if (d_1327_v99_) in (d_1336_v107_) else self.f7), self.f7, globalState)).cf48})
                d_1338_v109_: _dafny.MultiSet
                d_1338_v109_ = _dafny.MultiSet([len(default__.fm2(516, True, globalState))])
                index210_ = default__.safeIndex(202, (d_1334_v105_).length(0))
                (d_1334_v105_)[index210_] = (((d_1337_v108_)[d_1335_v106_] if (d_1335_v106_) in (d_1337_v108_) else d_1338_v109_)).cardinality
                index211_ = default__.safeIndex(202, (d_1334_v105_).length(0))
                (d_1334_v105_)[index211_] = len(_dafny.SeqWithoutIsStrInference([d_1334_v105_, (d_1334_v105_ if False else d_1334_v105_)]))
        d_1339_v110_: C0
        nw237_ = C0()
        nw237_.ctor__(self.f7, self.f7)
        d_1339_v110_ = nw237_
        d_1340_v111_: _dafny.Seq
        d_1340_v111_ = _dafny.SeqWithoutIsStrInference([(self).fm8((self).f14, globalState)])
        d_1341_v112_: _dafny.Array
        nw238_ = _dafny.Array(None, 20)
        nw238_[int(0)] = (d_1339_v110_).f12
        nw238_[int(1)] = self.f7
        nw238_[int(2)] = d_1339_v110_.f13
        nw238_[int(3)] = (d_1340_v111_)[default__.safeIndex(((d_1195_v1_)[(self).f14] if ((self).f14) in (d_1195_v1_) else (self).f14), len(d_1340_v111_))]
        nw238_[int(4)] = self.f7
        nw238_[int(5)] = (d_1339_v110_).f12
        nw238_[int(6)] = d_1339_v110_.f13
        nw238_[int(7)] = not((d_1339_v110_).fm16(self.f7, globalState))
        nw238_[int(8)] = (d_1339_v110_).f12
        nw238_[int(9)] = (d_1339_v110_).f12
        nw238_[int(10)] = d_1339_v110_.f13
        nw238_[int(11)] = self.f7
        nw238_[int(12)] = self.f7
        nw238_[int(13)] = (d_1339_v110_).f12
        nw238_[int(14)] = self.f7
        nw238_[int(15)] = self.f7
        nw238_[int(16)] = not((d_1339_v110_).f12)
        nw238_[int(17)] = (d_1339_v110_).f12
        nw238_[int(18)] = (d_1339_v110_).f12
        nw238_[int(19)] = self.f7
        d_1341_v112_ = nw238_
        d_1342_v113_: str
        d_1342_v113_ = _dafny.CodePoint('y')
        d_1343_v114_: _dafny.Map
        d_1343_v114_ = _dafny.Map({d_1341_v112_: d_1342_v113_})
        d_1344_v115_: _dafny.Seq
        d_1344_v115_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acjxxafo"))
        d_1343_v114_ = (d_1343_v114_).set(d_1341_v112_, (d_1342_v113_ if (D1_DC4(d_1339_v110_.f13, (d_1340_v111_)[default__.safeIndex(len(d_1344_v115_), len(d_1340_v111_))])).cf8 else d_1342_v113_))
        d_1342_v113_ = d_1342_v113_
        d_1345_v116_: _dafny.Set
        d_1345_v116_ = _dafny.Set({self.f7})
        d_1346_v117_: _dafny.Set
        d_1346_v117_ = _dafny.Set({d_1345_v116_, d_1345_v116_, d_1345_v116_, d_1345_v116_, (d_1345_v116_) - (d_1345_v116_)})
        d_1346_v117_ = d_1346_v117_
        d_1347_v118_: _dafny.MultiSet
        d_1347_v118_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cejtyhlrx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vervc"))])
        d_1348_v119_: _dafny.MultiSet
        d_1348_v119_ = _dafny.MultiSet([self.f7])
        d_1349_v120_: _dafny.Seq
        d_1349_v120_ = _dafny.SeqWithoutIsStrInference([d_1348_v119_, d_1348_v119_])
        d_1350_v121_: _dafny.MultiSet
        d_1350_v121_ = _dafny.MultiSet([(self).f14])
        d_1351_v122_: D2
        d_1351_v122_ = D2_DC8(False, d_1344_v115_)
        d_1352_v123_: _dafny.Array
        nw239_ = _dafny.Array(None, 28)
        nw239_[int(0)] = (self).f14
        nw239_[int(1)] = (self).f14
        nw239_[int(2)] = (d_1196_v2_)[default__.safeIndex((self).f14, len(d_1196_v2_))]
        nw239_[int(3)] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1347_v118_).cardinality: (self).fm9(d_1339_v110_.f13, (self).f14, d_1342_v113_, globalState)})])))
        nw239_[int(4)] = (0) - ((self).f14)
        nw239_[int(5)] = (self).f14
        nw239_[int(6)] = (self).f14
        nw239_[int(7)] = (self).f14
        nw239_[int(8)] = (self).f14
        nw239_[int(9)] = (self).f14
        nw239_[int(10)] = (self).f14
        nw239_[int(11)] = (self).f14
        nw239_[int(12)] = (self).f14
        nw239_[int(13)] = (self).f14
        nw239_[int(14)] = (self).f14
        nw239_[int(15)] = (self).f14
        nw239_[int(16)] = (self).f14
        nw239_[int(17)] = 84
        nw239_[int(18)] = (self).f14
        nw239_[int(19)] = (self).f14
        nw239_[int(20)] = len(d_1349_v120_)
        nw239_[int(21)] = (self).f14
        nw239_[int(22)] = (self).f14
        nw239_[int(23)] = ((d_1350_v121_)[(0) - ((self).f14)] if ((0) - ((self).f14)) in (d_1350_v121_) else (self).f14)
        nw239_[int(24)] = default__.fm0(globalState)
        nw239_[int(25)] = len(d_1345_v116_)
        nw239_[int(26)] = (self).f14
        nw239_[int(27)] = (0) - (len((d_1351_v122_).cf15))
        d_1352_v123_ = nw239_
        d_1353_v124_: _dafny.Map
        d_1353_v124_ = _dafny.Map({d_1352_v123_: _dafny.Set({(self).f14, default__.fm0(globalState)})})
        d_1354_v125_: D13
        d_1354_v125_ = D13_DC34(d_1353_v124_)
        (d_1339_v110_).f13 = (d_1352_v123_) not in ((d_1354_v125_).cf57)

    def m11(self, globalState):
        r0: bool = False
        r1: bool = False
        d_1355_v0_: D1
        d_1355_v0_ = D1_DC4(self.f7, self.f7)
        d_1356_i0_: int
        d_1356_i0_ = 0
        with _dafny.label("6"):
            def lambda86_(source10_):
                if source10_.is_DC4:
                    d_1360___mcc_h0_ = source10_.cf7
                    d_1361___mcc_h1_ = source10_.cf8
                    d_1362_cf8_ = d_1361___mcc_h1_
                    d_1363_cf7_ = d_1360___mcc_h0_
                    return d_1363_cf7_
                elif source10_.is_DC5:
                    d_1364___mcc_h2_ = source10_.cf9
                    d_1365___mcc_h3_ = source10_.cf10
                    d_1366___mcc_h4_ = source10_.cf11
                    d_1367_cf11_ = d_1366___mcc_h4_
                    d_1368_cf10_ = d_1365___mcc_h3_
                    d_1369_cf9_ = d_1364___mcc_h2_
                    return self.f7
                elif source10_.is_DC3:
                    d_1370___mcc_h5_ = source10_.cf6
                    d_1371_cf6_ = d_1370___mcc_h5_
                    return (self.f7) and (self.f7)
                elif True:
                    d_1372___mcc_h6_ = source10_.cf12
                    d_1373_cf12_ = d_1372___mcc_h6_
                    return not(((self).f14) <= (len(_dafny.Set({self.f7, self.f7, self.f7}))))

            while lambda86_(d_1355_v0_):
                with _dafny.c_label("6"):
                    if (d_1356_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1356_i0_ = (d_1356_i0_) + (1)
                    d_1357_v1_: int
                    d_1357_v1_ = 123
                    d_1357_v1_ = default__.safeModuloInt((self).f14, d_1357_v1_)
                    d_1358_v2_: _dafny.Seq
                    d_1358_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxhikfk"))
                    default__.m0((self).f14, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdhqaeccs"))) + (d_1358_v2_)), globalState)
                    d_1359_v3_: _dafny.Map
                    d_1359_v3_ = _dafny.Map({d_1358_v2_: d_1357_v1_})
                    d_1357_v1_ = ((d_1359_v3_)[d_1358_v2_] if (d_1358_v2_) in (d_1359_v3_) else (0) - (len(d_1358_v2_)))
                    d_1358_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                    pass
            pass
        d_1374_v4_: _dafny.Array
        def lambda87_(d_1375_i1_):
            return _dafny.Map({False: (self).f14})

        init48_ = lambda87_
        nw240_ = _dafny.Array(None, 1)
        for i0_48_ in range(nw240_.length(0)):
            nw240_[i0_48_] = init48_(i0_48_)
        d_1374_v4_ = nw240_
        d_1376_v5_: _dafny.Map
        d_1376_v5_ = _dafny.Map({self.f7: d_1374_v4_})
        d_1376_v5_ = (d_1376_v5_).set(self.f7, d_1374_v4_)
        d_1377_v6_: _dafny.Array
        nw241_ = _dafny.Array(int(0), 15)
        d_1377_v6_ = nw241_
        d_1378_v7_: _dafny.Map
        d_1378_v7_ = _dafny.Map({d_1377_v6_: (self).f14})
        d_1379_v8_: _dafny.MultiSet
        d_1379_v8_ = _dafny.MultiSet([self.f7, self.f7, True, self.f7])
        d_1380_v9_: _dafny.Seq
        d_1380_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qifiugigg"))
        d_1381_v10_: _dafny.Seq
        d_1381_v10_ = _dafny.SeqWithoutIsStrInference([self.f7])
        d_1382_v11_: D12
        d_1382_v11_ = D12_DC30((self).f14, not(self.f7), _dafny.SeqWithoutIsStrInference([(self).f14]), len(d_1381_v10_))
        d_1383_v12_: _dafny.Array
        nw242_ = _dafny.Array(None, 23)
        nw242_[int(0)] = (self).f14
        nw242_[int(1)] = (self).f14
        nw242_[int(2)] = (self).f14
        nw242_[int(3)] = ((self).f14 if self.f7 else (self).f14)
        nw242_[int(4)] = (self).f14
        nw242_[int(5)] = (len(d_1378_v7_)) + (len(default__.fm21(self.f7, globalState)))
        nw242_[int(6)] = (self).f14
        nw242_[int(7)] = (self).f14
        nw242_[int(8)] = default__.fm0(globalState)
        nw242_[int(9)] = (self).f14
        nw242_[int(10)] = (self).f14
        nw242_[int(11)] = (self).f14
        nw242_[int(12)] = default__.safeDivisionInt(((d_1379_v8_)[self.f7] if (self.f7) in (d_1379_v8_) else (self).f14), (self).f14)
        nw242_[int(13)] = (self).f14
        nw242_[int(14)] = (self).f14
        nw242_[int(15)] = (self).f14
        nw242_[int(16)] = (self).f14
        nw242_[int(17)] = (self).f14
        nw242_[int(18)] = len(d_1380_v9_)
        nw242_[int(19)] = (d_1382_v11_).cf52
        nw242_[int(20)] = (self).f14
        nw242_[int(21)] = (self).f14
        nw242_[int(22)] = (self).f14
        d_1383_v12_ = nw242_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1377_v6_).length(0)):
            d_1384_i2_: int = guard_loop_3_
            if (True) and (((0) <= (d_1384_i2_)) and ((d_1384_i2_) < ((d_1377_v6_).length(0)))):
                (d_1377_v6_)[(d_1384_i2_)] = default__.safeModuloInt(d_1384_i2_, (self).f14)
        d_1385_v13_: _dafny.MultiSet
        d_1385_v13_ = _dafny.MultiSet([(0) - ((self).f14)])
        d_1386_v14_: _dafny.MultiSet
        d_1386_v14_ = _dafny.MultiSet([(d_1385_v13_).cardinality])
        d_1387_i3_: int
        d_1387_i3_ = 0
        with _dafny.label("7"):
            while (d_1385_v13_).issubset(d_1385_v13_):
                with _dafny.c_label("7"):
                    if (d_1387_i3_) >= (100):
                        raise _dafny.Break("7")
                    d_1387_i3_ = (d_1387_i3_) + (1)
                    d_1388_v15_: _dafny.Map
                    d_1388_v15_ = _dafny.Map({(_dafny.MultiSet([self.f7])).cardinality: self.f7})
                    d_1389_v16_: _dafny.Map
                    d_1389_v16_ = _dafny.Map({self.f7: ((d_1388_v15_)[(0) - ((self).f14)] if ((0) - ((self).f14)) in (d_1388_v15_) else self.f7)})
                    d_1390_v17_: _dafny.Seq
                    d_1390_v17_ = _dafny.SeqWithoutIsStrInference([D0_DC0(308)])
                    d_1389_v16_ = (d_1389_v16_).set(self.f7, default__.fm1(d_1390_v17_, (self).f14, len(d_1388_v15_), globalState))
                    d_1391_v18_: _dafny.Array
                    def lambda88_(d_1392_v10_, d_1393_v8_):
                        def lambda89_(d_1394_i4_):
                            return (d_1392_v10_)[default__.safeIndex(len(_dafny.Set({(d_1393_v8_).cardinality})), len(d_1392_v10_))]

                        return lambda89_

                    init49_ = lambda88_(d_1381_v10_, d_1379_v8_)
                    nw243_ = _dafny.Array(None, 2)
                    for i0_49_ in range(nw243_.length(0)):
                        nw243_[i0_49_] = init49_(i0_49_)
                    d_1391_v18_ = nw243_
                    index212_ = default__.safeIndex(645, (d_1391_v18_).length(0))
                    (d_1391_v18_)[index212_] = (len(d_1381_v10_)) < ((self).f14)
                    index213_ = default__.safeIndex(645, (d_1391_v18_).length(0))
                    (d_1391_v18_)[index213_] = self.f7
                    d_1383_v12_ = d_1383_v12_
                    index214_ = default__.safeIndex(645, (d_1391_v18_).length(0))
                    (d_1391_v18_)[index214_] = ((d_1379_v8_) | (d_1379_v8_)).ispropersubset((d_1379_v8_).intersection(_dafny.MultiSet([(d_1391_v18_)[default__.safeIndex(645, (d_1391_v18_).length(0))], True, True])))
                    pass
            pass
        d_1395_v19_: _dafny.Map
        d_1395_v19_ = _dafny.Map({self.f7: d_1381_v10_})
        d_1396_v20_: _dafny.Map
        d_1396_v20_ = _dafny.Map({default__.fm49(False, d_1395_v19_, self.f7, globalState): self.f7})
        d_1397_v21_: _dafny.Map
        d_1397_v21_ = _dafny.Map({self.f7: self.f7})
        d_1398_v22_: _dafny.Seq
        d_1398_v22_ = _dafny.SeqWithoutIsStrInference([d_1397_v21_])
        d_1399_v23_: _dafny.Map
        d_1399_v23_ = _dafny.Map({self.f7: (d_1398_v22_)[default__.safeIndex((self).f14, len(d_1398_v22_))]})
        d_1400_v24_: _dafny.Seq
        d_1400_v24_ = _dafny.SeqWithoutIsStrInference([len(d_1396_v20_), len(d_1380_v9_), len(((d_1399_v23_)[self.f7] if (self.f7) in (d_1399_v23_) else (d_1397_v21_).set(self.f7, self.f7))), (self).f14, ((self).f14) + (573)])
        d_1401_v25_: _dafny.Array
        nw244_ = _dafny.Array(False, 25)
        d_1401_v25_ = nw244_
        index215_ = default__.safeIndex(425, (d_1401_v25_).length(0))
        (d_1401_v25_)[index215_] = not(self.f7)
        d_1402_v26_: str
        d_1402_v26_ = _dafny.CodePoint('s')
        d_1403_v27_: _dafny.Set
        d_1403_v27_ = _dafny.Set({(self).f14})
        index216_ = default__.safeIndex(425, (d_1401_v25_).length(0))
        rhs213_ = ((self.f7) not in (_dafny.Map({self.f7: (self).f14}))) == (self.f7)
        rhs214_ = d_1400_v24_
        rhs215_ = ((d_1403_v27_ if self.f7 else d_1403_v27_)).issubset(d_1403_v27_)
        rhs216_ = d_1402_v26_
        rhs217_ = self.f7
        lhs128_ = d_1401_v25_
        lhs129_ = default__.safeIndex(425, (d_1401_v25_).length(0))
        lhs130_ = self
        r1 = rhs213_
        d_1400_v24_ = rhs214_
        lhs128_[lhs129_] = rhs215_
        d_1402_v26_ = rhs216_
        lhs130_.f7 = rhs217_
        d_1404_v28_: _dafny.Array
        nw245_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_1404_v28_ = nw245_
        rhs218_ = d_1404_v28_
        rhs219_ = ((d_1381_v10_)[default__.safeIndex(len(d_1380_v9_), len(d_1381_v10_))] if not((not(self.f7) if (d_1401_v25_)[default__.safeIndex(425, (d_1401_v25_).length(0))] else self.f7)) else self.f7)
        d_1404_v28_ = rhs218_
        r0 = rhs219_
        d_1405_v29_: D6
        d_1405_v29_ = D6_DC15(d_1402_v26_)
        pat_let_tv35_ = d_1402_v26_
        d_1406_v30_: _dafny.Map
        def iife119_(_pat_let44_0):
            def iife120_(d_1407_dt__update__tmp_h0_):
                def iife121_(_pat_let45_0):
                    def iife122_(d_1408_dt__update_hcf29_h0_):
                        return D6_DC15(d_1408_dt__update_hcf29_h0_)
                    return iife122_(_pat_let45_0)
                return iife121_(pat_let_tv35_)
            return iife120_(_pat_let44_0)
        d_1406_v30_ = _dafny.Map({iife119_(d_1405_v29_): (d_1401_v25_)[default__.safeIndex(425, (d_1401_v25_).length(0))]})
        r0 = (d_1405_v29_) not in (d_1406_v30_)
        d_1409_v31_: _dafny.Seq
        d_1409_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1402_v26_ for d_1410_i5_ in range(default__.abs(-313))])])
        r1 = (((d_1409_v31_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1402_v26_ for d_1411_i6_ in range(default__.abs(928))])), len(d_1409_v31_)), _dafny.SeqWithoutIsStrInference([d_1402_v26_ for d_1412_i7_ in range(default__.abs(-208))]))) + (d_1409_v31_)) != ((_dafny.SeqWithoutIsStrInference([d_1380_v9_ for d_1413_i8_ in range(default__.abs(137))])) + (d_1409_v31_))
        return r0, r1

    def m12(self, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: bool = False
        r3: _dafny.MultiSet = _dafny.MultiSet({})
        d_1414_v0_: str
        d_1414_v0_ = _dafny.CodePoint('p')
        d_1415_v1_: _dafny.Map
        d_1415_v1_ = _dafny.Map({self.f7: (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1416_i0_ in range(default__.abs(43))])).set(default__.safeIndex((self).f14, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1417_i0_ in range(default__.abs(43))]))), d_1414_v0_)})
        d_1418_v2_: _dafny.Seq
        d_1418_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vau"))
        d_1415_v1_ = (d_1415_v1_).set(True, (default__.fm35((self).f14, (0) - ((self).f14), globalState)) + (d_1418_v2_))
        d_1419_v3_: _dafny.Seq
        d_1419_v3_ = _dafny.SeqWithoutIsStrInference([self.f7])
        d_1420_v4_: _dafny.Map
        d_1420_v4_ = _dafny.Map({(self).f14: len(d_1419_v3_)})
        d_1421_v5_: _dafny.Map
        d_1421_v5_ = _dafny.Map({len((d_1420_v4_) | (d_1420_v4_)): self.f7})
        d_1421_v5_ = (d_1421_v5_).set((0) - ((self).f14), self.f7)
        default__.m0((0) - (default__.safeDivisionInt((self).f14, (self).f14)), (self).f14, globalState)
        r0 = self.f7
        hi4_ = (self).f14
        for d_1422_i1_ in range(default__.fm0(globalState), hi4_):
            d_1423_v6_: D12
            d_1423_v6_ = D12_DC32(self.f7)
            d_1424_v7_: _dafny.Set
            d_1424_v7_ = _dafny.Set({d_1418_v2_})
            d_1425_v8_: D0
            d_1425_v8_ = D0_DC0((self).f14)
            d_1426_v9_: _dafny.Seq
            d_1426_v9_ = _dafny.SeqWithoutIsStrInference([d_1425_v8_])
            d_1427_v10_: _dafny.Map
            d_1427_v10_ = _dafny.Map({d_1418_v2_: self.f7})
            d_1428_v11_: _dafny.Array
            def lambda90_(d_1429_i1_):
                def lambda91_(d_1430_i2_):
                    return (d_1430_i2_) + (d_1429_i1_)

                return lambda91_

            init50_ = lambda90_(d_1422_i1_)
            nw246_ = _dafny.Array(None, 2)
            for i0_50_ in range(nw246_.length(0)):
                nw246_[i0_50_] = init50_(i0_50_)
            d_1428_v11_ = nw246_
            d_1431_v12_: _dafny.MultiSet
            d_1431_v12_ = _dafny.MultiSet([d_1428_v11_])
            d_1432_v13_: _dafny.Set
            d_1432_v13_ = _dafny.Set({d_1414_v0_})
            d_1433_v14_: _dafny.Seq
            d_1433_v14_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14, (self).f14])
            d_1434_v15_: _dafny.Seq
            d_1434_v15_ = _dafny.SeqWithoutIsStrInference([len(d_1433_v14_)])
            d_1435_v16_: _dafny.Array
            nw247_ = _dafny.Array(None, 23)
            nw247_[int(0)] = self.f7
            nw247_[int(1)] = self.f7
            nw247_[int(2)] = self.f7
            nw247_[int(3)] = (self.f7) or (self.f7)
            nw247_[int(4)] = default__.fm1(d_1426_v9_, d_1422_i1_, 503, globalState)
            nw247_[int(5)] = self.f7
            nw247_[int(6)] = not((d_1419_v3_)[default__.safeIndex((self).f14, len(d_1419_v3_))])
            nw247_[int(7)] = self.f7
            nw247_[int(8)] = self.f7
            nw247_[int(9)] = self.f7
            nw247_[int(10)] = self.f7
            nw247_[int(11)] = not (self.f7) or (self.f7)
            nw247_[int(12)] = self.f7
            nw247_[int(13)] = ((d_1427_v10_)[d_1418_v2_] if (d_1418_v2_) in (d_1427_v10_) else self.f7)
            nw247_[int(14)] = self.f7
            nw247_[int(15)] = True
            nw247_[int(16)] = not(self.f7)
            nw247_[int(17)] = ((d_1431_v12_).set(d_1428_v11_, default__.abs(d_1422_i1_))).ispropersubset(d_1431_v12_)
            nw247_[int(18)] = self.f7
            nw247_[int(19)] = default__.fm1(d_1426_v9_, len(((default__.fm50(d_1432_v13_, d_1424_v7_, (0) - (default__.fm0(globalState)), globalState)).set(default__.safeIndex((self).f14, len(default__.fm50(d_1432_v13_, d_1424_v7_, (0) - (default__.fm0(globalState)), globalState))), self.f7)).set(default__.safeIndex((0) - ((self).f14), len((default__.fm50(d_1432_v13_, d_1424_v7_, (0) - (default__.fm0(globalState)), globalState)).set(default__.safeIndex((self).f14, len(default__.fm50(d_1432_v13_, d_1424_v7_, (0) - (default__.fm0(globalState)), globalState))), self.f7))), self.f7)), 295, globalState)
            nw247_[int(20)] = self.f7
            nw247_[int(21)] = False
            nw247_[int(22)] = (d_1422_i1_) == (len(d_1434_v15_))
            d_1435_v16_ = nw247_
            rhs220_ = _dafny.SeqWithoutIsStrInference([(d_1419_v3_) < (d_1419_v3_)])
            rhs221_ = d_1423_v6_
            rhs222_ = d_1424_v7_
            rhs223_ = d_1435_v16_
            d_1419_v3_ = rhs220_
            d_1423_v6_ = rhs221_
            d_1424_v7_ = rhs222_
            d_1435_v16_ = rhs223_
            index217_ = default__.safeIndex(744, (d_1435_v16_).length(0))
            (d_1435_v16_)[index217_] = self.f7
            d_1436_v17_: _dafny.Seq
            d_1436_v17_ = _dafny.SeqWithoutIsStrInference([d_1418_v2_, _dafny.SeqWithoutIsStrInference([d_1414_v0_]), d_1418_v2_])
            index218_ = default__.safeIndex(744, (d_1435_v16_).length(0))
            (d_1435_v16_)[index218_] = default__.fm1(_dafny.SeqWithoutIsStrInference([d_1425_v8_ for d_1437_i3_ in range(default__.abs(822))]), default__.safeDivisionInt(d_1422_i1_, len(d_1436_v17_)), d_1422_i1_, globalState)
            d_1438_v18_: _dafny.MultiSet
            d_1438_v18_ = _dafny.MultiSet([(d_1436_v17_)[default__.safeIndex((self).f14, len(d_1436_v17_))], d_1418_v2_])
            d_1439_v19_: _dafny.Seq
            d_1439_v19_ = _dafny.SeqWithoutIsStrInference([d_1438_v18_])
            r2 = not ((_dafny.MultiSet(d_1436_v17_)) == ((d_1439_v19_)[default__.safeIndex((0) - ((self).f14), len(d_1439_v19_))])) or (self.f7)
            d_1418_v2_ = d_1418_v2_
        d_1440_v20_: _dafny.Array
        nw248_ = _dafny.Array(None, 8)
        nw248_[int(0)] = len(d_1419_v3_)
        nw248_[int(1)] = len(_dafny.SeqWithoutIsStrInference([True, not(self.f7)]))
        nw248_[int(2)] = (self).f14
        nw248_[int(3)] = (self).f14
        nw248_[int(4)] = (self).f14
        nw248_[int(5)] = (self).f14
        nw248_[int(6)] = (self).f14
        nw248_[int(7)] = (self).f14
        d_1440_v20_ = nw248_
        d_1441_v21_: D9
        d_1441_v21_ = D9_DC23(d_1440_v20_)
        d_1442_v22_: _dafny.Seq
        d_1442_v22_ = _dafny.SeqWithoutIsStrInference([(d_1441_v21_).cf44])
        d_1443_v23_: _dafny.Set
        d_1443_v23_ = _dafny.Set({((self).f14) - ((self).f14), len(d_1442_v22_), (self).f14})
        d_1444_v24_: D2
        d_1444_v24_ = D2_DC7(d_1443_v23_)
        pat_let_tv36_ = d_1443_v23_
        def iife123_(_pat_let46_0):
            def iife124_(d_1445_dt__update__tmp_h0_):
                def iife125_(_pat_let47_0):
                    def iife126_(d_1446_dt__update_hcf13_h0_):
                        return D2_DC7(d_1446_dt__update_hcf13_h0_)
                    return iife126_(_pat_let47_0)
                return iife125_(pat_let_tv36_)
            return iife124_(_pat_let46_0)
        d_1443_v23_ = (iife123_(d_1444_v24_)).cf13
        r0 = self.f7
        d_1447_v25_: _dafny.MultiSet
        d_1447_v25_ = _dafny.MultiSet([self.f7, self.f7])
        r1 = (_dafny.MultiSet([not(self.f7)])).intersection(d_1447_v25_)
        r2 = not(True)
        r3 = d_1447_v25_
        return r0, r1, r2, r3

    @property
    def f14(self):
        return self._f14

class C4(T0):
    def  __init__(self):
        self._f7: bool = False
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def ctor__(self, f11, f7):
        (self)._f11 = f11
        (self).f7 = f7

    def fm15(self, globalState):
        return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1448_i0_ in range(default__.abs(538))]))) + ((-417) * (len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmpsycksp")))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))}))))

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Set = _dafny.Set({})
        d_1449_v0_: _dafny.Map
        d_1449_v0_ = _dafny.Map({p0: p2})
        d_1450_v1_: _dafny.Map
        d_1450_v1_ = _dafny.Map({len(d_1449_v0_): self.f7})
        d_1451_v2_: C0
        nw249_ = C0()
        nw249_.ctor__(((d_1450_v1_)[(0) - (p1)] if ((0) - (p1)) in (d_1450_v1_) else p0), p2)
        d_1451_v2_ = nw249_
        d_1452_i0_: int
        d_1452_i0_ = 0
        with _dafny.label("8"):
            while (d_1451_v2_).f12:
                with _dafny.c_label("8"):
                    if (d_1452_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_1452_i0_ = (d_1452_i0_) + (1)
                    d_1453_v3_: _dafny.Seq
                    d_1453_v3_ = _dafny.SeqWithoutIsStrInference([p0, d_1451_v2_.f13, self.f7])
                    d_1454_v4_: C0
                    nw250_ = C0()
                    nw250_.ctor__((len(d_1453_v3_)) != (default__.fm0(globalState)), p0)
                    d_1454_v4_ = nw250_
                    d_1455_v5_: int
                    d_1455_v5_ = 503
                    d_1455_v5_ = d_1455_v5_
                    d_1456_v6_: _dafny.Set
                    d_1456_v6_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_1451_v2_.f13]))})
                    d_1457_v7_: _dafny.Map
                    d_1457_v7_ = _dafny.Map({(d_1451_v2_.f13) not in (_dafny.Map({False: d_1456_v6_})): d_1453_v3_})
                    d_1457_v7_ = (d_1457_v7_).set(not(not((d_1451_v2_).f12)), d_1453_v3_)
                    default__.m0(p1, 10, globalState)
                    pass
            pass
        hi5_ = p1
        for d_1458_i1_ in range(p1, hi5_):
            d_1459_v9_: _dafny.MultiSet
            d_1459_v9_ = _dafny.MultiSet([d_1451_v2_.f13])
            d_1460_v10_: _dafny.Map
            d_1460_v10_ = _dafny.Map({p2: d_1458_i1_})
            d_1461_v11_: _dafny.Set
            d_1461_v11_ = _dafny.Set({len(d_1460_v10_), len(d_1460_v10_)})
            d_1462_v12_: _dafny.Seq
            d_1462_v12_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(d_1461_v11_): 542}))])
            d_1463_v13_: _dafny.Map
            d_1463_v13_ = _dafny.Map({d_1462_v12_: (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1464_i2_ in range(default__.abs(263))])))})
            d_1465_v14_: _dafny.Seq
            d_1465_v14_ = _dafny.SeqWithoutIsStrInference([p1, d_1458_i1_, len(d_1463_v13_), default__.fm0(globalState)])
            d_1466_v15_: D1
            d_1466_v15_ = D1_DC5(True, d_1462_v12_, p1)
            d_1467_v16_: _dafny.Map
            d_1467_v16_ = _dafny.Map({d_1459_v9_: d_1466_v15_})
            d_1468_v17_: _dafny.Array
            nw251_ = _dafny.Array(None, 12)
            nw251_[int(0)] = (0) - (p1)
            def iife127_():
                coll31_ = _dafny.Map()
                compr_31_: _dafny.MultiSet
                for compr_31_ in (d_1467_v16_).keys.Elements:
                    d_1469_v8_: _dafny.MultiSet = compr_31_
                    if (d_1469_v8_) in (d_1467_v16_):
                        coll31_[d_1469_v8_] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p1: d_1458_i1_})), -527, p1]))
                return _dafny.Map(coll31_)
            nw251_[int(1)] = (0) - (len(iife127_()
            ))
            nw251_[int(2)] = p1
            nw251_[int(3)] = p1
            nw251_[int(4)] = d_1458_i1_
            nw251_[int(5)] = p1
            nw251_[int(6)] = p1
            nw251_[int(7)] = d_1458_i1_
            nw251_[int(8)] = (p1) - (669)
            nw251_[int(9)] = d_1458_i1_
            nw251_[int(10)] = d_1458_i1_
            nw251_[int(11)] = p1
            d_1468_v17_ = nw251_
            index219_ = default__.safeIndex(560, (d_1468_v17_).length(0))
            (d_1468_v17_)[index219_] = d_1458_i1_
            d_1470_v18_: _dafny.MultiSet
            d_1470_v18_ = _dafny.MultiSet([(0) - (d_1458_i1_), p1, 146])
            d_1471_v19_: _dafny.Map
            d_1471_v19_ = _dafny.Map({p2: d_1462_v12_})
            index220_ = default__.safeIndex(560, (d_1468_v17_).length(0))
            (d_1468_v17_)[index220_] = (324) * ((((d_1470_v18_)[-348] if (-348) in (d_1470_v18_) else len(((d_1471_v19_)[(d_1451_v2_).f12] if ((d_1451_v2_).f12) in (d_1471_v19_) else d_1465_v14_)))) * (d_1458_i1_))
            d_1472_v20_: str
            d_1472_v20_ = _dafny.CodePoint('t')
            d_1473_v21_: _dafny.Seq
            d_1473_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anfua"))
            d_1474_v22_: C0
            nw252_ = C0()
            nw252_.ctor__((d_1461_v11_) == (d_1461_v11_), (d_1472_v20_) in (d_1473_v21_))
            d_1474_v22_ = nw252_
            d_1475_v23_: D0
            d_1475_v23_ = D0_DC1((0) - (d_1458_i1_))
            d_1476_v24_: _dafny.Set
            d_1476_v24_ = _dafny.Set({not(False), d_1451_v2_.f13, p0, d_1474_v22_.f13, self.f7})
            d_1477_v25_: _dafny.Array
            nw253_ = _dafny.Array(None, 25)
            nw253_[int(0)] = d_1475_v23_
            nw253_[int(1)] = d_1475_v23_
            nw253_[int(2)] = d_1475_v23_
            nw253_[int(3)] = d_1475_v23_
            nw253_[int(4)] = d_1475_v23_
            nw253_[int(5)] = d_1475_v23_
            nw253_[int(6)] = d_1475_v23_
            nw253_[int(7)] = d_1475_v23_
            nw253_[int(8)] = d_1475_v23_
            nw253_[int(9)] = d_1475_v23_
            nw253_[int(10)] = d_1475_v23_
            nw253_[int(11)] = d_1475_v23_
            nw253_[int(12)] = D0_DC1(p1)
            nw253_[int(13)] = d_1475_v23_
            nw253_[int(14)] = D0_DC1((0) - (p1))
            nw253_[int(15)] = d_1475_v23_
            nw253_[int(16)] = d_1475_v23_
            nw253_[int(17)] = d_1475_v23_
            nw253_[int(18)] = d_1475_v23_
            nw253_[int(19)] = d_1475_v23_
            nw253_[int(20)] = D0_DC1(len(d_1476_v24_))
            nw253_[int(21)] = d_1475_v23_
            nw253_[int(22)] = d_1475_v23_
            nw253_[int(23)] = d_1475_v23_
            nw253_[int(24)] = d_1475_v23_
            d_1477_v25_ = nw253_
            index221_ = default__.safeIndex(930, (d_1477_v25_).length(0))
            (d_1477_v25_)[index221_] = d_1475_v23_
            d_1478_v26_: _dafny.Array
            nw254_ = _dafny.Array(False, 9)
            d_1478_v26_ = nw254_
            d_1479_v27_: _dafny.Map
            d_1479_v27_ = _dafny.Map({d_1474_v22_.f13: d_1478_v26_})
            index222_ = default__.safeIndex(560, (d_1468_v17_).length(0))
            index223_ = default__.safeIndex(560, (d_1468_v17_).length(0))
            index224_ = default__.safeIndex(930, (d_1477_v25_).length(0))
            index225_ = default__.safeIndex(560, (d_1468_v17_).length(0))
            rhs224_ = ((d_1468_v17_)[default__.safeIndex(560, (d_1468_v17_).length(0))]) + (len((d_1479_v27_) | (d_1479_v27_)))
            rhs225_ = (self).fm15(globalState)
            rhs226_ = d_1475_v23_
            rhs227_ = (d_1468_v17_)[default__.safeIndex(560, (d_1468_v17_).length(0))]
            rhs228_ = (d_1451_v2_).f12
            lhs131_ = d_1468_v17_
            lhs132_ = default__.safeIndex(560, (d_1468_v17_).length(0))
            lhs133_ = d_1468_v17_
            lhs134_ = default__.safeIndex(560, (d_1468_v17_).length(0))
            lhs135_ = d_1477_v25_
            lhs136_ = default__.safeIndex(930, (d_1477_v25_).length(0))
            lhs137_ = d_1468_v17_
            lhs138_ = default__.safeIndex(560, (d_1468_v17_).length(0))
            lhs139_ = d_1451_v2_
            lhs131_[lhs132_] = rhs224_
            lhs133_[lhs134_] = rhs225_
            lhs135_[lhs136_] = rhs226_
            lhs137_[lhs138_] = rhs227_
            lhs139_.f13 = rhs228_
            if (p1) != (423):
                d_1480_v28_: _dafny.Array
                nw255_ = _dafny.Array(_dafny.Map({}), 1)
                d_1480_v28_ = nw255_
                index226_ = default__.safeIndex(11, (d_1480_v28_).length(0))
                (d_1480_v28_)[index226_] = d_1471_v19_
                d_1481_v29_: _dafny.Seq
                d_1481_v29_ = _dafny.SeqWithoutIsStrInference([d_1471_v19_, d_1471_v19_, _dafny.Map({(d_1451_v2_).f12: d_1462_v12_})])
                index227_ = default__.safeIndex(11, (d_1480_v28_).length(0))
                (d_1480_v28_)[index227_] = (d_1481_v29_)[default__.safeIndex((d_1468_v17_)[default__.safeIndex(560, (d_1468_v17_).length(0))], len(d_1481_v29_))]
                d_1482_v30_: _dafny.Map
                d_1482_v30_ = _dafny.Map({p1: d_1461_v11_})
                d_1482_v30_ = (d_1482_v30_).set(p1, default__.fm2((d_1468_v17_)[default__.safeIndex(560, (d_1468_v17_).length(0))], not(d_1474_v22_.f13), globalState))
                d_1483_v31_: _dafny.Map
                d_1483_v31_ = _dafny.Map({(d_1449_v0_) | (d_1449_v0_): (d_1474_v22_).f12})
                d_1483_v31_ = (d_1483_v31_).set(d_1449_v0_, p0)
                (d_1474_v22_).f13 = not (d_1474_v22_.f13) or (p0)
                index228_ = default__.safeIndex(560, (d_1468_v17_).length(0))
                (d_1468_v17_)[index228_] = ((d_1468_v17_)[default__.safeIndex(560, (d_1468_v17_).length(0))]) * (d_1458_i1_)
            elif True:
                default__.m0(((d_1470_v18_)[753] if (753) in (d_1470_v18_) else p1), ((d_1460_v10_)[d_1451_v2_.f13] if (d_1451_v2_.f13) in (d_1460_v10_) else d_1458_i1_), globalState)
                d_1484_v32_: _dafny.Map
                d_1484_v32_ = _dafny.Map({d_1475_v23_: (d_1458_i1_) + (p1)})
                d_1484_v32_ = d_1484_v32_
                d_1485_v33_: _dafny.Array
                nw256_ = _dafny.Array(_dafny.Map({}), 4)
                d_1485_v33_ = nw256_
                d_1485_v33_ = d_1485_v33_
                index229_ = default__.safeIndex(746, (d_1478_v26_).length(0))
                (d_1478_v26_)[index229_] = (d_1451_v2_).f12
                index230_ = default__.safeIndex(746, (d_1478_v26_).length(0))
                (d_1478_v26_)[index230_] = ((d_1474_v22_).f12 if ((d_1468_v17_)[default__.safeIndex(560, (d_1468_v17_).length(0))]) in (d_1470_v18_) else ((d_1451_v2_).f12 if (d_1474_v22_).f12 else d_1451_v2_.f13))
                d_1486_v34_: _dafny.Map
                d_1486_v34_ = _dafny.Map({d_1459_v9_: 822})
                d_1486_v34_ = d_1486_v34_
        d_1487_v36_: _dafny.Seq
        d_1487_v36_ = _dafny.SeqWithoutIsStrInference([(d_1451_v2_).f12, p2, False, d_1451_v2_.f13, (d_1451_v2_).f12])
        d_1488_v37_: _dafny.MultiSet
        d_1488_v37_ = _dafny.MultiSet([d_1487_v36_])
        d_1489_v38_: str
        d_1489_v38_ = _dafny.CodePoint('x')
        d_1490_v39_: _dafny.Map
        def iife128_():
            coll32_ = _dafny.Map()
            compr_32_: _dafny.Seq
            for compr_32_ in (d_1488_v37_).Elements:
                d_1491_v35_: _dafny.Seq = compr_32_
                if (d_1491_v35_) in (d_1488_v37_):
                    coll32_[d_1491_v35_] = p2
            return _dafny.Map(coll32_)
        d_1490_v39_ = _dafny.Map({len(iife128_()
        ): d_1489_v38_})
        d_1490_v39_ = (d_1490_v39_).set((0) - (p1), d_1489_v38_)
        if (d_1451_v2_.f13) or (p2):
            d_1492_v40_: int
            d_1492_v40_ = 128
            d_1492_v40_ = p1
            d_1493_v41_: _dafny.Set
            d_1493_v41_ = _dafny.Set({p2, d_1451_v2_.f13, d_1451_v2_.f13})
            d_1494_v42_: D1
            d_1494_v42_ = D1_DC3(d_1493_v41_)
            d_1495_v43_: D1
            d_1495_v43_ = D1_DC6(d_1494_v42_)
            source11_ = d_1495_v43_
            if source11_.is_DC4:
                d_1496___mcc_h0_ = source11_.cf7
                d_1497___mcc_h1_ = source11_.cf8
                d_1498_cf8_ = d_1497___mcc_h1_
                d_1499_cf7_ = d_1496___mcc_h0_
                d_1500_v44_: _dafny.Seq
                d_1500_v44_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1501_v45_: _dafny.Map
                d_1501_v45_ = _dafny.Map({(d_1500_v44_) <= (d_1500_v44_): _dafny.Map({(d_1451_v2_).f12: p1})})
                d_1502_v46_: _dafny.Map
                d_1502_v46_ = _dafny.Map({(D1_DC5(p2, _dafny.SeqWithoutIsStrInference([len(d_1493_v41_)]), d_1492_v40_)).cf9: d_1492_v40_})
                d_1501_v45_ = _dafny.Map({p0: d_1502_v46_})
                d_1492_v40_ = -24
                d_1503_v47_: _dafny.MultiSet
                d_1503_v47_ = _dafny.MultiSet([d_1489_v38_, default__.fm18(globalState), d_1489_v38_])
                d_1504_v48_: _dafny.Map
                d_1504_v48_ = _dafny.Map({(d_1493_v41_) != (d_1493_v41_): ((d_1503_v47_).set(d_1489_v38_, default__.abs(d_1492_v40_))) | (d_1503_v47_)})
                d_1503_v47_ = ((d_1504_v48_)[d_1451_v2_.f13] if (d_1451_v2_.f13) in (d_1504_v48_) else _dafny.MultiSet([d_1489_v38_]))
                d_1505_v49_: _dafny.Array
                nw257_ = _dafny.Array(int(0), 26)
                d_1505_v49_ = nw257_
                d_1506_v50_: _dafny.Seq
                d_1506_v50_ = _dafny.SeqWithoutIsStrInference([d_1505_v49_])
                d_1507_v51_: _dafny.Map
                d_1507_v51_ = _dafny.Map({d_1492_v40_: d_1506_v50_})
                d_1507_v51_ = (d_1507_v51_).set(d_1492_v40_, (d_1506_v50_) + (_dafny.SeqWithoutIsStrInference([d_1505_v49_, d_1505_v49_, d_1505_v49_, d_1505_v49_])))
            elif source11_.is_DC5:
                d_1508___mcc_h2_ = source11_.cf9
                d_1509___mcc_h3_ = source11_.cf10
                d_1510___mcc_h4_ = source11_.cf11
                d_1511_cf11_ = d_1510___mcc_h4_
                d_1512_cf10_ = d_1509___mcc_h3_
                d_1513_cf9_ = d_1508___mcc_h2_
                d_1514_v52_: _dafny.Seq
                d_1514_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vu"))
                d_1514_v52_ = (d_1514_v52_).set(default__.safeIndex(d_1492_v40_, len(d_1514_v52_)), d_1489_v38_)
                d_1513_cf9_ = True
                d_1515_v53_: _dafny.MultiSet
                d_1515_v53_ = _dafny.MultiSet([d_1451_v2_])
                d_1516_v54_: _dafny.Seq
                d_1516_v54_ = _dafny.SeqWithoutIsStrInference([d_1451_v2_])
                d_1517_v55_: _dafny.Seq
                d_1517_v55_ = _dafny.SeqWithoutIsStrInference([(d_1516_v54_)[default__.safeIndex(d_1492_v40_, len(d_1516_v54_))]])
                d_1518_v56_: _dafny.Map
                d_1518_v56_ = _dafny.Map({p0: d_1516_v54_})
                d_1519_v57_: _dafny.Array
                nw258_ = _dafny.Array(None, 17)
                nw258_[int(0)] = d_1515_v53_
                nw258_[int(1)] = (_dafny.MultiSet([d_1451_v2_, d_1451_v2_])) - (d_1515_v53_)
                nw258_[int(2)] = (d_1515_v53_) - (d_1515_v53_)
                nw258_[int(3)] = (d_1515_v53_).intersection(d_1515_v53_)
                nw258_[int(4)] = d_1515_v53_
                nw258_[int(5)] = _dafny.MultiSet(d_1517_v55_)
                nw258_[int(6)] = _dafny.MultiSet((((d_1518_v56_)[(d_1451_v2_).f12] if ((d_1451_v2_).f12) in (d_1518_v56_) else d_1517_v55_) if d_1451_v2_.f13 else _dafny.SeqWithoutIsStrInference([d_1451_v2_, d_1451_v2_])))
                nw258_[int(7)] = d_1515_v53_
                nw258_[int(8)] = d_1515_v53_
                nw258_[int(9)] = d_1515_v53_
                nw258_[int(10)] = d_1515_v53_
                nw258_[int(11)] = _dafny.MultiSet(d_1517_v55_)
                nw258_[int(12)] = d_1515_v53_
                nw258_[int(13)] = _dafny.MultiSet([d_1451_v2_, d_1451_v2_])
                nw258_[int(14)] = d_1515_v53_
                nw258_[int(15)] = d_1515_v53_
                nw258_[int(16)] = d_1515_v53_
                d_1519_v57_ = nw258_
                index231_ = default__.safeIndex(324, (d_1519_v57_).length(0))
                (d_1519_v57_)[index231_] = d_1515_v53_
                index232_ = default__.safeIndex(324, (d_1519_v57_).length(0))
                (d_1519_v57_)[index232_] = d_1515_v53_
                d_1492_v40_ = d_1511_cf11_
            elif source11_.is_DC3:
                d_1520___mcc_h5_ = source11_.cf6
                d_1521_cf6_ = d_1520___mcc_h5_
                d_1522_v58_: C0
                nw259_ = C0()
                nw259_.ctor__((p0) and ((d_1451_v2_).f12), (d_1451_v2_).f12)
                d_1522_v58_ = nw259_
                d_1523_v59_: _dafny.Array
                def lambda92_(d_1524_v2_):
                    def lambda93_(d_1525_i3_):
                        return d_1524_v2_.f13

                    return lambda93_

                init51_ = lambda92_(d_1451_v2_)
                nw260_ = _dafny.Array(None, 17)
                for i0_51_ in range(nw260_.length(0)):
                    nw260_[i0_51_] = init51_(i0_51_)
                d_1523_v59_ = nw260_
                d_1526_v60_: _dafny.Seq
                d_1526_v60_ = _dafny.SeqWithoutIsStrInference([d_1523_v59_, d_1523_v59_, (d_1523_v59_ if (d_1451_v2_).f12 else d_1523_v59_), d_1523_v59_])
                d_1523_v59_ = (d_1526_v60_)[default__.safeIndex((d_1492_v40_) * (d_1492_v40_), len(d_1526_v60_))]
                d_1527_v61_: _dafny.Array
                nw261_ = _dafny.Array(int(0), 6)
                d_1527_v61_ = nw261_
                rhs229_ = False
                rhs230_ = d_1527_v61_
                rhs231_ = d_1492_v40_
                lhs140_ = d_1451_v2_
                lhs140_.f13 = rhs229_
                d_1527_v61_ = rhs230_
                d_1492_v40_ = rhs231_
                (self).f7 = d_1522_v58_.f13
            elif True:
                d_1528___mcc_h6_ = source11_.cf12
                d_1529_cf12_ = d_1528___mcc_h6_
                d_1530_v62_: _dafny.Set
                d_1530_v62_ = _dafny.Set({p1})
                d_1531_v63_: _dafny.Map
                d_1531_v63_ = _dafny.Map({(d_1451_v2_).f12: len(d_1530_v62_)})
                d_1532_v64_: _dafny.Map
                d_1532_v64_ = _dafny.Map({d_1492_v40_: p1})
                d_1533_v65_: _dafny.Map
                d_1533_v65_ = _dafny.Map({d_1532_v64_: (d_1451_v2_).f12})
                d_1534_v66_: _dafny.MultiSet
                d_1534_v66_ = _dafny.MultiSet([_dafny.CodePoint('n')])
                d_1535_v67_: _dafny.Seq
                d_1535_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfqf"))
                d_1536_v68_: _dafny.Map
                d_1536_v68_ = _dafny.Map({d_1535_v67_: (d_1451_v2_).f12})
                d_1537_v69_: _dafny.Array
                nw262_ = _dafny.Array(None, 28)
                nw262_[int(0)] = (d_1534_v66_).cardinality
                nw262_[int(1)] = d_1492_v40_
                nw262_[int(2)] = d_1492_v40_
                nw262_[int(3)] = d_1492_v40_
                nw262_[int(4)] = p1
                nw262_[int(5)] = p1
                nw262_[int(6)] = 986
                nw262_[int(7)] = len(d_1536_v68_)
                nw262_[int(8)] = p1
                nw262_[int(9)] = p1
                nw262_[int(10)] = p1
                nw262_[int(11)] = p1
                nw262_[int(12)] = p1
                nw262_[int(13)] = len(d_1535_v67_)
                nw262_[int(14)] = (0) - (p1)
                nw262_[int(15)] = p1
                nw262_[int(16)] = d_1492_v40_
                nw262_[int(17)] = p1
                nw262_[int(18)] = p1
                nw262_[int(19)] = d_1492_v40_
                nw262_[int(20)] = p1
                nw262_[int(21)] = len(d_1532_v64_)
                nw262_[int(22)] = len(d_1531_v63_)
                nw262_[int(23)] = p1
                nw262_[int(24)] = d_1492_v40_
                nw262_[int(25)] = p1
                nw262_[int(26)] = p1
                nw262_[int(27)] = d_1492_v40_
                d_1537_v69_ = nw262_
                d_1538_v70_: _dafny.Set
                d_1538_v70_ = _dafny.Set({d_1537_v69_})
                d_1539_v71_: _dafny.MultiSet
                d_1539_v71_ = _dafny.MultiSet([(D0_DC2(p2, d_1492_v40_, d_1538_v70_, p1)).cf2])
                d_1540_v72_: _dafny.Seq
                d_1540_v72_ = _dafny.SeqWithoutIsStrInference([d_1492_v40_])
                d_1541_v73_: _dafny.Map
                d_1541_v73_ = _dafny.Map({((d_1533_v65_)[_dafny.Map({p1: (d_1539_v71_).cardinality})] if (_dafny.Map({p1: (d_1539_v71_).cardinality})) in (d_1533_v65_) else (d_1451_v2_).f12): d_1540_v72_})
                d_1542_v74_: _dafny.Array
                nw263_ = _dafny.Array(None, 10)
                nw263_[int(0)] = ((d_1531_v63_)[(d_1451_v2_).f12] if ((d_1451_v2_).f12) in (d_1531_v63_) else p1)
                nw263_[int(1)] = default__.safeModuloInt((self).fm15(globalState), d_1492_v40_)
                nw263_[int(2)] = d_1492_v40_
                nw263_[int(3)] = d_1492_v40_
                nw263_[int(4)] = len(d_1541_v73_)
                nw263_[int(5)] = d_1492_v40_
                nw263_[int(6)] = p1
                nw263_[int(7)] = default__.safeModuloInt(701, p1)
                nw263_[int(8)] = 395
                nw263_[int(9)] = ((D0_DC1((0) - (p1))).cf1) - (d_1492_v40_)
                d_1542_v74_ = nw263_
                index233_ = default__.safeIndex(102, (d_1537_v69_).length(0))
                (d_1537_v69_)[index233_] = ((d_1532_v64_)[d_1492_v40_] if (d_1492_v40_) in (d_1532_v64_) else (d_1540_v72_)[default__.safeIndex(d_1492_v40_, len(d_1540_v72_))])
                d_1543_v75_: _dafny.Array
                nw264_ = _dafny.Array(False, 25)
                d_1543_v75_ = nw264_
                index234_ = default__.safeIndex(102, (d_1537_v69_).length(0))
                rhs232_ = default__.safeDivisionInt(p1, d_1492_v40_)
                rhs233_ = d_1492_v40_
                rhs234_ = d_1543_v75_
                rhs235_ = d_1487_v36_
                rhs236_ = (d_1451_v2_).f12
                lhs141_ = d_1537_v69_
                lhs142_ = default__.safeIndex(102, (d_1537_v69_).length(0))
                lhs143_ = self
                d_1492_v40_ = rhs232_
                lhs141_[lhs142_] = rhs233_
                d_1543_v75_ = rhs234_
                d_1487_v36_ = rhs235_
                lhs143_.f7 = rhs236_
                d_1544_v76_: _dafny.Array
                nw265_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_1544_v76_ = nw265_
                d_1545_v77_: _dafny.Array
                nw266_ = _dafny.Array(_dafny.MultiSet({}), 18)
                d_1545_v77_ = nw266_
                index235_ = default__.safeIndex(366, (d_1544_v76_).length(0))
                (d_1544_v76_)[index235_] = d_1545_v77_
                d_1546_v78_: _dafny.Array
                nw267_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_1546_v78_ = nw267_
                index236_ = default__.safeIndex(475, (d_1546_v78_).length(0))
                (d_1546_v78_)[index236_] = d_1543_v75_
                index237_ = default__.safeIndex(366, (d_1544_v76_).length(0))
                index238_ = default__.safeIndex(475, (d_1546_v78_).length(0))
                rhs237_ = d_1545_v77_
                rhs238_ = False
                rhs239_ = p2
                rhs240_ = (p0 if True else (d_1451_v2_.f13) or ((d_1451_v2_).f12))
                rhs241_ = d_1543_v75_
                lhs144_ = d_1544_v76_
                lhs145_ = default__.safeIndex(366, (d_1544_v76_).length(0))
                lhs146_ = d_1451_v2_
                lhs147_ = d_1451_v2_
                lhs148_ = d_1451_v2_
                lhs149_ = d_1546_v78_
                lhs150_ = default__.safeIndex(475, (d_1546_v78_).length(0))
                lhs144_[lhs145_] = rhs237_
                lhs146_.f13 = rhs238_
                lhs147_.f13 = rhs239_
                lhs148_.f13 = rhs240_
                lhs149_[lhs150_] = rhs241_
                d_1547_v79_: _dafny.Seq
                d_1547_v79_ = _dafny.SeqWithoutIsStrInference([(d_1546_v78_)[default__.safeIndex(475, (d_1546_v78_).length(0))], (d_1546_v78_)[default__.safeIndex(475, (d_1546_v78_).length(0))], (d_1546_v78_)[default__.safeIndex(475, (d_1546_v78_).length(0))], d_1543_v75_, (d_1546_v78_)[default__.safeIndex(475, (d_1546_v78_).length(0))]])
                d_1548_v80_: _dafny.Map
                d_1548_v80_ = _dafny.Map({(d_1451_v2_).f12: (d_1547_v79_)[default__.safeIndex((d_1537_v69_)[default__.safeIndex(102, (d_1537_v69_).length(0))], len(d_1547_v79_))]})
                rhs242_ = d_1540_v72_
                rhs243_ = ((d_1548_v80_)[False] if (False) in (d_1548_v80_) else (d_1546_v78_)[default__.safeIndex(475, (d_1546_v78_).length(0))])
                d_1540_v72_ = rhs242_
                d_1543_v75_ = rhs243_
                d_1549_v81_: D0
                d_1549_v81_ = D0_DC1(default__.safeModuloInt(default__.fm0(globalState), d_1492_v40_))
                index239_ = default__.safeIndex(102, (d_1537_v69_).length(0))
                rhs244_ = p2
                rhs245_ = (((d_1535_v67_) + (d_1535_v67_)).set(default__.safeIndex((d_1549_v81_).cf1, len((d_1535_v67_) + (d_1535_v67_))), d_1489_v38_)) + (d_1535_v67_)
                rhs246_ = (d_1540_v72_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1550_i4_ in range(default__.abs(301))])), len(d_1540_v72_))]
                rhs247_ = d_1549_v81_
                rhs248_ = (d_1451_v2_).f12
                lhs151_ = self
                lhs152_ = d_1537_v69_
                lhs153_ = default__.safeIndex(102, (d_1537_v69_).length(0))
                lhs154_ = d_1451_v2_
                lhs151_.f7 = rhs244_
                d_1535_v67_ = rhs245_
                lhs152_[lhs153_] = rhs246_
                d_1549_v81_ = rhs247_
                lhs154_.f13 = rhs248_
            d_1551_v82_: D0
            d_1551_v82_ = D0_DC0(p1)
            d_1492_v40_ = default__.safeModuloInt(default__.safeDivisionInt(p1, d_1492_v40_), (d_1551_v82_).cf0)
            d_1552_v83_: _dafny.Seq
            d_1552_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btmsw"))
            d_1553_v84_: D2
            d_1553_v84_ = D2_DC8((d_1492_v40_) >= (p1), d_1552_v83_)
            d_1554_v85_: _dafny.Seq
            d_1554_v85_ = _dafny.SeqWithoutIsStrInference([d_1551_v82_, d_1551_v82_])
            pat_let_tv37_ = d_1554_v85_
            pat_let_tv38_ = p1
            pat_let_tv39_ = d_1492_v40_
            pat_let_tv40_ = globalState
            def iife129_(_pat_let48_0):
                def iife130_(d_1555_dt__update__tmp_h0_):
                    def iife131_(_pat_let49_0):
                        def iife132_(d_1556_dt__update_hcf14_h0_):
                            return D2_DC8(d_1556_dt__update_hcf14_h0_, (d_1555_dt__update__tmp_h0_).cf15)
                        return iife132_(_pat_let49_0)
                    return iife131_(default__.fm1(pat_let_tv37_, pat_let_tv38_, pat_let_tv39_, pat_let_tv40_))
                return iife130_(_pat_let48_0)
            d_1553_v84_ = iife129_(d_1553_v84_)
            d_1557_v86_: _dafny.Array
            nw268_ = _dafny.Array(None, 24)
            nw268_[int(0)] = p2
            nw268_[int(1)] = p2
            nw268_[int(2)] = (d_1451_v2_).fm16(True, globalState)
            nw268_[int(3)] = p0
            nw268_[int(4)] = (d_1451_v2_).f12
            nw268_[int(5)] = (d_1451_v2_).f12
            nw268_[int(6)] = d_1451_v2_.f13
            nw268_[int(7)] = (False) == (True)
            nw268_[int(8)] = True
            nw268_[int(9)] = p0
            nw268_[int(10)] = ((d_1451_v2_).f12 if (d_1487_v36_)[default__.safeIndex(p1, len(d_1487_v36_))] else (d_1451_v2_).f12)
            nw268_[int(11)] = True
            nw268_[int(12)] = ((d_1451_v2_).f12 if d_1451_v2_.f13 else (d_1451_v2_).f12)
            nw268_[int(13)] = d_1451_v2_.f13
            nw268_[int(14)] = (d_1451_v2_).f12
            nw268_[int(15)] = (d_1451_v2_).f12
            nw268_[int(16)] = False
            nw268_[int(17)] = self.f7
            nw268_[int(18)] = (False if self.f7 else p0)
            nw268_[int(19)] = (d_1451_v2_).f12
            nw268_[int(20)] = p0
            nw268_[int(21)] = d_1451_v2_.f13
            nw268_[int(22)] = (p1) >= (d_1492_v40_)
            nw268_[int(23)] = d_1451_v2_.f13
            d_1557_v86_ = nw268_
            index240_ = default__.safeIndex(719, (d_1557_v86_).length(0))
            (d_1557_v86_)[index240_] = d_1451_v2_.f13
            index241_ = default__.safeIndex(719, (d_1557_v86_).length(0))
            (d_1557_v86_)[index241_] = (p0) or (self.f7)
        elif True:
            (d_1451_v2_).f13 = False
            (self).f7 = d_1451_v2_.f13
            d_1558_v87_: D1
            d_1558_v87_ = D1_DC5((p1) > (35), _dafny.SeqWithoutIsStrInference([p1]), p1)
            d_1558_v87_ = d_1558_v87_
            d_1559_v88_: _dafny.Array
            nw269_ = _dafny.Array(None, 21)
            nw269_[int(0)] = (d_1451_v2_).f12
            nw269_[int(1)] = (d_1451_v2_).f12
            nw269_[int(2)] = d_1451_v2_.f13
            nw269_[int(3)] = p0
            nw269_[int(4)] = d_1451_v2_.f13
            nw269_[int(5)] = (d_1451_v2_).f12
            nw269_[int(6)] = True
            nw269_[int(7)] = d_1451_v2_.f13
            nw269_[int(8)] = p0
            nw269_[int(9)] = (d_1451_v2_).f12
            nw269_[int(10)] = self.f7
            nw269_[int(11)] = d_1451_v2_.f13
            nw269_[int(12)] = d_1451_v2_.f13
            nw269_[int(13)] = d_1451_v2_.f13
            nw269_[int(14)] = d_1451_v2_.f13
            nw269_[int(15)] = (d_1451_v2_).f12
            nw269_[int(16)] = (d_1451_v2_).f12
            nw269_[int(17)] = self.f7
            nw269_[int(18)] = True
            nw269_[int(19)] = (d_1451_v2_).f12
            nw269_[int(20)] = False
            d_1559_v88_ = nw269_
            d_1560_v89_: _dafny.Seq
            d_1560_v89_ = _dafny.SeqWithoutIsStrInference([d_1559_v88_])
            if (d_1451_v2_).fm16((_dafny.SeqWithoutIsStrInference([d_1559_v88_])) == (d_1560_v89_), globalState):
                d_1561_v90_: _dafny.Array
                def lambda94_(d_1562_p1_):
                    def lambda95_(d_1563_i5_):
                        return (d_1563_i5_) + (d_1562_p1_)

                    return lambda95_

                init52_ = lambda94_(p1)
                nw270_ = _dafny.Array(None, 29)
                for i0_52_ in range(nw270_.length(0)):
                    nw270_[i0_52_] = init52_(i0_52_)
                d_1561_v90_ = nw270_
                d_1564_v91_: _dafny.Map
                d_1564_v91_ = _dafny.Map({d_1489_v38_: d_1561_v90_})
                d_1564_v91_ = ((d_1564_v91_) | (d_1564_v91_) if True else ((d_1564_v91_).set(d_1489_v38_, d_1561_v90_)) | (_dafny.Map({d_1489_v38_: d_1561_v90_})))
                d_1565_v92_: _dafny.Array
                nw271_ = _dafny.Array(_dafny.Map({}), 24)
                d_1565_v92_ = nw271_
                d_1566_v93_: _dafny.Map
                d_1566_v93_ = _dafny.Map({(D3_DC10(d_1559_v88_)).cf21: d_1565_v92_})
                d_1567_v94_: _dafny.Array
                d_1567_v94_ = d_1565_v92_
                d_1568_v95_: _dafny.Seq
                d_1568_v95_ = _dafny.SeqWithoutIsStrInference([d_1565_v92_])
                d_1569_v96_: _dafny.MultiSet
                d_1569_v96_ = _dafny.MultiSet([False])
                d_1570_v97_: _dafny.Array
                nw272_ = _dafny.Array(None, 26)
                nw272_[int(0)] = d_1565_v92_
                nw272_[int(1)] = d_1565_v92_
                nw272_[int(2)] = d_1565_v92_
                nw272_[int(3)] = ((d_1566_v93_)[d_1559_v88_] if (d_1559_v88_) in (d_1566_v93_) else d_1565_v92_)
                nw272_[int(4)] = d_1565_v92_
                nw272_[int(5)] = d_1565_v92_
                nw272_[int(6)] = d_1565_v92_
                nw272_[int(7)] = d_1565_v92_
                nw272_[int(8)] = (d_1567_v94_)
                nw272_[int(9)] = d_1565_v92_
                nw272_[int(10)] = d_1565_v92_
                nw272_[int(11)] = d_1565_v92_
                nw272_[int(12)] = d_1565_v92_
                nw272_[int(13)] = d_1565_v92_
                nw272_[int(14)] = d_1565_v92_
                nw272_[int(15)] = d_1565_v92_
                nw272_[int(16)] = d_1565_v92_
                nw272_[int(17)] = d_1565_v92_
                nw272_[int(18)] = (d_1568_v95_)[default__.safeIndex(((d_1569_v96_)[True] if (True) in (d_1569_v96_) else p1), len(d_1568_v95_))]
                nw272_[int(19)] = d_1565_v92_
                nw272_[int(20)] = d_1565_v92_
                nw272_[int(21)] = (d_1567_v94_)
                nw272_[int(22)] = d_1565_v92_
                nw272_[int(23)] = d_1565_v92_
                nw272_[int(24)] = d_1565_v92_
                nw272_[int(25)] = d_1565_v92_
                d_1570_v97_ = nw272_
                index242_ = default__.safeIndex(422, (d_1570_v97_).length(0))
                (d_1570_v97_)[index242_] = d_1565_v92_
                index243_ = default__.safeIndex(802, (d_1559_v88_).length(0))
                (d_1559_v88_)[index243_] = d_1451_v2_.f13
                d_1571_v98_: _dafny.MultiSet
                d_1571_v98_ = _dafny.MultiSet([d_1561_v90_])
                index244_ = default__.safeIndex(422, (d_1570_v97_).length(0))
                index245_ = default__.safeIndex(802, (d_1559_v88_).length(0))
                rhs249_ = d_1489_v38_
                rhs250_ = (d_1565_v92_ if (d_1451_v2_).f12 else d_1565_v92_)
                rhs251_ = (d_1451_v2_).f12
                rhs252_ = d_1571_v98_
                lhs155_ = d_1570_v97_
                lhs156_ = default__.safeIndex(422, (d_1570_v97_).length(0))
                lhs157_ = d_1559_v88_
                lhs158_ = default__.safeIndex(802, (d_1559_v88_).length(0))
                d_1489_v38_ = rhs249_
                lhs155_[lhs156_] = rhs250_
                lhs157_[lhs158_] = rhs251_
                d_1571_v98_ = rhs252_
                d_1572_v99_: _dafny.Array
                nw273_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_1572_v99_ = nw273_
                d_1573_v100_: _dafny.Array
                def lambda96_(d_1574_i6_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))

                init53_ = lambda96_
                nw274_ = _dafny.Array(None, 22)
                for i0_53_ in range(nw274_.length(0)):
                    nw274_[i0_53_] = init53_(i0_53_)
                d_1573_v100_ = nw274_
                index246_ = default__.safeIndex(380, (d_1572_v99_).length(0))
                (d_1572_v99_)[index246_] = d_1573_v100_
                d_1575_v101_: _dafny.Seq
                d_1575_v101_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1576_v102_: _dafny.Map
                d_1576_v102_ = _dafny.Map({d_1575_v101_: d_1573_v100_})
                index247_ = default__.safeIndex(380, (d_1572_v99_).length(0))
                (d_1572_v99_)[index247_] = ((d_1576_v102_)[(d_1575_v101_ if not(d_1451_v2_.f13) else d_1575_v101_)] if ((d_1575_v101_ if not(d_1451_v2_.f13) else d_1575_v101_)) in (d_1576_v102_) else d_1573_v100_)
                index248_ = default__.safeIndex(802, (d_1559_v88_).length(0))
                (d_1559_v88_)[index248_] = self.f7
                d_1577_v103_: int
                d_1577_v103_ = 951
                d_1577_v103_ = p1
            elif True:
                (d_1451_v2_).f13 = (p2 if self.f7 else (d_1451_v2_).f12)
                d_1578_v104_: _dafny.Seq
                d_1578_v104_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_1579_v105_: C0
                nw275_ = C0()
                nw275_.ctor__(not(d_1451_v2_.f13), (d_1487_v36_)[default__.safeIndex((d_1578_v104_)[default__.safeIndex(p1, len(d_1578_v104_))], len(d_1487_v36_))])
                d_1579_v105_ = nw275_
                d_1580_v106_: _dafny.Seq
                d_1580_v106_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukiih"))
                d_1580_v106_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgmrpobis"))) + ((d_1580_v106_) + (d_1580_v106_))
                d_1581_v107_: _dafny.Map
                d_1581_v107_ = _dafny.Map({p1: (d_1578_v104_)[default__.safeIndex(p1, len(d_1578_v104_))]})
                d_1582_v108_: int
                d_1582_v108_ = 410
                d_1583_v109_: D0
                d_1583_v109_ = D0_DC1(d_1582_v108_)
                pat_let_tv41_ = p2
                pat_let_tv42_ = p1
                d_1584_v110_: _dafny.MultiSet
                def iife133_(_pat_let50_0):
                    def iife134_(d_1585_dt__update__tmp_h1_):
                        def iife135_(_pat_let51_0):
                            def iife136_(d_1586_dt__update_hcf9_h0_):
                                def iife137_(_pat_let52_0):
                                    def iife138_(d_1587_dt__update_hcf11_h0_):
                                        return D1_DC5(d_1586_dt__update_hcf9_h0_, (d_1585_dt__update__tmp_h1_).cf10, d_1587_dt__update_hcf11_h0_)
                                    return iife138_(_pat_let52_0)
                                return iife137_((0) - (pat_let_tv42_))
                            return iife136_(_pat_let51_0)
                        return iife135_(pat_let_tv41_)
                    return iife134_(_pat_let50_0)
                d_1584_v110_ = _dafny.MultiSet([d_1558_v87_, d_1558_v87_, iife133_(d_1558_v87_)])
                d_1588_v111_: _dafny.Map
                d_1588_v111_ = _dafny.Map({default__.fm18(globalState): p1})
                d_1589_v112_: _dafny.Map
                d_1589_v112_ = _dafny.Map({d_1449_v0_: d_1588_v111_})
                rhs253_ = (d_1589_v112_) != (d_1589_v112_)
                rhs254_ = default__.fm19(p1, globalState)
                rhs255_ = d_1582_v108_
                rhs256_ = d_1583_v109_
                rhs257_ = d_1584_v110_
                lhs159_ = d_1451_v2_
                lhs159_.f13 = rhs253_
                d_1581_v107_ = rhs254_
                d_1582_v108_ = rhs255_
                d_1583_v109_ = rhs256_
                d_1584_v110_ = rhs257_
                (d_1451_v2_).f13 = self.f7
            d_1590_v113_: _dafny.Seq
            d_1590_v113_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
            (d_1451_v2_).f13 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))) < (d_1590_v113_) if d_1451_v2_.f13 else ((d_1451_v2_).f12 if False else d_1451_v2_.f13))
        (self).f7 = (d_1451_v2_).f12
        r0 = d_1449_v0_
        d_1591_v114_: _dafny.Set
        d_1591_v114_ = _dafny.Set({not(self.f7)})
        r1 = d_1591_v114_
        return r0, r1

    def m3(self, globalState):
        d_1592_v0_: int
        d_1592_v0_ = 361
        d_1593_v1_: _dafny.Seq
        d_1593_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xl"))
        d_1594_v2_: D2
        d_1594_v2_ = D2_DC8(self.f7, d_1593_v1_)
        d_1592_v0_ = len(((d_1594_v2_).cf15) + (d_1593_v1_))
        hi6_ = d_1592_v0_
        for d_1595_i0_ in range(d_1592_v0_, hi6_):
            d_1596_v3_: D0
            d_1596_v3_ = D0_DC0(d_1595_i0_)
            d_1597_v4_: _dafny.Seq
            d_1597_v4_ = _dafny.SeqWithoutIsStrInference([d_1596_v3_])
            d_1598_v5_: C0
            nw276_ = C0()
            nw276_.ctor__(default__.fm1(d_1597_v4_, d_1595_i0_, (0) - (len((d_1593_v1_).set(default__.safeIndex(d_1592_v0_, len(d_1593_v1_)), _dafny.CodePoint('e')))), globalState), self.f7)
            d_1598_v5_ = nw276_
            d_1599_v6_: str
            d_1599_v6_ = _dafny.CodePoint('b')
            d_1600_v7_: _dafny.Map
            d_1600_v7_ = _dafny.Map({d_1598_v5_.f13: d_1599_v6_})
            d_1601_v8_: _dafny.Map
            d_1601_v8_ = _dafny.Map({d_1598_v5_.f13: (((default__.fm20((d_1598_v5_).fm16(d_1598_v5_.f13, globalState), (d_1598_v5_).f12, globalState) if (d_1598_v5_).f12 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcosxqa")))).set(default__.safeIndex(d_1592_v0_, len((default__.fm20((d_1598_v5_).fm16(d_1598_v5_.f13, globalState), (d_1598_v5_).f12, globalState) if (d_1598_v5_).f12 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcosxqa"))))), d_1599_v6_)).set(default__.safeIndex(d_1592_v0_, len(((default__.fm20((d_1598_v5_).fm16(d_1598_v5_.f13, globalState), (d_1598_v5_).f12, globalState) if (d_1598_v5_).f12 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcosxqa")))).set(default__.safeIndex(d_1592_v0_, len((default__.fm20((d_1598_v5_).fm16(d_1598_v5_.f13, globalState), (d_1598_v5_).f12, globalState) if (d_1598_v5_).f12 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcosxqa"))))), d_1599_v6_))), ((d_1600_v7_)[self.f7] if (self.f7) in (d_1600_v7_) else d_1599_v6_))})
            d_1593_v1_ = ((d_1601_v8_)[self.f7] if (self.f7) in (d_1601_v8_) else d_1593_v1_)
            if (d_1598_v5_).fm16((d_1598_v5_).f12, globalState):
                d_1602_v9_: _dafny.Array
                nw277_ = _dafny.Array(D3.default()(), 5)
                d_1602_v9_ = nw277_
                d_1603_v10_: _dafny.Map
                d_1603_v10_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1599_v6_, d_1599_v6_])): d_1602_v9_})
                default__.m0(len(d_1603_v10_), len(d_1593_v1_), globalState)
                d_1592_v0_ = ((self).fm15(globalState)) + (d_1592_v0_)
                (self).f7 = (d_1598_v5_).f12
                d_1592_v0_ = default__.safeModuloInt(d_1595_i0_, d_1595_i0_)
                d_1604_v11_: _dafny.Set
                d_1604_v11_ = _dafny.Set({d_1595_i0_})
                d_1604_v11_ = default__.fm2(default__.fm0(globalState), (d_1598_v5_.f13 if (d_1598_v5_).f12 else d_1598_v5_.f13), globalState)
            elif True:
                d_1605_v12_: _dafny.Seq
                d_1605_v12_ = _dafny.SeqWithoutIsStrInference([d_1592_v0_])
                d_1606_v13_: _dafny.Set
                d_1606_v13_ = _dafny.Set({d_1598_v5_.f13})
                d_1607_v14_: _dafny.Map
                d_1607_v14_ = _dafny.Map({d_1605_v12_: len(d_1606_v13_)})
                d_1592_v0_ = (len((d_1607_v14_) | (d_1607_v14_))) * (d_1595_i0_)
                d_1593_v1_ = d_1593_v1_
                d_1608_v15_: _dafny.Set
                d_1608_v15_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_1592_v0_ for d_1609_i1_ in range(default__.abs(-993))])})
                d_1610_v16_: _dafny.Map
                d_1610_v16_ = _dafny.Map({(d_1608_v15_) - (d_1608_v15_): d_1598_v5_.f13})
                d_1610_v16_ = (d_1610_v16_).set(d_1608_v15_, d_1598_v5_.f13)
                d_1611_v17_: _dafny.Array
                nw278_ = _dafny.Array(_dafny.MultiSet({}), 16)
                d_1611_v17_ = nw278_
                d_1611_v17_ = d_1611_v17_
                d_1592_v0_ = (0) - (((d_1592_v0_) * (d_1595_i0_)) * (d_1592_v0_))
            d_1612_v18_: D2
            d_1612_v18_ = D2_DC9(d_1592_v0_, d_1592_v0_, d_1593_v1_, True, _dafny.Map({D0_DC0(d_1592_v0_): d_1592_v0_}))
            d_1613_v19_: _dafny.Map
            d_1613_v19_ = _dafny.Map({D0_DC0(d_1592_v0_): 459})
            d_1614_v20_: D2
            d_1614_v20_ = D2_DC9((d_1612_v18_).cf17, d_1595_i0_, d_1593_v1_, False, d_1613_v19_)
            d_1615_v21_: _dafny.Seq
            d_1615_v21_ = _dafny.SeqWithoutIsStrInference([(d_1612_v18_).cf17, 718])
            d_1592_v0_ = (len(d_1615_v21_)) + (d_1595_i0_)
        d_1616_v22_: _dafny.Array
        nw279_ = _dafny.Array(_dafny.Array(None, 0), 8)
        d_1616_v22_ = nw279_
        d_1617_v23_: _dafny.Array
        nw280_ = _dafny.Array(None, 3)
        nw280_[int(0)] = d_1592_v0_
        nw280_[int(1)] = (0) - (d_1592_v0_)
        nw280_[int(2)] = d_1592_v0_
        d_1617_v23_ = nw280_
        d_1618_v24_: _dafny.Seq
        d_1618_v24_ = _dafny.SeqWithoutIsStrInference([d_1617_v23_])
        index249_ = default__.safeIndex(575, (d_1616_v22_).length(0))
        (d_1616_v22_)[index249_] = (d_1618_v24_)[default__.safeIndex(d_1592_v0_, len(d_1618_v24_))]
        index250_ = default__.safeIndex(575, (d_1616_v22_).length(0))
        (d_1616_v22_)[index250_] = d_1617_v23_
        d_1619_v25_: _dafny.Array
        nw281_ = _dafny.Array(_dafny.Map({}), 16)
        d_1619_v25_ = nw281_
        d_1620_v26_: _dafny.Map
        d_1620_v26_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f7])): (D1_DC4(True, self.f7)).cf7})
        index251_ = default__.safeIndex(523, (d_1619_v25_).length(0))
        (d_1619_v25_)[index251_] = d_1620_v26_
        index252_ = default__.safeIndex(523, (d_1619_v25_).length(0))
        (d_1619_v25_)[index252_] = (d_1620_v26_) | (_dafny.Map({_dafny.MultiSet([self.f7]): default__.fm1(_dafny.SeqWithoutIsStrInference([D0_DC0(d_1592_v0_) for d_1621_i2_ in range(default__.abs(492))]), d_1592_v0_, d_1592_v0_, globalState)}))
        d_1622_v27_: _dafny.Set
        d_1622_v27_ = _dafny.Set({(d_1616_v22_)[default__.safeIndex(575, (d_1616_v22_).length(0))]})
        d_1623_v28_: D0
        d_1623_v28_ = D0_DC2(self.f7, d_1592_v0_, d_1622_v27_, d_1592_v0_)
        hi7_ = d_1592_v0_
        for d_1624_i3_ in range((d_1623_v28_).cf3, hi7_):
            d_1625_v29_: _dafny.Map
            d_1625_v29_ = _dafny.Map({self.f7: self.f7})
            d_1626_v30_: _dafny.Seq
            d_1626_v30_ = _dafny.SeqWithoutIsStrInference([d_1625_v29_, _dafny.Map({self.f7: self.f7}), d_1625_v29_, _dafny.Map({self.f7: self.f7})])
            d_1627_v31_: _dafny.Seq
            d_1627_v31_ = _dafny.SeqWithoutIsStrInference([d_1625_v29_, (d_1626_v30_)[default__.safeIndex(d_1592_v0_, len(d_1626_v30_))], d_1625_v29_])
            d_1628_v32_: _dafny.MultiSet
            d_1628_v32_ = _dafny.MultiSet([d_1624_i3_, (0) - (d_1592_v0_), len((d_1627_v31_)[default__.safeIndex(d_1624_i3_, len(d_1627_v31_))])])
            (self).f7 = not((d_1624_i3_) not in (d_1628_v32_))
            d_1592_v0_ = d_1592_v0_
            (self).f7 = self.f7
            (self).m10(d_1592_v0_, globalState)
        d_1629_v33_: C0
        nw282_ = C0()
        nw282_.ctor__(not(self.f7), self.f7)
        d_1629_v33_ = nw282_
        d_1630_v34_: _dafny.Map
        d_1630_v34_ = _dafny.Map({d_1629_v33_: d_1629_v33_.f13})
        d_1631_v35_: _dafny.MultiSet
        d_1631_v35_ = _dafny.MultiSet([default__.fm0(globalState)])
        if ((d_1630_v34_)[d_1629_v33_] if (d_1629_v33_) in (d_1630_v34_) else (d_1631_v35_).issubset(d_1631_v35_)):
            (d_1629_v33_).f13 = False
            d_1632_v36_: _dafny.Array
            def lambda97_(d_1633_i4_):
                return _dafny.CodePoint('q')

            init54_ = lambda97_
            nw283_ = _dafny.Array(None, 22)
            for i0_54_ in range(nw283_.length(0)):
                nw283_[i0_54_] = init54_(i0_54_)
            d_1632_v36_ = nw283_
            d_1634_v37_: str
            d_1634_v37_ = _dafny.CodePoint('h')
            index253_ = default__.safeIndex(668, (d_1632_v36_).length(0))
            (d_1632_v36_)[index253_] = (d_1634_v37_ if self.f7 else d_1634_v37_)
            index254_ = default__.safeIndex(668, (d_1632_v36_).length(0))
            (d_1632_v36_)[index254_] = d_1634_v37_
            (d_1629_v33_).f13 = (d_1592_v0_) != (d_1592_v0_)
            d_1635_v38_: _dafny.Set
            d_1635_v38_ = _dafny.Set({d_1592_v0_})
            (d_1629_v33_).f13 = (d_1635_v38_).issubset(d_1635_v38_)
            d_1636_v39_: _dafny.Array
            nw284_ = _dafny.Array(None, 7)
            d_1636_v39_ = nw284_
            d_1637_v40_: T1
            nw285_ = C3()
            nw285_.ctor__(d_1592_v0_, not(d_1629_v33_.f13))
            d_1637_v40_ = nw285_
            index255_ = default__.safeIndex(616, (d_1636_v39_).length(0))
            (d_1636_v39_)[index255_] = d_1637_v40_
            index256_ = default__.safeIndex(616, (d_1636_v39_).length(0))
            (d_1636_v39_)[index256_] = d_1637_v40_
        elif True:
            d_1592_v0_ = d_1592_v0_
            d_1638_v41_: D0
            d_1638_v41_ = D0_DC0(508)
            d_1639_v42_: _dafny.Set
            d_1639_v42_ = _dafny.Set({d_1592_v0_})
            d_1640_v43_: _dafny.Map
            d_1640_v43_ = _dafny.Map({d_1592_v0_: d_1592_v0_})
            def iife139_():
                coll33_ = _dafny.Set()
                compr_33_: int
                for compr_33_ in (d_1640_v43_).keys.Elements:
                    d_1641_v44_: int = compr_33_
                    if (d_1641_v44_) in (d_1640_v43_):
                        coll33_ = coll33_.union(_dafny.Set([(d_1641_v44_) + ((0) - (-336))]))
                return _dafny.Set(coll33_)
            rhs258_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efqt"))
            rhs259_ = d_1593_v1_
            rhs260_ = ((d_1592_v0_) - (d_1592_v0_)) * (d_1592_v0_)
            rhs261_ = (d_1638_v41_).cf0
            rhs262_ = (iife139_()
            ).issubset((d_1639_v42_) | (d_1639_v42_))
            lhs160_ = self
            d_1593_v1_ = rhs258_
            d_1593_v1_ = rhs259_
            d_1592_v0_ = rhs260_
            d_1592_v0_ = rhs261_
            lhs160_.f7 = rhs262_
            d_1642_v45_: _dafny.MultiSet
            d_1642_v45_ = _dafny.MultiSet([False, self.f7, d_1629_v33_.f13])
            d_1643_v46_: _dafny.Seq
            d_1643_v46_ = _dafny.SeqWithoutIsStrInference([d_1642_v45_])
            d_1644_v47_: _dafny.Set
            d_1644_v47_ = _dafny.Set({d_1629_v33_.f13, False, self.f7, not (d_1629_v33_.f13) or (d_1629_v33_.f13)})
            rhs263_ = d_1643_v46_
            rhs264_ = default__.fm27((d_1629_v33_).f12, _dafny.CodePoint('t'), not((d_1592_v0_) != (d_1592_v0_)), self.f7, globalState)
            rhs265_ = (d_1592_v0_) - (default__.safeDivisionInt(d_1592_v0_, d_1592_v0_))
            rhs266_ = len(d_1593_v1_)
            rhs267_ = ((self).fm15(globalState)) - (d_1592_v0_)
            d_1643_v46_ = rhs263_
            d_1644_v47_ = rhs264_
            d_1592_v0_ = rhs265_
            d_1592_v0_ = rhs266_
            d_1592_v0_ = rhs267_
            d_1645_v48_: str
            d_1645_v48_ = _dafny.CodePoint('c')
            d_1645_v48_ = d_1645_v48_
            if not(((d_1644_v47_) | (d_1644_v47_)).isdisjoint(d_1644_v47_)):
                d_1646_v49_: C3
                nw286_ = C3()
                nw286_.ctor__(d_1592_v0_, self.f7)
                d_1646_v49_ = nw286_
                d_1647_v50_: _dafny.Map
                d_1647_v50_ = _dafny.Map({(len(d_1644_v47_) if d_1629_v33_.f13 else d_1592_v0_): not(((d_1646_v49_).f14) > (-68))})
                (d_1629_v33_).f13 = ((d_1647_v50_)[(d_1592_v0_ if d_1629_v33_.f13 else d_1592_v0_)] if ((d_1592_v0_ if d_1629_v33_.f13 else d_1592_v0_)) in (d_1647_v50_) else self.f7)
                d_1648_v51_: _dafny.Seq
                d_1648_v51_ = _dafny.SeqWithoutIsStrInference([(d_1629_v33_).f12])
                d_1649_v52_: _dafny.Map
                d_1649_v52_ = _dafny.Map({(d_1648_v51_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1645_v48_ for d_1650_i5_ in range(default__.abs(-478))])), len(d_1648_v51_))]: d_1642_v45_})
                d_1649_v52_ = (d_1649_v52_).set((d_1629_v33_).f12, _dafny.MultiSet(d_1648_v51_))
                d_1651_v53_: _dafny.Seq
                d_1651_v53_ = _dafny.SeqWithoutIsStrInference([d_1593_v1_])
                d_1651_v53_ = d_1651_v53_
                d_1652_v54_: _dafny.Map
                d_1652_v54_ = _dafny.Map({(d_1629_v33_).f12: d_1640_v43_})
                d_1653_v55_: C1
                nw287_ = C1()
                nw287_.ctor__((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbw")))) not in (((d_1652_v54_)[d_1629_v33_.f13] if (d_1629_v33_.f13) in (d_1652_v54_) else _dafny.Map({d_1592_v0_: len(d_1593_v1_)}))))
                d_1653_v55_ = nw287_
            elif True:
                d_1593_v1_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), d_1645_v48_])) + (d_1593_v1_)) + ((d_1593_v1_) + (d_1593_v1_))
                index257_ = default__.safeIndex(955, (d_1617_v23_).length(0))
                (d_1617_v23_)[index257_] = len((d_1593_v1_) + (d_1593_v1_))
                index258_ = default__.safeIndex(955, (d_1617_v23_).length(0))
                (d_1617_v23_)[index258_] = (self).fm15(globalState)
                d_1654_v56_: _dafny.Seq
                d_1654_v56_ = _dafny.SeqWithoutIsStrInference([not(not(False))])
                d_1592_v0_ = len((d_1654_v56_ if (d_1629_v33_.f13 if not(not(d_1629_v33_.f13)) else d_1629_v33_.f13) else d_1654_v56_))
                d_1655_v57_: D13
                d_1655_v57_ = D13_DC35()
                d_1656_v58_: D13
                d_1656_v58_ = D13_DC37(d_1655_v57_)
                d_1657_v59_: D13
                d_1657_v59_ = D13_DC37(d_1656_v58_)
                d_1658_v60_: D13
                d_1658_v60_ = D13_DC37(d_1656_v58_)
                d_1659_v61_: _dafny.Map
                d_1659_v61_ = _dafny.Map({d_1658_v60_: d_1642_v45_})
                (self).f7 = (d_1642_v45_).issubset(((d_1659_v61_)[d_1658_v60_] if (d_1658_v60_) in (d_1659_v61_) else d_1642_v45_))
                d_1660_v62_: _dafny.Map
                d_1660_v62_ = _dafny.Map({(d_1616_v22_)[default__.safeIndex(575, (d_1616_v22_).length(0))]: D1_DC4(True, d_1629_v33_.f13)})
                d_1661_v63_: D1
                d_1661_v63_ = D1_DC4(d_1629_v33_.f13, d_1629_v33_.f13)
                (d_1629_v33_).f13 = (d_1660_v62_) == (_dafny.Map({(d_1616_v22_)[default__.safeIndex(575, (d_1616_v22_).length(0))]: d_1661_v63_}))

    def m9(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: _dafny.Map = _dafny.Map({})
        d_1662_v0_: _dafny.Seq
        d_1662_v0_ = _dafny.SeqWithoutIsStrInference([self.f7])
        d_1663_v1_: C1
        nw288_ = C1()
        nw288_.ctor__((d_1662_v0_) <= (d_1662_v0_))
        d_1663_v1_ = nw288_
        d_1664_v2_: str
        d_1664_v2_ = _dafny.CodePoint('w')
        d_1665_v3_: _dafny.Set
        d_1665_v3_ = _dafny.Set({d_1664_v2_})
        d_1666_v4_: _dafny.Set
        d_1666_v4_ = d_1665_v3_
        def lambda98_(source12_):
            d_1667___mcc_h0_ = source12_
            d_1668_cf47_ = d_1667___mcc_h0_
            return self.f7

        if lambda98_(d_1666_v4_):
            d_1669_v5_: int
            d_1669_v5_ = 877
            d_1670_v6_: _dafny.Seq
            d_1670_v6_ = _dafny.SeqWithoutIsStrInference([d_1669_v5_, d_1669_v5_])
            d_1671_v7_: _dafny.Map
            d_1671_v7_ = _dafny.Map({self.f7: (d_1670_v6_) != ((_dafny.SeqWithoutIsStrInference([d_1669_v5_ for d_1672_i0_ in range(default__.abs(637))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1664_v2_ for d_1673_i1_ in range(default__.abs(94))])), len(_dafny.SeqWithoutIsStrInference([d_1669_v5_ for d_1674_i0_ in range(default__.abs(637))]))), d_1669_v5_))})
            d_1671_v7_ = d_1671_v7_
            d_1675_v8_: _dafny.Array
            nw289_ = _dafny.Array(_dafny.Seq({}), 9)
            d_1675_v8_ = nw289_
            index259_ = default__.safeIndex(440, (d_1675_v8_).length(0))
            (d_1675_v8_)[index259_] = d_1670_v6_
            d_1676_v9_: _dafny.Set
            d_1676_v9_ = _dafny.Set({self.f7})
            index260_ = default__.safeIndex(440, (d_1675_v8_).length(0))
            rhs268_ = (self.f7) not in ((d_1676_v9_) | (d_1676_v9_))
            rhs269_ = d_1670_v6_
            rhs270_ = d_1670_v6_
            rhs271_ = True
            rhs272_ = self.f7
            lhs161_ = self
            lhs162_ = d_1675_v8_
            lhs163_ = default__.safeIndex(440, (d_1675_v8_).length(0))
            lhs164_ = self
            lhs165_ = self
            lhs161_.f7 = rhs268_
            d_1670_v6_ = rhs269_
            lhs162_[lhs163_] = rhs270_
            lhs164_.f7 = rhs271_
            lhs165_.f7 = rhs272_
            (self).f7 = self.f7
            d_1677_v10_: D9
            d_1677_v10_ = D9_DC26((d_1675_v8_)[default__.safeIndex(440, (d_1675_v8_).length(0))])
            d_1678_v11_: D12
            d_1678_v11_ = D12_DC31(d_1677_v10_, self.f7)
            if not ((d_1678_v11_).cf54) or (self.f7):
                r2 = _dafny.Map({self.f7: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suupjuvr"))})
                d_1679_v12_: D0
                d_1679_v12_ = D0_DC0(d_1669_v5_)
                d_1680_v13_: C3
                nw290_ = C3()
                nw290_.ctor__((d_1669_v5_ if default__.fm1(_dafny.SeqWithoutIsStrInference([d_1679_v12_, d_1679_v12_, d_1679_v12_, d_1679_v12_, d_1679_v12_]), d_1669_v5_, default__.fm0(globalState), globalState) else d_1669_v5_), self.f7)
                d_1680_v13_ = nw290_
                (self).f7 = (self.f7 if self.f7 else (d_1662_v0_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_1680_v13_).f14 for d_1681_i2_ in range(default__.abs(436))])), len(d_1662_v0_))])
                d_1682_v14_: _dafny.MultiSet
                d_1682_v14_ = _dafny.MultiSet([d_1669_v5_])
                d_1665_v3_ = default__.fm30((d_1682_v14_).isdisjoint(d_1682_v14_), d_1669_v5_, ((0) - ((0) - (d_1669_v5_))) >= (718), globalState)
                d_1683_v15_: _dafny.Seq
                d_1683_v15_ = _dafny.SeqWithoutIsStrInference([d_1679_v12_])
                (self).f7 = (self.f7 if default__.fm1(default__.fm51(default__.fm18(globalState), self.f7, globalState), d_1669_v5_, d_1669_v5_, globalState) else default__.fm1(d_1683_v15_, d_1669_v5_, len(d_1670_v6_), globalState))
            elif True:
                d_1669_v5_ = d_1669_v5_
                d_1684_v16_: _dafny.Seq
                d_1684_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mj"))
                d_1685_v17_: _dafny.Array
                nw291_ = _dafny.Array(None, 7)
                nw291_[int(0)] = (d_1684_v16_) + (d_1684_v16_)
                nw291_[int(1)] = d_1684_v16_
                nw291_[int(2)] = d_1684_v16_
                nw291_[int(3)] = (d_1684_v16_) + (d_1684_v16_)
                nw291_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1664_v2_ for d_1686_i3_ in range(default__.abs(460))])
                nw291_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1664_v2_ for d_1687_i4_ in range(default__.abs(686))])
                nw291_[int(6)] = d_1684_v16_
                d_1685_v17_ = nw291_
                index261_ = default__.safeIndex(154, (d_1685_v17_).length(0))
                (d_1685_v17_)[index261_] = _dafny.SeqWithoutIsStrInference([d_1664_v2_ for d_1688_i5_ in range(default__.abs(818))])
                d_1689_v18_: _dafny.Seq
                d_1689_v18_ = _dafny.SeqWithoutIsStrInference([d_1684_v16_, d_1684_v16_, d_1684_v16_])
                index262_ = default__.safeIndex(154, (d_1685_v17_).length(0))
                (d_1685_v17_)[index262_] = ((d_1684_v16_ if self.f7 else d_1684_v16_)) + (((d_1689_v18_)[default__.safeIndex(d_1669_v5_, len(d_1689_v18_))]) + (d_1684_v16_))
                (self).f7 = self.f7
                d_1690_v19_: _dafny.Set
                d_1690_v19_ = _dafny.Set({d_1676_v9_})
                d_1691_v20_: _dafny.Set
                d_1691_v20_ = _dafny.Set({(0) - (len(d_1690_v19_))})
                d_1692_v21_: _dafny.Map
                d_1692_v21_ = _dafny.Map({d_1691_v20_: _dafny.SeqWithoutIsStrInference([d_1664_v2_ for d_1693_i6_ in range(default__.abs(-716))])})
                d_1692_v21_ = d_1692_v21_
                d_1694_v22_: _dafny.Map
                d_1694_v22_ = _dafny.Map({d_1669_v5_: self.f7})
                (self).f7 = ((d_1694_v22_)[d_1669_v5_] if (d_1669_v5_) in (d_1694_v22_) else self.f7)
            d_1695_v23_: _dafny.MultiSet
            d_1695_v23_ = _dafny.MultiSet([False, self.f7, not(False), self.f7])
            d_1696_v24_: _dafny.Map
            d_1696_v24_ = _dafny.Map({476: (d_1695_v23_).set(self.f7, default__.abs(d_1669_v5_))})
            (self).f7 = (d_1695_v23_).ispropersubset(((d_1696_v24_)[d_1669_v5_] if (d_1669_v5_) in (d_1696_v24_) else d_1695_v23_))
        elif True:
            d_1697_v25_: _dafny.Array
            nw292_ = _dafny.Array(False, 6)
            d_1697_v25_ = nw292_
            d_1698_v26_: _dafny.Set
            d_1698_v26_ = _dafny.Set({d_1697_v25_})
            d_1699_v27_: int
            d_1699_v27_ = 72
            d_1700_v28_: _dafny.Set
            d_1700_v28_ = _dafny.Set({self.f7, self.f7, self.f7, self.f7, self.f7})
            d_1701_v29_: C2
            nw293_ = C2()
            nw293_.ctor__(d_1698_v26_, d_1664_v2_, (d_1699_v27_) != (len(d_1700_v28_)))
            d_1701_v29_ = nw293_
            (self).f7 = (d_1662_v0_)[default__.safeIndex(d_1699_v27_, len(d_1662_v0_))]
            (self).f7 = self.f7
            (self).f7 = self.f7
            (self).f7 = self.f7
        d_1702_v30_: int
        d_1702_v30_ = 928
        if (d_1662_v0_)[default__.safeIndex(d_1702_v30_, len(d_1662_v0_))]:
            if default__.fm1(default__.fm51(default__.fm18(globalState), self.f7, globalState), default__.safeDivisionInt(d_1702_v30_, d_1702_v30_), (d_1702_v30_) * (d_1702_v30_), globalState):
                d_1703_v31_: _dafny.Array
                nw294_ = _dafny.Array(_dafny.Set({}), 23)
                d_1703_v31_ = nw294_
                d_1704_v32_: _dafny.Array
                def lambda99_(d_1705_v30_, d_1706_v0_):
                    def lambda100_(d_1707_i7_):
                        return D3_DC11(263, d_1705_v30_, len(_dafny.Map({len(d_1706_v0_): self.f7})), self.f7)

                    return lambda100_

                init55_ = lambda99_(d_1702_v30_, d_1662_v0_)
                nw295_ = _dafny.Array(None, 19)
                for i0_55_ in range(nw295_.length(0)):
                    nw295_[i0_55_] = init55_(i0_55_)
                d_1704_v32_ = nw295_
                d_1708_v33_: _dafny.Set
                d_1708_v33_ = _dafny.Set({d_1704_v32_})
                index263_ = default__.safeIndex(352, (d_1703_v31_).length(0))
                (d_1703_v31_)[index263_] = d_1708_v33_
                index264_ = default__.safeIndex(352, (d_1703_v31_).length(0))
                (d_1703_v31_)[index264_] = ((d_1708_v33_) - (d_1708_v33_)).intersection(d_1708_v33_)
                d_1702_v30_ = d_1702_v30_
                d_1709_v34_: _dafny.Seq
                d_1709_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcfdxmh"))
                d_1710_v35_: _dafny.Map
                d_1710_v35_ = _dafny.Map({d_1664_v2_: (d_1709_v34_) + (d_1709_v34_)})
                d_1711_v36_: _dafny.Array
                nw296_ = _dafny.Array(int(0), 25)
                d_1711_v36_ = nw296_
                index265_ = default__.safeIndex(120, (d_1711_v36_).length(0))
                (d_1711_v36_)[index265_] = d_1702_v30_
                d_1712_v37_: _dafny.Seq
                d_1712_v37_ = _dafny.SeqWithoutIsStrInference([d_1702_v30_])
                d_1713_v38_: _dafny.Set
                d_1713_v38_ = _dafny.Set({d_1702_v30_, d_1702_v30_, (0) - (len(d_1712_v37_)), d_1702_v30_, d_1702_v30_})
                d_1714_v39_: _dafny.Map
                d_1714_v39_ = _dafny.Map({self.f7: d_1702_v30_})
                pat_let_tv43_ = d_1710_v35_
                index266_ = default__.safeIndex(120, (d_1711_v36_).length(0))
                def iife140_(_pat_let53_0):
                    def iife141_(d_1715_dt__update__tmp_h0_):
                        def iife142_(_pat_let54_0):
                            def iife143_(d_1716_dt__update_hcf62_h0_):
                                return D14_DC38(d_1716_dt__update_hcf62_h0_)
                            return iife143_(_pat_let54_0)
                        return iife142_(pat_let_tv43_)
                    return iife141_(_pat_let53_0)
                rhs273_ = default__.safeDivisionInt(d_1702_v30_, d_1702_v30_)
                rhs274_ = (iife140_(D14_DC38(d_1710_v35_))).cf62
                rhs275_ = default__.safeModuloInt((self).fm15(globalState), d_1702_v30_)
                rhs276_ = (0) - (d_1702_v30_)
                rhs277_ = (d_1663_v1_).fm32((d_1702_v30_) > (d_1702_v30_), d_1713_v38_, 728, len(d_1714_v39_), globalState)
                lhs166_ = d_1711_v36_
                lhs167_ = default__.safeIndex(120, (d_1711_v36_).length(0))
                lhs168_ = self
                r1 = rhs273_
                d_1710_v35_ = rhs274_
                lhs166_[lhs167_] = rhs275_
                r1 = rhs276_
                lhs168_.f7 = rhs277_
                d_1717_v40_: _dafny.Array
                nw297_ = _dafny.Array(None, 8)
                nw297_[int(0)] = True
                nw297_[int(1)] = False
                nw297_[int(2)] = self.f7
                nw297_[int(3)] = self.f7
                nw297_[int(4)] = self.f7
                nw297_[int(5)] = self.f7
                nw297_[int(6)] = self.f7
                nw297_[int(7)] = self.f7
                d_1717_v40_ = nw297_
                d_1718_v41_: _dafny.Set
                d_1718_v41_ = _dafny.Set({d_1717_v40_, d_1717_v40_})
                d_1719_v42_: C2
                nw298_ = C2()
                nw298_.ctor__(d_1718_v41_, d_1664_v2_, self.f7)
                d_1719_v42_ = nw298_
                d_1720_v43_: _dafny.MultiSet
                d_1720_v43_ = _dafny.MultiSet([self.f7])
                d_1721_v44_: D15
                d_1721_v44_ = D15_DC42(_dafny.MultiSet([self.f7]))
                d_1722_v45_: _dafny.Map
                d_1722_v45_ = _dafny.Map({self.f7: d_1662_v0_})
                d_1723_v46_: D0
                d_1723_v46_ = D0_DC0((d_1711_v36_)[default__.safeIndex(120, (d_1711_v36_).length(0))])
                d_1724_v47_: _dafny.Seq
                d_1724_v47_ = _dafny.SeqWithoutIsStrInference([d_1723_v46_])
                d_1725_v48_: _dafny.Seq
                d_1725_v48_ = _dafny.SeqWithoutIsStrInference([d_1720_v43_, (d_1721_v44_).cf69, (d_1720_v43_).set(True, default__.abs(d_1702_v30_)), (default__.fm49(not(self.f7), d_1722_v45_, self.f7, globalState)) - (_dafny.MultiSet([default__.fm1(d_1724_v47_, d_1702_v30_, d_1702_v30_, globalState)])), (d_1720_v43_) - (_dafny.MultiSet([False, self.f7]))])
                r1 = len(d_1725_v48_)
            elif True:
                d_1726_v49_: _dafny.Set
                d_1726_v49_ = _dafny.Set({self.f7})
                d_1727_v50_: D1
                d_1727_v50_ = D1_DC3(d_1726_v49_)
                d_1728_v51_: _dafny.Map
                d_1728_v51_ = _dafny.Map({(_dafny.MultiSet([d_1702_v30_])).cardinality: self.f7})
                d_1729_v52_: _dafny.Seq
                d_1729_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "syvodnk"))
                d_1730_v53_: D3
                d_1730_v53_ = D3_DC11(d_1702_v30_, d_1702_v30_, len(d_1729_v52_), self.f7)
                d_1731_v54_: _dafny.Map
                d_1731_v54_ = _dafny.Map({d_1702_v30_: d_1729_v52_})
                d_1732_v55_: _dafny.Set
                d_1732_v55_ = _dafny.Set({d_1702_v30_, (0) - (d_1702_v30_)})
                d_1733_v56_: _dafny.Array
                nw299_ = _dafny.Array(None, 26)
                nw299_[int(0)] = d_1702_v30_
                nw299_[int(1)] = 694
                nw299_[int(2)] = d_1702_v30_
                nw299_[int(3)] = default__.fm0(globalState)
                nw299_[int(4)] = len((_dafny.Set({self.f7, self.f7})) - ((d_1727_v50_).cf6))
                nw299_[int(5)] = (d_1702_v30_ if True else d_1702_v30_)
                nw299_[int(6)] = default__.safeModuloInt(138, len(d_1726_v49_))
                nw299_[int(7)] = (d_1702_v30_) + (d_1702_v30_)
                nw299_[int(8)] = (d_1702_v30_ if self.f7 else len(d_1662_v0_))
                nw299_[int(9)] = d_1702_v30_
                nw299_[int(10)] = len((d_1728_v51_ if self.f7 else d_1728_v51_))
                nw299_[int(11)] = d_1702_v30_
                nw299_[int(12)] = (919) + (d_1702_v30_)
                nw299_[int(13)] = d_1702_v30_
                nw299_[int(14)] = d_1702_v30_
                nw299_[int(15)] = (d_1730_v53_).cf23
                nw299_[int(16)] = d_1702_v30_
                nw299_[int(17)] = d_1702_v30_
                nw299_[int(18)] = (self).fm15(globalState)
                nw299_[int(19)] = 772
                nw299_[int(20)] = (default__.fm0(globalState) if self.f7 else d_1702_v30_)
                nw299_[int(21)] = d_1702_v30_
                nw299_[int(22)] = d_1702_v30_
                nw299_[int(23)] = d_1702_v30_
                nw299_[int(24)] = len((d_1729_v52_) + (((d_1731_v54_)[(0) - (len(d_1732_v55_))] if ((0) - (len(d_1732_v55_))) in (d_1731_v54_) else d_1729_v52_)))
                nw299_[int(25)] = d_1702_v30_
                d_1733_v56_ = nw299_
                index267_ = default__.safeIndex(947, (d_1733_v56_).length(0))
                (d_1733_v56_)[index267_] = (0) - (d_1702_v30_)
                index268_ = default__.safeIndex(947, (d_1733_v56_).length(0))
                (d_1733_v56_)[index268_] = len((d_1662_v0_) + (d_1662_v0_))
                d_1734_v57_: _dafny.Array
                def lambda101_(d_1735_v0_):
                    def lambda102_(d_1736_i8_):
                        return (_dafny.MultiSet([d_1735_v0_])).ispropersubset(_dafny.MultiSet([d_1735_v0_]))

                    return lambda102_

                init56_ = lambda101_(d_1662_v0_)
                nw300_ = _dafny.Array(None, 10)
                for i0_56_ in range(nw300_.length(0)):
                    nw300_[i0_56_] = init56_(i0_56_)
                d_1734_v57_ = nw300_
                d_1734_v57_ = d_1734_v57_
                d_1737_v58_: _dafny.Array
                nw301_ = _dafny.Array(D13.default()(), 19)
                d_1737_v58_ = nw301_
                d_1738_v60_: _dafny.Map
                def iife144_():
                    coll34_ = _dafny.Set()
                    compr_34_: int
                    for compr_34_ in _dafny.IntegerRange(262, 582):
                        d_1739_v59_: int = compr_34_
                        if ((262) <= (d_1739_v59_)) and ((d_1739_v59_) < (582)):
                            coll34_ = coll34_.union(_dafny.Set([default__.safeDivisionInt(d_1739_v59_, (d_1733_v56_)[default__.safeIndex(947, (d_1733_v56_).length(0))])]))
                    return _dafny.Set(coll34_)
                d_1738_v60_ = _dafny.Map({d_1733_v56_: iife144_()
                })
                d_1740_v61_: D13
                d_1740_v61_ = D13_DC34(d_1738_v60_)
                d_1741_v62_: D13
                d_1741_v62_ = D13_DC37(d_1740_v61_)
                index269_ = default__.safeIndex(875, (d_1737_v58_).length(0))
                (d_1737_v58_)[index269_] = d_1741_v62_
                index270_ = default__.safeIndex(875, (d_1737_v58_).length(0))
                (d_1737_v58_)[index270_] = D13_DC37(d_1740_v61_)
                d_1742_v63_: _dafny.Map
                d_1742_v63_ = _dafny.Map({default__.fm21(self.f7, globalState): (d_1729_v52_) <= (d_1729_v52_)})
                d_1743_v64_: _dafny.Seq
                d_1743_v64_ = _dafny.SeqWithoutIsStrInference([len(d_1729_v52_), (d_1733_v56_)[default__.safeIndex(947, (d_1733_v56_).length(0))], (d_1733_v56_)[default__.safeIndex(947, (d_1733_v56_).length(0))], (d_1733_v56_)[default__.safeIndex(947, (d_1733_v56_).length(0))], 28])
                d_1742_v63_ = (d_1742_v63_) | (_dafny.Map({d_1743_v64_: self.f7}))
                index271_ = default__.safeIndex(947, (d_1733_v56_).length(0))
                (d_1733_v56_)[index271_] = (0) - (len(d_1729_v52_))
            r1 = default__.safeDivisionInt((d_1702_v30_) + (d_1702_v30_), d_1702_v30_)
            d_1744_v65_: D6
            d_1744_v65_ = D6_DC15(d_1664_v2_)
            d_1744_v65_ = d_1744_v65_
            d_1662_v0_ = (d_1662_v0_) + (((d_1662_v0_).set(default__.safeIndex((_dafny.MultiSet(d_1662_v0_)).cardinality, len(d_1662_v0_)), self.f7)) + (_dafny.SeqWithoutIsStrInference([self.f7])))
            d_1745_v66_: _dafny.Array
            nw302_ = _dafny.Array(int(0), 13)
            d_1745_v66_ = nw302_
            d_1746_v67_: D0
            d_1746_v67_ = D0_DC2(not(self.f7), len(d_1662_v0_), _dafny.Set({d_1745_v66_}), d_1702_v30_)
            d_1747_v68_: D0
            d_1747_v68_ = D0_DC2(self.f7, (d_1746_v67_).cf5, _dafny.Set({d_1745_v66_, d_1745_v66_}), default__.fm0(globalState))
            r1 = (d_1746_v67_).cf5
        elif True:
            d_1748_v69_: _dafny.Map
            d_1748_v69_ = _dafny.Map({self.f7: (d_1662_v0_)[default__.safeIndex(d_1702_v30_, len(d_1662_v0_))]})
            d_1749_v70_: _dafny.MultiSet
            d_1749_v70_ = _dafny.MultiSet([len(d_1748_v69_)])
            d_1750_v71_: _dafny.MultiSet
            d_1750_v71_ = _dafny.MultiSet([d_1664_v2_])
            d_1751_v72_: _dafny.Map
            d_1751_v72_ = _dafny.Map({(_dafny.MultiSet([self.f7])).set(self.f7, default__.abs(d_1702_v30_)): self.f7})
            d_1752_v74_: _dafny.Map
            def iife145_():
                coll35_ = _dafny.Set()
                compr_35_: _dafny.MultiSet
                for compr_35_ in (d_1751_v72_).keys.Elements:
                    d_1753_v73_: _dafny.MultiSet = compr_35_
                    if (d_1753_v73_) in (d_1751_v72_):
                        coll35_ = coll35_.union(_dafny.Set([d_1753_v73_]))
                return _dafny.Set(coll35_)
            def iife146_():
                coll36_ = _dafny.Set()
                compr_36_: _dafny.MultiSet
                for compr_36_ in (d_1751_v72_).keys.Elements:
                    d_1754_v73_: _dafny.MultiSet = compr_36_
                    if (d_1754_v73_) in (d_1751_v72_):
                        coll36_ = coll36_.union(_dafny.Set([d_1754_v73_]))
                return _dafny.Set(coll36_)
            d_1752_v74_ = _dafny.Map({self.f7: ((d_1749_v70_)[((d_1750_v71_)[d_1664_v2_] if (d_1664_v2_) in (d_1750_v71_) else len(iife145_()
            ))] if (((d_1750_v71_)[d_1664_v2_] if (d_1664_v2_) in (d_1750_v71_) else len(iife146_()
            ))) in (d_1749_v70_) else d_1702_v30_)})
            d_1752_v74_ = (d_1752_v74_).set(self.f7, d_1702_v30_)
            d_1662_v0_ = (d_1662_v0_) + (d_1662_v0_)
            d_1755_v75_: _dafny.Map
            d_1755_v75_ = _dafny.Map({d_1662_v0_: d_1702_v30_})
            d_1755_v75_ = (d_1755_v75_) | ((d_1755_v75_) | (d_1755_v75_))
            d_1756_v76_: _dafny.MultiSet
            d_1756_v76_ = _dafny.MultiSet([self.f7, self.f7, self.f7, self.f7, self.f7])
            d_1757_v77_: _dafny.Seq
            d_1757_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snip"))
            d_1702_v30_ = default__.safeDivisionInt(((d_1756_v76_)[self.f7] if (self.f7) in (d_1756_v76_) else d_1702_v30_), len(d_1757_v77_))
            d_1702_v30_ = ((self).fm15(globalState)) - ((len(d_1757_v77_)) - (d_1702_v30_))
        d_1758_v78_: _dafny.Seq
        d_1758_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbunisxop"))
        d_1759_v79_: C3
        nw303_ = C3()
        nw303_.ctor__(d_1702_v30_, not(self.f7))
        d_1759_v79_ = nw303_
        d_1760_v80_: _dafny.Map
        d_1760_v80_ = _dafny.Map({self.f7: d_1759_v79_})
        d_1761_v81_: _dafny.MultiSet
        d_1761_v81_ = _dafny.MultiSet([self.f7])
        d_1762_v82_: D0
        d_1762_v82_ = D0_DC0((d_1761_v81_).cardinality)
        d_1763_v83_: _dafny.Seq
        d_1763_v83_ = _dafny.SeqWithoutIsStrInference([d_1762_v82_, d_1762_v82_, d_1762_v82_])
        d_1764_v84_: _dafny.Map
        d_1764_v84_ = _dafny.Map({False: _dafny.CodePoint('n')})
        if default__.fm1(default__.fm51(d_1664_v2_, self.f7, globalState), len((d_1758_v78_) + (default__.fm35(len(_dafny.SeqWithoutIsStrInference([d_1759_v79_, ((d_1760_v80_)[default__.fm1(d_1763_v83_, d_1702_v30_, d_1702_v30_, globalState)] if (default__.fm1(d_1763_v83_, d_1702_v30_, d_1702_v30_, globalState)) in (d_1760_v80_) else d_1759_v79_), d_1759_v79_, d_1759_v79_])), len(d_1764_v84_), globalState))), (len(d_1758_v78_)) - ((0) - (d_1702_v30_)), globalState):
            (self).f7 = not (not(self.f7)) or (self.f7)
            d_1702_v30_ = default__.safeModuloInt(d_1702_v30_, len(_dafny.SeqWithoutIsStrInference([d_1664_v2_ for d_1765_i9_ in range(default__.abs(-723))])))
            if (default__.fm30(not(self.f7), (d_1759_v79_).f14, False, globalState)) == (d_1665_v3_):
                d_1766_v85_: _dafny.Array
                nw304_ = _dafny.Array(int(0), 7)
                d_1766_v85_ = nw304_
                d_1767_v86_: D9
                d_1767_v86_ = D9_DC23(d_1766_v85_)
                rhs278_ = ((0) - ((d_1759_v79_).f14)) - ((d_1759_v79_).f14)
                rhs279_ = d_1767_v86_
                r1 = rhs278_
                d_1767_v86_ = rhs279_
                d_1702_v30_ = default__.safeDivisionInt((d_1759_v79_).f14, ((d_1759_v79_).f14) - ((d_1759_v79_).fm9(self.f7, d_1702_v30_, d_1664_v2_, globalState)))
                d_1768_v87_: D6
                d_1768_v87_ = D6_DC17(not(self.f7), (d_1759_v79_).f14, default__.fm26(self.f7, d_1702_v30_, (d_1759_v79_).f14, globalState), d_1664_v2_)
                d_1769_v88_: D6
                d_1769_v88_ = D6_DC18(d_1768_v87_)
                d_1770_v89_: D14
                d_1770_v89_ = D14_DC40(d_1769_v88_, self.f7, self.f7)
                d_1770_v89_ = d_1770_v89_
                d_1771_v90_: _dafny.Array
                def lambda103_(d_1772_v79_):
                    def lambda104_(d_1773_i10_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.Map({False: (d_1772_v79_).f14}), _dafny.Map({not(self.f7): 676}), _dafny.Map({self.f7: (_dafny.MultiSet([_dafny.Map({self.f7: self.f7}), _dafny.Map({self.f7: self.f7})])).cardinality})])

                    return lambda104_

                init57_ = lambda103_(d_1759_v79_)
                nw305_ = _dafny.Array(None, 12)
                for i0_57_ in range(nw305_.length(0)):
                    nw305_[i0_57_] = init57_(i0_57_)
                d_1771_v90_ = nw305_
                d_1774_v91_: _dafny.Map
                d_1774_v91_ = _dafny.Map({self.f7: (0) - ((d_1759_v79_).f14)})
                d_1775_v92_: _dafny.Seq
                d_1775_v92_ = _dafny.SeqWithoutIsStrInference([d_1774_v91_, d_1774_v91_, d_1774_v91_])
                index272_ = default__.safeIndex(950, (d_1771_v90_).length(0))
                (d_1771_v90_)[index272_] = d_1775_v92_
                d_1776_v93_: _dafny.Map
                d_1776_v93_ = _dafny.Map({len(d_1774_v91_): d_1662_v0_})
                d_1777_v94_: _dafny.Seq
                d_1777_v94_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f7, self.f7]), ((d_1776_v93_)[(d_1759_v79_).f14] if ((d_1759_v79_).f14) in (d_1776_v93_) else d_1662_v0_), d_1662_v0_])
                d_1778_v95_: _dafny.Map
                d_1778_v95_ = _dafny.Map({(d_1662_v0_)[default__.safeIndex((0) - (((d_1761_v81_)[self.f7] if (self.f7) in (d_1761_v81_) else (d_1761_v81_).cardinality)), len(d_1662_v0_))]: d_1777_v94_})
                d_1779_v96_: _dafny.Seq
                d_1779_v96_ = _dafny.SeqWithoutIsStrInference([d_1777_v94_])
                index273_ = default__.safeIndex(950, (d_1771_v90_).length(0))
                rhs280_ = d_1775_v92_
                rhs281_ = ((d_1778_v95_)[(d_1759_v79_).fm8((d_1759_v79_).f14, globalState)] if ((d_1759_v79_).fm8((d_1759_v79_).f14, globalState)) in (d_1778_v95_) else ((d_1779_v96_)[default__.safeIndex((d_1759_v79_).f14, len(d_1779_v96_))] if self.f7 else d_1777_v94_))
                lhs169_ = d_1771_v90_
                lhs170_ = default__.safeIndex(950, (d_1771_v90_).length(0))
                lhs169_[lhs170_] = rhs280_
                d_1777_v94_ = rhs281_
                d_1780_v97_: _dafny.Map
                d_1780_v97_ = _dafny.Map({self.f7: self.f7})
                d_1780_v97_ = d_1780_v97_
            elif True:
                d_1758_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kefhqejr"))
                d_1781_v98_: D12
                d_1781_v98_ = D12_DC32(self.f7)
                d_1782_v99_: _dafny.Map
                d_1782_v99_ = _dafny.Map({(d_1759_v79_).f14: self.f7})
                d_1783_v100_: _dafny.MultiSet
                d_1783_v100_ = _dafny.MultiSet([(d_1759_v79_).f14, len(d_1782_v99_)])
                d_1784_v101_: _dafny.Seq
                d_1784_v101_ = _dafny.SeqWithoutIsStrInference([d_1783_v100_])
                d_1785_v102_: _dafny.Seq
                d_1785_v102_ = _dafny.SeqWithoutIsStrInference([(d_1759_v79_).f14])
                d_1786_v103_: _dafny.Map
                d_1786_v103_ = _dafny.Map({(d_1759_v79_).f14: (d_1759_v79_).f14})
                d_1787_v104_: _dafny.Map
                d_1787_v104_ = _dafny.Map({d_1702_v30_: d_1786_v103_})
                d_1788_v105_: _dafny.Array
                nw306_ = _dafny.Array(False, 9)
                d_1788_v105_ = nw306_
                d_1789_v106_: _dafny.MultiSet
                d_1789_v106_ = _dafny.MultiSet([d_1788_v105_])
                d_1781_v98_ = default__.fm52((d_1784_v101_) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1785_v102_), d_1783_v100_, d_1783_v100_, (d_1783_v100_).set(len(((d_1787_v104_)[(d_1789_v106_).cardinality] if ((d_1789_v106_).cardinality) in (d_1787_v104_) else d_1786_v103_)), default__.abs((d_1759_v79_).f14)), _dafny.MultiSet(d_1785_v102_)])), d_1702_v30_, default__.safeModuloInt((d_1759_v79_).f14, d_1702_v30_), self.f7, globalState)
                d_1790_v107_: _dafny.Set
                d_1790_v107_ = _dafny.Set({d_1788_v105_, d_1788_v105_})
                d_1791_v108_: D9
                d_1791_v108_ = D9_DC25()
                d_1792_v109_: C2
                nw307_ = C2()
                nw307_.ctor__(d_1790_v107_, default__.fm37(self.f7, d_1791_v108_, globalState), (255) >= (826))
                d_1792_v109_ = nw307_
                r1 = default__.safeModuloInt((d_1759_v79_).f14, (d_1702_v30_) * ((d_1759_v79_).f14))
                d_1793_v110_: C0
                nw308_ = C0()
                nw308_.ctor__(self.f7, False)
                d_1793_v110_ = nw308_
            d_1794_v111_: _dafny.Seq
            d_1794_v111_ = _dafny.SeqWithoutIsStrInference([(d_1759_v79_).f14, (d_1759_v79_).f14])
            d_1795_v112_: D9
            d_1795_v112_ = D9_DC26(d_1794_v111_)
            d_1796_v113_: _dafny.Array
            nw309_ = _dafny.Array(None, 10)
            nw309_[int(0)] = d_1795_v112_
            nw309_[int(1)] = d_1795_v112_
            nw309_[int(2)] = d_1795_v112_
            nw309_[int(3)] = d_1795_v112_
            nw309_[int(4)] = d_1795_v112_
            nw309_[int(5)] = d_1795_v112_
            nw309_[int(6)] = (D9_DC26(d_1794_v111_) if self.f7 else d_1795_v112_)
            nw309_[int(7)] = d_1795_v112_
            nw309_[int(8)] = d_1795_v112_
            nw309_[int(9)] = d_1795_v112_
            d_1796_v113_ = nw309_
            index274_ = default__.safeIndex(525, (d_1796_v113_).length(0))
            (d_1796_v113_)[index274_] = d_1795_v112_
            index275_ = default__.safeIndex(525, (d_1796_v113_).length(0))
            (d_1796_v113_)[index275_] = D9_DC26(d_1794_v111_)
            (self).f7 = self.f7
        elif True:
            d_1797_v114_: _dafny.Set
            d_1797_v114_ = _dafny.Set({d_1702_v30_})
            d_1702_v30_ = len((d_1797_v114_) - (d_1797_v114_))
            if (_dafny.MultiSet(d_1662_v0_)).isdisjoint(d_1761_v81_):
                d_1702_v30_ = (d_1759_v79_).f14
                (self).f7 = (674) < ((d_1759_v79_).f14)
                d_1798_v115_: _dafny.Array
                nw310_ = _dafny.Array(False, 27)
                d_1798_v115_ = nw310_
                index276_ = default__.safeIndex(578, (d_1798_v115_).length(0))
                (d_1798_v115_)[index276_] = not (self.f7) or (True)
                index277_ = default__.safeIndex(578, (d_1798_v115_).length(0))
                (d_1798_v115_)[index277_] = True
                d_1798_v115_ = d_1798_v115_
                d_1799_v116_: _dafny.Map
                d_1799_v116_ = _dafny.Map({d_1664_v2_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))})
                d_1800_v117_: D14
                d_1800_v117_ = D14_DC38(d_1799_v116_)
                d_1801_v118_: _dafny.Map
                d_1801_v118_ = _dafny.Map({(d_1759_v79_).f14: d_1800_v117_})
                d_1802_v119_: D14
                d_1802_v119_ = D14_DC41(((d_1801_v118_)[d_1702_v30_] if (d_1702_v30_) in (d_1801_v118_) else d_1800_v117_))
                d_1802_v119_ = d_1802_v119_
            elif True:
                (self).f7 = self.f7
                (self).f7 = self.f7
                d_1803_v120_: _dafny.Array
                nw311_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1803_v120_ = nw311_
                d_1804_v121_: _dafny.Array
                nw312_ = _dafny.Array(_dafny.CodePoint('D'), 19)
                d_1804_v121_ = nw312_
                index278_ = default__.safeIndex(485, (d_1803_v120_).length(0))
                (d_1803_v120_)[index278_] = (d_1804_v121_ if self.f7 else d_1804_v121_)
                index279_ = default__.safeIndex(485, (d_1803_v120_).length(0))
                nw313_ = _dafny.Array(None, 9)
                nw313_[int(0)] = d_1664_v2_
                nw313_[int(1)] = d_1664_v2_
                nw313_[int(2)] = _dafny.CodePoint('t')
                nw313_[int(3)] = d_1664_v2_
                nw313_[int(4)] = d_1664_v2_
                nw313_[int(5)] = d_1664_v2_
                nw313_[int(6)] = ((d_1764_v84_)[not(self.f7)] if (not(self.f7)) in (d_1764_v84_) else d_1664_v2_)
                nw313_[int(7)] = _dafny.CodePoint('n')
                nw313_[int(8)] = d_1664_v2_
                (d_1803_v120_)[index279_] = nw313_
                d_1758_v78_ = d_1758_v78_
                d_1805_v122_: _dafny.MultiSet
                d_1805_v122_ = _dafny.MultiSet([209])
                d_1806_v123_: _dafny.Map
                d_1806_v123_ = _dafny.Map({not(self.f7): self.f7})
                d_1807_v124_: _dafny.Array
                nw314_ = _dafny.Array(None, 21)
                nw314_[int(0)] = self.f7
                nw314_[int(1)] = (_dafny.MultiSet([(d_1759_v79_).f14, (d_1759_v79_).f14])).ispropersubset(d_1805_v122_)
                nw314_[int(2)] = self.f7
                nw314_[int(3)] = True
                nw314_[int(4)] = self.f7
                nw314_[int(5)] = (self.f7) == (False)
                nw314_[int(6)] = self.f7
                nw314_[int(7)] = not(((d_1806_v123_)[self.f7] if (self.f7) in (d_1806_v123_) else self.f7))
                nw314_[int(8)] = (not(self.f7)) and (self.f7)
                nw314_[int(9)] = self.f7
                nw314_[int(10)] = self.f7
                nw314_[int(11)] = self.f7
                nw314_[int(12)] = self.f7
                nw314_[int(13)] = (not(self.f7) if self.f7 else self.f7)
                nw314_[int(14)] = ((d_1759_v79_).f14) < ((d_1759_v79_).f14)
                nw314_[int(15)] = ((0) - (d_1702_v30_)) < ((d_1759_v79_).f14)
                nw314_[int(16)] = self.f7
                nw314_[int(17)] = self.f7
                nw314_[int(18)] = self.f7
                nw314_[int(19)] = (d_1662_v0_)[default__.safeIndex((0) - ((0) - (default__.fm0(globalState))), len(d_1662_v0_))]
                nw314_[int(20)] = self.f7
                d_1807_v124_ = nw314_
                index280_ = default__.safeIndex(433, (d_1807_v124_).length(0))
                (d_1807_v124_)[index280_] = self.f7
                index281_ = default__.safeIndex(433, (d_1807_v124_).length(0))
                rhs282_ = (d_1758_v78_)[default__.safeIndex(default__.safeDivisionInt(len(d_1797_v114_), (d_1805_v122_).cardinality), len(d_1758_v78_))]
                rhs283_ = self.f7
                lhs171_ = d_1807_v124_
                lhs172_ = default__.safeIndex(433, (d_1807_v124_).length(0))
                r0 = rhs282_
                lhs171_[lhs172_] = rhs283_
            d_1808_v125_: _dafny.Array
            nw315_ = _dafny.Array(int(0), 25)
            d_1808_v125_ = nw315_
            if (d_1808_v125_) in (_dafny.SeqWithoutIsStrInference([d_1808_v125_, d_1808_v125_, d_1808_v125_])):
                index282_ = default__.safeIndex(219, (d_1808_v125_).length(0))
                (d_1808_v125_)[index282_] = (d_1759_v79_).f14
                d_1809_v126_: _dafny.Array
                nw316_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_1809_v126_ = nw316_
                d_1810_v127_: _dafny.Array
                def lambda105_(d_1811_i11_):
                    return self.f7

                init58_ = lambda105_
                nw317_ = _dafny.Array(None, 13)
                for i0_58_ in range(nw317_.length(0)):
                    nw317_[i0_58_] = init58_(i0_58_)
                d_1810_v127_ = nw317_
                index283_ = default__.safeIndex(841, (d_1809_v126_).length(0))
                (d_1809_v126_)[index283_] = d_1810_v127_
                d_1812_v128_: _dafny.MultiSet
                d_1812_v128_ = _dafny.MultiSet([d_1702_v30_])
                index284_ = default__.safeIndex(219, (d_1808_v125_).length(0))
                index285_ = default__.safeIndex(841, (d_1809_v126_).length(0))
                rhs284_ = ((d_1812_v128_)[(d_1759_v79_).fm9(self.f7, (0) - (d_1702_v30_), d_1664_v2_, globalState)] if ((d_1759_v79_).fm9(self.f7, (0) - (d_1702_v30_), d_1664_v2_, globalState)) in (d_1812_v128_) else default__.fm0(globalState))
                rhs285_ = self.f7
                rhs286_ = (d_1759_v79_).f14
                rhs287_ = d_1810_v127_
                lhs173_ = self
                lhs174_ = d_1808_v125_
                lhs175_ = default__.safeIndex(219, (d_1808_v125_).length(0))
                lhs176_ = d_1809_v126_
                lhs177_ = default__.safeIndex(841, (d_1809_v126_).length(0))
                d_1702_v30_ = rhs284_
                lhs173_.f7 = rhs285_
                lhs174_[lhs175_] = rhs286_
                lhs176_[lhs177_] = rhs287_
                d_1702_v30_ = (d_1759_v79_).f14
                d_1813_v129_: _dafny.Array
                def lambda106_(d_1814_v79_):
                    def lambda107_(d_1815_i12_):
                        return (d_1815_i12_) - ((d_1814_v79_).f14)

                    return lambda107_

                init59_ = lambda106_(d_1759_v79_)
                nw318_ = _dafny.Array(None, 7)
                for i0_59_ in range(nw318_.length(0)):
                    nw318_[i0_59_] = init59_(i0_59_)
                d_1813_v129_ = nw318_
                d_1813_v129_ = d_1813_v129_
                d_1816_v130_: D6
                d_1816_v130_ = D6_DC17(self.f7, (self).fm15(globalState), d_1664_v2_, d_1664_v2_)
                d_1817_v131_: D6
                d_1817_v131_ = D6_DC18(d_1816_v130_)
                d_1818_v132_: D6
                d_1818_v132_ = D6_DC18(d_1816_v130_)
                d_1819_v133_: D14
                d_1819_v133_ = D14_DC40(d_1818_v132_, False, self.f7)
                d_1820_v134_: _dafny.Array
                def lambda108_(d_1821_v2_):
                    def lambda109_(d_1822_i13_):
                        return d_1821_v2_

                    return lambda109_

                init60_ = lambda108_(d_1664_v2_)
                nw319_ = _dafny.Array(None, 5)
                for i0_60_ in range(nw319_.length(0)):
                    nw319_[i0_60_] = init60_(i0_60_)
                d_1820_v134_ = nw319_
                d_1823_v135_: _dafny.MultiSet
                d_1823_v135_ = _dafny.MultiSet([d_1820_v134_, d_1820_v134_])
                r1 = ((d_1759_v79_).f14 if (d_1819_v133_).cf67 else ((d_1823_v135_)[d_1820_v134_] if (d_1820_v134_) in (d_1823_v135_) else (d_1759_v79_).f14))
                d_1824_v136_: _dafny.Array
                nw320_ = _dafny.Array(_dafny.Seq({}), 15)
                d_1824_v136_ = nw320_
                d_1824_v136_ = d_1824_v136_
            elif True:
                index286_ = default__.safeIndex(92, (d_1808_v125_).length(0))
                (d_1808_v125_)[index286_] = (d_1759_v79_).f14
                index287_ = default__.safeIndex(92, (d_1808_v125_).length(0))
                (d_1808_v125_)[index287_] = default__.safeDivisionInt(((d_1759_v79_).f14) * (d_1702_v30_), ((d_1759_v79_).f14) - (239))
                d_1825_v137_: _dafny.Array
                def lambda110_(d_1826_i14_):
                    return self.f7

                init61_ = lambda110_
                nw321_ = _dafny.Array(None, 10)
                for i0_61_ in range(nw321_.length(0)):
                    nw321_[i0_61_] = init61_(i0_61_)
                d_1825_v137_ = nw321_
                index288_ = default__.safeIndex(831, (d_1825_v137_).length(0))
                (d_1825_v137_)[index288_] = self.f7
                index289_ = default__.safeIndex(831, (d_1825_v137_).length(0))
                rhs288_ = (self.f7) and (self.f7)
                rhs289_ = not (False) or (self.f7)
                rhs290_ = 690
                lhs178_ = d_1825_v137_
                lhs179_ = default__.safeIndex(831, (d_1825_v137_).length(0))
                lhs180_ = self
                lhs178_[lhs179_] = rhs288_
                lhs180_.f7 = rhs289_
                d_1702_v30_ = rhs290_
                index290_ = default__.safeIndex(831, (d_1825_v137_).length(0))
                (d_1825_v137_)[index290_] = True
                d_1827_v138_: C1
                nw322_ = C1()
                nw322_.ctor__(self.f7)
                d_1827_v138_ = nw322_
                rhs291_ = default__.safeModuloInt((d_1759_v79_).f14, (d_1808_v125_)[default__.safeIndex(92, (d_1808_v125_).length(0))])
                rhs292_ = default__.safeModuloInt((0) - ((d_1759_v79_).f14), (d_1808_v125_)[default__.safeIndex(92, (d_1808_v125_).length(0))])
                rhs293_ = (d_1759_v79_).f14
                r1 = rhs291_
                d_1702_v30_ = rhs292_
                d_1702_v30_ = rhs293_
            d_1828_v139_: _dafny.Array
            nw323_ = _dafny.Array(False, 14)
            d_1828_v139_ = nw323_
            d_1829_v140_: _dafny.Map
            d_1829_v140_ = _dafny.Map({self.f7: (d_1759_v79_).f14})
            d_1830_v141_: _dafny.Map
            d_1830_v141_ = _dafny.Map({len(d_1829_v140_): self.f7})
            d_1831_v142_: _dafny.Map
            d_1831_v142_ = _dafny.Map({d_1830_v141_: (d_1828_v139_ if self.f7 else d_1828_v139_)})
            d_1828_v139_ = ((d_1831_v142_)[d_1830_v141_] if (d_1830_v141_) in (d_1831_v142_) else d_1828_v139_)
            (self).f7 = ((d_1759_v79_).fm8(d_1702_v30_, globalState) if self.f7 else self.f7)
        d_1832_v143_: _dafny.Array
        nw324_ = _dafny.Array(False, 3)
        d_1832_v143_ = nw324_
        index291_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        (d_1832_v143_)[index291_] = not(self.f7)
        d_1833_v144_: D14
        d_1833_v144_ = D14_DC38(_dafny.Map({_dafny.CodePoint('e'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "balalo"))}))
        d_1834_v145_: D15
        d_1834_v145_ = D15_DC43((d_1759_v79_).f14)
        pat_let_tv44_ = d_1664_v2_
        pat_let_tv45_ = d_1758_v78_
        pat_let_tv46_ = d_1702_v30_
        index292_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        def lambda111_(source13_):
            if source13_.is_DC39:
                d_1835___mcc_h1_ = source13_.cf63
                d_1836___mcc_h2_ = source13_.cf64
                d_1837_cf64_ = d_1836___mcc_h2_
                d_1838_cf63_ = d_1835___mcc_h1_
                if self.f7:
                    return self.f7
                elif True:
                    return self.f7
            elif source13_.is_DC40:
                d_1839___mcc_h3_ = source13_.cf65
                d_1840___mcc_h4_ = source13_.cf66
                d_1841___mcc_h5_ = source13_.cf67
                d_1842_cf67_ = d_1841___mcc_h5_
                d_1843_cf66_ = d_1840___mcc_h4_
                d_1844_cf65_ = d_1839___mcc_h3_
                return d_1842_cf67_
            elif source13_.is_DC38:
                d_1845___mcc_h6_ = source13_.cf62
                d_1846_cf62_ = d_1845___mcc_h6_
                return self.f7
            elif True:
                d_1847___mcc_h7_ = source13_.cf68
                d_1848_cf68_ = d_1847___mcc_h7_
                return self.f7

        def lambda112_(source14_):
            if source14_.is_DC43:
                d_1849___mcc_h8_ = source14_.cf70
                d_1850_cf70_ = d_1849___mcc_h8_
                return False
            elif source14_.is_DC42:
                d_1851___mcc_h9_ = source14_.cf69
                d_1852_cf69_ = d_1851___mcc_h9_
                if self.f7:
                    return self.f7
                elif True:
                    return self.f7
            elif True:
                d_1853___mcc_h10_ = source14_.cf71
                d_1854_cf71_ = d_1853___mcc_h10_
                return not((pat_let_tv44_) not in (pat_let_tv45_))

        def iife147_(_pat_let55_0):
            def iife148_(d_1855_dt__update__tmp_h1_):
                def iife149_(_pat_let56_0):
                    def iife150_(d_1856_dt__update_hcf70_h0_):
                        return D15_DC43(d_1856_dt__update_hcf70_h0_)
                    return iife150_(_pat_let56_0)
                return iife149_(pat_let_tv46_)
            return iife148_(_pat_let55_0)
        rhs294_ = not(self.f7)
        rhs295_ = lambda111_(d_1833_v144_)
        rhs296_ = 688
        rhs297_ = lambda112_(iife147_(d_1834_v145_))
        lhs181_ = self
        lhs182_ = self
        lhs183_ = d_1832_v143_
        lhs184_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        lhs181_.f7 = rhs294_
        lhs182_.f7 = rhs295_
        r1 = rhs296_
        lhs183_[lhs184_] = rhs297_
        index293_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        index294_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        index295_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        def iife151_():
            coll37_ = _dafny.Set()
            compr_37_: int
            for compr_37_ in _dafny.IntegerRange(258, 576):
                d_1857_v146_: int = compr_37_
                if ((258) <= (d_1857_v146_)) and ((d_1857_v146_) < (576)):
                    coll37_ = coll37_.union(_dafny.Set([default__.safeDivisionInt(d_1857_v146_, (d_1759_v79_).f14)]))
            return _dafny.Set(coll37_)
        rhs298_ = ((d_1702_v30_) * (948)) - (d_1702_v30_)
        rhs299_ = (d_1832_v143_)[default__.safeIndex(392, (d_1832_v143_).length(0))]
        rhs300_ = (d_1663_v1_).fm32(not((d_1832_v143_)[default__.safeIndex(392, (d_1832_v143_).length(0))]), iife151_()
        , 831, d_1702_v30_, globalState)
        rhs301_ = (0) - (default__.safeDivisionInt((d_1834_v145_).cf70, d_1702_v30_))
        rhs302_ = self.f7
        lhs185_ = d_1832_v143_
        lhs186_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        lhs187_ = d_1832_v143_
        lhs188_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        lhs189_ = d_1832_v143_
        lhs190_ = default__.safeIndex(392, (d_1832_v143_).length(0))
        d_1702_v30_ = rhs298_
        lhs185_[lhs186_] = rhs299_
        lhs187_[lhs188_] = rhs300_
        r1 = rhs301_
        lhs189_[lhs190_] = rhs302_
        r0 = default__.fm18(globalState)
        r1 = d_1702_v30_
        d_1858_v147_: _dafny.Map
        d_1858_v147_ = _dafny.Map({self.f7: d_1758_v78_})
        r2 = (d_1858_v147_).set(False, d_1758_v78_)
        return r0, r1, r2

    def m10(self, p0, globalState):
        d_1859_v0_: int
        d_1859_v0_ = 921
        d_1860_v1_: _dafny.Map
        d_1860_v1_ = _dafny.Map({p0: self.f7})
        d_1859_v0_ = (0) - ((default__.safeDivisionInt(p0, d_1859_v0_)) - (len(d_1860_v1_)))
        (self).f7 = True
        d_1861_v2_: str
        d_1861_v2_ = _dafny.CodePoint('t')
        d_1862_v3_: _dafny.Seq
        d_1862_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oislpn"))
        d_1863_i0_: int
        d_1863_i0_ = 0
        with _dafny.label("9"):
            while ((d_1861_v2_ if self.f7 else default__.fm18(globalState))) in (d_1862_v3_):
                with _dafny.c_label("9"):
                    if (d_1863_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1863_i0_ = (d_1863_i0_) + (1)
                    d_1864_v4_: _dafny.Map
                    d_1864_v4_ = _dafny.Map({d_1862_v3_: self.f7})
                    d_1864_v4_ = (d_1864_v4_).set((d_1862_v3_) + (d_1862_v3_), self.f7)
                    (self).f7 = self.f7
                    d_1865_v5_: _dafny.Array
                    nw325_ = _dafny.Array(int(0), 28)
                    d_1865_v5_ = nw325_
                    d_1866_v6_: D0
                    d_1866_v6_ = D0_DC0(512)
                    d_1867_v7_: _dafny.MultiSet
                    d_1867_v7_ = _dafny.MultiSet([self.f7, self.f7, default__.fm1(_dafny.SeqWithoutIsStrInference([d_1866_v6_, d_1866_v6_, d_1866_v6_]), d_1859_v0_, d_1859_v0_, globalState)])
                    index296_ = default__.safeIndex(793, (d_1865_v5_).length(0))
                    (d_1865_v5_)[index296_] = (d_1867_v7_).cardinality
                    d_1868_v8_: _dafny.Array
                    def lambda113_(d_1869_i1_):
                        return self.f7

                    init62_ = lambda113_
                    nw326_ = _dafny.Array(None, 17)
                    for i0_62_ in range(nw326_.length(0)):
                        nw326_[i0_62_] = init62_(i0_62_)
                    d_1868_v8_ = nw326_
                    d_1870_v9_: _dafny.MultiSet
                    d_1870_v9_ = _dafny.MultiSet([d_1868_v8_, d_1868_v8_])
                    index297_ = default__.safeIndex(793, (d_1865_v5_).length(0))
                    (d_1865_v5_)[index297_] = ((d_1870_v9_) - ((d_1870_v9_).set(d_1868_v8_, default__.abs(p0)))).cardinality
                    d_1871_v10_: _dafny.MultiSet
                    d_1871_v10_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(d_1862_v3_) for d_1872_i2_ in range(default__.abs(644))]))])
                    index298_ = default__.safeIndex(793, (d_1865_v5_).length(0))
                    (d_1865_v5_)[index298_] = ((d_1871_v10_)[p0] if (p0) in (d_1871_v10_) else len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnw")))]) for d_1873_i3_ in range(default__.abs(474))])))
                    pass
            pass
        d_1874_v11_: _dafny.Set
        d_1874_v11_ = _dafny.Set({d_1861_v2_, d_1861_v2_})
        d_1875_v12_: _dafny.Map
        d_1875_v12_ = _dafny.Map({self.f7: d_1862_v3_})
        d_1876_v13_: _dafny.Map
        d_1876_v13_ = _dafny.Map({self.f7: True})
        d_1877_v15_: _dafny.Set
        d_1877_v15_ = _dafny.Set({d_1861_v2_})
        def iife152_():
            coll38_ = _dafny.Set()
            compr_38_: str
            for compr_38_ in (((d_1875_v12_)[((d_1876_v13_)[self.f7] if (self.f7) in (d_1876_v13_) else self.f7)] if (((d_1876_v13_)[self.f7] if (self.f7) in (d_1876_v13_) else self.f7)) in (d_1875_v12_) else d_1862_v3_)).Elements:
                d_1878_v14_: str = compr_38_
                if (d_1878_v14_) in (((d_1875_v12_)[((d_1876_v13_)[self.f7] if (self.f7) in (d_1876_v13_) else self.f7)] if (((d_1876_v13_)[self.f7] if (self.f7) in (d_1876_v13_) else self.f7)) in (d_1875_v12_) else d_1862_v3_)):
                    coll38_ = coll38_.union(_dafny.Set([d_1878_v14_]))
            return _dafny.Set(coll38_)
        source15_ = (iife152_()
        ) - (d_1877_v15_)
        d_1879___mcc_h0_ = source15_
        d_1880_cf47_ = d_1879___mcc_h0_
        d_1862_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
        d_1881_v16_: D1
        d_1881_v16_ = D1_DC3(_dafny.Set({self.f7}))
        d_1882_v17_: D0
        d_1882_v17_ = D0_DC0(-529)
        d_1883_v18_: D8
        d_1883_v18_ = D8_DC22(p0, d_1881_v16_, d_1882_v17_)
        d_1884_v19_: _dafny.Seq
        d_1884_v19_ = _dafny.SeqWithoutIsStrInference([d_1883_v18_, d_1883_v18_, d_1883_v18_])
        d_1884_v19_ = d_1884_v19_
        (self).f7 = self.f7
        d_1859_v0_ = p0
        if (self.f7) and (self.f7):
            d_1885_v21_: C0
            nw327_ = C0()
            def iife153_():
                coll39_ = _dafny.Set()
                compr_39_: str
                for compr_39_ in (d_1862_v3_).Elements:
                    d_1886_v20_: str = compr_39_
                    if (d_1886_v20_) in (d_1862_v3_):
                        coll39_ = coll39_.union(_dafny.Set([d_1886_v20_]))
                return _dafny.Set(coll39_)
            nw327_.ctor__((_dafny.Set({d_1861_v2_, d_1861_v2_})).ispropersubset(iife153_()
            ), self.f7)
            d_1885_v21_ = nw327_
            d_1887_v22_: D6
            d_1887_v22_ = D6_DC16(False)
            if (d_1887_v22_).cf30:
                d_1888_v23_: _dafny.Seq
                d_1888_v23_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1889_v24_: _dafny.Set
                d_1889_v24_ = _dafny.Set({d_1888_v23_})
                d_1890_v25_: _dafny.Map
                d_1890_v25_ = _dafny.Map({False: d_1889_v24_})
                d_1891_v26_: D0
                d_1891_v26_ = D0_DC0(p0)
                d_1892_v27_: _dafny.Seq
                d_1892_v27_ = _dafny.SeqWithoutIsStrInference([d_1891_v26_, d_1891_v26_])
                d_1893_v28_: _dafny.MultiSet
                d_1893_v28_ = _dafny.MultiSet([d_1885_v21_.f13, d_1885_v21_.f13])
                d_1894_v29_: _dafny.Set
                d_1894_v29_ = _dafny.Set({not((d_1885_v21_).f12)})
                d_1895_v30_: D1
                d_1895_v30_ = D1_DC3(d_1894_v29_)
                d_1896_v31_: D8
                d_1896_v31_ = D8_DC22(d_1859_v0_, d_1895_v30_, d_1891_v26_)
                rhs303_ = ((d_1889_v24_ if ((d_1876_v13_)[(d_1885_v21_).f12] if ((d_1885_v21_).f12) in (d_1876_v13_) else True) else ((d_1890_v25_)[d_1885_v21_.f13] if (d_1885_v21_.f13) in (d_1890_v25_) else d_1889_v24_))).ispropersubset((_dafny.Set({d_1888_v23_}) if d_1885_v21_.f13 else d_1889_v24_))
                rhs304_ = default__.fm1(d_1892_v27_, (d_1893_v28_).cardinality, (d_1896_v31_).cf41, globalState)
                rhs305_ = len((d_1888_v23_) + (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (p0)) for d_1897_i4_ in range(default__.abs(933))])))
                lhs191_ = d_1885_v21_
                lhs192_ = d_1885_v21_
                lhs191_.f13 = rhs303_
                lhs192_.f13 = rhs304_
                d_1859_v0_ = rhs305_
                d_1898_v32_: _dafny.Array
                nw328_ = _dafny.Array(int(0), 3)
                d_1898_v32_ = nw328_
                index299_ = default__.safeIndex(35, (d_1898_v32_).length(0))
                (d_1898_v32_)[index299_] = default__.safeModuloInt(p0, d_1859_v0_)
                index300_ = default__.safeIndex(35, (d_1898_v32_).length(0))
                (d_1898_v32_)[index300_] = d_1859_v0_
                d_1899_v33_: _dafny.Map
                d_1899_v33_ = _dafny.Map({d_1861_v2_: d_1898_v32_})
                d_1900_v34_: _dafny.Array
                nw329_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_1900_v34_ = nw329_
                d_1901_v35_: _dafny.Array
                def lambda114_(d_1902_v23_):
                    def lambda115_(d_1903_i5_):
                        return d_1902_v23_

                    return lambda115_

                init63_ = lambda114_(d_1888_v23_)
                nw330_ = _dafny.Array(None, 28)
                for i0_63_ in range(nw330_.length(0)):
                    nw330_[i0_63_] = init63_(i0_63_)
                d_1901_v35_ = nw330_
                index301_ = default__.safeIndex(107, (d_1900_v34_).length(0))
                (d_1900_v34_)[index301_] = d_1901_v35_
                index302_ = default__.safeIndex(107, (d_1900_v34_).length(0))
                rhs306_ = default__.safeModuloInt((d_1898_v32_)[default__.safeIndex(35, (d_1898_v32_).length(0))], (p0 if (d_1885_v21_).f12 else len(d_1876_v13_)))
                rhs307_ = ((d_1899_v33_) | (d_1899_v33_)) | (d_1899_v33_)
                rhs308_ = not(d_1885_v21_.f13)
                rhs309_ = default__.fm18(globalState)
                rhs310_ = d_1901_v35_
                lhs193_ = d_1885_v21_
                lhs194_ = d_1900_v34_
                lhs195_ = default__.safeIndex(107, (d_1900_v34_).length(0))
                d_1859_v0_ = rhs306_
                d_1899_v33_ = rhs307_
                lhs193_.f13 = rhs308_
                d_1861_v2_ = rhs309_
                lhs194_[lhs195_] = rhs310_
                d_1904_v36_: _dafny.Seq
                d_1904_v36_ = _dafny.SeqWithoutIsStrInference([not(not(self.f7)), self.f7])
                d_1905_v37_: _dafny.Seq
                d_1905_v37_ = _dafny.SeqWithoutIsStrInference([(d_1904_v36_)[default__.safeIndex(p0, len(d_1904_v36_))]])
                d_1906_v38_: _dafny.Map
                d_1906_v38_ = _dafny.Map({d_1885_v21_.f13: p0})
                d_1907_v39_: _dafny.Seq
                d_1907_v39_ = _dafny.SeqWithoutIsStrInference([d_1906_v38_, d_1906_v38_])
                d_1908_v40_: _dafny.Map
                d_1908_v40_ = _dafny.Map({len(d_1862_v3_): d_1861_v2_})
                d_1909_v41_: _dafny.Map
                d_1909_v41_ = _dafny.Map({d_1907_v39_: ((d_1905_v37_) + (_dafny.SeqWithoutIsStrInference([d_1885_v21_.f13, (d_1885_v21_).f12, (d_1885_v21_).f12, self.f7, (d_1885_v21_).f12]))).set(default__.safeIndex(len(d_1908_v40_), len((d_1905_v37_) + (_dafny.SeqWithoutIsStrInference([d_1885_v21_.f13, (d_1885_v21_).f12, (d_1885_v21_).f12, self.f7, (d_1885_v21_).f12])))), (d_1885_v21_).f12)})
                d_1905_v37_ = ((d_1909_v41_)[d_1907_v39_] if (d_1907_v39_) in (d_1909_v41_) else d_1905_v37_)
                index303_ = default__.safeIndex(35, (d_1898_v32_).length(0))
                rhs311_ = ((d_1898_v32_)[default__.safeIndex(35, (d_1898_v32_).length(0))]) - (d_1859_v0_)
                rhs312_ = d_1885_v21_.f13
                lhs196_ = d_1898_v32_
                lhs197_ = default__.safeIndex(35, (d_1898_v32_).length(0))
                lhs198_ = d_1885_v21_
                lhs196_[lhs197_] = rhs311_
                lhs198_.f13 = rhs312_
            elif True:
                d_1910_v42_: _dafny.Array
                nw331_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_1910_v42_ = nw331_
                d_1911_v43_: _dafny.Array
                nw332_ = _dafny.Array(int(0), 18)
                d_1911_v43_ = nw332_
                index304_ = default__.safeIndex(382, (d_1910_v42_).length(0))
                (d_1910_v42_)[index304_] = d_1911_v43_
                d_1912_v44_: _dafny.Map
                d_1912_v44_ = _dafny.Map({d_1859_v0_: d_1859_v0_})
                d_1913_v45_: _dafny.Map
                d_1913_v45_ = _dafny.Map({True: p0})
                d_1914_v46_: _dafny.Map
                d_1914_v46_ = _dafny.Map({D6_DC16(d_1885_v21_.f13): (d_1885_v21_).f12})
                index305_ = default__.safeIndex(382, (d_1910_v42_).length(0))
                nw333_ = _dafny.Array(None, 21)
                nw333_[int(0)] = p0
                nw333_[int(1)] = (d_1859_v0_ if d_1885_v21_.f13 else p0)
                nw333_[int(2)] = -965
                nw333_[int(3)] = d_1859_v0_
                nw333_[int(4)] = d_1859_v0_
                nw333_[int(5)] = d_1859_v0_
                nw333_[int(6)] = d_1859_v0_
                nw333_[int(7)] = default__.fm0(globalState)
                nw333_[int(8)] = p0
                nw333_[int(9)] = p0
                nw333_[int(10)] = p0
                nw333_[int(11)] = (0) - (default__.safeModuloInt(857, len(d_1912_v44_)))
                nw333_[int(12)] = p0
                nw333_[int(13)] = d_1859_v0_
                nw333_[int(14)] = (p0) - (len(d_1912_v44_))
                nw333_[int(15)] = (0) - ((p0) * (len((default__.fm53(len(d_1913_v45_), d_1885_v21_.f13, globalState)).cf15)))
                nw333_[int(16)] = p0
                nw333_[int(17)] = (0) - (default__.safeDivisionInt(d_1859_v0_, p0))
                nw333_[int(18)] = (d_1859_v0_) * (101)
                nw333_[int(19)] = d_1859_v0_
                nw333_[int(20)] = (507) * (len((d_1914_v46_).set(d_1887_v22_, (d_1885_v21_).f12)))
                (d_1910_v42_)[index305_] = nw333_
                d_1915_v47_: _dafny.Array
                nw334_ = _dafny.Array(None, 16)
                nw334_[int(0)] = not(d_1885_v21_.f13)
                nw334_[int(1)] = self.f7
                nw334_[int(2)] = (d_1885_v21_).f12
                nw334_[int(3)] = self.f7
                nw334_[int(4)] = d_1885_v21_.f13
                nw334_[int(5)] = d_1885_v21_.f13
                nw334_[int(6)] = (d_1885_v21_).f12
                nw334_[int(7)] = self.f7
                nw334_[int(8)] = (d_1885_v21_).f12
                nw334_[int(9)] = (d_1885_v21_).f12
                nw334_[int(10)] = self.f7
                nw334_[int(11)] = d_1885_v21_.f13
                nw334_[int(12)] = d_1885_v21_.f13
                nw334_[int(13)] = self.f7
                nw334_[int(14)] = (d_1885_v21_).f12
                nw334_[int(15)] = (d_1885_v21_).f12
                d_1915_v47_ = nw334_
                d_1916_v48_: _dafny.Set
                d_1916_v48_ = _dafny.Set({d_1915_v47_})
                d_1917_v49_: _dafny.Seq
                d_1917_v49_ = _dafny.SeqWithoutIsStrInference([d_1885_v21_.f13, (d_1885_v21_).f12])
                d_1918_v50_: D13
                d_1918_v50_ = D13_DC36(d_1861_v2_, d_1917_v49_, True)
                d_1919_v51_: C2
                nw335_ = C2()
                nw335_.ctor__(d_1916_v48_, (d_1918_v50_).cf58, self.f7)
                d_1919_v51_ = nw335_
                arr4_ = (d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]
                index306_ = default__.safeIndex(695, ((d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]).length(0))
                arr4_[index306_] = d_1859_v0_
                arr5_ = (d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]
                index307_ = default__.safeIndex(695, ((d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]).length(0))
                arr5_[index307_] = d_1859_v0_
                arr6_ = (d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]
                index308_ = default__.safeIndex(695, ((d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]).length(0))
                arr6_[index308_] = ((d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))])[default__.safeIndex(695, ((d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]).length(0))]
                d_1920_v52_: D9
                d_1920_v52_ = D9_DC24()
                arr7_ = (d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]
                index309_ = default__.safeIndex(695, ((d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]).length(0))
                arr7_[index309_] = len((D16_DC45(_dafny.Map({d_1920_v52_: (d_1910_v42_)[default__.safeIndex(382, (d_1910_v42_).length(0))]}))).cf72)
            d_1921_v53_: _dafny.Array
            def lambda116_(d_1922_p0_, d_1923_v0_):
                def lambda117_(d_1924_i6_):
                    return (D12_DC30((0) - (d_1922_p0_), False, _dafny.SeqWithoutIsStrInference([d_1923_v0_]), d_1922_p0_)).cf51

                return lambda117_

            init64_ = lambda116_(p0, d_1859_v0_)
            nw336_ = _dafny.Array(None, 5)
            for i0_64_ in range(nw336_.length(0)):
                nw336_[i0_64_] = init64_(i0_64_)
            d_1921_v53_ = nw336_
            d_1925_v54_: _dafny.Seq
            d_1925_v54_ = _dafny.SeqWithoutIsStrInference([d_1859_v0_])
            index310_ = default__.safeIndex(17, (d_1921_v53_).length(0))
            (d_1921_v53_)[index310_] = (_dafny.SeqWithoutIsStrInference([p0])) + (d_1925_v54_)
            d_1926_v55_: _dafny.MultiSet
            d_1926_v55_ = _dafny.MultiSet([self.f7])
            index311_ = default__.safeIndex(17, (d_1921_v53_).length(0))
            (d_1921_v53_)[index311_] = ((d_1925_v54_).set(default__.safeIndex(len(d_1862_v3_), len(d_1925_v54_)), (d_1926_v55_).cardinality)) + (_dafny.SeqWithoutIsStrInference([(0) - (p0) for d_1927_i7_ in range(default__.abs(983))]))
            d_1926_v55_ = d_1926_v55_
            d_1928_v56_: _dafny.Array
            nw337_ = _dafny.Array(False, 7)
            d_1928_v56_ = nw337_
            d_1929_v57_: _dafny.Map
            d_1929_v57_ = _dafny.Map({-271: d_1928_v56_})
            d_1930_v58_: _dafny.Array
            nw338_ = _dafny.Array(None, 3)
            nw338_[int(0)] = d_1885_v21_.f13
            nw338_[int(1)] = (len(d_1929_v57_)) >= (d_1859_v0_)
            nw338_[int(2)] = (p0) < (17)
            d_1930_v58_ = nw338_
            index312_ = default__.safeIndex(928, (d_1928_v56_).length(0))
            (d_1928_v56_)[index312_] = d_1885_v21_.f13
            d_1931_v59_: _dafny.Set
            d_1931_v59_ = _dafny.Set({(d_1885_v21_).f12, False})
            d_1932_v60_: D1
            d_1932_v60_ = D1_DC3(d_1931_v59_)
            d_1933_v61_: D0
            d_1933_v61_ = D0_DC0(len((d_1862_v3_).set(default__.safeIndex(p0, len(d_1862_v3_)), d_1861_v2_)))
            d_1934_v62_: D8
            d_1934_v62_ = D8_DC22(446, d_1932_v60_, d_1933_v61_)
            d_1935_v63_: _dafny.Seq
            d_1935_v63_ = _dafny.SeqWithoutIsStrInference([True, default__.fm1(_dafny.SeqWithoutIsStrInference([d_1933_v61_ for d_1936_i8_ in range(default__.abs(578))]), d_1859_v0_, p0, globalState), d_1885_v21_.f13])
            index313_ = default__.safeIndex(928, (d_1928_v56_).length(0))
            rhs313_ = (d_1885_v21_).f12
            rhs314_ = (d_1935_v63_)[default__.safeIndex(len(d_1862_v3_), len(d_1935_v63_))]
            rhs315_ = ((d_1862_v3_) + (d_1862_v3_)) <= (d_1862_v3_)
            rhs316_ = d_1861_v2_
            rhs317_ = d_1934_v62_
            lhs199_ = d_1928_v56_
            lhs200_ = default__.safeIndex(928, (d_1928_v56_).length(0))
            lhs201_ = d_1885_v21_
            lhs202_ = d_1885_v21_
            lhs199_[lhs200_] = rhs313_
            lhs201_.f13 = rhs314_
            lhs202_.f13 = rhs315_
            d_1861_v2_ = rhs316_
            d_1934_v62_ = rhs317_
        elif True:
            if (True if self.f7 else self.f7):
                d_1937_v64_: _dafny.Array
                def lambda118_(d_1938_i9_):
                    return _dafny.MultiSet([self.f7, not(self.f7)])

                init65_ = lambda118_
                nw339_ = _dafny.Array(None, 15)
                for i0_65_ in range(nw339_.length(0)):
                    nw339_[i0_65_] = init65_(i0_65_)
                d_1937_v64_ = nw339_
                d_1937_v64_ = d_1937_v64_
                d_1939_v65_: D0
                d_1939_v65_ = D0_DC0(p0)
                d_1940_v66_: _dafny.Seq
                d_1940_v66_ = _dafny.SeqWithoutIsStrInference([D0_DC0(p0), d_1939_v65_, d_1939_v65_])
                d_1941_v67_: _dafny.Seq
                d_1941_v67_ = _dafny.SeqWithoutIsStrInference([self.f7])
                d_1942_v68_: _dafny.Seq
                d_1942_v68_ = _dafny.SeqWithoutIsStrInference([len(d_1941_v67_)])
                d_1943_v69_: C1
                nw340_ = C1()
                nw340_.ctor__(default__.fm1(d_1940_v66_, len(d_1862_v3_), (0) - ((d_1942_v68_)[default__.safeIndex(p0, len(d_1942_v68_))]), globalState))
                d_1943_v69_ = nw340_
                d_1944_v70_: _dafny.Array
                nw341_ = _dafny.Array(False, 23)
                d_1944_v70_ = nw341_
                index314_ = default__.safeIndex(542, (d_1944_v70_).length(0))
                (d_1944_v70_)[index314_] = self.f7
                index315_ = default__.safeIndex(542, (d_1944_v70_).length(0))
                (d_1944_v70_)[index315_] = self.f7
                d_1945_v71_: _dafny.Map
                d_1945_v71_ = _dafny.Map({(d_1944_v70_)[default__.safeIndex(542, (d_1944_v70_).length(0))]: (d_1859_v0_) - (d_1859_v0_)})
                d_1945_v71_ = default__.fm43(d_1861_v2_, self.f7, self.f7, globalState)
                d_1946_v72_: _dafny.MultiSet
                d_1946_v72_ = _dafny.MultiSet([d_1859_v0_])
                d_1947_v73_: C3
                nw342_ = C3()
                nw342_.ctor__((d_1859_v0_) - (151), (d_1946_v72_).issubset(d_1946_v72_))
                d_1947_v73_ = nw342_
            elif True:
                d_1948_v74_: D6
                d_1948_v74_ = D6_DC17(self.f7, d_1859_v0_, d_1861_v2_, d_1861_v2_)
                d_1949_v75_: _dafny.Map
                d_1949_v75_ = _dafny.Map({default__.fm0(globalState): (0) - ((0) - (default__.fm0(globalState)))})
                d_1950_v77_: _dafny.Set
                d_1950_v77_ = _dafny.Set({len(d_1949_v75_)})
                d_1951_v78_: _dafny.Array
                nw343_ = _dafny.Array(None, 16)
                nw343_[int(0)] = p0
                nw343_[int(1)] = -252
                nw343_[int(2)] = d_1859_v0_
                nw343_[int(3)] = (250) + ((d_1948_v74_).cf32)
                nw343_[int(4)] = d_1859_v0_
                nw343_[int(5)] = (d_1859_v0_) + (d_1859_v0_)
                nw343_[int(6)] = p0
                nw343_[int(7)] = d_1859_v0_
                nw343_[int(8)] = len(d_1860_v1_)
                nw343_[int(9)] = d_1859_v0_
                def iife154_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in (d_1950_v77_).Elements:
                        d_1952_v76_: int = compr_40_
                        if (d_1952_v76_) in (d_1950_v77_):
                            coll40_[default__.safeModuloInt(d_1952_v76_, len(d_1862_v3_))] = self.f7
                    return _dafny.Map(coll40_)
                nw343_[int(10)] = ((d_1949_v75_)[p0] if (p0) in (d_1949_v75_) else len(iife154_()
                ))
                nw343_[int(11)] = p0
                nw343_[int(12)] = d_1859_v0_
                nw343_[int(13)] = (0) - ((default__.fm0(globalState)) + (len(d_1862_v3_)))
                nw343_[int(14)] = (d_1859_v0_) + (len(_dafny.SeqWithoutIsStrInference([p0 for d_1953_i10_ in range(default__.abs(81))])))
                nw343_[int(15)] = p0
                d_1951_v78_ = nw343_
                index316_ = default__.safeIndex(846, (d_1951_v78_).length(0))
                (d_1951_v78_)[index316_] = p0
                index317_ = default__.safeIndex(846, (d_1951_v78_).length(0))
                (d_1951_v78_)[index317_] = (276) + (p0)
                d_1954_v79_: _dafny.Seq
                d_1954_v79_ = _dafny.SeqWithoutIsStrInference([self.f7, False])
                d_1955_v80_: _dafny.Seq
                d_1955_v80_ = _dafny.SeqWithoutIsStrInference([self.f7, self.f7, (d_1954_v79_)[default__.safeIndex(d_1859_v0_, len(d_1954_v79_))], self.f7, self.f7])
                d_1956_v81_: _dafny.Seq
                d_1956_v81_ = _dafny.SeqWithoutIsStrInference([len(d_1955_v80_)])
                d_1957_v82_: _dafny.Map
                d_1957_v82_ = _dafny.Map({d_1951_v78_: d_1859_v0_})
                index318_ = default__.safeIndex(846, (d_1951_v78_).length(0))
                (d_1951_v78_)[index318_] = (d_1956_v81_)[default__.safeIndex(default__.safeModuloInt(((d_1957_v82_)[d_1951_v78_] if (d_1951_v78_) in (d_1957_v82_) else (d_1951_v78_)[default__.safeIndex(846, (d_1951_v78_).length(0))]), (d_1956_v81_)[default__.safeIndex(d_1859_v0_, len(d_1956_v81_))]), len(d_1956_v81_))]
                d_1958_v83_: _dafny.Seq
                d_1958_v83_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")), ((((d_1875_v12_)[((d_1876_v13_)[self.f7] if (self.f7) in (d_1876_v13_) else self.f7)] if (((d_1876_v13_)[self.f7] if (self.f7) in (d_1876_v13_) else self.f7)) in (d_1875_v12_) else d_1862_v3_)) + (d_1862_v3_)).set(default__.safeIndex((d_1951_v78_)[default__.safeIndex(846, (d_1951_v78_).length(0))], len((((d_1875_v12_)[((d_1876_v13_)[self.f7] if (self.f7) in (d_1876_v13_) else self.f7)] if (((d_1876_v13_)[self.f7] if (self.f7) in (d_1876_v13_) else self.f7)) in (d_1875_v12_) else d_1862_v3_)) + (d_1862_v3_))), d_1861_v2_), d_1862_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "volltvv")), d_1862_v3_])
                d_1959_v84_: D9
                d_1959_v84_ = D9_DC26(d_1956_v81_)
                d_1960_v85_: D12
                d_1960_v85_ = D12_DC31(d_1959_v84_, self.f7)
                d_1961_v86_: _dafny.Map
                d_1961_v86_ = _dafny.Map({True: d_1954_v79_})
                d_1958_v83_ = default__.fm3(self.f7, d_1859_v0_, (default__.fm49((d_1960_v85_).cf54, d_1961_v86_, self.f7, globalState)).cardinality, ((d_1951_v78_)[default__.safeIndex(846, (d_1951_v78_).length(0))] if self.f7 else d_1859_v0_), globalState)
                d_1962_v87_: _dafny.MultiSet
                d_1962_v87_ = _dafny.MultiSet([(p0) * (p0)])
                d_1962_v87_ = d_1962_v87_
                (self).f7 = self.f7
            (self).f7 = True
            d_1963_v88_: _dafny.Seq
            d_1963_v88_ = _dafny.SeqWithoutIsStrInference([382, (d_1859_v0_) + (d_1859_v0_), p0])
            d_1859_v0_ = (d_1963_v88_)[default__.safeIndex(p0, len(d_1963_v88_))]
            d_1862_v3_ = _dafny.SeqWithoutIsStrInference([d_1861_v2_ for d_1964_i11_ in range(default__.abs(380))])
            d_1965_v89_: _dafny.Set
            d_1965_v89_ = _dafny.Set({self.f7, self.f7, self.f7, self.f7, self.f7})
            d_1966_v90_: _dafny.Map
            d_1966_v90_ = _dafny.Map({len((d_1965_v89_) | (d_1965_v89_)): d_1862_v3_})
            d_1966_v90_ = (d_1966_v90_).set(default__.safeDivisionInt(p0, p0), d_1862_v3_)
        d_1967_v91_: C3
        nw344_ = C3()
        nw344_.ctor__(p0, self.f7)
        d_1967_v91_ = nw344_
        d_1968_v92_: _dafny.Map
        d_1968_v92_ = _dafny.Map({self.f7: d_1967_v91_})
        hi8_ = (-406 if self.f7 else len(d_1968_v92_))
        for d_1969_i12_ in range((self).fm15(globalState), hi8_):
            d_1970_v93_: _dafny.Array
            def lambda119_(d_1971_i13_):
                return self.f7

            init66_ = lambda119_
            nw345_ = _dafny.Array(None, 13)
            for i0_66_ in range(nw345_.length(0)):
                nw345_[i0_66_] = init66_(i0_66_)
            d_1970_v93_ = nw345_
            index319_ = default__.safeIndex(225, (d_1970_v93_).length(0))
            (d_1970_v93_)[index319_] = self.f7
            d_1972_v94_: _dafny.Array
            nw346_ = _dafny.Array(None, 1)
            nw346_[int(0)] = d_1859_v0_
            d_1972_v94_ = nw346_
            d_1973_v95_: _dafny.Set
            d_1973_v95_ = _dafny.Set({d_1972_v94_})
            d_1974_v96_: D0
            d_1974_v96_ = D0_DC2(self.f7, p0, d_1973_v95_, p0)
            d_1975_v97_: _dafny.MultiSet
            d_1975_v97_ = _dafny.MultiSet([self.f7])
            index320_ = default__.safeIndex(225, (d_1970_v93_).length(0))
            (d_1970_v93_)[index320_] = ((d_1974_v96_).cf3) == (((d_1975_v97_) | (d_1975_v97_)).cardinality)
            index321_ = default__.safeIndex(225, (d_1970_v93_).length(0))
            rhs318_ = False
            rhs319_ = (d_1970_v93_)[default__.safeIndex(225, (d_1970_v93_).length(0))]
            lhs203_ = d_1970_v93_
            lhs204_ = default__.safeIndex(225, (d_1970_v93_).length(0))
            lhs205_ = self
            lhs203_[lhs204_] = rhs318_
            lhs205_.f7 = rhs319_
            d_1859_v0_ = d_1859_v0_
            d_1976_v98_: _dafny.Seq
            d_1976_v98_ = _dafny.SeqWithoutIsStrInference([d_1860_v1_])
            (self).f7 = (len((d_1976_v98_)[default__.safeIndex((d_1967_v91_).f14, len(d_1976_v98_))])) >= (777)

    @property
    def f11(self):
        return self._f11

class C5:
    def  __init__(self):
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f10):
        (self)._f10 = f10

    def m7(self, globalState):
        d_1977_v0_: str
        d_1977_v0_ = _dafny.CodePoint('n')
        d_1978_v1_: _dafny.Seq
        d_1978_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xb"))
        d_1979_v2_: _dafny.Map
        d_1979_v2_ = _dafny.Map({d_1977_v0_: len(d_1978_v1_)})
        d_1980_v3_: int
        d_1980_v3_ = 68
        hi9_ = d_1980_v3_
        for d_1981_i0_ in range(((d_1979_v2_)[d_1977_v0_] if (d_1977_v0_) in (d_1979_v2_) else d_1980_v3_), hi9_):
            if (self).f10:
                d_1982_v4_: bool
                d_1982_v4_ = True
                d_1982_v4_ = (self).f10
                d_1983_v5_: _dafny.MultiSet
                d_1983_v5_ = _dafny.MultiSet([d_1981_i0_])
                rhs320_ = (self).f10
                rhs321_ = _dafny.SeqWithoutIsStrInference([d_1977_v0_ for d_1984_i1_ in range(default__.abs(204))])
                rhs322_ = d_1983_v5_
                d_1982_v4_ = rhs320_
                d_1978_v1_ = rhs321_
                d_1983_v5_ = rhs322_
                d_1985_v6_: _dafny.MultiSet
                d_1985_v6_ = _dafny.MultiSet([(d_1981_i0_) <= (default__.fm0(globalState)), not(((self).f10) and ((self).f10)), not(d_1982_v4_), (self).f10])
                d_1986_v8_: _dafny.Set
                d_1986_v8_ = _dafny.Set({d_1980_v3_})
                d_1987_v9_: _dafny.Seq
                d_1987_v9_ = _dafny.SeqWithoutIsStrInference([d_1982_v4_, d_1982_v4_, d_1982_v4_])
                d_1988_v10_: D0
                d_1988_v10_ = D0_DC1(len(d_1978_v1_))
                d_1989_v11_: _dafny.Map
                def iife155_(_pat_let57_0):
                    def iife156_(d_1990_dt__update__tmp_h0_):
                        def iife157_(_pat_let58_0):
                            def iife158_(d_1991_dt__update_hcf1_h0_):
                                return D0_DC1(d_1991_dt__update_hcf1_h0_)
                            return iife158_(_pat_let58_0)
                        return iife157_(-760)
                    return iife156_(_pat_let57_0)
                d_1989_v11_ = _dafny.Map({(iife155_(d_1988_v10_)).cf1: d_1980_v3_})
                def iife159_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in (d_1986_v8_).Elements:
                        d_1992_v7_: int = compr_41_
                        if (d_1992_v7_) in (d_1986_v8_):
                            coll41_[default__.safeModuloInt(d_1992_v7_, d_1981_i0_)] = _dafny.Map({len(d_1978_v1_): d_1982_v4_})
                    return _dafny.Map(coll41_)
                rhs323_ = (-407) in (iife159_()
                )
                rhs324_ = d_1985_v6_
                rhs325_ = (0) - (len(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_1982_v4_, (self).f10, (d_1987_v9_)[default__.safeIndex(len(d_1989_v11_), len(d_1987_v9_))]]), d_1987_v9_, d_1987_v9_})))
                d_1982_v4_ = rhs323_
                d_1985_v6_ = rhs324_
                d_1980_v3_ = rhs325_
                d_1982_v4_ = d_1982_v4_
                d_1993_v12_: D1
                d_1993_v12_ = D1_DC4(d_1982_v4_, not(d_1982_v4_))
                d_1994_v13_: _dafny.Seq
                def iife160_(_pat_let59_0):
                    def iife161_(d_1995_dt__update__tmp_h1_):
                        def iife162_(_pat_let60_0):
                            def iife163_(d_1996_dt__update_hcf8_h0_):
                                return D1_DC4((d_1995_dt__update__tmp_h1_).cf7, d_1996_dt__update_hcf8_h0_)
                            return iife163_(_pat_let60_0)
                        return iife162_(False)
                    return iife161_(_pat_let59_0)
                d_1994_v13_ = _dafny.SeqWithoutIsStrInference([iife160_(d_1993_v12_)])
                d_1994_v13_ = (d_1994_v13_) + ((d_1994_v13_) + (d_1994_v13_))
            elif True:
                d_1997_v14_: _dafny.Array
                nw347_ = _dafny.Array(False, 2)
                d_1997_v14_ = nw347_
                d_1998_v15_: _dafny.Map
                d_1998_v15_ = _dafny.Map({d_1997_v14_: d_1981_i0_})
                d_1998_v15_ = d_1998_v15_
                d_1999_v16_: bool
                d_1999_v16_ = True
                d_1999_v16_ = (self).f10
                d_2000_v17_: D0
                d_2000_v17_ = D0_DC0(d_1980_v3_)
                d_2001_v18_: _dafny.Seq
                d_2001_v18_ = _dafny.SeqWithoutIsStrInference([d_2000_v17_])
                d_1999_v16_ = default__.fm1(d_2001_v18_, d_1981_i0_, ((0) - (d_1980_v3_)) + (d_1980_v3_), globalState)
                index322_ = default__.safeIndex(400, (d_1997_v14_).length(0))
                (d_1997_v14_)[index322_] = not (d_1999_v16_) or (not(d_1999_v16_))
                index323_ = default__.safeIndex(400, (d_1997_v14_).length(0))
                rhs326_ = d_1999_v16_
                rhs327_ = True
                lhs206_ = d_1997_v14_
                lhs207_ = default__.safeIndex(400, (d_1997_v14_).length(0))
                lhs206_[lhs207_] = rhs326_
                d_1999_v16_ = rhs327_
                d_2002_v19_: _dafny.Set
                d_2002_v19_ = _dafny.Set({d_1981_i0_, default__.fm0(globalState), d_1980_v3_})
                d_2003_v20_: D2
                d_2003_v20_ = D2_DC7(d_2002_v19_)
                def iife164_(_pat_let61_0):
                    def iife165_(d_2004_dt__update__tmp_h2_):
                        def iife167_():
                            coll42_ = _dafny.Set()
                            compr_42_: int
                            for compr_42_ in _dafny.IntegerRange(133, -938):
                                d_2005_v21_: int = compr_42_
                                if ((133) <= (d_2005_v21_)) and ((d_2005_v21_) < (-938)):
                                    coll42_ = coll42_.union(_dafny.Set([default__.safeDivisionInt(d_2005_v21_, 462)]))
                            return _dafny.Set(coll42_)
                        def iife166_(_pat_let62_0):
                            def iife168_(d_2006_dt__update_hcf13_h0_):
                                return D2_DC7(d_2006_dt__update_hcf13_h0_)
                            return iife168_(_pat_let62_0)
                        return iife166_(iife167_()
                        )
                    return iife165_(_pat_let61_0)
                d_2002_v19_ = (iife164_(d_2003_v20_)).cf13
            d_2007_v22_: D1
            d_2007_v22_ = D1_DC4((self).f10, (self).f10)
            d_2008_v23_: _dafny.Array
            nw348_ = _dafny.Array(None, 7)
            nw348_[int(0)] = (self).f10
            nw348_[int(1)] = False
            nw348_[int(2)] = (self).f10
            nw348_[int(3)] = (self).f10
            nw348_[int(4)] = (self).f10
            nw348_[int(5)] = (self).f10
            nw348_[int(6)] = ((d_2007_v22_).cf8) and ((self).f10)
            d_2008_v23_ = nw348_
            index324_ = default__.safeIndex(584, (d_2008_v23_).length(0))
            (d_2008_v23_)[index324_] = False
            d_2009_v24_: bool
            d_2009_v24_ = False
            d_2010_v25_: D0
            d_2010_v25_ = D0_DC0(d_1980_v3_)
            d_2011_v26_: _dafny.Seq
            d_2011_v26_ = _dafny.SeqWithoutIsStrInference([d_2010_v25_, d_2010_v25_])
            d_2012_v27_: _dafny.Map
            d_2012_v27_ = _dafny.Map({d_1978_v1_: default__.fm1(d_2011_v26_, d_1981_i0_, d_1980_v3_, globalState)})
            d_2013_v28_: _dafny.Seq
            d_2013_v28_ = _dafny.SeqWithoutIsStrInference([d_1978_v1_, d_1978_v1_])
            d_2014_v29_: _dafny.Map
            d_2014_v29_ = _dafny.Map({d_1981_i0_: d_2009_v24_})
            d_2015_v30_: _dafny.Seq
            d_2015_v30_ = _dafny.SeqWithoutIsStrInference([d_1981_i0_])
            index325_ = default__.safeIndex(584, (d_2008_v23_).length(0))
            rhs328_ = (D2_DC8((self).f10, d_1978_v1_)).cf14
            rhs329_ = ((d_2012_v27_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akjinxm"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akjinxm"))) in (d_2012_v27_) else (not(d_2009_v24_)) or ((self).f10))
            rhs330_ = ((d_1978_v1_) + (d_1978_v1_)) + ((d_1978_v1_ if d_2009_v24_ else (d_2013_v28_)[default__.safeIndex(len(d_2014_v29_), len(d_2013_v28_))]))
            rhs331_ = (0) - (default__.fm0(globalState))
            rhs332_ = not ((self).f10) or (default__.fm1(d_2011_v26_, len(d_2015_v30_), d_1980_v3_, globalState))
            lhs208_ = d_2008_v23_
            lhs209_ = default__.safeIndex(584, (d_2008_v23_).length(0))
            lhs208_[lhs209_] = rhs328_
            d_2009_v24_ = rhs329_
            d_1978_v1_ = rhs330_
            d_1980_v3_ = rhs331_
            d_2009_v24_ = rhs332_
            d_2016_v31_: D1
            d_2016_v31_ = D1_DC4((self).f10, d_2009_v24_)
            d_2017_v32_: D1
            d_2017_v32_ = D1_DC6(d_2016_v31_)
            d_2018_v33_: _dafny.MultiSet
            d_2018_v33_ = _dafny.MultiSet([d_1980_v3_])
            d_2019_v34_: _dafny.Set
            d_2019_v34_ = _dafny.Set({(d_2008_v23_)[default__.safeIndex(584, (d_2008_v23_).length(0))], (self).f10})
            d_2020_v35_: _dafny.Seq
            d_2020_v35_ = _dafny.SeqWithoutIsStrInference([D1_DC6(d_2016_v31_), default__.fm14(d_1981_i0_, (d_2018_v33_).cardinality, len(d_2019_v34_), default__.fm1(_dafny.SeqWithoutIsStrInference([d_2010_v25_, d_2010_v25_, D0_DC0(d_1981_i0_), d_2010_v25_, d_2010_v25_]), len(d_1978_v1_), d_1980_v3_, globalState), globalState), d_2016_v31_, d_2016_v31_, d_2016_v31_])
            pat_let_tv47_ = d_2016_v31_
            d_2021_v36_: _dafny.Array
            nw349_ = _dafny.Array(None, 19)
            nw349_[int(0)] = d_2017_v32_
            nw349_[int(1)] = d_2017_v32_
            nw349_[int(2)] = d_2017_v32_
            nw349_[int(3)] = d_2017_v32_
            nw349_[int(4)] = d_2017_v32_
            nw349_[int(5)] = d_2017_v32_
            nw349_[int(6)] = d_2017_v32_
            nw349_[int(7)] = D1_DC6(d_2016_v31_)
            nw349_[int(8)] = d_2017_v32_
            nw349_[int(9)] = d_2017_v32_
            nw349_[int(10)] = d_2017_v32_
            nw349_[int(11)] = d_2017_v32_
            nw349_[int(12)] = d_2017_v32_
            nw349_[int(13)] = D1_DC6((d_2020_v35_)[default__.safeIndex(d_1981_i0_, len(d_2020_v35_))])
            nw349_[int(14)] = d_2017_v32_
            nw349_[int(15)] = D1_DC6(d_2016_v31_)
            nw349_[int(16)] = d_2017_v32_
            nw349_[int(17)] = d_2017_v32_
            def iife169_(_pat_let63_0):
                def iife170_(d_2022_dt__update__tmp_h3_):
                    def iife171_(_pat_let64_0):
                        def iife172_(d_2023_dt__update_hcf12_h0_):
                            return D1_DC6(d_2023_dt__update_hcf12_h0_)
                        return iife172_(_pat_let64_0)
                    return iife171_(pat_let_tv47_)
                return iife170_(_pat_let63_0)
            nw349_[int(18)] = iife169_(D1_DC6(d_2016_v31_))
            d_2021_v36_ = nw349_
            d_2024_v37_: _dafny.Set
            d_2024_v37_ = _dafny.Set({d_2021_v36_})
            d_2009_v24_ = (d_2024_v37_).ispropersubset(d_2024_v37_)
            d_2025_v38_: _dafny.Seq
            d_2025_v38_ = _dafny.SeqWithoutIsStrInference([(d_2015_v30_ if (d_2008_v23_)[default__.safeIndex(584, (d_2008_v23_).length(0))] else _dafny.SeqWithoutIsStrInference([d_1981_i0_ for d_2026_i2_ in range(default__.abs(938))])), (d_2015_v30_) + (d_2015_v30_), _dafny.SeqWithoutIsStrInference([d_1980_v3_])])
            d_2015_v30_ = (d_2025_v38_)[default__.safeIndex((d_1980_v3_) * (d_1980_v3_), len(d_2025_v38_))]
        d_2027_v39_: bool
        d_2027_v39_ = False
        d_2028_v41_: _dafny.Seq
        def iife173_():
            coll43_ = _dafny.Map()
            compr_43_: str
            for compr_43_ in (d_1978_v1_).Elements:
                d_2029_v40_: str = compr_43_
                if (d_2029_v40_) in (d_1978_v1_):
                    coll43_[d_2029_v40_] = d_1980_v3_
            return _dafny.Map(coll43_)
        d_2028_v41_ = _dafny.SeqWithoutIsStrInference([len(iife173_()
        )])
        d_2030_v42_: _dafny.MultiSet
        d_2030_v42_ = _dafny.MultiSet([d_1978_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fogjs")), d_1978_v1_])
        d_2027_v39_ = (len(d_2028_v41_)) > (((d_2030_v42_).cardinality) + (d_1980_v3_))
        d_2031_i3_: int
        d_2031_i3_ = 0
        with _dafny.label("10"):
            while d_2027_v39_:
                with _dafny.c_label("10"):
                    if (d_2031_i3_) >= (100):
                        raise _dafny.Break("10")
                    d_2031_i3_ = (d_2031_i3_) + (1)
                    d_2032_v43_: D2
                    d_2032_v43_ = D2_DC8(d_2027_v39_, d_1978_v1_)
                    d_1980_v3_ = len((d_2032_v43_).cf15)
                    d_2033_v44_: _dafny.Array
                    def lambda120_(d_2034_v3_):
                        def lambda121_(d_2035_i4_):
                            return _dafny.Map({(self).f10: d_2034_v3_})

                        return lambda121_

                    init67_ = lambda120_(d_1980_v3_)
                    nw350_ = _dafny.Array(None, 15)
                    for i0_67_ in range(nw350_.length(0)):
                        nw350_[i0_67_] = init67_(i0_67_)
                    d_2033_v44_ = nw350_
                    d_2036_v45_: _dafny.MultiSet
                    d_2036_v45_ = _dafny.MultiSet([d_1980_v3_])
                    d_2037_v46_: _dafny.Map
                    d_2037_v46_ = _dafny.Map({d_1980_v3_: d_1980_v3_})
                    d_2038_v47_: _dafny.Map
                    d_2038_v47_ = _dafny.Map({(self).f10: ((d_2036_v45_)[d_1980_v3_] if (d_1980_v3_) in (d_2036_v45_) else ((d_2037_v46_)[len(_dafny.Map({d_2027_v39_: d_1980_v3_}))] if (len(_dafny.Map({d_2027_v39_: d_1980_v3_}))) in (d_2037_v46_) else len(d_1978_v1_)))})
                    index326_ = default__.safeIndex(875, (d_2033_v44_).length(0))
                    (d_2033_v44_)[index326_] = d_2038_v47_
                    d_2039_v48_: _dafny.Map
                    d_2039_v48_ = _dafny.Map({d_1980_v3_: d_2027_v39_})
                    index327_ = default__.safeIndex(875, (d_2033_v44_).length(0))
                    (d_2033_v44_)[index327_] = (_dafny.Map({(self).f10: d_1980_v3_}) if ((d_2039_v48_)[543] if (543) in (d_2039_v48_) else d_2027_v39_) else _dafny.Map({not(d_2027_v39_): d_1980_v3_}))
                    d_2040_v49_: _dafny.Array
                    nw351_ = _dafny.Array(int(0), 24)
                    d_2040_v49_ = nw351_
                    d_2041_v50_: _dafny.Set
                    d_2041_v50_ = _dafny.Set({d_2040_v49_})
                    d_2042_v51_: D0
                    d_2042_v51_ = D0_DC2(d_2027_v39_, 693, d_2041_v50_, d_1980_v3_)
                    d_2043_v52_: _dafny.MultiSet
                    d_2043_v52_ = _dafny.MultiSet([d_2042_v51_])
                    d_2044_v53_: _dafny.Set
                    d_2044_v53_ = _dafny.Set({d_2043_v52_})
                    d_2044_v53_ = _dafny.Set({(d_2043_v52_).set(d_2042_v51_, default__.abs(285))})
                    d_2045_v54_: _dafny.Set
                    d_2045_v54_ = _dafny.Set({-256})
                    def iife174_():
                        coll44_ = _dafny.Set()
                        compr_44_: int
                        for compr_44_ in _dafny.IntegerRange(467, 541):
                            d_2046_v55_: int = compr_44_
                            if ((467) <= (d_2046_v55_)) and ((d_2046_v55_) < (541)):
                                coll44_ = coll44_.union(_dafny.Set([(d_2046_v55_) - (d_1980_v3_)]))
                        return _dafny.Set(coll44_)
                    d_2045_v54_ = (iife174_()
                    ) - (d_2045_v54_)
                    pass
            pass
        d_2047_v56_: _dafny.Array
        def lambda122_(d_2048_v1_):
            def lambda123_(d_2049_i5_):
                return (d_2049_i5_) - (len(d_2048_v1_))

            return lambda123_

        init68_ = lambda122_(d_1978_v1_)
        nw352_ = _dafny.Array(None, 22)
        for i0_68_ in range(nw352_.length(0)):
            nw352_[i0_68_] = init68_(i0_68_)
        d_2047_v56_ = nw352_
        d_2050_v57_: _dafny.Seq
        d_2050_v57_ = _dafny.SeqWithoutIsStrInference([(self).f10, d_2027_v39_])
        index328_ = default__.safeIndex(864, (d_2047_v56_).length(0))
        (d_2047_v56_)[index328_] = len(_dafny.Map({len(d_2050_v57_): default__.fm0(globalState)}))
        index329_ = default__.safeIndex(864, (d_2047_v56_).length(0))
        (d_2047_v56_)[index329_] = d_1980_v3_
        d_2051_i6_: int
        d_2051_i6_ = 0
        with _dafny.label("11"):
            while (self).f10:
                with _dafny.c_label("11"):
                    if (d_2051_i6_) >= (100):
                        raise _dafny.Break("11")
                    d_2051_i6_ = (d_2051_i6_) + (1)
                    d_1980_v3_ = d_1980_v3_
                    d_2027_v39_ = d_2027_v39_
                    d_2052_v58_: _dafny.Array
                    nw353_ = _dafny.Array(_dafny.Array(None, 0), 18)
                    d_2052_v58_ = nw353_
                    d_2053_v59_: _dafny.Map
                    d_2053_v59_ = _dafny.Map({(self).f10: d_2047_v56_})
                    index330_ = default__.safeIndex(486, (d_2052_v58_).length(0))
                    (d_2052_v58_)[index330_] = ((d_2053_v59_)[(self).f10] if ((self).f10) in (d_2053_v59_) else d_2047_v56_)
                    index331_ = default__.safeIndex(486, (d_2052_v58_).length(0))
                    (d_2052_v58_)[index331_] = d_2047_v56_
                    d_2054_v60_: _dafny.MultiSet
                    d_2054_v60_ = _dafny.MultiSet([(self).f10, d_2027_v39_, d_2027_v39_, not(d_2027_v39_), (self).f10])
                    d_2055_v61_: _dafny.Map
                    d_2055_v61_ = _dafny.Map({(self).f10: _dafny.MultiSet([d_2027_v39_])})
                    index332_ = default__.safeIndex(864, (d_2047_v56_).length(0))
                    index333_ = default__.safeIndex(864, (d_2047_v56_).length(0))
                    rhs333_ = (d_2047_v56_)[default__.safeIndex(864, (d_2047_v56_).length(0))]
                    rhs334_ = default__.fm0(globalState)
                    rhs335_ = ((d_2055_v61_)[d_2027_v39_] if (d_2027_v39_) in (d_2055_v61_) else d_2054_v60_)
                    rhs336_ = (len((d_2050_v57_) + (d_2050_v57_))) * (d_1980_v3_)
                    lhs210_ = d_2047_v56_
                    lhs211_ = default__.safeIndex(864, (d_2047_v56_).length(0))
                    lhs212_ = d_2047_v56_
                    lhs213_ = default__.safeIndex(864, (d_2047_v56_).length(0))
                    lhs210_[lhs211_] = rhs333_
                    d_1980_v3_ = rhs334_
                    d_2054_v60_ = rhs335_
                    lhs212_[lhs213_] = rhs336_
                    pass
            pass
        d_2056_v62_: _dafny.Set
        d_2056_v62_ = _dafny.Set({(d_2047_v56_)[default__.safeIndex(864, (d_2047_v56_).length(0))], d_1980_v3_, (d_2047_v56_)[default__.safeIndex(864, (d_2047_v56_).length(0))], d_1980_v3_})
        d_1980_v3_ = (len((d_2056_v62_) - (d_2056_v62_))) + (default__.safeDivisionInt((d_2047_v56_)[default__.safeIndex(864, (d_2047_v56_).length(0))], 148))

    def m8(self, p0, p1, globalState):
        r0: D0 = D0.default()()
        d_2057_v0_: int
        d_2057_v0_ = 984
        def iife175_():
            coll45_ = _dafny.Map()
            compr_45_: int
            for compr_45_ in _dafny.IntegerRange(219, 267):
                d_2058_v1_: int = compr_45_
                if ((219) <= (d_2058_v1_)) and ((d_2058_v1_) < (267)):
                    coll45_[default__.safeDivisionInt(d_2058_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kja"))))] = p1
            return _dafny.Map(coll45_)
        hi10_ = len(iife175_()
        )
        for d_2059_i0_ in range(default__.safeDivisionInt(default__.fm0(globalState), d_2057_v0_), hi10_):
            d_2060_v2_: _dafny.MultiSet
            d_2060_v2_ = _dafny.MultiSet([default__.safeModuloInt(641, d_2057_v0_), d_2057_v0_])
            d_2060_v2_ = d_2060_v2_
            d_2061_v3_: bool
            d_2061_v3_ = True
            d_2061_v3_ = p1
            d_2062_v4_: _dafny.Seq
            d_2062_v4_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2063_v5_: _dafny.Seq
            d_2063_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocvmfl"))
            d_2064_v6_: _dafny.MultiSet
            d_2064_v6_ = _dafny.MultiSet([(d_2062_v4_)[default__.safeIndex(d_2059_i0_, len(d_2062_v4_))], (d_2059_i0_) == (len(d_2063_v5_)), (self).f10, (self).f10, p1])
            d_2065_v7_: D0
            d_2065_v7_ = D0_DC1(d_2057_v0_)
            d_2066_v8_: _dafny.Map
            d_2066_v8_ = _dafny.Map({d_2065_v7_: d_2061_v3_})
            d_2067_v9_: str
            d_2067_v9_ = _dafny.CodePoint('x')
            d_2068_v10_: _dafny.Array
            nw354_ = _dafny.Array(None, 16)
            nw354_[int(0)] = d_2061_v3_
            nw354_[int(1)] = (d_2059_i0_) <= (858)
            nw354_[int(2)] = p1
            nw354_[int(3)] = p1
            nw354_[int(4)] = (d_2057_v0_) == (d_2059_i0_)
            nw354_[int(5)] = p1
            nw354_[int(6)] = not(((d_2066_v8_)[d_2065_v7_] if (d_2065_v7_) in (d_2066_v8_) else d_2061_v3_))
            nw354_[int(7)] = not((self).f10)
            nw354_[int(8)] = d_2061_v3_
            nw354_[int(9)] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2069_i1_ in range(default__.abs(166))])).set(default__.safeIndex(d_2057_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2070_i1_ in range(default__.abs(166))]))), d_2067_v9_)) < (d_2063_v5_)
            nw354_[int(10)] = False
            nw354_[int(11)] = (self).f10
            nw354_[int(12)] = not((d_2061_v3_) == (not(d_2061_v3_)))
            nw354_[int(13)] = (d_2059_i0_) <= (d_2059_i0_)
            nw354_[int(14)] = (self).f10
            nw354_[int(15)] = p1
            d_2068_v10_ = nw354_
            rhs337_ = d_2064_v6_
            rhs338_ = d_2068_v10_
            d_2064_v6_ = rhs337_
            d_2068_v10_ = rhs338_
            d_2071_v11_: _dafny.Array
            nw355_ = _dafny.Array(D2.default()(), 26)
            d_2071_v11_ = nw355_
            d_2072_v12_: C4
            nw356_ = C4()
            nw356_.ctor__(d_2071_v11_, False)
            d_2072_v12_ = nw356_
        d_2073_v13_: _dafny.Seq
        d_2073_v13_ = _dafny.SeqWithoutIsStrInference([not(p1), (self).f10, p1])
        d_2073_v13_ = d_2073_v13_
        d_2074_i2_: int
        d_2074_i2_ = 0
        with _dafny.label("12"):
            while (self).f10:
                with _dafny.c_label("12"):
                    if (d_2074_i2_) >= (100):
                        raise _dafny.Break("12")
                    d_2074_i2_ = (d_2074_i2_) + (1)
                    d_2075_v14_: _dafny.Map
                    d_2075_v14_ = _dafny.Map({(self).f10: (0) - (d_2057_v0_)})
                    d_2076_v15_: _dafny.Map
                    d_2076_v15_ = _dafny.Map({((d_2075_v14_)[p1] if (p1) in (d_2075_v14_) else len(d_2073_v13_)): (self).f10})
                    if ((d_2076_v15_)[d_2057_v0_] if (d_2057_v0_) in (d_2076_v15_) else not(p1)):
                        d_2077_v16_: bool
                        d_2077_v16_ = False
                        d_2077_v16_ = (default__.safeModuloInt(d_2057_v0_, d_2057_v0_)) < (d_2057_v0_)
                        d_2078_v17_: _dafny.Array
                        def lambda124_(d_2079_i3_):
                            return not(False)

                        init69_ = lambda124_
                        nw357_ = _dafny.Array(None, 14)
                        for i0_69_ in range(nw357_.length(0)):
                            nw357_[i0_69_] = init69_(i0_69_)
                        d_2078_v17_ = nw357_
                        d_2080_v18_: _dafny.Array
                        nw358_ = _dafny.Array(None, 1)
                        nw358_[int(0)] = d_2078_v17_
                        d_2080_v18_ = nw358_
                        index334_ = default__.safeIndex(810, (d_2080_v18_).length(0))
                        (d_2080_v18_)[index334_] = d_2078_v17_
                        index335_ = default__.safeIndex(810, (d_2080_v18_).length(0))
                        (d_2080_v18_)[index335_] = d_2078_v17_
                        d_2081_v19_: _dafny.Set
                        d_2081_v19_ = _dafny.Set({(self).f10, d_2077_v16_})
                        d_2082_v20_: _dafny.Seq
                        d_2082_v20_ = _dafny.SeqWithoutIsStrInference([len(d_2081_v19_)])
                        d_2083_v21_: D1
                        d_2083_v21_ = D1_DC5(d_2077_v16_, d_2082_v20_, d_2057_v0_)
                        d_2084_v22_: _dafny.MultiSet
                        d_2084_v22_ = _dafny.MultiSet([d_2077_v16_, True, d_2077_v16_, False, False])
                        d_2085_v23_: _dafny.Array
                        nw359_ = _dafny.Array(None, 26)
                        nw359_[int(0)] = _dafny.SeqWithoutIsStrInference([d_2057_v0_, d_2057_v0_])
                        nw359_[int(1)] = (d_2083_v21_).cf10
                        nw359_[int(2)] = _dafny.SeqWithoutIsStrInference([d_2057_v0_ for d_2086_i4_ in range(default__.abs(-441))])
                        nw359_[int(3)] = d_2082_v20_
                        nw359_[int(4)] = d_2082_v20_
                        nw359_[int(5)] = d_2082_v20_
                        nw359_[int(6)] = d_2082_v20_
                        nw359_[int(7)] = _dafny.SeqWithoutIsStrInference([d_2057_v0_])
                        nw359_[int(8)] = d_2082_v20_
                        nw359_[int(9)] = (d_2082_v20_).set(default__.safeIndex(d_2057_v0_, len(d_2082_v20_)), d_2057_v0_)
                        nw359_[int(10)] = d_2082_v20_
                        nw359_[int(11)] = (d_2082_v20_).set(default__.safeIndex(len(d_2073_v13_), len(d_2082_v20_)), d_2057_v0_)
                        nw359_[int(12)] = _dafny.SeqWithoutIsStrInference([d_2057_v0_])
                        nw359_[int(13)] = d_2082_v20_
                        nw359_[int(14)] = _dafny.SeqWithoutIsStrInference([len(d_2082_v20_) for d_2087_i5_ in range(default__.abs(803))])
                        nw359_[int(15)] = (d_2082_v20_).set(default__.safeIndex(d_2057_v0_, len(d_2082_v20_)), d_2057_v0_)
                        nw359_[int(16)] = d_2082_v20_
                        nw359_[int(17)] = _dafny.SeqWithoutIsStrInference([d_2057_v0_])
                        nw359_[int(18)] = _dafny.SeqWithoutIsStrInference([d_2057_v0_ for d_2088_i6_ in range(default__.abs(856))])
                        nw359_[int(19)] = d_2082_v20_
                        nw359_[int(20)] = _dafny.SeqWithoutIsStrInference([(0) - (d_2057_v0_), (d_2084_v22_).cardinality, len(_dafny.SeqWithoutIsStrInference([d_2057_v0_ for d_2089_i7_ in range(default__.abs(969))]))])
                        nw359_[int(21)] = d_2082_v20_
                        nw359_[int(22)] = d_2082_v20_
                        nw359_[int(23)] = d_2082_v20_
                        nw359_[int(24)] = d_2082_v20_
                        nw359_[int(25)] = d_2082_v20_
                        d_2085_v23_ = nw359_
                        d_2090_v24_: _dafny.Map
                        d_2090_v24_ = _dafny.Map({(d_2057_v0_) - (d_2057_v0_): d_2085_v23_})
                        d_2091_v25_: _dafny.Map
                        d_2091_v25_ = _dafny.Map({d_2057_v0_: _dafny.Map({(self).f10: 149})})
                        d_2090_v24_ = (d_2090_v24_).set((len((((d_2091_v25_)[d_2057_v0_] if (d_2057_v0_) in (d_2091_v25_) else d_2075_v14_)).set((self).f10, d_2057_v0_))) - (d_2057_v0_), d_2085_v23_)
                        d_2092_v26_: _dafny.MultiSet
                        d_2092_v26_ = _dafny.MultiSet([847])
                        d_2093_v27_: C1
                        nw360_ = C1()
                        nw360_.ctor__((d_2092_v26_).isdisjoint(d_2092_v26_))
                        d_2093_v27_ = nw360_
                        d_2094_v28_: _dafny.Array
                        nw361_ = _dafny.Array(int(0), 29)
                        d_2094_v28_ = nw361_
                        d_2095_v29_: _dafny.Set
                        d_2095_v29_ = _dafny.Set({d_2094_v28_})
                        d_2077_v16_ = not (not(d_2077_v16_)) or ((d_2095_v29_).ispropersubset(d_2095_v29_))
                    elif True:
                        d_2057_v0_ = (0) - (default__.fm0(globalState))
                        d_2096_v30_: D0
                        d_2096_v30_ = D0_DC0(d_2057_v0_)
                        d_2097_v31_: _dafny.Array
                        nw362_ = _dafny.Array(None, 5)
                        nw362_[int(0)] = default__.fm1(_dafny.SeqWithoutIsStrInference([d_2096_v30_, d_2096_v30_]), d_2057_v0_, d_2057_v0_, globalState)
                        nw362_[int(1)] = p1
                        nw362_[int(2)] = (self).f10
                        nw362_[int(3)] = True
                        nw362_[int(4)] = p1
                        d_2097_v31_ = nw362_
                        d_2098_v32_: D3
                        d_2098_v32_ = D3_DC10(d_2097_v31_)
                        d_2099_v33_: _dafny.Seq
                        d_2099_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hv"))
                        d_2100_v34_: str
                        d_2100_v34_ = _dafny.CodePoint('u')
                        rhs339_ = default__.fm0(globalState)
                        rhs340_ = d_2098_v32_
                        rhs341_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jph")) if (self).f10 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "di")))
                        rhs342_ = d_2100_v34_
                        d_2057_v0_ = rhs339_
                        d_2098_v32_ = rhs340_
                        d_2099_v33_ = rhs341_
                        d_2100_v34_ = rhs342_
                        d_2101_v35_: _dafny.Map
                        d_2101_v35_ = _dafny.Map({d_2096_v30_: 650})
                        d_2099_v33_ = (d_2099_v33_) + ((D2_DC9(d_2057_v0_, len(d_2075_v14_), d_2099_v33_, (self).f10, d_2101_v35_)).cf18)
                        d_2102_v36_: bool
                        d_2102_v36_ = False
                        d_2102_v36_ = d_2102_v36_
                        d_2103_v37_: _dafny.Array
                        nw363_ = _dafny.Array(_dafny.Array(None, 0), 15)
                        d_2103_v37_ = nw363_
                        d_2104_v38_: _dafny.Map
                        d_2104_v38_ = _dafny.Map({(0) - (d_2057_v0_): d_2097_v31_})
                        d_2105_v39_: _dafny.Array
                        nw364_ = _dafny.Array(None, 8)
                        nw364_[int(0)] = d_2097_v31_
                        nw364_[int(1)] = d_2097_v31_
                        nw364_[int(2)] = d_2097_v31_
                        nw364_[int(3)] = d_2097_v31_
                        nw364_[int(4)] = d_2097_v31_
                        nw364_[int(5)] = d_2097_v31_
                        nw364_[int(6)] = d_2097_v31_
                        nw364_[int(7)] = ((d_2104_v38_)[908] if (908) in (d_2104_v38_) else d_2097_v31_)
                        d_2105_v39_ = nw364_
                        d_2106_v40_: _dafny.Map
                        d_2106_v40_ = _dafny.Map({d_2057_v0_: d_2105_v39_})
                        index336_ = default__.safeIndex(271, (d_2103_v37_).length(0))
                        (d_2103_v37_)[index336_] = ((d_2106_v40_)[d_2057_v0_] if (d_2057_v0_) in (d_2106_v40_) else d_2105_v39_)
                        index337_ = default__.safeIndex(271, (d_2103_v37_).length(0))
                        (d_2103_v37_)[index337_] = d_2105_v39_
                    d_2057_v0_ = (d_2057_v0_) * (d_2057_v0_)
                    d_2107_v41_: _dafny.Array
                    def lambda125_(d_2108_i8_):
                        return ((self).f10) or (True)

                    init70_ = lambda125_
                    nw365_ = _dafny.Array(None, 3)
                    for i0_70_ in range(nw365_.length(0)):
                        nw365_[i0_70_] = init70_(i0_70_)
                    d_2107_v41_ = nw365_
                    index338_ = default__.safeIndex(66, (d_2107_v41_).length(0))
                    (d_2107_v41_)[index338_] = ((self).f10) not in (d_2073_v13_)
                    index339_ = default__.safeIndex(66, (d_2107_v41_).length(0))
                    (d_2107_v41_)[index339_] = not(p1)
                    d_2109_v42_: D6
                    d_2109_v42_ = D6_DC16(False)
                    d_2110_v43_: _dafny.MultiSet
                    d_2110_v43_ = _dafny.MultiSet([(d_2109_v42_).cf30, (d_2073_v13_)[default__.safeIndex(d_2057_v0_, len(d_2073_v13_))]])
                    d_2111_v44_: D3
                    d_2111_v44_ = D3_DC11((d_2110_v43_).cardinality, d_2057_v0_, d_2057_v0_, (self).f10)
                    if (d_2111_v44_).cf25:
                        d_2112_v46_: _dafny.Array
                        def lambda126_(d_2113_v15_):
                            def lambda127_(d_2114_i9_):
                                def iife176_():
                                    coll46_ = _dafny.Set()
                                    compr_46_: int
                                    for compr_46_ in (_dafny.MultiSet([len(d_2113_v15_)])).Elements:
                                        d_2115_v45_: int = compr_46_
                                        if (d_2115_v45_) in (_dafny.MultiSet([len(d_2113_v15_)])):
                                            coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_2115_v45_, (_dafny.MultiSet([322])).cardinality)]))
                                    return _dafny.Set(coll46_)
                                return D2_DC7(iife176_()
)

                            return lambda127_

                        init71_ = lambda126_(d_2076_v15_)
                        nw366_ = _dafny.Array(None, 29)
                        for i0_71_ in range(nw366_.length(0)):
                            nw366_[i0_71_] = init71_(i0_71_)
                        d_2112_v46_ = nw366_
                        d_2116_v47_: C4
                        nw367_ = C4()
                        nw367_.ctor__((d_2112_v46_ if (d_2107_v41_)[default__.safeIndex(66, (d_2107_v41_).length(0))] else d_2112_v46_), (self).f10)
                        d_2116_v47_ = nw367_
                        d_2057_v0_ = (d_2057_v0_) + (d_2057_v0_)
                        d_2117_v48_: _dafny.Map
                        d_2117_v48_ = _dafny.Map({d_2057_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krvabqmix"))})
                        d_2118_v49_: _dafny.Seq
                        d_2118_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rr"))
                        d_2117_v48_ = (d_2117_v48_).set(d_2057_v0_, d_2118_v49_)
                        d_2057_v0_ = (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djnwyp"))), d_2057_v0_)) - (d_2057_v0_)
                        d_2057_v0_ = d_2057_v0_
                    elif True:
                        index340_ = default__.safeIndex(66, (d_2107_v41_).length(0))
                        (d_2107_v41_)[index340_] = p1
                        d_2057_v0_ = (d_2057_v0_) * (d_2057_v0_)
                        d_2119_v50_: D0
                        d_2119_v50_ = D0_DC0(d_2057_v0_)
                        d_2120_v51_: _dafny.Seq
                        d_2120_v51_ = _dafny.SeqWithoutIsStrInference([d_2119_v50_])
                        index341_ = default__.safeIndex(66, (d_2107_v41_).length(0))
                        (d_2107_v41_)[index341_] = default__.fm1(d_2120_v51_, d_2057_v0_, (d_2057_v0_) * (d_2057_v0_), globalState)
                        index342_ = default__.safeIndex(66, (d_2107_v41_).length(0))
                        (d_2107_v41_)[index342_] = p1
                        d_2057_v0_ = d_2057_v0_
                    pass
            pass
        d_2121_v52_: _dafny.Set
        d_2121_v52_ = _dafny.Set({p1, p1})
        d_2122_v53_: C3
        nw368_ = C3()
        nw368_.ctor__(len(d_2121_v52_), p1)
        d_2122_v53_ = nw368_
        d_2123_v54_: _dafny.MultiSet
        d_2123_v54_ = _dafny.MultiSet([d_2122_v53_])
        d_2124_v55_: _dafny.Map
        d_2124_v55_ = _dafny.Map({659: d_2123_v54_})
        hi11_ = (d_2122_v53_).f14
        for d_2125_i10_ in range((len(d_2124_v55_) if p1 else len(d_2073_v13_)), hi11_):
            d_2126_v56_: _dafny.MultiSet
            d_2126_v56_ = _dafny.MultiSet([d_2125_i10_])
            d_2127_v57_: T0
            nw369_ = C1()
            nw369_.ctor__(p1)
            d_2127_v57_ = nw369_
            d_2057_v0_ = default__.safeModuloInt((d_2122_v53_).f14, default__.safeModuloInt(d_2125_i10_, ((d_2126_v56_)[len(_dafny.Map({d_2127_v57_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2128_i11_ in range(default__.abs(566))])}))] if (len(_dafny.Map({d_2127_v57_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2129_i11_ in range(default__.abs(566))])}))) in (d_2126_v56_) else d_2057_v0_)))
            rhs343_ = not(d_2127_v57_.f7)
            rhs344_ = p1
            lhs214_ = d_2127_v57_
            lhs215_ = d_2127_v57_
            lhs214_.f7 = rhs343_
            lhs215_.f7 = rhs344_
            d_2130_v58_: _dafny.Seq
            d_2130_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmfjacoit"))
            d_2131_v59_: _dafny.Array
            nw370_ = _dafny.Array(int(0), 29)
            d_2131_v59_ = nw370_
            d_2132_v60_: D14
            d_2132_v60_ = D14_DC39(d_2130_v58_, d_2131_v59_)
            d_2133_v61_: C0
            nw371_ = C0()
            nw371_.ctor__(((d_2132_v60_).cf63) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egn"))), d_2127_v57_.f7)
            d_2133_v61_ = nw371_
            d_2130_v58_ = d_2130_v58_
        d_2134_v62_: bool
        d_2134_v62_ = True
        d_2135_v63_: _dafny.Seq
        d_2135_v63_ = _dafny.SeqWithoutIsStrInference([d_2057_v0_, (d_2122_v53_).f14])
        d_2136_v64_: _dafny.Map
        d_2136_v64_ = _dafny.Map({(d_2135_v63_) != (d_2135_v63_): False})
        d_2137_v65_: _dafny.MultiSet
        d_2137_v65_ = _dafny.MultiSet([p1])
        rhs345_ = not(not(((d_2136_v64_)[p1] if (p1) in (d_2136_v64_) else (d_2137_v65_).issubset(d_2137_v65_))))
        rhs346_ = d_2073_v13_
        d_2134_v62_ = rhs345_
        d_2073_v13_ = rhs346_
        d_2138_v66_: D0
        d_2138_v66_ = D0_DC0(len(d_2073_v13_))
        d_2139_v67_: _dafny.MultiSet
        d_2139_v67_ = _dafny.MultiSet([d_2057_v0_, d_2057_v0_])
        d_2140_v68_: _dafny.Seq
        d_2140_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxc"))
        d_2141_v69_: str
        d_2141_v69_ = _dafny.CodePoint('i')
        d_2142_v70_: _dafny.Map
        d_2142_v70_ = _dafny.Map({d_2141_v69_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ab")))})
        d_2143_v71_: _dafny.Map
        d_2143_v71_ = _dafny.Map({True: d_2142_v70_})
        d_2144_v73_: _dafny.MultiSet
        d_2144_v73_ = _dafny.MultiSet([d_2141_v69_, d_2141_v69_])
        d_2145_i12_: int
        d_2145_i12_ = 0
        with _dafny.label("13"):
            def iife177_():
                coll47_ = _dafny.Map()
                compr_47_: str
                for compr_47_ in (d_2144_v73_).Elements:
                    d_2187_v72_: str = compr_47_
                    if (d_2187_v72_) in (d_2144_v73_):
                        coll47_[d_2187_v72_] = (0) - (-746)
                return _dafny.Map(coll47_)
            while default__.fm1((_dafny.SeqWithoutIsStrInference([d_2138_v66_, d_2138_v66_])).set(default__.safeIndex((d_2139_v67_).cardinality, len(_dafny.SeqWithoutIsStrInference([d_2138_v66_, d_2138_v66_]))), D0_DC0(len(d_2140_v68_))), (0) - (d_2057_v0_), len(((d_2143_v71_)[True] if (True) in (d_2143_v71_) else iife177_()
            )), globalState):
                with _dafny.c_label("13"):
                    if (d_2145_i12_) >= (100):
                        raise _dafny.Break("13")
                    d_2145_i12_ = (d_2145_i12_) + (1)
                    d_2073_v13_ = (d_2073_v13_) + (d_2073_v13_)
                    d_2146_v74_: _dafny.Map
                    d_2146_v74_ = _dafny.Map({d_2140_v68_: False})
                    d_2147_v75_: _dafny.Map
                    d_2147_v75_ = _dafny.Map({((d_2122_v53_).f14 if p1 else len(d_2146_v74_)): (p1 if p1 else (self).f10)})
                    d_2147_v75_ = (d_2147_v75_).set((d_2122_v53_).f14, not(((self).f10) == ((self).f10)))
                    d_2057_v0_ = d_2057_v0_
                    d_2148_v76_: _dafny.Array
                    def lambda128_(d_2149_v62_):
                        def lambda129_(d_2150_i13_):
                            return d_2149_v62_

                        return lambda129_

                    init72_ = lambda128_(d_2134_v62_)
                    nw372_ = _dafny.Array(None, 27)
                    for i0_72_ in range(nw372_.length(0)):
                        nw372_[i0_72_] = init72_(i0_72_)
                    d_2148_v76_ = nw372_
                    d_2151_v77_: _dafny.Seq
                    d_2151_v77_ = _dafny.SeqWithoutIsStrInference([d_2148_v76_])
                    d_2152_v78_: D6
                    d_2152_v78_ = D6_DC15(default__.fm26(d_2134_v62_, len(d_2151_v77_), (0) - ((d_2122_v53_).f14), globalState))
                    d_2153_v79_: D6
                    d_2153_v79_ = D6_DC18(d_2152_v78_)
                    source16_ = D14_DC40(d_2153_v79_, not(d_2134_v62_), (self).f10)
                    if source16_.is_DC39:
                        d_2154___mcc_h0_ = source16_.cf63
                        d_2155___mcc_h1_ = source16_.cf64
                        d_2156_cf64_ = d_2155___mcc_h1_
                        d_2157_cf63_ = d_2154___mcc_h0_
                        d_2158_v80_: _dafny.Map
                        d_2158_v80_ = _dafny.Map({(d_2122_v53_).f14: d_2136_v64_})
                        d_2159_v81_: _dafny.Array
                        nw373_ = _dafny.Array(None, 15)
                        nw373_[int(0)] = d_2136_v64_
                        nw373_[int(1)] = default__.fm34(d_2057_v0_, p1, d_2134_v62_, (d_2122_v53_).f14, globalState)
                        nw373_[int(2)] = d_2136_v64_
                        nw373_[int(3)] = d_2136_v64_
                        nw373_[int(4)] = default__.fm34(309, (self).f10, (self).f10, (d_2137_v65_).cardinality, globalState)
                        nw373_[int(5)] = default__.fm34(376, (self).f10, (self).f10, d_2057_v0_, globalState)
                        nw373_[int(6)] = ((_dafny.Map({p1: d_2134_v62_})).set(p1, (self).f10)).set((d_2122_v53_).fm8(d_2057_v0_, globalState), p1)
                        nw373_[int(7)] = d_2136_v64_
                        nw373_[int(8)] = (d_2136_v64_) | (((d_2158_v80_)[len(d_2140_v68_)] if (len(d_2140_v68_)) in (d_2158_v80_) else d_2136_v64_))
                        nw373_[int(9)] = (d_2136_v64_).set(p1, d_2134_v62_)
                        nw373_[int(10)] = _dafny.Map({d_2134_v62_: True})
                        nw373_[int(11)] = _dafny.Map({d_2134_v62_: d_2134_v62_})
                        nw373_[int(12)] = d_2136_v64_
                        nw373_[int(13)] = d_2136_v64_
                        nw373_[int(14)] = d_2136_v64_
                        d_2159_v81_ = nw373_
                        index343_ = default__.safeIndex(267, (d_2159_v81_).length(0))
                        (d_2159_v81_)[index343_] = d_2136_v64_
                        index344_ = default__.safeIndex(267, (d_2159_v81_).length(0))
                        (d_2159_v81_)[index344_] = d_2136_v64_
                        d_2141_v69_ = d_2141_v69_
                        d_2134_v62_ = False
                        d_2136_v64_ = (d_2136_v64_) | (d_2136_v64_)
                    elif source16_.is_DC40:
                        d_2160___mcc_h2_ = source16_.cf65
                        d_2161___mcc_h3_ = source16_.cf66
                        d_2162___mcc_h4_ = source16_.cf67
                        d_2163_cf67_ = d_2162___mcc_h4_
                        d_2164_cf66_ = d_2161___mcc_h3_
                        d_2165_cf65_ = d_2160___mcc_h2_
                        d_2166_v82_: _dafny.MultiSet
                        d_2166_v82_ = _dafny.MultiSet([d_2140_v68_, d_2140_v68_, (d_2140_v68_) + (d_2140_v68_)])
                        d_2057_v0_ = ((d_2166_v82_)[d_2140_v68_] if (d_2140_v68_) in (d_2166_v82_) else (d_2122_v53_).f14)
                        d_2140_v68_ = d_2140_v68_
                        d_2057_v0_ = d_2057_v0_
                        d_2167_v83_: _dafny.Map
                        d_2167_v83_ = _dafny.Map({False: d_2073_v13_})
                        d_2137_v65_ = (_dafny.MultiSet((d_2073_v13_) + (((d_2167_v83_)[d_2163_cf67_] if (d_2163_cf67_) in (d_2167_v83_) else d_2073_v13_)))) - (d_2137_v65_)
                    elif source16_.is_DC38:
                        d_2168___mcc_h5_ = source16_.cf62
                        d_2169_cf62_ = d_2168___mcc_h5_
                        d_2170_v84_: _dafny.Array
                        nw374_ = _dafny.Array(_dafny.Seq({}), 18)
                        d_2170_v84_ = nw374_
                        d_2171_v85_: _dafny.Array
                        nw375_ = _dafny.Array(D6.default()(), 16)
                        d_2171_v85_ = nw375_
                        d_2172_v86_: _dafny.Seq
                        d_2172_v86_ = _dafny.SeqWithoutIsStrInference([d_2171_v85_, d_2171_v85_])
                        d_2173_v87_: _dafny.Seq
                        d_2173_v87_ = _dafny.SeqWithoutIsStrInference([d_2171_v85_, d_2171_v85_, d_2171_v85_, (d_2172_v86_)[default__.safeIndex((d_2122_v53_).f14, len(d_2172_v86_))], d_2171_v85_])
                        rhs347_ = d_2057_v0_
                        rhs348_ = d_2170_v84_
                        rhs349_ = (d_2173_v87_) + (d_2173_v87_)
                        rhs350_ = default__.fm0(globalState)
                        d_2057_v0_ = rhs347_
                        d_2170_v84_ = rhs348_
                        d_2172_v86_ = rhs349_
                        d_2057_v0_ = rhs350_
                        d_2057_v0_ = ((-261 if False else d_2057_v0_)) - (725)
                        d_2174_v88_: _dafny.Array
                        def lambda130_(d_2175_v53_):
                            def lambda131_(d_2176_i14_):
                                return (d_2176_i14_) + ((d_2175_v53_).f14)

                            return lambda131_

                        init73_ = lambda130_(d_2122_v53_)
                        nw376_ = _dafny.Array(None, 3)
                        for i0_73_ in range(nw376_.length(0)):
                            nw376_[i0_73_] = init73_(i0_73_)
                        d_2174_v88_ = nw376_
                        index345_ = default__.safeIndex(232, (d_2174_v88_).length(0))
                        (d_2174_v88_)[index345_] = (d_2122_v53_).f14
                        index346_ = default__.safeIndex(232, (d_2174_v88_).length(0))
                        (d_2174_v88_)[index346_] = (d_2122_v53_).f14
                        d_2177_v89_: _dafny.Map
                        d_2177_v89_ = _dafny.Map({default__.safeDivisionInt(d_2057_v0_, (d_2174_v88_)[default__.safeIndex(232, (d_2174_v88_).length(0))]): default__.fm0(globalState)})
                        d_2177_v89_ = (d_2177_v89_).set(default__.fm0(globalState), d_2057_v0_)
                    elif True:
                        d_2178___mcc_h6_ = source16_.cf68
                        d_2179_cf68_ = d_2178___mcc_h6_
                        index347_ = default__.safeIndex(959, (d_2148_v76_).length(0))
                        (d_2148_v76_)[index347_] = (self).f10
                        d_2180_v90_: _dafny.Map
                        d_2180_v90_ = _dafny.Map({p1: d_2073_v13_})
                        index348_ = default__.safeIndex(959, (d_2148_v76_).length(0))
                        (d_2148_v76_)[index348_] = ((default__.fm49(p1, d_2180_v90_, False, globalState)).set(p1, default__.abs(((d_2139_v67_)[(d_2122_v53_).f14] if ((d_2122_v53_).f14) in (d_2139_v67_) else (0) - ((d_2122_v53_).f14))))) == (d_2137_v65_)
                        index349_ = default__.safeIndex(959, (d_2148_v76_).length(0))
                        (d_2148_v76_)[index349_] = d_2134_v62_
                        d_2134_v62_ = (d_2148_v76_)[default__.safeIndex(959, (d_2148_v76_).length(0))]
                        d_2181_v91_: _dafny.Array
                        nw377_ = _dafny.Array(int(0), 16)
                        d_2181_v91_ = nw377_
                        d_2182_v92_: D16
                        d_2182_v92_ = D16_DC46((d_2122_v53_).f14, d_2181_v91_, d_2141_v69_)
                        d_2183_v93_: D9
                        d_2183_v93_ = D9_DC23(d_2181_v91_)
                        d_2184_v94_: _dafny.Seq
                        d_2184_v94_ = _dafny.SeqWithoutIsStrInference([d_2181_v91_, d_2181_v91_, d_2181_v91_, d_2181_v91_, d_2181_v91_])
                        d_2185_v95_: D14
                        d_2185_v95_ = D14_DC39(default__.fm20(d_2134_v62_, d_2134_v62_, globalState), d_2181_v91_)
                        d_2186_v96_: _dafny.Array
                        nw378_ = _dafny.Array(None, 28)
                        nw378_[int(0)] = d_2181_v91_
                        nw378_[int(1)] = d_2181_v91_
                        nw378_[int(2)] = d_2181_v91_
                        nw378_[int(3)] = d_2181_v91_
                        nw378_[int(4)] = d_2181_v91_
                        nw378_[int(5)] = d_2181_v91_
                        nw378_[int(6)] = d_2181_v91_
                        nw378_[int(7)] = d_2181_v91_
                        nw378_[int(8)] = d_2181_v91_
                        nw378_[int(9)] = d_2181_v91_
                        nw378_[int(10)] = d_2181_v91_
                        nw378_[int(11)] = d_2181_v91_
                        nw378_[int(12)] = (d_2182_v92_).cf74
                        nw378_[int(13)] = d_2181_v91_
                        nw378_[int(14)] = (d_2183_v93_).cf44
                        nw378_[int(15)] = d_2181_v91_
                        nw378_[int(16)] = d_2181_v91_
                        nw378_[int(17)] = d_2181_v91_
                        nw378_[int(18)] = d_2181_v91_
                        nw378_[int(19)] = (d_2184_v94_)[default__.safeIndex((d_2122_v53_).f14, len(d_2184_v94_))]
                        nw378_[int(20)] = d_2181_v91_
                        nw378_[int(21)] = d_2181_v91_
                        nw378_[int(22)] = d_2181_v91_
                        nw378_[int(23)] = d_2181_v91_
                        nw378_[int(24)] = d_2181_v91_
                        nw378_[int(25)] = d_2181_v91_
                        nw378_[int(26)] = d_2181_v91_
                        nw378_[int(27)] = (d_2185_v95_).cf64
                        d_2186_v96_ = nw378_
                        d_2186_v96_ = d_2186_v96_
                    pass
            pass
        pat_let_tv48_ = d_2057_v0_
        def iife178_(_pat_let65_0):
            def iife179_(d_2188_dt__update__tmp_h0_):
                def iife180_(_pat_let66_0):
                    def iife181_(d_2189_dt__update_hcf0_h0_):
                        return D0_DC0(d_2189_dt__update_hcf0_h0_)
                    return iife181_(_pat_let66_0)
                return iife180_(pat_let_tv48_)
            return iife179_(_pat_let65_0)
        r0 = iife178_(d_2138_v66_)
        return r0

    @property
    def f10(self):
        return self._f10

class C6(T1, T0):
    def  __init__(self):
        self._f7: bool = False
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def ctor__(self, f9, f7):
        (self)._f9 = f9
        (self).f7 = f7

    def fm8(self, p0, globalState):
        return self.f7

    def fm9(self, p0, p1, p2, globalState):
        def iife182_():
            coll48_ = _dafny.Map()
            compr_48_: int
            for compr_48_ in _dafny.IntegerRange(888, 769):
                d_2192_v0_: int = compr_48_
                if ((888) <= (d_2192_v0_)) and ((d_2192_v0_) < (769)):
                    coll48_[(d_2192_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxffjj"))))] = (self).f9
            return _dafny.Map(coll48_)
        return (default__.safeDivisionInt(101, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f9, (self).f9])).cardinality for d_2190_i1_ in range(default__.abs(-704))]) for d_2191_i0_ in range(default__.abs(390))])))) * (len((iife182_()
        ) | (_dafny.Map({(D0_DC1((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-992: 215}))]))))).cf1: self.f7}))))

    def fm12(self, p0, p1, p2, p3, globalState):
        return 254

    def m4(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        d_2193_v0_: _dafny.Array
        def lambda132_(d_2194_i0_):
            return self.f7

        init74_ = lambda132_
        nw379_ = _dafny.Array(None, 27)
        for i0_74_ in range(nw379_.length(0)):
            nw379_[i0_74_] = init74_(i0_74_)
        d_2193_v0_ = nw379_
        index350_ = default__.safeIndex(505, (d_2193_v0_).length(0))
        (d_2193_v0_)[index350_] = (self).fm8(p0, globalState)
        index351_ = default__.safeIndex(505, (d_2193_v0_).length(0))
        (d_2193_v0_)[index351_] = True
        d_2195_i1_: int
        d_2195_i1_ = 0
        with _dafny.label("14"):
            while (p1) != ((p1) + (p1)):
                with _dafny.c_label("14"):
                    if (d_2195_i1_) >= (100):
                        raise _dafny.Break("14")
                    d_2195_i1_ = (d_2195_i1_) + (1)
                    r1 = (0) - (-752)
                    d_2196_v1_: _dafny.Seq
                    d_2196_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0])])
                    d_2197_v2_: _dafny.Seq
                    d_2197_v2_ = _dafny.SeqWithoutIsStrInference([(d_2193_v0_)[default__.safeIndex(505, (d_2193_v0_).length(0))]])
                    d_2198_v3_: _dafny.MultiSet
                    d_2198_v3_ = _dafny.MultiSet([(d_2193_v0_)[default__.safeIndex(505, (d_2193_v0_).length(0))], (self).f9])
                    d_2199_v4_: _dafny.Seq
                    d_2199_v4_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_2200_v5_: _dafny.Seq
                    d_2200_v5_ = _dafny.SeqWithoutIsStrInference([len(d_2199_v4_)])
                    d_2201_v6_: _dafny.Map
                    d_2201_v6_ = _dafny.Map({p0: (self).f9})
                    d_2202_v7_: _dafny.Array
                    nw380_ = _dafny.Array(None, 18)
                    nw380_[int(0)] = len((d_2196_v1_) + ((d_2196_v1_).set(default__.safeIndex(p0, len(d_2196_v1_)), _dafny.SeqWithoutIsStrInference([len(d_2197_v2_), p0]))))
                    nw380_[int(1)] = (d_2198_v3_).cardinality
                    nw380_[int(2)] = (d_2200_v5_)[default__.safeIndex(p0, len(d_2200_v5_))]
                    nw380_[int(3)] = p0
                    nw380_[int(4)] = p0
                    nw380_[int(5)] = (self).fm12(self.f7, len(d_2201_v6_), p0, (d_2193_v0_)[default__.safeIndex(505, (d_2193_v0_).length(0))], globalState)
                    nw380_[int(6)] = p0
                    nw380_[int(7)] = (764) + (-606)
                    nw380_[int(8)] = default__.fm0(globalState)
                    nw380_[int(9)] = p0
                    nw380_[int(10)] = (self).fm9((d_2193_v0_)[default__.safeIndex(505, (d_2193_v0_).length(0))], p0, _dafny.CodePoint('g'), globalState)
                    nw380_[int(11)] = (p0) + (len(p1))
                    nw380_[int(12)] = len(p1)
                    nw380_[int(13)] = p0
                    nw380_[int(14)] = p0
                    nw380_[int(15)] = 611
                    nw380_[int(16)] = len(d_2201_v6_)
                    nw380_[int(17)] = default__.safeDivisionInt(len(d_2199_v4_), p0)
                    d_2202_v7_ = nw380_
                    index352_ = default__.safeIndex(301, (d_2202_v7_).length(0))
                    (d_2202_v7_)[index352_] = 560
                    index353_ = default__.safeIndex(301, (d_2202_v7_).length(0))
                    (d_2202_v7_)[index353_] = p0
                    d_2203_v8_: _dafny.Map
                    d_2203_v8_ = _dafny.Map({p0: p1})
                    d_2204_v9_: str
                    d_2204_v9_ = _dafny.CodePoint('w')
                    index354_ = default__.safeIndex(301, (d_2202_v7_).length(0))
                    (d_2202_v7_)[index354_] = len((((d_2203_v8_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydewlhgde")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydewlhgde")))) in (d_2203_v8_) else p1)) + (((p1).set(default__.safeIndex(default__.fm0(globalState), len(p1)), d_2204_v9_)).set(default__.safeIndex((d_2202_v7_)[default__.safeIndex(301, (d_2202_v7_).length(0))], len((p1).set(default__.safeIndex(default__.fm0(globalState), len(p1)), d_2204_v9_))), d_2204_v9_)))
                    (self).f7 = not(self.f7)
                    pass
            pass
        index355_ = default__.safeIndex(505, (d_2193_v0_).length(0))
        (d_2193_v0_)[index355_] = True
        d_2205_v10_: str
        d_2205_v10_ = _dafny.CodePoint('m')
        d_2206_v11_: _dafny.Seq
        d_2206_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), d_2205_v10_, d_2205_v10_])
        d_2206_v11_ = (_dafny.SeqWithoutIsStrInference([d_2205_v10_])) + (_dafny.SeqWithoutIsStrInference([default__.fm13((self).f9, p0, (self).f9, globalState)]))
        d_2207_v12_: _dafny.MultiSet
        d_2207_v12_ = _dafny.MultiSet([self.f7, True])
        index356_ = default__.safeIndex(505, (d_2193_v0_).length(0))
        (d_2193_v0_)[index356_] = ((d_2207_v12_).issubset(d_2207_v12_) if (self.f7 if (self).f9 else (self).f9) else (self).f9)
        d_2208_v13_: _dafny.Set
        d_2208_v13_ = _dafny.Set({len(_dafny.Set({len(p1), p0, p0})), p0})
        d_2209_i2_: int
        d_2209_i2_ = 0
        with _dafny.label("15"):
            while (len(d_2208_v13_)) == (p0):
                with _dafny.c_label("15"):
                    if (d_2209_i2_) >= (100):
                        raise _dafny.Break("15")
                    d_2209_i2_ = (d_2209_i2_) + (1)
                    d_2210_v14_: _dafny.Array
                    nw381_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
                    d_2210_v14_ = nw381_
                    index357_ = default__.safeIndex(315, (d_2210_v14_).length(0))
                    (d_2210_v14_)[index357_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlg"))
                    d_2211_v15_: _dafny.Seq
                    d_2211_v15_ = _dafny.SeqWithoutIsStrInference([d_2206_v11_, (p1) + (d_2206_v11_)])
                    d_2212_v16_: _dafny.Seq
                    d_2212_v16_ = _dafny.SeqWithoutIsStrInference([p0])
                    index358_ = default__.safeIndex(315, (d_2210_v14_).length(0))
                    (d_2210_v14_)[index358_] = (d_2211_v15_)[default__.safeIndex((len(d_2212_v16_) if self.f7 else p0), len(d_2211_v15_))]
                    d_2205_v10_ = _dafny.CodePoint('m')
                    d_2213_v17_: _dafny.Array
                    def lambda133_(d_2214_p0_):
                        def lambda134_(d_2215_i3_):
                            return (d_2215_i3_) * (len(_dafny.SeqWithoutIsStrInference([d_2214_p0_])))

                        return lambda134_

                    init75_ = lambda133_(p0)
                    nw382_ = _dafny.Array(None, 4)
                    for i0_75_ in range(nw382_.length(0)):
                        nw382_[i0_75_] = init75_(i0_75_)
                    d_2213_v17_ = nw382_
                    d_2216_v18_: _dafny.Set
                    d_2216_v18_ = _dafny.Set({d_2213_v17_})
                    d_2217_v19_: _dafny.Seq
                    d_2217_v19_ = _dafny.SeqWithoutIsStrInference([d_2216_v18_, d_2216_v18_])
                    d_2218_v21_: D0
                    def iife183_():
                        coll49_ = _dafny.Set()
                        compr_49_: int
                        for compr_49_ in _dafny.IntegerRange(287, 279):
                            d_2219_v20_: int = compr_49_
                            if ((287) <= (d_2219_v20_)) and ((d_2219_v20_) < (279)):
                                coll49_ = coll49_.union(_dafny.Set([(d_2219_v20_) + (p0)]))
                        return _dafny.Set(coll49_)
                    d_2218_v21_ = D0_DC2(self.f7, (0) - (p0), (d_2217_v19_)[default__.safeIndex(len(iife183_()
), len(d_2217_v19_))], (self).fm9(self.f7, p0, d_2205_v10_, globalState))
                    index359_ = default__.safeIndex(505, (d_2193_v0_).length(0))
                    (d_2193_v0_)[index359_] = (d_2218_v21_).cf2
                    index360_ = default__.safeIndex(505, (d_2193_v0_).length(0))
                    (d_2193_v0_)[index360_] = not((self).f9)
                    pass
            pass
        d_2220_v22_: _dafny.Array
        nw383_ = _dafny.Array(_dafny.Array(None, 0), 19)
        d_2220_v22_ = nw383_
        r0 = d_2220_v22_
        d_2221_v23_: _dafny.Array
        def lambda135_(d_2222_p0_):
            def lambda136_(d_2223_i4_):
                return (d_2223_i4_) + (len(_dafny.Map({(self).f9: len(_dafny.SeqWithoutIsStrInference([d_2222_p0_]))})))

            return lambda136_

        init76_ = lambda135_(p0)
        nw384_ = _dafny.Array(None, 6)
        for i0_76_ in range(nw384_.length(0)):
            nw384_[i0_76_] = init76_(i0_76_)
        d_2221_v23_ = nw384_
        d_2224_v24_: _dafny.Seq
        d_2224_v24_ = _dafny.SeqWithoutIsStrInference([d_2193_v0_])
        d_2225_v25_: _dafny.Map
        d_2225_v25_ = _dafny.Map({d_2221_v23_: len(d_2224_v24_)})
        r1 = (((self).fm12(self.f7, p0, p0, (d_2193_v0_)[default__.safeIndex(505, (d_2193_v0_).length(0))], globalState)) + (p0)) + ((len(d_2206_v11_)) * (len(d_2225_v25_)))
        r2 = (d_2193_v0_)[default__.safeIndex(505, (d_2193_v0_).length(0))]
        r3 = (len(p1)) * (p0)
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Set = _dafny.Set({})
        (self).f7 = p2
        d_2226_v0_: _dafny.Seq
        d_2226_v0_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajaosb"))])
        d_2227_v1_: _dafny.Seq
        d_2227_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onrgrefdt"))
        hi12_ = (p1) + (p1)
        for d_2228_i0_ in range(len(((d_2226_v0_)[default__.safeIndex(p1, len(d_2226_v0_))]) + (d_2227_v1_)), hi12_):
            d_2229_v2_: _dafny.Array
            def lambda137_(d_2230_i1_):
                return (self).f9

            init77_ = lambda137_
            nw385_ = _dafny.Array(None, 25)
            for i0_77_ in range(nw385_.length(0)):
                nw385_[i0_77_] = init77_(i0_77_)
            d_2229_v2_ = nw385_
            index361_ = default__.safeIndex(602, (d_2229_v2_).length(0))
            (d_2229_v2_)[index361_] = self.f7
            d_2231_v3_: str
            d_2231_v3_ = _dafny.CodePoint('c')
            index362_ = default__.safeIndex(602, (d_2229_v2_).length(0))
            rhs351_ = p0
            rhs352_ = d_2231_v3_
            lhs216_ = d_2229_v2_
            lhs217_ = default__.safeIndex(602, (d_2229_v2_).length(0))
            lhs216_[lhs217_] = rhs351_
            d_2231_v3_ = rhs352_
            d_2232_v4_: _dafny.Map
            d_2232_v4_ = _dafny.Map({(d_2229_v2_)[default__.safeIndex(602, (d_2229_v2_).length(0))]: (d_2229_v2_)[default__.safeIndex(602, (d_2229_v2_).length(0))]})
            d_2233_v5_: C0
            nw386_ = C0()
            nw386_.ctor__(p0, (((d_2232_v4_)[p0] if (p0) in (d_2232_v4_) else (self).f9) if (self).f9 else p2))
            d_2233_v5_ = nw386_
            d_2234_v6_: int
            d_2234_v6_ = 88
            d_2234_v6_ = d_2228_i0_
            d_2234_v6_ = ((d_2234_v6_) * (d_2234_v6_)) - (default__.safeModuloInt(d_2228_i0_, d_2234_v6_))
        d_2235_v7_: _dafny.Array
        nw387_ = _dafny.Array(D14.default()(), 2)
        d_2235_v7_ = nw387_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_2235_v7_).length(0)):
            d_2236_i2_: int = guard_loop_4_
            if (True) and (((0) <= (d_2236_i2_)) and ((d_2236_i2_) < ((d_2235_v7_).length(0)))):
                (d_2235_v7_)[(d_2236_i2_)] = (D14_DC40(D6_DC18(D6_DC15(_dafny.CodePoint('o'))), p0, p2) if p2 else D14_DC40(D6_DC18(D6_DC15(_dafny.CodePoint('i'))), self.f7, True))
        if (self).fm8(p1, globalState):
            d_2237_v8_: int
            d_2237_v8_ = 754
            d_2238_v9_: _dafny.Seq
            d_2238_v9_ = _dafny.SeqWithoutIsStrInference([d_2237_v8_])
            d_2237_v8_ = (d_2238_v9_)[default__.safeIndex(p1, len(d_2238_v9_))]
            d_2239_v10_: _dafny.Seq
            d_2239_v10_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2237_v8_ = len((d_2239_v10_).set(default__.safeIndex((p1) * (len(_dafny.Map({self.f7: False}))), len(d_2239_v10_)), (False) == (p2)))
            d_2240_v11_: str
            d_2240_v11_ = _dafny.CodePoint('k')
            d_2241_v12_: D6
            d_2241_v12_ = D6_DC15(d_2240_v11_)
            d_2241_v12_ = D6_DC15((d_2227_v1_)[default__.safeIndex(len(d_2227_v1_), len(d_2227_v1_))])
            d_2242_v13_: _dafny.Array
            nw388_ = _dafny.Array(False, 4)
            d_2242_v13_ = nw388_
            d_2243_v14_: _dafny.Set
            d_2243_v14_ = _dafny.Set({d_2242_v13_, d_2242_v13_})
            d_2244_v15_: C2
            nw389_ = C2()
            nw389_.ctor__(d_2243_v14_, d_2240_v11_, p2)
            d_2244_v15_ = nw389_
            d_2237_v8_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkdxg")))
        elif True:
            d_2245_v16_: int
            d_2245_v16_ = 619
            d_2245_v16_ = p1
            d_2246_v17_: _dafny.Map
            d_2246_v17_ = _dafny.Map({p1: d_2227_v1_})
            d_2247_v18_: _dafny.Seq
            d_2247_v18_ = _dafny.SeqWithoutIsStrInference([False, p0])
            d_2248_v19_: D0
            d_2248_v19_ = D0_DC0(p1)
            d_2249_v20_: _dafny.Seq
            d_2249_v20_ = _dafny.SeqWithoutIsStrInference([d_2248_v19_, d_2248_v19_, d_2248_v19_, d_2248_v19_, d_2248_v19_])
            d_2250_v21_: _dafny.Seq
            d_2250_v21_ = _dafny.SeqWithoutIsStrInference([p1, d_2245_v16_])
            rhs353_ = ((d_2245_v16_) <= (len(_dafny.SeqWithoutIsStrInference([(d_2247_v18_)[default__.safeIndex(p1, len(d_2247_v18_))]])))) == ((self).f9)
            rhs354_ = ((self).fm12(p0, d_2245_v16_, 478, p2, globalState)) - (p1)
            rhs355_ = (_dafny.SeqWithoutIsStrInference([default__.fm1((d_2249_v20_).set(default__.safeIndex(len(d_2250_v21_), len(d_2249_v20_)), D0_DC0(len(d_2227_v1_))), p1, d_2245_v16_, globalState)])) < (d_2247_v18_)
            rhs356_ = ((d_2246_v17_) | (_dafny.Map({p1: d_2227_v1_}))) | (_dafny.Map({d_2245_v16_: d_2227_v1_}))
            lhs218_ = self
            lhs219_ = self
            lhs218_.f7 = rhs353_
            d_2245_v16_ = rhs354_
            lhs219_.f7 = rhs355_
            d_2246_v17_ = rhs356_
            d_2245_v16_ = 635
            d_2245_v16_ = d_2245_v16_
            d_2251_v22_: _dafny.Array
            def lambda138_(d_2252_p1_):
                def lambda139_(d_2253_i3_):
                    return (d_2253_i3_) + (d_2252_p1_)

                return lambda139_

            init78_ = lambda138_(p1)
            nw390_ = _dafny.Array(None, 14)
            for i0_78_ in range(nw390_.length(0)):
                nw390_[i0_78_] = init78_(i0_78_)
            d_2251_v22_ = nw390_
            index363_ = default__.safeIndex(816, (d_2251_v22_).length(0))
            (d_2251_v22_)[index363_] = (_dafny.MultiSet(d_2250_v21_)).cardinality
            index364_ = default__.safeIndex(816, (d_2251_v22_).length(0))
            (d_2251_v22_)[index364_] = len(((d_2246_v17_) | (d_2246_v17_)) | (d_2246_v17_))
        if p0:
            d_2254_v23_: D13
            d_2254_v23_ = D13_DC35()
            d_2255_v24_: _dafny.Seq
            d_2255_v24_ = _dafny.SeqWithoutIsStrInference([d_2254_v23_])
            d_2256_v25_: _dafny.Map
            d_2256_v25_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([D13_DC35() for d_2257_i4_ in range(default__.abs(-987))])) + (d_2255_v24_): p1})
            d_2258_v26_: str
            d_2258_v26_ = _dafny.CodePoint('u')
            d_2259_v27_: _dafny.Map
            d_2259_v27_ = _dafny.Map({p0: d_2258_v26_})
            d_2256_v25_ = (d_2256_v25_).set(d_2255_v24_, len(d_2259_v27_))
            d_2260_v28_: _dafny.Array
            nw391_ = _dafny.Array(_dafny.Set({}), 22)
            d_2260_v28_ = nw391_
            d_2261_v29_: _dafny.Seq
            d_2261_v29_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_2262_v30_: _dafny.Set
            d_2262_v30_ = _dafny.Set({d_2261_v29_})
            index365_ = default__.safeIndex(557, (d_2260_v28_).length(0))
            (d_2260_v28_)[index365_] = (d_2262_v30_) - (d_2262_v30_)
            d_2263_v31_: _dafny.Seq
            d_2263_v31_ = _dafny.SeqWithoutIsStrInference([p2, default__.fm1(_dafny.SeqWithoutIsStrInference([D0_DC0(p1) for d_2264_i5_ in range(default__.abs(693))]), p1, p1, globalState), (self).f9])
            d_2265_v32_: _dafny.Map
            d_2265_v32_ = _dafny.Map({d_2261_v29_: d_2263_v31_})
            index366_ = default__.safeIndex(557, (d_2260_v28_).length(0))
            def iife184_():
                coll50_ = _dafny.Set()
                compr_50_: _dafny.Seq
                for compr_50_ in (d_2265_v32_).keys.Elements:
                    d_2266_v33_: _dafny.Seq = compr_50_
                    if (d_2266_v33_) in (d_2265_v32_):
                        coll50_ = coll50_.union(_dafny.Set([d_2266_v33_]))
                return _dafny.Set(coll50_)
            (d_2260_v28_)[index366_] = (_dafny.Set({d_2261_v29_, _dafny.SeqWithoutIsStrInference([len(d_2263_v31_), p1]), d_2261_v29_})).intersection(iife184_()
            )
            d_2267_v34_: _dafny.Array
            nw392_ = _dafny.Array(int(0), 22)
            d_2267_v34_ = nw392_
            index367_ = default__.safeIndex(965, (d_2267_v34_).length(0))
            (d_2267_v34_)[index367_] = p1
            index368_ = default__.safeIndex(965, (d_2267_v34_).length(0))
            (d_2267_v34_)[index368_] = default__.safeModuloInt(p1, p1)
            index369_ = default__.safeIndex(965, (d_2267_v34_).length(0))
            (d_2267_v34_)[index369_] = p1
            index370_ = default__.safeIndex(965, (d_2267_v34_).length(0))
            (d_2267_v34_)[index370_] = (default__.safeDivisionInt((self).fm9((self).f9, (0) - ((d_2267_v34_)[default__.safeIndex(965, (d_2267_v34_).length(0))]), d_2258_v26_, globalState), (d_2267_v34_)[default__.safeIndex(965, (d_2267_v34_).length(0))]) if (self).f9 else (0) - ((d_2267_v34_)[default__.safeIndex(965, (d_2267_v34_).length(0))]))
        elif True:
            (self).f7 = p2
            d_2268_v35_: _dafny.Map
            d_2268_v35_ = _dafny.Map({769: _dafny.MultiSet([self.f7])})
            d_2269_v36_: _dafny.MultiSet
            d_2269_v36_ = _dafny.MultiSet([not(self.f7), p0])
            if not(((_dafny.MultiSet([(self).f9, True])).ispropersubset(((d_2268_v35_)[-296] if (-296) in (d_2268_v35_) else d_2269_v36_))) or ((self).f9)):
                d_2270_v37_: _dafny.Map
                d_2270_v37_ = _dafny.Map({not(p2): p1})
                d_2271_v38_: _dafny.Map
                d_2271_v38_ = _dafny.Map({p1: d_2270_v37_})
                d_2272_v39_: _dafny.Array
                nw393_ = _dafny.Array(False, 6)
                d_2272_v39_ = nw393_
                d_2273_v40_: _dafny.Seq
                d_2273_v40_ = _dafny.SeqWithoutIsStrInference([self.f7])
                d_2274_v41_: _dafny.Map
                d_2274_v41_ = _dafny.Map({d_2272_v39_: d_2273_v40_})
                d_2275_v42_: _dafny.MultiSet
                d_2275_v42_ = _dafny.MultiSet([len((d_2274_v41_) | (_dafny.Map({d_2272_v39_: d_2273_v40_}))), p1, ((_dafny.MultiSet([p2, (self).f9]) if p2 else _dafny.MultiSet([self.f7, p0, p0, p0, self.f7]))).cardinality, p1])
                d_2276_v43_: str
                d_2276_v43_ = _dafny.CodePoint('g')
                d_2277_v44_: _dafny.Set
                d_2277_v44_ = _dafny.Set({p1})
                d_2278_v45_: _dafny.Set
                d_2278_v45_ = _dafny.Set({p1, len(d_2277_v44_), p1, len(d_2227_v1_), p1})
                d_2279_v46_: _dafny.Map
                d_2279_v46_ = _dafny.Map({p1: d_2227_v1_})
                d_2280_v47_: _dafny.Map
                d_2280_v47_ = _dafny.Map({len(((d_2279_v46_)[p1] if (p1) in (d_2279_v46_) else d_2227_v1_)): (self).f9})
                d_2281_v48_: _dafny.Seq
                d_2281_v48_ = _dafny.SeqWithoutIsStrInference([len(d_2227_v1_)])
                d_2282_v49_: _dafny.Array
                nw394_ = _dafny.Array(None, 15)
                nw394_[int(0)] = p1
                nw394_[int(1)] = p1
                nw394_[int(2)] = p1
                nw394_[int(3)] = p1
                nw394_[int(4)] = p1
                nw394_[int(5)] = p1
                nw394_[int(6)] = p1
                nw394_[int(7)] = p1
                nw394_[int(8)] = p1
                nw394_[int(9)] = p1
                nw394_[int(10)] = p1
                nw394_[int(11)] = (_dafny.MultiSet([default__.fm0(globalState)])).cardinality
                nw394_[int(12)] = p1
                nw394_[int(13)] = (d_2281_v48_)[default__.safeIndex(548, len(d_2281_v48_))]
                nw394_[int(14)] = p1
                d_2282_v49_ = nw394_
                d_2283_v50_: _dafny.Map
                d_2283_v50_ = _dafny.Map({False: d_2282_v49_})
                d_2284_v51_: _dafny.Array
                nw395_ = _dafny.Array(None, 27)
                nw395_[int(0)] = len(d_2227_v1_)
                nw395_[int(1)] = p1
                nw395_[int(2)] = p1
                nw395_[int(3)] = p1
                nw395_[int(4)] = p1
                nw395_[int(5)] = p1
                nw395_[int(6)] = p1
                nw395_[int(7)] = p1
                nw395_[int(8)] = p1
                nw395_[int(9)] = p1
                nw395_[int(10)] = p1
                nw395_[int(11)] = p1
                nw395_[int(12)] = p1
                nw395_[int(13)] = p1
                nw395_[int(14)] = p1
                nw395_[int(15)] = len((d_2227_v1_).set(default__.safeIndex(len(d_2273_v40_), len(d_2227_v1_)), d_2276_v43_))
                nw395_[int(16)] = (0) - (p1)
                nw395_[int(17)] = p1
                nw395_[int(18)] = p1
                nw395_[int(19)] = p1
                nw395_[int(20)] = p1
                nw395_[int(21)] = len(d_2277_v44_)
                nw395_[int(22)] = 153
                nw395_[int(23)] = ((d_2275_v42_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkp"))), default__.abs(len(d_2280_v47_)))).cardinality
                nw395_[int(24)] = (0) - ((d_2275_v42_).cardinality)
                nw395_[int(25)] = 441
                nw395_[int(26)] = len(d_2283_v50_)
                d_2284_v51_ = nw395_
                d_2285_v52_: _dafny.MultiSet
                d_2285_v52_ = _dafny.MultiSet([d_2284_v51_, d_2282_v49_, d_2284_v51_])
                d_2286_v53_: int
                d_2286_v53_ = 882
                d_2287_v54_: _dafny.Map
                d_2287_v54_ = _dafny.Map({p1: d_2284_v51_})
                d_2288_v55_: D17
                d_2288_v55_ = D17_DC47(d_2285_v52_)
                rhs357_ = (d_2271_v38_).set(len(d_2287_v54_), d_2270_v37_)
                rhs358_ = d_2275_v42_
                rhs359_ = (((d_2288_v55_).cf76) - (d_2285_v52_)).set(d_2284_v51_, default__.abs(d_2286_v53_))
                rhs360_ = p1
                d_2271_v38_ = rhs357_
                d_2275_v42_ = rhs358_
                d_2285_v52_ = rhs359_
                d_2286_v53_ = rhs360_
                default__.m0(default__.safeDivisionInt((0) - (len(d_2273_v40_)), d_2286_v53_), d_2286_v53_, globalState)
                d_2286_v53_ = (d_2286_v53_) * (821)
                d_2289_v56_: _dafny.Array
                nw396_ = _dafny.Array(_dafny.Set({}), 9)
                d_2289_v56_ = nw396_
                d_2290_v57_: _dafny.Map
                d_2290_v57_ = _dafny.Map({default__.fm0(globalState): d_2289_v56_})
                d_2290_v57_ = (d_2290_v57_).set(p1, d_2289_v56_)
                d_2291_v58_: _dafny.Array
                nw397_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_2291_v58_ = nw397_
                index371_ = default__.safeIndex(87, (d_2291_v58_).length(0))
                (d_2291_v58_)[index371_] = default__.fm28(p2, d_2286_v53_, globalState)
                index372_ = default__.safeIndex(87, (d_2291_v58_).length(0))
                (d_2291_v58_)[index372_] = _dafny.SeqWithoutIsStrInference([d_2276_v43_ for d_2292_i6_ in range(default__.abs(819))])
            elif True:
                d_2293_v59_: _dafny.Map
                d_2293_v59_ = _dafny.Map({(p1) == (323): -596})
                d_2293_v59_ = (d_2293_v59_).set(self.f7, 147)
                d_2294_v60_: C3
                nw398_ = C3()
                nw398_.ctor__(p1, p2)
                d_2294_v60_ = nw398_
                d_2295_v61_: int
                d_2295_v61_ = 569
                d_2296_v62_: T1
                nw399_ = C3()
                nw399_.ctor__(p1, p0)
                d_2296_v62_ = nw399_
                d_2297_v63_: D18
                d_2297_v63_ = D18_DC50(d_2296_v62_)
                rhs361_ = d_2294_v60_
                rhs362_ = p1
                rhs363_ = (d_2295_v61_) >= (d_2295_v61_)
                rhs364_ = (d_2297_v63_).cf84
                rhs365_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hulf"))) + (d_2227_v1_)
                lhs220_ = self
                d_2294_v60_ = rhs361_
                d_2295_v61_ = rhs362_
                lhs220_.f7 = rhs363_
                d_2296_v62_ = rhs364_
                d_2227_v1_ = rhs365_
                d_2298_v64_: _dafny.Array
                nw400_ = _dafny.Array(int(0), 8)
                d_2298_v64_ = nw400_
                d_2299_v65_: _dafny.Seq
                d_2299_v65_ = _dafny.SeqWithoutIsStrInference([(d_2294_v60_).f14, p1])
                d_2300_v66_: D12
                d_2300_v66_ = D12_DC30(d_2295_v61_, not(d_2296_v62_.f7), d_2299_v65_, (d_2294_v60_).f14)
                d_2301_v67_: _dafny.Seq
                d_2301_v67_ = _dafny.SeqWithoutIsStrInference([d_2296_v62_.f7, p0, (d_2300_v66_).cf50, p2, d_2296_v62_.f7])
                d_2302_v68_: _dafny.Set
                d_2302_v68_ = _dafny.Set({(d_2301_v67_)[default__.safeIndex(p1, len(d_2301_v67_))]})
                d_2303_v69_: _dafny.MultiSet
                d_2303_v69_ = _dafny.MultiSet([p1, (d_2294_v60_).f14, (0) - (len(d_2302_v68_)), 100, d_2295_v61_])
                d_2304_v70_: _dafny.Array
                nw401_ = _dafny.Array(False, 8)
                d_2304_v70_ = nw401_
                d_2305_v71_: _dafny.Set
                d_2305_v71_ = _dafny.Set({d_2304_v70_})
                index373_ = default__.safeIndex(548, (d_2298_v64_).length(0))
                (d_2298_v64_)[index373_] = (((d_2303_v69_)[len(d_2305_v71_)] if (len(d_2305_v71_)) in (d_2303_v69_) else p1)) + (p1)
                index374_ = default__.safeIndex(548, (d_2298_v64_).length(0))
                (d_2298_v64_)[index374_] = (0) - ((d_2294_v60_).f14)
                (d_2296_v62_).f7 = not(d_2296_v62_.f7)
                d_2306_v72_: _dafny.Set
                d_2306_v72_ = _dafny.Set({(d_2294_v60_).f14, p1, (d_2294_v60_).f14})
                d_2307_v73_: _dafny.Seq
                d_2307_v73_ = _dafny.SeqWithoutIsStrInference([default__.fm2(d_2295_v61_, d_2296_v62_.f7, globalState), (d_2306_v72_) - (_dafny.Set({(_dafny.MultiSet([d_2296_v62_, d_2296_v62_])).cardinality}))])
                d_2308_v74_: _dafny.Array
                def lambda140_(d_2309_v67_):
                    def lambda141_(d_2310_i7_):
                        return d_2309_v67_

                    return lambda141_

                init79_ = lambda140_(d_2301_v67_)
                nw402_ = _dafny.Array(None, 1)
                for i0_79_ in range(nw402_.length(0)):
                    nw402_[i0_79_] = init79_(i0_79_)
                d_2308_v74_ = nw402_
                index375_ = default__.safeIndex(843, (d_2308_v74_).length(0))
                (d_2308_v74_)[index375_] = d_2301_v67_
                d_2311_v75_: _dafny.MultiSet
                d_2311_v75_ = _dafny.MultiSet([d_2304_v70_, d_2304_v70_, d_2304_v70_, d_2304_v70_])
                d_2312_v76_: _dafny.Map
                d_2312_v76_ = d_2268_v35_
                d_2313_v77_: _dafny.Seq
                d_2313_v77_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, (d_2294_v60_).fm8(d_2295_v61_, globalState), self.f7])])
                index376_ = default__.safeIndex(843, (d_2308_v74_).length(0))
                rhs366_ = default__.fm54(globalState)
                rhs367_ = (d_2313_v77_)[default__.safeIndex((d_2294_v60_).f14, len(d_2313_v77_))]
                rhs368_ = d_2311_v75_
                rhs369_ = d_2312_v76_
                rhs370_ = d_2298_v64_
                lhs221_ = d_2308_v74_
                lhs222_ = default__.safeIndex(843, (d_2308_v74_).length(0))
                d_2307_v73_ = rhs366_
                lhs221_[lhs222_] = rhs367_
                d_2311_v75_ = rhs368_
                d_2312_v76_ = rhs369_
                d_2298_v64_ = rhs370_
            d_2314_v78_: C3
            nw403_ = C3()
            nw403_.ctor__(962, not(p0))
            d_2314_v78_ = nw403_
            d_2315_v79_: str
            d_2315_v79_ = _dafny.CodePoint('k')
            d_2316_v80_: _dafny.Set
            d_2316_v80_ = _dafny.Set({p2})
            d_2317_v81_: C3
            nw404_ = C3()
            nw404_.ctor__(p1, (d_2316_v80_).ispropersubset(default__.fm27(self.f7, d_2315_v79_, p2, p0, globalState)))
            d_2317_v81_ = nw404_
            d_2318_v82_: _dafny.Array
            nw405_ = _dafny.Array(False, 18)
            d_2318_v82_ = nw405_
            d_2319_v83_: _dafny.Map
            d_2319_v83_ = _dafny.Map({(self).f9: d_2318_v82_})
            d_2318_v82_ = ((d_2319_v83_)[self.f7] if (self.f7) in (d_2319_v83_) else d_2318_v82_)
        d_2320_v84_: _dafny.Set
        d_2320_v84_ = _dafny.Set({False})
        r1 = (d_2320_v84_) - ((d_2320_v84_) | (d_2320_v84_))
        d_2321_v85_: _dafny.Map
        d_2321_v85_ = _dafny.Map({False: p2})
        r0 = (d_2321_v85_) | (d_2321_v85_)
        d_2322_v86_: str
        d_2322_v86_ = _dafny.CodePoint('i')
        d_2323_v87_: D1
        d_2323_v87_ = D1_DC3(d_2320_v84_)
        r1 = ((d_2323_v87_).cf6 if (d_2322_v86_) in (d_2227_v1_) else (d_2320_v84_).intersection(d_2320_v84_))
        return r0, r1

    def m3(self, globalState):
        if self.f7:
            d_2324_v0_: int
            d_2324_v0_ = -390
            d_2325_v1_: _dafny.Array
            nw406_ = _dafny.Array(None, 4)
            nw406_[int(0)] = self.f7
            nw406_[int(1)] = (self).f9
            nw406_[int(2)] = self.f7
            nw406_[int(3)] = False
            d_2325_v1_ = nw406_
            d_2326_v2_: _dafny.Set
            d_2326_v2_ = _dafny.Set({d_2325_v1_, d_2325_v1_, d_2325_v1_, d_2325_v1_, d_2325_v1_})
            d_2327_v3_: str
            d_2327_v3_ = _dafny.CodePoint('t')
            d_2328_v4_: T0
            nw407_ = C2()
            nw407_.ctor__(d_2326_v2_, d_2327_v3_, (self).f9)
            d_2328_v4_ = nw407_
            d_2329_v5_: _dafny.Seq
            d_2329_v5_ = _dafny.SeqWithoutIsStrInference([d_2328_v4_])
            d_2330_v6_: _dafny.Seq
            d_2330_v6_ = _dafny.SeqWithoutIsStrInference([len(d_2329_v5_), 127])
            d_2331_v7_: _dafny.Array
            def lambda142_(d_2332_v0_):
                def lambda143_(d_2333_i0_):
                    return default__.safeDivisionInt(d_2333_i0_, d_2332_v0_)

                return lambda143_

            init80_ = lambda142_(d_2324_v0_)
            nw408_ = _dafny.Array(None, 5)
            for i0_80_ in range(nw408_.length(0)):
                nw408_[i0_80_] = init80_(i0_80_)
            d_2331_v7_ = nw408_
            d_2334_v8_: _dafny.Map
            d_2334_v8_ = _dafny.Map({d_2324_v0_: d_2331_v7_})
            d_2335_v9_: _dafny.Map
            d_2335_v9_ = _dafny.Map({(d_2324_v0_) < (d_2324_v0_): default__.safeDivisionInt(len(d_2330_v6_), len(d_2334_v8_))})
            d_2335_v9_ = d_2335_v9_
            d_2336_v10_: _dafny.Array
            nw409_ = _dafny.Array(_dafny.CodePoint('D'), 1)
            d_2336_v10_ = nw409_
            index377_ = default__.safeIndex(358, (d_2336_v10_).length(0))
            (d_2336_v10_)[index377_] = (_dafny.CodePoint('r') if self.f7 else _dafny.CodePoint('o'))
            d_2337_v11_: _dafny.Seq
            d_2337_v11_ = _dafny.SeqWithoutIsStrInference([d_2328_v4_.f7])
            d_2338_v12_: _dafny.Set
            d_2338_v12_ = _dafny.Set({(self).f9, (d_2337_v11_) < (d_2337_v11_)})
            index378_ = default__.safeIndex(358, (d_2336_v10_).length(0))
            rhs371_ = default__.fm13(d_2328_v4_.f7, d_2324_v0_, (self).f9, globalState)
            rhs372_ = d_2338_v12_
            lhs223_ = d_2336_v10_
            lhs224_ = default__.safeIndex(358, (d_2336_v10_).length(0))
            lhs223_[lhs224_] = rhs371_
            d_2338_v12_ = rhs372_
            d_2339_v13_: _dafny.MultiSet
            d_2339_v13_ = _dafny.MultiSet([d_2328_v4_.f7])
            d_2340_v14_: D9
            d_2340_v14_ = D9_DC26(_dafny.SeqWithoutIsStrInference([(d_2339_v13_).cardinality]))
            pat_let_tv49_ = d_2330_v6_
            def iife185_(_pat_let67_0):
                def iife186_(d_2341_dt__update__tmp_h0_):
                    def iife187_(_pat_let68_0):
                        def iife188_(d_2342_dt__update_hcf45_h0_):
                            return D9_DC26(d_2342_dt__update_hcf45_h0_)
                        return iife188_(_pat_let68_0)
                    return iife187_(pat_let_tv49_)
                return iife186_(_pat_let67_0)
            rhs373_ = (self).f9
            rhs374_ = (self).f9
            rhs375_ = iife185_(d_2340_v14_)
            rhs376_ = (d_2337_v11_) + (d_2337_v11_)
            lhs225_ = self
            lhs226_ = d_2328_v4_
            lhs225_.f7 = rhs373_
            lhs226_.f7 = rhs374_
            d_2340_v14_ = rhs375_
            d_2337_v11_ = rhs376_
            d_2343_v15_: C5
            nw410_ = C5()
            nw410_.ctor__(d_2328_v4_.f7)
            d_2343_v15_ = nw410_
            d_2344_v16_: _dafny.Set
            d_2344_v16_ = _dafny.Set({d_2324_v0_, d_2324_v0_})
            d_2345_v17_: D2
            d_2345_v17_ = D2_DC7(d_2344_v16_)
            d_2346_v18_: _dafny.Seq
            d_2346_v18_ = _dafny.SeqWithoutIsStrInference([d_2344_v16_])
            pat_let_tv50_ = d_2344_v16_
            pat_let_tv51_ = d_2344_v16_
            d_2347_v19_: _dafny.Array
            nw411_ = _dafny.Array(None, 19)
            nw411_[int(0)] = d_2345_v17_
            nw411_[int(1)] = d_2345_v17_
            nw411_[int(2)] = D2_DC7((d_2346_v18_)[default__.safeIndex(d_2324_v0_, len(d_2346_v18_))])
            nw411_[int(3)] = d_2345_v17_
            def iife189_(_pat_let69_0):
                def iife190_(d_2348_dt__update__tmp_h1_):
                    def iife191_(_pat_let70_0):
                        def iife192_(d_2349_dt__update_hcf13_h0_):
                            return D2_DC7(d_2349_dt__update_hcf13_h0_)
                        return iife192_(_pat_let70_0)
                    return iife191_(pat_let_tv50_)
                return iife190_(_pat_let69_0)
            nw411_[int(4)] = iife189_(D2_DC7(d_2344_v16_))
            def iife193_(_pat_let71_0):
                def iife194_(d_2350_dt__update__tmp_h2_):
                    def iife195_(_pat_let72_0):
                        def iife196_(d_2351_dt__update_hcf13_h1_):
                            return D2_DC7(d_2351_dt__update_hcf13_h1_)
                        return iife196_(_pat_let72_0)
                    return iife195_(pat_let_tv51_)
                return iife194_(_pat_let71_0)
            nw411_[int(5)] = iife193_(D2_DC7(d_2344_v16_))
            nw411_[int(6)] = d_2345_v17_
            nw411_[int(7)] = d_2345_v17_
            nw411_[int(8)] = d_2345_v17_
            nw411_[int(9)] = d_2345_v17_
            nw411_[int(10)] = d_2345_v17_
            nw411_[int(11)] = D2_DC7(d_2344_v16_)
            nw411_[int(12)] = d_2345_v17_
            nw411_[int(13)] = d_2345_v17_
            nw411_[int(14)] = d_2345_v17_
            nw411_[int(15)] = d_2345_v17_
            nw411_[int(16)] = D2_DC7(_dafny.Set({d_2324_v0_, d_2324_v0_, d_2324_v0_, d_2324_v0_, d_2324_v0_}))
            nw411_[int(17)] = d_2345_v17_
            nw411_[int(18)] = D2_DC7(d_2344_v16_)
            d_2347_v19_ = nw411_
            d_2352_v20_: C4
            nw412_ = C4()
            nw412_.ctor__(d_2347_v19_, (self).f9)
            d_2352_v20_ = nw412_
            d_2353_v21_: _dafny.Map
            d_2353_v21_ = _dafny.Map({(self).f9: d_2352_v20_})
            d_2354_v22_: _dafny.Map
            d_2354_v22_ = _dafny.Map({d_2327_v3_: d_2352_v20_})
            d_2355_v23_: _dafny.Seq
            d_2355_v23_ = _dafny.SeqWithoutIsStrInference([((d_2354_v22_)[d_2327_v3_] if (d_2327_v3_) in (d_2354_v22_) else d_2352_v20_), d_2352_v20_])
            d_2356_v24_: _dafny.Array
            nw413_ = _dafny.Array(None, 22)
            nw413_[int(0)] = d_2352_v20_
            nw413_[int(1)] = d_2352_v20_
            nw413_[int(2)] = d_2352_v20_
            nw413_[int(3)] = d_2352_v20_
            nw413_[int(4)] = d_2352_v20_
            nw413_[int(5)] = ((d_2353_v21_)[(self).f9] if ((self).f9) in (d_2353_v21_) else d_2352_v20_)
            nw413_[int(6)] = d_2352_v20_
            nw413_[int(7)] = d_2352_v20_
            nw413_[int(8)] = d_2352_v20_
            nw413_[int(9)] = d_2352_v20_
            nw413_[int(10)] = d_2352_v20_
            nw413_[int(11)] = d_2352_v20_
            nw413_[int(12)] = d_2352_v20_
            nw413_[int(13)] = d_2352_v20_
            nw413_[int(14)] = d_2352_v20_
            nw413_[int(15)] = d_2352_v20_
            nw413_[int(16)] = d_2352_v20_
            nw413_[int(17)] = d_2352_v20_
            nw413_[int(18)] = d_2352_v20_
            nw413_[int(19)] = d_2352_v20_
            nw413_[int(20)] = (d_2352_v20_ if (self).f9 else d_2352_v20_)
            nw413_[int(21)] = (d_2355_v23_)[default__.safeIndex(d_2324_v0_, len(d_2355_v23_))]
            d_2356_v24_ = nw413_
            index379_ = default__.safeIndex(229, (d_2356_v24_).length(0))
            (d_2356_v24_)[index379_] = (d_2355_v23_)[default__.safeIndex(d_2324_v0_, len(d_2355_v23_))]
            index380_ = default__.safeIndex(229, (d_2356_v24_).length(0))
            (d_2356_v24_)[index380_] = d_2352_v20_
        elif True:
            d_2357_v25_: _dafny.Array
            def lambda144_(d_2358_i1_):
                return D2_DC7(_dafny.Set({-717}))

            init81_ = lambda144_
            nw414_ = _dafny.Array(None, 6)
            for i0_81_ in range(nw414_.length(0)):
                nw414_[i0_81_] = init81_(i0_81_)
            d_2357_v25_ = nw414_
            d_2359_v26_: C4
            nw415_ = C4()
            nw415_.ctor__(d_2357_v25_, (not(True)) or (False))
            d_2359_v26_ = nw415_
            d_2360_v27_: _dafny.Array
            nw416_ = _dafny.Array(False, 16)
            d_2360_v27_ = nw416_
            d_2361_v28_: _dafny.Set
            d_2361_v28_ = _dafny.Set({d_2360_v27_})
            d_2362_v29_: _dafny.Map
            d_2362_v29_ = _dafny.Map({(self).f9: d_2360_v27_})
            d_2363_v30_: _dafny.Array
            nw417_ = _dafny.Array(None, 6)
            nw417_[int(0)] = d_2361_v28_
            nw417_[int(1)] = (d_2361_v28_) - (d_2361_v28_)
            nw417_[int(2)] = d_2361_v28_
            nw417_[int(3)] = d_2361_v28_
            nw417_[int(4)] = (d_2361_v28_).intersection(d_2361_v28_)
            nw417_[int(5)] = _dafny.Set({d_2360_v27_, d_2360_v27_, d_2360_v27_, ((d_2362_v29_)[not((self).f9)] if (not((self).f9)) in (d_2362_v29_) else d_2360_v27_)})
            d_2363_v30_ = nw417_
            index381_ = default__.safeIndex(304, (d_2363_v30_).length(0))
            (d_2363_v30_)[index381_] = d_2361_v28_
            index382_ = default__.safeIndex(304, (d_2363_v30_).length(0))
            (d_2363_v30_)[index382_] = d_2361_v28_
            d_2364_v31_: str
            d_2364_v31_ = _dafny.CodePoint('d')
            d_2365_v32_: int
            d_2365_v32_ = 43
            d_2366_v33_: _dafny.Seq
            d_2366_v33_ = _dafny.SeqWithoutIsStrInference([d_2365_v32_])
            d_2367_v34_: _dafny.MultiSet
            d_2367_v34_ = _dafny.MultiSet([self.f7, not(self.f7), (self).f9])
            d_2368_v35_: _dafny.Seq
            d_2368_v35_ = _dafny.SeqWithoutIsStrInference([d_2364_v31_])
            d_2369_v36_: D0
            d_2369_v36_ = D0_DC0(d_2365_v32_)
            d_2370_v37_: _dafny.Seq
            d_2370_v37_ = _dafny.SeqWithoutIsStrInference([d_2369_v36_, d_2369_v36_])
            d_2371_v38_: _dafny.MultiSet
            d_2371_v38_ = _dafny.MultiSet([623])
            d_2372_v39_: _dafny.Array
            nw418_ = _dafny.Array(None, 15)
            nw418_[int(0)] = d_2365_v32_
            nw418_[int(1)] = (d_2365_v32_) * (len(d_2366_v33_))
            nw418_[int(2)] = (0) - (((d_2367_v34_)[self.f7] if (self.f7) in (d_2367_v34_) else 985))
            nw418_[int(3)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caerf")))
            nw418_[int(4)] = (841) - (len(d_2368_v35_))
            nw418_[int(5)] = d_2365_v32_
            nw418_[int(6)] = default__.safeDivisionInt(d_2365_v32_, (d_2367_v34_).cardinality)
            nw418_[int(7)] = len(_dafny.Set({(self).fm12((self).fm8(d_2365_v32_, globalState), d_2365_v32_, d_2365_v32_, default__.fm1(d_2370_v37_, d_2365_v32_, d_2365_v32_, globalState), globalState)}))
            nw418_[int(8)] = d_2365_v32_
            nw418_[int(9)] = d_2365_v32_
            nw418_[int(10)] = (((d_2371_v38_).set(len(d_2368_v35_), default__.abs(387))) | ((_dafny.MultiSet([(d_2371_v38_).cardinality])).set(d_2365_v32_, default__.abs((self).fm12((self).f9, d_2365_v32_, (d_2371_v38_).cardinality, not((self).f9), globalState))))).cardinality
            nw418_[int(11)] = ((d_2367_v34_)[(self).f9] if ((self).f9) in (d_2367_v34_) else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2373_i2_ in range(default__.abs(-608))])))
            nw418_[int(12)] = d_2365_v32_
            nw418_[int(13)] = ((0) - (d_2365_v32_)) - (d_2365_v32_)
            nw418_[int(14)] = 971
            d_2372_v39_ = nw418_
            index383_ = default__.safeIndex(512, (d_2372_v39_).length(0))
            (d_2372_v39_)[index383_] = default__.safeDivisionInt(d_2365_v32_, d_2365_v32_)
            index384_ = default__.safeIndex(512, (d_2372_v39_).length(0))
            rhs377_ = default__.fm13((not(self.f7) if self.f7 else False), d_2365_v32_, self.f7, globalState)
            rhs378_ = self.f7
            rhs379_ = ((len(_dafny.SeqWithoutIsStrInference([self.f7]))) * (default__.fm0(globalState))) + (-956)
            rhs380_ = (not((self).f9)) and ((self).f9)
            rhs381_ = (self).f9
            lhs227_ = self
            lhs228_ = d_2372_v39_
            lhs229_ = default__.safeIndex(512, (d_2372_v39_).length(0))
            lhs230_ = self
            lhs231_ = self
            d_2364_v31_ = rhs377_
            lhs227_.f7 = rhs378_
            lhs228_[lhs229_] = rhs379_
            lhs230_.f7 = rhs380_
            lhs231_.f7 = rhs381_
            if self.f7:
                d_2374_v40_: _dafny.Seq
                d_2374_v40_ = _dafny.SeqWithoutIsStrInference([self.f7])
                d_2375_v41_: _dafny.Set
                d_2375_v41_ = _dafny.Set({self.f7, (d_2374_v40_)[default__.safeIndex(903, len(d_2374_v40_))], (self).f9})
                d_2376_v42_: _dafny.Map
                d_2376_v42_ = _dafny.Map({self.f7: d_2375_v41_})
                index385_ = default__.safeIndex(512, (d_2372_v39_).length(0))
                (d_2372_v39_)[index385_] = (d_2365_v32_) - ((0) - (len(d_2376_v42_)))
                d_2377_v43_: _dafny.Map
                d_2377_v43_ = _dafny.Map({(self).f9: (d_2372_v39_)[default__.safeIndex(512, (d_2372_v39_).length(0))]})
                d_2378_v44_: C3
                nw419_ = C3()
                nw419_.ctor__(len((d_2377_v43_) | (default__.fm43(_dafny.CodePoint('k'), (self).f9, self.f7, globalState))), self.f7)
                d_2378_v44_ = nw419_
                index386_ = default__.safeIndex(512, (d_2372_v39_).length(0))
                (d_2372_v39_)[index386_] = (0) - ((0) - ((0) - ((0) - (d_2365_v32_))))
                (self).f7 = (self).fm8((d_2378_v44_).f14, globalState)
                (self).f7 = ((d_2365_v32_) + (918)) > (d_2365_v32_)
            elif True:
                d_2368_v35_ = d_2368_v35_
                d_2379_v45_: C5
                nw420_ = C5()
                nw420_.ctor__((self).f9)
                d_2379_v45_ = nw420_
                (self).f7 = (False if not (not((d_2379_v45_).f10)) or (self.f7) else (d_2379_v45_).f10)
                d_2365_v32_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpuhsuu"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkajlx")) if (d_2379_v45_).f10 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aj")))))
                index387_ = default__.safeIndex(512, (d_2372_v39_).length(0))
                (d_2372_v39_)[index387_] = (d_2372_v39_)[default__.safeIndex(512, (d_2372_v39_).length(0))]
            default__.m0((d_2372_v39_)[default__.safeIndex(512, (d_2372_v39_).length(0))], (d_2372_v39_)[default__.safeIndex(512, (d_2372_v39_).length(0))], globalState)
        d_2380_v46_: int
        d_2380_v46_ = 91
        d_2381_i3_: int
        d_2381_i3_ = 0
        with _dafny.label("16"):
            while (d_2380_v46_) != (d_2380_v46_):
                with _dafny.c_label("16"):
                    if (d_2381_i3_) >= (100):
                        raise _dafny.Break("16")
                    d_2381_i3_ = (d_2381_i3_) + (1)
                    d_2382_v47_: str
                    d_2382_v47_ = _dafny.CodePoint('s')
                    d_2383_v48_: _dafny.Seq
                    d_2383_v48_ = _dafny.SeqWithoutIsStrInference([d_2382_v47_, d_2382_v47_])
                    d_2384_v49_: _dafny.Seq
                    d_2384_v49_ = _dafny.SeqWithoutIsStrInference([d_2380_v46_, len(_dafny.SeqWithoutIsStrInference([d_2380_v46_])), d_2380_v46_, d_2380_v46_, -967])
                    d_2385_v50_: _dafny.Seq
                    d_2385_v50_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                    d_2386_v51_: _dafny.Array
                    nw421_ = _dafny.Array(None, 10)
                    nw421_[int(0)] = self.f7
                    nw421_[int(1)] = not(self.f7)
                    nw421_[int(2)] = not (self.f7) or ((self).f9)
                    nw421_[int(3)] = (self).f9
                    nw421_[int(4)] = not(((d_2383_v48_)[default__.safeIndex(len(d_2384_v49_), len(d_2383_v48_))]) not in (_dafny.SeqWithoutIsStrInference([d_2382_v47_ for d_2387_i4_ in range(default__.abs(500))])))
                    nw421_[int(5)] = self.f7
                    nw421_[int(6)] = (d_2382_v47_) in (d_2383_v48_)
                    nw421_[int(7)] = self.f7
                    nw421_[int(8)] = not((self).f9)
                    nw421_[int(9)] = (len(d_2385_v50_)) == ((0) - (d_2380_v46_))
                    d_2386_v51_ = nw421_
                    index388_ = default__.safeIndex(541, (d_2386_v51_).length(0))
                    (d_2386_v51_)[index388_] = self.f7
                    index389_ = default__.safeIndex(541, (d_2386_v51_).length(0))
                    (d_2386_v51_)[index389_] = self.f7
                    d_2380_v46_ = 625
                    d_2380_v46_ = default__.safeModuloInt((314 if (self).f9 else d_2380_v46_), -620)
                    d_2382_v47_ = d_2382_v47_
                    pass
            pass
        d_2388_v52_: _dafny.Array
        def lambda145_(d_2389_i5_):
            return self.f7

        init82_ = lambda145_
        nw422_ = _dafny.Array(None, 28)
        for i0_82_ in range(nw422_.length(0)):
            nw422_[i0_82_] = init82_(i0_82_)
        d_2388_v52_ = nw422_
        d_2390_v53_: _dafny.Map
        d_2390_v53_ = _dafny.Map({d_2388_v52_: d_2380_v46_})
        d_2390_v53_ = (d_2390_v53_).set(d_2388_v52_, d_2380_v46_)
        d_2391_v54_: _dafny.Seq
        d_2391_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnenprj"))
        d_2392_v55_: _dafny.Seq
        d_2392_v55_ = _dafny.SeqWithoutIsStrInference([d_2380_v46_])
        d_2393_v56_: D9
        d_2393_v56_ = D9_DC26(d_2392_v55_)
        pat_let_tv52_ = d_2391_v54_
        pat_let_tv53_ = d_2391_v54_
        pat_let_tv54_ = d_2391_v54_
        pat_let_tv55_ = d_2391_v54_
        def lambda146_(source17_):
            if source17_.is_DC24:
                return pat_let_tv52_
            elif source17_.is_DC25:
                return pat_let_tv53_
            elif source17_.is_DC26:
                d_2394___mcc_h0_ = source17_.cf45
                d_2395_cf45_ = d_2394___mcc_h0_
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2396_i6_ in range(default__.abs(34))])
            elif True:
                d_2397___mcc_h1_ = source17_.cf44
                d_2398_cf44_ = d_2397___mcc_h1_
                return (pat_let_tv54_) + (pat_let_tv55_)

        d_2391_v54_ = lambda146_(d_2393_v56_)
        d_2399_v57_: C3
        nw423_ = C3()
        nw423_.ctor__((d_2380_v46_ if self.f7 else d_2380_v46_), not((self).f9))
        d_2399_v57_ = nw423_
        hi13_ = len(_dafny.SeqWithoutIsStrInference([d_2380_v46_]))
        for d_2400_i7_ in range((d_2380_v46_) * (d_2380_v46_), hi13_):
            if ((d_2399_v57_).f14) >= ((d_2399_v57_).f14):
                d_2380_v46_ = d_2400_i7_
                d_2380_v46_ = d_2400_i7_
                d_2401_v58_: D0
                d_2401_v58_ = D0_DC0(d_2380_v46_)
                d_2402_v59_: _dafny.Seq
                d_2402_v59_ = _dafny.SeqWithoutIsStrInference([d_2401_v58_, d_2401_v58_])
                d_2403_v60_: _dafny.Map
                d_2403_v60_ = _dafny.Map({d_2391_v54_: d_2400_i7_})
                (self).f7 = ((self).f9 if default__.fm1(d_2402_v59_, len(d_2403_v60_), d_2380_v46_, globalState) else (self).f9)
                d_2404_v61_: _dafny.Array
                nw424_ = _dafny.Array(int(0), 8)
                d_2404_v61_ = nw424_
                index390_ = default__.safeIndex(936, (d_2404_v61_).length(0))
                (d_2404_v61_)[index390_] = len((_dafny.SeqWithoutIsStrInference([(self).f9])) + (_dafny.SeqWithoutIsStrInference([(self).f9])))
                index391_ = default__.safeIndex(936, (d_2404_v61_).length(0))
                (d_2404_v61_)[index391_] = len(d_2391_v54_)
                d_2380_v46_ = ((d_2399_v57_).f14) * ((d_2399_v57_).f14)
            elif True:
                d_2405_v62_: _dafny.Seq
                d_2405_v62_ = _dafny.SeqWithoutIsStrInference([self.f7, self.f7])
                d_2406_v63_: _dafny.Map
                d_2406_v63_ = _dafny.Map({729: len(d_2405_v62_)})
                d_2407_v64_: _dafny.Map
                d_2407_v64_ = _dafny.Map({d_2406_v63_: (d_2400_i7_) - ((d_2399_v57_).f14)})
                d_2407_v64_ = (_dafny.Map({d_2406_v63_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdo")))})).set(_dafny.Map({d_2380_v46_: (d_2399_v57_).f14}), (d_2392_v55_)[default__.safeIndex((d_2399_v57_).f14, len(d_2392_v55_))])
                (self).f7 = (self).fm8((d_2399_v57_).f14, globalState)
                (self).f7 = self.f7
                d_2408_v65_: _dafny.Array
                def lambda147_(d_2409_v57_):
                    def lambda148_(d_2410_i8_):
                        return (d_2410_i8_) * ((d_2409_v57_).f14)

                    return lambda148_

                init83_ = lambda147_(d_2399_v57_)
                nw425_ = _dafny.Array(None, 9)
                for i0_83_ in range(nw425_.length(0)):
                    nw425_[i0_83_] = init83_(i0_83_)
                d_2408_v65_ = nw425_
                index392_ = default__.safeIndex(365, (d_2408_v65_).length(0))
                (d_2408_v65_)[index392_] = d_2400_i7_
                index393_ = default__.safeIndex(365, (d_2408_v65_).length(0))
                (d_2408_v65_)[index393_] = 857
                d_2411_v66_: str
                d_2411_v66_ = _dafny.CodePoint('p')
                d_2412_v67_: _dafny.Set
                d_2412_v67_ = _dafny.Set({d_2411_v66_})
                d_2413_v68_: _dafny.Set
                d_2413_v68_ = _dafny.Set({d_2391_v54_})
                d_2414_v69_: _dafny.Array
                nw426_ = _dafny.Array(None, 20)
                nw426_[int(0)] = (_dafny.SeqWithoutIsStrInference([self.f7, (self).f9, (self).f9, (self).f9])) + (d_2405_v62_)
                nw426_[int(1)] = default__.fm50(d_2412_v67_, d_2413_v68_, (_dafny.MultiSet([D14_DC39(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")), d_2408_v65_)])).cardinality, globalState)
                nw426_[int(2)] = d_2405_v62_
                nw426_[int(3)] = d_2405_v62_
                nw426_[int(4)] = d_2405_v62_
                nw426_[int(5)] = d_2405_v62_
                nw426_[int(6)] = d_2405_v62_
                nw426_[int(7)] = (_dafny.SeqWithoutIsStrInference([True])) + (d_2405_v62_)
                nw426_[int(8)] = (d_2405_v62_) + (d_2405_v62_)
                nw426_[int(9)] = d_2405_v62_
                nw426_[int(10)] = d_2405_v62_
                nw426_[int(11)] = d_2405_v62_
                nw426_[int(12)] = (d_2405_v62_) + (_dafny.SeqWithoutIsStrInference([self.f7, (self).f9]))
                nw426_[int(13)] = (d_2405_v62_) + (_dafny.SeqWithoutIsStrInference([(self).f9]))
                nw426_[int(14)] = default__.fm50(d_2412_v67_, d_2413_v68_, len(_dafny.Set({self.f7})), globalState)
                nw426_[int(15)] = d_2405_v62_
                nw426_[int(16)] = d_2405_v62_
                nw426_[int(17)] = d_2405_v62_
                nw426_[int(18)] = _dafny.SeqWithoutIsStrInference([self.f7])
                nw426_[int(19)] = (d_2405_v62_) + (d_2405_v62_)
                d_2414_v69_ = nw426_
                index394_ = default__.safeIndex(447, (d_2414_v69_).length(0))
                (d_2414_v69_)[index394_] = d_2405_v62_
                index395_ = default__.safeIndex(447, (d_2414_v69_).length(0))
                (d_2414_v69_)[index395_] = _dafny.SeqWithoutIsStrInference([self.f7, (self).f9, (self).f9])
            d_2415_v70_: _dafny.Array
            def lambda149_(d_2416_v57_):
                def lambda150_(d_2417_i9_):
                    return _dafny.Set({(d_2416_v57_).f14})

                return lambda150_

            init84_ = lambda149_(d_2399_v57_)
            nw427_ = _dafny.Array(None, 25)
            for i0_84_ in range(nw427_.length(0)):
                nw427_[i0_84_] = init84_(i0_84_)
            d_2415_v70_ = nw427_
            index396_ = default__.safeIndex(640, (d_2415_v70_).length(0))
            (d_2415_v70_)[index396_] = _dafny.Set({(_dafny.MultiSet([(d_2399_v57_).f14, (d_2399_v57_).f14])).cardinality, 484})
            d_2418_v71_: _dafny.Set
            d_2418_v71_ = _dafny.Set({d_2400_i7_})
            index397_ = default__.safeIndex(640, (d_2415_v70_).length(0))
            (d_2415_v70_)[index397_] = d_2418_v71_
            d_2419_v72_: _dafny.Seq
            d_2419_v72_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_2420_v73_: _dafny.Map
            d_2420_v73_ = _dafny.Map({(d_2399_v57_).f14: d_2419_v72_})
            d_2421_v74_: _dafny.Seq
            d_2421_v74_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f7]), (((d_2420_v73_)[(d_2399_v57_).f14] if ((d_2399_v57_).f14) in (d_2420_v73_) else d_2419_v72_)) + (d_2419_v72_), d_2419_v72_])
            d_2421_v74_ = d_2421_v74_
            d_2422_v75_: str
            d_2422_v75_ = _dafny.CodePoint('b')
            (self).f7 = (d_2422_v75_) not in ((d_2391_v54_) + (d_2391_v54_))

    @property
    def f9(self):
        return self._f9

class C7(T0, T1):
    def  __init__(self):
        self._f7: bool = False
        self._f8: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def ctor__(self, f8, f7):
        (self)._f8 = f8
        (self).f7 = f7

    def fm8(self, p0, globalState):
        return self.f7

    def fm9(self, p0, p1, p2, globalState):
        def iife197_():
            def iife198_():
                coll52_ = _dafny.Map()
                compr_52_: int
                for compr_52_ in _dafny.IntegerRange(573, 822):
                    d_2424_v1_: int = compr_52_
                    if ((573) <= (d_2424_v1_)) and ((d_2424_v1_) < (822)):
                        coll52_[(d_2424_v1_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugnal"))))] = self.f7
                return _dafny.Map(coll52_)
            coll51_ = _dafny.Map()
            compr_51_: int
            for compr_51_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([628, (0) - (len(_dafny.SeqWithoutIsStrInference([self.f7])))])): len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, False, self.f7])).cardinality]))})).keys.Elements:
                d_2423_v0_: int = compr_51_
                if (d_2423_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([628, (0) - (len(_dafny.SeqWithoutIsStrInference([self.f7])))])): len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, False, self.f7])).cardinality]))})):
                    coll51_[(d_2423_v0_) - (len(iife198_()
                    ))] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f7, self.f7])), 734]))
            return _dafny.Map(coll51_)
        return ((0) - (default__.safeDivisionInt(-831, len(iife197_()
        )))) + (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyvc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))))

    def fm10(self, p0, globalState):
        return ((_dafny.MultiSet([485, -242])).cardinality) * ((0) - (len(_dafny.Map({self.f7: not(self.f7)}))))

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Set = _dafny.Set({})
        hi14_ = (default__.fm0(globalState)) - ((self).fm10(False, globalState))
        for d_2425_i0_ in range(p1, hi14_):
            d_2426_v0_: _dafny.Map
            d_2426_v0_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1 for d_2427_i1_ in range(default__.abs(100))]): self.f7})
            d_2428_v1_: _dafny.Seq
            d_2428_v1_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({self.f7: p0}))])
            d_2429_v2_: _dafny.Map
            d_2429_v2_ = _dafny.Map({d_2425_i0_: default__.fm11(((d_2426_v0_)[d_2428_v1_] if (d_2428_v1_) in (d_2426_v0_) else False), self.f7, d_2425_i0_, self.f7, globalState)})
            d_2429_v2_ = (d_2429_v2_).set(p1, _dafny.Map({p2: not(True)}))
            if self.f7:
                d_2430_v3_: int
                d_2430_v3_ = 565
                d_2431_v4_: _dafny.Map
                d_2431_v4_ = _dafny.Map({p0: d_2430_v3_})
                d_2432_v5_: _dafny.Map
                d_2432_v5_ = _dafny.Map({_dafny.Map({True: d_2430_v3_}): p2})
                d_2430_v3_ = len(((_dafny.Map({d_2431_v4_: p0})) | (d_2432_v5_)).set(d_2431_v4_, False))
                d_2430_v3_ = d_2430_v3_
                d_2430_v3_ = default__.safeDivisionInt(p1, (0) - ((p1) + (p1)))
                (self).f7 = self.f7
                d_2430_v3_ = p1
            elif True:
                (self).f7 = not(self.f7)
                d_2433_v6_: _dafny.MultiSet
                d_2433_v6_ = _dafny.MultiSet([d_2425_i0_])
                d_2434_v7_: _dafny.MultiSet
                d_2434_v7_ = _dafny.MultiSet([False, p0, True, self.f7])
                d_2435_v8_: _dafny.Seq
                d_2435_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmrn"))
                d_2436_v9_: _dafny.Array
                nw428_ = _dafny.Array(None, 10)
                nw428_[int(0)] = p1
                nw428_[int(1)] = (0) - (len(_dafny.Set({p2, self.f7, self.f7, self.f7, p2})))
                nw428_[int(2)] = default__.safeDivisionInt(p1, (0) - (d_2425_i0_))
                nw428_[int(3)] = ((d_2433_v6_ if self.f7 else d_2433_v6_)).cardinality
                nw428_[int(4)] = 888
                nw428_[int(5)] = (d_2425_i0_ if p2 else d_2425_i0_)
                nw428_[int(6)] = (d_2434_v7_).cardinality
                nw428_[int(7)] = (29) + (d_2425_i0_)
                nw428_[int(8)] = d_2425_i0_
                nw428_[int(9)] = len(d_2435_v8_)
                d_2436_v9_ = nw428_
                d_2437_v10_: _dafny.Array
                nw429_ = _dafny.Array(False, 27)
                d_2437_v10_ = nw429_
                index398_ = default__.safeIndex(526, (d_2437_v10_).length(0))
                (d_2437_v10_)[index398_] = self.f7
                index399_ = default__.safeIndex(526, (d_2437_v10_).length(0))
                rhs382_ = d_2436_v9_
                rhs383_ = p0
                lhs232_ = d_2437_v10_
                lhs233_ = default__.safeIndex(526, (d_2437_v10_).length(0))
                d_2436_v9_ = rhs382_
                lhs232_[lhs233_] = rhs383_
                d_2438_v11_: str
                d_2438_v11_ = _dafny.CodePoint('g')
                d_2438_v11_ = d_2438_v11_
                (self).m5(globalState)
                d_2439_v12_: int
                d_2439_v12_ = 455
                d_2439_v12_ = d_2439_v12_
            (self).f7 = p2
            d_2440_v13_: _dafny.Array
            nw430_ = _dafny.Array(D1.default()(), 19)
            d_2440_v13_ = nw430_
            d_2441_v14_: D1
            d_2441_v14_ = D1_DC6(D1_DC5(p0, _dafny.SeqWithoutIsStrInference([default__.fm0(globalState)]), d_2425_i0_))
            d_2442_v15_: D1
            d_2442_v15_ = D1_DC6(d_2441_v14_)
            index400_ = default__.safeIndex(895, (d_2440_v13_).length(0))
            (d_2440_v13_)[index400_] = d_2442_v15_
            index401_ = default__.safeIndex(895, (d_2440_v13_).length(0))
            (d_2440_v13_)[index401_] = d_2442_v15_
        hi15_ = p1
        for d_2443_i2_ in range(p1, hi15_):
            d_2444_v16_: _dafny.Seq
            d_2444_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxlh"))
            d_2444_v16_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsxh"))) + (d_2444_v16_)
            d_2445_v17_: C5
            nw431_ = C5()
            nw431_.ctor__((d_2443_i2_) >= (p1))
            d_2445_v17_ = nw431_
            if (d_2445_v17_).f10:
                d_2446_v18_: str
                d_2446_v18_ = _dafny.CodePoint('y')
                d_2447_v19_: _dafny.Seq
                d_2447_v19_ = _dafny.SeqWithoutIsStrInference([self.f7, p2])
                d_2448_v20_: D13
                d_2448_v20_ = D13_DC36(d_2446_v18_, d_2447_v19_, self.f7)
                d_2449_v21_: T1
                nw432_ = C6()
                nw432_.ctor__((d_2448_v20_).cf60, ((d_2445_v17_).f10 if p0 else p2))
                d_2449_v21_ = nw432_
                d_2449_v21_ = d_2449_v21_
                d_2450_v22_: _dafny.Set
                d_2450_v22_ = _dafny.Set({default__.fm21(p0, globalState), _dafny.SeqWithoutIsStrInference([p1 for d_2451_i3_ in range(default__.abs(562))])})
                d_2450_v22_ = d_2450_v22_
                d_2444_v16_ = d_2444_v16_
                d_2452_v23_: _dafny.Array
                nw433_ = _dafny.Array(int(0), 15)
                d_2452_v23_ = nw433_
                index402_ = default__.safeIndex(356, (d_2452_v23_).length(0))
                (d_2452_v23_)[index402_] = 450
                d_2453_v24_: _dafny.Array
                nw434_ = _dafny.Array(_dafny.Map({}), 10)
                d_2453_v24_ = nw434_
                d_2454_v25_: _dafny.Map
                d_2454_v25_ = _dafny.Map({d_2446_v18_: d_2453_v24_})
                d_2455_v26_: _dafny.Array
                d_2455_v26_ = d_2453_v24_
                d_2456_v27_: _dafny.Seq
                d_2456_v27_ = _dafny.SeqWithoutIsStrInference([d_2454_v25_, d_2454_v25_, d_2454_v25_, _dafny.Map({d_2446_v18_: d_2455_v26_})])
                index403_ = default__.safeIndex(356, (d_2452_v23_).length(0))
                (d_2452_v23_)[index403_] = len((((d_2456_v27_)[default__.safeIndex(d_2443_i2_, len(d_2456_v27_))]) | ((d_2454_v25_).set(d_2446_v18_, d_2453_v24_))) | (d_2454_v25_))
                index404_ = default__.safeIndex(356, (d_2452_v23_).length(0))
                (d_2452_v23_)[index404_] = p1
            elif True:
                d_2457_v28_: _dafny.Array
                def lambda151_(d_2458_i4_):
                    return default__.safeDivisionInt(d_2458_i4_, 93)

                init85_ = lambda151_
                nw435_ = _dafny.Array(None, 17)
                for i0_85_ in range(nw435_.length(0)):
                    nw435_[i0_85_] = init85_(i0_85_)
                d_2457_v28_ = nw435_
                index405_ = default__.safeIndex(558, (d_2457_v28_).length(0))
                (d_2457_v28_)[index405_] = d_2443_i2_
                index406_ = default__.safeIndex(558, (d_2457_v28_).length(0))
                (d_2457_v28_)[index406_] = default__.safeDivisionInt(((0) - (p1)) * (d_2443_i2_), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxawko")))) - (d_2443_i2_))
                d_2459_v29_: _dafny.Array
                nw436_ = _dafny.Array(None, 9)
                nw436_[int(0)] = self.f7
                nw436_[int(1)] = p2
                nw436_[int(2)] = (d_2445_v17_).f10
                nw436_[int(3)] = p0
                nw436_[int(4)] = (d_2445_v17_).f10
                nw436_[int(5)] = (d_2445_v17_).f10
                nw436_[int(6)] = ((d_2445_v17_).f10) == (p0)
                nw436_[int(7)] = (d_2444_v16_) != (d_2444_v16_)
                nw436_[int(8)] = p2
                d_2459_v29_ = nw436_
                d_2459_v29_ = (d_2459_v29_ if p0 else d_2459_v29_)
                (self).f7 = self.f7
                (d_2445_v17_).m7(globalState)
                d_2460_v30_: str
                d_2460_v30_ = _dafny.CodePoint('a')
                d_2460_v30_ = d_2460_v30_
            d_2461_v31_: int
            d_2461_v31_ = 753
            d_2461_v31_ = p1
        d_2462_v32_: _dafny.Seq
        d_2462_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbgnqbndj"))
        d_2462_v32_ = ((d_2462_v32_) + (d_2462_v32_) if p0 else (d_2462_v32_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngt"))))
        d_2463_v33_: D0
        d_2463_v33_ = D0_DC0(p1)
        d_2464_v34_: _dafny.Seq
        d_2464_v34_ = _dafny.SeqWithoutIsStrInference([d_2463_v33_])
        d_2465_v35_: str
        d_2465_v35_ = _dafny.CodePoint('e')
        d_2466_v36_: _dafny.Seq
        d_2466_v36_ = _dafny.SeqWithoutIsStrInference([d_2462_v32_, (_dafny.SeqWithoutIsStrInference([(d_2462_v32_)[default__.safeIndex(p1, len(d_2462_v32_))] for d_2467_i5_ in range(default__.abs(870))])).set(default__.safeIndex(len(_dafny.Map({default__.fm1(d_2464_v34_, p1, p1, globalState): self.f7})), len(_dafny.SeqWithoutIsStrInference([(d_2462_v32_)[default__.safeIndex(p1, len(d_2462_v32_))] for d_2468_i5_ in range(default__.abs(870))]))), d_2465_v35_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojoyengg")), (d_2462_v32_) + (d_2462_v32_), (d_2462_v32_) + (_dafny.SeqWithoutIsStrInference([d_2465_v35_ for d_2469_i6_ in range(default__.abs(588))]))])
        d_2462_v32_ = (d_2466_v36_)[default__.safeIndex(p1, len(d_2466_v36_))]
        d_2470_i7_: int
        d_2470_i7_ = 0
        with _dafny.label("17"):
            while (p1) >= (p1):
                with _dafny.c_label("17"):
                    if (d_2470_i7_) >= (100):
                        raise _dafny.Break("17")
                    d_2470_i7_ = (d_2470_i7_) + (1)
                    d_2471_v37_: _dafny.Array
                    def lambda152_(d_2472_i8_):
                        return False

                    init86_ = lambda152_
                    nw437_ = _dafny.Array(None, 9)
                    for i0_86_ in range(nw437_.length(0)):
                        nw437_[i0_86_] = init86_(i0_86_)
                    d_2471_v37_ = nw437_
                    d_2473_v38_: T1
                    nw438_ = C6()
                    nw438_.ctor__(self.f7, p0)
                    d_2473_v38_ = nw438_
                    d_2474_v39_: _dafny.Map
                    d_2474_v39_ = _dafny.Map({d_2473_v38_: p2})
                    index407_ = default__.safeIndex(315, (d_2471_v37_).length(0))
                    (d_2471_v37_)[index407_] = (d_2473_v38_) not in (d_2474_v39_)
                    d_2475_v40_: _dafny.Array
                    nw439_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
                    d_2475_v40_ = nw439_
                    index408_ = default__.safeIndex(86, (d_2475_v40_).length(0))
                    (d_2475_v40_)[index408_] = d_2462_v32_
                    index409_ = default__.safeIndex(315, (d_2471_v37_).length(0))
                    index410_ = default__.safeIndex(86, (d_2475_v40_).length(0))
                    rhs384_ = (d_2462_v32_) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "duoejddgs"))) + (d_2462_v32_))
                    rhs385_ = (d_2462_v32_) + ((d_2462_v32_) + (d_2462_v32_))
                    lhs234_ = d_2471_v37_
                    lhs235_ = default__.safeIndex(315, (d_2471_v37_).length(0))
                    lhs236_ = d_2475_v40_
                    lhs237_ = default__.safeIndex(86, (d_2475_v40_).length(0))
                    lhs234_[lhs235_] = rhs384_
                    lhs236_[lhs237_] = rhs385_
                    (d_2473_v38_).f7 = p2
                    if d_2473_v38_.f7:
                        d_2476_v41_: int
                        d_2476_v41_ = 479
                        d_2476_v41_ = (0) - ((p1) * (default__.safeDivisionInt(45, d_2476_v41_)))
                        d_2477_v42_: _dafny.Array
                        nw440_ = _dafny.Array(None, 6)
                        nw440_[int(0)] = d_2476_v41_
                        nw440_[int(1)] = 149
                        nw440_[int(2)] = p1
                        nw440_[int(3)] = d_2476_v41_
                        nw440_[int(4)] = p1
                        nw440_[int(5)] = len(d_2462_v32_)
                        d_2477_v42_ = nw440_
                        d_2478_v43_: D14
                        d_2478_v43_ = D14_DC39(d_2462_v32_, d_2477_v42_)
                        d_2479_v44_: _dafny.Map
                        d_2479_v44_ = _dafny.Map({d_2465_v35_: (d_2475_v40_)[default__.safeIndex(86, (d_2475_v40_).length(0))]})
                        d_2480_v45_: _dafny.Seq
                        d_2480_v45_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_2465_v35_])) <= (((d_2479_v44_)[d_2465_v35_] if (d_2465_v35_) in (d_2479_v44_) else (_dafny.SeqWithoutIsStrInference([d_2465_v35_])).set(default__.safeIndex(len((d_2462_v32_).set(default__.safeIndex(770, len(d_2462_v32_)), _dafny.CodePoint('v'))), len(_dafny.SeqWithoutIsStrInference([d_2465_v35_]))), d_2465_v35_)))])
                        index411_ = default__.safeIndex(315, (d_2471_v37_).length(0))
                        rhs386_ = ((d_2478_v43_).cf63) < ((d_2462_v32_ if d_2473_v38_.f7 else d_2462_v32_))
                        rhs387_ = (d_2480_v45_)[default__.safeIndex(d_2476_v41_, len(d_2480_v45_))]
                        lhs238_ = d_2473_v38_
                        lhs239_ = d_2471_v37_
                        lhs240_ = default__.safeIndex(315, (d_2471_v37_).length(0))
                        lhs238_.f7 = rhs386_
                        lhs239_[lhs240_] = rhs387_
                        d_2476_v41_ = (0) - (((0) - (p1)) - (d_2476_v41_))
                        d_2481_v46_: _dafny.Set
                        d_2481_v46_ = _dafny.Set({d_2473_v38_.f7, True})
                        d_2482_v47_: _dafny.Seq
                        d_2482_v47_ = _dafny.SeqWithoutIsStrInference([d_2476_v41_, 138])
                        d_2483_v48_: D1
                        d_2483_v48_ = D1_DC5(d_2473_v38_.f7, d_2482_v47_, 616)
                        r1 = (d_2481_v46_) - ((default__.fm27(d_2473_v38_.f7, d_2465_v35_, (d_2471_v37_)[default__.safeIndex(315, (d_2471_v37_).length(0))], not((d_2483_v48_).cf9), globalState) if False else d_2481_v46_))
                        d_2476_v41_ = default__.safeDivisionInt(p1, p1)
                    elif True:
                        d_2484_v49_: int
                        d_2484_v49_ = 329
                        d_2484_v49_ = p1
                        d_2485_v50_: _dafny.Array
                        nw441_ = _dafny.Array(int(0), 26)
                        d_2485_v50_ = nw441_
                        index412_ = default__.safeIndex(412, (d_2485_v50_).length(0))
                        (d_2485_v50_)[index412_] = 35
                        d_2486_v51_: _dafny.Map
                        d_2486_v51_ = _dafny.Map({self.f7: d_2484_v49_})
                        index413_ = default__.safeIndex(412, (d_2485_v50_).length(0))
                        rhs388_ = not(p0)
                        rhs389_ = ((d_2486_v51_)[d_2473_v38_.f7] if (d_2473_v38_.f7) in (d_2486_v51_) else d_2484_v49_)
                        rhs390_ = p1
                        rhs391_ = (d_2473_v38_).fm9(not((d_2471_v37_)[default__.safeIndex(315, (d_2471_v37_).length(0))]), p1, d_2465_v35_, globalState)
                        rhs392_ = d_2484_v49_
                        lhs241_ = self
                        lhs242_ = d_2485_v50_
                        lhs243_ = default__.safeIndex(412, (d_2485_v50_).length(0))
                        lhs241_.f7 = rhs388_
                        d_2484_v49_ = rhs389_
                        lhs242_[lhs243_] = rhs390_
                        d_2484_v49_ = rhs391_
                        d_2484_v49_ = rhs392_
                        (d_2473_v38_).f7 = (d_2471_v37_)[default__.safeIndex(315, (d_2471_v37_).length(0))]
                        pat_let_tv56_ = d_2484_v49_
                        pat_let_tv57_ = d_2484_v49_
                        pat_let_tv58_ = p1
                        index414_ = default__.safeIndex(315, (d_2471_v37_).length(0))
                        def iife199_(_pat_let73_0):
                            def iife200_(d_2487_dt__update__tmp_h1_):
                                def iife201_(_pat_let74_0):
                                    def iife202_(d_2488_dt__update_hcf22_h0_):
                                        def iife203_(_pat_let75_0):
                                            def iife204_(d_2489_dt__update_hcf23_h0_):
                                                def iife205_(_pat_let76_0):
                                                    def iife206_(d_2490_dt__update_hcf24_h0_):
                                                        return D3_DC11(d_2488_dt__update_hcf22_h0_, d_2489_dt__update_hcf23_h0_, d_2490_dt__update_hcf24_h0_, (d_2487_dt__update__tmp_h1_).cf25)
                                                    return iife206_(_pat_let76_0)
                                                return iife205_(pat_let_tv58_)
                                            return iife204_(_pat_let75_0)
                                        return iife203_(pat_let_tv57_)
                                    return iife202_(_pat_let74_0)
                                return iife201_(pat_let_tv56_)
                            return iife200_(_pat_let73_0)
                        (d_2471_v37_)[index414_] = (iife199_(D3_DC11(default__.fm0(globalState), (d_2485_v50_)[default__.safeIndex(412, (d_2485_v50_).length(0))], p1, d_2473_v38_.f7))).cf25
                        index415_ = default__.safeIndex(412, (d_2485_v50_).length(0))
                        (d_2485_v50_)[index415_] = p1
                    if p0:
                        d_2491_v52_: int
                        d_2491_v52_ = 150
                        d_2491_v52_ = d_2491_v52_
                        d_2492_v53_: _dafny.Seq
                        d_2492_v53_ = _dafny.SeqWithoutIsStrInference([self.f7])
                        d_2493_v54_: _dafny.Set
                        d_2493_v54_ = _dafny.Set({len(d_2492_v53_)})
                        d_2494_v55_: _dafny.Map
                        d_2494_v55_ = _dafny.Map({d_2463_v33_: len(d_2493_v54_)})
                        d_2495_v56_: D2
                        d_2495_v56_ = D2_DC9(d_2491_v52_, d_2491_v52_, (d_2475_v40_)[default__.safeIndex(86, (d_2475_v40_).length(0))], p0, d_2494_v55_)
                        d_2496_v57_: _dafny.Set
                        d_2496_v57_ = _dafny.Set({len((d_2495_v56_).cf18)})
                        d_2497_v58_: _dafny.MultiSet
                        d_2497_v58_ = _dafny.MultiSet([p1, (d_2491_v52_) - (len(d_2493_v54_))])
                        d_2497_v58_ = d_2497_v58_
                        d_2498_v59_: D1
                        d_2498_v59_ = D1_DC3(_dafny.Set({d_2473_v38_.f7, p0, d_2473_v38_.f7}))
                        d_2499_v60_: _dafny.Seq
                        d_2499_v60_ = _dafny.SeqWithoutIsStrInference([d_2491_v52_, p1, 872, len((d_2475_v40_)[default__.safeIndex(86, (d_2475_v40_).length(0))]), len(d_2492_v53_)])
                        d_2500_v61_: _dafny.Map
                        d_2500_v61_ = _dafny.Map({d_2498_v59_: len(d_2499_v60_)})
                        d_2500_v61_ = d_2500_v61_
                        d_2501_v62_: D2
                        d_2501_v62_ = D2_DC7(d_2493_v54_)
                        d_2502_v63_: _dafny.Map
                        d_2502_v63_ = _dafny.Map({d_2465_v35_: d_2501_v62_})
                        d_2503_v64_: _dafny.Set
                        d_2503_v64_ = _dafny.Set({d_2471_v37_, d_2471_v37_})
                        d_2504_v65_: C2
                        nw442_ = C2()
                        nw442_.ctor__(d_2503_v64_, d_2465_v35_, self.f7)
                        d_2504_v65_ = nw442_
                        d_2505_v66_: _dafny.Map
                        d_2505_v66_ = _dafny.Map({d_2491_v52_: d_2504_v65_})
                        d_2506_v67_: _dafny.Map
                        d_2506_v67_ = _dafny.Map({p2: d_2505_v66_})
                        d_2507_v68_: D3
                        d_2507_v68_ = D3_DC10(d_2471_v37_)
                        d_2508_v69_: _dafny.Array
                        nw443_ = _dafny.Array(None, 4)
                        nw443_[int(0)] = (d_2507_v68_).cf21
                        nw443_[int(1)] = d_2471_v37_
                        nw443_[int(2)] = d_2471_v37_
                        nw443_[int(3)] = d_2471_v37_
                        d_2508_v69_ = nw443_
                        index416_ = default__.safeIndex(351, (d_2508_v69_).length(0))
                        (d_2508_v69_)[index416_] = d_2471_v37_
                        index417_ = default__.safeIndex(351, (d_2508_v69_).length(0))
                        rhs393_ = ((d_2496_v57_) != (d_2496_v57_) if (_dafny.SeqWithoutIsStrInference([p1 for d_2509_i9_ in range(default__.abs(587))])) < (d_2499_v60_) else default__.fm1((d_2464_v34_).set(default__.safeIndex(d_2491_v52_, len(d_2464_v34_)), d_2463_v33_), (0) - (d_2491_v52_), 805, globalState))
                        rhs394_ = (d_2502_v63_) | (d_2502_v63_)
                        rhs395_ = d_2506_v67_
                        rhs396_ = d_2471_v37_
                        lhs244_ = d_2473_v38_
                        lhs245_ = d_2508_v69_
                        lhs246_ = default__.safeIndex(351, (d_2508_v69_).length(0))
                        lhs244_.f7 = rhs393_
                        d_2502_v63_ = rhs394_
                        d_2506_v67_ = rhs395_
                        lhs245_[lhs246_] = rhs396_
                        index418_ = default__.safeIndex(315, (d_2471_v37_).length(0))
                        (d_2471_v37_)[index418_] = (self.f7 if False else d_2473_v38_.f7)
                    elif True:
                        d_2510_v70_: int
                        d_2510_v70_ = 126
                        d_2511_v71_: _dafny.Map
                        d_2511_v71_ = _dafny.Map({True: d_2471_v37_})
                        d_2510_v70_ = len(d_2511_v71_)
                        d_2512_v72_: C5
                        nw444_ = C5()
                        nw444_.ctor__(p2)
                        d_2512_v72_ = nw444_
                        d_2513_v73_: _dafny.Seq
                        d_2513_v73_ = _dafny.SeqWithoutIsStrInference([d_2471_v37_])
                        d_2514_v74_: _dafny.Set
                        d_2514_v74_ = _dafny.Set({d_2471_v37_, d_2471_v37_, (d_2513_v73_)[default__.safeIndex(d_2510_v70_, len(d_2513_v73_))]})
                        d_2515_v75_: C2
                        nw445_ = C2()
                        nw445_.ctor__(d_2514_v74_, d_2465_v35_, (len(_dafny.SeqWithoutIsStrInference([(D13_DC36(d_2465_v35_, _dafny.SeqWithoutIsStrInference([p0, d_2473_v38_.f7]), True)).cf58 for d_2516_i10_ in range(default__.abs(534))]))) == (d_2510_v70_))
                        d_2515_v75_ = nw445_
                        d_2471_v37_ = d_2471_v37_
                        (d_2515_v75_).m3(globalState)
                    pass
            pass
        hi16_ = p1
        for d_2517_i11_ in range(p1, hi16_):
            d_2518_v76_: _dafny.MultiSet
            d_2518_v76_ = _dafny.MultiSet([621])
            d_2519_v77_: D12
            d_2519_v77_ = D12_DC29((d_2518_v76_) | (d_2518_v76_))
            source18_ = d_2519_v77_
            if source18_.is_DC30:
                d_2520___mcc_h0_ = source18_.cf49
                d_2521___mcc_h1_ = source18_.cf50
                d_2522___mcc_h2_ = source18_.cf51
                d_2523___mcc_h3_ = source18_.cf52
                d_2524_cf52_ = d_2523___mcc_h3_
                d_2525_cf51_ = d_2522___mcc_h2_
                d_2526_cf50_ = d_2521___mcc_h1_
                d_2527_cf49_ = d_2520___mcc_h0_
                d_2462_v32_ = d_2462_v32_
                d_2528_v78_: _dafny.Array
                nw446_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_2528_v78_ = nw446_
                d_2529_v79_: _dafny.Seq
                d_2529_v79_ = _dafny.SeqWithoutIsStrInference([d_2528_v78_, d_2528_v78_])
                d_2528_v78_ = (d_2529_v79_)[default__.safeIndex(d_2517_i11_, len(d_2529_v79_))]
                d_2527_cf49_ = default__.fm0(globalState)
                d_2462_v32_ = ((d_2462_v32_) + ((d_2462_v32_).set(default__.safeIndex(d_2524_cf52_, len(d_2462_v32_)), d_2465_v35_))) + (d_2462_v32_)
            elif source18_.is_DC31:
                d_2530___mcc_h4_ = source18_.cf53
                d_2531___mcc_h5_ = source18_.cf54
                d_2532_cf54_ = d_2531___mcc_h5_
                d_2533_cf53_ = d_2530___mcc_h4_
                d_2534_v80_: _dafny.MultiSet
                d_2534_v80_ = _dafny.MultiSet([self.f7, self.f7])
                d_2535_v81_: C3
                nw447_ = C3()
                nw447_.ctor__((((d_2534_v80_).set(p0, default__.abs(d_2517_i11_))).cardinality) - (p1), p2)
                d_2535_v81_ = nw447_
                d_2536_v82_: _dafny.Seq
                d_2536_v82_ = _dafny.SeqWithoutIsStrInference([d_2517_i11_])
                d_2537_v83_: _dafny.Seq
                d_2537_v83_ = _dafny.SeqWithoutIsStrInference([(self).fm8((d_2536_v82_)[default__.safeIndex((d_2535_v81_).f14, len(d_2536_v82_))], globalState), p2, self.f7])
                (self).f7 = (d_2537_v83_)[default__.safeIndex(((d_2518_v76_).cardinality) + ((d_2535_v81_).f14), len(d_2537_v83_))]
                d_2538_v84_: _dafny.Map
                d_2538_v84_ = _dafny.Map({self.f7: len(d_2462_v32_)})
                d_2539_v85_: _dafny.Map
                d_2539_v85_ = _dafny.Map({728: (d_2535_v81_).f14})
                d_2540_v86_: _dafny.Array
                nw448_ = _dafny.Array(None, 15)
                nw448_[int(0)] = d_2517_i11_
                nw448_[int(1)] = p1
                nw448_[int(2)] = (d_2535_v81_).f14
                nw448_[int(3)] = d_2517_i11_
                nw448_[int(4)] = d_2517_i11_
                nw448_[int(5)] = len(d_2538_v84_)
                nw448_[int(6)] = d_2517_i11_
                nw448_[int(7)] = (d_2535_v81_).fm9(p0, p1, _dafny.CodePoint('b'), globalState)
                nw448_[int(8)] = d_2517_i11_
                nw448_[int(9)] = d_2517_i11_
                nw448_[int(10)] = p1
                nw448_[int(11)] = (d_2535_v81_).f14
                nw448_[int(12)] = len((d_2539_v85_).set(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))))
                nw448_[int(13)] = 794
                nw448_[int(14)] = d_2517_i11_
                d_2540_v86_ = nw448_
                d_2541_v87_: _dafny.MultiSet
                d_2541_v87_ = _dafny.MultiSet([d_2540_v86_, d_2540_v86_, d_2540_v86_, d_2540_v86_, d_2540_v86_])
                d_2541_v87_ = d_2541_v87_
                d_2532_cf54_ = not(((_dafny.SeqWithoutIsStrInference([d_2465_v35_ for d_2542_i12_ in range(default__.abs(180))])) + (default__.fm20(p2, d_2532_cf54_, globalState))) < (d_2462_v32_))
            elif source18_.is_DC32:
                d_2543___mcc_h6_ = source18_.cf55
                d_2544_cf55_ = d_2543___mcc_h6_
                d_2545_v88_: _dafny.Seq
                d_2545_v88_ = _dafny.SeqWithoutIsStrInference([self.f7, (self).fm8(d_2517_i11_, globalState), p2])
                d_2546_v89_: _dafny.Set
                d_2546_v89_ = _dafny.Set({len((d_2545_v88_).set(default__.safeIndex(p1, len(d_2545_v88_)), self.f7)), d_2517_i11_, default__.fm0(globalState), d_2517_i11_, p1})
                d_2547_v90_: _dafny.Map
                d_2547_v90_ = _dafny.Map({p1: d_2546_v89_})
                d_2547_v90_ = d_2547_v90_
                rhs397_ = d_2544_cf55_
                rhs398_ = (p0 if (p1) <= ((0) - (p1)) else not(d_2544_cf55_))
                d_2544_cf55_ = rhs397_
                d_2544_cf55_ = rhs398_
                d_2465_v35_ = d_2465_v35_
                d_2548_v91_: int
                d_2548_v91_ = -542
                d_2549_v92_: _dafny.Seq
                d_2549_v92_ = _dafny.SeqWithoutIsStrInference([(0) - (p1)])
                d_2548_v91_ = default__.safeDivisionInt(p1, (d_2549_v92_)[default__.safeIndex(p1, len(d_2549_v92_))])
            elif source18_.is_DC29:
                d_2550___mcc_h7_ = source18_.cf48
                d_2551_cf48_ = d_2550___mcc_h7_
                (self).f7 = False
                (self).f7 = self.f7
                d_2552_v93_: _dafny.Set
                d_2552_v93_ = _dafny.Set({d_2517_i11_})
                r1 = default__.fm27((_dafny.Set({p1, d_2517_i11_, p1})).ispropersubset(d_2552_v93_), d_2465_v35_, self.f7, not (p2) or (p2), globalState)
                d_2553_v94_: _dafny.Array
                def lambda153_(d_2554_i11_):
                    def lambda154_(d_2555_i13_):
                        return (d_2555_i13_) - (d_2554_i11_)

                    return lambda154_

                init87_ = lambda153_(d_2517_i11_)
                nw449_ = _dafny.Array(None, 1)
                for i0_87_ in range(nw449_.length(0)):
                    nw449_[i0_87_] = init87_(i0_87_)
                d_2553_v94_ = nw449_
                index419_ = default__.safeIndex(693, (d_2553_v94_).length(0))
                (d_2553_v94_)[index419_] = d_2517_i11_
                index420_ = default__.safeIndex(693, (d_2553_v94_).length(0))
                (d_2553_v94_)[index420_] = d_2517_i11_
            elif True:
                d_2556___mcc_h8_ = source18_.cf56
                d_2557_cf56_ = d_2556___mcc_h8_
                d_2558_v95_: _dafny.Map
                d_2558_v95_ = _dafny.Map({d_2517_i11_: p2})
                d_2559_v96_: _dafny.Seq
                d_2559_v96_ = _dafny.SeqWithoutIsStrInference([len(default__.fm27(p2, d_2465_v35_, p0, ((d_2558_v95_)[d_2517_i11_] if (d_2517_i11_) in (d_2558_v95_) else self.f7), globalState))])
                rhs399_ = (d_2517_i11_) == ((d_2517_i11_) - (len(d_2559_v96_)))
                rhs400_ = (p0) and (p2)
                lhs247_ = self
                lhs248_ = self
                lhs247_.f7 = rhs399_
                lhs248_.f7 = rhs400_
                d_2560_v97_: int
                d_2560_v97_ = 971
                d_2561_v98_: _dafny.Map
                d_2561_v98_ = _dafny.Map({False: (self).fm8(len((d_2462_v32_).set(default__.safeIndex(379, len(d_2462_v32_)), d_2465_v35_)), globalState)})
                d_2560_v97_ = (0) - ((0) - ((d_2560_v97_) + (len((d_2561_v98_ if p0 else d_2561_v98_)))))
                d_2562_v99_: D6
                d_2562_v99_ = D6_DC15(d_2465_v35_)
                d_2563_v100_: D9
                d_2563_v100_ = D9_DC25()
                d_2564_v101_: _dafny.Map
                d_2564_v101_ = _dafny.Map({not(p0): d_2462_v32_})
                rhs401_ = d_2462_v32_
                rhs402_ = default__.fm37(False, d_2563_v100_, globalState)
                rhs403_ = p1
                rhs404_ = d_2562_v99_
                rhs405_ = ((d_2564_v101_)[(d_2560_v97_) == (d_2560_v97_)] if ((d_2560_v97_) == (d_2560_v97_)) in (d_2564_v101_) else (d_2462_v32_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrwpnkxg"))))
                d_2462_v32_ = rhs401_
                d_2465_v35_ = rhs402_
                d_2560_v97_ = rhs403_
                d_2562_v99_ = rhs404_
                d_2462_v32_ = rhs405_
                d_2565_v102_: _dafny.MultiSet
                d_2565_v102_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2465_v35_ for d_2566_i14_ in range(default__.abs(-34))]), d_2462_v32_])
                rhs406_ = d_2565_v102_
                rhs407_ = len(d_2462_v32_)
                d_2565_v102_ = rhs406_
                d_2560_v97_ = rhs407_
            if True:
                d_2567_v103_: _dafny.Array
                nw450_ = _dafny.Array(False, 27)
                d_2567_v103_ = nw450_
                d_2568_v104_: _dafny.Set
                d_2568_v104_ = _dafny.Set({d_2567_v103_, d_2567_v103_})
                d_2569_v105_: C2
                nw451_ = C2()
                nw451_.ctor__(d_2568_v104_, d_2465_v35_, not(p2))
                d_2569_v105_ = nw451_
                d_2569_v105_ = d_2569_v105_
                d_2570_v106_: _dafny.Array
                def lambda155_(d_2571_p2_):
                    def lambda156_(d_2572_i15_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.Map({493: d_2571_p2_})])

                    return lambda156_

                init88_ = lambda155_(p2)
                nw452_ = _dafny.Array(None, 1)
                for i0_88_ in range(nw452_.length(0)):
                    nw452_[i0_88_] = init88_(i0_88_)
                d_2570_v106_ = nw452_
                d_2573_v107_: _dafny.Map
                d_2573_v107_ = _dafny.Map({897: p2})
                d_2574_v108_: _dafny.Seq
                d_2574_v108_ = _dafny.SeqWithoutIsStrInference([d_2573_v107_, d_2573_v107_, d_2573_v107_, d_2573_v107_, d_2573_v107_])
                index421_ = default__.safeIndex(445, (d_2570_v106_).length(0))
                (d_2570_v106_)[index421_] = ((d_2574_v108_).set(default__.safeIndex(d_2517_i11_, len(d_2574_v108_)), d_2573_v107_)) + ((d_2574_v108_).set(default__.safeIndex(d_2517_i11_, len(d_2574_v108_)), d_2573_v107_))
                index422_ = default__.safeIndex(445, (d_2570_v106_).length(0))
                (d_2570_v106_)[index422_] = d_2574_v108_
                d_2575_v109_: _dafny.Array
                nw453_ = _dafny.Array(int(0), 23)
                d_2575_v109_ = nw453_
                d_2575_v109_ = d_2575_v109_
                d_2576_v110_: _dafny.Array
                nw454_ = _dafny.Array(_dafny.Seq({}), 13)
                d_2576_v110_ = nw454_
                index423_ = default__.safeIndex(320, (d_2576_v110_).length(0))
                (d_2576_v110_)[index423_] = d_2466_v36_
                d_2577_v111_: _dafny.Array
                nw455_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                d_2577_v111_ = nw455_
                d_2578_v112_: _dafny.Array
                def lambda157_(d_2579_v35_):
                    def lambda158_(d_2580_i16_):
                        return _dafny.SeqWithoutIsStrInference([d_2579_v35_ for d_2581_i17_ in range(default__.abs(187))])

                    return lambda158_

                init89_ = lambda157_(d_2465_v35_)
                nw456_ = _dafny.Array(None, 6)
                for i0_89_ in range(nw456_.length(0)):
                    nw456_[i0_89_] = init89_(i0_89_)
                d_2578_v112_ = nw456_
                index424_ = default__.safeIndex(320, (d_2576_v110_).length(0))
                rhs408_ = d_2466_v36_
                rhs409_ = d_2567_v103_
                rhs410_ = d_2577_v111_
                rhs411_ = (d_2569_v105_).f16
                rhs412_ = d_2578_v112_
                lhs249_ = d_2576_v110_
                lhs250_ = default__.safeIndex(320, (d_2576_v110_).length(0))
                lhs249_[lhs250_] = rhs408_
                d_2567_v103_ = rhs409_
                d_2577_v111_ = rhs410_
                d_2465_v35_ = rhs411_
                d_2578_v112_ = rhs412_
                d_2582_v113_: _dafny.Seq
                d_2582_v113_ = _dafny.SeqWithoutIsStrInference([self.f7])
                d_2583_v114_: C6
                nw457_ = C6()
                nw457_.ctor__(not(p2), not((d_2582_v113_) == (d_2582_v113_)))
                d_2583_v114_ = nw457_
            elif True:
                d_2584_v115_: _dafny.Map
                d_2584_v115_ = _dafny.Map({p2: self.f7})
                d_2585_v116_: _dafny.Map
                d_2585_v116_ = d_2584_v115_
                r0 = (d_2585_v116_)
                d_2586_v117_: _dafny.Seq
                d_2586_v117_ = _dafny.SeqWithoutIsStrInference([p0, p2])
                (self).f7 = (d_2586_v117_)[default__.safeIndex(default__.safeDivisionInt(default__.fm0(globalState), d_2517_i11_), len(d_2586_v117_))]
                d_2587_v118_: T1
                nw458_ = C6()
                nw458_.ctor__((d_2517_i11_) == ((_dafny.MultiSet([p0, p0, p0, self.f7, p2])).cardinality), p0)
                d_2587_v118_ = nw458_
                d_2587_v118_ = d_2587_v118_
                (d_2587_v118_).f7 = self.f7
                (self).f7 = (self).fm8(p1, globalState)
            d_2588_v119_: int
            d_2588_v119_ = 322
            d_2588_v119_ = ((0) - (default__.safeModuloInt(d_2517_i11_, d_2517_i11_))) * (d_2517_i11_)
            (self).f7 = self.f7
        d_2589_v120_: D6
        d_2589_v120_ = D6_DC17(p0, p1, d_2465_v35_, d_2465_v35_)
        d_2590_v121_: D6
        d_2590_v121_ = D6_DC18(d_2589_v120_)
        pat_let_tv59_ = p2
        pat_let_tv60_ = p2
        pat_let_tv61_ = p0
        pat_let_tv62_ = d_2462_v32_
        pat_let_tv63_ = p2
        pat_let_tv64_ = p2
        def lambda159_(source19_):
            if source19_.is_DC16:
                d_2591___mcc_h9_ = source19_.cf30
                d_2592_cf30_ = d_2591___mcc_h9_
                return (_dafny.Map({not(d_2592_cf30_): pat_let_tv59_})) | (_dafny.Map({True: self.f7}))
            elif source19_.is_DC17:
                d_2593___mcc_h10_ = source19_.cf31
                d_2594___mcc_h11_ = source19_.cf32
                d_2595___mcc_h12_ = source19_.cf33
                d_2596___mcc_h13_ = source19_.cf34
                d_2597_cf34_ = d_2596___mcc_h13_
                d_2598_cf33_ = d_2595___mcc_h12_
                d_2599_cf32_ = d_2594___mcc_h11_
                d_2600_cf31_ = d_2593___mcc_h10_
                return (_dafny.Map({self.f7: False})).set(pat_let_tv60_, not(d_2600_cf31_))
            elif source19_.is_DC15:
                d_2601___mcc_h14_ = source19_.cf29
                d_2602_cf29_ = d_2601___mcc_h14_
                return (_dafny.Map({(D2_DC8(pat_let_tv61_, pat_let_tv62_)).cf14: False})).set(not(not(True)), self.f7)
            elif True:
                d_2603___mcc_h15_ = source19_.cf35
                d_2604_cf35_ = d_2603___mcc_h15_
                return _dafny.Map({pat_let_tv63_: pat_let_tv64_})

        r0 = lambda159_(d_2590_v121_)
        d_2605_v122_: _dafny.Set
        d_2605_v122_ = _dafny.Set({p0})
        r1 = (d_2605_v122_) - (d_2605_v122_)
        return r0, r1

    def m3(self, globalState):
        d_2606_v0_: int
        d_2606_v0_ = -924
        d_2607_v1_: _dafny.Seq
        d_2607_v1_ = _dafny.SeqWithoutIsStrInference([self.f7, True])
        d_2608_v2_: str
        d_2608_v2_ = _dafny.CodePoint('q')
        d_2609_v3_: _dafny.Set
        d_2609_v3_ = _dafny.Set({d_2608_v2_, d_2608_v2_})
        d_2610_v5_: _dafny.Map
        def iife207_():
            coll53_ = _dafny.Set()
            compr_53_: str
            for compr_53_ in (d_2609_v3_).Elements:
                d_2611_v4_: str = compr_53_
                if (d_2611_v4_) in (d_2609_v3_):
                    coll53_ = coll53_.union(_dafny.Set([d_2611_v4_]))
            return _dafny.Set(coll53_)
        d_2610_v5_ = _dafny.Map({len(iife207_()
        ): self.f7})
        d_2612_v6_: _dafny.Map
        d_2612_v6_ = _dafny.Map({len(d_2610_v5_): d_2606_v0_})
        hi17_ = ((0) - (len(d_2607_v1_))) * (len(d_2612_v6_))
        for d_2613_i0_ in range(d_2606_v0_, hi17_):
            d_2606_v0_ = d_2606_v0_
            d_2614_v7_: _dafny.MultiSet
            d_2614_v7_ = _dafny.MultiSet([self.f7])
            d_2615_v8_: _dafny.Seq
            d_2615_v8_ = _dafny.SeqWithoutIsStrInference([d_2614_v7_])
            d_2616_v9_: _dafny.MultiSet
            d_2616_v9_ = _dafny.MultiSet([(d_2615_v8_)[default__.safeIndex(len(d_2607_v1_), len(d_2615_v8_))]])
            d_2616_v9_ = (d_2616_v9_) | (d_2616_v9_)
            (self).f7 = self.f7
            d_2617_v10_: _dafny.Seq
            d_2617_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygaj"))
            if (d_2608_v2_) not in (d_2617_v10_):
                d_2618_v11_: _dafny.MultiSet
                d_2618_v11_ = _dafny.MultiSet([d_2613_i0_])
                d_2618_v11_ = d_2618_v11_
                d_2619_v12_: _dafny.Array
                nw459_ = _dafny.Array(_dafny.Map({}), 7)
                d_2619_v12_ = nw459_
                index425_ = default__.safeIndex(81, (d_2619_v12_).length(0))
                (d_2619_v12_)[index425_] = (_dafny.Map({d_2613_i0_: -272})) | (d_2612_v6_)
                d_2620_v13_: _dafny.Seq
                d_2620_v13_ = _dafny.SeqWithoutIsStrInference([d_2612_v6_])
                index426_ = default__.safeIndex(81, (d_2619_v12_).length(0))
                (d_2619_v12_)[index426_] = (d_2620_v13_)[default__.safeIndex(d_2613_i0_, len(d_2620_v13_))]
                (self).f7 = (default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference([d_2608_v2_ for d_2621_i1_ in range(default__.abs(775))])).set(default__.safeIndex(d_2613_i0_, len(_dafny.SeqWithoutIsStrInference([d_2608_v2_ for d_2622_i1_ in range(default__.abs(775))]))), d_2608_v2_)), d_2606_v0_)) != (d_2613_i0_)
                d_2623_v14_: _dafny.Set
                d_2623_v14_ = _dafny.Set({d_2613_i0_, (0) - (d_2613_i0_), d_2606_v0_})
                d_2606_v0_ = (0) - (len(d_2623_v14_))
                d_2624_v15_: _dafny.Seq
                d_2624_v15_ = _dafny.SeqWithoutIsStrInference([d_2606_v0_])
                d_2606_v0_ = (d_2624_v15_)[default__.safeIndex((0) - ((0) - (d_2606_v0_)), len(d_2624_v15_))]
            elif True:
                (self).f7 = self.f7
                (self).f7 = self.f7
                d_2625_v16_: _dafny.Set
                d_2625_v16_ = d_2609_v3_
                d_2626_v17_: _dafny.Seq
                d_2626_v17_ = _dafny.SeqWithoutIsStrInference([d_2625_v16_])
                d_2627_v18_: _dafny.Map
                d_2627_v18_ = _dafny.Map({(((self).f8)[self.f7] if (self.f7) in ((self).f8) else _dafny.SeqWithoutIsStrInference([d_2613_i0_])): (d_2626_v17_).set(default__.safeIndex(len(d_2617_v10_), len(d_2626_v17_)), d_2625_v16_)})
                d_2627_v18_ = d_2627_v18_
                d_2628_v19_: _dafny.Map
                d_2628_v19_ = _dafny.Map({self.f7: d_2608_v2_})
                (self).f7 = (default__.safeDivisionInt(len(d_2628_v19_), d_2606_v0_)) < (((d_2612_v6_)[len(d_2617_v10_)] if (len(d_2617_v10_)) in (d_2612_v6_) else d_2606_v0_))
                d_2629_v20_: _dafny.Set
                d_2629_v20_ = _dafny.Set({self.f7})
                d_2630_v21_: D1
                d_2630_v21_ = D1_DC3(d_2629_v20_)
                d_2631_v22_: D0
                d_2631_v22_ = D0_DC0(d_2606_v0_)
                d_2632_v23_: D8
                d_2632_v23_ = D8_DC22(939, d_2630_v21_, d_2631_v22_)
                rhs413_ = d_2606_v0_
                rhs414_ = default__.fm55(d_2608_v2_, globalState)
                rhs415_ = d_2617_v10_
                rhs416_ = d_2610_v5_
                d_2606_v0_ = rhs413_
                d_2632_v23_ = rhs414_
                d_2617_v10_ = rhs415_
                d_2610_v5_ = rhs416_
        d_2633_v24_: _dafny.Set
        d_2633_v24_ = _dafny.Set({d_2606_v0_, len(_dafny.SeqWithoutIsStrInference([self.f7])), (self).fm9(self.f7, 958, d_2608_v2_, globalState)})
        if (d_2633_v24_).ispropersubset(d_2633_v24_):
            d_2634_v25_: _dafny.Array
            nw460_ = _dafny.Array(False, 26)
            d_2634_v25_ = nw460_
            index427_ = default__.safeIndex(997, (d_2634_v25_).length(0))
            (d_2634_v25_)[index427_] = self.f7
            index428_ = default__.safeIndex(997, (d_2634_v25_).length(0))
            (d_2634_v25_)[index428_] = self.f7
            d_2635_v26_: _dafny.Seq
            d_2635_v26_ = _dafny.SeqWithoutIsStrInference([-672])
            d_2636_v27_: _dafny.Seq
            d_2636_v27_ = _dafny.SeqWithoutIsStrInference([d_2635_v26_, d_2635_v26_, (_dafny.SeqWithoutIsStrInference([d_2606_v0_ for d_2637_i3_ in range(default__.abs(89))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owdl"))), len(_dafny.SeqWithoutIsStrInference([d_2606_v0_ for d_2638_i3_ in range(default__.abs(89))]))), len(default__.fm20(self.f7, self.f7, globalState)))])
            if ((_dafny.SeqWithoutIsStrInference([d_2606_v0_ for d_2639_i2_ in range(default__.abs(349))])) + (_dafny.SeqWithoutIsStrInference([d_2606_v0_]))) not in (d_2636_v27_):
                d_2640_v28_: _dafny.Array
                nw461_ = _dafny.Array(_dafny.Map({}), 19)
                d_2640_v28_ = nw461_
                d_2641_v29_: _dafny.MultiSet
                d_2641_v29_ = _dafny.MultiSet([True])
                d_2642_v30_: _dafny.Map
                d_2642_v30_ = _dafny.Map({d_2606_v0_: d_2641_v29_})
                index429_ = default__.safeIndex(806, (d_2640_v28_).length(0))
                (d_2640_v28_)[index429_] = d_2642_v30_
                index430_ = default__.safeIndex(806, (d_2640_v28_).length(0))
                (d_2640_v28_)[index430_] = d_2642_v30_
                d_2606_v0_ = (0) - ((0) - (d_2606_v0_))
                d_2606_v0_ = d_2606_v0_
                d_2606_v0_ = (0) - ((0) - (len(d_2607_v1_)))
                d_2643_v31_: _dafny.Seq
                d_2643_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cljigeqyq"))
                d_2643_v31_ = d_2643_v31_
            elif True:
                d_2644_v32_: _dafny.Array
                nw462_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_2644_v32_ = nw462_
                d_2645_v33_: D1
                d_2645_v33_ = D1_DC3(_dafny.Set({self.f7, self.f7}))
                d_2646_v34_: _dafny.Map
                d_2646_v34_ = _dafny.Map({(d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))]: (d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))]})
                d_2647_v35_: C3
                nw463_ = C3()
                nw463_.ctor__((d_2635_v26_)[default__.safeIndex(len(d_2646_v34_), len(d_2635_v26_))], (d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))])
                d_2647_v35_ = nw463_
                d_2648_v36_: _dafny.Map
                d_2648_v36_ = _dafny.Map({(d_2645_v33_).cf6: d_2647_v35_})
                d_2649_v37_: _dafny.Array
                nw464_ = _dafny.Array(None, 10)
                nw464_[int(0)] = d_2606_v0_
                nw464_[int(1)] = 339
                nw464_[int(2)] = -621
                nw464_[int(3)] = d_2606_v0_
                nw464_[int(4)] = len(d_2607_v1_)
                nw464_[int(5)] = d_2606_v0_
                nw464_[int(6)] = 585
                nw464_[int(7)] = len(d_2648_v36_)
                nw464_[int(8)] = (d_2647_v35_).f14
                nw464_[int(9)] = d_2606_v0_
                d_2649_v37_ = nw464_
                d_2650_v38_: _dafny.Seq
                d_2650_v38_ = _dafny.SeqWithoutIsStrInference([d_2649_v37_, d_2649_v37_])
                d_2651_v39_: _dafny.Map
                d_2651_v39_ = _dafny.Map({d_2634_v25_: d_2649_v37_})
                d_2652_v40_: _dafny.Array
                nw465_ = _dafny.Array(None, 15)
                nw465_[int(0)] = d_2649_v37_
                nw465_[int(1)] = (d_2650_v38_)[default__.safeIndex(d_2606_v0_, len(d_2650_v38_))]
                nw465_[int(2)] = d_2649_v37_
                nw465_[int(3)] = d_2649_v37_
                nw465_[int(4)] = d_2649_v37_
                nw465_[int(5)] = d_2649_v37_
                nw465_[int(6)] = d_2649_v37_
                nw465_[int(7)] = d_2649_v37_
                nw465_[int(8)] = d_2649_v37_
                nw465_[int(9)] = d_2649_v37_
                nw465_[int(10)] = d_2649_v37_
                nw465_[int(11)] = d_2649_v37_
                nw465_[int(12)] = ((d_2651_v39_)[d_2634_v25_] if (d_2634_v25_) in (d_2651_v39_) else d_2649_v37_)
                nw465_[int(13)] = d_2649_v37_
                nw465_[int(14)] = d_2649_v37_
                d_2652_v40_ = nw465_
                index431_ = default__.safeIndex(185, (d_2644_v32_).length(0))
                (d_2644_v32_)[index431_] = d_2652_v40_
                index432_ = default__.safeIndex(185, (d_2644_v32_).length(0))
                (d_2644_v32_)[index432_] = (d_2652_v40_ if self.f7 else d_2652_v40_)
                d_2653_v41_: _dafny.Set
                d_2653_v41_ = _dafny.Set({d_2649_v37_})
                index433_ = default__.safeIndex(997, (d_2634_v25_).length(0))
                rhs417_ = self.f7
                rhs418_ = (-325) + (d_2606_v0_)
                rhs419_ = ((d_2653_v41_) - (d_2653_v41_)) - (_dafny.Set({d_2649_v37_}))
                rhs420_ = (0) - ((d_2606_v0_) - ((d_2647_v35_).f14))
                lhs251_ = d_2634_v25_
                lhs252_ = default__.safeIndex(997, (d_2634_v25_).length(0))
                lhs251_[lhs252_] = rhs417_
                d_2606_v0_ = rhs418_
                d_2653_v41_ = rhs419_
                d_2606_v0_ = rhs420_
                d_2654_v42_: _dafny.Seq
                d_2654_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpswxaqat"))
                rhs421_ = len(d_2654_v42_)
                rhs422_ = (d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))]
                rhs423_ = (d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))]
                lhs253_ = self
                lhs254_ = self
                d_2606_v0_ = rhs421_
                lhs253_.f7 = rhs422_
                lhs254_.f7 = rhs423_
                d_2606_v0_ = (default__.safeDivisionInt((d_2647_v35_).f14, (_dafny.MultiSet([(d_2647_v35_).f14, d_2606_v0_])).cardinality)) * (((d_2647_v35_).f14) * (len(_dafny.SeqWithoutIsStrInference([self.f7]))))
                d_2655_v43_: D0
                d_2655_v43_ = D0_DC0((d_2647_v35_).f14)
                d_2656_v44_: C6
                nw466_ = C6()
                nw466_.ctor__(default__.fm1(_dafny.SeqWithoutIsStrInference([d_2655_v43_, default__.fm56(globalState)]), (d_2647_v35_).f14, (0) - (d_2606_v0_), globalState), (d_2634_v25_) not in (_dafny.SeqWithoutIsStrInference([d_2634_v25_, d_2634_v25_])))
                d_2656_v44_ = nw466_
            (self).f7 = (d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))]
            d_2657_v45_: _dafny.Set
            d_2657_v45_ = _dafny.Set({d_2634_v25_, d_2634_v25_, d_2634_v25_})
            d_2658_v46_: D8
            d_2658_v46_ = D8_DC20(d_2657_v45_)
            d_2658_v46_ = d_2658_v46_
            if self.f7:
                d_2659_v47_: C5
                nw467_ = C5()
                nw467_.ctor__(True)
                d_2659_v47_ = nw467_
                d_2606_v0_ = len(_dafny.Map({d_2634_v25_: d_2606_v0_}))
                d_2660_v48_: D0
                d_2660_v48_ = D0_DC1(d_2606_v0_)
                d_2660_v48_ = d_2660_v48_
                (self).f7 = (d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))]
                d_2661_v49_: _dafny.Array
                nw468_ = _dafny.Array(int(0), 6)
                d_2661_v49_ = nw468_
                index434_ = default__.safeIndex(144, (d_2661_v49_).length(0))
                (d_2661_v49_)[index434_] = (0) - (d_2606_v0_)
                d_2662_v50_: _dafny.Seq
                d_2662_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yiih"))
                index435_ = default__.safeIndex(144, (d_2661_v49_).length(0))
                (d_2661_v49_)[index435_] = len(d_2662_v50_)
            elif True:
                nw469_ = _dafny.Array(False, 9)
                d_2634_v25_ = nw469_
                d_2663_v51_: _dafny.Seq
                d_2663_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aykqqpd"))
                d_2663_v51_ = d_2663_v51_
                (self).f7 = (d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))]
                d_2664_v52_: _dafny.MultiSet
                d_2664_v52_ = _dafny.MultiSet([self.f7, self.f7])
                d_2665_v53_: _dafny.Array
                nw470_ = _dafny.Array(None, 17)
                nw470_[int(0)] = d_2606_v0_
                nw470_[int(1)] = d_2606_v0_
                nw470_[int(2)] = -997
                nw470_[int(3)] = len(d_2663_v51_)
                nw470_[int(4)] = default__.safeDivisionInt((d_2664_v52_).cardinality, d_2606_v0_)
                nw470_[int(5)] = (0) - (len(d_2663_v51_))
                nw470_[int(6)] = d_2606_v0_
                nw470_[int(7)] = (d_2606_v0_ if (d_2607_v1_)[default__.safeIndex(d_2606_v0_, len(d_2607_v1_))] else d_2606_v0_)
                nw470_[int(8)] = default__.safeModuloInt(len(d_2635_v26_), d_2606_v0_)
                nw470_[int(9)] = default__.safeDivisionInt((0) - (len(d_2663_v51_)), 802)
                nw470_[int(10)] = -196
                nw470_[int(11)] = d_2606_v0_
                nw470_[int(12)] = d_2606_v0_
                nw470_[int(13)] = d_2606_v0_
                nw470_[int(14)] = d_2606_v0_
                nw470_[int(15)] = len(d_2663_v51_)
                nw470_[int(16)] = d_2606_v0_
                d_2665_v53_ = nw470_
                index436_ = default__.safeIndex(716, (d_2665_v53_).length(0))
                (d_2665_v53_)[index436_] = d_2606_v0_
                index437_ = default__.safeIndex(716, (d_2665_v53_).length(0))
                (d_2665_v53_)[index437_] = d_2606_v0_
                d_2666_v54_: D1
                d_2666_v54_ = D1_DC4(self.f7, self.f7)
                d_2667_v55_: D1
                d_2667_v55_ = D1_DC6(d_2666_v54_)
                d_2668_v56_: D18
                d_2668_v56_ = D18_DC51(d_2607_v1_, d_2606_v0_, True)
                d_2669_v57_: _dafny.Seq
                d_2669_v57_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_2668_v56_).cf87]), d_2607_v1_, d_2607_v1_])
                rhs424_ = (d_2607_v1_).set(default__.safeIndex(d_2606_v0_, len(d_2607_v1_)), not((d_2634_v25_)[default__.safeIndex(997, (d_2634_v25_).length(0))]))
                rhs425_ = d_2665_v53_
                rhs426_ = (d_2606_v0_) >= (len(d_2669_v57_))
                rhs427_ = d_2667_v55_
                lhs255_ = self
                d_2607_v1_ = rhs424_
                d_2665_v53_ = rhs425_
                lhs255_.f7 = rhs426_
                d_2667_v55_ = rhs427_
        elif True:
            d_2670_v58_: _dafny.Map
            d_2670_v58_ = _dafny.Map({True: self.f7})
            d_2671_v59_: _dafny.Array
            nw471_ = _dafny.Array(False, 7)
            d_2671_v59_ = nw471_
            d_2672_v60_: C2
            nw472_ = C2()
            nw472_.ctor__(_dafny.Set({d_2671_v59_}), _dafny.CodePoint('t'), self.f7)
            d_2672_v60_ = nw472_
            d_2673_v61_: _dafny.Map
            d_2673_v61_ = _dafny.Map({self.f7: d_2672_v60_})
            d_2674_v62_: _dafny.Seq
            d_2674_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiamatg"))
            d_2670_v58_ = (d_2670_v58_).set((d_2606_v0_) < (len(d_2673_v61_)), (d_2674_v62_) != (d_2674_v62_))
            d_2675_v63_: _dafny.Seq
            d_2675_v63_ = _dafny.SeqWithoutIsStrInference([d_2674_v62_, d_2674_v62_])
            d_2676_v64_: C0
            nw473_ = C0()
            nw473_.ctor__((d_2606_v0_) < (d_2606_v0_), (_dafny.SeqWithoutIsStrInference([d_2608_v2_ for d_2677_i4_ in range(default__.abs(486))])) == ((d_2675_v63_)[default__.safeIndex(d_2606_v0_, len(d_2675_v63_))]))
            d_2676_v64_ = nw473_
            (self).f7 = True
            (d_2676_v64_).f13 = d_2676_v64_.f13
            d_2606_v0_ = (d_2606_v0_) * (((self).fm10(self.f7, globalState)) + (d_2606_v0_))
        d_2606_v0_ = d_2606_v0_
        d_2606_v0_ = d_2606_v0_
        d_2606_v0_ = 348
        d_2678_v65_: _dafny.Array
        def lambda160_(d_2679_v0_):
            def lambda161_(d_2680_i5_):
                return D9_DC26(_dafny.SeqWithoutIsStrInference([920, d_2679_v0_]))

            return lambda161_

        init90_ = lambda160_(d_2606_v0_)
        nw474_ = _dafny.Array(None, 4)
        for i0_90_ in range(nw474_.length(0)):
            nw474_[i0_90_] = init90_(i0_90_)
        d_2678_v65_ = nw474_
        d_2681_v66_: _dafny.Seq
        d_2681_v66_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_2606_v0_ for d_2682_i6_ in range(default__.abs(999))]))])
        index438_ = default__.safeIndex(579, (d_2678_v65_).length(0))
        (d_2678_v65_)[index438_] = D9_DC26(d_2681_v66_)
        d_2683_v67_: D9
        d_2683_v67_ = D9_DC26(d_2681_v66_)
        d_2684_v68_: _dafny.Map
        d_2684_v68_ = _dafny.Map({d_2606_v0_: (D9_DC26(d_2681_v66_)).cf45})
        d_2685_v69_: _dafny.Seq
        d_2685_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msetw"))
        pat_let_tv65_ = d_2684_v68_
        pat_let_tv66_ = d_2685_v69_
        pat_let_tv67_ = d_2685_v69_
        pat_let_tv68_ = d_2684_v68_
        pat_let_tv69_ = d_2681_v66_
        pat_let_tv70_ = d_2606_v0_
        index439_ = default__.safeIndex(579, (d_2678_v65_).length(0))
        def iife209_(_pat_let78_0):
            def iife210_(d_2686_dt__update__tmp_h0_):
                def iife211_(_pat_let79_0):
                    def iife212_(d_2687_dt__update_hcf45_h0_):
                        return D9_DC26(d_2687_dt__update_hcf45_h0_)
                    return iife212_(_pat_let79_0)
                return iife211_(((pat_let_tv65_)[len(pat_let_tv66_)] if (len(pat_let_tv67_)) in (pat_let_tv68_) else pat_let_tv69_))
            return iife210_(_pat_let78_0)
        def iife208_(_pat_let77_0):
            def iife213_(d_2688_dt__update__tmp_h1_):
                def iife214_(_pat_let80_0):
                    def iife215_(d_2690_dt__update_hcf45_h1_):
                        return D9_DC26(d_2690_dt__update_hcf45_h1_)
                    return iife215_(_pat_let80_0)
                return iife214_(_dafny.SeqWithoutIsStrInference([pat_let_tv70_ for d_2689_i7_ in range(default__.abs(944))]))
            return iife213_(_pat_let77_0)
        (d_2678_v65_)[index439_] = iife208_(iife209_(d_2683_v67_))

    def m4(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        (self).m5(globalState)
        d_2691_i0_: int
        d_2691_i0_ = 0
        with _dafny.label("18"):
            while (641) != ((0) - ((self).fm10(not(self.f7), globalState))):
                with _dafny.c_label("18"):
                    if (d_2691_i0_) >= (100):
                        raise _dafny.Break("18")
                    d_2691_i0_ = (d_2691_i0_) + (1)
                    d_2692_v0_: _dafny.Array
                    nw475_ = _dafny.Array(_dafny.MultiSet({}), 28)
                    d_2692_v0_ = nw475_
                    d_2693_v1_: D0
                    d_2693_v1_ = D0_DC0(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2694_i1_ in range(default__.abs(886))])))
                    d_2695_v2_: _dafny.Seq
                    d_2695_v2_ = _dafny.SeqWithoutIsStrInference([d_2693_v1_])
                    d_2696_v3_: _dafny.MultiSet
                    d_2696_v3_ = _dafny.MultiSet([self.f7, default__.fm1(d_2695_v2_, p0, len(p1), globalState)])
                    d_2697_v4_: _dafny.Seq
                    d_2697_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([self.f7, False]), d_2696_v3_])
                    index440_ = default__.safeIndex(467, (d_2692_v0_).length(0))
                    (d_2692_v0_)[index440_] = (d_2697_v4_)[default__.safeIndex(p0, len(d_2697_v4_))]
                    d_2698_v5_: D2
                    d_2698_v5_ = D2_DC8(self.f7, p1)
                    d_2699_v6_: _dafny.Seq
                    d_2699_v6_ = _dafny.SeqWithoutIsStrInference([(d_2698_v5_).cf14])
                    index441_ = default__.safeIndex(467, (d_2692_v0_).length(0))
                    (d_2692_v0_)[index441_] = (d_2696_v3_).set((d_2699_v6_)[default__.safeIndex(p0, len(d_2699_v6_))], default__.abs((p0) * (p0)))
                    d_2700_v7_: _dafny.Map
                    d_2700_v7_ = _dafny.Map({self.f7: self.f7})
                    (self).f7 = ((d_2700_v7_)[self.f7] if (self.f7) in (d_2700_v7_) else self.f7)
                    d_2701_v8_: _dafny.Array
                    nw476_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                    d_2701_v8_ = nw476_
                    d_2702_v9_: str
                    d_2702_v9_ = _dafny.CodePoint('a')
                    index442_ = default__.safeIndex(78, (d_2701_v8_).length(0))
                    (d_2701_v8_)[index442_] = d_2702_v9_
                    d_2703_v10_: _dafny.Array
                    def lambda162_(d_2704_p0_):
                        def lambda163_(d_2705_i2_):
                            return (41) > (d_2704_p0_)

                        return lambda163_

                    init91_ = lambda162_(p0)
                    nw477_ = _dafny.Array(None, 8)
                    for i0_91_ in range(nw477_.length(0)):
                        nw477_[i0_91_] = init91_(i0_91_)
                    d_2703_v10_ = nw477_
                    d_2706_v11_: _dafny.Set
                    d_2706_v11_ = _dafny.Set({True})
                    d_2707_v12_: _dafny.Seq
                    d_2707_v12_ = _dafny.SeqWithoutIsStrInference([d_2706_v11_])
                    index443_ = default__.safeIndex(542, (d_2703_v10_).length(0))
                    (d_2703_v10_)[index443_] = (len(p1)) < (len(d_2707_v12_))
                    d_2708_v13_: _dafny.Array
                    nw478_ = _dafny.Array(None, 4)
                    def iife216_(_pat_let81_0):
                        def iife217_(d_2709_dt__update__tmp_h0_):
                            def iife218_(_pat_let82_0):
                                def iife219_(d_2710_dt__update_hcf14_h0_):
                                    return D2_DC8(d_2710_dt__update_hcf14_h0_, (d_2709_dt__update__tmp_h0_).cf15)
                                return iife219_(_pat_let82_0)
                            return iife218_(self.f7)
                        return iife217_(_pat_let81_0)
                    nw478_[int(0)] = ((iife216_(D2_DC8(self.f7, p1))).cf15) + ((d_2698_v5_).cf15)
                    nw478_[int(1)] = p1
                    nw478_[int(2)] = p1
                    nw478_[int(3)] = p1
                    d_2708_v13_ = nw478_
                    index444_ = default__.safeIndex(380, (d_2708_v13_).length(0))
                    (d_2708_v13_)[index444_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2711_i3_ in range(default__.abs(291))])
                    index445_ = default__.safeIndex(78, (d_2701_v8_).length(0))
                    index446_ = default__.safeIndex(542, (d_2703_v10_).length(0))
                    index447_ = default__.safeIndex(380, (d_2708_v13_).length(0))
                    rhs428_ = (0) - (p0)
                    rhs429_ = (p1)[default__.safeIndex(p0, len(p1))]
                    rhs430_ = self.f7
                    rhs431_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgnlwup"))
                    lhs256_ = d_2701_v8_
                    lhs257_ = default__.safeIndex(78, (d_2701_v8_).length(0))
                    lhs258_ = d_2703_v10_
                    lhs259_ = default__.safeIndex(542, (d_2703_v10_).length(0))
                    lhs260_ = d_2708_v13_
                    lhs261_ = default__.safeIndex(380, (d_2708_v13_).length(0))
                    r1 = rhs428_
                    lhs256_[lhs257_] = rhs429_
                    lhs258_[lhs259_] = rhs430_
                    lhs260_[lhs261_] = rhs431_
                    d_2712_v14_: _dafny.Array
                    nw479_ = _dafny.Array(None, 6)
                    nw479_[int(0)] = p0
                    nw479_[int(1)] = p0
                    nw479_[int(2)] = -651
                    nw479_[int(3)] = p0
                    nw479_[int(4)] = p0
                    nw479_[int(5)] = 801
                    d_2712_v14_ = nw479_
                    d_2713_v15_: _dafny.Map
                    d_2713_v15_ = _dafny.Map({d_2712_v14_: _dafny.Set({p0})})
                    d_2714_v16_: D13
                    d_2714_v16_ = D13_DC34(d_2713_v15_)
                    d_2715_v17_: _dafny.Seq
                    d_2715_v17_ = _dafny.SeqWithoutIsStrInference([d_2714_v16_, d_2714_v16_, d_2714_v16_, d_2714_v16_, d_2714_v16_])
                    d_2716_v18_: D17
                    d_2716_v18_ = D17_DC49(p0, d_2715_v17_, p0, (d_2703_v10_)[default__.safeIndex(542, (d_2703_v10_).length(0))])
                    r1 = (d_2716_v18_).cf82
                    pass
            pass
        d_2717_v19_: _dafny.Set
        d_2717_v19_ = _dafny.Set({self.f7, self.f7})
        d_2718_v20_: str
        d_2718_v20_ = _dafny.CodePoint('n')
        rhs432_ = (self).fm9(self.f7, len(_dafny.Set({len(d_2717_v19_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnyywoju"))), 149, p0, -648})), d_2718_v20_, globalState)
        rhs433_ = 310
        r3 = rhs432_
        r3 = rhs433_
        d_2719_v21_: _dafny.Set
        d_2719_v21_ = _dafny.Set({p0, (self).fm9(self.f7, p0, d_2718_v20_, globalState)})
        d_2720_v22_: _dafny.Seq
        d_2720_v22_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2721_v23_: D1
        d_2721_v23_ = D1_DC5(self.f7, d_2720_v22_, p0)
        d_2722_v24_: D12
        d_2722_v24_ = D12_DC31(D9_DC26(d_2720_v22_), (d_2721_v23_).cf9)
        d_2723_v25_: _dafny.Seq
        d_2723_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2718_v20_ for d_2724_i4_ in range(default__.abs(-892))])])
        d_2725_v26_: _dafny.Map
        d_2725_v26_ = _dafny.Map({self.f7: _dafny.SeqWithoutIsStrInference([d_2718_v20_ for d_2726_i5_ in range(default__.abs(244))])})
        d_2727_v27_: _dafny.Seq
        d_2727_v27_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([p1])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p1]))), p1), d_2723_v25_, d_2723_v25_, _dafny.SeqWithoutIsStrInference([((d_2725_v26_)[self.f7] if (self.f7) in (d_2725_v26_) else _dafny.SeqWithoutIsStrInference([d_2718_v20_ for d_2728_i6_ in range(default__.abs(14))])), p1])])
        d_2729_v28_: D0
        d_2729_v28_ = D0_DC0(p0)
        d_2730_v29_: _dafny.Seq
        d_2730_v29_ = _dafny.SeqWithoutIsStrInference([d_2729_v28_, d_2729_v28_])
        d_2731_v30_: _dafny.Array
        def lambda164_(d_2732_p0_):
            def lambda165_(d_2733_i8_):
                return (d_2733_i8_) * (d_2732_p0_)

            return lambda165_

        init92_ = lambda164_(p0)
        nw480_ = _dafny.Array(None, 19)
        for i0_92_ in range(nw480_.length(0)):
            nw480_[i0_92_] = init92_(i0_92_)
        d_2731_v30_ = nw480_
        d_2734_v31_: _dafny.Map
        d_2734_v31_ = _dafny.Map({d_2731_v30_: d_2719_v21_})
        d_2735_v32_: D13
        d_2735_v32_ = D13_DC34(d_2734_v31_)
        d_2736_v33_: _dafny.Seq
        d_2736_v33_ = _dafny.SeqWithoutIsStrInference([d_2735_v32_])
        d_2737_v34_: _dafny.Seq
        d_2737_v34_ = _dafny.SeqWithoutIsStrInference([self.f7])
        d_2738_v35_: D17
        d_2738_v35_ = D17_DC49(p0, d_2736_v33_, len(d_2737_v34_), self.f7)
        d_2739_v36_: _dafny.Array
        nw481_ = _dafny.Array(None, 21)
        nw481_[int(0)] = self.f7
        nw481_[int(1)] = self.f7
        nw481_[int(2)] = self.f7
        nw481_[int(3)] = self.f7
        nw481_[int(4)] = (d_2722_v24_).cf54
        nw481_[int(5)] = (p1) not in ((((d_2727_v27_)[default__.safeIndex(p0, len(d_2727_v27_))]).set(default__.safeIndex(p0, len((d_2727_v27_)[default__.safeIndex(p0, len(d_2727_v27_))])), p1)).set(default__.safeIndex(p0, len(((d_2727_v27_)[default__.safeIndex(p0, len(d_2727_v27_))]).set(default__.safeIndex(p0, len((d_2727_v27_)[default__.safeIndex(p0, len(d_2727_v27_))])), p1))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eopjgc"))))
        nw481_[int(6)] = (default__.fm1(d_2730_v29_, p0, p0, globalState)) or (self.f7)
        nw481_[int(7)] = not(True)
        nw481_[int(8)] = not (self.f7) or (self.f7)
        nw481_[int(9)] = (self.f7 if self.f7 else self.f7)
        nw481_[int(10)] = self.f7
        nw481_[int(11)] = True
        nw481_[int(12)] = self.f7
        nw481_[int(13)] = self.f7
        nw481_[int(14)] = (_dafny.SeqWithoutIsStrInference([p1 for d_2740_i7_ in range(default__.abs(-331))])) <= (d_2723_v25_)
        nw481_[int(15)] = self.f7
        nw481_[int(16)] = (d_2738_v35_).cf83
        nw481_[int(17)] = self.f7
        nw481_[int(18)] = (d_2718_v20_) in (p1)
        nw481_[int(19)] = (d_2718_v20_) not in (_dafny.SeqWithoutIsStrInference([d_2718_v20_ for d_2741_i9_ in range(default__.abs(564))]))
        nw481_[int(20)] = True
        d_2739_v36_ = nw481_
        index448_ = default__.safeIndex(473, (d_2739_v36_).length(0))
        (d_2739_v36_)[index448_] = (self).fm8(p0, globalState)
        d_2742_v37_: _dafny.Map
        d_2742_v37_ = _dafny.Map({p0: not(self.f7)})
        index449_ = default__.safeIndex(473, (d_2739_v36_).length(0))
        rhs434_ = (d_2719_v21_).intersection((d_2719_v21_) | (d_2719_v21_))
        rhs435_ = (len(d_2720_v22_)) > (p0)
        rhs436_ = ((d_2742_v37_)[p0] if (p0) in (d_2742_v37_) else self.f7)
        rhs437_ = d_2717_v19_
        rhs438_ = self.f7
        lhs262_ = self
        lhs263_ = d_2739_v36_
        lhs264_ = default__.safeIndex(473, (d_2739_v36_).length(0))
        lhs265_ = self
        d_2719_v21_ = rhs434_
        lhs262_.f7 = rhs435_
        lhs263_[lhs264_] = rhs436_
        d_2717_v19_ = rhs437_
        lhs265_.f7 = rhs438_
        r2 = (d_2739_v36_)[default__.safeIndex(473, (d_2739_v36_).length(0))]
        d_2743_v38_: _dafny.Set
        d_2743_v38_ = _dafny.Set({d_2739_v36_, d_2739_v36_, d_2739_v36_, d_2739_v36_})
        d_2744_v39_: C2
        nw482_ = C2()
        nw482_.ctor__((d_2743_v38_).intersection(d_2743_v38_), d_2718_v20_, not(self.f7))
        d_2744_v39_ = nw482_
        d_2745_v40_: _dafny.Array
        nw483_ = _dafny.Array(_dafny.Array(None, 0), 8)
        d_2745_v40_ = nw483_
        r0 = d_2745_v40_
        r1 = len((p1) + (p1))
        r2 = (False if self.f7 else self.f7)
        d_2746_v41_: D9
        d_2746_v41_ = D9_DC25()
        pat_let_tv71_ = p0
        pat_let_tv72_ = p0
        pat_let_tv73_ = p0
        pat_let_tv74_ = p0
        pat_let_tv75_ = p0
        def lambda166_(source20_):
            if source20_.is_DC24:
                return default__.safeModuloInt(pat_let_tv71_, pat_let_tv72_)
            elif source20_.is_DC25:
                return pat_let_tv73_
            elif source20_.is_DC26:
                d_2747___mcc_h0_ = source20_.cf45
                d_2748_cf45_ = d_2747___mcc_h0_
                return (0) - ((0) - (pat_let_tv74_))
            elif True:
                d_2749___mcc_h1_ = source20_.cf44
                d_2750_cf44_ = d_2749___mcc_h1_
                return pat_let_tv75_

        r3 = lambda166_(d_2746_v41_)
        return r0, r1, r2, r3

    def m5(self, globalState):
        d_2751_v0_: C0
        nw484_ = C0()
        nw484_.ctor__(self.f7, self.f7)
        d_2751_v0_ = nw484_
        d_2752_v1_: int
        d_2752_v1_ = 593
        d_2752_v1_ = (d_2752_v1_) + (default__.safeModuloInt(d_2752_v1_, -343))
        d_2753_v2_: _dafny.Map
        d_2753_v2_ = _dafny.Map({d_2751_v0_.f13: True})
        d_2754_i0_: int
        d_2754_i0_ = 0
        with _dafny.label("19"):
            while (d_2752_v1_) > ((542) * (len(d_2753_v2_))):
                with _dafny.c_label("19"):
                    if (d_2754_i0_) >= (100):
                        raise _dafny.Break("19")
                    d_2754_i0_ = (d_2754_i0_) + (1)
                    (d_2751_v0_).f13 = self.f7
                    d_2755_v3_: _dafny.Map
                    d_2755_v3_ = (d_2753_v2_) | (d_2753_v2_)
                    d_2756_v4_: T1
                    nw485_ = C3()
                    nw485_.ctor__(default__.fm0(globalState), self.f7)
                    d_2756_v4_ = nw485_
                    d_2757_v5_: C6
                    nw486_ = C6()
                    nw486_.ctor__(not((d_2751_v0_).f12), d_2751_v0_.f13)
                    d_2757_v5_ = nw486_
                    rhs439_ = d_2755_v3_
                    rhs440_ = not((d_2751_v0_).f12)
                    rhs441_ = d_2752_v1_
                    rhs442_ = d_2756_v4_
                    rhs443_ = d_2757_v5_
                    lhs266_ = d_2751_v0_
                    d_2755_v3_ = rhs439_
                    lhs266_.f13 = rhs440_
                    d_2752_v1_ = rhs441_
                    d_2756_v4_ = rhs442_
                    d_2757_v5_ = rhs443_
                    d_2758_v6_: str
                    d_2758_v6_ = _dafny.CodePoint('i')
                    d_2759_v7_: _dafny.Map
                    d_2759_v7_ = _dafny.Map({(d_2751_v0_).f12: d_2752_v1_})
                    d_2760_v8_: _dafny.Seq
                    d_2760_v8_ = _dafny.SeqWithoutIsStrInference([d_2759_v7_])
                    d_2761_v9_: _dafny.Map
                    d_2761_v9_ = _dafny.Map({d_2758_v6_: (d_2760_v8_)[default__.safeIndex(d_2752_v1_, len(d_2760_v8_))]})
                    d_2762_v10_: _dafny.Seq
                    d_2762_v10_ = _dafny.SeqWithoutIsStrInference([d_2752_v1_, d_2752_v1_])
                    d_2761_v9_ = default__.fm57(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqkrpxyi")), d_2752_v1_, (d_2762_v10_)[default__.safeIndex(d_2752_v1_, len(d_2762_v10_))], globalState)
                    d_2763_v11_: _dafny.Seq
                    d_2763_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivk"))
                    d_2764_v12_: _dafny.Map
                    d_2764_v12_ = _dafny.Map({_dafny.Set({d_2752_v1_, d_2752_v1_}): d_2763_v11_})
                    d_2765_v13_: _dafny.Set
                    d_2765_v13_ = _dafny.Set({d_2752_v1_, d_2752_v1_, d_2752_v1_, d_2752_v1_, d_2752_v1_})
                    d_2764_v12_ = (d_2764_v12_).set((d_2765_v13_).intersection(d_2765_v13_), d_2763_v11_)
                    pass
            pass
        (self).f7 = d_2751_v0_.f13
        (d_2751_v0_).f13 = (d_2751_v0_).f12
        hi18_ = -902
        for d_2766_i1_ in range(d_2752_v1_, hi18_):
            d_2767_v14_: C1
            nw487_ = C1()
            nw487_.ctor__(d_2751_v0_.f13)
            d_2767_v14_ = nw487_
            d_2768_v15_: _dafny.Seq
            d_2768_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svdmeqnc"))
            d_2769_v16_: _dafny.Array
            nw488_ = _dafny.Array(int(0), 22)
            d_2769_v16_ = nw488_
            d_2770_v17_: _dafny.Seq
            d_2770_v17_ = _dafny.SeqWithoutIsStrInference([d_2769_v16_])
            d_2771_v18_: _dafny.Map
            d_2771_v18_ = _dafny.Map({len(d_2768_v15_): (d_2770_v17_)[default__.safeIndex(default__.fm0(globalState), len(d_2770_v17_))]})
            d_2772_v19_: _dafny.Set
            d_2773_v20_: bool
            out6_: _dafny.Set
            out7_: bool
            out6_, out7_ = (self).m6((62) + (d_2766_i1_), d_2771_v18_, d_2752_v1_, (d_2752_v1_) < (d_2766_i1_), globalState)
            d_2772_v19_ = out6_
            d_2773_v20_ = out7_
            d_2774_v21_: _dafny.Seq
            d_2774_v21_ = _dafny.SeqWithoutIsStrInference([d_2752_v1_])
            d_2775_v23_: _dafny.Array
            nw489_ = _dafny.Array(None, 11)
            nw489_[int(0)] = d_2751_v0_.f13
            nw489_[int(1)] = d_2773_v20_
            nw489_[int(2)] = False
            nw489_[int(3)] = d_2751_v0_.f13
            nw489_[int(4)] = (d_2751_v0_).f12
            nw489_[int(5)] = (d_2751_v0_).f12
            def iife220_():
                coll54_ = _dafny.Set()
                compr_54_: int
                for compr_54_ in (d_2774_v21_).Elements:
                    d_2776_v22_: int = compr_54_
                    if (d_2776_v22_) in (d_2774_v21_):
                        coll54_ = coll54_.union(_dafny.Set([default__.safeDivisionInt(d_2776_v22_, len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([110, (_dafny.MultiSet([False])).cardinality, len(_dafny.SeqWithoutIsStrInference([True, True]))])})))]))
                return _dafny.Set(coll54_)
            nw489_[int(6)] = (iife220_()
            ).isdisjoint(d_2772_v19_)
            nw489_[int(7)] = self.f7
            nw489_[int(8)] = False
            nw489_[int(9)] = False
            nw489_[int(10)] = d_2751_v0_.f13
            d_2775_v23_ = nw489_
            index450_ = default__.safeIndex(700, (d_2775_v23_).length(0))
            (d_2775_v23_)[index450_] = True
            index451_ = default__.safeIndex(700, (d_2775_v23_).length(0))
            rhs444_ = d_2752_v1_
            rhs445_ = len(((d_2768_v15_) + (d_2768_v15_)) + (d_2768_v15_))
            rhs446_ = (default__.safeDivisionInt((0) - (d_2752_v1_), d_2752_v1_)) < (d_2752_v1_)
            lhs267_ = d_2775_v23_
            lhs268_ = default__.safeIndex(700, (d_2775_v23_).length(0))
            d_2752_v1_ = rhs444_
            d_2752_v1_ = rhs445_
            lhs267_[lhs268_] = rhs446_
            if self.f7:
                d_2752_v1_ = d_2752_v1_
                d_2752_v1_ = len(d_2768_v15_)
                d_2777_v24_: _dafny.Map
                d_2777_v24_ = _dafny.Map({d_2769_v16_: d_2772_v19_})
                d_2778_v25_: D13
                d_2778_v25_ = D13_DC34(d_2777_v24_)
                d_2779_v26_: _dafny.Seq
                d_2779_v26_ = _dafny.SeqWithoutIsStrInference([d_2778_v25_, d_2778_v25_, d_2778_v25_, d_2778_v25_])
                d_2780_v27_: _dafny.Seq
                d_2780_v27_ = _dafny.SeqWithoutIsStrInference([(d_2775_v23_)[default__.safeIndex(700, (d_2775_v23_).length(0))]])
                (d_2751_v0_).f13 = (D17_DC49(default__.fm0(globalState), d_2779_v26_, d_2752_v1_, (d_2780_v27_)[default__.safeIndex(d_2766_i1_, len(d_2780_v27_))])).cf83
                d_2781_v28_: _dafny.Map
                d_2781_v28_ = _dafny.Map({(d_2775_v23_)[default__.safeIndex(700, (d_2775_v23_).length(0))]: d_2766_i1_})
                d_2781_v28_ = (d_2781_v28_).set((d_2751_v0_).f12, len((d_2767_v14_).fm31(d_2752_v1_, not((d_2751_v0_).f12), d_2768_v15_, (d_2751_v0_).f12, globalState)))
                d_2782_v29_: _dafny.Seq
                d_2782_v29_ = _dafny.SeqWithoutIsStrInference([d_2775_v23_])
                d_2775_v23_ = (d_2782_v29_)[default__.safeIndex(d_2752_v1_, len(d_2782_v29_))]
            elif True:
                d_2783_v30_: C6
                nw490_ = C6()
                nw490_.ctor__(d_2751_v0_.f13, not(False))
                d_2783_v30_ = nw490_
                index452_ = default__.safeIndex(776, (d_2769_v16_).length(0))
                (d_2769_v16_)[index452_] = d_2766_i1_
                index453_ = default__.safeIndex(776, (d_2769_v16_).length(0))
                (d_2769_v16_)[index453_] = d_2752_v1_
                d_2784_v31_: _dafny.Map
                d_2784_v31_ = _dafny.Map({(d_2751_v0_).f12: d_2775_v23_})
                d_2785_v32_: D3
                d_2785_v32_ = D3_DC10(d_2775_v23_)
                pat_let_tv76_ = d_2775_v23_
                pat_let_tv77_ = d_2775_v23_
                d_2786_v33_: _dafny.Array
                nw491_ = _dafny.Array(None, 23)
                nw491_[int(0)] = D3_DC10(((d_2784_v31_)[(d_2751_v0_).f12] if ((d_2751_v0_).f12) in (d_2784_v31_) else d_2775_v23_))
                nw491_[int(1)] = d_2785_v32_
                nw491_[int(2)] = D3_DC10(d_2775_v23_)
                def iife221_(_pat_let83_0):
                    def iife222_(d_2787_dt__update__tmp_h0_):
                        def iife223_(_pat_let84_0):
                            def iife224_(d_2788_dt__update_hcf21_h0_):
                                return D3_DC10(d_2788_dt__update_hcf21_h0_)
                            return iife224_(_pat_let84_0)
                        return iife223_(pat_let_tv76_)
                    return iife222_(_pat_let83_0)
                nw491_[int(3)] = iife221_(d_2785_v32_)
                def iife225_(_pat_let85_0):
                    def iife226_(d_2789_dt__update__tmp_h1_):
                        def iife227_(_pat_let86_0):
                            def iife228_(d_2790_dt__update_hcf21_h1_):
                                return D3_DC10(d_2790_dt__update_hcf21_h1_)
                            return iife228_(_pat_let86_0)
                        return iife227_(pat_let_tv77_)
                    return iife226_(_pat_let85_0)
                nw491_[int(4)] = iife225_(d_2785_v32_)
                nw491_[int(5)] = d_2785_v32_
                nw491_[int(6)] = d_2785_v32_
                nw491_[int(7)] = d_2785_v32_
                nw491_[int(8)] = D3_DC10(d_2775_v23_)
                nw491_[int(9)] = d_2785_v32_
                nw491_[int(10)] = d_2785_v32_
                nw491_[int(11)] = D3_DC10(d_2775_v23_)
                nw491_[int(12)] = D3_DC10(d_2775_v23_)
                nw491_[int(13)] = d_2785_v32_
                nw491_[int(14)] = (d_2785_v32_ if True else d_2785_v32_)
                nw491_[int(15)] = d_2785_v32_
                nw491_[int(16)] = d_2785_v32_
                nw491_[int(17)] = d_2785_v32_
                nw491_[int(18)] = d_2785_v32_
                nw491_[int(19)] = d_2785_v32_
                nw491_[int(20)] = d_2785_v32_
                nw491_[int(21)] = d_2785_v32_
                nw491_[int(22)] = D3_DC10(d_2775_v23_)
                d_2786_v33_ = nw491_
                index454_ = default__.safeIndex(470, (d_2786_v33_).length(0))
                (d_2786_v33_)[index454_] = d_2785_v32_
                index455_ = default__.safeIndex(470, (d_2786_v33_).length(0))
                (d_2786_v33_)[index455_] = d_2785_v32_
                (d_2751_v0_).f13 = (d_2775_v23_)[default__.safeIndex(700, (d_2775_v23_).length(0))]
                d_2791_v34_: _dafny.Map
                d_2791_v34_ = _dafny.Map({d_2751_v0_.f13: d_2753_v2_})
                d_2792_v35_: _dafny.Map
                d_2792_v35_ = _dafny.Map({(d_2769_v16_)[default__.safeIndex(776, (d_2769_v16_).length(0))]: d_2791_v34_})
                d_2793_v36_: _dafny.Map
                d_2793_v36_ = _dafny.Map({d_2766_i1_: (d_2783_v30_).f9})
                d_2792_v35_ = (d_2792_v35_).set(default__.safeDivisionInt(d_2752_v1_, (0) - (len(_dafny.Map({len(d_2768_v15_): d_2793_v36_})))), d_2791_v34_)

    def m6(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        d_2794_v0_: _dafny.Seq
        d_2794_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ck"))
        d_2794_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hosrfwfoe"))
        rhs447_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njabs"))
        rhs448_ = (d_2794_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uesfgn")))
        d_2794_v0_ = rhs447_
        d_2794_v0_ = rhs448_
        d_2795_v1_: D2
        d_2795_v1_ = D2_DC8(self.f7, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myneoghrv")))
        d_2796_v2_: _dafny.Seq
        d_2796_v2_ = _dafny.SeqWithoutIsStrInference([(d_2795_v1_).cf14])
        r1 = not((d_2796_v2_)[default__.safeIndex(len(d_2796_v2_), len(d_2796_v2_))])
        d_2797_v4_: _dafny.Array
        def lambda167_(d_2798_p0_):
            def lambda168_(d_2799_i0_):
                def iife229_():
                    coll55_ = _dafny.Set()
                    compr_55_: int
                    for compr_55_ in (_dafny.Map({d_2798_p0_: d_2798_p0_})).keys.Elements:
                        d_2800_v3_: int = compr_55_
                        if (d_2800_v3_) in (_dafny.Map({d_2798_p0_: d_2798_p0_})):
                            coll55_ = coll55_.union(_dafny.Set([(d_2800_v3_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcoa"))))]))
                    return _dafny.Set(coll55_)
                return _dafny.Map({iife229_()
                : True})

            return lambda168_

        init93_ = lambda167_(p0)
        nw492_ = _dafny.Array(None, 15)
        for i0_93_ in range(nw492_.length(0)):
            nw492_[i0_93_] = init93_(i0_93_)
        d_2797_v4_ = nw492_
        d_2801_v5_: _dafny.Set
        d_2801_v5_ = _dafny.Set({p2})
        d_2802_v6_: _dafny.Map
        d_2802_v6_ = _dafny.Map({d_2801_v5_: p3})
        index456_ = default__.safeIndex(837, (d_2797_v4_).length(0))
        (d_2797_v4_)[index456_] = (d_2802_v6_ if self.f7 else d_2802_v6_)
        d_2803_v7_: D0
        d_2803_v7_ = D0_DC0(default__.fm0(globalState))
        d_2804_v8_: _dafny.Seq
        d_2804_v8_ = _dafny.SeqWithoutIsStrInference([d_2803_v7_])
        d_2805_v9_: _dafny.Map
        d_2805_v9_ = _dafny.Map({(p3 if default__.fm1(d_2804_v8_, len(_dafny.Map({not(self.f7): self.f7})), p0, globalState) else p3): (d_2802_v6_) | (d_2802_v6_)})
        d_2806_v10_: _dafny.MultiSet
        d_2806_v10_ = _dafny.MultiSet([self.f7, self.f7])
        index457_ = default__.safeIndex(837, (d_2797_v4_).length(0))
        (d_2797_v4_)[index457_] = ((d_2805_v9_)[(self).fm8((0) - (p0), globalState)] if ((self).fm8((0) - (p0), globalState)) in (d_2805_v9_) else (d_2802_v6_).set(default__.fm2((d_2806_v10_).cardinality, self.f7, globalState), self.f7))
        d_2807_i1_: int
        d_2807_i1_ = 0
        with _dafny.label("20"):
            while True:
                with _dafny.c_label("20"):
                    if (d_2807_i1_) >= (100):
                        raise _dafny.Break("20")
                    d_2807_i1_ = (d_2807_i1_) + (1)
                    (self).f7 = self.f7
                    d_2808_v11_: int
                    d_2808_v11_ = -180
                    d_2808_v11_ = p2
                    d_2809_v12_: str
                    d_2809_v12_ = _dafny.CodePoint('u')
                    d_2809_v12_ = d_2809_v12_
                    default__.m0(215, d_2808_v11_, globalState)
                    pass
            pass
        d_2810_v13_: C6
        nw493_ = C6()
        nw493_.ctor__(p3, p3)
        d_2810_v13_ = nw493_
        r0 = (d_2801_v5_) - (d_2801_v5_)
        r1 = (d_2810_v13_).f9
        return r0, r1

    @property
    def f8(self):
        return self._f8

class C8:
    def  __init__(self):
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm5(self, globalState):
        return ((self).f6) - (387)

    def fm6(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), _dafny.CodePoint('q')])) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2811_i0_ in range(default__.abs(588))])))

    def m1(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        d_2812_v0_: bool
        d_2812_v0_ = True
        d_2813_v1_: D1
        d_2813_v1_ = D1_DC4(d_2812_v0_, d_2812_v0_)
        source21_ = (d_2813_v1_ if d_2812_v0_ else D1_DC4(d_2812_v0_, d_2812_v0_))
        if source21_.is_DC4:
            d_2814___mcc_h0_ = source21_.cf7
            d_2815___mcc_h1_ = source21_.cf8
            d_2816_cf8_ = d_2815___mcc_h1_
            d_2817_cf7_ = d_2814___mcc_h0_
            if d_2817_cf7_:
                d_2818_v2_: _dafny.Seq
                d_2818_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqrmug"))
                d_2818_v2_ = d_2818_v2_
                r3 = (0) - ((self).f6)
                d_2819_v3_: D0
                d_2819_v3_ = D0_DC0((self).f6)
                d_2820_v4_: _dafny.MultiSet
                d_2820_v4_ = _dafny.MultiSet([d_2819_v3_])
                d_2820_v4_ = (d_2820_v4_).intersection(_dafny.MultiSet([d_2819_v3_, d_2819_v3_]))
                d_2821_v5_: _dafny.MultiSet
                d_2821_v5_ = _dafny.MultiSet([(0) - ((self).f6), (self).f6])
                d_2822_v6_: _dafny.Seq
                d_2822_v6_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_2823_v7_: _dafny.Seq
                d_2823_v7_ = _dafny.SeqWithoutIsStrInference([d_2812_v0_, d_2817_cf7_])
                d_2824_v8_: _dafny.Set
                d_2824_v8_ = _dafny.Set({d_2817_cf7_})
                d_2825_v9_: _dafny.Array
                nw494_ = _dafny.Array(None, 29)
                nw494_[int(0)] = d_2821_v5_
                nw494_[int(1)] = (_dafny.MultiSet(d_2822_v6_) if True else d_2821_v5_)
                nw494_[int(2)] = d_2821_v5_
                nw494_[int(3)] = (d_2821_v5_) | (d_2821_v5_)
                nw494_[int(4)] = d_2821_v5_
                nw494_[int(5)] = _dafny.MultiSet(d_2822_v6_)
                nw494_[int(6)] = (d_2821_v5_).intersection(_dafny.MultiSet([(self).f6]))
                nw494_[int(7)] = d_2821_v5_
                nw494_[int(8)] = (d_2821_v5_).intersection(d_2821_v5_)
                nw494_[int(9)] = (d_2821_v5_) - (d_2821_v5_)
                nw494_[int(10)] = _dafny.MultiSet([(_dafny.MultiSet([d_2817_cf7_])).cardinality, (self).f6, (self).f6])
                nw494_[int(11)] = (d_2821_v5_).intersection(d_2821_v5_)
                nw494_[int(12)] = d_2821_v5_
                nw494_[int(13)] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(self).f6 for d_2826_i0_ in range(default__.abs(215))]) if d_2816_cf8_ else _dafny.SeqWithoutIsStrInference([714 for d_2827_i1_ in range(default__.abs(-510))])))
                nw494_[int(14)] = (_dafny.MultiSet([(self).f6])) | (d_2821_v5_)
                nw494_[int(15)] = _dafny.MultiSet([(self).f6, (self).f6])
                nw494_[int(16)] = default__.fm7(d_2816_cf8_, d_2812_v0_, d_2816_cf8_, globalState)
                nw494_[int(17)] = d_2821_v5_
                nw494_[int(18)] = (_dafny.MultiSet([len(d_2823_v7_), (self).f6])).set((self).f6, default__.abs((0) - (-677)))
                nw494_[int(19)] = d_2821_v5_
                nw494_[int(20)] = (d_2821_v5_) | (_dafny.MultiSet([(self).f6]))
                nw494_[int(21)] = _dafny.MultiSet(d_2822_v6_)
                nw494_[int(22)] = d_2821_v5_
                nw494_[int(23)] = default__.fm7(d_2812_v0_, d_2812_v0_, d_2817_cf7_, globalState)
                nw494_[int(24)] = d_2821_v5_
                nw494_[int(25)] = ((d_2821_v5_).set(len(d_2824_v8_), default__.abs((self).f6))).intersection(d_2821_v5_)
                nw494_[int(26)] = default__.fm7(d_2812_v0_, d_2812_v0_, d_2812_v0_, globalState)
                nw494_[int(27)] = d_2821_v5_
                nw494_[int(28)] = _dafny.MultiSet([(self).f6, (self).f6])
                d_2825_v9_ = nw494_
                d_2825_v9_ = d_2825_v9_
                d_2828_v10_: _dafny.Map
                d_2828_v10_ = _dafny.Map({d_2816_cf8_: -326})
                d_2828_v10_ = (d_2828_v10_).set(d_2812_v0_, (0) - (default__.safeDivisionInt((self).f6, (self).f6)))
            elif True:
                d_2829_v11_: C0
                nw495_ = C0()
                nw495_.ctor__(d_2817_cf7_, d_2817_cf7_)
                d_2829_v11_ = nw495_
                d_2830_v12_: str
                d_2830_v12_ = _dafny.CodePoint('w')
                d_2830_v12_ = d_2830_v12_
                d_2831_v13_: _dafny.Array
                nw496_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_2831_v13_ = nw496_
                d_2832_v14_: _dafny.Array
                nw497_ = _dafny.Array(False, 9)
                d_2832_v14_ = nw497_
                index458_ = default__.safeIndex(91, (d_2831_v13_).length(0))
                (d_2831_v13_)[index458_] = d_2832_v14_
                index459_ = default__.safeIndex(91, (d_2831_v13_).length(0))
                (d_2831_v13_)[index459_] = d_2832_v14_
                d_2833_v15_: _dafny.Seq
                d_2833_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etpudh"))
                d_2834_v16_: _dafny.MultiSet
                d_2834_v16_ = _dafny.MultiSet([len(d_2833_v15_)])
                d_2835_v17_: _dafny.MultiSet
                d_2835_v17_ = _dafny.MultiSet([(d_2829_v11_).f12])
                d_2836_v18_: _dafny.Map
                d_2836_v18_ = _dafny.Map({(d_2829_v11_).f12: (self).f6})
                d_2837_v19_: _dafny.Map
                d_2837_v19_ = _dafny.Map({(self).f6: d_2817_cf7_})
                d_2838_v20_: C6
                nw498_ = C6()
                nw498_.ctor__((d_2829_v11_).f12, d_2812_v0_)
                d_2838_v20_ = nw498_
                d_2839_v21_: D20
                d_2839_v21_ = D20_DC54(d_2838_v20_)
                d_2840_v22_: _dafny.Seq
                d_2840_v22_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
                d_2841_v23_: _dafny.Array
                nw499_ = _dafny.Array(None, 29)
                nw499_[int(0)] = 357
                nw499_[int(1)] = default__.safeModuloInt((self).f6, (self).f6)
                nw499_[int(2)] = (self).f6
                nw499_[int(3)] = default__.fm0(globalState)
                nw499_[int(4)] = (d_2834_v16_).cardinality
                nw499_[int(5)] = (self).f6
                nw499_[int(6)] = (self).f6
                nw499_[int(7)] = ((self).f6) - ((self).f6)
                nw499_[int(8)] = (self).f6
                nw499_[int(9)] = (self).f6
                nw499_[int(10)] = (-569) - ((self).f6)
                nw499_[int(11)] = (self).f6
                nw499_[int(12)] = 515
                nw499_[int(13)] = ((self).f6 if not((d_2829_v11_).f12) else (self).f6)
                nw499_[int(14)] = (self).f6
                nw499_[int(15)] = (self).f6
                nw499_[int(16)] = len(_dafny.SeqWithoutIsStrInference([d_2830_v12_ for d_2842_i2_ in range(default__.abs(-843))]))
                nw499_[int(17)] = (self).f6
                nw499_[int(18)] = (d_2835_v17_).cardinality
                nw499_[int(19)] = default__.safeDivisionInt((self).f6, len(d_2836_v18_))
                nw499_[int(20)] = (0) - ((self).f6)
                nw499_[int(21)] = (self).f6
                nw499_[int(22)] = (self).f6
                nw499_[int(23)] = (0) - (len((d_2829_v11_).fm17((self).f6, d_2812_v0_, globalState)))
                nw499_[int(24)] = (self).f6
                nw499_[int(25)] = (len(d_2837_v19_)) * (-587)
                nw499_[int(26)] = len(_dafny.Map({(d_2839_v21_).cf89: _dafny.MultiSet((d_2840_v22_).set(default__.safeIndex((self).f6, len(d_2840_v22_)), (self).f6))}))
                nw499_[int(27)] = (self).f6
                nw499_[int(28)] = (self).f6
                d_2841_v23_ = nw499_
                d_2843_v24_: _dafny.Seq
                d_2843_v24_ = _dafny.SeqWithoutIsStrInference([d_2832_v14_])
                index460_ = default__.safeIndex(434, (d_2841_v23_).length(0))
                (d_2841_v23_)[index460_] = (len(d_2843_v24_)) * ((self).f6)
                index461_ = default__.safeIndex(434, (d_2841_v23_).length(0))
                (d_2841_v23_)[index461_] = (0) - ((0) - ((0) - (((self).f6) + (-563))))
                d_2844_v25_: D0
                d_2844_v25_ = D0_DC1(((self).f6) - ((d_2841_v23_)[default__.safeIndex(434, (d_2841_v23_).length(0))]))
                d_2844_v25_ = (d_2844_v25_ if d_2812_v0_ else default__.fm58(d_2817_cf7_, d_2829_v11_.f13, globalState))
            d_2845_v26_: _dafny.Array
            nw500_ = _dafny.Array(None, 2)
            nw500_[int(0)] = False
            nw500_[int(1)] = d_2812_v0_
            d_2845_v26_ = nw500_
            d_2846_v27_: _dafny.Set
            d_2846_v27_ = _dafny.Set({d_2845_v26_})
            d_2847_v28_: str
            d_2847_v28_ = _dafny.CodePoint('q')
            d_2848_v29_: T0
            nw501_ = C2()
            nw501_.ctor__(d_2846_v27_, d_2847_v28_, d_2816_cf8_)
            d_2848_v29_ = nw501_
            d_2849_v30_: _dafny.Seq
            d_2849_v30_ = _dafny.SeqWithoutIsStrInference([d_2848_v29_, d_2848_v29_])
            d_2850_v31_: D0
            d_2850_v31_ = D0_DC0((self).f6)
            d_2851_v32_: _dafny.Seq
            d_2851_v32_ = _dafny.SeqWithoutIsStrInference([d_2850_v31_])
            d_2852_v33_: _dafny.Seq
            d_2852_v33_ = _dafny.SeqWithoutIsStrInference([d_2816_cf8_])
            d_2853_v34_: _dafny.Seq
            d_2853_v34_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_2851_v32_, len(d_2852_v33_), -168, globalState)])
            d_2854_v35_: _dafny.Map
            d_2854_v35_ = _dafny.Map({((d_2849_v30_) + ((d_2849_v30_).set(default__.safeIndex((self).f6, len(d_2849_v30_)), d_2848_v29_))).set(default__.safeIndex(default__.fm0(globalState), len((d_2849_v30_) + ((d_2849_v30_).set(default__.safeIndex((self).f6, len(d_2849_v30_)), d_2848_v29_)))), d_2848_v29_): (d_2853_v34_)[default__.safeIndex((self).f6, len(d_2853_v34_))]})
            d_2855_v36_: _dafny.Map
            d_2855_v36_ = _dafny.Map({d_2816_cf8_: True})
            d_2856_v37_: C6
            nw502_ = C6()
            nw502_.ctor__(False, default__.fm1(d_2851_v32_, len(_dafny.Map({d_2816_cf8_: d_2855_v36_})), (self).f6, globalState))
            d_2856_v37_ = nw502_
            d_2857_v38_: _dafny.MultiSet
            d_2857_v38_ = _dafny.MultiSet([d_2856_v37_, d_2856_v37_])
            d_2854_v35_ = (d_2854_v35_).set((d_2849_v30_).set(default__.safeIndex((self).f6, len(d_2849_v30_)), d_2848_v29_), (d_2857_v38_).isdisjoint(d_2857_v38_))
            d_2858_v39_: _dafny.Array
            def lambda169_(d_2859_i3_):
                return D2_DC7(_dafny.Set({(self).f6, (self).f6}))

            init94_ = lambda169_
            nw503_ = _dafny.Array(None, 21)
            for i0_94_ in range(nw503_.length(0)):
                nw503_[i0_94_] = init94_(i0_94_)
            d_2858_v39_ = nw503_
            d_2860_v40_: C4
            nw504_ = C4()
            nw504_.ctor__(d_2858_v39_, d_2816_cf8_)
            d_2860_v40_ = nw504_
            d_2861_v41_: _dafny.Set
            d_2861_v41_ = _dafny.Set({d_2847_v28_, d_2847_v28_})
            if not((d_2847_v28_) not in (d_2861_v41_)):
                d_2862_v42_: _dafny.MultiSet
                d_2862_v42_ = _dafny.MultiSet([(self).f6])
                rhs449_ = d_2812_v0_
                rhs450_ = d_2848_v29_.f7
                rhs451_ = not ((len(_dafny.SeqWithoutIsStrInference([(self).f6, 912]))) not in (d_2862_v42_)) or (not((d_2852_v33_) < (d_2852_v33_)))
                r2 = rhs449_
                r0 = rhs450_
                d_2817_cf7_ = rhs451_
                d_2863_v43_: _dafny.Array
                nw505_ = _dafny.Array(int(0), 6)
                d_2863_v43_ = nw505_
                index462_ = default__.safeIndex(541, (d_2863_v43_).length(0))
                (d_2863_v43_)[index462_] = (self).f6
                index463_ = default__.safeIndex(541, (d_2863_v43_).length(0))
                (d_2863_v43_)[index463_] = 727
                d_2864_v44_: _dafny.Set
                d_2864_v44_ = _dafny.Set({d_2863_v43_, d_2863_v43_})
                d_2865_v45_: D0
                d_2865_v45_ = D0_DC2(d_2848_v29_.f7, (d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))], d_2864_v44_, (d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))])
                d_2866_v46_: _dafny.Set
                d_2866_v46_ = _dafny.Set({d_2848_v29_.f7, (d_2865_v45_).cf2, (d_2856_v37_).f9, (d_2856_v37_).f9, d_2848_v29_.f7})
                d_2867_v47_: _dafny.Set
                d_2867_v47_ = _dafny.Set({d_2866_v46_, d_2866_v46_})
                d_2868_v48_: C2
                nw506_ = C2()
                nw506_.ctor__(d_2846_v27_, d_2847_v28_, (d_2856_v37_).f9)
                d_2868_v48_ = nw506_
                d_2869_v49_: _dafny.Map
                d_2869_v49_ = _dafny.Map({(d_2867_v47_).isdisjoint(d_2867_v47_): d_2868_v48_})
                d_2870_v50_: _dafny.Seq
                d_2870_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubha"))
                d_2871_v51_: C5
                nw507_ = C5()
                nw507_.ctor__(d_2848_v29_.f7)
                d_2871_v51_ = nw507_
                d_2872_v52_: _dafny.Map
                d_2872_v52_ = _dafny.Map({d_2850_v31_: (d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))]})
                d_2873_v53_: D2
                d_2873_v53_ = D2_DC9(default__.safeModuloInt((self).f6, (0) - ((self).f6)), (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([default__.fm1(_dafny.SeqWithoutIsStrInference([D0_DC0((self).f6)]), (self).f6, (d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))], globalState)])))), d_2870_v50_, ((d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))]) < (len(_dafny.Map({d_2871_v51_: d_2866_v46_}))), d_2872_v52_)
                d_2874_v54_: _dafny.Map
                d_2874_v54_ = _dafny.Map({(d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))]: (self).f6})
                rhs452_ = d_2869_v49_
                rhs453_ = d_2873_v53_
                rhs454_ = ((d_2874_v54_)[(d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))]] if ((d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))]) in (d_2874_v54_) else (d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))])
                d_2869_v49_ = rhs452_
                d_2873_v53_ = rhs453_
                r1 = rhs454_
                d_2875_v55_: _dafny.Map
                d_2875_v55_ = _dafny.Map({(d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))]: d_2817_cf7_})
                d_2876_v56_: _dafny.Map
                d_2876_v56_ = _dafny.Map({d_2875_v55_: d_2812_v0_})
                d_2877_v57_: _dafny.Seq
                d_2877_v57_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_2878_v58_: _dafny.Map
                d_2878_v58_ = _dafny.Map({d_2812_v0_: d_2877_v57_})
                d_2879_v59_: C7
                nw508_ = C7()
                nw508_.ctor__(d_2878_v58_, False)
                d_2879_v59_ = nw508_
                d_2880_v60_: _dafny.Map
                d_2880_v60_ = _dafny.Map({(self).f6: d_2879_v59_})
                d_2881_v61_: D12
                d_2881_v61_ = D12_DC30((self).f6, (d_2856_v37_).f9, default__.fm21(((d_2876_v56_)[d_2875_v55_] if (d_2875_v55_) in (d_2876_v56_) else (d_2856_v37_).fm8((d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))], globalState)), globalState), len(d_2880_v60_))
                d_2882_v62_: _dafny.Seq
                d_2882_v62_ = _dafny.SeqWithoutIsStrInference([d_2881_v61_, d_2881_v61_])
                d_2883_v63_: _dafny.Seq
                d_2883_v63_ = _dafny.SeqWithoutIsStrInference([d_2882_v62_])
                d_2883_v63_ = d_2883_v63_
                r2 = ((d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))]) >= ((d_2863_v43_)[default__.safeIndex(541, (d_2863_v43_).length(0))])
            elif True:
                d_2884_v64_: _dafny.Array
                nw509_ = _dafny.Array(None, 22)
                d_2884_v64_ = nw509_
                d_2885_v65_: C2
                nw510_ = C2()
                nw510_.ctor__(d_2846_v27_, d_2847_v28_, (d_2853_v34_)[default__.safeIndex((self).f6, len(d_2853_v34_))])
                d_2885_v65_ = nw510_
                index464_ = default__.safeIndex(950, (d_2884_v64_).length(0))
                (d_2884_v64_)[index464_] = d_2885_v65_
                index465_ = default__.safeIndex(950, (d_2884_v64_).length(0))
                nw511_ = C2()
                nw511_.ctor__(d_2846_v27_, d_2847_v28_, d_2816_cf8_)
                (d_2884_v64_)[index465_] = nw511_
                d_2886_v66_: _dafny.Array
                def lambda170_(d_2887_i4_):
                    return default__.safeModuloInt(d_2887_i4_, (self).f6)

                init95_ = lambda170_
                nw512_ = _dafny.Array(None, 29)
                for i0_95_ in range(nw512_.length(0)):
                    nw512_[i0_95_] = init95_(i0_95_)
                d_2886_v66_ = nw512_
                d_2888_v67_: _dafny.Seq
                d_2888_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsi"))
                index466_ = default__.safeIndex(690, (d_2886_v66_).length(0))
                (d_2886_v66_)[index466_] = (0) - (len(d_2888_v67_))
                d_2889_v68_: _dafny.Map
                d_2889_v68_ = _dafny.Map({d_2888_v67_: (d_2856_v37_).f9})
                d_2890_v69_: _dafny.MultiSet
                d_2890_v69_ = _dafny.MultiSet([(453) - ((self).f6), (0) - ((self).f6), len(d_2889_v68_)])
                index467_ = default__.safeIndex(690, (d_2886_v66_).length(0))
                rhs455_ = (self).f6
                rhs456_ = 923
                rhs457_ = d_2890_v69_
                lhs269_ = d_2886_v66_
                lhs270_ = default__.safeIndex(690, (d_2886_v66_).length(0))
                r1 = rhs455_
                lhs269_[lhs270_] = rhs456_
                d_2890_v69_ = rhs457_
                d_2891_v70_: _dafny.Map
                d_2891_v70_ = _dafny.Map({(self).f6: (self).f6})
                d_2892_v71_: D6
                d_2892_v71_ = D6_DC15(d_2847_v28_)
                d_2893_v72_: D6
                d_2893_v72_ = D6_DC18(d_2892_v71_)
                d_2894_v73_: _dafny.Seq
                d_2894_v73_ = _dafny.SeqWithoutIsStrInference([d_2893_v72_])
                d_2895_v74_: _dafny.Map
                d_2895_v74_ = _dafny.Map({d_2894_v73_: d_2817_cf7_})
                d_2891_v70_ = (d_2891_v70_).set((self).f6, len(d_2895_v74_))
                d_2896_v75_: _dafny.Map
                d_2896_v75_ = _dafny.Map({(d_2856_v37_).f9: (d_2886_v66_)[default__.safeIndex(690, (d_2886_v66_).length(0))]})
                d_2896_v75_ = d_2896_v75_
                d_2897_v76_: C4
                nw513_ = C4()
                nw513_.ctor__(d_2858_v39_, (d_2856_v37_).f9)
                d_2897_v76_ = nw513_
        elif source21_.is_DC5:
            d_2898___mcc_h2_ = source21_.cf9
            d_2899___mcc_h3_ = source21_.cf10
            d_2900___mcc_h4_ = source21_.cf11
            d_2901_cf11_ = d_2900___mcc_h4_
            d_2902_cf10_ = d_2899___mcc_h3_
            d_2903_cf9_ = d_2898___mcc_h2_
            d_2904_v77_: T1
            nw514_ = C6()
            nw514_.ctor__(d_2812_v0_, d_2812_v0_)
            d_2904_v77_ = nw514_
            d_2905_v78_: _dafny.Set
            d_2905_v78_ = _dafny.Set({d_2904_v77_, d_2904_v77_})
            d_2906_v79_: _dafny.Set
            d_2906_v79_ = _dafny.Set({len(d_2905_v78_)})
            d_2907_v80_: _dafny.Seq
            d_2907_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sraop"))
            d_2908_v81_: D2
            d_2908_v81_ = D2_DC7(d_2906_v79_)
            d_2909_v82_: _dafny.Array
            nw515_ = _dafny.Array(None, 19)
            nw515_[int(0)] = d_2906_v79_
            nw515_[int(1)] = (d_2906_v79_) | (d_2906_v79_)
            nw515_[int(2)] = d_2906_v79_
            nw515_[int(3)] = d_2906_v79_
            nw515_[int(4)] = (d_2906_v79_) | (_dafny.Set({len(d_2907_v80_)}))
            nw515_[int(5)] = d_2906_v79_
            nw515_[int(6)] = d_2906_v79_
            nw515_[int(7)] = d_2906_v79_
            nw515_[int(8)] = d_2906_v79_
            nw515_[int(9)] = (d_2908_v81_).cf13
            nw515_[int(10)] = d_2906_v79_
            nw515_[int(11)] = d_2906_v79_
            nw515_[int(12)] = _dafny.Set({(self).f6})
            nw515_[int(13)] = (d_2906_v79_) - (d_2906_v79_)
            nw515_[int(14)] = _dafny.Set({(self).f6, -181})
            nw515_[int(15)] = default__.fm2((self).f6, d_2903_cf9_, globalState)
            nw515_[int(16)] = d_2906_v79_
            nw515_[int(17)] = (default__.fm2(len(d_2907_v80_), False, globalState)) | (d_2906_v79_)
            nw515_[int(18)] = d_2906_v79_
            d_2909_v82_ = nw515_
            d_2909_v82_ = d_2909_v82_
            d_2910_v83_: str
            d_2910_v83_ = _dafny.CodePoint('f')
            d_2911_v84_: _dafny.Set
            d_2911_v84_ = _dafny.Set({d_2910_v83_})
            d_2912_v85_: D0
            d_2912_v85_ = D0_DC0(d_2901_cf11_)
            d_2913_v86_: _dafny.Map
            d_2913_v86_ = _dafny.Map({len((self).fm6(d_2912_v85_, globalState)): d_2907_v80_})
            d_2914_v87_: _dafny.MultiSet
            d_2914_v87_ = _dafny.MultiSet([(0) - (d_2901_cf11_)])
            d_2915_v88_: _dafny.Array
            nw516_ = _dafny.Array(None, 26)
            nw516_[int(0)] = -691
            nw516_[int(1)] = (self).f6
            nw516_[int(2)] = default__.safeModuloInt(d_2901_cf11_, (self).f6)
            nw516_[int(3)] = default__.safeDivisionInt((self).f6, -107)
            nw516_[int(4)] = len(((d_2907_v80_).set(default__.safeIndex((self).f6, len(d_2907_v80_)), d_2910_v83_)) + (d_2907_v80_))
            nw516_[int(5)] = d_2901_cf11_
            nw516_[int(6)] = (76) - ((self).f6)
            nw516_[int(7)] = d_2901_cf11_
            nw516_[int(8)] = (self).f6
            nw516_[int(9)] = len(d_2906_v79_)
            nw516_[int(10)] = len((d_2911_v84_) | (d_2911_v84_))
            nw516_[int(11)] = d_2901_cf11_
            nw516_[int(12)] = d_2901_cf11_
            nw516_[int(13)] = (0) - (((_dafny.MultiSet([d_2907_v80_, _dafny.SeqWithoutIsStrInference([d_2910_v83_, d_2910_v83_])])).set(d_2907_v80_, default__.abs((self).f6))).cardinality)
            nw516_[int(14)] = d_2901_cf11_
            nw516_[int(15)] = d_2901_cf11_
            nw516_[int(16)] = len((d_2907_v80_) + (((d_2913_v86_)[(self).f6] if ((self).f6) in (d_2913_v86_) else d_2907_v80_)))
            nw516_[int(17)] = (len(default__.fm21(d_2903_cf9_, globalState))) + ((self).f6)
            nw516_[int(18)] = len(d_2907_v80_)
            nw516_[int(19)] = (self).f6
            nw516_[int(20)] = (self).f6
            nw516_[int(21)] = (self).f6
            nw516_[int(22)] = default__.safeModuloInt(len(d_2907_v80_), (0) - (d_2901_cf11_))
            nw516_[int(23)] = ((d_2914_v87_).intersection(d_2914_v87_)).cardinality
            nw516_[int(24)] = 957
            nw516_[int(25)] = d_2901_cf11_
            d_2915_v88_ = nw516_
            index468_ = default__.safeIndex(537, (d_2915_v88_).length(0))
            (d_2915_v88_)[index468_] = (0) - ((self).f6)
            index469_ = default__.safeIndex(537, (d_2915_v88_).length(0))
            (d_2915_v88_)[index469_] = (self).f6
            d_2916_v89_: _dafny.Seq
            d_2916_v89_ = _dafny.SeqWithoutIsStrInference([not(not(d_2812_v0_)), d_2903_cf9_])
            d_2916_v89_ = d_2916_v89_
            d_2917_v90_: _dafny.Array
            def lambda171_(d_2918_v81_):
                def lambda172_(d_2919_i5_):
                    return d_2918_v81_

                return lambda172_

            init96_ = lambda171_(d_2908_v81_)
            nw517_ = _dafny.Array(None, 8)
            for i0_96_ in range(nw517_.length(0)):
                nw517_[i0_96_] = init96_(i0_96_)
            d_2917_v90_ = nw517_
            d_2920_v91_: C4
            nw518_ = C4()
            nw518_.ctor__(d_2917_v90_, d_2812_v0_)
            d_2920_v91_ = nw518_
            d_2921_v92_: _dafny.Set
            d_2921_v92_ = _dafny.Set({d_2920_v91_})
            d_2922_v93_: _dafny.Map
            d_2922_v93_ = _dafny.Map({d_2812_v0_: d_2921_v92_})
            d_2923_v94_: _dafny.Map
            d_2923_v94_ = _dafny.Map({(self).f6: (self).f6})
            d_2924_v95_: _dafny.MultiSet
            d_2924_v95_ = _dafny.MultiSet([d_2923_v94_, d_2923_v94_])
            d_2925_v96_: _dafny.Seq
            d_2925_v96_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({-743: (d_2915_v88_)[default__.safeIndex(537, (d_2915_v88_).length(0))]})])
            d_2926_v97_: _dafny.Seq
            d_2926_v97_ = _dafny.SeqWithoutIsStrInference([d_2914_v87_, (d_2914_v87_).set(len(d_2907_v80_), default__.abs(len(_dafny.SeqWithoutIsStrInference([not(d_2903_cf9_)]))))])
            d_2922_v93_ = (d_2922_v93_).set(not(((_dafny.MultiSet(d_2925_v96_)).set(_dafny.Map({len(_dafny.Map({d_2915_v88_: d_2904_v77_.f7})): (d_2915_v88_)[default__.safeIndex(537, (d_2915_v88_).length(0))]}), default__.abs(((d_2926_v97_)[default__.safeIndex(271, len(d_2926_v97_))]).cardinality))).ispropersubset(d_2924_v95_)), _dafny.Set({d_2920_v91_, d_2920_v91_}))
        elif source21_.is_DC3:
            d_2927___mcc_h5_ = source21_.cf6
            d_2928_cf6_ = d_2927___mcc_h5_
            d_2929_v98_: _dafny.Array
            nw519_ = _dafny.Array(None, 29)
            d_2929_v98_ = nw519_
            d_2930_v99_: C6
            nw520_ = C6()
            nw520_.ctor__(d_2812_v0_, d_2812_v0_)
            d_2930_v99_ = nw520_
            index470_ = default__.safeIndex(741, (d_2929_v98_).length(0))
            (d_2929_v98_)[index470_] = d_2930_v99_
            index471_ = default__.safeIndex(741, (d_2929_v98_).length(0))
            (d_2929_v98_)[index471_] = d_2930_v99_
            d_2931_v100_: C1
            nw521_ = C1()
            nw521_.ctor__(d_2812_v0_)
            d_2931_v100_ = nw521_
            d_2932_v101_: _dafny.MultiSet
            d_2932_v101_ = _dafny.MultiSet([d_2931_v100_])
            d_2933_v102_: C6
            nw522_ = C6()
            nw522_.ctor__(not((d_2812_v0_ if (d_2930_v99_).f9 else d_2812_v0_)), (d_2932_v101_).ispropersubset(_dafny.MultiSet([d_2931_v100_, d_2931_v100_])))
            d_2933_v102_ = nw522_
            d_2934_v103_: C5
            nw523_ = C5()
            nw523_.ctor__((d_2930_v99_).f9)
            d_2934_v103_ = nw523_
            r2 = (d_2933_v102_).fm8((self).f6, globalState)
        elif True:
            d_2935___mcc_h6_ = source21_.cf12
            d_2936_cf12_ = d_2935___mcc_h6_
            d_2937_v104_: _dafny.Map
            d_2937_v104_ = _dafny.Map({d_2812_v0_: default__.fm21(not(False), globalState)})
            d_2938_v105_: T0
            nw524_ = C7()
            nw524_.ctor__(d_2937_v104_, d_2812_v0_)
            d_2938_v105_ = nw524_
            d_2939_v106_: _dafny.Map
            d_2939_v106_ = _dafny.Map({(self).f6: d_2938_v105_})
            d_2940_v107_: _dafny.Map
            d_2940_v107_ = _dafny.Map({d_2938_v105_.f7: (self).f6})
            d_2941_v108_: _dafny.Seq
            d_2941_v108_ = _dafny.SeqWithoutIsStrInference([(d_2939_v106_) | ((d_2939_v106_).set(((d_2940_v107_)[d_2938_v105_.f7] if (d_2938_v105_.f7) in (d_2940_v107_) else (self).f6), d_2938_v105_)), d_2939_v106_, (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2938_v105_.f7])): d_2938_v105_})) | (d_2939_v106_), d_2939_v106_])
            d_2941_v108_ = d_2941_v108_
            d_2942_v109_: _dafny.Array
            nw525_ = _dafny.Array(D13.default()(), 25)
            d_2942_v109_ = nw525_
            d_2943_v110_: _dafny.Array
            nw526_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_2943_v110_ = nw526_
            d_2944_v111_: D0
            d_2944_v111_ = D0_DC0((self).f6)
            d_2945_v112_: _dafny.Seq
            d_2945_v112_ = _dafny.SeqWithoutIsStrInference([d_2944_v111_])
            d_2946_v113_: D12
            d_2946_v113_ = D12_DC33(D12_DC30(171, default__.fm1(d_2945_v112_, (self).f6, (self).f6, globalState), _dafny.SeqWithoutIsStrInference([(self).f6 for d_2947_i6_ in range(default__.abs(0))]), (self).f6))
            d_2948_v114_: D12
            d_2948_v114_ = D12_DC32(False)
            pat_let_tv78_ = d_2948_v114_
            d_2949_v115_: _dafny.Array
            nw527_ = _dafny.Array(None, 15)
            nw527_[int(0)] = d_2946_v113_
            nw527_[int(1)] = d_2946_v113_
            nw527_[int(2)] = d_2946_v113_
            nw527_[int(3)] = D12_DC33(d_2948_v114_)
            nw527_[int(4)] = d_2946_v113_
            nw527_[int(5)] = d_2946_v113_
            nw527_[int(6)] = d_2946_v113_
            nw527_[int(7)] = d_2946_v113_
            def iife230_(_pat_let87_0):
                def iife231_(d_2950_dt__update__tmp_h0_):
                    def iife232_(_pat_let88_0):
                        def iife233_(d_2951_dt__update_hcf56_h0_):
                            return D12_DC33(d_2951_dt__update_hcf56_h0_)
                        return iife233_(_pat_let88_0)
                    return iife232_(pat_let_tv78_)
                return iife231_(_pat_let87_0)
            nw527_[int(8)] = iife230_(d_2946_v113_)
            nw527_[int(9)] = d_2946_v113_
            nw527_[int(10)] = D12_DC33(d_2948_v114_)
            nw527_[int(11)] = d_2946_v113_
            nw527_[int(12)] = D12_DC33(d_2948_v114_)
            nw527_[int(13)] = d_2946_v113_
            nw527_[int(14)] = d_2946_v113_
            d_2949_v115_ = nw527_
            index472_ = default__.safeIndex(247, (d_2943_v110_).length(0))
            (d_2943_v110_)[index472_] = d_2949_v115_
            d_2952_v116_: D21
            d_2952_v116_ = D21_DC57(d_2949_v115_)
            index473_ = default__.safeIndex(247, (d_2943_v110_).length(0))
            rhs458_ = False
            rhs459_ = d_2942_v109_
            rhs460_ = (d_2952_v116_).cf98
            lhs271_ = d_2943_v110_
            lhs272_ = default__.safeIndex(247, (d_2943_v110_).length(0))
            d_2812_v0_ = rhs458_
            d_2942_v109_ = rhs459_
            lhs271_[lhs272_] = rhs460_
            r3 = (self).f6
            d_2953_v117_: _dafny.Seq
            d_2953_v117_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ao"))
            d_2954_v118_: D2
            d_2954_v118_ = D2_DC8(((self).f6) < ((self).f6), d_2953_v117_)
            source22_ = d_2954_v118_
            if source22_.is_DC8:
                d_2955___mcc_h7_ = source22_.cf14
                d_2956___mcc_h8_ = source22_.cf15
                d_2957_cf15_ = d_2956___mcc_h8_
                d_2958_cf14_ = d_2955___mcc_h7_
                d_2959_v119_: _dafny.Set
                d_2959_v119_ = _dafny.Set({(self).f6, (self).f6})
                d_2960_v120_: D2
                d_2960_v120_ = D2_DC7(d_2959_v119_)
                d_2961_v121_: _dafny.Array
                nw528_ = _dafny.Array(None, 18)
                nw528_[int(0)] = d_2960_v120_
                nw528_[int(1)] = d_2960_v120_
                nw528_[int(2)] = d_2960_v120_
                nw528_[int(3)] = d_2960_v120_
                nw528_[int(4)] = d_2960_v120_
                nw528_[int(5)] = d_2960_v120_
                nw528_[int(6)] = d_2960_v120_
                nw528_[int(7)] = d_2960_v120_
                nw528_[int(8)] = D2_DC7(d_2959_v119_)
                nw528_[int(9)] = d_2960_v120_
                nw528_[int(10)] = d_2960_v120_
                nw528_[int(11)] = D2_DC7(d_2959_v119_)
                nw528_[int(12)] = d_2960_v120_
                nw528_[int(13)] = d_2960_v120_
                nw528_[int(14)] = d_2960_v120_
                nw528_[int(15)] = d_2960_v120_
                nw528_[int(16)] = d_2960_v120_
                nw528_[int(17)] = d_2960_v120_
                d_2961_v121_ = nw528_
                d_2962_v122_: D22
                d_2962_v122_ = D22_DC61(d_2961_v121_)
                pat_let_tv79_ = d_2961_v121_
                d_2963_v123_: _dafny.Seq
                def iife234_(_pat_let89_0):
                    def iife235_(d_2964_dt__update__tmp_h1_):
                        def iife236_(_pat_let90_0):
                            def iife237_(d_2965_dt__update_hcf105_h0_):
                                return D22_DC61(d_2965_dt__update_hcf105_h0_)
                            return iife237_(_pat_let90_0)
                        return iife236_(pat_let_tv79_)
                    return iife235_(_pat_let89_0)
                d_2963_v123_ = _dafny.SeqWithoutIsStrInference([(iife234_(d_2962_v122_)).cf105])
                d_2966_v124_: C4
                nw529_ = C4()
                nw529_.ctor__((d_2963_v123_)[default__.safeIndex((self).f6, len(d_2963_v123_))], not (False) or (d_2812_v0_))
                d_2966_v124_ = nw529_
                (d_2938_v105_).f7 = d_2812_v0_
                d_2958_cf14_ = (not(d_2812_v0_)) or (d_2958_cf14_)
                d_2967_v125_: D15
                d_2967_v125_ = D15_DC43(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnka"))))
                d_2968_v126_: D15
                d_2968_v126_ = D15_DC44(d_2967_v125_)
                d_2969_v127_: _dafny.Seq
                d_2969_v127_ = _dafny.SeqWithoutIsStrInference([d_2968_v126_, d_2968_v126_, d_2968_v126_, d_2968_v126_, d_2968_v126_])
                d_2969_v127_ = default__.fm59(globalState)
            elif source22_.is_DC9:
                d_2970___mcc_h9_ = source22_.cf16
                d_2971___mcc_h10_ = source22_.cf17
                d_2972___mcc_h11_ = source22_.cf18
                d_2973___mcc_h12_ = source22_.cf19
                d_2974___mcc_h13_ = source22_.cf20
                d_2975_cf20_ = d_2974___mcc_h13_
                d_2976_cf19_ = d_2973___mcc_h12_
                d_2977_cf18_ = d_2972___mcc_h11_
                d_2978_cf17_ = d_2971___mcc_h10_
                d_2979_cf16_ = d_2970___mcc_h9_
                d_2978_cf17_ = (d_2979_cf16_) + ((self).f6)
                d_2980_v128_: _dafny.Map
                d_2980_v128_ = _dafny.Map({d_2976_cf19_: d_2952_v116_})
                d_2981_v129_: _dafny.Seq
                d_2981_v129_ = _dafny.SeqWithoutIsStrInference([d_2980_v128_])
                d_2981_v129_ = (_dafny.SeqWithoutIsStrInference([(d_2980_v128_).set(d_2812_v0_, d_2952_v116_), d_2980_v128_, d_2980_v128_]) if d_2812_v0_ else d_2981_v129_)
                d_2982_v130_: str
                d_2982_v130_ = _dafny.CodePoint('l')
                d_2982_v130_ = d_2982_v130_
                d_2983_v131_: _dafny.Map
                d_2983_v131_ = _dafny.Map({d_2979_cf16_: d_2938_v105_.f7})
                d_2984_v132_: D23
                d_2984_v132_ = D23_DC64(_dafny.Map({d_2979_cf16_: d_2982_v130_}))
                d_2985_v133_: _dafny.Map
                d_2985_v133_ = _dafny.Map({(self).f6: d_2982_v130_})
                d_2986_v135_: _dafny.Seq
                d_2986_v135_ = _dafny.SeqWithoutIsStrInference([d_2978_cf17_])
                def iife238_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in (d_2986_v135_).Elements:
                        d_2987_v134_: int = compr_56_
                        if (d_2987_v134_) in (d_2986_v135_):
                            coll56_[default__.safeDivisionInt(d_2987_v134_, d_2979_cf16_)] = not(d_2938_v105_.f7)
                    return _dafny.Map(coll56_)
                rhs461_ = (d_2938_v105_.f7 if d_2976_cf19_ else ((d_2984_v132_).cf109) != (d_2985_v133_))
                rhs462_ = d_2978_cf17_
                rhs463_ = (iife238_()
                ) | ((d_2983_v131_) | (default__.fm45(d_2978_cf17_, (self).f6, globalState)))
                r0 = rhs461_
                d_2979_cf16_ = rhs462_
                d_2983_v131_ = rhs463_
            elif True:
                d_2988___mcc_h14_ = source22_.cf13
                d_2989_cf13_ = d_2988___mcc_h14_
                d_2990_v136_: _dafny.Set
                d_2990_v136_ = _dafny.Set({d_2812_v0_})
                r3 = len((default__.fm60(globalState)) | (_dafny.Map({d_2990_v136_: 315})))
                d_2991_v137_: _dafny.Array
                nw530_ = _dafny.Array(False, 28)
                d_2991_v137_ = nw530_
                d_2992_v138_: _dafny.Array
                nw531_ = _dafny.Array(None, 9)
                nw531_[int(0)] = d_2991_v137_
                nw531_[int(1)] = d_2991_v137_
                nw531_[int(2)] = d_2991_v137_
                nw531_[int(3)] = d_2991_v137_
                nw531_[int(4)] = d_2991_v137_
                nw531_[int(5)] = d_2991_v137_
                nw531_[int(6)] = d_2991_v137_
                nw531_[int(7)] = d_2991_v137_
                nw531_[int(8)] = d_2991_v137_
                d_2992_v138_ = nw531_
                index474_ = default__.safeIndex(849, (d_2992_v138_).length(0))
                (d_2992_v138_)[index474_] = d_2991_v137_
                d_2993_v139_: _dafny.Seq
                d_2993_v139_ = _dafny.SeqWithoutIsStrInference([d_2938_v105_.f7, not(not(d_2812_v0_)), True])
                index475_ = default__.safeIndex(695, (d_2991_v137_).length(0))
                (d_2991_v137_)[index475_] = (d_2812_v0_) not in (d_2993_v139_)
                index476_ = default__.safeIndex(121, (d_2991_v137_).length(0))
                (d_2991_v137_)[index476_] = d_2938_v105_.f7
                d_2994_v140_: _dafny.Array
                def lambda173_(d_2995_v117_):
                    def lambda174_(d_2996_i7_):
                        return default__.safeDivisionInt(d_2996_i7_, len(d_2995_v117_))

                    return lambda174_

                init97_ = lambda173_(d_2953_v117_)
                nw532_ = _dafny.Array(None, 29)
                for i0_97_ in range(nw532_.length(0)):
                    nw532_[i0_97_] = init97_(i0_97_)
                d_2994_v140_ = nw532_
                index477_ = default__.safeIndex(728, (d_2994_v140_).length(0))
                (d_2994_v140_)[index477_] = 13
                d_2997_v141_: _dafny.Map
                d_2997_v141_ = _dafny.Map({(self).f6: d_2938_v105_.f7})
                index478_ = default__.safeIndex(849, (d_2992_v138_).length(0))
                index479_ = default__.safeIndex(695, (d_2991_v137_).length(0))
                index480_ = default__.safeIndex(121, (d_2991_v137_).length(0))
                index481_ = default__.safeIndex(728, (d_2994_v140_).length(0))
                rhs464_ = d_2991_v137_
                rhs465_ = (self).f6
                rhs466_ = d_2812_v0_
                rhs467_ = ((d_2997_v141_)[(self).f6] if ((self).f6) in (d_2997_v141_) else d_2812_v0_)
                rhs468_ = (0) - ((self).f6)
                lhs273_ = d_2992_v138_
                lhs274_ = default__.safeIndex(849, (d_2992_v138_).length(0))
                lhs275_ = d_2991_v137_
                lhs276_ = default__.safeIndex(695, (d_2991_v137_).length(0))
                lhs277_ = d_2991_v137_
                lhs278_ = default__.safeIndex(121, (d_2991_v137_).length(0))
                lhs279_ = d_2994_v140_
                lhs280_ = default__.safeIndex(728, (d_2994_v140_).length(0))
                lhs273_[lhs274_] = rhs464_
                r1 = rhs465_
                lhs275_[lhs276_] = rhs466_
                lhs277_[lhs278_] = rhs467_
                lhs279_[lhs280_] = rhs468_
                d_2998_v142_: C1
                nw533_ = C1()
                nw533_.ctor__(((_dafny.Map({d_2938_v105_.f7: d_2938_v105_.f7}))) != (_dafny.Map({d_2938_v105_.f7: d_2938_v105_.f7})))
                d_2998_v142_ = nw533_
                d_2998_v142_ = d_2998_v142_
        d_2999_v143_: _dafny.Seq
        d_2999_v143_ = _dafny.SeqWithoutIsStrInference([(self).f6, -167])
        d_3000_v144_: _dafny.Seq
        d_3000_v144_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cosrqbc"))
        d_3001_v145_: _dafny.Map
        d_3001_v145_ = _dafny.Map({(self).f6: (self).f6})
        d_3002_v146_: _dafny.MultiSet
        d_3002_v146_ = _dafny.MultiSet([d_2812_v0_, d_2812_v0_, d_2812_v0_])
        d_3003_v147_: _dafny.Array
        nw534_ = _dafny.Array(False, 23)
        d_3003_v147_ = nw534_
        d_3004_v148_: _dafny.Set
        d_3004_v148_ = _dafny.Set({d_3003_v147_, d_3003_v147_, d_3003_v147_, d_3003_v147_})
        d_3005_v149_: str
        d_3005_v149_ = _dafny.CodePoint('l')
        d_3006_v150_: T0
        nw535_ = C2()
        nw535_.ctor__(d_3004_v148_, d_3005_v149_, d_2812_v0_)
        d_3006_v150_ = nw535_
        d_3007_v151_: _dafny.MultiSet
        d_3007_v151_ = _dafny.MultiSet([d_3006_v150_, d_3006_v150_])
        d_3008_v152_: D0
        d_3008_v152_ = D0_DC0(default__.fm0(globalState))
        d_3009_v153_: _dafny.Array
        nw536_ = _dafny.Array(None, 24)
        nw536_[int(0)] = len(d_2999_v143_)
        nw536_[int(1)] = (self).f6
        nw536_[int(2)] = (self).fm5(globalState)
        nw536_[int(3)] = (self).f6
        nw536_[int(4)] = (self).f6
        nw536_[int(5)] = (len(d_3000_v144_)) * (((d_3001_v145_)[(self).f6] if ((self).f6) in (d_3001_v145_) else (self).f6))
        nw536_[int(6)] = ((0) - ((self).f6)) + ((self).f6)
        nw536_[int(7)] = (self).f6
        nw536_[int(8)] = (d_3002_v146_).cardinality
        nw536_[int(9)] = (self).f6
        nw536_[int(10)] = default__.safeModuloInt((0) - ((self).f6), (self).f6)
        nw536_[int(11)] = ((self).f6) * ((self).f6)
        nw536_[int(12)] = (self).f6
        nw536_[int(13)] = (self).f6
        nw536_[int(14)] = (self).f6
        nw536_[int(15)] = (self).f6
        nw536_[int(16)] = ((d_3007_v151_)[d_3006_v150_] if (d_3006_v150_) in (d_3007_v151_) else (self).f6)
        nw536_[int(17)] = (self).f6
        nw536_[int(18)] = (self).f6
        nw536_[int(19)] = (self).f6
        nw536_[int(20)] = 176
        nw536_[int(21)] = (self).f6
        def iife239_(_pat_let91_0):
            def iife240_(d_3010_dt__update__tmp_h2_):
                def iife241_(_pat_let92_0):
                    def iife242_(d_3011_dt__update_hcf0_h0_):
                        return D0_DC0(d_3011_dt__update_hcf0_h0_)
                    return iife242_(_pat_let92_0)
                return iife241_((self).f6)
            return iife240_(_pat_let91_0)
        nw536_[int(22)] = len(((self).fm6(iife239_(d_3008_v152_), globalState)) + (d_3000_v144_))
        nw536_[int(23)] = -750
        d_3009_v153_ = nw536_
        d_3012_v154_: _dafny.Map
        d_3012_v154_ = _dafny.Map({d_3009_v153_: (0) - ((self).f6)})
        index482_ = default__.safeIndex(456, (d_3009_v153_).length(0))
        (d_3009_v153_)[index482_] = default__.safeDivisionInt(((d_3012_v154_)[d_3009_v153_] if (d_3009_v153_) in (d_3012_v154_) else 666), (self).f6)
        index483_ = default__.safeIndex(456, (d_3009_v153_).length(0))
        (d_3009_v153_)[index483_] = default__.safeDivisionInt((self).f6, (self).f6)
        r3 = 435
        d_3013_v155_: D2
        d_3013_v155_ = D2_DC8(d_2812_v0_, d_3000_v144_)
        hi19_ = -443
        for d_3014_i8_ in range(len(_dafny.Map({d_3013_v155_: d_3000_v144_})), hi19_):
            d_3015_v156_: _dafny.Seq
            d_3015_v156_ = _dafny.SeqWithoutIsStrInference([d_2812_v0_, d_2812_v0_])
            d_3016_v157_: D13
            d_3016_v157_ = D13_DC36(d_3005_v149_, d_3015_v156_, d_3006_v150_.f7)
            d_3017_v158_: D13
            d_3017_v158_ = D13_DC37(d_3016_v157_)
            source23_ = d_3017_v158_
            if source23_.is_DC35:
                default__.m0((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], default__.safeDivisionInt((self).f6, (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]), globalState)
                d_3018_v159_: _dafny.Seq
                d_3018_v159_ = _dafny.SeqWithoutIsStrInference([d_3008_v152_, d_3008_v152_])
                pat_let_tv80_ = d_3009_v153_
                pat_let_tv81_ = d_3009_v153_
                def iife243_(_pat_let93_0):
                    def iife244_(d_3019_dt__update__tmp_h3_):
                        def iife245_(_pat_let94_0):
                            def iife246_(d_3020_dt__update_hcf0_h1_):
                                return D0_DC0(d_3020_dt__update_hcf0_h1_)
                            return iife246_(_pat_let94_0)
                        return iife245_((pat_let_tv81_)[default__.safeIndex(456, (pat_let_tv80_).length(0))])
                    return iife244_(_pat_let93_0)
                r0 = (d_2812_v0_ if default__.fm1((d_3018_v159_).set(default__.safeIndex(d_3014_i8_, len(d_3018_v159_)), iife243_(d_3008_v152_)), len(d_3015_v156_), (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], globalState) else False)
                d_3021_v160_: _dafny.Map
                d_3021_v160_ = _dafny.Map({(d_2999_v143_) + (d_2999_v143_): d_3006_v150_})
                d_3021_v160_ = (d_3021_v160_).set(d_2999_v143_, d_3006_v150_)
                d_3022_v161_: _dafny.Map
                d_3022_v161_ = _dafny.Map({205: d_3005_v149_})
                d_3023_v162_: _dafny.Set
                d_3023_v162_ = _dafny.Set({((d_3022_v161_)[d_3014_i8_] if (d_3014_i8_) in (d_3022_v161_) else _dafny.CodePoint('a')), d_3005_v149_})
                d_3024_v164_: _dafny.Seq
                def iife247_():
                    coll57_ = _dafny.Set()
                    compr_57_: _dafny.Seq
                    for compr_57_ in (_dafny.SeqWithoutIsStrInference([d_3000_v144_ for d_3025_i9_ in range(default__.abs(-994))])).Elements:
                        d_3026_v163_: _dafny.Seq = compr_57_
                        if (d_3026_v163_) in (_dafny.SeqWithoutIsStrInference([d_3000_v144_ for d_3025_i9_ in range(default__.abs(-994))])):
                            coll57_ = coll57_.union(_dafny.Set([d_3026_v163_]))
                    return _dafny.Set(coll57_)
                d_3024_v164_ = _dafny.SeqWithoutIsStrInference([iife247_()
                ])
                d_3027_v165_: _dafny.Map
                d_3027_v165_ = _dafny.Map({d_3006_v150_.f7: d_3006_v150_.f7})
                d_3015_v156_ = (d_3015_v156_) + (default__.fm50(d_3023_v162_, (d_3024_v164_)[default__.safeIndex(d_3014_i8_, len(d_3024_v164_))], len(d_3027_v165_), globalState))
            elif source23_.is_DC36:
                d_3028___mcc_h15_ = source23_.cf58
                d_3029___mcc_h16_ = source23_.cf59
                d_3030___mcc_h17_ = source23_.cf60
                d_3031_cf60_ = d_3030___mcc_h17_
                d_3032_cf59_ = d_3029___mcc_h16_
                d_3033_cf58_ = d_3028___mcc_h15_
                index484_ = default__.safeIndex(746, (d_3003_v147_).length(0))
                (d_3003_v147_)[index484_] = d_2812_v0_
                index485_ = default__.safeIndex(746, (d_3003_v147_).length(0))
                (d_3003_v147_)[index485_] = d_2812_v0_
                r3 = default__.safeDivisionInt((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])
                index486_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                (d_3009_v153_)[index486_] = ((self).f6) + ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])
                r3 = d_3014_i8_
            elif source23_.is_DC34:
                d_3034___mcc_h18_ = source23_.cf57
                d_3035_cf57_ = d_3034___mcc_h18_
                d_3036_v166_: T1
                nw537_ = C6()
                nw537_.ctor__(d_3006_v150_.f7, d_3006_v150_.f7)
                d_3036_v166_ = nw537_
                d_3037_v167_: _dafny.Array
                nw538_ = _dafny.Array(None, 12)
                nw538_[int(0)] = d_3036_v166_
                nw538_[int(1)] = d_3036_v166_
                nw538_[int(2)] = d_3036_v166_
                nw538_[int(3)] = d_3036_v166_
                nw538_[int(4)] = d_3036_v166_
                nw538_[int(5)] = d_3036_v166_
                nw538_[int(6)] = d_3036_v166_
                nw538_[int(7)] = d_3036_v166_
                nw538_[int(8)] = d_3036_v166_
                nw538_[int(9)] = d_3036_v166_
                nw538_[int(10)] = d_3036_v166_
                nw538_[int(11)] = d_3036_v166_
                d_3037_v167_ = nw538_
                index487_ = default__.safeIndex(969, (d_3037_v167_).length(0))
                (d_3037_v167_)[index487_] = d_3036_v166_
                d_3038_v168_: _dafny.Array
                nw539_ = _dafny.Array(None, 18)
                nw539_[int(0)] = d_3000_v144_
                nw539_[int(1)] = d_3000_v144_
                nw539_[int(2)] = d_3000_v144_
                nw539_[int(3)] = d_3000_v144_
                nw539_[int(4)] = d_3000_v144_
                nw539_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oviiryq"))
                nw539_[int(6)] = d_3000_v144_
                nw539_[int(7)] = d_3000_v144_
                nw539_[int(8)] = (_dafny.SeqWithoutIsStrInference([d_3005_v149_ for d_3039_i10_ in range(default__.abs(388))]) if not(False) else d_3000_v144_)
                nw539_[int(9)] = d_3000_v144_
                nw539_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwnc"))
                nw539_[int(11)] = d_3000_v144_
                nw539_[int(12)] = d_3000_v144_
                nw539_[int(13)] = d_3000_v144_
                nw539_[int(14)] = (d_3000_v144_).set(default__.safeIndex(d_3014_i8_, len(d_3000_v144_)), d_3005_v149_)
                nw539_[int(15)] = d_3000_v144_
                nw539_[int(16)] = d_3000_v144_
                nw539_[int(17)] = d_3000_v144_
                d_3038_v168_ = nw539_
                index488_ = default__.safeIndex(878, (d_3038_v168_).length(0))
                (d_3038_v168_)[index488_] = d_3000_v144_
                d_3040_v169_: _dafny.Seq
                d_3040_v169_ = _dafny.SeqWithoutIsStrInference([d_3036_v166_, d_3036_v166_])
                d_3041_v170_: _dafny.Set
                d_3041_v170_ = _dafny.Set({d_3005_v149_, d_3005_v149_, d_3005_v149_, d_3005_v149_, d_3005_v149_})
                d_3042_v171_: _dafny.Set
                d_3042_v171_ = d_3041_v170_
                d_3043_v172_: _dafny.Map
                d_3043_v172_ = _dafny.Map({(self).f6: d_3042_v171_})
                d_3044_v173_: _dafny.Map
                d_3044_v173_ = _dafny.Map({d_3043_v172_: default__.fm26(not(d_3006_v150_.f7), (self).f6, len(d_3015_v156_), globalState)})
                d_3045_v174_: _dafny.Map
                d_3045_v174_ = _dafny.Map({(self).fm5(globalState): d_3043_v172_})
                d_3046_v175_: D21
                d_3046_v175_ = D21_DC58((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])
                d_3047_v176_: _dafny.Array
                nw540_ = _dafny.Array(D2.default()(), 22)
                d_3047_v176_ = nw540_
                d_3048_v177_: C4
                nw541_ = C4()
                nw541_.ctor__(d_3047_v176_, d_2812_v0_)
                d_3048_v177_ = nw541_
                d_3049_v178_: _dafny.Map
                d_3049_v178_ = _dafny.Map({d_3036_v166_.f7: d_3048_v177_})
                index489_ = default__.safeIndex(969, (d_3037_v167_).length(0))
                index490_ = default__.safeIndex(878, (d_3038_v168_).length(0))
                index491_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                rhs469_ = (d_3040_v169_)[default__.safeIndex((self).f6, len(d_3040_v169_))]
                rhs470_ = ((d_3044_v173_)[((d_3045_v174_)[(d_3046_v175_).cf99] if ((d_3046_v175_).cf99) in (d_3045_v174_) else _dafny.Map({len(_dafny.Map({(self).f6: d_3049_v178_})): d_3042_v171_}))] if (((d_3045_v174_)[(d_3046_v175_).cf99] if ((d_3046_v175_).cf99) in (d_3045_v174_) else _dafny.Map({len(_dafny.Map({(self).f6: d_3049_v178_})): d_3042_v171_}))) in (d_3044_v173_) else default__.fm13(d_3006_v150_.f7, len(d_3000_v144_), d_2812_v0_, globalState))
                rhs471_ = d_3000_v144_
                rhs472_ = ((0) - ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])) == ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])
                rhs473_ = ((self).f6) - ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])
                lhs281_ = d_3037_v167_
                lhs282_ = default__.safeIndex(969, (d_3037_v167_).length(0))
                lhs283_ = d_3038_v168_
                lhs284_ = default__.safeIndex(878, (d_3038_v168_).length(0))
                lhs285_ = d_3009_v153_
                lhs286_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                lhs281_[lhs282_] = rhs469_
                d_3005_v149_ = rhs470_
                lhs283_[lhs284_] = rhs471_
                r2 = rhs472_
                lhs285_[lhs286_] = rhs473_
                d_3006_v150_ = d_3006_v150_
                d_2813_v1_ = d_2813_v1_
                r3 = (self).f6
            elif True:
                d_3050___mcc_h19_ = source23_.cf61
                d_3051_cf61_ = d_3050___mcc_h19_
                d_3052_v179_: _dafny.Array
                nw542_ = _dafny.Array(None, 5)
                d_3052_v179_ = nw542_
                d_3053_v180_: _dafny.Array
                nw543_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_3053_v180_ = nw543_
                index492_ = default__.safeIndex(949, (d_3053_v180_).length(0))
                (d_3053_v180_)[index492_] = d_3009_v153_
                index493_ = default__.safeIndex(877, (d_3003_v147_).length(0))
                (d_3003_v147_)[index493_] = (386) != ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])
                index494_ = default__.safeIndex(949, (d_3053_v180_).length(0))
                index495_ = default__.safeIndex(877, (d_3003_v147_).length(0))
                rhs474_ = d_3052_v179_
                rhs475_ = (d_3000_v144_)[default__.safeIndex(len(d_3000_v144_), len(d_3000_v144_))]
                rhs476_ = d_2812_v0_
                rhs477_ = d_3009_v153_
                rhs478_ = d_3006_v150_.f7
                lhs287_ = d_3053_v180_
                lhs288_ = default__.safeIndex(949, (d_3053_v180_).length(0))
                lhs289_ = d_3003_v147_
                lhs290_ = default__.safeIndex(877, (d_3003_v147_).length(0))
                d_3052_v179_ = rhs474_
                d_3005_v149_ = rhs475_
                r2 = rhs476_
                lhs287_[lhs288_] = rhs477_
                lhs289_[lhs290_] = rhs478_
                d_3054_v181_: _dafny.Map
                d_3054_v181_ = _dafny.Map({False: d_3014_i8_})
                pat_let_tv82_ = d_3054_v181_
                pat_let_tv83_ = d_3006_v150_
                pat_let_tv84_ = d_3000_v144_
                pat_let_tv85_ = d_3000_v144_
                def iife248_(_pat_let95_0):
                    def iife249_(d_3055_dt__update__tmp_h4_):
                        def iife250_(_pat_let96_0):
                            def iife251_(d_3056_dt__update_hcf0_h2_):
                                return D0_DC0(d_3056_dt__update_hcf0_h2_)
                            return iife251_(_pat_let96_0)
                        return iife250_(len((pat_let_tv82_).set(not(pat_let_tv83_.f7), len(pat_let_tv84_))))
                    return iife249_(_pat_let95_0)
                def iife253_(_pat_let98_0):
                    def iife254_(d_3057_dt__update__tmp_h5_):
                        def iife255_(_pat_let99_0):
                            def iife256_(d_3058_dt__update_hcf0_h3_):
                                return D0_DC0(d_3058_dt__update_hcf0_h3_)
                            return iife256_(_pat_let99_0)
                        return iife255_(len(pat_let_tv85_))
                    return iife254_(_pat_let98_0)
                def iife252_(_pat_let97_0):
                    def iife257_(d_3059_dt__update__tmp_h6_):
                        def iife258_(_pat_let100_0):
                            def iife259_(d_3060_dt__update_hcf0_h4_):
                                return D0_DC0(d_3060_dt__update_hcf0_h4_)
                            return iife259_(_pat_let100_0)
                        return iife258_(d_3014_i8_)
                    return iife257_(_pat_let97_0)
                (d_3006_v150_).f7 = default__.fm1(_dafny.SeqWithoutIsStrInference([d_3008_v152_, d_3008_v152_, iife248_(d_3008_v152_), iife252_(iife253_(d_3008_v152_))]), (0) - (len(d_3001_v145_)), (self).f6, globalState)
                d_3061_v182_: _dafny.Array
                def lambda175_(d_3062_i11_):
                    return D2_DC7(_dafny.Set({(self).f6}))

                init98_ = lambda175_
                nw544_ = _dafny.Array(None, 8)
                for i0_98_ in range(nw544_.length(0)):
                    nw544_[i0_98_] = init98_(i0_98_)
                d_3061_v182_ = nw544_
                d_3063_v183_: D22
                d_3063_v183_ = D22_DC61(d_3061_v182_)
                d_3064_v184_: C5
                nw545_ = C5()
                nw545_.ctor__((d_3003_v147_)[default__.safeIndex(877, (d_3003_v147_).length(0))])
                d_3064_v184_ = nw545_
                d_3065_v185_: C4
                nw546_ = C4()
                nw546_.ctor__(d_3061_v182_, d_3006_v150_.f7)
                d_3065_v185_ = nw546_
                rhs479_ = (d_3001_v145_).set(d_3014_i8_, default__.safeModuloInt((0) - ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]), (self).f6))
                rhs480_ = False
                rhs481_ = D22_DC61((d_3065_v185_).f11)
                rhs482_ = d_3064_v184_
                rhs483_ = d_3065_v185_
                lhs291_ = d_3006_v150_
                d_3001_v145_ = rhs479_
                lhs291_.f7 = rhs480_
                d_3063_v183_ = rhs481_
                d_3064_v184_ = rhs482_
                d_3065_v185_ = rhs483_
                index496_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                (d_3009_v153_)[index496_] = (0) - ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])
            d_3066_v186_: _dafny.Map
            d_3066_v186_ = _dafny.Map({d_3006_v150_.f7: d_2999_v143_})
            d_3067_v187_: C7
            nw547_ = C7()
            nw547_.ctor__(d_3066_v186_, not(d_3006_v150_.f7))
            d_3067_v187_ = nw547_
            d_3067_v187_ = d_3067_v187_
            index497_ = default__.safeIndex(530, (d_3003_v147_).length(0))
            (d_3003_v147_)[index497_] = d_3006_v150_.f7
            index498_ = default__.safeIndex(530, (d_3003_v147_).length(0))
            (d_3003_v147_)[index498_] = ((self).f6) < ((self).f6)
            if (d_3067_v187_).fm8(((self).f6 if d_3006_v150_.f7 else 391), globalState):
                d_3068_v188_: _dafny.Map
                d_3068_v188_ = _dafny.Map({d_3014_i8_: _dafny.CodePoint('u')})
                d_3069_v189_: D23
                d_3069_v189_ = D23_DC64((d_3068_v188_) | (d_3068_v188_))
                pat_let_tv86_ = d_3068_v188_
                def iife260_(_pat_let101_0):
                    def iife261_(d_3070_dt__update__tmp_h7_):
                        def iife262_(_pat_let102_0):
                            def iife263_(d_3071_dt__update_hcf109_h0_):
                                return D23_DC64(d_3071_dt__update_hcf109_h0_)
                            return iife263_(_pat_let102_0)
                        return iife262_(pat_let_tv86_)
                    return iife261_(_pat_let101_0)
                rhs484_ = _dafny.SeqWithoutIsStrInference([-903])
                rhs485_ = iife260_(d_3069_v189_)
                d_2999_v143_ = rhs484_
                d_3069_v189_ = rhs485_
                d_2999_v143_ = (((d_2999_v143_) + (_dafny.SeqWithoutIsStrInference([(self).f6, (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], (self).f6, len(d_2999_v143_), (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]]))) + (d_2999_v143_)).set(default__.safeIndex((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], len(((d_2999_v143_) + (_dafny.SeqWithoutIsStrInference([(self).f6, (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], (self).f6, len(d_2999_v143_), (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]]))) + (d_2999_v143_))), (self).f6)
                d_3072_v190_: _dafny.Map
                d_3072_v190_ = _dafny.Map({d_3006_v150_.f7: d_2812_v0_})
                d_3073_v191_: _dafny.Map
                d_3073_v191_ = d_3072_v190_
                d_3074_v192_: _dafny.MultiSet
                d_3074_v192_ = _dafny.MultiSet([default__.fm21(d_3006_v150_.f7, globalState), _dafny.SeqWithoutIsStrInference([689 for d_3075_i12_ in range(default__.abs(14))])])
                rhs486_ = (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]
                rhs487_ = default__.safeDivisionInt(len(d_3000_v144_), (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])
                rhs488_ = d_3073_v191_
                rhs489_ = ((default__.fm61((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], False, globalState)).intersection(d_3074_v192_)).cardinality
                rhs490_ = d_3006_v150_.f7
                lhs292_ = d_3006_v150_
                r3 = rhs486_
                r3 = rhs487_
                d_3073_v191_ = rhs488_
                r1 = rhs489_
                lhs292_.f7 = rhs490_
                default__.m0((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], (0) - (d_3014_i8_), globalState)
                d_3076_v194_: _dafny.Array
                def lambda176_(d_3077_v188_, d_3078_v144_, d_3079_v149_):
                    def lambda177_(d_3080_i13_):
                        def iife264_():
                            coll58_ = _dafny.Map()
                            compr_58_: _dafny.Seq
                            for compr_58_ in (_dafny.SeqWithoutIsStrInference([d_3078_v144_])).Elements:
                                d_3081_v193_: _dafny.Seq = compr_58_
                                if (d_3081_v193_) in (_dafny.SeqWithoutIsStrInference([d_3078_v144_])):
                                    coll58_[d_3081_v193_] = d_3079_v149_
                            return _dafny.Map(coll58_)
                        return (iife264_()
                        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([((d_3077_v188_)[(self).f6] if ((self).f6) in (d_3077_v188_) else d_3079_v149_) for d_3082_i14_ in range(default__.abs(-365))]): d_3079_v149_}))

                    return lambda177_

                init99_ = lambda176_(d_3068_v188_, d_3000_v144_, d_3005_v149_)
                nw548_ = _dafny.Array(None, 26)
                for i0_99_ in range(nw548_.length(0)):
                    nw548_[i0_99_] = init99_(i0_99_)
                d_3076_v194_ = nw548_
                d_3083_v195_: _dafny.Map
                d_3083_v195_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plrvjxso")): d_3005_v149_})
                index499_ = default__.safeIndex(865, (d_3076_v194_).length(0))
                (d_3076_v194_)[index499_] = d_3083_v195_
                d_3084_v197_: _dafny.MultiSet
                d_3084_v197_ = _dafny.MultiSet([d_3000_v144_, d_3000_v144_])
                index500_ = default__.safeIndex(865, (d_3076_v194_).length(0))
                def iife265_():
                    coll59_ = _dafny.Map()
                    compr_59_: _dafny.Seq
                    for compr_59_ in (d_3084_v197_).Elements:
                        d_3085_v196_: _dafny.Seq = compr_59_
                        if (d_3085_v196_) in (d_3084_v197_):
                            coll59_[d_3085_v196_] = d_3005_v149_
                    return _dafny.Map(coll59_)
                (d_3076_v194_)[index500_] = (d_3083_v195_ if not(not(d_3006_v150_.f7)) else iife265_()
                )
            elif True:
                r3 = (d_2999_v143_)[default__.safeIndex(d_3014_i8_, len(d_2999_v143_))]
                d_3086_v198_: _dafny.Seq
                d_3086_v198_ = _dafny.SeqWithoutIsStrInference([default__.fm50(default__.fm30((d_3003_v147_)[default__.safeIndex(530, (d_3003_v147_).length(0))], (self).f6, (d_3003_v147_)[default__.safeIndex(530, (d_3003_v147_).length(0))], globalState), default__.fm62(d_3006_v150_.f7, (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], globalState), d_3014_i8_, globalState)])
                d_3086_v198_ = default__.fm41(d_3006_v150_.f7, globalState)
                r3 = (0) - ((0) - (d_3014_i8_))
                d_3001_v145_ = (d_3001_v145_).set((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], (0) - (default__.safeModuloInt((self).f6, (0) - ((0) - ((d_3002_v146_).cardinality)))))
                d_3087_v199_: C5
                nw549_ = C5()
                nw549_.ctor__(d_2812_v0_)
                d_3087_v199_ = nw549_
        d_3088_v200_: _dafny.Set
        d_3088_v200_ = _dafny.Set({d_2812_v0_, d_2812_v0_})
        d_3089_v201_: D1
        d_3089_v201_ = D1_DC3(d_3088_v200_)
        pat_let_tv87_ = d_3006_v150_
        def lambda178_(source24_):
            if source24_.is_DC4:
                d_3090___mcc_h20_ = source24_.cf7
                d_3091___mcc_h21_ = source24_.cf8
                d_3092_cf8_ = d_3091___mcc_h21_
                d_3093_cf7_ = d_3090___mcc_h20_
                return d_3093_cf7_
            elif source24_.is_DC5:
                d_3094___mcc_h22_ = source24_.cf9
                d_3095___mcc_h23_ = source24_.cf10
                d_3096___mcc_h24_ = source24_.cf11
                d_3097_cf11_ = d_3096___mcc_h24_
                d_3098_cf10_ = d_3095___mcc_h23_
                d_3099_cf9_ = d_3094___mcc_h22_
                return pat_let_tv87_.f7
            elif source24_.is_DC3:
                d_3100___mcc_h25_ = source24_.cf6
                d_3101_cf6_ = d_3100___mcc_h25_
                return False
            elif True:
                d_3102___mcc_h26_ = source24_.cf12
                d_3103_cf12_ = d_3102___mcc_h26_
                return False

        if lambda178_((d_3089_v201_ if True else d_3089_v201_)):
            index501_ = default__.safeIndex(456, (d_3009_v153_).length(0))
            (d_3009_v153_)[index501_] = (self).f6
            r2 = d_3006_v150_.f7
            d_3104_v202_: D6
            d_3104_v202_ = D6_DC16(d_3006_v150_.f7)
            if (d_3104_v202_).cf30:
                (d_3006_v150_).f7 = d_3006_v150_.f7
                d_3105_v203_: C2
                nw550_ = C2()
                nw550_.ctor__(_dafny.Set({d_3003_v147_, d_3003_v147_, d_3003_v147_, d_3003_v147_}), d_3005_v149_, (len(d_2999_v143_)) < ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]))
                d_3105_v203_ = nw550_
                d_3105_v203_ = d_3105_v203_
                d_3106_v204_: _dafny.MultiSet
                d_3106_v204_ = _dafny.MultiSet([(d_3000_v144_)[default__.safeIndex((0) - ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]), len(d_3000_v144_))]])
                d_3107_v205_: _dafny.MultiSet
                d_3107_v205_ = _dafny.MultiSet([d_3106_v204_])
                d_3108_v206_: _dafny.Seq
                d_3108_v206_ = _dafny.SeqWithoutIsStrInference([d_3107_v205_])
                d_3107_v205_ = ((d_3108_v206_)[default__.safeIndex((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))], len(d_3108_v206_))]).intersection(d_3107_v205_)
                r1 = ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]) * ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tin")))) - ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]))
                d_3109_v207_: _dafny.Map
                d_3109_v207_ = _dafny.Map({default__.safeDivisionInt(((d_3002_v146_)[not(d_3006_v150_.f7)] if (not(d_3006_v150_.f7)) in (d_3002_v146_) else default__.fm0(globalState)), (self).f6): (d_3105_v203_).f16})
                d_3109_v207_ = (d_3109_v207_).set((self).f6, d_3005_v149_)
            elif True:
                d_3110_v208_: _dafny.Array
                def lambda179_(d_3111_v153_):
                    def lambda180_(d_3112_i15_):
                        return _dafny.SeqWithoutIsStrInference([(d_3111_v153_)[default__.safeIndex(456, (d_3111_v153_).length(0))] for d_3113_i16_ in range(default__.abs(966))])

                    return lambda180_

                init100_ = lambda179_(d_3009_v153_)
                nw551_ = _dafny.Array(None, 6)
                for i0_100_ in range(nw551_.length(0)):
                    nw551_[i0_100_] = init100_(i0_100_)
                d_3110_v208_ = nw551_
                index502_ = default__.safeIndex(488, (d_3110_v208_).length(0))
                (d_3110_v208_)[index502_] = _dafny.SeqWithoutIsStrInference([(D6_DC17(d_3006_v150_.f7, (self).f6, d_3005_v149_, d_3005_v149_)).cf32 for d_3114_i17_ in range(default__.abs(746))])
                d_3115_v209_: _dafny.Set
                d_3115_v209_ = _dafny.Set({default__.safeDivisionInt(756, (0) - (len(d_2999_v143_))), (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]})
                d_3116_v210_: _dafny.Seq
                d_3116_v210_ = _dafny.SeqWithoutIsStrInference([d_2999_v143_])
                index503_ = default__.safeIndex(488, (d_3110_v208_).length(0))
                rhs491_ = ((d_3116_v210_)[default__.safeIndex(len(d_3000_v144_), len(d_3116_v210_))]) + ((_dafny.SeqWithoutIsStrInference([-261, (self).f6])) + (d_2999_v143_))
                rhs492_ = ((d_3115_v209_).intersection(d_3115_v209_)) - (d_3115_v209_)
                lhs293_ = d_3110_v208_
                lhs294_ = default__.safeIndex(488, (d_3110_v208_).length(0))
                lhs293_[lhs294_] = rhs491_
                d_3115_v209_ = rhs492_
                d_3117_v211_: _dafny.Array
                def lambda181_(d_3118_v209_):
                    def lambda182_(d_3119_i18_):
                        return D2_DC7(d_3118_v209_)

                    return lambda182_

                init101_ = lambda181_(d_3115_v209_)
                nw552_ = _dafny.Array(None, 5)
                for i0_101_ in range(nw552_.length(0)):
                    nw552_[i0_101_] = init101_(i0_101_)
                d_3117_v211_ = nw552_
                d_3120_v212_: C4
                nw553_ = C4()
                nw553_.ctor__(d_3117_v211_, d_2812_v0_)
                d_3120_v212_ = nw553_
                d_3120_v212_ = d_3120_v212_
                index504_ = default__.safeIndex(306, (d_3003_v147_).length(0))
                (d_3003_v147_)[index504_] = (d_3006_v150_.f7 if d_2812_v0_ else d_3006_v150_.f7)
                index505_ = default__.safeIndex(306, (d_3003_v147_).length(0))
                (d_3003_v147_)[index505_] = d_3006_v150_.f7
                d_3121_v213_: _dafny.Seq
                d_3121_v213_ = _dafny.SeqWithoutIsStrInference([d_3008_v152_, d_3008_v152_])
                r0 = default__.fm1(d_3121_v213_, (778 if d_3006_v150_.f7 else (self).f6), (len(d_3000_v144_)) + ((0) - ((self).f6)), globalState)
                d_3122_v214_: _dafny.Array
                nw554_ = _dafny.Array(None, 19)
                d_3122_v214_ = nw554_
                d_3123_v215_: _dafny.Map
                d_3123_v215_ = _dafny.Map({d_3088_v200_: d_3122_v214_})
                d_3122_v214_ = ((d_3123_v215_)[d_3088_v200_] if (d_3088_v200_) in (d_3123_v215_) else d_3122_v214_)
            if ((self).f6) >= (default__.safeModuloInt((self).f6, (self).f6)):
                index506_ = default__.safeIndex(996, (d_3003_v147_).length(0))
                (d_3003_v147_)[index506_] = d_2812_v0_
                d_3124_v216_: _dafny.Array
                nw555_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_3124_v216_ = nw555_
                index507_ = default__.safeIndex(459, (d_3124_v216_).length(0))
                (d_3124_v216_)[index507_] = d_3005_v149_
                d_3125_v217_: _dafny.MultiSet
                d_3125_v217_ = _dafny.MultiSet([d_3005_v149_, d_3005_v149_])
                d_3126_v218_: _dafny.Seq
                d_3126_v218_ = _dafny.SeqWithoutIsStrInference([d_3002_v146_])
                index508_ = default__.safeIndex(996, (d_3003_v147_).length(0))
                index509_ = default__.safeIndex(459, (d_3124_v216_).length(0))
                index510_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                rhs493_ = (self).f6
                rhs494_ = (d_3125_v217_).isdisjoint(_dafny.MultiSet([d_3005_v149_, _dafny.CodePoint('n'), d_3005_v149_]))
                rhs495_ = (d_3005_v149_ if d_3006_v150_.f7 else d_3005_v149_)
                rhs496_ = default__.fm0(globalState)
                rhs497_ = not((d_3126_v218_) < (d_3126_v218_))
                lhs295_ = d_3003_v147_
                lhs296_ = default__.safeIndex(996, (d_3003_v147_).length(0))
                lhs297_ = d_3124_v216_
                lhs298_ = default__.safeIndex(459, (d_3124_v216_).length(0))
                lhs299_ = d_3009_v153_
                lhs300_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                r3 = rhs493_
                lhs295_[lhs296_] = rhs494_
                lhs297_[lhs298_] = rhs495_
                lhs299_[lhs300_] = rhs496_
                r0 = rhs497_
                r2 = True
                d_3127_v219_: _dafny.Map
                d_3128_v220_: _dafny.Set
                out8_: _dafny.Map
                out9_: _dafny.Set
                out8_, out9_ = (d_3006_v150_).m2((d_3003_v147_)[default__.safeIndex(996, (d_3003_v147_).length(0))], (self).f6, ((d_3124_v216_)[default__.safeIndex(459, (d_3124_v216_).length(0))]) not in (d_3000_v144_), globalState)
                d_3127_v219_ = out8_
                d_3128_v220_ = out9_
                index511_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                (d_3009_v153_)[index511_] = (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]
                d_3129_v221_: C2
                nw556_ = C2()
                nw556_.ctor__(d_3004_v148_, d_3005_v149_, d_3006_v150_.f7)
                d_3129_v221_ = nw556_
            elif True:
                d_3003_v147_ = d_3003_v147_
                d_3130_v222_: _dafny.Array
                nw557_ = _dafny.Array(None, 6)
                nw557_[int(0)] = (d_3000_v144_) + (d_3000_v144_)
                nw557_[int(1)] = d_3000_v144_
                nw557_[int(2)] = (d_3000_v144_) + (d_3000_v144_)
                nw557_[int(3)] = d_3000_v144_
                nw557_[int(4)] = d_3000_v144_
                nw557_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wc"))
                d_3130_v222_ = nw557_
                index512_ = default__.safeIndex(702, (d_3130_v222_).length(0))
                (d_3130_v222_)[index512_] = (d_3000_v144_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbrvad")))
                index513_ = default__.safeIndex(702, (d_3130_v222_).length(0))
                rhs498_ = ((self).fm6(d_3008_v152_, globalState)) + (d_3000_v144_)
                rhs499_ = d_3006_v150_
                rhs500_ = default__.safeModuloInt(len(d_3000_v144_), len((d_3001_v145_).set((self).f6, (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))])))
                rhs501_ = not(not((len(d_2999_v143_)) <= ((self).f6)))
                rhs502_ = d_2812_v0_
                lhs301_ = d_3130_v222_
                lhs302_ = default__.safeIndex(702, (d_3130_v222_).length(0))
                lhs303_ = d_3006_v150_
                lhs301_[lhs302_] = rhs498_
                d_3006_v150_ = rhs499_
                r3 = rhs500_
                r2 = rhs501_
                lhs303_.f7 = rhs502_
                (d_3006_v150_).f7 = d_2812_v0_
                index514_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                rhs503_ = (self).f6
                rhs504_ = d_3005_v149_
                lhs304_ = d_3009_v153_
                lhs305_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                lhs304_[lhs305_] = rhs503_
                d_3005_v149_ = rhs504_
                index515_ = default__.safeIndex(456, (d_3009_v153_).length(0))
                (d_3009_v153_)[index515_] = (0) - ((self).f6)
            d_3131_v223_: _dafny.Map
            d_3131_v223_ = _dafny.Map({False: d_3006_v150_.f7})
            index516_ = default__.safeIndex(700, (d_3003_v147_).length(0))
            (d_3003_v147_)[index516_] = ((d_3131_v223_)[d_3006_v150_.f7] if (d_3006_v150_.f7) in (d_3131_v223_) else d_2812_v0_)
            index517_ = default__.safeIndex(700, (d_3003_v147_).length(0))
            (d_3003_v147_)[index517_] = not (d_3006_v150_.f7) or (d_2812_v0_)
        elif True:
            r3 = (self).f6
            index518_ = default__.safeIndex(456, (d_3009_v153_).length(0))
            (d_3009_v153_)[index518_] = (default__.safeDivisionInt((self).f6, (self).fm5(globalState))) + (default__.fm0(globalState))
            d_3132_v224_: _dafny.MultiSet
            d_3132_v224_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f6 for d_3133_i19_ in range(default__.abs(668))])])
            index519_ = default__.safeIndex(456, (d_3009_v153_).length(0))
            (d_3009_v153_)[index519_] = default__.safeModuloInt((self).f6, (d_3132_v224_).cardinality)
            d_3134_v225_: D21
            d_3134_v225_ = D21_DC59(d_3003_v147_, (d_3000_v144_).set(default__.safeIndex(len(d_3000_v144_), len(d_3000_v144_)), d_3005_v149_), d_3009_v153_, True)
            d_3135_v226_: _dafny.Seq
            d_3135_v226_ = _dafny.SeqWithoutIsStrInference([d_3134_v225_, d_3134_v225_])
            index520_ = default__.safeIndex(456, (d_3009_v153_).length(0))
            rhs505_ = (0) - (default__.safeModuloInt((self).f6, len(d_3000_v144_)))
            rhs506_ = (d_3135_v226_) < ((_dafny.SeqWithoutIsStrInference([d_3134_v225_])) + (d_3135_v226_))
            lhs306_ = d_3009_v153_
            lhs307_ = default__.safeIndex(456, (d_3009_v153_).length(0))
            lhs308_ = d_3006_v150_
            lhs306_[lhs307_] = rhs505_
            lhs308_.f7 = rhs506_
            d_3136_v227_: D9
            d_3136_v227_ = D9_DC24()
            d_3136_v227_ = d_3136_v227_
        d_3137_v228_: _dafny.Map
        d_3137_v228_ = _dafny.Map({d_2812_v0_: d_3006_v150_.f7})
        d_3137_v228_ = (d_3137_v228_).set(d_2812_v0_, ((d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]) < ((self).f6))
        d_3138_v229_: T1
        nw558_ = C3()
        nw558_.ctor__((0) - ((self).f6), d_3006_v150_.f7)
        d_3138_v229_ = nw558_
        r0 = ((_dafny.MultiSet([d_3138_v229_, d_3138_v229_, d_3138_v229_])).cardinality) == (((self).f6) + (207))
        r1 = (self).f6
        r2 = d_3138_v229_.f7
        r3 = (d_3009_v153_)[default__.safeIndex(456, (d_3009_v153_).length(0))]
        return r0, r1, r2, r3

    @property
    def f6(self):
        return self._f6
