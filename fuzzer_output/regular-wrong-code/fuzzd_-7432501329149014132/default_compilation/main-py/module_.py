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
    def fm0(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(650, 384):
                d_0_v0_: int = compr_0_
                if ((650) <= (d_0_v0_)) and ((d_0_v0_) < (384)):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeDivisionInt(d_0_v0_, len(_dafny.Map({True: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))).cardinality})))]))
            return _dafny.Set(coll0_)
        return not((_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffhh"))), 917]))), 146})).issubset((_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False])), len(iife0_()
        )})) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))}))))

    @staticmethod
    def fm1(globalState):
        return D3_DC6(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): (D21_DC43(False, True)).cf59}))

    @staticmethod
    def fm2(globalState):
        return 823

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return _dafny.Set({(-334) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmgef"))))})

    @staticmethod
    def fm13(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([not(True), False, False, True, False]) if False else _dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([False, True])) + (_dafny.SeqWithoutIsStrInference([False, not(False), False, False, not(True)])))

    @staticmethod
    def fm14(globalState):
        return _dafny.Map({(D4_DC8(_dafny.CodePoint('f'))).cf13: 133})

    @staticmethod
    def fm15(p0, globalState):
        return ((_dafny.Set({_dafny.CodePoint('v')})) - (_dafny.Set({_dafny.CodePoint('n')}))) - (_dafny.Set({_dafny.CodePoint('f')}))

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([573, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egtxgco")))])])).Elements:
                d_1_v0_: _dafny.Seq = compr_1_
                if (d_1_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([573, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egtxgco")))])])):
                    coll1_ = coll1_.union(_dafny.Set([d_1_v0_]))
            return _dafny.Set(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uosh"))), len(_dafny.Map({_dafny.CodePoint('r'): False}))])).Elements:
                d_2_v1_: int = compr_2_
                if (d_2_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uosh"))), len(_dafny.Map({_dafny.CodePoint('r'): False}))])):
                    coll2_[default__.safeModuloInt(d_2_v1_, 574)] = len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True)])) for d_3_i0_ in range(default__.abs(883))])): True}))
            return _dafny.Map(coll2_)
        return ((iife1_()
        ) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([212, 481, len(_dafny.SeqWithoutIsStrInference([False])), len(iife2_()
        ), 864]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_4_i1_ in range(default__.abs(-263))])}))])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wq")))])}))) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([446 for d_5_i2_ in range(default__.abs(-188))]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsydd"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhwldry")))])).cardinality])}))

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        return _dafny.Set({(False) not in (_dafny.MultiSet([not(False), not(True)]))})

    @staticmethod
    def fm18(p0, globalState):
        return ((_dafny.MultiSet([393])) - (_dafny.MultiSet([len(_dafny.Set({False}))]))) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-554: 431})) for d_6_i0_ in range(default__.abs(991))])))

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([-88, ((0) - (len(_dafny.Set({not(True)})))) - (len(_dafny.Map({len(_dafny.Map({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfb")))]): 419})): len(_dafny.Set({True}))}))), ((0) - (len(_dafny.Map({False: not(not(True))})))) + (len(_dafny.Set({True, True, True, False, False}))), (-161) + (656)])

    @staticmethod
    def fm23(globalState):
        return D5_DC13(default__.safeModuloInt(471, len(_dafny.SeqWithoutIsStrInference([not(not(not(True))), True, False]))))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([True, (D23_DC46(False)).cf63])) + (_dafny.SeqWithoutIsStrInference([not(not(True))])))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(0) - ((529) + (877))])

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eltc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlqr")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))

    @staticmethod
    def fm31(globalState):
        if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), True])) for d_7_i0_ in range(default__.abs(809))]))).issubset(_dafny.MultiSet([_dafny.MultiSet([True])])):
            return (_dafny.MultiSet([False, False, False])) - (_dafny.MultiSet([not(True)]))
        elif True:
            return _dafny.MultiSet([False])

    @staticmethod
    def fm32(p0, globalState):
        return D4_DC8(_dafny.CodePoint('g'))

    @staticmethod
    def fm33(p0, p1, globalState):
        return (_dafny.Map({True: len(_dafny.Map({len(_dafny.Map({not(True): 186})): 788}))})) | (_dafny.Map({True: 446}))

    @staticmethod
    def fm34(p0, p1, globalState):
        return (_dafny.Set({True, not(False), not(True), True, True})).intersection((_dafny.Set({False, (D7_DC16(len(_dafny.Set({False, not(False), True, False, False})), True)).cf23})) - (_dafny.Set({False})))

    @staticmethod
    def fm35(p0, globalState):
        if (941) != (-607):
            return _dafny.CodePoint('p')
        elif True:
            return _dafny.CodePoint('g')

    @staticmethod
    def fm36(p0, p1, globalState):
        return ((_dafny.Map({False: _dafny.Set({-217})})) | (_dafny.Map({False: _dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))})}))) | ((_dafny.Map({True: _dafny.Set({-894})})) | (_dafny.Map({False: _dafny.Set({509, 517})})))

    @staticmethod
    def fm37(p0, globalState):
        return (_dafny.Map({50: not(not(True))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')])): True}))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(953, -151):
                    d_9_v1_: int = compr_5_
                    if ((953) <= (d_9_v1_)) and ((d_9_v1_) < (-151)):
                        coll5_[default__.safeModuloInt(d_9_v1_, 543)] = 922
                return _dafny.Map(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(953, -151):
                    d_9_v1_: int = compr_4_
                    if ((953) <= (d_9_v1_)) and ((d_9_v1_) < (-151)):
                        coll4_[default__.safeModuloInt(d_9_v1_, 543)] = 922
                return _dafny.Map(coll4_)
            compr_3_: _dafny.Map
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({622: False}), _dafny.Map({len(iife4_()
            ): not(True)})])).Elements:
                d_10_v0_: _dafny.Map = compr_3_
                if (d_10_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({622: False}), _dafny.Map({len(iife5_()
                ): not(True)})])):
                    coll3_[d_10_v0_] = len(_dafny.SeqWithoutIsStrInference([False, False, True, not(True)]))
            return _dafny.Map(coll3_)
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyem"))), 266])).Elements:
                d_11_v2_: int = compr_6_
                if (d_11_v2_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyem"))), 266])):
                    coll6_ = coll6_.union(_dafny.Set([(d_11_v2_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({31, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uu"))), 822}))}))])), (_dafny.MultiSet([_dafny.CodePoint('u'), _dafny.CodePoint('a')])).cardinality}))}))]))}))])))]))
            return _dafny.Set(coll6_)
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Map
            for compr_7_ in ((D27_DC52(_dafny.Map({_dafny.Map({-774: False}): len(_dafny.Set({True}))}))).cf70).keys.Elements:
                d_12_v3_: _dafny.Map = compr_7_
                if (d_12_v3_) in ((D27_DC52(_dafny.Map({_dafny.Map({-774: False}): len(_dafny.Set({True}))}))).cf70):
                    coll7_[d_12_v3_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "norftn")))
            return _dafny.Map(coll7_)
        return ((_dafny.Map({_dafny.Map({-297: True}): len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqqq"))), 867, (_dafny.MultiSet([301])).cardinality]) for d_8_i0_ in range(default__.abs(401))]))}) if False else iife3_()
        )) | ((_dafny.Map({_dafny.Map({534: (D25_DC50(False, _dafny.SeqWithoutIsStrInference([False]))).cf67}): len(iife6_()
        )})) | (iife7_()
        ))

    @staticmethod
    def fm39(p0, globalState):
        source0_ = D27_DC53((_dafny.MultiSet([850])).cardinality, len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False])})), _dafny.CodePoint('h'))
        if source0_.is_DC53:
            d_13___mcc_h0_ = source0_.cf71
            d_14___mcc_h1_ = source0_.cf72
            d_15___mcc_h2_ = source0_.cf73
            d_16_cf73_ = d_15___mcc_h2_
            d_17_cf72_ = d_14___mcc_h1_
            d_18_cf71_ = d_13___mcc_h0_
            return D11_DC27(d_17_cf72_, d_17_cf72_, True)
        elif True:
            d_19___mcc_h3_ = source0_.cf70
            d_20_cf70_ = d_19___mcc_h3_
            return D11_DC27(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_21_i0_ in range(default__.abs(501))])), len(_dafny.Map({False: len(_dafny.Set({True}))})), False)

    @staticmethod
    def fm40(p0, p1, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in (_dafny.Set({127, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpib")))})).Elements:
                d_22_v0_: int = compr_8_
                if (d_22_v0_) in (_dafny.Set({127, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpib")))})):
                    coll8_[default__.safeModuloInt(d_22_v0_, -920)] = 540
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: _dafny.Seq
            for compr_9_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_23_i0_ in range(default__.abs(384))]): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.Set({True}))}))})).keys.Elements:
                d_24_v1_: _dafny.Seq = compr_9_
                if (d_24_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_23_i0_ in range(default__.abs(384))]): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.Set({True}))}))})):
                    coll9_[d_24_v1_] = 404
            return _dafny.Map(coll9_)
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-324, 333, (0) - (len(_dafny.Set({(_dafny.MultiSet([False])).cardinality}))), -428, len(iife8_()
        )])])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(0) - (len(iife9_()
        ))])]))).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({not(not(not(False))), False})), 621, 321]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('i'): -144}))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-303]) for d_25_i1_ in range(default__.abs(784))]))))

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(70, 200):
                d_26_v0_: int = compr_10_
                if ((70) <= (d_26_v0_)) and ((d_26_v0_) < (200)):
                    coll10_[(d_26_v0_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdyuas"))))] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pufrg"))
            return _dafny.Map(coll10_)
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(661, 452):
                d_27_v1_: int = compr_11_
                if ((661) <= (d_27_v1_)) and ((d_27_v1_) < (452)):
                    coll11_[(d_27_v1_) + (len(_dafny.Set({541, -226})))] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_28_i0_ in range(default__.abs(807))])
            return _dafny.Map(coll11_)
        return (iife10_()
        ) | (iife11_()
        )

    @staticmethod
    def fm42(globalState):
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: int
            for compr_12_ in _dafny.IntegerRange(134, -601):
                d_29_v0_: int = compr_12_
                if ((134) <= (d_29_v0_)) and ((d_29_v0_) < (-601)):
                    coll12_ = coll12_.union(_dafny.Set([(d_29_v0_) - (649)]))
            return _dafny.Set(coll12_)
        return iife12_()
        

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]), _dafny.MultiSet([False]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))])), (0) - ((len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True}))]))) + (-806))])

    @staticmethod
    def fm44(p0, p1, globalState):
        return ((_dafny.Map({not(not(False)): _dafny.Map({False: (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhmtdera"))])).cardinality})}) if not(False) else _dafny.Map({not(True): _dafny.Map({True: 719})}))) | ((_dafny.Map({False: _dafny.Map({False: (_dafny.MultiSet([-528])).cardinality})}) if True else _dafny.Map({True: _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktdot")))})})))

    @staticmethod
    def fm45(globalState):
        return _dafny.Map({(len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bn")))])}))) * (62): default__.safeModuloInt(len(_dafny.Map({True: len(_dafny.Map({True: False}))})), 184)})

    @staticmethod
    def fm46(globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_30_i0_ in range(default__.abs(305))])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_31_i1_ in range(default__.abs(358))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shpnlys")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_32_i2_ in range(default__.abs(71))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_33_i3_ in range(default__.abs(197))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "grebnnqgo"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnn"))]))))

    @staticmethod
    def fm47(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))})

    @staticmethod
    def fm48(p0, p1, globalState):
        return D12_DC29()

    @staticmethod
    def fm49(p0, p1, globalState):
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in (_dafny.Map({len(_dafny.Map({len(_dafny.Map({False: 901})): False})): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frk"))): 749}))})).keys.Elements:
                d_34_v0_: int = compr_13_
                if (d_34_v0_) in (_dafny.Map({len(_dafny.Map({len(_dafny.Map({False: 901})): False})): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frk"))): 749}))})):
                    coll13_[default__.safeModuloInt(d_34_v0_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality)] = 381
            return _dafny.Map(coll13_)
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([iife13_()
        ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({754: 464})])))

    @staticmethod
    def fm50(globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_35_i0_ in range(default__.abs(838))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))])

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(435, 483):
                d_36_v0_: int = compr_14_
                if ((435) <= (d_36_v0_)) and ((d_36_v0_) < (483)):
                    coll14_[(d_36_v0_) * (963)] = True
            return _dafny.Map(coll14_)
        def iife15_():
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: _dafny.Seq
                for compr_17_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maokai"))]))).Elements:
                    d_37_v1_: _dafny.Seq = compr_17_
                    if (d_37_v1_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maokai"))]))):
                        coll17_[d_37_v1_] = True
                return _dafny.Map(coll17_)
            coll15_ = _dafny.Set()
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: _dafny.Seq
                for compr_16_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maokai"))]))).Elements:
                    d_37_v1_: _dafny.Seq = compr_16_
                    if (d_37_v1_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maokai"))]))):
                        coll16_[d_37_v1_] = True
                return _dafny.Map(coll16_)
            compr_15_: _dafny.Seq
            for compr_15_ in (iife16_()
            ).keys.Elements:
                d_38_v2_: _dafny.Seq = compr_15_
                if (d_38_v2_) in (iife17_()
                ):
                    coll15_ = coll15_.union(_dafny.Set([d_38_v2_]))
            return _dafny.Set(coll15_)
        return ((_dafny.MultiSet([_dafny.Map({346: True})])) | (_dafny.MultiSet([iife14_()
        ]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([439, len(iife15_()
        )])): True})])))

    @staticmethod
    def fm52(globalState):
        source1_ = D7_DC16(len(_dafny.Set({-805})), (D8_DC19(not(not(False)), True, False)).cf28)
        if source1_.is_DC16:
            d_39___mcc_h0_ = source1_.cf22
            d_40___mcc_h1_ = source1_.cf23
            d_41_cf23_ = d_40___mcc_h1_
            d_42_cf22_ = d_39___mcc_h0_
            return (_dafny.Set({_dafny.SeqWithoutIsStrInference([d_41_cf23_])})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([d_41_cf23_]), _dafny.SeqWithoutIsStrInference([d_41_cf23_]), _dafny.SeqWithoutIsStrInference([d_41_cf23_])}))
        elif source1_.is_DC17:
            d_43___mcc_h2_ = source1_.cf24
            d_44___mcc_h3_ = source1_.cf25
            d_45___mcc_h4_ = source1_.cf26
            d_46_cf26_ = d_45___mcc_h4_
            d_47_cf25_ = d_44___mcc_h3_
            d_48_cf24_ = d_43___mcc_h2_
            return _dafny.Set({_dafny.SeqWithoutIsStrInference([d_48_cf24_]), _dafny.SeqWithoutIsStrInference([True])})
        elif True:
            d_49___mcc_h5_ = source1_.cf21
            d_50_cf21_ = d_49___mcc_h5_
            return _dafny.Set({_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True])})

    @staticmethod
    def fm53(p0, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: _dafny.Set
            for compr_18_ in (_dafny.Set({_dafny.Set({True, False})})).Elements:
                d_51_v0_: _dafny.Set = compr_18_
                if (d_51_v0_) in (_dafny.Set({_dafny.Set({True, False})})):
                    coll18_[d_51_v0_] = _dafny.SeqWithoutIsStrInference([not(True), True])
            return _dafny.Map(coll18_)
        return ((_dafny.Map({_dafny.Set({False, not(not(False)), not(False)}): _dafny.SeqWithoutIsStrInference([False, (D5_DC12(False, True)).cf18])})) | (_dafny.Map({_dafny.Set({True, True}): _dafny.SeqWithoutIsStrInference([False, not(True)])}))) | ((_dafny.Map({_dafny.Set({False, not(True)}): _dafny.SeqWithoutIsStrInference([False, True])})) | (iife18_()
        ))

    @staticmethod
    def fm54(p0, globalState):
        if False:
            def iife19_():
                def iife21_():
                    coll21_ = _dafny.Map()
                    compr_21_: int
                    for compr_21_ in _dafny.IntegerRange(-639, -619):
                        d_52_v1_: int = compr_21_
                        if ((-639) <= (d_52_v1_)) and ((d_52_v1_) < (-619)):
                            coll21_[(d_52_v1_) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))).cardinality)] = False
                    return _dafny.Map(coll21_)
                coll19_ = _dafny.Map()
                def iife20_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(-639, -619):
                        d_52_v1_: int = compr_20_
                        if ((-639) <= (d_52_v1_)) and ((d_52_v1_) < (-619)):
                            coll20_[(d_52_v1_) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))).cardinality)] = False
                    return _dafny.Map(coll20_)
                compr_19_: _dafny.Map
                for compr_19_ in (_dafny.SeqWithoutIsStrInference([iife20_()
                , _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): False})])).Elements:
                    d_53_v0_: _dafny.Map = compr_19_
                    if (d_53_v0_) in (_dafny.SeqWithoutIsStrInference([iife21_()
                    , _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): False})])):
                        coll19_[d_53_v0_] = (_dafny.MultiSet([not(True)])).cardinality
                return _dafny.Map(coll19_)
            return D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teiiyw")), len((D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjdmjhv")), 672, 508, -508, _dafny.SeqWithoutIsStrInference([True, True, False]))).cf8), len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))])])), (_dafny.MultiSet([len(iife19_()
), len(_dafny.Map({198: _dafny.CodePoint('c')}))])).cardinality, _dafny.SeqWithoutIsStrInference([False]))
        elif True:
            return D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwblcko")), (0) - (-939), (0) - (len(_dafny.Map({len(_dafny.Map({False: len(_dafny.Set({True, not(False), True}))})): (0) - (len(_dafny.SeqWithoutIsStrInference([579])))}))), 929, _dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return D9_DC23()

    @staticmethod
    def fm56(p0, p1, p2, globalState):
        def iife22_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(838, 851):
                d_54_v0_: int = compr_22_
                if ((838) <= (d_54_v0_)) and ((d_54_v0_) < (851)):
                    coll22_[default__.safeDivisionInt(d_54_v0_, -661)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False]))])
            return _dafny.Map(coll22_)
        source2_ = D25_DC49(iife22_()
)
        if source2_.is_DC50:
            d_55___mcc_h0_ = source2_.cf67
            d_56___mcc_h1_ = source2_.cf68
            d_57_cf68_ = d_56___mcc_h1_
            d_58_cf67_ = d_55___mcc_h0_
            def iife23_():
                coll23_ = _dafny.Set()
                compr_23_: int
                for compr_23_ in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([625])).cardinality])).Elements:
                    d_59_v1_: int = compr_23_
                    if (d_59_v1_) in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([625])).cardinality])):
                        coll23_ = coll23_.union(_dafny.Set([default__.safeModuloInt(d_59_v1_, len(_dafny.Map({True: True})))]))
                return _dafny.Set(coll23_)
            return D8_DC18(iife23_()
)
        elif True:
            d_60___mcc_h2_ = source2_.cf66
            d_61_cf66_ = d_60___mcc_h2_
            return D8_DC18(_dafny.Set({-446, 19}))

    @staticmethod
    def fm57(p0, globalState):
        def iife24_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in _dafny.IntegerRange(802, -719):
                d_62_v0_: int = compr_24_
                if ((802) <= (d_62_v0_)) and ((d_62_v0_) < (-719)):
                    coll24_[(d_62_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lffpro"))))] = True
            return _dafny.Map(coll24_)
        if ((0) - (len(iife24_()
        ))) != (len(_dafny.SeqWithoutIsStrInference([False]))):
            if True:
                return D20_DC41(663)
            elif True:
                return D20_DC41(800)
        elif True:
            return D20_DC41(len(_dafny.SeqWithoutIsStrInference([not(False)])))

    @staticmethod
    def fm58(p0, globalState):
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: int
            for compr_25_ in _dafny.IntegerRange(16, 756):
                d_63_v0_: int = compr_25_
                if ((16) <= (d_63_v0_)) and ((d_63_v0_) < (756)):
                    coll25_[default__.safeModuloInt(d_63_v0_, 992)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvpt")))
            return _dafny.Map(coll25_)
        def iife26_():
            coll26_ = _dafny.Set()
            compr_26_: int
            for compr_26_ in (_dafny.SeqWithoutIsStrInference([-294 for d_64_i0_ in range(default__.abs(705))])).Elements:
                d_65_v1_: int = compr_26_
                if (d_65_v1_) in (_dafny.SeqWithoutIsStrInference([-294 for d_64_i0_ in range(default__.abs(705))])):
                    coll26_ = coll26_.union(_dafny.Set([(d_65_v1_) + (-584)]))
            return _dafny.Set(coll26_)
        return (_dafny.Map({_dafny.MultiSet([(D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwdmh")), 207, len(iife25_()
), (_dafny.MultiSet([True])).cardinality, _dafny.SeqWithoutIsStrInference([True]))).cf11]): _dafny.Set({False, False})})) | (_dafny.Map({_dafny.MultiSet([len(iife26_()
        )]): _dafny.Set({False, not(False), False, False})}))

    @staticmethod
    def fm59(p0, p1, globalState):
        return D4_DC10((D5_DC13(434)).cf19)

    @staticmethod
    def fm60(p0, p1, globalState):
        return D7_DC17((-471) < (-136), (-8) <= ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D7_DC16(-228, False), D7_DC16(-913, False), D7_DC16((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))).cardinality, False), D7_DC16(len(_dafny.Map({True: False})), True)]))).cardinality), (0) - (len(_dafny.Map({True: False}))))

    @staticmethod
    def fm61(globalState):
        source3_ = D4_DC10(-757)
        if source3_.is_DC9:
            d_66___mcc_h0_ = source3_.cf14
            d_67_cf14_ = d_66___mcc_h0_
            return (_dafny.Map({D5_DC11(True): _dafny.CodePoint('a')})) | (_dafny.Map({D5_DC11(True): _dafny.CodePoint('t')}))
        elif source3_.is_DC10:
            d_68___mcc_h1_ = source3_.cf15
            d_69_cf15_ = d_68___mcc_h1_
            return (_dafny.Map({D5_DC11(not(True)): _dafny.CodePoint('i')})) | (_dafny.Map({D5_DC11(True): _dafny.CodePoint('j')}))
        elif True:
            d_70___mcc_h2_ = source3_.cf13
            d_71_cf13_ = d_70___mcc_h2_
            def iife27_():
                coll27_ = _dafny.Map()
                compr_27_: D5
                for compr_27_ in (_dafny.Map({D5_DC11(False): len(_dafny.SeqWithoutIsStrInference([91, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_72_i0_ in range(default__.abs(313))])), 130, len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})), 936]))})).keys.Elements:
                    d_73_v0_: D5 = compr_27_
                    if (d_73_v0_) in (_dafny.Map({D5_DC11(False): len(_dafny.SeqWithoutIsStrInference([91, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_72_i0_ in range(default__.abs(313))])), 130, len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})), 936]))})):
                        coll27_[d_73_v0_] = d_71_cf13_
                return _dafny.Map(coll27_)
            return iife27_()
            

    @staticmethod
    def fm62(p0, p1, globalState):
        return _dafny.Set({D2_DC5(), D2_DC5(), D2_DC5(), D2_DC5()})

    @staticmethod
    def fm63(globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('e')])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s'), _dafny.CodePoint('u'), _dafny.CodePoint('a')])))) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('v')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o')]))))

    @staticmethod
    def fm64(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_74_i0_ in range(default__.abs(718))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gprxnyjj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gx"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okdgos")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfhrbisl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onpcunmrj")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_75_i1_ in range(default__.abs(512))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgthnk"))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kumprfk")) for d_76_i2_ in range(default__.abs(873))]))

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, not(True)])])

    @staticmethod
    def fm66(p0, p1, globalState):
        return _dafny.Map({_dafny.Set({True, True}): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryyq")))})

    @staticmethod
    def fm67(globalState):
        return D19_DC38((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_77_i0_ in range(default__.abs(901))]))) + (132), not (not(not(True))) or (True), ((0) - (len(_dafny.SeqWithoutIsStrInference([388])))) >= (len(_dafny.Map({not(True): True}))))

    @staticmethod
    def fm68(p0, p1, globalState):
        return ((_dafny.MultiSet([D5_DC12(False, False), D5_DC12(True, not(True)), D5_DC12(False, False), D5_DC12(False, True)])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D5_DC12(True, True), D5_DC12(not(True), True)])))).intersection(_dafny.MultiSet([D5_DC12(False, False), D5_DC12(True, not(True))]))

    @staticmethod
    def m0(globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        d_78_v0_: bool
        d_78_v0_ = False
        d_79_v1_: int
        d_79_v1_ = 716
        d_80_v2_: _dafny.Map
        d_80_v2_ = _dafny.Map({d_78_v0_: d_79_v1_})
        d_81_v3_: T1
        nw0_ = C1()
        nw0_.ctor__(d_80_v2_)
        d_81_v3_ = nw0_
        d_82_v4_: _dafny.Map
        d_82_v4_ = _dafny.Map({d_81_v3_: d_78_v0_})
        d_83_v5_: _dafny.Seq
        d_83_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsg"))
        d_84_v6_: _dafny.Set
        d_84_v6_ = _dafny.Set({len(d_83_v5_)})
        d_85_v7_: _dafny.Seq
        d_85_v7_ = _dafny.SeqWithoutIsStrInference([d_84_v6_, d_84_v6_, d_84_v6_])
        d_86_v8_: _dafny.Seq
        d_86_v8_ = _dafny.SeqWithoutIsStrInference([d_78_v0_, d_78_v0_])
        d_87_v9_: _dafny.Map
        d_87_v9_ = _dafny.Map({d_86_v8_: d_78_v0_})
        d_88_v10_: D3
        d_88_v10_ = D3_DC6(d_87_v9_)
        d_89_v11_: _dafny.Seq
        d_89_v11_ = _dafny.SeqWithoutIsStrInference([d_79_v1_, d_79_v1_, (D11_DC27(d_79_v1_, d_79_v1_, d_78_v0_)).cf41, (0) - (d_79_v1_)])
        d_90_v12_: _dafny.Map
        d_90_v12_ = _dafny.Map({d_89_v11_: len(d_83_v5_)})
        d_91_v13_: _dafny.Array
        nw1_ = _dafny.Array(None, 27)
        nw1_[int(0)] = d_78_v0_
        nw1_[int(1)] = (d_78_v0_) or (True)
        nw1_[int(2)] = not(((0) - (default__.fm2(globalState))) < (d_79_v1_))
        nw1_[int(3)] = d_78_v0_
        nw1_[int(4)] = not(not (d_78_v0_) or (not(d_78_v0_)))
        nw1_[int(5)] = ((0) - (d_79_v1_)) <= (663)
        nw1_[int(6)] = not (d_78_v0_) or (d_78_v0_)
        nw1_[int(7)] = d_78_v0_
        nw1_[int(8)] = d_78_v0_
        nw1_[int(9)] = d_78_v0_
        nw1_[int(10)] = (-288) == (len(d_82_v4_))
        nw1_[int(11)] = not((_dafny.Set({d_79_v1_})).issubset((d_85_v7_)[default__.safeIndex(d_79_v1_, len(d_85_v7_))]))
        nw1_[int(12)] = d_78_v0_
        nw1_[int(13)] = d_78_v0_
        nw1_[int(14)] = default__.fm0(len(d_86_v8_), globalState)
        nw1_[int(15)] = d_78_v0_
        nw1_[int(16)] = d_78_v0_
        nw1_[int(17)] = d_78_v0_
        nw1_[int(18)] = d_78_v0_
        nw1_[int(19)] = d_78_v0_
        nw1_[int(20)] = (default__.fm7((0) - (-201), d_88_v10_, d_78_v0_, globalState)).isdisjoint(default__.fm42(globalState))
        nw1_[int(21)] = d_78_v0_
        nw1_[int(22)] = d_78_v0_
        nw1_[int(23)] = not(True)
        nw1_[int(24)] = d_78_v0_
        nw1_[int(25)] = (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_79_v1_, d_79_v1_]): d_79_v1_})) != (d_90_v12_)
        nw1_[int(26)] = d_78_v0_
        d_91_v13_ = nw1_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_91_v13_).length(0)):
            d_92_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_92_i0_)) and ((d_92_i0_) < ((d_91_v13_).length(0)))):
                (d_91_v13_)[(d_92_i0_)] = (d_79_v1_) == (d_79_v1_)
        d_93_v14_: _dafny.Array
        nw2_ = _dafny.Array(int(0), 24)
        d_93_v14_ = nw2_
        d_94_v15_: _dafny.Seq
        d_94_v15_ = _dafny.SeqWithoutIsStrInference([d_93_v14_])
        d_95_v16_: _dafny.Seq
        d_95_v16_ = _dafny.SeqWithoutIsStrInference([d_91_v13_, d_91_v13_, d_91_v13_])
        rhs0_ = (d_79_v1_) * (d_79_v1_)
        rhs1_ = (d_93_v14_) not in (d_94_v15_)
        rhs2_ = (d_95_v16_)[default__.safeIndex(default__.fm2(globalState), len(d_95_v16_))]
        rhs3_ = d_78_v0_
        rhs4_ = d_79_v1_
        lhs0_ = globalState
        d_79_v1_ = rhs0_
        r0 = rhs1_
        d_91_v13_ = rhs2_
        r1 = rhs3_
        lhs0_.f2 = rhs4_
        d_84_v6_ = (d_84_v6_) - (d_84_v6_)
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_91_v13_).length(0)):
            d_96_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_96_i1_)) and ((d_96_i1_) < ((d_91_v13_).length(0)))):
                (d_91_v13_)[(d_96_i1_)] = d_78_v0_
        r3 = (d_79_v1_ if d_78_v0_ else d_79_v1_)
        index0_ = default__.safeIndex(627, (d_91_v13_).length(0))
        (d_91_v13_)[index0_] = not(d_78_v0_)
        d_97_v17_: str
        d_97_v17_ = _dafny.CodePoint('o')
        d_98_v18_: _dafny.Map
        d_98_v18_ = _dafny.Map({d_79_v1_: d_97_v17_})
        d_99_v19_: _dafny.Map
        d_99_v19_ = _dafny.Map({d_79_v1_: d_79_v1_})
        d_100_v20_: _dafny.MultiSet
        d_100_v20_ = _dafny.MultiSet([d_99_v19_])
        d_101_v21_: _dafny.Map
        d_101_v21_ = _dafny.Map({True: _dafny.CodePoint('h')})
        d_102_v22_: _dafny.Array
        nw3_ = _dafny.Array(None, 26)
        nw3_[int(0)] = d_97_v17_
        nw3_[int(1)] = d_97_v17_
        nw3_[int(2)] = d_97_v17_
        nw3_[int(3)] = default__.fm35(d_79_v1_, globalState)
        nw3_[int(4)] = d_97_v17_
        nw3_[int(5)] = ((d_98_v18_)[default__.fm2(globalState)] if (default__.fm2(globalState)) in (d_98_v18_) else default__.fm35((d_100_v20_).cardinality, globalState))
        nw3_[int(6)] = _dafny.CodePoint('r')
        nw3_[int(7)] = d_97_v17_
        nw3_[int(8)] = (d_83_v5_)[default__.safeIndex(len(d_86_v8_), len(d_83_v5_))]
        nw3_[int(9)] = ((d_101_v21_)[d_78_v0_] if (d_78_v0_) in (d_101_v21_) else d_97_v17_)
        nw3_[int(10)] = d_97_v17_
        nw3_[int(11)] = d_97_v17_
        nw3_[int(12)] = d_97_v17_
        nw3_[int(13)] = d_97_v17_
        nw3_[int(14)] = d_97_v17_
        nw3_[int(15)] = d_97_v17_
        nw3_[int(16)] = d_97_v17_
        nw3_[int(17)] = d_97_v17_
        nw3_[int(18)] = d_97_v17_
        nw3_[int(19)] = d_97_v17_
        nw3_[int(20)] = d_97_v17_
        nw3_[int(21)] = d_97_v17_
        nw3_[int(22)] = d_97_v17_
        nw3_[int(23)] = d_97_v17_
        nw3_[int(24)] = d_97_v17_
        nw3_[int(25)] = _dafny.CodePoint('d')
        d_102_v22_ = nw3_
        d_103_v23_: D0
        d_103_v23_ = D0_DC0(d_102_v22_)
        index1_ = default__.safeIndex(627, (d_91_v13_).length(0))
        rhs5_ = d_78_v0_
        rhs6_ = d_103_v23_
        lhs1_ = d_91_v13_
        lhs2_ = default__.safeIndex(627, (d_91_v13_).length(0))
        lhs1_[lhs2_] = rhs5_
        d_103_v23_ = rhs6_
        r0 = not((d_86_v8_) <= ((d_86_v8_) + (d_86_v8_)))
        r1 = (d_91_v13_)[default__.safeIndex(627, (d_91_v13_).length(0))]
        d_104_v24_: _dafny.Set
        d_104_v24_ = _dafny.Set({d_97_v17_})
        d_105_v25_: _dafny.Map
        d_105_v25_ = _dafny.Map({(d_79_v1_) > (len(d_104_v24_)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqm"))})
        r2 = d_105_v25_
        r3 = (0) - (d_79_v1_)
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_106_v0_: bool
        d_106_v0_ = True
        d_107_v1_: _dafny.Seq
        d_107_v1_ = _dafny.SeqWithoutIsStrInference([174])
        d_108_v2_: _dafny.Seq
        d_108_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jobjqyr"))
        d_109_v3_: _dafny.Array
        nw4_ = _dafny.Array(_dafny.CodePoint('D'), 24)
        d_109_v3_ = nw4_
        d_110_v4_: D0
        d_110_v4_ = D0_DC0(d_109_v3_)
        d_111_v5_: _dafny.Array
        nw5_ = _dafny.Array(int(0), 4)
        d_111_v5_ = nw5_
        d_112_v6_: int
        d_112_v6_ = 766
        d_113_v7_: _dafny.Map
        d_113_v7_ = _dafny.Map({d_111_v5_: d_112_v6_})
        d_114_v8_: _dafny.Seq
        d_114_v8_ = _dafny.SeqWithoutIsStrInference([d_108_v2_])
        d_115_v9_: _dafny.Array
        nw6_ = _dafny.Array(None, 11)
        nw6_[int(0)] = d_108_v2_
        nw6_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_116_i0_ in range(default__.abs(660))])
        nw6_[int(2)] = d_108_v2_
        nw6_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgw"))
        nw6_[int(4)] = d_108_v2_
        nw6_[int(5)] = d_108_v2_
        nw6_[int(6)] = (d_114_v8_)[default__.safeIndex(d_112_v6_, len(d_114_v8_))]
        nw6_[int(7)] = d_108_v2_
        nw6_[int(8)] = d_108_v2_
        nw6_[int(9)] = d_108_v2_
        nw6_[int(10)] = d_108_v2_
        d_115_v9_ = nw6_
        d_117_v10_: _dafny.Map
        d_117_v10_ = _dafny.Map({d_112_v6_: d_112_v6_})
        d_118_v11_: _dafny.Map
        d_118_v11_ = _dafny.Map({d_111_v5_: d_106_v0_})
        d_119_v12_: _dafny.Array
        nw7_ = _dafny.Array(False, 21)
        d_119_v12_ = nw7_
        d_120_v13_: _dafny.Seq
        d_120_v13_ = _dafny.SeqWithoutIsStrInference([d_106_v0_])
        d_121_v14_: _dafny.Map
        d_121_v14_ = _dafny.Map({d_119_v12_: d_120_v13_})
        d_122_globalState_: GlobalState
        nw8_ = GlobalState()
        nw8_.ctor__(2, False, 503, (d_107_v1_ if d_106_v0_ else (d_107_v1_).set(default__.safeIndex(len(d_108_v2_), len(d_107_v1_)), len(_dafny.SeqWithoutIsStrInference([d_106_v0_, not(d_106_v0_)])))), (d_110_v4_).cf0, 65, False, (d_113_v7_), d_115_v9_, 502, 969, True, d_117_v10_, False, (d_118_v11_ if d_106_v0_ else _dafny.Map({d_111_v5_: d_106_v0_})), 379, 874, 885, (d_121_v14_) | (d_121_v14_), 184, True, 592, 436, 68, 644, -277)
        d_122_globalState_ = nw8_
        d_123_v15_: _dafny.Seq
        d_123_v15_ = _dafny.SeqWithoutIsStrInference([d_109_v3_])
        d_124_v16_: _dafny.MultiSet
        d_124_v16_ = _dafny.MultiSet([d_109_v3_, (d_123_v15_)[default__.safeIndex(len(d_108_v2_), len(d_123_v15_))], d_109_v3_, d_109_v3_, d_109_v3_])
        (d_122_globalState_).f1 = (d_124_v16_).ispropersubset(d_124_v16_)
        hi0_ = len((d_120_v13_) + (_dafny.SeqWithoutIsStrInference([not(d_106_v0_), d_106_v0_])))
        for d_125_i1_ in range(d_112_v6_, hi0_):
            d_126_v17_: bool
            d_127_v18_: bool
            d_128_v19_: _dafny.Map
            d_129_v20_: int
            out0_: bool
            out1_: bool
            out2_: _dafny.Map
            out3_: int
            out0_, out1_, out2_, out3_ = default__.m0(d_122_globalState_)
            d_126_v17_ = out0_
            d_127_v18_ = out1_
            d_128_v19_ = out2_
            d_129_v20_ = out3_
            d_130_v21_: _dafny.MultiSet
            d_130_v21_ = _dafny.MultiSet([d_125_i1_])
            d_112_v6_ = (d_130_v21_).cardinality
            nw9_ = _dafny.Array(_dafny.CodePoint('D'), 6)
            d_109_v3_ = nw9_
            d_131_v22_: str
            d_131_v22_ = _dafny.CodePoint('w')
            d_131_v22_ = d_131_v22_
        hi1_ = d_112_v6_
        for d_132_i2_ in range(73, hi1_):
            (d_122_globalState_).f13 = (d_132_i2_) > (d_112_v6_)
            d_106_v0_ = not(d_106_v0_)
            d_106_v0_ = (d_106_v0_ if (d_106_v0_ if not(d_106_v0_) else default__.fm0(d_112_v6_, d_122_globalState_)) else d_106_v0_)
            d_133_v23_: _dafny.Array
            nw10_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_133_v23_ = nw10_
            d_134_v24_: _dafny.Map
            d_134_v24_ = _dafny.Map({d_133_v23_: not(True)})
            d_134_v24_ = (d_134_v24_).set(d_133_v23_, (d_106_v0_) and (d_106_v0_))
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_119_v12_).length(0)):
            d_135_i3_: int = guard_loop_2_
            if (True) and (((0) <= (d_135_i3_)) and ((d_135_i3_) < ((d_119_v12_).length(0)))):
                (d_119_v12_)[(d_135_i3_)] = not((d_106_v0_) or (d_106_v0_))
        d_120_v13_ = d_120_v13_
        d_136_v25_: _dafny.Array
        nw11_ = _dafny.Array(_dafny.MultiSet({}), 25)
        d_136_v25_ = nw11_
        d_137_v26_: D2
        d_137_v26_ = D2_DC4(d_136_v25_)
        d_136_v25_ = ((d_137_v26_ if default__.fm0(d_112_v6_, d_122_globalState_) else d_137_v26_)).cf6
        d_138_v27_: bool
        d_139_v28_: bool
        d_140_v29_: _dafny.Map
        d_141_v30_: int
        out4_: bool
        out5_: bool
        out6_: _dafny.Map
        out7_: int
        out4_, out5_, out6_, out7_ = default__.m0(d_122_globalState_)
        d_138_v27_ = out4_
        d_139_v28_ = out5_
        d_140_v29_ = out6_
        d_141_v30_ = out7_
        d_142_v31_: bool
        d_143_v32_: bool
        d_144_v33_: _dafny.Map
        d_145_v34_: int
        out8_: bool
        out9_: bool
        out10_: _dafny.Map
        out11_: int
        out8_, out9_, out10_, out11_ = default__.m0(d_122_globalState_)
        d_142_v31_ = out8_
        d_143_v32_ = out9_
        d_144_v33_ = out10_
        d_145_v34_ = out11_
        d_146_v36_: _dafny.Array
        def lambda0_(d_147_v13_, d_148_v32_):
            def lambda1_(d_149_i4_):
                def iife28_():
                    coll28_ = _dafny.Map()
                    compr_28_: _dafny.Seq
                    for compr_28_ in (_dafny.MultiSet([d_147_v13_, d_147_v13_])).Elements:
                        d_150_v35_: _dafny.Seq = compr_28_
                        if (d_150_v35_) in (_dafny.MultiSet([d_147_v13_, d_147_v13_])):
                            coll28_[d_150_v35_] = d_148_v32_
                    return _dafny.Map(coll28_)
                return (D3_DC6(iife28_()
)).cf7

            return lambda1_

        init0_ = lambda0_(d_120_v13_, d_143_v32_)
        nw12_ = _dafny.Array(None, 14)
        for i0_0_ in range(nw12_.length(0)):
            nw12_[i0_0_] = init0_(i0_0_)
        d_146_v36_ = nw12_
        d_151_v37_: _dafny.Map
        d_151_v37_ = _dafny.Map({d_120_v13_: d_139_v28_})
        index2_ = default__.safeIndex(798, (d_146_v36_).length(0))
        (d_146_v36_)[index2_] = (d_151_v37_) | (d_151_v37_)
        d_152_v39_: _dafny.Set
        d_152_v39_ = _dafny.Set({d_120_v13_})
        index3_ = default__.safeIndex(798, (d_146_v36_).length(0))
        def iife29_():
            coll29_ = _dafny.Map()
            compr_29_: _dafny.Seq
            for compr_29_ in (d_152_v39_).Elements:
                d_153_v38_: _dafny.Seq = compr_29_
                if (d_153_v38_) in (d_152_v39_):
                    coll29_[d_153_v38_] = d_139_v28_
            return _dafny.Map(coll29_)
        (d_146_v36_)[index3_] = (((default__.fm1(d_122_globalState_)).cf7) | (iife29_()
        )) | (d_151_v37_)
        d_154_v40_: D0
        d_154_v40_ = D0_DC0(d_109_v3_)
        d_155_v41_: D0
        d_155_v41_ = D0_DC2(d_154_v40_)
        source4_ = d_155_v41_
        if source4_.is_DC1:
            d_156___mcc_h0_ = source4_.cf1
            d_157___mcc_h1_ = source4_.cf2
            d_158___mcc_h2_ = source4_.cf3
            d_159_cf3_ = d_158___mcc_h2_
            d_160_cf2_ = d_157___mcc_h1_
            d_161_cf1_ = d_156___mcc_h0_
            d_162_v42_: _dafny.Map
            d_162_v42_ = _dafny.Map({(d_141_v30_) < (d_159_cf3_): default__.fm0(d_159_cf3_, d_122_globalState_)})
            d_162_v42_ = (d_162_v42_).set(True, not (d_142_v31_) or (d_106_v0_))
            (d_122_globalState_).f2 = (default__.fm2(d_122_globalState_)) * (default__.safeDivisionInt(d_112_v6_, d_145_v34_))
            d_119_v12_ = d_119_v12_
            d_108_v2_ = d_108_v2_
        elif source4_.is_DC0:
            d_163___mcc_h3_ = source4_.cf0
            d_164_cf0_ = d_163___mcc_h3_
            d_165_v43_: C0
            nw13_ = C0()
            nw13_.ctor__(d_109_v3_)
            d_165_v43_ = nw13_
            d_166_v44_: C3
            nw14_ = C3()
            nw14_.ctor__(d_111_v5_, d_106_v0_, d_117_v10_)
            d_166_v44_ = nw14_
            index4_ = default__.safeIndex(835, ((d_166_v44_).f39).length(0))
            ((d_166_v44_).f39)[index4_] = d_112_v6_
            d_167_v45_: _dafny.Map
            d_167_v45_ = _dafny.Map({d_141_v30_: d_139_v28_})
            d_168_v46_: D3
            d_168_v46_ = D3_DC7(d_108_v2_, len(_dafny.SeqWithoutIsStrInference([d_138_v27_])), d_145_v34_, (d_107_v1_)[default__.safeIndex(len(d_167_v45_), len(d_107_v1_))], d_120_v13_)
            index5_ = default__.safeIndex(835, ((d_166_v44_).f39).length(0))
            ((d_166_v44_).f39)[index5_] = len((d_168_v46_).cf8)
            (d_122_globalState_).f16 = ((d_166_v44_).f39)[default__.safeIndex(835, ((d_166_v44_).f39).length(0))]
        elif True:
            d_169___mcc_h4_ = source4_.cf4
            d_170_cf4_ = d_169___mcc_h4_
            d_171_v47_: D7
            d_171_v47_ = D7_DC16(default__.safeDivisionInt(d_112_v6_, d_141_v30_), not((not(not(d_143_v32_))) and (d_138_v27_)))
            d_172_v48_: C7
            nw15_ = C7()
            nw15_.ctor__(d_106_v0_)
            d_172_v48_ = nw15_
            d_173_v49_: _dafny.MultiSet
            d_173_v49_ = _dafny.MultiSet([d_138_v27_, False])
            d_174_v50_: _dafny.Map
            d_174_v50_ = _dafny.Map({d_106_v0_: ((d_173_v49_).cardinality) == (((d_173_v49_)[d_138_v27_] if (d_138_v27_) in (d_173_v49_) else d_141_v30_))})
            rhs7_ = d_171_v47_
            rhs8_ = (0) - ((d_145_v34_ if d_138_v27_ else (d_141_v30_) + (d_145_v34_)))
            rhs9_ = len(d_174_v50_)
            rhs10_ = d_172_v48_
            lhs3_ = d_122_globalState_
            d_171_v47_ = rhs7_
            d_112_v6_ = rhs8_
            lhs3_.f0 = rhs9_
            d_172_v48_ = rhs10_
            d_175_v51_: _dafny.Map
            d_175_v51_ = _dafny.Map({default__.fm0((0) - (d_112_v6_), d_122_globalState_): 480})
            d_176_v52_: C1
            nw16_ = C1()
            nw16_.ctor__(d_175_v51_)
            d_176_v52_ = nw16_
            d_177_v53_: _dafny.Seq
            d_177_v53_ = _dafny.SeqWithoutIsStrInference([d_176_v52_])
            d_178_v54_: _dafny.Map
            d_178_v54_ = _dafny.Map({d_112_v6_: d_176_v52_})
            rhs11_ = d_141_v30_
            rhs12_ = ((((d_177_v53_) + (d_177_v53_)).set(default__.safeIndex(d_141_v30_, len((d_177_v53_) + (d_177_v53_))), ((d_178_v54_)[d_112_v6_] if (d_112_v6_) in (d_178_v54_) else d_176_v52_))).set(default__.safeIndex((0) - (d_145_v34_), len(((d_177_v53_) + (d_177_v53_)).set(default__.safeIndex(d_141_v30_, len((d_177_v53_) + (d_177_v53_))), ((d_178_v54_)[d_112_v6_] if (d_112_v6_) in (d_178_v54_) else d_176_v52_)))), d_176_v52_)) + ((d_177_v53_) + (d_177_v53_))
            lhs4_ = d_122_globalState_
            lhs4_.f2 = rhs11_
            d_177_v53_ = rhs12_
            d_179_v55_: C3
            nw17_ = C3()
            nw17_.ctor__(d_111_v5_, d_139_v28_, d_117_v10_)
            d_179_v55_ = nw17_
            d_180_v56_: C3
            d_180_v56_ = d_179_v55_
            d_181_v57_: _dafny.MultiSet
            d_181_v57_ = _dafny.MultiSet([(d_180_v56_)])
            d_182_v58_: C11
            nw18_ = C11()
            nw18_.ctor__((d_181_v57_).cardinality, _dafny.Map({d_141_v30_: 685}))
            d_182_v58_ = nw18_
            d_183_v59_: C7
            nw19_ = C7()
            nw19_.ctor__(default__.fm0(((d_175_v51_)[not(False)] if (not(False)) in (d_175_v51_) else d_145_v34_), d_122_globalState_))
            d_183_v59_ = nw19_
        d_184_v60_: str
        d_184_v60_ = _dafny.CodePoint('p')
        d_108_v2_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghaefwn")) if d_106_v0_ else (d_108_v2_).set(default__.safeIndex(d_141_v30_, len(d_108_v2_)), d_184_v60_))
        if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwxqtd"))) + (d_108_v2_)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftqmi"))):
            d_139_v28_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrbedsljc"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_185_i5_ in range(default__.abs(888))]))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ck")))
            nw20_ = _dafny.Array(int(0), 6)
            d_111_v5_ = nw20_
            source5_ = D10_DC25(d_112_v6_, d_112_v6_, False, default__.fm0(d_112_v6_, d_122_globalState_))
            if source5_.is_DC25:
                d_186___mcc_h5_ = source5_.cf35
                d_187___mcc_h6_ = source5_.cf36
                d_188___mcc_h7_ = source5_.cf37
                d_189___mcc_h8_ = source5_.cf38
                d_190_cf38_ = d_189___mcc_h8_
                d_191_cf37_ = d_188___mcc_h7_
                d_192_cf36_ = d_187___mcc_h6_
                d_193_cf35_ = d_186___mcc_h5_
                d_194_v61_: _dafny.Array
                nw21_ = _dafny.Array(D11.default()(), 26)
                d_194_v61_ = nw21_
                d_194_v61_ = d_194_v61_
                d_195_v62_: _dafny.Map
                d_195_v62_ = _dafny.Map({d_138_v27_: d_190_cf38_})
                d_196_v63_: _dafny.Seq
                d_196_v63_ = _dafny.SeqWithoutIsStrInference([(d_117_v10_).set(len(d_195_v62_), len(default__.fm34(_dafny.SeqWithoutIsStrInference([d_143_v32_]), d_193_cf35_, d_122_globalState_)))])
                d_196_v63_ = (((d_196_v63_) + (d_196_v63_)) + ((_dafny.SeqWithoutIsStrInference([d_117_v10_ for d_197_i6_ in range(default__.abs(436))])) + (d_196_v63_))).set(default__.safeIndex((d_192_cf36_ if d_190_cf38_ else (_dafny.MultiSet([d_145_v34_, d_193_cf35_])).cardinality), len(((d_196_v63_) + (d_196_v63_)) + ((_dafny.SeqWithoutIsStrInference([d_117_v10_ for d_198_i6_ in range(default__.abs(436))])) + (d_196_v63_)))), d_117_v10_)
                d_184_v60_ = d_184_v60_
                d_111_v5_ = d_111_v5_
            elif True:
                d_199___mcc_h9_ = source5_.cf34
                d_200_cf34_ = d_199___mcc_h9_
                d_184_v60_ = _dafny.CodePoint('o')
                d_201_v64_: C0
                nw22_ = C0()
                nw22_.ctor__(d_109_v3_)
                d_201_v64_ = nw22_
                d_202_v65_: C9
                nw23_ = C9()
                nw23_.ctor__(d_119_v12_, len(d_108_v2_), d_117_v10_)
                d_202_v65_ = nw23_
                d_203_v66_: _dafny.Map
                d_203_v66_ = _dafny.Map({d_145_v34_: d_202_v65_})
                d_204_v67_: _dafny.Map
                d_204_v67_ = _dafny.Map({d_112_v6_: d_203_v66_})
                d_205_v68_: _dafny.Map
                d_205_v68_ = _dafny.Map({d_106_v0_: d_201_v64_})
                rhs13_ = d_112_v6_
                rhs14_ = d_184_v60_
                rhs15_ = len((d_204_v67_) | ((d_204_v67_) | (d_204_v67_)))
                rhs16_ = (805) >= (d_112_v6_)
                rhs17_ = ((d_205_v68_)[d_106_v0_] if (d_106_v0_) in (d_205_v68_) else d_201_v64_)
                lhs5_ = d_122_globalState_
                lhs6_ = d_122_globalState_
                lhs7_ = d_122_globalState_
                lhs5_.f16 = rhs13_
                d_184_v60_ = rhs14_
                lhs6_.f5 = rhs15_
                lhs7_.f13 = rhs16_
                d_201_v64_ = rhs17_
                d_206_v69_: _dafny.Array
                nw24_ = _dafny.Array(None, 11)
                d_206_v69_ = nw24_
                d_207_v70_: C10
                nw25_ = C10()
                nw25_.ctor__(d_112_v6_, True)
                d_207_v70_ = nw25_
                index6_ = default__.safeIndex(775, (d_206_v69_).length(0))
                (d_206_v69_)[index6_] = d_207_v70_
                index7_ = default__.safeIndex(775, (d_206_v69_).length(0))
                nw26_ = C10()
                nw26_.ctor__((d_202_v65_).f34, d_142_v31_)
                (d_206_v69_)[index7_] = nw26_
                index8_ = default__.safeIndex(641, (d_111_v5_).length(0))
                (d_111_v5_)[index8_] = d_112_v6_
                d_208_v71_: _dafny.MultiSet
                d_208_v71_ = _dafny.MultiSet([len(d_120_v13_), d_112_v6_])
                index9_ = default__.safeIndex(641, (d_111_v5_).length(0))
                (d_111_v5_)[index9_] = ((d_208_v71_).cardinality) + (default__.fm2(d_122_globalState_))
            d_209_v72_: bool
            d_210_v73_: bool
            d_211_v74_: _dafny.Map
            d_212_v75_: int
            out12_: bool
            out13_: bool
            out14_: _dafny.Map
            out15_: int
            out12_, out13_, out14_, out15_ = default__.m0(d_122_globalState_)
            d_209_v72_ = out12_
            d_210_v73_ = out13_
            d_211_v74_ = out14_
            d_212_v75_ = out15_
            d_115_v9_ = d_115_v9_
        elif True:
            (d_122_globalState_).f1 = d_142_v31_
            d_213_v76_: _dafny.Map
            d_213_v76_ = _dafny.Map({d_139_v28_: False})
            d_214_v77_: D20
            d_214_v77_ = D20_DC41(d_145_v34_)
            pat_let_tv0_ = d_151_v37_
            d_215_v78_: _dafny.Map
            def iife30_(_pat_let0_0):
                def iife31_(d_216_dt__update__tmp_h1_):
                    def iife32_(_pat_let1_0):
                        def iife33_(d_217_dt__update_hcf7_h0_):
                            return D3_DC6(d_217_dt__update_hcf7_h0_)
                        return iife33_(_pat_let1_0)
                    return iife32_(pat_let_tv0_)
                return iife31_(_pat_let0_0)
            d_215_v78_ = _dafny.Map({default__.safeDivisionInt(len(d_213_v76_), len(default__.fm7((d_214_v77_).cf57, iife30_(D3_DC6(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_142_v31_]): d_143_v32_}))), d_142_v31_, d_122_globalState_))): d_108_v2_})
            d_215_v78_ = (d_215_v78_).set(d_145_v34_, (d_108_v2_) + (d_108_v2_))
            d_218_v79_: C4
            nw27_ = C4()
            nw27_.ctor__()
            d_218_v79_ = nw27_
            d_219_v80_: _dafny.Set
            d_219_v80_ = _dafny.Set({d_141_v30_})
            d_220_v81_: _dafny.Seq
            d_221_v82_: _dafny.MultiSet
            d_222_v83_: int
            out16_: _dafny.Seq
            out17_: _dafny.MultiSet
            out18_: int
            out16_, out17_, out18_ = (d_218_v79_).m18(d_219_v80_, d_122_globalState_)
            d_220_v81_ = out16_
            d_221_v82_ = out17_
            d_222_v83_ = out18_
            (d_122_globalState_).f0 = d_141_v30_
        d_223_i7_: int
        d_223_i7_ = 0
        with _dafny.label("0"):
            while d_138_v27_:
                with _dafny.c_label("0"):
                    if (d_223_i7_) >= (100):
                        raise _dafny.Break("0")
                    d_223_i7_ = (d_223_i7_) + (1)
                    d_224_v84_: C6
                    nw28_ = C6()
                    nw28_.ctor__()
                    d_224_v84_ = nw28_
                    (d_122_globalState_).f5 = d_112_v6_
                    d_225_v85_: D5
                    d_225_v85_ = D5_DC12((d_120_v13_)[default__.safeIndex(289, len(d_120_v13_))], d_106_v0_)
                    d_226_v86_: C7
                    nw29_ = C7()
                    nw29_.ctor__((d_225_v85_).cf18)
                    d_226_v86_ = nw29_
                    d_226_v86_ = d_226_v86_
                    index10_ = default__.safeIndex(615, (d_111_v5_).length(0))
                    (d_111_v5_)[index10_] = len(d_108_v2_)
                    index11_ = default__.safeIndex(615, (d_111_v5_).length(0))
                    (d_111_v5_)[index11_] = d_112_v6_
                    pass
            pass
        (d_122_globalState_).f5 = d_112_v6_
        (d_122_globalState_).f5 = (default__.safeDivisionInt(len(_dafny.Set({d_112_v6_})), len(d_107_v1_)) if d_142_v31_ else d_141_v30_)
        index12_ = default__.safeIndex(885, (d_136_v25_).length(0))
        (d_136_v25_)[index12_] = _dafny.MultiSet([d_184_v60_, d_184_v60_])
        d_227_v87_: _dafny.MultiSet
        d_227_v87_ = _dafny.MultiSet([d_184_v60_])
        index13_ = default__.safeIndex(885, (d_136_v25_).length(0))
        (d_136_v25_)[index13_] = (d_227_v87_) - (d_227_v87_)
        _dafny.print(_dafny.string_of(d_106_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_107_v1_) == (_dafny.SeqWithoutIsStrInference([174]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_108_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_112_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_113_v7_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_v8_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jobjqyr"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v9_)[1]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_115_v9_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_117_v10_) == (_dafny.Map({766: 766}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_118_v11_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v12_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v13_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_121_v14_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_globalState_).f3) == (_dafny.SeqWithoutIsStrInference([174]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_122_globalState_).f7)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_122_globalState_).f8)[1]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_globalState_).f8)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_.f12) == (_dafny.Map({766: 766}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_122_globalState_).f14)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_122_globalState_).f18)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_122_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_globalState_).f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_123_v15_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v16_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_136_v25_)[10]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_137_v26_).cf6)[10]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_v27_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_139_v28_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v29_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqm"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_141_v30_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_142_v31_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_v32_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_144_v33_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqm"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_v34_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[0]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[1]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[2]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[3]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[4]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[5]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[6]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[7]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[8]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[9]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[10]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[11]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[12]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v36_)[13]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v37_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v39_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_184_v60_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_223_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v87_) == (_dafny.MultiSet([_dafny.CodePoint('p')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.Array(None, 0), _dafny.CodePoint('D'), int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC5(D2, NamedTuple('DC5', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), int(0), int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC7(D3, NamedTuple('DC7', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({self.cf8.VerbatimString(True)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC9(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)

class D4_DC9(D4, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC12(D5, NamedTuple('DC12', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC14(D6, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC16(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)

class D7_DC16(D7, NamedTuple('DC16', [('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC19(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)

class D8_DC19(D8, NamedTuple('DC19', [('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)

class D9_DC23(D9, NamedTuple('DC23', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC25(int(0), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)

class D10_DC25(D10, NamedTuple('DC25', [('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC24(D10, NamedTuple('DC24', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC27(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)

class D11_DC27(D11, NamedTuple('DC27', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC26(D11, NamedTuple('DC26', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC29()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)

class D12_DC29(D12, NamedTuple('DC29', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC30(D12, NamedTuple('DC30', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)

class D13_DC31(D13, NamedTuple('DC31', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)

class D14_DC32(D14, NamedTuple('DC32', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)

class D15_DC33(D15, NamedTuple('DC33', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D16_DC34)

class D16_DC34(D16, NamedTuple('DC34', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC34({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC34) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D17_DC35)

class D17_DC35(D17, NamedTuple('DC35', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC35({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC35) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D18_DC36)

class D18_DC36(D18, NamedTuple('DC36', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC36({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC36) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC38(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D19_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D19_DC37)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D19_DC39)

class D19_DC38(D19, NamedTuple('DC38', [('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC38({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC38) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC37(D19, NamedTuple('DC37', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC37({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC37) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC39(D19, NamedTuple('DC39', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC39({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC39) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC41(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D20_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D20_DC40)

class D20_DC41(D20, NamedTuple('DC41', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC41({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC41) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC40(D20, NamedTuple('DC40', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC40({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC40) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC43(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D21_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D21_DC42)

class D21_DC43(D21, NamedTuple('DC43', [('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC43({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC43) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC42(D21, NamedTuple('DC42', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC42({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC42) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D22_DC44)

class D22_DC44(D22, NamedTuple('DC44', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC44({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC44) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC46(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D23_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D23_DC45)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D23_DC47)

class D23_DC46(D23, NamedTuple('DC46', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC46({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC46) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC45(D23, NamedTuple('DC45', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC45({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC45) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC47(D23, NamedTuple('DC47', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC47({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC47) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D24_DC48)

class D24_DC48(D24, NamedTuple('DC48', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC48({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC48) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC50(False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D25_DC50)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D25_DC49)

class D25_DC50(D25, NamedTuple('DC50', [('cf67', Any), ('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC50({_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC50) and self.cf67 == __o.cf67 and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC49(D25, NamedTuple('DC49', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC49({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC49) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D26_DC51)

class D26_DC51(D26, NamedTuple('DC51', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC51({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC51) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC53(int(0), int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D27_DC53)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D27_DC52)

class D27_DC53(D27, NamedTuple('DC53', [('cf71', Any), ('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC53({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC53) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC52(D27, NamedTuple('DC52', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC52({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC52) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f26(self):
        return self._f26
    @f26.setter
    def f26(self, value):
        self._f26 = value
    def m1(self, p0, p1, p2, globalState):
        pass

    def m2(self, globalState):
        pass


class T1:
    pass
    def m7(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f1: bool = False
        self.f2: int = int(0)
        self.f5: int = int(0)
        self.f12: _dafny.Map = _dafny.Map({})
        self.f13: bool = False
        self.f16: int = int(0)
        self.f21: int = int(0)
        self.f22: int = int(0)
        self._f3: _dafny.Seq = _dafny.Seq({})
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f6: bool = False
        self._f7: _dafny.Map = _dafny.Map({})
        self._f8: _dafny.Array = _dafny.Array(None, 0)
        self._f9: int = int(0)
        self._f10: int = int(0)
        self._f11: bool = False
        self._f14: _dafny.Map = _dafny.Map({})
        self._f15: int = int(0)
        self._f17: int = int(0)
        self._f18: _dafny.Map = _dafny.Map({})
        self._f19: int = int(0)
        self._f20: bool = False
        self._f23: int = int(0)
        self._f24: int = int(0)
        self._f25: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25):
        (self).f0 = f0
        (self).f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self).f21 = f21
        (self).f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24
        (self)._f25 = f25

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
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    @property
    def f23(self):
        return self._f23
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25

class C0:
    def  __init__(self):
        self.f44: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f44):
        (self).f44 = f44

    def fm26(self, p0, p1, p2, globalState):
        return (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdupim"))))

    def fm27(self, p0, globalState):
        def iife34_():
            def iife36_():
                coll32_ = _dafny.Map()
                compr_32_: int
                for compr_32_ in (_dafny.Map({(_dafny.MultiSet([False, False, not(True), not(False), True])).cardinality: 8})).keys.Elements:
                    d_228_v1_: int = compr_32_
                    if (d_228_v1_) in (_dafny.Map({(_dafny.MultiSet([False, False, not(True), not(False), True])).cardinality: 8})):
                        coll32_[default__.safeDivisionInt(d_228_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpowsvsjo"))))] = not(not(True))
                return _dafny.Map(coll32_)
            coll30_ = _dafny.Map()
            def iife35_():
                coll31_ = _dafny.Map()
                compr_31_: int
                for compr_31_ in (_dafny.Map({(_dafny.MultiSet([False, False, not(True), not(False), True])).cardinality: 8})).keys.Elements:
                    d_228_v1_: int = compr_31_
                    if (d_228_v1_) in (_dafny.Map({(_dafny.MultiSet([False, False, not(True), not(False), True])).cardinality: 8})):
                        coll31_[default__.safeDivisionInt(d_228_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpowsvsjo"))))] = not(not(True))
                return _dafny.Map(coll31_)
            compr_30_: _dafny.Map
            for compr_30_ in (_dafny.SeqWithoutIsStrInference([iife35_()
            ])).Elements:
                d_229_v0_: _dafny.Map = compr_30_
                if (d_229_v0_) in (_dafny.SeqWithoutIsStrInference([iife36_()
                ])):
                    coll30_[d_229_v0_] = _dafny.Map({True: 506})
            return _dafny.Map(coll30_)
        def iife37_():
            coll33_ = _dafny.Map()
            compr_33_: _dafny.Map
            for compr_33_ in (_dafny.MultiSet([_dafny.Map({446: True}), _dafny.Map({-369: False})])).Elements:
                d_230_v2_: _dafny.Map = compr_33_
                if (d_230_v2_) in (_dafny.MultiSet([_dafny.Map({446: True}), _dafny.Map({-369: False})])):
                    coll33_[d_230_v2_] = _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjfblq")))})
            return _dafny.Map(coll33_)
        return (iife34_()
        ) | (((D9_DC22(iife37_()
)).cf33) | (_dafny.Map({_dafny.Map({514: True}): _dafny.Map({False: 681})})))


class C1(T1):
    def  __init__(self):
        self.f43: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f43):
        (self).f43 = f43

    def fm8(self, p0, globalState):
        return ((_dafny.Map({55: _dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([-856])).cardinality)]))]))})})) | (_dafny.Map({len(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_231_i0_ in range(default__.abs(504))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvheitv"))})): _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))})}))) | ((_dafny.Map({343: _dafny.Set({992})})) | (_dafny.Map({182: _dafny.Set({len(self.f43), -804})})))

    def fm25(self, globalState):
        return D2_DC5()

    def m7(self, p0, globalState):
        r0: bool = False
        d_232_v0_: bool
        d_232_v0_ = False
        d_233_v1_: _dafny.Seq
        d_233_v1_ = _dafny.SeqWithoutIsStrInference([d_232_v0_, d_232_v0_])
        d_234_v2_: _dafny.Map
        d_234_v2_ = _dafny.Map({d_233_v1_: d_232_v0_})
        d_235_v3_: D3
        d_235_v3_ = D3_DC6(d_234_v2_)
        d_235_v3_ = D3_DC6((d_234_v2_) | (d_234_v2_))
        rhs18_ = not((d_232_v0_ if d_232_v0_ else d_232_v0_))
        rhs19_ = p0
        lhs8_ = globalState
        r0 = rhs18_
        lhs8_.f22 = rhs19_
        d_236_v4_: _dafny.Array
        nw30_ = _dafny.Array(None, 14)
        nw30_[int(0)] = d_232_v0_
        nw30_[int(1)] = d_232_v0_
        nw30_[int(2)] = d_232_v0_
        nw30_[int(3)] = d_232_v0_
        nw30_[int(4)] = (d_233_v1_)[default__.safeIndex(p0, len(d_233_v1_))]
        nw30_[int(5)] = d_232_v0_
        nw30_[int(6)] = d_232_v0_
        nw30_[int(7)] = d_232_v0_
        nw30_[int(8)] = d_232_v0_
        nw30_[int(9)] = d_232_v0_
        nw30_[int(10)] = d_232_v0_
        nw30_[int(11)] = d_232_v0_
        nw30_[int(12)] = d_232_v0_
        nw30_[int(13)] = d_232_v0_
        d_236_v4_ = nw30_
        d_237_v5_: _dafny.Map
        d_237_v5_ = _dafny.Map({d_236_v4_: (_dafny.MultiSet([d_232_v0_])).set(d_232_v0_, default__.abs(p0))})
        hi2_ = default__.safeDivisionInt(default__.fm2(globalState), len(d_237_v5_))
        for d_238_i0_ in range(p0, hi2_):
            d_239_v6_: _dafny.Array
            def lambda2_(d_240_i0_):
                def lambda3_(d_241_i1_):
                    return (d_241_i1_) + (d_240_i0_)

                return lambda3_

            init1_ = lambda2_(d_238_i0_)
            nw31_ = _dafny.Array(None, 20)
            for i0_1_ in range(nw31_.length(0)):
                nw31_[i0_1_] = init1_(i0_1_)
            d_239_v6_ = nw31_
            index14_ = default__.safeIndex(815, (d_239_v6_).length(0))
            (d_239_v6_)[index14_] = (0) - (d_238_i0_)
            index15_ = default__.safeIndex(815, (d_239_v6_).length(0))
            (d_239_v6_)[index15_] = d_238_i0_
            d_236_v4_ = (d_236_v4_ if d_232_v0_ else d_236_v4_)
            (globalState).f0 = d_238_i0_
            d_242_v7_: _dafny.Seq
            d_242_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhkr"))
            d_242_v7_ = d_242_v7_
        hi3_ = p0
        for d_243_i2_ in range((0) - (len(d_233_v1_)), hi3_):
            d_236_v4_ = (d_236_v4_ if d_232_v0_ else d_236_v4_)
            d_232_v0_ = (p0) <= (p0)
            d_244_v8_: _dafny.Set
            d_244_v8_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptmp"))), d_243_i2_})
            d_245_v9_: D8
            d_245_v9_ = D8_DC18(d_244_v8_)
            d_232_v0_ = (len((d_245_v9_).cf27)) <= (default__.fm2(globalState))
            d_246_v10_: _dafny.MultiSet
            d_246_v10_ = _dafny.MultiSet([default__.safeModuloInt(p0, p0)])
            d_247_v11_: _dafny.MultiSet
            d_247_v11_ = _dafny.MultiSet([d_232_v0_])
            d_246_v10_ = _dafny.MultiSet([(d_247_v11_).cardinality, (0) - ((0) - (((0) - (d_243_i2_)) - (d_243_i2_))), d_243_i2_])
        d_248_v12_: _dafny.Map
        d_248_v12_ = _dafny.Map({len(d_233_v1_): p0})
        d_249_v13_: _dafny.Map
        d_249_v13_ = _dafny.Map({d_248_v12_: d_232_v0_})
        d_249_v13_ = (d_249_v13_).set((d_248_v12_) | (d_248_v12_), True)
        if True:
            d_250_v14_: _dafny.Array
            def lambda4_(d_251_i3_):
                return _dafny.CodePoint('a')

            init2_ = lambda4_
            nw32_ = _dafny.Array(None, 18)
            for i0_2_ in range(nw32_.length(0)):
                nw32_[i0_2_] = init2_(i0_2_)
            d_250_v14_ = nw32_
            d_252_v15_: C0
            nw33_ = C0()
            nw33_.ctor__(d_250_v14_)
            d_252_v15_ = nw33_
            (globalState).f16 = default__.fm2(globalState)
            d_253_v16_: _dafny.Array
            def lambda5_(d_254_i4_):
                return default__.safeDivisionInt(d_254_i4_, len(self.f43))

            init3_ = lambda5_
            nw34_ = _dafny.Array(None, 27)
            for i0_3_ in range(nw34_.length(0)):
                nw34_[i0_3_] = init3_(i0_3_)
            d_253_v16_ = nw34_
            index16_ = default__.safeIndex(46, (d_253_v16_).length(0))
            (d_253_v16_)[index16_] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxutpcxq"))), p0)
            index17_ = default__.safeIndex(46, (d_253_v16_).length(0))
            (d_253_v16_)[index17_] = p0
            d_255_v17_: D7
            d_255_v17_ = D7_DC16(353, d_232_v0_)
            pat_let_tv1_ = d_253_v16_
            pat_let_tv2_ = d_253_v16_
            pat_let_tv3_ = globalState
            d_256_v18_: _dafny.Set
            def iife38_(_pat_let2_0):
                def iife39_(d_257_dt__update__tmp_h0_):
                    def iife40_(_pat_let3_0):
                        def iife41_(d_258_dt__update_hcf23_h0_):
                            return D7_DC16((d_257_dt__update__tmp_h0_).cf22, d_258_dt__update_hcf23_h0_)
                        return iife41_(_pat_let3_0)
                    return iife40_(default__.fm0((pat_let_tv2_)[default__.safeIndex(46, (pat_let_tv1_).length(0))], pat_let_tv3_))
                return iife39_(_pat_let2_0)
            d_256_v18_ = _dafny.Set({d_255_v17_, iife38_(d_255_v17_), d_255_v17_, d_255_v17_})
            d_256_v18_ = d_256_v18_
            d_259_v19_: _dafny.Map
            d_259_v19_ = _dafny.Map({p0: not (d_232_v0_) or (not(d_232_v0_))})
            def iife42_():
                coll34_ = _dafny.Map()
                compr_34_: int
                for compr_34_ in (_dafny.Map({p0: (d_253_v16_)[default__.safeIndex(46, (d_253_v16_).length(0))]})).keys.Elements:
                    d_260_v20_: int = compr_34_
                    if (d_260_v20_) in (_dafny.Map({p0: (d_253_v16_)[default__.safeIndex(46, (d_253_v16_).length(0))]})):
                        coll34_[(d_260_v20_) - (len(_dafny.Map({p0: _dafny.CodePoint('x')})))] = d_232_v0_
                return _dafny.Map(coll34_)
            d_259_v19_ = iife42_()
            
        elif True:
            d_261_v21_: _dafny.Array
            nw35_ = _dafny.Array(int(0), 16)
            d_261_v21_ = nw35_
            index18_ = default__.safeIndex(642, (d_261_v21_).length(0))
            (d_261_v21_)[index18_] = p0
            d_262_v22_: _dafny.Map
            d_262_v22_ = _dafny.Map({p0: d_232_v0_})
            d_263_v23_: _dafny.MultiSet
            d_263_v23_ = _dafny.MultiSet([((d_262_v22_)[p0] if (p0) in (d_262_v22_) else d_232_v0_), d_232_v0_, d_232_v0_])
            d_264_v24_: _dafny.Map
            d_264_v24_ = _dafny.Map({len(d_248_v12_): default__.fm28(d_232_v0_, d_232_v0_, (d_263_v23_).cardinality, globalState)})
            d_265_v25_: D5
            d_265_v25_ = D5_DC13((0) - (len((d_233_v1_) + (((d_264_v24_)[-962] if (-962) in (d_264_v24_) else d_233_v1_)))))
            d_266_v26_: str
            d_266_v26_ = _dafny.CodePoint('h')
            d_267_v27_: _dafny.Seq
            d_267_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxamgsoh"))
            index19_ = default__.safeIndex(642, (d_261_v21_).length(0))
            rhs20_ = (p0 if d_232_v0_ else p0)
            rhs21_ = p0
            rhs22_ = d_265_v25_
            rhs23_ = d_266_v26_
            rhs24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pleyxhvvw"))
            lhs9_ = globalState
            lhs10_ = d_261_v21_
            lhs11_ = default__.safeIndex(642, (d_261_v21_).length(0))
            lhs9_.f5 = rhs20_
            lhs10_[lhs11_] = rhs21_
            d_265_v25_ = rhs22_
            d_266_v26_ = rhs23_
            d_267_v27_ = rhs24_
            d_268_v28_: _dafny.Seq
            d_268_v28_ = _dafny.SeqWithoutIsStrInference([(d_261_v21_)[default__.safeIndex(642, (d_261_v21_).length(0))]])
            d_269_v29_: _dafny.Array
            nw36_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_269_v29_ = nw36_
            index20_ = default__.safeIndex(266, (d_269_v29_).length(0))
            (d_269_v29_)[index20_] = d_261_v21_
            d_270_v30_: _dafny.Seq
            d_270_v30_ = _dafny.SeqWithoutIsStrInference([(d_268_v28_) + (d_268_v28_), d_268_v28_, _dafny.SeqWithoutIsStrInference([p0, p0, 46, (d_268_v28_)[default__.safeIndex((d_261_v21_)[default__.safeIndex(642, (d_261_v21_).length(0))], len(d_268_v28_))]]), d_268_v28_])
            d_271_v31_: _dafny.Map
            d_271_v31_ = _dafny.Map({(d_267_v27_) == (d_267_v27_): d_268_v28_})
            index21_ = default__.safeIndex(266, (d_269_v29_).length(0))
            rhs25_ = len((d_270_v30_)[default__.safeIndex((d_268_v28_)[default__.safeIndex(p0, len(d_268_v28_))], len(d_270_v30_))])
            rhs26_ = (d_232_v0_) == (not(d_232_v0_))
            rhs27_ = ((d_271_v31_)[False] if (False) in (d_271_v31_) else d_268_v28_)
            rhs28_ = d_261_v21_
            rhs29_ = (d_261_v21_)[default__.safeIndex(642, (d_261_v21_).length(0))]
            lhs12_ = globalState
            lhs13_ = d_269_v29_
            lhs14_ = default__.safeIndex(266, (d_269_v29_).length(0))
            lhs15_ = globalState
            lhs12_.f22 = rhs25_
            r0 = rhs26_
            d_268_v28_ = rhs27_
            lhs13_[lhs14_] = rhs28_
            lhs15_.f5 = rhs29_
            d_272_v32_: _dafny.Array
            nw37_ = _dafny.Array(None, 5)
            nw37_[int(0)] = d_266_v26_
            nw37_[int(1)] = d_266_v26_
            nw37_[int(2)] = d_266_v26_
            nw37_[int(3)] = _dafny.CodePoint('k')
            nw37_[int(4)] = d_266_v26_
            d_272_v32_ = nw37_
            d_273_v33_: C0
            nw38_ = C0()
            nw38_.ctor__(d_272_v32_)
            d_273_v33_ = nw38_
            (globalState).f2 = default__.safeDivisionInt(453, ((d_261_v21_)[default__.safeIndex(642, (d_261_v21_).length(0))]) - (-260))
            d_274_v34_: _dafny.Set
            d_274_v34_ = _dafny.Set({not(d_232_v0_)})
            d_268_v28_ = (default__.fm29(d_263_v23_, (d_268_v28_)[default__.safeIndex(len(d_274_v34_), len(d_268_v28_))], d_232_v0_, (d_261_v21_)[default__.safeIndex(642, (d_261_v21_).length(0))], globalState)) + (d_268_v28_)
        r0 = not (d_232_v0_) or ((d_233_v1_) != (_dafny.SeqWithoutIsStrInference([d_232_v0_, False, d_232_v0_, d_232_v0_])))
        return r0

    def m21(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        r3: bool = False
        d_275_v0_: bool
        d_275_v0_ = False
        d_276_v1_: int
        d_276_v1_ = -285
        if not (d_275_v0_) or (default__.fm0(d_276_v1_, globalState)):
            d_277_v2_: _dafny.Array
            def lambda6_(d_278_v1_):
                def lambda7_(d_279_i0_):
                    return (d_279_i0_) + (d_278_v1_)

                return lambda7_

            init4_ = lambda6_(d_276_v1_)
            nw39_ = _dafny.Array(None, 3)
            for i0_4_ in range(nw39_.length(0)):
                nw39_[i0_4_] = init4_(i0_4_)
            d_277_v2_ = nw39_
            d_280_v3_: _dafny.Seq
            d_280_v3_ = _dafny.SeqWithoutIsStrInference([d_275_v0_])
            index22_ = default__.safeIndex(772, (d_277_v2_).length(0))
            (d_277_v2_)[index22_] = len(d_280_v3_)
            d_281_v4_: _dafny.Seq
            d_281_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvpok"))
            d_282_v5_: _dafny.Map
            d_282_v5_ = _dafny.Map({d_280_v3_: d_275_v0_})
            d_283_v6_: D3
            d_283_v6_ = D3_DC6(d_282_v5_)
            d_284_v7_: _dafny.Map
            d_284_v7_ = _dafny.Map({(d_275_v0_) == (d_275_v0_): d_283_v6_})
            index23_ = default__.safeIndex(772, (d_277_v2_).length(0))
            rhs30_ = len(d_284_v7_)
            rhs31_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jftqty"))) + (d_281_v4_)
            lhs16_ = d_277_v2_
            lhs17_ = default__.safeIndex(772, (d_277_v2_).length(0))
            lhs16_[lhs17_] = rhs30_
            d_281_v4_ = rhs31_
            d_285_v8_: _dafny.Array
            def lambda8_(d_286_v4_):
                def lambda9_(d_287_i1_):
                    return d_286_v4_

                return lambda9_

            init5_ = lambda8_(d_281_v4_)
            nw40_ = _dafny.Array(None, 9)
            for i0_5_ in range(nw40_.length(0)):
                nw40_[i0_5_] = init5_(i0_5_)
            d_285_v8_ = nw40_
            def lambda10_(d_288_i2_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpfst"))

            init6_ = lambda10_
            nw41_ = _dafny.Array(None, 27)
            for i0_6_ in range(nw41_.length(0)):
                nw41_[i0_6_] = init6_(i0_6_)
            d_285_v8_ = nw41_
            d_289_v9_: _dafny.Array
            nw42_ = _dafny.Array(None, 25)
            nw42_[int(0)] = d_275_v0_
            nw42_[int(1)] = d_275_v0_
            nw42_[int(2)] = default__.fm0(d_276_v1_, globalState)
            nw42_[int(3)] = d_275_v0_
            nw42_[int(4)] = d_275_v0_
            nw42_[int(5)] = d_275_v0_
            nw42_[int(6)] = d_275_v0_
            nw42_[int(7)] = d_275_v0_
            nw42_[int(8)] = d_275_v0_
            nw42_[int(9)] = d_275_v0_
            nw42_[int(10)] = d_275_v0_
            nw42_[int(11)] = d_275_v0_
            nw42_[int(12)] = d_275_v0_
            nw42_[int(13)] = not(d_275_v0_)
            nw42_[int(14)] = d_275_v0_
            nw42_[int(15)] = d_275_v0_
            nw42_[int(16)] = d_275_v0_
            nw42_[int(17)] = default__.fm0(d_276_v1_, globalState)
            nw42_[int(18)] = d_275_v0_
            nw42_[int(19)] = d_275_v0_
            nw42_[int(20)] = not(default__.fm0((d_277_v2_)[default__.safeIndex(772, (d_277_v2_).length(0))], globalState))
            nw42_[int(21)] = d_275_v0_
            nw42_[int(22)] = False
            nw42_[int(23)] = d_275_v0_
            nw42_[int(24)] = d_275_v0_
            d_289_v9_ = nw42_
            d_290_v10_: _dafny.Map
            d_290_v10_ = _dafny.Map({d_275_v0_: d_289_v9_})
            d_291_v11_: _dafny.Map
            d_291_v11_ = _dafny.Map({d_275_v0_: d_290_v10_})
            d_292_v12_: _dafny.Array
            nw43_ = _dafny.Array(None, 7)
            nw43_[int(0)] = d_290_v10_
            nw43_[int(1)] = d_290_v10_
            nw43_[int(2)] = d_290_v10_
            nw43_[int(3)] = (d_290_v10_) | (d_290_v10_)
            nw43_[int(4)] = ((d_291_v11_)[d_275_v0_] if (d_275_v0_) in (d_291_v11_) else d_290_v10_)
            nw43_[int(5)] = (d_290_v10_) | (d_290_v10_)
            nw43_[int(6)] = d_290_v10_
            d_292_v12_ = nw43_
            index24_ = default__.safeIndex(618, (d_292_v12_).length(0))
            (d_292_v12_)[index24_] = d_290_v10_
            index25_ = default__.safeIndex(618, (d_292_v12_).length(0))
            rhs32_ = (d_281_v4_) + (d_281_v4_)
            rhs33_ = (_dafny.Map({d_275_v0_: d_289_v9_})) | (d_290_v10_)
            lhs18_ = d_292_v12_
            lhs19_ = default__.safeIndex(618, (d_292_v12_).length(0))
            d_281_v4_ = rhs32_
            lhs18_[lhs19_] = rhs33_
            d_293_v13_: _dafny.MultiSet
            d_293_v13_ = _dafny.MultiSet([838])
            d_294_v14_: _dafny.Map
            d_294_v14_ = _dafny.Map({d_293_v13_: (d_277_v2_)[default__.safeIndex(772, (d_277_v2_).length(0))]})
            (globalState).f21 = default__.safeModuloInt(len(d_280_v3_), ((d_294_v14_)[d_293_v13_] if (d_293_v13_) in (d_294_v14_) else d_276_v1_))
            d_295_v15_: _dafny.Map
            d_295_v15_ = _dafny.Map({d_289_v9_: d_277_v2_})
            d_295_v15_ = (_dafny.Map({d_289_v9_: d_277_v2_})) | (d_295_v15_)
        elif True:
            (globalState).f16 = (d_276_v1_ if d_275_v0_ else d_276_v1_)
            d_296_v16_: str
            d_296_v16_ = _dafny.CodePoint('p')
            r0 = d_296_v16_
            d_297_v17_: _dafny.Array
            nw44_ = _dafny.Array(int(0), 22)
            d_297_v17_ = nw44_
            d_297_v17_ = d_297_v17_
            d_298_v18_: _dafny.Map
            d_298_v18_ = _dafny.Map({default__.fm0(d_276_v1_, globalState): d_275_v0_})
            d_298_v18_ = (d_298_v18_).set((d_276_v1_) <= (-798), not(default__.fm0(len(_dafny.Set({d_276_v1_, d_276_v1_})), globalState)))
            (self).f43 = (self.f43).set(default__.fm0(d_276_v1_, globalState), d_276_v1_)
        d_299_v19_: _dafny.Set
        d_299_v19_ = _dafny.Set({d_275_v0_, d_275_v0_})
        d_300_i3_: int
        d_300_i3_ = 0
        with _dafny.label("1"):
            while (d_299_v19_).issubset((d_299_v19_) | (d_299_v19_)):
                with _dafny.c_label("1"):
                    if (d_300_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_300_i3_ = (d_300_i3_) + (1)
                    d_301_v20_: _dafny.Array
                    nw45_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                    d_301_v20_ = nw45_
                    d_302_v21_: C0
                    nw46_ = C0()
                    nw46_.ctor__(d_301_v20_)
                    d_302_v21_ = nw46_
                    d_303_v22_: _dafny.Map
                    d_303_v22_ = _dafny.Map({not(d_275_v0_): d_302_v21_})
                    d_303_v22_ = (d_303_v22_).set(not((D7_DC17(d_275_v0_, d_275_v0_, d_276_v1_)).cf25), d_302_v21_)
                    (globalState).f2 = d_276_v1_
                    d_304_v23_: str
                    d_304_v23_ = _dafny.CodePoint('l')
                    d_305_v25_: _dafny.MultiSet
                    d_305_v25_ = _dafny.MultiSet([d_276_v1_])
                    d_306_v27_: D8
                    def iife43_():
                        coll35_ = _dafny.Set()
                        compr_35_: int
                        for compr_35_ in (d_305_v25_).Elements:
                            d_307_v26_: int = compr_35_
                            if (d_307_v26_) in (d_305_v25_):
                                coll35_ = coll35_.union(_dafny.Set([default__.safeDivisionInt(d_307_v26_, (D7_DC16(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('h')])), False)).cf22)]))
                        return _dafny.Set(coll35_)
                    d_306_v27_ = D8_DC18(iife43_()
)
                    d_308_v28_: _dafny.Map
                    d_308_v28_ = _dafny.Map({d_306_v27_: d_275_v0_})
                    d_309_v29_: _dafny.Set
                    d_309_v29_ = _dafny.Set({(d_302_v21_).fm26(d_308_v28_, d_276_v1_, -24, globalState)})
                    d_310_v30_: _dafny.Map
                    d_310_v30_ = _dafny.Map({len(d_299_v19_): d_275_v0_})
                    d_311_v31_: D5
                    d_311_v31_ = D5_DC12(d_275_v0_, default__.fm0(len(d_310_v30_), globalState))
                    d_312_v32_: _dafny.Map
                    d_312_v32_ = _dafny.Map({d_275_v0_: (d_311_v31_).cf18})
                    d_313_v33_: _dafny.Map
                    d_313_v33_ = _dafny.Map({len(d_299_v19_): D9_DC23()})
                    d_314_v34_: _dafny.Seq
                    d_314_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmigbbu"))
                    d_315_v35_: _dafny.Map
                    d_315_v35_ = _dafny.Map({d_304_v23_: d_275_v0_})
                    d_316_v36_: _dafny.Array
                    nw47_ = _dafny.Array(None, 19)
                    nw47_[int(0)] = d_275_v0_
                    nw47_[int(1)] = d_275_v0_
                    nw47_[int(2)] = d_275_v0_
                    nw47_[int(3)] = d_275_v0_
                    nw47_[int(4)] = default__.fm0(d_276_v1_, globalState)
                    nw47_[int(5)] = d_275_v0_
                    nw47_[int(6)] = d_275_v0_
                    nw47_[int(7)] = d_275_v0_
                    nw47_[int(8)] = d_275_v0_
                    nw47_[int(9)] = d_275_v0_
                    nw47_[int(10)] = d_275_v0_
                    nw47_[int(11)] = d_275_v0_
                    nw47_[int(12)] = False
                    nw47_[int(13)] = d_275_v0_
                    nw47_[int(14)] = d_275_v0_
                    nw47_[int(15)] = d_275_v0_
                    nw47_[int(16)] = default__.fm0(len(d_315_v35_), globalState)
                    nw47_[int(17)] = True
                    nw47_[int(18)] = d_275_v0_
                    d_316_v36_ = nw47_
                    d_317_v37_: _dafny.Map
                    d_317_v37_ = _dafny.Map({d_316_v36_: 501})
                    d_318_v38_: _dafny.Map
                    d_318_v38_ = _dafny.Map({not(d_275_v0_): d_317_v37_})
                    d_319_v39_: _dafny.Map
                    d_319_v39_ = _dafny.Map({(d_302_v21_).fm26(d_308_v28_, len(d_313_v33_), len(d_314_v34_), globalState): ((d_318_v38_)[True] if (True) in (d_318_v38_) else d_317_v37_)})
                    d_320_v40_: _dafny.MultiSet
                    d_320_v40_ = _dafny.MultiSet([d_275_v0_])
                    d_321_v41_: _dafny.Map
                    d_321_v41_ = _dafny.Map({d_275_v0_: _dafny.Map({d_275_v0_: d_275_v0_})})
                    d_322_v42_: _dafny.Seq
                    d_322_v42_ = _dafny.SeqWithoutIsStrInference([d_275_v0_, not(d_275_v0_), not(d_275_v0_), d_275_v0_])
                    d_323_v43_: _dafny.Map
                    d_323_v43_ = _dafny.Map({d_276_v1_: d_322_v42_})
                    d_324_v44_: D7
                    d_324_v44_ = D7_DC17(d_275_v0_, d_275_v0_, d_276_v1_)
                    d_325_v45_: _dafny.Array
                    nw48_ = _dafny.Array(None, 29)
                    nw48_[int(0)] = 612
                    def iife44_():
                        coll36_ = _dafny.Set()
                        compr_36_: str
                        for compr_36_ in (_dafny.MultiSet([d_304_v23_])).Elements:
                            d_326_v24_: str = compr_36_
                            if (d_326_v24_) in (_dafny.MultiSet([d_304_v23_])):
                                coll36_ = coll36_.union(_dafny.Set([d_326_v24_]))
                        return _dafny.Set(coll36_)
                    nw48_[int(1)] = default__.safeModuloInt(len(iife44_()
                    ), d_276_v1_)
                    nw48_[int(2)] = d_276_v1_
                    nw48_[int(3)] = d_276_v1_
                    nw48_[int(4)] = len(d_309_v29_)
                    nw48_[int(5)] = d_276_v1_
                    nw48_[int(6)] = len(d_299_v19_)
                    nw48_[int(7)] = d_276_v1_
                    nw48_[int(8)] = default__.safeModuloInt(len((d_312_v32_).set(d_275_v0_, d_275_v0_)), d_276_v1_)
                    nw48_[int(9)] = -764
                    nw48_[int(10)] = (d_276_v1_) - (d_276_v1_)
                    nw48_[int(11)] = (d_302_v21_).fm26(d_308_v28_, len(((d_319_v39_)[d_276_v1_] if (d_276_v1_) in (d_319_v39_) else d_317_v37_)), d_276_v1_, globalState)
                    nw48_[int(12)] = (d_320_v40_).cardinality
                    nw48_[int(13)] = d_276_v1_
                    nw48_[int(14)] = len(d_321_v41_)
                    nw48_[int(15)] = d_276_v1_
                    nw48_[int(16)] = len(((d_322_v42_) + (((d_323_v43_)[d_276_v1_] if (d_276_v1_) in (d_323_v43_) else d_322_v42_))).set(default__.safeIndex(d_276_v1_, len((d_322_v42_) + (((d_323_v43_)[d_276_v1_] if (d_276_v1_) in (d_323_v43_) else d_322_v42_)))), True))
                    nw48_[int(17)] = (d_276_v1_) - ((0) - (len(d_322_v42_)))
                    nw48_[int(18)] = default__.fm2(globalState)
                    nw48_[int(19)] = 77
                    nw48_[int(20)] = d_276_v1_
                    nw48_[int(21)] = d_276_v1_
                    nw48_[int(22)] = len(d_314_v34_)
                    nw48_[int(23)] = d_276_v1_
                    nw48_[int(24)] = 242
                    nw48_[int(25)] = d_276_v1_
                    nw48_[int(26)] = len(_dafny.SeqWithoutIsStrInference([d_276_v1_]))
                    nw48_[int(27)] = (d_324_v44_).cf26
                    nw48_[int(28)] = d_276_v1_
                    d_325_v45_ = nw48_
                    index26_ = default__.safeIndex(809, (d_325_v45_).length(0))
                    (d_325_v45_)[index26_] = len(d_314_v34_)
                    index27_ = default__.safeIndex(809, (d_325_v45_).length(0))
                    (d_325_v45_)[index27_] = d_276_v1_
                    pass
            pass
        d_327_v46_: _dafny.Array
        def lambda11_(d_328_i4_):
            return _dafny.CodePoint('v')

        init7_ = lambda11_
        nw49_ = _dafny.Array(None, 27)
        for i0_7_ in range(nw49_.length(0)):
            nw49_[i0_7_] = init7_(i0_7_)
        d_327_v46_ = nw49_
        d_329_v47_: C0
        nw50_ = C0()
        nw50_.ctor__(d_327_v46_)
        d_329_v47_ = nw50_
        d_330_v48_: D4
        d_330_v48_ = D4_DC9(d_276_v1_)
        source6_ = d_330_v48_
        if source6_.is_DC9:
            d_331___mcc_h0_ = source6_.cf14
            d_332_cf14_ = d_331___mcc_h0_
            d_333_v49_: _dafny.Array
            nw51_ = _dafny.Array(int(0), 10)
            d_333_v49_ = nw51_
            index28_ = default__.safeIndex(844, (d_333_v49_).length(0))
            (d_333_v49_)[index28_] = d_332_cf14_
            index29_ = default__.safeIndex(844, (d_333_v49_).length(0))
            (d_333_v49_)[index29_] = (0) - (d_276_v1_)
            if True:
                (globalState).f13 = (d_276_v1_) > (d_332_cf14_)
                r2 = d_275_v0_
                d_334_v50_: _dafny.Seq
                d_334_v50_ = _dafny.SeqWithoutIsStrInference([(d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))], (d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))], default__.safeModuloInt(d_276_v1_, d_332_cf14_)])
                d_334_v50_ = ((d_334_v50_ if d_275_v0_ else d_334_v50_)) + ((d_334_v50_).set(default__.safeIndex(d_276_v1_, len(d_334_v50_)), (0) - (d_276_v1_)))
                d_335_v51_: C0
                nw52_ = C0()
                nw52_.ctor__(d_329_v47_.f44)
                d_335_v51_ = nw52_
                d_336_v52_: D5
                d_336_v52_ = D5_DC12(d_275_v0_, d_275_v0_)
                d_337_v53_: _dafny.MultiSet
                d_337_v53_ = _dafny.MultiSet([d_336_v52_, d_336_v52_])
                def iife45_():
                    coll37_ = _dafny.Set()
                    compr_37_: D5
                    for compr_37_ in (d_337_v53_).Elements:
                        d_338_v54_: D5 = compr_37_
                        if (d_338_v54_) in (d_337_v53_):
                            coll37_ = coll37_.union(_dafny.Set([d_338_v54_]))
                    return _dafny.Set(coll37_)
                d_275_v0_ = not((iife45_()
                ).issubset(_dafny.Set({d_336_v52_, d_336_v52_, d_336_v52_})))
            elif True:
                d_333_v49_ = d_333_v49_
                index30_ = default__.safeIndex(844, (d_333_v49_).length(0))
                (d_333_v49_)[index30_] = (d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))]
                d_333_v49_ = d_333_v49_
                (globalState).f13 = True
                d_339_v55_: str
                d_339_v55_ = _dafny.CodePoint('b')
                r0 = (d_339_v55_ if d_275_v0_ else _dafny.CodePoint('p'))
            d_340_v56_: _dafny.Array
            def lambda12_(d_341_i5_):
                return True

            init8_ = lambda12_
            nw53_ = _dafny.Array(None, 25)
            for i0_8_ in range(nw53_.length(0)):
                nw53_[i0_8_] = init8_(i0_8_)
            d_340_v56_ = nw53_
            index31_ = default__.safeIndex(866, (d_340_v56_).length(0))
            (d_340_v56_)[index31_] = (d_275_v0_) and (d_275_v0_)
            d_342_v57_: _dafny.Map
            d_342_v57_ = _dafny.Map({-106: (d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))]})
            d_343_v58_: _dafny.MultiSet
            d_343_v58_ = _dafny.MultiSet([len(d_342_v57_), d_332_cf14_])
            d_344_v59_: _dafny.Map
            d_344_v59_ = _dafny.Map({(d_343_v58_).cardinality: d_275_v0_})
            d_345_v60_: _dafny.Map
            d_345_v60_ = _dafny.Map({d_276_v1_: len(d_344_v59_)})
            index32_ = default__.safeIndex(866, (d_340_v56_).length(0))
            (d_340_v56_)[index32_] = (d_332_cf14_) < (((d_345_v60_)[(d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))]] if ((d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))]) in (d_345_v60_) else (d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))]))
            if not (d_275_v0_) or (d_275_v0_):
                d_346_v61_: _dafny.Seq
                d_346_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtgjmrwhb"))
                d_347_v62_: _dafny.Map
                d_347_v62_ = _dafny.Map({default__.safeDivisionInt(d_332_cf14_, len(_dafny.Map({d_346_v61_: len(_dafny.SeqWithoutIsStrInference([d_332_cf14_]))}))): d_346_v61_})
                d_347_v62_ = d_347_v62_
                d_348_v63_: _dafny.Array
                nw54_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_348_v63_ = nw54_
                d_349_v64_: _dafny.MultiSet
                d_349_v64_ = _dafny.MultiSet([(d_340_v56_)[default__.safeIndex(866, (d_340_v56_).length(0))], False])
                rhs34_ = default__.safeModuloInt(len(default__.fm30(d_332_cf14_, d_346_v61_, (d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))], globalState)), default__.safeDivisionInt(d_332_cf14_, (d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))]))
                rhs35_ = d_348_v63_
                rhs36_ = (d_349_v64_).issubset(default__.fm31(globalState))
                lhs20_ = globalState
                lhs21_ = globalState
                lhs20_.f5 = rhs34_
                d_348_v63_ = rhs35_
                lhs21_.f1 = rhs36_
                d_350_v65_: _dafny.Map
                d_350_v65_ = _dafny.Map({d_346_v61_: d_329_v47_.f44})
                d_351_v66_: C0
                nw55_ = C0()
                nw55_.ctor__(((d_350_v65_)[d_346_v61_] if (d_346_v61_) in (d_350_v65_) else d_329_v47_.f44))
                d_351_v66_ = nw55_
                d_346_v61_ = d_346_v61_
                d_352_v67_: _dafny.Seq
                d_352_v67_ = _dafny.SeqWithoutIsStrInference([(d_340_v56_)[default__.safeIndex(866, (d_340_v56_).length(0))]])
                d_353_v68_: _dafny.Map
                d_353_v68_ = _dafny.Map({True: d_352_v67_})
                index33_ = default__.safeIndex(844, (d_333_v49_).length(0))
                (d_333_v49_)[index33_] = len((d_352_v67_) + (((d_353_v68_)[d_275_v0_] if (d_275_v0_) in (d_353_v68_) else _dafny.SeqWithoutIsStrInference([d_275_v0_]))))
            elif True:
                r2 = d_275_v0_
                d_340_v56_ = d_340_v56_
                d_354_v69_: _dafny.Seq
                d_354_v69_ = _dafny.SeqWithoutIsStrInference([(d_340_v56_)[default__.safeIndex(866, (d_340_v56_).length(0))]])
                index34_ = default__.safeIndex(866, (d_340_v56_).length(0))
                (d_340_v56_)[index34_] = ((d_340_v56_)[default__.safeIndex(866, (d_340_v56_).length(0))] if False else (d_354_v69_)[default__.safeIndex(-248, len(d_354_v69_))])
                pat_let_tv4_ = d_332_cf14_
                def iife46_(_pat_let4_0):
                    def iife47_(d_355_dt__update__tmp_h0_):
                        def iife48_(_pat_let5_0):
                            def iife49_(d_356_dt__update_hcf19_h0_):
                                return D5_DC13(d_356_dt__update_hcf19_h0_)
                            return iife49_(_pat_let5_0)
                        return iife48_(pat_let_tv4_)
                    return iife47_(_pat_let4_0)
                (globalState).f2 = (0) - (default__.safeDivisionInt((d_333_v49_)[default__.safeIndex(844, (d_333_v49_).length(0))], (iife46_(D5_DC13(d_332_cf14_))).cf19))
                index35_ = default__.safeIndex(866, (d_340_v56_).length(0))
                (d_340_v56_)[index35_] = not(not(d_275_v0_))
        elif source6_.is_DC10:
            d_357___mcc_h1_ = source6_.cf15
            d_358_cf15_ = d_357___mcc_h1_
            d_359_v70_: _dafny.Seq
            d_359_v70_ = _dafny.SeqWithoutIsStrInference([d_358_cf15_ for d_360_i6_ in range(default__.abs(320))])
            d_361_v71_: _dafny.Map
            d_361_v71_ = _dafny.Map({False: (_dafny.SeqWithoutIsStrInference([d_358_cf15_]))})
            d_362_v72_: _dafny.Seq
            d_362_v72_ = _dafny.SeqWithoutIsStrInference([d_276_v1_])
            d_361_v71_ = (d_361_v71_).set(d_275_v0_, (d_362_v72_) + (((d_361_v71_)[d_275_v0_] if (d_275_v0_) in (d_361_v71_) else d_362_v72_)))
            d_363_v73_: _dafny.Map
            d_363_v73_ = _dafny.Map({d_358_cf15_: d_358_cf15_})
            d_364_v75_: _dafny.MultiSet
            d_364_v75_ = _dafny.MultiSet([-86])
            d_365_v76_: _dafny.MultiSet
            def iife50_():
                coll38_ = _dafny.Map()
                compr_38_: int
                for compr_38_ in (d_364_v75_).Elements:
                    d_366_v74_: int = compr_38_
                    if (d_366_v74_) in (d_364_v75_):
                        coll38_[default__.safeModuloInt(d_366_v74_, len(_dafny.Map({d_275_v0_: d_275_v0_})))] = d_358_cf15_
                return _dafny.Map(coll38_)
            d_365_v76_ = _dafny.MultiSet([d_363_v73_, d_363_v73_, d_363_v73_, d_363_v73_, iife50_()
            ])
            d_367_v77_: str
            d_367_v77_ = _dafny.CodePoint('y')
            d_368_v78_: _dafny.MultiSet
            d_368_v78_ = _dafny.MultiSet([d_367_v77_])
            d_369_v79_: _dafny.Seq
            d_369_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdchfnll"))
            (globalState).f21 = (((d_365_v76_)[d_363_v73_] if (d_363_v73_) in (d_365_v76_) else ((d_368_v78_)[d_367_v77_] if (d_367_v77_) in (d_368_v78_) else 686)) if d_275_v0_ else (len(d_369_v79_)) + (d_358_cf15_))
            (globalState).f22 = -23
            d_370_v80_: _dafny.Set
            d_370_v80_ = _dafny.Set({d_276_v1_, d_276_v1_})
            d_371_v81_: D8
            d_371_v81_ = D8_DC18(d_370_v80_)
            d_371_v81_ = d_371_v81_
        elif True:
            d_372___mcc_h2_ = source6_.cf13
            d_373_cf13_ = d_372___mcc_h2_
            d_276_v1_ = d_276_v1_
            d_374_v82_: _dafny.Array
            def lambda13_(d_375_v1_):
                def lambda14_(d_376_i7_):
                    return default__.safeModuloInt(d_376_i7_, d_375_v1_)

                return lambda14_

            init9_ = lambda13_(d_276_v1_)
            nw56_ = _dafny.Array(None, 14)
            for i0_9_ in range(nw56_.length(0)):
                nw56_[i0_9_] = init9_(i0_9_)
            d_374_v82_ = nw56_
            d_377_v83_: _dafny.Seq
            d_377_v83_ = _dafny.SeqWithoutIsStrInference([d_374_v82_, d_374_v82_, d_374_v82_])
            d_378_v84_: _dafny.Seq
            d_378_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htmvxaccj"))
            d_379_v85_: _dafny.Map
            d_379_v85_ = _dafny.Map({d_275_v0_: d_378_v84_})
            d_374_v82_ = (d_377_v83_)[default__.safeIndex(len(d_379_v85_), len(d_377_v83_))]
            d_380_v86_: _dafny.Array
            nw57_ = _dafny.Array(_dafny.Seq({}), 12)
            d_380_v86_ = nw57_
            d_381_v87_: _dafny.Seq
            d_381_v87_ = _dafny.SeqWithoutIsStrInference([d_275_v0_])
            index36_ = default__.safeIndex(657, (d_380_v86_).length(0))
            (d_380_v86_)[index36_] = (d_381_v87_ if d_275_v0_ else d_381_v87_)
            index37_ = default__.safeIndex(657, (d_380_v86_).length(0))
            (d_380_v86_)[index37_] = d_381_v87_
            r0 = d_373_cf13_
        hi4_ = d_276_v1_
        for d_382_i8_ in range(d_276_v1_, hi4_):
            d_383_v88_: _dafny.Map
            d_383_v88_ = _dafny.Map({not(d_275_v0_): d_275_v0_})
            d_384_v89_: _dafny.MultiSet
            d_384_v89_ = _dafny.MultiSet([len(d_383_v88_), d_276_v1_, d_382_i8_])
            d_385_v90_: _dafny.Seq
            d_385_v90_ = _dafny.SeqWithoutIsStrInference([d_275_v0_, d_275_v0_, (d_384_v89_).isdisjoint(d_384_v89_)])
            d_386_v91_: _dafny.Map
            d_386_v91_ = _dafny.Map({d_275_v0_: (d_385_v90_) + (d_385_v90_)})
            d_385_v90_ = ((d_386_v91_)[d_275_v0_] if (d_275_v0_) in (d_386_v91_) else d_385_v90_)
            d_383_v88_ = (d_383_v88_).set(d_275_v0_, False)
            d_387_v92_: C0
            nw58_ = C0()
            nw58_.ctor__(d_327_v46_)
            d_387_v92_ = nw58_
            if d_275_v0_:
                (globalState).f0 = d_276_v1_
                (globalState).f2 = (d_382_i8_) * (d_276_v1_)
                (globalState).f16 = default__.fm2(globalState)
                r3 = not(d_275_v0_)
                (globalState).f0 = default__.safeModuloInt(d_382_i8_, d_276_v1_)
            elif True:
                (globalState).f13 = d_275_v0_
                d_388_v93_: str
                d_388_v93_ = _dafny.CodePoint('m')
                d_389_v94_: _dafny.MultiSet
                d_389_v94_ = _dafny.MultiSet([d_388_v93_])
                d_390_v95_: _dafny.Map
                d_390_v95_ = _dafny.Map({d_276_v1_: (d_389_v94_).cardinality})
                d_391_v96_: _dafny.MultiSet
                d_391_v96_ = _dafny.MultiSet([d_390_v95_])
                (self).f43 = (self.f43).set(((d_391_v96_).cardinality) >= (d_276_v1_), d_276_v1_)
                (globalState).f21 = default__.safeDivisionInt(d_382_i8_, default__.safeDivisionInt(d_276_v1_, 281))
                d_392_v97_: _dafny.Map
                d_392_v97_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_276_v1_ for d_393_i9_ in range(default__.abs(926))]): d_275_v0_})
                d_394_v98_: _dafny.Seq
                d_394_v98_ = _dafny.SeqWithoutIsStrInference([d_276_v1_, d_382_i8_])
                d_395_v99_: _dafny.Array
                def lambda15_(d_396_v1_):
                    def lambda16_(d_397_i10_):
                        return default__.safeModuloInt(d_397_i10_, d_396_v1_)

                    return lambda16_

                init10_ = lambda15_(d_276_v1_)
                nw59_ = _dafny.Array(None, 19)
                for i0_10_ in range(nw59_.length(0)):
                    nw59_[i0_10_] = init10_(i0_10_)
                d_395_v99_ = nw59_
                d_398_v100_: _dafny.Map
                d_398_v100_ = _dafny.Map({d_395_v99_: d_382_i8_})
                d_399_v101_: _dafny.Map
                d_399_v101_ = d_398_v100_
                d_400_v102_: _dafny.Map
                d_400_v102_ = _dafny.Map({d_394_v98_: d_399_v101_})
                d_401_v103_: _dafny.Map
                d_401_v103_ = _dafny.Map({False: (d_400_v102_).set(d_394_v98_, d_399_v101_)})
                d_402_v104_: _dafny.Map
                d_402_v104_ = _dafny.Map({d_392_v97_: d_401_v103_})
                d_403_v105_: _dafny.Array
                nw60_ = _dafny.Array(False, 12)
                d_403_v105_ = nw60_
                d_404_v106_: _dafny.Map
                d_404_v106_ = _dafny.Map({d_403_v105_: d_382_i8_})
                d_405_v107_: _dafny.Seq
                d_405_v107_ = _dafny.SeqWithoutIsStrInference([len(d_404_v106_)])
                d_402_v104_ = (d_402_v104_).set((((d_392_v97_).set(_dafny.SeqWithoutIsStrInference([d_276_v1_]), (d_385_v90_)[default__.safeIndex(d_382_i8_, len(d_385_v90_))])).set(d_405_v107_, d_275_v0_)) | (d_392_v97_), (d_401_v103_) | (d_401_v103_))
                r2 = d_275_v0_
        d_275_v0_ = (D8_DC19(d_275_v0_, d_275_v0_, d_275_v0_)).cf28
        d_406_v108_: str
        d_406_v108_ = _dafny.CodePoint('t')
        r0 = d_406_v108_
        d_407_v109_: _dafny.Array
        nw61_ = _dafny.Array(D4.default()(), 1)
        d_407_v109_ = nw61_
        r1 = d_407_v109_
        d_408_v110_: _dafny.Array
        def lambda17_(d_409_v0_):
            def lambda18_(d_410_i11_):
                return d_409_v0_

            return lambda18_

        init11_ = lambda17_(d_275_v0_)
        nw62_ = _dafny.Array(None, 2)
        for i0_11_ in range(nw62_.length(0)):
            nw62_[i0_11_] = init11_(i0_11_)
        d_408_v110_ = nw62_
        d_411_v111_: _dafny.Seq
        d_411_v111_ = _dafny.SeqWithoutIsStrInference([d_408_v110_, d_408_v110_])
        d_412_v112_: _dafny.Map
        d_412_v112_ = _dafny.Map({(d_411_v111_)[default__.safeIndex(d_276_v1_, len(d_411_v111_))]: d_276_v1_})
        r2 = (d_408_v110_) not in (d_412_v112_)
        d_413_v113_: D8
        d_413_v113_ = D8_DC20(d_275_v0_)
        pat_let_tv5_ = d_276_v1_
        pat_let_tv6_ = globalState
        def iife51_(_pat_let6_0):
            def iife52_(d_414_dt__update__tmp_h2_):
                def iife53_(_pat_let7_0):
                    def iife54_(d_415_dt__update_hcf31_h0_):
                        return D8_DC20(d_415_dt__update_hcf31_h0_)
                    return iife54_(_pat_let7_0)
                return iife53_(default__.fm0(pat_let_tv5_, pat_let_tv6_))
            return iife52_(_pat_let6_0)
        r3 = (iife51_(d_413_v113_)).cf31
        return r0, r1, r2, r3


class C2:
    def  __init__(self):
        self._f41: bool = False
        self._f42: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f41, f42):
        (self)._f41 = f41
        (self)._f42 = f42

    def fm24(self, globalState):
        return (((0) - ((0) - ((self).f42))) + ((0) - (len(_dafny.Set({(self).f41}))))) != ((self).f42)

    def m19(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_416_v0_: _dafny.Array
        def lambda19_(d_417_i0_):
            return (d_417_i0_) - ((self).f42)

        init12_ = lambda19_
        nw63_ = _dafny.Array(None, 8)
        for i0_12_ in range(nw63_.length(0)):
            nw63_[i0_12_] = init12_(i0_12_)
        d_416_v0_ = nw63_
        index38_ = default__.safeIndex(746, (d_416_v0_).length(0))
        (d_416_v0_)[index38_] = p0
        index39_ = default__.safeIndex(746, (d_416_v0_).length(0))
        (d_416_v0_)[index39_] = (self).f42
        d_418_i1_: int
        d_418_i1_ = 0
        with _dafny.label("2"):
            while not ((self).f41) or ((self).f41):
                with _dafny.c_label("2"):
                    if (d_418_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_418_i1_ = (d_418_i1_) + (1)
                    d_419_v1_: _dafny.Seq
                    d_419_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwpw"))
                    d_420_v2_: _dafny.Array
                    nw64_ = _dafny.Array(False, 3)
                    d_420_v2_ = nw64_
                    d_421_v3_: _dafny.Map
                    d_421_v3_ = _dafny.Map({d_419_v1_: d_420_v2_})
                    d_421_v3_ = (d_421_v3_).set(d_419_v1_, (p1)[default__.safeIndex((d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], len(p1))])
                    d_422_v4_: str
                    d_422_v4_ = _dafny.CodePoint('m')
                    d_423_v5_: _dafny.Array
                    nw65_ = _dafny.Array(None, 28)
                    nw65_[int(0)] = d_422_v4_
                    nw65_[int(1)] = d_422_v4_
                    nw65_[int(2)] = d_422_v4_
                    nw65_[int(3)] = d_422_v4_
                    nw65_[int(4)] = (default__.fm32((self).f41, globalState)).cf13
                    nw65_[int(5)] = d_422_v4_
                    nw65_[int(6)] = (d_419_v1_)[default__.safeIndex((d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], len(d_419_v1_))]
                    nw65_[int(7)] = d_422_v4_
                    nw65_[int(8)] = d_422_v4_
                    nw65_[int(9)] = d_422_v4_
                    nw65_[int(10)] = d_422_v4_
                    nw65_[int(11)] = d_422_v4_
                    nw65_[int(12)] = d_422_v4_
                    nw65_[int(13)] = d_422_v4_
                    nw65_[int(14)] = d_422_v4_
                    nw65_[int(15)] = d_422_v4_
                    nw65_[int(16)] = _dafny.CodePoint('p')
                    nw65_[int(17)] = _dafny.CodePoint('e')
                    nw65_[int(18)] = d_422_v4_
                    nw65_[int(19)] = d_422_v4_
                    nw65_[int(20)] = d_422_v4_
                    nw65_[int(21)] = (d_419_v1_)[default__.safeIndex(p0, len(d_419_v1_))]
                    nw65_[int(22)] = d_422_v4_
                    nw65_[int(23)] = d_422_v4_
                    nw65_[int(24)] = _dafny.CodePoint('q')
                    nw65_[int(25)] = d_422_v4_
                    nw65_[int(26)] = d_422_v4_
                    nw65_[int(27)] = (d_419_v1_)[default__.safeIndex((d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], len(d_419_v1_))]
                    d_423_v5_ = nw65_
                    d_424_v6_: C0
                    nw66_ = C0()
                    nw66_.ctor__(d_423_v5_)
                    d_424_v6_ = nw66_
                    d_425_v7_: _dafny.Map
                    d_425_v7_ = _dafny.Map({(0) - ((self).f42): (self).f41})
                    (globalState).f13 = not ((d_425_v7_) != (d_425_v7_)) or ((self).f41)
                    d_426_v8_: _dafny.Array
                    nw67_ = _dafny.Array(_dafny.Map({}), 13)
                    d_426_v8_ = nw67_
                    d_427_v9_: _dafny.Set
                    d_427_v9_ = _dafny.Set({p0})
                    d_428_v10_: _dafny.Map
                    d_428_v10_ = _dafny.Map({d_427_v9_: (self).f41})
                    index40_ = default__.safeIndex(960, (d_426_v8_).length(0))
                    (d_426_v8_)[index40_] = d_428_v10_
                    index41_ = default__.safeIndex(960, (d_426_v8_).length(0))
                    (d_426_v8_)[index41_] = _dafny.Map({d_427_v9_: (self).f41})
                    pass
            pass
        d_429_v11_: _dafny.Seq
        d_429_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywiqjxetm"))
        if (d_429_v11_) == (d_429_v11_):
            d_430_v12_: _dafny.Set
            d_430_v12_ = _dafny.Set({(self).f41})
            d_431_v13_: D10
            d_431_v13_ = D10_DC24(d_430_v12_)
            d_432_v14_: D5
            d_432_v14_ = D5_DC12((self).f41, (self).f41)
            d_433_v15_: _dafny.Array
            nw68_ = _dafny.Array(None, 5)
            nw68_[int(0)] = (d_430_v12_).ispropersubset((d_431_v13_).cf34)
            nw68_[int(1)] = (p0) != (p0)
            nw68_[int(2)] = (self).f41
            nw68_[int(3)] = (self).f41
            nw68_[int(4)] = (d_432_v14_).cf17
            d_433_v15_ = nw68_
            index42_ = default__.safeIndex(74, (d_433_v15_).length(0))
            (d_433_v15_)[index42_] = (self).f41
            index43_ = default__.safeIndex(74, (d_433_v15_).length(0))
            (d_433_v15_)[index43_] = (self).f41
            (globalState).f1 = (d_433_v15_)[default__.safeIndex(74, (d_433_v15_).length(0))]
            index44_ = default__.safeIndex(74, (d_433_v15_).length(0))
            (d_433_v15_)[index44_] = (d_433_v15_)[default__.safeIndex(74, (d_433_v15_).length(0))]
            d_434_v16_: _dafny.Array
            def lambda20_(d_435_v0_):
                def lambda21_(d_436_i2_):
                    return _dafny.SeqWithoutIsStrInference([(d_435_v0_)[default__.safeIndex(746, (d_435_v0_).length(0))]])

                return lambda21_

            init13_ = lambda20_(d_416_v0_)
            nw69_ = _dafny.Array(None, 4)
            for i0_13_ in range(nw69_.length(0)):
                nw69_[i0_13_] = init13_(i0_13_)
            d_434_v16_ = nw69_
            index45_ = default__.safeIndex(618, (d_434_v16_).length(0))
            (d_434_v16_)[index45_] = _dafny.SeqWithoutIsStrInference([163])
            d_437_v17_: _dafny.Seq
            d_437_v17_ = _dafny.SeqWithoutIsStrInference([(self).f42, (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]])
            d_438_v18_: _dafny.MultiSet
            d_438_v18_ = _dafny.MultiSet([True])
            index46_ = default__.safeIndex(618, (d_434_v16_).length(0))
            (d_434_v16_)[index46_] = (d_437_v17_) + (default__.fm29(d_438_v18_, (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], (self).f41, (self).f42, globalState))
            (globalState).f13 = not(not((self).f41))
        elif True:
            d_439_v19_: _dafny.Array
            nw70_ = _dafny.Array(_dafny.Set({}), 1)
            d_439_v19_ = nw70_
            d_440_v20_: _dafny.Set
            d_440_v20_ = _dafny.Set({(self).f42})
            index47_ = default__.safeIndex(424, (d_439_v19_).length(0))
            (d_439_v19_)[index47_] = d_440_v20_
            index48_ = default__.safeIndex(424, (d_439_v19_).length(0))
            (d_439_v19_)[index48_] = (d_440_v20_) | (d_440_v20_)
            rhs37_ = (self).f41
            rhs38_ = (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]
            lhs22_ = globalState
            lhs22_.f1 = rhs37_
            r1 = rhs38_
            d_441_v21_: _dafny.Array
            def lambda22_(d_442_i3_):
                return _dafny.CodePoint('h')

            init14_ = lambda22_
            nw71_ = _dafny.Array(None, 21)
            for i0_14_ in range(nw71_.length(0)):
                nw71_[i0_14_] = init14_(i0_14_)
            d_441_v21_ = nw71_
            d_443_v22_: C0
            nw72_ = C0()
            nw72_.ctor__(d_441_v21_)
            d_443_v22_ = nw72_
            d_429_v11_ = (d_429_v11_) + (d_429_v11_)
            d_444_v23_: _dafny.Seq
            d_444_v23_ = _dafny.SeqWithoutIsStrInference([(self).f41, False])
            (globalState).f5 = ((d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]) + ((0) - (default__.safeModuloInt((0) - (len(d_444_v23_)), p0)))
        if (self).f41:
            d_445_v24_: _dafny.Seq
            d_445_v24_ = _dafny.SeqWithoutIsStrInference([(self).f42])
            d_446_v25_: C1
            nw73_ = C1()
            nw73_.ctor__(default__.fm33(d_445_v24_, default__.fm0(44, globalState), globalState))
            d_446_v25_ = nw73_
            d_447_v26_: _dafny.MultiSet
            d_447_v26_ = _dafny.MultiSet([False])
            if ((len(d_429_v11_)) <= ((self).f42)) in (d_447_v26_):
                d_448_v27_: _dafny.Map
                d_448_v27_ = _dafny.Map({d_429_v11_: d_447_v26_})
                d_449_v28_: _dafny.Seq
                d_449_v28_ = _dafny.SeqWithoutIsStrInference([(self).f41, not((self).f41)])
                d_450_v29_: D3
                d_450_v29_ = D3_DC7(d_429_v11_, (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], p0, (d_449_v28_).set(default__.safeIndex((self).f42, len(d_449_v28_)), (self).f41))
                rhs39_ = (not((self).f41)) or (((d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]) == (len(d_449_v28_)))
                rhs40_ = ((d_448_v27_) | (d_448_v27_)).set((d_450_v29_).cf8, _dafny.MultiSet([(self).f41]))
                rhs41_ = (0) - (p0)
                lhs23_ = globalState
                lhs24_ = globalState
                lhs23_.f1 = rhs39_
                d_448_v27_ = rhs40_
                lhs24_.f0 = rhs41_
                (globalState).f13 = not(not(not ((d_429_v11_) <= (d_429_v11_)) or (((self).f42) >= ((0) - ((self).f42)))))
                d_429_v11_ = d_429_v11_
                d_451_v30_: C1
                nw74_ = C1()
                nw74_.ctor__((d_446_v25_.f43) | (d_446_v25_.f43))
                d_451_v30_ = nw74_
                d_452_v31_: _dafny.Array
                def lambda23_(d_453_v26_):
                    def lambda24_(d_454_i4_):
                        return (d_453_v26_).issubset(_dafny.MultiSet([(self).f41]))

                    return lambda24_

                init15_ = lambda23_(d_447_v26_)
                nw75_ = _dafny.Array(None, 14)
                for i0_15_ in range(nw75_.length(0)):
                    nw75_[i0_15_] = init15_(i0_15_)
                d_452_v31_ = nw75_
                index49_ = default__.safeIndex(290, (d_452_v31_).length(0))
                (d_452_v31_)[index49_] = default__.fm0((0) - ((self).f42), globalState)
                index50_ = default__.safeIndex(290, (d_452_v31_).length(0))
                (d_452_v31_)[index50_] = (self).f41
            elif True:
                (globalState).f0 = (873 if (self).f41 else (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))])
                d_455_v32_: _dafny.Set
                d_455_v32_ = _dafny.Set({(0) - ((self).f42), (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_456_i5_ in range(default__.abs(53))]))) + (p0), (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], (self).f42, (self).f42})
                d_457_v33_: _dafny.Seq
                d_457_v33_ = _dafny.SeqWithoutIsStrInference([not((self).f41), False, (self).f41, (self).f41, (self).f41])
                d_458_v34_: _dafny.Set
                d_458_v34_ = _dafny.Set({(self).f41})
                d_459_v35_: _dafny.Map
                d_459_v35_ = _dafny.Map({d_429_v11_: (self).f41})
                d_460_v36_: _dafny.Map
                d_460_v36_ = _dafny.Map({(self).f41: (self).f41})
                d_461_v37_: _dafny.Array
                nw76_ = _dafny.Array(None, 27)
                nw76_[int(0)] = (default__.fm34(d_457_v33_, 191, globalState)).issubset(d_458_v34_)
                nw76_[int(1)] = (self).f41
                nw76_[int(2)] = ((self).f42) == ((self).f42)
                nw76_[int(3)] = (False if (d_457_v33_)[default__.safeIndex(733, len(d_457_v33_))] else (self).f41)
                nw76_[int(4)] = (self).f41
                nw76_[int(5)] = (self).f41
                nw76_[int(6)] = (self).f41
                nw76_[int(7)] = (self).f41
                nw76_[int(8)] = (self).f41
                nw76_[int(9)] = (self).f41
                nw76_[int(10)] = ((self).f41) == ((self).f41)
                nw76_[int(11)] = ((d_459_v35_)[d_429_v11_] if (d_429_v11_) in (d_459_v35_) else ((d_460_v36_)[True] if (True) in (d_460_v36_) else (self).f41))
                nw76_[int(12)] = (self).f41
                nw76_[int(13)] = (self).f41
                nw76_[int(14)] = default__.fm0(p0, globalState)
                nw76_[int(15)] = (p0) == ((self).f42)
                nw76_[int(16)] = (d_455_v32_).issubset(d_455_v32_)
                nw76_[int(17)] = not((self).f41)
                nw76_[int(18)] = (self).f41
                nw76_[int(19)] = (self).f41
                nw76_[int(20)] = not(default__.fm0((self).f42, globalState))
                nw76_[int(21)] = (self).f41
                nw76_[int(22)] = (self).f41
                nw76_[int(23)] = (self).f41
                nw76_[int(24)] = (self).f41
                nw76_[int(25)] = (self).f41
                nw76_[int(26)] = (self).f41
                d_461_v37_ = nw76_
                index51_ = default__.safeIndex(660, (d_461_v37_).length(0))
                (d_461_v37_)[index51_] = ((self).f41) not in (d_446_v25_.f43)
                index52_ = default__.safeIndex(660, (d_461_v37_).length(0))
                rhs42_ = (d_455_v32_).intersection(d_455_v32_)
                rhs43_ = default__.fm2(globalState)
                rhs44_ = ((self).f41 if (96) != (373) else True)
                lhs25_ = globalState
                lhs26_ = d_461_v37_
                lhs27_ = default__.safeIndex(660, (d_461_v37_).length(0))
                d_455_v32_ = rhs42_
                lhs25_.f2 = rhs43_
                lhs26_[lhs27_] = rhs44_
                d_462_v38_: str
                d_462_v38_ = _dafny.CodePoint('q')
                d_463_v39_: D0
                d_463_v39_ = D0_DC1(d_416_v0_, d_462_v38_, (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))])
                d_464_v40_: D0
                d_464_v40_ = D0_DC2(d_463_v39_)
                d_465_v41_: _dafny.Map
                d_465_v41_ = _dafny.Map({(d_461_v37_ if (d_461_v37_)[default__.safeIndex(660, (d_461_v37_).length(0))] else d_461_v37_): (d_464_v40_ if (d_461_v37_)[default__.safeIndex(660, (d_461_v37_).length(0))] else d_464_v40_)})
                d_465_v41_ = (d_465_v41_).set(d_461_v37_, d_464_v40_)
                (globalState).f13 = (self).f41
                d_466_v42_: D4
                d_466_v42_ = D4_DC8(d_462_v38_)
                d_467_v43_: _dafny.Seq
                d_467_v43_ = _dafny.SeqWithoutIsStrInference([d_466_v42_, D4_DC8(_dafny.CodePoint('o'))])
                d_468_v44_: _dafny.Map
                d_468_v44_ = _dafny.Map({d_462_v38_: d_429_v11_})
                d_469_v45_: _dafny.Map
                d_469_v45_ = _dafny.Map({d_467_v43_: _dafny.MultiSet([len(((d_468_v44_)[d_462_v38_] if (d_462_v38_) in (d_468_v44_) else d_429_v11_))])})
                d_470_v46_: _dafny.Map
                d_470_v46_ = _dafny.Map({default__.fm2(globalState): (self).f42})
                d_471_v47_: _dafny.MultiSet
                d_471_v47_ = _dafny.MultiSet([p0, ((d_470_v46_)[(d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]] if ((d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]) in (d_470_v46_) else len(_dafny.Set({(d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], ((d_447_v26_)[(self).f41] if ((self).f41) in (d_447_v26_) else (self).f42)}))), default__.fm2(globalState), (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]])
                pat_let_tv7_ = p0
                pat_let_tv8_ = globalState
                def iife55_(_pat_let8_0):
                    def iife56_(d_472_dt__update__tmp_h0_):
                        def iife57_(_pat_let9_0):
                            def iife58_(d_473_dt__update_hcf13_h0_):
                                return D4_DC8(d_473_dt__update_hcf13_h0_)
                            return iife58_(_pat_let9_0)
                        return iife57_(default__.fm35(pat_let_tv7_, pat_let_tv8_))
                    return iife56_(_pat_let8_0)
                d_469_v45_ = (d_469_v45_).set(_dafny.SeqWithoutIsStrInference([iife55_(d_466_v42_), d_466_v42_, default__.fm32((d_461_v37_)[default__.safeIndex(660, (d_461_v37_).length(0))], globalState)]), d_471_v47_)
            d_474_v48_: str
            d_474_v48_ = _dafny.CodePoint('o')
            d_475_v49_: _dafny.Array
            nw77_ = _dafny.Array(None, 11)
            nw77_[int(0)] = d_474_v48_
            nw77_[int(1)] = d_474_v48_
            nw77_[int(2)] = d_474_v48_
            nw77_[int(3)] = d_474_v48_
            nw77_[int(4)] = d_474_v48_
            nw77_[int(5)] = d_474_v48_
            nw77_[int(6)] = _dafny.CodePoint('g')
            nw77_[int(7)] = d_474_v48_
            nw77_[int(8)] = d_474_v48_
            nw77_[int(9)] = (d_429_v11_)[default__.safeIndex((d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))], len(d_429_v11_))]
            nw77_[int(10)] = d_474_v48_
            d_475_v49_ = nw77_
            d_476_v50_: _dafny.Map
            d_476_v50_ = _dafny.Map({d_475_v49_: _dafny.Map({(self).f42: (self).f41})})
            d_476_v50_ = d_476_v50_
            (globalState).f22 = (0) - ((self).f42)
            d_477_v51_: _dafny.Map
            d_477_v51_ = _dafny.Map({(d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]: False})
            d_478_v52_: _dafny.Map
            d_478_v52_ = _dafny.Map({d_477_v51_: _dafny.Map({(self).f41: len(_dafny.Map({(self).f41: len(d_445_v24_)}))})})
            d_479_v53_: D9
            d_479_v53_ = D9_DC22(d_478_v52_)
            source7_ = d_479_v53_
            if source7_.is_DC23:
                index53_ = default__.safeIndex(746, (d_416_v0_).length(0))
                (d_416_v0_)[index53_] = (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]
                d_480_v54_: _dafny.MultiSet
                d_480_v54_ = _dafny.MultiSet([d_416_v0_, d_416_v0_])
                r3 = ((_dafny.MultiSet([d_416_v0_, d_416_v0_])) | (d_480_v54_)).issubset(_dafny.MultiSet([d_416_v0_]))
                index54_ = default__.safeIndex(746, (d_416_v0_).length(0))
                (d_416_v0_)[index54_] = default__.safeModuloInt(default__.fm2(globalState), (self).f42)
                (globalState).f1 = (self).f41
            elif True:
                d_481___mcc_h0_ = source7_.cf33
                d_482_cf33_ = d_481___mcc_h0_
                d_483_v55_: _dafny.Array
                nw78_ = _dafny.Array(None, 10)
                nw78_[int(0)] = (self).f41
                nw78_[int(1)] = (self).f41
                nw78_[int(2)] = False
                nw78_[int(3)] = (self).f41
                nw78_[int(4)] = (self).f41
                nw78_[int(5)] = not((p0) < ((d_445_v24_)[default__.safeIndex((self).f42, len(d_445_v24_))]))
                nw78_[int(6)] = (self).f41
                nw78_[int(7)] = (self).f41
                nw78_[int(8)] = (self).f41
                nw78_[int(9)] = (len(d_429_v11_)) != (p0)
                d_483_v55_ = nw78_
                index55_ = default__.safeIndex(735, (d_483_v55_).length(0))
                (d_483_v55_)[index55_] = (self).f41
                d_484_v56_: _dafny.Seq
                d_484_v56_ = _dafny.SeqWithoutIsStrInference([(self).f41])
                index56_ = default__.safeIndex(735, (d_483_v55_).length(0))
                (d_483_v55_)[index56_] = ((self).f41) not in (d_484_v56_)
                index57_ = default__.safeIndex(735, (d_483_v55_).length(0))
                (d_483_v55_)[index57_] = (self).f41
                d_477_v51_ = (_dafny.Map({(d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]: ((d_477_v51_)[p0] if (p0) in (d_477_v51_) else not((self).f41))})) | (d_477_v51_)
                index58_ = default__.safeIndex(735, (d_483_v55_).length(0))
                (d_483_v55_)[index58_] = (len(d_429_v11_)) != (p0)
        elif True:
            (globalState).f5 = default__.fm2(globalState)
            d_485_v57_: _dafny.Map
            d_485_v57_ = _dafny.Map({(self).f42: (0) - (p0)})
            d_486_v58_: _dafny.Map
            d_486_v58_ = _dafny.Map({p0: d_485_v57_})
            (globalState).f12 = (d_485_v57_) | (((d_486_v58_)[p0] if (p0) in (d_486_v58_) else d_485_v57_))
            d_487_v59_: bool
            d_488_v60_: bool
            d_489_v61_: _dafny.Map
            d_490_v62_: int
            out19_: bool
            out20_: bool
            out21_: _dafny.Map
            out22_: int
            out19_, out20_, out21_, out22_ = default__.m0(globalState)
            d_487_v59_ = out19_
            d_488_v60_ = out20_
            d_489_v61_ = out21_
            d_490_v62_ = out22_
            (globalState).f0 = (0) - ((d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))])
            d_491_v63_: str
            d_491_v63_ = _dafny.CodePoint('w')
            d_491_v63_ = d_491_v63_
        d_492_v64_: _dafny.MultiSet
        d_492_v64_ = _dafny.MultiSet([(self).f41])
        d_493_v65_: _dafny.Map
        d_493_v65_ = _dafny.Map({not(not ((self).f41) or ((self).f41)): d_492_v64_})
        d_494_v66_: D7
        d_494_v66_ = D7_DC17((self).f41, (self).f41, (self).f42)
        d_493_v65_ = (d_493_v65_).set((d_494_v66_).cf25, d_492_v64_)
        d_495_v67_: _dafny.Seq
        d_495_v67_ = _dafny.SeqWithoutIsStrInference([not((self).f41)])
        if (d_492_v64_).ispropersubset(_dafny.MultiSet(d_495_v67_)):
            r2 = (self).f41
            d_496_v68_: D7
            d_497_v69_: int
            d_498_v70_: bool
            out23_: D7
            out24_: int
            out25_: bool
            out23_, out24_, out25_ = (self).m20((self).f41, globalState)
            d_496_v68_ = out23_
            d_497_v69_ = out24_
            d_498_v70_ = out25_
            d_416_v0_ = d_416_v0_
            d_499_v71_: str
            d_499_v71_ = _dafny.CodePoint('e')
            d_500_v72_: D4
            d_500_v72_ = D4_DC8(d_499_v71_)
            d_501_v73_: D8
            d_501_v73_ = D8_DC19((self).f41, d_498_v70_, d_498_v70_)
            pat_let_tv9_ = d_499_v71_
            def iife59_(_pat_let10_0):
                def iife60_(d_502_dt__update__tmp_h1_):
                    def iife61_(_pat_let11_0):
                        def iife62_(d_503_dt__update_hcf13_h1_):
                            return D4_DC8(d_503_dt__update_hcf13_h1_)
                        return iife62_(_pat_let11_0)
                    return iife61_(pat_let_tv9_)
                return iife60_(_pat_let10_0)
            rhs45_ = (d_501_v73_).cf29
            rhs46_ = default__.safeDivisionInt(d_497_v69_, p0)
            rhs47_ = (self).f42
            rhs48_ = True
            rhs49_ = iife59_(d_500_v72_)
            lhs28_ = globalState
            lhs29_ = globalState
            lhs30_ = globalState
            lhs31_ = globalState
            lhs28_.f13 = rhs45_
            lhs29_.f2 = rhs46_
            lhs30_.f0 = rhs47_
            lhs31_.f1 = rhs48_
            d_500_v72_ = rhs49_
            d_504_v74_: _dafny.Array
            nw79_ = _dafny.Array(_dafny.CodePoint('D'), 1)
            d_504_v74_ = nw79_
            d_505_v75_: C0
            nw80_ = C0()
            nw80_.ctor__(d_504_v74_)
            d_505_v75_ = nw80_
        elif True:
            d_506_v77_: _dafny.MultiSet
            d_506_v77_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsardqrq")), d_429_v11_])
            d_507_v78_: str
            d_507_v78_ = _dafny.CodePoint('t')
            d_508_v79_: D4
            d_508_v79_ = D4_DC8(d_507_v78_)
            d_509_v80_: _dafny.Set
            d_509_v80_ = _dafny.Set({d_508_v79_, d_508_v79_})
            d_510_v82_: _dafny.Set
            def iife63_():
                coll39_ = _dafny.Set()
                compr_39_: D4
                for compr_39_ in (d_509_v80_).Elements:
                    d_511_v81_: D4 = compr_39_
                    if (d_511_v81_) in (d_509_v80_):
                        coll39_ = coll39_.union(_dafny.Set([d_511_v81_]))
                return _dafny.Set(coll39_)
            d_510_v82_ = _dafny.Set({_dafny.Set({d_508_v79_, d_508_v79_}), d_509_v80_, iife63_()
            , d_509_v80_, d_509_v80_})
            d_512_v83_: _dafny.Array
            nw81_ = _dafny.Array(None, 13)
            nw81_[int(0)] = ((self).f41) and (True)
            nw81_[int(1)] = (self).f41
            nw81_[int(2)] = (self).f41
            nw81_[int(3)] = (152) != ((self).f42)
            nw81_[int(4)] = not(not((self).f41))
            def iife64_():
                coll40_ = _dafny.Map()
                compr_40_: _dafny.Seq
                for compr_40_ in (d_506_v77_).Elements:
                    d_513_v76_: _dafny.Seq = compr_40_
                    if (d_513_v76_) in (d_506_v77_):
                        coll40_[d_513_v76_] = (self).f41
                return _dafny.Map(coll40_)
            nw81_[int(5)] = (d_429_v11_) not in (iife64_()
            )
            nw81_[int(6)] = (self).f41
            nw81_[int(7)] = (self).f41
            nw81_[int(8)] = (d_510_v82_).issubset(d_510_v82_)
            nw81_[int(9)] = (self).f41
            nw81_[int(10)] = (self).f41
            nw81_[int(11)] = False
            nw81_[int(12)] = (d_429_v11_) == (d_429_v11_)
            d_512_v83_ = nw81_
            d_512_v83_ = d_512_v83_
            (globalState).f1 = (self).f41
            r3 = (D8_DC20((self).f41)).cf31
            if (self).f41:
                d_514_v84_: _dafny.Set
                d_514_v84_ = _dafny.Set({p0, (self).f42, p0, p0, p0})
                d_515_v85_: _dafny.Map
                d_515_v85_ = _dafny.Map({(d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]: d_514_v84_})
                d_516_v86_: _dafny.Array
                nw82_ = _dafny.Array(None, 14)
                nw82_[int(0)] = d_507_v78_
                nw82_[int(1)] = d_507_v78_
                nw82_[int(2)] = default__.fm35(len(d_515_v85_), globalState)
                nw82_[int(3)] = d_507_v78_
                nw82_[int(4)] = d_507_v78_
                nw82_[int(5)] = d_507_v78_
                nw82_[int(6)] = d_507_v78_
                nw82_[int(7)] = d_507_v78_
                nw82_[int(8)] = d_507_v78_
                nw82_[int(9)] = _dafny.CodePoint('u')
                nw82_[int(10)] = d_507_v78_
                nw82_[int(11)] = d_507_v78_
                nw82_[int(12)] = d_507_v78_
                nw82_[int(13)] = d_507_v78_
                d_516_v86_ = nw82_
                d_517_v87_: C0
                nw83_ = C0()
                nw83_.ctor__(d_516_v86_)
                d_517_v87_ = nw83_
                d_518_v88_: _dafny.Map
                d_518_v88_ = _dafny.Map({(self).f41: p0})
                d_518_v88_ = (d_518_v88_).set((self).f41, (self).f42)
                r3 = (self).f41
                index59_ = default__.safeIndex(746, (d_416_v0_).length(0))
                rhs50_ = (self).f42
                rhs51_ = (self).f41
                rhs52_ = (d_429_v11_) < (d_429_v11_)
                lhs32_ = d_416_v0_
                lhs33_ = default__.safeIndex(746, (d_416_v0_).length(0))
                lhs34_ = globalState
                lhs32_[lhs33_] = rhs50_
                lhs34_.f1 = rhs51_
                r3 = rhs52_
                d_519_v90_: _dafny.Map
                def iife65_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in (_dafny.Set({p0})).Elements:
                        d_520_v89_: int = compr_41_
                        if (d_520_v89_) in (_dafny.Set({p0})):
                            coll41_[default__.safeModuloInt(d_520_v89_, p0)] = p0
                    return _dafny.Map(coll41_)
                d_519_v90_ = _dafny.Map({(self).f41: iife65_()
                })
                d_521_v91_: _dafny.Map
                d_521_v91_ = _dafny.Map({(self).f41: (self).f41})
                rhs53_ = len((d_519_v90_) | ((d_519_v90_) | (d_519_v90_)))
                rhs54_ = len(d_521_v91_)
                lhs35_ = globalState
                lhs36_ = globalState
                lhs35_.f21 = rhs53_
                lhs36_.f16 = rhs54_
            elif True:
                (globalState).f13 = (self).f41
                r3 = (d_492_v64_).issubset(d_492_v64_)
                d_522_v92_: _dafny.Array
                def lambda25_(d_523_v67_):
                    def lambda26_(d_524_i6_):
                        return d_523_v67_

                    return lambda26_

                init16_ = lambda25_(d_495_v67_)
                nw84_ = _dafny.Array(None, 5)
                for i0_16_ in range(nw84_.length(0)):
                    nw84_[i0_16_] = init16_(i0_16_)
                d_522_v92_ = nw84_
                d_522_v92_ = d_522_v92_
                d_525_v93_: _dafny.Map
                d_525_v93_ = _dafny.Map({(self).f42: d_429_v11_})
                d_429_v11_ = ((d_525_v93_)[default__.safeModuloInt((self).f42, (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))])] if (default__.safeModuloInt((self).f42, (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))])) in (d_525_v93_) else d_429_v11_)
                d_507_v78_ = d_507_v78_
            d_526_v94_: _dafny.Array
            nw85_ = _dafny.Array(None, 3)
            nw85_[int(0)] = d_507_v78_
            nw85_[int(1)] = d_507_v78_
            nw85_[int(2)] = d_507_v78_
            d_526_v94_ = nw85_
            d_527_v95_: C0
            nw86_ = C0()
            nw86_.ctor__(d_526_v94_)
            d_527_v95_ = nw86_
        r0 = (d_416_v0_)[default__.safeIndex(746, (d_416_v0_).length(0))]
        d_528_v96_: _dafny.Seq
        d_528_v96_ = _dafny.SeqWithoutIsStrInference([(0) - (-406)])
        r1 = ((d_528_v96_)[default__.safeIndex(p0, len(d_528_v96_))]) + (len(d_528_v96_))
        d_529_v97_: _dafny.Seq
        d_529_v97_ = default__.fm29(d_492_v64_, (self).f42, (self).f41, (self).f42, globalState)
        def lambda27_(source8_):
            d_530___mcc_h1_ = source8_
            d_531_cf20_ = d_530___mcc_h1_
            return (self).f41

        r2 = lambda27_(d_529_v97_)
        r3 = not((self).f41)
        return r0, r1, r2, r3

    def m20(self, p0, globalState):
        r0: D7 = D7.default()()
        r1: int = int(0)
        r2: bool = False
        d_532_v0_: _dafny.Seq
        d_532_v0_ = _dafny.SeqWithoutIsStrInference([(self).f42])
        d_533_v1_: _dafny.Array
        def lambda28_(d_534_i0_):
            return (self).f41

        init17_ = lambda28_
        nw87_ = _dafny.Array(None, 29)
        for i0_17_ in range(nw87_.length(0)):
            nw87_[i0_17_] = init17_(i0_17_)
        d_533_v1_ = nw87_
        d_535_v2_: _dafny.Map
        d_535_v2_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ieyd"))): d_533_v1_})
        d_536_v3_: str
        d_536_v3_ = _dafny.CodePoint('w')
        d_537_v4_: D7
        d_537_v4_ = D7_DC17(p0, p0, (self).f42)
        d_538_v5_: _dafny.Map
        d_538_v5_ = _dafny.Map({(0) - ((self).f42): (self).f42})
        d_539_v6_: _dafny.Seq
        d_539_v6_ = _dafny.SeqWithoutIsStrInference([p0])
        d_540_v7_: _dafny.Seq
        d_540_v7_ = _dafny.SeqWithoutIsStrInference([d_539_v6_, d_539_v6_, (d_539_v6_).set(default__.safeIndex((self).f42, len(d_539_v6_)), (self).f41)])
        d_541_v8_: _dafny.Set
        d_541_v8_ = _dafny.Set({((d_538_v5_)[len(d_540_v7_)] if (len(d_540_v7_)) in (d_538_v5_) else (self).f42), (self).f42})
        d_542_v9_: _dafny.MultiSet
        d_542_v9_ = _dafny.MultiSet([d_536_v3_])
        d_543_v10_: _dafny.Array
        nw88_ = _dafny.Array(None, 24)
        nw88_[int(0)] = (self).f42
        nw88_[int(1)] = len(d_532_v0_)
        nw88_[int(2)] = (self).f42
        nw88_[int(3)] = (self).f42
        nw88_[int(4)] = len(d_535_v2_)
        nw88_[int(5)] = (self).f42
        nw88_[int(6)] = (self).f42
        nw88_[int(7)] = (self).f42
        nw88_[int(8)] = (self).f42
        nw88_[int(9)] = 614
        nw88_[int(10)] = (self).f42
        nw88_[int(11)] = (self).f42
        nw88_[int(12)] = len(_dafny.Map({True: d_536_v3_}))
        nw88_[int(13)] = -868
        nw88_[int(14)] = (d_537_v4_).cf26
        nw88_[int(15)] = len(d_541_v8_)
        nw88_[int(16)] = len(d_539_v6_)
        nw88_[int(17)] = (self).f42
        nw88_[int(18)] = (self).f42
        nw88_[int(19)] = (d_542_v9_).cardinality
        nw88_[int(20)] = (d_532_v0_)[default__.safeIndex(225, len(d_532_v0_))]
        nw88_[int(21)] = (self).f42
        nw88_[int(22)] = (self).f42
        nw88_[int(23)] = (self).f42
        d_543_v10_ = nw88_
        d_544_v11_: D4
        d_544_v11_ = D4_DC10(default__.fm2(globalState))
        d_545_v12_: _dafny.Map
        d_545_v12_ = _dafny.Map({d_543_v10_: d_544_v11_})
        d_545_v12_ = d_545_v12_
        (globalState).f22 = default__.safeDivisionInt((0) - ((self).f42), -665)
        d_546_v14_: _dafny.Map
        d_546_v14_ = _dafny.Map({len(d_532_v0_): p0})
        def iife66_():
            coll42_ = _dafny.Map()
            compr_42_: int
            for compr_42_ in (d_532_v0_).Elements:
                d_547_v13_: int = compr_42_
                if (d_547_v13_) in (d_532_v0_):
                    coll42_[default__.safeDivisionInt(d_547_v13_, (self).f42)] = p0
            return _dafny.Map(coll42_)
        if ((self).f41 if (not(p0)) == (p0) else (iife66_()
        ) != (d_546_v14_)):
            d_548_v15_: bool
            d_549_v16_: bool
            d_550_v17_: _dafny.Map
            d_551_v18_: int
            out26_: bool
            out27_: bool
            out28_: _dafny.Map
            out29_: int
            out26_, out27_, out28_, out29_ = default__.m0(globalState)
            d_548_v15_ = out26_
            d_549_v16_ = out27_
            d_550_v17_ = out28_
            d_551_v18_ = out29_
            d_552_v19_: D8
            d_552_v19_ = D8_DC18((d_541_v8_) | (_dafny.Set({d_551_v18_})))
            source9_ = d_552_v19_
            if source9_.is_DC19:
                d_553___mcc_h0_ = source9_.cf28
                d_554___mcc_h1_ = source9_.cf29
                d_555___mcc_h2_ = source9_.cf30
                d_556_cf30_ = d_555___mcc_h2_
                d_557_cf29_ = d_554___mcc_h1_
                d_558_cf28_ = d_553___mcc_h0_
                d_559_v20_: _dafny.Seq
                d_559_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sg"))
                d_559_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bikyhya"))
                d_557_cf29_ = (d_539_v6_)[default__.safeIndex((d_551_v18_) + ((self).f42), len(d_539_v6_))]
                d_560_v21_: _dafny.Array
                nw89_ = _dafny.Array(_dafny.Seq({}), 5)
                d_560_v21_ = nw89_
                index60_ = default__.safeIndex(241, (d_560_v21_).length(0))
                (d_560_v21_)[index60_] = _dafny.SeqWithoutIsStrInference([not(p0), (self).f41, d_548_v15_, False, False])
                d_561_v22_: _dafny.MultiSet
                d_561_v22_ = _dafny.MultiSet([d_549_v16_])
                d_562_v23_: _dafny.MultiSet
                d_562_v23_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_551_v18_]), (d_532_v0_) + (d_532_v0_), default__.fm29(d_561_v22_, 570, False, (self).f42, globalState), _dafny.SeqWithoutIsStrInference([d_551_v18_, -758]), d_532_v0_])
                d_563_v24_: _dafny.Map
                d_563_v24_ = _dafny.Map({d_548_v15_: d_556_cf30_})
                index61_ = default__.safeIndex(241, (d_560_v21_).length(0))
                rhs55_ = (d_539_v6_) + ((d_539_v6_) + ((d_539_v6_).set(default__.safeIndex((0) - (len(d_563_v24_)), len(d_539_v6_)), d_548_v15_)))
                rhs56_ = d_539_v6_
                rhs57_ = (0) - ((0) - (d_551_v18_))
                rhs58_ = default__.fm2(globalState)
                rhs59_ = d_562_v23_
                lhs37_ = d_560_v21_
                lhs38_ = default__.safeIndex(241, (d_560_v21_).length(0))
                lhs39_ = globalState
                lhs40_ = globalState
                lhs37_[lhs38_] = rhs55_
                d_539_v6_ = rhs56_
                lhs39_.f21 = rhs57_
                lhs40_.f22 = rhs58_
                d_562_v23_ = rhs59_
                d_564_v25_: _dafny.Array
                def lambda29_(d_565_i1_):
                    return _dafny.MultiSet([(self).f42])

                init18_ = lambda29_
                nw90_ = _dafny.Array(None, 4)
                for i0_18_ in range(nw90_.length(0)):
                    nw90_[i0_18_] = init18_(i0_18_)
                d_564_v25_ = nw90_
                d_566_v26_: _dafny.MultiSet
                d_566_v26_ = _dafny.MultiSet([d_551_v18_, (self).f42])
                index62_ = default__.safeIndex(585, (d_564_v25_).length(0))
                (d_564_v25_)[index62_] = d_566_v26_
                d_567_v27_: _dafny.Array
                nw91_ = _dafny.Array(_dafny.Map({}), 5)
                d_567_v27_ = nw91_
                d_568_v28_: _dafny.Map
                d_568_v28_ = _dafny.Map({d_557_cf29_: d_541_v8_})
                d_569_v29_: _dafny.Map
                d_569_v29_ = _dafny.Map({457: d_568_v28_})
                index63_ = default__.safeIndex(395, (d_567_v27_).length(0))
                (d_567_v27_)[index63_] = ((d_569_v29_)[(self).f42] if ((self).f42) in (d_569_v29_) else _dafny.Map({d_556_cf30_: d_541_v8_}))
                d_570_v30_: D10
                d_570_v30_ = D10_DC25(282, len(d_541_v8_), not((self).f41), d_556_cf30_)
                index64_ = default__.safeIndex(585, (d_564_v25_).length(0))
                index65_ = default__.safeIndex(395, (d_567_v27_).length(0))
                rhs60_ = (d_566_v26_) - ((d_566_v26_).intersection(d_566_v26_))
                rhs61_ = (self).f42
                rhs62_ = (default__.fm36((self).f42, d_570_v30_, globalState)) | (d_568_v28_)
                lhs41_ = d_564_v25_
                lhs42_ = default__.safeIndex(585, (d_564_v25_).length(0))
                lhs43_ = globalState
                lhs44_ = d_567_v27_
                lhs45_ = default__.safeIndex(395, (d_567_v27_).length(0))
                lhs41_[lhs42_] = rhs60_
                lhs43_.f22 = rhs61_
                lhs44_[lhs45_] = rhs62_
            elif source9_.is_DC20:
                d_571___mcc_h3_ = source9_.cf31
                d_572_cf31_ = d_571___mcc_h3_
                d_573_v31_: _dafny.Array
                def lambda30_(d_574_v16_):
                    def lambda31_(d_575_i2_):
                        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dg")) if d_574_v16_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "giqcy")))

                    return lambda31_

                init19_ = lambda30_(d_549_v16_)
                nw92_ = _dafny.Array(None, 20)
                for i0_19_ in range(nw92_.length(0)):
                    nw92_[i0_19_] = init19_(i0_19_)
                d_573_v31_ = nw92_
                d_576_v32_: _dafny.MultiSet
                d_576_v32_ = _dafny.MultiSet([d_543_v10_, d_543_v10_, d_543_v10_, d_543_v10_])
                rhs63_ = (self).f42
                rhs64_ = ((self).f42) + (len(d_539_v6_))
                rhs65_ = ((_dafny.MultiSet([d_543_v10_, d_543_v10_, d_543_v10_])) - ((d_576_v32_).set(d_543_v10_, default__.abs(len(d_539_v6_))))).cardinality
                rhs66_ = (self).fm24(globalState)
                rhs67_ = d_573_v31_
                lhs46_ = globalState
                lhs47_ = globalState
                lhs48_ = globalState
                lhs49_ = globalState
                lhs46_.f0 = rhs63_
                lhs47_.f22 = rhs64_
                lhs48_.f5 = rhs65_
                lhs49_.f1 = rhs66_
                d_573_v31_ = rhs67_
                d_543_v10_ = d_543_v10_
                rhs68_ = d_536_v3_
                rhs69_ = (D7_DC17(p0, p0, d_551_v18_)).cf26
                lhs50_ = globalState
                d_536_v3_ = rhs68_
                lhs50_.f5 = rhs69_
                d_577_v33_: _dafny.Array
                nw93_ = _dafny.Array(_dafny.Set({}), 2)
                d_577_v33_ = nw93_
                d_578_v34_: _dafny.Array
                nw94_ = _dafny.Array(_dafny.Set({}), 11)
                d_578_v34_ = nw94_
                d_579_v35_: D7
                d_579_v35_ = D7_DC15(d_578_v34_)
                d_580_v36_: _dafny.MultiSet
                d_580_v36_ = _dafny.MultiSet([d_579_v35_])
                d_581_v37_: _dafny.Set
                d_581_v37_ = _dafny.Set({d_580_v36_, d_580_v36_, d_580_v36_})
                index66_ = default__.safeIndex(824, (d_577_v33_).length(0))
                (d_577_v33_)[index66_] = d_581_v37_
                index67_ = default__.safeIndex(824, (d_577_v33_).length(0))
                (d_577_v33_)[index67_] = _dafny.Set({d_580_v36_, _dafny.MultiSet([d_579_v35_]), (d_580_v36_).set(D7_DC15(d_578_v34_), default__.abs((self).f42)), d_580_v36_, (d_580_v36_) - (d_580_v36_)})
            elif source9_.is_DC18:
                d_582___mcc_h4_ = source9_.cf27
                d_583_cf27_ = d_582___mcc_h4_
                d_584_v38_: _dafny.Seq
                d_584_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "toboqogtk"))
                d_584_v38_ = ((d_584_v38_ if p0 else _dafny.SeqWithoutIsStrInference([d_536_v3_ for d_585_i3_ in range(default__.abs(98))]))).set(default__.safeIndex(((self).f42) - ((self).f42), len((d_584_v38_ if p0 else _dafny.SeqWithoutIsStrInference([d_536_v3_ for d_586_i3_ in range(default__.abs(98))])))), d_536_v3_)
                d_587_v39_: _dafny.Array
                nw95_ = _dafny.Array(_dafny.Seq({}), 15)
                d_587_v39_ = nw95_
                d_588_v40_: _dafny.Seq
                d_588_v40_ = _dafny.SeqWithoutIsStrInference([d_546_v14_, d_546_v14_, d_546_v14_, default__.fm37(d_549_v16_, globalState)])
                index68_ = default__.safeIndex(303, (d_587_v39_).length(0))
                (d_587_v39_)[index68_] = (d_588_v40_) + (d_588_v40_)
                index69_ = default__.safeIndex(303, (d_587_v39_).length(0))
                def iife67_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in (d_532_v0_).Elements:
                        d_590_v41_: int = compr_43_
                        if (d_590_v41_) in (d_532_v0_):
                            coll43_[(d_590_v41_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcyyk"))))] = False
                    return _dafny.Map(coll43_)
                rhs70_ = (_dafny.SeqWithoutIsStrInference([d_546_v14_ for d_589_i4_ in range(default__.abs(41))])) + ((d_588_v40_).set(default__.safeIndex((self).f42, len(d_588_v40_)), iife67_()
                ))
                rhs71_ = (default__.fm2(globalState)) + ((self).f42)
                rhs72_ = default__.safeModuloInt(default__.safeDivisionInt(630, d_551_v18_), d_551_v18_)
                lhs51_ = d_587_v39_
                lhs52_ = default__.safeIndex(303, (d_587_v39_).length(0))
                lhs53_ = globalState
                lhs54_ = globalState
                lhs51_[lhs52_] = rhs70_
                lhs53_.f21 = rhs71_
                lhs54_.f22 = rhs72_
                d_591_v42_: _dafny.Map
                d_591_v42_ = _dafny.Map({d_548_v15_: (0) - ((self).f42)})
                (globalState).f22 = ((d_538_v5_)[(self).f42] if ((self).f42) in (d_538_v5_) else ((d_591_v42_)[(self).f41] if ((self).f41) in (d_591_v42_) else len(_dafny.Map({858: d_551_v18_}))))
                d_592_v44_: _dafny.Seq
                d_592_v44_ = _dafny.SeqWithoutIsStrInference([d_532_v0_])
                d_593_v45_: _dafny.MultiSet
                d_593_v45_ = _dafny.MultiSet([len(d_584_v38_), d_551_v18_, d_551_v18_])
                def iife68_():
                    coll44_ = _dafny.Map()
                    compr_44_: int
                    for compr_44_ in (((d_592_v44_)[default__.safeIndex((d_593_v45_).cardinality, len(d_592_v44_))]) + (_dafny.SeqWithoutIsStrInference([d_551_v18_]))).Elements:
                        d_594_v43_: int = compr_44_
                        if (d_594_v43_) in (((d_592_v44_)[default__.safeIndex((d_593_v45_).cardinality, len(d_592_v44_))]) + (_dafny.SeqWithoutIsStrInference([d_551_v18_]))):
                            coll44_[(d_594_v43_) * (-654)] = d_548_v15_
                    return _dafny.Map(coll44_)
                d_546_v14_ = iife68_()
                
            elif True:
                d_595___mcc_h5_ = source9_.cf32
                d_596_cf32_ = d_595___mcc_h5_
                (globalState).f22 = default__.safeDivisionInt((self).f42, 897)
                d_597_v46_: _dafny.Set
                d_597_v46_ = _dafny.Set({(d_539_v6_)[default__.safeIndex(d_551_v18_, len(d_539_v6_))], d_548_v15_, d_549_v16_})
                (globalState).f1 = (d_597_v46_).issubset(_dafny.Set({d_549_v16_}))
                index70_ = default__.safeIndex(82, (d_533_v1_).length(0))
                (d_533_v1_)[index70_] = (d_539_v6_)[default__.safeIndex(d_551_v18_, len(d_539_v6_))]
                index71_ = default__.safeIndex(82, (d_533_v1_).length(0))
                (d_533_v1_)[index71_] = d_549_v16_
                d_598_v47_: _dafny.Array
                nw96_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_598_v47_ = nw96_
                d_599_v48_: _dafny.Array
                nw97_ = _dafny.Array(None, 25)
                nw97_[int(0)] = not((d_533_v1_)[default__.safeIndex(82, (d_533_v1_).length(0))])
                nw97_[int(1)] = (d_533_v1_)[default__.safeIndex(82, (d_533_v1_).length(0))]
                nw97_[int(2)] = p0
                nw97_[int(3)] = (d_533_v1_)[default__.safeIndex(82, (d_533_v1_).length(0))]
                nw97_[int(4)] = d_549_v16_
                nw97_[int(5)] = d_548_v15_
                nw97_[int(6)] = (d_533_v1_)[default__.safeIndex(82, (d_533_v1_).length(0))]
                nw97_[int(7)] = (d_533_v1_)[default__.safeIndex(82, (d_533_v1_).length(0))]
                nw97_[int(8)] = d_548_v15_
                nw97_[int(9)] = d_549_v16_
                nw97_[int(10)] = d_549_v16_
                nw97_[int(11)] = not((self).f41)
                nw97_[int(12)] = p0
                nw97_[int(13)] = not((d_533_v1_)[default__.safeIndex(82, (d_533_v1_).length(0))])
                nw97_[int(14)] = p0
                nw97_[int(15)] = d_548_v15_
                nw97_[int(16)] = False
                nw97_[int(17)] = (self).f41
                nw97_[int(18)] = p0
                nw97_[int(19)] = p0
                nw97_[int(20)] = not(p0)
                nw97_[int(21)] = (d_533_v1_)[default__.safeIndex(82, (d_533_v1_).length(0))]
                nw97_[int(22)] = p0
                nw97_[int(23)] = d_549_v16_
                nw97_[int(24)] = p0
                d_599_v48_ = nw97_
                index72_ = default__.safeIndex(648, (d_598_v47_).length(0))
                (d_598_v47_)[index72_] = d_599_v48_
                index73_ = default__.safeIndex(648, (d_598_v47_).length(0))
                (d_598_v47_)[index73_] = d_599_v48_
            if (self).f41:
                d_600_v49_: _dafny.MultiSet
                d_600_v49_ = _dafny.MultiSet([d_549_v16_, d_548_v15_])
                d_601_v50_: _dafny.Map
                d_601_v50_ = _dafny.Map({p0: len(default__.fm29(d_600_v49_, (self).f42, d_548_v15_, d_551_v18_, globalState))})
                d_602_v51_: _dafny.Map
                d_602_v51_ = _dafny.Map({-743: (d_601_v50_) | (d_601_v50_)})
                d_602_v51_ = (d_602_v51_).set((len(d_546_v14_)) * (d_551_v18_), d_601_v50_)
                d_603_v52_: _dafny.MultiSet
                d_603_v52_ = _dafny.MultiSet([225])
                d_604_v53_: _dafny.Map
                d_604_v53_ = _dafny.Map({d_543_v10_: _dafny.SeqWithoutIsStrInference([d_536_v3_])})
                d_605_v54_: _dafny.Seq
                d_605_v54_ = _dafny.SeqWithoutIsStrInference([d_536_v3_, d_536_v3_])
                index74_ = default__.safeIndex(849, (d_543_v10_).length(0))
                (d_543_v10_)[index74_] = default__.safeModuloInt((0) - (((d_603_v52_)[d_551_v18_] if (d_551_v18_) in (d_603_v52_) else (self).f42)), len(((d_604_v53_)[d_543_v10_] if (d_543_v10_) in (d_604_v53_) else d_605_v54_)))
                index75_ = default__.safeIndex(849, (d_543_v10_).length(0))
                (d_543_v10_)[index75_] = (836) + (d_551_v18_)
                (globalState).f13 = not(d_548_v15_)
                d_548_v15_ = (d_549_v16_) or (p0)
                pat_let_tv10_ = d_541_v8_
                def iife69_(_pat_let12_0):
                    def iife70_(d_606_dt__update__tmp_h0_):
                        def iife71_(_pat_let13_0):
                            def iife72_(d_607_dt__update_hcf27_h0_):
                                return D8_DC18(d_607_dt__update_hcf27_h0_)
                            return iife72_(_pat_let13_0)
                        return iife71_(pat_let_tv10_)
                    return iife70_(_pat_let12_0)
                d_552_v19_ = iife69_(d_552_v19_)
            elif True:
                d_608_v55_: _dafny.Seq
                d_608_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                d_609_v56_: D10
                d_609_v56_ = D10_DC25(d_551_v18_, (self).f42, d_549_v16_, True)
                d_610_v57_: _dafny.MultiSet
                d_610_v57_ = _dafny.MultiSet([d_551_v18_, (self).f42, (0) - (len(d_608_v55_)), (d_609_v56_).cf36])
                d_611_v58_: _dafny.Map
                d_611_v58_ = _dafny.Map({d_533_v1_: d_610_v57_})
                d_611_v58_ = d_611_v58_
                d_549_v16_ = (d_551_v18_) <= (((0) - (len(d_541_v8_))) + ((0) - (d_551_v18_)))
                (globalState).f16 = (0) - ((d_551_v18_) + (d_551_v18_))
                d_532_v0_ = (d_532_v0_) + (_dafny.SeqWithoutIsStrInference([len(d_546_v14_) for d_612_i5_ in range(default__.abs(226))]))
                d_537_v4_ = d_537_v4_
            d_613_v59_: _dafny.Array
            nw98_ = _dafny.Array(_dafny.CodePoint('D'), 19)
            d_613_v59_ = nw98_
            d_614_v60_: C0
            nw99_ = C0()
            nw99_.ctor__(d_613_v59_)
            d_614_v60_ = nw99_
            d_615_v61_: _dafny.Seq
            d_615_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fl"))
            d_616_v62_: D3
            d_616_v62_ = D3_DC7(d_615_v61_, d_551_v18_, d_551_v18_, d_551_v18_, d_539_v6_)
            pat_let_tv11_ = d_551_v18_
            pat_let_tv12_ = d_539_v6_
            pat_let_tv13_ = d_551_v18_
            pat_let_tv14_ = d_539_v6_
            def iife73_(_pat_let14_0):
                def iife74_(d_617_dt__update__tmp_h1_):
                    def iife75_(_pat_let15_0):
                        def iife76_(d_618_dt__update_hcf10_h0_):
                            def iife77_(_pat_let16_0):
                                def iife78_(d_619_dt__update_hcf12_h0_):
                                    return D3_DC7((d_617_dt__update__tmp_h1_).cf8, (d_617_dt__update__tmp_h1_).cf9, d_618_dt__update_hcf10_h0_, (d_617_dt__update__tmp_h1_).cf11, d_619_dt__update_hcf12_h0_)
                                return iife78_(_pat_let16_0)
                            return iife77_(pat_let_tv12_)
                        return iife76_(_pat_let15_0)
                    return iife75_(pat_let_tv11_)
                return iife74_(_pat_let14_0)
            def iife79_(_pat_let17_0):
                def iife80_(d_620_dt__update__tmp_h2_):
                    def iife81_(_pat_let18_0):
                        def iife82_(d_621_dt__update_hcf10_h1_):
                            def iife83_(_pat_let19_0):
                                def iife84_(d_622_dt__update_hcf12_h1_):
                                    return D3_DC7((d_620_dt__update__tmp_h2_).cf8, (d_620_dt__update__tmp_h2_).cf9, d_621_dt__update_hcf10_h1_, (d_620_dt__update__tmp_h2_).cf11, d_622_dt__update_hcf12_h1_)
                                return iife84_(_pat_let19_0)
                            return iife83_(pat_let_tv14_)
                        return iife82_(_pat_let18_0)
                    return iife81_(pat_let_tv13_)
                return iife80_(_pat_let17_0)
            d_615_v61_ = ((iife73_(d_616_v62_)).cf8).set(default__.safeIndex((self).f42, len((iife79_(d_616_v62_)).cf8)), d_536_v3_)
        elif True:
            (globalState).f13 = p0
            d_623_v63_: D4
            d_623_v63_ = D4_DC8(default__.fm35((self).f42, globalState))
            d_624_v64_: _dafny.Array
            def lambda32_(d_625_p0_):
                def lambda33_(d_626_i6_):
                    return (_dafny.MultiSet([d_625_p0_])).intersection(_dafny.MultiSet([(self).f41, d_625_p0_]))

                return lambda33_

            init20_ = lambda32_(p0)
            nw100_ = _dafny.Array(None, 3)
            for i0_20_ in range(nw100_.length(0)):
                nw100_[i0_20_] = init20_(i0_20_)
            d_624_v64_ = nw100_
            index76_ = default__.safeIndex(261, (d_624_v64_).length(0))
            (d_624_v64_)[index76_] = default__.fm31(globalState)
            d_627_v65_: _dafny.Array
            nw101_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
            d_627_v65_ = nw101_
            d_628_v66_: _dafny.Seq
            d_628_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmcs"))
            index77_ = default__.safeIndex(605, (d_627_v65_).length(0))
            (d_627_v65_)[index77_] = (d_628_v66_) + (d_628_v66_)
            index78_ = default__.safeIndex(261, (d_624_v64_).length(0))
            index79_ = default__.safeIndex(605, (d_627_v65_).length(0))
            rhs73_ = default__.fm32(p0, globalState)
            rhs74_ = _dafny.MultiSet([p0, (self).fm24(globalState), (self).f41])
            rhs75_ = d_536_v3_
            rhs76_ = (d_628_v66_ if p0 else d_628_v66_)
            lhs55_ = d_624_v64_
            lhs56_ = default__.safeIndex(261, (d_624_v64_).length(0))
            lhs57_ = d_627_v65_
            lhs58_ = default__.safeIndex(605, (d_627_v65_).length(0))
            d_623_v63_ = rhs73_
            lhs55_[lhs56_] = rhs74_
            d_536_v3_ = rhs75_
            lhs57_[lhs58_] = rhs76_
            d_629_v67_: _dafny.Map
            d_629_v67_ = _dafny.Map({(self).f41: _dafny.Map({(self).f42: d_533_v1_})})
            d_630_v68_: D11
            d_630_v68_ = D11_DC26(d_533_v1_)
            d_631_v69_: _dafny.Array
            nw102_ = _dafny.Array(None, 29)
            nw102_[int(0)] = d_535_v2_
            nw102_[int(1)] = _dafny.Map({(self).f42: d_533_v1_})
            nw102_[int(2)] = d_535_v2_
            nw102_[int(3)] = (d_535_v2_) | (((d_629_v67_)[(self).f41] if ((self).f41) in (d_629_v67_) else d_535_v2_))
            nw102_[int(4)] = (d_535_v2_) | (_dafny.Map({880: d_533_v1_}))
            nw102_[int(5)] = d_535_v2_
            nw102_[int(6)] = (d_535_v2_) | (d_535_v2_)
            nw102_[int(7)] = (d_535_v2_).set((self).f42, (d_630_v68_).cf39)
            nw102_[int(8)] = d_535_v2_
            nw102_[int(9)] = d_535_v2_
            nw102_[int(10)] = (d_535_v2_ if p0 else _dafny.Map({(0) - ((0) - ((self).f42)): d_533_v1_}))
            nw102_[int(11)] = d_535_v2_
            nw102_[int(12)] = (d_535_v2_) | (d_535_v2_)
            nw102_[int(13)] = d_535_v2_
            nw102_[int(14)] = d_535_v2_
            nw102_[int(15)] = (d_535_v2_ if p0 else d_535_v2_)
            nw102_[int(16)] = (d_535_v2_ if p0 else d_535_v2_)
            nw102_[int(17)] = d_535_v2_
            nw102_[int(18)] = d_535_v2_
            nw102_[int(19)] = d_535_v2_
            nw102_[int(20)] = _dafny.Map({(self).f42: d_533_v1_})
            nw102_[int(21)] = d_535_v2_
            nw102_[int(22)] = _dafny.Map({(self).f42: d_533_v1_})
            nw102_[int(23)] = (d_535_v2_) | (d_535_v2_)
            nw102_[int(24)] = _dafny.Map({(self).f42: d_533_v1_})
            nw102_[int(25)] = d_535_v2_
            nw102_[int(26)] = (_dafny.Map({(self).f42: d_533_v1_})).set((self).f42, d_533_v1_)
            nw102_[int(27)] = d_535_v2_
            nw102_[int(28)] = d_535_v2_
            d_631_v69_ = nw102_
            index80_ = default__.safeIndex(12, (d_631_v69_).length(0))
            (d_631_v69_)[index80_] = d_535_v2_
            index81_ = default__.safeIndex(12, (d_631_v69_).length(0))
            (d_631_v69_)[index81_] = d_535_v2_
            d_632_v70_: _dafny.Map
            d_632_v70_ = _dafny.Map({d_533_v1_: d_536_v3_})
            d_633_v71_: _dafny.Seq
            d_633_v71_ = _dafny.SeqWithoutIsStrInference([d_628_v66_, (d_627_v65_)[default__.safeIndex(605, (d_627_v65_).length(0))], _dafny.SeqWithoutIsStrInference([d_536_v3_ for d_634_i7_ in range(default__.abs(851))])])
            d_635_v72_: _dafny.Seq
            d_635_v72_ = _dafny.SeqWithoutIsStrInference([d_633_v71_])
            index82_ = default__.safeIndex(605, (d_627_v65_).length(0))
            rhs77_ = True
            rhs78_ = d_632_v70_
            rhs79_ = (d_635_v72_)[default__.safeIndex(775, len(d_635_v72_))]
            rhs80_ = (d_628_v66_) + ((d_628_v66_) + ((d_627_v65_)[default__.safeIndex(605, (d_627_v65_).length(0))]))
            lhs59_ = d_627_v65_
            lhs60_ = default__.safeIndex(605, (d_627_v65_).length(0))
            r2 = rhs77_
            d_632_v70_ = rhs78_
            d_633_v71_ = rhs79_
            lhs59_[lhs60_] = rhs80_
            index83_ = default__.safeIndex(228, (d_543_v10_).length(0))
            (d_543_v10_)[index83_] = (self).f42
            index84_ = default__.safeIndex(228, (d_543_v10_).length(0))
            (d_543_v10_)[index84_] = (self).f42
        (globalState).f13 = p0
        (globalState).f13 = ((self).f42) in (_dafny.SeqWithoutIsStrInference([(0) - ((self).f42)]))
        d_636_v73_: _dafny.Seq
        d_636_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
        (globalState).f22 = (default__.fm2(globalState)) * ((915) - (len(d_636_v73_)))
        d_637_v74_: _dafny.Array
        nw103_ = _dafny.Array(_dafny.Set({}), 29)
        d_637_v74_ = nw103_
        d_638_v75_: D7
        d_638_v75_ = D7_DC15(d_637_v74_)
        r0 = d_638_v75_
        r1 = (self).f42
        r2 = (self).f41
        return r0, r1, r2

    @property
    def f41(self):
        return self._f41
    @property
    def f42(self):
        return self._f42

class C3(T0):
    def  __init__(self):
        self._f26: _dafny.Map = _dafny.Map({})
        self.f40: bool = False
        self._f39: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f26(self):
        return self._f26
    @f26.setter
    def f26(self, value):
        self._f26 = value
    def ctor__(self, f39, f40, f26):
        (self)._f39 = f39
        (self).f40 = f40
        (self).f26 = f26

    def fm3(self, globalState):
        return 646

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife85_():
            coll45_ = _dafny.Set()
            compr_45_: int
            for compr_45_ in _dafny.IntegerRange(423, 193):
                d_639_v0_: int = compr_45_
                if ((423) <= (d_639_v0_)) and ((d_639_v0_) < (193)):
                    coll45_ = coll45_.union(_dafny.Set([(d_639_v0_) * (902)]))
            return _dafny.Set(coll45_)
        return _dafny.Map({default__.safeDivisionInt(-64, len(_dafny.SeqWithoutIsStrInference([iife85_()
        , _dafny.Set({(0) - (-80)})]))): _dafny.Map({_dafny.CodePoint('m'): len(_dafny.Set({self.f40, self.f40, self.f40, self.f40}))})})

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        d_640_v0_: D2
        d_640_v0_ = D2_DC5()
        d_641_v1_: _dafny.Array
        nw104_ = _dafny.Array(D4.default()(), 4)
        d_641_v1_ = nw104_
        d_642_v2_: _dafny.Seq
        d_642_v2_ = _dafny.SeqWithoutIsStrInference([p2])
        d_643_v3_: _dafny.Seq
        d_643_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
        d_644_v4_: _dafny.Array
        nw105_ = _dafny.Array(None, 6)
        nw105_[int(0)] = (len(self.f26)) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nolfdh"))))
        nw105_[int(1)] = (d_643_v3_) != (d_643_v3_)
        nw105_[int(2)] = default__.fm0(p1, globalState)
        nw105_[int(3)] = self.f40
        nw105_[int(4)] = self.f40
        nw105_[int(5)] = self.f40
        d_644_v4_ = nw105_
        index85_ = default__.safeIndex(159, (d_644_v4_).length(0))
        (d_644_v4_)[index85_] = not (not(self.f40)) or (True)
        d_645_v5_: D7
        d_645_v5_ = D7_DC17(self.f40, self.f40, p2)
        pat_let_tv15_ = d_640_v0_
        pat_let_tv16_ = d_640_v0_
        pat_let_tv17_ = d_640_v0_
        index86_ = default__.safeIndex(159, (d_644_v4_).length(0))
        def lambda34_(source10_):
            if source10_.is_DC16:
                d_646___mcc_h0_ = source10_.cf22
                d_647___mcc_h1_ = source10_.cf23
                d_648_cf23_ = d_647___mcc_h1_
                d_649_cf22_ = d_646___mcc_h0_
                return pat_let_tv15_
            elif source10_.is_DC17:
                d_650___mcc_h2_ = source10_.cf24
                d_651___mcc_h3_ = source10_.cf25
                d_652___mcc_h4_ = source10_.cf26
                d_653_cf26_ = d_652___mcc_h4_
                d_654_cf25_ = d_651___mcc_h3_
                d_655_cf24_ = d_650___mcc_h2_
                return pat_let_tv16_
            elif True:
                d_656___mcc_h5_ = source10_.cf21
                d_657_cf21_ = d_656___mcc_h5_
                return pat_let_tv17_

        rhs81_ = lambda34_(d_645_v5_)
        rhs82_ = self.f40
        rhs83_ = (d_641_v1_ if self.f40 else d_641_v1_)
        rhs84_ = d_642_v2_
        rhs85_ = False
        lhs61_ = globalState
        lhs62_ = d_644_v4_
        lhs63_ = default__.safeIndex(159, (d_644_v4_).length(0))
        d_640_v0_ = rhs81_
        lhs61_.f1 = rhs82_
        d_641_v1_ = rhs83_
        d_642_v2_ = rhs84_
        lhs62_[lhs63_] = rhs85_
        index87_ = default__.safeIndex(189, ((self).f39).length(0))
        ((self).f39)[index87_] = default__.safeDivisionInt(625, p1)
        index88_ = default__.safeIndex(189, ((self).f39).length(0))
        ((self).f39)[index88_] = (default__.fm23(globalState)).cf19
        d_658_i0_: int
        d_658_i0_ = 0
        with _dafny.label("3"):
            while self.f40:
                with _dafny.c_label("3"):
                    if (d_658_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_658_i0_ = (d_658_i0_) + (1)
                    d_659_v6_: _dafny.Seq
                    d_659_v6_ = _dafny.SeqWithoutIsStrInference([(d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]])
                    d_660_v7_: _dafny.Map
                    d_660_v7_ = _dafny.Map({p1: (d_659_v6_) + (d_659_v6_)})
                    index89_ = default__.safeIndex(159, (d_644_v4_).length(0))
                    rhs86_ = ((_dafny.Map({((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]: _dafny.SeqWithoutIsStrInference([(d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]])})) | (d_660_v7_)) | (d_660_v7_)
                    rhs87_ = (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]
                    lhs64_ = d_644_v4_
                    lhs65_ = default__.safeIndex(159, (d_644_v4_).length(0))
                    d_660_v7_ = rhs86_
                    lhs64_[lhs65_] = rhs87_
                    d_661_v8_: _dafny.Array
                    nw106_ = _dafny.Array(_dafny.Array(None, 0), 10)
                    d_661_v8_ = nw106_
                    d_662_v9_: _dafny.Array
                    nw107_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                    d_662_v9_ = nw107_
                    index90_ = default__.safeIndex(515, (d_661_v8_).length(0))
                    (d_661_v8_)[index90_] = d_662_v9_
                    index91_ = default__.safeIndex(515, (d_661_v8_).length(0))
                    nw108_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                    (d_661_v8_)[index91_] = nw108_
                    index92_ = default__.safeIndex(159, (d_644_v4_).length(0))
                    (d_644_v4_)[index92_] = (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]
                    d_663_v10_: _dafny.MultiSet
                    d_663_v10_ = _dafny.MultiSet([p2])
                    if ((d_663_v10_).intersection(d_663_v10_)).issubset(d_663_v10_):
                        d_664_v11_: _dafny.Array
                        nw109_ = _dafny.Array(_dafny.MultiSet({}), 5)
                        d_664_v11_ = nw109_
                        d_665_v12_: D2
                        d_665_v12_ = D2_DC4(d_664_v11_)
                        pat_let_tv18_ = d_664_v11_
                        pat_let_tv19_ = d_664_v11_
                        pat_let_tv20_ = d_664_v11_
                        d_666_v13_: _dafny.Array
                        nw110_ = _dafny.Array(None, 16)
                        nw110_[int(0)] = D2_DC4(d_664_v11_)
                        nw110_[int(1)] = d_665_v12_
                        nw110_[int(2)] = d_665_v12_
                        nw110_[int(3)] = d_665_v12_
                        def iife86_(_pat_let20_0):
                            def iife87_(d_667_dt__update__tmp_h0_):
                                def iife88_(_pat_let21_0):
                                    def iife89_(d_668_dt__update_hcf6_h0_):
                                        return D2_DC4(d_668_dt__update_hcf6_h0_)
                                    return iife89_(_pat_let21_0)
                                return iife88_(pat_let_tv18_)
                            return iife87_(_pat_let20_0)
                        nw110_[int(4)] = iife86_(d_665_v12_)
                        def iife90_(_pat_let22_0):
                            def iife91_(d_669_dt__update__tmp_h1_):
                                def iife92_(_pat_let23_0):
                                    def iife93_(d_670_dt__update_hcf6_h1_):
                                        return D2_DC4(d_670_dt__update_hcf6_h1_)
                                    return iife93_(_pat_let23_0)
                                return iife92_(pat_let_tv19_)
                            return iife91_(_pat_let22_0)
                        nw110_[int(5)] = iife90_(d_665_v12_)
                        nw110_[int(6)] = d_665_v12_
                        nw110_[int(7)] = d_665_v12_
                        nw110_[int(8)] = d_665_v12_
                        def iife94_(_pat_let24_0):
                            def iife95_(d_671_dt__update__tmp_h2_):
                                def iife96_(_pat_let25_0):
                                    def iife97_(d_672_dt__update_hcf6_h2_):
                                        return D2_DC4(d_672_dt__update_hcf6_h2_)
                                    return iife97_(_pat_let25_0)
                                return iife96_(pat_let_tv20_)
                            return iife95_(_pat_let24_0)
                        nw110_[int(9)] = (d_665_v12_ if (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))] else iife94_(d_665_v12_))
                        nw110_[int(10)] = d_665_v12_
                        nw110_[int(11)] = d_665_v12_
                        nw110_[int(12)] = d_665_v12_
                        nw110_[int(13)] = (D2_DC4(d_664_v11_) if (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))] else d_665_v12_)
                        nw110_[int(14)] = d_665_v12_
                        nw110_[int(15)] = D2_DC4(d_664_v11_)
                        d_666_v13_ = nw110_
                        index93_ = default__.safeIndex(896, (d_666_v13_).length(0))
                        (d_666_v13_)[index93_] = D2_DC4(d_664_v11_)
                        pat_let_tv21_ = d_664_v11_
                        index94_ = default__.safeIndex(896, (d_666_v13_).length(0))
                        def iife98_(_pat_let26_0):
                            def iife99_(d_673_dt__update__tmp_h3_):
                                def iife100_(_pat_let27_0):
                                    def iife101_(d_674_dt__update_hcf6_h3_):
                                        return D2_DC4(d_674_dt__update_hcf6_h3_)
                                    return iife101_(_pat_let27_0)
                                return iife100_(pat_let_tv21_)
                            return iife99_(_pat_let26_0)
                        (d_666_v13_)[index94_] = iife98_(d_665_v12_)
                        d_675_v14_: _dafny.Map
                        d_675_v14_ = _dafny.Map({(d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]: (0) - (p2)})
                        d_676_v15_: C1
                        nw111_ = C1()
                        nw111_.ctor__((d_675_v14_) | (default__.fm33(d_642_v2_, (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))], globalState)))
                        d_676_v15_ = nw111_
                        (globalState).f0 = 425
                        d_677_v16_: _dafny.MultiSet
                        d_677_v16_ = _dafny.MultiSet([False])
                        d_678_v17_: D5
                        d_678_v17_ = D5_DC11(self.f40)
                        d_679_v18_: _dafny.Seq
                        d_679_v18_ = _dafny.SeqWithoutIsStrInference([(d_677_v16_).intersection(d_677_v16_), d_677_v16_, ((d_677_v16_).set(self.f40, default__.abs(((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]))) | (d_677_v16_), (_dafny.MultiSet([(d_678_v17_).cf16, self.f40])).set((d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))], default__.abs(p1))])
                        d_680_v19_: str
                        d_680_v19_ = _dafny.CodePoint('s')
                        pat_let_tv22_ = d_680_v19_
                        d_681_v20_: _dafny.Map
                        def iife102_(_pat_let28_0):
                            def iife103_(d_682_dt__update__tmp_h4_):
                                def iife104_(_pat_let29_0):
                                    def iife105_(d_683_dt__update_hcf13_h0_):
                                        return D4_DC8(d_683_dt__update_hcf13_h0_)
                                    return iife105_(_pat_let29_0)
                                return iife104_(pat_let_tv22_)
                            return iife103_(_pat_let28_0)
                        d_681_v20_ = _dafny.Map({d_644_v4_: iife102_(D4_DC8(d_680_v19_))})
                        rhs88_ = d_679_v18_
                        rhs89_ = default__.fm0((d_642_v2_)[default__.safeIndex(332, len(d_642_v2_))], globalState)
                        rhs90_ = d_642_v2_
                        rhs91_ = len(((default__.fm30(p1, d_643_v3_, len(default__.fm38(len(d_659_v6_), p2, p1, self.f40, globalState)), globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfpkd")))).set(default__.safeIndex((d_642_v2_)[default__.safeIndex(p2, len(d_642_v2_))], len((default__.fm30(p1, d_643_v3_, len(default__.fm38(len(d_659_v6_), p2, p1, self.f40, globalState)), globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfpkd"))))), d_680_v19_))
                        rhs92_ = d_681_v20_
                        lhs66_ = self
                        d_679_v18_ = rhs88_
                        lhs66_.f40 = rhs89_
                        d_642_v2_ = rhs90_
                        r2 = rhs91_
                        d_681_v20_ = rhs92_
                        d_684_v21_: _dafny.Array
                        nw112_ = _dafny.Array(_dafny.Map({}), 29)
                        d_684_v21_ = nw112_
                        rhs93_ = d_684_v21_
                        rhs94_ = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        lhs67_ = globalState
                        d_684_v21_ = rhs93_
                        lhs67_.f22 = rhs94_
                    elif True:
                        d_685_v22_: _dafny.Map
                        d_685_v22_ = _dafny.Map({(d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]: p2})
                        d_686_v23_: _dafny.Map
                        d_686_v23_ = _dafny.Map({((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]: d_685_v22_})
                        d_687_v24_: D3
                        d_687_v24_ = D3_DC7(d_643_v3_, p1, len(default__.fm30(p1, d_643_v3_, p1, globalState)), len(d_643_v3_), d_659_v6_)
                        d_688_v25_: _dafny.Seq
                        d_688_v25_ = _dafny.SeqWithoutIsStrInference([(d_687_v24_).cf12, d_659_v6_])
                        d_686_v23_ = (d_686_v23_).set(len((d_659_v6_ if (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))] else (d_688_v25_)[default__.safeIndex(p2, len(d_688_v25_))])), (d_685_v22_) | (d_685_v22_))
                        d_689_v26_: C1
                        nw113_ = C1()
                        nw113_.ctor__((_dafny.Map({(d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]: ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]})) | (d_685_v22_))
                        d_689_v26_ = nw113_
                        d_690_v27_: _dafny.Array
                        def lambda35_(d_691_v10_):
                            def lambda36_(d_692_i1_):
                                return d_691_v10_

                            return lambda36_

                        init21_ = lambda35_(d_663_v10_)
                        nw114_ = _dafny.Array(None, 15)
                        for i0_21_ in range(nw114_.length(0)):
                            nw114_[i0_21_] = init21_(i0_21_)
                        d_690_v27_ = nw114_
                        d_690_v27_ = d_690_v27_
                        d_693_v28_: _dafny.MultiSet
                        d_693_v28_ = _dafny.MultiSet([(d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]])
                        d_694_v29_: _dafny.Map
                        d_694_v29_ = _dafny.Map({d_693_v28_: p1})
                        d_695_v30_: _dafny.Set
                        d_695_v30_ = _dafny.Set({p2, len(d_643_v3_)})
                        d_696_v31_: _dafny.Array
                        nw115_ = _dafny.Array(None, 22)
                        nw115_[int(0)] = len(self.f26)
                        nw115_[int(1)] = 850
                        nw115_[int(2)] = len(d_695_v30_)
                        nw115_[int(3)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw115_[int(4)] = len(d_643_v3_)
                        nw115_[int(5)] = 747
                        nw115_[int(6)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw115_[int(7)] = p2
                        nw115_[int(8)] = p1
                        nw115_[int(9)] = p1
                        nw115_[int(10)] = len(d_685_v22_)
                        nw115_[int(11)] = (d_642_v2_)[default__.safeIndex(p1, len(d_642_v2_))]
                        nw115_[int(12)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw115_[int(13)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw115_[int(14)] = p1
                        nw115_[int(15)] = (d_693_v28_).cardinality
                        nw115_[int(16)] = len(d_659_v6_)
                        nw115_[int(17)] = p2
                        nw115_[int(18)] = (_dafny.MultiSet(d_642_v2_)).cardinality
                        nw115_[int(19)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw115_[int(20)] = p2
                        nw115_[int(21)] = 995
                        d_696_v31_ = nw115_
                        d_697_v32_: str
                        d_697_v32_ = _dafny.CodePoint('h')
                        d_698_v33_: D0
                        d_698_v33_ = D0_DC1(d_696_v31_, d_697_v32_, p1)
                        d_699_v34_: C2
                        nw116_ = C2()
                        nw116_.ctor__(True, (0) - (p2))
                        d_699_v34_ = nw116_
                        d_700_v35_: _dafny.Array
                        nw117_ = _dafny.Array(None, 29)
                        nw117_[int(0)] = d_699_v34_
                        nw117_[int(1)] = d_699_v34_
                        nw117_[int(2)] = d_699_v34_
                        nw117_[int(3)] = d_699_v34_
                        nw117_[int(4)] = d_699_v34_
                        nw117_[int(5)] = d_699_v34_
                        nw117_[int(6)] = d_699_v34_
                        nw117_[int(7)] = d_699_v34_
                        nw117_[int(8)] = d_699_v34_
                        nw117_[int(9)] = d_699_v34_
                        nw117_[int(10)] = d_699_v34_
                        nw117_[int(11)] = d_699_v34_
                        nw117_[int(12)] = d_699_v34_
                        nw117_[int(13)] = d_699_v34_
                        nw117_[int(14)] = d_699_v34_
                        nw117_[int(15)] = d_699_v34_
                        nw117_[int(16)] = d_699_v34_
                        nw117_[int(17)] = d_699_v34_
                        nw117_[int(18)] = d_699_v34_
                        nw117_[int(19)] = d_699_v34_
                        nw117_[int(20)] = d_699_v34_
                        nw117_[int(21)] = d_699_v34_
                        nw117_[int(22)] = d_699_v34_
                        nw117_[int(23)] = d_699_v34_
                        nw117_[int(24)] = d_699_v34_
                        nw117_[int(25)] = d_699_v34_
                        nw117_[int(26)] = d_699_v34_
                        nw117_[int(27)] = d_699_v34_
                        nw117_[int(28)] = d_699_v34_
                        d_700_v35_ = nw117_
                        d_701_v36_: _dafny.Map
                        d_701_v36_ = _dafny.Map({p1: d_700_v35_})
                        d_702_v37_: _dafny.Map
                        d_702_v37_ = _dafny.Map({p1: d_644_v4_})
                        d_703_v38_: D12
                        d_703_v38_ = D12_DC28((d_693_v28_).set((d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))], default__.abs(p2)))
                        d_704_v39_: _dafny.Array
                        nw118_ = _dafny.Array(None, 27)
                        nw118_[int(0)] = p1
                        nw118_[int(1)] = ((d_694_v29_)[_dafny.MultiSet([self.f40, (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))], (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))], False, (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]])] if (_dafny.MultiSet([self.f40, (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))], (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))], False, (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]])) in (d_694_v29_) else p1)
                        nw118_[int(2)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw118_[int(3)] = (d_698_v33_).cf3
                        nw118_[int(4)] = 772
                        nw118_[int(5)] = (0) - (p1)
                        nw118_[int(6)] = p1
                        nw118_[int(7)] = default__.fm2(globalState)
                        nw118_[int(8)] = p1
                        nw118_[int(9)] = 456
                        nw118_[int(10)] = p1
                        nw118_[int(11)] = p2
                        nw118_[int(12)] = p2
                        nw118_[int(13)] = (_dafny.MultiSet([self.f40])).cardinality
                        nw118_[int(14)] = len(d_701_v36_)
                        nw118_[int(15)] = (0) - (len(d_702_v37_))
                        nw118_[int(16)] = p1
                        nw118_[int(17)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw118_[int(18)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oylgyhryt")))
                        nw118_[int(19)] = ((d_703_v38_).cf43).cardinality
                        nw118_[int(20)] = p1
                        nw118_[int(21)] = (self).fm3(globalState)
                        nw118_[int(22)] = len(d_643_v3_)
                        nw118_[int(23)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw118_[int(24)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw118_[int(25)] = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
                        nw118_[int(26)] = -717
                        d_704_v39_ = nw118_
                        d_705_v40_: _dafny.Seq
                        d_705_v40_ = _dafny.SeqWithoutIsStrInference([d_696_v31_, d_704_v39_])
                        rhs95_ = default__.safeDivisionInt(p1, default__.safeModuloInt(p1, (d_699_v34_).f42))
                        rhs96_ = d_705_v40_
                        rhs97_ = (d_699_v34_).f42
                        lhs68_ = globalState
                        lhs69_ = globalState
                        lhs68_.f21 = rhs95_
                        d_705_v40_ = rhs96_
                        lhs69_.f21 = rhs97_
                        d_705_v40_ = d_705_v40_
                    pass
            pass
        d_706_v41_: str
        d_706_v41_ = _dafny.CodePoint('i')
        d_707_v42_: _dafny.Map
        d_707_v42_ = _dafny.Map({True: (default__.fm39(p2, globalState)).cf42})
        d_708_v43_: _dafny.Array
        nw119_ = _dafny.Array(None, 20)
        nw119_[int(0)] = d_706_v41_
        nw119_[int(1)] = d_706_v41_
        nw119_[int(2)] = d_706_v41_
        nw119_[int(3)] = _dafny.CodePoint('c')
        nw119_[int(4)] = d_706_v41_
        nw119_[int(5)] = d_706_v41_
        nw119_[int(6)] = d_706_v41_
        nw119_[int(7)] = d_706_v41_
        nw119_[int(8)] = d_706_v41_
        nw119_[int(9)] = d_706_v41_
        nw119_[int(10)] = d_706_v41_
        nw119_[int(11)] = _dafny.CodePoint('r')
        nw119_[int(12)] = d_706_v41_
        nw119_[int(13)] = d_706_v41_
        nw119_[int(14)] = d_706_v41_
        nw119_[int(15)] = d_706_v41_
        nw119_[int(16)] = d_706_v41_
        nw119_[int(17)] = (d_643_v3_)[default__.safeIndex(len(d_707_v42_), len(d_643_v3_))]
        nw119_[int(18)] = d_706_v41_
        nw119_[int(19)] = (D4_DC8(d_706_v41_)).cf13
        d_708_v43_ = nw119_
        d_709_v44_: C0
        nw120_ = C0()
        nw120_.ctor__(d_708_v43_)
        d_709_v44_ = nw120_
        d_710_v45_: _dafny.Seq
        d_710_v45_ = _dafny.SeqWithoutIsStrInference([(d_643_v3_ if not((d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]) else d_643_v3_), (d_643_v3_).set(default__.safeIndex(((self).f39)[default__.safeIndex(189, ((self).f39).length(0))], len(d_643_v3_)), d_706_v41_)])
        rhs98_ = (d_644_v4_)[default__.safeIndex(159, (d_644_v4_).length(0))]
        rhs99_ = d_710_v45_
        rhs100_ = p1
        lhs70_ = self
        lhs71_ = globalState
        lhs70_.f40 = rhs98_
        d_710_v45_ = rhs99_
        lhs71_.f22 = rhs100_
        d_711_v46_: _dafny.MultiSet
        d_711_v46_ = _dafny.MultiSet([((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]])
        d_712_v47_: _dafny.Array
        nw121_ = _dafny.Array(None, 6)
        nw121_[int(0)] = len(_dafny.Set({p1, (0) - (p2)}))
        nw121_[int(1)] = len(_dafny.SeqWithoutIsStrInference([((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]]))
        nw121_[int(2)] = p2
        nw121_[int(3)] = p2
        nw121_[int(4)] = default__.safeDivisionInt((d_711_v46_).cardinality, p2)
        nw121_[int(5)] = p2
        d_712_v47_ = nw121_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_712_v47_).length(0)):
            d_713_i2_: int = guard_loop_3_
            if (True) and (((0) <= (d_713_i2_)) and ((d_713_i2_) < ((d_712_v47_).length(0)))):
                (d_712_v47_)[(d_713_i2_)] = (d_713_i2_) * (p2)
        r0 = ((self).f39)[default__.safeIndex(189, ((self).f39).length(0))]
        d_714_v48_: _dafny.Map
        d_714_v48_ = _dafny.Map({d_643_v3_: d_644_v4_})
        r1 = d_714_v48_
        r2 = (p2 if self.f40 else (p2) + (len(d_643_v3_)))
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        d_715_v0_: _dafny.Seq
        d_715_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ld"))
        d_716_v1_: int
        d_716_v1_ = -799
        d_717_v2_: _dafny.Seq
        d_717_v2_ = _dafny.SeqWithoutIsStrInference([d_716_v1_, d_716_v1_])
        (globalState).f16 = ((d_717_v2_)[default__.safeIndex((d_717_v2_)[default__.safeIndex(d_716_v1_, len(d_717_v2_))], len(d_717_v2_))] if (d_715_v0_) != (d_715_v0_) else d_716_v1_)
        d_718_v3_: _dafny.Seq
        d_718_v3_ = _dafny.SeqWithoutIsStrInference([self.f40, self.f40])
        d_719_v4_: _dafny.Map
        d_719_v4_ = _dafny.Map({(d_715_v0_) + (d_715_v0_): _dafny.MultiSet([(d_718_v3_)[default__.safeIndex(d_716_v1_, len(d_718_v3_))]])})
        d_719_v4_ = d_719_v4_
        d_720_v5_: _dafny.Seq
        d_720_v5_ = d_717_v2_
        d_721_v6_: _dafny.Seq
        d_721_v6_ = _dafny.SeqWithoutIsStrInference([d_720_v5_, d_717_v2_, d_720_v5_, d_720_v5_, d_720_v5_])
        d_722_v7_: _dafny.Array
        nw122_ = _dafny.Array(None, 6)
        nw122_[int(0)] = d_721_v6_
        nw122_[int(1)] = d_721_v6_
        nw122_[int(2)] = (d_721_v6_) + (_dafny.SeqWithoutIsStrInference([d_720_v5_ for d_723_i1_ in range(default__.abs(293))]))
        nw122_[int(3)] = (d_721_v6_) + (_dafny.SeqWithoutIsStrInference([d_720_v5_ for d_724_i2_ in range(default__.abs(741))]))
        nw122_[int(4)] = (d_721_v6_) + (_dafny.SeqWithoutIsStrInference([d_720_v5_ for d_725_i3_ in range(default__.abs(952))]))
        nw122_[int(5)] = (d_721_v6_) + (d_721_v6_)
        d_722_v7_ = nw122_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_722_v7_).length(0)):
            d_726_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_726_i0_)) and ((d_726_i0_) < ((d_722_v7_).length(0)))):
                (d_722_v7_)[(d_726_i0_)] = d_721_v6_
        d_727_v8_: _dafny.Map
        d_727_v8_ = _dafny.Map({d_716_v1_: d_717_v2_})
        r2 = (d_716_v1_) not in ((d_717_v2_) + (((d_727_v8_)[len(d_718_v3_)] if (len(d_718_v3_)) in (d_727_v8_) else d_717_v2_)))
        d_728_v9_: str
        d_728_v9_ = _dafny.CodePoint('p')
        d_729_v10_: D4
        d_729_v10_ = D4_DC8(d_728_v9_)
        pat_let_tv23_ = d_728_v9_
        def iife106_(_pat_let30_0):
            def iife107_(d_730_dt__update__tmp_h1_):
                def iife108_(_pat_let31_0):
                    def iife109_(d_731_dt__update_hcf13_h0_):
                        return D4_DC8(d_731_dt__update_hcf13_h0_)
                    return iife109_(_pat_let31_0)
                return iife108_(pat_let_tv23_)
            return iife107_(_pat_let30_0)
        d_729_v10_ = iife106_(d_729_v10_)
        (globalState).f21 = len(d_715_v0_)
        r0 = default__.fm0(-206, globalState)
        d_732_v11_: _dafny.Set
        d_732_v11_ = _dafny.Set({self.f26, self.f26})
        r1 = d_732_v11_
        d_733_v12_: _dafny.Map
        d_733_v12_ = _dafny.Map({d_716_v1_: _dafny.MultiSet([self.f40, self.f40])})
        d_734_v13_: _dafny.Map
        d_734_v13_ = _dafny.Map({len(d_733_v12_): self.f40})
        d_735_v14_: _dafny.Map
        d_735_v14_ = _dafny.Map({self.f40: d_716_v1_})
        d_736_v15_: _dafny.Map
        d_736_v15_ = _dafny.Map({d_734_v13_: d_735_v14_})
        d_737_v16_: D9
        d_737_v16_ = D9_DC22(d_736_v15_)
        def lambda37_(source11_):
            if source11_.is_DC23:
                return self.f40
            elif True:
                d_738___mcc_h0_ = source11_.cf33
                d_739_cf33_ = d_738___mcc_h0_
                return (_dafny.Set({False})).issubset(_dafny.Set({self.f40}))

        r2 = lambda37_((d_737_v16_ if self.f40 else d_737_v16_))
        return r0, r1, r2

    @property
    def f39(self):
        return self._f39

class C4:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm22(self, p0, globalState):
        return _dafny.CodePoint('y')

    def m17(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Set = _dafny.Set({})
        r3: bool = False
        d_740_v0_: _dafny.Array
        nw123_ = _dafny.Array(_dafny.Array(None, 0), 2)
        d_740_v0_ = nw123_
        d_740_v0_ = d_740_v0_
        d_741_v1_: int
        d_741_v1_ = 121
        (globalState).f22 = d_741_v1_
        d_742_v2_: _dafny.Map
        d_742_v2_ = _dafny.Map({p0: 385})
        d_743_v3_: C1
        nw124_ = C1()
        nw124_.ctor__(d_742_v2_)
        d_743_v3_ = nw124_
        if p0:
            (globalState).f2 = d_741_v1_
            d_744_v4_: str
            d_744_v4_ = _dafny.CodePoint('u')
            d_745_v5_: _dafny.Seq
            d_745_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "woynvoj"))
            d_746_v6_: _dafny.Seq
            d_746_v6_ = _dafny.SeqWithoutIsStrInference([p0, False])
            d_747_v7_: _dafny.Map
            d_747_v7_ = _dafny.Map({d_744_v4_: p0})
            d_748_v8_: _dafny.MultiSet
            d_748_v8_ = _dafny.MultiSet([p0])
            d_749_v9_: _dafny.Set
            d_749_v9_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ueiucop")), d_745_v5_, d_745_v5_, d_745_v5_, d_745_v5_})
            d_750_v10_: C2
            nw125_ = C2()
            nw125_.ctor__((d_744_v4_) not in (d_745_v5_), ((d_748_v8_).cardinality if (d_746_v6_)[default__.safeIndex(len(d_747_v7_), len(d_746_v6_))] else len(d_749_v9_)))
            d_750_v10_ = nw125_
            d_750_v10_ = d_750_v10_
            (globalState).f1 = (((d_750_v10_).f42) < (d_741_v1_)) in ((d_743_v3_.f43) | (d_742_v2_))
            (globalState).f21 = d_741_v1_
            (globalState).f13 = (d_750_v10_).f41
        elif True:
            d_751_v11_: C2
            nw126_ = C2()
            nw126_.ctor__(p0, default__.safeModuloInt(default__.fm2(globalState), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_752_i0_ in range(default__.abs(41))]))))
            d_751_v11_ = nw126_
            r3 = not((d_751_v11_).f41)
            (globalState).f22 = -426
            d_753_v12_: _dafny.Seq
            d_753_v12_ = _dafny.SeqWithoutIsStrInference([d_741_v1_, d_741_v1_, (d_751_v11_).f42])
            d_754_v13_: _dafny.Array
            nw127_ = _dafny.Array(None, 19)
            nw127_[int(0)] = d_741_v1_
            nw127_[int(1)] = ((d_743_v3_.f43)[(d_751_v11_).f41] if ((d_751_v11_).f41) in (d_743_v3_.f43) else (d_751_v11_).f42)
            nw127_[int(2)] = (d_751_v11_).f42
            nw127_[int(3)] = len(d_753_v12_)
            nw127_[int(4)] = len(d_753_v12_)
            nw127_[int(5)] = d_741_v1_
            nw127_[int(6)] = (d_751_v11_).f42
            nw127_[int(7)] = (d_751_v11_).f42
            nw127_[int(8)] = default__.fm2(globalState)
            nw127_[int(9)] = (d_751_v11_).f42
            nw127_[int(10)] = 263
            nw127_[int(11)] = len(d_753_v12_)
            nw127_[int(12)] = (d_751_v11_).f42
            nw127_[int(13)] = (d_751_v11_).f42
            nw127_[int(14)] = d_741_v1_
            nw127_[int(15)] = d_741_v1_
            nw127_[int(16)] = (d_751_v11_).f42
            nw127_[int(17)] = (d_751_v11_).f42
            nw127_[int(18)] = d_741_v1_
            d_754_v13_ = nw127_
            d_755_v14_: _dafny.Map
            d_755_v14_ = _dafny.Map({736: len(d_753_v12_)})
            d_756_v15_: C3
            nw128_ = C3()
            nw128_.ctor__(d_754_v13_, ((0) - (default__.fm2(globalState))) > (d_741_v1_), d_755_v14_)
            d_756_v15_ = nw128_
            d_757_v16_: _dafny.Seq
            d_757_v16_ = d_753_v12_
            d_758_v17_: _dafny.Map
            d_758_v17_ = _dafny.Map({d_757_v16_: p0})
            d_759_v18_: _dafny.Seq
            d_759_v18_ = _dafny.SeqWithoutIsStrInference([((d_758_v17_)[d_757_v16_] if (d_757_v16_) in (d_758_v17_) else (d_751_v11_).f41), d_756_v15_.f40, d_756_v15_.f40, d_756_v15_.f40])
            d_760_v19_: _dafny.Array
            nw129_ = _dafny.Array(None, 3)
            nw129_[int(0)] = d_759_v18_
            nw129_[int(1)] = d_759_v18_
            nw129_[int(2)] = (d_759_v18_) + (d_759_v18_)
            d_760_v19_ = nw129_
            index95_ = default__.safeIndex(857, (d_760_v19_).length(0))
            (d_760_v19_)[index95_] = d_759_v18_
            index96_ = default__.safeIndex(857, (d_760_v19_).length(0))
            (d_760_v19_)[index96_] = d_759_v18_
        hi5_ = 300
        for d_761_i1_ in range(len(default__.fm28(p0, p0, d_741_v1_, globalState)), hi5_):
            d_762_v20_: str
            d_762_v20_ = _dafny.CodePoint('u')
            d_762_v20_ = d_762_v20_
            d_763_v21_: _dafny.Array
            nw130_ = _dafny.Array(None, 15)
            nw130_[int(0)] = d_762_v20_
            nw130_[int(1)] = d_762_v20_
            nw130_[int(2)] = d_762_v20_
            nw130_[int(3)] = d_762_v20_
            nw130_[int(4)] = d_762_v20_
            nw130_[int(5)] = d_762_v20_
            nw130_[int(6)] = d_762_v20_
            nw130_[int(7)] = d_762_v20_
            nw130_[int(8)] = d_762_v20_
            nw130_[int(9)] = d_762_v20_
            nw130_[int(10)] = d_762_v20_
            nw130_[int(11)] = d_762_v20_
            nw130_[int(12)] = d_762_v20_
            nw130_[int(13)] = d_762_v20_
            nw130_[int(14)] = d_762_v20_
            d_763_v21_ = nw130_
            index97_ = default__.safeIndex(445, (d_763_v21_).length(0))
            (d_763_v21_)[index97_] = d_762_v20_
            d_764_v22_: _dafny.Seq
            d_764_v22_ = _dafny.SeqWithoutIsStrInference([p0, p0, not(p0), default__.fm0(d_761_i1_, globalState), p0])
            d_765_v23_: _dafny.Set
            d_765_v23_ = _dafny.Set({True})
            d_766_v24_: _dafny.Array
            nw131_ = _dafny.Array(None, 11)
            nw131_[int(0)] = p0
            nw131_[int(1)] = p0
            nw131_[int(2)] = p0
            nw131_[int(3)] = not((_dafny.Set({p0})).isdisjoint(d_765_v23_))
            nw131_[int(4)] = p0
            nw131_[int(5)] = p0
            nw131_[int(6)] = p0
            nw131_[int(7)] = p0
            nw131_[int(8)] = (d_764_v22_) == (_dafny.SeqWithoutIsStrInference([not(p0)]))
            nw131_[int(9)] = (d_765_v23_).ispropersubset(d_765_v23_)
            nw131_[int(10)] = p0
            d_766_v24_ = nw131_
            index98_ = default__.safeIndex(421, (d_766_v24_).length(0))
            (d_766_v24_)[index98_] = not (p0) or (False)
            d_767_v25_: _dafny.MultiSet
            d_767_v25_ = _dafny.MultiSet([True])
            d_768_v26_: _dafny.Seq
            d_768_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))
            index99_ = default__.safeIndex(445, (d_763_v21_).length(0))
            index100_ = default__.safeIndex(421, (d_766_v24_).length(0))
            rhs101_ = d_762_v20_
            rhs102_ = default__.fm28((p0) in (d_767_v25_), p0, d_741_v1_, globalState)
            rhs103_ = (d_761_i1_) * (len(d_768_v26_))
            rhs104_ = (d_765_v23_).issubset(d_765_v23_)
            lhs72_ = d_763_v21_
            lhs73_ = default__.safeIndex(445, (d_763_v21_).length(0))
            lhs74_ = globalState
            lhs75_ = d_766_v24_
            lhs76_ = default__.safeIndex(421, (d_766_v24_).length(0))
            lhs72_[lhs73_] = rhs101_
            d_764_v22_ = rhs102_
            lhs74_.f22 = rhs103_
            lhs75_[lhs76_] = rhs104_
            d_769_v27_: _dafny.MultiSet
            d_769_v27_ = _dafny.MultiSet([d_768_v26_])
            d_770_v28_: _dafny.MultiSet
            d_770_v28_ = d_769_v27_
            if ((d_769_v27_).set((d_768_v26_).set(default__.safeIndex(default__.fm2(globalState), len(d_768_v26_)), d_762_v20_), default__.abs(d_741_v1_))).issubset((d_769_v27_)):
                d_771_v29_: _dafny.Set
                d_771_v29_ = _dafny.Set({d_741_v1_, d_761_i1_, d_741_v1_, 876})
                d_772_v30_: _dafny.Map
                d_772_v30_ = _dafny.Map({d_761_i1_: len(d_768_v26_)})
                d_773_v31_: _dafny.Map
                d_773_v31_ = _dafny.Map({(d_771_v29_).issubset(d_771_v29_): default__.fm0(len(d_772_v30_), globalState)})
                d_774_v32_: _dafny.Seq
                d_774_v32_ = _dafny.SeqWithoutIsStrInference([d_741_v1_])
                d_773_v31_ = (d_773_v31_).set((d_774_v32_) < (d_774_v32_), ((d_766_v24_)[default__.safeIndex(421, (d_766_v24_).length(0))] if (d_766_v24_)[default__.safeIndex(421, (d_766_v24_).length(0))] else p0))
                d_775_v33_: _dafny.Array
                nw132_ = _dafny.Array(int(0), 13)
                d_775_v33_ = nw132_
                index101_ = default__.safeIndex(772, (d_775_v33_).length(0))
                (d_775_v33_)[index101_] = 952
                index102_ = default__.safeIndex(772, (d_775_v33_).length(0))
                (d_775_v33_)[index102_] = (d_741_v1_) + (d_761_i1_)
                d_776_v34_: _dafny.Map
                d_776_v34_ = _dafny.Map({(d_775_v33_)[default__.safeIndex(772, (d_775_v33_).length(0))]: d_762_v20_})
                (globalState).f1 = ((d_761_i1_) not in (d_776_v34_)) not in (d_764_v22_)
                d_777_v35_: _dafny.MultiSet
                d_777_v35_ = _dafny.MultiSet([d_761_i1_, d_741_v1_, d_741_v1_])
                d_778_v36_: _dafny.Map
                d_778_v36_ = _dafny.Map({d_768_v26_: (d_777_v35_).ispropersubset(d_777_v35_)})
                d_778_v36_ = (d_778_v36_).set(d_768_v26_, False)
                (globalState).f13 = (d_766_v24_)[default__.safeIndex(421, (d_766_v24_).length(0))]
            elif True:
                d_779_v37_: _dafny.Array
                nw133_ = _dafny.Array(D5.default()(), 25)
                d_779_v37_ = nw133_
                d_780_v38_: D5
                d_780_v38_ = D5_DC13(d_761_i1_)
                index103_ = default__.safeIndex(478, (d_779_v37_).length(0))
                (d_779_v37_)[index103_] = d_780_v38_
                index104_ = default__.safeIndex(478, (d_779_v37_).length(0))
                (d_779_v37_)[index104_] = d_780_v38_
                d_781_v39_: _dafny.Map
                d_781_v39_ = _dafny.Map({d_741_v1_: default__.fm40(p0, d_741_v1_, globalState)})
                d_782_v40_: _dafny.Map
                d_782_v40_ = _dafny.Map({314: (d_763_v21_)[default__.safeIndex(445, (d_763_v21_).length(0))]})
                d_783_v41_: _dafny.Map
                d_783_v41_ = _dafny.Map({p0: _dafny.Map({d_782_v40_: p0})})
                d_784_v42_: _dafny.Map
                d_784_v42_ = _dafny.Map({d_782_v40_: (d_766_v24_)[default__.safeIndex(421, (d_766_v24_).length(0))]})
                d_785_v43_: _dafny.Seq
                d_785_v43_ = _dafny.SeqWithoutIsStrInference([len(((d_783_v41_)[p0] if (p0) in (d_783_v41_) else d_784_v42_))])
                d_786_v44_: _dafny.Seq
                d_786_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_741_v1_, d_761_i1_]), d_785_v43_])
                d_781_v39_ = (d_781_v39_).set(-23, _dafny.MultiSet(d_786_v44_))
                index105_ = default__.safeIndex(931, (d_740_v0_).length(0))
                (d_740_v0_)[index105_] = d_766_v24_
                d_787_v45_: D11
                d_787_v45_ = D11_DC26(d_766_v24_)
                index106_ = default__.safeIndex(931, (d_740_v0_).length(0))
                (d_740_v0_)[index106_] = (d_787_v45_).cf39
                (globalState).f13 = not(default__.fm0(751, globalState))
                (globalState).f5 = d_761_i1_
            d_788_v46_: _dafny.Array
            nw134_ = _dafny.Array(D9.default()(), 25)
            d_788_v46_ = nw134_
            d_789_v47_: _dafny.Map
            d_789_v47_ = d_743_v3_.f43
            rhs105_ = d_788_v46_
            rhs106_ = (0) - (d_741_v1_)
            rhs107_ = ((_dafny.MultiSet([(d_766_v24_)[default__.safeIndex(421, (d_766_v24_).length(0))]])) == (d_767_v25_)) in (d_764_v22_)
            rhs108_ = (d_766_v24_)[default__.safeIndex(421, (d_766_v24_).length(0))]
            rhs109_ = (0) - ((((0) - (d_741_v1_)) + (len((d_789_v47_)))) * (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_741_v1_: len(d_743_v3_.f43)}) for d_790_i2_ in range(default__.abs(-185))]))))
            lhs77_ = globalState
            lhs78_ = globalState
            d_788_v46_ = rhs105_
            lhs77_.f2 = rhs106_
            r3 = rhs107_
            r3 = rhs108_
            lhs78_.f2 = rhs109_
        if not (p0) or ((d_741_v1_) >= ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhxxksugn")))))):
            d_791_v48_: _dafny.Array
            nw135_ = _dafny.Array(int(0), 25)
            d_791_v48_ = nw135_
            index107_ = default__.safeIndex(428, (d_791_v48_).length(0))
            (d_791_v48_)[index107_] = d_741_v1_
            index108_ = default__.safeIndex(428, (d_791_v48_).length(0))
            (d_791_v48_)[index108_] = default__.safeDivisionInt(((0) - (d_741_v1_)) + (d_741_v1_), (0) - (default__.safeDivisionInt(d_741_v1_, d_741_v1_)))
            index109_ = default__.safeIndex(428, (d_791_v48_).length(0))
            (d_791_v48_)[index109_] = default__.safeDivisionInt(771, d_741_v1_)
            d_792_v49_: _dafny.Map
            d_792_v49_ = _dafny.Map({d_741_v1_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_793_i3_ in range(default__.abs(883))])})
            d_794_v50_: _dafny.Set
            d_794_v50_ = _dafny.Set({(d_791_v48_)[default__.safeIndex(428, (d_791_v48_).length(0))], -882})
            index110_ = default__.safeIndex(428, (d_791_v48_).length(0))
            rhs110_ = default__.fm2(globalState)
            rhs111_ = p0
            rhs112_ = p0
            rhs113_ = len((d_792_v49_) | (default__.fm41(d_741_v1_, (d_791_v48_)[default__.safeIndex(428, (d_791_v48_).length(0))], d_794_v50_, p0, globalState)))
            lhs79_ = globalState
            lhs80_ = globalState
            lhs81_ = globalState
            lhs82_ = d_791_v48_
            lhs83_ = default__.safeIndex(428, (d_791_v48_).length(0))
            lhs79_.f22 = rhs110_
            lhs80_.f1 = rhs111_
            lhs81_.f1 = rhs112_
            lhs82_[lhs83_] = rhs113_
            d_795_v51_: _dafny.MultiSet
            d_795_v51_ = _dafny.MultiSet([False])
            d_796_v52_: _dafny.Map
            d_796_v52_ = _dafny.Map({((d_795_v51_)[p0] if (p0) in (d_795_v51_) else 300): p0})
            d_797_v53_: _dafny.Set
            d_797_v53_ = _dafny.Set({p0})
            d_798_v54_: _dafny.Map
            d_798_v54_ = _dafny.Map({len(d_797_v53_): (d_791_v48_)[default__.safeIndex(428, (d_791_v48_).length(0))]})
            d_799_v55_: T0
            nw136_ = C3()
            nw136_.ctor__(d_791_v48_, ((d_796_v52_)[default__.fm2(globalState)] if (default__.fm2(globalState)) in (d_796_v52_) else p0), (d_798_v54_) | (d_798_v54_))
            d_799_v55_ = nw136_
            d_799_v55_ = d_799_v55_
            d_800_v56_: _dafny.Array
            def lambda38_(d_801_v1_):
                def lambda39_(d_802_i4_):
                    return (d_801_v1_) >= (d_801_v1_)

                return lambda39_

            init22_ = lambda38_(d_741_v1_)
            nw137_ = _dafny.Array(None, 24)
            for i0_22_ in range(nw137_.length(0)):
                nw137_[i0_22_] = init22_(i0_22_)
            d_800_v56_ = nw137_
            index111_ = default__.safeIndex(176, (d_800_v56_).length(0))
            (d_800_v56_)[index111_] = p0
            index112_ = default__.safeIndex(428, (d_791_v48_).length(0))
            index113_ = default__.safeIndex(176, (d_800_v56_).length(0))
            rhs114_ = ((d_791_v48_)[default__.safeIndex(428, (d_791_v48_).length(0))]) + (d_741_v1_)
            rhs115_ = p0
            rhs116_ = not(p0)
            rhs117_ = True
            lhs84_ = d_791_v48_
            lhs85_ = default__.safeIndex(428, (d_791_v48_).length(0))
            lhs86_ = d_800_v56_
            lhs87_ = default__.safeIndex(176, (d_800_v56_).length(0))
            lhs84_[lhs85_] = rhs114_
            lhs86_[lhs87_] = rhs115_
            r3 = rhs116_
            r3 = rhs117_
        elif True:
            if True:
                d_803_v57_: _dafny.MultiSet
                d_803_v57_ = _dafny.MultiSet([p0, p0])
                d_804_v58_: _dafny.Map
                d_804_v58_ = _dafny.Map({_dafny.CodePoint('q'): p0})
                d_805_v59_: _dafny.Map
                d_805_v59_ = _dafny.Map({d_803_v57_: len((default__.fm29(d_803_v57_, d_741_v1_, p0, len(d_804_v58_), globalState) if True else _dafny.SeqWithoutIsStrInference([d_741_v1_, default__.fm2(globalState), d_741_v1_, ((d_803_v57_)[p0] if (p0) in (d_803_v57_) else d_741_v1_), d_741_v1_])))})
                d_805_v59_ = (d_805_v59_).set(d_803_v57_, d_741_v1_)
                d_806_v60_: _dafny.Seq
                d_806_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpwe"))
                (globalState).f5 = (default__.safeDivisionInt(len(d_806_v60_), 389)) - (d_741_v1_)
                d_807_v61_: _dafny.Array
                nw138_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_807_v61_ = nw138_
                d_807_v61_ = d_807_v61_
                d_808_v62_: _dafny.Array
                nw139_ = _dafny.Array(False, 29)
                d_808_v62_ = nw139_
                index114_ = default__.safeIndex(996, (d_808_v62_).length(0))
                (d_808_v62_)[index114_] = p0
                index115_ = default__.safeIndex(996, (d_808_v62_).length(0))
                rhs118_ = (default__.fm0(d_741_v1_, globalState)) and (p0)
                rhs119_ = (p0) or (p0)
                rhs120_ = (0) - ((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_809_i5_ in range(default__.abs(-9))]))) + (default__.safeDivisionInt(d_741_v1_, d_741_v1_)))
                rhs121_ = d_808_v62_
                rhs122_ = d_741_v1_
                lhs88_ = d_808_v62_
                lhs89_ = default__.safeIndex(996, (d_808_v62_).length(0))
                lhs90_ = globalState
                lhs91_ = globalState
                lhs92_ = globalState
                lhs88_[lhs89_] = rhs118_
                lhs90_.f1 = rhs119_
                lhs91_.f0 = rhs120_
                d_808_v62_ = rhs121_
                lhs92_.f22 = rhs122_
                (globalState).f22 = len(d_806_v60_)
            elif True:
                (globalState).f0 = d_741_v1_
                d_810_v63_: _dafny.Seq
                d_810_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                d_811_v64_: _dafny.Set
                d_811_v64_ = _dafny.Set({d_741_v1_})
                d_812_v65_: _dafny.Map
                d_812_v65_ = _dafny.Map({d_810_v63_: d_811_v64_})
                (globalState).f22 = len((d_812_v65_).set(d_810_v63_, d_811_v64_))
                d_813_v66_: C2
                nw140_ = C2()
                nw140_.ctor__(not (p0) or (not(p0)), d_741_v1_)
                d_813_v66_ = nw140_
                d_814_v67_: _dafny.Seq
                d_814_v67_ = _dafny.SeqWithoutIsStrInference([(d_813_v66_).f41])
                d_815_v68_: _dafny.Map
                d_815_v68_ = _dafny.Map({d_814_v67_: (d_813_v66_).fm24(globalState)})
                d_816_v69_: _dafny.Set
                d_816_v69_ = _dafny.Set({default__.fm42(globalState), _dafny.Set({d_741_v1_})})
                d_817_v70_: _dafny.Seq
                d_817_v70_ = _dafny.SeqWithoutIsStrInference([d_816_v69_, d_816_v69_])
                d_818_v72_: _dafny.Set
                d_818_v72_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([p0])})
                def iife110_():
                    coll46_ = _dafny.Map()
                    compr_46_: _dafny.Seq
                    for compr_46_ in (d_818_v72_).Elements:
                        d_819_v71_: _dafny.Seq = compr_46_
                        if (d_819_v71_) in (d_818_v72_):
                            coll46_[d_819_v71_] = (d_813_v66_).f41
                    return _dafny.Map(coll46_)
                d_815_v68_ = (d_815_v68_ if (d_816_v69_).ispropersubset((d_817_v70_)[default__.safeIndex(len(d_810_v63_), len(d_817_v70_))]) else (d_815_v68_) | (iife110_()
                ))
                (globalState).f0 = (d_813_v66_).f42
            d_820_v73_: _dafny.Seq
            d_820_v73_ = _dafny.SeqWithoutIsStrInference([p0])
            d_821_v74_: D3
            d_821_v74_ = D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uv")), d_741_v1_, d_741_v1_, default__.fm2(globalState), d_820_v73_)
            d_822_v75_: _dafny.Seq
            d_822_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jye"))
            d_823_v76_: _dafny.Array
            def lambda40_(d_824_p0_):
                def lambda41_(d_825_i6_):
                    return d_824_p0_

                return lambda41_

            init23_ = lambda40_(p0)
            nw141_ = _dafny.Array(None, 20)
            for i0_23_ in range(nw141_.length(0)):
                nw141_[i0_23_] = init23_(i0_23_)
            d_823_v76_ = nw141_
            d_826_v77_: _dafny.Map
            d_826_v77_ = _dafny.Map({((d_821_v74_).cf8) + (d_822_v75_): d_823_v76_})
            d_826_v77_ = (d_826_v77_).set(d_822_v75_, d_823_v76_)
            d_827_v78_: _dafny.Array
            nw142_ = _dafny.Array(int(0), 10)
            d_827_v78_ = nw142_
            d_828_v79_: str
            d_828_v79_ = _dafny.CodePoint('b')
            d_829_v80_: D0
            d_829_v80_ = D0_DC1(d_827_v78_, d_828_v79_, 964)
            d_830_v81_: D0
            d_830_v81_ = D0_DC2(d_829_v80_)
            d_830_v81_ = D0_DC2(d_829_v80_)
            (globalState).f22 = (len(d_822_v75_)) + (d_741_v1_)
            d_831_v82_: D5
            d_831_v82_ = D5_DC13(444)
            d_832_v83_: _dafny.Array
            nw143_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_832_v83_ = nw143_
            d_833_v84_: _dafny.Array
            def lambda42_(d_834_v79_):
                def lambda43_(d_835_i7_):
                    return d_834_v79_

                return lambda43_

            init24_ = lambda42_(d_828_v79_)
            nw144_ = _dafny.Array(None, 14)
            for i0_24_ in range(nw144_.length(0)):
                nw144_[i0_24_] = init24_(i0_24_)
            d_833_v84_ = nw144_
            index116_ = default__.safeIndex(444, (d_832_v83_).length(0))
            (d_832_v83_)[index116_] = d_833_v84_
            index117_ = default__.safeIndex(444, (d_832_v83_).length(0))
            rhs123_ = d_741_v1_
            rhs124_ = 763
            rhs125_ = (d_831_v82_ if default__.fm0(d_741_v1_, globalState) else (d_831_v82_ if p0 else d_831_v82_))
            rhs126_ = d_833_v84_
            rhs127_ = (d_822_v75_) == (d_822_v75_)
            lhs93_ = globalState
            lhs94_ = globalState
            lhs95_ = d_832_v83_
            lhs96_ = default__.safeIndex(444, (d_832_v83_).length(0))
            lhs97_ = globalState
            lhs93_.f0 = rhs123_
            lhs94_.f2 = rhs124_
            d_831_v82_ = rhs125_
            lhs95_[lhs96_] = rhs126_
            lhs97_.f13 = rhs127_
        d_836_v85_: _dafny.Map
        d_836_v85_ = _dafny.Map({(_dafny.MultiSet([d_741_v1_, default__.fm2(globalState)])).cardinality: (0) - (d_741_v1_)})
        d_837_v86_: _dafny.Array
        nw145_ = _dafny.Array(int(0), 16)
        d_837_v86_ = nw145_
        d_838_v87_: _dafny.MultiSet
        d_838_v87_ = _dafny.MultiSet([d_837_v86_, d_837_v86_, d_837_v86_])
        d_839_v88_: _dafny.Set
        d_839_v88_ = _dafny.Set({p0})
        d_840_v89_: _dafny.Map
        d_840_v89_ = _dafny.Map({d_839_v88_: d_741_v1_})
        d_841_v90_: _dafny.Seq
        d_841_v90_ = _dafny.SeqWithoutIsStrInference([((d_840_v89_)[d_839_v88_] if (d_839_v88_) in (d_840_v89_) else d_741_v1_)])
        d_842_v91_: _dafny.Map
        d_842_v91_ = _dafny.Map({d_838_v87_: d_841_v90_})
        d_843_v92_: _dafny.Seq
        d_843_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xx"))
        d_844_v93_: _dafny.Set
        d_844_v93_ = _dafny.Set({d_741_v1_})
        d_845_v94_: D8
        d_845_v94_ = D8_DC18(d_844_v93_)
        d_846_v95_: str
        d_846_v95_ = _dafny.CodePoint('s')
        d_847_v96_: _dafny.MultiSet
        d_847_v96_ = _dafny.MultiSet([d_741_v1_])
        d_848_v97_: _dafny.Map
        d_848_v97_ = _dafny.Map({(d_742_v2_).set(False, d_741_v1_): p0})
        d_849_v98_: _dafny.Seq
        d_849_v98_ = _dafny.SeqWithoutIsStrInference([p0, default__.fm0(d_741_v1_, globalState)])
        d_850_v99_: D3
        d_850_v99_ = D3_DC7(d_843_v92_, (d_847_v96_).cardinality, len(d_848_v97_), 467, d_849_v98_)
        d_851_v100_: _dafny.Array
        nw146_ = _dafny.Array(None, 20)
        nw146_[int(0)] = -431
        nw146_[int(1)] = d_741_v1_
        nw146_[int(2)] = d_741_v1_
        nw146_[int(3)] = d_741_v1_
        nw146_[int(4)] = d_741_v1_
        nw146_[int(5)] = d_741_v1_
        nw146_[int(6)] = ((d_836_v85_)[d_741_v1_] if (d_741_v1_) in (d_836_v85_) else 2)
        nw146_[int(7)] = d_741_v1_
        nw146_[int(8)] = len(((d_842_v91_)[d_838_v87_] if (d_838_v87_) in (d_842_v91_) else d_841_v90_))
        nw146_[int(9)] = d_741_v1_
        nw146_[int(10)] = ((0) - (len(_dafny.Map({_dafny.CodePoint('w'): d_741_v1_})))) - (d_741_v1_)
        nw146_[int(11)] = d_741_v1_
        nw146_[int(12)] = (len((d_843_v92_).set(default__.safeIndex(len((d_845_v94_).cf27), len(d_843_v92_)), default__.fm35(d_741_v1_, globalState)))) - (d_741_v1_)
        nw146_[int(13)] = d_741_v1_
        nw146_[int(14)] = d_741_v1_
        nw146_[int(15)] = 346
        nw146_[int(16)] = d_741_v1_
        nw146_[int(17)] = len(d_843_v92_)
        nw146_[int(18)] = default__.safeModuloInt((0) - (len((d_843_v92_).set(default__.safeIndex(d_741_v1_, len(d_843_v92_)), d_846_v95_))), (d_850_v99_).cf11)
        nw146_[int(19)] = d_741_v1_
        d_851_v100_ = nw146_
        r0 = d_851_v100_
        r1 = d_847_v96_
        r2 = d_844_v93_
        r3 = not(p0)
        return r0, r1, r2, r3

    def m18(self, p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: int = int(0)
        d_852_v0_: bool
        d_852_v0_ = True
        if d_852_v0_:
            d_853_v1_: int
            d_853_v1_ = 23
            (globalState).f16 = (d_853_v1_) - (d_853_v1_)
            d_854_v2_: _dafny.MultiSet
            d_854_v2_ = _dafny.MultiSet([d_852_v0_, not(d_852_v0_), d_852_v0_])
            d_854_v2_ = ((d_854_v2_) - (d_854_v2_)) | (d_854_v2_)
            d_855_v4_: _dafny.MultiSet
            d_855_v4_ = _dafny.MultiSet([d_853_v1_])
            d_856_v5_: _dafny.Map
            d_856_v5_ = _dafny.Map({d_853_v1_: d_852_v0_})
            d_857_v6_: _dafny.Map
            d_857_v6_ = _dafny.Map({d_853_v1_: d_856_v5_})
            def iife111_():
                coll47_ = _dafny.Map()
                compr_47_: int
                for compr_47_ in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_855_v4_])).cardinality, d_853_v1_])).Elements:
                    d_858_v3_: int = compr_47_
                    if (d_858_v3_) in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_855_v4_])).cardinality, d_853_v1_])):
                        coll47_[(d_858_v3_) * (len(_dafny.Set({d_852_v0_, d_852_v0_, True})))] = True
                return _dafny.Map(coll47_)
            (globalState).f2 = (len((iife111_()
            ) | (((d_857_v6_)[d_853_v1_] if (d_853_v1_) in (d_857_v6_) else d_856_v5_))) if d_852_v0_ else d_853_v1_)
            d_859_v7_: _dafny.Seq
            d_859_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dglsmhy"))
            d_860_v9_: D8
            def iife112_():
                coll48_ = _dafny.Set()
                compr_48_: int
                for compr_48_ in _dafny.IntegerRange(965, -300):
                    d_861_v8_: int = compr_48_
                    if ((965) <= (d_861_v8_)) and ((d_861_v8_) < (-300)):
                        coll48_ = coll48_.union(_dafny.Set([default__.safeDivisionInt(d_861_v8_, d_853_v1_)]))
                return _dafny.Set(coll48_)
            d_860_v9_ = D8_DC18(iife112_()
)
            d_862_v10_: _dafny.Map
            d_862_v10_ = _dafny.Map({d_860_v9_: d_855_v4_})
            d_863_v11_: _dafny.Array
            nw147_ = _dafny.Array(None, 15)
            nw147_[int(0)] = d_853_v1_
            nw147_[int(1)] = (0) - (len(d_859_v7_))
            nw147_[int(2)] = len(d_859_v7_)
            nw147_[int(3)] = d_853_v1_
            nw147_[int(4)] = d_853_v1_
            nw147_[int(5)] = -179
            nw147_[int(6)] = d_853_v1_
            nw147_[int(7)] = d_853_v1_
            nw147_[int(8)] = (d_854_v2_).cardinality
            nw147_[int(9)] = d_853_v1_
            nw147_[int(10)] = d_853_v1_
            nw147_[int(11)] = len(_dafny.Set({d_852_v0_, d_852_v0_, d_852_v0_, d_852_v0_}))
            nw147_[int(12)] = default__.fm2(globalState)
            nw147_[int(13)] = len(d_862_v10_)
            nw147_[int(14)] = d_853_v1_
            d_863_v11_ = nw147_
            d_864_v12_: _dafny.Map
            d_864_v12_ = _dafny.Map({d_852_v0_: d_852_v0_})
            d_865_v13_: _dafny.Seq
            d_865_v13_ = _dafny.SeqWithoutIsStrInference([d_864_v12_, d_864_v12_])
            d_866_v14_: _dafny.Map
            d_866_v14_ = _dafny.Map({d_853_v1_: d_853_v1_})
            d_867_v15_: _dafny.Map
            d_867_v15_ = _dafny.Map({len(d_866_v14_): d_853_v1_})
            d_868_v16_: _dafny.Seq
            d_868_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_853_v1_: d_853_v1_}), d_866_v14_])
            d_869_v17_: C3
            nw148_ = C3()
            nw148_.ctor__(d_863_v11_, default__.fm0(len(d_865_v13_), globalState), (d_866_v14_ if d_852_v0_ else (d_868_v16_)[default__.safeIndex(d_853_v1_, len(d_868_v16_))]))
            d_869_v17_ = nw148_
            d_870_v18_: _dafny.Seq
            d_870_v18_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_853_v1_, globalState)])
            d_871_v19_: _dafny.Map
            d_871_v19_ = _dafny.Map({d_870_v18_: d_869_v17_.f40})
            d_872_v20_: D3
            d_872_v20_ = D3_DC6(d_871_v19_)
            source12_ = d_872_v20_
            if source12_.is_DC7:
                d_873___mcc_h0_ = source12_.cf8
                d_874___mcc_h1_ = source12_.cf9
                d_875___mcc_h2_ = source12_.cf10
                d_876___mcc_h3_ = source12_.cf11
                d_877___mcc_h4_ = source12_.cf12
                d_878_cf12_ = d_877___mcc_h4_
                d_879_cf11_ = d_876___mcc_h3_
                d_880_cf10_ = d_875___mcc_h2_
                d_881_cf9_ = d_874___mcc_h1_
                d_882_cf8_ = d_873___mcc_h0_
                d_883_v21_: _dafny.Seq
                d_883_v21_ = _dafny.SeqWithoutIsStrInference([d_853_v1_, d_880_cf10_])
                d_884_v22_: _dafny.Array
                nw149_ = _dafny.Array(None, 17)
                nw149_[int(0)] = ((d_883_v21_)[default__.safeIndex((d_869_v17_).fm3(globalState), len(d_883_v21_))]) < (d_881_cf9_)
                nw149_[int(1)] = d_852_v0_
                nw149_[int(2)] = default__.fm0(d_853_v1_, globalState)
                nw149_[int(3)] = d_869_v17_.f40
                nw149_[int(4)] = (d_879_cf11_) != (d_879_cf11_)
                nw149_[int(5)] = ((d_856_v5_)[d_853_v1_] if (d_853_v1_) in (d_856_v5_) else d_869_v17_.f40)
                nw149_[int(6)] = d_869_v17_.f40
                nw149_[int(7)] = d_869_v17_.f40
                nw149_[int(8)] = (default__.fm43(len(_dafny.SeqWithoutIsStrInference([default__.fm2(globalState), d_853_v1_])), d_853_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbhscjr")), globalState)).issubset(_dafny.MultiSet([-599, d_853_v1_]))
                nw149_[int(9)] = d_869_v17_.f40
                nw149_[int(10)] = d_852_v0_
                nw149_[int(11)] = (d_869_v17_.f40) or (((d_864_v12_)[d_869_v17_.f40] if (d_869_v17_.f40) in (d_864_v12_) else d_869_v17_.f40))
                nw149_[int(12)] = d_869_v17_.f40
                nw149_[int(13)] = (d_880_cf10_) > (d_881_cf9_)
                nw149_[int(14)] = d_852_v0_
                nw149_[int(15)] = d_869_v17_.f40
                nw149_[int(16)] = d_852_v0_
                d_884_v22_ = nw149_
                d_885_v23_: _dafny.Map
                d_885_v23_ = _dafny.Map({d_883_v21_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "el")) if d_869_v17_.f40 else d_882_cf8_)})
                d_886_v25_: _dafny.Set
                d_886_v25_ = _dafny.Set({d_883_v21_})
                def iife113_():
                    coll49_ = _dafny.Map()
                    compr_49_: _dafny.Seq
                    for compr_49_ in (d_886_v25_).Elements:
                        d_887_v24_: _dafny.Seq = compr_49_
                        if (d_887_v24_) in (d_886_v25_):
                            coll49_[d_887_v24_] = d_859_v7_
                    return _dafny.Map(coll49_)
                rhs128_ = d_853_v1_
                rhs129_ = d_884_v22_
                rhs130_ = iife113_()

                lhs98_ = globalState
                lhs98_.f22 = rhs128_
                d_884_v22_ = rhs129_
                d_885_v23_ = rhs130_
                d_888_v26_: _dafny.Array
                def lambda44_(d_889_i0_):
                    return _dafny.MultiSet([_dafny.CodePoint('b'), _dafny.CodePoint('i'), _dafny.CodePoint('v')])

                init25_ = lambda44_
                nw150_ = _dafny.Array(None, 1)
                for i0_25_ in range(nw150_.length(0)):
                    nw150_[i0_25_] = init25_(i0_25_)
                d_888_v26_ = nw150_
                d_890_v27_: D2
                d_890_v27_ = D2_DC4(d_888_v26_)
                d_890_v27_ = d_890_v27_
                d_891_v28_: _dafny.Map
                d_891_v28_ = _dafny.Map({_dafny.Set({d_879_cf11_, len(d_882_cf8_)}): d_853_v1_})
                d_892_v29_: C3
                nw151_ = C3()
                nw151_.ctor__(d_863_v11_, (_dafny.MultiSet(d_878_cf12_)).isdisjoint((d_854_v2_).set(d_869_v17_.f40, default__.abs(((d_891_v28_)[p0] if (p0) in (d_891_v28_) else d_881_cf9_)))), _dafny.Map({d_881_cf9_: d_881_cf9_}))
                d_892_v29_ = nw151_
                d_880_cf10_ = (d_869_v17_).fm3(globalState)
            elif True:
                d_893___mcc_h5_ = source12_.cf7
                d_894_cf7_ = d_893___mcc_h5_
                d_895_v30_: _dafny.Map
                d_895_v30_ = _dafny.Map({d_853_v1_: d_860_v9_})
                d_896_v32_: _dafny.Seq
                def iife114_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(-102, -901):
                        d_897_v31_: int = compr_50_
                        if ((-102) <= (d_897_v31_)) and ((d_897_v31_) < (-901)):
                            coll50_[(d_897_v31_) + (d_853_v1_)] = d_860_v9_
                    return _dafny.Map(coll50_)
                d_896_v32_ = _dafny.SeqWithoutIsStrInference([iife114_()
                ])
                (globalState).f16 = ((d_853_v1_) * (d_853_v1_)) * (len((d_895_v30_) | ((d_896_v32_)[default__.safeIndex(d_853_v1_, len(d_896_v32_))])))
                d_898_v33_: str
                d_898_v33_ = _dafny.CodePoint('w')
                d_898_v33_ = d_898_v33_
                (globalState).f0 = d_853_v1_
                index118_ = default__.safeIndex(462, ((d_869_v17_).f39).length(0))
                ((d_869_v17_).f39)[index118_] = len(d_859_v7_)
                d_899_v34_: _dafny.Seq
                d_899_v34_ = _dafny.SeqWithoutIsStrInference([d_859_v7_, d_859_v7_])
                index119_ = default__.safeIndex(462, ((d_869_v17_).f39).length(0))
                ((d_869_v17_).f39)[index119_] = (0) - (len(((d_870_v18_) + (default__.fm28(d_869_v17_.f40, ((d_856_v5_)[len(d_899_v34_)] if (len(d_899_v34_)) in (d_856_v5_) else d_852_v0_), d_853_v1_, globalState))) + (_dafny.SeqWithoutIsStrInference([d_869_v17_.f40]))))
        elif True:
            d_900_v35_: int
            d_900_v35_ = 725
            (globalState).f5 = d_900_v35_
            d_901_v36_: bool
            d_902_v37_: bool
            d_903_v38_: _dafny.Map
            d_904_v39_: int
            out30_: bool
            out31_: bool
            out32_: _dafny.Map
            out33_: int
            out30_, out31_, out32_, out33_ = default__.m0(globalState)
            d_901_v36_ = out30_
            d_902_v37_ = out31_
            d_903_v38_ = out32_
            d_904_v39_ = out33_
            if d_852_v0_:
                d_905_v40_: _dafny.Seq
                d_905_v40_ = _dafny.SeqWithoutIsStrInference([d_904_v39_, d_900_v35_, d_904_v39_, 615])
                d_906_v41_: _dafny.Map
                d_906_v41_ = _dafny.Map({d_901_v36_: len(d_905_v40_)})
                d_907_v42_: C1
                nw152_ = C1()
                nw152_.ctor__(d_906_v41_)
                d_907_v42_ = nw152_
                (globalState).f0 = default__.fm2(globalState)
                d_908_v43_: C1
                nw153_ = C1()
                nw153_.ctor__(d_906_v41_)
                d_908_v43_ = nw153_
                d_909_v44_: _dafny.Set
                d_909_v44_ = _dafny.Set({(-786) >= (d_900_v35_)})
                d_909_v44_ = d_909_v44_
                d_910_v45_: _dafny.Map
                d_910_v45_ = _dafny.Map({d_902_v37_: d_901_v36_})
                d_911_v46_: _dafny.MultiSet
                d_911_v46_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))), d_904_v39_, d_900_v35_, len(_dafny.SeqWithoutIsStrInference([d_904_v39_ for d_912_i1_ in range(default__.abs(320))]))])
                rhs131_ = ((d_910_v45_)[d_852_v0_] if (d_852_v0_) in (d_910_v45_) else d_901_v36_)
                rhs132_ = ((d_900_v35_) * (d_904_v39_)) * (((d_911_v46_)[(d_905_v40_)[default__.safeIndex(d_904_v39_, len(d_905_v40_))]] if ((d_905_v40_)[default__.safeIndex(d_904_v39_, len(d_905_v40_))]) in (d_911_v46_) else d_900_v35_))
                lhs99_ = globalState
                d_901_v36_ = rhs131_
                lhs99_.f0 = rhs132_
            elif True:
                d_913_v47_: _dafny.Seq
                d_913_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfjknnwd"))
                d_914_v48_: _dafny.Seq
                d_914_v48_ = _dafny.SeqWithoutIsStrInference([d_902_v37_, d_852_v0_, (d_913_v47_) != (d_913_v47_)])
                d_915_v49_: _dafny.Seq
                d_915_v49_ = _dafny.SeqWithoutIsStrInference([d_904_v39_, d_900_v35_, d_900_v35_, d_900_v35_, d_904_v39_])
                d_914_v48_ = ((d_914_v48_).set(default__.safeIndex((D5_DC13(len(d_915_v49_))).cf19, len(d_914_v48_)), d_901_v36_)) + ((d_914_v48_) + (d_914_v48_))
                (globalState).f16 = 991
                rhs133_ = ((d_914_v48_) + (d_914_v48_)) < (d_914_v48_)
                rhs134_ = not((not (d_902_v37_) or (not(d_852_v0_))) == (d_852_v0_))
                rhs135_ = d_904_v39_
                lhs100_ = globalState
                lhs101_ = globalState
                lhs102_ = globalState
                lhs100_.f1 = rhs133_
                lhs101_.f13 = rhs134_
                lhs102_.f5 = rhs135_
                (globalState).f16 = len(d_913_v47_)
                d_916_v50_: str
                d_916_v50_ = _dafny.CodePoint('i')
                d_917_v51_: C2
                nw154_ = C2()
                nw154_.ctor__((d_916_v50_) in (_dafny.SeqWithoutIsStrInference([d_916_v50_ for d_918_i2_ in range(default__.abs(234))])), d_900_v35_)
                d_917_v51_ = nw154_
            d_919_v52_: _dafny.Seq
            d_919_v52_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_904_v39_, globalState)])
            d_919_v52_ = (d_919_v52_) + ((_dafny.SeqWithoutIsStrInference([not(d_852_v0_)]) if d_901_v36_ else d_919_v52_))
            (globalState).f22 = default__.safeModuloInt(d_904_v39_, d_900_v35_)
        d_920_v53_: int
        d_920_v53_ = 499
        hi6_ = d_920_v53_
        for d_921_i3_ in range(d_920_v53_, hi6_):
            d_922_v54_: _dafny.Array
            nw155_ = _dafny.Array(D9.default()(), 12)
            d_922_v54_ = nw155_
            d_923_v55_: _dafny.Array
            nw156_ = _dafny.Array(_dafny.CodePoint('D'), 13)
            d_923_v55_ = nw156_
            d_924_v56_: _dafny.Array
            d_924_v56_ = d_922_v54_
            rhs136_ = (d_924_v56_)
            rhs137_ = d_923_v55_
            rhs138_ = d_920_v53_
            lhs103_ = globalState
            d_922_v54_ = rhs136_
            d_923_v55_ = rhs137_
            lhs103_.f21 = rhs138_
            d_925_v57_: _dafny.Array
            nw157_ = _dafny.Array(_dafny.Seq({}), 14)
            d_925_v57_ = nw157_
            d_926_v58_: _dafny.Seq
            d_926_v58_ = _dafny.SeqWithoutIsStrInference([d_920_v53_])
            index120_ = default__.safeIndex(405, (d_925_v57_).length(0))
            (d_925_v57_)[index120_] = d_926_v58_
            index121_ = default__.safeIndex(405, (d_925_v57_).length(0))
            (d_925_v57_)[index121_] = d_926_v58_
            d_927_v59_: _dafny.Array
            def lambda45_(d_928_i4_):
                return (d_928_i4_) - (958)

            init26_ = lambda45_
            nw158_ = _dafny.Array(None, 12)
            for i0_26_ in range(nw158_.length(0)):
                nw158_[i0_26_] = init26_(i0_26_)
            d_927_v59_ = nw158_
            index122_ = default__.safeIndex(901, (d_927_v59_).length(0))
            (d_927_v59_)[index122_] = (d_926_v58_)[default__.safeIndex(d_921_i3_, len(d_926_v58_))]
            index123_ = default__.safeIndex(901, (d_927_v59_).length(0))
            (d_927_v59_)[index123_] = d_921_i3_
            if d_852_v0_:
                d_929_v60_: _dafny.Array
                d_930_v61_: _dafny.MultiSet
                d_931_v62_: _dafny.Set
                d_932_v63_: bool
                out34_: _dafny.Array
                out35_: _dafny.MultiSet
                out36_: _dafny.Set
                out37_: bool
                out34_, out35_, out36_, out37_ = (self).m17(d_852_v0_, globalState)
                d_929_v60_ = out34_
                d_930_v61_ = out35_
                d_931_v62_ = out36_
                d_932_v63_ = out37_
                d_933_v64_: _dafny.Map
                d_933_v64_ = _dafny.Map({False: d_921_i3_})
                d_934_v65_: _dafny.Map
                d_934_v65_ = _dafny.Map({d_852_v0_: (_dafny.Map({d_932_v63_: d_920_v53_})) | (d_933_v64_)})
                d_935_v66_: str
                d_935_v66_ = _dafny.CodePoint('k')
                d_936_v67_: D0
                d_936_v67_ = D0_DC1(d_927_v59_, d_935_v66_, d_920_v53_)
                d_937_v68_: _dafny.Seq
                d_937_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnpbluy"))
                pat_let_tv24_ = d_852_v0_
                pat_let_tv25_ = globalState
                def iife115_(_pat_let32_0):
                    def iife116_(d_938_dt__update__tmp_h0_):
                        def iife117_(_pat_let33_0):
                            def iife118_(d_939_dt__update_hcf2_h0_):
                                return D0_DC1((d_938_dt__update__tmp_h0_).cf1, d_939_dt__update_hcf2_h0_, (d_938_dt__update__tmp_h0_).cf3)
                            return iife118_(_pat_let33_0)
                        return iife117_((self).fm22(not(pat_let_tv24_), pat_let_tv25_))
                    return iife116_(_pat_let32_0)
                rhs139_ = d_852_v0_
                rhs140_ = (iife115_(d_936_v67_)).cf1
                rhs141_ = (_dafny.Map({d_852_v0_: d_933_v64_})) | (((d_934_v65_).set(not(False), default__.fm33(_dafny.SeqWithoutIsStrInference([d_921_i3_ for d_940_i5_ in range(default__.abs(-334))]), d_932_v63_, globalState))) | (default__.fm44(len(d_937_v68_), d_932_v63_, globalState)))
                d_932_v63_ = rhs139_
                d_927_v59_ = rhs140_
                d_934_v65_ = rhs141_
                d_941_v69_: _dafny.Map
                d_941_v69_ = _dafny.Map({(d_927_v59_)[default__.safeIndex(901, (d_927_v59_).length(0))]: d_920_v53_})
                d_941_v69_ = (d_941_v69_).set(d_920_v53_, d_921_i3_)
                (globalState).f1 = d_932_v63_
                index124_ = default__.safeIndex(901, (d_927_v59_).length(0))
                (d_927_v59_)[index124_] = ((d_925_v57_)[default__.safeIndex(405, (d_925_v57_).length(0))])[default__.safeIndex(d_920_v53_, len((d_925_v57_)[default__.safeIndex(405, (d_925_v57_).length(0))]))]
            elif True:
                (globalState).f16 = (d_921_i3_) + (len(d_926_v58_))
                d_942_v70_: _dafny.Map
                d_942_v70_ = _dafny.Map({not(True): d_852_v0_})
                d_943_v71_: _dafny.MultiSet
                d_943_v71_ = _dafny.MultiSet([len(d_942_v70_), (len((d_925_v57_)[default__.safeIndex(405, (d_925_v57_).length(0))])) - (d_921_i3_), ((d_927_v59_)[default__.safeIndex(901, (d_927_v59_).length(0))]) * (d_921_i3_)])
                d_943_v71_ = (default__.fm43((0) - (d_921_i3_), d_920_v53_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_944_i6_ in range(default__.abs(-881))]), globalState)).intersection(d_943_v71_)
                d_945_v72_: _dafny.Array
                def lambda46_(d_946_i7_):
                    return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_947_i8_ in range(default__.abs(477))])

                init27_ = lambda46_
                nw159_ = _dafny.Array(None, 19)
                for i0_27_ in range(nw159_.length(0)):
                    nw159_[i0_27_] = init27_(i0_27_)
                d_945_v72_ = nw159_
                index125_ = default__.safeIndex(408, (d_945_v72_).length(0))
                (d_945_v72_)[index125_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvkxjmg"))
                d_948_v73_: _dafny.Seq
                d_948_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrvvto"))
                d_949_v74_: str
                d_949_v74_ = _dafny.CodePoint('f')
                index126_ = default__.safeIndex(408, (d_945_v72_).length(0))
                rhs142_ = ((d_948_v73_).set(default__.safeIndex((0) - ((d_927_v59_)[default__.safeIndex(901, (d_927_v59_).length(0))]), len(d_948_v73_)), d_949_v74_)) + (d_948_v73_)
                rhs143_ = d_920_v53_
                rhs144_ = d_945_v72_
                rhs145_ = d_852_v0_
                lhs104_ = d_945_v72_
                lhs105_ = default__.safeIndex(408, (d_945_v72_).length(0))
                lhs104_[lhs105_] = rhs142_
                d_920_v53_ = rhs143_
                d_945_v72_ = rhs144_
                d_852_v0_ = rhs145_
                d_950_v75_: C0
                nw160_ = C0()
                nw160_.ctor__(d_923_v55_)
                d_950_v75_ = nw160_
                index127_ = default__.safeIndex(405, (d_925_v57_).length(0))
                (d_925_v57_)[index127_] = d_926_v58_
        r2 = (0) - (d_920_v53_)
        d_951_v76_: _dafny.Array
        nw161_ = _dafny.Array(_dafny.Seq({}), 12)
        d_951_v76_ = nw161_
        d_952_v77_: _dafny.Map
        d_952_v77_ = _dafny.Map({d_852_v0_: default__.fm2(globalState)})
        d_953_v78_: _dafny.Seq
        d_953_v78_ = _dafny.SeqWithoutIsStrInference([len(d_952_v77_), d_920_v53_, d_920_v53_, default__.fm2(globalState), d_920_v53_])
        index128_ = default__.safeIndex(48, (d_951_v76_).length(0))
        (d_951_v76_)[index128_] = d_953_v78_
        d_954_v79_: D12
        d_954_v79_ = D12_DC29()
        pat_let_tv26_ = d_953_v78_
        pat_let_tv27_ = d_953_v78_
        pat_let_tv28_ = d_953_v78_
        pat_let_tv29_ = d_953_v78_
        index129_ = default__.safeIndex(48, (d_951_v76_).length(0))
        def lambda47_(source13_):
            if source13_.is_DC29:
                return (pat_let_tv26_) + (pat_let_tv27_)
            elif source13_.is_DC28:
                d_955___mcc_h6_ = source13_.cf43
                d_956_cf43_ = d_955___mcc_h6_
                return pat_let_tv28_
            elif True:
                d_957___mcc_h7_ = source13_.cf44
                d_958_cf44_ = d_957___mcc_h7_
                return pat_let_tv29_

        (d_951_v76_)[index129_] = lambda47_(d_954_v79_)
        d_959_v80_: D8
        d_959_v80_ = D8_DC18(p0)
        d_960_v81_: _dafny.MultiSet
        d_960_v81_ = _dafny.MultiSet([d_959_v80_])
        pat_let_tv30_ = p0
        def iife119_(_pat_let34_0):
            def iife120_(d_961_dt__update__tmp_h1_):
                def iife121_(_pat_let35_0):
                    def iife122_(d_962_dt__update_hcf27_h0_):
                        return D8_DC18(d_962_dt__update_hcf27_h0_)
                    return iife122_(_pat_let35_0)
                return iife121_(pat_let_tv30_)
            return iife120_(_pat_let34_0)
        if (d_960_v81_).ispropersubset((d_960_v81_).intersection(_dafny.MultiSet([d_959_v80_, iife119_(d_959_v80_), D8_DC18(p0), d_959_v80_, d_959_v80_]))):
            d_963_v82_: _dafny.Array
            def lambda48_(d_964_v53_):
                def lambda49_(d_965_i9_):
                    return default__.safeModuloInt(d_965_i9_, d_964_v53_)

                return lambda49_

            init28_ = lambda48_(d_920_v53_)
            nw162_ = _dafny.Array(None, 4)
            for i0_28_ in range(nw162_.length(0)):
                nw162_[i0_28_] = init28_(i0_28_)
            d_963_v82_ = nw162_
            d_966_v83_: _dafny.Map
            d_966_v83_ = _dafny.Map({(0) - (-865): d_920_v53_})
            d_967_v84_: C3
            nw163_ = C3()
            nw163_.ctor__(d_963_v82_, d_852_v0_, (d_966_v83_) | (d_966_v83_))
            d_967_v84_ = nw163_
            d_968_v85_: C2
            nw164_ = C2()
            nw164_.ctor__(d_967_v84_.f40, (len((d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))])) - (d_920_v53_))
            d_968_v85_ = nw164_
            d_969_v86_: str
            d_969_v86_ = _dafny.CodePoint('v')
            d_969_v86_ = d_969_v86_
            d_970_v87_: _dafny.MultiSet
            d_970_v87_ = _dafny.MultiSet([d_967_v84_.f40])
            d_971_v88_: D12
            d_971_v88_ = D12_DC28(d_970_v87_)
            d_972_v89_: D12
            d_972_v89_ = D12_DC30(d_971_v88_)
            source14_ = d_972_v89_
            if source14_.is_DC29:
                d_973_v90_: D4
                d_973_v90_ = D4_DC8(_dafny.CodePoint('p'))
                d_974_v91_: _dafny.Array
                nw165_ = _dafny.Array(None, 26)
                nw165_[int(0)] = d_969_v86_
                nw165_[int(1)] = _dafny.CodePoint('y')
                nw165_[int(2)] = d_969_v86_
                nw165_[int(3)] = d_969_v86_
                nw165_[int(4)] = d_969_v86_
                nw165_[int(5)] = d_969_v86_
                nw165_[int(6)] = d_969_v86_
                nw165_[int(7)] = d_969_v86_
                nw165_[int(8)] = d_969_v86_
                nw165_[int(9)] = d_969_v86_
                nw165_[int(10)] = d_969_v86_
                nw165_[int(11)] = d_969_v86_
                nw165_[int(12)] = d_969_v86_
                nw165_[int(13)] = d_969_v86_
                nw165_[int(14)] = d_969_v86_
                nw165_[int(15)] = d_969_v86_
                nw165_[int(16)] = d_969_v86_
                nw165_[int(17)] = d_969_v86_
                nw165_[int(18)] = d_969_v86_
                nw165_[int(19)] = d_969_v86_
                nw165_[int(20)] = d_969_v86_
                nw165_[int(21)] = _dafny.CodePoint('j')
                nw165_[int(22)] = d_969_v86_
                nw165_[int(23)] = d_969_v86_
                nw165_[int(24)] = d_969_v86_
                nw165_[int(25)] = (d_973_v90_).cf13
                d_974_v91_ = nw165_
                d_975_v92_: C0
                nw166_ = C0()
                nw166_.ctor__(d_974_v91_)
                d_975_v92_ = nw166_
                d_976_v93_: _dafny.Array
                def lambda50_(d_977_v85_):
                    def lambda51_(d_978_i10_):
                        return (_dafny.MultiSet([(d_977_v85_).f42])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - ((d_977_v85_).f42) for d_979_i11_ in range(default__.abs(346))])))

                    return lambda51_

                init29_ = lambda50_(d_968_v85_)
                nw167_ = _dafny.Array(None, 17)
                for i0_29_ in range(nw167_.length(0)):
                    nw167_[i0_29_] = init29_(i0_29_)
                d_976_v93_ = nw167_
                d_980_v94_: _dafny.Seq
                d_980_v94_ = _dafny.SeqWithoutIsStrInference([(d_968_v85_).f41])
                d_981_v95_: _dafny.MultiSet
                d_981_v95_ = _dafny.MultiSet([d_980_v94_, d_980_v94_])
                index130_ = default__.safeIndex(19, (d_976_v93_).length(0))
                (d_976_v93_)[index130_] = _dafny.MultiSet([(d_968_v85_).f42, (d_981_v95_).cardinality, (d_968_v85_).f42])
                d_982_v96_: _dafny.Array
                nw168_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_982_v96_ = nw168_
                d_983_v97_: _dafny.Array
                nw169_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
                d_983_v97_ = nw169_
                index131_ = default__.safeIndex(272, (d_982_v96_).length(0))
                (d_982_v96_)[index131_] = d_983_v97_
                d_984_v98_: _dafny.Set
                d_984_v98_ = _dafny.Set({d_967_v84_.f40, d_852_v0_})
                d_985_v99_: _dafny.MultiSet
                d_985_v99_ = _dafny.MultiSet([-601])
                d_986_v100_: _dafny.Map
                d_986_v100_ = _dafny.Map({d_852_v0_: d_984_v98_})
                index132_ = default__.safeIndex(19, (d_976_v93_).length(0))
                index133_ = default__.safeIndex(272, (d_982_v96_).length(0))
                rhs146_ = d_985_v99_
                rhs147_ = d_983_v97_
                rhs148_ = default__.fm34((default__.fm28(d_967_v84_.f40, False, (d_968_v85_).f42, globalState)).set(default__.safeIndex(len(((d_986_v100_)[d_967_v84_.f40] if (d_967_v84_.f40) in (d_986_v100_) else _dafny.Set({(d_968_v85_).f41}))), len(default__.fm28(d_967_v84_.f40, False, (d_968_v85_).f42, globalState))), (d_968_v85_).fm24(globalState)), ((d_968_v85_).f42 if d_852_v0_ else d_920_v53_), globalState)
                rhs149_ = (d_967_v84_).fm3(globalState)
                rhs150_ = (0) - (d_920_v53_)
                lhs106_ = d_976_v93_
                lhs107_ = default__.safeIndex(19, (d_976_v93_).length(0))
                lhs108_ = d_982_v96_
                lhs109_ = default__.safeIndex(272, (d_982_v96_).length(0))
                lhs110_ = globalState
                lhs111_ = globalState
                lhs106_[lhs107_] = rhs146_
                lhs108_[lhs109_] = rhs147_
                d_984_v98_ = rhs148_
                lhs110_.f5 = rhs149_
                lhs111_.f5 = rhs150_
                (globalState).f1 = default__.fm0(default__.safeDivisionInt(d_920_v53_, d_920_v53_), globalState)
                d_987_v101_: T1
                nw170_ = C1()
                nw170_.ctor__(_dafny.Map({d_967_v84_.f40: d_920_v53_}))
                d_987_v101_ = nw170_
                d_988_v102_: _dafny.Array
                nw171_ = _dafny.Array(None, 9)
                nw171_[int(0)] = d_987_v101_
                nw171_[int(1)] = d_987_v101_
                nw171_[int(2)] = d_987_v101_
                nw171_[int(3)] = d_987_v101_
                nw171_[int(4)] = d_987_v101_
                nw171_[int(5)] = d_987_v101_
                nw171_[int(6)] = d_987_v101_
                nw171_[int(7)] = d_987_v101_
                nw171_[int(8)] = d_987_v101_
                d_988_v102_ = nw171_
                d_989_v103_: _dafny.Map
                d_989_v103_ = _dafny.Map({(d_968_v85_).f42: d_987_v101_})
                index134_ = default__.safeIndex(791, (d_988_v102_).length(0))
                (d_988_v102_)[index134_] = ((d_989_v103_)[(d_968_v85_).f42] if ((d_968_v85_).f42) in (d_989_v103_) else d_987_v101_)
                index135_ = default__.safeIndex(791, (d_988_v102_).length(0))
                (d_988_v102_)[index135_] = (d_987_v101_ if d_852_v0_ else d_987_v101_)
            elif source14_.is_DC28:
                d_990___mcc_h8_ = source14_.cf43
                d_991_cf43_ = d_990___mcc_h8_
                d_992_v104_: C2
                nw172_ = C2()
                nw172_.ctor__((d_968_v85_).f41, d_920_v53_)
                d_992_v104_ = nw172_
                index136_ = default__.safeIndex(15, ((d_967_v84_).f39).length(0))
                ((d_967_v84_).f39)[index136_] = (0) - ((d_968_v85_).f42)
                d_993_v105_: _dafny.Set
                d_993_v105_ = _dafny.Set({False, False, (d_992_v104_).f41, d_967_v84_.f40})
                d_994_v106_: _dafny.Map
                d_994_v106_ = _dafny.Map({d_852_v0_: False})
                d_995_v107_: D4
                d_995_v107_ = D4_DC8(d_969_v86_)
                index137_ = default__.safeIndex(15, ((d_967_v84_).f39).length(0))
                rhs151_ = default__.safeModuloInt(((d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))])[default__.safeIndex(len(d_993_v105_), len((d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))]))], len((d_994_v106_).set((d_968_v85_).f41, not(False))))
                rhs152_ = (d_995_v107_).cf13
                lhs112_ = (d_967_v84_).f39
                lhs113_ = default__.safeIndex(15, ((d_967_v84_).f39).length(0))
                lhs112_[lhs113_] = rhs151_
                d_969_v86_ = rhs152_
                (globalState).f2 = (d_920_v53_) * (d_920_v53_)
                (globalState).f22 = default__.safeDivisionInt(d_920_v53_, (_dafny.MultiSet(d_953_v78_)).cardinality)
            elif True:
                d_996___mcc_h9_ = source14_.cf44
                d_997_cf44_ = d_996___mcc_h9_
                d_998_v108_: _dafny.Array
                nw173_ = _dafny.Array(False, 6)
                d_998_v108_ = nw173_
                index138_ = default__.safeIndex(141, (d_998_v108_).length(0))
                (d_998_v108_)[index138_] = (d_968_v85_).f41
                index139_ = default__.safeIndex(141, (d_998_v108_).length(0))
                (d_998_v108_)[index139_] = (d_968_v85_).f41
                index140_ = default__.safeIndex(141, (d_998_v108_).length(0))
                (d_998_v108_)[index140_] = d_967_v84_.f40
                (globalState).f5 = (d_968_v85_).f42
                d_999_v109_: _dafny.Set
                d_999_v109_ = _dafny.Set({d_952_v77_})
                d_1000_v110_: _dafny.Set
                d_1000_v110_ = _dafny.Set({(d_968_v85_).f42, (d_968_v85_).f42, d_920_v53_, len(d_999_v109_)})
                d_1001_v111_: _dafny.Seq
                d_1001_v111_ = (d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))]
                def iife123_():
                    coll51_ = _dafny.Set()
                    compr_51_: int
                    for compr_51_ in ((d_1001_v111_)).Elements:
                        d_1002_v112_: int = compr_51_
                        if (d_1002_v112_) in ((d_1001_v111_)):
                            coll51_ = coll51_.union(_dafny.Set([(d_1002_v112_) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iiscweu")))])))]))
                    return _dafny.Set(coll51_)
                rhs153_ = (iife123_()
                ) - (default__.fm42(globalState))
                rhs154_ = 999
                lhs114_ = globalState
                d_1000_v110_ = rhs153_
                lhs114_.f0 = rhs154_
            d_1003_v113_: _dafny.MultiSet
            d_1003_v113_ = _dafny.MultiSet([(d_968_v85_).f42])
            d_1004_v114_: D7
            d_1004_v114_ = D7_DC16((d_1003_v113_).cardinality, d_852_v0_)
            d_1005_v115_: _dafny.Seq
            d_1005_v115_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eytt"))
            if ((d_1004_v114_).cf23 if False else (d_1005_v115_) <= (d_1005_v115_)):
                (globalState).f1 = not(d_967_v84_.f40)
                d_1006_v116_: C1
                nw174_ = C1()
                nw174_.ctor__((d_952_v77_) | (d_952_v77_))
                d_1006_v116_ = nw174_
                d_1007_v117_: _dafny.Set
                d_1007_v117_ = _dafny.Set({True})
                d_1008_v119_: C3
                nw175_ = C3()
                def iife124_():
                    coll52_ = _dafny.Map()
                    compr_52_: int
                    for compr_52_ in _dafny.IntegerRange(-117, 796):
                        d_1009_v118_: int = compr_52_
                        if ((-117) <= (d_1009_v118_)) and ((d_1009_v118_) < (796)):
                            coll52_[default__.safeModuloInt(d_1009_v118_, len(_dafny.SeqWithoutIsStrInference([not(True)])))] = d_920_v53_
                    return _dafny.Map(coll52_)
                nw175_.ctor__((d_967_v84_).f39, (_dafny.Set({(d_968_v85_).f41, d_852_v0_, d_967_v84_.f40, (d_968_v85_).f41, (d_968_v85_).f41})).ispropersubset(d_1007_v117_), iife124_()
                )
                d_1008_v119_ = nw175_
                d_969_v86_ = d_969_v86_
                d_1010_v120_: _dafny.Map
                d_1010_v120_ = _dafny.Map({(d_968_v85_).f42: d_967_v84_.f40})
                d_1010_v120_ = (d_1010_v120_).set((d_953_v78_)[default__.safeIndex(d_920_v53_, len(d_953_v78_))], d_852_v0_)
            elif True:
                d_969_v86_ = (d_969_v86_ if d_852_v0_ else d_969_v86_)
                d_1011_v121_: _dafny.Array
                def lambda52_(d_1012_v85_):
                    def lambda53_(d_1013_i12_):
                        return (d_1012_v85_).f41

                    return lambda53_

                init30_ = lambda52_(d_968_v85_)
                nw176_ = _dafny.Array(None, 28)
                for i0_30_ in range(nw176_.length(0)):
                    nw176_[i0_30_] = init30_(i0_30_)
                d_1011_v121_ = nw176_
                d_1014_v122_: D11
                d_1014_v122_ = D11_DC26(d_1011_v121_)
                d_1014_v122_ = D11_DC26(d_1011_v121_)
                d_920_v53_ = (d_968_v85_).f42
                index141_ = default__.safeIndex(963, (d_1011_v121_).length(0))
                (d_1011_v121_)[index141_] = (d_968_v85_).f41
                index142_ = default__.safeIndex(963, (d_1011_v121_).length(0))
                (d_1011_v121_)[index142_] = (d_967_v84_.f40) or ((d_968_v85_).f41)
                (globalState).f0 = ((d_952_v77_)[((d_968_v85_).f42) < ((d_968_v85_).f42)] if (((d_968_v85_).f42) < ((d_968_v85_).f42)) in (d_952_v77_) else len(d_952_v77_))
        elif True:
            d_1015_v123_: _dafny.Set
            d_1015_v123_ = _dafny.Set({True})
            d_1016_v124_: _dafny.Map
            d_1016_v124_ = _dafny.Map({d_1015_v123_: default__.safeModuloInt(d_920_v53_, d_920_v53_)})
            d_1017_v125_: _dafny.Map
            d_1017_v125_ = _dafny.Map({d_1015_v123_: d_920_v53_})
            d_1016_v124_ = (d_1017_v125_)
            d_1018_v126_: _dafny.Array
            nw177_ = _dafny.Array(_dafny.MultiSet({}), 6)
            d_1018_v126_ = nw177_
            d_1019_v127_: _dafny.Map
            d_1019_v127_ = _dafny.Map({d_920_v53_: d_852_v0_})
            d_1020_v128_: _dafny.MultiSet
            d_1020_v128_ = _dafny.MultiSet([default__.fm37(d_852_v0_, globalState), d_1019_v127_])
            index143_ = default__.safeIndex(53, (d_1018_v126_).length(0))
            (d_1018_v126_)[index143_] = (d_1020_v128_)
            index144_ = default__.safeIndex(53, (d_1018_v126_).length(0))
            (d_1018_v126_)[index144_] = _dafny.MultiSet([(d_1019_v127_).set(d_920_v53_, not(d_852_v0_))])
            d_1021_v129_: str
            d_1021_v129_ = _dafny.CodePoint('a')
            rhs155_ = default__.fm35((d_920_v53_) + (190), globalState)
            rhs156_ = d_920_v53_
            rhs157_ = d_852_v0_
            lhs115_ = globalState
            lhs116_ = globalState
            d_1021_v129_ = rhs155_
            lhs115_.f0 = rhs156_
            lhs116_.f1 = rhs157_
            d_1022_v130_: _dafny.Seq
            d_1022_v130_ = _dafny.SeqWithoutIsStrInference([d_852_v0_])
            d_1023_v131_: _dafny.Map
            d_1023_v131_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psxxviri")): d_1022_v130_})
            d_1024_v132_: _dafny.Seq
            d_1024_v132_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
            d_1025_v133_: _dafny.Set
            d_1025_v133_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggkejsbpl"))})
            d_1026_v134_: _dafny.MultiSet
            d_1026_v134_ = _dafny.MultiSet([d_1024_v132_])
            d_1027_v135_: _dafny.Seq
            d_1027_v135_ = _dafny.SeqWithoutIsStrInference([d_920_v53_])
            d_1028_v136_: _dafny.Array
            nw178_ = _dafny.Array(None, 28)
            nw178_[int(0)] = len((_dafny.SeqWithoutIsStrInference([d_920_v53_, d_920_v53_, d_920_v53_, d_920_v53_, d_920_v53_])))
            nw178_[int(1)] = len((d_1023_v131_) | (_dafny.Map({d_1024_v132_: (d_1022_v130_).set(default__.safeIndex(515, len(d_1022_v130_)), d_852_v0_)})))
            nw178_[int(2)] = d_920_v53_
            nw178_[int(3)] = default__.fm2(globalState)
            nw178_[int(4)] = (0) - (default__.fm2(globalState))
            nw178_[int(5)] = len((_dafny.Set({d_1024_v132_})) | (d_1025_v133_))
            nw178_[int(6)] = d_920_v53_
            nw178_[int(7)] = d_920_v53_
            nw178_[int(8)] = (0) - (d_920_v53_)
            nw178_[int(9)] = d_920_v53_
            nw178_[int(10)] = d_920_v53_
            nw178_[int(11)] = ((d_1026_v134_)[d_1024_v132_] if (d_1024_v132_) in (d_1026_v134_) else d_920_v53_)
            nw178_[int(12)] = default__.safeModuloInt(d_920_v53_, d_920_v53_)
            nw178_[int(13)] = d_920_v53_
            nw178_[int(14)] = (0) - (((d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))])[default__.safeIndex(len(default__.fm30(d_920_v53_, d_1024_v132_, d_920_v53_, globalState)), len((d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))]))])
            nw178_[int(15)] = d_920_v53_
            nw178_[int(16)] = default__.fm2(globalState)
            nw178_[int(17)] = d_920_v53_
            nw178_[int(18)] = d_920_v53_
            nw178_[int(19)] = d_920_v53_
            nw178_[int(20)] = len(d_953_v78_)
            nw178_[int(21)] = d_920_v53_
            nw178_[int(22)] = d_920_v53_
            nw178_[int(23)] = ((0) - ((0) - (d_920_v53_))) + (d_920_v53_)
            nw178_[int(24)] = d_920_v53_
            nw178_[int(25)] = len((((d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_920_v53_])) for d_1029_i13_ in range(default__.abs(887))])))
            nw178_[int(26)] = d_920_v53_
            nw178_[int(27)] = d_920_v53_
            d_1028_v136_ = nw178_
            d_1030_v137_: D4
            d_1030_v137_ = D4_DC10(d_920_v53_)
            index145_ = default__.safeIndex(730, (d_1028_v136_).length(0))
            (d_1028_v136_)[index145_] = (d_1030_v137_).cf15
            d_1031_v138_: _dafny.MultiSet
            d_1031_v138_ = _dafny.MultiSet([d_1015_v123_])
            index146_ = default__.safeIndex(730, (d_1028_v136_).length(0))
            (d_1028_v136_)[index146_] = ((0) - ((0) - (((_dafny.MultiSet([default__.fm34(d_1022_v130_, d_920_v53_, globalState)])) | (d_1031_v138_)).cardinality))) * (d_920_v53_)
            d_952_v77_ = (d_952_v77_).set(d_852_v0_, (d_1028_v136_)[default__.safeIndex(730, (d_1028_v136_).length(0))])
        d_1032_v139_: _dafny.Array
        nw179_ = _dafny.Array(int(0), 9)
        d_1032_v139_ = nw179_
        index147_ = default__.safeIndex(940, (d_1032_v139_).length(0))
        (d_1032_v139_)[index147_] = d_920_v53_
        d_1033_v140_: _dafny.Array
        def lambda54_(d_1034_v0_):
            def lambda55_(d_1035_i14_):
                return _dafny.SeqWithoutIsStrInference([d_1034_v0_])

            return lambda55_

        init31_ = lambda54_(d_852_v0_)
        nw180_ = _dafny.Array(None, 3)
        for i0_31_ in range(nw180_.length(0)):
            nw180_[i0_31_] = init31_(i0_31_)
        d_1033_v140_ = nw180_
        d_1036_v141_: _dafny.Seq
        d_1036_v141_ = _dafny.SeqWithoutIsStrInference([not(d_852_v0_)])
        index148_ = default__.safeIndex(571, (d_1033_v140_).length(0))
        (d_1033_v140_)[index148_] = d_1036_v141_
        d_1037_v142_: str
        d_1037_v142_ = _dafny.CodePoint('i')
        d_1038_v143_: _dafny.Seq
        d_1038_v143_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjtjwqdko"))
        d_1039_v144_: _dafny.Map
        d_1039_v144_ = _dafny.Map({(d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))]: d_1038_v143_})
        index149_ = default__.safeIndex(940, (d_1032_v139_).length(0))
        index150_ = default__.safeIndex(571, (d_1033_v140_).length(0))
        rhs158_ = d_1032_v139_
        rhs159_ = (d_1037_v142_) not in (((d_1039_v144_)[(d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))]] if ((d_951_v76_)[default__.safeIndex(48, (d_951_v76_).length(0))]) in (d_1039_v144_) else d_1038_v143_))
        rhs160_ = (0) - (d_920_v53_)
        rhs161_ = ((d_1036_v141_) + (_dafny.SeqWithoutIsStrInference([d_852_v0_, d_852_v0_]))) + (d_1036_v141_)
        lhs117_ = globalState
        lhs118_ = d_1032_v139_
        lhs119_ = default__.safeIndex(940, (d_1032_v139_).length(0))
        lhs120_ = d_1033_v140_
        lhs121_ = default__.safeIndex(571, (d_1033_v140_).length(0))
        d_1032_v139_ = rhs158_
        lhs117_.f13 = rhs159_
        lhs118_[lhs119_] = rhs160_
        lhs120_[lhs121_] = rhs161_
        r0 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wt"))
        r1 = _dafny.MultiSet([(self).fm22(d_852_v0_, globalState), d_1037_v142_, d_1037_v142_])
        r2 = default__.safeDivisionInt((0) - (d_920_v53_), default__.safeDivisionInt(d_920_v53_, (d_1032_v139_)[default__.safeIndex(940, (d_1032_v139_).length(0))]))
        return r0, r1, r2


class C5(T1):
    def  __init__(self):
        self._f38: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f38):
        (self)._f38 = f38

    def fm8(self, p0, globalState):
        def iife125_():
            coll53_ = _dafny.Map()
            compr_53_: int
            for compr_53_ in (_dafny.Map({(self).f38: (self).f38})).keys.Elements:
                d_1040_v0_: int = compr_53_
                if (d_1040_v0_) in (_dafny.Map({(self).f38: (self).f38})):
                    coll53_[(d_1040_v0_) + (27)] = _dafny.Set({(self).f38})
            return _dafny.Map(coll53_)
        def iife126_():
            coll54_ = _dafny.Set()
            compr_54_: int
            for compr_54_ in (_dafny.Map({(self).f38: -422})).keys.Elements:
                d_1041_v1_: int = compr_54_
                if (d_1041_v1_) in (_dafny.Map({(self).f38: -422})):
                    coll54_ = coll54_.union(_dafny.Set([default__.safeDivisionInt(d_1041_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yiusrtu"))))]))
            return _dafny.Set(coll54_)
        return ((iife125_()
        ) | (_dafny.Map({(self).f38: _dafny.Set({(self).f38})}))) | (_dafny.Map({(self).f38: iife126_()
        }))

    def m7(self, p0, globalState):
        r0: bool = False
        d_1042_v0_: _dafny.Seq
        d_1042_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asyt"))
        (globalState).f2 = (0) - ((0) - (default__.safeModuloInt(((self).f38) - (len(d_1042_v0_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jb"))))))
        d_1043_v1_: bool
        d_1043_v1_ = False
        d_1044_i0_: int
        d_1044_i0_ = 0
        with _dafny.label("4"):
            while not(d_1043_v1_):
                with _dafny.c_label("4"):
                    if (d_1044_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_1044_i0_ = (d_1044_i0_) + (1)
                    (globalState).f16 = p0
                    d_1045_v2_: _dafny.Map
                    d_1045_v2_ = _dafny.Map({d_1043_v1_: p0})
                    d_1046_v3_: str
                    d_1046_v3_ = _dafny.CodePoint('a')
                    d_1047_v4_: _dafny.MultiSet
                    d_1047_v4_ = _dafny.MultiSet([d_1046_v3_])
                    d_1048_v5_: D4
                    d_1048_v5_ = D4_DC10(p0)
                    d_1049_v6_: _dafny.Seq
                    d_1049_v6_ = _dafny.SeqWithoutIsStrInference([True, d_1043_v1_])
                    d_1050_v7_: _dafny.Seq
                    d_1050_v7_ = _dafny.SeqWithoutIsStrInference([(self).f38, (self).f38, (self).f38, p0])
                    d_1051_v8_: _dafny.Map
                    d_1051_v8_ = _dafny.Map({len(d_1049_v6_): d_1050_v7_})
                    d_1052_v9_: _dafny.Seq
                    d_1052_v9_ = _dafny.SeqWithoutIsStrInference([p0, len(d_1051_v8_)])
                    d_1053_v10_: _dafny.MultiSet
                    d_1053_v10_ = _dafny.MultiSet([d_1052_v9_])
                    d_1054_v11_: _dafny.Array
                    nw181_ = _dafny.Array(None, 15)
                    nw181_[int(0)] = p0
                    nw181_[int(1)] = (840) - (p0)
                    nw181_[int(2)] = (((d_1045_v2_)[d_1043_v1_] if (d_1043_v1_) in (d_1045_v2_) else (self).f38)) - (789)
                    nw181_[int(3)] = p0
                    nw181_[int(4)] = default__.safeDivisionInt(((d_1047_v4_)[d_1046_v3_] if (d_1046_v3_) in (d_1047_v4_) else p0), p0)
                    nw181_[int(5)] = (self).f38
                    nw181_[int(6)] = (self).f38
                    nw181_[int(7)] = (d_1048_v5_).cf15
                    nw181_[int(8)] = (p0) + (p0)
                    nw181_[int(9)] = (0) - (len(_dafny.Set({p0})))
                    nw181_[int(10)] = (0) - (len(d_1045_v2_))
                    nw181_[int(11)] = (self).f38
                    nw181_[int(12)] = (d_1053_v10_).cardinality
                    nw181_[int(13)] = (self).f38
                    nw181_[int(14)] = ((self).f38) * (p0)
                    d_1054_v11_ = nw181_
                    index151_ = default__.safeIndex(792, (d_1054_v11_).length(0))
                    (d_1054_v11_)[index151_] = p0
                    d_1055_v12_: D7
                    d_1055_v12_ = D7_DC17(d_1043_v1_, d_1043_v1_, (self).f38)
                    d_1056_v13_: _dafny.Seq
                    d_1056_v13_ = _dafny.SeqWithoutIsStrInference([d_1055_v12_])
                    d_1057_v14_: _dafny.Set
                    d_1057_v14_ = _dafny.Set({p0})
                    index152_ = default__.safeIndex(792, (d_1054_v11_).length(0))
                    (d_1054_v11_)[index152_] = ((0) - (len(d_1056_v13_))) * (len((_dafny.Set({(self).f38})).intersection(d_1057_v14_)))
                    d_1058_v15_: _dafny.Set
                    d_1058_v15_ = _dafny.Set({d_1042_v0_})
                    d_1059_v16_: D3
                    d_1059_v16_ = D3_DC7(d_1042_v0_, (self).f38, len(d_1057_v14_), len(d_1058_v15_), d_1049_v6_)
                    d_1060_v17_: _dafny.Seq
                    d_1060_v17_ = default__.fm21(d_1046_v3_, (0) - ((self).f38), d_1059_v16_, globalState)
                    d_1061_v18_: _dafny.Map
                    d_1061_v18_ = _dafny.Map({d_1042_v0_: ((d_1060_v17_)) + (d_1052_v9_)})
                    d_1061_v18_ = (d_1061_v18_).set((d_1042_v0_) + (d_1042_v0_), d_1052_v9_)
                    d_1062_v19_: D0
                    d_1062_v19_ = D0_DC1(d_1054_v11_, d_1046_v3_, (self).f38)
                    d_1063_v20_: D0
                    d_1063_v20_ = D0_DC2(d_1062_v19_)
                    pat_let_tv31_ = d_1062_v19_
                    def iife127_(_pat_let36_0):
                        def iife128_(d_1064_dt__update__tmp_h0_):
                            def iife129_(_pat_let37_0):
                                def iife130_(d_1065_dt__update_hcf4_h0_):
                                    return D0_DC2(d_1065_dt__update_hcf4_h0_)
                                return iife130_(_pat_let37_0)
                            return iife129_(pat_let_tv31_)
                        return iife128_(_pat_let36_0)
                    rhs162_ = (self).f38
                    rhs163_ = iife127_(d_1063_v20_)
                    lhs122_ = globalState
                    lhs122_.f22 = rhs162_
                    d_1063_v20_ = rhs163_
                    pass
            pass
        d_1066_v21_: _dafny.Seq
        d_1066_v21_ = _dafny.SeqWithoutIsStrInference([d_1043_v1_, d_1043_v1_])
        d_1067_i1_: int
        d_1067_i1_ = 0
        with _dafny.label("5"):
            while not(not(((0) - (default__.safeDivisionInt(p0, p0))) == (len((d_1066_v21_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([default__.fm0(p0, globalState), False, d_1043_v1_])), len(d_1066_v21_)), d_1043_v1_))))):
                with _dafny.c_label("5"):
                    if (d_1067_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_1067_i1_ = (d_1067_i1_) + (1)
                    if (d_1043_v1_) and (d_1043_v1_):
                        d_1068_v22_: _dafny.Map
                        d_1068_v22_ = _dafny.Map({d_1043_v1_: d_1043_v1_})
                        d_1068_v22_ = (d_1068_v22_).set((d_1043_v1_) == (d_1043_v1_), (d_1043_v1_ if d_1043_v1_ else d_1043_v1_))
                        d_1069_v23_: _dafny.Array
                        nw182_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
                        d_1069_v23_ = nw182_
                        d_1069_v23_ = (d_1069_v23_ if d_1043_v1_ else (d_1069_v23_ if d_1043_v1_ else d_1069_v23_))
                        d_1070_v24_: _dafny.Array
                        nw183_ = _dafny.Array(False, 17)
                        d_1070_v24_ = nw183_
                        index153_ = default__.safeIndex(771, (d_1070_v24_).length(0))
                        (d_1070_v24_)[index153_] = ((0) - ((self).f38)) != ((self).f38)
                        index154_ = default__.safeIndex(771, (d_1070_v24_).length(0))
                        rhs164_ = d_1043_v1_
                        rhs165_ = ((self).f38) * ((self).f38)
                        lhs123_ = d_1070_v24_
                        lhs124_ = default__.safeIndex(771, (d_1070_v24_).length(0))
                        lhs125_ = globalState
                        lhs123_[lhs124_] = rhs164_
                        lhs125_.f2 = rhs165_
                        index155_ = default__.safeIndex(771, (d_1070_v24_).length(0))
                        (d_1070_v24_)[index155_] = ((d_1070_v24_)[default__.safeIndex(771, (d_1070_v24_).length(0))] if (d_1070_v24_)[default__.safeIndex(771, (d_1070_v24_).length(0))] else ((self).f38) > (p0))
                        d_1071_v25_: _dafny.Set
                        d_1071_v25_ = _dafny.Set({(self).f38})
                        (globalState).f13 = (d_1071_v25_).issubset(d_1071_v25_)
                    elif True:
                        d_1072_v26_: D4
                        d_1072_v26_ = D4_DC10(p0)
                        d_1073_v27_: _dafny.Map
                        d_1073_v27_ = _dafny.Map({d_1072_v26_: p0})
                        d_1073_v27_ = d_1073_v27_
                        (globalState).f2 = p0
                        d_1074_v28_: _dafny.Array
                        nw184_ = _dafny.Array(False, 28)
                        d_1074_v28_ = nw184_
                        d_1075_v29_: _dafny.Seq
                        d_1075_v29_ = _dafny.SeqWithoutIsStrInference([d_1074_v28_])
                        d_1074_v28_ = (d_1075_v29_)[default__.safeIndex((self).f38, len(d_1075_v29_))]
                        d_1076_v30_: _dafny.Seq
                        d_1076_v30_ = _dafny.SeqWithoutIsStrInference([p0])
                        (globalState).f16 = (d_1076_v30_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1077_i2_ in range(default__.abs(-983))])), len(d_1076_v30_))]
                        (globalState).f16 = default__.safeDivisionInt(default__.safeDivisionInt(len(d_1042_v0_), (self).f38), p0)
                    d_1078_v31_: _dafny.Map
                    d_1078_v31_ = _dafny.Map({d_1043_v1_: d_1043_v1_})
                    d_1079_v32_: _dafny.Set
                    d_1079_v32_ = _dafny.Set({p0})
                    d_1080_v33_: _dafny.Map
                    d_1080_v33_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1081_i3_ in range(default__.abs(240))]): _dafny.CodePoint('w')})
                    d_1078_v31_ = (d_1078_v31_).set((d_1079_v32_).issubset(d_1079_v32_), default__.fm0(len(d_1080_v33_), globalState))
                    d_1082_v35_: _dafny.Array
                    def lambda56_(d_1083_v1_):
                        def lambda57_(d_1084_i4_):
                            def iife131_():
                                coll55_ = _dafny.Map()
                                compr_55_: int
                                for compr_55_ in _dafny.IntegerRange(345, 913):
                                    d_1085_v34_: int = compr_55_
                                    if ((345) <= (d_1085_v34_)) and ((d_1085_v34_) < (913)):
                                        coll55_[default__.safeModuloInt(d_1085_v34_, (self).f38)] = len(_dafny.Set({(D10_DC25((self).f38, 202, d_1083_v1_, d_1083_v1_)).cf37}))
                                return _dafny.Map(coll55_)
                            return (d_1084_i4_) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([iife131_()
                            ]))])))

                        return lambda57_

                    init32_ = lambda56_(d_1043_v1_)
                    nw185_ = _dafny.Array(None, 6)
                    for i0_32_ in range(nw185_.length(0)):
                        nw185_[i0_32_] = init32_(i0_32_)
                    d_1082_v35_ = nw185_
                    d_1086_v36_: C3
                    nw186_ = C3()
                    nw186_.ctor__(d_1082_v35_, d_1043_v1_, _dafny.Map({740: p0}))
                    d_1086_v36_ = nw186_
                    if d_1086_v36_.f40:
                        d_1087_v37_: _dafny.Set
                        d_1087_v37_ = _dafny.Set({d_1043_v1_, d_1043_v1_, d_1086_v36_.f40, not(d_1086_v36_.f40), d_1086_v36_.f40})
                        (globalState).f0 = default__.safeDivisionInt((0) - (len(d_1087_v37_)), (len(d_1042_v0_)) - ((self).f38))
                        d_1088_v38_: C4
                        nw187_ = C4()
                        nw187_.ctor__()
                        d_1088_v38_ = nw187_
                        d_1089_v39_: _dafny.Set
                        d_1089_v39_ = _dafny.Set({d_1088_v38_})
                        d_1090_v40_: D5
                        d_1090_v40_ = D5_DC13(len(d_1089_v39_))
                        pat_let_tv32_ = p0
                        d_1091_v41_: _dafny.Map
                        def iife132_(_pat_let38_0):
                            def iife133_(d_1092_dt__update__tmp_h1_):
                                def iife134_(_pat_let39_0):
                                    def iife135_(d_1093_dt__update_hcf19_h0_):
                                        return D5_DC13(d_1093_dt__update_hcf19_h0_)
                                    return iife135_(_pat_let39_0)
                                return iife134_(pat_let_tv32_)
                            return iife133_(_pat_let38_0)
                        d_1091_v41_ = _dafny.Map({(iife132_(d_1090_v40_)).cf19: p0})
                        d_1094_v42_: _dafny.Map
                        d_1094_v42_ = _dafny.Map({d_1043_v1_: len(d_1091_v41_)})
                        d_1095_v43_: C1
                        nw188_ = C1()
                        nw188_.ctor__(d_1094_v42_)
                        d_1095_v43_ = nw188_
                        d_1096_v44_: _dafny.Array
                        def lambda58_(d_1097_v36_):
                            def lambda59_(d_1098_i5_):
                                return d_1097_v36_.f40

                            return lambda59_

                        init33_ = lambda58_(d_1086_v36_)
                        nw189_ = _dafny.Array(None, 8)
                        for i0_33_ in range(nw189_.length(0)):
                            nw189_[i0_33_] = init33_(i0_33_)
                        d_1096_v44_ = nw189_
                        d_1096_v44_ = d_1096_v44_
                        d_1099_v45_: C2
                        nw190_ = C2()
                        nw190_.ctor__(d_1086_v36_.f40, p0)
                        d_1099_v45_ = nw190_
                        d_1096_v44_ = d_1096_v44_
                    elif True:
                        d_1100_v47_: _dafny.Set
                        d_1100_v47_ = _dafny.Set({(d_1086_v36_).f39, (d_1086_v36_).f39, d_1082_v35_})
                        d_1101_v48_: _dafny.Map
                        d_1101_v48_ = _dafny.Map({697: d_1086_v36_.f40})
                        d_1102_v49_: _dafny.Array
                        nw191_ = _dafny.Array(None, 25)
                        nw191_[int(0)] = False
                        nw191_[int(1)] = not(True)
                        nw191_[int(2)] = (p0) >= (p0)
                        def iife136_():
                            coll56_ = _dafny.Set()
                            compr_56_: int
                            for compr_56_ in _dafny.IntegerRange(-53, 257):
                                d_1103_v46_: int = compr_56_
                                if ((-53) <= (d_1103_v46_)) and ((d_1103_v46_) < (257)):
                                    coll56_ = coll56_.union(_dafny.Set([default__.safeDivisionInt(d_1103_v46_, p0)]))
                            return _dafny.Set(coll56_)
                        nw191_[int(3)] = (iife136_()
                        ) != (d_1079_v32_)
                        nw191_[int(4)] = (d_1100_v47_).issubset(d_1100_v47_)
                        nw191_[int(5)] = (len(d_1101_v48_)) < ((self).f38)
                        nw191_[int(6)] = d_1086_v36_.f40
                        nw191_[int(7)] = d_1043_v1_
                        nw191_[int(8)] = (False) or (d_1086_v36_.f40)
                        nw191_[int(9)] = not(d_1043_v1_)
                        nw191_[int(10)] = True
                        nw191_[int(11)] = d_1043_v1_
                        nw191_[int(12)] = d_1086_v36_.f40
                        nw191_[int(13)] = d_1086_v36_.f40
                        nw191_[int(14)] = d_1086_v36_.f40
                        nw191_[int(15)] = (p0) == ((self).f38)
                        nw191_[int(16)] = not(d_1086_v36_.f40)
                        nw191_[int(17)] = d_1043_v1_
                        nw191_[int(18)] = d_1086_v36_.f40
                        nw191_[int(19)] = not(d_1043_v1_)
                        nw191_[int(20)] = d_1043_v1_
                        nw191_[int(21)] = d_1086_v36_.f40
                        nw191_[int(22)] = d_1086_v36_.f40
                        nw191_[int(23)] = d_1086_v36_.f40
                        nw191_[int(24)] = d_1086_v36_.f40
                        d_1102_v49_ = nw191_
                        index156_ = default__.safeIndex(596, (d_1102_v49_).length(0))
                        (d_1102_v49_)[index156_] = d_1086_v36_.f40
                        d_1104_v50_: _dafny.Map
                        d_1104_v50_ = _dafny.Map({(0) - (p0): p0})
                        index157_ = default__.safeIndex(596, (d_1102_v49_).length(0))
                        (d_1102_v49_)[index157_] = (default__.fm2(globalState)) not in (d_1104_v50_)
                        index158_ = default__.safeIndex(596, (d_1102_v49_).length(0))
                        (d_1102_v49_)[index158_] = d_1086_v36_.f40
                        d_1105_v51_: str
                        d_1105_v51_ = _dafny.CodePoint('a')
                        d_1106_v52_: D4
                        d_1106_v52_ = D4_DC8(d_1105_v51_)
                        d_1107_v53_: D4
                        d_1107_v53_ = D4_DC8((d_1105_v51_ if d_1043_v1_ else (d_1106_v52_).cf13))
                        d_1108_v54_: _dafny.MultiSet
                        d_1108_v54_ = _dafny.MultiSet([not (d_1043_v1_) or ((d_1102_v49_)[default__.safeIndex(596, (d_1102_v49_).length(0))])])
                        d_1109_v55_: _dafny.Map
                        d_1109_v55_ = _dafny.Map({p0: d_1108_v54_})
                        rhs166_ = default__.fm32(d_1086_v36_.f40, globalState)
                        rhs167_ = ((self).f38) >= (p0)
                        rhs168_ = default__.safeDivisionInt(default__.safeDivisionInt(983, default__.fm2(globalState)), -450)
                        rhs169_ = d_1086_v36_.f40
                        rhs170_ = ((((d_1109_v55_)[(self).f38] if ((self).f38) in (d_1109_v55_) else d_1108_v54_)) - (d_1108_v54_)) - (((d_1108_v54_).set((d_1066_v21_)[default__.safeIndex(p0, len(d_1066_v21_))], default__.abs((self).f38))).intersection(d_1108_v54_))
                        lhs126_ = globalState
                        lhs127_ = globalState
                        d_1107_v53_ = rhs166_
                        d_1043_v1_ = rhs167_
                        lhs126_.f2 = rhs168_
                        lhs127_.f1 = rhs169_
                        d_1108_v54_ = rhs170_
                        d_1110_v56_: _dafny.MultiSet
                        d_1110_v56_ = _dafny.MultiSet([d_1105_v51_])
                        (globalState).f22 = (d_1110_v56_).cardinality
                        def iife137_():
                            coll57_ = _dafny.Map()
                            compr_57_: int
                            for compr_57_ in (d_1079_v32_).Elements:
                                d_1111_v57_: int = compr_57_
                                if (d_1111_v57_) in (d_1079_v32_):
                                    coll57_[(d_1111_v57_) * ((self).f38)] = (d_1102_v49_)[default__.safeIndex(596, (d_1102_v49_).length(0))]
                            return _dafny.Map(coll57_)
                        d_1101_v48_ = (iife137_()
                        ) | ((d_1101_v48_) | ((d_1101_v48_).set(p0, d_1086_v36_.f40)))
                    pass
            pass
        d_1112_v58_: _dafny.Array
        nw192_ = _dafny.Array(int(0), 13)
        d_1112_v58_ = nw192_
        index159_ = default__.safeIndex(599, (d_1112_v58_).length(0))
        (d_1112_v58_)[index159_] = default__.safeDivisionInt(542, p0)
        d_1113_v59_: _dafny.Set
        d_1113_v59_ = _dafny.Set({d_1043_v1_, d_1043_v1_, not(d_1043_v1_), d_1043_v1_})
        d_1114_v60_: D10
        d_1114_v60_ = D10_DC24(d_1113_v59_)
        d_1115_v61_: _dafny.Set
        d_1115_v61_ = _dafny.Set({d_1114_v60_})
        d_1116_v62_: _dafny.Map
        d_1116_v62_ = _dafny.Map({d_1043_v1_: d_1043_v1_})
        d_1117_v63_: _dafny.Array
        nw193_ = _dafny.Array(None, 15)
        nw193_[int(0)] = d_1043_v1_
        nw193_[int(1)] = d_1043_v1_
        nw193_[int(2)] = d_1043_v1_
        nw193_[int(3)] = d_1043_v1_
        nw193_[int(4)] = d_1043_v1_
        nw193_[int(5)] = d_1043_v1_
        nw193_[int(6)] = default__.fm0(p0, globalState)
        nw193_[int(7)] = d_1043_v1_
        nw193_[int(8)] = default__.fm0(p0, globalState)
        nw193_[int(9)] = d_1043_v1_
        nw193_[int(10)] = d_1043_v1_
        nw193_[int(11)] = (d_1115_v61_).isdisjoint(d_1115_v61_)
        nw193_[int(12)] = (D8_DC20(((d_1116_v62_)[False] if (False) in (d_1116_v62_) else False))).cf31
        nw193_[int(13)] = d_1043_v1_
        nw193_[int(14)] = (d_1043_v1_) == (d_1043_v1_)
        d_1117_v63_ = nw193_
        d_1118_v64_: _dafny.Map
        d_1118_v64_ = _dafny.Map({d_1043_v1_: _dafny.CodePoint('t')})
        d_1119_v65_: D10
        d_1119_v65_ = D10_DC25(p0, p0, d_1043_v1_, not(False))
        d_1120_v66_: _dafny.Map
        d_1120_v66_ = _dafny.Map({d_1117_v63_: d_1117_v63_})
        d_1121_v67_: _dafny.Array
        def lambda60_(d_1122_i6_):
            return _dafny.CodePoint('v')

        init34_ = lambda60_
        nw194_ = _dafny.Array(None, 8)
        for i0_34_ in range(nw194_.length(0)):
            nw194_[i0_34_] = init34_(i0_34_)
        d_1121_v67_ = nw194_
        d_1123_v68_: D0
        d_1123_v68_ = D0_DC2(D0_DC0(d_1121_v67_))
        d_1124_v69_: D0
        d_1124_v69_ = D0_DC2(d_1123_v68_)
        d_1125_v70_: D0
        d_1125_v70_ = D0_DC2((d_1124_v69_).cf4)
        d_1126_v71_: D0
        d_1126_v71_ = D0_DC2(d_1125_v70_)
        d_1127_v72_: _dafny.Set
        d_1127_v72_ = _dafny.Set({d_1124_v69_})
        index160_ = default__.safeIndex(599, (d_1112_v58_).length(0))
        rhs171_ = default__.safeModuloInt(p0, (0) - (default__.safeModuloInt(p0, len(d_1118_v64_))))
        rhs172_ = (d_1119_v65_).cf38
        rhs173_ = (((self).f38) - (p0)) + (default__.fm2(globalState))
        rhs174_ = ((d_1120_v66_)[d_1117_v63_] if (d_1117_v63_) in (d_1120_v66_) else d_1117_v63_)
        rhs175_ = (d_1127_v72_).ispropersubset((d_1127_v72_ if d_1043_v1_ else _dafny.Set({d_1124_v69_})))
        lhs128_ = d_1112_v58_
        lhs129_ = default__.safeIndex(599, (d_1112_v58_).length(0))
        lhs130_ = globalState
        lhs131_ = globalState
        lhs132_ = globalState
        lhs128_[lhs129_] = rhs171_
        lhs130_.f13 = rhs172_
        lhs131_.f22 = rhs173_
        d_1117_v63_ = rhs174_
        lhs132_.f13 = rhs175_
        d_1128_v73_: _dafny.Seq
        d_1128_v73_ = _dafny.SeqWithoutIsStrInference([(self).f38, default__.fm2(globalState)])
        (globalState).f1 = (d_1128_v73_) != (d_1128_v73_)
        if d_1043_v1_:
            d_1129_v74_: _dafny.Set
            d_1129_v74_ = _dafny.Set({d_1112_v58_, d_1112_v58_, d_1112_v58_})
            d_1130_v75_: C2
            nw195_ = C2()
            nw195_.ctor__(d_1043_v1_, len(d_1129_v74_))
            d_1130_v75_ = nw195_
            d_1131_v76_: str
            d_1131_v76_ = _dafny.CodePoint('r')
            index161_ = default__.safeIndex(630, (d_1121_v67_).length(0))
            (d_1121_v67_)[index161_] = d_1131_v76_
            index162_ = default__.safeIndex(630, (d_1121_v67_).length(0))
            (d_1121_v67_)[index162_] = d_1131_v76_
            d_1132_v77_: C4
            nw196_ = C4()
            nw196_.ctor__()
            d_1132_v77_ = nw196_
            d_1133_v78_: _dafny.Map
            d_1133_v78_ = _dafny.Map({d_1132_v77_: d_1042_v0_})
            d_1133_v78_ = (d_1133_v78_) | ((d_1133_v78_) | (d_1133_v78_))
            d_1134_v79_: _dafny.MultiSet
            d_1134_v79_ = _dafny.MultiSet([p0, (d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))]])
            d_1135_v80_: _dafny.Map
            d_1135_v80_ = _dafny.Map({(d_1134_v79_).set(p0, default__.abs(309)): (d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))]})
            d_1136_v81_: bool
            d_1137_v82_: D5
            d_1138_v83_: bool
            out38_: bool
            out39_: D5
            out40_: bool
            out38_, out39_, out40_ = (self).m15((0) - ((0) - (((d_1135_v80_)[d_1134_v79_] if (d_1134_v79_) in (d_1135_v80_) else len((d_1042_v0_).set(default__.safeIndex((d_1130_v75_).f42, len(d_1042_v0_)), _dafny.CodePoint('c')))))), p0, globalState)
            d_1136_v81_ = out38_
            d_1137_v82_ = out39_
            d_1138_v83_ = out40_
            index163_ = default__.safeIndex(599, (d_1112_v58_).length(0))
            (d_1112_v58_)[index163_] = default__.safeDivisionInt((self).f38, (d_1130_v75_).f42)
        elif True:
            d_1139_v84_: _dafny.Array
            nw197_ = _dafny.Array(None, 1)
            nw197_[int(0)] = d_1112_v58_
            d_1139_v84_ = nw197_
            index164_ = default__.safeIndex(456, (d_1139_v84_).length(0))
            (d_1139_v84_)[index164_] = d_1112_v58_
            index165_ = default__.safeIndex(456, (d_1139_v84_).length(0))
            (d_1139_v84_)[index165_] = d_1112_v58_
            if d_1043_v1_:
                d_1140_v85_: _dafny.Map
                d_1140_v85_ = _dafny.Map({len(d_1042_v0_): (d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))]})
                (globalState).f12 = d_1140_v85_
                d_1141_v86_: _dafny.Map
                d_1141_v86_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1043_v1_: (d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))]}) for d_1142_i7_ in range(default__.abs(-471))]): D7_DC16(500, d_1043_v1_)})
                d_1143_v87_: _dafny.Map
                d_1143_v87_ = _dafny.Map({d_1043_v1_: (d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))]})
                d_1144_v88_: D7
                d_1144_v88_ = D7_DC16(p0, default__.fm0((self).f38, globalState))
                d_1141_v86_ = (d_1141_v86_).set(_dafny.SeqWithoutIsStrInference([d_1143_v87_]), d_1144_v88_)
                (globalState).f1 = ((self).f38) >= ((d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))])
                (globalState).f13 = d_1043_v1_
                d_1145_v89_: _dafny.MultiSet
                d_1145_v89_ = _dafny.MultiSet([(p0) >= ((self).f38), d_1043_v1_, d_1043_v1_])
                d_1145_v89_ = ((d_1145_v89_) | (d_1145_v89_)) - ((d_1145_v89_).intersection(d_1145_v89_))
            elif True:
                d_1146_v90_: str
                d_1146_v90_ = _dafny.CodePoint('s')
                d_1146_v90_ = d_1146_v90_
                (globalState).f16 = (0) - (len(default__.fm30((0) - (default__.fm2(globalState)), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "quu"))) + (default__.fm30((d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))], d_1042_v0_, (self).f38, globalState)), 356, globalState)))
                d_1147_v91_: _dafny.Set
                d_1147_v91_ = _dafny.Set({p0})
                d_1148_v92_: D3
                d_1148_v92_ = D3_DC7(d_1042_v0_, len(d_1042_v0_), len(d_1147_v91_), (self).f38, d_1066_v21_)
                (globalState).f22 = len((d_1148_v92_).cf8)
                d_1128_v73_ = (_dafny.SeqWithoutIsStrInference([610])).set(default__.safeIndex(default__.fm2(globalState), len(_dafny.SeqWithoutIsStrInference([610]))), (p0) - ((0) - ((self).f38)))
                d_1149_v93_: _dafny.Seq
                d_1149_v93_ = d_1128_v73_
                d_1150_v94_: _dafny.Map
                d_1150_v94_ = _dafny.Map({d_1043_v1_: p0})
                d_1151_v95_: _dafny.Array
                nw198_ = _dafny.Array(None, 23)
                nw198_[int(0)] = (d_1128_v73_) + (d_1128_v73_)
                nw198_[int(1)] = d_1128_v73_
                nw198_[int(2)] = d_1128_v73_
                nw198_[int(3)] = (d_1128_v73_) + (d_1128_v73_)
                nw198_[int(4)] = d_1128_v73_
                nw198_[int(5)] = (d_1128_v73_).set(default__.safeIndex((d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))], len(d_1128_v73_)), (self).f38)
                nw198_[int(6)] = (d_1149_v93_)
                nw198_[int(7)] = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
                nw198_[int(8)] = d_1128_v73_
                nw198_[int(9)] = (d_1128_v73_) + (d_1128_v73_)
                nw198_[int(10)] = (d_1128_v73_) + (d_1128_v73_)
                nw198_[int(11)] = (d_1128_v73_) + (_dafny.SeqWithoutIsStrInference([(d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))] for d_1152_i8_ in range(default__.abs(959))]))
                nw198_[int(12)] = (d_1128_v73_) + (d_1128_v73_)
                nw198_[int(13)] = d_1128_v73_
                nw198_[int(14)] = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({341: d_1043_v1_})) for d_1153_i9_ in range(default__.abs(-943))])
                nw198_[int(15)] = (d_1128_v73_) + (d_1128_v73_)
                nw198_[int(16)] = d_1128_v73_
                nw198_[int(17)] = _dafny.SeqWithoutIsStrInference([(0) - ((self).f38)])
                nw198_[int(18)] = _dafny.SeqWithoutIsStrInference([p0, p0])
                nw198_[int(19)] = (d_1128_v73_).set(default__.safeIndex((d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))], len(d_1128_v73_)), 542)
                nw198_[int(20)] = _dafny.SeqWithoutIsStrInference([p0, (self).f38])
                nw198_[int(21)] = d_1128_v73_
                nw198_[int(22)] = ((d_1128_v73_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([True])), len(d_1128_v73_)), len(d_1150_v94_))) + (d_1128_v73_)
                d_1151_v95_ = nw198_
                d_1151_v95_ = d_1151_v95_
            d_1154_v96_: _dafny.Array
            nw199_ = _dafny.Array(_dafny.Map({}), 15)
            d_1154_v96_ = nw199_
            index166_ = default__.safeIndex(77, (d_1154_v96_).length(0))
            (d_1154_v96_)[index166_] = (d_1120_v66_) | (d_1120_v66_)
            index167_ = default__.safeIndex(77, (d_1154_v96_).length(0))
            (d_1154_v96_)[index167_] = d_1120_v66_
            (globalState).f2 = default__.safeDivisionInt(p0, default__.safeDivisionInt((d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))], (d_1112_v58_)[default__.safeIndex(599, (d_1112_v58_).length(0))]))
            (globalState).f21 = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddqgct"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usxodm"))))
        r0 = True
        return r0

    def m15(self, p0, p1, globalState):
        r0: bool = False
        r1: D5 = D5.default()()
        r2: bool = False
        d_1155_v0_: _dafny.Array
        nw200_ = _dafny.Array(_dafny.MultiSet({}), 12)
        d_1155_v0_ = nw200_
        d_1156_v1_: D2
        d_1156_v1_ = D2_DC4(d_1155_v0_)
        source15_ = d_1156_v1_
        if source15_.is_DC5:
            d_1157_v2_: bool
            d_1157_v2_ = False
            d_1158_v3_: _dafny.Seq
            d_1158_v3_ = _dafny.SeqWithoutIsStrInference([d_1157_v2_])
            d_1159_v4_: _dafny.Seq
            d_1159_v4_ = _dafny.SeqWithoutIsStrInference([(self).f38])
            d_1160_v5_: _dafny.Array
            nw201_ = _dafny.Array(None, 12)
            nw201_[int(0)] = d_1157_v2_
            nw201_[int(1)] = False
            nw201_[int(2)] = ((d_1158_v3_)[default__.safeIndex(962, len(d_1158_v3_))]) == (d_1157_v2_)
            nw201_[int(3)] = d_1157_v2_
            nw201_[int(4)] = d_1157_v2_
            nw201_[int(5)] = d_1157_v2_
            nw201_[int(6)] = d_1157_v2_
            nw201_[int(7)] = d_1157_v2_
            nw201_[int(8)] = ((0) - ((self).f38)) not in ((d_1159_v4_))
            nw201_[int(9)] = True
            nw201_[int(10)] = d_1157_v2_
            nw201_[int(11)] = d_1157_v2_
            d_1160_v5_ = nw201_
            index168_ = default__.safeIndex(393, (d_1160_v5_).length(0))
            (d_1160_v5_)[index168_] = d_1157_v2_
            d_1161_v6_: _dafny.Array
            nw202_ = _dafny.Array(int(0), 20)
            d_1161_v6_ = nw202_
            index169_ = default__.safeIndex(393, (d_1160_v5_).length(0))
            rhs176_ = (p0) * (default__.fm2(globalState))
            rhs177_ = d_1157_v2_
            rhs178_ = d_1157_v2_
            rhs179_ = d_1161_v6_
            lhs133_ = globalState
            lhs134_ = d_1160_v5_
            lhs135_ = default__.safeIndex(393, (d_1160_v5_).length(0))
            lhs133_.f5 = rhs176_
            lhs134_[lhs135_] = rhs177_
            r2 = rhs178_
            d_1161_v6_ = rhs179_
            d_1162_v7_: _dafny.Map
            d_1162_v7_ = _dafny.Map({d_1160_v5_: d_1161_v6_})
            d_1162_v7_ = (d_1162_v7_).set(d_1160_v5_, d_1161_v6_)
            d_1163_v8_: _dafny.Seq
            d_1163_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tb"))
            d_1164_v9_: _dafny.Array
            nw203_ = _dafny.Array(_dafny.Seq({}), 29)
            d_1164_v9_ = nw203_
            d_1165_v10_: _dafny.Map
            d_1165_v10_ = _dafny.Map({(d_1163_v8_) == (d_1163_v8_): d_1164_v9_})
            d_1165_v10_ = (d_1165_v10_).set((d_1160_v5_)[default__.safeIndex(393, (d_1160_v5_).length(0))], ((d_1165_v10_)[d_1157_v2_] if (d_1157_v2_) in (d_1165_v10_) else d_1164_v9_))
            d_1166_v11_: _dafny.Map
            d_1166_v11_ = _dafny.Map({d_1157_v2_: p0})
            d_1167_v12_: C1
            nw204_ = C1()
            nw204_.ctor__(d_1166_v11_)
            d_1167_v12_ = nw204_
        elif True:
            d_1168___mcc_h0_ = source15_.cf6
            d_1169_cf6_ = d_1168___mcc_h0_
            d_1170_v13_: bool
            d_1170_v13_ = True
            if d_1170_v13_:
                d_1171_v14_: _dafny.Set
                d_1171_v14_ = _dafny.Set({519, p1})
                d_1172_v15_: _dafny.Array
                nw205_ = _dafny.Array(None, 13)
                nw205_[int(0)] = (d_1170_v13_) or (d_1170_v13_)
                nw205_[int(1)] = d_1170_v13_
                nw205_[int(2)] = d_1170_v13_
                nw205_[int(3)] = d_1170_v13_
                nw205_[int(4)] = d_1170_v13_
                nw205_[int(5)] = d_1170_v13_
                nw205_[int(6)] = not(d_1170_v13_)
                nw205_[int(7)] = d_1170_v13_
                nw205_[int(8)] = d_1170_v13_
                nw205_[int(9)] = default__.fm0(p0, globalState)
                nw205_[int(10)] = (d_1171_v14_).isdisjoint(default__.fm42(globalState))
                nw205_[int(11)] = False
                nw205_[int(12)] = d_1170_v13_
                d_1172_v15_ = nw205_
                d_1173_v16_: _dafny.Array
                nw206_ = _dafny.Array(int(0), 25)
                d_1173_v16_ = nw206_
                index170_ = default__.safeIndex(306, (d_1173_v16_).length(0))
                (d_1173_v16_)[index170_] = p1
                index171_ = default__.safeIndex(368, (d_1173_v16_).length(0))
                (d_1173_v16_)[index171_] = (self).f38
                d_1174_v17_: str
                d_1174_v17_ = _dafny.CodePoint('q')
                d_1175_v18_: _dafny.Seq
                d_1175_v18_ = _dafny.SeqWithoutIsStrInference([d_1174_v17_])
                index172_ = default__.safeIndex(306, (d_1173_v16_).length(0))
                index173_ = default__.safeIndex(368, (d_1173_v16_).length(0))
                rhs180_ = d_1172_v15_
                rhs181_ = (d_1172_v15_ if d_1170_v13_ else d_1172_v15_)
                rhs182_ = default__.safeDivisionInt(len(d_1175_v18_), (-833) - (p0))
                rhs183_ = p1
                rhs184_ = p1
                lhs136_ = d_1173_v16_
                lhs137_ = default__.safeIndex(306, (d_1173_v16_).length(0))
                lhs138_ = d_1173_v16_
                lhs139_ = default__.safeIndex(368, (d_1173_v16_).length(0))
                lhs140_ = globalState
                d_1172_v15_ = rhs180_
                d_1172_v15_ = rhs181_
                lhs136_[lhs137_] = rhs182_
                lhs138_[lhs139_] = rhs183_
                lhs140_.f21 = rhs184_
                (globalState).f21 = p0
                d_1176_v19_: bool
                d_1177_v20_: bool
                d_1178_v21_: _dafny.Map
                d_1179_v22_: int
                out41_: bool
                out42_: bool
                out43_: _dafny.Map
                out44_: int
                out41_, out42_, out43_, out44_ = default__.m0(globalState)
                d_1176_v19_ = out41_
                d_1177_v20_ = out42_
                d_1178_v21_ = out43_
                d_1179_v22_ = out44_
                nw207_ = _dafny.Array(False, 21)
                d_1172_v15_ = nw207_
                d_1180_v23_: _dafny.Map
                d_1180_v23_ = _dafny.Map({d_1174_v17_: d_1176_v19_})
                d_1181_v24_: _dafny.MultiSet
                d_1181_v24_ = _dafny.MultiSet([False])
                d_1182_v25_: _dafny.Map
                d_1182_v25_ = _dafny.Map({(p0 if ((d_1180_v23_)[d_1174_v17_] if (d_1174_v17_) in (d_1180_v23_) else d_1170_v13_) else (d_1181_v24_).cardinality): (d_1177_v20_ if d_1177_v20_ else d_1177_v20_)})
                d_1182_v25_ = d_1182_v25_
            elif True:
                d_1183_v26_: _dafny.Seq
                d_1183_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otaieyd"))
                d_1184_v27_: _dafny.Array
                nw208_ = _dafny.Array(None, 7)
                nw208_[int(0)] = p0
                nw208_[int(1)] = len(d_1183_v26_)
                nw208_[int(2)] = (self).f38
                nw208_[int(3)] = 904
                nw208_[int(4)] = (self).f38
                nw208_[int(5)] = 40
                nw208_[int(6)] = p0
                d_1184_v27_ = nw208_
                d_1185_v28_: _dafny.Set
                d_1185_v28_ = _dafny.Set({d_1184_v27_})
                d_1186_v29_: _dafny.Array
                nw209_ = _dafny.Array(D9.default()(), 26)
                d_1186_v29_ = nw209_
                d_1187_v30_: _dafny.Array
                d_1187_v30_ = d_1186_v29_
                d_1188_v31_: _dafny.Map
                d_1188_v31_ = _dafny.Map({d_1187_v30_: p0})
                (globalState).f5 = default__.safeDivisionInt(len(d_1185_v28_), ((d_1188_v31_)[d_1187_v30_] if (d_1187_v30_) in (d_1188_v31_) else p1))
                d_1189_v32_: _dafny.Array
                nw210_ = _dafny.Array(_dafny.Seq({}), 18)
                d_1189_v32_ = nw210_
                d_1190_v33_: _dafny.Seq
                d_1190_v33_ = _dafny.SeqWithoutIsStrInference([(self).f38, (self).f38])
                index174_ = default__.safeIndex(700, (d_1189_v32_).length(0))
                (d_1189_v32_)[index174_] = d_1190_v33_
                d_1191_v34_: _dafny.Set
                d_1191_v34_ = _dafny.Set({d_1170_v13_})
                d_1192_v35_: _dafny.Set
                d_1192_v35_ = _dafny.Set({p1, (len(d_1191_v34_)) * (p0)})
                index175_ = default__.safeIndex(700, (d_1189_v32_).length(0))
                rhs185_ = d_1190_v33_
                rhs186_ = not (not(d_1170_v13_)) or (d_1170_v13_)
                rhs187_ = d_1192_v35_
                rhs188_ = (self).f38
                rhs189_ = (self).f38
                lhs141_ = d_1189_v32_
                lhs142_ = default__.safeIndex(700, (d_1189_v32_).length(0))
                lhs143_ = globalState
                lhs144_ = globalState
                lhs141_[lhs142_] = rhs185_
                r2 = rhs186_
                d_1192_v35_ = rhs187_
                lhs143_.f16 = rhs188_
                lhs144_.f16 = rhs189_
                d_1193_v36_: _dafny.Map
                d_1193_v36_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1170_v13_: d_1170_v13_}))]) for d_1194_i0_ in range(default__.abs(317))]): d_1170_v13_})
                d_1193_v36_ = (d_1193_v36_).set(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1 for d_1195_i1_ in range(default__.abs(209))])]), d_1170_v13_)
                d_1196_v37_: _dafny.Map
                d_1196_v37_ = _dafny.Map({d_1170_v13_: d_1170_v13_})
                d_1197_v38_: _dafny.Map
                d_1197_v38_ = _dafny.Map({len(d_1196_v37_): (0) - ((self).f38)})
                (globalState).f12 = ((d_1197_v38_).set((self).f38, 484)) | ((d_1197_v38_) | (d_1197_v38_))
                index176_ = default__.safeIndex(357, (d_1184_v27_).length(0))
                (d_1184_v27_)[index176_] = p1
                index177_ = default__.safeIndex(357, (d_1184_v27_).length(0))
                (d_1184_v27_)[index177_] = (0) - (-617)
            d_1198_v39_: _dafny.Array
            nw211_ = _dafny.Array(int(0), 25)
            d_1198_v39_ = nw211_
            d_1199_v40_: _dafny.MultiSet
            d_1199_v40_ = _dafny.MultiSet([d_1170_v13_])
            d_1200_v41_: _dafny.Map
            d_1200_v41_ = _dafny.Map({(self).f38: p1})
            d_1201_v42_: C3
            nw212_ = C3()
            nw212_.ctor__(d_1198_v39_, True, (_dafny.Map({p1: (d_1199_v40_).cardinality})) | (d_1200_v41_))
            d_1201_v42_ = nw212_
            r0 = d_1170_v13_
            d_1202_v43_: _dafny.Seq
            d_1202_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhlhhi"))
            d_1202_v43_ = d_1202_v43_
        (globalState).f16 = (-763) + (p0)
        d_1203_v44_: _dafny.Seq
        d_1203_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhcywahv"))
        (globalState).f1 = (d_1203_v44_) == (d_1203_v44_)
        d_1204_v45_: bool
        d_1204_v45_ = True
        d_1205_v46_: _dafny.Seq
        d_1205_v46_ = _dafny.SeqWithoutIsStrInference([d_1204_v45_, d_1204_v45_, d_1204_v45_])
        d_1206_v47_: str
        d_1206_v47_ = _dafny.CodePoint('w')
        d_1207_v48_: _dafny.Seq
        d_1207_v48_ = _dafny.SeqWithoutIsStrInference([len(((d_1203_v44_).set(default__.safeIndex(len(d_1205_v46_), len(d_1203_v44_)), d_1206_v47_)).set(default__.safeIndex(969, len((d_1203_v44_).set(default__.safeIndex(len(d_1205_v46_), len(d_1203_v44_)), d_1206_v47_))), d_1206_v47_))])
        d_1208_i2_: int
        d_1208_i2_ = 0
        with _dafny.label("6"):
            while (d_1205_v46_)[default__.safeIndex((d_1207_v48_)[default__.safeIndex((self).f38, len(d_1207_v48_))], len(d_1205_v46_))]:
                with _dafny.c_label("6"):
                    if (d_1208_i2_) >= (100):
                        raise _dafny.Break("6")
                    d_1208_i2_ = (d_1208_i2_) + (1)
                    (globalState).f13 = d_1204_v45_
                    d_1209_v49_: _dafny.Array
                    nw213_ = _dafny.Array(D9.default()(), 3)
                    d_1209_v49_ = nw213_
                    d_1210_v50_: _dafny.Seq
                    d_1210_v50_ = _dafny.SeqWithoutIsStrInference([d_1209_v49_, d_1209_v49_, d_1209_v49_])
                    d_1211_v51_: _dafny.Array
                    d_1211_v51_ = (d_1210_v50_)[default__.safeIndex((self).f38, len(d_1210_v50_))]
                    d_1211_v51_ = d_1211_v51_
                    rhs190_ = (d_1206_v47_) in ((_dafny.SeqWithoutIsStrInference([d_1206_v47_ for d_1212_i3_ in range(default__.abs(534))]) if d_1204_v45_ else _dafny.SeqWithoutIsStrInference([d_1206_v47_ for d_1213_i4_ in range(default__.abs(504))])))
                    rhs191_ = d_1204_v45_
                    rhs192_ = default__.fm0(p1, globalState)
                    lhs145_ = globalState
                    lhs146_ = globalState
                    r2 = rhs190_
                    lhs145_.f13 = rhs191_
                    lhs146_.f1 = rhs192_
                    d_1214_v52_: D5
                    d_1214_v52_ = D5_DC11(d_1204_v45_)
                    d_1215_v53_: _dafny.Map
                    d_1215_v53_ = _dafny.Map({d_1204_v45_: (d_1214_v52_).cf16})
                    (globalState).f13 = not (not(((d_1215_v53_)[d_1204_v45_] if (d_1204_v45_) in (d_1215_v53_) else d_1204_v45_))) or (d_1204_v45_)
                    pass
            pass
        d_1216_v54_: _dafny.Array
        nw214_ = _dafny.Array(False, 29)
        d_1216_v54_ = nw214_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1216_v54_).length(0)):
            d_1217_i5_: int = guard_loop_5_
            if (True) and (((0) <= (d_1217_i5_)) and ((d_1217_i5_) < ((d_1216_v54_).length(0)))):
                (d_1216_v54_)[(d_1217_i5_)] = (_dafny.CodePoint('s')) in (d_1203_v44_)
        d_1218_v55_: _dafny.MultiSet
        d_1218_v55_ = _dafny.MultiSet([d_1203_v44_, _dafny.SeqWithoutIsStrInference([d_1206_v47_])])
        d_1219_v56_: _dafny.Map
        d_1219_v56_ = _dafny.Map({(d_1218_v55_).cardinality: default__.fm0(p1, globalState)})
        d_1220_v57_: _dafny.Map
        d_1220_v57_ = _dafny.Map({(self).f38: len(d_1219_v56_)})
        (globalState).f12 = ((d_1220_v57_) | (default__.fm45(globalState))) | (d_1220_v57_)
        r0 = not (d_1204_v45_) or (not(d_1204_v45_))
        d_1221_v58_: D5
        d_1221_v58_ = D5_DC13(p1)
        r1 = d_1221_v58_
        r2 = d_1204_v45_
        return r0, r1, r2

    def m16(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: _dafny.Array = _dafny.Array(None, 0)
        d_1222_v0_: _dafny.Array
        nw215_ = _dafny.Array(_dafny.Array(None, 0), 23)
        d_1222_v0_ = nw215_
        d_1223_v1_: _dafny.Array
        nw216_ = _dafny.Array(False, 24)
        d_1223_v1_ = nw216_
        index178_ = default__.safeIndex(889, (d_1222_v0_).length(0))
        (d_1222_v0_)[index178_] = d_1223_v1_
        index179_ = default__.safeIndex(889, (d_1222_v0_).length(0))
        (d_1222_v0_)[index179_] = d_1223_v1_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1223_v1_).length(0)):
            d_1224_i0_: int = guard_loop_6_
            if (True) and (((0) <= (d_1224_i0_)) and ((d_1224_i0_) < ((d_1223_v1_).length(0)))):
                (d_1223_v1_)[(d_1224_i0_)] = p1
        if p1:
            d_1225_v2_: _dafny.Seq
            d_1225_v2_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1226_v3_: D3
            d_1226_v3_ = D3_DC7(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1227_i1_ in range(default__.abs(928))]), (0) - (p0), p0, default__.fm2(globalState), d_1225_v2_)
            (globalState).f21 = (d_1226_v3_).cf9
            if p1:
                (globalState).f0 = (self).f38
                d_1228_v4_: C4
                nw217_ = C4()
                nw217_.ctor__()
                d_1228_v4_ = nw217_
                (globalState).f22 = ((p0 if p1 else -727)) + (-934)
                d_1223_v1_ = d_1223_v1_
                d_1229_v5_: _dafny.Seq
                d_1229_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iergc"))
                d_1229_v5_ = d_1229_v5_
            elif True:
                d_1230_v6_: _dafny.Array
                def lambda61_(d_1231_i2_):
                    return (d_1231_i2_) + ((self).f38)

                init35_ = lambda61_
                nw218_ = _dafny.Array(None, 15)
                for i0_35_ in range(nw218_.length(0)):
                    nw218_[i0_35_] = init35_(i0_35_)
                d_1230_v6_ = nw218_
                d_1232_v7_: D5
                d_1232_v7_ = D5_DC13(p0)
                d_1233_v8_: _dafny.Map
                d_1233_v8_ = _dafny.Map({(self).f38: d_1232_v7_})
                d_1234_v9_: _dafny.Map
                d_1234_v9_ = _dafny.Map({len(d_1233_v8_): p1})
                d_1235_v10_: _dafny.Seq
                d_1235_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvogptsi"))
                d_1236_v11_: _dafny.Seq
                d_1236_v11_ = _dafny.SeqWithoutIsStrInference([d_1235_v10_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlol"))])
                d_1237_v12_: _dafny.Map
                d_1237_v12_ = _dafny.Map({len(d_1234_v9_): _dafny.MultiSet(d_1236_v11_)})
                d_1238_v13_: _dafny.MultiSet
                d_1238_v13_ = _dafny.MultiSet([d_1235_v10_])
                index180_ = default__.safeIndex(558, (d_1230_v6_).length(0))
                (d_1230_v6_)[index180_] = (((d_1237_v12_)[(self).f38] if ((self).f38) in (d_1237_v12_) else d_1238_v13_)).cardinality
                d_1239_v14_: str
                d_1239_v14_ = _dafny.CodePoint('l')
                d_1240_v15_: _dafny.Seq
                d_1240_v15_ = _dafny.SeqWithoutIsStrInference([p0])
                index181_ = default__.safeIndex(889, (d_1223_v1_).length(0))
                (d_1223_v1_)[index181_] = (len(default__.fm21(d_1239_v14_, p0, d_1226_v3_, globalState))) < (len(d_1240_v15_))
                index182_ = default__.safeIndex(2, (d_1230_v6_).length(0))
                (d_1230_v6_)[index182_] = 615
                index183_ = default__.safeIndex(632, (d_1230_v6_).length(0))
                (d_1230_v6_)[index183_] = (self).f38
                d_1241_v16_: _dafny.Set
                d_1241_v16_ = _dafny.Set({p0, len(d_1225_v2_), -627, default__.fm2(globalState)})
                d_1242_v17_: D7
                d_1242_v17_ = D7_DC16((self).f38, not(p1))
                index184_ = default__.safeIndex(558, (d_1230_v6_).length(0))
                index185_ = default__.safeIndex(889, (d_1223_v1_).length(0))
                index186_ = default__.safeIndex(2, (d_1230_v6_).length(0))
                index187_ = default__.safeIndex(632, (d_1230_v6_).length(0))
                rhs193_ = (d_1241_v16_).ispropersubset(d_1241_v16_)
                rhs194_ = p0
                rhs195_ = p1
                rhs196_ = (d_1242_v17_).cf22
                rhs197_ = ((self).f38) * (592)
                lhs147_ = globalState
                lhs148_ = d_1230_v6_
                lhs149_ = default__.safeIndex(558, (d_1230_v6_).length(0))
                lhs150_ = d_1223_v1_
                lhs151_ = default__.safeIndex(889, (d_1223_v1_).length(0))
                lhs152_ = d_1230_v6_
                lhs153_ = default__.safeIndex(2, (d_1230_v6_).length(0))
                lhs154_ = d_1230_v6_
                lhs155_ = default__.safeIndex(632, (d_1230_v6_).length(0))
                lhs147_.f1 = rhs193_
                lhs148_[lhs149_] = rhs194_
                lhs150_[lhs151_] = rhs195_
                lhs152_[lhs153_] = rhs196_
                lhs154_[lhs155_] = rhs197_
                d_1243_v18_: _dafny.MultiSet
                d_1243_v18_ = _dafny.MultiSet([d_1239_v14_])
                rhs198_ = (d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]
                rhs199_ = (((d_1243_v18_)[_dafny.CodePoint('d')] if (_dafny.CodePoint('d')) in (d_1243_v18_) else len(d_1241_v16_))) * (p0)
                lhs156_ = globalState
                r3 = rhs198_
                lhs156_.f2 = rhs199_
                d_1244_v19_: _dafny.Seq
                d_1244_v19_ = _dafny.SeqWithoutIsStrInference([d_1240_v15_])
                d_1245_v20_: _dafny.MultiSet
                d_1245_v20_ = _dafny.MultiSet([d_1240_v15_, (d_1244_v19_)[default__.safeIndex((d_1230_v6_)[default__.safeIndex(558, (d_1230_v6_).length(0))], len(d_1244_v19_))]])
                def iife138_():
                    coll58_ = _dafny.Set()
                    compr_58_: _dafny.Seq
                    for compr_58_ in (d_1245_v20_).Elements:
                        d_1246_v21_: _dafny.Seq = compr_58_
                        if (d_1246_v21_) in (d_1245_v20_):
                            coll58_ = coll58_.union(_dafny.Set([d_1246_v21_]))
                    return _dafny.Set(coll58_)
                (globalState).f2 = len(iife138_()
                )
                rhs200_ = (d_1223_v1_)[default__.safeIndex(889, (d_1223_v1_).length(0))]
                rhs201_ = d_1235_v10_
                rhs202_ = p1
                rhs203_ = d_1239_v14_
                rhs204_ = len(d_1235_v10_)
                lhs157_ = globalState
                lhs158_ = globalState
                lhs159_ = globalState
                lhs157_.f13 = rhs200_
                d_1235_v10_ = rhs201_
                lhs158_.f13 = rhs202_
                d_1239_v14_ = rhs203_
                lhs159_.f21 = rhs204_
                d_1247_v22_: _dafny.Map
                d_1247_v22_ = _dafny.Map({p1: (self).f38})
                d_1248_v23_: C1
                nw219_ = C1()
                nw219_.ctor__(d_1247_v22_)
                d_1248_v23_ = nw219_
            d_1249_v24_: _dafny.Seq
            d_1249_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcgk"))
            (globalState).f1 = (default__.fm2(globalState)) != (default__.safeDivisionInt((self).f38, len(d_1249_v24_)))
            d_1250_v25_: _dafny.Map
            d_1250_v25_ = _dafny.Map({default__.fm34((d_1225_v2_).set(default__.safeIndex((self).f38, len(d_1225_v2_)), p1), p0, globalState): p0})
            d_1251_v26_: _dafny.Map
            d_1251_v26_ = (_dafny.Map({default__.fm34(_dafny.SeqWithoutIsStrInference([p1, default__.fm0(p0, globalState)]), default__.fm2(globalState), globalState): (self).f38})) | (d_1250_v25_)
            source16_ = d_1251_v26_
            d_1252___mcc_h0_ = source16_
            d_1253_cf48_ = d_1252___mcc_h0_
            d_1254_v27_: _dafny.Array
            nw220_ = _dafny.Array(_dafny.CodePoint('D'), 11)
            d_1254_v27_ = nw220_
            d_1255_v28_: D0
            d_1255_v28_ = D0_DC0(d_1254_v27_)
            d_1256_v29_: _dafny.Map
            d_1256_v29_ = _dafny.Map({D0_DC2(d_1255_v28_): p0})
            d_1257_v30_: _dafny.Seq
            d_1257_v30_ = _dafny.SeqWithoutIsStrInference([len(d_1249_v24_), (0) - (len(d_1256_v29_)), default__.fm2(globalState), len(d_1249_v24_)])
            d_1258_v31_: _dafny.MultiSet
            d_1258_v31_ = _dafny.MultiSet([d_1257_v30_, d_1257_v30_])
            d_1259_v32_: _dafny.Seq
            d_1259_v32_ = _dafny.SeqWithoutIsStrInference([((d_1258_v31_)[d_1257_v30_] if (d_1257_v30_) in (d_1258_v31_) else (self).f38)])
            d_1260_v33_: _dafny.Map
            d_1260_v33_ = _dafny.Map({d_1259_v32_: d_1249_v24_})
            rhs205_ = p1
            rhs206_ = (((d_1249_v24_) + (d_1249_v24_)) + (d_1249_v24_)).set(default__.safeIndex((self).f38, len(((d_1249_v24_) + (d_1249_v24_)) + (d_1249_v24_))), _dafny.CodePoint('d'))
            rhs207_ = ((d_1260_v33_)[d_1257_v30_] if (d_1257_v30_) in (d_1260_v33_) else d_1249_v24_)
            lhs160_ = globalState
            lhs160_.f13 = rhs205_
            d_1249_v24_ = rhs206_
            d_1249_v24_ = rhs207_
            d_1261_v34_: C4
            nw221_ = C4()
            nw221_.ctor__()
            d_1261_v34_ = nw221_
            d_1262_v35_: _dafny.Map
            d_1262_v35_ = _dafny.Map({(p0) <= (p0): d_1261_v34_})
            d_1262_v35_ = (d_1262_v35_).set(p1, d_1261_v34_)
            d_1263_v36_: _dafny.Array
            def lambda62_(d_1264_p0_):
                def lambda63_(d_1265_i3_):
                    return default__.safeDivisionInt(d_1265_i3_, d_1264_p0_)

                return lambda63_

            init36_ = lambda62_(p0)
            nw222_ = _dafny.Array(None, 12)
            for i0_36_ in range(nw222_.length(0)):
                nw222_[i0_36_] = init36_(i0_36_)
            d_1263_v36_ = nw222_
            d_1266_v37_: str
            d_1266_v37_ = _dafny.CodePoint('c')
            d_1263_v36_ = (D0_DC1(d_1263_v36_, d_1266_v37_, p0)).cf1
            (globalState).f16 = (self).f38
            d_1267_v38_: D12
            d_1267_v38_ = D12_DC29()
            source17_ = d_1267_v38_
            if source17_.is_DC29:
                (globalState).f1 = not(p1)
                (globalState).f1 = (p1 if p1 else (856) in ((_dafny.Map({(self).f38: True})).set(len(d_1249_v24_), p1)))
                d_1268_v39_: _dafny.Array
                nw223_ = _dafny.Array(int(0), 14)
                d_1268_v39_ = nw223_
                index188_ = default__.safeIndex(684, (d_1268_v39_).length(0))
                (d_1268_v39_)[index188_] = (self).f38
                index189_ = default__.safeIndex(684, (d_1268_v39_).length(0))
                (d_1268_v39_)[index189_] = default__.fm2(globalState)
                index190_ = default__.safeIndex(684, (d_1268_v39_).length(0))
                (d_1268_v39_)[index190_] = (582) + (p0)
            elif source17_.is_DC28:
                d_1269___mcc_h1_ = source17_.cf43
                d_1270_cf43_ = d_1269___mcc_h1_
                (globalState).f1 = p1
                d_1271_v40_: _dafny.Array
                nw224_ = _dafny.Array(int(0), 15)
                d_1271_v40_ = nw224_
                index191_ = default__.safeIndex(782, (d_1271_v40_).length(0))
                (d_1271_v40_)[index191_] = 945
                index192_ = default__.safeIndex(782, (d_1271_v40_).length(0))
                (d_1271_v40_)[index192_] = len(_dafny.Map({(default__.fm28(p1, False, p0, globalState)) + (d_1225_v2_): p1}))
                d_1272_v41_: _dafny.Map
                d_1272_v41_ = _dafny.Map({p1: p1})
                d_1273_v42_: _dafny.Seq
                d_1273_v42_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p1: p1}), d_1272_v41_, d_1272_v41_])
                d_1274_v43_: _dafny.Map
                d_1274_v43_ = _dafny.Map({p1: (d_1273_v42_)[default__.safeIndex(p0, len(d_1273_v42_))]})
                d_1274_v43_ = d_1274_v43_
                (globalState).f5 = default__.safeModuloInt(default__.fm2(globalState), (self).f38)
            elif True:
                d_1275___mcc_h2_ = source17_.cf44
                d_1276_cf44_ = d_1275___mcc_h2_
                d_1277_v44_: str
                d_1277_v44_ = _dafny.CodePoint('m')
                d_1278_v45_: _dafny.Array
                nw225_ = _dafny.Array(None, 18)
                nw225_[int(0)] = d_1249_v24_
                nw225_[int(1)] = d_1249_v24_
                nw225_[int(2)] = d_1249_v24_
                nw225_[int(3)] = d_1249_v24_
                nw225_[int(4)] = (d_1249_v24_) + (d_1249_v24_)
                nw225_[int(5)] = d_1249_v24_
                nw225_[int(6)] = d_1249_v24_
                nw225_[int(7)] = d_1249_v24_
                nw225_[int(8)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1279_i4_ in range(default__.abs(896))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drtc")))
                nw225_[int(9)] = d_1249_v24_
                nw225_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwtfq"))) + (d_1249_v24_)
                nw225_[int(11)] = d_1249_v24_
                nw225_[int(12)] = d_1249_v24_
                nw225_[int(13)] = d_1249_v24_
                nw225_[int(14)] = d_1249_v24_
                nw225_[int(15)] = default__.fm30(p0, (d_1226_v3_).cf8, (self).f38, globalState)
                nw225_[int(16)] = ((d_1249_v24_) + (default__.fm30((self).f38, d_1249_v24_, (0) - ((self).f38), globalState))).set(default__.safeIndex(p0, len((d_1249_v24_) + (default__.fm30((self).f38, d_1249_v24_, (0) - ((self).f38), globalState)))), d_1277_v44_)
                nw225_[int(17)] = default__.fm30(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcpv"))), _dafny.SeqWithoutIsStrInference([d_1277_v44_ for d_1280_i5_ in range(default__.abs(593))]), 132, globalState)
                d_1278_v45_ = nw225_
                index193_ = default__.safeIndex(319, (d_1278_v45_).length(0))
                (d_1278_v45_)[index193_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbrp"))) + (d_1249_v24_)
                d_1281_v46_: _dafny.Array
                nw226_ = _dafny.Array(int(0), 2)
                d_1281_v46_ = nw226_
                d_1282_v47_: _dafny.Set
                d_1282_v47_ = _dafny.Set({p1})
                d_1283_v48_: _dafny.Map
                d_1283_v48_ = _dafny.Map({len(d_1282_v47_): (0) - (p0)})
                d_1284_v49_: T0
                nw227_ = C3()
                nw227_.ctor__(d_1281_v46_, False, d_1283_v48_)
                d_1284_v49_ = nw227_
                index194_ = default__.safeIndex(319, (d_1278_v45_).length(0))
                rhs208_ = (d_1249_v24_) + (d_1249_v24_)
                rhs209_ = d_1284_v49_
                rhs210_ = d_1249_v24_
                lhs161_ = d_1278_v45_
                lhs162_ = default__.safeIndex(319, (d_1278_v45_).length(0))
                lhs161_[lhs162_] = rhs208_
                d_1284_v49_ = rhs209_
                d_1249_v24_ = rhs210_
                d_1285_v50_: _dafny.Map
                d_1285_v50_ = _dafny.Map({p1: (self).f38})
                d_1286_v51_: _dafny.Seq
                d_1286_v51_ = _dafny.SeqWithoutIsStrInference([(self).f38, p0])
                d_1285_v50_ = (d_1285_v50_).set((_dafny.SeqWithoutIsStrInference([(self).f38, p0])) != (d_1286_v51_), len(_dafny.SeqWithoutIsStrInference([d_1277_v44_ for d_1287_i6_ in range(default__.abs(830))])))
                d_1288_v52_: _dafny.MultiSet
                d_1288_v52_ = _dafny.MultiSet([(self).f38, p0, p0, p0])
                d_1289_v53_: _dafny.Seq
                d_1289_v53_ = _dafny.SeqWithoutIsStrInference([(d_1288_v52_ if True else d_1288_v52_)])
                d_1289_v53_ = d_1289_v53_
                d_1290_v54_: D12
                d_1290_v54_ = D12_DC30(D12_DC29())
                d_1290_v54_ = d_1290_v54_
        elif True:
            if p1:
                r2 = (0) - (p0)
                (globalState).f13 = (p1) or (p1)
                d_1291_v55_: _dafny.Array
                nw228_ = _dafny.Array(int(0), 17)
                d_1291_v55_ = nw228_
                d_1292_v56_: str
                d_1292_v56_ = _dafny.CodePoint('c')
                d_1293_v57_: _dafny.MultiSet
                d_1293_v57_ = _dafny.MultiSet([p0, 704])
                d_1294_v58_: D0
                d_1294_v58_ = D0_DC1(d_1291_v55_, d_1292_v56_, ((d_1293_v57_)[(self).f38] if ((self).f38) in (d_1293_v57_) else (self).f38))
                d_1294_v58_ = D0_DC1(d_1291_v55_, _dafny.CodePoint('e'), p0)
                index195_ = default__.safeIndex(778, (d_1291_v55_).length(0))
                (d_1291_v55_)[index195_] = (self).f38
                d_1295_v60_: _dafny.Array
                def lambda64_(d_1296_p1_, d_1297_p0_):
                    def lambda65_(d_1298_i7_):
                        def iife139_():
                            coll59_ = _dafny.Map()
                            compr_59_: int
                            for compr_59_ in (_dafny.Set({d_1297_p0_})).Elements:
                                d_1299_v59_: int = compr_59_
                                if (d_1299_v59_) in (_dafny.Set({d_1297_p0_})):
                                    coll59_[default__.safeDivisionInt(d_1299_v59_, (D7_DC17(True, d_1296_p1_, 196)).cf26)] = d_1296_p1_
                            return _dafny.Map(coll59_)
                        return iife139_()
                        

                    return lambda65_

                init37_ = lambda64_(p1, p0)
                nw229_ = _dafny.Array(None, 14)
                for i0_37_ in range(nw229_.length(0)):
                    nw229_[i0_37_] = init37_(i0_37_)
                d_1295_v60_ = nw229_
                d_1300_v61_: _dafny.Map
                d_1300_v61_ = _dafny.Map({(self).f38: default__.fm0((self).f38, globalState)})
                index196_ = default__.safeIndex(908, (d_1295_v60_).length(0))
                (d_1295_v60_)[index196_] = d_1300_v61_
                index197_ = default__.safeIndex(778, (d_1291_v55_).length(0))
                index198_ = default__.safeIndex(908, (d_1295_v60_).length(0))
                rhs211_ = (self).f38
                rhs212_ = d_1300_v61_
                lhs163_ = d_1291_v55_
                lhs164_ = default__.safeIndex(778, (d_1291_v55_).length(0))
                lhs165_ = d_1295_v60_
                lhs166_ = default__.safeIndex(908, (d_1295_v60_).length(0))
                lhs163_[lhs164_] = rhs211_
                lhs165_[lhs166_] = rhs212_
                d_1301_v62_: _dafny.Seq
                d_1301_v62_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.safeModuloInt((d_1291_v55_)[default__.safeIndex(778, (d_1291_v55_).length(0))], (d_1291_v55_)[default__.safeIndex(778, (d_1291_v55_).length(0))])), (d_1291_v55_)[default__.safeIndex(778, (d_1291_v55_).length(0))], (self).f38, -504, (d_1291_v55_)[default__.safeIndex(778, (d_1291_v55_).length(0))]])
                d_1301_v62_ = d_1301_v62_
            elif True:
                d_1302_v63_: _dafny.Array
                def lambda66_(d_1303_p0_, d_1304_p1_):
                    def lambda67_(d_1305_i8_):
                        return (_dafny.Map({d_1303_p0_: d_1304_p1_})) | (_dafny.Map({d_1303_p0_: True}))

                    return lambda67_

                init38_ = lambda66_(p0, p1)
                nw230_ = _dafny.Array(None, 27)
                for i0_38_ in range(nw230_.length(0)):
                    nw230_[i0_38_] = init38_(i0_38_)
                d_1302_v63_ = nw230_
                d_1306_v64_: _dafny.Map
                d_1306_v64_ = _dafny.Map({p0: not(p1)})
                index199_ = default__.safeIndex(927, (d_1302_v63_).length(0))
                (d_1302_v63_)[index199_] = d_1306_v64_
                index200_ = default__.safeIndex(927, (d_1302_v63_).length(0))
                (d_1302_v63_)[index200_] = d_1306_v64_
                d_1307_v65_: _dafny.Array
                nw231_ = _dafny.Array(int(0), 11)
                d_1307_v65_ = nw231_
                d_1308_v66_: _dafny.Seq
                d_1308_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukmyudn"))
                d_1309_v67_: _dafny.Map
                d_1309_v67_ = _dafny.Map({len(d_1308_v66_): (0) - ((self).f38)})
                d_1310_v68_: C3
                nw232_ = C3()
                nw232_.ctor__(d_1307_v65_, (p1 if p1 else not(p1)), (d_1309_v67_).set((0) - ((self).f38), p0))
                d_1310_v68_ = nw232_
                d_1311_v69_: str
                d_1311_v69_ = _dafny.CodePoint('c')
                d_1312_v70_: _dafny.Seq
                d_1312_v70_ = _dafny.SeqWithoutIsStrInference([p1, ((d_1308_v66_).set(default__.safeIndex(-351, len(d_1308_v66_)), d_1311_v69_)) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psyfxcw")))])
                (globalState).f13 = (d_1312_v70_)[default__.safeIndex(len(_dafny.Set({d_1307_v65_, (d_1310_v68_).f39, d_1307_v65_})), len(d_1312_v70_))]
                d_1313_v71_: _dafny.MultiSet
                d_1313_v71_ = _dafny.MultiSet([not(default__.fm0(335, globalState)), d_1310_v68_.f40, p1, p1])
                d_1313_v71_ = _dafny.MultiSet(d_1312_v70_)
                d_1306_v64_ = (d_1306_v64_).set(p0, d_1310_v68_.f40)
            d_1314_v72_: _dafny.Array
            nw233_ = _dafny.Array(None, 7)
            nw233_[int(0)] = p0
            nw233_[int(1)] = p0
            nw233_[int(2)] = p0
            nw233_[int(3)] = (self).f38
            nw233_[int(4)] = p0
            nw233_[int(5)] = (self).f38
            nw233_[int(6)] = (self).f38
            d_1314_v72_ = nw233_
            d_1315_v73_: D0
            d_1315_v73_ = D0_DC1(d_1314_v72_, _dafny.CodePoint('u'), 489)
            d_1316_v74_: D0
            d_1316_v74_ = D0_DC2(d_1315_v73_)
            d_1317_v75_: D0
            d_1317_v75_ = D0_DC2(d_1316_v74_)
            d_1318_v76_: D0
            d_1318_v76_ = D0_DC2(d_1316_v74_)
            d_1319_v77_: D0
            d_1319_v77_ = D0_DC2(d_1318_v76_)
            d_1320_v78_: _dafny.Map
            d_1320_v78_ = _dafny.Map({p1: d_1319_v77_})
            d_1321_v79_: _dafny.Set
            d_1321_v79_ = _dafny.Set({d_1320_v78_})
            d_1321_v79_ = d_1321_v79_
            index201_ = default__.safeIndex(889, (d_1222_v0_).length(0))
            (d_1222_v0_)[index201_] = (d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]
            d_1322_v80_: _dafny.MultiSet
            d_1322_v80_ = _dafny.MultiSet([not(p1)])
            d_1323_v81_: _dafny.Map
            d_1323_v81_ = _dafny.Map({d_1322_v80_: p0})
            d_1323_v81_ = (d_1323_v81_).set(d_1322_v80_, p0)
            r2 = default__.fm2(globalState)
        d_1324_v82_: str
        d_1324_v82_ = _dafny.CodePoint('a')
        d_1324_v82_ = _dafny.CodePoint('t')
        d_1325_v83_: _dafny.Array
        def lambda68_(d_1326_i10_):
            return default__.safeModuloInt(d_1326_i10_, (self).f38)

        init39_ = lambda68_
        nw234_ = _dafny.Array(None, 24)
        for i0_39_ in range(nw234_.length(0)):
            nw234_[i0_39_] = init39_(i0_39_)
        d_1325_v83_ = nw234_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1325_v83_).length(0)):
            d_1327_i9_: int = guard_loop_7_
            if (True) and (((0) <= (d_1327_i9_)) and ((d_1327_i9_) < ((d_1325_v83_).length(0)))):
                (d_1325_v83_)[(d_1327_i9_)] = (d_1327_i9_) * ((p0) * (1))
        if p1:
            d_1328_v84_: D8
            d_1328_v84_ = D8_DC20(p1)
            d_1329_v85_: D7
            d_1329_v85_ = D7_DC16((self).f38, (D7_DC17(p1, (d_1328_v84_).cf31, (self).f38)).cf25)
            d_1330_v87_: _dafny.Map
            def iife140_():
                coll60_ = _dafny.Map()
                compr_60_: int
                for compr_60_ in _dafny.IntegerRange(363, 620):
                    d_1331_v86_: int = compr_60_
                    if ((363) <= (d_1331_v86_)) and ((d_1331_v86_) < (620)):
                        coll60_[(d_1331_v86_) + (p0)] = (self).f38
                return _dafny.Map(coll60_)
            d_1330_v87_ = _dafny.Map({d_1325_v83_: iife140_()
            })
            d_1332_v88_: _dafny.Map
            d_1332_v88_ = _dafny.Map({p0: (0) - (default__.fm2(globalState))})
            rhs213_ = (d_1329_v85_).cf22
            rhs214_ = ((len(((d_1330_v87_)[d_1325_v83_] if (d_1325_v83_) in (d_1330_v87_) else d_1332_v88_))) < (p0)) or (p1)
            lhs167_ = globalState
            r2 = rhs213_
            lhs167_.f1 = rhs214_
            d_1333_v89_: _dafny.Array
            nw235_ = _dafny.Array(None, 5)
            nw235_[int(0)] = d_1324_v82_
            nw235_[int(1)] = d_1324_v82_
            nw235_[int(2)] = d_1324_v82_
            nw235_[int(3)] = d_1324_v82_
            nw235_[int(4)] = d_1324_v82_
            d_1333_v89_ = nw235_
            d_1333_v89_ = d_1333_v89_
            (globalState).f21 = p0
            (globalState).f1 = False
            d_1334_v90_: D8
            d_1334_v90_ = D8_DC19(p1, False, p1)
            source18_ = d_1334_v90_
            if source18_.is_DC19:
                d_1335___mcc_h3_ = source18_.cf28
                d_1336___mcc_h4_ = source18_.cf29
                d_1337___mcc_h5_ = source18_.cf30
                d_1338_cf30_ = d_1337___mcc_h5_
                d_1339_cf29_ = d_1336___mcc_h4_
                d_1340_cf28_ = d_1335___mcc_h3_
                d_1341_v91_: _dafny.Seq
                d_1341_v91_ = _dafny.SeqWithoutIsStrInference([d_1339_cf29_, d_1338_cf30_])
                d_1342_v92_: D3
                d_1342_v92_ = D3_DC7(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1343_i11_ in range(default__.abs(33))]), (self).f38, p0, (self).f38, d_1341_v91_)
                d_1344_v93_: _dafny.Seq
                d_1344_v93_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f38: d_1338_cf30_})])
                d_1345_v94_: _dafny.Map
                d_1345_v94_ = _dafny.Map({d_1223_v1_: (self).f38})
                d_1346_v95_: _dafny.MultiSet
                d_1346_v95_ = _dafny.MultiSet([default__.fm2(globalState), (d_1342_v92_).cf9, (len((d_1344_v93_)[default__.safeIndex(default__.fm2(globalState), len(d_1344_v93_))]) if True else p0), len(d_1345_v94_)])
                d_1346_v95_ = d_1346_v95_
                d_1347_v96_: _dafny.Map
                d_1347_v96_ = _dafny.Map({d_1339_cf29_: p0})
                d_1348_v97_: C1
                nw236_ = C1()
                nw236_.ctor__((d_1347_v96_ if (d_1328_v84_).cf31 else _dafny.Map({True: (self).f38})))
                d_1348_v97_ = nw236_
                d_1346_v95_ = d_1346_v95_
                (globalState).f0 = (0) - ((self).f38)
            elif source18_.is_DC20:
                d_1349___mcc_h6_ = source18_.cf31
                d_1350_cf31_ = d_1349___mcc_h6_
                d_1351_v98_: _dafny.Seq
                d_1351_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gan"))
                d_1352_v99_: _dafny.MultiSet
                d_1352_v99_ = _dafny.MultiSet([d_1351_v98_])
                d_1353_v100_: _dafny.MultiSet
                d_1353_v100_ = d_1352_v99_
                d_1354_v101_: _dafny.MultiSet
                d_1354_v101_ = _dafny.MultiSet([p0])
                arr0_ = (d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]
                index202_ = default__.safeIndex(483, ((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]).length(0))
                arr0_[index202_] = not(((d_1354_v101_).cardinality) <= (p0))
                arr1_ = (d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]
                index203_ = default__.safeIndex(483, ((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]).length(0))
                rhs215_ = d_1353_v100_
                rhs216_ = p0
                rhs217_ = default__.fm0(len(default__.fm30((self).f38, d_1351_v98_, (self).f38, globalState)), globalState)
                rhs218_ = d_1350_cf31_
                lhs168_ = globalState
                lhs169_ = (d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]
                lhs170_ = default__.safeIndex(483, ((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]).length(0))
                d_1353_v100_ = rhs215_
                lhs168_.f16 = rhs216_
                d_1350_cf31_ = rhs217_
                lhs169_[lhs170_] = rhs218_
                nw237_ = _dafny.Array(int(0), 14)
                d_1325_v83_ = nw237_
                d_1355_v102_: _dafny.Seq
                d_1355_v102_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1356_v103_: _dafny.Map
                d_1356_v103_ = _dafny.Map({d_1223_v1_: d_1324_v82_})
                d_1357_v104_: _dafny.Seq
                d_1357_v104_ = _dafny.SeqWithoutIsStrInference([d_1223_v1_])
                d_1358_v105_: _dafny.Set
                d_1358_v105_ = _dafny.Set({True, p1})
                d_1359_v106_: _dafny.Array
                nw238_ = _dafny.Array(None, 19)
                nw238_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1324_v82_ for d_1360_i12_ in range(default__.abs(789))])
                nw238_[int(1)] = d_1351_v98_
                nw238_[int(2)] = (d_1351_v98_) + (default__.fm30(p0, d_1351_v98_, (0) - (p0), globalState))
                nw238_[int(3)] = d_1351_v98_
                nw238_[int(4)] = d_1351_v98_
                nw238_[int(5)] = d_1351_v98_
                nw238_[int(6)] = d_1351_v98_
                nw238_[int(7)] = d_1351_v98_
                nw238_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwpa"))
                nw238_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtmdl"))) + (d_1351_v98_)
                nw238_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjehcqr"))
                nw238_[int(11)] = d_1351_v98_
                nw238_[int(12)] = _dafny.SeqWithoutIsStrInference([d_1324_v82_ for d_1361_i13_ in range(default__.abs(900))])
                nw238_[int(13)] = (d_1351_v98_) + (d_1351_v98_)
                nw238_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybwgl"))
                nw238_[int(15)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlpsyhu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whtao")))
                nw238_[int(16)] = d_1351_v98_
                nw238_[int(17)] = (d_1351_v98_).set(default__.safeIndex((0) - (len(d_1355_v102_)), len(d_1351_v98_)), ((d_1356_v103_)[(d_1357_v104_)[default__.safeIndex(len(d_1358_v105_), len(d_1357_v104_))]] if ((d_1357_v104_)[default__.safeIndex(len(d_1358_v105_), len(d_1357_v104_))]) in (d_1356_v103_) else d_1324_v82_))
                nw238_[int(18)] = d_1351_v98_
                d_1359_v106_ = nw238_
                d_1362_v107_: _dafny.Seq
                d_1362_v107_ = _dafny.SeqWithoutIsStrInference([d_1350_cf31_, ((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))])[default__.safeIndex(483, ((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]).length(0))], ((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))])[default__.safeIndex(483, ((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]).length(0))], d_1350_cf31_, True])
                d_1363_v108_: D3
                d_1363_v108_ = D3_DC7(d_1351_v98_, (self).f38, p0, p0, d_1362_v107_)
                index204_ = default__.safeIndex(139, (d_1359_v106_).length(0))
                (d_1359_v106_)[index204_] = (d_1363_v108_).cf8
                index205_ = default__.safeIndex(139, (d_1359_v106_).length(0))
                (d_1359_v106_)[index205_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wyafye"))) + ((d_1351_v98_).set(default__.safeIndex((self).f38, len(d_1351_v98_)), d_1324_v82_))
                d_1364_v109_: _dafny.Map
                d_1364_v109_ = _dafny.Map({((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))])[default__.safeIndex(483, ((d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]).length(0))]: (d_1362_v107_)[default__.safeIndex((self).f38, len(d_1362_v107_))]})
                index206_ = default__.safeIndex(217, (d_1325_v83_).length(0))
                (d_1325_v83_)[index206_] = (self).f38
                d_1365_v110_: _dafny.Set
                d_1365_v110_ = _dafny.Set({d_1325_v83_})
                index207_ = default__.safeIndex(217, (d_1325_v83_).length(0))
                rhs219_ = d_1364_v109_
                rhs220_ = d_1362_v107_
                rhs221_ = (_dafny.Set({d_1325_v83_, d_1325_v83_})) == (d_1365_v110_)
                rhs222_ = default__.fm2(globalState)
                rhs223_ = d_1352_v99_
                lhs171_ = d_1325_v83_
                lhs172_ = default__.safeIndex(217, (d_1325_v83_).length(0))
                d_1364_v109_ = rhs219_
                d_1362_v107_ = rhs220_
                r0 = rhs221_
                lhs171_[lhs172_] = rhs222_
                d_1352_v99_ = rhs223_
            elif source18_.is_DC18:
                d_1366___mcc_h7_ = source18_.cf27
                d_1367_cf27_ = d_1366___mcc_h7_
                d_1368_v111_: _dafny.Array
                def lambda69_(d_1369_i14_):
                    return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))])) + (_dafny.SeqWithoutIsStrInference([534]))

                init40_ = lambda69_
                nw239_ = _dafny.Array(None, 27)
                for i0_40_ in range(nw239_.length(0)):
                    nw239_[i0_40_] = init40_(i0_40_)
                d_1368_v111_ = nw239_
                d_1368_v111_ = d_1368_v111_
                d_1370_v112_: _dafny.Seq
                d_1370_v112_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aj"))
                (globalState).f22 = (0) - (len(d_1370_v112_))
                d_1371_v113_: _dafny.Map
                d_1371_v113_ = _dafny.Map({len(d_1370_v112_): (D7_DC16(default__.fm2(globalState), p1)).cf23})
                d_1372_v114_: _dafny.Map
                d_1372_v114_ = d_1371_v113_
                d_1373_v115_: _dafny.MultiSet
                d_1373_v115_ = _dafny.MultiSet([(d_1372_v114_)])
                d_1374_v117_: _dafny.MultiSet
                d_1374_v117_ = _dafny.MultiSet([(self).f38])
                d_1375_v118_: _dafny.Seq
                def iife141_():
                    coll61_ = _dafny.Map()
                    compr_61_: int
                    for compr_61_ in (d_1374_v117_).Elements:
                        d_1376_v116_: int = compr_61_
                        if (d_1376_v116_) in (d_1374_v117_):
                            coll61_[default__.safeDivisionInt(d_1376_v116_, p0)] = True
                    return _dafny.Map(coll61_)
                d_1375_v118_ = _dafny.SeqWithoutIsStrInference([(d_1373_v115_).issubset(_dafny.MultiSet([d_1371_v113_, iife141_()
                ]))])
                d_1375_v118_ = (_dafny.SeqWithoutIsStrInference([p1]) if (D7_DC17(p1, p1, (self).f38)).cf24 else d_1375_v118_)
                index208_ = default__.safeIndex(12, (d_1325_v83_).length(0))
                (d_1325_v83_)[index208_] = len((d_1371_v113_).set((self).f38, not(p1)))
                index209_ = default__.safeIndex(12, (d_1325_v83_).length(0))
                (d_1325_v83_)[index209_] = (p0) - (default__.safeModuloInt((self).f38, len(d_1370_v112_)))
            elif True:
                d_1377___mcc_h8_ = source18_.cf32
                d_1378_cf32_ = d_1377___mcc_h8_
                d_1379_v119_: _dafny.Seq
                d_1379_v119_ = _dafny.SeqWithoutIsStrInference([p0, (self).f38, p0, p0])
                d_1379_v119_ = d_1379_v119_
                d_1380_v120_: C4
                nw240_ = C4()
                nw240_.ctor__()
                d_1380_v120_ = nw240_
                d_1381_v121_: _dafny.MultiSet
                d_1381_v121_ = _dafny.MultiSet([(0) - (p0), (p0) * ((0) - (len(_dafny.SeqWithoutIsStrInference([d_1325_v83_])))), 957, (self).f38, p0])
                d_1381_v121_ = _dafny.MultiSet([-232, default__.safeModuloInt(p0, (self).f38)])
                d_1382_v122_: _dafny.Seq
                d_1382_v122_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_1383_v123_: _dafny.Map
                d_1383_v123_ = _dafny.Map({d_1223_v1_: len(d_1382_v122_)})
                d_1384_v124_: _dafny.Map
                d_1384_v124_ = _dafny.Map({not(p1): ((d_1381_v121_)[-257] if (-257) in (d_1381_v121_) else p0)})
                d_1385_v125_: T1
                nw241_ = C1()
                nw241_.ctor__(d_1384_v124_)
                d_1385_v125_ = nw241_
                d_1386_v126_: _dafny.Map
                d_1386_v126_ = _dafny.Map({d_1385_v125_: d_1383_v123_})
                d_1387_v127_: _dafny.Array
                nw242_ = _dafny.Array(None, 9)
                nw242_[int(0)] = (d_1383_v123_) | (d_1383_v123_)
                nw242_[int(1)] = (d_1383_v123_) | (d_1383_v123_)
                nw242_[int(2)] = d_1383_v123_
                nw242_[int(3)] = d_1383_v123_
                nw242_[int(4)] = d_1383_v123_
                nw242_[int(5)] = (((d_1386_v126_)[d_1385_v125_] if (d_1385_v125_) in (d_1386_v126_) else d_1383_v123_)) | (_dafny.Map({(d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]: p0}))
                nw242_[int(6)] = (d_1383_v123_) | (d_1383_v123_)
                nw242_[int(7)] = _dafny.Map({(d_1222_v0_)[default__.safeIndex(889, (d_1222_v0_).length(0))]: p0})
                nw242_[int(8)] = _dafny.Map({d_1223_v1_: len(d_1379_v119_)})
                d_1387_v127_ = nw242_
                index210_ = default__.safeIndex(809, (d_1387_v127_).length(0))
                (d_1387_v127_)[index210_] = d_1383_v123_
                index211_ = default__.safeIndex(809, (d_1387_v127_).length(0))
                (d_1387_v127_)[index211_] = d_1383_v123_
        elif True:
            d_1388_v128_: _dafny.Seq
            d_1388_v128_ = _dafny.SeqWithoutIsStrInference([p1, p1, True])
            d_1389_v129_: _dafny.MultiSet
            d_1389_v129_ = _dafny.MultiSet([(0) - ((self).f38)])
            d_1390_v130_: _dafny.Map
            d_1390_v130_ = _dafny.Map({p1: d_1389_v129_})
            d_1391_v131_: _dafny.Map
            d_1391_v131_ = _dafny.Map({default__.fm0(len(d_1390_v130_), globalState): p1})
            d_1392_v132_: _dafny.Set
            d_1392_v132_ = _dafny.Set({d_1391_v131_})
            r0 = (d_1388_v128_)[default__.safeIndex(len(d_1392_v132_), len(d_1388_v128_))]
            d_1393_v133_: _dafny.Seq
            d_1393_v133_ = _dafny.SeqWithoutIsStrInference([p0, p0, 564, (self).f38])
            d_1394_v134_: _dafny.Seq
            d_1394_v134_ = (d_1393_v133_) + (d_1393_v133_)
            d_1394_v134_ = d_1394_v134_
            d_1395_v135_: _dafny.MultiSet
            d_1395_v135_ = _dafny.MultiSet([p1, p1])
            rhs224_ = (p1) == ((d_1395_v135_).ispropersubset(d_1395_v135_))
            rhs225_ = d_1324_v82_
            rhs226_ = p0
            rhs227_ = -331
            rhs228_ = _dafny.Map({(len(d_1393_v133_)) - ((self).f38): (p0) + ((0) - (p0))})
            lhs173_ = globalState
            lhs174_ = globalState
            lhs175_ = globalState
            lhs176_ = globalState
            lhs173_.f13 = rhs224_
            d_1324_v82_ = rhs225_
            lhs174_.f0 = rhs226_
            lhs175_.f21 = rhs227_
            lhs176_.f12 = rhs228_
            if p1:
                (globalState).f13 = p1
                (globalState).f0 = default__.fm2(globalState)
                d_1396_v136_: C4
                nw243_ = C4()
                nw243_.ctor__()
                d_1396_v136_ = nw243_
                d_1324_v82_ = _dafny.CodePoint('v')
                d_1397_v137_: D4
                d_1397_v137_ = D4_DC8(d_1324_v82_)
                pat_let_tv33_ = d_1324_v82_
                def iife142_(_pat_let40_0):
                    def iife143_(d_1398_dt__update__tmp_h0_):
                        def iife144_(_pat_let41_0):
                            def iife145_(d_1399_dt__update_hcf13_h0_):
                                return D4_DC8(d_1399_dt__update_hcf13_h0_)
                            return iife145_(_pat_let41_0)
                        return iife144_(pat_let_tv33_)
                    return iife143_(_pat_let40_0)
                d_1397_v137_ = iife142_(d_1397_v137_)
            elif True:
                d_1400_v138_: _dafny.Seq
                d_1400_v138_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
                (globalState).f0 = default__.safeDivisionInt(((self).f38) - (len(d_1400_v138_)), -939)
                (globalState).f13 = p1
                d_1401_v139_: D4
                d_1401_v139_ = D4_DC9(p0)
                d_1402_v140_: _dafny.Map
                d_1402_v140_ = _dafny.Map({(self).f38: d_1401_v139_})
                d_1402_v140_ = (d_1402_v140_).set(p0, d_1401_v139_)
                rhs229_ = p0
                rhs230_ = False
                lhs177_ = globalState
                lhs178_ = globalState
                lhs177_.f0 = rhs229_
                lhs178_.f13 = rhs230_
                d_1403_v141_: C1
                nw244_ = C1()
                nw244_.ctor__(default__.fm33(_dafny.SeqWithoutIsStrInference([p0]), False, globalState))
                d_1403_v141_ = nw244_
                d_1404_v142_: _dafny.Map
                d_1404_v142_ = _dafny.Map({(0) - (p0): d_1403_v141_})
                d_1405_v143_: _dafny.MultiSet
                d_1405_v143_ = default__.fm46(globalState)
                d_1406_v144_: _dafny.Map
                d_1406_v144_ = _dafny.Map({d_1404_v142_: d_1405_v143_})
                d_1407_v145_: _dafny.MultiSet
                d_1407_v145_ = _dafny.MultiSet([d_1400_v138_, d_1400_v138_])
                d_1406_v144_ = (d_1406_v144_).set(_dafny.Map({(self).f38: d_1403_v141_}), d_1407_v145_)
            (globalState).f16 = default__.safeModuloInt((self).f38, (self).f38)
        d_1408_v146_: _dafny.Seq
        d_1408_v146_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1324_v82_ for d_1409_i15_ in range(default__.abs(870))])])
        d_1410_v147_: _dafny.MultiSet
        d_1410_v147_ = _dafny.MultiSet([(self).f38])
        d_1411_v148_: _dafny.Map
        d_1411_v148_ = _dafny.Map({p1: p1})
        d_1412_v149_: _dafny.Seq
        d_1412_v149_ = _dafny.SeqWithoutIsStrInference([d_1411_v148_])
        r0 = ((_dafny.MultiSet([len(d_1412_v149_)])).set((self).f38, default__.abs(p0))).issubset((default__.fm43((self).f38, p0, (d_1408_v146_)[default__.safeIndex(p0, len(d_1408_v146_))], globalState)) - (d_1410_v147_))
        d_1413_v150_: _dafny.Map
        d_1413_v150_ = _dafny.Map({p1: p0})
        d_1414_v151_: _dafny.Seq
        d_1414_v151_ = _dafny.SeqWithoutIsStrInference([(self).f38, (self).f38, (self).f38, p0, (self).f38])
        d_1415_v152_: _dafny.Seq
        d_1415_v152_ = _dafny.SeqWithoutIsStrInference([d_1413_v150_, default__.fm33(d_1414_v151_, p1, globalState)])
        d_1416_v153_: _dafny.Seq
        d_1416_v153_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1417_v154_: _dafny.Map
        d_1417_v154_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([-567]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfli"))})
        d_1418_v155_: _dafny.Array
        nw245_ = _dafny.Array(None, 15)
        nw245_[int(0)] = d_1413_v150_
        nw245_[int(1)] = (d_1413_v150_) | (d_1413_v150_)
        nw245_[int(2)] = (d_1413_v150_) | (d_1413_v150_)
        nw245_[int(3)] = (d_1413_v150_) | (d_1413_v150_)
        nw245_[int(4)] = _dafny.Map({p1: 486})
        nw245_[int(5)] = d_1413_v150_
        nw245_[int(6)] = (d_1413_v150_) | (d_1413_v150_)
        nw245_[int(7)] = (d_1415_v152_)[default__.safeIndex(p0, len(d_1415_v152_))]
        nw245_[int(8)] = d_1413_v150_
        nw245_[int(9)] = _dafny.Map({p1: p0})
        nw245_[int(10)] = (d_1413_v150_ if (d_1416_v153_)[default__.safeIndex(p0, len(d_1416_v153_))] else d_1413_v150_)
        nw245_[int(11)] = (d_1413_v150_) | (_dafny.Map({False: (self).f38}))
        nw245_[int(12)] = d_1413_v150_
        nw245_[int(13)] = _dafny.Map({p1: len(d_1417_v154_)})
        nw245_[int(14)] = (d_1413_v150_) | (_dafny.Map({p1: (self).f38}))
        d_1418_v155_ = nw245_
        r1 = d_1418_v155_
        d_1419_v156_: _dafny.Seq
        d_1419_v156_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orggosuyu"))
        r2 = len(default__.fm30(p0, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsslvtq"))) + (d_1419_v156_), (self).f38, globalState))
        d_1420_v157_: _dafny.Set
        d_1420_v157_ = _dafny.Set({p1})
        d_1421_v158_: D5
        d_1421_v158_ = D5_DC12(False, p1)
        nw246_ = _dafny.Array(None, 27)
        nw246_[int(0)] = p1
        nw246_[int(1)] = p1
        nw246_[int(2)] = p1
        nw246_[int(3)] = p1
        nw246_[int(4)] = p1
        nw246_[int(5)] = p1
        nw246_[int(6)] = not(p1)
        nw246_[int(7)] = True
        nw246_[int(8)] = (d_1420_v157_) != (d_1420_v157_)
        nw246_[int(9)] = p1
        nw246_[int(10)] = True
        nw246_[int(11)] = (p1) == (not(False))
        nw246_[int(12)] = p1
        nw246_[int(13)] = not(p1)
        nw246_[int(14)] = (p0) <= (len(d_1414_v151_))
        nw246_[int(15)] = p1
        nw246_[int(16)] = p1
        nw246_[int(17)] = (False) or (p1)
        nw246_[int(18)] = p1
        nw246_[int(19)] = p1
        nw246_[int(20)] = not(((d_1416_v153_)[default__.safeIndex((self).f38, len(d_1416_v153_))] if p1 else p1))
        nw246_[int(21)] = ((d_1421_v158_).cf18) or (p1)
        nw246_[int(22)] = False
        nw246_[int(23)] = p1
        nw246_[int(24)] = (p1) == (p1)
        nw246_[int(25)] = p1
        nw246_[int(26)] = not(not(p1))
        r3 = nw246_
        return r0, r1, r2, r3

    @property
    def f38(self):
        return self._f38

class C6:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm20(self, p0, p1, globalState):
        return _dafny.Map({not(False): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqxjgj"))) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1422_i0_ in range(default__.abs(487))]))})

    def m14(self, p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: _dafny.Set = _dafny.Set({})
        d_1423_i0_: int
        d_1423_i0_ = 0
        with _dafny.label("7"):
            while p2:
                with _dafny.c_label("7"):
                    if (d_1423_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_1423_i0_ = (d_1423_i0_) + (1)
                    (globalState).f2 = p1
                    d_1424_v0_: _dafny.Array
                    def lambda70_(d_1425_i1_):
                        return _dafny.CodePoint('q')

                    init41_ = lambda70_
                    nw247_ = _dafny.Array(None, 22)
                    for i0_41_ in range(nw247_.length(0)):
                        nw247_[i0_41_] = init41_(i0_41_)
                    d_1424_v0_ = nw247_
                    d_1426_v1_: _dafny.Map
                    d_1426_v1_ = _dafny.Map({p2: d_1424_v0_})
                    d_1427_v2_: C0
                    nw248_ = C0()
                    nw248_.ctor__(((d_1426_v1_)[p2] if (p2) in (d_1426_v1_) else d_1424_v0_))
                    d_1427_v2_ = nw248_
                    d_1428_v3_: _dafny.Seq
                    d_1428_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfdgttee"))
                    d_1428_v3_ = d_1428_v3_
                    d_1429_v4_: str
                    d_1429_v4_ = _dafny.CodePoint('y')
                    d_1430_v5_: _dafny.MultiSet
                    d_1430_v5_ = _dafny.MultiSet([d_1429_v4_, d_1429_v4_])
                    d_1431_v6_: _dafny.Array
                    nw249_ = _dafny.Array(None, 9)
                    nw249_[int(0)] = d_1430_v5_
                    nw249_[int(1)] = d_1430_v5_
                    nw249_[int(2)] = d_1430_v5_
                    nw249_[int(3)] = _dafny.MultiSet([d_1429_v4_, d_1429_v4_])
                    nw249_[int(4)] = d_1430_v5_
                    nw249_[int(5)] = d_1430_v5_
                    nw249_[int(6)] = d_1430_v5_
                    nw249_[int(7)] = d_1430_v5_
                    nw249_[int(8)] = _dafny.MultiSet([_dafny.CodePoint('h')])
                    d_1431_v6_ = nw249_
                    pat_let_tv34_ = d_1431_v6_
                    def iife146_(_pat_let42_0):
                        def iife147_(d_1432_dt__update__tmp_h0_):
                            def iife148_(_pat_let43_0):
                                def iife149_(d_1433_dt__update_hcf6_h0_):
                                    return D2_DC4(d_1433_dt__update_hcf6_h0_)
                                return iife149_(_pat_let43_0)
                            return iife148_(pat_let_tv34_)
                        return iife147_(_pat_let42_0)
                    source19_ = iife146_(D2_DC4(d_1431_v6_))
                    if source19_.is_DC5:
                        (globalState).f1 = p2
                        (globalState).f2 = len(default__.fm30(p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), (p1) + (p1), globalState))
                        d_1434_v7_: _dafny.Map
                        d_1434_v7_ = _dafny.Map({(0) - (p1): default__.fm0((0) - (p1), globalState)})
                        d_1435_v8_: C4
                        nw250_ = C4()
                        nw250_.ctor__()
                        d_1435_v8_ = nw250_
                        d_1436_v9_: _dafny.MultiSet
                        d_1436_v9_ = _dafny.MultiSet([d_1435_v8_, d_1435_v8_])
                        d_1434_v7_ = (d_1434_v7_) | ((d_1434_v7_).set((d_1436_v9_).cardinality, False))
                        d_1437_v10_: _dafny.Set
                        d_1437_v10_ = _dafny.Set({p1})
                        d_1438_v12_: _dafny.Seq
                        d_1439_v13_: _dafny.MultiSet
                        d_1440_v14_: int
                        out45_: _dafny.Seq
                        out46_: _dafny.MultiSet
                        out47_: int
                        def iife150_():
                            coll62_ = _dafny.Set()
                            compr_62_: int
                            for compr_62_ in _dafny.IntegerRange(749, 526):
                                d_1441_v11_: int = compr_62_
                                if ((749) <= (d_1441_v11_)) and ((d_1441_v11_) < (526)):
                                    coll62_ = coll62_.union(_dafny.Set([(d_1441_v11_) + (-824)]))
                            return _dafny.Set(coll62_)
                        out45_, out46_, out47_ = (d_1435_v8_).m18((d_1437_v10_).intersection(iife150_()
                        ), globalState)
                        d_1438_v12_ = out45_
                        d_1439_v13_ = out46_
                        d_1440_v14_ = out47_
                    elif True:
                        d_1442___mcc_h0_ = source19_.cf6
                        d_1443_cf6_ = d_1442___mcc_h0_
                        (globalState).f13 = p2
                        d_1444_v15_: _dafny.Map
                        d_1444_v15_ = _dafny.Map({p2: p2})
                        d_1445_v16_: _dafny.Seq
                        d_1445_v16_ = _dafny.SeqWithoutIsStrInference([d_1428_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjudtuuv"))])
                        d_1446_v17_: _dafny.Set
                        d_1446_v17_ = _dafny.Set({((d_1445_v16_)[default__.safeIndex(p1, len(d_1445_v16_))]).set(default__.safeIndex(p1, len((d_1445_v16_)[default__.safeIndex(p1, len(d_1445_v16_))])), d_1429_v4_)})
                        d_1447_v18_: D8
                        d_1447_v18_ = D8_DC19(p2, p2, p2)
                        d_1448_v19_: _dafny.MultiSet
                        d_1448_v19_ = _dafny.MultiSet([d_1447_v18_, d_1447_v18_])
                        d_1449_v20_: _dafny.Array
                        nw251_ = _dafny.Array(None, 28)
                        nw251_[int(0)] = p2
                        nw251_[int(1)] = False
                        nw251_[int(2)] = default__.fm0(p1, globalState)
                        nw251_[int(3)] = True
                        nw251_[int(4)] = p2
                        nw251_[int(5)] = default__.fm0(-670, globalState)
                        nw251_[int(6)] = (default__.fm47(p1, d_1444_v15_, p1, globalState)).isdisjoint(d_1446_v17_)
                        nw251_[int(7)] = default__.fm0(p1, globalState)
                        nw251_[int(8)] = not(p2)
                        nw251_[int(9)] = (p1) > (737)
                        nw251_[int(10)] = not(p2)
                        nw251_[int(11)] = p2
                        nw251_[int(12)] = (351) == (p1)
                        nw251_[int(13)] = p2
                        nw251_[int(14)] = not(p2)
                        nw251_[int(15)] = False
                        nw251_[int(16)] = p2
                        nw251_[int(17)] = (d_1448_v19_).issubset(d_1448_v19_)
                        nw251_[int(18)] = p2
                        nw251_[int(19)] = p2
                        nw251_[int(20)] = (p1) == (p1)
                        nw251_[int(21)] = p2
                        nw251_[int(22)] = (p2) == (p2)
                        nw251_[int(23)] = p2
                        nw251_[int(24)] = not(p2)
                        nw251_[int(25)] = p2
                        nw251_[int(26)] = p2
                        nw251_[int(27)] = p2
                        d_1449_v20_ = nw251_
                        d_1450_v21_: D11
                        d_1450_v21_ = D11_DC26(d_1449_v20_)
                        d_1449_v20_ = (d_1450_v21_).cf39
                        d_1451_v22_: _dafny.Array
                        nw252_ = _dafny.Array(None, 2)
                        nw252_[int(0)] = d_1449_v20_
                        nw252_[int(1)] = d_1449_v20_
                        d_1451_v22_ = nw252_
                        rhs231_ = (0) - (default__.fm2(globalState))
                        rhs232_ = d_1451_v22_
                        lhs179_ = globalState
                        lhs179_.f0 = rhs231_
                        d_1451_v22_ = rhs232_
                        (globalState).f21 = p1
                    pass
            pass
        d_1452_v23_: _dafny.Array
        nw253_ = _dafny.Array(int(0), 29)
        d_1452_v23_ = nw253_
        d_1453_v24_: str
        d_1453_v24_ = _dafny.CodePoint('q')
        d_1454_v25_: D0
        d_1454_v25_ = D0_DC1(d_1452_v23_, d_1453_v24_, p1)
        source20_ = d_1454_v25_
        if source20_.is_DC1:
            d_1455___mcc_h1_ = source20_.cf1
            d_1456___mcc_h2_ = source20_.cf2
            d_1457___mcc_h3_ = source20_.cf3
            d_1458_cf3_ = d_1457___mcc_h3_
            d_1459_cf2_ = d_1456___mcc_h2_
            d_1460_cf1_ = d_1455___mcc_h1_
            d_1461_v26_: bool
            d_1462_v27_: bool
            d_1463_v28_: _dafny.Map
            d_1464_v29_: int
            out48_: bool
            out49_: bool
            out50_: _dafny.Map
            out51_: int
            out48_, out49_, out50_, out51_ = default__.m0(globalState)
            d_1461_v26_ = out48_
            d_1462_v27_ = out49_
            d_1463_v28_ = out50_
            d_1464_v29_ = out51_
            d_1465_v30_: _dafny.Map
            d_1465_v30_ = _dafny.Map({default__.fm2(globalState): -13})
            d_1465_v30_ = (d_1465_v30_).set(248, d_1458_cf3_)
            d_1466_v31_: bool
            d_1467_v32_: bool
            d_1468_v33_: _dafny.Map
            d_1469_v34_: int
            out52_: bool
            out53_: bool
            out54_: _dafny.Map
            out55_: int
            out52_, out53_, out54_, out55_ = default__.m0(globalState)
            d_1466_v31_ = out52_
            d_1467_v32_ = out53_
            d_1468_v33_ = out54_
            d_1469_v34_ = out55_
            d_1459_cf2_ = d_1459_cf2_
        elif source20_.is_DC0:
            d_1470___mcc_h4_ = source20_.cf0
            d_1471_cf0_ = d_1470___mcc_h4_
            d_1472_v35_: _dafny.Set
            d_1472_v35_ = _dafny.Set({p2})
            d_1473_v36_: _dafny.Map
            d_1473_v36_ = _dafny.Map({p0: default__.fm0(len(d_1472_v35_), globalState)})
            d_1473_v36_ = (d_1473_v36_).set(p0, default__.fm0(p1, globalState))
            d_1474_v37_: _dafny.Seq
            d_1474_v37_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_1475_v38_: _dafny.Map
            d_1475_v38_ = _dafny.Map({p1: d_1452_v23_})
            d_1476_v39_: _dafny.MultiSet
            d_1476_v39_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1453_v24_ for d_1477_i2_ in range(default__.abs(-668))])), len(d_1475_v38_), default__.fm2(globalState)])
            d_1478_v40_: _dafny.Seq
            d_1478_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            rhs233_ = (len(d_1474_v37_)) > ((d_1476_v39_).cardinality)
            rhs234_ = len((default__.fm30(len(d_1478_v40_), d_1478_v40_, p1, globalState)) + ((d_1478_v40_) + (_dafny.SeqWithoutIsStrInference([d_1453_v24_ for d_1479_i3_ in range(default__.abs(841))]))))
            lhs180_ = globalState
            lhs181_ = globalState
            lhs180_.f1 = rhs233_
            lhs181_.f2 = rhs234_
            d_1480_v41_: _dafny.Map
            d_1480_v41_ = _dafny.Map({p2: p2})
            if ((d_1480_v41_) | (d_1480_v41_)) != (_dafny.Map({p2: p2})):
                index212_ = default__.safeIndex(10, (d_1452_v23_).length(0))
                (d_1452_v23_)[index212_] = p1
                d_1481_v42_: _dafny.Array
                def lambda71_(d_1482_v35_):
                    def lambda72_(d_1483_i4_):
                        return d_1482_v35_

                    return lambda72_

                init42_ = lambda71_(d_1472_v35_)
                nw254_ = _dafny.Array(None, 17)
                for i0_42_ in range(nw254_.length(0)):
                    nw254_[i0_42_] = init42_(i0_42_)
                d_1481_v42_ = nw254_
                d_1484_v43_: D7
                d_1484_v43_ = D7_DC15(d_1481_v42_)
                d_1485_v44_: C3
                nw255_ = C3()
                nw255_.ctor__(d_1452_v23_, p2, _dafny.Map({p1: p1}))
                d_1485_v44_ = nw255_
                d_1486_v45_: _dafny.Seq
                d_1486_v45_ = _dafny.SeqWithoutIsStrInference([d_1485_v44_, d_1485_v44_])
                d_1487_v46_: _dafny.Map
                d_1487_v46_ = _dafny.Map({d_1484_v43_: ((_dafny.SeqWithoutIsStrInference([d_1485_v44_, d_1485_v44_, d_1485_v44_, d_1485_v44_, d_1485_v44_])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_1485_v44_, d_1485_v44_, d_1485_v44_, d_1485_v44_, d_1485_v44_]))), d_1485_v44_)) + (d_1486_v45_)})
                d_1488_v47_: _dafny.Set
                d_1488_v47_ = _dafny.Set({p0})
                index213_ = default__.safeIndex(10, (d_1452_v23_).length(0))
                rhs235_ = (329) - (p1)
                rhs236_ = p1
                rhs237_ = d_1487_v46_
                rhs238_ = (d_1488_v47_).ispropersubset(d_1488_v47_)
                lhs182_ = d_1452_v23_
                lhs183_ = default__.safeIndex(10, (d_1452_v23_).length(0))
                lhs184_ = globalState
                lhs185_ = d_1485_v44_
                lhs182_[lhs183_] = rhs235_
                lhs184_.f21 = rhs236_
                d_1487_v46_ = rhs237_
                lhs185_.f40 = rhs238_
                d_1489_v48_: _dafny.MultiSet
                d_1489_v48_ = _dafny.MultiSet([p2, d_1485_v44_.f40])
                d_1490_v49_: D10
                d_1490_v49_ = D10_DC25(247, p1, d_1485_v44_.f40, d_1485_v44_.f40)
                d_1491_v50_: _dafny.Seq
                d_1491_v50_ = _dafny.SeqWithoutIsStrInference([d_1490_v49_, d_1490_v49_, d_1490_v49_])
                d_1492_v51_: _dafny.Map
                d_1492_v51_ = _dafny.Map({((d_1489_v48_)[d_1485_v44_.f40] if (d_1485_v44_.f40) in (d_1489_v48_) else p1): d_1491_v50_})
                d_1493_v52_: _dafny.Map
                d_1493_v52_ = _dafny.Map({(((d_1492_v51_)[p1] if (p1) in (d_1492_v51_) else d_1491_v50_)) + (d_1491_v50_): _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([p1])).set(default__.safeIndex(default__.fm2(globalState), len(_dafny.SeqWithoutIsStrInference([p1]))), (d_1452_v23_)[default__.safeIndex(10, (d_1452_v23_).length(0))]))})
                d_1494_v53_: _dafny.Map
                d_1494_v53_ = _dafny.Map({p2: d_1476_v39_})
                d_1493_v52_ = (d_1493_v52_).set(d_1491_v50_, ((d_1494_v53_)[d_1485_v44_.f40] if (d_1485_v44_.f40) in (d_1494_v53_) else d_1476_v39_))
                d_1495_v54_: D4
                d_1495_v54_ = D4_DC10(len(d_1478_v40_))
                d_1496_v55_: _dafny.Seq
                d_1496_v55_ = _dafny.SeqWithoutIsStrInference([True])
                d_1497_v56_: _dafny.Map
                d_1497_v56_ = _dafny.Map({d_1496_v55_: p1})
                index214_ = default__.safeIndex(10, (d_1452_v23_).length(0))
                (d_1452_v23_)[index214_] = (((d_1495_v54_).cf15) - (p1)) - (len(d_1497_v56_))
                (globalState).f1 = not((d_1485_v44_.f40 if (d_1485_v44_.f40 if p2 else not(d_1485_v44_.f40)) else p2))
                (globalState).f22 = p1
            elif True:
                d_1498_v57_: _dafny.Map
                d_1498_v57_ = _dafny.Map({p1: not(default__.fm0(p1, globalState))})
                d_1499_v58_: _dafny.Set
                d_1499_v58_ = _dafny.Set({d_1498_v57_})
                d_1499_v58_ = d_1499_v58_
                index215_ = default__.safeIndex(42, (p0).length(0))
                (p0)[index215_] = False
                index216_ = default__.safeIndex(42, (p0).length(0))
                (p0)[index216_] = p2
                (globalState).f16 = (default__.fm2(globalState) if p2 else 462)
                (globalState).f21 = (p1) - (default__.fm2(globalState))
                d_1500_v59_: _dafny.Array
                nw256_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_1500_v59_ = nw256_
                index217_ = default__.safeIndex(279, (d_1500_v59_).length(0))
                (d_1500_v59_)[index217_] = d_1478_v40_
                index218_ = default__.safeIndex(279, (d_1500_v59_).length(0))
                (d_1500_v59_)[index218_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdjqocpn"))
            index219_ = default__.safeIndex(106, (p0).length(0))
            (p0)[index219_] = (d_1474_v37_) <= (_dafny.SeqWithoutIsStrInference([p1]))
            index220_ = default__.safeIndex(106, (p0).length(0))
            (p0)[index220_] = p2
        elif True:
            d_1501___mcc_h5_ = source20_.cf4
            d_1502_cf4_ = d_1501___mcc_h5_
            (globalState).f1 = (p1) > ((p1) - ((0) - (p1)))
            index221_ = default__.safeIndex(884, (p0).length(0))
            (p0)[index221_] = p2
            index222_ = default__.safeIndex(884, (p0).length(0))
            (p0)[index222_] = not(p2)
            d_1503_v60_: _dafny.Array
            nw257_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_1503_v60_ = nw257_
            d_1503_v60_ = d_1503_v60_
            if p2:
                d_1504_v61_: _dafny.Map
                d_1504_v61_ = _dafny.Map({p2: p1})
                d_1505_v62_: C1
                nw258_ = C1()
                nw258_.ctor__(d_1504_v61_)
                d_1505_v62_ = nw258_
                d_1506_v63_: _dafny.Seq
                d_1506_v63_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(884, (p0).length(0))]])
                (globalState).f13 = (d_1506_v63_)[default__.safeIndex(p1, len(d_1506_v63_))]
                d_1507_v64_: _dafny.Map
                d_1507_v64_ = _dafny.Map({(p1 if p2 else p1): default__.fm0(p1, globalState)})
                d_1508_v65_: _dafny.Set
                d_1508_v65_ = _dafny.Set({d_1453_v24_})
                d_1509_v66_: _dafny.Set
                d_1509_v66_ = _dafny.Set({p1, len(d_1507_v64_)})
                rhs239_ = ((d_1507_v64_)[(p1 if (p0)[default__.safeIndex(884, (p0).length(0))] else len(d_1508_v65_))] if ((p1 if (p0)[default__.safeIndex(884, (p0).length(0))] else len(d_1508_v65_))) in (d_1507_v64_) else (d_1509_v66_).ispropersubset(d_1509_v66_))
                rhs240_ = (p0)[default__.safeIndex(884, (p0).length(0))]
                lhs186_ = globalState
                lhs187_ = globalState
                lhs186_.f13 = rhs239_
                lhs187_.f13 = rhs240_
                (globalState).f13 = (p0)[default__.safeIndex(884, (p0).length(0))]
                d_1510_v67_: _dafny.Array
                nw259_ = _dafny.Array(_dafny.Set({}), 27)
                d_1510_v67_ = nw259_
                d_1510_v67_ = d_1510_v67_
            elif True:
                (globalState).f13 = True
                d_1511_v68_: _dafny.Seq
                d_1511_v68_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p1, globalState)])
                d_1512_v69_: _dafny.Map
                d_1512_v69_ = _dafny.Map({p1: (p0)[default__.safeIndex(884, (p0).length(0))]})
                d_1513_v70_: _dafny.Map
                d_1513_v70_ = d_1512_v69_
                d_1514_v71_: _dafny.Map
                d_1514_v71_ = _dafny.Map({not((True) in (d_1511_v68_)): d_1513_v70_})
                d_1514_v71_ = (d_1514_v71_).set((p1) == (89), d_1513_v70_)
                d_1515_v72_: _dafny.Seq
                d_1515_v72_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1516_v73_: _dafny.Array
                def lambda73_(d_1517_p1_):
                    def lambda74_(d_1518_i5_):
                        return _dafny.Set({(0) - (d_1517_p1_), d_1517_p1_, 518, d_1517_p1_, d_1517_p1_})

                    return lambda74_

                init43_ = lambda73_(p1)
                nw260_ = _dafny.Array(None, 2)
                for i0_43_ in range(nw260_.length(0)):
                    nw260_[i0_43_] = init43_(i0_43_)
                d_1516_v73_ = nw260_
                d_1519_v74_: _dafny.Array
                def lambda75_(d_1520_v72_):
                    def lambda76_(d_1521_i6_):
                        return (d_1520_v72_) + (d_1520_v72_)

                    return lambda76_

                init44_ = lambda75_(d_1515_v72_)
                nw261_ = _dafny.Array(None, 3)
                for i0_44_ in range(nw261_.length(0)):
                    nw261_[i0_44_] = init44_(i0_44_)
                d_1519_v74_ = nw261_
                d_1522_v75_: _dafny.MultiSet
                d_1522_v75_ = _dafny.MultiSet([p1])
                rhs241_ = default__.fm2(globalState)
                rhs242_ = d_1515_v72_
                rhs243_ = d_1516_v73_
                rhs244_ = (d_1519_v74_ if p2 else d_1519_v74_)
                rhs245_ = not((d_1522_v75_).issubset(d_1522_v75_))
                lhs188_ = globalState
                lhs189_ = globalState
                lhs188_.f0 = rhs241_
                d_1515_v72_ = rhs242_
                d_1516_v73_ = rhs243_
                d_1519_v74_ = rhs244_
                lhs189_.f1 = rhs245_
                index223_ = default__.safeIndex(884, (p0).length(0))
                (p0)[index223_] = p2
                d_1523_v76_: _dafny.Array
                nw262_ = _dafny.Array(None, 9)
                nw262_[int(0)] = d_1453_v24_
                nw262_[int(1)] = d_1453_v24_
                nw262_[int(2)] = d_1453_v24_
                nw262_[int(3)] = d_1453_v24_
                nw262_[int(4)] = default__.fm35(p1, globalState)
                nw262_[int(5)] = d_1453_v24_
                nw262_[int(6)] = d_1453_v24_
                nw262_[int(7)] = _dafny.CodePoint('y')
                nw262_[int(8)] = _dafny.CodePoint('m')
                d_1523_v76_ = nw262_
                d_1524_v77_: C0
                nw263_ = C0()
                nw263_.ctor__(d_1523_v76_)
                d_1524_v77_ = nw263_
        d_1525_v78_: _dafny.Seq
        d_1525_v78_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1526_v79_: _dafny.Map
        d_1526_v79_ = _dafny.Map({p1: _dafny.MultiSet([len(d_1525_v78_)])})
        d_1527_v80_: _dafny.Map
        d_1527_v80_ = _dafny.Map({len(d_1526_v79_): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1528_i7_ in range(default__.abs(284))])})
        d_1529_v81_: _dafny.Seq
        d_1529_v81_ = _dafny.SeqWithoutIsStrInference([d_1453_v24_, d_1453_v24_])
        d_1530_v82_: _dafny.MultiSet
        d_1530_v82_ = _dafny.MultiSet([default__.safeModuloInt(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kftyswf")))), len((((d_1527_v80_)[default__.fm2(globalState)] if (default__.fm2(globalState)) in (d_1527_v80_) else d_1529_v81_)) + (d_1529_v81_))])
        d_1530_v82_ = d_1530_v82_
        index224_ = default__.safeIndex(672, (p0).length(0))
        (p0)[index224_] = (p1) <= (p1)
        index225_ = default__.safeIndex(672, (p0).length(0))
        (p0)[index225_] = ((p2) == (default__.fm0(p1, globalState))) == (p2)
        (globalState).f13 = (p1) <= (default__.safeModuloInt(default__.fm2(globalState), -457))
        d_1531_v83_: _dafny.Map
        d_1531_v83_ = _dafny.Map({(p0)[default__.safeIndex(672, (p0).length(0))]: 743})
        d_1532_v84_: D11
        d_1532_v84_ = D11_DC27(p1, len(d_1525_v78_), False)
        d_1533_v85_: _dafny.Seq
        d_1533_v85_ = _dafny.SeqWithoutIsStrInference([not(p2), not((p0)[default__.safeIndex(672, (p0).length(0))]), (p0)[default__.safeIndex(672, (p0).length(0))], (p0)[default__.safeIndex(672, (p0).length(0))], p2])
        d_1534_v86_: D11
        d_1534_v86_ = D11_DC27(((d_1531_v83_)[(p0)[default__.safeIndex(672, (p0).length(0))]] if ((p0)[default__.safeIndex(672, (p0).length(0))]) in (d_1531_v83_) else p1), len((d_1529_v81_) + (d_1529_v81_)), ((d_1533_v85_)[default__.safeIndex(p1, len(d_1533_v85_))] if not(default__.fm0((d_1532_v84_).cf40, globalState)) else (p0)[default__.safeIndex(672, (p0).length(0))]))
        source21_ = d_1532_v84_
        if source21_.is_DC27:
            d_1535___mcc_h6_ = source21_.cf40
            d_1536___mcc_h7_ = source21_.cf41
            d_1537___mcc_h8_ = source21_.cf42
            d_1538_cf42_ = d_1537___mcc_h8_
            d_1539_cf41_ = d_1536___mcc_h7_
            d_1540_cf40_ = d_1535___mcc_h6_
            d_1541_v87_: _dafny.Set
            d_1541_v87_ = _dafny.Set({d_1538_cf42_})
            index226_ = default__.safeIndex(672, (p0).length(0))
            (p0)[index226_] = (d_1541_v87_).ispropersubset(_dafny.Set({False}))
            d_1542_v88_: _dafny.Set
            d_1542_v88_ = _dafny.Set({d_1452_v23_, d_1452_v23_, d_1452_v23_})
            d_1542_v88_ = (d_1542_v88_) | ((d_1542_v88_) - (d_1542_v88_))
            d_1540_cf40_ = p1
            d_1543_v89_: _dafny.Array
            nw264_ = _dafny.Array(None, 15)
            nw264_[int(0)] = d_1453_v24_
            nw264_[int(1)] = d_1453_v24_
            nw264_[int(2)] = d_1453_v24_
            nw264_[int(3)] = d_1453_v24_
            nw264_[int(4)] = d_1453_v24_
            nw264_[int(5)] = d_1453_v24_
            nw264_[int(6)] = (d_1529_v81_)[default__.safeIndex(d_1539_cf41_, len(d_1529_v81_))]
            nw264_[int(7)] = d_1453_v24_
            nw264_[int(8)] = d_1453_v24_
            nw264_[int(9)] = d_1453_v24_
            nw264_[int(10)] = d_1453_v24_
            nw264_[int(11)] = d_1453_v24_
            nw264_[int(12)] = d_1453_v24_
            nw264_[int(13)] = d_1453_v24_
            nw264_[int(14)] = d_1453_v24_
            d_1543_v89_ = nw264_
            d_1544_v90_: _dafny.Seq
            d_1544_v90_ = _dafny.SeqWithoutIsStrInference([d_1543_v89_])
            d_1545_v91_: C0
            nw265_ = C0()
            nw265_.ctor__((d_1544_v90_)[default__.safeIndex(d_1540_cf40_, len(d_1544_v90_))])
            d_1545_v91_ = nw265_
        elif True:
            d_1546___mcc_h9_ = source21_.cf39
            d_1547_cf39_ = d_1546___mcc_h9_
            d_1548_v92_: _dafny.Array
            def lambda77_(d_1549_v81_):
                def lambda78_(d_1550_i8_):
                    return (d_1549_v81_) + (d_1549_v81_)

                return lambda78_

            init45_ = lambda77_(d_1529_v81_)
            nw266_ = _dafny.Array(None, 17)
            for i0_45_ in range(nw266_.length(0)):
                nw266_[i0_45_] = init45_(i0_45_)
            d_1548_v92_ = nw266_
            index227_ = default__.safeIndex(898, (d_1548_v92_).length(0))
            (d_1548_v92_)[index227_] = d_1529_v81_
            d_1551_v93_: _dafny.MultiSet
            d_1551_v93_ = _dafny.MultiSet([(p0)[default__.safeIndex(672, (p0).length(0))], p2, (d_1533_v85_)[default__.safeIndex(-757, len(d_1533_v85_))]])
            d_1552_v94_: D7
            d_1552_v94_ = D7_DC16((d_1551_v93_).cardinality, (p0)[default__.safeIndex(672, (p0).length(0))])
            index228_ = default__.safeIndex(898, (d_1548_v92_).length(0))
            (d_1548_v92_)[index228_] = (default__.fm30(p1, d_1529_v81_, (d_1552_v94_).cf22, globalState)).set(default__.safeIndex((d_1525_v78_)[default__.safeIndex(len(d_1529_v81_), len(d_1525_v78_))], len(default__.fm30(p1, d_1529_v81_, (d_1552_v94_).cf22, globalState))), d_1453_v24_)
            d_1553_v95_: _dafny.Array
            nw267_ = _dafny.Array(_dafny.Map({}), 19)
            d_1553_v95_ = nw267_
            d_1554_v96_: D2
            d_1554_v96_ = D2_DC5()
            d_1555_v97_: _dafny.Map
            d_1555_v97_ = _dafny.Map({d_1554_v96_: p2})
            index229_ = default__.safeIndex(684, (d_1553_v95_).length(0))
            (d_1553_v95_)[index229_] = d_1555_v97_
            d_1556_v98_: C1
            nw268_ = C1()
            nw268_.ctor__(d_1531_v83_)
            d_1556_v98_ = nw268_
            d_1557_v99_: D12
            d_1557_v99_ = D12_DC28(d_1551_v93_)
            d_1558_v100_: _dafny.Map
            d_1558_v100_ = _dafny.Map({d_1556_v98_: d_1557_v99_})
            d_1559_v101_: _dafny.Seq
            d_1559_v101_ = _dafny.SeqWithoutIsStrInference([d_1558_v100_, d_1558_v100_])
            index230_ = default__.safeIndex(684, (d_1553_v95_).length(0))
            index231_ = default__.safeIndex(672, (p0).length(0))
            rhs246_ = d_1555_v97_
            rhs247_ = (d_1453_v24_) in ((d_1548_v92_)[default__.safeIndex(898, (d_1548_v92_).length(0))])
            rhs248_ = 55
            rhs249_ = d_1559_v101_
            lhs190_ = d_1553_v95_
            lhs191_ = default__.safeIndex(684, (d_1553_v95_).length(0))
            lhs192_ = p0
            lhs193_ = default__.safeIndex(672, (p0).length(0))
            lhs194_ = globalState
            lhs190_[lhs191_] = rhs246_
            lhs192_[lhs193_] = rhs247_
            lhs194_.f16 = rhs248_
            d_1559_v101_ = rhs249_
            (globalState).f5 = (p1) * (p1)
            if (p0)[default__.safeIndex(672, (p0).length(0))]:
                d_1560_v102_: D10
                d_1560_v102_ = D10_DC25((0) - (p1), p1, (p0)[default__.safeIndex(672, (p0).length(0))], p2)
                d_1561_v103_: _dafny.Map
                d_1561_v103_ = _dafny.Map({default__.fm48(d_1560_v102_, p1, globalState): d_1530_v82_})
                d_1562_v104_: _dafny.Set
                d_1562_v104_ = _dafny.Set({not((p0)[default__.safeIndex(672, (p0).length(0))])})
                d_1563_v105_: D5
                d_1563_v105_ = D5_DC12(p2, True)
                d_1564_v106_: _dafny.Set
                d_1564_v106_ = _dafny.Set({p1, p1})
                rhs250_ = (len(d_1564_v106_)) - (p1)
                rhs251_ = (d_1561_v103_) | ((d_1561_v103_) | (d_1561_v103_))
                rhs252_ = _dafny.Set({(d_1551_v93_).ispropersubset(d_1551_v93_)})
                rhs253_ = d_1529_v81_
                rhs254_ = d_1563_v105_
                lhs195_ = globalState
                lhs195_.f0 = rhs250_
                d_1561_v103_ = rhs251_
                d_1562_v104_ = rhs252_
                d_1529_v81_ = rhs253_
                d_1563_v105_ = rhs254_
                (globalState).f22 = p1
                (globalState).f1 = (((0) - ((0) - (default__.fm2(globalState)))) <= (p1)) or ((p0)[default__.safeIndex(672, (p0).length(0))])
                (globalState).f0 = p1
                d_1565_v107_: _dafny.Seq
                d_1565_v107_ = _dafny.SeqWithoutIsStrInference([d_1525_v78_, _dafny.SeqWithoutIsStrInference([p1]), d_1525_v78_])
                d_1566_v108_: _dafny.Map
                d_1566_v108_ = _dafny.Map({(d_1565_v107_ if False else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1]) for d_1567_i9_ in range(default__.abs(261))])): p1})
                d_1566_v108_ = (d_1566_v108_).set(d_1565_v107_, p1)
            elif True:
                d_1568_v109_: C1
                nw269_ = C1()
                nw269_.ctor__((d_1531_v83_) | (d_1531_v83_))
                d_1568_v109_ = nw269_
                d_1569_v110_: D5
                d_1569_v110_ = D5_DC13((p1 if p2 else p1))
                d_1570_v111_: _dafny.Array
                nw270_ = _dafny.Array(_dafny.CodePoint('D'), 13)
                d_1570_v111_ = nw270_
                index232_ = default__.safeIndex(448, (d_1570_v111_).length(0))
                (d_1570_v111_)[index232_] = (d_1529_v81_)[default__.safeIndex((d_1530_v82_).cardinality, len(d_1529_v81_))]
                d_1571_v112_: _dafny.Map
                d_1571_v112_ = _dafny.Map({p2: d_1452_v23_})
                d_1572_v113_: D19
                d_1572_v113_ = D19_DC37(d_1571_v112_)
                index233_ = default__.safeIndex(448, (d_1570_v111_).length(0))
                rhs255_ = d_1569_v110_
                rhs256_ = (len((d_1572_v113_).cf51)) == (p1)
                rhs257_ = d_1453_v24_
                rhs258_ = len((d_1548_v92_)[default__.safeIndex(898, (d_1548_v92_).length(0))])
                lhs196_ = globalState
                lhs197_ = d_1570_v111_
                lhs198_ = default__.safeIndex(448, (d_1570_v111_).length(0))
                lhs199_ = globalState
                d_1569_v110_ = rhs255_
                lhs196_.f1 = rhs256_
                lhs197_[lhs198_] = rhs257_
                lhs199_.f22 = rhs258_
                d_1573_v114_: _dafny.Map
                d_1573_v114_ = _dafny.Map({p1: (_dafny.MultiSet([p1])).cardinality})
                d_1574_v115_: _dafny.MultiSet
                d_1574_v115_ = _dafny.MultiSet([d_1573_v114_, d_1573_v114_, d_1573_v114_])
                d_1574_v115_ = (default__.fm49(len(d_1529_v81_), _dafny.Map({p2: d_1525_v78_}), globalState)).intersection(d_1574_v115_)
                d_1575_v116_: _dafny.Map
                d_1575_v116_ = _dafny.Map({d_1529_v81_: p1})
                d_1576_v117_: D3
                d_1576_v117_ = D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhsklbgg")), p1, p1, len(d_1575_v116_), d_1533_v85_)
                index234_ = default__.safeIndex(898, (d_1548_v92_).length(0))
                rhs259_ = d_1530_v82_
                rhs260_ = (d_1529_v81_) + ((default__.fm30(p1, _dafny.SeqWithoutIsStrInference([(d_1570_v111_)[default__.safeIndex(448, (d_1570_v111_).length(0))] for d_1577_i10_ in range(default__.abs(620))]), p1, globalState)).set(default__.safeIndex(p1, len(default__.fm30(p1, _dafny.SeqWithoutIsStrInference([(d_1570_v111_)[default__.safeIndex(448, (d_1570_v111_).length(0))] for d_1578_i10_ in range(default__.abs(620))]), p1, globalState))), (d_1570_v111_)[default__.safeIndex(448, (d_1570_v111_).length(0))]))
                rhs261_ = p1
                rhs262_ = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlianpqdv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eem")))) + ((d_1576_v117_).cf8))
                rhs263_ = p2
                lhs200_ = d_1548_v92_
                lhs201_ = default__.safeIndex(898, (d_1548_v92_).length(0))
                lhs202_ = globalState
                lhs203_ = globalState
                lhs204_ = globalState
                d_1530_v82_ = rhs259_
                lhs200_[lhs201_] = rhs260_
                lhs202_.f2 = rhs261_
                lhs203_.f22 = rhs262_
                lhs204_.f13 = rhs263_
                d_1579_v118_: C0
                nw271_ = C0()
                nw271_.ctor__(d_1570_v111_)
                d_1579_v118_ = nw271_
                d_1579_v118_ = d_1579_v118_
        d_1580_v119_: _dafny.Array
        nw272_ = _dafny.Array(_dafny.CodePoint('D'), 2)
        d_1580_v119_ = nw272_
        d_1581_v120_: D0
        d_1581_v120_ = D0_DC0(d_1580_v119_)
        r0 = d_1581_v120_
        d_1582_v121_: _dafny.Set
        d_1582_v121_ = _dafny.Set({default__.safeModuloInt(p1, p1), p1, p1, p1, p1})
        r1 = d_1582_v121_
        return r0, r1


class C7:
    def  __init__(self):
        self._f37: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self, f37):
        (self)._f37 = f37

    def fm19(self, p0, p1, globalState):
        return -69

    def m13(self, p0, globalState):
        r0: bool = False
        d_1583_v0_: _dafny.Array
        def lambda79_(d_1584_i0_):
            return (self).f37

        init46_ = lambda79_
        nw273_ = _dafny.Array(None, 10)
        for i0_46_ in range(nw273_.length(0)):
            nw273_[i0_46_] = init46_(i0_46_)
        d_1583_v0_ = nw273_
        index235_ = default__.safeIndex(372, (d_1583_v0_).length(0))
        (d_1583_v0_)[index235_] = (self).f37
        index236_ = default__.safeIndex(372, (d_1583_v0_).length(0))
        (d_1583_v0_)[index236_] = (self).f37
        d_1585_v1_: _dafny.Seq
        d_1585_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thpj"))
        (globalState).f13 = (d_1585_v1_) <= ((d_1585_v1_) + (d_1585_v1_))
        index237_ = default__.safeIndex(372, (d_1583_v0_).length(0))
        (d_1583_v0_)[index237_] = (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]
        d_1586_v2_: D7
        d_1586_v2_ = D7_DC16(p0, (self).f37)
        pat_let_tv35_ = p0
        pat_let_tv36_ = p0
        pat_let_tv37_ = p0
        def lambda80_(source22_):
            if source22_.is_DC16:
                d_1587___mcc_h0_ = source22_.cf22
                d_1588___mcc_h1_ = source22_.cf23
                d_1589_cf23_ = d_1588___mcc_h1_
                d_1590_cf22_ = d_1587___mcc_h0_
                return (pat_let_tv35_) - ((_dafny.MultiSet([d_1590_cf22_])).cardinality)
            elif source22_.is_DC17:
                d_1591___mcc_h2_ = source22_.cf24
                d_1592___mcc_h3_ = source22_.cf25
                d_1593___mcc_h4_ = source22_.cf26
                d_1594_cf26_ = d_1593___mcc_h4_
                d_1595_cf25_ = d_1592___mcc_h3_
                d_1596_cf24_ = d_1591___mcc_h2_
                return pat_let_tv36_
            elif True:
                d_1597___mcc_h5_ = source22_.cf21
                d_1598_cf21_ = d_1597___mcc_h5_
                return pat_let_tv37_

        (globalState).f21 = lambda80_(d_1586_v2_)
        d_1599_i1_: int
        d_1599_i1_ = 0
        with _dafny.label("8"):
            while (self).f37:
                with _dafny.c_label("8"):
                    if (d_1599_i1_) >= (100):
                        raise _dafny.Break("8")
                    d_1599_i1_ = (d_1599_i1_) + (1)
                    if (self).f37:
                        d_1600_v3_: C2
                        nw274_ = C2()
                        nw274_.ctor__((self).f37, p0)
                        d_1600_v3_ = nw274_
                        d_1601_v4_: _dafny.Seq
                        d_1601_v4_ = _dafny.SeqWithoutIsStrInference([(d_1600_v3_).f42])
                        d_1602_v5_: _dafny.Map
                        d_1602_v5_ = _dafny.Map({(d_1601_v4_)[default__.safeIndex((d_1600_v3_).f42, len(d_1601_v4_))]: True})
                        d_1603_v6_: D7
                        d_1604_v7_: int
                        d_1605_v8_: bool
                        out56_: D7
                        out57_: int
                        out58_: bool
                        out56_, out57_, out58_ = (d_1600_v3_).m20(((d_1602_v5_)[p0] if (p0) in (d_1602_v5_) else (self).f37), globalState)
                        d_1603_v6_ = out56_
                        d_1604_v7_ = out57_
                        d_1605_v8_ = out58_
                        (globalState).f1 = d_1605_v8_
                        d_1606_v9_: _dafny.Seq
                        d_1606_v9_ = _dafny.SeqWithoutIsStrInference([not(True), (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]])
                        d_1606_v9_ = (_dafny.SeqWithoutIsStrInference([(self).f37])) + (d_1606_v9_)
                        d_1607_v10_: _dafny.Map
                        d_1607_v10_ = _dafny.Map({(self).f37: d_1604_v7_})
                        d_1608_v11_: _dafny.Map
                        d_1608_v11_ = _dafny.Map({d_1585_v1_: p0})
                        d_1609_v12_: str
                        d_1609_v12_ = _dafny.CodePoint('d')
                        d_1610_v13_: _dafny.MultiSet
                        d_1610_v13_ = _dafny.MultiSet([d_1609_v12_])
                        d_1611_v14_: _dafny.MultiSet
                        d_1611_v14_ = _dafny.MultiSet([(d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]])
                        d_1612_v15_: _dafny.MultiSet
                        d_1612_v15_ = _dafny.MultiSet([(0) - (d_1604_v7_)])
                        d_1613_v16_: _dafny.Array
                        nw275_ = _dafny.Array(None, 25)
                        nw275_[int(0)] = (d_1600_v3_).f42
                        nw275_[int(1)] = d_1604_v7_
                        nw275_[int(2)] = 971
                        nw275_[int(3)] = p0
                        nw275_[int(4)] = -283
                        nw275_[int(5)] = p0
                        nw275_[int(6)] = 732
                        nw275_[int(7)] = (d_1600_v3_).f42
                        nw275_[int(8)] = (d_1600_v3_).f42
                        nw275_[int(9)] = d_1604_v7_
                        nw275_[int(10)] = d_1604_v7_
                        nw275_[int(11)] = p0
                        nw275_[int(12)] = p0
                        nw275_[int(13)] = 520
                        nw275_[int(14)] = (d_1600_v3_).f42
                        nw275_[int(15)] = d_1604_v7_
                        nw275_[int(16)] = p0
                        nw275_[int(17)] = (d_1611_v14_).cardinality
                        nw275_[int(18)] = d_1604_v7_
                        nw275_[int(19)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdsgsrlsa")))
                        nw275_[int(20)] = (d_1612_v15_).cardinality
                        nw275_[int(21)] = d_1604_v7_
                        nw275_[int(22)] = d_1604_v7_
                        nw275_[int(23)] = d_1604_v7_
                        nw275_[int(24)] = p0
                        d_1613_v16_ = nw275_
                        d_1614_v18_: C3
                        nw276_ = C3()
                        def iife151_():
                            coll63_ = _dafny.Map()
                            compr_63_: int
                            for compr_63_ in _dafny.IntegerRange(926, -66):
                                d_1615_v17_: int = compr_63_
                                if ((926) <= (d_1615_v17_)) and ((d_1615_v17_) < (-66)):
                                    coll63_[(d_1615_v17_) - (len(d_1601_v4_))] = d_1604_v7_
                            return _dafny.Map(coll63_)
                        nw276_.ctor__(d_1613_v16_, (d_1600_v3_).f41, iife151_()
                        )
                        d_1614_v18_ = nw276_
                        d_1616_v19_: _dafny.Map
                        d_1616_v19_ = _dafny.Map({d_1614_v18_: 857})
                        d_1617_v20_: _dafny.Map
                        d_1617_v20_ = _dafny.Map({(d_1600_v3_).f42: d_1614_v18_})
                        d_1618_v21_: _dafny.Array
                        nw277_ = _dafny.Array(None, 20)
                        nw277_[int(0)] = p0
                        nw277_[int(1)] = (d_1600_v3_).f42
                        nw277_[int(2)] = d_1604_v7_
                        nw277_[int(3)] = (d_1600_v3_).f42
                        nw277_[int(4)] = d_1604_v7_
                        nw277_[int(5)] = len((d_1616_v19_).set(((d_1617_v20_)[p0] if (p0) in (d_1617_v20_) else d_1614_v18_), 831))
                        nw277_[int(6)] = len(d_1601_v4_)
                        nw277_[int(7)] = d_1604_v7_
                        nw277_[int(8)] = (d_1600_v3_).f42
                        nw277_[int(9)] = p0
                        nw277_[int(10)] = len(d_1607_v10_)
                        nw277_[int(11)] = p0
                        nw277_[int(12)] = -667
                        nw277_[int(13)] = p0
                        nw277_[int(14)] = d_1604_v7_
                        nw277_[int(15)] = d_1604_v7_
                        nw277_[int(16)] = p0
                        nw277_[int(17)] = (d_1600_v3_).f42
                        nw277_[int(18)] = 560
                        nw277_[int(19)] = p0
                        d_1618_v21_ = nw277_
                        d_1619_v22_: _dafny.Map
                        d_1619_v22_ = _dafny.Map({(d_1586_v2_).cf22: len(d_1585_v1_)})
                        d_1620_v23_: T0
                        nw278_ = C3()
                        nw278_.ctor__(d_1613_v16_, d_1614_v18_.f40, d_1619_v22_)
                        d_1620_v23_ = nw278_
                        d_1621_v24_: _dafny.Seq
                        d_1621_v24_ = _dafny.SeqWithoutIsStrInference([d_1620_v23_, d_1620_v23_, d_1620_v23_])
                        d_1622_v25_: _dafny.Map
                        d_1622_v25_ = _dafny.Map({d_1614_v18_.f40: (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]})
                        d_1623_v26_: _dafny.Map
                        d_1623_v26_ = _dafny.Map({(d_1600_v3_).f42: d_1622_v25_})
                        d_1624_v27_: _dafny.Array
                        nw279_ = _dafny.Array(None, 28)
                        nw279_[int(0)] = (d_1600_v3_).f42
                        nw279_[int(1)] = 164
                        nw279_[int(2)] = len(d_1607_v10_)
                        nw279_[int(3)] = len(d_1608_v11_)
                        nw279_[int(4)] = d_1604_v7_
                        nw279_[int(5)] = (d_1601_v4_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0 for d_1625_i2_ in range(default__.abs(767))])), len(d_1601_v4_))]
                        nw279_[int(6)] = (d_1600_v3_).f42
                        nw279_[int(7)] = (102) * ((d_1610_v13_).cardinality)
                        nw279_[int(8)] = (d_1600_v3_).f42
                        nw279_[int(9)] = d_1604_v7_
                        nw279_[int(10)] = (0) - ((0) - (d_1604_v7_))
                        nw279_[int(11)] = d_1604_v7_
                        nw279_[int(12)] = 494
                        nw279_[int(13)] = len(d_1601_v4_)
                        nw279_[int(14)] = ((d_1611_v14_)[(d_1600_v3_).f41] if ((d_1600_v3_).f41) in (d_1611_v14_) else -264)
                        nw279_[int(15)] = (len(d_1606_v9_) if (d_1606_v9_)[default__.safeIndex((0) - (len(d_1585_v1_)), len(d_1606_v9_))] else p0)
                        nw279_[int(16)] = (d_1600_v3_).f42
                        nw279_[int(17)] = d_1604_v7_
                        nw279_[int(18)] = default__.safeModuloInt(p0, (d_1600_v3_).f42)
                        nw279_[int(19)] = 271
                        nw279_[int(20)] = default__.safeDivisionInt(d_1604_v7_, (d_1600_v3_).f42)
                        nw279_[int(21)] = d_1604_v7_
                        nw279_[int(22)] = (d_1612_v15_).cardinality
                        nw279_[int(23)] = len(d_1621_v24_)
                        nw279_[int(24)] = 991
                        nw279_[int(25)] = ((d_1611_v14_)[True] if (True) in (d_1611_v14_) else 515)
                        nw279_[int(26)] = p0
                        nw279_[int(27)] = (d_1601_v4_)[default__.safeIndex(len(d_1623_v26_), len(d_1601_v4_))]
                        d_1624_v27_ = nw279_
                        index238_ = default__.safeIndex(428, (d_1618_v21_).length(0))
                        (d_1618_v21_)[index238_] = (d_1604_v7_) * (len(d_1606_v9_))
                        d_1626_v28_: _dafny.Map
                        d_1626_v28_ = _dafny.Map({(d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]: d_1606_v9_})
                        index239_ = default__.safeIndex(428, (d_1618_v21_).length(0))
                        rhs264_ = (d_1600_v3_).f42
                        rhs265_ = d_1611_v14_
                        rhs266_ = d_1614_v18_.f40
                        rhs267_ = (self).f37
                        rhs268_ = (d_1604_v7_) == ((0) - (((d_1600_v3_).f42) - (len(((d_1626_v28_)[(d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]] if ((d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]) in (d_1626_v28_) else d_1606_v9_)))))
                        lhs205_ = d_1618_v21_
                        lhs206_ = default__.safeIndex(428, (d_1618_v21_).length(0))
                        lhs207_ = globalState
                        lhs208_ = globalState
                        lhs209_ = globalState
                        lhs205_[lhs206_] = rhs264_
                        d_1611_v14_ = rhs265_
                        lhs207_.f13 = rhs266_
                        lhs208_.f13 = rhs267_
                        lhs209_.f13 = rhs268_
                    elif True:
                        d_1627_v29_: bool
                        d_1628_v30_: bool
                        d_1629_v31_: _dafny.Map
                        d_1630_v32_: int
                        out59_: bool
                        out60_: bool
                        out61_: _dafny.Map
                        out62_: int
                        out59_, out60_, out61_, out62_ = default__.m0(globalState)
                        d_1627_v29_ = out59_
                        d_1628_v30_ = out60_
                        d_1629_v31_ = out61_
                        d_1630_v32_ = out62_
                        d_1583_v0_ = d_1583_v0_
                        d_1631_v33_: _dafny.Seq
                        d_1631_v33_ = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
                        d_1632_v34_: _dafny.Seq
                        d_1632_v34_ = d_1631_v33_
                        d_1633_v35_: _dafny.Seq
                        d_1633_v35_ = _dafny.SeqWithoutIsStrInference([d_1632_v34_, _dafny.SeqWithoutIsStrInference([d_1630_v32_, p0])])
                        d_1634_v36_: _dafny.Map
                        d_1634_v36_ = _dafny.Map({d_1630_v32_: True})
                        d_1635_v37_: _dafny.Map
                        d_1635_v37_ = _dafny.Map({(d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]: len((d_1634_v36_).set(p0, (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]))})
                        d_1636_v38_: _dafny.Map
                        d_1636_v38_ = _dafny.Map({(d_1633_v35_).set(default__.safeIndex(((d_1635_v37_)[d_1628_v30_] if (d_1628_v30_) in (d_1635_v37_) else p0), len(d_1633_v35_)), (d_1631_v33_).set(default__.safeIndex(p0, len(d_1631_v33_)), len(d_1631_v33_))): (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]})
                        d_1636_v38_ = d_1636_v38_
                        d_1637_v39_: D8
                        d_1637_v39_ = D8_DC20((self).f37)
                        (globalState).f13 = not(not((d_1637_v39_).cf31))
                        d_1638_v40_: D8
                        d_1638_v40_ = D8_DC19((self).f37, (d_1627_v29_ if (d_1637_v39_).cf31 else d_1627_v29_), (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))])
                        d_1638_v40_ = D8_DC19(False, d_1628_v30_, (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))])
                    (globalState).f1 = False
                    d_1639_v41_: C6
                    nw280_ = C6()
                    nw280_.ctor__()
                    d_1639_v41_ = nw280_
                    (globalState).f21 = p0
                    pass
            pass
        d_1640_i3_: int
        d_1640_i3_ = 0
        with _dafny.label("9"):
            while (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]:
                with _dafny.c_label("9"):
                    if (d_1640_i3_) >= (100):
                        raise _dafny.Break("9")
                    d_1640_i3_ = (d_1640_i3_) + (1)
                    if not((d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]):
                        (globalState).f0 = default__.safeDivisionInt(p0, p0)
                        d_1641_v42_: D8
                        d_1641_v42_ = D8_DC20((d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))])
                        pat_let_tv38_ = d_1583_v0_
                        pat_let_tv39_ = d_1583_v0_
                        d_1642_v43_: _dafny.Array
                        nw281_ = _dafny.Array(None, 1)
                        def iife152_(_pat_let44_0):
                            def iife153_(d_1643_dt__update__tmp_h0_):
                                def iife154_(_pat_let45_0):
                                    def iife155_(d_1644_dt__update_hcf31_h0_):
                                        return D8_DC20(d_1644_dt__update_hcf31_h0_)
                                    return iife155_(_pat_let45_0)
                                return iife154_((pat_let_tv39_)[default__.safeIndex(372, (pat_let_tv38_).length(0))])
                            return iife153_(_pat_let44_0)
                        nw281_[int(0)] = iife152_(d_1641_v42_)
                        d_1642_v43_ = nw281_
                        d_1645_v44_: _dafny.Map
                        d_1645_v44_ = _dafny.Map({(d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]: d_1642_v43_})
                        d_1646_v45_: _dafny.Map
                        d_1646_v45_ = _dafny.Map({((d_1645_v44_)[True] if (True) in (d_1645_v44_) else d_1642_v43_): (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]})
                        d_1647_v46_: _dafny.Seq
                        d_1647_v46_ = _dafny.SeqWithoutIsStrInference([(d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]])
                        d_1648_v47_: _dafny.Map
                        d_1648_v47_ = _dafny.Map({(self).f37: len(_dafny.Set({p0, p0, len(d_1647_v46_), p0, p0}))})
                        rhs269_ = d_1646_v45_
                        rhs270_ = (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]
                        rhs271_ = (d_1648_v47_) | (d_1648_v47_)
                        rhs272_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqaxjujl"))
                        lhs210_ = globalState
                        d_1646_v45_ = rhs269_
                        lhs210_.f1 = rhs270_
                        d_1648_v47_ = rhs271_
                        d_1585_v1_ = rhs272_
                        d_1649_v48_: C2
                        nw282_ = C2()
                        nw282_.ctor__((self).f37, p0)
                        d_1649_v48_ = nw282_
                        (globalState).f1 = (p0) > ((p0) * ((0) - (p0)))
                        d_1585_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1650_i4_ in range(default__.abs(445))])
                    elif True:
                        d_1651_v49_: _dafny.Seq
                        d_1651_v49_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_1652_v50_: _dafny.Seq
                        d_1652_v50_ = _dafny.SeqWithoutIsStrInference([True, (self).f37])
                        d_1653_v51_: _dafny.Map
                        d_1653_v51_ = _dafny.Map({(d_1651_v49_) + (_dafny.SeqWithoutIsStrInference([len(d_1652_v50_), p0, 410])): default__.safeModuloInt(p0, p0)})
                        d_1653_v51_ = (d_1653_v51_).set(d_1651_v49_, p0)
                        (globalState).f13 = (self).f37
                        d_1654_v52_: C5
                        nw283_ = C5()
                        nw283_.ctor__((0) - (p0))
                        d_1654_v52_ = nw283_
                        (globalState).f21 = (d_1654_v52_).f38
                        d_1655_v53_: _dafny.Array
                        nw284_ = _dafny.Array(int(0), 19)
                        d_1655_v53_ = nw284_
                        index240_ = default__.safeIndex(547, (d_1655_v53_).length(0))
                        (d_1655_v53_)[index240_] = (0) - ((d_1654_v52_).f38)
                        index241_ = default__.safeIndex(547, (d_1655_v53_).length(0))
                        (d_1655_v53_)[index241_] = 998
                    d_1656_v54_: str
                    d_1656_v54_ = _dafny.CodePoint('a')
                    d_1585_v1_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtqnx"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtqnx")))), d_1656_v54_)) + (d_1585_v1_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynoheu")) if (self).f37 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "othhkkvx"))))
                    (globalState).f21 = 259
                    d_1657_v56_: _dafny.MultiSet
                    def iife156_():
                        coll64_ = _dafny.Map()
                        compr_64_: int
                        for compr_64_ in _dafny.IntegerRange(73, 159):
                            d_1658_v55_: int = compr_64_
                            if ((73) <= (d_1658_v55_)) and ((d_1658_v55_) < (159)):
                                coll64_[default__.safeModuloInt(d_1658_v55_, p0)] = (self).f37
                        return _dafny.Map(coll64_)
                    d_1657_v56_ = _dafny.MultiSet([len(iife156_()
                    )])
                    d_1659_v57_: _dafny.Map
                    d_1659_v57_ = _dafny.Map({(d_1657_v56_).ispropersubset(default__.fm43(p0, len(_dafny.Map({(self).f37: True})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mf")), globalState)): (p0) - (p0)})
                    d_1660_v58_: _dafny.Set
                    d_1660_v58_ = _dafny.Set({(d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]})
                    d_1659_v57_ = (d_1659_v57_).set((self).f37, len(d_1660_v58_))
                    pass
            pass
        r0 = (d_1583_v0_)[default__.safeIndex(372, (d_1583_v0_).length(0))]
        return r0

    @property
    def f37(self):
        return self._f37

class C8(T1, T0):
    def  __init__(self):
        self._f26: _dafny.Map = _dafny.Map({})
        self.f35: bool = False
        self._f36: D5 = D5.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f26(self):
        return self._f26
    @f26.setter
    def f26(self, value):
        self._f26 = value
    def ctor__(self, f35, f36, f26):
        (self).f35 = f35
        (self)._f36 = f36
        (self).f26 = f26

    def fm8(self, p0, globalState):
        def iife157_():
            coll65_ = _dafny.Map()
            compr_65_: int
            for compr_65_ in _dafny.IntegerRange(928, 493):
                d_1661_v0_: int = compr_65_
                if ((928) <= (d_1661_v0_)) and ((d_1661_v0_) < (493)):
                    coll65_[default__.safeModuloInt(d_1661_v0_, len(_dafny.Map({self.f35: _dafny.MultiSet([self.f35])})))] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tf"))])))
            return _dafny.Map(coll65_)
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(iife157_()
        ), (0) - (len(_dafny.Map({self.f35: 270}))), 751])): _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({self.f35: self.f35})]))})})

    def fm3(self, globalState):
        return (len(_dafny.Set({False}))) + (default__.safeDivisionInt((0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f35]))).cardinality), 161))

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife158_():
            def iife160_():
                coll68_ = _dafny.Map()
                compr_68_: _dafny.Seq
                for compr_68_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, self.f35]): 430})).keys.Elements:
                    d_1662_v1_: _dafny.Seq = compr_68_
                    if (d_1662_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, self.f35]): 430})):
                        coll68_[d_1662_v1_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')])
                return _dafny.Map(coll68_)
            coll66_ = _dafny.Map()
            def iife159_():
                coll67_ = _dafny.Map()
                compr_67_: _dafny.Seq
                for compr_67_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, self.f35]): 430})).keys.Elements:
                    d_1662_v1_: _dafny.Seq = compr_67_
                    if (d_1662_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, self.f35]): 430})):
                        coll67_[d_1662_v1_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')])
                return _dafny.Map(coll67_)
            compr_66_: _dafny.Seq
            for compr_66_ in (iife159_()
            ).keys.Elements:
                d_1663_v0_: _dafny.Seq = compr_66_
                if (d_1663_v0_) in (iife160_()
                ):
                    coll66_[d_1663_v0_] = self.f35
            return _dafny.Map(coll66_)
        return (_dafny.Map({len(_dafny.Map({D3_DC6(iife158_()
): 652})): _dafny.Map({_dafny.CodePoint('n'): 819})})) | (_dafny.Map({713: _dafny.Map({_dafny.CodePoint('p'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kg")))})}))

    def fm12(self, globalState):
        def iife161_():
            coll69_ = _dafny.Map()
            compr_69_: int
            for compr_69_ in _dafny.IntegerRange(462, 670):
                d_1664_v0_: int = compr_69_
                if ((462) <= (d_1664_v0_)) and ((d_1664_v0_) < (670)):
                    coll69_[default__.safeModuloInt(d_1664_v0_, (0) - ((0) - ((0) - ((0) - (-923)))))] = _dafny.CodePoint('i')
            return _dafny.Map(coll69_)
        def iife162_():
            coll70_ = _dafny.Set()
            compr_70_: int
            for compr_70_ in (_dafny.Map({988: self.f35})).keys.Elements:
                d_1665_v1_: int = compr_70_
                if (d_1665_v1_) in (_dafny.Map({988: self.f35})):
                    coll70_ = coll70_.union(_dafny.Set([default__.safeDivisionInt(d_1665_v1_, len(_dafny.Map({True: 607})))]))
            return _dafny.Set(coll70_)
        source23_ = (D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebngtivx")), len(iife161_()
), len(iife162_()
), 612, _dafny.SeqWithoutIsStrInference([self.f35])) if self.f35 else D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvpq")), len(_dafny.Set({999, len(_dafny.SeqWithoutIsStrInference([self.f35, self.f35, self.f35, self.f35, self.f35]))})), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f35]))).cardinality, (0) - (len(_dafny.Map({self.f35: self.f35}))), _dafny.SeqWithoutIsStrInference([self.f35, self.f35])))
        if source23_.is_DC7:
            d_1666___mcc_h0_ = source23_.cf8
            d_1667___mcc_h1_ = source23_.cf9
            d_1668___mcc_h2_ = source23_.cf10
            d_1669___mcc_h3_ = source23_.cf11
            d_1670___mcc_h4_ = source23_.cf12
            d_1671_cf12_ = d_1670___mcc_h4_
            d_1672_cf11_ = d_1669___mcc_h3_
            d_1673_cf10_ = d_1668___mcc_h2_
            d_1674_cf9_ = d_1667___mcc_h1_
            d_1675_cf8_ = d_1666___mcc_h0_
            if self.f35:
                return D2_DC5()
            elif True:
                return D2_DC5()
        elif True:
            d_1676___mcc_h5_ = source23_.cf7
            d_1677_cf7_ = d_1676___mcc_h5_
            return D2_DC5()

    def m7(self, p0, globalState):
        r0: bool = False
        (globalState).f22 = p0
        d_1678_v0_: _dafny.Array
        nw285_ = _dafny.Array(False, 12)
        d_1678_v0_ = nw285_
        index242_ = default__.safeIndex(789, (d_1678_v0_).length(0))
        (d_1678_v0_)[index242_] = False
        index243_ = default__.safeIndex(789, (d_1678_v0_).length(0))
        (d_1678_v0_)[index243_] = self.f35
        d_1678_v0_ = d_1678_v0_
        d_1679_v1_: _dafny.Set
        d_1679_v1_ = _dafny.Set({p0, (p0) * (-194)})
        d_1680_v2_: _dafny.Seq
        d_1680_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ns"))
        d_1681_v3_: D4
        d_1681_v3_ = D4_DC10(len(d_1680_v2_))
        pat_let_tv40_ = d_1679_v1_
        pat_let_tv41_ = d_1679_v1_
        pat_let_tv42_ = d_1679_v1_
        pat_let_tv43_ = p0
        def lambda81_(source24_):
            if source24_.is_DC9:
                d_1682___mcc_h0_ = source24_.cf14
                d_1683_cf14_ = d_1682___mcc_h0_
                return (pat_let_tv40_).intersection(pat_let_tv41_)
            elif source24_.is_DC10:
                d_1684___mcc_h1_ = source24_.cf15
                d_1685_cf15_ = d_1684___mcc_h1_
                return pat_let_tv42_
            elif True:
                d_1686___mcc_h2_ = source24_.cf13
                d_1687_cf13_ = d_1686___mcc_h2_
                return _dafny.Set({pat_let_tv43_})

        d_1679_v1_ = lambda81_(d_1681_v3_)
        (globalState).f0 = default__.fm2(globalState)
        d_1688_i0_: int
        d_1688_i0_ = 0
        with _dafny.label("10"):
            while (d_1678_v0_)[default__.safeIndex(789, (d_1678_v0_).length(0))]:
                with _dafny.c_label("10"):
                    if (d_1688_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1688_i0_ = (d_1688_i0_) + (1)
                    (globalState).f1 = (self.f35) == (not(self.f35))
                    (globalState).f1 = self.f35
                    d_1689_v4_: _dafny.Seq
                    d_1689_v4_ = _dafny.SeqWithoutIsStrInference([p0])
                    (globalState).f21 = ((0) - ((0) - ((p0) + (p0)))) + (len((d_1689_v4_) + ((_dafny.SeqWithoutIsStrInference([p0])))))
                    d_1690_v5_: D5
                    d_1690_v5_ = D5_DC11((d_1679_v1_).issubset(d_1679_v1_))
                    pat_let_tv44_ = d_1678_v0_
                    pat_let_tv45_ = d_1678_v0_
                    def iife164_(_pat_let47_0):
                        def iife165_(d_1691_dt__update__tmp_h0_):
                            def iife166_(_pat_let48_0):
                                def iife167_(d_1692_dt__update_hcf16_h0_):
                                    return D5_DC11(d_1692_dt__update_hcf16_h0_)
                                return iife167_(_pat_let48_0)
                            return iife166_((pat_let_tv45_)[default__.safeIndex(789, (pat_let_tv44_).length(0))])
                        return iife165_(_pat_let47_0)
                    def iife163_(_pat_let46_0):
                        def iife168_(d_1693_dt__update__tmp_h1_):
                            def iife169_(_pat_let49_0):
                                def iife170_(d_1694_dt__update_hcf16_h1_):
                                    return D5_DC11(d_1694_dt__update_hcf16_h1_)
                                return iife170_(_pat_let49_0)
                            return iife169_(self.f35)
                        return iife168_(_pat_let46_0)
                    d_1690_v5_ = iife163_(iife164_((self).f36))
                    pass
            pass
        pat_let_tv46_ = p0
        pat_let_tv47_ = p0
        pat_let_tv48_ = p0
        pat_let_tv49_ = p0
        pat_let_tv50_ = d_1678_v0_
        pat_let_tv51_ = d_1678_v0_
        def lambda82_(source25_):
            if source25_.is_DC12:
                d_1695___mcc_h3_ = source25_.cf17
                d_1696___mcc_h4_ = source25_.cf18
                d_1697_cf18_ = d_1696___mcc_h4_
                d_1698_cf17_ = d_1695___mcc_h3_
                return (_dafny.SeqWithoutIsStrInference([pat_let_tv46_ for d_1699_i1_ in range(default__.abs(-367))])) not in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([pat_let_tv47_ for d_1700_i2_ in range(default__.abs(750))]), _dafny.SeqWithoutIsStrInference([pat_let_tv48_, 445]), _dafny.SeqWithoutIsStrInference([755]), _dafny.SeqWithoutIsStrInference([pat_let_tv49_])])))
            elif source25_.is_DC13:
                d_1701___mcc_h5_ = source25_.cf19
                d_1702_cf19_ = d_1701___mcc_h5_
                return (pat_let_tv51_)[default__.safeIndex(789, (pat_let_tv50_).length(0))]
            elif True:
                d_1703___mcc_h6_ = source25_.cf16
                d_1704_cf16_ = d_1703___mcc_h6_
                return d_1704_cf16_

        r0 = lambda82_((self).f36)
        return r0

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        d_1705_v0_: _dafny.Set
        d_1705_v0_ = _dafny.Set({self.f35})
        d_1706_v1_: _dafny.Array
        nw286_ = _dafny.Array(None, 5)
        nw286_[int(0)] = p1
        nw286_[int(1)] = p2
        nw286_[int(2)] = p1
        nw286_[int(3)] = -186
        nw286_[int(4)] = len(d_1705_v0_)
        d_1706_v1_ = nw286_
        d_1707_v2_: _dafny.MultiSet
        d_1707_v2_ = _dafny.MultiSet([self.f35])
        index244_ = default__.safeIndex(209, (d_1706_v1_).length(0))
        (d_1706_v1_)[index244_] = (d_1707_v2_).cardinality
        d_1708_v3_: str
        d_1708_v3_ = _dafny.CodePoint('j')
        d_1709_v4_: _dafny.MultiSet
        d_1709_v4_ = _dafny.MultiSet([p2])
        d_1710_v5_: _dafny.Map
        d_1710_v5_ = _dafny.Map({self.f35: self.f35})
        d_1711_v6_: _dafny.Seq
        d_1711_v6_ = _dafny.SeqWithoutIsStrInference([((d_1710_v5_)[self.f35] if (self.f35) in (d_1710_v5_) else True), self.f35])
        d_1712_v7_: _dafny.Map
        d_1712_v7_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebkd"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebkd")))), d_1708_v3_)): (d_1709_v4_) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([p2, p2, p2])), len(d_1711_v6_)]))})
        index245_ = default__.safeIndex(209, (d_1706_v1_).length(0))
        rhs273_ = (p1) != ((len(d_1711_v6_) if self.f35 else p2))
        rhs274_ = p1
        rhs275_ = len(default__.fm13(len(d_1710_v5_), globalState))
        rhs276_ = p2
        rhs277_ = d_1712_v7_
        lhs211_ = self
        lhs212_ = d_1706_v1_
        lhs213_ = default__.safeIndex(209, (d_1706_v1_).length(0))
        lhs214_ = globalState
        lhs211_.f35 = rhs273_
        lhs212_[lhs213_] = rhs274_
        r2 = rhs275_
        lhs214_.f16 = rhs276_
        d_1712_v7_ = rhs277_
        d_1713_v8_: _dafny.Seq
        d_1713_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmxywdk"))
        (self).f26 = (self.f26).set(len((d_1713_v8_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psxknie")))), (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))])
        hi7_ = p1
        for d_1714_i0_ in range(p2, hi7_):
            nw287_ = _dafny.Array(int(0), 5)
            d_1706_v1_ = nw287_
            d_1715_v9_: _dafny.Set
            d_1715_v9_ = _dafny.Set({D0_DC1(d_1706_v1_, d_1708_v3_, d_1714_i0_)})
            (globalState).f13 = not((d_1715_v9_).isdisjoint((d_1715_v9_) - (d_1715_v9_)))
            d_1716_v11_: _dafny.Map
            d_1716_v11_ = _dafny.Map({d_1708_v3_: p2})
            d_1717_v12_: _dafny.Seq
            def iife171_():
                coll71_ = _dafny.Map()
                compr_71_: str
                for compr_71_ in (d_1716_v11_).keys.Elements:
                    d_1718_v10_: str = compr_71_
                    if (d_1718_v10_) in (d_1716_v11_):
                        coll71_[d_1718_v10_] = (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]
                return _dafny.Map(coll71_)
            d_1717_v12_ = _dafny.SeqWithoutIsStrInference([default__.fm14(globalState), iife171_()
            ])
            d_1717_v12_ = d_1717_v12_
            d_1719_v13_: _dafny.Set
            d_1719_v13_ = _dafny.Set({d_1708_v3_})
            d_1720_v14_: _dafny.Map
            d_1720_v14_ = _dafny.Map({334: d_1706_v1_})
            index246_ = default__.safeIndex(209, (d_1706_v1_).length(0))
            rhs278_ = (d_1719_v13_) - (default__.fm15(self.f35, globalState))
            rhs279_ = (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]
            rhs280_ = (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]
            rhs281_ = ((-365) - (d_1714_i0_)) in (d_1720_v14_)
            lhs215_ = d_1706_v1_
            lhs216_ = default__.safeIndex(209, (d_1706_v1_).length(0))
            lhs217_ = globalState
            lhs218_ = self
            d_1719_v13_ = rhs278_
            lhs215_[lhs216_] = rhs279_
            lhs217_.f21 = rhs280_
            lhs218_.f35 = rhs281_
        d_1721_v15_: D5
        d_1721_v15_ = D5_DC12(self.f35, (_dafny.SeqWithoutIsStrInference([d_1708_v3_ for d_1722_i1_ in range(default__.abs(239))])) <= (d_1713_v8_))
        d_1721_v15_ = d_1721_v15_
        if (self.f35 if self.f35 else not(self.f35)):
            d_1723_v16_: _dafny.Array
            nw288_ = _dafny.Array(False, 4)
            d_1723_v16_ = nw288_
            index247_ = default__.safeIndex(417, (d_1723_v16_).length(0))
            (d_1723_v16_)[index247_] = self.f35
            d_1724_v17_: D4
            d_1724_v17_ = D4_DC10((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))])
            index248_ = default__.safeIndex(417, (d_1723_v16_).length(0))
            (d_1723_v16_)[index248_] = (((0) - ((0) - (p1))) - ((d_1724_v17_).cf15)) < (p2)
            d_1725_v18_: D3
            d_1725_v18_ = D3_DC6(_dafny.Map({d_1711_v6_: (d_1723_v16_)[default__.safeIndex(417, (d_1723_v16_).length(0))]}))
            source26_ = (d_1725_v18_ if self.f35 else d_1725_v18_)
            if source26_.is_DC7:
                d_1726___mcc_h0_ = source26_.cf8
                d_1727___mcc_h1_ = source26_.cf9
                d_1728___mcc_h2_ = source26_.cf10
                d_1729___mcc_h3_ = source26_.cf11
                d_1730___mcc_h4_ = source26_.cf12
                d_1731_cf12_ = d_1730___mcc_h4_
                d_1732_cf11_ = d_1729___mcc_h3_
                d_1733_cf10_ = d_1728___mcc_h2_
                d_1734_cf9_ = d_1727___mcc_h1_
                d_1735_cf8_ = d_1726___mcc_h0_
                index249_ = default__.safeIndex(417, (d_1723_v16_).length(0))
                (d_1723_v16_)[index249_] = not((_dafny.SeqWithoutIsStrInference([self.f35, self.f35, (d_1723_v16_)[default__.safeIndex(417, (d_1723_v16_).length(0))], not(self.f35)])) < ((d_1731_cf12_) + (d_1731_cf12_)))
                (globalState).f1 = (d_1723_v16_)[default__.safeIndex(417, (d_1723_v16_).length(0))]
                d_1736_v19_: _dafny.Array
                nw289_ = _dafny.Array(_dafny.Set({}), 20)
                d_1736_v19_ = nw289_
                d_1737_v20_: D7
                d_1737_v20_ = D7_DC15(d_1736_v19_)
                d_1738_v21_: _dafny.Map
                d_1738_v21_ = _dafny.Map({(d_1737_v20_).cf21: d_1733_cf10_})
                d_1738_v21_ = (d_1738_v21_).set(d_1736_v19_, d_1732_cf11_)
                d_1734_cf9_ = (0) - (default__.fm2(globalState))
            elif True:
                d_1739___mcc_h5_ = source26_.cf7
                d_1740_cf7_ = d_1739___mcc_h5_
                d_1741_v22_: _dafny.Array
                def lambda83_(d_1742_v8_):
                    def lambda84_(d_1743_i2_):
                        return d_1742_v8_

                    return lambda84_

                init47_ = lambda83_(d_1713_v8_)
                nw290_ = _dafny.Array(None, 3)
                for i0_47_ in range(nw290_.length(0)):
                    nw290_[i0_47_] = init47_(i0_47_)
                d_1741_v22_ = nw290_
                index250_ = default__.safeIndex(383, (d_1741_v22_).length(0))
                (d_1741_v22_)[index250_] = (d_1713_v8_) + (d_1713_v8_)
                index251_ = default__.safeIndex(383, (d_1741_v22_).length(0))
                (d_1741_v22_)[index251_] = (d_1713_v8_) + ((d_1713_v8_) + (d_1713_v8_))
                (self).f35 = not ((self.f35) and ((d_1723_v16_)[default__.safeIndex(417, (d_1723_v16_).length(0))])) or ((d_1723_v16_)[default__.safeIndex(417, (d_1723_v16_).length(0))])
                (globalState).f2 = default__.fm2(globalState)
                (globalState).f1 = not (self.f35) or ((d_1723_v16_)[default__.safeIndex(417, (d_1723_v16_).length(0))])
            (self).f35 = (d_1723_v16_)[default__.safeIndex(417, (d_1723_v16_).length(0))]
            d_1705_v0_ = d_1705_v0_
            (globalState).f22 = p1
        elif True:
            d_1744_v23_: bool
            d_1745_v24_: bool
            d_1746_v25_: _dafny.Map
            d_1747_v26_: int
            out63_: bool
            out64_: bool
            out65_: _dafny.Map
            out66_: int
            out63_, out64_, out65_, out66_ = default__.m0(globalState)
            d_1744_v23_ = out63_
            d_1745_v24_ = out64_
            d_1746_v25_ = out65_
            d_1747_v26_ = out66_
            d_1748_v27_: D4
            d_1748_v27_ = D4_DC9((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))])
            index252_ = default__.safeIndex(209, (d_1706_v1_).length(0))
            (d_1706_v1_)[index252_] = (default__.safeDivisionInt(p1, p1)) * ((d_1748_v27_).cf14)
            d_1749_v28_: _dafny.Array
            nw291_ = _dafny.Array(None, 21)
            nw291_[int(0)] = self.f35
            nw291_[int(1)] = (d_1709_v4_).issubset(d_1709_v4_)
            nw291_[int(2)] = (d_1747_v26_) == (231)
            nw291_[int(3)] = True
            nw291_[int(4)] = self.f35
            nw291_[int(5)] = d_1745_v24_
            nw291_[int(6)] = d_1744_v23_
            nw291_[int(7)] = d_1744_v23_
            nw291_[int(8)] = self.f35
            nw291_[int(9)] = d_1744_v23_
            nw291_[int(10)] = (D7_DC16(((self.f26)[p1] if (p1) in (self.f26) else (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]), d_1745_v24_)).cf23
            nw291_[int(11)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjluddw")))) <= (d_1747_v26_)
            nw291_[int(12)] = False
            nw291_[int(13)] = ((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]) == (d_1747_v26_)
            nw291_[int(14)] = d_1744_v23_
            nw291_[int(15)] = d_1745_v24_
            nw291_[int(16)] = d_1745_v24_
            nw291_[int(17)] = ((d_1710_v5_)[d_1744_v23_] if (d_1744_v23_) in (d_1710_v5_) else d_1745_v24_)
            nw291_[int(18)] = not((len(d_1711_v6_)) != (len(self.f26)))
            nw291_[int(19)] = d_1744_v23_
            nw291_[int(20)] = d_1745_v24_
            d_1749_v28_ = nw291_
            index253_ = default__.safeIndex(812, (d_1749_v28_).length(0))
            (d_1749_v28_)[index253_] = (d_1744_v23_ if d_1744_v23_ else d_1745_v24_)
            index254_ = default__.safeIndex(812, (d_1749_v28_).length(0))
            (d_1749_v28_)[index254_] = d_1745_v24_
            (globalState).f21 = (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]
            d_1750_v29_: _dafny.Set
            d_1750_v29_ = _dafny.Set({p1})
            d_1751_v30_: _dafny.Map
            d_1751_v30_ = _dafny.Map({self.f35: d_1709_v4_})
            r0 = default__.safeDivisionInt((p2 if (d_1749_v28_)[default__.safeIndex(812, (d_1749_v28_).length(0))] else len(d_1711_v6_)), default__.safeModuloInt(len(d_1750_v29_), len(d_1751_v30_)))
        if default__.fm0(p1, globalState):
            if self.f35:
                (globalState).f1 = self.f35
                d_1752_v31_: D7
                d_1752_v31_ = D7_DC16(default__.safeModuloInt(635, (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]), ((self).fm3(globalState)) >= (96))
                d_1753_v32_: _dafny.Seq
                d_1753_v32_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1754_v33_: _dafny.Set
                d_1754_v33_ = _dafny.Set({d_1753_v32_, d_1753_v32_, d_1753_v32_})
                d_1755_v34_: _dafny.Map
                d_1755_v34_ = _dafny.Map({d_1713_v8_: _dafny.Set({d_1753_v32_})})
                rhs282_ = d_1752_v31_
                rhs283_ = (d_1754_v33_).ispropersubset(((d_1755_v34_)[d_1713_v8_] if (d_1713_v8_) in (d_1755_v34_) else default__.fm16((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))], self.f35, (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))], globalState)))
                lhs219_ = globalState
                d_1752_v31_ = rhs282_
                lhs219_.f13 = rhs283_
                d_1756_v35_: _dafny.Set
                d_1756_v35_ = _dafny.Set({p2})
                (globalState).f22 = len(d_1756_v35_)
                (globalState).f16 = default__.safeDivisionInt(default__.safeDivisionInt(p2, p2), p1)
                d_1757_v36_: _dafny.Array
                def lambda85_(d_1758_v3_):
                    def lambda86_(d_1759_i3_):
                        return d_1758_v3_

                    return lambda86_

                init48_ = lambda85_(d_1708_v3_)
                nw292_ = _dafny.Array(None, 14)
                for i0_48_ in range(nw292_.length(0)):
                    nw292_[i0_48_] = init48_(i0_48_)
                d_1757_v36_ = nw292_
                nw293_ = _dafny.Array(_dafny.CodePoint('D'), 23)
                d_1757_v36_ = nw293_
            elif True:
                d_1760_v38_: _dafny.Seq
                def iife172_():
                    coll72_ = _dafny.Map()
                    compr_72_: int
                    for compr_72_ in _dafny.IntegerRange(156, 921):
                        d_1761_v37_: int = compr_72_
                        if ((156) <= (d_1761_v37_)) and ((d_1761_v37_) < (921)):
                            coll72_[(d_1761_v37_) + (((d_1709_v4_)[p1] if (p1) in (d_1709_v4_) else (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]))] = self.f35
                    return _dafny.Map(coll72_)
                d_1760_v38_ = _dafny.SeqWithoutIsStrInference([iife172_()
                ])
                d_1760_v38_ = d_1760_v38_
                (globalState).f16 = (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]
                (globalState).f1 = self.f35
                d_1762_v39_: _dafny.Array
                nw294_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_1762_v39_ = nw294_
                index255_ = default__.safeIndex(98, (d_1762_v39_).length(0))
                (d_1762_v39_)[index255_] = d_1708_v3_
                d_1763_v40_: _dafny.Array
                nw295_ = _dafny.Array(False, 14)
                d_1763_v40_ = nw295_
                index256_ = default__.safeIndex(981, (d_1763_v40_).length(0))
                (d_1763_v40_)[index256_] = self.f35
                d_1764_v41_: _dafny.Array
                nw296_ = _dafny.Array(_dafny.MultiSet({}), 28)
                d_1764_v41_ = nw296_
                d_1765_v42_: D2
                d_1765_v42_ = D2_DC4(d_1764_v41_)
                d_1766_v43_: _dafny.Map
                d_1766_v43_ = _dafny.Map({775: self.f35})
                d_1767_v44_: _dafny.Map
                d_1767_v44_ = _dafny.Map({d_1765_v42_: d_1766_v43_})
                index257_ = default__.safeIndex(98, (d_1762_v39_).length(0))
                index258_ = default__.safeIndex(209, (d_1706_v1_).length(0))
                index259_ = default__.safeIndex(981, (d_1763_v40_).length(0))
                rhs284_ = _dafny.CodePoint('b')
                rhs285_ = default__.safeDivisionInt(p1, ((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]) - (p2))
                rhs286_ = (0) - (len(d_1767_v44_))
                rhs287_ = self.f35
                lhs220_ = d_1762_v39_
                lhs221_ = default__.safeIndex(98, (d_1762_v39_).length(0))
                lhs222_ = globalState
                lhs223_ = d_1706_v1_
                lhs224_ = default__.safeIndex(209, (d_1706_v1_).length(0))
                lhs225_ = d_1763_v40_
                lhs226_ = default__.safeIndex(981, (d_1763_v40_).length(0))
                lhs220_[lhs221_] = rhs284_
                lhs222_.f5 = rhs285_
                lhs223_[lhs224_] = rhs286_
                lhs225_[lhs226_] = rhs287_
                (globalState).f13 = self.f35
            (globalState).f22 = (0) - (((d_1707_v2_)[self.f35] if (self.f35) in (d_1707_v2_) else len(d_1710_v5_)))
            d_1768_v45_: _dafny.Seq
            d_1768_v45_ = _dafny.SeqWithoutIsStrInference([(default__.fm18(self.f35, globalState)).cardinality, p2, p2, p2, 898])
            d_1769_v46_: _dafny.Set
            d_1769_v46_ = _dafny.Set({p1, p2})
            d_1705_v0_ = default__.fm17(d_1768_v45_, self.f35, (d_1769_v46_).intersection(d_1769_v46_), globalState)
            d_1770_v47_: _dafny.Array
            nw297_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_1770_v47_ = nw297_
            d_1771_v48_: _dafny.Array
            nw298_ = _dafny.Array(_dafny.CodePoint('D'), 13)
            d_1771_v48_ = nw298_
            index260_ = default__.safeIndex(534, (d_1770_v47_).length(0))
            (d_1770_v47_)[index260_] = d_1771_v48_
            d_1772_v49_: D2
            d_1772_v49_ = D2_DC5()
            index261_ = default__.safeIndex(534, (d_1770_v47_).length(0))
            rhs288_ = (p2) - ((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))])
            rhs289_ = d_1771_v48_
            rhs290_ = d_1710_v5_
            rhs291_ = (self).fm12(globalState)
            rhs292_ = 352
            lhs227_ = globalState
            lhs228_ = d_1770_v47_
            lhs229_ = default__.safeIndex(534, (d_1770_v47_).length(0))
            lhs227_.f21 = rhs288_
            lhs228_[lhs229_] = rhs289_
            d_1710_v5_ = rhs290_
            d_1772_v49_ = rhs291_
            r0 = rhs292_
            d_1773_v50_: _dafny.MultiSet
            d_1773_v50_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbmjqav")), d_1713_v8_, d_1713_v8_])
            d_1774_v51_: C5
            nw299_ = C5()
            nw299_.ctor__(((d_1773_v50_)[d_1713_v8_] if (d_1713_v8_) in (d_1773_v50_) else p2))
            d_1774_v51_ = nw299_
        elif True:
            index262_ = default__.safeIndex(209, (d_1706_v1_).length(0))
            (d_1706_v1_)[index262_] = p1
            if True:
                d_1775_v52_: _dafny.Map
                d_1775_v52_ = _dafny.Map({self.f35: d_1713_v8_})
                (globalState).f0 = default__.safeModuloInt((0) - ((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]), (len(d_1775_v52_)) + (p1))
                (globalState).f13 = self.f35
                d_1776_v53_: _dafny.Seq
                d_1776_v53_ = _dafny.SeqWithoutIsStrInference([(d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]])
                (self).f35 = ((d_1776_v53_)[default__.safeIndex((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))], len(d_1776_v53_))]) != (p2)
                d_1777_v54_: _dafny.Array
                nw300_ = _dafny.Array(None, 19)
                nw300_[int(0)] = (not(not(self.f35))) and (True)
                nw300_[int(1)] = True
                nw300_[int(2)] = self.f35
                nw300_[int(3)] = self.f35
                nw300_[int(4)] = self.f35
                nw300_[int(5)] = self.f35
                nw300_[int(6)] = self.f35
                nw300_[int(7)] = self.f35
                nw300_[int(8)] = self.f35
                nw300_[int(9)] = self.f35
                nw300_[int(10)] = (d_1721_v15_).cf18
                nw300_[int(11)] = self.f35
                nw300_[int(12)] = self.f35
                nw300_[int(13)] = self.f35
                nw300_[int(14)] = (d_1711_v6_)[default__.safeIndex(p2, len(d_1711_v6_))]
                nw300_[int(15)] = self.f35
                nw300_[int(16)] = self.f35
                nw300_[int(17)] = self.f35
                nw300_[int(18)] = not(self.f35)
                d_1777_v54_ = nw300_
                d_1778_v55_: _dafny.Seq
                d_1778_v55_ = _dafny.SeqWithoutIsStrInference([d_1713_v8_, (d_1713_v8_).set(default__.safeIndex(default__.fm2(globalState), len(d_1713_v8_)), (d_1713_v8_)[default__.safeIndex(p1, len(d_1713_v8_))])])
                index263_ = default__.safeIndex(209, (d_1706_v1_).length(0))
                rhs293_ = ((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]) + (876)
                rhs294_ = (d_1777_v54_ if self.f35 else d_1777_v54_)
                rhs295_ = (d_1778_v55_)[default__.safeIndex(p2, len(d_1778_v55_))]
                rhs296_ = self.f35
                rhs297_ = d_1777_v54_
                lhs230_ = d_1706_v1_
                lhs231_ = default__.safeIndex(209, (d_1706_v1_).length(0))
                lhs232_ = globalState
                lhs230_[lhs231_] = rhs293_
                d_1777_v54_ = rhs294_
                d_1713_v8_ = rhs295_
                lhs232_.f13 = rhs296_
                d_1777_v54_ = rhs297_
                (globalState).f1 = self.f35
            elif True:
                index264_ = default__.safeIndex(209, (d_1706_v1_).length(0))
                (d_1706_v1_)[index264_] = (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]
                d_1779_v56_: bool
                d_1780_v57_: bool
                d_1781_v58_: _dafny.Map
                d_1782_v59_: int
                out67_: bool
                out68_: bool
                out69_: _dafny.Map
                out70_: int
                out67_, out68_, out69_, out70_ = default__.m0(globalState)
                d_1779_v56_ = out67_
                d_1780_v57_ = out68_
                d_1781_v58_ = out69_
                d_1782_v59_ = out70_
                d_1713_v8_ = d_1713_v8_
                d_1783_v60_: _dafny.Array
                nw301_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_1783_v60_ = nw301_
                d_1784_v61_: _dafny.Map
                d_1784_v61_ = _dafny.Map({d_1783_v60_: d_1708_v3_})
                d_1784_v61_ = (d_1784_v61_).set(d_1783_v60_, d_1708_v3_)
                (globalState).f2 = p2
            (self).f35 = ((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]) >= ((d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))])
            (globalState).f1 = self.f35
            d_1785_v62_: _dafny.Map
            d_1785_v62_ = _dafny.Map({(d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]: self.f35})
            d_1786_v63_: _dafny.Array
            nw302_ = _dafny.Array(None, 10)
            nw302_[int(0)] = ((d_1785_v62_)[p2] if (p2) in (d_1785_v62_) else self.f35)
            nw302_[int(1)] = self.f35
            nw302_[int(2)] = default__.fm0(p1, globalState)
            nw302_[int(3)] = self.f35
            nw302_[int(4)] = self.f35
            nw302_[int(5)] = self.f35
            nw302_[int(6)] = self.f35
            nw302_[int(7)] = self.f35
            nw302_[int(8)] = self.f35
            nw302_[int(9)] = self.f35
            d_1786_v63_ = nw302_
            d_1787_v64_: D11
            d_1787_v64_ = D11_DC26(d_1786_v63_)
            d_1786_v63_ = (d_1787_v64_).cf39
        r0 = default__.fm2(globalState)
        d_1788_v65_: _dafny.Array
        nw303_ = _dafny.Array(False, 28)
        d_1788_v65_ = nw303_
        d_1789_v66_: _dafny.Seq
        d_1789_v66_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")): d_1788_v65_})])
        d_1790_v67_: _dafny.Map
        d_1790_v67_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1708_v3_ for d_1791_i4_ in range(default__.abs(174))]): d_1788_v65_})
        r1 = ((d_1789_v66_)[default__.safeIndex(p1, len(d_1789_v66_))]) | ((d_1790_v67_) | (d_1790_v67_))
        r2 = (d_1706_v1_)[default__.safeIndex(209, (d_1706_v1_).length(0))]
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        d_1792_v0_: _dafny.Array
        nw304_ = _dafny.Array(int(0), 2)
        d_1792_v0_ = nw304_
        d_1793_v1_: int
        d_1793_v1_ = 945
        d_1794_v2_: D8
        d_1794_v2_ = D8_DC18(_dafny.Set({d_1793_v1_}))
        d_1795_v3_: _dafny.Seq
        d_1795_v3_ = _dafny.SeqWithoutIsStrInference([self.f35])
        d_1796_v4_: _dafny.Map
        d_1796_v4_ = _dafny.Map({d_1795_v3_: False})
        pat_let_tv52_ = d_1796_v4_
        pat_let_tv53_ = d_1793_v1_
        index265_ = default__.safeIndex(266, (d_1792_v0_).length(0))
        def iife173_(_pat_let50_0):
            def iife174_(d_1797_dt__update__tmp_h0_):
                def iife175_(_pat_let51_0):
                    def iife176_(d_1798_dt__update_hcf27_h0_):
                        return D8_DC18(d_1798_dt__update_hcf27_h0_)
                    return iife176_(_pat_let51_0)
                return iife175_(_dafny.Set({len(pat_let_tv52_), pat_let_tv53_}))
            return iife174_(_pat_let50_0)
        (d_1792_v0_)[index265_] = len((iife173_(d_1794_v2_)).cf27)
        index266_ = default__.safeIndex(266, (d_1792_v0_).length(0))
        (d_1792_v0_)[index266_] = default__.safeModuloInt(d_1793_v1_, (d_1793_v1_) * (d_1793_v1_))
        d_1799_v5_: _dafny.Array
        def lambda87_(d_1800_i0_):
            return self.f35

        init49_ = lambda87_
        nw305_ = _dafny.Array(None, 13)
        for i0_49_ in range(nw305_.length(0)):
            nw305_[i0_49_] = init49_(i0_49_)
        d_1799_v5_ = nw305_
        index267_ = default__.safeIndex(566, (d_1799_v5_).length(0))
        (d_1799_v5_)[index267_] = self.f35
        index268_ = default__.safeIndex(566, (d_1799_v5_).length(0))
        (d_1799_v5_)[index268_] = self.f35
        d_1801_v6_: _dafny.Seq
        d_1801_v6_ = _dafny.SeqWithoutIsStrInference([839])
        hi8_ = d_1793_v1_
        for d_1802_i1_ in range((d_1793_v1_) + (len(d_1801_v6_)), hi8_):
            d_1803_v7_: _dafny.Array
            def lambda88_(d_1804_i2_):
                return _dafny.CodePoint('t')

            init50_ = lambda88_
            nw306_ = _dafny.Array(None, 6)
            for i0_50_ in range(nw306_.length(0)):
                nw306_[i0_50_] = init50_(i0_50_)
            d_1803_v7_ = nw306_
            d_1805_v8_: C0
            nw307_ = C0()
            nw307_.ctor__(d_1803_v7_)
            d_1805_v8_ = nw307_
            d_1806_v9_: _dafny.MultiSet
            d_1806_v9_ = _dafny.MultiSet([(d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))]])
            (self).f35 = (self.f35) == ((_dafny.MultiSet(d_1795_v3_)).issubset(d_1806_v9_))
            d_1807_v10_: _dafny.Set
            d_1807_v10_ = _dafny.Set({self.f35, (True) or (self.f35), not (self.f35) or ((d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))])})
            rhs298_ = d_1802_i1_
            rhs299_ = d_1807_v10_
            lhs233_ = globalState
            lhs233_.f0 = rhs298_
            d_1807_v10_ = rhs299_
            (globalState).f5 = (d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]
        d_1808_v11_: _dafny.MultiSet
        d_1808_v11_ = _dafny.MultiSet([(d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))], ((self).f36).cf16, (d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))], (d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))], self.f35])
        (globalState).f0 = (((d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]) - ((d_1808_v11_).cardinality)) + (-651)
        d_1809_v12_: _dafny.Seq
        d_1809_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "miulhuko"))
        d_1810_i3_: int
        d_1810_i3_ = 0
        with _dafny.label("11"):
            while (((d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]) + (d_1793_v1_)) > (len(d_1809_v12_)):
                with _dafny.c_label("11"):
                    if (d_1810_i3_) >= (100):
                        raise _dafny.Break("11")
                    d_1810_i3_ = (d_1810_i3_) + (1)
                    d_1799_v5_ = d_1799_v5_
                    d_1811_v13_: D4
                    d_1811_v13_ = D4_DC10((0) - ((d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]))
                    d_1811_v13_ = d_1811_v13_
                    index269_ = default__.safeIndex(566, (d_1799_v5_).length(0))
                    (d_1799_v5_)[index269_] = (d_1809_v12_) <= ((d_1809_v12_) + (d_1809_v12_))
                    if self.f35:
                        d_1812_v14_: _dafny.Map
                        d_1812_v14_ = _dafny.Map({self.f35: d_1792_v0_})
                        d_1813_v15_: D19
                        d_1813_v15_ = D19_DC37(d_1812_v14_)
                        d_1814_v16_: D19
                        d_1814_v16_ = D19_DC39(d_1813_v15_)
                        d_1815_v17_: D19
                        d_1815_v17_ = D19_DC39(d_1814_v16_)
                        d_1816_v18_: D19
                        d_1816_v18_ = D19_DC39(d_1814_v16_)
                        d_1816_v18_ = d_1816_v18_
                        d_1817_v19_: _dafny.Array
                        nw308_ = _dafny.Array(_dafny.Array(None, 0), 7)
                        d_1817_v19_ = nw308_
                        index270_ = default__.safeIndex(611, (d_1817_v19_).length(0))
                        (d_1817_v19_)[index270_] = d_1792_v0_
                        index271_ = default__.safeIndex(611, (d_1817_v19_).length(0))
                        (d_1817_v19_)[index271_] = d_1792_v0_
                        index272_ = default__.safeIndex(566, (d_1799_v5_).length(0))
                        (d_1799_v5_)[index272_] = self.f35
                        d_1818_v20_: _dafny.Set
                        d_1818_v20_ = _dafny.Set({self.f35})
                        index273_ = default__.safeIndex(266, (d_1792_v0_).length(0))
                        rhs300_ = (0) - ((((d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]) + ((d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]) if ((d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))]) or ((d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))]) else (d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]))
                        rhs301_ = self.f35
                        rhs302_ = d_1818_v20_
                        lhs234_ = d_1792_v0_
                        lhs235_ = default__.safeIndex(266, (d_1792_v0_).length(0))
                        lhs234_[lhs235_] = rhs300_
                        r2 = rhs301_
                        d_1818_v20_ = rhs302_
                        d_1819_v21_: _dafny.Map
                        d_1819_v21_ = _dafny.Map({(d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))]: self.f35})
                        (self).f26 = (self.f26).set((0) - (d_1793_v1_), (len(_dafny.Map({(d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]: d_1793_v1_}))) - (len(d_1819_v21_)))
                    elif True:
                        (globalState).f1 = self.f35
                        def lambda89_(d_1820_v0_):
                            def lambda90_(d_1821_i4_):
                                return (d_1821_i4_) - ((d_1820_v0_)[default__.safeIndex(266, (d_1820_v0_).length(0))])

                            return lambda90_

                        init51_ = lambda89_(d_1792_v0_)
                        nw309_ = _dafny.Array(None, 19)
                        for i0_51_ in range(nw309_.length(0)):
                            nw309_[i0_51_] = init51_(i0_51_)
                        d_1792_v0_ = nw309_
                        index274_ = default__.safeIndex(566, (d_1799_v5_).length(0))
                        (d_1799_v5_)[index274_] = default__.fm0(len((_dafny.SeqWithoutIsStrInference([(d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]])) + (d_1801_v6_)), globalState)
                        (globalState).f2 = d_1793_v1_
                        d_1822_v22_: _dafny.Set
                        d_1822_v22_ = _dafny.Set({(d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))]})
                        d_1823_v23_: _dafny.Map
                        d_1823_v23_ = _dafny.Map({self.f35: 241})
                        d_1824_v24_: _dafny.Set
                        d_1824_v24_ = _dafny.Set({len(d_1822_v22_), 641, len(d_1823_v23_), (d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))], (d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]})
                        (globalState).f0 = len((d_1824_v24_ if self.f35 else (_dafny.Set({855, (d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))]})).intersection(d_1824_v24_)))
                    pass
            pass
        r2 = (d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))]
        r0 = default__.fm0((d_1792_v0_)[default__.safeIndex(266, (d_1792_v0_).length(0))], globalState)
        d_1825_v25_: _dafny.Set
        d_1825_v25_ = _dafny.Set({d_1799_v5_})
        d_1826_v26_: _dafny.Set
        d_1826_v26_ = _dafny.Set({_dafny.Map({d_1793_v1_: len(d_1825_v25_)})})
        r1 = ((d_1826_v26_ if self.f35 else d_1826_v26_)) | ((d_1826_v26_).intersection(_dafny.Set({self.f26})))
        r2 = (d_1795_v3_)[default__.safeIndex((len(_dafny.SeqWithoutIsStrInference([(d_1799_v5_)[default__.safeIndex(566, (d_1799_v5_).length(0))]]))) - ((0) - (len(d_1801_v6_))), len(d_1795_v3_))]
        return r0, r1, r2

    def m12(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_1827_v0_: _dafny.Seq
        d_1827_v0_ = _dafny.SeqWithoutIsStrInference([(self).fm3(globalState)])
        d_1828_v1_: int
        d_1828_v1_ = 159
        (self).f35 = default__.fm0(default__.safeDivisionInt((d_1827_v0_)[default__.safeIndex(d_1828_v1_, len(d_1827_v0_))], d_1828_v1_), globalState)
        d_1829_v2_: _dafny.Array
        nw310_ = _dafny.Array(int(0), 6)
        d_1829_v2_ = nw310_
        d_1830_v3_: D7
        d_1830_v3_ = D7_DC16(d_1828_v1_, self.f35)
        d_1831_v4_: _dafny.MultiSet
        d_1831_v4_ = _dafny.MultiSet([d_1830_v3_])
        d_1832_v5_: C3
        nw311_ = C3()
        nw311_.ctor__(d_1829_v2_, not((d_1831_v4_).ispropersubset(d_1831_v4_)), self.f26)
        d_1832_v5_ = nw311_
        (globalState).f1 = d_1832_v5_.f40
        d_1833_v6_: _dafny.Map
        d_1833_v6_ = _dafny.Map({True: d_1832_v5_.f40})
        d_1834_v7_: _dafny.MultiSet
        d_1834_v7_ = _dafny.MultiSet([d_1832_v5_.f40, (d_1832_v5_.f40 if not(((d_1833_v6_)[False] if (False) in (d_1833_v6_) else self.f35)) else self.f35)])
        d_1835_v8_: D12
        d_1835_v8_ = D12_DC28(d_1834_v7_)
        d_1834_v7_ = (d_1835_v8_).cf43
        d_1836_i0_: int
        d_1836_i0_ = 0
        with _dafny.label("12"):
            while (self.f35) == (False):
                with _dafny.c_label("12"):
                    if (d_1836_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_1836_i0_ = (d_1836_i0_) + (1)
                    d_1837_v9_: _dafny.Array
                    nw312_ = _dafny.Array(_dafny.CodePoint('D'), 19)
                    d_1837_v9_ = nw312_
                    d_1838_v10_: D0
                    d_1838_v10_ = D0_DC0(d_1837_v9_)
                    d_1839_v11_: int
                    d_1840_v12_: _dafny.Map
                    d_1841_v13_: int
                    out71_: int
                    out72_: _dafny.Map
                    out73_: int
                    out71_, out72_, out73_ = (d_1832_v5_).m1(d_1838_v10_, 850, (0) - ((d_1828_v1_) + (d_1828_v1_)), globalState)
                    d_1839_v11_ = out71_
                    d_1840_v12_ = out72_
                    d_1841_v13_ = out73_
                    d_1842_v14_: _dafny.Array
                    nw313_ = _dafny.Array(False, 27)
                    d_1842_v14_ = nw313_
                    index275_ = default__.safeIndex(942, (d_1842_v14_).length(0))
                    (d_1842_v14_)[index275_] = False
                    index276_ = default__.safeIndex(942, (d_1842_v14_).length(0))
                    (d_1842_v14_)[index276_] = d_1832_v5_.f40
                    d_1843_v15_: _dafny.Seq
                    d_1843_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dretpb"))
                    if default__.fm0(len(d_1843_v15_), globalState):
                        d_1844_v16_: _dafny.MultiSet
                        d_1844_v16_ = _dafny.MultiSet([d_1841_v13_])
                        index277_ = default__.safeIndex(154, ((d_1832_v5_).f39).length(0))
                        ((d_1832_v5_).f39)[index277_] = (d_1841_v13_ if (d_1842_v14_)[default__.safeIndex(942, (d_1842_v14_).length(0))] else ((d_1844_v16_)[d_1828_v1_] if (d_1828_v1_) in (d_1844_v16_) else d_1841_v13_))
                        index278_ = default__.safeIndex(154, ((d_1832_v5_).f39).length(0))
                        ((d_1832_v5_).f39)[index278_] = d_1839_v11_
                        d_1845_v17_: _dafny.MultiSet
                        d_1845_v17_ = _dafny.MultiSet([d_1843_v15_])
                        d_1846_v18_: _dafny.Seq
                        d_1846_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfss")), d_1843_v15_])
                        d_1847_v19_: _dafny.Seq
                        d_1847_v19_ = _dafny.SeqWithoutIsStrInference([d_1832_v5_.f40, d_1832_v5_.f40, default__.fm0(d_1839_v11_, globalState)])
                        d_1848_v20_: _dafny.Seq
                        d_1848_v20_ = _dafny.SeqWithoutIsStrInference([d_1847_v19_])
                        d_1849_v21_: _dafny.Map
                        d_1849_v21_ = _dafny.Map({(d_1845_v17_).intersection(_dafny.MultiSet((d_1846_v18_).set(default__.safeIndex(((d_1832_v5_).f39)[default__.safeIndex(154, ((d_1832_v5_).f39).length(0))], len(d_1846_v18_)), default__.fm30(len((d_1848_v20_)[default__.safeIndex(d_1839_v11_, len(d_1848_v20_))]), d_1843_v15_, ((d_1832_v5_).f39)[default__.safeIndex(154, ((d_1832_v5_).f39).length(0))], globalState)))): (d_1842_v14_)[default__.safeIndex(942, (d_1842_v14_).length(0))]})
                        d_1849_v21_ = (d_1849_v21_).set(d_1845_v17_, (d_1847_v19_)[default__.safeIndex(d_1828_v1_, len(d_1847_v19_))])
                        d_1847_v19_ = (D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpfks")), -716, d_1841_v13_, len(d_1846_v18_), _dafny.SeqWithoutIsStrInference([d_1832_v5_.f40]))).cf12
                        d_1829_v2_ = (d_1832_v5_).f39
                        d_1850_v22_: bool
                        d_1851_v23_: bool
                        d_1852_v24_: _dafny.Map
                        d_1853_v25_: int
                        out74_: bool
                        out75_: bool
                        out76_: _dafny.Map
                        out77_: int
                        out74_, out75_, out76_, out77_ = default__.m0(globalState)
                        d_1850_v22_ = out74_
                        d_1851_v23_ = out75_
                        d_1852_v24_ = out76_
                        d_1853_v25_ = out77_
                    elif True:
                        d_1827_v0_ = (d_1827_v0_) + (_dafny.SeqWithoutIsStrInference([d_1828_v1_, 100, d_1828_v1_]))
                        d_1854_v26_: bool
                        d_1855_v27_: bool
                        d_1856_v28_: _dafny.Map
                        d_1857_v29_: int
                        out78_: bool
                        out79_: bool
                        out80_: _dafny.Map
                        out81_: int
                        out78_, out79_, out80_, out81_ = default__.m0(globalState)
                        d_1854_v26_ = out78_
                        d_1855_v27_ = out79_
                        d_1856_v28_ = out80_
                        d_1857_v29_ = out81_
                        d_1858_v30_: C0
                        nw314_ = C0()
                        nw314_.ctor__(d_1837_v9_)
                        d_1858_v30_ = nw314_
                        d_1859_v31_: C2
                        nw315_ = C2()
                        nw315_.ctor__((d_1842_v14_)[default__.safeIndex(942, (d_1842_v14_).length(0))], len(_dafny.Map({d_1828_v1_: d_1832_v5_.f40})))
                        d_1859_v31_ = nw315_
                        d_1860_v32_: D3
                        d_1860_v32_ = D3_DC7(d_1843_v15_, d_1828_v1_, len(d_1833_v6_), d_1841_v13_, _dafny.SeqWithoutIsStrInference([d_1832_v5_.f40, d_1832_v5_.f40]))
                        d_1861_v33_: _dafny.Map
                        d_1861_v33_ = _dafny.Map({len((d_1860_v32_).cf12): True})
                        d_1855_v27_ = (self.f35 if ((d_1861_v33_)[(self).fm3(globalState)] if ((self).fm3(globalState)) in (d_1861_v33_) else False) else d_1832_v5_.f40)
                    d_1862_v34_: _dafny.Array
                    def lambda91_(d_1863_v13_, d_1864_v1_, d_1865_v11_):
                        def lambda92_(d_1866_i1_):
                            return (_dafny.MultiSet([d_1863_v13_, d_1864_v1_])) - (_dafny.MultiSet([d_1864_v1_, d_1865_v11_]))

                        return lambda92_

                    init52_ = lambda91_(d_1841_v13_, d_1828_v1_, d_1839_v11_)
                    nw316_ = _dafny.Array(None, 11)
                    for i0_52_ in range(nw316_.length(0)):
                        nw316_[i0_52_] = init52_(i0_52_)
                    d_1862_v34_ = nw316_
                    d_1867_v35_: _dafny.Map
                    d_1867_v35_ = _dafny.Map({self.f35: d_1843_v15_})
                    d_1868_v36_: _dafny.Seq
                    d_1868_v36_ = _dafny.SeqWithoutIsStrInference([self.f35])
                    d_1869_v37_: _dafny.Seq
                    d_1869_v37_ = _dafny.SeqWithoutIsStrInference([d_1867_v35_])
                    d_1870_v38_: _dafny.Seq
                    d_1870_v38_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1832_v5_.f40: d_1843_v15_}), (d_1869_v37_)[default__.safeIndex((d_1827_v0_)[default__.safeIndex(len(d_1843_v15_), len(d_1827_v0_))], len(d_1869_v37_))], (d_1867_v35_ if d_1832_v5_.f40 else d_1867_v35_)])
                    rhs303_ = d_1862_v34_
                    rhs304_ = (d_1842_v14_)[default__.safeIndex(942, (d_1842_v14_).length(0))]
                    rhs305_ = ((d_1868_v36_) + (d_1868_v36_)) == ((d_1868_v36_) + ((d_1868_v36_).set(default__.safeIndex(d_1841_v13_, len(d_1868_v36_)), self.f35)))
                    rhs306_ = (d_1870_v38_)[default__.safeIndex(((d_1834_v7_)[False] if (False) in (d_1834_v7_) else (0) - (-254)), len(d_1870_v38_))]
                    lhs236_ = self
                    d_1862_v34_ = rhs303_
                    lhs236_.f35 = rhs304_
                    r2 = rhs305_
                    d_1867_v35_ = rhs306_
                    pass
            pass
        d_1871_v39_: _dafny.Array
        nw317_ = _dafny.Array(False, 18)
        d_1871_v39_ = nw317_
        index279_ = default__.safeIndex(692, (d_1871_v39_).length(0))
        (d_1871_v39_)[index279_] = False
        d_1872_v40_: _dafny.Seq
        d_1872_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ieogrfj"))
        index280_ = default__.safeIndex(692, (d_1871_v39_).length(0))
        rhs307_ = not((d_1872_v40_) <= (d_1872_v40_))
        rhs308_ = not(self.f35)
        lhs237_ = d_1871_v39_
        lhs238_ = default__.safeIndex(692, (d_1871_v39_).length(0))
        lhs239_ = globalState
        lhs237_[lhs238_] = rhs307_
        lhs239_.f1 = rhs308_
        r0 = ((d_1828_v1_) - (110)) - ((0) - (d_1828_v1_))
        d_1873_v41_: str
        d_1873_v41_ = _dafny.CodePoint('j')
        d_1874_v42_: _dafny.Map
        d_1874_v42_ = _dafny.Map({d_1873_v41_: d_1828_v1_})
        r1 = ((d_1874_v42_)[d_1873_v41_] if (d_1873_v41_) in (d_1874_v42_) else d_1828_v1_)
        d_1875_v43_: _dafny.MultiSet
        d_1875_v43_ = _dafny.MultiSet([d_1828_v1_])
        d_1876_v44_: _dafny.Map
        d_1876_v44_ = _dafny.Map({d_1875_v43_: d_1832_v5_.f40})
        r2 = (d_1875_v43_) in (d_1876_v44_)
        r3 = _dafny.SeqWithoutIsStrInference([not(True)])
        return r0, r1, r2, r3

    @property
    def f36(self):
        return self._f36

class C9(T0):
    def  __init__(self):
        self._f26: _dafny.Map = _dafny.Map({})
        self.f33: _dafny.Array = _dafny.Array(None, 0)
        self._f34: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f26(self):
        return self._f26
    @f26.setter
    def f26(self, value):
        self._f26 = value
    def ctor__(self, f33, f34, f26):
        (self).f33 = f33
        (self)._f34 = f34
        (self).f26 = f26

    def fm3(self, globalState):
        return (self).f34

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife177_():
            coll73_ = _dafny.Map()
            compr_73_: int
            for compr_73_ in _dafny.IntegerRange(-644, 404):
                d_1877_v0_: int = compr_73_
                if ((-644) <= (d_1877_v0_)) and ((d_1877_v0_) < (404)):
                    coll73_[(d_1877_v0_) - ((self).f34)] = _dafny.Map({_dafny.CodePoint('m'): (self).f34})
            return _dafny.Map(coll73_)
        def iife178_():
            coll74_ = _dafny.Map()
            compr_74_: str
            for compr_74_ in (_dafny.Map({_dafny.CodePoint('s'): not(True)})).keys.Elements:
                d_1878_v1_: str = compr_74_
                if (d_1878_v1_) in (_dafny.Map({_dafny.CodePoint('s'): not(True)})):
                    coll74_[d_1878_v1_] = (self).f34
            return _dafny.Map(coll74_)
        return (iife177_()
        ) | ((_dafny.Map({524: iife178_()
        })) | (_dafny.Map({(_dafny.MultiSet([(0) - (len(_dafny.Map({-38: True})))])).cardinality: _dafny.Map({_dafny.CodePoint('e'): (self).f34})})))

    def fm11(self, p0, p1, p2, globalState):
        return ((self).f34) in ((_dafny.Map({(self).f34: not(False)})) | (_dafny.Map({(self).f34: True})))

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        d_1879_v0_: bool
        d_1879_v0_ = True
        d_1880_v1_: _dafny.Map
        d_1880_v1_ = _dafny.Map({d_1879_v0_: p1})
        d_1881_v2_: C1
        nw318_ = C1()
        nw318_.ctor__((d_1880_v1_) | (d_1880_v1_))
        d_1881_v2_ = nw318_
        if d_1879_v0_:
            d_1882_v3_: _dafny.Array
            def lambda93_(d_1883_i0_):
                return _dafny.CodePoint('n')

            init53_ = lambda93_
            nw319_ = _dafny.Array(None, 10)
            for i0_53_ in range(nw319_.length(0)):
                nw319_[i0_53_] = init53_(i0_53_)
            d_1882_v3_ = nw319_
            pat_let_tv54_ = d_1882_v3_
            pat_let_tv55_ = d_1882_v3_
            d_1884_v4_: _dafny.Array
            nw320_ = _dafny.Array(None, 23)
            nw320_[int(0)] = p0
            nw320_[int(1)] = p0
            nw320_[int(2)] = p0
            def iife179_(_pat_let52_0):
                def iife180_(d_1885_dt__update__tmp_h0_):
                    def iife181_(_pat_let53_0):
                        def iife182_(d_1886_dt__update_hcf0_h0_):
                            return D0_DC0(d_1886_dt__update_hcf0_h0_)
                        return iife182_(_pat_let53_0)
                    return iife181_(pat_let_tv54_)
                return iife180_(_pat_let52_0)
            nw320_[int(3)] = iife179_(p0)
            nw320_[int(4)] = p0
            nw320_[int(5)] = p0
            nw320_[int(6)] = p0
            nw320_[int(7)] = p0
            nw320_[int(8)] = p0
            def iife183_(_pat_let54_0):
                def iife184_(d_1887_dt__update__tmp_h1_):
                    def iife185_(_pat_let55_0):
                        def iife186_(d_1888_dt__update_hcf0_h1_):
                            return D0_DC0(d_1888_dt__update_hcf0_h1_)
                        return iife186_(_pat_let55_0)
                    return iife185_(pat_let_tv55_)
                return iife184_(_pat_let54_0)
            nw320_[int(9)] = iife183_(p0)
            nw320_[int(10)] = p0
            nw320_[int(11)] = p0
            nw320_[int(12)] = p0
            nw320_[int(13)] = p0
            nw320_[int(14)] = p0
            nw320_[int(15)] = p0
            nw320_[int(16)] = p0
            nw320_[int(17)] = p0
            nw320_[int(18)] = p0
            nw320_[int(19)] = p0
            nw320_[int(20)] = p0
            nw320_[int(21)] = p0
            nw320_[int(22)] = D0_DC0(d_1882_v3_)
            d_1884_v4_ = nw320_
            index281_ = default__.safeIndex(621, (d_1884_v4_).length(0))
            (d_1884_v4_)[index281_] = p0
            index282_ = default__.safeIndex(621, (d_1884_v4_).length(0))
            (d_1884_v4_)[index282_] = p0
            (globalState).f16 = (self).f34
            d_1889_v5_: _dafny.Array
            nw321_ = _dafny.Array(int(0), 9)
            d_1889_v5_ = nw321_
            d_1890_v6_: _dafny.Map
            d_1890_v6_ = _dafny.Map({d_1889_v5_: self.f33})
            d_1891_v7_: D7
            d_1891_v7_ = D7_DC17(d_1879_v0_, d_1879_v0_, 752)
            d_1892_v8_: _dafny.Seq
            d_1892_v8_ = _dafny.SeqWithoutIsStrInference([(self).f34, (d_1891_v7_).cf26, -917, p1])
            d_1893_v9_: _dafny.Seq
            d_1893_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwiqa"))
            d_1894_v10_: str
            d_1894_v10_ = _dafny.CodePoint('d')
            d_1895_v11_: _dafny.Seq
            d_1895_v11_ = _dafny.SeqWithoutIsStrInference([d_1879_v0_, d_1879_v0_, d_1879_v0_, d_1879_v0_])
            d_1896_v12_: D3
            d_1896_v12_ = D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "en")), p1, p2, (self).f34, d_1895_v11_)
            d_1897_v13_: _dafny.Map
            d_1897_v13_ = _dafny.Map({(self).f34: d_1879_v0_})
            d_1898_v14_: _dafny.Set
            d_1898_v14_ = _dafny.Set({False, d_1879_v0_, default__.fm0(len(d_1897_v13_), globalState), d_1879_v0_, not(d_1879_v0_)})
            d_1899_v15_: _dafny.Map
            d_1899_v15_ = _dafny.Map({d_1879_v0_: d_1898_v14_})
            d_1900_v16_: D3
            d_1900_v16_ = D3_DC7((d_1896_v12_).cf8, p2, (self).f34, len(d_1899_v15_), d_1895_v11_)
            rhs309_ = d_1890_v6_
            rhs310_ = default__.fm21(d_1894_v10_, (_dafny.MultiSet([d_1889_v5_])).cardinality, d_1896_v12_, globalState)
            rhs311_ = default__.fm0((self).f34, globalState)
            rhs312_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdiqiqha"))) + (d_1893_v9_)
            rhs313_ = d_1893_v9_
            lhs240_ = globalState
            d_1890_v6_ = rhs309_
            d_1892_v8_ = rhs310_
            lhs240_.f13 = rhs311_
            d_1893_v9_ = rhs312_
            d_1893_v9_ = rhs313_
            (globalState).f13 = d_1879_v0_
            d_1879_v0_ = (False if not (d_1879_v0_) or (d_1879_v0_) else (p2) != (default__.fm2(globalState)))
        elif True:
            (globalState).f1 = ((d_1879_v0_) and (d_1879_v0_)) == ((d_1879_v0_ if True else d_1879_v0_))
            d_1901_v17_: _dafny.Array
            def lambda94_(d_1902_p2_):
                def lambda95_(d_1903_i1_):
                    return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpvnqrd"))))]), _dafny.MultiSet([194, d_1902_p2_])])

                return lambda95_

            init54_ = lambda94_(p2)
            nw322_ = _dafny.Array(None, 1)
            for i0_54_ in range(nw322_.length(0)):
                nw322_[i0_54_] = init54_(i0_54_)
            d_1901_v17_ = nw322_
            d_1904_v18_: _dafny.MultiSet
            d_1904_v18_ = _dafny.MultiSet([p2])
            d_1905_v19_: _dafny.Seq
            d_1905_v19_ = _dafny.SeqWithoutIsStrInference([d_1904_v18_, d_1904_v18_])
            index283_ = default__.safeIndex(296, (d_1901_v17_).length(0))
            (d_1901_v17_)[index283_] = d_1905_v19_
            index284_ = default__.safeIndex(296, (d_1901_v17_).length(0))
            (d_1901_v17_)[index284_] = ((_dafny.SeqWithoutIsStrInference([d_1904_v18_])) + (d_1905_v19_)) + (d_1905_v19_)
            d_1906_v20_: _dafny.Array
            nw323_ = _dafny.Array(D9.default()(), 20)
            d_1906_v20_ = nw323_
            d_1907_v21_: _dafny.Array
            d_1907_v21_ = d_1906_v20_
            d_1908_v22_: _dafny.Array
            nw324_ = _dafny.Array(None, 23)
            nw324_[int(0)] = d_1907_v21_
            nw324_[int(1)] = d_1907_v21_
            nw324_[int(2)] = d_1907_v21_
            nw324_[int(3)] = d_1907_v21_
            nw324_[int(4)] = d_1907_v21_
            nw324_[int(5)] = d_1907_v21_
            nw324_[int(6)] = d_1907_v21_
            nw324_[int(7)] = d_1907_v21_
            nw324_[int(8)] = d_1907_v21_
            nw324_[int(9)] = d_1907_v21_
            nw324_[int(10)] = d_1907_v21_
            nw324_[int(11)] = d_1907_v21_
            nw324_[int(12)] = (d_1907_v21_ if True else d_1906_v20_)
            nw324_[int(13)] = d_1906_v20_
            nw324_[int(14)] = d_1907_v21_
            nw324_[int(15)] = d_1907_v21_
            nw324_[int(16)] = d_1907_v21_
            nw324_[int(17)] = d_1907_v21_
            nw324_[int(18)] = d_1907_v21_
            nw324_[int(19)] = d_1907_v21_
            nw324_[int(20)] = d_1906_v20_
            nw324_[int(21)] = d_1907_v21_
            nw324_[int(22)] = d_1907_v21_
            d_1908_v22_ = nw324_
            index285_ = default__.safeIndex(775, (d_1908_v22_).length(0))
            (d_1908_v22_)[index285_] = d_1907_v21_
            index286_ = default__.safeIndex(775, (d_1908_v22_).length(0))
            (d_1908_v22_)[index286_] = d_1907_v21_
            source27_ = default__.fm50(globalState)
            d_1909___mcc_h0_ = source27_
            d_1910_cf45_ = d_1909___mcc_h0_
            d_1911_v23_: _dafny.Seq
            d_1911_v23_ = _dafny.SeqWithoutIsStrInference([p2, (self).f34])
            d_1912_v24_: _dafny.MultiSet
            d_1912_v24_ = _dafny.MultiSet([d_1911_v23_, d_1911_v23_])
            d_1913_v25_: _dafny.Set
            d_1913_v25_ = _dafny.Set({len(d_1911_v23_)})
            d_1914_v26_: D8
            d_1914_v26_ = D8_DC21(D8_DC18(d_1913_v25_))
            d_1915_v27_: D8
            d_1915_v27_ = D8_DC20(d_1879_v0_)
            pat_let_tv56_ = d_1915_v27_
            d_1916_v28_: _dafny.MultiSet
            def iife187_(_pat_let56_0):
                def iife188_(d_1917_dt__update__tmp_h2_):
                    def iife189_(_pat_let57_0):
                        def iife190_(d_1918_dt__update_hcf32_h0_):
                            return D8_DC21(d_1918_dt__update_hcf32_h0_)
                        return iife190_(_pat_let57_0)
                    return iife189_(pat_let_tv56_)
                return iife188_(_pat_let56_0)
            d_1916_v28_ = default__.fm51((self).f34, d_1879_v0_, d_1912_v24_, iife187_(d_1914_v26_), globalState)
            d_1916_v28_ = d_1916_v28_
            (globalState).f5 = len(d_1911_v23_)
            d_1919_v29_: _dafny.Seq
            d_1919_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhv"))
            (globalState).f21 = len(d_1919_v29_)
            (globalState).f16 = p2
            arr2_ = self.f33
            index287_ = default__.safeIndex(820, (self.f33).length(0))
            arr2_[index287_] = d_1879_v0_
            d_1920_v30_: _dafny.Seq
            d_1920_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpwlu"))
            d_1921_v31_: str
            d_1921_v31_ = _dafny.CodePoint('k')
            d_1922_v32_: _dafny.Set
            d_1922_v32_ = _dafny.Set({d_1921_v31_})
            d_1923_v33_: D20
            d_1923_v33_ = D20_DC40(d_1922_v32_)
            arr3_ = self.f33
            index288_ = default__.safeIndex(820, (self.f33).length(0))
            rhs314_ = not (d_1879_v0_) or ((d_1920_v30_) <= ((d_1920_v30_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f34])), len(d_1920_v30_)), _dafny.CodePoint('s'))))
            rhs315_ = ((d_1923_v33_).cf56).issubset(d_1922_v32_)
            lhs241_ = globalState
            lhs242_ = self.f33
            lhs243_ = default__.safeIndex(820, (self.f33).length(0))
            lhs241_.f1 = rhs314_
            lhs242_[lhs243_] = rhs315_
        hi9_ = p2
        for d_1924_i2_ in range(p1, hi9_):
            d_1925_v34_: _dafny.Seq
            d_1925_v34_ = _dafny.SeqWithoutIsStrInference([d_1879_v0_, d_1879_v0_, d_1879_v0_, d_1879_v0_, True])
            d_1926_v35_: _dafny.MultiSet
            d_1926_v35_ = _dafny.MultiSet([d_1925_v34_, d_1925_v34_])
            d_1926_v35_ = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_1925_v34_, d_1925_v34_, d_1925_v34_, d_1925_v34_])) + (_dafny.SeqWithoutIsStrInference([d_1925_v34_])))
            d_1927_v36_: _dafny.Seq
            d_1927_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdk"))
            d_1928_v37_: _dafny.Map
            d_1928_v37_ = _dafny.Map({len(d_1927_v36_): d_1927_v36_})
            d_1929_v38_: _dafny.Set
            d_1929_v38_ = _dafny.Set({p2})
            d_1930_v39_: _dafny.Map
            d_1930_v39_ = _dafny.Map({d_1928_v37_: len(d_1929_v38_)})
            d_1930_v39_ = (d_1930_v39_).set(d_1928_v37_, default__.safeModuloInt(p2, d_1924_i2_))
            (globalState).f13 = d_1879_v0_
            d_1931_v40_: _dafny.MultiSet
            d_1931_v40_ = _dafny.MultiSet([d_1879_v0_])
            d_1932_v41_: _dafny.Set
            d_1932_v41_ = _dafny.Set({d_1879_v0_})
            d_1933_v42_: _dafny.Seq
            d_1933_v42_ = _dafny.SeqWithoutIsStrInference([len(d_1932_v41_)])
            d_1934_v43_: _dafny.Map
            d_1934_v43_ = _dafny.Map({(d_1931_v40_).cardinality: d_1933_v42_})
            d_1935_v44_: _dafny.Map
            d_1935_v44_ = _dafny.Map({(0) - ((self).f34): d_1925_v34_})
            d_1936_v45_: _dafny.Map
            d_1936_v45_ = _dafny.Map({d_1879_v0_: d_1879_v0_})
            d_1937_v46_: _dafny.MultiSet
            d_1937_v46_ = _dafny.MultiSet([len(d_1936_v45_)])
            d_1938_v47_: _dafny.Seq
            d_1938_v47_ = _dafny.SeqWithoutIsStrInference([d_1937_v46_, d_1937_v46_])
            d_1939_v48_: D5
            d_1939_v48_ = D5_DC12(not(False), d_1879_v0_)
            d_1940_v49_: _dafny.Map
            d_1940_v49_ = _dafny.Map({d_1939_v48_: not(True)})
            d_1941_v50_: D7
            d_1941_v50_ = D7_DC17(d_1879_v0_, not(d_1879_v0_), p1)
            d_1942_v51_: _dafny.Array
            nw325_ = _dafny.Array(None, 25)
            nw325_[int(0)] = p2
            nw325_[int(1)] = p2
            nw325_[int(2)] = len(((d_1934_v43_)[p2] if (p2) in (d_1934_v43_) else d_1933_v42_))
            nw325_[int(3)] = p2
            nw325_[int(4)] = 374
            nw325_[int(5)] = (self).f34
            nw325_[int(6)] = (d_1933_v42_)[default__.safeIndex(p1, len(d_1933_v42_))]
            nw325_[int(7)] = (0) - ((len(d_1933_v42_)) - ((self).f34))
            nw325_[int(8)] = len(d_1935_v44_)
            nw325_[int(9)] = ((self).f34) * (((self.f26)[p1] if (p1) in (self.f26) else len(d_1938_v47_)))
            nw325_[int(10)] = p1
            nw325_[int(11)] = (0) - (p2)
            nw325_[int(12)] = len(default__.fm29(_dafny.MultiSet(d_1925_v34_), p1, d_1879_v0_, 368, globalState))
            nw325_[int(13)] = p2
            nw325_[int(14)] = (p2 if d_1879_v0_ else p1)
            nw325_[int(15)] = (len(d_1940_v49_)) - (p2)
            nw325_[int(16)] = default__.fm2(globalState)
            nw325_[int(17)] = (self).f34
            nw325_[int(18)] = (self).f34
            nw325_[int(19)] = p1
            nw325_[int(20)] = d_1924_i2_
            nw325_[int(21)] = (d_1941_v50_).cf26
            nw325_[int(22)] = default__.safeModuloInt(d_1924_i2_, 36)
            nw325_[int(23)] = p1
            nw325_[int(24)] = default__.safeModuloInt((self).f34, len(d_1927_v36_))
            d_1942_v51_ = nw325_
            index289_ = default__.safeIndex(58, (d_1942_v51_).length(0))
            (d_1942_v51_)[index289_] = p2
            index290_ = default__.safeIndex(58, (d_1942_v51_).length(0))
            (d_1942_v51_)[index290_] = d_1924_i2_
        d_1943_v52_: _dafny.Seq
        d_1943_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
        d_1944_v53_: _dafny.Set
        d_1944_v53_ = _dafny.Set({d_1943_v52_, d_1943_v52_, (d_1943_v52_) + (d_1943_v52_), d_1943_v52_})
        d_1945_v54_: _dafny.MultiSet
        d_1945_v54_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1946_i3_ in range(default__.abs(507))]), d_1943_v52_, d_1943_v52_])
        def iife191_():
            coll75_ = _dafny.Set()
            compr_75_: _dafny.Seq
            for compr_75_ in (d_1945_v54_).Elements:
                d_1947_v55_: _dafny.Seq = compr_75_
                if (d_1947_v55_) in (d_1945_v54_):
                    coll75_ = coll75_.union(_dafny.Set([d_1947_v55_]))
            return _dafny.Set(coll75_)
        d_1944_v53_ = (iife191_()
        ) | (d_1944_v53_)
        d_1948_i4_: int
        d_1948_i4_ = 0
        with _dafny.label("13"):
            while d_1879_v0_:
                with _dafny.c_label("13"):
                    if (d_1948_i4_) >= (100):
                        raise _dafny.Break("13")
                    d_1948_i4_ = (d_1948_i4_) + (1)
                    d_1949_v56_: _dafny.Array
                    def lambda96_(d_1950_i5_):
                        return (d_1950_i5_) + (len(_dafny.SeqWithoutIsStrInference([(self).f34])))

                    init55_ = lambda96_
                    nw326_ = _dafny.Array(None, 25)
                    for i0_55_ in range(nw326_.length(0)):
                        nw326_[i0_55_] = init55_(i0_55_)
                    d_1949_v56_ = nw326_
                    d_1951_v57_: C3
                    nw327_ = C3()
                    nw327_.ctor__(d_1949_v56_, d_1879_v0_, self.f26)
                    d_1951_v57_ = nw327_
                    if (default__.safeModuloInt(p1, (self).f34)) < ((self).f34):
                        d_1952_v58_: _dafny.MultiSet
                        d_1952_v58_ = _dafny.MultiSet([p2])
                        d_1953_v59_: _dafny.Map
                        d_1953_v59_ = _dafny.Map({p1: (d_1951_v57_).f39})
                        d_1954_v60_: str
                        d_1954_v60_ = _dafny.CodePoint('u')
                        d_1955_v61_: D0
                        d_1955_v61_ = D0_DC1(((d_1953_v59_)[p1] if (p1) in (d_1953_v59_) else ((d_1953_v59_)[(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality] if ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality) in (d_1953_v59_) else (d_1951_v57_).f39)), d_1954_v60_, p2)
                        d_1956_v62_: _dafny.Seq
                        d_1956_v62_ = _dafny.SeqWithoutIsStrInference([(((d_1952_v58_)[(self).f34] if ((self).f34) in (d_1952_v58_) else (self).f34)) * ((d_1955_v61_).cf3)])
                        (globalState).f16 = len(d_1956_v62_)
                        d_1957_v63_: _dafny.Map
                        d_1957_v63_ = _dafny.Map({p1: d_1951_v57_.f40})
                        (globalState).f0 = default__.safeDivisionInt(len((_dafny.Map({p2: d_1951_v57_.f40})) | (d_1957_v63_)), 768)
                        d_1958_v64_: _dafny.Array
                        nw328_ = _dafny.Array(_dafny.Set({}), 1)
                        d_1958_v64_ = nw328_
                        d_1959_v65_: _dafny.Set
                        d_1959_v65_ = _dafny.Set({(d_1955_v61_).cf1, d_1949_v56_})
                        index291_ = default__.safeIndex(18, (d_1958_v64_).length(0))
                        (d_1958_v64_)[index291_] = (d_1959_v65_).intersection(d_1959_v65_)
                        index292_ = default__.safeIndex(18, (d_1958_v64_).length(0))
                        (d_1958_v64_)[index292_] = d_1959_v65_
                        (globalState).f21 = ((d_1881_v2_.f43)[d_1951_v57_.f40] if (d_1951_v57_.f40) in (d_1881_v2_.f43) else (0) - (p1))
                        arr4_ = self.f33
                        index293_ = default__.safeIndex(487, (self.f33).length(0))
                        arr4_[index293_] = True
                        arr5_ = self.f33
                        index294_ = default__.safeIndex(487, (self.f33).length(0))
                        arr5_[index294_] = default__.fm0(p2, globalState)
                    elif True:
                        d_1960_v66_: str
                        d_1960_v66_ = _dafny.CodePoint('r')
                        d_1961_v67_: _dafny.MultiSet
                        d_1961_v67_ = _dafny.MultiSet([d_1960_v66_, d_1960_v66_, d_1960_v66_])
                        (globalState).f13 = (_dafny.MultiSet([d_1960_v66_, (d_1943_v52_)[default__.safeIndex((self).f34, len(d_1943_v52_))]])).ispropersubset(d_1961_v67_)
                        (globalState).f1 = (p1) < (p2)
                        d_1962_v68_: D7
                        d_1962_v68_ = D7_DC16((d_1951_v57_).fm3(globalState), d_1879_v0_)
                        d_1962_v68_ = d_1962_v68_
                        (globalState).f2 = (default__.safeModuloInt(p1, (self).f34)) + (p1)
                        d_1963_v69_: _dafny.Array
                        nw329_ = _dafny.Array(_dafny.Array(None, 0), 5)
                        d_1963_v69_ = nw329_
                        d_1964_v71_: _dafny.Array
                        nw330_ = _dafny.Array(None, 27)
                        nw330_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnpqseima"))
                        nw330_[int(1)] = d_1943_v52_
                        nw330_[int(2)] = d_1943_v52_
                        nw330_[int(3)] = d_1943_v52_
                        nw330_[int(4)] = d_1943_v52_
                        nw330_[int(5)] = d_1943_v52_
                        nw330_[int(6)] = d_1943_v52_
                        nw330_[int(7)] = d_1943_v52_
                        nw330_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th"))
                        nw330_[int(9)] = d_1943_v52_
                        nw330_[int(10)] = d_1943_v52_
                        nw330_[int(11)] = d_1943_v52_
                        nw330_[int(12)] = d_1943_v52_
                        nw330_[int(13)] = _dafny.SeqWithoutIsStrInference([d_1960_v66_ for d_1965_i6_ in range(default__.abs(34))])
                        nw330_[int(14)] = d_1943_v52_
                        def iife192_():
                            coll76_ = _dafny.Set()
                            compr_76_: int
                            for compr_76_ in _dafny.IntegerRange(-785, 88):
                                d_1966_v70_: int = compr_76_
                                if ((-785) <= (d_1966_v70_)) and ((d_1966_v70_) < (88)):
                                    coll76_ = coll76_.union(_dafny.Set([(d_1966_v70_) * ((self).f34)]))
                            return _dafny.Set(coll76_)
                        nw330_[int(15)] = default__.fm30(p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), len(iife192_()
                        ), globalState)
                        nw330_[int(16)] = d_1943_v52_
                        nw330_[int(17)] = d_1943_v52_
                        nw330_[int(18)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1967_i7_ in range(default__.abs(246))])
                        nw330_[int(19)] = d_1943_v52_
                        nw330_[int(20)] = d_1943_v52_
                        nw330_[int(21)] = _dafny.SeqWithoutIsStrInference([d_1960_v66_ for d_1968_i8_ in range(default__.abs(865))])
                        nw330_[int(22)] = d_1943_v52_
                        nw330_[int(23)] = _dafny.SeqWithoutIsStrInference([d_1960_v66_ for d_1969_i9_ in range(default__.abs(-355))])
                        nw330_[int(24)] = d_1943_v52_
                        nw330_[int(25)] = _dafny.SeqWithoutIsStrInference([d_1960_v66_ for d_1970_i10_ in range(default__.abs(624))])
                        nw330_[int(26)] = d_1943_v52_
                        d_1964_v71_ = nw330_
                        index295_ = default__.safeIndex(139, (d_1963_v69_).length(0))
                        (d_1963_v69_)[index295_] = d_1964_v71_
                        index296_ = default__.safeIndex(139, (d_1963_v69_).length(0))
                        (d_1963_v69_)[index296_] = d_1964_v71_
                    d_1971_v72_: _dafny.Map
                    d_1971_v72_ = _dafny.Map({p1: d_1879_v0_})
                    d_1972_v73_: _dafny.Map
                    d_1972_v73_ = _dafny.Map({d_1971_v72_: d_1879_v0_})
                    d_1973_v74_: _dafny.Map
                    d_1973_v74_ = _dafny.Map({d_1879_v0_: (d_1972_v73_) | (d_1972_v73_)})
                    d_1974_v75_: _dafny.Seq
                    d_1974_v75_ = _dafny.SeqWithoutIsStrInference([((d_1971_v72_)[p1] if (p1) in (d_1971_v72_) else d_1951_v57_.f40), True, d_1879_v0_, d_1879_v0_, d_1879_v0_])
                    d_1973_v74_ = (d_1973_v74_).set((d_1974_v75_) != (default__.fm13((self).f34, globalState)), (d_1972_v73_) | (d_1972_v73_))
                    d_1975_v76_: _dafny.Map
                    d_1975_v76_ = _dafny.Map({(self).f34: (d_1943_v52_) + (d_1943_v52_)})
                    d_1975_v76_ = d_1975_v76_
                    pass
            pass
        hi10_ = p1
        for d_1976_i11_ in range((0) - ((0) - (len((d_1943_v52_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afqvtspln")))))), hi10_):
            d_1879_v0_ = d_1879_v0_
            d_1977_v77_: str
            d_1977_v77_ = _dafny.CodePoint('l')
            d_1978_v78_: _dafny.Map
            d_1978_v78_ = _dafny.Map({self.f33: d_1943_v52_})
            d_1879_v0_ = (d_1977_v77_) not in (((d_1978_v78_)[self.f33] if (self.f33) in (d_1978_v78_) else d_1943_v52_))
            d_1979_v79_: _dafny.Set
            d_1979_v79_ = _dafny.Set({d_1879_v0_, d_1879_v0_})
            arr6_ = self.f33
            index297_ = default__.safeIndex(853, (self.f33).length(0))
            arr6_[index297_] = (d_1879_v0_) not in (d_1979_v79_)
            d_1980_v80_: _dafny.Array
            nw331_ = _dafny.Array(int(0), 21)
            d_1980_v80_ = nw331_
            d_1981_v81_: _dafny.Seq
            d_1981_v81_ = _dafny.SeqWithoutIsStrInference([p1])
            arr7_ = self.f33
            index298_ = default__.safeIndex(853, (self.f33).length(0))
            rhs316_ = d_1977_v77_
            rhs317_ = (d_1943_v52_).set(default__.safeIndex(134, len(d_1943_v52_)), d_1977_v77_)
            rhs318_ = (d_1976_i11_) * ((0) - ((D0_DC1(d_1980_v80_, d_1977_v77_, (self).f34)).cf3))
            rhs319_ = d_1879_v0_
            rhs320_ = ((d_1880_v1_ if d_1879_v0_ else d_1880_v1_)) == (default__.fm33(d_1981_v81_, d_1879_v0_, globalState))
            lhs244_ = self.f33
            lhs245_ = default__.safeIndex(853, (self.f33).length(0))
            d_1977_v77_ = rhs316_
            d_1943_v52_ = rhs317_
            r0 = rhs318_
            d_1879_v0_ = rhs319_
            lhs244_[lhs245_] = rhs320_
            d_1982_v82_: _dafny.Array
            nw332_ = _dafny.Array(None, 4)
            nw332_[int(0)] = d_1980_v80_
            nw332_[int(1)] = d_1980_v80_
            nw332_[int(2)] = d_1980_v80_
            nw332_[int(3)] = d_1980_v80_
            d_1982_v82_ = nw332_
            index299_ = default__.safeIndex(398, (d_1982_v82_).length(0))
            (d_1982_v82_)[index299_] = d_1980_v80_
            index300_ = default__.safeIndex(398, (d_1982_v82_).length(0))
            (d_1982_v82_)[index300_] = d_1980_v80_
        r0 = (self).f34
        d_1983_v83_: _dafny.Map
        d_1983_v83_ = _dafny.Map({d_1943_v52_: self.f33})
        r1 = ((d_1983_v83_) | (_dafny.Map({d_1943_v52_: self.f33}))) | (d_1983_v83_)
        r2 = -639
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        (globalState).f0 = 664
        d_1984_v0_: bool
        d_1984_v0_ = False
        d_1985_v1_: D11
        d_1985_v1_ = D11_DC27((self).f34, default__.fm2(globalState), d_1984_v0_)
        source28_ = d_1985_v1_
        if source28_.is_DC27:
            d_1986___mcc_h0_ = source28_.cf40
            d_1987___mcc_h1_ = source28_.cf41
            d_1988___mcc_h2_ = source28_.cf42
            d_1989_cf42_ = d_1988___mcc_h2_
            d_1990_cf41_ = d_1987___mcc_h1_
            d_1991_cf40_ = d_1986___mcc_h0_
            (globalState).f0 = d_1990_cf41_
            d_1992_v2_: _dafny.Array
            nw333_ = _dafny.Array(_dafny.Seq({}), 8)
            d_1992_v2_ = nw333_
            d_1993_v3_: _dafny.Seq
            d_1993_v3_ = _dafny.SeqWithoutIsStrInference([self.f26])
            index301_ = default__.safeIndex(597, (d_1992_v2_).length(0))
            (d_1992_v2_)[index301_] = d_1993_v3_
            index302_ = default__.safeIndex(597, (d_1992_v2_).length(0))
            (d_1992_v2_)[index302_] = d_1993_v3_
            if d_1989_cf42_:
                arr8_ = self.f33
                index303_ = default__.safeIndex(755, (self.f33).length(0))
                arr8_[index303_] = (d_1989_cf42_) and (False)
                d_1994_v4_: _dafny.Seq
                d_1994_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ko"))
                d_1995_v5_: _dafny.Map
                d_1995_v5_ = _dafny.Map({d_1984_v0_: len(d_1994_v4_)})
                d_1996_v6_: D5
                d_1996_v6_ = D5_DC11(d_1984_v0_)
                d_1997_v7_: C8
                nw334_ = C8()
                nw334_.ctor__(d_1984_v0_, d_1996_v6_, _dafny.Map({d_1990_cf41_: (self).f34}))
                d_1997_v7_ = nw334_
                d_1998_v8_: _dafny.MultiSet
                d_1998_v8_ = _dafny.MultiSet([d_1997_v7_, d_1997_v7_, d_1997_v7_])
                d_1999_v9_: _dafny.Set
                d_1999_v9_ = _dafny.Set({d_1998_v8_, d_1998_v8_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1997_v7_])), d_1998_v8_, d_1998_v8_})
                arr9_ = self.f33
                index304_ = default__.safeIndex(755, (self.f33).length(0))
                rhs321_ = ((d_1985_v1_).cf42) == ((d_1995_v5_) == (_dafny.Map({d_1989_cf42_: d_1990_cf41_})))
                rhs322_ = d_1990_cf41_
                rhs323_ = (0) - (d_1991_cf40_)
                rhs324_ = (d_1999_v9_).isdisjoint(d_1999_v9_)
                lhs246_ = self.f33
                lhs247_ = default__.safeIndex(755, (self.f33).length(0))
                lhs248_ = globalState
                lhs249_ = globalState
                lhs250_ = globalState
                lhs246_[lhs247_] = rhs321_
                lhs248_.f0 = rhs322_
                lhs249_.f0 = rhs323_
                lhs250_.f1 = rhs324_
                d_2000_v10_: C7
                nw335_ = C7()
                nw335_.ctor__(d_1997_v7_.f35)
                d_2000_v10_ = nw335_
                d_2001_v11_: _dafny.Seq
                d_2001_v11_ = _dafny.SeqWithoutIsStrInference([d_2000_v10_])
                d_2002_v12_: _dafny.Array
                nw336_ = _dafny.Array(None, 21)
                nw336_[int(0)] = d_2000_v10_
                nw336_[int(1)] = d_2000_v10_
                nw336_[int(2)] = d_2000_v10_
                nw336_[int(3)] = d_2000_v10_
                nw336_[int(4)] = d_2000_v10_
                nw336_[int(5)] = d_2000_v10_
                nw336_[int(6)] = d_2000_v10_
                nw336_[int(7)] = d_2000_v10_
                nw336_[int(8)] = d_2000_v10_
                nw336_[int(9)] = d_2000_v10_
                nw336_[int(10)] = d_2000_v10_
                nw336_[int(11)] = d_2000_v10_
                nw336_[int(12)] = d_2000_v10_
                nw336_[int(13)] = d_2000_v10_
                nw336_[int(14)] = (d_2001_v11_)[default__.safeIndex(d_1991_cf40_, len(d_2001_v11_))]
                nw336_[int(15)] = d_2000_v10_
                nw336_[int(16)] = d_2000_v10_
                nw336_[int(17)] = d_2000_v10_
                nw336_[int(18)] = d_2000_v10_
                nw336_[int(19)] = d_2000_v10_
                nw336_[int(20)] = d_2000_v10_
                d_2002_v12_ = nw336_
                index305_ = default__.safeIndex(470, (d_2002_v12_).length(0))
                (d_2002_v12_)[index305_] = d_2000_v10_
                d_2003_v13_: _dafny.Array
                def lambda97_(d_2004_i0_):
                    return (d_2004_i0_) + ((self).f34)

                init56_ = lambda97_
                nw337_ = _dafny.Array(None, 16)
                for i0_56_ in range(nw337_.length(0)):
                    nw337_[i0_56_] = init56_(i0_56_)
                d_2003_v13_ = nw337_
                index306_ = default__.safeIndex(470, (d_2002_v12_).length(0))
                rhs325_ = d_2000_v10_
                rhs326_ = (0) - ((self).f34)
                rhs327_ = d_2003_v13_
                rhs328_ = not(d_1989_cf42_)
                rhs329_ = False
                lhs251_ = d_2002_v12_
                lhs252_ = default__.safeIndex(470, (d_2002_v12_).length(0))
                lhs253_ = globalState
                lhs254_ = globalState
                lhs251_[lhs252_] = rhs325_
                lhs253_.f22 = rhs326_
                d_2003_v13_ = rhs327_
                lhs254_.f13 = rhs328_
                r0 = rhs329_
                (globalState).f5 = (self).fm3(globalState)
                nw338_ = _dafny.Array(int(0), 13)
                d_2003_v13_ = nw338_
                d_2005_v14_: _dafny.Array
                nw339_ = _dafny.Array(_dafny.Seq({}), 17)
                d_2005_v14_ = nw339_
                d_2006_v15_: _dafny.Seq
                d_2006_v15_ = _dafny.SeqWithoutIsStrInference([(self.f33)[default__.safeIndex(755, (self.f33).length(0))]])
                index307_ = default__.safeIndex(663, (d_2005_v14_).length(0))
                (d_2005_v14_)[index307_] = (d_2006_v15_) + (d_2006_v15_)
                index308_ = default__.safeIndex(663, (d_2005_v14_).length(0))
                (d_2005_v14_)[index308_] = (d_2006_v15_) + (d_2006_v15_)
            elif True:
                (globalState).f16 = (d_1991_cf40_ if d_1984_v0_ else d_1990_cf41_)
                d_2007_v16_: _dafny.Seq
                d_2007_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srpfggo"))
                d_2008_v17_: _dafny.Seq
                d_2008_v17_ = _dafny.SeqWithoutIsStrInference([False])
                d_2009_v18_: _dafny.Map
                d_2009_v18_ = _dafny.Map({d_2007_v16_: (_dafny.MultiSet(d_2008_v17_)).cardinality})
                d_2010_v19_: _dafny.Map
                d_2010_v19_ = _dafny.Map({d_2009_v18_: default__.fm0((self).f34, globalState)})
                d_2011_v21_: _dafny.Seq
                d_2011_v21_ = _dafny.SeqWithoutIsStrInference([d_2007_v16_])
                def iife193_():
                    coll77_ = _dafny.Map()
                    compr_77_: _dafny.Seq
                    for compr_77_ in (d_2011_v21_).Elements:
                        d_2012_v20_: _dafny.Seq = compr_77_
                        if (d_2012_v20_) in (d_2011_v21_):
                            coll77_[d_2012_v20_] = d_1991_cf40_
                    return _dafny.Map(coll77_)
                d_2010_v19_ = (d_2010_v19_).set(iife193_()
                , d_1984_v0_)
                r2 = d_1989_cf42_
                (globalState).f22 = 42
                r2 = d_1989_cf42_
            d_2013_v22_: _dafny.Array
            nw340_ = _dafny.Array(int(0), 13)
            d_2013_v22_ = nw340_
            d_2013_v22_ = d_2013_v22_
        elif True:
            d_2014___mcc_h3_ = source28_.cf39
            d_2015_cf39_ = d_2014___mcc_h3_
            d_2016_v23_: str
            d_2016_v23_ = _dafny.CodePoint('y')
            d_2017_v25_: _dafny.Map
            d_2017_v25_ = _dafny.Map({default__.fm2(globalState): d_1984_v0_})
            d_2018_v26_: D9
            def iife194_():
                coll78_ = _dafny.Map()
                compr_78_: _dafny.Map
                for compr_78_ in (_dafny.SeqWithoutIsStrInference([d_2017_v25_, d_2017_v25_])).Elements:
                    d_2019_v24_: _dafny.Map = compr_78_
                    if (d_2019_v24_) in (_dafny.SeqWithoutIsStrInference([d_2017_v25_, d_2017_v25_])):
                        coll78_[d_2019_v24_] = _dafny.Map({d_1984_v0_: (self).f34})
                return _dafny.Map(coll78_)
            d_2018_v26_ = D9_DC22(iife194_()
)
            rhs330_ = d_2016_v23_
            rhs331_ = d_2018_v26_
            d_2016_v23_ = rhs330_
            d_2018_v26_ = rhs331_
            d_2020_v27_: D8
            d_2020_v27_ = D8_DC18(default__.fm42(globalState))
            d_2021_v28_: _dafny.Set
            d_2021_v28_ = _dafny.Set({(self).f34, (self).f34})
            (globalState).f1 = (((d_2020_v27_).cf27).isdisjoint(d_2021_v28_)) or (not(d_1984_v0_))
            d_2022_v29_: _dafny.Map
            d_2022_v29_ = _dafny.Map({d_1984_v0_: d_2016_v23_})
            d_2022_v29_ = (d_2022_v29_).set((d_1984_v0_ if d_1984_v0_ else d_1984_v0_), d_2016_v23_)
            d_1984_v0_ = d_1984_v0_
        d_2023_v30_: _dafny.Seq
        d_2023_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whytt"))
        d_2024_v31_: _dafny.Array
        nw341_ = _dafny.Array(int(0), 3)
        d_2024_v31_ = nw341_
        d_2025_v32_: _dafny.Map
        d_2025_v32_ = _dafny.Map({d_2024_v31_: d_2023_v30_})
        d_2026_v33_: str
        d_2026_v33_ = _dafny.CodePoint('j')
        d_2027_v34_: _dafny.Seq
        d_2027_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cl"))])
        d_2028_v35_: _dafny.Seq
        d_2028_v35_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lflvu")), (d_2027_v34_)[default__.safeIndex((self).f34, len(d_2027_v34_))], d_2023_v30_])
        d_2029_v36_: _dafny.Array
        nw342_ = _dafny.Array(None, 29)
        nw342_[int(0)] = (d_2023_v30_) + (d_2023_v30_)
        nw342_[int(1)] = ((d_2025_v32_)[d_2024_v31_] if (d_2024_v31_) in (d_2025_v32_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2030_i1_ in range(default__.abs(489))]))
        nw342_[int(2)] = (default__.fm30((self).f34, d_2023_v30_, (self).f34, globalState)).set(default__.safeIndex((self).fm3(globalState), len(default__.fm30((self).f34, d_2023_v30_, (self).f34, globalState))), d_2026_v33_)
        nw342_[int(3)] = (d_2023_v30_) + (d_2023_v30_)
        nw342_[int(4)] = ((d_2027_v34_)[default__.safeIndex((self).f34, len(d_2027_v34_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjs")))
        nw342_[int(5)] = d_2023_v30_
        nw342_[int(6)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjebxpj"))) + (default__.fm30((self).f34, d_2023_v30_, (self).f34, globalState))
        nw342_[int(7)] = (d_2023_v30_) + (d_2023_v30_)
        nw342_[int(8)] = d_2023_v30_
        nw342_[int(9)] = (d_2023_v30_) + (_dafny.SeqWithoutIsStrInference([d_2026_v33_ for d_2031_i2_ in range(default__.abs(365))]))
        nw342_[int(10)] = d_2023_v30_
        nw342_[int(11)] = d_2023_v30_
        nw342_[int(12)] = d_2023_v30_
        nw342_[int(13)] = d_2023_v30_
        nw342_[int(14)] = _dafny.SeqWithoutIsStrInference([d_2026_v33_ for d_2032_i3_ in range(default__.abs(875))])
        nw342_[int(15)] = d_2023_v30_
        nw342_[int(16)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")) if d_1984_v0_ else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2033_i4_ in range(default__.abs(317))]))
        nw342_[int(17)] = (d_2023_v30_).set(default__.safeIndex((self).f34, len(d_2023_v30_)), d_2026_v33_)
        nw342_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shwka"))
        nw342_[int(19)] = (default__.fm30((self).f34, (d_2023_v30_).set(default__.safeIndex((self).f34, len(d_2023_v30_)), d_2026_v33_), (self).f34, globalState)) + (d_2023_v30_)
        nw342_[int(20)] = d_2023_v30_
        nw342_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efcxkrb"))
        nw342_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljillqg"))
        nw342_[int(23)] = d_2023_v30_
        nw342_[int(24)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2034_i5_ in range(default__.abs(902))])) + (d_2023_v30_)
        nw342_[int(25)] = d_2023_v30_
        nw342_[int(26)] = d_2023_v30_
        nw342_[int(27)] = d_2023_v30_
        nw342_[int(28)] = default__.fm30((self).f34, d_2023_v30_, 80, globalState)
        d_2029_v36_ = nw342_
        index309_ = default__.safeIndex(553, (d_2029_v36_).length(0))
        (d_2029_v36_)[index309_] = (d_2023_v30_) + (d_2023_v30_)
        index310_ = default__.safeIndex(553, (d_2029_v36_).length(0))
        (d_2029_v36_)[index310_] = (d_2023_v30_) + (d_2023_v30_)
        d_2035_v37_: _dafny.Map
        d_2035_v37_ = _dafny.Map({not(d_1984_v0_): d_1984_v0_})
        d_2036_v38_: _dafny.Seq
        d_2036_v38_ = _dafny.SeqWithoutIsStrInference([d_1984_v0_, not(d_1984_v0_), ((d_2035_v37_)[d_1984_v0_] if (d_1984_v0_) in (d_2035_v37_) else d_1984_v0_)])
        (globalState).f13 = ((d_1984_v0_ if d_1984_v0_ else False)) == (not((d_2036_v38_)[default__.safeIndex((self).f34, len(d_2036_v38_))]))
        (globalState).f22 = (self).f34
        index311_ = default__.safeIndex(427, (d_2024_v31_).length(0))
        (d_2024_v31_)[index311_] = (self).f34
        d_2037_v39_: _dafny.Array
        nw343_ = _dafny.Array(_dafny.CodePoint('D'), 29)
        d_2037_v39_ = nw343_
        index312_ = default__.safeIndex(755, (d_2037_v39_).length(0))
        (d_2037_v39_)[index312_] = d_2026_v33_
        index313_ = default__.safeIndex(427, (d_2024_v31_).length(0))
        index314_ = default__.safeIndex(755, (d_2037_v39_).length(0))
        rhs332_ = (0) - ((0) - ((self).f34))
        rhs333_ = (self).f34
        rhs334_ = d_2026_v33_
        lhs255_ = d_2024_v31_
        lhs256_ = default__.safeIndex(427, (d_2024_v31_).length(0))
        lhs257_ = globalState
        lhs258_ = d_2037_v39_
        lhs259_ = default__.safeIndex(755, (d_2037_v39_).length(0))
        lhs255_[lhs256_] = rhs332_
        lhs257_.f0 = rhs333_
        lhs258_[lhs259_] = rhs334_
        d_2038_v40_: C5
        nw344_ = C5()
        nw344_.ctor__(len((d_2029_v36_)[default__.safeIndex(553, (d_2029_v36_).length(0))]))
        d_2038_v40_ = nw344_
        d_2039_v41_: _dafny.Seq
        d_2039_v41_ = _dafny.SeqWithoutIsStrInference([d_2038_v40_, d_2038_v40_])
        d_2040_v42_: _dafny.Seq
        d_2040_v42_ = _dafny.SeqWithoutIsStrInference([d_2039_v41_, d_2039_v41_])
        r0 = (d_1984_v0_ if (d_2040_v42_) < (d_2040_v42_) else d_1984_v0_)
        d_2041_v44_: _dafny.MultiSet
        d_2041_v44_ = _dafny.MultiSet([d_1984_v0_])
        d_2042_v45_: _dafny.Set
        d_2042_v45_ = _dafny.Set({(self).f34, (d_2041_v44_).cardinality, (d_2038_v40_).f38, -185})
        d_2043_v47_: D21
        d_2043_v47_ = D21_DC42(self.f26)
        def iife195_():
            coll79_ = _dafny.Map()
            compr_79_: int
            for compr_79_ in (d_2042_v45_).Elements:
                d_2044_v43_: int = compr_79_
                if (d_2044_v43_) in (d_2042_v45_):
                    coll79_[default__.safeDivisionInt(d_2044_v43_, (self).f34)] = (d_2038_v40_).f38
            return _dafny.Map(coll79_)
        def iife196_():
            coll80_ = _dafny.Map()
            compr_80_: int
            for compr_80_ in _dafny.IntegerRange(352, 960):
                d_2045_v46_: int = compr_80_
                if ((352) <= (d_2045_v46_)) and ((d_2045_v46_) < (960)):
                    coll80_[default__.safeDivisionInt(d_2045_v46_, 615)] = (d_2024_v31_)[default__.safeIndex(427, (d_2024_v31_).length(0))]
            return _dafny.Map(coll80_)
        r1 = ((_dafny.Set({self.f26, self.f26})) - (_dafny.Set({iife195_()
        , self.f26, _dafny.Map({(d_2024_v31_)[default__.safeIndex(427, (d_2024_v31_).length(0))]: len(_dafny.Set({d_1984_v0_, not(d_1984_v0_)}))}), iife196_()
        , self.f26}))) - (_dafny.Set({(d_2043_v47_).cf58, self.f26, self.f26}))
        d_2046_v48_: _dafny.Seq
        d_2046_v48_ = _dafny.SeqWithoutIsStrInference([self.f33])
        r2 = ((_dafny.SeqWithoutIsStrInference([self.f33, self.f33])) + (d_2046_v48_)) == ((d_2046_v48_) + (d_2046_v48_))
        return r0, r1, r2

    @property
    def f34(self):
        return self._f34

class C10(T1):
    def  __init__(self):
        self._f31: int = int(0)
        self._f32: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    def ctor__(self, f31, f32):
        (self)._f31 = f31
        (self)._f32 = f32

    def fm8(self, p0, globalState):
        return ((_dafny.Map({(self).f31: _dafny.Set({len(_dafny.Map({(self).f32: (self).f31})), (self).f31})})) | (_dafny.Map({(self).f31: _dafny.Set({(self).f31})}))) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f31, (self).f31, (self).f31])): _dafny.Set({(0) - (-246)})})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjhl"))): _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtiuel")))})})))

    def fm10(self, p0, p1, p2, globalState):
        def iife197_():
            coll81_ = _dafny.Map()
            compr_81_: int
            for compr_81_ in (_dafny.Map({(self).f31: (self).f32})).keys.Elements:
                d_2047_v0_: int = compr_81_
                if (d_2047_v0_) in (_dafny.Map({(self).f31: (self).f32})):
                    coll81_[default__.safeDivisionInt(d_2047_v0_, (self).f31)] = True
            return _dafny.Map(coll81_)
        return ((_dafny.Map({(self).f31: (self).f32})) | (_dafny.Map({(self).f31: (D5_DC12(False, (self).f32)).cf17}))) | (iife197_()
        )

    def m7(self, p0, globalState):
        r0: bool = False
        d_2048_v0_: _dafny.Seq
        d_2048_v0_ = _dafny.SeqWithoutIsStrInference([67, p0, 606, p0, p0])
        d_2049_v1_: _dafny.Map
        d_2049_v1_ = _dafny.Map({(self).f32: (0) - (len(d_2048_v0_))})
        d_2050_i0_: int
        d_2050_i0_ = 0
        with _dafny.label("14"):
            while (943) > (default__.safeModuloInt(len(d_2049_v1_), p0)):
                with _dafny.c_label("14"):
                    if (d_2050_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2050_i0_ = (d_2050_i0_) + (1)
                    d_2051_v2_: D5
                    d_2051_v2_ = D5_DC12((self).f32, (self).f32)
                    source29_ = d_2051_v2_
                    if source29_.is_DC12:
                        d_2052___mcc_h0_ = source29_.cf17
                        d_2053___mcc_h1_ = source29_.cf18
                        d_2054_cf18_ = d_2053___mcc_h1_
                        d_2055_cf17_ = d_2052___mcc_h0_
                        d_2056_v3_: str
                        d_2056_v3_ = _dafny.CodePoint('g')
                        d_2056_v3_ = d_2056_v3_
                        d_2057_v4_: _dafny.Seq
                        d_2057_v4_ = _dafny.SeqWithoutIsStrInference([d_2055_cf17_, d_2055_cf17_, d_2055_cf17_, (self).f32])
                        d_2058_v5_: _dafny.Array
                        nw345_ = _dafny.Array(None, 4)
                        nw345_[int(0)] = (self).f31
                        nw345_[int(1)] = p0
                        nw345_[int(2)] = len(d_2057_v4_)
                        nw345_[int(3)] = 604
                        d_2058_v5_ = nw345_
                        d_2059_v6_: _dafny.Map
                        d_2059_v6_ = _dafny.Map({p0: 426})
                        d_2060_v7_: T0
                        nw346_ = C3()
                        nw346_.ctor__(d_2058_v5_, False, d_2059_v6_)
                        d_2060_v7_ = nw346_
                        d_2061_v8_: _dafny.Seq
                        d_2061_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgfrioqs"))
                        d_2062_v9_: _dafny.Map
                        d_2062_v9_ = _dafny.Map({d_2060_v7_: d_2061_v8_})
                        d_2063_v10_: _dafny.Map
                        d_2063_v10_ = _dafny.Map({(0) - (len(((d_2062_v9_)[d_2060_v7_] if (d_2060_v7_) in (d_2062_v9_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxtp"))))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yy"))})
                        d_2064_v11_: _dafny.Map
                        d_2064_v11_ = _dafny.Map({d_2063_v10_: (len(default__.fm30((self).f31, d_2061_v8_, (d_2060_v7_).fm3(globalState), globalState)) if d_2055_cf17_ else (self).f31)})
                        d_2064_v11_ = (d_2064_v11_).set(d_2063_v10_, p0)
                        d_2065_v12_: D19
                        d_2065_v12_ = D19_DC39(D19_DC38(p0, not(d_2054_cf18_), (self).f32))
                        d_2066_v13_: D19
                        d_2066_v13_ = D19_DC38(615, d_2055_cf17_, d_2055_cf17_)
                        pat_let_tv57_ = d_2066_v13_
                        d_2067_v14_: _dafny.Seq
                        def iife198_(_pat_let58_0):
                            def iife199_(d_2068_dt__update__tmp_h0_):
                                def iife200_(_pat_let59_0):
                                    def iife201_(d_2069_dt__update_hcf55_h0_):
                                        return D19_DC39(d_2069_dt__update_hcf55_h0_)
                                    return iife201_(_pat_let59_0)
                                return iife200_(pat_let_tv57_)
                            return iife199_(_pat_let58_0)
                        d_2067_v14_ = _dafny.SeqWithoutIsStrInference([d_2065_v12_, d_2065_v12_, d_2065_v12_, iife198_(d_2065_v12_), D19_DC39(d_2066_v13_)])
                        index315_ = default__.safeIndex(496, (d_2058_v5_).length(0))
                        (d_2058_v5_)[index315_] = len(d_2067_v14_)
                        d_2070_v15_: _dafny.Map
                        d_2070_v15_ = _dafny.Map({(self).f31: False})
                        d_2071_v16_: _dafny.Map
                        d_2071_v16_ = _dafny.Map({d_2055_cf17_: d_2070_v15_})
                        d_2072_v17_: _dafny.Map
                        d_2072_v17_ = d_2070_v15_
                        d_2073_v19_: _dafny.Seq
                        def iife202_():
                            coll82_ = _dafny.Map()
                            compr_82_: int
                            for compr_82_ in _dafny.IntegerRange(120, 400):
                                d_2074_v18_: int = compr_82_
                                if ((120) <= (d_2074_v18_)) and ((d_2074_v18_) < (400)):
                                    coll82_[default__.safeModuloInt(d_2074_v18_, p0)] = (self).f32
                            return _dafny.Map(coll82_)
                        d_2073_v19_ = _dafny.SeqWithoutIsStrInference([((d_2071_v16_)[d_2055_cf17_] if (d_2055_cf17_) in (d_2071_v16_) else (d_2072_v17_)), d_2070_v15_, ((d_2070_v15_).set(-501, d_2054_cf18_)) | (d_2070_v15_), (d_2070_v15_) | (iife202_()
                        )])
                        d_2075_v21_: D3
                        d_2075_v21_ = D3_DC6(_dafny.Map({d_2057_v4_: d_2055_cf17_}))
                        d_2076_v22_: _dafny.MultiSet
                        def iife203_():
                            coll83_ = _dafny.Map()
                            compr_83_: _dafny.Seq
                            for compr_83_ in (default__.fm52(globalState)).Elements:
                                d_2077_v20_: _dafny.Seq = compr_83_
                                if (d_2077_v20_) in (default__.fm52(globalState)):
                                    coll83_[d_2077_v20_] = not(d_2055_cf17_)
                            return _dafny.Map(coll83_)
                        d_2076_v22_ = _dafny.MultiSet([D3_DC6(iife203_()
), d_2075_v21_])
                        d_2078_v24_: _dafny.Map
                        d_2078_v24_ = _dafny.Map({d_2049_v1_: (self).f32})
                        index316_ = default__.safeIndex(496, (d_2058_v5_).length(0))
                        def iife204_():
                            coll84_ = _dafny.Map()
                            compr_84_: int
                            for compr_84_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_2054_cf18_, (self).f32})) for d_2079_i1_ in range(default__.abs(739))])).Elements:
                                d_2080_v23_: int = compr_84_
                                if (d_2080_v23_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_2054_cf18_, (self).f32})) for d_2079_i1_ in range(default__.abs(739))])):
                                    coll84_[default__.safeDivisionInt(d_2080_v23_, (self).f31)] = d_2055_cf17_
                            return _dafny.Map(coll84_)
                        rhs335_ = (D4_DC10(len(d_2061_v8_))).cf15
                        rhs336_ = _dafny.SeqWithoutIsStrInference([d_2070_v15_, iife204_()
                        , d_2070_v15_])
                        rhs337_ = ((_dafny.MultiSet([D7_DC16(len(d_2078_v24_), (self).f32)])).cardinality) - ((self).f31)
                        rhs338_ = ((_dafny.SeqWithoutIsStrInference([d_2056_v3_ for d_2081_i2_ in range(default__.abs(-510))])) + (d_2061_v8_)) + ((d_2061_v8_) + (d_2061_v8_))
                        rhs339_ = d_2076_v22_
                        lhs260_ = d_2058_v5_
                        lhs261_ = default__.safeIndex(496, (d_2058_v5_).length(0))
                        lhs262_ = globalState
                        lhs260_[lhs261_] = rhs335_
                        d_2073_v19_ = rhs336_
                        lhs262_.f2 = rhs337_
                        d_2061_v8_ = rhs338_
                        d_2076_v22_ = rhs339_
                        d_2082_v25_: _dafny.Array
                        nw347_ = _dafny.Array(None, 2)
                        nw347_[int(0)] = d_2055_cf17_
                        nw347_[int(1)] = (True) and (not(d_2055_cf17_))
                        d_2082_v25_ = nw347_
                        index317_ = default__.safeIndex(830, (d_2082_v25_).length(0))
                        (d_2082_v25_)[index317_] = not(not((self).f32))
                        index318_ = default__.safeIndex(830, (d_2082_v25_).length(0))
                        (d_2082_v25_)[index318_] = ((0) - (p0)) != ((d_2058_v5_)[default__.safeIndex(496, (d_2058_v5_).length(0))])
                    elif source29_.is_DC13:
                        d_2083___mcc_h2_ = source29_.cf19
                        d_2084_cf19_ = d_2083___mcc_h2_
                        d_2085_v26_: _dafny.Array
                        def lambda98_(d_2086_cf19_):
                            def lambda99_(d_2087_i3_):
                                return _dafny.Set({(self).f31, d_2086_cf19_})

                            return lambda99_

                        init57_ = lambda98_(d_2084_cf19_)
                        nw348_ = _dafny.Array(None, 21)
                        for i0_57_ in range(nw348_.length(0)):
                            nw348_[i0_57_] = init57_(i0_57_)
                        d_2085_v26_ = nw348_
                        d_2088_v27_: _dafny.Set
                        d_2088_v27_ = _dafny.Set({d_2084_cf19_})
                        index319_ = default__.safeIndex(876, (d_2085_v26_).length(0))
                        (d_2085_v26_)[index319_] = d_2088_v27_
                        index320_ = default__.safeIndex(876, (d_2085_v26_).length(0))
                        (d_2085_v26_)[index320_] = d_2088_v27_
                        d_2089_v28_: _dafny.Seq
                        d_2089_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bggng"))
                        d_2090_v29_: D4
                        d_2090_v29_ = D4_DC10(p0)
                        d_2091_v30_: _dafny.Seq
                        d_2091_v30_ = _dafny.SeqWithoutIsStrInference([(self).f32])
                        d_2092_v31_: D3
                        d_2092_v31_ = D3_DC7(d_2089_v28_, (d_2090_v29_).cf15, (self).f31, d_2084_cf19_, d_2091_v30_)
                        d_2093_v32_: str
                        d_2093_v32_ = _dafny.CodePoint('u')
                        (globalState).f13 = ((d_2089_v28_) + (((d_2092_v31_).cf8).set(default__.safeIndex((0) - (d_2084_cf19_), len((d_2092_v31_).cf8)), d_2093_v32_))) == (((_dafny.SeqWithoutIsStrInference([d_2093_v32_ for d_2094_i4_ in range(default__.abs(437))])) + (_dafny.SeqWithoutIsStrInference([d_2093_v32_ for d_2095_i5_ in range(default__.abs(412))]))).set(default__.safeIndex(d_2084_cf19_, len((_dafny.SeqWithoutIsStrInference([d_2093_v32_ for d_2096_i4_ in range(default__.abs(437))])) + (_dafny.SeqWithoutIsStrInference([d_2093_v32_ for d_2097_i5_ in range(default__.abs(412))])))), d_2093_v32_))
                        d_2098_v33_: _dafny.Array
                        nw349_ = _dafny.Array(False, 14)
                        d_2098_v33_ = nw349_
                        index321_ = default__.safeIndex(115, (d_2098_v33_).length(0))
                        (d_2098_v33_)[index321_] = False
                        index322_ = default__.safeIndex(115, (d_2098_v33_).length(0))
                        (d_2098_v33_)[index322_] = default__.fm0(((self).f31) - ((self).f31), globalState)
                        (globalState).f1 = (d_2098_v33_)[default__.safeIndex(115, (d_2098_v33_).length(0))]
                    elif True:
                        d_2099___mcc_h3_ = source29_.cf16
                        d_2100_cf16_ = d_2099___mcc_h3_
                        d_2101_v34_: _dafny.Seq
                        d_2101_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylrtvd"))
                        d_2101_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2102_i6_ in range(default__.abs(-268))])
                        (globalState).f21 = (-304) + (len(d_2048_v0_))
                        (globalState).f13 = (p0) <= (p0)
                        (globalState).f1 = (self).f32
                    d_2103_v35_: _dafny.Array
                    nw350_ = _dafny.Array(_dafny.Array(None, 0), 5)
                    d_2103_v35_ = nw350_
                    d_2104_v36_: _dafny.Array
                    def lambda100_(d_2105_i7_):
                        return (self).f32

                    init58_ = lambda100_
                    nw351_ = _dafny.Array(None, 20)
                    for i0_58_ in range(nw351_.length(0)):
                        nw351_[i0_58_] = init58_(i0_58_)
                    d_2104_v36_ = nw351_
                    index323_ = default__.safeIndex(879, (d_2103_v35_).length(0))
                    (d_2103_v35_)[index323_] = d_2104_v36_
                    index324_ = default__.safeIndex(879, (d_2103_v35_).length(0))
                    (d_2103_v35_)[index324_] = (d_2104_v36_ if not(True) else d_2104_v36_)
                    d_2106_v37_: _dafny.Set
                    d_2106_v37_ = _dafny.Set({p0})
                    d_2107_v38_: _dafny.Seq
                    d_2107_v38_ = _dafny.SeqWithoutIsStrInference([(self).f32, (len(d_2106_v37_)) not in (_dafny.Map({(self).f31: (self).f32})), (self).f32, not((self).f32), not((self).f32)])
                    (globalState).f13 = (d_2107_v38_)[default__.safeIndex((0) - (p0), len(d_2107_v38_))]
                    r0 = not ((self).f32) or ((self).f32)
                    pass
            pass
        (globalState).f5 = (self).f31
        d_2108_v39_: _dafny.Array
        nw352_ = _dafny.Array(int(0), 4)
        d_2108_v39_ = nw352_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2108_v39_).length(0)):
            d_2109_i8_: int = guard_loop_8_
            if (True) and (((0) <= (d_2109_i8_)) and ((d_2109_i8_) < ((d_2108_v39_).length(0)))):
                (d_2108_v39_)[(d_2109_i8_)] = default__.safeModuloInt(d_2109_i8_, len((_dafny.Map({_dafny.Map({p0: (self).f32}): (self).f31}) if (self).f32 else _dafny.Map({_dafny.Map({(self).f31: (self).f32}): (self).f31}))))
        d_2110_v40_: _dafny.MultiSet
        d_2110_v40_ = _dafny.MultiSet([True])
        d_2111_v41_: _dafny.Seq
        d_2111_v41_ = _dafny.SeqWithoutIsStrInference([True])
        d_2112_i9_: int
        d_2112_i9_ = 0
        with _dafny.label("15"):
            while ((d_2110_v40_) | (d_2110_v40_)).ispropersubset((d_2110_v40_ if (self).f32 else _dafny.MultiSet(d_2111_v41_))):
                with _dafny.c_label("15"):
                    if (d_2112_i9_) >= (100):
                        raise _dafny.Break("15")
                    d_2112_i9_ = (d_2112_i9_) + (1)
                    d_2113_v42_: _dafny.Array
                    nw353_ = _dafny.Array(False, 6)
                    d_2113_v42_ = nw353_
                    index325_ = default__.safeIndex(619, (d_2113_v42_).length(0))
                    (d_2113_v42_)[index325_] = ((self).f31) == (p0)
                    d_2114_v43_: _dafny.MultiSet
                    d_2114_v43_ = _dafny.MultiSet([d_2108_v39_, d_2108_v39_])
                    index326_ = default__.safeIndex(619, (d_2113_v42_).length(0))
                    (d_2113_v42_)[index326_] = (d_2114_v43_).issubset(_dafny.MultiSet([d_2108_v39_, d_2108_v39_]))
                    d_2115_v44_: _dafny.Seq
                    d_2115_v44_ = _dafny.SeqWithoutIsStrInference([d_2108_v39_, d_2108_v39_])
                    d_2116_v45_: _dafny.Array
                    nw354_ = _dafny.Array(None, 23)
                    nw354_[int(0)] = d_2108_v39_
                    nw354_[int(1)] = d_2108_v39_
                    nw354_[int(2)] = d_2108_v39_
                    nw354_[int(3)] = d_2108_v39_
                    nw354_[int(4)] = d_2108_v39_
                    nw354_[int(5)] = d_2108_v39_
                    nw354_[int(6)] = d_2108_v39_
                    nw354_[int(7)] = d_2108_v39_
                    nw354_[int(8)] = d_2108_v39_
                    nw354_[int(9)] = d_2108_v39_
                    nw354_[int(10)] = d_2108_v39_
                    nw354_[int(11)] = (d_2115_v44_)[default__.safeIndex(p0, len(d_2115_v44_))]
                    nw354_[int(12)] = d_2108_v39_
                    nw354_[int(13)] = d_2108_v39_
                    nw354_[int(14)] = d_2108_v39_
                    nw354_[int(15)] = (d_2108_v39_ if (self).f32 else d_2108_v39_)
                    nw354_[int(16)] = d_2108_v39_
                    nw354_[int(17)] = d_2108_v39_
                    nw354_[int(18)] = d_2108_v39_
                    nw354_[int(19)] = d_2108_v39_
                    nw354_[int(20)] = d_2108_v39_
                    nw354_[int(21)] = d_2108_v39_
                    nw354_[int(22)] = d_2108_v39_
                    d_2116_v45_ = nw354_
                    index327_ = default__.safeIndex(482, (d_2116_v45_).length(0))
                    (d_2116_v45_)[index327_] = d_2108_v39_
                    index328_ = default__.safeIndex(482, (d_2116_v45_).length(0))
                    (d_2116_v45_)[index328_] = (d_2115_v44_)[default__.safeIndex((0) - ((self).f31), len(d_2115_v44_))]
                    d_2117_v46_: _dafny.Seq
                    d_2117_v46_ = _dafny.SeqWithoutIsStrInference([d_2110_v40_])
                    d_2118_v47_: _dafny.Set
                    d_2118_v47_ = _dafny.Set({len((d_2049_v1_) | (d_2049_v1_)), default__.safeModuloInt(len(d_2117_v46_), p0), (self).f31})
                    d_2118_v47_ = d_2118_v47_
                    (globalState).f21 = 26
                    pass
            pass
        if not((self).f32):
            d_2119_v48_: C2
            nw355_ = C2()
            nw355_.ctor__((self).f32, (self).f31)
            d_2119_v48_ = nw355_
            (globalState).f1 = (self).f32
            d_2120_v49_: _dafny.Seq
            d_2120_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vltn"))
            if (d_2111_v41_)[default__.safeIndex((len(_dafny.Set({d_2120_v49_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")), d_2120_v49_}))) * ((self).f31), len(d_2111_v41_))]:
                d_2111_v41_ = d_2111_v41_
                (globalState).f1 = (d_2119_v48_).f41
                d_2121_v50_: _dafny.Array
                nw356_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_2121_v50_ = nw356_
                d_2122_v51_: _dafny.Array
                nw357_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_2122_v51_ = nw357_
                index329_ = default__.safeIndex(272, (d_2121_v50_).length(0))
                (d_2121_v50_)[index329_] = d_2122_v51_
                index330_ = default__.safeIndex(272, (d_2121_v50_).length(0))
                (d_2121_v50_)[index330_] = d_2122_v51_
                d_2123_v52_: _dafny.MultiSet
                d_2123_v52_ = _dafny.MultiSet([p0])
                (globalState).f0 = ((d_2123_v52_).intersection((_dafny.MultiSet([(self).f31])).intersection(_dafny.MultiSet([(0) - ((self).f31), (d_2119_v48_).f42])))).cardinality
                (globalState).f0 = (self).f31
            elif True:
                d_2124_v53_: _dafny.Map
                d_2124_v53_ = _dafny.Map({d_2110_v40_: (d_2119_v48_).f41})
                d_2125_v54_: _dafny.Set
                d_2125_v54_ = _dafny.Set({True, ((d_2124_v53_)[d_2110_v40_] if (d_2110_v40_) in (d_2124_v53_) else (d_2119_v48_).f41), (self).f32})
                d_2126_v55_: D10
                d_2126_v55_ = D10_DC24(d_2125_v54_)
                d_2127_v56_: _dafny.Map
                d_2127_v56_ = _dafny.Map({(d_2126_v55_).cf34: d_2111_v41_})
                d_2127_v56_ = default__.fm53((not((self).f32)) == (True), globalState)
                d_2128_v57_: _dafny.Array
                def lambda101_(d_2129_i10_):
                    return _dafny.CodePoint('j')

                init59_ = lambda101_
                nw358_ = _dafny.Array(None, 2)
                for i0_59_ in range(nw358_.length(0)):
                    nw358_[i0_59_] = init59_(i0_59_)
                d_2128_v57_ = nw358_
                index331_ = default__.safeIndex(653, (d_2128_v57_).length(0))
                (d_2128_v57_)[index331_] = default__.fm35((d_2119_v48_).f42, globalState)
                index332_ = default__.safeIndex(653, (d_2128_v57_).length(0))
                (d_2128_v57_)[index332_] = (d_2120_v49_)[default__.safeIndex((d_2119_v48_).f42, len(d_2120_v49_))]
                index333_ = default__.safeIndex(400, (d_2108_v39_).length(0))
                (d_2108_v39_)[index333_] = (self).f31
                index334_ = default__.safeIndex(400, (d_2108_v39_).length(0))
                (d_2108_v39_)[index334_] = (0) - (len(default__.fm42(globalState)))
                (globalState).f22 = (d_2108_v39_)[default__.safeIndex(400, (d_2108_v39_).length(0))]
                d_2130_v58_: D20
                d_2130_v58_ = D20_DC40(_dafny.Set({(d_2128_v57_)[default__.safeIndex(653, (d_2128_v57_).length(0))], (d_2128_v57_)[default__.safeIndex(653, (d_2128_v57_).length(0))]}))
                d_2131_v59_: D7
                d_2131_v59_ = D7_DC17(not((self).f32), (d_2119_v48_).f41, p0)
                d_2132_v60_: _dafny.Map
                d_2132_v60_ = _dafny.Map({(self).f32: d_2120_v49_})
                d_2133_v61_: _dafny.MultiSet
                d_2133_v61_ = _dafny.MultiSet([(d_2120_v49_) + (((d_2132_v60_)[(self).f32] if ((self).f32) in (d_2132_v60_) else d_2120_v49_))])
                rhs340_ = ((d_2133_v61_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxu"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxu"))) in (d_2133_v61_) else (p0 if (d_2119_v48_).f41 else 695))
                rhs341_ = d_2130_v58_
                rhs342_ = (d_2119_v48_).f41
                rhs343_ = (D4_DC9((self).f31)).cf14
                rhs344_ = d_2131_v59_
                lhs263_ = globalState
                lhs264_ = globalState
                lhs265_ = globalState
                lhs263_.f2 = rhs340_
                d_2130_v58_ = rhs341_
                lhs264_.f13 = rhs342_
                lhs265_.f5 = rhs343_
                d_2131_v59_ = rhs344_
            d_2134_v62_: _dafny.Set
            d_2134_v62_ = _dafny.Set({p0, p0})
            d_2135_v63_: _dafny.MultiSet
            d_2135_v63_ = _dafny.MultiSet([d_2134_v62_])
            d_2136_v64_: _dafny.Map
            d_2136_v64_ = _dafny.Map({(d_2135_v63_).cardinality: (self).f31})
            d_2136_v64_ = (d_2136_v64_).set((d_2048_v0_)[default__.safeIndex(((d_2049_v1_)[(self).f32] if ((self).f32) in (d_2049_v1_) else (self).f31), len(d_2048_v0_))], p0)
            d_2137_v65_: _dafny.Map
            d_2137_v65_ = _dafny.Map({d_2111_v41_: True})
            d_2138_v66_: _dafny.Map
            d_2138_v66_ = _dafny.Map({(self).fm10((d_2119_v48_).f41, D3_DC6(d_2137_v65_), (d_2119_v48_).f41, globalState): d_2049_v1_})
            d_2139_v67_: D9
            d_2139_v67_ = D9_DC22(d_2138_v66_)
            d_2139_v67_ = D9_DC22(d_2138_v66_)
        elif True:
            d_2140_v68_: _dafny.Seq
            d_2140_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evhpvgw"))
            d_2141_v69_: C7
            nw359_ = C7()
            nw359_.ctor__((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2142_i11_ in range(default__.abs(974))])) < (d_2140_v68_))
            d_2141_v69_ = nw359_
            d_2143_v70_: str
            d_2143_v70_ = _dafny.CodePoint('e')
            (globalState).f2 = (0) - (default__.safeDivisionInt((self).f31, len((d_2140_v68_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0, len(d_2140_v68_)])), len(d_2140_v68_)), d_2143_v70_))))
            d_2144_v71_: D8
            d_2144_v71_ = D8_DC19(not((d_2141_v69_).f37), (d_2141_v69_).f37, (self).f32)
            d_2145_v72_: C7
            nw360_ = C7()
            nw360_.ctor__(((d_2111_v41_).set(default__.safeIndex((self).f31, len(d_2111_v41_)), (d_2141_v69_).f37)) == (_dafny.SeqWithoutIsStrInference([(self).f32, (d_2144_v71_).cf30])))
            d_2145_v72_ = nw360_
            r0 = (d_2141_v69_).f37
            d_2146_v73_: D10
            d_2146_v73_ = D10_DC25(p0, ((d_2049_v1_)[(d_2111_v41_)[default__.safeIndex((self).f31, len(d_2111_v41_))]] if ((d_2111_v41_)[default__.safeIndex((self).f31, len(d_2111_v41_))]) in (d_2049_v1_) else p0), (d_2141_v69_).f37, (d_2141_v69_).f37)
            d_2147_v74_: D5
            d_2147_v74_ = D5_DC11((d_2146_v73_).cf37)
            source30_ = d_2147_v74_
            if source30_.is_DC12:
                d_2148___mcc_h4_ = source30_.cf17
                d_2149___mcc_h5_ = source30_.cf18
                d_2150_cf18_ = d_2149___mcc_h5_
                d_2151_cf17_ = d_2148___mcc_h4_
                d_2152_v75_: _dafny.MultiSet
                d_2152_v75_ = _dafny.MultiSet([(self).f31])
                d_2153_v76_: _dafny.Map
                d_2153_v76_ = _dafny.Map({(d_2152_v75_) | (_dafny.MultiSet([(self).f31])): (d_2141_v69_).f37})
                d_2153_v76_ = (d_2153_v76_).set((_dafny.MultiSet([len(d_2140_v68_)])) - (d_2152_v75_), d_2151_cf17_)
                d_2154_v77_: C7
                nw361_ = C7()
                nw361_.ctor__((d_2141_v69_).f37)
                d_2154_v77_ = nw361_
                d_2155_v78_: _dafny.Set
                d_2155_v78_ = _dafny.Set({(self).f31})
                d_2155_v78_ = d_2155_v78_
                (globalState).f5 = (self).f31
            elif source30_.is_DC13:
                d_2156___mcc_h6_ = source30_.cf19
                d_2157_cf19_ = d_2156___mcc_h6_
                d_2158_v79_: _dafny.Array
                nw362_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_2158_v79_ = nw362_
                d_2159_v80_: _dafny.Array
                nw363_ = _dafny.Array(False, 21)
                d_2159_v80_ = nw363_
                index335_ = default__.safeIndex(14, (d_2158_v79_).length(0))
                (d_2158_v79_)[index335_] = d_2159_v80_
                index336_ = default__.safeIndex(14, (d_2158_v79_).length(0))
                (d_2158_v79_)[index336_] = d_2159_v80_
                d_2048_v0_ = _dafny.SeqWithoutIsStrInference([(d_2157_cf19_) * ((0) - (len(d_2048_v0_))) for d_2160_i12_ in range(default__.abs(-808))])
                (globalState).f13 = (d_2141_v69_).f37
                arr10_ = (d_2158_v79_)[default__.safeIndex(14, (d_2158_v79_).length(0))]
                index337_ = default__.safeIndex(869, ((d_2158_v79_)[default__.safeIndex(14, (d_2158_v79_).length(0))]).length(0))
                arr10_[index337_] = (d_2111_v41_)[default__.safeIndex(p0, len(d_2111_v41_))]
                d_2161_v81_: _dafny.Seq
                d_2161_v81_ = _dafny.SeqWithoutIsStrInference([d_2140_v68_, d_2140_v68_])
                d_2162_v82_: _dafny.Set
                d_2162_v82_ = _dafny.Set({(d_2141_v69_).f37, (d_2141_v69_).f37})
                d_2163_v83_: _dafny.Map
                d_2163_v83_ = _dafny.Map({len((d_2161_v81_)[default__.safeIndex(d_2157_cf19_, len(d_2161_v81_))]): d_2162_v82_})
                d_2164_v84_: _dafny.Map
                d_2164_v84_ = _dafny.Map({(_dafny.Set({(d_2145_v72_).f37})) - (((d_2163_v83_)[len(d_2162_v82_)] if (len(d_2162_v82_)) in (d_2163_v83_) else d_2162_v82_)): (d_2162_v82_).isdisjoint(d_2162_v82_)})
                d_2165_v85_: _dafny.Map
                d_2165_v85_ = _dafny.Map({(d_2141_v69_).f37: (p0) >= (d_2157_cf19_)})
                d_2166_v86_: _dafny.Set
                d_2166_v86_ = _dafny.Set({551, (d_2141_v69_).fm19((d_2048_v0_)[default__.safeIndex(p0, len(d_2048_v0_))], p0, globalState)})
                arr11_ = (d_2158_v79_)[default__.safeIndex(14, (d_2158_v79_).length(0))]
                index338_ = default__.safeIndex(869, ((d_2158_v79_)[default__.safeIndex(14, (d_2158_v79_).length(0))]).length(0))
                rhs345_ = (d_2145_v72_).f37
                rhs346_ = d_2164_v84_
                rhs347_ = ((d_2165_v85_)[not(not((d_2166_v86_).issubset(d_2166_v86_)))] if (not(not((d_2166_v86_).issubset(d_2166_v86_)))) in (d_2165_v85_) else (d_2141_v69_).f37)
                lhs266_ = (d_2158_v79_)[default__.safeIndex(14, (d_2158_v79_).length(0))]
                lhs267_ = default__.safeIndex(869, ((d_2158_v79_)[default__.safeIndex(14, (d_2158_v79_).length(0))]).length(0))
                lhs268_ = globalState
                lhs266_[lhs267_] = rhs345_
                d_2164_v84_ = rhs346_
                lhs268_.f1 = rhs347_
            elif True:
                d_2167___mcc_h7_ = source30_.cf16
                d_2168_cf16_ = d_2167___mcc_h7_
                d_2169_v87_: _dafny.Map
                d_2169_v87_ = _dafny.Map({(d_2145_v72_).f37: d_2168_cf16_})
                d_2170_v88_: _dafny.Map
                d_2170_v88_ = _dafny.Map({len(d_2169_v87_): not(((self).f31) == ((self).f31))})
                d_2170_v88_ = ((d_2170_v88_).set((self).f31, ((d_2170_v88_)[p0] if (p0) in (d_2170_v88_) else (d_2145_v72_).f37))) | (d_2170_v88_)
                d_2171_v89_: _dafny.Map
                d_2171_v89_ = _dafny.Map({d_2110_v40_: p0})
                d_2172_v90_: D12
                d_2172_v90_ = D12_DC28(d_2110_v40_)
                d_2171_v89_ = (d_2171_v89_).set(((d_2172_v90_).cf43) - (_dafny.MultiSet(default__.fm13((self).f31, globalState))), (self).f31)
                (globalState).f21 = p0
                d_2173_v91_: _dafny.Array
                nw364_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_2173_v91_ = nw364_
                index339_ = default__.safeIndex(538, (d_2173_v91_).length(0))
                (d_2173_v91_)[index339_] = d_2140_v68_
                index340_ = default__.safeIndex(538, (d_2173_v91_).length(0))
                (d_2173_v91_)[index340_] = (default__.fm54(p0, globalState)).cf8
        d_2174_v92_: _dafny.Seq
        d_2174_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rc"))
        d_2175_v93_: D5
        d_2175_v93_ = D5_DC13(len(d_2174_v92_))
        d_2176_v94_: _dafny.Seq
        d_2176_v94_ = _dafny.SeqWithoutIsStrInference([d_2175_v93_])
        d_2176_v94_ = d_2176_v94_
        d_2177_v95_: _dafny.Array
        nw365_ = _dafny.Array(False, 4)
        d_2177_v95_ = nw365_
        d_2178_v96_: _dafny.MultiSet
        d_2178_v96_ = _dafny.MultiSet([d_2177_v95_, d_2177_v95_])
        d_2179_v97_: _dafny.Map
        d_2179_v97_ = _dafny.Map({(False) and ((self).f32): (d_2178_v96_).ispropersubset(d_2178_v96_)})
        r0 = ((d_2179_v97_)[(self).f32] if ((self).f32) in (d_2179_v97_) else (self).f32)
        return r0

    def m10(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: bool = False
        d_2180_v0_: _dafny.Seq
        d_2180_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2181_v1_: D5
        d_2181_v1_ = D5_DC13((self).f31)
        d_2182_v2_: _dafny.Seq
        d_2182_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hy"))
        r1 = (d_2182_v2_ if not((d_2180_v0_)[default__.safeIndex((d_2181_v1_).cf19, len(d_2180_v0_))]) else d_2182_v2_)
        d_2180_v0_ = d_2180_v0_
        (globalState).f1 = (self).f32
        d_2183_v3_: D10
        d_2183_v3_ = D10_DC25(p1, p1, (self).f32, p0)
        d_2184_v4_: _dafny.Map
        d_2184_v4_ = _dafny.Map({(d_2183_v3_).cf36: p1})
        d_2185_v5_: _dafny.Map
        d_2185_v5_ = _dafny.Map({d_2184_v4_: 58})
        hi11_ = len(d_2185_v5_)
        for d_2186_i0_ in range((self).f31, hi11_):
            (globalState).f13 = (d_2180_v0_)[default__.safeIndex(p1, len(d_2180_v0_))]
            d_2187_v6_: _dafny.Array
            nw366_ = _dafny.Array(int(0), 1)
            d_2187_v6_ = nw366_
            d_2188_v7_: C3
            nw367_ = C3()
            nw367_.ctor__(d_2187_v6_, (self).f32, default__.fm45(globalState))
            d_2188_v7_ = nw367_
            d_2189_v8_: _dafny.Array
            def lambda102_(d_2190_v7_):
                def lambda103_(d_2191_i1_):
                    return (_dafny.SeqWithoutIsStrInference([(self).f31 for d_2192_i2_ in range(default__.abs(625))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: d_2190_v7_.f40})), 297]))

                return lambda103_

            init60_ = lambda102_(d_2188_v7_)
            nw368_ = _dafny.Array(None, 5)
            for i0_60_ in range(nw368_.length(0)):
                nw368_[i0_60_] = init60_(i0_60_)
            d_2189_v8_ = nw368_
            d_2193_v9_: D3
            d_2193_v9_ = D3_DC7(d_2182_v2_, (self).f31, (0) - (p1), p1, d_2180_v0_)
            d_2194_v10_: _dafny.MultiSet
            d_2194_v10_ = _dafny.MultiSet([_dafny.CodePoint('g')])
            pat_let_tv58_ = d_2182_v2_
            pat_let_tv59_ = d_2194_v10_
            index341_ = default__.safeIndex(708, (d_2189_v8_).length(0))
            def iife205_(_pat_let60_0):
                def iife206_(d_2195_dt__update__tmp_h0_):
                    def iife207_(_pat_let61_0):
                        def iife208_(d_2196_dt__update_hcf8_h0_):
                            def iife209_(_pat_let62_0):
                                def iife210_(d_2197_dt__update_hcf9_h0_):
                                    return D3_DC7(d_2196_dt__update_hcf8_h0_, d_2197_dt__update_hcf9_h0_, (d_2195_dt__update__tmp_h0_).cf10, (d_2195_dt__update__tmp_h0_).cf11, (d_2195_dt__update__tmp_h0_).cf12)
                                return iife210_(_pat_let62_0)
                            return iife209_((pat_let_tv59_).cardinality)
                        return iife208_(_pat_let61_0)
                    return iife207_(pat_let_tv58_)
                return iife206_(_pat_let60_0)
            (d_2189_v8_)[index341_] = default__.fm21(_dafny.CodePoint('p'), p1, iife205_(d_2193_v9_), globalState)
            d_2198_v11_: _dafny.Seq
            d_2198_v11_ = _dafny.SeqWithoutIsStrInference([p1, d_2186_i0_])
            index342_ = default__.safeIndex(708, (d_2189_v8_).length(0))
            (d_2189_v8_)[index342_] = (d_2198_v11_) + (d_2198_v11_)
            (globalState).f22 = (0) - (((self).f31) - ((d_2181_v1_).cf19))
        pat_let_tv60_ = p1
        pat_let_tv61_ = p1
        d_2199_v12_: _dafny.Array
        nw369_ = _dafny.Array(None, 27)
        def iife211_(_pat_let63_0):
            def iife212_(d_2200_dt__update__tmp_h1_):
                def iife213_(_pat_let64_0):
                    def iife214_(d_2201_dt__update_hcf19_h0_):
                        return D5_DC13(d_2201_dt__update_hcf19_h0_)
                    return iife214_(_pat_let64_0)
                return iife213_((self).f31)
            return iife212_(_pat_let63_0)
        nw369_[int(0)] = iife211_(d_2181_v1_)
        nw369_[int(1)] = d_2181_v1_
        nw369_[int(2)] = D5_DC13(len(d_2180_v0_))
        nw369_[int(3)] = d_2181_v1_
        nw369_[int(4)] = d_2181_v1_
        nw369_[int(5)] = d_2181_v1_
        nw369_[int(6)] = d_2181_v1_
        def iife215_(_pat_let65_0):
            def iife216_(d_2202_dt__update__tmp_h2_):
                def iife217_(_pat_let66_0):
                    def iife218_(d_2203_dt__update_hcf19_h1_):
                        return D5_DC13(d_2203_dt__update_hcf19_h1_)
                    return iife218_(_pat_let66_0)
                return iife217_((self).f31)
            return iife216_(_pat_let65_0)
        nw369_[int(7)] = iife215_(d_2181_v1_)
        nw369_[int(8)] = d_2181_v1_
        nw369_[int(9)] = d_2181_v1_
        nw369_[int(10)] = D5_DC13(len(d_2180_v0_))
        nw369_[int(11)] = d_2181_v1_
        nw369_[int(12)] = d_2181_v1_
        def iife219_(_pat_let67_0):
            def iife220_(d_2204_dt__update__tmp_h3_):
                def iife221_(_pat_let68_0):
                    def iife222_(d_2205_dt__update_hcf19_h2_):
                        return D5_DC13(d_2205_dt__update_hcf19_h2_)
                    return iife222_(_pat_let68_0)
                return iife221_((self).f31)
            return iife220_(_pat_let67_0)
        nw369_[int(13)] = iife219_(d_2181_v1_)
        nw369_[int(14)] = d_2181_v1_
        nw369_[int(15)] = d_2181_v1_
        nw369_[int(16)] = d_2181_v1_
        nw369_[int(17)] = d_2181_v1_
        nw369_[int(18)] = d_2181_v1_
        def iife223_(_pat_let69_0):
            def iife224_(d_2206_dt__update__tmp_h4_):
                def iife225_(_pat_let70_0):
                    def iife226_(d_2207_dt__update_hcf19_h3_):
                        return D5_DC13(d_2207_dt__update_hcf19_h3_)
                    return iife226_(_pat_let70_0)
                return iife225_(pat_let_tv60_)
            return iife224_(_pat_let69_0)
        nw369_[int(19)] = iife223_(d_2181_v1_)
        nw369_[int(20)] = default__.fm23(globalState)
        nw369_[int(21)] = d_2181_v1_
        nw369_[int(22)] = d_2181_v1_
        nw369_[int(23)] = d_2181_v1_
        nw369_[int(24)] = (d_2181_v1_ if p0 else D5_DC13(p1))
        nw369_[int(25)] = default__.fm23(globalState)
        def iife227_(_pat_let71_0):
            def iife228_(d_2208_dt__update__tmp_h5_):
                def iife229_(_pat_let72_0):
                    def iife230_(d_2209_dt__update_hcf19_h4_):
                        return D5_DC13(d_2209_dt__update_hcf19_h4_)
                    return iife230_(_pat_let72_0)
                return iife229_(pat_let_tv61_)
            return iife228_(_pat_let71_0)
        nw369_[int(26)] = iife227_(d_2181_v1_)
        d_2199_v12_ = nw369_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_2199_v12_).length(0)):
            d_2210_i3_: int = guard_loop_9_
            if (True) and (((0) <= (d_2210_i3_)) and ((d_2210_i3_) < ((d_2199_v12_).length(0)))):
                (d_2199_v12_)[(d_2210_i3_)] = d_2181_v1_
        d_2211_v13_: _dafny.Array
        def lambda104_(d_2212_i6_):
            return (d_2212_i6_) + ((self).f31)

        init61_ = lambda104_
        nw370_ = _dafny.Array(None, 24)
        for i0_61_ in range(nw370_.length(0)):
            nw370_[i0_61_] = init61_(i0_61_)
        d_2211_v13_ = nw370_
        d_2213_v14_: _dafny.Map
        d_2213_v14_ = _dafny.Map({d_2182_v2_: d_2211_v13_})
        d_2214_i4_: int
        d_2214_i4_ = 0
        with _dafny.label("16"):
            while (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2264_i5_ in range(default__.abs(-309))])) in (d_2213_v14_):
                with _dafny.c_label("16"):
                    if (d_2214_i4_) >= (100):
                        raise _dafny.Break("16")
                    d_2214_i4_ = (d_2214_i4_) + (1)
                    (globalState).f1 = p0
                    d_2215_v15_: _dafny.Array
                    def lambda105_(d_2216_i7_):
                        return (self).f32

                    init62_ = lambda105_
                    nw371_ = _dafny.Array(None, 5)
                    for i0_62_ in range(nw371_.length(0)):
                        nw371_[i0_62_] = init62_(i0_62_)
                    d_2215_v15_ = nw371_
                    d_2217_v16_: _dafny.MultiSet
                    d_2217_v16_ = _dafny.MultiSet([d_2215_v15_])
                    d_2218_v17_: _dafny.Map
                    d_2218_v17_ = _dafny.Map({(d_2217_v16_).cardinality: p0})
                    d_2219_v18_: D8
                    d_2219_v18_ = D8_DC20(default__.fm0(len(d_2218_v17_), globalState))
                    def iife231_(_pat_let73_0):
                        def iife232_(d_2220_dt__update__tmp_h6_):
                            def iife233_(_pat_let74_0):
                                def iife234_(d_2221_dt__update_hcf31_h0_):
                                    return D8_DC20(d_2221_dt__update_hcf31_h0_)
                                return iife234_(_pat_let74_0)
                            return iife233_((self).f32)
                        return iife232_(_pat_let73_0)
                    source31_ = iife231_(d_2219_v18_)
                    if source31_.is_DC19:
                        d_2222___mcc_h0_ = source31_.cf28
                        d_2223___mcc_h1_ = source31_.cf29
                        d_2224___mcc_h2_ = source31_.cf30
                        d_2225_cf30_ = d_2224___mcc_h2_
                        d_2226_cf29_ = d_2223___mcc_h1_
                        d_2227_cf28_ = d_2222___mcc_h0_
                        d_2228_v19_: str
                        d_2228_v19_ = _dafny.CodePoint('e')
                        d_2228_v19_ = d_2228_v19_
                        (globalState).f5 = ((self).f31) - (p1)
                        d_2229_v20_: _dafny.Map
                        d_2229_v20_ = _dafny.Map({len(d_2180_v0_): d_2182_v2_})
                        d_2230_v21_: C5
                        nw372_ = C5()
                        nw372_.ctor__(306)
                        d_2230_v21_ = nw372_
                        d_2231_v22_: _dafny.Map
                        d_2231_v22_ = _dafny.Map({d_2229_v20_: d_2230_v21_})
                        def iife235_():
                            def iife236_():
                                coll86_ = _dafny.Map()
                                compr_86_: int
                                for compr_86_ in _dafny.IntegerRange(889, -605):
                                    d_2233_v24_: int = compr_86_
                                    if ((889) <= (d_2233_v24_)) and ((d_2233_v24_) < (-605)):
                                        coll86_[(d_2233_v24_) - (p1)] = (self).f31
                                return _dafny.Map(coll86_)
                            coll85_ = _dafny.Map()
                            compr_85_: int
                            for compr_85_ in _dafny.IntegerRange(6, 573):
                                d_2232_v23_: int = compr_85_
                                if ((6) <= (d_2232_v23_)) and ((d_2232_v23_) < (573)):
                                    coll85_[(d_2232_v23_) - (len(iife236_()
                                    ))] = d_2182_v2_
                            return _dafny.Map(coll85_)
                        d_2231_v22_ = (d_2231_v22_).set((iife235_()
                        ) | (d_2229_v20_), d_2230_v21_)
                        d_2234_v25_: _dafny.MultiSet
                        d_2234_v25_ = _dafny.MultiSet([(d_2230_v21_).f38, (_dafny.MultiSet([d_2228_v19_])).cardinality])
                        d_2235_v26_: _dafny.MultiSet
                        d_2235_v26_ = _dafny.MultiSet([p0, default__.fm0((d_2234_v25_).cardinality, globalState), d_2227_cf28_])
                        d_2235_v26_ = d_2235_v26_
                    elif source31_.is_DC20:
                        d_2236___mcc_h3_ = source31_.cf31
                        d_2237_cf31_ = d_2236___mcc_h3_
                        d_2238_v27_: _dafny.Array
                        def lambda106_(d_2239_v0_):
                            def lambda107_(d_2240_i8_):
                                return d_2239_v0_

                            return lambda107_

                        init63_ = lambda106_(d_2180_v0_)
                        nw373_ = _dafny.Array(None, 23)
                        for i0_63_ in range(nw373_.length(0)):
                            nw373_[i0_63_] = init63_(i0_63_)
                        d_2238_v27_ = nw373_
                        index343_ = default__.safeIndex(768, (d_2238_v27_).length(0))
                        (d_2238_v27_)[index343_] = d_2180_v0_
                        index344_ = default__.safeIndex(768, (d_2238_v27_).length(0))
                        (d_2238_v27_)[index344_] = d_2180_v0_
                        d_2215_v15_ = d_2215_v15_
                        (globalState).f13 = ((d_2238_v27_)[default__.safeIndex(768, (d_2238_v27_).length(0))])[default__.safeIndex((self).f31, len((d_2238_v27_)[default__.safeIndex(768, (d_2238_v27_).length(0))]))]
                        index345_ = default__.safeIndex(404, (d_2211_v13_).length(0))
                        (d_2211_v13_)[index345_] = p1
                        d_2241_v28_: D3
                        d_2241_v28_ = D3_DC7(d_2182_v2_, p1, default__.fm2(globalState), (self).f31, d_2180_v0_)
                        index346_ = default__.safeIndex(404, (d_2211_v13_).length(0))
                        (d_2211_v13_)[index346_] = len(default__.fm30(p1, (d_2241_v28_).cf8, p1, globalState))
                    elif source31_.is_DC18:
                        d_2242___mcc_h4_ = source31_.cf27
                        d_2243_cf27_ = d_2242___mcc_h4_
                        d_2244_v29_: D9
                        d_2244_v29_ = D9_DC23()
                        d_2245_v30_: _dafny.Array
                        nw374_ = _dafny.Array(None, 29)
                        nw374_[int(0)] = d_2244_v29_
                        nw374_[int(1)] = d_2244_v29_
                        nw374_[int(2)] = d_2244_v29_
                        nw374_[int(3)] = d_2244_v29_
                        nw374_[int(4)] = d_2244_v29_
                        nw374_[int(5)] = d_2244_v29_
                        nw374_[int(6)] = d_2244_v29_
                        nw374_[int(7)] = d_2244_v29_
                        nw374_[int(8)] = d_2244_v29_
                        nw374_[int(9)] = d_2244_v29_
                        nw374_[int(10)] = D9_DC23()
                        nw374_[int(11)] = d_2244_v29_
                        nw374_[int(12)] = d_2244_v29_
                        nw374_[int(13)] = d_2244_v29_
                        nw374_[int(14)] = default__.fm55((self).f32, len(d_2182_v2_), p1, globalState)
                        nw374_[int(15)] = d_2244_v29_
                        nw374_[int(16)] = default__.fm55(p0, p1, p1, globalState)
                        nw374_[int(17)] = d_2244_v29_
                        nw374_[int(18)] = d_2244_v29_
                        nw374_[int(19)] = d_2244_v29_
                        nw374_[int(20)] = d_2244_v29_
                        nw374_[int(21)] = d_2244_v29_
                        nw374_[int(22)] = d_2244_v29_
                        nw374_[int(23)] = D9_DC23()
                        nw374_[int(24)] = d_2244_v29_
                        nw374_[int(25)] = d_2244_v29_
                        nw374_[int(26)] = d_2244_v29_
                        nw374_[int(27)] = d_2244_v29_
                        nw374_[int(28)] = d_2244_v29_
                        d_2245_v30_ = nw374_
                        index347_ = default__.safeIndex(629, (d_2245_v30_).length(0))
                        (d_2245_v30_)[index347_] = d_2244_v29_
                        index348_ = default__.safeIndex(629, (d_2245_v30_).length(0))
                        (d_2245_v30_)[index348_] = d_2244_v29_
                        d_2246_v31_: _dafny.MultiSet
                        d_2246_v31_ = _dafny.MultiSet([p0, (self).f32, (self).f32, p0])
                        d_2246_v31_ = (d_2246_v31_).intersection(d_2246_v31_)
                        rhs348_ = (default__.safeModuloInt(44, len(default__.fm30((self).f31, d_2182_v2_, p1, globalState)))) < ((self).f31)
                        rhs349_ = p1
                        lhs269_ = globalState
                        r0 = rhs348_
                        lhs269_.f21 = rhs349_
                        d_2247_v32_: _dafny.Seq
                        d_2247_v32_ = _dafny.SeqWithoutIsStrInference([d_2184_v4_])
                        d_2248_v33_: _dafny.Map
                        d_2248_v33_ = _dafny.Map({d_2215_v15_: (d_2247_v32_)[default__.safeIndex((0) - ((self).f31), len(d_2247_v32_))]})
                        d_2249_v34_: _dafny.Seq
                        d_2249_v34_ = _dafny.SeqWithoutIsStrInference([(self).f31, (0) - (len(((d_2248_v33_)[d_2215_v15_] if (d_2215_v15_) in (d_2248_v33_) else d_2184_v4_))), (self).f31, len(_dafny.Map({p1: default__.fm2(globalState)})), (self).f31])
                        r0 = (d_2249_v34_) < ((d_2249_v34_).set(default__.safeIndex(-89, len(d_2249_v34_)), p1))
                    elif True:
                        d_2250___mcc_h5_ = source31_.cf32
                        d_2251_cf32_ = d_2250___mcc_h5_
                        (globalState).f0 = ((self).f31) - (default__.safeModuloInt((self).f31, p1))
                        d_2252_v35_: str
                        d_2252_v35_ = _dafny.CodePoint('f')
                        d_2253_v36_: _dafny.Set
                        d_2253_v36_ = _dafny.Set({d_2252_v35_})
                        d_2254_v37_: _dafny.Seq
                        d_2254_v37_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_2252_v35_}), d_2253_v36_, d_2253_v36_])
                        d_2255_v39_: _dafny.Map
                        d_2255_v39_ = _dafny.Map({(self).f32: _dafny.CodePoint('y')})
                        d_2256_v40_: _dafny.Map
                        def iife237_():
                            coll87_ = _dafny.Set()
                            compr_87_: str
                            for compr_87_ in (d_2182_v2_).Elements:
                                d_2257_v38_: str = compr_87_
                                if (d_2257_v38_) in (d_2182_v2_):
                                    coll87_ = coll87_.union(_dafny.Set([d_2257_v38_]))
                            return _dafny.Set(coll87_)
                        def iife238_():
                            coll88_ = _dafny.Set()
                            compr_88_: str
                            for compr_88_ in (d_2182_v2_).Elements:
                                d_2258_v38_: str = compr_88_
                                if (d_2258_v38_) in (d_2182_v2_):
                                    coll88_ = coll88_.union(_dafny.Set([d_2258_v38_]))
                            return _dafny.Set(coll88_)
                        d_2256_v40_ = _dafny.Map({(d_2254_v37_) + ((_dafny.SeqWithoutIsStrInference([iife237_()
                        , _dafny.Set({default__.fm35((self).f31, globalState)}), d_2253_v36_, d_2253_v36_])).set(default__.safeIndex((self).f31, len(_dafny.SeqWithoutIsStrInference([iife238_()
                        , _dafny.Set({default__.fm35((self).f31, globalState)}), d_2253_v36_, d_2253_v36_]))), d_2253_v36_)): ((self).f32) not in (d_2255_v39_)})
                        d_2259_v41_: _dafny.Seq
                        d_2259_v41_ = _dafny.SeqWithoutIsStrInference([d_2254_v37_, _dafny.SeqWithoutIsStrInference([default__.fm15((self).f32, globalState), _dafny.Set({d_2252_v35_, d_2252_v35_})])])
                        d_2256_v40_ = (d_2256_v40_).set((d_2259_v41_)[default__.safeIndex(p1, len(d_2259_v41_))], p0)
                        index349_ = default__.safeIndex(758, (d_2215_v15_).length(0))
                        (d_2215_v15_)[index349_] = (d_2253_v36_).ispropersubset(d_2253_v36_)
                        index350_ = default__.safeIndex(758, (d_2215_v15_).length(0))
                        (d_2215_v15_)[index350_] = p0
                        d_2260_v42_: D7
                        d_2260_v42_ = D7_DC17((d_2215_v15_)[default__.safeIndex(758, (d_2215_v15_).length(0))], False, len(d_2180_v0_))
                        index351_ = default__.safeIndex(758, (d_2215_v15_).length(0))
                        (d_2215_v15_)[index351_] = ((d_2260_v42_).cf26) > ((self).f31)
                    d_2261_v43_: _dafny.Map
                    d_2261_v43_ = _dafny.Map({d_2182_v2_: (self).f31})
                    d_2261_v43_ = (d_2261_v43_).set(d_2182_v2_, (self).f31)
                    d_2262_v44_: _dafny.Set
                    d_2262_v44_ = _dafny.Set({(self).f32, False})
                    d_2263_v45_: _dafny.Seq
                    d_2263_v45_ = _dafny.SeqWithoutIsStrInference([d_2262_v44_, d_2262_v44_])
                    (globalState).f16 = ((len((d_2263_v45_)[default__.safeIndex((self).f31, len(d_2263_v45_))])) * ((self).f31)) - (p1)
                    pass
            pass
        r0 = not((self).f32)
        d_2265_v46_: str
        d_2265_v46_ = _dafny.CodePoint('f')
        d_2266_v47_: _dafny.Map
        d_2266_v47_ = _dafny.Map({d_2182_v2_: d_2265_v46_})
        r1 = default__.fm30(default__.safeModuloInt(len(d_2266_v47_), p1), d_2182_v2_, p1, globalState)
        d_2267_v48_: _dafny.MultiSet
        d_2267_v48_ = _dafny.MultiSet([d_2182_v2_])
        d_2268_v49_: _dafny.Seq
        d_2268_v49_ = _dafny.SeqWithoutIsStrInference([(d_2267_v48_).cardinality, default__.fm2(globalState)])
        d_2269_v50_: _dafny.Set
        d_2269_v50_ = _dafny.Set({(self).f31, p1, len((d_2268_v49_).set(default__.safeIndex((d_2268_v49_)[default__.safeIndex(len(d_2182_v2_), len(d_2268_v49_))], len(d_2268_v49_)), p1)), (self).f31, (self).f31})
        r2 = ((d_2269_v50_) - (_dafny.Set({(d_2268_v49_)[default__.safeIndex(p1, len(d_2268_v49_))]}))).isdisjoint(d_2269_v50_)
        return r0, r1, r2

    def m11(self, p0, p1, p2, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_2270_v0_: _dafny.Seq
        d_2270_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2271_i0_: int
        d_2271_i0_ = 0
        with _dafny.label("17"):
            while not (((self).f31) <= (len((d_2270_v0_).set(default__.safeIndex((self).f31, len(d_2270_v0_)), p1)))) or (p2):
                with _dafny.c_label("17"):
                    if (d_2271_i0_) >= (100):
                        raise _dafny.Break("17")
                    d_2271_i0_ = (d_2271_i0_) + (1)
                    d_2272_v1_: bool
                    d_2273_v2_: bool
                    d_2274_v3_: _dafny.Map
                    d_2275_v4_: int
                    out82_: bool
                    out83_: bool
                    out84_: _dafny.Map
                    out85_: int
                    out82_, out83_, out84_, out85_ = default__.m0(globalState)
                    d_2272_v1_ = out82_
                    d_2273_v2_ = out83_
                    d_2274_v3_ = out84_
                    d_2275_v4_ = out85_
                    d_2276_v5_: _dafny.MultiSet
                    d_2276_v5_ = _dafny.MultiSet([p2])
                    d_2277_v6_: _dafny.Map
                    d_2277_v6_ = _dafny.Map({d_2272_v1_: d_2275_v4_})
                    d_2278_v7_: _dafny.Seq
                    d_2278_v7_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.fm2(globalState)), (self).f31, (self).f31, len(d_2277_v6_), 133])
                    d_2279_v8_: _dafny.Set
                    d_2279_v8_ = _dafny.Set({False})
                    d_2280_v9_: _dafny.Seq
                    d_2280_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdsh"))
                    d_2281_v10_: _dafny.Array
                    nw375_ = _dafny.Array(None, 6)
                    nw375_[int(0)] = ((d_2276_v5_)[d_2273_v2_] if (d_2273_v2_) in (d_2276_v5_) else (self).f31)
                    nw375_[int(1)] = (self).f31
                    nw375_[int(2)] = len(_dafny.Map({True: (self).f32}))
                    nw375_[int(3)] = default__.safeModuloInt((d_2278_v7_)[default__.safeIndex(len(d_2279_v8_), len(d_2278_v7_))], d_2275_v4_)
                    nw375_[int(4)] = (self).f31
                    nw375_[int(5)] = len(d_2280_v9_)
                    d_2281_v10_ = nw375_
                    d_2282_v11_: _dafny.Map
                    d_2282_v11_ = _dafny.Map({default__.fm0((self).f31, globalState): d_2272_v1_})
                    index352_ = default__.safeIndex(97, (d_2281_v10_).length(0))
                    (d_2281_v10_)[index352_] = default__.safeDivisionInt(len(d_2282_v11_), len(d_2279_v8_))
                    d_2283_v12_: D19
                    d_2283_v12_ = D19_DC38(d_2275_v4_, (self).f32, (self).f32)
                    d_2284_v13_: _dafny.Seq
                    def iife239_(_pat_let75_0):
                        def iife240_(d_2285_dt__update__tmp_h0_):
                            def iife241_(_pat_let76_0):
                                def iife242_(d_2286_dt__update_hcf54_h0_):
                                    def iife243_(_pat_let77_0):
                                        def iife244_(d_2287_dt__update_hcf53_h0_):
                                            return D19_DC38((d_2285_dt__update__tmp_h0_).cf52, d_2287_dt__update_hcf53_h0_, d_2286_dt__update_hcf54_h0_)
                                        return iife244_(_pat_let77_0)
                                    return iife243_((self).f32)
                                return iife242_(_pat_let76_0)
                            return iife241_((self).f32)
                        return iife240_(_pat_let75_0)
                    d_2284_v13_ = _dafny.SeqWithoutIsStrInference([d_2283_v12_, d_2283_v12_, iife239_(d_2283_v12_)])
                    d_2288_v14_: _dafny.Map
                    d_2288_v14_ = _dafny.Map({d_2281_v10_: d_2284_v13_})
                    index353_ = default__.safeIndex(97, (d_2281_v10_).length(0))
                    (d_2281_v10_)[index353_] = len(((d_2288_v14_).set(d_2281_v10_, _dafny.SeqWithoutIsStrInference([d_2283_v12_])) if d_2272_v1_ else d_2288_v14_))
                    d_2289_v15_: D5
                    d_2289_v15_ = D5_DC11(p2)
                    d_2290_v16_: _dafny.Map
                    d_2290_v16_ = _dafny.Map({d_2275_v4_: (d_2281_v10_)[default__.safeIndex(97, (d_2281_v10_).length(0))]})
                    d_2291_v17_: C8
                    nw376_ = C8()
                    nw376_.ctor__(d_2272_v1_, d_2289_v15_, d_2290_v16_)
                    d_2291_v17_ = nw376_
                    d_2292_v18_: _dafny.Map
                    d_2292_v18_ = _dafny.Map({(d_2281_v10_)[default__.safeIndex(97, (d_2281_v10_).length(0))]: True})
                    (globalState).f0 = len(((d_2292_v18_) | (d_2292_v18_)) | ((d_2292_v18_) | (d_2292_v18_)))
                    pass
            pass
        d_2293_v19_: C2
        nw377_ = C2()
        nw377_.ctor__(not(not(p2)), ((self).f31) + ((self).f31))
        d_2293_v19_ = nw377_
        d_2294_v20_: _dafny.Set
        d_2294_v20_ = _dafny.Set({(self).f31})
        d_2295_v21_: D8
        d_2295_v21_ = D8_DC18((_dafny.Set({(d_2293_v19_).f42, 538})) - (d_2294_v20_))
        d_2296_v22_: _dafny.Seq
        d_2296_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpqchg"))
        d_2297_v23_: _dafny.MultiSet
        d_2297_v23_ = _dafny.MultiSet([(self).f32, (D5_DC11(p2)).cf16])
        d_2298_v24_: _dafny.Array
        nw378_ = _dafny.Array(None, 20)
        nw378_[int(0)] = (d_2293_v19_).f42
        nw378_[int(1)] = len(d_2296_v22_)
        nw378_[int(2)] = (d_2293_v19_).f42
        nw378_[int(3)] = len(d_2294_v20_)
        nw378_[int(4)] = (self).f31
        nw378_[int(5)] = (d_2293_v19_).f42
        nw378_[int(6)] = 870
        nw378_[int(7)] = (d_2293_v19_).f42
        nw378_[int(8)] = (d_2293_v19_).f42
        nw378_[int(9)] = (self).f31
        nw378_[int(10)] = (self).f31
        nw378_[int(11)] = (0) - ((self).f31)
        nw378_[int(12)] = (d_2297_v23_).cardinality
        nw378_[int(13)] = (d_2293_v19_).f42
        nw378_[int(14)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2299_i1_ in range(default__.abs(342))]))
        nw378_[int(15)] = (d_2293_v19_).f42
        nw378_[int(16)] = (self).f31
        nw378_[int(17)] = (self).f31
        nw378_[int(18)] = (self).f31
        nw378_[int(19)] = (self).f31
        d_2298_v24_ = nw378_
        d_2300_v25_: _dafny.Map
        d_2300_v25_ = _dafny.Map({(self).f32: (d_2293_v19_).f42})
        d_2301_v26_: _dafny.Map
        d_2301_v26_ = _dafny.Map({d_2298_v24_: d_2300_v25_})
        d_2302_v27_: _dafny.Seq
        d_2302_v27_ = _dafny.SeqWithoutIsStrInference([(self).f31, len(d_2296_v22_)])
        d_2295_v21_ = (d_2295_v21_ if not((d_2293_v19_).f41) else default__.fm56((d_2293_v19_).f42, (d_2293_v19_).f41, len((((d_2301_v26_)[d_2298_v24_] if (d_2298_v24_) in (d_2301_v26_) else d_2300_v25_)).set((self).f32, len(d_2302_v27_))), globalState))
        d_2303_v28_: D5
        d_2303_v28_ = D5_DC12((d_2293_v19_).f41, p2)
        (globalState).f1 = (d_2303_v28_).cf18
        hi12_ = (d_2297_v23_).cardinality
        for d_2304_i2_ in range((d_2293_v19_).f42, hi12_):
            (globalState).f13 = True
            index354_ = default__.safeIndex(44, (p1).length(0))
            (p1)[index354_] = p2
            index355_ = default__.safeIndex(44, (p1).length(0))
            (p1)[index355_] = p2
            d_2305_v29_: _dafny.Array
            nw379_ = _dafny.Array(None, 28)
            nw379_[int(0)] = (d_2293_v19_).f41
            nw379_[int(1)] = (d_2293_v19_).f41
            nw379_[int(2)] = (p1)[default__.safeIndex(44, (p1).length(0))]
            nw379_[int(3)] = (p1)[default__.safeIndex(44, (p1).length(0))]
            nw379_[int(4)] = (d_2293_v19_).f41
            nw379_[int(5)] = (d_2293_v19_).fm24(globalState)
            nw379_[int(6)] = not((d_2293_v19_).f41)
            nw379_[int(7)] = (d_2293_v19_).f41
            nw379_[int(8)] = p2
            nw379_[int(9)] = p2
            nw379_[int(10)] = (p1)[default__.safeIndex(44, (p1).length(0))]
            nw379_[int(11)] = (self).f32
            nw379_[int(12)] = default__.fm0(211, globalState)
            nw379_[int(13)] = p2
            nw379_[int(14)] = True
            nw379_[int(15)] = False
            nw379_[int(16)] = (d_2293_v19_).f41
            nw379_[int(17)] = p2
            nw379_[int(18)] = (self).f32
            nw379_[int(19)] = (self).f32
            nw379_[int(20)] = (self).f32
            nw379_[int(21)] = (d_2293_v19_).f41
            nw379_[int(22)] = (d_2293_v19_).f41
            nw379_[int(23)] = (d_2293_v19_).f41
            nw379_[int(24)] = (d_2293_v19_).f41
            nw379_[int(25)] = p2
            nw379_[int(26)] = default__.fm0(d_2304_i2_, globalState)
            nw379_[int(27)] = (p1)[default__.safeIndex(44, (p1).length(0))]
            d_2305_v29_ = nw379_
            d_2306_v31_: C9
            nw380_ = C9()
            def iife245_():
                coll89_ = _dafny.Map()
                compr_89_: int
                for compr_89_ in _dafny.IntegerRange(746, 383):
                    d_2307_v30_: int = compr_89_
                    if ((746) <= (d_2307_v30_)) and ((d_2307_v30_) < (383)):
                        coll89_[(d_2307_v30_) * ((d_2293_v19_).f42)] = d_2304_i2_
                return _dafny.Map(coll89_)
            nw380_.ctor__(d_2305_v29_, (d_2293_v19_).f42, iife245_()
            )
            d_2306_v31_ = nw380_
            d_2308_v32_: _dafny.Array
            nw381_ = _dafny.Array(_dafny.MultiSet({}), 17)
            d_2308_v32_ = nw381_
            d_2309_v33_: _dafny.MultiSet
            d_2309_v33_ = _dafny.MultiSet([(_dafny.MultiSet([640, (d_2293_v19_).f42, (d_2306_v31_).fm3(globalState), (d_2293_v19_).f42])).cardinality, d_2304_i2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oik"))), (d_2293_v19_).f42])
            index356_ = default__.safeIndex(668, (d_2308_v32_).length(0))
            (d_2308_v32_)[index356_] = d_2309_v33_
            index357_ = default__.safeIndex(668, (d_2308_v32_).length(0))
            rhs350_ = (d_2293_v19_).f42
            rhs351_ = ((d_2306_v31_).f34) - (-870)
            rhs352_ = default__.safeModuloInt((0) - ((d_2306_v31_).f34), ((d_2306_v31_).f34) * ((0) - ((d_2306_v31_).f34)))
            rhs353_ = d_2309_v33_
            lhs270_ = globalState
            lhs271_ = globalState
            lhs272_ = d_2308_v32_
            lhs273_ = default__.safeIndex(668, (d_2308_v32_).length(0))
            lhs270_.f0 = rhs350_
            lhs271_.f2 = rhs351_
            r3 = rhs352_
            lhs272_[lhs273_] = rhs353_
        (globalState).f1 = (d_2293_v19_).f41
        d_2310_v34_: _dafny.Seq
        d_2310_v34_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(self).f31 for d_2311_i3_ in range(default__.abs(910))])).set(default__.safeIndex(553, len(_dafny.SeqWithoutIsStrInference([(self).f31 for d_2312_i3_ in range(default__.abs(910))]))), (d_2297_v23_).cardinality), _dafny.SeqWithoutIsStrInference([(d_2293_v19_).f42]), d_2302_v27_, d_2302_v27_, d_2302_v27_])
        d_2313_v35_: _dafny.MultiSet
        d_2313_v35_ = _dafny.MultiSet([(self).f31, len(d_2310_v34_), (self).f31, (self).f31])
        def iife246_():
            coll90_ = _dafny.Map()
            compr_90_: int
            for compr_90_ in (d_2302_v27_).Elements:
                d_2314_v36_: int = compr_90_
                if (d_2314_v36_) in (d_2302_v27_):
                    coll90_[default__.safeDivisionInt(d_2314_v36_, (self).f31)] = (d_2293_v19_).f42
            return _dafny.Map(coll90_)
        r0 = (d_2313_v35_).set(len(default__.fm29(d_2297_v23_, (0) - (len(iife246_()
        )), (d_2293_v19_).f41, -990, globalState)), default__.abs(-978))
        r1 = default__.safeModuloInt(783, (self).f31)
        r2 = (d_2298_v24_ if (self).f32 else d_2298_v24_)
        r3 = (default__.safeModuloInt((0) - ((0) - ((0) - ((self).f31))), (d_2313_v35_).cardinality)) * ((d_2293_v19_).f42)
        return r0, r1, r2, r3

    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32

class C11(T0, T1):
    def  __init__(self):
        self._f26: _dafny.Map = _dafny.Map({})
        self._f30: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f26(self):
        return self._f26
    @f26.setter
    def f26(self, value):
        self._f26 = value
    def ctor__(self, f30, f26):
        (self)._f30 = f30
        (self).f26 = f26

    def fm3(self, globalState):
        return (self).f30

    def fm4(self, p0, p1, p2, p3, globalState):
        source32_ = (D4_DC10(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnlbgagj")))) if True else D4_DC10((D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qf")), (self).f30, (self).f30, (self).f30, _dafny.SeqWithoutIsStrInference([True, not(False)]))).cf10))
        if source32_.is_DC9:
            d_2315___mcc_h0_ = source32_.cf14
            d_2316_cf14_ = d_2315___mcc_h0_
            return (_dafny.Map({(self).f30: _dafny.Map({_dafny.CodePoint('u'): d_2316_cf14_})})) | (_dafny.Map({len(_dafny.Map({False: False})): _dafny.Map({_dafny.CodePoint('w'): (self).f30})}))
        elif source32_.is_DC10:
            d_2317___mcc_h1_ = source32_.cf15
            d_2318_cf15_ = d_2317___mcc_h1_
            return _dafny.Map({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([True])))): _dafny.Map({_dafny.CodePoint('t'): len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30])]))})})
        elif True:
            d_2319___mcc_h2_ = source32_.cf13
            d_2320_cf13_ = d_2319___mcc_h2_
            def iife247_():
                coll91_ = _dafny.Map()
                compr_91_: str
                for compr_91_ in (_dafny.Set({d_2320_cf13_, d_2320_cf13_, d_2320_cf13_})).Elements:
                    d_2321_v0_: str = compr_91_
                    if (d_2321_v0_) in (_dafny.Set({d_2320_cf13_, d_2320_cf13_, d_2320_cf13_})):
                        coll91_[d_2321_v0_] = (self).f30
                return _dafny.Map(coll91_)
            return _dafny.Map({(self).f30: iife247_()
            })

    def fm8(self, p0, globalState):
        def iife248_():
            coll92_ = _dafny.Map()
            compr_92_: int
            for compr_92_ in _dafny.IntegerRange(-927, -284):
                d_2322_v0_: int = compr_92_
                if ((-927) <= (d_2322_v0_)) and ((d_2322_v0_) < (-284)):
                    coll92_[(d_2322_v0_) - (len(_dafny.SeqWithoutIsStrInference([True])))] = (_dafny.Set({-821, (self).f30, (self).f30, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2323_i0_ in range(default__.abs(91))]))])), (self).f30})) - (_dafny.Set({539}))
            return _dafny.Map(coll92_)
        return iife248_()
        

    def fm9(self, p0, p1, globalState):
        return len(_dafny.Set({(self).f30, (self).f30, len(_dafny.Set({False, False, not(True)}))}))

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        (globalState).f2 = (default__.safeDivisionInt(p1, p2)) * (default__.safeDivisionInt(p1, p2))
        d_2324_v0_: bool
        d_2324_v0_ = False
        d_2325_v1_: str
        d_2325_v1_ = _dafny.CodePoint('v')
        d_2326_v2_: _dafny.Map
        d_2326_v2_ = _dafny.Map({d_2325_v1_: d_2324_v0_})
        d_2327_v3_: _dafny.Array
        nw382_ = _dafny.Array(None, 14)
        nw382_[int(0)] = d_2324_v0_
        nw382_[int(1)] = d_2324_v0_
        nw382_[int(2)] = d_2324_v0_
        nw382_[int(3)] = d_2324_v0_
        nw382_[int(4)] = d_2324_v0_
        nw382_[int(5)] = d_2324_v0_
        nw382_[int(6)] = ((d_2326_v2_)[d_2325_v1_] if (d_2325_v1_) in (d_2326_v2_) else False)
        nw382_[int(7)] = d_2324_v0_
        nw382_[int(8)] = d_2324_v0_
        nw382_[int(9)] = d_2324_v0_
        nw382_[int(10)] = d_2324_v0_
        nw382_[int(11)] = d_2324_v0_
        nw382_[int(12)] = d_2324_v0_
        nw382_[int(13)] = default__.fm0((0) - (p1), globalState)
        d_2327_v3_ = nw382_
        d_2328_v4_: D4
        d_2328_v4_ = D4_DC9((self).f30)
        d_2329_v5_: C9
        nw383_ = C9()
        nw383_.ctor__(d_2327_v3_, (d_2328_v4_).cf14, self.f26)
        d_2329_v5_ = nw383_
        d_2324_v0_ = d_2324_v0_
        d_2330_i0_: int
        d_2330_i0_ = 0
        with _dafny.label("18"):
            while d_2324_v0_:
                with _dafny.c_label("18"):
                    if (d_2330_i0_) >= (100):
                        raise _dafny.Break("18")
                    d_2330_i0_ = (d_2330_i0_) + (1)
                    (globalState).f2 = p2
                    d_2331_v6_: _dafny.MultiSet
                    d_2331_v6_ = _dafny.MultiSet([d_2324_v0_])
                    d_2331_v6_ = d_2331_v6_
                    (globalState).f2 = 782
                    d_2332_v7_: _dafny.Seq
                    d_2332_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xagp"))
                    index358_ = default__.safeIndex(179, (d_2327_v3_).length(0))
                    (d_2327_v3_)[index358_] = (p1) < (p1)
                    d_2333_v8_: D20
                    d_2333_v8_ = D20_DC41((d_2329_v5_).f34)
                    pat_let_tv62_ = d_2329_v5_
                    pat_let_tv63_ = p2
                    d_2334_v9_: _dafny.Array
                    nw384_ = _dafny.Array(None, 23)
                    nw384_[int(0)] = d_2333_v8_
                    nw384_[int(1)] = d_2333_v8_
                    nw384_[int(2)] = d_2333_v8_
                    nw384_[int(3)] = default__.fm57(d_2324_v0_, globalState)
                    nw384_[int(4)] = d_2333_v8_
                    nw384_[int(5)] = (d_2333_v8_ if d_2324_v0_ else d_2333_v8_)
                    nw384_[int(6)] = d_2333_v8_
                    nw384_[int(7)] = d_2333_v8_
                    nw384_[int(8)] = default__.fm57(not(d_2324_v0_), globalState)
                    nw384_[int(9)] = d_2333_v8_
                    nw384_[int(10)] = d_2333_v8_
                    nw384_[int(11)] = d_2333_v8_
                    nw384_[int(12)] = D20_DC41((_dafny.MultiSet([p1])).cardinality)
                    nw384_[int(13)] = d_2333_v8_
                    nw384_[int(14)] = d_2333_v8_
                    nw384_[int(15)] = d_2333_v8_
                    nw384_[int(16)] = d_2333_v8_
                    def iife249_(_pat_let78_0):
                        def iife250_(d_2335_dt__update__tmp_h0_):
                            def iife251_(_pat_let79_0):
                                def iife252_(d_2336_dt__update_hcf57_h0_):
                                    return D20_DC41(d_2336_dt__update_hcf57_h0_)
                                return iife252_(_pat_let79_0)
                            return iife251_((pat_let_tv62_).f34)
                        return iife250_(_pat_let78_0)
                    nw384_[int(17)] = iife249_(d_2333_v8_)
                    nw384_[int(18)] = d_2333_v8_
                    nw384_[int(19)] = d_2333_v8_
                    nw384_[int(20)] = d_2333_v8_
                    nw384_[int(21)] = d_2333_v8_
                    def iife253_(_pat_let80_0):
                        def iife254_(d_2337_dt__update__tmp_h1_):
                            def iife255_(_pat_let81_0):
                                def iife256_(d_2338_dt__update_hcf57_h1_):
                                    return D20_DC41(d_2338_dt__update_hcf57_h1_)
                                return iife256_(_pat_let81_0)
                            return iife255_(pat_let_tv63_)
                        return iife254_(_pat_let80_0)
                    nw384_[int(22)] = iife253_(d_2333_v8_)
                    d_2334_v9_ = nw384_
                    index359_ = default__.safeIndex(587, (d_2334_v9_).length(0))
                    (d_2334_v9_)[index359_] = d_2333_v8_
                    index360_ = default__.safeIndex(179, (d_2327_v3_).length(0))
                    index361_ = default__.safeIndex(587, (d_2334_v9_).length(0))
                    rhs354_ = default__.fm30(-813, default__.fm30(p1, d_2332_v7_, (d_2329_v5_).f34, globalState), (self).fm3(globalState), globalState)
                    rhs355_ = not(d_2324_v0_)
                    rhs356_ = d_2333_v8_
                    lhs274_ = d_2327_v3_
                    lhs275_ = default__.safeIndex(179, (d_2327_v3_).length(0))
                    lhs276_ = d_2334_v9_
                    lhs277_ = default__.safeIndex(587, (d_2334_v9_).length(0))
                    d_2332_v7_ = rhs354_
                    lhs274_[lhs275_] = rhs355_
                    lhs276_[lhs277_] = rhs356_
                    pass
            pass
        d_2339_v10_: _dafny.Array
        nw385_ = _dafny.Array(None, 13)
        nw385_[int(0)] = 616
        nw385_[int(1)] = (((_dafny.MultiSet([d_2324_v0_])).set(True, default__.abs((self).f30))).cardinality) * (992)
        nw385_[int(2)] = (d_2329_v5_).f34
        nw385_[int(3)] = ((self).f30) + ((self).f30)
        nw385_[int(4)] = ((self).f30) - ((self).f30)
        nw385_[int(5)] = (0) - ((0) - (p1))
        nw385_[int(6)] = (self).fm9((self).f30, (0) - ((self).f30), globalState)
        nw385_[int(7)] = p1
        nw385_[int(8)] = p2
        nw385_[int(9)] = p2
        nw385_[int(10)] = (d_2329_v5_).fm3(globalState)
        nw385_[int(11)] = 187
        nw385_[int(12)] = (d_2329_v5_).f34
        d_2339_v10_ = nw385_
        index362_ = default__.safeIndex(144, (d_2339_v10_).length(0))
        (d_2339_v10_)[index362_] = p1
        index363_ = default__.safeIndex(144, (d_2339_v10_).length(0))
        (d_2339_v10_)[index363_] = p2
        d_2340_v11_: _dafny.Seq
        d_2340_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwfy"))
        hi13_ = len((d_2340_v11_) + (d_2340_v11_))
        for d_2341_i1_ in range(len(d_2340_v11_), hi13_):
            d_2342_v12_: C3
            nw386_ = C3()
            nw386_.ctor__((d_2339_v10_ if not(d_2324_v0_) else d_2339_v10_), d_2324_v0_, self.f26)
            d_2342_v12_ = nw386_
            d_2342_v12_ = d_2342_v12_
            (globalState).f1 = not((-452) == (34))
            d_2343_v13_: _dafny.Array
            nw387_ = _dafny.Array(None, 12)
            d_2343_v13_ = nw387_
            d_2344_v14_: _dafny.Seq
            d_2344_v14_ = _dafny.SeqWithoutIsStrInference([d_2324_v0_, d_2342_v12_.f40])
            d_2345_v15_: _dafny.MultiSet
            d_2345_v15_ = _dafny.MultiSet([p2])
            d_2346_v16_: _dafny.Map
            d_2346_v16_ = _dafny.Map({(d_2339_v10_)[default__.safeIndex(144, (d_2339_v10_).length(0))]: d_2324_v0_})
            d_2347_v17_: _dafny.Set
            d_2347_v17_ = _dafny.Set({False, ((d_2346_v16_)[387] if (387) in (d_2346_v16_) else d_2342_v12_.f40), d_2324_v0_})
            d_2348_v18_: _dafny.Map
            d_2348_v18_ = _dafny.Map({(d_2345_v15_) | (_dafny.MultiSet([d_2341_i1_])): d_2347_v17_})
            rhs357_ = d_2343_v13_
            rhs358_ = d_2344_v14_
            rhs359_ = (default__.fm58(d_2325_v1_, globalState)) | ((d_2348_v18_ if d_2342_v12_.f40 else d_2348_v18_))
            d_2343_v13_ = rhs357_
            d_2344_v14_ = rhs358_
            d_2348_v18_ = rhs359_
            arr12_ = d_2329_v5_.f33
            index364_ = default__.safeIndex(733, (d_2329_v5_.f33).length(0))
            arr12_[index364_] = d_2342_v12_.f40
            arr13_ = d_2329_v5_.f33
            index365_ = default__.safeIndex(733, (d_2329_v5_.f33).length(0))
            arr13_[index365_] = d_2324_v0_
        r0 = (d_2339_v10_)[default__.safeIndex(144, (d_2339_v10_).length(0))]
        d_2349_v19_: _dafny.Map
        d_2349_v19_ = _dafny.Map({d_2340_v11_: d_2327_v3_})
        d_2350_v20_: _dafny.Map
        d_2350_v20_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(D4_DC8(d_2325_v1_)).cf13 for d_2351_i2_ in range(default__.abs(819))]): d_2349_v19_})
        d_2352_v21_: _dafny.Map
        d_2352_v21_ = ((d_2350_v20_)[d_2340_v11_] if (d_2340_v11_) in (d_2350_v20_) else d_2349_v19_)
        r1 = (d_2352_v21_)
        r2 = (self).fm3(globalState)
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        d_2353_v0_: bool
        d_2353_v0_ = False
        (globalState).f1 = d_2353_v0_
        d_2354_v1_: _dafny.Map
        d_2354_v1_ = _dafny.Map({(684) - ((self).f30): d_2353_v0_})
        d_2355_v2_: _dafny.Seq
        d_2355_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqibxhn"))
        d_2356_i0_: int
        d_2356_i0_ = 0
        with _dafny.label("19"):
            while ((d_2354_v1_)[len(d_2355_v2_)] if (len(d_2355_v2_)) in (d_2354_v1_) else not(d_2353_v0_)):
                with _dafny.c_label("19"):
                    if (d_2356_i0_) >= (100):
                        raise _dafny.Break("19")
                    d_2356_i0_ = (d_2356_i0_) + (1)
                    if d_2353_v0_:
                        d_2357_v3_: _dafny.Array
                        nw388_ = _dafny.Array(int(0), 10)
                        d_2357_v3_ = nw388_
                        d_2358_v4_: str
                        d_2358_v4_ = _dafny.CodePoint('y')
                        d_2357_v3_ = (D0_DC1(d_2357_v3_, d_2358_v4_, (self).f30)).cf1
                        (globalState).f5 = (self).f30
                        (globalState).f13 = default__.fm0((self).f30, globalState)
                        d_2359_v5_: _dafny.Array
                        nw389_ = _dafny.Array(_dafny.Array(None, 0), 3)
                        d_2359_v5_ = nw389_
                        d_2359_v5_ = d_2359_v5_
                        (globalState).f1 = d_2353_v0_
                    elif True:
                        d_2360_v6_: bool
                        d_2361_v7_: bool
                        d_2362_v8_: _dafny.Map
                        d_2363_v9_: int
                        out86_: bool
                        out87_: bool
                        out88_: _dafny.Map
                        out89_: int
                        out86_, out87_, out88_, out89_ = default__.m0(globalState)
                        d_2360_v6_ = out86_
                        d_2361_v7_ = out87_
                        d_2362_v8_ = out88_
                        d_2363_v9_ = out89_
                        d_2364_v10_: str
                        d_2364_v10_ = _dafny.CodePoint('y')
                        d_2365_v11_: D4
                        d_2365_v11_ = D4_DC8(d_2364_v10_)
                        d_2366_v12_: _dafny.Array
                        nw390_ = _dafny.Array(None, 29)
                        nw390_[int(0)] = _dafny.CodePoint('o')
                        nw390_[int(1)] = d_2364_v10_
                        nw390_[int(2)] = (d_2365_v11_).cf13
                        nw390_[int(3)] = default__.fm35((0) - (default__.fm2(globalState)), globalState)
                        nw390_[int(4)] = d_2364_v10_
                        nw390_[int(5)] = d_2364_v10_
                        nw390_[int(6)] = d_2364_v10_
                        nw390_[int(7)] = d_2364_v10_
                        nw390_[int(8)] = d_2364_v10_
                        nw390_[int(9)] = d_2364_v10_
                        nw390_[int(10)] = (d_2364_v10_ if d_2361_v7_ else _dafny.CodePoint('o'))
                        nw390_[int(11)] = d_2364_v10_
                        nw390_[int(12)] = d_2364_v10_
                        nw390_[int(13)] = _dafny.CodePoint('c')
                        nw390_[int(14)] = (d_2364_v10_ if d_2353_v0_ else _dafny.CodePoint('b'))
                        nw390_[int(15)] = d_2364_v10_
                        nw390_[int(16)] = d_2364_v10_
                        nw390_[int(17)] = d_2364_v10_
                        nw390_[int(18)] = d_2364_v10_
                        nw390_[int(19)] = d_2364_v10_
                        nw390_[int(20)] = d_2364_v10_
                        nw390_[int(21)] = d_2364_v10_
                        nw390_[int(22)] = d_2364_v10_
                        nw390_[int(23)] = d_2364_v10_
                        nw390_[int(24)] = d_2364_v10_
                        nw390_[int(25)] = _dafny.CodePoint('f')
                        nw390_[int(26)] = _dafny.CodePoint('b')
                        nw390_[int(27)] = d_2364_v10_
                        nw390_[int(28)] = d_2364_v10_
                        d_2366_v12_ = nw390_
                        index366_ = default__.safeIndex(496, (d_2366_v12_).length(0))
                        (d_2366_v12_)[index366_] = d_2364_v10_
                        index367_ = default__.safeIndex(496, (d_2366_v12_).length(0))
                        (d_2366_v12_)[index367_] = _dafny.CodePoint('j')
                        d_2367_v13_: D4
                        d_2367_v13_ = D4_DC10((self).f30)
                        d_2368_v14_: _dafny.MultiSet
                        d_2368_v14_ = _dafny.MultiSet([d_2355_v2_, (d_2355_v2_).set(default__.safeIndex((self).f30, len(d_2355_v2_)), (d_2366_v12_)[default__.safeIndex(496, (d_2366_v12_).length(0))])])
                        d_2369_v15_: _dafny.MultiSet
                        d_2369_v15_ = d_2368_v14_
                        d_2370_v16_: _dafny.MultiSet
                        d_2370_v16_ = _dafny.MultiSet([d_2369_v15_])
                        d_2367_v13_ = default__.fm59((d_2370_v16_).issubset(d_2370_v16_), (self).f30, globalState)
                        d_2371_v17_: _dafny.Map
                        d_2371_v17_ = _dafny.Map({d_2360_v6_: len(d_2355_v2_)})
                        d_2372_v18_: C1
                        nw391_ = C1()
                        nw391_.ctor__(d_2371_v17_)
                        d_2372_v18_ = nw391_
                        d_2373_v19_: _dafny.Array
                        def lambda108_(d_2374_v6_):
                            def lambda109_(d_2375_i1_):
                                return d_2374_v6_

                            return lambda109_

                        init64_ = lambda108_(d_2360_v6_)
                        nw392_ = _dafny.Array(None, 19)
                        for i0_64_ in range(nw392_.length(0)):
                            nw392_[i0_64_] = init64_(i0_64_)
                        d_2373_v19_ = nw392_
                        d_2376_v22_: C9
                        nw393_ = C9()
                        def iife257_():
                            coll93_ = _dafny.Map()
                            compr_93_: int
                            for compr_93_ in _dafny.IntegerRange(808, 459):
                                d_2377_v20_: int = compr_93_
                                if ((808) <= (d_2377_v20_)) and ((d_2377_v20_) < (459)):
                                    coll93_[default__.safeModuloInt(d_2377_v20_, 145)] = (self).f30
                            return _dafny.Map(coll93_)
                        def iife258_():
                            coll94_ = _dafny.Map()
                            compr_94_: int
                            for compr_94_ in _dafny.IntegerRange(902, -754):
                                d_2378_v21_: int = compr_94_
                                if ((902) <= (d_2378_v21_)) and ((d_2378_v21_) < (-754)):
                                    coll94_[(d_2378_v21_) + (d_2363_v9_)] = (self).f30
                            return _dafny.Map(coll94_)
                        nw393_.ctor__(d_2373_v19_, (default__.fm60((self).f30, (self).f30, globalState)).cf26, (iife257_()
                        ) | (iife258_()
                        ))
                        d_2376_v22_ = nw393_
                        d_2379_v23_: _dafny.Seq
                        d_2379_v23_ = _dafny.SeqWithoutIsStrInference([d_2376_v22_])
                        d_2376_v22_ = (d_2379_v23_)[default__.safeIndex((self).f30, len(d_2379_v23_))]
                    d_2380_v24_: D8
                    d_2380_v24_ = D8_DC20(d_2353_v0_)
                    d_2381_v25_: _dafny.Array
                    nw394_ = _dafny.Array(None, 22)
                    nw394_[int(0)] = d_2380_v24_
                    nw394_[int(1)] = d_2380_v24_
                    nw394_[int(2)] = d_2380_v24_
                    nw394_[int(3)] = d_2380_v24_
                    nw394_[int(4)] = d_2380_v24_
                    nw394_[int(5)] = d_2380_v24_
                    nw394_[int(6)] = d_2380_v24_
                    nw394_[int(7)] = d_2380_v24_
                    nw394_[int(8)] = d_2380_v24_
                    def iife259_(_pat_let82_0):
                        def iife260_(d_2382_dt__update__tmp_h0_):
                            def iife261_(_pat_let83_0):
                                def iife262_(d_2383_dt__update_hcf31_h0_):
                                    return D8_DC20(d_2383_dt__update_hcf31_h0_)
                                return iife262_(_pat_let83_0)
                            return iife261_(False)
                        return iife260_(_pat_let82_0)
                    nw394_[int(9)] = iife259_(d_2380_v24_)
                    nw394_[int(10)] = d_2380_v24_
                    nw394_[int(11)] = d_2380_v24_
                    nw394_[int(12)] = D8_DC20(d_2353_v0_)
                    nw394_[int(13)] = d_2380_v24_
                    nw394_[int(14)] = d_2380_v24_
                    nw394_[int(15)] = D8_DC20(d_2353_v0_)
                    nw394_[int(16)] = d_2380_v24_
                    nw394_[int(17)] = d_2380_v24_
                    nw394_[int(18)] = D8_DC20(d_2353_v0_)
                    nw394_[int(19)] = d_2380_v24_
                    nw394_[int(20)] = D8_DC20(d_2353_v0_)
                    nw394_[int(21)] = d_2380_v24_
                    d_2381_v25_ = nw394_
                    index368_ = default__.safeIndex(315, (d_2381_v25_).length(0))
                    (d_2381_v25_)[index368_] = d_2380_v24_
                    d_2384_v27_: _dafny.Seq
                    d_2384_v27_ = _dafny.SeqWithoutIsStrInference([len(default__.fm30(len(d_2355_v2_), d_2355_v2_, (self).f30, globalState))])
                    d_2385_v28_: _dafny.Map
                    d_2385_v28_ = _dafny.Map({(self).f30: d_2384_v27_})
                    d_2386_v29_: _dafny.Seq
                    d_2386_v29_ = _dafny.SeqWithoutIsStrInference([d_2384_v27_, ((d_2385_v28_)[len(self.f26)] if (len(self.f26)) in (d_2385_v28_) else d_2384_v27_)])
                    d_2387_v30_: _dafny.MultiSet
                    d_2387_v30_ = _dafny.MultiSet([(self).f30, (self).f30])
                    d_2388_v31_: D11
                    d_2388_v31_ = D11_DC27(default__.fm2(globalState), (self).f30, d_2353_v0_)
                    index369_ = default__.safeIndex(315, (d_2381_v25_).length(0))
                    def iife263_():
                        coll95_ = _dafny.Set()
                        compr_95_: int
                        for compr_95_ in _dafny.IntegerRange(610, 622):
                            d_2389_v26_: int = compr_95_
                            if ((610) <= (d_2389_v26_)) and ((d_2389_v26_) < (622)):
                                coll95_ = coll95_.union(_dafny.Set([default__.safeModuloInt(d_2389_v26_, 780)]))
                        return _dafny.Set(coll95_)
                    rhs360_ = d_2380_v24_
                    rhs361_ = (iife263_()
                    ).issubset(_dafny.Set({len(d_2386_v29_), (self).f30, (self).f30, (self).f30, ((d_2387_v30_)[(self).f30] if ((self).f30) in (d_2387_v30_) else (self).f30)}))
                    rhs362_ = (0) - ((0) - (default__.safeDivisionInt((self).f30, (d_2388_v31_).cf41)))
                    lhs278_ = d_2381_v25_
                    lhs279_ = default__.safeIndex(315, (d_2381_v25_).length(0))
                    lhs280_ = globalState
                    lhs281_ = globalState
                    lhs278_[lhs279_] = rhs360_
                    lhs280_.f1 = rhs361_
                    lhs281_.f5 = rhs362_
                    (globalState).f21 = default__.safeDivisionInt((self).f30, default__.safeDivisionInt(526, (self).f30))
                    d_2390_v32_: _dafny.Map
                    d_2390_v32_ = _dafny.Map({d_2353_v0_: (self).f30})
                    d_2391_v33_: C1
                    nw395_ = C1()
                    nw395_.ctor__(d_2390_v32_)
                    d_2391_v33_ = nw395_
                    pass
            pass
        if not(not (d_2353_v0_) or (not(d_2353_v0_))):
            d_2353_v0_ = (not(d_2353_v0_)) or (d_2353_v0_)
            d_2392_v34_: _dafny.Set
            d_2392_v34_ = _dafny.Set({d_2353_v0_})
            d_2393_v35_: _dafny.Array
            nw396_ = _dafny.Array(None, 15)
            nw396_[int(0)] = (self).f30
            nw396_[int(1)] = (self).f30
            nw396_[int(2)] = (self).f30
            nw396_[int(3)] = (self).f30
            nw396_[int(4)] = (self).fm3(globalState)
            nw396_[int(5)] = ((self.f26)[(0) - ((self).f30)] if ((0) - ((self).f30)) in (self.f26) else (self).f30)
            nw396_[int(6)] = (self).f30
            nw396_[int(7)] = len(d_2355_v2_)
            nw396_[int(8)] = (self).f30
            nw396_[int(9)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2394_i2_ in range(default__.abs(226))])) + (d_2355_v2_))
            nw396_[int(10)] = ((self).f30) - ((self).f30)
            nw396_[int(11)] = (0) - (default__.safeDivisionInt(len(d_2392_v34_), (0) - ((self).f30)))
            nw396_[int(12)] = ((self).f30) + ((self).f30)
            nw396_[int(13)] = ((self).f30) * ((self).f30)
            nw396_[int(14)] = (self).f30
            d_2393_v35_ = nw396_
            index370_ = default__.safeIndex(854, (d_2393_v35_).length(0))
            (d_2393_v35_)[index370_] = (self).f30
            d_2395_v36_: _dafny.MultiSet
            d_2395_v36_ = _dafny.MultiSet([d_2353_v0_, d_2353_v0_, d_2353_v0_])
            d_2396_v37_: _dafny.Array
            nw397_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_2396_v37_ = nw397_
            index371_ = default__.safeIndex(854, (d_2393_v35_).length(0))
            rhs363_ = (self).f30
            rhs364_ = d_2395_v36_
            rhs365_ = d_2396_v37_
            lhs282_ = d_2393_v35_
            lhs283_ = default__.safeIndex(854, (d_2393_v35_).length(0))
            lhs282_[lhs283_] = rhs363_
            d_2395_v36_ = rhs364_
            d_2396_v37_ = rhs365_
            d_2397_v38_: _dafny.Seq
            d_2397_v38_ = _dafny.SeqWithoutIsStrInference([d_2395_v36_])
            d_2398_v39_: _dafny.Set
            d_2398_v39_ = _dafny.Set({len(d_2397_v38_)})
            d_2398_v39_ = ((d_2398_v39_) | (d_2398_v39_)) | (d_2398_v39_)
            (globalState).f1 = default__.fm0(default__.safeModuloInt((d_2393_v35_)[default__.safeIndex(854, (d_2393_v35_).length(0))], (self).f30), globalState)
            d_2399_v40_: str
            d_2399_v40_ = _dafny.CodePoint('u')
            d_2400_v41_: _dafny.MultiSet
            d_2400_v41_ = _dafny.MultiSet([d_2399_v40_, _dafny.CodePoint('k')])
            d_2401_v42_: _dafny.Array
            nw398_ = _dafny.Array(None, 15)
            nw398_[int(0)] = d_2400_v41_
            nw398_[int(1)] = d_2400_v41_
            nw398_[int(2)] = _dafny.MultiSet(d_2355_v2_)
            nw398_[int(3)] = _dafny.MultiSet([d_2399_v40_, _dafny.CodePoint('q'), d_2399_v40_, default__.fm35((0) - (default__.fm2(globalState)), globalState), d_2399_v40_])
            nw398_[int(4)] = d_2400_v41_
            nw398_[int(5)] = d_2400_v41_
            nw398_[int(6)] = d_2400_v41_
            nw398_[int(7)] = _dafny.MultiSet([d_2399_v40_, d_2399_v40_, d_2399_v40_, _dafny.CodePoint('g')])
            nw398_[int(8)] = d_2400_v41_
            nw398_[int(9)] = d_2400_v41_
            nw398_[int(10)] = d_2400_v41_
            nw398_[int(11)] = d_2400_v41_
            nw398_[int(12)] = d_2400_v41_
            nw398_[int(13)] = d_2400_v41_
            nw398_[int(14)] = (d_2400_v41_).set(d_2399_v40_, default__.abs((d_2393_v35_)[default__.safeIndex(854, (d_2393_v35_).length(0))]))
            d_2401_v42_ = nw398_
            d_2402_v43_: D2
            d_2402_v43_ = D2_DC4(d_2401_v42_)
            d_2403_v44_: _dafny.Seq
            d_2403_v44_ = _dafny.SeqWithoutIsStrInference([d_2353_v0_, d_2353_v0_, d_2353_v0_, not(d_2353_v0_), (d_2395_v36_).issubset(d_2395_v36_)])
            rhs366_ = d_2392_v34_
            rhs367_ = d_2402_v43_
            rhs368_ = ((d_2355_v2_) + (d_2355_v2_)) + (default__.fm30(len(d_2355_v2_), d_2355_v2_, 529, globalState))
            rhs369_ = (d_2403_v44_) + (_dafny.SeqWithoutIsStrInference([d_2353_v0_]))
            rhs370_ = (self).f30
            lhs284_ = globalState
            d_2392_v34_ = rhs366_
            d_2402_v43_ = rhs367_
            d_2355_v2_ = rhs368_
            d_2403_v44_ = rhs369_
            lhs284_.f16 = rhs370_
        elif True:
            d_2404_v45_: _dafny.Seq
            d_2404_v45_ = _dafny.SeqWithoutIsStrInference([d_2353_v0_])
            if (_dafny.MultiSet([((d_2354_v1_)[(self).f30] if ((self).f30) in (d_2354_v1_) else d_2353_v0_), d_2353_v0_, not(d_2353_v0_), d_2353_v0_, d_2353_v0_])).isdisjoint(_dafny.MultiSet(d_2404_v45_)):
                d_2405_v46_: _dafny.Array
                nw399_ = _dafny.Array(False, 5)
                d_2405_v46_ = nw399_
                index372_ = default__.safeIndex(507, (d_2405_v46_).length(0))
                (d_2405_v46_)[index372_] = d_2353_v0_
                index373_ = default__.safeIndex(507, (d_2405_v46_).length(0))
                (d_2405_v46_)[index373_] = d_2353_v0_
                (globalState).f13 = (_dafny.SeqWithoutIsStrInference([default__.fm0((self).f30, globalState)])) < (d_2404_v45_)
                d_2406_v47_: _dafny.Array
                def lambda110_(d_2407_v45_):
                    def lambda111_(d_2408_i3_):
                        return default__.safeDivisionInt(d_2408_i3_, len(d_2407_v45_))

                    return lambda111_

                init65_ = lambda110_(d_2404_v45_)
                nw400_ = _dafny.Array(None, 22)
                for i0_65_ in range(nw400_.length(0)):
                    nw400_[i0_65_] = init65_(i0_65_)
                d_2406_v47_ = nw400_
                index374_ = default__.safeIndex(685, (d_2406_v47_).length(0))
                (d_2406_v47_)[index374_] = (self).f30
                index375_ = default__.safeIndex(685, (d_2406_v47_).length(0))
                (d_2406_v47_)[index375_] = (self).f30
                index376_ = default__.safeIndex(685, (d_2406_v47_).length(0))
                (d_2406_v47_)[index376_] = 513
                r2 = not((d_2355_v2_) <= (d_2355_v2_))
            elif True:
                d_2409_v48_: _dafny.MultiSet
                d_2409_v48_ = _dafny.MultiSet([d_2353_v0_])
                d_2410_v49_: _dafny.Seq
                d_2410_v49_ = _dafny.SeqWithoutIsStrInference([d_2409_v48_, d_2409_v48_])
                d_2411_v50_: _dafny.Map
                d_2411_v50_ = _dafny.Map({(d_2410_v49_)[default__.safeIndex((d_2409_v48_).cardinality, len(d_2410_v49_))]: self.f26})
                d_2411_v50_ = d_2411_v50_
                d_2412_v51_: str
                d_2412_v51_ = _dafny.CodePoint('h')
                d_2413_v52_: _dafny.Array
                nw401_ = _dafny.Array(None, 20)
                nw401_[int(0)] = not (d_2353_v0_) or (d_2353_v0_)
                nw401_[int(1)] = d_2353_v0_
                nw401_[int(2)] = not (not(d_2353_v0_)) or (d_2353_v0_)
                nw401_[int(3)] = d_2353_v0_
                nw401_[int(4)] = d_2353_v0_
                nw401_[int(5)] = (d_2404_v45_)[default__.safeIndex(len(self.f26), len(d_2404_v45_))]
                nw401_[int(6)] = d_2353_v0_
                nw401_[int(7)] = d_2353_v0_
                nw401_[int(8)] = d_2353_v0_
                nw401_[int(9)] = d_2353_v0_
                nw401_[int(10)] = (d_2355_v2_) < (d_2355_v2_)
                nw401_[int(11)] = d_2353_v0_
                nw401_[int(12)] = d_2353_v0_
                nw401_[int(13)] = d_2353_v0_
                nw401_[int(14)] = d_2353_v0_
                nw401_[int(15)] = d_2353_v0_
                nw401_[int(16)] = (d_2353_v0_) == (not(d_2353_v0_))
                nw401_[int(17)] = True
                nw401_[int(18)] = d_2353_v0_
                nw401_[int(19)] = d_2353_v0_
                d_2413_v52_ = nw401_
                d_2414_v53_: D4
                d_2414_v53_ = D4_DC8(d_2412_v51_)
                rhs371_ = not(d_2353_v0_)
                rhs372_ = (d_2414_v53_).cf13
                rhs373_ = (self).f30
                rhs374_ = d_2413_v52_
                lhs285_ = globalState
                lhs286_ = globalState
                lhs285_.f1 = rhs371_
                d_2412_v51_ = rhs372_
                lhs286_.f22 = rhs373_
                d_2413_v52_ = rhs374_
                index377_ = default__.safeIndex(138, (d_2413_v52_).length(0))
                (d_2413_v52_)[index377_] = d_2353_v0_
                index378_ = default__.safeIndex(138, (d_2413_v52_).length(0))
                (d_2413_v52_)[index378_] = (default__.fm2(globalState)) >= (-84)
                d_2415_v54_: _dafny.Set
                d_2415_v54_ = _dafny.Set({d_2355_v2_, d_2355_v2_})
                d_2416_v55_: _dafny.Map
                d_2416_v55_ = _dafny.Map({d_2353_v0_: default__.fm0((self).f30, globalState)})
                d_2417_v56_: _dafny.Map
                d_2417_v56_ = _dafny.Map({(d_2413_v52_)[default__.safeIndex(138, (d_2413_v52_).length(0))]: (self).f30})
                d_2415_v54_ = default__.fm47(-475, d_2416_v55_, ((self).f30) - (((d_2417_v56_)[(d_2404_v45_)[default__.safeIndex((self).f30, len(d_2404_v45_))]] if ((d_2404_v45_)[default__.safeIndex((self).f30, len(d_2404_v45_))]) in (d_2417_v56_) else (self).f30)), globalState)
                d_2418_v57_: _dafny.Array
                def lambda112_(d_2419_v45_):
                    def lambda113_(d_2420_i4_):
                        return d_2419_v45_

                    return lambda113_

                init66_ = lambda112_(d_2404_v45_)
                nw402_ = _dafny.Array(None, 14)
                for i0_66_ in range(nw402_.length(0)):
                    nw402_[i0_66_] = init66_(i0_66_)
                d_2418_v57_ = nw402_
                index379_ = default__.safeIndex(216, (d_2418_v57_).length(0))
                (d_2418_v57_)[index379_] = d_2404_v45_
                index380_ = default__.safeIndex(216, (d_2418_v57_).length(0))
                (d_2418_v57_)[index380_] = d_2404_v45_
            (globalState).f0 = default__.safeModuloInt((self).f30, (0) - ((self).f30))
            d_2421_v58_: _dafny.Array
            def lambda114_(d_2422_i5_):
                return (d_2422_i5_) * ((self).f30)

            init67_ = lambda114_
            nw403_ = _dafny.Array(None, 7)
            for i0_67_ in range(nw403_.length(0)):
                nw403_[i0_67_] = init67_(i0_67_)
            d_2421_v58_ = nw403_
            d_2423_v59_: _dafny.Map
            d_2423_v59_ = _dafny.Map({d_2353_v0_: d_2421_v58_})
            d_2424_v60_: D19
            d_2424_v60_ = D19_DC37(d_2423_v59_)
            d_2424_v60_ = d_2424_v60_
            d_2425_v61_: D5
            d_2425_v61_ = D5_DC11(d_2353_v0_)
            d_2426_v62_: C8
            nw404_ = C8()
            nw404_.ctor__(d_2353_v0_, d_2425_v61_, self.f26)
            d_2426_v62_ = nw404_
            index381_ = default__.safeIndex(348, (d_2421_v58_).length(0))
            (d_2421_v58_)[index381_] = ((self).f30 if d_2353_v0_ else (0) - ((self).f30))
            d_2427_v63_: str
            d_2427_v63_ = _dafny.CodePoint('v')
            d_2428_v64_: _dafny.MultiSet
            d_2428_v64_ = _dafny.MultiSet([d_2353_v0_])
            index382_ = default__.safeIndex(348, (d_2421_v58_).length(0))
            rhs375_ = False
            rhs376_ = default__.fm30((self).f30, (d_2355_v2_).set(default__.safeIndex((self).f30, len(d_2355_v2_)), d_2427_v63_), (((d_2428_v64_).set(d_2353_v0_, default__.abs((self).f30))).intersection(d_2428_v64_)).cardinality, globalState)
            rhs377_ = (self).f30
            lhs287_ = globalState
            lhs288_ = d_2421_v58_
            lhs289_ = default__.safeIndex(348, (d_2421_v58_).length(0))
            lhs287_.f13 = rhs375_
            d_2355_v2_ = rhs376_
            lhs288_[lhs289_] = rhs377_
        d_2429_v65_: _dafny.Seq
        d_2429_v65_ = _dafny.SeqWithoutIsStrInference([(self).f30])
        rhs378_ = (self).f30
        rhs379_ = d_2429_v65_
        rhs380_ = d_2353_v0_
        lhs290_ = globalState
        lhs291_ = globalState
        lhs290_.f22 = rhs378_
        d_2429_v65_ = rhs379_
        lhs291_.f13 = rhs380_
        d_2430_v66_: D5
        d_2430_v66_ = D5_DC11(d_2353_v0_)
        d_2431_v67_: T1
        nw405_ = C8()
        nw405_.ctor__(True, d_2430_v66_, self.f26)
        d_2431_v67_ = nw405_
        d_2432_v68_: _dafny.MultiSet
        d_2432_v68_ = _dafny.MultiSet([(0) - (default__.safeModuloInt(len(_dafny.Set({d_2431_v67_, d_2431_v67_, d_2431_v67_, d_2431_v67_})), (0) - ((self).f30))), 106, (self).f30, (self).f30])
        d_2433_v69_: _dafny.Map
        d_2433_v69_ = _dafny.Map({d_2353_v0_: (self).f30})
        d_2434_v70_: _dafny.Array
        def lambda115_(d_2435_v2_):
            def lambda116_(d_2436_i7_):
                return (d_2436_i7_) - ((D5_DC13(len(d_2435_v2_))).cf19)

            return lambda116_

        init68_ = lambda115_(d_2355_v2_)
        nw406_ = _dafny.Array(None, 7)
        for i0_68_ in range(nw406_.length(0)):
            nw406_[i0_68_] = init68_(i0_68_)
        d_2434_v70_ = nw406_
        d_2437_v71_: _dafny.Map
        d_2437_v71_ = _dafny.Map({True: d_2434_v70_})
        d_2438_v72_: _dafny.Map
        d_2438_v72_ = _dafny.Map({((d_2429_v65_) + (_dafny.SeqWithoutIsStrInference([(0) - ((self).f30) for d_2439_i6_ in range(default__.abs(375))]))).set(default__.safeIndex(((d_2433_v69_)[d_2353_v0_] if (d_2353_v0_) in (d_2433_v69_) else (self).f30), len((d_2429_v65_) + (_dafny.SeqWithoutIsStrInference([(0) - ((self).f30) for d_2440_i6_ in range(default__.abs(375))])))), len(d_2437_v71_)): d_2432_v68_})
        d_2432_v68_ = ((d_2438_v72_)[(d_2429_v65_) + (d_2429_v65_)] if ((d_2429_v65_) + (d_2429_v65_)) in (d_2438_v72_) else _dafny.MultiSet(d_2429_v65_))
        if d_2353_v0_:
            d_2441_v73_: str
            d_2441_v73_ = _dafny.CodePoint('i')
            d_2355_v2_ = ((d_2355_v2_) + ((d_2355_v2_) + (d_2355_v2_))).set(default__.safeIndex((self).f30, len((d_2355_v2_) + ((d_2355_v2_) + (d_2355_v2_)))), d_2441_v73_)
            d_2442_v74_: _dafny.Set
            d_2442_v74_ = _dafny.Set({(d_2433_v69_) | (d_2433_v69_)})
            rhs381_ = (_dafny.Set({d_2433_v69_, d_2433_v69_})) | (d_2442_v74_)
            rhs382_ = (self).f30
            lhs292_ = globalState
            d_2442_v74_ = rhs381_
            lhs292_.f22 = rhs382_
            d_2429_v65_ = d_2429_v65_
            (globalState).f2 = ((self).f30) * (((self).f30) * ((self).f30))
            (globalState).f16 = default__.safeModuloInt((self).f30, (self).f30)
        elif True:
            r2 = d_2353_v0_
            d_2355_v2_ = default__.fm30(((self).f30) + ((self).f30), (d_2355_v2_) + (d_2355_v2_), (self).f30, globalState)
            (globalState).f16 = ((self).f30) * (459)
            d_2443_v75_: _dafny.Array
            nw407_ = _dafny.Array(False, 22)
            d_2443_v75_ = nw407_
            d_2444_v76_: _dafny.Map
            d_2444_v76_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2443_v75_, d_2443_v75_]): self.f26})
            d_2445_v77_: _dafny.Seq
            d_2445_v77_ = _dafny.SeqWithoutIsStrInference([d_2443_v75_, d_2443_v75_])
            d_2444_v76_ = (d_2444_v76_).set(d_2445_v77_, self.f26)
            d_2446_v78_: _dafny.Array
            nw408_ = _dafny.Array(_dafny.CodePoint('D'), 25)
            d_2446_v78_ = nw408_
            d_2447_v79_: D0
            d_2447_v79_ = D0_DC0(d_2446_v78_)
            d_2448_v80_: D11
            d_2448_v80_ = D11_DC26(d_2443_v75_)
            rhs383_ = d_2447_v79_
            rhs384_ = d_2353_v0_
            rhs385_ = (d_2448_v80_).cf39
            rhs386_ = (not((d_2353_v0_) and (d_2353_v0_))) or ((_dafny.CodePoint('a')) in (d_2355_v2_))
            lhs293_ = globalState
            lhs294_ = globalState
            d_2447_v79_ = rhs383_
            lhs293_.f13 = rhs384_
            d_2443_v75_ = rhs385_
            lhs294_.f1 = rhs386_
        r0 = d_2353_v0_
        d_2449_v81_: _dafny.Map
        d_2449_v81_ = _dafny.Map({d_2353_v0_: d_2353_v0_})
        d_2450_v82_: _dafny.Map
        d_2450_v82_ = _dafny.Map({d_2449_v81_: d_2353_v0_})
        d_2451_v83_: _dafny.Seq
        d_2451_v83_ = _dafny.SeqWithoutIsStrInference([((d_2450_v82_)[d_2449_v81_] if (d_2449_v81_) in (d_2450_v82_) else True)])
        d_2452_v84_: _dafny.Set
        d_2452_v84_ = _dafny.Set({((self.f26).set((self).f30, len(d_2451_v83_))).set((self).f30, (self).f30), self.f26})
        r1 = d_2452_v84_
        r2 = d_2353_v0_
        return r0, r1, r2

    def m7(self, p0, globalState):
        r0: bool = False
        d_2453_v0_: bool
        d_2453_v0_ = False
        d_2454_v1_: _dafny.Map
        d_2454_v1_ = _dafny.Map({d_2453_v0_: False})
        d_2454_v1_ = (d_2454_v1_).set(d_2453_v0_, d_2453_v0_)
        def iife264_():
            coll96_ = _dafny.Map()
            compr_96_: int
            for compr_96_ in (self.f26).keys.Elements:
                d_2455_v2_: int = compr_96_
                if (d_2455_v2_) in (self.f26):
                    def iife265_():
                        coll97_ = _dafny.Set()
                        compr_97_: int
                        for compr_97_ in _dafny.IntegerRange(448, 708):
                            d_2456_v3_: int = compr_97_
                            if ((448) <= (d_2456_v3_)) and ((d_2456_v3_) < (708)):
                                coll97_ = coll97_.union(_dafny.Set([(d_2456_v3_) * (695)]))
                        return _dafny.Set(coll97_)
                    coll96_[(d_2455_v2_) - (((self.f26)[len(_dafny.Set({p0, (self).f30}))] if (len(_dafny.Set({p0, (self).f30}))) in (self.f26) else p0))] = (0) - (len(iife265_()
                    ))
            return _dafny.Map(coll96_)
        (globalState).f12 = (self.f26) | (iife264_()
        )
        if d_2453_v0_:
            d_2457_v4_: _dafny.Array
            nw409_ = _dafny.Array(int(0), 11)
            d_2457_v4_ = nw409_
            d_2458_v5_: _dafny.Seq
            d_2458_v5_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2459_v6_: _dafny.Seq
            d_2459_v6_ = _dafny.SeqWithoutIsStrInference([p0, len(d_2458_v5_)])
            d_2460_v7_: _dafny.Map
            d_2460_v7_ = _dafny.Map({d_2457_v4_: len(d_2459_v6_)})
            d_2461_v8_: _dafny.Map
            d_2461_v8_ = ((d_2460_v7_).set(d_2457_v4_, (self).fm3(globalState))) | (d_2460_v7_)
            source33_ = d_2461_v8_
            d_2462___mcc_h0_ = source33_
            d_2463_cf5_ = d_2462___mcc_h0_
            d_2464_v9_: str
            d_2464_v9_ = _dafny.CodePoint('g')
            d_2465_v10_: _dafny.Seq
            d_2465_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usrimhp"))
            d_2466_v11_: _dafny.Seq
            d_2466_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfirpf")), d_2465_v10_, (default__.fm30(len(self.f26), d_2465_v10_, -180, globalState)).set(default__.safeIndex((self).f30, len(default__.fm30(len(self.f26), d_2465_v10_, -180, globalState))), d_2464_v9_)])
            rhs387_ = d_2464_v9_
            rhs388_ = d_2453_v0_
            rhs389_ = d_2466_v11_
            d_2464_v9_ = rhs387_
            d_2453_v0_ = rhs388_
            d_2466_v11_ = rhs389_
            d_2467_v12_: bool
            d_2468_v13_: bool
            d_2469_v14_: _dafny.Map
            d_2470_v15_: int
            out90_: bool
            out91_: bool
            out92_: _dafny.Map
            out93_: int
            out90_, out91_, out92_, out93_ = default__.m0(globalState)
            d_2467_v12_ = out90_
            d_2468_v13_ = out91_
            d_2469_v14_ = out92_
            d_2470_v15_ = out93_
            (globalState).f1 = d_2453_v0_
            d_2471_v16_: _dafny.Seq
            d_2471_v16_ = _dafny.SeqWithoutIsStrInference([(d_2453_v0_) and (d_2468_v13_), d_2467_v12_, d_2467_v12_, d_2468_v13_])
            d_2471_v16_ = _dafny.SeqWithoutIsStrInference([d_2453_v0_, d_2467_v12_])
            d_2472_v17_: _dafny.Map
            d_2472_v17_ = _dafny.Map({d_2453_v0_: p0})
            d_2473_v18_: _dafny.Map
            d_2473_v18_ = (_dafny.Map({True: -874})) | (d_2472_v17_)
            source34_ = d_2473_v18_
            d_2474___mcc_h1_ = source34_
            d_2475_cf46_ = d_2474___mcc_h1_
            d_2476_v19_: _dafny.Array
            nw410_ = _dafny.Array(False, 26)
            d_2476_v19_ = nw410_
            index383_ = default__.safeIndex(127, (d_2476_v19_).length(0))
            (d_2476_v19_)[index383_] = d_2453_v0_
            index384_ = default__.safeIndex(127, (d_2476_v19_).length(0))
            (d_2476_v19_)[index384_] = d_2453_v0_
            (globalState).f1 = (d_2476_v19_)[default__.safeIndex(127, (d_2476_v19_).length(0))]
            (globalState).f16 = (p0) - ((self).f30)
            d_2477_v20_: _dafny.MultiSet
            d_2477_v20_ = _dafny.MultiSet([d_2453_v0_])
            d_2453_v0_ = (False) == ((d_2477_v20_).isdisjoint(d_2477_v20_))
            d_2478_v21_: _dafny.Map
            d_2478_v21_ = _dafny.Map({p0: False})
            d_2479_v22_: _dafny.MultiSet
            d_2479_v22_ = _dafny.MultiSet([d_2453_v0_, not(not(not(not(d_2453_v0_))))])
            d_2480_v23_: _dafny.Seq
            d_2480_v23_ = _dafny.SeqWithoutIsStrInference([d_2479_v22_])
            d_2481_v24_: D8
            d_2481_v24_ = D8_DC19(d_2453_v0_, d_2453_v0_, d_2453_v0_)
            d_2454_v1_ = (d_2454_v1_).set(d_2453_v0_, ((d_2478_v21_)[((d_2480_v23_)[default__.safeIndex(p0, len(d_2480_v23_))]).cardinality] if (((d_2480_v23_)[default__.safeIndex(p0, len(d_2480_v23_))]).cardinality) in (d_2478_v21_) else (d_2481_v24_).cf30))
            (globalState).f5 = p0
            d_2482_v25_: _dafny.Seq
            d_2482_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpfm"))
            d_2482_v25_ = (d_2482_v25_) + ((d_2482_v25_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csrqqsvim"))))
        elif True:
            d_2483_v26_: str
            d_2483_v26_ = _dafny.CodePoint('p')
            d_2484_v27_: _dafny.Array
            def lambda117_(d_2485_p0_, d_2486_v0_):
                def lambda118_(d_2487_i0_):
                    return (_dafny.Map({d_2485_p0_: d_2486_v0_})) | (_dafny.Map({(self).f30: d_2486_v0_}))

                return lambda118_

            init69_ = lambda117_(p0, d_2453_v0_)
            nw411_ = _dafny.Array(None, 8)
            for i0_69_ in range(nw411_.length(0)):
                nw411_[i0_69_] = init69_(i0_69_)
            d_2484_v27_ = nw411_
            d_2488_v28_: _dafny.Map
            d_2488_v28_ = _dafny.Map({749: d_2453_v0_})
            index385_ = default__.safeIndex(807, (d_2484_v27_).length(0))
            (d_2484_v27_)[index385_] = (d_2488_v28_) | (d_2488_v28_)
            d_2489_v29_: _dafny.Seq
            d_2489_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vu"))
            index386_ = default__.safeIndex(807, (d_2484_v27_).length(0))
            def iife266_():
                coll98_ = _dafny.Map()
                compr_98_: int
                for compr_98_ in _dafny.IntegerRange(285, 611):
                    d_2490_v30_: int = compr_98_
                    if ((285) <= (d_2490_v30_)) and ((d_2490_v30_) < (611)):
                        coll98_[(d_2490_v30_) - ((self).f30)] = d_2453_v0_
                return _dafny.Map(coll98_)
            rhs390_ = (len((_dafny.Set({d_2453_v0_})).intersection(_dafny.Set({d_2453_v0_})))) > (p0)
            rhs391_ = default__.fm0((self).fm9((self).f30, (self).f30, globalState), globalState)
            rhs392_ = (d_2489_v29_)[default__.safeIndex(p0, len(d_2489_v29_))]
            rhs393_ = iife266_()

            lhs295_ = globalState
            lhs296_ = d_2484_v27_
            lhs297_ = default__.safeIndex(807, (d_2484_v27_).length(0))
            lhs295_.f13 = rhs390_
            d_2453_v0_ = rhs391_
            d_2483_v26_ = rhs392_
            lhs296_[lhs297_] = rhs393_
            d_2491_v31_: _dafny.Seq
            d_2491_v31_ = _dafny.SeqWithoutIsStrInference([d_2453_v0_, ((d_2454_v1_)[not(d_2453_v0_)] if (not(d_2453_v0_)) in (d_2454_v1_) else d_2453_v0_), d_2453_v0_, d_2453_v0_])
            d_2489_v29_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpjifo"))).set(default__.safeIndex((670) + (len(d_2491_v31_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpjifo")))), d_2483_v26_)
            d_2492_v32_: _dafny.Array
            nw412_ = _dafny.Array(False, 10)
            d_2492_v32_ = nw412_
            index387_ = default__.safeIndex(794, (d_2492_v32_).length(0))
            (d_2492_v32_)[index387_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhktow"))).set(default__.safeIndex(331, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhktow")))), _dafny.CodePoint('p'))) == (d_2489_v29_)
            index388_ = default__.safeIndex(794, (d_2492_v32_).length(0))
            (d_2492_v32_)[index388_] = d_2453_v0_
            (globalState).f0 = (((self).f30) * (p0)) + (p0)
            d_2493_v34_: _dafny.Map
            d_2493_v34_ = _dafny.Map({(d_2492_v32_)[default__.safeIndex(794, (d_2492_v32_).length(0))]: _dafny.Set({(d_2492_v32_)[default__.safeIndex(794, (d_2492_v32_).length(0))]})})
            d_2494_v35_: _dafny.Map
            d_2494_v35_ = _dafny.Map({((d_2493_v34_)[(d_2492_v32_)[default__.safeIndex(794, (d_2492_v32_).length(0))]] if ((d_2492_v32_)[default__.safeIndex(794, (d_2492_v32_).length(0))]) in (d_2493_v34_) else _dafny.Set({True, d_2453_v0_})): p0})
            d_2495_v36_: _dafny.Set
            def iife267_():
                coll99_ = _dafny.Map()
                compr_99_: _dafny.Set
                for compr_99_ in (d_2494_v35_).keys.Elements:
                    d_2496_v33_: _dafny.Set = compr_99_
                    if (d_2496_v33_) in (d_2494_v35_):
                        coll99_[d_2496_v33_] = p0
                return _dafny.Map(coll99_)
            d_2495_v36_ = _dafny.Set({self.f26, _dafny.Map({len(_dafny.Set({iife267_()
            })): 168})})
            d_2497_v38_: D11
            def iife268_():
                coll100_ = _dafny.Map()
                compr_100_: int
                for compr_100_ in _dafny.IntegerRange(106, 964):
                    d_2498_v37_: int = compr_100_
                    if ((106) <= (d_2498_v37_)) and ((d_2498_v37_) < (964)):
                        coll100_[default__.safeModuloInt(d_2498_v37_, p0)] = p0
                return _dafny.Map(coll100_)
            d_2497_v38_ = D11_DC27((0) - (len((d_2495_v36_).intersection(_dafny.Set({iife268_()
})))), 699, not ((d_2492_v32_)[default__.safeIndex(794, (d_2492_v32_).length(0))]) or (d_2453_v0_))
            d_2497_v38_ = d_2497_v38_
        d_2499_v39_: _dafny.Array
        def lambda119_(d_2500_v0_):
            def lambda120_(d_2501_i1_):
                return d_2500_v0_

            return lambda120_

        init70_ = lambda119_(d_2453_v0_)
        nw413_ = _dafny.Array(None, 26)
        for i0_70_ in range(nw413_.length(0)):
            nw413_[i0_70_] = init70_(i0_70_)
        d_2499_v39_ = nw413_
        d_2502_v40_: C9
        nw414_ = C9()
        nw414_.ctor__(d_2499_v39_, (p0) - (p0), self.f26)
        d_2502_v40_ = nw414_
        d_2503_v41_: D12
        d_2503_v41_ = D12_DC29()
        d_2504_v42_: D12
        d_2504_v42_ = D12_DC30(d_2503_v41_)
        pat_let_tv64_ = d_2503_v41_
        def iife269_(_pat_let84_0):
            def iife270_(d_2505_dt__update__tmp_h0_):
                def iife271_(_pat_let85_0):
                    def iife272_(d_2506_dt__update_hcf44_h0_):
                        return D12_DC30(d_2506_dt__update_hcf44_h0_)
                    return iife272_(_pat_let85_0)
                return iife271_(pat_let_tv64_)
            return iife270_(_pat_let84_0)
        d_2504_v42_ = iife269_(d_2504_v42_)
        d_2507_i2_: int
        d_2507_i2_ = 0
        with _dafny.label("20"):
            while d_2453_v0_:
                with _dafny.c_label("20"):
                    if (d_2507_i2_) >= (100):
                        raise _dafny.Break("20")
                    d_2507_i2_ = (d_2507_i2_) + (1)
                    d_2508_v43_: _dafny.Seq
                    d_2508_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtbcu"))
                    d_2509_v44_: _dafny.Seq
                    d_2509_v44_ = _dafny.SeqWithoutIsStrInference([d_2508_v43_])
                    d_2508_v43_ = (d_2509_v44_)[default__.safeIndex(p0, len(d_2509_v44_))]
                    d_2510_v45_: _dafny.Seq
                    d_2510_v45_ = _dafny.SeqWithoutIsStrInference([False, d_2453_v0_])
                    d_2511_v46_: _dafny.Seq
                    d_2511_v46_ = _dafny.SeqWithoutIsStrInference([d_2510_v45_])
                    d_2512_v47_: _dafny.Map
                    d_2512_v47_ = _dafny.Map({((d_2511_v46_)[default__.safeIndex((d_2502_v40_).f34, len(d_2511_v46_))]) + (d_2510_v45_): d_2453_v0_})
                    (globalState).f2 = len(d_2512_v47_)
                    d_2513_v48_: _dafny.Array
                    def lambda121_(d_2514_i3_):
                        return default__.safeModuloInt(d_2514_i3_, (self).f30)

                    init71_ = lambda121_
                    nw415_ = _dafny.Array(None, 19)
                    for i0_71_ in range(nw415_.length(0)):
                        nw415_[i0_71_] = init71_(i0_71_)
                    d_2513_v48_ = nw415_
                    d_2515_v49_: _dafny.Map
                    d_2515_v49_ = _dafny.Map({d_2513_v48_: default__.safeModuloInt((d_2502_v40_).f34, p0)})
                    d_2515_v49_ = d_2515_v49_
                    if d_2453_v0_:
                        d_2516_v50_: _dafny.Set
                        d_2516_v50_ = _dafny.Set({d_2453_v0_, d_2453_v0_})
                        d_2517_v51_: _dafny.Set
                        d_2517_v51_ = _dafny.Set({d_2453_v0_, (len(_dafny.Map({d_2453_v0_: d_2453_v0_}))) == ((d_2502_v40_).f34), d_2453_v0_, ((d_2502_v40_).f34) >= (len(d_2516_v50_)), (p0) <= ((d_2502_v40_).f34)})
                        d_2518_v52_: _dafny.Map
                        d_2518_v52_ = _dafny.Map({_dafny.CodePoint('p'): (d_2502_v40_).f34})
                        d_2519_v53_: _dafny.Seq
                        d_2519_v53_ = _dafny.SeqWithoutIsStrInference([d_2517_v51_, d_2516_v50_])
                        d_2520_v54_: _dafny.Map
                        d_2520_v54_ = _dafny.Map({len(d_2518_v52_): (d_2519_v53_)[default__.safeIndex((0) - ((self).f30), len(d_2519_v53_))]})
                        d_2517_v51_ = ((d_2520_v54_)[default__.safeDivisionInt((self).f30, p0)] if (default__.safeDivisionInt((self).f30, p0)) in (d_2520_v54_) else _dafny.Set({False, d_2453_v0_}))
                        d_2521_v55_: C5
                        nw416_ = C5()
                        nw416_.ctor__(default__.safeModuloInt((self).fm3(globalState), len(_dafny.Set({d_2499_v39_}))))
                        d_2521_v55_ = nw416_
                        d_2522_v56_: D5
                        d_2522_v56_ = D5_DC11(d_2453_v0_)
                        d_2523_v57_: T0
                        nw417_ = C8()
                        nw417_.ctor__(d_2453_v0_, d_2522_v56_, default__.fm45(globalState))
                        d_2523_v57_ = nw417_
                        d_2523_v57_ = d_2523_v57_
                        arr14_ = d_2502_v40_.f33
                        index389_ = default__.safeIndex(666, (d_2502_v40_.f33).length(0))
                        arr14_[index389_] = ((0) - ((d_2502_v40_).f34)) <= (p0)
                        arr15_ = d_2502_v40_.f33
                        index390_ = default__.safeIndex(666, (d_2502_v40_.f33).length(0))
                        arr15_[index390_] = (d_2510_v45_)[default__.safeIndex((d_2523_v57_).fm3(globalState), len(d_2510_v45_))]
                        d_2524_v58_: _dafny.Array
                        def lambda122_(d_2525_i4_):
                            return _dafny.CodePoint('o')

                        init72_ = lambda122_
                        nw418_ = _dafny.Array(None, 9)
                        for i0_72_ in range(nw418_.length(0)):
                            nw418_[i0_72_] = init72_(i0_72_)
                        d_2524_v58_ = nw418_
                        d_2526_v59_: D0
                        d_2526_v59_ = D0_DC0(d_2524_v58_)
                        d_2526_v59_ = d_2526_v59_
                    elif True:
                        d_2527_v60_: C6
                        nw419_ = C6()
                        nw419_.ctor__()
                        d_2527_v60_ = nw419_
                        d_2528_v61_: D3
                        d_2528_v61_ = D3_DC7(d_2508_v43_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2529_i5_ in range(default__.abs(299))])), len(d_2508_v43_), (d_2502_v40_).fm3(globalState), d_2510_v45_)
                        d_2530_v62_: str
                        d_2530_v62_ = _dafny.CodePoint('s')
                        d_2531_v63_: _dafny.Set
                        d_2531_v63_ = _dafny.Set({default__.safeModuloInt((d_2502_v40_).f34, (d_2502_v40_).f34), default__.safeDivisionInt((self).fm3(globalState), (self).f30), len(_dafny.SeqWithoutIsStrInference([((D3_DC7(d_2508_v43_, len((d_2528_v61_).cf8), 567, len(d_2508_v43_), _dafny.SeqWithoutIsStrInference([True]))).cf8).set(default__.safeIndex((d_2502_v40_).f34, len((D3_DC7(d_2508_v43_, len((d_2528_v61_).cf8), 567, len(d_2508_v43_), _dafny.SeqWithoutIsStrInference([True]))).cf8)), d_2530_v62_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agpj"))]))})
                        d_2531_v63_ = d_2531_v63_
                        (globalState).f13 = d_2453_v0_
                        d_2532_v64_: _dafny.Map
                        d_2532_v64_ = _dafny.Map({d_2530_v62_: d_2453_v0_})
                        d_2533_v65_: _dafny.MultiSet
                        d_2533_v65_ = _dafny.MultiSet([d_2453_v0_, ((d_2532_v64_)[d_2530_v62_] if (d_2530_v62_) in (d_2532_v64_) else d_2453_v0_), d_2453_v0_, d_2453_v0_, d_2453_v0_])
                        (globalState).f12 = _dafny.Map({((d_2533_v65_)[default__.fm0((d_2502_v40_).f34, globalState)] if (default__.fm0((d_2502_v40_).f34, globalState)) in (d_2533_v65_) else p0): 210})
                        (globalState).f21 = (d_2502_v40_).f34
                    pass
            pass
        r0 = d_2453_v0_
        return r0

    def m8(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        (globalState).f13 = p1
        (self).f26 = (self.f26).set(p0, p0)
        (globalState).f13 = (p0) <= (p0)
        d_2534_v0_: _dafny.Seq
        d_2534_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2535_i0_: int
        d_2535_i0_ = 0
        with _dafny.label("21"):
            while (d_2534_v0_)[default__.safeIndex(len(_dafny.Set({p1})), len(d_2534_v0_))]:
                with _dafny.c_label("21"):
                    if (d_2535_i0_) >= (100):
                        raise _dafny.Break("21")
                    d_2535_i0_ = (d_2535_i0_) + (1)
                    d_2536_v1_: _dafny.Seq
                    d_2536_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idoh"))
                    rhs394_ = p1
                    rhs395_ = not (p1) or ((p1) == (p1))
                    rhs396_ = (self).f30
                    rhs397_ = (d_2536_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2537_i1_ in range(default__.abs(689))]))
                    lhs298_ = globalState
                    lhs299_ = globalState
                    lhs300_ = globalState
                    lhs298_.f13 = rhs394_
                    lhs299_.f1 = rhs395_
                    lhs300_.f22 = rhs396_
                    d_2536_v1_ = rhs397_
                    (globalState).f1 = p1
                    d_2536_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2538_i2_ in range(default__.abs(806))])
                    (globalState).f0 = default__.safeModuloInt(p0, (self).f30)
                    pass
            pass
        d_2539_v2_: _dafny.Array
        nw420_ = _dafny.Array(int(0), 23)
        d_2539_v2_ = nw420_
        index391_ = default__.safeIndex(6, (d_2539_v2_).length(0))
        (d_2539_v2_)[index391_] = (self).f30
        index392_ = default__.safeIndex(6, (d_2539_v2_).length(0))
        (d_2539_v2_)[index392_] = (self).f30
        (globalState).f21 = ((self).f30) + ((self).f30)
        r0 = d_2539_v2_
        return r0

    def m9(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_2540_v0_: _dafny.Seq
        d_2540_v0_ = _dafny.SeqWithoutIsStrInference([(self).f30])
        d_2541_v1_: _dafny.MultiSet
        d_2541_v1_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f30]), d_2540_v0_])
        hi14_ = (0) - ((d_2541_v1_).cardinality)
        for d_2542_i0_ in range(default__.safeDivisionInt((self).f30, (self).f30), hi14_):
            d_2543_v2_: _dafny.Seq
            d_2543_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sir"))
            d_2544_v4_: D4
            d_2544_v4_ = D4_DC9(d_2542_i0_)
            d_2545_v5_: _dafny.Set
            d_2545_v5_ = _dafny.Set({D4_DC9(d_2542_i0_), d_2544_v4_})
            d_2546_v6_: _dafny.Array
            def lambda123_(d_2547_p1_):
                def lambda124_(d_2548_i1_):
                    return d_2547_p1_

                return lambda124_

            init73_ = lambda123_(p1)
            nw421_ = _dafny.Array(None, 29)
            for i0_73_ in range(nw421_.length(0)):
                nw421_[i0_73_] = init73_(i0_73_)
            d_2546_v6_ = nw421_
            d_2549_v7_: _dafny.Set
            d_2549_v7_ = _dafny.Set({d_2546_v6_})
            def iife273_():
                coll101_ = _dafny.Map()
                compr_101_: D4
                for compr_101_ in (d_2545_v5_).Elements:
                    d_2550_v3_: D4 = compr_101_
                    if (d_2550_v3_) in (d_2545_v5_):
                        coll101_[d_2550_v3_] = (self).f30
                return _dafny.Map(coll101_)
            rhs398_ = (len(iife273_()
            )) <= ((len(d_2543_v2_)) + (len(d_2543_v2_)))
            rhs399_ = default__.fm0(len(d_2540_v0_), globalState)
            rhs400_ = d_2542_i0_
            rhs401_ = d_2543_v2_
            rhs402_ = len(d_2549_v7_)
            lhs301_ = globalState
            lhs302_ = globalState
            lhs303_ = globalState
            lhs304_ = globalState
            lhs301_.f1 = rhs398_
            lhs302_.f1 = rhs399_
            lhs303_.f21 = rhs400_
            d_2543_v2_ = rhs401_
            lhs304_.f21 = rhs402_
            d_2551_v8_: _dafny.Map
            d_2551_v8_ = _dafny.Map({default__.safeModuloInt((self).f30, d_2542_i0_): (p2 if p2 else p2)})
            d_2551_v8_ = default__.fm37(((self).f30) == ((self).f30), globalState)
            d_2543_v2_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kn")) if p0 else d_2543_v2_)
            d_2552_v9_: _dafny.Array
            nw422_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
            d_2552_v9_ = nw422_
            d_2552_v9_ = d_2552_v9_
        d_2553_v10_: _dafny.Seq
        d_2553_v10_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, p1])
        d_2554_v11_: C5
        nw423_ = C5()
        nw423_.ctor__(len(d_2553_v10_))
        d_2554_v11_ = nw423_
        d_2555_v12_: C10
        nw424_ = C10()
        nw424_.ctor__(((d_2554_v11_).f38) * (342), (p1 if p0 else True))
        d_2555_v12_ = nw424_
        d_2556_v13_: _dafny.Map
        d_2556_v13_ = _dafny.Map({(d_2555_v12_).f31: (d_2555_v12_).f32})
        def iife274_():
            coll102_ = _dafny.Map()
            compr_102_: int
            for compr_102_ in _dafny.IntegerRange(819, 936):
                d_2557_v14_: int = compr_102_
                if ((819) <= (d_2557_v14_)) and ((d_2557_v14_) < (936)):
                    coll102_[default__.safeDivisionInt(d_2557_v14_, 741)] = p1
            return _dafny.Map(coll102_)
        d_2556_v13_ = ((iife274_()
        ) | (d_2556_v13_)) | (d_2556_v13_)
        hi15_ = (self).f30
        for d_2558_i2_ in range((d_2555_v12_).f31, hi15_):
            d_2559_v15_: D7
            d_2559_v15_ = D7_DC16((d_2554_v11_).f38, p1)
            (globalState).f16 = ((self).f30) * ((d_2559_v15_).cf22)
            d_2560_v16_: _dafny.Seq
            d_2560_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qiocrb"))
            d_2561_v18_: _dafny.Array
            nw425_ = _dafny.Array(None, 11)
            nw425_[int(0)] = 154
            nw425_[int(1)] = d_2558_i2_
            nw425_[int(2)] = (self).f30
            nw425_[int(3)] = (d_2554_v11_).f38
            nw425_[int(4)] = (self).f30
            nw425_[int(5)] = len(d_2560_v16_)
            nw425_[int(6)] = (d_2554_v11_).f38
            def iife275_():
                coll103_ = _dafny.Set()
                compr_103_: str
                for compr_103_ in (d_2560_v16_).Elements:
                    d_2562_v17_: str = compr_103_
                    if (d_2562_v17_) in (d_2560_v16_):
                        coll103_ = coll103_.union(_dafny.Set([d_2562_v17_]))
                return _dafny.Set(coll103_)
            nw425_[int(7)] = (0) - (len(iife275_()
            ))
            nw425_[int(8)] = (d_2554_v11_).f38
            nw425_[int(9)] = (d_2555_v12_).f31
            nw425_[int(10)] = d_2558_i2_
            d_2561_v18_ = nw425_
            d_2563_v19_: _dafny.Map
            d_2563_v19_ = _dafny.Map({p1: (0) - ((d_2555_v12_).f31)})
            d_2564_v20_: _dafny.Map
            d_2564_v20_ = _dafny.Map({d_2563_v19_: p0})
            index393_ = default__.safeIndex(386, (d_2561_v18_).length(0))
            (d_2561_v18_)[index393_] = len((d_2564_v20_) | (d_2564_v20_))
            index394_ = default__.safeIndex(386, (d_2561_v18_).length(0))
            (d_2561_v18_)[index394_] = (len(d_2553_v10_)) - (d_2558_i2_)
            d_2565_v21_: T0
            nw426_ = C3()
            nw426_.ctor__(d_2561_v18_, p1, self.f26)
            d_2565_v21_ = nw426_
            d_2566_v22_: _dafny.Array
            nw427_ = _dafny.Array(False, 5)
            d_2566_v22_ = nw427_
            nw428_ = C9()
            nw428_.ctor__(d_2566_v22_, (self).f30, d_2565_v21_.f26)
            d_2565_v21_ = nw428_
            d_2567_v23_: str
            d_2567_v23_ = _dafny.CodePoint('w')
            d_2568_v24_: D5
            d_2568_v24_ = D5_DC11(p1)
            d_2569_v25_: _dafny.Array
            nw429_ = _dafny.Array(None, 22)
            nw429_[int(0)] = default__.fm35((d_2561_v18_)[default__.safeIndex(386, (d_2561_v18_).length(0))], globalState)
            nw429_[int(1)] = d_2567_v23_
            nw429_[int(2)] = (d_2567_v23_ if (d_2568_v24_).cf16 else d_2567_v23_)
            nw429_[int(3)] = d_2567_v23_
            nw429_[int(4)] = d_2567_v23_
            nw429_[int(5)] = default__.fm35(d_2558_i2_, globalState)
            nw429_[int(6)] = d_2567_v23_
            nw429_[int(7)] = d_2567_v23_
            nw429_[int(8)] = d_2567_v23_
            nw429_[int(9)] = (d_2560_v16_)[default__.safeIndex((d_2555_v12_).f31, len(d_2560_v16_))]
            nw429_[int(10)] = d_2567_v23_
            nw429_[int(11)] = d_2567_v23_
            nw429_[int(12)] = (d_2567_v23_ if p1 else d_2567_v23_)
            nw429_[int(13)] = d_2567_v23_
            nw429_[int(14)] = d_2567_v23_
            nw429_[int(15)] = d_2567_v23_
            nw429_[int(16)] = d_2567_v23_
            nw429_[int(17)] = d_2567_v23_
            nw429_[int(18)] = _dafny.CodePoint('h')
            nw429_[int(19)] = d_2567_v23_
            nw429_[int(20)] = d_2567_v23_
            nw429_[int(21)] = d_2567_v23_
            d_2569_v25_ = nw429_
            index395_ = default__.safeIndex(463, (d_2569_v25_).length(0))
            (d_2569_v25_)[index395_] = d_2567_v23_
            index396_ = default__.safeIndex(463, (d_2569_v25_).length(0))
            (d_2569_v25_)[index396_] = d_2567_v23_
        (globalState).f13 = default__.fm0((self).f30, globalState)
        r0 = (d_2555_v12_).f31
        return r0

    @property
    def f30(self):
        return self._f30

class C12(T0):
    def  __init__(self):
        self._f26: _dafny.Map = _dafny.Map({})
        self.f28: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C12"
    @property
    def f26(self):
        return self._f26
    @f26.setter
    def f26(self, value):
        self._f26 = value
    def ctor__(self, f28, f26):
        (self).f28 = f28
        (self).f26 = f26

    def fm3(self, globalState):
        return default__.safeModuloInt(117, 909)

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife276_():
            coll104_ = _dafny.Map()
            compr_104_: int
            for compr_104_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgrhnris")))])).Elements:
                d_2570_v0_: int = compr_104_
                if (d_2570_v0_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgrhnris")))])):
                    coll104_[(d_2570_v0_) + (-494)] = len(_dafny.Set({685}))
            return _dafny.Map(coll104_)
        def iife277_():
            coll105_ = _dafny.Map()
            compr_105_: str
            for compr_105_ in (_dafny.Map({(D4_DC8(_dafny.CodePoint('n'))).cf13: True})).keys.Elements:
                d_2571_v1_: str = compr_105_
                if (d_2571_v1_) in (_dafny.Map({(D4_DC8(_dafny.CodePoint('n'))).cf13: True})):
                    coll105_[d_2571_v1_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')]))
            return _dafny.Map(coll105_)
        return ((_dafny.Map({(0) - (len(iife276_()
        )): _dafny.Map({_dafny.CodePoint('r'): 774})}) if not(not(True)) else _dafny.Map({593: _dafny.Map({_dafny.CodePoint('r'): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality})}))) | ((_dafny.Map({(_dafny.MultiSet([False, True, False])).cardinality: _dafny.Map({_dafny.CodePoint('v'): 78})})) | (_dafny.Map({619: iife277_()
        })))

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        d_2572_v0_: bool
        d_2572_v0_ = False
        if d_2572_v0_:
            d_2573_v1_: _dafny.Seq
            d_2573_v1_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(0) - (p1), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2574_i0_ in range(default__.abs(358))])), p2, p1])) + (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2575_i1_ in range(default__.abs(769))]))))]))])
            (globalState).f16 = len((d_2573_v1_)[default__.safeIndex(((0) - (default__.fm2(globalState))) - (641), len(d_2573_v1_))])
            d_2576_v2_: _dafny.Array
            def lambda125_(d_2577_v0_):
                def lambda126_(d_2578_i2_):
                    return d_2577_v0_

                return lambda126_

            init74_ = lambda125_(d_2572_v0_)
            nw430_ = _dafny.Array(None, 24)
            for i0_74_ in range(nw430_.length(0)):
                nw430_[i0_74_] = init74_(i0_74_)
            d_2576_v2_ = nw430_
            d_2579_v3_: C9
            nw431_ = C9()
            nw431_.ctor__((d_2576_v2_ if d_2572_v0_ else d_2576_v2_), p2, self.f26)
            d_2579_v3_ = nw431_
            arr16_ = d_2579_v3_.f33
            index397_ = default__.safeIndex(312, (d_2579_v3_.f33).length(0))
            arr16_[index397_] = (d_2572_v0_) == (d_2572_v0_)
            arr17_ = d_2579_v3_.f33
            index398_ = default__.safeIndex(312, (d_2579_v3_.f33).length(0))
            arr17_[index398_] = not(d_2572_v0_)
            d_2580_v4_: _dafny.Array
            nw432_ = _dafny.Array(int(0), 20)
            d_2580_v4_ = nw432_
            index399_ = default__.safeIndex(176, (d_2580_v4_).length(0))
            (d_2580_v4_)[index399_] = (0) - (p1)
            index400_ = default__.safeIndex(176, (d_2580_v4_).length(0))
            (d_2580_v4_)[index400_] = p2
            if (d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))]:
                rhs403_ = (D8_DC20(d_2572_v0_)).cf31
                rhs404_ = default__.fm0((d_2580_v4_)[default__.safeIndex(176, (d_2580_v4_).length(0))], globalState)
                rhs405_ = (d_2579_v3_).fm3(globalState)
                rhs406_ = d_2580_v4_
                lhs305_ = globalState
                lhs306_ = globalState
                lhs307_ = globalState
                lhs305_.f1 = rhs403_
                lhs306_.f13 = rhs404_
                lhs307_.f22 = rhs405_
                d_2580_v4_ = rhs406_
                d_2581_v5_: _dafny.Seq
                d_2581_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ce"))
                d_2582_v6_: bool
                d_2583_v7_: str
                out94_: bool
                out95_: str
                out94_, out95_ = (self).m6((0) - (default__.safeDivisionInt(-546, p2)), len((d_2581_v5_) + (d_2581_v5_)), globalState)
                d_2582_v6_ = out94_
                d_2583_v7_ = out95_
                d_2584_v8_: D0
                d_2584_v8_ = D0_DC1(d_2580_v4_, _dafny.CodePoint('p'), (d_2580_v4_)[default__.safeIndex(176, (d_2580_v4_).length(0))])
                d_2585_v9_: D0
                d_2585_v9_ = D0_DC2(d_2584_v8_)
                d_2586_v10_: D0
                d_2586_v10_ = D0_DC2(d_2585_v9_)
                d_2587_v11_: _dafny.Set
                d_2587_v11_ = _dafny.Set({D0_DC2(d_2585_v9_), d_2586_v10_})
                d_2588_v12_: _dafny.Seq
                d_2588_v12_ = _dafny.SeqWithoutIsStrInference([d_2582_v6_])
                d_2589_v13_: _dafny.Map
                d_2589_v13_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([(d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))]])})
                arr18_ = d_2579_v3_.f33
                index401_ = default__.safeIndex(312, (d_2579_v3_.f33).length(0))
                rhs407_ = not((d_2588_v12_) < ((_dafny.SeqWithoutIsStrInference([d_2572_v0_])) + (((d_2589_v13_)[p1] if (p1) in (d_2589_v13_) else d_2588_v12_))))
                rhs408_ = d_2587_v11_
                rhs409_ = (d_2581_v5_).set(default__.safeIndex(len(d_2581_v5_), len(d_2581_v5_)), d_2583_v7_)
                lhs308_ = d_2579_v3_.f33
                lhs309_ = default__.safeIndex(312, (d_2579_v3_.f33).length(0))
                lhs308_[lhs309_] = rhs407_
                d_2587_v11_ = rhs408_
                d_2581_v5_ = rhs409_
                d_2590_v14_: _dafny.Array
                def lambda127_(d_2591_v6_, d_2592_v0_, d_2593_v3_):
                    def lambda128_(d_2594_i3_):
                        return D7_DC17(d_2591_v6_, d_2592_v0_, (d_2593_v3_).f34)

                    return lambda128_

                init75_ = lambda127_(d_2582_v6_, d_2572_v0_, d_2579_v3_)
                nw433_ = _dafny.Array(None, 1)
                for i0_75_ in range(nw433_.length(0)):
                    nw433_[i0_75_] = init75_(i0_75_)
                d_2590_v14_ = nw433_
                d_2595_v15_: D7
                d_2595_v15_ = D7_DC17(d_2572_v0_, (d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))], (d_2579_v3_).f34)
                index402_ = default__.safeIndex(772, (d_2590_v14_).length(0))
                (d_2590_v14_)[index402_] = d_2595_v15_
                d_2596_v16_: _dafny.Seq
                d_2596_v16_ = _dafny.SeqWithoutIsStrInference([d_2580_v4_, d_2580_v4_, d_2580_v4_])
                index403_ = default__.safeIndex(772, (d_2590_v14_).length(0))
                rhs410_ = D7_DC17(d_2582_v6_, (d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))], len(d_2596_v16_))
                rhs411_ = ((p2) - ((0) - (p2))) * ((0) - ((d_2579_v3_).f34))
                lhs310_ = d_2590_v14_
                lhs311_ = default__.safeIndex(772, (d_2590_v14_).length(0))
                lhs310_[lhs311_] = rhs410_
                r2 = rhs411_
                (d_2579_v3_).f33 = d_2576_v2_
            elif True:
                (globalState).f22 = p2
                arr19_ = d_2579_v3_.f33
                index404_ = default__.safeIndex(312, (d_2579_v3_.f33).length(0))
                arr19_[index404_] = (d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))]
                d_2576_v2_ = d_2579_v3_.f33
                d_2597_v17_: _dafny.Set
                d_2597_v17_ = _dafny.Set({(d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))], (d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))], (d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))], d_2572_v0_, d_2572_v0_})
                d_2598_v18_: D10
                d_2598_v18_ = D10_DC25(p1, (d_2580_v4_)[default__.safeIndex(176, (d_2580_v4_).length(0))], (d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))], d_2572_v0_)
                (globalState).f0 = len((d_2597_v17_ if d_2572_v0_ else (_dafny.Set({(d_2598_v18_).cf37})) - (d_2597_v17_)))
                d_2599_v19_: _dafny.Array
                nw434_ = _dafny.Array(_dafny.MultiSet({}), 21)
                d_2599_v19_ = nw434_
                index405_ = default__.safeIndex(331, (d_2599_v19_).length(0))
                (d_2599_v19_)[index405_] = _dafny.MultiSet([(d_2579_v3_.f33)[default__.safeIndex(312, (d_2579_v3_.f33).length(0))]])
                index406_ = default__.safeIndex(331, (d_2599_v19_).length(0))
                (d_2599_v19_)[index406_] = default__.fm31(globalState)
        elif True:
            if d_2572_v0_:
                d_2600_v20_: _dafny.Seq
                d_2600_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bddfsr"))
                d_2601_v21_: _dafny.MultiSet
                d_2601_v21_ = _dafny.MultiSet([d_2600_v20_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2602_i4_ in range(default__.abs(698))]), d_2600_v20_, d_2600_v20_, d_2600_v20_])
                d_2603_v22_: _dafny.MultiSet
                d_2603_v22_ = d_2601_v21_
                d_2603_v22_ = default__.fm50(globalState)
                d_2604_v23_: _dafny.Array
                def lambda129_(d_2605_v0_):
                    def lambda130_(d_2606_i5_):
                        return d_2605_v0_

                    return lambda130_

                init76_ = lambda129_(d_2572_v0_)
                nw435_ = _dafny.Array(None, 20)
                for i0_76_ in range(nw435_.length(0)):
                    nw435_[i0_76_] = init76_(i0_76_)
                d_2604_v23_ = nw435_
                d_2607_v24_: _dafny.Map
                d_2607_v24_ = _dafny.Map({d_2604_v23_: d_2572_v0_})
                d_2607_v24_ = (d_2607_v24_).set(d_2604_v23_, not(not (d_2572_v0_) or (not(d_2572_v0_))))
                d_2600_v20_ = (d_2600_v20_) + (d_2600_v20_)
                d_2608_v25_: _dafny.Map
                d_2608_v25_ = _dafny.Map({d_2572_v0_: d_2572_v0_})
                d_2609_v26_: _dafny.MultiSet
                d_2609_v26_ = _dafny.MultiSet([d_2608_v25_])
                (globalState).f5 = ((d_2609_v26_)[(_dafny.Map({d_2572_v0_: d_2572_v0_})) | (d_2608_v25_)] if ((_dafny.Map({d_2572_v0_: d_2572_v0_})) | (d_2608_v25_)) in (d_2609_v26_) else 281)
                d_2610_v27_: _dafny.Set
                d_2610_v27_ = _dafny.Set({d_2572_v0_})
                rhs412_ = d_2610_v27_
                rhs413_ = d_2572_v0_
                lhs312_ = globalState
                d_2610_v27_ = rhs412_
                lhs312_.f13 = rhs413_
            elif True:
                d_2611_v28_: _dafny.Seq
                d_2611_v28_ = _dafny.SeqWithoutIsStrInference([d_2572_v0_])
                d_2611_v28_ = (d_2611_v28_) + (d_2611_v28_)
                d_2612_v29_: _dafny.Seq
                d_2612_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icrth"))
                d_2613_v30_: _dafny.Array
                nw436_ = _dafny.Array(None, 4)
                nw436_[int(0)] = p2
                nw436_[int(1)] = len((_dafny.Map({p1: p1})) | (self.f26))
                nw436_[int(2)] = (0) - ((default__.fm2(globalState)) - ((0) - (len(d_2612_v29_))))
                nw436_[int(3)] = (p1) * (p1)
                d_2613_v30_ = nw436_
                index407_ = default__.safeIndex(433, (d_2613_v30_).length(0))
                (d_2613_v30_)[index407_] = p2
                index408_ = default__.safeIndex(433, (d_2613_v30_).length(0))
                (d_2613_v30_)[index408_] = (p1) + (default__.safeModuloInt(len(d_2612_v29_), p1))
                d_2614_v31_: _dafny.Array
                nw437_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_2614_v31_ = nw437_
                index409_ = default__.safeIndex(72, (d_2614_v31_).length(0))
                (d_2614_v31_)[index409_] = (d_2612_v29_) + (d_2612_v29_)
                d_2615_v32_: _dafny.Map
                d_2615_v32_ = _dafny.Map({775: d_2612_v29_})
                index410_ = default__.safeIndex(72, (d_2614_v31_).length(0))
                (d_2614_v31_)[index410_] = ((d_2612_v29_ if False else d_2612_v29_)) + (((d_2615_v32_)[(0) - (p2)] if ((0) - (p2)) in (d_2615_v32_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))
                d_2616_v33_: _dafny.Array
                nw438_ = _dafny.Array(_dafny.Map({}), 18)
                d_2616_v33_ = nw438_
                d_2617_v34_: D5
                d_2617_v34_ = D5_DC11(d_2572_v0_)
                d_2618_v35_: str
                d_2618_v35_ = _dafny.CodePoint('c')
                d_2619_v36_: _dafny.Map
                d_2619_v36_ = _dafny.Map({d_2617_v34_: d_2618_v35_})
                index411_ = default__.safeIndex(286, (d_2616_v33_).length(0))
                (d_2616_v33_)[index411_] = d_2619_v36_
                d_2620_v37_: _dafny.Seq
                d_2620_v37_ = _dafny.SeqWithoutIsStrInference([d_2619_v36_, d_2619_v36_, default__.fm61(globalState)])
                index412_ = default__.safeIndex(286, (d_2616_v33_).length(0))
                index413_ = default__.safeIndex(433, (d_2613_v30_).length(0))
                rhs414_ = True
                rhs415_ = d_2572_v0_
                rhs416_ = (d_2620_v37_)[default__.safeIndex(len(d_2611_v28_), len(d_2620_v37_))]
                rhs417_ = p1
                lhs313_ = globalState
                lhs314_ = globalState
                lhs315_ = d_2616_v33_
                lhs316_ = default__.safeIndex(286, (d_2616_v33_).length(0))
                lhs317_ = d_2613_v30_
                lhs318_ = default__.safeIndex(433, (d_2613_v30_).length(0))
                lhs313_.f1 = rhs414_
                lhs314_.f1 = rhs415_
                lhs315_[lhs316_] = rhs416_
                lhs317_[lhs318_] = rhs417_
                index414_ = default__.safeIndex(72, (d_2614_v31_).length(0))
                (d_2614_v31_)[index414_] = d_2612_v29_
            d_2572_v0_ = not (False) or (d_2572_v0_)
            d_2621_v38_: _dafny.Array
            nw439_ = _dafny.Array(D11.default()(), 13)
            d_2621_v38_ = nw439_
            d_2622_v39_: _dafny.Set
            d_2622_v39_ = _dafny.Set({self.f28})
            d_2623_v40_: _dafny.Seq
            d_2623_v40_ = _dafny.SeqWithoutIsStrInference([d_2572_v0_, d_2572_v0_])
            d_2624_v41_: D11
            d_2624_v41_ = D11_DC27(p1, len(d_2622_v39_), (d_2623_v40_)[default__.safeIndex((self).fm3(globalState), len(d_2623_v40_))])
            index415_ = default__.safeIndex(399, (d_2621_v38_).length(0))
            (d_2621_v38_)[index415_] = d_2624_v41_
            d_2625_v42_: D5
            d_2625_v42_ = D5_DC11(False)
            d_2626_v43_: D8
            d_2626_v43_ = D8_DC19(d_2572_v0_, d_2572_v0_, (d_2625_v42_).cf16)
            d_2627_v44_: _dafny.MultiSet
            d_2627_v44_ = _dafny.MultiSet([(d_2626_v43_).cf29, d_2572_v0_, not(d_2572_v0_), d_2572_v0_])
            d_2628_v45_: _dafny.Set
            d_2628_v45_ = _dafny.Set({p2, (0) - (p2)})
            d_2629_v46_: _dafny.Map
            d_2629_v46_ = _dafny.Map({d_2572_v0_: default__.fm0((0) - (p1), globalState)})
            d_2630_v47_: _dafny.Seq
            d_2630_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ciwma"))
            d_2631_v48_: str
            d_2631_v48_ = _dafny.CodePoint('d')
            d_2632_v49_: _dafny.Array
            nw440_ = _dafny.Array(None, 18)
            nw440_[int(0)] = p1
            nw440_[int(1)] = p2
            nw440_[int(2)] = p2
            nw440_[int(3)] = p2
            nw440_[int(4)] = 479
            nw440_[int(5)] = p1
            nw440_[int(6)] = p2
            nw440_[int(7)] = p2
            nw440_[int(8)] = p1
            nw440_[int(9)] = p1
            nw440_[int(10)] = (p1) - (p2)
            nw440_[int(11)] = (p2) * (len(d_2628_v45_))
            nw440_[int(12)] = (p2) - (p2)
            nw440_[int(13)] = p1
            nw440_[int(14)] = p1
            nw440_[int(15)] = (p1) - (len(d_2629_v46_))
            nw440_[int(16)] = len((d_2630_v47_) + ((d_2630_v47_).set(default__.safeIndex(len(self.f28), len(d_2630_v47_)), d_2631_v48_)))
            nw440_[int(17)] = p2
            d_2632_v49_ = nw440_
            index416_ = default__.safeIndex(659, (d_2632_v49_).length(0))
            (d_2632_v49_)[index416_] = p2
            index417_ = default__.safeIndex(399, (d_2621_v38_).length(0))
            index418_ = default__.safeIndex(659, (d_2632_v49_).length(0))
            rhs418_ = not(d_2572_v0_)
            rhs419_ = p1
            rhs420_ = d_2624_v41_
            rhs421_ = ((d_2627_v44_).intersection(d_2627_v44_)).intersection(d_2627_v44_)
            rhs422_ = p1
            lhs319_ = globalState
            lhs320_ = globalState
            lhs321_ = d_2621_v38_
            lhs322_ = default__.safeIndex(399, (d_2621_v38_).length(0))
            lhs323_ = d_2632_v49_
            lhs324_ = default__.safeIndex(659, (d_2632_v49_).length(0))
            lhs319_.f13 = rhs418_
            lhs320_.f0 = rhs419_
            lhs321_[lhs322_] = rhs420_
            d_2627_v44_ = rhs421_
            lhs323_[lhs324_] = rhs422_
            (globalState).f1 = default__.fm0(p1, globalState)
            d_2633_v50_: C0
            nw441_ = C0()
            nw441_.ctor__((p0).cf0)
            d_2633_v50_ = nw441_
        d_2634_v51_: C4
        nw442_ = C4()
        nw442_.ctor__()
        d_2634_v51_ = nw442_
        d_2635_v52_: _dafny.Array
        nw443_ = _dafny.Array(int(0), 21)
        d_2635_v52_ = nw443_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_2635_v52_).length(0)):
            d_2636_i6_: int = guard_loop_10_
            if (True) and (((0) <= (d_2636_i6_)) and ((d_2636_i6_) < ((d_2635_v52_).length(0)))):
                (d_2635_v52_)[(d_2636_i6_)] = default__.safeModuloInt(d_2636_i6_, len((_dafny.Map({d_2572_v0_: p2})) | (_dafny.Map({d_2572_v0_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faetrw")))}))))
        d_2572_v0_ = not((d_2572_v0_ if d_2572_v0_ else d_2572_v0_))
        d_2637_v53_: _dafny.Seq
        d_2637_v53_ = _dafny.SeqWithoutIsStrInference([p1, p2, p1, p2])
        rhs423_ = not((p1) not in (d_2637_v53_))
        rhs424_ = d_2635_v52_
        rhs425_ = (739) + (p1)
        lhs325_ = globalState
        lhs326_ = globalState
        lhs325_.f1 = rhs423_
        d_2635_v52_ = rhs424_
        lhs326_.f0 = rhs425_
        d_2638_v54_: D21
        d_2638_v54_ = D21_DC43(d_2572_v0_, not(d_2572_v0_))
        (globalState).f13 = (d_2638_v54_).cf60
        r0 = ((p2) * (p2)) + (p2)
        d_2639_v55_: _dafny.Seq
        d_2639_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrfskm"))
        d_2640_v56_: _dafny.Array
        nw444_ = _dafny.Array(None, 23)
        nw444_[int(0)] = True
        nw444_[int(1)] = d_2572_v0_
        nw444_[int(2)] = not(d_2572_v0_)
        nw444_[int(3)] = d_2572_v0_
        nw444_[int(4)] = False
        nw444_[int(5)] = d_2572_v0_
        nw444_[int(6)] = d_2572_v0_
        nw444_[int(7)] = d_2572_v0_
        nw444_[int(8)] = d_2572_v0_
        nw444_[int(9)] = d_2572_v0_
        nw444_[int(10)] = False
        nw444_[int(11)] = d_2572_v0_
        nw444_[int(12)] = True
        nw444_[int(13)] = d_2572_v0_
        nw444_[int(14)] = d_2572_v0_
        nw444_[int(15)] = d_2572_v0_
        nw444_[int(16)] = d_2572_v0_
        nw444_[int(17)] = d_2572_v0_
        nw444_[int(18)] = d_2572_v0_
        nw444_[int(19)] = d_2572_v0_
        nw444_[int(20)] = d_2572_v0_
        nw444_[int(21)] = d_2572_v0_
        nw444_[int(22)] = False
        d_2640_v56_ = nw444_
        d_2641_v57_: _dafny.Map
        d_2641_v57_ = _dafny.Map({(d_2639_v55_ if d_2572_v0_ else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2642_i7_ in range(default__.abs(599))])): d_2640_v56_})
        r1 = d_2641_v57_
        r2 = p2
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        d_2643_v0_: _dafny.Seq
        d_2643_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmniyppy"))
        source35_ = _dafny.MultiSet([d_2643_v0_])
        d_2644___mcc_h0_ = source35_
        d_2645_cf45_ = d_2644___mcc_h0_
        d_2646_v1_: bool
        d_2646_v1_ = True
        d_2647_v2_: _dafny.Seq
        d_2647_v2_ = _dafny.SeqWithoutIsStrInference([d_2646_v1_])
        d_2648_v3_: int
        d_2648_v3_ = 607
        d_2649_v4_: D4
        d_2649_v4_ = D4_DC10(d_2648_v3_)
        d_2650_v5_: _dafny.Array
        nw445_ = _dafny.Array(None, 6)
        nw445_[int(0)] = d_2647_v2_
        nw445_[int(1)] = d_2647_v2_
        nw445_[int(2)] = d_2647_v2_
        nw445_[int(3)] = d_2647_v2_
        nw445_[int(4)] = _dafny.SeqWithoutIsStrInference([d_2646_v1_, d_2646_v1_, not(default__.fm0((d_2649_v4_).cf15, globalState)), d_2646_v1_, True])
        nw445_[int(5)] = d_2647_v2_
        d_2650_v5_ = nw445_
        index419_ = default__.safeIndex(234, (d_2650_v5_).length(0))
        (d_2650_v5_)[index419_] = _dafny.SeqWithoutIsStrInference([d_2646_v1_, d_2646_v1_])
        index420_ = default__.safeIndex(234, (d_2650_v5_).length(0))
        (d_2650_v5_)[index420_] = (_dafny.SeqWithoutIsStrInference([d_2646_v1_, d_2646_v1_])) + (_dafny.SeqWithoutIsStrInference([d_2646_v1_]))
        r2 = d_2646_v1_
        (self).f28 = (self.f28).set(d_2648_v3_, default__.fm0(d_2648_v3_, globalState))
        if ((d_2650_v5_)[default__.safeIndex(234, (d_2650_v5_).length(0))])[default__.safeIndex((d_2648_v3_ if True else d_2648_v3_), len((d_2650_v5_)[default__.safeIndex(234, (d_2650_v5_).length(0))]))]:
            (globalState).f22 = d_2648_v3_
            d_2651_v6_: _dafny.Array
            nw446_ = _dafny.Array(int(0), 11)
            d_2651_v6_ = nw446_
            def lambda131_(d_2652_v3_):
                def lambda132_(d_2653_i0_):
                    return (d_2653_i0_) - (d_2652_v3_)

                return lambda132_

            init77_ = lambda131_(d_2648_v3_)
            nw447_ = _dafny.Array(None, 18)
            for i0_77_ in range(nw447_.length(0)):
                nw447_[i0_77_] = init77_(i0_77_)
            d_2651_v6_ = nw447_
            d_2654_v7_: _dafny.Set
            d_2654_v7_ = _dafny.Set({(0) - (d_2648_v3_)})
            d_2654_v7_ = d_2654_v7_
            (globalState).f1 = d_2646_v1_
            d_2655_v8_: D8
            d_2655_v8_ = D8_DC19(((self.f28)[d_2648_v3_] if (d_2648_v3_) in (self.f28) else d_2646_v1_), False, True)
            d_2656_v9_: D8
            d_2656_v9_ = D8_DC21(d_2655_v8_)
            d_2656_v9_ = d_2656_v9_
        elif True:
            d_2657_v10_: C5
            nw448_ = C5()
            nw448_.ctor__(d_2648_v3_)
            d_2657_v10_ = nw448_
            d_2658_v11_: D2
            d_2658_v11_ = D2_DC5()
            d_2659_v12_: _dafny.Set
            d_2659_v12_ = _dafny.Set({D2_DC5(), d_2658_v11_, d_2658_v11_, d_2658_v11_, d_2658_v11_})
            d_2660_v13_: _dafny.Array
            def lambda133_(d_2661_v1_):
                def lambda134_(d_2662_i1_):
                    return d_2661_v1_

                return lambda134_

            init78_ = lambda133_(d_2646_v1_)
            nw449_ = _dafny.Array(None, 2)
            for i0_78_ in range(nw449_.length(0)):
                nw449_[i0_78_] = init78_(i0_78_)
            d_2660_v13_ = nw449_
            d_2663_v14_: _dafny.Set
            d_2663_v14_ = _dafny.Set({d_2660_v13_, d_2660_v13_})
            d_2664_v15_: D23
            d_2664_v15_ = D23_DC45(d_2663_v14_)
            d_2665_v16_: _dafny.Array
            nw450_ = _dafny.Array(int(0), 27)
            d_2665_v16_ = nw450_
            d_2666_v17_: _dafny.Map
            d_2666_v17_ = _dafny.Map({d_2665_v16_: d_2648_v3_})
            d_2667_v18_: _dafny.Map
            d_2667_v18_ = _dafny.Map({d_2646_v1_: d_2660_v13_})
            d_2668_v19_: _dafny.Set
            d_2668_v19_ = _dafny.Set({d_2646_v1_})
            d_2669_v20_: _dafny.Seq
            d_2669_v20_ = _dafny.SeqWithoutIsStrInference([(d_2657_v10_).f38, d_2648_v3_, d_2648_v3_])
            d_2670_v21_: _dafny.Array
            nw451_ = _dafny.Array(None, 29)
            nw451_[int(0)] = d_2646_v1_
            nw451_[int(1)] = (d_2659_v12_).issubset(default__.fm62(self.f26, d_2646_v1_, globalState))
            nw451_[int(2)] = d_2646_v1_
            nw451_[int(3)] = (d_2663_v14_) != ((d_2664_v15_).cf62)
            nw451_[int(4)] = ((d_2650_v5_)[default__.safeIndex(234, (d_2650_v5_).length(0))])[default__.safeIndex(d_2648_v3_, len((d_2650_v5_)[default__.safeIndex(234, (d_2650_v5_).length(0))]))]
            nw451_[int(5)] = d_2646_v1_
            nw451_[int(6)] = d_2646_v1_
            nw451_[int(7)] = not(d_2646_v1_)
            nw451_[int(8)] = d_2646_v1_
            nw451_[int(9)] = d_2646_v1_
            nw451_[int(10)] = d_2646_v1_
            nw451_[int(11)] = (d_2648_v3_) != (((d_2666_v17_)[d_2665_v16_] if (d_2665_v16_) in (d_2666_v17_) else len(d_2667_v18_)))
            nw451_[int(12)] = (d_2648_v3_) < ((d_2657_v10_).f38)
            nw451_[int(13)] = (d_2668_v19_).ispropersubset(d_2668_v19_)
            nw451_[int(14)] = d_2646_v1_
            nw451_[int(15)] = not((len(d_2643_v0_)) <= (d_2648_v3_))
            nw451_[int(16)] = d_2646_v1_
            nw451_[int(17)] = True
            nw451_[int(18)] = d_2646_v1_
            nw451_[int(19)] = d_2646_v1_
            nw451_[int(20)] = d_2646_v1_
            nw451_[int(21)] = d_2646_v1_
            nw451_[int(22)] = d_2646_v1_
            nw451_[int(23)] = d_2646_v1_
            nw451_[int(24)] = d_2646_v1_
            nw451_[int(25)] = d_2646_v1_
            nw451_[int(26)] = (d_2669_v20_) <= ((d_2669_v20_).set(default__.safeIndex(624, len(d_2669_v20_)), (d_2657_v10_).f38))
            nw451_[int(27)] = d_2646_v1_
            nw451_[int(28)] = d_2646_v1_
            d_2670_v21_ = nw451_
            index421_ = default__.safeIndex(707, (d_2670_v21_).length(0))
            (d_2670_v21_)[index421_] = default__.fm0(d_2648_v3_, globalState)
            d_2671_v22_: _dafny.Array
            nw452_ = _dafny.Array(None, 21)
            d_2671_v22_ = nw452_
            d_2672_v23_: C2
            nw453_ = C2()
            nw453_.ctor__(d_2646_v1_, (0) - ((d_2657_v10_).f38))
            d_2672_v23_ = nw453_
            index422_ = default__.safeIndex(663, (d_2671_v22_).length(0))
            (d_2671_v22_)[index422_] = d_2672_v23_
            index423_ = default__.safeIndex(707, (d_2670_v21_).length(0))
            index424_ = default__.safeIndex(663, (d_2671_v22_).length(0))
            rhs426_ = ((self.f28)[d_2648_v3_] if (d_2648_v3_) in (self.f28) else (d_2672_v23_).f41)
            rhs427_ = d_2672_v23_
            lhs327_ = d_2670_v21_
            lhs328_ = default__.safeIndex(707, (d_2670_v21_).length(0))
            lhs329_ = d_2671_v22_
            lhs330_ = default__.safeIndex(663, (d_2671_v22_).length(0))
            lhs327_[lhs328_] = rhs426_
            lhs329_[lhs330_] = rhs427_
            d_2673_v24_: str
            d_2673_v24_ = _dafny.CodePoint('b')
            (globalState).f5 = (len(((d_2643_v0_) + (d_2643_v0_)).set(default__.safeIndex(d_2648_v3_, len((d_2643_v0_) + (d_2643_v0_))), d_2673_v24_))) - (((d_2672_v23_).f42) + (718))
            d_2674_v25_: _dafny.Set
            d_2674_v25_ = _dafny.Set({d_2648_v3_})
            d_2675_v26_: _dafny.Array
            nw454_ = _dafny.Array(None, 3)
            nw454_[int(0)] = d_2674_v25_
            nw454_[int(1)] = d_2674_v25_
            nw454_[int(2)] = d_2674_v25_
            d_2675_v26_ = nw454_
            d_2676_v27_: _dafny.Map
            d_2676_v27_ = _dafny.Map({True: d_2675_v26_})
            d_2676_v27_ = (d_2676_v27_).set(((d_2650_v5_)[default__.safeIndex(234, (d_2650_v5_).length(0))])[default__.safeIndex(271, len((d_2650_v5_)[default__.safeIndex(234, (d_2650_v5_).length(0))]))], d_2675_v26_)
            d_2677_v28_: D7
            d_2677_v28_ = D7_DC16(len((d_2650_v5_)[default__.safeIndex(234, (d_2650_v5_).length(0))]), True)
            d_2678_v29_: _dafny.Map
            d_2678_v29_ = _dafny.Map({d_2677_v28_: (d_2670_v21_)[default__.safeIndex(707, (d_2670_v21_).length(0))]})
            d_2678_v29_ = (d_2678_v29_).set(D7_DC16(len(_dafny.SeqWithoutIsStrInference([(d_2670_v21_)[default__.safeIndex(707, (d_2670_v21_).length(0))]])), d_2646_v1_), d_2646_v1_)
        d_2679_v30_: bool
        d_2679_v30_ = False
        d_2680_v31_: int
        d_2680_v31_ = 619
        r0 = not((100) <= ((len(d_2643_v0_) if d_2679_v30_ else (0) - (d_2680_v31_))))
        (globalState).f21 = d_2680_v31_
        d_2681_v32_: _dafny.Seq
        d_2681_v32_ = _dafny.SeqWithoutIsStrInference([d_2679_v30_])
        d_2682_v33_: D21
        d_2682_v33_ = D21_DC43(d_2679_v30_, d_2679_v30_)
        d_2683_v34_: _dafny.Set
        d_2683_v34_ = _dafny.Set({len(d_2681_v32_)})
        d_2684_v35_: _dafny.Map
        d_2684_v35_ = _dafny.Map({False: d_2679_v30_})
        d_2685_v36_: _dafny.Array
        nw455_ = _dafny.Array(None, 15)
        nw455_[int(0)] = (_dafny.MultiSet(d_2681_v32_)).issubset(_dafny.MultiSet([d_2679_v30_]))
        nw455_[int(1)] = (d_2682_v33_).cf60
        nw455_[int(2)] = d_2679_v30_
        nw455_[int(3)] = d_2679_v30_
        nw455_[int(4)] = d_2679_v30_
        nw455_[int(5)] = d_2679_v30_
        nw455_[int(6)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pr"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unl")))
        nw455_[int(7)] = (d_2683_v34_).issubset(d_2683_v34_)
        nw455_[int(8)] = d_2679_v30_
        nw455_[int(9)] = d_2679_v30_
        nw455_[int(10)] = False
        nw455_[int(11)] = not((d_2679_v30_ if d_2679_v30_ else d_2679_v30_))
        nw455_[int(12)] = ((d_2684_v35_)[d_2679_v30_] if (d_2679_v30_) in (d_2684_v35_) else d_2679_v30_)
        nw455_[int(13)] = d_2679_v30_
        nw455_[int(14)] = d_2679_v30_
        d_2685_v36_ = nw455_
        index425_ = default__.safeIndex(117, (d_2685_v36_).length(0))
        (d_2685_v36_)[index425_] = d_2679_v30_
        index426_ = default__.safeIndex(117, (d_2685_v36_).length(0))
        (d_2685_v36_)[index426_] = d_2679_v30_
        d_2686_v37_: _dafny.Array
        nw456_ = _dafny.Array(_dafny.Set({}), 8)
        d_2686_v37_ = nw456_
        d_2687_v38_: _dafny.Seq
        d_2687_v38_ = _dafny.SeqWithoutIsStrInference([d_2686_v37_, d_2686_v37_, d_2686_v37_, d_2686_v37_, d_2686_v37_])
        d_2688_v39_: _dafny.Array
        nw457_ = _dafny.Array(None, 19)
        nw457_[int(0)] = d_2686_v37_
        nw457_[int(1)] = d_2686_v37_
        nw457_[int(2)] = d_2686_v37_
        nw457_[int(3)] = d_2686_v37_
        nw457_[int(4)] = d_2686_v37_
        nw457_[int(5)] = d_2686_v37_
        nw457_[int(6)] = d_2686_v37_
        nw457_[int(7)] = d_2686_v37_
        nw457_[int(8)] = d_2686_v37_
        nw457_[int(9)] = (d_2687_v38_)[default__.safeIndex(174, len(d_2687_v38_))]
        nw457_[int(10)] = (d_2686_v37_ if (d_2685_v36_)[default__.safeIndex(117, (d_2685_v36_).length(0))] else d_2686_v37_)
        nw457_[int(11)] = d_2686_v37_
        nw457_[int(12)] = d_2686_v37_
        nw457_[int(13)] = d_2686_v37_
        nw457_[int(14)] = d_2686_v37_
        nw457_[int(15)] = d_2686_v37_
        nw457_[int(16)] = d_2686_v37_
        nw457_[int(17)] = d_2686_v37_
        nw457_[int(18)] = d_2686_v37_
        d_2688_v39_ = nw457_
        index427_ = default__.safeIndex(862, (d_2688_v39_).length(0))
        (d_2688_v39_)[index427_] = d_2686_v37_
        index428_ = default__.safeIndex(862, (d_2688_v39_).length(0))
        (d_2688_v39_)[index428_] = d_2686_v37_
        d_2689_v40_: _dafny.Array
        nw458_ = _dafny.Array(int(0), 18)
        d_2689_v40_ = nw458_
        d_2690_v41_: str
        d_2690_v41_ = _dafny.CodePoint('v')
        d_2691_v42_: D0
        d_2691_v42_ = D0_DC1(d_2689_v40_, d_2690_v41_, d_2680_v31_)
        d_2692_i2_: int
        d_2692_i2_ = 0
        with _dafny.label("22"):
            while (((d_2691_v42_).cf3) + (d_2680_v31_)) == (d_2680_v31_):
                with _dafny.c_label("22"):
                    if (d_2692_i2_) >= (100):
                        raise _dafny.Break("22")
                    d_2692_i2_ = (d_2692_i2_) + (1)
                    (globalState).f5 = d_2680_v31_
                    d_2643_v0_ = _dafny.SeqWithoutIsStrInference([d_2690_v41_ for d_2693_i3_ in range(default__.abs(601))])
                    r0 = not ((d_2685_v36_)[default__.safeIndex(117, (d_2685_v36_).length(0))]) or (default__.fm0(d_2680_v31_, globalState))
                    d_2694_v43_: _dafny.Array
                    nw459_ = _dafny.Array(_dafny.Array(None, 0), 15)
                    d_2694_v43_ = nw459_
                    index429_ = default__.safeIndex(569, (d_2694_v43_).length(0))
                    (d_2694_v43_)[index429_] = d_2689_v40_
                    d_2695_v44_: _dafny.MultiSet
                    d_2695_v44_ = _dafny.MultiSet([d_2679_v30_, d_2679_v30_])
                    d_2696_v47_: T0
                    nw460_ = C9()
                    def iife278_():
                        def iife280_():
                            coll108_ = _dafny.Set()
                            compr_108_: int
                            for compr_108_ in _dafny.IntegerRange(741, 528):
                                d_2699_v46_: int = compr_108_
                                if ((741) <= (d_2699_v46_)) and ((d_2699_v46_) < (528)):
                                    coll108_ = coll108_.union(_dafny.Set([(d_2699_v46_) - (d_2680_v31_)]))
                            return _dafny.Set(coll108_)
                        coll106_ = _dafny.Map()
                        def iife279_():
                            coll107_ = _dafny.Set()
                            compr_107_: int
                            for compr_107_ in _dafny.IntegerRange(741, 528):
                                d_2697_v46_: int = compr_107_
                                if ((741) <= (d_2697_v46_)) and ((d_2697_v46_) < (528)):
                                    coll107_ = coll107_.union(_dafny.Set([(d_2697_v46_) - (d_2680_v31_)]))
                            return _dafny.Set(coll107_)
                        compr_106_: int
                        for compr_106_ in (iife279_()
                        ).Elements:
                            d_2698_v45_: int = compr_106_
                            if (d_2698_v45_) in (iife280_()
                            ):
                                coll106_[(d_2698_v45_) - (len(_dafny.Set({_dafny.Map({True: d_2680_v31_})})))] = d_2680_v31_
                        return _dafny.Map(coll106_)
                    nw460_.ctor__(d_2685_v36_, d_2680_v31_, iife278_()
                    )
                    d_2696_v47_ = nw460_
                    d_2700_v48_: _dafny.MultiSet
                    d_2700_v48_ = _dafny.MultiSet([d_2696_v47_])
                    d_2701_v49_: _dafny.Set
                    d_2701_v49_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_2679_v30_: len(_dafny.Map({d_2679_v30_: d_2680_v31_}))})) for d_2702_i4_ in range(default__.abs(671))])})
                    index430_ = default__.safeIndex(569, (d_2694_v43_).length(0))
                    nw461_ = _dafny.Array(None, 21)
                    nw461_[int(0)] = d_2680_v31_
                    nw461_[int(1)] = d_2680_v31_
                    nw461_[int(2)] = 568
                    nw461_[int(3)] = len((_dafny.Set({d_2679_v30_, not((d_2685_v36_)[default__.safeIndex(117, (d_2685_v36_).length(0))]), d_2679_v30_})) | (_dafny.Set({(d_2685_v36_)[default__.safeIndex(117, (d_2685_v36_).length(0))]})))
                    nw461_[int(4)] = d_2680_v31_
                    nw461_[int(5)] = 940
                    nw461_[int(6)] = (d_2680_v31_) - (((d_2695_v44_)[d_2679_v30_] if (d_2679_v30_) in (d_2695_v44_) else d_2680_v31_))
                    nw461_[int(7)] = -865
                    nw461_[int(8)] = (d_2700_v48_).cardinality
                    nw461_[int(9)] = (self).fm3(globalState)
                    nw461_[int(10)] = d_2680_v31_
                    nw461_[int(11)] = d_2680_v31_
                    nw461_[int(12)] = len(d_2643_v0_)
                    nw461_[int(13)] = (d_2680_v31_) * (len(d_2643_v0_))
                    nw461_[int(14)] = default__.safeModuloInt(len(d_2701_v49_), d_2680_v31_)
                    nw461_[int(15)] = len(d_2681_v32_)
                    nw461_[int(16)] = d_2680_v31_
                    nw461_[int(17)] = len((d_2643_v0_) + (d_2643_v0_))
                    nw461_[int(18)] = d_2680_v31_
                    nw461_[int(19)] = d_2680_v31_
                    nw461_[int(20)] = (0) - (default__.safeModuloInt(len(d_2696_v47_.f26), (0) - ((0) - (d_2680_v31_))))
                    (d_2694_v43_)[index430_] = nw461_
                    pass
            pass
        r0 = d_2679_v30_
        d_2703_v50_: _dafny.Set
        d_2703_v50_ = _dafny.Set({self.f26})
        r1 = ((d_2703_v50_).intersection(d_2703_v50_)) | ((d_2703_v50_).intersection(d_2703_v50_))
        d_2704_v51_: _dafny.Seq
        d_2704_v51_ = _dafny.SeqWithoutIsStrInference([default__.fm15((d_2685_v36_)[default__.safeIndex(117, (d_2685_v36_).length(0))], globalState), _dafny.Set({d_2690_v41_})])
        d_2705_v52_: _dafny.Seq
        d_2705_v52_ = _dafny.SeqWithoutIsStrInference([d_2680_v31_])
        d_2706_v53_: _dafny.Seq
        d_2706_v53_ = _dafny.SeqWithoutIsStrInference([default__.fm15((d_2685_v36_)[default__.safeIndex(117, (d_2685_v36_).length(0))], globalState), (d_2704_v51_)[default__.safeIndex((d_2705_v52_)[default__.safeIndex(len(_dafny.Map({False: (d_2685_v36_)[default__.safeIndex(117, (d_2685_v36_).length(0))]})), len(d_2705_v52_))], len(d_2704_v51_))]])
        d_2707_v54_: _dafny.Seq
        d_2707_v54_ = _dafny.SeqWithoutIsStrInference([d_2704_v51_, d_2706_v53_])
        r2 = ((d_2680_v31_) * (d_2680_v31_)) != (len((d_2707_v54_)[default__.safeIndex(d_2680_v31_, len(d_2707_v54_))]))
        return r0, r1, r2

    def m5(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        (globalState).f21 = p2
        d_2708_v0_: _dafny.Array
        def lambda135_(d_2709_p3_):
            def lambda136_(d_2710_i0_):
                return not((d_2709_p3_) not in (_dafny.Map({d_2709_p3_: d_2709_p3_})))

            return lambda136_

        init79_ = lambda135_(p3)
        nw462_ = _dafny.Array(None, 2)
        for i0_79_ in range(nw462_.length(0)):
            nw462_[i0_79_] = init79_(i0_79_)
        d_2708_v0_ = nw462_
        index431_ = default__.safeIndex(67, (d_2708_v0_).length(0))
        (d_2708_v0_)[index431_] = p3
        index432_ = default__.safeIndex(67, (d_2708_v0_).length(0))
        (d_2708_v0_)[index432_] = (p1) <= (p1)
        d_2711_v1_: D11
        d_2711_v1_ = D11_DC27(p2, p2, p3)
        d_2711_v1_ = d_2711_v1_
        d_2712_v2_: _dafny.Set
        d_2712_v2_ = _dafny.Set({p3, (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))], (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]})
        if (d_2712_v2_).ispropersubset(d_2712_v2_):
            (globalState).f0 = p1
            if not(p3):
                r0 = (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]
                (globalState).f21 = p1
                (globalState).f13 = ((p1) - (p1)) == (p1)
                index433_ = default__.safeIndex(67, (d_2708_v0_).length(0))
                (d_2708_v0_)[index433_] = p3
                d_2713_v3_: _dafny.Seq
                d_2713_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sv"))
                d_2714_v4_: _dafny.Map
                d_2714_v4_ = _dafny.Map({(d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]: p1})
                d_2715_v5_: _dafny.Seq
                d_2715_v5_ = _dafny.SeqWithoutIsStrInference([len(d_2714_v4_)])
                d_2716_v6_: _dafny.Map
                d_2716_v6_ = _dafny.Map({p3: (d_2715_v5_)[default__.safeIndex((p0).fm3(globalState), len(d_2715_v5_))]})
                d_2717_v7_: _dafny.Seq
                d_2717_v7_ = _dafny.SeqWithoutIsStrInference([d_2713_v3_])
                d_2718_v8_: _dafny.MultiSet
                d_2718_v8_ = _dafny.MultiSet([False])
                d_2719_v9_: _dafny.Array
                nw463_ = _dafny.Array(None, 23)
                nw463_[int(0)] = p1
                nw463_[int(1)] = (len(d_2713_v3_)) + (len(d_2714_v4_))
                nw463_[int(2)] = default__.safeModuloInt(p2, p2)
                nw463_[int(3)] = default__.safeModuloInt((p0).fm3(globalState), p2)
                nw463_[int(4)] = default__.safeModuloInt(p2, p2)
                nw463_[int(5)] = p1
                nw463_[int(6)] = (0) - (default__.safeModuloInt(p1, 9))
                nw463_[int(7)] = len(d_2715_v5_)
                nw463_[int(8)] = p1
                nw463_[int(9)] = (0) - (default__.safeModuloInt(p1, p1))
                nw463_[int(10)] = (p1) * (len(d_2717_v7_))
                nw463_[int(11)] = p1
                nw463_[int(12)] = (p1) + (p1)
                nw463_[int(13)] = p1
                nw463_[int(14)] = default__.safeModuloInt(len(d_2715_v5_), p1)
                nw463_[int(15)] = p2
                nw463_[int(16)] = p1
                nw463_[int(17)] = (735) + (p1)
                nw463_[int(18)] = len(d_2713_v3_)
                nw463_[int(19)] = (0) - (p2)
                nw463_[int(20)] = default__.safeDivisionInt(len(d_2713_v3_), p1)
                nw463_[int(21)] = ((d_2716_v6_)[(d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]] if ((d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]) in (d_2716_v6_) else len(p0.f26))
                nw463_[int(22)] = (d_2718_v8_).cardinality
                d_2719_v9_ = nw463_
                d_2720_v10_: _dafny.Map
                d_2720_v10_ = _dafny.Map({True: (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]})
                index434_ = default__.safeIndex(121, (d_2719_v9_).length(0))
                (d_2719_v9_)[index434_] = len((d_2720_v10_) | (_dafny.Map({(d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]: (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]})))
                d_2721_v11_: str
                d_2721_v11_ = _dafny.CodePoint('o')
                d_2722_v12_: _dafny.Array
                nw464_ = _dafny.Array(None, 7)
                nw464_[int(0)] = _dafny.CodePoint('r')
                nw464_[int(1)] = _dafny.CodePoint('p')
                nw464_[int(2)] = d_2721_v11_
                nw464_[int(3)] = _dafny.CodePoint('t')
                nw464_[int(4)] = d_2721_v11_
                nw464_[int(5)] = d_2721_v11_
                nw464_[int(6)] = d_2721_v11_
                d_2722_v12_ = nw464_
                d_2723_v13_: D0
                d_2723_v13_ = D0_DC0(d_2722_v12_)
                d_2724_v14_: _dafny.Set
                d_2724_v14_ = _dafny.Set({p1, p1})
                d_2725_v15_: D8
                d_2725_v15_ = D8_DC18(d_2724_v14_)
                d_2726_v16_: _dafny.Map
                d_2726_v16_ = _dafny.Map({d_2725_v15_: p1})
                index435_ = default__.safeIndex(121, (d_2719_v9_).length(0))
                (d_2719_v9_)[index435_] = ((len(_dafny.SeqWithoutIsStrInference([p2]))) - (-428)) * (len((((d_2715_v5_).set(default__.safeIndex(p1, len(d_2715_v5_)), p2)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvvyhlgf"))), len((d_2715_v5_).set(default__.safeIndex(p1, len(d_2715_v5_)), p2))), p2)).set(default__.safeIndex((_dafny.MultiSet([d_2723_v13_, d_2723_v13_])).cardinality, len(((d_2715_v5_).set(default__.safeIndex(p1, len(d_2715_v5_)), p2)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvvyhlgf"))), len((d_2715_v5_).set(default__.safeIndex(p1, len(d_2715_v5_)), p2))), p2))), ((d_2726_v16_)[d_2725_v15_] if (d_2725_v15_) in (d_2726_v16_) else len(d_2712_v2_)))))
            elif True:
                d_2727_v17_: str
                d_2727_v17_ = _dafny.CodePoint('j')
                d_2728_v18_: _dafny.Map
                d_2728_v18_ = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p2, p2]))).cardinality: (_dafny.CodePoint('f') if False else d_2727_v17_)})
                d_2728_v18_ = d_2728_v18_
                d_2729_v19_: _dafny.Array
                nw465_ = _dafny.Array(D23.default()(), 12)
                d_2729_v19_ = nw465_
                d_2730_v20_: _dafny.Map
                d_2730_v20_ = _dafny.Map({p2: d_2729_v19_})
                d_2731_v21_: _dafny.Array
                nw466_ = _dafny.Array(None, 19)
                nw466_[int(0)] = d_2729_v19_
                nw466_[int(1)] = d_2729_v19_
                nw466_[int(2)] = d_2729_v19_
                nw466_[int(3)] = d_2729_v19_
                nw466_[int(4)] = d_2729_v19_
                nw466_[int(5)] = d_2729_v19_
                nw466_[int(6)] = d_2729_v19_
                nw466_[int(7)] = d_2729_v19_
                nw466_[int(8)] = ((d_2730_v20_)[p1] if (p1) in (d_2730_v20_) else d_2729_v19_)
                nw466_[int(9)] = d_2729_v19_
                nw466_[int(10)] = d_2729_v19_
                nw466_[int(11)] = d_2729_v19_
                nw466_[int(12)] = d_2729_v19_
                nw466_[int(13)] = d_2729_v19_
                nw466_[int(14)] = d_2729_v19_
                nw466_[int(15)] = d_2729_v19_
                nw466_[int(16)] = d_2729_v19_
                nw466_[int(17)] = (d_2729_v19_ if (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))] else d_2729_v19_)
                nw466_[int(18)] = d_2729_v19_
                d_2731_v21_ = nw466_
                index436_ = default__.safeIndex(340, (d_2731_v21_).length(0))
                (d_2731_v21_)[index436_] = d_2729_v19_
                index437_ = default__.safeIndex(340, (d_2731_v21_).length(0))
                (d_2731_v21_)[index437_] = d_2729_v19_
                d_2732_v22_: _dafny.Array
                nw467_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_2732_v22_ = nw467_
                index438_ = default__.safeIndex(248, (d_2732_v22_).length(0))
                (d_2732_v22_)[index438_] = d_2727_v17_
                d_2733_v23_: _dafny.Seq
                d_2733_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aahhnp"))
                index439_ = default__.safeIndex(67, (d_2708_v0_).length(0))
                index440_ = default__.safeIndex(248, (d_2732_v22_).length(0))
                rhs428_ = (p2) < (843)
                rhs429_ = (d_2733_v23_) in (_dafny.Set({d_2733_v23_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))).set(default__.safeIndex(default__.fm2(globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))), d_2727_v17_), d_2733_v23_, d_2733_v23_, d_2733_v23_}))
                rhs430_ = d_2708_v0_
                rhs431_ = d_2727_v17_
                lhs331_ = globalState
                lhs332_ = d_2708_v0_
                lhs333_ = default__.safeIndex(67, (d_2708_v0_).length(0))
                lhs334_ = d_2732_v22_
                lhs335_ = default__.safeIndex(248, (d_2732_v22_).length(0))
                lhs331_.f13 = rhs428_
                lhs332_[lhs333_] = rhs429_
                d_2708_v0_ = rhs430_
                lhs334_[lhs335_] = rhs431_
                index441_ = default__.safeIndex(67, (d_2708_v0_).length(0))
                (d_2708_v0_)[index441_] = (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]
                d_2734_v24_: _dafny.Array
                nw468_ = _dafny.Array(None, 26)
                nw468_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
                nw468_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2727_v17_ for d_2735_i1_ in range(default__.abs(611))])
                nw468_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ix"))
                nw468_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
                nw468_[int(4)] = d_2733_v23_
                nw468_[int(5)] = d_2733_v23_
                nw468_[int(6)] = d_2733_v23_
                nw468_[int(7)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfbq"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfbq")))), (d_2732_v22_)[default__.safeIndex(248, (d_2732_v22_).length(0))])
                nw468_[int(8)] = d_2733_v23_
                nw468_[int(9)] = d_2733_v23_
                nw468_[int(10)] = d_2733_v23_
                nw468_[int(11)] = d_2733_v23_
                nw468_[int(12)] = d_2733_v23_
                nw468_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwxeq"))
                nw468_[int(14)] = d_2733_v23_
                nw468_[int(15)] = d_2733_v23_
                nw468_[int(16)] = d_2733_v23_
                nw468_[int(17)] = d_2733_v23_
                nw468_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbmp"))
                nw468_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnktipv"))
                nw468_[int(20)] = d_2733_v23_
                nw468_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aomttc"))
                nw468_[int(22)] = d_2733_v23_
                nw468_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrnluwmnu"))
                nw468_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkl"))
                nw468_[int(25)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmkel"))
                d_2734_v24_ = nw468_
                d_2736_v25_: _dafny.Map
                d_2736_v25_ = _dafny.Map({(p2) * ((p0).fm3(globalState)): d_2734_v24_})
                d_2736_v25_ = (d_2736_v25_).set(p1, d_2734_v24_)
            d_2737_v26_: _dafny.Seq
            d_2737_v26_ = _dafny.SeqWithoutIsStrInference([p1, 139])
            d_2738_v27_: _dafny.Set
            d_2738_v27_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([p1, p1]), (d_2737_v26_) + (d_2737_v26_)})
            d_2739_v28_: C11
            nw469_ = C11()
            nw469_.ctor__(p1, (p0.f26 if (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))] else default__.fm45(globalState)))
            d_2739_v28_ = nw469_
            d_2740_v29_: D7
            d_2740_v29_ = D7_DC17(p3, False, p1)
            rhs432_ = default__.fm16((p2 if p3 else -951), not(not((d_2740_v29_).cf25)), (0) - ((d_2739_v28_).f30), globalState)
            rhs433_ = d_2739_v28_
            d_2738_v27_ = rhs432_
            d_2739_v28_ = rhs433_
            d_2741_v30_: str
            d_2741_v30_ = _dafny.CodePoint('f')
            d_2742_v31_: _dafny.Seq
            d_2742_v31_ = _dafny.SeqWithoutIsStrInference([(d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))], p3, p3])
            d_2743_v32_: _dafny.Array
            nw470_ = _dafny.Array(None, 20)
            nw470_[int(0)] = d_2741_v30_
            nw470_[int(1)] = _dafny.CodePoint('q')
            nw470_[int(2)] = default__.fm35(p2, globalState)
            nw470_[int(3)] = d_2741_v30_
            nw470_[int(4)] = _dafny.CodePoint('r')
            nw470_[int(5)] = d_2741_v30_
            nw470_[int(6)] = d_2741_v30_
            nw470_[int(7)] = d_2741_v30_
            nw470_[int(8)] = d_2741_v30_
            nw470_[int(9)] = d_2741_v30_
            nw470_[int(10)] = d_2741_v30_
            nw470_[int(11)] = d_2741_v30_
            nw470_[int(12)] = d_2741_v30_
            nw470_[int(13)] = d_2741_v30_
            nw470_[int(14)] = d_2741_v30_
            nw470_[int(15)] = d_2741_v30_
            nw470_[int(16)] = default__.fm35(len(d_2742_v31_), globalState)
            nw470_[int(17)] = d_2741_v30_
            nw470_[int(18)] = d_2741_v30_
            nw470_[int(19)] = d_2741_v30_
            d_2743_v32_ = nw470_
            d_2744_v33_: C0
            nw471_ = C0()
            nw471_.ctor__(d_2743_v32_)
            d_2744_v33_ = nw471_
            d_2745_v34_: D5
            d_2745_v34_ = D5_DC11((d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))])
            d_2746_v35_: C8
            nw472_ = C8()
            nw472_.ctor__((p1) <= ((d_2739_v28_).f30), d_2745_v34_, (p0.f26 if p3 else p0.f26))
            d_2746_v35_ = nw472_
        elif True:
            d_2747_v36_: _dafny.Seq
            d_2747_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmd"))
            d_2748_v38_: T1
            nw473_ = C11()
            def iife281_():
                coll109_ = _dafny.Map()
                compr_109_: int
                for compr_109_ in (_dafny.SeqWithoutIsStrInference([-413 for d_2749_i3_ in range(default__.abs(-641))])).Elements:
                    d_2750_v37_: int = compr_109_
                    if (d_2750_v37_) in (_dafny.SeqWithoutIsStrInference([-413 for d_2749_i3_ in range(default__.abs(-641))])):
                        coll109_[default__.safeModuloInt(d_2750_v37_, p1)] = p1
                return _dafny.Map(coll109_)
            nw473_.ctor__(p2, iife281_()
            )
            d_2748_v38_ = nw473_
            d_2751_v39_: _dafny.MultiSet
            d_2751_v39_ = _dafny.MultiSet([d_2748_v38_])
            d_2752_v40_: str
            d_2752_v40_ = _dafny.CodePoint('y')
            d_2747_v36_ = ((_dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('y') if True else _dafny.CodePoint('k')) for d_2753_i2_ in range(default__.abs(-838))])).set(default__.safeIndex(default__.safeDivisionInt((d_2751_v39_).cardinality, p2), len(_dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('y') if True else _dafny.CodePoint('k')) for d_2754_i2_ in range(default__.abs(-838))]))), d_2752_v40_)).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('y') if True else _dafny.CodePoint('k')) for d_2755_i2_ in range(default__.abs(-838))])).set(default__.safeIndex(default__.safeDivisionInt((d_2751_v39_).cardinality, p2), len(_dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('y') if True else _dafny.CodePoint('k')) for d_2756_i2_ in range(default__.abs(-838))]))), d_2752_v40_))), d_2752_v40_)
            (globalState).f5 = p1
            d_2757_v41_: _dafny.Array
            nw474_ = _dafny.Array(int(0), 20)
            d_2757_v41_ = nw474_
            d_2758_v42_: _dafny.Map
            d_2758_v42_ = _dafny.Map({p3: d_2752_v40_})
            d_2759_v43_: _dafny.MultiSet
            d_2759_v43_ = _dafny.MultiSet([p2])
            d_2760_v44_: _dafny.Map
            d_2760_v44_ = _dafny.Map({d_2758_v42_: ((d_2759_v43_)[p1] if (p1) in (d_2759_v43_) else (0) - (p1))})
            d_2761_v45_: _dafny.Seq
            d_2761_v45_ = _dafny.SeqWithoutIsStrInference([len(p0.f26), len(d_2760_v44_), p2])
            d_2762_v46_: _dafny.Map
            d_2762_v46_ = _dafny.Map({d_2757_v41_: not((d_2761_v45_) < (d_2761_v45_))})
            d_2762_v46_ = _dafny.Map({d_2757_v41_: (p1) != (p2)})
            d_2763_v47_: _dafny.Seq
            d_2763_v47_ = _dafny.SeqWithoutIsStrInference([False])
            r0 = ((d_2763_v47_)[default__.safeIndex(p2, len(d_2763_v47_))]) or ((d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))])
            (self).f28 = (self.f28) | (_dafny.Map({p1: (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]}))
        if (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]:
            r0 = p3
            d_2764_v48_: D11
            d_2764_v48_ = D11_DC26(d_2708_v0_)
            d_2708_v0_ = (d_2708_v0_ if (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))] else (d_2764_v48_).cf39)
            d_2765_v49_: _dafny.Seq
            d_2765_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpqlkg"))
            r0 = not((d_2765_v49_) <= ((d_2765_v49_) + (d_2765_v49_)))
            d_2766_v50_: C4
            nw475_ = C4()
            nw475_.ctor__()
            d_2766_v50_ = nw475_
            (globalState).f0 = (p1) * (p1)
        elif True:
            d_2767_v51_: _dafny.Set
            d_2767_v51_ = _dafny.Set({p1, (self).fm3(globalState)})
            d_2768_v52_: _dafny.Seq
            d_2768_v52_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_2769_v53_: _dafny.Set
            d_2769_v53_ = _dafny.Set({len((d_2767_v51_) - (d_2767_v51_)), default__.fm2(globalState), len((d_2768_v52_).set(default__.safeIndex(p2, len(d_2768_v52_)), len(_dafny.Set({p1})))), p2})
            rhs434_ = d_2767_v51_
            rhs435_ = (p0).fm3(globalState)
            rhs436_ = p2
            lhs336_ = globalState
            lhs337_ = globalState
            d_2769_v53_ = rhs434_
            lhs336_.f2 = rhs435_
            lhs337_.f16 = rhs436_
            if (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]:
                d_2770_v54_: _dafny.Array
                nw476_ = _dafny.Array(int(0), 11)
                d_2770_v54_ = nw476_
                index442_ = default__.safeIndex(678, (d_2770_v54_).length(0))
                (d_2770_v54_)[index442_] = ((_dafny.MultiSet([p1])).cardinality) * (p2)
                d_2771_v55_: _dafny.Seq
                d_2771_v55_ = _dafny.SeqWithoutIsStrInference([p3, True])
                d_2772_v56_: _dafny.Seq
                d_2772_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                d_2773_v57_: D7
                d_2773_v57_ = D7_DC17(not ((d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]) or (p3), (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))], p2)
                d_2774_v58_: _dafny.Map
                d_2774_v58_ = _dafny.Map({d_2770_v54_: 186})
                d_2775_v59_: _dafny.Map
                d_2775_v59_ = _dafny.Map({d_2772_v56_: (_dafny.MultiSet([d_2774_v58_, d_2774_v58_])).intersection(_dafny.MultiSet([d_2774_v58_]))})
                index443_ = default__.safeIndex(678, (d_2770_v54_).length(0))
                rhs437_ = len((d_2769_v53_).intersection(d_2769_v53_))
                rhs438_ = (_dafny.SeqWithoutIsStrInference([p3])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([p3]))), not(not (p3) or (p3)))
                rhs439_ = (d_2772_v56_) + (d_2772_v56_)
                rhs440_ = d_2773_v57_
                rhs441_ = d_2775_v59_
                lhs338_ = d_2770_v54_
                lhs339_ = default__.safeIndex(678, (d_2770_v54_).length(0))
                lhs338_[lhs339_] = rhs437_
                d_2771_v55_ = rhs438_
                d_2772_v56_ = rhs439_
                d_2773_v57_ = rhs440_
                d_2775_v59_ = rhs441_
                d_2776_v60_: _dafny.Seq
                d_2776_v60_ = _dafny.SeqWithoutIsStrInference([d_2708_v0_])
                d_2777_v61_: _dafny.Seq
                d_2777_v61_ = _dafny.SeqWithoutIsStrInference([d_2772_v56_])
                d_2778_v62_: _dafny.Map
                d_2778_v62_ = _dafny.Map({d_2770_v54_: p1})
                d_2779_v63_: D10
                d_2779_v63_ = D10_DC25(len(d_2778_v62_), (d_2770_v54_)[default__.safeIndex(678, (d_2770_v54_).length(0))], not(p3), p3)
                d_2780_v64_: _dafny.Map
                d_2780_v64_ = _dafny.Map({(d_2776_v60_) <= (d_2776_v60_): (len((d_2777_v61_)[default__.safeIndex((d_2779_v63_).cf35, len(d_2777_v61_))])) > (p1)})
                d_2781_v65_: _dafny.Seq
                d_2781_v65_ = _dafny.SeqWithoutIsStrInference([p0.f26])
                d_2780_v64_ = (d_2780_v64_).set(p3, ((d_2768_v52_)[default__.safeIndex(p1, len(d_2768_v52_))]) not in ((d_2781_v65_)[default__.safeIndex(p1, len(d_2781_v65_))]))
                d_2782_v66_: C11
                nw477_ = C11()
                nw477_.ctor__((d_2770_v54_)[default__.safeIndex(678, (d_2770_v54_).length(0))], p0.f26)
                d_2782_v66_ = nw477_
                (globalState).f16 = (d_2770_v54_)[default__.safeIndex(678, (d_2770_v54_).length(0))]
                d_2771_v55_ = d_2771_v55_
            elif True:
                (globalState).f13 = (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]
                d_2783_v67_: _dafny.Array
                nw478_ = _dafny.Array(int(0), 21)
                d_2783_v67_ = nw478_
                d_2783_v67_ = d_2783_v67_
                d_2708_v0_ = d_2708_v0_
                d_2784_v68_: _dafny.Array
                nw479_ = _dafny.Array(_dafny.Map({}), 18)
                d_2784_v68_ = nw479_
                d_2785_v69_: _dafny.Seq
                d_2785_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uq"))
                d_2786_v70_: _dafny.Map
                d_2786_v70_ = _dafny.Map({d_2785_v69_: d_2708_v0_})
                d_2787_v71_: _dafny.Map
                d_2787_v71_ = d_2786_v70_
                index444_ = default__.safeIndex(717, (d_2784_v68_).length(0))
                (d_2784_v68_)[index444_] = d_2787_v71_
                index445_ = default__.safeIndex(717, (d_2784_v68_).length(0))
                (d_2784_v68_)[index445_] = d_2787_v71_
                d_2788_v72_: str
                d_2788_v72_ = _dafny.CodePoint('a')
                d_2789_v73_: _dafny.Array
                nw480_ = _dafny.Array(None, 20)
                nw480_[int(0)] = d_2788_v72_
                nw480_[int(1)] = d_2788_v72_
                nw480_[int(2)] = default__.fm35(p1, globalState)
                nw480_[int(3)] = d_2788_v72_
                nw480_[int(4)] = d_2788_v72_
                nw480_[int(5)] = d_2788_v72_
                nw480_[int(6)] = d_2788_v72_
                nw480_[int(7)] = d_2788_v72_
                nw480_[int(8)] = d_2788_v72_
                nw480_[int(9)] = d_2788_v72_
                nw480_[int(10)] = d_2788_v72_
                nw480_[int(11)] = d_2788_v72_
                nw480_[int(12)] = (d_2785_v69_)[default__.safeIndex(644, len(d_2785_v69_))]
                nw480_[int(13)] = d_2788_v72_
                nw480_[int(14)] = _dafny.CodePoint('w')
                nw480_[int(15)] = d_2788_v72_
                nw480_[int(16)] = d_2788_v72_
                nw480_[int(17)] = d_2788_v72_
                nw480_[int(18)] = default__.fm35(p1, globalState)
                nw480_[int(19)] = d_2788_v72_
                d_2789_v73_ = nw480_
                d_2790_v74_: _dafny.MultiSet
                d_2790_v74_ = _dafny.MultiSet([d_2789_v73_])
                (globalState).f2 = (d_2790_v74_).cardinality
            d_2791_v75_: _dafny.Array
            nw481_ = _dafny.Array(int(0), 10)
            d_2791_v75_ = nw481_
            d_2792_v76_: _dafny.Map
            d_2792_v76_ = _dafny.Map({d_2768_v52_: d_2791_v75_})
            d_2793_v77_: _dafny.Seq
            d_2793_v77_ = _dafny.SeqWithoutIsStrInference([d_2792_v76_])
            d_2794_v78_: _dafny.Map
            d_2794_v78_ = _dafny.Map({p3: p2})
            d_2795_v79_: _dafny.Seq
            d_2795_v79_ = _dafny.SeqWithoutIsStrInference([d_2791_v75_])
            d_2796_v80_: _dafny.Seq
            d_2796_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnbuvlu"))
            d_2797_v81_: _dafny.Map
            d_2797_v81_ = _dafny.Map({len(d_2796_v80_): d_2791_v75_})
            d_2798_v82_: _dafny.MultiSet
            d_2798_v82_ = _dafny.MultiSet([p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))])
            d_2799_v83_: _dafny.Array
            nw482_ = _dafny.Array(None, 27)
            nw482_[int(0)] = (d_2792_v76_) | (d_2792_v76_)
            nw482_[int(1)] = (d_2792_v76_) | (d_2792_v76_)
            nw482_[int(2)] = (d_2792_v76_).set(d_2768_v52_, d_2791_v75_)
            nw482_[int(3)] = d_2792_v76_
            nw482_[int(4)] = (d_2792_v76_) | (d_2792_v76_)
            nw482_[int(5)] = (d_2792_v76_) | (d_2792_v76_)
            nw482_[int(6)] = d_2792_v76_
            nw482_[int(7)] = ((d_2793_v77_)[default__.safeIndex(p1, len(d_2793_v77_))]) | (d_2792_v76_)
            nw482_[int(8)] = d_2792_v76_
            nw482_[int(9)] = (d_2792_v76_) | (d_2792_v76_)
            nw482_[int(10)] = (d_2792_v76_) | (d_2792_v76_)
            nw482_[int(11)] = (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_2712_v2_)]): d_2791_v75_})) | (d_2792_v76_)
            nw482_[int(12)] = d_2792_v76_
            nw482_[int(13)] = d_2792_v76_
            nw482_[int(14)] = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality for d_2800_i4_ in range(default__.abs(726))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality for d_2801_i4_ in range(default__.abs(726))]))), (0) - (len(d_2794_v78_))): d_2791_v75_})
            nw482_[int(15)] = d_2792_v76_
            nw482_[int(16)] = d_2792_v76_
            nw482_[int(17)] = ((d_2792_v76_).set(d_2768_v52_, (d_2795_v79_)[default__.safeIndex(p1, len(d_2795_v79_))])) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): d_2791_v75_}))
            nw482_[int(18)] = d_2792_v76_
            nw482_[int(19)] = _dafny.Map({d_2768_v52_: d_2791_v75_})
            nw482_[int(20)] = _dafny.Map({d_2768_v52_: ((d_2797_v81_)[p2] if (p2) in (d_2797_v81_) else d_2791_v75_)})
            nw482_[int(21)] = (_dafny.Map({d_2768_v52_: d_2791_v75_})) | (d_2792_v76_)
            nw482_[int(22)] = (d_2792_v76_) | (d_2792_v76_)
            nw482_[int(23)] = (d_2792_v76_ if p3 else _dafny.Map({d_2768_v52_: d_2791_v75_}))
            nw482_[int(24)] = _dafny.Map({(_dafny.SeqWithoutIsStrInference([-633 for d_2802_i5_ in range(default__.abs(844))])).set(default__.safeIndex(621, len(_dafny.SeqWithoutIsStrInference([-633 for d_2803_i5_ in range(default__.abs(844))]))), (d_2798_v82_).cardinality): d_2791_v75_})
            nw482_[int(25)] = d_2792_v76_
            nw482_[int(26)] = d_2792_v76_
            d_2799_v83_ = nw482_
            index446_ = default__.safeIndex(450, (d_2799_v83_).length(0))
            (d_2799_v83_)[index446_] = (_dafny.Map({d_2768_v52_: d_2791_v75_})) | (d_2792_v76_)
            d_2804_v84_: _dafny.Array
            nw483_ = _dafny.Array(_dafny.Map({}), 3)
            d_2804_v84_ = nw483_
            d_2805_v85_: _dafny.Map
            d_2805_v85_ = _dafny.Map({d_2791_v75_: d_2791_v75_})
            index447_ = default__.safeIndex(13, (d_2804_v84_).length(0))
            (d_2804_v84_)[index447_] = (d_2805_v85_) | (d_2805_v85_)
            index448_ = default__.safeIndex(450, (d_2799_v83_).length(0))
            index449_ = default__.safeIndex(13, (d_2804_v84_).length(0))
            index450_ = default__.safeIndex(67, (d_2708_v0_).length(0))
            rhs442_ = ((d_2792_v76_).set(d_2768_v52_, d_2791_v75_)) | (d_2792_v76_)
            rhs443_ = (d_2805_v85_) | (d_2805_v85_)
            rhs444_ = (p0).fm3(globalState)
            rhs445_ = ((True) and (True) if True else not(p3))
            rhs446_ = (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]
            lhs340_ = d_2799_v83_
            lhs341_ = default__.safeIndex(450, (d_2799_v83_).length(0))
            lhs342_ = d_2804_v84_
            lhs343_ = default__.safeIndex(13, (d_2804_v84_).length(0))
            lhs344_ = globalState
            lhs345_ = d_2708_v0_
            lhs346_ = default__.safeIndex(67, (d_2708_v0_).length(0))
            lhs347_ = globalState
            lhs340_[lhs341_] = rhs442_
            lhs342_[lhs343_] = rhs443_
            lhs344_.f21 = rhs444_
            lhs345_[lhs346_] = rhs445_
            lhs347_.f1 = rhs446_
            d_2806_v86_: _dafny.MultiSet
            d_2806_v86_ = _dafny.MultiSet([p3, default__.fm0(p2, globalState)])
            d_2807_v87_: _dafny.Map
            d_2807_v87_ = _dafny.Map({d_2806_v86_: _dafny.SeqWithoutIsStrInference([p3])})
            d_2808_v88_: _dafny.Array
            nw484_ = _dafny.Array(_dafny.CodePoint('D'), 11)
            d_2808_v88_ = nw484_
            d_2809_v89_: _dafny.Map
            d_2809_v89_ = _dafny.Map({d_2808_v88_: p2})
            d_2810_v90_: _dafny.Seq
            d_2810_v90_ = _dafny.SeqWithoutIsStrInference([True, (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]])
            d_2807_v87_ = (d_2807_v87_).set((((d_2806_v86_).set((d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))], default__.abs(p2))).set(p3, default__.abs(((d_2809_v89_)[d_2808_v88_] if (d_2808_v88_) in (d_2809_v89_) else (0) - (p1))))).set(p3, default__.abs((self).fm3(globalState))), d_2810_v90_)
            (globalState).f5 = p1
        d_2811_v91_: str
        d_2811_v91_ = _dafny.CodePoint('v')
        d_2812_v92_: D19
        d_2812_v92_ = D19_DC38(p1, default__.fm0((_dafny.MultiSet([d_2811_v91_, d_2811_v91_, d_2811_v91_])).cardinality, globalState), (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))])
        d_2813_v93_: _dafny.Map
        d_2813_v93_ = _dafny.Map({d_2712_v2_: (d_2812_v92_).cf52})
        d_2813_v93_ = d_2813_v93_
        d_2814_v94_: _dafny.MultiSet
        d_2814_v94_ = _dafny.MultiSet([p1])
        d_2815_v95_: _dafny.Map
        d_2815_v95_ = _dafny.Map({p3: (d_2814_v94_).cardinality})
        d_2816_v96_: _dafny.Map
        d_2816_v96_ = _dafny.Map({_dafny.Map({((d_2815_v95_)[(d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]] if ((d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]) in (d_2815_v95_) else p1): (d_2708_v0_)[default__.safeIndex(67, (d_2708_v0_).length(0))]}): d_2815_v95_})
        d_2817_v97_: D9
        d_2817_v97_ = D9_DC22(d_2816_v96_)
        pat_let_tv65_ = p3
        pat_let_tv66_ = d_2712_v2_
        pat_let_tv67_ = p1
        pat_let_tv68_ = d_2816_v96_
        def lambda137_(source36_):
            if source36_.is_DC23:
                return pat_let_tv65_
            elif True:
                d_2818___mcc_h0_ = source36_.cf33
                d_2819_cf33_ = d_2818___mcc_h0_
                return (len(pat_let_tv66_)) != (pat_let_tv67_)

        def iife282_(_pat_let86_0):
            def iife283_(d_2820_dt__update__tmp_h0_):
                def iife284_(_pat_let87_0):
                    def iife285_(d_2821_dt__update_hcf33_h0_):
                        return D9_DC22(d_2821_dt__update_hcf33_h0_)
                    return iife285_(_pat_let87_0)
                return iife284_(pat_let_tv68_)
            return iife283_(_pat_let86_0)
        r0 = lambda137_(iife282_(d_2817_v97_))
        return r0

    def m6(self, p0, p1, globalState):
        r0: bool = False
        r1: str = _dafny.CodePoint('D')
        d_2822_v0_: _dafny.Array
        nw485_ = _dafny.Array(int(0), 22)
        d_2822_v0_ = nw485_
        index451_ = default__.safeIndex(146, (d_2822_v0_).length(0))
        (d_2822_v0_)[index451_] = p1
        index452_ = default__.safeIndex(146, (d_2822_v0_).length(0))
        (d_2822_v0_)[index452_] = (p0) + (p0)
        d_2823_v1_: bool
        d_2824_v2_: bool
        d_2825_v3_: _dafny.Map
        d_2826_v4_: int
        out96_: bool
        out97_: bool
        out98_: _dafny.Map
        out99_: int
        out96_, out97_, out98_, out99_ = default__.m0(globalState)
        d_2823_v1_ = out96_
        d_2824_v2_ = out97_
        d_2825_v3_ = out98_
        d_2826_v4_ = out99_
        d_2827_v5_: _dafny.Seq
        d_2827_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mr"))
        d_2828_v6_: C3
        nw486_ = C3()
        nw486_.ctor__(d_2822_v0_, d_2823_v1_, self.f26)
        d_2828_v6_ = nw486_
        d_2829_v7_: _dafny.Map
        d_2829_v7_ = _dafny.Map({d_2824_v2_: d_2828_v6_})
        d_2830_v8_: _dafny.MultiSet
        d_2830_v8_ = _dafny.MultiSet([p1, -473, (d_2822_v0_)[default__.safeIndex(146, (d_2822_v0_).length(0))], p1, len(d_2829_v7_)])
        d_2831_v9_: str
        d_2831_v9_ = _dafny.CodePoint('b')
        d_2832_v10_: _dafny.Seq
        d_2832_v10_ = _dafny.SeqWithoutIsStrInference([d_2824_v2_, d_2824_v2_])
        d_2833_v11_: D3
        d_2833_v11_ = D3_DC7(d_2827_v5_, (((d_2830_v8_)[d_2826_v4_] if (d_2826_v4_) in (d_2830_v8_) else (d_2822_v0_)[default__.safeIndex(146, (d_2822_v0_).length(0))])) + ((d_2822_v0_)[default__.safeIndex(146, (d_2822_v0_).length(0))]), (0) - (p0), len((d_2827_v5_) + ((d_2827_v5_).set(default__.safeIndex((d_2822_v0_)[default__.safeIndex(146, (d_2822_v0_).length(0))], len(d_2827_v5_)), d_2831_v9_))), (d_2832_v10_) + (d_2832_v10_))
        d_2834_v12_: _dafny.Set
        d_2834_v12_ = _dafny.Set({d_2823_v1_})
        d_2835_v13_: _dafny.Seq
        d_2835_v13_ = _dafny.SeqWithoutIsStrInference([d_2826_v4_])
        d_2836_v14_: _dafny.Set
        d_2836_v14_ = _dafny.Set({d_2835_v13_, d_2835_v13_})
        d_2833_v11_ = D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nq")), default__.safeDivisionInt((d_2822_v0_)[default__.safeIndex(146, (d_2822_v0_).length(0))], len(d_2834_v12_)), len((d_2827_v5_) + (d_2827_v5_)), default__.safeModuloInt(len(d_2836_v14_), p0), (d_2832_v10_) + (d_2832_v10_))
        hi16_ = p1
        for d_2837_i0_ in range(len(default__.fm45(globalState)), hi16_):
            d_2838_v15_: _dafny.Map
            d_2838_v15_ = _dafny.Map({d_2824_v2_: (0) - (d_2826_v4_)})
            rhs447_ = (d_2828_v6_.f40) not in ((d_2838_v15_) | (d_2838_v15_))
            rhs448_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")) if not (d_2828_v6_.f40) or (d_2824_v2_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxhkbqojr")))
            lhs348_ = globalState
            lhs348_.f1 = rhs447_
            d_2827_v5_ = rhs448_
            if d_2828_v6_.f40:
                (globalState).f21 = p0
                d_2834_v12_ = (d_2834_v12_) | (d_2834_v12_)
                d_2839_v16_: bool
                d_2840_v17_: bool
                d_2841_v18_: _dafny.Map
                d_2842_v19_: int
                out100_: bool
                out101_: bool
                out102_: _dafny.Map
                out103_: int
                out100_, out101_, out102_, out103_ = default__.m0(globalState)
                d_2839_v16_ = out100_
                d_2840_v17_ = out101_
                d_2841_v18_ = out102_
                d_2842_v19_ = out103_
                d_2843_v20_: _dafny.Map
                d_2843_v20_ = _dafny.Map({not(default__.fm0(d_2842_v19_, globalState)): d_2823_v1_})
                d_2844_v21_: D5
                d_2844_v21_ = D5_DC11(d_2824_v2_)
                d_2845_v22_: _dafny.Set
                d_2845_v22_ = _dafny.Set({(_dafny.Map({False: len(d_2843_v20_)})).set(not(d_2828_v6_.f40), len(d_2832_v10_)), (d_2838_v15_) | (d_2838_v15_), d_2838_v15_, ((_dafny.Map({(d_2844_v21_).cf16: d_2837_i0_})).set(d_2824_v2_, (0) - (len(d_2843_v20_)))).set(d_2824_v2_, p1)})
                d_2845_v22_ = (d_2845_v22_) | (d_2845_v22_)
                d_2846_v23_: _dafny.Array
                nw487_ = _dafny.Array(False, 4)
                d_2846_v23_ = nw487_
                d_2847_v24_: _dafny.Set
                d_2847_v24_ = _dafny.Set({d_2846_v23_})
                (globalState).f1 = (_dafny.Set({d_2846_v23_})).ispropersubset(d_2847_v24_)
            elif True:
                d_2848_v25_: D23
                d_2848_v25_ = D23_DC46(d_2823_v1_)
                d_2849_v26_: _dafny.MultiSet
                d_2849_v26_ = _dafny.MultiSet([d_2824_v2_, True])
                d_2850_v27_: _dafny.Map
                d_2850_v27_ = _dafny.Map({d_2828_v6_.f40: False})
                d_2851_v28_: _dafny.Array
                nw488_ = _dafny.Array(None, 21)
                nw488_[int(0)] = (d_2848_v25_).cf63
                nw488_[int(1)] = (d_2832_v10_)[default__.safeIndex((d_2849_v26_).cardinality, len(d_2832_v10_))]
                nw488_[int(2)] = d_2828_v6_.f40
                nw488_[int(3)] = d_2823_v1_
                nw488_[int(4)] = ((d_2850_v27_)[d_2823_v1_] if (d_2823_v1_) in (d_2850_v27_) else d_2828_v6_.f40)
                nw488_[int(5)] = d_2824_v2_
                nw488_[int(6)] = not(not(d_2828_v6_.f40))
                nw488_[int(7)] = not (False) or (d_2828_v6_.f40)
                nw488_[int(8)] = d_2824_v2_
                nw488_[int(9)] = d_2828_v6_.f40
                nw488_[int(10)] = not((d_2827_v5_) <= (_dafny.SeqWithoutIsStrInference([d_2831_v9_ for d_2852_i1_ in range(default__.abs(121))])))
                nw488_[int(11)] = d_2828_v6_.f40
                nw488_[int(12)] = d_2824_v2_
                nw488_[int(13)] = not(((d_2835_v13_)[default__.safeIndex(p0, len(d_2835_v13_))]) in (self.f28))
                nw488_[int(14)] = d_2828_v6_.f40
                nw488_[int(15)] = d_2828_v6_.f40
                nw488_[int(16)] = d_2828_v6_.f40
                nw488_[int(17)] = d_2828_v6_.f40
                nw488_[int(18)] = d_2823_v1_
                nw488_[int(19)] = d_2828_v6_.f40
                nw488_[int(20)] = d_2828_v6_.f40
                d_2851_v28_ = nw488_
                index453_ = default__.safeIndex(522, (d_2851_v28_).length(0))
                (d_2851_v28_)[index453_] = (d_2828_v6_.f40 if d_2828_v6_.f40 else d_2824_v2_)
                index454_ = default__.safeIndex(522, (d_2851_v28_).length(0))
                (d_2851_v28_)[index454_] = default__.fm0(default__.safeModuloInt((d_2822_v0_)[default__.safeIndex(146, (d_2822_v0_).length(0))], (0) - (p1)), globalState)
                r0 = d_2828_v6_.f40
                (self).f28 = (self.f28).set((d_2822_v0_)[default__.safeIndex(146, (d_2822_v0_).length(0))], d_2823_v1_)
                d_2824_v2_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "feihmdkj"))).set(default__.safeIndex(d_2826_v4_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "feihmdkj")))), d_2831_v9_)) == (d_2827_v5_)
                d_2853_v29_: _dafny.Array
                nw489_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_2853_v29_ = nw489_
                d_2854_v30_: C0
                nw490_ = C0()
                nw490_.ctor__((d_2853_v29_ if (d_2851_v28_)[default__.safeIndex(522, (d_2851_v28_).length(0))] else d_2853_v29_))
                d_2854_v30_ = nw490_
            d_2855_v31_: _dafny.MultiSet
            d_2855_v31_ = _dafny.MultiSet([d_2824_v2_])
            d_2855_v31_ = (_dafny.MultiSet([True]) if (d_2828_v6_.f40 if d_2828_v6_.f40 else ((self.f28)[-411] if (-411) in (self.f28) else d_2823_v1_)) else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2824_v2_])))
            (globalState).f2 = default__.fm2(globalState)
        (globalState).f2 = (d_2822_v0_)[default__.safeIndex(146, (d_2822_v0_).length(0))]
        (d_2828_v6_).f40 = not ((d_2823_v1_ if d_2824_v2_ else d_2824_v2_)) or (d_2828_v6_.f40)
        r0 = d_2824_v2_
        r1 = d_2831_v9_
        return r0, r1


class C13(T0):
    def  __init__(self):
        self._f26: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C13"
    @property
    def f26(self):
        return self._f26
    @f26.setter
    def f26(self, value):
        self._f26 = value
    def ctor__(self, f26):
        (self).f26 = f26

    def fm3(self, globalState):
        return (len((_dafny.Set({not(not(False))})).intersection(_dafny.Set({True, False})))) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))))

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Map({943: _dafny.Map({_dafny.CodePoint('h'): -677})})) | (_dafny.Map({-565: _dafny.Map({_dafny.CodePoint('i'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th")))})}))) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rse"))): _dafny.Map({_dafny.CodePoint('i'): 266})})) | (_dafny.Map({370: _dafny.Map({_dafny.CodePoint('c'): -749})})))

    def fm5(self, p0, p1, globalState):
        return ((_dafny.MultiSet([883, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jayf")))]) if False else _dafny.MultiSet([(_dafny.MultiSet([len(_dafny.Map({not(True): True})), len(_dafny.Set({-569, (_dafny.MultiSet([True])).cardinality, len(self.f26)})), 818, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2856_i0_ in range(default__.abs(338))])), 542])).cardinality]))).isdisjoint((_dafny.MultiSet([(D3_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsgmlrg")), len(_dafny.Map({not(not(False)): -112})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsjq"))), len(_dafny.Set({True})), _dafny.SeqWithoutIsStrInference([not(False)]))).cf11, 146])) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cblohdo")))])))

    def fm6(self, p0, globalState):
        return (self.f26) | (self.f26)

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        d_2857_v0_: _dafny.Set
        d_2857_v0_ = _dafny.Set({default__.safeModuloInt(p2, p1)})
        d_2857_v0_ = d_2857_v0_
        (globalState).f16 = default__.safeDivisionInt(p2, (self).fm3(globalState))
        d_2858_v1_: _dafny.Array
        nw491_ = _dafny.Array(_dafny.Array(None, 0), 9)
        d_2858_v1_ = nw491_
        d_2859_v2_: _dafny.Array
        def lambda138_(d_2860_i0_):
            return _dafny.Set({True, True})

        init80_ = lambda138_
        nw492_ = _dafny.Array(None, 18)
        for i0_80_ in range(nw492_.length(0)):
            nw492_[i0_80_] = init80_(i0_80_)
        d_2859_v2_ = nw492_
        index455_ = default__.safeIndex(839, (d_2858_v1_).length(0))
        (d_2858_v1_)[index455_] = d_2859_v2_
        index456_ = default__.safeIndex(839, (d_2858_v1_).length(0))
        (d_2858_v1_)[index456_] = d_2859_v2_
        d_2861_v3_: _dafny.Array
        def lambda139_(d_2862_i1_):
            return not(False)

        init81_ = lambda139_
        nw493_ = _dafny.Array(None, 12)
        for i0_81_ in range(nw493_.length(0)):
            nw493_[i0_81_] = init81_(i0_81_)
        d_2861_v3_ = nw493_
        d_2863_v4_: bool
        d_2863_v4_ = False
        index457_ = default__.safeIndex(130, (d_2861_v3_).length(0))
        (d_2861_v3_)[index457_] = d_2863_v4_
        index458_ = default__.safeIndex(130, (d_2861_v3_).length(0))
        (d_2861_v3_)[index458_] = False
        d_2863_v4_ = False
        if (p2) > (p2):
            d_2864_v5_: _dafny.Seq
            d_2864_v5_ = _dafny.SeqWithoutIsStrInference([(d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]])
            d_2865_v8_: _dafny.Seq
            d_2865_v8_ = _dafny.SeqWithoutIsStrInference([d_2864_v5_])
            d_2866_v9_: D3
            def iife286_():
                coll110_ = _dafny.Map()
                compr_110_: _dafny.Seq
                for compr_110_ in (d_2865_v8_).Elements:
                    d_2867_v7_: _dafny.Seq = compr_110_
                    if (d_2867_v7_) in (d_2865_v8_):
                        coll110_[d_2867_v7_] = (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]
                return _dafny.Map(coll110_)
            d_2866_v9_ = D3_DC6(iife286_()
)
            d_2868_v10_: _dafny.MultiSet
            d_2868_v10_ = _dafny.MultiSet([d_2863_v4_])
            d_2869_v11_: _dafny.Array
            def lambda140_(d_2870_p2_):
                def lambda141_(d_2871_i2_):
                    return (d_2871_i2_) - (d_2870_p2_)

                return lambda141_

            init82_ = lambda140_(p2)
            nw494_ = _dafny.Array(None, 9)
            for i0_82_ in range(nw494_.length(0)):
                nw494_[i0_82_] = init82_(i0_82_)
            d_2869_v11_ = nw494_
            d_2872_v12_: str
            d_2872_v12_ = _dafny.CodePoint('x')
            d_2873_v13_: D0
            d_2873_v13_ = D0_DC1(d_2869_v11_, d_2872_v12_, p1)
            d_2874_v14_: _dafny.Seq
            d_2874_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkjv"))
            d_2875_v15_: _dafny.Map
            d_2875_v15_ = _dafny.Map({d_2873_v13_: d_2874_v14_})
            d_2876_v16_: T0
            nw495_ = C9()
            nw495_.ctor__(d_2861_v3_, p2, self.f26)
            d_2876_v16_ = nw495_
            d_2877_v17_: _dafny.MultiSet
            d_2877_v17_ = _dafny.MultiSet([len(d_2874_v14_)])
            d_2878_v18_: _dafny.Map
            d_2878_v18_ = _dafny.Map({d_2876_v16_: d_2877_v17_})
            d_2879_v19_: C8
            nw496_ = C8()
            nw496_.ctor__(d_2863_v4_, D5_DC11((d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]), d_2876_v16_.f26)
            d_2879_v19_ = nw496_
            d_2880_v20_: _dafny.Seq
            d_2880_v20_ = _dafny.SeqWithoutIsStrInference([d_2879_v19_])
            d_2881_v21_: _dafny.Map
            d_2881_v21_ = _dafny.Map({d_2863_v4_: p1})
            d_2882_v22_: T1
            nw497_ = C5()
            nw497_.ctor__(p2)
            d_2882_v22_ = nw497_
            d_2883_v23_: _dafny.Set
            d_2883_v23_ = _dafny.Set({d_2882_v22_, d_2882_v22_})
            d_2884_v24_: _dafny.Array
            nw498_ = _dafny.Array(None, 27)
            nw498_[int(0)] = p1
            nw498_[int(1)] = p2
            nw498_[int(2)] = p2
            nw498_[int(3)] = len(d_2864_v5_)
            nw498_[int(4)] = p1
            nw498_[int(5)] = p1
            nw498_[int(6)] = p1
            nw498_[int(7)] = p2
            def iife287_():
                coll111_ = _dafny.Map()
                compr_111_: int
                for compr_111_ in _dafny.IntegerRange(191, 302):
                    d_2885_v6_: int = compr_111_
                    if ((191) <= (d_2885_v6_)) and ((d_2885_v6_) < (302)):
                        coll111_[(d_2885_v6_) * (p2)] = d_2863_v4_
                return _dafny.Map(coll111_)
            nw498_[int(8)] = (len(iife287_()
            )) * (p1)
            nw498_[int(9)] = (p2) - (p2)
            nw498_[int(10)] = len((d_2857_v0_).intersection(default__.fm7(len(self.f26), d_2866_v9_, d_2863_v4_, globalState)))
            nw498_[int(11)] = p1
            nw498_[int(12)] = default__.safeModuloInt(p1, ((d_2868_v10_)[d_2863_v4_] if (d_2863_v4_) in (d_2868_v10_) else p2))
            nw498_[int(13)] = default__.safeDivisionInt(len(d_2875_v15_), len(d_2878_v18_))
            nw498_[int(14)] = default__.safeModuloInt(p1, p1)
            nw498_[int(15)] = -285
            nw498_[int(16)] = default__.safeModuloInt(len((d_2880_v20_).set(default__.safeIndex(((d_2881_v21_)[d_2879_v19_.f35] if (d_2879_v19_.f35) in (d_2881_v21_) else p2), len(d_2880_v20_)), d_2879_v19_)), p1)
            nw498_[int(17)] = p1
            nw498_[int(18)] = p1
            nw498_[int(19)] = (0) - (p1)
            nw498_[int(20)] = (p2) - (len(d_2874_v14_))
            nw498_[int(21)] = p1
            nw498_[int(22)] = len(d_2883_v23_)
            nw498_[int(23)] = (p1 if not(d_2879_v19_.f35) else (0) - (p1))
            nw498_[int(24)] = p1
            nw498_[int(25)] = p1
            nw498_[int(26)] = p2
            d_2884_v24_ = nw498_
            index459_ = default__.safeIndex(966, (d_2869_v11_).length(0))
            (d_2869_v11_)[index459_] = p2
            d_2886_v26_: _dafny.Seq
            def iife288_():
                coll112_ = _dafny.Map()
                compr_112_: int
                for compr_112_ in (d_2857_v0_).Elements:
                    d_2887_v25_: int = compr_112_
                    if (d_2887_v25_) in (d_2857_v0_):
                        coll112_[(d_2887_v25_) + (p1)] = (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]
                return _dafny.Map(coll112_)
            d_2886_v26_ = _dafny.SeqWithoutIsStrInference([iife288_()
            ])
            index460_ = default__.safeIndex(966, (d_2869_v11_).length(0))
            (d_2869_v11_)[index460_] = default__.safeDivisionInt(((d_2876_v16_.f26)[p1] if (p1) in (d_2876_v16_.f26) else p2), (0) - (len(d_2886_v26_)))
            (globalState).f1 = d_2879_v19_.f35
            d_2888_v27_: _dafny.Array
            nw499_ = _dafny.Array(_dafny.Seq({}), 25)
            d_2888_v27_ = nw499_
            index461_ = default__.safeIndex(259, (d_2888_v27_).length(0))
            (d_2888_v27_)[index461_] = d_2864_v5_
            d_2889_v28_: _dafny.Map
            d_2889_v28_ = _dafny.Map({(d_2869_v11_)[default__.safeIndex(966, (d_2869_v11_).length(0))]: d_2863_v4_})
            index462_ = default__.safeIndex(259, (d_2888_v27_).length(0))
            index463_ = default__.safeIndex(130, (d_2861_v3_).length(0))
            rhs449_ = d_2863_v4_
            rhs450_ = default__.fm13(default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference([(d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]])).set(default__.safeIndex((d_2869_v11_)[default__.safeIndex(966, (d_2869_v11_).length(0))], len(_dafny.SeqWithoutIsStrInference([(d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]]))), d_2863_v4_)), len(d_2889_v28_)), globalState)
            rhs451_ = (_dafny.MultiSet([d_2872_v12_, _dafny.CodePoint('e'), (d_2874_v14_)[default__.safeIndex((d_2869_v11_)[default__.safeIndex(966, (d_2869_v11_).length(0))], len(d_2874_v14_))], d_2872_v12_])).ispropersubset(default__.fm63(globalState))
            lhs349_ = globalState
            lhs350_ = d_2888_v27_
            lhs351_ = default__.safeIndex(259, (d_2888_v27_).length(0))
            lhs352_ = d_2861_v3_
            lhs353_ = default__.safeIndex(130, (d_2861_v3_).length(0))
            lhs349_.f1 = rhs449_
            lhs350_[lhs351_] = rhs450_
            lhs352_[lhs353_] = rhs451_
            index464_ = default__.safeIndex(259, (d_2888_v27_).length(0))
            (d_2888_v27_)[index464_] = (d_2888_v27_)[default__.safeIndex(259, (d_2888_v27_).length(0))]
            d_2890_v29_: int
            d_2891_v30_: _dafny.Array
            d_2892_v31_: bool
            out104_: int
            out105_: _dafny.Array
            out106_: bool
            out104_, out105_, out106_ = (self).m4((d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))], _dafny.Set({False, d_2879_v19_.f35}), d_2874_v14_, d_2881_v21_, globalState)
            d_2890_v29_ = out104_
            d_2891_v30_ = out105_
            d_2892_v31_ = out106_
        elif True:
            d_2893_v32_: _dafny.Seq
            d_2893_v32_ = _dafny.SeqWithoutIsStrInference([d_2863_v4_, (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]])
            if not (d_2863_v4_) or ((d_2893_v32_) != (_dafny.SeqWithoutIsStrInference([True, d_2863_v4_]))):
                d_2894_v33_: _dafny.Array
                def lambda142_(d_2895_i3_):
                    return default__.safeModuloInt(d_2895_i3_, -413)

                init83_ = lambda142_
                nw500_ = _dafny.Array(None, 10)
                for i0_83_ in range(nw500_.length(0)):
                    nw500_[i0_83_] = init83_(i0_83_)
                d_2894_v33_ = nw500_
                index465_ = default__.safeIndex(232, (d_2894_v33_).length(0))
                (d_2894_v33_)[index465_] = p2
                index466_ = default__.safeIndex(232, (d_2894_v33_).length(0))
                (d_2894_v33_)[index466_] = p2
                d_2896_v34_: _dafny.Map
                d_2896_v34_ = _dafny.Map({not(not(not(not((d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))])))): (d_2894_v33_)[default__.safeIndex(232, (d_2894_v33_).length(0))]})
                d_2897_v35_: C1
                nw501_ = C1()
                nw501_.ctor__(d_2896_v34_)
                d_2897_v35_ = nw501_
                (globalState).f5 = (d_2894_v33_)[default__.safeIndex(232, (d_2894_v33_).length(0))]
                d_2898_v36_: _dafny.Array
                nw502_ = _dafny.Array(_dafny.Seq({}), 18)
                d_2898_v36_ = nw502_
                d_2899_v37_: _dafny.Seq
                d_2899_v37_ = _dafny.SeqWithoutIsStrInference([(d_2894_v33_)[default__.safeIndex(232, (d_2894_v33_).length(0))]])
                index467_ = default__.safeIndex(774, (d_2898_v36_).length(0))
                (d_2898_v36_)[index467_] = d_2899_v37_
                d_2900_v38_: _dafny.Seq
                d_2900_v38_ = (((d_2899_v37_)).set(default__.safeIndex((d_2894_v33_)[default__.safeIndex(232, (d_2894_v33_).length(0))], len((d_2899_v37_))), p1)) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpoimf"))) for d_2901_i4_ in range(default__.abs(630))]))
                index468_ = default__.safeIndex(774, (d_2898_v36_).length(0))
                (d_2898_v36_)[index468_] = d_2900_v38_
                d_2902_v39_: _dafny.MultiSet
                d_2902_v39_ = _dafny.MultiSet([p1])
                index469_ = default__.safeIndex(232, (d_2894_v33_).length(0))
                rhs452_ = ((d_2902_v39_) - ((d_2902_v39_).set(p2, default__.abs(len(_dafny.Map({(d_2894_v33_)[default__.safeIndex(232, (d_2894_v33_).length(0))]: (d_2894_v33_)[default__.safeIndex(232, (d_2894_v33_).length(0))]})))))).cardinality
                rhs453_ = (d_2863_v4_ if (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))] else d_2863_v4_)
                rhs454_ = (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]
                lhs354_ = d_2894_v33_
                lhs355_ = default__.safeIndex(232, (d_2894_v33_).length(0))
                lhs356_ = globalState
                lhs357_ = globalState
                lhs354_[lhs355_] = rhs452_
                lhs356_.f13 = rhs453_
                lhs357_.f13 = rhs454_
            elif True:
                d_2903_v40_: _dafny.Array
                def lambda143_(d_2904_i5_):
                    return _dafny.CodePoint('a')

                init84_ = lambda143_
                nw503_ = _dafny.Array(None, 13)
                for i0_84_ in range(nw503_.length(0)):
                    nw503_[i0_84_] = init84_(i0_84_)
                d_2903_v40_ = nw503_
                d_2905_v41_: C0
                nw504_ = C0()
                nw504_.ctor__(d_2903_v40_)
                d_2905_v41_ = nw504_
                d_2906_v42_: C6
                nw505_ = C6()
                nw505_.ctor__()
                d_2906_v42_ = nw505_
                d_2907_v43_: D7
                d_2907_v43_ = D7_DC16(p2, d_2863_v4_)
                d_2908_v44_: _dafny.Array
                def lambda144_(d_2909_p2_):
                    def lambda145_(d_2910_i6_):
                        return (d_2910_i6_) - (d_2909_p2_)

                    return lambda145_

                init85_ = lambda144_(p2)
                nw506_ = _dafny.Array(None, 24)
                for i0_85_ in range(nw506_.length(0)):
                    nw506_[i0_85_] = init85_(i0_85_)
                d_2908_v44_ = nw506_
                d_2911_v45_: _dafny.Map
                d_2911_v45_ = _dafny.Map({(d_2907_v43_).cf22: (d_2908_v44_ if d_2863_v4_ else d_2908_v44_)})
                d_2911_v45_ = d_2911_v45_
                (globalState).f1 = True
                d_2912_v46_: _dafny.Seq
                d_2912_v46_ = _dafny.SeqWithoutIsStrInference([p2])
                d_2913_v47_: _dafny.MultiSet
                d_2913_v47_ = _dafny.MultiSet([len(d_2912_v46_), p2])
                index470_ = default__.safeIndex(130, (d_2861_v3_).length(0))
                index471_ = default__.safeIndex(130, (d_2861_v3_).length(0))
                rhs455_ = (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]
                rhs456_ = not((d_2913_v47_).isdisjoint(d_2913_v47_))
                rhs457_ = d_2863_v4_
                rhs458_ = (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]
                rhs459_ = (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]
                lhs358_ = d_2861_v3_
                lhs359_ = default__.safeIndex(130, (d_2861_v3_).length(0))
                lhs360_ = globalState
                lhs361_ = d_2861_v3_
                lhs362_ = default__.safeIndex(130, (d_2861_v3_).length(0))
                lhs363_ = globalState
                lhs364_ = globalState
                lhs358_[lhs359_] = rhs455_
                lhs360_.f1 = rhs456_
                lhs361_[lhs362_] = rhs457_
                lhs363_.f13 = rhs458_
                lhs364_.f13 = rhs459_
            d_2914_v48_: _dafny.Seq
            d_2914_v48_ = _dafny.SeqWithoutIsStrInference([p2, p2, len(d_2893_v32_)])
            d_2915_v49_: _dafny.Seq
            d_2915_v49_ = _dafny.SeqWithoutIsStrInference([d_2914_v48_])
            d_2916_v50_: D5
            d_2916_v50_ = D5_DC11(d_2863_v4_)
            d_2917_v51_: C8
            nw507_ = C8()
            nw507_.ctor__(d_2863_v4_, d_2916_v50_, self.f26)
            d_2917_v51_ = nw507_
            d_2918_v52_: _dafny.Map
            d_2918_v52_ = _dafny.Map({d_2863_v4_: d_2917_v51_})
            d_2919_v53_: _dafny.MultiSet
            d_2919_v53_ = _dafny.MultiSet([p1, p2, 698])
            d_2920_v54_: _dafny.MultiSet
            d_2920_v54_ = _dafny.MultiSet([(d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]])
            d_2921_v55_: _dafny.Array
            nw508_ = _dafny.Array(None, 16)
            nw508_[int(0)] = d_2914_v48_
            nw508_[int(1)] = _dafny.SeqWithoutIsStrInference([p1 for d_2922_i7_ in range(default__.abs(-730))])
            nw508_[int(2)] = ((d_2915_v49_)[default__.safeIndex(p2, len(d_2915_v49_))]).set(default__.safeIndex(p1, len((d_2915_v49_)[default__.safeIndex(p2, len(d_2915_v49_))])), p2)
            nw508_[int(3)] = (d_2914_v48_) + (d_2914_v48_)
            nw508_[int(4)] = d_2914_v48_
            nw508_[int(5)] = (d_2914_v48_) + (d_2914_v48_)
            nw508_[int(6)] = (_dafny.SeqWithoutIsStrInference([p2, p2, p2, len(d_2914_v48_)])) + (d_2914_v48_)
            nw508_[int(7)] = (d_2914_v48_).set(default__.safeIndex(len(d_2918_v52_), len(d_2914_v48_)), (d_2919_v53_).cardinality)
            nw508_[int(8)] = d_2914_v48_
            nw508_[int(9)] = _dafny.SeqWithoutIsStrInference([p2, p1])
            nw508_[int(10)] = (d_2914_v48_) + (d_2914_v48_)
            nw508_[int(11)] = default__.fm29(d_2920_v54_, p2, d_2863_v4_, p2, globalState)
            nw508_[int(12)] = (d_2914_v48_) + (_dafny.SeqWithoutIsStrInference([p2 for d_2923_i8_ in range(default__.abs(181))]))
            nw508_[int(13)] = (_dafny.SeqWithoutIsStrInference([p1, p1])) + (d_2914_v48_)
            nw508_[int(14)] = (d_2914_v48_) + (d_2914_v48_)
            nw508_[int(15)] = (_dafny.SeqWithoutIsStrInference([p2 for d_2924_i9_ in range(default__.abs(174))])) + (d_2914_v48_)
            d_2921_v55_ = nw508_
            d_2921_v55_ = d_2921_v55_
            d_2925_v56_: _dafny.Seq
            d_2925_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkfooslw"))
            if not(not ((self).fm5((d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))], p1, globalState)) or ((d_2925_v56_) < (d_2925_v56_))):
                d_2926_v57_: _dafny.Array
                def lambda146_(d_2927_v56_, d_2928_p2_, d_2929_p1_, d_2930_v32_):
                    def lambda147_(d_2931_i10_):
                        return (d_2931_i10_) - (len((D3_DC7(d_2927_v56_, d_2928_p2_, d_2928_p2_, d_2929_p1_, d_2930_v32_)).cf8))

                    return lambda147_

                init86_ = lambda146_(d_2925_v56_, p2, p1, d_2893_v32_)
                nw509_ = _dafny.Array(None, 18)
                for i0_86_ in range(nw509_.length(0)):
                    nw509_[i0_86_] = init86_(i0_86_)
                d_2926_v57_ = nw509_
                index472_ = default__.safeIndex(620, (d_2926_v57_).length(0))
                (d_2926_v57_)[index472_] = (p2) - (898)
                index473_ = default__.safeIndex(620, (d_2926_v57_).length(0))
                (d_2926_v57_)[index473_] = (0) - (p2)
                d_2932_v58_: str
                d_2932_v58_ = _dafny.CodePoint('s')
                d_2932_v58_ = _dafny.CodePoint('f')
                d_2933_v59_: bool
                d_2934_v60_: bool
                d_2935_v61_: _dafny.Map
                d_2936_v62_: int
                out107_: bool
                out108_: bool
                out109_: _dafny.Map
                out110_: int
                out107_, out108_, out109_, out110_ = default__.m0(globalState)
                d_2933_v59_ = out107_
                d_2934_v60_ = out108_
                d_2935_v61_ = out109_
                d_2936_v62_ = out110_
                (globalState).f16 = p2
                (globalState).f1 = d_2933_v59_
            elif True:
                (d_2917_v51_).f35 = (d_2857_v0_).issubset(d_2857_v0_)
                d_2937_v63_: str
                d_2937_v63_ = _dafny.CodePoint('c')
                d_2938_v64_: _dafny.Set
                d_2938_v64_ = _dafny.Set({d_2863_v4_})
                d_2939_v65_: D10
                d_2939_v65_ = D10_DC24((d_2938_v64_) - (_dafny.Set({d_2863_v4_})))
                d_2940_v66_: D20
                d_2940_v66_ = D20_DC41(596)
                rhs460_ = d_2937_v63_
                rhs461_ = (p2) * ((172 if default__.fm0(p2, globalState) else (d_2940_v66_).cf57))
                rhs462_ = D10_DC24(d_2938_v64_)
                lhs365_ = globalState
                d_2937_v63_ = rhs460_
                lhs365_.f21 = rhs461_
                d_2939_v65_ = rhs462_
                d_2941_v67_: C11
                nw510_ = C11()
                nw510_.ctor__(p2, self.f26)
                d_2941_v67_ = nw510_
                d_2942_v68_: _dafny.Map
                d_2942_v68_ = _dafny.Map({d_2863_v4_: (0) - (p1)})
                d_2943_v69_: _dafny.Map
                d_2943_v69_ = _dafny.Map({d_2863_v4_: (d_2893_v32_)[default__.safeIndex(p1, len(d_2893_v32_))]})
                d_2942_v68_ = (d_2942_v68_).set(((d_2943_v69_)[(d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]] if ((d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]) in (d_2943_v69_) else (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]), (d_2941_v67_).f30)
                (self).f26 = (self.f26).set(p2, (d_2941_v67_).f30)
            d_2944_v70_: str
            d_2944_v70_ = _dafny.CodePoint('l')
            d_2945_v71_: _dafny.Array
            nw511_ = _dafny.Array(None, 16)
            nw511_[int(0)] = _dafny.CodePoint('b')
            nw511_[int(1)] = d_2944_v70_
            nw511_[int(2)] = d_2944_v70_
            nw511_[int(3)] = d_2944_v70_
            nw511_[int(4)] = d_2944_v70_
            nw511_[int(5)] = d_2944_v70_
            nw511_[int(6)] = d_2944_v70_
            nw511_[int(7)] = d_2944_v70_
            nw511_[int(8)] = d_2944_v70_
            nw511_[int(9)] = _dafny.CodePoint('l')
            nw511_[int(10)] = d_2944_v70_
            nw511_[int(11)] = d_2944_v70_
            nw511_[int(12)] = _dafny.CodePoint('j')
            nw511_[int(13)] = d_2944_v70_
            nw511_[int(14)] = d_2944_v70_
            nw511_[int(15)] = d_2944_v70_
            d_2945_v71_ = nw511_
            d_2946_v72_: _dafny.Map
            d_2946_v72_ = _dafny.Map({D0_DC0(d_2945_v71_): (d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]})
            d_2947_v73_: _dafny.Map
            d_2947_v73_ = _dafny.Map({d_2946_v72_: d_2893_v32_})
            d_2947_v73_ = (d_2947_v73_).set((d_2946_v72_) | (d_2946_v72_), d_2893_v32_)
            d_2948_v75_: _dafny.MultiSet
            def iife289_():
                coll113_ = _dafny.Set()
                compr_113_: int
                for compr_113_ in _dafny.IntegerRange(825, 494):
                    d_2949_v74_: int = compr_113_
                    if ((825) <= (d_2949_v74_)) and ((d_2949_v74_) < (494)):
                        coll113_ = coll113_.union(_dafny.Set([(d_2949_v74_) + (p2)]))
                return _dafny.Set(coll113_)
            d_2948_v75_ = _dafny.MultiSet([d_2857_v0_, (d_2857_v0_) - (iife289_()
            )])
            d_2948_v75_ = d_2948_v75_
        r0 = (p2) + (p2)
        d_2950_v76_: _dafny.Seq
        d_2950_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijmbxg"))
        d_2951_v77_: _dafny.Seq
        d_2951_v77_ = _dafny.SeqWithoutIsStrInference([d_2950_v76_])
        d_2952_v78_: _dafny.Map
        d_2952_v78_ = _dafny.Map({(d_2861_v3_)[default__.safeIndex(130, (d_2861_v3_).length(0))]: (d_2951_v77_)[default__.safeIndex(p2, len(d_2951_v77_))]})
        d_2953_v79_: _dafny.Map
        d_2953_v79_ = _dafny.Map({((d_2952_v78_)[d_2863_v4_] if (d_2863_v4_) in (d_2952_v78_) else d_2950_v76_): d_2861_v3_})
        r1 = (d_2953_v79_ if default__.fm0(p2, globalState) else d_2953_v79_)
        r2 = p1
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        d_2954_v0_: bool
        d_2954_v0_ = False
        r0 = d_2954_v0_
        d_2955_v1_: _dafny.Set
        d_2955_v1_ = _dafny.Set({d_2954_v0_, d_2954_v0_})
        if ((d_2955_v1_).intersection(d_2955_v1_)).ispropersubset(d_2955_v1_):
            d_2956_v2_: int
            d_2956_v2_ = 964
            (globalState).f2 = default__.safeDivisionInt(d_2956_v2_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_2956_v2_: d_2954_v0_})) for d_2957_i0_ in range(default__.abs(57))])))
            d_2958_v3_: _dafny.Array
            def lambda148_(d_2959_v0_):
                def lambda149_(d_2960_i1_):
                    return d_2959_v0_

                return lambda149_

            init87_ = lambda148_(d_2954_v0_)
            nw512_ = _dafny.Array(None, 22)
            for i0_87_ in range(nw512_.length(0)):
                nw512_[i0_87_] = init87_(i0_87_)
            d_2958_v3_ = nw512_
            index474_ = default__.safeIndex(265, (d_2958_v3_).length(0))
            (d_2958_v3_)[index474_] = d_2954_v0_
            index475_ = default__.safeIndex(265, (d_2958_v3_).length(0))
            (d_2958_v3_)[index475_] = False
            d_2961_v4_: C10
            nw513_ = C10()
            nw513_.ctor__(d_2956_v2_, (d_2958_v3_)[default__.safeIndex(265, (d_2958_v3_).length(0))])
            d_2961_v4_ = nw513_
            d_2962_v5_: _dafny.Map
            d_2962_v5_ = _dafny.Map({(d_2958_v3_)[default__.safeIndex(265, (d_2958_v3_).length(0))]: d_2961_v4_})
            d_2962_v5_ = (d_2962_v5_).set(not ((d_2958_v3_)[default__.safeIndex(265, (d_2958_v3_).length(0))]) or (d_2954_v0_), d_2961_v4_)
            rhs463_ = not(default__.fm0(d_2956_v2_, globalState))
            rhs464_ = d_2954_v0_
            lhs366_ = globalState
            lhs367_ = globalState
            lhs366_.f1 = rhs463_
            lhs367_.f13 = rhs464_
            (globalState).f5 = (d_2956_v2_) * ((d_2961_v4_).f31)
        elif True:
            d_2954_v0_ = d_2954_v0_
            (globalState).f13 = d_2954_v0_
            d_2963_v6_: D12
            d_2963_v6_ = D12_DC29()
            d_2964_v7_: _dafny.MultiSet
            d_2964_v7_ = _dafny.MultiSet([d_2963_v6_, d_2963_v6_, d_2963_v6_, d_2963_v6_, d_2963_v6_])
            (globalState).f13 = (d_2964_v7_) == (d_2964_v7_)
            d_2965_v8_: int
            d_2965_v8_ = -968
            (globalState).f5 = default__.safeModuloInt(d_2965_v8_, d_2965_v8_)
            d_2966_v9_: _dafny.Seq
            d_2966_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bok"))
            d_2967_v10_: _dafny.Array
            def lambda150_(d_2968_v8_):
                def lambda151_(d_2969_i2_):
                    return (d_2969_i2_) + (d_2968_v8_)

                return lambda151_

            init88_ = lambda150_(d_2965_v8_)
            nw514_ = _dafny.Array(None, 4)
            for i0_88_ in range(nw514_.length(0)):
                nw514_[i0_88_] = init88_(i0_88_)
            d_2967_v10_ = nw514_
            d_2970_v11_: _dafny.Map
            d_2970_v11_ = _dafny.Map({d_2954_v0_: d_2967_v10_})
            d_2971_v12_: _dafny.Map
            d_2971_v12_ = _dafny.Map({d_2954_v0_: d_2954_v0_})
            d_2972_v13_: str
            d_2972_v13_ = _dafny.CodePoint('f')
            d_2973_v14_: _dafny.MultiSet
            d_2973_v14_ = _dafny.MultiSet([d_2972_v13_, d_2972_v13_])
            d_2974_v15_: _dafny.Seq
            d_2974_v15_ = _dafny.SeqWithoutIsStrInference([d_2954_v0_, False])
            d_2975_v16_: _dafny.Seq
            d_2975_v16_ = _dafny.SeqWithoutIsStrInference([len(d_2974_v15_), d_2965_v8_])
            d_2976_v17_: _dafny.Array
            nw515_ = _dafny.Array(None, 25)
            nw515_[int(0)] = d_2954_v0_
            nw515_[int(1)] = (_dafny.Set({d_2954_v0_})) != (_dafny.Set({d_2954_v0_}))
            nw515_[int(2)] = (len(d_2966_v9_)) <= (789)
            nw515_[int(3)] = d_2954_v0_
            nw515_[int(4)] = d_2954_v0_
            nw515_[int(5)] = d_2954_v0_
            nw515_[int(6)] = (d_2955_v1_).ispropersubset(d_2955_v1_)
            nw515_[int(7)] = d_2954_v0_
            nw515_[int(8)] = (d_2970_v11_) == (_dafny.Map({d_2954_v0_: d_2967_v10_}))
            nw515_[int(9)] = (d_2965_v8_) < (len(d_2955_v1_))
            nw515_[int(10)] = d_2954_v0_
            nw515_[int(11)] = ((d_2971_v12_)[d_2954_v0_] if (d_2954_v0_) in (d_2971_v12_) else d_2954_v0_)
            nw515_[int(12)] = d_2954_v0_
            nw515_[int(13)] = d_2954_v0_
            nw515_[int(14)] = d_2954_v0_
            nw515_[int(15)] = not(d_2954_v0_)
            nw515_[int(16)] = d_2954_v0_
            nw515_[int(17)] = (d_2973_v14_).isdisjoint(d_2973_v14_)
            nw515_[int(18)] = d_2954_v0_
            nw515_[int(19)] = (len((d_2975_v16_).set(default__.safeIndex(d_2965_v8_, len(d_2975_v16_)), d_2965_v8_))) > (d_2965_v8_)
            nw515_[int(20)] = True
            nw515_[int(21)] = d_2954_v0_
            nw515_[int(22)] = d_2954_v0_
            nw515_[int(23)] = d_2954_v0_
            nw515_[int(24)] = d_2954_v0_
            d_2976_v17_ = nw515_
            nw516_ = _dafny.Array(False, 4)
            d_2976_v17_ = nw516_
        d_2977_v18_: int
        d_2977_v18_ = 461
        d_2978_v19_: str
        d_2978_v19_ = _dafny.CodePoint('a')
        d_2979_v20_: _dafny.Seq
        d_2979_v20_ = _dafny.SeqWithoutIsStrInference([d_2978_v19_])
        d_2980_v21_: _dafny.Map
        d_2980_v21_ = _dafny.Map({d_2977_v18_: (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2981_i3_ in range(default__.abs(297))])) + (d_2979_v20_)})
        d_2980_v21_ = (d_2980_v21_).set(default__.safeModuloInt(len(d_2979_v20_), d_2977_v18_), d_2979_v20_)
        d_2982_v22_: _dafny.Seq
        d_2982_v22_ = _dafny.SeqWithoutIsStrInference([d_2977_v18_])
        d_2983_v23_: _dafny.Seq
        d_2983_v23_ = (d_2982_v22_) + (_dafny.SeqWithoutIsStrInference([106, d_2977_v18_]))
        source37_ = d_2983_v23_
        d_2984___mcc_h0_ = source37_
        d_2985_cf20_ = d_2984___mcc_h0_
        (globalState).f16 = (0) - (d_2977_v18_)
        d_2986_v24_: _dafny.Array
        nw517_ = _dafny.Array(int(0), 20)
        d_2986_v24_ = nw517_
        index476_ = default__.safeIndex(219, (d_2986_v24_).length(0))
        (d_2986_v24_)[index476_] = d_2977_v18_
        d_2987_v25_: _dafny.Seq
        d_2987_v25_ = _dafny.SeqWithoutIsStrInference([d_2954_v0_])
        d_2988_v26_: _dafny.MultiSet
        d_2988_v26_ = _dafny.MultiSet([d_2987_v25_])
        index477_ = default__.safeIndex(219, (d_2986_v24_).length(0))
        (d_2986_v24_)[index477_] = (0) - (((d_2988_v26_).intersection(d_2988_v26_)).cardinality)
        d_2989_v27_: _dafny.Map
        d_2989_v27_ = _dafny.Map({d_2977_v18_: d_2954_v0_})
        d_2990_v28_: T0
        nw518_ = C11()
        nw518_.ctor__(len(d_2987_v25_), self.f26)
        d_2990_v28_ = nw518_
        d_2991_v29_: _dafny.Map
        d_2991_v29_ = _dafny.Map({d_2990_v28_: d_2977_v18_})
        d_2992_v30_: _dafny.Map
        d_2992_v30_ = _dafny.Map({False: ((d_2991_v29_)[d_2990_v28_] if (d_2990_v28_) in (d_2991_v29_) else (d_2986_v24_)[default__.safeIndex(219, (d_2986_v24_).length(0))])})
        d_2993_v31_: _dafny.Map
        d_2993_v31_ = _dafny.Map({d_2989_v27_: d_2992_v30_})
        d_2994_v32_: D9
        d_2994_v32_ = D9_DC22(d_2993_v31_)
        d_2995_v33_: _dafny.MultiSet
        d_2995_v33_ = _dafny.MultiSet([d_2994_v32_])
        (globalState).f13 = not((d_2995_v33_).issubset(d_2995_v33_))
        d_2996_v34_: D7
        d_2996_v34_ = D7_DC17(not(d_2954_v0_), d_2954_v0_, (d_2986_v24_)[default__.safeIndex(219, (d_2986_v24_).length(0))])
        (globalState).f21 = (((d_2982_v22_)[default__.safeIndex((d_2996_v34_).cf26, len(d_2982_v22_))]) - ((d_2986_v24_)[default__.safeIndex(219, (d_2986_v24_).length(0))])) * (d_2977_v18_)
        d_2997_i4_: int
        d_2997_i4_ = 0
        with _dafny.label("23"):
            while (d_2954_v0_) or (d_2954_v0_):
                with _dafny.c_label("23"):
                    if (d_2997_i4_) >= (100):
                        raise _dafny.Break("23")
                    d_2997_i4_ = (d_2997_i4_) + (1)
                    d_2998_v35_: bool
                    d_2999_v36_: bool
                    d_3000_v37_: _dafny.Map
                    d_3001_v38_: int
                    out111_: bool
                    out112_: bool
                    out113_: _dafny.Map
                    out114_: int
                    out111_, out112_, out113_, out114_ = default__.m0(globalState)
                    d_2998_v35_ = out111_
                    d_2999_v36_ = out112_
                    d_3000_v37_ = out113_
                    d_3001_v38_ = out114_
                    d_3002_v39_: _dafny.Seq
                    d_3002_v39_ = _dafny.SeqWithoutIsStrInference([d_2999_v36_])
                    d_3003_v40_: _dafny.Map
                    d_3003_v40_ = _dafny.Map({d_2977_v18_: d_2954_v0_})
                    d_3004_v41_: C12
                    nw519_ = C12()
                    nw519_.ctor__(d_3003_v40_, _dafny.Map({d_2977_v18_: len(d_2979_v20_)}))
                    d_3004_v41_ = nw519_
                    d_3005_v42_: _dafny.Map
                    d_3005_v42_ = _dafny.Map({not(d_2998_v35_): d_3004_v41_})
                    d_3006_v43_: _dafny.Seq
                    d_3006_v43_ = _dafny.SeqWithoutIsStrInference([d_3005_v42_])
                    d_3007_v44_: _dafny.Array
                    nw520_ = _dafny.Array(None, 25)
                    nw520_[int(0)] = d_3005_v42_
                    nw520_[int(1)] = d_3005_v42_
                    nw520_[int(2)] = (d_3005_v42_) | (d_3005_v42_)
                    nw520_[int(3)] = d_3005_v42_
                    nw520_[int(4)] = (d_3005_v42_) | ((d_3006_v43_)[default__.safeIndex(d_3001_v38_, len(d_3006_v43_))])
                    nw520_[int(5)] = (d_3005_v42_) | (d_3005_v42_)
                    nw520_[int(6)] = d_3005_v42_
                    nw520_[int(7)] = (d_3006_v43_)[default__.safeIndex(d_3001_v38_, len(d_3006_v43_))]
                    nw520_[int(8)] = d_3005_v42_
                    nw520_[int(9)] = (_dafny.Map({not(not(d_2954_v0_)): d_3004_v41_})) | (d_3005_v42_)
                    nw520_[int(10)] = d_3005_v42_
                    nw520_[int(11)] = _dafny.Map({d_2954_v0_: d_3004_v41_})
                    nw520_[int(12)] = d_3005_v42_
                    nw520_[int(13)] = ((d_3005_v42_).set(d_2999_v36_, d_3004_v41_)) | (d_3005_v42_)
                    nw520_[int(14)] = (d_3005_v42_) | (d_3005_v42_)
                    nw520_[int(15)] = d_3005_v42_
                    nw520_[int(16)] = (d_3005_v42_) | (d_3005_v42_)
                    nw520_[int(17)] = d_3005_v42_
                    nw520_[int(18)] = (d_3005_v42_) | (d_3005_v42_)
                    nw520_[int(19)] = d_3005_v42_
                    nw520_[int(20)] = d_3005_v42_
                    nw520_[int(21)] = d_3005_v42_
                    nw520_[int(22)] = d_3005_v42_
                    nw520_[int(23)] = d_3005_v42_
                    nw520_[int(24)] = d_3005_v42_
                    d_3007_v44_ = nw520_
                    index478_ = default__.safeIndex(625, (d_3007_v44_).length(0))
                    (d_3007_v44_)[index478_] = (d_3005_v42_) | (d_3005_v42_)
                    d_3008_v45_: _dafny.MultiSet
                    d_3008_v45_ = _dafny.MultiSet([d_2999_v36_])
                    index479_ = default__.safeIndex(625, (d_3007_v44_).length(0))
                    rhs465_ = (d_3002_v39_)[default__.safeIndex((d_3001_v38_) * (len(d_2979_v20_)), len(d_3002_v39_))]
                    rhs466_ = d_3002_v39_
                    rhs467_ = d_2954_v0_
                    rhs468_ = (d_3005_v42_) | ((d_3005_v42_) | ((_dafny.Map({d_2954_v0_: d_3004_v41_})).set(d_2999_v36_, d_3004_v41_)))
                    rhs469_ = not (not((d_3008_v45_).ispropersubset(_dafny.MultiSet(d_3002_v39_)))) or (not(True))
                    lhs368_ = globalState
                    lhs369_ = globalState
                    lhs370_ = d_3007_v44_
                    lhs371_ = default__.safeIndex(625, (d_3007_v44_).length(0))
                    lhs372_ = globalState
                    lhs368_.f1 = rhs465_
                    d_3002_v39_ = rhs466_
                    lhs369_.f13 = rhs467_
                    lhs370_[lhs371_] = rhs468_
                    lhs372_.f13 = rhs469_
                    d_3009_v46_: _dafny.Array
                    nw521_ = _dafny.Array(None, 5)
                    nw521_[int(0)] = d_2998_v35_
                    nw521_[int(1)] = False
                    nw521_[int(2)] = d_2998_v35_
                    nw521_[int(3)] = d_2998_v35_
                    nw521_[int(4)] = d_2954_v0_
                    d_3009_v46_ = nw521_
                    d_3010_v47_: _dafny.MultiSet
                    d_3010_v47_ = _dafny.MultiSet([d_3009_v46_, d_3009_v46_])
                    if (d_3010_v47_).isdisjoint(d_3010_v47_):
                        d_3011_v48_: _dafny.Map
                        d_3011_v48_ = _dafny.Map({d_2977_v18_: d_3009_v46_})
                        (globalState).f22 = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))).set(default__.safeIndex(d_2977_v18_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))), _dafny.CodePoint('c')) for d_3012_i5_ in range(default__.abs(459))])), (0) - (len(d_3011_v48_)))
                        d_3013_v49_: _dafny.Array
                        def lambda152_(d_3014_v19_):
                            def lambda153_(d_3015_i6_):
                                return d_3014_v19_

                            return lambda153_

                        init89_ = lambda152_(d_2978_v19_)
                        nw522_ = _dafny.Array(None, 3)
                        for i0_89_ in range(nw522_.length(0)):
                            nw522_[i0_89_] = init89_(i0_89_)
                        d_3013_v49_ = nw522_
                        d_3016_v50_: C0
                        nw523_ = C0()
                        nw523_.ctor__(d_3013_v49_)
                        d_3016_v50_ = nw523_
                        d_2979_v20_ = (d_2979_v20_) + (d_2979_v20_)
                        d_3017_v51_: C9
                        nw524_ = C9()
                        nw524_.ctor__(d_3009_v46_, (((d_3008_v45_)[d_2998_v35_] if (d_2998_v35_) in (d_3008_v45_) else d_2977_v18_)) * ((0) - (d_2977_v18_)), self.f26)
                        d_3017_v51_ = nw524_
                        rhs470_ = d_3017_v51_
                        rhs471_ = d_2979_v20_
                        rhs472_ = (len(d_2979_v20_)) != ((d_3017_v51_).f34)
                        d_3017_v51_ = rhs470_
                        d_2979_v20_ = rhs471_
                        d_2954_v0_ = rhs472_
                        (globalState).f2 = len(_dafny.SeqWithoutIsStrInference([d_2978_v19_ for d_3018_i7_ in range(default__.abs(425))]))
                    elif True:
                        d_3019_v52_: _dafny.Map
                        d_3019_v52_ = _dafny.Map({d_2999_v36_: d_2998_v35_})
                        (globalState).f0 = ((self.f26)[(0) - ((d_2977_v18_) + (len(d_3019_v52_)))] if ((0) - ((d_2977_v18_) + (len(d_3019_v52_)))) in (self.f26) else d_3001_v38_)
                        d_3020_v53_: _dafny.Array
                        nw525_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                        d_3020_v53_ = nw525_
                        index480_ = default__.safeIndex(528, (d_3020_v53_).length(0))
                        (d_3020_v53_)[index480_] = d_2978_v19_
                        index481_ = default__.safeIndex(528, (d_3020_v53_).length(0))
                        (d_3020_v53_)[index481_] = d_2978_v19_
                        d_3021_v54_: _dafny.Array
                        def lambda154_(d_3022_v18_):
                            def lambda155_(d_3023_i8_):
                                return (d_3023_i8_) * (d_3022_v18_)

                            return lambda155_

                        init90_ = lambda154_(d_2977_v18_)
                        nw526_ = _dafny.Array(None, 1)
                        for i0_90_ in range(nw526_.length(0)):
                            nw526_[i0_90_] = init90_(i0_90_)
                        d_3021_v54_ = nw526_
                        rhs473_ = (self).fm3(globalState)
                        rhs474_ = d_3021_v54_
                        rhs475_ = ((d_2979_v20_) + ((d_2979_v20_).set(default__.safeIndex(d_3001_v38_, len(d_2979_v20_)), d_2978_v19_))) + ((d_2979_v20_) + (d_2979_v20_))
                        rhs476_ = d_2954_v0_
                        lhs373_ = globalState
                        lhs373_.f2 = rhs473_
                        d_3021_v54_ = rhs474_
                        d_2979_v20_ = rhs475_
                        d_2998_v35_ = rhs476_
                        d_3024_v55_: _dafny.Map
                        d_3024_v55_ = _dafny.Map({_dafny.MultiSet((d_3002_v39_) + (_dafny.SeqWithoutIsStrInference([d_2999_v36_, d_2998_v35_]))): d_2999_v36_})
                        d_3024_v55_ = _dafny.Map({d_3008_v45_: (len(d_2979_v20_)) >= (d_2977_v18_)})
                        d_3025_v56_: C6
                        nw527_ = C6()
                        nw527_.ctor__()
                        d_3025_v56_ = nw527_
                    d_3026_v57_: _dafny.Array
                    def lambda156_(d_3027_v38_):
                        def lambda157_(d_3028_i9_):
                            return (d_3028_i9_) + (d_3027_v38_)

                        return lambda157_

                    init91_ = lambda156_(d_3001_v38_)
                    nw528_ = _dafny.Array(None, 27)
                    for i0_91_ in range(nw528_.length(0)):
                        nw528_[i0_91_] = init91_(i0_91_)
                    d_3026_v57_ = nw528_
                    index482_ = default__.safeIndex(250, (d_3026_v57_).length(0))
                    (d_3026_v57_)[index482_] = -720
                    d_3029_v58_: _dafny.Map
                    d_3029_v58_ = _dafny.Map({d_2998_v35_: d_2998_v35_})
                    d_3030_v59_: _dafny.Seq
                    d_3030_v59_ = _dafny.SeqWithoutIsStrInference([d_3029_v58_, (d_3029_v58_).set(False, False), d_3029_v58_])
                    index483_ = default__.safeIndex(250, (d_3026_v57_).length(0))
                    rhs477_ = 60
                    rhs478_ = (0) - (default__.safeDivisionInt(d_2977_v18_, 926))
                    rhs479_ = d_2977_v18_
                    rhs480_ = (d_3030_v59_) == ((_dafny.SeqWithoutIsStrInference([d_3029_v58_])) + (_dafny.SeqWithoutIsStrInference([d_3029_v58_, d_3029_v58_, d_3029_v58_])))
                    lhs374_ = globalState
                    lhs375_ = d_3026_v57_
                    lhs376_ = default__.safeIndex(250, (d_3026_v57_).length(0))
                    lhs377_ = globalState
                    lhs374_.f21 = rhs477_
                    lhs375_[lhs376_] = rhs478_
                    lhs377_.f5 = rhs479_
                    d_2998_v35_ = rhs480_
                    pass
            pass
        hi17_ = d_2977_v18_
        for d_3031_i10_ in range(len((d_2979_v20_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltncbpc")))), hi17_):
            (globalState).f1 = d_2954_v0_
            d_3032_v60_: C4
            nw529_ = C4()
            nw529_.ctor__()
            d_3032_v60_ = nw529_
            (globalState).f1 = d_2954_v0_
            if not((False if (self).fm5(False, (0) - (d_2977_v18_), globalState) else d_2954_v0_)):
                d_3033_v61_: _dafny.Array
                def lambda158_(d_3034_v0_):
                    def lambda159_(d_3035_i11_):
                        return not(not (d_3034_v0_) or (True))

                    return lambda159_

                init92_ = lambda158_(d_2954_v0_)
                nw530_ = _dafny.Array(None, 16)
                for i0_92_ in range(nw530_.length(0)):
                    nw530_[i0_92_] = init92_(i0_92_)
                d_3033_v61_ = nw530_
                d_3033_v61_ = d_3033_v61_
                d_3036_v62_: D8
                d_3036_v62_ = D8_DC19(d_2954_v0_, d_2954_v0_, d_2954_v0_)
                d_3036_v62_ = d_3036_v62_
                d_3037_v63_: _dafny.MultiSet
                d_3037_v63_ = _dafny.MultiSet([d_2979_v20_])
                d_3038_v64_: _dafny.MultiSet
                d_3038_v64_ = (d_3037_v63_).intersection(_dafny.MultiSet([(d_2979_v20_).set(default__.safeIndex(len(d_2955_v1_), len(d_2979_v20_)), d_2978_v19_)]))
                d_3038_v64_ = default__.fm50(globalState)
                d_3039_v65_: _dafny.Set
                d_3039_v65_ = _dafny.Set({d_3033_v61_})
                d_3040_v66_: _dafny.Array
                nw531_ = _dafny.Array(None, 27)
                nw531_[int(0)] = len(d_2982_v22_)
                nw531_[int(1)] = len(d_2955_v1_)
                nw531_[int(2)] = d_2977_v18_
                nw531_[int(3)] = d_2977_v18_
                nw531_[int(4)] = d_3031_i10_
                nw531_[int(5)] = d_2977_v18_
                nw531_[int(6)] = d_2977_v18_
                nw531_[int(7)] = (0) - (d_2977_v18_)
                nw531_[int(8)] = d_3031_i10_
                nw531_[int(9)] = d_2977_v18_
                nw531_[int(10)] = d_2977_v18_
                nw531_[int(11)] = d_2977_v18_
                nw531_[int(12)] = default__.fm2(globalState)
                nw531_[int(13)] = d_3031_i10_
                nw531_[int(14)] = d_2977_v18_
                nw531_[int(15)] = d_3031_i10_
                nw531_[int(16)] = d_3031_i10_
                nw531_[int(17)] = d_2977_v18_
                nw531_[int(18)] = d_3031_i10_
                nw531_[int(19)] = len(d_2955_v1_)
                nw531_[int(20)] = len(d_2979_v20_)
                nw531_[int(21)] = d_3031_i10_
                nw531_[int(22)] = d_2977_v18_
                nw531_[int(23)] = len(d_2982_v22_)
                nw531_[int(24)] = d_3031_i10_
                nw531_[int(25)] = len(d_3039_v65_)
                nw531_[int(26)] = -431
                d_3040_v66_ = nw531_
                d_3041_v68_: D0
                def iife290_():
                    coll114_ = _dafny.Map()
                    compr_114_: int
                    for compr_114_ in _dafny.IntegerRange(509, 803):
                        d_3042_v67_: int = compr_114_
                        if ((509) <= (d_3042_v67_)) and ((d_3042_v67_) < (803)):
                            coll114_[default__.safeModuloInt(d_3042_v67_, d_3031_i10_)] = _dafny.SeqWithoutIsStrInference([d_2954_v0_, d_2954_v0_])
                    return _dafny.Map(coll114_)
                d_3041_v68_ = D0_DC1(d_3040_v66_, _dafny.CodePoint('i'), len(iife290_()
))
                d_3043_v69_: D0
                d_3043_v69_ = D0_DC2(d_3041_v68_)
                d_3044_v70_: _dafny.Array
                nw532_ = _dafny.Array(None, 23)
                nw532_[int(0)] = d_3043_v69_
                nw532_[int(1)] = d_3043_v69_
                nw532_[int(2)] = d_3043_v69_
                nw532_[int(3)] = d_3043_v69_
                nw532_[int(4)] = d_3043_v69_
                nw532_[int(5)] = d_3043_v69_
                nw532_[int(6)] = d_3043_v69_
                nw532_[int(7)] = d_3043_v69_
                nw532_[int(8)] = D0_DC2(d_3041_v68_)
                nw532_[int(9)] = d_3043_v69_
                nw532_[int(10)] = D0_DC2(d_3041_v68_)
                nw532_[int(11)] = d_3043_v69_
                nw532_[int(12)] = d_3043_v69_
                nw532_[int(13)] = D0_DC2(d_3041_v68_)
                nw532_[int(14)] = d_3043_v69_
                nw532_[int(15)] = d_3043_v69_
                nw532_[int(16)] = d_3043_v69_
                nw532_[int(17)] = d_3043_v69_
                nw532_[int(18)] = d_3043_v69_
                nw532_[int(19)] = d_3043_v69_
                nw532_[int(20)] = d_3043_v69_
                nw532_[int(21)] = d_3043_v69_
                nw532_[int(22)] = d_3043_v69_
                d_3044_v70_ = nw532_
                d_3045_v71_: _dafny.Seq
                d_3045_v71_ = _dafny.SeqWithoutIsStrInference([d_3044_v70_, d_3044_v70_, d_3044_v70_, d_3044_v70_])
                d_3045_v71_ = d_3045_v71_
                d_2954_v0_ = (((d_2982_v22_)[default__.safeIndex(d_3031_i10_, len(d_2982_v22_))]) - ((0) - (d_2977_v18_))) < (default__.safeDivisionInt(d_3031_i10_, d_3031_i10_))
            elif True:
                d_3046_v72_: C10
                nw533_ = C10()
                nw533_.ctor__((_dafny.MultiSet([not(d_2954_v0_), d_2954_v0_, not(d_2954_v0_), d_2954_v0_, d_2954_v0_])).cardinality, d_2954_v0_)
                d_3046_v72_ = nw533_
                d_3047_v73_: _dafny.Set
                d_3047_v73_ = _dafny.Set({len(d_2979_v20_)})
                def iife291_():
                    coll115_ = _dafny.Set()
                    compr_115_: int
                    for compr_115_ in _dafny.IntegerRange(201, 558):
                        d_3048_v74_: int = compr_115_
                        if ((201) <= (d_3048_v74_)) and ((d_3048_v74_) < (558)):
                            coll115_ = coll115_.union(_dafny.Set([default__.safeDivisionInt(d_3048_v74_, (_dafny.MultiSet([d_2977_v18_])).cardinality)]))
                    return _dafny.Set(coll115_)
                def iife292_():
                    coll116_ = _dafny.Set()
                    compr_116_: int
                    for compr_116_ in (self.f26).keys.Elements:
                        d_3049_v75_: int = compr_116_
                        if (d_3049_v75_) in (self.f26):
                            coll116_ = coll116_.union(_dafny.Set([default__.safeDivisionInt(d_3049_v75_, -331)]))
                    return _dafny.Set(coll116_)
                rhs481_ = ((d_3047_v73_).intersection(_dafny.Set({d_3031_i10_}))).intersection((iife291_()
                ) - (iife292_()
                ))
                rhs482_ = (d_3046_v72_).f32
                rhs483_ = (default__.safeDivisionInt((d_3046_v72_).f31, d_3031_i10_) if d_2954_v0_ else d_3031_i10_)
                rhs484_ = d_2977_v18_
                lhs378_ = globalState
                lhs379_ = globalState
                lhs380_ = globalState
                d_3047_v73_ = rhs481_
                lhs378_.f1 = rhs482_
                lhs379_.f2 = rhs483_
                lhs380_.f0 = rhs484_
                (globalState).f1 = (d_3046_v72_).f32
                d_3050_v76_: _dafny.Map
                d_3050_v76_ = _dafny.Map({d_2954_v0_: not((d_3046_v72_).f32)})
                r2 = (self).fm5(((d_3050_v76_)[d_2954_v0_] if (d_2954_v0_) in (d_3050_v76_) else (d_3046_v72_).f32), d_2977_v18_, globalState)
                d_3051_v77_: _dafny.MultiSet
                d_3051_v77_ = _dafny.MultiSet([d_2977_v18_])
                d_3052_v78_: _dafny.Map
                d_3052_v78_ = _dafny.Map({(d_3051_v77_).cardinality: True})
                d_3053_v79_: _dafny.Map
                d_3053_v79_ = _dafny.Map({d_3052_v78_: _dafny.Map({d_2954_v0_: (d_3046_v72_).f31})})
                d_3054_v80_: D9
                d_3054_v80_ = D9_DC22((d_3053_v79_ if (d_3046_v72_).f32 else d_3053_v79_))
                d_3054_v80_ = d_3054_v80_
        d_3055_v81_: _dafny.Seq
        d_3055_v81_ = _dafny.SeqWithoutIsStrInference([d_2982_v22_, d_2982_v22_])
        r0 = ((self).fm3(globalState)) > (len(d_3055_v81_))
        d_3056_v82_: _dafny.Set
        d_3056_v82_ = _dafny.Set({self.f26})
        r1 = d_3056_v82_
        d_3057_v83_: _dafny.MultiSet
        d_3057_v83_ = _dafny.MultiSet([self.f26])
        r2 = (self).fm5((d_2977_v18_) >= (d_2977_v18_), (d_3057_v83_).cardinality, globalState)
        return r0, r1, r2

    def m3(self, p0, p1, p2, p3, globalState):
        d_3058_v0_: _dafny.Seq
        d_3058_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgfmtkbt"))
        d_3059_v1_: _dafny.Set
        d_3059_v1_ = _dafny.Set({d_3058_v0_})
        d_3060_v2_: int
        d_3060_v2_ = 156
        d_3061_v3_: D7
        d_3061_v3_ = D7_DC16((0) - (d_3060_v2_), True)
        d_3062_v4_: _dafny.Set
        d_3062_v4_ = _dafny.Set({d_3061_v3_})
        d_3063_v5_: _dafny.Seq
        d_3063_v5_ = _dafny.SeqWithoutIsStrInference([False, p3])
        d_3064_v6_: D3
        d_3064_v6_ = D3_DC7(d_3058_v0_, len(d_3062_v4_), d_3060_v2_, d_3060_v2_, d_3063_v5_)
        d_3065_v7_: _dafny.Seq
        d_3065_v7_ = _dafny.SeqWithoutIsStrInference([d_3059_v1_, _dafny.Set({d_3058_v0_, (d_3064_v6_).cf8})])
        d_3066_v8_: _dafny.Map
        d_3066_v8_ = _dafny.Map({((d_3065_v7_)[default__.safeIndex(d_3060_v2_, len(d_3065_v7_))]).intersection(d_3059_v1_): d_3060_v2_})
        (globalState).f22 = ((d_3066_v8_)[(_dafny.Set({d_3058_v0_, d_3058_v0_, d_3058_v0_, d_3058_v0_})).intersection(d_3059_v1_)] if ((_dafny.Set({d_3058_v0_, d_3058_v0_, d_3058_v0_, d_3058_v0_})).intersection(d_3059_v1_)) in (d_3066_v8_) else 161)
        d_3067_v9_: _dafny.Array
        nw534_ = _dafny.Array(int(0), 27)
        d_3067_v9_ = nw534_
        d_3068_v10_: str
        d_3068_v10_ = _dafny.CodePoint('l')
        index484_ = default__.safeIndex(479, (d_3067_v9_).length(0))
        (d_3067_v9_)[index484_] = (0) - (len(((d_3058_v0_).set(default__.safeIndex(812, len(d_3058_v0_)), d_3068_v10_)).set(default__.safeIndex(d_3060_v2_, len((d_3058_v0_).set(default__.safeIndex(812, len(d_3058_v0_)), d_3068_v10_))), d_3068_v10_)))
        index485_ = default__.safeIndex(479, (d_3067_v9_).length(0))
        (d_3067_v9_)[index485_] = d_3060_v2_
        if ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) >= ((d_3060_v2_ if False else (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])):
            index486_ = default__.safeIndex(479, (d_3067_v9_).length(0))
            rhs485_ = d_3058_v0_
            rhs486_ = (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]
            rhs487_ = (False if p2 else p1)
            lhs381_ = d_3067_v9_
            lhs382_ = default__.safeIndex(479, (d_3067_v9_).length(0))
            lhs383_ = globalState
            d_3058_v0_ = rhs485_
            lhs381_[lhs382_] = rhs486_
            lhs383_.f13 = rhs487_
            d_3069_v11_: _dafny.Seq
            d_3069_v11_ = _dafny.SeqWithoutIsStrInference([d_3063_v5_, d_3063_v5_, d_3063_v5_])
            d_3070_v12_: _dafny.Array
            def lambda160_(d_3071_p3_):
                def lambda161_(d_3072_i0_):
                    return not (d_3071_p3_) or (d_3071_p3_)

                return lambda161_

            init93_ = lambda160_(p3)
            nw535_ = _dafny.Array(None, 20)
            for i0_93_ in range(nw535_.length(0)):
                nw535_[i0_93_] = init93_(i0_93_)
            d_3070_v12_ = nw535_
            d_3073_v13_: _dafny.Map
            d_3073_v13_ = _dafny.Map({d_3060_v2_: d_3068_v10_})
            d_3074_v14_: C1
            nw536_ = C1()
            nw536_.ctor__(_dafny.Map({(d_3063_v5_)[default__.safeIndex((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], len(d_3063_v5_))]: (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]}))
            d_3074_v14_ = nw536_
            d_3075_v15_: _dafny.Set
            d_3075_v15_ = _dafny.Set({d_3074_v14_})
            rhs488_ = ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1, p3]), d_3063_v5_])) + (d_3069_v11_) if (d_3075_v15_).isdisjoint(d_3075_v15_) else (d_3069_v11_) + (d_3069_v11_))
            rhs489_ = d_3070_v12_
            rhs490_ = default__.safeModuloInt(len(d_3063_v5_), (0) - (d_3060_v2_))
            rhs491_ = (d_3073_v13_ if not(False) else d_3073_v13_)
            rhs492_ = p3
            lhs384_ = globalState
            lhs385_ = globalState
            d_3069_v11_ = rhs488_
            d_3070_v12_ = rhs489_
            lhs384_.f2 = rhs490_
            d_3073_v13_ = rhs491_
            lhs385_.f13 = rhs492_
            d_3058_v0_ = d_3058_v0_
            (globalState).f0 = (0) - ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])
            if p2:
                index487_ = default__.safeIndex(320, (d_3070_v12_).length(0))
                (d_3070_v12_)[index487_] = p0
                index488_ = default__.safeIndex(320, (d_3070_v12_).length(0))
                (d_3070_v12_)[index488_] = (p2) == ((p3 if p0 else p0))
                index489_ = default__.safeIndex(320, (d_3070_v12_).length(0))
                (d_3070_v12_)[index489_] = not((d_3060_v2_) != ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]))
                d_3076_v16_: T0
                nw537_ = C3()
                nw537_.ctor__(d_3067_v9_, not(p0), self.f26)
                d_3076_v16_ = nw537_
                d_3077_v17_: _dafny.Map
                d_3077_v17_ = _dafny.Map({(d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]: d_3076_v16_})
                d_3077_v17_ = (d_3077_v17_).set((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], d_3076_v16_)
                d_3060_v2_ = (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]
                d_3058_v0_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "us"))) + ((d_3058_v0_) + (d_3058_v0_))
            elif True:
                d_3058_v0_ = d_3058_v0_
                d_3058_v0_ = _dafny.SeqWithoutIsStrInference([d_3068_v10_ for d_3078_i1_ in range(default__.abs(497))])
                d_3079_v18_: _dafny.Map
                d_3079_v18_ = _dafny.Map({d_3058_v0_: d_3067_v9_})
                d_3080_v19_: _dafny.Map
                d_3080_v19_ = _dafny.Map({((d_3079_v18_)[d_3058_v0_] if (d_3058_v0_) in (d_3079_v18_) else d_3067_v9_): default__.fm2(globalState)})
                d_3080_v19_ = d_3080_v19_
                d_3081_v20_: _dafny.Array
                nw538_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_3081_v20_ = nw538_
                d_3082_v21_: D0
                d_3082_v21_ = D0_DC0(d_3081_v20_)
                d_3082_v21_ = (d_3082_v21_ if (self).fm5(p0, d_3060_v2_, globalState) else D0_DC0(d_3081_v20_))
                (globalState).f13 = p1
        elif True:
            d_3083_v22_: _dafny.Map
            d_3083_v22_ = _dafny.Map({d_3063_v5_: d_3060_v2_})
            (globalState).f22 = len((d_3083_v22_).set(d_3063_v5_, (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]))
            d_3084_v23_: _dafny.Seq
            d_3084_v23_ = _dafny.SeqWithoutIsStrInference([d_3060_v2_, d_3060_v2_, d_3060_v2_])
            (globalState).f5 = (d_3084_v23_)[default__.safeIndex((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], len(d_3084_v23_))]
            if p1:
                d_3085_v24_: _dafny.Array
                nw539_ = _dafny.Array(D23.default()(), 19)
                d_3085_v24_ = nw539_
                d_3086_v25_: _dafny.Map
                d_3086_v25_ = _dafny.Map({(d_3084_v23_)[default__.safeIndex((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], len(d_3084_v23_))]: d_3085_v24_})
                d_3086_v25_ = d_3086_v25_
                (globalState).f1 = (p3) and (p1)
                d_3087_v26_: C10
                nw540_ = C10()
                nw540_.ctor__(len(self.f26), p3)
                d_3087_v26_ = nw540_
                d_3088_v27_: _dafny.Map
                d_3088_v27_ = _dafny.Map({d_3087_v26_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bea"))})
                d_3089_v28_: _dafny.Map
                d_3089_v28_ = _dafny.Map({(d_3087_v26_).f32: (d_3087_v26_).f31})
                d_3090_v29_: _dafny.Seq
                d_3090_v29_ = _dafny.SeqWithoutIsStrInference([(d_3058_v0_) + (d_3058_v0_), _dafny.SeqWithoutIsStrInference([d_3068_v10_ for d_3091_i2_ in range(default__.abs(187))]), ((d_3088_v27_)[d_3087_v26_] if (d_3087_v26_) in (d_3088_v27_) else default__.fm30(len(d_3089_v28_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxhamhvro")), (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], globalState))])
                d_3092_v30_: _dafny.Seq
                d_3092_v30_ = default__.fm64(globalState)
                rhs493_ = (0) - (d_3060_v2_)
                rhs494_ = (p2) and (((d_3087_v26_).f31) > ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]))
                rhs495_ = d_3060_v2_
                rhs496_ = ((d_3092_v30_)).set(default__.safeIndex((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], len((d_3092_v30_))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ry")))
                lhs386_ = globalState
                lhs387_ = globalState
                lhs388_ = globalState
                lhs386_.f2 = rhs493_
                lhs387_.f1 = rhs494_
                lhs388_.f22 = rhs495_
                d_3090_v29_ = rhs496_
                (globalState).f22 = d_3060_v2_
                d_3093_v31_: _dafny.Array
                nw541_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_3093_v31_ = nw541_
                d_3094_v32_: _dafny.Map
                d_3094_v32_ = _dafny.Map({d_3093_v31_: False})
                d_3094_v32_ = (d_3094_v32_).set(d_3093_v31_, True)
            elif True:
                d_3095_v33_: _dafny.Map
                d_3095_v33_ = _dafny.Map({d_3060_v2_: (d_3063_v5_)[default__.safeIndex(len(d_3063_v5_), len(d_3063_v5_))]})
                d_3096_v34_: C12
                nw542_ = C12()
                nw542_.ctor__(d_3095_v33_, self.f26)
                d_3096_v34_ = nw542_
                d_3097_v35_: D4
                d_3097_v35_ = D4_DC10(d_3060_v2_)
                (globalState).f21 = (len(_dafny.SeqWithoutIsStrInference([d_3068_v10_ for d_3098_i3_ in range(default__.abs(895))]))) - ((d_3097_v35_).cf15)
                (globalState).f13 = p0
                d_3058_v0_ = (d_3064_v6_).cf8
                (globalState).f1 = ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) >= (d_3060_v2_)
            d_3099_v36_: _dafny.Array
            nw543_ = _dafny.Array(False, 25)
            d_3099_v36_ = nw543_
            d_3100_v37_: _dafny.Map
            d_3100_v37_ = _dafny.Map({len(d_3058_v0_): p2})
            index490_ = default__.safeIndex(344, (d_3099_v36_).length(0))
            (d_3099_v36_)[index490_] = ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) <= (len(d_3100_v37_))
            d_3101_v38_: _dafny.MultiSet
            d_3101_v38_ = _dafny.MultiSet([d_3063_v5_])
            d_3102_v39_: _dafny.MultiSet
            d_3102_v39_ = _dafny.MultiSet([558])
            d_3103_v40_: D7
            d_3103_v40_ = D7_DC17(False, p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wceihuayy"))))
            index491_ = default__.safeIndex(344, (d_3099_v36_).length(0))
            rhs497_ = ((d_3100_v37_)[(0) - (len((d_3058_v0_) + (default__.fm30(((d_3102_v39_)[(d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]] if ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) in (d_3102_v39_) else (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]), d_3058_v0_, d_3060_v2_, globalState))))] if ((0) - (len((d_3058_v0_) + (default__.fm30(((d_3102_v39_)[(d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]] if ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) in (d_3102_v39_) else (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]), d_3058_v0_, d_3060_v2_, globalState))))) in (d_3100_v37_) else (d_3063_v5_)[default__.safeIndex(d_3060_v2_, len(d_3063_v5_))])
            rhs498_ = p2
            rhs499_ = default__.fm65(False, d_3060_v2_, d_3058_v0_, globalState)
            rhs500_ = (d_3103_v40_).cf25
            rhs501_ = (self.f26 if (True) or (p2) else self.f26)
            lhs389_ = globalState
            lhs390_ = d_3099_v36_
            lhs391_ = default__.safeIndex(344, (d_3099_v36_).length(0))
            lhs392_ = globalState
            lhs393_ = globalState
            lhs389_.f13 = rhs497_
            lhs390_[lhs391_] = rhs498_
            d_3101_v38_ = rhs499_
            lhs392_.f1 = rhs500_
            lhs393_.f12 = rhs501_
            index492_ = default__.safeIndex(479, (d_3067_v9_).length(0))
            (d_3067_v9_)[index492_] = (self).fm3(globalState)
        d_3104_v41_: _dafny.Seq
        d_3104_v41_ = _dafny.SeqWithoutIsStrInference([(d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], 96])
        d_3105_v42_: _dafny.Map
        d_3105_v42_ = _dafny.Map({d_3104_v41_: p0})
        if (d_3105_v42_) != (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_3060_v2_]): p2})):
            (globalState).f1 = (True) == (True)
            (globalState).f16 = (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]
            d_3106_v43_: _dafny.Map
            d_3106_v43_ = _dafny.Map({default__.safeModuloInt(d_3060_v2_, (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]): p2})
            def iife293_():
                coll117_ = _dafny.Map()
                compr_117_: int
                for compr_117_ in (d_3104_v41_).Elements:
                    d_3107_v44_: int = compr_117_
                    if (d_3107_v44_) in (d_3104_v41_):
                        coll117_[default__.safeDivisionInt(d_3107_v44_, 206)] = p0
                return _dafny.Map(coll117_)
            d_3106_v43_ = iife293_()
            
            d_3108_v45_: _dafny.Array
            nw544_ = _dafny.Array(False, 21)
            d_3108_v45_ = nw544_
            d_3108_v45_ = d_3108_v45_
            (globalState).f13 = p1
        elif True:
            d_3109_v46_: _dafny.Seq
            d_3109_v46_ = _dafny.SeqWithoutIsStrInference([default__.fm66(d_3058_v0_, d_3060_v2_, globalState)])
            d_3110_v47_: C4
            nw545_ = C4()
            nw545_.ctor__()
            d_3110_v47_ = nw545_
            d_3111_v48_: _dafny.Set
            d_3111_v48_ = _dafny.Set({p3, (d_3063_v5_)[default__.safeIndex((0) - ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]), len(d_3063_v5_))]})
            d_3112_v49_: C5
            nw546_ = C5()
            nw546_.ctor__(d_3060_v2_)
            d_3112_v49_ = nw546_
            d_3113_v50_: _dafny.Seq
            d_3113_v50_ = _dafny.SeqWithoutIsStrInference([d_3112_v49_])
            d_3114_v51_: _dafny.Map
            d_3114_v51_ = _dafny.Map({len(d_3111_v48_): d_3113_v50_})
            d_3115_v52_: _dafny.Map
            d_3115_v52_ = _dafny.Map({d_3111_v48_: d_3060_v2_})
            d_3116_v53_: _dafny.Map
            d_3116_v53_ = d_3115_v52_
            rhs502_ = (((d_3109_v46_) + ((d_3109_v46_).set(default__.safeIndex(len(((d_3114_v51_)[(d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]] if ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) in (d_3114_v51_) else _dafny.SeqWithoutIsStrInference([d_3112_v49_]))), len(d_3109_v46_)), d_3116_v53_))).set(default__.safeIndex(len(d_3058_v0_), len((d_3109_v46_) + ((d_3109_v46_).set(default__.safeIndex(len(((d_3114_v51_)[(d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]] if ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) in (d_3114_v51_) else _dafny.SeqWithoutIsStrInference([d_3112_v49_]))), len(d_3109_v46_)), d_3116_v53_)))), d_3116_v53_)) + (d_3109_v46_)
            rhs503_ = (default__.fm67(globalState)).cf52
            rhs504_ = p1
            rhs505_ = (d_3110_v47_ if p1 else d_3110_v47_)
            lhs394_ = globalState
            lhs395_ = globalState
            d_3109_v46_ = rhs502_
            lhs394_.f0 = rhs503_
            lhs395_.f13 = rhs504_
            d_3110_v47_ = rhs505_
            d_3117_v54_: _dafny.Array
            nw547_ = _dafny.Array(None, 7)
            nw547_[int(0)] = d_3068_v10_
            nw547_[int(1)] = d_3068_v10_
            nw547_[int(2)] = d_3068_v10_
            nw547_[int(3)] = d_3068_v10_
            nw547_[int(4)] = d_3068_v10_
            nw547_[int(5)] = d_3068_v10_
            nw547_[int(6)] = d_3068_v10_
            d_3117_v54_ = nw547_
            def lambda162_(d_3118_v10_):
                def lambda163_(d_3119_i4_):
                    return d_3118_v10_

                return lambda163_

            init94_ = lambda162_(d_3068_v10_)
            nw548_ = _dafny.Array(None, 17)
            for i0_94_ in range(nw548_.length(0)):
                nw548_[i0_94_] = init94_(i0_94_)
            d_3117_v54_ = nw548_
            d_3120_v55_: C5
            nw549_ = C5()
            nw549_.ctor__(default__.fm2(globalState))
            d_3120_v55_ = nw549_
            d_3121_v56_: _dafny.Set
            d_3121_v56_ = _dafny.Set({(d_3120_v55_).f38})
            d_3122_v57_: D8
            d_3122_v57_ = D8_DC18(d_3121_v56_)
            d_3123_v58_: D8
            d_3123_v58_ = D8_DC21(d_3122_v57_)
            d_3124_v59_: D8
            d_3124_v59_ = D8_DC21(d_3123_v58_)
            pat_let_tv69_ = d_3121_v56_
            d_3125_v60_: _dafny.Set
            def iife294_(_pat_let88_0):
                def iife295_(d_3126_dt__update__tmp_h0_):
                    def iife296_(_pat_let89_0):
                        def iife297_(d_3127_dt__update_hcf32_h0_):
                            return D8_DC21(d_3127_dt__update_hcf32_h0_)
                        return iife297_(_pat_let89_0)
                    return iife296_(D8_DC18(pat_let_tv69_))
                return iife295_(_pat_let88_0)
            d_3125_v60_ = _dafny.Set({iife294_(D8_DC21(d_3122_v57_)), d_3124_v59_, d_3124_v59_})
            d_3125_v60_ = ((d_3125_v60_) - (d_3125_v60_)) | ((d_3125_v60_) - (d_3125_v60_))
            if (self).fm5(p3, (d_3112_v49_).f38, globalState):
                d_3128_v61_: D4
                d_3128_v61_ = D4_DC10(d_3060_v2_)
                index493_ = default__.safeIndex(479, (d_3067_v9_).length(0))
                (d_3067_v9_)[index493_] = (0) - ((d_3128_v61_).cf15)
                d_3129_v62_: _dafny.Array
                nw550_ = _dafny.Array(False, 16)
                d_3129_v62_ = nw550_
                d_3130_v63_: _dafny.Map
                d_3130_v63_ = _dafny.Map({(p1 if p3 else p3): d_3129_v62_})
                d_3129_v62_ = ((d_3130_v63_)[(p3) and (p0)] if ((p3) and (p0)) in (d_3130_v63_) else d_3129_v62_)
                d_3063_v5_ = default__.fm13(len(d_3058_v0_), globalState)
                d_3131_v64_: _dafny.Array
                nw551_ = _dafny.Array(_dafny.MultiSet({}), 18)
                d_3131_v64_ = nw551_
                d_3132_v65_: D2
                d_3132_v65_ = D2_DC4(d_3131_v64_)
                d_3132_v65_ = d_3132_v65_
                d_3133_v66_: D5
                d_3133_v66_ = D5_DC12(p3, p1)
                d_3134_v67_: _dafny.MultiSet
                d_3134_v67_ = _dafny.MultiSet([D5_DC12(p1, p2), d_3133_v66_])
                d_3135_v68_: _dafny.Seq
                d_3135_v68_ = _dafny.SeqWithoutIsStrInference([default__.fm68(d_3068_v10_, p1, globalState), d_3134_v67_])
                d_3136_v69_: _dafny.Map
                d_3136_v69_ = _dafny.Map({p1: (d_3112_v49_).f38})
                rhs506_ = (((_dafny.MultiSet([d_3133_v66_])) - ((d_3135_v68_)[default__.safeIndex((d_3112_v49_).f38, len(d_3135_v68_))])).set(d_3133_v66_, default__.abs((len((d_3136_v69_).set(p1, (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]))) * (d_3060_v2_)))).cardinality
                rhs507_ = ((d_3060_v2_) + ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])) >= ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])
                rhs508_ = not(p3)
                lhs396_ = globalState
                lhs397_ = globalState
                lhs398_ = globalState
                lhs396_.f2 = rhs506_
                lhs397_.f13 = rhs507_
                lhs398_.f13 = rhs508_
            elif True:
                d_3137_v70_: _dafny.Map
                d_3137_v70_ = _dafny.Map({len(d_3058_v0_): p0})
                d_3138_v71_: _dafny.Array
                nw552_ = _dafny.Array(None, 4)
                nw552_[int(0)] = p1
                nw552_[int(1)] = ((d_3137_v70_)[-474] if (-474) in (d_3137_v70_) else p1)
                nw552_[int(2)] = ((d_3120_v55_).f38) >= ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])
                nw552_[int(3)] = not(p0)
                d_3138_v71_ = nw552_
                index494_ = default__.safeIndex(463, (d_3138_v71_).length(0))
                (d_3138_v71_)[index494_] = (d_3058_v0_) < (d_3058_v0_)
                index495_ = default__.safeIndex(463, (d_3138_v71_).length(0))
                (d_3138_v71_)[index495_] = p3
                d_3139_v72_: D0
                d_3139_v72_ = D0_DC1(d_3067_v9_, d_3068_v10_, (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])
                d_3139_v72_ = d_3139_v72_
                (globalState).f13 = p2
                (globalState).f22 = (0) - (default__.fm2(globalState))
                d_3140_v73_: _dafny.Map
                d_3140_v73_ = _dafny.Map({d_3060_v2_: (d_3110_v47_).fm22(p1, globalState)})
                d_3141_v74_: D4
                d_3141_v74_ = D4_DC8(((d_3140_v73_)[(d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]] if ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) in (d_3140_v73_) else d_3068_v10_))
                d_3141_v74_ = d_3141_v74_
        d_3142_i5_: int
        d_3142_i5_ = 0
        with _dafny.label("24"):
            while p1:
                with _dafny.c_label("24"):
                    if (d_3142_i5_) >= (100):
                        raise _dafny.Break("24")
                    d_3142_i5_ = (d_3142_i5_) + (1)
                    rhs509_ = (default__.fm28(p1, p1, (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], globalState)) + (_dafny.SeqWithoutIsStrInference([p0]))
                    rhs510_ = d_3063_v5_
                    rhs511_ = p1
                    rhs512_ = (((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) + (len(d_3063_v5_))) + (d_3060_v2_)
                    lhs399_ = globalState
                    lhs400_ = globalState
                    d_3063_v5_ = rhs509_
                    d_3063_v5_ = rhs510_
                    lhs399_.f1 = rhs511_
                    lhs400_.f0 = rhs512_
                    if True:
                        d_3143_v75_: _dafny.Array
                        nw553_ = _dafny.Array(False, 19)
                        d_3143_v75_ = nw553_
                        d_3144_v77_: C9
                        nw554_ = C9()
                        def iife298_():
                            coll118_ = _dafny.Map()
                            compr_118_: int
                            for compr_118_ in _dafny.IntegerRange(442, 936):
                                d_3145_v76_: int = compr_118_
                                if ((442) <= (d_3145_v76_)) and ((d_3145_v76_) < (936)):
                                    coll118_[(d_3145_v76_) * (237)] = (_dafny.MultiSet([True])).cardinality
                            return _dafny.Map(coll118_)
                        nw554_.ctor__((D11_DC26(d_3143_v75_)).cf39, len(d_3104_v41_), iife298_()
                        )
                        d_3144_v77_ = nw554_
                        arr20_ = d_3144_v77_.f33
                        index496_ = default__.safeIndex(978, (d_3144_v77_.f33).length(0))
                        arr20_[index496_] = p3
                        d_3146_v78_: D2
                        d_3146_v78_ = D2_DC5()
                        index497_ = default__.safeIndex(647, (d_3143_v75_).length(0))
                        (d_3143_v75_)[index497_] = (d_3144_v77_).fm11(d_3146_v78_, -404, d_3068_v10_, globalState)
                        arr21_ = d_3144_v77_.f33
                        index498_ = default__.safeIndex(978, (d_3144_v77_.f33).length(0))
                        index499_ = default__.safeIndex(479, (d_3067_v9_).length(0))
                        index500_ = default__.safeIndex(647, (d_3143_v75_).length(0))
                        rhs513_ = p0
                        rhs514_ = (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]
                        rhs515_ = p0
                        rhs516_ = default__.fm0((d_3144_v77_).f34, globalState)
                        lhs401_ = d_3144_v77_.f33
                        lhs402_ = default__.safeIndex(978, (d_3144_v77_.f33).length(0))
                        lhs403_ = d_3067_v9_
                        lhs404_ = default__.safeIndex(479, (d_3067_v9_).length(0))
                        lhs405_ = globalState
                        lhs406_ = d_3143_v75_
                        lhs407_ = default__.safeIndex(647, (d_3143_v75_).length(0))
                        lhs401_[lhs402_] = rhs513_
                        lhs403_[lhs404_] = rhs514_
                        lhs405_.f1 = rhs515_
                        lhs406_[lhs407_] = rhs516_
                        (globalState).f13 = not(not(p0))
                        d_3147_v79_: bool
                        d_3148_v80_: _dafny.Set
                        d_3149_v81_: bool
                        out115_: bool
                        out116_: _dafny.Set
                        out117_: bool
                        out115_, out116_, out117_ = (d_3144_v77_).m2(globalState)
                        d_3147_v79_ = out115_
                        d_3148_v80_ = out116_
                        d_3149_v81_ = out117_
                        d_3150_v82_: _dafny.Seq
                        d_3150_v82_ = _dafny.SeqWithoutIsStrInference([d_3143_v75_])
                        d_3143_v75_ = (d_3150_v82_)[default__.safeIndex(len(d_3063_v5_), len(d_3150_v82_))]
                    elif True:
                        def lambda164_(d_3151_p2_, d_3152_v2_):
                            def lambda165_(d_3153_i6_):
                                return (d_3153_i6_) + (len(_dafny.Map({d_3151_p2_: d_3152_v2_})))

                            return lambda165_

                        init95_ = lambda164_(p2, d_3060_v2_)
                        nw555_ = _dafny.Array(None, 15)
                        for i0_95_ in range(nw555_.length(0)):
                            nw555_[i0_95_] = init95_(i0_95_)
                        d_3067_v9_ = nw555_
                        (globalState).f22 = default__.fm2(globalState)
                        (globalState).f2 = d_3060_v2_
                        d_3154_v83_: C2
                        nw556_ = C2()
                        nw556_.ctor__(p2, d_3060_v2_)
                        d_3154_v83_ = nw556_
                        (globalState).f22 = (default__.safeModuloInt((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], d_3060_v2_)) - ((d_3154_v83_).f42)
                    d_3155_v84_: _dafny.Set
                    d_3155_v84_ = _dafny.Set({d_3060_v2_})
                    d_3156_v85_: _dafny.Map
                    d_3156_v85_ = _dafny.Map({p2: d_3058_v0_})
                    d_3157_v86_: _dafny.MultiSet
                    d_3157_v86_ = _dafny.MultiSet([(0) - ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]), len(((d_3156_v85_)[p2] if (p2) in (d_3156_v85_) else d_3058_v0_))])
                    d_3158_v87_: _dafny.Array
                    nw557_ = _dafny.Array(None, 19)
                    nw557_[int(0)] = False
                    nw557_[int(1)] = True
                    nw557_[int(2)] = (p2) == (default__.fm0((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], globalState))
                    nw557_[int(3)] = p3
                    nw557_[int(4)] = not((d_3155_v84_).ispropersubset(d_3155_v84_))
                    nw557_[int(5)] = p2
                    nw557_[int(6)] = p1
                    nw557_[int(7)] = p3
                    nw557_[int(8)] = p2
                    nw557_[int(9)] = (d_3157_v86_).ispropersubset(_dafny.MultiSet([(d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]]))
                    nw557_[int(10)] = (p2) == (p0)
                    nw557_[int(11)] = (d_3060_v2_) > (d_3060_v2_)
                    nw557_[int(12)] = p1
                    nw557_[int(13)] = False
                    nw557_[int(14)] = p1
                    nw557_[int(15)] = False
                    nw557_[int(16)] = p0
                    nw557_[int(17)] = default__.fm0(len(d_3058_v0_), globalState)
                    nw557_[int(18)] = p3
                    d_3158_v87_ = nw557_
                    d_3159_v88_: _dafny.Map
                    d_3159_v88_ = _dafny.Map({False: d_3158_v87_})
                    d_3158_v87_ = ((d_3159_v88_)[not(p2)] if (not(p2)) in (d_3159_v88_) else d_3158_v87_)
                    index501_ = default__.safeIndex(479, (d_3067_v9_).length(0))
                    (d_3067_v9_)[index501_] = default__.fm2(globalState)
                    pass
            pass
        if p0:
            (globalState).f16 = d_3060_v2_
            d_3160_v89_: _dafny.Map
            d_3160_v89_ = _dafny.Map({False: True})
            d_3160_v89_ = (d_3160_v89_).set(p2, p2)
            d_3161_v90_: _dafny.Set
            d_3161_v90_ = _dafny.Set({d_3063_v5_, _dafny.SeqWithoutIsStrInference([p0]), d_3063_v5_})
            d_3162_v91_: _dafny.Map
            d_3162_v91_ = _dafny.Map({d_3161_v90_: p2})
            d_3163_v92_: _dafny.Seq
            d_3163_v92_ = _dafny.SeqWithoutIsStrInference([d_3161_v90_, d_3161_v90_])
            d_3164_v93_: _dafny.Map
            d_3164_v93_ = _dafny.Map({p0: (d_3163_v92_)[default__.safeIndex(d_3060_v2_, len(d_3163_v92_))]})
            d_3165_v94_: _dafny.Map
            d_3165_v94_ = _dafny.Map({d_3060_v2_: d_3059_v1_})
            d_3166_v95_: _dafny.Map
            d_3166_v95_ = _dafny.Map({d_3058_v0_: p2})
            def iife299_():
                coll119_ = _dafny.Set()
                compr_119_: _dafny.Seq
                for compr_119_ in (d_3166_v95_).keys.Elements:
                    d_3167_v96_: _dafny.Seq = compr_119_
                    if (d_3167_v96_) in (d_3166_v95_):
                        coll119_ = coll119_.union(_dafny.Set([d_3167_v96_]))
                return _dafny.Set(coll119_)
            d_3162_v91_ = (d_3162_v91_).set((d_3161_v90_) | (((d_3164_v93_)[False] if (False) in (d_3164_v93_) else d_3161_v90_)), (((d_3165_v94_)[-420] if (-420) in (d_3165_v94_) else iife299_()
            )).isdisjoint(_dafny.Set({d_3058_v0_, d_3058_v0_})))
            d_3168_v97_: D5
            d_3168_v97_ = D5_DC12(p0, p2)
            d_3169_v98_: _dafny.Seq
            d_3169_v98_ = _dafny.SeqWithoutIsStrInference([d_3058_v0_])
            pat_let_tv70_ = p0
            d_3170_v99_: _dafny.Array
            nw558_ = _dafny.Array(None, 22)
            nw558_[int(0)] = d_3168_v97_
            nw558_[int(1)] = D5_DC12(not(True), p0)
            nw558_[int(2)] = d_3168_v97_
            nw558_[int(3)] = d_3168_v97_
            nw558_[int(4)] = d_3168_v97_
            nw558_[int(5)] = d_3168_v97_
            nw558_[int(6)] = d_3168_v97_
            nw558_[int(7)] = d_3168_v97_
            nw558_[int(8)] = d_3168_v97_
            nw558_[int(9)] = D5_DC12(p1, not(p2))
            def iife300_(_pat_let90_0):
                def iife301_(d_3171_dt__update__tmp_h1_):
                    def iife302_(_pat_let91_0):
                        def iife303_(d_3172_dt__update_hcf17_h0_):
                            return D5_DC12(d_3172_dt__update_hcf17_h0_, (d_3171_dt__update__tmp_h1_).cf18)
                        return iife303_(_pat_let91_0)
                    return iife302_(pat_let_tv70_)
                return iife301_(_pat_let90_0)
            nw558_[int(10)] = iife300_(d_3168_v97_)
            nw558_[int(11)] = D5_DC12((d_3063_v5_)[default__.safeIndex(d_3060_v2_, len(d_3063_v5_))], p0)
            nw558_[int(12)] = D5_DC12(p0, p0)
            nw558_[int(13)] = d_3168_v97_
            nw558_[int(14)] = D5_DC12(p0, p2)
            nw558_[int(15)] = d_3168_v97_
            nw558_[int(16)] = d_3168_v97_
            nw558_[int(17)] = (d_3168_v97_ if not(p0) else d_3168_v97_)
            nw558_[int(18)] = d_3168_v97_
            nw558_[int(19)] = d_3168_v97_
            nw558_[int(20)] = d_3168_v97_
            nw558_[int(21)] = D5_DC12(default__.fm0(len(d_3169_v98_), globalState), not(p3))
            d_3170_v99_ = nw558_
            index502_ = default__.safeIndex(183, (d_3170_v99_).length(0))
            (d_3170_v99_)[index502_] = d_3168_v97_
            d_3173_v100_: _dafny.Array
            nw559_ = _dafny.Array(D9.default()(), 27)
            d_3173_v100_ = nw559_
            index503_ = default__.safeIndex(183, (d_3170_v99_).length(0))
            rhs517_ = d_3168_v97_
            rhs518_ = default__.fm30((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odgkbpiss"))) + ((d_3169_v98_)[default__.safeIndex(d_3060_v2_, len(d_3169_v98_))]), (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], globalState)
            rhs519_ = d_3173_v100_
            rhs520_ = default__.safeDivisionInt((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], d_3060_v2_)
            lhs408_ = d_3170_v99_
            lhs409_ = default__.safeIndex(183, (d_3170_v99_).length(0))
            lhs410_ = globalState
            lhs408_[lhs409_] = rhs517_
            d_3058_v0_ = rhs518_
            d_3173_v100_ = rhs519_
            lhs410_.f21 = rhs520_
            d_3174_v101_: _dafny.Array
            nw560_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_3174_v101_ = nw560_
            d_3175_v102_: _dafny.Map
            d_3175_v102_ = _dafny.Map({p2: len(_dafny.SeqWithoutIsStrInference([d_3068_v10_ for d_3176_i7_ in range(default__.abs(912))]))})
            d_3177_v103_: D4
            d_3177_v103_ = D4_DC9(d_3060_v2_)
            d_3178_v104_: _dafny.Map
            d_3178_v104_ = (d_3175_v102_).set(p2, 909)
            d_3179_v105_: _dafny.Seq
            d_3179_v105_ = _dafny.SeqWithoutIsStrInference([d_3175_v102_])
            d_3180_v106_: _dafny.Array
            nw561_ = _dafny.Array(None, 14)
            nw561_[int(0)] = default__.fm33(d_3104_v41_, p2, globalState)
            nw561_[int(1)] = _dafny.Map({p3: len(_dafny.Map({p1: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ut")))}))})
            nw561_[int(2)] = d_3175_v102_
            nw561_[int(3)] = (d_3175_v102_).set(p0, (d_3177_v103_).cf14)
            nw561_[int(4)] = d_3175_v102_
            nw561_[int(5)] = _dafny.Map({False: (d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]})
            nw561_[int(6)] = (d_3178_v104_)
            nw561_[int(7)] = d_3175_v102_
            nw561_[int(8)] = d_3175_v102_
            nw561_[int(9)] = _dafny.Map({p2: 941})
            nw561_[int(10)] = d_3175_v102_
            nw561_[int(11)] = d_3175_v102_
            nw561_[int(12)] = (d_3179_v105_)[default__.safeIndex(d_3060_v2_, len(d_3179_v105_))]
            nw561_[int(13)] = d_3175_v102_
            d_3180_v106_ = nw561_
            index504_ = default__.safeIndex(997, (d_3174_v101_).length(0))
            (d_3174_v101_)[index504_] = d_3180_v106_
            index505_ = default__.safeIndex(997, (d_3174_v101_).length(0))
            (d_3174_v101_)[index505_] = d_3180_v106_
        elif True:
            if ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) <= (d_3060_v2_):
                d_3068_v10_ = d_3068_v10_
                (globalState).f13 = (d_3060_v2_) != ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])
                (globalState).f13 = ((d_3104_v41_) + (d_3104_v41_)) < (d_3104_v41_)
                d_3181_v107_: _dafny.MultiSet
                d_3181_v107_ = _dafny.MultiSet([d_3060_v2_, len(_dafny.Map({d_3060_v2_: p0}))])
                d_3182_v108_: _dafny.Map
                d_3182_v108_ = _dafny.Map({(_dafny.Map({(d_3181_v107_).cardinality: p1})) | (_dafny.Map({d_3060_v2_: p2})): (d_3058_v0_).set(default__.safeIndex((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], len(d_3058_v0_)), _dafny.CodePoint('h'))})
                d_3182_v108_ = d_3182_v108_
                d_3183_v109_: _dafny.MultiSet
                d_3183_v109_ = _dafny.MultiSet([True, p2, p1])
                d_3184_v110_: _dafny.Set
                d_3184_v110_ = _dafny.Set({p1})
                (globalState).f1 = (not(p3)) and (not((((d_3183_v109_)[p1] if (p1) in (d_3183_v109_) else len(d_3184_v110_))) >= (d_3060_v2_)))
            elif True:
                d_3185_v111_: _dafny.Map
                d_3185_v111_ = _dafny.Map({p2: d_3060_v2_})
                index506_ = default__.safeIndex(479, (d_3067_v9_).length(0))
                rhs521_ = (d_3185_v111_) | ((d_3185_v111_) | (d_3185_v111_))
                rhs522_ = (0) - ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])
                lhs411_ = d_3067_v9_
                lhs412_ = default__.safeIndex(479, (d_3067_v9_).length(0))
                d_3185_v111_ = rhs521_
                lhs411_[lhs412_] = rhs522_
                (globalState).f16 = d_3060_v2_
                d_3058_v0_ = (d_3064_v6_).cf8
                d_3186_v112_: D8
                d_3186_v112_ = D8_DC19(p2, p2, p3)
                d_3186_v112_ = (d_3186_v112_ if p2 else D8_DC19(p1, False, p2))
                (globalState).f1 = True
            d_3187_v113_: C6
            nw562_ = C6()
            nw562_.ctor__()
            d_3187_v113_ = nw562_
            d_3188_v114_: _dafny.Set
            d_3188_v114_ = _dafny.Set({p0, p3})
            d_3189_v115_: _dafny.Seq
            d_3189_v115_ = _dafny.SeqWithoutIsStrInference([(d_3188_v114_) | (_dafny.Set({p1})), d_3188_v114_, d_3188_v114_])
            d_3189_v115_ = d_3189_v115_
            if p1:
                d_3190_v116_: _dafny.Array
                nw563_ = _dafny.Array(False, 1)
                d_3190_v116_ = nw563_
                index507_ = default__.safeIndex(996, (d_3190_v116_).length(0))
                (d_3190_v116_)[index507_] = p1
                index508_ = default__.safeIndex(996, (d_3190_v116_).length(0))
                (d_3190_v116_)[index508_] = not(False)
                (globalState).f22 = (0) - ((((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]) - (len(default__.fm30((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))], d_3058_v0_, len(d_3058_v0_), globalState)))) + ((0) - ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))])))
                index509_ = default__.safeIndex(479, (d_3067_v9_).length(0))
                (d_3067_v9_)[index509_] = d_3060_v2_
                d_3191_v117_: _dafny.Map
                d_3191_v117_ = _dafny.Map({p2: (0) - (d_3060_v2_)})
                d_3192_v118_: _dafny.Set
                d_3192_v118_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ey")))})
                d_3191_v117_ = (d_3191_v117_).set((d_3190_v116_)[default__.safeIndex(996, (d_3190_v116_).length(0))], (len(d_3192_v118_)) + ((d_3067_v9_)[default__.safeIndex(479, (d_3067_v9_).length(0))]))
                (globalState).f21 = d_3060_v2_
            elif True:
                (globalState).f1 = (d_3063_v5_)[default__.safeIndex(640, len(d_3063_v5_))]
                d_3193_v119_: _dafny.Array
                nw564_ = _dafny.Array(False, 19)
                d_3193_v119_ = nw564_
                d_3194_v120_: D0
                d_3195_v121_: _dafny.Set
                out118_: D0
                out119_: _dafny.Set
                out118_, out119_ = (d_3187_v113_).m14(d_3193_v119_, d_3060_v2_, p0, globalState)
                d_3194_v120_ = out118_
                d_3195_v121_ = out119_
                d_3196_v122_: _dafny.Map
                d_3196_v122_ = _dafny.Map({p0: (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_3197_i8_ in range(default__.abs(641))])) == ((d_3058_v0_).set(default__.safeIndex(d_3060_v2_, len(d_3058_v0_)), d_3068_v10_))})
                d_3196_v122_ = d_3196_v122_
                d_3198_v123_: _dafny.Array
                nw565_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_3198_v123_ = nw565_
                d_3199_v124_: C0
                nw566_ = C0()
                nw566_.ctor__(d_3198_v123_)
                d_3199_v124_ = nw566_
                (globalState).f1 = p1
            d_3200_v125_: _dafny.Map
            d_3200_v125_ = _dafny.Map({p1: len(d_3063_v5_)})
            d_3200_v125_ = (d_3200_v125_).set(p2, (0) - (d_3060_v2_))

    def m4(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        d_3201_v0_: _dafny.Seq
        d_3201_v0_ = _dafny.SeqWithoutIsStrInference([p3, p3])
        d_3202_v1_: int
        d_3202_v1_ = 369
        d_3203_v2_: _dafny.MultiSet
        d_3203_v2_ = _dafny.MultiSet([p2, p2, p2])
        d_3204_v3_: _dafny.Seq
        d_3204_v3_ = _dafny.SeqWithoutIsStrInference([((d_3203_v2_)[p2] if (p2) in (d_3203_v2_) else len(p2))])
        d_3205_v4_: _dafny.Set
        d_3205_v4_ = _dafny.Set({(d_3204_v3_)[default__.safeIndex(d_3202_v1_, len(d_3204_v3_))], d_3202_v1_})
        d_3206_v5_: _dafny.Set
        d_3206_v5_ = _dafny.Set({d_3205_v4_})
        d_3207_v6_: _dafny.Array
        nw567_ = _dafny.Array(None, 4)
        nw567_[int(0)] = p3
        nw567_[int(1)] = ((d_3201_v0_)[default__.safeIndex(d_3202_v1_, len(d_3201_v0_))]) | (p3)
        nw567_[int(2)] = p3
        nw567_[int(3)] = _dafny.Map({p0: len(d_3206_v5_)})
        d_3207_v6_ = nw567_
        d_3207_v6_ = d_3207_v6_
        d_3208_v7_: _dafny.Array
        nw568_ = _dafny.Array(None, 3)
        nw568_[int(0)] = (self).fm3(globalState)
        nw568_[int(1)] = d_3202_v1_
        nw568_[int(2)] = d_3202_v1_
        d_3208_v7_ = nw568_
        d_3209_v8_: str
        d_3209_v8_ = _dafny.CodePoint('w')
        d_3210_v9_: D0
        d_3210_v9_ = D0_DC1(d_3208_v7_, d_3209_v8_, d_3202_v1_)
        source38_ = d_3210_v9_
        if source38_.is_DC1:
            d_3211___mcc_h0_ = source38_.cf1
            d_3212___mcc_h1_ = source38_.cf2
            d_3213___mcc_h2_ = source38_.cf3
            d_3214_cf3_ = d_3213___mcc_h2_
            d_3215_cf2_ = d_3212___mcc_h1_
            d_3216_cf1_ = d_3211___mcc_h0_
            d_3217_v10_: C5
            nw569_ = C5()
            nw569_.ctor__(default__.fm2(globalState))
            d_3217_v10_ = nw569_
            d_3218_v11_: D5
            d_3218_v11_ = D5_DC11(p0)
            d_3219_v12_: _dafny.Map
            d_3219_v12_ = _dafny.Map({d_3215_cf2_: (d_3218_v11_).cf16})
            d_3219_v12_ = d_3219_v12_
            d_3220_v13_: _dafny.MultiSet
            d_3220_v13_ = _dafny.MultiSet([p0])
            d_3204_v3_ = default__.fm29(d_3220_v13_, (0) - ((0) - (len(p2))), p0, d_3214_cf3_, globalState)
            (globalState).f21 = default__.fm2(globalState)
        elif source38_.is_DC0:
            d_3221___mcc_h3_ = source38_.cf0
            d_3222_cf0_ = d_3221___mcc_h3_
            if p0:
                (globalState).f13 = p0
                d_3223_v14_: D11
                d_3223_v14_ = D11_DC27(d_3202_v1_, d_3202_v1_, p0)
                (globalState).f13 = ((d_3223_v14_).cf40) >= ((d_3202_v1_) + (d_3202_v1_))
                (globalState).f21 = len(d_3204_v3_)
                d_3209_v8_ = d_3209_v8_
                d_3224_v15_: _dafny.Seq
                d_3224_v15_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
                (globalState).f22 = default__.safeModuloInt(d_3202_v1_, len((d_3224_v15_) + (_dafny.SeqWithoutIsStrInference([p0, p0]))))
            elif True:
                d_3225_v16_: _dafny.Map
                d_3225_v16_ = _dafny.Map({True: d_3208_v7_})
                d_3226_v17_: _dafny.MultiSet
                d_3226_v17_ = _dafny.MultiSet([94])
                d_3227_v18_: _dafny.Seq
                d_3227_v18_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_3228_v19_: _dafny.Array
                nw570_ = _dafny.Array(None, 15)
                nw570_[int(0)] = True
                nw570_[int(1)] = True
                nw570_[int(2)] = default__.fm0(d_3202_v1_, globalState)
                nw570_[int(3)] = p0
                nw570_[int(4)] = p0
                nw570_[int(5)] = (self).fm5(p0, (d_3226_v17_).cardinality, globalState)
                nw570_[int(6)] = p0
                nw570_[int(7)] = p0
                nw570_[int(8)] = True
                nw570_[int(9)] = p0
                nw570_[int(10)] = p0
                nw570_[int(11)] = p0
                nw570_[int(12)] = p0
                nw570_[int(13)] = (d_3227_v18_)[default__.safeIndex(d_3202_v1_, len(d_3227_v18_))]
                nw570_[int(14)] = p0
                d_3228_v19_ = nw570_
                d_3229_v20_: _dafny.Map
                d_3229_v20_ = _dafny.Map({((d_3225_v16_)[p0] if (p0) in (d_3225_v16_) else d_3208_v7_): d_3228_v19_})
                d_3229_v20_ = (d_3229_v20_).set(d_3208_v7_, d_3228_v19_)
                d_3230_v21_: _dafny.Seq
                d_3230_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
                d_3231_v22_: D3
                d_3231_v22_ = D3_DC7(d_3230_v21_, d_3202_v1_, default__.fm2(globalState), 276, _dafny.SeqWithoutIsStrInference([p0, False]))
                d_3230_v21_ = ((d_3231_v22_).cf8) + (d_3230_v21_)
                (globalState).f2 = d_3202_v1_
                (globalState).f1 = not(p0)
                (globalState).f16 = default__.safeDivisionInt(d_3202_v1_, d_3202_v1_)
            d_3232_v23_: _dafny.Seq
            d_3232_v23_ = _dafny.SeqWithoutIsStrInference([d_3208_v7_])
            d_3233_v24_: _dafny.Map
            d_3233_v24_ = _dafny.Map({d_3209_v8_: (d_3232_v23_)[default__.safeIndex(d_3202_v1_, len(d_3232_v23_))]})
            d_3233_v24_ = (d_3233_v24_).set(d_3209_v8_, d_3208_v7_)
            d_3234_v25_: _dafny.Seq
            d_3234_v25_ = _dafny.SeqWithoutIsStrInference([p0, p0, not(True)])
            d_3235_v26_: _dafny.Map
            d_3235_v26_ = _dafny.Map({d_3234_v25_: p0})
            d_3236_v27_: D3
            d_3236_v27_ = D3_DC6(d_3235_v26_)
            source39_ = d_3236_v27_
            if source39_.is_DC7:
                d_3237___mcc_h5_ = source39_.cf8
                d_3238___mcc_h6_ = source39_.cf9
                d_3239___mcc_h7_ = source39_.cf10
                d_3240___mcc_h8_ = source39_.cf11
                d_3241___mcc_h9_ = source39_.cf12
                d_3242_cf12_ = d_3241___mcc_h9_
                d_3243_cf11_ = d_3240___mcc_h8_
                d_3244_cf10_ = d_3239___mcc_h7_
                d_3245_cf9_ = d_3238___mcc_h6_
                d_3246_cf8_ = d_3237___mcc_h5_
                d_3247_v28_: D5
                d_3247_v28_ = D5_DC12(False, p0)
                d_3248_v29_: _dafny.Map
                d_3248_v29_ = _dafny.Map({(p0) or ((d_3247_v28_).cf18): d_3245_cf9_})
                d_3248_v29_ = (d_3248_v29_).set(p0, len(d_3204_v3_))
                d_3249_v30_: D3
                d_3249_v30_ = D3_DC7(p2, len(d_3246_cf8_), 559, ((self.f26)[589] if (589) in (self.f26) else ((self.f26)[(d_3204_v3_)[default__.safeIndex(d_3244_cf10_, len(d_3204_v3_))]] if ((d_3204_v3_)[default__.safeIndex(d_3244_cf10_, len(d_3204_v3_))]) in (self.f26) else d_3202_v1_)), d_3242_cf12_)
                d_3250_v31_: _dafny.Map
                d_3250_v31_ = _dafny.Map({(d_3204_v3_)[default__.safeIndex(d_3245_cf9_, len(d_3204_v3_))]: default__.fm21(d_3209_v8_, 185, d_3249_v30_, globalState)})
                d_3251_v32_: _dafny.Map
                d_3251_v32_ = _dafny.Map({d_3204_v3_: d_3245_cf9_})
                d_3252_v33_: D25
                d_3252_v33_ = D25_DC49(_dafny.Map({d_3245_cf9_: d_3204_v3_}))
                rhs523_ = default__.safeModuloInt(len(d_3251_v32_), d_3202_v1_)
                rhs524_ = ((d_3252_v33_).cf66) | (d_3250_v31_)
                rhs525_ = p0
                lhs413_ = globalState
                lhs414_ = globalState
                lhs413_.f16 = rhs523_
                d_3250_v31_ = rhs524_
                lhs414_.f13 = rhs525_
                d_3208_v7_ = d_3208_v7_
                d_3253_v34_: C3
                nw571_ = C3()
                nw571_.ctor__(d_3208_v7_, p0, self.f26)
                d_3253_v34_ = nw571_
                d_3254_v35_: _dafny.Seq
                d_3254_v35_ = _dafny.SeqWithoutIsStrInference([d_3253_v34_, d_3253_v34_, d_3253_v34_, d_3253_v34_])
                (globalState).f1 = (default__.safeModuloInt(d_3244_cf10_, (0) - (((self.f26)[d_3202_v1_] if (d_3202_v1_) in (self.f26) else (0) - (d_3245_cf9_))))) < (((0) - (len(p1))) + (len(d_3254_v35_)))
            elif True:
                d_3255___mcc_h10_ = source39_.cf7
                d_3256_cf7_ = d_3255___mcc_h10_
                d_3257_v36_: C1
                nw572_ = C1()
                nw572_.ctor__(_dafny.Map({p0: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igncu")))}))
                d_3257_v36_ = nw572_
                d_3258_v37_: _dafny.Map
                d_3258_v37_ = _dafny.Map({True: p1})
                (globalState).f13 = ((d_3202_v1_) <= (((self.f26)[d_3202_v1_] if (d_3202_v1_) in (self.f26) else d_3202_v1_))) not in (((d_3258_v37_).set(p0, p1) if p0 else d_3258_v37_))
                d_3259_v38_: D3
                d_3259_v38_ = D3_DC7(_dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3260_i0_ in range(default__.abs(828))]), d_3202_v1_, d_3202_v1_, d_3202_v1_, _dafny.SeqWithoutIsStrInference([p0, p0, p0]))
                (globalState).f0 = (d_3204_v3_)[default__.safeIndex(len(default__.fm21(d_3209_v8_, d_3202_v1_, d_3259_v38_, globalState)), len(d_3204_v3_))]
                (self).f26 = (self.f26).set((d_3202_v1_) + (d_3202_v1_), default__.safeModuloInt((0) - (d_3202_v1_), d_3202_v1_))
            d_3261_v39_: _dafny.Map
            d_3261_v39_ = p3
            d_3262_v40_: T1
            nw573_ = C10()
            nw573_.ctor__(d_3202_v1_, (p2) != (p2))
            d_3262_v40_ = nw573_
            rhs526_ = (_dafny.Map({p0: len(p2)})) | (p3)
            rhs527_ = d_3262_v40_
            rhs528_ = d_3202_v1_
            lhs415_ = globalState
            d_3261_v39_ = rhs526_
            d_3262_v40_ = rhs527_
            lhs415_.f0 = rhs528_
        elif True:
            d_3263___mcc_h4_ = source38_.cf4
            d_3264_cf4_ = d_3263___mcc_h4_
            d_3265_v41_: _dafny.Map
            d_3265_v41_ = _dafny.Map({p0: p0})
            d_3266_v42_: _dafny.Seq
            d_3266_v42_ = _dafny.SeqWithoutIsStrInference([((d_3265_v41_)[p0] if (p0) in (d_3265_v41_) else p0)])
            d_3267_v43_: _dafny.Array
            nw574_ = _dafny.Array(None, 21)
            nw574_[int(0)] = p0
            nw574_[int(1)] = p0
            nw574_[int(2)] = False
            nw574_[int(3)] = not(p0)
            nw574_[int(4)] = p0
            nw574_[int(5)] = (d_3202_v1_) >= (902)
            nw574_[int(6)] = p0
            nw574_[int(7)] = True
            nw574_[int(8)] = ((d_3265_v41_)[p0] if (p0) in (d_3265_v41_) else p0)
            nw574_[int(9)] = (p2) != (p2)
            nw574_[int(10)] = True
            nw574_[int(11)] = False
            nw574_[int(12)] = False
            nw574_[int(13)] = (d_3266_v42_)[default__.safeIndex(436, len(d_3266_v42_))]
            nw574_[int(14)] = p0
            nw574_[int(15)] = p0
            nw574_[int(16)] = True
            nw574_[int(17)] = (len(self.f26)) != (d_3202_v1_)
            nw574_[int(18)] = p0
            nw574_[int(19)] = False
            nw574_[int(20)] = p0
            d_3267_v43_ = nw574_
            d_3267_v43_ = d_3267_v43_
            d_3268_v44_: _dafny.Array
            nw575_ = _dafny.Array(_dafny.Seq({}), 16)
            d_3268_v44_ = nw575_
            index510_ = default__.safeIndex(236, (d_3268_v44_).length(0))
            (d_3268_v44_)[index510_] = d_3266_v42_
            index511_ = default__.safeIndex(236, (d_3268_v44_).length(0))
            (d_3268_v44_)[index511_] = _dafny.SeqWithoutIsStrInference([(d_3266_v42_)[default__.safeIndex(d_3202_v1_, len(d_3266_v42_))], (p0 if p0 else p0)])
            if not (p0) or (p0):
                (globalState).f0 = ((759 if not(not(p0)) else d_3202_v1_)) - (d_3202_v1_)
                d_3269_v45_: _dafny.Seq
                d_3269_v45_ = d_3204_v3_
                d_3270_v46_: _dafny.Map
                d_3270_v46_ = _dafny.Map({d_3269_v45_: d_3209_v8_})
                rhs529_ = ((d_3270_v46_)[d_3204_v3_] if (d_3204_v3_) in (d_3270_v46_) else d_3209_v8_)
                rhs530_ = not ((d_3204_v3_) <= (d_3204_v3_)) or ((p2) < (p2))
                lhs416_ = globalState
                d_3209_v8_ = rhs529_
                lhs416_.f13 = rhs530_
                (globalState).f1 = p0
                (globalState).f16 = default__.fm2(globalState)
                d_3271_v47_: _dafny.Map
                d_3271_v47_ = _dafny.Map({d_3202_v1_: p0})
                (self).f26 = (self.f26).set(default__.safeDivisionInt(len(d_3271_v47_), d_3202_v1_), 288)
            elif True:
                d_3272_v48_: C10
                nw576_ = C10()
                nw576_.ctor__(d_3202_v1_, (p2) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "no"))))
                d_3272_v48_ = nw576_
                index512_ = default__.safeIndex(253, (d_3267_v43_).length(0))
                (d_3267_v43_)[index512_] = (d_3265_v41_) == (d_3265_v41_)
                d_3273_v49_: _dafny.Map
                d_3273_v49_ = _dafny.Map({(d_3272_v48_).f31: p2})
                d_3274_v50_: D3
                d_3274_v50_ = D3_DC7(((d_3273_v49_)[len(_dafny.SeqWithoutIsStrInference([d_3202_v1_ for d_3275_i1_ in range(default__.abs(431))]))] if (len(_dafny.SeqWithoutIsStrInference([d_3202_v1_ for d_3276_i1_ in range(default__.abs(431))]))) in (d_3273_v49_) else p2), (d_3272_v48_).f31, (d_3272_v48_).f31, len(p2), _dafny.SeqWithoutIsStrInference([(d_3272_v48_).f32]))
                d_3277_v51_: _dafny.Array
                nw577_ = _dafny.Array(None, 28)
                nw577_[int(0)] = p2
                nw577_[int(1)] = p2
                nw577_[int(2)] = (p2) + (p2)
                nw577_[int(3)] = ((d_3274_v50_).cf8).set(default__.safeIndex((d_3272_v48_).f31, len((d_3274_v50_).cf8)), d_3209_v8_)
                nw577_[int(4)] = _dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3278_i2_ in range(default__.abs(203))])
                nw577_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpohneur"))).set(default__.safeIndex(len(p2), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpohneur")))), _dafny.CodePoint('p'))
                nw577_[int(6)] = p2
                nw577_[int(7)] = p2
                nw577_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "od"))
                nw577_[int(9)] = _dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3279_i3_ in range(default__.abs(42))])
                nw577_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rktuv"))
                nw577_[int(11)] = (p2 if False else _dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3280_i4_ in range(default__.abs(683))]))
                nw577_[int(12)] = p2
                nw577_[int(13)] = p2
                nw577_[int(14)] = (p2) + (p2)
                nw577_[int(15)] = p2
                nw577_[int(16)] = default__.fm30((d_3272_v48_).f31, p2, (d_3272_v48_).f31, globalState)
                nw577_[int(17)] = _dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3281_i5_ in range(default__.abs(659))])
                nw577_[int(18)] = p2
                nw577_[int(19)] = p2
                nw577_[int(20)] = (p2) + (p2)
                nw577_[int(21)] = (p2 if p0 else p2)
                nw577_[int(22)] = p2
                nw577_[int(23)] = (p2).set(default__.safeIndex(d_3202_v1_, len(p2)), d_3209_v8_)
                nw577_[int(24)] = p2
                nw577_[int(25)] = p2
                nw577_[int(26)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvifqtkpl"))) + (p2)
                nw577_[int(27)] = p2
                d_3277_v51_ = nw577_
                index513_ = default__.safeIndex(705, (d_3277_v51_).length(0))
                (d_3277_v51_)[index513_] = (p2) + (default__.fm30(len(p2), p2, 878, globalState))
                d_3282_v52_: _dafny.Seq
                d_3282_v52_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                d_3283_v53_: D7
                d_3283_v53_ = D7_DC17((d_3272_v48_).f32, p0, len((d_3282_v52_)[default__.safeIndex((d_3272_v48_).f31, len(d_3282_v52_))]))
                index514_ = default__.safeIndex(253, (d_3267_v43_).length(0))
                index515_ = default__.safeIndex(705, (d_3277_v51_).length(0))
                rhs531_ = ((d_3283_v53_).cf26) > ((D4_DC9(len(d_3204_v3_))).cf14)
                rhs532_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjpsltsq"))
                lhs417_ = d_3267_v43_
                lhs418_ = default__.safeIndex(253, (d_3267_v43_).length(0))
                lhs419_ = d_3277_v51_
                lhs420_ = default__.safeIndex(705, (d_3277_v51_).length(0))
                lhs417_[lhs418_] = rhs531_
                lhs419_[lhs420_] = rhs532_
                d_3284_v54_: _dafny.Array
                nw578_ = _dafny.Array(False, 21)
                d_3284_v54_ = nw578_
                d_3285_v55_: T0
                nw579_ = C9()
                nw579_.ctor__(d_3284_v54_, len((_dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3286_i6_ in range(default__.abs(162))])) + (p2)), (self.f26) | (_dafny.Map({(d_3272_v48_).f31: d_3202_v1_})))
                d_3285_v55_ = nw579_
                d_3285_v55_ = d_3285_v55_
                index516_ = default__.safeIndex(253, (d_3267_v43_).length(0))
                (d_3267_v43_)[index516_] = (len(p2)) >= (d_3202_v1_)
                (globalState).f13 = p0
            index517_ = default__.safeIndex(248, (d_3267_v43_).length(0))
            (d_3267_v43_)[index517_] = p0
            index518_ = default__.safeIndex(248, (d_3267_v43_).length(0))
            (d_3267_v43_)[index518_] = p0
        d_3287_v56_: D23
        d_3287_v56_ = D23_DC46(p0)
        d_3288_i7_: int
        d_3288_i7_ = 0
        with _dafny.label("25"):
            while (d_3287_v56_).cf63:
                with _dafny.c_label("25"):
                    if (d_3288_i7_) >= (100):
                        raise _dafny.Break("25")
                    d_3288_i7_ = (d_3288_i7_) + (1)
                    index519_ = default__.safeIndex(515, (d_3208_v7_).length(0))
                    (d_3208_v7_)[index519_] = (self).fm3(globalState)
                    index520_ = default__.safeIndex(515, (d_3208_v7_).length(0))
                    (d_3208_v7_)[index520_] = d_3202_v1_
                    d_3289_v57_: _dafny.Seq
                    d_3289_v57_ = _dafny.SeqWithoutIsStrInference([True])
                    pat_let_tv71_ = d_3289_v57_
                    def iife304_(_pat_let92_0):
                        def iife305_(d_3290_dt__update__tmp_h0_):
                            def iife306_(_pat_let93_0):
                                def iife307_(d_3291_dt__update_hcf68_h0_):
                                    return D25_DC50((d_3290_dt__update__tmp_h0_).cf67, d_3291_dt__update_hcf68_h0_)
                                return iife307_(_pat_let93_0)
                            return iife306_(pat_let_tv71_)
                        return iife305_(_pat_let92_0)
                    source40_ = iife304_(D25_DC50(p0, d_3289_v57_))
                    if source40_.is_DC50:
                        d_3292___mcc_h11_ = source40_.cf67
                        d_3293___mcc_h12_ = source40_.cf68
                        d_3294_cf68_ = d_3293___mcc_h12_
                        d_3295_cf67_ = d_3292___mcc_h11_
                        (globalState).f13 = d_3295_cf67_
                        (globalState).f2 = (d_3208_v7_)[default__.safeIndex(515, (d_3208_v7_).length(0))]
                        d_3296_v58_: _dafny.Map
                        d_3296_v58_ = _dafny.Map({(d_3208_v7_)[default__.safeIndex(515, (d_3208_v7_).length(0))]: d_3295_cf67_})
                        d_3297_v59_: T0
                        nw580_ = C8()
                        nw580_.ctor__(d_3295_cf67_, D5_DC11(d_3295_cf67_), self.f26)
                        d_3297_v59_ = nw580_
                        d_3298_v60_: _dafny.Map
                        d_3298_v60_ = _dafny.Map({d_3202_v1_: d_3297_v59_})
                        d_3299_v61_: _dafny.Map
                        d_3299_v61_ = _dafny.Map({d_3294_cf68_: len(d_3298_v60_)})
                        d_3300_v62_: _dafny.Map
                        d_3300_v62_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3301_i8_ in range(default__.abs(-207))])).set(default__.safeIndex(d_3202_v1_, len(_dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3302_i8_ in range(default__.abs(-207))]))), d_3209_v8_): d_3209_v8_})
                        d_3303_v63_: _dafny.MultiSet
                        d_3303_v63_ = _dafny.MultiSet([(0) - (len(d_3300_v62_))])
                        d_3304_v64_: _dafny.Array
                        nw581_ = _dafny.Array(None, 17)
                        nw581_[int(0)] = (d_3295_cf67_) == (d_3295_cf67_)
                        nw581_[int(1)] = (p1).issubset(p1)
                        nw581_[int(2)] = p0
                        nw581_[int(3)] = (d_3202_v1_) >= (len(d_3296_v58_))
                        nw581_[int(4)] = d_3295_cf67_
                        nw581_[int(5)] = (d_3294_cf68_)[default__.safeIndex(len(d_3299_v61_), len(d_3294_cf68_))]
                        nw581_[int(6)] = True
                        nw581_[int(7)] = (d_3303_v63_).isdisjoint(d_3303_v63_)
                        nw581_[int(8)] = False
                        nw581_[int(9)] = False
                        nw581_[int(10)] = (p1).isdisjoint(p1)
                        nw581_[int(11)] = d_3295_cf67_
                        nw581_[int(12)] = (p0 if default__.fm0(d_3202_v1_, globalState) else d_3295_cf67_)
                        nw581_[int(13)] = d_3295_cf67_
                        nw581_[int(14)] = p0
                        nw581_[int(15)] = (d_3295_cf67_) and (not(d_3295_cf67_))
                        nw581_[int(16)] = d_3295_cf67_
                        d_3304_v64_ = nw581_
                        index521_ = default__.safeIndex(68, (d_3304_v64_).length(0))
                        (d_3304_v64_)[index521_] = d_3295_cf67_
                        index522_ = default__.safeIndex(68, (d_3304_v64_).length(0))
                        (d_3304_v64_)[index522_] = False
                        d_3305_v65_: _dafny.Seq
                        d_3305_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntm"))
                        d_3305_v65_ = (d_3305_v65_) + (p2)
                    elif True:
                        d_3306___mcc_h13_ = source40_.cf66
                        d_3307_cf66_ = d_3306___mcc_h13_
                        d_3308_v66_: C6
                        nw582_ = C6()
                        nw582_.ctor__()
                        d_3308_v66_ = nw582_
                        d_3309_v67_: _dafny.Map
                        d_3309_v67_ = _dafny.Map({d_3308_v66_: (d_3208_v7_)[default__.safeIndex(515, (d_3208_v7_).length(0))]})
                        d_3309_v67_ = (d_3309_v67_).set((d_3308_v66_ if p0 else d_3308_v66_), len(d_3204_v3_))
                        d_3310_v68_: C10
                        nw583_ = C10()
                        nw583_.ctor__(d_3202_v1_, p0)
                        d_3310_v68_ = nw583_
                        rhs533_ = not ((d_3310_v68_).f32) or (p0)
                        rhs534_ = d_3310_v68_
                        r2 = rhs533_
                        d_3310_v68_ = rhs534_
                        d_3311_v69_: bool
                        d_3312_v70_: bool
                        d_3313_v71_: _dafny.Map
                        d_3314_v72_: int
                        out120_: bool
                        out121_: bool
                        out122_: _dafny.Map
                        out123_: int
                        out120_, out121_, out122_, out123_ = default__.m0(globalState)
                        d_3311_v69_ = out120_
                        d_3312_v70_ = out121_
                        d_3313_v71_ = out122_
                        d_3314_v72_ = out123_
                        d_3315_v73_: _dafny.Map
                        d_3315_v73_ = _dafny.Map({len(p2): p2})
                        (self).f26 = (self.f26).set(914, len(d_3315_v73_))
                    d_3289_v57_ = (_dafny.SeqWithoutIsStrInference([p0])) + (default__.fm13(d_3202_v1_, globalState))
                    d_3316_v74_: _dafny.MultiSet
                    d_3316_v74_ = _dafny.MultiSet([p0])
                    d_3317_v75_: _dafny.Map
                    d_3317_v75_ = _dafny.Map({(default__.fm0(d_3202_v1_, globalState)) not in (d_3316_v74_): d_3202_v1_})
                    d_3317_v75_ = (d_3317_v75_).set(False, (d_3208_v7_)[default__.safeIndex(515, (d_3208_v7_).length(0))])
                    pass
            pass
        d_3318_v76_: _dafny.Map
        d_3318_v76_ = _dafny.Map({p0: p0})
        if (d_3318_v76_) == ((d_3318_v76_).set(p0, not(((d_3318_v76_)[p0] if (p0) in (d_3318_v76_) else p0)))):
            if p0:
                (globalState).f21 = (414) + (d_3202_v1_)
                (globalState).f5 = len(d_3204_v3_)
                r2 = (not (p0) or (p0)) == (True)
                (globalState).f0 = d_3202_v1_
                (globalState).f2 = d_3202_v1_
            elif True:
                (globalState).f16 = -761
                (globalState).f22 = 414
                (globalState).f1 = p0
                d_3319_v77_: _dafny.Array
                nw584_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                d_3319_v77_ = nw584_
                d_3320_v78_: C0
                nw585_ = C0()
                nw585_.ctor__(d_3319_v77_)
                d_3320_v78_ = nw585_
                rhs535_ = p0
                rhs536_ = len(p2)
                rhs537_ = d_3202_v1_
                lhs421_ = globalState
                lhs422_ = globalState
                lhs423_ = globalState
                lhs421_.f13 = rhs535_
                lhs422_.f16 = rhs536_
                lhs423_.f16 = rhs537_
            d_3321_v79_: _dafny.Array
            def lambda166_(d_3322_p0_):
                def lambda167_(d_3323_i10_):
                    return _dafny.Map({d_3322_p0_: d_3322_p0_})

                return lambda167_

            init96_ = lambda166_(p0)
            nw586_ = _dafny.Array(None, 27)
            for i0_96_ in range(nw586_.length(0)):
                nw586_[i0_96_] = init96_(i0_96_)
            d_3321_v79_ = nw586_
            d_3324_v80_: _dafny.Map
            d_3324_v80_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3325_i9_ in range(default__.abs(474))]): d_3321_v79_})
            d_3324_v80_ = (d_3324_v80_).set((p2) + (p2), d_3321_v79_)
            d_3326_v81_: _dafny.Array
            nw587_ = _dafny.Array(None, 28)
            nw587_[int(0)] = d_3209_v8_
            nw587_[int(1)] = d_3209_v8_
            nw587_[int(2)] = d_3209_v8_
            nw587_[int(3)] = d_3209_v8_
            nw587_[int(4)] = d_3209_v8_
            nw587_[int(5)] = d_3209_v8_
            nw587_[int(6)] = d_3209_v8_
            nw587_[int(7)] = d_3209_v8_
            nw587_[int(8)] = (p2)[default__.safeIndex(d_3202_v1_, len(p2))]
            nw587_[int(9)] = d_3209_v8_
            nw587_[int(10)] = d_3209_v8_
            nw587_[int(11)] = d_3209_v8_
            nw587_[int(12)] = d_3209_v8_
            nw587_[int(13)] = d_3209_v8_
            nw587_[int(14)] = _dafny.CodePoint('p')
            nw587_[int(15)] = d_3209_v8_
            nw587_[int(16)] = d_3209_v8_
            nw587_[int(17)] = _dafny.CodePoint('v')
            nw587_[int(18)] = default__.fm35(d_3202_v1_, globalState)
            nw587_[int(19)] = d_3209_v8_
            nw587_[int(20)] = d_3209_v8_
            nw587_[int(21)] = d_3209_v8_
            nw587_[int(22)] = d_3209_v8_
            nw587_[int(23)] = d_3209_v8_
            nw587_[int(24)] = d_3209_v8_
            nw587_[int(25)] = d_3209_v8_
            nw587_[int(26)] = (p2)[default__.safeIndex(d_3202_v1_, len(p2))]
            nw587_[int(27)] = d_3209_v8_
            d_3326_v81_ = nw587_
            d_3327_v82_: C0
            nw588_ = C0()
            nw588_.ctor__(d_3326_v81_)
            d_3327_v82_ = nw588_
            d_3328_v83_: _dafny.Array
            nw589_ = _dafny.Array(None, 12)
            nw589_[int(0)] = _dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3329_i11_ in range(default__.abs(26))])
            nw589_[int(1)] = p2
            nw589_[int(2)] = p2
            nw589_[int(3)] = p2
            nw589_[int(4)] = p2
            nw589_[int(5)] = p2
            nw589_[int(6)] = p2
            nw589_[int(7)] = p2
            nw589_[int(8)] = p2
            nw589_[int(9)] = p2
            nw589_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmnmrn"))
            nw589_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jschioeii"))
            d_3328_v83_ = nw589_
            index523_ = default__.safeIndex(348, (d_3328_v83_).length(0))
            (d_3328_v83_)[index523_] = p2
            index524_ = default__.safeIndex(348, (d_3328_v83_).length(0))
            (d_3328_v83_)[index524_] = p2
            rhs538_ = d_3204_v3_
            rhs539_ = _dafny.SeqWithoutIsStrInference([d_3202_v1_, d_3202_v1_])
            d_3204_v3_ = rhs538_
            d_3204_v3_ = rhs539_
        elif True:
            d_3330_v84_: D8
            d_3330_v84_ = D8_DC18(d_3205_v4_)
            d_3331_v85_: D8
            d_3331_v85_ = D8_DC21(d_3330_v84_)
            d_3332_v86_: D8
            d_3332_v86_ = D8_DC21(d_3330_v84_)
            source41_ = d_3332_v86_
            if source41_.is_DC19:
                d_3333___mcc_h14_ = source41_.cf28
                d_3334___mcc_h15_ = source41_.cf29
                d_3335___mcc_h16_ = source41_.cf30
                d_3336_cf30_ = d_3335___mcc_h16_
                d_3337_cf29_ = d_3334___mcc_h15_
                d_3338_cf28_ = d_3333___mcc_h14_
                (globalState).f16 = d_3202_v1_
                d_3339_v87_: _dafny.Array
                nw590_ = _dafny.Array(None, 20)
                nw590_[int(0)] = d_3208_v7_
                nw590_[int(1)] = d_3208_v7_
                nw590_[int(2)] = d_3208_v7_
                nw590_[int(3)] = d_3208_v7_
                nw590_[int(4)] = (d_3208_v7_ if p0 else d_3208_v7_)
                nw590_[int(5)] = d_3208_v7_
                nw590_[int(6)] = d_3208_v7_
                nw590_[int(7)] = d_3208_v7_
                nw590_[int(8)] = d_3208_v7_
                nw590_[int(9)] = d_3208_v7_
                nw590_[int(10)] = d_3208_v7_
                nw590_[int(11)] = d_3208_v7_
                nw590_[int(12)] = d_3208_v7_
                nw590_[int(13)] = d_3208_v7_
                nw590_[int(14)] = d_3208_v7_
                nw590_[int(15)] = d_3208_v7_
                nw590_[int(16)] = d_3208_v7_
                nw590_[int(17)] = d_3208_v7_
                nw590_[int(18)] = d_3208_v7_
                nw590_[int(19)] = d_3208_v7_
                d_3339_v87_ = nw590_
                index525_ = default__.safeIndex(838, (d_3339_v87_).length(0))
                (d_3339_v87_)[index525_] = d_3208_v7_
                d_3340_v88_: _dafny.Set
                d_3340_v88_ = _dafny.Set({p2})
                d_3341_v89_: _dafny.Map
                d_3341_v89_ = _dafny.Map({d_3202_v1_: d_3208_v7_})
                d_3342_v90_: _dafny.Seq
                d_3342_v90_ = _dafny.SeqWithoutIsStrInference([d_3208_v7_])
                d_3343_v91_: _dafny.Map
                d_3343_v91_ = _dafny.Map({p2: _dafny.Set({d_3336_cf30_})})
                index526_ = default__.safeIndex(838, (d_3339_v87_).length(0))
                def iife308_():
                    coll120_ = _dafny.Set()
                    compr_120_: _dafny.Seq
                    for compr_120_ in (d_3343_v91_).keys.Elements:
                        d_3344_v92_: _dafny.Seq = compr_120_
                        if (d_3344_v92_) in (d_3343_v91_):
                            coll120_ = coll120_.union(_dafny.Set([d_3344_v92_]))
                    return _dafny.Set(coll120_)
                rhs540_ = (d_3202_v1_) > (d_3202_v1_)
                rhs541_ = ((d_3341_v89_)[263] if (263) in (d_3341_v89_) else (d_3342_v90_)[default__.safeIndex(d_3202_v1_, len(d_3342_v90_))])
                rhs542_ = d_3209_v8_
                rhs543_ = iife308_()

                lhs424_ = globalState
                lhs425_ = d_3339_v87_
                lhs426_ = default__.safeIndex(838, (d_3339_v87_).length(0))
                lhs424_.f13 = rhs540_
                lhs425_[lhs426_] = rhs541_
                d_3209_v8_ = rhs542_
                d_3340_v88_ = rhs543_
                d_3345_v93_: _dafny.Seq
                d_3345_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqtcm"))
                d_3346_v94_: _dafny.Seq
                d_3346_v94_ = _dafny.SeqWithoutIsStrInference([d_3338_cf28_])
                d_3347_v95_: _dafny.MultiSet
                d_3347_v95_ = _dafny.MultiSet([d_3202_v1_, len(d_3346_v94_)])
                rhs544_ = default__.fm35(d_3202_v1_, globalState)
                rhs545_ = p2
                rhs546_ = default__.fm28((d_3347_v95_).ispropersubset(_dafny.MultiSet([d_3202_v1_, d_3202_v1_, d_3202_v1_, 669, d_3202_v1_])), (d_3205_v4_) == (d_3205_v4_), d_3202_v1_, globalState)
                d_3209_v8_ = rhs544_
                d_3345_v93_ = rhs545_
                d_3346_v94_ = rhs546_
                d_3348_v96_: _dafny.Array
                nw591_ = _dafny.Array(False, 20)
                d_3348_v96_ = nw591_
                d_3348_v96_ = d_3348_v96_
            elif source41_.is_DC20:
                d_3349___mcc_h17_ = source41_.cf31
                d_3350_cf31_ = d_3349___mcc_h17_
                d_3351_v97_: _dafny.Set
                d_3351_v97_ = _dafny.Set({(d_3204_v3_ if p0 else d_3204_v3_), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_3202_v1_ for d_3352_i12_ in range(default__.abs(594))])), d_3202_v1_])})
                d_3353_v98_: _dafny.Seq
                d_3353_v98_ = _dafny.SeqWithoutIsStrInference([p0])
                d_3354_v99_: _dafny.Seq
                d_3354_v99_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_3202_v1_, d_3202_v1_]), d_3204_v3_])
                def iife309_():
                    coll121_ = _dafny.Set()
                    compr_121_: _dafny.Seq
                    for compr_121_ in (d_3354_v99_).Elements:
                        d_3355_v100_: _dafny.Seq = compr_121_
                        if (d_3355_v100_) in (d_3354_v99_):
                            coll121_ = coll121_.union(_dafny.Set([d_3355_v100_]))
                    return _dafny.Set(coll121_)
                rhs547_ = (iife309_()
                ).intersection(d_3351_v97_)
                rhs548_ = (d_3353_v98_).set(default__.safeIndex(144, len(d_3353_v98_)), (default__.fm0(d_3202_v1_, globalState) if p0 else p0))
                d_3351_v97_ = rhs547_
                d_3353_v98_ = rhs548_
                d_3356_v101_: _dafny.Seq
                d_3356_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wawaugdw"))
                d_3357_v102_: _dafny.Seq
                d_3357_v102_ = _dafny.SeqWithoutIsStrInference([(d_3356_v101_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iq"))), _dafny.SeqWithoutIsStrInference([d_3209_v8_ for d_3358_i13_ in range(default__.abs(624))])])
                rhs549_ = (d_3357_v102_)[default__.safeIndex(d_3202_v1_, len(d_3357_v102_))]
                rhs550_ = d_3350_cf31_
                lhs427_ = globalState
                d_3356_v101_ = rhs549_
                lhs427_.f13 = rhs550_
                d_3359_v103_: D7
                d_3359_v103_ = D7_DC17((True if (self).fm5(False, d_3202_v1_, globalState) else not(p0)), d_3350_cf31_, d_3202_v1_)
                pat_let_tv72_ = d_3202_v1_
                pat_let_tv73_ = d_3202_v1_
                def iife310_(_pat_let94_0):
                    def iife311_(d_3360_dt__update__tmp_h1_):
                        def iife312_(_pat_let95_0):
                            def iife313_(d_3361_dt__update_hcf26_h0_):
                                return D7_DC17((d_3360_dt__update__tmp_h1_).cf24, (d_3360_dt__update__tmp_h1_).cf25, d_3361_dt__update_hcf26_h0_)
                            return iife313_(_pat_let95_0)
                        return iife312_(((0) - (pat_let_tv72_)) + (pat_let_tv73_))
                    return iife311_(_pat_let94_0)
                d_3359_v103_ = iife310_(D7_DC17(d_3350_cf31_, d_3350_cf31_, d_3202_v1_))
                index527_ = default__.safeIndex(13, (d_3208_v7_).length(0))
                (d_3208_v7_)[index527_] = default__.safeModuloInt((0) - (d_3202_v1_), d_3202_v1_)
                index528_ = default__.safeIndex(13, (d_3208_v7_).length(0))
                (d_3208_v7_)[index528_] = default__.fm2(globalState)
            elif source41_.is_DC18:
                d_3362___mcc_h18_ = source41_.cf27
                d_3363_cf27_ = d_3362___mcc_h18_
                d_3202_v1_ = (d_3202_v1_) * (len((d_3363_cf27_) | (d_3205_v4_)))
                (globalState).f0 = d_3202_v1_
                d_3364_v104_: C7
                nw592_ = C7()
                nw592_.ctor__(p0)
                d_3364_v104_ = nw592_
                d_3364_v104_ = d_3364_v104_
                d_3365_v105_: _dafny.MultiSet
                d_3365_v105_ = _dafny.MultiSet([p0])
                d_3366_v106_: _dafny.Set
                d_3366_v106_ = _dafny.Set({not((d_3364_v104_).f37), not(not((default__.fm0(len(p2), globalState)) not in (d_3365_v105_))), ((d_3364_v104_).f37) and (not((d_3364_v104_).f37))})
                d_3366_v106_ = d_3366_v106_
            elif True:
                d_3367___mcc_h19_ = source41_.cf32
                d_3368_cf32_ = d_3367___mcc_h19_
                d_3369_v107_: _dafny.Map
                d_3369_v107_ = _dafny.Map({d_3202_v1_: True})
                d_3369_v107_ = (d_3369_v107_).set(d_3202_v1_, p0)
                d_3370_v108_: _dafny.MultiSet
                d_3370_v108_ = _dafny.MultiSet([p0, p0])
                (globalState).f13 = ((d_3369_v107_)[d_3202_v1_] if (d_3202_v1_) in (d_3369_v107_) else (d_3370_v108_).isdisjoint(d_3370_v108_))
                (globalState).f13 = p0
                (globalState).f0 = 290
            (globalState).f13 = (p2) <= ((p2) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))).set(default__.safeIndex(d_3202_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))), _dafny.CodePoint('q'))))
            d_3371_v109_: C1
            nw593_ = C1()
            nw593_.ctor__((_dafny.Map({p0: d_3202_v1_})) | (_dafny.Map({p0: d_3202_v1_})))
            d_3371_v109_ = nw593_
            (globalState).f13 = p0
            d_3205_v4_ = _dafny.Set({d_3202_v1_, 433, d_3202_v1_})
        d_3372_v110_: bool
        d_3373_v111_: bool
        d_3374_v112_: _dafny.Map
        d_3375_v113_: int
        out124_: bool
        out125_: bool
        out126_: _dafny.Map
        out127_: int
        out124_, out125_, out126_, out127_ = default__.m0(globalState)
        d_3372_v110_ = out124_
        d_3373_v111_ = out125_
        d_3374_v112_ = out126_
        d_3375_v113_ = out127_
        pat_let_tv74_ = d_3375_v113_
        pat_let_tv75_ = d_3375_v113_
        pat_let_tv76_ = d_3202_v1_
        pat_let_tv77_ = d_3375_v113_
        pat_let_tv78_ = d_3204_v3_
        pat_let_tv79_ = d_3375_v113_
        pat_let_tv80_ = d_3204_v3_
        pat_let_tv81_ = d_3375_v113_
        def lambda168_(source42_):
            if source42_.is_DC19:
                d_3376___mcc_h20_ = source42_.cf28
                d_3377___mcc_h21_ = source42_.cf29
                d_3378___mcc_h22_ = source42_.cf30
                d_3379_cf30_ = d_3378___mcc_h22_
                d_3380_cf29_ = d_3377___mcc_h21_
                d_3381_cf28_ = d_3376___mcc_h20_
                return pat_let_tv74_
            elif source42_.is_DC20:
                d_3382___mcc_h23_ = source42_.cf31
                d_3383_cf31_ = d_3382___mcc_h23_
                return pat_let_tv75_
            elif source42_.is_DC18:
                d_3384___mcc_h24_ = source42_.cf27
                d_3385_cf27_ = d_3384___mcc_h24_
                return default__.safeModuloInt(pat_let_tv76_, pat_let_tv77_)
            elif True:
                d_3386___mcc_h25_ = source42_.cf32
                d_3387_cf32_ = d_3386___mcc_h25_
                return default__.safeModuloInt((pat_let_tv78_)[default__.safeIndex(pat_let_tv79_, len(pat_let_tv80_))], pat_let_tv81_)

        d_3375_v113_ = lambda168_(D8_DC21(D8_DC20(d_3372_v110_)))
        r0 = default__.fm2(globalState)
        r1 = d_3208_v7_
        r2 = d_3372_v110_
        return r0, r1, r2

