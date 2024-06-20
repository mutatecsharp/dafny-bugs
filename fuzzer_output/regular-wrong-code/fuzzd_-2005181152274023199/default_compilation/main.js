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
    static fm0(globalState) {
      return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(253)), function (_0_i0) {
        return new BigNumber(357);
      })).length),true), (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(-341))).length),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)))).length),false)));
    };
    static fm1(p0, p1, p2, globalState) {
      if (true) {
        return (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.UnicodeFromString("gggckn")).length)).minus(new BigNumber(-433)));
      } else if (true) {
        return new BigNumber(356);
      } else {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),new BigNumber(-565))).length);
      }
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return (new BigNumber((_dafny.Seq.UnicodeFromString("kvqorjhh")).length)).isLessThan(new BigNumber(-411));
    };
    static fm7(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(76))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-449)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(692), new BigNumber(418))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(692)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(418)))) {
            _coll0.push([_module.__default.safeDivisionInt(_1_v0, new BigNumber((_dafny.Seq.of(!(!(true)))).length)),false]);
          }
        }
        return _coll0;
      }()).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(713))));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(937)));
    };
    static fm9(p0, globalState) {
      return _dafny.Set.fromElements((true) === (true), false, _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("qylrshxjb")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), function (_2_i0) {
        return _dafny.Seq.UnicodeFromString("tp");
      })), !(true) || (true), false);
    };
    static fm12(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lsqi"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), function (_3_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }));
    };
    static fm13(p0, p1, globalState) {
      if (_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("b"), new _dafny.CodePoint('m'.codePointAt(0)))) {
        return _dafny.Seq.UnicodeFromString("ip");
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("j"), _dafny.Seq.UnicodeFromString("lucm"));
      }
    };
    static fm16(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(477)), function (_4_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(866)), function (_5_i1) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(731)), function (_6_i2) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })));
    };
    static fm23(p0, p1, globalState) {
      return !(((_dafny.ZERO).minus((_module.D29.create_DC78(new BigNumber(236), new BigNumber((function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of _dafny.IntegerRange(new BigNumber(861), new BigNumber(243))) {
    let _7_v0 = _compr_1;
    if (((new BigNumber(861)).isLessThanOrEqualTo(_7_v0)) && ((_7_v0).isLessThan(new BigNumber(243)))) {
      _coll1.push([(_7_v0).plus(new BigNumber(959)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(867)), function (_8_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })).length)]);
    }
  }
  return _coll1;
}()).length), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(926)))).dtor_cf135)).isLessThanOrEqualTo(new BigNumber(228)));
    };
    static fm24(p0, p1, p2, globalState) {
      if (!(new BigNumber(892)).isEqualTo(new BigNumber(-451))) {
        return _module.D1.create_DC4(false);
      } else {
        return _module.D1.create_DC4(true);
      }
    };
    static fm25(p0, globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber(-381)).multipliedBy(new BigNumber(-192)));
    };
    static fm26(p0, p1, globalState) {
      return _module.D2.create_DC8(!((_module.D1.create_DC4(false)).dtor_cf12), !(((_module.D25.create_DC67(!(true), new BigNumber(81), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("wa")).length)), true)).dtor_cf117).isLessThanOrEqualTo(new BigNumber(779))), !(false));
    };
    static fm27(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).Elements) {
          let _9_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0))), _9_v0)) {
            _coll2.add(_9_v0);
          }
        }
        return _coll2;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0))))));
    };
    static fm28(p0, globalState) {
      return _module.D1.create_DC3(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-560)), function (_10_i0) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})).length), new BigNumber((_dafny.Seq.UnicodeFromString("fvskreag")).length)), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-92)), new BigNumber(-899), new BigNumber((_dafny.Set.fromElements(new BigNumber(83))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-512),new _dafny.CodePoint('w'.codePointAt(0)))).length), new BigNumber(377))).length), new BigNumber((_dafny.Seq.UnicodeFromString("nf")).length)), _module.__default.safeDivisionInt(_dafny.ONE, new BigNumber(-868)), new BigNumber(636), (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('i'.codePointAt(0)))).length))).minus(new BigNumber(-203))));
    };
    static fm29(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(308)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-361)))).length),true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(312)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus((_module.D1.create_DC3(new BigNumber(112), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wasaeyknb")).length),new BigNumber(-405))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(39)), function (_11_i0) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}))).length), new BigNumber(-633))).dtor_cf8))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(473))));
    };
    static fm30(p0, p1, globalState) {
      return function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(76)), _dafny.Seq.of(new BigNumber((function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of (_dafny.Seq.of(new BigNumber(451), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(829))).length))).Elements) {
            let _12_v1 = _compr_4;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(451), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(829))).length)), _12_v1)) {
              _coll4.add((_12_v1).plus(new BigNumber((_dafny.Set.fromElements(false, !(false))).length)));
            }
          }
          return _coll4;
        }()).length)))).Elements) {
          let _13_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(76)), _dafny.Seq.of(new BigNumber((function () {
            let _coll5 = new _dafny.Set();
            for (const _compr_5 of (_dafny.Seq.of(new BigNumber(451), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(829))).length))).Elements) {
              let _14_v1 = _compr_5;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(451), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(829))).length)), _14_v1)) {
                _coll5.add((_14_v1).plus(new BigNumber((_dafny.Set.fromElements(false, !(false))).length)));
              }
            }
            return _coll5;
          }()).length))), _13_v0)) {
            _coll3.push([(_13_v0).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("rg")).length))).length))),(new BigNumber(530)).plus(new BigNumber((_dafny.Seq.of(new BigNumber(459), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-871)), function (_15_i0) {
              return new _dafny.CodePoint('n'.codePointAt(0));
            })).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false)).length)), new BigNumber(63))).length))).length))]);
          }
        }
        return _coll3;
      }();
    };
    static fm31(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_module.D3.create_DC13(new BigNumber((function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(471), new BigNumber(379))) {
    let _16_v0 = _compr_6;
    if (((new BigNumber(471)).isLessThanOrEqualTo(_16_v0)) && ((_16_v0).isLessThan(new BigNumber(379)))) {
      _coll6.push([_module.__default.safeDivisionInt(_16_v0, new BigNumber((function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-833)), function (_17_i0) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("vcf")).length);
        })).Elements) {
          let _18_v1 = _compr_7;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-833)), function (_17_i0) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("vcf")).length);
          }), _18_v1)) {
            _coll7.add((_18_v1).multipliedBy(new BigNumber(-187)));
          }
        }
        return _coll7;
      }()).length)),false]);
    }
  }
  return _coll6;
}()).length), _dafny.Seq.UnicodeFromString("fahmrgw"), new BigNumber((_dafny.Set.fromElements(new BigNumber(354), new BigNumber(29))).length), new BigNumber((_dafny.Seq.UnicodeFromString("dnvpfxa")).length), new _dafny.CodePoint('s'.codePointAt(0)))), _dafny.Seq.of(_module.D3.create_DC13(new BigNumber(364), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-922)), function (_19_i1) {
  return new _dafny.CodePoint('y'.codePointAt(0));
}), new BigNumber(-517), new BigNumber((function () {
  let _coll8 = new _dafny.Map();
  for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-609), new BigNumber(922))) {
    let _20_v2 = _compr_8;
    if (((new BigNumber(-609)).isLessThanOrEqualTo(_20_v2)) && ((_20_v2).isLessThan(new BigNumber(922)))) {
      _coll8.push([_module.__default.safeDivisionInt(_20_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-510))).length)),true]);
    }
  }
  return _coll8;
}()).length), new _dafny.CodePoint('q'.codePointAt(0))), _module.D3.create_DC13(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), _dafny.Seq.UnicodeFromString("tosl"), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(972), new _dafny.CodePoint('q'.codePointAt(0))))))).Union((_dafny.MultiSet.fromElements(_module.D3.create_DC13(new BigNumber(830), _dafny.Seq.UnicodeFromString("xg"), new BigNumber((_dafny.Set.fromElements(true, true)).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of _dafny.IntegerRange(new BigNumber(438), new BigNumber(597))) {
    let _21_v3 = _compr_9;
    if (((new BigNumber(438)).isLessThanOrEqualTo(_21_v3)) && ((_21_v3).isLessThan(new BigNumber(597)))) {
      _coll9.push([(_21_v3).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(963)), function (_22_i2) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length)),true]);
    }
  }
  return _coll9;
}())).length))).length), new _dafny.CodePoint('r'.codePointAt(0))))).Union(_dafny.MultiSet.fromElements(_module.D3.create_DC13(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new BigNumber(202), new BigNumber(457))).length))).length), _dafny.Seq.UnicodeFromString("aiq"), new BigNumber(82), new BigNumber((_dafny.Seq.UnicodeFromString("aynd")).length), new _dafny.CodePoint('p'.codePointAt(0))), _module.D3.create_DC13(new BigNumber(493), _dafny.Seq.UnicodeFromString("uw"), new BigNumber(462), new BigNumber((_dafny.Seq.of(true, true)).length), new _dafny.CodePoint('f'.codePointAt(0))))));
    };
    static fm32(globalState) {
      if (((true) ? (false) : (false))) {
        return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true)));
      } else {
        return (_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)));
      }
    };
    static fm33(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(911), new BigNumber(616))) {
          let _23_v0 = _compr_10;
          if (((new BigNumber(911)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(616)))) {
            _coll10.add((_23_v0).minus(new BigNumber(-58)));
          }
        }
        return _coll10;
      }()).length),new BigNumber(527))).length), new BigNumber((_dafny.Seq.of(new BigNumber(94))).length), new BigNumber((_dafny.Seq.UnicodeFromString("pbmngnpin")).length), new BigNumber(-658), new BigNumber((_dafny.Seq.of(_module.D2.create_DC8(!(true), true, false))).length)), _dafny.Seq.of(new BigNumber(783)));
    };
    static fm34(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(424),_module.D2.create_DC8(true, true, false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-165),_module.D2.create_DC8(true, true, true)));
    };
    static fm35(p0, p1, globalState) {
      return _module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("mphupn")).length), _dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0))), !(_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), function (_24_i0) {
  return (_module.D16.create_DC40(new BigNumber((_dafny.Set.fromElements(false, false)).length), new BigNumber(450), new BigNumber(211), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true, true)).length),false)).length))).dtor_cf72;
}), new BigNumber(368))), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(710), new BigNumber(14))));
    };
    static fm36(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(false, false, true)).Intersect(_dafny.Set.fromElements(false))).Intersect(((!(true)) ? (_dafny.Set.fromElements(false, (_module.D10.create_DC25(false)).dtor_cf48, false, false)) : (_dafny.Set.fromElements(false, true, !(true)))));
    };
    static fm39(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(829)), function (_25_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), function (_26_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        });
      }), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cykrhjkg"), _dafny.Seq.UnicodeFromString("fe"))), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cwmhocws"), _dafny.Seq.UnicodeFromString("boaslvno")));
    };
    static fm40(p0, globalState) {
      return _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lmsi"), _dafny.Seq.UnicodeFromString("vh"))).length), new BigNumber(-127), (_dafny.ZERO).minus((new BigNumber(-289)).plus(new BigNumber(-497))));
    };
    static fm41(globalState) {
      return _module.D0.create_DC0(_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(47), new BigNumber((_dafny.Set.fromElements(false)).length))), _dafny.Seq.of(new BigNumber(624), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("mfvja")).length)), new BigNumber(-280))));
    };
    static fm42(p0, p1, p2, globalState) {
      return _module.D9.create_DC22(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-498), new BigNumber((_dafny.Seq.UnicodeFromString("tr")).length))).cardinality())), new BigNumber((_dafny.Seq.of(true, false)).length)));
    };
    static fm43(p0, p1, globalState) {
      return (((true) ? (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(715))) : (_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(true)))),new BigNumber(163))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber(423), new BigNumber(527), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, false, true, false, false)).length)))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(23)))).cardinality()),new BigNumber((function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, false, false)).cardinality()), new BigNumber(71))).Elements) {
          let _27_v0 = _compr_11;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, false, false)).cardinality()), new BigNumber(71))).contains(_27_v0)) {
            _coll11.push([_module.__default.safeModuloInt(_27_v0, new BigNumber((_dafny.Seq.UnicodeFromString("bgpxdfu")).length)),true]);
          }
        }
        return _coll11;
      }()).length))).length),new BigNumber((_dafny.Seq.UnicodeFromString("gmkulofbe")).length))).length))));
    };
    static fm44(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(!(!(false))), _dafny.Seq.of(!(!(!(true))), true));
    };
    static fm45(globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("frqcl"), _dafny.Seq.UnicodeFromString("njtvu"), _dafny.Seq.UnicodeFromString("lb"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_28_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }))).Intersect(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), function (_29_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), function (_30_i2) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("anl")))).Union((_module.D30.create_DC79(function () {
  let _coll12 = new _dafny.Set();
  for (const _compr_12 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), function (_31_i3) {
    return new _dafny.CodePoint('q'.codePointAt(0));
  }), _dafny.Seq.UnicodeFromString("kvx"), _dafny.Seq.UnicodeFromString("wkyiedaj"))).Elements) {
    let _32_v0 = _compr_12;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), function (_31_i3) {
      return new _dafny.CodePoint('q'.codePointAt(0));
    }), _dafny.Seq.UnicodeFromString("kvx"), _dafny.Seq.UnicodeFromString("wkyiedaj")), _32_v0)) {
      _coll12.add(_32_v0);
    }
  }
  return _coll12;
}())).dtor_cf137);
    };
    static fm46(globalState) {
      return new _dafny.CodePoint('y'.codePointAt(0));
    };
    static fm47(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false)));
    };
    static fm48(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-505)), function (_33_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      });
    };
    static fm49(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("opxh"), _dafny.Seq.UnicodeFromString("tgmmnul"))).cardinality()), new BigNumber(201), new BigNumber((_dafny.Seq.UnicodeFromString("t")).length), new BigNumber(666)))).length),new BigNumber(26))).length)).plus(new BigNumber((_dafny.Seq.UnicodeFromString("albnktp")).length)));
    };
    static fm50(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(857),true),!_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("rmbiql"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-948)), function (_34_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("yog")));
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of((function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of _dafny.IntegerRange(new BigNumber(594), new BigNumber(461))) {
          let _35_v0 = _compr_13;
          if (((new BigNumber(594)).isLessThanOrEqualTo(_35_v0)) && ((_35_v0).isLessThan(new BigNumber(461)))) {
            _coll13.push([_module.__default.safeDivisionInt(_35_v0, new BigNumber(331)),new BigNumber(-447)]);
          }
        }
        return _coll13;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(350),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-505),false)).length))));
    };
    static fm52(p0, p1, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("u"), _dafny.Seq.UnicodeFromString("dsxh"), _dafny.Seq.UnicodeFromString("yagibnn"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(947)), function (_36_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), function (_37_i1) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })))).Intersect((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("seqmtp"), _dafny.Seq.UnicodeFromString("fi"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), function (_38_i2) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("aahqhps"), _dafny.Seq.UnicodeFromString("ie")))));
    };
    static fm53(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC1(false, new BigNumber(952), _dafny.Seq.Create(_module.__default.abs(new BigNumber(572)), function (_39_i0) {
  return new _dafny.CodePoint('c'.codePointAt(0));
}), true, new BigNumber((_dafny.Seq.of(new BigNumber(430))).length)), _module.D0.create_DC1(false, new BigNumber(482), _dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0))), false, new BigNumber(280))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), function (_40_i1) {
        return _module.D0.create_DC1(false, new BigNumber(943), _dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0))), true, new BigNumber(27));
      })), _dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC1(true, new BigNumber(822), _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))), true, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(983)), function (_41_i2) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})).length)), _module.D0.create_DC1(true, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-896),true)).length),true)).length))).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(400)), function (_42_i3) {
  return new _dafny.CodePoint('w'.codePointAt(0));
}), true, new BigNumber((function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of _dafny.IntegerRange(new BigNumber(-241), new BigNumber(903))) {
    let _43_v0 = _compr_14;
    if (((new BigNumber(-241)).isLessThanOrEqualTo(_43_v0)) && ((_43_v0).isLessThan(new BigNumber(903)))) {
      _coll14.push([(_43_v0).multipliedBy((_module.D30.create_DC80(false, _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())), true, new BigNumber(-938))).dtor_cf141),new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of (_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)))).Elements) {
          let _44_v1 = _compr_15;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0))), _44_v1)) {
            _coll15.push([_44_v1,new BigNumber(-854)]);
          }
        }
        return _coll15;
      }()).length))).length)]);
    }
  }
  return _coll14;
}()).length))), _dafny.Seq.of(_module.D0.create_DC1(!(!(!(false))), new BigNumber(430), _dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0))), !(true), new BigNumber(36)), _module.D0.create_DC1(!(!(!(false))), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-918)), function (_45_i4) {
  return new _dafny.CodePoint('u'.codePointAt(0));
})).length))).length), _dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))), false, new BigNumber(435)), _module.D0.create_DC1(false, new BigNumber(167), _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0))), false, new BigNumber(-988)), _module.D0.create_DC1(true, new BigNumber(334), _dafny.Seq.of(new _dafny.CodePoint('o'.codePointAt(0))), true, new BigNumber((_dafny.Seq.of(true, true)).length)))));
    };
    static fm54(p0, p1, globalState) {
      return function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Seq.of(function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of (_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).Elements) {
            let _46_v1 = _compr_17;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))), _46_v1)) {
              _coll17.push([_46_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, true)).length),true)).length)]);
            }
          }
          return _coll17;
        }(), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(499)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(771)), function (_47_i0) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        })).length)))).Elements) {
          let _48_v0 = _compr_16;
          if (_dafny.Seq.contains(_dafny.Seq.of(function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of (_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).Elements) {
              let _46_v1 = _compr_18;
              if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))), _46_v1)) {
                _coll18.push([_46_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, true)).length),true)).length)]);
              }
            }
            return _coll18;
          }(), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(499)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(771)), function (_47_i0) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          })).length))), _48_v0)) {
            _coll16.push([_48_v0,(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-647)), function (_49_i1) {
              return new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)))).cardinality());
            })).length)).plus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(322), new BigNumber((_dafny.Seq.of(false)).length)))).cardinality()))]);
          }
        }
        return _coll16;
      }();
    };
    static fm55(p0, p1, p2, p3, globalState) {
      if (!((_dafny.MultiSet.fromElements(new BigNumber(890))).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ptwyd")).length), (_dafny.ZERO).minus(new BigNumber(-377)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(false))).length), new BigNumber(263), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qnmsblqx")).length)))).length)))))) {
        return _module.D9.create_DC23(false, !(false), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), function (_50_i0) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})).length)))).length)));
      } else {
        return _module.D9.create_DC23(true, true, new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,true))).length));
      }
    };
    static fm56(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D17.create_DC42(true);
      if (_source0.is_DC42) {
        let _51___mcc_h0 = (_source0).cf74;
        let _52_cf74 = _51___mcc_h0;
        return _module.D13.create_DC33(_52_cf74, new BigNumber((_dafny.Seq.of(_52_cf74)).length), _52_cf74, new BigNumber(117));
      } else if (_source0.is_DC43) {
        let _53___mcc_h1 = (_source0).cf75;
        let _54___mcc_h2 = (_source0).cf76;
        let _55_cf76 = _54___mcc_h2;
        let _56_cf75 = _53___mcc_h1;
        return _module.D13.create_DC33(false, new BigNumber(399), false, new BigNumber(520));
      } else if (_source0.is_DC44) {
        let _57___mcc_h3 = (_source0).cf77;
        let _58___mcc_h4 = (_source0).cf78;
        let _59_cf78 = _58___mcc_h4;
        let _60_cf77 = _57___mcc_h3;
        return _module.D13.create_DC33(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), !(true), new BigNumber((_dafny.Seq.UnicodeFromString("fn")).length));
      } else {
        let _61___mcc_h5 = (_source0).cf73;
        let _62_cf73 = _61___mcc_h5;
        if (!(false)) {
          return _module.D13.create_DC33(false, new BigNumber((_dafny.Seq.of(!(false))).length), false, new BigNumber(203));
        } else {
          return _module.D13.create_DC33(false, new BigNumber(-622), true, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(313)), function (_63_i0) {
  return new _dafny.CodePoint('u'.codePointAt(0));
})).length)),_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(true)).length))).length)))))).length)));
        }
      }
    };
    static fm57(p0, p1, p2, globalState) {
      let _source1 = _module.D3.create_DC13(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(489))).cardinality()), _dafny.Seq.UnicodeFromString("ybukhe"), new BigNumber((_dafny.Set.fromElements(true, (_module.D3.create_DC11(!(true), true)).dtor_cf24)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ujk")).length),new BigNumber(-931))).length), (_module.D8.create_DC21(new _dafny.CodePoint('n'.codePointAt(0)), new BigNumber(-482))).dtor_cf41);
      if (_source1.is_DC11) {
        let _64___mcc_h0 = (_source1).cf23;
        let _65___mcc_h1 = (_source1).cf24;
        let _66_cf24 = _65___mcc_h1;
        let _67_cf23 = _64___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(391),_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_66_cf24,new BigNumber(903))).length)))).length),_66_cf24)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(424),_67_cf23));
      } else if (_source1.is_DC12) {
        let _68___mcc_h2 = (_source1).cf25;
        let _69___mcc_h3 = (_source1).cf26;
        let _70___mcc_h4 = (_source1).cf27;
        let _71___mcc_h5 = (_source1).cf28;
        let _72___mcc_h6 = (_source1).cf29;
        let _73_cf29 = _72___mcc_h6;
        let _74_cf28 = _71___mcc_h5;
        let _75_cf27 = _70___mcc_h4;
        let _76_cf26 = _69___mcc_h3;
        let _77_cf25 = _68___mcc_h2;
        return (_module.D17.create_DC41(_dafny.Map.Empty.slice().updateUnsafe(_75_cf27,_76_cf26))).dtor_cf73;
      } else if (_source1.is_DC13) {
        let _78___mcc_h7 = (_source1).cf30;
        let _79___mcc_h8 = (_source1).cf31;
        let _80___mcc_h9 = (_source1).cf32;
        let _81___mcc_h10 = (_source1).cf33;
        let _82___mcc_h11 = (_source1).cf34;
        let _83_cf34 = _82___mcc_h11;
        let _84_cf33 = _81___mcc_h10;
        let _85_cf32 = _80___mcc_h9;
        let _86_cf31 = _79___mcc_h8;
        let _87_cf30 = _78___mcc_h7;
        if (false) {
          return _dafny.Map.Empty.slice().updateUnsafe(_84_cf33,false);
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_84_cf33),false);
        }
      } else if (_source1.is_DC10) {
        let _88___mcc_h12 = (_source1).cf22;
        let _89_cf22 = _88___mcc_h12;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("w")).length),!(!(false)));
      } else {
        let _90___mcc_h13 = (_source1).cf35;
        let _91_cf35 = _90___mcc_h13;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dvhocxjgf")).length),true)).length), new BigNumber(502), new BigNumber(-996), (_dafny.ZERO).minus(new BigNumber(-394)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length),!(false));
      }
    };
    static fm58(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.D19.create_DC50(false),true);
    };
    static fm59(globalState) {
      return _module.D3.create_DC13((new BigNumber((_dafny.Seq.of(false, false)).length)).minus(new BigNumber((_dafny.Seq.of(false, false)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(149)), function (_92_i0) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}), ((false) ? (new BigNumber(654)) : (new BigNumber((_dafny.Seq.UnicodeFromString("wblaoerrr")).length))), _module.__default.safeModuloInt(new BigNumber(77), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(((_module.D25.create_DC67(true, new BigNumber(399), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber(-158), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length)), true)).dtor_cf118).length))).length))), new _dafny.CodePoint('u'.codePointAt(0)));
    };
    static fm60(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_module.D29.create_DC78(new BigNumber(165), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("tcb"))).length), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(931)))).dtor_cf135,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ww")).length),new BigNumber(-340)));
    };
    static fm61(p0, globalState) {
      let _source2 = _module.D20.create_DC53(true, (_module.D19.create_DC50(true)).dtor_cf84, new _dafny.CodePoint('l'.codePointAt(0)));
      if (_source2.is_DC53) {
        let _93___mcc_h0 = (_source2).cf87;
        let _94___mcc_h1 = (_source2).cf88;
        let _95___mcc_h2 = (_source2).cf89;
        let _96_cf89 = _95___mcc_h2;
        let _97_cf88 = _94___mcc_h1;
        let _98_cf87 = _93___mcc_h0;
        if ((_module.D19.create_DC49(false)).dtor_cf83) {
          return _module.D21.create_DC55(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_98_cf87,_96_cf89)).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_97_cf88,new BigNumber(-677))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(213),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("uoqthk")).length))).length))));
        } else {
          return _module.D21.create_DC55(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-230),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_98_cf87, _98_cf87)).cardinality())),_98_cf87),(_module.D13.create_DC33(_97_cf88, new BigNumber((_dafny.Seq.UnicodeFromString("t")).length), _98_cf87, new BigNumber(-347))).dtor_cf56)).length))).length))));
        }
      } else if (_source2.is_DC52) {
        let _99___mcc_h3 = (_source2).cf86;
        let _100_cf86 = _99___mcc_h3;
        return _module.D21.create_DC55(_dafny.Set.fromElements(_100_cf86, _100_cf86, _100_cf86, _100_cf86));
      } else {
        let _101___mcc_h4 = (_source2).cf90;
        let _102_cf90 = _101___mcc_h4;
        return _module.D21.create_DC55(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-390)), function (_103_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length),new BigNumber((_dafny.Set.fromElements(true)).length))));
      }
    };
    static fm62(p0, p1, globalState) {
      return (_module.D31.create_DC83(_dafny.Seq.of(_dafny.Set.fromElements(true)))).dtor_cf143;
    };
    static fm63(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false, false)).length))).isEqualTo(new BigNumber((_dafny.Set.fromElements(!(true), true, true)).length)),new _dafny.CodePoint('c'.codePointAt(0)));
    };
    static fm64(p0, p1, p2, p3, globalState) {
      if ((false) && (!(true))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(391))).length))).length), new BigNumber(622), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(672), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D31.create_DC84(false, false)).dtor_cf144,new _dafny.CodePoint('m'.codePointAt(0)))).length), new BigNumber(718), new BigNumber((_dafny.Seq.of(!(false), true)).length))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("tsuf")).length)), (_dafny.ZERO).minus(new BigNumber(-435)))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(224)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-9))).length))).length)));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(250))).length))).length));
      }
    };
    static fm65(globalState) {
      return function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-409)), function (_104_i0) {
          return _dafny.Seq.UnicodeFromString("hw");
        }), _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-614)), function (_105_i1) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })))).Elements) {
          let _106_v0 = _compr_19;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-409)), function (_104_i0) {
            return _dafny.Seq.UnicodeFromString("hw");
          }), _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-614)), function (_105_i1) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          }))), _106_v0)) {
            _coll19.push([_106_v0,new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("isiwmcovd"), _dafny.Seq.UnicodeFromString("dkdiy"))).length)]);
          }
        }
        return _coll19;
      }();
    };
    static m0(globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let r3 = _dafny.Map.Empty;
      let _107_v0;
      _107_v0 = false;
      let _108_v1;
      let _nw0 = Array((new BigNumber(12)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _107_v0;
      _nw0[(_dafny.ONE).toNumber()] = _107_v0;
      _nw0[(new BigNumber(2)).toNumber()] = _107_v0;
      _nw0[(new BigNumber(3)).toNumber()] = _107_v0;
      _nw0[(new BigNumber(4)).toNumber()] = _107_v0;
      _nw0[(new BigNumber(5)).toNumber()] = _107_v0;
      _nw0[(new BigNumber(6)).toNumber()] = _107_v0;
      _nw0[(new BigNumber(7)).toNumber()] = _107_v0;
      _nw0[(new BigNumber(8)).toNumber()] = false;
      _nw0[(new BigNumber(9)).toNumber()] = _107_v0;
      _nw0[(new BigNumber(10)).toNumber()] = _107_v0;
      _nw0[(new BigNumber(11)).toNumber()] = _107_v0;
      _108_v1 = _nw0;
      let _109_v2;
      _109_v2 = _dafny.Seq.UnicodeFromString("cvivu");
      let _110_v3;
      _110_v3 = _module.D13.create_DC34(_108_v1, _109_v2, _107_v0, _107_v0);
      if ((_110_v3).dtor_cf61) {
        let _111_v4;
        _111_v4 = _dafny.Map.Empty.slice().updateUnsafe(_108_v1,_107_v0);
        let _112_v5;
        _112_v5 = _module.D13.create_DC32(_111_v4);
        let _113_v6;
        _113_v6 = new BigNumber(-108);
        let _114_v7;
        _114_v7 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jm"),(_dafny.ZERO).minus(_113_v6));
        let _115_v8;
        let _nw1 = new _module.C5();
        _nw1.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(980)), function (_116_i0) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("lf"), !(_module.__default.fm65(globalState)).equals(_114_v7));
        _115_v8 = _nw1;
        let _117_v9;
        _117_v9 = _dafny.MultiSet.fromElements((_113_v6).multipliedBy(_113_v6));
        let _118_v10;
        let _nw2 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _118_v10 = _nw2;
        let _119_v11;
        _119_v11 = _dafny.Seq.of(_108_v1);
        let _120_v12;
        _120_v12 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_108_v1), _119_v11, _dafny.Seq.update(_119_v11, _module.__default.safeIndex(_module.__default.fm1(_module.__default.fm1(_113_v6, _107_v0, _107_v0, globalState), _107_v0, _107_v0, globalState), new BigNumber((_119_v11).length)), _108_v1));
        let _rhs0 = _112_v5;
        let _rhs1 = _115_v8;
        let _rhs2 = _118_v10;
        let _rhs3 = (((_120_v12).contains(_dafny.Seq.Concat(_119_v11, _119_v11))) ? ((_120_v12).get(_dafny.Seq.Concat(_119_v11, _119_v11))) : ((_113_v6).plus(_113_v6)));
        let _rhs4 = _117_v9;
        _112_v5 = _rhs0;
        _115_v8 = _rhs1;
        r2 = _rhs2;
        _113_v6 = _rhs3;
        _117_v9 = _rhs4;
        let _121_v13;
        _121_v13 = new _dafny.CodePoint('n'.codePointAt(0));
        _121_v13 = _module.__default.fm46(globalState);
        _107_v0 = ((_113_v6).minus(_113_v6)).isEqualTo(_113_v6);
        let _122_v14;
        _122_v14 = _dafny.Seq.of(_107_v0);
        let _123_v15;
        _123_v15 = _dafny.Map.Empty.slice().updateUnsafe(_113_v6,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(100)), ((_124_v13) => function (_125_i1) {
          return _124_v13;
        })(_121_v13))).length));
        let _126_v16;
        _126_v16 = _dafny.Seq.of(_123_v15, _dafny.Map.Empty.slice().updateUnsafe(_113_v6,_113_v6));
        let _127_v17;
        let _nw3 = new _module.C6();
        _nw3.__ctor(_122_v14, _113_v6, !(true), (_115_v8).f28, _126_v16);
        _127_v17 = _nw3;
        let _128_v18;
        _128_v18 = _dafny.Map.Empty.slice().updateUnsafe(_127_v17,_121_v13);
        let _rhs5 = (((_128_v18).contains(_127_v17)) ? ((_128_v18).get(_127_v17)) : (_121_v13));
        let _rhs6 = (_127_v17).f19;
        let _lhs0 = globalState;
        _121_v13 = _rhs5;
        _lhs0.f9 = _rhs6;
        let _129_v19;
        let _nw4 = Array((new BigNumber(3)).toNumber());
        _129_v19 = _nw4;
        let _nw5 = Array((new BigNumber(4)).toNumber());
        _129_v19 = _nw5;
      } else {
        (globalState).f12 = _dafny.Seq.Concat(_109_v2, _109_v2);
        let _130_v20;
        _130_v20 = new BigNumber(825);
        let _131_v21;
        let _nw6 = new _module.C0();
        _nw6.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_130_v20, _107_v0, _107_v0, globalState),_107_v0)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(692)), function (_132_i2) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        })).length));
        _131_v21 = _nw6;
        let _133_v22;
        _133_v22 = _dafny.MultiSet.fromElements(_131_v21, _131_v21);
        (globalState).f15 = (((_133_v22).contains(_131_v21)) ? ((_133_v22).get(_131_v21)) : (new BigNumber(664)));
        let _134_v23;
        _134_v23 = _dafny.Seq.of(_130_v20);
        let _135_v24;
        _135_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_107_v0)).length),(_131_v21).f32);
        let _136_v25;
        _136_v25 = _dafny.Seq.of(_135_v24, _135_v24, _135_v24, _135_v24, _135_v24);
        let _137_v26;
        let _nw7 = new _module.C3();
        _nw7.__ctor(_134_v23, _109_v2, _dafny.Seq.Concat(_136_v25, _136_v25), ((_107_v0) ? (_107_v0) : (false)));
        _137_v26 = _nw7;
        _137_v26 = _137_v26;
        (globalState).f15 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_131_v21).f32), _134_v23)).length);
        let _138_v27;
        _138_v27 = _module.D21.create_DC56((_137_v26).f18, _130_v20, new BigNumber(229));
        if ((_138_v27).dtor_cf92) {
          let _139_v28;
          _139_v28 = new _dafny.CodePoint('k'.codePointAt(0));
          _139_v28 = _139_v28;
          let _index0 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_108_v1).length));
          (_108_v1)[_index0] = _107_v0;
          let _140_v29;
          _140_v29 = _module.D16.create_DC39(_107_v0, new BigNumber(306));
          let _index1 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_108_v1).length));
          (_108_v1)[_index1] = (_140_v29).dtor_cf67;
          let _141_v30;
          let _nw8 = new _module.C3();
          _nw8.__ctor(_module.__default.fm33((_108_v1)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_108_v1).length))], globalState), _dafny.Seq.UnicodeFromString("xj"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_142_v24) => function (_143_i3) {
            return _142_v24;
          })(_135_v24)), _136_v25), (_137_v26).f18);
          _141_v30 = _nw8;
          let _144_v31;
          _144_v31 = _dafny.Set.fromElements(!((_137_v26).f18), (_137_v26).f18);
          let _index2 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_108_v1).length));
          (_108_v1)[_index2] = (_dafny.Set.fromElements((_108_v1)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_108_v1).length))], _107_v0)).IsSubsetOf(_144_v31);
          let _145_v32;
          _145_v32 = _dafny.MultiSet.fromElements((_137_v26).f18);
          (globalState).f2 = (_145_v32).IsProperSubsetOf(_145_v32);
        } else {
          _107_v0 = (_137_v26).f18;
          let _146_v33;
          _146_v33 = _module.D9.create_DC23(_107_v0, (_137_v26).f18, (_131_v21).f31);
          let _147_v34;
          let _nw9 = new _module.C8();
          _nw9.__ctor((_131_v21).f32, (_137_v26).f18);
          _147_v34 = _nw9;
          let _148_v35;
          _148_v35 = _dafny.MultiSet.fromElements(_147_v34);
          let _149_v36;
          let _nw10 = new _module.C3();
          _nw10.__ctor(_dafny.Seq.of(new BigNumber(807), (_131_v21).f32, new BigNumber(-416), new BigNumber((_148_v35).cardinality())), (_137_v26).f19, (_137_v26).f20, (_137_v26).f18);
          _149_v36 = _nw10;
          let _150_v37;
          let _nw11 = new _module.C7();
          _nw11.__ctor(_130_v20, (_137_v26).f18);
          _150_v37 = _nw11;
          let _151_v38;
          _151_v38 = _dafny.Set.fromElements((_137_v26).f18, _107_v0);
          let _rhs7 = _146_v33;
          let _rhs8 = ((_151_v38).IsSubsetOf(_151_v38)) === ((_137_v26).f18);
          let _rhs9 = (_131_v21).f31;
          let _rhs10 = _149_v36;
          let _rhs11 = _150_v37;
          let _lhs1 = globalState;
          let _lhs2 = globalState;
          _146_v33 = _rhs7;
          _lhs1.f2 = _rhs8;
          _lhs2.f3 = _rhs9;
          _149_v36 = _rhs10;
          _150_v37 = _rhs11;
          let _index3 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_108_v1).length));
          (_108_v1)[_index3] = false;
          let _index4 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_108_v1).length));
          (_108_v1)[_index4] = (_137_v26).f18;
          (globalState).f0 = _147_v34.f23;
          let _152_v39;
          let _nw12 = new _module.C6();
          _nw12.__ctor(_dafny.Seq.of((_137_v26).f18), new BigNumber(((_137_v26).f19).length), (_137_v26).f18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-34)), function (_153_i4) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }), (_137_v26).f20);
          _152_v39 = _nw12;
        }
      }
      let _154_i5;
      _154_i5 = _dafny.ZERO;
      L0: {
        while (!(_107_v0)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_154_i5)) {
              break L0;
            }
            _154_i5 = (_154_i5).plus(_dafny.ONE);
            let _155_v40;
            let _nw13 = new _module.C9();
            _nw13.__ctor(_108_v1, _107_v0);
            _155_v40 = _nw13;
            let _156_v41;
            let _nw14 = Array((new BigNumber(7)).toNumber());
            _nw14[(_dafny.ZERO).toNumber()] = _155_v40;
            _nw14[(_dafny.ONE).toNumber()] = _155_v40;
            _nw14[(new BigNumber(2)).toNumber()] = _155_v40;
            _nw14[(new BigNumber(3)).toNumber()] = _155_v40;
            _nw14[(new BigNumber(4)).toNumber()] = _155_v40;
            _nw14[(new BigNumber(5)).toNumber()] = _155_v40;
            _nw14[(new BigNumber(6)).toNumber()] = _155_v40;
            _156_v41 = _nw14;
            let _index5 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_156_v41).length));
            (_156_v41)[_index5] = _155_v40;
            let _157_v42;
            _157_v42 = new BigNumber(725);
            let _index6 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_156_v41).length));
            let _rhs12 = _157_v42;
            let _rhs13 = _155_v40;
            let _lhs3 = _156_v41;
            let _lhs4 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_156_v41).length));
            r1 = _rhs12;
            _lhs3[_lhs4] = _rhs13;
            let _158_v43;
            _158_v43 = _dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(_157_v42)), (_dafny.ZERO).minus(_157_v42));
            _158_v43 = _158_v43;
            let _159_v44;
            _159_v44 = _dafny.Seq.of(_107_v0, _107_v0, _107_v0, _107_v0);
            if ((!(!(_107_v0))) || ((_159_v44)[_module.__default.safeIndex(_157_v42, new BigNumber((_159_v44).length))])) {
              let _160_v45;
              let _nw15 = new _module.C0();
              _nw15.__ctor((_157_v42).multipliedBy(_157_v42), new BigNumber((_109_v2).length));
              _160_v45 = _nw15;
              let _161_v46;
              _161_v46 = _dafny.Map.Empty.slice().updateUnsafe((_160_v45).f32,_157_v42);
              let _162_v47;
              _162_v47 = _dafny.Map.Empty.slice().updateUnsafe(_107_v0,new BigNumber(815));
              let _163_v48;
              _163_v48 = _module.D3.create_DC13(_157_v42, _109_v2, _157_v42, (((_162_v47).contains((_155_v40).f18)) ? ((_162_v47).get((_155_v40).f18)) : (new BigNumber(507))), new _dafny.CodePoint('c'.codePointAt(0)));
              let _164_v49;
              _164_v49 = _dafny.Seq.of(_161_v46, _161_v46, _161_v46);
              let _165_v50;
              let _nw16 = new _module.C6();
              _nw16.__ctor(_module.__default.fm44(new BigNumber((_161_v46).length), (_155_v40).f18, (_163_v48).dtor_cf34, (((_162_v47).contains((_155_v40).f18)) ? ((_162_v47).get((_155_v40).f18)) : (new BigNumber(550))), globalState), ((_160_v45).f31).multipliedBy(_157_v42), (_155_v40).f18, _109_v2, _164_v49);
              _165_v50 = _nw16;
              _165_v50 = _165_v50;
              _109_v2 = (_110_v3).dtor_cf60;
              (globalState).f2 = true;
              (globalState).f2 = !((_160_v45).f32).isEqualTo((_dafny.ZERO).minus((_160_v45).f31));
            } else {
              (globalState).f2 = false;
              _107_v0 = true;
              (globalState).f2 = _107_v0;
              _107_v0 = _107_v0;
              r0 = _157_v42;
            }
            let _166_v51;
            let _nw17 = new _module.C1();
            _nw17.__ctor(_107_v0);
            _166_v51 = _nw17;
            let _167_v52;
            _167_v52 = _dafny.Map.Empty.slice().updateUnsafe(_166_v51,(_155_v40).f18);
            let _168_v53;
            _168_v53 = _module.D11.create_DC27((_167_v52).update(_166_v51, (_155_v40).f18));
            let _169_v54;
            _169_v54 = _module.D11.create_DC29(_168_v53);
            _169_v54 = _169_v54;
          }
        }
      }
      if (_107_v0) {
        let _170_v55;
        _170_v55 = new BigNumber(-816);
        let _171_v56;
        _171_v56 = _dafny.Seq.of((_dafny.ZERO).minus(_170_v55));
        let _172_v57;
        _172_v57 = _dafny.Set.fromElements(_107_v0, _107_v0);
        let _173_v58;
        _173_v58 = _dafny.Map.Empty.slice().updateUnsafe(_170_v55,new BigNumber((_172_v57).length));
        let _174_v59;
        _174_v59 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_170_v55,_170_v55), _173_v58);
        let _175_v60;
        let _nw18 = new _module.C3();
        _nw18.__ctor(_171_v56, _109_v2, _174_v59, !(((_107_v0) ? (_107_v0) : (_107_v0))));
        _175_v60 = _nw18;
        _175_v60 = _175_v60;
        (globalState).f2 = (_107_v0) || (_107_v0);
        let _176_v61;
        _176_v61 = new _dafny.CodePoint('b'.codePointAt(0));
        let _177_v62;
        _177_v62 = _dafny.Map.Empty.slice().updateUnsafe(false,false);
        let _178_v63;
        _178_v63 = _dafny.MultiSet.fromElements(new BigNumber((_177_v62).length));
        let _179_v64;
        _179_v64 = _dafny.MultiSet.fromElements(_107_v0);
        let _180_v65;
        _180_v65 = _dafny.Seq.of(_179_v64);
        let _181_v66;
        _181_v66 = _dafny.Map.Empty.slice().updateUnsafe((_180_v65)[_module.__default.safeIndex(new BigNumber(791), new BigNumber((_180_v65).length))],_170_v55);
        let _182_v67;
        let _nw19 = Array((new BigNumber(25)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _109_v2;
        _nw19[(_dafny.ONE).toNumber()] = _109_v2;
        _nw19[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_109_v2, _109_v2), _module.__default.safeIndex(((_175_v60).f35)[_module.__default.safeIndex((_dafny.ZERO).minus(_170_v55), new BigNumber(((_175_v60).f35).length))], new BigNumber((_dafny.Seq.Concat(_109_v2, _109_v2)).length)), _176_v61);
        _nw19[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("cgsgthyqi");
        _nw19[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(952)), ((_183_v61) => function (_184_i6) {
          return _183_v61;
        })(_176_v61));
        _nw19[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mfhtk"), _dafny.Seq.UnicodeFromString("fkjwvty"));
        _nw19[(new BigNumber(6)).toNumber()] = _109_v2;
        _nw19[(new BigNumber(7)).toNumber()] = _109_v2;
        _nw19[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_109_v2, _module.__default.safeIndex(_170_v55, new BigNumber((_109_v2).length)), _176_v61);
        _nw19[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("gwuxuka");
        _nw19[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("yftjej");
        _nw19[(new BigNumber(11)).toNumber()] = _module.__default.fm16(_170_v55, true, _176_v61, _170_v55, globalState);
        _nw19[(new BigNumber(12)).toNumber()] = _109_v2;
        _nw19[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_109_v2, _109_v2);
        _nw19[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("nvtv");
        _nw19[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(((_107_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), ((_185_v61) => function (_186_i7) {
          return _185_v61;
        })(_176_v61))) : (_dafny.Seq.UnicodeFromString("nnpdry"))), _module.__default.safeIndex((((_178_v63).contains((_dafny.ZERO).minus(_170_v55))) ? ((_178_v63).get((_dafny.ZERO).minus(_170_v55))) : (_170_v55)), new BigNumber((((_107_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), ((_187_v61) => function (_188_i7) {
          return _187_v61;
        })(_176_v61))) : (_dafny.Seq.UnicodeFromString("nnpdry")))).length)), _176_v61);
        _nw19[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("vbke");
        _nw19[(new BigNumber(17)).toNumber()] = ((_module.__default.fm23(_dafny.Set.fromElements(_170_v55), new BigNumber(((_181_v66).update(_179_v64, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(459)), function (_189_i8) {
          return new BigNumber(954);
        })).length))).length), globalState)) ? (_109_v2) : (_109_v2));
        _nw19[(new BigNumber(18)).toNumber()] = (_175_v60).fm37(_170_v55, globalState);
        _nw19[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), ((_190_v61) => function (_191_i9) {
          return _190_v61;
        })(_176_v61));
        _nw19[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_109_v2, _dafny.Seq.update(_109_v2, _module.__default.safeIndex(_170_v55, new BigNumber((_109_v2).length)), _176_v61));
        _nw19[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(609)), ((_192_v61) => function (_193_i10) {
          return _192_v61;
        })(_176_v61));
        _nw19[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_109_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(219)), ((_194_v61) => function (_195_i11) {
          return _194_v61;
        })(_176_v61)));
        _nw19[(new BigNumber(23)).toNumber()] = _109_v2;
        _nw19[(new BigNumber(24)).toNumber()] = _109_v2;
        _182_v67 = _nw19;
        _182_v67 = _182_v67;
        let _196_v68;
        let _nw20 = new _module.C0();
        _nw20.__ctor(_170_v55, _170_v55);
        _196_v68 = _nw20;
        let _rhs14 = _107_v0;
        let _rhs15 = _170_v55;
        _107_v0 = _rhs14;
        _170_v55 = _rhs15;
      } else {
        let _197_v69;
        _197_v69 = _dafny.Set.fromElements(_107_v0);
        let _198_v70;
        let _nw21 = new _module.C5();
        _nw21.__ctor(_109_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(853)), function (_199_i12) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), (_197_v69).IsSubsetOf(_197_v69));
        _198_v70 = _nw21;
        _198_v70 = _198_v70;
        let _200_v71;
        let _nw22 = new _module.C5();
        _nw22.__ctor(_109_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(821)), function (_201_i13) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        }), (((_198_v70).f18) ? ((_198_v70).f18) : ((_198_v70).f18)));
        _200_v71 = _nw22;
        _200_v71 = _200_v71;
        let _202_v72;
        _202_v72 = new BigNumber(999);
        let _203_v73;
        _203_v73 = _dafny.Map.Empty.slice().updateUnsafe((_198_v70).f18,_202_v72);
        (globalState).f2 = (_203_v73).contains((((_198_v70).f18) ? (_107_v0) : ((_198_v70).f18)));
        let _index7 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_108_v1).length));
        (_108_v1)[_index7] = (_198_v70).f18;
        let _204_v74;
        _204_v74 = _dafny.Seq.of(true);
        let _index8 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_108_v1).length));
        (_108_v1)[_index8] = (!((_204_v74)[_module.__default.safeIndex(_202_v72, new BigNumber((_204_v74).length))]) || (true)) || ((_198_v70).f18);
        (globalState).f12 = _dafny.Seq.UnicodeFromString("tvvi");
      }
      let _205_v75;
      _205_v75 = new BigNumber(690);
      let _206_v76;
      _206_v76 = _module.D17.create_DC44(_108_v1, _205_v75);
      _206_v76 = _206_v76;
      r1 = _205_v75;
      let _207_v77;
      _207_v77 = _dafny.Set.fromElements(_205_v75, _205_v75);
      let _208_v78;
      _208_v78 = _dafny.Map.Empty.slice().updateUnsafe(_107_v0,_207_v77);
      let _209_i14;
      _209_i14 = _dafny.ZERO;
      L1: {
        while (_module.__default.fm23((((_208_v78).contains(_107_v0)) ? ((_208_v78).get(_107_v0)) : (_207_v77)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-424)), function (_255_i15) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }), _109_v2)).length), globalState)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_209_i14)) {
              break L1;
            }
            _209_i14 = (_209_i14).plus(_dafny.ONE);
            let _210_v79;
            _210_v79 = _dafny.MultiSet.fromElements(_205_v75, _module.__default.fm1(new BigNumber(-202), _107_v0, !(!(_107_v0)), globalState));
            _210_v79 = _210_v79;
            let _211_v80;
            _211_v80 = _dafny.MultiSet.fromElements(_107_v0);
            let _212_v81;
            let _init0 = function (_213_i16) {
              return (_213_i16).multipliedBy(new BigNumber(392));
            };
            let _nw23 = Array((new BigNumber(10)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw23.length); _i0_0++) {
              _nw23[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _212_v81 = _nw23;
            let _214_v82;
            _214_v82 = _dafny.Set.fromElements(_212_v81);
            let _215_v83;
            let _nw24 = Array((new BigNumber(3)).toNumber());
            _nw24[(_dafny.ZERO).toNumber()] = _205_v75;
            _nw24[(_dafny.ONE).toNumber()] = new BigNumber((_211_v80).cardinality());
            _nw24[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_109_v2).length), new BigNumber((_214_v82).length)));
            _215_v83 = _nw24;
            let _index9 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length));
            (_212_v81)[_index9] = _205_v75;
            let _216_v84;
            let _nw25 = new _module.C1();
            _nw25.__ctor(_107_v0);
            _216_v84 = _nw25;
            let _217_v85;
            _217_v85 = _dafny.Set.fromElements(_216_v84);
            let _218_v86;
            _218_v86 = _module.D1.create_DC3(new BigNumber((_217_v85).length), _205_v75, _205_v75, _205_v75, new BigNumber((_109_v2).length));
            let _219_v87;
            _219_v87 = _dafny.Seq.of(_218_v86, _module.D1.create_DC3(_205_v75, new BigNumber((_109_v2).length), new BigNumber(494), _module.__default.fm1(_205_v75, _107_v0, _107_v0, globalState), _205_v75), _218_v86, _218_v86, _218_v86);
            let _index10 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_212_v81).length));
            (_212_v81)[_index10] = _module.__default.fm1(new BigNumber((_dafny.MultiSet.FromArray(_219_v87)).cardinality()), false, _107_v0, globalState);
            let _220_v88;
            _220_v88 = new _dafny.CodePoint('n'.codePointAt(0));
            let _221_v89;
            _221_v89 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_107_v0, _107_v0), _module.__default.safeIndex(new BigNumber(398), new BigNumber((_dafny.Seq.of(_107_v0, _107_v0)).length)), _107_v0)).length),_205_v75);
            let _222_v90;
            _222_v90 = _dafny.Seq.of(_221_v89, _221_v89);
            let _223_v91;
            let _nw26 = new _module.C4();
            _nw26.__ctor(_220_v88, _205_v75, _107_v0, _109_v2, _222_v90);
            _223_v91 = _nw26;
            let _224_v92;
            let _nw27 = new _module.C4();
            _nw27.__ctor(_220_v88, _205_v75, _107_v0, _109_v2, _222_v90);
            _224_v92 = _nw27;
            let _225_v93;
            _225_v93 = _dafny.Map.Empty.slice().updateUnsafe(_224_v92,_223_v91);
            let _226_v94;
            _226_v94 = _dafny.Seq.of(_223_v91, (((_225_v93).contains(_224_v92)) ? ((_225_v93).get(_224_v92)) : (_223_v91)), _223_v91);
            let _227_v95;
            _227_v95 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_223_v91, _223_v91), _226_v94, _dafny.Seq.of(_223_v91, _223_v91));
            let _228_v96;
            _228_v96 = _module.D28.create_DC74(((_107_v0) ? (_227_v95) : (_227_v95)));
            let _index11 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_108_v1).length));
            (_108_v1)[_index11] = !(false);
            let _229_v97;
            _229_v97 = _module.D23.create_DC61(_108_v1, _107_v0, true, (_223_v91).f18);
            let _230_v98;
            _230_v98 = _dafny.Seq.of(_229_v97, _229_v97);
            let _231_v99;
            _231_v99 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(103),_dafny.Seq.of(_229_v97, _229_v97, _229_v97, _module.D23.create_DC61(_108_v1, _107_v0, _107_v0, _107_v0)));
            let _232_v100;
            _232_v100 = _dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(341),_230_v98)).Merge(_231_v99));
            let _233_v101;
            _233_v101 = _dafny.Set.fromElements((_224_v92).f18, (_223_v91).f18, _107_v0, _107_v0);
            let _234_v102;
            _234_v102 = _dafny.Seq.of(_dafny.Set.fromElements(_205_v75, _205_v75, _205_v75, _205_v75, new BigNumber(255)), _207_v77, _207_v77, _207_v77, _module.__default.fm49(new BigNumber((_233_v101).length), new BigNumber(810), _205_v75, globalState));
            let _235_v103;
            _235_v103 = _dafny.Seq.of(_205_v75);
            let _236_v104;
            _236_v104 = _module.D1.create_DC5(_module.D1.create_DC3(_205_v75, _205_v75, new BigNumber((_221_v89).length), new BigNumber((_dafny.MultiSet.fromElements((_224_v92).f18)).cardinality()), (_235_v103)[_module.__default.safeIndex(_205_v75, new BigNumber((_235_v103).length))]));
            let _237_v105;
            _237_v105 = _module.D1.create_DC5(_236_v104);
            let _238_v106;
            _238_v106 = _module.D1.create_DC5(_237_v105);
            let _239_v107;
            _239_v107 = _module.D1.create_DC5(_module.D1.create_DC5(_238_v106));
            let _240_v108;
            _240_v108 = _dafny.Map.Empty.slice().updateUnsafe(_239_v107,_205_v75);
            let _index12 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length));
            let _index13 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_212_v81).length));
            let _index14 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_108_v1).length));
            let _rhs16 = (((_232_v100).contains((_231_v99).Merge(_231_v99))) ? ((_232_v100).get((_231_v99).Merge(_231_v99))) : (new BigNumber((_234_v102).length)));
            let _rhs17 = (((_240_v108).contains(_module.D1.create_DC5(_237_v105))) ? ((_240_v108).get(_module.D1.create_DC5(_237_v105))) : (_205_v75));
            let _rhs18 = _228_v96;
            let _rhs19 = !(false);
            let _lhs5 = _212_v81;
            let _lhs6 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length));
            let _lhs7 = _212_v81;
            let _lhs8 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_212_v81).length));
            let _lhs9 = _108_v1;
            let _lhs10 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_108_v1).length));
            _lhs5[_lhs6] = _rhs16;
            _lhs7[_lhs8] = _rhs17;
            _228_v96 = _rhs18;
            _lhs9[_lhs10] = _rhs19;
            let _241_v109;
            _241_v109 = _module.D16.create_DC40((_212_v81)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length))], _205_v75, (_212_v81)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length))], (_212_v81)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length))]);
            let _242_v110;
            _242_v110 = _dafny.Map.Empty.slice().updateUnsafe((_212_v81)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length))],_241_v109);
            let _243_v113;
            let _nw28 = Array((new BigNumber(9)).toNumber());
            _nw28[(_dafny.ZERO).toNumber()] = _242_v110;
            _nw28[(_dafny.ONE).toNumber()] = _242_v110;
            _nw28[(new BigNumber(2)).toNumber()] = _242_v110;
            _nw28[(new BigNumber(3)).toNumber()] = _242_v110;
            _nw28[(new BigNumber(4)).toNumber()] = _242_v110;
            _nw28[(new BigNumber(5)).toNumber()] = function () {
              let _coll20 = new _dafny.Map();
              for (const _compr_20 of _dafny.IntegerRange(new BigNumber(992), new BigNumber(-148))) {
                let _244_v111 = _compr_20;
                if (((new BigNumber(992)).isLessThanOrEqualTo(_244_v111)) && ((_244_v111).isLessThan(new BigNumber(-148)))) {
                  _coll20.push([_module.__default.safeDivisionInt(_244_v111, new BigNumber(-711)),_241_v109]);
                }
              }
              return _coll20;
            }();
            _nw28[(new BigNumber(6)).toNumber()] = _242_v110;
            _nw28[(new BigNumber(7)).toNumber()] = function () {
              let _coll21 = new _dafny.Map();
              for (const _compr_21 of (_221_v89).Keys.Elements) {
                let _245_v112 = _compr_21;
                if ((_221_v89).contains(_245_v112)) {
                  _coll21.push([(_245_v112).multipliedBy(_205_v75),_241_v109]);
                }
              }
              return _coll21;
            }();
            _nw28[(new BigNumber(8)).toNumber()] = _242_v110;
            _243_v113 = _nw28;
            let _246_v114;
            _246_v114 = _module.D26.create_DC70(_205_v75, (_212_v81)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length))], _107_v0, _243_v113);
            let _247_v115;
            let _nw29 = new _module.C5();
            _nw29.__ctor(_109_v2, _109_v2, (_223_v91).f18);
            _247_v115 = _nw29;
            let _248_v116;
            _248_v116 = _dafny.Seq.of(_247_v115, _247_v115, _247_v115);
            let _249_v117;
            _249_v117 = _dafny.Map.Empty.slice().updateUnsafe((_223_v91).f18,_107_v0);
            let _250_v118;
            let _nw30 = Array((new BigNumber(21)).toNumber());
            _nw30[(_dafny.ZERO).toNumber()] = _246_v114;
            _nw30[(_dafny.ONE).toNumber()] = _module.D26.create_DC70(_205_v75, new BigNumber(617), !((_223_v91).f18), _243_v113);
            _nw30[(new BigNumber(2)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(3)).toNumber()] = function (_pat_let0_0) {
              return function (_251_dt__update__tmp_h0) {
                return function (_pat_let1_0) {
                  return function (_252_dt__update_hcf123_h0) {
                    return _module.D26.create_DC70((_251_dt__update__tmp_h0).dtor_cf122, _252_dt__update_hcf123_h0, (_251_dt__update__tmp_h0).dtor_cf124, (_251_dt__update__tmp_h0).dtor_cf125);
                  }(_pat_let1_0);
                }(new BigNumber(871));
              }(_pat_let0_0);
            }(_246_v114);
            _nw30[(new BigNumber(4)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(5)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(6)).toNumber()] = _module.D26.create_DC70((_dafny.ZERO).minus(_205_v75), (_212_v81)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length))], (_223_v91).f18, _243_v113);
            _nw30[(new BigNumber(7)).toNumber()] = _module.D26.create_DC70(_205_v75, new BigNumber((_248_v116).length), false, _243_v113);
            _nw30[(new BigNumber(8)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(9)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(10)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(11)).toNumber()] = _module.D26.create_DC70(new BigNumber((_dafny.Set.fromElements((_108_v1)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_108_v1).length))])).length), new BigNumber((_211_v80).cardinality()), (_223_v91).f18, _243_v113);
            _nw30[(new BigNumber(12)).toNumber()] = _module.D26.create_DC70(new BigNumber(((_221_v89).update((_212_v81)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length))], (_212_v81)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_212_v81).length))])).length), _205_v75, (_223_v91).f18, _243_v113);
            _nw30[(new BigNumber(13)).toNumber()] = _module.D26.create_DC70(new BigNumber((_dafny.Seq.of((_108_v1)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_108_v1).length))])).length), new BigNumber((_249_v117).length), _107_v0, _243_v113);
            _nw30[(new BigNumber(14)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(15)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(16)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(17)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(18)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(19)).toNumber()] = _246_v114;
            _nw30[(new BigNumber(20)).toNumber()] = _246_v114;
            _250_v118 = _nw30;
            let _index15 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_250_v118).length));
            (_250_v118)[_index15] = (((_247_v115).fm15(new BigNumber((_207_v77).length), _205_v75, globalState)) ? (_246_v114) : (_246_v114));
            let _253_v119;
            let _nw31 = Array((new BigNumber(19)).toNumber()).fill(false);
            _253_v119 = _nw31;
            let _254_v120;
            _254_v120 = _dafny.Map.Empty.slice().updateUnsafe(_253_v119,_220_v88);
            let _index16 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_250_v118).length));
            (_250_v118)[_index16] = _module.D26.create_DC70(_205_v75, new BigNumber((_254_v120).length), _107_v0, _243_v113);
            _220_v88 = _220_v88;
          }
        }
      }
      r0 = (_dafny.ZERO).minus(_205_v75);
      r1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_205_v75));
      let _256_v121;
      _256_v121 = _dafny.Seq.of(_107_v0, _107_v0);
      let _257_v122;
      _257_v122 = _dafny.Seq.of(_205_v75, _205_v75, new BigNumber((_256_v121).length));
      let _258_v123;
      _258_v123 = _dafny.Map.Empty.slice().updateUnsafe(_107_v0,_257_v122);
      let _259_v124;
      _259_v124 = _dafny.Set.fromElements(_107_v0, _107_v0);
      let _260_v125;
      _260_v125 = _dafny.Map.Empty.slice().updateUnsafe(_259_v124,_205_v75);
      let _261_v126;
      _261_v126 = new _dafny.CodePoint('v'.codePointAt(0));
      let _262_v127;
      _262_v127 = _dafny.Map.Empty.slice().updateUnsafe(_205_v75,new BigNumber((_109_v2).length));
      let _263_v128;
      let _nw32 = new _module.C9();
      _nw32.__ctor(_108_v1, true);
      _263_v128 = _nw32;
      let _264_v129;
      _264_v129 = _dafny.Map.Empty.slice().updateUnsafe(_263_v128,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(105)), ((_265_v126) => function (_266_i18) {
        return _265_v126;
      })(_261_v126))).length));
      let _267_v130;
      _267_v130 = _module.D8.create_DC21(_261_v126, _module.__default.fm1(_205_v75, _107_v0, (_263_v128).f18, globalState));
      let _268_v131;
      _268_v131 = _dafny.MultiSet.fromElements(_module.__default.fm1(_205_v75, _107_v0, _107_v0, globalState));
      let _269_v132;
      _269_v132 = _dafny.MultiSet.fromElements(_261_v126);
      let _270_v133;
      let _nw33 = Array((new BigNumber(26)).toNumber());
      _nw33[(_dafny.ZERO).toNumber()] = new BigNumber(87);
      _nw33[(_dafny.ONE).toNumber()] = new BigNumber(((_258_v123).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-21)), ((_271_v75) => function (_272_i17) {
        return _271_v75;
      })(_205_v75))))).length);
      _nw33[(new BigNumber(2)).toNumber()] = _205_v75;
      _nw33[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((((_260_v125).contains(_259_v124)) ? ((_260_v125).get(_259_v124)) : (_205_v75)));
      _nw33[(new BigNumber(4)).toNumber()] = _205_v75;
      _nw33[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("sihya"), _module.__default.safeIndex(_205_v75, new BigNumber((_dafny.Seq.UnicodeFromString("sihya")).length)), _261_v126)).length);
      _nw33[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_205_v75, new BigNumber((_257_v122).length))));
      _nw33[(new BigNumber(7)).toNumber()] = _205_v75;
      _nw33[(new BigNumber(8)).toNumber()] = new BigNumber(735);
      _nw33[(new BigNumber(9)).toNumber()] = _205_v75;
      _nw33[(new BigNumber(10)).toNumber()] = (new BigNumber((_109_v2).length)).multipliedBy(_205_v75);
      _nw33[(new BigNumber(11)).toNumber()] = new BigNumber(-405);
      _nw33[(new BigNumber(12)).toNumber()] = new BigNumber(232);
      _nw33[(new BigNumber(13)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_262_v127).length))).plus((((_264_v129).contains(_263_v128)) ? ((_264_v129).get(_263_v128)) : ((_dafny.ZERO).minus(_205_v75))));
      _nw33[(new BigNumber(14)).toNumber()] = new BigNumber((_109_v2).length);
      _nw33[(new BigNumber(15)).toNumber()] = _205_v75;
      _nw33[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_256_v121, _256_v121)).length);
      _nw33[(new BigNumber(17)).toNumber()] = (_267_v130).dtor_cf42;
      _nw33[(new BigNumber(18)).toNumber()] = (_205_v75).multipliedBy(_205_v75);
      _nw33[(new BigNumber(19)).toNumber()] = _205_v75;
      _nw33[(new BigNumber(20)).toNumber()] = _205_v75;
      _nw33[(new BigNumber(21)).toNumber()] = new BigNumber((_268_v131).cardinality());
      _nw33[(new BigNumber(22)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(702), new BigNumber((_269_v132).cardinality()));
      _nw33[(new BigNumber(23)).toNumber()] = (((_268_v131).contains(new BigNumber((_109_v2).length))) ? ((_268_v131).get(new BigNumber((_109_v2).length))) : (new BigNumber((_module.__default.fm49(_205_v75, _205_v75, _205_v75, globalState)).length)));
      _nw33[(new BigNumber(24)).toNumber()] = _module.__default.safeDivisionInt(_205_v75, (_dafny.ZERO).minus(new BigNumber((_109_v2).length)));
      _nw33[(new BigNumber(25)).toNumber()] = (_205_v75).plus(_205_v75);
      _270_v133 = _nw33;
      r2 = _270_v133;
      r3 = function () {
        let _coll22 = new _dafny.Map();
        for (const _compr_22 of (_207_v77).Elements) {
          let _273_v134 = _compr_22;
          if ((_207_v77).contains(_273_v134)) {
            _coll22.push([_module.__default.safeModuloInt(_273_v134, _205_v75),_107_v0]);
          }
        }
        return _coll22;
      }();
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _274_v0;
      _274_v0 = _dafny.Seq.UnicodeFromString("ewld");
      let _275_v1;
      _275_v1 = _dafny.Map.Empty.slice().updateUnsafe(true,true);
      let _276_v2;
      _276_v2 = new BigNumber(-284);
      let _277_v3;
      _277_v3 = false;
      let _278_v4;
      let _nw34 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
      _278_v4 = _nw34;
      let _279_v5;
      _279_v5 = _dafny.Seq.of(_276_v2);
      let _280_v6;
      _280_v6 = _dafny.Map.Empty.slice().updateUnsafe(_277_v3,_276_v2);
      let _281_v7;
      _281_v7 = _dafny.Seq.of(_277_v3, _277_v3);
      let _282_v8;
      let _nw35 = Array((new BigNumber(27)).toNumber());
      _nw35[(_dafny.ZERO).toNumber()] = new BigNumber((_279_v5).length);
      _nw35[(_dafny.ONE).toNumber()] = new BigNumber(51);
      _nw35[(new BigNumber(2)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(3)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(4)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(5)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(6)).toNumber()] = (((_280_v6).contains(!(_277_v3))) ? ((_280_v6).get(!(_277_v3))) : (_276_v2));
      _nw35[(new BigNumber(7)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(8)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(9)).toNumber()] = new BigNumber((_281_v7).length);
      _nw35[(new BigNumber(10)).toNumber()] = new BigNumber(-987);
      _nw35[(new BigNumber(11)).toNumber()] = new BigNumber(466);
      _nw35[(new BigNumber(12)).toNumber()] = new BigNumber((_274_v0).length);
      _nw35[(new BigNumber(13)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(14)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(15)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(16)).toNumber()] = new BigNumber(666);
      _nw35[(new BigNumber(17)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(18)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(19)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(20)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(21)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(22)).toNumber()] = new BigNumber((_274_v0).length);
      _nw35[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.of(_277_v3)).length);
      _nw35[(new BigNumber(24)).toNumber()] = _276_v2;
      _nw35[(new BigNumber(25)).toNumber()] = new BigNumber(-606);
      _nw35[(new BigNumber(26)).toNumber()] = _276_v2;
      _282_v8 = _nw35;
      let _283_v9;
      _283_v9 = _dafny.Seq.of(_282_v8);
      let _284_v10;
      _284_v10 = _dafny.MultiSet.fromElements(_282_v8);
      let _285_v11;
      let _nw36 = Array((new BigNumber(5)).toNumber());
      _nw36[(_dafny.ZERO).toNumber()] = false;
      _nw36[(_dafny.ONE).toNumber()] = _277_v3;
      _nw36[(new BigNumber(2)).toNumber()] = !(!(true));
      _nw36[(new BigNumber(3)).toNumber()] = _277_v3;
      _nw36[(new BigNumber(4)).toNumber()] = _277_v3;
      _285_v11 = _nw36;
      let _286_v12;
      _286_v12 = new _dafny.CodePoint('l'.codePointAt(0));
      let _287_v13;
      _287_v13 = _dafny.Map.Empty.slice().updateUnsafe(_285_v11,_286_v12);
      let _288_globalState;
      let _nw37 = new _module.GlobalState();
      _nw37.__ctor(new BigNumber(644), false, true, new BigNumber(888), new BigNumber(99), false, new BigNumber(83), false, new BigNumber(-470), _274_v0, _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_275_v1).length)), _dafny.Seq.of(_276_v2, _276_v2, (_dafny.ZERO).minus(_276_v2), new BigNumber((_dafny.MultiSet.fromElements(_277_v3, _277_v3, _277_v3)).cardinality()))), true, _274_v0, ((true) ? (_278_v4) : (_278_v4)), _dafny.Seq.Concat(_283_v9, _283_v9), new BigNumber(523), _284_v10, _287_v13);
      _288_globalState = _nw37;
      let _289_v14;
      _289_v14 = _dafny.Set.fromElements(_277_v3);
      let _hi0 = new BigNumber((_289_v14).length);
      for (let _290_i0 = _276_v2; _290_i0.isLessThan(_hi0); _290_i0 = _290_i0.plus(_dafny.ONE)) {
        if (_277_v3) {
          let _291_v15;
          _291_v15 = _dafny.Map.Empty.slice().updateUnsafe(_276_v2,!(_277_v3));
          let _292_v16;
          _292_v16 = _dafny.Set.fromElements(new BigNumber((_280_v6).length));
          _291_v15 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_279_v5).length)).plus(new BigNumber((_292_v16).length)),_277_v3);
          let _293_v17;
          _293_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_288_globalState),(_dafny.ZERO).minus(new BigNumber((_292_v16).length)));
          let _294_v18;
          _294_v18 = _dafny.Seq.of(_291_v15);
          _293_v17 = (_293_v17).update(_294_v18, _276_v2);
          let _295_v19;
          let _nw38 = Array((new BigNumber(9)).toNumber()).fill([]);
          _295_v19 = _nw38;
          let _296_v20;
          _296_v20 = _dafny.Map.Empty.slice().updateUnsafe(_276_v2,_285_v11);
          let _index17 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_295_v19).length));
          (_295_v19)[_index17] = (((_296_v20).contains(new BigNumber(152))) ? ((_296_v20).get(new BigNumber(152))) : (_285_v11));
          let _index18 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_295_v19).length));
          (_295_v19)[_index18] = _285_v11;
          (_288_globalState).f0 = _290_i0;
          let _297_v21;
          let _298_v22;
          let _299_v23;
          let _300_v24;
          let _out0;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = _module.__default.m0(_288_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _out3 = _outcollector0[3];
          _297_v21 = _out0;
          _298_v22 = _out1;
          _299_v23 = _out2;
          _300_v24 = _out3;
        } else {
          let _index19 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_285_v11).length));
          (_285_v11)[_index19] = _277_v3;
          let _index20 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_285_v11).length));
          (_285_v11)[_index20] = ((_277_v3) ? (_277_v3) : (!(true)));
          let _301_v25;
          let _302_v26;
          let _303_v27;
          let _304_v28;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0(_288_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _301_v25 = _out4;
          _302_v26 = _out5;
          _303_v27 = _out6;
          _304_v28 = _out7;
          let _index21 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_303_v27).length));
          (_303_v27)[_index21] = _301_v25;
          let _index22 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_303_v27).length));
          (_303_v27)[_index22] = _module.__default.fm1(((_277_v3) ? (_module.__default.fm1(new BigNumber(-45), (_285_v11)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_285_v11).length))], false, _288_globalState)) : (_276_v2)), _277_v3, _277_v3, _288_globalState);
          let _index23 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_285_v11).length));
          (_285_v11)[_index23] = (_277_v3) === (true);
          let _305_v29;
          _305_v29 = _dafny.Map.Empty.slice().updateUnsafe(_290_i0,(_303_v27)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_303_v27).length))]);
          let _index24 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_285_v11).length));
          let _rhs20 = ((_305_v29).Merge(_305_v29)).Merge(_305_v29);
          let _rhs21 = (_301_v25).isLessThanOrEqualTo((new BigNumber(748)).multipliedBy(_301_v25));
          let _rhs22 = _dafny.Seq.update(_281_v7, _module.__default.safeIndex((_dafny.ZERO).minus((new BigNumber(210)).minus(new BigNumber((_304_v28).length))), new BigNumber((_281_v7).length)), (_285_v11)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_285_v11).length))]);
          let _lhs11 = _285_v11;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_285_v11).length));
          _305_v29 = _rhs20;
          _lhs11[_lhs12] = _rhs21;
          _281_v7 = _rhs22;
        }
        _285_v11 = _285_v11;
        (_288_globalState).f3 = _276_v2;
        let _306_v30;
        _306_v30 = _dafny.Map.Empty.slice().updateUnsafe(_290_i0,_290_i0);
        let _307_v31;
        _307_v31 = _dafny.Seq.of(_306_v30);
        let _308_v32;
        let _nw39 = new _module.C4();
        _nw39.__ctor(new _dafny.CodePoint('p'.codePointAt(0)), _276_v2, _dafny.areEqual(_274_v0, _274_v0), _dafny.Seq.UnicodeFromString("yennni"), _dafny.Seq.Concat(_307_v31, _307_v31));
        _308_v32 = _nw39;
      }
      if (!(!((_276_v2).isEqualTo(_276_v2))) || (_277_v3)) {
        let _309_v33;
        let _310_v34;
        let _311_v35;
        let _312_v36;
        let _out8;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector2 = _module.__default.m0(_288_globalState);
        _out8 = _outcollector2[0];
        _out9 = _outcollector2[1];
        _out10 = _outcollector2[2];
        _out11 = _outcollector2[3];
        _309_v33 = _out8;
        _310_v34 = _out9;
        _311_v35 = _out10;
        _312_v36 = _out11;
        let _313_v37;
        let _314_v38;
        let _315_v39;
        let _316_v40;
        let _out12;
        let _out13;
        let _out14;
        let _out15;
        let _outcollector3 = _module.__default.m0(_288_globalState);
        _out12 = _outcollector3[0];
        _out13 = _outcollector3[1];
        _out14 = _outcollector3[2];
        _out15 = _outcollector3[3];
        _313_v37 = _out12;
        _314_v38 = _out13;
        _315_v39 = _out14;
        _316_v40 = _out15;
        _275_v1 = (_275_v1).update(_277_v3, _277_v3);
        _285_v11 = _285_v11;
        let _rhs23 = new BigNumber((_274_v0).length);
        let _rhs24 = _310_v34;
        let _lhs13 = _288_globalState;
        _310_v34 = _rhs23;
        _lhs13.f0 = _rhs24;
      } else {
        let _317_v41;
        _317_v41 = _module.D9.create_DC23(_277_v3, _277_v3, (_dafny.ZERO).minus(_276_v2));
        let _318_v42;
        _318_v42 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_277_v3, _277_v3)).length));
        let _319_v43;
        _319_v43 = _dafny.MultiSet.fromElements((_317_v41).dtor_cf44, _277_v3, _module.__default.fm23(_318_v42, _276_v2, _288_globalState), _277_v3);
        let _320_v44;
        _320_v44 = _module.D0.create_DC0(_277_v3);
        (_288_globalState).f0 = (new BigNumber((_module.__default.fm47(_319_v43, (_dafny.ZERO).minus(_276_v2), _276_v2, (_320_v44).dtor_cf0, _288_globalState)).length)).multipliedBy(_276_v2);
        (_288_globalState).f3 = new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(404)), ((_321_v12) => function (_322_i1) {
          return _321_v12;
        })(_286_v12)), _274_v0)).length);
        let _323_v45;
        let _nw40 = Array((new BigNumber(15)).toNumber()).fill([]);
        _323_v45 = _nw40;
        let _324_v46;
        let _nw41 = Array((_dafny.ONE).toNumber());
        _nw41[(_dafny.ZERO).toNumber()] = _285_v11;
        _324_v46 = _nw41;
        let _325_v47;
        _325_v47 = _module.D26.create_DC71(new BigNumber((_289_v14).length), _324_v46, _276_v2);
        let _326_v48;
        _326_v48 = _dafny.MultiSet.fromElements(_276_v2);
        let _327_v49;
        _327_v49 = _dafny.Seq.of(_323_v45, _323_v45);
        let _rhs25 = _277_v3;
        let _rhs26 = new BigNumber((_dafny.Seq.of(_277_v3, ((_325_v47).dtor_cf128).isLessThanOrEqualTo(_276_v2), _277_v3)).length);
        let _rhs27 = (((_326_v48).IsSubsetOf(_dafny.MultiSet.fromElements(_276_v2))) ? (_323_v45) : ((_327_v49)[_module.__default.safeIndex(_276_v2, new BigNumber((_327_v49).length))]));
        let _lhs14 = _288_globalState;
        let _lhs15 = _288_globalState;
        _lhs14.f2 = _rhs25;
        _lhs15.f0 = _rhs26;
        _323_v45 = _rhs27;
        let _328_v50;
        _328_v50 = _dafny.Map.Empty.slice().updateUnsafe(_276_v2,_module.__default.safeModuloInt(_276_v2, _276_v2));
        _328_v50 = _dafny.Map.Empty.slice().updateUnsafe(_276_v2,new BigNumber(30));
        (_288_globalState).f9 = _dafny.Seq.Concat(_274_v0, _274_v0);
      }
      if ((_dafny.MultiSet.fromElements(_277_v3)).contains((_277_v3) === (_277_v3))) {
        _277_v3 = _277_v3;
        let _329_v51;
        _329_v51 = _dafny.Map.Empty.slice().updateUnsafe(_277_v3,_286_v12);
        let _330_v52;
        _330_v52 = _dafny.Map.Empty.slice().updateUnsafe(_329_v51,_277_v3);
        (_288_globalState).f2 = (((_330_v52).contains(_dafny.Map.Empty.slice().updateUnsafe(_277_v3,_286_v12))) ? ((_330_v52).get(_dafny.Map.Empty.slice().updateUnsafe(_277_v3,_286_v12))) : (_277_v3));
        (_288_globalState).f15 = _module.__default.safeModuloInt(_276_v2, _276_v2);
        let _index25 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_282_v8).length));
        (_282_v8)[_index25] = new BigNumber((_dafny.Seq.Concat(_279_v5, _279_v5)).length);
        let _index26 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_282_v8).length));
        (_282_v8)[_index26] = _276_v2;
        let _331_v53;
        let _332_v54;
        let _333_v55;
        let _334_v56;
        let _out16;
        let _out17;
        let _out18;
        let _out19;
        let _outcollector4 = _module.__default.m0(_288_globalState);
        _out16 = _outcollector4[0];
        _out17 = _outcollector4[1];
        _out18 = _outcollector4[2];
        _out19 = _outcollector4[3];
        _331_v53 = _out16;
        _332_v54 = _out17;
        _333_v55 = _out18;
        _334_v56 = _out19;
      } else {
        let _335_v57;
        _335_v57 = _module.D1.create_DC4(_277_v3);
        let _336_v58;
        _336_v58 = _module.D2.create_DC7(_module.__default.fm1(new BigNumber(135), (_335_v57).dtor_cf12, _277_v3, _288_globalState), _277_v3, ((_dafny.ZERO).minus(new BigNumber(-40))).plus(new BigNumber((_274_v0).length)));
        _336_v58 = _336_v58;
        (_288_globalState).f2 = _277_v3;
        (_288_globalState).f12 = _dafny.Seq.Concat(_274_v0, _274_v0);
        let _337_v59;
        _337_v59 = _module.D1.create_DC3((_dafny.ZERO).minus(_276_v2), new BigNumber((_274_v0).length), _276_v2, _276_v2, _276_v2);
        let _338_v60;
        let _nw42 = new _module.C5();
        _nw42.__ctor(_274_v0, _274_v0, _277_v3);
        _338_v60 = _nw42;
        let _339_v61;
        _339_v61 = _module.D1.create_DC5(((_277_v3) ? (_337_v59) : (_module.D1.create_DC2(_338_v60))));
        _339_v61 = _339_v61;
        _275_v1 = (_275_v1).update((_338_v60).f18, (_338_v60).f18);
      }
      _286_v12 = _286_v12;
      let _hi1 = _276_v2;
      for (let _340_i2 = (_276_v2).minus(_276_v2); _340_i2.isLessThan(_hi1); _340_i2 = _340_i2.plus(_dafny.ONE)) {
        let _341_v62;
        let _nw43 = Array((new BigNumber(21)).toNumber());
        _341_v62 = _nw43;
        let _342_v63;
        _342_v63 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_341_v62);
        let _343_v64;
        _343_v64 = _dafny.Set.fromElements((_dafny.ZERO).minus(_276_v2), new BigNumber((_342_v63).length));
        if (_module.__default.fm23(_343_v64, _276_v2, _288_globalState)) {
          (_288_globalState).f2 = !(_277_v3);
          _277_v3 = (_module.D0.create_DC1(_277_v3, new BigNumber((_274_v0).length), _274_v0, _277_v3, _340_i2)).dtor_cf1;
          (_288_globalState).f0 = _340_i2;
          (_288_globalState).f3 = (_dafny.ZERO).minus(_340_i2);
          let _344_v65;
          _344_v65 = _dafny.MultiSet.fromElements(_340_i2, _340_i2);
          let _345_v66;
          let _nw44 = new _module.C8();
          _nw44.__ctor(new BigNumber((_344_v65).cardinality()), _277_v3);
          _345_v66 = _nw44;
          let _346_v67;
          let _nw45 = new _module.C9();
          _nw45.__ctor(_285_v11, _277_v3);
          _346_v67 = _nw45;
          let _rhs28 = (_346_v67).f18;
          let _rhs29 = _345_v66;
          let _rhs30 = _346_v67;
          let _rhs31 = _345_v66.f23;
          let _lhs16 = _288_globalState;
          _277_v3 = _rhs28;
          _345_v66 = _rhs29;
          _346_v67 = _rhs30;
          _lhs16.f15 = _rhs31;
        } else {
          let _347_v68;
          let _nw46 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
          _347_v68 = _nw46;
          let _index27 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_347_v68).length));
          (_347_v68)[_index27] = _343_v64;
          let _348_v69;
          _348_v69 = _dafny.Seq.of(_279_v5, _279_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-267)), ((_349_i2) => function (_350_i3) {
            return _349_i2;
          })(_340_i2)), _279_v5);
          let _index28 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_347_v68).length));
          let _rhs32 = _343_v64;
          let _rhs33 = _module.__default.fm1(_340_i2, _277_v3, false, _288_globalState);
          let _rhs34 = _348_v69;
          let _rhs35 = _282_v8;
          let _lhs17 = _347_v68;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_347_v68).length));
          let _lhs19 = _288_globalState;
          _lhs17[_lhs18] = _rhs32;
          _lhs19.f15 = _rhs33;
          _348_v69 = _rhs34;
          _282_v8 = _rhs35;
          let _index29 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_282_v8).length));
          (_282_v8)[_index29] = (_340_i2).plus(_276_v2);
          let _351_v70;
          _351_v70 = _dafny.MultiSet.fromElements((function (_pat_let2_0) {
            return function (_352_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_353_dt__update_hcf33_h0) {
                  return _module.D3.create_DC13((_352_dt__update__tmp_h0).dtor_cf30, (_352_dt__update__tmp_h0).dtor_cf31, (_352_dt__update__tmp_h0).dtor_cf32, _353_dt__update_hcf33_h0, (_352_dt__update__tmp_h0).dtor_cf34);
                }(_pat_let3_0);
              }(new BigNumber(828));
            }(_pat_let2_0);
          }(_module.D3.create_DC13(_276_v2, _dafny.Seq.UnicodeFromString("ppao"), _340_i2, new BigNumber(281), _286_v12))).dtor_cf31);
          let _index30 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_282_v8).length));
          (_282_v8)[_index30] = ((!(true)) ? (((_277_v3) ? (_276_v2) : (_340_i2))) : ((new BigNumber((_351_v70).cardinality())).minus(_276_v2)));
          let _354_v71;
          let _nw47 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _354_v71 = _nw47;
          let _355_v72;
          _355_v72 = _dafny.Set.fromElements(_354_v71, _354_v71, _354_v71, ((_277_v3) ? (_354_v71) : (_354_v71)));
          let _356_v73;
          let _init1 = ((_357_v0) => function (_358_i4) {
            return _357_v0;
          })(_274_v0);
          let _nw48 = Array((new BigNumber(23)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw48.length); _i0_1++) {
            _nw48[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _356_v73 = _nw48;
          let _index31 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_356_v73).length));
          (_356_v73)[_index31] = _dafny.Seq.Concat(_274_v0, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), ((_359_v12) => function (_360_i5) {
            return _359_v12;
          })(_286_v12)), _module.__default.safeIndex(_276_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), ((_361_v12) => function (_362_i5) {
            return _361_v12;
          })(_286_v12))).length)), _286_v12));
          let _index32 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_356_v73).length));
          let _rhs36 = _module.__default.safeModuloInt((_282_v8)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_282_v8).length))], _276_v2);
          let _rhs37 = _355_v72;
          let _rhs38 = (_279_v5)[_module.__default.safeIndex((_282_v8)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_282_v8).length))], new BigNumber((_279_v5).length))];
          let _rhs39 = _module.__default.fm16((new BigNumber((_274_v0).length)).plus(new BigNumber(661)), (_289_v14).IsDisjointFrom(_289_v14), _286_v12, (_282_v8)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_282_v8).length))], _288_globalState);
          let _lhs20 = _288_globalState;
          let _lhs21 = _288_globalState;
          let _lhs22 = _356_v73;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_356_v73).length));
          _lhs20.f15 = _rhs36;
          _355_v72 = _rhs37;
          _lhs21.f0 = _rhs38;
          _lhs22[_lhs23] = _rhs39;
          let _363_v74;
          _363_v74 = _dafny.Set.fromElements(_279_v5);
          let _364_v75;
          let _nw49 = new _module.C6();
          _nw49.__ctor(_281_v7, _340_i2, !(_277_v3), _dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), ((_365_v12) => function (_366_i6) {
            return _365_v12;
          })(_286_v12)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(999)), ((_367_i2) => function (_368_i7) {
            return _dafny.Map.Empty.slice().updateUnsafe(_367_i2,_367_i2);
          })(_340_i2)));
          _364_v75 = _nw49;
          let _369_v76;
          _369_v76 = _dafny.Map.Empty.slice().updateUnsafe((_363_v74).IsProperSubsetOf(_363_v74),_dafny.Map.Empty.slice().updateUnsafe(_364_v75,_277_v3));
          let _370_v77;
          _370_v77 = _dafny.Map.Empty.slice().updateUnsafe(_277_v3,_369_v76);
          let _371_v78;
          _371_v78 = _module.D29.create_DC76((((_370_v77).contains(_277_v3)) ? ((_370_v77).get(_277_v3)) : (_369_v76)));
          let _pat_let_tv0 = _369_v76;
          _369_v76 = (_369_v76).Merge((function (_pat_let4_0) {
            return function (_372_dt__update__tmp_h1) {
              return function (_pat_let5_0) {
                return function (_373_dt__update_hcf132_h0) {
                  return _module.D29.create_DC76(_373_dt__update_hcf132_h0);
                }(_pat_let5_0);
              }(_pat_let_tv0);
            }(_pat_let4_0);
          }(_371_v78)).dtor_cf132);
          let _374_v79;
          _374_v79 = _dafny.MultiSet.fromElements(new BigNumber(661));
          let _375_v80;
          let _nw50 = new _module.C7();
          _nw50.__ctor((((_374_v79).contains(_340_i2)) ? ((_374_v79).get(_340_i2)) : (new BigNumber((_dafny.Seq.of(_340_i2)).length))), !_dafny.Seq.contains(_274_v0, _286_v12));
          _375_v80 = _nw50;
          _375_v80 = _375_v80;
        }
        let _376_v81;
        _376_v81 = _dafny.Map.Empty.slice().updateUnsafe(_277_v3,_286_v12);
        let _377_v82;
        _377_v82 = _dafny.Map.Empty.slice().updateUnsafe(_277_v3,_376_v81);
        let _378_v83;
        _378_v83 = _module.D25.create_DC66(_377_v82);
        let _379_v84;
        _379_v84 = _module.D25.create_DC68(_378_v83);
        let _source3 = _379_v84;
        if (_source3.is_DC67) {
          let _380___mcc_h0 = (_source3).cf116;
          let _381___mcc_h1 = (_source3).cf117;
          let _382___mcc_h2 = (_source3).cf118;
          let _383___mcc_h3 = (_source3).cf119;
          let _384_cf119 = _383___mcc_h3;
          let _385_cf118 = _382___mcc_h2;
          let _386_cf117 = _381___mcc_h1;
          let _387_cf116 = _380___mcc_h0;
          let _388_v85;
          _388_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(580),!(_384_cf119));
          let _389_v86;
          _389_v86 = _dafny.Seq.of(_388_v85);
          _388_v85 = (_389_v86)[_module.__default.safeIndex(_276_v2, new BigNumber((_389_v86).length))];
          let _390_v87;
          let _391_v88;
          let _392_v89;
          let _393_v90;
          let _out20;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector5 = _module.__default.m0(_288_globalState);
          _out20 = _outcollector5[0];
          _out21 = _outcollector5[1];
          _out22 = _outcollector5[2];
          _out23 = _outcollector5[3];
          _390_v87 = _out20;
          _391_v88 = _out21;
          _392_v89 = _out22;
          _393_v90 = _out23;
          (_288_globalState).f2 = _384_cf119;
          let _index33 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_285_v11).length));
          (_285_v11)[_index33] = !_dafny.Seq.contains(_274_v0, new _dafny.CodePoint('b'.codePointAt(0)));
          let _index34 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_285_v11).length));
          (_285_v11)[_index34] = (_391_v88).isLessThanOrEqualTo(new BigNumber(242));
        } else if (_source3.is_DC66) {
          let _394___mcc_h4 = (_source3).cf115;
          let _395_cf115 = _394___mcc_h4;
          _289_v14 = _289_v14;
          (_288_globalState).f2 = !(_277_v3);
          let _396_v91;
          _396_v91 = _module.D17.create_DC44(_285_v11, _276_v2);
          let _397_v92;
          _397_v92 = _dafny.Set.fromElements(_285_v11, (_396_v91).dtor_cf77);
          (_288_globalState).f2 = (_397_v92).IsSubsetOf((_397_v92).Difference(_dafny.Set.fromElements(_285_v11)));
          let _398_v93;
          _398_v93 = _dafny.Map.Empty.slice().updateUnsafe(_340_i2,_285_v11);
          _398_v93 = _398_v93;
        } else {
          let _399___mcc_h5 = (_source3).cf120;
          let _400_cf120 = _399___mcc_h5;
          let _401_v94;
          _401_v94 = _dafny.MultiSet.fromElements(_286_v12, new _dafny.CodePoint('r'.codePointAt(0)));
          let _402_v95;
          let _init2 = ((_403_v2, _404_v0) => function (_405_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe(_403_v2,_module.D16.create_DC40(_403_v2, new BigNumber(-738), _403_v2, new BigNumber((_404_v0).length)));
          })(_276_v2, _274_v0);
          let _nw51 = Array((new BigNumber(14)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw51.length); _i0_2++) {
            _nw51[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _402_v95 = _nw51;
          let _406_v96;
          _406_v96 = _module.D26.create_DC70(_340_i2, _340_i2, _277_v3, _402_v95);
          (_288_globalState).f9 = _dafny.Seq.update(_dafny.Seq.Concat(_274_v0, _dafny.Seq.Concat(_274_v0, _274_v0)), _module.__default.safeIndex(((((_401_v94).contains(_286_v12)) ? ((_401_v94).get(_286_v12)) : (_276_v2))).multipliedBy((_dafny.ZERO).minus((_406_v96).dtor_cf123)), new BigNumber((_dafny.Seq.Concat(_274_v0, _dafny.Seq.Concat(_274_v0, _274_v0))).length)), _286_v12);
          let _407_v97;
          _407_v97 = _282_v8;
          let _408_v98;
          let _nw52 = Array((new BigNumber(20)).toNumber());
          _nw52[(_dafny.ZERO).toNumber()] = _282_v8;
          _nw52[(_dafny.ONE).toNumber()] = _282_v8;
          _nw52[(new BigNumber(2)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(3)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(4)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(5)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(6)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(7)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(8)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(9)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(10)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(11)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(12)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(13)).toNumber()] = (_407_v97);
          _nw52[(new BigNumber(14)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(15)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(16)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(17)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(18)).toNumber()] = _282_v8;
          _nw52[(new BigNumber(19)).toNumber()] = _282_v8;
          _408_v98 = _nw52;
          _408_v98 = _408_v98;
          let _409_v99;
          _409_v99 = _dafny.Map.Empty.slice().updateUnsafe(_340_i2,_277_v3);
          _409_v99 = (_409_v99).update(new BigNumber((_279_v5).length), true);
          let _410_v100;
          _410_v100 = _module.D2.create_DC7(_276_v2, _277_v3, _340_i2);
          let _411_v101;
          let _nw53 = new _module.C2();
          _nw53.__ctor(new BigNumber((_343_v64).length), _module.__default.fm40(_410_v100, _288_globalState), _277_v3);
          _411_v101 = _nw53;
          let _412_v102;
          _412_v102 = _dafny.Seq.of(_411_v101, _411_v101, _411_v101);
          let _413_v103;
          _413_v103 = _dafny.Map.Empty.slice().updateUnsafe((_412_v102)[_module.__default.safeIndex((_dafny.ZERO).minus(_340_i2), new BigNumber((_412_v102).length))],new BigNumber(155));
          _413_v103 = (_413_v103).update(_411_v101, _340_i2);
        }
        let _414_v104;
        let _415_v105;
        let _416_v106;
        let _417_v107;
        let _out24;
        let _out25;
        let _out26;
        let _out27;
        let _outcollector6 = _module.__default.m0(_288_globalState);
        _out24 = _outcollector6[0];
        _out25 = _outcollector6[1];
        _out26 = _outcollector6[2];
        _out27 = _outcollector6[3];
        _414_v104 = _out24;
        _415_v105 = _out25;
        _416_v106 = _out26;
        _417_v107 = _out27;
        (_288_globalState).f2 = !(_277_v3);
      }
      let _index35 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_282_v8).length));
      (_282_v8)[_index35] = _276_v2;
      let _index36 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_282_v8).length));
      (_282_v8)[_index36] = _276_v2;
      let _418_v108;
      let _nw54 = new _module.C7();
      _nw54.__ctor((_279_v5)[_module.__default.safeIndex(_276_v2, new BigNumber((_279_v5).length))], _277_v3);
      _418_v108 = _nw54;
      _418_v108 = _418_v108;
      (_288_globalState).f0 = (_282_v8)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_282_v8).length))];
      (_288_globalState).f15 = _276_v2;
      let _419_v109;
      _419_v109 = _dafny.Map.Empty.slice().updateUnsafe(_277_v3,_274_v0);
      let _420_v110;
      let _nw55 = Array((new BigNumber(11)).toNumber());
      _nw55[(_dafny.ZERO).toNumber()] = _274_v0;
      _nw55[(_dafny.ONE).toNumber()] = _274_v0;
      _nw55[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-354)), ((_421_v12) => function (_422_i9) {
        return _421_v12;
      })(_286_v12));
      _nw55[(new BigNumber(3)).toNumber()] = _274_v0;
      _nw55[(new BigNumber(4)).toNumber()] = _274_v0;
      _nw55[(new BigNumber(5)).toNumber()] = _274_v0;
      _nw55[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((((_419_v109).contains(_277_v3)) ? ((_419_v109).get(_277_v3)) : (_dafny.Seq.UnicodeFromString("fbdttde"))), _274_v0);
      _nw55[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_274_v0, _274_v0);
      _nw55[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("mp");
      _nw55[(new BigNumber(9)).toNumber()] = ((_277_v3) ? (_dafny.Seq.UnicodeFromString("sstrrcv")) : (_274_v0));
      _nw55[(new BigNumber(10)).toNumber()] = _274_v0;
      _420_v110 = _nw55;
      let _index37 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_420_v110).length));
      (_420_v110)[_index37] = _274_v0;
      let _index38 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_420_v110).length));
      (_420_v110)[_index38] = _274_v0;
      let _423_v111;
      _423_v111 = _module.D17.create_DC42(_277_v3);
      let _424_v112;
      _424_v112 = _dafny.Map.Empty.slice().updateUnsafe(_276_v2,(_423_v111).dtor_cf74);
      let _hi2 = ((_277_v3) ? (_418_v108.f24) : (new BigNumber(130)));
      for (let _425_i10 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(_277_v3)).length), new BigNumber((_424_v112).length)); _425_i10.isLessThan(_hi2); _425_i10 = _425_i10.plus(_dafny.ONE)) {
        (_418_v108).m1(_dafny.Seq.UnicodeFromString("scoeuba"), _425_i10, _dafny.Seq.Concat((_420_v110)[_module.__default.safeIndex(new BigNumber(283), new BigNumber((_420_v110).length))], _274_v0), _288_globalState);
        let _426_v113;
        let _nw56 = new _module.C5();
        _nw56.__ctor((_420_v110)[_module.__default.safeIndex(new BigNumber(283), new BigNumber((_420_v110).length))], (_420_v110)[_module.__default.safeIndex(new BigNumber(283), new BigNumber((_420_v110).length))], _277_v3);
        _426_v113 = _nw56;
        let _427_v114;
        _427_v114 = _dafny.Map.Empty.slice().updateUnsafe((_282_v8)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_282_v8).length))],_276_v2);
        let _428_v115;
        _428_v115 = _dafny.Map.Empty.slice().updateUnsafe(_426_v113,new BigNumber((_427_v114).length));
        let _429_v116;
        _429_v116 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_277_v3,new BigNumber((_428_v115).length)));
        (_288_globalState).f2 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_280_v6, _280_v6), _429_v116));
        (_426_v113).m7((_425_i10).isLessThanOrEqualTo(new BigNumber((_279_v5).length)), _418_v108.f24, _288_globalState);
        let _430_v117;
        let _nw57 = Array((new BigNumber(4)).toNumber());
        _nw57[(_dafny.ZERO).toNumber()] = (_424_v112).Merge(_424_v112);
        _nw57[(_dafny.ONE).toNumber()] = _424_v112;
        _nw57[(new BigNumber(2)).toNumber()] = _424_v112;
        _nw57[(new BigNumber(3)).toNumber()] = _424_v112;
        _430_v117 = _nw57;
        let _init3 = ((_431_v112) => function (_432_i11) {
          return _431_v112;
        })(_424_v112);
        let _nw58 = Array((new BigNumber(14)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw58.length); _i0_3++) {
          _nw58[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _430_v117 = _nw58;
      }
      let _433_v118;
      _433_v118 = _dafny.Map.Empty.slice().updateUnsafe(_286_v12,(_420_v110)[_module.__default.safeIndex(new BigNumber(283), new BigNumber((_420_v110).length))]);
      let _434_v119;
      let _nw59 = Array((new BigNumber(24)).toNumber());
      _434_v119 = _nw59;
      let _435_v120;
      _435_v120 = _dafny.Map.Empty.slice().updateUnsafe(_434_v119,(_420_v110)[_module.__default.safeIndex(new BigNumber(283), new BigNumber((_420_v110).length))]);
      _433_v118 = (_433_v118).update(_286_v12, (((_435_v120).contains(_434_v119)) ? ((_435_v120).get(_434_v119)) : (_dafny.Seq.UnicodeFromString("malcvx"))));
      let _index39 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_285_v11).length));
      (_285_v11)[_index39] = _277_v3;
      let _436_v121;
      let _init4 = ((_437_v12) => function (_438_i12) {
        return _437_v12;
      })(_286_v12);
      let _nw60 = Array((new BigNumber(8)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw60.length); _i0_4++) {
        _nw60[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _436_v121 = _nw60;
      let _index40 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_436_v121).length));
      (_436_v121)[_index40] = _286_v12;
      let _439_v122;
      _439_v122 = _dafny.MultiSet.fromElements((_277_v3) === (!(true)));
      let _440_v123;
      _440_v123 = _dafny.Set.fromElements(_418_v108.f24, _276_v2, new BigNumber(469), _276_v2);
      let _index41 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_285_v11).length));
      let _index42 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_436_v121).length));
      let _rhs40 = (((_277_v3) ? (_439_v122) : (_439_v122))).IsDisjointFrom(_dafny.MultiSet.fromElements(false));
      let _rhs41 = _286_v12;
      let _rhs42 = (_dafny.MultiSet.FromArray(((false) ? (_281_v7) : (_281_v7)))).Union(_439_v122);
      let _rhs43 = _module.__default.fm23(_440_v123, new BigNumber(-471), _288_globalState);
      let _rhs44 = _dafny.Seq.IsProperPrefixOf(_274_v0, _274_v0);
      let _lhs24 = _285_v11;
      let _lhs25 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_285_v11).length));
      let _lhs26 = _436_v121;
      let _lhs27 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_436_v121).length));
      let _lhs28 = _288_globalState;
      _lhs24[_lhs25] = _rhs40;
      _lhs26[_lhs27] = _rhs41;
      _439_v122 = _rhs42;
      _277_v3 = _rhs43;
      _lhs28.f2 = _rhs44;
      (_288_globalState).f2 = (_423_v111).dtor_cf74;
      let _441_v124;
      _441_v124 = _dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe((_285_v11)[_module.__default.safeIndex(new BigNumber(635), new BigNumber((_285_v11).length))],new BigNumber(871))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_285_v11)[_module.__default.safeIndex(new BigNumber(635), new BigNumber((_285_v11).length))],new BigNumber((_439_v122).cardinality()))));
      _441_v124 = (_441_v124).update(_dafny.Map.Empty.slice().updateUnsafe(false,(_282_v8)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_282_v8).length))]), _module.__default.abs(_276_v2));
      let _index43 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_282_v8).length));
      (_282_v8)[_index43] = (_418_v108.f24).plus(_module.__default.fm1(_276_v2, _277_v3, !(_277_v3), _288_globalState));
      process.stdout.write((_274_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v1).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true).updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_276_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_277_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_279_v5, _dafny.Seq.of(new BigNumber(-284)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_280_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_281_v7, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v8)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_283_v9).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_284_v10).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_285_v11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_285_v11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_285_v11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_285_v11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_285_v11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_286_v12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_287_v13).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_288_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_288_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_288_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_288_globalState.f9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_288_globalState.f10, _dafny.Seq.of(_dafny.ONE, new BigNumber(-284), new BigNumber(-284), new BigNumber(284), new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_288_globalState.f12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_288_globalState).f14).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_288_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_288_globalState.f16).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_288_globalState.f17).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_v14).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_418_v108.f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_419_v109).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("ewld")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_420_v110)[new BigNumber(2)], _dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_420_v110)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_423_v111).dtor_cf74));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v112).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-284),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_433_v118).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),_dafny.Seq.UnicodeFromString("ewld")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_435_v120).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_436_v121)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_436_v121)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_436_v121)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_436_v121)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_436_v121)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_436_v121)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_436_v121)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_436_v121)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_439_v122).equals(_dafny.MultiSet.fromElements(false, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_440_v123).equals(_dafny.Set.fromElements(new BigNumber(-284), new BigNumber(469)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_441_v124).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(3)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-284))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4, cf5) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
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
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + this.cf3.toVerbatimString(true) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO);
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
    static create_DC3(cf7, cf8, cf9, cf10, cf11) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC4(cf12) {
      let $dt = new D1(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC2(cf6) {
      let $dt = new D1(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf13) {
      let $dt = new D1(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf12 === other.cf12;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf6 === other.cf6;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC7(cf15, cf16, cf17) {
      let $dt = new D2(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC8(cf18, cf19, cf20) {
      let $dt = new D2(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC6(cf14) {
      let $dt = new D2(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC9(cf21) {
      let $dt = new D2(3);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC9() { return this.$tag === 3; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf18 === other.cf18 && this.cf19 === other.cf19 && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC11(cf23, cf24) {
      let $dt = new D3(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC12(cf25, cf26, cf27, cf28, cf29) {
      let $dt = new D3(1);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC13(cf30, cf31, cf32, cf33, cf34) {
      let $dt = new D3(2);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC10(cf22) {
      let $dt = new D3(3);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC14(cf35) {
      let $dt = new D3(4);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get is_DC10() { return this.$tag === 3; }
    get is_DC14() { return this.$tag === 4; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC13" + "(" + _dafny.toString(this.cf30) + ", " + this.cf31.toVerbatimString(true) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 4) {
        return "D3.DC14" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23 && this.cf24 === other.cf24;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf25 === other.cf25 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf22 === other.cf22;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(false, false);
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
    static create_DC15(cf36) {
      let $dt = new D4(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
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
    static create_DC16(cf37) {
      let $dt = new D5(0);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf37 === other.cf37;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
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
    static create_DC18() {
      let $dt = new D6(0);
      return $dt;
    }
    static create_DC17(cf38) {
      let $dt = new D6(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC18";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf38 === other.cf38;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC18();
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
    static create_DC19(cf39) {
      let $dt = new D7(0);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf39 === other.cf39;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
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
    static create_DC21(cf41, cf42) {
      let $dt = new D8(0);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC20(cf40) {
      let $dt = new D8(1);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC23(cf44, cf45, cf46) {
      let $dt = new D9(0);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC22(cf43) {
      let $dt = new D9(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf44 === other.cf44 && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC23(false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf48) {
      let $dt = new D10(0);
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC24(cf47) {
      let $dt = new D10(1);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC26(cf49) {
      let $dt = new D10(2);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf48 === other.cf48;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC25(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf51) {
      let $dt = new D11(0);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC27(cf50) {
      let $dt = new D11(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC29(cf52) {
      let $dt = new D11(2);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf51 === other.cf51;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC28(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC31() {
      let $dt = new D12(0);
      return $dt;
    }
    static create_DC30(cf53) {
      let $dt = new D12(1);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC31";
      } else if (this.$tag === 1) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf53, other.cf53);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC31();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC33(cf55, cf56, cf57, cf58) {
      let $dt = new D13(0);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC34(cf59, cf60, cf61, cf62) {
      let $dt = new D13(1);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC32(cf54) {
      let $dt = new D13(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC35(cf63) {
      let $dt = new D13(3);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get is_DC35() { return this.$tag === 3; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf59) + ", " + this.cf60.toVerbatimString(true) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC32" + "(" + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf55 === other.cf55 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf59 === other.cf59 && _dafny.areEqual(this.cf60, other.cf60) && this.cf61 === other.cf61 && this.cf62 === other.cf62;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf63, other.cf63);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC33(false, _dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC36(cf64) {
      let $dt = new D14(0);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC37(cf65) {
      let $dt = new D15(0);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf65, other.cf65);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC39(cf67, cf68) {
      let $dt = new D16(0);
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC40(cf69, cf70, cf71, cf72) {
      let $dt = new D16(1);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC38(cf66) {
      let $dt = new D16(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC40" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC38" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf67 === other.cf67 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70) && _dafny.areEqual(this.cf71, other.cf71) && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC39(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC42(cf74) {
      let $dt = new D17(0);
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC43(cf75, cf76) {
      let $dt = new D17(1);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC44(cf77, cf78) {
      let $dt = new D17(2);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC41(cf73) {
      let $dt = new D17(3);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get is_DC41() { return this.$tag === 3; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC42" + "(" + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC43" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC44" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC41" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf75, other.cf75) && this.cf76 === other.cf76;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf77 === other.cf77 && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC42(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC46(cf80) {
      let $dt = new D18(0);
      $dt.cf80 = cf80;
      return $dt;
    }
    static create_DC45(cf79) {
      let $dt = new D18(1);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC46" + "(" + _dafny.toString(this.cf80) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC45" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf80 === other.cf80;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC46([]);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC48(cf82) {
      let $dt = new D19(0);
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC49(cf83) {
      let $dt = new D19(1);
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC50(cf84) {
      let $dt = new D19(2);
      $dt.cf84 = cf84;
      return $dt;
    }
    static create_DC47(cf81) {
      let $dt = new D19(3);
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC51(cf85) {
      let $dt = new D19(4);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get is_DC47() { return this.$tag === 3; }
    get is_DC51() { return this.$tag === 4; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC49" + "(" + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC50" + "(" + _dafny.toString(this.cf84) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC47" + "(" + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 4) {
        return "D19.DC51" + "(" + _dafny.toString(this.cf85) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf82 === other.cf82;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf83 === other.cf83;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf84 === other.cf84;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf85, other.cf85);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC48(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC53(cf87, cf88, cf89) {
      let $dt = new D20(0);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      return $dt;
    }
    static create_DC52(cf86) {
      let $dt = new D20(1);
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC54(cf90) {
      let $dt = new D20(2);
      $dt.cf90 = cf90;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC52() { return this.$tag === 1; }
    get is_DC54() { return this.$tag === 2; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf90() { return this.cf90; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC53" + "(" + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC52" + "(" + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC54" + "(" + _dafny.toString(this.cf90) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf87 === other.cf87 && this.cf88 === other.cf88 && _dafny.areEqual(this.cf89, other.cf89);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf86, other.cf86);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf90, other.cf90);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC53(false, false, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC56(cf92, cf93, cf94) {
      let $dt = new D21(0);
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC55(cf91) {
      let $dt = new D21(1);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC56" + "(" + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC55" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf92 === other.cf92 && _dafny.areEqual(this.cf93, other.cf93) && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf91, other.cf91);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC56(false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC58(cf96) {
      let $dt = new D22(0);
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC57(cf95) {
      let $dt = new D22(1);
      $dt.cf95 = cf95;
      return $dt;
    }
    static create_DC59(cf97) {
      let $dt = new D22(2);
      $dt.cf97 = cf97;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC59() { return this.$tag === 2; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf97() { return this.cf97; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC58" + "(" + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC57" + "(" + _dafny.toString(this.cf95) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC59" + "(" + _dafny.toString(this.cf97) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf95 === other.cf95;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf97, other.cf97);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC58(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC61(cf99, cf100, cf101, cf102) {
      let $dt = new D23(0);
      $dt.cf99 = cf99;
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      return $dt;
    }
    static create_DC62(cf103) {
      let $dt = new D23(1);
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC60(cf98) {
      let $dt = new D23(2);
      $dt.cf98 = cf98;
      return $dt;
    }
    get is_DC61() { return this.$tag === 0; }
    get is_DC62() { return this.$tag === 1; }
    get is_DC60() { return this.$tag === 2; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf98() { return this.cf98; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC61" + "(" + _dafny.toString(this.cf99) + ", " + _dafny.toString(this.cf100) + ", " + _dafny.toString(this.cf101) + ", " + _dafny.toString(this.cf102) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC62" + "(" + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC60" + "(" + _dafny.toString(this.cf98) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf99 === other.cf99 && this.cf100 === other.cf100 && this.cf101 === other.cf101 && this.cf102 === other.cf102;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf103, other.cf103);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf98, other.cf98);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC61([], false, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC64(cf105, cf106, cf107, cf108, cf109) {
      let $dt = new D24(0);
      $dt.cf105 = cf105;
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      $dt.cf109 = cf109;
      return $dt;
    }
    static create_DC65(cf110, cf111, cf112, cf113, cf114) {
      let $dt = new D24(1);
      $dt.cf110 = cf110;
      $dt.cf111 = cf111;
      $dt.cf112 = cf112;
      $dt.cf113 = cf113;
      $dt.cf114 = cf114;
      return $dt;
    }
    static create_DC63(cf104) {
      let $dt = new D24(2);
      $dt.cf104 = cf104;
      return $dt;
    }
    get is_DC64() { return this.$tag === 0; }
    get is_DC65() { return this.$tag === 1; }
    get is_DC63() { return this.$tag === 2; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf104() { return this.cf104; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC64" + "(" + _dafny.toString(this.cf105) + ", " + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ", " + _dafny.toString(this.cf108) + ", " + _dafny.toString(this.cf109) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC65" + "(" + _dafny.toString(this.cf110) + ", " + _dafny.toString(this.cf111) + ", " + _dafny.toString(this.cf112) + ", " + _dafny.toString(this.cf113) + ", " + _dafny.toString(this.cf114) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC63" + "(" + _dafny.toString(this.cf104) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf105 === other.cf105 && this.cf106 === other.cf106 && this.cf107 === other.cf107 && _dafny.areEqual(this.cf108, other.cf108) && _dafny.areEqual(this.cf109, other.cf109);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf110 === other.cf110 && _dafny.areEqual(this.cf111, other.cf111) && this.cf112 === other.cf112 && _dafny.areEqual(this.cf113, other.cf113) && this.cf114 === other.cf114;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf104, other.cf104);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC64(false, false, null, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC67(cf116, cf117, cf118, cf119) {
      let $dt = new D25(0);
      $dt.cf116 = cf116;
      $dt.cf117 = cf117;
      $dt.cf118 = cf118;
      $dt.cf119 = cf119;
      return $dt;
    }
    static create_DC66(cf115) {
      let $dt = new D25(1);
      $dt.cf115 = cf115;
      return $dt;
    }
    static create_DC68(cf120) {
      let $dt = new D25(2);
      $dt.cf120 = cf120;
      return $dt;
    }
    get is_DC67() { return this.$tag === 0; }
    get is_DC66() { return this.$tag === 1; }
    get is_DC68() { return this.$tag === 2; }
    get dtor_cf116() { return this.cf116; }
    get dtor_cf117() { return this.cf117; }
    get dtor_cf118() { return this.cf118; }
    get dtor_cf119() { return this.cf119; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf120() { return this.cf120; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC67" + "(" + _dafny.toString(this.cf116) + ", " + _dafny.toString(this.cf117) + ", " + _dafny.toString(this.cf118) + ", " + _dafny.toString(this.cf119) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC66" + "(" + _dafny.toString(this.cf115) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC68" + "(" + _dafny.toString(this.cf120) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf116 === other.cf116 && _dafny.areEqual(this.cf117, other.cf117) && _dafny.areEqual(this.cf118, other.cf118) && this.cf119 === other.cf119;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf115, other.cf115);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf120, other.cf120);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC67(false, _dafny.ZERO, _dafny.Seq.of(), false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC70(cf122, cf123, cf124, cf125) {
      let $dt = new D26(0);
      $dt.cf122 = cf122;
      $dt.cf123 = cf123;
      $dt.cf124 = cf124;
      $dt.cf125 = cf125;
      return $dt;
    }
    static create_DC71(cf126, cf127, cf128) {
      let $dt = new D26(1);
      $dt.cf126 = cf126;
      $dt.cf127 = cf127;
      $dt.cf128 = cf128;
      return $dt;
    }
    static create_DC69(cf121) {
      let $dt = new D26(2);
      $dt.cf121 = cf121;
      return $dt;
    }
    static create_DC72(cf129) {
      let $dt = new D26(3);
      $dt.cf129 = cf129;
      return $dt;
    }
    get is_DC70() { return this.$tag === 0; }
    get is_DC71() { return this.$tag === 1; }
    get is_DC69() { return this.$tag === 2; }
    get is_DC72() { return this.$tag === 3; }
    get dtor_cf122() { return this.cf122; }
    get dtor_cf123() { return this.cf123; }
    get dtor_cf124() { return this.cf124; }
    get dtor_cf125() { return this.cf125; }
    get dtor_cf126() { return this.cf126; }
    get dtor_cf127() { return this.cf127; }
    get dtor_cf128() { return this.cf128; }
    get dtor_cf121() { return this.cf121; }
    get dtor_cf129() { return this.cf129; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC70" + "(" + _dafny.toString(this.cf122) + ", " + _dafny.toString(this.cf123) + ", " + _dafny.toString(this.cf124) + ", " + _dafny.toString(this.cf125) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC71" + "(" + _dafny.toString(this.cf126) + ", " + _dafny.toString(this.cf127) + ", " + _dafny.toString(this.cf128) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC69" + "(" + _dafny.toString(this.cf121) + ")";
      } else if (this.$tag === 3) {
        return "D26.DC72" + "(" + _dafny.toString(this.cf129) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf122, other.cf122) && _dafny.areEqual(this.cf123, other.cf123) && this.cf124 === other.cf124 && this.cf125 === other.cf125;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf126, other.cf126) && this.cf127 === other.cf127 && _dafny.areEqual(this.cf128, other.cf128);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf121, other.cf121);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf129, other.cf129);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC70(_dafny.ZERO, _dafny.ZERO, false, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC73(cf130) {
      let $dt = new D27(0);
      $dt.cf130 = cf130;
      return $dt;
    }
    get is_DC73() { return this.$tag === 0; }
    get dtor_cf130() { return this.cf130; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC73" + "(" + _dafny.toString(this.cf130) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf130, other.cf130);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
        }
      };
    }
  }

  $module.D28 = class D28 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC75() {
      let $dt = new D28(0);
      return $dt;
    }
    static create_DC74(cf131) {
      let $dt = new D28(1);
      $dt.cf131 = cf131;
      return $dt;
    }
    get is_DC75() { return this.$tag === 0; }
    get is_DC74() { return this.$tag === 1; }
    get dtor_cf131() { return this.cf131; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC75";
      } else if (this.$tag === 1) {
        return "D28.DC74" + "(" + _dafny.toString(this.cf131) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf131, other.cf131);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D28.create_DC75();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D28.Default();
        }
      };
    }
  }

  $module.D29 = class D29 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC77(cf133) {
      let $dt = new D29(0);
      $dt.cf133 = cf133;
      return $dt;
    }
    static create_DC78(cf134, cf135, cf136) {
      let $dt = new D29(1);
      $dt.cf134 = cf134;
      $dt.cf135 = cf135;
      $dt.cf136 = cf136;
      return $dt;
    }
    static create_DC76(cf132) {
      let $dt = new D29(2);
      $dt.cf132 = cf132;
      return $dt;
    }
    get is_DC77() { return this.$tag === 0; }
    get is_DC78() { return this.$tag === 1; }
    get is_DC76() { return this.$tag === 2; }
    get dtor_cf133() { return this.cf133; }
    get dtor_cf134() { return this.cf134; }
    get dtor_cf135() { return this.cf135; }
    get dtor_cf136() { return this.cf136; }
    get dtor_cf132() { return this.cf132; }
    toString() {
      if (this.$tag === 0) {
        return "D29.DC77" + "(" + _dafny.toString(this.cf133) + ")";
      } else if (this.$tag === 1) {
        return "D29.DC78" + "(" + _dafny.toString(this.cf134) + ", " + _dafny.toString(this.cf135) + ", " + _dafny.toString(this.cf136) + ")";
      } else if (this.$tag === 2) {
        return "D29.DC76" + "(" + _dafny.toString(this.cf132) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf133, other.cf133);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf134, other.cf134) && _dafny.areEqual(this.cf135, other.cf135) && _dafny.areEqual(this.cf136, other.cf136);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf132, other.cf132);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D29.create_DC77(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D29.Default();
        }
      };
    }
  }

  $module.D30 = class D30 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC80(cf138, cf139, cf140, cf141) {
      let $dt = new D30(0);
      $dt.cf138 = cf138;
      $dt.cf139 = cf139;
      $dt.cf140 = cf140;
      $dt.cf141 = cf141;
      return $dt;
    }
    static create_DC81() {
      let $dt = new D30(1);
      return $dt;
    }
    static create_DC79(cf137) {
      let $dt = new D30(2);
      $dt.cf137 = cf137;
      return $dt;
    }
    static create_DC82(cf142) {
      let $dt = new D30(3);
      $dt.cf142 = cf142;
      return $dt;
    }
    get is_DC80() { return this.$tag === 0; }
    get is_DC81() { return this.$tag === 1; }
    get is_DC79() { return this.$tag === 2; }
    get is_DC82() { return this.$tag === 3; }
    get dtor_cf138() { return this.cf138; }
    get dtor_cf139() { return this.cf139; }
    get dtor_cf140() { return this.cf140; }
    get dtor_cf141() { return this.cf141; }
    get dtor_cf137() { return this.cf137; }
    get dtor_cf142() { return this.cf142; }
    toString() {
      if (this.$tag === 0) {
        return "D30.DC80" + "(" + _dafny.toString(this.cf138) + ", " + _dafny.toString(this.cf139) + ", " + _dafny.toString(this.cf140) + ", " + _dafny.toString(this.cf141) + ")";
      } else if (this.$tag === 1) {
        return "D30.DC81";
      } else if (this.$tag === 2) {
        return "D30.DC79" + "(" + _dafny.toString(this.cf137) + ")";
      } else if (this.$tag === 3) {
        return "D30.DC82" + "(" + _dafny.toString(this.cf142) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf138 === other.cf138 && _dafny.areEqual(this.cf139, other.cf139) && this.cf140 === other.cf140 && _dafny.areEqual(this.cf141, other.cf141);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf137, other.cf137);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf142, other.cf142);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D30.create_DC80(false, _dafny.Set.Empty, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D30.Default();
        }
      };
    }
  }

  $module.D31 = class D31 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC84(cf144, cf145) {
      let $dt = new D31(0);
      $dt.cf144 = cf144;
      $dt.cf145 = cf145;
      return $dt;
    }
    static create_DC83(cf143) {
      let $dt = new D31(1);
      $dt.cf143 = cf143;
      return $dt;
    }
    get is_DC84() { return this.$tag === 0; }
    get is_DC83() { return this.$tag === 1; }
    get dtor_cf144() { return this.cf144; }
    get dtor_cf145() { return this.cf145; }
    get dtor_cf143() { return this.cf143; }
    toString() {
      if (this.$tag === 0) {
        return "D31.DC84" + "(" + _dafny.toString(this.cf144) + ", " + _dafny.toString(this.cf145) + ")";
      } else if (this.$tag === 1) {
        return "D31.DC83" + "(" + _dafny.toString(this.cf143) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf144 === other.cf144 && this.cf145 === other.cf145;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf143, other.cf143);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D31.create_DC84(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D31.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f2 = false;
      this.f3 = _dafny.ZERO;
      this.f9 = _dafny.Seq.UnicodeFromString("");
      this.f10 = _dafny.Seq.of();
      this.f12 = _dafny.Seq.UnicodeFromString("");
      this.f13 = [];
      this.f15 = _dafny.ZERO;
      this.f16 = _dafny.MultiSet.Empty;
      this.f17 = _dafny.Map.Empty;
      this._f1 = false;
      this._f4 = _dafny.ZERO;
      this._f5 = false;
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this._f8 = _dafny.ZERO;
      this._f11 = false;
      this._f14 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this).f16 = f16;
      (_this).f17 = f17;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f31 = _dafny.ZERO;
      this._f32 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f31, f32) {
      let _this = this;
      (_this)._f31 = f31;
      (_this)._f32 = f32;
      return;
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f18 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f18) {
      let _this = this;
      (_this)._f18 = f18;
      return;
    }
    fm21(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat((((_this).f18) ? (_dafny.Seq.UnicodeFromString("fsarcis")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-683)), function (_442_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("k"), _dafny.Seq.UnicodeFromString("cjfcrawcl")));
    };
    fm22(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(918), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-951), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,new BigNumber(618))).length))));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _443_i0;
      _443_i0 = _dafny.ZERO;
      L2: {
        while (!((_this).f18) || (!((_this).f18))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_443_i0)) {
              break L2;
            }
            _443_i0 = (_443_i0).plus(_dafny.ONE);
            (globalState).f2 = (_module.__default.safeModuloInt(new BigNumber(916), p1)).isLessThanOrEqualTo(p1);
            (globalState).f2 = !(new BigNumber(164)).isEqualTo(p1);
            let _444_v0;
            _444_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
            let _445_v1;
            _445_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_444_v0).length),(_this).f18);
            let _446_v2;
            _446_v2 = _dafny.Seq.of((_445_v1).update(p1, (_this).f18));
            let _447_v3;
            _447_v3 = _module.D0.create_DC1((_this).f18, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(856)), function (_448_i1) {
  return new _dafny.CodePoint('a'.codePointAt(0));
}), (_this).f18, p1);
            let _449_v4;
            _449_v4 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_446_v2).length)).isLessThanOrEqualTo(p1),(p1).isLessThan((_this).fm22(p2, _447_v3, globalState)));
            _444_v0 = (_444_v0).update(true, (_this).f18);
            let _450_v5;
            let _init5 = function (_451_i2) {
              return _module.__default.safeDivisionInt(_451_i2, new BigNumber((_dafny.Set.fromElements((_this).f18, (_this).f18)).length));
            };
            let _nw61 = Array((new BigNumber(18)).toNumber());
            for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw61.length); _i0_5++) {
              _nw61[_i0_5] = _init5(new BigNumber(_i0_5));
            }
            _450_v5 = _nw61;
            let _452_v6;
            _452_v6 = _dafny.Seq.of((_this).f18);
            let _453_v7;
            _453_v7 = _dafny.Map.Empty.slice().updateUnsafe((_452_v6)[_module.__default.safeIndex(p1, new BigNumber((_452_v6).length))],_450_v5);
            let _454_v8;
            let _nw62 = Array((new BigNumber(9)).toNumber());
            _nw62[(_dafny.ZERO).toNumber()] = _450_v5;
            _nw62[(_dafny.ONE).toNumber()] = (((_453_v7).contains((_this).f18)) ? ((_453_v7).get((_this).f18)) : (_450_v5));
            _nw62[(new BigNumber(2)).toNumber()] = _450_v5;
            _nw62[(new BigNumber(3)).toNumber()] = _450_v5;
            _nw62[(new BigNumber(4)).toNumber()] = _450_v5;
            _nw62[(new BigNumber(5)).toNumber()] = _450_v5;
            _nw62[(new BigNumber(6)).toNumber()] = _450_v5;
            _nw62[(new BigNumber(7)).toNumber()] = _450_v5;
            _nw62[(new BigNumber(8)).toNumber()] = _450_v5;
            _454_v8 = _nw62;
            let _index44 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_454_v8).length));
            (_454_v8)[_index44] = _450_v5;
            let _455_v9;
            let _nw63 = Array((new BigNumber(11)).toNumber()).fill(false);
            _455_v9 = _nw63;
            let _index45 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_455_v9).length));
            (_455_v9)[_index45] = (_this).f18;
            let _456_v10;
            _456_v10 = _dafny.Set.fromElements(p1);
            let _index46 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_454_v8).length));
            let _index47 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_455_v9).length));
            let _rhs45 = p1;
            let _rhs46 = _450_v5;
            let _rhs47 = _module.__default.fm23(_456_v10, (_module.__default.fm1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(188)), function (_457_i3) {
              return new _dafny.CodePoint('j'.codePointAt(0));
            })).length), (_this).f18, (_this).f18, globalState)).multipliedBy(p1), globalState);
            let _lhs29 = globalState;
            let _lhs30 = _454_v8;
            let _lhs31 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_454_v8).length));
            let _lhs32 = _455_v9;
            let _lhs33 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_455_v9).length));
            _lhs29.f0 = _rhs45;
            _lhs30[_lhs31] = _rhs46;
            _lhs32[_lhs33] = _rhs47;
          }
        }
      }
      let _458_v11;
      _458_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p1);
      _458_v11 = (_458_v11).Merge(_458_v11);
      let _459_v12;
      _459_v12 = _module.D1.create_DC4(((_this).f18) === ((_this).f18));
      let _460_v13;
      _460_v13 = _dafny.Seq.of(p1);
      let _461_v14;
      _461_v14 = new _dafny.CodePoint('o'.codePointAt(0));
      let _462_v15;
      _462_v15 = _dafny.Set.fromElements(_461_v14);
      let _463_v16;
      _463_v16 = _module.D2.create_DC7((_dafny.ZERO).minus(new BigNumber((p2).length)), (_this).f18, p1);
      _459_v12 = _module.__default.fm24((_module.__default.fm25(p1, globalState)).contains(new BigNumber((_460_v13).length)), (_462_v15).Union(_462_v15), ((!((_463_v16).dtor_cf16)) ? ((_this).f18) : ((_this).f18)), globalState);
      let _464_v17;
      _464_v17 = _module.D2.create_DC8((_this).f18, (_this).f18, (_this).f18);
      let _465_v18;
      let _nw64 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
      _465_v18 = _nw64;
      let _466_v19;
      _466_v19 = _dafny.Seq.of((_this).f18, (_this).f18, !((_this).f18), (_this).f18, (_this).f18);
      let _index48 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_465_v18).length));
      (_465_v18)[_index48] = _dafny.Seq.update(_466_v19, _module.__default.safeIndex(p1, new BigNumber((_466_v19).length)), (_this).f18);
      let _467_v20;
      let _init6 = ((_468_v17) => function (_469_i4) {
        return (_468_v17).dtor_cf19;
      })(_464_v17);
      let _nw65 = Array((new BigNumber(8)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw65.length); _i0_6++) {
        _nw65[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _467_v20 = _nw65;
      let _index49 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
      (_467_v20)[_index49] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_466_v19)).contains((_this).f18);
      let _index50 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_465_v18).length));
      let _index51 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
      let _rhs48 = _module.__default.fm26(_460_v13, (p1).multipliedBy(p1), globalState);
      let _rhs49 = _466_v19;
      let _rhs50 = (_this).f18;
      let _rhs51 = (p1).isEqualTo((_module.__default.fm1(p1, (_this).f18, (_this).f18, globalState)).minus((_dafny.ZERO).minus(_module.__default.fm1(new BigNumber((p2).length), (_this).f18, (_this).f18, globalState))));
      let _lhs34 = _465_v18;
      let _lhs35 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_465_v18).length));
      let _lhs36 = _467_v20;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
      let _lhs38 = globalState;
      _464_v17 = _rhs48;
      _lhs34[_lhs35] = _rhs49;
      _lhs36[_lhs37] = _rhs50;
      _lhs38.f2 = _rhs51;
      let _index52 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
      (_467_v20)[_index52] = (_this).f18;
      if ((_463_v16).dtor_cf16) {
        let _470_v21;
        _470_v21 = _module.D0.create_DC1((_this).f18, p1, (((_467_v20)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length))]) ? (_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)))) : (_dafny.Seq.of(_461_v14, _461_v14, _461_v14, _461_v14))), (_this).f18, p1);
        let _pat_let_tv1 = _467_v20;
        let _pat_let_tv2 = _467_v20;
        let _pat_let_tv3 = p1;
        let _pat_let_tv4 = p1;
        let _index53 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
        let _rhs52 = _470_v21;
        let _rhs53 = function (_pat_let6_0) {
          return function (_471_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_472_dt__update_hcf16_h0) {
                return function (_pat_let8_0) {
                  return function (_473_dt__update_hcf15_h0) {
                    return _module.D2.create_DC7(_473_dt__update_hcf15_h0, _472_dt__update_hcf16_h0, (_471_dt__update__tmp_h0).dtor_cf17);
                  }(_pat_let8_0);
                }(_module.__default.safeDivisionInt(_pat_let_tv3, _pat_let_tv4));
              }(_pat_let7_0);
            }((_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_pat_let_tv1).length))]);
          }(_pat_let6_0);
        }(_module.D2.create_DC7(p1, (_467_v20)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length))], p1));
        let _rhs54 = true;
        let _lhs39 = _467_v20;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
        _470_v21 = _rhs52;
        _463_v16 = _rhs53;
        _lhs39[_lhs40] = _rhs54;
        let _474_v22;
        let _nw66 = new _module.C0();
        _nw66.__ctor(p1, p1);
        _474_v22 = _nw66;
        let _index54 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
        (_467_v20)[_index54] = ((p1).minus(p1)).isLessThanOrEqualTo(new BigNumber(50));
        let _475_v23;
        let _nw67 = new _module.C0();
        _nw67.__ctor((_474_v22).f31, (_474_v22).f31);
        _475_v23 = _nw67;
        (globalState).f0 = (_470_v21).dtor_cf5;
      } else {
        let _476_v24;
        _476_v24 = _dafny.MultiSet.fromElements(p1);
        (globalState).f2 = (_476_v24).IsDisjointFrom(_476_v24);
        let _477_v25;
        let _478_v26;
        let _479_v27;
        let _480_v28;
        let _out28;
        let _out29;
        let _out30;
        let _out31;
        let _outcollector7 = _module.__default.m0(globalState);
        _out28 = _outcollector7[0];
        _out29 = _outcollector7[1];
        _out30 = _outcollector7[2];
        _out31 = _outcollector7[3];
        _477_v25 = _out28;
        _478_v26 = _out29;
        _479_v27 = _out30;
        _480_v28 = _out31;
        let _481_v29;
        _481_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_462_v15);
        (globalState).f12 = (_this).fm21(_478_v26, p2, _481_v29, _466_v19, globalState);
        (globalState).f0 = _477_v25;
        let _index55 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_479_v27).length));
        (_479_v27)[_index55] = new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((p2).length), new BigNumber((p0).length)), _461_v14)).length);
        let _482_v30;
        _482_v30 = _dafny.MultiSet.fromElements((_467_v20)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length))]);
        let _483_v31;
        _483_v31 = _module.D1.create_DC3(_477_v25, _478_v26, _477_v25, (((_482_v30).contains((_this).f18)) ? ((_482_v30).get((_this).f18)) : (_478_v26)), _478_v26);
        let _484_v32;
        _484_v32 = _module.D1.create_DC5(_module.D1.create_DC5(_483_v31));
        let _485_v33;
        _485_v33 = _dafny.Map.Empty.slice().updateUnsafe((_467_v20)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length))],(_458_v11).update((_this).f18, _477_v25));
        let _486_v34;
        _486_v34 = _dafny.Seq.of(_458_v11, _458_v11);
        let _487_v35;
        _487_v35 = _module.D0.create_DC1(!((_467_v20)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length))]), new BigNumber(((((_485_v33).contains(true)) ? ((_485_v33).get(true)) : ((_486_v34)[_module.__default.safeIndex((((_482_v30).contains((_this).f18)) ? ((_482_v30).get((_this).f18)) : (_477_v25)), new BigNumber((_486_v34).length))]))).length), (_this).fm21(new BigNumber((p2).length), p2, _module.__default.fm27(globalState), _466_v19, globalState), false, new BigNumber((_460_v13).length));
        let _index56 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
        let _index57 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_479_v27).length));
        let _rhs55 = _478_v26;
        let _rhs56 = _477_v25;
        let _rhs57 = false;
        let _rhs58 = (_this).fm22(_dafny.Seq.Concat(p2, p0), _487_v35, globalState);
        let _rhs59 = _484_v32;
        let _lhs41 = globalState;
        let _lhs42 = _467_v20;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_467_v20).length));
        let _lhs44 = _479_v27;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_479_v27).length));
        _478_v26 = _rhs55;
        _lhs41.f0 = _rhs56;
        _lhs42[_lhs43] = _rhs57;
        _lhs44[_lhs45] = _rhs58;
        _484_v32 = _rhs59;
      }
      return;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f18 = false;
      this.f33 = _dafny.ZERO;
      this._f34 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f33, f34, f18) {
      let _this = this;
      (_this).f33 = f33;
      (_this)._f34 = f34;
      (_this)._f18 = f18;
      return;
    }
    fm19(p0, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm20(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("qa"), _dafny.Seq.UnicodeFromString("pj"), _dafny.Seq.UnicodeFromString("wx"))).Intersect((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(401)), function (_488_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("wpnvrbl"))).Intersect(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("lftn"), _dafny.Seq.UnicodeFromString("ogsxkngw"))));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _hi3 = _module.__default.fm1(_this.f33, (_this).f18, (_this).f18, globalState);
      for (let _489_i0 = p1; _489_i0.isLessThan(_hi3); _489_i0 = _489_i0.plus(_dafny.ONE)) {
        let _490_v0;
        _490_v0 = _dafny.MultiSet.fromElements((_this).f18);
        (globalState).f2 = !(((((_this).f18) ? (_490_v0) : (_490_v0))).contains((_this).f18));
        let _491_v1;
        _491_v1 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
        let _492_v2;
        _492_v2 = new _dafny.CodePoint('d'.codePointAt(0));
        _491_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,new BigNumber((_dafny.Seq.Concat(p2, _dafny.Seq.update(p2, _module.__default.safeIndex(_489_i0, new BigNumber((p2).length)), _492_v2))).length));
        let _493_v3;
        _493_v3 = _module.D1.create_DC4((_this).f18);
        let _source4 = function (_pat_let9_0) {
          return function (_494_dt__update__tmp_h0) {
            return function (_pat_let10_0) {
              return function (_495_dt__update_hcf12_h0) {
                return _module.D1.create_DC4(_495_dt__update_hcf12_h0);
              }(_pat_let10_0);
            }((_this).f18);
          }(_pat_let9_0);
        }(_493_v3);
        if (_source4.is_DC3) {
          let _496___mcc_h0 = (_source4).cf7;
          let _497___mcc_h1 = (_source4).cf8;
          let _498___mcc_h2 = (_source4).cf9;
          let _499___mcc_h3 = (_source4).cf10;
          let _500___mcc_h4 = (_source4).cf11;
          let _501_cf11 = _500___mcc_h4;
          let _502_cf10 = _499___mcc_h3;
          let _503_cf9 = _498___mcc_h2;
          let _504_cf8 = _497___mcc_h1;
          let _505_cf7 = _496___mcc_h0;
          let _506_v4;
          _506_v4 = _dafny.Seq.of((_this).f18, (_this).f18);
          let _507_v5;
          _507_v5 = _module.D2.create_DC7(new BigNumber((_490_v0).cardinality()), (_this).f18, new BigNumber(601));
          let _508_v6;
          let _init7 = ((_509_cf9) => function (_510_i1) {
            return (_510_i1).minus(_509_cf9);
          })(_503_cf9);
          let _nw68 = Array((new BigNumber(8)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw68.length); _i0_7++) {
            _nw68[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _508_v6 = _nw68;
          let _511_v7;
          _511_v7 = _dafny.Set.fromElements(_508_v6, _508_v6);
          let _512_v8;
          let _nw69 = Array((new BigNumber(10)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = (_this).f18;
          _nw69[(_dafny.ONE).toNumber()] = (_506_v4)[_module.__default.safeIndex(new BigNumber(-876), new BigNumber((_506_v4).length))];
          _nw69[(new BigNumber(2)).toNumber()] = (_this).f18;
          _nw69[(new BigNumber(3)).toNumber()] = (_this).fm19(_507_v5, globalState);
          _nw69[(new BigNumber(4)).toNumber()] = (new BigNumber((p2).length)).isLessThan(_503_cf9);
          _nw69[(new BigNumber(5)).toNumber()] = (_this).f18;
          _nw69[(new BigNumber(6)).toNumber()] = (_this).f18;
          _nw69[(new BigNumber(7)).toNumber()] = !((_511_v7).IsDisjointFrom(_511_v7));
          _nw69[(new BigNumber(8)).toNumber()] = !((_this).f18);
          _nw69[(new BigNumber(9)).toNumber()] = (_this).f18;
          _512_v8 = _nw69;
          let _index58 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_512_v8).length));
          (_512_v8)[_index58] = (_this).f18;
          let _513_v9;
          _513_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_491_v1);
          let _514_v10;
          _514_v10 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(((_490_v0).contains((_this).f18)) ? ((_490_v0).get((_this).f18)) : (_504_cf8))), _491_v1, _491_v1, _491_v1);
          let _515_v11;
          _515_v11 = _module.D0.create_DC1((_this).f18, _502_cf10, p0, !((_this).f18), p1);
          let _516_v12;
          _516_v12 = _dafny.MultiSet.fromElements(p0, (_515_v11).dtor_cf3);
          let _index59 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_512_v8).length));
          let _rhs60 = _dafny.Seq.IsPrefixOf(p0, _dafny.Seq.Concat(p2, p0));
          let _rhs61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_514_v10)[_module.__default.safeIndex(new BigNumber((_516_v12).cardinality()), new BigNumber((_514_v10).length))]);
          let _lhs46 = _512_v8;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_512_v8).length));
          _lhs46[_lhs47] = _rhs60;
          _513_v9 = _rhs61;
          _505_cf7 = _504_cf8;
          let _517_v13;
          _517_v13 = _dafny.Set.fromElements((((_this).f18) ? ((p0)[_module.__default.safeIndex(new BigNumber(732), new BigNumber((p0).length))]) : (_492_v2)), _492_v2, _492_v2, _492_v2, _492_v2);
          _517_v13 = (_dafny.Set.fromElements(_492_v2)).Union((_517_v13).Intersect(_517_v13));
          let _index60 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_508_v6).length));
          (_508_v6)[_index60] = p1;
          let _index61 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_508_v6).length));
          (_508_v6)[_index61] = _module.__default.fm1(new BigNumber((_490_v0).cardinality()), (_512_v8)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_512_v8).length))], (_this).f18, globalState);
        } else if (_source4.is_DC4) {
          let _518___mcc_h5 = (_source4).cf12;
          let _519_cf12 = _518___mcc_h5;
          let _520_v14;
          _520_v14 = _dafny.Seq.of((_this).f18, _519_cf12, (_this).f18, (_this).f18);
          _520_v14 = _520_v14;
          (globalState).f3 = new BigNumber(272);
          let _521_v15;
          let _nw70 = Array((new BigNumber(6)).toNumber()).fill(false);
          _521_v15 = _nw70;
          let _index62 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_521_v15).length));
          (_521_v15)[_index62] = !(_519_cf12);
          let _522_v16;
          _522_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_519_cf12);
          let _index63 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_521_v15).length));
          (_521_v15)[_index63] = (((_522_v16).contains(true)) ? ((_522_v16).get(true)) : (_519_cf12));
          let _523_v17;
          let _nw71 = new _module.C1();
          _nw71.__ctor((_521_v15)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_521_v15).length))]);
          _523_v17 = _nw71;
          let _nw72 = new _module.C1();
          _nw72.__ctor((_523_v17).f18);
          _523_v17 = _nw72;
        } else if (_source4.is_DC2) {
          let _524___mcc_h6 = (_source4).cf6;
          let _525_cf6 = _524___mcc_h6;
          let _526_v19;
          _526_v19 = _dafny.Map.Empty.slice().updateUnsafe(_489_i0,new BigNumber(-697));
          let _527_v20;
          _527_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_526_v19).length),new BigNumber((p2).length));
          let _528_v21;
          _528_v21 = _dafny.Seq.of(_527_v20);
          let _529_v22;
          _529_v22 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(779)), ((_530_p1) => function (_531_i2) {
            return _530_p1;
          })(p1)));
          (globalState).f3 = new BigNumber((function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of ((_527_v20).Merge((_528_v21)[_module.__default.safeIndex((((_529_v22).contains(_dafny.Seq.of(_489_i0, _489_i0))) ? ((_529_v22).get(_dafny.Seq.of(_489_i0, _489_i0))) : (_489_i0)), new BigNumber((_528_v21).length))])).Keys.Elements) {
              let _532_v18 = _compr_23;
              if (((_527_v20).Merge((_528_v21)[_module.__default.safeIndex((((_529_v22).contains(_dafny.Seq.of(_489_i0, _489_i0))) ? ((_529_v22).get(_dafny.Seq.of(_489_i0, _489_i0))) : (_489_i0)), new BigNumber((_528_v21).length))])).contains(_532_v18)) {
                _coll23.push([_module.__default.safeDivisionInt(_532_v18, new BigNumber(-808)),(_this).f18]);
              }
            }
            return _coll23;
          }()).length);
          (globalState).f12 = p0;
          let _533_v23;
          let _nw73 = Array((new BigNumber(18)).toNumber()).fill([]);
          _533_v23 = _nw73;
          let _534_v24;
          let _nw74 = Array((new BigNumber(18)).toNumber()).fill(false);
          _534_v24 = _nw74;
          let _index64 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_533_v23).length));
          (_533_v23)[_index64] = _534_v24;
          let _535_v25;
          let _nw75 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
          _535_v25 = _nw75;
          let _536_v26;
          _536_v26 = _dafny.Set.fromElements((_this).f18);
          let _537_v27;
          _537_v27 = _dafny.MultiSet.fromElements(new BigNumber(((_this).f34).length));
          let _index65 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_533_v23).length));
          let _rhs62 = _module.__default.safeModuloInt(new BigNumber((_536_v26).length), ((((_491_v1).contains((_this).f18)) ? ((_491_v1).get((_this).f18)) : (p1))).minus(new BigNumber((_537_v27).cardinality())));
          let _rhs63 = _534_v24;
          let _rhs64 = _535_v25;
          let _lhs48 = globalState;
          let _lhs49 = _533_v23;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_533_v23).length));
          _lhs48.f0 = _rhs62;
          _lhs49[_lhs50] = _rhs63;
          _535_v25 = _rhs64;
          let _538_v28;
          _538_v28 = _dafny.Map.Empty.slice().updateUnsafe(_489_i0,(_dafny.MultiSet.fromElements(p1, new BigNumber(731))).Difference(_dafny.MultiSet.fromElements(_489_i0)));
          _538_v28 = (_538_v28).update((_dafny.ZERO).minus(_489_i0), _537_v27);
        } else {
          let _539___mcc_h7 = (_source4).cf13;
          let _540_cf13 = _539___mcc_h7;
          let _541_v29;
          _541_v29 = _module.D1.create_DC3(p1, p1, _this.f33, new BigNumber(87), _489_i0);
          let _542_v30;
          _542_v30 = _dafny.Map.Empty.slice().updateUnsafe((((_490_v0).contains((_this).f18)) ? ((_490_v0).get((_this).f18)) : (_489_i0)),new BigNumber(841));
          let _543_v31;
          let _nw76 = Array((new BigNumber(3)).toNumber());
          _nw76[(_dafny.ZERO).toNumber()] = _this.f33;
          _nw76[(_dafny.ONE).toNumber()] = new BigNumber(783);
          _nw76[(new BigNumber(2)).toNumber()] = p1;
          _543_v31 = _nw76;
          let _544_v32;
          _544_v32 = _dafny.Set.fromElements(_543_v31);
          let _pat_let_tv5 = p1;
          let _pat_let_tv6 = _544_v32;
          let _545_v33;
          let _nw77 = Array((new BigNumber(26)).toNumber());
          _nw77[(_dafny.ZERO).toNumber()] = _541_v29;
          _nw77[(_dafny.ONE).toNumber()] = _541_v29;
          _nw77[(new BigNumber(2)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(3)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(4)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(5)).toNumber()] = function (_pat_let11_0) {
            return function (_546_dt__update__tmp_h1) {
              return function (_pat_let12_0) {
                return function (_547_dt__update_hcf9_h0) {
                  return _module.D1.create_DC3((_546_dt__update__tmp_h1).dtor_cf7, (_546_dt__update__tmp_h1).dtor_cf8, _547_dt__update_hcf9_h0, (_546_dt__update__tmp_h1).dtor_cf10, (_546_dt__update__tmp_h1).dtor_cf11);
                }(_pat_let12_0);
              }(new BigNumber(939));
            }(_pat_let11_0);
          }(_541_v29);
          _nw77[(new BigNumber(6)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(7)).toNumber()] = function (_pat_let13_0) {
            return function (_548_dt__update__tmp_h2) {
              return function (_pat_let14_0) {
                return function (_549_dt__update_hcf10_h0) {
                  return function (_pat_let15_0) {
                    return function (_550_dt__update_hcf7_h0) {
                      return function (_pat_let16_0) {
                        return function (_551_dt__update_hcf8_h0) {
                          return _module.D1.create_DC3(_550_dt__update_hcf7_h0, _551_dt__update_hcf8_h0, (_548_dt__update__tmp_h2).dtor_cf9, _549_dt__update_hcf10_h0, (_548_dt__update__tmp_h2).dtor_cf11);
                        }(_pat_let16_0);
                      }(_this.f33);
                    }(_pat_let15_0);
                  }(_pat_let_tv5);
                }(_pat_let14_0);
              }(_this.f33);
            }(_pat_let13_0);
          }(_module.D1.create_DC3(_489_i0, p1, _this.f33, _this.f33, new BigNumber(((_this).f34).length)));
          _nw77[(new BigNumber(8)).toNumber()] = _module.D1.create_DC3((_dafny.ZERO).minus(p1), new BigNumber((_542_v30).length), p1, _this.f33, _this.f33);
          _nw77[(new BigNumber(9)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(10)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(11)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(12)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(13)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(14)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(15)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(16)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(17)).toNumber()] = _module.D1.create_DC3((_dafny.ZERO).minus(_this.f33), new BigNumber((_dafny.Seq.of((_this).f18)).length), p1, new BigNumber(484), new BigNumber(572));
          _nw77[(new BigNumber(18)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(19)).toNumber()] = function (_pat_let17_0) {
            return function (_552_dt__update__tmp_h3) {
              return function (_pat_let18_0) {
                return function (_553_dt__update_hcf10_h1) {
                  return _module.D1.create_DC3((_552_dt__update__tmp_h3).dtor_cf7, (_552_dt__update__tmp_h3).dtor_cf8, (_552_dt__update__tmp_h3).dtor_cf9, _553_dt__update_hcf10_h1, (_552_dt__update__tmp_h3).dtor_cf11);
                }(_pat_let18_0);
              }(new BigNumber((_pat_let_tv6).length));
            }(_pat_let17_0);
          }(_541_v29);
          _nw77[(new BigNumber(20)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(21)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(22)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(23)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(24)).toNumber()] = _541_v29;
          _nw77[(new BigNumber(25)).toNumber()] = _module.D1.create_DC3(new BigNumber(727), _this.f33, new BigNumber(-359), _module.__default.fm1(_this.f33, (_this).f18, (_this).f18, globalState), _this.f33);
          _545_v33 = _nw77;
          let _index66 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_545_v33).length));
          (_545_v33)[_index66] = _541_v29;
          let _554_v34;
          _554_v34 = _module.D0.create_DC0((_this).f18);
          let _index67 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_545_v33).length));
          (_545_v33)[_index67] = _module.__default.fm28(function (_pat_let19_0) {
            return function (_555_dt__update__tmp_h4) {
              return function (_pat_let20_0) {
                return function (_556_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_556_dt__update_hcf0_h0);
                }(_pat_let20_0);
              }((_this).f18);
            }(_pat_let19_0);
          }(_554_v34), globalState);
          (globalState).f0 = (_dafny.ZERO).minus(_this.f33);
          (globalState).f3 = p1;
          let _557_v35;
          let _nw78 = Array((new BigNumber(25)).toNumber()).fill(false);
          _557_v35 = _nw78;
          let _558_v36;
          _558_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_557_v35);
          _558_v36 = (_558_v36).update((_this).f18, _557_v35);
        }
        let _559_v37;
        let _init8 = function (_560_i3) {
          return (_this).f18;
        };
        let _nw79 = Array((new BigNumber(19)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw79.length); _i0_8++) {
          _nw79[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _559_v37 = _nw79;
        let _561_v38;
        _561_v38 = _dafny.Set.fromElements((_this).f18);
        let _562_v39;
        _562_v39 = _dafny.Seq.of(_561_v38);
        let _563_v40;
        _563_v40 = _dafny.MultiSet.fromElements(_492_v2);
        let _564_v41;
        let _nw80 = Array((new BigNumber(10)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = p1;
        _nw80[(_dafny.ONE).toNumber()] = _489_i0;
        _nw80[(new BigNumber(2)).toNumber()] = _489_i0;
        _nw80[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_489_i0), _module.__default.fm1(new BigNumber((p2).length), (_this).f18, (_this).f18, globalState));
        _nw80[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_489_i0);
        _nw80[(new BigNumber(5)).toNumber()] = (_489_i0).multipliedBy(new BigNumber((_562_v39).length));
        _nw80[(new BigNumber(6)).toNumber()] = _this.f33;
        _nw80[(new BigNumber(7)).toNumber()] = _this.f33;
        _nw80[(new BigNumber(8)).toNumber()] = _489_i0;
        _nw80[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(_489_i0, new BigNumber((_563_v40).cardinality()));
        _564_v41 = _nw80;
        let _index68 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_564_v41).length));
        (_564_v41)[_index68] = new BigNumber(79);
        let _565_v42;
        _565_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f34).length),(_this).f18);
        let _566_v43;
        _566_v43 = _dafny.Seq.of(_565_v42);
        let _567_v44;
        _567_v44 = _dafny.Map.Empty.slice().updateUnsafe(_566_v43,p1);
        let _568_v45;
        _568_v45 = _module.D2.create_DC7(p1, (_this).f18, _489_i0);
        let _index69 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_564_v41).length));
        let _rhs65 = _559_v37;
        let _rhs66 = (_489_i0).multipliedBy(_this.f33);
        let _rhs67 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update((_this).f34, _module.__default.safeIndex(_489_i0, new BigNumber(((_this).f34).length)), new BigNumber((_490_v0).cardinality())), (_this).f34)).length)).minus((((_567_v44).contains(_566_v43)) ? ((_567_v44).get(_566_v43)) : (_489_i0)));
        let _rhs68 = !((_this).fm19(_568_v45, globalState)) || ((_this).f18);
        let _lhs51 = _564_v41;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_564_v41).length));
        let _lhs53 = globalState;
        let _lhs54 = globalState;
        _559_v37 = _rhs65;
        _lhs51[_lhs52] = _rhs66;
        _lhs53.f0 = _rhs67;
        _lhs54.f2 = _rhs68;
      }
      let _569_v46;
      _569_v46 = _dafny.MultiSet.fromElements(!((_this).f18), (_this).f18);
      let _570_v47;
      _570_v47 = _module.D2.create_DC7(p1, (_module.D0.create_DC1((_this).f18, (_dafny.ZERO).minus(_this.f33), p2, (_this).f18, _this.f33)).dtor_cf1, (_dafny.ZERO).minus(_this.f33));
      let _571_v48;
      _571_v48 = _dafny.Map.Empty.slice().updateUnsafe(_this.f33,(_570_v47).dtor_cf15);
      let _572_v49;
      _572_v49 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_this).f18);
      let _573_v50;
      _573_v50 = _dafny.Seq.of(true, true, (_this).f18, (_this).f18, false);
      let _574_v51;
      _574_v51 = _dafny.Map.Empty.slice().updateUnsafe(_573_v50,true);
      let _575_v52;
      _575_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_574_v51).length),(_this).f18);
      let _576_v53;
      _576_v53 = _dafny.MultiSet.fromElements(p1);
      let _577_v54;
      _577_v54 = _dafny.Map.Empty.slice().updateUnsafe(_576_v53,(_this).f18);
      let _578_v55;
      _578_v55 = _dafny.Seq.of((_this).f18, (_573_v50)[_module.__default.safeIndex(p1, new BigNumber((_573_v50).length))], (((_575_v52).contains((_dafny.ZERO).minus(new BigNumber((_577_v54).length)))) ? ((_575_v52).get((_dafny.ZERO).minus(new BigNumber((_577_v54).length)))) : (!(true))));
      let _579_v56;
      _579_v56 = _module.D1.create_DC3((_dafny.ZERO).minus(p1), _this.f33, p1, _this.f33, (_dafny.ZERO).minus(new BigNumber((_573_v50).length)));
      let _580_v57;
      _580_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,new BigNumber((_576_v53).cardinality()));
      let _581_v58;
      _581_v58 = _dafny.Set.fromElements(new BigNumber((_569_v46).cardinality()));
      let _582_v59;
      let _nw81 = Array((new BigNumber(23)).toNumber());
      _nw81[(_dafny.ZERO).toNumber()] = _this.f33;
      _nw81[(_dafny.ONE).toNumber()] = (new BigNumber((_module.__default.fm29((_this).f18, _this.f33, _dafny.Seq.UnicodeFromString("deqxsk"), globalState)).length)).minus(_this.f33);
      _nw81[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((((_569_v46).contains((_this).f18)) ? ((_569_v46).get((_this).f18)) : (new BigNumber((_571_v48).length))));
      _nw81[(new BigNumber(3)).toNumber()] = _module.__default.fm1(_this.f33, (_this).f18, (_570_v47).dtor_cf16, globalState);
      _nw81[(new BigNumber(4)).toNumber()] = new BigNumber((_572_v49).length);
      _nw81[(new BigNumber(5)).toNumber()] = new BigNumber((_573_v50).length);
      _nw81[(new BigNumber(6)).toNumber()] = p1;
      _nw81[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_this.f33, (_this).f18, (_this).f18, globalState);
      _nw81[(new BigNumber(8)).toNumber()] = (function (_pat_let21_0) {
        return function (_583_dt__update__tmp_h5) {
          return function (_pat_let22_0) {
            return function (_584_dt__update_hcf10_h2) {
              return function (_pat_let23_0) {
                return function (_585_dt__update_hcf11_h0) {
                  return _module.D1.create_DC3((_583_dt__update__tmp_h5).dtor_cf7, (_583_dt__update__tmp_h5).dtor_cf8, (_583_dt__update__tmp_h5).dtor_cf9, _584_dt__update_hcf10_h2, _585_dt__update_hcf11_h0);
                }(_pat_let23_0);
              }(new BigNumber(-280));
            }(_pat_let22_0);
          }(_this.f33);
        }(_pat_let21_0);
      }(_579_v56)).dtor_cf10;
      _nw81[(new BigNumber(9)).toNumber()] = _this.f33;
      _nw81[(new BigNumber(10)).toNumber()] = _this.f33;
      _nw81[(new BigNumber(11)).toNumber()] = p1;
      _nw81[(new BigNumber(12)).toNumber()] = _this.f33;
      _nw81[(new BigNumber(13)).toNumber()] = ((_this).f34)[_module.__default.safeIndex(new BigNumber((_580_v57).length), new BigNumber(((_this).f34).length))];
      _nw81[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(593)), function (_586_i4) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }))).length);
      _nw81[(new BigNumber(15)).toNumber()] = _this.f33;
      _nw81[(new BigNumber(16)).toNumber()] = (_this.f33).plus(_this.f33);
      _nw81[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_581_v58).length)).minus((_dafny.ZERO).minus((((_569_v46).contains((_this).f18)) ? ((_569_v46).get((_this).f18)) : (new BigNumber((p2).length))))));
      _nw81[(new BigNumber(18)).toNumber()] = _this.f33;
      _nw81[(new BigNumber(19)).toNumber()] = p1;
      _nw81[(new BigNumber(20)).toNumber()] = p1;
      _nw81[(new BigNumber(21)).toNumber()] = new BigNumber(((_this).f34).length);
      _nw81[(new BigNumber(22)).toNumber()] = (((_this).f18) ? (_this.f33) : (p1));
      _582_v59 = _nw81;
      _582_v59 = _582_v59;
      let _587_v60;
      let _nw82 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
      _587_v60 = _nw82;
      let _588_v61;
      _588_v61 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pb"), p0), p2);
      let _589_v62;
      _589_v62 = _dafny.Set.fromElements(!((_this).f18));
      let _590_v63;
      let _nw83 = Array((new BigNumber(8)).toNumber());
      _nw83[(_dafny.ZERO).toNumber()] = (_this.f33).isLessThan(new BigNumber((_589_v62).length));
      _nw83[(_dafny.ONE).toNumber()] = (_this).f18;
      _nw83[(new BigNumber(2)).toNumber()] = (_this).f18;
      _nw83[(new BigNumber(3)).toNumber()] = (_576_v53).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_580_v57).length))));
      _nw83[(new BigNumber(4)).toNumber()] = (_this).f18;
      _nw83[(new BigNumber(5)).toNumber()] = (_this).f18;
      _nw83[(new BigNumber(6)).toNumber()] = (_this).f18;
      _nw83[(new BigNumber(7)).toNumber()] = (_this).f18;
      _590_v63 = _nw83;
      let _index70 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length));
      (_590_v63)[_index70] = (((_575_v52).contains(_this.f33)) ? ((_575_v52).get(_this.f33)) : (true));
      let _591_v64;
      _591_v64 = _module.D2.create_DC8((_this).f18, false, (_this).f18);
      let _592_v65;
      _592_v65 = _dafny.Map.Empty.slice().updateUnsafe(_591_v64,_590_v63);
      let _593_v66;
      _593_v66 = new _dafny.CodePoint('n'.codePointAt(0));
      let _594_v67;
      _594_v67 = _dafny.Map.Empty.slice().updateUnsafe(_593_v66,new BigNumber(790));
      let _index71 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length));
      let _rhs69 = (_module.D3.create_DC10(_587_v60)).dtor_cf22;
      let _rhs70 = _dafny.Seq.Concat(_588_v61, _588_v61);
      let _rhs71 = _dafny.Seq.Concat((_this).f34, _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex((((_594_v67).contains(_593_v66)) ? ((_594_v67).get(_593_v66)) : (p1)), new BigNumber((p0).length)), _593_v66)).length)), (_this).f34));
      let _rhs72 = !((_this).fm19(_570_v47, globalState)) || ((_581_v58).IsDisjointFrom(function () {
        let _coll24 = new _dafny.Set();
        for (const _compr_24 of (_571_v48).Keys.Elements) {
          let _595_v68 = _compr_24;
          if ((_571_v48).contains(_595_v68)) {
            _coll24.add(_module.__default.safeModuloInt(_595_v68, new BigNumber((_dafny.Seq.UnicodeFromString("u")).length)));
          }
        }
        return _coll24;
      }()));
      let _rhs73 = _592_v65;
      let _lhs55 = globalState;
      let _lhs56 = _590_v63;
      let _lhs57 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length));
      _587_v60 = _rhs69;
      _588_v61 = _rhs70;
      _lhs55.f10 = _rhs71;
      _lhs56[_lhs57] = _rhs72;
      _592_v65 = _rhs73;
      let _596_i5;
      _596_i5 = _dafny.ZERO;
      L3: {
        while (true) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_596_i5)) {
              break L3;
            }
            _596_i5 = (_596_i5).plus(_dafny.ONE);
            let _rhs74 = new BigNumber((p0).length);
            let _rhs75 = (p1).plus(_this.f33);
            let _lhs58 = globalState;
            let _lhs59 = globalState;
            _lhs58.f15 = _rhs74;
            _lhs59.f3 = _rhs75;
            let _index72 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length));
            (_590_v63)[_index72] = (_this.f33).isLessThan(_this.f33);
            if (((new BigNumber((p2).length)).multipliedBy(new BigNumber(-440))).isLessThan(_this.f33)) {
              let _rhs76 = _593_v66;
              let _rhs77 = _589_v62;
              let _rhs78 = new BigNumber(815);
              let _lhs60 = globalState;
              _593_v66 = _rhs76;
              _589_v62 = _rhs77;
              _lhs60.f15 = _rhs78;
              let _597_v69;
              _597_v69 = _dafny.Set.fromElements(_569_v46, (_569_v46).update((_this).f18, _module.__default.abs(p1)), _569_v46);
              let _598_v70;
              let _nw84 = new _module.C0();
              _nw84.__ctor(_this.f33, new BigNumber((_597_v69).length));
              _598_v70 = _nw84;
              let _599_v71;
              let _nw85 = new _module.C0();
              _nw85.__ctor(new BigNumber((_573_v50).length), _this.f33);
              _599_v71 = _nw85;
              let _600_v72;
              _600_v72 = _dafny.Map.Empty.slice().updateUnsafe(_590_v63,(_590_v63)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length))]);
              _600_v72 = (_600_v72).update(_590_v63, (_this).f18);
              let _601_v73;
              _601_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p2);
              let _602_v74;
              _602_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(((_601_v73).contains(false)) ? ((_601_v73).get(false)) : (p2)));
              _601_v73 = (_601_v73).update((_this).f18, p0);
            } else {
              let _603_v75;
              let _nw86 = new _module.C0();
              _nw86.__ctor(_this.f33, (_dafny.ZERO).minus(new BigNumber((p0).length)));
              _603_v75 = _nw86;
              let _604_v76;
              _604_v76 = _dafny.Map.Empty.slice().updateUnsafe((_590_v63)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length))],_580_v57);
              let _605_v77;
              _605_v77 = _dafny.Map.Empty.slice().updateUnsafe((_604_v76).Merge(_604_v76),(_603_v75).f32);
              _605_v77 = (_605_v77).update(_604_v76, new BigNumber(870));
              let _606_v78;
              _606_v78 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_603_v75).f32);
              (globalState).f3 = (((_576_v53).contains((((_this).f18) ? (new BigNumber((_606_v78).length)) : ((_603_v75).f31)))) ? ((_576_v53).get((((_this).f18) ? (new BigNumber((_606_v78).length)) : ((_603_v75).f31)))) : ((_603_v75).f32));
              let _607_v79;
              _607_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,true);
              _607_v79 = ((_607_v79).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,(_590_v63)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length))]))).Merge((_607_v79).update(true, true));
              let _608_v80;
              let _nw87 = Array((new BigNumber(11)).toNumber());
              _nw87[(_dafny.ZERO).toNumber()] = _579_v56;
              _nw87[(_dafny.ONE).toNumber()] = _579_v56;
              _nw87[(new BigNumber(2)).toNumber()] = function (_pat_let24_0) {
                return function (_609_dt__update__tmp_h6) {
                  return function (_pat_let25_0) {
                    return function (_610_dt__update_hcf10_h3) {
                      return _module.D1.create_DC3((_609_dt__update__tmp_h6).dtor_cf7, (_609_dt__update__tmp_h6).dtor_cf8, (_609_dt__update__tmp_h6).dtor_cf9, _610_dt__update_hcf10_h3, (_609_dt__update__tmp_h6).dtor_cf11);
                    }(_pat_let25_0);
                  }(_dafny.ONE);
                }(_pat_let24_0);
              }(_579_v56);
              _nw87[(new BigNumber(3)).toNumber()] = _579_v56;
              _nw87[(new BigNumber(4)).toNumber()] = _579_v56;
              _nw87[(new BigNumber(5)).toNumber()] = _579_v56;
              _nw87[(new BigNumber(6)).toNumber()] = _module.D1.create_DC3(new BigNumber((_575_v52).length), (_603_v75).f31, p1, (_dafny.ZERO).minus(new BigNumber((_571_v48).length)), (_603_v75).f32);
              _nw87[(new BigNumber(7)).toNumber()] = _579_v56;
              _nw87[(new BigNumber(8)).toNumber()] = _579_v56;
              _nw87[(new BigNumber(9)).toNumber()] = _579_v56;
              _nw87[(new BigNumber(10)).toNumber()] = _579_v56;
              _608_v80 = _nw87;
              let _611_v81;
              _611_v81 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,p1);
              let _index73 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_608_v80).length));
              (_608_v80)[_index73] = (((_590_v63)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length))]) ? (_579_v56) : (_module.D1.create_DC3((_603_v75).f32, new BigNumber((_611_v81).length), _this.f33, p1, p1)));
              let _index74 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_608_v80).length));
              (_608_v80)[_index74] = (((_this).f18) ? (_579_v56) : (_579_v56));
            }
            (globalState).f2 = (_this).f18;
          }
        }
      }
      let _hi4 = _this.f33;
      for (let _612_i6 = _module.__default.fm1(_this.f33, (_this).f18, (_590_v63)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length))], globalState); _612_i6.isLessThan(_hi4); _612_i6 = _612_i6.plus(_dafny.ONE)) {
        (globalState).f3 = _this.f33;
        let _613_v82;
        _613_v82 = _dafny.Map.Empty.slice().updateUnsafe(true,_590_v63);
        let _614_v83;
        _614_v83 = _dafny.Seq.of(_613_v82, _613_v82, _613_v82);
        let _615_v84;
        _615_v84 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(883),(_613_v82).Merge((_614_v83)[_module.__default.safeIndex(p1, new BigNumber((_614_v83).length))]));
        _615_v84 = (_615_v84).update(_612_i6, _dafny.Map.Empty.slice().updateUnsafe(!((_590_v63)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length))]),_590_v63));
        let _616_v85;
        let _nw88 = new _module.C0();
        _nw88.__ctor((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-68)), ((_617_v66) => function (_618_i7) {
          return _617_v66;
        })(_593_v66))).length)).multipliedBy(p1), _this.f33);
        _616_v85 = _nw88;
        let _index75 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length));
        (_590_v63)[_index75] = (_590_v63)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_590_v63).length))];
      }
      (globalState).f2 = (_this).f18;
      return;
    }
    get f34() {
      let _this = this;
      return _this._f34;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f18 = false;
      this._f19 = _dafny.Seq.UnicodeFromString("");
      this._f20 = _dafny.Seq.of();
      this._f35 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f35, f19, f20, f18) {
      let _this = this;
      (_this)._f35 = f35;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f18 = f18;
      return;
    }
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return !((_this).f18) || (!_dafny.Seq.contains(_dafny.Seq.of((_this).f18, true), (_module.D3.create_DC11((_this).f18, !((_this).f18))).dtor_cf23));
    };
    fm3(p0, globalState) {
      let _this = this;
      if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))).IsSubsetOf(_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0))))) {
        return (_this).f18;
      } else {
        return (_this).f18;
      }
    };
    fm37(p0, globalState) {
      let _this = this;
      return (_this).f19;
    };
    fm38(p0, globalState) {
      let _this = this;
      return new BigNumber(814);
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      (globalState).f2 = p2;
      let _619_v0;
      _619_v0 = _dafny.MultiSet.fromElements(p1);
      if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(p1, p1, p2, p1))).equals(_619_v0)) {
        (globalState).f2 = false;
        let _620_v1;
        let _nw89 = Array((new BigNumber(15)).toNumber()).fill([]);
        _620_v1 = _nw89;
        let _621_v2;
        _621_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _622_v3;
        let _nw90 = Array((new BigNumber(16)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = (((_621_v2).contains(p2)) ? ((_621_v2).get(p2)) : ((_this).f18));
        _nw90[(_dafny.ONE).toNumber()] = p1;
        _nw90[(new BigNumber(2)).toNumber()] = p1;
        _nw90[(new BigNumber(3)).toNumber()] = p2;
        _nw90[(new BigNumber(4)).toNumber()] = p2;
        _nw90[(new BigNumber(5)).toNumber()] = p1;
        _nw90[(new BigNumber(6)).toNumber()] = (_this).f18;
        _nw90[(new BigNumber(7)).toNumber()] = p2;
        _nw90[(new BigNumber(8)).toNumber()] = p2;
        _nw90[(new BigNumber(9)).toNumber()] = (_this).f18;
        _nw90[(new BigNumber(10)).toNumber()] = (((_621_v2).contains(p1)) ? ((_621_v2).get(p1)) : (p2));
        _nw90[(new BigNumber(11)).toNumber()] = p2;
        _nw90[(new BigNumber(12)).toNumber()] = p1;
        _nw90[(new BigNumber(13)).toNumber()] = p2;
        _nw90[(new BigNumber(14)).toNumber()] = (_this).f18;
        _nw90[(new BigNumber(15)).toNumber()] = p1;
        _622_v3 = _nw90;
        let _index76 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_620_v1).length));
        (_620_v1)[_index76] = ((false) ? (_622_v3) : (_622_v3));
        let _index77 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_620_v1).length));
        (_620_v1)[_index77] = _622_v3;
        (globalState).f2 = (_this).f18;
        let _623_v4;
        _623_v4 = new BigNumber(640);
        (globalState).f0 = (_623_v4).minus(new BigNumber(((_this).fm37(new BigNumber(((_this).f19).length), globalState)).length));
        let _624_v5;
        _624_v5 = _dafny.Seq.of(p1);
        let _625_v6;
        _625_v6 = _dafny.Map.Empty.slice().updateUnsafe(_623_v4,(_this).f18);
        let _626_v7;
        let _nw91 = Array((new BigNumber(19)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = _619_v0;
        _nw91[(_dafny.ONE).toNumber()] = _619_v0;
        _nw91[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements((_this).f18)).Difference(_619_v0);
        _nw91[(new BigNumber(3)).toNumber()] = (_619_v0).update(p1, _module.__default.abs(_623_v4));
        _nw91[(new BigNumber(4)).toNumber()] = _619_v0;
        _nw91[(new BigNumber(5)).toNumber()] = _619_v0;
        _nw91[(new BigNumber(6)).toNumber()] = _619_v0;
        _nw91[(new BigNumber(7)).toNumber()] = _619_v0;
        _nw91[(new BigNumber(8)).toNumber()] = _619_v0;
        _nw91[(new BigNumber(9)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f18))).Intersect(_619_v0);
        _nw91[(new BigNumber(10)).toNumber()] = (_619_v0).Intersect(_619_v0);
        _nw91[(new BigNumber(11)).toNumber()] = (_619_v0).Difference(_619_v0);
        _nw91[(new BigNumber(12)).toNumber()] = (_619_v0).Intersect(_619_v0);
        _nw91[(new BigNumber(13)).toNumber()] = _619_v0;
        _nw91[(new BigNumber(14)).toNumber()] = _619_v0;
        _nw91[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.FromArray(_624_v5);
        _nw91[(new BigNumber(16)).toNumber()] = (_dafny.MultiSet.FromArray(_624_v5)).update((((_625_v6).contains(_623_v4)) ? ((_625_v6).get(_623_v4)) : (p1)), _module.__default.abs(_623_v4));
        _nw91[(new BigNumber(17)).toNumber()] = _619_v0;
        _nw91[(new BigNumber(18)).toNumber()] = _619_v0;
        _626_v7 = _nw91;
        let _627_v8;
        _627_v8 = _module.D6.create_DC17(_626_v7);
        _626_v7 = (_627_v8).dtor_cf38;
      } else {
        let _628_v9;
        let _nw92 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
        _628_v9 = _nw92;
        let _629_v10;
        _629_v10 = new BigNumber(164);
        let _630_v11;
        _630_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,_629_v10);
        let _index78 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_628_v9).length));
        (_628_v9)[_index78] = _630_v11;
        let _index79 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_628_v9).length));
        (_628_v9)[_index79] = _630_v11;
        let _631_v12;
        let _init9 = ((_632_v10) => function (_633_i0) {
          return (_632_v10).isLessThan(_632_v10);
        })(_629_v10);
        let _nw93 = Array((new BigNumber(6)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw93.length); _i0_9++) {
          _nw93[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _631_v12 = _nw93;
        _631_v12 = _631_v12;
        (globalState).f0 = _629_v10;
        if (p2) {
          let _634_v13;
          let _nw94 = Array((new BigNumber(2)).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("r")).length);
          _nw94[(_dafny.ONE).toNumber()] = (new BigNumber((_module.__default.fm39(p1, new BigNumber(288), _629_v10, globalState)).length)).multipliedBy(_629_v10);
          _634_v13 = _nw94;
          let _index80 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_634_v13).length));
          (_634_v13)[_index80] = _629_v10;
          let _635_v15;
          _635_v15 = _dafny.Map.Empty.slice().updateUnsafe(_629_v10,_629_v10);
          let _636_v16;
          _636_v16 = _dafny.Set.fromElements(!(function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of _dafny.IntegerRange(new BigNumber(309), new BigNumber(139))) {
              let _637_v14 = _compr_25;
              if (((new BigNumber(309)).isLessThanOrEqualTo(_637_v14)) && ((_637_v14).isLessThan(new BigNumber(139)))) {
                _coll25.push([_module.__default.safeDivisionInt(_637_v14, _629_v10),_629_v10]);
              }
            }
            return _coll25;
          }()).equals(_635_v15), (_this).f18);
          let _638_v17;
          _638_v17 = _dafny.Map.Empty.slice().updateUnsafe((_630_v11).Merge((_628_v9)[_module.__default.safeIndex(new BigNumber(679), new BigNumber((_628_v9).length))]),_636_v16);
          let _index81 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_634_v13).length));
          let _rhs79 = _629_v10;
          let _rhs80 = (((_638_v17).contains(_630_v11)) ? ((_638_v17).get(_630_v11)) : (_636_v16));
          let _rhs81 = _629_v10;
          let _lhs61 = _634_v13;
          let _lhs62 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_634_v13).length));
          let _lhs63 = globalState;
          _lhs61[_lhs62] = _rhs79;
          _636_v16 = _rhs80;
          _lhs63.f15 = _rhs81;
          let _639_v18;
          _639_v18 = _dafny.Map.Empty.slice().updateUnsafe(true,(p1) === (p1));
          let _640_v19;
          _640_v19 = _631_v12;
          let _641_v20;
          _641_v20 = _dafny.Map.Empty.slice().updateUnsafe(_631_v12,_629_v10);
          _639_v18 = (_639_v18).update((_641_v20).contains((_640_v19)), p1);
          let _642_v21;
          _642_v21 = _module.D2.create_DC7(_629_v10, !(p1), _629_v10);
          (globalState).f10 = _dafny.Seq.update(_module.__default.fm40(_642_v21, globalState), _module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber((_module.__default.fm40(_642_v21, globalState)).length)), _629_v10);
          let _643_v22;
          _643_v22 = _dafny.Seq.of(p1);
          let _644_v23;
          _644_v23 = _dafny.Map.Empty.slice().updateUnsafe(_643_v22,p1);
          _629_v10 = (((((_644_v23).contains(_643_v22)) ? ((_644_v23).get(_643_v22)) : (p2))) ? ((_634_v13)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_634_v13).length))]) : (_629_v10));
          let _645_v24;
          _645_v24 = _dafny.Seq.of(_631_v12);
          let _646_v25;
          let _nw95 = Array((new BigNumber(22)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = _631_v12;
          _nw95[(_dafny.ONE).toNumber()] = _631_v12;
          _nw95[(new BigNumber(2)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(3)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(4)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(5)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(6)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(7)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(8)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(9)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(10)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(11)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(12)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(13)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(14)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(15)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(16)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(17)).toNumber()] = (_645_v24)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(839)), ((_647_v10) => function (_648_i1) {
            return _647_v10;
          })(_629_v10))).length), new BigNumber((_645_v24).length))];
          _nw95[(new BigNumber(18)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(19)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(20)).toNumber()] = _631_v12;
          _nw95[(new BigNumber(21)).toNumber()] = _631_v12;
          _646_v25 = _nw95;
          let _index82 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_646_v25).length));
          (_646_v25)[_index82] = _631_v12;
          let _649_v26;
          _649_v26 = _dafny.Map.Empty.slice().updateUnsafe(p1,_643_v22);
          let _650_v27;
          _650_v27 = _dafny.MultiSet.fromElements(_630_v11);
          let _651_v28;
          _651_v28 = _dafny.Map.Empty.slice().updateUnsafe(_629_v10,p2);
          let _652_v29;
          _652_v29 = _dafny.MultiSet.fromElements(_629_v10, new BigNumber(((_this).f19).length), new BigNumber((_651_v28).length));
          let _index83 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_646_v25).length));
          let _rhs82 = ((p2) ? (_631_v12) : (_631_v12));
          let _rhs83 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(596)), function (_653_i2) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          }), _module.__default.safeIndex(new BigNumber(((_649_v26).update((_this).f18, _643_v22)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(596)), function (_654_i2) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length)), new _dafny.CodePoint('j'.codePointAt(0)));
          let _rhs84 = (((_650_v27).contains((_628_v9)[_module.__default.safeIndex(new BigNumber(679), new BigNumber((_628_v9).length))])) ? ((_650_v27).get((_628_v9)[_module.__default.safeIndex(new BigNumber(679), new BigNumber((_628_v9).length))])) : ((((_652_v29).contains(_629_v10)) ? ((_652_v29).get(_629_v10)) : (new BigNumber(((_this).f19).length)))));
          let _rhs85 = (((_619_v0).contains(!(!(!(p2))) || (p2))) ? ((_619_v0).get(!(!(!(p2))) || (p2))) : (_629_v10));
          let _lhs64 = _646_v25;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_646_v25).length));
          let _lhs66 = globalState;
          let _lhs67 = globalState;
          let _lhs68 = globalState;
          _lhs64[_lhs65] = _rhs82;
          _lhs66.f9 = _rhs83;
          _lhs67.f15 = _rhs84;
          _lhs68.f15 = _rhs85;
        } else {
          let _655_v30;
          let _nw96 = new _module.C1();
          _nw96.__ctor(p2);
          _655_v30 = _nw96;
          let _656_v31;
          let _nw97 = new _module.C1();
          _nw97.__ctor(!(p1));
          _656_v31 = _nw97;
          (globalState).f3 = _module.__default.safeDivisionInt((new BigNumber(-898)).multipliedBy(_629_v10), _629_v10);
          let _657_v32;
          let _nw98 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _657_v32 = _nw98;
          let _658_v33;
          _658_v33 = _657_v32;
          _658_v33 = _658_v33;
          let _659_v34;
          _659_v34 = _dafny.Seq.of(p2);
          let _index84 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_657_v32).length));
          (_657_v32)[_index84] = new BigNumber((_659_v34).length);
          let _index85 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_657_v32).length));
          (_657_v32)[_index85] = (_dafny.ZERO).minus(_629_v10);
        }
        (globalState).f2 = ((true) ? (p1) : ((_this).f18));
      }
      (globalState).f2 = (_this).f18;
      let _660_v35;
      let _init10 = ((_661_p1, _662_p2) => function (_663_i3) {
        return (_dafny.Map.Empty.slice().updateUnsafe(!(_661_p1),(_this).f18)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_662_p2));
      })(p1, p2);
      let _nw99 = Array((new BigNumber(4)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw99.length); _i0_10++) {
        _nw99[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _660_v35 = _nw99;
      let _664_v36;
      _664_v36 = _module.D3.create_DC11(p1, (_this).f18);
      let _pat_let_tv7 = p1;
      let _pat_let_tv8 = p2;
      let _pat_let_tv9 = p2;
      let _pat_let_tv10 = p1;
      let _pat_let_tv11 = _619_v0;
      let _pat_let_tv12 = _619_v0;
      let _pat_let_tv13 = _619_v0;
      let _pat_let_tv14 = p1;
      let _rhs86 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f19, (_this).f19), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jnb"), (_this).f19));
      let _rhs87 = _660_v35;
      let _rhs88 = function (_source5) {
        if (_source5.is_DC11) {
          let _665___mcc_h0 = (_source5).cf23;
          let _666___mcc_h1 = (_source5).cf24;
          let _667_cf24 = _666___mcc_h1;
          let _668_cf23 = _665___mcc_h0;
          return new BigNumber((((!(_668_cf23)) ? (_dafny.Seq.of(_667_cf24)) : (_dafny.Seq.of(_pat_let_tv7)))).length);
        } else if (_source5.is_DC12) {
          let _669___mcc_h2 = (_source5).cf25;
          let _670___mcc_h3 = (_source5).cf26;
          let _671___mcc_h4 = (_source5).cf27;
          let _672___mcc_h5 = (_source5).cf28;
          let _673___mcc_h6 = (_source5).cf29;
          let _674_cf29 = _673___mcc_h6;
          let _675_cf28 = _672___mcc_h5;
          let _676_cf27 = _671___mcc_h4;
          let _677_cf26 = _670___mcc_h3;
          let _678_cf25 = _669___mcc_h2;
          return _module.__default.safeDivisionInt(_676_cf27, _676_cf27);
        } else if (_source5.is_DC13) {
          let _679___mcc_h7 = (_source5).cf30;
          let _680___mcc_h8 = (_source5).cf31;
          let _681___mcc_h9 = (_source5).cf32;
          let _682___mcc_h10 = (_source5).cf33;
          let _683___mcc_h11 = (_source5).cf34;
          let _684_cf34 = _683___mcc_h11;
          let _685_cf33 = _682___mcc_h10;
          let _686_cf32 = _681___mcc_h9;
          let _687_cf31 = _680___mcc_h8;
          let _688_cf30 = _679___mcc_h7;
          return _module.__default.safeDivisionInt(_685_cf33, new BigNumber((_dafny.Set.fromElements(_pat_let_tv8)).length));
        } else if (_source5.is_DC10) {
          let _689___mcc_h12 = (_source5).cf22;
          let _690_cf22 = _689___mcc_h12;
          return (_module.D0.create_DC1((_module.D0.create_DC0((_this).f18)).dtor_cf0, new BigNumber(971), (_this).f19, _pat_let_tv9, new BigNumber(-299))).dtor_cf5;
        } else {
          let _691___mcc_h13 = (_source5).cf35;
          let _692_cf35 = _691___mcc_h13;
          return new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber(((_this).f19).length))).Union(_dafny.MultiSet.fromElements(new BigNumber(((_module.D3.create_DC13(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv10,new BigNumber(((_this).f19).length))).length), _dafny.Seq.UnicodeFromString("ukjh"), new BigNumber((_pat_let_tv11).cardinality()), (((_pat_let_tv13).contains(true)) ? ((_pat_let_tv12).get(true)) : (new BigNumber(((_this).f19).length))), new _dafny.CodePoint('i'.codePointAt(0)))).dtor_cf31).length)))).cardinality());
        }
      }(function (_pat_let26_0) {
        return function (_693_dt__update__tmp_h0) {
          return function (_pat_let27_0) {
            return function (_694_dt__update_hcf23_h0) {
              return _module.D3.create_DC11(_694_dt__update_hcf23_h0, (_693_dt__update__tmp_h0).dtor_cf24);
            }(_pat_let27_0);
          }(_pat_let_tv14);
        }(_pat_let26_0);
      }(_664_v36));
      let _lhs69 = globalState;
      let _lhs70 = globalState;
      _lhs69.f12 = _rhs86;
      _660_v35 = _rhs87;
      _lhs70.f3 = _rhs88;
      let _695_v37;
      _695_v37 = new BigNumber(129);
      let _696_i4;
      _696_i4 = _dafny.ZERO;
      L4: {
        while ((_this).fm3(_695_v37, globalState)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_696_i4)) {
              break L4;
            }
            _696_i4 = (_696_i4).plus(_dafny.ONE);
            let _697_v38;
            let _nw100 = Array((new BigNumber(9)).toNumber());
            _nw100[(_dafny.ZERO).toNumber()] = _695_v37;
            _nw100[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_695_v37));
            _nw100[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_695_v37);
            _nw100[(new BigNumber(3)).toNumber()] = new BigNumber(-119);
            _nw100[(new BigNumber(4)).toNumber()] = _695_v37;
            _nw100[(new BigNumber(5)).toNumber()] = _695_v37;
            _nw100[(new BigNumber(6)).toNumber()] = _695_v37;
            _nw100[(new BigNumber(7)).toNumber()] = _695_v37;
            _nw100[(new BigNumber(8)).toNumber()] = _695_v37;
            _697_v38 = _nw100;
            let _698_v39;
            _698_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_697_v38);
            _698_v39 = (_698_v39).update((_this).f19, _697_v38);
            let _699_v40;
            _699_v40 = new _dafny.CodePoint('u'.codePointAt(0));
            let _index86 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_697_v38).length));
            (_697_v38)[_index86] = _695_v37;
            let _index87 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_697_v38).length));
            let _rhs89 = _699_v40;
            let _rhs90 = _695_v37;
            let _lhs71 = _697_v38;
            let _lhs72 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_697_v38).length));
            _699_v40 = _rhs89;
            _lhs71[_lhs72] = _rhs90;
            (globalState).f0 = _695_v37;
            (globalState).f3 = new BigNumber(117);
          }
        }
      }
      let _700_i5;
      _700_i5 = _dafny.ZERO;
      L5: {
        while (p2) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_700_i5)) {
              break L5;
            }
            _700_i5 = (_700_i5).plus(_dafny.ONE);
            let _701_v41;
            let _init11 = ((_702_v37) => function (_703_i6) {
              return (_703_i6).multipliedBy(_702_v37);
            })(_695_v37);
            let _nw101 = Array((new BigNumber(4)).toNumber());
            for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw101.length); _i0_11++) {
              _nw101[_i0_11] = _init11(new BigNumber(_i0_11));
            }
            _701_v41 = _nw101;
            let _704_v42;
            _704_v42 = _dafny.Seq.of((_this).f18, p1, (_this).f18, (_this).f18, p2);
            let _index88 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_701_v41).length));
            (_701_v41)[_index88] = new BigNumber((_704_v42).length);
            let _705_v43;
            let _nw102 = Array((new BigNumber(13)).toNumber()).fill(false);
            _705_v43 = _nw102;
            let _706_v44;
            _706_v44 = _dafny.Seq.of(_705_v43, _705_v43, _705_v43);
            let _707_v45;
            _707_v45 = _dafny.Seq.of(_706_v44, _dafny.Seq.Concat(_706_v44, _706_v44));
            let _708_v46;
            _708_v46 = _dafny.MultiSet.fromElements(_695_v37);
            let _index89 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_701_v41).length));
            let _rhs91 = _695_v37;
            let _rhs92 = (_707_v45)[_module.__default.safeIndex((((_708_v46).contains(_695_v37)) ? ((_708_v46).get(_695_v37)) : (_695_v37)), new BigNumber((_707_v45).length))];
            let _rhs93 = new BigNumber((_dafny.Seq.Concat(_704_v42, _dafny.Seq.update(_704_v42, _module.__default.safeIndex(new BigNumber(368), new BigNumber((_704_v42).length)), p2))).length);
            let _lhs73 = _701_v41;
            let _lhs74 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_701_v41).length));
            _lhs73[_lhs74] = _rhs91;
            _706_v44 = _rhs92;
            _695_v37 = _rhs93;
            let _index90 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_701_v41).length));
            (_701_v41)[_index90] = _module.__default.safeModuloInt(((_this).f35)[_module.__default.safeIndex(new BigNumber(771), new BigNumber(((_this).f35).length))], (new BigNumber(287)).minus(new BigNumber(583)));
            (globalState).f2 = false;
            let _709_v47;
            _709_v47 = _dafny.Map.Empty.slice().updateUnsafe(true,_705_v43);
            let _710_v48;
            _710_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(227),(_this).f35);
            let _711_v49;
            let _nw103 = new _module.C2();
            _nw103.__ctor(new BigNumber((_709_v47).length), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(774)), function (_712_i7) {
              return ((_this).f35)[_module.__default.safeIndex(new BigNumber(((_this).f35).length), new BigNumber(((_this).f35).length))];
            }), (((_710_v48).contains((_701_v41)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_701_v41).length))])) ? ((_710_v48).get((_701_v41)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_701_v41).length))])) : ((_this).f35))), (_this).f18);
            _711_v49 = _nw103;
          }
        }
      }
      r0 = (_this).fm38(_dafny.Seq.UnicodeFromString("vhgblgpc"), globalState);
      return r0;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _713_v0;
      let _init12 = function (_714_i1) {
        return (_this).f18;
      };
      let _nw104 = Array((new BigNumber(24)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw104.length); _i0_12++) {
        _nw104[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _713_v0 = _nw104;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_713_v0).length))) {
        let _715_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_715_i0)) && ((_715_i0).isLessThan(new BigNumber((_713_v0).length))))) {
          (_713_v0)[(_715_i0)] = ((false) ? ((_this).f18) : (((_dafny.ZERO).minus(p1)).isLessThanOrEqualTo(p1)));
        }
      }
      let _716_v1;
      let _init13 = function (_717_i2) {
        return _dafny.Seq.Concat(_dafny.Seq.of((_this).f18, (_this).f18), _dafny.Seq.of((_this).f18, true, true, (_this).f18));
      };
      let _nw105 = Array((new BigNumber(10)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw105.length); _i0_13++) {
        _nw105[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _716_v1 = _nw105;
      let _718_v2;
      _718_v2 = _dafny.Seq.of((_this).f18, (_this).f18);
      let _index91 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_716_v1).length));
      (_716_v1)[_index91] = _718_v2;
      let _index92 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_716_v1).length));
      (_716_v1)[_index92] = _718_v2;
      (globalState).f0 = p1;
      (globalState).f2 = false;
      let _719_v3;
      let _nw106 = new _module.C1();
      _nw106.__ctor((_this).f18);
      _719_v3 = _nw106;
      (globalState).f2 = (_this).f18;
      return;
    }
    m9(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = false;
      let _720_v0;
      _720_v0 = new BigNumber(154);
      let _721_v1;
      _721_v1 = _dafny.Seq.of(!((_this).fm3(_720_v0, globalState)));
      let _722_i0;
      _722_i0 = _dafny.ZERO;
      L6: {
        while (_dafny.Seq.contains(_721_v1, (_this).f18)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_722_i0)) {
              break L6;
            }
            _722_i0 = (_722_i0).plus(_dafny.ONE);
            let _rhs94 = (_dafny.ZERO).minus(new BigNumber(-718));
            let _rhs95 = _720_v0;
            let _rhs96 = (_this).f18;
            let _lhs75 = globalState;
            let _lhs76 = globalState;
            let _lhs77 = globalState;
            _lhs75.f0 = _rhs94;
            _lhs76.f3 = _rhs95;
            _lhs77.f2 = _rhs96;
            (globalState).f0 = new BigNumber(((_this).f19).length);
            let _723_v2;
            let _init14 = ((_724_v0, _725_v1) => function (_726_i1) {
              return ((_module.D8.create_DC20(_dafny.Set.fromElements(_724_v0))).dtor_cf40).Difference(_dafny.Set.fromElements(new BigNumber(((_this).f19).length), (_module.D0.create_DC1((_this).f18, _724_v0, (_this).f19, true, new BigNumber((_dafny.MultiSet.FromArray((_this).f35)).cardinality()))).dtor_cf5, (_module.D2.create_DC7(new BigNumber((_725_v1).length), (_this).f18, new BigNumber((_dafny.Set.fromElements((_this).f18, (_this).f18)).length))).dtor_cf15));
            })(_720_v0, _721_v1);
            let _nw107 = Array((new BigNumber(20)).toNumber());
            for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw107.length); _i0_14++) {
              _nw107[_i0_14] = _init14(new BigNumber(_i0_14));
            }
            _723_v2 = _nw107;
            _723_v2 = _723_v2;
            _720_v0 = ((_this).f35)[_module.__default.safeIndex(_720_v0, new BigNumber(((_this).f35).length))];
          }
        }
      }
      let _727_v3;
      _727_v3 = _module.D1.create_DC4((_this).f18);
      let _728_v4;
      let _nw108 = Array((new BigNumber(16)).toNumber());
      _nw108[(_dafny.ZERO).toNumber()] = (_this).f18;
      _nw108[(_dafny.ONE).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(2)).toNumber()] = (_727_v3).dtor_cf12;
      _nw108[(new BigNumber(3)).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(4)).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(5)).toNumber()] = true;
      _nw108[(new BigNumber(6)).toNumber()] = (_module.__default.fm41(globalState)).dtor_cf0;
      _nw108[(new BigNumber(7)).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(8)).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(9)).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(10)).toNumber()] = (_721_v1)[_module.__default.safeIndex(_720_v0, new BigNumber((_721_v1).length))];
      _nw108[(new BigNumber(11)).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(12)).toNumber()] = !((_this).f18);
      _nw108[(new BigNumber(13)).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(14)).toNumber()] = (_this).f18;
      _nw108[(new BigNumber(15)).toNumber()] = (_this).f18;
      _728_v4 = _nw108;
      let _729_v5;
      _729_v5 = _dafny.Map.Empty.slice().updateUnsafe(_728_v4,_728_v4);
      let _730_v6;
      let _nw109 = Array((new BigNumber(14)).toNumber());
      _nw109[(_dafny.ZERO).toNumber()] = (_729_v5).Merge(_dafny.Map.Empty.slice().updateUnsafe(_728_v4,_728_v4));
      _nw109[(_dafny.ONE).toNumber()] = _729_v5;
      _nw109[(new BigNumber(2)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(3)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(4)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(5)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(6)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(7)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(8)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(9)).toNumber()] = (_729_v5).Merge(_729_v5);
      _nw109[(new BigNumber(10)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(11)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(12)).toNumber()] = _729_v5;
      _nw109[(new BigNumber(13)).toNumber()] = (_729_v5).update(_728_v4, _728_v4);
      _730_v6 = _nw109;
      let _index93 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_730_v6).length));
      (_730_v6)[_index93] = (((_this).f18) ? (_729_v5) : ((_729_v5).update(_728_v4, _728_v4)));
      let _index94 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_730_v6).length));
      (_730_v6)[_index94] = _729_v5;
      let _731_v7;
      let _nw110 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _731_v7 = _nw110;
      let _index95 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
      (_731_v7)[_index95] = _720_v0;
      let _index96 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
      (_731_v7)[_index96] = new BigNumber(257);
      let _732_v8;
      _732_v8 = _dafny.Set.fromElements(_720_v0);
      let _733_i2;
      _733_i2 = _dafny.ZERO;
      L7: {
        while ((_732_v8).equals((_732_v8).Union(_732_v8))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_733_i2)) {
              break L7;
            }
            _733_i2 = (_733_i2).plus(_dafny.ONE);
            let _734_v9;
            _734_v9 = _dafny.Map.Empty.slice().updateUnsafe((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))],(_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))]);
            let _735_v10;
            _735_v10 = _dafny.Set.fromElements((_this).f18, false);
            let _736_v11;
            _736_v11 = new _dafny.CodePoint('b'.codePointAt(0));
            let _737_v12;
            _737_v12 = _module.D3.create_DC13((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))], (_this).f19, _module.__default.safeModuloInt(new BigNumber(-519), (_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))]), (((_734_v9).contains(_720_v0)) ? ((_734_v9).get(_720_v0)) : (new BigNumber((_735_v10).length))), _736_v11);
            let _source6 = _737_v12;
            if (_source6.is_DC11) {
              let _738___mcc_h0 = (_source6).cf23;
              let _739___mcc_h1 = (_source6).cf24;
              let _740_cf24 = _739___mcc_h1;
              let _741_cf23 = _738___mcc_h0;
              (globalState).f15 = _720_v0;
              (globalState).f2 = !(((_740_cf24) ? (_740_cf24) : ((_this).fm2((_this).f18, (_this).f18, _740_cf24, globalState))));
              let _742_v13;
              let _nw111 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
              _742_v13 = _nw111;
              let _743_v14;
              _743_v14 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("tvysx"), _dafny.Seq.UnicodeFromString("mfamwn"));
              let _index97 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_742_v13).length));
              (_742_v13)[_index97] = (_743_v14).Union(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), ((_744_v11) => function (_745_i3) {
                return _744_v11;
              })(_736_v11))));
              let _index98 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_742_v13).length));
              (_742_v13)[_index98] = (_743_v14).Difference(_dafny.MultiSet.fromElements((_this).f19, (_this).f19, (_this).f19, (_this).fm37(_720_v0, globalState)));
              (globalState).f9 = _dafny.Seq.Concat((_this).f19, (_this).f19);
            } else if (_source6.is_DC12) {
              let _746___mcc_h2 = (_source6).cf25;
              let _747___mcc_h3 = (_source6).cf26;
              let _748___mcc_h4 = (_source6).cf27;
              let _749___mcc_h5 = (_source6).cf28;
              let _750___mcc_h6 = (_source6).cf29;
              let _751_cf29 = _750___mcc_h6;
              let _752_cf28 = _749___mcc_h5;
              let _753_cf27 = _748___mcc_h4;
              let _754_cf26 = _747___mcc_h3;
              let _755_cf25 = _746___mcc_h2;
              _735_v10 = (_735_v10).Union(_735_v10);
              let _index99 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
              (_731_v7)[_index99] = (_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))];
              (globalState).f2 = _754_cf26;
              (globalState).f3 = new BigNumber(-860);
            } else if (_source6.is_DC13) {
              let _756___mcc_h7 = (_source6).cf30;
              let _757___mcc_h8 = (_source6).cf31;
              let _758___mcc_h9 = (_source6).cf32;
              let _759___mcc_h10 = (_source6).cf33;
              let _760___mcc_h11 = (_source6).cf34;
              let _761_cf34 = _760___mcc_h11;
              let _762_cf33 = _759___mcc_h10;
              let _763_cf32 = _758___mcc_h9;
              let _764_cf31 = _757___mcc_h8;
              let _765_cf30 = _756___mcc_h7;
              (globalState).f3 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_765_cf30));
              let _766_v15;
              _766_v15 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))], _765_cf30),_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), ((_767_v0) => function (_768_i4) {
                return _767_v0;
              })(_720_v0)));
              _766_v15 = (_766_v15).update((_dafny.ZERO).minus(_765_cf30), (_this).f35);
              let _769_v16;
              let _nw112 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
              _769_v16 = _nw112;
              let _index100 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_769_v16).length));
              (_769_v16)[_index100] = _dafny.Map.Empty.slice().updateUnsafe(_763_cf32,new BigNumber((_dafny.Seq.of(_736_v11, _736_v11)).length));
              let _index101 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_769_v16).length));
              (_769_v16)[_index101] = _734_v9;
              (globalState).f10 = (_module.__default.fm42((_this).f18, (_this).f18, (_this).f18, globalState)).dtor_cf43;
            } else if (_source6.is_DC10) {
              let _770___mcc_h12 = (_source6).cf22;
              let _771_cf22 = _770___mcc_h12;
              let _index102 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length));
              (_728_v4)[_index102] = (_this).fm2(false, true, (_this).f18, globalState);
              let _772_v17;
              _772_v17 = _module.D0.create_DC1(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).fm38((_this).f19, globalState)),(_dafny.ZERO).minus(_720_v0))).length), (_this).f19, (_this).f18, (_this).fm38((_this).f19, globalState));
              let _index103 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
              let _index104 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length));
              let _rhs97 = (_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))];
              let _rhs98 = (_772_v17).dtor_cf2;
              let _rhs99 = true;
              let _lhs78 = _731_v7;
              let _lhs79 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
              let _lhs80 = _728_v4;
              let _lhs81 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length));
              _720_v0 = _rhs97;
              _lhs78[_lhs79] = _rhs98;
              _lhs80[_lhs81] = _rhs99;
              let _773_v18;
              _773_v18 = _dafny.Seq.of(_731_v7);
              let _774_v19;
              let _init15 = ((_775_v11) => function (_776_i5) {
                return _775_v11;
              })(_736_v11);
              let _nw113 = Array((_dafny.ONE).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw113.length); _i0_15++) {
                _nw113[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _774_v19 = _nw113;
              let _index105 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_774_v19).length));
              (_774_v19)[_index105] = _736_v11;
              let _777_v20;
              _777_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_773_v18);
              let _index106 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
              let _index107 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_774_v19).length));
              let _index108 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length));
              let _rhs100 = _dafny.Seq.Concat(_dafny.Seq.Concat(_773_v18, _773_v18), _dafny.Seq.update((((_777_v20).contains((_728_v4)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length))])) ? ((_777_v20).get((_728_v4)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length))])) : (_773_v18)), _module.__default.safeIndex((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))], new BigNumber(((((_777_v20).contains((_728_v4)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length))])) ? ((_777_v20).get((_728_v4)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length))])) : (_773_v18))).length)), _731_v7));
              let _rhs101 = (_720_v0).plus((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))]);
              let _rhs102 = _736_v11;
              let _rhs103 = (_728_v4)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length))];
              let _lhs82 = _731_v7;
              let _lhs83 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
              let _lhs84 = _774_v19;
              let _lhs85 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_774_v19).length));
              let _lhs86 = _728_v4;
              let _lhs87 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_728_v4).length));
              _773_v18 = _rhs100;
              _lhs82[_lhs83] = _rhs101;
              _lhs84[_lhs85] = _rhs102;
              _lhs86[_lhs87] = _rhs103;
              let _index109 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
              (_731_v7)[_index109] = new BigNumber(((_this).fm37((_720_v0).minus(_720_v0), globalState)).length);
              let _778_v21;
              let _nw114 = new _module.C0();
              _nw114.__ctor(new BigNumber(((_this).f19).length), (((_734_v9).contains((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))])) ? ((_734_v9).get((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))])) : ((_dafny.ZERO).minus(_720_v0))));
              _778_v21 = _nw114;
            } else {
              let _779___mcc_h13 = (_source6).cf35;
              let _780_cf35 = _779___mcc_h13;
              let _781_v22;
              _781_v22 = _dafny.MultiSet.fromElements(_736_v11, new _dafny.CodePoint('f'.codePointAt(0)), _736_v11);
              _720_v0 = ((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))]).plus((((_781_v22).contains(new _dafny.CodePoint('u'.codePointAt(0)))) ? ((_781_v22).get(new _dafny.CodePoint('u'.codePointAt(0)))) : ((((_734_v9).contains(_720_v0)) ? ((_734_v9).get(_720_v0)) : (_720_v0)))));
              let _782_v23;
              _782_v23 = _dafny.MultiSet.fromElements((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))], _720_v0, new BigNumber((_dafny.Seq.UnicodeFromString("vuiukyo")).length), new BigNumber((_721_v1).length));
              let _783_v24;
              _783_v24 = _dafny.Seq.of(_721_v1, _721_v1);
              r1 = ((((_782_v23).contains(_720_v0)) ? ((_782_v23).get(_720_v0)) : (new BigNumber((_dafny.MultiSet.FromArray(_783_v24)).cardinality())))).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_720_v0, (_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))]));
              let _784_v25;
              let _nw115 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
              _784_v25 = _nw115;
              let _index110 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_784_v25).length));
              (_784_v25)[_index110] = _dafny.MultiSet.fromElements((_this).f18);
              let _index111 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_784_v25).length));
              let _rhs104 = !_dafny.areEqual(_dafny.Seq.of((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))], _720_v0, new BigNumber(-405)), (_this).f35);
              let _rhs105 = _dafny.Seq.of(new BigNumber(746));
              let _rhs106 = !((_this).f18);
              let _rhs107 = _720_v0;
              let _rhs108 = _dafny.MultiSet.fromElements((_this).f18);
              let _lhs88 = globalState;
              let _lhs89 = globalState;
              let _lhs90 = globalState;
              let _lhs91 = _784_v25;
              let _lhs92 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_784_v25).length));
              _lhs88.f2 = _rhs104;
              _lhs89.f10 = _rhs105;
              r1 = _rhs106;
              _lhs90.f15 = _rhs107;
              _lhs91[_lhs92] = _rhs108;
              let _785_v26;
              let _nw116 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
              _785_v26 = _nw116;
              let _index112 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_785_v26).length));
              (_785_v26)[_index112] = (_this).f35;
              let _786_v27;
              _786_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f35).length),(_this).f18);
              let _index113 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_785_v26).length));
              let _rhs109 = (_this).f18;
              let _rhs110 = _720_v0;
              let _rhs111 = !(((_this).f18) || ((_this).f18));
              let _rhs112 = _dafny.Seq.Concat((_this).f35, _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_786_v27).length)), (_this).f35));
              let _rhs113 = _735_v10;
              let _lhs93 = globalState;
              let _lhs94 = globalState;
              let _lhs95 = _785_v26;
              let _lhs96 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_785_v26).length));
              r2 = _rhs109;
              _lhs93.f0 = _rhs110;
              _lhs94.f2 = _rhs111;
              _lhs95[_lhs96] = _rhs112;
              _735_v10 = _rhs113;
            }
            _735_v10 = _dafny.Set.fromElements(true, (!((_this).f18)) || (!((_this).f18)), (_this).f18, (_this).f18);
            let _index114 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
            (_731_v7)[_index114] = (_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))];
            let _index115 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length));
            (_731_v7)[_index115] = (new BigNumber(-145)).multipliedBy((_720_v0).minus((_731_v7)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_731_v7).length))]));
          }
        }
      }
      (globalState).f10 = (_this).f35;
      (globalState).f0 = (_720_v0).multipliedBy(_720_v0);
      r0 = (_this).f19;
      r1 = (_this).f18;
      r2 = (_this).f18;
      return [r0, r1, r2];
    }
    m10(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _787_v0;
      _787_v0 = _module.D10.create_DC24(_dafny.Set.fromElements((_this).f18, p0));
      let _788_v1;
      _788_v1 = _dafny.Set.fromElements((_this).f18);
      if (!((((_787_v0).dtor_cf47).Intersect(_788_v1)).equals(_788_v1))) {
        let _789_v2;
        _789_v2 = new BigNumber(910);
        let _790_v3;
        let _nw117 = new _module.C2();
        _nw117.__ctor(_789_v2, (_this).f35, (_this).fm2((_this).f18, true, p0, globalState));
        _790_v3 = _nw117;
        (globalState).f2 = (_this).fm3((_this).fm38((_this).f19, globalState), globalState);
        let _791_v4;
        let _nw118 = Array((new BigNumber(4)).toNumber()).fill(false);
        _791_v4 = _nw118;
        let _index116 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_791_v4).length));
        (_791_v4)[_index116] = (_this).f18;
        let _index117 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_791_v4).length));
        (_791_v4)[_index117] = p0;
        let _792_v5;
        _792_v5 = _module.D9.create_DC23((_this).f18, p0, _790_v3.f33);
        let _793_v6;
        let _nw119 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _793_v6 = _nw119;
        let _794_v7;
        _794_v7 = _793_v6;
        let _795_v8;
        let _nw120 = Array((new BigNumber(13)).toNumber());
        _nw120[(_dafny.ZERO).toNumber()] = _793_v6;
        _nw120[(_dafny.ONE).toNumber()] = _793_v6;
        _nw120[(new BigNumber(2)).toNumber()] = _793_v6;
        _nw120[(new BigNumber(3)).toNumber()] = _794_v7;
        _nw120[(new BigNumber(4)).toNumber()] = _794_v7;
        _nw120[(new BigNumber(5)).toNumber()] = _793_v6;
        _nw120[(new BigNumber(6)).toNumber()] = _794_v7;
        _nw120[(new BigNumber(7)).toNumber()] = _794_v7;
        _nw120[(new BigNumber(8)).toNumber()] = _793_v6;
        _nw120[(new BigNumber(9)).toNumber()] = _794_v7;
        _nw120[(new BigNumber(10)).toNumber()] = _794_v7;
        _nw120[(new BigNumber(11)).toNumber()] = _793_v6;
        _nw120[(new BigNumber(12)).toNumber()] = _794_v7;
        _795_v8 = _nw120;
        let _796_v9;
        _796_v9 = _dafny.Seq.of(_795_v8);
        let _797_v10;
        _797_v10 = _dafny.Seq.of(_793_v6);
        let _798_v11;
        _798_v11 = _dafny.Seq.of(p0);
        let _799_v12;
        _799_v12 = _dafny.Map.Empty.slice().updateUnsafe(_790_v3.f33,new BigNumber((_798_v11).length));
        let _800_v13;
        _800_v13 = _module.D0.create_DC1(p0, new BigNumber((_799_v12).length), (_this).f19, (_this).f18, _789_v2);
        let _801_v14;
        _801_v14 = _module.D1.create_DC3((_792_v5).dtor_cf46, new BigNumber(589), _module.__default.fm1((_dafny.ZERO).minus(_790_v3.f33), true, (_791_v4)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_791_v4).length))], globalState), (new BigNumber((_796_v9).length)).plus(_789_v2), _module.__default.safeModuloInt(new BigNumber((_797_v10).length), (_800_v13).dtor_cf2));
        _801_v14 = _801_v14;
        (globalState).f3 = _module.__default.safeDivisionInt(_790_v3.f33, _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(53)), ((_802_v2) => function (_803_i0) {
          return (_dafny.ZERO).minus(_802_v2);
        })(_789_v2))).length)), new BigNumber(((_this).f19).length)));
      } else {
        (globalState).f2 = (_this).f18;
        (globalState).f2 = (_this).f18;
        let _804_v15;
        _804_v15 = _dafny.MultiSet.fromElements(p0);
        let _805_v16;
        let _nw121 = new _module.C2();
        _nw121.__ctor(new BigNumber((_804_v15).cardinality()), (_this).f35, (_this).f18);
        _805_v16 = _nw121;
        let _nw122 = new _module.C2();
        _nw122.__ctor(_module.__default.safeDivisionInt(_805_v16.f33, _805_v16.f33), (_805_v16).f34, p0);
        _805_v16 = _nw122;
        let _806_v17;
        _806_v17 = _dafny.Map.Empty.slice().updateUnsafe(_805_v16.f33,(_805_v16.f33).isLessThan(new BigNumber(505)));
        let _807_v18;
        _807_v18 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_805_v16.f33,(_this).fm3(_805_v16.f33, globalState)));
        _806_v17 = (_806_v17).update(_805_v16.f33, (_807_v18).IsDisjointFrom(_807_v18));
        let _808_v19;
        let _init16 = function (_809_i1) {
          return (_this).f18;
        };
        let _nw123 = Array((new BigNumber(3)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw123.length); _i0_16++) {
          _nw123[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _808_v19 = _nw123;
        let _810_v20;
        _810_v20 = _dafny.Seq.of(_808_v19, _808_v19, _808_v19);
        let _811_v21;
        _811_v21 = _808_v19;
        let _812_v22;
        _812_v22 = new _dafny.CodePoint('a'.codePointAt(0));
        let _813_v23;
        _813_v23 = _dafny.Map.Empty.slice().updateUnsafe(_812_v22,_808_v19);
        let _814_v24;
        let _nw124 = Array((new BigNumber(25)).toNumber());
        _nw124[(_dafny.ZERO).toNumber()] = _808_v19;
        _nw124[(_dafny.ONE).toNumber()] = _808_v19;
        _nw124[(new BigNumber(2)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(3)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(4)).toNumber()] = ((p0) ? (_808_v19) : (_808_v19));
        _nw124[(new BigNumber(5)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(6)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(7)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(8)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(9)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(10)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(11)).toNumber()] = (_810_v20)[_module.__default.safeIndex(_805_v16.f33, new BigNumber((_810_v20).length))];
        _nw124[(new BigNumber(12)).toNumber()] = (_811_v21);
        _nw124[(new BigNumber(13)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(14)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(15)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(16)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(17)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(18)).toNumber()] = (((_813_v23).contains(_812_v22)) ? ((_813_v23).get(_812_v22)) : (_808_v19));
        _nw124[(new BigNumber(19)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(20)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(21)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(22)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(23)).toNumber()] = _808_v19;
        _nw124[(new BigNumber(24)).toNumber()] = _808_v19;
        _814_v24 = _nw124;
        _814_v24 = _814_v24;
      }
      let _815_v25;
      let _nw125 = Array((_dafny.ONE).toNumber());
      _nw125[(_dafny.ZERO).toNumber()] = (_this).f18;
      _815_v25 = _nw125;
      let _index118 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
      (_815_v25)[_index118] = true;
      let _index119 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
      (_815_v25)[_index119] = p0;
      let _816_v26;
      _816_v26 = new BigNumber(89);
      let _817_v27;
      let _nw126 = new _module.C0();
      _nw126.__ctor(_module.__default.safeDivisionInt(_816_v26, _816_v26), _816_v26);
      _817_v27 = _nw126;
      _817_v27 = _817_v27;
      let _818_v28;
      _818_v28 = _dafny.Seq.of(p0);
      let _819_v29;
      _819_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_818_v28).length),(_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]);
      if ((!(_819_v29).contains((_dafny.ZERO).minus((_817_v27).f32))) || ((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))])) {
        (globalState).f0 = new BigNumber(832);
        let _820_v30;
        _820_v30 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f18) ? (_815_v25) : (_815_v25)),((_this).f35)[_module.__default.safeIndex(new BigNumber(571), new BigNumber(((_this).f35).length))]);
        _820_v30 = (_820_v30).update(_815_v25, (_dafny.ZERO).minus(_816_v26));
        if ((((_this).f18) ? (p0) : (((_817_v27).f32).isEqualTo(_816_v26)))) {
          _819_v29 = (_819_v29).update(_816_v26, (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]);
          _818_v28 = _818_v28;
          (globalState).f3 = (_817_v27).f32;
          (globalState).f3 = (_817_v27).f31;
          let _index120 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
          (_815_v25)[_index120] = (_this).f18;
        } else {
          (globalState).f15 = (_817_v27).f31;
          let _index121 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
          (_815_v25)[_index121] = false;
          let _821_v31;
          _821_v31 = new _dafny.CodePoint('m'.codePointAt(0));
          let _822_v32;
          _822_v32 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_818_v28, (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]),_821_v31);
          _822_v32 = (_822_v32).update((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], _821_v31);
          let _index122 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
          let _rhs114 = (((_this).f18) ? ((_818_v28)[_module.__default.safeIndex(new BigNumber((_module.__default.fm43((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], p0, globalState)).length), new BigNumber((_818_v28).length))]) : (p0));
          let _rhs115 = (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))];
          let _lhs97 = globalState;
          let _lhs98 = _815_v25;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
          _lhs97.f2 = _rhs114;
          _lhs98[_lhs99] = _rhs115;
          let _823_v33;
          let _init17 = ((_824_v27) => function (_825_i2) {
            return (_825_i2).minus((_824_v27).f32);
          })(_817_v27);
          let _nw127 = Array((new BigNumber(19)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw127.length); _i0_17++) {
            _nw127[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _823_v33 = _nw127;
          let _826_v34;
          _826_v34 = _dafny.Map.Empty.slice().updateUnsafe(_823_v33,_823_v33);
          _826_v34 = (_826_v34).update((((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]) ? (_823_v33) : (_823_v33)), _823_v33);
        }
        let _827_v35;
        _827_v35 = _dafny.MultiSet.fromElements((_817_v27).f31, _816_v26, new BigNumber((_818_v28).length), (_817_v27).f32, _816_v26);
        let _828_v36;
        _828_v36 = new _dafny.CodePoint('j'.codePointAt(0));
        let _829_v37;
        _829_v37 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("dwvvdp"),new BigNumber(364));
        let _830_v38;
        _830_v38 = _module.D8.create_DC21(_828_v36, new BigNumber((_829_v37).length));
        let _831_v39;
        let _nw128 = Array((new BigNumber(24)).toNumber());
        _nw128[(_dafny.ZERO).toNumber()] = _module.__default.fm1((((_827_v35).contains((_817_v27).f32)) ? ((_827_v35).get((_817_v27).f32)) : ((_817_v27).f31)), (_this).fm2((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], globalState), (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], globalState);
        _nw128[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_818_v28, _818_v28)).length);
        _nw128[(new BigNumber(2)).toNumber()] = ((_817_v27).f31).minus(new BigNumber(((_this).f19).length));
        _nw128[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_817_v27).f31, (_dafny.ZERO).minus(new BigNumber(((_this).f35).length)));
        _nw128[(new BigNumber(4)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(5)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(6)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(7)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(8)).toNumber()] = (_817_v27).f32;
        _nw128[(new BigNumber(9)).toNumber()] = (_830_v38).dtor_cf42;
        _nw128[(new BigNumber(10)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(11)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(12)).toNumber()] = (_817_v27).f32;
        _nw128[(new BigNumber(13)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(14)).toNumber()] = (_816_v26).plus((_817_v27).f31);
        _nw128[(new BigNumber(15)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(16)).toNumber()] = _module.__default.safeDivisionInt(_816_v26, new BigNumber(986));
        _nw128[(new BigNumber(17)).toNumber()] = ((_817_v27).f31).plus((_dafny.ZERO).minus((_817_v27).f32));
        _nw128[(new BigNumber(18)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(19)).toNumber()] = (_817_v27).f32;
        _nw128[(new BigNumber(20)).toNumber()] = (_817_v27).f31;
        _nw128[(new BigNumber(21)).toNumber()] = new BigNumber(816);
        _nw128[(new BigNumber(22)).toNumber()] = new BigNumber(((_this).f19).length);
        _nw128[(new BigNumber(23)).toNumber()] = (_817_v27).f31;
        _831_v39 = _nw128;
        let _index123 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_831_v39).length));
        (_831_v39)[_index123] = new BigNumber((_819_v29).length);
        let _index124 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_831_v39).length));
        (_831_v39)[_index124] = (_817_v27).f32;
        let _832_v40;
        _832_v40 = _module.D1.create_DC4((_this).f18);
        let _source7 = _832_v40;
        if (_source7.is_DC3) {
          let _833___mcc_h0 = (_source7).cf7;
          let _834___mcc_h1 = (_source7).cf8;
          let _835___mcc_h2 = (_source7).cf9;
          let _836___mcc_h3 = (_source7).cf10;
          let _837___mcc_h4 = (_source7).cf11;
          let _838_cf11 = _837___mcc_h4;
          let _839_cf10 = _836___mcc_h3;
          let _840_cf9 = _835___mcc_h2;
          let _841_cf8 = _834___mcc_h1;
          let _842_cf7 = _833___mcc_h0;
          let _843_v41;
          let _nw129 = new _module.C1();
          _nw129.__ctor(!((_this).f18));
          _843_v41 = _nw129;
          let _844_v42;
          _844_v42 = _module.D1.create_DC2(_843_v41);
          let _845_v43;
          _845_v43 = _dafny.Map.Empty.slice().updateUnsafe(_844_v42,_dafny.Seq.UnicodeFromString("nr"));
          let _846_v44;
          _846_v44 = _module.D3.create_DC13((_817_v27).f32, (_this).f19, _842_cf7, (_817_v27).f31, _828_v36);
          let _847_v45;
          _847_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(((_845_v43).contains(_module.D1.create_DC2(_843_v41))) ? ((_845_v43).get(_module.D1.create_DC2(_843_v41))) : ((_846_v44).dtor_cf31)));
          _847_v45 = (_847_v45).update(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("imyjlnfkg"), (_this).f19), (_this).f19);
          _788_v1 = _788_v1;
          let _848_v46;
          _848_v46 = _dafny.Set.fromElements(_828_v36, _828_v36);
          let _849_v47;
          _849_v47 = _module.D2.create_DC7((_dafny.ZERO).minus(_842_cf7), (_843_v41).f18, new BigNumber((_848_v46).length));
          let _850_v48;
          _850_v48 = _module.D2.create_DC9(_849_v47);
          let _851_v49;
          _851_v49 = _module.D2.create_DC9(_850_v48);
          let _852_v50;
          _852_v50 = _dafny.Map.Empty.slice().updateUnsafe(_830_v38,_851_v49);
          let _853_v52;
          _853_v52 = _dafny.Seq.of(_830_v38);
          let _854_v54;
          let _nw130 = new _module.C1();
          _nw130.__ctor((_this).f18);
          _854_v54 = _nw130;
          let _855_v55;
          _855_v55 = _dafny.Map.Empty.slice().updateUnsafe(_854_v54,false);
          let _856_v56;
          _856_v56 = _module.D11.create_DC27(_855_v55);
          let _857_v57;
          let _init18 = ((_858_cf9, _859_cf11) => function (_860_i3) {
            return _dafny.Set.fromElements(_858_cf9, _859_cf11, new BigNumber(((_this).f19).length));
          })(_840_cf9, _838_cf11);
          let _nw131 = Array((new BigNumber(16)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw131.length); _i0_18++) {
            _nw131[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _857_v57 = _nw131;
          let _861_v58;
          _861_v58 = _dafny.Map.Empty.slice().updateUnsafe(_857_v57,(_852_v50).update(_830_v38, _851_v49));
          let _862_v59;
          let _nw132 = Array((new BigNumber(13)).toNumber());
          _nw132[(_dafny.ZERO).toNumber()] = _852_v50;
          _nw132[(_dafny.ONE).toNumber()] = function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of (_853_v52).Elements) {
              let _863_v51 = _compr_26;
              if (_dafny.Seq.contains(_853_v52, _863_v51)) {
                _coll26.push([_863_v51,_851_v49]);
              }
            }
            return _coll26;
          }();
          _nw132[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_830_v38,_851_v49)).Merge(_852_v50);
          _nw132[(new BigNumber(3)).toNumber()] = _852_v50;
          _nw132[(new BigNumber(4)).toNumber()] = _852_v50;
          _nw132[(new BigNumber(5)).toNumber()] = _852_v50;
          _nw132[(new BigNumber(6)).toNumber()] = _852_v50;
          _nw132[(new BigNumber(7)).toNumber()] = function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of (_dafny.Seq.update(_dafny.Seq.of(_830_v38), _module.__default.safeIndex(new BigNumber(((_856_v56).dtor_cf50).length), new BigNumber((_dafny.Seq.of(_830_v38)).length)), _830_v38)).Elements) {
              let _864_v53 = _compr_27;
              if (_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.of(_830_v38), _module.__default.safeIndex(new BigNumber(((_856_v56).dtor_cf50).length), new BigNumber((_dafny.Seq.of(_830_v38)).length)), _830_v38), _864_v53)) {
                _coll27.push([_864_v53,_851_v49]);
              }
            }
            return _coll27;
          }();
          _nw132[(new BigNumber(8)).toNumber()] = _852_v50;
          _nw132[(new BigNumber(9)).toNumber()] = _852_v50;
          _nw132[(new BigNumber(10)).toNumber()] = ((((_861_v58).contains(_857_v57)) ? ((_861_v58).get(_857_v57)) : (_dafny.Map.Empty.slice().updateUnsafe(_830_v38,_851_v49)))).update(_830_v38, _851_v49);
          _nw132[(new BigNumber(11)).toNumber()] = _852_v50;
          _nw132[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_830_v38,_851_v49);
          _862_v59 = _nw132;
          let _index125 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_862_v59).length));
          (_862_v59)[_index125] = (_852_v50).Merge(_dafny.Map.Empty.slice().updateUnsafe(_830_v38,_851_v49));
          let _index126 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_862_v59).length));
          (_862_v59)[_index126] = (_852_v50).Merge(_852_v50);
          let _865_v60;
          _865_v60 = _dafny.Map.Empty.slice().updateUnsafe(_815_v25,(_this).f19);
          _865_v60 = (_865_v60).update(_815_v25, (_this).f19);
        } else if (_source7.is_DC4) {
          let _866___mcc_h5 = (_source7).cf12;
          let _867_cf12 = _866___mcc_h5;
          _828_v36 = _828_v36;
          (globalState).f2 = (_this).f18;
          let _868_v61;
          _868_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,_831_v39);
          let _869_v62;
          _869_v62 = _dafny.Map.Empty.slice().updateUnsafe(_816_v26,_868_v61);
          let _870_v63;
          _870_v63 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_816_v26,(_817_v27).f32)).update((_817_v27).f32, (_831_v39)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_831_v39).length))])).length)).isLessThan((_817_v27).f32),((((_869_v62).contains(new BigNumber(((_this).f19).length))) ? ((_869_v62).get(new BigNumber(((_this).f19).length))) : (_868_v61))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f35,_831_v39)));
          _870_v63 = (_870_v63).update((_this).f18, (_868_v61).Merge(_868_v61));
          let _871_v64;
          let _nw133 = Array((new BigNumber(5)).toNumber()).fill([]);
          _871_v64 = _nw133;
          let _index127 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_871_v64).length));
          (_871_v64)[_index127] = _831_v39;
          let _index128 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_871_v64).length));
          (_871_v64)[_index128] = _831_v39;
        } else if (_source7.is_DC2) {
          let _872___mcc_h6 = (_source7).cf6;
          let _873_cf6 = _872___mcc_h6;
          let _874_v65;
          let _nw134 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _874_v65 = _nw134;
          let _index129 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_874_v65).length));
          (_874_v65)[_index129] = _828_v36;
          let _875_v66;
          _875_v66 = _dafny.MultiSet.fromElements(!(!((_this).f18)), p0, ((_817_v27).f31).isLessThan(_816_v26));
          let _index130 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_874_v65).length));
          let _index131 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_831_v39).length));
          let _rhs116 = _828_v36;
          let _rhs117 = (_dafny.MultiSet.FromArray(_818_v28)).Union(_875_v66);
          let _rhs118 = (_817_v27).f32;
          let _lhs100 = _874_v65;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_874_v65).length));
          let _lhs102 = _831_v39;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_831_v39).length));
          _lhs100[_lhs101] = _rhs116;
          _875_v66 = _rhs117;
          _lhs102[_lhs103] = _rhs118;
          (globalState).f3 = (((_875_v66).contains(true)) ? ((_875_v66).get(true)) : (_module.__default.safeDivisionInt(new BigNumber(-414), (_817_v27).f32)));
          let _index132 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
          (_815_v25)[_index132] = (_this).f18;
          r3 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uyagbxxr"), (_this).f19), (_this).f19);
        } else {
          let _876___mcc_h7 = (_source7).cf13;
          let _877_cf13 = _876___mcc_h7;
          let _878_v67;
          _878_v67 = _dafny.Map.Empty.slice().updateUnsafe((_817_v27).f31,_818_v28);
          let _879_v68;
          _879_v68 = _dafny.Map.Empty.slice().updateUnsafe((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))],_818_v28);
          let _880_v69;
          _880_v69 = _module.D12.create_DC30(_818_v28);
          let _881_v70;
          let _nw135 = Array((new BigNumber(25)).toNumber());
          _nw135[(_dafny.ZERO).toNumber()] = _818_v28;
          _nw135[(_dafny.ONE).toNumber()] = _818_v28;
          _nw135[(new BigNumber(2)).toNumber()] = _dafny.Seq.of((_this).f18);
          _nw135[(new BigNumber(3)).toNumber()] = _818_v28;
          _nw135[(new BigNumber(4)).toNumber()] = _dafny.Seq.of((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]);
          _nw135[(new BigNumber(5)).toNumber()] = (((_878_v67).contains((_817_v27).f32)) ? ((_878_v67).get((_817_v27).f32)) : (_818_v28));
          _nw135[(new BigNumber(6)).toNumber()] = _818_v28;
          _nw135[(new BigNumber(7)).toNumber()] = _module.__default.fm44((_817_v27).f31, (_this).f18, _828_v36, new BigNumber((_dafny.MultiSet.fromElements(_module.__default.fm1((_817_v27).f32, (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], p0, globalState), (_817_v27).f31, (_dafny.ZERO).minus((_817_v27).f31))).cardinality()), globalState);
          _nw135[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_818_v28, _818_v28);
          _nw135[(new BigNumber(9)).toNumber()] = _818_v28;
          _nw135[(new BigNumber(10)).toNumber()] = _818_v28;
          _nw135[(new BigNumber(11)).toNumber()] = _module.__default.fm44((_817_v27).f31, (_this).f18, _828_v36, _816_v26, globalState);
          _nw135[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_818_v28, _module.__default.safeIndex((_831_v39)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_831_v39).length))], new BigNumber((_818_v28).length)), (_818_v28)[_module.__default.safeIndex(_816_v26, new BigNumber((_818_v28).length))]);
          _nw135[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_818_v28, _818_v28);
          _nw135[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_818_v28, _818_v28);
          _nw135[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_818_v28, _module.__default.safeIndex(new BigNumber(-627), new BigNumber((_818_v28).length)), !((_this).f18));
          _nw135[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_818_v28, _818_v28);
          _nw135[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat((((_879_v68).contains((_this).f18)) ? ((_879_v68).get((_this).f18)) : (_818_v28)), _818_v28);
          _nw135[(new BigNumber(18)).toNumber()] = _818_v28;
          _nw135[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_818_v28, _818_v28);
          _nw135[(new BigNumber(20)).toNumber()] = (_880_v69).dtor_cf53;
          _nw135[(new BigNumber(21)).toNumber()] = _818_v28;
          _nw135[(new BigNumber(22)).toNumber()] = _818_v28;
          _nw135[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_818_v28, _818_v28);
          _nw135[(new BigNumber(24)).toNumber()] = _818_v28;
          _881_v70 = _nw135;
          let _882_v71;
          _882_v71 = _dafny.Set.fromElements((_this).f19);
          let _883_v73;
          let _nw136 = Array((new BigNumber(21)).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements((_this).f19);
          _nw136[(_dafny.ONE).toNumber()] = _882_v71;
          _nw136[(new BigNumber(2)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(3)).toNumber()] = (_882_v71).Difference(_882_v71);
          _nw136[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("dr"));
          _nw136[(new BigNumber(5)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(235)), ((_884_v36) => function (_885_i4) {
            return _884_v36;
          })(_828_v36)), (_this).f19);
          _nw136[(new BigNumber(7)).toNumber()] = _module.__default.fm45(globalState);
          _nw136[(new BigNumber(8)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(9)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(10)).toNumber()] = (_module.__default.fm45(globalState)).Difference(function () {
            let _coll28 = new _dafny.Set();
            for (const _compr_28 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_817_v27).f32)).Keys.Elements) {
              let _886_v72 = _compr_28;
              if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_817_v27).f32)).contains(_886_v72)) {
                _coll28.add(_886_v72);
              }
            }
            return _coll28;
          }());
          _nw136[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements((_this).f19);
          _nw136[(new BigNumber(12)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(13)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(14)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(15)).toNumber()] = (_882_v71).Intersect(_dafny.Set.fromElements((_this).f19));
          _nw136[(new BigNumber(16)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(17)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(18)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(19)).toNumber()] = _882_v71;
          _nw136[(new BigNumber(20)).toNumber()] = _882_v71;
          _883_v73 = _nw136;
          let _index133 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_883_v73).length));
          (_883_v73)[_index133] = _882_v71;
          let _887_v74;
          _887_v74 = _dafny.Map.Empty.slice().updateUnsafe(_815_v25,false);
          let _888_v75;
          _888_v75 = _module.D13.create_DC32(_887_v74);
          let _889_v76;
          _889_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_887_v74).Merge((_888_v75).dtor_cf54)).length));
          let _index134 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_883_v73).length));
          let _rhs119 = _787_v0;
          let _rhs120 = _881_v70;
          let _rhs121 = _882_v71;
          let _rhs122 = (_817_v27).f31;
          let _rhs123 = (_889_v76).update((_this).fm2((_this).fm2(false, (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], p0, globalState), (_this).fm2((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], !(p0), p0, globalState), true, globalState), (_817_v27).f31);
          let _lhs104 = _883_v73;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_883_v73).length));
          let _lhs106 = globalState;
          _787_v0 = _rhs119;
          _881_v70 = _rhs120;
          _lhs104[_lhs105] = _rhs121;
          _lhs106.f3 = _rhs122;
          _889_v76 = _rhs123;
          _828_v36 = _828_v36;
          let _890_v77;
          let _nw137 = Array((new BigNumber(13)).toNumber()).fill([]);
          _890_v77 = _nw137;
          let _index135 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_890_v77).length));
          (_890_v77)[_index135] = _815_v25;
          let _891_v78;
          let _nw138 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _891_v78 = _nw138;
          let _892_v79;
          _892_v79 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18),_891_v78);
          let _index136 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_890_v77).length));
          let _index137 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_831_v39).length));
          let _rhs124 = (((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]) ? (_815_v25) : (_815_v25));
          let _rhs125 = new BigNumber(335);
          let _rhs126 = (_this).f18;
          let _rhs127 = _892_v79;
          let _lhs107 = _890_v77;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_890_v77).length));
          let _lhs109 = _831_v39;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_831_v39).length));
          let _lhs111 = globalState;
          _lhs107[_lhs108] = _rhs124;
          _lhs109[_lhs110] = _rhs125;
          _lhs111.f2 = _rhs126;
          _892_v79 = _rhs127;
          let _893_v80;
          let _nw139 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
          _893_v80 = _nw139;
          let _index138 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_893_v80).length));
          (_893_v80)[_index138] = (_this).f35;
          let _index139 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_893_v80).length));
          let _rhs128 = !(!((_818_v28)[_module.__default.safeIndex(_816_v26, new BigNumber((_818_v28).length))])) || ((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]);
          let _rhs129 = p0;
          let _rhs130 = _dafny.Seq.Concat((_this).f35, (_this).f35);
          let _rhs131 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(593)), ((_894_v27) => function (_895_i5) {
            return ((_this).f19)[_module.__default.safeIndex((_894_v27).f31, new BigNumber(((_this).f19).length))];
          })(_817_v27));
          let _lhs112 = globalState;
          let _lhs113 = globalState;
          let _lhs114 = _893_v80;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_893_v80).length));
          let _lhs116 = globalState;
          _lhs112.f2 = _rhs128;
          _lhs113.f2 = _rhs129;
          _lhs114[_lhs115] = _rhs130;
          _lhs116.f9 = _rhs131;
        }
      } else {
        (globalState).f15 = (_dafny.ZERO).minus(_816_v26);
        if ((_this).fm2(p0, (_this).f18, (_this).f18, globalState)) {
          (globalState).f15 = (_817_v27).f32;
          let _896_v81;
          _896_v81 = _dafny.Map.Empty.slice().updateUnsafe(_815_v25,(_dafny.ZERO).minus((_817_v27).f31));
          let _897_v82;
          _897_v82 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(178)), function (_898_i6) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          }),(_817_v27).f31);
          _896_v81 = (_896_v81).update(_815_v25, new BigNumber(((_897_v82).Merge(_897_v82)).length));
          let _899_v83;
          let _nw140 = new _module.C1();
          _nw140.__ctor(p0);
          _899_v83 = _nw140;
          let _900_v84;
          let _nw141 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _900_v84 = _nw141;
          _900_v84 = _900_v84;
          let _index140 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
          (_815_v25)[_index140] = p0;
        } else {
          let _nw142 = Array((new BigNumber(29)).toNumber()).fill(false);
          _815_v25 = _nw142;
          let _index141 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length));
          (_815_v25)[_index141] = ((false) ? ((_this).f18) : ((_this).f18));
          let _rhs132 = ((_817_v27).f31).isEqualTo((_817_v27).f32);
          let _rhs133 = (_this).fm38(_dafny.Seq.UnicodeFromString("rsjphe"), globalState);
          let _lhs117 = globalState;
          let _lhs118 = globalState;
          _lhs117.f2 = _rhs132;
          _lhs118.f15 = _rhs133;
          let _901_v85;
          _901_v85 = new _dafny.CodePoint('w'.codePointAt(0));
          let _902_v86;
          _902_v86 = _dafny.MultiSet.fromElements(_901_v85, _module.__default.fm46(globalState), ((_this).f19)[_module.__default.safeIndex((_817_v27).f32, new BigNumber(((_this).f19).length))], _901_v85);
          _902_v86 = ((_902_v86).Difference(_dafny.MultiSet.FromArray((_this).f19))).Union(_902_v86);
          (globalState).f2 = p0;
        }
        (globalState).f2 = (_818_v28)[_module.__default.safeIndex((_817_v27).f32, new BigNumber((_818_v28).length))];
        if (!((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))])) {
          let _903_v87;
          let _nw143 = new _module.C1();
          _nw143.__ctor((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]);
          _903_v87 = _nw143;
          (globalState).f2 = (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))];
          (globalState).f3 = (_817_v27).f32;
          (_903_v87).m1((_this).f19, _module.__default.safeDivisionInt(_module.__default.fm1((_817_v27).f32, (_this).f18, false, globalState), (_817_v27).f31), _dafny.Seq.UnicodeFromString("hkpnyc"), globalState);
          let _904_v88;
          _904_v88 = _dafny.Set.fromElements((_dafny.ZERO).minus((_817_v27).f31));
          let _905_v89;
          _905_v89 = _dafny.Map.Empty.slice().updateUnsafe(_904_v88,p0);
          _905_v89 = (_905_v89).update(_904_v88, false);
        } else {
          (globalState).f2 = (_this).f18;
          r1 = _module.__default.safeDivisionInt(_module.__default.fm1((_817_v27).f31, (_this).fm3(new BigNumber((_818_v28).length), globalState), (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], globalState), (_817_v27).f31);
          let _906_v90;
          let _nw144 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _906_v90 = _nw144;
          let _index142 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_906_v90).length));
          (_906_v90)[_index142] = new BigNumber(237);
          let _907_v91;
          let _nw145 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
          _907_v91 = _nw145;
          let _908_v92;
          let _nw146 = new _module.C1();
          _nw146.__ctor((_this).f18);
          _908_v92 = _nw146;
          let _909_v93;
          _909_v93 = _dafny.Map.Empty.slice().updateUnsafe(_908_v92,!(p0));
          let _910_v94;
          _910_v94 = _module.D11.create_DC27(_909_v93);
          let _911_v95;
          _911_v95 = _dafny.Set.fromElements(_910_v94);
          let _index143 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_907_v91).length));
          (_907_v91)[_index143] = _911_v95;
          let _index144 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_906_v90).length));
          let _index145 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_907_v91).length));
          let _rhs134 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_817_v27).f32, _816_v26), new BigNumber(((_this).f19).length));
          let _rhs135 = _module.__default.fm1(_816_v26, (_this).fm2(false, (_this).f18, p0, globalState), (_this).f18, globalState);
          let _rhs136 = (_911_v95).Union(_911_v95);
          let _rhs137 = (_this).f18;
          let _lhs119 = globalState;
          let _lhs120 = _906_v90;
          let _lhs121 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_906_v90).length));
          let _lhs122 = _907_v91;
          let _lhs123 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_907_v91).length));
          let _lhs124 = globalState;
          _lhs119.f15 = _rhs134;
          _lhs120[_lhs121] = _rhs135;
          _lhs122[_lhs123] = _rhs136;
          _lhs124.f2 = _rhs137;
          (globalState).f12 = (_this).f19;
          let _912_v96;
          _912_v96 = new _dafny.CodePoint('v'.codePointAt(0));
          _912_v96 = _912_v96;
        }
        let _913_v97;
        let _nw147 = new _module.C2();
        _nw147.__ctor((_817_v27).f32, (_this).f35, p0);
        _913_v97 = _nw147;
      }
      let _914_v98;
      let _nw148 = new _module.C2();
      _nw148.__ctor(new BigNumber(((_this).f19).length), (_this).f35, (_this).f18);
      _914_v98 = _nw148;
      let _915_v99;
      _915_v99 = _module.D13.create_DC34(_815_v25, _dafny.Seq.UnicodeFromString("qgpak"), (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], !((_this).f18));
      let _916_v100;
      _916_v100 = _dafny.Seq.of(_915_v99, _module.D13.create_DC34((_815_v25), _dafny.Seq.UnicodeFromString("eivjeqg"), p0, p0), _915_v99);
      _916_v100 = _dafny.Seq.Concat(_916_v100, _916_v100);
      r0 = (_788_v1).Intersect(_788_v1);
      let _917_v101;
      _917_v101 = _module.D2.create_DC8((_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))], (_815_v25)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_815_v25).length))]);
      let _918_v102;
      _918_v102 = _module.D2.create_DC9(_917_v101);
      let _pat_let_tv15 = _914_v98;
      let _pat_let_tv16 = _817_v27;
      let _pat_let_tv17 = _914_v98;
      let _pat_let_tv18 = _914_v98;
      let _pat_let_tv19 = _817_v27;
      r1 = function (_source8) {
        if (_source8.is_DC7) {
          let _919___mcc_h8 = (_source8).cf15;
          let _920___mcc_h9 = (_source8).cf16;
          let _921___mcc_h10 = (_source8).cf17;
          let _922_cf17 = _921___mcc_h10;
          let _923_cf16 = _920___mcc_h9;
          let _924_cf15 = _919___mcc_h8;
          return _module.__default.safeDivisionInt((_dafny.ZERO).minus(_pat_let_tv15.f33), (_pat_let_tv16).f32);
        } else if (_source8.is_DC8) {
          let _925___mcc_h11 = (_source8).cf18;
          let _926___mcc_h12 = (_source8).cf19;
          let _927___mcc_h13 = (_source8).cf20;
          let _928_cf20 = _927___mcc_h13;
          let _929_cf19 = _926___mcc_h12;
          let _930_cf18 = _925___mcc_h11;
          return _pat_let_tv17.f33;
        } else if (_source8.is_DC6) {
          let _931___mcc_h14 = (_source8).cf14;
          let _932_cf14 = _931___mcc_h14;
          return _pat_let_tv18.f33;
        } else {
          let _933___mcc_h15 = (_source8).cf21;
          let _934_cf21 = _933___mcc_h15;
          return (_pat_let_tv19).f31;
        }
      }(((!(true)) ? (_918_v102) : (_918_v102)));
      r2 = (_dafny.ZERO).minus(_816_v26);
      r3 = _dafny.Seq.Concat((_this).f19, _dafny.Seq.update(_dafny.Seq.Concat((_this).f19, (_this).f19), _module.__default.safeIndex((_817_v27).f32, new BigNumber((_dafny.Seq.Concat((_this).f19, (_this).f19)).length)), new _dafny.CodePoint('j'.codePointAt(0))));
      return [r0, r1, r2, r3];
    }
    get f35() {
      let _this = this;
      return _this._f35;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f18 = false;
      this._f19 = _dafny.Seq.UnicodeFromString("");
      this._f20 = _dafny.Seq.of();
      this._f29 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f30 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f29, f30, f18, f19, f20) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      return;
    }
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return !(((_this).f30).isEqualTo(((_dafny.ZERO).minus((_this).f30)).minus((_this).f30)));
    };
    fm3(p0, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm17(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber(-617)).minus((_dafny.ZERO).minus((_this).f30));
    };
    fm18(globalState) {
      let _this = this;
      return (_this).f30;
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      if ((new BigNumber((_dafny.Seq.UnicodeFromString("pcybxl")).length)).isLessThan((_this).f30)) {
        let _935_v0;
        _935_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f18);
        let _936_v1;
        _936_v1 = _dafny.Set.fromElements(true);
        let _937_v2;
        _937_v2 = _dafny.Set.fromElements(_936_v1, _936_v1, _936_v1, _936_v1);
        let _938_v3;
        _938_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,(p1).multipliedBy(new BigNumber((_937_v2).length)));
        let _rhs138 = (_this).f18;
        let _rhs139 = (_this).f18;
        let _rhs140 = _935_v0;
        let _rhs141 = ((_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f30)).Merge((_938_v3).update(p1, new BigNumber((_935_v0).length)))).Merge(_938_v3);
        let _lhs125 = globalState;
        let _lhs126 = globalState;
        _lhs125.f2 = _rhs138;
        _lhs126.f2 = _rhs139;
        _935_v0 = _rhs140;
        _938_v3 = _rhs141;
        (globalState).f2 = !((_this).f18);
        let _939_v4;
        let _nw149 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _939_v4 = _nw149;
        let _index146 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_939_v4).length));
        (_939_v4)[_index146] = p1;
        let _index147 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_939_v4).length));
        (_939_v4)[_index147] = p1;
        let _940_v5;
        _940_v5 = _dafny.Seq.of((_this).f18, (_this).f18, (_this).f18);
        let _941_v6;
        let _nw150 = new _module.C0();
        _nw150.__ctor(new BigNumber((_940_v5).length), (_this).f30);
        _941_v6 = _nw150;
        let _index148 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_939_v4).length));
        (_939_v4)[_index148] = (_939_v4)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_939_v4).length))];
      } else {
        let _942_v7;
        _942_v7 = _dafny.Seq.of(p1);
        let _943_v8;
        _943_v8 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of((_this).f30), _dafny.Seq.of((_this).f30, (_this).f30)),_dafny.Seq.contains(_942_v7, (_this).f30));
        let _944_v9;
        _944_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f18);
        _943_v8 = (_943_v8).update(_dafny.Seq.of(p1, (_this).f30), !((((_944_v9).contains((_this).f30)) ? ((_944_v9).get((_this).f30)) : ((_this).f18))));
        let _945_v10;
        _945_v10 = _module.D0.create_DC1((_this).f18, p1, _dafny.Seq.of((_this).f29, new _dafny.CodePoint('x'.codePointAt(0))), !((_this).f18), (_this).f30);
        let _946_v11;
        _946_v11 = _dafny.Map.Empty.slice().updateUnsafe(_945_v10,(_this).f18);
        let _947_v12;
        _947_v12 = _dafny.MultiSet.fromElements((_this).f18, true, (_this).f18, (_this).f18);
        _946_v11 = (_946_v11).update(_945_v10, (_947_v12).IsSubsetOf(_947_v12));
        let _948_v13;
        _948_v13 = _module.D2.create_DC7(p1, (_this).f18, (_this).f30);
        if ((_948_v13).dtor_cf16) {
          (globalState).f3 = _module.__default.safeModuloInt(((_this).f30).plus(p1), (((_this).f18) ? ((_this).f30) : (p1)));
          let _949_v14;
          _949_v14 = _dafny.Set.fromElements((_this).f30);
          let _950_v15;
          _950_v15 = _module.D2.create_DC8(!_dafny.Seq.contains(_dafny.Seq.of(p1, new BigNumber((_949_v14).length), new BigNumber(810)), p1), true, _dafny.Seq.IsPrefixOf(p0, p0));
          _950_v15 = _950_v15;
          let _951_v16;
          let _nw151 = Array((new BigNumber(28)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = (_950_v15).dtor_cf19;
          _nw151[(_dafny.ONE).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(2)).toNumber()] = false;
          _nw151[(new BigNumber(3)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(4)).toNumber()] = ((_dafny.ZERO).minus((_dafny.ZERO).minus(p1))).isEqualTo((_this).f30);
          _nw151[(new BigNumber(5)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(6)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(7)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(8)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(9)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(10)).toNumber()] = true;
          _nw151[(new BigNumber(11)).toNumber()] = true;
          _nw151[(new BigNumber(12)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(13)).toNumber()] = !_dafny.Seq.contains(_942_v7, p1);
          _nw151[(new BigNumber(14)).toNumber()] = !(!((_this).f18));
          _nw151[(new BigNumber(15)).toNumber()] = !((_this).f18) || ((_this).f18);
          _nw151[(new BigNumber(16)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(17)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(18)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(19)).toNumber()] = false;
          _nw151[(new BigNumber(20)).toNumber()] = !((_this).f18);
          _nw151[(new BigNumber(21)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(22)).toNumber()] = true;
          _nw151[(new BigNumber(23)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(24)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(25)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(26)).toNumber()] = (_this).f18;
          _nw151[(new BigNumber(27)).toNumber()] = false;
          _951_v16 = _nw151;
          let _index149 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_951_v16).length));
          (_951_v16)[_index149] = !((_this).f18);
          let _index150 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_951_v16).length));
          (_951_v16)[_index150] = true;
          let _952_v17;
          let _init19 = ((_953_v14) => function (_954_i0) {
            return (_954_i0).plus(new BigNumber((_953_v14).length));
          })(_949_v14);
          let _nw152 = Array((new BigNumber(10)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw152.length); _i0_19++) {
            _nw152[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _952_v17 = _nw152;
          let _index151 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_952_v17).length));
          (_952_v17)[_index151] = new BigNumber((p2).length);
          let _index152 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_952_v17).length));
          (_952_v17)[_index152] = (_this).f30;
          let _955_v18;
          _955_v18 = _dafny.Map.Empty.slice().updateUnsafe(_952_v17,(_951_v16)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_951_v16).length))]);
          let _956_v19;
          _956_v19 = _dafny.Seq.of((_this).f18);
          let _957_v20;
          _957_v20 = _module.D1.create_DC3(new BigNumber(309), new BigNumber((_956_v19).length), new BigNumber(-878), p1, (_952_v17)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_952_v17).length))]);
          let _index153 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_951_v16).length));
          let _rhs142 = (((_955_v18).contains(_952_v17)) ? ((_955_v18).get(_952_v17)) : ((_951_v16)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_951_v16).length))]));
          let _rhs143 = (_this).f18;
          let _rhs144 = true;
          let _rhs145 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_957_v20).dtor_cf8));
          let _lhs127 = _951_v16;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_951_v16).length));
          let _lhs129 = globalState;
          let _lhs130 = globalState;
          let _lhs131 = globalState;
          _lhs127[_lhs128] = _rhs142;
          _lhs129.f2 = _rhs143;
          _lhs130.f2 = _rhs144;
          _lhs131.f0 = _rhs145;
        } else {
          let _958_v21;
          _958_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(25),new BigNumber((_947_v12).cardinality()));
          let _959_v22;
          _959_v22 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_942_v7).length),_module.__default.fm1(new BigNumber((_942_v7).length), (_this).f18, (_this).f18, globalState)), _958_v21, _958_v21);
          let _960_v23;
          _960_v23 = _dafny.Map.Empty.slice().updateUnsafe(true,_959_v22);
          _960_v23 = (_960_v23).update((_this).f18, _959_v22);
          (globalState).f15 = (_this).f30;
          let _961_v24;
          let _nw153 = new _module.C1();
          _nw153.__ctor(true);
          _961_v24 = _nw153;
          let _962_v25;
          _962_v25 = _dafny.Seq.of(_947_v12);
          let _963_v27;
          _963_v27 = _dafny.Seq.of(p2, (_this).f19);
          let _964_v28;
          _964_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_961_v24).f18);
          let _rhs146 = _module.__default.fm30(_962_v25, (function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of (_963_v27).Elements) {
              let _965_v26 = _compr_29;
              if (_dafny.Seq.contains(_963_v27, _965_v26)) {
                _coll29.push([_965_v26,(_961_v24).f18]);
              }
            }
            return _coll29;
          }()).Merge(_964_v28), globalState);
          let _rhs147 = _961_v24;
          _958_v21 = _rhs146;
          _961_v24 = _rhs147;
          (globalState).f2 = (_this).f18;
          (globalState).f15 = (_this).f30;
        }
        if ((_this).f18) {
          let _966_v29;
          _966_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
          let _967_v30;
          let _nw154 = Array((new BigNumber(27)).toNumber()).fill(false);
          _967_v30 = _nw154;
          let _968_v31;
          _968_v31 = _module.D2.create_DC9(_module.D2.create_DC8((_this).f18, (_this).f18, (((_944_v9).contains((_this).f30)) ? ((_944_v9).get((_this).f30)) : (true))));
          let _969_v32;
          _969_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f29);
          let _970_v33;
          let _nw155 = Array((new BigNumber(22)).toNumber());
          _nw155[(_dafny.ZERO).toNumber()] = (_this).f29;
          _nw155[(_dafny.ONE).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(2)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(3)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(4)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(5)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(6)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(7)).toNumber()] = (((_969_v32).contains(new BigNumber((p2).length))) ? ((_969_v32).get(new BigNumber((p2).length))) : (new _dafny.CodePoint('q'.codePointAt(0))));
          _nw155[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
          _nw155[(new BigNumber(9)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(10)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(11)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(12)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(13)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(14)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(15)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(16)).toNumber()] = ((_this).f19)[_module.__default.safeIndex(new BigNumber((_942_v7).length), new BigNumber(((_this).f19).length))];
          _nw155[(new BigNumber(17)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(18)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(19)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(20)).toNumber()] = (_this).f29;
          _nw155[(new BigNumber(21)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
          _970_v33 = _nw155;
          let _971_v34;
          _971_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f29);
          let _972_v35;
          _972_v35 = _dafny.Map.Empty.slice().updateUnsafe(_970_v33,_971_v34);
          let _973_v36;
          _973_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_dafny.ZERO).minus(p1));
          let _974_v37;
          _974_v37 = _dafny.MultiSet.fromElements((_this).f30, (_this).f30);
          let _975_v38;
          let _nw156 = Array((new BigNumber(24)).toNumber());
          _nw156[(_dafny.ZERO).toNumber()] = (_this).f30;
          _nw156[(_dafny.ONE).toNumber()] = new BigNumber((_972_v35).length);
          _nw156[(new BigNumber(2)).toNumber()] = (_this).f30;
          _nw156[(new BigNumber(3)).toNumber()] = _module.__default.fm1(p1, true, (_this).f18, globalState);
          _nw156[(new BigNumber(4)).toNumber()] = p1;
          _nw156[(new BigNumber(5)).toNumber()] = (_this).f30;
          _nw156[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_944_v9).length), (_this).f30);
          _nw156[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_this).f30);
          _nw156[(new BigNumber(8)).toNumber()] = (_this).f30;
          _nw156[(new BigNumber(9)).toNumber()] = (new BigNumber(755)).multipliedBy(new BigNumber((p0).length));
          _nw156[(new BigNumber(10)).toNumber()] = ((_this).f30).minus(new BigNumber(21));
          _nw156[(new BigNumber(11)).toNumber()] = new BigNumber(((_this).f19).length);
          _nw156[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_this).f30);
          _nw156[(new BigNumber(13)).toNumber()] = new BigNumber(922);
          _nw156[(new BigNumber(14)).toNumber()] = p1;
          _nw156[(new BigNumber(15)).toNumber()] = p1;
          _nw156[(new BigNumber(16)).toNumber()] = (p1).minus((_this).f30);
          _nw156[(new BigNumber(17)).toNumber()] = (((_973_v36).contains((_this).f18)) ? ((_973_v36).get((_this).f18)) : (new BigNumber(614)));
          _nw156[(new BigNumber(18)).toNumber()] = (_this).f30;
          _nw156[(new BigNumber(19)).toNumber()] = (_this).f30;
          _nw156[(new BigNumber(20)).toNumber()] = new BigNumber((_974_v37).cardinality());
          _nw156[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
          _nw156[(new BigNumber(22)).toNumber()] = (_this).f30;
          _nw156[(new BigNumber(23)).toNumber()] = (_this).f30;
          _975_v38 = _nw156;
          let _976_v39;
          _976_v39 = _dafny.Seq.of(_966_v29, (_966_v29).update((((_944_v9).contains(new BigNumber((_dafny.Seq.of((_this).f18, (_this).f18)).length))) ? ((_944_v9).get(new BigNumber((_dafny.Seq.of((_this).f18, (_this).f18)).length))) : ((_this).f18)), (_this).f18), _dafny.Map.Empty.slice().updateUnsafe((_this).f18,!((_this).f18)));
          let _977_v40;
          _977_v40 = _dafny.Seq.of((_this).f18);
          let _978_v41;
          _978_v41 = _module.D2.create_DC8((_this).fm3(p1, globalState), (_this).f18, (_this).f18);
          let _pat_let_tv20 = _978_v41;
          let _rhs148 = (_976_v39)[_module.__default.safeIndex(p1, new BigNumber((_976_v39).length))];
          let _rhs149 = _967_v30;
          let _rhs150 = (new BigNumber((p0).length)).minus((_this).f30);
          let _rhs151 = (((_977_v40)[_module.__default.safeIndex((_this).f30, new BigNumber((_977_v40).length))]) ? (function (_pat_let28_0) {
            return function (_979_dt__update__tmp_h0) {
              return function (_pat_let29_0) {
                return function (_980_dt__update_hcf21_h0) {
                  return _module.D2.create_DC9(_980_dt__update_hcf21_h0);
                }(_pat_let29_0);
              }(_pat_let_tv20);
            }(_pat_let28_0);
          }(_968_v31)) : (_968_v31));
          let _rhs152 = _975_v38;
          let _lhs132 = globalState;
          _966_v29 = _rhs148;
          _967_v30 = _rhs149;
          _lhs132.f0 = _rhs150;
          _968_v31 = _rhs151;
          _975_v38 = _rhs152;
          let _index154 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_967_v30).length));
          (_967_v30)[_index154] = (new BigNumber(254)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("vd")).length));
          let _index155 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_967_v30).length));
          (_967_v30)[_index155] = !(_module.__default.fm23(_dafny.Set.fromElements(p1), new BigNumber(274), globalState)) || ((_947_v12).IsProperSubsetOf(_947_v12));
          let _981_v42;
          _981_v42 = new _dafny.CodePoint('v'.codePointAt(0));
          _981_v42 = (_this).f29;
          let _index156 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_967_v30).length));
          (_967_v30)[_index156] = (_this).f18;
          let _982_v43;
          _982_v43 = _module.D0.create_DC0((_this).f18);
          let _index157 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_967_v30).length));
          (_967_v30)[_index157] = (_982_v43).dtor_cf0;
        } else {
          let _983_v44;
          _983_v44 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(114))).cardinality()));
          let _984_v45;
          _984_v45 = _dafny.Map.Empty.slice().updateUnsafe(_942_v7,new BigNumber((_983_v44).cardinality()));
          let _985_v46;
          _985_v46 = _module.D1.create_DC4((_this).f18);
          let _986_v47;
          let _nw157 = Array((new BigNumber(23)).toNumber());
          _nw157[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_this).f30, p1);
          _nw157[(_dafny.ONE).toNumber()] = (p1).minus(p1);
          _nw157[(new BigNumber(2)).toNumber()] = (_this).f30;
          _nw157[(new BigNumber(3)).toNumber()] = new BigNumber((_module.__default.fm31(new BigNumber(((_984_v45)).length), (_this).f18, _985_v46, globalState)).cardinality());
          _nw157[(new BigNumber(4)).toNumber()] = (_this).f30;
          _nw157[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt((_this).f30, (_this).f30);
          _nw157[(new BigNumber(6)).toNumber()] = new BigNumber(197);
          _nw157[(new BigNumber(7)).toNumber()] = (_this).f30;
          _nw157[(new BigNumber(8)).toNumber()] = (_this).f30;
          _nw157[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, p1));
          _nw157[(new BigNumber(10)).toNumber()] = (_this).f30;
          _nw157[(new BigNumber(11)).toNumber()] = (_this).f30;
          _nw157[(new BigNumber(12)).toNumber()] = (_this).f30;
          _nw157[(new BigNumber(13)).toNumber()] = new BigNumber((p2).length);
          _nw157[(new BigNumber(14)).toNumber()] = (_this).f30;
          _nw157[(new BigNumber(15)).toNumber()] = new BigNumber((p0).length);
          _nw157[(new BigNumber(16)).toNumber()] = new BigNumber(672);
          _nw157[(new BigNumber(17)).toNumber()] = ((_this).f30).multipliedBy(new BigNumber((_942_v7).length));
          _nw157[(new BigNumber(18)).toNumber()] = p1;
          _nw157[(new BigNumber(19)).toNumber()] = new BigNumber((_942_v7).length);
          _nw157[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex(new BigNumber((_947_v12).cardinality()), new BigNumber((p2).length)), (_this).f29)).length);
          _nw157[(new BigNumber(21)).toNumber()] = p1;
          _nw157[(new BigNumber(22)).toNumber()] = (_this).f30;
          _986_v47 = _nw157;
          let _index158 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_986_v47).length));
          (_986_v47)[_index158] = p1;
          let _987_v48;
          _987_v48 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nla"), (_this).f19), _dafny.Seq.Create(_module.__default.abs(new BigNumber(621)), function (_988_i1) {
            return (_this).f29;
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(228)), function (_989_i2) {
            return (_this).f29;
          }));
          let _index159 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_986_v47).length));
          let _rhs153 = ((_this).f18) || ((_this).f18);
          let _rhs154 = (((_this).f30).minus(p1)).plus((_dafny.ZERO).minus(p1));
          let _rhs155 = new BigNumber((_947_v12).cardinality());
          let _rhs156 = ((((_this).f18) ? (_987_v48) : ((_987_v48).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(77)), function (_990_i3) {
            return (_this).f29;
          }), _module.__default.abs(p1))))).Difference((_987_v48).update(p2, _module.__default.abs((_this).f30)));
          let _lhs133 = globalState;
          let _lhs134 = globalState;
          let _lhs135 = _986_v47;
          let _lhs136 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_986_v47).length));
          _lhs133.f2 = _rhs153;
          _lhs134.f3 = _rhs154;
          _lhs135[_lhs136] = _rhs155;
          _987_v48 = _rhs156;
          let _991_v49;
          let _init20 = ((_992_v47) => function (_993_i4) {
            return _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_992_v47)[_module.__default.safeIndex(new BigNumber(343), new BigNumber((_992_v47).length))]);
          })(_986_v47);
          let _nw158 = Array((new BigNumber(26)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw158.length); _i0_20++) {
            _nw158[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _991_v49 = _nw158;
          let _nw159 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _991_v49 = _nw159;
          let _994_v50;
          _994_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
          _994_v50 = (_994_v50).update((_this).f18, !((_this).f18));
          let _995_v51;
          let _nw160 = Array((new BigNumber(6)).toNumber()).fill(false);
          _995_v51 = _nw160;
          let _996_v52;
          _996_v52 = _986_v47;
          let _997_v53;
          _997_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_996_v52));
          let _index160 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_995_v51).length));
          (_995_v51)[_index160] = !((_this).fm3(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(745),_986_v47)).update((_986_v47)[_module.__default.safeIndex(new BigNumber(343), new BigNumber((_986_v47).length))], (((_997_v53).contains((_this).f18)) ? ((_997_v53).get((_this).f18)) : (_986_v47)))).length), globalState));
          let _index161 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_995_v51).length));
          (_995_v51)[_index161] = (_this).f18;
          let _index162 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_995_v51).length));
          (_995_v51)[_index162] = (_995_v51)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_995_v51).length))];
        }
        let _998_v54;
        let _nw161 = Array((_dafny.ONE).toNumber()).fill(false);
        _998_v54 = _nw161;
        let _index163 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_998_v54).length));
        (_998_v54)[_index163] = (_this).f18;
        let _index164 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_998_v54).length));
        (_998_v54)[_index164] = (_this).fm3(new BigNumber(((_this).f19).length), globalState);
      }
      let _999_v55;
      _999_v55 = _dafny.MultiSet.fromElements((_this).f30);
      let _hi5 = new BigNumber(821);
      for (let _1000_i5 = new BigNumber((_999_v55).cardinality()); _1000_i5.isLessThan(_hi5); _1000_i5 = _1000_i5.plus(_dafny.ONE)) {
        let _1001_v56;
        let _nw162 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _1001_v56 = _nw162;
        let _1002_v57;
        _1002_v57 = _dafny.Seq.of((_this).f18);
        let _1003_v58;
        _1003_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1002_v57);
        let _index165 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length));
        (_1001_v56)[_index165] = _module.__default.safeDivisionInt(new BigNumber(((((_1003_v58).contains((_this).f18)) ? ((_1003_v58).get((_this).f18)) : (_1002_v57))).length), _1000_i5);
        let _1004_v59;
        let _init21 = ((_1005_p1) => function (_1006_i6) {
          return (_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f30)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_1005_p1));
        })(p1);
        let _nw163 = Array((new BigNumber(28)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw163.length); _i0_21++) {
          _nw163[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _1004_v59 = _nw163;
        let _1007_v60;
        _1007_v60 = _dafny.Seq.of(p1);
        let _1008_v61;
        _1008_v61 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18),new BigNumber((_dafny.Seq.update(_1007_v60, _module.__default.safeIndex((_this).f30, new BigNumber((_1007_v60).length)), _1000_i5)).length));
        let _index166 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1004_v59).length));
        (_1004_v59)[_index166] = _1008_v61;
        let _1009_v62;
        _1009_v62 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f18), _module.__default.fm32(globalState));
        let _1010_v64;
        _1010_v64 = _dafny.Seq.of(p2);
        let _1011_v65;
        _1011_v65 = _dafny.MultiSet.fromElements(true);
        let _1012_v66;
        _1012_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1000_i5,(_1007_v60)[_module.__default.safeIndex(new BigNumber((_1011_v65).cardinality()), new BigNumber((_1007_v60).length))]);
        let _index167 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length));
        let _index168 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1004_v59).length));
        let _rhs157 = new BigNumber(((_module.__default.fm30(_1009_v62, function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of (_1010_v64).Elements) {
            let _1013_v63 = _compr_30;
            if (_dafny.Seq.contains(_1010_v64, _1013_v63)) {
              _coll30.push([_1013_v63,(_this).f18]);
            }
          }
          return _coll30;
        }(), globalState)).Merge((((_this).f18) ? (_1012_v66) : (_1012_v66)))).length);
        let _rhs158 = (_dafny.ZERO).minus((_1000_i5).minus(_1000_i5));
        let _rhs159 = _1000_i5;
        let _rhs160 = _1008_v61;
        let _lhs137 = globalState;
        let _lhs138 = _1001_v56;
        let _lhs139 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length));
        let _lhs140 = globalState;
        let _lhs141 = _1004_v59;
        let _lhs142 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1004_v59).length));
        _lhs137.f0 = _rhs157;
        _lhs138[_lhs139] = _rhs158;
        _lhs140.f15 = _rhs159;
        _lhs141[_lhs142] = _rhs160;
        (globalState).f2 = (_this).fm2((_this).f18, false, (_this).f18, globalState);
        if ((_this).f18) {
          let _1014_v67;
          let _nw164 = new _module.C2();
          _nw164.__ctor(_module.__default.fm1((_1001_v56)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length))], (_this).f18, (_this).f18, globalState), _dafny.Seq.of(new BigNumber((p2).length)), !((_1002_v57)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of((_this).f30, (_this).f30)).length), (_this).f30)).cardinality()), new BigNumber((_1002_v57).length))]));
          _1014_v67 = _nw164;
          let _1015_v68;
          _1015_v68 = _module.D1.create_DC2(_1014_v67);
          let _1016_v69;
          _1016_v69 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(_1015_v68),p1);
          let _1017_v70;
          _1017_v70 = _module.D1.create_DC5(_1015_v68);
          _1016_v69 = (_1016_v69).update(_1017_v70, p1);
          let _1018_v71;
          let _init22 = ((_1019_v61) => function (_1020_i7) {
            return _module.D2.create_DC6(_1019_v61);
          })(_1008_v61);
          let _nw165 = Array((new BigNumber(24)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw165.length); _i0_22++) {
            _nw165[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _1018_v71 = _nw165;
          let _1021_v72;
          let _nw166 = Array((new BigNumber(2)).toNumber());
          _nw166[(_dafny.ZERO).toNumber()] = _1018_v71;
          _nw166[(_dafny.ONE).toNumber()] = _1018_v71;
          _1021_v72 = _nw166;
          let _rhs161 = p1;
          let _rhs162 = _1021_v72;
          let _rhs163 = p1;
          let _rhs164 = (_this).f18;
          let _lhs143 = globalState;
          let _lhs144 = globalState;
          let _lhs145 = globalState;
          _lhs143.f3 = _rhs161;
          _1021_v72 = _rhs162;
          _lhs144.f3 = _rhs163;
          _lhs145.f2 = _rhs164;
          let _1022_v73;
          _1022_v73 = _dafny.MultiSet.fromElements((_1012_v66).update((_1001_v56)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length))], new BigNumber(695)), (_1012_v66).update(_1000_i5, (_this).f30));
          _1012_v66 = (_1012_v66).update((_this).f30, new BigNumber((_1022_v73).cardinality()));
          let _index169 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length));
          (_1001_v56)[_index169] = _module.__default.safeDivisionInt((_1001_v56)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length))], p1);
          _1001_v56 = _1001_v56;
        } else {
          let _1023_v74;
          _1023_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(827),(_999_v55).IsSubsetOf(_999_v55));
          _1023_v74 = (_1023_v74).update((_1001_v56)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length))], (_this).f18);
          (globalState).f2 = (_this).f18;
          let _index170 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length));
          (_1001_v56)[_index170] = ((_1000_i5).minus(p1)).minus(new BigNumber((p2).length));
          (globalState).f3 = ((_this).f30).plus(_1000_i5);
          (globalState).f2 = (_this).f18;
        }
        let _1024_v75;
        _1024_v75 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f18);
        let _1025_v76;
        let _nw167 = new _module.C0();
        _nw167.__ctor(p1, _module.__default.safeModuloInt(new BigNumber(((_module.D0.create_DC1((_this).fm3((_1001_v56)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length))], globalState), _1000_i5, p2, (_this).f18, new BigNumber((_1024_v75).length))).dtor_cf3).length), (_1001_v56)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_1001_v56).length))]));
        _1025_v76 = _nw167;
      }
      (globalState).f2 = (_this).f18;
      let _1026_v77;
      _1026_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).fm3((_this).f30, globalState));
      let _1027_i8;
      _1027_i8 = _dafny.ZERO;
      L8: {
        while ((((_1026_v77).contains((_this).f30)) ? ((_1026_v77).get((_this).f30)) : ((_this).f18))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1027_i8)) {
              break L8;
            }
            _1027_i8 = (_1027_i8).plus(_dafny.ONE);
            (globalState).f2 = (_this).f18;
            if (!((_this).f18)) {
              let _1028_v78;
              _1028_v78 = _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(161), p1)));
              let _1029_v79;
              _1029_v79 = _dafny.Seq.of((_this).f18, (_this).f18, (_this).f18, (_this).f18);
              let _rhs165 = _dafny.Seq.update(_1028_v78, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat((_this).f19, _dafny.Seq.UnicodeFromString("tdxtowp"))).length), new BigNumber((_1028_v78).length)), p1);
              let _rhs166 = ((_dafny.ZERO).minus(p1)).isLessThanOrEqualTo(_module.__default.fm1((_this).f30, (_this).f18, true, globalState));
              let _rhs167 = (_dafny.ZERO).minus((new BigNumber(972)).multipliedBy((_this).f30));
              let _rhs168 = ((_this).f18) === ((_this).f18);
              let _rhs169 = (new BigNumber((_dafny.Seq.update(_1029_v79, _module.__default.safeIndex(p1, new BigNumber((_1029_v79).length)), (_this).f18)).length)).multipliedBy(p1);
              let _lhs146 = globalState;
              let _lhs147 = globalState;
              let _lhs148 = globalState;
              let _lhs149 = globalState;
              let _lhs150 = globalState;
              _lhs146.f10 = _rhs165;
              _lhs147.f2 = _rhs166;
              _lhs148.f15 = _rhs167;
              _lhs149.f2 = _rhs168;
              _lhs150.f0 = _rhs169;
              let _1030_v80;
              let _nw168 = new _module.C2();
              _nw168.__ctor((_this).f30, _module.__default.fm33((_this).f18, globalState), !(new BigNumber(67)).isEqualTo(new BigNumber(246)));
              _1030_v80 = _nw168;
              let _1031_v81;
              let _init23 = function (_1032_i9) {
                return (_this).f18;
              };
              let _nw169 = Array((new BigNumber(9)).toNumber());
              for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw169.length); _i0_23++) {
                _nw169[_i0_23] = _init23(new BigNumber(_i0_23));
              }
              _1031_v81 = _nw169;
              let _1033_v82;
              _1033_v82 = _dafny.MultiSet.fromElements(_1031_v81, _1031_v81, _1031_v81, _1031_v81);
              _1033_v82 = (((_this).f18) ? (_1033_v82) : (_1033_v82));
              let _1034_v84;
              _1034_v84 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f18),(_this).f30);
              let _1035_v85;
              _1035_v85 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f18),(_this).fm3(p1, globalState));
              let _1036_v86;
              let _nw170 = new _module.C1();
              _nw170.__ctor(!(function () {
                let _coll31 = new _dafny.Map();
                for (const _compr_31 of (_1034_v84).Keys.Elements) {
                  let _1037_v83 = _compr_31;
                  if ((_1034_v84).contains(_1037_v83)) {
                    _coll31.push([_1037_v83,(_this).f18]);
                  }
                }
                return _coll31;
              }()).equals(_1035_v85));
              _1036_v86 = _nw170;
              let _1038_v87;
              _1038_v87 = _module.D2.create_DC8((_this).f18, (_this).f18, !((_this).f18));
              let _1039_v88;
              let _nw171 = new _module.C1();
              _nw171.__ctor(((_1038_v87).dtor_cf20) || (false));
              _1039_v88 = _nw171;
            } else {
              let _1040_v89;
              _1040_v89 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f18, (_this).f18)).length));
              (globalState).f0 = _module.__default.safeModuloInt((_this).f30, ((_dafny.ZERO).minus(new BigNumber((_1040_v89).length))).multipliedBy((_this).f30));
              let _1041_v90;
              let _init24 = function (_1042_i10) {
                return (_this).f18;
              };
              let _nw172 = Array((new BigNumber(15)).toNumber());
              for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw172.length); _i0_24++) {
                _nw172[_i0_24] = _init24(new BigNumber(_i0_24));
              }
              _1041_v90 = _nw172;
              let _1043_v91;
              _1043_v91 = _dafny.Seq.of(true, (_this).f18);
              let _index171 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1041_v90).length));
              (_1041_v90)[_index171] = (_1043_v91)[_module.__default.safeIndex(p1, new BigNumber((_1043_v91).length))];
              let _index172 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1041_v90).length));
              let _rhs170 = _dafny.Seq.UnicodeFromString("vqgd");
              let _rhs171 = (_this).f18;
              let _rhs172 = p1;
              let _rhs173 = !(true) || (!((_this).f18));
              let _rhs174 = (_this).fm2((_this).f18, (_this).f18, (_this).f18, globalState);
              let _lhs151 = globalState;
              let _lhs152 = globalState;
              let _lhs153 = globalState;
              let _lhs154 = globalState;
              let _lhs155 = _1041_v90;
              let _lhs156 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1041_v90).length));
              _lhs151.f12 = _rhs170;
              _lhs152.f2 = _rhs171;
              _lhs153.f15 = _rhs172;
              _lhs154.f2 = _rhs173;
              _lhs155[_lhs156] = _rhs174;
              (globalState).f2 = (_1041_v90)[_module.__default.safeIndex(new BigNumber(31), new BigNumber((_1041_v90).length))];
              let _1044_v92;
              let _1045_v93;
              let _1046_v94;
              let _1047_v95;
              let _out32;
              let _out33;
              let _out34;
              let _out35;
              let _outcollector8 = _module.__default.m0(globalState);
              _out32 = _outcollector8[0];
              _out33 = _outcollector8[1];
              _out34 = _outcollector8[2];
              _out35 = _outcollector8[3];
              _1044_v92 = _out32;
              _1045_v93 = _out33;
              _1046_v94 = _out34;
              _1047_v95 = _out35;
              (globalState).f2 = (_this).f18;
            }
            (globalState).f15 = p1;
            _1026_v77 = (_1026_v77).update(((_this).f30).minus(new BigNumber(403)), (_this).f18);
          }
        }
      }
      let _1048_v96;
      _1048_v96 = _dafny.Set.fromElements(p1);
      let _1049_v97;
      _1049_v97 = _dafny.Seq.of((_this).f18, (_this).f18);
      let _1050_v98;
      _1050_v98 = _dafny.Seq.of(new BigNumber((_1048_v96).length), new BigNumber((_1049_v97).length), p1, (_this).f30, p1);
      let _hi6 = new BigNumber((_1050_v98).length);
      for (let _1051_i11 = new BigNumber(528); _1051_i11.isLessThan(_hi6); _1051_i11 = _1051_i11.plus(_dafny.ONE)) {
        let _1052_v99;
        let _nw173 = Array((new BigNumber(4)).toNumber());
        _nw173[(_dafny.ZERO).toNumber()] = (_this).f18;
        _nw173[(_dafny.ONE).toNumber()] = (_this).f18;
        _nw173[(new BigNumber(2)).toNumber()] = (_this).f18;
        _nw173[(new BigNumber(3)).toNumber()] = (_1049_v97)[_module.__default.safeIndex(_1051_i11, new BigNumber((_1049_v97).length))];
        _1052_v99 = _nw173;
        let _1053_v100;
        _1053_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1049_v97, _1049_v97)).length),_1052_v99);
        _1053_v100 = (_1053_v100).update(new BigNumber(20), _1052_v99);
        let _1054_v101;
        _1054_v101 = _dafny.MultiSet.fromElements((_this).f18);
        let _1055_v102;
        _1055_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p1);
        let _1056_v104;
        _1056_v104 = _dafny.Seq.of(_1048_v96);
        let _1057_v105;
        let _nw174 = Array((new BigNumber(29)).toNumber());
        _nw174[(_dafny.ZERO).toNumber()] = (_this).f30;
        _nw174[(_dafny.ONE).toNumber()] = p1;
        _nw174[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("btuvfidmd")).length);
        _nw174[(new BigNumber(3)).toNumber()] = new BigNumber(-309);
        _nw174[(new BigNumber(4)).toNumber()] = (_this).fm17((_dafny.ZERO).minus((_this).f30), _1054_v101, globalState);
        _nw174[(new BigNumber(5)).toNumber()] = p1;
        _nw174[(new BigNumber(6)).toNumber()] = new BigNumber(725);
        _nw174[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw174[(new BigNumber(8)).toNumber()] = p1;
        _nw174[(new BigNumber(9)).toNumber()] = _1051_i11;
        _nw174[(new BigNumber(10)).toNumber()] = p1;
        _nw174[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex(_1051_i11, new BigNumber((p2).length)), (_this).f29)).length);
        _nw174[(new BigNumber(12)).toNumber()] = p1;
        _nw174[(new BigNumber(13)).toNumber()] = p1;
        _nw174[(new BigNumber(14)).toNumber()] = new BigNumber((_1055_v102).length);
        _nw174[(new BigNumber(15)).toNumber()] = (_this).f30;
        _nw174[(new BigNumber(16)).toNumber()] = (_this).f30;
        _nw174[(new BigNumber(17)).toNumber()] = _1051_i11;
        _nw174[(new BigNumber(18)).toNumber()] = (_this).f30;
        _nw174[(new BigNumber(19)).toNumber()] = (_this).fm18(globalState);
        _nw174[(new BigNumber(20)).toNumber()] = _1051_i11;
        _nw174[(new BigNumber(21)).toNumber()] = new BigNumber((function () {
          let _coll32 = new _dafny.Set();
          for (const _compr_32 of _dafny.IntegerRange(new BigNumber(711), new BigNumber(606))) {
            let _1058_v103 = _compr_32;
            if (((new BigNumber(711)).isLessThanOrEqualTo(_1058_v103)) && ((_1058_v103).isLessThan(new BigNumber(606)))) {
              _coll32.add((_1058_v103).minus(_1051_i11));
            }
          }
          return _coll32;
        }()).length);
        _nw174[(new BigNumber(22)).toNumber()] = p1;
        _nw174[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus((_this).f30);
        _nw174[(new BigNumber(24)).toNumber()] = (_this).f30;
        _nw174[(new BigNumber(25)).toNumber()] = p1;
        _nw174[(new BigNumber(26)).toNumber()] = new BigNumber((_1056_v104).length);
        _nw174[(new BigNumber(27)).toNumber()] = p1;
        _nw174[(new BigNumber(28)).toNumber()] = p1;
        _1057_v105 = _nw174;
        let _1059_v106;
        _1059_v106 = _dafny.Map.Empty.slice().updateUnsafe((_1049_v97)[_module.__default.safeIndex(new BigNumber(932), new BigNumber((_1049_v97).length))],_1057_v105);
        _1059_v106 = (_1059_v106).update((_this).f18, _1057_v105);
        (globalState).f15 = _1051_i11;
        let _1060_v107;
        _1060_v107 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_1049_v97)[_module.__default.safeIndex(_1051_i11, new BigNumber((_1049_v97).length))]);
        _1060_v107 = (_1060_v107).update(false, (_this).f18);
      }
      let _1061_v108;
      _1061_v108 = _module.D2.create_DC8((_this).f18, (_this).f18, (_this).f18);
      let _1062_v109;
      _1062_v109 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1049_v97).length),_1061_v108);
      _1062_v109 = (_module.__default.fm34((_this).f30, (_this).f30, globalState)).Merge(_1062_v109);
      return;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1063_v0;
      _1063_v0 = _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f30));
      let _source9 = _1063_v0;
      if (_source9.is_DC7) {
        let _1064___mcc_h0 = (_source9).cf15;
        let _1065___mcc_h1 = (_source9).cf16;
        let _1066___mcc_h2 = (_source9).cf17;
        let _1067_cf17 = _1066___mcc_h2;
        let _1068_cf16 = _1065___mcc_h1;
        let _1069_cf15 = _1064___mcc_h0;
        let _1070_v1;
        _1070_v1 = _module.D3.create_DC14(_module.D3.create_DC11(p1, p2));
        let _source10 = _1070_v1;
        if (_source10.is_DC11) {
          let _1071___mcc_h8 = (_source10).cf23;
          let _1072___mcc_h9 = (_source10).cf24;
          let _1073_cf24 = _1072___mcc_h9;
          let _1074_cf23 = _1071___mcc_h8;
          r0 = _1069_cf15;
          let _1075_v2;
          let _nw175 = Array((new BigNumber(25)).toNumber()).fill(false);
          _1075_v2 = _nw175;
          let _index173 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1075_v2).length));
          (_1075_v2)[_index173] = (_this).f18;
          let _index174 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1075_v2).length));
          (_1075_v2)[_index174] = (_this).f18;
          let _1076_v3;
          _1076_v3 = _dafny.Set.fromElements((_this).f30);
          (globalState).f15 = _module.__default.fm1((new BigNumber((_1076_v3).length)).minus(_1067_cf17), (_module.__default.fm35(_1067_cf17, _1073_cf24, globalState)).dtor_cf4, _1073_cf24, globalState);
          (globalState).f0 = _module.__default.safeDivisionInt(_1069_cf15, new BigNumber((_module.__default.fm32(globalState)).cardinality()));
        } else if (_source10.is_DC12) {
          let _1077___mcc_h10 = (_source10).cf25;
          let _1078___mcc_h11 = (_source10).cf26;
          let _1079___mcc_h12 = (_source10).cf27;
          let _1080___mcc_h13 = (_source10).cf28;
          let _1081___mcc_h14 = (_source10).cf29;
          let _1082_cf29 = _1081___mcc_h14;
          let _1083_cf28 = _1080___mcc_h13;
          let _1084_cf27 = _1079___mcc_h12;
          let _1085_cf26 = _1078___mcc_h11;
          let _1086_cf25 = _1077___mcc_h10;
          let _1087_v4;
          _1087_v4 = _dafny.Seq.of((_this).f18);
          let _1088_v5;
          _1088_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1067_cf17,new BigNumber((_1087_v4).length));
          let _1089_v6;
          _1089_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1088_v5,new BigNumber(505));
          let _1090_v7;
          _1090_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1067_cf17,_1088_v5);
          let _1091_v8;
          _1091_v8 = _dafny.Set.fromElements((_this).f30, (_this).f30);
          let _1092_v9;
          _1092_v9 = _dafny.Seq.of(new BigNumber(72), new BigNumber((_1091_v8).length));
          let _1093_v10;
          _1093_v10 = _dafny.Set.fromElements(p1);
          let _1094_v11;
          _1094_v11 = _dafny.Seq.of(_1093_v10, _module.__default.fm36((_this).f18, _1084_cf27, (_this).f18, globalState));
          _1089_v6 = (_1089_v6).update((((_1090_v7).contains(_1067_cf17)) ? ((_1090_v7).get(_1067_cf17)) : (_dafny.Map.Empty.slice().updateUnsafe(_1069_cf15,new BigNumber((_1092_v9).length)))), new BigNumber((_1094_v11).length));
          let _1095_v12;
          _1095_v12 = _dafny.Seq.of(_dafny.Seq.Concat(_1092_v9, _1092_v9));
          _1095_v12 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(88)), function (_1096_i0) {
            return _dafny.Seq.of(new BigNumber(590));
          });
          let _1097_v13;
          _1097_v13 = new _dafny.CodePoint('v'.codePointAt(0));
          _1097_v13 = ((_this).f19)[_module.__default.safeIndex(_1069_cf15, new BigNumber(((_this).f19).length))];
          let _1098_v14;
          _1098_v14 = _module.D2.create_DC7(new BigNumber(485), _1085_cf26, new BigNumber(((_this).f19).length));
          _1068_cf16 = (_module.__default.fm1(_1067_cf17, _1082_cf29, (_1098_v14).dtor_cf16, globalState)).isLessThanOrEqualTo(_1084_cf27);
        } else if (_source10.is_DC13) {
          let _1099___mcc_h15 = (_source10).cf30;
          let _1100___mcc_h16 = (_source10).cf31;
          let _1101___mcc_h17 = (_source10).cf32;
          let _1102___mcc_h18 = (_source10).cf33;
          let _1103___mcc_h19 = (_source10).cf34;
          let _1104_cf34 = _1103___mcc_h19;
          let _1105_cf33 = _1102___mcc_h18;
          let _1106_cf32 = _1101___mcc_h17;
          let _1107_cf31 = _1100___mcc_h16;
          let _1108_cf30 = _1099___mcc_h15;
          let _1109_v15;
          _1109_v15 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _1110_v16;
          let _nw176 = Array((new BigNumber(3)).toNumber());
          _nw176[(_dafny.ZERO).toNumber()] = _1109_v15;
          _nw176[(_dafny.ONE).toNumber()] = _1109_v15;
          _nw176[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
          _1110_v16 = _nw176;
          let _index175 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_1110_v16).length));
          (_1110_v16)[_index175] = _1109_v15;
          let _index176 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_1110_v16).length));
          (_1110_v16)[_index176] = ((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,p2)).Merge(_1109_v15)).Merge(_1109_v15);
          let _1111_v17;
          _1111_v17 = _dafny.MultiSet.fromElements(p2);
          let _1112_v18;
          let _nw177 = new _module.C0();
          _nw177.__ctor((new BigNumber(((_this).f19).length)).multipliedBy(_1108_cf30), (((_1111_v17).contains(p2)) ? ((_1111_v17).get(p2)) : (_1105_cf33)));
          _1112_v18 = _nw177;
          let _1113_v19;
          let _nw178 = Array((new BigNumber(13)).toNumber()).fill([]);
          _1113_v19 = _nw178;
          let _1114_v20;
          let _nw179 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1114_v20 = _nw179;
          let _index177 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1113_v19).length));
          (_1113_v19)[_index177] = _1114_v20;
          let _index178 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1113_v19).length));
          let _rhs175 = !((p2) && ((_this).f18));
          let _rhs176 = _1114_v20;
          let _rhs177 = p2;
          let _lhs157 = _1113_v19;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1113_v19).length));
          let _lhs159 = globalState;
          _1068_cf16 = _rhs175;
          _lhs157[_lhs158] = _rhs176;
          _lhs159.f2 = _rhs177;
          let _1115_v21;
          _1115_v21 = _dafny.Seq.of(_1108_cf30);
          (globalState).f3 = new BigNumber((_1115_v21).length);
        } else if (_source10.is_DC10) {
          let _1116___mcc_h20 = (_source10).cf22;
          let _1117_cf22 = _1116___mcc_h20;
          let _1118_v22;
          let _init25 = ((_1119_cf15) => function (_1120_i1) {
            return ((_this).f30).isLessThanOrEqualTo(_1119_cf15);
          })(_1069_cf15);
          let _nw180 = Array((new BigNumber(22)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw180.length); _i0_25++) {
            _nw180[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1118_v22 = _nw180;
          let _index179 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_1118_v22).length));
          (_1118_v22)[_index179] = p1;
          let _1121_v23;
          _1121_v23 = _dafny.Seq.of(_1067_cf17);
          let _1122_v24;
          let _nw181 = new _module.C3();
          _nw181.__ctor(_dafny.Seq.Concat(_1121_v23, _1121_v23), (_this).f19, (_this).f20, _1068_cf16);
          _1122_v24 = _nw181;
          let _index180 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_1118_v22).length));
          let _rhs178 = p2;
          let _rhs179 = (_this).f19;
          let _rhs180 = !_dafny.Seq.contains(_1121_v23, (_this).f30);
          let _rhs181 = _1122_v24;
          let _lhs160 = _1118_v22;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_1118_v22).length));
          let _lhs162 = globalState;
          let _lhs163 = globalState;
          _lhs160[_lhs161] = _rhs178;
          _lhs162.f12 = _rhs179;
          _lhs163.f2 = _rhs180;
          _1122_v24 = _rhs181;
          let _1123_v25;
          _1123_v25 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f18) === (_1068_cf16),(_1069_cf15).isEqualTo(_1067_cf17));
          let _1124_v26;
          _1124_v26 = _dafny.MultiSet.fromElements((_1121_v23)[_module.__default.safeIndex(_1069_cf15, new BigNumber((_1121_v23).length))]);
          let _1125_v27;
          _1125_v27 = _dafny.Set.fromElements(_1068_cf16, ((_1124_v26).update(_1069_cf15, _module.__default.abs(_module.__default.fm1(new BigNumber(444), (_this).f18, _1068_cf16, globalState)))).IsProperSubsetOf(_1124_v26), p2, false);
          let _1126_v28;
          _1126_v28 = _dafny.MultiSet.fromElements(!((_this).fm3((_this).f30, globalState)));
          let _1127_v29;
          _1127_v29 = _dafny.Seq.of((_this).f18, p1);
          let _1128_v30;
          _1128_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1127_v29);
          let _1129_v31;
          _1129_v31 = _1128_v30;
          let _rhs182 = true;
          let _rhs183 = (_module.__default.fm47(_1126_v28, _1067_cf17, new BigNumber(((_1129_v31)).length), (_1118_v22)[_module.__default.safeIndex(new BigNumber(415), new BigNumber((_1118_v22).length))], globalState)).update((_1118_v22)[_module.__default.safeIndex(new BigNumber(415), new BigNumber((_1118_v22).length))], (_this).f18);
          let _rhs184 = _1125_v27;
          let _lhs164 = globalState;
          _lhs164.f2 = _rhs182;
          _1123_v25 = _rhs183;
          _1125_v27 = _rhs184;
          let _1130_v32;
          let _nw182 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1130_v32 = _nw182;
          let _index181 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_1130_v32).length));
          (_1130_v32)[_index181] = (_this).f30;
          let _index182 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_1130_v32).length));
          (_1130_v32)[_index182] = (_1069_cf15).plus(_1067_cf17);
          let _1131_v33;
          _1131_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_1130_v32);
          let _1132_v34;
          _1132_v34 = _module.D2.create_DC7(_1067_cf17, false, new BigNumber((_1131_v33).length));
          let _1133_v35;
          let _nw183 = new _module.C2();
          _nw183.__ctor((_1069_cf15).multipliedBy((_1130_v32)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_1130_v32).length))]), _dafny.Seq.Concat(_1121_v23, _1121_v23), (_1132_v34).dtor_cf16);
          _1133_v35 = _nw183;
        } else {
          let _1134___mcc_h21 = (_source10).cf35;
          let _1135_cf35 = _1134___mcc_h21;
          let _1136_v36;
          _1136_v36 = _dafny.Seq.of(false, _1068_cf16, (_this).f18);
          let _1137_v37;
          _1137_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1068_cf16,_dafny.Seq.update(_1136_v36, _module.__default.safeIndex(_1069_cf15, new BigNumber((_1136_v36).length)), true));
          _1137_v37 = (_1137_v37).update(p1, _dafny.Seq.Concat(_1136_v36, _dafny.Seq.of(p1)));
          let _1138_v38;
          let _nw184 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Set.Empty);
          _1138_v38 = _nw184;
          let _1139_v39;
          _1139_v39 = _dafny.Set.fromElements((_this).f30, _1069_cf15);
          let _index183 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_1138_v38).length));
          (_1138_v38)[_index183] = _1139_v39;
          let _index184 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_1138_v38).length));
          (_1138_v38)[_index184] = _1139_v39;
          let _1140_v40;
          _1140_v40 = _dafny.Seq.of(_1067_cf17, (_this).f30);
          let _1141_v41;
          _1141_v41 = _dafny.Seq.of(new BigNumber((_1140_v40).length), _1067_cf17);
          let _1142_v42;
          _1142_v42 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),(_this).fm3(_1067_cf17, globalState));
          (globalState).f3 = _module.__default.fm1(((_this).f30).plus(_module.__default.fm1(new BigNumber((_1141_v41).length), _1068_cf16, p1, globalState)), _1068_cf16, (((_1142_v42).contains(((_this).f19)[_module.__default.safeIndex((_dafny.ZERO).minus(_1067_cf17), new BigNumber(((_this).f19).length))])) ? ((_1142_v42).get(((_this).f19)[_module.__default.safeIndex((_dafny.ZERO).minus(_1067_cf17), new BigNumber(((_this).f19).length))])) : (p1)), globalState);
          let _1143_v43;
          let _nw185 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1143_v43 = _nw185;
          let _index185 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_1143_v43).length));
          (_1143_v43)[_index185] = _dafny.areEqual((_this).f19, _module.__default.fm48(globalState));
          let _index186 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_1143_v43).length));
          (_1143_v43)[_index186] = (_this).fm3(new BigNumber(((_module.__default.fm49(new BigNumber(708), (_this).f30, _1069_cf15, globalState)).Intersect(_dafny.Set.fromElements(new BigNumber(414), _1067_cf17, _1067_cf17))).length), globalState);
        }
        let _1144_v44;
        let _nw186 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
        _1144_v44 = _nw186;
        let _1145_v45;
        _1145_v45 = _module.D3.create_DC10(_1144_v44);
        let _pat_let_tv21 = _1144_v44;
        let _source11 = function (_pat_let30_0) {
          return function (_1146_dt__update__tmp_h0) {
            return function (_pat_let31_0) {
              return function (_1147_dt__update_hcf22_h0) {
                return _module.D3.create_DC10(_1147_dt__update_hcf22_h0);
              }(_pat_let31_0);
            }(_pat_let_tv21);
          }(_pat_let30_0);
        }(_1145_v45);
        if (_source11.is_DC11) {
          let _1148___mcc_h22 = (_source11).cf23;
          let _1149___mcc_h23 = (_source11).cf24;
          let _1150_cf24 = _1149___mcc_h23;
          let _1151_cf23 = _1148___mcc_h22;
          let _1152_v46;
          _1152_v46 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1150_cf24);
          _1068_cf16 = (_1150_cf24) && ((((_1152_v46).contains(_1151_cf23)) ? ((_1152_v46).get(_1151_cf23)) : ((_this).f18)));
          (globalState).f2 = p2;
          let _1153_v47;
          let _nw187 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _1153_v47 = _nw187;
          let _1154_v48;
          _1154_v48 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_this).f29)).length));
          let _1155_v50;
          _1155_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1153_v47,function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of (_1154_v48).Elements) {
              let _1156_v49 = _compr_33;
              if ((_1154_v48).contains(_1156_v49)) {
                _coll33.add(_module.__default.safeModuloInt(_1156_v49, new BigNumber((_dafny.Seq.of(false)).length)));
              }
            }
            return _coll33;
          }());
          _1155_v50 = (_1155_v50).Merge(_1155_v50);
          let _1157_v51;
          let _nw188 = new _module.C1();
          _nw188.__ctor(((_dafny.ZERO).minus(_1069_cf15)).isLessThan(new BigNumber(((_this).f19).length)));
          _1157_v51 = _nw188;
        } else if (_source11.is_DC12) {
          let _1158___mcc_h24 = (_source11).cf25;
          let _1159___mcc_h25 = (_source11).cf26;
          let _1160___mcc_h26 = (_source11).cf27;
          let _1161___mcc_h27 = (_source11).cf28;
          let _1162___mcc_h28 = (_source11).cf29;
          let _1163_cf29 = _1162___mcc_h28;
          let _1164_cf28 = _1161___mcc_h27;
          let _1165_cf27 = _1160___mcc_h26;
          let _1166_cf26 = _1159___mcc_h25;
          let _1167_cf25 = _1158___mcc_h24;
          let _1168_v52;
          _1168_v52 = _module.D2.create_DC8((_this).f18, false, !(true));
          let _1169_v53;
          _1169_v53 = _module.D2.create_DC9(_1168_v52);
          let _1170_v54;
          _1170_v54 = _dafny.MultiSet.fromElements(_1163_cf29);
          let _1171_v55;
          _1171_v55 = _dafny.MultiSet.fromElements(_1069_cf15, new BigNumber((_1170_v54).cardinality()));
          let _1172_v56;
          _1172_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,new BigNumber((_1170_v54).cardinality()));
          let _1173_v57;
          _1173_v57 = _dafny.Seq.of(_1171_v55);
          let _1174_v58;
          _1174_v58 = _dafny.Map.Empty.slice().updateUnsafe(!(_dafny.Map.Empty.slice().updateUnsafe(_1067_cf17,(_this).fm18(globalState))).equals(_1172_v56),(_1173_v57)[_module.__default.safeIndex(_1067_cf17, new BigNumber((_1173_v57).length))]);
          let _rhs185 = _1169_v53;
          let _rhs186 = _1163_cf29;
          let _rhs187 = (((_1174_v58).contains(_1166_cf26)) ? ((_1174_v58).get(_1166_cf26)) : (_1171_v55));
          let _lhs165 = globalState;
          _1169_v53 = _rhs185;
          _lhs165.f2 = _rhs186;
          _1171_v55 = _rhs187;
          let _1175_v59;
          _1175_v59 = _dafny.Set.fromElements(true);
          let _1176_v60;
          _1176_v60 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC8(_1068_cf16, _1068_cf16, false),p2);
          let _1177_v61;
          _1177_v61 = new _dafny.CodePoint('c'.codePointAt(0));
          let _rhs188 = _1175_v59;
          let _rhs189 = _1176_v60;
          let _rhs190 = (_1165_cf27).isEqualTo(_1069_cf15);
          let _rhs191 = ((p1) ? ((_this).f29) : ((((_this).f18) ? (_1177_v61) : ((_this).f29))));
          let _rhs192 = (_dafny.ZERO).minus(_1069_cf15);
          let _lhs166 = globalState;
          _1175_v59 = _rhs188;
          _1176_v60 = _rhs189;
          _1163_cf29 = _rhs190;
          _1177_v61 = _rhs191;
          _lhs166.f3 = _rhs192;
          let _1178_v62;
          let _nw189 = new _module.C0();
          _nw189.__ctor((_this).f30, (_this).f30);
          _1178_v62 = _nw189;
          (globalState).f3 = (_dafny.ZERO).minus(_1067_cf17);
        } else if (_source11.is_DC13) {
          let _1179___mcc_h29 = (_source11).cf30;
          let _1180___mcc_h30 = (_source11).cf31;
          let _1181___mcc_h31 = (_source11).cf32;
          let _1182___mcc_h32 = (_source11).cf33;
          let _1183___mcc_h33 = (_source11).cf34;
          let _1184_cf34 = _1183___mcc_h33;
          let _1185_cf33 = _1182___mcc_h32;
          let _1186_cf32 = _1181___mcc_h31;
          let _1187_cf31 = _1180___mcc_h30;
          let _1188_cf30 = _1179___mcc_h29;
          let _1189_v63;
          let _init26 = ((_1190_cf33) => function (_1191_i2) {
            return (_1191_i2).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1190_cf33,new BigNumber(-538))).length));
          })(_1185_cf33);
          let _nw190 = Array((new BigNumber(21)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw190.length); _i0_26++) {
            _nw190[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1189_v63 = _nw190;
          _1189_v63 = _1189_v63;
          let _1192_v64;
          _1192_v64 = _dafny.MultiSet.fromElements((_this).f30);
          let _1193_v65;
          _1193_v65 = _dafny.Seq.of((_1192_v64).IsProperSubsetOf(_1192_v64));
          _1193_v65 = _dafny.Seq.update(((_1068_cf16) ? (_1193_v65) : (_1193_v65)), _module.__default.safeIndex(new BigNumber(243), new BigNumber((((_1068_cf16) ? (_1193_v65) : (_1193_v65))).length)), p1);
          _1068_cf16 = p2;
          let _1194_v66;
          _1194_v66 = _dafny.Seq.of(new BigNumber(547), _1186_cf32, _1188_cf30);
          let _1195_v67;
          let _nw191 = new _module.C3();
          _nw191.__ctor(_1194_v66, _1187_cf31, _dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), ((_1196_cf30) => function (_1197_i3) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1196_cf30,new BigNumber((_dafny.MultiSet.fromElements((_this).f29)).cardinality()));
          })(_1188_cf30)), p2);
          _1195_v67 = _nw191;
        } else if (_source11.is_DC10) {
          let _1198___mcc_h34 = (_source11).cf22;
          let _1199_cf22 = _1198___mcc_h34;
          let _1200_v68;
          _1200_v68 = new _dafny.CodePoint('v'.codePointAt(0));
          _1200_v68 = ((p2) ? (_1200_v68) : (_1200_v68));
          let _1201_v69;
          let _nw192 = new _module.C0();
          _nw192.__ctor((_this).f30, _1069_cf15);
          _1201_v69 = _nw192;
          let _1202_v70;
          _1202_v70 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18),false);
          let _1203_v71;
          _1203_v71 = _dafny.Set.fromElements((_1201_v69).f32, new BigNumber((_1202_v70).length));
          let _1204_v72;
          let _nw193 = new _module.C0();
          _nw193.__ctor((_1201_v69).f31, new BigNumber((_module.__default.fm50((((_1202_v70).contains(_module.__default.fm23(_1203_v71, (_1201_v69).f31, globalState))) ? ((_1202_v70).get(_module.__default.fm23(_1203_v71, (_1201_v69).f31, globalState))) : (p2)), (_this).f19, _1202_v70, _1200_v68, globalState)).length));
          _1204_v72 = _nw193;
          let _1205_v73;
          let _init27 = ((_1206_p2, _1207_v72) => function (_1208_i4) {
            return _module.__default.safeModuloInt(_1208_i4, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1206_p2,(_1207_v72).f31)).length));
          })(p2, _1204_v72);
          let _nw194 = Array((new BigNumber(5)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw194.length); _i0_27++) {
            _nw194[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1205_v73 = _nw194;
          let _1209_v74;
          _1209_v74 = _dafny.Seq.of(!(_1068_cf16));
          let _1210_v75;
          _1210_v75 = _dafny.MultiSet.fromElements(true);
          let _1211_v76;
          let _nw195 = Array((new BigNumber(27)).toNumber());
          _nw195[(_dafny.ZERO).toNumber()] = (_this).f29;
          _nw195[(_dafny.ONE).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(2)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(3)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(4)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(5)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(6)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(7)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(8)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(9)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(10)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(11)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(12)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(13)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(14)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(15)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(16)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _nw195[(new BigNumber(18)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(19)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(20)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(21)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(22)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(23)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(24)).toNumber()] = (_this).f29;
          _nw195[(new BigNumber(25)).toNumber()] = _1200_v68;
          _nw195[(new BigNumber(26)).toNumber()] = _1200_v68;
          _1211_v76 = _nw195;
          let _1212_v77;
          _1212_v77 = _module.D3.create_DC12(_1211_v76, p1, _module.__default.fm1(new BigNumber((_1210_v75).cardinality()), p2, true, globalState), _dafny.Map.Empty.slice().updateUnsafe(_1210_v75,(_this).f18), p2);
          let _1213_v78;
          _1213_v78 = _dafny.Set.fromElements((_this).f29);
          let _1214_v79;
          _1214_v79 = _dafny.Set.fromElements(false);
          let _1215_v80;
          _1215_v80 = _dafny.Seq.of((_this).f19);
          let _1216_v81;
          _1216_v81 = _dafny.Seq.of((_1215_v80)[_module.__default.safeIndex((_1201_v69).f31, new BigNumber((_1215_v80).length))]);
          let _1217_v82;
          let _nw196 = Array((new BigNumber(24)).toNumber());
          _nw196[(_dafny.ZERO).toNumber()] = p2;
          _nw196[(_dafny.ONE).toNumber()] = p1;
          _nw196[(new BigNumber(2)).toNumber()] = p2;
          _nw196[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_1209_v74, _module.__default.fm44(new BigNumber((_1210_v75).cardinality()), p1, _1200_v68, (_1204_v72).f32, globalState));
          _nw196[(new BigNumber(4)).toNumber()] = _1068_cf16;
          _nw196[(new BigNumber(5)).toNumber()] = ((_1210_v75)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1209_v74));
          _nw196[(new BigNumber(6)).toNumber()] = p1;
          _nw196[(new BigNumber(7)).toNumber()] = p1;
          _nw196[(new BigNumber(8)).toNumber()] = !(!((_this).f18) || ((_this).f18));
          _nw196[(new BigNumber(9)).toNumber()] = p2;
          _nw196[(new BigNumber(10)).toNumber()] = true;
          _nw196[(new BigNumber(11)).toNumber()] = (_this).f18;
          _nw196[(new BigNumber(12)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1212_v77).dtor_cf27,_1213_v78)).length)).isLessThan((_1201_v69).f31);
          _nw196[(new BigNumber(13)).toNumber()] = (_1214_v79).IsSubsetOf(_1214_v79);
          _nw196[(new BigNumber(14)).toNumber()] = !((_1209_v74)[_module.__default.safeIndex(_1067_cf17, new BigNumber((_1209_v74).length))]);
          _nw196[(new BigNumber(15)).toNumber()] = !(p2);
          _nw196[(new BigNumber(16)).toNumber()] = _1068_cf16;
          _nw196[(new BigNumber(17)).toNumber()] = ((_dafny.ZERO).minus((_1204_v72).f31)).isLessThan((_dafny.ZERO).minus((_1204_v72).f31));
          _nw196[(new BigNumber(18)).toNumber()] = _dafny.Seq.IsProperPrefixOf((_this).f19, (_this).f19);
          _nw196[(new BigNumber(19)).toNumber()] = (_this).f18;
          _nw196[(new BigNumber(20)).toNumber()] = !_dafny.Seq.contains(_1216_v81, (_this).f19);
          _nw196[(new BigNumber(21)).toNumber()] = true;
          _nw196[(new BigNumber(22)).toNumber()] = p1;
          _nw196[(new BigNumber(23)).toNumber()] = p1;
          _1217_v82 = _nw196;
          let _index187 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_1217_v82).length));
          (_1217_v82)[_index187] = p1;
          let _index188 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_1217_v82).length));
          let _rhs193 = new _dafny.CodePoint('v'.codePointAt(0));
          let _rhs194 = _module.__default.fm1(_1069_cf15, ((_dafny.ZERO).minus((_1204_v72).f32)).isEqualTo((_1204_v72).f32), _1068_cf16, globalState);
          let _rhs195 = ((p2) ? ((_1204_v72).f31) : (_1067_cf17));
          let _rhs196 = _1205_v73;
          let _rhs197 = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("tv"), (_this).f19);
          let _lhs167 = globalState;
          let _lhs168 = globalState;
          let _lhs169 = _1217_v82;
          let _lhs170 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_1217_v82).length));
          _1200_v68 = _rhs193;
          _lhs167.f0 = _rhs194;
          _lhs168.f15 = _rhs195;
          _1205_v73 = _rhs196;
          _lhs169[_lhs170] = _rhs197;
        } else {
          let _1218___mcc_h35 = (_source11).cf35;
          let _1219_cf35 = _1218___mcc_h35;
          (globalState).f0 = _1067_cf17;
          _1068_cf16 = !((_this).f18) || (false);
          (globalState).f2 = (_this).f18;
          let _1220_v83;
          _1220_v83 = _dafny.Seq.of(false, p2, false);
          let _1221_v84;
          _1221_v84 = _dafny.Seq.of(_1220_v83);
          (globalState).f2 = (_module.__default.safeDivisionInt(_1069_cf15, _1067_cf17)).isLessThanOrEqualTo(new BigNumber(((_1221_v84)[_module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber((_1221_v84).length))]).length));
        }
        (globalState).f9 = _dafny.Seq.UnicodeFromString("oqwb");
        let _1222_v85;
        _1222_v85 = _dafny.Seq.of(_1069_cf15, _1069_cf15, new BigNumber(148), new BigNumber(505), _1069_cf15);
        let _1223_v86;
        let _nw197 = new _module.C3();
        _nw197.__ctor(_1222_v85, (_this).f19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(935)), ((_1224_cf17) => function (_1225_i5) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1224_cf17,(_this).f30);
        })(_1067_cf17)), _1068_cf16);
        _1223_v86 = _nw197;
      } else if (_source9.is_DC8) {
        let _1226___mcc_h3 = (_source9).cf18;
        let _1227___mcc_h4 = (_source9).cf19;
        let _1228___mcc_h5 = (_source9).cf20;
        let _1229_cf20 = _1228___mcc_h5;
        let _1230_cf19 = _1227___mcc_h4;
        let _1231_cf18 = _1226___mcc_h3;
        let _1232_v87;
        _1232_v87 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f30);
        let _1233_v88;
        _1233_v88 = _dafny.Set.fromElements((((_1232_v87).contains((_this).f30)) ? ((_1232_v87).get((_this).f30)) : ((_dafny.ZERO).minus((_this).f30))), new BigNumber(391), (_this).f30, (_this).f30);
        let _1234_v89;
        _1234_v89 = _dafny.Set.fromElements((_dafny.Set.fromElements((_this).f30, (_this).f30)).IsDisjointFrom(_1233_v88), (_this).f18);
        let _1235_v90;
        _1235_v90 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1236_v91;
        _1236_v91 = _dafny.Seq.of(p1);
        let _1237_v92;
        _1237_v92 = _dafny.Map.Empty.slice().updateUnsafe((((_this).fm2((_this).f18, true, false, globalState)) ? (_1233_v88) : (_dafny.Set.fromElements(new BigNumber((_1236_v91).length)))),_module.__default.fm23(_1233_v88, (_this).f30, globalState));
        let _rhs198 = _module.__default.fm23(_1233_v88, (_dafny.ZERO).minus((_this).f30), globalState);
        let _rhs199 = _1234_v89;
        let _rhs200 = (_this).f18;
        let _rhs201 = (_this).f29;
        let _rhs202 = _1237_v92;
        let _lhs171 = globalState;
        let _lhs172 = globalState;
        _lhs171.f2 = _rhs198;
        _1234_v89 = _rhs199;
        _lhs172.f2 = _rhs200;
        _1235_v90 = _rhs201;
        _1237_v92 = _rhs202;
        let _1238_v93;
        _1238_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1235_v90,new BigNumber((_1232_v87).length));
        let _1239_v94;
        _1239_v94 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1238_v93).length),(_this).f18);
        let _1240_v95;
        _1240_v95 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1229_cf20);
        let _1241_v96;
        let _nw198 = Array((new BigNumber(20)).toNumber());
        _nw198[(_dafny.ZERO).toNumber()] = _1229_cf20;
        _nw198[(_dafny.ONE).toNumber()] = _1229_cf20;
        _nw198[(new BigNumber(2)).toNumber()] = _1231_cf18;
        _nw198[(new BigNumber(3)).toNumber()] = _1230_cf19;
        _nw198[(new BigNumber(4)).toNumber()] = _1230_cf19;
        _nw198[(new BigNumber(5)).toNumber()] = p2;
        _nw198[(new BigNumber(6)).toNumber()] = (_this).f18;
        _nw198[(new BigNumber(7)).toNumber()] = p1;
        _nw198[(new BigNumber(8)).toNumber()] = !((_this).f18);
        _nw198[(new BigNumber(9)).toNumber()] = (((_1239_v94).contains((_dafny.ZERO).minus(_module.__default.fm1((_this).f30, _1229_cf20, _1231_cf18, globalState)))) ? ((_1239_v94).get((_dafny.ZERO).minus(_module.__default.fm1((_this).f30, _1229_cf20, _1231_cf18, globalState)))) : (_1229_cf20));
        _nw198[(new BigNumber(10)).toNumber()] = p2;
        _nw198[(new BigNumber(11)).toNumber()] = _1231_cf18;
        _nw198[(new BigNumber(12)).toNumber()] = (_1236_v91)[_module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber((_1236_v91).length))];
        _nw198[(new BigNumber(13)).toNumber()] = _1231_cf18;
        _nw198[(new BigNumber(14)).toNumber()] = (((_1240_v95).contains(p2)) ? ((_1240_v95).get(p2)) : (_1230_cf19));
        _nw198[(new BigNumber(15)).toNumber()] = p2;
        _nw198[(new BigNumber(16)).toNumber()] = _1229_cf20;
        _nw198[(new BigNumber(17)).toNumber()] = _1230_cf19;
        _nw198[(new BigNumber(18)).toNumber()] = (_this).f18;
        _nw198[(new BigNumber(19)).toNumber()] = !(_1229_cf20);
        _1241_v96 = _nw198;
        let _1242_v97;
        _1242_v97 = _dafny.Set.fromElements(_1241_v96);
        let _1243_v98;
        _1243_v98 = _dafny.Seq.of(new BigNumber(416), (_this).f30);
        let _1244_v99;
        let _nw199 = new _module.C2();
        _nw199.__ctor(new BigNumber(((_1242_v97).Intersect(_1242_v97)).length), _1243_v98, (_1234_v89).IsProperSubsetOf(_dafny.Set.fromElements(p2)));
        _1244_v99 = _nw199;
        let _1245_v100;
        let _nw200 = new _module.C2();
        _nw200.__ctor(new BigNumber(((_this).f19).length), (_1244_v99).f34, _1229_cf20);
        _1245_v100 = _nw200;
        let _1246_v101;
        let _nw201 = new _module.C2();
        _nw201.__ctor((_1245_v100.f33).multipliedBy(new BigNumber((_dafny.Seq.of(p2, p1)).length)), (_1245_v100).f34, p1);
        _1246_v101 = _nw201;
        let _1247_v102;
        _1247_v102 = _module.D11.create_DC28(((true) ? ((_this).f18) : (_1231_cf18)));
        let _rhs203 = _1236_v91;
        let _rhs204 = (_dafny.ZERO).minus(_1245_v100.f33);
        let _rhs205 = _1246_v101;
        let _rhs206 = _1247_v102;
        let _lhs173 = _1244_v99;
        _1236_v91 = _rhs203;
        _lhs173.f33 = _rhs204;
        _1246_v101 = _rhs205;
        _1247_v102 = _rhs206;
      } else if (_source9.is_DC6) {
        let _1248___mcc_h6 = (_source9).cf14;
        let _1249_cf14 = _1248___mcc_h6;
        let _1250_v103;
        let _init28 = ((_1251_p2) => function (_1252_i6) {
          return _1251_p2;
        })(p2);
        let _nw202 = Array((new BigNumber(26)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw202.length); _i0_28++) {
          _nw202[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1250_v103 = _nw202;
        let _1253_v104;
        let _init29 = ((_1254_p1) => function (_1255_i7) {
          return _module.D9.create_DC23(_1254_p1, false, (_this).f30);
        })(p1);
        let _nw203 = Array((new BigNumber(27)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw203.length); _i0_29++) {
          _nw203[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1253_v104 = _nw203;
        let _1256_v105;
        _1256_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1253_v104,_1250_v103);
        _1250_v103 = (((_1256_v105).contains(_1253_v104)) ? ((_1256_v105).get(_1253_v104)) : (_1250_v103));
        let _rhs207 = (_this).f18;
        let _rhs208 = ((_this).f30).multipliedBy(_module.__default.safeModuloInt((_this).f30, (_this).f30));
        let _lhs174 = globalState;
        let _lhs175 = globalState;
        _lhs174.f2 = _rhs207;
        _lhs175.f3 = _rhs208;
        let _1257_v106;
        _1257_v106 = _dafny.MultiSet.fromElements((_this).f30);
        let _1258_v107;
        _1258_v107 = _dafny.Set.fromElements(_1257_v106, _1257_v106);
        let _1259_v108;
        _1259_v108 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_1257_v106);
        _1258_v107 = (_1258_v107).Union(_dafny.Set.fromElements(_1257_v106, (((_1259_v108).contains((_this).f30)) ? ((_1259_v108).get((_this).f30)) : (_1257_v106))));
        let _1260_v109;
        _1260_v109 = new _dafny.CodePoint('l'.codePointAt(0));
        _1260_v109 = ((_this).f19)[_module.__default.safeIndex((new BigNumber((_module.__default.fm48(globalState)).length)).plus((_this).f30), new BigNumber(((_this).f19).length))];
      } else {
        let _1261___mcc_h7 = (_source9).cf21;
        let _1262_cf21 = _1261___mcc_h7;
        let _1263_v110;
        _1263_v110 = _dafny.MultiSet.fromElements((_this).f30);
        let _1264_v111;
        _1264_v111 = _dafny.Seq.of((_this).f30, new BigNumber(((_this).f19).length), (_this).f30, new BigNumber((_1263_v110).cardinality()));
        let _1265_v112;
        let _nw204 = new _module.C0();
        _nw204.__ctor((_this).f30, (_1264_v111)[_module.__default.safeIndex((_this).f30, new BigNumber((_1264_v111).length))]);
        _1265_v112 = _nw204;
        let _1266_v113;
        let _nw205 = new _module.C1();
        _nw205.__ctor(p1);
        _1266_v113 = _nw205;
        let _1267_v114;
        let _nw206 = Array((new BigNumber(19)).toNumber()).fill(false);
        _1267_v114 = _nw206;
        let _1268_v115;
        _1268_v115 = _module.D0.create_DC1(p1, (_1265_v112).f32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), function (_1269_i8) {
  return (_this).f29;
}), true, (_1265_v112).f32);
        let _index189 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1267_v114).length));
        (_1267_v114)[_index189] = (_1268_v115).dtor_cf4;
        let _index190 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1267_v114).length));
        (_1267_v114)[_index190] = (p2) || (p1);
        let _1270_v116;
        _1270_v116 = _dafny.Seq.of((_1267_v114)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_1267_v114).length))]);
        let _index191 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1267_v114).length));
        let _rhs209 = p1;
        let _rhs210 = _dafny.Seq.Concat(_dafny.Seq.of(p1), _1270_v116);
        let _lhs176 = _1267_v114;
        let _lhs177 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1267_v114).length));
        _lhs176[_lhs177] = _rhs209;
        _1270_v116 = _rhs210;
      }
      let _1271_i9;
      _1271_i9 = _dafny.ZERO;
      L9: {
        while ((_this).f18) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1271_i9)) {
              break L9;
            }
            _1271_i9 = (_1271_i9).plus(_dafny.ONE);
            let _1272_v117;
            let _nw207 = new _module.C1();
            _nw207.__ctor(!(false));
            _1272_v117 = _nw207;
            let _1273_v119;
            _1273_v119 = _dafny.Map.Empty.slice().updateUnsafe(function () {
              let _coll34 = new _dafny.Map();
              for (const _compr_34 of (_dafny.Seq.update((_this).f19, _module.__default.safeIndex((_this).f30, new BigNumber(((_this).f19).length)), new _dafny.CodePoint('r'.codePointAt(0)))).Elements) {
                let _1274_v118 = _compr_34;
                if (_dafny.Seq.contains(_dafny.Seq.update((_this).f19, _module.__default.safeIndex((_this).f30, new BigNumber(((_this).f19).length)), new _dafny.CodePoint('r'.codePointAt(0))), _1274_v118)) {
                  _coll34.push([_1274_v118,(_this).f19]);
                }
              }
              return _coll34;
            }(),_1272_v117);
            let _1275_v120;
            _1275_v120 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f19);
            let _1276_v121;
            let _nw208 = Array((new BigNumber(28)).toNumber());
            _nw208[(_dafny.ZERO).toNumber()] = _1272_v117;
            _nw208[(_dafny.ONE).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(2)).toNumber()] = (((_1273_v119).contains(_1275_v120)) ? ((_1273_v119).get(_1275_v120)) : (_1272_v117));
            _nw208[(new BigNumber(3)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(4)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(5)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(6)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(7)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(8)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(9)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(10)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(11)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(12)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(13)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(14)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(15)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(16)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(17)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(18)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(19)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(20)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(21)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(22)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(23)).toNumber()] = ((p2) ? (_1272_v117) : (_1272_v117));
            _nw208[(new BigNumber(24)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(25)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(26)).toNumber()] = _1272_v117;
            _nw208[(new BigNumber(27)).toNumber()] = _1272_v117;
            _1276_v121 = _nw208;
            let _index192 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_1276_v121).length));
            (_1276_v121)[_index192] = _1272_v117;
            let _index193 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_1276_v121).length));
            (_1276_v121)[_index193] = _1272_v117;
            let _1277_v122;
            let _nw209 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
            _1277_v122 = _nw209;
            _1277_v122 = _1277_v122;
            let _1278_v123;
            _1278_v123 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
            (globalState).f2 = !((((_1278_v123).contains(true)) ? ((_1278_v123).get(true)) : (p1))) || (false);
            let _1279_v124;
            _1279_v124 = _dafny.MultiSet.fromElements((_this).f30);
            _1279_v124 = (_1279_v124).Union((((_this).f18) ? (_1279_v124) : (_1279_v124)));
          }
        }
      }
      let _1280_v125;
      _1280_v125 = _dafny.Set.fromElements(p2, !(true), p1, p1, (_this).f18);
      let _1281_v126;
      _1281_v126 = _dafny.Seq.of(new BigNumber((_1280_v125).length), (_this).f30, new BigNumber(603), (_dafny.ZERO).minus((_this).f30));
      let _1282_v127;
      _1282_v127 = _dafny.Seq.of(p2, p1);
      let _1283_v128;
      _1283_v128 = _module.D10.create_DC24(_1280_v125);
      let _1284_v129;
      let _nw210 = new _module.C3();
      _nw210.__ctor(_1281_v126, (_this).f19, _module.__default.fm51(p2, (_this).f30, (_1282_v127)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f30), new BigNumber((_1282_v127).length))], _1283_v128, globalState), !(p1) || ((_this).f18));
      _1284_v129 = _nw210;
      let _1285_v130;
      let _nw211 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
      _1285_v130 = _nw211;
      let _1286_v131;
      _1286_v131 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f19);
      let _index194 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_1285_v130).length));
      (_1285_v130)[_index194] = (_1286_v131).Merge(_1286_v131);
      let _index195 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_1285_v130).length));
      (_1285_v130)[_index195] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f30, (_this).f30),(_this).f19);
      let _1287_v132;
      let _nw212 = new _module.C1();
      _nw212.__ctor(p1);
      _1287_v132 = _nw212;
      (globalState).f2 = p2;
      r0 = ((_this).f30).plus(new BigNumber((_1282_v127).length));
      return r0;
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _1288_v0;
      _1288_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
      let _1289_v1;
      _1289_v1 = _dafny.MultiSet.fromElements((((_1288_v0).contains((_this).f18)) ? ((_1288_v0).get((_this).f18)) : ((_this).f18)));
      let _1290_v2;
      _1290_v2 = _dafny.Seq.of(_1289_v1);
      let _1291_v3;
      _1291_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_1289_v1).Union((_1290_v2)[_module.__default.safeIndex(p0, new BigNumber((_1290_v2).length))])).cardinality()));
      _1291_v3 = function () {
        let _coll35 = new _dafny.Map();
        for (const _compr_35 of _dafny.IntegerRange(new BigNumber(584), new BigNumber(278))) {
          let _1292_v4 = _compr_35;
          if (((new BigNumber(584)).isLessThanOrEqualTo(_1292_v4)) && ((_1292_v4).isLessThan(new BigNumber(278)))) {
            _coll35.push([_module.__default.safeDivisionInt(_1292_v4, (_dafny.ZERO).minus(p0)),new BigNumber((_dafny.Seq.of(new BigNumber(895), p0)).length)]);
          }
        }
        return _coll35;
      }();
      let _rhs211 = (_this).f18;
      let _rhs212 = (((p0).isLessThan((_this).f30)) ? ((_this).f30) : (p0));
      let _lhs178 = globalState;
      let _lhs179 = globalState;
      _lhs178.f2 = _rhs211;
      _lhs179.f3 = _rhs212;
      (globalState).f2 = (_this).f18;
      let _1293_v5;
      _1293_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f18);
      let _1294_v7;
      _1294_v7 = _dafny.Seq.of(_1293_v5, function () {
        let _coll36 = new _dafny.Map();
        for (const _compr_36 of (_dafny.Seq.of((_this).f29)).Elements) {
          let _1295_v6 = _compr_36;
          if (_dafny.Seq.contains(_dafny.Seq.of((_this).f29), _1295_v6)) {
            _coll36.push([_1295_v6,(_this).f18]);
          }
        }
        return _coll36;
      }(), _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f18), _1293_v5);
      let _1296_v8;
      _1296_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(222),_1291_v3);
      let _1297_v10;
      _1297_v10 = _dafny.Seq.of(_module.__default.fm48(globalState));
      let _1298_v11;
      _1298_v11 = _dafny.MultiSet.fromElements(_1291_v3, _1291_v3, ((((_1296_v8).contains(new BigNumber((_1288_v0).length))) ? ((_1296_v8).get(new BigNumber((_1288_v0).length))) : (_1291_v3))).Merge(_module.__default.fm30(_1290_v2, function () {
        let _coll37 = new _dafny.Map();
        for (const _compr_37 of (_1297_v10).Elements) {
          let _1299_v9 = _compr_37;
          if (_dafny.Seq.contains(_1297_v10, _1299_v9)) {
            _coll37.push([_1299_v9,!((_this).f18)]);
          }
        }
        return _coll37;
      }(), globalState)));
      let _rhs213 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1294_v7, _1294_v7), _dafny.Seq.Concat(_1294_v7, _1294_v7));
      let _rhs214 = _1298_v11;
      _1294_v7 = _rhs213;
      _1298_v11 = _rhs214;
      let _1300_v12;
      _1300_v12 = _dafny.Seq.of((_this).f18, (_this).f18, (_this).f18);
      let _1301_v13;
      _1301_v13 = _dafny.Map.Empty.slice().updateUnsafe((_1300_v12)[_module.__default.safeIndex(p0, new BigNumber((_1300_v12).length))],(_this).f29);
      let _1302_v14;
      _1302_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1301_v13).length),(_this).f18);
      let _1303_i0;
      _1303_i0 = _dafny.ZERO;
      L10: {
        while ((((_this).f18) ? ((((_1302_v14).contains(p0)) ? ((_1302_v14).get(p0)) : ((_this).f18))) : ((_this).f18))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1303_i0)) {
              break L10;
            }
            _1303_i0 = (_1303_i0).plus(_dafny.ONE);
            let _1304_v15;
            let _nw213 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _1304_v15 = _nw213;
            let _index196 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1304_v15).length));
            (_1304_v15)[_index196] = (_this).f30;
            let _index197 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1304_v15).length));
            (_1304_v15)[_index197] = (_this).f30;
            (globalState).f9 = (_this).f19;
            let _1305_v16;
            _1305_v16 = new _dafny.CodePoint('t'.codePointAt(0));
            _1305_v16 = (_this).f29;
            (globalState).f2 = !((_this).f18);
          }
        }
      }
      let _1306_v17;
      _1306_v17 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f30), new BigNumber(((_this).f19).length));
      let _1307_v18;
      _1307_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.update(_1306_v17, _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f30))), new BigNumber((_1306_v17).length)), (_this).f30));
      let _1308_v19;
      _1308_v19 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
      let _1309_v20;
      _1309_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Set.fromElements((_this).f18));
      let _1310_v21;
      _1310_v21 = _module.D9.create_DC23((_this).f18, (_this).f18, (_this).f30);
      let _1311_v22;
      let _init30 = ((_1312_v12) => function (_1313_i4) {
        return _module.__default.safeModuloInt(_1313_i4, new BigNumber((_1312_v12).length));
      })(_1300_v12);
      let _nw214 = Array((new BigNumber(7)).toNumber());
      for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw214.length); _i0_30++) {
        _nw214[_i0_30] = _init30(new BigNumber(_i0_30));
      }
      _1311_v22 = _nw214;
      let _1314_v23;
      _1314_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1311_v22,(_this).f18);
      let _1315_v24;
      _1315_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_dafny.Seq.of(new BigNumber(235)));
      let _1316_v25;
      _1316_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,new BigNumber((_1315_v24).length));
      let _1317_v26;
      let _nw215 = Array((new BigNumber(24)).toNumber());
      _nw215[(_dafny.ZERO).toNumber()] = _1306_v17;
      _nw215[(_dafny.ONE).toNumber()] = _dafny.Seq.of(p0, (_this).f30);
      _nw215[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-488)), ((_1318_v1, _1319_v3, _1320_p0) => function (_1321_i2) {
        return new BigNumber((_dafny.Seq.of(_module.D2.create_DC9(_module.D2.create_DC9(_module.D2.create_DC7((_this).f30, (_this).f18, (((_1318_v1).contains((_this).f18)) ? ((_1318_v1).get((_this).f18)) : (new BigNumber(62)))))), _module.D2.create_DC9(_module.D2.create_DC7((_this).f30, (_this).f18, new BigNumber((_1319_v3).length))), _module.D2.create_DC9(_module.D2.create_DC9(_module.D2.create_DC9(_module.D2.create_DC7(new BigNumber(((_this).f19).length), (_this).f18, _1320_p0)))), _module.D2.create_DC9(_module.D2.create_DC9(_module.D2.create_DC8((_this).f18, (_this).f18, (_this).f18))), _module.D2.create_DC9(_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1320_p0))))).length);
      })(_1289_v1, _1291_v3, p0));
      _nw215[(new BigNumber(3)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(4)).toNumber()] = _dafny.Seq.of((_this).f30);
      _nw215[(new BigNumber(5)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1306_v17, _1306_v17);
      _nw215[(new BigNumber(7)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(8)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1306_v17, _1306_v17);
      _nw215[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_1306_v17, _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1306_v17).length)), p0);
      _nw215[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(656)), function (_1322_i3) {
        return (_this).f30;
      });
      _nw215[(new BigNumber(12)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(p0);
      _nw215[(new BigNumber(14)).toNumber()] = _dafny.Seq.update(_1306_v17, _module.__default.safeIndex((_this).f30, new BigNumber((_1306_v17).length)), p0);
      _nw215[(new BigNumber(15)).toNumber()] = (((_1307_v18).contains(true)) ? ((_1307_v18).get(true)) : (_dafny.Seq.of(p0)));
      _nw215[(new BigNumber(16)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(17)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(18)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(19)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(20)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of((_this).f30, (_this).f30, (_this).f30, new BigNumber((_1308_v19).cardinality())), _module.__default.safeIndex(new BigNumber((_1309_v20).length), new BigNumber((_dafny.Seq.of((_this).f30, (_this).f30, (_this).f30, new BigNumber((_1308_v19).cardinality()))).length)), (_1310_v21).dtor_cf46), _module.__default.safeIndex(new BigNumber((_1314_v23).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f30, (_this).f30, (_this).f30, new BigNumber((_1308_v19).cardinality())), _module.__default.safeIndex(new BigNumber((_1309_v20).length), new BigNumber((_dafny.Seq.of((_this).f30, (_this).f30, (_this).f30, new BigNumber((_1308_v19).cardinality()))).length)), (_1310_v21).dtor_cf46)).length)), (((_1316_v25).contains((_this).f18)) ? ((_1316_v25).get((_this).f18)) : ((_1306_v17)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f18)).length), new BigNumber((_1306_v17).length))])));
      _nw215[(new BigNumber(22)).toNumber()] = _1306_v17;
      _nw215[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_1306_v17, _dafny.Seq.of(p0));
      _1317_v26 = _nw215;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1317_v26).length))) {
        let _1323_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1323_i1)) && ((_1323_i1).isLessThan(new BigNumber((_1317_v26).length))))) {
          (_1317_v26)[(_1323_i1)] = _dafny.Seq.of(_module.__default.safeModuloInt((_this).f30, p0), p0);
        }
      }
      r0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(431)), function (_1324_i5) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      });
      return r0;
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f18 = false;
      this.f27 = _dafny.Seq.UnicodeFromString("");
      this._f28 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f27, f28, f18) {
      let _this = this;
      (_this).f27 = f27;
      (_this)._f28 = f28;
      (_this)._f18 = f18;
      return;
    }
    fm14(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm15(p0, p1, globalState) {
      let _this = this;
      return !(!(new BigNumber(461)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18)).length))) || (false);
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      (globalState).f2 = (_this).f18;
      let _1325_v0;
      _1325_v0 = _module.D0.create_DC0(true);
      let _1326_v1;
      _1326_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_dafny.ZERO).minus(p1));
      let _1327_v2;
      _1327_v2 = _dafny.Map.Empty.slice().updateUnsafe((_1325_v0).dtor_cf0,_1326_v1);
      let _1328_v3;
      _1328_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(((((_1327_v2).contains((_this).f18)) ? ((_1327_v2).get((_this).f18)) : (_1326_v1))).length));
      let _1329_v4;
      _1329_v4 = _dafny.Seq.of((_this).f18, (_this).f18);
      let _rhs215 = (new BigNumber((_1328_v3).length)).plus(p1);
      let _rhs216 = _module.__default.safeDivisionInt(p1, new BigNumber((_1329_v4).length));
      let _rhs217 = (_this).f18;
      let _rhs218 = (_this).f18;
      let _rhs219 = (_this).f18;
      let _lhs180 = globalState;
      let _lhs181 = globalState;
      let _lhs182 = globalState;
      let _lhs183 = globalState;
      let _lhs184 = globalState;
      _lhs180.f0 = _rhs215;
      _lhs181.f3 = _rhs216;
      _lhs182.f2 = _rhs217;
      _lhs183.f2 = _rhs218;
      _lhs184.f2 = _rhs219;
      if ((p1).isLessThanOrEqualTo(new BigNumber(582))) {
        (globalState).f15 = p1;
        let _1330_v5;
        _1330_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
        (globalState).f2 = !((new BigNumber((_1330_v5).length)).multipliedBy((_dafny.ZERO).minus(_module.__default.fm1(p1, (_this).f18, (_this).f18, globalState)))).isEqualTo(p1);
        (globalState).f2 = (_this).f18;
        let _1331_v6;
        let _nw216 = Array((new BigNumber(18)).toNumber()).fill(false);
        _1331_v6 = _nw216;
        let _index198 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length));
        (_1331_v6)[_index198] = (_this).f18;
        let _index199 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length));
        (_1331_v6)[_index199] = !(_dafny.Seq.IsPrefixOf(_1329_v4, _dafny.Seq.Concat(_1329_v4, _1329_v4)));
        if ((_1329_v4)[_module.__default.safeIndex(p1, new BigNumber((_1329_v4).length))]) {
          let _1332_v7;
          _1332_v7 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1333_v8;
          _1333_v8 = _dafny.Set.fromElements(p1);
          let _1334_v9;
          let _nw217 = Array((new BigNumber(24)).toNumber());
          _nw217[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("m"), _module.__default.safeIndex(new BigNumber(224), new BigNumber((_dafny.Seq.UnicodeFromString("m")).length)), _1332_v7);
          _nw217[(_dafny.ONE).toNumber()] = _module.__default.fm16(p1, (_this).f18, (_this.f27)[_module.__default.safeIndex(p1, new BigNumber((_this.f27).length))], p1, globalState);
          _nw217[(new BigNumber(2)).toNumber()] = p2;
          _nw217[(new BigNumber(3)).toNumber()] = p2;
          _nw217[(new BigNumber(4)).toNumber()] = _this.f27;
          _nw217[(new BigNumber(5)).toNumber()] = p0;
          _nw217[(new BigNumber(6)).toNumber()] = _this.f27;
          _nw217[(new BigNumber(7)).toNumber()] = p0;
          _nw217[(new BigNumber(8)).toNumber()] = (_this).f28;
          _nw217[(new BigNumber(9)).toNumber()] = (_this).f28;
          _nw217[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("t"), (_this).f28);
          _nw217[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("gdjcvu");
          _nw217[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("y");
          _nw217[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("kxv"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("kxv")).length)), _1332_v7), p0);
          _nw217[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(p0, _dafny.Seq.UnicodeFromString("tclvbcm"));
          _nw217[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("c");
          _nw217[(new BigNumber(16)).toNumber()] = _this.f27;
          _nw217[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dsde"), (_this).f28);
          _nw217[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(p2, (_this).f28);
          _nw217[(new BigNumber(19)).toNumber()] = (_this).f28;
          _nw217[(new BigNumber(20)).toNumber()] = ((false) ? ((_this).f28) : (p2));
          _nw217[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(57)), ((_1335_v7) => function (_1336_i0) {
            return _1335_v7;
          })(_1332_v7));
          _nw217[(new BigNumber(22)).toNumber()] = (_this).f28;
          _nw217[(new BigNumber(23)).toNumber()] = _dafny.Seq.update(_this.f27, _module.__default.safeIndex(new BigNumber((_1333_v8).length), new BigNumber((_this.f27).length)), _1332_v7);
          _1334_v9 = _nw217;
          let _index200 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_1334_v9).length));
          (_1334_v9)[_index200] = p2;
          let _1337_v11;
          _1337_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
          let _1338_v13;
          _1338_v13 = _dafny.Seq.of(p1, (_dafny.ZERO).minus(p1), p1, p1, p1);
          let _1339_v14;
          _1339_v14 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe(p1,true));
          let _1340_v15;
          _1340_v15 = _dafny.MultiSet.fromElements((function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of _dafny.IntegerRange(new BigNumber(120), new BigNumber(247))) {
              let _1341_v10 = _compr_38;
              if (((new BigNumber(120)).isLessThanOrEqualTo(_1341_v10)) && ((_1341_v10).isLessThan(new BigNumber(247)))) {
                _coll38.push([(_1341_v10).plus(new BigNumber(23)),(_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))]]);
              }
            }
            return _coll38;
          }()).Merge(_1337_v11), (((_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))]) ? (function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of (_1338_v13).Elements) {
              let _1342_v12 = _compr_39;
              if (_dafny.Seq.contains(_1338_v13, _1342_v12)) {
                _coll39.push([_module.__default.safeDivisionInt(_1342_v12, new BigNumber(163)),(_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))]]);
              }
            }
            return _coll39;
          }()) : (_1337_v11)), ((((_1339_v14).contains(p1)) ? ((_1339_v14).get(p1)) : (_1337_v11))).Merge(_1337_v11));
          let _1343_v16;
          _1343_v16 = _dafny.Seq.of(p0);
          let _index201 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_1334_v9).length));
          let _rhs220 = p0;
          let _rhs221 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(475)), ((_1344_p2, _1345_v7) => function (_1346_i1) {
            return (_1344_p2)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1345_v7,(_this).f18)).length), new BigNumber((_1344_p2).length))];
          })(p2, _1332_v7)), (_1343_v16)[_module.__default.safeIndex(new BigNumber((_this.f27).length), new BigNumber((_1343_v16).length))]);
          let _rhs222 = _1340_v15;
          let _lhs185 = _1334_v9;
          let _lhs186 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_1334_v9).length));
          let _lhs187 = globalState;
          _lhs185[_lhs186] = _rhs220;
          _lhs187.f12 = _rhs221;
          _1340_v15 = _rhs222;
          (globalState).f0 = (_module.D2.create_DC7(new BigNumber((_1333_v8).length), (_this).f18, p1)).dtor_cf15;
          (globalState).f15 = ((p1).multipliedBy(p1)).plus(p1);
          let _1347_v17;
          let _nw218 = Array((new BigNumber(9)).toNumber());
          _nw218[(_dafny.ZERO).toNumber()] = p1;
          _nw218[(_dafny.ONE).toNumber()] = p1;
          _nw218[(new BigNumber(2)).toNumber()] = ((true) ? (p1) : (p1));
          _nw218[(new BigNumber(3)).toNumber()] = p1;
          _nw218[(new BigNumber(4)).toNumber()] = p1;
          _nw218[(new BigNumber(5)).toNumber()] = (p1).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))],_1332_v7)).length));
          _nw218[(new BigNumber(6)).toNumber()] = p1;
          _nw218[(new BigNumber(7)).toNumber()] = _module.__default.fm1(p1, (_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))], (_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))], globalState);
          _nw218[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p2, _this.f27)).length);
          _1347_v17 = _nw218;
          let _index202 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1347_v17).length));
          (_1347_v17)[_index202] = (p1).minus(p1);
          let _index203 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1347_v17).length));
          (_1347_v17)[_index203] = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(p1, p1), new BigNumber(-138));
          let _rhs223 = p1;
          let _rhs224 = ((!((_this).f18)) ? (new _dafny.CodePoint('s'.codePointAt(0))) : (_1332_v7));
          let _lhs188 = globalState;
          _lhs188.f0 = _rhs223;
          _1332_v7 = _rhs224;
        } else {
          let _index204 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length));
          (_1331_v6)[_index204] = (_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))];
          let _1348_v18;
          let _nw219 = new _module.C1();
          _nw219.__ctor((_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))]);
          _1348_v18 = _nw219;
          let _1349_v19;
          _1349_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1329_v4,(p1).multipliedBy(p1));
          let _1350_v20;
          let _init31 = function (_1351_i2) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          };
          let _nw220 = Array((new BigNumber(17)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw220.length); _i0_31++) {
            _nw220[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1350_v20 = _nw220;
          let _1352_v22;
          _1352_v22 = _dafny.MultiSet.fromElements(true);
          let _1353_v23;
          _1353_v23 = _dafny.Set.fromElements((_module.D3.create_DC12(_1350_v20, (_this).f18, (_dafny.ZERO).minus(p1), function () {
  let _coll40 = new _dafny.Map();
  for (const _compr_40 of (_dafny.MultiSet.fromElements(_1352_v22)).Elements) {
    let _1354_v21 = _compr_40;
    if ((_dafny.MultiSet.fromElements(_1352_v22)).contains(_1354_v21)) {
      _coll40.push([_1354_v21,(_1331_v6)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_1331_v6).length))]]);
    }
  }
  return _coll40;
}(), (_this).f18)).dtor_cf27, new BigNumber(-198), new BigNumber(35));
          (globalState).f0 = (((_1349_v19).contains(_1329_v4)) ? ((_1349_v19).get(_1329_v4)) : (_module.__default.safeDivisionInt(new BigNumber(((_this).f28).length), new BigNumber((_1353_v23).length))));
          _1353_v23 = (function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of _dafny.IntegerRange(new BigNumber(290), new BigNumber(598))) {
              let _1355_v24 = _compr_41;
              if (((new BigNumber(290)).isLessThanOrEqualTo(_1355_v24)) && ((_1355_v24).isLessThan(new BigNumber(598)))) {
                _coll41.add((_1355_v24).multipliedBy(p1));
              }
            }
            return _coll41;
          }()).Union(_1353_v23);
          let _1356_v25;
          _1356_v25 = _dafny.Seq.of(p1);
          (globalState).f2 = ((p1).multipliedBy(p1)).isLessThanOrEqualTo((((_1328_v3).contains(p1)) ? ((_1328_v3).get(p1)) : (new BigNumber((_1356_v25).length))));
        }
      } else {
        let _1357_v26;
        _1357_v26 = _module.D9.create_DC23((_1329_v4)[_module.__default.safeIndex(p1, new BigNumber((_1329_v4).length))], (_this).f18, p1);
        let _source12 = _module.D0.create_DC1((_1329_v4)[_module.__default.safeIndex(p1, new BigNumber((_1329_v4).length))], p1, p0, (_1357_v26).dtor_cf44, new BigNumber((_module.__default.fm43((_this).f18, false, globalState)).length));
        if (_source12.is_DC1) {
          let _1358___mcc_h0 = (_source12).cf1;
          let _1359___mcc_h1 = (_source12).cf2;
          let _1360___mcc_h2 = (_source12).cf3;
          let _1361___mcc_h3 = (_source12).cf4;
          let _1362___mcc_h4 = (_source12).cf5;
          let _1363_cf5 = _1362___mcc_h4;
          let _1364_cf4 = _1361___mcc_h3;
          let _1365_cf3 = _1360___mcc_h2;
          let _1366_cf2 = _1359___mcc_h1;
          let _1367_cf1 = _1358___mcc_h0;
          let _1368_v27;
          _1368_v27 = _dafny.Map.Empty.slice().updateUnsafe(p2,(new BigNumber((p2).length)).minus((_dafny.ZERO).minus(_1366_cf2)));
          _1368_v27 = (_1368_v27).update(_1365_cf3, p1);
          let _1369_v28;
          _1369_v28 = _dafny.Set.fromElements(_1364_cf4, (_this).f18);
          let _rhs225 = _1366_cf2;
          let _rhs226 = (_this).f18;
          let _rhs227 = ((((_1326_v1).contains(true)) ? ((_1326_v1).get(true)) : (_1366_cf2))).minus(new BigNumber((_this.f27).length));
          let _rhs228 = (_1369_v28).IsProperSubsetOf(_1369_v28);
          let _rhs229 = p1;
          let _lhs189 = globalState;
          let _lhs190 = globalState;
          let _lhs191 = globalState;
          _lhs189.f15 = _rhs225;
          _1367_cf1 = _rhs226;
          _1366_cf2 = _rhs227;
          _lhs190.f2 = _rhs228;
          _lhs191.f3 = _rhs229;
          let _1370_v29;
          _1370_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1364_cf4,_1367_cf1);
          let _1371_v30;
          _1371_v30 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1367_cf1)).update((_this).f18, _1367_cf1), _1370_v29);
          let _1372_v31;
          _1372_v31 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1371_v30);
          _1372_v31 = (_1372_v31).update((_dafny.ZERO).minus((((_1328_v3).contains(_1366_cf2)) ? ((_1328_v3).get(_1366_cf2)) : (_1363_cf5))), _dafny.Seq.Concat(_1371_v30, _1371_v30));
          let _1373_v32;
          _1373_v32 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(499)), function (_1374_i3) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length), _1366_cf2);
          let _1375_v33;
          _1375_v33 = _dafny.Set.fromElements((_1373_v32)[_module.__default.safeIndex(new BigNumber(734), new BigNumber((_1373_v32).length))]);
          let _1376_v34;
          let _init32 = function (_1377_i4) {
            return !((_module.D11.create_DC28((_this).f18)).dtor_cf51);
          };
          let _nw221 = Array((new BigNumber(21)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw221.length); _i0_32++) {
            _nw221[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1376_v34 = _nw221;
          let _1378_v35;
          _1378_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1376_v34,false);
          let _1379_v36;
          let _nw222 = Array((new BigNumber(26)).toNumber());
          _nw222[(_dafny.ZERO).toNumber()] = _1366_cf2;
          _nw222[(_dafny.ONE).toNumber()] = p1;
          _nw222[(new BigNumber(2)).toNumber()] = _1363_cf5;
          _nw222[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(_1366_cf2, new BigNumber((_1375_v33).length));
          _nw222[(new BigNumber(4)).toNumber()] = (_1366_cf2).minus(p1);
          _nw222[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_1366_cf2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1367_cf1,false)).length));
          _nw222[(new BigNumber(6)).toNumber()] = (_module.__default.fm1(new BigNumber((p0).length), !(_1367_cf1), (_this).f18, globalState)).plus(p1);
          _nw222[(new BigNumber(7)).toNumber()] = _1366_cf2;
          _nw222[(new BigNumber(8)).toNumber()] = _1366_cf2;
          _nw222[(new BigNumber(9)).toNumber()] = p1;
          _nw222[(new BigNumber(10)).toNumber()] = _1366_cf2;
          _nw222[(new BigNumber(11)).toNumber()] = ((_1373_v32)[_module.__default.safeIndex(_module.__default.fm1(p1, true, (_this).f18, globalState), new BigNumber((_1373_v32).length))]).plus(p1);
          _nw222[(new BigNumber(12)).toNumber()] = new BigNumber((_1373_v32).length);
          _nw222[(new BigNumber(13)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("kseym")).length)).minus(_1363_cf5);
          _nw222[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt(_1366_cf2, _1363_cf5);
          _nw222[(new BigNumber(15)).toNumber()] = _1366_cf2;
          _nw222[(new BigNumber(16)).toNumber()] = new BigNumber((_1370_v29).length);
          _nw222[(new BigNumber(17)).toNumber()] = (((_1328_v3).contains(new BigNumber(69))) ? ((_1328_v3).get(new BigNumber(69))) : (p1));
          _nw222[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1373_v32, _1373_v32)).length);
          _nw222[(new BigNumber(19)).toNumber()] = _1363_cf5;
          _nw222[(new BigNumber(20)).toNumber()] = _1366_cf2;
          _nw222[(new BigNumber(21)).toNumber()] = new BigNumber(680);
          _nw222[(new BigNumber(22)).toNumber()] = (p1).plus(p1);
          _nw222[(new BigNumber(23)).toNumber()] = new BigNumber(264);
          _nw222[(new BigNumber(24)).toNumber()] = new BigNumber((_1378_v35).length);
          _nw222[(new BigNumber(25)).toNumber()] = _1366_cf2;
          _1379_v36 = _nw222;
          let _index205 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_1379_v36).length));
          (_1379_v36)[_index205] = _1363_cf5;
          let _index206 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_1379_v36).length));
          (_1379_v36)[_index206] = _1366_cf2;
        } else {
          let _1380___mcc_h5 = (_source12).cf0;
          let _1381_cf0 = _1380___mcc_h5;
          (globalState).f3 = _module.__default.safeDivisionInt(p1, p1);
          let _1382_v37;
          _1382_v37 = _dafny.MultiSet.fromElements(_1381_cf0);
          let _1383_v38;
          _1383_v38 = _dafny.Seq.of((_dafny.ZERO).minus(p1), new BigNumber((_1382_v37).cardinality()));
          let _1384_v39;
          _1384_v39 = _dafny.Seq.of(_1328_v3, _1328_v3);
          let _1385_v40;
          let _nw223 = new _module.C3();
          _nw223.__ctor(_dafny.Seq.Concat(_1383_v38, _1383_v38), _this.f27, _1384_v39, (_1381_cf0) === (_1381_cf0));
          _1385_v40 = _nw223;
          let _1386_v41;
          let _nw224 = new _module.C2();
          _nw224.__ctor(new BigNumber(998), (_1385_v40).f35, (((_this).f18) ? ((_this).f18) : ((_1329_v4)[_module.__default.safeIndex(p1, new BigNumber((_1329_v4).length))])));
          _1386_v41 = _nw224;
          let _1387_v42;
          _1387_v42 = _dafny.Set.fromElements(_1381_cf0);
          let _1388_v43;
          _1388_v43 = new _dafny.CodePoint('k'.codePointAt(0));
          let _pat_let_tv22 = _1381_cf0;
          let _1389_v44;
          _1389_v44 = _dafny.MultiSet.fromElements(function (_pat_let32_0) {
            return function (_1390_dt__update__tmp_h0) {
              return function (_pat_let33_0) {
                return function (_1391_dt__update_hcf18_h0) {
                  return function (_pat_let34_0) {
                    return function (_1392_dt__update_hcf20_h0) {
                      return _module.D2.create_DC8(_1391_dt__update_hcf18_h0, (_1390_dt__update__tmp_h0).dtor_cf19, _1392_dt__update_hcf20_h0);
                    }(_pat_let34_0);
                  }(_pat_let_tv22);
                }(_pat_let33_0);
              }((_this).f18);
            }(_pat_let32_0);
          }(_module.D2.create_DC8(_1381_cf0, _1381_cf0, !(_1381_cf0))));
          let _1393_v45;
          _1393_v45 = _module.D0.create_DC1(_1381_cf0, new BigNumber((_1387_v42).length), _dafny.Seq.of(_1388_v43), (_this).f18, new BigNumber((_1389_v44).cardinality()));
          let _1394_v46;
          _1394_v46 = _dafny.Set.fromElements(_1386_v41.f33);
          let _1395_v47;
          _1395_v47 = _dafny.MultiSet.fromElements(new BigNumber(364), _1386_v41.f33, new BigNumber(-990));
          let _1396_v48;
          let _nw225 = Array((new BigNumber(18)).toNumber());
          _nw225[(_dafny.ZERO).toNumber()] = _1381_cf0;
          _nw225[(_dafny.ONE).toNumber()] = _1381_cf0;
          _nw225[(new BigNumber(2)).toNumber()] = _1381_cf0;
          _nw225[(new BigNumber(3)).toNumber()] = (_1381_cf0) === (!((_1393_v45).dtor_cf4));
          _nw225[(new BigNumber(4)).toNumber()] = !(_1381_cf0);
          _nw225[(new BigNumber(5)).toNumber()] = _1381_cf0;
          _nw225[(new BigNumber(6)).toNumber()] = _1381_cf0;
          _nw225[(new BigNumber(7)).toNumber()] = (_module.__default.fm1(_1386_v41.f33, (_this).f18, (_this).f18, globalState)).isLessThanOrEqualTo(p1);
          _nw225[(new BigNumber(8)).toNumber()] = (_1394_v46).equals(_1394_v46);
          _nw225[(new BigNumber(9)).toNumber()] = (_this).fm14(_1386_v41.f33, _1329_v4, _1395_v47, globalState);
          _nw225[(new BigNumber(10)).toNumber()] = _1381_cf0;
          _nw225[(new BigNumber(11)).toNumber()] = !(!(_1381_cf0)) || ((_1385_v40).fm2(_1381_cf0, false, (_this).f18, globalState));
          _nw225[(new BigNumber(12)).toNumber()] = false;
          _nw225[(new BigNumber(13)).toNumber()] = (_this).f18;
          _nw225[(new BigNumber(14)).toNumber()] = false;
          _nw225[(new BigNumber(15)).toNumber()] = (_1394_v46).IsDisjointFrom(_1394_v46);
          _nw225[(new BigNumber(16)).toNumber()] = (_this).f18;
          _nw225[(new BigNumber(17)).toNumber()] = _1381_cf0;
          _1396_v48 = _nw225;
          _1396_v48 = _1396_v48;
        }
        let _1397_v50;
        _1397_v50 = _dafny.MultiSet.fromElements(p1);
        let _1398_v52;
        _1398_v52 = _dafny.MultiSet.fromElements((_this).f18);
        let _1399_v54;
        _1399_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of _dafny.IntegerRange(new BigNumber(359), new BigNumber(764))) {
            let _1400_v53 = _compr_42;
            if (((new BigNumber(359)).isLessThanOrEqualTo(_1400_v53)) && ((_1400_v53).isLessThan(new BigNumber(764)))) {
              _coll42.push([(_1400_v53).multipliedBy(new BigNumber((_1326_v1).length)),(_dafny.ZERO).minus(p1)]);
            }
          }
          return _coll42;
        }());
        let _1401_v55;
        _1401_v55 = _dafny.Seq.of(p1, (_dafny.ZERO).minus(p1), p1, p1);
        let _1402_v56;
        _1402_v56 = new _dafny.CodePoint('w'.codePointAt(0));
        let _1403_v57;
        let _nw226 = Array((new BigNumber(25)).toNumber());
        _nw226[(_dafny.ZERO).toNumber()] = _1328_v3;
        _nw226[(_dafny.ONE).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(2)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(3)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(4)).toNumber()] = function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of _dafny.IntegerRange(new BigNumber(-938), new BigNumber(151))) {
            let _1404_v49 = _compr_43;
            if (((new BigNumber(-938)).isLessThanOrEqualTo(_1404_v49)) && ((_1404_v49).isLessThan(new BigNumber(151)))) {
              _coll43.push([_module.__default.safeDivisionInt(_1404_v49, p1),p1]);
            }
          }
          return _coll43;
        }();
        _nw226[(new BigNumber(5)).toNumber()] = (_1328_v3).update(new BigNumber((_1397_v50).cardinality()), p1);
        _nw226[(new BigNumber(6)).toNumber()] = (_1328_v3).update(p1, _module.__default.fm1(p1, !((_this).f18), (_this).f18, globalState));
        _nw226[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),p1);
        _nw226[(new BigNumber(8)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(9)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(10)).toNumber()] = function () {
          let _coll44 = new _dafny.Map();
          for (const _compr_44 of _dafny.IntegerRange(new BigNumber(504), new BigNumber(-597))) {
            let _1405_v51 = _compr_44;
            if (((new BigNumber(504)).isLessThanOrEqualTo(_1405_v51)) && ((_1405_v51).isLessThan(new BigNumber(-597)))) {
              _coll44.push([_module.__default.safeDivisionInt(_1405_v51, p1),p1]);
            }
          }
          return _coll44;
        }();
        _nw226[(new BigNumber(11)).toNumber()] = (_1328_v3).update(p1, new BigNumber((_dafny.Seq.UnicodeFromString("pxqm")).length));
        _nw226[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_1398_v52).contains(false)) ? ((_1398_v52).get(false)) : (new BigNumber(699))));
        _nw226[(new BigNumber(13)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(14)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(15)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(16)).toNumber()] = (((_1399_v54).contains(new BigNumber((_module.__default.fm16(new BigNumber((_1401_v55).length), (_this).f18, _1402_v56, new BigNumber(858), globalState)).length))) ? ((_1399_v54).get(new BigNumber((_module.__default.fm16(new BigNumber((_1401_v55).length), (_this).f18, _1402_v56, new BigNumber(858), globalState)).length))) : (_1328_v3));
        _nw226[(new BigNumber(17)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(p1));
        _nw226[(new BigNumber(19)).toNumber()] = (_1328_v3).Merge(_1328_v3);
        _nw226[(new BigNumber(20)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(21)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(22)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(23)).toNumber()] = _1328_v3;
        _nw226[(new BigNumber(24)).toNumber()] = _1328_v3;
        _1403_v57 = _nw226;
        let _index207 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_1403_v57).length));
        (_1403_v57)[_index207] = _1328_v3;
        let _index208 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_1403_v57).length));
        (_1403_v57)[_index208] = _1328_v3;
        if ((_this).f18) {
          (globalState).f15 = _module.__default.fm1(p1, (_1325_v0).dtor_cf0, false, globalState);
          _1398_v52 = (_1398_v52).Intersect(_1398_v52);
          (globalState).f2 = (_this).f18;
          let _1406_v58;
          _1406_v58 = _dafny.Map.Empty.slice().updateUnsafe(p1,!(true));
          let _1407_v59;
          _1407_v59 = _dafny.Set.fromElements(new BigNumber(-892), p1);
          let _1408_v60;
          _1408_v60 = _dafny.Seq.of(_1407_v59);
          let _1409_v61;
          _1409_v61 = _dafny.Map.Empty.slice().updateUnsafe(true,_1397_v50);
          _1406_v58 = (_1406_v58).update((p1).plus(new BigNumber(((_1408_v60)[_module.__default.safeIndex(new BigNumber((_1409_v61).length), new BigNumber((_1408_v60).length))]).length)), false);
          let _1410_v62;
          _1410_v62 = _dafny.Seq.of(_module.__default.fm43((_this).f18, !((_this).f18), globalState), _1326_v1, _1326_v1, _1326_v1);
          let _1411_v63;
          let _nw227 = new _module.C0();
          _nw227.__ctor(p1, p1);
          _1411_v63 = _nw227;
          let _1412_v64;
          _1412_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1411_v63,(_this).f18);
          let _1413_v65;
          let _nw228 = Array((new BigNumber(11)).toNumber());
          _nw228[(_dafny.ZERO).toNumber()] = _1326_v1;
          _nw228[(_dafny.ONE).toNumber()] = _1326_v1;
          _nw228[(new BigNumber(2)).toNumber()] = _1326_v1;
          _nw228[(new BigNumber(3)).toNumber()] = _1326_v1;
          _nw228[(new BigNumber(4)).toNumber()] = _1326_v1;
          _nw228[(new BigNumber(5)).toNumber()] = (_1326_v1).update(false, p1);
          _nw228[(new BigNumber(6)).toNumber()] = (_1410_v62)[_module.__default.safeIndex(new BigNumber(((((_1399_v54).contains(p1)) ? ((_1399_v54).get(p1)) : (_dafny.Map.Empty.slice().updateUnsafe(p1,p1)))).length), new BigNumber((_1410_v62).length))];
          _nw228[(new BigNumber(7)).toNumber()] = _module.__default.fm29(true, p1, _dafny.Seq.UnicodeFromString("rsbldutix"), globalState);
          _nw228[(new BigNumber(8)).toNumber()] = (_1326_v1).Merge(_1326_v1);
          _nw228[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,new BigNumber((_1412_v64).length));
          _nw228[(new BigNumber(10)).toNumber()] = (_1326_v1).update(!((_this).f18), new BigNumber(988));
          _1413_v65 = _nw228;
          _1413_v65 = _1413_v65;
        } else {
          (globalState).f15 = p1;
          (globalState).f2 = (_this).f18;
          let _1414_v66;
          _1414_v66 = _dafny.MultiSet.fromElements(_1401_v55);
          let _1415_v67;
          let _nw229 = Array((new BigNumber(19)).toNumber());
          _nw229[(_dafny.ZERO).toNumber()] = (new BigNumber(488)).plus(new BigNumber(-771));
          _nw229[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("gmvtc")).length);
          _nw229[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-186));
          _nw229[(new BigNumber(3)).toNumber()] = new BigNumber(-800);
          _nw229[(new BigNumber(4)).toNumber()] = p1;
          _nw229[(new BigNumber(5)).toNumber()] = p1;
          _nw229[(new BigNumber(6)).toNumber()] = p1;
          _nw229[(new BigNumber(7)).toNumber()] = p1;
          _nw229[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(p1, (((_1414_v66).contains(_1401_v55)) ? ((_1414_v66).get(_1401_v55)) : (p1)));
          _nw229[(new BigNumber(9)).toNumber()] = new BigNumber(-40);
          _nw229[(new BigNumber(10)).toNumber()] = new BigNumber(-801);
          _nw229[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1401_v55).length));
          _nw229[(new BigNumber(12)).toNumber()] = new BigNumber(577);
          _nw229[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((((_1326_v1).contains((_this).f18)) ? ((_1326_v1).get((_this).f18)) : (p1)), (_dafny.ZERO).minus(new BigNumber((_this.f27).length))));
          _nw229[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
          _nw229[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw229[(new BigNumber(16)).toNumber()] = p1;
          _nw229[(new BigNumber(17)).toNumber()] = p1;
          _nw229[(new BigNumber(18)).toNumber()] = new BigNumber(426);
          _1415_v67 = _nw229;
          let _1416_v68;
          let _nw230 = Array((new BigNumber(9)).toNumber());
          _nw230[(_dafny.ZERO).toNumber()] = _1402_v56;
          _nw230[(_dafny.ONE).toNumber()] = (p0)[_module.__default.safeIndex(p1, new BigNumber((p0).length))];
          _nw230[(new BigNumber(2)).toNumber()] = _1402_v56;
          _nw230[(new BigNumber(3)).toNumber()] = _1402_v56;
          _nw230[(new BigNumber(4)).toNumber()] = _1402_v56;
          _nw230[(new BigNumber(5)).toNumber()] = _1402_v56;
          _nw230[(new BigNumber(6)).toNumber()] = _1402_v56;
          _nw230[(new BigNumber(7)).toNumber()] = _1402_v56;
          _nw230[(new BigNumber(8)).toNumber()] = _module.__default.fm46(globalState);
          _1416_v68 = _nw230;
          let _1417_v69;
          _1417_v69 = _dafny.Seq.of(_1398_v52);
          let _1418_v70;
          _1418_v70 = _dafny.Map.Empty.slice().updateUnsafe((_1417_v69)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1417_v69).length))],true);
          let _1419_v71;
          _1419_v71 = _module.D3.create_DC12(_1416_v68, !(!((_this).f18)), p1, _1418_v70, (_this).f18);
          let _1420_v72;
          _1420_v72 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f18);
          let _1421_v73;
          _1421_v73 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(_1326_v1),_1402_v56);
          let _1422_v74;
          let _nw231 = new _module.C4();
          _nw231.__ctor(_1402_v56, p1, (_this).f18, p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(574)), ((_1423_v57) => function (_1424_i5) {
            return (_1423_v57)[_module.__default.safeIndex(new BigNumber(393), new BigNumber((_1423_v57).length))];
          })(_1403_v57)));
          _1422_v74 = _nw231;
          let _1425_v75;
          _1425_v75 = _dafny.Map.Empty.slice().updateUnsafe((((_1326_v1).contains(true)) ? ((_1326_v1).get(true)) : (new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f18, (_this).f18), _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_dafny.Seq.of((_this).f18, (_this).f18)).length)), (_this).f18)).length))),_1422_v74);
          let _1426_v76;
          _1426_v76 = _dafny.Seq.of(_1425_v75);
          let _1427_v77;
          _1427_v77 = _module.D13.create_DC33((_1422_v74).f18, p1, (_this).f18, (_dafny.ZERO).minus(new BigNumber((_1401_v55).length)));
          let _nw232 = Array((new BigNumber(28)).toNumber());
          _nw232[(_dafny.ZERO).toNumber()] = p1;
          _nw232[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(p1, new BigNumber(836));
          _nw232[(new BigNumber(2)).toNumber()] = p1;
          _nw232[(new BigNumber(3)).toNumber()] = p1;
          _nw232[(new BigNumber(4)).toNumber()] = (((_this).f18) ? (p1) : (p1));
          _nw232[(new BigNumber(5)).toNumber()] = (new BigNumber((_1326_v1).length)).minus(p1);
          _nw232[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(486), p1);
          _nw232[(new BigNumber(7)).toNumber()] = (((_this).f18) ? (p1) : (p1));
          _nw232[(new BigNumber(8)).toNumber()] = (_1419_v71).dtor_cf27;
          _nw232[(new BigNumber(9)).toNumber()] = p1;
          _nw232[(new BigNumber(10)).toNumber()] = new BigNumber((_1420_v72).length);
          _nw232[(new BigNumber(11)).toNumber()] = new BigNumber((_1421_v73).length);
          _nw232[(new BigNumber(12)).toNumber()] = (new BigNumber(-963)).plus(p1);
          _nw232[(new BigNumber(13)).toNumber()] = p1;
          _nw232[(new BigNumber(14)).toNumber()] = new BigNumber((_1397_v50).cardinality());
          _nw232[(new BigNumber(15)).toNumber()] = (p1).minus(p1);
          _nw232[(new BigNumber(16)).toNumber()] = new BigNumber(942);
          _nw232[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((new BigNumber(379)).plus(new BigNumber((_1426_v76).length)));
          _nw232[(new BigNumber(18)).toNumber()] = (((_1397_v50).contains(p1)) ? ((_1397_v50).get(p1)) : (new BigNumber((_1401_v55).length)));
          _nw232[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw232[(new BigNumber(20)).toNumber()] = ((_dafny.ZERO).minus(p1)).minus((_dafny.ZERO).minus(p1));
          _nw232[(new BigNumber(21)).toNumber()] = (_1427_v77).dtor_cf56;
          _nw232[(new BigNumber(22)).toNumber()] = _module.__default.safeModuloInt(p1, p1);
          _nw232[(new BigNumber(23)).toNumber()] = p1;
          _nw232[(new BigNumber(24)).toNumber()] = p1;
          _nw232[(new BigNumber(25)).toNumber()] = p1;
          _nw232[(new BigNumber(26)).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
          _nw232[(new BigNumber(27)).toNumber()] = (p1).plus(new BigNumber((_1401_v55).length));
          _1415_v67 = _nw232;
          let _1428_v78;
          let _1429_v79;
          let _1430_v80;
          let _1431_v81;
          let _out36;
          let _out37;
          let _out38;
          let _out39;
          let _outcollector9 = _module.__default.m0(globalState);
          _out36 = _outcollector9[0];
          _out37 = _outcollector9[1];
          _out38 = _outcollector9[2];
          _out39 = _outcollector9[3];
          _1428_v78 = _out36;
          _1429_v79 = _out37;
          _1430_v80 = _out38;
          _1431_v81 = _out39;
          let _1432_v82;
          let _nw233 = Array((new BigNumber(15)).toNumber()).fill(false);
          _1432_v82 = _nw233;
          _1432_v82 = _1432_v82;
        }
        let _1433_v83;
        _1433_v83 = _1398_v52;
        _1433_v83 = _1433_v83;
        let _1434_v84;
        let _nw234 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _1434_v84 = _nw234;
        let _index209 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_1434_v84).length));
        (_1434_v84)[_index209] = new BigNumber((p0).length);
        let _index210 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_1434_v84).length));
        let _rhs230 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(556)), ((_1435_v56) => function (_1436_i6) {
          return _1435_v56;
        })(_1402_v56));
        let _rhs231 = p1;
        let _lhs192 = globalState;
        let _lhs193 = _1434_v84;
        let _lhs194 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_1434_v84).length));
        _lhs192.f12 = _rhs230;
        _lhs193[_lhs194] = _rhs231;
      }
      let _1437_v85;
      _1437_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_module.__default.fm48(globalState));
      _1437_v85 = (_1437_v85).update((_this).f18, _this.f27);
      (globalState).f0 = _module.__default.fm1(p1, (_this).f18, (_this).f18, globalState);
      (globalState).f3 = p1;
      return;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let _1438_v0;
      _1438_v0 = _dafny.Set.fromElements(p1, p1, p1, p1, p1);
      let _1439_v1;
      _1439_v1 = _dafny.MultiSet.fromElements(_module.__default.fm23(_1438_v0, p1, globalState), (_this).f18);
      let _1440_v2;
      _1440_v2 = _dafny.Seq.of((_1439_v1).IsDisjointFrom(_1439_v1));
      (globalState).f2 = (_1440_v2)[_module.__default.safeIndex(p1, new BigNumber((_1440_v2).length))];
      if (((_this).f18) || (p0)) {
        let _1441_v3;
        let _nw235 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1441_v3 = _nw235;
        let _index211 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length));
        (_1441_v3)[_index211] = p0;
        let _index212 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length));
        (_1441_v3)[_index212] = (p0) && (true);
        let _1442_v4;
        let _init33 = ((_1443_v2) => function (_1444_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(false,_1443_v2);
        })(_1440_v2);
        let _nw236 = Array((new BigNumber(24)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw236.length); _i0_33++) {
          _nw236[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1442_v4 = _nw236;
        let _1445_v5;
        _1445_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1440_v2);
        let _index213 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1442_v4).length));
        (_1442_v4)[_index213] = _1445_v5;
        let _index214 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1442_v4).length));
        (_1442_v4)[_index214] = _1445_v5;
        let _index215 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length));
        let _index216 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length));
        let _rhs232 = _module.__default.safeModuloInt(p1, p1);
        let _rhs233 = (_this).f18;
        let _rhs234 = !_dafny.Seq.contains(_dafny.Seq.Concat(_1440_v2, _1440_v2), p0);
        let _rhs235 = _dafny.Seq.IsProperPrefixOf(_this.f27, _dafny.Seq.UnicodeFromString("gnynibgy"));
        let _rhs236 = _1441_v3;
        let _lhs195 = globalState;
        let _lhs196 = _1441_v3;
        let _lhs197 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length));
        let _lhs198 = globalState;
        let _lhs199 = _1441_v3;
        let _lhs200 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length));
        _lhs195.f3 = _rhs232;
        _lhs196[_lhs197] = _rhs233;
        _lhs198.f2 = _rhs234;
        _lhs199[_lhs200] = _rhs235;
        _1441_v3 = _rhs236;
        if (false) {
          let _1446_v6;
          _1446_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1441_v3,_dafny.Seq.contains(_1440_v2, (_1441_v3)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length))]));
          _1446_v6 = (_1446_v6).update(_1441_v3, p0);
          let _1447_v7;
          _1447_v7 = _dafny.Set.fromElements((_this).f18, (_this).f18, (_this).f18);
          let _1448_v8;
          _1448_v8 = _dafny.Map.Empty.slice().updateUnsafe((_1441_v3)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length))],_1447_v7);
          _1448_v8 = (_1448_v8).update((_1441_v3)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length))], (_dafny.Set.fromElements((_1441_v3)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length))], !(!((_this).f18)))).Difference(_1447_v7));
          let _index217 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length));
          (_1441_v3)[_index217] = (_1441_v3)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length))];
          let _index218 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length));
          (_1441_v3)[_index218] = (_1441_v3)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length))];
          (globalState).f0 = (((_this).f18) ? (p1) : (new BigNumber(((_1438_v0).Intersect(_1438_v0)).length)));
        } else {
          let _1449_v9;
          _1449_v9 = _dafny.Seq.of((p1).minus(p1));
          (globalState).f0 = (_1449_v9)[_module.__default.safeIndex(p1, new BigNumber((_1449_v9).length))];
          (globalState).f15 = p1;
          let _1450_v10;
          let _init34 = ((_1451_v9) => function (_1452_i1) {
            return _1451_v9;
          })(_1449_v9);
          let _nw237 = Array((new BigNumber(16)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw237.length); _i0_34++) {
            _nw237[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1450_v10 = _nw237;
          let _index219 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1450_v10).length));
          (_1450_v10)[_index219] = _1449_v9;
          let _index220 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1450_v10).length));
          (_1450_v10)[_index220] = _dafny.Seq.Concat(_1449_v9, _dafny.Seq.Concat(_dafny.Seq.of(p1), _1449_v9));
          let _1453_v11;
          _1453_v11 = _dafny.Seq.of(_1438_v0);
          (globalState).f2 = _module.__default.fm23((_1453_v11)[_module.__default.safeIndex(p1, new BigNumber((_1453_v11).length))], (_dafny.ZERO).minus(p1), globalState);
          let _1454_v12;
          _1454_v12 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1455_v13;
          _1455_v13 = _dafny.MultiSet.fromElements((_this).f28, _dafny.Seq.update(_this.f27, _module.__default.safeIndex(p1, new BigNumber((_this.f27).length)), _1454_v12), _dafny.Seq.update(_this.f27, _module.__default.safeIndex(new BigNumber(98), new BigNumber((_this.f27).length)), _1454_v12));
          let _1456_v14;
          let _nw238 = Array((new BigNumber(20)).toNumber());
          _nw238[(_dafny.ZERO).toNumber()] = _1455_v13;
          _nw238[(_dafny.ONE).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(2)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_this.f27, _this.f27, _this.f27);
          _nw238[(new BigNumber(4)).toNumber()] = (_1455_v13).Intersect(_dafny.MultiSet.fromElements(_this.f27));
          _nw238[(new BigNumber(5)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(6)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(7)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(8)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(9)).toNumber()] = _module.__default.fm52(p1, p1, globalState);
          _nw238[(new BigNumber(10)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(11)).toNumber()] = _module.__default.fm52(p1, p1, globalState);
          _nw238[(new BigNumber(12)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(13)).toNumber()] = (_1455_v13).Intersect(_1455_v13);
          _nw238[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(_this.f27, _this.f27, (_this).f28);
          _nw238[(new BigNumber(15)).toNumber()] = (_1455_v13).Union(_1455_v13);
          _nw238[(new BigNumber(16)).toNumber()] = (_1455_v13).Difference(_1455_v13);
          _nw238[(new BigNumber(17)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(18)).toNumber()] = _1455_v13;
          _nw238[(new BigNumber(19)).toNumber()] = _module.__default.fm52(new BigNumber((_1439_v1).cardinality()), new BigNumber(798), globalState);
          _1456_v14 = _nw238;
          let _index221 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_1456_v14).length));
          (_1456_v14)[_index221] = _1455_v13;
          let _1457_v15;
          _1457_v15 = _module.D12.create_DC31();
          let _index222 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_1456_v14).length));
          let _rhs237 = !((_this).f18);
          let _rhs238 = (_dafny.MultiSet.fromElements((_this).f28, (_this).f28, (_this).f28)).Difference(_1455_v13);
          let _rhs239 = _dafny.Seq.contains(_this.f27, _1454_v12);
          let _rhs240 = _1457_v15;
          let _lhs201 = globalState;
          let _lhs202 = _1456_v14;
          let _lhs203 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_1456_v14).length));
          let _lhs204 = globalState;
          _lhs201.f2 = _rhs237;
          _lhs202[_lhs203] = _rhs238;
          _lhs204.f2 = _rhs239;
          _1457_v15 = _rhs240;
        }
        _1440_v2 = _dafny.Seq.of((_1441_v3)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_1441_v3).length))]);
      } else {
        let _1458_v16;
        let _nw239 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
        _1458_v16 = _nw239;
        let _index223 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1458_v16).length));
        (_1458_v16)[_index223] = _1440_v2;
        let _1459_v17;
        _1459_v17 = _dafny.Seq.of(_dafny.Seq.of((_this).f18), _dafny.Seq.Concat(_1440_v2, _dafny.Seq.of(p0)), _dafny.Seq.update(_1440_v2, _module.__default.safeIndex(p1, new BigNumber((_1440_v2).length)), p0), _dafny.Seq.Concat(_1440_v2, _1440_v2));
        let _1460_v18;
        _1460_v18 = _dafny.Seq.of(p1);
        let _index224 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1458_v16).length));
        let _rhs241 = p0;
        let _rhs242 = ((_this).f18) && ((_this).f18);
        let _rhs243 = (_1459_v17)[_module.__default.safeIndex(_module.__default.fm1(new BigNumber((_1460_v18).length), p0, p0, globalState), new BigNumber((_1459_v17).length))];
        let _rhs244 = ((_1438_v0).Difference(_1438_v0)).Difference(_1438_v0);
        let _lhs205 = globalState;
        let _lhs206 = globalState;
        let _lhs207 = _1458_v16;
        let _lhs208 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1458_v16).length));
        _lhs205.f2 = _rhs241;
        _lhs206.f2 = _rhs242;
        _lhs207[_lhs208] = _rhs243;
        _1438_v0 = _rhs244;
        let _1461_v19;
        let _nw240 = Array((new BigNumber(24)).toNumber()).fill(false);
        _1461_v19 = _nw240;
        let _index225 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1461_v19).length));
        (_1461_v19)[_index225] = p0;
        let _1462_v20;
        _1462_v20 = _module.D0.create_DC1(p0, new BigNumber(-422), _this.f27, p0, p1);
        let _1463_v21;
        _1463_v21 = _dafny.Seq.of(_module.__default.fm35(p1, (_this).f18, globalState), function (_pat_let35_0) {
          return function (_1464_dt__update__tmp_h0) {
            return function (_pat_let36_0) {
              return function (_1465_dt__update_hcf4_h0) {
                return function (_pat_let37_0) {
                  return function (_1466_dt__update_hcf1_h0) {
                    return _module.D0.create_DC1(_1466_dt__update_hcf1_h0, (_1464_dt__update__tmp_h0).dtor_cf2, (_1464_dt__update__tmp_h0).dtor_cf3, _1465_dt__update_hcf4_h0, (_1464_dt__update__tmp_h0).dtor_cf5);
                  }(_pat_let37_0);
                }((_this).f18);
              }(_pat_let36_0);
            }(false);
          }(_pat_let35_0);
        }(_1462_v20));
        let _1467_v22;
        _1467_v22 = _dafny.Seq.of(_1463_v21);
        let _1468_v23;
        let _nw241 = Array((new BigNumber(7)).toNumber());
        _nw241[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_1463_v21, _module.__default.safeIndex(p1, new BigNumber((_1463_v21).length)), _1462_v20);
        _nw241[(_dafny.ONE).toNumber()] = _1463_v21;
        _nw241[(new BigNumber(2)).toNumber()] = _module.__default.fm53(globalState);
        _nw241[(new BigNumber(3)).toNumber()] = _1463_v21;
        _nw241[(new BigNumber(4)).toNumber()] = (_1467_v22)[_module.__default.safeIndex(p1, new BigNumber((_1467_v22).length))];
        _nw241[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1463_v21, _1463_v21);
        _nw241[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_1463_v21, _module.__default.safeIndex(p1, new BigNumber((_1463_v21).length)), _1462_v20);
        _1468_v23 = _nw241;
        let _index226 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1468_v23).length));
        (_1468_v23)[_index226] = _module.__default.fm53(globalState);
        let _1469_v24;
        _1469_v24 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1470_v25;
        _1470_v25 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wbmoton"));
        let _1471_v26;
        _1471_v26 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
        let _1472_v27;
        _1472_v27 = _dafny.MultiSet.fromElements(_1469_v24);
        let _1473_v28;
        _1473_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1472_v27);
        let _1474_v29;
        _1474_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(((((_1473_v28).contains(p0)) ? ((_1473_v28).get(p0)) : (_1472_v27))).cardinality()));
        let _1475_v30;
        _1475_v30 = _dafny.Seq.of(_1474_v29);
        let _1476_v31;
        let _nw242 = new _module.C4();
        _nw242.__ctor(_1469_v24, new BigNumber((_dafny.MultiSet.FromArray(_1460_v18)).cardinality()), (_this).f18, _dafny.Seq.update((_1470_v25)[_module.__default.safeIndex((((_1471_v26).contains((_this).f18)) ? ((_1471_v26).get((_this).f18)) : (p1)), new BigNumber((_1470_v25).length))], _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1474_v29).length),p0)).length), new BigNumber(((_1470_v25)[_module.__default.safeIndex((((_1471_v26).contains((_this).f18)) ? ((_1471_v26).get((_this).f18)) : (p1)), new BigNumber((_1470_v25).length))]).length)), _1469_v24), _1475_v30);
        _1476_v31 = _nw242;
        let _1477_v32;
        _1477_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1476_v31,p0);
        let _index227 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1461_v19).length));
        let _index228 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1468_v23).length));
        let _rhs245 = (new BigNumber((_module.__default.fm16((_1476_v31).f30, (_this).f18, (_1476_v31).f29, (_1476_v31).f30, globalState)).length)).isLessThan(new BigNumber((_1477_v32).length));
        let _rhs246 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), ((_1478_p0, _1479_p1, _1480_v31) => function (_1481_i2) {
          return function (_pat_let38_0) {
            return function (_1482_dt__update__tmp_h1) {
              return function (_pat_let39_0) {
                return function (_1483_dt__update_hcf4_h1) {
                  return function (_pat_let40_0) {
                    return function (_1484_dt__update_hcf2_h0) {
                      return _module.D0.create_DC1((_1482_dt__update__tmp_h1).dtor_cf1, _1484_dt__update_hcf2_h0, (_1482_dt__update__tmp_h1).dtor_cf3, _1483_dt__update_hcf4_h1, (_1482_dt__update__tmp_h1).dtor_cf5);
                    }(_pat_let40_0);
                  }((_dafny.ZERO).minus((_1480_v31).f30));
                }(_pat_let39_0);
              }(_1478_p0);
            }(_pat_let38_0);
          }(_module.D0.create_DC1((_this).f18, new BigNumber(484), (_this).f28, _1478_p0, _1479_p1));
        })(p0, p1, _1476_v31)), _1463_v21);
        let _lhs209 = _1461_v19;
        let _lhs210 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1461_v19).length));
        let _lhs211 = _1468_v23;
        let _lhs212 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1468_v23).length));
        _lhs209[_lhs210] = _rhs245;
        _lhs211[_lhs212] = _rhs246;
        let _1485_v33;
        let _nw243 = new _module.C0();
        _nw243.__ctor(new BigNumber(((_1458_v16)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_1458_v16).length))]).length), (_1476_v31).f30);
        _1485_v33 = _nw243;
        let _1486_v34;
        _1486_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1461_v19,p1);
        (globalState).f15 = _module.__default.safeModuloInt((_1485_v33).f32, new BigNumber((_1486_v34).length));
        (globalState).f3 = _module.__default.safeModuloInt((_1485_v33).f32, p1);
      }
      let _1487_v35;
      let _nw244 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _1487_v35 = _nw244;
      let _1488_v36;
      _1488_v36 = _dafny.MultiSet.fromElements(_1487_v35);
      let _1489_v37;
      _1489_v37 = _dafny.Seq.of(_1487_v35);
      let _hi7 = _module.__default.safeModuloInt(p1, p1);
      for (let _1490_i3 = (((_1488_v36).contains((_1489_v37)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_1489_v37).length))])) ? ((_1488_v36).get((_1489_v37)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_1489_v37).length))])) : (p1)); _1490_i3.isLessThan(_hi7); _1490_i3 = _1490_i3.plus(_dafny.ONE)) {
        let _1491_v38;
        let _nw245 = Array((new BigNumber(23)).toNumber()).fill(false);
        _1491_v38 = _nw245;
        let _index229 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_1491_v38).length));
        (_1491_v38)[_index229] = p0;
        let _index230 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_1491_v38).length));
        let _rhs247 = _dafny.Seq.Concat((_this).f28, _dafny.Seq.Concat((_this).f28, (_this).f28));
        let _rhs248 = !((_1439_v1).equals(_1439_v1)) || (_dafny.Seq.IsProperPrefixOf(_this.f27, (_this).f28));
        let _lhs213 = globalState;
        let _lhs214 = _1491_v38;
        let _lhs215 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_1491_v38).length));
        _lhs213.f12 = _rhs247;
        _lhs214[_lhs215] = _rhs248;
        let _1492_v39;
        let _nw246 = Array((new BigNumber(22)).toNumber()).fill([]);
        _1492_v39 = _nw246;
        let _1493_v40;
        _1493_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1492_v39,_module.__default.fm48(globalState));
        _1493_v40 = (_1493_v40).update(_1492_v39, (_this).f28);
        let _1494_v41;
        _1494_v41 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("inthwlmb"), (_this).f28, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xgf"), _this.f27));
        _1494_v41 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(769)), function (_1495_i4) {
          return (_this).f28;
        });
        let _1496_v42;
        _1496_v42 = new _dafny.CodePoint('w'.codePointAt(0));
        _1496_v42 = _module.__default.fm46(globalState);
      }
      let _1497_v43;
      _1497_v43 = _1439_v1;
      let _1498_i5;
      _1498_i5 = _dafny.ZERO;
      L11: {
        while ((((_1497_v43)).Difference(_1439_v1)).IsDisjointFrom(_1439_v1)) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1498_i5)) {
              break L11;
            }
            _1498_i5 = (_1498_i5).plus(_dafny.ONE);
            (globalState).f3 = new BigNumber(188);
            if ((_this).f18) {
              (globalState).f0 = (_dafny.ZERO).minus((p1).multipliedBy(p1));
              let _1499_v44;
              let _nw247 = Array((new BigNumber(2)).toNumber()).fill(_dafny.MultiSet.Empty);
              _1499_v44 = _nw247;
              _1499_v44 = _1499_v44;
              (_this).f27 = (_this).f28;
              (globalState).f3 = p1;
              let _1500_v45;
              let _1501_v46;
              let _1502_v47;
              let _1503_v48;
              let _out40;
              let _out41;
              let _out42;
              let _out43;
              let _outcollector10 = _module.__default.m0(globalState);
              _out40 = _outcollector10[0];
              _out41 = _outcollector10[1];
              _out42 = _outcollector10[2];
              _out43 = _outcollector10[3];
              _1500_v45 = _out40;
              _1501_v46 = _out41;
              _1502_v47 = _out42;
              _1503_v48 = _out43;
            } else {
              let _1504_v49;
              let _nw248 = new _module.C0();
              _nw248.__ctor(p1, _module.__default.safeDivisionInt(p1, p1));
              _1504_v49 = _nw248;
              (globalState).f3 = (_1504_v49).f31;
              (globalState).f0 = p1;
              let _1505_v50;
              let _nw249 = Array((new BigNumber(3)).toNumber()).fill(_module.D2.Default());
              _1505_v50 = _nw249;
              _1505_v50 = (((_this).f18) ? (_1505_v50) : (_1505_v50));
              (globalState).f3 = new BigNumber(511);
            }
            let _1506_v51;
            let _nw250 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
            _1506_v51 = _nw250;
            let _index231 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1506_v51).length));
            (_1506_v51)[_index231] = _1439_v1;
            let _1507_v52;
            let _nw251 = Array((new BigNumber(26)).toNumber()).fill([]);
            _1507_v52 = _nw251;
            let _1508_v53;
            let _nw252 = Array((new BigNumber(20)).toNumber()).fill(false);
            _1508_v53 = _nw252;
            let _1509_v54;
            _1509_v54 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_1508_v53);
            let _1510_v55;
            _1510_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
            let _index232 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1506_v51).length));
            let _rhs249 = new BigNumber((_module.__default.fm36(!(!((_this).f18) || ((_this).f18)), p1, ((_dafny.ZERO).minus(p1)).isLessThan(new BigNumber(447)), globalState)).length);
            let _rhs250 = (_1439_v1).Difference(_dafny.MultiSet.fromElements((((_1510_v55).contains(p0)) ? ((_1510_v55).get(p0)) : ((_this).f18))));
            let _rhs251 = _1507_v52;
            let _rhs252 = _1440_v2;
            let _rhs253 = (_1509_v54).update((_this).f18, _1508_v53);
            let _lhs216 = globalState;
            let _lhs217 = _1506_v51;
            let _lhs218 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1506_v51).length));
            _lhs216.f3 = _rhs249;
            _lhs217[_lhs218] = _rhs250;
            _1507_v52 = _rhs251;
            _1440_v2 = _rhs252;
            _1509_v54 = _rhs253;
            let _1511_v56;
            _1511_v56 = _dafny.Seq.of(p1);
            let _1512_v57;
            _1512_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1511_v56).length),_1438_v0);
            (globalState).f0 = new BigNumber(((_1512_v57).Merge(_1512_v57)).length);
          }
        }
      }
      let _hi8 = p1;
      for (let _1513_i6 = p1; _1513_i6.isLessThan(_hi8); _1513_i6 = _1513_i6.plus(_dafny.ONE)) {
        let _1514_v58;
        let _nw253 = new _module.C1();
        _nw253.__ctor((_this).f18);
        _1514_v58 = _nw253;
        let _1515_v59;
        _1515_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1513_i6,p1);
        let _1516_v60;
        _1516_v60 = _module.D13.create_DC33(p0, (((_1515_v59).contains(p1)) ? ((_1515_v59).get(p1)) : (new BigNumber(-363))), (_this).f18, p1);
        let _source13 = _1516_v60;
        if (_source13.is_DC33) {
          let _1517___mcc_h0 = (_source13).cf55;
          let _1518___mcc_h1 = (_source13).cf56;
          let _1519___mcc_h2 = (_source13).cf57;
          let _1520___mcc_h3 = (_source13).cf58;
          let _1521_cf58 = _1520___mcc_h3;
          let _1522_cf57 = _1519___mcc_h2;
          let _1523_cf56 = _1518___mcc_h1;
          let _1524_cf55 = _1517___mcc_h0;
          let _1525_v61;
          let _nw254 = Array((new BigNumber(3)).toNumber()).fill([]);
          _1525_v61 = _nw254;
          _1525_v61 = _1525_v61;
          let _index233 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1487_v35).length));
          (_1487_v35)[_index233] = _1523_cf56;
          let _1526_v62;
          let _nw255 = Array((_dafny.ONE).toNumber()).fill(false);
          _1526_v62 = _nw255;
          let _1527_v63;
          _1527_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(611),_1526_v62);
          let _1528_v64;
          _1528_v64 = _dafny.Seq.of(_1523_cf56);
          let _index234 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1487_v35).length));
          let _rhs254 = new BigNumber((_module.__default.fm16(_1521_cf58, (_this).f18, new _dafny.CodePoint('o'.codePointAt(0)), (_dafny.ZERO).minus(_1523_cf56), globalState)).length);
          let _rhs255 = _1523_cf56;
          let _rhs256 = p1;
          let _rhs257 = (_dafny.Map.Empty.slice().updateUnsafe((_1528_v64)[_module.__default.safeIndex(new BigNumber((_1438_v0).length), new BigNumber((_1528_v64).length))],_1526_v62)).update(_1523_cf56, _1526_v62);
          let _lhs219 = _1487_v35;
          let _lhs220 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1487_v35).length));
          _lhs219[_lhs220] = _rhs254;
          _1521_cf58 = _rhs255;
          _1521_cf58 = _rhs256;
          _1527_v63 = _rhs257;
          let _1529_v65;
          _1529_v65 = _dafny.Set.fromElements(p0, (_this).f18);
          let _1530_v66;
          _1530_v66 = _module.D10.create_DC24((((_this).f18) ? (_1529_v65) : (_1529_v65)));
          let _pat_let_tv23 = _1529_v65;
          let _pat_let_tv24 = _1529_v65;
          _1530_v66 = function (_pat_let41_0) {
            return function (_1533_dt__update__tmp_h3) {
              return function (_pat_let44_0) {
                return function (_1534_dt__update_hcf47_h1) {
                  return _module.D10.create_DC24(_1534_dt__update_hcf47_h1);
                }(_pat_let44_0);
              }(_pat_let_tv24);
            }(_pat_let41_0);
          }(function (_pat_let42_0) {
            return function (_1531_dt__update__tmp_h2) {
              return function (_pat_let43_0) {
                return function (_1532_dt__update_hcf47_h0) {
                  return _module.D10.create_DC24(_1532_dt__update_hcf47_h0);
                }(_pat_let43_0);
              }(_pat_let_tv23);
            }(_pat_let42_0);
          }(_1530_v66));
          (globalState).f12 = _this.f27;
        } else if (_source13.is_DC34) {
          let _1535___mcc_h4 = (_source13).cf59;
          let _1536___mcc_h5 = (_source13).cf60;
          let _1537___mcc_h6 = (_source13).cf61;
          let _1538___mcc_h7 = (_source13).cf62;
          let _1539_cf62 = _1538___mcc_h7;
          let _1540_cf61 = _1537___mcc_h6;
          let _1541_cf60 = _1536___mcc_h5;
          let _1542_cf59 = _1535___mcc_h4;
          (globalState).f15 = _module.__default.safeDivisionInt((_1513_i6).plus(p1), (_dafny.ZERO).minus(_1513_i6));
          let _1543_v67;
          let _1544_v68;
          let _1545_v69;
          let _1546_v70;
          let _out44;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector11 = _module.__default.m0(globalState);
          _out44 = _outcollector11[0];
          _out45 = _outcollector11[1];
          _out46 = _outcollector11[2];
          _out47 = _outcollector11[3];
          _1543_v67 = _out44;
          _1544_v68 = _out45;
          _1545_v69 = _out46;
          _1546_v70 = _out47;
          let _1547_v71;
          _1547_v71 = _module.D0.create_DC1((_this).f18, new BigNumber((_dafny.Seq.of(_1539_cf62, (_this).f18)).length), _this.f27, (_this).f18, new BigNumber((_1515_v59).length));
          let _index235 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1545_v69).length));
          (_1545_v69)[_index235] = (_1543_v67).plus((_1547_v71).dtor_cf5);
          let _index236 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1545_v69).length));
          (_1545_v69)[_index236] = _1543_v67;
          _1440_v2 = _dafny.Seq.Concat(_dafny.Seq.update(_1440_v2, _module.__default.safeIndex(_1543_v67, new BigNumber((_1440_v2).length)), !(_1539_cf62)), _1440_v2);
        } else if (_source13.is_DC32) {
          let _1548___mcc_h8 = (_source13).cf54;
          let _1549_cf54 = _1548___mcc_h8;
          let _1550_v72;
          let _nw256 = new _module.C1();
          _nw256.__ctor(!(false));
          _1550_v72 = _nw256;
          let _1551_v73;
          _1551_v73 = _dafny.Seq.of(_1513_i6);
          let _1552_v75;
          _1552_v75 = _dafny.Set.fromElements(p0, (_this).f18);
          let _1553_v78;
          _1553_v78 = _module.D0.create_DC1(p0, _1513_i6, _this.f27, false, _1513_i6);
          let _1554_v79;
          _1554_v79 = _dafny.MultiSet.fromElements(p1);
          let _1555_v80;
          let _nw257 = Array((new BigNumber(27)).toNumber());
          _nw257[(_dafny.ZERO).toNumber()] = (_this).f18;
          _nw257[(_dafny.ONE).toNumber()] = p0;
          _nw257[(new BigNumber(2)).toNumber()] = p0;
          _nw257[(new BigNumber(3)).toNumber()] = _dafny.areEqual((_this).f28, _this.f27);
          _nw257[(new BigNumber(4)).toNumber()] = (function () {
            let _coll45 = new _dafny.Set();
            for (const _compr_45 of (_1551_v73).Elements) {
              let _1556_v74 = _compr_45;
              if (_dafny.Seq.contains(_1551_v73, _1556_v74)) {
                _coll45.add(_module.__default.safeModuloInt(_1556_v74, new BigNumber(742)));
              }
            }
            return _coll45;
          }()).IsDisjointFrom(_1438_v0);
          _nw257[(new BigNumber(5)).toNumber()] = (_this).f18;
          _nw257[(new BigNumber(6)).toNumber()] = (_this).f18;
          _nw257[(new BigNumber(7)).toNumber()] = (_1552_v75).equals(_1552_v75);
          _nw257[(new BigNumber(8)).toNumber()] = p0;
          _nw257[(new BigNumber(9)).toNumber()] = _module.__default.fm23(function () {
            let _coll46 = new _dafny.Set();
            for (const _compr_46 of (_1438_v0).Elements) {
              let _1557_v76 = _compr_46;
              if ((_1438_v0).contains(_1557_v76)) {
                _coll46.add(_module.__default.safeModuloInt(_1557_v76, new BigNumber((_dafny.Seq.UnicodeFromString("gm")).length)));
              }
            }
            return _coll46;
          }(), (_1551_v73)[_module.__default.safeIndex(_1513_i6, new BigNumber((_1551_v73).length))], globalState);
          _nw257[(new BigNumber(10)).toNumber()] = p0;
          _nw257[(new BigNumber(11)).toNumber()] = (_this).f18;
          _nw257[(new BigNumber(12)).toNumber()] = (_this).f18;
          _nw257[(new BigNumber(13)).toNumber()] = false;
          _nw257[(new BigNumber(14)).toNumber()] = _module.__default.fm23(function () {
            let _coll47 = new _dafny.Set();
            for (const _compr_47 of _dafny.IntegerRange(new BigNumber(868), new BigNumber(822))) {
              let _1558_v77 = _compr_47;
              if (((new BigNumber(868)).isLessThanOrEqualTo(_1558_v77)) && ((_1558_v77).isLessThan(new BigNumber(822)))) {
                _coll47.add(_module.__default.safeModuloInt(_1558_v77, _1513_i6));
              }
            }
            return _coll47;
          }(), p1, globalState);
          _nw257[(new BigNumber(15)).toNumber()] = (_this).fm15(new BigNumber(179), _1513_i6, globalState);
          _nw257[(new BigNumber(16)).toNumber()] = p0;
          _nw257[(new BigNumber(17)).toNumber()] = _module.__default.fm23(_1438_v0, new BigNumber((_this.f27).length), globalState);
          _nw257[(new BigNumber(18)).toNumber()] = p0;
          _nw257[(new BigNumber(19)).toNumber()] = ((_1550_v72).fm22(_dafny.Seq.UnicodeFromString("bm"), _1553_v78, globalState)).isLessThan(new BigNumber((_1551_v73).length));
          _nw257[(new BigNumber(20)).toNumber()] = p0;
          _nw257[(new BigNumber(21)).toNumber()] = (_this).f18;
          _nw257[(new BigNumber(22)).toNumber()] = (_this).f18;
          _nw257[(new BigNumber(23)).toNumber()] = (((_this).f18) ? (p0) : (!((_this).f18)));
          _nw257[(new BigNumber(24)).toNumber()] = _module.__default.fm23(_1438_v0, (_dafny.ZERO).minus(p1), globalState);
          _nw257[(new BigNumber(25)).toNumber()] = (_module.__default.fm25(p1, globalState)).IsDisjointFrom(_1554_v79);
          _nw257[(new BigNumber(26)).toNumber()] = true;
          _1555_v80 = _nw257;
          let _index237 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_1555_v80).length));
          (_1555_v80)[_index237] = p0;
          let _1559_v81;
          _1559_v81 = _module.D9.create_DC23(p0, (_this).f18, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-102)), ((_1560_p1) => function (_1561_i7) {
  return _1560_p1;
})(p1))).length));
          let _1562_v82;
          _1562_v82 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_this.f27).length));
          let _index238 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_1555_v80).length));
          let _rhs258 = (_this).f18;
          let _rhs259 = (((p0) ? (_1559_v81) : (_1559_v81))).dtor_cf44;
          let _rhs260 = _1439_v1;
          let _rhs261 = (p1).multipliedBy(p1);
          let _rhs262 = (((_this).f18) ? ((_1513_i6).isLessThanOrEqualTo((((_1562_v82).contains(p0)) ? ((_1562_v82).get(p0)) : (p1)))) : ((_this).f18));
          let _lhs221 = globalState;
          let _lhs222 = globalState;
          let _lhs223 = globalState;
          let _lhs224 = _1555_v80;
          let _lhs225 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_1555_v80).length));
          _lhs221.f2 = _rhs258;
          _lhs222.f2 = _rhs259;
          _1439_v1 = _rhs260;
          _lhs223.f0 = _rhs261;
          _lhs224[_lhs225] = _rhs262;
          (globalState).f2 = p0;
          (globalState).f3 = (((_this).fm15(_1513_i6, _1513_i6, globalState)) ? (_1513_i6) : (p1));
        } else {
          let _1563___mcc_h9 = (_source13).cf63;
          let _1564_cf63 = _1563___mcc_h9;
          let _1565_v83;
          let _nw258 = Array((new BigNumber(17)).toNumber()).fill(_module.D0.Default());
          _1565_v83 = _nw258;
          let _1566_v84;
          _1566_v84 = _module.D0.create_DC0(p0);
          let _index239 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1565_v83).length));
          (_1565_v83)[_index239] = ((false) ? (_1566_v84) : (_1566_v84));
          let _1567_v85;
          _1567_v85 = new _dafny.CodePoint('j'.codePointAt(0));
          let _1568_v86;
          _1568_v86 = _dafny.Seq.of(p1);
          let _index240 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1565_v83).length));
          let _rhs263 = !(p0) || ((_this).f18);
          let _rhs264 = (new BigNumber((_module.__default.fm44(p1, p0, _1567_v85, _1513_i6, globalState)).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_1568_v86, _1568_v86)).length));
          let _rhs265 = function (_pat_let45_0) {
            return function (_1569_dt__update__tmp_h4) {
              return function (_pat_let46_0) {
                return function (_1570_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_1570_dt__update_hcf0_h0);
                }(_pat_let46_0);
              }(true);
            }(_pat_let45_0);
          }(_1566_v84);
          let _rhs266 = p0;
          let _lhs226 = globalState;
          let _lhs227 = globalState;
          let _lhs228 = _1565_v83;
          let _lhs229 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1565_v83).length));
          let _lhs230 = globalState;
          _lhs226.f2 = _rhs263;
          _lhs227.f2 = _rhs264;
          _lhs228[_lhs229] = _rhs265;
          _lhs230.f2 = _rhs266;
          let _1571_v87;
          _1571_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1567_v85,p0);
          (globalState).f15 = (new BigNumber(450)).plus(new BigNumber(((_1571_v87).Merge(_1571_v87)).length));
          let _1572_v88;
          _1572_v88 = _dafny.Seq.of(_1515_v59, _1515_v59);
          let _1573_v89;
          let _nw259 = new _module.C3();
          _nw259.__ctor(_1568_v86, _dafny.Seq.Concat(_this.f27, _this.f27), _1572_v88, (p1).isEqualTo(_1513_i6));
          _1573_v89 = _nw259;
          let _1574_v90;
          _1574_v90 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(926)), function (_1575_i8) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }));
          let _1576_v91;
          let _nw260 = new _module.C4();
          _nw260.__ctor(_1567_v85, p1, false, _dafny.Seq.Concat(_this.f27, (_1574_v90)[_module.__default.safeIndex(_1513_i6, new BigNumber((_1574_v90).length))]), _dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), ((_1577_v59) => function (_1578_i9) {
            return _1577_v59;
          })(_1515_v59)));
          _1576_v91 = _nw260;
        }
        (globalState).f3 = p1;
        let _1579_v92;
        _1579_v92 = _dafny.Seq.of(p1, new BigNumber(571), (((_1515_v59).contains(p1)) ? ((_1515_v59).get(p1)) : (_1513_i6)), new BigNumber((_module.__default.fm54(_1513_i6, p1, globalState)).length));
        (globalState).f2 = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-129)), function (_1580_i10) {
          return new BigNumber((_this.f27).length);
        }), _1579_v92));
      }
      let _source14 = _1497_v43;
      let _1581___mcc_h10 = _source14;
      let _1582_cf65 = _1581___mcc_h10;
      let _1583_v93;
      _1583_v93 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1584_v94;
      _1584_v94 = _dafny.MultiSet.fromElements(_dafny.Seq.update((_this).f28, _module.__default.safeIndex(new BigNumber((_this.f27).length), new BigNumber(((_this).f28).length)), _1583_v93), _this.f27, _dafny.Seq.UnicodeFromString("mnrphd"), (_this).f28, (_this).f28);
      (globalState).f0 = _module.__default.safeModuloInt(p1, (((_1584_v94).contains((_this).f28)) ? ((_1584_v94).get((_this).f28)) : (p1)));
      let _index241 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_1487_v35).length));
      (_1487_v35)[_index241] = (p1).plus(p1);
      let _index242 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_1487_v35).length));
      (_1487_v35)[_index242] = p1;
      let _1585_v95;
      let _nw261 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
      _1585_v95 = _nw261;
      let _1586_v96;
      _1586_v96 = _dafny.Seq.of((_1487_v35)[_module.__default.safeIndex(new BigNumber(49), new BigNumber((_1487_v35).length))], new BigNumber((_dafny.MultiSet.fromElements(p0, (_this).f18)).cardinality()), (_1487_v35)[_module.__default.safeIndex(new BigNumber(49), new BigNumber((_1487_v35).length))]);
      let _index243 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1585_v95).length));
      (_1585_v95)[_index243] = _1586_v96;
      let _index244 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1585_v95).length));
      let _rhs267 = !(p0);
      let _rhs268 = _dafny.Seq.Concat(_1586_v96, _1586_v96);
      let _rhs269 = p0;
      let _lhs231 = globalState;
      let _lhs232 = _1585_v95;
      let _lhs233 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1585_v95).length));
      let _lhs234 = globalState;
      _lhs231.f2 = _rhs267;
      _lhs232[_lhs233] = _rhs268;
      _lhs234.f2 = _rhs269;
      let _1587_v97;
      let _nw262 = Array((new BigNumber(8)).toNumber()).fill(_module.D13.Default());
      _1587_v97 = _nw262;
      let _1588_v98;
      let _nw263 = Array((new BigNumber(25)).toNumber()).fill(false);
      _1588_v98 = _nw263;
      let _index245 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1588_v98).length));
      (_1588_v98)[_index245] = p0;
      let _index246 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1588_v98).length));
      let _rhs270 = _1587_v97;
      let _rhs271 = p0;
      let _rhs272 = p0;
      let _rhs273 = p1;
      let _lhs235 = _1588_v98;
      let _lhs236 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1588_v98).length));
      let _lhs237 = globalState;
      let _lhs238 = globalState;
      _1587_v97 = _rhs270;
      _lhs235[_lhs236] = _rhs271;
      _lhs237.f2 = _rhs272;
      _lhs238.f15 = _rhs273;
      return;
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f18 = false;
      this._f19 = _dafny.Seq.UnicodeFromString("");
      this._f20 = _dafny.Seq.of();
      this.f26 = _dafny.ZERO;
      this._f25 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f25, f26, f18, f19, f20) {
      let _this = this;
      (_this)._f25 = f25;
      (_this).f26 = f26;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      return;
    }
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm3(p0, globalState) {
      let _this = this;
      return ((_this).f18) && ((((_this).f18) ? ((_this).f18) : ((_this).f18)));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _1589_v0;
      _1589_v0 = _dafny.Set.fromElements(p1);
      (globalState).f2 = (_1589_v0).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber(30), new BigNumber(-321)));
      let _1590_i0;
      _1590_i0 = _dafny.ZERO;
      L12: {
        while (((((_this).f18) ? (new BigNumber(431)) : (p1))).isLessThan(p1)) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1590_i0)) {
              break L12;
            }
            _1590_i0 = (_1590_i0).plus(_dafny.ONE);
            let _1591_v1;
            _1591_v1 = _dafny.MultiSet.fromElements((_this).f18);
            let _1592_v2;
            _1592_v2 = _module.D0.create_DC1((_this).f18, new BigNumber((p0).length), _dafny.Seq.Concat((_this).f19, p0), ((_this).fm2((_this).f18, (_this).f18, (_module.D2.create_DC8((_this).f18, (_this).f18, (_this).f18)).dtor_cf18, globalState)) === ((_this).f18), (_dafny.ZERO).minus(new BigNumber((_1591_v1).cardinality())));
            let _source15 = _1592_v2;
            if (_source15.is_DC1) {
              let _1593___mcc_h0 = (_source15).cf1;
              let _1594___mcc_h1 = (_source15).cf2;
              let _1595___mcc_h2 = (_source15).cf3;
              let _1596___mcc_h3 = (_source15).cf4;
              let _1597___mcc_h4 = (_source15).cf5;
              let _1598_cf5 = _1597___mcc_h4;
              let _1599_cf4 = _1596___mcc_h3;
              let _1600_cf3 = _1595___mcc_h2;
              let _1601_cf2 = _1594___mcc_h1;
              let _1602_cf1 = _1593___mcc_h0;
              let _1603_v3;
              _1603_v3 = _dafny.MultiSet.fromElements(_1601_cf2);
              let _1604_v4;
              _1604_v4 = _dafny.Seq.of(_1598_cf5);
              let _1605_v5;
              _1605_v5 = new _dafny.CodePoint('f'.codePointAt(0));
              let _1606_v6;
              let _nw264 = Array((new BigNumber(14)).toNumber());
              _nw264[(_dafny.ZERO).toNumber()] = (_this.f26).multipliedBy(_module.__default.fm1(p1, _1602_cf1, _1602_cf1, globalState));
              _nw264[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1603_v3).cardinality()), _module.__default.fm1((_dafny.ZERO).minus(_this.f26), (_this).f18, (_this).f18, globalState));
              _nw264[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(p1), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1591_v1).cardinality()))).cardinality())));
              _nw264[(new BigNumber(3)).toNumber()] = _1601_cf2;
              _nw264[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), function (_1607_i1) {
                return new _dafny.CodePoint('k'.codePointAt(0));
              }), _module.__default.fm13(_1599_cf4, p1, globalState)), _module.__default.safeIndex((_1604_v4)[_module.__default.safeIndex(_this.f26, new BigNumber((_1604_v4).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), function (_1608_i1) {
                return new _dafny.CodePoint('k'.codePointAt(0));
              }), _module.__default.fm13(_1599_cf4, p1, globalState))).length)), _1605_v5)).length);
              _nw264[(new BigNumber(5)).toNumber()] = p1;
              _nw264[(new BigNumber(6)).toNumber()] = (((_1591_v1).contains(_1599_cf4)) ? ((_1591_v1).get(_1599_cf4)) : (p1));
              _nw264[(new BigNumber(7)).toNumber()] = _1601_cf2;
              _nw264[(new BigNumber(8)).toNumber()] = ((_1602_cf1) ? (_1601_cf2) : ((_dafny.ZERO).minus(_1598_cf5)));
              _nw264[(new BigNumber(9)).toNumber()] = _1598_cf5;
              _nw264[(new BigNumber(10)).toNumber()] = _this.f26;
              _nw264[(new BigNumber(11)).toNumber()] = new BigNumber(((_this).f25).length);
              _nw264[(new BigNumber(12)).toNumber()] = _1598_cf5;
              _nw264[(new BigNumber(13)).toNumber()] = _this.f26;
              _1606_v6 = _nw264;
              let _index247 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1606_v6).length));
              (_1606_v6)[_index247] = (_dafny.ZERO).minus((_1598_cf5).plus(_this.f26));
              let _index248 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1606_v6).length));
              (_1606_v6)[_index248] = _this.f26;
              _1603_v3 = _1603_v3;
              let _1609_v7;
              let _nw265 = new _module.C2();
              _nw265.__ctor(new BigNumber(170), _1604_v4, (_this).f18);
              _1609_v7 = _nw265;
              let _1610_v8;
              let _nw266 = Array((new BigNumber(29)).toNumber()).fill(false);
              _1610_v8 = _nw266;
              let _1611_v9;
              _1611_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1609_v7,_1610_v8);
              let _1612_v10;
              _1612_v10 = _dafny.Seq.of((((_1611_v9).contains(_1609_v7)) ? ((_1611_v9).get(_1609_v7)) : (_1610_v8)), _1610_v8);
              let _1613_v11;
              _1613_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1599_cf4,_1612_v10);
              let _1614_v12;
              let _nw267 = Array((new BigNumber(20)).toNumber());
              _nw267[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1612_v10, _1612_v10);
              _nw267[(_dafny.ONE).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1612_v10, _1612_v10);
              _nw267[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1612_v10, _dafny.Seq.update(_1612_v10, _module.__default.safeIndex((_1606_v6)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_1606_v6).length))], new BigNumber((_1612_v10).length)), _1610_v8));
              _nw267[(new BigNumber(4)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(5)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(6)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(7)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(8)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(9)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(10)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_1612_v10, _1612_v10);
              _nw267[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat((((_1613_v11).contains(_1602_cf1)) ? ((_1613_v11).get(_1602_cf1)) : (_1612_v10)), _dafny.Seq.update(_1612_v10, _module.__default.safeIndex(_this.f26, new BigNumber((_1612_v10).length)), _1610_v8));
              _nw267[(new BigNumber(13)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1612_v10, _1612_v10);
              _nw267[(new BigNumber(15)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(16)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_1612_v10, _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1612_v10).length)), _1610_v8);
              _nw267[(new BigNumber(18)).toNumber()] = _1612_v10;
              _nw267[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(_1610_v8);
              _1614_v12 = _nw267;
              let _index249 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1614_v12).length));
              (_1614_v12)[_index249] = _1612_v10;
              let _index250 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1614_v12).length));
              (_1614_v12)[_index250] = _1612_v10;
              let _1615_v13;
              let _nw268 = Array((new BigNumber(7)).toNumber());
              _nw268[(_dafny.ZERO).toNumber()] = (_this).f25;
              _nw268[(_dafny.ONE).toNumber()] = _module.__default.fm44((_1606_v6)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_1606_v6).length))], _1602_cf1, _1605_v5, new BigNumber((_dafny.Seq.of((_this).f18, _1599_cf4)).length), globalState);
              _nw268[(new BigNumber(2)).toNumber()] = (_this).f25;
              _nw268[(new BigNumber(3)).toNumber()] = (_this).f25;
              _nw268[(new BigNumber(4)).toNumber()] = (_this).f25;
              _nw268[(new BigNumber(5)).toNumber()] = (_this).f25;
              _nw268[(new BigNumber(6)).toNumber()] = (_this).f25;
              _1615_v13 = _nw268;
              let _1616_v14;
              _1616_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1610_v8,_1615_v13);
              _1616_v14 = (_1616_v14).update(_1610_v8, _1615_v13);
            } else {
              let _1617___mcc_h5 = (_source15).cf0;
              let _1618_cf0 = _1617___mcc_h5;
              (globalState).f15 = (_dafny.ZERO).minus(p1);
              let _1619_v15;
              let _nw269 = Array((new BigNumber(25)).toNumber()).fill(false);
              _1619_v15 = _nw269;
              let _1620_v16;
              let _nw270 = new _module.C1();
              _nw270.__ctor(_1618_cf0);
              _1620_v16 = _nw270;
              let _1621_v17;
              _1621_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1619_v15,_1620_v16);
              _1621_v17 = (_1621_v17).update(_1619_v15, _1620_v16);
              let _1622_v18;
              let _nw271 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
              _1622_v18 = _nw271;
              let _1623_v19;
              _1623_v19 = _dafny.Set.fromElements(_1618_cf0, _1618_cf0);
              let _rhs274 = p1;
              let _rhs275 = _module.__default.fm1(new BigNumber((_1623_v19).length), !(_1618_cf0) || ((_this).f18), (_this).f18, globalState);
              let _rhs276 = (_this).f18;
              let _rhs277 = _1622_v18;
              let _lhs239 = globalState;
              let _lhs240 = globalState;
              let _lhs241 = globalState;
              _lhs239.f15 = _rhs274;
              _lhs240.f0 = _rhs275;
              _lhs241.f2 = _rhs276;
              _1622_v18 = _rhs277;
              let _1624_v20;
              let _1625_v21;
              let _out48;
              let _out49;
              let _outcollector12 = (_this).m6(globalState);
              _out48 = _outcollector12[0];
              _out49 = _outcollector12[1];
              _1624_v20 = _out48;
              _1625_v21 = _out49;
            }
            let _1626_v22;
            _1626_v22 = _dafny.Set.fromElements(p2, (_this).f19);
            let _1627_v23;
            _1627_v23 = _module.D11.create_DC28((_this).f18);
            let _1628_v24;
            _1628_v24 = _dafny.Seq.of(_module.D11.create_DC29(_1627_v23), _1627_v23);
            let _1629_v25;
            _1629_v25 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (_1626_v22) : (_1626_v22)),_module.D11.create_DC29((_1628_v24)[_module.__default.safeIndex(p1, new BigNumber((_1628_v24).length))]));
            _1629_v25 = (_1629_v25).update((_1626_v22).Intersect(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("vfwjy"))), _module.D11.create_DC29(_1627_v23));
            let _1630_v26;
            _1630_v26 = _dafny.Seq.of(p1);
            let _1631_v27;
            let _nw272 = new _module.C3();
            _nw272.__ctor(_1630_v26, _dafny.Seq.Concat(p2, (_this).f19), _dafny.Seq.Concat((_this).f20, (_this).f20), (_this).f18);
            _1631_v27 = _nw272;
            let _1632_v28;
            _1632_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
            let _1633_v29;
            _1633_v29 = _dafny.Set.fromElements((_1632_v28).Merge(_1632_v28));
            _1633_v29 = (_1633_v29).Union(_1633_v29);
          }
        }
      }
      let _1634_v30;
      let _nw273 = new _module.C1();
      _nw273.__ctor(false);
      _1634_v30 = _nw273;
      let _1635_v31;
      _1635_v31 = _module.D11.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(_1634_v30,(_this).f18));
      let _1636_v32;
      _1636_v32 = _module.D11.create_DC29(_1635_v31);
      let _1637_v33;
      _1637_v33 = _module.D11.create_DC29(_1635_v31);
      let _source16 = _1637_v33;
      if (_source16.is_DC28) {
        let _1638___mcc_h6 = (_source16).cf51;
        let _1639_cf51 = _1638___mcc_h6;
        (globalState).f0 = new BigNumber(287);
        let _1640_v34;
        _1640_v34 = _dafny.MultiSet.fromElements(_this.f26);
        let _1641_v35;
        _1641_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1639_cf51,new BigNumber((_1640_v34).cardinality()));
        let _1642_v36;
        _1642_v36 = _dafny.Set.fromElements(_1641_v35);
        let _1643_v37;
        _1643_v37 = _dafny.Seq.of((_dafny.ZERO).minus(_this.f26), p1, (_dafny.ZERO).minus(_this.f26));
        let _rhs278 = _module.__default.safeDivisionInt(_this.f26, ((_1643_v37)[_module.__default.safeIndex(new BigNumber((_1640_v34).cardinality()), new BigNumber((_1643_v37).length))]).minus(p1));
        let _rhs279 = p1;
        let _rhs280 = _1642_v36;
        let _rhs281 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(424)), function (_1644_i2) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(839)), function (_1645_i3) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }));
        let _lhs242 = globalState;
        let _lhs243 = _this;
        let _lhs244 = globalState;
        _lhs242.f0 = _rhs278;
        _lhs243.f26 = _rhs279;
        _1642_v36 = _rhs280;
        _lhs244.f12 = _rhs281;
        if (!((_this).f18) || (!(false))) {
          let _1646_v38;
          _1646_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f19).length),new BigNumber((_1643_v37).length));
          let _1647_v39;
          let _nw274 = new _module.C2();
          _nw274.__ctor(new BigNumber(119), _1643_v37, _1639_cf51);
          _1647_v39 = _nw274;
          let _1648_v40;
          _1648_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1647_v39,false);
          _1646_v38 = (_1646_v38).update((_dafny.ZERO).minus(_this.f26), new BigNumber((_1648_v40).length));
          let _1649_v41;
          let _nw275 = Array((new BigNumber(17)).toNumber()).fill(false);
          _1649_v41 = _nw275;
          let _index251 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_1649_v41).length));
          (_1649_v41)[_index251] = (_this).f18;
          let _1650_v42;
          let _nw276 = new _module.C3();
          _nw276.__ctor(_dafny.Seq.of(new BigNumber(((_1647_v39).f34).length)), _dafny.Seq.UnicodeFromString("cutdr"), _dafny.Seq.Concat((_this).f20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(639)), ((_1651_v38, _1652_p1) => function (_1653_i4) {
            return (_1651_v38).update(_1652_p1, _this.f26);
          })(_1646_v38, p1))), ((_this).f25)[_module.__default.safeIndex(new BigNumber(327), new BigNumber(((_this).f25).length))]);
          _1650_v42 = _nw276;
          let _1654_v43;
          let _nw277 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
          _1654_v43 = _nw277;
          let _index252 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_1654_v43).length));
          (_1654_v43)[_index252] = _module.__default.fm49(new BigNumber(((_1650_v42).f19).length), _1647_v39.f33, p1, globalState);
          let _index253 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_1649_v41).length));
          let _index254 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_1654_v43).length));
          let _rhs282 = (_1650_v42).f18;
          let _rhs283 = (_1650_v42).f18;
          let _rhs284 = _1650_v42;
          let _rhs285 = _this.f26;
          let _rhs286 = _1589_v0;
          let _lhs245 = _1649_v41;
          let _lhs246 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_1649_v41).length));
          let _lhs247 = globalState;
          let _lhs248 = _1654_v43;
          let _lhs249 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_1654_v43).length));
          _1639_cf51 = _rhs282;
          _lhs245[_lhs246] = _rhs283;
          _1650_v42 = _rhs284;
          _lhs247.f3 = _rhs285;
          _lhs248[_lhs249] = _rhs286;
          (globalState).f0 = ((_this.f26).plus(_1647_v39.f33)).minus(_module.__default.fm1(new BigNumber(((_1654_v43)[_module.__default.safeIndex(new BigNumber(898), new BigNumber((_1654_v43).length))]).length), !(!((_1649_v41)[_module.__default.safeIndex(new BigNumber(957), new BigNumber((_1649_v41).length))])), (_1650_v42).f18, globalState));
          let _1655_v44;
          let _nw278 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1655_v44 = _nw278;
          let _1656_v45;
          _1656_v45 = new _dafny.CodePoint('p'.codePointAt(0));
          let _index255 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_1655_v44).length));
          (_1655_v44)[_index255] = _dafny.Seq.update((_this).f19, _module.__default.safeIndex(_1647_v39.f33, new BigNumber(((_this).f19).length)), _1656_v45);
          let _index256 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_1655_v44).length));
          (_1655_v44)[_index256] = (((_1650_v42).f18) ? ((_this).f19) : (p0));
          _1641_v35 = (_1641_v35).update(_1639_cf51, (_this.f26).minus(_this.f26));
        } else {
          let _1657_v46;
          let _init35 = ((_1658_cf51) => function (_1659_i5) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1658_cf51,new BigNumber(908));
          })(_1639_cf51);
          let _nw279 = Array((new BigNumber(11)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw279.length); _i0_35++) {
            _nw279[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1657_v46 = _nw279;
          let _index257 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1657_v46).length));
          (_1657_v46)[_index257] = _1641_v35;
          let _1660_v47;
          let _nw280 = Array((new BigNumber(9)).toNumber());
          _nw280[(_dafny.ZERO).toNumber()] = (_this).f18;
          _nw280[(_dafny.ONE).toNumber()] = _1639_cf51;
          _nw280[(new BigNumber(2)).toNumber()] = _1639_cf51;
          _nw280[(new BigNumber(3)).toNumber()] = (_this).f18;
          _nw280[(new BigNumber(4)).toNumber()] = _1639_cf51;
          _nw280[(new BigNumber(5)).toNumber()] = !(_1639_cf51);
          _nw280[(new BigNumber(6)).toNumber()] = _1639_cf51;
          _nw280[(new BigNumber(7)).toNumber()] = _1639_cf51;
          _nw280[(new BigNumber(8)).toNumber()] = (_this).f18;
          _1660_v47 = _nw280;
          let _1661_v48;
          _1661_v48 = _dafny.Map.Empty.slice().updateUnsafe(false,_1660_v47);
          let _1662_v49;
          _1662_v49 = _dafny.Map.Empty.slice().updateUnsafe((_1661_v48).Merge(_1661_v48),true);
          let _index258 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1657_v46).length));
          let _rhs287 = ((_this).f18) || (false);
          let _rhs288 = (((_1662_v49).contains(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1660_v47))) ? ((_1662_v49).get(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1660_v47))) : (_1639_cf51));
          let _rhs289 = _1639_cf51;
          let _rhs290 = (_1641_v35).Merge(_1641_v35);
          let _lhs250 = globalState;
          let _lhs251 = globalState;
          let _lhs252 = _1657_v46;
          let _lhs253 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1657_v46).length));
          _lhs250.f2 = _rhs287;
          _lhs251.f2 = _rhs288;
          _1639_cf51 = _rhs289;
          _lhs252[_lhs253] = _rhs290;
          (globalState).f15 = new BigNumber((_1641_v35).length);
          (globalState).f2 = (_this).f18;
          _1640_v34 = (_1640_v34).Intersect(_1640_v34);
          let _rhs291 = (_this).f18;
          let _rhs292 = !((_this).f18);
          let _rhs293 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-768))), new BigNumber(453))));
          let _lhs254 = globalState;
          let _lhs255 = globalState;
          let _lhs256 = globalState;
          _lhs254.f2 = _rhs291;
          _lhs255.f2 = _rhs292;
          _lhs256.f3 = _rhs293;
        }
        if (_1639_cf51) {
          let _1663_v50;
          let _init36 = function (_1664_i6) {
            return false;
          };
          let _nw281 = Array((new BigNumber(9)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw281.length); _i0_36++) {
            _nw281[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1663_v50 = _nw281;
          let _1665_v51;
          _1665_v51 = _dafny.Seq.of(_1663_v50);
          let _1666_v52;
          _1666_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1639_cf51,_1639_cf51);
          let _1667_v53;
          let _nw282 = Array((new BigNumber(26)).toNumber());
          _nw282[(_dafny.ZERO).toNumber()] = _1663_v50;
          _nw282[(_dafny.ONE).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(2)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(3)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(4)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(5)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(6)).toNumber()] = (_1665_v51)[_module.__default.safeIndex(new BigNumber(((_this).f25).length), new BigNumber((_1665_v51).length))];
          _nw282[(new BigNumber(7)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(8)).toNumber()] = (_1665_v51)[_module.__default.safeIndex(new BigNumber((_1666_v52).length), new BigNumber((_1665_v51).length))];
          _nw282[(new BigNumber(9)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(10)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(11)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(12)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(13)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(14)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(15)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(16)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(17)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(18)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(19)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(20)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(21)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(22)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(23)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(24)).toNumber()] = _1663_v50;
          _nw282[(new BigNumber(25)).toNumber()] = _1663_v50;
          _1667_v53 = _nw282;
          let _index259 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1667_v53).length));
          (_1667_v53)[_index259] = _1663_v50;
          let _index260 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1667_v53).length));
          (_1667_v53)[_index260] = _1663_v50;
          let _1668_v54;
          _1668_v54 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1669_v55;
          let _nw283 = new _module.C4();
          _nw283.__ctor(_1668_v54, _this.f26, (_this).f18, _dafny.Seq.Concat(p2, (_this).f19), (_this).f20);
          _1669_v55 = _nw283;
          (globalState).f3 = _this.f26;
          let _arr0 = (_1667_v53)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_1667_v53).length))];
          let _index261 = _module.__default.safeIndex(new BigNumber(93), new BigNumber(((_1667_v53)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_1667_v53).length))]).length));
          _arr0[_index261] = true;
          let _arr1 = (_1667_v53)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_1667_v53).length))];
          let _index262 = _module.__default.safeIndex(new BigNumber(93), new BigNumber(((_1667_v53)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_1667_v53).length))]).length));
          _arr1[_index262] = (_this).f18;
          let _1670_v56;
          _1670_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1666_v52).length),true);
          _1670_v56 = (_1670_v56).update(p1, true);
        } else {
          (globalState).f3 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(p1, _this.f26), p1);
          let _1671_v57;
          _1671_v57 = _dafny.MultiSet.fromElements(p0);
          _1671_v57 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ukw"), p0), _dafny.Seq.Concat(p2, (_this).f19), (_this).f19, _dafny.Seq.Concat(p2, (_this).f19), p2);
          (globalState).f0 = p1;
          let _1672_v58;
          _1672_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,!(_1639_cf51));
          (globalState).f2 = (p1).isEqualTo(_module.__default.safeDivisionInt(new BigNumber(664), new BigNumber((_1672_v58).length)));
          let _1673_v59;
          _1673_v59 = new _dafny.CodePoint('m'.codePointAt(0));
          let _1674_v60;
          _1674_v60 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1643_v37)[_module.__default.safeIndex(p1, new BigNumber((_1643_v37).length))]),_1673_v59);
          (globalState).f0 = _module.__default.fm1(new BigNumber(((_1674_v60).Merge(_1674_v60)).length), _1639_cf51, true, globalState);
        }
      } else if (_source16.is_DC27) {
        let _1675___mcc_h7 = (_source16).cf50;
        let _1676_cf50 = _1675___mcc_h7;
        (globalState).f0 = p1;
        let _1677_v61;
        let _nw284 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _1677_v61 = _nw284;
        let _index263 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_1677_v61).length));
        (_1677_v61)[_index263] = (_this).f25;
        let _index264 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_1677_v61).length));
        (_1677_v61)[_index264] = _dafny.Seq.Concat((_this).f25, _dafny.Seq.Concat((_this).f25, (_this).f25));
        if ((p1).isLessThanOrEqualTo(new BigNumber((p0).length))) {
          (globalState).f9 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), function (_1678_i7) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("p"));
          (globalState).f2 = (_this).f18;
          (globalState).f2 = false;
          (globalState).f2 = ((_this).f25)[_module.__default.safeIndex(p1, new BigNumber(((_this).f25).length))];
          let _1679_v62;
          let _init37 = ((_1680_p1) => function (_1681_i8) {
            return (_1681_i8).minus(new BigNumber((_dafny.Seq.of((_this).f18, (_module.D0.create_DC1((_this).f18, _1680_p1, (_this).f19, (_this).f18, new BigNumber(-3))).dtor_cf1, (_this).f18)).length));
          })(p1);
          let _nw285 = Array((new BigNumber(12)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw285.length); _i0_37++) {
            _nw285[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1679_v62 = _nw285;
          let _index265 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1679_v62).length));
          (_1679_v62)[_index265] = new BigNumber((((!((_this).f18)) ? (_dafny.Seq.UnicodeFromString("a")) : ((_this).f19))).length);
          let _index266 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1679_v62).length));
          (_1679_v62)[_index266] = new BigNumber(411);
        } else {
          let _1682_v63;
          _1682_v63 = new _dafny.CodePoint('t'.codePointAt(0));
          let _1683_v64;
          _1683_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1682_v63,(_this).f18);
          let _1684_v65;
          _1684_v65 = _dafny.Set.fromElements(_1683_v64, _1683_v64);
          let _1685_v66;
          _1685_v66 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,(_dafny.Set.fromElements(_1683_v64, _1683_v64)).IsDisjointFrom(_1684_v65));
          let _1686_v67;
          _1686_v67 = _module.D8.create_DC20(_1589_v0);
          (globalState).f2 = (((_1685_v66).contains(_module.__default.safeModuloInt(_this.f26, p1))) ? ((_1685_v66).get(_module.__default.safeModuloInt(_this.f26, p1))) : (_module.__default.fm23((_1686_v67).dtor_cf40, _this.f26, globalState)));
          let _index267 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_1677_v61).length));
          (_1677_v61)[_index267] = _dafny.Seq.Concat((_1677_v61)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1677_v61).length))], (_this).f25);
          let _1687_v68;
          let _nw286 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1687_v68 = _nw286;
          let _index268 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1687_v68).length));
          (_1687_v68)[_index268] = _1682_v63;
          let _index269 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1687_v68).length));
          (_1687_v68)[_index269] = _1682_v63;
          let _1688_v69;
          _1688_v69 = _dafny.Map.Empty.slice().updateUnsafe(((_1677_v61)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1677_v61).length))])[_module.__default.safeIndex(_this.f26, new BigNumber(((_1677_v61)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1677_v61).length))]).length))],(_this).f18);
          let _1689_v70;
          _1689_v70 = _dafny.MultiSet.fromElements(_1688_v69, _1688_v69);
          let _1690_v71;
          _1690_v71 = _dafny.MultiSet.fromElements((_this).f18);
          let _1691_v72;
          let _nw287 = Array((new BigNumber(19)).toNumber());
          _nw287[(_dafny.ZERO).toNumber()] = _dafny.Seq.IsProperPrefixOf((_this).f19, p0);
          _nw287[(_dafny.ONE).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(2)).toNumber()] = !(((_1677_v61)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1677_v61).length))])[_module.__default.safeIndex(new BigNumber(403), new BigNumber(((_1677_v61)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1677_v61).length))]).length))]) || ((_this).f18);
          _nw287[(new BigNumber(3)).toNumber()] = ((_this).f25)[_module.__default.safeIndex(p1, new BigNumber(((_this).f25).length))];
          _nw287[(new BigNumber(4)).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(5)).toNumber()] = _module.__default.fm23(_1589_v0, p1, globalState);
          _nw287[(new BigNumber(6)).toNumber()] = !(_this.f26).isEqualTo(_this.f26);
          _nw287[(new BigNumber(7)).toNumber()] = (((_1688_v69).contains((_this).f18)) ? ((_1688_v69).get((_this).f18)) : ((_this).f18));
          _nw287[(new BigNumber(8)).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(9)).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(10)).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(11)).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(12)).toNumber()] = (new BigNumber((_1689_v70).cardinality())).isLessThanOrEqualTo(p1);
          _nw287[(new BigNumber(13)).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(14)).toNumber()] = true;
          _nw287[(new BigNumber(15)).toNumber()] = (_1690_v71).IsProperSubsetOf(_1690_v71);
          _nw287[(new BigNumber(16)).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(17)).toNumber()] = (_this).f18;
          _nw287[(new BigNumber(18)).toNumber()] = (_this).f18;
          _1691_v72 = _nw287;
          let _index270 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_1691_v72).length));
          (_1691_v72)[_index270] = (((_this).f18) ? ((_this).f18) : ((_this).fm2((_this).f18, (_this).f18, !((_this).f18), globalState)));
          let _index271 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_1691_v72).length));
          (_1691_v72)[_index271] = !((_this).f18);
          let _1692_v73;
          _1692_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(890),_1685_v66);
          (globalState).f2 = !((_1685_v66).Merge(_1685_v66)).equals(((((_1692_v73).contains(_this.f26)) ? ((_1692_v73).get(_this.f26)) : (_1685_v66))).update(_this.f26, (_1691_v72)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_1691_v72).length))]));
        }
        (globalState).f0 = _this.f26;
      } else {
        let _1693___mcc_h8 = (_source16).cf52;
        let _1694_cf52 = _1693___mcc_h8;
        (globalState).f2 = (_this).f18;
        (globalState).f2 = (_this).f18;
        if (!((new BigNumber((p2).length)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f26))))) {
          let _1695_v74;
          let _nw288 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _1695_v74 = _nw288;
          let _init38 = function (_1696_i9) {
            return _module.__default.safeDivisionInt(_1696_i9, _this.f26);
          };
          let _nw289 = Array((new BigNumber(17)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw289.length); _i0_38++) {
            _nw289[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1695_v74 = _nw289;
          (globalState).f3 = p1;
          let _1697_v75;
          _1697_v75 = _dafny.MultiSet.fromElements(new BigNumber((_1589_v0).length));
          let _1698_v76;
          _1698_v76 = _dafny.MultiSet.fromElements(new BigNumber(((_module.D16.create_DC38(_1697_v75)).dtor_cf66).cardinality()));
          let _1699_v77;
          _1699_v77 = _dafny.Seq.of(p1, _this.f26);
          let _1700_v78;
          _1700_v78 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1701_v79;
          let _nw290 = new _module.C4();
          _nw290.__ctor(_1700_v78, new BigNumber(563), (_this).f18, _dafny.Seq.UnicodeFromString("stxe"), (_this).f20);
          _1701_v79 = _nw290;
          let _1702_v80;
          _1702_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(436),_1701_v79);
          let _1703_v81;
          _1703_v81 = _dafny.Seq.of((_this).f19);
          let _1704_v82;
          _1704_v82 = _dafny.Map.Empty.slice().updateUnsafe((_1699_v77)[_module.__default.safeIndex(new BigNumber((_1702_v80).length), new BigNumber((_1699_v77).length))],new BigNumber((_1703_v81).length));
          let _1705_v83;
          _1705_v83 = _dafny.Set.fromElements(_1699_v77);
          let _1706_v84;
          _1706_v84 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1705_v83).length),(_1701_v79).f19);
          let _rhs294 = false;
          let _rhs295 = !((_this.f26).isLessThanOrEqualTo(_this.f26));
          let _rhs296 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(700)), ((_1707_p1) => function (_1708_i10) {
            return _1707_p1;
          })(p1)), _module.__default.safeIndex(new BigNumber((_1697_v75).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(700)), ((_1709_p1) => function (_1710_i10) {
            return _1709_p1;
          })(p1))).length)), (_1699_v77)[_module.__default.safeIndex(_this.f26, new BigNumber((_1699_v77).length))]), _module.__default.safeIndex((((_1704_v82).contains(_this.f26)) ? ((_1704_v82).get(_this.f26)) : (p1)), new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(700)), ((_1711_p1) => function (_1712_i10) {
            return _1711_p1;
          })(p1)), _module.__default.safeIndex(new BigNumber((_1697_v75).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(700)), ((_1713_p1) => function (_1714_i10) {
            return _1713_p1;
          })(p1))).length)), (_1699_v77)[_module.__default.safeIndex(_this.f26, new BigNumber((_1699_v77).length))])).length)), new BigNumber((_1706_v84).length));
          let _rhs297 = (_this.f26).minus(new BigNumber(-826));
          let _lhs257 = globalState;
          let _lhs258 = globalState;
          let _lhs259 = globalState;
          let _lhs260 = globalState;
          _lhs257.f2 = _rhs294;
          _lhs258.f2 = _rhs295;
          _lhs259.f10 = _rhs296;
          _lhs260.f0 = _rhs297;
          let _1715_v85;
          let _nw291 = Array((new BigNumber(9)).toNumber());
          _nw291[(_dafny.ZERO).toNumber()] = _1634_v30;
          _nw291[(_dafny.ONE).toNumber()] = _1634_v30;
          _nw291[(new BigNumber(2)).toNumber()] = _1634_v30;
          _nw291[(new BigNumber(3)).toNumber()] = _1634_v30;
          _nw291[(new BigNumber(4)).toNumber()] = _1634_v30;
          _nw291[(new BigNumber(5)).toNumber()] = _1634_v30;
          _nw291[(new BigNumber(6)).toNumber()] = _1634_v30;
          _nw291[(new BigNumber(7)).toNumber()] = _1634_v30;
          _nw291[(new BigNumber(8)).toNumber()] = _1634_v30;
          _1715_v85 = _nw291;
          let _index272 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1715_v85).length));
          (_1715_v85)[_index272] = _1634_v30;
          let _index273 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1715_v85).length));
          (_1715_v85)[_index273] = _1634_v30;
          let _1716_v86;
          let _init39 = ((_1717_v79) => function (_1718_i11) {
            return !((_1717_v79).f18) || ((_this).f18);
          })(_1701_v79);
          let _nw292 = Array((new BigNumber(28)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw292.length); _i0_39++) {
            _nw292[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1716_v86 = _nw292;
          let _index274 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1716_v86).length));
          (_1716_v86)[_index274] = (_1701_v79).f18;
          let _index275 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1716_v86).length));
          (_1716_v86)[_index275] = !_dafny.Seq.contains(_dafny.Seq.of(!((_this).f18)), (_this).f18);
        } else {
          let _1719_v87;
          let _1720_v88;
          let _out50;
          let _out51;
          let _outcollector13 = (_this).m6(globalState);
          _out50 = _outcollector13[0];
          _out51 = _outcollector13[1];
          _1719_v87 = _out50;
          _1720_v88 = _out51;
          let _1721_v89;
          _1721_v89 = new _dafny.CodePoint('x'.codePointAt(0));
          _1721_v89 = _1721_v89;
          let _1722_v90;
          let _nw293 = Array((new BigNumber(4)).toNumber());
          _nw293[(_dafny.ZERO).toNumber()] = _1721_v89;
          _nw293[(_dafny.ONE).toNumber()] = _1721_v89;
          _nw293[(new BigNumber(2)).toNumber()] = _1721_v89;
          _nw293[(new BigNumber(3)).toNumber()] = _module.__default.fm46(globalState);
          _1722_v90 = _nw293;
          _1722_v90 = _1722_v90;
          let _1723_v91;
          _1723_v91 = _module.D2.create_DC8(_1719_v87, false, (_this).f18);
          let _1724_v92;
          let _nw294 = Array((new BigNumber(15)).toNumber());
          _nw294[(_dafny.ZERO).toNumber()] = _1719_v87;
          _nw294[(_dafny.ONE).toNumber()] = false;
          _nw294[(new BigNumber(2)).toNumber()] = _1719_v87;
          _nw294[(new BigNumber(3)).toNumber()] = false;
          _nw294[(new BigNumber(4)).toNumber()] = (_this).f18;
          _nw294[(new BigNumber(5)).toNumber()] = false;
          _nw294[(new BigNumber(6)).toNumber()] = (_this).f18;
          _nw294[(new BigNumber(7)).toNumber()] = (_this).f18;
          _nw294[(new BigNumber(8)).toNumber()] = (_this).f18;
          _nw294[(new BigNumber(9)).toNumber()] = _1719_v87;
          _nw294[(new BigNumber(10)).toNumber()] = (_this).f18;
          _nw294[(new BigNumber(11)).toNumber()] = (_this).f18;
          _nw294[(new BigNumber(12)).toNumber()] = _1719_v87;
          _nw294[(new BigNumber(13)).toNumber()] = _1719_v87;
          _nw294[(new BigNumber(14)).toNumber()] = _1719_v87;
          _1724_v92 = _nw294;
          let _1725_v93;
          _1725_v93 = _module.D13.create_DC34(_1724_v92, _dafny.Seq.UnicodeFromString("s"), !((_this).f18), false);
          let _1726_v94;
          _1726_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1723_v91,_module.D13.create_DC35(_1725_v93));
          _1726_v94 = (_1726_v94).update(function (_pat_let47_0) {
            return function (_1727_dt__update__tmp_h0) {
              return function (_pat_let48_0) {
                return function (_1728_dt__update_hcf19_h0) {
                  return _module.D2.create_DC8((_1727_dt__update__tmp_h0).dtor_cf18, _1728_dt__update_hcf19_h0, (_1727_dt__update__tmp_h0).dtor_cf20);
                }(_pat_let48_0);
              }(!(true));
            }(_pat_let47_0);
          }(_module.D2.create_DC8((_this).f18, _1719_v87, _1719_v87)), _module.D13.create_DC35(_1725_v93));
          (globalState).f15 = _1720_v88;
        }
        (globalState).f0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
      }
      let _1729_v95;
      _1729_v95 = new _dafny.CodePoint('c'.codePointAt(0));
      _1729_v95 = _1729_v95;
      let _hi9 = new BigNumber((_dafny.Seq.UnicodeFromString("tsjsavsim")).length);
      for (let _1730_i12 = new BigNumber(2); _1730_i12.isLessThan(_hi9); _1730_i12 = _1730_i12.plus(_dafny.ONE)) {
        (globalState).f15 = _1730_i12;
        (globalState).f12 = p0;
        let _1731_v96;
        _1731_v96 = _module.D0.create_DC1((_this).f18, _this.f26, (_this).f19, !((_this).f18), p1);
        let _1732_v97;
        _1732_v97 = _dafny.Seq.of(_this.f26, p1);
        let _1733_v98;
        _1733_v98 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p2);
        let _1734_v99;
        let _nw295 = Array((new BigNumber(21)).toNumber());
        _nw295[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(((_this).f25).length), _this.f26));
        _nw295[(_dafny.ONE).toNumber()] = _this.f26;
        _nw295[(new BigNumber(2)).toNumber()] = (_1730_i12).minus(p1);
        _nw295[(new BigNumber(3)).toNumber()] = _1730_i12;
        _nw295[(new BigNumber(4)).toNumber()] = _module.__default.fm1(_this.f26, (_this).f18, (_this).f18, globalState);
        _nw295[(new BigNumber(5)).toNumber()] = _1730_i12;
        _nw295[(new BigNumber(6)).toNumber()] = p1;
        _nw295[(new BigNumber(7)).toNumber()] = _1730_i12;
        _nw295[(new BigNumber(8)).toNumber()] = _this.f26;
        _nw295[(new BigNumber(9)).toNumber()] = p1;
        _nw295[(new BigNumber(10)).toNumber()] = _this.f26;
        _nw295[(new BigNumber(11)).toNumber()] = _this.f26;
        _nw295[(new BigNumber(12)).toNumber()] = _this.f26;
        _nw295[(new BigNumber(13)).toNumber()] = p1;
        _nw295[(new BigNumber(14)).toNumber()] = new BigNumber(487);
        _nw295[(new BigNumber(15)).toNumber()] = _this.f26;
        _nw295[(new BigNumber(16)).toNumber()] = (_1634_v30).fm22((_this).f19, _1731_v96, globalState);
        _nw295[(new BigNumber(17)).toNumber()] = _1730_i12;
        _nw295[(new BigNumber(18)).toNumber()] = _1730_i12;
        _nw295[(new BigNumber(19)).toNumber()] = (new BigNumber(((_module.D0.create_DC1((_this).f18, new BigNumber((_1732_v97).length), _module.__default.fm48(globalState), (_this).f18, _1730_i12)).dtor_cf3).length)).multipliedBy(new BigNumber(((((_1733_v98).contains((_this).f18)) ? ((_1733_v98).get((_this).f18)) : (_dafny.Seq.update(p0, _module.__default.safeIndex(p1, new BigNumber((p0).length)), new _dafny.CodePoint('a'.codePointAt(0)))))).length));
        _nw295[(new BigNumber(20)).toNumber()] = p1;
        _1734_v99 = _nw295;
        _1734_v99 = _1734_v99;
        let _1735_v100;
        _1735_v100 = _module.D0.create_DC0((_this).f18);
        let _1736_v101;
        _1736_v101 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,_1735_v100);
        let _1737_v102;
        _1737_v102 = _dafny.Map.Empty.slice().updateUnsafe(_1736_v101,(_this).f18);
        let _1738_v103;
        _1738_v103 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(((_1737_v102).contains(_1736_v101)) ? ((_1737_v102).get(_1736_v101)) : ((_this).f18)));
        (globalState).f0 = new BigNumber((_1738_v103).length);
      }
      let _1739_v104;
      _1739_v104 = _dafny.Set.fromElements(_1729_v95, _1729_v95);
      let _1740_v105;
      _1740_v105 = _dafny.Map.Empty.slice().updateUnsafe(true,_1739_v104);
      let _1741_v106;
      let _nw296 = Array((new BigNumber(26)).toNumber());
      _nw296[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(p2, p0);
      _nw296[(_dafny.ONE).toNumber()] = p2;
      _nw296[(new BigNumber(2)).toNumber()] = (((_this).f18) ? (_dafny.Seq.UnicodeFromString("mamkhfj")) : (p2));
      _nw296[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), ((_1742_v95) => function (_1743_i13) {
        return _1742_v95;
      })(_1729_v95)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), ((_1744_v95) => function (_1745_i13) {
        return _1744_v95;
      })(_1729_v95))).length)), _1729_v95);
      _nw296[(new BigNumber(4)).toNumber()] = p0;
      _nw296[(new BigNumber(5)).toNumber()] = (_this).f19;
      _nw296[(new BigNumber(6)).toNumber()] = p2;
      _nw296[(new BigNumber(7)).toNumber()] = p2;
      _nw296[(new BigNumber(8)).toNumber()] = (_1634_v30).fm21(new BigNumber((p0).length), (_this).f19, _1740_v105, (_this).f25, globalState);
      _nw296[(new BigNumber(9)).toNumber()] = (((_this).f18) ? (_dafny.Seq.update((_this).f19, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), function (_1746_i14) {
        return (_dafny.ZERO).minus(_this.f26);
      })).length), new BigNumber(((_this).f19).length)), _1729_v95)) : (_dafny.Seq.UnicodeFromString("b")));
      _nw296[(new BigNumber(10)).toNumber()] = (_this).f19;
      _nw296[(new BigNumber(11)).toNumber()] = (_this).f19;
      _nw296[(new BigNumber(12)).toNumber()] = (((_this).f18) ? (p0) : (_module.__default.fm48(globalState)));
      _nw296[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("yntu");
      _nw296[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-992)), ((_1747_v95) => function (_1748_i15) {
        return _1747_v95;
      })(_1729_v95)), p2);
      _nw296[(new BigNumber(15)).toNumber()] = (_this).f19;
      _nw296[(new BigNumber(16)).toNumber()] = (_this).f19;
      _nw296[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(new BigNumber(179))).length), new BigNumber((p0).length)), _module.__default.fm46(globalState)), _dafny.Seq.UnicodeFromString("ohly"));
      _nw296[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("mssdb");
      _nw296[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(p0, p2);
      _nw296[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("leqkgaus");
      _nw296[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-677)), ((_1749_v95) => function (_1750_i16) {
        return _1749_v95;
      })(_1729_v95)), (_this).f19);
      _nw296[(new BigNumber(22)).toNumber()] = _dafny.Seq.UnicodeFromString("nkvl");
      _nw296[(new BigNumber(23)).toNumber()] = p2;
      _nw296[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(p2, p2);
      _nw296[(new BigNumber(25)).toNumber()] = _dafny.Seq.UnicodeFromString("ocvvqga");
      _1741_v106 = _nw296;
      let _index276 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1741_v106).length));
      (_1741_v106)[_index276] = p0;
      let _index277 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1741_v106).length));
      (_1741_v106)[_index277] = ((!((_this).f18) || ((_this).f18)) ? ((_this).f19) : (p2));
      return;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if (false) {
        if ((_this).f18) {
          let _1751_v0;
          _1751_v0 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1752_v1;
          _1752_v1 = _module.D3.create_DC13(new BigNumber(((_this).f19).length), (_this).f19, new BigNumber(972), _this.f26, _1751_v0);
          let _1753_v2;
          _1753_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(238)), function (_1754_i0) {
            return _module.D3.create_DC13((_dafny.ZERO).minus(_this.f26), _dafny.Seq.UnicodeFromString("hpkoolb"), _this.f26, new BigNumber(809), new _dafny.CodePoint('f'.codePointAt(0)));
          }), _dafny.Seq.of(_1752_v1, _1752_v1, _1752_v1, _1752_v1)));
          let _1755_v3;
          _1755_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,(_this).f18);
          let _1756_v4;
          _1756_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1755_v3,new BigNumber(((_this).f19).length));
          let _1757_v5;
          _1757_v5 = _dafny.Seq.of(_module.D3.create_DC13(_this.f26, _dafny.Seq.Create(_module.__default.abs(new BigNumber(984)), ((_1758_v0) => function (_1759_i1) {
  return _1758_v0;
})(_1751_v0)), _this.f26, (((_1756_v4).contains(_1755_v3)) ? ((_1756_v4).get(_1755_v3)) : (new BigNumber(505))), _1751_v0), _1752_v1);
          _1753_v2 = (_1753_v2).update(true, _dafny.Seq.Concat(_1757_v5, _1757_v5));
          let _1760_v6;
          _1760_v6 = _dafny.MultiSet.fromElements(_this.f26, _this.f26);
          let _1761_v7;
          _1761_v7 = _dafny.Set.fromElements((_this).f18, p1);
          let _1762_v8;
          _1762_v8 = _dafny.Seq.of(new BigNumber(-53), new BigNumber((_1761_v7).length), _module.__default.fm1(_this.f26, p2, p2, globalState), _this.f26);
          let _1763_v9;
          _1763_v9 = _dafny.Set.fromElements(_this.f26);
          let _rhs298 = _this.f26;
          let _rhs299 = p1;
          let _rhs300 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update(_1762_v8, _module.__default.safeIndex(_this.f26, new BigNumber((_1762_v8).length)), new BigNumber((_1763_v9).length)), _1762_v8))).IsSubsetOf((_1760_v6).Intersect(_1760_v6));
          let _lhs261 = globalState;
          let _lhs262 = globalState;
          let _lhs263 = globalState;
          _lhs261.f0 = _rhs298;
          _lhs262.f2 = _rhs299;
          _lhs263.f2 = _rhs300;
          let _1764_v10;
          _1764_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this.f26).plus((_dafny.ZERO).minus(_this.f26)),_module.__default.fm52(_this.f26, (_1762_v8)[_module.__default.safeIndex(_this.f26, new BigNumber((_1762_v8).length))], globalState));
          let _1765_v11;
          _1765_v11 = _dafny.MultiSet.fromElements((_this).f19, (_this).f19, (_this).f19, (_this).f19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), function (_1766_i2) {
            return ((_this).f19)[_module.__default.safeIndex(_this.f26, new BigNumber(((_this).f19).length))];
          }));
          _1764_v10 = (_1764_v10).update(_this.f26, (_1765_v11).Union(_1765_v11));
          let _1767_v12;
          let _1768_v13;
          let _1769_v14;
          let _1770_v15;
          let _out52;
          let _out53;
          let _out54;
          let _out55;
          let _outcollector14 = _module.__default.m0(globalState);
          _out52 = _outcollector14[0];
          _out53 = _outcollector14[1];
          _out54 = _outcollector14[2];
          _out55 = _outcollector14[3];
          _1767_v12 = _out52;
          _1768_v13 = _out53;
          _1769_v14 = _out54;
          _1770_v15 = _out55;
          let _index278 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1769_v14).length));
          (_1769_v14)[_index278] = new BigNumber(((_this).f19).length);
          let _1771_v16;
          let _nw297 = Array((new BigNumber(25)).toNumber());
          _nw297[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
          _nw297[(_dafny.ONE).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(2)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(3)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(4)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(5)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(6)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(7)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(8)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(9)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(10)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(11)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(12)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(13)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(14)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(15)).toNumber()] = _module.__default.fm46(globalState);
          _nw297[(new BigNumber(16)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(17)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(18)).toNumber()] = _module.__default.fm46(globalState);
          _nw297[(new BigNumber(19)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(20)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(21)).toNumber()] = _1751_v0;
          _nw297[(new BigNumber(22)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw297[(new BigNumber(23)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw297[(new BigNumber(24)).toNumber()] = _1751_v0;
          _1771_v16 = _nw297;
          let _1772_v17;
          _1772_v17 = _dafny.Set.fromElements(_1771_v16, _1771_v16, (((_this).fm2(false, p1, p2, globalState)) ? (_1771_v16) : (_1771_v16)), _1771_v16, _1771_v16);
          let _index279 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1769_v14).length));
          let _rhs301 = _this.f26;
          let _rhs302 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_1768_v13, _1768_v13));
          let _rhs303 = _dafny.Set.fromElements(_1771_v16);
          let _lhs264 = _1769_v14;
          let _lhs265 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1769_v14).length));
          let _lhs266 = globalState;
          _lhs264[_lhs265] = _rhs301;
          _lhs266.f0 = _rhs302;
          _1772_v17 = _rhs303;
        } else {
          let _1773_v18;
          let _nw298 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1773_v18 = _nw298;
          let _1774_v19;
          _1774_v19 = _module.D13.create_DC34(_1773_v18, _module.__default.fm48(globalState), p1, p1);
          let _1775_v20;
          _1775_v20 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1773_v18);
          let _1776_v21;
          _1776_v21 = _dafny.Set.fromElements(new BigNumber((_1775_v20).length), _this.f26);
          let _1777_v22;
          _1777_v22 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (_1774_v19) : (_1774_v19)),!(_1776_v21).contains(_this.f26));
          _1777_v22 = (_1777_v22).update(_1774_v19, ((_this).f18) || (!((_this).f18)));
          let _1778_v23;
          let _nw299 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1778_v23 = _nw299;
          let _index280 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_1778_v23).length));
          (_1778_v23)[_index280] = new BigNumber(813);
          let _1779_v24;
          _1779_v24 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,true);
          let _index281 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_1778_v23).length));
          let _rhs304 = (_this).f18;
          let _rhs305 = (_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1779_v24).length)), _this.f26)).minus(_this.f26);
          let _rhs306 = _dafny.Seq.UnicodeFromString("fphpatnfw");
          let _rhs307 = (_this.f26).isLessThan(new BigNumber((_module.__default.fm36(!((_this).f18), _this.f26, p1, globalState)).length));
          let _rhs308 = _this.f26;
          let _lhs267 = globalState;
          let _lhs268 = _1778_v23;
          let _lhs269 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_1778_v23).length));
          let _lhs270 = globalState;
          let _lhs271 = globalState;
          let _lhs272 = globalState;
          _lhs267.f2 = _rhs304;
          _lhs268[_lhs269] = _rhs305;
          _lhs270.f12 = _rhs306;
          _lhs271.f2 = _rhs307;
          _lhs272.f3 = _rhs308;
          let _1780_v25;
          _1780_v25 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1781_v26;
          let _nw300 = new _module.C4();
          _nw300.__ctor(_1780_v25, (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f26)), p1, (_this).f19, (_this).f20);
          _1781_v26 = _nw300;
          let _1782_v27;
          _1782_v27 = _dafny.Seq.of(_1781_v26);
          let _1783_v28;
          let _nw301 = new _module.C3();
          _nw301.__ctor(_dafny.Seq.of(_this.f26, new BigNumber((_1782_v27).length)), (_1781_v26).f19, (_this).f20, p1);
          _1783_v28 = _nw301;
          let _1784_v29;
          _1784_v29 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm1(_this.f26, p1, p1, globalState)).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(841)), function (_1785_i3) {
            return _this.f26;
          })).length)),_1781_v26);
          let _1786_v30;
          _1786_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1783_v28).f18,(_1783_v28).f18);
          _1784_v29 = (_1784_v29).update((_1781_v26).fm2((_module.D0.create_DC0((_1781_v26).f18)).dtor_cf0, (_1783_v28).f18, ((_this).f25)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1786_v30).length),_1780_v25)).length), new BigNumber(((_this).f25).length))], globalState), _1781_v26);
          let _1787_v31;
          let _nw302 = Array((new BigNumber(2)).toNumber());
          _nw302[(_dafny.ZERO).toNumber()] = _1778_v23;
          _nw302[(_dafny.ONE).toNumber()] = (((_1781_v26).f18) ? (_1778_v23) : (_1778_v23));
          _1787_v31 = _nw302;
          let _1788_v32;
          _1788_v32 = _dafny.MultiSet.fromElements(_this.f26, _module.__default.fm1(_this.f26, p2, p2, globalState));
          let _1789_v33;
          _1789_v33 = _module.D16.create_DC39(!((_1788_v32).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f26, (_dafny.ZERO).minus((_1778_v23)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_1778_v23).length))]))))), new BigNumber(951));
          let _1790_v34;
          _1790_v34 = _module.D2.create_DC8(p2, (_this).f18, true);
          let _1791_v35;
          _1791_v35 = _dafny.MultiSet.fromElements(p1, (_1790_v34).dtor_cf19, p1, (_1781_v26).f18);
          let _1792_v36;
          _1792_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1781_v26,_1780_v25);
          let _pat_let_tv25 = _1778_v23;
          let _pat_let_tv26 = _1778_v23;
          let _rhs309 = (_1778_v23)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_1778_v23).length))];
          let _rhs310 = _1787_v31;
          let _rhs311 = function (_pat_let49_0) {
            return function (_1793_dt__update__tmp_h0) {
              return function (_pat_let50_0) {
                return function (_1794_dt__update_hcf67_h0) {
                  return _module.D16.create_DC39(_1794_dt__update_hcf67_h0, (_1793_dt__update__tmp_h0).dtor_cf68);
                }(_pat_let50_0);
              }((_this.f26).isLessThan((_pat_let_tv26)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_pat_let_tv25).length))]));
            }(_pat_let49_0);
          }(_1789_v33);
          let _rhs312 = (_module.__default.fm32(globalState)).Difference((_dafny.MultiSet.fromElements((_1781_v26).fm2(false, (_1781_v26).f18, (_1783_v28).f18, globalState))).Union((_1791_v35).update(p2, _module.__default.abs(new BigNumber((_1792_v36).length)))));
          let _rhs313 = (_module.__default.fm55(_this.f26, p2, new BigNumber((_1786_v30).length), (_1783_v28).f18, globalState)).dtor_cf46;
          let _lhs273 = globalState;
          let _lhs274 = globalState;
          _lhs273.f15 = _rhs309;
          _1787_v31 = _rhs310;
          _1789_v33 = _rhs311;
          _1791_v35 = _rhs312;
          _lhs274.f15 = _rhs313;
          let _rhs314 = (_1783_v28).f19;
          let _rhs315 = ((_dafny.ZERO).minus((_1778_v23)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_1778_v23).length))])).multipliedBy(new BigNumber(-479));
          let _lhs275 = globalState;
          let _lhs276 = globalState;
          _lhs275.f9 = _rhs314;
          _lhs276.f3 = _rhs315;
        }
        let _1795_v37;
        _1795_v37 = _dafny.Set.fromElements(_this.f26);
        let _1796_v38;
        _1796_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p1);
        let _1797_v39;
        _1797_v39 = _dafny.Seq.of(new BigNumber((_1795_v37).length), _module.__default.safeModuloInt(_this.f26, _this.f26), _this.f26, _module.__default.fm1(new BigNumber(((_this).f25).length), (((_1796_v38).contains((_this).f25)) ? ((_1796_v38).get((_this).f25)) : (p1)), (_this).f18, globalState));
        (globalState).f0 = (_1797_v39)[_module.__default.safeIndex((_this.f26).multipliedBy(_this.f26), new BigNumber((_1797_v39).length))];
        let _1798_v40;
        let _nw303 = new _module.C1();
        _nw303.__ctor(p1);
        _1798_v40 = _nw303;
        (globalState).f9 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(838)), function (_1799_i4) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }), (_this).f19);
        let _1800_v41;
        let _nw304 = Array((new BigNumber(3)).toNumber());
        _nw304[(_dafny.ZERO).toNumber()] = (_this).f18;
        _nw304[(_dafny.ONE).toNumber()] = p1;
        _nw304[(new BigNumber(2)).toNumber()] = p1;
        _1800_v41 = _nw304;
        let _1801_v42;
        _1801_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1800_v41,true);
        (globalState).f0 = new BigNumber((_1801_v42).length);
      } else {
        let _1802_v43;
        _1802_v43 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1803_v44;
        let _nw305 = new _module.C4();
        _nw305.__ctor(_1802_v43, _this.f26, p1, _dafny.Seq.Concat((_this).f19, (_this).f19), (_this).f20);
        _1803_v44 = _nw305;
        let _1804_v45;
        _1804_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
        (globalState).f2 = (_1804_v45).equals((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).Merge((_1804_v45).update(p2, true)));
        if (p1) {
          (_this).f26 = (_1803_v44).f30;
          (globalState).f15 = _module.__default.safeDivisionInt(_this.f26, new BigNumber(941));
          let _1805_v46;
          let _nw306 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1805_v46 = _nw306;
          let _index282 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1805_v46).length));
          (_1805_v46)[_index282] = (new BigNumber(((_this).f19).length)).minus(_this.f26);
          let _1806_v47;
          _1806_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_1803_v44).f30);
          let _index283 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1805_v46).length));
          (_1805_v46)[_index283] = (new BigNumber(183)).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(((_this).f25).length))).Merge(_1806_v47)).length));
          (globalState).f2 = p1;
          (globalState).f12 = (_this).f19;
        } else {
          let _1807_v48;
          let _nw307 = new _module.C1();
          _nw307.__ctor(false);
          _1807_v48 = _nw307;
          let _1808_v49;
          _1808_v49 = _dafny.MultiSet.fromElements(_1807_v48, _1807_v48, _1807_v48);
          (globalState).f2 = ((_dafny.MultiSet.fromElements(_1807_v48, _1807_v48)).Intersect(_1808_v49)).IsProperSubsetOf(_1808_v49);
          let _1809_v51;
          let _init40 = function (_1810_i5) {
            return function () {
              let _coll48 = new _dafny.Set();
              for (const _compr_48 of _dafny.IntegerRange(new BigNumber(186), new BigNumber(24))) {
                let _1811_v50 = _compr_48;
                if (((new BigNumber(186)).isLessThanOrEqualTo(_1811_v50)) && ((_1811_v50).isLessThan(new BigNumber(24)))) {
                  _coll48.add((_1811_v50).minus((_dafny.ZERO).minus(_this.f26)));
                }
              }
              return _coll48;
            }();
          };
          let _nw308 = Array((new BigNumber(22)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw308.length); _i0_40++) {
            _nw308[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1809_v51 = _nw308;
          let _1812_v52;
          _1812_v52 = _dafny.Map.Empty.slice().updateUnsafe((_1803_v44).f30,p2);
          let _1813_v53;
          _1813_v53 = _dafny.Set.fromElements(new BigNumber((_1812_v52).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-716)), ((_1814_v44) => function (_1815_i6) {
            return (_1814_v44).f29;
          })(_1803_v44))).length), (_1803_v44).f30);
          let _index284 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1809_v51).length));
          (_1809_v51)[_index284] = _1813_v53;
          let _index285 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1809_v51).length));
          (_1809_v51)[_index285] = _1813_v53;
          let _1816_v54;
          _1816_v54 = _dafny.Set.fromElements(p2, p2);
          let _1817_v55;
          _1817_v55 = _dafny.MultiSet.fromElements(_1816_v54, _1816_v54);
          let _1818_v56;
          _1818_v56 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),new BigNumber(((_1817_v55).Union(_1817_v55)).cardinality()));
          _1818_v56 = (_1818_v56).update((_this).f18, (_1803_v44).f30);
          let _1819_v57;
          let _nw309 = new _module.C1();
          _nw309.__ctor((_this).f18);
          _1819_v57 = _nw309;
          let _1820_v58;
          _1820_v58 = _module.D17.create_DC41(_1812_v52);
          let _1821_v59;
          _1821_v59 = _dafny.Seq.of((_1812_v52).Merge((_1820_v58).dtor_cf73));
          _1821_v59 = _dafny.Seq.Concat(_1821_v59, _1821_v59);
        }
        let _1822_v60;
        let _nw310 = Array((new BigNumber(14)).toNumber()).fill(false);
        _1822_v60 = _nw310;
        let _1823_v61;
        _1823_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1822_v60,_1822_v60);
        let _1824_v62;
        let _init41 = function (_1825_i7) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        };
        let _nw311 = Array((new BigNumber(17)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw311.length); _i0_41++) {
          _nw311[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1824_v62 = _nw311;
        let _index286 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_1824_v62).length));
        (_1824_v62)[_index286] = new _dafny.CodePoint('h'.codePointAt(0));
        let _1826_v63;
        _1826_v63 = _module.D13.create_DC33((_1803_v44).fm3(_this.f26, globalState), _module.__default.safeModuloInt((_1803_v44).f30, new BigNumber(788)), p1, (_this.f26).plus((_1803_v44).f30));
        let _index287 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_1824_v62).length));
        let _rhs316 = _1823_v61;
        let _rhs317 = (_1803_v44).f29;
        let _rhs318 = (((_this).f18) ? (_1802_v43) : ((_1803_v44).f29));
        let _rhs319 = _module.__default.fm56(p2, (_this).f18, _this.f26, _dafny.Seq.Create(_module.__default.abs(new BigNumber(798)), ((_1827_v44) => function (_1828_i8) {
          return (_1827_v44).f29;
        })(_1803_v44)), globalState);
        let _lhs277 = _1824_v62;
        let _lhs278 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_1824_v62).length));
        _1823_v61 = _rhs316;
        _1802_v43 = _rhs317;
        _lhs277[_lhs278] = _rhs318;
        _1826_v63 = _rhs319;
        let _1829_v64;
        let _nw312 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1829_v64 = _nw312;
        let _1830_v65;
        _1830_v65 = _dafny.MultiSet.fromElements(_1829_v64);
        let _1831_v66;
        let _nw313 = new _module.C4();
        _nw313.__ctor(new _dafny.CodePoint('m'.codePointAt(0)), (_1803_v44).f30, p2, _dafny.Seq.UnicodeFromString("swvwp"), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1830_v65).cardinality()),_this.f26)));
        _1831_v66 = _nw313;
      }
      (_this).f26 = _this.f26;
      if ((p2) || ((!((_this).f18)) || (!((_this).f18)))) {
        let _1832_v67;
        let _init42 = function (_1833_i9) {
          return (_this).f18;
        };
        let _nw314 = Array((new BigNumber(14)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw314.length); _i0_42++) {
          _nw314[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1832_v67 = _nw314;
        let _index288 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length));
        (_1832_v67)[_index288] = (_this).f18;
        let _index289 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length));
        (_1832_v67)[_index289] = (_this).f18;
        let _1834_v68;
        _1834_v68 = _dafny.Seq.of(new BigNumber(471), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))],false)).length));
        let _1835_v69;
        let _nw315 = new _module.C2();
        _nw315.__ctor(_this.f26, ((p2) ? (_1834_v68) : (_1834_v68)), p1);
        _1835_v69 = _nw315;
        if (((_dafny.ZERO).minus((_dafny.ZERO).minus(_1835_v69.f33))).isEqualTo(new BigNumber(((_this).f19).length))) {
          let _1836_v70;
          _1836_v70 = _dafny.MultiSet.fromElements(!((_this).f18), p2, true);
          let _1837_v71;
          _1837_v71 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1838_v72;
          _1838_v72 = _module.D9.create_DC23(p2, p2, _1835_v69.f33);
          let _1839_v73;
          _1839_v73 = _dafny.Set.fromElements(((_this).f25)[_module.__default.safeIndex(_1835_v69.f33, new BigNumber(((_this).f25).length))]);
          let _1840_v74;
          _1840_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1835_v69.f33,_1835_v69.f33);
          let _1841_v75;
          _1841_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1840_v74).length),_this.f26);
          let _1842_v76;
          _1842_v76 = _dafny.MultiSet.fromElements(_1840_v74);
          let _1843_v77;
          _1843_v77 = _module.D17.create_DC43(_1842_v76, (_this).f18);
          let _1844_v78;
          _1844_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1835_v69.f33);
          let _1845_v79;
          _1845_v79 = _module.D2.create_DC7(_1835_v69.f33, (_this).f18, _this.f26);
          let _1846_v80;
          let _nw316 = Array((new BigNumber(29)).toNumber());
          _nw316[(_dafny.ZERO).toNumber()] = _this.f26;
          _nw316[(_dafny.ONE).toNumber()] = new BigNumber(432);
          _nw316[(new BigNumber(2)).toNumber()] = _1835_v69.f33;
          _nw316[(new BigNumber(3)).toNumber()] = (((_1836_v70).contains((_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))])) ? ((_1836_v70).get((_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))])) : (_this.f26));
          _nw316[(new BigNumber(4)).toNumber()] = _1835_v69.f33;
          _nw316[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.of(_1837_v71, _1837_v71)).length);
          _nw316[(new BigNumber(6)).toNumber()] = _1835_v69.f33;
          _nw316[(new BigNumber(7)).toNumber()] = (_module.__default.fm1(((_1835_v69).f34)[_module.__default.safeIndex(_this.f26, new BigNumber(((_1835_v69).f34).length))], p2, (_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))], globalState)).plus(new BigNumber((_dafny.MultiSet.fromElements(_1834_v68, _module.__default.fm33(true, globalState))).cardinality()));
          _nw316[(new BigNumber(8)).toNumber()] = new BigNumber(884);
          _nw316[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(_this.f26, _1835_v69.f33);
          _nw316[(new BigNumber(10)).toNumber()] = _1835_v69.f33;
          _nw316[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt((_1838_v72).dtor_cf46, _1835_v69.f33);
          _nw316[(new BigNumber(12)).toNumber()] = _1835_v69.f33;
          _nw316[(new BigNumber(13)).toNumber()] = (_module.D17.create_DC44(_1832_v67, _1835_v69.f33)).dtor_cf78;
          _nw316[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1839_v73).length), _1835_v69.f33);
          _nw316[(new BigNumber(15)).toNumber()] = _this.f26;
          _nw316[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_1843_v77).dtor_cf76)).length);
          _nw316[(new BigNumber(17)).toNumber()] = _this.f26;
          _nw316[(new BigNumber(18)).toNumber()] = _this.f26;
          _nw316[(new BigNumber(19)).toNumber()] = (_1835_v69.f33).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(323)), ((_1847_v71) => function (_1848_i10) {
            return _1847_v71;
          })(_1837_v71))).length));
          _nw316[(new BigNumber(20)).toNumber()] = _this.f26;
          _nw316[(new BigNumber(21)).toNumber()] = (((_1836_v70).contains((_this).f18)) ? ((_1836_v70).get((_this).f18)) : (_module.__default.fm1((_dafny.ZERO).minus(_this.f26), (_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))], p2, globalState)));
          _nw316[(new BigNumber(22)).toNumber()] = _1835_v69.f33;
          _nw316[(new BigNumber(23)).toNumber()] = ((_1835_v69).f34)[_module.__default.safeIndex(_1835_v69.f33, new BigNumber(((_1835_v69).f34).length))];
          _nw316[(new BigNumber(24)).toNumber()] = (((_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))]) ? ((((_1844_v78).contains((_this).f18)) ? ((_1844_v78).get((_this).f18)) : (new BigNumber(56)))) : (_module.__default.fm1(new BigNumber(903), (_this).f18, (_1835_v69).fm19(_1845_v79, globalState), globalState)));
          _nw316[(new BigNumber(25)).toNumber()] = _1835_v69.f33;
          _nw316[(new BigNumber(26)).toNumber()] = _1835_v69.f33;
          _nw316[(new BigNumber(27)).toNumber()] = (new BigNumber(845)).minus(_1835_v69.f33);
          _nw316[(new BigNumber(28)).toNumber()] = _this.f26;
          _1846_v80 = _nw316;
          let _index290 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_1846_v80).length));
          (_1846_v80)[_index290] = _1835_v69.f33;
          let _index291 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_1846_v80).length));
          (_1846_v80)[_index291] = _1835_v69.f33;
          let _1849_v81;
          _1849_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(236),_1846_v80);
          _1849_v81 = ((_1849_v81).Merge(_1849_v81)).Merge(_1849_v81);
          let _1850_v82;
          _1850_v82 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.fm57(_1835_v69.f33, new BigNumber(((_this).f19).length), true, globalState));
          _1850_v82 = _1850_v82;
          let _1851_v83;
          _1851_v83 = _module.D8.create_DC21(_1837_v71, new BigNumber((_dafny.Seq.of(p2)).length));
          _1837_v71 = (_1851_v83).dtor_cf41;
          _1841_v75 = (_1841_v75).update((new BigNumber(((_this).f25).length)).minus((_1846_v80)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_1846_v80).length))]), _1835_v69.f33);
        } else {
          (globalState).f0 = _this.f26;
          (globalState).f2 = (_this).f18;
          let _index292 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length));
          (_1832_v67)[_index292] = (_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))];
          let _index293 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length));
          (_1832_v67)[_index293] = ((_this).f25)[_module.__default.safeIndex(new BigNumber((((p2) ? (_dafny.Seq.update(_1834_v68, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-208)), new BigNumber((_1834_v68).length)), new BigNumber((_1834_v68).length))) : (_1834_v68))).length), new BigNumber(((_this).f25).length))];
          let _1852_v84;
          _1852_v84 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1853_v85;
          _1853_v85 = _dafny.Set.fromElements(new BigNumber((_module.__default.fm43(p1, p2, globalState)).length));
          let _rhs320 = _module.__default.fm46(globalState);
          let _rhs321 = ((!(_1853_v85).contains(_1835_v69.f33)) ? (_1835_v69.f33) : (_1835_v69.f33));
          let _rhs322 = (_dafny.ZERO).minus(_1835_v69.f33);
          let _lhs279 = _1835_v69;
          let _lhs280 = _1835_v69;
          _1852_v84 = _rhs320;
          _lhs279.f33 = _rhs321;
          _lhs280.f33 = _rhs322;
        }
        let _1854_v86;
        _1854_v86 = _dafny.Set.fromElements(true);
        let _1855_v87;
        _1855_v87 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1856_v88;
        _1856_v88 = _dafny.Set.fromElements(_1855_v87, _1855_v87);
        let _source17 = _module.__default.fm24((_1854_v86).IsSubsetOf(_1854_v86), _1856_v88, (_this.f26).isEqualTo(_1835_v69.f33), globalState);
        if (_source17.is_DC3) {
          let _1857___mcc_h0 = (_source17).cf7;
          let _1858___mcc_h1 = (_source17).cf8;
          let _1859___mcc_h2 = (_source17).cf9;
          let _1860___mcc_h3 = (_source17).cf10;
          let _1861___mcc_h4 = (_source17).cf11;
          let _1862_cf11 = _1861___mcc_h4;
          let _1863_cf10 = _1860___mcc_h3;
          let _1864_cf9 = _1859___mcc_h2;
          let _1865_cf8 = _1858___mcc_h1;
          let _1866_cf7 = _1857___mcc_h0;
          let _1867_v89;
          let _init43 = function (_1868_i11) {
            return _dafny.MultiSet.fromElements((_this).f18);
          };
          let _nw317 = Array((new BigNumber(4)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw317.length); _i0_43++) {
            _nw317[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1867_v89 = _nw317;
          let _1869_v90;
          _1869_v90 = _module.D6.create_DC17(_1867_v89);
          let _1870_v91;
          _1870_v91 = _dafny.Map.Empty.slice().updateUnsafe(_1869_v90,(_this).f18);
          let _1871_v92;
          _1871_v92 = _dafny.Seq.of(_1870_v91);
          let _1872_v93;
          let _nw318 = Array((new BigNumber(22)).toNumber());
          _nw318[(_dafny.ZERO).toNumber()] = _1870_v91;
          _nw318[(_dafny.ONE).toNumber()] = (_1871_v92)[_module.__default.safeIndex(_1864_cf9, new BigNumber((_1871_v92).length))];
          _nw318[(new BigNumber(2)).toNumber()] = _1870_v91;
          _nw318[(new BigNumber(3)).toNumber()] = _1870_v91;
          _nw318[(new BigNumber(4)).toNumber()] = (_1870_v91).Merge(_1870_v91);
          _nw318[(new BigNumber(5)).toNumber()] = (_1870_v91).Merge(_1870_v91);
          _nw318[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1869_v90,p2)).update(_1869_v90, p1);
          _nw318[(new BigNumber(7)).toNumber()] = (_1870_v91).Merge(_1870_v91);
          _nw318[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1869_v90,(_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))]);
          _nw318[(new BigNumber(9)).toNumber()] = (_1870_v91).update(_1869_v90, (_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))]);
          _nw318[(new BigNumber(10)).toNumber()] = (_1871_v92)[_module.__default.safeIndex(_1865_cf8, new BigNumber((_1871_v92).length))];
          _nw318[(new BigNumber(11)).toNumber()] = (_1870_v91).update(_1869_v90, (_this).f18);
          _nw318[(new BigNumber(12)).toNumber()] = (_1870_v91).Merge(_1870_v91);
          _nw318[(new BigNumber(13)).toNumber()] = _1870_v91;
          _nw318[(new BigNumber(14)).toNumber()] = _1870_v91;
          _nw318[(new BigNumber(15)).toNumber()] = (((_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))]) ? (_1870_v91) : (_1870_v91));
          _nw318[(new BigNumber(16)).toNumber()] = (((_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))]) ? (_1870_v91) : (_dafny.Map.Empty.slice().updateUnsafe(_1869_v90,(_this).f18)));
          _nw318[(new BigNumber(17)).toNumber()] = _1870_v91;
          _nw318[(new BigNumber(18)).toNumber()] = _1870_v91;
          _nw318[(new BigNumber(19)).toNumber()] = (_1870_v91).Merge(_1870_v91);
          _nw318[(new BigNumber(20)).toNumber()] = (_1870_v91).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC17(_1867_v89),(_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))]));
          _nw318[(new BigNumber(21)).toNumber()] = (_1870_v91).update(_module.D6.create_DC17(_1867_v89), (_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))]);
          _1872_v93 = _nw318;
          let _1873_v94;
          let _nw319 = new _module.C1();
          _nw319.__ctor((_this).f18);
          _1873_v94 = _nw319;
          let _rhs323 = _1872_v93;
          let _rhs324 = _1855_v87;
          let _rhs325 = p2;
          let _rhs326 = _this.f26;
          let _rhs327 = _1873_v94;
          let _lhs281 = globalState;
          let _lhs282 = globalState;
          _1872_v93 = _rhs323;
          _1855_v87 = _rhs324;
          _lhs281.f2 = _rhs325;
          _lhs282.f15 = _rhs326;
          _1873_v94 = _rhs327;
          let _1874_v95;
          _1874_v95 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1865_cf8);
          let _rhs328 = (_1874_v95).Merge(_1874_v95);
          let _rhs329 = _1864_cf9;
          let _lhs283 = _1835_v69;
          _1874_v95 = _rhs328;
          _lhs283.f33 = _rhs329;
          (globalState).f2 = p2;
          (globalState).f2 = p2;
        } else if (_source17.is_DC4) {
          let _1875___mcc_h5 = (_source17).cf12;
          let _1876_cf12 = _1875___mcc_h5;
          let _1877_v96;
          _1877_v96 = _dafny.MultiSet.fromElements(_1835_v69.f33, new BigNumber(364), new BigNumber(-216), _this.f26);
          _1877_v96 = _module.__default.fm25(((_dafny.ZERO).minus(_1835_v69.f33)).multipliedBy(_1835_v69.f33), globalState);
          let _1878_v97;
          let _init44 = ((_1879_v69) => function (_1880_i12) {
            return _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(new BigNumber(-914)), _module.__default.safeIndex(_1879_v69.f33, new BigNumber((_dafny.Seq.of(new BigNumber(-914))).length)), _1879_v69.f33), (_1879_v69).f34);
          })(_1835_v69);
          let _nw320 = Array((new BigNumber(18)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw320.length); _i0_44++) {
            _nw320[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1878_v97 = _nw320;
          let _1881_v98;
          _1881_v98 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,p1);
          let _1882_v99;
          _1882_v99 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,(((_1881_v98).contains(new BigNumber(564))) ? ((_1881_v98).get(new BigNumber(564))) : (_1876_cf12)));
          let _index294 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1878_v97).length));
          (_1878_v97)[_index294] = _module.__default.fm33((((_1881_v98).contains(new BigNumber(((_this).f19).length))) ? ((_1881_v98).get(new BigNumber(((_this).f19).length))) : (p2)), globalState);
          let _index295 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1878_v97).length));
          (_1878_v97)[_index295] = ((p1) ? (_dafny.Seq.Concat((_1835_v69).f34, (_1835_v69).f34)) : (_dafny.Seq.Concat(_1834_v68, _1834_v68)));
          _1882_v99 = (_1882_v99).update(_this.f26, !(_1835_v69.f33).isEqualTo(new BigNumber(((_this).f19).length)));
          let _1883_v100;
          _1883_v100 = _module.D16.create_DC38(_1877_v96);
          _1883_v100 = _1883_v100;
        } else if (_source17.is_DC2) {
          let _1884___mcc_h6 = (_source17).cf6;
          let _1885_cf6 = _1884___mcc_h6;
          (globalState).f15 = new BigNumber(((_this).f19).length);
          let _index296 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length));
          (_1832_v67)[_index296] = ((_dafny.ZERO).minus(_1835_v69.f33)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm1(_1835_v69.f33, p2, (_this).f18, globalState))));
          let _index297 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length));
          (_1832_v67)[_index297] = !((_module.D9.create_DC23((_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))], p1, new BigNumber(-803))).dtor_cf44);
          let _1886_v101;
          let _nw321 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _1886_v101 = _nw321;
          let _index298 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_1886_v101).length));
          (_1886_v101)[_index298] = _this.f26;
          let _index299 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_1886_v101).length));
          (_1886_v101)[_index299] = _module.__default.safeModuloInt(((!(p2)) ? (_this.f26) : (_this.f26)), _1835_v69.f33);
        } else {
          let _1887___mcc_h7 = (_source17).cf13;
          let _1888_cf13 = _1887___mcc_h7;
          let _1889_v102;
          _1889_v102 = _dafny.Map.Empty.slice().updateUnsafe((_1832_v67)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length))],((true) ? (_1832_v67) : (_1832_v67)));
          _1889_v102 = (_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1832_v67)).Merge(_1889_v102);
          let _1890_v103;
          _1890_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1855_v87,_1835_v69.f33);
          _1890_v103 = (_1890_v103).update(new _dafny.CodePoint('s'.codePointAt(0)), _1835_v69.f33);
          (_1835_v69).m1(_dafny.Seq.Concat((_this).f19, (_this).f19), _this.f26, (_this).f19, globalState);
          let _1891_v104;
          let _nw322 = Array((new BigNumber(10)).toNumber());
          _1891_v104 = _nw322;
          let _index300 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1832_v67).length));
          (_1832_v67)[_index300] = (new BigNumber(696)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f26,_1891_v104)).length));
        }
        r0 = _1835_v69.f33;
      } else {
        let _1892_v105;
        _1892_v105 = new _dafny.CodePoint('g'.codePointAt(0));
        let _1893_v106;
        _1893_v106 = _module.D8.create_DC21(_1892_v105, _this.f26);
        let _1894_v107;
        _1894_v107 = _dafny.MultiSet.fromElements(_1893_v106);
        let _1895_v108;
        _1895_v108 = _dafny.Seq.of(_1894_v107);
        let _1896_v109;
        _1896_v109 = _dafny.Set.fromElements(new BigNumber(-316));
        let _1897_v110;
        _1897_v110 = _module.D18.create_DC45(_1895_v108);
        let _1898_v111;
        _1898_v111 = _dafny.Seq.of(_1895_v108, _dafny.Seq.update(_1895_v108, _module.__default.safeIndex(new BigNumber((_1896_v109).length), new BigNumber((_1895_v108).length)), _1894_v107), _dafny.Seq.Create(_module.__default.abs(new BigNumber(50)), ((_1899_v105, _1900_v106) => function (_1901_i13) {
          return _dafny.MultiSet.fromElements(_module.D8.create_DC21(_1899_v105, _this.f26), _1900_v106);
        })(_1892_v105, _1893_v106)), (_1897_v110).dtor_cf79);
        _1895_v108 = (_1898_v111)[_module.__default.safeIndex((_this.f26).minus(_this.f26), new BigNumber((_1898_v111).length))];
        let _1902_v112;
        let _nw323 = Array((new BigNumber(13)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1902_v112 = _nw323;
        let _index301 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_1902_v112).length));
        (_1902_v112)[_index301] = _1892_v105;
        let _index302 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_1902_v112).length));
        (_1902_v112)[_index302] = _1892_v105;
        if (((_this).f25)[_module.__default.safeIndex(_this.f26, new BigNumber(((_this).f25).length))]) {
          let _1903_v113;
          let _1904_v114;
          let _out56;
          let _out57;
          let _outcollector15 = (_this).m6(globalState);
          _out56 = _outcollector15[0];
          _out57 = _outcollector15[1];
          _1903_v113 = _out56;
          _1904_v114 = _out57;
          let _1905_v115;
          let _nw324 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1905_v115 = _nw324;
          let _index303 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1905_v115).length));
          (_1905_v115)[_index303] = (new BigNumber(998)).isLessThan(_1904_v114);
          let _index304 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1905_v115).length));
          (_1905_v115)[_index304] = p2;
          let _1906_v116;
          _1906_v116 = _dafny.Seq.of(new BigNumber(690), new BigNumber(320), new BigNumber(((_this).f25).length), _this.f26);
          let _index305 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1905_v115).length));
          (_1905_v115)[_index305] = ((_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_1904_v114, new BigNumber(912), _this.f26, _1904_v114), _1906_v116)) ? (p2) : (false));
          let _1907_v117;
          _1907_v117 = _dafny.MultiSet.fromElements(p1);
          let _1908_v118;
          _1908_v118 = _dafny.Seq.of(_1907_v117, _1907_v117);
          let _1909_v119;
          _1909_v119 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,(_1907_v117).IsProperSubsetOf((_1908_v118)[_module.__default.safeIndex(_1904_v114, new BigNumber((_1908_v118).length))]));
          let _index306 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1905_v115).length));
          let _rhs330 = (_1905_v115)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_1905_v115).length))];
          let _rhs331 = new BigNumber((_1909_v119).length);
          let _lhs284 = _1905_v115;
          let _lhs285 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1905_v115).length));
          _lhs284[_lhs285] = _rhs330;
          _1904_v114 = _rhs331;
          let _index307 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1905_v115).length));
          (_1905_v115)[_index307] = (!((_this).f18)) === (_module.__default.fm23(_1896_v109, _1904_v114, globalState));
        } else {
          let _1910_v120;
          _1910_v120 = _dafny.MultiSet.fromElements(p2, false);
          (globalState).f0 = (((_1910_v120).contains((_this).f18)) ? ((_1910_v120).get((_this).f18)) : (new BigNumber(577)));
          let _1911_v121;
          _1911_v121 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2((_this).f18, p2, p1, globalState),p2);
          let _1912_v122;
          let _nw325 = Array((new BigNumber(15)).toNumber());
          _nw325[(_dafny.ZERO).toNumber()] = (_this).f18;
          _nw325[(_dafny.ONE).toNumber()] = ((p2) ? ((_this).f18) : (p2));
          _nw325[(new BigNumber(2)).toNumber()] = _module.__default.fm23(_1896_v109, _this.f26, globalState);
          _nw325[(new BigNumber(3)).toNumber()] = (_this).f18;
          _nw325[(new BigNumber(4)).toNumber()] = p1;
          _nw325[(new BigNumber(5)).toNumber()] = p2;
          _nw325[(new BigNumber(6)).toNumber()] = p1;
          _nw325[(new BigNumber(7)).toNumber()] = (_this).f18;
          _nw325[(new BigNumber(8)).toNumber()] = (_this).f18;
          _nw325[(new BigNumber(9)).toNumber()] = p2;
          _nw325[(new BigNumber(10)).toNumber()] = !((new BigNumber(64)).isLessThanOrEqualTo(_this.f26));
          _nw325[(new BigNumber(11)).toNumber()] = !(true);
          _nw325[(new BigNumber(12)).toNumber()] = (((_1911_v121).contains(p1)) ? ((_1911_v121).get(p1)) : (!(p2)));
          _nw325[(new BigNumber(13)).toNumber()] = p1;
          _nw325[(new BigNumber(14)).toNumber()] = (_1910_v120).contains(p2);
          _1912_v122 = _nw325;
          let _index308 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length));
          (_1912_v122)[_index308] = p1;
          let _index309 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length));
          (_1912_v122)[_index309] = (_this).f18;
          let _index310 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length));
          (_1912_v122)[_index310] = p1;
          let _1913_v123;
          _1913_v123 = _module.D17.create_DC42((_this).f18);
          let _index311 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length));
          let _index312 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length));
          let _rhs332 = (_1912_v122)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length))];
          let _rhs333 = (_1912_v122)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length))];
          let _rhs334 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-542)), ((_1914_v112) => function (_1915_i14) {
            return (_1914_v112)[_module.__default.safeIndex(new BigNumber(521), new BigNumber((_1914_v112).length))];
          })(_1902_v112)), _dafny.Seq.Concat((_this).f19, (_this).f19));
          let _rhs335 = _1913_v123;
          let _rhs336 = (_1912_v122)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length))];
          let _lhs286 = _1912_v122;
          let _lhs287 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length));
          let _lhs288 = globalState;
          let _lhs289 = globalState;
          let _lhs290 = _1912_v122;
          let _lhs291 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_1912_v122).length));
          _lhs286[_lhs287] = _rhs332;
          _lhs288.f2 = _rhs333;
          _lhs289.f12 = _rhs334;
          _1913_v123 = _rhs335;
          _lhs290[_lhs291] = _rhs336;
          let _1916_v124;
          _1916_v124 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,(_this).f19);
          (globalState).f0 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_this.f26), new BigNumber((_1916_v124).length));
        }
        let _1917_v125;
        _1917_v125 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,p2);
        let _1918_v126;
        _1918_v126 = _dafny.Seq.of(new BigNumber(171), new BigNumber((_1917_v125).length));
        let _1919_v127;
        _1919_v127 = _dafny.Map.Empty.slice().updateUnsafe(_1918_v126,!(p1));
        _1919_v127 = _1919_v127;
        let _1920_v128;
        let _nw326 = new _module.C1();
        _nw326.__ctor(p2);
        _1920_v128 = _nw326;
        let _1921_v129;
        _1921_v129 = _dafny.Map.Empty.slice().updateUnsafe(_1920_v128,p1);
        let _1922_v130;
        _1922_v130 = _module.D11.create_DC27(((_1921_v129).update(_1920_v128, p2)).Merge(_1921_v129));
        _1922_v130 = _module.D11.create_DC27(_1921_v129);
      }
      (globalState).f9 = _dafny.Seq.Concat((_this).f19, (_this).f19);
      if ((_this).f18) {
        let _1923_v131;
        let _nw327 = Array((new BigNumber(20)).toNumber()).fill(false);
        _1923_v131 = _nw327;
        let _index313 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1923_v131).length));
        (_1923_v131)[_index313] = p1;
        let _index314 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1923_v131).length));
        (_1923_v131)[_index314] = p2;
        let _1924_v132;
        let _init45 = function (_1925_i15) {
          return (_1925_i15).minus(new BigNumber((_dafny.Seq.of(_this.f26)).length));
        };
        let _nw328 = Array((new BigNumber(26)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw328.length); _i0_45++) {
          _nw328[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1924_v132 = _nw328;
        let _index315 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1924_v132).length));
        (_1924_v132)[_index315] = _module.__default.safeModuloInt(_this.f26, new BigNumber(536));
        let _1926_v133;
        _1926_v133 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18),(_dafny.ZERO).minus(_this.f26));
        let _index316 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1924_v132).length));
        let _index317 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1923_v131).length));
        let _rhs337 = (_this.f26).isLessThan(_this.f26);
        let _rhs338 = _this.f26;
        let _rhs339 = (_this).fm2(p2, (new BigNumber(-148)).isLessThan(new BigNumber((_1926_v133).length)), ((_dafny.ZERO).minus(_module.__default.fm1(new BigNumber((_dafny.Seq.of(_this.f26)).length), (_1923_v131)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_1923_v131).length))], (_1923_v131)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_1923_v131).length))], globalState))).isLessThan(_this.f26), globalState);
        let _lhs292 = globalState;
        let _lhs293 = _1924_v132;
        let _lhs294 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1924_v132).length));
        let _lhs295 = _1923_v131;
        let _lhs296 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1923_v131).length));
        _lhs292.f2 = _rhs337;
        _lhs293[_lhs294] = _rhs338;
        _lhs295[_lhs296] = _rhs339;
        let _1927_v134;
        _1927_v134 = _dafny.Set.fromElements((_1923_v131)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_1923_v131).length))]);
        let _1928_v135;
        _1928_v135 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1927_v134).IsDisjointFrom(_1927_v134));
        _1928_v135 = (_1928_v135).update(p2, (false) && (p2));
        let _1929_v136;
        let _nw329 = Array((new BigNumber(20)).toNumber()).fill([]);
        _1929_v136 = _nw329;
        let _1930_v137;
        _1930_v137 = _1924_v132;
        let _index318 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1929_v136).length));
        (_1929_v136)[_index318] = ((false) ? ((_1930_v137)) : (_1924_v132));
        let _index319 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1929_v136).length));
        (_1929_v136)[_index319] = _1924_v132;
        let _1931_v138;
        _1931_v138 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1932_v139;
        _1932_v139 = _dafny.Set.fromElements(_1931_v138);
        r0 = ((p1) ? (_module.__default.safeDivisionInt((_1924_v132)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_1924_v132).length))], new BigNumber((_1932_v139).length))) : (_this.f26));
      } else {
        (globalState).f0 = _this.f26;
        let _1933_v140;
        _1933_v140 = _dafny.MultiSet.fromElements((_this).f18);
        let _1934_v141;
        _1934_v141 = _dafny.MultiSet.fromElements(_1933_v140, _1933_v140, _1933_v140, _1933_v140);
        let _1935_v142;
        _1935_v142 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,p1);
        let _1936_v143;
        _1936_v143 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm49(new BigNumber(((_module.D19.create_DC47(_1934_v141)).dtor_cf81).cardinality()), new BigNumber((_1935_v142).length), _this.f26, globalState)).length), new BigNumber((_dafny.MultiSet.fromElements((_this).f18, p1)).cardinality()), _this.f26);
        let _1937_v144;
        _1937_v144 = _dafny.Seq.of(_this.f26);
        let _1938_v145;
        _1938_v145 = _dafny.Set.fromElements(p1, p2);
        let _1939_v146;
        _1939_v146 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1940_v147;
        _1940_v147 = _dafny.Map.Empty.slice().updateUnsafe(_1939_v146,_1939_v146);
        let _1941_v148;
        _1941_v148 = _module.D13.create_DC33(false, _this.f26, (_this).f18, new BigNumber(((_this).f19).length));
        let _1942_v149;
        _1942_v149 = _dafny.Map.Empty.slice().updateUnsafe((_1941_v148).dtor_cf55,_this.f26);
        let _1943_v150;
        _1943_v150 = _dafny.Map.Empty.slice().updateUnsafe(_1940_v147,new BigNumber((_1942_v149).length));
        let _1944_v151;
        let _nw330 = Array((new BigNumber(27)).toNumber());
        _nw330[(_dafny.ZERO).toNumber()] = !(p2);
        _nw330[(_dafny.ONE).toNumber()] = !(!((_this).f18));
        _nw330[(new BigNumber(2)).toNumber()] = p1;
        _nw330[(new BigNumber(3)).toNumber()] = (_this).f18;
        _nw330[(new BigNumber(4)).toNumber()] = (_1936_v143).equals(_dafny.MultiSet.fromElements(_this.f26));
        _nw330[(new BigNumber(5)).toNumber()] = p2;
        _nw330[(new BigNumber(6)).toNumber()] = !((_this).f18);
        _nw330[(new BigNumber(7)).toNumber()] = p2;
        _nw330[(new BigNumber(8)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1937_v144, _1937_v144);
        _nw330[(new BigNumber(9)).toNumber()] = p2;
        _nw330[(new BigNumber(10)).toNumber()] = (_this).f18;
        _nw330[(new BigNumber(11)).toNumber()] = (_1938_v145).IsProperSubsetOf(_1938_v145);
        _nw330[(new BigNumber(12)).toNumber()] = (_1938_v145).IsDisjointFrom(_module.__default.fm36(p1, new BigNumber((_1943_v150).length), false, globalState));
        _nw330[(new BigNumber(13)).toNumber()] = (_this).f18;
        _nw330[(new BigNumber(14)).toNumber()] = p2;
        _nw330[(new BigNumber(15)).toNumber()] = p1;
        _nw330[(new BigNumber(16)).toNumber()] = false;
        _nw330[(new BigNumber(17)).toNumber()] = p1;
        _nw330[(new BigNumber(18)).toNumber()] = p2;
        _nw330[(new BigNumber(19)).toNumber()] = p1;
        _nw330[(new BigNumber(20)).toNumber()] = p1;
        _nw330[(new BigNumber(21)).toNumber()] = p1;
        _nw330[(new BigNumber(22)).toNumber()] = p1;
        _nw330[(new BigNumber(23)).toNumber()] = p2;
        _nw330[(new BigNumber(24)).toNumber()] = !(_1933_v140).contains(p2);
        _nw330[(new BigNumber(25)).toNumber()] = (_this.f26).isLessThan(_this.f26);
        _nw330[(new BigNumber(26)).toNumber()] = (_this).f18;
        _1944_v151 = _nw330;
        let _1945_v152;
        _1945_v152 = _dafny.Set.fromElements(_this.f26);
        let _index320 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1944_v151).length));
        (_1944_v151)[_index320] = !(_module.__default.fm23(_1945_v152, _this.f26, globalState));
        let _index321 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1944_v151).length));
        (_1944_v151)[_index321] = !(p1);
        let _1946_v153;
        _1946_v153 = _dafny.Seq.of((_this).f25);
        let _1947_v154;
        _1947_v154 = _dafny.Set.fromElements((_1946_v153)[_module.__default.safeIndex(_this.f26, new BigNumber((_1946_v153).length))]);
        let _1948_v155;
        _1948_v155 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f26).isEqualTo(_this.f26),_1947_v154);
        _1948_v155 = (_1948_v155).update(!((_1944_v151)[_module.__default.safeIndex(new BigNumber(61), new BigNumber((_1944_v151).length))]), _1947_v154);
        let _1949_v156;
        let _nw331 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1949_v156 = _nw331;
        let _1950_v157;
        _1950_v157 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rfnihgwpf"),(_1941_v148).dtor_cf58);
        let _index322 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_1949_v156).length));
        (_1949_v156)[_index322] = new BigNumber((_1950_v157).length);
        let _index323 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_1949_v156).length));
        (_1949_v156)[_index323] = new BigNumber((_module.__default.fm58((_this.f26).isLessThanOrEqualTo(_this.f26), p1, globalState)).length);
        (globalState).f2 = (((_1936_v143).IsSubsetOf(_1936_v143)) ? ((true) === ((_this).f18)) : ((_1944_v151)[_module.__default.safeIndex(new BigNumber(61), new BigNumber((_1944_v151).length))]));
      }
      if ((_this).fm3(_this.f26, globalState)) {
        (globalState).f0 = _this.f26;
        let _1951_v158;
        _1951_v158 = new _dafny.CodePoint('g'.codePointAt(0));
        _1951_v158 = _1951_v158;
        let _1952_v159;
        _1952_v159 = _dafny.Set.fromElements(_this.f26);
        _1952_v159 = function () {
          let _coll49 = new _dafny.Set();
          for (const _compr_49 of _dafny.IntegerRange(new BigNumber(981), new BigNumber(989))) {
            let _1953_v160 = _compr_49;
            if (((new BigNumber(981)).isLessThanOrEqualTo(_1953_v160)) && ((_1953_v160).isLessThan(new BigNumber(989)))) {
              _coll49.add((_1953_v160).multipliedBy(_this.f26));
            }
          }
          return _coll49;
        }();
        let _1954_v161;
        _1954_v161 = _dafny.Seq.of(_this.f26);
        let _1955_v162;
        _1955_v162 = _dafny.Seq.of((_this).f19);
        let _1956_v163;
        let _nw332 = new _module.C3();
        _nw332.__ctor(_1954_v161, _dafny.Seq.update((_1955_v162)[_module.__default.safeIndex(new BigNumber((_1952_v159).length), new BigNumber((_1955_v162).length))], _module.__default.safeIndex(new BigNumber(4), new BigNumber(((_1955_v162)[_module.__default.safeIndex(new BigNumber((_1952_v159).length), new BigNumber((_1955_v162).length))]).length)), _1951_v158), (_this).f20, (new BigNumber(290)).isLessThanOrEqualTo(_this.f26));
        _1956_v163 = _nw332;
        let _1957_v164;
        let _init46 = function (_1958_i16) {
          return (_1958_i16).minus(_this.f26);
        };
        let _nw333 = Array((new BigNumber(26)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw333.length); _i0_46++) {
          _nw333[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1957_v164 = _nw333;
        let _1959_v165;
        _1959_v165 = _dafny.MultiSet.fromElements((_this).f18);
        let _rhs340 = _1957_v164;
        let _rhs341 = _1959_v165;
        _1957_v164 = _rhs340;
        _1959_v165 = _rhs341;
      } else {
        let _1960_v166;
        _1960_v166 = _module.D13.create_DC33(p1, _this.f26, p2, _this.f26);
        let _1961_v167;
        _1961_v167 = _dafny.MultiSet.fromElements(p1, (_this).f18);
        let _1962_v168;
        let _nw334 = Array((new BigNumber(20)).toNumber());
        _nw334[(_dafny.ZERO).toNumber()] = (_1960_v166).dtor_cf57;
        _nw334[(_dafny.ONE).toNumber()] = ((p1) ? (false) : ((_this).f18));
        _nw334[(new BigNumber(2)).toNumber()] = p1;
        _nw334[(new BigNumber(3)).toNumber()] = p2;
        _nw334[(new BigNumber(4)).toNumber()] = !((_this).f18);
        _nw334[(new BigNumber(5)).toNumber()] = p2;
        _nw334[(new BigNumber(6)).toNumber()] = p1;
        _nw334[(new BigNumber(7)).toNumber()] = (_this).f18;
        _nw334[(new BigNumber(8)).toNumber()] = p1;
        _nw334[(new BigNumber(9)).toNumber()] = (_this).f18;
        _nw334[(new BigNumber(10)).toNumber()] = !(false) || ((_this).f18);
        _nw334[(new BigNumber(11)).toNumber()] = !((_this).f18);
        _nw334[(new BigNumber(12)).toNumber()] = p1;
        _nw334[(new BigNumber(13)).toNumber()] = p1;
        _nw334[(new BigNumber(14)).toNumber()] = p1;
        _nw334[(new BigNumber(15)).toNumber()] = p1;
        _nw334[(new BigNumber(16)).toNumber()] = !(!(_1961_v167).equals(_1961_v167));
        _nw334[(new BigNumber(17)).toNumber()] = (_this).f18;
        _nw334[(new BigNumber(18)).toNumber()] = (_this).fm3(new BigNumber(204), globalState);
        _nw334[(new BigNumber(19)).toNumber()] = p1;
        _1962_v168 = _nw334;
        let _index324 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_1962_v168).length));
        (_1962_v168)[_index324] = p2;
        let _index325 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_1962_v168).length));
        (_1962_v168)[_index325] = (_this).fm2(true, p2, p1, globalState);
        r0 = new BigNumber(487);
        let _1963_v169;
        _1963_v169 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1964_v170;
        _1964_v170 = _dafny.Seq.of((_this).f19, (_this).f19);
        let _1965_v171;
        let _nw335 = Array((new BigNumber(26)).toNumber());
        _nw335[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("xlqtclqmw");
        _nw335[(_dafny.ONE).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(2)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(610)), function (_1966_i17) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), _module.__default.safeIndex(_this.f26, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(610)), function (_1967_i17) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        })).length)), _1963_v169);
        _nw335[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update((_this).f19, _module.__default.safeIndex(_this.f26, new BigNumber(((_this).f19).length)), _1963_v169), (_this).f19);
        _nw335[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("rskscm");
        _nw335[(new BigNumber(6)).toNumber()] = _module.__default.fm48(globalState);
        _nw335[(new BigNumber(7)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(618)), ((_1968_v169) => function (_1969_i18) {
          return _1968_v169;
        })(_1963_v169));
        _nw335[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), ((_1970_v169) => function (_1971_i19) {
          return _1970_v169;
        })(_1963_v169)), (_this).f19);
        _nw335[(new BigNumber(10)).toNumber()] = (_1964_v170)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1964_v170).length))];
        _nw335[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat((_this).f19, _dafny.Seq.UnicodeFromString("jfpoexew"));
        _nw335[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("qkvw");
        _nw335[(new BigNumber(13)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(14)).toNumber()] = _module.__default.fm16(_this.f26, false, _1963_v169, _this.f26, globalState);
        _nw335[(new BigNumber(15)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(16)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(17)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(18)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat((_this).f19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(855)), function (_1972_i20) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        }));
        _nw335[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), ((_1973_v169) => function (_1974_i21) {
          return _1973_v169;
        })(_1963_v169));
        _nw335[(new BigNumber(21)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update((_this).f19, _module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber(((_this).f19).length)), _1963_v169), (_this).f19);
        _nw335[(new BigNumber(23)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(24)).toNumber()] = (_this).f19;
        _nw335[(new BigNumber(25)).toNumber()] = (_this).f19;
        _1965_v171 = _nw335;
        let _index326 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1965_v171).length));
        (_1965_v171)[_index326] = (_this).f19;
        let _index327 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1965_v171).length));
        let _rhs342 = false;
        let _rhs343 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f19, (_this).f19), _dafny.Seq.UnicodeFromString("mrmigu"));
        let _rhs344 = (_this.f26).isLessThan((_dafny.ZERO).minus(_this.f26));
        let _rhs345 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), ((_1975_v169) => function (_1976_i22) {
          return _1975_v169;
        })(_1963_v169)), (_this).f19);
        let _rhs346 = _1963_v169;
        let _lhs297 = globalState;
        let _lhs298 = globalState;
        let _lhs299 = globalState;
        let _lhs300 = _1965_v171;
        let _lhs301 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1965_v171).length));
        _lhs297.f2 = _rhs342;
        _lhs298.f9 = _rhs343;
        _lhs299.f2 = _rhs344;
        _lhs300[_lhs301] = _rhs345;
        _1963_v169 = _rhs346;
        _1963_v169 = _1963_v169;
        (globalState).f15 = (_module.__default.fm59(globalState)).dtor_cf33;
      }
      r0 = _this.f26;
      return r0;
    }
    m6(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _1977_v0;
      let _nw336 = new _module.C1();
      _nw336.__ctor((_this).f18);
      _1977_v0 = _nw336;
      let _1978_v1;
      _1978_v1 = _dafny.Set.fromElements(_1977_v0);
      let _1979_v2;
      _1979_v2 = _dafny.Seq.of(new BigNumber(((_this).f19).length), new BigNumber((_1978_v1).length));
      let _1980_v3;
      _1980_v3 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm33((_1977_v0).f18, globalState),_1979_v2);
      let _1981_v4;
      let _nw337 = Array((new BigNumber(12)).toNumber());
      _nw337[(_dafny.ZERO).toNumber()] = _1979_v2;
      _nw337[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1979_v2, _1979_v2);
      _nw337[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f26, _this.f26, _this.f26, new BigNumber(((_this).f19).length)), _1979_v2);
      _nw337[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat((((_1980_v3).contains(_1979_v2)) ? ((_1980_v3).get(_1979_v2)) : (_1979_v2)), _1979_v2);
      _nw337[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(964)), function (_1982_i0) {
        return _this.f26;
      }), _1979_v2);
      _nw337[(new BigNumber(5)).toNumber()] = _1979_v2;
      _nw337[(new BigNumber(6)).toNumber()] = _1979_v2;
      _nw337[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(new BigNumber(((_this).f19).length), _this.f26, _this.f26);
      _nw337[(new BigNumber(8)).toNumber()] = _1979_v2;
      _nw337[(new BigNumber(9)).toNumber()] = (((_1977_v0).f18) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(696)), function (_1983_i1) {
        return _this.f26;
      })) : (_dafny.Seq.update(_1979_v2, _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1979_v2).length)), _this.f26)));
      _nw337[(new BigNumber(10)).toNumber()] = _1979_v2;
      _nw337[(new BigNumber(11)).toNumber()] = _1979_v2;
      _1981_v4 = _nw337;
      _1981_v4 = _1981_v4;
      let _hi10 = _this.f26;
      for (let _1984_i2 = new BigNumber(((((_1977_v0).f18) ? ((_this).f19) : ((_this).f19))).length); _1984_i2.isLessThan(_hi10); _1984_i2 = _1984_i2.plus(_dafny.ONE)) {
        r0 = (((_this).f18) ? ((_this).f18) : (((_1977_v0).f18) === ((_1977_v0).f18)));
        (globalState).f0 = _this.f26;
        let _1985_v5;
        _1985_v5 = new _dafny.CodePoint('a'.codePointAt(0));
        (globalState).f12 = _dafny.Seq.Concat((_this).f19, _dafny.Seq.update((_this).f19, _module.__default.safeIndex(new BigNumber(538), new BigNumber(((_this).f19).length)), _1985_v5));
        r0 = (_this).fm2((_1977_v0).f18, (_this).f18, (_1984_i2).isLessThanOrEqualTo(_this.f26), globalState);
      }
      if ((_this).f18) {
        let _1986_v6;
        let _init47 = function (_1987_i3) {
          return _dafny.Seq.IsProperPrefixOf((_this).f25, (_this).f25);
        };
        let _nw338 = Array((new BigNumber(22)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw338.length); _i0_47++) {
          _nw338[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1986_v6 = _nw338;
        let _1988_v7;
        _1988_v7 = _dafny.Set.fromElements(_this.f26);
        let _index328 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1986_v6).length));
        (_1986_v6)[_index328] = _module.__default.fm23(_1988_v7, new BigNumber(((_this).f19).length), globalState);
        let _index329 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1986_v6).length));
        (_1986_v6)[_index329] = (_this).f18;
        let _index330 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1986_v6).length));
        (_1986_v6)[_index330] = !(!(_dafny.Seq.IsProperPrefixOf((_this).f19, (_this).f19)));
        (globalState).f9 = _dafny.Seq.Concat((_this).f19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), function (_1989_i4) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }));
        let _1990_v8;
        let _nw339 = Array((new BigNumber(26)).toNumber());
        _1990_v8 = _nw339;
        _1990_v8 = _1990_v8;
        if ((_1977_v0).f18) {
          let _1991_v9;
          _1991_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(532),_this.f26);
          let _1992_v10;
          _1992_v10 = _dafny.Set.fromElements((_module.D20.create_DC52(_1991_v9)).dtor_cf86);
          let _1993_v12;
          _1993_v12 = _module.D21.create_DC55(function () {
  let _coll50 = new _dafny.Set();
  for (const _compr_50 of ((_this).f20).Elements) {
    let _1994_v11 = _compr_50;
    if (_dafny.Seq.contains((_this).f20, _1994_v11)) {
      _coll50.add(_1994_v11);
    }
  }
  return _coll50;
}());
          _1992_v10 = (_1993_v12).dtor_cf91;
          let _1995_v13;
          _1995_v13 = _dafny.MultiSet.fromElements((_this).f18);
          (globalState).f3 = new BigNumber(((((_1977_v0).f18) ? (_dafny.MultiSet.fromElements((_1977_v0).f18)) : ((_1995_v13).Intersect(_dafny.MultiSet.FromArray((_this).f25))))).cardinality());
          let _index331 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1986_v6).length));
          (_1986_v6)[_index331] = !(!(!((_this).f18)));
          let _1996_v14;
          let _nw340 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _1996_v14 = _nw340;
          let _1997_v15;
          _1997_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f19).length),_1991_v9);
          let _index332 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_1996_v14).length));
          (_1996_v14)[_index332] = new BigNumber(((((_1997_v15).contains(_this.f26)) ? ((_1997_v15).get(_this.f26)) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f26,new BigNumber(888))))).length);
          let _index333 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_1996_v14).length));
          (_1996_v14)[_index333] = (_this.f26).plus(_this.f26);
          let _index334 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_1996_v14).length));
          (_1996_v14)[_index334] = (_1996_v14)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_1996_v14).length))];
        } else {
          let _1998_v16;
          _1998_v16 = _dafny.Seq.of(_dafny.Seq.update((_this).f25, _module.__default.safeIndex(_this.f26, new BigNumber(((_this).f25).length)), (_1977_v0).f18));
          let _1999_v17;
          _1999_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1998_v16,_this.f26);
          _1999_v17 = (_1999_v17).update(_1998_v16, _this.f26);
          let _2000_v18;
          _2000_v18 = _dafny.Map.Empty.slice().updateUnsafe((_1986_v6)[_module.__default.safeIndex(new BigNumber(110), new BigNumber((_1986_v6).length))],(_this).f18);
          let _2001_v19;
          _2001_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2000_v18).length),_this.f26);
          let _2002_v20;
          _2002_v20 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.update(_1979_v2, _module.__default.safeIndex(new BigNumber((_2001_v19).length), new BigNumber((_1979_v2).length)), _this.f26), _module.__default.safeIndex(_this.f26, new BigNumber((_dafny.Seq.update(_1979_v2, _module.__default.safeIndex(new BigNumber((_2001_v19).length), new BigNumber((_1979_v2).length)), _this.f26)).length)), _this.f26), _1979_v2);
          let _2003_v21;
          let _nw341 = new _module.C3();
          _nw341.__ctor((_2002_v20)[_module.__default.safeIndex(_this.f26, new BigNumber((_2002_v20).length))], _dafny.Seq.Concat((_this).f19, (_this).f19), (_this).f20, (_this).f18);
          _2003_v21 = _nw341;
          let _2004_v22;
          let _nw342 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _2004_v22 = _nw342;
          let _index335 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_2004_v22).length));
          (_2004_v22)[_index335] = _this.f26;
          let _index336 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_2004_v22).length));
          (_2004_v22)[_index336] = new BigNumber(((_this).f19).length);
          (globalState).f2 = (_1977_v0).f18;
          (globalState).f3 = _this.f26;
        }
      } else {
        let _2005_v23;
        _2005_v23 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18);
        let _2006_v24;
        _2006_v24 = _dafny.Seq.of((_this).f18);
        let _2007_v25;
        let _nw343 = Array((new BigNumber(24)).toNumber()).fill([]);
        _2007_v25 = _nw343;
        let _2008_v26;
        _2008_v26 = _2005_v23;
        let _2009_v27;
        _2009_v27 = new _dafny.CodePoint('b'.codePointAt(0));
        let _rhs347 = (_1977_v0).f18;
        let _rhs348 = (_2008_v26);
        let _rhs349 = (_this.f26).isLessThan(_this.f26);
        let _rhs350 = _module.__default.fm44((_this.f26).minus(_this.f26), (_1977_v0).f18, _2009_v27, new BigNumber(-143), globalState);
        let _rhs351 = _2007_v25;
        let _lhs302 = globalState;
        let _lhs303 = globalState;
        _lhs302.f2 = _rhs347;
        _2005_v23 = _rhs348;
        _lhs303.f2 = _rhs349;
        _2006_v24 = _rhs350;
        _2007_v25 = _rhs351;
        (globalState).f2 = (_1977_v0).f18;
        let _2010_v28;
        _2010_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,_this.f26);
        (globalState).f15 = (_this.f26).plus(_module.__default.safeDivisionInt(new BigNumber((_1979_v2).length), new BigNumber((_2010_v28).length)));
        if (false) {
          let _2011_v29;
          let _init48 = function (_2012_i5) {
            return (_2012_i5).multipliedBy(_this.f26);
          };
          let _nw344 = Array((new BigNumber(10)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw344.length); _i0_48++) {
            _nw344[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _2011_v29 = _nw344;
          let _index337 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2011_v29).length));
          (_2011_v29)[_index337] = _this.f26;
          let _2013_v30;
          let _nw345 = Array((new BigNumber(21)).toNumber()).fill(false);
          _2013_v30 = _nw345;
          let _2014_v31;
          _2014_v31 = _dafny.Set.fromElements(_2013_v30, _2013_v30);
          let _2015_v32;
          _2015_v32 = _dafny.MultiSet.fromElements(_this.f26, _this.f26);
          let _index338 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2011_v29).length));
          let _rhs352 = (_1977_v0).f18;
          let _rhs353 = _this.f26;
          let _rhs354 = (_2014_v31).IsDisjointFrom(_2014_v31);
          let _rhs355 = !((new BigNumber((_2015_v32).cardinality())).isEqualTo(_this.f26));
          let _rhs356 = _module.__default.safeModuloInt((_this.f26).plus((_1979_v2)[_module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber((_1979_v2).length))]), _this.f26);
          let _lhs304 = globalState;
          let _lhs305 = _2011_v29;
          let _lhs306 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2011_v29).length));
          let _lhs307 = globalState;
          let _lhs308 = globalState;
          let _lhs309 = globalState;
          _lhs304.f2 = _rhs352;
          _lhs305[_lhs306] = _rhs353;
          _lhs307.f2 = _rhs354;
          _lhs308.f2 = _rhs355;
          _lhs309.f3 = _rhs356;
          let _2016_v33;
          _2016_v33 = _2013_v30;
          let _2017_v34;
          _2017_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2016_v33,(_1977_v0).f18);
          r0 = !((((_2017_v34).contains(_2013_v30)) ? ((_2017_v34).get(_2013_v30)) : ((_1977_v0).f18)));
          let _2018_v35;
          let _nw346 = Array((new BigNumber(14)).toNumber());
          _nw346[(_dafny.ZERO).toNumber()] = _2009_v27;
          _nw346[(_dafny.ONE).toNumber()] = ((_this).f19)[_module.__default.safeIndex((_2011_v29)[_module.__default.safeIndex(new BigNumber(31), new BigNumber((_2011_v29).length))], new BigNumber(((_this).f19).length))];
          _nw346[(new BigNumber(2)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(3)).toNumber()] = ((false) ? (_2009_v27) : (_2009_v27));
          _nw346[(new BigNumber(4)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(5)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(6)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(7)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(8)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(9)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(10)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(11)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(12)).toNumber()] = _2009_v27;
          _nw346[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _2018_v35 = _nw346;
          let _index339 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2018_v35).length));
          (_2018_v35)[_index339] = ((_this).f19)[_module.__default.safeIndex((_2011_v29)[_module.__default.safeIndex(new BigNumber(31), new BigNumber((_2011_v29).length))], new BigNumber(((_this).f19).length))];
          let _index340 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2018_v35).length));
          (_2018_v35)[_index340] = _module.__default.fm46(globalState);
          let _rhs357 = _2011_v29;
          let _rhs358 = _2009_v27;
          _2011_v29 = _rhs357;
          _2009_v27 = _rhs358;
          let _index341 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2011_v29).length));
          (_2011_v29)[_index341] = _module.__default.fm1(((_2011_v29)[_module.__default.safeIndex(new BigNumber(31), new BigNumber((_2011_v29).length))]).multipliedBy(new BigNumber(623)), false, (_1977_v0).f18, globalState);
        } else {
          (globalState).f3 = _this.f26;
          let _2019_v36;
          _2019_v36 = _module.D12.create_DC30((_this).f25);
          let _2020_v37;
          _2020_v37 = _dafny.Seq.of(_2019_v36);
          let _2021_v38;
          _2021_v38 = _dafny.Set.fromElements(_this.f26);
          let _2022_v39;
          _2022_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,_2021_v38);
          let _2023_v40;
          _2023_v40 = _dafny.Map.Empty.slice().updateUnsafe((_1977_v0).f18,_this.f26);
          let _2024_v41;
          _2024_v41 = _dafny.Set.fromElements((_module.D19.create_DC50((_1977_v0).f18)).dtor_cf84, (_1977_v0).f18);
          let _2025_v42;
          _2025_v42 = _dafny.Map.Empty.slice().updateUnsafe((_1977_v0).f18,_1979_v2);
          let _2026_v43;
          _2026_v43 = _dafny.MultiSet.fromElements(_this.f26, new BigNumber((_2024_v41).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-538)), function (_2027_i6) {
            return new BigNumber(-917);
          })).length));
          let _2028_v44;
          let _nw347 = Array((new BigNumber(27)).toNumber());
          _nw347[(_dafny.ZERO).toNumber()] = _this.f26;
          _nw347[(_dafny.ONE).toNumber()] = new BigNumber((_2022_v39).length);
          _nw347[(new BigNumber(2)).toNumber()] = _this.f26;
          _nw347[(new BigNumber(3)).toNumber()] = new BigNumber(302);
          _nw347[(new BigNumber(4)).toNumber()] = new BigNumber((_2023_v40).length);
          _nw347[(new BigNumber(5)).toNumber()] = _this.f26;
          _nw347[(new BigNumber(6)).toNumber()] = _module.__default.fm1(new BigNumber((_2024_v41).length), (_1977_v0).f18, !((_1977_v0).f18), globalState);
          _nw347[(new BigNumber(7)).toNumber()] = _this.f26;
          _nw347[(new BigNumber(8)).toNumber()] = _this.f26;
          _nw347[(new BigNumber(9)).toNumber()] = _module.__default.fm1(_this.f26, (_1977_v0).f18, (_1977_v0).f18, globalState);
          _nw347[(new BigNumber(10)).toNumber()] = _this.f26;
          _nw347[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_2024_v41).length), _this.f26));
          _nw347[(new BigNumber(12)).toNumber()] = _this.f26;
          _nw347[(new BigNumber(13)).toNumber()] = (((_1977_v0).f18) ? (_this.f26) : (_this.f26));
          _nw347[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2010_v28).length), _this.f26);
          _nw347[(new BigNumber(15)).toNumber()] = (_this.f26).multipliedBy(_this.f26);
          _nw347[(new BigNumber(16)).toNumber()] = (_this.f26).multipliedBy(_this.f26);
          _nw347[(new BigNumber(17)).toNumber()] = _this.f26;
          _nw347[(new BigNumber(18)).toNumber()] = _this.f26;
          _nw347[(new BigNumber(19)).toNumber()] = new BigNumber((_2006_v24).length);
          _nw347[(new BigNumber(20)).toNumber()] = new BigNumber(((((_2025_v42).contains((_1977_v0).f18)) ? ((_2025_v42).get((_1977_v0).f18)) : (_dafny.Seq.of(new BigNumber((_2026_v43).cardinality()))))).length);
          _nw347[(new BigNumber(21)).toNumber()] = (((_this).f18) ? (_this.f26) : (_this.f26));
          _nw347[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(_this.f26);
          _nw347[(new BigNumber(23)).toNumber()] = new BigNumber(((_this).f25).length);
          _nw347[(new BigNumber(24)).toNumber()] = new BigNumber(822);
          _nw347[(new BigNumber(25)).toNumber()] = new BigNumber((_module.__default.fm32(globalState)).cardinality());
          _nw347[(new BigNumber(26)).toNumber()] = _this.f26;
          _2028_v44 = _nw347;
          let _2029_v45;
          _2029_v45 = _dafny.Map.Empty.slice().updateUnsafe((_1977_v0).f18,(_1977_v0).f18);
          let _2030_v46;
          _2030_v46 = _dafny.MultiSet.fromElements(_2028_v44, _2028_v44, _2028_v44);
          let _rhs359 = _2020_v37;
          let _rhs360 = (_this).f18;
          let _rhs361 = (_1977_v0).f18;
          let _rhs362 = (((_2029_v45).contains((_2030_v46).IsDisjointFrom(_2030_v46))) ? ((_2029_v45).get((_2030_v46).IsDisjointFrom(_2030_v46))) : ((_1977_v0).f18));
          let _rhs363 = _2028_v44;
          let _lhs310 = globalState;
          let _lhs311 = globalState;
          let _lhs312 = globalState;
          _2020_v37 = _rhs359;
          _lhs310.f2 = _rhs360;
          _lhs311.f2 = _rhs361;
          _lhs312.f2 = _rhs362;
          _2028_v44 = _rhs363;
          let _2031_v47;
          _2031_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("o"),_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm35(_this.f26, (_this).f18, globalState),_this.f26));
          let _2032_v48;
          _2032_v48 = _module.D0.create_DC1((_this).f18, _this.f26, (_this).f19, (_1977_v0).f18, _this.f26);
          let _2033_v49;
          _2033_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2032_v48,new BigNumber((_2006_v24).length));
          _2031_v47 = (_2031_v47).update((_this).f19, _2033_v49);
          let _2034_v50;
          let _nw348 = new _module.C0();
          _nw348.__ctor(_this.f26, ((_dafny.ZERO).minus(_this.f26)).minus(_module.__default.fm1(_this.f26, (_1977_v0).f18, !((_1977_v0).f18), globalState)));
          _2034_v50 = _nw348;
          let _2035_v51;
          let _init49 = ((_2036_v24) => function (_2037_i7) {
            return _2036_v24;
          })(_2006_v24);
          let _nw349 = Array((new BigNumber(28)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw349.length); _i0_49++) {
            _nw349[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _2035_v51 = _nw349;
          let _2038_v52;
          _2038_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(825),_2035_v51);
          _2038_v52 = (_2038_v52).update(((_2034_v50).f32).plus(_this.f26), _2035_v51);
        }
        let _2039_v53;
        _2039_v53 = _dafny.Set.fromElements(new BigNumber((_1979_v2).length));
        let _2040_v54;
        _2040_v54 = _dafny.Seq.of((_this).f19);
        let _2041_v55;
        _2041_v55 = _dafny.MultiSet.fromElements((_2040_v54)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f26), new BigNumber((_2040_v54).length))], (_this).f19);
        _2006_v24 = _module.__default.fm44(_this.f26, _module.__default.fm23(_2039_v53, _this.f26, globalState), _2009_v27, (_1979_v2)[_module.__default.safeIndex((((_2041_v55).contains(_dafny.Seq.of(_2009_v27))) ? ((_2041_v55).get(_dafny.Seq.of(_2009_v27))) : (_module.__default.fm1(_this.f26, (_1977_v0).f18, (_1977_v0).f18, globalState))), new BigNumber((_1979_v2).length))], globalState);
      }
      let _2042_v56;
      let _init50 = function (_2043_i8) {
        return (_this).f19;
      };
      let _nw350 = Array((new BigNumber(12)).toNumber());
      for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw350.length); _i0_50++) {
        _nw350[_i0_50] = _init50(new BigNumber(_i0_50));
      }
      _2042_v56 = _nw350;
      let _index342 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_2042_v56).length));
      (_2042_v56)[_index342] = (_this).f19;
      let _index343 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_2042_v56).length));
      (_2042_v56)[_index343] = _dafny.Seq.Concat((_this).f19, (_this).f19);
      let _2044_i9;
      _2044_i9 = _dafny.ZERO;
      L13: {
        while ((_1977_v0).f18) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2044_i9)) {
              break L13;
            }
            _2044_i9 = (_2044_i9).plus(_dafny.ONE);
            let _2045_v57;
            let _init51 = function (_2046_i10) {
              return false;
            };
            let _nw351 = Array((new BigNumber(24)).toNumber());
            for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw351.length); _i0_51++) {
              _nw351[_i0_51] = _init51(new BigNumber(_i0_51));
            }
            _2045_v57 = _nw351;
            let _2047_v58;
            _2047_v58 = _dafny.Set.fromElements(_2045_v57, _2045_v57);
            let _2048_v59;
            _2048_v59 = _2045_v57;
            let _2049_v60;
            _2049_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2047_v58).Intersect(_2047_v58)).length),(_2048_v59));
            _2049_v60 = (_2049_v60).Merge(_2049_v60);
            let _2050_v61;
            _2050_v61 = new _dafny.CodePoint('i'.codePointAt(0));
            let _2051_v62;
            let _nw352 = new _module.C5();
            _nw352.__ctor(_dafny.Seq.Concat((_2042_v56)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_2042_v56).length))], (_this).f19), _dafny.Seq.Concat((_2042_v56)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_2042_v56).length))], (_2042_v56)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_2042_v56).length))]), !_dafny.Seq.contains((_this).f19, _2050_v61));
            _2051_v62 = _nw352;
            let _2052_v63;
            _2052_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2050_v61,new BigNumber((_dafny.Seq.of((_this).f18, (_this).f18)).length));
            _2052_v63 = (_2052_v63).update(_module.__default.fm46(globalState), _this.f26);
            let _2053_v65;
            let _nw353 = new _module.C4();
            _nw353.__ctor(_2050_v61, _this.f26, (_this.f26).isEqualTo((_1979_v2)[_module.__default.safeIndex(_this.f26, new BigNumber((_1979_v2).length))]), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wgfvxeb"), (_2051_v62).f28), _dafny.Seq.update((_this).f20, _module.__default.safeIndex(_this.f26, new BigNumber(((_this).f20).length)), function () {
              let _coll51 = new _dafny.Map();
              for (const _compr_51 of (_dafny.Seq.of(_this.f26)).Elements) {
                let _2054_v64 = _compr_51;
                if (_dafny.Seq.contains(_dafny.Seq.of(_this.f26), _2054_v64)) {
                  _coll51.push([(_2054_v64).minus(_this.f26),new BigNumber(((_this).f25).length)]);
                }
              }
              return _coll51;
            }()));
            _2053_v65 = _nw353;
          }
        }
      }
      (globalState).f2 = !((((true) ? (_this.f26) : (_this.f26))).isLessThanOrEqualTo(_this.f26));
      r0 = (_this).f18;
      let _2055_v66;
      _2055_v66 = _dafny.Map.Empty.slice().updateUnsafe(!((_1977_v0).f18),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f26,new BigNumber(876))).length));
      r1 = new BigNumber((_2055_v66).length);
      return [r0, r1];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f18 = false;
      this.f24 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f24, f18) {
      let _this = this;
      (_this).f24 = f24;
      (_this)._f18 = f18;
      return;
    }
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm11(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.UnicodeFromString("ifgrjv"))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.UnicodeFromString("hpppgtxgx")))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("jdc"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("rgxqth"))));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      (globalState).f12 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("oukov"), _dafny.Seq.UnicodeFromString("pdkrn")), p2);
      let _2056_v0;
      let _nw354 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _2056_v0 = _nw354;
      let _2057_v1;
      _2057_v1 = _dafny.Map.Empty.slice().updateUnsafe(true,_2056_v0);
      let _2058_v2;
      _2058_v2 = _dafny.MultiSet.fromElements(_2057_v1);
      (globalState).f3 = (((_2058_v2).contains((_2057_v1).Merge(_2057_v1))) ? ((_2058_v2).get((_2057_v1).Merge(_2057_v1))) : (_this.f24));
      let _2059_v3;
      _2059_v3 = new _dafny.CodePoint('a'.codePointAt(0));
      let _2060_v4;
      _2060_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_2059_v3);
      let _2061_v5;
      _2061_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2060_v4,(_this).f18);
      let _2062_v6;
      _2062_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2059_v3);
      let _2063_v7;
      _2063_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2061_v5,_2062_v6);
      let _2064_v9;
      _2064_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("qv"),false);
      _2063_v7 = (_2063_v7).update(_2061_v5, function () {
        let _coll52 = new _dafny.Map();
        for (const _compr_52 of (_2064_v9).Keys.Elements) {
          let _2065_v8 = _compr_52;
          if ((_2064_v9).contains(_2065_v8)) {
            _coll52.push([_2065_v8,_2059_v3]);
          }
        }
        return _coll52;
      }());
      let _2066_v10;
      let _init52 = function (_2067_i0) {
        return (_this).f18;
      };
      let _nw355 = Array((new BigNumber(26)).toNumber());
      for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw355.length); _i0_52++) {
        _nw355[_i0_52] = _init52(new BigNumber(_i0_52));
      }
      _2066_v10 = _nw355;
      let _index344 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length));
      (_2066_v10)[_index344] = !(((_this).f18) === (!(false)));
      let _index345 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length));
      (_2066_v10)[_index345] = !((((_this).f18) ? ((_this).fm10((_this).f18, (_this).f18, (_this).f18, globalState)) : ((p1).isLessThanOrEqualTo(_this.f24))));
      _2059_v3 = _2059_v3;
      let _2068_v11;
      _2068_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
      let _2069_v12;
      _2069_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(((_2068_v11).contains((_this).f18)) ? ((_2068_v11).get((_this).f18)) : (true)));
      let _2070_v13;
      _2070_v13 = _dafny.Set.fromElements(new BigNumber((_2069_v12).length));
      let _2071_v14;
      _2071_v14 = _dafny.Map.Empty.slice().updateUnsafe((_2066_v10)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length))],p1);
      let _2072_v16;
      _2072_v16 = _dafny.Seq.of(_this.f24);
      let _2073_i1;
      _2073_i1 = _dafny.ZERO;
      L14: {
        while (((_2070_v13).equals(_2070_v13)) === (!((_2071_v14).update((_2066_v10)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length))], new BigNumber((function () {
          let _coll53 = new _dafny.Map();
          for (const _compr_53 of (_2072_v16).Elements) {
            let _2096_v15 = _compr_53;
            if (_dafny.Seq.contains(_2072_v16, _2096_v15)) {
              _coll53.push([(_2096_v15).minus(p1),(_this).f18]);
            }
          }
          return _coll53;
        }()).length))).equals(_2071_v14))) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2073_i1)) {
              break L14;
            }
            _2073_i1 = (_2073_i1).plus(_dafny.ONE);
            let _2074_v17;
            _2074_v17 = _dafny.Seq.of(p0);
            let _2075_v18;
            let _nw356 = Array((new BigNumber(28)).toNumber());
            _nw356[(_dafny.ZERO).toNumber()] = p2;
            _nw356[(_dafny.ONE).toNumber()] = p2;
            _nw356[(new BigNumber(2)).toNumber()] = p0;
            _nw356[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(p2, p0);
            _nw356[(new BigNumber(4)).toNumber()] = p2;
            _nw356[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(470)), ((_2076_v3) => function (_2077_i2) {
              return _2076_v3;
            })(_2059_v3));
            _nw356[(new BigNumber(6)).toNumber()] = p0;
            _nw356[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(p0, p2);
            _nw356[(new BigNumber(8)).toNumber()] = p0;
            _nw356[(new BigNumber(9)).toNumber()] = p2;
            _nw356[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gvm"), _dafny.Seq.UnicodeFromString("qmjimlo"));
            _nw356[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), ((_2078_v3) => function (_2079_i3) {
              return _2078_v3;
            })(_2059_v3));
            _nw356[(new BigNumber(12)).toNumber()] = p0;
            _nw356[(new BigNumber(13)).toNumber()] = p2;
            _nw356[(new BigNumber(14)).toNumber()] = p2;
            _nw356[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(p2, (_2074_v17)[_module.__default.safeIndex(new BigNumber((p2).length), new BigNumber((_2074_v17).length))]);
            _nw356[(new BigNumber(16)).toNumber()] = p0;
            _nw356[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm12(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(853)), ((_2080_v3) => function (_2081_i4) {
              return _2080_v3;
            })(_2059_v3))).length), (_2066_v10)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length))], globalState), p0);
            _nw356[(new BigNumber(18)).toNumber()] = p0;
            _nw356[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(p0, _dafny.Seq.UnicodeFromString("r"));
            _nw356[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("etpqxgf");
            _nw356[(new BigNumber(21)).toNumber()] = p0;
            _nw356[(new BigNumber(22)).toNumber()] = p2;
            _nw356[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("nuxgycs");
            _nw356[(new BigNumber(24)).toNumber()] = p2;
            _nw356[(new BigNumber(25)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-679)), ((_2082_v3) => function (_2083_i5) {
              return _2082_v3;
            })(_2059_v3)), _module.__default.fm12(_this.f24, true, globalState)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-679)), ((_2084_v3) => function (_2085_i5) {
              return _2084_v3;
            })(_2059_v3)), _module.__default.fm12(_this.f24, true, globalState))).length)), _2059_v3);
            _nw356[(new BigNumber(26)).toNumber()] = p2;
            _nw356[(new BigNumber(27)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(59)), ((_2086_v3) => function (_2087_i6) {
              return _2086_v3;
            })(_2059_v3));
            _2075_v18 = _nw356;
            let _2088_v19;
            _2088_v19 = _dafny.Seq.of((_this).f18, !((_2066_v10)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length))]));
            let _index346 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_2075_v18).length));
            (_2075_v18)[_index346] = _dafny.Seq.update(_dafny.Seq.update(p2, _module.__default.safeIndex(p1, new BigNumber((p2).length)), new _dafny.CodePoint('s'.codePointAt(0))), _module.__default.safeIndex(new BigNumber((_2088_v19).length), new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex(p1, new BigNumber((p2).length)), new _dafny.CodePoint('s'.codePointAt(0)))).length)), _2059_v3);
            let _index347 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_2075_v18).length));
            (_2075_v18)[_index347] = _dafny.Seq.update(p0, _module.__default.safeIndex(p1, new BigNumber((p0).length)), _2059_v3);
            (globalState).f3 = (((_this).f18) ? (p1) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(586)), function (_2089_i7) {
              return _dafny.Map.Empty.slice().updateUnsafe(_this.f24,new BigNumber(512));
            })).length)));
            let _2090_v20;
            _2090_v20 = _dafny.MultiSet.fromElements((p1).isLessThan(p1));
            _2090_v20 = _2090_v20;
            let _2091_v21;
            _2091_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(341),_module.__default.fm1(p1, (_2066_v10)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length))], (_2066_v10)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length))], globalState));
            let _2092_v22;
            _2092_v22 = _dafny.Seq.of(_2091_v21, _2091_v21);
            let _2093_v23;
            let _nw357 = new _module.C6();
            _nw357.__ctor(_dafny.Seq.of((_2066_v10)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length))], (_2066_v10)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_2066_v10).length))]), _this.f24, !(_dafny.areEqual((_2075_v18)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_2075_v18).length))], p2)), p2, _dafny.Seq.Concat(_2092_v22, _2092_v22));
            _2093_v23 = _nw357;
            let _2094_v24;
            _2094_v24 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2093_v23);
            let _2095_v25;
            _2095_v25 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),_2093_v23);
            _2093_v23 = (((_2094_v24).contains((_2093_v23).f19)) ? ((_2094_v24).get((_2093_v23).f19)) : ((((_2095_v25).contains(_2059_v3)) ? ((_2095_v25).get(_2059_v3)) : (_2093_v23))));
          }
        }
      }
      return;
    }
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _2097_v0;
      _2097_v0 = _module.D17.create_DC42((_this).f18);
      _2097_v0 = _2097_v0;
      (globalState).f2 = !((_this).f18);
      let _hi11 = p3;
      for (let _2098_i0 = _module.__default.safeDivisionInt(_this.f24, new BigNumber(599)); _2098_i0.isLessThan(_hi11); _2098_i0 = _2098_i0.plus(_dafny.ONE)) {
        let _2099_v1;
        _2099_v1 = _dafny.Seq.of(new BigNumber((_module.__default.fm43((_this).f18, !((_this).f18), globalState)).length), _module.__default.safeModuloInt(p1, p3));
        (globalState).f10 = _2099_v1;
        let _2100_v2;
        _2100_v2 = _dafny.Seq.UnicodeFromString("dqugt");
        let _2101_v3;
        _2101_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2100_v2,(_this).f18);
        let _2102_v4;
        let _nw358 = new _module.C0();
        _nw358.__ctor((((p2).contains(p1)) ? ((p2).get(p1)) : (p1)), (_2099_v1)[_module.__default.safeIndex(new BigNumber((_2101_v3).length), new BigNumber((_2099_v1).length))]);
        _2102_v4 = _nw358;
        let _2103_v5;
        _2103_v5 = _module.D11.create_DC29(_module.D11.create_DC28((_this).f18));
        _2103_v5 = _2103_v5;
        let _2104_v6;
        let _init53 = function (_2105_i1) {
          return (_dafny.Map.Empty.slice().updateUnsafe((_this).f18,!(!((_this).f18)))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18));
        };
        let _nw359 = Array((new BigNumber(26)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw359.length); _i0_53++) {
          _nw359[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _2104_v6 = _nw359;
        let _2106_v7;
        _2106_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,!((_this).f18));
        let _index348 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2104_v6).length));
        (_2104_v6)[_index348] = _2106_v7;
        let _index349 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2104_v6).length));
        (_2104_v6)[_index349] = (((_this).f18) ? ((_2106_v7).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,true))) : (_2106_v7));
      }
      let _2107_v8;
      let _nw360 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _2107_v8 = _nw360;
      let _index350 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length));
      (_2107_v8)[_index350] = (p3).multipliedBy(_this.f24);
      let _2108_v9;
      let _nw361 = Array((new BigNumber(6)).toNumber());
      _nw361[(_dafny.ZERO).toNumber()] = ((_dafny.ZERO).minus(_this.f24)).isLessThanOrEqualTo(p1);
      _nw361[(_dafny.ONE).toNumber()] = (_this).f18;
      _nw361[(new BigNumber(2)).toNumber()] = (_this).f18;
      _nw361[(new BigNumber(3)).toNumber()] = true;
      _nw361[(new BigNumber(4)).toNumber()] = (_this).f18;
      _nw361[(new BigNumber(5)).toNumber()] = (_this).f18;
      _2108_v9 = _nw361;
      let _index351 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length));
      let _rhs364 = p3;
      let _rhs365 = _2108_v9;
      let _rhs366 = (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("ic")).length), _this.f24)).plus(p3);
      let _lhs313 = _2107_v8;
      let _lhs314 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length));
      let _lhs315 = globalState;
      _lhs313[_lhs314] = _rhs364;
      _2108_v9 = _rhs365;
      _lhs315.f3 = _rhs366;
      let _2109_v10;
      _2109_v10 = new _dafny.CodePoint('e'.codePointAt(0));
      let _2110_v11;
      _2110_v11 = _dafny.Seq.UnicodeFromString("uermrgdr");
      let _2111_v12;
      _2111_v12 = _dafny.Seq.of((_this).f18);
      let _2112_v13;
      _2112_v13 = _dafny.Seq.of(new BigNumber((_2111_v12).length), new BigNumber((_dafny.Set.fromElements((_this).f18, true)).length));
      let _2113_v14;
      let _nw362 = new _module.C2();
      _nw362.__ctor(p3, _2112_v13, (_this).f18);
      _2113_v14 = _nw362;
      let _2114_v15;
      _2114_v15 = _module.D12.create_DC31();
      let _2115_v16;
      _2115_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2113_v14,_2114_v15);
      let _2116_v17;
      _2116_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2115_v16).length),_this.f24);
      let _2117_v18;
      _2117_v18 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p3,p1), _2116_v17, _2116_v17, _2116_v17);
      let _2118_v19;
      let _nw363 = new _module.C4();
      _nw363.__ctor(_2109_v10, ((_2107_v8)[_module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length))]).plus(p3), (_this).f18, _2110_v11, _2117_v18);
      _2118_v19 = _nw363;
      let _2119_v20;
      _2119_v20 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18);
      if ((_2111_v12)[_module.__default.safeIndex((_this.f24).minus(new BigNumber((_2119_v20).cardinality())), new BigNumber((_2111_v12).length))]) {
        let _2120_v21;
        let _nw364 = new _module.C0();
        _nw364.__ctor(new BigNumber(74), (_2118_v19).fm18(globalState));
        _2120_v21 = _nw364;
        _2109_v10 = _module.__default.fm46(globalState);
        let _2121_v22;
        _2121_v22 = _module.D17.create_DC44(_2108_v9, _this.f24);
        let _2122_v23;
        _2122_v23 = _dafny.Set.fromElements((_2107_v8)[_module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length))]);
        let _pat_let_tv27 = _2108_v9;
        let _pat_let_tv28 = p1;
        let _pat_let_tv29 = _2122_v23;
        let _2123_v24;
        let _nw365 = Array((new BigNumber(20)).toNumber());
        _nw365[(_dafny.ZERO).toNumber()] = function (_pat_let51_0) {
          return function (_2124_dt__update__tmp_h0) {
            return function (_pat_let52_0) {
              return function (_2125_dt__update_hcf77_h0) {
                return _module.D17.create_DC44(_2125_dt__update_hcf77_h0, (_2124_dt__update__tmp_h0).dtor_cf78);
              }(_pat_let52_0);
            }(_pat_let_tv27);
          }(_pat_let51_0);
        }(_2121_v22);
        _nw365[(_dafny.ONE).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(2)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(3)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(4)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(5)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(6)).toNumber()] = (((_this).f18) ? (function (_pat_let53_0) {
          return function (_2126_dt__update__tmp_h1) {
            return function (_pat_let54_0) {
              return function (_2127_dt__update_hcf78_h0) {
                return _module.D17.create_DC44((_2126_dt__update__tmp_h1).dtor_cf77, _2127_dt__update_hcf78_h0);
              }(_pat_let54_0);
            }(_pat_let_tv28);
          }(_pat_let53_0);
        }(_2121_v22)) : (_2121_v22));
        _nw365[(new BigNumber(7)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(8)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(9)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(10)).toNumber()] = function (_pat_let55_0) {
          return function (_2128_dt__update__tmp_h2) {
            return function (_pat_let56_0) {
              return function (_2129_dt__update_hcf78_h1) {
                return _module.D17.create_DC44((_2128_dt__update__tmp_h2).dtor_cf77, _2129_dt__update_hcf78_h1);
              }(_pat_let56_0);
            }(new BigNumber((_pat_let_tv29).length));
          }(_pat_let55_0);
        }(_module.D17.create_DC44(_2108_v9, _module.__default.fm1(_2113_v14.f33, (_this).f18, true, globalState)));
        _nw365[(new BigNumber(11)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(12)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(13)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(14)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(15)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(16)).toNumber()] = _module.D17.create_DC44(_2108_v9, _2113_v14.f33);
        _nw365[(new BigNumber(17)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(18)).toNumber()] = _2121_v22;
        _nw365[(new BigNumber(19)).toNumber()] = _2121_v22;
        _2123_v24 = _nw365;
        let _index352 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_2123_v24).length));
        (_2123_v24)[_index352] = _2121_v22;
        let _2130_v25;
        let _nw366 = new _module.C6();
        _nw366.__ctor(_dafny.Seq.Concat(_2111_v12, _dafny.Seq.of((_this).f18, false, (_this).f18)), p1, (_this).f18, _2110_v11, _dafny.Seq.update(_2117_v18, _module.__default.safeIndex(_this.f24, new BigNumber((_2117_v18).length)), _2116_v17));
        _2130_v25 = _nw366;
        let _2131_v26;
        _2131_v26 = _module.D13.create_DC33((_this).f18, (_2118_v19).f30, !((_this).f18), _2113_v14.f33);
        let _2132_v27;
        _2132_v27 = _module.D3.create_DC11((_this).f18, (_this).f18);
        let _index353 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_2123_v24).length));
        let _rhs367 = ((_this).f18) === ((_2118_v19).fm2((_this).f18, (_this).f18, (_this).f18, globalState));
        let _rhs368 = !((_2131_v26).dtor_cf57);
        let _rhs369 = _dafny.Seq.update(_2110_v11, _module.__default.safeIndex(new BigNumber((_2116_v17).length), new BigNumber((_2110_v11).length)), _2109_v10);
        let _rhs370 = _module.D17.create_DC44(_2108_v9, ((!((_2132_v27).dtor_cf24)) ? ((_2120_v21).f31) : (_this.f24)));
        let _rhs371 = _2130_v25;
        let _lhs316 = globalState;
        let _lhs317 = globalState;
        let _lhs318 = globalState;
        let _lhs319 = _2123_v24;
        let _lhs320 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_2123_v24).length));
        _lhs316.f2 = _rhs367;
        _lhs317.f2 = _rhs368;
        _lhs318.f12 = _rhs369;
        _lhs319[_lhs320] = _rhs370;
        _2130_v25 = _rhs371;
        let _2133_v28;
        _2133_v28 = _dafny.Set.fromElements((_2118_v19).f29, (_2118_v19).f29);
        let _2134_v29;
        _2134_v29 = _dafny.Map.Empty.slice().updateUnsafe((_2118_v19).fm3(new BigNumber(72), globalState),new BigNumber((_2133_v28).length));
        let _2135_v30;
        _2135_v30 = _dafny.Map.Empty.slice().updateUnsafe((((_2134_v29).contains(false)) ? ((_2134_v29).get(false)) : ((_2120_v21).f32)),(_this).f18);
        _2135_v30 = (_2135_v30).update(new BigNumber((_2110_v11).length), !((_this).f18));
        _2135_v30 = (_2135_v30).Merge(function () {
          let _coll54 = new _dafny.Map();
          for (const _compr_54 of (_2112_v13).Elements) {
            let _2136_v31 = _compr_54;
            if (_dafny.Seq.contains(_2112_v13, _2136_v31)) {
              _coll54.push([(_2136_v31).minus((_2120_v21).f31),(_this).f18]);
            }
          }
          return _coll54;
        }());
      } else {
        (globalState).f2 = (_this).f18;
        (globalState).f2 = !((_this).f18);
        let _2137_v32;
        let _nw367 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2137_v32 = _nw367;
        let _index354 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_2137_v32).length));
        (_2137_v32)[_index354] = (_module.D13.create_DC34(_2108_v9, _2110_v11, (_this).f18, (_this).f18)).dtor_cf60;
        let _2138_v33;
        _2138_v33 = _dafny.Seq.of(_2110_v11);
        let _2139_v34;
        _2139_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_2138_v33)[_module.__default.safeIndex(_2113_v14.f33, new BigNumber((_2138_v33).length))]);
        let _index355 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_2137_v32).length));
        (_2137_v32)[_index355] = (((_2139_v34).contains(((_2107_v8)[_module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length))]).isLessThan(new BigNumber(((_2113_v14).f34).length)))) ? ((_2139_v34).get(((_2107_v8)[_module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length))]).isLessThan(new BigNumber(((_2113_v14).f34).length)))) : (_dafny.Seq.UnicodeFromString("lnxxapdgd")));
        let _2140_v35;
        let _init54 = ((_2141_p3, _2142_v17) => function (_2143_i2) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_2141_p3,new BigNumber(86))).Merge(_2142_v17);
        })(p3, _2116_v17);
        let _nw368 = Array((new BigNumber(12)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw368.length); _i0_54++) {
          _nw368[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _2140_v35 = _nw368;
        let _index356 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_2140_v35).length));
        (_2140_v35)[_index356] = _2116_v17;
        let _2144_v36;
        _2144_v36 = _dafny.Seq.of(_2119_v20);
        let _index357 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_2140_v35).length));
        (_2140_v35)[_index357] = _module.__default.fm30(_2144_v36, function () {
          let _coll55 = new _dafny.Map();
          for (const _compr_55 of (_2138_v33).Elements) {
            let _2145_v37 = _compr_55;
            if (_dafny.Seq.contains(_2138_v33, _2145_v37)) {
              _coll55.push([_2145_v37,(_this).f18]);
            }
          }
          return _coll55;
        }(), globalState);
        let _2146_v38;
        _2146_v38 = _dafny.Set.fromElements((_this).f18);
        let _2147_v39;
        _2147_v39 = _module.D10.create_DC24(_2146_v38);
        let _2148_v40;
        _2148_v40 = _module.D10.create_DC26(_module.D10.create_DC26(_2147_v39));
        let _source18 = _2148_v40;
        if (_source18.is_DC25) {
          let _2149___mcc_h0 = (_source18).cf48;
          let _2150_cf48 = _2149___mcc_h0;
          let _2151_v41;
          _2151_v41 = _dafny.Seq.of(_module.__default.fm44(_2113_v14.f33, _2150_cf48, _module.__default.fm46(globalState), _2113_v14.f33, globalState));
          let _2152_v42;
          _2152_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_2151_v41)[_module.__default.safeIndex((_2118_v19).f30, new BigNumber((_2151_v41).length))]);
          let _2153_v43;
          _2153_v43 = _dafny.Map.Empty.slice().updateUnsafe(_2108_v9,!(!(true)));
          _2111_v12 = _dafny.Seq.update((((_2152_v42).contains(_2150_cf48)) ? ((_2152_v42).get(_2150_cf48)) : (_dafny.Seq.Concat(_2111_v12, _2111_v12))), _module.__default.safeIndex(new BigNumber((_2153_v43).length), new BigNumber(((((_2152_v42).contains(_2150_cf48)) ? ((_2152_v42).get(_2150_cf48)) : (_dafny.Seq.Concat(_2111_v12, _2111_v12)))).length)), _2150_cf48);
          let _2154_v44;
          _2154_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2113_v14.f33,_2150_cf48);
          _2154_v44 = (_2154_v44).update((_2118_v19).f30, ((_2111_v12)[_module.__default.safeIndex(new BigNumber((_2112_v13).length), new BigNumber((_2111_v12).length))]) || ((_this).f18));
          (globalState).f2 = (((_this).f18) ? ((_this).f18) : ((_this).f18));
          (globalState).f3 = (_2107_v8)[_module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length))];
        } else if (_source18.is_DC24) {
          let _2155___mcc_h1 = (_source18).cf47;
          let _2156_cf47 = _2155___mcc_h1;
          let _2157_v45;
          let _nw369 = new _module.C4();
          _nw369.__ctor((_2118_v19).f29, (_2118_v19).f30, (_this).f18, (_2137_v32)[_module.__default.safeIndex(new BigNumber(283), new BigNumber((_2137_v32).length))], _2117_v18);
          _2157_v45 = _nw369;
          let _index358 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length));
          let _rhs372 = (_this).f18;
          let _rhs373 = new BigNumber((_2110_v11).length);
          let _rhs374 = (false) === ((_this).f18);
          let _rhs375 = ((_2118_v19).f30).plus(new BigNumber(((_2137_v32)[_module.__default.safeIndex(new BigNumber(283), new BigNumber((_2137_v32).length))]).length));
          let _lhs321 = globalState;
          let _lhs322 = _2107_v8;
          let _lhs323 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length));
          let _lhs324 = globalState;
          let _lhs325 = _2113_v14;
          _lhs321.f2 = _rhs372;
          _lhs322[_lhs323] = _rhs373;
          _lhs324.f2 = _rhs374;
          _lhs325.f33 = _rhs375;
          (globalState).f0 = (new BigNumber(265)).plus(_2113_v14.f33);
          let _index359 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_2140_v35).length));
          (_2140_v35)[_index359] = ((_2140_v35)[_module.__default.safeIndex(new BigNumber(951), new BigNumber((_2140_v35).length))]).update(((_2157_v45).f30).plus(p1), p1);
        } else {
          let _2158___mcc_h2 = (_source18).cf49;
          let _2159_cf49 = _2158___mcc_h2;
          let _index360 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_2107_v8).length));
          (_2107_v8)[_index360] = _module.__default.safeDivisionInt(_module.__default.fm1((_2118_v19).f30, (_this).f18, (_this).f18, globalState), (_2113_v14.f33).minus(new BigNumber(42)));
          let _2160_v46;
          _2160_v46 = _dafny.Seq.of((_2113_v14).f34, (_2113_v14).f34);
          (globalState).f2 = _dafny.Seq.contains(_2160_v46, _dafny.Seq.Concat((_2113_v14).f34, _2112_v13));
          (globalState).f0 = (((_2113_v14).f34)[_module.__default.safeIndex(_2113_v14.f33, new BigNumber(((_2113_v14).f34).length))]).multipliedBy(_2113_v14.f33);
          (globalState).f2 = (_this).f18;
        }
      }
      return;
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f18 = false;
      this.f23 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f23, f18) {
      let _this = this;
      (_this).f23 = f23;
      (_this)._f18 = f18;
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _2161_v0;
      _2161_v0 = _dafny.Set.fromElements(_module.__default.fm1(p1, (_this).f18, (_this).f18, globalState));
      let _hi12 = (new BigNumber(486)).plus(new BigNumber((_2161_v0).length));
      for (let _2162_i0 = p1; _2162_i0.isLessThan(_hi12); _2162_i0 = _2162_i0.plus(_dafny.ONE)) {
        let _2163_v1;
        _2163_v1 = new _dafny.CodePoint('w'.codePointAt(0));
        let _2164_v2;
        _2164_v2 = _dafny.MultiSet.fromElements(_2163_v1, _2163_v1, _2163_v1);
        _2164_v2 = (_2164_v2).Intersect((_2164_v2).Intersect((_2164_v2).update(_2163_v1, _module.__default.abs(_2162_i0))));
        (globalState).f0 = p1;
        let _2165_v3;
        _2165_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
        let _2166_v4;
        _2166_v4 = _dafny.Seq.of(_this.f23, new BigNumber((_2165_v3).length));
        let _2167_v5;
        _2167_v5 = _dafny.MultiSet.fromElements(p1);
        if (!(!(_module.__default.fm6(_dafny.Seq.Concat(_2166_v4, _2166_v4), false, (_2167_v5).Difference(_2167_v5), p1, globalState)))) {
          let _2168_v6;
          let _nw370 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _2168_v6 = _nw370;
          let _2169_v7;
          _2169_v7 = _dafny.Seq.of(_2168_v6);
          let _2170_v8;
          _2170_v8 = _dafny.Seq.of(_2169_v7, _2169_v7, _dafny.Seq.of(_2168_v6), _2169_v7, _2169_v7);
          (globalState).f2 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat((_2170_v8)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_2170_v8).length))], _2169_v7), _2169_v7);
          let _2171_v9;
          let _nw371 = Array((new BigNumber(17)).toNumber()).fill(false);
          _2171_v9 = _nw371;
          let _2172_v10;
          _2172_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2171_v9,p1);
          (globalState).f2 = ((_2162_i0).plus(_2162_i0)).isLessThanOrEqualTo((((_2172_v10).contains(_2171_v9)) ? ((_2172_v10).get(_2171_v9)) : (new BigNumber(9))));
          let _2173_v11;
          _2173_v11 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_2162_i0, (_this).f18, (_this).f18, globalState),(_this).f18);
          let _2174_v12;
          _2174_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_2173_v11).update(new BigNumber(670), (_this).f18));
          let _2175_v13;
          _2175_v13 = _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_module.__default.fm1(_module.__default.fm1(p1, (_this).f18, (_this).f18, globalState), _module.__default.fm6(_2166_v4, (_this).f18, _2167_v5, _2162_i0, globalState), (_this).f18, globalState)));
          let _2176_v14;
          _2176_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,new BigNumber((_dafny.MultiSet.fromElements((_this).f18, (_this).f18, false)).cardinality()));
          let _2177_v15;
          _2177_v15 = _dafny.Set.fromElements(!((_this).f18));
          let _2178_v16;
          _2178_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2165_v3);
          let _pat_let_tv30 = _2176_v14;
          let _pat_let_tv31 = _2176_v14;
          let _pat_let_tv32 = _2176_v14;
          let _2179_v17;
          let _nw372 = Array((new BigNumber(25)).toNumber());
          _nw372[(_dafny.ZERO).toNumber()] = _module.D2.create_DC6(_module.__default.fm7(new BigNumber((_2174_v12).length), (_this).f18, globalState));
          _nw372[(_dafny.ONE).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(2)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(3)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(4)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(5)).toNumber()] = _module.D2.create_DC6(_2176_v14);
          _nw372[(new BigNumber(6)).toNumber()] = function (_pat_let57_0) {
            return function (_2180_dt__update__tmp_h0) {
              return function (_pat_let58_0) {
                return function (_2181_dt__update_hcf14_h0) {
                  return _module.D2.create_DC6(_2181_dt__update_hcf14_h0);
                }(_pat_let58_0);
              }(_pat_let_tv30);
            }(_pat_let57_0);
          }(_2175_v13);
          _nw372[(new BigNumber(7)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(8)).toNumber()] = _module.D2.create_DC6(_2176_v14);
          _nw372[(new BigNumber(9)).toNumber()] = _module.__default.fm8(_2177_v15, (_this).f18, (_this).f18, _2163_v1, globalState);
          _nw372[(new BigNumber(10)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(11)).toNumber()] = function (_pat_let59_0) {
            return function (_2184_dt__update__tmp_h2) {
              return function (_pat_let62_0) {
                return function (_2185_dt__update_hcf14_h2) {
                  return _module.D2.create_DC6(_2185_dt__update_hcf14_h2);
                }(_pat_let62_0);
              }(_pat_let_tv32);
            }(_pat_let59_0);
          }(function (_pat_let60_0) {
            return function (_2182_dt__update__tmp_h1) {
              return function (_pat_let61_0) {
                return function (_2183_dt__update_hcf14_h1) {
                  return _module.D2.create_DC6(_2183_dt__update_hcf14_h1);
                }(_pat_let61_0);
              }(_pat_let_tv31);
            }(_pat_let60_0);
          }(_2175_v13));
          _nw372[(new BigNumber(12)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(13)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(14)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(15)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(16)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(17)).toNumber()] = ((_module.__default.fm6(_2166_v4, !((_this).f18), _dafny.MultiSet.fromElements(new BigNumber((_2178_v16).length), new BigNumber(-262)), new BigNumber((_2177_v15).length), globalState)) ? (_module.__default.fm8(_2177_v15, (_this).f18, !((_this).f18), _2163_v1, globalState)) : (_2175_v13));
          _nw372[(new BigNumber(18)).toNumber()] = _module.D2.create_DC6(_module.__default.fm7(new BigNumber(56), false, globalState));
          _nw372[(new BigNumber(19)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(20)).toNumber()] = _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,new BigNumber((_2166_v4).length)));
          _nw372[(new BigNumber(21)).toNumber()] = _module.__default.fm8(_2177_v15, (_this).f18, (_this).f18, _2163_v1, globalState);
          _nw372[(new BigNumber(22)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(23)).toNumber()] = _2175_v13;
          _nw372[(new BigNumber(24)).toNumber()] = _module.D2.create_DC6(_2176_v14);
          _2179_v17 = _nw372;
          let _index361 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_2179_v17).length));
          (_2179_v17)[_index361] = _2175_v13;
          let _index362 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_2179_v17).length));
          (_2179_v17)[_index362] = _2175_v13;
          (globalState).f2 = ((_module.__default.fm9((_this).f18, globalState)).Union(_dafny.Set.fromElements((_this).f18, (_this).f18, false, (_this).f18))).IsDisjointFrom(_module.__default.fm9((_this).f18, globalState));
          let _2186_v18;
          let _nw373 = Array((new BigNumber(11)).toNumber());
          _nw373[(_dafny.ZERO).toNumber()] = p0;
          _nw373[(_dafny.ONE).toNumber()] = p0;
          _nw373[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(p0, p2);
          _nw373[(new BigNumber(3)).toNumber()] = p0;
          _nw373[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("e");
          _nw373[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(747)), ((_2187_v1) => function (_2188_i1) {
            return _2187_v1;
          })(_2163_v1));
          _nw373[(new BigNumber(6)).toNumber()] = p0;
          _nw373[(new BigNumber(7)).toNumber()] = p0;
          _nw373[(new BigNumber(8)).toNumber()] = (_module.D0.create_DC1((_this).f18, new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex(_this.f23, new BigNumber((p2).length)), _2163_v1)).length), p2, (_this).f18, _this.f23)).dtor_cf3;
          _nw373[(new BigNumber(9)).toNumber()] = p0;
          _nw373[(new BigNumber(10)).toNumber()] = p2;
          _2186_v18 = _nw373;
          let _index363 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_2186_v18).length));
          (_2186_v18)[_index363] = _dafny.Seq.Concat(p2, p2);
          let _index364 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_2186_v18).length));
          (_2186_v18)[_index364] = _dafny.Seq.Concat(p2, _dafny.Seq.UnicodeFromString("msaxakyd"));
        } else {
          let _2189_v19;
          _2189_v19 = _dafny.Set.fromElements(true, ((_this).f18) && ((_this).f18), (_this).f18, !_dafny.Seq.contains(p0, _2163_v1));
          _2189_v19 = (((!((_this).f18)) ? (_dafny.Set.fromElements(!((_this).f18), _module.__default.fm6(_dafny.Seq.of(_this.f23, _this.f23), (_this).f18, _2167_v5, new BigNumber(800), globalState), (_this).f18)) : (_dafny.Set.fromElements((_this).f18)))).Union(_2189_v19);
          (globalState).f2 = (_this).f18;
          (globalState).f3 = (_dafny.ZERO).minus((new BigNumber(417)).minus((_this.f23).multipliedBy(p1)));
          let _2190_v20;
          _2190_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_this.f23);
          let _2191_v21;
          let _nw374 = new _module.C4();
          _nw374.__ctor(_2163_v1, new BigNumber(611), (_this).f18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(842)), ((_2192_v1) => function (_2193_i2) {
            return _2192_v1;
          })(_2163_v1)), _dafny.Seq.of(_2190_v20, _2190_v20));
          _2191_v21 = _nw374;
          let _2194_v22;
          _2194_v22 = _module.D1.create_DC2(_2191_v21);
          let _2195_v23;
          _2195_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2194_v22,new BigNumber((_dafny.Seq.UnicodeFromString("fdy")).length));
          let _2196_v24;
          _2196_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p2);
          let _2197_v25;
          _2197_v25 = _dafny.Map.Empty.slice().updateUnsafe((((_2195_v23).contains(_module.D1.create_DC2(_2191_v21))) ? ((_2195_v23).get(_module.D1.create_DC2(_2191_v21))) : (_this.f23)),new BigNumber((_2196_v24).length));
          let _2198_v26;
          _2198_v26 = _dafny.Map.Empty.slice().updateUnsafe((_2191_v21).f18,new BigNumber(442));
          let _2199_v27;
          _2199_v27 = _dafny.Seq.of((_this).f18, (_2191_v21).f18, (_2191_v21).f18, (_2191_v21).f18, (_2191_v21).f18);
          _2190_v20 = ((_2190_v20).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2198_v26).length),new BigNumber((_2199_v27).length)))).Merge((_2190_v20).Merge(_2190_v20));
          _2197_v25 = (_2197_v25).update(new BigNumber(929), _2162_i0);
        }
        let _2200_v28;
        _2200_v28 = _module.D19.create_DC50((_this).f18);
        let _2201_v29;
        _2201_v29 = _module.D2.create_DC8((_2200_v28).dtor_cf84, (_this).f18, true);
        let _2202_v30;
        _2202_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2201_v29,(_this).f18);
        let _2203_v31;
        let _init55 = function (_2204_i3) {
          return (_2204_i3).multipliedBy(_this.f23);
        };
        let _nw375 = Array((new BigNumber(8)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw375.length); _i0_55++) {
          _nw375[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _2203_v31 = _nw375;
        let _2205_v32;
        _2205_v32 = _dafny.Seq.of(_2203_v31, _2203_v31);
        let _2206_v33;
        _2206_v33 = _dafny.MultiSet.fromElements(_module.__default.fm6(_2166_v4, (_this).f18, _2167_v5, _this.f23, globalState), (_this).f18);
        let _rhs376 = _2202_v30;
        let _rhs377 = _dafny.Seq.Concat(_dafny.Seq.of(_2203_v31, _2203_v31), _2205_v32);
        let _rhs378 = _module.__default.fm1(_2162_i0, (_2206_v33).IsDisjointFrom(_2206_v33), (_this).f18, globalState);
        let _lhs326 = globalState;
        _2202_v30 = _rhs376;
        _2205_v32 = _rhs377;
        _lhs326.f3 = _rhs378;
      }
      let _2207_v34;
      let _nw376 = new _module.C7();
      _nw376.__ctor(_this.f23, (_this).f18);
      _2207_v34 = _nw376;
      if ((_this).f18) {
        let _2208_v35;
        let _nw377 = new _module.C1();
        _nw377.__ctor((_this).f18);
        _2208_v35 = _nw377;
        let _2209_v36;
        _2209_v36 = _dafny.Seq.of(_2207_v34.f24, new BigNumber((p2).length));
        (globalState).f2 = !(_this.f23).isEqualTo(_module.__default.fm1(new BigNumber((_2209_v36).length), (_this).f18, (_this).f18, globalState));
        if ((_this).f18) {
          let _2210_v37;
          _2210_v37 = _dafny.Seq.of(p2);
          let _2211_v38;
          _2211_v38 = _dafny.MultiSet.fromElements((_this).f18);
          let _2212_v39;
          _2212_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_2210_v37)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_2211_v38).cardinality())), new BigNumber((_2210_v37).length))]);
          _2212_v39 = (_2212_v39).update((p1).minus(_2207_v34.f24), _dafny.Seq.UnicodeFromString("bko"));
          let _2213_v40;
          _2213_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2209_v36,(_this).f18);
          _2213_v40 = (_2213_v40).update(_dafny.Seq.Concat(_2209_v36, _2209_v36), (_this).f18);
          (globalState).f2 = ((_this).f18) || (!(!((_2207_v34).fm10(true, (_this).f18, (_this).f18, globalState))) || (false));
          (globalState).f10 = _2209_v36;
          (globalState).f2 = (_this).f18;
        } else {
          let _2214_v41;
          _2214_v41 = _dafny.MultiSet.fromElements(p1);
          _2214_v41 = _2214_v41;
          let _2215_v42;
          _2215_v42 = _dafny.MultiSet.fromElements((_this).f18);
          let _2216_v43;
          _2216_v43 = _dafny.Seq.of((_this).f18);
          let _2217_v44;
          _2217_v44 = _module.D12.create_DC30(_2216_v43);
          let _2218_v45;
          _2218_v45 = _2215_v42;
          let _2219_v46;
          _2219_v46 = _dafny.Map.Empty.slice().updateUnsafe(_2214_v41,_2215_v42);
          let _2220_v47;
          let _nw378 = Array((new BigNumber(20)).toNumber());
          _nw378[(_dafny.ZERO).toNumber()] = _2215_v42;
          _nw378[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements((_this).f18, (_this).f18)).update(!((_this).f18), _module.__default.abs(_this.f23));
          _nw378[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.FromArray((_2217_v44).dtor_cf53)).Intersect(_2215_v42);
          _nw378[(new BigNumber(3)).toNumber()] = (_2218_v45);
          _nw378[(new BigNumber(4)).toNumber()] = _2215_v42;
          _nw378[(new BigNumber(5)).toNumber()] = (_2215_v42).update(false, _module.__default.abs(_2207_v34.f24));
          _nw378[(new BigNumber(6)).toNumber()] = _2215_v42;
          _nw378[(new BigNumber(7)).toNumber()] = (_2215_v42).Intersect(_2215_v42);
          _nw378[(new BigNumber(8)).toNumber()] = _2215_v42;
          _nw378[(new BigNumber(9)).toNumber()] = (((_2219_v46).contains(_dafny.MultiSet.FromArray(_2209_v36))) ? ((_2219_v46).get(_dafny.MultiSet.FromArray(_2209_v36))) : (_2215_v42));
          _nw378[(new BigNumber(10)).toNumber()] = (_2215_v42).Union(_dafny.MultiSet.fromElements((_this).f18, false, !((_this).f18)));
          _nw378[(new BigNumber(11)).toNumber()] = _module.__default.fm32(globalState);
          _nw378[(new BigNumber(12)).toNumber()] = _2215_v42;
          _nw378[(new BigNumber(13)).toNumber()] = _2215_v42;
          _nw378[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements((_this).f18);
          _nw378[(new BigNumber(15)).toNumber()] = _2215_v42;
          _nw378[(new BigNumber(16)).toNumber()] = _2215_v42;
          _nw378[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.FromArray(_2216_v43);
          _nw378[(new BigNumber(18)).toNumber()] = (_2215_v42).update((_this).f18, _module.__default.abs(new BigNumber(((_module.D16.create_DC38(_2214_v41)).dtor_cf66).cardinality())));
          _nw378[(new BigNumber(19)).toNumber()] = (_dafny.MultiSet.FromArray(_2216_v43)).Union(_2215_v42);
          _2220_v47 = _nw378;
          let _index365 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_2220_v47).length));
          (_2220_v47)[_index365] = _dafny.MultiSet.fromElements((_this).f18, false, (_this).f18);
          let _index366 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_2220_v47).length));
          (_2220_v47)[_index366] = _dafny.MultiSet.fromElements((((_this).f18) ? (!((_this).f18)) : ((_this).f18)));
          let _2221_v48;
          _2221_v48 = new _dafny.CodePoint('s'.codePointAt(0));
          let _2222_v49;
          _2222_v49 = _module.D3.create_DC13(_2207_v34.f24, p2, _this.f23, _module.__default.fm1(_this.f23, (_this).f18, (_this).f18, globalState), _2221_v48);
          let _2223_v50;
          _2223_v50 = _dafny.Map.Empty.slice().updateUnsafe((_2222_v49).dtor_cf34,!((_2220_v47)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_2220_v47).length))]).contains((_this).f18));
          _2223_v50 = _2223_v50;
          let _2224_v51;
          let _nw379 = Array((new BigNumber(3)).toNumber());
          _nw379[(_dafny.ZERO).toNumber()] = (_this).f18;
          _nw379[(_dafny.ONE).toNumber()] = true;
          _nw379[(new BigNumber(2)).toNumber()] = _module.__default.fm23(_2161_v0, p1, globalState);
          _2224_v51 = _nw379;
          let _index367 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_2224_v51).length));
          (_2224_v51)[_index367] = false;
          let _index368 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_2224_v51).length));
          (_2224_v51)[_index368] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), ((_2225_v48) => function (_2226_i4) {
            return _2225_v48;
          })(_2221_v48)), p2), _dafny.Seq.UnicodeFromString("afcymv"));
          (globalState).f2 = !((_this).f18);
        }
        let _2227_v52;
        _2227_v52 = _module.D1.create_DC4(true);
        let _2228_v53;
        _2228_v53 = _module.D1.create_DC5(_2227_v52);
        let _source19 = _2228_v53;
        if (_source19.is_DC3) {
          let _2229___mcc_h0 = (_source19).cf7;
          let _2230___mcc_h1 = (_source19).cf8;
          let _2231___mcc_h2 = (_source19).cf9;
          let _2232___mcc_h3 = (_source19).cf10;
          let _2233___mcc_h4 = (_source19).cf11;
          let _2234_cf11 = _2233___mcc_h4;
          let _2235_cf10 = _2232___mcc_h3;
          let _2236_cf9 = _2231___mcc_h2;
          let _2237_cf8 = _2230___mcc_h1;
          let _2238_cf7 = _2229___mcc_h0;
          let _2239_v54;
          _2239_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f18);
          _2239_v54 = (_2239_v54).update((_2207_v34.f24).plus(_2207_v34.f24), (_this).f18);
          (globalState).f2 = _dafny.Seq.IsProperPrefixOf(p0, _dafny.Seq.Concat(p2, _dafny.Seq.UnicodeFromString("sunv")));
          (globalState).f3 = _this.f23;
          let _2240_v55;
          let _init56 = function (_2241_i5) {
            return (_this).f18;
          };
          let _nw380 = Array((new BigNumber(29)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw380.length); _i0_56++) {
            _nw380[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2240_v55 = _nw380;
          let _index369 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2240_v55).length));
          (_2240_v55)[_index369] = (_this).f18;
          let _index370 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2240_v55).length));
          (_2240_v55)[_index370] = !_dafny.areEqual(p0, p2);
        } else if (_source19.is_DC4) {
          let _2242___mcc_h5 = (_source19).cf12;
          let _2243_cf12 = _2242___mcc_h5;
          let _2244_v56;
          _2244_v56 = _dafny.MultiSet.fromElements(p0);
          (_2207_v34).m5((_2244_v56).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sbqo"), p0)), _module.__default.safeModuloInt(_2207_v34.f24, p1), _module.__default.fm25(new BigNumber(-768), globalState), new BigNumber((_dafny.Seq.Concat(p0, p2)).length), globalState);
          (globalState).f2 = _dafny.Seq.IsProperPrefixOf(_2209_v36, _2209_v36);
          let _2245_v57;
          _2245_v57 = _dafny.MultiSet.fromElements(_2243_cf12);
          let _2246_v58;
          _2246_v58 = _dafny.Map.Empty.slice().updateUnsafe(((_2243_cf12) ? (_2207_v34.f24) : (_this.f23)),new BigNumber((_2161_v0).length));
          let _2247_v59;
          _2247_v59 = _module.D0.create_DC1((_this).f18, new BigNumber(-99), p2, _2243_cf12, _module.__default.fm1(_2207_v34.f24, false, _2243_cf12, globalState));
          let _rhs379 = _dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-1)), function (_2248_i6) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }));
          let _rhs380 = _2245_v57;
          let _rhs381 = (((_2246_v58).contains(new BigNumber((_2209_v36).length))) ? ((_2246_v58).get(new BigNumber((_2209_v36).length))) : ((_dafny.ZERO).minus(_2207_v34.f24)));
          let _rhs382 = (_2208_v35).fm22(_dafny.Seq.Create(_module.__default.abs(new BigNumber(721)), function (_2249_i7) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          }), _2247_v59, globalState);
          let _rhs383 = _2243_cf12;
          let _lhs327 = globalState;
          let _lhs328 = _2207_v34;
          let _lhs329 = globalState;
          let _lhs330 = globalState;
          _lhs327.f12 = _rhs379;
          _2245_v57 = _rhs380;
          _lhs328.f24 = _rhs381;
          _lhs329.f3 = _rhs382;
          _lhs330.f2 = _rhs383;
          let _2250_v60;
          let _init57 = ((_2251_v58) => function (_2252_i8) {
            return _module.D20.create_DC52(_2251_v58);
          })(_2246_v58);
          let _nw381 = Array((new BigNumber(14)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw381.length); _i0_57++) {
            _nw381[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2250_v60 = _nw381;
          let _2253_v63;
          _2253_v63 = _module.D20.create_DC52(function () {
  let _coll56 = new _dafny.Map();
  for (const _compr_56 of _dafny.IntegerRange(new BigNumber(-108), new BigNumber(430))) {
    let _2254_v61 = _compr_56;
    if (((new BigNumber(-108)).isLessThanOrEqualTo(_2254_v61)) && ((_2254_v61).isLessThan(new BigNumber(430)))) {
      _coll56.push([_module.__default.safeModuloInt(_2254_v61, _this.f23),new BigNumber((function () {
        let _coll57 = new _dafny.Map();
        for (const _compr_57 of _dafny.IntegerRange(new BigNumber(424), new BigNumber(779))) {
          let _2255_v62 = _compr_57;
          if (((new BigNumber(424)).isLessThanOrEqualTo(_2255_v62)) && ((_2255_v62).isLessThan(new BigNumber(779)))) {
            _coll57.push([(_2255_v62).plus(p1),true]);
          }
        }
        return _coll57;
      }()).length)]);
    }
  }
  return _coll56;
}());
          let _index371 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_2250_v60).length));
          (_2250_v60)[_index371] = _2253_v63;
          let _index372 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_2250_v60).length));
          (_2250_v60)[_index372] = _2253_v63;
        } else if (_source19.is_DC2) {
          let _2256___mcc_h6 = (_source19).cf6;
          let _2257_cf6 = _2256___mcc_h6;
          (globalState).f2 = (_this).f18;
          let _2258_v64;
          _2258_v64 = new _dafny.CodePoint('f'.codePointAt(0));
          let _2259_v65;
          _2259_v65 = _dafny.MultiSet.fromElements(p0, p0, _dafny.Seq.update(((false) ? (p0) : (p0)), _module.__default.safeIndex(_this.f23, new BigNumber((((false) ? (p0) : (p0))).length)), _2258_v64), _dafny.Seq.UnicodeFromString("dlsydos"), _dafny.Seq.Concat(p0, p2));
          _2259_v65 = _module.__default.fm52(_2207_v34.f24, _2207_v34.f24, globalState);
          _2258_v64 = _2258_v64;
          let _2260_v66;
          _2260_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2207_v34.f24,(_2257_cf6).f18);
          (_2207_v34).f24 = (_2207_v34.f24).plus(new BigNumber(((_2260_v66).Merge(_2260_v66)).length));
        } else {
          let _2261___mcc_h7 = (_source19).cf13;
          let _2262_cf13 = _2261___mcc_h7;
          let _2263_v67;
          let _nw382 = Array((new BigNumber(2)).toNumber()).fill(false);
          _2263_v67 = _nw382;
          let _2264_v68;
          _2264_v68 = _dafny.Set.fromElements(_2263_v67, _2263_v67);
          (globalState).f2 = ((_2264_v68).Difference(_2264_v68)).IsSubsetOf(_2264_v68);
          let _2265_v69;
          let _nw383 = new _module.C0();
          _nw383.__ctor(_2207_v34.f24, (_dafny.ZERO).minus(p1));
          _2265_v69 = _nw383;
          let _2266_v70;
          let _init58 = function (_2267_i9) {
            return _dafny.Seq.Concat(_dafny.Seq.of((_this).f18), _dafny.Seq.of(!(true)));
          };
          let _nw384 = Array((new BigNumber(7)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw384.length); _i0_58++) {
            _nw384[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _2266_v70 = _nw384;
          let _index373 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2266_v70).length));
          (_2266_v70)[_index373] = _dafny.Seq.of(_module.__default.fm23(_2161_v0, _2207_v34.f24, globalState), (_this).f18, !((_this).f18), !((_this).f18), (_this).f18);
          let _2268_v71;
          _2268_v71 = _dafny.Seq.of((_this).f18, (_this).f18);
          let _index374 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2266_v70).length));
          (_2266_v70)[_index374] = _dafny.Seq.Concat(_dafny.Seq.of(!((_this).f18)), _2268_v71);
          (_2207_v34).f24 = (_dafny.ZERO).minus(_2207_v34.f24);
        }
        let _2269_v72;
        _2269_v72 = _module.D6.create_DC18();
        let _source20 = _2269_v72;
        if (_source20.is_DC18) {
          let _2270_v73;
          let _init59 = function (_2271_i10) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          };
          let _nw385 = Array((new BigNumber(4)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw385.length); _i0_59++) {
            _nw385[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _2270_v73 = _nw385;
          let _2272_v74;
          _2272_v74 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18);
          let _2273_v75;
          _2273_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2272_v74,(_2207_v34).fm10((_this).f18, (_this).f18, (_this).f18, globalState));
          let _2274_v76;
          _2274_v76 = _module.D3.create_DC12(_2270_v73, (_this).f18, p1, _2273_v75, (_this).f18);
          (globalState).f3 = ((_dafny.ZERO).minus(((_2209_v36)[_module.__default.safeIndex(p1, new BigNumber((_2209_v36).length))]).minus((_2274_v76).dtor_cf27))).multipliedBy(_this.f23);
          let _2275_v77;
          _2275_v77 = new _dafny.CodePoint('i'.codePointAt(0));
          let _2276_v78;
          _2276_v78 = _module.D20.create_DC53((_this).f18, (_this).f18, _2275_v77);
          let _2277_v79;
          _2277_v79 = _dafny.Seq.of(_2208_v35, _2208_v35, _2208_v35, _2208_v35, _2208_v35);
          let _rhs384 = _2276_v78;
          let _rhs385 = _dafny.Seq.Concat(_2277_v79, _2277_v79);
          _2276_v78 = _rhs384;
          _2277_v79 = _rhs385;
          (_2207_v34).f24 = _2207_v34.f24;
          let _2278_v80;
          let _nw386 = Array((new BigNumber(15)).toNumber()).fill(false);
          _2278_v80 = _nw386;
          let _2279_v81;
          let _nw387 = Array((new BigNumber(6)).toNumber()).fill(_module.D1.Default());
          _2279_v81 = _nw387;
          let _2280_v82;
          _2280_v82 = _module.D1.create_DC4((_this).f18);
          let _index375 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_2279_v81).length));
          (_2279_v81)[_index375] = _2280_v82;
          let _2281_v83;
          _2281_v83 = _dafny.Seq.of(p0);
          let _2282_v84;
          _2282_v84 = _dafny.Seq.of((_this).f18, (_this).f18);
          let _2283_v85;
          _2283_v85 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,new BigNumber(((_2281_v83)[_module.__default.safeIndex(new BigNumber((_2282_v84).length), new BigNumber((_2281_v83).length))]).length));
          let _2284_v86;
          _2284_v86 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2283_v85);
          let _2285_v87;
          _2285_v87 = _dafny.Seq.of(_2284_v86);
          let _2286_v90;
          let _nw388 = Array((new BigNumber(25)).toNumber());
          _nw388[(_dafny.ZERO).toNumber()] = _2284_v86;
          _nw388[(_dafny.ONE).toNumber()] = (_2284_v86).Merge((_2284_v86).update(new BigNumber((_2283_v85).length), _2283_v85));
          _nw388[(new BigNumber(2)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(3)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2207_v34.f24,_2283_v85);
          _nw388[(new BigNumber(5)).toNumber()] = (_2285_v87)[_module.__default.safeIndex(new BigNumber(951), new BigNumber((_2285_v87).length))];
          _nw388[(new BigNumber(6)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_2283_v85);
          _nw388[(new BigNumber(8)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(9)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(10)).toNumber()] = (((_2282_v84)[_module.__default.safeIndex(_2207_v34.f24, new BigNumber((_2282_v84).length))]) ? (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_this.f23, (_this).f18, (_this).f18, globalState),function () {
            let _coll58 = new _dafny.Map();
            for (const _compr_58 of _dafny.IntegerRange(new BigNumber(-146), new BigNumber(-195))) {
              let _2287_v88 = _compr_58;
              if (((new BigNumber(-146)).isLessThanOrEqualTo(_2287_v88)) && ((_2287_v88).isLessThan(new BigNumber(-195)))) {
                _coll58.push([(_2287_v88).multipliedBy((_dafny.ZERO).minus(new BigNumber((_2283_v85).length))),(_2209_v36)[_module.__default.safeIndex(new BigNumber((_2282_v84).length), new BigNumber((_2209_v36).length))]]);
              }
            }
            return _coll58;
          }())) : (_2284_v86));
          _nw388[(new BigNumber(11)).toNumber()] = (_2284_v86).Merge(_2284_v86);
          _nw388[(new BigNumber(12)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(13)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(14)).toNumber()] = _module.__default.fm60(_2282_v84, (_this).f18, (_this).f18, globalState);
          _nw388[(new BigNumber(15)).toNumber()] = function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(884), new BigNumber(857))) {
              let _2288_v89 = _compr_59;
              if (((new BigNumber(884)).isLessThanOrEqualTo(_2288_v89)) && ((_2288_v89).isLessThan(new BigNumber(857)))) {
                _coll59.push([_module.__default.safeModuloInt(_2288_v89, _2207_v34.f24),_2283_v85]);
              }
            }
            return _coll59;
          }();
          _nw388[(new BigNumber(16)).toNumber()] = (_2284_v86).Merge(_2284_v86);
          _nw388[(new BigNumber(17)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(18)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(19)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(20)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(21)).toNumber()] = (_2284_v86).Merge(_2284_v86);
          _nw388[(new BigNumber(22)).toNumber()] = _2284_v86;
          _nw388[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe(_2207_v34.f24,new BigNumber((_dafny.Seq.UnicodeFromString("babsckc")).length)));
          _nw388[(new BigNumber(24)).toNumber()] = (_2284_v86).update(_2207_v34.f24, _2283_v85);
          _2286_v90 = _nw388;
          let _index376 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_2286_v90).length));
          (_2286_v90)[_index376] = _2284_v86;
          let _2289_v91;
          _2289_v91 = _module.D2.create_DC7(_2207_v34.f24, (_this).f18, new BigNumber(956));
          let _index377 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_2279_v81).length));
          let _index378 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_2286_v90).length));
          let _rhs386 = _2278_v80;
          let _rhs387 = (new BigNumber(-41)).multipliedBy(new BigNumber(974));
          let _rhs388 = _module.D1.create_DC4(!((_this).f18));
          let _rhs389 = _module.__default.fm40(_2289_v91, globalState);
          let _rhs390 = _2284_v86;
          let _lhs331 = globalState;
          let _lhs332 = _2279_v81;
          let _lhs333 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_2279_v81).length));
          let _lhs334 = globalState;
          let _lhs335 = _2286_v90;
          let _lhs336 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_2286_v90).length));
          _2278_v80 = _rhs386;
          _lhs331.f0 = _rhs387;
          _lhs332[_lhs333] = _rhs388;
          _lhs334.f10 = _rhs389;
          _lhs335[_lhs336] = _rhs390;
        } else {
          let _2290___mcc_h8 = (_source20).cf38;
          let _2291_cf38 = _2290___mcc_h8;
          let _rhs391 = (_this).f18;
          let _rhs392 = _2207_v34.f24;
          let _rhs393 = (_this).f18;
          let _rhs394 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(763)), function (_2292_i11) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          });
          let _lhs337 = globalState;
          let _lhs338 = globalState;
          let _lhs339 = globalState;
          let _lhs340 = globalState;
          _lhs337.f2 = _rhs391;
          _lhs338.f3 = _rhs392;
          _lhs339.f2 = _rhs393;
          _lhs340.f9 = _rhs394;
          (globalState).f2 = (_this).f18;
          (globalState).f2 = (_this).f18;
          let _2293_v92;
          let _2294_v93;
          let _2295_v94;
          let _2296_v95;
          let _out58;
          let _out59;
          let _out60;
          let _out61;
          let _outcollector16 = _module.__default.m0(globalState);
          _out58 = _outcollector16[0];
          _out59 = _outcollector16[1];
          _out60 = _outcollector16[2];
          _out61 = _outcollector16[3];
          _2293_v92 = _out58;
          _2294_v93 = _out59;
          _2295_v94 = _out60;
          _2296_v95 = _out61;
        }
      } else {
        let _2297_v96;
        _2297_v96 = _dafny.Seq.of(true, false);
        let _2298_v97;
        _2298_v97 = _dafny.MultiSet.fromElements(!((_this).f18));
        let _2299_v98;
        _2299_v98 = _dafny.Map.Empty.slice().updateUnsafe(_2207_v34.f24,(_this).f18);
        let _2300_v99;
        _2300_v99 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
        let _2301_v100;
        _2301_v100 = _dafny.Map.Empty.slice().updateUnsafe(_2300_v99,p1);
        let _2302_v101;
        _2302_v101 = _dafny.Seq.of(_this.f23, _2207_v34.f24, new BigNumber((_2301_v100).length), p1, _2207_v34.f24);
        let _2303_v102;
        let _nw389 = new _module.C3();
        _nw389.__ctor(_2302_v101, p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), ((_2304_v34) => function (_2305_i12) {
          return _dafny.Map.Empty.slice().updateUnsafe(_2304_v34.f24,new BigNumber((_dafny.Set.fromElements((_this).f18, (_this).f18, (_this).f18)).length));
        })(_2207_v34)), (_this).f18);
        _2303_v102 = _nw389;
        let _2306_v103;
        _2306_v103 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ljsg"),_2207_v34.f24);
        let _2307_v104;
        let _nw390 = Array((new BigNumber(17)).toNumber());
        _nw390[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.of(_2297_v96, _2297_v96)).length);
        _nw390[(_dafny.ONE).toNumber()] = (((_2298_v97).contains((((_2299_v98).contains(new BigNumber((p0).length))) ? ((_2299_v98).get(new BigNumber((p0).length))) : ((_2297_v96)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_2297_v96).length))])))) ? ((_2298_v97).get((((_2299_v98).contains(new BigNumber((p0).length))) ? ((_2299_v98).get(new BigNumber((p0).length))) : ((_2297_v96)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_2297_v96).length))])))) : (new BigNumber((_2302_v101).length)));
        _nw390[(new BigNumber(2)).toNumber()] = p1;
        _nw390[(new BigNumber(3)).toNumber()] = _2207_v34.f24;
        _nw390[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.of(_2303_v102)).length);
        _nw390[(new BigNumber(5)).toNumber()] = _this.f23;
        _nw390[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2303_v102).f18,(_2303_v102).f18)).length);
        _nw390[(new BigNumber(7)).toNumber()] = new BigNumber((_2306_v103).length);
        _nw390[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_2207_v34.f24);
        _nw390[(new BigNumber(9)).toNumber()] = _2207_v34.f24;
        _nw390[(new BigNumber(10)).toNumber()] = _this.f23;
        _nw390[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2207_v34.f24,_2207_v34.f24)).length);
        _nw390[(new BigNumber(12)).toNumber()] = new BigNumber((_2297_v96).length);
        _nw390[(new BigNumber(13)).toNumber()] = p1;
        _nw390[(new BigNumber(14)).toNumber()] = new BigNumber((_2298_v97).cardinality());
        _nw390[(new BigNumber(15)).toNumber()] = new BigNumber(373);
        _nw390[(new BigNumber(16)).toNumber()] = _2207_v34.f24;
        _2307_v104 = _nw390;
        let _2308_v105;
        _2308_v105 = _dafny.Map.Empty.slice().updateUnsafe(_2307_v104,p2);
        _2308_v105 = (_2308_v105).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2307_v104,p0));
        let _index379 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_2307_v104).length));
        (_2307_v104)[_index379] = (((_2303_v102).f18) ? (_2207_v34.f24) : (p1));
        let _index380 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_2307_v104).length));
        (_2307_v104)[_index380] = (new BigNumber(297)).multipliedBy(_2207_v34.f24);
        let _2309_v106;
        _2309_v106 = new _dafny.CodePoint('m'.codePointAt(0));
        let _2310_v107;
        _2310_v107 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_2309_v106);
        _2310_v107 = (_2310_v107).update(_this.f23, _2309_v106);
        (globalState).f2 = (_this).f18;
        let _2311_v108;
        let _nw391 = Array((new BigNumber(12)).toNumber()).fill([]);
        _2311_v108 = _nw391;
        let _2312_v109;
        let _init60 = function (_2313_i13) {
          return _dafny.MultiSet.fromElements(_this.f23);
        };
        let _nw392 = Array((new BigNumber(9)).toNumber());
        for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw392.length); _i0_60++) {
          _nw392[_i0_60] = _init60(new BigNumber(_i0_60));
        }
        _2312_v109 = _nw392;
        let _2314_v110;
        _2314_v110 = _module.D22.create_DC57(_2312_v109);
        let _index381 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_2311_v108).length));
        (_2311_v108)[_index381] = (_2314_v110).dtor_cf95;
        let _2315_v111;
        _2315_v111 = _module.D17.create_DC41(_2299_v98);
        let _2316_v112;
        let _init61 = function (_2317_i14) {
          return true;
        };
        let _nw393 = Array((new BigNumber(24)).toNumber());
        for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw393.length); _i0_61++) {
          _nw393[_i0_61] = _init61(new BigNumber(_i0_61));
        }
        _2316_v112 = _nw393;
        let _2318_v113;
        _2318_v113 = _module.D10.create_DC25((_2303_v102).f18);
        let _index382 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_2316_v112).length));
        (_2316_v112)[_index382] = (_2318_v113).dtor_cf48;
        let _index383 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_2311_v108).length));
        let _index384 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_2316_v112).length));
        let _rhs395 = _2312_v109;
        let _rhs396 = _2315_v111;
        let _rhs397 = (_this).f18;
        let _rhs398 = (_this).f18;
        let _lhs341 = _2311_v108;
        let _lhs342 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_2311_v108).length));
        let _lhs343 = _2316_v112;
        let _lhs344 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_2316_v112).length));
        let _lhs345 = globalState;
        _lhs341[_lhs342] = _rhs395;
        _2315_v111 = _rhs396;
        _lhs343[_lhs344] = _rhs397;
        _lhs345.f2 = _rhs398;
      }
      let _2319_v114;
      _2319_v114 = _module.D1.create_DC4((_this).f18);
      (globalState).f2 = function (_source21) {
        if (_source21.is_DC3) {
          let _2320___mcc_h9 = (_source21).cf7;
          let _2321___mcc_h10 = (_source21).cf8;
          let _2322___mcc_h11 = (_source21).cf9;
          let _2323___mcc_h12 = (_source21).cf10;
          let _2324___mcc_h13 = (_source21).cf11;
          let _2325_cf11 = _2324___mcc_h13;
          let _2326_cf10 = _2323___mcc_h12;
          let _2327_cf9 = _2322___mcc_h11;
          let _2328_cf8 = _2321___mcc_h10;
          let _2329_cf7 = _2320___mcc_h9;
          return !((_this).f18) || (!((_this).f18));
        } else if (_source21.is_DC4) {
          let _2330___mcc_h14 = (_source21).cf12;
          let _2331_cf12 = _2330___mcc_h14;
          return (_this).f18;
        } else if (_source21.is_DC2) {
          let _2332___mcc_h15 = (_source21).cf6;
          let _2333_cf6 = _2332___mcc_h15;
          return (_dafny.MultiSet.fromElements((_2333_cf6).f18)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f18, (_this).f18)));
        } else {
          let _2334___mcc_h16 = (_source21).cf13;
          let _2335_cf13 = _2334___mcc_h16;
          return (_this).f18;
        }
      }(_2319_v114);
      (globalState).f2 = _module.__default.fm23(_2161_v0, _2207_v34.f24, globalState);
      if ((_this).f18) {
        let _2336_v115;
        _2336_v115 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(p1)).isLessThan(p1),_module.__default.safeDivisionInt(_2207_v34.f24, _this.f23));
        _2336_v115 = (_2336_v115).update((_this).f18, _this.f23);
        let _2337_v116;
        let _init62 = ((_2338_p1) => function (_2339_i15) {
          return _dafny.MultiSet.FromArray(_dafny.Seq.of(_2338_p1));
        })(p1);
        let _nw394 = Array((new BigNumber(28)).toNumber());
        for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw394.length); _i0_62++) {
          _nw394[_i0_62] = _init62(new BigNumber(_i0_62));
        }
        _2337_v116 = _nw394;
        let _2340_v117;
        _2340_v117 = _dafny.MultiSet.fromElements(p1);
        let _index385 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_2337_v116).length));
        (_2337_v116)[_index385] = (_2340_v117).Union(_2340_v117);
        let _index386 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_2337_v116).length));
        (_2337_v116)[_index386] = (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(692)), function (_2341_i16) {
          return new BigNumber(40);
        }))).update(p1, _module.__default.abs(_module.__default.safeModuloInt(new BigNumber(472), _this.f23)));
        (globalState).f9 = _module.__default.fm13((_this).f18, p1, globalState);
        (globalState).f9 = _dafny.Seq.Concat(p2, _dafny.Seq.UnicodeFromString("b"));
        _2336_v115 = _module.__default.fm43((_this).f18, false, globalState);
      } else {
        let _2342_v118;
        let _nw395 = Array((_dafny.ONE).toNumber());
        _nw395[(_dafny.ZERO).toNumber()] = _2207_v34.f24;
        _2342_v118 = _nw395;
        let _2343_v119;
        _2343_v119 = _module.D17.create_DC42((_this).f18);
        let _2344_v120;
        _2344_v120 = _dafny.Map.Empty.slice().updateUnsafe((_2343_v119).dtor_cf74,_2207_v34.f24);
        let _2345_v121;
        _2345_v121 = _dafny.Set.fromElements(_2344_v120);
        let _2346_v122;
        _2346_v122 = _dafny.Map.Empty.slice().updateUnsafe(_2342_v118,(p1).plus(new BigNumber((_2345_v121).length)));
        _2346_v122 = (_2346_v122).update(_2342_v118, (_this.f23).minus(_this.f23));
        (globalState).f2 = !(((_this).f18) && ((_this).f18)) || ((_this).f18);
        let _2347_v123;
        _2347_v123 = _dafny.Set.fromElements((_this).f18, (_this).f18);
        let _2348_v124;
        _2348_v124 = _dafny.Seq.of(_2347_v123, _2347_v123);
        let _2349_v125;
        _2349_v125 = new _dafny.CodePoint('h'.codePointAt(0));
        let _2350_v126;
        _2350_v126 = _module.D3.create_DC13(_2207_v34.f24, _dafny.Seq.Concat(p2, p0), (_2207_v34.f24).plus(_module.__default.fm1(new BigNumber((_dafny.MultiSet.FromArray(_2348_v124)).cardinality()), (_this).f18, (_this).f18, globalState)), _2207_v34.f24, _2349_v125);
        let _source22 = _2350_v126;
        if (_source22.is_DC11) {
          let _2351___mcc_h17 = (_source22).cf23;
          let _2352___mcc_h18 = (_source22).cf24;
          let _2353_cf24 = _2352___mcc_h18;
          let _2354_cf23 = _2351___mcc_h17;
          let _2355_v127;
          let _init63 = ((_2356_cf23) => function (_2357_i17) {
            return _2356_cf23;
          })(_2354_cf23);
          let _nw396 = Array((new BigNumber(9)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw396.length); _i0_63++) {
            _nw396[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _2355_v127 = _nw396;
          let _2358_v128;
          _2358_v128 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12((_dafny.ZERO).minus(_2207_v34.f24), _2354_cf23, globalState),_2355_v127);
          let _2359_v129;
          _2359_v129 = _dafny.MultiSet.fromElements(new BigNumber(670), _2207_v34.f24);
          let _rhs399 = (_2358_v128).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ocfdgltdd"),_2355_v127));
          let _rhs400 = (_2359_v129).IsSubsetOf(_2359_v129);
          _2358_v128 = _rhs399;
          _2354_cf23 = _rhs400;
          let _2360_v130;
          _2360_v130 = _module.D13.create_DC33((_this).f18, _2207_v34.f24, _2354_cf23, new BigNumber(-352));
          let _2361_v131;
          _2361_v131 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_2354_cf23);
          let _2362_v132;
          _2362_v132 = _module.D16.create_DC39(false, _this.f23);
          let _2363_v133;
          _2363_v133 = _module.D2.create_DC7(new BigNumber((_2359_v129).cardinality()), (_2362_v132).dtor_cf67, p1);
          let _2364_v134;
          _2364_v134 = _dafny.Map.Empty.slice().updateUnsafe(_2354_cf23,_2353_cf24);
          let _pat_let_tv33 = _2353_cf24;
          let _2365_v135;
          let _nw397 = Array((new BigNumber(14)).toNumber());
          _nw397[(_dafny.ZERO).toNumber()] = _2360_v130;
          _nw397[(_dafny.ONE).toNumber()] = _2360_v130;
          _nw397[(new BigNumber(2)).toNumber()] = _module.D13.create_DC33((((_2361_v131).contains(_2207_v34.f24)) ? ((_2361_v131).get(_2207_v34.f24)) : (_2354_cf23)), new BigNumber(553), true, _2207_v34.f24);
          _nw397[(new BigNumber(3)).toNumber()] = _2360_v130;
          _nw397[(new BigNumber(4)).toNumber()] = function (_pat_let63_0) {
            return function (_2366_dt__update__tmp_h3) {
              return function (_pat_let64_0) {
                return function (_2367_dt__update_hcf57_h0) {
                  return function (_pat_let65_0) {
                    return function (_2368_dt__update_hcf55_h0) {
                      return _module.D13.create_DC33(_2368_dt__update_hcf55_h0, (_2366_dt__update__tmp_h3).dtor_cf56, _2367_dt__update_hcf57_h0, (_2366_dt__update__tmp_h3).dtor_cf58);
                    }(_pat_let65_0);
                  }((_this).f18);
                }(_pat_let64_0);
              }(_pat_let_tv33);
            }(_pat_let63_0);
          }(_2360_v130);
          _nw397[(new BigNumber(5)).toNumber()] = _2360_v130;
          _nw397[(new BigNumber(6)).toNumber()] = _2360_v130;
          _nw397[(new BigNumber(7)).toNumber()] = _2360_v130;
          _nw397[(new BigNumber(8)).toNumber()] = _2360_v130;
          _nw397[(new BigNumber(9)).toNumber()] = _2360_v130;
          _nw397[(new BigNumber(10)).toNumber()] = ((_module.__default.fm23(_2161_v0, _this.f23, globalState)) ? (_module.D13.create_DC33(true, p1, _2354_cf23, _2207_v34.f24)) : (_2360_v130));
          _nw397[(new BigNumber(11)).toNumber()] = _2360_v130;
          _nw397[(new BigNumber(12)).toNumber()] = _module.D13.create_DC33((_2363_v133).dtor_cf16, p1, _2353_cf24, new BigNumber((_2364_v134).length));
          _nw397[(new BigNumber(13)).toNumber()] = _2360_v130;
          _2365_v135 = _nw397;
          _2365_v135 = _2365_v135;
          let _2369_v136;
          _2369_v136 = _dafny.Seq.of(_2354_cf23, (_this).f18);
          let _2370_v137;
          _2370_v137 = _dafny.Map.Empty.slice().updateUnsafe(_2369_v136,p1);
          let _2371_v138;
          _2371_v138 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_2369_v136,(_dafny.ZERO).minus(p1))).update(_2369_v136, new BigNumber(151)));
          let _rhs401 = (_this).f18;
          let _rhs402 = (_2371_v138)[_module.__default.safeIndex(_2207_v34.f24, new BigNumber((_2371_v138).length))];
          let _rhs403 = (_dafny.ZERO).minus((new BigNumber((p2).length)).multipliedBy(p1));
          let _lhs346 = globalState;
          let _lhs347 = globalState;
          _lhs346.f2 = _rhs401;
          _2370_v137 = _rhs402;
          _lhs347.f0 = _rhs403;
          let _index387 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_2342_v118).length));
          (_2342_v118)[_index387] = (new BigNumber((p0).length)).minus(_2207_v34.f24);
          let _index388 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_2342_v118).length));
          (_2342_v118)[_index388] = new BigNumber(813);
        } else if (_source22.is_DC12) {
          let _2372___mcc_h19 = (_source22).cf25;
          let _2373___mcc_h20 = (_source22).cf26;
          let _2374___mcc_h21 = (_source22).cf27;
          let _2375___mcc_h22 = (_source22).cf28;
          let _2376___mcc_h23 = (_source22).cf29;
          let _2377_cf29 = _2376___mcc_h23;
          let _2378_cf28 = _2375___mcc_h22;
          let _2379_cf27 = _2374___mcc_h21;
          let _2380_cf26 = _2373___mcc_h20;
          let _2381_cf25 = _2372___mcc_h19;
          (globalState).f15 = (_dafny.ZERO).minus((_2207_v34.f24).minus((new BigNumber(-724)).multipliedBy(p1)));
          (globalState).f2 = (_this).f18;
          let _2382_v139;
          let _nw398 = Array((new BigNumber(20)).toNumber()).fill(false);
          _2382_v139 = _nw398;
          let _index389 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_2382_v139).length));
          (_2382_v139)[_index389] = _2377_cf29;
          let _index390 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_2382_v139).length));
          (_2382_v139)[_index390] = (_this).f18;
          _2379_cf27 = new BigNumber(3);
        } else if (_source22.is_DC13) {
          let _2383___mcc_h24 = (_source22).cf30;
          let _2384___mcc_h25 = (_source22).cf31;
          let _2385___mcc_h26 = (_source22).cf32;
          let _2386___mcc_h27 = (_source22).cf33;
          let _2387___mcc_h28 = (_source22).cf34;
          let _2388_cf34 = _2387___mcc_h28;
          let _2389_cf33 = _2386___mcc_h27;
          let _2390_cf32 = _2385___mcc_h26;
          let _2391_cf31 = _2384___mcc_h25;
          let _2392_cf30 = _2383___mcc_h24;
          let _2393_v140;
          let _nw399 = Array((new BigNumber(8)).toNumber()).fill([]);
          _2393_v140 = _nw399;
          let _2394_v141;
          let _nw400 = Array((new BigNumber(23)).toNumber());
          _nw400[(_dafny.ZERO).toNumber()] = _2388_cf34;
          _nw400[(_dafny.ONE).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(2)).toNumber()] = _2349_v125;
          _nw400[(new BigNumber(3)).toNumber()] = _2349_v125;
          _nw400[(new BigNumber(4)).toNumber()] = (p2)[_module.__default.safeIndex(_this.f23, new BigNumber((p2).length))];
          _nw400[(new BigNumber(5)).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(6)).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(7)).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(8)).toNumber()] = (_2350_v126).dtor_cf34;
          _nw400[(new BigNumber(9)).toNumber()] = _2349_v125;
          _nw400[(new BigNumber(10)).toNumber()] = _2349_v125;
          _nw400[(new BigNumber(11)).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(12)).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(13)).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(14)).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw400[(new BigNumber(16)).toNumber()] = _2349_v125;
          _nw400[(new BigNumber(17)).toNumber()] = (_2391_cf31)[_module.__default.safeIndex(_2390_cf32, new BigNumber((_2391_cf31).length))];
          _nw400[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
          _nw400[(new BigNumber(19)).toNumber()] = _2388_cf34;
          _nw400[(new BigNumber(20)).toNumber()] = _2349_v125;
          _nw400[(new BigNumber(21)).toNumber()] = _2349_v125;
          _nw400[(new BigNumber(22)).toNumber()] = _2349_v125;
          _2394_v141 = _nw400;
          let _index391 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_2393_v140).length));
          (_2393_v140)[_index391] = _2394_v141;
          let _index392 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_2393_v140).length));
          (_2393_v140)[_index392] = _2394_v141;
          let _rhs404 = _dafny.Set.fromElements(_2392_cf30, _module.__default.fm1((_dafny.ZERO).minus(_2392_cf30), (_this).f18, (_this).f18, globalState), _2389_cf33, _module.__default.safeModuloInt(_this.f23, _2207_v34.f24), _module.__default.safeDivisionInt(p1, _2207_v34.f24));
          let _rhs405 = (_this).f18;
          let _rhs406 = _2349_v125;
          let _lhs348 = globalState;
          _2161_v0 = _rhs404;
          _lhs348.f2 = _rhs405;
          _2388_cf34 = _rhs406;
          (globalState).f9 = p2;
          (_this).f23 = new BigNumber(442);
        } else if (_source22.is_DC10) {
          let _2395___mcc_h29 = (_source22).cf22;
          let _2396_cf22 = _2395___mcc_h29;
          (globalState).f2 = (_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(312)), ((_2397_v125) => function (_2398_i18) {
            return _2397_v125;
          })(_2349_v125)), _module.__default.safeIndex(_2207_v34.f24, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(312)), ((_2399_v125) => function (_2400_i18) {
            return _2399_v125;
          })(_2349_v125))).length)), new _dafny.CodePoint('l'.codePointAt(0))), _dafny.Seq.UnicodeFromString("hbxr"))) || ((_this).f18);
          let _2401_v142;
          _2401_v142 = _dafny.MultiSet.fromElements((_this).f18);
          _2401_v142 = _2401_v142;
          let _2402_v143;
          _2402_v143 = _dafny.Map.Empty.slice().updateUnsafe(_2207_v34.f24,(_this).f18);
          _2402_v143 = (_2402_v143).update((_this.f23).minus(_2207_v34.f24), (_this).f18);
          let _2403_v144;
          let _init64 = function (_2404_i19) {
            return (_this).f18;
          };
          let _nw401 = Array((new BigNumber(24)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw401.length); _i0_64++) {
            _nw401[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2403_v144 = _nw401;
          let _2405_v145;
          _2405_v145 = _module.D3.create_DC11((_this).f18, (_this).f18);
          let _2406_v146;
          _2406_v146 = _dafny.Seq.of(!((_2405_v145).dtor_cf23));
          let _2407_v147;
          _2407_v147 = _dafny.Map.Empty.slice().updateUnsafe(_2207_v34.f24,_2207_v34.f24);
          let _2408_v148;
          _2408_v148 = _dafny.Set.fromElements(_2407_v147);
          let _2409_v149;
          _2409_v149 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2406_v146).length),_2408_v148);
          let _2410_v150;
          _2410_v150 = _module.D21.create_DC55(((((_2409_v149).contains(_this.f23)) ? ((_2409_v149).get(_this.f23)) : (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2161_v0).length),_2207_v34.f24), _2407_v147)))).Intersect(_2408_v148));
          let _2411_v151;
          _2411_v151 = _dafny.Seq.of(_2207_v34.f24);
          let _rhs407 = _2403_v144;
          let _rhs408 = _2349_v125;
          let _rhs409 = _dafny.Seq.IsProperPrefixOf(_2411_v151, _2411_v151);
          let _rhs410 = _module.__default.fm61((_this).f18, globalState);
          let _rhs411 = _2342_v118;
          let _lhs349 = globalState;
          _2403_v144 = _rhs407;
          _2349_v125 = _rhs408;
          _lhs349.f2 = _rhs409;
          _2410_v150 = _rhs410;
          _2342_v118 = _rhs411;
        } else {
          let _2412___mcc_h30 = (_source22).cf35;
          let _2413_cf35 = _2412___mcc_h30;
          let _2414_v152;
          let _nw402 = Array((new BigNumber(15)).toNumber()).fill([]);
          _2414_v152 = _nw402;
          let _2415_v153;
          let _nw403 = Array((_dafny.ONE).toNumber()).fill([]);
          _2415_v153 = _nw403;
          let _index393 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_2414_v152).length));
          (_2414_v152)[_index393] = _2415_v153;
          let _index394 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_2414_v152).length));
          (_2414_v152)[_index394] = _2415_v153;
          (globalState).f2 = (_this).f18;
          let _2416_v154;
          _2416_v154 = _module.D1.create_DC3(_2207_v34.f24, new BigNumber(-308), (_dafny.ZERO).minus((_dafny.ZERO).minus(_2207_v34.f24)), p1, p1);
          let _2417_v155;
          _2417_v155 = _module.D1.create_DC5(_2416_v154);
          _2417_v155 = _2417_v155;
          (globalState).f2 = (_this).f18;
        }
        (globalState).f0 = _module.__default.fm1(_2207_v34.f24, (_this).f18, (_this).f18, globalState);
        let _2418_v156;
        let _init65 = function (_2419_i20) {
          return false;
        };
        let _nw404 = Array((new BigNumber(19)).toNumber());
        for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw404.length); _i0_65++) {
          _nw404[_i0_65] = _init65(new BigNumber(_i0_65));
        }
        _2418_v156 = _nw404;
        let _2420_v157;
        _2420_v157 = _dafny.Seq.of(_2418_v156, _2418_v156, _2418_v156, _2418_v156);
        _2420_v157 = _2420_v157;
      }
      return;
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f18 = false;
      this.f22 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f22, f18) {
      let _this = this;
      (_this).f22 = f22;
      (_this)._f18 = f18;
      return;
    }
    fm4(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber(835)).multipliedBy((((_this).f18) ? (new BigNumber(127)) : (new BigNumber((_dafny.Seq.of((_this).f18, (_this).f18)).length))));
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return (((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-689),(_this).f18)).length))).cardinality()))).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("maumun")).length), new BigNumber(-107))).length))).minus(new BigNumber(848));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _2421_v0;
      _2421_v0 = _module.D0.create_DC0((_this).f18);
      let _2422_i0;
      _2422_i0 = _dafny.ZERO;
      L15: {
        while ((_2421_v0).dtor_cf0) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2422_i0)) {
              break L15;
            }
            _2422_i0 = (_2422_i0).plus(_dafny.ONE);
            let _rhs412 = (_this).f18;
            let _rhs413 = _dafny.areEqual(_dafny.Seq.Concat(p2, _dafny.Seq.UnicodeFromString("srlae")), _dafny.Seq.Concat(p2, p2));
            let _lhs350 = globalState;
            let _lhs351 = globalState;
            _lhs350.f2 = _rhs412;
            _lhs351.f2 = _rhs413;
            let _2423_v1;
            _2423_v1 = _module.D1.create_DC3(p1, p1, (_dafny.ZERO).minus(p1), p1, p1);
            let _2424_v2;
            _2424_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2423_v1);
            _2424_v2 = (_2424_v2).update(_module.__default.safeModuloInt(p1, p1), _2423_v1);
            let _source23 = _2423_v1;
            if (_source23.is_DC3) {
              let _2425___mcc_h0 = (_source23).cf7;
              let _2426___mcc_h1 = (_source23).cf8;
              let _2427___mcc_h2 = (_source23).cf9;
              let _2428___mcc_h3 = (_source23).cf10;
              let _2429___mcc_h4 = (_source23).cf11;
              let _2430_cf11 = _2429___mcc_h4;
              let _2431_cf10 = _2428___mcc_h3;
              let _2432_cf9 = _2427___mcc_h2;
              let _2433_cf8 = _2426___mcc_h1;
              let _2434_cf7 = _2425___mcc_h0;
              let _2435_v3;
              _2435_v3 = _dafny.Set.fromElements((_this).f18);
              let _2436_v4;
              _2436_v4 = _dafny.Seq.of(_module.__default.safeModuloInt(p1, _2433_cf8), ((true) ? (_2431_cf10) : (new BigNumber((_2435_v3).length))), _2431_cf10);
              let _2437_v5;
              _2437_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p1);
              let _2438_v6;
              _2438_v6 = _module.D2.create_DC6(_2437_v5);
              (globalState).f3 = (_2436_v4)[_module.__default.safeIndex((_2433_cf8).multipliedBy(_module.__default.fm1(new BigNumber(((_2438_v6).dtor_cf14).length), (_this).f18, (_this).f18, globalState)), new BigNumber((_2436_v4).length))];
              (globalState).f2 = !(new BigNumber((p0).length)).isEqualTo(_2431_cf10);
              let _2439_v7;
              _2439_v7 = new _dafny.CodePoint('r'.codePointAt(0));
              let _2440_v8;
              _2440_v8 = _dafny.Set.fromElements(_2439_v7);
              let _2441_v9;
              _2441_v9 = _module.D23.create_DC60(_2440_v8);
              let _2442_v10;
              let _nw405 = new _module.C7();
              _nw405.__ctor(new BigNumber((((_2441_v9).dtor_cf98).Union(_2440_v8)).length), (_this).f18);
              _2442_v10 = _nw405;
              (globalState).f15 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_2434_cf7, _2430_cf11), _module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), p1));
            } else if (_source23.is_DC4) {
              let _2443___mcc_h5 = (_source23).cf12;
              let _2444_cf12 = _2443___mcc_h5;
              let _2445_v11;
              let _nw406 = Array((new BigNumber(19)).toNumber());
              _nw406[(_dafny.ZERO).toNumber()] = p2;
              _nw406[(_dafny.ONE).toNumber()] = p0;
              _nw406[(new BigNumber(2)).toNumber()] = p0;
              _nw406[(new BigNumber(3)).toNumber()] = p0;
              _nw406[(new BigNumber(4)).toNumber()] = p2;
              _nw406[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(p2, p0);
              _nw406[(new BigNumber(6)).toNumber()] = p2;
              _nw406[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-357)), function (_2446_i1) {
                return new _dafny.CodePoint('r'.codePointAt(0));
              });
              _nw406[(new BigNumber(8)).toNumber()] = p2;
              _nw406[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("bm");
              _nw406[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nvgq"), p0);
              _nw406[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-59)), function (_2447_i2) {
                return new _dafny.CodePoint('u'.codePointAt(0));
              });
              _nw406[(new BigNumber(12)).toNumber()] = p0;
              _nw406[(new BigNumber(13)).toNumber()] = p0;
              _nw406[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(p0, p2);
              _nw406[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(p2, p2);
              _nw406[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(349)), function (_2448_i3) {
                return new _dafny.CodePoint('u'.codePointAt(0));
              });
              _nw406[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-196)), function (_2449_i4) {
                return new _dafny.CodePoint('j'.codePointAt(0));
              });
              _nw406[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(p2, p0), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-781)), function (_2450_i5) {
                return new _dafny.CodePoint('p'.codePointAt(0));
              })).length), new BigNumber((_dafny.Seq.Concat(p2, p0)).length)), new _dafny.CodePoint('i'.codePointAt(0)));
              _2445_v11 = _nw406;
              let _index395 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_2445_v11).length));
              (_2445_v11)[_index395] = p2;
              let _2451_v12;
              _2451_v12 = _dafny.Set.fromElements(p1, p1);
              let _2452_v13;
              _2452_v13 = _dafny.Set.fromElements(_2451_v12, _2451_v12);
              let _2453_v14;
              _2453_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2444_cf12,_2452_v13);
              let _2454_v15;
              _2454_v15 = _dafny.Seq.of(_2451_v12);
              let _2455_v17;
              _2455_v17 = _dafny.Seq.of(function () {
                let _coll60 = new _dafny.Set();
                for (const _compr_60 of (_2454_v15).Elements) {
                  let _2456_v16 = _compr_60;
                  if (_dafny.Seq.contains(_2454_v15, _2456_v16)) {
                    _coll60.add(_2456_v16);
                  }
                }
                return _coll60;
              }());
              let _2457_v18;
              _2457_v18 = _module.D24.create_DC63(_2452_v13);
              let _2458_v19;
              _2458_v19 = _dafny.Seq.of(p1);
              let _2459_v20;
              _2459_v20 = _dafny.Seq.of(p1, (_2458_v19)[_module.__default.safeIndex(new BigNumber((_2451_v12).length), new BigNumber((_2458_v19).length))]);
              let _2460_v22;
              let _nw407 = Array((new BigNumber(17)).toNumber());
              _nw407[(_dafny.ZERO).toNumber()] = _2452_v13;
              _nw407[(_dafny.ONE).toNumber()] = (((_2453_v14).contains(!((_this).f18))) ? ((_2453_v14).get(!((_this).f18))) : (_dafny.Set.fromElements(_module.__default.fm49(p1, p1, p1, globalState))));
              _nw407[(new BigNumber(2)).toNumber()] = (((_this).f18) ? (_2452_v13) : ((_2455_v17)[_module.__default.safeIndex(p1, new BigNumber((_2455_v17).length))]));
              _nw407[(new BigNumber(3)).toNumber()] = _2452_v13;
              _nw407[(new BigNumber(4)).toNumber()] = (_2457_v18).dtor_cf104;
              _nw407[(new BigNumber(5)).toNumber()] = _2452_v13;
              _nw407[(new BigNumber(6)).toNumber()] = _2452_v13;
              _nw407[(new BigNumber(7)).toNumber()] = _2452_v13;
              _nw407[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(function () {
                let _coll61 = new _dafny.Set();
                for (const _compr_61 of (_2459_v20).Elements) {
                  let _2461_v21 = _compr_61;
                  if (_dafny.Seq.contains(_2459_v20, _2461_v21)) {
                    _coll61.add((_2461_v21).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))).length)));
                  }
                }
                return _coll61;
              }(), _2451_v12);
              _nw407[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(_2451_v12, _dafny.Set.fromElements(p1));
              _nw407[(new BigNumber(10)).toNumber()] = _2452_v13;
              _nw407[(new BigNumber(11)).toNumber()] = _2452_v13;
              _nw407[(new BigNumber(12)).toNumber()] = (_2452_v13).Intersect(_2452_v13);
              _nw407[(new BigNumber(13)).toNumber()] = (_2452_v13).Difference(_2452_v13);
              _nw407[(new BigNumber(14)).toNumber()] = (((_this).f18) ? (_2452_v13) : (_2452_v13));
              _nw407[(new BigNumber(15)).toNumber()] = _2452_v13;
              _nw407[(new BigNumber(16)).toNumber()] = _2452_v13;
              _2460_v22 = _nw407;
              let _index396 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_2460_v22).length));
              (_2460_v22)[_index396] = _2452_v13;
              let _index397 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_2445_v11).length));
              let _index398 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_2460_v22).length));
              let _rhs414 = p0;
              let _rhs415 = p1;
              let _rhs416 = (((_this).f18) ? ((_2452_v13).Union(_2452_v13)) : (_2452_v13));
              let _lhs352 = _2445_v11;
              let _lhs353 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_2445_v11).length));
              let _lhs354 = globalState;
              let _lhs355 = _2460_v22;
              let _lhs356 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_2460_v22).length));
              _lhs352[_lhs353] = _rhs414;
              _lhs354.f3 = _rhs415;
              _lhs355[_lhs356] = _rhs416;
              (globalState).f15 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), p1);
              let _2462_v23;
              let _nw408 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
              _2462_v23 = _nw408;
              _2462_v23 = _2462_v23;
              (globalState).f15 = p1;
            } else if (_source23.is_DC2) {
              let _2463___mcc_h6 = (_source23).cf6;
              let _2464_cf6 = _2463___mcc_h6;
              let _2465_v24;
              _2465_v24 = _dafny.Set.fromElements((_2464_cf6).f18);
              let _2466_v25;
              let _nw409 = new _module.C5();
              _nw409.__ctor(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("w"), p0), p2, (_2465_v24).IsProperSubsetOf(_2465_v24));
              _2466_v25 = _nw409;
              (globalState).f3 = p1;
              let _2467_v26;
              _2467_v26 = _dafny.Seq.of(_module.__default.fm36((_2464_cf6).f18, _module.__default.fm1(p1, (_2464_cf6).f18, (_2464_cf6).f18, globalState), false, globalState), _2465_v24, _2465_v24);
              let _2468_v27;
              _2468_v27 = new _dafny.CodePoint('n'.codePointAt(0));
              let _2469_v28;
              _2469_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
              let _2470_v29;
              _2470_v29 = _dafny.Map.Empty.slice().updateUnsafe(false,_2469_v28);
              let _2471_v30;
              _2471_v30 = _dafny.MultiSet.fromElements(p1, new BigNumber(561), new BigNumber((_2470_v29).length), p1);
              let _2472_v31;
              _2472_v31 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
              let _2473_v32;
              _2473_v32 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).fm5(_2471_v30, new BigNumber((_2472_v31).length), globalState)),_2468_v27);
              let _2474_v33;
              let _nw410 = Array((new BigNumber(12)).toNumber());
              _nw410[(_dafny.ZERO).toNumber()] = _2468_v27;
              _nw410[(_dafny.ONE).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(2)).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(3)).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(4)).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
              _nw410[(new BigNumber(6)).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(7)).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(8)).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(9)).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(10)).toNumber()] = _2468_v27;
              _nw410[(new BigNumber(11)).toNumber()] = (((_2473_v32).contains(p1)) ? ((_2473_v32).get(p1)) : (_2468_v27));
              _2474_v33 = _nw410;
              let _2475_v34;
              _2475_v34 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2474_v33);
              let _2476_v35;
              _2476_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2475_v34).length),p1);
              _2467_v26 = _module.__default.fm62(p1, _2476_v35, globalState);
              let _2477_v36;
              _2477_v36 = _dafny.Seq.of(_this.f22);
              _2477_v36 = _2477_v36;
            } else {
              let _2478___mcc_h7 = (_source23).cf13;
              let _2479_cf13 = _2478___mcc_h7;
              let _2480_v37;
              _2480_v37 = _dafny.MultiSet.fromElements(true, (_this).f18, (_this).f18);
              let _2481_v38;
              _2481_v38 = _2480_v37;
              let _2482_v39;
              _2482_v39 = _dafny.Seq.of(p1, new BigNumber(539));
              let _2483_v40;
              _2483_v40 = _dafny.Seq.of((_this).f18, (_this).f18);
              let _2484_v41;
              _2484_v41 = _dafny.MultiSet.fromElements(p1, new BigNumber((_2483_v40).length));
              let _2485_v42;
              _2485_v42 = _module.D16.create_DC38(_2484_v41);
              let _2486_v43;
              let _nw411 = Array((new BigNumber(15)).toNumber());
              _nw411[(_dafny.ZERO).toNumber()] = (_2480_v37).update((_this).f18, _module.__default.abs(p1));
              _nw411[(_dafny.ONE).toNumber()] = _2480_v37;
              _nw411[(new BigNumber(2)).toNumber()] = _2480_v37;
              _nw411[(new BigNumber(3)).toNumber()] = (_2480_v37).Union(_dafny.MultiSet.fromElements((_this).f18, true, (_this).f18, (_this).f18, (_this).f18));
              _nw411[(new BigNumber(4)).toNumber()] = _2480_v37;
              _nw411[(new BigNumber(5)).toNumber()] = _2480_v37;
              _nw411[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(true);
              _nw411[(new BigNumber(7)).toNumber()] = (_2481_v38);
              _nw411[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.fromElements((_this).f18, (_this).f18)).Difference(_2480_v37);
              _nw411[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_module.__default.fm6(_2482_v39, (_this).f18, (_2485_v42).dtor_cf66, p1, globalState));
              _nw411[(new BigNumber(10)).toNumber()] = _2480_v37;
              _nw411[(new BigNumber(11)).toNumber()] = (_2480_v37).Difference((_module.__default.fm32(globalState)).update((_this).f18, _module.__default.abs(new BigNumber(213))));
              _nw411[(new BigNumber(12)).toNumber()] = (_module.__default.fm32(globalState)).Difference(_dafny.MultiSet.FromArray(_2483_v40));
              _nw411[(new BigNumber(13)).toNumber()] = _2480_v37;
              _nw411[(new BigNumber(14)).toNumber()] = _2480_v37;
              _2486_v43 = _nw411;
              let _index399 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2486_v43).length));
              (_2486_v43)[_index399] = _2480_v37;
              let _2487_v44;
              let _init66 = ((_2488_p1) => function (_2489_i6) {
                return _module.__default.safeDivisionInt(_2489_i6, _2488_p1);
              })(p1);
              let _nw412 = Array((new BigNumber(9)).toNumber());
              for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw412.length); _i0_66++) {
                _nw412[_i0_66] = _init66(new BigNumber(_i0_66));
              }
              _2487_v44 = _nw412;
              let _index400 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_2487_v44).length));
              (_2487_v44)[_index400] = (_dafny.ZERO).minus(p1);
              let _2490_v45;
              let _nw413 = new _module.C0();
              _nw413.__ctor(p1, new BigNumber((_2482_v39).length));
              _2490_v45 = _nw413;
              let _2491_v46;
              _2491_v46 = _dafny.Seq.of(_2490_v45);
              let _index401 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_2487_v44).length));
              (_2487_v44)[_index401] = (_2490_v45).f32;
              let _index402 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2486_v43).length));
              let _index403 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_2487_v44).length));
              let _index404 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_2487_v44).length));
              let _rhs417 = _2480_v37;
              let _rhs418 = (_2490_v45).f31;
              let _rhs419 = _2491_v46;
              let _rhs420 = (_this).f18;
              let _rhs421 = (_dafny.ZERO).minus(p1);
              let _lhs357 = _2486_v43;
              let _lhs358 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2486_v43).length));
              let _lhs359 = _2487_v44;
              let _lhs360 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_2487_v44).length));
              let _lhs361 = globalState;
              let _lhs362 = _2487_v44;
              let _lhs363 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_2487_v44).length));
              _lhs357[_lhs358] = _rhs417;
              _lhs359[_lhs360] = _rhs418;
              _2491_v46 = _rhs419;
              _lhs361.f2 = _rhs420;
              _lhs362[_lhs363] = _rhs421;
              (globalState).f3 = (_2487_v44)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_2487_v44).length))];
              let _2492_v47;
              let _nw414 = new _module.C1();
              _nw414.__ctor((_this).f18);
              _2492_v47 = _nw414;
              let _2493_v48;
              _2493_v48 = new _dafny.CodePoint('c'.codePointAt(0));
              let _2494_v49;
              _2494_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2493_v48,(_this).f18);
              let _2495_v50;
              _2495_v50 = _dafny.Set.fromElements(_2494_v49);
              let _2496_v51;
              _2496_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2482_v39,_2495_v50);
              _2496_v51 = (_2496_v51).update(_2482_v39, (_2495_v50).Difference(_2495_v50));
            }
            (globalState).f3 = _module.__default.safeDivisionInt(p1, p1);
          }
        }
      }
      let _2497_v52;
      _2497_v52 = _module.D17.create_DC41(_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f18));
      let _2498_v53;
      _2498_v53 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f18);
      let _2499_v54;
      _2499_v54 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p1));
      let _2500_v55;
      _2500_v55 = _dafny.Seq.of(p1);
      let _2501_v56;
      _2501_v56 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_2500_v55).length)), (_dafny.ZERO).minus(p1));
      let _2502_v57;
      _2502_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(172),p1);
      let _2503_v58;
      _2503_v58 = _dafny.Seq.of(_2499_v54);
      let _2504_v59;
      _2504_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.update(_2500_v55, _module.__default.safeIndex(p1, new BigNumber((_2500_v55).length)), new BigNumber(302)));
      let _pat_let_tv34 = p1;
      let _pat_let_tv35 = _2498_v53;
      let _2505_v60;
      let _nw415 = Array((new BigNumber(16)).toNumber());
      _nw415[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(p1, new BigNumber(((function (_pat_let66_0) {
        return function (_2508_dt__update__tmp_h1) {
          return function (_pat_let69_0) {
            return function (_2509_dt__update_hcf73_h1) {
              return _module.D17.create_DC41(_2509_dt__update_hcf73_h1);
            }(_pat_let69_0);
          }(_pat_let_tv35);
        }(_pat_let66_0);
      }(function (_pat_let67_0) {
        return function (_2506_dt__update__tmp_h0) {
          return function (_pat_let68_0) {
            return function (_2507_dt__update_hcf73_h0) {
              return _module.D17.create_DC41(_2507_dt__update_hcf73_h0);
            }(_pat_let68_0);
          }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv34,(_this).f18));
        }(_pat_let67_0);
      }(_2497_v52))).dtor_cf73).length), new BigNumber((p2).length)));
      _nw415[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(p1);
      _nw415[(new BigNumber(2)).toNumber()] = _2499_v54;
      _nw415[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(p1);
      _nw415[(new BigNumber(4)).toNumber()] = _2499_v54;
      _nw415[(new BigNumber(5)).toNumber()] = _module.__default.fm25((_dafny.ZERO).minus(p1), globalState);
      _nw415[(new BigNumber(6)).toNumber()] = _2499_v54;
      _nw415[(new BigNumber(7)).toNumber()] = _2499_v54;
      _nw415[(new BigNumber(8)).toNumber()] = _2499_v54;
      _nw415[(new BigNumber(9)).toNumber()] = ((_module.__default.fm6(_2501_v56, (_this).f18, _2499_v54, new BigNumber((_2502_v57).length), globalState)) ? (_dafny.MultiSet.fromElements((_2500_v55)[_module.__default.safeIndex(p1, new BigNumber((_2500_v55).length))], p1, p1)) : (_2499_v54));
      _nw415[(new BigNumber(10)).toNumber()] = (_2503_v58)[_module.__default.safeIndex(p1, new BigNumber((_2503_v58).length))];
      _nw415[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.update((((_2504_v59).contains((_this).f18)) ? ((_2504_v59).get((_this).f18)) : (_2500_v55)), _module.__default.safeIndex(p1, new BigNumber(((((_2504_v59).contains((_this).f18)) ? ((_2504_v59).get((_this).f18)) : (_2500_v55))).length)), p1))).Intersect(_2499_v54);
      _nw415[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(p1);
      _nw415[(new BigNumber(13)).toNumber()] = _2499_v54;
      _nw415[(new BigNumber(14)).toNumber()] = ((true) ? (_2499_v54) : (_2499_v54));
      _nw415[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p1));
      _2505_v60 = _nw415;
      let _2510_v61;
      _2510_v61 = new _dafny.CodePoint('y'.codePointAt(0));
      let _2511_v62;
      _2511_v62 = _dafny.MultiSet.fromElements(_2510_v61);
      let _index405 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_2505_v60).length));
      (_2505_v60)[_index405] = _dafny.MultiSet.fromElements((((_2511_v62).contains(_2510_v61)) ? ((_2511_v62).get(_2510_v61)) : (new BigNumber((_dafny.Set.fromElements(_2510_v61, _2510_v61)).length))), new BigNumber(108), new BigNumber(273), p1);
      let _index406 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_2505_v60).length));
      (_2505_v60)[_index406] = (_dafny.MultiSet.FromArray(_2500_v55)).Union((_2499_v54).Difference(_2499_v54));
      let _hi13 = p1;
      for (let _2512_i7 = p1; _2512_i7.isLessThan(_hi13); _2512_i7 = _2512_i7.plus(_dafny.ONE)) {
        let _2513_v63;
        _2513_v63 = _dafny.Seq.of(_2502_v57);
        let _2514_v64;
        let _nw416 = new _module.C4();
        _nw416.__ctor(_2510_v61, _2512_i7, (_2512_i7).isLessThanOrEqualTo(p1), p2, _2513_v63);
        _2514_v64 = _nw416;
        let _arr2 = _this.f22;
        let _index407 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_this.f22).length));
        _arr2[_index407] = ((_dafny.ZERO).minus((_2514_v64).f30)).isLessThan(p1);
        let _arr3 = _this.f22;
        let _index408 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_this.f22).length));
        _arr3[_index408] = !((new BigNumber(((_2505_v60)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_2505_v60).length))]).cardinality())).isLessThan(_2512_i7));
        (globalState).f3 = (p1).multipliedBy(new BigNumber(905));
        let _arr4 = _this.f22;
        let _index409 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_this.f22).length));
        _arr4[_index409] = true;
      }
      let _2515_v65;
      let _2516_v66;
      let _2517_v67;
      let _2518_v68;
      let _out62;
      let _out63;
      let _out64;
      let _out65;
      let _outcollector17 = _module.__default.m0(globalState);
      _out62 = _outcollector17[0];
      _out63 = _outcollector17[1];
      _out64 = _outcollector17[2];
      _out65 = _outcollector17[3];
      _2515_v65 = _out62;
      _2516_v66 = _out63;
      _2517_v67 = _out64;
      _2518_v68 = _out65;
      let _2519_v69;
      _2519_v69 = _dafny.Map.Empty.slice().updateUnsafe((_2516_v66).isLessThanOrEqualTo(p1),_2510_v61);
      _2519_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_2510_v61);
      let _2520_v70;
      _2520_v70 = _dafny.Seq.of((_dafny.MultiSet.fromElements((_this).f18)).Difference(_module.__default.fm32(globalState)));
      _2520_v70 = _2520_v70;
      return;
    }
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f18 = false;
      this._f19 = _dafny.Seq.UnicodeFromString("");
      this._f20 = _dafny.Seq.of();
      this._f21 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f21, f19, f20, f18) {
      let _this = this;
      (_this)._f21 = f21;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f18 = f18;
      return;
    }
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return true;
    };
    fm3(p0, globalState) {
      let _this = this;
      return (_module.D0.create_DC1((_this).f21, new BigNumber(((_this).f19).length), _dafny.Seq.update((_this).f19, _module.__default.safeIndex(new BigNumber(207), new BigNumber(((_this).f19).length)), new _dafny.CodePoint('t'.codePointAt(0))), (_this).f18, new BigNumber(-625))).dtor_cf1;
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2521_i0;
      _2521_i0 = _dafny.ZERO;
      L16: {
        while ((_this).f21) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2521_i0)) {
              break L16;
            }
            _2521_i0 = (_2521_i0).plus(_dafny.ONE);
            if (p1) {
              (globalState).f15 = new BigNumber(-738);
              let _2522_v0;
              _2522_v0 = new BigNumber(-840);
              let _2523_v1;
              _2523_v1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), function (_2524_i1) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(136)), function (_2525_i2) {
                return new _dafny.CodePoint('g'.codePointAt(0));
              })),_2522_v0);
              _2523_v1 = (_2523_v1).update(p2, _2522_v0);
              let _2526_v2;
              let _nw417 = Array((new BigNumber(4)).toNumber()).fill([]);
              _2526_v2 = _nw417;
              let _2527_v3;
              let _nw418 = Array((new BigNumber(20)).toNumber()).fill(false);
              _2527_v3 = _nw418;
              let _index410 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_2526_v2).length));
              (_2526_v2)[_index410] = _2527_v3;
              let _2528_v4;
              let _nw419 = new _module.C1();
              _nw419.__ctor((_this).f21);
              _2528_v4 = _nw419;
              let _2529_v5;
              _2529_v5 = _module.D1.create_DC2(_2528_v4);
              let _pat_let_tv36 = _2528_v4;
              let _2530_v6;
              _2530_v6 = _dafny.MultiSet.fromElements((_module.D1.create_DC2((function (_pat_let70_0) {
  return function (_2531_dt__update__tmp_h0) {
    return function (_pat_let71_0) {
      return function (_2532_dt__update_hcf6_h0) {
        return _module.D1.create_DC2(_2532_dt__update_hcf6_h0);
      }(_pat_let71_0);
    }(_pat_let_tv36);
  }(_pat_let70_0);
}(_2529_v5)).dtor_cf6)).dtor_cf6);
              let _2533_v7;
              _2533_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2530_v6).cardinality()),(_dafny.ZERO).minus(_2522_v0));
              let _2534_v8;
              _2534_v8 = _dafny.MultiSet.fromElements(new BigNumber((_2533_v7).length));
              let _index411 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_2526_v2).length));
              let _rhs422 = (_dafny.ZERO).minus((new BigNumber(161)).plus(_2522_v0));
              let _rhs423 = _2527_v3;
              let _rhs424 = (_this).f18;
              let _rhs425 = (((_2534_v8).contains(_2522_v0)) ? ((_2534_v8).get(_2522_v0)) : ((new BigNumber(52)).multipliedBy(_2522_v0)));
              let _lhs364 = globalState;
              let _lhs365 = _2526_v2;
              let _lhs366 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_2526_v2).length));
              let _lhs367 = globalState;
              _lhs364.f15 = _rhs422;
              _lhs365[_lhs366] = _rhs423;
              _lhs367.f2 = _rhs424;
              _2522_v0 = _rhs425;
              let _2535_v9;
              _2535_v9 = _dafny.MultiSet.fromElements(p2);
              let _2536_v10;
              _2536_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,p2);
              let _2537_v11;
              _2537_v11 = _dafny.Set.fromElements(_2522_v0, _2522_v0);
              let _2538_v12;
              _2538_v12 = _dafny.Seq.of(_2537_v11, _2537_v11, _2537_v11);
              let _2539_v13;
              _2539_v13 = _dafny.Map.Empty.slice().updateUnsafe((_2528_v4).f18,(_2538_v12)[_module.__default.safeIndex(_2522_v0, new BigNumber((_2538_v12).length))]);
              (_this).m3((_module.__default.fm1(_2522_v0, p2, (_2528_v4).f18, globalState)).plus(new BigNumber((_module.__default.fm30(_dafny.Seq.of(_module.__default.fm32(globalState), _2535_v9), _2536_v10, globalState)).length)), _2539_v13, ((p2) ? (p2) : ((_this).f18)), globalState);
              let _2540_v14;
              _2540_v14 = _dafny.Map.Empty.slice().updateUnsafe((_2522_v0).minus(_2522_v0),(_this).f18);
              let _2541_v15;
              _2541_v15 = _dafny.Seq.of(true, (_2528_v4).f18);
              let _2542_v16;
              _2542_v16 = new _dafny.CodePoint('t'.codePointAt(0));
              let _2543_v17;
              _2543_v17 = _dafny.Seq.of(_module.__default.fm1(_2522_v0, p2, p1, globalState), _2522_v0, _2522_v0, _2522_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2528_v4).f18,_2542_v16)).length));
              let _2544_v18;
              _2544_v18 = _module.D16.create_DC38((_2534_v8).update(new BigNumber((_2543_v17).length), _module.__default.abs(new BigNumber(((_this).f19).length))));
              let _2545_v19;
              _2545_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2544_v18,((_this).fm3(_2522_v0, globalState)) && ((_this).f18));
              let _rhs426 = ((!(p2)) ? (!(p1)) : (_dafny.Seq.contains(_2541_v15, p2)));
              let _rhs427 = (_dafny.Map.Empty.slice().updateUnsafe(_2522_v0,(_this).f18)).update(_2522_v0, (_2528_v4).f18);
              let _rhs428 = (_this).f21;
              let _rhs429 = new BigNumber((_2545_v19).length);
              let _rhs430 = (_2528_v4).f18;
              let _lhs368 = globalState;
              let _lhs369 = globalState;
              let _lhs370 = globalState;
              let _lhs371 = globalState;
              _lhs368.f2 = _rhs426;
              _2540_v14 = _rhs427;
              _lhs369.f2 = _rhs428;
              _lhs370.f15 = _rhs429;
              _lhs371.f2 = _rhs430;
            } else {
              let _2546_v20;
              _2546_v20 = new BigNumber(320);
              let _2547_v21;
              _2547_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p1);
              let _2548_v22;
              _2548_v22 = _dafny.Seq.of((_this).f18, true, (_this).f18, (_this).fm3(_2546_v20, globalState), (((_2547_v21).contains(p1)) ? ((_2547_v21).get(p1)) : ((_this).f18)));
              let _2549_v24;
              _2549_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2546_v20,(_this).f21);
              (globalState).f2 = !(!(function () {
                let _coll62 = new _dafny.Map();
                for (const _compr_62 of (_2549_v24).Keys.Elements) {
                  let _2550_v23 = _compr_62;
                  if ((_2549_v24).contains(_2550_v23)) {
                    _coll62.push([(_2550_v23).multipliedBy(_2546_v20),_2546_v20]);
                  }
                }
                return _coll62;
              }()).contains(new BigNumber((_2548_v22).length)));
              (globalState).f2 = p2;
              let _2551_v25;
              _2551_v25 = new _dafny.CodePoint('m'.codePointAt(0));
              let _2552_v26;
              let _nw420 = Array((new BigNumber(13)).toNumber());
              _nw420[(_dafny.ZERO).toNumber()] = ((_this).f19)[_module.__default.safeIndex(new BigNumber((_2548_v22).length), new BigNumber(((_this).f19).length))];
              _nw420[(_dafny.ONE).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
              _nw420[(new BigNumber(3)).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(4)).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(5)).toNumber()] = _module.__default.fm46(globalState);
              _nw420[(new BigNumber(6)).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(7)).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(8)).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(9)).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(10)).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(11)).toNumber()] = _2551_v25;
              _nw420[(new BigNumber(12)).toNumber()] = _2551_v25;
              _2552_v26 = _nw420;
              let _index412 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_2552_v26).length));
              (_2552_v26)[_index412] = _2551_v25;
              let _index413 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_2552_v26).length));
              (_2552_v26)[_index413] = ((p1) ? (_2551_v25) : (_2551_v25));
              let _2553_v27;
              _2553_v27 = _dafny.Map.Empty.slice().updateUnsafe(false,_2551_v25);
              _2546_v20 = _module.__default.safeDivisionInt(_2546_v20, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2546_v20,(((_2553_v27).contains(!((_this).fm3(_2546_v20, globalState)))) ? ((_2553_v27).get(!((_this).fm3(_2546_v20, globalState)))) : (_2551_v25)))).length));
              let _2554_v28;
              let _init67 = function (_2555_i3) {
                return _module.__default.safeModuloInt(_2555_i3, new BigNumber(408));
              };
              let _nw421 = Array((new BigNumber(9)).toNumber());
              for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw421.length); _i0_67++) {
                _nw421[_i0_67] = _init67(new BigNumber(_i0_67));
              }
              _2554_v28 = _nw421;
              let _index414 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_2554_v28).length));
              (_2554_v28)[_index414] = _2546_v20;
              let _2556_v29;
              _2556_v29 = _dafny.Seq.of(_2554_v28);
              let _2557_v30;
              _2557_v30 = _dafny.Seq.of(_2546_v20);
              let _2558_v31;
              let _nw422 = new _module.C3();
              _nw422.__ctor(_2557_v30, (_this).f19, (_this).f20, (_this).f18);
              _2558_v31 = _nw422;
              let _2559_v32;
              _2559_v32 = _dafny.Set.fromElements(_2558_v31, _2558_v31, _2558_v31);
              let _2560_v33;
              _2560_v33 = _dafny.Set.fromElements(_2559_v32);
              let _index415 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_2554_v28).length));
              (_2554_v28)[_index415] = (((p2) ? (new BigNumber((_2556_v29).length)) : (_2546_v20))).plus(new BigNumber((_2560_v33).length));
            }
            let _2561_v34;
            let _init68 = ((_2562_p2) => function (_2563_i4) {
              return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),_2562_p2);
            })(p2);
            let _nw423 = Array((new BigNumber(24)).toNumber());
            for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw423.length); _i0_68++) {
              _nw423[_i0_68] = _init68(new BigNumber(_i0_68));
            }
            _2561_v34 = _nw423;
            let _2564_v35;
            _2564_v35 = _module.D3.create_DC10(_2561_v34);
            let _source24 = _2564_v35;
            if (_source24.is_DC11) {
              let _2565___mcc_h0 = (_source24).cf23;
              let _2566___mcc_h1 = (_source24).cf24;
              let _2567_cf24 = _2566___mcc_h1;
              let _2568_cf23 = _2565___mcc_h0;
              let _2569_v36;
              _2569_v36 = _dafny.Seq.of(!(_2568_cf23));
              let _2570_v37;
              _2570_v37 = new BigNumber(884);
              let _2571_v38;
              _2571_v38 = _dafny.Set.fromElements(new BigNumber(331), _2570_v37);
              let _2572_v39;
              let _nw424 = new _module.C6();
              _nw424.__ctor(_dafny.Seq.Concat(_2569_v36, _dafny.Seq.of(p2, p2)), _2570_v37, _module.__default.fm23(_2571_v38, _2570_v37, globalState), (_this).f19, (_this).f20);
              _2572_v39 = _nw424;
              _2568_cf23 = _2568_cf23;
              let _2573_v40;
              let _nw425 = new _module.C8();
              _nw425.__ctor(_2572_v39.f26, !(_2567_cf24));
              _2573_v40 = _nw425;
              let _2574_v41;
              _2574_v41 = _dafny.Seq.of(_2573_v40);
              let _2575_v42;
              _2575_v42 = _dafny.Seq.of(_2572_v39.f26);
              let _2576_v43;
              _2576_v43 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements((_2574_v41)[_module.__default.safeIndex(_2572_v39.f26, new BigNumber((_2574_v41).length))])).length), new BigNumber((_2575_v42).length));
              let _2577_v44;
              _2577_v44 = new _dafny.CodePoint('p'.codePointAt(0));
              let _2578_v45;
              _2578_v45 = _dafny.Set.fromElements((_2573_v40).f18, p1);
              (globalState).f12 = _dafny.Seq.update(_dafny.Seq.update((_this).f19, _module.__default.safeIndex((new BigNumber(807)).plus((((_2576_v43).contains(_2570_v37)) ? ((_2576_v43).get(_2570_v37)) : (_2572_v39.f26))), new BigNumber(((_this).f19).length)), _2577_v44), _module.__default.safeIndex((_2570_v37).minus(new BigNumber((_2578_v45).length)), new BigNumber((_dafny.Seq.update((_this).f19, _module.__default.safeIndex((new BigNumber(807)).plus((((_2576_v43).contains(_2570_v37)) ? ((_2576_v43).get(_2570_v37)) : (_2572_v39.f26))), new BigNumber(((_this).f19).length)), _2577_v44)).length)), _2577_v44);
              let _2579_v46;
              let _init69 = ((_2580_p2) => function (_2581_i5) {
                return _2580_p2;
              })(p2);
              let _nw426 = Array((new BigNumber(28)).toNumber());
              for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw426.length); _i0_69++) {
                _nw426[_i0_69] = _init69(new BigNumber(_i0_69));
              }
              _2579_v46 = _nw426;
              _2579_v46 = ((_2567_cf24) ? (_2579_v46) : (_2579_v46));
            } else if (_source24.is_DC12) {
              let _2582___mcc_h2 = (_source24).cf25;
              let _2583___mcc_h3 = (_source24).cf26;
              let _2584___mcc_h4 = (_source24).cf27;
              let _2585___mcc_h5 = (_source24).cf28;
              let _2586___mcc_h6 = (_source24).cf29;
              let _2587_cf29 = _2586___mcc_h6;
              let _2588_cf28 = _2585___mcc_h5;
              let _2589_cf27 = _2584___mcc_h4;
              let _2590_cf26 = _2583___mcc_h3;
              let _2591_cf25 = _2582___mcc_h2;
              let _2592_v47;
              _2592_v47 = _dafny.Seq.of(_2589_cf27);
              let _2593_v48;
              _2593_v48 = _module.D9.create_DC22(_2592_v47);
              let _2594_v49;
              _2594_v49 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat((_this).f19, (_this).f19),_2593_v48);
              _2594_v49 = (_2594_v49).update((_this).f19, _2593_v48);
              (globalState).f2 = (_this).f18;
              let _2595_v50;
              _2595_v50 = _dafny.Seq.of(_dafny.Set.fromElements(_2564_v35, _2564_v35, _2564_v35, _2564_v35));
              let _2596_v51;
              let _nw427 = new _module.C8();
              _nw427.__ctor(_2589_cf27, _2590_cf26);
              _2596_v51 = _nw427;
              let _2597_v52;
              _2597_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_2596_v51,(_dafny.ZERO).minus(_2589_cf27)),_2595_v50);
              let _2598_v53;
              _2598_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2596_v51,_2589_cf27);
              let _2599_v54;
              _2599_v54 = new _dafny.CodePoint('d'.codePointAt(0));
              let _2600_v55;
              let _nw428 = new _module.C4();
              _nw428.__ctor(_2599_v54, _2589_cf27, true, (_this).f19, (_this).f20);
              _2600_v55 = _nw428;
              let _2601_v56;
              _2601_v56 = _dafny.Set.fromElements(_2564_v35);
              _2595_v50 = (((_2597_v52).contains(_2598_v53)) ? ((_2597_v52).get(_2598_v53)) : (_dafny.Seq.update(_2595_v50, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.of(_2600_v55), _module.__default.safeIndex((_2600_v55).f30, new BigNumber((_dafny.Seq.of(_2600_v55)).length)), _2600_v55),(_2600_v55).f30)).length), new BigNumber((_2595_v50).length)), _2601_v56)));
              let _2602_v57;
              _2602_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_2589_cf27);
              let _2603_v58;
              _2603_v58 = _module.D3.create_DC13((_2600_v55).f30, (_this).f19, (_2600_v55).f30, new BigNumber(-495), ((_this).f19)[_module.__default.safeIndex((((_2602_v57).contains(true)) ? ((_2602_v57).get(true)) : ((_2600_v55).f30)), new BigNumber(((_this).f19).length))]);
              let _2604_v59;
              _2604_v59 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2603_v58).dtor_cf30);
              _2604_v59 = (_2604_v59).update((_this).f18, (_2600_v55).f30);
            } else if (_source24.is_DC13) {
              let _2605___mcc_h7 = (_source24).cf30;
              let _2606___mcc_h8 = (_source24).cf31;
              let _2607___mcc_h9 = (_source24).cf32;
              let _2608___mcc_h10 = (_source24).cf33;
              let _2609___mcc_h11 = (_source24).cf34;
              let _2610_cf34 = _2609___mcc_h11;
              let _2611_cf33 = _2608___mcc_h10;
              let _2612_cf32 = _2607___mcc_h9;
              let _2613_cf31 = _2606___mcc_h8;
              let _2614_cf30 = _2605___mcc_h7;
              let _2615_v60;
              let _nw429 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _2615_v60 = _nw429;
              let _2616_v61;
              _2616_v61 = _module.D13.create_DC33((_this).f21, new BigNumber(195), p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length));
              let _2617_v62;
              _2617_v62 = _dafny.Seq.of((_2616_v61).dtor_cf55);
              let _2618_v63;
              _2618_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2615_v60,_2617_v62);
              _2618_v63 = ((_2618_v63).update(_2615_v60, _2617_v62)).Merge(_2618_v63);
              let _2619_v64;
              _2619_v64 = _dafny.Set.fromElements((_this).f18);
              let _2620_v65;
              _2620_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_2611_cf33);
              let _2621_v66;
              _2621_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2612_cf32,new BigNumber((_2620_v65).length));
              let _2622_v67;
              let _nw430 = Array((new BigNumber(21)).toNumber());
              _nw430[(_dafny.ZERO).toNumber()] = new BigNumber(482);
              _nw430[(_dafny.ONE).toNumber()] = _2614_cf30;
              _nw430[(new BigNumber(2)).toNumber()] = _2611_cf33;
              _nw430[(new BigNumber(3)).toNumber()] = _2612_cf32;
              _nw430[(new BigNumber(4)).toNumber()] = _2612_cf32;
              _nw430[(new BigNumber(5)).toNumber()] = _2614_cf30;
              _nw430[(new BigNumber(6)).toNumber()] = new BigNumber(746);
              _nw430[(new BigNumber(7)).toNumber()] = new BigNumber((_2619_v64).length);
              _nw430[(new BigNumber(8)).toNumber()] = new BigNumber(-278);
              _nw430[(new BigNumber(9)).toNumber()] = _2612_cf32;
              _nw430[(new BigNumber(10)).toNumber()] = _2612_cf32;
              _nw430[(new BigNumber(11)).toNumber()] = _2614_cf30;
              _nw430[(new BigNumber(12)).toNumber()] = _2612_cf32;
              _nw430[(new BigNumber(13)).toNumber()] = new BigNumber((_2621_v66).length);
              _nw430[(new BigNumber(14)).toNumber()] = _2611_cf33;
              _nw430[(new BigNumber(15)).toNumber()] = new BigNumber(((_this).f19).length);
              _nw430[(new BigNumber(16)).toNumber()] = _2612_cf32;
              _nw430[(new BigNumber(17)).toNumber()] = _2612_cf32;
              _nw430[(new BigNumber(18)).toNumber()] = _2611_cf33;
              _nw430[(new BigNumber(19)).toNumber()] = new BigNumber((_2617_v62).length);
              _nw430[(new BigNumber(20)).toNumber()] = _2611_cf33;
              _2622_v67 = _nw430;
              let _2623_v68;
              let _nw431 = Array((_dafny.ONE).toNumber());
              _nw431[(_dafny.ZERO).toNumber()] = _2622_v67;
              _2623_v68 = _nw431;
              let _2624_v69;
              _2624_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2612_cf32,p2);
              let _2625_v70;
              _2625_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_2624_v69);
              let _2626_v71;
              _2626_v71 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2624_v69).Merge((((_2625_v70).contains(p2)) ? ((_2625_v70).get(p2)) : (_2624_v69)))).length),_2623_v68);
              let _rhs431 = (((_2626_v71).contains((_module.__default.fm1(_2612_cf32, (_this).f18, p1, globalState)).minus(_2614_cf30))) ? ((_2626_v71).get((_module.__default.fm1(_2612_cf32, (_this).f18, p1, globalState)).minus(_2614_cf30))) : (_2623_v68));
              let _rhs432 = ((_2612_cf32).plus((_dafny.ZERO).minus(_2612_cf32))).minus(_2611_cf33);
              let _lhs372 = globalState;
              _2623_v68 = _rhs431;
              _lhs372.f3 = _rhs432;
              let _2627_v72;
              _2627_v72 = _dafny.MultiSet.fromElements(p2);
              let _2628_v73;
              _2628_v73 = _dafny.MultiSet.fromElements(_2627_v72, _2627_v72, _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_2617_v62, _2617_v62)));
              _2628_v73 = _2628_v73;
              let _2629_v74;
              _2629_v74 = _dafny.Set.fromElements(_2620_v65);
              _2612_cf32 = new BigNumber((_2629_v74).length);
            } else if (_source24.is_DC10) {
              let _2630___mcc_h12 = (_source24).cf22;
              let _2631_cf22 = _2630___mcc_h12;
              let _2632_v75;
              _2632_v75 = new BigNumber(14);
              let _2633_v76;
              _2633_v76 = new _dafny.CodePoint('e'.codePointAt(0));
              let _2634_v77;
              let _nw432 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
              _2634_v77 = _nw432;
              let _2635_v78;
              _2635_v78 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_2632_v75, new BigNumber((_dafny.Seq.update((_this).f19, _module.__default.safeIndex(_2632_v75, new BigNumber(((_this).f19).length)), _2633_v76)).length)),(((_this).f21) ? (_2634_v77) : (_2634_v77)));
              _2635_v78 = (_2635_v78).update(_2632_v75, _2634_v77);
              let _2636_v79;
              _2636_v79 = _dafny.MultiSet.fromElements(_2632_v75);
              let _2637_v80;
              _2637_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2636_v79).cardinality()),_2632_v75);
              let _2638_v81;
              let _nw433 = new _module.C2();
              _nw433.__ctor(_2632_v75, _dafny.Seq.of(new BigNumber((_2637_v80).length)), p1);
              _2638_v81 = _nw433;
              _2638_v81 = _2638_v81;
              let _2639_v82;
              _2639_v82 = _dafny.Seq.of((_this).f19, (_this).f19);
              let _2640_v83;
              _2640_v83 = _dafny.Set.fromElements(p2);
              let _2641_v84;
              let _nw434 = new _module.C4();
              _nw434.__ctor(_2633_v76, _2632_v75, p2, (_2639_v82)[_module.__default.safeIndex(new BigNumber((_2640_v83).length), new BigNumber((_2639_v82).length))], (_this).f20);
              _2641_v84 = _nw434;
              let _2642_v85;
              _2642_v85 = _dafny.Map.Empty.slice().updateUnsafe(false,_2641_v84);
              let _2643_v86;
              _2643_v86 = _dafny.Seq.of(_2641_v84);
              let _2644_v87;
              _2644_v87 = _dafny.Seq.of((_this).f18);
              let _2645_v88;
              let _nw435 = Array((new BigNumber(24)).toNumber());
              _nw435[(_dafny.ZERO).toNumber()] = _2641_v84;
              _nw435[(_dafny.ONE).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(2)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(3)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(4)).toNumber()] = (((_2642_v85).contains(false)) ? ((_2642_v85).get(false)) : (_2641_v84));
              _nw435[(new BigNumber(5)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(6)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(7)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(8)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(9)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(10)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(11)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(12)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(13)).toNumber()] = (_2643_v86)[_module.__default.safeIndex(_2632_v75, new BigNumber((_2643_v86).length))];
              _nw435[(new BigNumber(14)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(15)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(16)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(17)).toNumber()] = (((_this).f18) ? (_2641_v84) : (_2641_v84));
              _nw435[(new BigNumber(18)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(19)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(20)).toNumber()] = (((_this).fm3(new BigNumber((_2644_v87).length), globalState)) ? (_2641_v84) : (_2641_v84));
              _nw435[(new BigNumber(21)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(22)).toNumber()] = _2641_v84;
              _nw435[(new BigNumber(23)).toNumber()] = _2641_v84;
              _2645_v88 = _nw435;
              let _index416 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_2645_v88).length));
              (_2645_v88)[_index416] = _2641_v84;
              let _index417 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_2645_v88).length));
              (_2645_v88)[_index417] = _2641_v84;
              (_2641_v84).m1(_dafny.Seq.UnicodeFromString("atpvbkx"), (_dafny.ZERO).minus(new BigNumber(-305)), _dafny.Seq.UnicodeFromString("ib"), globalState);
            } else {
              let _2646___mcc_h13 = (_source24).cf35;
              let _2647_cf35 = _2646___mcc_h13;
              (globalState).f2 = p1;
              (globalState).f2 = (_this).f21;
              let _2648_v89;
              _2648_v89 = new BigNumber(638);
              (globalState).f0 = ((!(p1)) ? ((_dafny.ZERO).minus(_2648_v89)) : ((_dafny.ZERO).minus((_2648_v89).minus(_2648_v89))));
              let _2649_v90;
              _2649_v90 = _dafny.Set.fromElements(new BigNumber(((_this).f19).length), _2648_v89, _2648_v89);
              let _2650_v91;
              _2650_v91 = _dafny.Seq.of(new BigNumber((_2649_v90).length));
              let _2651_v92;
              _2651_v92 = _dafny.Set.fromElements((_this).f21);
              let _2652_v93;
              _2652_v93 = _module.D10.create_DC24(_2651_v92);
              let _2653_v94;
              _2653_v94 = _module.D13.create_DC33(false, _2648_v89, !((_this).f21), _2648_v89);
              let _2654_v95;
              let _nw436 = new _module.C3();
              _nw436.__ctor(_2650_v91, _dafny.Seq.UnicodeFromString("akrqthbvy"), _module.__default.fm51(p2, _2648_v89, p1, _2652_v93, globalState), (_2653_v94).dtor_cf57);
              _2654_v95 = _nw436;
            }
            let _2655_v96;
            _2655_v96 = _dafny.MultiSet.fromElements(new BigNumber(-876));
            let _2656_v97;
            _2656_v97 = new BigNumber(259);
            (globalState).f0 = ((((_2655_v96).contains(_2656_v97)) ? ((_2655_v96).get(_2656_v97)) : (_2656_v97))).multipliedBy((_2656_v97).plus(_2656_v97));
            let _2657_v98;
            _2657_v98 = _dafny.Set.fromElements(!(p2), (_this).f18);
            let _2658_v99;
            _2658_v99 = new _dafny.CodePoint('b'.codePointAt(0));
            let _2659_v100;
            _2659_v100 = _dafny.Seq.of((_this).f18, p2);
            let _2660_v101;
            _2660_v101 = _dafny.Map.Empty.slice().updateUnsafe(_2659_v100,new BigNumber(-758));
            let _2661_v102;
            _2661_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f19);
            let _2662_v103;
            _2662_v103 = _dafny.Map.Empty.slice().updateUnsafe(_2656_v97,new BigNumber(((((_2661_v102).contains((_this).f18)) ? ((_2661_v102).get((_this).f18)) : ((_this).f19))).length));
            let _2663_v104;
            _2663_v104 = _dafny.Set.fromElements(new BigNumber(-682));
            let _2664_v105;
            _2664_v105 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_2656_v97);
            let _2665_v106;
            _2665_v106 = _dafny.MultiSet.fromElements(_module.__default.fm23(_2663_v104, new BigNumber((_2664_v105).length), globalState), true, p2, p2, true);
            let _2666_v108;
            _2666_v108 = _dafny.MultiSet.fromElements(_2658_v99);
            let _2667_v109;
            _2667_v109 = _module.D0.create_DC0(false);
            let _2668_v110;
            _2668_v110 = _dafny.Seq.of(new BigNumber(-670), _2656_v97);
            let _2669_v111;
            _2669_v111 = _dafny.Seq.of(_2656_v97, (_2668_v110)[_module.__default.safeIndex(_2656_v97, new BigNumber((_2668_v110).length))], new BigNumber(((_this).f19).length), (_dafny.ZERO).minus(_2656_v97));
            let _2670_v112;
            let _nw437 = new _module.C2();
            _nw437.__ctor(_2656_v97, _2668_v110, p2);
            _2670_v112 = _nw437;
            let _2671_v113;
            _2671_v113 = _dafny.Map.Empty.slice().updateUnsafe((_2667_v109).dtor_cf0,_2670_v112);
            let _2672_v114;
            let _nw438 = Array((new BigNumber(23)).toNumber());
            _nw438[(_dafny.ZERO).toNumber()] = _2656_v97;
            _nw438[(_dafny.ONE).toNumber()] = _2656_v97;
            _nw438[(new BigNumber(2)).toNumber()] = (new BigNumber((_2657_v98).length)).plus(new BigNumber((_dafny.MultiSet.fromElements(_2658_v99)).cardinality()));
            _nw438[(new BigNumber(3)).toNumber()] = (((_2660_v101).contains(_dafny.Seq.of(p1))) ? ((_2660_v101).get(_dafny.Seq.of(p1))) : (_2656_v97));
            _nw438[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_2656_v97, _2656_v97);
            _nw438[(new BigNumber(5)).toNumber()] = _2656_v97;
            _nw438[(new BigNumber(6)).toNumber()] = _2656_v97;
            _nw438[(new BigNumber(7)).toNumber()] = _2656_v97;
            _nw438[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(888)), ((_2673_v97, _2674_p1) => function (_2675_i6) {
              return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2673_v97, new BigNumber(111)),_2674_p1)).length);
            })(_2656_v97, p1))).length);
            _nw438[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(119), _2656_v97);
            _nw438[(new BigNumber(10)).toNumber()] = new BigNumber(((_this).f19).length);
            _nw438[(new BigNumber(11)).toNumber()] = (((_2662_v103).contains((_dafny.ZERO).minus(_2656_v97))) ? ((_2662_v103).get((_dafny.ZERO).minus(_2656_v97))) : (_2656_v97));
            _nw438[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2665_v106).cardinality()));
            _nw438[(new BigNumber(13)).toNumber()] = new BigNumber(((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((function () {
              let _coll63 = new _dafny.Map();
              for (const _compr_63 of (_2666_v108).Elements) {
                let _2676_v107 = _compr_63;
                if ((_2666_v108).contains(_2676_v107)) {
                  _coll63.push([_2676_v107,(_this).f21]);
                }
              }
              return _coll63;
            }()).length)), new BigNumber((_2671_v113).length), new BigNumber(((_this).f19).length))).Union(_2663_v104)).length);
            _nw438[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_2670_v112.f33, _2670_v112.f33));
            _nw438[(new BigNumber(15)).toNumber()] = _2670_v112.f33;
            _nw438[(new BigNumber(16)).toNumber()] = new BigNumber(366);
            _nw438[(new BigNumber(17)).toNumber()] = _2670_v112.f33;
            _nw438[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt(_2670_v112.f33, _2656_v97);
            _nw438[(new BigNumber(19)).toNumber()] = _2670_v112.f33;
            _nw438[(new BigNumber(20)).toNumber()] = _2656_v97;
            _nw438[(new BigNumber(21)).toNumber()] = (new BigNumber((_2659_v100).length)).plus(_2656_v97);
            _nw438[(new BigNumber(22)).toNumber()] = new BigNumber(291);
            _2672_v114 = _nw438;
            let _index418 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2672_v114).length));
            (_2672_v114)[_index418] = (_2656_v97).plus(_2656_v97);
            let _2677_v115;
            _2677_v115 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p1);
            let _2678_v116;
            _2678_v116 = _dafny.Map.Empty.slice().updateUnsafe(_2677_v115,_2656_v97);
            let _index419 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2672_v114).length));
            (_2672_v114)[_index419] = (((_2678_v116).contains((_2677_v115).Merge((_2677_v115).update(_module.__default.fm23(_2663_v104, _2656_v97, globalState), (_this).f18)))) ? ((_2678_v116).get((_2677_v115).Merge((_2677_v115).update(_module.__default.fm23(_2663_v104, _2656_v97, globalState), (_this).f18)))) : (_2670_v112.f33));
          }
        }
      }
      let _2679_v117;
      _2679_v117 = _module.D19.create_DC49(!((_this).f21));
      let _2680_v118;
      let _nw439 = new _module.C1();
      _nw439.__ctor((_2679_v117).dtor_cf83);
      _2680_v118 = _nw439;
      let _2681_v119;
      _2681_v119 = _dafny.Set.fromElements(_2680_v118);
      let _2682_v120;
      _2682_v120 = new BigNumber(-915);
      (globalState).f0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_2681_v119).length), _2682_v120));
      if ((_2680_v118).f18) {
        let _2683_v122;
        let _init70 = ((_2684_v120) => function (_2685_i7) {
          return (_2685_i7).minus(new BigNumber((function () {
            let _coll64 = new _dafny.Set();
            for (const _compr_64 of _dafny.IntegerRange(new BigNumber(980), new BigNumber(-303))) {
              let _2686_v121 = _compr_64;
              if (((new BigNumber(980)).isLessThanOrEqualTo(_2686_v121)) && ((_2686_v121).isLessThan(new BigNumber(-303)))) {
                _coll64.add(_module.__default.safeDivisionInt(_2686_v121, _2684_v120));
              }
            }
            return _coll64;
          }()).length));
        })(_2682_v120);
        let _nw440 = Array((new BigNumber(18)).toNumber());
        for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw440.length); _i0_70++) {
          _nw440[_i0_70] = _init70(new BigNumber(_i0_70));
        }
        _2683_v122 = _nw440;
        let _index420 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length));
        (_2683_v122)[_index420] = _2682_v120;
        let _2687_v123;
        _2687_v123 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2688_v124;
        _2688_v124 = _dafny.Map.Empty.slice().updateUnsafe(_2687_v123,(_module.__default.fm63(globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_2680_v118).f18,_2687_v123)));
        let _2689_v125;
        _2689_v125 = _module.D16.create_DC39((_2680_v118).f18, new BigNumber(553));
        let _index421 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length));
        let _rhs433 = _2682_v120;
        let _rhs434 = _2682_v120;
        let _rhs435 = (_2689_v125).dtor_cf67;
        let _rhs436 = function () {
          let _coll65 = new _dafny.Map();
          for (const _compr_65 of (_dafny.Map.Empty.slice().updateUnsafe(_2687_v123,_2687_v123)).Keys.Elements) {
            let _2690_v126 = _compr_65;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_2687_v123,_2687_v123)).contains(_2690_v126)) {
              _coll65.push([_2690_v126,_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_2687_v123)]);
            }
          }
          return _coll65;
        }();
        let _lhs373 = _2683_v122;
        let _lhs374 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length));
        let _lhs375 = globalState;
        let _lhs376 = globalState;
        _lhs373[_lhs374] = _rhs433;
        _lhs375.f15 = _rhs434;
        _lhs376.f2 = _rhs435;
        _2688_v124 = _rhs436;
        let _index422 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length));
        (_2683_v122)[_index422] = new BigNumber(703);
        let _2691_v127;
        _2691_v127 = _module.D6.create_DC18();
        let _source25 = _2691_v127;
        if (_source25.is_DC18) {
          let _2692_v128;
          let _nw441 = Array((new BigNumber(23)).toNumber()).fill([]);
          _2692_v128 = _nw441;
          let _index423 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_2692_v128).length));
          (_2692_v128)[_index423] = _2683_v122;
          let _2693_v129;
          _2693_v129 = _dafny.Seq.of((_2680_v118).f18, (_this).f18);
          let _2694_v130;
          _2694_v130 = _dafny.Seq.of(_2693_v129);
          let _2695_v131;
          _2695_v131 = _dafny.Map.Empty.slice().updateUnsafe(_2694_v130,_2683_v122);
          let _2696_v132;
          _2696_v132 = _2683_v122;
          let _index424 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_2692_v128).length));
          (_2692_v128)[_index424] = (((_2695_v131).contains(_2694_v130)) ? ((_2695_v131).get(_2694_v130)) : ((_2696_v132)));
          let _2697_v133;
          let _nw442 = Array((new BigNumber(12)).toNumber());
          _2697_v133 = _nw442;
          let _2698_v134;
          _2698_v134 = _dafny.Seq.of(_2697_v133, _2697_v133);
          let _2699_v135;
          _2699_v135 = _dafny.Set.fromElements(new BigNumber((_2698_v134).length));
          let _2700_v136;
          _2700_v136 = _dafny.Map.Empty.slice().updateUnsafe(false,(_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))]);
          let _index425 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length));
          let _rhs437 = false;
          let _rhs438 = (_2699_v135).Intersect((_dafny.Set.fromElements(_2682_v120)).Union(_dafny.Set.fromElements(new BigNumber((_2700_v136).length))));
          let _rhs439 = p2;
          let _rhs440 = (_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))];
          let _lhs377 = globalState;
          let _lhs378 = globalState;
          let _lhs379 = _2683_v122;
          let _lhs380 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length));
          _lhs377.f2 = _rhs437;
          _2699_v135 = _rhs438;
          _lhs378.f2 = _rhs439;
          _lhs379[_lhs380] = _rhs440;
          let _2701_v137;
          _2701_v137 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(p2), _module.__default.safeIndex(_2682_v120, new BigNumber((_dafny.Seq.of(p2)).length)), p1)).length),(_this).f21);
          _2701_v137 = _2701_v137;
          let _2702_v138;
          _2702_v138 = _dafny.Seq.of(_2682_v120, new BigNumber(496));
          let _2703_v139;
          _2703_v139 = _dafny.Map.Empty.slice().updateUnsafe((_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))],_2682_v120);
          let _2704_v140;
          _2704_v140 = _dafny.MultiSet.fromElements(true);
          let _2705_v141;
          _2705_v141 = _dafny.Map.Empty.slice().updateUnsafe(_2704_v140,(_2680_v118).f18);
          let _2706_v142;
          _2706_v142 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2705_v141).length),_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(p1),(_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))]),_2687_v123)).length)));
          let _2707_v143;
          let _nw443 = new _module.C5();
          _nw443.__ctor((_this).f19, (_this).f19, (_2680_v118).f18);
          _2707_v143 = _nw443;
          let _2708_v144;
          _2708_v144 = _dafny.MultiSet.fromElements((_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))]);
          let _2709_v145;
          _2709_v145 = _module.D24.create_DC64(_module.__default.fm23((((_2706_v142).contains(_2682_v120)) ? ((_2706_v142).get(_2682_v120)) : (_2699_v135)), new BigNumber((_2693_v129).length), globalState), (_this).f21, _2707_v143, _2682_v120, new BigNumber((_2708_v144).cardinality()));
          let _2710_v146;
          let _nw444 = new _module.C3();
          _nw444.__ctor(_2702_v138, _module.__default.fm48(globalState), _dafny.Seq.Concat((_this).f20, _dafny.Seq.of(_2703_v139)), (_2709_v145).dtor_cf106);
          _2710_v146 = _nw444;
        } else {
          let _2711___mcc_h14 = (_source25).cf38;
          let _2712_cf38 = _2711___mcc_h14;
          let _2713_v147;
          _2713_v147 = _module.D22.create_DC58((_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))]);
          _2713_v147 = _2713_v147;
          let _2714_v148;
          let _nw445 = Array((new BigNumber(24)).toNumber());
          _2714_v148 = _nw445;
          let _2715_v149;
          let _nw446 = new _module.C5();
          _nw446.__ctor(_dafny.Seq.update((_this).f19, _module.__default.safeIndex((_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))], new BigNumber(((_this).f19).length)), _2687_v123), _dafny.Seq.Create(_module.__default.abs(new BigNumber(957)), function (_2716_i8) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }), (_2680_v118).f18);
          _2715_v149 = _nw446;
          let _index426 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_2714_v148).length));
          (_2714_v148)[_index426] = (_module.D24.create_DC64(!((_this).f18), (_2680_v118).f18, _2715_v149, (_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))], (_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))])).dtor_cf107;
          let _index427 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_2714_v148).length));
          (_2714_v148)[_index427] = _2715_v149;
          let _2717_v150;
          let _nw447 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2717_v150 = _nw447;
          let _2718_v151;
          _2718_v151 = _dafny.Set.fromElements((_2714_v148)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_2714_v148).length))]);
          let _2719_v152;
          _2719_v152 = _dafny.Seq.of((_2714_v148)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_2714_v148).length))]);
          let _2720_v153;
          _2720_v153 = _dafny.MultiSet.fromElements(_2718_v151, _dafny.Set.fromElements((_2719_v152)[_module.__default.safeIndex((_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))], new BigNumber((_2719_v152).length))]), _2718_v151);
          let _index428 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_2717_v150).length));
          (_2717_v150)[_index428] = _2720_v153;
          let _2721_v154;
          _2721_v154 = _dafny.MultiSet.fromElements(p1);
          let _2722_v155;
          _2722_v155 = _dafny.Map.Empty.slice().updateUnsafe((((_2721_v154).contains(!(false))) ? ((_2721_v154).get(!(false))) : ((_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))])),_2718_v151);
          let _2723_v156;
          _2723_v156 = _dafny.Seq.of(_2718_v151, _2718_v151, _2718_v151, _dafny.Set.fromElements((_2714_v148)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_2714_v148).length))], _2715_v149));
          let _2724_v157;
          _2724_v157 = _module.D24.create_DC64(p2, !((_this).f18), _2715_v149, (_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))], (_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))]);
          let _2725_v158;
          _2725_v158 = _dafny.Seq.of(_dafny.Seq.of((((_2722_v155).contains(_2682_v120)) ? ((_2722_v155).get(_2682_v120)) : (_2718_v151))), _dafny.Seq.Concat(_2723_v156, _2723_v156), _dafny.Seq.of(_2718_v151, _dafny.Set.fromElements((_2724_v157).dtor_cf107, _2715_v149)));
          let _index429 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_2717_v150).length));
          (_2717_v150)[_index429] = _dafny.MultiSet.FromArray((_2725_v158)[_module.__default.safeIndex(_2682_v120, new BigNumber((_2725_v158).length))]);
          _2682_v120 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat((_this).f19, _dafny.Seq.UnicodeFromString("csklyy"))).length));
        }
        (globalState).f2 = (_this).fm3(((!((_this).f18)) ? ((_2683_v122)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_2683_v122).length))]) : (_2682_v120)), globalState);
        let _2726_v159;
        _2726_v159 = _dafny.MultiSet.fromElements(!((_this).f18));
        let _rhs441 = !((_this).f21);
        let _rhs442 = _2726_v159;
        let _lhs381 = globalState;
        _lhs381.f2 = _rhs441;
        _2726_v159 = _rhs442;
      } else {
        let _2727_v160;
        let _nw448 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _2727_v160 = _nw448;
        let _index430 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length));
        (_2727_v160)[_index430] = _2682_v120;
        let _index431 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length));
        (_2727_v160)[_index431] = _2682_v120;
        let _2728_v161;
        _2728_v161 = _module.D10.create_DC25(p1);
        (globalState).f2 = (_2728_v161).dtor_cf48;
        let _2729_v162;
        _2729_v162 = _dafny.Seq.of((_this).f21);
        let _2730_v163;
        _2730_v163 = _module.D21.create_DC56(_dafny.Seq.IsPrefixOf(_2729_v162, _2729_v162), _2682_v120, (new BigNumber(778)).multipliedBy((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))]));
        let _source26 = _2730_v163;
        if (_source26.is_DC56) {
          let _2731___mcc_h15 = (_source26).cf92;
          let _2732___mcc_h16 = (_source26).cf93;
          let _2733___mcc_h17 = (_source26).cf94;
          let _2734_cf94 = _2733___mcc_h17;
          let _2735_cf93 = _2732___mcc_h16;
          let _2736_cf92 = _2731___mcc_h15;
          let _2737_v164;
          _2737_v164 = _module.D2.create_DC7((_dafny.ZERO).minus(_2682_v120), (_this).f21, _2735_cf93);
          (globalState).f10 = _module.__default.fm40(function (_pat_let72_0) {
            return function (_2738_dt__update__tmp_h1) {
              return function (_pat_let73_0) {
                return function (_2739_dt__update_hcf17_h0) {
                  return _module.D2.create_DC7((_2738_dt__update__tmp_h1).dtor_cf15, (_2738_dt__update__tmp_h1).dtor_cf16, _2739_dt__update_hcf17_h0);
                }(_pat_let73_0);
              }(new BigNumber(555));
            }(_pat_let72_0);
          }(_2737_v164), globalState);
          let _2740_v165;
          _2740_v165 = _dafny.MultiSet.fromElements((_2680_v118).f18);
          let _2741_v166;
          _2741_v166 = _dafny.MultiSet.fromElements((_2740_v165).update((_2680_v118).f18, _module.__default.abs((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))])));
          let _2742_v167;
          _2742_v167 = _module.D19.create_DC47(_2741_v166);
          let _2743_v168;
          let _nw449 = new _module.C1();
          _nw449.__ctor((_this).f18);
          _2743_v168 = _nw449;
          let _2744_v169;
          _2744_v169 = _dafny.Map.Empty.slice().updateUnsafe(_2742_v167,_2743_v168);
          let _2745_v170;
          _2745_v170 = _dafny.Seq.of(_2735_cf93, _2682_v120, _2735_cf93, (new BigNumber((_2744_v169).length)).multipliedBy(_2734_cf94));
          (globalState).f3 = (_2745_v170)[_module.__default.safeIndex(_module.__default.fm1(_2734_cf94, p2, (_2680_v118).f18, globalState), new BigNumber((_2745_v170).length))];
          let _2746_v171;
          let _nw450 = new _module.C7();
          _nw450.__ctor(_2734_cf94, (_2680_v118).f18);
          _2746_v171 = _nw450;
          _2746_v171 = _2746_v171;
          let _2747_v172;
          let _nw451 = new _module.C6();
          _nw451.__ctor(_2729_v162, _2734_cf94, (_2680_v118).f18, (_this).f19, (_this).f20);
          _2747_v172 = _nw451;
          let _2748_v173;
          _2748_v173 = _dafny.Seq.of(_2747_v172);
          let _2749_v174;
          let _nw452 = new _module.C6();
          _nw452.__ctor(_2729_v162, new BigNumber((_2748_v173).length), false, (_this).f19, (_this).f20);
          _2749_v174 = _nw452;
          let _2750_v175;
          _2750_v175 = _dafny.Map.Empty.slice().updateUnsafe(_2747_v172,_2735_cf93);
          let _2751_v176;
          _2751_v176 = new _dafny.CodePoint('u'.codePointAt(0));
          let _2752_v177;
          _2752_v177 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2750_v175).length),_2751_v176);
          let _2753_v178;
          let _nw453 = Array((new BigNumber(23)).toNumber());
          _nw453[(_dafny.ZERO).toNumber()] = true;
          _nw453[(_dafny.ONE).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(2)).toNumber()] = p2;
          _nw453[(new BigNumber(3)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(4)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(5)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(6)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(7)).toNumber()] = false;
          _nw453[(new BigNumber(8)).toNumber()] = (_2734_cf94).isLessThanOrEqualTo(_2682_v120);
          _nw453[(new BigNumber(9)).toNumber()] = ((_this).f21) === ((_this).f18);
          _nw453[(new BigNumber(10)).toNumber()] = _2736_cf92;
          _nw453[(new BigNumber(11)).toNumber()] = _2736_cf92;
          _nw453[(new BigNumber(12)).toNumber()] = _2736_cf92;
          _nw453[(new BigNumber(13)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(14)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(15)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(16)).toNumber()] = (_2729_v162)[_module.__default.safeIndex(_2682_v120, new BigNumber((_2729_v162).length))];
          _nw453[(new BigNumber(17)).toNumber()] = _2736_cf92;
          _nw453[(new BigNumber(18)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(19)).toNumber()] = (_2680_v118).f18;
          _nw453[(new BigNumber(20)).toNumber()] = (_this).f18;
          _nw453[(new BigNumber(21)).toNumber()] = !((_this).f18);
          _nw453[(new BigNumber(22)).toNumber()] = !(_2752_v177).equals(_2752_v177);
          _2753_v178 = _nw453;
          let _index432 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_2753_v178).length));
          (_2753_v178)[_index432] = (true) || (p1);
          let _index433 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_2753_v178).length));
          (_2753_v178)[_index433] = (_this).f21;
        } else {
          let _2754___mcc_h18 = (_source26).cf91;
          let _2755_cf91 = _2754___mcc_h18;
          let _2756_v179;
          let _init71 = function (_2757_i9) {
            return (_this).f21;
          };
          let _nw454 = Array((new BigNumber(28)).toNumber());
          for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw454.length); _i0_71++) {
            _nw454[_i0_71] = _init71(new BigNumber(_i0_71));
          }
          _2756_v179 = _nw454;
          let _2758_v180;
          _2758_v180 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_2756_v179);
          let _2759_v181;
          _2759_v181 = _dafny.Map.Empty.slice().updateUnsafe((_2758_v180).update((_this).f18, _2756_v179),(_2680_v118).f18);
          (globalState).f2 = !((((_2759_v181).contains(_dafny.Map.Empty.slice().updateUnsafe((_2680_v118).f18,_2756_v179))) ? ((_2759_v181).get(_dafny.Map.Empty.slice().updateUnsafe((_2680_v118).f18,_2756_v179))) : (p2)));
          (globalState).f2 = (_2680_v118).f18;
          let _index434 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length));
          (_2727_v160)[_index434] = (_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))];
          let _2760_v182;
          _2760_v182 = new _dafny.CodePoint('s'.codePointAt(0));
          _2760_v182 = _2760_v182;
        }
        (globalState).f2 = (_this).f21;
        if (true) {
          let _2761_v183;
          _2761_v183 = new _dafny.CodePoint('r'.codePointAt(0));
          let _2762_v184;
          _2762_v184 = _dafny.MultiSet.fromElements(new BigNumber(722));
          let _2763_v185;
          let _nw455 = new _module.C4();
          _nw455.__ctor(_2761_v183, (((_2762_v184).contains((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))])) ? ((_2762_v184).get((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))])) : ((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))])), (_this).f18, _dafny.Seq.UnicodeFromString("jgqnwe"), _dafny.Seq.update(_dafny.Seq.Concat((_this).f20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), ((_2764_v160) => function (_2765_i10) {
            return _dafny.Map.Empty.slice().updateUnsafe((_2764_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2764_v160).length))],new BigNumber(-459));
          })(_2727_v160))), _module.__default.safeIndex((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))], new BigNumber((_dafny.Seq.Concat((_this).f20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), ((_2766_v160) => function (_2767_i10) {
            return _dafny.Map.Empty.slice().updateUnsafe((_2766_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2766_v160).length))],new BigNumber(-459));
          })(_2727_v160)))).length)), _dafny.Map.Empty.slice().updateUnsafe((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))],new BigNumber((_dafny.Seq.update((_this).f19, _module.__default.safeIndex(_2682_v120, new BigNumber(((_this).f19).length)), new _dafny.CodePoint('c'.codePointAt(0)))).length))));
          _2763_v185 = _nw455;
          _2763_v185 = _2763_v185;
          let _2768_v186;
          let _nw456 = Array((new BigNumber(26)).toNumber()).fill(false);
          _2768_v186 = _nw456;
          let _index435 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_2768_v186).length));
          (_2768_v186)[_index435] = (new BigNumber((_dafny.Seq.of((_2680_v118).f18)).length)).isLessThanOrEqualTo((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))]);
          let _index436 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_2768_v186).length));
          let _rhs443 = !((_this).f21) || (p2);
          let _rhs444 = (_2680_v118).f18;
          let _lhs382 = globalState;
          let _lhs383 = _2768_v186;
          let _lhs384 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_2768_v186).length));
          _lhs382.f2 = _rhs443;
          _lhs383[_lhs384] = _rhs444;
          (globalState).f2 = (_this).f21;
          let _2769_v187;
          _2769_v187 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2680_v118).f18);
          let _2770_v188;
          _2770_v188 = _dafny.Map.Empty.slice().updateUnsafe(_2768_v186,_2769_v187);
          (globalState).f0 = new BigNumber((((p1) ? (_2770_v188) : (((_2770_v188).update(_2768_v186, _2769_v187)).Merge(_2770_v188)))).length);
          let _2771_v189;
          _2771_v189 = _module.D9.create_DC23(p1, p1, (_2763_v185).f30);
          let _2772_v190;
          _2772_v190 = _dafny.MultiSet.fromElements(_2771_v189, _2771_v189);
          _2772_v190 = _2772_v190;
        } else {
          let _2773_v191;
          _2773_v191 = new _dafny.CodePoint('k'.codePointAt(0));
          let _2774_v192;
          _2774_v192 = _dafny.Map.Empty.slice().updateUnsafe((_2680_v118).f18,p2);
          let _2775_v193;
          let _init72 = function (_2776_i11) {
            return (_this).f21;
          };
          let _nw457 = Array((new BigNumber(23)).toNumber());
          for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw457.length); _i0_72++) {
            _nw457[_i0_72] = _init72(new BigNumber(_i0_72));
          }
          _2775_v193 = _nw457;
          let _2777_v194;
          _2777_v194 = _dafny.Map.Empty.slice().updateUnsafe(_2775_v193,_dafny.Seq.Create(_module.__default.abs(new BigNumber(823)), ((_2778_v191) => function (_2779_i12) {
            return _2778_v191;
          })(_2773_v191)));
          let _2780_v195;
          _2780_v195 = _module.D0.create_DC1((((_2774_v192).contains(p2)) ? ((_2774_v192).get(p2)) : ((_this).f18)), (_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))], (((_2777_v194).contains(_2775_v193)) ? ((_2777_v194).get(_2775_v193)) : ((_this).f19)), (_2680_v118).f18, _2682_v120);
          let _2781_v196;
          _2781_v196 = _dafny.Map.Empty.slice().updateUnsafe((_2780_v195).dtor_cf1,(_this).f21);
          let _2782_v197;
          _2782_v197 = _dafny.Map.Empty.slice().updateUnsafe(_2682_v120,(_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))]);
          let _2783_v198;
          let _nw458 = new _module.C4();
          _nw458.__ctor(_2773_v191, _2682_v120, (((_2774_v192).contains((_2680_v118).f18)) ? ((_2774_v192).get((_2680_v118).f18)) : ((_2680_v118).f18)), (_this).f19, _dafny.Seq.update(_dafny.Seq.of(_2782_v197), _module.__default.safeIndex(new BigNumber((_2781_v196).length), new BigNumber((_dafny.Seq.of(_2782_v197)).length)), _dafny.Map.Empty.slice().updateUnsafe(_2682_v120,(_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))])));
          _2783_v198 = _nw458;
          let _2784_v199;
          _2784_v199 = _dafny.Map.Empty.slice().updateUnsafe((((_2774_v192).contains(p1)) ? ((_2774_v192).get(p1)) : ((_2680_v118).f18)),new BigNumber((_2729_v162).length));
          let _2785_v200;
          _2785_v200 = _dafny.Map.Empty.slice().updateUnsafe(_2773_v191,_2784_v199);
          let _index437 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length));
          (_2727_v160)[_index437] = _module.__default.fm1(new BigNumber(749), (_2680_v118).f18, !((_dafny.Map.Empty.slice().updateUnsafe(_2773_v191,_2784_v199)).update(new _dafny.CodePoint('b'.codePointAt(0)), _2784_v199)).equals(_2785_v200), globalState);
          let _2786_v201;
          _2786_v201 = _dafny.Set.fromElements((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))]);
          let _2787_v202;
          _2787_v202 = _dafny.MultiSet.fromElements(_2786_v201);
          let _2788_v203;
          _2788_v203 = _dafny.MultiSet.fromElements(p2);
          let _2789_v204;
          _2789_v204 = _dafny.Map.Empty.slice().updateUnsafe((_2680_v118).f18,_2773_v191);
          let _2790_v205;
          _2790_v205 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2789_v204);
          let _2791_v206;
          _2791_v206 = _module.D25.create_DC66(_2790_v205);
          let _2792_v207;
          _2792_v207 = _dafny.Map.Empty.slice().updateUnsafe((_2787_v202).update(_dafny.Set.fromElements(_2682_v120), _module.__default.abs((((_2788_v203).contains((_2680_v118).f18)) ? ((_2788_v203).get((_2680_v118).f18)) : (_2682_v120)))),((p2) ? ((_2791_v206).dtor_cf115) : (_2790_v205)));
          _2792_v207 = (_2792_v207).update((((_this).f18) ? (_dafny.MultiSet.fromElements(_2786_v201)) : (_2787_v202)), _2790_v205);
          let _2793_v209;
          _2793_v209 = _dafny.Seq.of(_2682_v120, (_dafny.ZERO).minus(new BigNumber((function () {
            let _coll66 = new _dafny.Map();
            for (const _compr_66 of _dafny.IntegerRange(new BigNumber(448), new BigNumber(295))) {
              let _2794_v208 = _compr_66;
              if (((new BigNumber(448)).isLessThanOrEqualTo(_2794_v208)) && ((_2794_v208).isLessThan(new BigNumber(295)))) {
                _coll66.push([(_2794_v208).multipliedBy(new BigNumber(168)),new BigNumber(-990)]);
              }
            }
            return _coll66;
          }()).length)));
          let _2795_v210;
          let _nw459 = new _module.C3();
          _nw459.__ctor(_2793_v209, _dafny.Seq.UnicodeFromString("cedehmcug"), (_this).f20, (_this).f21);
          _2795_v210 = _nw459;
          let _2796_v211;
          let _nw460 = new _module.C1();
          _nw460.__ctor(p2);
          _2796_v211 = _nw460;
          let _2797_v212;
          _2797_v212 = _dafny.Map.Empty.slice().updateUnsafe(_2795_v210,_2796_v211);
          _2797_v212 = (_2797_v212).update(_2795_v210, _2796_v211);
          _2773_v191 = ((_this).f19)[_module.__default.safeIndex((_2727_v160)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_2727_v160).length))], new BigNumber(((_this).f19).length))];
        }
      }
      (globalState).f0 = _2682_v120;
      (globalState).f2 = (_this).f21;
      (globalState).f2 = !(!(_2682_v120).isEqualTo(_module.__default.fm1(_2682_v120, (_this).f18, p1, globalState)));
      r0 = _2682_v120;
      return r0;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      (globalState).f2 = (_this).f21;
      let _2798_v0;
      let _nw461 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _2798_v0 = _nw461;
      let _2799_v1;
      _2799_v1 = new _dafny.CodePoint('q'.codePointAt(0));
      let _index438 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_2798_v0).length));
      (_2798_v0)[_index438] = _2799_v1;
      let _index439 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_2798_v0).length));
      (_2798_v0)[_index439] = _2799_v1;
      let _2800_v2;
      let _nw462 = new _module.C1();
      _nw462.__ctor((_this).f21);
      _2800_v2 = _nw462;
      let _2801_v3;
      _2801_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2800_v2,(_this).f18);
      _2801_v3 = (_2801_v3).update(_2800_v2, (_this).fm3(p1, globalState));
      let _2802_v4;
      _2802_v4 = _dafny.Seq.of(_2798_v0);
      let _2803_v5;
      _2803_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(260),(_2802_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-354)), new BigNumber((_2802_v4).length))]);
      _2803_v5 = (((_this).f21) ? ((_2803_v5).Merge(_2803_v5)) : ((_2803_v5).Merge(_2803_v5)));
      let _2804_v6;
      _2804_v6 = _dafny.Set.fromElements((_this).f21);
      let _2805_v7;
      _2805_v7 = _dafny.Set.fromElements(_2804_v6);
      let _2806_v8;
      _2806_v8 = _module.D0.create_DC1(false, p1, p2, (_this).f18, new BigNumber(655));
      let _rhs445 = ((_dafny.ZERO).minus(new BigNumber(-851))).plus(p1);
      let _rhs446 = !(_2805_v7).equals(_2805_v7);
      let _rhs447 = ((_2800_v2).fm22((_this).f19, _2806_v8, globalState)).multipliedBy(p1);
      let _lhs385 = globalState;
      let _lhs386 = globalState;
      let _lhs387 = globalState;
      _lhs385.f3 = _rhs445;
      _lhs386.f2 = _rhs446;
      _lhs387.f3 = _rhs447;
      (globalState).f15 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), function (_2807_i0) {
        return (_this).f19;
      })).length), p1);
      return;
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let _2808_v0;
      _2808_v0 = new _dafny.CodePoint('h'.codePointAt(0));
      let _2809_v1;
      _2809_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2808_v0);
      let _2810_v2;
      _2810_v2 = _dafny.Seq.of(p0);
      let _2811_v3;
      let _nw463 = new _module.C4();
      _nw463.__ctor((((_2809_v1).contains(new BigNumber((_2810_v2).length))) ? ((_2809_v1).get(new BigNumber((_2810_v2).length))) : (_2808_v0)), p0, (_this).f21, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pjdlrtgp"), (_this).f19), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pjdlrtgp"), (_this).f19)).length)), _2808_v0), (_this).f20);
      _2811_v3 = _nw463;
      let _2812_v4;
      _2812_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(181),_2811_v3);
      _2811_v3 = (((_2812_v4).contains(p0)) ? ((_2812_v4).get(p0)) : (_2811_v3));
      let _2813_v5;
      let _nw464 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _2813_v5 = _nw464;
      let _2814_v6;
      let _nw465 = new _module.C5();
      _nw465.__ctor(_dafny.Seq.UnicodeFromString("ued"), (_this).f19, (_this).f18);
      _2814_v6 = _nw465;
      let _2815_v7;
      let _nw466 = Array((new BigNumber(6)).toNumber());
      _nw466[(_dafny.ZERO).toNumber()] = _2814_v6;
      _nw466[(_dafny.ONE).toNumber()] = _2814_v6;
      _nw466[(new BigNumber(2)).toNumber()] = _2814_v6;
      _nw466[(new BigNumber(3)).toNumber()] = _2814_v6;
      _nw466[(new BigNumber(4)).toNumber()] = _2814_v6;
      _nw466[(new BigNumber(5)).toNumber()] = _2814_v6;
      _2815_v7 = _nw466;
      let _2816_v8;
      _2816_v8 = _dafny.Seq.of(_2815_v7);
      let _2817_v9;
      _2817_v9 = _dafny.Seq.of(_2816_v8);
      let _2818_v10;
      _2818_v10 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_dafny.Seq.of(_2815_v7), (_2817_v9)[_module.__default.safeIndex(_module.__default.fm1(p0, false, true, globalState), new BigNumber((_2817_v9).length))]),_2813_v5);
      _2813_v5 = (((_2818_v10).contains(false)) ? ((_2818_v10).get(false)) : (_2813_v5));
      (globalState).f2 = (_2814_v6).fm15(p0, p0, globalState);
      let _hi14 = new BigNumber((_dafny.Seq.Concat(_2814_v6.f27, _dafny.Seq.UnicodeFromString("rwis"))).length);
      for (let _2819_i0 = (p0).plus((_dafny.ZERO).minus(_module.__default.fm1(p0, true, (_2811_v3).f18, globalState))); _2819_i0.isLessThan(_hi14); _2819_i0 = _2819_i0.plus(_dafny.ONE)) {
        let _2820_v11;
        _2820_v11 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(_2819_i0)),!_dafny.areEqual(_dafny.Seq.update((_2814_v6).f28, _module.__default.safeIndex(p0, new BigNumber(((_2814_v6).f28).length)), new _dafny.CodePoint('w'.codePointAt(0))), _2814_v6.f27));
        (globalState).f2 = (((_2820_v11).contains(p0)) ? ((_2820_v11).get(p0)) : (p2));
        let _2821_v12;
        let _init73 = ((_2822_v0) => function (_2823_i1) {
          return _2822_v0;
        })(_2808_v0);
        let _nw467 = Array((new BigNumber(27)).toNumber());
        for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw467.length); _i0_73++) {
          _nw467[_i0_73] = _init73(new BigNumber(_i0_73));
        }
        _2821_v12 = _nw467;
        let _index440 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2821_v12).length));
        (_2821_v12)[_index440] = new _dafny.CodePoint('w'.codePointAt(0));
        let _index441 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2821_v12).length));
        (_2821_v12)[_index441] = _2808_v0;
        _2813_v5 = _2813_v5;
        let _2824_v13;
        _2824_v13 = _dafny.Seq.of(false, true, p2);
        let _2825_v14;
        let _nw468 = Array((new BigNumber(29)).toNumber());
        _nw468[(_dafny.ZERO).toNumber()] = !((_this).f21);
        _nw468[(_dafny.ONE).toNumber()] = !((_2824_v13)[_module.__default.safeIndex(p0, new BigNumber((_2824_v13).length))]);
        _nw468[(new BigNumber(2)).toNumber()] = (_this).f18;
        _nw468[(new BigNumber(3)).toNumber()] = (_this).f21;
        _nw468[(new BigNumber(4)).toNumber()] = (_this).f18;
        _nw468[(new BigNumber(5)).toNumber()] = (_this).f18;
        _nw468[(new BigNumber(6)).toNumber()] = (_2811_v3).f18;
        _nw468[(new BigNumber(7)).toNumber()] = (_this).f18;
        _nw468[(new BigNumber(8)).toNumber()] = (_this).f18;
        _nw468[(new BigNumber(9)).toNumber()] = false;
        _nw468[(new BigNumber(10)).toNumber()] = (_2811_v3).f18;
        _nw468[(new BigNumber(11)).toNumber()] = (_2811_v3).f18;
        _nw468[(new BigNumber(12)).toNumber()] = p2;
        _nw468[(new BigNumber(13)).toNumber()] = (_2811_v3).f18;
        _nw468[(new BigNumber(14)).toNumber()] = true;
        _nw468[(new BigNumber(15)).toNumber()] = false;
        _nw468[(new BigNumber(16)).toNumber()] = p2;
        _nw468[(new BigNumber(17)).toNumber()] = !((_this).f21);
        _nw468[(new BigNumber(18)).toNumber()] = (_this).f21;
        _nw468[(new BigNumber(19)).toNumber()] = (_2811_v3).f18;
        _nw468[(new BigNumber(20)).toNumber()] = !((_2811_v3).f18);
        _nw468[(new BigNumber(21)).toNumber()] = p2;
        _nw468[(new BigNumber(22)).toNumber()] = (_this).f21;
        _nw468[(new BigNumber(23)).toNumber()] = (_2811_v3).f18;
        _nw468[(new BigNumber(24)).toNumber()] = (_2811_v3).f18;
        _nw468[(new BigNumber(25)).toNumber()] = (_2811_v3).f18;
        _nw468[(new BigNumber(26)).toNumber()] = p2;
        _nw468[(new BigNumber(27)).toNumber()] = p2;
        _nw468[(new BigNumber(28)).toNumber()] = (_2811_v3).f18;
        _2825_v14 = _nw468;
        let _2826_v15;
        let _nw469 = new _module.C9();
        _nw469.__ctor(_2825_v14, (_2811_v3).f18);
        _2826_v15 = _nw469;
      }
      let _index442 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2813_v5).length));
      (_2813_v5)[_index442] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, p0));
      let _index443 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2813_v5).length));
      (_2813_v5)[_index443] = p0;
      (globalState).f2 = (_this).f18;
      return;
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let _2827_v0;
      _2827_v0 = new BigNumber(534);
      let _2828_v1;
      let _nw470 = new _module.C7();
      _nw470.__ctor(_2827_v0, (_this).f18);
      _2828_v1 = _nw470;
      let _2829_i0;
      _2829_i0 = _dafny.ZERO;
      L17: {
        while ((_this).f18) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2829_i0)) {
              break L17;
            }
            _2829_i0 = (_2829_i0).plus(_dafny.ONE);
            (globalState).f9 = _module.__default.fm12(_2827_v0, (_this).f21, globalState);
            let _2830_v2;
            _2830_v2 = _dafny.Set.fromElements(new BigNumber(((_this).f19).length), _2827_v0);
            (_2828_v1).f24 = _module.__default.fm1(_2828_v1.f24, _module.__default.fm23(_2830_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length), globalState), (_this).f21, globalState);
            let _2831_v3;
            let _nw471 = new _module.C1();
            _nw471.__ctor(p0);
            _2831_v3 = _nw471;
            let _2832_v4;
            let _nw472 = new _module.C7();
            _nw472.__ctor((new BigNumber(451)).minus(_2828_v1.f24), (_2827_v0).isEqualTo(new BigNumber(79)));
            _2832_v4 = _nw472;
          }
        }
      }
      let _2833_v5;
      let _nw473 = new _module.C3();
      _nw473.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(866)), ((_2834_v1) => function (_2835_i1) {
        return _2834_v1.f24;
      })(_2828_v1)), (_this).f19, (_this).f20, p0);
      _2833_v5 = _nw473;
      let _2836_v6;
      _2836_v6 = _dafny.Seq.of(_2833_v5, _2833_v5);
      _2836_v6 = _dafny.Seq.Concat(_2836_v6, _dafny.Seq.of(_2833_v5, _2833_v5));
      let _2837_v7;
      _2837_v7 = _dafny.MultiSet.fromElements((_2827_v0).minus(new BigNumber(-987)));
      _2837_v7 = _2837_v7;
      let _2838_v8;
      _2838_v8 = _dafny.Set.fromElements((_this).f19, (_this).f19, (_this).f19, (_this).f19, (_this).f19);
      let _2839_v9;
      let _init74 = ((_2840_v1) => function (_2841_i2) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),_2840_v1.f24)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),_2840_v1.f24));
      })(_2828_v1);
      let _nw474 = Array((new BigNumber(26)).toNumber());
      for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw474.length); _i0_74++) {
        _nw474[_i0_74] = _init74(new BigNumber(_i0_74));
      }
      _2839_v9 = _nw474;
      let _2842_v10;
      _2842_v10 = new _dafny.CodePoint('i'.codePointAt(0));
      let _2843_v11;
      _2843_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2842_v10,new BigNumber(165));
      let _index444 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_2839_v9).length));
      (_2839_v9)[_index444] = (_2843_v11).Merge(_2843_v11);
      let _2844_v12;
      _2844_v12 = _dafny.Set.fromElements(_2827_v0);
      let _2845_v13;
      _2845_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-385)), ((_2846_v10) => function (_2847_i3) {
        return _2846_v10;
      })(_2842_v10))),_2827_v0);
      let _2848_v15;
      _2848_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2842_v10,true);
      let _index445 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_2839_v9).length));
      let _rhs448 = _module.__default.fm23(_2844_v12, new BigNumber((_2845_v13).length), globalState);
      let _rhs449 = _2827_v0;
      let _rhs450 = (_dafny.Set.fromElements((_this).f19, _dafny.Seq.UnicodeFromString("pewnmqk"), _dafny.Seq.UnicodeFromString("cswc"), _dafny.Seq.UnicodeFromString("erclvwh"), (_this).f19)).Union(_2838_v8);
      let _rhs451 = _2827_v0;
      let _rhs452 = ((_module.__default.fm64(_2842_v10, _2828_v1.f24, p0, p0, globalState)).Merge(_2843_v11)).Merge(function () {
        let _coll67 = new _dafny.Map();
        for (const _compr_67 of (_2848_v15).Keys.Elements) {
          let _2849_v14 = _compr_67;
          if ((_2848_v15).contains(_2849_v14)) {
            _coll67.push([_2849_v14,(_dafny.ZERO).minus(_2827_v0)]);
          }
        }
        return _coll67;
      }());
      let _lhs388 = globalState;
      let _lhs389 = globalState;
      let _lhs390 = _2839_v9;
      let _lhs391 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_2839_v9).length));
      r2 = _rhs448;
      _lhs388.f15 = _rhs449;
      _2838_v8 = _rhs450;
      _lhs389.f0 = _rhs451;
      _lhs390[_lhs391] = _rhs452;
      let _2850_v16;
      _2850_v16 = _module.D17.create_DC42((_this).f18);
      let _source27 = _2850_v16;
      if (_source27.is_DC42) {
        let _2851___mcc_h0 = (_source27).cf74;
        let _2852_cf74 = _2851___mcc_h0;
        let _2853_v17;
        let _init75 = ((_2854_p0) => function (_2855_i4) {
          return (true) || (!(_2854_p0));
        })(p0);
        let _nw475 = Array((new BigNumber(16)).toNumber());
        for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw475.length); _i0_75++) {
          _nw475[_i0_75] = _init75(new BigNumber(_i0_75));
        }
        _2853_v17 = _nw475;
        let _index446 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2853_v17).length));
        (_2853_v17)[_index446] = (_this).f18;
        let _index447 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2853_v17).length));
        (_2853_v17)[_index447] = (_2852_cf74) && ((_this).f21);
        let _2856_v18;
        _2856_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2828_v1.f24,(_this).f18);
        let _2857_v19;
        _2857_v19 = _module.D17.create_DC41(_2856_v18);
        let _2858_v20;
        _2858_v20 = _dafny.Set.fromElements(_2857_v19, _2857_v19);
        let _2859_v21;
        let _init76 = ((_2860_v1) => function (_2861_i5) {
          return (_2861_i5).plus(_2860_v1.f24);
        })(_2828_v1);
        let _nw476 = Array((new BigNumber(10)).toNumber());
        for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw476.length); _i0_76++) {
          _nw476[_i0_76] = _init76(new BigNumber(_i0_76));
        }
        _2859_v21 = _nw476;
        let _2862_v22;
        _2862_v22 = _module.D24.create_DC65(p0, _2858_v20, (_this).f21, _2827_v0, _2859_v21);
        if ((_2862_v22).dtor_cf112) {
          let _2863_v23;
          _2863_v23 = _dafny.Set.fromElements(p0, _2852_cf74);
          let _2864_v24;
          _2864_v24 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), ((_2865_v17) => function (_2866_i6) {
            return new BigNumber((_dafny.Set.fromElements(!((_2865_v17)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2865_v17).length))]))).length);
          })(_2853_v17)), _dafny.Seq.update((_2833_v5).f35, _module.__default.safeIndex(_2827_v0, new BigNumber(((_2833_v5).f35).length)), new BigNumber((_2863_v23).length)));
          let _2867_v25;
          _2867_v25 = _dafny.Map.Empty.slice().updateUnsafe((_2864_v24)[_module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber((_2864_v24).length))],_2852_cf74);
          _2867_v25 = ((_2867_v25).Merge((_2867_v25).update((_2833_v5).f35, p0))).Merge((((_this).f18) ? (_2867_v25) : (_2867_v25)));
          let _2868_v26;
          let _nw477 = Array((new BigNumber(6)).toNumber());
          _2868_v26 = _nw477;
          let _2869_v27;
          let _nw478 = new _module.C2();
          _nw478.__ctor(_2828_v1.f24, _dafny.Seq.of(_2828_v1.f24, new BigNumber(904)), false);
          _2869_v27 = _nw478;
          let _index448 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_2868_v26).length));
          (_2868_v26)[_index448] = _2869_v27;
          let _index449 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_2868_v26).length));
          (_2868_v26)[_index449] = _2869_v27;
          (globalState).f2 = !(_module.__default.safeModuloInt(_2869_v27.f33, (_dafny.ZERO).minus(_2828_v1.f24))).isEqualTo((_dafny.ZERO).minus((_2833_v5).fm38(_dafny.Seq.UnicodeFromString("dqcnaoqeg"), globalState)));
          let _index450 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2859_v21).length));
          (_2859_v21)[_index450] = _2828_v1.f24;
          let _2870_v28;
          _2870_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2853_v17,(_2853_v17)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2853_v17).length))]);
          let _2871_v29;
          _2871_v29 = _module.D13.create_DC32(_2870_v28);
          let _pat_let_tv37 = _2870_v28;
          let _pat_let_tv38 = _2870_v28;
          let _pat_let_tv39 = _2853_v17;
          let _2872_v30;
          let _nw479 = Array((new BigNumber(12)).toNumber());
          _nw479[(_dafny.ZERO).toNumber()] = _2871_v29;
          _nw479[(_dafny.ONE).toNumber()] = function (_pat_let74_0) {
            return function (_2873_dt__update__tmp_h0) {
              return function (_pat_let75_0) {
                return function (_2874_dt__update_hcf54_h0) {
                  return _module.D13.create_DC32(_2874_dt__update_hcf54_h0);
                }(_pat_let75_0);
              }(_pat_let_tv37);
            }(_pat_let74_0);
          }(_2871_v29);
          _nw479[(new BigNumber(2)).toNumber()] = _2871_v29;
          _nw479[(new BigNumber(3)).toNumber()] = _2871_v29;
          _nw479[(new BigNumber(4)).toNumber()] = _2871_v29;
          _nw479[(new BigNumber(5)).toNumber()] = function (_pat_let76_0) {
            return function (_2875_dt__update__tmp_h1) {
              return function (_pat_let77_0) {
                return function (_2876_dt__update_hcf54_h1) {
                  return _module.D13.create_DC32(_2876_dt__update_hcf54_h1);
                }(_pat_let77_0);
              }((_pat_let_tv38).update(_pat_let_tv39, (_this).f21));
            }(_pat_let76_0);
          }(_2871_v29);
          _nw479[(new BigNumber(6)).toNumber()] = _module.D13.create_DC32(_2870_v28);
          _nw479[(new BigNumber(7)).toNumber()] = _2871_v29;
          _nw479[(new BigNumber(8)).toNumber()] = _2871_v29;
          _nw479[(new BigNumber(9)).toNumber()] = ((p0) ? (_2871_v29) : (_2871_v29));
          _nw479[(new BigNumber(10)).toNumber()] = _module.D13.create_DC32(_2870_v28);
          _nw479[(new BigNumber(11)).toNumber()] = _2871_v29;
          _2872_v30 = _nw479;
          let _index451 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_2872_v30).length));
          (_2872_v30)[_index451] = _2871_v29;
          let _pat_let_tv40 = _2870_v28;
          let _index452 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2859_v21).length));
          let _index453 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_2872_v30).length));
          let _index454 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2853_v17).length));
          let _rhs453 = (new BigNumber(658)).plus(new BigNumber(236));
          let _rhs454 = ((!(_2852_cf74) || ((_this).f21)) ? (_2871_v29) : (function (_pat_let78_0) {
            return function (_2877_dt__update__tmp_h2) {
              return function (_pat_let79_0) {
                return function (_2878_dt__update_hcf54_h2) {
                  return _module.D13.create_DC32(_2878_dt__update_hcf54_h2);
                }(_pat_let79_0);
              }(_pat_let_tv40);
            }(_pat_let78_0);
          }(_2871_v29)));
          let _rhs455 = (_this).f18;
          let _rhs456 = !_dafny.Seq.contains((_this).f19, _2842_v10);
          let _rhs457 = (_2833_v5).f35;
          let _lhs392 = _2859_v21;
          let _lhs393 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2859_v21).length));
          let _lhs394 = _2872_v30;
          let _lhs395 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_2872_v30).length));
          let _lhs396 = _2853_v17;
          let _lhs397 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2853_v17).length));
          let _lhs398 = globalState;
          _lhs392[_lhs393] = _rhs453;
          _lhs394[_lhs395] = _rhs454;
          _lhs396[_lhs397] = _rhs455;
          _2852_cf74 = _rhs456;
          _lhs398.f10 = _rhs457;
          let _index455 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2859_v21).length));
          (_2859_v21)[_index455] = _2828_v1.f24;
        } else {
          r2 = (_this).f21;
          let _2879_v31;
          _2879_v31 = _dafny.Seq.of((_this).f18, false);
          let _2880_v32;
          let _nw480 = new _module.C5();
          _nw480.__ctor((_this).f19, (_this).f19, true);
          _2880_v32 = _nw480;
          let _2881_v33;
          _2881_v33 = _module.D24.create_DC64((_2879_v31)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_2879_v31).length))], _2852_cf74, _2880_v32, new BigNumber(((_this).f19).length), _module.__default.safeModuloInt(_2828_v1.f24, (_dafny.ZERO).minus(_2827_v0)));
          _2881_v33 = function (_pat_let80_0) {
            return function (_2882_dt__update__tmp_h3) {
              return function (_pat_let81_0) {
                return function (_2883_dt__update_hcf106_h0) {
                  return _module.D24.create_DC64((_2882_dt__update__tmp_h3).dtor_cf105, _2883_dt__update_hcf106_h0, (_2882_dt__update__tmp_h3).dtor_cf107, (_2882_dt__update__tmp_h3).dtor_cf108, (_2882_dt__update__tmp_h3).dtor_cf109);
                }(_pat_let81_0);
              }((_this).f18);
            }(_pat_let80_0);
          }(_module.D24.create_DC64(p0, _2852_cf74, _2880_v32, _2828_v1.f24, _2828_v1.f24));
          let _2884_v34;
          let _nw481 = new _module.C1();
          _nw481.__ctor(_2852_cf74);
          _2884_v34 = _nw481;
          let _2885_v35;
          let _init77 = ((_2886_v10) => function (_2887_i7) {
            return _2886_v10;
          })(_2842_v10);
          let _nw482 = Array((new BigNumber(3)).toNumber());
          for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw482.length); _i0_77++) {
            _nw482[_i0_77] = _init77(new BigNumber(_i0_77));
          }
          _2885_v35 = _nw482;
          let _2888_v36;
          _2888_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe(_2885_v35,_2828_v1.f24));
          let _2889_v37;
          _2889_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2885_v35,_2828_v1.f24);
          _2888_v36 = (_2888_v36).update(p0, (_2889_v37).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2885_v35,new BigNumber((_2856_v18).length))));
          let _2890_v38;
          let _2891_v39;
          let _2892_v40;
          let _2893_v41;
          let _out66;
          let _out67;
          let _out68;
          let _out69;
          let _outcollector18 = (_2833_v5).m10(p0, globalState);
          _out66 = _outcollector18[0];
          _out67 = _outcollector18[1];
          _out68 = _outcollector18[2];
          _out69 = _outcollector18[3];
          _2890_v38 = _out66;
          _2891_v39 = _out67;
          _2892_v40 = _out68;
          _2893_v41 = _out69;
        }
        (globalState).f0 = new BigNumber(950);
        let _2894_v42;
        _2894_v42 = _dafny.Seq.of((_this).f19, (_this).f19);
        let _2895_v43;
        _2895_v43 = _module.D26.create_DC69(_2894_v42);
        _2894_v42 = (_2895_v43).dtor_cf121;
      } else if (_source27.is_DC43) {
        let _2896___mcc_h1 = (_source27).cf75;
        let _2897___mcc_h2 = (_source27).cf76;
        let _2898_cf76 = _2897___mcc_h2;
        let _2899_cf75 = _2896___mcc_h1;
        (_2828_v1).f24 = (_dafny.ZERO).minus(((_2898_cf76) ? (_2827_v0) : ((_2833_v5).fm38((_this).f19, globalState))));
        let _2900_v44;
        _2900_v44 = _dafny.Seq.of((_this).f21, false);
        _2900_v44 = _2900_v44;
        let _2901_v45;
        _2901_v45 = _dafny.Set.fromElements((_this).f18, _2898_cf76, _2898_cf76, p0, !((_this).f18));
        let _2902_v46;
        _2902_v46 = _dafny.Seq.of((_dafny.Set.fromElements(p0)).Intersect(_2901_v45), _2901_v45);
        let _2903_v47;
        _2903_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2827_v0,_2828_v1.f24);
        _2902_v46 = _module.__default.fm62(new BigNumber((_dafny.MultiSet.fromElements(_2828_v1.f24)).cardinality()), _2903_v47, globalState);
        _2900_v44 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2900_v44, _2900_v44), (((_this).f21) ? (_2900_v44) : (_dafny.Seq.update(_2900_v44, _module.__default.safeIndex(_2827_v0, new BigNumber((_2900_v44).length)), (_this).f18))));
      } else if (_source27.is_DC44) {
        let _2904___mcc_h3 = (_source27).cf77;
        let _2905___mcc_h4 = (_source27).cf78;
        let _2906_cf78 = _2905___mcc_h4;
        let _2907_cf77 = _2904___mcc_h3;
        _2842_v10 = _2842_v10;
        r1 = (_this).f18;
        let _2908_v48;
        _2908_v48 = _module.D19.create_DC50(!((_this).f18));
        let _source28 = _module.D19.create_DC51(_2908_v48);
        if (_source28.is_DC48) {
          let _2909___mcc_h6 = (_source28).cf82;
          let _2910_cf82 = _2909___mcc_h6;
          let _2911_v49;
          _2911_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,true);
          _2911_v49 = (_2911_v49).update(_2910_cf82, (_this).f21);
          let _2912_v50;
          let _nw483 = new _module.C2();
          _nw483.__ctor(((p0) ? (new BigNumber((_2837_v7).cardinality())) : (_2828_v1.f24)), (_2833_v5).f35, (_this).f21);
          _2912_v50 = _nw483;
          (globalState).f15 = (_2828_v1.f24).minus(_2828_v1.f24);
          let _2913_v51;
          _2913_v51 = _dafny.Seq.of((_this).f19, _dafny.Seq.update((_this).f19, _module.__default.safeIndex(_2828_v1.f24, new BigNumber(((_this).f19).length)), _2842_v10), (_this).f19);
          (globalState).f3 = new BigNumber(((((_this).f21) ? (_dafny.Seq.Concat((_this).f19, _dafny.Seq.UnicodeFromString("eofv"))) : ((_2913_v51)[_module.__default.safeIndex(_2828_v1.f24, new BigNumber((_2913_v51).length))]))).length);
        } else if (_source28.is_DC49) {
          let _2914___mcc_h7 = (_source28).cf83;
          let _2915_cf83 = _2914___mcc_h7;
          (globalState).f15 = new BigNumber(((_2833_v5).f35).length);
          (_2828_v1).f24 = (new BigNumber((_module.__default.fm48(globalState)).length)).multipliedBy(_2828_v1.f24);
          let _index456 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_2907_cf77).length));
          (_2907_cf77)[_index456] = p0;
          let _2916_v52;
          _2916_v52 = _module.D10.create_DC25(_2915_cf83);
          let _index457 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_2907_cf77).length));
          (_2907_cf77)[_index457] = (_2916_v52).dtor_cf48;
          (globalState).f15 = _2906_cf78;
        } else if (_source28.is_DC50) {
          let _2917___mcc_h8 = (_source28).cf84;
          let _2918_cf84 = _2917___mcc_h8;
          let _2919_v53;
          _2919_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f19).length),_2918_cf84);
          let _2920_v54;
          _2920_v54 = _dafny.MultiSet.fromElements((_this).f19, (_this).f19, (_this).f19, _dafny.Seq.update((_this).f19, _module.__default.safeIndex(new BigNumber((_2919_v53).length), new BigNumber(((_this).f19).length)), _2842_v10), _dafny.Seq.Create(_module.__default.abs(new BigNumber(186)), ((_2921_v10) => function (_2922_i8) {
            return _2921_v10;
          })(_2842_v10)));
          (_2828_v1).m5(_2920_v54, _2828_v1.f24, _2837_v7, _module.__default.safeDivisionInt(_2828_v1.f24, new BigNumber(782)), globalState);
          let _2923_v55;
          let _nw484 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _2923_v55 = _nw484;
          let _2924_v56;
          _2924_v56 = _dafny.Seq.of(_2919_v53);
          let _index458 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_2923_v55).length));
          (_2923_v55)[_index458] = new BigNumber((_2924_v56).length);
          let _index459 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_2923_v55).length));
          (_2923_v55)[_index459] = _module.__default.safeModuloInt(_2828_v1.f24, _2827_v0);
          let _2925_v57;
          let _nw485 = new _module.C3();
          _nw485.__ctor(_dafny.Seq.of(_2828_v1.f24, new BigNumber(754)), (_2833_v5).fm37(_2828_v1.f24, globalState), (_this).f20, (_2833_v5).fm2(!(_2918_cf84), _2918_cf84, _2918_cf84, globalState));
          _2925_v57 = _nw485;
          let _2926_v58;
          _2926_v58 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_2925_v57));
          let _2927_v59;
          _2927_v59 = _dafny.Seq.of(_2925_v57);
          let _2928_v60;
          _2928_v60 = _dafny.Seq.of(_2927_v59);
          let _2929_v61;
          _2929_v61 = _dafny.Seq.of(_2926_v58, _dafny.MultiSet.FromArray(_2928_v60), _2926_v58, _2926_v58);
          let _2930_v62;
          _2930_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2927_v59);
          let _2931_v63;
          _2931_v63 = _2928_v60;
          let _2932_v64;
          let _nw486 = Array((new BigNumber(20)).toNumber());
          _nw486[(_dafny.ZERO).toNumber()] = _2926_v58;
          _nw486[(_dafny.ONE).toNumber()] = (_2926_v58).Difference(_2926_v58);
          _nw486[(new BigNumber(2)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(3)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(4)).toNumber()] = (_2926_v58).Difference(_2926_v58);
          _nw486[(new BigNumber(5)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(6)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(7)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(8)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(9)).toNumber()] = (_2926_v58).Union((_2929_v61)[_module.__default.safeIndex(_2828_v1.f24, new BigNumber((_2929_v61).length))]);
          _nw486[(new BigNumber(10)).toNumber()] = (_2926_v58).Difference(_dafny.MultiSet.fromElements((((_2930_v62).contains(_2918_cf84)) ? ((_2930_v62).get(_2918_cf84)) : (_2927_v59))));
          _nw486[(new BigNumber(11)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_2927_v59);
          _nw486[(new BigNumber(13)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(14)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(15)).toNumber()] = (_2926_v58).update(_2927_v59, _module.__default.abs(new BigNumber(((_2925_v57).f19).length)));
          _nw486[(new BigNumber(16)).toNumber()] = _2926_v58;
          _nw486[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.fromElements(_dafny.Seq.of(_2925_v57, _2925_v57), _dafny.Seq.update(_2927_v59, _module.__default.safeIndex(_2827_v0, new BigNumber((_2927_v59).length)), _2925_v57));
          _nw486[(new BigNumber(18)).toNumber()] = (_2926_v58).Difference(_dafny.MultiSet.FromArray((_2931_v63)));
          _nw486[(new BigNumber(19)).toNumber()] = (_2926_v58).Intersect(_2926_v58);
          _2932_v64 = _nw486;
          let _2933_v65;
          _2933_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2918_cf84,_2926_v58);
          let _index460 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_2932_v64).length));
          (_2932_v64)[_index460] = (((_2933_v65).contains((_this).f21)) ? ((_2933_v65).get((_this).f21)) : (_2926_v58));
          let _2934_v66;
          _2934_v66 = _module.D28.create_DC74(_2926_v58);
          let _index461 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_2932_v64).length));
          let _rhs458 = _module.__default.fm12(new BigNumber(((_2833_v5).f35).length), (_2925_v57).f18, globalState);
          let _rhs459 = (_2934_v66).dtor_cf131;
          let _lhs399 = globalState;
          let _lhs400 = _2932_v64;
          let _lhs401 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_2932_v64).length));
          _lhs399.f9 = _rhs458;
          _lhs400[_lhs401] = _rhs459;
          let _2935_v67;
          let _nw487 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _2935_v67 = _nw487;
          let _2936_v68;
          _2936_v68 = _module.D20.create_DC53((_this).f18, (_2925_v57).f18, _2842_v10);
          let _index462 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_2935_v67).length));
          (_2935_v67)[_index462] = (_dafny.Map.Empty.slice().updateUnsafe(_2918_cf84,p0)).update((_2936_v68).dtor_cf87, _2918_cf84);
          let _2937_v69;
          _2937_v69 = _dafny.Seq.of((_this).f21);
          let _2938_v70;
          _2938_v70 = _dafny.Seq.of((_2937_v69)[_module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber((_2937_v69).length))]);
          let _2939_v71;
          _2939_v71 = _dafny.Map.Empty.slice().updateUnsafe((_2937_v69)[_module.__default.safeIndex((_dafny.ZERO).minus(_2828_v1.f24), new BigNumber((_2937_v69).length))],_2918_cf84);
          let _2940_v72;
          _2940_v72 = _dafny.Seq.of(_2939_v71);
          let _2941_v73;
          _2941_v73 = _dafny.Seq.of((_2939_v71).Merge(_2939_v71), _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f21), _2939_v71, (_2940_v72)[_module.__default.safeIndex(_2828_v1.f24, new BigNumber((_2940_v72).length))], _2939_v71);
          let _index463 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_2935_v67).length));
          (_2935_v67)[_index463] = (_2940_v72)[_module.__default.safeIndex((new BigNumber((_dafny.MultiSet.fromElements((_2925_v57).f18)).cardinality())).minus(new BigNumber(((_this).f19).length)), new BigNumber((_2940_v72).length))];
        } else if (_source28.is_DC47) {
          let _2942___mcc_h9 = (_source28).cf81;
          let _2943_cf81 = _2942___mcc_h9;
          let _2944_v74;
          let _nw488 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _2944_v74 = _nw488;
          let _2945_v75;
          _2945_v75 = _dafny.Seq.of(_2944_v74, _2944_v74);
          let _2946_v76;
          _2946_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2944_v74);
          let _2947_v77;
          _2947_v77 = _dafny.Seq.of((_this).f21);
          let _2948_v78;
          _2948_v78 = _dafny.Map.Empty.slice().updateUnsafe(_2944_v74,new BigNumber((_2947_v77).length));
          let _2949_v79;
          _2949_v79 = _dafny.Set.fromElements(true);
          let _2950_v80;
          _2950_v80 = _module.D16.create_DC40(new BigNumber((_module.__default.fm25(new BigNumber(632), globalState)).cardinality()), (_2833_v5).fm38((_this).f19, globalState), _2828_v1.f24, new BigNumber((_2949_v79).length));
          let _2951_v81;
          _2951_v81 = _dafny.Map.Empty.slice().updateUnsafe(_2828_v1.f24,_2950_v80);
          let _2952_v83;
          let _nw489 = Array((new BigNumber(17)).toNumber());
          _nw489[(_dafny.ZERO).toNumber()] = _2951_v81;
          _nw489[(_dafny.ONE).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(2)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(3)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(4)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(5)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(6)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(7)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(8)).toNumber()] = function () {
            let _coll68 = new _dafny.Map();
            for (const _compr_68 of _dafny.IntegerRange(new BigNumber(579), new BigNumber(785))) {
              let _2953_v82 = _compr_68;
              if (((new BigNumber(579)).isLessThanOrEqualTo(_2953_v82)) && ((_2953_v82).isLessThan(new BigNumber(785)))) {
                _coll68.push([(_2953_v82).multipliedBy(_2828_v1.f24),_2950_v80]);
              }
            }
            return _coll68;
          }();
          _nw489[(new BigNumber(9)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(10)).toNumber()] = (_2951_v81).update(_2828_v1.f24, _2950_v80);
          _nw489[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2827_v0,_2950_v80);
          _nw489[(new BigNumber(12)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(13)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(14)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(15)).toNumber()] = _2951_v81;
          _nw489[(new BigNumber(16)).toNumber()] = _2951_v81;
          _2952_v83 = _nw489;
          let _2954_v84;
          _2954_v84 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2948_v78).length),_2952_v83);
          let _2955_v85;
          _2955_v85 = _module.D26.create_DC70(_2828_v1.f24, _2906_cf78, p0, (((_2954_v84).contains(_2828_v1.f24)) ? ((_2954_v84).get(_2828_v1.f24)) : (_2952_v83)));
          let _2956_v86;
          _2956_v86 = _dafny.Map.Empty.slice().updateUnsafe(_2955_v85,(_this).f18);
          let _2957_v87;
          _2957_v87 = _dafny.Map.Empty.slice().updateUnsafe(_2906_cf78,_2944_v74);
          let _2958_v88;
          let _nw490 = Array((new BigNumber(27)).toNumber());
          _nw490[(_dafny.ZERO).toNumber()] = _2944_v74;
          _nw490[(_dafny.ONE).toNumber()] = (_2945_v75)[_module.__default.safeIndex(_2906_cf78, new BigNumber((_2945_v75).length))];
          _nw490[(new BigNumber(2)).toNumber()] = (((_2946_v76).contains((((_2956_v86).contains(_2955_v85)) ? ((_2956_v86).get(_2955_v85)) : ((_this).f18)))) ? ((_2946_v76).get((((_2956_v86).contains(_2955_v85)) ? ((_2956_v86).get(_2955_v85)) : ((_this).f18)))) : (_2944_v74));
          _nw490[(new BigNumber(3)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(4)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(5)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(6)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(7)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(8)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(9)).toNumber()] = ((false) ? (_2944_v74) : (_2944_v74));
          _nw490[(new BigNumber(10)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(11)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(12)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(13)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(14)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(15)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(16)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(17)).toNumber()] = ((p0) ? (_2944_v74) : (_2944_v74));
          _nw490[(new BigNumber(18)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(19)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(20)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(21)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(22)).toNumber()] = (((_2957_v87).contains(_2827_v0)) ? ((_2957_v87).get(_2827_v0)) : (_2944_v74));
          _nw490[(new BigNumber(23)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(24)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(25)).toNumber()] = _2944_v74;
          _nw490[(new BigNumber(26)).toNumber()] = _2944_v74;
          _2958_v88 = _nw490;
          let _index464 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_2958_v88).length));
          (_2958_v88)[_index464] = _2944_v74;
          let _2959_v89;
          _2959_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2828_v1.f24,p0);
          let _2960_v90;
          _2960_v90 = _module.D17.create_DC41(_2959_v89);
          let _2961_v92;
          _2961_v92 = _module.D24.create_DC65(false, function () {
  let _coll69 = new _dafny.Set();
  for (const _compr_69 of (_dafny.MultiSet.fromElements(_2960_v90)).Elements) {
    let _2962_v91 = _compr_69;
    if ((_dafny.MultiSet.fromElements(_2960_v90)).contains(_2962_v91)) {
      _coll69.add(_2962_v91);
    }
  }
  return _coll69;
}(), (_this).f18, _2828_v1.f24, _2944_v74);
          let _index465 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_2958_v88).length));
          (_2958_v88)[_index465] = (_2961_v92).dtor_cf114;
          let _2963_v93;
          let _nw491 = Array((new BigNumber(3)).toNumber());
          _nw491[(_dafny.ZERO).toNumber()] = _2837_v7;
          _nw491[(_dafny.ONE).toNumber()] = _2837_v7;
          _nw491[(new BigNumber(2)).toNumber()] = _2837_v7;
          _2963_v93 = _nw491;
          let _2964_v94;
          _2964_v94 = _module.D22.create_DC59(_module.D22.create_DC57(_2963_v93));
          let _2965_v95;
          _2965_v95 = _module.D22.create_DC59(_2964_v94);
          let _arr5 = (_2958_v88)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_2958_v88).length))];
          let _index466 = _module.__default.safeIndex(new BigNumber(885), new BigNumber(((_2958_v88)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_2958_v88).length))]).length));
          _arr5[_index466] = _2828_v1.f24;
          let _arr6 = (_2958_v88)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_2958_v88).length))];
          let _index467 = _module.__default.safeIndex(new BigNumber(885), new BigNumber(((_2958_v88)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_2958_v88).length))]).length));
          let _rhs460 = (_this).f21;
          let _rhs461 = _2965_v95;
          let _rhs462 = _2907_cf77;
          let _rhs463 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_2828_v1.f24, (((_this).f21) ? (_module.__default.fm1(_2828_v1.f24, p0, p0, globalState)) : (_2906_cf78))));
          let _lhs402 = (_2958_v88)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_2958_v88).length))];
          let _lhs403 = _module.__default.safeIndex(new BigNumber(885), new BigNumber(((_2958_v88)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_2958_v88).length))]).length));
          r1 = _rhs460;
          _2965_v95 = _rhs461;
          _2907_cf77 = _rhs462;
          _lhs402[_lhs403] = _rhs463;
          let _2966_v96;
          _2966_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(449),new BigNumber(551));
          (globalState).f3 = new BigNumber((_2966_v96).length);
          let _2967_v97;
          _2967_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p0);
          let _2968_v98;
          _2968_v98 = _dafny.Map.Empty.slice().updateUnsafe(_2827_v0,_2967_v97);
          _2968_v98 = (_2968_v98).update(_2828_v1.f24, _2967_v97);
        } else {
          let _2969___mcc_h10 = (_source28).cf85;
          let _2970_cf85 = _2969___mcc_h10;
          (globalState).f3 = _module.__default.safeModuloInt(_2828_v1.f24, ((p0) ? (_2906_cf78) : (new BigNumber((function () {
            let _coll70 = new _dafny.Set();
            for (const _compr_70 of (_dafny.MultiSet.fromElements(_2827_v0)).Elements) {
              let _2971_v99 = _compr_70;
              if ((_dafny.MultiSet.fromElements(_2827_v0)).contains(_2971_v99)) {
                _coll70.add((_2971_v99).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.of((_module.D17.create_DC43(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length),new BigNumber(-961))), false)).dtor_cf76, false, true, false), _dafny.Seq.of(false, false))).length))));
              }
            }
            return _coll70;
          }()).length))));
          (globalState).f15 = _2906_cf78;
          r1 = false;
          let _2972_v100;
          _2972_v100 = _module.D25.create_DC67((_this).f21, new BigNumber(((_this).f19).length), (_2833_v5).f35, (_this).f21);
          let _2973_v101;
          let _nw492 = Array((new BigNumber(6)).toNumber());
          _nw492[(_dafny.ZERO).toNumber()] = (_2833_v5).f35;
          _nw492[(_dafny.ONE).toNumber()] = (_2833_v5).f35;
          _nw492[(new BigNumber(2)).toNumber()] = (_2833_v5).f35;
          _nw492[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat((_2833_v5).f35, _dafny.Seq.of(_2828_v1.f24, _2906_cf78));
          _nw492[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(606)), ((_2974_v1) => function (_2975_i9) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2974_v1.f24,true)).length);
          })(_2828_v1)), _dafny.Seq.of(_2828_v1.f24));
          _nw492[(new BigNumber(5)).toNumber()] = (_2972_v100).dtor_cf118;
          _2973_v101 = _nw492;
          let _index468 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_2973_v101).length));
          (_2973_v101)[_index468] = _dafny.Seq.Concat((_2833_v5).f35, (_2972_v100).dtor_cf118);
          let _index469 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_2973_v101).length));
          (_2973_v101)[_index469] = _dafny.Seq.update(_dafny.Seq.of(new BigNumber(((_this).f19).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f21,_2828_v1.f24)).length)), _module.__default.safeIndex((_2827_v0).plus(_2828_v1.f24), new BigNumber((_dafny.Seq.of(new BigNumber(((_this).f19).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f21,_2828_v1.f24)).length))).length)), (_2906_cf78).plus(new BigNumber((function () {
            let _coll71 = new _dafny.Set();
            for (const _compr_71 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(698)), ((_2976_v5) => function (_2977_i10) {
              return new BigNumber(((_2976_v5).f35).length);
            })(_2833_v5))).Elements) {
              let _2978_v102 = _compr_71;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(698)), ((_2979_v5) => function (_2977_i10) {
                return new BigNumber(((_2979_v5).f35).length);
              })(_2833_v5)), _2978_v102)) {
                _coll71.add((_2978_v102).minus(new BigNumber(341)));
              }
            }
            return _coll71;
          }()).length)));
        }
        let _index470 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_2907_cf77).length));
        (_2907_cf77)[_index470] = !(p0);
        let _index471 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_2907_cf77).length));
        (_2907_cf77)[_index471] = p0;
      } else {
        let _2980___mcc_h5 = (_source27).cf73;
        let _2981_cf73 = _2980___mcc_h5;
        let _2982_v103;
        _2982_v103 = _dafny.Seq.of((_this).f18);
        let _2983_v104;
        _2983_v104 = _dafny.Map.Empty.slice().updateUnsafe(_2842_v10,_2842_v10);
        let _2984_v105;
        _2984_v105 = _dafny.MultiSet.fromElements(p0, (_this).f18, !((_this).f18));
        let _2985_v106;
        _2985_v106 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_2984_v105);
        if (!(_dafny.MultiSet.FromArray((((_this).f18) ? (_2982_v103) : (_module.__default.fm44(_2827_v0, (_this).f18, (((_2983_v104).contains(_2842_v10)) ? ((_2983_v104).get(_2842_v10)) : (_2842_v10)), new BigNumber(173), globalState))))).equals((((_2985_v106).contains(p0)) ? ((_2985_v106).get(p0)) : (_2984_v105)))) {
          _2827_v0 = (_2828_v1.f24).multipliedBy(new BigNumber((_2837_v7).cardinality()));
          let _2986_v107;
          let _nw493 = new _module.C5();
          _nw493.__ctor(_dafny.Seq.Concat((_this).f19, (_this).f19), (_this).f19, false);
          _2986_v107 = _nw493;
          r2 = (_this).f18;
          (globalState).f12 = (_2986_v107).f28;
          (globalState).f0 = _module.__default.safeModuloInt(new BigNumber(675), new BigNumber(773));
        } else {
          (_2828_v1).f24 = _2828_v1.f24;
          (_2828_v1).m1((_this).f19, _2828_v1.f24, (_this).f19, globalState);
          let _rhs464 = _dafny.Seq.Concat((((_this).f18) ? (_dafny.Seq.UnicodeFromString("qnpnbbeug")) : ((_this).f19)), _module.__default.fm12(_2828_v1.f24, p0, globalState));
          let _rhs465 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_2828_v1.f24, _2827_v0), _2828_v1.f24);
          let _lhs404 = globalState;
          let _lhs405 = _2828_v1;
          _lhs404.f12 = _rhs464;
          _lhs405.f24 = _rhs465;
          r1 = (_this).f18;
          let _2987_v108;
          let _nw494 = Array((new BigNumber(16)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2987_v108 = _nw494;
          _2987_v108 = _2987_v108;
        }
        let _2988_v109;
        _2988_v109 = _module.D20.create_DC53((_this).f18, (_this).f18, _2842_v10);
        let _2989_v110;
        let _nw495 = Array((new BigNumber(24)).toNumber());
        _nw495[(_dafny.ZERO).toNumber()] = ((_this).f19)[_module.__default.safeIndex(_2828_v1.f24, new BigNumber(((_this).f19).length))];
        _nw495[(_dafny.ONE).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
        _nw495[(new BigNumber(3)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(4)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(5)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(6)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(7)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
        _nw495[(new BigNumber(9)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(10)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(11)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(12)).toNumber()] = ((_this).f19)[_module.__default.safeIndex((_dafny.ZERO).minus(_2827_v0), new BigNumber(((_this).f19).length))];
        _nw495[(new BigNumber(13)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(14)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(15)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(16)).toNumber()] = (((_this).f18) ? (_2842_v10) : (_2842_v10));
        _nw495[(new BigNumber(17)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(18)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(19)).toNumber()] = (_2988_v109).dtor_cf89;
        _nw495[(new BigNumber(20)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(21)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(22)).toNumber()] = _2842_v10;
        _nw495[(new BigNumber(23)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _2989_v110 = _nw495;
        _2989_v110 = _2989_v110;
        let _2990_v111;
        let _nw496 = Array((new BigNumber(4)).toNumber()).fill(false);
        _2990_v111 = _nw496;
        let _2991_v112;
        _2991_v112 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2828_v1.f24);
        let _2992_v113;
        _2992_v113 = _dafny.Map.Empty.slice().updateUnsafe(_2991_v112,_2827_v0);
        let _2993_v114;
        let _nw497 = new _module.C3();
        _nw497.__ctor(_dafny.Seq.of(new BigNumber(917)), _dafny.Seq.UnicodeFromString("x"), (_this).f20, p0);
        _2993_v114 = _nw497;
        let _index472 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_2990_v111).length));
        (_2990_v111)[_index472] = !(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2992_v113).length),_2993_v114)).length)).isEqualTo(_2828_v1.f24);
        let _index473 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_2990_v111).length));
        (_2990_v111)[_index473] = (_2982_v103)[_module.__default.safeIndex(new BigNumber(((_this).f19).length), new BigNumber((_2982_v103).length))];
        _2990_v111 = _2990_v111;
      }
      r0 = true;
      r1 = (_this).f21;
      let _2994_v115;
      _2994_v115 = _dafny.Map.Empty.slice().updateUnsafe(_2828_v1.f24,(_2833_v5).fm2(p0, (_this).f21, p0, globalState));
      r2 = (((_2994_v115).contains(_2828_v1.f24)) ? ((_2994_v115).get(_2828_v1.f24)) : (true));
      return [r0, r1, r2];
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
