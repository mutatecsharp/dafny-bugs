// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, globalState) {
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0))),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_module.D9.create_DC24(!(false), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), function (_0_i0) {
  return new _dafny.CodePoint('s'.codePointAt(0));
})).length)))).length), new BigNumber(596)))).length);
    };
    static fm1(p0, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("hpvvkqoa")).length))), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(false)).length))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(-524)))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(766)))));
    };
    static fm2(p0, globalState) {
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(false), !(false), true, true, false), _dafny.Seq.of(false))).length)).isLessThanOrEqualTo((new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(!(true),false))).Keys.Elements) {
          let _1_v0 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(!(true),false))).contains(_1_v0)) {
            _coll0.push([_1_v0,false]);
          }
        }
        return _coll0;
      }()).length)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-803),new BigNumber((_dafny.Seq.of(new BigNumber(-538), new BigNumber(-54))).length))).length)));
    };
    static fm3(p0, globalState) {
      return new _dafny.CodePoint('y'.codePointAt(0));
    };
    static fm4(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!((_module.D1.create_DC2(false)).dtor_cf2) || (!(!(true))),_dafny.Seq.of(true));
    };
    static fm5(p0, p1, p2, globalState) {
      return _module.D0.create_DC1(new _dafny.CodePoint('c'.codePointAt(0)));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, !(true), !(true)), _dafny.Seq.of(false)), _dafny.Seq.of(false, true, false));
    };
    static fm11(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("c");
    };
    static fm12(p0, p1, p2, globalState) {
      return _module.D3.create_DC7((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber(409)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(176))).cardinality()), new BigNumber(766)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true, false)).length)))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber(-517), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, true)).length))), _dafny.MultiSet.fromElements(new BigNumber(264)))));
    };
    static fm13(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),!(false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),false)));
    };
    static fm14(p0, p1, globalState) {
      return _dafny.Set.fromElements(_module.D1.create_DC2(true));
    };
    static fm15(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false));
    };
    static fm16(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-115)), _dafny.Seq.of(new BigNumber(879), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-823)), function (_2_i0) {
        return (_module.D7.create_DC19(new _dafny.CodePoint('n'.codePointAt(0)), false)).dtor_cf29;
      })).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), function (_3_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(!(true),false);
      })).length)), _dafny.Seq.of(new BigNumber(373))));
    };
    static fm17(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(false, false)).Intersect((_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(false)));
    };
    static fm18(globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_module.D4.create_DC11(new BigNumber((_dafny.Seq.of(_module.D1.create_DC3(new _dafny.CodePoint('t'.codePointAt(0))))).length), new _dafny.CodePoint('f'.codePointAt(0)))).dtor_cf14)).length),new BigNumber((_dafny.MultiSet.FromArray((_module.D6.create_DC16(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(916),new BigNumber(337))).length))))).dtor_cf22)).cardinality())), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(44),new BigNumber(907)));
    };
    static fm19(globalState) {
      let _source0 = ((false) ? (_module.D4.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(516),new BigNumber(525)))) : (_module.D4.create_DC10(function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (function () {
    let _coll2 = new _dafny.Map();
    for (const _compr_2 of _dafny.IntegerRange(new BigNumber(375), new BigNumber(105))) {
      let _4_v1 = _compr_2;
      if (((new BigNumber(375)).isLessThanOrEqualTo(_4_v1)) && ((_4_v1).isLessThan(new BigNumber(105)))) {
        _coll2.push([(_4_v1).multipliedBy(new BigNumber((function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(916), new BigNumber(690))) {
            let _6_v2 = _compr_3;
            if (((new BigNumber(916)).isLessThanOrEqualTo(_6_v2)) && ((_6_v2).isLessThan(new BigNumber(690)))) {
              _coll3.push([(_6_v2).minus(new BigNumber(893)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(90))).length),new _dafny.CodePoint('c'.codePointAt(0)))).length)]);
            }
          }
          return _coll3;
        }()).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(434)), function (_5_i0) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        })).length)]);
      }
    }
    return _coll2;
  }()).Keys.Elements) {
    let _7_v0 = _compr_1;
    if ((function () {
      let _coll4 = new _dafny.Map();
      for (const _compr_4 of _dafny.IntegerRange(new BigNumber(375), new BigNumber(105))) {
        let _4_v1 = _compr_4;
        if (((new BigNumber(375)).isLessThanOrEqualTo(_4_v1)) && ((_4_v1).isLessThan(new BigNumber(105)))) {
          _coll4.push([(_4_v1).multipliedBy(new BigNumber((function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of _dafny.IntegerRange(new BigNumber(916), new BigNumber(690))) {
              let _6_v2 = _compr_5;
              if (((new BigNumber(916)).isLessThanOrEqualTo(_6_v2)) && ((_6_v2).isLessThan(new BigNumber(690)))) {
                _coll5.push([(_6_v2).minus(new BigNumber(893)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(90))).length),new _dafny.CodePoint('c'.codePointAt(0)))).length)]);
              }
            }
            return _coll5;
          }()).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(434)), function (_5_i0) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          })).length)]);
        }
      }
      return _coll4;
    }()).contains(_7_v0)) {
      _coll1.push([_module.__default.safeModuloInt(_7_v0, new BigNumber(-34)),new BigNumber(360)]);
    }
  }
  return _coll1;
}())));
      if (_source0.is_DC11) {
        let _8___mcc_h0 = (_source0).cf13;
        let _9___mcc_h1 = (_source0).cf14;
        let _10_cf14 = _9___mcc_h1;
        let _11_cf13 = _8___mcc_h0;
        return _module.D3.create_DC9(_module.D3.create_DC9(_module.D3.create_DC7(function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_11_cf13, _11_cf13, _11_cf13))).Elements) {
    let _12_v3 = _compr_6;
    if ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_11_cf13, _11_cf13, _11_cf13))).contains(_12_v3)) {
      _coll6.add(_12_v3);
    }
  }
  return _coll6;
}())));
      } else if (_source0.is_DC12) {
        let _13___mcc_h2 = (_source0).cf15;
        let _14___mcc_h3 = (_source0).cf16;
        let _15___mcc_h4 = (_source0).cf17;
        let _16_cf17 = _15___mcc_h4;
        let _17_cf16 = _14___mcc_h3;
        let _18_cf15 = _13___mcc_h2;
        return _module.D3.create_DC9(_module.D3.create_DC8(_18_cf15));
      } else if (_source0.is_DC10) {
        let _19___mcc_h5 = (_source0).cf12;
        let _20_cf12 = _19___mcc_h5;
        return _module.D3.create_DC9(_module.D3.create_DC8(new BigNumber(923)));
      } else {
        let _21___mcc_h6 = (_source0).cf18;
        let _22_cf18 = _21___mcc_h6;
        return _module.D3.create_DC9(_module.D3.create_DC7(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("uh")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(573),true)).length))).length)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(!(true)))).length)))));
      }
    };
    static fm20(p0, p1, globalState) {
      let _source1 = ((!(true)) ? (_module.D7.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),false),function () {
  let _coll7 = new _dafny.Map();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(910), new BigNumber(290))) {
    let _23_v0 = _compr_7;
    if (((new BigNumber(910)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(290)))) {
      _coll7.push([_module.__default.safeModuloInt(_23_v0, new BigNumber(50)),new _dafny.CodePoint('q'.codePointAt(0))]);
    }
  }
  return _coll7;
}()))) : (_module.D7.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(37),new _dafny.CodePoint('y'.codePointAt(0)))))));
      if (_source1.is_DC19) {
        let _24___mcc_h0 = (_source1).cf29;
        let _25___mcc_h1 = (_source1).cf30;
        let _26_cf30 = _25___mcc_h1;
        let _27_cf29 = _24___mcc_h0;
        return (_dafny.Set.fromElements(_26_cf30, false, _26_cf30)).Intersect(_dafny.Set.fromElements(_26_cf30));
      } else {
        let _28___mcc_h2 = (_source1).cf28;
        let _29_cf28 = _28___mcc_h2;
        return (_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(true));
      }
    };
    static fm21(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(false,(new BigNumber((_dafny.Seq.UnicodeFromString("pqgqi")).length)).multipliedBy(new BigNumber(-128)));
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(719), new BigNumber(317))) {
          let _30_v0 = _compr_8;
          if (((new BigNumber(719)).isLessThanOrEqualTo(_30_v0)) && ((_30_v0).isLessThan(new BigNumber(317)))) {
            _coll8.push([(_30_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('r'.codePointAt(0)))).length))),((_module.D6.create_DC17(_dafny.Seq.UnicodeFromString("uyc"), false, new BigNumber((_dafny.Seq.UnicodeFromString("eda")).length), true, false)).dtor_cf25).minus(new BigNumber(47))]);
          }
        }
        return _coll8;
      }();
    };
    static fm23(globalState) {
      return _module.D4.create_DC10((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber((_dafny.Seq.of(true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(118))).length)),new BigNumber((_dafny.Seq.UnicodeFromString("xbwtruw")).length))).length),new BigNumber(304))));
    };
    static m0(globalState) {
      let _31_v0;
      _31_v0 = true;
      (globalState).f3 = _31_v0;
      let _32_v1;
      _32_v1 = new BigNumber(524);
      let _33_v2;
      _33_v2 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new _dafny.CodePoint('x'.codePointAt(0)), globalState),_32_v1);
      _33_v2 = (_33_v2).update(_32_v1, ((_dafny.ZERO).minus(_32_v1)).minus(_32_v1));
      let _34_v3;
      let _init0 = ((_35_v0) => function (_36_i0) {
        return (_module.D7.create_DC19(new _dafny.CodePoint('o'.codePointAt(0)), _35_v0)).dtor_cf30;
      })(_31_v0);
      let _nw0 = Array((new BigNumber(26)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _34_v3 = _nw0;
      let _37_v4;
      _37_v4 = _module.D8.create_DC21(_32_v1);
      let _pat_let_tv0 = _31_v0;
      let _pat_let_tv1 = _31_v0;
      let _pat_let_tv2 = _31_v0;
      let _pat_let_tv3 = _31_v0;
      let _rhs0 = (_32_v1).isEqualTo(_32_v1);
      let _rhs1 = function (_source2) {
        if (_source2.is_DC21) {
          let _38___mcc_h0 = (_source2).cf32;
          let _39_cf32 = _38___mcc_h0;
          return !((_dafny.MultiSet.fromElements(new BigNumber(665), _39_cf32)).IsDisjointFrom(_dafny.MultiSet.fromElements(_39_cf32, (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.D6.create_DC17(_dafny.Seq.UnicodeFromString("ekoegwtab"), _pat_let_tv0, new BigNumber((_dafny.Seq.UnicodeFromString("mkjxih")).length), _pat_let_tv1, _pat_let_tv2)).dtor_cf25)))));
        } else {
          let _40___mcc_h1 = (_source2).cf31;
          let _41_cf31 = _40___mcc_h1;
          return _pat_let_tv3;
        }
      }(_37_v4);
      let _rhs2 = ((_31_v0) ? (_34_v3) : (_34_v3));
      let _lhs0 = globalState;
      _lhs0.f3 = _rhs0;
      _31_v0 = _rhs1;
      _34_v3 = _rhs2;
      let _42_v5;
      _42_v5 = _dafny.Map.Empty.slice().updateUnsafe(_31_v0,new BigNumber((_dafny.Seq.UnicodeFromString("jouwqcwf")).length));
      let _43_v6;
      _43_v6 = new _dafny.CodePoint('x'.codePointAt(0));
      let _44_v7;
      _44_v7 = _module.D4.create_DC11(_32_v1, _43_v6);
      let _45_v8;
      _45_v8 = _dafny.MultiSet.fromElements(_31_v0);
      let _46_v9;
      let _nw1 = new _module.C2();
      _nw1.__ctor((_44_v7).dtor_cf13, _32_v1, (((_45_v8).contains(!(!(_31_v0)))) ? ((_45_v8).get(!(!(_31_v0)))) : (_32_v1)), new BigNumber((_33_v2).length));
      _46_v9 = _nw1;
      let _47_v10;
      _47_v10 = _dafny.Map.Empty.slice().updateUnsafe((_42_v5).Merge(_42_v5),_46_v9);
      let _48_v11;
      _48_v11 = _dafny.Seq.UnicodeFromString("vyfp");
      _47_v10 = (_47_v10).update(((_42_v5).update(_31_v0, (_dafny.ZERO).minus(new BigNumber((_48_v11).length)))).Merge(_42_v5), _46_v9);
      let _49_v12;
      _49_v12 = _dafny.Set.fromElements(_34_v3, _34_v3, _34_v3);
      let _hi0 = (_46_v9).f9;
      for (let _50_i1 = new BigNumber((_49_v12).length); _50_i1.isLessThan(_hi0); _50_i1 = _50_i1.plus(_dafny.ONE)) {
        if (_31_v0) {
          let _51_v13;
          let _nw2 = new _module.C2();
          _nw2.__ctor(_46_v9.f10, _32_v1, (_46_v9).f9, (_46_v9).f9);
          _51_v13 = _nw2;
          let _52_v14;
          _52_v14 = _dafny.Seq.of(_51_v13);
          let _53_v15;
          _53_v15 = _dafny.Map.Empty.slice().updateUnsafe(_31_v0,_52_v14);
          (_46_v9).f10 = new BigNumber((_53_v15).length);
          (globalState).f3 = !((_31_v0) === (_31_v0)) || (_31_v0);
          (_46_v9).f10 = (_51_v13).f7;
          _34_v3 = ((_31_v0) ? (_34_v3) : (((_31_v0) ? (_34_v3) : (_34_v3))));
          let _54_v16;
          _54_v16 = _dafny.Seq.of(_32_v1, (_46_v9).f9);
          _48_v11 = _dafny.Seq.update(_module.__default.fm11(_31_v0, (_dafny.ZERO).minus(_50_i1), globalState), _module.__default.safeIndex((_54_v16)[_module.__default.safeIndex(_50_i1, new BigNumber((_54_v16).length))], new BigNumber((_module.__default.fm11(_31_v0, (_dafny.ZERO).minus(_50_i1), globalState)).length)), _43_v6);
        } else {
          let _55_v17;
          _55_v17 = _dafny.Map.Empty.slice().updateUnsafe(true,_45_v8);
          (globalState).f1 = !(((((_55_v17).contains(!(_31_v0))) ? ((_55_v17).get(!(_31_v0))) : (_45_v8))).IsSubsetOf(_module.__default.fm17(_31_v0, (_46_v9).f9, globalState)));
          let _56_v18;
          _56_v18 = _dafny.Map.Empty.slice().updateUnsafe(_43_v6,_31_v0);
          let _57_v20;
          _57_v20 = _dafny.Map.Empty.slice().updateUnsafe(_56_v18,function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of _dafny.IntegerRange(new BigNumber(791), new BigNumber(145))) {
              let _58_v19 = _compr_9;
              if (((new BigNumber(791)).isLessThanOrEqualTo(_58_v19)) && ((_58_v19).isLessThan(new BigNumber(145)))) {
                _coll9.push([(_58_v19).plus(new BigNumber(293)),_43_v6]);
              }
            }
            return _coll9;
          }());
          let _59_v21;
          let _nw3 = new _module.C4();
          _nw3.__ctor((_57_v20).Merge(_57_v20));
          _59_v21 = _nw3;
          let _60_v22;
          _60_v22 = _dafny.Seq.of(_31_v0);
          let _61_v23;
          _61_v23 = _module.D4.create_DC12((_46_v9).f9, _60_v22, _31_v0);
          let _62_v24;
          let _nw4 = new _module.C1();
          _nw4.__ctor(_31_v0, (_dafny.ZERO).minus(new BigNumber((_45_v8).cardinality())), (_46_v9).f9, new BigNumber(((_61_v23).dtor_cf16).length));
          _62_v24 = _nw4;
          let _63_v25;
          let _nw5 = Array((new BigNumber(3)).toNumber());
          _nw5[(_dafny.ZERO).toNumber()] = _62_v24;
          _nw5[(_dafny.ONE).toNumber()] = _62_v24;
          _nw5[(new BigNumber(2)).toNumber()] = _62_v24;
          _63_v25 = _nw5;
          let _index0 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_63_v25).length));
          (_63_v25)[_index0] = _62_v24;
          let _index1 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_63_v25).length));
          (_63_v25)[_index1] = _62_v24;
          _32_v1 = (_dafny.ZERO).minus(_50_i1);
          let _64_v26;
          let _nw6 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _64_v26 = _nw6;
          _64_v26 = ((_31_v0) ? (_64_v26) : (_64_v26));
        }
        let _65_v27;
        let _nw7 = Array((new BigNumber(25)).toNumber()).fill([]);
        _65_v27 = _nw7;
        let _66_v28;
        _66_v28 = _module.D4.create_DC13(_module.__default.fm23(globalState));
        let _67_v29;
        _67_v29 = _module.D4.create_DC13(_66_v28);
        let _68_v30;
        _68_v30 = _dafny.Seq.of(_67_v29);
        let _69_v31;
        _69_v31 = _dafny.Map.Empty.slice().updateUnsafe(_65_v27,_68_v30);
        _69_v31 = (_69_v31).update(_65_v27, _68_v30);
        let _70_v32;
        _70_v32 = _dafny.Seq.of(_31_v0, true, _31_v0, false);
        (_46_v9).f10 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(((_31_v0) ? (_70_v32) : (_70_v32)), _70_v32)).length));
        let _71_v34;
        _71_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_33_v2);
        let _72_v35;
        _72_v35 = _dafny.Seq.of((_46_v9).f9, _50_i1, new BigNumber(((((_71_v34).contains(true)) ? ((_71_v34).get(true)) : (_33_v2))).length), (_46_v9).fm7(_dafny.Seq.update(_module.__default.fm11(_31_v0, (_46_v9).f9, globalState), _module.__default.safeIndex(_50_i1, new BigNumber((_module.__default.fm11(_31_v0, (_46_v9).f9, globalState)).length)), _43_v6), (_46_v9).f9, globalState));
        let _73_v36;
        let _nw8 = new _module.C1();
        _nw8.__ctor(((((_42_v5).contains(_31_v0)) ? ((_42_v5).get(_31_v0)) : ((_46_v9).f9))).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_31_v0)).length))), new BigNumber((function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of (_72_v35).Elements) {
            let _74_v33 = _compr_10;
            if (_dafny.Seq.contains(_72_v35, _74_v33)) {
              _coll10.push([(_74_v33).plus(_46_v9.f10),_46_v9.f10]);
            }
          }
          return _coll10;
        }()).length), (_46_v9.f10).plus((_46_v9).f9), _46_v9.f10);
        _73_v36 = _nw8;
        _73_v36 = _73_v36;
      }
      let _75_v37;
      _75_v37 = _dafny.Set.fromElements(_31_v0, _31_v0, _31_v0);
      let _index2 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_34_v3).length));
      (_34_v3)[_index2] = (_75_v37).IsProperSubsetOf(_75_v37);
      let _index3 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_34_v3).length));
      (_34_v3)[_index3] = _module.__default.fm2(_31_v0, globalState);
      return;
    }
    static Main(__noArgsParameter) {
      let _76_globalState;
      let _nw9 = new _module.GlobalState();
      _nw9.__ctor(true, true, new BigNumber(135), true, false, new _dafny.CodePoint('e'.codePointAt(0)));
      _76_globalState = _nw9;
      let _77_v0;
      _77_v0 = _dafny.Seq.UnicodeFromString("jc");
      _77_v0 = _dafny.Seq.UnicodeFromString("pqitbtm");
      let _78_v1;
      _78_v1 = new BigNumber(414);
      let _79_v2;
      _79_v2 = true;
      let _80_v3;
      _80_v3 = _dafny.Seq.of(_79_v2, true);
      let _81_v4;
      _81_v4 = new _dafny.CodePoint('j'.codePointAt(0));
      let _82_v5;
      _82_v5 = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,new BigNumber((_80_v3).length));
      let _83_v6;
      _83_v6 = _dafny.Map.Empty.slice().updateUnsafe(_80_v3,(((_82_v5).contains(!(true))) ? ((_82_v5).get(!(true))) : (_78_v1)));
      _78_v1 = _module.__default.safeDivisionInt(_78_v1, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_80_v3,_module.__default.fm0(_81_v4, _76_globalState))).Merge(_83_v6)).length));
      if (!(_79_v2)) {
        let _84_v7;
        let _nw10 = Array((new BigNumber(2)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _78_v1;
        _nw10[(_dafny.ONE).toNumber()] = _78_v1;
        _84_v7 = _nw10;
        let _85_v8;
        _85_v8 = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,_81_v4);
        let _86_v9;
        _86_v9 = _dafny.Map.Empty.slice().updateUnsafe((_85_v8).update(_79_v2, _81_v4),_module.__default.fm2(_79_v2, _76_globalState));
        let _87_v10;
        _87_v10 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((((_86_v9).contains(_85_v8)) ? ((_86_v9).get(_85_v8)) : (_79_v2)), _76_globalState),_84_v7);
        let _88_v11;
        _88_v11 = _dafny.MultiSet.fromElements(_82_v5);
        let _89_v12;
        _89_v12 = _module.D0.create_DC0(_84_v7);
        let _90_v13;
        let _nw11 = Array((new BigNumber(28)).toNumber());
        _nw11[(_dafny.ZERO).toNumber()] = _84_v7;
        _nw11[(_dafny.ONE).toNumber()] = _84_v7;
        _nw11[(new BigNumber(2)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(3)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(4)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(5)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(6)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(7)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(8)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(9)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(10)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(11)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(12)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(13)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(14)).toNumber()] = (((_87_v10).contains(_88_v11)) ? ((_87_v10).get(_88_v11)) : (_84_v7));
        _nw11[(new BigNumber(15)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(16)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(17)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(18)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(19)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(20)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(21)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(22)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(23)).toNumber()] = ((_79_v2) ? (_84_v7) : ((_89_v12).dtor_cf0));
        _nw11[(new BigNumber(24)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(25)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(26)).toNumber()] = _84_v7;
        _nw11[(new BigNumber(27)).toNumber()] = _84_v7;
        _90_v13 = _nw11;
        let _rhs3 = _90_v13;
        let _rhs4 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-42)), ((_91_v4) => function (_92_i0) {
          return _91_v4;
        })(_81_v4)), _dafny.Seq.Concat(_77_v0, _77_v0));
        _90_v13 = _rhs3;
        _77_v0 = _rhs4;
        (_76_globalState).f3 = _79_v2;
        _79_v2 = _79_v2;
        let _93_v14;
        _93_v14 = _module.D0.create_DC1(_81_v4);
        _93_v14 = _93_v14;
        _78_v1 = (new BigNumber(466)).plus(new BigNumber((_77_v0).length));
      } else {
        _module.__default.m0(_76_globalState);
        if ((_80_v3)[_module.__default.safeIndex(((_79_v2) ? (_78_v1) : (_78_v1)), new BigNumber((_80_v3).length))]) {
          let _94_v15;
          _94_v15 = _dafny.Set.fromElements(_81_v4, _81_v4, _81_v4, _module.__default.fm3(_79_v2, _76_globalState));
          let _95_v16;
          _95_v16 = _dafny.Seq.of(_94_v15, _94_v15);
          _95_v16 = _dafny.Seq.Concat(_95_v16, _dafny.Seq.of(_94_v15, _94_v15, _94_v15));
          _77_v0 = _77_v0;
          let _96_v17;
          let _init1 = ((_97_v1) => function (_98_i1) {
            return (_98_i1).minus(_97_v1);
          })(_78_v1);
          let _nw12 = Array((new BigNumber(10)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw12.length); _i0_1++) {
            _nw12[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _96_v17 = _nw12;
          let _index4 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_96_v17).length));
          (_96_v17)[_index4] = (_78_v1).multipliedBy(_78_v1);
          let _99_v18;
          let _nw13 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _99_v18 = _nw13;
          let _index5 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_99_v18).length));
          (_99_v18)[_index5] = _77_v0;
          let _100_v19;
          _100_v19 = _dafny.Set.fromElements(_79_v2);
          let _101_v20;
          let _nw14 = Array((new BigNumber(8)).toNumber());
          _nw14[(_dafny.ZERO).toNumber()] = _81_v4;
          _nw14[(_dafny.ONE).toNumber()] = (_77_v0)[_module.__default.safeIndex(new BigNumber((_100_v19).length), new BigNumber((_77_v0).length))];
          _nw14[(new BigNumber(2)).toNumber()] = _module.__default.fm3(_79_v2, _76_globalState);
          _nw14[(new BigNumber(3)).toNumber()] = _81_v4;
          _nw14[(new BigNumber(4)).toNumber()] = _81_v4;
          _nw14[(new BigNumber(5)).toNumber()] = _81_v4;
          _nw14[(new BigNumber(6)).toNumber()] = _81_v4;
          _nw14[(new BigNumber(7)).toNumber()] = _81_v4;
          _101_v20 = _nw14;
          let _index6 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_101_v20).length));
          (_101_v20)[_index6] = new _dafny.CodePoint('t'.codePointAt(0));
          let _index7 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_96_v17).length));
          let _index8 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_99_v18).length));
          let _index9 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_101_v20).length));
          let _rhs5 = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm4((_dafny.ZERO).minus(_78_v1), _79_v2, _76_globalState)).length));
          let _rhs6 = _77_v0;
          let _rhs7 = _81_v4;
          let _rhs8 = _81_v4;
          let _rhs9 = _module.__default.safeModuloInt(_module.__default.fm0(_81_v4, _76_globalState), _module.__default.fm0(_81_v4, _76_globalState));
          let _lhs1 = _96_v17;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_96_v17).length));
          let _lhs3 = _99_v18;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_99_v18).length));
          let _lhs5 = _101_v20;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_101_v20).length));
          _lhs1[_lhs2] = _rhs5;
          _lhs3[_lhs4] = _rhs6;
          _81_v4 = _rhs7;
          _lhs5[_lhs6] = _rhs8;
          _78_v1 = _rhs9;
          let _102_v21;
          _102_v21 = _module.D0.create_DC1(_81_v4);
          _102_v21 = _module.D0.create_DC1((_101_v20)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_101_v20).length))]);
          let _index10 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_99_v18).length));
          (_99_v18)[_index10] = _dafny.Seq.Concat((_99_v18)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_99_v18).length))], (_99_v18)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_99_v18).length))]);
        } else {
          _module.__default.m0(_76_globalState);
          let _103_v22;
          _103_v22 = _dafny.Seq.of(_module.__default.fm0(_81_v4, _76_globalState), new BigNumber(800));
          let _104_v23;
          let _nw15 = new _module.C1();
          _nw15.__ctor(_79_v2, _78_v1, new BigNumber((_103_v22).length), new BigNumber((_dafny.Seq.Concat(_77_v0, _77_v0)).length));
          _104_v23 = _nw15;
          let _105_v24;
          _105_v24 = _dafny.MultiSet.fromElements(_79_v2, (_104_v23).f12);
          let _106_v25;
          _106_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_105_v24).update((_104_v23).f12, _module.__default.abs(_78_v1))).cardinality()),_dafny.Seq.update(_77_v0, _module.__default.safeIndex(_104_v23.f13, new BigNumber((_77_v0).length)), new _dafny.CodePoint('q'.codePointAt(0))));
          let _107_v26;
          let _nw16 = new _module.C2();
          _nw16.__ctor(new BigNumber((_106_v25).length), _module.__default.safeDivisionInt(new BigNumber(852), _78_v1), (((_104_v23).f12) ? (_104_v23.f13) : (_104_v23.f13)), (_103_v22)[_module.__default.safeIndex(_78_v1, new BigNumber((_103_v22).length))]);
          _107_v26 = _nw16;
          _107_v26 = _107_v26;
          _81_v4 = _81_v4;
          (_76_globalState).f3 = !(new BigNumber(105)).isEqualTo((_107_v26).f7);
        }
        let _108_v28;
        _108_v28 = _module.D1.create_DC3(_81_v4);
        let _109_v29;
        _109_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm13(_78_v1, _76_globalState),_108_v28);
        let _110_v30;
        let _nw17 = new _module.C4();
        _nw17.__ctor(function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of (_109_v29).Keys.Elements) {
            let _111_v27 = _compr_11;
            if ((_109_v29).contains(_111_v27)) {
              _coll11.push([_111_v27,_dafny.Map.Empty.slice().updateUnsafe(_78_v1,new _dafny.CodePoint('y'.codePointAt(0)))]);
            }
          }
          return _coll11;
        }());
        _110_v30 = _nw17;
        _110_v30 = _110_v30;
        let _112_v31;
        let _nw18 = new _module.C0();
        _nw18.__ctor(_79_v2);
        _112_v31 = _nw18;
        let _113_v32;
        let _nw19 = Array((new BigNumber(12)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _112_v31;
        _nw19[(_dafny.ONE).toNumber()] = _112_v31;
        _nw19[(new BigNumber(2)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(3)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(4)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(5)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(6)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(7)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(8)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(9)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(10)).toNumber()] = _112_v31;
        _nw19[(new BigNumber(11)).toNumber()] = _112_v31;
        _113_v32 = _nw19;
        _113_v32 = _113_v32;
        _78_v1 = _module.__default.safeModuloInt(_78_v1, (_78_v1).multipliedBy((_dafny.ZERO).minus(new BigNumber((_module.__default.fm21(new BigNumber((_80_v3).length), _78_v1, _76_globalState)).length))));
      }
      _78_v1 = new BigNumber(-958);
      let _114_v34;
      _114_v34 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_77_v0).Elements) {
          let _115_v33 = _compr_12;
          if (_dafny.Seq.contains(_77_v0, _115_v33)) {
            _coll12.push([_115_v33,_79_v2]);
          }
        }
        return _coll12;
      }(),_dafny.Map.Empty.slice().updateUnsafe(_78_v1,new _dafny.CodePoint('q'.codePointAt(0))));
      let _116_v35;
      _116_v35 = _module.D7.create_DC18(_114_v34);
      let _117_v36;
      let _nw20 = new _module.C4();
      _nw20.__ctor((_module.D7.create_DC18((_116_v35).dtor_cf28)).dtor_cf28);
      _117_v36 = _nw20;
      if (_79_v2) {
        if ((_79_v2) || (_79_v2)) {
          let _118_v37;
          let _nw21 = new _module.C0();
          _nw21.__ctor(_79_v2);
          _118_v37 = _nw21;
          let _119_v38;
          _119_v38 = _dafny.Seq.of(_118_v37);
          let _120_v39;
          _120_v39 = _dafny.Map.Empty.slice().updateUnsafe((_118_v37).f11,_119_v38);
          _119_v38 = _dafny.Seq.Concat(_119_v38, (((_120_v39).contains((_118_v37).f11)) ? ((_120_v39).get((_118_v37).f11)) : (_119_v38)));
          let _121_v40;
          let _nw22 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _121_v40 = _nw22;
          let _122_v41;
          _122_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_79_v2, _76_globalState),!(false));
          let _123_v42;
          _123_v42 = _dafny.Seq.of(new BigNumber(897));
          let _124_v43;
          _124_v43 = _module.D0.create_DC1(_81_v4);
          let _125_v44;
          let _nw23 = Array((new BigNumber(22)).toNumber()).fill(false);
          _125_v44 = _nw23;
          let _126_v45;
          _126_v45 = _dafny.Seq.of(_125_v44, _125_v44, _125_v44);
          let _127_v46;
          _127_v46 = _module.D8.create_DC20(_126_v45);
          let _nw24 = Array((new BigNumber(17)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = new BigNumber(935);
          _nw24[(_dafny.ONE).toNumber()] = _78_v1;
          _nw24[(new BigNumber(2)).toNumber()] = new BigNumber(((_122_v41).update(_79_v2, (_118_v37).f11)).length);
          _nw24[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_123_v42, _123_v42)).length);
          _nw24[(new BigNumber(4)).toNumber()] = _78_v1;
          _nw24[(new BigNumber(5)).toNumber()] = (_78_v1).minus(_78_v1);
          _nw24[(new BigNumber(6)).toNumber()] = (new BigNumber(195)).minus(_78_v1);
          _nw24[(new BigNumber(7)).toNumber()] = _78_v1;
          _nw24[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm0((_124_v43).dtor_cf1, _76_globalState));
          _nw24[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_78_v1), _78_v1);
          _nw24[(new BigNumber(10)).toNumber()] = (new BigNumber((_dafny.Seq.update(_77_v0, _module.__default.safeIndex(_78_v1, new BigNumber((_77_v0).length)), _81_v4)).length)).multipliedBy(_78_v1);
          _nw24[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_79_v2,true)).length);
          _nw24[(new BigNumber(12)).toNumber()] = new BigNumber((((_79_v2) ? (_123_v42) : (_dafny.Seq.of(_78_v1, _78_v1, _78_v1)))).length);
          _nw24[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(_78_v1);
          _nw24[(new BigNumber(14)).toNumber()] = _78_v1;
          _nw24[(new BigNumber(15)).toNumber()] = new BigNumber(((_127_v46).dtor_cf31).length);
          _nw24[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(_78_v1, _78_v1);
          _121_v40 = _nw24;
          let _128_v47;
          let _nw25 = new _module.C1();
          _nw25.__ctor((_118_v37).f11, _78_v1, _78_v1, _78_v1);
          _128_v47 = _nw25;
          let _129_v48;
          _129_v48 = _dafny.Map.Empty.slice().updateUnsafe(_128_v47,(_128_v47).f12);
          let _130_v49;
          _130_v49 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,_128_v47);
          _129_v48 = (_129_v48).update((((_130_v49).contains(_module.__default.fm0(_81_v4, _76_globalState))) ? ((_130_v49).get(_module.__default.fm0(_81_v4, _76_globalState))) : (_128_v47)), _79_v2);
          let _index11 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_125_v44).length));
          (_125_v44)[_index11] = _79_v2;
          let _index12 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_125_v44).length));
          (_125_v44)[_index12] = false;
          let _131_v50;
          _131_v50 = _module.D6.create_DC17(_77_v0, _79_v2, new BigNumber((_80_v3).length), _module.__default.fm2((_125_v44)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_125_v44).length))], _76_globalState), (_125_v44)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_125_v44).length))]);
          (_128_v47).f13 = (_131_v50).dtor_cf25;
        } else {
          let _132_v51;
          let _nw26 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _132_v51 = _nw26;
          let _133_v52;
          _133_v52 = _dafny.Set.fromElements(_78_v1);
          let _134_v53;
          _134_v53 = _dafny.Seq.of(new BigNumber((_82_v5).length), _78_v1);
          let _index13 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_132_v51).length));
          (_132_v51)[_index13] = _module.__default.safeModuloInt(new BigNumber((_133_v52).length), (_134_v53)[_module.__default.safeIndex(_78_v1, new BigNumber((_134_v53).length))]);
          let _index14 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_132_v51).length));
          let _rhs10 = (_dafny.ZERO).minus((new BigNumber(937)).minus((_78_v1).multipliedBy(_78_v1)));
          let _rhs11 = _79_v2;
          let _rhs12 = _78_v1;
          let _lhs7 = _76_globalState;
          let _lhs8 = _132_v51;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_132_v51).length));
          _78_v1 = _rhs10;
          _lhs7.f1 = _rhs11;
          _lhs8[_lhs9] = _rhs12;
          _132_v51 = _132_v51;
          (_76_globalState).f3 = ((_132_v51)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_132_v51).length))]).isLessThanOrEqualTo(new BigNumber(172));
          let _135_v54;
          _135_v54 = _dafny.Set.fromElements(_79_v2);
          let _136_v55;
          _136_v55 = _dafny.Seq.of(_135_v54);
          let _137_v56;
          let _nw27 = Array((new BigNumber(13)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _135_v54;
          _nw27[(_dafny.ONE).toNumber()] = _135_v54;
          _nw27[(new BigNumber(2)).toNumber()] = _135_v54;
          _nw27[(new BigNumber(3)).toNumber()] = (_dafny.Set.fromElements(_79_v2, _79_v2, (_module.D6.create_DC17(_77_v0, _79_v2, _78_v1, (_80_v3)[_module.__default.safeIndex((_132_v51)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_132_v51).length))], new BigNumber((_80_v3).length))], !(true))).dtor_cf24)).Difference(_135_v54);
          _nw27[(new BigNumber(4)).toNumber()] = (_136_v55)[_module.__default.safeIndex((_132_v51)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_132_v51).length))], new BigNumber((_136_v55).length))];
          _nw27[(new BigNumber(5)).toNumber()] = _135_v54;
          _nw27[(new BigNumber(6)).toNumber()] = _135_v54;
          _nw27[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_79_v2, _79_v2);
          _nw27[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_79_v2, _79_v2, _79_v2);
          _nw27[(new BigNumber(9)).toNumber()] = _135_v54;
          _nw27[(new BigNumber(10)).toNumber()] = _135_v54;
          _nw27[(new BigNumber(11)).toNumber()] = _135_v54;
          _nw27[(new BigNumber(12)).toNumber()] = (_dafny.Set.fromElements(_79_v2, _79_v2)).Intersect(_135_v54);
          _137_v56 = _nw27;
          let _index15 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_137_v56).length));
          (_137_v56)[_index15] = _135_v54;
          let _138_v57;
          let _nw28 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _138_v57 = _nw28;
          let _index16 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_138_v57).length));
          (_138_v57)[_index16] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vyygsh"), _77_v0);
          let _139_v58;
          _139_v58 = _module.D7.create_DC19((_77_v0)[_module.__default.safeIndex(_78_v1, new BigNumber((_77_v0).length))], _79_v2);
          let _index17 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_137_v56).length));
          let _index18 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_138_v57).length));
          let _rhs13 = (_139_v58).dtor_cf30;
          let _rhs14 = _135_v54;
          let _rhs15 = _dafny.Seq.Concat(_77_v0, _77_v0);
          let _lhs10 = _76_globalState;
          let _lhs11 = _137_v56;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_137_v56).length));
          let _lhs13 = _138_v57;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_138_v57).length));
          _lhs10.f1 = _rhs13;
          _lhs11[_lhs12] = _rhs14;
          _lhs13[_lhs14] = _rhs15;
          let _140_v59;
          let _nw29 = new _module.C3();
          _nw29.__ctor(_78_v1, _78_v1);
          _140_v59 = _nw29;
          _140_v59 = _140_v59;
        }
        let _141_v60;
        let _nw30 = new _module.C0();
        _nw30.__ctor(_79_v2);
        _141_v60 = _nw30;
        _141_v60 = _141_v60;
        _77_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_77_v0, _77_v0), _dafny.Seq.Concat(_77_v0, _77_v0));
        _77_v0 = _77_v0;
        let _142_v61;
        let _nw31 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
        _142_v61 = _nw31;
        let _143_v62;
        _143_v62 = _dafny.Seq.of(_78_v1, _78_v1, _78_v1);
        let _144_v63;
        let _nw32 = new _module.C2();
        _nw32.__ctor(_78_v1, (_143_v62)[_module.__default.safeIndex(_78_v1, new BigNumber((_143_v62).length))], _78_v1, _78_v1);
        _144_v63 = _nw32;
        let _index19 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_142_v61).length));
        (_142_v61)[_index19] = _dafny.Map.Empty.slice().updateUnsafe(_144_v63,(_144_v63).fm7(_77_v0, _144_v63.f10, _76_globalState));
        let _145_v64;
        _145_v64 = _dafny.Map.Empty.slice().updateUnsafe(_144_v63,_144_v63.f10);
        let _index20 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_142_v61).length));
        (_142_v61)[_index20] = _145_v64;
      } else {
        _78_v1 = _78_v1;
        let _146_v65;
        _146_v65 = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.fm0(_81_v4, _76_globalState)),_78_v1));
        let _147_v66;
        _147_v66 = _module.D1.create_DC2(_79_v2);
        let _148_v67;
        _148_v67 = _dafny.Map.Empty.slice().updateUnsafe((((_82_v5).contains(true)) ? ((_82_v5).get(true)) : (_78_v1)),_78_v1);
        _146_v65 = (_146_v65).update((_147_v66).dtor_cf2, _148_v67);
        let _149_v68;
        let _nw33 = Array((new BigNumber(12)).toNumber()).fill([]);
        _149_v68 = _nw33;
        let _150_v69;
        let _nw34 = new _module.C3();
        _nw34.__ctor(new BigNumber(777), _78_v1);
        _150_v69 = _nw34;
        let _151_v70;
        let _nw35 = Array((new BigNumber(4)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = _150_v69;
        _nw35[(_dafny.ONE).toNumber()] = _150_v69;
        _nw35[(new BigNumber(2)).toNumber()] = _150_v69;
        _nw35[(new BigNumber(3)).toNumber()] = _150_v69;
        _151_v70 = _nw35;
        let _152_v71;
        _152_v71 = _module.D6.create_DC17(_77_v0, _79_v2, _78_v1, _79_v2, _79_v2);
        let _153_v72;
        _153_v72 = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,true);
        let _154_v73;
        let _nw36 = Array((new BigNumber(14)).toNumber());
        _nw36[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_78_v1);
        _nw36[(_dafny.ONE).toNumber()] = _78_v1;
        _nw36[(new BigNumber(2)).toNumber()] = _78_v1;
        _nw36[(new BigNumber(3)).toNumber()] = _78_v1;
        _nw36[(new BigNumber(4)).toNumber()] = _78_v1;
        _nw36[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_151_v70, _151_v70, _151_v70)).cardinality()));
        _nw36[(new BigNumber(6)).toNumber()] = (_150_v69).fm7((_152_v71).dtor_cf23, _78_v1, _76_globalState);
        _nw36[(new BigNumber(7)).toNumber()] = _78_v1;
        _nw36[(new BigNumber(8)).toNumber()] = (_150_v69).fm7(_77_v0, _78_v1, _76_globalState);
        _nw36[(new BigNumber(9)).toNumber()] = new BigNumber(608);
        _nw36[(new BigNumber(10)).toNumber()] = new BigNumber((_153_v72).length);
        _nw36[(new BigNumber(11)).toNumber()] = (_150_v69).fm7(_77_v0, _78_v1, _76_globalState);
        _nw36[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_82_v5).length));
        _nw36[(new BigNumber(13)).toNumber()] = _78_v1;
        _154_v73 = _nw36;
        let _index21 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_149_v68).length));
        (_149_v68)[_index21] = _154_v73;
        let _155_v74;
        _155_v74 = _dafny.MultiSet.fromElements(_79_v2);
        let _index22 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_149_v68).length));
        let _nw37 = Array((new BigNumber(10)).toNumber());
        _nw37[(_dafny.ZERO).toNumber()] = _78_v1;
        _nw37[(_dafny.ONE).toNumber()] = (_78_v1).minus(new BigNumber(591));
        _nw37[(new BigNumber(2)).toNumber()] = _78_v1;
        _nw37[(new BigNumber(3)).toNumber()] = (_78_v1).multipliedBy(new BigNumber((_155_v74).cardinality()));
        _nw37[(new BigNumber(4)).toNumber()] = (_78_v1).minus(new BigNumber((_153_v72).length));
        _nw37[(new BigNumber(5)).toNumber()] = _78_v1;
        _nw37[(new BigNumber(6)).toNumber()] = (_78_v1).multipliedBy(new BigNumber(-336));
        _nw37[(new BigNumber(7)).toNumber()] = _78_v1;
        _nw37[(new BigNumber(8)).toNumber()] = _module.__default.fm0((_77_v0)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_77_v0).length))], _76_globalState);
        _nw37[(new BigNumber(9)).toNumber()] = _78_v1;
        (_149_v68)[_index22] = _nw37;
        let _156_v75;
        _156_v75 = _module.D4.create_DC12(_78_v1, _dafny.Seq.of(_79_v2), !(!(_79_v2)));
        let _157_v76;
        _157_v76 = _module.D7.create_DC19(_81_v4, _79_v2);
        let _158_v77;
        _158_v77 = _dafny.Set.fromElements(_78_v1, (_dafny.ZERO).minus((((_155_v74).contains(_79_v2)) ? ((_155_v74).get(_79_v2)) : (new BigNumber(179)))), _78_v1);
        let _159_v78;
        let _nw38 = Array((new BigNumber(25)).toNumber());
        _nw38[(_dafny.ZERO).toNumber()] = _79_v2;
        _nw38[(_dafny.ONE).toNumber()] = _79_v2;
        _nw38[(new BigNumber(2)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(3)).toNumber()] = (_79_v2) || (false);
        _nw38[(new BigNumber(4)).toNumber()] = (_79_v2) || (_79_v2);
        _nw38[(new BigNumber(5)).toNumber()] = _dafny.Seq.IsPrefixOf(_80_v3, _80_v3);
        _nw38[(new BigNumber(6)).toNumber()] = !(_79_v2) || (_79_v2);
        _nw38[(new BigNumber(7)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(8)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(9)).toNumber()] = !(_79_v2) || ((_156_v75).dtor_cf17);
        _nw38[(new BigNumber(10)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(11)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(12)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(13)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(14)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(15)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(16)).toNumber()] = (_78_v1).isLessThan((_dafny.ZERO).minus((_156_v75).dtor_cf15));
        _nw38[(new BigNumber(17)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(18)).toNumber()] = (_157_v76).dtor_cf30;
        _nw38[(new BigNumber(19)).toNumber()] = !(!((_80_v3)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_158_v77).length)), new BigNumber((_80_v3).length))]) || (_79_v2));
        _nw38[(new BigNumber(20)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(21)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(22)).toNumber()] = _79_v2;
        _nw38[(new BigNumber(23)).toNumber()] = !(_79_v2) || (_79_v2);
        _nw38[(new BigNumber(24)).toNumber()] = !(_79_v2);
        _159_v78 = _nw38;
        let _index23 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_159_v78).length));
        (_159_v78)[_index23] = _79_v2;
        let _index24 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_159_v78).length));
        (_159_v78)[_index24] = _79_v2;
        let _index25 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_154_v73).length));
        (_154_v73)[_index25] = _78_v1;
        let _index26 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_154_v73).length));
        (_154_v73)[_index26] = ((_78_v1).plus(_78_v1)).minus(_78_v1);
      }
      (_76_globalState).f3 = _79_v2;
      (_76_globalState).f1 = _79_v2;
      let _160_v79;
      _160_v79 = _dafny.Seq.of(_78_v1);
      _160_v79 = _160_v79;
      let _161_v80;
      let _nw39 = Array((new BigNumber(16)).toNumber());
      _nw39[(_dafny.ZERO).toNumber()] = _81_v4;
      _nw39[(_dafny.ONE).toNumber()] = (_77_v0)[_module.__default.safeIndex(new BigNumber((_80_v3).length), new BigNumber((_77_v0).length))];
      _nw39[(new BigNumber(2)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
      _nw39[(new BigNumber(4)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(5)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
      _nw39[(new BigNumber(7)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(8)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(9)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(10)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(11)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(12)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(13)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(14)).toNumber()] = _81_v4;
      _nw39[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
      _161_v80 = _nw39;
      let _162_v81;
      _162_v81 = _module.D9.create_DC22(_161_v80);
      _161_v80 = (_162_v81).dtor_cf33;
      let _hi1 = ((_79_v2) ? (_78_v1) : (new BigNumber((_dafny.Set.fromElements(new BigNumber((_module.__default.fm16(_76_globalState)).length), _78_v1)).length)));
      for (let _163_i2 = _78_v1; _163_i2.isLessThan(_hi1); _163_i2 = _163_i2.plus(_dafny.ONE)) {
        if (_79_v2) {
          _78_v1 = _module.__default.safeModuloInt(_163_i2, _163_i2);
          let _164_v82;
          _164_v82 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_78_v1), _163_i2),_163_i2);
          _164_v82 = _164_v82;
          let _165_v83;
          let _nw40 = new _module.C3();
          _nw40.__ctor(_163_i2, _163_i2);
          _165_v83 = _nw40;
          _165_v83 = _165_v83;
          let _166_v84;
          let _nw41 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _166_v84 = _nw41;
          let _167_v85;
          let _168_v86;
          let _169_v87;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = (_165_v83).m4(_166_v84, _76_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _167_v85 = _out0;
          _168_v86 = _out1;
          _169_v87 = _out2;
          let _170_v88;
          let _nw42 = Array((new BigNumber(4)).toNumber());
          _nw42[(_dafny.ZERO).toNumber()] = _168_v86;
          _nw42[(_dafny.ONE).toNumber()] = _79_v2;
          _nw42[(new BigNumber(2)).toNumber()] = _168_v86;
          _nw42[(new BigNumber(3)).toNumber()] = _79_v2;
          _170_v88 = _nw42;
          let _171_v89;
          _171_v89 = _dafny.Map.Empty.slice().updateUnsafe(_81_v4,_79_v2);
          let _index27 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_170_v88).length));
          (_170_v88)[_index27] = (((_171_v89).contains(_81_v4)) ? ((_171_v89).get(_81_v4)) : (_79_v2));
          let _index28 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_170_v88).length));
          let _rhs16 = _79_v2;
          let _rhs17 = _81_v4;
          let _rhs18 = !(_168_v86) || (_168_v86);
          let _lhs15 = _170_v88;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_170_v88).length));
          let _lhs17 = _76_globalState;
          _lhs15[_lhs16] = _rhs16;
          _81_v4 = _rhs17;
          _lhs17.f3 = _rhs18;
        } else {
          let _172_v90;
          let _nw43 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _172_v90 = _nw43;
          let _index29 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_172_v90).length));
          (_172_v90)[_index29] = _module.__default.fm0(_81_v4, _76_globalState);
          let _173_v91;
          _173_v91 = _dafny.Seq.of(_77_v0);
          let _index30 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_172_v90).length));
          (_172_v90)[_index30] = (_dafny.ZERO).minus((new BigNumber((_173_v91).length)).multipliedBy(_163_i2));
          _78_v1 = ((((true) ? (_79_v2) : (_module.__default.fm2(_79_v2, _76_globalState)))) ? (_module.__default.safeDivisionInt(_78_v1, _163_i2)) : (_78_v1));
          let _index31 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_172_v90).length));
          (_172_v90)[_index31] = ((_dafny.ZERO).minus(_163_i2)).plus(_module.__default.fm0(_81_v4, _76_globalState));
          let _174_v92;
          _174_v92 = _dafny.Map.Empty.slice().updateUnsafe((_172_v90)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_172_v90).length))],_79_v2);
          let _175_v93;
          _175_v93 = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,_79_v2);
          let _176_v94;
          let _nw44 = Array((new BigNumber(7)).toNumber());
          _nw44[(_dafny.ZERO).toNumber()] = true;
          _nw44[(_dafny.ONE).toNumber()] = !(_79_v2);
          _nw44[(new BigNumber(2)).toNumber()] = _79_v2;
          _nw44[(new BigNumber(3)).toNumber()] = _79_v2;
          _nw44[(new BigNumber(4)).toNumber()] = (((_174_v92).contains(_78_v1)) ? ((_174_v92).get(_78_v1)) : (_79_v2));
          _nw44[(new BigNumber(5)).toNumber()] = (((_175_v93).contains(_79_v2)) ? ((_175_v93).get(_79_v2)) : (_79_v2));
          _nw44[(new BigNumber(6)).toNumber()] = _79_v2;
          _176_v94 = _nw44;
          let _index32 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_176_v94).length));
          (_176_v94)[_index32] = _79_v2;
          let _index33 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_176_v94).length));
          (_176_v94)[_index33] = (_80_v3)[_module.__default.safeIndex(_module.__default.safeModuloInt(_78_v1, _163_i2), new BigNumber((_80_v3).length))];
          (_76_globalState).f1 = ((_79_v2) ? (_79_v2) : ((_dafny.Set.fromElements((_176_v94)[_module.__default.safeIndex(new BigNumber(732), new BigNumber((_176_v94).length))])).IsProperSubsetOf(_dafny.Set.fromElements((_176_v94)[_module.__default.safeIndex(new BigNumber(732), new BigNumber((_176_v94).length))], true))));
        }
        let _source3 = _module.D8.create_DC21(_module.__default.safeModuloInt(new BigNumber(-625), _163_i2));
        if (_source3.is_DC21) {
          let _177___mcc_h0 = (_source3).cf32;
          let _178_cf32 = _177___mcc_h0;
          let _179_v95;
          let _init2 = ((_180_v79) => function (_181_i3) {
            return _180_v79;
          })(_160_v79);
          let _nw45 = Array((new BigNumber(11)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw45.length); _i0_2++) {
            _nw45[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _179_v95 = _nw45;
          let _index34 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_179_v95).length));
          (_179_v95)[_index34] = _160_v79;
          let _index35 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_179_v95).length));
          (_179_v95)[_index35] = _dafny.Seq.Concat(_160_v79, _160_v79);
          (_117_v36).m1(_76_globalState);
          let _182_v97;
          _182_v97 = _dafny.MultiSet.fromElements(_78_v1, _163_i2);
          let _183_v98;
          let _nw46 = new _module.C2();
          _nw46.__ctor(new BigNumber((function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of (_182_v97).Elements) {
              let _184_v96 = _compr_13;
              if ((_182_v97).contains(_184_v96)) {
                _coll13.push([_module.__default.safeDivisionInt(_184_v96, (_dafny.ZERO).minus(_163_i2)),_79_v2]);
              }
            }
            return _coll13;
          }()).length), (_dafny.ZERO).minus(_178_cf32), _178_cf32, _178_cf32);
          _183_v98 = _nw46;
          let _185_v99;
          let _nw47 = new _module.C0();
          _nw47.__ctor(_79_v2);
          _185_v99 = _nw47;
          let _186_v100;
          _186_v100 = _dafny.Seq.of(_185_v99, _185_v99);
          let _187_v101;
          _187_v101 = _module.D6.create_DC17(_77_v0, _79_v2, new BigNumber(-505), true, _79_v2);
          _82_v5 = (_82_v5).update((new BigNumber((_186_v100).length)).isLessThanOrEqualTo((_183_v98).fm7((_187_v101).dtor_cf23, _78_v1, _76_globalState)), _163_i2);
        } else {
          let _188___mcc_h1 = (_source3).cf31;
          let _189_cf31 = _188___mcc_h1;
          let _190_v102;
          _190_v102 = _dafny.MultiSet.fromElements(_81_v4, _81_v4);
          let _191_v103;
          _191_v103 = _module.D4.create_DC12(_78_v1, _80_v3, _79_v2);
          let _192_v104;
          _192_v104 = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,(_82_v5).update(false, _163_i2));
          let _193_v105;
          let _nw48 = Array((new BigNumber(27)).toNumber());
          _nw48[(_dafny.ZERO).toNumber()] = _82_v5;
          _nw48[(_dafny.ONE).toNumber()] = _82_v5;
          _nw48[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,new BigNumber(797));
          _nw48[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,_78_v1);
          _nw48[(new BigNumber(4)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,_163_i2);
          _nw48[(new BigNumber(6)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(7)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,_163_i2);
          _nw48[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,_78_v1);
          _nw48[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_78_v1);
          _nw48[(new BigNumber(11)).toNumber()] = _module.__default.fm21(new BigNumber((_82_v5).length), (((_190_v102).contains(_81_v4)) ? ((_190_v102).get(_81_v4)) : (new BigNumber((_dafny.Seq.of((_191_v103).dtor_cf15)).length))), _76_globalState);
          _nw48[(new BigNumber(12)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(13)).toNumber()] = _module.__default.fm21((_dafny.ZERO).minus(_163_i2), _78_v1, _76_globalState);
          _nw48[(new BigNumber(14)).toNumber()] = (((_192_v104).contains(_79_v2)) ? ((_192_v104).get(_79_v2)) : (_82_v5));
          _nw48[(new BigNumber(15)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(16)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(17)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(18)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(19)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(20)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(21)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(22)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(23)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(24)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(25)).toNumber()] = _82_v5;
          _nw48[(new BigNumber(26)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,_78_v1);
          _193_v105 = _nw48;
          let _194_v106;
          let _nw49 = Array((new BigNumber(11)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = _193_v105;
          _nw49[(_dafny.ONE).toNumber()] = _193_v105;
          _nw49[(new BigNumber(2)).toNumber()] = _193_v105;
          _nw49[(new BigNumber(3)).toNumber()] = _193_v105;
          _nw49[(new BigNumber(4)).toNumber()] = _193_v105;
          _nw49[(new BigNumber(5)).toNumber()] = _193_v105;
          _nw49[(new BigNumber(6)).toNumber()] = _193_v105;
          _nw49[(new BigNumber(7)).toNumber()] = _193_v105;
          _nw49[(new BigNumber(8)).toNumber()] = _193_v105;
          _nw49[(new BigNumber(9)).toNumber()] = _193_v105;
          _nw49[(new BigNumber(10)).toNumber()] = _193_v105;
          _194_v106 = _nw49;
          let _index36 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_194_v106).length));
          (_194_v106)[_index36] = _193_v105;
          let _index37 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_194_v106).length));
          (_194_v106)[_index37] = _193_v105;
          _78_v1 = new BigNumber((_77_v0).length);
          let _195_v107;
          let _init3 = function (_196_i4) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(117)), function (_197_i5) {
              return new _dafny.CodePoint('m'.codePointAt(0));
            });
          };
          let _nw50 = Array((new BigNumber(24)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw50.length); _i0_3++) {
            _nw50[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _195_v107 = _nw50;
          let _index38 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_195_v107).length));
          (_195_v107)[_index38] = _77_v0;
          let _198_v108;
          let _init4 = function (_199_i6) {
            return (_199_i6).multipliedBy(new BigNumber(281));
          };
          let _nw51 = Array((new BigNumber(29)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw51.length); _i0_4++) {
            _nw51[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _198_v108 = _nw51;
          let _index39 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_198_v108).length));
          (_198_v108)[_index39] = ((_79_v2) ? (_78_v1) : (_78_v1));
          let _index40 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_195_v107).length));
          let _index41 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_198_v108).length));
          let _rhs19 = _79_v2;
          let _rhs20 = _77_v0;
          let _rhs21 = _78_v1;
          let _rhs22 = _79_v2;
          let _lhs18 = _76_globalState;
          let _lhs19 = _195_v107;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_195_v107).length));
          let _lhs21 = _198_v108;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_198_v108).length));
          let _lhs23 = _76_globalState;
          _lhs18.f1 = _rhs19;
          _lhs19[_lhs20] = _rhs20;
          _lhs21[_lhs22] = _rhs21;
          _lhs23.f3 = _rhs22;
          let _200_v109;
          let _nw52 = new _module.C2();
          _nw52.__ctor(_163_i2, _78_v1, _163_i2, (_198_v108)[_module.__default.safeIndex(new BigNumber(352), new BigNumber((_198_v108).length))]);
          _200_v109 = _nw52;
          let _201_v110;
          _201_v110 = _dafny.Seq.of(_200_v109);
          let _index42 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_198_v108).length));
          (_198_v108)[_index42] = (_dafny.ZERO).minus(((_79_v2) ? ((_198_v108)[_module.__default.safeIndex(new BigNumber(352), new BigNumber((_198_v108).length))]) : (((_79_v2) ? (new BigNumber((_201_v110).length)) : ((_200_v109).f8)))));
        }
        (_76_globalState).f1 = (_80_v3)[_module.__default.safeIndex(_78_v1, new BigNumber((_80_v3).length))];
        if (!((_163_i2).isLessThan(_78_v1))) {
          let _202_v111;
          _202_v111 = _dafny.Map.Empty.slice().updateUnsafe(_163_i2,_163_i2);
          let _203_v112;
          _203_v112 = _dafny.Set.fromElements(true);
          let _204_v113;
          _204_v113 = _dafny.Set.fromElements(new BigNumber((_203_v112).length), _163_i2, _163_i2);
          let _205_v114;
          _205_v114 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(_79_v2)).update(_79_v2, _module.__default.abs(new BigNumber((_204_v113).length))),new BigNumber((_dafny.Seq.UnicodeFromString("rxrjsdrs")).length));
          _202_v111 = _module.__default.fm22(_78_v1, (_163_i2).minus(new BigNumber((_205_v114).length)), true, _160_v79, _76_globalState);
          let _206_v115;
          let _init5 = ((_207_v2) => function (_208_i7) {
            return _207_v2;
          })(_79_v2);
          let _nw53 = Array((new BigNumber(14)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw53.length); _i0_5++) {
            _nw53[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _206_v115 = _nw53;
          let _209_v116;
          _209_v116 = _dafny.MultiSet.fromElements(_206_v115, _206_v115, _206_v115);
          let _210_v117;
          let _nw54 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _210_v117 = _nw54;
          let _211_v118;
          _211_v118 = _dafny.Map.Empty.slice().updateUnsafe(_209_v116,_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm11(_79_v2, _163_i2, _76_globalState),_210_v117));
          let _212_v119;
          _212_v119 = _dafny.Map.Empty.slice().updateUnsafe(_77_v0,_210_v117);
          _211_v118 = (_211_v118).update((_dafny.MultiSet.fromElements(_206_v115, _206_v115)).Difference(_209_v116), _212_v119);
          let _213_v120;
          _213_v120 = _module.D0.create_DC0(_210_v117);
          let _214_v121;
          _214_v121 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(338),_79_v2));
          let _215_v122;
          _215_v122 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,_79_v2);
          let _216_v123;
          let _nw55 = new _module.C0();
          _nw55.__ctor(_79_v2);
          _216_v123 = _nw55;
          let _217_v124;
          _217_v124 = _dafny.Map.Empty.slice().updateUnsafe((((_214_v121).contains(_163_i2)) ? ((_214_v121).get(_163_i2)) : (_215_v122)),_216_v123);
          let _218_v125;
          let _nw56 = new _module.C3();
          _nw56.__ctor(new BigNumber((_217_v124).length), _78_v1);
          _218_v125 = _nw56;
          let _219_v126;
          _219_v126 = _dafny.Map.Empty.slice().updateUnsafe(_163_i2,((_79_v2) ? (_218_v125) : (_218_v125)));
          let _rhs23 = _213_v120;
          let _rhs24 = (_dafny.Map.Empty.slice().updateUnsafe((((_82_v5).contains((_216_v123).f11)) ? ((_82_v5).get((_216_v123).f11)) : ((_dafny.ZERO).minus((_218_v125).f8))),_218_v125)).Merge((_219_v126).Merge(_dafny.Map.Empty.slice().updateUnsafe(_78_v1,_218_v125)));
          let _rhs25 = _79_v2;
          let _rhs26 = (_218_v125).fm7(_77_v0, (new BigNumber((_77_v0).length)).multipliedBy((_218_v125).f7), _76_globalState);
          let _lhs24 = _76_globalState;
          _213_v120 = _rhs23;
          _219_v126 = _rhs24;
          _lhs24.f3 = _rhs25;
          _78_v1 = _rhs26;
          _215_v122 = (_215_v122).update(new BigNumber((_77_v0).length), !(_79_v2));
          let _220_v127;
          _220_v127 = _dafny.MultiSet.fromElements(false);
          let _221_v128;
          let _nw57 = new _module.C1();
          _nw57.__ctor(!(new BigNumber((_220_v127).cardinality())).isEqualTo(_78_v1), _78_v1, new BigNumber(220), (_218_v125).f7);
          _221_v128 = _nw57;
        } else {
          _79_v2 = _79_v2;
          let _222_v129;
          let _nw58 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _222_v129 = _nw58;
          let _index43 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_222_v129).length));
          (_222_v129)[_index43] = _78_v1;
          let _index44 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_222_v129).length));
          (_222_v129)[_index44] = (_dafny.ZERO).minus(_163_i2);
          let _223_v130;
          let _nw59 = Array((new BigNumber(5)).toNumber()).fill(false);
          _223_v130 = _nw59;
          let _index45 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_223_v130).length));
          (_223_v130)[_index45] = false;
          let _index46 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_223_v130).length));
          let _rhs27 = (_80_v3)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_160_v79).length), _163_i2), new BigNumber((_80_v3).length))];
          let _rhs28 = (((_222_v129)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_222_v129).length))]).minus(_78_v1)).isLessThanOrEqualTo((_222_v129)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_222_v129).length))]);
          let _lhs25 = _223_v130;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_223_v130).length));
          let _lhs27 = _76_globalState;
          _lhs25[_lhs26] = _rhs27;
          _lhs27.f3 = _rhs28;
          let _index47 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_223_v130).length));
          (_223_v130)[_index47] = (_223_v130)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_223_v130).length))];
          let _224_v131;
          _224_v131 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,(_223_v130)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_223_v130).length))]);
          (_117_v36).m2(_224_v131, _dafny.Seq.Concat(_80_v3, _80_v3), !((_223_v130)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_223_v130).length))]), _76_globalState);
        }
      }
      _78_v1 = (_dafny.ZERO).minus((_78_v1).multipliedBy((_78_v1).plus(_78_v1)));
      _79_v2 = _79_v2;
      let _225_v132;
      _225_v132 = _dafny.Set.fromElements((_77_v0)[_module.__default.safeIndex(_78_v1, new BigNumber((_77_v0).length))]);
      let _226_v134;
      _226_v134 = _dafny.MultiSet.fromElements(_225_v132, function () {
        let _coll14 = new _dafny.Set();
        for (const _compr_14 of (_225_v132).Elements) {
          let _227_v133 = _compr_14;
          if ((_225_v132).contains(_227_v133)) {
            _coll14.add(_227_v133);
          }
        }
        return _coll14;
      }(), _225_v132);
      let _228_i8;
      _228_i8 = _dafny.ZERO;
      L0: {
        while ((_226_v134).IsSubsetOf(_226_v134)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_228_i8)) {
              break L0;
            }
            _228_i8 = (_228_i8).plus(_dafny.ONE);
            let _229_v135;
            let _nw60 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
            _229_v135 = _nw60;
            let _230_v136;
            let _nw61 = new _module.C1();
            _nw61.__ctor(_79_v2, _78_v1, _78_v1, _78_v1);
            _230_v136 = _nw61;
            let _231_v137;
            _231_v137 = _module.D2.create_DC5(new BigNumber((_dafny.Set.fromElements(_79_v2)).length), _77_v0, _230_v136);
            let _232_v138;
            _232_v138 = _module.D2.create_DC6(_231_v137);
            let _233_v139;
            _233_v139 = _dafny.Map.Empty.slice().updateUnsafe(_81_v4,_232_v138);
            let _234_v140;
            _234_v140 = _dafny.Set.fromElements(_233_v139);
            let _index48 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_229_v135).length));
            (_229_v135)[_index48] = _234_v140;
            let _index49 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_229_v135).length));
            (_229_v135)[_index49] = _234_v140;
            let _235_v141;
            _235_v141 = _module.D6.create_DC17(_77_v0, _79_v2, _78_v1, _79_v2, _79_v2);
            let _236_v142;
            let _nw62 = new _module.C1();
            _nw62.__ctor(!((_235_v141).dtor_cf26), (_230_v136).f7, _78_v1, ((_230_v136).f7).plus((_230_v136).fm7(_77_v0, (_230_v136).f7, _76_globalState)));
            _236_v142 = _nw62;
            _236_v142 = _236_v142;
            let _237_v143;
            let _init6 = ((_238_v136) => function (_239_i9) {
              return (_239_i9).plus((_238_v136).f7);
            })(_230_v136);
            let _nw63 = Array((new BigNumber(2)).toNumber());
            for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw63.length); _i0_6++) {
              _nw63[_i0_6] = _init6(new BigNumber(_i0_6));
            }
            _237_v143 = _nw63;
            let _index50 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_237_v143).length));
            (_237_v143)[_index50] = (_230_v136).f7;
            let _index51 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_237_v143).length));
            (_237_v143)[_index51] = _78_v1;
            let _240_v144;
            _240_v144 = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,_79_v2);
            let _index52 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_237_v143).length));
            (_237_v143)[_index52] = (new BigNumber((_240_v144).length)).plus((_78_v1).multipliedBy((_236_v142).fm7(_dafny.Seq.UnicodeFromString("qyb"), new BigNumber((_160_v79).length), _76_globalState)));
          }
        }
      }
      let _241_v145;
      let _init7 = function (_242_i10) {
        return _module.D1.create_DC2(false);
      };
      let _nw64 = Array((new BigNumber(3)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw64.length); _i0_7++) {
        _nw64[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _241_v145 = _nw64;
      let _243_v146;
      _243_v146 = _module.D1.create_DC2(false);
      let _index53 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_241_v145).length));
      (_241_v145)[_index53] = _243_v146;
      let _index54 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_241_v145).length));
      (_241_v145)[_index54] = _module.D1.create_DC2(_79_v2);
      if (_79_v2) {
        (_76_globalState).f3 = !((_79_v2) || (_79_v2)) || (_79_v2);
        let _244_v147;
        _244_v147 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,_79_v2);
        let _245_v148;
        _245_v148 = _dafny.Map.Empty.slice().updateUnsafe(_81_v4,_80_v3);
        let _246_v149;
        let _nw65 = Array((new BigNumber(29)).toNumber()).fill(false);
        _246_v149 = _nw65;
        let _247_v150;
        _247_v150 = _dafny.Set.fromElements(_246_v149, _246_v149, _246_v149);
        (_117_v36).m2(_244_v147, (((_245_v148).contains(_81_v4)) ? ((_245_v148).get(_81_v4)) : (_80_v3)), (_247_v150).IsProperSubsetOf(_dafny.Set.fromElements(_246_v149, _246_v149)), _76_globalState);
        let _248_v151;
        _248_v151 = _module.D6.create_DC17(_dafny.Seq.Create(_module.__default.abs(new BigNumber(833)), ((_249_v4) => function (_250_i11) {
  return _249_v4;
})(_81_v4)), _79_v2, new BigNumber((_77_v0).length), _79_v2, _79_v2);
        _248_v151 = _248_v151;
        (_117_v36).m1(_76_globalState);
        let _251_v152;
        let _nw66 = new _module.C4();
        _nw66.__ctor((_117_v36).f6);
        _251_v152 = _nw66;
      } else {
        _module.__default.m0(_76_globalState);
        (_76_globalState).f1 = _79_v2;
        let _252_v153;
        _252_v153 = _dafny.Map.Empty.slice().updateUnsafe(_82_v5,!(_79_v2));
        _78_v1 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_78_v1), new BigNumber((_252_v153).length));
        let _253_v154;
        let _nw67 = Array((new BigNumber(2)).toNumber());
        _253_v154 = _nw67;
        _253_v154 = _253_v154;
        (_117_v36).m1(_76_globalState);
      }
      process.stdout.write(_dafny.toString((_76_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_76_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_76_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_77_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_78_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_79_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_80_v3, _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_81_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_82_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_83_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, true),new BigNumber(414)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v34).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-958),new _dafny.CodePoint('q'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_v35).dtor_cf28).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-958),new _dafny.CodePoint('q'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_117_v36).f6).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),true).updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-958),new _dafny.CodePoint('q'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_160_v79, _dafny.Seq.of(new BigNumber(-958)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_v81).dtor_cf33)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_225_v132).equals(_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_226_v134).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_i8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_241_v145)[_dafny.ZERO]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_241_v145)[_dafny.ONE]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_241_v145)[new BigNumber(2)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v146).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC3(cf3) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D1(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC5(cf5, cf6, cf7) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC4(cf4) {
      let $dt = new D2(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D2(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ", " + this.cf6.toVerbatimString(true) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), null);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC8(cf10) {
      let $dt = new D3(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC7(cf9) {
      let $dt = new D3(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC9(cf11) {
      let $dt = new D3(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf13, cf14) {
      let $dt = new D4(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC12(cf15, cf16, cf17) {
      let $dt = new D4(1);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC10(cf12) {
      let $dt = new D4(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC13(cf18) {
      let $dt = new D4(3);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get is_DC13() { return this.$tag === 3; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf20, cf21) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC14(cf19) {
      let $dt = new D5(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + this.cf20.toVerbatimString(true) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf23, cf24, cf25, cf26, cf27) {
      let $dt = new D6(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC16(cf22) {
      let $dt = new D6(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + this.cf23.toVerbatimString(true) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24 && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(_dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC19(cf29, cf30) {
      let $dt = new D7(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC18(cf28) {
      let $dt = new D7(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(new _dafny.CodePoint('D'.codePointAt(0)), false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf32) {
      let $dt = new D8(0);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC20(cf31) {
      let $dt = new D8(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf34, cf35, cf36) {
      let $dt = new D9(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC24(cf37, cf38) {
      let $dt = new D9(1);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC22(cf33) {
      let $dt = new D9(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf33 === other.cf33;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC23(_dafny.ZERO, _dafny.ZERO, _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = false;
      this.f3 = false;
      this._f0 = false;
      this._f2 = _dafny.ZERO;
      this._f4 = false;
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f11 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this.f13 = _dafny.ZERO;
      this._f12 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f12, f13, f7, f8) {
      let _this = this;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(175)), function (_254_i0) {
        return _dafny.MultiSet.fromElements((_this).f12);
      })).length)))).Difference((_module.D3.create_DC7(_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_this).f12)).length), (_this).f8))))).dtor_cf9)).Union((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber(-398)), _dafny.MultiSet.fromElements((_this).f8, (_this).f8), _dafny.MultiSet.fromElements((_this).f7), _dafny.MultiSet.fromElements(new BigNumber(639)), _dafny.MultiSet.fromElements((_this).f7))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f12,new _dafny.CodePoint('v'.codePointAt(0)))).length), (_this).f7, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(239)), function (_255_i1) {
        return _this.f13;
      })).length)))));
    };
    fm7(p0, p1, globalState) {
      let _this = this;
      return (_this).f7;
    };
    m6(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _module.D3.Default();
      let _hi2 = (_this).f7;
      for (let _256_i0 = p0; _256_i0.isLessThan(_hi2); _256_i0 = _256_i0.plus(_dafny.ONE)) {
        (globalState).f1 = ((_this).f7).isLessThan(p0);
        let _257_v0;
        let _nw68 = new _module.C0();
        _nw68.__ctor((_this).f12);
        _257_v0 = _nw68;
        let _258_v1;
        let _nw69 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _258_v1 = _nw69;
        let _index55 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_258_v1).length));
        (_258_v1)[_index55] = _module.__default.safeModuloInt((_this).f7, (_this).f8);
        let _259_v2;
        let _nw70 = Array((new BigNumber(26)).toNumber()).fill(_module.D3.Default());
        _259_v2 = _nw70;
        let _260_v3;
        _260_v3 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("mpakoftbq")).length), _this.f13, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), function (_261_i1) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length));
        let _262_v4;
        _262_v4 = _dafny.Seq.of(_260_v3);
        let _263_v6;
        _263_v6 = _module.D3.create_DC7(function () {
  let _coll15 = new _dafny.Set();
  for (const _compr_15 of (_262_v4).Elements) {
    let _264_v5 = _compr_15;
    if (_dafny.Seq.contains(_262_v4, _264_v5)) {
      _coll15.add(_264_v5);
    }
  }
  return _coll15;
}());
        let _index56 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_259_v2).length));
        (_259_v2)[_index56] = _263_v6;
        let _265_v7;
        _265_v7 = _dafny.MultiSet.fromElements((_this).f12, (_257_v0).f11, (_this).f12);
        let _266_v8;
        _266_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(440),_265_v7);
        let _267_v10;
        _267_v10 = _module.D4.create_DC10(function () {
  let _coll16 = new _dafny.Map();
  for (const _compr_16 of _dafny.IntegerRange(new BigNumber(674), new BigNumber(811))) {
    let _268_v9 = _compr_16;
    if (((new BigNumber(674)).isLessThanOrEqualTo(_268_v9)) && ((_268_v9).isLessThan(new BigNumber(811)))) {
      _coll16.push([(_268_v9).plus(new BigNumber(68)),_256_i0]);
    }
  }
  return _coll16;
}());
        let _269_v11;
        _269_v11 = _dafny.Set.fromElements((_this).f12);
        let _index57 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_258_v1).length));
        let _index58 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_259_v2).length));
        let _rhs29 = new BigNumber(((((_266_v8).contains(new BigNumber(((_267_v10).dtor_cf12).length))) ? ((_266_v8).get(new BigNumber(((_267_v10).dtor_cf12).length))) : (_265_v7))).cardinality());
        let _rhs30 = _module.__default.fm12((!((_257_v0).f11)) || (true), (_257_v0).f11, _module.__default.fm3((_this).f12, globalState), globalState);
        let _rhs31 = (((_this).f12) ? (!((_this).f7).isEqualTo(new BigNumber((_269_v11).length))) : ((p0).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_257_v0).f11,(_this).f8)).length))));
        let _lhs28 = _258_v1;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_258_v1).length));
        let _lhs30 = _259_v2;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_259_v2).length));
        let _lhs32 = globalState;
        _lhs28[_lhs29] = _rhs29;
        _lhs30[_lhs31] = _rhs30;
        _lhs32.f3 = _rhs31;
        _258_v1 = _258_v1;
      }
      let _270_v12;
      _270_v12 = _dafny.Seq.UnicodeFromString("nhflilcl");
      _270_v12 = _270_v12;
      let _271_v13;
      let _nw71 = new _module.C0();
      _nw71.__ctor((_this).f12);
      _271_v13 = _nw71;
      let _272_v14;
      _272_v14 = _module.D1.create_DC2((_271_v13).f11);
      let _273_v15;
      let _nw72 = new _module.C0();
      _nw72.__ctor((((_272_v14).dtor_cf2) ? ((_this).f12) : ((_this).f12)));
      _273_v15 = _nw72;
      let _274_v16;
      _274_v16 = _dafny.Set.fromElements(_module.__default.safeModuloInt(_this.f13, _this.f13), _this.f13, (_this).f7, (_this).f7);
      _274_v16 = _274_v16;
      (_this).f13 = _module.__default.safeDivisionInt(p0, (_this).fm7(_270_v12, (_this).f7, globalState));
      let _275_v17;
      _275_v17 = _dafny.Seq.of((_271_v13).f11);
      r0 = _dafny.Seq.Concat(_275_v17, _dafny.Seq.of((_this).f12));
      let _276_v18;
      _276_v18 = _module.D3.create_DC8(((_this).f8).minus(p0));
      r1 = _276_v18;
      return [r0, r1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this.f10 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f9, f10, f7, f8) {
      let _this = this;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, false, !(false), true)).length)), _dafny.MultiSet.fromElements((_this).f8))).Difference(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("ql")).length))).length),(_this).f8)).length)), _dafny.MultiSet.fromElements((_this).f8, (_this).f7, (_this).f9), _dafny.MultiSet.fromElements(new BigNumber(846), (_this).f8), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(23)), function (_277_i0) {
        return _dafny.Set.fromElements((_this).f9);
      })).length))));
    };
    fm7(p0, p1, globalState) {
      let _this = this;
      return (_this).f8;
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return ((_module.D1.create_DC2(!(false))).dtor_cf2) && (!_dafny.areEqual(_dafny.Seq.UnicodeFromString("see"), _dafny.Seq.UnicodeFromString("mu")));
    };
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(false)).IsSubsetOf((_dafny.MultiSet.fromElements(true, false)).Intersect(_dafny.MultiSet.fromElements(true)));
    };
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      (globalState).f1 = (_this.f10).isLessThanOrEqualTo((_this.f10).multipliedBy((_this).f9));
      let _278_i0;
      _278_i0 = _dafny.ZERO;
      L1: {
        while (false) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_278_i0)) {
              break L1;
            }
            _278_i0 = (_278_i0).plus(_dafny.ONE);
            let _279_v0;
            _279_v0 = true;
            if (!(_279_v0) || (true)) {
              (_this).f10 = (_this).f7;
              let _280_v1;
              let _init8 = ((_281_v0) => function (_282_i1) {
                return _281_v0;
              })(_279_v0);
              let _nw73 = Array((new BigNumber(8)).toNumber());
              for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw73.length); _i0_8++) {
                _nw73[_i0_8] = _init8(new BigNumber(_i0_8));
              }
              _280_v1 = _nw73;
              _280_v1 = _280_v1;
              let _283_v2;
              _283_v2 = new _dafny.CodePoint('e'.codePointAt(0));
              let _284_v3;
              let _nw74 = new _module.C0();
              _nw74.__ctor((_this).fm10(!(_279_v0), _279_v0, _279_v0, _283_v2, globalState));
              _284_v3 = _nw74;
              let _285_v4;
              _285_v4 = _dafny.Seq.of(_279_v0);
              let _286_v5;
              _286_v5 = _module.D1.create_DC3(_283_v2);
              let _287_v6;
              _287_v6 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f7, new BigNumber((_dafny.MultiSet.FromArray(_285_v4)).cardinality())),_286_v5);
              _287_v6 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(561),_286_v5)).Merge(_287_v6);
              let _288_v7;
              let _nw75 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
              _288_v7 = _nw75;
              _288_v7 = _288_v7;
            } else {
              let _289_v8;
              _289_v8 = _dafny.Set.fromElements((_this).f8, _this.f10);
              _289_v8 = _289_v8;
              let _290_v9;
              _290_v9 = _module.D2.create_DC4((_this).f7);
              (_this).f10 = (_290_v9).dtor_cf4;
              let _291_v10;
              _291_v10 = _dafny.Set.fromElements(_279_v0, _module.__default.fm2(_279_v0, globalState), _279_v0, _279_v0);
              let _292_v11;
              _292_v11 = _dafny.Map.Empty.slice().updateUnsafe(((!(_279_v0)) ? (_279_v0) : (_279_v0)),new BigNumber((_291_v10).length));
              let _293_v12;
              _293_v12 = new _dafny.CodePoint('w'.codePointAt(0));
              let _294_v13;
              let _init9 = ((_295_v12) => function (_296_i2) {
                return _295_v12;
              })(_293_v12);
              let _nw76 = Array((_dafny.ONE).toNumber());
              for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw76.length); _i0_9++) {
                _nw76[_i0_9] = _init9(new BigNumber(_i0_9));
              }
              _294_v13 = _nw76;
              let _297_v14;
              _297_v14 = _module.D1.create_DC2(_279_v0);
              let _rhs32 = p1;
              let _rhs33 = _279_v0;
              let _rhs34 = (((_297_v14).dtor_cf2) ? (new _dafny.CodePoint('x'.codePointAt(0))) : (_293_v12));
              let _rhs35 = _294_v13;
              let _rhs36 = _279_v0;
              let _lhs33 = globalState;
              let _lhs34 = globalState;
              _292_v11 = _rhs32;
              _lhs33.f1 = _rhs33;
              _293_v12 = _rhs34;
              _294_v13 = _rhs35;
              _lhs34.f3 = _rhs36;
              (_this).f10 = (_this).f8;
              let _298_v15;
              let _nw77 = new _module.C0();
              _nw77.__ctor(_279_v0);
              _298_v15 = _nw77;
            }
            (globalState).f3 = (_this.f10).isLessThan(((_this).f8).plus(new BigNumber(-750)));
            let _299_v16;
            _299_v16 = _module.D2.create_DC4((_this).f9);
            let _300_v17;
            _300_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,function (_pat_let0_0) {
              return function (_301_dt__update__tmp_h0) {
                return function (_pat_let1_0) {
                  return function (_302_dt__update_hcf4_h0) {
                    return _module.D2.create_DC4(_302_dt__update_hcf4_h0);
                  }(_pat_let1_0);
                }((_this).f9);
              }(_pat_let0_0);
            }(_299_v16));
            let _303_v18;
            _303_v18 = new _dafny.CodePoint('h'.codePointAt(0));
            let _304_v19;
            _304_v19 = _dafny.Seq.of(_279_v0, _279_v0, _279_v0);
            let _305_v20;
            _305_v20 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_279_v0), _304_v19, _304_v19, _304_v19);
            let _306_v21;
            let _nw78 = Array((new BigNumber(19)).toNumber());
            _nw78[(_dafny.ZERO).toNumber()] = (_this).f8;
            _nw78[(_dafny.ONE).toNumber()] = (_this).f9;
            _nw78[(new BigNumber(2)).toNumber()] = (_this).f7;
            _nw78[(new BigNumber(3)).toNumber()] = (_this).f7;
            _nw78[(new BigNumber(4)).toNumber()] = (_this).f9;
            _nw78[(new BigNumber(5)).toNumber()] = _this.f10;
            _nw78[(new BigNumber(6)).toNumber()] = (_this).f7;
            _nw78[(new BigNumber(7)).toNumber()] = (_this).f9;
            _nw78[(new BigNumber(8)).toNumber()] = _this.f10;
            _nw78[(new BigNumber(9)).toNumber()] = (_this).f7;
            _nw78[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("sdqprjgij")).length);
            _nw78[(new BigNumber(11)).toNumber()] = (_this).f8;
            _nw78[(new BigNumber(12)).toNumber()] = new BigNumber((_300_v17).length);
            _nw78[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm0(_303_v18, globalState));
            _nw78[(new BigNumber(14)).toNumber()] = new BigNumber(-826);
            _nw78[(new BigNumber(15)).toNumber()] = (_this).f9;
            _nw78[(new BigNumber(16)).toNumber()] = (_this).f7;
            _nw78[(new BigNumber(17)).toNumber()] = (_this).f7;
            _nw78[(new BigNumber(18)).toNumber()] = new BigNumber((_305_v20).cardinality());
            _306_v21 = _nw78;
            let _307_v22;
            _307_v22 = _module.D0.create_DC0(_306_v21);
            let _pat_let_tv4 = _306_v21;
            let _308_v23;
            let _nw79 = Array((new BigNumber(29)).toNumber());
            _nw79[(_dafny.ZERO).toNumber()] = function (_pat_let2_0) {
              return function (_309_dt__update__tmp_h1) {
                return function (_pat_let3_0) {
                  return function (_310_dt__update_hcf0_h0) {
                    return _module.D0.create_DC0(_310_dt__update_hcf0_h0);
                  }(_pat_let3_0);
                }(_pat_let_tv4);
              }(_pat_let2_0);
            }(_307_v22);
            _nw79[(_dafny.ONE).toNumber()] = _307_v22;
            _nw79[(new BigNumber(2)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(3)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(4)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(5)).toNumber()] = _module.D0.create_DC0((_307_v22).dtor_cf0);
            _nw79[(new BigNumber(6)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(7)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(8)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(9)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(10)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(11)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(12)).toNumber()] = _module.D0.create_DC0(_306_v21);
            _nw79[(new BigNumber(13)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(14)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(15)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(16)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(17)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(18)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(19)).toNumber()] = _module.D0.create_DC0(_306_v21);
            _nw79[(new BigNumber(20)).toNumber()] = _module.D0.create_DC0(_306_v21);
            _nw79[(new BigNumber(21)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(22)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(23)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(24)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(25)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(26)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(27)).toNumber()] = _307_v22;
            _nw79[(new BigNumber(28)).toNumber()] = _307_v22;
            _308_v23 = _nw79;
            _308_v23 = _308_v23;
            let _311_v24;
            _311_v24 = _module.D0.create_DC1(_303_v18);
            let _312_v25;
            _312_v25 = _dafny.Set.fromElements(_311_v24, _311_v24);
            let _pat_let_tv5 = _303_v18;
            let _313_v26;
            _313_v26 = _dafny.Map.Empty.slice().updateUnsafe(_279_v0,(_312_v25).IsProperSubsetOf(_dafny.Set.fromElements(_311_v24, function (_pat_let4_0) {
              return function (_314_dt__update__tmp_h2) {
                return function (_pat_let5_0) {
                  return function (_315_dt__update_hcf1_h0) {
                    return _module.D0.create_DC1(_315_dt__update_hcf1_h0);
                  }(_pat_let5_0);
                }(_pat_let_tv5);
              }(_pat_let4_0);
            }(_311_v24))));
            let _316_v27;
            let _nw80 = new _module.C1();
            _nw80.__ctor(_279_v0, _this.f10, (_this).f7, (_dafny.ZERO).minus(_this.f10));
            _316_v27 = _nw80;
            let _317_v28;
            _317_v28 = _module.D2.create_DC5((_dafny.ZERO).minus((_this).f8), _module.__default.fm11(_279_v0, _this.f10, globalState), _316_v27);
            let _318_v29;
            _318_v29 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let6_0) {
              return function (_319_dt__update__tmp_h3) {
                return function (_pat_let7_0) {
                  return function (_320_dt__update_hcf5_h0) {
                    return _module.D2.create_DC5(_320_dt__update_hcf5_h0, (_319_dt__update__tmp_h3).dtor_cf6, (_319_dt__update__tmp_h3).dtor_cf7);
                  }(_pat_let7_0);
                }(_this.f10);
              }(_pat_let6_0);
            }(_module.D2.create_DC5((_this).f9, p0, _316_v27)),_279_v0);
            if ((((_313_v26).contains(!(_318_v29).contains(_317_v28))) ? ((_313_v26).get(!(_318_v29).contains(_317_v28))) : (!(_305_v20).equals(_305_v20)))) {
              let _321_v30;
              _321_v30 = _dafny.MultiSet.fromElements((_this).f8);
              let _322_v31;
              _322_v31 = _dafny.Set.fromElements(_321_v30);
              let _323_v32;
              _323_v32 = _module.D3.create_DC9(_module.D3.create_DC7(_322_v31));
              let _324_v33;
              _324_v33 = _dafny.Seq.of(_323_v32, _323_v32, _323_v32);
              (globalState).f1 = !_dafny.Seq.contains(_324_v33, _323_v32);
              (globalState).f1 = _module.__default.fm2(((_279_v0) ? (!(_279_v0)) : (_279_v0)), globalState);
              let _325_v34;
              let _init10 = ((_326_v27) => function (_327_i3) {
                return ((_dafny.ZERO).minus((_326_v27).f7)).isLessThanOrEqualTo((_326_v27).f7);
              })(_316_v27);
              let _nw81 = Array((new BigNumber(3)).toNumber());
              for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw81.length); _i0_10++) {
                _nw81[_i0_10] = _init10(new BigNumber(_i0_10));
              }
              _325_v34 = _nw81;
              let _328_v35;
              _328_v35 = _dafny.Seq.of(_316_v27);
              let _index59 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_325_v34).length));
              (_325_v34)[_index59] = !(new BigNumber((_328_v35).length)).isEqualTo((_316_v27).f8);
              let _index60 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_325_v34).length));
              (_325_v34)[_index60] = _279_v0;
              let _index61 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_306_v21).length));
              (_306_v21)[_index61] = (_316_v27).f8;
              let _index62 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_306_v21).length));
              (_306_v21)[_index62] = _module.__default.safeModuloInt((_this).f9, (_316_v27).f8);
              (_this).f10 = ((((_325_v34)[_module.__default.safeIndex(new BigNumber(951), new BigNumber((_325_v34).length))]) ? ((_this).f8) : (_module.__default.fm0(_303_v18, globalState)))).multipliedBy((_dafny.ZERO).minus((_316_v27).f8));
            } else {
              let _329_v36;
              let _init11 = ((_330_v0) => function (_331_i4) {
                return _330_v0;
              })(_279_v0);
              let _nw82 = Array((new BigNumber(22)).toNumber());
              for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw82.length); _i0_11++) {
                _nw82[_i0_11] = _init11(new BigNumber(_i0_11));
              }
              _329_v36 = _nw82;
              let _332_v37;
              _332_v37 = _dafny.Map.Empty.slice().updateUnsafe(_279_v0,_329_v36);
              _332_v37 = _332_v37;
              (globalState).f1 = ((_316_v27).f7).isEqualTo((_this).f7);
              let _333_v38;
              _333_v38 = _dafny.Set.fromElements((_this).f9);
              let _334_v39;
              _334_v39 = _dafny.Seq.of(new BigNumber(-828), new BigNumber((_333_v38).length), (_316_v27).f8, (_316_v27).f7, new BigNumber((p0).length));
              _334_v39 = _dafny.Seq.Concat(_dafny.Seq.Concat(_334_v39, _334_v39), _334_v39);
              let _index63 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_306_v21).length));
              (_306_v21)[_index63] = (_this).f8;
              let _index64 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_306_v21).length));
              let _rhs37 = _303_v18;
              let _rhs38 = (_316_v27).f8;
              let _lhs35 = _306_v21;
              let _lhs36 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_306_v21).length));
              _303_v18 = _rhs37;
              _lhs35[_lhs36] = _rhs38;
              (globalState).f1 = _279_v0;
            }
          }
        }
      }
      let _335_v40;
      let _nw83 = Array((_dafny.ONE).toNumber()).fill([]);
      _335_v40 = _nw83;
      let _nw84 = Array((new BigNumber(3)).toNumber()).fill([]);
      _335_v40 = _nw84;
      let _336_v41;
      _336_v41 = true;
      (globalState).f1 = (_336_v41) || (_336_v41);
      if (((_336_v41) ? (false) : (_336_v41))) {
        (globalState).f1 = _336_v41;
        if (_336_v41) {
          let _337_v42;
          let _init12 = ((_338_v41) => function (_339_i5) {
            return (_338_v41) && (_338_v41);
          })(_336_v41);
          let _nw85 = Array((new BigNumber(20)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw85.length); _i0_12++) {
            _nw85[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _337_v42 = _nw85;
          _337_v42 = _337_v42;
          let _340_v43;
          let _nw86 = Array((new BigNumber(2)).toNumber()).fill(_module.D0.Default());
          _340_v43 = _nw86;
          let _341_v44;
          _341_v44 = _dafny.MultiSet.fromElements(_340_v43);
          _341_v44 = ((_dafny.MultiSet.fromElements(_340_v43, _340_v43)).Intersect(_341_v44)).update(_340_v43, _module.__default.abs(new BigNumber(319)));
          r0 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cget"), p0);
          let _index65 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_337_v42).length));
          (_337_v42)[_index65] = _336_v41;
          let _index66 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_337_v42).length));
          let _rhs39 = !(_336_v41);
          let _rhs40 = _this.f10;
          let _rhs41 = (_this).f9;
          let _rhs42 = true;
          let _rhs43 = _module.__default.fm0(new _dafny.CodePoint('l'.codePointAt(0)), globalState);
          let _lhs37 = _337_v42;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_337_v42).length));
          let _lhs39 = _this;
          let _lhs40 = _this;
          let _lhs41 = globalState;
          let _lhs42 = _this;
          _lhs37[_lhs38] = _rhs39;
          _lhs39.f10 = _rhs40;
          _lhs40.f10 = _rhs41;
          _lhs41.f1 = _rhs42;
          _lhs42.f10 = _rhs43;
          let _342_v45;
          _342_v45 = _dafny.Seq.of((_337_v42)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_337_v42).length))]);
          let _343_v46;
          _343_v46 = _dafny.Set.fromElements(_342_v45, _342_v45, _342_v45);
          _336_v41 = !(!(new BigNumber(((_343_v46).Intersect(_dafny.Set.fromElements(_342_v45))).length)).isEqualTo(_module.__default.safeModuloInt(new BigNumber(-139), new BigNumber(-421))));
        } else {
          (_this).f10 = new BigNumber((p0).length);
          let _rhs44 = _this.f10;
          let _rhs45 = _336_v41;
          let _lhs43 = _this;
          let _lhs44 = globalState;
          _lhs43.f10 = _rhs44;
          _lhs44.f1 = _rhs45;
          (globalState).f1 = _336_v41;
          let _344_v47;
          let _nw87 = Array((new BigNumber(3)).toNumber()).fill(false);
          _344_v47 = _nw87;
          let _index67 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_344_v47).length));
          (_344_v47)[_index67] = _336_v41;
          let _index68 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_344_v47).length));
          (_344_v47)[_index68] = _336_v41;
          let _345_v48;
          let _nw88 = new _module.C0();
          _nw88.__ctor(_336_v41);
          _345_v48 = _nw88;
        }
        let _346_v49;
        _346_v49 = _dafny.Seq.of(_336_v41);
        let _347_v50;
        _347_v50 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_336_v41, globalState),_346_v49);
        let _348_v51;
        _348_v51 = _dafny.Seq.of(new BigNumber((p0).length), (_this).f7, _this.f10, _this.f10, (_this).f8);
        let _349_v52;
        let _nw89 = new _module.C1();
        _nw89.__ctor(_336_v41, (_this).f8, (_this).f9, new BigNumber((_346_v49).length));
        _349_v52 = _nw89;
        let _350_v53;
        _350_v53 = _dafny.Map.Empty.slice().updateUnsafe(_336_v41,_349_v52);
        let _351_v54;
        _351_v54 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_346_v49)).cardinality()), new BigNumber((_346_v49).length));
        let _352_v55;
        _352_v55 = new _dafny.CodePoint('y'.codePointAt(0));
        let _353_v56;
        let _nw90 = Array((new BigNumber(29)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = _this.f10;
        _nw90[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_this.f10);
        _nw90[(new BigNumber(2)).toNumber()] = (new BigNumber((_346_v49).length)).plus((_this).f8);
        _nw90[(new BigNumber(3)).toNumber()] = (_this).f7;
        _nw90[(new BigNumber(4)).toNumber()] = new BigNumber((_347_v50).length);
        _nw90[(new BigNumber(5)).toNumber()] = _this.f10;
        _nw90[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_this).f8);
        _nw90[(new BigNumber(7)).toNumber()] = (_this).f7;
        _nw90[(new BigNumber(8)).toNumber()] = (_this).f7;
        _nw90[(new BigNumber(9)).toNumber()] = ((_336_v41) ? (new BigNumber((_348_v51).length)) : ((_this).f8));
        _nw90[(new BigNumber(10)).toNumber()] = (_this).fm7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(389)), function (_354_i6) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }), new BigNumber(454), globalState);
        _nw90[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f7), (_this).f9);
        _nw90[(new BigNumber(12)).toNumber()] = _this.f10;
        _nw90[(new BigNumber(13)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_350_v53).length))).minus((_this).f7);
        _nw90[(new BigNumber(14)).toNumber()] = (_349_v52).f8;
        _nw90[(new BigNumber(15)).toNumber()] = (_this).f9;
        _nw90[(new BigNumber(16)).toNumber()] = (_349_v52).f8;
        _nw90[(new BigNumber(17)).toNumber()] = (_this).f7;
        _nw90[(new BigNumber(18)).toNumber()] = (_this).f8;
        _nw90[(new BigNumber(19)).toNumber()] = _this.f10;
        _nw90[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((_this).f7);
        _nw90[(new BigNumber(21)).toNumber()] = (_this).f9;
        _nw90[(new BigNumber(22)).toNumber()] = (_this).f7;
        _nw90[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(((_this).f8).minus(new BigNumber(-46)));
        _nw90[(new BigNumber(24)).toNumber()] = (_349_v52).f8;
        _nw90[(new BigNumber(25)).toNumber()] = ((((_351_v54).contains((_this).f9)) ? ((_351_v54).get((_this).f9)) : ((_this).f8))).plus((_this).f9);
        _nw90[(new BigNumber(26)).toNumber()] = _this.f10;
        _nw90[(new BigNumber(27)).toNumber()] = _module.__default.fm0(_352_v55, globalState);
        _nw90[(new BigNumber(28)).toNumber()] = _this.f10;
        _353_v56 = _nw90;
        let _index69 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_353_v56).length));
        (_353_v56)[_index69] = new BigNumber((_348_v51).length);
        let _355_v57;
        _355_v57 = _dafny.Map.Empty.slice().updateUnsafe(_352_v55,_336_v41);
        let _356_v58;
        _356_v58 = _dafny.Seq.of(_349_v52, _349_v52);
        let _357_v59;
        _357_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(516)), ((_358_v55) => function (_359_i7) {
          return _358_v55;
        })(_352_v55))).length),_356_v58);
        let _index70 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_353_v56).length));
        (_353_v56)[_index70] = new BigNumber(((_355_v57).Merge(_module.__default.fm13(new BigNumber(((((_357_v59).contains(_this.f10)) ? ((_357_v59).get(_this.f10)) : (_356_v58))).length), globalState))).length);
        let _360_v60;
        _360_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_336_v41),_346_v49);
        let _361_v61;
        let _nw91 = new _module.C1();
        _nw91.__ctor(_336_v41, (_dafny.ZERO).minus(new BigNumber((p0).length)), _this.f10, new BigNumber((_360_v60).length));
        _361_v61 = _nw91;
        let _362_v62;
        _362_v62 = _module.D1.create_DC2(_336_v41);
        let _pat_let_tv6 = _336_v41;
        let _363_v63;
        _363_v63 = _dafny.Set.fromElements(function (_pat_let8_0) {
          return function (_364_dt__update__tmp_h4) {
            return function (_pat_let9_0) {
              return function (_365_dt__update_hcf2_h0) {
                return _module.D1.create_DC2(_365_dt__update_hcf2_h0);
              }(_pat_let9_0);
            }(_pat_let_tv6);
          }(_pat_let8_0);
        }(_362_v62));
        (globalState).f3 = (_363_v63).IsProperSubsetOf(_module.__default.fm14(_361_v61.f13, !((_361_v61).f12), globalState));
      } else {
        if (_dafny.Seq.IsPrefixOf(p0, p0)) {
          let _366_v64;
          _366_v64 = new _dafny.CodePoint('s'.codePointAt(0));
          let _367_v65;
          _367_v65 = _dafny.Map.Empty.slice().updateUnsafe(_336_v41,_366_v64);
          let _368_v66;
          _368_v66 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),(((_367_v65).contains(_336_v41)) ? ((_367_v65).get(_336_v41)) : (_366_v64)));
          let _369_v67;
          _369_v67 = _dafny.Map.Empty.slice().updateUnsafe(_368_v66,(_this).f9);
          let _370_v68;
          _370_v68 = _dafny.Seq.of(_336_v41);
          let _371_v69;
          let _nw92 = new _module.C1();
          _nw92.__ctor(!(_336_v41), (((_369_v67).contains(_368_v66)) ? ((_369_v67).get(_368_v66)) : (new BigNumber((_module.__default.fm15(globalState)).length))), (_this).f9, new BigNumber((_370_v68).length));
          _371_v69 = _nw92;
          let _372_v70;
          let _373_v71;
          let _out3;
          let _out4;
          let _outcollector1 = (_371_v69).m6(_371_v69.f13, globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _372_v70 = _out3;
          _373_v71 = _out4;
          let _374_v73;
          let _nw93 = new _module.C0();
          _nw93.__ctor(!(_336_v41));
          _374_v73 = _nw93;
          let _375_v74;
          _375_v74 = _dafny.MultiSet.fromElements(_374_v73, _374_v73);
          let _376_v75;
          _376_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm7(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("glko"), _module.__default.safeIndex(new BigNumber(283), new BigNumber((_dafny.Seq.UnicodeFromString("glko")).length)), new _dafny.CodePoint('o'.codePointAt(0))), new BigNumber(-128), globalState),_375_v74);
          let _377_v76;
          _377_v76 = _dafny.Set.fromElements((_371_v69).f12);
          let _378_v77;
          let _nw94 = Array((new BigNumber(19)).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = (_371_v69).f12;
          _nw94[(_dafny.ONE).toNumber()] = _336_v41;
          _nw94[(new BigNumber(2)).toNumber()] = (_371_v69).f12;
          _nw94[(new BigNumber(3)).toNumber()] = _336_v41;
          _nw94[(new BigNumber(4)).toNumber()] = (_371_v69).f12;
          _nw94[(new BigNumber(5)).toNumber()] = (_374_v73).f11;
          _nw94[(new BigNumber(6)).toNumber()] = _336_v41;
          _nw94[(new BigNumber(7)).toNumber()] = _336_v41;
          _nw94[(new BigNumber(8)).toNumber()] = _336_v41;
          _nw94[(new BigNumber(9)).toNumber()] = (_371_v69).f12;
          _nw94[(new BigNumber(10)).toNumber()] = !((_371_v69).f12);
          _nw94[(new BigNumber(11)).toNumber()] = (_371_v69).f12;
          _nw94[(new BigNumber(12)).toNumber()] = _336_v41;
          _nw94[(new BigNumber(13)).toNumber()] = !((_374_v73).f11);
          _nw94[(new BigNumber(14)).toNumber()] = (_371_v69).f12;
          _nw94[(new BigNumber(15)).toNumber()] = (_374_v73).f11;
          _nw94[(new BigNumber(16)).toNumber()] = false;
          _nw94[(new BigNumber(17)).toNumber()] = (_371_v69).f12;
          _nw94[(new BigNumber(18)).toNumber()] = _336_v41;
          _378_v77 = _nw94;
          let _379_v78;
          _379_v78 = _dafny.Seq.of(_378_v77);
          let _380_v79;
          _380_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_374_v73).f11);
          let _381_v80;
          let _nw95 = Array((new BigNumber(19)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = (new BigNumber((p0).length)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), ((_382_v64) => function (_383_i8) {
            return _382_v64;
          })(_366_v64))).length));
          _nw95[(_dafny.ONE).toNumber()] = (new BigNumber(((function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_dafny.Seq.of((_this).f7)).Elements) {
              let _384_v72 = _compr_17;
              if (_dafny.Seq.contains(_dafny.Seq.of((_this).f7), _384_v72)) {
                _coll17.push([_module.__default.safeDivisionInt(_384_v72, new BigNumber(950)),_this.f10]);
              }
            }
            return _coll17;
          }()).update(_371_v69.f13, _371_v69.f13)).length)).multipliedBy((_this).f7);
          _nw95[(new BigNumber(2)).toNumber()] = (_this).f9;
          _nw95[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_376_v75).contains((_this).f8)) ? ((_376_v75).get((_this).f8)) : (_375_v74)),(_this).f9)).length);
          _nw95[(new BigNumber(4)).toNumber()] = ((_336_v41) ? ((_this).f7) : (_this.f10));
          _nw95[(new BigNumber(5)).toNumber()] = (_this).f9;
          _nw95[(new BigNumber(6)).toNumber()] = (new BigNumber(76)).minus((_this).f9);
          _nw95[(new BigNumber(7)).toNumber()] = ((_this).f7).minus(_this.f10);
          _nw95[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p0, p0)).length);
          _nw95[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_371_v69.f13);
          _nw95[(new BigNumber(10)).toNumber()] = new BigNumber(-325);
          _nw95[(new BigNumber(11)).toNumber()] = (_this).f9;
          _nw95[(new BigNumber(12)).toNumber()] = ((_this).f8).multipliedBy(new BigNumber((_377_v76).length));
          _nw95[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f10)), _371_v69.f13);
          _nw95[(new BigNumber(14)).toNumber()] = _371_v69.f13;
          _nw95[(new BigNumber(15)).toNumber()] = (_this).f7;
          _nw95[(new BigNumber(16)).toNumber()] = (_this).f8;
          _nw95[(new BigNumber(17)).toNumber()] = (new BigNumber((_dafny.Seq.of(_this.f10, (_this).f8)).length)).multipliedBy(new BigNumber((_379_v78).length));
          _nw95[(new BigNumber(18)).toNumber()] = new BigNumber(((_380_v79).Merge(_380_v79)).length);
          _381_v80 = _nw95;
          let _index71 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_381_v80).length));
          (_381_v80)[_index71] = new BigNumber(793);
          let _index72 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_381_v80).length));
          (_381_v80)[_index72] = new BigNumber((_377_v76).length);
          (globalState).f1 = _336_v41;
          (_this).f10 = (_dafny.ZERO).minus(((_381_v80)[_module.__default.safeIndex(new BigNumber(530), new BigNumber((_381_v80).length))]).minus(new BigNumber(-91)));
        } else {
          (globalState).f3 = (!(_336_v41) || (_336_v41)) || (_336_v41);
          let _385_v81;
          _385_v81 = new _dafny.CodePoint('y'.codePointAt(0));
          r0 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), function (_386_i9) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }), p0), _module.__default.safeIndex((_this).f9, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), function (_387_i9) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }), p0)).length)), _385_v81), p0);
          (globalState).f1 = false;
          let _388_v82;
          let _nw96 = Array((new BigNumber(8)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = (_this).f8;
          _nw96[(_dafny.ONE).toNumber()] = (_this).f9;
          _nw96[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), function (_389_i10) {
            return (_this).f7;
          })).length);
          _nw96[(new BigNumber(3)).toNumber()] = (_this).f8;
          _nw96[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f7)).cardinality()))).length);
          _nw96[(new BigNumber(5)).toNumber()] = new BigNumber(404);
          _nw96[(new BigNumber(6)).toNumber()] = _this.f10;
          _nw96[(new BigNumber(7)).toNumber()] = (_this).f8;
          _388_v82 = _nw96;
          let _390_v83;
          _390_v83 = _dafny.Seq.of(_388_v82, _388_v82);
          let _391_v84;
          let _nw97 = Array((new BigNumber(8)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = _336_v41;
          _nw97[(_dafny.ONE).toNumber()] = _336_v41;
          _nw97[(new BigNumber(2)).toNumber()] = false;
          _nw97[(new BigNumber(3)).toNumber()] = _336_v41;
          _nw97[(new BigNumber(4)).toNumber()] = _336_v41;
          _nw97[(new BigNumber(5)).toNumber()] = _336_v41;
          _nw97[(new BigNumber(6)).toNumber()] = _336_v41;
          _nw97[(new BigNumber(7)).toNumber()] = _336_v41;
          _391_v84 = _nw97;
          let _392_v85;
          _392_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_390_v83).length),_391_v84);
          let _393_v86;
          _393_v86 = _dafny.Seq.of(new BigNumber((_392_v85).length), (_this).f9, (_this).f8);
          let _394_v87;
          _394_v87 = _dafny.Seq.of((_393_v86)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_393_v86).length))], ((_this).f7).minus((_this).f7));
          let _395_v88;
          _395_v88 = _dafny.MultiSet.fromElements(_module.__default.fm16(globalState), _module.__default.fm16(globalState), _393_v86, _394_v87);
          let _396_v89;
          _396_v89 = _dafny.Map.Empty.slice().updateUnsafe(_395_v88,(_this).f7);
          let _rhs46 = !(_336_v41);
          let _rhs47 = _dafny.Seq.update(_dafny.Seq.of(_this.f10, _this.f10, (((_396_v89).contains(_395_v88)) ? ((_396_v89).get(_395_v88)) : (_this.f10))), _module.__default.safeIndex((new BigNumber(498)).plus((_this).f8), new BigNumber((_dafny.Seq.of(_this.f10, _this.f10, (((_396_v89).contains(_395_v88)) ? ((_396_v89).get(_395_v88)) : (_this.f10)))).length)), (_this).f9);
          _336_v41 = _rhs46;
          _394_v87 = _rhs47;
          r0 = p0;
        }
        if (_336_v41) {
          let _397_v90;
          _397_v90 = _module.D2.create_DC4(_this.f10);
          let _398_v91;
          _398_v91 = _dafny.Set.fromElements(_336_v41);
          let _399_v92;
          let _nw98 = Array((new BigNumber(4)).toNumber());
          _nw98[(_dafny.ZERO).toNumber()] = (_this).f8;
          _nw98[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_module.__default.fm17(_336_v41, _this.f10, globalState)).cardinality()), (_this).f8);
          _nw98[(new BigNumber(2)).toNumber()] = (_397_v90).dtor_cf4;
          _nw98[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f10, new BigNumber((_398_v91).length)));
          _399_v92 = _nw98;
          let _rhs48 = _399_v92;
          let _rhs49 = _this.f10;
          let _lhs45 = _this;
          _399_v92 = _rhs48;
          _lhs45.f10 = _rhs49;
          (globalState).f1 = ((_dafny.Set.fromElements(_336_v41)).Intersect(_398_v91)).IsSubsetOf(_398_v91);
          (globalState).f3 = _336_v41;
          _397_v90 = _397_v90;
          (globalState).f1 = _336_v41;
        } else {
          let _400_v93;
          _400_v93 = _dafny.Seq.of(new BigNumber((p0).length));
          let _401_v94;
          _401_v94 = _dafny.Set.fromElements(p0);
          let _402_v95;
          _402_v95 = _dafny.Seq.of(_336_v41);
          let _403_v96;
          let _nw99 = Array((new BigNumber(2)).toNumber()).fill(false);
          _403_v96 = _nw99;
          let _404_v97;
          _404_v97 = _dafny.Map.Empty.slice().updateUnsafe(_336_v41,_403_v96);
          let _405_v98;
          _405_v98 = _dafny.Seq.of(_403_v96);
          let _406_v99;
          _406_v99 = _dafny.MultiSet.fromElements(_404_v97, (_404_v97).update(_336_v41, (_405_v98)[_module.__default.safeIndex(new BigNumber(-514), new BigNumber((_405_v98).length))]), _404_v97, _dafny.Map.Empty.slice().updateUnsafe(false,_403_v96), _404_v97);
          let _407_v100;
          let _nw100 = Array((new BigNumber(16)).toNumber());
          _nw100[(_dafny.ZERO).toNumber()] = _dafny.Seq.IsProperPrefixOf(_module.__default.fm16(globalState), _400_v93);
          _nw100[(_dafny.ONE).toNumber()] = _336_v41;
          _nw100[(new BigNumber(2)).toNumber()] = _336_v41;
          _nw100[(new BigNumber(3)).toNumber()] = _336_v41;
          _nw100[(new BigNumber(4)).toNumber()] = (_dafny.Set.fromElements(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(93)), function (_408_i11) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          }))).IsSubsetOf(_401_v94);
          _nw100[(new BigNumber(5)).toNumber()] = !((_this).f7).isEqualTo((_this).f8);
          _nw100[(new BigNumber(6)).toNumber()] = ((_336_v41) ? (_336_v41) : (true));
          _nw100[(new BigNumber(7)).toNumber()] = (_336_v41) && (_336_v41);
          _nw100[(new BigNumber(8)).toNumber()] = _336_v41;
          _nw100[(new BigNumber(9)).toNumber()] = ((_module.__default.fm2(_336_v41, globalState)) ? (_336_v41) : (_336_v41));
          _nw100[(new BigNumber(10)).toNumber()] = _336_v41;
          _nw100[(new BigNumber(11)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_402_v95, _402_v95);
          _nw100[(new BigNumber(12)).toNumber()] = _336_v41;
          _nw100[(new BigNumber(13)).toNumber()] = ((_this).f9).isLessThan((((_406_v99).contains(_404_v97)) ? ((_406_v99).get(_404_v97)) : ((_this).f8)));
          _nw100[(new BigNumber(14)).toNumber()] = _336_v41;
          _nw100[(new BigNumber(15)).toNumber()] = _336_v41;
          _407_v100 = _nw100;
          let _index73 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_403_v96).length));
          (_403_v96)[_index73] = !(_336_v41);
          let _409_v101;
          _409_v101 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(p0, p0),(_this).f7);
          let _410_v102;
          _410_v102 = new _dafny.CodePoint('l'.codePointAt(0));
          let _411_v103;
          let _nw101 = Array((new BigNumber(10)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = _410_v102;
          _nw101[(_dafny.ONE).toNumber()] = _410_v102;
          _nw101[(new BigNumber(2)).toNumber()] = _410_v102;
          _nw101[(new BigNumber(3)).toNumber()] = _410_v102;
          _nw101[(new BigNumber(4)).toNumber()] = _410_v102;
          _nw101[(new BigNumber(5)).toNumber()] = _410_v102;
          _nw101[(new BigNumber(6)).toNumber()] = ((_336_v41) ? (_410_v102) : (_410_v102));
          _nw101[(new BigNumber(7)).toNumber()] = _410_v102;
          _nw101[(new BigNumber(8)).toNumber()] = _module.__default.fm3(_336_v41, globalState);
          _nw101[(new BigNumber(9)).toNumber()] = _410_v102;
          _411_v103 = _nw101;
          let _index74 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_403_v96).length));
          let _rhs50 = _336_v41;
          let _rhs51 = _409_v101;
          let _rhs52 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f8)).length);
          let _rhs53 = _411_v103;
          let _rhs54 = _403_v96;
          let _lhs46 = _403_v96;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_403_v96).length));
          let _lhs48 = _this;
          _lhs46[_lhs47] = _rhs50;
          _409_v101 = _rhs51;
          _lhs48.f10 = _rhs52;
          _411_v103 = _rhs53;
          _403_v96 = _rhs54;
          let _412_v104;
          _412_v104 = _dafny.Set.fromElements(false, false);
          let _413_v105;
          let _nw102 = new _module.C1();
          _nw102.__ctor(true, new BigNumber((_412_v104).length), ((_this).f7).minus(_this.f10), ((!(_336_v41)) ? ((_this).f9) : ((_dafny.ZERO).minus((_this).f8))));
          _413_v105 = _nw102;
          let _414_v106;
          let _nw103 = Array((new BigNumber(15)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = _412_v104;
          _nw103[(_dafny.ONE).toNumber()] = (_412_v104).Union(_412_v104);
          _nw103[(new BigNumber(2)).toNumber()] = _412_v104;
          _nw103[(new BigNumber(3)).toNumber()] = _412_v104;
          _nw103[(new BigNumber(4)).toNumber()] = _412_v104;
          _nw103[(new BigNumber(5)).toNumber()] = (((_403_v96)[_module.__default.safeIndex(new BigNumber(867), new BigNumber((_403_v96).length))]) ? (_412_v104) : (_412_v104));
          _nw103[(new BigNumber(6)).toNumber()] = (_412_v104).Difference(_412_v104);
          _nw103[(new BigNumber(7)).toNumber()] = _412_v104;
          _nw103[(new BigNumber(8)).toNumber()] = (((_403_v96)[_module.__default.safeIndex(new BigNumber(867), new BigNumber((_403_v96).length))]) ? (_412_v104) : (_412_v104));
          _nw103[(new BigNumber(9)).toNumber()] = _412_v104;
          _nw103[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements((_413_v105).f12);
          _nw103[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements((_413_v105).f12, _336_v41, (_413_v105).f12, false, _336_v41);
          _nw103[(new BigNumber(12)).toNumber()] = _412_v104;
          _nw103[(new BigNumber(13)).toNumber()] = (_412_v104).Difference(_412_v104);
          _nw103[(new BigNumber(14)).toNumber()] = _412_v104;
          _414_v106 = _nw103;
          let _index75 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_414_v106).length));
          (_414_v106)[_index75] = _412_v104;
          let _index76 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_414_v106).length));
          let _rhs55 = _412_v104;
          let _rhs56 = _407_v100;
          let _rhs57 = (p0)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), ((_415_v105) => function (_416_i12) {
            return _415_v105.f13;
          })(_413_v105))).length), new BigNumber((p0).length))];
          let _lhs49 = _414_v106;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_414_v106).length));
          _lhs49[_lhs50] = _rhs55;
          _407_v100 = _rhs56;
          _410_v102 = _rhs57;
          r0 = p0;
          (globalState).f1 = false;
        }
        (_this).f10 = _this.f10;
        let _417_v107;
        let _nw104 = Array((new BigNumber(6)).toNumber()).fill(false);
        _417_v107 = _nw104;
        let _418_v108;
        _418_v108 = new _dafny.CodePoint('i'.codePointAt(0));
        let _419_v109;
        _419_v109 = _module.D4.create_DC11((_this).f7, _418_v108);
        let _index77 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_417_v107).length));
        (_417_v107)[_index77] = !((_419_v109).dtor_cf13).isEqualTo(new BigNumber(284));
        let _420_v110;
        _420_v110 = _dafny.MultiSet.fromElements((_this).fm9(_336_v41, _this.f10, globalState), _336_v41);
        let _421_v111;
        _421_v111 = _dafny.Seq.of(_336_v41, _336_v41);
        let _422_v112;
        _422_v112 = _dafny.Seq.of(_420_v110, _dafny.MultiSet.fromElements(_336_v41));
        let _423_v113;
        let _nw105 = Array((new BigNumber(12)).toNumber());
        _nw105[(_dafny.ZERO).toNumber()] = _420_v110;
        _nw105[(_dafny.ONE).toNumber()] = (_420_v110).Intersect(_420_v110);
        _nw105[(new BigNumber(2)).toNumber()] = _420_v110;
        _nw105[(new BigNumber(3)).toNumber()] = _420_v110;
        _nw105[(new BigNumber(4)).toNumber()] = (_dafny.MultiSet.FromArray(_421_v111)).update(_336_v41, _module.__default.abs((_this).f9));
        _nw105[(new BigNumber(5)).toNumber()] = ((_336_v41) ? (_420_v110) : (_420_v110));
        _nw105[(new BigNumber(6)).toNumber()] = _420_v110;
        _nw105[(new BigNumber(7)).toNumber()] = _420_v110;
        _nw105[(new BigNumber(8)).toNumber()] = _420_v110;
        _nw105[(new BigNumber(9)).toNumber()] = _420_v110;
        _nw105[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(_336_v41, _336_v41, _336_v41);
        _nw105[(new BigNumber(11)).toNumber()] = (_422_v112)[_module.__default.safeIndex((_this).f7, new BigNumber((_422_v112).length))];
        _423_v113 = _nw105;
        let _index78 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_423_v113).length));
        (_423_v113)[_index78] = _dafny.MultiSet.fromElements(_336_v41);
        let _424_v114;
        _424_v114 = _dafny.Set.fromElements(_418_v108, _418_v108);
        let _425_v115;
        _425_v115 = _dafny.Map.Empty.slice().updateUnsafe(_424_v114,_this.f10);
        let _index79 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_417_v107).length));
        let _index80 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_423_v113).length));
        let _rhs58 = ((true) ? (_336_v41) : (!((_425_v115).equals(_425_v115))));
        let _rhs59 = _420_v110;
        let _rhs60 = _module.__default.safeDivisionInt(((_this).f8).multipliedBy((_dafny.ZERO).minus(_this.f10)), (_this).f7);
        let _lhs51 = _417_v107;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_417_v107).length));
        let _lhs53 = _423_v113;
        let _lhs54 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_423_v113).length));
        let _lhs55 = _this;
        _lhs51[_lhs52] = _rhs58;
        _lhs53[_lhs54] = _rhs59;
        _lhs55.f10 = _rhs60;
        let _426_v116;
        _426_v116 = _dafny.Set.fromElements((_this).f9, new BigNumber(658), new BigNumber(163));
        let _427_v117;
        _427_v117 = _dafny.Map.Empty.slice().updateUnsafe(_426_v116,true);
        let _428_v118;
        _428_v118 = _dafny.Map.Empty.slice().updateUnsafe((_417_v107)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_417_v107).length))],_427_v117);
        let _429_v119;
        _429_v119 = _dafny.Map.Empty.slice().updateUnsafe((((_428_v118).contains(_336_v41)) ? ((_428_v118).get(_336_v41)) : (_427_v117)),_336_v41);
        let _index81 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_417_v107).length));
        let _rhs61 = (_417_v107)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_417_v107).length))];
        let _rhs62 = ((_336_v41) ? (_417_v107) : (_417_v107));
        let _rhs63 = false;
        let _rhs64 = (((_429_v119).contains(_427_v117)) ? ((_429_v119).get(_427_v117)) : ((_417_v107)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_417_v107).length))]));
        let _lhs56 = _417_v107;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_417_v107).length));
        _336_v41 = _rhs61;
        _417_v107 = _rhs62;
        _336_v41 = _rhs63;
        _lhs56[_lhs57] = _rhs64;
      }
      let _430_v120;
      _430_v120 = _dafny.Seq.of((_this).f8);
      let _431_v121;
      _431_v121 = new _dafny.CodePoint('q'.codePointAt(0));
      r0 = _dafny.Seq.Concat(_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_430_v120).length), new BigNumber((p0).length)), _431_v121), _dafny.Seq.UnicodeFromString("vhgenyj"));
      r0 = p0;
      return r0;
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    __ctor(f7, f8) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length), (_this).f7, (_dafny.ZERO).minus((_this).f8)), _dafny.MultiSet.fromElements((_this).f8, (_this).f8), _dafny.MultiSet.fromElements((_this).f7, (_this).f7, (_this).f8, (_this).f7))).Difference(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("qt"))).length)), _dafny.MultiSet.fromElements(new BigNumber(670))));
    };
    fm7(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).Union((_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)))).Union(_dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)))))).cardinality());
    };
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.ZERO;
      let r3 = false;
      r2 = _module.__default.safeDivisionInt(p3, (_this).f8);
      let _432_v0;
      let _nw106 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _432_v0 = _nw106;
      _432_v0 = _432_v0;
      let _433_v1;
      _433_v1 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f8, new BigNumber((p0).length))).length), ((_this).f8).multipliedBy(p3));
      let _rhs65 = _433_v1;
      let _rhs66 = p1;
      _433_v1 = _rhs65;
      r3 = _rhs66;
      let _434_v2;
      _434_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length),(_dafny.ZERO).minus(new BigNumber((p0).length)));
      _434_v2 = (_434_v2).update((_dafny.ZERO).minus((_this).f8), (_this).f7);
      let _435_v3;
      _435_v3 = _dafny.Seq.of(!(p1), true);
      let _436_v4;
      _436_v4 = _dafny.MultiSet.fromElements(_module.D1.create_DC2(p1), p2);
      _435_v3 = _module.__default.fm8(_433_v1, new BigNumber(-163), p1, _436_v4, globalState);
      let _437_v5;
      _437_v5 = new _dafny.CodePoint('c'.codePointAt(0));
      let _438_v6;
      _438_v6 = _module.D1.create_DC3(_437_v5);
      let _source4 = ((_module.__default.fm2(p1, globalState)) ? (function (_pat_let10_0) {
        return function (_439_dt__update__tmp_h0) {
          return function (_pat_let11_0) {
            return function (_440_dt__update_hcf3_h0) {
              return _module.D1.create_DC3(_440_dt__update_hcf3_h0);
            }(_pat_let11_0);
          }(new _dafny.CodePoint('f'.codePointAt(0)));
        }(_pat_let10_0);
      }(_438_v6)) : (_438_v6));
      if (_source4.is_DC3) {
        let _441___mcc_h0 = (_source4).cf3;
        let _442_cf3 = _441___mcc_h0;
        let _443_v7;
        _443_v7 = _dafny.MultiSet.fromElements(p1);
        let _444_v8;
        _444_v8 = _dafny.Seq.of(_443_v7, _443_v7);
        _444_v8 = _444_v8;
        if (p1) {
          let _445_v9;
          let _nw107 = new _module.C2();
          _nw107.__ctor(p3, p3, _module.__default.fm0(_module.__default.fm3(false, globalState), globalState), _module.__default.safeDivisionInt((_dafny.ZERO).minus(p3), (_this).f7));
          _445_v9 = _nw107;
          (_445_v9).f10 = ((new BigNumber((_dafny.Seq.UnicodeFromString("r")).length)).minus(new BigNumber((_dafny.Seq.update(_435_v3, _module.__default.safeIndex((_this).f8, new BigNumber((_435_v3).length)), p1)).length))).multipliedBy(((p1) ? (new BigNumber(761)) : ((_445_v9).f9)));
          let _446_v10;
          let _nw108 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _446_v10 = _nw108;
          let _index82 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_446_v10).length));
          (_446_v10)[_index82] = _445_v9.f10;
          let _index83 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_446_v10).length));
          (_446_v10)[_index83] = _module.__default.fm0(_442_cf3, globalState);
          let _447_v11;
          _447_v11 = _module.D2.create_DC4((_this).f8);
          let _448_v12;
          _448_v12 = _module.D2.create_DC6(_447_v11);
          let _449_v13;
          _449_v13 = _module.D2.create_DC6(_448_v12);
          let _450_v14;
          _450_v14 = _module.D2.create_DC6(_447_v11);
          let _451_v15;
          _451_v15 = _module.D2.create_DC6(_450_v14);
          let _452_v16;
          _452_v16 = _dafny.Map.Empty.slice().updateUnsafe(_451_v15,(_this).f7);
          let _453_v17;
          _453_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(_450_v14),new BigNumber(-176)));
          let _454_v18;
          _454_v18 = _dafny.Map.Empty.slice().updateUnsafe(_442_cf3,p1);
          let _455_v19;
          _455_v19 = _dafny.Map.Empty.slice().updateUnsafe((_452_v16).Merge((((_453_v17).contains(p1)) ? ((_453_v17).get(p1)) : (_452_v16))),(((_454_v18).contains(_442_cf3)) ? ((_454_v18).get(_442_cf3)) : (p1)));
          _455_v19 = _dafny.Map.Empty.slice().updateUnsafe(_452_v16,p1);
          let _456_v20;
          let _nw109 = new _module.C0();
          _nw109.__ctor(p1);
          _456_v20 = _nw109;
        } else {
          let _457_v21;
          _457_v21 = _dafny.Map.Empty.slice().updateUnsafe((false) === (false),p1);
          _457_v21 = (_457_v21).update(p1, p1);
          let _458_v22;
          let _nw110 = new _module.C1();
          _nw110.__ctor(p1, _module.__default.fm0(_442_cf3, globalState), (_this).f7, (_this).f7);
          _458_v22 = _nw110;
          let _459_v23;
          _459_v23 = _module.D2.create_DC5((_this).f7, p0, _458_v22);
          let _460_v24;
          _460_v24 = _module.D2.create_DC6(_459_v23);
          let _461_v25;
          _461_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_dafny.Map.Empty.slice().updateUnsafe(p3,_460_v24));
          _461_v25 = _461_v25;
          let _462_v26;
          let _nw111 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
          _462_v26 = _nw111;
          _462_v26 = (_module.D0.create_DC0(_462_v26)).dtor_cf0;
          let _463_v27;
          _463_v27 = _module.D0.create_DC0(_462_v26);
          let _464_v28;
          _464_v28 = _module.D0.create_DC0((_463_v27).dtor_cf0);
          _463_v27 = _463_v27;
          r2 = _module.__default.safeDivisionInt(new BigNumber(486), (_458_v22).f7);
        }
        let _index84 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_432_v0).length));
        (_432_v0)[_index84] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hql"), p0);
        let _465_v29;
        let _nw112 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _465_v29 = _nw112;
        let _466_v30;
        _466_v30 = _dafny.Set.fromElements(_465_v29, _465_v29);
        let _index85 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_465_v29).length));
        (_465_v29)[_index85] = _module.__default.safeDivisionInt(new BigNumber((_466_v30).length), p3);
        let _467_v32;
        _467_v32 = _dafny.Set.fromElements(_442_cf3);
        let _468_v33;
        _468_v33 = _dafny.Set.fromElements(new BigNumber((_467_v32).length));
        let _index86 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_432_v0).length));
        let _index87 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_465_v29).length));
        let _rhs67 = _module.__default.fm2(p1, globalState);
        let _rhs68 = _dafny.Seq.Concat(p0, _module.__default.fm11(p1, new BigNumber((p0).length), globalState));
        let _rhs69 = p3;
        let _rhs70 = new BigNumber((((_434_v2).Merge(_434_v2)).Merge(function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of (_468_v33).Elements) {
            let _469_v31 = _compr_18;
            if ((_468_v33).contains(_469_v31)) {
              _coll18.push([(_469_v31).minus((_this).f7),(_module.D2.create_DC4(p3)).dtor_cf4]);
            }
          }
          return _coll18;
        }())).length);
        let _lhs58 = globalState;
        let _lhs59 = _432_v0;
        let _lhs60 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_432_v0).length));
        let _lhs61 = _465_v29;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_465_v29).length));
        _lhs58.f1 = _rhs67;
        _lhs59[_lhs60] = _rhs68;
        _lhs61[_lhs62] = _rhs69;
        r0 = _rhs70;
        _module.__default.m0(globalState);
      } else {
        let _470___mcc_h1 = (_source4).cf2;
        let _471_cf2 = _470___mcc_h1;
        let _472_v34;
        let _nw113 = Array((new BigNumber(28)).toNumber()).fill(false);
        _472_v34 = _nw113;
        let _index88 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_472_v34).length));
        (_472_v34)[_index88] = p1;
        let _index89 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_472_v34).length));
        (_472_v34)[_index89] = _471_cf2;
        let _473_v35;
        let _nw114 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _473_v35 = _nw114;
        let _index90 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_473_v35).length));
        (_473_v35)[_index90] = _module.__default.safeModuloInt(_module.__default.fm0(new _dafny.CodePoint('d'.codePointAt(0)), globalState), p3);
        let _index91 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_473_v35).length));
        (_473_v35)[_index91] = (p3).minus((_dafny.ZERO).minus((_this).f7));
        (globalState).f3 = _471_cf2;
        _471_cf2 = _dafny.Seq.IsPrefixOf(_433_v1, _dafny.Seq.Concat(_433_v1, _433_v1));
      }
      r0 = _module.__default.safeDivisionInt(new BigNumber((_435_v3).length), (_this).f7);
      let _474_v36;
      _474_v36 = _module.D4.create_DC11((_this).f7, _437_v5);
      let _475_v37;
      _475_v37 = _module.D4.create_DC13(_474_v36);
      let _pat_let_tv7 = _437_v5;
      let _pat_let_tv8 = p0;
      let _pat_let_tv9 = _437_v5;
      let _pat_let_tv10 = _437_v5;
      r1 = function (_source5) {
        if (_source5.is_DC11) {
          let _476___mcc_h2 = (_source5).cf13;
          let _477___mcc_h3 = (_source5).cf14;
          let _478_cf14 = _477___mcc_h3;
          let _479_cf13 = _476___mcc_h2;
          return _dafny.Seq.of(_478_cf14, _pat_let_tv7);
        } else if (_source5.is_DC12) {
          let _480___mcc_h4 = (_source5).cf15;
          let _481___mcc_h5 = (_source5).cf16;
          let _482___mcc_h6 = (_source5).cf17;
          let _483_cf17 = _482___mcc_h6;
          let _484_cf16 = _481___mcc_h5;
          let _485_cf15 = _480___mcc_h4;
          return _pat_let_tv8;
        } else if (_source5.is_DC10) {
          let _486___mcc_h7 = (_source5).cf12;
          let _487_cf12 = _486___mcc_h7;
          return _dafny.Seq.of(_pat_let_tv9);
        } else {
          let _488___mcc_h8 = (_source5).cf18;
          let _489_cf18 = _488___mcc_h8;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), ((_490_v5) => function (_491_i0) {
            return _490_v5;
          })(_pat_let_tv10));
        }
      }(_475_v37);
      r2 = (_this).f8;
      r3 = p1;
      return [r0, r1, r2, r3];
    }
    m4(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let _492_v0;
      _492_v0 = false;
      let _493_v1;
      _493_v1 = _dafny.Seq.UnicodeFromString("xeoexah");
      let _494_v2;
      let _nw115 = new _module.C1();
      _nw115.__ctor(_492_v0, (_this).f7, new BigNumber((_493_v1).length), (_this).f7);
      _494_v2 = _nw115;
      let _495_v3;
      _495_v3 = _dafny.Seq.of(_494_v2);
      _495_v3 = _495_v3;
      let _496_v4;
      _496_v4 = new BigNumber(96);
      _496_v4 = _module.__default.safeDivisionInt(_496_v4, new BigNumber(-305));
      let _497_v5;
      let _nw116 = Array((new BigNumber(6)).toNumber()).fill([]);
      _497_v5 = _nw116;
      let _498_v6;
      let _nw117 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
      _498_v6 = _nw117;
      let _index92 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_497_v5).length));
      (_497_v5)[_index92] = _498_v6;
      let _index93 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_497_v5).length));
      (_497_v5)[_index93] = _498_v6;
      let _hi3 = _496_v4;
      for (let _499_i0 = _496_v4; _499_i0.isLessThan(_hi3); _499_i0 = _499_i0.plus(_dafny.ONE)) {
        let _500_v7;
        _500_v7 = _dafny.Map.Empty.slice().updateUnsafe(_492_v0,_492_v0);
        let _501_v8;
        let _nw118 = Array((new BigNumber(25)).toNumber());
        _nw118[(_dafny.ZERO).toNumber()] = false;
        _nw118[(_dafny.ONE).toNumber()] = _492_v0;
        _nw118[(new BigNumber(2)).toNumber()] = false;
        _nw118[(new BigNumber(3)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(4)).toNumber()] = (((_500_v7).contains((((_500_v7).contains(_492_v0)) ? ((_500_v7).get(_492_v0)) : (_492_v0)))) ? ((_500_v7).get((((_500_v7).contains(_492_v0)) ? ((_500_v7).get(_492_v0)) : (_492_v0)))) : (_492_v0));
        _nw118[(new BigNumber(5)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(6)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(7)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(8)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(9)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(10)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(11)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(12)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(13)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(14)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(15)).toNumber()] = false;
        _nw118[(new BigNumber(16)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(17)).toNumber()] = true;
        _nw118[(new BigNumber(18)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(19)).toNumber()] = !(_492_v0);
        _nw118[(new BigNumber(20)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(21)).toNumber()] = false;
        _nw118[(new BigNumber(22)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(23)).toNumber()] = _492_v0;
        _nw118[(new BigNumber(24)).toNumber()] = _492_v0;
        _501_v8 = _nw118;
        let _502_v9;
        _502_v9 = _dafny.Map.Empty.slice().updateUnsafe(_492_v0,_501_v8);
        let _503_v10;
        _503_v10 = _module.D5.create_DC14(_501_v8);
        let _504_v11;
        let _nw119 = Array((new BigNumber(24)).toNumber());
        _nw119[(_dafny.ZERO).toNumber()] = _501_v8;
        _nw119[(_dafny.ONE).toNumber()] = _501_v8;
        _nw119[(new BigNumber(2)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(3)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(4)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(5)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(6)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(7)).toNumber()] = (((_502_v9).contains(_492_v0)) ? ((_502_v9).get(_492_v0)) : (_501_v8));
        _nw119[(new BigNumber(8)).toNumber()] = (_503_v10).dtor_cf19;
        _nw119[(new BigNumber(9)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(10)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(11)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(12)).toNumber()] = (_503_v10).dtor_cf19;
        _nw119[(new BigNumber(13)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(14)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(15)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(16)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(17)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(18)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(19)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(20)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(21)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(22)).toNumber()] = _501_v8;
        _nw119[(new BigNumber(23)).toNumber()] = _501_v8;
        _504_v11 = _nw119;
        let _505_v12;
        _505_v12 = new _dafny.CodePoint('y'.codePointAt(0));
        let _506_v13;
        _506_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Seq.update(_493_v1, _module.__default.safeIndex(new BigNumber(-796), new BigNumber((_493_v1).length)), _505_v12)).length));
        let _507_v14;
        _507_v14 = _dafny.Set.fromElements(_493_v1, _493_v1);
        let _rhs71 = _504_v11;
        let _rhs72 = ((_dafny.ZERO).minus(new BigNumber((_506_v13).length))).minus(_module.__default.safeModuloInt((_494_v2).f8, new BigNumber((_507_v14).length)));
        let _rhs73 = (_494_v2).f8;
        _504_v11 = _rhs71;
        _496_v4 = _rhs72;
        _496_v4 = _rhs73;
        let _508_v15;
        _508_v15 = _dafny.Map.Empty.slice().updateUnsafe((_494_v2).f8,_dafny.Seq.of(p0));
        _508_v15 = (_508_v15).update((_494_v2).f8, _dafny.Seq.of(p0));
        (globalState).f3 = _module.__default.fm2(_492_v0, globalState);
        _501_v8 = ((_492_v0) ? (_501_v8) : (_501_v8));
      }
      let _509_v16;
      _509_v16 = new _dafny.CodePoint('s'.codePointAt(0));
      let _510_v17;
      let _nw120 = new _module.C2();
      _nw120.__ctor(_module.__default.fm0(_509_v16, globalState), _496_v4, (_494_v2).f8, new BigNumber(250));
      _510_v17 = _nw120;
      let _511_v18;
      let _nw121 = new _module.C2();
      _nw121.__ctor((_494_v2).f8, _496_v4, (_494_v2).f8, (_this).f7);
      _511_v18 = _nw121;
      r0 = ((_492_v0) ? (((_492_v0) ? (_498_v6) : ((_497_v5)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_497_v5).length))]))) : ((_497_v5)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_497_v5).length))]));
      let _512_v19;
      _512_v19 = _dafny.MultiSet.fromElements(!(_492_v0));
      r1 = (_512_v19).contains(((_492_v0) ? (_492_v0) : (_492_v0)));
      let _513_v20;
      _513_v20 = _dafny.Seq.of(new BigNumber(996), _511_v18.f10);
      let _514_v21;
      _514_v21 = _module.D6.create_DC16(_513_v20);
      r2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_513_v20, (_514_v21).dtor_cf22), _dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), ((_515_v0) => function (_516_i1) {
        return new BigNumber((function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(new _dafny.CodePoint('n'.codePointAt(0))),(_this).f8)).Keys.Elements) {
            let _517_v22 = _compr_19;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(new _dafny.CodePoint('n'.codePointAt(0))),(_this).f8)).contains(_517_v22)) {
              _coll19.push([_517_v22,_515_v0]);
            }
          }
          return _coll19;
        }()).length);
      })(_492_v0)));
      return [r0, r1, r2];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f6 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    m1(globalState) {
      let _this = this;
      let _518_v0;
      _518_v0 = _module.D0.create_DC1(new _dafny.CodePoint('e'.codePointAt(0)));
      let _519_v1;
      _519_v1 = true;
      let _520_v2;
      _520_v2 = new _dafny.CodePoint('t'.codePointAt(0));
      let _521_v3;
      _521_v3 = _module.D1.create_DC2(_519_v1);
      let _522_v4;
      let _nw122 = Array((new BigNumber(23)).toNumber());
      _nw122[(_dafny.ZERO).toNumber()] = _518_v0;
      _nw122[(_dafny.ONE).toNumber()] = _518_v0;
      _nw122[(new BigNumber(2)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(3)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(4)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(5)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(6)).toNumber()] = _module.__default.fm5(_519_v1, _519_v1, _519_v1, globalState);
      _nw122[(new BigNumber(7)).toNumber()] = _module.D0.create_DC1(_520_v2);
      _nw122[(new BigNumber(8)).toNumber()] = _module.D0.create_DC1(_520_v2);
      _nw122[(new BigNumber(9)).toNumber()] = _module.D0.create_DC1(_520_v2);
      _nw122[(new BigNumber(10)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(11)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(12)).toNumber()] = _module.__default.fm5(_519_v1, _module.__default.fm2((_521_v3).dtor_cf2, globalState), _519_v1, globalState);
      _nw122[(new BigNumber(13)).toNumber()] = _module.D0.create_DC1(_520_v2);
      _nw122[(new BigNumber(14)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(15)).toNumber()] = _module.__default.fm5(_519_v1, _519_v1, false, globalState);
      _nw122[(new BigNumber(16)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(17)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(18)).toNumber()] = _module.__default.fm5(_519_v1, _519_v1, _519_v1, globalState);
      _nw122[(new BigNumber(19)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(20)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(21)).toNumber()] = _518_v0;
      _nw122[(new BigNumber(22)).toNumber()] = _module.D0.create_DC1(_520_v2);
      _522_v4 = _nw122;
      let _523_v5;
      _523_v5 = new BigNumber(-361);
      let _524_v6;
      _524_v6 = _dafny.MultiSet.fromElements(_523_v5);
      let _525_v7;
      _525_v7 = _dafny.Seq.of(_519_v1, _519_v1);
      let _526_v8;
      _526_v8 = _dafny.MultiSet.fromElements(new BigNumber((_524_v6).cardinality()), new BigNumber((_525_v7).length));
      let _527_v9;
      _527_v9 = _dafny.Map.Empty.slice().updateUnsafe(_522_v4,new BigNumber((_524_v6).cardinality()));
      _527_v9 = (_527_v9).update(_522_v4, new BigNumber(59));
      _519_v1 = _519_v1;
      let _528_i0;
      _528_i0 = _dafny.ZERO;
      L2: {
        while (_519_v1) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_528_i0)) {
              break L2;
            }
            _528_i0 = (_528_i0).plus(_dafny.ONE);
            let _529_v10;
            let _init13 = ((_530_v5) => function (_531_i1) {
              return _module.__default.safeModuloInt(_531_i1, (_dafny.ZERO).minus(_530_v5));
            })(_523_v5);
            let _nw123 = Array((new BigNumber(18)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw123.length); _i0_13++) {
              _nw123[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _529_v10 = _nw123;
            let _532_v11;
            _532_v11 = _module.D0.create_DC0(_529_v10);
            let _533_v12;
            _533_v12 = _dafny.Seq.of(_529_v10);
            let _pat_let_tv11 = _533_v12;
            let _pat_let_tv12 = _523_v5;
            let _pat_let_tv13 = _533_v12;
            let _source6 = function (_pat_let12_0) {
              return function (_534_dt__update__tmp_h0) {
                return function (_pat_let13_0) {
                  return function (_535_dt__update_hcf0_h0) {
                    return _module.D0.create_DC0(_535_dt__update_hcf0_h0);
                  }(_pat_let13_0);
                }((_pat_let_tv11)[_module.__default.safeIndex(_pat_let_tv12, new BigNumber((_pat_let_tv13).length))]);
              }(_pat_let12_0);
            }(_532_v11);
            if (_source6.is_DC1) {
              let _536___mcc_h0 = (_source6).cf1;
              let _537_cf1 = _536___mcc_h0;
              let _538_v13;
              let _nw124 = new _module.C3();
              _nw124.__ctor(_523_v5, ((_519_v1) ? (_523_v5) : (_523_v5)));
              _538_v13 = _nw124;
              let _539_v14;
              _539_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(524),_523_v5);
              let _540_v15;
              _540_v15 = _dafny.Map.Empty.slice().updateUnsafe(_519_v1,_dafny.Map.Empty.slice().updateUnsafe(_523_v5,_523_v5));
              let _541_v17;
              _541_v17 = _dafny.Seq.of((((_540_v15).contains(_519_v1)) ? ((_540_v15).get(_519_v1)) : (function () {
                let _coll20 = new _dafny.Map();
                for (const _compr_20 of _dafny.IntegerRange(new BigNumber(856), new BigNumber(532))) {
                  let _542_v16 = _compr_20;
                  if (((new BigNumber(856)).isLessThanOrEqualTo(_542_v16)) && ((_542_v16).isLessThan(new BigNumber(532)))) {
                    _coll20.push([_module.__default.safeModuloInt(_542_v16, _523_v5),_523_v5]);
                  }
                }
                return _coll20;
              }())));
              let _543_v18;
              _543_v18 = _dafny.Map.Empty.slice().updateUnsafe(_520_v2,!(_dafny.Seq.contains(_541_v17, _539_v14)));
              _543_v18 = (_543_v18).update(_520_v2, true);
              let _544_v19;
              let _init14 = ((_545_v1) => function (_546_i2) {
                return _545_v1;
              })(_519_v1);
              let _nw125 = Array((new BigNumber(26)).toNumber());
              for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw125.length); _i0_14++) {
                _nw125[_i0_14] = _init14(new BigNumber(_i0_14));
              }
              _544_v19 = _nw125;
              let _index94 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_544_v19).length));
              (_544_v19)[_index94] = ((_519_v1) ? (_519_v1) : (_519_v1));
              let _index95 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_544_v19).length));
              (_544_v19)[_index95] = _519_v1;
              let _547_v20;
              _547_v20 = _dafny.Seq.UnicodeFromString("xjhkseht");
              let _548_v21;
              _548_v21 = _module.D6.create_DC17(_547_v20, _519_v1, _523_v5, false, !((_544_v19)[_module.__default.safeIndex(new BigNumber(679), new BigNumber((_544_v19).length))]));
              (globalState).f3 = !_dafny.areEqual(_547_v20, (_548_v21).dtor_cf23);
            } else {
              let _549___mcc_h1 = (_source6).cf0;
              let _550_cf0 = _549___mcc_h1;
              let _551_v22;
              _551_v22 = _dafny.Seq.UnicodeFromString("jsstabmc");
              _551_v22 = _dafny.Seq.Concat(_551_v22, ((_519_v1) ? (_551_v22) : (_dafny.Seq.UnicodeFromString("f"))));
              let _552_v23;
              let _init15 = ((_553_v1) => function (_554_i3) {
                return _553_v1;
              })(_519_v1);
              let _nw126 = Array((new BigNumber(2)).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw126.length); _i0_15++) {
                _nw126[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _552_v23 = _nw126;
              _552_v23 = _552_v23;
              _551_v22 = _module.__default.fm11(_519_v1, _523_v5, globalState);
              _523_v5 = (_dafny.ZERO).minus(_523_v5);
            }
            _520_v2 = _520_v2;
            let _555_v24;
            let _nw127 = Array((new BigNumber(26)).toNumber());
            _nw127[(_dafny.ZERO).toNumber()] = _519_v1;
            _nw127[(_dafny.ONE).toNumber()] = _519_v1;
            _nw127[(new BigNumber(2)).toNumber()] = !(_519_v1);
            _nw127[(new BigNumber(3)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(4)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(5)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(6)).toNumber()] = _module.__default.fm2(_519_v1, globalState);
            _nw127[(new BigNumber(7)).toNumber()] = false;
            _nw127[(new BigNumber(8)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(9)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(10)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(11)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(12)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(13)).toNumber()] = true;
            _nw127[(new BigNumber(14)).toNumber()] = !(false);
            _nw127[(new BigNumber(15)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(16)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(17)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(18)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(19)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(20)).toNumber()] = true;
            _nw127[(new BigNumber(21)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(22)).toNumber()] = _module.__default.fm2(_519_v1, globalState);
            _nw127[(new BigNumber(23)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(24)).toNumber()] = _519_v1;
            _nw127[(new BigNumber(25)).toNumber()] = !(!(_519_v1));
            _555_v24 = _nw127;
            let _556_v25;
            _556_v25 = _dafny.MultiSet.fromElements(_555_v24, _555_v24);
            let _557_v26;
            _557_v26 = _dafny.Map.Empty.slice().updateUnsafe(_556_v25,_519_v1);
            _557_v26 = (_557_v26).update(_556_v25, (_523_v5).isEqualTo(_523_v5));
            let _558_v27;
            _558_v27 = _dafny.Seq.of(_523_v5);
            let _index96 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_529_v10).length));
            (_529_v10)[_index96] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-759)), ((_559_v2) => function (_560_i4) {
              return _559_v2;
            })(_520_v2)), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("t"), _module.__default.safeIndex(new BigNumber((_558_v27).length), new BigNumber((_dafny.Seq.UnicodeFromString("t")).length)), _520_v2))).length);
            let _561_v28;
            _561_v28 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_525_v7).length),_523_v5));
            let _index97 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_529_v10).length));
            (_529_v10)[_index97] = new BigNumber(((_561_v28).Intersect((_561_v28).Union(_module.__default.fm18(globalState)))).cardinality());
          }
        }
      }
      let _562_v29;
      _562_v29 = _dafny.Seq.UnicodeFromString("geidfsen");
      let _563_v30;
      _563_v30 = _dafny.Map.Empty.slice().updateUnsafe(_523_v5,new BigNumber(188));
      let _pat_let_tv14 = _519_v1;
      let _564_v32;
      _564_v32 = _module.D6.create_DC17(_562_v29, (function (_pat_let14_0) {
  return function (_566_dt__update__tmp_h1) {
    return function (_pat_let15_0) {
      return function (_567_dt__update_hcf17_h0) {
        return _module.D4.create_DC12((_566_dt__update__tmp_h1).dtor_cf15, (_566_dt__update__tmp_h1).dtor_cf16, _567_dt__update_hcf17_h0);
      }(_pat_let15_0);
    }(!(_pat_let_tv14));
  }(_pat_let14_0);
}(_module.D4.create_DC12(new BigNumber((function () {
  let _coll21 = new _dafny.Set();
  for (const _compr_21 of (_563_v30).Keys.Elements) {
    let _565_v31 = _compr_21;
    if ((_563_v30).contains(_565_v31)) {
      _coll21.add((_565_v31).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(143))).cardinality())));
    }
  }
  return _coll21;
}()).length), _dafny.Seq.of(_519_v1), false))).dtor_cf17, (((_526_v8).contains(_523_v5)) ? ((_526_v8).get(_523_v5)) : (_523_v5)), (_525_v7)[_module.__default.safeIndex((_dafny.ZERO).minus(_523_v5), new BigNumber((_525_v7).length))], false);
      _519_v1 = (_564_v32).dtor_cf27;
      let _568_v33;
      let _init16 = ((_569_v1) => function (_570_i5) {
        return _569_v1;
      })(_519_v1);
      let _nw128 = Array((new BigNumber(10)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw128.length); _i0_16++) {
        _nw128[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _568_v33 = _nw128;
      let _index98 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_568_v33).length));
      (_568_v33)[_index98] = _519_v1;
      let _index99 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_568_v33).length));
      (_568_v33)[_index99] = !((_525_v7)[_module.__default.safeIndex(_523_v5, new BigNumber((_525_v7).length))]);
      (globalState).f1 = (_568_v33)[_module.__default.safeIndex(new BigNumber(61), new BigNumber((_568_v33).length))];
      return;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let _571_v0;
      let _nw129 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _571_v0 = _nw129;
      let _572_v1;
      _572_v1 = _module.D0.create_DC0(_571_v0);
      let _573_v2;
      let _nw130 = Array((new BigNumber(6)).toNumber()).fill(false);
      _573_v2 = _nw130;
      let _574_v3;
      _574_v3 = _dafny.Map.Empty.slice().updateUnsafe((_572_v1).dtor_cf0,_573_v2);
      _574_v3 = (_574_v3).update(_571_v0, _573_v2);
      let _575_v4;
      let _nw131 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
      _575_v4 = _nw131;
      let _576_v5;
      _576_v5 = new BigNumber(-872);
      let _577_v6;
      _577_v6 = _dafny.Seq.of(_576_v5);
      let _578_v7;
      _578_v7 = _dafny.Map.Empty.slice().updateUnsafe(_576_v5,new BigNumber(((_module.D6.create_DC16(_577_v6)).dtor_cf22).length));
      let _index100 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_575_v4).length));
      (_575_v4)[_index100] = _578_v7;
      let _579_v8;
      _579_v8 = new _dafny.CodePoint('a'.codePointAt(0));
      let _580_v9;
      _580_v9 = _dafny.MultiSet.fromElements(_571_v0, _571_v0, _571_v0, _571_v0, _571_v0);
      let _581_v10;
      _581_v10 = _dafny.Seq.UnicodeFromString("jmjl");
      let _582_v11;
      _582_v11 = _dafny.Map.Empty.slice().updateUnsafe(true,p2);
      let _index101 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_575_v4).length));
      let _rhs74 = _578_v7;
      let _rhs75 = _579_v8;
      let _rhs76 = (_module.D4.create_DC12(new BigNumber((_581_v10).length), p1, p2)).dtor_cf17;
      let _rhs77 = new BigNumber((_582_v11).length);
      let _rhs78 = _dafny.MultiSet.fromElements(_571_v0);
      let _lhs63 = _575_v4;
      let _lhs64 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_575_v4).length));
      let _lhs65 = globalState;
      _lhs63[_lhs64] = _rhs74;
      _579_v8 = _rhs75;
      _lhs65.f3 = _rhs76;
      _576_v5 = _rhs77;
      _580_v9 = _rhs78;
      let _583_v12;
      _583_v12 = _dafny.MultiSet.fromElements(_576_v5, _576_v5);
      let _584_v13;
      _584_v13 = _dafny.MultiSet.fromElements(_579_v8, _579_v8);
      let _585_v14;
      _585_v14 = _dafny.MultiSet.fromElements((((_583_v12).contains(_576_v5)) ? ((_583_v12).get(_576_v5)) : (new BigNumber((_584_v13).cardinality()))));
      let _586_v15;
      _586_v15 = _dafny.Set.fromElements(_583_v12);
      let _587_v16;
      _587_v16 = _module.D3.create_DC9(_module.D3.create_DC7(_586_v15));
      _587_v16 = _module.__default.fm19(globalState);
      (globalState).f1 = (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), ((_588_v8) => function (_589_i0) {
        return _588_v8;
      })(_579_v8)), _581_v10)) || (!(p2) || (p2));
      if (!((p1)[_module.__default.safeIndex(_576_v5, new BigNumber((p1).length))])) {
        let _index102 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_573_v2).length));
        (_573_v2)[_index102] = p2;
        let _index103 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_573_v2).length));
        (_573_v2)[_index103] = !((p1)[_module.__default.safeIndex(_576_v5, new BigNumber((p1).length))]);
        let _index104 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_573_v2).length));
        let _index105 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_573_v2).length));
        let _rhs79 = p2;
        let _rhs80 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), ((_590_v8) => function (_591_i1) {
          return _590_v8;
        })(_579_v8)), _581_v10)).length);
        let _rhs81 = _module.__default.fm0(_579_v8, globalState);
        let _rhs82 = _576_v5;
        let _rhs83 = p2;
        let _lhs66 = _573_v2;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_573_v2).length));
        let _lhs68 = _573_v2;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_573_v2).length));
        _lhs66[_lhs67] = _rhs79;
        _576_v5 = _rhs80;
        _576_v5 = _rhs81;
        _576_v5 = _rhs82;
        _lhs68[_lhs69] = _rhs83;
        (globalState).f1 = !(_576_v5).isEqualTo((_576_v5).minus(_576_v5));
        let _index106 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_573_v2).length));
        (_573_v2)[_index106] = !(p2);
        (globalState).f1 = _module.__default.fm2((_573_v2)[_module.__default.safeIndex(new BigNumber(747), new BigNumber((_573_v2).length))], globalState);
        let _index107 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_575_v4).length));
        (_575_v4)[_index107] = ((_575_v4)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_575_v4).length))]).update(_576_v5, _module.__default.safeModuloInt(_576_v5, _576_v5));
      } else {
        let _592_v17;
        let _nw132 = new _module.C2();
        _nw132.__ctor(_576_v5, new BigNumber((_dafny.MultiSet.FromArray(_577_v6)).cardinality()), _576_v5, _576_v5);
        _592_v17 = _nw132;
        (globalState).f3 = _module.__default.fm2(false, globalState);
        (globalState).f3 = p2;
        let _593_v18;
        _593_v18 = _module.D6.create_DC17(_581_v10, false, _592_v17.f10, p2, p2);
        _576_v5 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_592_v17).f9, (_592_v17).f9), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), ((_594_v8) => function (_595_i2) {
          return _594_v8;
        })(_579_v8)), (_593_v18).dtor_cf23)).length));
        let _596_v19;
        _596_v19 = _module.D3.create_DC8(_592_v17.f10);
        let _pat_let_tv15 = _592_v17;
        let _597_v20;
        _597_v20 = _dafny.MultiSet.fromElements(_596_v19, _596_v19, function (_pat_let16_0) {
          return function (_598_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_599_dt__update_hcf10_h0) {
                return _module.D3.create_DC8(_599_dt__update_hcf10_h0);
              }(_pat_let17_0);
            }((_pat_let_tv15).f9);
          }(_pat_let16_0);
        }(_596_v19));
        let _600_v21;
        _600_v21 = _module.D4.create_DC11(new BigNumber((_dafny.MultiSet.fromElements(_576_v5, (_592_v17).f9)).cardinality()), _579_v8);
        let _601_v22;
        _601_v22 = _dafny.Set.fromElements(true);
        let _602_v23;
        _602_v23 = _module.D2.create_DC4(new BigNumber(776));
        let _603_v24;
        _603_v24 = _dafny.Seq.of(_601_v22, _module.__default.fm20(_602_v23, p2, globalState));
        let _604_v25;
        let _nw133 = Array((new BigNumber(24)).toNumber());
        _nw133[(_dafny.ZERO).toNumber()] = _581_v10;
        _nw133[(_dafny.ONE).toNumber()] = _581_v10;
        _nw133[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_581_v10, _module.__default.safeIndex(new BigNumber(172), new BigNumber((_581_v10).length)), _579_v8);
        _nw133[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_581_v10, _module.__default.safeIndex((((_597_v20).contains(_596_v19)) ? ((_597_v20).get(_596_v19)) : ((_dafny.ZERO).minus((_600_v21).dtor_cf13))), new BigNumber((_581_v10).length)), _579_v8), _module.__default.safeIndex(new BigNumber(((_603_v24)[_module.__default.safeIndex(_592_v17.f10, new BigNumber((_603_v24).length))]).length), new BigNumber((_dafny.Seq.update(_581_v10, _module.__default.safeIndex((((_597_v20).contains(_596_v19)) ? ((_597_v20).get(_596_v19)) : ((_dafny.ZERO).minus((_600_v21).dtor_cf13))), new BigNumber((_581_v10).length)), _579_v8)).length)), _579_v8);
        _nw133[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("t"), _581_v10);
        _nw133[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-519)), ((_605_v8) => function (_606_i3) {
          return _605_v8;
        })(_579_v8));
        _nw133[(new BigNumber(6)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("sval");
        _nw133[(new BigNumber(8)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(9)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_581_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(743)), ((_607_v8) => function (_608_i4) {
          return _607_v8;
        })(_579_v8)));
        _nw133[(new BigNumber(11)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_581_v10, _581_v10);
        _nw133[(new BigNumber(13)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(14)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(15)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(16)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(17)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_581_v10, _dafny.Seq.update(_581_v10, _module.__default.safeIndex(_576_v5, new BigNumber((_581_v10).length)), _579_v8));
        _nw133[(new BigNumber(19)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(20)).toNumber()] = _581_v10;
        _nw133[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_581_v10, _dafny.Seq.UnicodeFromString("tswjrvwl"));
        _nw133[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_581_v10, _581_v10);
        _nw133[(new BigNumber(23)).toNumber()] = _581_v10;
        _604_v25 = _nw133;
        let _index108 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_604_v25).length));
        (_604_v25)[_index108] = _581_v10;
        let _index109 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_604_v25).length));
        (_604_v25)[_index109] = ((!(!(p2) || (!(!(p2))))) ? (_581_v10) : (_dafny.Seq.Concat(_581_v10, _dafny.Seq.UnicodeFromString("hupyfkie"))));
      }
      let _609_v26;
      let _nw134 = Array((new BigNumber(4)).toNumber());
      _609_v26 = _nw134;
      let _610_v27;
      let _nw135 = new _module.C3();
      _nw135.__ctor((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(p2, true)).length)), new BigNumber(383));
      _610_v27 = _nw135;
      let _index110 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_609_v26).length));
      (_609_v26)[_index110] = _610_v27;
      let _index111 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_609_v26).length));
      (_609_v26)[_index111] = _610_v27;
      return;
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
