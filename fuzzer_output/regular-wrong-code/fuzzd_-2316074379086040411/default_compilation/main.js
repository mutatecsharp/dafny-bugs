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
    static fm0(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(573)), function (_0_i0) {
        return new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality());
      }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),true)).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-708)), _dafny.Seq.of(new BigNumber(224), new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality()))));
    };
    static fm1(p0, globalState) {
      return _dafny.MultiSet.fromElements(((true) ? (new _dafny.CodePoint('c'.codePointAt(0))) : (new _dafny.CodePoint('h'.codePointAt(0)))));
    };
    static fm2(p0, p1, p2, globalState) {
      let _source0 = _module.D5.create_DC14(_module.D5.create_DC13(new _dafny.CodePoint('y'.codePointAt(0)), false, new BigNumber(349), new _dafny.CodePoint('u'.codePointAt(0))));
      if (_source0.is_DC12) {
        let _1___mcc_h0 = (_source0).cf17;
        let _2_cf17 = _1___mcc_h0;
        return _module.D0.create_DC0(new _dafny.CodePoint('d'.codePointAt(0)));
      } else if (_source0.is_DC13) {
        let _3___mcc_h1 = (_source0).cf18;
        let _4___mcc_h2 = (_source0).cf19;
        let _5___mcc_h3 = (_source0).cf20;
        let _6___mcc_h4 = (_source0).cf21;
        let _7_cf21 = _6___mcc_h4;
        let _8_cf20 = _5___mcc_h3;
        let _9_cf19 = _4___mcc_h2;
        let _10_cf18 = _3___mcc_h1;
        return _module.D0.create_DC0(new _dafny.CodePoint('v'.codePointAt(0)));
      } else if (_source0.is_DC11) {
        let _11___mcc_h5 = (_source0).cf16;
        let _12_cf16 = _11___mcc_h5;
        return _module.D0.create_DC0(new _dafny.CodePoint('o'.codePointAt(0)));
      } else {
        let _13___mcc_h6 = (_source0).cf22;
        let _14_cf22 = _13___mcc_h6;
        return _module.D0.create_DC0(new _dafny.CodePoint('y'.codePointAt(0)));
      }
    };
    static fm4(p0, globalState) {
      let _source1 = _module.D5.create_DC11(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), function (_15_i0) {
  return new _dafny.CodePoint('b'.codePointAt(0));
})));
      if (_source1.is_DC12) {
        let _16___mcc_h0 = (_source1).cf17;
        let _17_cf17 = _16___mcc_h0;
        return _17_cf17;
      } else if (_source1.is_DC13) {
        let _18___mcc_h1 = (_source1).cf18;
        let _19___mcc_h2 = (_source1).cf19;
        let _20___mcc_h3 = (_source1).cf20;
        let _21___mcc_h4 = (_source1).cf21;
        let _22_cf21 = _21___mcc_h4;
        let _23_cf20 = _20___mcc_h3;
        let _24_cf19 = _19___mcc_h2;
        let _25_cf18 = _18___mcc_h1;
        return new BigNumber((_dafny.Seq.of(_23_cf20)).length);
      } else if (_source1.is_DC11) {
        let _26___mcc_h5 = (_source1).cf16;
        let _27_cf16 = _26___mcc_h5;
        return new BigNumber(890);
      } else {
        let _28___mcc_h6 = (_source1).cf22;
        let _29_cf22 = _28___mcc_h6;
        return new BigNumber(-862);
      }
    };
    static fm5(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("hab");
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(245),new BigNumber((_dafny.Seq.of(true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(478),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(30)), function (_30_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()),new BigNumber(247))));
    };
    static fm7(p0, p1, globalState) {
      return (function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),new BigNumber(994))).Keys.Elements) {
          let _31_v0 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),new BigNumber(994))).contains(_31_v0)) {
            _coll0.push([_31_v0,new BigNumber((_dafny.Seq.UnicodeFromString("af")).length)]);
          }
        }
        return _coll0;
      }()).Merge((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), function (_32_i0) {
          return _dafny.Seq.of(false);
        }))).Elements) {
          let _33_v1 = _compr_1;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), function (_32_i0) {
            return _dafny.Seq.of(false);
          }))).contains(_33_v1)) {
            _coll1.push([_33_v1,new BigNumber(467)]);
          }
        }
        return _coll1;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(false), false, false, false, true),new BigNumber(461))));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('y'.codePointAt(0));
    };
    static fm9(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(!(false), true, false, false, true), _dafny.Seq.of(false, false)), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(!(false))));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_34_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("fuonoct"))));
    };
    static fm11(p0, p1, p2, globalState) {
      return ((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).Elements) {
          let _35_v0 = _compr_2;
          if ((_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).contains(_35_v0)) {
            _coll2.push([_35_v0,new BigNumber((_dafny.Seq.of(new BigNumber(149), new BigNumber((_dafny.Seq.of(false)).length))).length)]);
          }
        }
        return _coll2;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(46)))).Merge((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),false)).Keys.Elements) {
          let _36_v1 = _compr_3;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),false)).contains(_36_v1)) {
            _coll3.push([_36_v1,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(72))).cardinality())]);
          }
        }
        return _coll3;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber((_dafny.Seq.of(new BigNumber(518))).length))));
    };
    static fm12(globalState) {
      return !(true) || ((new BigNumber(-711)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(841), new BigNumber(902))) {
          let _37_v0 = _compr_4;
          if (((new BigNumber(841)).isLessThanOrEqualTo(_37_v0)) && ((_37_v0).isLessThan(new BigNumber(902)))) {
            _coll4.push([(_37_v0).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-948)), function (_38_i0) {
              return new _dafny.CodePoint('w'.codePointAt(0));
            })).length)]);
          }
        }
        return _coll4;
      }())).length)));
    };
    static fm13(globalState) {
      return (_dafny.MultiSet.fromElements(true, true)).Union(((true) ? (_dafny.MultiSet.fromElements(false)) : (_dafny.MultiSet.fromElements(false, true))));
    };
    static m1(p0, p1, p2, p3, globalState) {
      let r0 = undefined;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _39_v0;
      _39_v0 = _dafny.Set.fromElements(true, _module.__default.fm12(globalState));
      let _40_v1;
      _40_v1 = new BigNumber(952);
      let _41_v4;
      _41_v4 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(327), new BigNumber(247))) {
          let _42_v2 = _compr_5;
          if (((new BigNumber(327)).isLessThanOrEqualTo(_42_v2)) && ((_42_v2).isLessThan(new BigNumber(247)))) {
            _coll5.add((_42_v2).multipliedBy(new BigNumber((function () {
              let _coll6 = new _dafny.Set();
              for (const _compr_6 of (_dafny.MultiSet.FromArray(p3)).Elements) {
                let _43_v3 = _compr_6;
                if ((_dafny.MultiSet.FromArray(p3)).contains(_43_v3)) {
                  _coll6.add(_module.__default.safeDivisionInt(_43_v3, new BigNumber(665)));
                }
              }
              return _coll6;
            }()).length)));
          }
        }
        return _coll5;
      }(),p2);
      let _44_v5;
      _44_v5 = _dafny.Map.Empty.slice().updateUnsafe(_40_v1,new BigNumber((_41_v4).length));
      let _45_i0;
      _45_i0 = _dafny.ZERO;
      L0: {
        while (((_dafny.ZERO).minus((((_44_v5).contains(_40_v1)) ? ((_44_v5).get(_40_v1)) : (_40_v1)))).isLessThan(new BigNumber((_39_v0).length))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_45_i0)) {
              break L0;
            }
            _45_i0 = (_45_i0).plus(_dafny.ONE);
            (globalState).f2 = _40_v1;
            _44_v5 = (_44_v5).update(_40_v1, _module.__default.safeModuloInt(_40_v1, new BigNumber(196)));
            let _46_v6;
            _46_v6 = _module.D4.create_DC9(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(234)), ((_47_v1) => function (_48_i1) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_47_v1)).length);
})(_40_v1))).length), p2);
            let _source2 = ((p2) ? (_46_v6) : (_46_v6));
            if (_source2.is_DC9) {
              let _49___mcc_h0 = (_source2).cf14;
              let _50___mcc_h1 = (_source2).cf15;
              let _51_cf15 = _50___mcc_h1;
              let _52_cf14 = _49___mcc_h0;
              let _53_v7;
              let _nw0 = Array((_dafny.ONE).toNumber());
              _nw0[(_dafny.ZERO).toNumber()] = _52_cf14;
              _53_v7 = _nw0;
              let _54_v8;
              _54_v8 = _dafny.Seq.UnicodeFromString("e");
              let _55_v9;
              _55_v9 = _dafny.MultiSet.fromElements(_54_v8);
              let _56_v10;
              _56_v10 = _module.D5.create_DC11(_55_v9);
              let _index0 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_53_v7).length));
              (_53_v7)[_index0] = new BigNumber(((_55_v9).Union((_56_v10).dtor_cf16)).cardinality());
              let _index1 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_53_v7).length));
              (_53_v7)[_index1] = (((_44_v5).contains(_52_cf14)) ? ((_44_v5).get(_52_cf14)) : (_40_v1));
              (globalState).f10 = p2;
              let _57_v11;
              let _nw1 = new _module.C0();
              _nw1.__ctor();
              _57_v11 = _nw1;
              r0 = _57_v11;
              let _58_v12;
              let _nw2 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _58_v12 = _nw2;
              let _index2 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_58_v12).length));
              (_58_v12)[_index2] = _54_v8;
              let _index3 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_58_v12).length));
              (_58_v12)[_index3] = _dafny.Seq.Concat(_54_v8, _54_v8);
            } else if (_source2.is_DC10) {
              let _59_v13;
              _59_v13 = _dafny.Set.fromElements(_40_v1, _40_v1, _40_v1, _40_v1, _40_v1);
              let _60_v14;
              _60_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_59_v13).length),p2);
              let _61_v15;
              _61_v15 = _module.D0.create_DC1(!((((_60_v14).contains(_40_v1)) ? ((_60_v14).get(_40_v1)) : (p2))), false, p2);
              let _62_v16;
              _62_v16 = _dafny.Seq.UnicodeFromString("tgxywina");
              let _63_v17;
              _63_v17 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
              r3 = (((_61_v15).dtor_cf3) ? (((p2) ? (new BigNumber((_62_v16).length)) : (_40_v1))) : (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,(_46_v6).dtor_cf15)).Merge(_63_v17)).length)));
              (globalState).f9 = new BigNumber(-423);
              let _64_v18;
              _64_v18 = _dafny.Seq.of(p2);
              let _65_v19;
              let _nw3 = Array((new BigNumber(4)).toNumber());
              _nw3[(_dafny.ZERO).toNumber()] = _64_v18;
              _nw3[(_dafny.ONE).toNumber()] = _dafny.Seq.of((_64_v18)[_module.__default.safeIndex((_dafny.ZERO).minus(_40_v1), new BigNumber((_64_v18).length))]);
              _nw3[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(p2, !(p2));
              _nw3[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_64_v18, _64_v18);
              _65_v19 = _nw3;
              r1 = _65_v19;
              let _66_v20;
              let _nw4 = new _module.C0();
              _nw4.__ctor();
              _66_v20 = _nw4;
            } else {
              let _67___mcc_h2 = (_source2).cf13;
              let _68_cf13 = _67___mcc_h2;
              let _69_v21;
              let _nw5 = new _module.C0();
              _nw5.__ctor();
              _69_v21 = _nw5;
              r0 = ((p2) ? (_69_v21) : (_69_v21));
              (globalState).f2 = ((_40_v1).multipliedBy(_40_v1)).multipliedBy((_40_v1).minus(new BigNumber(-153)));
              (globalState).f9 = _40_v1;
              (globalState).f2 = (p3)[_module.__default.safeIndex(_40_v1, new BigNumber((p3).length))];
            }
            let _70_v22;
            _70_v22 = _dafny.Seq.UnicodeFromString("cnyebc");
            _40_v1 = _module.__default.safeModuloInt(new BigNumber((_70_v22).length), (new BigNumber(-808)).plus(new BigNumber(133)));
          }
        }
      }
      let _71_v23;
      let _nw6 = new _module.C0();
      _nw6.__ctor();
      _71_v23 = _nw6;
      let _72_v24;
      let _nw7 = Array((new BigNumber(8)).toNumber());
      _nw7[(_dafny.ZERO).toNumber()] = _71_v23;
      _nw7[(_dafny.ONE).toNumber()] = _71_v23;
      _nw7[(new BigNumber(2)).toNumber()] = _71_v23;
      _nw7[(new BigNumber(3)).toNumber()] = _71_v23;
      _nw7[(new BigNumber(4)).toNumber()] = _71_v23;
      _nw7[(new BigNumber(5)).toNumber()] = _71_v23;
      _nw7[(new BigNumber(6)).toNumber()] = _71_v23;
      _nw7[(new BigNumber(7)).toNumber()] = _71_v23;
      _72_v24 = _nw7;
      let _rhs0 = _40_v1;
      let _rhs1 = _72_v24;
      let _rhs2 = p2;
      let _rhs3 = _40_v1;
      let _rhs4 = _71_v23;
      let _lhs0 = globalState;
      let _lhs1 = globalState;
      r3 = _rhs0;
      _72_v24 = _rhs1;
      _lhs0.f13 = _rhs2;
      _lhs1.f2 = _rhs3;
      _71_v23 = _rhs4;
      let _73_v25;
      _73_v25 = _dafny.Seq.of(p2);
      let _74_v26;
      _74_v26 = _dafny.Map.Empty.slice().updateUnsafe(_73_v25,_40_v1);
      let _hi0 = _40_v1;
      for (let _75_i2 = _module.__default.fm4(_74_v26, globalState); _75_i2.isLessThan(_hi0); _75_i2 = _75_i2.plus(_dafny.ONE)) {
        let _76_v27;
        let _nw8 = new _module.C0();
        _nw8.__ctor();
        _76_v27 = _nw8;
        (globalState).f10 = p2;
        let _77_v28;
        _77_v28 = _dafny.Seq.UnicodeFromString("dw");
        let _78_v29;
        _78_v29 = _dafny.Map.Empty.slice().updateUnsafe(_40_v1,_77_v28);
        let _79_v30;
        let _nw9 = Array((new BigNumber(6)).toNumber());
        _nw9[(_dafny.ZERO).toNumber()] = ((p2) ? (_77_v28) : (_77_v28));
        _nw9[(_dafny.ONE).toNumber()] = _77_v28;
        _nw9[(new BigNumber(2)).toNumber()] = (((_78_v29).contains(new BigNumber(619))) ? ((_78_v29).get(new BigNumber(619))) : (_77_v28));
        _nw9[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_77_v28, _77_v28);
        _nw9[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("hicn");
        _nw9[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("etaeup"), _77_v28);
        _79_v30 = _nw9;
        _79_v30 = _79_v30;
        let _80_v31;
        _80_v31 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _81_v32;
        _81_v32 = _dafny.Map.Empty.slice().updateUnsafe((((_80_v31).contains(p2)) ? ((_80_v31).get(p2)) : (p2)),p2);
        let _82_v33;
        _82_v33 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.fm10(p2, p2, _40_v1, p0, globalState));
        (globalState).f10 = (((_80_v31).contains(true)) ? ((_80_v31).get(true)) : ((_dafny.Set.fromElements(p2, p2, p2, p2, p2)).IsSubsetOf((((_82_v33).contains(p2)) ? ((_82_v33).get(p2)) : (_dafny.Set.fromElements(p2, p2))))));
      }
      let _83_v34;
      _83_v34 = new _dafny.CodePoint('s'.codePointAt(0));
      _83_v34 = p0;
      let _84_v35;
      let _init0 = ((_85_v1) => function (_86_i4) {
        return (_86_i4).minus(_85_v1);
      })(_40_v1);
      let _nw10 = Array((new BigNumber(23)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw10.length); _i0_0++) {
        _nw10[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _84_v35 = _nw10;
      let _87_v36;
      _87_v36 = _module.D1.create_DC3(p3, p2);
      let _88_v37;
      _88_v37 = _dafny.Map.Empty.slice().updateUnsafe(_84_v35,(_87_v36).dtor_cf5);
      let _89_v38;
      _89_v38 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(-464));
      let _90_v39;
      _90_v39 = _dafny.Seq.of(_89_v38);
      let _91_v40;
      let _nw11 = Array((new BigNumber(28)).toNumber());
      _nw11[(_dafny.ZERO).toNumber()] = p2;
      _nw11[(_dafny.ONE).toNumber()] = p2;
      _nw11[(new BigNumber(2)).toNumber()] = !(p2);
      _nw11[(new BigNumber(3)).toNumber()] = p2;
      _nw11[(new BigNumber(4)).toNumber()] = _module.__default.fm12(globalState);
      _nw11[(new BigNumber(5)).toNumber()] = false;
      _nw11[(new BigNumber(6)).toNumber()] = false;
      _nw11[(new BigNumber(7)).toNumber()] = p2;
      _nw11[(new BigNumber(8)).toNumber()] = !(p2);
      _nw11[(new BigNumber(9)).toNumber()] = true;
      _nw11[(new BigNumber(10)).toNumber()] = _module.__default.fm12(globalState);
      _nw11[(new BigNumber(11)).toNumber()] = (_88_v37).equals(_dafny.Map.Empty.slice().updateUnsafe(_84_v35,p3));
      _nw11[(new BigNumber(12)).toNumber()] = p2;
      _nw11[(new BigNumber(13)).toNumber()] = p2;
      _nw11[(new BigNumber(14)).toNumber()] = (_71_v23).fm3(!(p2), p2, _module.__default.fm13(globalState), _73_v25, globalState);
      _nw11[(new BigNumber(15)).toNumber()] = p2;
      _nw11[(new BigNumber(16)).toNumber()] = p2;
      _nw11[(new BigNumber(17)).toNumber()] = p2;
      _nw11[(new BigNumber(18)).toNumber()] = p2;
      _nw11[(new BigNumber(19)).toNumber()] = true;
      _nw11[(new BigNumber(20)).toNumber()] = p2;
      _nw11[(new BigNumber(21)).toNumber()] = p2;
      _nw11[(new BigNumber(22)).toNumber()] = _dafny.Seq.IsPrefixOf(_73_v25, _dafny.Seq.of(p2, p2));
      _nw11[(new BigNumber(23)).toNumber()] = p2;
      _nw11[(new BigNumber(24)).toNumber()] = p2;
      _nw11[(new BigNumber(25)).toNumber()] = _dafny.Seq.IsPrefixOf(_90_v39, _90_v39);
      _nw11[(new BigNumber(26)).toNumber()] = p2;
      _nw11[(new BigNumber(27)).toNumber()] = (_40_v1).isLessThan(_module.__default.fm4(_74_v26, globalState));
      _91_v40 = _nw11;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_91_v40).length))) {
        let _92_i3 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_92_i3)) && ((_92_i3).isLessThan(new BigNumber((_91_v40).length))))) {
          (_91_v40)[(_92_i3)] = (_dafny.MultiSet.fromElements(_40_v1)).IsSubsetOf(((true) ? (_dafny.MultiSet.fromElements(_40_v1)) : (_dafny.MultiSet.fromElements(_40_v1, new BigNumber((_dafny.Seq.of(_40_v1)).length), _40_v1, _40_v1, _40_v1))));
        }
      }
      let _93_v41;
      let _nw12 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
      _93_v41 = _nw12;
      let _index4 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_93_v41).length));
      (_93_v41)[_index4] = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(842));
      let _index5 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_93_v41).length));
      (_93_v41)[_index5] = (_89_v38).Merge((_89_v38).Merge(_89_v38));
      let _nw13 = new _module.C0();
      _nw13.__ctor();
      r0 = _nw13;
      let _94_v42;
      let _nw14 = Array((new BigNumber(14)).toNumber());
      _nw14[(_dafny.ZERO).toNumber()] = _73_v25;
      _nw14[(_dafny.ONE).toNumber()] = _dafny.Seq.of(p2);
      _nw14[(new BigNumber(2)).toNumber()] = _73_v25;
      _nw14[(new BigNumber(3)).toNumber()] = _73_v25;
      _nw14[(new BigNumber(4)).toNumber()] = _73_v25;
      _nw14[(new BigNumber(5)).toNumber()] = _module.__default.fm9(globalState);
      _nw14[(new BigNumber(6)).toNumber()] = _73_v25;
      _nw14[(new BigNumber(7)).toNumber()] = _73_v25;
      _nw14[(new BigNumber(8)).toNumber()] = _73_v25;
      _nw14[(new BigNumber(9)).toNumber()] = _module.__default.fm9(globalState);
      _nw14[(new BigNumber(10)).toNumber()] = _73_v25;
      _nw14[(new BigNumber(11)).toNumber()] = _73_v25;
      _nw14[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_73_v25, _module.__default.safeIndex((_dafny.ZERO).minus(_40_v1), new BigNumber((_73_v25).length)), _module.__default.fm12(globalState)), _module.__default.safeIndex(_40_v1, new BigNumber((_dafny.Seq.update(_73_v25, _module.__default.safeIndex((_dafny.ZERO).minus(_40_v1), new BigNumber((_73_v25).length)), _module.__default.fm12(globalState))).length)), !(p2));
      _nw14[(new BigNumber(13)).toNumber()] = _73_v25;
      _94_v42 = _nw14;
      let _95_v43;
      _95_v43 = _dafny.Map.Empty.slice().updateUnsafe(_71_v23,_94_v42);
      r1 = (((_95_v43).contains(_71_v23)) ? ((_95_v43).get(_71_v23)) : (_94_v42));
      r2 = _module.__default.safeModuloInt(_40_v1, _40_v1);
      r3 = _40_v1;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _96_v0;
      let _nw15 = Array((new BigNumber(20)).toNumber()).fill(false);
      _96_v0 = _nw15;
      let _97_v1;
      _97_v1 = _dafny.Seq.UnicodeFromString("uwjfy");
      let _98_globalState;
      let _nw16 = new _module.GlobalState();
      _nw16.__ctor(new BigNumber(511), new _dafny.CodePoint('o'.codePointAt(0)), new BigNumber(-937), new BigNumber(29), _96_v0, false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), function (_99_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(402)), function (_100_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }), new _dafny.CodePoint('y'.codePointAt(0)), new BigNumber(949), false, new _dafny.CodePoint('j'.codePointAt(0)), new BigNumber(820), true, _dafny.Seq.Concat(_97_v1, _97_v1), false);
      _98_globalState = _nw16;
      let _101_v2;
      _101_v2 = false;
      let _102_v3;
      _102_v3 = new BigNumber(397);
      if (((_101_v2) === (_101_v2)) === (!((_102_v3).isLessThanOrEqualTo(_102_v3)))) {
        let _103_v4;
        _103_v4 = _dafny.Seq.of(_102_v3, new BigNumber(92), _102_v3, _102_v3, new BigNumber((_dafny.MultiSet.fromElements(_101_v2)).cardinality()));
        (_98_globalState).f0 = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm0(_103_v4, _module.__default.safeModuloInt(_102_v3, _102_v3), _98_globalState)).length));
        let _104_v5;
        _104_v5 = _dafny.MultiSet.fromElements(new BigNumber(-498), _102_v3);
        (_98_globalState).f10 = !(((_104_v5).Union(_104_v5)).IsSubsetOf(_104_v5));
        if (_101_v2) {
          (_98_globalState).f7 = _97_v1;
          let _105_v6;
          _105_v6 = _dafny.Map.Empty.slice().updateUnsafe(((_101_v2) ? (new BigNumber(241)) : (new BigNumber((_104_v5).cardinality()))),_module.__default.fm1(_102_v3, _98_globalState));
          let _106_v7;
          _106_v7 = new _dafny.CodePoint('a'.codePointAt(0));
          let _107_v8;
          _107_v8 = _dafny.MultiSet.fromElements((_module.__default.fm2(_101_v2, _101_v2, _102_v3, _98_globalState)).dtor_cf0, _106_v7, new _dafny.CodePoint('t'.codePointAt(0)), _106_v7);
          _105_v6 = (_105_v6).update(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), function (_108_i2) {
            return new BigNumber(936);
          })).length), _107_v8);
          let _109_v9;
          let _nw17 = new _module.C0();
          _nw17.__ctor();
          _109_v9 = _nw17;
          let _110_v10;
          let _nw18 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _110_v10 = _nw18;
          let _111_v11;
          _111_v11 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_110_v10);
          let _112_v12;
          _112_v12 = _dafny.Seq.of(_110_v10, (((_111_v11).contains(_101_v2)) ? ((_111_v11).get(_101_v2)) : (_110_v10)), _110_v10);
          (_98_globalState).f9 = new BigNumber((_112_v12).length);
          let _113_v13;
          let _nw19 = new _module.C0();
          _nw19.__ctor();
          _113_v13 = _nw19;
        } else {
          let _114_v14;
          _114_v14 = new _dafny.CodePoint('l'.codePointAt(0));
          let _115_v15;
          _115_v15 = _module.D0.create_DC0(_114_v14);
          let _116_v16;
          _116_v16 = _dafny.Seq.of(_115_v15);
          let _117_v17;
          let _118_v18;
          let _119_v19;
          let _120_v20;
          let _out0;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = _module.__default.m1(new _dafny.CodePoint('j'.codePointAt(0)), _116_v16, _101_v2, _dafny.Seq.Concat(_103_v4, _dafny.Seq.of(_102_v3)), _98_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _out3 = _outcollector0[3];
          _117_v17 = _out0;
          _118_v18 = _out1;
          _119_v19 = _out2;
          _120_v20 = _out3;
          (_98_globalState).f10 = _101_v2;
          let _121_v21;
          _121_v21 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),_101_v2);
          let _122_v22;
          _122_v22 = _dafny.Set.fromElements(_121_v21);
          let _123_v23;
          _123_v23 = _dafny.Seq.of(_97_v1, _dafny.Seq.UnicodeFromString("ckdgwau"), _dafny.Seq.UnicodeFromString("vu"));
          let _124_v24;
          _124_v24 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_123_v23);
          let _125_v25;
          let _nw20 = Array((new BigNumber(29)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = new BigNumber((_122_v22).length);
          _nw20[(_dafny.ONE).toNumber()] = _119_v19;
          _nw20[(new BigNumber(2)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(3)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_120_v20);
          _nw20[(new BigNumber(5)).toNumber()] = _102_v3;
          _nw20[(new BigNumber(6)).toNumber()] = new BigNumber((_97_v1).length);
          _nw20[(new BigNumber(7)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(8)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(9)).toNumber()] = _102_v3;
          _nw20[(new BigNumber(10)).toNumber()] = _120_v20;
          _nw20[(new BigNumber(11)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(12)).toNumber()] = _102_v3;
          _nw20[(new BigNumber(13)).toNumber()] = _120_v20;
          _nw20[(new BigNumber(14)).toNumber()] = new BigNumber(597);
          _nw20[(new BigNumber(15)).toNumber()] = new BigNumber(((((_124_v24).contains(_101_v2)) ? ((_124_v24).get(_101_v2)) : (_123_v23))).length);
          _nw20[(new BigNumber(16)).toNumber()] = _102_v3;
          _nw20[(new BigNumber(17)).toNumber()] = _102_v3;
          _nw20[(new BigNumber(18)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(_120_v20);
          _nw20[(new BigNumber(20)).toNumber()] = _102_v3;
          _nw20[(new BigNumber(21)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(22)).toNumber()] = _120_v20;
          _nw20[(new BigNumber(23)).toNumber()] = _102_v3;
          _nw20[(new BigNumber(24)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(25)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_101_v2)).cardinality());
          _nw20[(new BigNumber(26)).toNumber()] = _119_v19;
          _nw20[(new BigNumber(27)).toNumber()] = _120_v20;
          _nw20[(new BigNumber(28)).toNumber()] = new BigNumber(291);
          _125_v25 = _nw20;
          let _126_v26;
          let _nw21 = Array((new BigNumber(26)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = _125_v25;
          _nw21[(_dafny.ONE).toNumber()] = _125_v25;
          _nw21[(new BigNumber(2)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(3)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(4)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(5)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(6)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(7)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(8)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(9)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(10)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(11)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(12)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(13)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(14)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(15)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(16)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(17)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(18)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(19)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(20)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(21)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(22)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(23)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(24)).toNumber()] = _125_v25;
          _nw21[(new BigNumber(25)).toNumber()] = _125_v25;
          _126_v26 = _nw21;
          let _index6 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_126_v26).length));
          (_126_v26)[_index6] = _125_v25;
          let _index7 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_126_v26).length));
          (_126_v26)[_index7] = _125_v25;
          let _127_v27;
          let _nw22 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _127_v27 = _nw22;
          let _nw23 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _127_v27 = _nw23;
          let _128_v28;
          _128_v28 = _dafny.Seq.of(_101_v2);
          (_98_globalState).f0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_97_v1).length), new BigNumber((_128_v28).length))));
        }
        let _129_v29;
        _129_v29 = new _dafny.CodePoint('d'.codePointAt(0));
        let _130_v30;
        _130_v30 = _module.D0.create_DC0(_129_v29);
        let _131_v31;
        _131_v31 = _dafny.Seq.of(_130_v30, _130_v30, _module.D0.create_DC0(_129_v29));
        let _132_v32;
        _132_v32 = _dafny.Map.Empty.slice().updateUnsafe(_102_v3,_102_v3);
        let _133_v33;
        let _134_v34;
        let _135_v35;
        let _136_v36;
        let _out4;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector1 = _module.__default.m1(_129_v29, _131_v31, _101_v2, _dafny.Seq.of(new BigNumber((_132_v32).length), _102_v3), _98_globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _out7 = _outcollector1[3];
        _133_v33 = _out4;
        _134_v34 = _out5;
        _135_v35 = _out6;
        _136_v36 = _out7;
        let _137_v37;
        let _init1 = ((_138_v29) => function (_139_i3) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), ((_140_v29) => function (_141_i4) {
            return _140_v29;
          })(_138_v29));
        })(_129_v29);
        let _nw24 = Array((_dafny.ONE).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw24.length); _i0_1++) {
          _nw24[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _137_v37 = _nw24;
        let _index8 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_137_v37).length));
        (_137_v37)[_index8] = _97_v1;
        let _142_v38;
        _142_v38 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("alnbjixe"), _97_v1);
        let _143_v39;
        _143_v39 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_101_v2);
        let _144_v40;
        _144_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_143_v39).length),_97_v1);
        let _145_v41;
        _145_v41 = _dafny.Seq.of(_101_v2, _101_v2);
        let _index9 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_137_v37).length));
        let _rhs5 = _dafny.Seq.IsPrefixOf((_142_v38)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_142_v38).length))], _dafny.Seq.Concat(_module.__default.fm5(_101_v2, _101_v2, _98_globalState), _dafny.Seq.update(_97_v1, _module.__default.safeIndex(new BigNumber((_104_v5).cardinality()), new BigNumber((_97_v1).length)), _129_v29)));
        let _rhs6 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-282)), ((_146_v29) => function (_147_i5) {
          return _146_v29;
        })(_129_v29)), (((_144_v40).contains((_dafny.ZERO).minus(new BigNumber((_145_v41).length)))) ? ((_144_v40).get((_dafny.ZERO).minus(new BigNumber((_145_v41).length)))) : (_97_v1)));
        let _rhs7 = _97_v1;
        let _lhs2 = _98_globalState;
        let _lhs3 = _98_globalState;
        let _lhs4 = _137_v37;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_137_v37).length));
        _lhs2.f15 = _rhs5;
        _lhs3.f7 = _rhs6;
        _lhs4[_lhs5] = _rhs7;
      } else {
        let _148_v42;
        _148_v42 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,new BigNumber(726));
        let _149_v43;
        _149_v43 = _dafny.Seq.of((((_148_v42).contains(_101_v2)) ? ((_148_v42).get(_101_v2)) : (new BigNumber(-708))), new BigNumber((_97_v1).length), new BigNumber(-797), _102_v3);
        let _150_v44;
        _150_v44 = _dafny.Seq.of(_149_v43, _149_v43, _dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), ((_151_v3) => function (_152_i6) {
          return _151_v3;
        })(_102_v3)));
        _150_v44 = _150_v44;
        let _index10 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_96_v0).length));
        (_96_v0)[_index10] = _101_v2;
        let _index11 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_96_v0).length));
        (_96_v0)[_index11] = true;
        let _153_v45;
        _153_v45 = _module.D1.create_DC4((_96_v0)[_module.__default.safeIndex(new BigNumber(47), new BigNumber((_96_v0).length))], _102_v3, _97_v1);
        let _154_v46;
        _154_v46 = _dafny.Set.fromElements(_153_v45, _153_v45, _153_v45);
        _154_v46 = _154_v46;
        let _155_v47;
        _155_v47 = _dafny.MultiSet.fromElements(_101_v2);
        (_98_globalState).f0 = (new BigNumber(((_155_v47).Union(_155_v47)).cardinality())).multipliedBy(_102_v3);
        let _156_v48;
        _156_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_148_v42).length),_102_v3);
        _156_v48 = (_156_v48).update(_102_v3, _102_v3);
      }
      if (_101_v2) {
        (_98_globalState).f13 = (((false) ? (_101_v2) : (_101_v2))) === (false);
        let _157_v49;
        let _nw25 = new _module.C0();
        _nw25.__ctor();
        _157_v49 = _nw25;
        let _158_v50;
        _158_v50 = new _dafny.CodePoint('f'.codePointAt(0));
        _158_v50 = _158_v50;
        let _159_v51;
        _159_v51 = _dafny.Set.fromElements(_101_v2);
        let _160_v52;
        _160_v52 = _dafny.Seq.of(_102_v3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_101_v2,_102_v3)).length), new BigNumber((_159_v51).length), new BigNumber(298));
        let _161_v53;
        _161_v53 = _module.D1.create_DC3(_160_v52, _101_v2);
        if (!((_161_v53).dtor_cf6)) {
          let _162_v54;
          _162_v54 = _dafny.Seq.of(true);
          (_98_globalState).f10 = ((_102_v3).multipliedBy(_102_v3)).isEqualTo(_module.__default.fm4(_dafny.Map.Empty.slice().updateUnsafe(_162_v54,new BigNumber(904)), _98_globalState));
          (_98_globalState).f10 = _101_v2;
          let _163_v55;
          let _nw26 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
          _163_v55 = _nw26;
          let _164_v56;
          _164_v56 = _dafny.Seq.of(_157_v49);
          let _165_v57;
          _165_v57 = _dafny.Map.Empty.slice().updateUnsafe(_102_v3,new BigNumber((_164_v56).length));
          let _166_v58;
          _166_v58 = _module.D1.create_DC4(_101_v2, _102_v3, _97_v1);
          let _167_v59;
          _167_v59 = _dafny.Seq.of(_165_v57);
          let _168_v60;
          _168_v60 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_157_v49);
          let _169_v61;
          _169_v61 = _165_v57;
          let _170_v62;
          _170_v62 = (_169_v61);
          let _nw27 = Array((new BigNumber(17)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _165_v57;
          _nw27[(_dafny.ONE).toNumber()] = _165_v57;
          _nw27[(new BigNumber(2)).toNumber()] = _165_v57;
          _nw27[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_102_v3,(_166_v58).dtor_cf8);
          _nw27[(new BigNumber(4)).toNumber()] = ((_165_v57).update(_102_v3, _102_v3)).Merge(_165_v57);
          _nw27[(new BigNumber(5)).toNumber()] = _module.__default.fm6(false, _101_v2, _102_v3, new BigNumber(702), _98_globalState);
          _nw27[(new BigNumber(6)).toNumber()] = (_165_v57).Merge((_167_v59)[_module.__default.safeIndex(new BigNumber(-488), new BigNumber((_167_v59).length))]);
          _nw27[(new BigNumber(7)).toNumber()] = _165_v57;
          _nw27[(new BigNumber(8)).toNumber()] = (_165_v57).update(_102_v3, new BigNumber((_168_v60).length));
          _nw27[(new BigNumber(9)).toNumber()] = (_169_v61);
          _nw27[(new BigNumber(10)).toNumber()] = _165_v57;
          _nw27[(new BigNumber(11)).toNumber()] = function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-838), new BigNumber(-263))) {
              let _171_v63 = _compr_7;
              if (((new BigNumber(-838)).isLessThanOrEqualTo(_171_v63)) && ((_171_v63).isLessThan(new BigNumber(-263)))) {
                _coll7.push([(_171_v63).minus(_102_v3),_102_v3]);
              }
            }
            return _coll7;
          }();
          _nw27[(new BigNumber(12)).toNumber()] = _165_v57;
          _nw27[(new BigNumber(13)).toNumber()] = (_165_v57).update(_102_v3, _102_v3);
          _nw27[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(446),_102_v3);
          _nw27[(new BigNumber(15)).toNumber()] = _165_v57;
          _nw27[(new BigNumber(16)).toNumber()] = function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_160_v52).Elements) {
              let _172_v64 = _compr_8;
              if (_dafny.Seq.contains(_160_v52, _172_v64)) {
                _coll8.push([_module.__default.safeModuloInt(_172_v64, _102_v3),new BigNumber(991)]);
              }
            }
            return _coll8;
          }();
          _163_v55 = _nw27;
          let _173_v65;
          _173_v65 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,new BigNumber((_dafny.Seq.update(_160_v52, _module.__default.safeIndex(_102_v3, new BigNumber((_160_v52).length)), _102_v3)).length));
          let _174_v66;
          _174_v66 = _dafny.MultiSet.fromElements(_101_v2);
          _173_v65 = (_173_v65).update(_dafny.Seq.IsPrefixOf(_97_v1, _97_v1), ((!((_157_v49).fm3(_101_v2, false, _174_v66, _dafny.Seq.of(_101_v2), _98_globalState))) ? (_102_v3) : (_102_v3)));
          let _index12 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_96_v0).length));
          (_96_v0)[_index12] = _101_v2;
          let _index13 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_96_v0).length));
          (_96_v0)[_index13] = false;
        } else {
          _96_v0 = _96_v0;
          let _175_v67;
          let _nw28 = new _module.C0();
          _nw28.__ctor();
          _175_v67 = _nw28;
          (_98_globalState).f0 = _102_v3;
          (_98_globalState).f10 = _101_v2;
          _102_v3 = (_dafny.ZERO).minus((_dafny.ZERO).minus(((_101_v2) ? (_102_v3) : (new BigNumber(421)))));
        }
        (_98_globalState).f15 = _101_v2;
      } else {
        (_98_globalState).f2 = _102_v3;
        let _176_v68;
        _176_v68 = _dafny.Seq.of(true);
        let _177_v69;
        _177_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(_101_v2, _101_v2), _176_v68),_101_v2);
        _177_v69 = _177_v69;
        let _178_v70;
        _178_v70 = _dafny.MultiSet.fromElements(_101_v2, _101_v2);
        let _179_v71;
        let _nw29 = Array((new BigNumber(4)).toNumber());
        _nw29[(_dafny.ZERO).toNumber()] = _178_v70;
        _nw29[(_dafny.ONE).toNumber()] = (_178_v70).Difference(_178_v70);
        _nw29[(new BigNumber(2)).toNumber()] = _178_v70;
        _nw29[(new BigNumber(3)).toNumber()] = (_178_v70).update(_101_v2, _module.__default.abs(_102_v3));
        _179_v71 = _nw29;
        let _index14 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_179_v71).length));
        (_179_v71)[_index14] = (_178_v70).Intersect(_178_v70);
        let _index15 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_179_v71).length));
        (_179_v71)[_index15] = (_dafny.MultiSet.fromElements(_101_v2)).Union(_178_v70);
        let _180_v72;
        let _nw30 = new _module.C0();
        _nw30.__ctor();
        _180_v72 = _nw30;
        let _181_v73;
        _181_v73 = _180_v72;
        _180_v72 = ((_101_v2) ? (_180_v72) : (((_101_v2) ? ((_181_v73)) : (_180_v72))));
        if (_101_v2) {
          (_180_v72).m0(_176_v68, _102_v3, _98_globalState);
          let _182_v74;
          let _init2 = function (_183_i7) {
            return _module.__default.safeDivisionInt(_183_i7, new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('y'.codePointAt(0)))).length));
          };
          let _nw31 = Array((new BigNumber(18)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw31.length); _i0_2++) {
            _nw31[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _182_v74 = _nw31;
          let _184_v75;
          _184_v75 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5(_101_v2, true, _98_globalState),_182_v74);
          let _185_v76;
          _185_v76 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_97_v1,_182_v74), (_184_v75).update(_97_v1, _182_v74), _184_v75);
          let _186_v77;
          _186_v77 = _dafny.Seq.of(_102_v3, new BigNumber(510), new BigNumber(-366));
          let _187_v78;
          _187_v78 = _dafny.Map.Empty.slice().updateUnsafe((_185_v76)[_module.__default.safeIndex(new BigNumber((_97_v1).length), new BigNumber((_185_v76).length))],_module.__default.safeDivisionInt(_102_v3, (_186_v77)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_101_v2, _101_v2)).length)), new BigNumber((_186_v77).length))]));
          _187_v78 = (_187_v78).update((_184_v75).Merge(_184_v75), new BigNumber(-798));
          (_98_globalState).f13 = _101_v2;
          let _index16 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_182_v74).length));
          (_182_v74)[_index16] = _102_v3;
          let _188_v79;
          _188_v79 = _dafny.MultiSet.fromElements(_102_v3, _102_v3);
          let _189_v80;
          _189_v80 = _dafny.Map.Empty.slice().updateUnsafe(_180_v72,_188_v79);
          let _index17 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_182_v74).length));
          (_182_v74)[_index17] = (new BigNumber(618)).plus((new BigNumber((_189_v80).length)).multipliedBy(_102_v3));
          let _190_v81;
          _190_v81 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_101_v2, _101_v2))).cardinality()));
          let _191_v82;
          _191_v82 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,(_182_v74)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_182_v74).length))]);
          let _192_v83;
          _192_v83 = _dafny.Map.Empty.slice().updateUnsafe((_190_v81).IsSubsetOf(_190_v81),(_97_v1)[_module.__default.safeIndex(_module.__default.fm4(_module.__default.fm7(new BigNumber((_module.__default.fm5(true, false, _98_globalState)).length), new BigNumber((_191_v82).length), _98_globalState), _98_globalState), new BigNumber((_97_v1).length))]);
          let _193_v84;
          let _init3 = function (_194_i8) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          };
          let _nw32 = Array((new BigNumber(16)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw32.length); _i0_3++) {
            _nw32[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _193_v84 = _nw32;
          let _195_v85;
          _195_v85 = _dafny.MultiSet.fromElements(_193_v84);
          let _196_v86;
          _196_v86 = _dafny.Seq.of(_195_v85, _195_v85);
          let _197_v87;
          _197_v87 = _dafny.Seq.of(_193_v84);
          let _198_v88;
          _198_v88 = new _dafny.CodePoint('s'.codePointAt(0));
          _192_v83 = (_192_v83).update(((_196_v86)[_module.__default.safeIndex(new BigNumber(976), new BigNumber((_196_v86).length))]).IsDisjointFrom(_dafny.MultiSet.fromElements((_197_v87)[_module.__default.safeIndex(_102_v3, new BigNumber((_197_v87).length))], _193_v84, _193_v84)), _198_v88);
        } else {
          let _199_v89;
          _199_v89 = _dafny.Seq.of(_102_v3);
          let _200_v90;
          _200_v90 = _dafny.MultiSet.fromElements(_102_v3);
          let _201_v91;
          _201_v91 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_101_v2),_102_v3);
          let _202_v92;
          _202_v92 = new _dafny.CodePoint('r'.codePointAt(0));
          let _203_v93;
          _203_v93 = _dafny.MultiSet.fromElements(_module.__default.fm8(_199_v89, (_200_v90).update(new BigNumber(-538), _module.__default.abs(new BigNumber((_dafny.MultiSet.fromElements(_102_v3, _102_v3, _102_v3, _102_v3)).cardinality()))), _101_v2, _module.__default.fm4(_201_v91, _98_globalState), _98_globalState), _202_v92, _202_v92, _202_v92);
          let _204_v94;
          _204_v94 = _dafny.Seq.of(_102_v3, new BigNumber((_203_v93).cardinality()), _module.__default.fm4(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),_102_v3), _98_globalState), _102_v3);
          _102_v3 = (_dafny.ZERO).minus((_204_v94)[_module.__default.safeIndex((_dafny.ZERO).minus(_102_v3), new BigNumber((_204_v94).length))]);
          let _205_v95;
          _205_v95 = _dafny.Set.fromElements(_101_v2);
          let _206_v96;
          _206_v96 = _dafny.Seq.of(_205_v95, _205_v95, _205_v95, _205_v95);
          let _207_v97;
          _207_v97 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(((_206_v96)[_module.__default.safeIndex(_102_v3, new BigNumber((_206_v96).length))]).length)),_180_v72);
          let _208_v98;
          let _209_v99;
          let _210_v100;
          let _211_v101;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m1(_202_v92, _dafny.Seq.Create(_module.__default.abs(new BigNumber(981)), ((_212_v92) => function (_213_i9) {
            return _module.D0.create_DC0(_212_v92);
          })(_202_v92)), _101_v2, _dafny.Seq.update(_199_v89, _module.__default.safeIndex(_102_v3, new BigNumber((_199_v89).length)), new BigNumber((_207_v97).length)), _98_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _208_v98 = _out8;
          _209_v99 = _out9;
          _210_v100 = _out10;
          _211_v101 = _out11;
          _102_v3 = _module.__default.fm4(_module.__default.fm7(_210_v100, _211_v101, _98_globalState), _98_globalState);
          let _214_v102;
          _214_v102 = _dafny.Seq.of(_176_v68, _module.__default.fm9(_98_globalState));
          let _215_v103;
          _215_v103 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_101_v2);
          let _216_v104;
          _216_v104 = _dafny.Seq.of(_215_v103);
          (_180_v72).m0((_214_v102)[_module.__default.safeIndex(new BigNumber((_216_v104).length), new BigNumber((_214_v102).length))], _module.__default.fm4(_201_v91, _98_globalState), _98_globalState);
          let _217_v105;
          _217_v105 = _module.D0.create_DC0(_202_v92);
          let _218_v106;
          _218_v106 = _dafny.Seq.of(_217_v105);
          let _219_v108;
          let _220_v109;
          let _221_v110;
          let _222_v111;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = _module.__default.m1(_202_v92, _218_v106, _101_v2, _dafny.Seq.of((_dafny.ZERO).minus(_211_v101), _module.__default.fm4(function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of (_dafny.Seq.of(_176_v68)).Elements) {
              let _223_v107 = _compr_9;
              if (_dafny.Seq.contains(_dafny.Seq.of(_176_v68), _223_v107)) {
                _coll9.push([_223_v107,new BigNumber((_205_v95).length)]);
              }
            }
            return _coll9;
          }(), _98_globalState)), _98_globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _219_v108 = _out12;
          _220_v109 = _out13;
          _221_v110 = _out14;
          _222_v111 = _out15;
        }
      }
      let _224_v112;
      let _nw33 = Array((new BigNumber(17)).toNumber());
      _224_v112 = _nw33;
      let _225_v113;
      _225_v113 = _dafny.Seq.of(_224_v112);
      (_98_globalState).f2 = new BigNumber((_225_v113).length);
      let _226_v114;
      let _nw34 = new _module.C0();
      _nw34.__ctor();
      _226_v114 = _nw34;
      if (_101_v2) {
        let _227_v115;
        let _nw35 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _227_v115 = _nw35;
        let _index18 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length));
        (_227_v115)[_index18] = _102_v3;
        let _index19 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length));
        (_227_v115)[_index19] = _module.__default.safeDivisionInt(_102_v3, _102_v3);
        if (!(_101_v2) || ((!(_101_v2)) && (_101_v2))) {
          let _228_v116;
          let _nw36 = Array((new BigNumber(9)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = _96_v0;
          _nw36[(_dafny.ONE).toNumber()] = _96_v0;
          _nw36[(new BigNumber(2)).toNumber()] = _96_v0;
          _nw36[(new BigNumber(3)).toNumber()] = _96_v0;
          _nw36[(new BigNumber(4)).toNumber()] = _96_v0;
          _nw36[(new BigNumber(5)).toNumber()] = _96_v0;
          _nw36[(new BigNumber(6)).toNumber()] = _96_v0;
          _nw36[(new BigNumber(7)).toNumber()] = _96_v0;
          _nw36[(new BigNumber(8)).toNumber()] = _96_v0;
          _228_v116 = _nw36;
          let _229_v117;
          _229_v117 = _dafny.Map.Empty.slice().updateUnsafe(_228_v116,false);
          let _230_v118;
          _230_v118 = _dafny.MultiSet.fromElements(_101_v2);
          let _231_v119;
          _231_v119 = _dafny.Seq.of(true, _101_v2, _101_v2);
          _229_v117 = (_229_v117).update(_228_v116, !((_226_v114).fm3(_101_v2, _101_v2, _230_v118, _231_v119, _98_globalState)));
          (_98_globalState).f10 = !((_101_v2) && (_101_v2));
          (_226_v114).m0(_231_v119, _102_v3, _98_globalState);
          let _index20 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length));
          (_227_v115)[_index20] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber(142), new BigNumber(898)), (_227_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length))]));
          (_98_globalState).f2 = (_dafny.ZERO).minus(((_101_v2) ? ((_dafny.ZERO).minus((_227_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length))])) : (_102_v3)));
        } else {
          _97_v1 = _dafny.Seq.UnicodeFromString("jkgtl");
          _226_v114 = _226_v114;
          let _index21 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length));
          (_227_v115)[_index21] = (_227_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length))];
          let _232_v120;
          _232_v120 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_101_v2);
          let _233_v121;
          _233_v121 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_232_v120);
          let _234_v122;
          _234_v122 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_232_v120, (((_233_v121).contains(true)) ? ((_233_v121).get(true)) : (_232_v120))),_101_v2);
          let _235_v123;
          _235_v123 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_101_v2,_101_v2), _232_v120, _232_v120);
          (_98_globalState).f13 = (((_234_v122).contains(_235_v123)) ? ((_234_v122).get(_235_v123)) : (_101_v2));
          let _236_v124;
          _236_v124 = new _dafny.CodePoint('w'.codePointAt(0));
          let _237_v125;
          _237_v125 = _dafny.Map.Empty.slice().updateUnsafe(_236_v124,_96_v0);
          (_98_globalState).f13 = (_237_v125).contains(_236_v124);
        }
        let _nw37 = Array((new BigNumber(11)).toNumber()).fill(false);
        _96_v0 = _nw37;
        if (_101_v2) {
          _102_v3 = new BigNumber(285);
          let _238_v126;
          _238_v126 = _dafny.Seq.of(_101_v2, _101_v2, _101_v2);
          (_226_v114).m0(_238_v126, (_227_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length))], _98_globalState);
          let _239_v127;
          _239_v127 = _dafny.Set.fromElements(_101_v2, _101_v2);
          let _240_v128;
          _240_v128 = _dafny.Map.Empty.slice().updateUnsafe((_227_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length))],new BigNumber((_239_v127).length));
          let _241_v129;
          _241_v129 = _dafny.MultiSet.fromElements(_101_v2);
          let _242_v130;
          _242_v130 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,(_226_v114).fm3(false, _101_v2, _241_v129, _238_v126, _98_globalState));
          let _243_v131;
          _243_v131 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,(((_240_v128).contains(new BigNumber(577))) ? ((_240_v128).get(new BigNumber(577))) : (new BigNumber((_242_v130).length))));
          let _244_v132;
          _244_v132 = _dafny.Map.Empty.slice().updateUnsafe(_102_v3,_243_v131);
          let _245_v133;
          _245_v133 = _module.D1.create_DC4(false, (_227_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length))], _97_v1);
          let _246_v134;
          _246_v134 = _dafny.Seq.of(_245_v133);
          let _247_v135;
          _247_v135 = _module.D1.create_DC4(_101_v2, new BigNumber((_246_v134).length), _97_v1);
          _244_v132 = (_244_v132).update(((_247_v135).dtor_cf8).minus(_102_v3), _243_v131);
          let _248_v136;
          _248_v136 = _dafny.Set.fromElements(_226_v114);
          let _249_v137;
          _249_v137 = _dafny.Map.Empty.slice().updateUnsafe(_102_v3,_248_v136);
          _249_v137 = (_dafny.Map.Empty.slice().updateUnsafe((_227_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length))],_248_v136)).Merge(_249_v137);
          (_98_globalState).f10 = _101_v2;
        } else {
          let _250_v139;
          _250_v139 = _dafny.Seq.of(new BigNumber((function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of _dafny.IntegerRange(new BigNumber(-420), new BigNumber(960))) {
              let _251_v138 = _compr_10;
              if (((new BigNumber(-420)).isLessThanOrEqualTo(_251_v138)) && ((_251_v138).isLessThan(new BigNumber(960)))) {
                _coll10.push([(_251_v138).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_227_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length))],_101_v2)).length)),_dafny.MultiSet.fromElements(_101_v2, _101_v2, _101_v2, _101_v2, !(_101_v2))]);
              }
            }
            return _coll10;
          }()).length), new BigNumber((_97_v1).length));
          let _252_v140;
          _252_v140 = _module.D1.create_DC3(_250_v139, _101_v2);
          (_98_globalState).f10 = (_252_v140).dtor_cf6;
          let _253_v141;
          let _nw38 = new _module.C0();
          _nw38.__ctor();
          _253_v141 = _nw38;
          (_98_globalState).f15 = false;
          let _254_v142;
          let _init4 = ((_255_v1, _256_v115) => function (_257_i10) {
            return _dafny.Seq.update(_255_v1, _module.__default.safeIndex((_256_v115)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_256_v115).length))], new BigNumber((_255_v1).length)), new _dafny.CodePoint('m'.codePointAt(0)));
          })(_97_v1, _227_v115);
          let _nw39 = Array((new BigNumber(23)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw39.length); _i0_4++) {
            _nw39[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _254_v142 = _nw39;
          let _258_v143;
          _258_v143 = _dafny.Map.Empty.slice().updateUnsafe(!(_101_v2),_97_v1);
          let _index22 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_254_v142).length));
          (_254_v142)[_index22] = (((_258_v143).contains(_101_v2)) ? ((_258_v143).get(_101_v2)) : (_dafny.Seq.UnicodeFromString("pojsssw")));
          let _index23 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_254_v142).length));
          (_254_v142)[_index23] = _97_v1;
          let _259_v144;
          _259_v144 = _module.D1.create_DC2(_102_v3);
          let _index24 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_227_v115).length));
          (_227_v115)[_index24] = (_259_v144).dtor_cf4;
        }
        let _260_v145;
        _260_v145 = _dafny.MultiSet.fromElements(_101_v2);
        let _261_v146;
        _261_v146 = _dafny.Seq.of(_101_v2);
        let _262_v147;
        _262_v147 = _dafny.Map.Empty.slice().updateUnsafe((_226_v114).fm3(_101_v2, _101_v2, _260_v145, _261_v146, _98_globalState),_102_v3);
        _262_v147 = _262_v147;
      } else {
        let _263_v148;
        _263_v148 = _dafny.Seq.of(!(_101_v2));
        (_226_v114).m0(_263_v148, _module.__default.safeDivisionInt(_102_v3, _102_v3), _98_globalState);
        let _index25 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_224_v112).length));
        (_224_v112)[_index25] = _226_v114;
        let _264_v149;
        _264_v149 = _dafny.Map.Empty.slice().updateUnsafe(_97_v1,_226_v114);
        let _index26 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_224_v112).length));
        let _rhs8 = true;
        let _rhs9 = (((_264_v149).contains(_dafny.Seq.UnicodeFromString("ucwu"))) ? ((_264_v149).get(_dafny.Seq.UnicodeFromString("ucwu"))) : (_226_v114));
        let _lhs6 = _98_globalState;
        let _lhs7 = _224_v112;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_224_v112).length));
        _lhs6.f13 = _rhs8;
        _lhs7[_lhs8] = _rhs9;
        let _265_v150;
        _265_v150 = _dafny.Map.Empty.slice().updateUnsafe(_102_v3,new BigNumber((_dafny.Seq.UnicodeFromString("dioapv")).length));
        _265_v150 = (_265_v150).update(_102_v3, _102_v3);
        let _266_v151;
        _266_v151 = new _dafny.CodePoint('v'.codePointAt(0));
        let _rhs10 = _266_v151;
        let _rhs11 = _96_v0;
        _266_v151 = _rhs10;
        _96_v0 = _rhs11;
        (_98_globalState).f10 = (((_102_v3).isLessThan(_102_v3)) ? (!(_101_v2) || ((_263_v148)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_97_v1).length)), new BigNumber((_263_v148).length))])) : (_101_v2));
      }
      let _267_i11;
      _267_i11 = _dafny.ZERO;
      L1: {
        while (_101_v2) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_267_i11)) {
              break L1;
            }
            _267_i11 = (_267_i11).plus(_dafny.ONE);
            let _268_v152;
            let _init5 = ((_269_v2) => function (_270_i12) {
              return _dafny.Seq.Concat(_dafny.Seq.of(_269_v2, _269_v2), _dafny.Seq.of(false));
            })(_101_v2);
            let _nw40 = Array((new BigNumber(17)).toNumber());
            for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw40.length); _i0_5++) {
              _nw40[_i0_5] = _init5(new BigNumber(_i0_5));
            }
            _268_v152 = _nw40;
            let _index27 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_268_v152).length));
            (_268_v152)[_index27] = _dafny.Seq.of(_101_v2);
            let _271_v153;
            _271_v153 = _dafny.Seq.of(_101_v2, !(_101_v2));
            let _index28 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_268_v152).length));
            (_268_v152)[_index28] = _dafny.Seq.Concat(_271_v153, _271_v153);
            let _272_v154;
            _272_v154 = _dafny.Map.Empty.slice().updateUnsafe(_271_v153,new BigNumber((_97_v1).length));
            let _273_v155;
            _273_v155 = _dafny.Set.fromElements(false);
            let _index29 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_268_v152).length));
            let _index30 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_268_v152).length));
            let _rhs12 = _102_v3;
            let _rhs13 = _102_v3;
            let _rhs14 = _271_v153;
            let _rhs15 = _module.__default.fm4(_272_v154, _98_globalState);
            let _rhs16 = (((_dafny.Set.fromElements(true, _101_v2, _101_v2, _101_v2)).IsProperSubsetOf(_273_v155)) ? (_271_v153) : (_271_v153));
            let _lhs9 = _98_globalState;
            let _lhs10 = _268_v152;
            let _lhs11 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_268_v152).length));
            let _lhs12 = _98_globalState;
            let _lhs13 = _268_v152;
            let _lhs14 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_268_v152).length));
            _102_v3 = _rhs12;
            _lhs9.f0 = _rhs13;
            _lhs10[_lhs11] = _rhs14;
            _lhs12.f9 = _rhs15;
            _lhs13[_lhs14] = _rhs16;
            (_98_globalState).f9 = _module.__default.safeDivisionInt(((_101_v2) ? (_102_v3) : (_102_v3)), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_102_v3, _102_v3)));
            (_98_globalState).f2 = _102_v3;
            let _274_v156;
            _274_v156 = _module.D1.create_DC2(_102_v3);
            let _275_v157;
            _275_v157 = _module.D1.create_DC2((_dafny.ZERO).minus((_dafny.ZERO).minus(((_101_v2) ? ((_274_v156).dtor_cf4) : (new BigNumber((_273_v155).length))))));
            let _source3 = _274_v156;
            if (_source3.is_DC3) {
              let _276___mcc_h0 = (_source3).cf5;
              let _277___mcc_h1 = (_source3).cf6;
              let _278_cf6 = _277___mcc_h1;
              let _279_cf5 = _276___mcc_h0;
              (_98_globalState).f10 = !(_278_cf6) || (_dafny.areEqual(_97_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(679)), function (_280_i13) {
                return new _dafny.CodePoint('m'.codePointAt(0));
              })));
              let _281_v158;
              _281_v158 = _dafny.Seq.of(_dafny.Seq.Concat(_97_v1, _dafny.Seq.UnicodeFromString("ijprsydy")), _dafny.Seq.Concat(_97_v1, _dafny.Seq.UnicodeFromString("h")), _dafny.Seq.UnicodeFromString("xceaadkiw"));
              let _282_v159;
              _282_v159 = _dafny.MultiSet.fromElements(_101_v2, _101_v2);
              let _283_v160;
              _283_v160 = _dafny.Map.Empty.slice().updateUnsafe(_102_v3,new _dafny.CodePoint('b'.codePointAt(0)));
              let _284_v161;
              _284_v161 = new _dafny.CodePoint('n'.codePointAt(0));
              let _rhs17 = _281_v158;
              let _rhs18 = _226_v114;
              let _rhs19 = (_102_v3).minus((((_282_v159).contains(_278_cf6)) ? ((_282_v159).get(_278_cf6)) : (_102_v3)));
              let _rhs20 = _dafny.Seq.of((((_283_v160).contains(new BigNumber(-218))) ? ((_283_v160).get(new BigNumber(-218))) : (_284_v161)), _284_v161, _284_v161, _284_v161, _284_v161);
              let _rhs21 = (_102_v3).isLessThanOrEqualTo(_102_v3);
              let _lhs15 = _98_globalState;
              let _lhs16 = _98_globalState;
              _281_v158 = _rhs17;
              _226_v114 = _rhs18;
              _lhs15.f0 = _rhs19;
              _97_v1 = _rhs20;
              _lhs16.f15 = _rhs21;
              let _index31 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_96_v0).length));
              (_96_v0)[_index31] = _278_cf6;
              let _285_v162;
              let _init6 = ((_286_v155) => function (_287_i14) {
                return _286_v155;
              })(_273_v155);
              let _nw41 = Array((new BigNumber(2)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw41.length); _i0_6++) {
                _nw41[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _285_v162 = _nw41;
              let _index32 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_285_v162).length));
              (_285_v162)[_index32] = (_273_v155).Intersect(_dafny.Set.fromElements(_278_cf6, _101_v2));
              let _index33 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_96_v0).length));
              let _index34 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_285_v162).length));
              let _rhs22 = _101_v2;
              let _rhs23 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("rugsxcnc")).length), new BigNumber((_279_cf5).length));
              let _rhs24 = (_273_v155).Difference(_module.__default.fm10(_101_v2, _278_cf6, _102_v3, _284_v161, _98_globalState));
              let _rhs25 = _279_cf5;
              let _lhs17 = _96_v0;
              let _lhs18 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_96_v0).length));
              let _lhs19 = _98_globalState;
              let _lhs20 = _285_v162;
              let _lhs21 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_285_v162).length));
              _lhs17[_lhs18] = _rhs22;
              _lhs19.f2 = _rhs23;
              _lhs20[_lhs21] = _rhs24;
              _279_cf5 = _rhs25;
              let _index35 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_96_v0).length));
              let _rhs26 = _226_v114;
              let _rhs27 = _278_cf6;
              let _rhs28 = (_96_v0)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_96_v0).length))];
              let _lhs22 = _98_globalState;
              let _lhs23 = _96_v0;
              let _lhs24 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_96_v0).length));
              _226_v114 = _rhs26;
              _lhs22.f15 = _rhs27;
              _lhs23[_lhs24] = _rhs28;
            } else if (_source3.is_DC4) {
              let _288___mcc_h2 = (_source3).cf7;
              let _289___mcc_h3 = (_source3).cf8;
              let _290___mcc_h4 = (_source3).cf9;
              let _291_cf9 = _290___mcc_h4;
              let _292_cf8 = _289___mcc_h3;
              let _293_cf7 = _288___mcc_h2;
              let _294_v163;
              _294_v163 = _dafny.MultiSet.fromElements(_module.__default.fm4(_272_v154, _98_globalState));
              let _295_v164;
              _295_v164 = new _dafny.CodePoint('n'.codePointAt(0));
              let _296_v165;
              _296_v165 = _dafny.Map.Empty.slice().updateUnsafe(!(_101_v2),_295_v164);
              let _297_v166;
              _297_v166 = _dafny.MultiSet.fromElements(_101_v2);
              let _298_v167;
              _298_v167 = _dafny.Seq.of((((_294_v163).contains(_102_v3)) ? ((_294_v163).get(_102_v3)) : (new BigNumber((_297_v166).cardinality()))), _292_cf8);
              let _299_v168;
              _299_v168 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_dafny.MultiSet.FromArray(_298_v167));
              let _300_v169;
              _300_v169 = _dafny.Seq.of(new BigNumber((_296_v165).length), _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("mfd")).length)), _102_v3), (_292_cf8).minus(new BigNumber((_299_v168).length)));
              let _rhs29 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_295_v164, _295_v164), _dafny.Seq.update(_291_cf9, _module.__default.safeIndex(_292_cf8, new BigNumber((_291_cf9).length)), _295_v164)), ((!(_101_v2)) ? (_97_v1) : (_97_v1)));
              let _rhs30 = (_300_v169)[_module.__default.safeIndex(_102_v3, new BigNumber((_300_v169).length))];
              let _rhs31 = _292_cf8;
              let _rhs32 = (_dafny.MultiSet.fromElements(_292_cf8, _292_cf8, _102_v3, (_dafny.ZERO).minus(_292_cf8))).Intersect((_294_v163).Intersect(_294_v163));
              let _lhs25 = _98_globalState;
              let _lhs26 = _98_globalState;
              let _lhs27 = _98_globalState;
              _lhs25.f7 = _rhs29;
              _lhs26.f9 = _rhs30;
              _lhs27.f2 = _rhs31;
              _294_v163 = _rhs32;
              let _301_v170;
              let _nw42 = Array((new BigNumber(8)).toNumber());
              _nw42[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_97_v1, _97_v1);
              _nw42[(_dafny.ONE).toNumber()] = _97_v1;
              _nw42[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("ws");
              _nw42[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_97_v1, _291_cf9);
              _nw42[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(741)), ((_302_v164) => function (_303_i15) {
                return _302_v164;
              })(_295_v164)), _291_cf9), _module.__default.safeIndex(_102_v3, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(741)), ((_304_v164) => function (_305_i15) {
                return _304_v164;
              })(_295_v164)), _291_cf9)).length)), _295_v164);
              _nw42[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-212)), function (_306_i16) {
                return new _dafny.CodePoint('f'.codePointAt(0));
              });
              _nw42[(new BigNumber(6)).toNumber()] = _291_cf9;
              _nw42[(new BigNumber(7)).toNumber()] = _97_v1;
              _301_v170 = _nw42;
              let _index36 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_301_v170).length));
              (_301_v170)[_index36] = _291_cf9;
              let _307_v171;
              _307_v171 = _dafny.Seq.of(_297_v166);
              let _pat_let_tv0 = _101_v2;
              let _index37 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_301_v170).length));
              (_301_v170)[_index37] = (((_226_v114).fm3(_293_cf7, _101_v2, (_307_v171)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_307_v171).length))], _271_v153, _98_globalState)) ? (_dafny.Seq.Concat(_291_cf9, _291_cf9)) : (_module.__default.fm5((function (_pat_let0_0) {
                return function (_308_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_309_dt__update_hcf2_h0) {
                      return _module.D0.create_DC1((_308_dt__update__tmp_h0).dtor_cf1, _309_dt__update_hcf2_h0, (_308_dt__update__tmp_h0).dtor_cf3);
                    }(_pat_let1_0);
                  }(_pat_let_tv0);
                }(_pat_let0_0);
              }(_module.D0.create_DC1((_226_v114).fm3(true, _101_v2, _297_v166, _271_v153, _98_globalState), _101_v2, _101_v2))).dtor_cf2, (_226_v114).fm3(!(_293_cf7), _101_v2, _297_v166, (_268_v152)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_268_v152).length))], _98_globalState), _98_globalState)));
              let _310_v172;
              _310_v172 = _dafny.Seq.of(_271_v153);
              _310_v172 = _dafny.Seq.Concat(_310_v172, _310_v172);
              _295_v164 = _295_v164;
            } else if (_source3.is_DC2) {
              let _311___mcc_h5 = (_source3).cf4;
              let _312_cf4 = _311___mcc_h5;
              let _313_v173;
              let _nw43 = new _module.C0();
              _nw43.__ctor();
              _313_v173 = _nw43;
              let _314_v174;
              let _nw44 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
              _314_v174 = _nw44;
              let _315_v175;
              _315_v175 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(671),_312_cf4);
              let _316_v176;
              _316_v176 = _dafny.Set.fromElements(_315_v175, _315_v175, _dafny.Map.Empty.slice().updateUnsafe(_312_cf4,_module.__default.fm4(_272_v154, _98_globalState)));
              let _index38 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_314_v174).length));
              (_314_v174)[_index38] = _316_v176;
              let _index39 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_314_v174).length));
              (_314_v174)[_index39] = _316_v176;
              let _317_v177;
              _317_v177 = _module.D1.create_DC4((_273_v155).IsDisjointFrom(_273_v155), _102_v3, _97_v1);
              _317_v177 = _317_v177;
              (_98_globalState).f9 = _module.__default.fm4(_272_v154, _98_globalState);
            } else {
              let _318___mcc_h6 = (_source3).cf10;
              let _319_cf10 = _318___mcc_h6;
              (_226_v114).m0(_dafny.Seq.of(_101_v2), _module.__default.safeModuloInt(_102_v3, _102_v3), _98_globalState);
              let _320_v178;
              _320_v178 = _dafny.Seq.of(_226_v114, _226_v114);
              let _321_v179;
              _321_v179 = _dafny.Map.Empty.slice().updateUnsafe(_226_v114,_320_v178);
              _321_v179 = (_321_v179).update(_226_v114, ((_101_v2) ? (_320_v178) : (_320_v178)));
              let _322_v180;
              _322_v180 = new _dafny.CodePoint('b'.codePointAt(0));
              let _323_v181;
              _323_v181 = _module.D0.create_DC0(_322_v180);
              let _324_v182;
              _324_v182 = _dafny.Seq.of(_module.D0.create_DC0(_322_v180), _323_v181, _323_v181);
              let _325_v183;
              _325_v183 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_dafny.Seq.of(_102_v3));
              let _326_v184;
              _326_v184 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("jgeiw")).length));
              let _327_v185;
              let _328_v186;
              let _329_v187;
              let _330_v188;
              let _out16;
              let _out17;
              let _out18;
              let _out19;
              let _outcollector4 = _module.__default.m1(_322_v180, ((_101_v2) ? (_dafny.Seq.of(_module.D0.create_DC0(_322_v180), function (_pat_let2_0) {
                return function (_331_dt__update__tmp_h1) {
                  return function (_pat_let3_0) {
                    return function (_332_dt__update_hcf0_h0) {
                      return _module.D0.create_DC0(_332_dt__update_hcf0_h0);
                    }(_pat_let3_0);
                  }(new _dafny.CodePoint('m'.codePointAt(0)));
                }(_pat_let2_0);
              }(_module.D0.create_DC0(_322_v180)), _module.D0.create_DC0(new _dafny.CodePoint('o'.codePointAt(0))), _module.D0.create_DC0(_322_v180))) : (_324_v182)), _101_v2, (((_325_v183).contains(_101_v2)) ? ((_325_v183).get(_101_v2)) : (_326_v184)), _98_globalState);
              _out16 = _outcollector4[0];
              _out17 = _outcollector4[1];
              _out18 = _outcollector4[2];
              _out19 = _outcollector4[3];
              _327_v185 = _out16;
              _328_v186 = _out17;
              _329_v187 = _out18;
              _330_v188 = _out19;
              (_98_globalState).f2 = (new BigNumber(389)).multipliedBy(_module.__default.fm4((_272_v154).update(_271_v153, _102_v3), _98_globalState));
            }
          }
        }
      }
      let _333_v189;
      _333_v189 = new _dafny.CodePoint('p'.codePointAt(0));
      let _334_v190;
      _334_v190 = _dafny.Map.Empty.slice().updateUnsafe(_333_v189,_102_v3);
      let _335_v191;
      _335_v191 = _dafny.Map.Empty.slice().updateUnsafe(false,_101_v2);
      let _336_v192;
      _336_v192 = _dafny.MultiSet.fromElements(_101_v2, _101_v2);
      let _337_v193;
      _337_v193 = _dafny.Seq.of(_101_v2);
      let _338_v194;
      _338_v194 = _dafny.Seq.of((_226_v114).fm3((((_335_v191).contains(_101_v2)) ? ((_335_v191).get(_101_v2)) : (_101_v2)), _101_v2, _336_v192, _337_v193, _98_globalState), false);
      let _339_v195;
      _339_v195 = _dafny.Seq.of(_334_v190, _module.__default.fm11(_dafny.MultiSet.FromArray(_337_v193), _102_v3, _101_v2, _98_globalState));
      let _340_i17;
      _340_i17 = _dafny.ZERO;
      L2: {
        while ((_226_v114).fm3(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("ipj"), _97_v1), _dafny.areEqual(_dafny.Seq.update(_dafny.Seq.update(_339_v195, _module.__default.safeIndex(_102_v3, new BigNumber((_339_v195).length)), _334_v190), _module.__default.safeIndex(_102_v3, new BigNumber((_dafny.Seq.update(_339_v195, _module.__default.safeIndex(_102_v3, new BigNumber((_339_v195).length)), _334_v190)).length)), _334_v190), _dafny.Seq.update(_dafny.Seq.of(_334_v190), _module.__default.safeIndex(_102_v3, new BigNumber((_dafny.Seq.of(_334_v190)).length)), _module.__default.fm11(_dafny.MultiSet.fromElements(_101_v2), _102_v3, _101_v2, _98_globalState))), (_336_v192).Intersect(_dafny.MultiSet.fromElements(_101_v2, _101_v2)), _338_v194, _98_globalState)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_340_i17)) {
              break L2;
            }
            _340_i17 = (_340_i17).plus(_dafny.ONE);
            let _341_v196;
            let _init7 = function (_342_i18) {
              return _dafny.Seq.UnicodeFromString("cixtikg");
            };
            let _nw45 = Array((new BigNumber(27)).toNumber());
            for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw45.length); _i0_7++) {
              _nw45[_i0_7] = _init7(new BigNumber(_i0_7));
            }
            _341_v196 = _nw45;
            let _index40 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_341_v196).length));
            (_341_v196)[_index40] = _97_v1;
            let _index41 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_341_v196).length));
            (_341_v196)[_index41] = _97_v1;
            let _343_v197;
            let _nw46 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
            _343_v197 = _nw46;
            let _344_v198;
            let _nw47 = Array((new BigNumber(22)).toNumber());
            _nw47[(_dafny.ZERO).toNumber()] = _343_v197;
            _nw47[(_dafny.ONE).toNumber()] = _343_v197;
            _nw47[(new BigNumber(2)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(3)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(4)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(5)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(6)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(7)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(8)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(9)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(10)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(11)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(12)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(13)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(14)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(15)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(16)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(17)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(18)).toNumber()] = ((_101_v2) ? (_343_v197) : (_343_v197));
            _nw47[(new BigNumber(19)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(20)).toNumber()] = _343_v197;
            _nw47[(new BigNumber(21)).toNumber()] = _343_v197;
            _344_v198 = _nw47;
            let _index42 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_344_v198).length));
            (_344_v198)[_index42] = _343_v197;
            let _index43 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_344_v198).length));
            (_344_v198)[_index43] = _343_v197;
            _96_v0 = _96_v0;
            let _345_v199;
            let _nw48 = new _module.C0();
            _nw48.__ctor();
            _345_v199 = _nw48;
          }
        }
      }
      let _346_v200;
      _346_v200 = _dafny.Seq.of(_102_v3);
      (_226_v114).m0(_338_v194, (_346_v200)[_module.__default.safeIndex(new BigNumber((_97_v1).length), new BigNumber((_346_v200).length))], _98_globalState);
      _338_v194 = _dafny.Seq.Concat(_dafny.Seq.of(_101_v2, _101_v2, _101_v2, _101_v2), ((_101_v2) ? (_dafny.Seq.of(_101_v2, _101_v2, _101_v2)) : (_338_v194)));
      let _347_i19;
      _347_i19 = _dafny.ZERO;
      L3: {
        while ((!((_226_v114).fm3(_101_v2, _101_v2, _dafny.MultiSet.fromElements(_101_v2, _101_v2, _101_v2), _337_v193, _98_globalState))) && (_101_v2)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_347_i19)) {
              break L3;
            }
            _347_i19 = (_347_i19).plus(_dafny.ONE);
            (_98_globalState).f10 = _101_v2;
            _101_v2 = _101_v2;
            (_98_globalState).f15 = !(_101_v2);
            (_98_globalState).f15 = ((_336_v192).Union(_336_v192)).IsProperSubsetOf((_336_v192).Difference(_336_v192));
          }
        }
      }
      let _348_i20;
      _348_i20 = _dafny.ZERO;
      L4: {
        while ((_101_v2) === (_101_v2)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_348_i20)) {
              break L4;
            }
            _348_i20 = (_348_i20).plus(_dafny.ONE);
            let _349_v201;
            _349_v201 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_336_v192);
            (_98_globalState).f15 = ((((_349_v201).contains(_101_v2)) ? ((_349_v201).get(_101_v2)) : (_dafny.MultiSet.fromElements(_101_v2, _101_v2)))).IsSubsetOf(_336_v192);
            (_98_globalState).f15 = (_102_v3).isLessThanOrEqualTo(_102_v3);
            let _350_v202;
            let _nw49 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
            _350_v202 = _nw49;
            let _351_v203;
            _351_v203 = _dafny.Map.Empty.slice().updateUnsafe(_101_v2,_102_v3);
            let _352_v204;
            _352_v204 = _dafny.Set.fromElements((_351_v203).Merge(_dafny.Map.Empty.slice().updateUnsafe(_101_v2,_102_v3)), (_351_v203).update(_101_v2, _102_v3), _351_v203);
            let _353_v205;
            _353_v205 = _dafny.Map.Empty.slice().updateUnsafe(_97_v1,_97_v1);
            let _rhs33 = _226_v114;
            let _rhs34 = _350_v202;
            let _rhs35 = !((_353_v205).update(_97_v1, _module.__default.fm5(_101_v2, _101_v2, _98_globalState))).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-480)), ((_354_v189) => function (_355_i21) {
              return _354_v189;
            })(_333_v189)));
            let _rhs36 = _352_v204;
            let _lhs28 = _98_globalState;
            _226_v114 = _rhs33;
            _350_v202 = _rhs34;
            _lhs28.f10 = _rhs35;
            _352_v204 = _rhs36;
            if ((_102_v3).isLessThanOrEqualTo(_102_v3)) {
              (_98_globalState).f9 = new BigNumber((_dafny.Seq.update(_97_v1, _module.__default.safeIndex(new BigNumber(793), new BigNumber((_97_v1).length)), _333_v189)).length);
              let _356_v206;
              let _nw50 = Array((new BigNumber(5)).toNumber()).fill([]);
              _356_v206 = _nw50;
              let _index44 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_356_v206).length));
              (_356_v206)[_index44] = _350_v202;
              let _index45 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_356_v206).length));
              (_356_v206)[_index45] = _350_v202;
              let _357_v207;
              _357_v207 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(72),_102_v3);
              let _358_v208;
              _358_v208 = _dafny.Seq.of(_357_v207);
              let _359_v209;
              _359_v209 = _dafny.MultiSet.fromElements(_102_v3, (new BigNumber(((_358_v208)[_module.__default.safeIndex(_102_v3, new BigNumber((_358_v208).length))]).length)).plus(_102_v3), (_102_v3).plus(_102_v3), _module.__default.safeModuloInt(_102_v3, _102_v3), _102_v3);
              let _360_v211;
              _360_v211 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_359_v209).cardinality()),(_226_v114).fm3(_101_v2, _101_v2, _dafny.MultiSet.fromElements(_101_v2, _101_v2), _337_v193, _98_globalState));
              let _361_v212;
              _361_v212 = _dafny.Map.Empty.slice().updateUnsafe(_360_v211,(_module.D1.create_DC4(_101_v2, _102_v3, _97_v1)).dtor_cf7);
              let _362_v215;
              _362_v215 = _dafny.Set.fromElements(_360_v211);
              let _363_v216;
              _363_v216 = _dafny.Seq.of(_362_v215);
              let _364_v217;
              _364_v217 = _dafny.Seq.of(function () {
                let _coll11 = new _dafny.Map();
                for (const _compr_11 of ((_363_v216)[_module.__default.safeIndex(_102_v3, new BigNumber((_363_v216).length))]).Elements) {
                  let _365_v214 = _compr_11;
                  if (((_363_v216)[_module.__default.safeIndex(_102_v3, new BigNumber((_363_v216).length))]).contains(_365_v214)) {
                    _coll11.push([_365_v214,_101_v2]);
                  }
                }
                return _coll11;
              }());
              let _366_v218;
              _366_v218 = _dafny.Set.fromElements(_101_v2);
              let _367_v219;
              _367_v219 = _module.D4.create_DC8(_361_v212);
              let _368_v220;
              let _nw51 = Array((new BigNumber(20)).toNumber());
              _nw51[(_dafny.ZERO).toNumber()] = (function () {
                let _coll12 = new _dafny.Map();
                for (const _compr_12 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-660)), ((_369_v3, _370_v2) => function (_371_i22) {
                  return _dafny.Map.Empty.slice().updateUnsafe(_369_v3,_370_v2);
                })(_102_v3, _101_v2))).Elements) {
                  let _372_v210 = _compr_12;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-660)), ((_373_v3, _374_v2) => function (_371_i22) {
                    return _dafny.Map.Empty.slice().updateUnsafe(_373_v3,_374_v2);
                  })(_102_v3, _101_v2)), _372_v210)) {
                    _coll12.push([_372_v210,_101_v2]);
                  }
                }
                return _coll12;
              }()).Merge(_361_v212);
              _nw51[(_dafny.ONE).toNumber()] = _361_v212;
              _nw51[(new BigNumber(2)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(3)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(4)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(5)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(6)).toNumber()] = (_361_v212).update(function () {
                let _coll13 = new _dafny.Map();
                for (const _compr_13 of (_360_v211).Keys.Elements) {
                  let _375_v213 = _compr_13;
                  if ((_360_v211).contains(_375_v213)) {
                    _coll13.push([(_375_v213).plus(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length)),true]);
                  }
                }
                return _coll13;
              }(), _101_v2);
              _nw51[(new BigNumber(7)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(8)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(9)).toNumber()] = (_361_v212).update(_360_v211, _101_v2);
              _nw51[(new BigNumber(10)).toNumber()] = (_361_v212).update(_360_v211, !(_101_v2));
              _nw51[(new BigNumber(11)).toNumber()] = ((_364_v217)[_module.__default.safeIndex(_102_v3, new BigNumber((_364_v217).length))]).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_102_v3,_101_v2)).update(_102_v3, _101_v2),(_226_v114).fm3(_101_v2, _101_v2, _dafny.MultiSet.fromElements(_101_v2), _338_v194, _98_globalState)));
              _nw51[(new BigNumber(12)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(13)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(14)).toNumber()] = (_361_v212).Merge(_361_v212);
              _nw51[(new BigNumber(15)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_366_v218).length)),false),(((_360_v211).contains(new BigNumber((_335_v191).length))) ? ((_360_v211).get(new BigNumber((_335_v191).length))) : (true)));
              _nw51[(new BigNumber(17)).toNumber()] = _361_v212;
              _nw51[(new BigNumber(18)).toNumber()] = (_367_v219).dtor_cf13;
              _nw51[(new BigNumber(19)).toNumber()] = _361_v212;
              _368_v220 = _nw51;
              let _index46 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_368_v220).length));
              (_368_v220)[_index46] = (_dafny.Map.Empty.slice().updateUnsafe(_360_v211,_101_v2)).Merge((_361_v212).update(function () {
                let _coll14 = new _dafny.Map();
                for (const _compr_14 of _dafny.IntegerRange(new BigNumber(-640), new BigNumber(23))) {
                  let _376_v221 = _compr_14;
                  if (((new BigNumber(-640)).isLessThanOrEqualTo(_376_v221)) && ((_376_v221).isLessThan(new BigNumber(23)))) {
                    _coll14.push([(_376_v221).minus(new BigNumber((_351_v203).length)),_101_v2]);
                  }
                }
                return _coll14;
              }(), _101_v2));
              let _index47 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_368_v220).length));
              let _rhs37 = _dafny.MultiSet.fromElements((_102_v3).minus((_dafny.ZERO).minus((((_351_v203).contains(_101_v2)) ? ((_351_v203).get(_101_v2)) : (_102_v3)))));
              let _rhs38 = _226_v114;
              let _rhs39 = ((!(_101_v2)) ? (_361_v212) : (((_364_v217)[_module.__default.safeIndex(_102_v3, new BigNumber((_364_v217).length))]).Merge(function () {
                let _coll15 = new _dafny.Map();
                for (const _compr_15 of (_dafny.Set.fromElements(_360_v211)).Elements) {
                  let _377_v222 = _compr_15;
                  if ((_dafny.Set.fromElements(_360_v211)).contains(_377_v222)) {
                    _coll15.push([_377_v222,false]);
                  }
                }
                return _coll15;
              }())));
              let _rhs40 = _226_v114;
              let _lhs29 = _368_v220;
              let _lhs30 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_368_v220).length));
              _359_v209 = _rhs37;
              _226_v114 = _rhs38;
              _lhs29[_lhs30] = _rhs39;
              _226_v114 = _rhs40;
              let _378_v223;
              let _nw52 = Array((new BigNumber(3)).toNumber()).fill(_module.D1.Default());
              _378_v223 = _nw52;
              let _index48 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_96_v0).length));
              (_96_v0)[_index48] = _101_v2;
              let _index49 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_96_v0).length));
              let _rhs41 = (_dafny.ZERO).minus((_102_v3).multipliedBy(_102_v3));
              let _rhs42 = new _dafny.CodePoint('v'.codePointAt(0));
              let _rhs43 = _378_v223;
              let _rhs44 = (_101_v2) && ((_337_v193)[_module.__default.safeIndex(_102_v3, new BigNumber((_337_v193).length))]);
              let _lhs31 = _98_globalState;
              let _lhs32 = _96_v0;
              let _lhs33 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_96_v0).length));
              _lhs31.f2 = _rhs41;
              _333_v189 = _rhs42;
              _378_v223 = _rhs43;
              _lhs32[_lhs33] = _rhs44;
              let _379_v224;
              let _380_v225;
              let _381_v226;
              let _382_v227;
              let _out20;
              let _out21;
              let _out22;
              let _out23;
              let _outcollector5 = _module.__default.m1(_333_v189, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-33)), ((_383_v189) => function (_384_i23) {
                return _module.D0.create_DC0(_383_v189);
              })(_333_v189)), _101_v2, _346_v200, _98_globalState);
              _out20 = _outcollector5[0];
              _out21 = _outcollector5[1];
              _out22 = _outcollector5[2];
              _out23 = _outcollector5[3];
              _379_v224 = _out20;
              _380_v225 = _out21;
              _381_v226 = _out22;
              _382_v227 = _out23;
            } else {
              (_226_v114).m0(_dafny.Seq.Concat(_dafny.Seq.of(_101_v2), _338_v194), new BigNumber((_346_v200).length), _98_globalState);
              _226_v114 = _226_v114;
              let _385_v228;
              let _nw53 = new _module.C0();
              _nw53.__ctor();
              _385_v228 = _nw53;
              let _386_v229;
              _386_v229 = _dafny.Map.Empty.slice().updateUnsafe(_102_v3,_dafny.Seq.UnicodeFromString("emrwvbq"));
              let _387_v230;
              _387_v230 = _dafny.MultiSet.fromElements(new BigNumber((_346_v200).length));
              let _388_v231;
              _388_v231 = _dafny.MultiSet.fromElements(_387_v230, _387_v230);
              (_98_globalState).f7 = _dafny.Seq.Concat(_dafny.Seq.Concat((((_386_v229).contains(new BigNumber((_388_v231).cardinality()))) ? ((_386_v229).get(new BigNumber((_388_v231).cardinality()))) : (_dafny.Seq.UnicodeFromString("skwwff"))), _module.__default.fm5(_101_v2, _101_v2, _98_globalState)), _97_v1);
              (_98_globalState).f2 = _module.__default.safeModuloInt(_102_v3, (_102_v3).multipliedBy(_102_v3));
            }
          }
        }
      }
      _102_v3 = _102_v3;
      let _389_v232;
      let _init8 = ((_390_v3) => function (_391_i24) {
        return _module.__default.safeModuloInt(_391_i24, (_dafny.ZERO).minus(_390_v3));
      })(_102_v3);
      let _nw54 = Array((new BigNumber(20)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw54.length); _i0_8++) {
        _nw54[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _389_v232 = _nw54;
      _389_v232 = _389_v232;
      let _rhs45 = _101_v2;
      let _rhs46 = _dafny.Seq.Concat(_97_v1, _dafny.Seq.update(_97_v1, _module.__default.safeIndex(_102_v3, new BigNumber((_97_v1).length)), _333_v189));
      let _lhs34 = _98_globalState;
      _101_v2 = _rhs45;
      _lhs34.f7 = _rhs46;
      _101_v2 = (_336_v192).IsSubsetOf(_dafny.MultiSet.fromElements(_101_v2));
      _389_v232 = _389_v232;
      process.stdout.write(_dafny.toString((_96_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_97_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_98_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_98_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_98_globalState).f4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_98_globalState).f4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_98_globalState).f6, _dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_98_globalState.f7).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_98_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_98_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_98_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_98_globalState).f14).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_98_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_101_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_225_v113).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_267_i11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_333_v189));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_334_v190).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber(890)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v191).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_336_v192).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_337_v193, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_338_v194, _dafny.Seq.of(false, false, false, false, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_339_v195, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber(890)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),new BigNumber(2)).updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(46)).updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),_dafny.ONE).updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_340_i17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_346_v200, _dafny.Seq.of(new BigNumber(890)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_347_i19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_348_i20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v232)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
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
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, false, false);
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
    static create_DC3(cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC4(cf7, cf8, cf9) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC5(cf10) {
      let $dt = new D1(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + this.cf9.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.Seq.of(), false);
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
    static create_DC6(cf11) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf12) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
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
    static create_DC9(cf14, cf15) {
      let $dt = new D4(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC10() {
      let $dt = new D4(1);
      return $dt;
    }
    static create_DC8(cf13) {
      let $dt = new D4(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10";
      } else if (this.$tag === 2) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC9(_dafny.ZERO, false);
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
    static create_DC12(cf17) {
      let $dt = new D5(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC13(cf18, cf19, cf20, cf21) {
      let $dt = new D5(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf16) {
      let $dt = new D5(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC14(cf22) {
      let $dt = new D5(3);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f2 = _dafny.ZERO;
      this.f7 = _dafny.Seq.UnicodeFromString("");
      this.f9 = _dafny.ZERO;
      this.f10 = false;
      this.f13 = false;
      this.f15 = false;
      this._f1 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f3 = _dafny.ZERO;
      this._f4 = [];
      this._f5 = false;
      this._f6 = _dafny.Seq.UnicodeFromString("");
      this._f8 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f11 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f12 = _dafny.ZERO;
      this._f14 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f3() {
      let _this = this;
      return _this._f3;
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
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return true;
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let _392_v0;
      _392_v0 = _module.D1.create_DC2(p1);
      let _393_v1;
      _393_v1 = true;
      let _394_v2;
      _394_v2 = _dafny.Map.Empty.slice().updateUnsafe(_393_v1,new BigNumber((_dafny.Seq.UnicodeFromString("snasecdf")).length));
      let _395_v3;
      _395_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(385),_dafny.Map.Empty.slice().updateUnsafe(false,p1));
      let _396_v4;
      _396_v4 = _dafny.Seq.of(_394_v2, (((_395_v3).contains(new BigNumber((p0).length))) ? ((_395_v3).get(new BigNumber((p0).length))) : (_394_v2)));
      let _397_v5;
      let _nw55 = Array((new BigNumber(14)).toNumber());
      _nw55[(_dafny.ZERO).toNumber()] = p1;
      _nw55[(_dafny.ONE).toNumber()] = p1;
      _nw55[(new BigNumber(2)).toNumber()] = p1;
      _nw55[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_392_v0).dtor_cf4);
      _nw55[(new BigNumber(4)).toNumber()] = new BigNumber((_396_v4).length);
      _nw55[(new BigNumber(5)).toNumber()] = p1;
      _nw55[(new BigNumber(6)).toNumber()] = (_392_v0).dtor_cf4;
      _nw55[(new BigNumber(7)).toNumber()] = p1;
      _nw55[(new BigNumber(8)).toNumber()] = p1;
      _nw55[(new BigNumber(9)).toNumber()] = p1;
      _nw55[(new BigNumber(10)).toNumber()] = p1;
      _nw55[(new BigNumber(11)).toNumber()] = (new BigNumber((p0).length)).multipliedBy(p1);
      _nw55[(new BigNumber(12)).toNumber()] = p1;
      _nw55[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(p1), p1);
      _397_v5 = _nw55;
      let _index50 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length));
      (_397_v5)[_index50] = (p1).multipliedBy(p1);
      let _398_v6;
      _398_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _index51 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length));
      (_397_v5)[_index51] = _module.__default.fm4(_398_v6, globalState);
      let _399_v7;
      _399_v7 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))]),_393_v1);
      let _400_v8;
      _400_v8 = _dafny.Set.fromElements(_399_v7);
      let _401_v9;
      _401_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_400_v8).length),p1);
      let _402_v10;
      _402_v10 = _dafny.Set.fromElements(_393_v1);
      let _403_v11;
      _403_v11 = _dafny.Seq.of(_402_v10, _402_v10);
      _401_v9 = (_401_v9).update(_module.__default.safeModuloInt((_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))], new BigNumber((_dafny.Seq.UnicodeFromString("t")).length)), (_dafny.ZERO).minus(new BigNumber((_403_v11).length)));
      let _404_v12;
      _404_v12 = _dafny.Map.Empty.slice().updateUnsafe(_393_v1,_397_v5);
      let _index52 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length));
      (_397_v5)[_index52] = (new BigNumber(((_404_v12).Merge(_404_v12)).length)).minus(new BigNumber(800));
      let _405_v13;
      let _nw56 = Array((new BigNumber(15)).toNumber()).fill(false);
      _405_v13 = _nw56;
      let _index53 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_405_v13).length));
      (_405_v13)[_index53] = true;
      let _index54 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_405_v13).length));
      (_405_v13)[_index54] = _393_v1;
      let _hi1 = (_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))];
      for (let _406_i0 = p1; _406_i0.isLessThan(_hi1); _406_i0 = _406_i0.plus(_dafny.ONE)) {
        (globalState).f0 = _module.__default.safeModuloInt(((_393_v1) ? (p1) : ((_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))])), (_dafny.ZERO).minus((_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))]));
        let _407_v14;
        _407_v14 = _dafny.Seq.UnicodeFromString("cshphu");
        let _index55 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_405_v13).length));
        let _index56 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length));
        let _rhs47 = (((_401_v9).contains((new BigNumber((_402_v10).length)).multipliedBy((_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))]))) ? ((_401_v9).get((new BigNumber((_402_v10).length)).multipliedBy((_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))]))) : (_module.__default.safeModuloInt(_406_i0, (_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))])));
        let _rhs48 = _dafny.areEqual(_407_v14, _407_v14);
        let _rhs49 = new BigNumber(686);
        let _lhs35 = globalState;
        let _lhs36 = _405_v13;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_405_v13).length));
        let _lhs38 = _397_v5;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length));
        _lhs35.f9 = _rhs47;
        _lhs36[_lhs37] = _rhs48;
        _lhs38[_lhs39] = _rhs49;
        (globalState).f13 = false;
        (globalState).f0 = new BigNumber(-981);
      }
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_397_v5).length))) {
        let _408_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_408_i1)) && ((_408_i1).isLessThan(new BigNumber((_397_v5).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_397_v5, (_408_i1).toNumber(), (_408_i1).multipliedBy(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(381)), function (_409_i2) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          })).length), (_397_v5)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_397_v5).length))]))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      return;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
